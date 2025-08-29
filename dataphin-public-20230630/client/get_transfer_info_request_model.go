// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTransferInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *GetTransferInfoRequest
	GetOpTenantId() *int64
	SetProposalId(v int64) *GetTransferInfoRequest
	GetProposalId() *int64
}

type GetTransferInfoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	ProposalId *int64 `json:"ProposalId,omitempty" xml:"ProposalId,omitempty"`
}

func (s GetTransferInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTransferInfoRequest) GoString() string {
	return s.String()
}

func (s *GetTransferInfoRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetTransferInfoRequest) GetProposalId() *int64 {
	return s.ProposalId
}

func (s *GetTransferInfoRequest) SetOpTenantId(v int64) *GetTransferInfoRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetTransferInfoRequest) SetProposalId(v int64) *GetTransferInfoRequest {
	s.ProposalId = &v
	return s
}

func (s *GetTransferInfoRequest) Validate() error {
	return dara.Validate(s)
}
