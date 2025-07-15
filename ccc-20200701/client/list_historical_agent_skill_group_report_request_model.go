// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHistoricalAgentSkillGroupReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentIdList(v string) *ListHistoricalAgentSkillGroupReportRequest
	GetAgentIdList() *string
	SetEndTime(v int64) *ListHistoricalAgentSkillGroupReportRequest
	GetEndTime() *int64
	SetInstanceId(v string) *ListHistoricalAgentSkillGroupReportRequest
	GetInstanceId() *string
	SetMediaType(v string) *ListHistoricalAgentSkillGroupReportRequest
	GetMediaType() *string
	SetPageNumber(v int32) *ListHistoricalAgentSkillGroupReportRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListHistoricalAgentSkillGroupReportRequest
	GetPageSize() *int32
	SetSkillGroupIdList(v string) *ListHistoricalAgentSkillGroupReportRequest
	GetSkillGroupIdList() *string
	SetStartTime(v int64) *ListHistoricalAgentSkillGroupReportRequest
	GetStartTime() *int64
}

type ListHistoricalAgentSkillGroupReportRequest struct {
	// example:
	//
	// ["agent1@ccc-test", "agent2@ccc-test"]
	AgentIdList *string `json:"AgentIdList,omitempty" xml:"AgentIdList,omitempty"`
	// example:
	//
	// 1620273600000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
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
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// [
	//
	//       "skg1@ccc-test",
	//
	//       "skg2@ccc-test"
	//
	// ]
	SkillGroupIdList *string `json:"SkillGroupIdList,omitempty" xml:"SkillGroupIdList,omitempty"`
	// example:
	//
	// 1634140800000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListHistoricalAgentSkillGroupReportRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHistoricalAgentSkillGroupReportRequest) GoString() string {
	return s.String()
}

func (s *ListHistoricalAgentSkillGroupReportRequest) GetAgentIdList() *string {
	return s.AgentIdList
}

func (s *ListHistoricalAgentSkillGroupReportRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListHistoricalAgentSkillGroupReportRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListHistoricalAgentSkillGroupReportRequest) GetMediaType() *string {
	return s.MediaType
}

func (s *ListHistoricalAgentSkillGroupReportRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListHistoricalAgentSkillGroupReportRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListHistoricalAgentSkillGroupReportRequest) GetSkillGroupIdList() *string {
	return s.SkillGroupIdList
}

func (s *ListHistoricalAgentSkillGroupReportRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListHistoricalAgentSkillGroupReportRequest) SetAgentIdList(v string) *ListHistoricalAgentSkillGroupReportRequest {
	s.AgentIdList = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportRequest) SetEndTime(v int64) *ListHistoricalAgentSkillGroupReportRequest {
	s.EndTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportRequest) SetInstanceId(v string) *ListHistoricalAgentSkillGroupReportRequest {
	s.InstanceId = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportRequest) SetMediaType(v string) *ListHistoricalAgentSkillGroupReportRequest {
	s.MediaType = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportRequest) SetPageNumber(v int32) *ListHistoricalAgentSkillGroupReportRequest {
	s.PageNumber = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportRequest) SetPageSize(v int32) *ListHistoricalAgentSkillGroupReportRequest {
	s.PageSize = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportRequest) SetSkillGroupIdList(v string) *ListHistoricalAgentSkillGroupReportRequest {
	s.SkillGroupIdList = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportRequest) SetStartTime(v int64) *ListHistoricalAgentSkillGroupReportRequest {
	s.StartTime = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportRequest) Validate() error {
	return dara.Validate(s)
}
