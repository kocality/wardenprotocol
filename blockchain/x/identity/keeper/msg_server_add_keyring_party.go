package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"gitlab.qredo.com/qrdochain/fusionchain/x/identity/types"
)

func (k msgServer) AddKeyringParty(goCtx context.Context, msg *types.MsgAddKeyringParty) (*types.MsgAddKeyringPartyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	kr, found := k.KeyringsRepo().Get(ctx, msg.KeyringId)
	if !found {
		return nil, fmt.Errorf("keyring not found")
	}

	if kr.IsParty(msg.Party) {
		return nil, fmt.Errorf("party is already a party of the keyring")
	}

	kr.AddParty(msg.Party)
	k.KeyringsRepo().Set(ctx, kr)

	return &types.MsgAddKeyringPartyResponse{}, nil
}