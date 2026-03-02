// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLibraryReviewsListQuery interface {
	dara.Model
	String() string
	GoString() string
	SetApplicant(v string) *LibraryReviewsListQuery
	GetApplicant() *string
	SetMarketId(v int64) *LibraryReviewsListQuery
	GetMarketId() *int64
	SetOrderBy(v string) *LibraryReviewsListQuery
	GetOrderBy() *string
	SetOrderDirection(v string) *LibraryReviewsListQuery
	GetOrderDirection() *string
	SetPageNumber(v int32) *LibraryReviewsListQuery
	GetPageNumber() *int32
	SetPageSize(v int32) *LibraryReviewsListQuery
	GetPageSize() *int32
	SetReviewer(v string) *LibraryReviewsListQuery
	GetReviewer() *string
}

type LibraryReviewsListQuery struct {
	Applicant      *string `json:"applicant,omitempty" xml:"applicant,omitempty"`
	MarketId       *int64  `json:"marketId,omitempty" xml:"marketId,omitempty"`
	OrderBy        *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	OrderDirection *string `json:"orderDirection,omitempty" xml:"orderDirection,omitempty"`
	PageNumber     *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize       *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Reviewer       *string `json:"reviewer,omitempty" xml:"reviewer,omitempty"`
}

func (s LibraryReviewsListQuery) String() string {
	return dara.Prettify(s)
}

func (s LibraryReviewsListQuery) GoString() string {
	return s.String()
}

func (s *LibraryReviewsListQuery) GetApplicant() *string {
	return s.Applicant
}

func (s *LibraryReviewsListQuery) GetMarketId() *int64 {
	return s.MarketId
}

func (s *LibraryReviewsListQuery) GetOrderBy() *string {
	return s.OrderBy
}

func (s *LibraryReviewsListQuery) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *LibraryReviewsListQuery) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *LibraryReviewsListQuery) GetPageSize() *int32 {
	return s.PageSize
}

func (s *LibraryReviewsListQuery) GetReviewer() *string {
	return s.Reviewer
}

func (s *LibraryReviewsListQuery) SetApplicant(v string) *LibraryReviewsListQuery {
	s.Applicant = &v
	return s
}

func (s *LibraryReviewsListQuery) SetMarketId(v int64) *LibraryReviewsListQuery {
	s.MarketId = &v
	return s
}

func (s *LibraryReviewsListQuery) SetOrderBy(v string) *LibraryReviewsListQuery {
	s.OrderBy = &v
	return s
}

func (s *LibraryReviewsListQuery) SetOrderDirection(v string) *LibraryReviewsListQuery {
	s.OrderDirection = &v
	return s
}

func (s *LibraryReviewsListQuery) SetPageNumber(v int32) *LibraryReviewsListQuery {
	s.PageNumber = &v
	return s
}

func (s *LibraryReviewsListQuery) SetPageSize(v int32) *LibraryReviewsListQuery {
	s.PageSize = &v
	return s
}

func (s *LibraryReviewsListQuery) SetReviewer(v string) *LibraryReviewsListQuery {
	s.Reviewer = &v
	return s
}

func (s *LibraryReviewsListQuery) Validate() error {
	return dara.Validate(s)
}
