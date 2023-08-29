package keeper

import (
	"github.com/lightmos/empyrean/x/restaking/types"
)

var _ types.QueryServer = Keeper{}
