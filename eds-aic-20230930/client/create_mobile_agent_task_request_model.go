// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMobileAgentTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *CreateMobileAgentTaskRequest
	GetAgentId() *string
	SetInputParam(v string) *CreateMobileAgentTaskRequest
	GetInputParam() *string
	SetName(v string) *CreateMobileAgentTaskRequest
	GetName() *string
	SetNodeIds(v string) *CreateMobileAgentTaskRequest
	GetNodeIds() *string
	SetSessionId(v string) *CreateMobileAgentTaskRequest
	GetSessionId() *string
	SetType(v string) *CreateMobileAgentTaskRequest
	GetType() *string
}

type CreateMobileAgentTaskRequest struct {
	// The agent ID.
	//
	// example:
	//
	// agt-mobileagent
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// The task input.
	//
	// example:
	//
	// Open DingTalk and reply to the first 5 unread messages
	InputParam *string `json:"InputParam,omitempty" xml:"InputParam,omitempty"`
	// The task name.
	//
	// example:
	//
	// Auto-reply to DingTalk messages
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The list of node IDs.
	//
	// example:
	//
	// cpn-csbdhjcbvfd****
	NodeIds *string `json:"NodeIds,omitempty" xml:"NodeIds,omitempty"`
	// The session ID. Tasks with the same session ID share the same context.
	//
	// example:
	//
	// ses-261005-6ekkb1m9ls00
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The task type.
	//
	// example:
	//
	// mobile_use_agent
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateMobileAgentTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMobileAgentTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateMobileAgentTaskRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *CreateMobileAgentTaskRequest) GetInputParam() *string {
	return s.InputParam
}

func (s *CreateMobileAgentTaskRequest) GetName() *string {
	return s.Name
}

func (s *CreateMobileAgentTaskRequest) GetNodeIds() *string {
	return s.NodeIds
}

func (s *CreateMobileAgentTaskRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *CreateMobileAgentTaskRequest) GetType() *string {
	return s.Type
}

func (s *CreateMobileAgentTaskRequest) SetAgentId(v string) *CreateMobileAgentTaskRequest {
	s.AgentId = &v
	return s
}

func (s *CreateMobileAgentTaskRequest) SetInputParam(v string) *CreateMobileAgentTaskRequest {
	s.InputParam = &v
	return s
}

func (s *CreateMobileAgentTaskRequest) SetName(v string) *CreateMobileAgentTaskRequest {
	s.Name = &v
	return s
}

func (s *CreateMobileAgentTaskRequest) SetNodeIds(v string) *CreateMobileAgentTaskRequest {
	s.NodeIds = &v
	return s
}

func (s *CreateMobileAgentTaskRequest) SetSessionId(v string) *CreateMobileAgentTaskRequest {
	s.SessionId = &v
	return s
}

func (s *CreateMobileAgentTaskRequest) SetType(v string) *CreateMobileAgentTaskRequest {
	s.Type = &v
	return s
}

func (s *CreateMobileAgentTaskRequest) Validate() error {
	return dara.Validate(s)
}
