package zcncore

import (
	"context"

	"github.com/0chain/gosdk/core/util"
)

func GetInfoFromSharders(urlSuffix string, op int, cb GetInfoCallback) {

	tq, err := NewTransactionQuery(util.Shuffle(_config.chain.Sharders))
	if err != nil {
		cb.OnInfoAvailable(op, StatusError, "", err.Error())
		return
	}

	qr, err := tq.GetInfo(context.TODO(), urlSuffix)
	if err != nil {
		cb.OnInfoAvailable(op, StatusError, "", err.Error())
		return
	}

	cb.OnInfoAvailable(op, StatusSuccess, string(qr.Content), "")
}

func GetInfoFromAnySharder(urlSuffix string, op int, cb GetInfoCallback) {

	tq, err := NewTransactionQuery(util.Shuffle(_config.chain.Sharders))
	if err != nil {
		cb.OnInfoAvailable(op, StatusError, "", err.Error())
		return
	}

	qr, err := tq.FromAny(context.TODO(), urlSuffix)
	if err != nil {
		cb.OnInfoAvailable(op, StatusError, "", err.Error())
		return
	}

	cb.OnInfoAvailable(op, StatusSuccess, string(qr.Content), "")
}

func GetEvents(cb GetInfoCallback, filters map[string]string) (err error) {
	if err = CheckConfig(); err != nil {
		return
	}
	go GetInfoFromSharders(WithParams(GET_MINERSC_EVENTS, Params{
		"block_number": filters["block_number"],
		"tx_hash":      filters["tx_hash"],
		"type":         filters["type"],
		"tag":          filters["tag"],
	}), 0, cb)
	return
}

func WithParams(uri string, params Params) string {
	return withParams(uri, params)
}
