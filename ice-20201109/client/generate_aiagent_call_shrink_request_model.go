// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateAIAgentCallShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAIAgentId(v string) *GenerateAIAgentCallShrinkRequest
	GetAIAgentId() *string
	SetAgentConfigShrink(v string) *GenerateAIAgentCallShrinkRequest
	GetAgentConfigShrink() *string
	SetChatSyncConfigShrink(v string) *GenerateAIAgentCallShrinkRequest
	GetChatSyncConfigShrink() *string
	SetExpire(v int64) *GenerateAIAgentCallShrinkRequest
	GetExpire() *int64
	SetSessionId(v string) *GenerateAIAgentCallShrinkRequest
	GetSessionId() *string
	SetTemplateConfigShrink(v string) *GenerateAIAgentCallShrinkRequest
	GetTemplateConfigShrink() *string
	SetUserData(v string) *GenerateAIAgentCallShrinkRequest
	GetUserData() *string
	SetUserId(v string) *GenerateAIAgentCallShrinkRequest
	GetUserId() *string
}

type GenerateAIAgentCallShrinkRequest struct {
	// The ID of the AI agent.
	//
	// This parameter is required.
	//
	// example:
	//
	// 39f8e0bc005e4f309379701645f4****
	AIAgentId            *string `json:"AIAgentId,omitempty" xml:"AIAgentId,omitempty"`
	AgentConfigShrink    *string `json:"AgentConfig,omitempty" xml:"AgentConfig,omitempty"`
	ChatSyncConfigShrink *string `json:"ChatSyncConfig,omitempty" xml:"ChatSyncConfig,omitempty"`
	// The time when the token expires. Unit: seconds. Default value: 3600. Valid values: 0 to 604800.
	//
	// example:
	//
	// 3600
	Expire *int64 `json:"Expire,omitempty" xml:"Expire,omitempty"`
	// example:
	//
	// fw1gr0bc005e4f309379701645f4****
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// Deprecated
	//
	// The template configurations of the AI agent. The specified configurations are merged with the template configurations that are specified in the console. If you do not specify this parameter, the system uses the default configurations for an AI agent created in the console.
	TemplateConfigShrink *string `json:"TemplateConfig,omitempty" xml:"TemplateConfig,omitempty"`
	// example:
	//
	// {"Email":"johndoe@example.com","Preferences":{"Language":"en"}}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// The username of the AI agent in the channel. If you do not specify this parameter, the system automatically generates a username. The value can be up to 64 characters in length.
	//
	// example:
	//
	// 877ae632caae49b1afc81c2e8194ffb4
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GenerateAIAgentCallShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateAIAgentCallShrinkRequest) GoString() string {
	return s.String()
}

func (s *GenerateAIAgentCallShrinkRequest) GetAIAgentId() *string {
	return s.AIAgentId
}

func (s *GenerateAIAgentCallShrinkRequest) GetAgentConfigShrink() *string {
	return s.AgentConfigShrink
}

func (s *GenerateAIAgentCallShrinkRequest) GetChatSyncConfigShrink() *string {
	return s.ChatSyncConfigShrink
}

func (s *GenerateAIAgentCallShrinkRequest) GetExpire() *int64 {
	return s.Expire
}

func (s *GenerateAIAgentCallShrinkRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *GenerateAIAgentCallShrinkRequest) GetTemplateConfigShrink() *string {
	return s.TemplateConfigShrink
}

func (s *GenerateAIAgentCallShrinkRequest) GetUserData() *string {
	return s.UserData
}

func (s *GenerateAIAgentCallShrinkRequest) GetUserId() *string {
	return s.UserId
}

func (s *GenerateAIAgentCallShrinkRequest) SetAIAgentId(v string) *GenerateAIAgentCallShrinkRequest {
	s.AIAgentId = &v
	return s
}

func (s *GenerateAIAgentCallShrinkRequest) SetAgentConfigShrink(v string) *GenerateAIAgentCallShrinkRequest {
	s.AgentConfigShrink = &v
	return s
}

func (s *GenerateAIAgentCallShrinkRequest) SetChatSyncConfigShrink(v string) *GenerateAIAgentCallShrinkRequest {
	s.ChatSyncConfigShrink = &v
	return s
}

func (s *GenerateAIAgentCallShrinkRequest) SetExpire(v int64) *GenerateAIAgentCallShrinkRequest {
	s.Expire = &v
	return s
}

func (s *GenerateAIAgentCallShrinkRequest) SetSessionId(v string) *GenerateAIAgentCallShrinkRequest {
	s.SessionId = &v
	return s
}

func (s *GenerateAIAgentCallShrinkRequest) SetTemplateConfigShrink(v string) *GenerateAIAgentCallShrinkRequest {
	s.TemplateConfigShrink = &v
	return s
}

func (s *GenerateAIAgentCallShrinkRequest) SetUserData(v string) *GenerateAIAgentCallShrinkRequest {
	s.UserData = &v
	return s
}

func (s *GenerateAIAgentCallShrinkRequest) SetUserId(v string) *GenerateAIAgentCallShrinkRequest {
	s.UserId = &v
	return s
}

func (s *GenerateAIAgentCallShrinkRequest) Validate() error {
	return dara.Validate(s)
}
