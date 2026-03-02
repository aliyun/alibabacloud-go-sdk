// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPbcInvokeReviewsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicant(v string) *ListPbcInvokeReviewsRequest
	GetApplicant() *string
	SetCompanyId(v int64) *ListPbcInvokeReviewsRequest
	GetCompanyId() *int64
	SetMarketId(v int64) *ListPbcInvokeReviewsRequest
	GetMarketId() *int64
	SetOrderDirection(v int32) *ListPbcInvokeReviewsRequest
	GetOrderDirection() *int32
	SetOrderby(v int32) *ListPbcInvokeReviewsRequest
	GetOrderby() *int32
	SetReviewer(v string) *ListPbcInvokeReviewsRequest
	GetReviewer() *string
}

type ListPbcInvokeReviewsRequest struct {
	Applicant      *string `json:"applicant,omitempty" xml:"applicant,omitempty"`
	CompanyId      *int64  `json:"companyId,omitempty" xml:"companyId,omitempty"`
	MarketId       *int64  `json:"marketId,omitempty" xml:"marketId,omitempty"`
	OrderDirection *int32  `json:"orderDirection,omitempty" xml:"orderDirection,omitempty"`
	Orderby        *int32  `json:"orderby,omitempty" xml:"orderby,omitempty"`
	Reviewer       *string `json:"reviewer,omitempty" xml:"reviewer,omitempty"`
}

func (s ListPbcInvokeReviewsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPbcInvokeReviewsRequest) GoString() string {
	return s.String()
}

func (s *ListPbcInvokeReviewsRequest) GetApplicant() *string {
	return s.Applicant
}

func (s *ListPbcInvokeReviewsRequest) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *ListPbcInvokeReviewsRequest) GetMarketId() *int64 {
	return s.MarketId
}

func (s *ListPbcInvokeReviewsRequest) GetOrderDirection() *int32 {
	return s.OrderDirection
}

func (s *ListPbcInvokeReviewsRequest) GetOrderby() *int32 {
	return s.Orderby
}

func (s *ListPbcInvokeReviewsRequest) GetReviewer() *string {
	return s.Reviewer
}

func (s *ListPbcInvokeReviewsRequest) SetApplicant(v string) *ListPbcInvokeReviewsRequest {
	s.Applicant = &v
	return s
}

func (s *ListPbcInvokeReviewsRequest) SetCompanyId(v int64) *ListPbcInvokeReviewsRequest {
	s.CompanyId = &v
	return s
}

func (s *ListPbcInvokeReviewsRequest) SetMarketId(v int64) *ListPbcInvokeReviewsRequest {
	s.MarketId = &v
	return s
}

func (s *ListPbcInvokeReviewsRequest) SetOrderDirection(v int32) *ListPbcInvokeReviewsRequest {
	s.OrderDirection = &v
	return s
}

func (s *ListPbcInvokeReviewsRequest) SetOrderby(v int32) *ListPbcInvokeReviewsRequest {
	s.Orderby = &v
	return s
}

func (s *ListPbcInvokeReviewsRequest) SetReviewer(v string) *ListPbcInvokeReviewsRequest {
	s.Reviewer = &v
	return s
}

func (s *ListPbcInvokeReviewsRequest) Validate() error {
	return dara.Validate(s)
}
