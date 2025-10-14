// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEventItems(v []*DescribeEventsResponseBodyEventItems) *DescribeEventsResponseBody
	GetEventItems() []*DescribeEventsResponseBodyEventItems
	SetPageNumber(v int64) *DescribeEventsResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeEventsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeEventsResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int64) *DescribeEventsResponseBody
	GetTotalRecordCount() *int64
}

type DescribeEventsResponseBody struct {
	EventItems []*DescribeEventsResponseBodyEventItems `json:"EventItems,omitempty" xml:"EventItems,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 4748127A-6D50-432C-B635-433467074C27
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 20
	TotalRecordCount *int64 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeEventsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEventsResponseBody) GetEventItems() []*DescribeEventsResponseBodyEventItems {
	return s.EventItems
}

func (s *DescribeEventsResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeEventsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeEventsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEventsResponseBody) GetTotalRecordCount() *int64 {
	return s.TotalRecordCount
}

func (s *DescribeEventsResponseBody) SetEventItems(v []*DescribeEventsResponseBodyEventItems) *DescribeEventsResponseBody {
	s.EventItems = v
	return s
}

func (s *DescribeEventsResponseBody) SetPageNumber(v int64) *DescribeEventsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeEventsResponseBody) SetPageSize(v int64) *DescribeEventsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeEventsResponseBody) SetRequestId(v string) *DescribeEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEventsResponseBody) SetTotalRecordCount(v int64) *DescribeEventsResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeEventsResponseBody) Validate() error {
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

type DescribeEventsResponseBodyEventItems struct {
	// example:
	//
	// 50421290
	EventId *int64 `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// example:
	//
	// ModifySecurityIps
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// example:
	//
	// {\\"Domain\\": \\"rds-cn-hangzhou.aliyuncs.com\\"}
	EventPayload *string `json:"EventPayload,omitempty" xml:"EventPayload,omitempty"`
	// example:
	//
	// FROM_USER
	EventReason *string `json:"EventReason,omitempty" xml:"EventReason,omitempty"`
	// example:
	//
	// 2021-10-15T06:39:49Z
	EventRecordTime *string `json:"EventRecordTime,omitempty" xml:"EventRecordTime,omitempty"`
	// example:
	//
	// 2021-10-15T06:35:00Z
	EventTime *string `json:"EventTime,omitempty" xml:"EventTime,omitempty"`
	// example:
	//
	// SecurityManagement
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// example:
	//
	// USRE
	EventUserType *string `json:"EventUserType,omitempty" xml:"EventUserType,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// px-bp1v8udesc89g156g
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// example:
	//
	// instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s DescribeEventsResponseBodyEventItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventsResponseBodyEventItems) GoString() string {
	return s.String()
}

func (s *DescribeEventsResponseBodyEventItems) GetEventId() *int64 {
	return s.EventId
}

func (s *DescribeEventsResponseBodyEventItems) GetEventName() *string {
	return s.EventName
}

func (s *DescribeEventsResponseBodyEventItems) GetEventPayload() *string {
	return s.EventPayload
}

func (s *DescribeEventsResponseBodyEventItems) GetEventReason() *string {
	return s.EventReason
}

func (s *DescribeEventsResponseBodyEventItems) GetEventRecordTime() *string {
	return s.EventRecordTime
}

func (s *DescribeEventsResponseBodyEventItems) GetEventTime() *string {
	return s.EventTime
}

func (s *DescribeEventsResponseBodyEventItems) GetEventType() *string {
	return s.EventType
}

func (s *DescribeEventsResponseBodyEventItems) GetEventUserType() *string {
	return s.EventUserType
}

func (s *DescribeEventsResponseBodyEventItems) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeEventsResponseBodyEventItems) GetResourceName() *string {
	return s.ResourceName
}

func (s *DescribeEventsResponseBodyEventItems) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeEventsResponseBodyEventItems) SetEventId(v int64) *DescribeEventsResponseBodyEventItems {
	s.EventId = &v
	return s
}

func (s *DescribeEventsResponseBodyEventItems) SetEventName(v string) *DescribeEventsResponseBodyEventItems {
	s.EventName = &v
	return s
}

func (s *DescribeEventsResponseBodyEventItems) SetEventPayload(v string) *DescribeEventsResponseBodyEventItems {
	s.EventPayload = &v
	return s
}

func (s *DescribeEventsResponseBodyEventItems) SetEventReason(v string) *DescribeEventsResponseBodyEventItems {
	s.EventReason = &v
	return s
}

func (s *DescribeEventsResponseBodyEventItems) SetEventRecordTime(v string) *DescribeEventsResponseBodyEventItems {
	s.EventRecordTime = &v
	return s
}

func (s *DescribeEventsResponseBodyEventItems) SetEventTime(v string) *DescribeEventsResponseBodyEventItems {
	s.EventTime = &v
	return s
}

func (s *DescribeEventsResponseBodyEventItems) SetEventType(v string) *DescribeEventsResponseBodyEventItems {
	s.EventType = &v
	return s
}

func (s *DescribeEventsResponseBodyEventItems) SetEventUserType(v string) *DescribeEventsResponseBodyEventItems {
	s.EventUserType = &v
	return s
}

func (s *DescribeEventsResponseBodyEventItems) SetRegionId(v string) *DescribeEventsResponseBodyEventItems {
	s.RegionId = &v
	return s
}

func (s *DescribeEventsResponseBodyEventItems) SetResourceName(v string) *DescribeEventsResponseBodyEventItems {
	s.ResourceName = &v
	return s
}

func (s *DescribeEventsResponseBodyEventItems) SetResourceType(v string) *DescribeEventsResponseBodyEventItems {
	s.ResourceType = &v
	return s
}

func (s *DescribeEventsResponseBodyEventItems) Validate() error {
	return dara.Validate(s)
}
