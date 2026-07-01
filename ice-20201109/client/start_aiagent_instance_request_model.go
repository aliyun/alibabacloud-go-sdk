// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartAIAgentInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAIAgentId(v string) *StartAIAgentInstanceRequest
	GetAIAgentId() *string
	SetAgentConfig(v *AIAgentConfig) *StartAIAgentInstanceRequest
	GetAgentConfig() *AIAgentConfig
	SetChatSyncConfig(v *StartAIAgentInstanceRequestChatSyncConfig) *StartAIAgentInstanceRequest
	GetChatSyncConfig() *StartAIAgentInstanceRequestChatSyncConfig
	SetRuntimeConfig(v *AIAgentRuntimeConfig) *StartAIAgentInstanceRequest
	GetRuntimeConfig() *AIAgentRuntimeConfig
	SetSessionId(v string) *StartAIAgentInstanceRequest
	GetSessionId() *string
	SetTemplateConfig(v *AIAgentTemplateConfig) *StartAIAgentInstanceRequest
	GetTemplateConfig() *AIAgentTemplateConfig
	SetUserData(v string) *StartAIAgentInstanceRequest
	GetUserData() *string
}

type StartAIAgentInstanceRequest struct {
	// The agent ID configured in the [IMS console](https://ims.console.aliyun.com/ai/robot/list).
	//
	// This parameter is required.
	//
	// example:
	//
	// 39f8e0bc005e4f309379701645f4****
	AIAgentId *string `json:"AIAgentId,omitempty" xml:"AIAgentId,omitempty"`
	// The agent template configuration. Values you provide merge with the template configuration set in the console. If you omit this parameter, the agent uses its default configuration from the console.
	//
	// > This field is compatible with TemplateConfig. Fields in AgentConfig take precedence. If TemplateConfig contains fields not defined in AgentConfig, those fields are used. Use AgentConfig instead of TemplateConfig.
	AgentConfig *AIAgentConfig `json:"AgentConfig,omitempty" xml:"AgentConfig,omitempty"`
	// The chat history synchronization configuration.
	ChatSyncConfig *StartAIAgentInstanceRequestChatSyncConfig `json:"ChatSyncConfig,omitempty" xml:"ChatSyncConfig,omitempty" type:"Struct"`
	// The configuration required for the agent at runtime.
	//
	// This parameter is required.
	RuntimeConfig *AIAgentRuntimeConfig `json:"RuntimeConfig,omitempty" xml:"RuntimeConfig,omitempty"`
	// A unique identifier for the chat session. This parameter is optional.
	//
	// example:
	//
	// f213fbc005e4f309379701645f4****
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// Deprecated
	//
	// The agent template configuration. Values you provide merge with the template configuration set in the console. If you omit this parameter, the agent uses its default configuration from the console.
	//
	// > The agent template configuration. This field is deprecated. See the AgentConfig field.
	TemplateConfig *AIAgentTemplateConfig `json:"TemplateConfig,omitempty" xml:"TemplateConfig,omitempty"`
	// User-defined data.
	//
	// example:
	//
	// {"Email":"johndoe@example.com","Preferences":{"Language":"en"}}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s StartAIAgentInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s StartAIAgentInstanceRequest) GoString() string {
	return s.String()
}

func (s *StartAIAgentInstanceRequest) GetAIAgentId() *string {
	return s.AIAgentId
}

func (s *StartAIAgentInstanceRequest) GetAgentConfig() *AIAgentConfig {
	return s.AgentConfig
}

func (s *StartAIAgentInstanceRequest) GetChatSyncConfig() *StartAIAgentInstanceRequestChatSyncConfig {
	return s.ChatSyncConfig
}

func (s *StartAIAgentInstanceRequest) GetRuntimeConfig() *AIAgentRuntimeConfig {
	return s.RuntimeConfig
}

func (s *StartAIAgentInstanceRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *StartAIAgentInstanceRequest) GetTemplateConfig() *AIAgentTemplateConfig {
	return s.TemplateConfig
}

func (s *StartAIAgentInstanceRequest) GetUserData() *string {
	return s.UserData
}

func (s *StartAIAgentInstanceRequest) SetAIAgentId(v string) *StartAIAgentInstanceRequest {
	s.AIAgentId = &v
	return s
}

func (s *StartAIAgentInstanceRequest) SetAgentConfig(v *AIAgentConfig) *StartAIAgentInstanceRequest {
	s.AgentConfig = v
	return s
}

func (s *StartAIAgentInstanceRequest) SetChatSyncConfig(v *StartAIAgentInstanceRequestChatSyncConfig) *StartAIAgentInstanceRequest {
	s.ChatSyncConfig = v
	return s
}

func (s *StartAIAgentInstanceRequest) SetRuntimeConfig(v *AIAgentRuntimeConfig) *StartAIAgentInstanceRequest {
	s.RuntimeConfig = v
	return s
}

func (s *StartAIAgentInstanceRequest) SetSessionId(v string) *StartAIAgentInstanceRequest {
	s.SessionId = &v
	return s
}

func (s *StartAIAgentInstanceRequest) SetTemplateConfig(v *AIAgentTemplateConfig) *StartAIAgentInstanceRequest {
	s.TemplateConfig = v
	return s
}

func (s *StartAIAgentInstanceRequest) SetUserData(v string) *StartAIAgentInstanceRequest {
	s.UserData = &v
	return s
}

func (s *StartAIAgentInstanceRequest) Validate() error {
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
	if s.RuntimeConfig != nil {
		if err := s.RuntimeConfig.Validate(); err != nil {
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

type StartAIAgentInstanceRequestChatSyncConfig struct {
	// The IM agent ID.
	//
	// example:
	//
	// ******005e4f309379701645f4****
	IMAIAgentId *string `json:"IMAIAgentId,omitempty" xml:"IMAIAgentId,omitempty"`
	// The receiver user ID.
	//
	// example:
	//
	// 4167626d312034b2b1c3b7f2f3e41884
	ReceiverId *string `json:"ReceiverId,omitempty" xml:"ReceiverId,omitempty"`
}

func (s StartAIAgentInstanceRequestChatSyncConfig) String() string {
	return dara.Prettify(s)
}

func (s StartAIAgentInstanceRequestChatSyncConfig) GoString() string {
	return s.String()
}

func (s *StartAIAgentInstanceRequestChatSyncConfig) GetIMAIAgentId() *string {
	return s.IMAIAgentId
}

func (s *StartAIAgentInstanceRequestChatSyncConfig) GetReceiverId() *string {
	return s.ReceiverId
}

func (s *StartAIAgentInstanceRequestChatSyncConfig) SetIMAIAgentId(v string) *StartAIAgentInstanceRequestChatSyncConfig {
	s.IMAIAgentId = &v
	return s
}

func (s *StartAIAgentInstanceRequestChatSyncConfig) SetReceiverId(v string) *StartAIAgentInstanceRequestChatSyncConfig {
	s.ReceiverId = &v
	return s
}

func (s *StartAIAgentInstanceRequestChatSyncConfig) Validate() error {
	return dara.Validate(s)
}
