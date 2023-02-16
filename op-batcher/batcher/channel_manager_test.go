package batcher_test

import (
	"math/big"
	"testing"

	"github.com/ethereum-optimism/optimism/op-batcher/batcher"
	"github.com/ethereum-optimism/optimism/op-node/testlog"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/stretchr/testify/require"
)

func TestChannelManagerReturnsErrReorg(t *testing.T) {
	log := testlog.Logger(t, log.LvlTrace)
	m := batcher.NewChannelManager(log, batcher.ChannelConfig{})

	a := types.NewBlock(&types.Header{
		Number: big.NewInt(0),
	}, nil, nil, nil, nil)
	b := types.NewBlock(&types.Header{
		Number:     big.NewInt(1),
		ParentHash: a.Hash(),
	}, nil, nil, nil, nil)
	c := types.NewBlock(&types.Header{
		Number:     big.NewInt(2),
		ParentHash: b.Hash(),
	}, nil, nil, nil, nil)
	x := types.NewBlock(&types.Header{
		Number:     big.NewInt(2),
		ParentHash: common.Hash{0xff},
	}, nil, nil, nil, nil)

	err := m.AddL2Block(a)
	require.Nil(t, err)
	err = m.AddL2Block(b)
	require.Nil(t, err)
	err = m.AddL2Block(c)
	require.Nil(t, err)
	err = m.AddL2Block(x)
	require.ErrorIs(t, err, batcher.ErrReorg)
}
