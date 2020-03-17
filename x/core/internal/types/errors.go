package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrAssetDoesNotExist = sdkerrors.Register(ModuleName, 1, "asset does not exist")
)
