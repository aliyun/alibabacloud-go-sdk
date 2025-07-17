// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDIJobEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDIJobId(v int64) *ListDIJobEventsRequest
	GetDIJobId() *int64
	SetEndTime(v int64) *ListDIJobEventsRequest
	GetEndTime() *int64
	SetEventType(v string) *ListDIJobEventsRequest
	GetEventType() *string
	SetPageNumber(v int64) *ListDIJobEventsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListDIJobEventsRequest
	GetPageSize() *int64
	SetStartTime(v int64) *ListDIJobEventsRequest
	GetStartTime() *int64
}

type ListDIJobEventsRequest struct {
	// The ID of the synchronization task.
	//
	// example:
	//
	// 11588
	DIJobId *int64 `json:"DIJobId,omitempty" xml:"DIJobId,omitempty"`
	// The end of the time range to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1717971005
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The type of event that you want to query. Valid values: Failover, Alarm, and DDL.
	//
	// This parameter is required.
	//
	// example:
	//
	// Alarm
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The beginning of the time range to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1716971005
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListDIJobEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDIJobEventsRequest) GoString() string {
	return s.String()
}

func (s *ListDIJobEventsRequest) GetDIJobId() *int64 {
	return s.DIJobId
}

func (s *ListDIJobEventsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListDIJobEventsRequest) GetEventType() *string {
	return s.EventType
}

func (s *ListDIJobEventsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListDIJobEventsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListDIJobEventsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListDIJobEventsRequest) SetDIJobId(v int64) *ListDIJobEventsRequest {
	s.DIJobId = &v
	return s
}

func (s *ListDIJobEventsRequest) SetEndTime(v int64) *ListDIJobEventsRequest {
	s.EndTime = &v
	return s
}

func (s *ListDIJobEventsRequest) SetEventType(v string) *ListDIJobEventsRequest {
	s.EventType = &v
	return s
}

func (s *ListDIJobEventsRequest) SetPageNumber(v int64) *ListDIJobEventsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDIJobEventsRequest) SetPageSize(v int64) *ListDIJobEventsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDIJobEventsRequest) SetStartTime(v int64) *ListDIJobEventsRequest {
	s.StartTime = &v
	return s
}

func (s *ListDIJobEventsRequest) Validate() error {
	return dara.Validate(s)
}
