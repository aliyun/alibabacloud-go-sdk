// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPbcReviewsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicant(v string) *ListPbcReviewsRequest
	GetApplicant() *string
	SetMarketId(v int64) *ListPbcReviewsRequest
	GetMarketId() *int64
	SetOrderBy(v int32) *ListPbcReviewsRequest
	GetOrderBy() *int32
	SetOrderDirection(v int32) *ListPbcReviewsRequest
	GetOrderDirection() *int32
	SetPageNumber(v string) *ListPbcReviewsRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListPbcReviewsRequest
	GetPageSize() *string
	SetReviewer(v string) *ListPbcReviewsRequest
	GetReviewer() *string
}

type ListPbcReviewsRequest struct {
	Applicant *string `json:"applicant,omitempty" xml:"applicant,omitempty"`
	// This parameter is required.
	MarketId       *int64  `json:"marketId,omitempty" xml:"marketId,omitempty"`
	OrderBy        *int32  `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	OrderDirection *int32  `json:"orderDirection,omitempty" xml:"orderDirection,omitempty"`
	PageNumber     *string `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize       *string `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Reviewer       *string `json:"reviewer,omitempty" xml:"reviewer,omitempty"`
}

func (s ListPbcReviewsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPbcReviewsRequest) GoString() string {
	return s.String()
}

func (s *ListPbcReviewsRequest) GetApplicant() *string {
	return s.Applicant
}

func (s *ListPbcReviewsRequest) GetMarketId() *int64 {
	return s.MarketId
}

func (s *ListPbcReviewsRequest) GetOrderBy() *int32 {
	return s.OrderBy
}

func (s *ListPbcReviewsRequest) GetOrderDirection() *int32 {
	return s.OrderDirection
}

func (s *ListPbcReviewsRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListPbcReviewsRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListPbcReviewsRequest) GetReviewer() *string {
	return s.Reviewer
}

func (s *ListPbcReviewsRequest) SetApplicant(v string) *ListPbcReviewsRequest {
	s.Applicant = &v
	return s
}

func (s *ListPbcReviewsRequest) SetMarketId(v int64) *ListPbcReviewsRequest {
	s.MarketId = &v
	return s
}

func (s *ListPbcReviewsRequest) SetOrderBy(v int32) *ListPbcReviewsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListPbcReviewsRequest) SetOrderDirection(v int32) *ListPbcReviewsRequest {
	s.OrderDirection = &v
	return s
}

func (s *ListPbcReviewsRequest) SetPageNumber(v string) *ListPbcReviewsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPbcReviewsRequest) SetPageSize(v string) *ListPbcReviewsRequest {
	s.PageSize = &v
	return s
}

func (s *ListPbcReviewsRequest) SetReviewer(v string) *ListPbcReviewsRequest {
	s.Reviewer = &v
	return s
}

func (s *ListPbcReviewsRequest) Validate() error {
	return dara.Validate(s)
}
