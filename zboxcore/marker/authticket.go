package marker

import (
	"fmt"

	"github.com/0chain/gosdk/core/encryption"
	"github.com/0chain/gosdk/zboxcore/client"
)

// AuthTicket authentication ticket for file sharing
type AuthTicket struct {
	ClientID        string `json:"client_id"`
	OwnerID         string `json:"owner_id"`
	AllocationID    string `json:"allocation_id"`
	FilePathHash    string `json:"file_path_hash"`
	ActualFileHash  string `json:"actual_file_hash"`
	FileName        string `json:"file_name"`
	RefType         string `json:"reference_type"`
	Expiration      int64  `json:"expiration"`
	Timestamp       int64  `json:"timestamp"`
	ReEncryptionKey string `json:"re_encryption_key,omitempty"`
	Encrypted       bool   `json:"encrypted"`
	Signature       string `json:"signature"`
}

// NewAuthTicket returns the MPT hash of the AuthTicket
func (at *AuthTicket) GetHashData() string {
	hashData := fmt.Sprintf("%v:%v:%v:%v:%v:%v:%v:%v:%v:%v:%v",
		at.AllocationID,
		at.ClientID,
		at.OwnerID,
		at.FilePathHash,
		at.FileName,
		at.RefType,
		at.ReEncryptionKey,
		at.Expiration,
		at.Timestamp,
		at.ActualFileHash,
		at.Encrypted,
	)
	return hashData
}

// Sign signs the AuthTicket
func (at *AuthTicket) Sign() error {
	var err error
	hash := encryption.Hash(at.GetHashData())
	at.Signature, err = client.Sign(hash)
	return err
}
