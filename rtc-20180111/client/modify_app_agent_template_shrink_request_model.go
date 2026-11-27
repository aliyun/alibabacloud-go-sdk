// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppAgentTemplateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentSilenceConfigShrink(v string) *ModifyAppAgentTemplateShrinkRequest
	GetAgentSilenceConfigShrink() *string
	SetAmbientSoundConfigShrink(v string) *ModifyAppAgentTemplateShrinkRequest
	GetAmbientSoundConfigShrink() *string
	SetAppId(v string) *ModifyAppAgentTemplateShrinkRequest
	GetAppId() *string
	SetAsrConfigShrink(v string) *ModifyAppAgentTemplateShrinkRequest
	GetAsrConfigShrink() *string
	SetBackChannelConfigShrink(v string) *ModifyAppAgentTemplateShrinkRequest
	GetBackChannelConfigShrink() *string
	SetChatMode(v int32) *ModifyAppAgentTemplateShrinkRequest
	GetChatMode() *int32
	SetEnableVideoUnderstanding(v bool) *ModifyAppAgentTemplateShrinkRequest
	GetEnableVideoUnderstanding() *bool
	SetGreeting(v string) *ModifyAppAgentTemplateShrinkRequest
	GetGreeting() *string
	SetId(v string) *ModifyAppAgentTemplateShrinkRequest
	GetId() *string
	SetInterruptConfigShrink(v string) *ModifyAppAgentTemplateShrinkRequest
	GetInterruptConfigShrink() *string
	SetInterruptMode(v int32) *ModifyAppAgentTemplateShrinkRequest
	GetInterruptMode() *int32
	SetLlmConfigShrink(v string) *ModifyAppAgentTemplateShrinkRequest
	GetLlmConfigShrink() *string
	SetName(v string) *ModifyAppAgentTemplateShrinkRequest
	GetName() *string
	SetPreferVideo(v int32) *ModifyAppAgentTemplateShrinkRequest
	GetPreferVideo() *int32
	SetTtsConfigShrink(v string) *ModifyAppAgentTemplateShrinkRequest
	GetTtsConfigShrink() *string
	SetType(v int32) *ModifyAppAgentTemplateShrinkRequest
	GetType() *int32
}

