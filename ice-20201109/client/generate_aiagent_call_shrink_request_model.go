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
	// The AI agent ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 39f8e0bc005e4f309379701645f4****
	AIAgentId *string `json:"AIAgentId,omitempty" xml:"AIAgentId,omitempty"`
	// The agent template configuration. The configuration you provide merges with the agent template configuration in the console. If you omit this parameter, the agent uses the default configuration from the console.
	//
	// > Compatibility with `TemplateConfig`: Fields in `AgentConfig` take precedence. If a field is specified in `TemplateConfig` but not in `AgentConfig`, the system uses the value from `TemplateConfig`. We recommend using `AgentConfig` instead of `TemplateConfig`.
	AgentConfigShrink *string `json:"AgentConfig,omitempty" xml:"AgentConfig,omitempty"`
	// The chat synchronization configuration.
	ChatSyncConfigShrink *string `json:"ChatSyncConfig,omitempty" xml:"ChatSyncConfig,omitempty"`
	// Optional. The expiration time of the token in seconds. Default value: 3600. Value range: 0 to 604800.
	//
	// example:
	//
	// 3600
	Expire *int64 `json:"Expire,omitempty" xml:"Expire,omitempty"`
	// A unique identifier for the session. If not provided, a new session is created.
	//
	// example:
	//
	// fw1gr0bc005e4f309379701645f4****
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// Deprecated
	//
	// - This configuration merges with the agent template configuration in the console.
	//
	// - If you omit this parameter, the agent uses the default configuration from the console.
	//
	// > The agent template configuration. This parameter is deprecated. Use the AgentConfig parameter instead.
	TemplateConfigShrink *string `json:"TemplateConfig,omitempty" xml:"TemplateConfig,omitempty"`
	// User data.
	//
	// example:
	//
	// {"Email":"johndoe@example.com","Preferences":{"Language":"en"}}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// The username in the channel. If you do not specify a username, one is automatically generated. The username can be up to 64 characters in length.
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
