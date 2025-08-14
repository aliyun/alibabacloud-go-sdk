// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartAIAgentInstanceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAIAgentId(v string) *StartAIAgentInstanceShrinkRequest
	GetAIAgentId() *string
	SetAgentConfigShrink(v string) *StartAIAgentInstanceShrinkRequest
	GetAgentConfigShrink() *string
	SetChatSyncConfigShrink(v string) *StartAIAgentInstanceShrinkRequest
	GetChatSyncConfigShrink() *string
	SetRuntimeConfigShrink(v string) *StartAIAgentInstanceShrinkRequest
	GetRuntimeConfigShrink() *string
	SetSessionId(v string) *StartAIAgentInstanceShrinkRequest
	GetSessionId() *string
	SetTemplateConfigShrink(v string) *StartAIAgentInstanceShrinkRequest
	GetTemplateConfigShrink() *string
	SetUserData(v string) *StartAIAgentInstanceShrinkRequest
	GetUserData() *string
}

type StartAIAgentInstanceShrinkRequest struct {
	// The ID of the AI agent created in the [IMS](https://ims.console.aliyun.com/ai/robot/list) console.
	//
	// This parameter is required.
	//
	// example:
	//
	// 39f8e0bc005e4f309379701645f4****
	AIAgentId         *string `json:"AIAgentId,omitempty" xml:"AIAgentId,omitempty"`
	AgentConfigShrink *string `json:"AgentConfig,omitempty" xml:"AgentConfig,omitempty"`
	// 同步聊天记录配置。
	ChatSyncConfigShrink *string `json:"ChatSyncConfig,omitempty" xml:"ChatSyncConfig,omitempty"`
	// This parameter is required.
	RuntimeConfigShrink *string `json:"RuntimeConfig,omitempty" xml:"RuntimeConfig,omitempty"`
	// example:
	//
	// f213fbc005e4f309379701645f4****
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// Deprecated
	TemplateConfigShrink *string `json:"TemplateConfig,omitempty" xml:"TemplateConfig,omitempty"`
	// example:
	//
	// {"Email":"johndoe@example.com","Preferences":{"Language":"en"}}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s StartAIAgentInstanceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s StartAIAgentInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *StartAIAgentInstanceShrinkRequest) GetAIAgentId() *string {
	return s.AIAgentId
}

func (s *StartAIAgentInstanceShrinkRequest) GetAgentConfigShrink() *string {
	return s.AgentConfigShrink
}

func (s *StartAIAgentInstanceShrinkRequest) GetChatSyncConfigShrink() *string {
	return s.ChatSyncConfigShrink
}

func (s *StartAIAgentInstanceShrinkRequest) GetRuntimeConfigShrink() *string {
	return s.RuntimeConfigShrink
}

func (s *StartAIAgentInstanceShrinkRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *StartAIAgentInstanceShrinkRequest) GetTemplateConfigShrink() *string {
	return s.TemplateConfigShrink
}

func (s *StartAIAgentInstanceShrinkRequest) GetUserData() *string {
	return s.UserData
}

func (s *StartAIAgentInstanceShrinkRequest) SetAIAgentId(v string) *StartAIAgentInstanceShrinkRequest {
	s.AIAgentId = &v
	return s
}

func (s *StartAIAgentInstanceShrinkRequest) SetAgentConfigShrink(v string) *StartAIAgentInstanceShrinkRequest {
	s.AgentConfigShrink = &v
	return s
}

func (s *StartAIAgentInstanceShrinkRequest) SetChatSyncConfigShrink(v string) *StartAIAgentInstanceShrinkRequest {
	s.ChatSyncConfigShrink = &v
	return s
}

func (s *StartAIAgentInstanceShrinkRequest) SetRuntimeConfigShrink(v string) *StartAIAgentInstanceShrinkRequest {
	s.RuntimeConfigShrink = &v
	return s
}

func (s *StartAIAgentInstanceShrinkRequest) SetSessionId(v string) *StartAIAgentInstanceShrinkRequest {
	s.SessionId = &v
	return s
}

func (s *StartAIAgentInstanceShrinkRequest) SetTemplateConfigShrink(v string) *StartAIAgentInstanceShrinkRequest {
	s.TemplateConfigShrink = &v
	return s
}

func (s *StartAIAgentInstanceShrinkRequest) SetUserData(v string) *StartAIAgentInstanceShrinkRequest {
	s.UserData = &v
	return s
}

func (s *StartAIAgentInstanceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
