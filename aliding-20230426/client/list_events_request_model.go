// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCalendarId(v string) *ListEventsRequest
	GetCalendarId() *string
	SetMaxAttendees(v int32) *ListEventsRequest
	GetMaxAttendees() *int32
	SetMaxResults(v int32) *ListEventsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListEventsRequest
	GetNextToken() *string
	SetSeriesMasterId(v string) *ListEventsRequest
	GetSeriesMasterId() *string
	SetShowDeleted(v bool) *ListEventsRequest
	GetShowDeleted() *bool
	SetSyncToken(v string) *ListEventsRequest
	GetSyncToken() *string
	SetTimeMax(v string) *ListEventsRequest
	GetTimeMax() *string
	SetTimeMin(v string) *ListEventsRequest
	GetTimeMin() *string
}

type ListEventsRequest struct {
	// This parameter is required.
	CalendarId     *string `json:"CalendarId,omitempty" xml:"CalendarId,omitempty"`
	MaxAttendees   *int32  `json:"MaxAttendees,omitempty" xml:"MaxAttendees,omitempty"`
	MaxResults     *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken      *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	SeriesMasterId *string `json:"SeriesMasterId,omitempty" xml:"SeriesMasterId,omitempty"`
	ShowDeleted    *bool   `json:"ShowDeleted,omitempty" xml:"ShowDeleted,omitempty"`
	SyncToken      *string `json:"SyncToken,omitempty" xml:"SyncToken,omitempty"`
	TimeMax        *string `json:"TimeMax,omitempty" xml:"TimeMax,omitempty"`
	TimeMin        *string `json:"TimeMin,omitempty" xml:"TimeMin,omitempty"`
}

func (s ListEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEventsRequest) GoString() string {
	return s.String()
}

func (s *ListEventsRequest) GetCalendarId() *string {
	return s.CalendarId
}

func (s *ListEventsRequest) GetMaxAttendees() *int32 {
	return s.MaxAttendees
}

func (s *ListEventsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListEventsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListEventsRequest) GetSeriesMasterId() *string {
	return s.SeriesMasterId
}

func (s *ListEventsRequest) GetShowDeleted() *bool {
	return s.ShowDeleted
}

func (s *ListEventsRequest) GetSyncToken() *string {
	return s.SyncToken
}

func (s *ListEventsRequest) GetTimeMax() *string {
	return s.TimeMax
}

func (s *ListEventsRequest) GetTimeMin() *string {
	return s.TimeMin
}

func (s *ListEventsRequest) SetCalendarId(v string) *ListEventsRequest {
	s.CalendarId = &v
	return s
}

func (s *ListEventsRequest) SetMaxAttendees(v int32) *ListEventsRequest {
	s.MaxAttendees = &v
	return s
}

func (s *ListEventsRequest) SetMaxResults(v int32) *ListEventsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListEventsRequest) SetNextToken(v string) *ListEventsRequest {
	s.NextToken = &v
	return s
}

func (s *ListEventsRequest) SetSeriesMasterId(v string) *ListEventsRequest {
	s.SeriesMasterId = &v
	return s
}

func (s *ListEventsRequest) SetShowDeleted(v bool) *ListEventsRequest {
	s.ShowDeleted = &v
	return s
}

func (s *ListEventsRequest) SetSyncToken(v string) *ListEventsRequest {
	s.SyncToken = &v
	return s
}

func (s *ListEventsRequest) SetTimeMax(v string) *ListEventsRequest {
	s.TimeMax = &v
	return s
}

func (s *ListEventsRequest) SetTimeMin(v string) *ListEventsRequest {
	s.TimeMin = &v
	return s
}

func (s *ListEventsRequest) Validate() error {
	return dara.Validate(s)
}
