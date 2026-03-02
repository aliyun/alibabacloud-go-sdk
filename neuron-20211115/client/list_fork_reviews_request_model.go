// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListForkReviewsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicant(v string) *ListForkReviewsRequest
	GetApplicant() *string
	SetCompanyId(v int64) *ListForkReviewsRequest
	GetCompanyId() *int64
	SetMarketId(v int64) *ListForkReviewsRequest
	GetMarketId() *int64
	SetOrderBy(v string) *ListForkReviewsRequest
	GetOrderBy() *string
	SetOrderDirection(v string) *ListForkReviewsRequest
	GetOrderDirection() *string
	SetPageNumber(v int32) *ListForkReviewsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListForkReviewsRequest
	GetPageSize() *int32
	SetReviewer(v string) *ListForkReviewsRequest
	GetReviewer() *string
}

type ListForkReviewsRequest struct {
	Applicant      *string `json:"applicant,omitempty" xml:"applicant,omitempty"`
	CompanyId      *int64  `json:"companyId,omitempty" xml:"companyId,omitempty"`
	MarketId       *int64  `json:"marketId,omitempty" xml:"marketId,omitempty"`
	OrderBy        *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	OrderDirection *string `json:"orderDirection,omitempty" xml:"orderDirection,omitempty"`
	PageNumber     *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize       *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Reviewer       *string `json:"reviewer,omitempty" xml:"reviewer,omitempty"`
}

func (s ListForkReviewsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListForkReviewsRequest) GoString() string {
	return s.String()
}

func (s *ListForkReviewsRequest) GetApplicant() *string {
	return s.Applicant
}

func (s *ListForkReviewsRequest) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *ListForkReviewsRequest) GetMarketId() *int64 {
	return s.MarketId
}

func (s *ListForkReviewsRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListForkReviewsRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *ListForkReviewsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListForkReviewsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListForkReviewsRequest) GetReviewer() *string {
	return s.Reviewer
}

func (s *ListForkReviewsRequest) SetApplicant(v string) *ListForkReviewsRequest {
	s.Applicant = &v
	return s
}

func (s *ListForkReviewsRequest) SetCompanyId(v int64) *ListForkReviewsRequest {
	s.CompanyId = &v
	return s
}

func (s *ListForkReviewsRequest) SetMarketId(v int64) *ListForkReviewsRequest {
	s.MarketId = &v
	return s
}

func (s *ListForkReviewsRequest) SetOrderBy(v string) *ListForkReviewsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListForkReviewsRequest) SetOrderDirection(v string) *ListForkReviewsRequest {
	s.OrderDirection = &v
	return s
}

func (s *ListForkReviewsRequest) SetPageNumber(v int32) *ListForkReviewsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListForkReviewsRequest) SetPageSize(v int32) *ListForkReviewsRequest {
	s.PageSize = &v
	return s
}

func (s *ListForkReviewsRequest) SetReviewer(v string) *ListForkReviewsRequest {
	s.Reviewer = &v
	return s
}

func (s *ListForkReviewsRequest) Validate() error {
	return dara.Validate(s)
}
