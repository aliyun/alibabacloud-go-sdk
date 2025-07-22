// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFundAccountPayRelationShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEcIdAccountIdsShrink(v string) *CreateFundAccountPayRelationShrinkRequest
	GetEcIdAccountIdsShrink() *string
	SetFundAccountId(v string) *CreateFundAccountPayRelationShrinkRequest
	GetFundAccountId() *string
	SetNbid(v string) *CreateFundAccountPayRelationShrinkRequest
	GetNbid() *string
}

type CreateFundAccountPayRelationShrinkRequest struct {
	// This parameter is required.
	EcIdAccountIdsShrink *string `json:"EcIdAccountIds,omitempty" xml:"EcIdAccountIds,omitempty"`
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

func (s CreateFundAccountPayRelationShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFundAccountPayRelationShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateFundAccountPayRelationShrinkRequest) GetEcIdAccountIdsShrink() *string {
	return s.EcIdAccountIdsShrink
}

func (s *CreateFundAccountPayRelationShrinkRequest) GetFundAccountId() *string {
	return s.FundAccountId
}

func (s *CreateFundAccountPayRelationShrinkRequest) GetNbid() *string {
	return s.Nbid
}

func (s *CreateFundAccountPayRelationShrinkRequest) SetEcIdAccountIdsShrink(v string) *CreateFundAccountPayRelationShrinkRequest {
	s.EcIdAccountIdsShrink = &v
	return s
}

func (s *CreateFundAccountPayRelationShrinkRequest) SetFundAccountId(v string) *CreateFundAccountPayRelationShrinkRequest {
	s.FundAccountId = &v
	return s
}

func (s *CreateFundAccountPayRelationShrinkRequest) SetNbid(v string) *CreateFundAccountPayRelationShrinkRequest {
	s.Nbid = &v
	return s
}

func (s *CreateFundAccountPayRelationShrinkRequest) Validate() error {
	return dara.Validate(s)
}
