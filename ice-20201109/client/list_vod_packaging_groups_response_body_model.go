// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVodPackagingGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPackagingGroups(v []*VodPackagingGroup) *ListVodPackagingGroupsResponseBody
	GetPackagingGroups() []*VodPackagingGroup
	SetPageNo(v int32) *ListVodPackagingGroupsResponseBody
	GetPageNo() *int32
	SetPageSize(v int32) *ListVodPackagingGroupsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListVodPackagingGroupsResponseBody
	GetRequestId() *string
	SetSortBy(v string) *ListVodPackagingGroupsResponseBody
	GetSortBy() *string
	SetTotalCount(v int64) *ListVodPackagingGroupsResponseBody
	GetTotalCount() *int64
}

type ListVodPackagingGroupsResponseBody struct {
	// The packaging groups.
	PackagingGroups []*VodPackagingGroup `json:"PackagingGroups,omitempty" xml:"PackagingGroups,omitempty" type:"Repeated"`
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
	// The ID of the request.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The sorting order of the packaging groups based on the time when they were created. Valid values:
	//
	// 	- desc: descending order.
	//
	// 	- asc: ascending order.
	//
	// example:
	//
	// desc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 5
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListVodPackagingGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVodPackagingGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListVodPackagingGroupsResponseBody) GetPackagingGroups() []*VodPackagingGroup {
	return s.PackagingGroups
}

func (s *ListVodPackagingGroupsResponseBody) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListVodPackagingGroupsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListVodPackagingGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVodPackagingGroupsResponseBody) GetSortBy() *string {
	return s.SortBy
}

func (s *ListVodPackagingGroupsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListVodPackagingGroupsResponseBody) SetPackagingGroups(v []*VodPackagingGroup) *ListVodPackagingGroupsResponseBody {
	s.PackagingGroups = v
	return s
}

func (s *ListVodPackagingGroupsResponseBody) SetPageNo(v int32) *ListVodPackagingGroupsResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListVodPackagingGroupsResponseBody) SetPageSize(v int32) *ListVodPackagingGroupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListVodPackagingGroupsResponseBody) SetRequestId(v string) *ListVodPackagingGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVodPackagingGroupsResponseBody) SetSortBy(v string) *ListVodPackagingGroupsResponseBody {
	s.SortBy = &v
	return s
}

func (s *ListVodPackagingGroupsResponseBody) SetTotalCount(v int64) *ListVodPackagingGroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListVodPackagingGroupsResponseBody) Validate() error {
	return dara.Validate(s)
}
