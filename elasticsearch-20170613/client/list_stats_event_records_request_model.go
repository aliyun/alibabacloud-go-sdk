// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStatsEventRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventType(v string) *ListStatsEventRecordsRequest
	GetEventType() *string
	SetLevel(v string) *ListStatsEventRecordsRequest
	GetLevel() *string
	SetStatus(v string) *ListStatsEventRecordsRequest
	GetStatus() *string
}

type ListStatsEventRecordsRequest struct {
	// example:
	//
	// UserOperator
	EventType *string `json:"eventType,omitempty" xml:"eventType,omitempty"`
	// example:
	//
	// Info
	Level *string `json:"level,omitempty" xml:"level,omitempty"`
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

func (s *ListStatsEventRecordsRequest) GetEventType() *string {
	return s.EventType
}

func (s *ListStatsEventRecordsRequest) GetLevel() *string {
	return s.Level
}

func (s *ListStatsEventRecordsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListStatsEventRecordsRequest) SetEventType(v string) *ListStatsEventRecordsRequest {
	s.EventType = &v
	return s
}

func (s *ListStatsEventRecordsRequest) SetLevel(v string) *ListStatsEventRecordsRequest {
	s.Level = &v
	return s
}

func (s *ListStatsEventRecordsRequest) SetStatus(v string) *ListStatsEventRecordsRequest {
	s.Status = &v
	return s
}

func (s *ListStatsEventRecordsRequest) Validate() error {
	return dara.Validate(s)
}
