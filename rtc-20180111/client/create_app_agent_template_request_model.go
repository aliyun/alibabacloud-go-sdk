// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppAgentTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentSilenceConfig(v *CreateAppAgentTemplateRequestAgentSilenceConfig) *CreateAppAgentTemplateRequest
	GetAgentSilenceConfig() *CreateAppAgentTemplateRequestAgentSilenceConfig
	SetAmbientSoundConfig(v *CreateAppAgentTemplateRequestAmbientSoundConfig) *CreateAppAgentTemplateRequest
	GetAmbientSoundConfig() *CreateAppAgentTemplateRequestAmbientSoundConfig
	SetAppId(v string) *CreateAppAgentTemplateRequest
	GetAppId() *string
	SetAsrConfig(v *CreateAppAgentTemplateRequestAsrConfig) *CreateAppAgentTemplateRequest
	GetAsrConfig() *CreateAppAgentTemplateRequestAsrConfig
	SetBackChannelConfig(v *CreateAppAgentTemplateRequestBackChannelConfig) *CreateAppAgentTemplateRequest
	GetBackChannelConfig() *CreateAppAgentTemplateRequestBackChannelConfig
	SetChatMode(v int32) *CreateAppAgentTemplateRequest
	GetChatMode() *int32
	SetEnableVideoUnderstanding(v bool) *CreateAppAgentTemplateRequest
	GetEnableVideoUnderstanding() *bool
	SetGreeting(v string) *CreateAppAgentTemplateRequest
	GetGreeting() *string
	SetInterruptConfig(v *CreateAppAgentTemplateRequestInterruptConfig) *CreateAppAgentTemplateRequest
	GetInterruptConfig() *CreateAppAgentTemplateRequestInterruptConfig
	SetInterruptMode(v int32) *CreateAppAgentTemplateRequest
	GetInterruptMode() *int32
	SetLlmConfig(v *CreateAppAgentTemplateRequestLlmConfig) *CreateAppAgentTemplateRequest
	GetLlmConfig() *CreateAppAgentTemplateRequestLlmConfig
	SetName(v string) *CreateAppAgentTemplateRequest
	GetName() *string
	SetPreferVideo(v int32) *CreateAppAgentTemplateRequest
	GetPreferVideo() *int32
	SetTtsConfig(v *CreateAppAgentTemplateRequestTtsConfig) *CreateAppAgentTemplateRequest
	GetTtsConfig() *CreateAppAgentTemplateRequestTtsConfig
	SetType(v int32) *CreateAppAgentTemplateRequest
	GetType() *int32
}

type CreateAppAgentTemplateRequest struct {
	// The agent silence configuration.
	AgentSilenceConfig *CreateAppAgentTemplateRequestAgentSilenceConfig `json:"AgentSilenceConfig,omitempty" xml:"AgentSilenceConfig,omitempty" type:"Struct"`
	// The ambient environment configuration for the agent. If this parameter is not specified, the ambient environment configuration is not used.
	AmbientSoundConfig *CreateAppAgentTemplateRequestAmbientSoundConfig `json:"AmbientSoundConfig,omitempty" xml:"AmbientSoundConfig,omitempty" type:"Struct"`
	// The application ID. You can create and query the ID in the console. Only a single ID is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// ac7N****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The speech-to-text settings.
	AsrConfig *CreateAppAgentTemplateRequestAsrConfig `json:"AsrConfig,omitempty" xml:"AsrConfig,omitempty" type:"Struct"`
	// The agent backchannel configuration.
	BackChannelConfig *CreateAppAgentTemplateRequestBackChannelConfig `json:"BackChannelConfig,omitempty" xml:"BackChannelConfig,omitempty" type:"Struct"`
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
	InterruptConfig *CreateAppAgentTemplateRequestInterruptConfig `json:"InterruptConfig,omitempty" xml:"InterruptConfig,omitempty" type:"Struct"`
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
	LlmConfig *CreateAppAgentTemplateRequestLlmConfig `json:"LlmConfig,omitempty" xml:"LlmConfig,omitempty" type:"Struct"`
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
	TtsConfig *CreateAppAgentTemplateRequestTtsConfig `json:"TtsConfig,omitempty" xml:"TtsConfig,omitempty" type:"Struct"`
	// The agent type. Valid values:
	//
	// - 1: Voice call (default).
	//
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateAppAgentTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAppAgentTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateAppAgentTemplateRequest) GetAgentSilenceConfig() *CreateAppAgentTemplateRequestAgentSilenceConfig {
	return s.AgentSilenceConfig
}

