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
	// example:
	//
	// test
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// example:
	//
	// test
	AiProId *string `json:"AiProId,omitempty" xml:"AiProId,omitempty"`
	// example:
	//
	// 1
	ChatHistoryStatus *int32 `json:"ChatHistoryStatus,omitempty" xml:"ChatHistoryStatus,omitempty"`
	// example:
	//
	// test
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// example:
	//
	// 1
	MemoryEnabled *int32 `json:"MemoryEnabled,omitempty" xml:"MemoryEnabled,omitempty"`
	// example:
	//
	// test
	ShowName *string `json:"ShowName,omitempty" xml:"ShowName,omitempty"`
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
