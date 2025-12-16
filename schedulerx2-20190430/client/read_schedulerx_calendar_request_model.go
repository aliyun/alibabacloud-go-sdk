// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadSchedulerxCalendarRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCalendarName(v string) *ReadSchedulerxCalendarRequest
	GetCalendarName() *string
	SetFetchCalendarDetail(v bool) *ReadSchedulerxCalendarRequest
	GetFetchCalendarDetail() *bool
	SetFetchSystemCalendar(v bool) *ReadSchedulerxCalendarRequest
	GetFetchSystemCalendar() *bool
	SetMaxResults(v int32) *ReadSchedulerxCalendarRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ReadSchedulerxCalendarRequest
	GetNextToken() *string
	SetRegionId(v string) *ReadSchedulerxCalendarRequest
	GetRegionId() *string
	SetYear(v int32) *ReadSchedulerxCalendarRequest
	GetYear() *int32
}

type ReadSchedulerxCalendarRequest struct {
	// The calendar name.
	//
	// example:
	//
	// 2025workday
	CalendarName *string `json:"CalendarName,omitempty" xml:"CalendarName,omitempty"`
	// Specifies whether to retrieve calendar details. The default value is false.
	//
	// 	- false: Returns only basic information without the detailed list of days for each month.
	//
	// 	- true: Returns the detailed list of days for each month.
	//
	// example:
	//
	// false
	FetchCalendarDetail *bool `json:"FetchCalendarDetail,omitempty" xml:"FetchCalendarDetail,omitempty"`
	// Specifies whether to retrieve system calendars. The default value is false.
	//
	// 	- false: Returns only user-defined calendars.
	//
	// 	- true: Returns both system calendars and user-defined calendars.
	//
	// example:
	//
	// false
	FetchSystemCalendar *bool `json:"FetchSystemCalendar,omitempty" xml:"FetchSystemCalendar,omitempty"`
	// The maximum number of entries to return on this call. The default value is 20.
	//
	// example:
	//
	// 15
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A token that specifies the starting position for retrieving the next page of data. You do not need to provide this parameter for the first request. An empty value indicates that all data has been read.
	//
	// example:
	//
	// AAAAAdYzT97YjSXWT8TQmxIAI5g=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The year.
	//
	// example:
	//
	// 2025
	Year *int32 `json:"Year,omitempty" xml:"Year,omitempty"`
}

func (s ReadSchedulerxCalendarRequest) String() string {
	return dara.Prettify(s)
}

func (s ReadSchedulerxCalendarRequest) GoString() string {
	return s.String()
}

func (s *ReadSchedulerxCalendarRequest) GetCalendarName() *string {
	return s.CalendarName
}

func (s *ReadSchedulerxCalendarRequest) GetFetchCalendarDetail() *bool {
	return s.FetchCalendarDetail
}

func (s *ReadSchedulerxCalendarRequest) GetFetchSystemCalendar() *bool {
	return s.FetchSystemCalendar
}

func (s *ReadSchedulerxCalendarRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ReadSchedulerxCalendarRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ReadSchedulerxCalendarRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ReadSchedulerxCalendarRequest) GetYear() *int32 {
	return s.Year
}

func (s *ReadSchedulerxCalendarRequest) SetCalendarName(v string) *ReadSchedulerxCalendarRequest {
	s.CalendarName = &v
	return s
}

func (s *ReadSchedulerxCalendarRequest) SetFetchCalendarDetail(v bool) *ReadSchedulerxCalendarRequest {
	s.FetchCalendarDetail = &v
	return s
}

func (s *ReadSchedulerxCalendarRequest) SetFetchSystemCalendar(v bool) *ReadSchedulerxCalendarRequest {
	s.FetchSystemCalendar = &v
	return s
}

func (s *ReadSchedulerxCalendarRequest) SetMaxResults(v int32) *ReadSchedulerxCalendarRequest {
	s.MaxResults = &v
	return s
}

func (s *ReadSchedulerxCalendarRequest) SetNextToken(v string) *ReadSchedulerxCalendarRequest {
	s.NextToken = &v
	return s
}

func (s *ReadSchedulerxCalendarRequest) SetRegionId(v string) *ReadSchedulerxCalendarRequest {
	s.RegionId = &v
	return s
}

func (s *ReadSchedulerxCalendarRequest) SetYear(v int32) *ReadSchedulerxCalendarRequest {
	s.Year = &v
	return s
}

func (s *ReadSchedulerxCalendarRequest) Validate() error {
	return dara.Validate(s)
}
