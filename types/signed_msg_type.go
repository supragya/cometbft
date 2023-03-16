package types

import cmtproto "github.com/cometbft/cometbft/proto/cometbft/types/v1"

type SignedMsgType = cmtproto.SignedMsgType

const (
	SignedMsgType_UNKNOWN SignedMsgType = cmdproto.UnknownType
	SignedMsgType_PREVOTE SignedMsgType = cmdproto.PrevoteType
	SignedMsgType_PRECOMMIT SignedMsgType = cmdproto.PrecommitType
	SignedMsgType_PROPOSAL SignedMsgType = cmdproto.ProposalType
)

// IsVoteTypeValid returns true if t is a valid vote type.
func IsVoteTypeValid(t SignedMsgType) bool {
	switch t {
	case SignedMsgType_PREVOTE, SignedMsgType_PRECOMMIT:
		return true
	default:
		return false
	}
}
