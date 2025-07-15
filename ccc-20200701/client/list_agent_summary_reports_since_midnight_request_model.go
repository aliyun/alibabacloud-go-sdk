// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentSummaryReportsSinceMidnightRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentIds(v string) *ListAgentSummaryReportsSinceMidnightRequest
	GetAgentIds() *string
	SetInstanceId(v string) *ListAgentSummaryReportsSinceMidnightRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *ListAgentSummaryReportsSinceMidnightRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAgentSummaryReportsSinceMidnightRequest
	GetPageSize() *int32
	SetSkillGroupId(v string) *ListAgentSummaryReportsSinceMidnightRequest
	GetSkillGroupId() *string
}

type ListAgentSummaryReportsSinceMidnightRequest struct {
	// example:
	//
	// ["agent1@ccc-test", "agent2@ccc-test"]
	AgentIds *string `json:"AgentIds,omitempty" xml:"AgentIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// skillgroup@ccc-test
	SkillGroupId *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
}

func (s ListAgentSummaryReportsSinceMidnightRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAgentSummaryReportsSinceMidnightRequest) GoString() string {
	return s.String()
}

func (s *ListAgentSummaryReportsSinceMidnightRequest) GetAgentIds() *string {
	return s.AgentIds
}

func (s *ListAgentSummaryReportsSinceMidnightRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListAgentSummaryReportsSinceMidnightRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAgentSummaryReportsSinceMidnightRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAgentSummaryReportsSinceMidnightRequest) GetSkillGroupId() *string {
	return s.SkillGroupId
}

func (s *ListAgentSummaryReportsSinceMidnightRequest) SetAgentIds(v string) *ListAgentSummaryReportsSinceMidnightRequest {
	s.AgentIds = &v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightRequest) SetInstanceId(v string) *ListAgentSummaryReportsSinceMidnightRequest {
	s.InstanceId = &v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightRequest) SetPageNumber(v int32) *ListAgentSummaryReportsSinceMidnightRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightRequest) SetPageSize(v int32) *ListAgentSummaryReportsSinceMidnightRequest {
	s.PageSize = &v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightRequest) SetSkillGroupId(v string) *ListAgentSummaryReportsSinceMidnightRequest {
	s.SkillGroupId = &v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightRequest) Validate() error {
	return dara.Validate(s)
}
