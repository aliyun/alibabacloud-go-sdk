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
	// List of agent IDs, formatted as a JSON array string. The array can contain up to 20 elements. This parameter is optional and defaults to empty, which matches all agents under the current instance.
	//
	// example:
	//
	// ["agent1@ccc-test", "agent2@ccc-test"]
	AgentIdList *string `json:"AgentIdList,omitempty" xml:"AgentIdList,omitempty"`
	// Perform fuzzy matching by agent name.
	//
	// example:
	//
	// agent
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// Filters by a list of call types. The value is a string in JSON array format, where each array element is a call type. This parameter is optional and defaults to empty, which matches all call types.
	//
	// example:
	//
	// ["Inbound", "Outbound"]
	CallTypeList *string `json:"CallTypeList,omitempty" xml:"CallTypeList,omitempty"`
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Media type. The default is Audio. Other options include Chat (text), Video, and ALL.
	//
	// example:
	//
	// AUDIO
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// Filters agents who are in outbound-only mode. This parameter is optional and defaults to empty, which means no filtering by outbound-only mode is applied.
	//
	// example:
	//
	// false
	OutboundScenario *bool `json:"OutboundScenario,omitempty" xml:"OutboundScenario,omitempty"`
	// Page number, ranging from 1 to 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Page size, ranging from 1 to 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Performs fuzzy filtering based on the full or partial agent display name, agent ID, or agent extension number. This parameter is optional and defaults to empty, which means no filtering is applied.
	//
	// example:
	//
	// agent
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// Filter by skill group ID. This parameter is optional and defaults to empty, which means no filtering is applied.
	//
	// example:
	//
	// skillgroup@ccc-test
	SkillGroupId *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	// Filters by a list of statuses. This parameter is optional and defaults to empty, which matches all statuses.
	//
	// example:
	//
	// ["ACW", "Dialing"]
	StateList *string `json:"StateList,omitempty" xml:"StateList,omitempty"`
	// Filter by work mode list. This parameter is optional and defaults to empty, which means all work modes are matched.
	//
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
