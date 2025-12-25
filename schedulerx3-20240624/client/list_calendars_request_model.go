// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCalendarsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCalendarName(v string) *ListCalendarsRequest
	GetCalendarName() *string
	SetClusterId(v string) *ListCalendarsRequest
	GetClusterId() *string
	SetFetchCalendarDetail(v bool) *ListCalendarsRequest
	GetFetchCalendarDetail() *bool
	SetMaxResults(v int32) *ListCalendarsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListCalendarsRequest
	GetNextToken() *string
	SetYear(v int32) *ListCalendarsRequest
	GetYear() *int32
}

type ListCalendarsRequest struct {
	// example:
	//
	// workday
	CalendarName *string `json:"CalendarName,omitempty" xml:"CalendarName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// false
	FetchCalendarDetail *bool `json:"FetchCalendarDetail,omitempty" xml:"FetchCalendarDetail,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// eCKqVlS5FKF5EWGGOo8EgQ==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 2024
	Year *int32 `json:"Year,omitempty" xml:"Year,omitempty"`
}

func (s ListCalendarsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCalendarsRequest) GoString() string {
	return s.String()
}

func (s *ListCalendarsRequest) GetCalendarName() *string {
	return s.CalendarName
}

func (s *ListCalendarsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListCalendarsRequest) GetFetchCalendarDetail() *bool {
	return s.FetchCalendarDetail
}

func (s *ListCalendarsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListCalendarsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListCalendarsRequest) GetYear() *int32 {
	return s.Year
}

func (s *ListCalendarsRequest) SetCalendarName(v string) *ListCalendarsRequest {
	s.CalendarName = &v
	return s
}

func (s *ListCalendarsRequest) SetClusterId(v string) *ListCalendarsRequest {
	s.ClusterId = &v
	return s
}

func (s *ListCalendarsRequest) SetFetchCalendarDetail(v bool) *ListCalendarsRequest {
	s.FetchCalendarDetail = &v
	return s
}

func (s *ListCalendarsRequest) SetMaxResults(v int32) *ListCalendarsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListCalendarsRequest) SetNextToken(v string) *ListCalendarsRequest {
	s.NextToken = &v
	return s
}

func (s *ListCalendarsRequest) SetYear(v int32) *ListCalendarsRequest {
	s.Year = &v
	return s
}

func (s *ListCalendarsRequest) Validate() error {
	return dara.Validate(s)
}