func (s *CreateAppAgentTemplateRequest) GetAmbientSoundConfig() *CreateAppAgentTemplateRequestAmbientSoundConfig {
	return s.AmbientSoundConfig
}

func (s *CreateAppAgentTemplateRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateAppAgentTemplateRequest) GetAsrConfig() *CreateAppAgentTemplateRequestAsrConfig {
	return s.AsrConfig
}

func (s *CreateAppAgentTemplateRequest) GetBackChannelConfig() *CreateAppAgentTemplateRequestBackChannelConfig {
	return s.BackChannelConfig
}

func (s *CreateAppAgentTemplateRequest) GetChatMode() *int32 {
	return s.ChatMode
}

func (s *CreateAppAgentTemplateRequest) GetEnableVideoUnderstanding() *bool {
	return s.EnableVideoUnderstanding
}

func (s *CreateAppAgentTemplateRequest) GetGreeting() *string {
	return s.Greeting
}

func (s *CreateAppAgentTemplateRequest) GetInterruptConfig() *CreateAppAgentTemplateRequestInterruptConfig {
	return s.InterruptConfig
}

func (s *CreateAppAgentTemplateRequest) GetInterruptMode() *int32 {
	return s.InterruptMode
}

func (s *CreateAppAgentTemplateRequest) GetLlmConfig() *CreateAppAgentTemplateRequestLlmConfig {
	return s.LlmConfig
}

func (s *CreateAppAgentTemplateRequest) GetName() *string {
	return s.Name
}

func (s *CreateAppAgentTemplateRequest) GetPreferVideo() *int32 {
	return s.PreferVideo
}

func (s *CreateAppAgentTemplateRequest) GetTtsConfig() *CreateAppAgentTemplateRequestTtsConfig {
	return s.TtsConfig
}

func (s *CreateAppAgentTemplateRequest) GetType() *int32 {
	return s.Type
}

func (s *CreateAppAgentTemplateRequest) SetAgentSilenceConfig(v *CreateAppAgentTemplateRequestAgentSilenceConfig) *CreateAppAgentTemplateRequest {
	s.AgentSilenceConfig = v
	return s
}

func (s *CreateAppAgentTemplateRequest) SetAmbientSoundConfig(v *CreateAppAgentTemplateRequestAmbientSoundConfig) *CreateAppAgentTemplateRequest {
	s.AmbientSoundConfig = v
	return s
}

func (s *CreateAppAgentTemplateRequest) SetAppId(v string) *CreateAppAgentTemplateRequest {
	s.AppId = &v
	return s
}

func (s *CreateAppAgentTemplateRequest) SetAsrConfig(v *CreateAppAgentTemplateRequestAsrConfig) *CreateAppAgentTemplateRequest {
	s.AsrConfig = v
	return s
}

func (s *CreateAppAgentTemplateRequest) SetBackChannelConfig(v *CreateAppAgentTemplateRequestBackChannelConfig) *CreateAppAgentTemplateRequest {
	s.BackChannelConfig = v
	return s
}

func (s *CreateAppAgentTemplateRequest) SetChatMode(v int32) *CreateAppAgentTemplateRequest {
	s.ChatMode = &v
	return s
}

func (s *CreateAppAgentTemplateRequest) SetEnableVideoUnderstanding(v bool) *CreateAppAgentTemplateRequest {
	s.EnableVideoUnderstanding = &v
	return s
}

func (s *CreateAppAgentTemplateRequest) SetGreeting(v string) *CreateAppAgentTemplateRequest {
	s.Greeting = &v
	return s
}

