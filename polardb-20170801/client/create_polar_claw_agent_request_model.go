// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePolarClawAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *CreatePolarClawAgentRequest
	GetAgentId() *string
	SetApplicationId(v string) *CreatePolarClawAgentRequest
	GetApplicationId() *string
	SetAvatar(v string) *CreatePolarClawAgentRequest
	GetAvatar() *string
	SetEmoji(v string) *CreatePolarClawAgentRequest
	GetEmoji() *string
	SetRestart(v bool) *CreatePolarClawAgentRequest
	GetRestart() *bool
	SetWorkspace(v string) *CreatePolarClawAgentRequest
	GetWorkspace() *string
}

type CreatePolarClawAgentRequest struct {
	// The agent ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// work
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The avatar\\"s URL or path.
	//
	// example:
	//
	// test
	Avatar *string `json:"Avatar,omitempty" xml:"Avatar,omitempty"`
	// The emoji character.
	//
	// example:
	//
	// U+1F99E
	Emoji *string `json:"Emoji,omitempty" xml:"Emoji,omitempty"`
	// Specifies whether to restart the gateway after the agent is created. The default value is `true`.
	//
	// example:
	//
	// true
	Restart *bool `json:"Restart,omitempty" xml:"Restart,omitempty"`
	// The absolute path of the agent workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// /home/node/.openclaw/workspace-work
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s CreatePolarClawAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePolarClawAgentRequest) GoString() string {
	return s.String()
}

func (s *CreatePolarClawAgentRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *CreatePolarClawAgentRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *CreatePolarClawAgentRequest) GetAvatar() *string {
	return s.Avatar
}

func (s *CreatePolarClawAgentRequest) GetEmoji() *string {
	return s.Emoji
}

func (s *CreatePolarClawAgentRequest) GetRestart() *bool {
	return s.Restart
}

func (s *CreatePolarClawAgentRequest) GetWorkspace() *string {
	return s.Workspace
}

func (s *CreatePolarClawAgentRequest) SetAgentId(v string) *CreatePolarClawAgentRequest {
	s.AgentId = &v
	return s
}

func (s *CreatePolarClawAgentRequest) SetApplicationId(v string) *CreatePolarClawAgentRequest {
	s.ApplicationId = &v
	return s
}

func (s *CreatePolarClawAgentRequest) SetAvatar(v string) *CreatePolarClawAgentRequest {
	s.Avatar = &v
	return s
}

func (s *CreatePolarClawAgentRequest) SetEmoji(v string) *CreatePolarClawAgentRequest {
	s.Emoji = &v
	return s
}

func (s *CreatePolarClawAgentRequest) SetRestart(v bool) *CreatePolarClawAgentRequest {
	s.Restart = &v
	return s
}

func (s *CreatePolarClawAgentRequest) SetWorkspace(v string) *CreatePolarClawAgentRequest {
	s.Workspace = &v
	return s
}

func (s *CreatePolarClawAgentRequest) Validate() error {
	return dara.Validate(s)
}
