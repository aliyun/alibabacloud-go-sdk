// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEventList(v []*DescribeEventsResponseBodyEventList) *DescribeEventsResponseBody
	GetEventList() []*DescribeEventsResponseBodyEventList
	SetMaxResults(v int32) *DescribeEventsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeEventsResponseBody
	GetNextToken() *string
	SetPageNumber(v int32) *DescribeEventsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeEventsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeEventsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeEventsResponseBody
	GetTotalCount() *int32
}

type DescribeEventsResponseBody struct {
	// The list of events on the current page, returned in descending order of the event log time.
	EventList []*DescribeEventsResponseBodyEventList `json:"EventList,omitempty" xml:"EventList,omitempty" type:"Repeated"`
	// A compatibility field. The current version does not set this field. Do not rely on its value.
	//
	// example:
	//
	// 0
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A compatibility field. The current version does not generate a next page token. Do not use this field for pagination.
	//
	// example:
	//
	// unused
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number corresponding to the current response.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of events per page used in the current response.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The unique identifier of the request, which can be used to troubleshoot request issues.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of events that meet all filter conditions within the 30-day retrievable window. Unit: entries.
	//
	// example:
	//
	// 36
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeEventsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEventsResponseBody) GetEventList() []*DescribeEventsResponseBodyEventList {
	return s.EventList
}

func (s *DescribeEventsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeEventsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeEventsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeEventsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeEventsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEventsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeEventsResponseBody) SetEventList(v []*DescribeEventsResponseBodyEventList) *DescribeEventsResponseBody {
	s.EventList = v
	return s
}

func (s *DescribeEventsResponseBody) SetMaxResults(v int32) *DescribeEventsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeEventsResponseBody) SetNextToken(v string) *DescribeEventsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeEventsResponseBody) SetPageNumber(v int32) *DescribeEventsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeEventsResponseBody) SetPageSize(v int32) *DescribeEventsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeEventsResponseBody) SetRequestId(v string) *DescribeEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEventsResponseBody) SetTotalCount(v int32) *DescribeEventsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeEventsResponseBody) Validate() error {
	if s.EventList != nil {
		for _, item := range s.EventList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeEventsResponseBodyEventList struct {
	// The ID of the delivery group associated with the event. This field may be empty if the event is not associated with a delivery group.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The name of the delivery group associated with the event. This field may be empty if the event is not associated with a delivery group.
	//
	// example:
	//
	// R&D workstation group
	AppInstanceGroupName *string `json:"AppInstanceGroupName,omitempty" xml:"AppInstanceGroupName,omitempty"`
	// The event detail text. The content varies based on the `EventType` and may be empty. For example, a session release event may return a JSON text that describes the release cause.
	//
	// example:
	//
	// {"Reason":"ByUser"}
	DetailInfo *string `json:"DetailInfo,omitempty" xml:"DetailInfo,omitempty"`
	// The unique identifier of the event.
	//
	// example:
	//
	// evt-9ciijz60n4xsv****
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The event type identifier. Different products may return different values. You can use this value for subsequent exact filtering.
	//
	// example:
	//
	// SessionDisconnected
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The private IP address of the member ENI of the cloud resource associated with the event. This field may be empty if it is not applicable.
	//
	// example:
	//
	// 192.168.1.10
	MemberEniIp *string `json:"MemberEniIp,omitempty" xml:"MemberEniIp,omitempty"`
	// The time when the event occurred. Format: `yyyy-MM-dd HH:mm:ss`. Time zone: UTC+8 (Asia/Shanghai).
	//
	// example:
	//
	// 2026-09-29 14:17:33
	OccurredTime *string `json:"OccurredTime,omitempty" xml:"OccurredTime,omitempty"`
	// The ID of the office network associated with the event. This field may be empty if it is not applicable.
	//
	// example:
	//
	// cn-shanghai+dir-123456****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The product type to which the event belongs. This field is an open string and may be extended as supported event sources expand. Callers should not restrict the value to the products currently listed in the `ProductType` request parameter.
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The ID of the end user who generated the event. This field may be empty if the event is not associated with an end user.
	//
	// example:
	//
	// user01
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeEventsResponseBodyEventList) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventsResponseBodyEventList) GoString() string {
	return s.String()
}

func (s *DescribeEventsResponseBodyEventList) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *DescribeEventsResponseBodyEventList) GetAppInstanceGroupName() *string {
	return s.AppInstanceGroupName
}

func (s *DescribeEventsResponseBodyEventList) GetDetailInfo() *string {
	return s.DetailInfo
}

func (s *DescribeEventsResponseBodyEventList) GetEventId() *string {
	return s.EventId
}

func (s *DescribeEventsResponseBodyEventList) GetEventType() *string {
	return s.EventType
}

func (s *DescribeEventsResponseBodyEventList) GetMemberEniIp() *string {
	return s.MemberEniIp
}

func (s *DescribeEventsResponseBodyEventList) GetOccurredTime() *string {
	return s.OccurredTime
}

func (s *DescribeEventsResponseBodyEventList) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *DescribeEventsResponseBodyEventList) GetProductType() *string {
	return s.ProductType
}

func (s *DescribeEventsResponseBodyEventList) GetUserId() *string {
	return s.UserId
}

func (s *DescribeEventsResponseBodyEventList) SetAppInstanceGroupId(v string) *DescribeEventsResponseBodyEventList {
	s.AppInstanceGroupId = &v
	return s
}

func (s *DescribeEventsResponseBodyEventList) SetAppInstanceGroupName(v string) *DescribeEventsResponseBodyEventList {
	s.AppInstanceGroupName = &v
	return s
}

func (s *DescribeEventsResponseBodyEventList) SetDetailInfo(v string) *DescribeEventsResponseBodyEventList {
	s.DetailInfo = &v
	return s
}

func (s *DescribeEventsResponseBodyEventList) SetEventId(v string) *DescribeEventsResponseBodyEventList {
	s.EventId = &v
	return s
}

func (s *DescribeEventsResponseBodyEventList) SetEventType(v string) *DescribeEventsResponseBodyEventList {
	s.EventType = &v
	return s
}

func (s *DescribeEventsResponseBodyEventList) SetMemberEniIp(v string) *DescribeEventsResponseBodyEventList {
	s.MemberEniIp = &v
	return s
}

func (s *DescribeEventsResponseBodyEventList) SetOccurredTime(v string) *DescribeEventsResponseBodyEventList {
	s.OccurredTime = &v
	return s
}

func (s *DescribeEventsResponseBodyEventList) SetOfficeSiteId(v string) *DescribeEventsResponseBodyEventList {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeEventsResponseBodyEventList) SetProductType(v string) *DescribeEventsResponseBodyEventList {
	s.ProductType = &v
	return s
}

func (s *DescribeEventsResponseBodyEventList) SetUserId(v string) *DescribeEventsResponseBodyEventList {
	s.UserId = &v
	return s
}

func (s *DescribeEventsResponseBodyEventList) Validate() error {
	return dara.Validate(s)
}
