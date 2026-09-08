// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIntervalAgentReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *ListIntervalAgentReportRequest
	GetAgentId() *string
	SetEndTime(v int64) *ListIntervalAgentReportRequest
	GetEndTime() *int64
	SetInstanceId(v string) *ListIntervalAgentReportRequest
	GetInstanceId() *string
	SetInterval(v string) *ListIntervalAgentReportRequest
	GetInterval() *string
	SetMediaType(v string) *ListIntervalAgentReportRequest
	GetMediaType() *string
	SetStartTime(v int64) *ListIntervalAgentReportRequest
	GetStartTime() *int64
}

type ListIntervalAgentReportRequest struct {
	// Agent ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// agent@ccc-test
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// End time, formatted as a UNIX timestamp in milliseconds. This parameter is optional. The default value is the current time. If Interval is Daily, the maximum interval between StartTime and EndTime is 180 days. If Interval is Hourly, the maximum interval is 10 days. Time precision for statistics is at the hour level, rounded down to the next full hour, using an open interval. For example, if the start time is 11:12:20 and the end time is 11:45:50, the aligned input time range becomes [11:00:00, 12:00:00), meaning greater than or equal to 11:00:00 and less than 12:00:00.
	//
	// example:
	//
	// 1532707199000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Segment statistics type. Optional. Default value is Daily (aggregated by Day).
	//
	// example:
	//
	// Hourly
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// Media type. The default value is Audio. Other valid values include Chat and Video.
	//
	// example:
	//
	// VIDEO
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// Start time, formatted as a UNIX timestamp in milliseconds. This parameter is optional. The default value is 00:00:00 of the current day. Time precision for statistics is at the hour level, rounded down to the previous full hour, using a closed interval.
	//
	// example:
	//
	// 1532448000000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListIntervalAgentReportRequest) String() string {
	return dara.Prettify(s)
}

func (s ListIntervalAgentReportRequest) GoString() string {
	return s.String()
}

func (s *ListIntervalAgentReportRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *ListIntervalAgentReportRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListIntervalAgentReportRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListIntervalAgentReportRequest) GetInterval() *string {
	return s.Interval
}

func (s *ListIntervalAgentReportRequest) GetMediaType() *string {
	return s.MediaType
}

func (s *ListIntervalAgentReportRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListIntervalAgentReportRequest) SetAgentId(v string) *ListIntervalAgentReportRequest {
	s.AgentId = &v
	return s
}

func (s *ListIntervalAgentReportRequest) SetEndTime(v int64) *ListIntervalAgentReportRequest {
	s.EndTime = &v
	return s
}

func (s *ListIntervalAgentReportRequest) SetInstanceId(v string) *ListIntervalAgentReportRequest {
	s.InstanceId = &v
	return s
}

func (s *ListIntervalAgentReportRequest) SetInterval(v string) *ListIntervalAgentReportRequest {
	s.Interval = &v
	return s
}

func (s *ListIntervalAgentReportRequest) SetMediaType(v string) *ListIntervalAgentReportRequest {
	s.MediaType = &v
	return s
}

func (s *ListIntervalAgentReportRequest) SetStartTime(v int64) *ListIntervalAgentReportRequest {
	s.StartTime = &v
	return s
}

func (s *ListIntervalAgentReportRequest) Validate() error {
	return dara.Validate(s)
}