func (s *CreateAppAgentTemplateRequest) SetInterruptConfig(v *CreateAppAgentTemplateRequestInterruptConfig) *CreateAppAgentTemplateRequest {
	s.InterruptConfig = v
	return s
}

func (s *CreateAppAgentTemplateRequest) SetInterruptMode(v int32) *CreateAppAgentTemplateRequest {
	s.InterruptMode = &v
	return s
}

func (s *CreateAppAgentTemplateRequest) SetLlmConfig(v *CreateAppAgentTemplateRequestLlmConfig) *CreateAppAgentTemplateRequest {
	s.LlmConfig = v
	return s
}

func (s *CreateAppAgentTemplateRequest) SetName(v string) *CreateAppAgentTemplateRequest {
	s.Name = &v
	return s
}

func (s *CreateAppAgentTemplateRequest) SetPreferVideo(v int32) *CreateAppAgentTemplateRequest {
	s.PreferVideo = &v
	return s
}

func (s *CreateAppAgentTemplateRequest) SetTtsConfig(v *CreateAppAgentTemplateRequestTtsConfig) *CreateAppAgentTemplateRequest {
	s.TtsConfig = v
	return s
}

func (s *CreateAppAgentTemplateRequest) SetType(v int32) *CreateAppAgentTemplateRequest {
	s.Type = &v
	return s
}

