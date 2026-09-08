// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIntervalInstanceReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *ListIntervalInstanceReportRequest
	GetEndTime() *int64
	SetInstanceId(v string) *ListIntervalInstanceReportRequest
	GetInstanceId() *string
	SetInterval(v string) *ListIntervalInstanceReportRequest
	GetInterval() *string
	SetStartTime(v int64) *ListIntervalInstanceReportRequest
	GetStartTime() *int64
}

type ListIntervalInstanceReportRequest struct {
	// End Time, formatted as a UNIX timestamp in milliseconds. This parameter is optional. The default value is the current time. If Interval is Daily, the maximum interval between StartTime and EndTime is 180 days. If Interval is Hourly, the maximum interval is 10 days. The time precision for statistics is hourly, snapped backward to the start of the hour, using an open interval. For example, if the original Start Time is 11:12:20 and End Time is 11:45:50, the aligned time range becomes [11:00:00, 12:00:00), meaning greater than or equal to 11:00:00 and less than 12:00:00.
	//
	// example:
	//
	// 1620316799000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Segment statistics type. This parameter is optional. The default value is Daily (daily aggregation).
	//
	// example:
	//
	// Hourly
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// Start Time, in UNIX timestamp format with millisecond precision. This parameter is optional. The default value is 00:00 of the current day. The time granularity for statistics is hourly, rounded down to the nearest hour, and uses a closed interval.
	//
	// example:
	//
	// 1620230400000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListIntervalInstanceReportRequest) String() string {
	return dara.Prettify(s)
}

func (s ListIntervalInstanceReportRequest) GoString() string {
	return s.String()
}

func (s *ListIntervalInstanceReportRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListIntervalInstanceReportRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListIntervalInstanceReportRequest) GetInterval() *string {
	return s.Interval
}

func (s *ListIntervalInstanceReportRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListIntervalInstanceReportRequest) SetEndTime(v int64) *ListIntervalInstanceReportRequest {
	s.EndTime = &v
	return s
}

func (s *ListIntervalInstanceReportRequest) SetInstanceId(v string) *ListIntervalInstanceReportRequest {
	s.InstanceId = &v
	return s
}

func (s *ListIntervalInstanceReportRequest) SetInterval(v string) *ListIntervalInstanceReportRequest {
	s.Interval = &v
	return s
}

func (s *ListIntervalInstanceReportRequest) SetStartTime(v int64) *ListIntervalInstanceReportRequest {
	s.StartTime = &v
	return s
}

func (s *ListIntervalInstanceReportRequest) Validate() error {
	return dara.Validate(s)
}
