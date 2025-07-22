// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppAgentTemplateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CreateAppAgentTemplateShrinkRequest
	GetAppId() *string
	SetAsrConfigShrink(v string) *CreateAppAgentTemplateShrinkRequest
	GetAsrConfigShrink() *string
	SetChatMode(v int32) *CreateAppAgentTemplateShrinkRequest
	GetChatMode() *int32
	SetGreeting(v string) *CreateAppAgentTemplateShrinkRequest
	GetGreeting() *string
	SetInterruptMode(v int32) *CreateAppAgentTemplateShrinkRequest
	GetInterruptMode() *int32
	SetLlmConfigShrink(v string) *CreateAppAgentTemplateShrinkRequest
	GetLlmConfigShrink() *string
	SetName(v string) *CreateAppAgentTemplateShrinkRequest
	GetName() *string
	SetTtsConfigShrink(v string) *CreateAppAgentTemplateShrinkRequest
	GetTtsConfigShrink() *string
	SetType(v int32) *CreateAppAgentTemplateShrinkRequest
	GetType() *int32
}

type CreateAppAgentTemplateShrinkRequest struct {
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

func (s CreateAppAgentTemplateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAppAgentTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateAppAgentTemplateShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateAppAgentTemplateShrinkRequest) GetAsrConfigShrink() *string {
	return s.AsrConfigShrink
}

func (s *CreateAppAgentTemplateShrinkRequest) GetChatMode() *int32 {
	return s.ChatMode
}

func (s *CreateAppAgentTemplateShrinkRequest) GetGreeting() *string {
	return s.Greeting
}

func (s *CreateAppAgentTemplateShrinkRequest) GetInterruptMode() *int32 {
	return s.InterruptMode
}

func (s *CreateAppAgentTemplateShrinkRequest) GetLlmConfigShrink() *string {
	return s.LlmConfigShrink
}

func (s *CreateAppAgentTemplateShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateAppAgentTemplateShrinkRequest) GetTtsConfigShrink() *string {
	return s.TtsConfigShrink
}

func (s *CreateAppAgentTemplateShrinkRequest) GetType() *int32 {
	return s.Type
}

func (s *CreateAppAgentTemplateShrinkRequest) SetAppId(v string) *CreateAppAgentTemplateShrinkRequest {
	s.AppId = &v
	return s
}

func (s *CreateAppAgentTemplateShrinkRequest) SetAsrConfigShrink(v string) *CreateAppAgentTemplateShrinkRequest {
	s.AsrConfigShrink = &v
	return s
}

func (s *CreateAppAgentTemplateShrinkRequest) SetChatMode(v int32) *CreateAppAgentTemplateShrinkRequest {
	s.ChatMode = &v
	return s
}

func (s *CreateAppAgentTemplateShrinkRequest) SetGreeting(v string) *CreateAppAgentTemplateShrinkRequest {
	s.Greeting = &v
	return s
}

func (s *CreateAppAgentTemplateShrinkRequest) SetInterruptMode(v int32) *CreateAppAgentTemplateShrinkRequest {
	s.InterruptMode = &v
	return s
}

func (s *CreateAppAgentTemplateShrinkRequest) SetLlmConfigShrink(v string) *CreateAppAgentTemplateShrinkRequest {
	s.LlmConfigShrink = &v
	return s
}

func (s *CreateAppAgentTemplateShrinkRequest) SetName(v string) *CreateAppAgentTemplateShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateAppAgentTemplateShrinkRequest) SetTtsConfigShrink(v string) *CreateAppAgentTemplateShrinkRequest {
	s.TtsConfigShrink = &v
	return s
}

func (s *CreateAppAgentTemplateShrinkRequest) SetType(v int32) *CreateAppAgentTemplateShrinkRequest {
	s.Type = &v
	return s
}

func (s *CreateAppAgentTemplateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
