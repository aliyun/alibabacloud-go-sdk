// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePbcInvokeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicant(v string) *DeletePbcInvokeRequest
	GetApplicant() *string
	SetCompanyId(v int64) *DeletePbcInvokeRequest
	GetCompanyId() *int64
	SetMarketId(v int64) *DeletePbcInvokeRequest
	GetMarketId() *int64
	SetPbcId(v int64) *DeletePbcInvokeRequest
	GetPbcId() *int64
}

type DeletePbcInvokeRequest struct {
	Applicant *string `json:"applicant,omitempty" xml:"applicant,omitempty"`
	CompanyId *int64  `json:"companyId,omitempty" xml:"companyId,omitempty"`
	MarketId  *int64  `json:"marketId,omitempty" xml:"marketId,omitempty"`
	PbcId     *int64  `json:"pbcId,omitempty" xml:"pbcId,omitempty"`
}

func (s DeletePbcInvokeRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePbcInvokeRequest) GoString() string {
	return s.String()
}

func (s *DeletePbcInvokeRequest) GetApplicant() *string {
	return s.Applicant
}

func (s *DeletePbcInvokeRequest) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *DeletePbcInvokeRequest) GetMarketId() *int64 {
	return s.MarketId
}

func (s *DeletePbcInvokeRequest) GetPbcId() *int64 {
	return s.PbcId
}

func (s *DeletePbcInvokeRequest) SetApplicant(v string) *DeletePbcInvokeRequest {
	s.Applicant = &v
	return s
}

func (s *DeletePbcInvokeRequest) SetCompanyId(v int64) *DeletePbcInvokeRequest {
	s.CompanyId = &v
	return s
}

func (s *DeletePbcInvokeRequest) SetMarketId(v int64) *DeletePbcInvokeRequest {
	s.MarketId = &v
	return s
}

func (s *DeletePbcInvokeRequest) SetPbcId(v int64) *DeletePbcInvokeRequest {
	s.PbcId = &v
	return s
}

func (s *DeletePbcInvokeRequest) Validate() error {
	return dara.Validate(s)
}
