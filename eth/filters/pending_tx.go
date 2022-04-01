package filters

import (
	"encoding/hex"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
)

type pendingTx struct {
	Hash     string    `json:"hash"`
	From     string    `json:"from"`
	To       string    `json:"to"`
	Data     string    `json:"data"`
	GasPrice uint64    `json:"gasPrice"`
	Gas      uint64    `json:"gas"`
	Time     time.Time `json:"time"`
	Value    *big.Int  `json:"value"`
	Nonce    uint64    `json:"nonce"`
	// AnnounceTime time.Time `json:"announceTime"`
	// Peer         string    `json:"peer"`
}

func makePendingTx(tx *types.Transaction) *pendingTx {
	ptx := pendingTx{Hash: tx.Hash().String()}
	if tx.Data() != nil {
		ptx.Data = hex.EncodeToString(tx.Data())
	}
	if tx.To() != nil {
		ptx.To = tx.To().String()
	}
	if tx.GasPrice() != nil {
		ptx.GasPrice = tx.GasPrice().Uint64()
	}

	s := types.NewEIP2930Signer(tx.ChainId())
	if from, err := s.Sender(tx); err != nil {
		//set from to some default value
		ptx.From = "0x0000000000000000000000000000000000000000"
	} else {
		ptx.From = from.String()
	}

	// ptx.Peer = tx.PeerID
	// ptx.AnnounceTime = tx.AnnounceTIme
	ptx.Time = tx.Time()
	ptx.Value = tx.Value()
	ptx.Nonce = tx.Nonce()
	ptx.Gas = tx.Gas()

	return &ptx //& means a pointer to the struct
}
