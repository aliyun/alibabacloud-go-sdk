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
	// This parameter is required.
	//
	// example:
	//
	// agent@ccc-test
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// example:
	//
	// 1558443508000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// Daily
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// skg-default@ccc-test
	SkillGroupId *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
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
