// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInfiniteCanvasesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *ListInfiniteCanvasesRequest
	GetKeyword() *string
	SetPageNo(v int32) *ListInfiniteCanvasesRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListInfiniteCanvasesRequest
	GetPageSize() *int32
	SetSortBy(v string) *ListInfiniteCanvasesRequest
	GetSortBy() *string
	SetSortOrder(v string) *ListInfiniteCanvasesRequest
	GetSortOrder() *string
}

type ListInfiniteCanvasesRequest struct {
	// The keyword for querying site monitoring tasks. Supports fuzzy match based on task name or task address.
	//
	// example:
	//
	// v2_
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The current page number. Default value: 1.
	//
	// example:
	//
	// 16
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The sort field and sort order. Separate multiple values with commas (,).
	//
	// example:
	//
	// utcCreate:Desc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The sort direction.
	//
	// Valid values:
	//
	// - Asc: Ascending order.
	//
	// - Desc: Descending order.
	//
	// Default value: Desc.
	//
	// example:
	//
	// Ascending
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
}

func (s ListInfiniteCanvasesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInfiniteCanvasesRequest) GoString() string {
	return s.String()
}

func (s *ListInfiniteCanvasesRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListInfiniteCanvasesRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListInfiniteCanvasesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListInfiniteCanvasesRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListInfiniteCanvasesRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListInfiniteCanvasesRequest) SetKeyword(v string) *ListInfiniteCanvasesRequest {
	s.Keyword = &v
	return s
}

func (s *ListInfiniteCanvasesRequest) SetPageNo(v int32) *ListInfiniteCanvasesRequest {
	s.PageNo = &v
	return s
}

func (s *ListInfiniteCanvasesRequest) SetPageSize(v int32) *ListInfiniteCanvasesRequest {
	s.PageSize = &v
	return s
}

func (s *ListInfiniteCanvasesRequest) SetSortBy(v string) *ListInfiniteCanvasesRequest {
	s.SortBy = &v
	return s
}

func (s *ListInfiniteCanvasesRequest) SetSortOrder(v string) *ListInfiniteCanvasesRequest {
	s.SortOrder = &v
	return s
}

func (s *ListInfiniteCanvasesRequest) Validate() error {
	return dara.Validate(s)
}
