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
	// The ID of the AI agent.
	//
	// This parameter is required.
	//
	// example:
	//
	// 39f8e0bc005e4f309379701645f4****
	AIAgentId      *string                                   `json:"AIAgentId,omitempty" xml:"AIAgentId,omitempty"`
	AgentConfig    *AIAgentConfig                            `json:"AgentConfig,omitempty" xml:"AgentConfig,omitempty"`
	ChatSyncConfig *GenerateAIAgentCallRequestChatSyncConfig `json:"ChatSyncConfig,omitempty" xml:"ChatSyncConfig,omitempty" type:"Struct"`
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
	TemplateConfig *AIAgentTemplateConfig `json:"TemplateConfig,omitempty" xml:"TemplateConfig,omitempty"`
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
	return dara.Validate(s)
}

type GenerateAIAgentCallRequestChatSyncConfig struct {
	// example:
	//
	// ******005e4f309379701645f4****
	IMAIAgentId *string `json:"IMAIAgentId,omitempty" xml:"IMAIAgentId,omitempty"`
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
