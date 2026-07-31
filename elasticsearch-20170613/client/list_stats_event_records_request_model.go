// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStatsEventRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *ListStatsEventRecordsRequest
	GetEndTime() *int64
	SetEventType(v string) *ListStatsEventRecordsRequest
	GetEventType() *string
	SetLevel(v string) *ListStatsEventRecordsRequest
	GetLevel() *string
	SetStartTime(v int64) *ListStatsEventRecordsRequest
	GetStartTime() *int64
	SetStatus(v string) *ListStatsEventRecordsRequest
	GetStatus() *string
}

type ListStatsEventRecordsRequest struct {
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The event type.
	//
	// example:
	//
	// UserOperator
	EventType *string `json:"eventType,omitempty" xml:"eventType,omitempty"`
	// The event level.
	//
	// example:
	//
	// Info
	Level     *string `json:"level,omitempty" xml:"level,omitempty"`
	StartTime *int64  `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// The event status.
	//
	// example:
	//
	// Executed
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListStatsEventRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListStatsEventRecordsRequest) GoString() string {
	return s.String()
}

func (s *ListStatsEventRecordsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListStatsEventRecordsRequest) GetEventType() *string {
	return s.EventType
}

func (s *ListStatsEventRecordsRequest) GetLevel() *string {
	return s.Level
}

func (s *ListStatsEventRecordsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListStatsEventRecordsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListStatsEventRecordsRequest) SetEndTime(v int64) *ListStatsEventRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *ListStatsEventRecordsRequest) SetEventType(v string) *ListStatsEventRecordsRequest {
	s.EventType = &v
	return s
}

func (s *ListStatsEventRecordsRequest) SetLevel(v string) *ListStatsEventRecordsRequest {
	s.Level = &v
	return s
}

func (s *ListStatsEventRecordsRequest) SetStartTime(v int64) *ListStatsEventRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *ListStatsEventRecordsRequest) SetStatus(v string) *ListStatsEventRecordsRequest {
	s.Status = &v
	return s
}

func (s *ListStatsEventRecordsRequest) Validate() error {
	return dara.Validate(s)
}
