package keeper_test

import (
	"testing"

	testkeeper "github.com/lightmos/empyrean/testutil/keeper"
	"github.com/lightmos/empyrean/x/restaking/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.RestakingKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
