// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPbcInvokesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicant(v string) *ListPbcInvokesRequest
	GetApplicant() *string
	SetCompanyId(v int64) *ListPbcInvokesRequest
	GetCompanyId() *int64
	SetMarketId(v int64) *ListPbcInvokesRequest
	GetMarketId() *int64
	SetPbcId(v int64) *ListPbcInvokesRequest
	GetPbcId() *int64
}

type ListPbcInvokesRequest struct {
	// example:
	//
	// 223352752411587433
	Applicant *string `json:"applicant,omitempty" xml:"applicant,omitempty"`
	CompanyId *int64  `json:"companyId,omitempty" xml:"companyId,omitempty"`
	MarketId  *int64  `json:"marketId,omitempty" xml:"marketId,omitempty"`
	PbcId     *int64  `json:"pbcId,omitempty" xml:"pbcId,omitempty"`
}

func (s ListPbcInvokesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPbcInvokesRequest) GoString() string {
	return s.String()
}

func (s *ListPbcInvokesRequest) GetApplicant() *string {
	return s.Applicant
}

func (s *ListPbcInvokesRequest) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *ListPbcInvokesRequest) GetMarketId() *int64 {
	return s.MarketId
}

func (s *ListPbcInvokesRequest) GetPbcId() *int64 {
	return s.PbcId
}

func (s *ListPbcInvokesRequest) SetApplicant(v string) *ListPbcInvokesRequest {
	s.Applicant = &v
	return s
}

func (s *ListPbcInvokesRequest) SetCompanyId(v int64) *ListPbcInvokesRequest {
	s.CompanyId = &v
	return s
}

func (s *ListPbcInvokesRequest) SetMarketId(v int64) *ListPbcInvokesRequest {
	s.MarketId = &v
	return s
}

func (s *ListPbcInvokesRequest) SetPbcId(v int64) *ListPbcInvokesRequest {
	s.PbcId = &v
	return s
}

func (s *ListPbcInvokesRequest) Validate() error {
	return dara.Validate(s)
}
