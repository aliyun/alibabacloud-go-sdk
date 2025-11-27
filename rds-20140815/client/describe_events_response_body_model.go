// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEventItems(v *DescribeEventsResponseBodyEventItems) *DescribeEventsResponseBody
	GetEventItems() *DescribeEventsResponseBodyEventItems
	SetPageNumber(v int32) *DescribeEventsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeEventsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeEventsResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int32) *DescribeEventsResponseBody
	GetTotalRecordCount() *int32
}

type DescribeEventsResponseBody struct {
	// The events.
	EventItems *DescribeEventsResponseBodyEventItems `json:"EventItems,omitempty" xml:"EventItems,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A103039D-B1B2-4C57-B989-7D7C0DA95426
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 40
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeEventsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEventsResponseBody) GetEventItems() *DescribeEventsResponseBodyEventItems {
	return s.EventItems
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

func (s *DescribeEventsResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeEventsResponseBody) SetEventItems(v *DescribeEventsResponseBodyEventItems) *DescribeEventsResponseBody {
	s.EventItems = v
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

func (s *DescribeEventsResponseBody) SetTotalRecordCount(v int32) *DescribeEventsResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeEventsResponseBody) Validate() error {
	if s.EventItems != nil {
		if err := s.EventItems.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeEventsResponseBodyEventItems struct {
	EventItems []*DescribeEventsResponseBodyEventItemsEventItems `json:"EventItems,omitempty" xml:"EventItems,omitempty" type:"Repeated"`
}

func (s DescribeEventsResponseBodyEventItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventsResponseBodyEventItems) GoString() string {
	return s.String()
}

func (s *DescribeEventsResponseBodyEventItems) GetEventItems() []*DescribeEventsResponseBodyEventItemsEventItems {
	return s.EventItems
}

func (s *DescribeEventsResponseBodyEventItems) SetEventItems(v []*DescribeEventsResponseBodyEventItemsEventItems) *DescribeEventsResponseBodyEventItems {
	s.EventItems = v
	return s
}

func (s *DescribeEventsResponseBodyEventItems) Validate() error {
	if s.EventItems != nil {
		for _, item := range s.EventItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeEventsResponseBodyEventItemsEventItems struct {
	// The ID of the user who executed the event.
	//
	// example:
	//
	// 22973492**********
	CallerUid *int64 `json:"CallerUid,omitempty" xml:"CallerUid,omitempty"`
	// The event ID.
	//
	// example:
	//
	// 11000053
	EventId *int32 `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The event name.
	//
	// example:
	//
	// ModifySecurityIPList
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// The request or context parameters of the event.
	//
	// example:
	//
	// {\\"Domain\\": \\"rds-inc-share.aliyuncs.com\\", \\"Api\\": \\"ReleaseInstancePublicConnection\\"}
	EventPayload *string `json:"EventPayload,omitempty" xml:"EventPayload,omitempty"`
	// The source of the event.
	//
	// example:
	//
	// FROM_USER
	EventReason *string `json:"EventReason,omitempty" xml:"EventReason,omitempty"`
	// The time when the event was recorded. The time is slightly later than the time the event occurred.
	//
	// example:
	//
	// 2019-08-20T01:12:49Z
	EventRecordTime *string `json:"EventRecordTime,omitempty" xml:"EventRecordTime,omitempty"`
	// The time when the event occurred.
	//
	// example:
	//
	// 2019-08-20T01:08:22Z
	EventTime *string `json:"EventTime,omitempty" xml:"EventTime,omitempty"`
	// The event type.
	//
	// example:
	//
	// NetworkManagement
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The type of the user who executed the event.
	//
	// example:
	//
	// SYSTEM
	EventUserType *string `json:"EventUserType,omitempty" xml:"EventUserType,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the resource associated with the event. Only instance IDs are supported for this parameter.
	//
	// example:
	//
	// rm-bp1z3065m9976ix8a
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The type of the resource associated with the event. Only instances are supported for this parameter.
	//
	// example:
	//
	// instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s DescribeEventsResponseBodyEventItemsEventItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventsResponseBodyEventItemsEventItems) GoString() string {
	return s.String()
}

func (s *DescribeEventsResponseBodyEventItemsEventItems) GetCallerUid() *int64 {
	return s.CallerUid
}

func (s *DescribeEventsResponseBodyEventItemsEventItems) GetEventId() *int32 {
	return s.EventId
}

func (s *DescribeEventsResponseBodyEventItemsEventItems) GetEventName() *string {
	return s.EventName
}

func (s *DescribeEventsResponseBodyEventItemsEventItems) GetEventPayload() *string {
	return s.EventPayload
}

func (s *DescribeEventsResponseBodyEventItemsEventItems) GetEventReason() *string {
	return s.EventReason
}

func (s *DescribeEventsResponseBodyEventItemsEventItems) GetEventRecordTime() *string {
	return s.EventRecordTime
}

func (s *DescribeEventsResponseBodyEventItemsEventItems) GetEventTime() *string {
	return s.EventTime
}

func (s *DescribeEventsResponseBodyEventItemsEventItems) GetEventType() *string {
	return s.EventType
}

func (s *DescribeEventsResponseBodyEventItemsEventItems) GetEventUserType() *string {
	return s.EventUserType
}

func (s *DescribeEventsResponseBodyEventItemsEventItems) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeEventsResponseBodyEventItemsEventItems) GetResourceName() *string {
	return s.ResourceName
}

func (s *DescribeEventsResponseBodyEventItemsEventItems) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeEventsResponseBodyEventItemsEventItems) SetCallerUid(v int64) *DescribeEventsResponseBodyEventItemsEventItems {
	s.CallerUid = &v
	return s
}

func (s *DescribeEventsResponseBodyEventItemsEventItems) SetEventId(v int32) *DescribeEventsResponseBodyEventItemsEventItems {
	s.EventId = &v
	return s
}

func (s *DescribeEventsResponseBodyEventItemsEventItems) SetEventName(v string) *DescribeEventsResponseBodyEventItemsEventItems {
	s.EventName = &v
	return s
}

func (s *DescribeEventsResponseBodyEventItemsEventItems) SetEventPayload(v string) *DescribeEventsResponseBodyEventItemsEventItems {
	s.EventPayload = &v
	return s
}

func (s *DescribeEventsResponseBodyEventItemsEventItems) SetEventReason(v string) *DescribeEventsResponseBodyEventItemsEventItems {
	s.EventReason = &v
	return s
}

func (s *DescribeEventsResponseBodyEventItemsEventItems) SetEventRecordTime(v string) *DescribeEventsResponseBodyEventItemsEventItems {
	s.EventRecordTime = &v
	return s
}

func (s *DescribeEventsResponseBodyEventItemsEventItems) SetEventTime(v string) *DescribeEventsResponseBodyEventItemsEventItems {
	s.EventTime = &v
	return s
}

func (s *DescribeEventsResponseBodyEventItemsEventItems) SetEventType(v string) *DescribeEventsResponseBodyEventItemsEventItems {
	s.EventType = &v
	return s
}

func (s *DescribeEventsResponseBodyEventItemsEventItems) SetEventUserType(v string) *DescribeEventsResponseBodyEventItemsEventItems {
	s.EventUserType = &v
	return s
}

func (s *DescribeEventsResponseBodyEventItemsEventItems) SetRegionId(v string) *DescribeEventsResponseBodyEventItemsEventItems {
	s.RegionId = &v
	return s
}

func (s *DescribeEventsResponseBodyEventItemsEventItems) SetResourceName(v string) *DescribeEventsResponseBodyEventItemsEventItems {
	s.ResourceName = &v
	return s
}

func (s *DescribeEventsResponseBodyEventItemsEventItems) SetResourceType(v string) *DescribeEventsResponseBodyEventItemsEventItems {
	s.ResourceType = &v
	return s
}

func (s *DescribeEventsResponseBodyEventItemsEventItems) Validate() error {
	return dara.Validate(s)
}
