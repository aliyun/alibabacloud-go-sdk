// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLivePackageChannelGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *ListLivePackageChannelGroupsRequest
	GetKeyword() *string
	SetPageNo(v int64) *ListLivePackageChannelGroupsRequest
	GetPageNo() *int64
	SetPageSize(v int64) *ListLivePackageChannelGroupsRequest
	GetPageSize() *int64
	SetSortBy(v string) *ListLivePackageChannelGroupsRequest
	GetSortBy() *string
}

type ListLivePackageChannelGroupsRequest struct {
	// The channel group name or description. Fuzzy match is supported.
	//
	// example:
	//
	// channel-group
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The sort order by creation time. Default value: desc.
	//
	// example:
	//
	// desc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
}

func (s ListLivePackageChannelGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLivePackageChannelGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListLivePackageChannelGroupsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListLivePackageChannelGroupsRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *ListLivePackageChannelGroupsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListLivePackageChannelGroupsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListLivePackageChannelGroupsRequest) SetKeyword(v string) *ListLivePackageChannelGroupsRequest {
	s.Keyword = &v
	return s
}

func (s *ListLivePackageChannelGroupsRequest) SetPageNo(v int64) *ListLivePackageChannelGroupsRequest {
	s.PageNo = &v
	return s
}

func (s *ListLivePackageChannelGroupsRequest) SetPageSize(v int64) *ListLivePackageChannelGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *ListLivePackageChannelGroupsRequest) SetSortBy(v string) *ListLivePackageChannelGroupsRequest {
	s.SortBy = &v
	return s
}

func (s *ListLivePackageChannelGroupsRequest) Validate() error {
	return dara.Validate(s)
}
