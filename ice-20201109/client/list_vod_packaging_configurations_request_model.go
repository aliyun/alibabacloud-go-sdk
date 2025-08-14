// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVodPackagingConfigurationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupName(v string) *ListVodPackagingConfigurationsRequest
	GetGroupName() *string
	SetKeyword(v string) *ListVodPackagingConfigurationsRequest
	GetKeyword() *string
	SetPageNo(v int64) *ListVodPackagingConfigurationsRequest
	GetPageNo() *int64
	SetPageSize(v int64) *ListVodPackagingConfigurationsRequest
	GetPageSize() *int64
	SetSortBy(v string) *ListVodPackagingConfigurationsRequest
	GetSortBy() *string
}

type ListVodPackagingConfigurationsRequest struct {
	// The name of the packaging group.
	//
	// example:
	//
	// vod_hls
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The search keyword. The names of the returned packaging configurations contain the keyword.
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
	// The sorting order of the packaging configurations based on the time when they were created. Valid values:
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

func (s ListVodPackagingConfigurationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVodPackagingConfigurationsRequest) GoString() string {
	return s.String()
}

func (s *ListVodPackagingConfigurationsRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *ListVodPackagingConfigurationsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListVodPackagingConfigurationsRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *ListVodPackagingConfigurationsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListVodPackagingConfigurationsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListVodPackagingConfigurationsRequest) SetGroupName(v string) *ListVodPackagingConfigurationsRequest {
	s.GroupName = &v
	return s
}

func (s *ListVodPackagingConfigurationsRequest) SetKeyword(v string) *ListVodPackagingConfigurationsRequest {
	s.Keyword = &v
	return s
}

func (s *ListVodPackagingConfigurationsRequest) SetPageNo(v int64) *ListVodPackagingConfigurationsRequest {
	s.PageNo = &v
	return s
}

func (s *ListVodPackagingConfigurationsRequest) SetPageSize(v int64) *ListVodPackagingConfigurationsRequest {
	s.PageSize = &v
	return s
}

func (s *ListVodPackagingConfigurationsRequest) SetSortBy(v string) *ListVodPackagingConfigurationsRequest {
	s.SortBy = &v
	return s
}

func (s *ListVodPackagingConfigurationsRequest) Validate() error {
	return dara.Validate(s)
}
