// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"github.com/google/wire"
	"go.uber.org/zap"
	"source.quilibrium.com/quilibrium/monorepo/node/config"
	"source.quilibrium.com/quilibrium/monorepo/node/consensus"
	"source.quilibrium.com/quilibrium/monorepo/node/consensus/master"
	"source.quilibrium.com/quilibrium/monorepo/node/execution/nop"
	"source.quilibrium.com/quilibrium/monorepo/node/keys"
	"source.quilibrium.com/quilibrium/monorepo/node/p2p"
)

// Injectors from wire.go:

func NewNode(configConfig *config.Config) (*Node, error) {
	zapLogger := logger()
	nopExecutionEngine := nop.NewNopExecutionEngine(zapLogger)
	engineConfig := configConfig.Engine
	keyConfig := configConfig.Key
	fileKeyManager := keys.NewFileKeyManager(keyConfig, zapLogger)
	p2PConfig := configConfig.P2P
	blossomSub := p2p.NewBlossomSub(p2PConfig, zapLogger)
	masterClockConsensusEngine := master.NewMasterClockConsensusEngine(engineConfig, zapLogger, fileKeyManager, blossomSub)
	node, err := newNode(nopExecutionEngine, masterClockConsensusEngine)
	if err != nil {
		return nil, err
	}
	return node, nil
}

// wire.go:

func logger() *zap.Logger {
	log, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	return log
}

var loggerSet = wire.NewSet(
	logger,
)

var keyManagerSet = wire.NewSet(wire.FieldsOf(new(*config.Config), "Key"), keys.NewFileKeyManager, wire.Bind(new(keys.KeyManager), new(*keys.FileKeyManager)))

var pubSubSet = wire.NewSet(wire.FieldsOf(new(*config.Config), "P2P"), p2p.NewBlossomSub, wire.Bind(new(p2p.PubSub), new(*p2p.BlossomSub)))

var engineSet = wire.NewSet(nop.NewNopExecutionEngine)

var consensusSet = wire.NewSet(wire.FieldsOf(new(*config.Config), "Engine"), master.NewMasterClockConsensusEngine, wire.Bind(new(consensus.ConsensusEngine), new(*master.MasterClockConsensusEngine)))
