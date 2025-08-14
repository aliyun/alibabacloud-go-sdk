// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVodPackagingGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *ListVodPackagingGroupsRequest
	GetKeyword() *string
	SetPageNo(v int64) *ListVodPackagingGroupsRequest
	GetPageNo() *int64
	SetPageSize(v int64) *ListVodPackagingGroupsRequest
	GetPageSize() *int64
	SetSortBy(v string) *ListVodPackagingGroupsRequest
	GetSortBy() *string
}

type ListVodPackagingGroupsRequest struct {
	// The search keyword. The names of the returned packaging groups contain the keyword.
	//
	// example:
	//
	// hls
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The sorting order of the packaging groups based on the time when they were created. Valid values:
	//
	// 	- desc (default): descending order.
	//
	// 	- asc: ascending order.
	//
	// example:
	//
	// desc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
}

func (s ListVodPackagingGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVodPackagingGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListVodPackagingGroupsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListVodPackagingGroupsRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *ListVodPackagingGroupsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListVodPackagingGroupsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListVodPackagingGroupsRequest) SetKeyword(v string) *ListVodPackagingGroupsRequest {
	s.Keyword = &v
	return s
}

func (s *ListVodPackagingGroupsRequest) SetPageNo(v int64) *ListVodPackagingGroupsRequest {
	s.PageNo = &v
	return s
}

func (s *ListVodPackagingGroupsRequest) SetPageSize(v int64) *ListVodPackagingGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *ListVodPackagingGroupsRequest) SetSortBy(v string) *ListVodPackagingGroupsRequest {
	s.SortBy = &v
	return s
}

func (s *ListVodPackagingGroupsRequest) Validate() error {
	return dara.Validate(s)
}
