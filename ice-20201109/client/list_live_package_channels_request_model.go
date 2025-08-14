// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLivePackageChannelsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupName(v string) *ListLivePackageChannelsRequest
	GetGroupName() *string
	SetKeyword(v string) *ListLivePackageChannelsRequest
	GetKeyword() *string
	SetPageNo(v int64) *ListLivePackageChannelsRequest
	GetPageNo() *int64
	SetPageSize(v int64) *ListLivePackageChannelsRequest
	GetPageSize() *int64
	SetSortBy(v string) *ListLivePackageChannelsRequest
	GetSortBy() *string
}

type ListLivePackageChannelsRequest struct {
	// The channel group name.
	//
	// This parameter is required.
	//
	// example:
	//
	// channel-group-1
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The channel name or description. Fuzzy match is supported.
	//
	// example:
	//
	// group-1
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
	// The sort order by creation time. Default value: desc.
	//
	// Valid values:
	//
	// 	- asc
	//
	// 	- desc
	//
	// example:
	//
	// desc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
}

func (s ListLivePackageChannelsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLivePackageChannelsRequest) GoString() string {
	return s.String()
}

func (s *ListLivePackageChannelsRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *ListLivePackageChannelsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListLivePackageChannelsRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *ListLivePackageChannelsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListLivePackageChannelsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListLivePackageChannelsRequest) SetGroupName(v string) *ListLivePackageChannelsRequest {
	s.GroupName = &v
	return s
}

func (s *ListLivePackageChannelsRequest) SetKeyword(v string) *ListLivePackageChannelsRequest {
	s.Keyword = &v
	return s
}

func (s *ListLivePackageChannelsRequest) SetPageNo(v int64) *ListLivePackageChannelsRequest {
	s.PageNo = &v
	return s
}

func (s *ListLivePackageChannelsRequest) SetPageSize(v int64) *ListLivePackageChannelsRequest {
	s.PageSize = &v
	return s
}

func (s *ListLivePackageChannelsRequest) SetSortBy(v string) *ListLivePackageChannelsRequest {
	s.SortBy = &v
	return s
}

func (s *ListLivePackageChannelsRequest) Validate() error {
	return dara.Validate(s)
}
