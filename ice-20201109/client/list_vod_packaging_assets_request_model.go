// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVodPackagingAssetsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupName(v string) *ListVodPackagingAssetsRequest
	GetGroupName() *string
	SetKeyword(v string) *ListVodPackagingAssetsRequest
	GetKeyword() *string
	SetPageNo(v int32) *ListVodPackagingAssetsRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListVodPackagingAssetsRequest
	GetPageSize() *int32
	SetSortBy(v string) *ListVodPackagingAssetsRequest
	GetSortBy() *string
}

type ListVodPackagingAssetsRequest struct {
	// The name of the packaging group.
	//
	// example:
	//
	// vod_hls
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The search keyword. The names of the returned assets are prefixed with this keyword.
	//
	// example:
	//
	// movie
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The sorting order of the assets based on the time when they were ingested. Valid values:
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

func (s ListVodPackagingAssetsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVodPackagingAssetsRequest) GoString() string {
	return s.String()
}

func (s *ListVodPackagingAssetsRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *ListVodPackagingAssetsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListVodPackagingAssetsRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListVodPackagingAssetsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListVodPackagingAssetsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListVodPackagingAssetsRequest) SetGroupName(v string) *ListVodPackagingAssetsRequest {
	s.GroupName = &v
	return s
}

func (s *ListVodPackagingAssetsRequest) SetKeyword(v string) *ListVodPackagingAssetsRequest {
	s.Keyword = &v
	return s
}

func (s *ListVodPackagingAssetsRequest) SetPageNo(v int32) *ListVodPackagingAssetsRequest {
	s.PageNo = &v
	return s
}

func (s *ListVodPackagingAssetsRequest) SetPageSize(v int32) *ListVodPackagingAssetsRequest {
	s.PageSize = &v
	return s
}

func (s *ListVodPackagingAssetsRequest) SetSortBy(v string) *ListVodPackagingAssetsRequest {
	s.SortBy = &v
	return s
}

func (s *ListVodPackagingAssetsRequest) Validate() error {
	return dara.Validate(s)
}