func (s *CreateAppAgentTemplateRequest) Validate() error {
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
	if s.AsrConfig != nil {
		if err := s.AsrConfig.Validate(); err != nil {
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
	if s.LlmConfig != nil {
		if err := s.LlmConfig.Validate(); err != nil {
			return err
		}
	}
	if s.TtsConfig != nil {
		if err := s.TtsConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAppAgentTemplateRequestAgentSilenceConfig struct {
	// The timeout period for silence alerts, in seconds. This alert can be triggered multiple times. If the agent remains silent after the timeout period elapses, the alert is triggered again upon the next timeout. Valid values: 5 to 600. Default value: 0. A value of 0 indicates that no alert is triggered.
	//
	// example:
	//
	// 30
	AlertTimeout *int32 `json:"AlertTimeout,omitempty" xml:"AlertTimeout,omitempty"`
	// The content of the silence prompt message.
	//
	// example:
	//
	// Hello, please try again later?
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The handling strategy for silence timeout. Valid values:
	//
	// - 1: Broadcasts the prompt message specified by Content to the user via voice.
	//
	// - 2: Sends the prompt message specified by Content and the conversation history to the LLM, and broadcasts the reply generated by the model.
	//
	// Default value: 1.
	//
	// example:
	//
	// 2
	Strategy *int32 `json:"Strategy,omitempty" xml:"Strategy,omitempty"`
	// The timeout period for triggering the webhook upon agent silence, in seconds. This webhook can be triggered multiple times. If the agent remains silent after the webhook timeout period elapses, the webhook is triggered again upon the next timeout. Valid values: 10 to 600. Default value: 0. A value of 0 indicates that the webhook is not triggered.
	//
	// example:
	//
	// 30
	WebhookTriggerTimeout *int32 `json:"WebhookTriggerTimeout,omitempty" xml:"WebhookTriggerTimeout,omitempty"`
}

func (s CreateAppAgentTemplateRequestAgentSilenceConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateAppAgentTemplateRequestAgentSilenceConfig) GoString() string {
	return s.String()
}

func (s *CreateAppAgentTemplateRequestAgentSilenceConfig) GetAlertTimeout() *int32 {
	return s.AlertTimeout
}

func (s *CreateAppAgentTemplateRequestAgentSilenceConfig) GetContent() *string {
	return s.Content
}

func (s *CreateAppAgentTemplateRequestAgentSilenceConfig) GetStrategy() *int32 {
	return s.Strategy
}

func (s *CreateAppAgentTemplateRequestAgentSilenceConfig) GetWebhookTriggerTimeout() *int32 {
	return s.WebhookTriggerTimeout
}

func (s *CreateAppAgentTemplateRequestAgentSilenceConfig) SetAlertTimeout(v int32) *CreateAppAgentTemplateRequestAgentSilenceConfig {
	s.AlertTimeout = &v
	return s
}

func (s *CreateAppAgentTemplateRequestAgentSilenceConfig) SetContent(v string) *CreateAppAgentTemplateRequestAgentSilenceConfig {
	s.Content = &v
	return s
}

func (s *CreateAppAgentTemplateRequestAgentSilenceConfig) SetStrategy(v int32) *CreateAppAgentTemplateRequestAgentSilenceConfig {
	s.Strategy = &v
	return s
}

func (s *CreateAppAgentTemplateRequestAgentSilenceConfig) SetWebhookTriggerTimeout(v int32) *CreateAppAgentTemplateRequestAgentSilenceConfig {
	s.WebhookTriggerTimeout = &v
	return s
}

func (s *CreateAppAgentTemplateRequestAgentSilenceConfig) Validate() error {
	return dara.Validate(s)
}

type CreateAppAgentTemplateRequestAmbientSoundConfig struct {
	// The default ambient sounds provided by the agent. Valid values:
	//
	// - white_noise
	//
	// - office
	//
	// example:
	//
	// white_noise
	SoundId *string `json:"SoundId,omitempty" xml:"SoundId,omitempty"`
	// The volume level.
	//
	// example:
	//
	// 100
	Volume *int32 `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s CreateAppAgentTemplateRequestAmbientSoundConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateAppAgentTemplateRequestAmbientSoundConfig) GoString() string {
	return s.String()
}

func (s *CreateAppAgentTemplateRequestAmbientSoundConfig) GetSoundId() *string {
	return s.SoundId
}

func (s *CreateAppAgentTemplateRequestAmbientSoundConfig) GetVolume() *int32 {
	return s.Volume
}

func (s *CreateAppAgentTemplateRequestAmbientSoundConfig) SetSoundId(v string) *CreateAppAgentTemplateRequestAmbientSoundConfig {
	s.SoundId = &v
	return s
}

func (s *CreateAppAgentTemplateRequestAmbientSoundConfig) SetVolume(v int32) *CreateAppAgentTemplateRequestAmbientSoundConfig {
	s.Volume = &v
	return s
}

func (s *CreateAppAgentTemplateRequestAmbientSoundConfig) Validate() error {
	return dara.Validate(s)
}

type CreateAppAgentTemplateRequestAsrConfig struct {
	// The sentence silence detection threshold. Default value: 500 ms. Valid values: 200 to 6000.
	//
	// example:
	//
	// 300
	MaxSentenceSilence *int32 `json:"MaxSentenceSilence,omitempty" xml:"MaxSentenceSilence,omitempty"`
	// The name.
	//
	// This parameter is required.
	//
	// example:
	//
	// STT
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The voice activity detection (VAD) configuration.
	VadConfig *CreateAppAgentTemplateRequestAsrConfigVadConfig `json:"VadConfig,omitempty" xml:"VadConfig,omitempty" type:"Struct"`
	// The hot words. Each vocabulary can contain up to 500 hot words, and each hot word can be up to 10 Chinese characters in length. Hot words must be UTF-8 encoded and cannot contain punctuation marks or special characters. Otherwise, the hot words do not take effect.
	WordWeights []*CreateAppAgentTemplateRequestAsrConfigWordWeights `json:"WordWeights,omitempty" xml:"WordWeights,omitempty" type:"Repeated"`
}

func (s CreateAppAgentTemplateRequestAsrConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateAppAgentTemplateRequestAsrConfig) GoString() string {
	return s.String()
}

func (s *CreateAppAgentTemplateRequestAsrConfig) GetMaxSentenceSilence() *int32 {
	return s.MaxSentenceSilence
}

func (s *CreateAppAgentTemplateRequestAsrConfig) GetName() *string {
	return s.Name
}

func (s *CreateAppAgentTemplateRequestAsrConfig) GetVadConfig() *CreateAppAgentTemplateRequestAsrConfigVadConfig {
	return s.VadConfig
}

func (s *CreateAppAgentTemplateRequestAsrConfig) GetWordWeights() []*CreateAppAgentTemplateRequestAsrConfigWordWeights {
	return s.WordWeights
}

func (s *CreateAppAgentTemplateRequestAsrConfig) SetMaxSentenceSilence(v int32) *CreateAppAgentTemplateRequestAsrConfig {
	s.MaxSentenceSilence = &v
	return s
}

func (s *CreateAppAgentTemplateRequestAsrConfig) SetName(v string) *CreateAppAgentTemplateRequestAsrConfig {
	s.Name = &v
	return s
}

func (s *CreateAppAgentTemplateRequestAsrConfig) SetVadConfig(v *CreateAppAgentTemplateRequestAsrConfigVadConfig) *CreateAppAgentTemplateRequestAsrConfig {
	s.VadConfig = v
	return s
}

func (s *CreateAppAgentTemplateRequestAsrConfig) SetWordWeights(v []*CreateAppAgentTemplateRequestAsrConfigWordWeights) *CreateAppAgentTemplateRequestAsrConfig {
	s.WordWeights = v
	return s
}

func (s *CreateAppAgentTemplateRequestAsrConfig) Validate() error {
	if s.VadConfig != nil {
		if err := s.VadConfig.Validate(); err != nil {
			return err
		}
	}
	if s.WordWeights != nil {
		for _, item := range s.WordWeights {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateAppAgentTemplateRequestAsrConfigVadConfig struct {
	// The minimum duration of continuous voice signals in milliseconds to avoid false interruptions. Default value: 300. Valid values: 0 to 3000.
	//
	// example:
	//
	// 1000
	InterruptSpeechDuration *int32 `json:"InterruptSpeechDuration,omitempty" xml:"InterruptSpeechDuration,omitempty"`
}

func (s CreateAppAgentTemplateRequestAsrConfigVadConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateAppAgentTemplateRequestAsrConfigVadConfig) GoString() string {
	return s.String()
}

func (s *CreateAppAgentTemplateRequestAsrConfigVadConfig) GetInterruptSpeechDuration() *int32 {
	return s.InterruptSpeechDuration
}

func (s *CreateAppAgentTemplateRequestAsrConfigVadConfig) SetInterruptSpeechDuration(v int32) *CreateAppAgentTemplateRequestAsrConfigVadConfig {
	s.InterruptSpeechDuration = &v
	return s
}

func (s *CreateAppAgentTemplateRequestAsrConfigVadConfig) Validate() error {
	return dara.Validate(s)
}

type CreateAppAgentTemplateRequestAsrConfigWordWeights struct {
	// The language code. This parameter is optional. The large language model automatically identifies the language. Valid values:
	//
	// - zh: Chinese.
	//
	// - en: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The hot word weight. Valid values: integers from 1 to 5. Common value: 4. If the effect is not obvious, you can increase the weight. However, a large weight may cause negative effects and lead to inaccurate recognition of other words.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
	// The hot word text. Each word can be up to 10 Chinese characters or English words in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// Apple
	Word *string `json:"Word,omitempty" xml:"Word,omitempty"`
}

func (s CreateAppAgentTemplateRequestAsrConfigWordWeights) String() string {
	return dara.Prettify(s)
}

func (s CreateAppAgentTemplateRequestAsrConfigWordWeights) GoString() string {
	return s.String()
}

func (s *CreateAppAgentTemplateRequestAsrConfigWordWeights) GetLang() *string {
	return s.Lang
}

func (s *CreateAppAgentTemplateRequestAsrConfigWordWeights) GetWeight() *int32 {
	return s.Weight
}

func (s *CreateAppAgentTemplateRequestAsrConfigWordWeights) GetWord() *string {
	return s.Word
}

func (s *CreateAppAgentTemplateRequestAsrConfigWordWeights) SetLang(v string) *CreateAppAgentTemplateRequestAsrConfigWordWeights {
	s.Lang = &v
	return s
}

func (s *CreateAppAgentTemplateRequestAsrConfigWordWeights) SetWeight(v int32) *CreateAppAgentTemplateRequestAsrConfigWordWeights {
	s.Weight = &v
	return s
}

func (s *CreateAppAgentTemplateRequestAsrConfigWordWeights) SetWord(v string) *CreateAppAgentTemplateRequestAsrConfigWordWeights {
	s.Word = &v
	return s
}

func (s *CreateAppAgentTemplateRequestAsrConfigWordWeights) Validate() error {
	return dara.Validate(s)
}

type CreateAppAgentTemplateRequestBackChannelConfig struct {
	// Specifies whether to enable backchanneling.
	UserTurnEnd *bool `json:"UserTurnEnd,omitempty" xml:"UserTurnEnd,omitempty"`
}

func (s CreateAppAgentTemplateRequestBackChannelConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateAppAgentTemplateRequestBackChannelConfig) GoString() string {
	return s.String()
}

func (s *CreateAppAgentTemplateRequestBackChannelConfig) GetUserTurnEnd() *bool {
	return s.UserTurnEnd
}

func (s *CreateAppAgentTemplateRequestBackChannelConfig) SetUserTurnEnd(v bool) *CreateAppAgentTemplateRequestBackChannelConfig {
	s.UserTurnEnd = &v
	return s
}

func (s *CreateAppAgentTemplateRequestBackChannelConfig) Validate() error {
	return dara.Validate(s)
}

type CreateAppAgentTemplateRequestInterruptConfig struct {
	// The semantic interruption configuration.
	SemanticsInterrupt *bool `json:"SemanticsInterrupt,omitempty" xml:"SemanticsInterrupt,omitempty"`
}

func (s CreateAppAgentTemplateRequestInterruptConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateAppAgentTemplateRequestInterruptConfig) GoString() string {
	return s.String()
}

func (s *CreateAppAgentTemplateRequestInterruptConfig) GetSemanticsInterrupt() *bool {
	return s.SemanticsInterrupt
}

func (s *CreateAppAgentTemplateRequestInterruptConfig) SetSemanticsInterrupt(v bool) *CreateAppAgentTemplateRequestInterruptConfig {
	s.SemanticsInterrupt = &v
	return s
}

func (s *CreateAppAgentTemplateRequestInterruptConfig) Validate() error {
	return dara.Validate(s)
}

type CreateAppAgentTemplateRequestLlmConfig struct {
	// The ID of the Model Studio agent.
	AgentAppId *string `json:"AgentAppId,omitempty" xml:"AgentAppId,omitempty"`
	// The API key required by the large language model service provider.
	//
	// This parameter is required.
	//
	// example:
	//
	// ak-1213123123132123131
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// The number of historical conversation turns provided to the large language model. Maximum value: 100. Default value: 50.
	//
	// example:
	//
	// 5
	HistoryDepth *int32 `json:"HistoryDepth,omitempty" xml:"HistoryDepth,omitempty"`
	// The maximum number of tokens for the output text. Default value: 8192.
	//
	// example:
	//
	// 1024
	MaxToken *int32 `json:"MaxToken,omitempty" xml:"MaxToken,omitempty"`
	// The name.
	//
	// This parameter is required.
	//
	// example:
	//
	// llm
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The prompt for the large language model. The maximum length is 5,000 characters.
	//
	// example:
	//
	// You are a senior psychiatrist
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// The sampling temperature. Default value: 0.7. Valid values: [0, 2).
	//
	// example:
	//
	// 0.9
	Temperature *float32 `json:"Temperature,omitempty" xml:"Temperature,omitempty"`
	// The top-P sampling range. Default value: 0.8. Valid values: (0, 1].
	//
	// example:
	//
	// 0.8
	TopP *float32 `json:"TopP,omitempty" xml:"TopP,omitempty"`
	// The endpoint URL for connecting to a third-party large language model.
	//
	// example:
	//
	// https://llm.example.aliyuns.com
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// The provider. Alibaba Cloud Model Studio and third-party configurations that comply with OpenAI API specifications are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// thirdparty
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s CreateAppAgentTemplateRequestLlmConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateAppAgentTemplateRequestLlmConfig) GoString() string {
	return s.String()
}

func (s *CreateAppAgentTemplateRequestLlmConfig) GetAgentAppId() *string {
	return s.AgentAppId
}

func (s *CreateAppAgentTemplateRequestLlmConfig) GetApiKey() *string {
	return s.ApiKey
}

func (s *CreateAppAgentTemplateRequestLlmConfig) GetHistoryDepth() *int32 {
	return s.HistoryDepth
}

func (s *CreateAppAgentTemplateRequestLlmConfig) GetMaxToken() *int32 {
	return s.MaxToken
}

func (s *CreateAppAgentTemplateRequestLlmConfig) GetName() *string {
	return s.Name
}

func (s *CreateAppAgentTemplateRequestLlmConfig) GetPrompt() *string {
	return s.Prompt
}

func (s *CreateAppAgentTemplateRequestLlmConfig) GetTemperature() *float32 {
	return s.Temperature
}

func (s *CreateAppAgentTemplateRequestLlmConfig) GetTopP() *float32 {
	return s.TopP
}

func (s *CreateAppAgentTemplateRequestLlmConfig) GetUrl() *string {
	return s.Url
}

func (s *CreateAppAgentTemplateRequestLlmConfig) GetVendor() *string {
	return s.Vendor
}

func (s *CreateAppAgentTemplateRequestLlmConfig) SetAgentAppId(v string) *CreateAppAgentTemplateRequestLlmConfig {
	s.AgentAppId = &v
	return s
}

func (s *CreateAppAgentTemplateRequestLlmConfig) SetApiKey(v string) *CreateAppAgentTemplateRequestLlmConfig {
	s.ApiKey = &v
	return s
}

func (s *CreateAppAgentTemplateRequestLlmConfig) SetHistoryDepth(v int32) *CreateAppAgentTemplateRequestLlmConfig {
	s.HistoryDepth = &v
	return s
}

func (s *CreateAppAgentTemplateRequestLlmConfig) SetMaxToken(v int32) *CreateAppAgentTemplateRequestLlmConfig {
	s.MaxToken = &v
	return s
}

func (s *CreateAppAgentTemplateRequestLlmConfig) SetName(v string) *CreateAppAgentTemplateRequestLlmConfig {
	s.Name = &v
	return s
}

func (s *CreateAppAgentTemplateRequestLlmConfig) SetPrompt(v string) *CreateAppAgentTemplateRequestLlmConfig {
	s.Prompt = &v
	return s
}

func (s *CreateAppAgentTemplateRequestLlmConfig) SetTemperature(v float32) *CreateAppAgentTemplateRequestLlmConfig {
	s.Temperature = &v
	return s
}

func (s *CreateAppAgentTemplateRequestLlmConfig) SetTopP(v float32) *CreateAppAgentTemplateRequestLlmConfig {
	s.TopP = &v
	return s
}

func (s *CreateAppAgentTemplateRequestLlmConfig) SetUrl(v string) *CreateAppAgentTemplateRequestLlmConfig {
	s.Url = &v
	return s
}

func (s *CreateAppAgentTemplateRequestLlmConfig) SetVendor(v string) *CreateAppAgentTemplateRequestLlmConfig {
	s.Vendor = &v
	return s
}

func (s *CreateAppAgentTemplateRequestLlmConfig) Validate() error {
	return dara.Validate(s)
}

type CreateAppAgentTemplateRequestTtsConfig struct {
	// The API key required by the large language model service provider.
	//
	// This parameter is required.
	//
	// example:
	//
	// ak-1213123123132123131
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// The filter settings. Valid values:
	//
	// - 1: Chinese parentheses （）
	//
	// - 2: English parentheses ()
	//
	// - 3: Chinese square brackets 【】
	//
	// - 4: English square brackets []
	//
	// - 5: English curly braces {}
	FilterBrackets []*int32 `json:"FilterBrackets,omitempty" xml:"FilterBrackets,omitempty" type:"Repeated"`
	// The name.
	//
	// This parameter is required.
	//
	// example:
	//
	// Tts
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The pitch. Default value: 1.0. Valid values: 0.5 to 2.0.
	//
	// example:
	//
	// 0.8
	Pitch *float32 `json:"Pitch,omitempty" xml:"Pitch,omitempty"`
	// The speech rate. Default value: 1.0. Valid values: 0.5 to 2.0.
	//
	// example:
	//
	// 0.8
	Rate *float32 `json:"Rate,omitempty" xml:"Rate,omitempty"`
	// The provider. Alibaba Cloud Model Studio is supported.
	//
	// example:
	//
	// aliyun
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
	// The voice. Default value: longxiaoxia. Valid values:
	//
	// - longwan
	//
	// - longcheng
	//
	// - longhua
	//
	// - longxiaochun
	//
	// - longxiaoxia
	//
	// - longxiaocheng
	//
	// - longxiaobai
	//
	// - longlaotie
	//
	// - longshu
	//
	// - longshuo
	//
	// - longjing
	//
	// - longmiao
	//
	// - longyue
	//
	// - longyuan
	//
	// - longfei
	//
	// - longjielidou
	//
	// - longtong
	//
	// - longxiang
	//
	// - loongstella
	//
	// - loongbella
	//
	// example:
	//
	// longcheng
	Voice *string `json:"Voice,omitempty" xml:"Voice,omitempty"`
	// The volume. Default value: 50. Valid values: 0 to 100.
	//
	// example:
	//
	// 70
	Volume *int32 `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s CreateAppAgentTemplateRequestTtsConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateAppAgentTemplateRequestTtsConfig) GoString() string {
	return s.String()
}

func (s *CreateAppAgentTemplateRequestTtsConfig) GetApiKey() *string {
	return s.ApiKey
}

func (s *CreateAppAgentTemplateRequestTtsConfig) GetFilterBrackets() []*int32 {
	return s.FilterBrackets
}

func (s *CreateAppAgentTemplateRequestTtsConfig) GetName() *string {
	return s.Name
}

func (s *CreateAppAgentTemplateRequestTtsConfig) GetPitch() *float32 {
	return s.Pitch
}

func (s *CreateAppAgentTemplateRequestTtsConfig) GetRate() *float32 {
	return s.Rate
}

func (s *CreateAppAgentTemplateRequestTtsConfig) GetVendor() *string {
	return s.Vendor
}

func (s *CreateAppAgentTemplateRequestTtsConfig) GetVoice() *string {
	return s.Voice
}

func (s *CreateAppAgentTemplateRequestTtsConfig) GetVolume() *int32 {
	return s.Volume
}

func (s *CreateAppAgentTemplateRequestTtsConfig) SetApiKey(v string) *CreateAppAgentTemplateRequestTtsConfig {
	s.ApiKey = &v
	return s
}

func (s *CreateAppAgentTemplateRequestTtsConfig) SetFilterBrackets(v []*int32) *CreateAppAgentTemplateRequestTtsConfig {
	s.FilterBrackets = v
	return s
}

func (s *CreateAppAgentTemplateRequestTtsConfig) SetName(v string) *CreateAppAgentTemplateRequestTtsConfig {
	s.Name = &v
	return s
}

func (s *CreateAppAgentTemplateRequestTtsConfig) SetPitch(v float32) *CreateAppAgentTemplateRequestTtsConfig {
	s.Pitch = &v
	return s
}

func (s *CreateAppAgentTemplateRequestTtsConfig) SetRate(v float32) *CreateAppAgentTemplateRequestTtsConfig {
	s.Rate = &v
	return s
}

func (s *CreateAppAgentTemplateRequestTtsConfig) SetVendor(v string) *CreateAppAgentTemplateRequestTtsConfig {
	s.Vendor = &v
	return s
}

func (s *CreateAppAgentTemplateRequestTtsConfig) SetVoice(v string) *CreateAppAgentTemplateRequestTtsConfig {
	s.Voice = &v
	return s
}

func (s *CreateAppAgentTemplateRequestTtsConfig) SetVolume(v int32) *CreateAppAgentTemplateRequestTtsConfig {
	s.Volume = &v
	return s
}

func (s *CreateAppAgentTemplateRequestTtsConfig) Validate() error {
	return dara.Validate(s)
}
