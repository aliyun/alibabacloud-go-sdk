// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAsyncTaskEvent interface {
	dara.Model
	String() string
	GoString() string
	SetEventDetail(v string) *AsyncTaskEvent
	GetEventDetail() *string
	SetEventId(v int64) *AsyncTaskEvent
	GetEventId() *int64
	SetStatus(v string) *AsyncTaskEvent
	GetStatus() *string
	SetTimestamp(v int64) *AsyncTaskEvent
	GetTimestamp() *int64
}

type AsyncTaskEvent struct {
	// example:
	//
	// body
	EventDetail *string `json:"eventDetail,omitempty" xml:"eventDetail,omitempty"`
	// example:
	//
	// 1
	EventId *int64 `json:"eventId,omitempty" xml:"eventId,omitempty"`
	// example:
	//
	// Succeeded
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 1647420449721
	Timestamp *int64 `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
}

func (s AsyncTaskEvent) String() string {
	return dara.Prettify(s)
}

func (s AsyncTaskEvent) GoString() string {
	return s.String()
}

func (s *AsyncTaskEvent) GetEventDetail() *string {
	return s.EventDetail
}

func (s *AsyncTaskEvent) GetEventId() *int64 {
	return s.EventId
}

func (s *AsyncTaskEvent) GetStatus() *string {
	return s.Status
}

func (s *AsyncTaskEvent) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *AsyncTaskEvent) SetEventDetail(v string) *AsyncTaskEvent {
	s.EventDetail = &v
	return s
}

func (s *AsyncTaskEvent) SetEventId(v int64) *AsyncTaskEvent {
	s.EventId = &v
	return s
}

func (s *AsyncTaskEvent) SetStatus(v string) *AsyncTaskEvent {
	s.Status = &v
	return s
}

func (s *AsyncTaskEvent) SetTimestamp(v int64) *AsyncTaskEvent {
	s.Timestamp = &v
	return s
}

func (s *AsyncTaskEvent) Validate() error {
	return dara.Validate(s)
}
