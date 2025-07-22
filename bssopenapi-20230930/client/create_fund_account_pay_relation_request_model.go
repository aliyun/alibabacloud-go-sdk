// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFundAccountPayRelationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEcIdAccountIds(v []*CreateFundAccountPayRelationRequestEcIdAccountIds) *CreateFundAccountPayRelationRequest
	GetEcIdAccountIds() []*CreateFundAccountPayRelationRequestEcIdAccountIds
	SetFundAccountId(v string) *CreateFundAccountPayRelationRequest
	GetFundAccountId() *string
	SetNbid(v string) *CreateFundAccountPayRelationRequest
	GetNbid() *string
}

type CreateFundAccountPayRelationRequest struct {
	// This parameter is required.
	EcIdAccountIds []*CreateFundAccountPayRelationRequestEcIdAccountIds `json:"EcIdAccountIds,omitempty" xml:"EcIdAccountIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 12332112
	FundAccountId *string `json:"FundAccountId,omitempty" xml:"FundAccountId,omitempty"`
	// example:
	//
	// 2684201000001
	Nbid *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
}

func (s CreateFundAccountPayRelationRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFundAccountPayRelationRequest) GoString() string {
	return s.String()
}

func (s *CreateFundAccountPayRelationRequest) GetEcIdAccountIds() []*CreateFundAccountPayRelationRequestEcIdAccountIds {
	return s.EcIdAccountIds
}

func (s *CreateFundAccountPayRelationRequest) GetFundAccountId() *string {
	return s.FundAccountId
}

func (s *CreateFundAccountPayRelationRequest) GetNbid() *string {
	return s.Nbid
}

func (s *CreateFundAccountPayRelationRequest) SetEcIdAccountIds(v []*CreateFundAccountPayRelationRequestEcIdAccountIds) *CreateFundAccountPayRelationRequest {
	s.EcIdAccountIds = v
	return s
}

func (s *CreateFundAccountPayRelationRequest) SetFundAccountId(v string) *CreateFundAccountPayRelationRequest {
	s.FundAccountId = &v
	return s
}

func (s *CreateFundAccountPayRelationRequest) SetNbid(v string) *CreateFundAccountPayRelationRequest {
	s.Nbid = &v
	return s
}

func (s *CreateFundAccountPayRelationRequest) Validate() error {
	return dara.Validate(s)
}

type CreateFundAccountPayRelationRequestEcIdAccountIds struct {
	// This parameter is required.
	AccountIds []*int64 `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	// example:
	//
	// 1501603440974415
	EcId *string `json:"EcId,omitempty" xml:"EcId,omitempty"`
}

func (s CreateFundAccountPayRelationRequestEcIdAccountIds) String() string {
	return dara.Prettify(s)
}

func (s CreateFundAccountPayRelationRequestEcIdAccountIds) GoString() string {
	return s.String()
}

func (s *CreateFundAccountPayRelationRequestEcIdAccountIds) GetAccountIds() []*int64 {
	return s.AccountIds
}

func (s *CreateFundAccountPayRelationRequestEcIdAccountIds) GetEcId() *string {
	return s.EcId
}

func (s *CreateFundAccountPayRelationRequestEcIdAccountIds) SetAccountIds(v []*int64) *CreateFundAccountPayRelationRequestEcIdAccountIds {
	s.AccountIds = v
	return s
}

func (s *CreateFundAccountPayRelationRequestEcIdAccountIds) SetEcId(v string) *CreateFundAccountPayRelationRequestEcIdAccountIds {
	s.EcId = &v
	return s
}

func (s *CreateFundAccountPayRelationRequestEcIdAccountIds) Validate() error {
	return dara.Validate(s)
}
