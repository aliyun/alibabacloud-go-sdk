// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppAgentTemplateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentSilenceConfigShrink(v string) *CreateAppAgentTemplateShrinkRequest
	GetAgentSilenceConfigShrink() *string
	SetAmbientSoundConfigShrink(v string) *CreateAppAgentTemplateShrinkRequest
	GetAmbientSoundConfigShrink() *string
	SetAppId(v string) *CreateAppAgentTemplateShrinkRequest
	GetAppId() *string
	SetAsrConfigShrink(v string) *CreateAppAgentTemplateShrinkRequest
	GetAsrConfigShrink() *string
	SetBackChannelConfigShrink(v string) *CreateAppAgentTemplateShrinkRequest
	GetBackChannelConfigShrink() *string
	SetChatMode(v int32) *CreateAppAgentTemplateShrinkRequest
	GetChatMode() *int32
	SetEnableVideoUnderstanding(v bool) *CreateAppAgentTemplateShrinkRequest
	GetEnableVideoUnderstanding() *bool
	SetGreeting(v string) *CreateAppAgentTemplateShrinkRequest
	GetGreeting() *string
	SetInterruptConfigShrink(v string) *CreateAppAgentTemplateShrinkRequest
	GetInterruptConfigShrink() *string
	SetInterruptMode(v int32) *CreateAppAgentTemplateShrinkRequest
	GetInterruptMode() *int32
	SetLlmConfigShrink(v string) *CreateAppAgentTemplateShrinkRequest
	GetLlmConfigShrink() *string
	SetName(v string) *CreateAppAgentTemplateShrinkRequest
	GetName() *string
	SetPreferVideo(v int32) *CreateAppAgentTemplateShrinkRequest
	GetPreferVideo() *int32
	SetTtsConfigShrink(v string) *CreateAppAgentTemplateShrinkRequest
	GetTtsConfigShrink() *string
	SetType(v int32) *CreateAppAgentTemplateShrinkRequest
	GetType() *int32
}

type CreateAppAgentTemplateShrinkRequest struct {
	// The agent silence configuration.
	AgentSilenceConfigShrink *string `json:"AgentSilenceConfig,omitempty" xml:"AgentSilenceConfig,omitempty"`
	// The ambient environment configuration for the agent. If this parameter is not specified, the ambient environment configuration is not used.
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
	// Specifies whether to enable video understanding. If enabled, the LLM must support multimodal data.
	EnableVideoUnderstanding *bool `json:"EnableVideoUnderstanding,omitempty" xml:"EnableVideoUnderstanding,omitempty"`
	// The greeting message. The maximum length is 500 characters.
	//
	// example:
	//
	// Good morning, my friend!
	Greeting *string `json:"Greeting,omitempty" xml:"Greeting,omitempty"`
	// The agent interruption configuration.
	InterruptConfigShrink *string `json:"InterruptConfig,omitempty" xml:"InterruptConfig,omitempty"`
	// The interruption mode. Valid values:
	//
	// - 1: Interruption allowed (default).
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
	// The preferred video stream for the agent to process when visual understanding is enabled. The agent can process only one video stream at a time. When both the camera stream and the screen sharing stream are pushed simultaneously, this parameter specifies the preferred stream. Valid values:
	//
	// - 1: The agent preferentially processes the camera stream.
	//
	// - 2: The agent preferentially processes the screen sharing stream.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
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

func (s CreateAppAgentTemplateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAppAgentTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateAppAgentTemplateShrinkRequest) GetAgentSilenceConfigShrink() *string {
	return s.AgentSilenceConfigShrink
}

func (s *CreateAppAgentTemplateShrinkRequest) GetAmbientSoundConfigShrink() *string {
	return s.AmbientSoundConfigShrink
}

func (s *CreateAppAgentTemplateShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateAppAgentTemplateShrinkRequest) GetAsrConfigShrink() *string {
	return s.AsrConfigShrink
}

func (s *CreateAppAgentTemplateShrinkRequest) GetBackChannelConfigShrink() *string {
	return s.BackChannelConfigShrink
}

func (s *CreateAppAgentTemplateShrinkRequest) GetChatMode() *int32 {
	return s.ChatMode
}

func (s *CreateAppAgentTemplateShrinkRequest) GetEnableVideoUnderstanding() *bool {
	return s.EnableVideoUnderstanding
}

func (s *CreateAppAgentTemplateShrinkRequest) GetGreeting() *string {
	return s.Greeting
}

func (s *CreateAppAgentTemplateShrinkRequest) GetInterruptConfigShrink() *string {
	return s.InterruptConfigShrink
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

func (s *CreateAppAgentTemplateShrinkRequest) GetPreferVideo() *int32 {
	return s.PreferVideo
}

func (s *CreateAppAgentTemplateShrinkRequest) GetTtsConfigShrink() *string {
	return s.TtsConfigShrink
}

func (s *CreateAppAgentTemplateShrinkRequest) GetType() *int32 {
	return s.Type
}

func (s *CreateAppAgentTemplateShrinkRequest) SetAgentSilenceConfigShrink(v string) *CreateAppAgentTemplateShrinkRequest {
	s.AgentSilenceConfigShrink = &v
	return s
}

func (s *CreateAppAgentTemplateShrinkRequest) SetAmbientSoundConfigShrink(v string) *CreateAppAgentTemplateShrinkRequest {
	s.AmbientSoundConfigShrink = &v
	return s
}

func (s *CreateAppAgentTemplateShrinkRequest) SetAppId(v string) *CreateAppAgentTemplateShrinkRequest {
	s.AppId = &v
	return s
}

func (s *CreateAppAgentTemplateShrinkRequest) SetAsrConfigShrink(v string) *CreateAppAgentTemplateShrinkRequest {
	s.AsrConfigShrink = &v
	return s
}

func (s *CreateAppAgentTemplateShrinkRequest) SetBackChannelConfigShrink(v string) *CreateAppAgentTemplateShrinkRequest {
	s.BackChannelConfigShrink = &v
	return s
}

func (s *CreateAppAgentTemplateShrinkRequest) SetChatMode(v int32) *CreateAppAgentTemplateShrinkRequest {
	s.ChatMode = &v
	return s
}

func (s *CreateAppAgentTemplateShrinkRequest) SetEnableVideoUnderstanding(v bool) *CreateAppAgentTemplateShrinkRequest {
	s.EnableVideoUnderstanding = &v
	return s
}

func (s *CreateAppAgentTemplateShrinkRequest) SetGreeting(v string) *CreateAppAgentTemplateShrinkRequest {
	s.Greeting = &v
	return s
}

func (s *CreateAppAgentTemplateShrinkRequest) SetInterruptConfigShrink(v string) *CreateAppAgentTemplateShrinkRequest {
	s.InterruptConfigShrink = &v
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

func (s *CreateAppAgentTemplateShrinkRequest) SetPreferVideo(v int32) *CreateAppAgentTemplateShrinkRequest {
	s.PreferVideo = &v
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
