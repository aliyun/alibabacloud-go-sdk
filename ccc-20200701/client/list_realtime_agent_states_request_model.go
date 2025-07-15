// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRealtimeAgentStatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentIdList(v string) *ListRealtimeAgentStatesRequest
	GetAgentIdList() *string
	SetAgentName(v string) *ListRealtimeAgentStatesRequest
	GetAgentName() *string
	SetCallTypeList(v string) *ListRealtimeAgentStatesRequest
	GetCallTypeList() *string
	SetInstanceId(v string) *ListRealtimeAgentStatesRequest
	GetInstanceId() *string
	SetMediaType(v string) *ListRealtimeAgentStatesRequest
	GetMediaType() *string
	SetOutboundScenario(v bool) *ListRealtimeAgentStatesRequest
	GetOutboundScenario() *bool
	SetPageNumber(v int32) *ListRealtimeAgentStatesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListRealtimeAgentStatesRequest
	GetPageSize() *int32
	SetQuery(v string) *ListRealtimeAgentStatesRequest
	GetQuery() *string
	SetSkillGroupId(v string) *ListRealtimeAgentStatesRequest
	GetSkillGroupId() *string
	SetStateList(v string) *ListRealtimeAgentStatesRequest
	GetStateList() *string
	SetWorkModeList(v string) *ListRealtimeAgentStatesRequest
	GetWorkModeList() *string
}

type ListRealtimeAgentStatesRequest struct {
	// example:
	//
	// ["agent1@ccc-test", "agent2@ccc-test"]
	AgentIdList *string `json:"AgentIdList,omitempty" xml:"AgentIdList,omitempty"`
	// example:
	//
	// agent
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// example:
	//
	// ["INBOUND", "OUTBOUND"]
	CallTypeList *string `json:"CallTypeList,omitempty" xml:"CallTypeList,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MediaType  *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// example:
	//
	// false
	OutboundScenario *bool `json:"OutboundScenario,omitempty" xml:"OutboundScenario,omitempty"`
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
	// agent
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// example:
	//
	// skillgroup@ccc-test
	SkillGroupId *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	// example:
	//
	// ["ACW", "Dialing"]
	StateList *string `json:"StateList,omitempty" xml:"StateList,omitempty"`
	// example:
	//
	// ["OFFICE_PHONE","ON_SITE"]
	WorkModeList *string `json:"WorkModeList,omitempty" xml:"WorkModeList,omitempty"`
}

func (s ListRealtimeAgentStatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRealtimeAgentStatesRequest) GoString() string {
	return s.String()
}

func (s *ListRealtimeAgentStatesRequest) GetAgentIdList() *string {
	return s.AgentIdList
}

func (s *ListRealtimeAgentStatesRequest) GetAgentName() *string {
	return s.AgentName
}

func (s *ListRealtimeAgentStatesRequest) GetCallTypeList() *string {
	return s.CallTypeList
}

func (s *ListRealtimeAgentStatesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListRealtimeAgentStatesRequest) GetMediaType() *string {
	return s.MediaType
}

func (s *ListRealtimeAgentStatesRequest) GetOutboundScenario() *bool {
	return s.OutboundScenario
}

func (s *ListRealtimeAgentStatesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListRealtimeAgentStatesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRealtimeAgentStatesRequest) GetQuery() *string {
	return s.Query
}

func (s *ListRealtimeAgentStatesRequest) GetSkillGroupId() *string {
	return s.SkillGroupId
}

func (s *ListRealtimeAgentStatesRequest) GetStateList() *string {
	return s.StateList
}

func (s *ListRealtimeAgentStatesRequest) GetWorkModeList() *string {
	return s.WorkModeList
}

func (s *ListRealtimeAgentStatesRequest) SetAgentIdList(v string) *ListRealtimeAgentStatesRequest {
	s.AgentIdList = &v
	return s
}

func (s *ListRealtimeAgentStatesRequest) SetAgentName(v string) *ListRealtimeAgentStatesRequest {
	s.AgentName = &v
	return s
}

func (s *ListRealtimeAgentStatesRequest) SetCallTypeList(v string) *ListRealtimeAgentStatesRequest {
	s.CallTypeList = &v
	return s
}

func (s *ListRealtimeAgentStatesRequest) SetInstanceId(v string) *ListRealtimeAgentStatesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListRealtimeAgentStatesRequest) SetMediaType(v string) *ListRealtimeAgentStatesRequest {
	s.MediaType = &v
	return s
}

func (s *ListRealtimeAgentStatesRequest) SetOutboundScenario(v bool) *ListRealtimeAgentStatesRequest {
	s.OutboundScenario = &v
	return s
}

func (s *ListRealtimeAgentStatesRequest) SetPageNumber(v int32) *ListRealtimeAgentStatesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRealtimeAgentStatesRequest) SetPageSize(v int32) *ListRealtimeAgentStatesRequest {
	s.PageSize = &v
	return s
}

func (s *ListRealtimeAgentStatesRequest) SetQuery(v string) *ListRealtimeAgentStatesRequest {
	s.Query = &v
	return s
}

func (s *ListRealtimeAgentStatesRequest) SetSkillGroupId(v string) *ListRealtimeAgentStatesRequest {
	s.SkillGroupId = &v
	return s
}

func (s *ListRealtimeAgentStatesRequest) SetStateList(v string) *ListRealtimeAgentStatesRequest {
	s.StateList = &v
	return s
}

func (s *ListRealtimeAgentStatesRequest) SetWorkModeList(v string) *ListRealtimeAgentStatesRequest {
	s.WorkModeList = &v
	return s
}

func (s *ListRealtimeAgentStatesRequest) Validate() error {
	return dara.Validate(s)
}
