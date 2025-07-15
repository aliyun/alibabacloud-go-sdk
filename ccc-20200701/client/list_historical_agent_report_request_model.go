// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHistoricalAgentReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentIdList(v string) *ListHistoricalAgentReportRequest
	GetAgentIdList() *string
	SetInstanceId(v string) *ListHistoricalAgentReportRequest
	GetInstanceId() *string
	SetMediaType(v string) *ListHistoricalAgentReportRequest
	GetMediaType() *string
	SetPageNumber(v int32) *ListHistoricalAgentReportRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListHistoricalAgentReportRequest
	GetPageSize() *int32
	SetSkillGroupId(v string) *ListHistoricalAgentReportRequest
	GetSkillGroupId() *string
	SetStartTime(v int64) *ListHistoricalAgentReportRequest
	GetStartTime() *int64
	SetStopTime(v int64) *ListHistoricalAgentReportRequest
	GetStopTime() *int64
}

type ListHistoricalAgentReportRequest struct {
	// example:
	//
	// ["agent1@ccc-test", "agent2@ccc-test"]
	AgentIdList *string `json:"AgentIdList,omitempty" xml:"AgentIdList,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MediaType  *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 100
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SkillGroupId *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	// example:
	//
	// 1532448000000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 1532707199000
	StopTime *int64 `json:"StopTime,omitempty" xml:"StopTime,omitempty"`
}

func (s ListHistoricalAgentReportRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHistoricalAgentReportRequest) GoString() string {
	return s.String()
}

func (s *ListHistoricalAgentReportRequest) GetAgentIdList() *string {
	return s.AgentIdList
}

func (s *ListHistoricalAgentReportRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListHistoricalAgentReportRequest) GetMediaType() *string {
	return s.MediaType
}

func (s *ListHistoricalAgentReportRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListHistoricalAgentReportRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListHistoricalAgentReportRequest) GetSkillGroupId() *string {
	return s.SkillGroupId
}

func (s *ListHistoricalAgentReportRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListHistoricalAgentReportRequest) GetStopTime() *int64 {
	return s.StopTime
}

func (s *ListHistoricalAgentReportRequest) SetAgentIdList(v string) *ListHistoricalAgentReportRequest {
	s.AgentIdList = &v
	return s
}

func (s *ListHistoricalAgentReportRequest) SetInstanceId(v string) *ListHistoricalAgentReportRequest {
	s.InstanceId = &v
	return s
}

func (s *ListHistoricalAgentReportRequest) SetMediaType(v string) *ListHistoricalAgentReportRequest {
	s.MediaType = &v
	return s
}

func (s *ListHistoricalAgentReportRequest) SetPageNumber(v int32) *ListHistoricalAgentReportRequest {
	s.PageNumber = &v
	return s
}

func (s *ListHistoricalAgentReportRequest) SetPageSize(v int32) *ListHistoricalAgentReportRequest {
	s.PageSize = &v
	return s
}

func (s *ListHistoricalAgentReportRequest) SetSkillGroupId(v string) *ListHistoricalAgentReportRequest {
	s.SkillGroupId = &v
	return s
}

func (s *ListHistoricalAgentReportRequest) SetStartTime(v int64) *ListHistoricalAgentReportRequest {
	s.StartTime = &v
	return s
}

func (s *ListHistoricalAgentReportRequest) SetStopTime(v int64) *ListHistoricalAgentReportRequest {
	s.StopTime = &v
	return s
}

func (s *ListHistoricalAgentReportRequest) Validate() error {
	return dara.Validate(s)
}
