// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *StartAgentRequest
	GetAppId() *string
	SetChannelId(v string) *StartAgentRequest
	GetChannelId() *string
	SetRtcConfig(v *StartAgentRequestRtcConfig) *StartAgentRequest
	GetRtcConfig() *StartAgentRequestRtcConfig
	SetTaskId(v string) *StartAgentRequest
	GetTaskId() *string
	SetTemplateId(v string) *StartAgentRequest
	GetTemplateId() *string
	SetVoiceChatConfig(v *StartAgentRequestVoiceChatConfig) *StartAgentRequest
	GetVoiceChatConfig() *StartAgentRequestVoiceChatConfig
}

type StartAgentRequest struct {
	// The application ID. You can create and query the ID in the console. Only a single ID is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// aoe****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the channel to join. Only a single ID is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// yourChannelId
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// The configuration for the agent to join a Real-Time Communication (RTC) channel.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	RtcConfig *StartAgentRequestRtcConfig `json:"RtcConfig,omitempty" xml:"RtcConfig,omitempty" type:"Struct"`
	// The task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// yourTaskId
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 76dasgb****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The configuration for AI voice chat.
	VoiceChatConfig *StartAgentRequestVoiceChatConfig `json:"VoiceChatConfig,omitempty" xml:"VoiceChatConfig,omitempty" type:"Struct"`
}

func (s StartAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s StartAgentRequest) GoString() string {
	return s.String()
}

func (s *StartAgentRequest) GetAppId() *string {
	return s.AppId
}

func (s *StartAgentRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *StartAgentRequest) GetRtcConfig() *StartAgentRequestRtcConfig {
	return s.RtcConfig
}

func (s *StartAgentRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *StartAgentRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *StartAgentRequest) GetVoiceChatConfig() *StartAgentRequestVoiceChatConfig {
	return s.VoiceChatConfig
}

func (s *StartAgentRequest) SetAppId(v string) *StartAgentRequest {
	s.AppId = &v
	return s
}

func (s *StartAgentRequest) SetChannelId(v string) *StartAgentRequest {
	s.ChannelId = &v
	return s
}

func (s *StartAgentRequest) SetRtcConfig(v *StartAgentRequestRtcConfig) *StartAgentRequest {
	s.RtcConfig = v
	return s
}

func (s *StartAgentRequest) SetTaskId(v string) *StartAgentRequest {
	s.TaskId = &v
	return s
}

func (s *StartAgentRequest) SetTemplateId(v string) *StartAgentRequest {
	s.TemplateId = &v
	return s
}

func (s *StartAgentRequest) SetVoiceChatConfig(v *StartAgentRequestVoiceChatConfig) *StartAgentRequest {
	s.VoiceChatConfig = v
	return s
}

