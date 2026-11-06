// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppInstanceGroupId(v string) *DescribeEventsRequest
	GetAppInstanceGroupId() *string
	SetAppInstanceGroupName(v string) *DescribeEventsRequest
	GetAppInstanceGroupName() *string
	SetAppInstanceId(v string) *DescribeEventsRequest
	GetAppInstanceId() *string
	SetBizRegionId(v string) *DescribeEventsRequest
	GetBizRegionId() *string
	SetEndTime(v string) *DescribeEventsRequest
	GetEndTime() *string
	SetEventId(v string) *DescribeEventsRequest
	GetEventId() *string
	SetEventType(v string) *DescribeEventsRequest
	GetEventType() *string
	SetMaxResults(v int32) *DescribeEventsRequest
	GetMaxResults() *int32
	SetMemberEniIp(v string) *DescribeEventsRequest
	GetMemberEniIp() *string
	SetNextToken(v string) *DescribeEventsRequest
	GetNextToken() *string
	SetOfficeSiteId(v string) *DescribeEventsRequest
	GetOfficeSiteId() *string
	SetPageNumber(v int32) *DescribeEventsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeEventsRequest
	GetPageSize() *int32
	SetProductType(v string) *DescribeEventsRequest
	GetProductType() *string
	SetStartTime(v string) *DescribeEventsRequest
	GetStartTime() *string
	SetUserId(v string) *DescribeEventsRequest
	GetUserId() *string
}

