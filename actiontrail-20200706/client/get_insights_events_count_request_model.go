// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInsightsEventsCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDate(v string) *GetInsightsEventsCountRequest
	GetDate() *string
	SetEndTime(v string) *GetInsightsEventsCountRequest
	GetEndTime() *string
	SetStartTime(v string) *GetInsightsEventsCountRequest
	GetStartTime() *string
}

type GetInsightsEventsCountRequest struct {
	// The date to query. The format is `yyyy-MM-dd`.
	//
	// example:
	//
	// 2026-01-07
	Date *string `json:"Date,omitempty" xml:"Date,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time must be in UTC.
	//
	// > - - If Date, StartTime, and EndTime are all left empty, the system queries the number of events in the last 24 hours.
	//
	// >
	//
	// >   - If Date is specified, the StartTime and EndTime parameters are ignored. The system queries the number of events on the specified date.
	//
	// >
	//
	// >   - If Date is left empty and both StartTime and EndTime are specified, the system queries the number of events in the specified time range.
	//
	// example:
	//
	// 2026-01-07T06:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time must be in UTC.
	//
	// example:
	//
	// 2025-12-01T02:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetInsightsEventsCountRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInsightsEventsCountRequest) GoString() string {
	return s.String()
}

func (s *GetInsightsEventsCountRequest) GetDate() *string {
	return s.Date
}

func (s *GetInsightsEventsCountRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *GetInsightsEventsCountRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *GetInsightsEventsCountRequest) SetDate(v string) *GetInsightsEventsCountRequest {
	s.Date = &v
	return s
}

func (s *GetInsightsEventsCountRequest) SetEndTime(v string) *GetInsightsEventsCountRequest {
	s.EndTime = &v
	return s
}

func (s *GetInsightsEventsCountRequest) SetStartTime(v string) *GetInsightsEventsCountRequest {
	s.StartTime = &v
	return s
}

func (s *GetInsightsEventsCountRequest) Validate() error {
	return dara.Validate(s)
}
