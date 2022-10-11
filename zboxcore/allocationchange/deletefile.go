package allocationchange

import (
	"path"

	"github.com/0chain/errors"
	"github.com/0chain/gosdk/core/common"
	"github.com/0chain/gosdk/zboxcore/fileref"
)

type DeleteFileChange struct {
	change
	ObjectTree fileref.RefEntity
}

func (ch *DeleteFileChange) ProcessChange(
	rootRef *fileref.Ref, _ int64) (
	map[string]int64, int64, error) {

	if ch.ObjectTree.GetPath() == "/" {
		rootRef.Children = nil
		rootRef.CalculateHash()
		return nil, 0, nil
	}

	parentPath := path.Dir(ch.ObjectTree.GetPath())

	fields, err := common.GetPathFields(parentPath)
	if err != nil {
		return nil, 0, err
	}

	dirRef := rootRef
	for _, name := range fields {
		found := false
		for _, child := range dirRef.Children {
			if child.GetName() == name {
				dirRef = child.(*fileref.Ref)
				found = true
				break
			}
		}

		if !found {
			return nil, 0, errors.New("invalid_reference_path", "Invalid reference path from the blobber")
		}
	}

	for i, child := range dirRef.Children {
		if child.GetName() == ch.ObjectTree.GetName() {
			dirRef.RemoveChild(i)
			rootRef.CalculateHash()
			return nil, 0, nil
		}
	}
	return nil, 0, errors.New("file_not_found", "File to delete not found in blobber")
}

func (n *DeleteFileChange) GetAffectedPath() string {
	if n.ObjectTree != nil {
		return n.ObjectTree.GetPath()
	}
	return ""
}

func (n *DeleteFileChange) GetSize() int64 {
	if n.ObjectTree != nil {
		return 0 - n.ObjectTree.GetSize()
	}
	return int64(0)
}
