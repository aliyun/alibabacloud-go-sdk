// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEventsViewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCalendarId(v string) *ListEventsViewRequest
	GetCalendarId() *string
	SetMaxAttendees(v int32) *ListEventsViewRequest
	GetMaxAttendees() *int32
	SetMaxResults(v int32) *ListEventsViewRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListEventsViewRequest
	GetNextToken() *string
	SetTimeMax(v string) *ListEventsViewRequest
	GetTimeMax() *string
	SetTimeMin(v string) *ListEventsViewRequest
	GetTimeMin() *string
}

type ListEventsViewRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// primary
	CalendarId *string `json:"CalendarId,omitempty" xml:"CalendarId,omitempty"`
	// example:
	//
	// 100
	MaxAttendees *int32 `json:"MaxAttendees,omitempty" xml:"MaxAttendees,omitempty"`
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// cnNTbW1xxx
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 2022-11-28T00:00:00+08:00
	TimeMax *string `json:"TimeMax,omitempty" xml:"TimeMax,omitempty"`
	// example:
	//
	// 2022-11-27T00:00:00+08:00
	TimeMin *string `json:"TimeMin,omitempty" xml:"TimeMin,omitempty"`
}

func (s ListEventsViewRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEventsViewRequest) GoString() string {
	return s.String()
}

func (s *ListEventsViewRequest) GetCalendarId() *string {
	return s.CalendarId
}

func (s *ListEventsViewRequest) GetMaxAttendees() *int32 {
	return s.MaxAttendees
}

func (s *ListEventsViewRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListEventsViewRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListEventsViewRequest) GetTimeMax() *string {
	return s.TimeMax
}

func (s *ListEventsViewRequest) GetTimeMin() *string {
	return s.TimeMin
}

func (s *ListEventsViewRequest) SetCalendarId(v string) *ListEventsViewRequest {
	s.CalendarId = &v
	return s
}

func (s *ListEventsViewRequest) SetMaxAttendees(v int32) *ListEventsViewRequest {
	s.MaxAttendees = &v
	return s
}

func (s *ListEventsViewRequest) SetMaxResults(v int32) *ListEventsViewRequest {
	s.MaxResults = &v
	return s
}

func (s *ListEventsViewRequest) SetNextToken(v string) *ListEventsViewRequest {
	s.NextToken = &v
	return s
}

func (s *ListEventsViewRequest) SetTimeMax(v string) *ListEventsViewRequest {
	s.TimeMax = &v
	return s
}

func (s *ListEventsViewRequest) SetTimeMin(v string) *ListEventsViewRequest {
	s.TimeMin = &v
	return s
}

func (s *ListEventsViewRequest) Validate() error {
	return dara.Validate(s)
}
