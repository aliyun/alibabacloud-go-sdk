// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagOptionsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFiltersShrink(v string) *ListTagOptionsShrinkRequest
	GetFiltersShrink() *string
	SetMaxResults(v int32) *ListTagOptionsShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListTagOptionsShrinkRequest
	GetNextToken() *string
	SetPageNumber(v int32) *ListTagOptionsShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTagOptionsShrinkRequest
	GetPageSize() *int32
	SetSortBy(v string) *ListTagOptionsShrinkRequest
	GetSortBy() *string
	SetSortOrder(v string) *ListTagOptionsShrinkRequest
	GetSortOrder() *string
}

type ListTagOptionsShrinkRequest struct {
	// The filter condition.
	FiltersShrink *string `json:"Filters,omitempty" xml:"Filters,omitempty"`
	MaxResults    *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken     *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The number of the page to return.
	//
	// Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: 1 to 100. Minimum value: 1. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The information based on which you want to sort the query results.
	//
	// Set the value to CreateTime, which specifies the creation time of tag options.
	//
	// example:
	//
	// CreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The order in which you want to sort the query results. Valid values:
	//
	// 	- Asc: the ascending order
	//
	// 	- Desc (default): the descending order
	//
	// example:
	//
	// Desc
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
}

func (s ListTagOptionsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTagOptionsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListTagOptionsShrinkRequest) GetFiltersShrink() *string {
	return s.FiltersShrink
}

func (s *ListTagOptionsShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTagOptionsShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTagOptionsShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTagOptionsShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTagOptionsShrinkRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListTagOptionsShrinkRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListTagOptionsShrinkRequest) SetFiltersShrink(v string) *ListTagOptionsShrinkRequest {
	s.FiltersShrink = &v
	return s
}

func (s *ListTagOptionsShrinkRequest) SetMaxResults(v int32) *ListTagOptionsShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTagOptionsShrinkRequest) SetNextToken(v string) *ListTagOptionsShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagOptionsShrinkRequest) SetPageNumber(v int32) *ListTagOptionsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTagOptionsShrinkRequest) SetPageSize(v int32) *ListTagOptionsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListTagOptionsShrinkRequest) SetSortBy(v string) *ListTagOptionsShrinkRequest {
	s.SortBy = &v
	return s
}

func (s *ListTagOptionsShrinkRequest) SetSortOrder(v string) *ListTagOptionsShrinkRequest {
	s.SortOrder = &v
	return s
}

func (s *ListTagOptionsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
