// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppAgentTemplateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ModifyAppAgentTemplateShrinkRequest
	GetAppId() *string
	SetAsrConfigShrink(v string) *ModifyAppAgentTemplateShrinkRequest
	GetAsrConfigShrink() *string
	SetChatMode(v int32) *ModifyAppAgentTemplateShrinkRequest
	GetChatMode() *int32
	SetGreeting(v string) *ModifyAppAgentTemplateShrinkRequest
	GetGreeting() *string
	SetId(v string) *ModifyAppAgentTemplateShrinkRequest
	GetId() *string
	SetInterruptMode(v int32) *ModifyAppAgentTemplateShrinkRequest
	GetInterruptMode() *int32
	SetLlmConfigShrink(v string) *ModifyAppAgentTemplateShrinkRequest
	GetLlmConfigShrink() *string
	SetName(v string) *ModifyAppAgentTemplateShrinkRequest
	GetName() *string
	SetTtsConfigShrink(v string) *ModifyAppAgentTemplateShrinkRequest
	GetTtsConfigShrink() *string
	SetType(v int32) *ModifyAppAgentTemplateShrinkRequest
	GetType() *int32
}

type ModifyAppAgentTemplateShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ac7N****
	AppId           *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AsrConfigShrink *string `json:"AsrConfig,omitempty" xml:"AsrConfig,omitempty"`
	// example:
	//
	// 2
	ChatMode *int32  `json:"ChatMode,omitempty" xml:"ChatMode,omitempty"`
	Greeting *string `json:"Greeting,omitempty" xml:"Greeting,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1231231312312131231
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 2
	InterruptMode   *int32  `json:"InterruptMode,omitempty" xml:"InterruptMode,omitempty"`
	LlmConfigShrink *string `json:"LlmConfig,omitempty" xml:"LlmConfig,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 智能体模版
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	TtsConfigShrink *string `json:"TtsConfig,omitempty" xml:"TtsConfig,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ModifyAppAgentTemplateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppAgentTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyAppAgentTemplateShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *ModifyAppAgentTemplateShrinkRequest) GetAsrConfigShrink() *string {
	return s.AsrConfigShrink
}

func (s *ModifyAppAgentTemplateShrinkRequest) GetChatMode() *int32 {
	return s.ChatMode
}

func (s *ModifyAppAgentTemplateShrinkRequest) GetGreeting() *string {
	return s.Greeting
}

func (s *ModifyAppAgentTemplateShrinkRequest) GetId() *string {
	return s.Id
}

func (s *ModifyAppAgentTemplateShrinkRequest) GetInterruptMode() *int32 {
	return s.InterruptMode
}

func (s *ModifyAppAgentTemplateShrinkRequest) GetLlmConfigShrink() *string {
	return s.LlmConfigShrink
}

func (s *ModifyAppAgentTemplateShrinkRequest) GetName() *string {
	return s.Name
}

func (s *ModifyAppAgentTemplateShrinkRequest) GetTtsConfigShrink() *string {
	return s.TtsConfigShrink
}

func (s *ModifyAppAgentTemplateShrinkRequest) GetType() *int32 {
	return s.Type
}

func (s *ModifyAppAgentTemplateShrinkRequest) SetAppId(v string) *ModifyAppAgentTemplateShrinkRequest {
	s.AppId = &v
	return s
}

func (s *ModifyAppAgentTemplateShrinkRequest) SetAsrConfigShrink(v string) *ModifyAppAgentTemplateShrinkRequest {
	s.AsrConfigShrink = &v
	return s
}

func (s *ModifyAppAgentTemplateShrinkRequest) SetChatMode(v int32) *ModifyAppAgentTemplateShrinkRequest {
	s.ChatMode = &v
	return s
}

func (s *ModifyAppAgentTemplateShrinkRequest) SetGreeting(v string) *ModifyAppAgentTemplateShrinkRequest {
	s.Greeting = &v
	return s
}

func (s *ModifyAppAgentTemplateShrinkRequest) SetId(v string) *ModifyAppAgentTemplateShrinkRequest {
	s.Id = &v
	return s
}

func (s *ModifyAppAgentTemplateShrinkRequest) SetInterruptMode(v int32) *ModifyAppAgentTemplateShrinkRequest {
	s.InterruptMode = &v
	return s
}

func (s *ModifyAppAgentTemplateShrinkRequest) SetLlmConfigShrink(v string) *ModifyAppAgentTemplateShrinkRequest {
	s.LlmConfigShrink = &v
	return s
}

func (s *ModifyAppAgentTemplateShrinkRequest) SetName(v string) *ModifyAppAgentTemplateShrinkRequest {
	s.Name = &v
	return s
}

func (s *ModifyAppAgentTemplateShrinkRequest) SetTtsConfigShrink(v string) *ModifyAppAgentTemplateShrinkRequest {
	s.TtsConfigShrink = &v
	return s
}

func (s *ModifyAppAgentTemplateShrinkRequest) SetType(v int32) *ModifyAppAgentTemplateShrinkRequest {
	s.Type = &v
	return s
}

func (s *ModifyAppAgentTemplateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