type DescribeEventsRequest struct {
	// The delivery group ID, used to filter events of a specific delivery group. If you do not specify this parameter, events are not filtered by delivery group. You can call the [ListAppInstanceGroup](~~ListAppInstanceGroup~~) operation to obtain the ID.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The delivery group name. If the name contains only letters, digits, underscores, or hyphens, events are filtered by name prefix. If the name contains spaces or other special characters, events are filtered by the exact name. If you do not specify this parameter, events are not filtered by name.
	//
	// example:
	//
	// R&D workstation group
	AppInstanceGroupName *string `json:"AppInstanceGroupName,omitempty" xml:"AppInstanceGroupName,omitempty"`
	// The application instance ID, used to filter events associated with a specific instance. If you do not specify this parameter, events are not filtered by application instance. To obtain this value, first call the [ListAppInstanceGroup](~~ListAppInstanceGroup~~) operation to obtain the delivery group ID, and then use the ID as `AppInstanceGroupId` to call the [ListAppInstances](~~ListAppInstances~~) operation.
	//
	// example:
	//
	// ai-azn3kmwruh1vl****
	AppInstanceId *string `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	// The business region ID of the event, used for exact filtering. If you do not specify this parameter, events are not filtered by business region. You can call the [ListRegions](~~ListRegions~~) operation to obtain the ID.
	//
	// example:
	//
	// cn-shanghai
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The end time of the query. **Required**. The format is the same as that of `StartTime`. We recommend that you always include the time zone offset. The query includes this time, and the end time cannot be earlier than the `StartTime`. A time range later than the call time does not generate future events.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-09-29T15:59:59Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The event ID, used to query a specific event. If you do not specify this parameter, events are not filtered by event ID.
	//
	// example:
	//
	// evt-9ciijz60n4xsv****
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The event type identifier, used to filter events. Different products may generate different event types. This operation does not restrict the value to a fixed enumeration. You should use the EventType values retrieved from EventList as filter values.
	//
	// example:
	//
	// SessionDisconnected
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// A compatibility field. The current version does not read this parameter. Use `PageSize` to control the number of records on each page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The private IP address of the member elastic network interface (ENI) of the cloud resource, used to filter related events. If you do not specify this parameter, events are not filtered by IP address.
	//
	// example:
	//
	// 192.168.1.10
	MemberEniIp *string `json:"MemberEniIp,omitempty" xml:"MemberEniIp,omitempty"`
	// A compatibility field. The current version does not read this parameter and does not support token-based pagination. Use `PageNumber`.
	//
	// example:
	//
	// unused
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The office network ID, used to filter events generated within a specific office network. If you do not specify this parameter, events are not filtered by office network.
	//
	// example:
	//
	// cn-shanghai+dir-123456****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The page number to query. Pages start from 1. Default value: 1. This parameter is used together with `PageSize`.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of events to return on each page. Default value: 10. This operation does not declare a fixed upper limit. We recommend that you set this parameter to a reasonable positive integer based on your actual needs.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The product type. The value is case-insensitive. If you specify an invalid value, the request fails. If you do not specify this parameter, events are not filtered by product type. Only the following public product values are supported:
	//
	// - `CloudApp`: Wuying Cloud Application.
	//
	// - `CloudBrowser`: Cloud Browser.
	//
	// - `WuyingServer`: Enterprise Edition Workstation.
	//
	// - `WuyingWorkstation`: Personal Edition Linggou Container Workstation.
	//
	// - `WuyingWorkstationTeam`: Linggou Team Edition Container Workstation.
	//
	// - `WuyingWorkstationBusiness`: Linggou Dedicated Edition Container Workstation.
	//
	// - `AndroidCloud`: Cloud Phone.
	//
	// - `AIAgent`: AgentBay (AI agent).
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The start time of the query. **Required**. Use the ISO 8601 format: `yyyy-MM-dd\\"T\\"HH:mm:ss`. You can include three-digit milliseconds and use `Z` or `±HH:mm` to specify the time zone. We recommend that you always include the time zone offset. The query includes this time, and the start time must be earlier than or equal to the `EndTime`. The service only retains and retrieves events that occurred within 30 days before the call time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-09-28T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The end user ID, used to filter events generated by a specific user. If you do not specify this parameter, events are not filtered by user.
	//
	// example:
	//
	// user01
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventsRequest) GoString() string {
	return s.String()
}

func (s *DescribeEventsRequest) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *DescribeEventsRequest) GetAppInstanceGroupName() *string {
	return s.AppInstanceGroupName
}

func (s *DescribeEventsRequest) GetAppInstanceId() *string {
	return s.AppInstanceId
}

func (s *DescribeEventsRequest) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *DescribeEventsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeEventsRequest) GetEventId() *string {
	return s.EventId
}

func (s *DescribeEventsRequest) GetEventType() *string {
	return s.EventType
}

func (s *DescribeEventsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeEventsRequest) GetMemberEniIp() *string {
	return s.MemberEniIp
}

func (s *DescribeEventsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeEventsRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *DescribeEventsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeEventsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeEventsRequest) GetProductType() *string {
	return s.ProductType
}

func (s *DescribeEventsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeEventsRequest) GetUserId() *string {
	return s.UserId
}

func (s *DescribeEventsRequest) SetAppInstanceGroupId(v string) *DescribeEventsRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *DescribeEventsRequest) SetAppInstanceGroupName(v string) *DescribeEventsRequest {
	s.AppInstanceGroupName = &v
	return s
}

func (s *DescribeEventsRequest) SetAppInstanceId(v string) *DescribeEventsRequest {
	s.AppInstanceId = &v
	return s
}

func (s *DescribeEventsRequest) SetBizRegionId(v string) *DescribeEventsRequest {
	s.BizRegionId = &v
	return s
}

func (s *DescribeEventsRequest) SetEndTime(v string) *DescribeEventsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeEventsRequest) SetEventId(v string) *DescribeEventsRequest {
	s.EventId = &v
	return s
}

func (s *DescribeEventsRequest) SetEventType(v string) *DescribeEventsRequest {
	s.EventType = &v
	return s
}

func (s *DescribeEventsRequest) SetMaxResults(v int32) *DescribeEventsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeEventsRequest) SetMemberEniIp(v string) *DescribeEventsRequest {
	s.MemberEniIp = &v
	return s
}

func (s *DescribeEventsRequest) SetNextToken(v string) *DescribeEventsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeEventsRequest) SetOfficeSiteId(v string) *DescribeEventsRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeEventsRequest) SetPageNumber(v int32) *DescribeEventsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeEventsRequest) SetPageSize(v int32) *DescribeEventsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeEventsRequest) SetProductType(v string) *DescribeEventsRequest {
	s.ProductType = &v
	return s
}

func (s *DescribeEventsRequest) SetStartTime(v string) *DescribeEventsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeEventsRequest) SetUserId(v string) *DescribeEventsRequest {
	s.UserId = &v
	return s
}

func (s *DescribeEventsRequest) Validate() error {
	return dara.Validate(s)
}
