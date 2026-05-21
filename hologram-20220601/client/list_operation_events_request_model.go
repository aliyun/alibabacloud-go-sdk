// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOperationEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventName(v string) *ListOperationEventsRequest
	GetEventName() *string
	SetEventNameDesc(v bool) *ListOperationEventsRequest
	GetEventNameDesc() *bool
	SetEventType(v string) *ListOperationEventsRequest
	GetEventType() *string
	SetInstanceId(v string) *ListOperationEventsRequest
	GetInstanceId() *string
	SetPageNumber(v int64) *ListOperationEventsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListOperationEventsRequest
	GetPageSize() *int64
	SetScheduleTimeDesc(v bool) *ListOperationEventsRequest
	GetScheduleTimeDesc() *bool
	SetState(v string) *ListOperationEventsRequest
	GetState() *string
}

type ListOperationEventsRequest struct {
	// example:
	//
	// COLD_UPGRADE
	EventName *string `json:"eventName,omitempty" xml:"eventName,omitempty"`
	// example:
	//
	// true
	EventNameDesc *bool `json:"eventNameDesc,omitempty" xml:"eventNameDesc,omitempty"`
	// example:
	//
	// SYSTEM_MAINTENANCE
	EventType *string `json:"eventType,omitempty" xml:"eventType,omitempty"`
	// example:
	//
	// hgpost-cn-xxx
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 30
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// true
	ScheduleTimeDesc *bool `json:"scheduleTimeDesc,omitempty" xml:"scheduleTimeDesc,omitempty"`
	// example:
	//
	// success
	State *string `json:"state,omitempty" xml:"state,omitempty"`
}

func (s ListOperationEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOperationEventsRequest) GoString() string {
	return s.String()
}

func (s *ListOperationEventsRequest) GetEventName() *string {
	return s.EventName
}

func (s *ListOperationEventsRequest) GetEventNameDesc() *bool {
	return s.EventNameDesc
}

func (s *ListOperationEventsRequest) GetEventType() *string {
	return s.EventType
}

func (s *ListOperationEventsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListOperationEventsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListOperationEventsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListOperationEventsRequest) GetScheduleTimeDesc() *bool {
	return s.ScheduleTimeDesc
}

func (s *ListOperationEventsRequest) GetState() *string {
	return s.State
}

func (s *ListOperationEventsRequest) SetEventName(v string) *ListOperationEventsRequest {
	s.EventName = &v
	return s
}

func (s *ListOperationEventsRequest) SetEventNameDesc(v bool) *ListOperationEventsRequest {
	s.EventNameDesc = &v
	return s
}

func (s *ListOperationEventsRequest) SetEventType(v string) *ListOperationEventsRequest {
	s.EventType = &v
	return s
}

func (s *ListOperationEventsRequest) SetInstanceId(v string) *ListOperationEventsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListOperationEventsRequest) SetPageNumber(v int64) *ListOperationEventsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListOperationEventsRequest) SetPageSize(v int64) *ListOperationEventsRequest {
	s.PageSize = &v
	return s
}

func (s *ListOperationEventsRequest) SetScheduleTimeDesc(v bool) *ListOperationEventsRequest {
	s.ScheduleTimeDesc = &v
	return s
}

func (s *ListOperationEventsRequest) SetState(v string) *ListOperationEventsRequest {
	s.State = &v
	return s
}

func (s *ListOperationEventsRequest) Validate() error {
	return dara.Validate(s)
}
