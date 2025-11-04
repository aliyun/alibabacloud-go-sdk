// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLivePackageChannelGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLivePackageChannelGroups(v []*ListLivePackageChannelGroupsResponseBodyLivePackageChannelGroups) *ListLivePackageChannelGroupsResponseBody
	GetLivePackageChannelGroups() []*ListLivePackageChannelGroupsResponseBodyLivePackageChannelGroups
	SetPageNo(v int64) *ListLivePackageChannelGroupsResponseBody
	GetPageNo() *int64
	SetPageSize(v int64) *ListLivePackageChannelGroupsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListLivePackageChannelGroupsResponseBody
	GetRequestId() *string
	SetSortBy(v string) *ListLivePackageChannelGroupsResponseBody
	GetSortBy() *string
	SetTotalCount(v int64) *ListLivePackageChannelGroupsResponseBody
	GetTotalCount() *int64
}

type ListLivePackageChannelGroupsResponseBody struct {
	// The channel groups returned.
	LivePackageChannelGroups []*ListLivePackageChannelGroupsResponseBodyLivePackageChannelGroups `json:"LivePackageChannelGroups,omitempty" xml:"LivePackageChannelGroups,omitempty" type:"Repeated"`
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
	// The request ID.
	//
	// example:
	//
	// 5D87B753-0250-5D9D-B248-D40C3271F864
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The sort order.
	//
	// example:
	//
	// desc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListLivePackageChannelGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLivePackageChannelGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListLivePackageChannelGroupsResponseBody) GetLivePackageChannelGroups() []*ListLivePackageChannelGroupsResponseBodyLivePackageChannelGroups {
	return s.LivePackageChannelGroups
}

func (s *ListLivePackageChannelGroupsResponseBody) GetPageNo() *int64 {
	return s.PageNo
}

func (s *ListLivePackageChannelGroupsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListLivePackageChannelGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLivePackageChannelGroupsResponseBody) GetSortBy() *string {
	return s.SortBy
}

func (s *ListLivePackageChannelGroupsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListLivePackageChannelGroupsResponseBody) SetLivePackageChannelGroups(v []*ListLivePackageChannelGroupsResponseBodyLivePackageChannelGroups) *ListLivePackageChannelGroupsResponseBody {
	s.LivePackageChannelGroups = v
	return s
}

func (s *ListLivePackageChannelGroupsResponseBody) SetPageNo(v int64) *ListLivePackageChannelGroupsResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListLivePackageChannelGroupsResponseBody) SetPageSize(v int64) *ListLivePackageChannelGroupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListLivePackageChannelGroupsResponseBody) SetRequestId(v string) *ListLivePackageChannelGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLivePackageChannelGroupsResponseBody) SetSortBy(v string) *ListLivePackageChannelGroupsResponseBody {
	s.SortBy = &v
	return s
}

func (s *ListLivePackageChannelGroupsResponseBody) SetTotalCount(v int64) *ListLivePackageChannelGroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListLivePackageChannelGroupsResponseBody) Validate() error {
	if s.LivePackageChannelGroups != nil {
		for _, item := range s.LivePackageChannelGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListLivePackageChannelGroupsResponseBodyLivePackageChannelGroups struct {
	// The time when the channel group was created. It is in the `yyyy-MM-ddTHH:mm:ssZ` format and displayed in UTC.
	//
	// example:
	//
	// 2023-04-01T12:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The channel group description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The channel group name.
	//
	// example:
	//
	// testChannelGroup
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The time when the channel group was last modified. It is in the `yyyy-MM-ddTHH:mm:ssZ` format and displayed in UTC.
	//
	// example:
	//
	// 2023-04-02T12:00:00Z
	LastModified *string `json:"LastModified,omitempty" xml:"LastModified,omitempty"`
	// The origin domain.
	//
	// example:
	//
	// origin.example.com
	OriginDomain *string `json:"OriginDomain,omitempty" xml:"OriginDomain,omitempty"`
}

func (s ListLivePackageChannelGroupsResponseBodyLivePackageChannelGroups) String() string {
	return dara.Prettify(s)
}

func (s ListLivePackageChannelGroupsResponseBodyLivePackageChannelGroups) GoString() string {
	return s.String()
}

func (s *ListLivePackageChannelGroupsResponseBodyLivePackageChannelGroups) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListLivePackageChannelGroupsResponseBodyLivePackageChannelGroups) GetDescription() *string {
	return s.Description
}

func (s *ListLivePackageChannelGroupsResponseBodyLivePackageChannelGroups) GetGroupName() *string {
	return s.GroupName
}

func (s *ListLivePackageChannelGroupsResponseBodyLivePackageChannelGroups) GetLastModified() *string {
	return s.LastModified
}

func (s *ListLivePackageChannelGroupsResponseBodyLivePackageChannelGroups) GetOriginDomain() *string {
	return s.OriginDomain
}

func (s *ListLivePackageChannelGroupsResponseBodyLivePackageChannelGroups) SetCreateTime(v string) *ListLivePackageChannelGroupsResponseBodyLivePackageChannelGroups {
	s.CreateTime = &v
	return s
}

func (s *ListLivePackageChannelGroupsResponseBodyLivePackageChannelGroups) SetDescription(v string) *ListLivePackageChannelGroupsResponseBodyLivePackageChannelGroups {
	s.Description = &v
	return s
}

func (s *ListLivePackageChannelGroupsResponseBodyLivePackageChannelGroups) SetGroupName(v string) *ListLivePackageChannelGroupsResponseBodyLivePackageChannelGroups {
	s.GroupName = &v
	return s
}

func (s *ListLivePackageChannelGroupsResponseBodyLivePackageChannelGroups) SetLastModified(v string) *ListLivePackageChannelGroupsResponseBodyLivePackageChannelGroups {
	s.LastModified = &v
	return s
}

func (s *ListLivePackageChannelGroupsResponseBodyLivePackageChannelGroups) SetOriginDomain(v string) *ListLivePackageChannelGroupsResponseBodyLivePackageChannelGroups {
	s.OriginDomain = &v
	return s
}

func (s *ListLivePackageChannelGroupsResponseBodyLivePackageChannelGroups) Validate() error {
	return dara.Validate(s)
}
