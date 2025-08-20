// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStackEventsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEvents(v []*ListStackEventsResponseBodyEvents) *ListStackEventsResponseBody
	GetEvents() []*ListStackEventsResponseBodyEvents
	SetPageNumber(v int32) *ListStackEventsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListStackEventsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListStackEventsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListStackEventsResponseBody
	GetTotalCount() *int32
}

type ListStackEventsResponseBody struct {
	// The events.
	Events []*ListStackEventsResponseBodyEvents `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	// The page number of the returned page.\\
	//
	// Pages start from page 1.\\
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.\\
	//
	// Maximum value: 50.\\
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B288A0BE-D927-4888-B0F7-B35EF84B6E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned events.
	//
	// example:
	//
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListStackEventsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListStackEventsResponseBody) GoString() string {
	return s.String()
}

func (s *ListStackEventsResponseBody) GetEvents() []*ListStackEventsResponseBodyEvents {
	return s.Events
}

func (s *ListStackEventsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListStackEventsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListStackEventsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListStackEventsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListStackEventsResponseBody) SetEvents(v []*ListStackEventsResponseBodyEvents) *ListStackEventsResponseBody {
	s.Events = v
	return s
}

func (s *ListStackEventsResponseBody) SetPageNumber(v int32) *ListStackEventsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListStackEventsResponseBody) SetPageSize(v int32) *ListStackEventsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListStackEventsResponseBody) SetRequestId(v string) *ListStackEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListStackEventsResponseBody) SetTotalCount(v int32) *ListStackEventsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListStackEventsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListStackEventsResponseBodyEvents struct {
	// The time when the event was created. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-08-01T04:07:39
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The event ID.
	//
	// example:
	//
	// 0086612d-77cf-4056-b0b5-ff8e94ad****
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The logical ID of the resource. The logical ID indicates the name of the resource that is defined in the template.
	//
	// example:
	//
	// WebServer
	LogicalResourceId *string `json:"LogicalResourceId,omitempty" xml:"LogicalResourceId,omitempty"`
	// The physical ID of the resource.
	//
	// example:
	//
	// i-m5e3tfdbinchnexh****
	PhysicalResourceId *string `json:"PhysicalResourceId,omitempty" xml:"PhysicalResourceId,omitempty"`
	// The resource type.
	//
	// example:
	//
	// ALIYUN::ECS::Instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The stack ID.
	//
	// example:
	//
	// 4a6c9851-3b0f-4f5f-b4ca-a14bf691****
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	// The stack name.
	//
	// example:
	//
	// StackName
	StackName *string `json:"StackName,omitempty" xml:"StackName,omitempty"`
	// The state of the resource.
	//
	// example:
	//
	// CREATE_COMPLETE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The reason why the resource is in the current state.
	//
	// example:
	//
	// state changed
	StatusReason *string `json:"StatusReason,omitempty" xml:"StatusReason,omitempty"`
}

func (s ListStackEventsResponseBodyEvents) String() string {
	return dara.Prettify(s)
}

func (s ListStackEventsResponseBodyEvents) GoString() string {
	return s.String()
}

func (s *ListStackEventsResponseBodyEvents) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListStackEventsResponseBodyEvents) GetEventId() *string {
	return s.EventId
}

func (s *ListStackEventsResponseBodyEvents) GetLogicalResourceId() *string {
	return s.LogicalResourceId
}

func (s *ListStackEventsResponseBodyEvents) GetPhysicalResourceId() *string {
	return s.PhysicalResourceId
}

func (s *ListStackEventsResponseBodyEvents) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListStackEventsResponseBodyEvents) GetStackId() *string {
	return s.StackId
}

func (s *ListStackEventsResponseBodyEvents) GetStackName() *string {
	return s.StackName
}

func (s *ListStackEventsResponseBodyEvents) GetStatus() *string {
	return s.Status
}

func (s *ListStackEventsResponseBodyEvents) GetStatusReason() *string {
	return s.StatusReason
}

func (s *ListStackEventsResponseBodyEvents) SetCreateTime(v string) *ListStackEventsResponseBodyEvents {
	s.CreateTime = &v
	return s
}

func (s *ListStackEventsResponseBodyEvents) SetEventId(v string) *ListStackEventsResponseBodyEvents {
	s.EventId = &v
	return s
}

func (s *ListStackEventsResponseBodyEvents) SetLogicalResourceId(v string) *ListStackEventsResponseBodyEvents {
	s.LogicalResourceId = &v
	return s
}

func (s *ListStackEventsResponseBodyEvents) SetPhysicalResourceId(v string) *ListStackEventsResponseBodyEvents {
	s.PhysicalResourceId = &v
	return s
}

func (s *ListStackEventsResponseBodyEvents) SetResourceType(v string) *ListStackEventsResponseBodyEvents {
	s.ResourceType = &v
	return s
}

func (s *ListStackEventsResponseBodyEvents) SetStackId(v string) *ListStackEventsResponseBodyEvents {
	s.StackId = &v
	return s
}

func (s *ListStackEventsResponseBodyEvents) SetStackName(v string) *ListStackEventsResponseBodyEvents {
	s.StackName = &v
	return s
}

func (s *ListStackEventsResponseBodyEvents) SetStatus(v string) *ListStackEventsResponseBodyEvents {
	s.Status = &v
	return s
}

func (s *ListStackEventsResponseBodyEvents) SetStatusReason(v string) *ListStackEventsResponseBodyEvents {
	s.StatusReason = &v
	return s
}

func (s *ListStackEventsResponseBodyEvents) Validate() error {
	return dara.Validate(s)
}
