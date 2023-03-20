package privval

import (
	"fmt"

	"github.com/cometbft/cometbft/crypto"
	cryptoenc "github.com/cometbft/cometbft/crypto/encoding"
	cryptoproto "github.com/cometbft/cometbft/proto/cometbft/crypto/v1"
	pvproto1 "github.com/cometbft/cometbft/proto/cometbft/privval/v1"
	pvproto2 "github.com/cometbft/cometbft/proto/cometbft/privval/v2"
	cmtproto "github.com/cometbft/cometbft/proto/cometbft/types/v3"
	cmtproto1 "github.com/cometbft/cometbft/proto/cometbft/types/v1"
	"github.com/cometbft/cometbft/types"
)

func DefaultValidationRequestHandler(
	privVal types.PrivValidator,
	req pvproto2.Message,
	chainID string,
) (pvproto2.Message, error) {
	var (
		res pvproto2.Message
		err error
	)

	switch r := req.Sum.(type) {
	case *pvproto2.Message_PubKeyRequest:
		if r.PubKeyRequest.GetChainId() != chainID {
			res = mustWrapMsg(&pvproto1.PubKeyResponse{
				PubKey: cryptoproto.PublicKey{}, Error: &pvproto1.RemoteSignerError{
					Code: 0, Description: "unable to provide pubkey"}})
			return res, fmt.Errorf("want chainID: %s, got chainID: %s", r.PubKeyRequest.GetChainId(), chainID)
		}

		var pubKey crypto.PubKey
		pubKey, err = privVal.GetPubKey()
		if err != nil {
			return res, err
		}
		pk, err := cryptoenc.PubKeyToProto(pubKey)
		if err != nil {
			return res, err
		}

		if err != nil {
			res = mustWrapMsg(&pvproto1.PubKeyResponse{
				PubKey: cryptoproto.PublicKey{}, Error: &pvproto1.RemoteSignerError{Code: 0, Description: err.Error()}})
		} else {
			res = mustWrapMsg(&pvproto1.PubKeyResponse{PubKey: pk, Error: nil})
		}

	case *pvproto2.Message_SignVoteRequest:
		if r.SignVoteRequest.ChainId != chainID {
			res = mustWrapMsg(&pvproto2.SignedVoteResponse{
				Vote: cmtproto.Vote{}, Error: &pvproto1.RemoteSignerError{
					Code: 0, Description: "unable to sign vote"}})
			return res, fmt.Errorf("want chainID: %s, got chainID: %s", r.SignVoteRequest.GetChainId(), chainID)
		}

		vote := r.SignVoteRequest.Vote

		err = privVal.SignVote(chainID, vote)
		if err != nil {
			res = mustWrapMsg(&pvproto2.SignedVoteResponse{
				Vote: cmtproto.Vote{}, Error: &pvproto1.RemoteSignerError{Code: 0, Description: err.Error()}})
		} else {
			res = mustWrapMsg(&pvproto2.SignedVoteResponse{Vote: *vote, Error: nil})
		}

	case *pvproto2.Message_SignProposalRequest:
		if r.SignProposalRequest.GetChainId() != chainID {
			res = mustWrapMsg(&pvproto1.SignedProposalResponse{
				Proposal: cmtproto1.Proposal{}, Error: &pvproto1.RemoteSignerError{
					Code:        0,
					Description: "unable to sign proposal"}})
			return res, fmt.Errorf("want chainID: %s, got chainID: %s", r.SignProposalRequest.GetChainId(), chainID)
		}

		proposal := r.SignProposalRequest.Proposal

		err = privVal.SignProposal(chainID, proposal)
		if err != nil {
			res = mustWrapMsg(&pvproto1.SignedProposalResponse{
				Proposal: cmtproto1.Proposal{}, Error: &pvproto1.RemoteSignerError{Code: 0, Description: err.Error()}})
		} else {
			res = mustWrapMsg(&pvproto1.SignedProposalResponse{Proposal: *proposal, Error: nil})
		}
	case *pvproto2.Message_PingRequest:
		err, res = nil, mustWrapMsg(&pvproto1.PingResponse{})

	default:
		err = fmt.Errorf("unknown msg: %v", r)
	}

	return res, err
}