func (s *StartAgentRequest) Validate() error {
	if s.RtcConfig != nil {
		if err := s.RtcConfig.Validate(); err != nil {
			return err
		}
	}
	if s.VoiceChatConfig != nil {
		if err := s.VoiceChatConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type StartAgentRequestRtcConfig struct {
	// The list of user IDs subscribed by the agent. Currently, only one user ID is supported.
	TargetUserIds []*string `json:"TargetUserIds,omitempty" xml:"TargetUserIds,omitempty" type:"Repeated"`
	// The user ID of the agent in the RTC channel.
	//
	// This parameter is required.
	//
	// example:
	//
	// 423341
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The maximum duration threshold for user inactivity, in seconds.
	//
	// When the duration of any of the following states reaches this threshold, the agent automatically stops and exits the channel:
	//
	// 1. All users in the targetUserIds list have left the channel.
	//
	// 2. The users in the targetUserIds list have not joined the channel.
	//
	// Default value: 30.
	//
	// Valid values: 5 to 180.
	//
	// example:
	//
	// 30
	UserInactivityTimeout *int32 `json:"UserInactivityTimeout,omitempty" xml:"UserInactivityTimeout,omitempty"`
}

func (s StartAgentRequestRtcConfig) String() string {
	return dara.Prettify(s)
}

func (s StartAgentRequestRtcConfig) GoString() string {
	return s.String()
}

func (s *StartAgentRequestRtcConfig) GetTargetUserIds() []*string {
	return s.TargetUserIds
}

func (s *StartAgentRequestRtcConfig) GetUserId() *string {
	return s.UserId
}

func (s *StartAgentRequestRtcConfig) GetUserInactivityTimeout() *int32 {
	return s.UserInactivityTimeout
}

func (s *StartAgentRequestRtcConfig) SetTargetUserIds(v []*string) *StartAgentRequestRtcConfig {
	s.TargetUserIds = v
	return s
}

func (s *StartAgentRequestRtcConfig) SetUserId(v string) *StartAgentRequestRtcConfig {
	s.UserId = &v
	return s
}

func (s *StartAgentRequestRtcConfig) SetUserInactivityTimeout(v int32) *StartAgentRequestRtcConfig {
	s.UserInactivityTimeout = &v
	return s
}

func (s *StartAgentRequestRtcConfig) Validate() error {
	return dara.Validate(s)
}

type StartAgentRequestVoiceChatConfig struct {
	// The Automatic Speech Recognition (ASR) configuration.
	ASRConfig *StartAgentRequestVoiceChatConfigASRConfig `json:"ASRConfig,omitempty" xml:"ASRConfig,omitempty" type:"Struct"`
	// The configuration for agent behavior during silence timeout. When a user joins the channel but does not interact effectively with the agent (the user does not speak), the agent is in a silent state. You can configure the agent to broadcast a voice messaging reminder or send a webhook notification during this silent state.
	AgentSilenceConfig *StartAgentRequestVoiceChatConfigAgentSilenceConfig `json:"AgentSilenceConfig,omitempty" xml:"AgentSilenceConfig,omitempty" type:"Struct"`
	// The ambient sound settings.
	AmbientSoundConfig *StartAgentRequestVoiceChatConfigAmbientSoundConfig `json:"AmbientSoundConfig,omitempty" xml:"AmbientSoundConfig,omitempty" type:"Struct"`
	// The backchannel settings.
	BackChannelConfig *StartAgentRequestVoiceChatConfigBackChannelConfig `json:"BackChannelConfig,omitempty" xml:"BackChannelConfig,omitempty" type:"Struct"`
	// The chat mode. Valid values:
	//
	// - 1: natural conversation mode. This is the default value.
	//
	// - 2: walkie-talkie mode.
	//
	// example:
	//
	// 1
	ChatMode *int32 `json:"ChatMode,omitempty" xml:"ChatMode,omitempty"`
	// Specifies whether to enable visual understanding. If enabled, the large language model (LLM) must support multimodal data.
	EnableVideoUnderstanding *bool `json:"EnableVideoUnderstanding,omitempty" xml:"EnableVideoUnderstanding,omitempty"`
	// The greeting message.
	//
	// example:
	//
	// Good morning, my friend
	Greeting *string `json:"Greeting,omitempty" xml:"Greeting,omitempty"`
	// The interrupt settings.
	InterruptConfig *StartAgentRequestVoiceChatConfigInterruptConfig `json:"InterruptConfig,omitempty" xml:"InterruptConfig,omitempty" type:"Struct"`
	// The interrupt mode. Valid values:
	//
	// - 1: voice interruption supported. This is the default value.
	//
	// - 2: voice interruption disabled.
	//
	// example:
	//
	// 1
	InterruptMode *int32 `json:"InterruptMode,omitempty" xml:"InterruptMode,omitempty"`
	// The large language model (LLM) configuration.
	LLMConfig *StartAgentRequestVoiceChatConfigLLMConfig `json:"LLMConfig,omitempty" xml:"LLMConfig,omitempty" type:"Struct"`
	// The text-to-speech (TTS) configuration.
	TTSConfig *StartAgentRequestVoiceChatConfigTTSConfig `json:"TTSConfig,omitempty" xml:"TTSConfig,omitempty" type:"Struct"`
	// After visual understanding is enabled, you can specify the video stream that the agent prioritizes for understanding. The agent can understand only one video stream when both the camera stream and the screen sharing stream are pushed simultaneously.
	//
	// 1. The agent prioritizes the camera stream.
	//
	// 2. The agent prioritizes the screen sharing stream.
	//
	// By default, the camera stream is prioritized.
	//
	// example:
	//
	// 1
	PreferVideo *int32 `json:"preferVideo,omitempty" xml:"preferVideo,omitempty"`
}

func (s StartAgentRequestVoiceChatConfig) String() string {
	return dara.Prettify(s)
}

func (s StartAgentRequestVoiceChatConfig) GoString() string {
	return s.String()
}

func (s *StartAgentRequestVoiceChatConfig) GetASRConfig() *StartAgentRequestVoiceChatConfigASRConfig {
	return s.ASRConfig
}

func (s *StartAgentRequestVoiceChatConfig) GetAgentSilenceConfig() *StartAgentRequestVoiceChatConfigAgentSilenceConfig {
	return s.AgentSilenceConfig
}

func (s *StartAgentRequestVoiceChatConfig) GetAmbientSoundConfig() *StartAgentRequestVoiceChatConfigAmbientSoundConfig {
	return s.AmbientSoundConfig
}

func (s *StartAgentRequestVoiceChatConfig) GetBackChannelConfig() *StartAgentRequestVoiceChatConfigBackChannelConfig {
	return s.BackChannelConfig
}

func (s *StartAgentRequestVoiceChatConfig) GetChatMode() *int32 {
	return s.ChatMode
}

func (s *StartAgentRequestVoiceChatConfig) GetEnableVideoUnderstanding() *bool {
	return s.EnableVideoUnderstanding
}

func (s *StartAgentRequestVoiceChatConfig) GetGreeting() *string {
	return s.Greeting
}

func (s *StartAgentRequestVoiceChatConfig) GetInterruptConfig() *StartAgentRequestVoiceChatConfigInterruptConfig {
	return s.InterruptConfig
}

func (s *StartAgentRequestVoiceChatConfig) GetInterruptMode() *int32 {
	return s.InterruptMode
}

func (s *StartAgentRequestVoiceChatConfig) GetLLMConfig() *StartAgentRequestVoiceChatConfigLLMConfig {
	return s.LLMConfig
}

func (s *StartAgentRequestVoiceChatConfig) GetTTSConfig() *StartAgentRequestVoiceChatConfigTTSConfig {
	return s.TTSConfig
}

func (s *StartAgentRequestVoiceChatConfig) GetPreferVideo() *int32 {
	return s.PreferVideo
}

func (s *StartAgentRequestVoiceChatConfig) SetASRConfig(v *StartAgentRequestVoiceChatConfigASRConfig) *StartAgentRequestVoiceChatConfig {
	s.ASRConfig = v
	return s
}

func (s *StartAgentRequestVoiceChatConfig) SetAgentSilenceConfig(v *StartAgentRequestVoiceChatConfigAgentSilenceConfig) *StartAgentRequestVoiceChatConfig {
	s.AgentSilenceConfig = v
	return s
}

func (s *StartAgentRequestVoiceChatConfig) SetAmbientSoundConfig(v *StartAgentRequestVoiceChatConfigAmbientSoundConfig) *StartAgentRequestVoiceChatConfig {
	s.AmbientSoundConfig = v
	return s
}

func (s *StartAgentRequestVoiceChatConfig) SetBackChannelConfig(v *StartAgentRequestVoiceChatConfigBackChannelConfig) *StartAgentRequestVoiceChatConfig {
	s.BackChannelConfig = v
	return s
}

func (s *StartAgentRequestVoiceChatConfig) SetChatMode(v int32) *StartAgentRequestVoiceChatConfig {
	s.ChatMode = &v
	return s
}

func (s *StartAgentRequestVoiceChatConfig) SetEnableVideoUnderstanding(v bool) *StartAgentRequestVoiceChatConfig {
	s.EnableVideoUnderstanding = &v
	return s
}

func (s *StartAgentRequestVoiceChatConfig) SetGreeting(v string) *StartAgentRequestVoiceChatConfig {
	s.Greeting = &v
	return s
}

func (s *StartAgentRequestVoiceChatConfig) SetInterruptConfig(v *StartAgentRequestVoiceChatConfigInterruptConfig) *StartAgentRequestVoiceChatConfig {
	s.InterruptConfig = v
	return s
}

func (s *StartAgentRequestVoiceChatConfig) SetInterruptMode(v int32) *StartAgentRequestVoiceChatConfig {
	s.InterruptMode = &v
	return s
}

func (s *StartAgentRequestVoiceChatConfig) SetLLMConfig(v *StartAgentRequestVoiceChatConfigLLMConfig) *StartAgentRequestVoiceChatConfig {
	s.LLMConfig = v
	return s
}

func (s *StartAgentRequestVoiceChatConfig) SetTTSConfig(v *StartAgentRequestVoiceChatConfigTTSConfig) *StartAgentRequestVoiceChatConfig {
	s.TTSConfig = v
	return s
}

func (s *StartAgentRequestVoiceChatConfig) SetPreferVideo(v int32) *StartAgentRequestVoiceChatConfig {
	s.PreferVideo = &v
	return s
}

func (s *StartAgentRequestVoiceChatConfig) Validate() error {
	if s.ASRConfig != nil {
		if err := s.ASRConfig.Validate(); err != nil {
			return err
		}
	}
	if s.AgentSilenceConfig != nil {
		if err := s.AgentSilenceConfig.Validate(); err != nil {
			return err
		}
	}
	if s.AmbientSoundConfig != nil {
		if err := s.AmbientSoundConfig.Validate(); err != nil {
			return err
		}
	}
	if s.BackChannelConfig != nil {
		if err := s.BackChannelConfig.Validate(); err != nil {
			return err
		}
	}
	if s.InterruptConfig != nil {
		if err := s.InterruptConfig.Validate(); err != nil {
			return err
		}
	}
	if s.LLMConfig != nil {
		if err := s.LLMConfig.Validate(); err != nil {
			return err
		}
	}
	if s.TTSConfig != nil {
		if err := s.TTSConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type StartAgentRequestVoiceChatConfigASRConfig struct {
	// If this parameter is not specified, the model automatically identifies the language.
	//
	// Currently supported language codes:
	//
	// - zh: Chinese, including Mandarin and dialects such as Shanghainese, Wu, Minnan, Northeastern, Gansu, Guizhou, Henan, Hubei, Hunan, Jiangxi, Ningxia, Shanxi, Shaanxi, Shandong, Sichuan, Tianjin, Yunnan, and Cantonese.
	//
	// - en: English.
	//
	// - ja: Japanese.
	//
	// - yue: Cantonese.
	//
	// - ko: Korean.
	LanguageHints []*string `json:"LanguageHints,omitempty" xml:"LanguageHints,omitempty" type:"Repeated"`
	// The voice segmentation detection threshold. Default value: 800. Unit: milliseconds. Recommended range: 200 to 6000.
	//
	// example:
	//
	// 800
	MaxSentenceSilence *int32 `json:"MaxSentenceSilence,omitempty" xml:"MaxSentenceSilence,omitempty"`
	// Specifies whether semantic or voice segmentation is enabled.
	//
	// example:
	//
	// false
	SemanticPunctuationEnabled *bool `json:"SemanticPunctuationEnabled,omitempty" xml:"SemanticPunctuationEnabled,omitempty"`
	// No configuration is required.
	//
	// example:
	//
	// zh
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	// The Voice Activity Detection (VAD) configuration.
	VadConfig *StartAgentRequestVoiceChatConfigASRConfigVadConfig `json:"VadConfig,omitempty" xml:"VadConfig,omitempty" type:"Struct"`
	// The ID of the supported hot word.
	//
	// example:
	//
	// vocab-xxx-24ee19fa8cfb4d52902170a0xxxxxxxx
	VocabularyId *string `json:"VocabularyId,omitempty" xml:"VocabularyId,omitempty"`
}

func (s StartAgentRequestVoiceChatConfigASRConfig) String() string {
	return dara.Prettify(s)
}

func (s StartAgentRequestVoiceChatConfigASRConfig) GoString() string {
	return s.String()
}

func (s *StartAgentRequestVoiceChatConfigASRConfig) GetLanguageHints() []*string {
	return s.LanguageHints
}

func (s *StartAgentRequestVoiceChatConfigASRConfig) GetMaxSentenceSilence() *int32 {
	return s.MaxSentenceSilence
}

func (s *StartAgentRequestVoiceChatConfigASRConfig) GetSemanticPunctuationEnabled() *bool {
	return s.SemanticPunctuationEnabled
}

func (s *StartAgentRequestVoiceChatConfigASRConfig) GetSourceLanguage() *string {
	return s.SourceLanguage
}

func (s *StartAgentRequestVoiceChatConfigASRConfig) GetVadConfig() *StartAgentRequestVoiceChatConfigASRConfigVadConfig {
	return s.VadConfig
}

func (s *StartAgentRequestVoiceChatConfigASRConfig) GetVocabularyId() *string {
	return s.VocabularyId
}

func (s *StartAgentRequestVoiceChatConfigASRConfig) SetLanguageHints(v []*string) *StartAgentRequestVoiceChatConfigASRConfig {
	s.LanguageHints = v
	return s
}

func (s *StartAgentRequestVoiceChatConfigASRConfig) SetMaxSentenceSilence(v int32) *StartAgentRequestVoiceChatConfigASRConfig {
	s.MaxSentenceSilence = &v
	return s
}

func (s *StartAgentRequestVoiceChatConfigASRConfig) SetSemanticPunctuationEnabled(v bool) *StartAgentRequestVoiceChatConfigASRConfig {
	s.SemanticPunctuationEnabled = &v
	return s
}

func (s *StartAgentRequestVoiceChatConfigASRConfig) SetSourceLanguage(v string) *StartAgentRequestVoiceChatConfigASRConfig {
	s.SourceLanguage = &v
	return s
}

func (s *StartAgentRequestVoiceChatConfigASRConfig) SetVadConfig(v *StartAgentRequestVoiceChatConfigASRConfigVadConfig) *StartAgentRequestVoiceChatConfigASRConfig {
	s.VadConfig = v
	return s
}

func (s *StartAgentRequestVoiceChatConfigASRConfig) SetVocabularyId(v string) *StartAgentRequestVoiceChatConfigASRConfig {
	s.VocabularyId = &v
	return s
}

func (s *StartAgentRequestVoiceChatConfigASRConfig) Validate() error {
	if s.VadConfig != nil {
		if err := s.VadConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type StartAgentRequestVoiceChatConfigASRConfigVadConfig struct {
	// The minimum duration of continuous human voice signals required to avoid false interruptions, in milliseconds.
	//
	// Valid values: 0 to 3000.
	//
	// Default value: 300.
	//
	// example:
	//
	// 300
	InterruptSpeechDuration *int32 `json:"InterruptSpeechDuration,omitempty" xml:"InterruptSpeechDuration,omitempty"`
}

func (s StartAgentRequestVoiceChatConfigASRConfigVadConfig) String() string {
	return dara.Prettify(s)
}

func (s StartAgentRequestVoiceChatConfigASRConfigVadConfig) GoString() string {
	return s.String()
}

func (s *StartAgentRequestVoiceChatConfigASRConfigVadConfig) GetInterruptSpeechDuration() *int32 {
	return s.InterruptSpeechDuration
}

func (s *StartAgentRequestVoiceChatConfigASRConfigVadConfig) SetInterruptSpeechDuration(v int32) *StartAgentRequestVoiceChatConfigASRConfigVadConfig {
	s.InterruptSpeechDuration = &v
	return s
}

func (s *StartAgentRequestVoiceChatConfigASRConfigVadConfig) Validate() error {
	return dara.Validate(s)
}

type StartAgentRequestVoiceChatConfigAgentSilenceConfig struct {
	// The timeout period for silence reminders, in seconds.
	//
	// This reminder can be triggered multiple times. If the agent remains silent after a silence reminder timeout, the reminder is triggered again upon the next timeout.
	//
	// Valid values: 5 to 600.
	//
	// Default value: 0. A value of 0 indicates that no reminder is sent.
	//
	// example:
	//
	// 0
	AlertTimeout *int32 `json:"AlertTimeout,omitempty" xml:"AlertTimeout,omitempty"`
	// The content of the silence prompt message.
	//
	// example:
	//
	// **
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The switch.
	//
	// example:
	//
	// false
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The script strategy for silence reminders. Valid values:
	//
	// - 1: Broadcast the prompt message specified by Content to the user by voice.
	//
	// - 2: Send the prompt message specified by Content and the conversation history to the LLM, and broadcast the generated reply by voice.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	Strategy *int32 `json:"Strategy,omitempty" xml:"Strategy,omitempty"`
	// The silence timeout period for triggering the agent silence timeout webhook, in seconds.
	//
	// This webhook can be triggered multiple times. If the agent remains silent after a webhook silence timeout, the webhook is triggered again upon the next timeout.
	//
	// Valid values: 10 to 600.
	//
	// Default value: 0. A value of 0 indicates that the webhook is not triggered.
	//
	// example:
	//
	// 0
	WebhookTriggerTimeout *int32 `json:"WebhookTriggerTimeout,omitempty" xml:"WebhookTriggerTimeout,omitempty"`
}

func (s StartAgentRequestVoiceChatConfigAgentSilenceConfig) String() string {
	return dara.Prettify(s)
}

func (s StartAgentRequestVoiceChatConfigAgentSilenceConfig) GoString() string {
	return s.String()
}

func (s *StartAgentRequestVoiceChatConfigAgentSilenceConfig) GetAlertTimeout() *int32 {
	return s.AlertTimeout
}

func (s *StartAgentRequestVoiceChatConfigAgentSilenceConfig) GetContent() *string {
	return s.Content
}

func (s *StartAgentRequestVoiceChatConfigAgentSilenceConfig) GetEnable() *bool {
	return s.Enable
}

func (s *StartAgentRequestVoiceChatConfigAgentSilenceConfig) GetStrategy() *int32 {
	return s.Strategy
}

func (s *StartAgentRequestVoiceChatConfigAgentSilenceConfig) GetWebhookTriggerTimeout() *int32 {
	return s.WebhookTriggerTimeout
}

func (s *StartAgentRequestVoiceChatConfigAgentSilenceConfig) SetAlertTimeout(v int32) *StartAgentRequestVoiceChatConfigAgentSilenceConfig {
	s.AlertTimeout = &v
	return s
}

func (s *StartAgentRequestVoiceChatConfigAgentSilenceConfig) SetContent(v string) *StartAgentRequestVoiceChatConfigAgentSilenceConfig {
	s.Content = &v
	return s
}

func (s *StartAgentRequestVoiceChatConfigAgentSilenceConfig) SetEnable(v bool) *StartAgentRequestVoiceChatConfigAgentSilenceConfig {
	s.Enable = &v
	return s
}

func (s *StartAgentRequestVoiceChatConfigAgentSilenceConfig) SetStrategy(v int32) *StartAgentRequestVoiceChatConfigAgentSilenceConfig {
	s.Strategy = &v
	return s
}

func (s *StartAgentRequestVoiceChatConfigAgentSilenceConfig) SetWebhookTriggerTimeout(v int32) *StartAgentRequestVoiceChatConfigAgentSilenceConfig {
	s.WebhookTriggerTimeout = &v
	return s
}

func (s *StartAgentRequestVoiceChatConfigAgentSilenceConfig) Validate() error {
	return dara.Validate(s)
}

type StartAgentRequestVoiceChatConfigAmbientSoundConfig struct {
	// The ambient sound ID. Built-in and download modes are planned. Currently, only the built-in mode is supported.
	//
	// - If this parameter is not specified or is left empty, ambient sound is disabled.
	//
	// - Built-in mode: The agent provides the following default ambient sounds: white_noise (white noise) and office (office environment).
	//
	// - Download mode: You can specify a custom soundId in the console (which must be different from built-in IDs) and provide a URL. The agent will download and bind the sound internally.
	//
	// example:
	//
	// white_noise
	SoundId *string `json:"SoundId,omitempty" xml:"SoundId,omitempty"`
	// The volume of the ambient sound.
	//
	// Default value: 50 (original volume).
	//
	// Valid values: 0 to 100.
	//
	// example:
	//
	// 50
	Volume *int32 `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s StartAgentRequestVoiceChatConfigAmbientSoundConfig) String() string {
	return dara.Prettify(s)
}

func (s StartAgentRequestVoiceChatConfigAmbientSoundConfig) GoString() string {
	return s.String()
}

func (s *StartAgentRequestVoiceChatConfigAmbientSoundConfig) GetSoundId() *string {
	return s.SoundId
}

func (s *StartAgentRequestVoiceChatConfigAmbientSoundConfig) GetVolume() *int32 {
	return s.Volume
}

func (s *StartAgentRequestVoiceChatConfigAmbientSoundConfig) SetSoundId(v string) *StartAgentRequestVoiceChatConfigAmbientSoundConfig {
	s.SoundId = &v
	return s
}

func (s *StartAgentRequestVoiceChatConfigAmbientSoundConfig) SetVolume(v int32) *StartAgentRequestVoiceChatConfigAmbientSoundConfig {
	s.Volume = &v
	return s
}

func (s *StartAgentRequestVoiceChatConfigAmbientSoundConfig) Validate() error {
	return dara.Validate(s)
}

type StartAgentRequestVoiceChatConfigBackChannelConfig struct {
	// Specifies whether to enable backchannel responses after the user finishes speaking. Default value: false.
	//
	// example:
	//
	// false
	UserTurnEnd *bool `json:"UserTurnEnd,omitempty" xml:"UserTurnEnd,omitempty"`
	// example:
	//
	// 1
	Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s StartAgentRequestVoiceChatConfigBackChannelConfig) String() string {
	return dara.Prettify(s)
}

func (s StartAgentRequestVoiceChatConfigBackChannelConfig) GoString() string {
	return s.String()
}

func (s *StartAgentRequestVoiceChatConfigBackChannelConfig) GetUserTurnEnd() *bool {
	return s.UserTurnEnd
}

func (s *StartAgentRequestVoiceChatConfigBackChannelConfig) GetVersion() *int32 {
	return s.Version
}

func (s *StartAgentRequestVoiceChatConfigBackChannelConfig) SetUserTurnEnd(v bool) *StartAgentRequestVoiceChatConfigBackChannelConfig {
	s.UserTurnEnd = &v
	return s
}

func (s *StartAgentRequestVoiceChatConfigBackChannelConfig) SetVersion(v int32) *StartAgentRequestVoiceChatConfigBackChannelConfig {
	s.Version = &v
	return s
}

func (s *StartAgentRequestVoiceChatConfigBackChannelConfig) Validate() error {
	return dara.Validate(s)
}

type StartAgentRequestVoiceChatConfigInterruptConfig struct {
	// Specifies whether to enable semantic interruption. Default value: true.
	//
	// example:
	//
	// true
	SemanticsInterrupt *bool `json:"SemanticsInterrupt,omitempty" xml:"SemanticsInterrupt,omitempty"`
}

func (s StartAgentRequestVoiceChatConfigInterruptConfig) String() string {
	return dara.Prettify(s)
}

func (s StartAgentRequestVoiceChatConfigInterruptConfig) GoString() string {
	return s.String()
}

func (s *StartAgentRequestVoiceChatConfigInterruptConfig) GetSemanticsInterrupt() *bool {
	return s.SemanticsInterrupt
}

func (s *StartAgentRequestVoiceChatConfigInterruptConfig) SetSemanticsInterrupt(v bool) *StartAgentRequestVoiceChatConfigInterruptConfig {
	s.SemanticsInterrupt = &v
	return s
}

func (s *StartAgentRequestVoiceChatConfigInterruptConfig) Validate() error {
	return dara.Validate(s)
}

type StartAgentRequestVoiceChatConfigLLMConfig struct {
	// The API key of the service provider.
	//
	// example:
	//
	// xxxxxxxxxxx
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// The agent application ID. You must specify either Model or AppId. If both are specified, AppId takes precedence.
	//
	// example:
	//
	// e****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The custom parameters.
	BizParams interface{} `json:"BizParams,omitempty" xml:"BizParams,omitempty"`
	// The number of historical conversation turns provided to the LLM.
	//
	// Default value: 3.
	//
	// example:
	//
	// 3
	HistoryDepth *int32 `json:"HistoryDepth,omitempty" xml:"HistoryDepth,omitempty"`
	// The maximum number of tokens for the output text.
	//
	// Default value: 8192.
	//
	// The maximum value for DeepSeek and Qwen models is 8192.
	//
	// example:
	//
	// 500
	MaxToken *int32 `json:"MaxToken,omitempty" xml:"MaxToken,omitempty"`
	// The model name.
	//
	// Supported models under the aliyun provider:
	//
	// ["qwen-plus", "deepseek-r1", "qwen-omni-turbo"]
	//
	// example:
	//
	// qwen-plus
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// The custom user parameters in key-value pairs. This parameter is deprecated. Use the BizParams parameter instead.
	//
	// example:
	//
	// {\"verification_code\": \"xxx\"}
	Params map[string]interface{} `json:"Params,omitempty" xml:"Params,omitempty"`
	// The prompt used to instruct the model on the role or behavior to adopt.
	//
	// example:
	//
	// xxxx
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// The sampling temperature used to control the randomness and creativity of the generated text. A higher value indicates higher randomness.
	//
	// Default value: 0.7.
	//
	// Valid values: 0 to 2, excluding 2.
	//
	// Note: This parameter does not take effect for deepseek-r1.
	//
	// example:
	//
	// 0.7
	Temperature         *float64      `json:"Temperature,omitempty" xml:"Temperature,omitempty"`
	ToolExecutionConfig interface{}   `json:"ToolExecutionConfig,omitempty" xml:"ToolExecutionConfig,omitempty"`
	Tools               []interface{} `json:"Tools,omitempty" xml:"Tools,omitempty" type:"Repeated"`
	// The sampling range used to control the diversity of output tokens. The model samples from tokens whose cumulative probability exceeds this value to ensure the sampling range is not too broad. A higher value indicates a richer variety of output tokens.
	//
	// Valid values: 0 to 1, excluding 0.
	//
	// Default value: 0.8.
	//
	// example:
	//
	// 0.8
	TopP *float64 `json:"TopP,omitempty" xml:"TopP,omitempty"`
	// The endpoint URL for integrating a third-party LLM.
	//
	// example:
	//
	// https://xxxxx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// The name of the service provider when a cloud service provider is integrated. Examples: aliyun for Model Studio, thirdparty.
	//
	// example:
	//
	// aliyun
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s StartAgentRequestVoiceChatConfigLLMConfig) String() string {
	return dara.Prettify(s)
}

func (s StartAgentRequestVoiceChatConfigLLMConfig) GoString() string {
	return s.String()
}

func (s *StartAgentRequestVoiceChatConfigLLMConfig) GetApiKey() *string {
	return s.ApiKey
}

func (s *StartAgentRequestVoiceChatConfigLLMConfig) GetAppId() *string {
	return s.AppId
}

func (s *StartAgentRequestVoiceChatConfigLLMConfig) GetBizParams() interface{} {
	return s.BizParams
}

func (s *StartAgentRequestVoiceChatConfigLLMConfig) GetHistoryDepth() *int32 {
	return s.HistoryDepth
}

func (s *StartAgentRequestVoiceChatConfigLLMConfig) GetMaxToken() *int32 {
	return s.MaxToken
}

func (s *StartAgentRequestVoiceChatConfigLLMConfig) GetModel() *string {
	return s.Model
}

func (s *StartAgentRequestVoiceChatConfigLLMConfig) GetParams() map[string]interface{} {
	return s.Params
}

func (s *StartAgentRequestVoiceChatConfigLLMConfig) GetPrompt() *string {
	return s.Prompt
}

func (s *StartAgentRequestVoiceChatConfigLLMConfig) GetTemperature() *float64 {
	return s.Temperature
}

func (s *StartAgentRequestVoiceChatConfigLLMConfig) GetToolExecutionConfig() interface{} {
	return s.ToolExecutionConfig
}

func (s *StartAgentRequestVoiceChatConfigLLMConfig) GetTools() []interface{} {
	return s.Tools
}

func (s *StartAgentRequestVoiceChatConfigLLMConfig) GetTopP() *float64 {
	return s.TopP
}

func (s *StartAgentRequestVoiceChatConfigLLMConfig) GetUrl() *string {
	return s.Url
}

func (s *StartAgentRequestVoiceChatConfigLLMConfig) GetVendor() *string {
	return s.Vendor
}

func (s *StartAgentRequestVoiceChatConfigLLMConfig) SetApiKey(v string) *StartAgentRequestVoiceChatConfigLLMConfig {
	s.ApiKey = &v
	return s
}

func (s *StartAgentRequestVoiceChatConfigLLMConfig) SetAppId(v string) *StartAgentRequestVoiceChatConfigLLMConfig {
	s.AppId = &v
	return s
}

func (s *StartAgentRequestVoiceChatConfigLLMConfig) SetBizParams(v interface{}) *StartAgentRequestVoiceChatConfigLLMConfig {
	s.BizParams = v
	return s
}

func (s *StartAgentRequestVoiceChatConfigLLMConfig) SetHistoryDepth(v int32) *StartAgentRequestVoiceChatConfigLLMConfig {
	s.HistoryDepth = &v
	return s
}

func (s *StartAgentRequestVoiceChatConfigLLMConfig) SetMaxToken(v int32) *StartAgentRequestVoiceChatConfigLLMConfig {
	s.MaxToken = &v
	return s
}

func (s *StartAgentRequestVoiceChatConfigLLMConfig) SetModel(v string) *StartAgentRequestVoiceChatConfigLLMConfig {
	s.Model = &v
	return s
}

func (s *StartAgentRequestVoiceChatConfigLLMConfig) SetParams(v map[string]interface{}) *StartAgentRequestVoiceChatConfigLLMConfig {
	s.Params = v
	return s
}

func (s *StartAgentRequestVoiceChatConfigLLMConfig) SetPrompt(v string) *StartAgentRequestVoiceChatConfigLLMConfig {
	s.Prompt = &v
	return s
}

func (s *StartAgentRequestVoiceChatConfigLLMConfig) SetTemperature(v float64) *StartAgentRequestVoiceChatConfigLLMConfig {
	s.Temperature = &v
	return s
}

func (s *StartAgentRequestVoiceChatConfigLLMConfig) SetToolExecutionConfig(v interface{}) *StartAgentRequestVoiceChatConfigLLMConfig {
	s.ToolExecutionConfig = v
	return s
}

func (s *StartAgentRequestVoiceChatConfigLLMConfig) SetTools(v []interface{}) *StartAgentRequestVoiceChatConfigLLMConfig {
	s.Tools = v
	return s
}

func (s *StartAgentRequestVoiceChatConfigLLMConfig) SetTopP(v float64) *StartAgentRequestVoiceChatConfigLLMConfig {
	s.TopP = &v
	return s
}

func (s *StartAgentRequestVoiceChatConfigLLMConfig) SetUrl(v string) *StartAgentRequestVoiceChatConfigLLMConfig {
	s.Url = &v
	return s
}

func (s *StartAgentRequestVoiceChatConfigLLMConfig) SetVendor(v string) *StartAgentRequestVoiceChatConfigLLMConfig {
	s.Vendor = &v
	return s
}

func (s *StartAgentRequestVoiceChatConfigLLMConfig) Validate() error {
	return dara.Validate(s)
}

type StartAgentRequestVoiceChatConfigTTSConfig struct {
	// The API key of the service provider.
	//
	// example:
	//
	// xxxxxx
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// The configuration to filter out characters within specified brackets from text-to-speech conversion.
	//
	// example:
	//
	// [1,2]
	FilterBrackets []*int32 `json:"FilterBrackets,omitempty" xml:"FilterBrackets,omitempty" type:"Repeated"`
	// The model name. Example: cosyvoice-v1.
	//
	// example:
	//
	// cosyvoice-v1
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// The pitch.
	//
	// Valid values: 0.5 to 2.0.
	//
	// Default value: 1.0.
	//
	// example:
	//
	// 1
	Pitch *float64 `json:"Pitch,omitempty" xml:"Pitch,omitempty"`
	// The speech rate.
	//
	// Valid values: 0.5 to 2.0.
	//
	// Default value: 1.0.
	//
	// example:
	//
	// 1
	Rate *float64 `json:"Rate,omitempty" xml:"Rate,omitempty"`
	// The service provider. Example: aliyun for Model Studio.
	//
	// example:
	//
	// aliyun
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
	// The voice. If you select the Alibaba Cloud cosyvoice-v1 model, the default voice is longxiaoxia.
	//
	// Available voices:
	//
	// longwan, longcheng, longhua, longxiaochun, longxiaoxia, longxiaocheng, longxiaobai, longlaotie, longshu, longshuo, longjing, longmiao, longyue, longyuan, longfei, longjielidou, longtong, longxiang, loongstella, and loongbella.
	//
	// example:
	//
	// longxiaoxia
	Voice *string `json:"Voice,omitempty" xml:"Voice,omitempty"`
	// The volume.
	//
	// Valid values: 0 to 100.
	//
	// Default value: 50.
	//
	// example:
	//
	// 50
	Volume *int32 `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s StartAgentRequestVoiceChatConfigTTSConfig) String() string {
	return dara.Prettify(s)
}

func (s StartAgentRequestVoiceChatConfigTTSConfig) GoString() string {
	return s.String()
}

func (s *StartAgentRequestVoiceChatConfigTTSConfig) GetApiKey() *string {
	return s.ApiKey
}

func (s *StartAgentRequestVoiceChatConfigTTSConfig) GetFilterBrackets() []*int32 {
	return s.FilterBrackets
}

func (s *StartAgentRequestVoiceChatConfigTTSConfig) GetModel() *string {
	return s.Model
}

func (s *StartAgentRequestVoiceChatConfigTTSConfig) GetPitch() *float64 {
	return s.Pitch
}

func (s *StartAgentRequestVoiceChatConfigTTSConfig) GetRate() *float64 {
	return s.Rate
}

func (s *StartAgentRequestVoiceChatConfigTTSConfig) GetVendor() *string {
	return s.Vendor
}

func (s *StartAgentRequestVoiceChatConfigTTSConfig) GetVoice() *string {
	return s.Voice
}

func (s *StartAgentRequestVoiceChatConfigTTSConfig) GetVolume() *int32 {
	return s.Volume
}

func (s *StartAgentRequestVoiceChatConfigTTSConfig) SetApiKey(v string) *StartAgentRequestVoiceChatConfigTTSConfig {
	s.ApiKey = &v
	return s
}

func (s *StartAgentRequestVoiceChatConfigTTSConfig) SetFilterBrackets(v []*int32) *StartAgentRequestVoiceChatConfigTTSConfig {
	s.FilterBrackets = v
	return s
}

func (s *StartAgentRequestVoiceChatConfigTTSConfig) SetModel(v string) *StartAgentRequestVoiceChatConfigTTSConfig {
	s.Model = &v
	return s
}

func (s *StartAgentRequestVoiceChatConfigTTSConfig) SetPitch(v float64) *StartAgentRequestVoiceChatConfigTTSConfig {
	s.Pitch = &v
	return s
}

func (s *StartAgentRequestVoiceChatConfigTTSConfig) SetRate(v float64) *StartAgentRequestVoiceChatConfigTTSConfig {
	s.Rate = &v
	return s
}

func (s *StartAgentRequestVoiceChatConfigTTSConfig) SetVendor(v string) *StartAgentRequestVoiceChatConfigTTSConfig {
	s.Vendor = &v
	return s
}

func (s *StartAgentRequestVoiceChatConfigTTSConfig) SetVoice(v string) *StartAgentRequestVoiceChatConfigTTSConfig {
	s.Voice = &v
	return s
}

func (s *StartAgentRequestVoiceChatConfigTTSConfig) SetVolume(v int32) *StartAgentRequestVoiceChatConfigTTSConfig {
	s.Volume = &v
	return s
}

func (s *StartAgentRequestVoiceChatConfigTTSConfig) Validate() error {
	return dara.Validate(s)
}
