package sedachain_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "seda-chain/testutil/keeper"
	"seda-chain/testutil/nullify"
	"seda-chain/x/sedachain"
	"seda-chain/x/sedachain/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SedachainKeeper(t)
	sedachain.InitGenesis(ctx, *k, genesisState)
	got := sedachain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
