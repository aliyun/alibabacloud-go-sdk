// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLivePackageOriginEndpointsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelName(v string) *ListLivePackageOriginEndpointsRequest
	GetChannelName() *string
	SetGroupName(v string) *ListLivePackageOriginEndpointsRequest
	GetGroupName() *string
	SetKeyword(v string) *ListLivePackageOriginEndpointsRequest
	GetKeyword() *string
	SetPageNo(v int64) *ListLivePackageOriginEndpointsRequest
	GetPageNo() *int64
	SetPageSize(v int64) *ListLivePackageOriginEndpointsRequest
	GetPageSize() *int64
	SetSortBy(v string) *ListLivePackageOriginEndpointsRequest
	GetSortBy() *string
}

type ListLivePackageOriginEndpointsRequest struct {
	// The channel name.
	//
	// This parameter is required.
	//
	// example:
	//
	// channel-1
	ChannelName *string `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	// The channel group name.
	//
	// This parameter is required.
	//
	// example:
	//
	// channel-group-1
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The endpoint name or description. Fuzzy match is supported.
	//
	// example:
	//
	// endpoint-
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The sort order by creation time. Valid values: asc and desc (default).
	//
	// example:
	//
	// desc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
}

func (s ListLivePackageOriginEndpointsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLivePackageOriginEndpointsRequest) GoString() string {
	return s.String()
}

func (s *ListLivePackageOriginEndpointsRequest) GetChannelName() *string {
	return s.ChannelName
}

func (s *ListLivePackageOriginEndpointsRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *ListLivePackageOriginEndpointsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListLivePackageOriginEndpointsRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *ListLivePackageOriginEndpointsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListLivePackageOriginEndpointsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListLivePackageOriginEndpointsRequest) SetChannelName(v string) *ListLivePackageOriginEndpointsRequest {
	s.ChannelName = &v
	return s
}

func (s *ListLivePackageOriginEndpointsRequest) SetGroupName(v string) *ListLivePackageOriginEndpointsRequest {
	s.GroupName = &v
	return s
}

func (s *ListLivePackageOriginEndpointsRequest) SetKeyword(v string) *ListLivePackageOriginEndpointsRequest {
	s.Keyword = &v
	return s
}

func (s *ListLivePackageOriginEndpointsRequest) SetPageNo(v int64) *ListLivePackageOriginEndpointsRequest {
	s.PageNo = &v
	return s
}

func (s *ListLivePackageOriginEndpointsRequest) SetPageSize(v int64) *ListLivePackageOriginEndpointsRequest {
	s.PageSize = &v
	return s
}

func (s *ListLivePackageOriginEndpointsRequest) SetSortBy(v string) *ListLivePackageOriginEndpointsRequest {
	s.SortBy = &v
	return s
}

func (s *ListLivePackageOriginEndpointsRequest) Validate() error {
	return dara.Validate(s)
}
