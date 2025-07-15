// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentStatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentIds(v string) *ListAgentStatesRequest
	GetAgentIds() *string
	SetExcludeOfflineUsers(v bool) *ListAgentStatesRequest
	GetExcludeOfflineUsers() *bool
	SetInstanceId(v string) *ListAgentStatesRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *ListAgentStatesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAgentStatesRequest
	GetPageSize() *int32
	SetSkillGroupId(v string) *ListAgentStatesRequest
	GetSkillGroupId() *string
	SetState(v string) *ListAgentStatesRequest
	GetState() *string
}

type ListAgentStatesRequest struct {
	// example:
	//
	// agent@ccc-test
	AgentIds *string `json:"AgentIds,omitempty" xml:"AgentIds,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// true
	ExcludeOfflineUsers *bool `json:"ExcludeOfflineUsers,omitempty" xml:"ExcludeOfflineUsers,omitempty"`
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
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// skillgroup@ccc-test
	SkillGroupId *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	// example:
	//
	// Ready
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ListAgentStatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAgentStatesRequest) GoString() string {
	return s.String()
}

func (s *ListAgentStatesRequest) GetAgentIds() *string {
	return s.AgentIds
}

func (s *ListAgentStatesRequest) GetExcludeOfflineUsers() *bool {
	return s.ExcludeOfflineUsers
}

func (s *ListAgentStatesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListAgentStatesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAgentStatesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAgentStatesRequest) GetSkillGroupId() *string {
	return s.SkillGroupId
}

func (s *ListAgentStatesRequest) GetState() *string {
	return s.State
}

func (s *ListAgentStatesRequest) SetAgentIds(v string) *ListAgentStatesRequest {
	s.AgentIds = &v
	return s
}

func (s *ListAgentStatesRequest) SetExcludeOfflineUsers(v bool) *ListAgentStatesRequest {
	s.ExcludeOfflineUsers = &v
	return s
}

func (s *ListAgentStatesRequest) SetInstanceId(v string) *ListAgentStatesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListAgentStatesRequest) SetPageNumber(v int32) *ListAgentStatesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAgentStatesRequest) SetPageSize(v int32) *ListAgentStatesRequest {
	s.PageSize = &v
	return s
}

func (s *ListAgentStatesRequest) SetSkillGroupId(v string) *ListAgentStatesRequest {
	s.SkillGroupId = &v
	return s
}

func (s *ListAgentStatesRequest) SetState(v string) *ListAgentStatesRequest {
	s.State = &v
	return s
}

func (s *ListAgentStatesRequest) Validate() error {
	return dara.Validate(s)
}
