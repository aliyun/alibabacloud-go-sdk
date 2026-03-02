// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLibraryReviewsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicant(v string) *ListLibraryReviewsRequest
	GetApplicant() *string
	SetMarketId(v int64) *ListLibraryReviewsRequest
	GetMarketId() *int64
	SetOrderBy(v string) *ListLibraryReviewsRequest
	GetOrderBy() *string
	SetOrderDirection(v string) *ListLibraryReviewsRequest
	GetOrderDirection() *string
	SetPageNumber(v int32) *ListLibraryReviewsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListLibraryReviewsRequest
	GetPageSize() *int32
	SetReviewer(v string) *ListLibraryReviewsRequest
	GetReviewer() *string
}

type ListLibraryReviewsRequest struct {
	Applicant      *string `json:"applicant,omitempty" xml:"applicant,omitempty"`
	MarketId       *int64  `json:"marketId,omitempty" xml:"marketId,omitempty"`
	OrderBy        *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	OrderDirection *string `json:"orderDirection,omitempty" xml:"orderDirection,omitempty"`
	PageNumber     *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize       *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Reviewer       *string `json:"reviewer,omitempty" xml:"reviewer,omitempty"`
}

func (s ListLibraryReviewsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLibraryReviewsRequest) GoString() string {
	return s.String()
}

func (s *ListLibraryReviewsRequest) GetApplicant() *string {
	return s.Applicant
}

func (s *ListLibraryReviewsRequest) GetMarketId() *int64 {
	return s.MarketId
}

func (s *ListLibraryReviewsRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListLibraryReviewsRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *ListLibraryReviewsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListLibraryReviewsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListLibraryReviewsRequest) GetReviewer() *string {
	return s.Reviewer
}

func (s *ListLibraryReviewsRequest) SetApplicant(v string) *ListLibraryReviewsRequest {
	s.Applicant = &v
	return s
}

func (s *ListLibraryReviewsRequest) SetMarketId(v int64) *ListLibraryReviewsRequest {
	s.MarketId = &v
	return s
}

func (s *ListLibraryReviewsRequest) SetOrderBy(v string) *ListLibraryReviewsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListLibraryReviewsRequest) SetOrderDirection(v string) *ListLibraryReviewsRequest {
	s.OrderDirection = &v
	return s
}

func (s *ListLibraryReviewsRequest) SetPageNumber(v int32) *ListLibraryReviewsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListLibraryReviewsRequest) SetPageSize(v int32) *ListLibraryReviewsRequest {
	s.PageSize = &v
	return s
}

func (s *ListLibraryReviewsRequest) SetReviewer(v string) *ListLibraryReviewsRequest {
	s.Reviewer = &v
	return s
}

func (s *ListLibraryReviewsRequest) Validate() error {
	return dara.Validate(s)
}