type ModifyAppAgentTemplateShrinkRequest struct {
	// The agent silence configuration.
	AgentSilenceConfigShrink *string `json:"AgentSilenceConfig,omitempty" xml:"AgentSilenceConfig,omitempty"`
	// The agent ambient environment configuration. If this parameter is not specified, the ambient environment configuration is not used.
	AmbientSoundConfigShrink *string `json:"AmbientSoundConfig,omitempty" xml:"AmbientSoundConfig,omitempty"`
	// The application ID. You can create and query the ID in the console. Only a single ID is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// ac7N****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The speech-to-text settings.
	AsrConfigShrink *string `json:"AsrConfig,omitempty" xml:"AsrConfig,omitempty"`
	// The agent backchannel configuration.
	BackChannelConfigShrink *string `json:"BackChannelConfig,omitempty" xml:"BackChannelConfig,omitempty"`
	// The conversation mode. Valid values:
	//
	// - 1: Natural conversation mode (default).
	//
	// - 2: Walkie-talkie mode.
	//
	// example:
	//
	// 2
	ChatMode *int32 `json:"ChatMode,omitempty" xml:"ChatMode,omitempty"`
	// If enabled, the LLM must support multimodal data.
	EnableVideoUnderstanding *bool `json:"EnableVideoUnderstanding,omitempty" xml:"EnableVideoUnderstanding,omitempty"`
	// The greeting message. The maximum length is 500 characters.
	//
	// example:
	//
	// Good morning, my friend
	Greeting *string `json:"Greeting,omitempty" xml:"Greeting,omitempty"`
	// The template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1231231312312131231
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The agent interruption configuration.
	InterruptConfigShrink *string `json:"InterruptConfig,omitempty" xml:"InterruptConfig,omitempty"`
	// The interruption mode. Valid values:
	//
	// - 1: Interruption supported (default).
	//
	// - 2: Interruption disabled.
	//
	// example:
	//
	// 2
	InterruptMode *int32 `json:"InterruptMode,omitempty" xml:"InterruptMode,omitempty"`
	// The large language model settings.
	LlmConfigShrink *string `json:"LlmConfig,omitempty" xml:"LlmConfig,omitempty"`
	// The template name.
	//
	// This parameter is required.
	//
	// example:
	//
	// Agent template
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The preferred video stream for the agent to understand after visual understanding is enabled. The agent can understand only one video stream. When both the camera stream and the screen sharing stream are pushed:
	//
	// 1 - The agent prioritizes understanding the camera stream.
	//
	// 2 - The agent prioritizes understanding the screen sharing stream.
	//
	// By default, the agent prioritizes understanding the camera stream.
	//
	// example:
	//
	// 2
	PreferVideo *int32 `json:"PreferVideo,omitempty" xml:"PreferVideo,omitempty"`
	// The text-to-speech settings.
	TtsConfigShrink *string `json:"TtsConfig,omitempty" xml:"TtsConfig,omitempty"`
	// The agent type. Valid values:
	//
	// - 1: Voice call (default).
	//
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

func (s *ModifyAppAgentTemplateShrinkRequest) GetAgentSilenceConfigShrink() *string {
	return s.AgentSilenceConfigShrink
}

func (s *ModifyAppAgentTemplateShrinkRequest) GetAmbientSoundConfigShrink() *string {
	return s.AmbientSoundConfigShrink
}

func (s *ModifyAppAgentTemplateShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *ModifyAppAgentTemplateShrinkRequest) GetAsrConfigShrink() *string {
	return s.AsrConfigShrink
}

func (s *ModifyAppAgentTemplateShrinkRequest) GetBackChannelConfigShrink() *string {
	return s.BackChannelConfigShrink
}

func (s *ModifyAppAgentTemplateShrinkRequest) GetChatMode() *int32 {
	return s.ChatMode
}

func (s *ModifyAppAgentTemplateShrinkRequest) GetEnableVideoUnderstanding() *bool {
	return s.EnableVideoUnderstanding
}

func (s *ModifyAppAgentTemplateShrinkRequest) GetGreeting() *string {
	return s.Greeting
}

func (s *ModifyAppAgentTemplateShrinkRequest) GetId() *string {
	return s.Id
}

func (s *ModifyAppAgentTemplateShrinkRequest) GetInterruptConfigShrink() *string {
	return s.InterruptConfigShrink
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

func (s *ModifyAppAgentTemplateShrinkRequest) GetPreferVideo() *int32 {
	return s.PreferVideo
}

func (s *ModifyAppAgentTemplateShrinkRequest) GetTtsConfigShrink() *string {
	return s.TtsConfigShrink
}

func (s *ModifyAppAgentTemplateShrinkRequest) GetType() *int32 {
	return s.Type
}

func (s *ModifyAppAgentTemplateShrinkRequest) SetAgentSilenceConfigShrink(v string) *ModifyAppAgentTemplateShrinkRequest {
	s.AgentSilenceConfigShrink = &v
	return s
}

func (s *ModifyAppAgentTemplateShrinkRequest) SetAmbientSoundConfigShrink(v string) *ModifyAppAgentTemplateShrinkRequest {
	s.AmbientSoundConfigShrink = &v
	return s
}

func (s *ModifyAppAgentTemplateShrinkRequest) SetAppId(v string) *ModifyAppAgentTemplateShrinkRequest {
	s.AppId = &v
	return s
}

func (s *ModifyAppAgentTemplateShrinkRequest) SetAsrConfigShrink(v string) *ModifyAppAgentTemplateShrinkRequest {
	s.AsrConfigShrink = &v
	return s
}

func (s *ModifyAppAgentTemplateShrinkRequest) SetBackChannelConfigShrink(v string) *ModifyAppAgentTemplateShrinkRequest {
	s.BackChannelConfigShrink = &v
	return s
}

func (s *ModifyAppAgentTemplateShrinkRequest) SetChatMode(v int32) *ModifyAppAgentTemplateShrinkRequest {
	s.ChatMode = &v
	return s
}

func (s *ModifyAppAgentTemplateShrinkRequest) SetEnableVideoUnderstanding(v bool) *ModifyAppAgentTemplateShrinkRequest {
	s.EnableVideoUnderstanding = &v
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

func (s *ModifyAppAgentTemplateShrinkRequest) SetInterruptConfigShrink(v string) *ModifyAppAgentTemplateShrinkRequest {
	s.InterruptConfigShrink = &v
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

func (s *ModifyAppAgentTemplateShrinkRequest) SetPreferVideo(v int32) *ModifyAppAgentTemplateShrinkRequest {
	s.PreferVideo = &v
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
