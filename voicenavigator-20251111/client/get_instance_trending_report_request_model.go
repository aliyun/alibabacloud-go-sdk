// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceTrendingReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *GetInstanceTrendingReportRequest
	GetEndTime() *int64
	SetInstanceId(v string) *GetInstanceTrendingReportRequest
	GetInstanceId() *string
	SetStartTime(v int64) *GetInstanceTrendingReportRequest
	GetStartTime() *int64
	SetTimeInterval(v string) *GetInstanceTrendingReportRequest
	GetTimeInterval() *string
}

type GetInstanceTrendingReportRequest struct {
	// The end time of the statistical period. The value is a millisecond-level UNIX timestamp that represents the number of milliseconds elapsed since 1970-01-01 00:00:00 UTC. The following constraints apply:
	//
	// Time range constraint: The maximum time range is 7 days for hourly queries and 180 days for daily queries.
	//
	// Time granularity alignment: For hourly queries, the minutes, seconds, and milliseconds must be 0. For daily queries, the hours, minutes, seconds, and milliseconds must be 0.
	//
	// If the specified timestamp is not aligned with the corresponding granularity, the server automatically rounds it up.
	//
	// example:
	//
	// 1582103299434
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// 82b2eaae-ce5c-45f8-8231-f15b6b27e55c
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The start time of the statistical period. The value is a millisecond-level UNIX timestamp that represents the number of milliseconds elapsed since 1970-01-01 00:00:00 UTC. The following constraints apply:
	//
	// 1. Earliest boundary: The start time can trace back to 00:00:00 of the date 365 calendar days before the current date.
	//
	// 2. Time granularity alignment: For hourly queries, the minutes, seconds, and milliseconds must be 0. For daily queries, the hours, minutes, seconds, and milliseconds must be 0.
	//
	// 3. If the specified timestamp is not aligned with the corresponding granularity, the server automatically rounds it up.
	//
	// example:
	//
	// 1582267398628
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The time interval.
	//
	// example:
	//
	// 1d
	TimeInterval *string `json:"TimeInterval,omitempty" xml:"TimeInterval,omitempty"`
}

func (s GetInstanceTrendingReportRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceTrendingReportRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceTrendingReportRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetInstanceTrendingReportRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceTrendingReportRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetInstanceTrendingReportRequest) GetTimeInterval() *string {
	return s.TimeInterval
}

func (s *GetInstanceTrendingReportRequest) SetEndTime(v int64) *GetInstanceTrendingReportRequest {
	s.EndTime = &v
	return s
}

func (s *GetInstanceTrendingReportRequest) SetInstanceId(v string) *GetInstanceTrendingReportRequest {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceTrendingReportRequest) SetStartTime(v int64) *GetInstanceTrendingReportRequest {
	s.StartTime = &v
	return s
}

func (s *GetInstanceTrendingReportRequest) SetTimeInterval(v string) *GetInstanceTrendingReportRequest {
	s.TimeInterval = &v
	return s
}

func (s *GetInstanceTrendingReportRequest) Validate() error {
	return dara.Validate(s)
}
