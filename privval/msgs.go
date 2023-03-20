package privval

import (
	"fmt"

	"github.com/cosmos/gogoproto/proto"

	pvproto1 "github.com/cometbft/cometbft/proto/cometbft/privval/v1"
	pvproto2 "github.com/cometbft/cometbft/proto/cometbft/privval/v2"
)

// TODO: Add ChainIDRequest

func mustWrapMsg(pb proto.Message) pvproto2.Message {
	msg := pvproto2.Message{}

	switch pb := pb.(type) {
	case *pvproto2.Message:
		msg = *pb
	case *pvproto1.PubKeyRequest:
		msg.Sum = &pvproto2.Message_PubKeyRequest{PubKeyRequest: pb}
	case *pvproto1.PubKeyResponse:
		msg.Sum = &pvproto2.Message_PubKeyResponse{PubKeyResponse: pb}
	case *pvproto2.SignVoteRequest:
		msg.Sum = &pvproto2.Message_SignVoteRequest{SignVoteRequest: pb}
	case *pvproto2.SignedVoteResponse:
		msg.Sum = &pvproto2.Message_SignedVoteResponse{SignedVoteResponse: pb}
	case *pvproto1.SignedProposalResponse:
		msg.Sum = &pvproto2.Message_SignedProposalResponse{SignedProposalResponse: pb}
	case *pvproto1.SignProposalRequest:
		msg.Sum = &pvproto2.Message_SignProposalRequest{SignProposalRequest: pb}
	case *pvproto1.PingRequest:
		msg.Sum = &pvproto2.Message_PingRequest{PingRequest: pb}
	case *pvproto1.PingResponse:
		msg.Sum = &pvproto2.Message_PingResponse{PingResponse: pb}
	default:
		panic(fmt.Errorf("unknown message type %T", pb))
	}

	return msg
}
