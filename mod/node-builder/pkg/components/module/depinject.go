// SPDX-License-Identifier: MIT
//
// Copyright (c) 2024 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following
// conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

package beacon

import (
	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/depinject"
	"cosmossdk.io/depinject/appconfig"
	"github.com/berachain/beacon-kit/mod/consensus-types/pkg/types"
	dastore "github.com/berachain/beacon-kit/mod/da/pkg/store"
	engineclient "github.com/berachain/beacon-kit/mod/execution/pkg/client"
	"github.com/berachain/beacon-kit/mod/node-builder/pkg/components"
	"github.com/berachain/beacon-kit/mod/node-builder/pkg/components/metrics"
	modulev1alpha1 "github.com/berachain/beacon-kit/mod/node-builder/pkg/components/module/api/module/v1alpha1"
	"github.com/berachain/beacon-kit/mod/node-builder/pkg/components/storage"
	"github.com/berachain/beacon-kit/mod/node-builder/pkg/config"
	"github.com/berachain/beacon-kit/mod/primitives"
	engineprimitives "github.com/berachain/beacon-kit/mod/primitives-engine"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/crypto"
	"github.com/berachain/beacon-kit/mod/state-transition/pkg/core"
	"github.com/berachain/beacon-kit/mod/storage/pkg/beacondb"
	depositdb "github.com/berachain/beacon-kit/mod/storage/pkg/deposit"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	gokzg4844 "github.com/crate-crypto/go-kzg-4844"
)

// TODO: we don't allow generics here? Why? Is it fixable?
//
//nolint:gochecknoinits // required by sdk.
func init() {
	appconfig.RegisterModule(&modulev1alpha1.Module{},
		appconfig.Provide(ProvideModule),
	)
}

// DepInjectInput is the input for the dep inject framework.
type DepInjectInput struct {
	depinject.In

	// Cosmos components
	AppOpts     servertypes.AppOptions
	Environment appmodule.Environment

	// BeaconKit components
	AvailabilityStore *dastore.Store[types.BeaconBlockBody]
	BeaconConfig      *config.Config
	ChainSpec         primitives.ChainSpec
	DepositStore      *depositdb.KVStore
	EngineClient      *engineclient.EngineClient[*types.ExecutableDataDeneb]
	KzgTrustedSetup   *gokzg4844.JSONTrustedSetup
	Signer            crypto.BLSSigner
	TelemetrySink     *metrics.TelemetrySink
}

// DepInjectOutput is the output for the dep inject framework.
type DepInjectOutput struct {
	depinject.Out
	Module appmodule.AppModule
}

// ProvideModule is a function that provides the module to the application.
func ProvideModule(in DepInjectInput) (DepInjectOutput, error) {
	storageBackend := storage.NewBackend[
		*dastore.Store[types.BeaconBlockBody], core.BeaconState[*types.Validator],
	](
		in.ChainSpec,
		in.AvailabilityStore,
		beacondb.New[
			*types.Fork,
			*types.BeaconBlockHeader,
			engineprimitives.ExecutionPayloadHeader,
			*types.Eth1Data,
			*types.Validator,
		](in.Environment.KVStoreService, DenebPayloadFactory),
		in.DepositStore,
	)

	// TODO: this is hood as fuck.
	if in.BeaconConfig.KZG.Implementation == "" {
		in.BeaconConfig.KZG.Implementation = "crate-crypto/go-kzg-4844"
	}

	runtime, err := components.ProvideRuntime(
		in.BeaconConfig,
		in.ChainSpec,
		in.Signer,
		in.EngineClient,
		in.KzgTrustedSetup,
		storageBackend,
		in.TelemetrySink,
		in.Environment.Logger.With("module", "beacon-kit"),
	)
	if err != nil {
		return DepInjectOutput{}, err
	}

	return DepInjectOutput{
		Module: NewAppModule(runtime),
	}, nil
}

// TODO: move this.
func DenebPayloadFactory() engineprimitives.ExecutionPayloadHeader {
	return &types.ExecutionPayloadHeaderDeneb{}
}