package keeper

import (
	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"

	connectionv7 "github.com/cosmos/ibc-go/v8/modules/core/03-connection/migrations/v7"
	ibcerrors "github.com/cosmos/ibc-go/v8/modules/core/errors"
)

// Migrator is a struct for handling in-place store migrations.
type Migrator struct {
	keeper Keeper
}

// NewMigrator returns a new Migrator.
func NewMigrator(keeper Keeper) Migrator {
	return Migrator{keeper: keeper}
}

// Migrate3to4 migrates from version 3 to 4.
// This migration writes the sentinel localhost connection end to state.
func (m Migrator) Migrate3to4(ctx sdk.Context) error {
	connectionv7.MigrateLocalhostConnection(ctx, m.keeper)
	return nil
}

// MigrateParams migrates from consensus version 4 to 5.
// This migration takes the parameters that are currently stored and managed by x/params
// and stores them directly in the ibc module's state.
func (m Migrator) MigrateParams(ctx sdk.Context) error {
	return errorsmod.Wrap(ibcerrors.ErrInvalidVersion, "must migrate to ibc-go v8.x first")
}
