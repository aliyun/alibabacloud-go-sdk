// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAiproEmbedConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentName(v string) *ModifyAiproEmbedConfigRequest
	GetAgentName() *string
	SetAiProId(v string) *ModifyAiproEmbedConfigRequest
	GetAiProId() *string
	SetChatHistoryStatus(v int32) *ModifyAiproEmbedConfigRequest
	GetChatHistoryStatus() *int32
	SetConfig(v string) *ModifyAiproEmbedConfigRequest
	GetConfig() *string
	SetMemoryEnabled(v int32) *ModifyAiproEmbedConfigRequest
	GetMemoryEnabled() *int32
	SetShowName(v string) *ModifyAiproEmbedConfigRequest
	GetShowName() *string
	SetUploadEnabled(v int32) *ModifyAiproEmbedConfigRequest
	GetUploadEnabled() *int32
}

type ModifyAiproEmbedConfigRequest struct {
	// The nickname of the agent. The default value is Intelligent Data Analytics Assistant when you create a configuration.
	//
	// example:
	//
	// test
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// The AIPro configuration ID. If this parameter is not specified or is left empty, a new configuration is created. If this parameter is specified, the existing configuration is updated.
	//
	// example:
	//
	// test
	AiProId *string `json:"AiProId,omitempty" xml:"AiProId,omitempty"`
	// Specifies whether to display the chat history. Valid values:
	//
	// - 1: Display.
	//
	// - 0: Do not display.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	ChatHistoryStatus *int32 `json:"ChatHistoryStatus,omitempty" xml:"ChatHistoryStatus,omitempty"`
	// The JSON string of the embedding appearance configuration saved on the frontend.
	//
	// example:
	//
	// test
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// Specifies whether to enable memory. Valid values:
	//
	// - 1: Enable.
	//
	// - 0: Disable.
	//
	// Default value: 0.
	//
	// example:
	//
	// 1
	MemoryEnabled *int32 `json:"MemoryEnabled,omitempty" xml:"MemoryEnabled,omitempty"`
	// The name of the configuration. This parameter is required and must be unique within the organization when you create a configuration. If this parameter is left empty when you update a configuration, the name is not updated.
	//
	// example:
	//
	// test
	ShowName *string `json:"ShowName,omitempty" xml:"ShowName,omitempty"`
	// Specifies whether to allow file uploads. Valid values:
	//
	// - 1: Allow.
	//
	// - 0: Do not allow.
	//
	// Default value: 0.
	//
	// example:
	//
	// 1
	UploadEnabled *int32 `json:"UploadEnabled,omitempty" xml:"UploadEnabled,omitempty"`
}

func (s ModifyAiproEmbedConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAiproEmbedConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyAiproEmbedConfigRequest) GetAgentName() *string {
	return s.AgentName
}

func (s *ModifyAiproEmbedConfigRequest) GetAiProId() *string {
	return s.AiProId
}

func (s *ModifyAiproEmbedConfigRequest) GetChatHistoryStatus() *int32 {
	return s.ChatHistoryStatus
}

func (s *ModifyAiproEmbedConfigRequest) GetConfig() *string {
	return s.Config
}

func (s *ModifyAiproEmbedConfigRequest) GetMemoryEnabled() *int32 {
	return s.MemoryEnabled
}

func (s *ModifyAiproEmbedConfigRequest) GetShowName() *string {
	return s.ShowName
}

func (s *ModifyAiproEmbedConfigRequest) GetUploadEnabled() *int32 {
	return s.UploadEnabled
}

func (s *ModifyAiproEmbedConfigRequest) SetAgentName(v string) *ModifyAiproEmbedConfigRequest {
	s.AgentName = &v
	return s
}

func (s *ModifyAiproEmbedConfigRequest) SetAiProId(v string) *ModifyAiproEmbedConfigRequest {
	s.AiProId = &v
	return s
}

func (s *ModifyAiproEmbedConfigRequest) SetChatHistoryStatus(v int32) *ModifyAiproEmbedConfigRequest {
	s.ChatHistoryStatus = &v
	return s
}

func (s *ModifyAiproEmbedConfigRequest) SetConfig(v string) *ModifyAiproEmbedConfigRequest {
	s.Config = &v
	return s
}

func (s *ModifyAiproEmbedConfigRequest) SetMemoryEnabled(v int32) *ModifyAiproEmbedConfigRequest {
	s.MemoryEnabled = &v
	return s
}

func (s *ModifyAiproEmbedConfigRequest) SetShowName(v string) *ModifyAiproEmbedConfigRequest {
	s.ShowName = &v
	return s
}

func (s *ModifyAiproEmbedConfigRequest) SetUploadEnabled(v int32) *ModifyAiproEmbedConfigRequest {
	s.UploadEnabled = &v
	return s
}

func (s *ModifyAiproEmbedConfigRequest) Validate() error {
	return dara.Validate(s)
}
