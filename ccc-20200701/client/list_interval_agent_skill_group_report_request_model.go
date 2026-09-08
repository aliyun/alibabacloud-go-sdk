// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIntervalAgentSkillGroupReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *ListIntervalAgentSkillGroupReportRequest
	GetAgentId() *string
	SetEndTime(v int64) *ListIntervalAgentSkillGroupReportRequest
	GetEndTime() *int64
	SetInstanceId(v string) *ListIntervalAgentSkillGroupReportRequest
	GetInstanceId() *string
	SetInterval(v string) *ListIntervalAgentSkillGroupReportRequest
	GetInterval() *string
	SetSkillGroupId(v string) *ListIntervalAgentSkillGroupReportRequest
	GetSkillGroupId() *string
	SetStartTime(v int64) *ListIntervalAgentSkillGroupReportRequest
	GetStartTime() *int64
}

type ListIntervalAgentSkillGroupReportRequest struct {
	// The agent ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// agent@ccc-test
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// The end time. This is a UNIX timestamp in milliseconds. This parameter is not required. The default value is the current time. If Interval is set to Daily, the maximum interval between StartTime and EndTime is 180 days. If Interval is set to Hourly, the maximum interval is 10 days. The statistics are measured in hours and rounded up to the nearest hour. This is an open interval. For example, if the start time is 11:12:20 and the end time is 11:45:50, the aligned time range for the input parameters is [11:00:00, 12:00:00), which means greater than or equal to 11:00 and less than 12:00.
	//
	// example:
	//
	// 1558443508000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of interval-based statistics. This parameter is not required. The default value is Daily (summarized by day).
	//
	// example:
	//
	// Daily
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The skill group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// skg-default@ccc-test
	SkillGroupId *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	// The start timestamp. The default value is 00:00 on the current day. The statistics are measured in hours and rounded down to the nearest hour. This is a closed interval.
	//
	// example:
	//
	// 1532448000000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListIntervalAgentSkillGroupReportRequest) String() string {
	return dara.Prettify(s)
}

func (s ListIntervalAgentSkillGroupReportRequest) GoString() string {
	return s.String()
}

func (s *ListIntervalAgentSkillGroupReportRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *ListIntervalAgentSkillGroupReportRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListIntervalAgentSkillGroupReportRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListIntervalAgentSkillGroupReportRequest) GetInterval() *string {
	return s.Interval
}

func (s *ListIntervalAgentSkillGroupReportRequest) GetSkillGroupId() *string {
	return s.SkillGroupId
}

func (s *ListIntervalAgentSkillGroupReportRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListIntervalAgentSkillGroupReportRequest) SetAgentId(v string) *ListIntervalAgentSkillGroupReportRequest {
	s.AgentId = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportRequest) SetEndTime(v int64) *ListIntervalAgentSkillGroupReportRequest {
	s.EndTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportRequest) SetInstanceId(v string) *ListIntervalAgentSkillGroupReportRequest {
	s.InstanceId = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportRequest) SetInterval(v string) *ListIntervalAgentSkillGroupReportRequest {
	s.Interval = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportRequest) SetSkillGroupId(v string) *ListIntervalAgentSkillGroupReportRequest {
	s.SkillGroupId = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportRequest) SetStartTime(v int64) *ListIntervalAgentSkillGroupReportRequest {
	s.StartTime = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportRequest) Validate() error {
	return dara.Validate(s)
}
