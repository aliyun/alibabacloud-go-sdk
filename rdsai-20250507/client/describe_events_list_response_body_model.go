// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventsListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEventCodeCounts(v string) *DescribeEventsListResponseBody
	GetEventCodeCounts() *string
	SetEvents(v []*DescribeEventsListResponseBodyEvents) *DescribeEventsListResponseBody
	GetEvents() []*DescribeEventsListResponseBodyEvents
	SetPageCount(v int64) *DescribeEventsListResponseBody
	GetPageCount() *int64
	SetPageNumber(v int64) *DescribeEventsListResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeEventsListResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeEventsListResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeEventsListResponseBody
	GetTotalCount() *int64
	SetTotalPages(v int64) *DescribeEventsListResponseBody
	GetTotalPages() *int64
}

type DescribeEventsListResponseBody struct {
	// The count for each event code.
	//
	// example:
	//
	// OtherException:3,MysqlIOException:1
	EventCodeCounts *string `json:"EventCodeCounts,omitempty" xml:"EventCodeCounts,omitempty"`
	// A list of events.
	Events []*DescribeEventsListResponseBodyEvents `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	// The number of pages returned.
	//
	// example:
	//
	// 1
	PageCount *int64 `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
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
	// FE9C65D7-930F-57A5-A207-8C396329****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 20
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The total number of pages.
	//
	// example:
	//
	// 10
	TotalPages *int64 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s DescribeEventsListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventsListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEventsListResponseBody) GetEventCodeCounts() *string {
	return s.EventCodeCounts
}

func (s *DescribeEventsListResponseBody) GetEvents() []*DescribeEventsListResponseBodyEvents {
	return s.Events
}

func (s *DescribeEventsListResponseBody) GetPageCount() *int64 {
	return s.PageCount
}

func (s *DescribeEventsListResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeEventsListResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeEventsListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEventsListResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeEventsListResponseBody) GetTotalPages() *int64 {
	return s.TotalPages
}

func (s *DescribeEventsListResponseBody) SetEventCodeCounts(v string) *DescribeEventsListResponseBody {
	s.EventCodeCounts = &v
	return s
}

func (s *DescribeEventsListResponseBody) SetEvents(v []*DescribeEventsListResponseBodyEvents) *DescribeEventsListResponseBody {
	s.Events = v
	return s
}

func (s *DescribeEventsListResponseBody) SetPageCount(v int64) *DescribeEventsListResponseBody {
	s.PageCount = &v
	return s
}

func (s *DescribeEventsListResponseBody) SetPageNumber(v int64) *DescribeEventsListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeEventsListResponseBody) SetPageSize(v int64) *DescribeEventsListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeEventsListResponseBody) SetRequestId(v string) *DescribeEventsListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEventsListResponseBody) SetTotalCount(v int64) *DescribeEventsListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeEventsListResponseBody) SetTotalPages(v int64) *DescribeEventsListResponseBody {
	s.TotalPages = &v
	return s
}

func (s *DescribeEventsListResponseBody) Validate() error {
	if s.Events != nil {
		for _, item := range s.Events {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeEventsListResponseBodyEvents struct {
	// The event code.
	//
	// example:
	//
	// ExceptionEvent
	EventCode *string `json:"EventCode,omitempty" xml:"EventCode,omitempty"`
	// The event status.
	//
	// example:
	//
	// None
	EventStatus *string `json:"EventStatus,omitempty" xml:"EventStatus,omitempty"`
	// A list of event times.
	EventTimeList []*string `json:"EventTimeList,omitempty" xml:"EventTimeList,omitempty" type:"Repeated"`
	// The instance description.
	//
	// example:
	//
	// 测试实例
	InstanceDescription *string `json:"InstanceDescription,omitempty" xml:"InstanceDescription,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rm-2zecnb327gp36****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The recovery time.
	//
	// example:
	//
	// 2025-07-23T02:11:07Z
	RecoveryTime *string `json:"RecoveryTime,omitempty" xml:"RecoveryTime,omitempty"`
	// The region ID.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeEventsListResponseBodyEvents) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventsListResponseBodyEvents) GoString() string {
	return s.String()
}

func (s *DescribeEventsListResponseBodyEvents) GetEventCode() *string {
	return s.EventCode
}

func (s *DescribeEventsListResponseBodyEvents) GetEventStatus() *string {
	return s.EventStatus
}

func (s *DescribeEventsListResponseBodyEvents) GetEventTimeList() []*string {
	return s.EventTimeList
}

func (s *DescribeEventsListResponseBodyEvents) GetInstanceDescription() *string {
	return s.InstanceDescription
}

func (s *DescribeEventsListResponseBodyEvents) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeEventsListResponseBodyEvents) GetRecoveryTime() *string {
	return s.RecoveryTime
}

func (s *DescribeEventsListResponseBodyEvents) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeEventsListResponseBodyEvents) SetEventCode(v string) *DescribeEventsListResponseBodyEvents {
	s.EventCode = &v
	return s
}

func (s *DescribeEventsListResponseBodyEvents) SetEventStatus(v string) *DescribeEventsListResponseBodyEvents {
	s.EventStatus = &v
	return s
}

func (s *DescribeEventsListResponseBodyEvents) SetEventTimeList(v []*string) *DescribeEventsListResponseBodyEvents {
	s.EventTimeList = v
	return s
}

func (s *DescribeEventsListResponseBodyEvents) SetInstanceDescription(v string) *DescribeEventsListResponseBodyEvents {
	s.InstanceDescription = &v
	return s
}

func (s *DescribeEventsListResponseBodyEvents) SetInstanceId(v string) *DescribeEventsListResponseBodyEvents {
	s.InstanceId = &v
	return s
}

func (s *DescribeEventsListResponseBodyEvents) SetRecoveryTime(v string) *DescribeEventsListResponseBodyEvents {
	s.RecoveryTime = &v
	return s
}

func (s *DescribeEventsListResponseBodyEvents) SetRegionId(v string) *DescribeEventsListResponseBodyEvents {
	s.RegionId = &v
	return s
}

func (s *DescribeEventsListResponseBodyEvents) Validate() error {
	return dara.Validate(s)
}
