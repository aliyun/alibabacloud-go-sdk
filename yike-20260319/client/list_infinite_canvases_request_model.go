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
	// The query keyword. Currently, only searching by infinite canvas ID is supported.
	//
	// example:
	//
	// canvas_xxx
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The current page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The field by which the results are sorted.
	//
	// example:
	//
	// CreatedTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The sort order. Valid values:
	//
	//  	- asc: ascending order
	//
	//  	- desc: descending order
	//
	// example:
	//
	// asc
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
