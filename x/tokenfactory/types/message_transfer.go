package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgTransfer{}

func NewMsgTransfer(owner string, denom string, newOwner string) *MsgTransfer {
	return &MsgTransfer{
		Owner:    owner,
		Denom:    denom,
		NewOwner: newOwner,
	}
}

func (msg *MsgTransfer) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}
	return nil
}
