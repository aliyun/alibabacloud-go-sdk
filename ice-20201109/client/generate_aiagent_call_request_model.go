// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateAIAgentCallRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAIAgentId(v string) *GenerateAIAgentCallRequest
	GetAIAgentId() *string
	SetAgentConfig(v *AIAgentConfig) *GenerateAIAgentCallRequest
	GetAgentConfig() *AIAgentConfig
	SetChatSyncConfig(v *GenerateAIAgentCallRequestChatSyncConfig) *GenerateAIAgentCallRequest
	GetChatSyncConfig() *GenerateAIAgentCallRequestChatSyncConfig
	SetExpire(v int64) *GenerateAIAgentCallRequest
	GetExpire() *int64
	SetSessionId(v string) *GenerateAIAgentCallRequest
	GetSessionId() *string
	SetTemplateConfig(v *AIAgentTemplateConfig) *GenerateAIAgentCallRequest
	GetTemplateConfig() *AIAgentTemplateConfig
	SetUserData(v string) *GenerateAIAgentCallRequest
	GetUserData() *string
	SetUserId(v string) *GenerateAIAgentCallRequest
	GetUserId() *string
}

type GenerateAIAgentCallRequest struct {
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
	AgentConfig *AIAgentConfig `json:"AgentConfig,omitempty" xml:"AgentConfig,omitempty"`
	// The chat synchronization configuration.
	ChatSyncConfig *GenerateAIAgentCallRequestChatSyncConfig `json:"ChatSyncConfig,omitempty" xml:"ChatSyncConfig,omitempty" type:"Struct"`
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
	TemplateConfig *AIAgentTemplateConfig `json:"TemplateConfig,omitempty" xml:"TemplateConfig,omitempty"`
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

func (s GenerateAIAgentCallRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateAIAgentCallRequest) GoString() string {
	return s.String()
}

func (s *GenerateAIAgentCallRequest) GetAIAgentId() *string {
	return s.AIAgentId
}

func (s *GenerateAIAgentCallRequest) GetAgentConfig() *AIAgentConfig {
	return s.AgentConfig
}

func (s *GenerateAIAgentCallRequest) GetChatSyncConfig() *GenerateAIAgentCallRequestChatSyncConfig {
	return s.ChatSyncConfig
}

func (s *GenerateAIAgentCallRequest) GetExpire() *int64 {
	return s.Expire
}

func (s *GenerateAIAgentCallRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *GenerateAIAgentCallRequest) GetTemplateConfig() *AIAgentTemplateConfig {
	return s.TemplateConfig
}

func (s *GenerateAIAgentCallRequest) GetUserData() *string {
	return s.UserData
}

func (s *GenerateAIAgentCallRequest) GetUserId() *string {
	return s.UserId
}

func (s *GenerateAIAgentCallRequest) SetAIAgentId(v string) *GenerateAIAgentCallRequest {
	s.AIAgentId = &v
	return s
}

func (s *GenerateAIAgentCallRequest) SetAgentConfig(v *AIAgentConfig) *GenerateAIAgentCallRequest {
	s.AgentConfig = v
	return s
}

func (s *GenerateAIAgentCallRequest) SetChatSyncConfig(v *GenerateAIAgentCallRequestChatSyncConfig) *GenerateAIAgentCallRequest {
	s.ChatSyncConfig = v
	return s
}

func (s *GenerateAIAgentCallRequest) SetExpire(v int64) *GenerateAIAgentCallRequest {
	s.Expire = &v
	return s
}

func (s *GenerateAIAgentCallRequest) SetSessionId(v string) *GenerateAIAgentCallRequest {
	s.SessionId = &v
	return s
}

func (s *GenerateAIAgentCallRequest) SetTemplateConfig(v *AIAgentTemplateConfig) *GenerateAIAgentCallRequest {
	s.TemplateConfig = v
	return s
}

func (s *GenerateAIAgentCallRequest) SetUserData(v string) *GenerateAIAgentCallRequest {
	s.UserData = &v
	return s
}

func (s *GenerateAIAgentCallRequest) SetUserId(v string) *GenerateAIAgentCallRequest {
	s.UserId = &v
	return s
}

func (s *GenerateAIAgentCallRequest) Validate() error {
	if s.AgentConfig != nil {
		if err := s.AgentConfig.Validate(); err != nil {
			return err
		}
	}
	if s.ChatSyncConfig != nil {
		if err := s.ChatSyncConfig.Validate(); err != nil {
			return err
		}
	}
	if s.TemplateConfig != nil {
		if err := s.TemplateConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GenerateAIAgentCallRequestChatSyncConfig struct {
	// The ID of the Instant Messaging (IM) agent.
	//
	// example:
	//
	// ******005e4f309379701645f4****
	IMAIAgentId *string `json:"IMAIAgentId,omitempty" xml:"IMAIAgentId,omitempty"`
	// The user ID of the recipient.
	//
	// example:
	//
	// 4167626d312034b2b1c3b7f2f3e41884
	ReceiverId *string `json:"ReceiverId,omitempty" xml:"ReceiverId,omitempty"`
}

func (s GenerateAIAgentCallRequestChatSyncConfig) String() string {
	return dara.Prettify(s)
}

func (s GenerateAIAgentCallRequestChatSyncConfig) GoString() string {
	return s.String()
}

func (s *GenerateAIAgentCallRequestChatSyncConfig) GetIMAIAgentId() *string {
	return s.IMAIAgentId
}

func (s *GenerateAIAgentCallRequestChatSyncConfig) GetReceiverId() *string {
	return s.ReceiverId
}

func (s *GenerateAIAgentCallRequestChatSyncConfig) SetIMAIAgentId(v string) *GenerateAIAgentCallRequestChatSyncConfig {
	s.IMAIAgentId = &v
	return s
}

func (s *GenerateAIAgentCallRequestChatSyncConfig) SetReceiverId(v string) *GenerateAIAgentCallRequestChatSyncConfig {
	s.ReceiverId = &v
	return s
}

func (s *GenerateAIAgentCallRequestChatSyncConfig) Validate() error {
	return dara.Validate(s)
}
