// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppAgentTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeAppAgentTemplatesResponseBody
	GetRequestId() *string
	SetTemplates(v []*DescribeAppAgentTemplatesResponseBodyTemplates) *DescribeAppAgentTemplatesResponseBody
	GetTemplates() []*DescribeAppAgentTemplatesResponseBodyTemplates
	SetTotalNum(v int64) *DescribeAppAgentTemplatesResponseBody
	GetTotalNum() *int64
	SetTotalPage(v int64) *DescribeAppAgentTemplatesResponseBody
	GetTotalPage() *int64
}

type DescribeAppAgentTemplatesResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 30D41049-D02D-1C21-86AE-B3E5FD805C27
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of templates.
	Templates []*DescribeAppAgentTemplatesResponseBodyTemplates `json:"Templates,omitempty" xml:"Templates,omitempty" type:"Repeated"`
	// The total number of records.
	//
	// example:
	//
	// 10
	TotalNum *int64 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// The total number of pages.
	//
	// example:
	//
	// 1
	TotalPage *int64 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s DescribeAppAgentTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppAgentTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppAgentTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAppAgentTemplatesResponseBody) GetTemplates() []*DescribeAppAgentTemplatesResponseBodyTemplates {
	return s.Templates
}

func (s *DescribeAppAgentTemplatesResponseBody) GetTotalNum() *int64 {
	return s.TotalNum
}

func (s *DescribeAppAgentTemplatesResponseBody) GetTotalPage() *int64 {
	return s.TotalPage
}

func (s *DescribeAppAgentTemplatesResponseBody) SetRequestId(v string) *DescribeAppAgentTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBody) SetTemplates(v []*DescribeAppAgentTemplatesResponseBodyTemplates) *DescribeAppAgentTemplatesResponseBody {
	s.Templates = v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBody) SetTotalNum(v int64) *DescribeAppAgentTemplatesResponseBody {
	s.TotalNum = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBody) SetTotalPage(v int64) *DescribeAppAgentTemplatesResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBody) Validate() error {
	if s.Templates != nil {
		for _, item := range s.Templates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAppAgentTemplatesResponseBodyTemplates struct {
	// The silence configuration for the agent.
	AgentSilenceConfig *DescribeAppAgentTemplatesResponseBodyTemplatesAgentSilenceConfig `json:"AgentSilenceConfig,omitempty" xml:"AgentSilenceConfig,omitempty" type:"Struct"`
	// The ambient sound configuration.
	AmbientSoundConfig *DescribeAppAgentTemplatesResponseBodyTemplatesAmbientSoundConfig `json:"AmbientSoundConfig,omitempty" xml:"AmbientSoundConfig,omitempty" type:"Struct"`
	// The speech-to-text configuration.
	AsrConfig *DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfig `json:"AsrConfig,omitempty" xml:"AsrConfig,omitempty" type:"Struct"`
	// The backchannel configuration.
	BackChannelConfig *DescribeAppAgentTemplatesResponseBodyTemplatesBackChannelConfig `json:"BackChannelConfig,omitempty" xml:"BackChannelConfig,omitempty" type:"Struct"`
	// The conversation mode. Valid values:
	//
	// - 1: natural conversation mode (default)
	//
	// - 2: walkie-talkie mode
	//
	// example:
	//
	// 1
	ChatMode *int32 `json:"ChatMode,omitempty" xml:"ChatMode,omitempty"`
	// The creation time. Format: yyyy-MM-ddTHH:mm:ssZ (UTC).
	//
	// example:
	//
	// 2020-09-04T06:22:15Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Specifies whether to enable video understanding. If enabled, the large language model must support multimodal data.
	EnableVideoUnderstanding *bool `json:"EnableVideoUnderstanding,omitempty" xml:"EnableVideoUnderstanding,omitempty"`
	// The greeting.
	//
	// example:
	//
	// Hello, robot
	Greeting *string `json:"Greeting,omitempty" xml:"Greeting,omitempty"`
	// The template ID.
	//
	// example:
	//
	// wv7N****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The interruption configuration.
	InterruptConfig *DescribeAppAgentTemplatesResponseBodyTemplatesInterruptConfig `json:"InterruptConfig,omitempty" xml:"InterruptConfig,omitempty" type:"Struct"`
	// The interruption mode. Valid values:
	//
	// - 1: Allow interruption (default).
	//
	// - 2: Disable interruption.
	//
	// example:
	//
	// 1
	InterruptMode *int32 `json:"InterruptMode,omitempty" xml:"InterruptMode,omitempty"`
	// The large language model (LLM) configuration.
	LlmConfig *DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig `json:"LlmConfig,omitempty" xml:"LlmConfig,omitempty" type:"Struct"`
	// The template name.
	//
	// example:
	//
	// Test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The preferred video stream for the agent to process when visual understanding is enabled. The agent can process only one video stream. When both the camera stream and the screen sharing stream are pushed, valid values are:
	//
	// - 1: The agent prioritizes the camera stream.
	//
	// - 2: The agent prioritizes the screen sharing stream.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PreferVideo *int32 `json:"PreferVideo,omitempty" xml:"PreferVideo,omitempty"`
	// The text-to-speech settings.
	TtsConfig *DescribeAppAgentTemplatesResponseBodyTemplatesTtsConfig `json:"TtsConfig,omitempty" xml:"TtsConfig,omitempty" type:"Struct"`
	// The type of the agent. Valid values:
	//
	// - 1: Voice call (default).
	//
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeAppAgentTemplatesResponseBodyTemplates) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppAgentTemplatesResponseBodyTemplates) GoString() string {
	return s.String()
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplates) GetAgentSilenceConfig() *DescribeAppAgentTemplatesResponseBodyTemplatesAgentSilenceConfig {
	return s.AgentSilenceConfig
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplates) GetAmbientSoundConfig() *DescribeAppAgentTemplatesResponseBodyTemplatesAmbientSoundConfig {
	return s.AmbientSoundConfig
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplates) GetAsrConfig() *DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfig {
	return s.AsrConfig
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplates) GetBackChannelConfig() *DescribeAppAgentTemplatesResponseBodyTemplatesBackChannelConfig {
	return s.BackChannelConfig
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplates) GetChatMode() *int32 {
	return s.ChatMode
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplates) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplates) GetEnableVideoUnderstanding() *bool {
	return s.EnableVideoUnderstanding
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplates) GetGreeting() *string {
	return s.Greeting
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplates) GetId() *string {
	return s.Id
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplates) GetInterruptConfig() *DescribeAppAgentTemplatesResponseBodyTemplatesInterruptConfig {
	return s.InterruptConfig
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplates) GetInterruptMode() *int32 {
	return s.InterruptMode
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplates) GetLlmConfig() *DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig {
	return s.LlmConfig
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplates) GetName() *string {
	return s.Name
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplates) GetPreferVideo() *int32 {
	return s.PreferVideo
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplates) GetTtsConfig() *DescribeAppAgentTemplatesResponseBodyTemplatesTtsConfig {
	return s.TtsConfig
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplates) GetType() *int32 {
	return s.Type
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplates) SetAgentSilenceConfig(v *DescribeAppAgentTemplatesResponseBodyTemplatesAgentSilenceConfig) *DescribeAppAgentTemplatesResponseBodyTemplates {
	s.AgentSilenceConfig = v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplates) SetAmbientSoundConfig(v *DescribeAppAgentTemplatesResponseBodyTemplatesAmbientSoundConfig) *DescribeAppAgentTemplatesResponseBodyTemplates {
	s.AmbientSoundConfig = v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplates) SetAsrConfig(v *DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfig) *DescribeAppAgentTemplatesResponseBodyTemplates {
	s.AsrConfig = v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplates) SetBackChannelConfig(v *DescribeAppAgentTemplatesResponseBodyTemplatesBackChannelConfig) *DescribeAppAgentTemplatesResponseBodyTemplates {
	s.BackChannelConfig = v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplates) SetChatMode(v int32) *DescribeAppAgentTemplatesResponseBodyTemplates {
	s.ChatMode = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplates) SetCreateTime(v string) *DescribeAppAgentTemplatesResponseBodyTemplates {
	s.CreateTime = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplates) SetEnableVideoUnderstanding(v bool) *DescribeAppAgentTemplatesResponseBodyTemplates {
	s.EnableVideoUnderstanding = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplates) SetGreeting(v string) *DescribeAppAgentTemplatesResponseBodyTemplates {
	s.Greeting = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplates) SetId(v string) *DescribeAppAgentTemplatesResponseBodyTemplates {
	s.Id = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplates) SetInterruptConfig(v *DescribeAppAgentTemplatesResponseBodyTemplatesInterruptConfig) *DescribeAppAgentTemplatesResponseBodyTemplates {
	s.InterruptConfig = v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplates) SetInterruptMode(v int32) *DescribeAppAgentTemplatesResponseBodyTemplates {
	s.InterruptMode = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplates) SetLlmConfig(v *DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig) *DescribeAppAgentTemplatesResponseBodyTemplates {
	s.LlmConfig = v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplates) SetName(v string) *DescribeAppAgentTemplatesResponseBodyTemplates {
	s.Name = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplates) SetPreferVideo(v int32) *DescribeAppAgentTemplatesResponseBodyTemplates {
	s.PreferVideo = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplates) SetTtsConfig(v *DescribeAppAgentTemplatesResponseBodyTemplatesTtsConfig) *DescribeAppAgentTemplatesResponseBodyTemplates {
	s.TtsConfig = v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplates) SetType(v int32) *DescribeAppAgentTemplatesResponseBodyTemplates {
	s.Type = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplates) Validate() error {
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

type DescribeAppAgentTemplatesResponseBodyTemplatesAgentSilenceConfig struct {
	// The timeout period for the silence alert, in seconds. This alert can be triggered multiple times. If the agent remains silent after the silence alert timeout, it will be triggered again upon the next timeout.
	//
	// example:
	//
	// 60
	AlertTimeout *int32 `json:"AlertTimeout,omitempty" xml:"AlertTimeout,omitempty"`
	// The content of the silence prompt message.
	//
	// example:
	//
	// Hello, are you still there?
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The silence timeout handling strategy. Valid values:
	//
	// - 1: Broadcast the prompt message specified by content to the user via voice.
	//
	// - 2: Send the prompt message specified by content and the conversation history to the LLM, and broadcast the reply generated by the model.
	//
	// example:
	//
	// 1
	Strategy *int32 `json:"Strategy,omitempty" xml:"Strategy,omitempty"`
	// The silence timeout period for triggering the agent silence timeout webhook, in seconds. This webhook can be triggered multiple times. If the agent remains silent after the webhook silence timeout, it will be triggered again upon the next timeout.
	//
	// example:
	//
	// 50
	WebhookTriggerTimeout *int32 `json:"WebhookTriggerTimeout,omitempty" xml:"WebhookTriggerTimeout,omitempty"`
}

func (s DescribeAppAgentTemplatesResponseBodyTemplatesAgentSilenceConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppAgentTemplatesResponseBodyTemplatesAgentSilenceConfig) GoString() string {
	return s.String()
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesAgentSilenceConfig) GetAlertTimeout() *int32 {
	return s.AlertTimeout
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesAgentSilenceConfig) GetContent() *string {
	return s.Content
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesAgentSilenceConfig) GetStrategy() *int32 {
	return s.Strategy
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesAgentSilenceConfig) GetWebhookTriggerTimeout() *int32 {
	return s.WebhookTriggerTimeout
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesAgentSilenceConfig) SetAlertTimeout(v int32) *DescribeAppAgentTemplatesResponseBodyTemplatesAgentSilenceConfig {
	s.AlertTimeout = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesAgentSilenceConfig) SetContent(v string) *DescribeAppAgentTemplatesResponseBodyTemplatesAgentSilenceConfig {
	s.Content = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesAgentSilenceConfig) SetStrategy(v int32) *DescribeAppAgentTemplatesResponseBodyTemplatesAgentSilenceConfig {
	s.Strategy = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesAgentSilenceConfig) SetWebhookTriggerTimeout(v int32) *DescribeAppAgentTemplatesResponseBodyTemplatesAgentSilenceConfig {
	s.WebhookTriggerTimeout = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesAgentSilenceConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeAppAgentTemplatesResponseBodyTemplatesAmbientSoundConfig struct {
	// The ID of the ambient sound.
	//
	// example:
	//
	// office
	SoundId *string `json:"SoundId,omitempty" xml:"SoundId,omitempty"`
	// The volume of the ambient sound.
	//
	// example:
	//
	// 50
	Volume *int32 `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s DescribeAppAgentTemplatesResponseBodyTemplatesAmbientSoundConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppAgentTemplatesResponseBodyTemplatesAmbientSoundConfig) GoString() string {
	return s.String()
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesAmbientSoundConfig) GetSoundId() *string {
	return s.SoundId
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesAmbientSoundConfig) GetVolume() *int32 {
	return s.Volume
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesAmbientSoundConfig) SetSoundId(v string) *DescribeAppAgentTemplatesResponseBodyTemplatesAmbientSoundConfig {
	s.SoundId = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesAmbientSoundConfig) SetVolume(v int32) *DescribeAppAgentTemplatesResponseBodyTemplatesAmbientSoundConfig {
	s.Volume = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesAmbientSoundConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfig struct {
	// The threshold for detecting sentence pauses in speech, in milliseconds. Default value: 500. Valid values: 200 to 6000.
	//
	// example:
	//
	// 300
	MaxSentenceSilence *int32 `json:"MaxSentenceSilence,omitempty" xml:"MaxSentenceSilence,omitempty"`
	// The name.
	//
	// example:
	//
	// ASR configuration
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The Voice Activity Detection (VAD) configuration.
	VadConfig *DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfigVadConfig `json:"VadConfig,omitempty" xml:"VadConfig,omitempty" type:"Struct"`
	// The ID of the hotword vocabulary.
	//
	// example:
	//
	// ecfadace11114cf08a7f07aceee798ad
	VocabularyId *string `json:"VocabularyId,omitempty" xml:"VocabularyId,omitempty"`
	// The hotword weights. Each vocabulary can contain up to 500 hotwords, and each hotword can be up to 10 Chinese characters in length. Hotwords must be UTF-8 encoded and cannot contain punctuation marks or special characters. Otherwise, the hotword feature will not take effect.
	WordWeights []*DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfigWordWeights `json:"WordWeights,omitempty" xml:"WordWeights,omitempty" type:"Repeated"`
}

func (s DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfig) GoString() string {
	return s.String()
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfig) GetMaxSentenceSilence() *int32 {
	return s.MaxSentenceSilence
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfig) GetName() *string {
	return s.Name
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfig) GetVadConfig() *DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfigVadConfig {
	return s.VadConfig
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfig) GetVocabularyId() *string {
	return s.VocabularyId
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfig) GetWordWeights() []*DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfigWordWeights {
	return s.WordWeights
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfig) SetMaxSentenceSilence(v int32) *DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfig {
	s.MaxSentenceSilence = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfig) SetName(v string) *DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfig {
	s.Name = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfig) SetVadConfig(v *DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfigVadConfig) *DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfig {
	s.VadConfig = v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfig) SetVocabularyId(v string) *DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfig {
	s.VocabularyId = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfig) SetWordWeights(v []*DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfigWordWeights) *DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfig {
	s.WordWeights = v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfig) Validate() error {
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

type DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfigVadConfig struct {
	// The minimum duration of continuous human voice signal detection required to avoid false interruptions, in milliseconds. Valid values: 0 to 3000.
	//
	// example:
	//
	// 1200
	InterruptSpeechDuration *int32 `json:"InterruptSpeechDuration,omitempty" xml:"InterruptSpeechDuration,omitempty"`
}

func (s DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfigVadConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfigVadConfig) GoString() string {
	return s.String()
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfigVadConfig) GetInterruptSpeechDuration() *int32 {
	return s.InterruptSpeechDuration
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfigVadConfig) SetInterruptSpeechDuration(v int32) *DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfigVadConfig {
	s.InterruptSpeechDuration = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfigVadConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfigWordWeights struct {
	// The language code.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The weight of the hotword. Valid values: integers from 1 to 5. Common value: 4. If the effect is not obvious, you can increase the weight. However, a large weight may cause negative effects and lead to inaccurate recognition of other words.
	//
	// example:
	//
	// 1
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
	// The text of the hotword. Each word can be up to 10 Chinese characters or English words in length.
	//
	// example:
	//
	// Dezhou
	Word *string `json:"Word,omitempty" xml:"Word,omitempty"`
}

func (s DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfigWordWeights) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfigWordWeights) GoString() string {
	return s.String()
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfigWordWeights) GetLang() *string {
	return s.Lang
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfigWordWeights) GetWeight() *int32 {
	return s.Weight
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfigWordWeights) GetWord() *string {
	return s.Word
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfigWordWeights) SetLang(v string) *DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfigWordWeights {
	s.Lang = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfigWordWeights) SetWeight(v int32) *DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfigWordWeights {
	s.Weight = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfigWordWeights) SetWord(v string) *DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfigWordWeights {
	s.Word = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesAsrConfigWordWeights) Validate() error {
	return dara.Validate(s)
}

type DescribeAppAgentTemplatesResponseBodyTemplatesBackChannelConfig struct {
	// Specifies whether to enable backchanneling after the user finishes speaking.
	UserTurnEnd *bool `json:"UserTurnEnd,omitempty" xml:"UserTurnEnd,omitempty"`
}

func (s DescribeAppAgentTemplatesResponseBodyTemplatesBackChannelConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppAgentTemplatesResponseBodyTemplatesBackChannelConfig) GoString() string {
	return s.String()
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesBackChannelConfig) GetUserTurnEnd() *bool {
	return s.UserTurnEnd
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesBackChannelConfig) SetUserTurnEnd(v bool) *DescribeAppAgentTemplatesResponseBodyTemplatesBackChannelConfig {
	s.UserTurnEnd = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesBackChannelConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeAppAgentTemplatesResponseBodyTemplatesInterruptConfig struct {
	// Specifies whether semantic interruption is enabled.
	SemanticsInterrupt *bool `json:"SemanticsInterrupt,omitempty" xml:"SemanticsInterrupt,omitempty"`
}

func (s DescribeAppAgentTemplatesResponseBodyTemplatesInterruptConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppAgentTemplatesResponseBodyTemplatesInterruptConfig) GoString() string {
	return s.String()
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesInterruptConfig) GetSemanticsInterrupt() *bool {
	return s.SemanticsInterrupt
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesInterruptConfig) SetSemanticsInterrupt(v bool) *DescribeAppAgentTemplatesResponseBodyTemplatesInterruptConfig {
	s.SemanticsInterrupt = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesInterruptConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig struct {
	// The application ID of the agent when the model is a Model Studio agent.
	//
	// example:
	//
	// 1233312
	AgentAppId *string `json:"AgentAppId,omitempty" xml:"AgentAppId,omitempty"`
	// The API key required by the LLM service provider.
	//
	// example:
	//
	// qW8GpBOdHK/pv9gdUSVLvQ==
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// The number of historical conversation turns provided to the LLM. Maximum value: 100. Default value: 3.
	//
	// example:
	//
	// 8
	HistoryDepth *int32 `json:"HistoryDepth,omitempty" xml:"HistoryDepth,omitempty"`
	// The maximum number of tokens for the output text. Default value: 8192.
	//
	// example:
	//
	// 1024
	MaxToken *int32 `json:"MaxToken,omitempty" xml:"MaxToken,omitempty"`
	// The model provided by Alibaba Cloud Model Studio. Currently, qwen-plus and deepseek-r1 are supported.
	//
	// example:
	//
	// deepseek-r1
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// The name.
	//
	// example:
	//
	// llm
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The prompt for the large language model. The maximum length is 5,000 characters.
	//
	// example:
	//
	// You are a senior psychologist
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// The sampling temperature. Valid values: 0 to 2 (exclusive). Default value: 0.7.
	//
	// example:
	//
	// 0.8
	Temperature *float32 `json:"Temperature,omitempty" xml:"Temperature,omitempty"`
	// The top-p sampling range. Valid values: (0, 1]. Default value: 0.8.
	//
	// example:
	//
	// 0.8
	TopP *float32 `json:"TopP,omitempty" xml:"TopP,omitempty"`
	// The endpoint URL used to connect to a third-party large language model.
	//
	// example:
	//
	// https://test.com
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// The vendor. Currently, Alibaba Cloud Model Studio and third-party configurations that comply with the OpenAI API specification are supported.
	//
	// example:
	//
	// aliyun
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig) GoString() string {
	return s.String()
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig) GetAgentAppId() *string {
	return s.AgentAppId
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig) GetApiKey() *string {
	return s.ApiKey
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig) GetHistoryDepth() *int32 {
	return s.HistoryDepth
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig) GetMaxToken() *int32 {
	return s.MaxToken
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig) GetModel() *string {
	return s.Model
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig) GetName() *string {
	return s.Name
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig) GetPrompt() *string {
	return s.Prompt
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig) GetTemperature() *float32 {
	return s.Temperature
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig) GetTopP() *float32 {
	return s.TopP
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig) GetUrl() *string {
	return s.Url
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig) GetVendor() *string {
	return s.Vendor
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig) SetAgentAppId(v string) *DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig {
	s.AgentAppId = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig) SetApiKey(v string) *DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig {
	s.ApiKey = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig) SetHistoryDepth(v int32) *DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig {
	s.HistoryDepth = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig) SetMaxToken(v int32) *DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig {
	s.MaxToken = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig) SetModel(v string) *DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig {
	s.Model = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig) SetName(v string) *DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig {
	s.Name = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig) SetPrompt(v string) *DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig {
	s.Prompt = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig) SetTemperature(v float32) *DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig {
	s.Temperature = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig) SetTopP(v float32) *DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig {
	s.TopP = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig) SetUrl(v string) *DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig {
	s.Url = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig) SetVendor(v string) *DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig {
	s.Vendor = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesLlmConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeAppAgentTemplatesResponseBodyTemplatesTtsConfig struct {
	// The API key required by the LLM service provider.
	//
	// example:
	//
	// N5448VFGI2mXJU8a/A03VQ==
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// The bracket filtering settings. Valid values:
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
	// The model provided by Alibaba Cloud Model Studio. Currently, only the cosyvoice-v1 model is supported.
	//
	// example:
	//
	// cosyvoice-v1
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// The name.
	//
	// example:
	//
	// Tts
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The pitch. Valid values: 0.5 to 2.0. Default value: 1.0.
	//
	// example:
	//
	// 0.8
	Pitch *float32 `json:"Pitch,omitempty" xml:"Pitch,omitempty"`
	// The speech rate. Valid values: 0.5 to 2.0. Default value: 1.0.
	//
	// example:
	//
	// 50
	Rate *float32 `json:"Rate,omitempty" xml:"Rate,omitempty"`
	// The vendor. Currently, only Alibaba Cloud Model Studio is supported.
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
	// longwan
	Voice *string `json:"Voice,omitempty" xml:"Voice,omitempty"`
	// The volume. Valid values: 0 to 100. Default value: 50.
	//
	// example:
	//
	// 50
	Volume *int32 `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s DescribeAppAgentTemplatesResponseBodyTemplatesTtsConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppAgentTemplatesResponseBodyTemplatesTtsConfig) GoString() string {
	return s.String()
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesTtsConfig) GetApiKey() *string {
	return s.ApiKey
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesTtsConfig) GetFilterBrackets() []*int32 {
	return s.FilterBrackets
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesTtsConfig) GetModel() *string {
	return s.Model
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesTtsConfig) GetName() *string {
	return s.Name
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesTtsConfig) GetPitch() *float32 {
	return s.Pitch
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesTtsConfig) GetRate() *float32 {
	return s.Rate
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesTtsConfig) GetVendor() *string {
	return s.Vendor
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesTtsConfig) GetVoice() *string {
	return s.Voice
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesTtsConfig) GetVolume() *int32 {
	return s.Volume
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesTtsConfig) SetApiKey(v string) *DescribeAppAgentTemplatesResponseBodyTemplatesTtsConfig {
	s.ApiKey = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesTtsConfig) SetFilterBrackets(v []*int32) *DescribeAppAgentTemplatesResponseBodyTemplatesTtsConfig {
	s.FilterBrackets = v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesTtsConfig) SetModel(v string) *DescribeAppAgentTemplatesResponseBodyTemplatesTtsConfig {
	s.Model = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesTtsConfig) SetName(v string) *DescribeAppAgentTemplatesResponseBodyTemplatesTtsConfig {
	s.Name = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesTtsConfig) SetPitch(v float32) *DescribeAppAgentTemplatesResponseBodyTemplatesTtsConfig {
	s.Pitch = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesTtsConfig) SetRate(v float32) *DescribeAppAgentTemplatesResponseBodyTemplatesTtsConfig {
	s.Rate = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesTtsConfig) SetVendor(v string) *DescribeAppAgentTemplatesResponseBodyTemplatesTtsConfig {
	s.Vendor = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesTtsConfig) SetVoice(v string) *DescribeAppAgentTemplatesResponseBodyTemplatesTtsConfig {
	s.Voice = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesTtsConfig) SetVolume(v int32) *DescribeAppAgentTemplatesResponseBodyTemplatesTtsConfig {
	s.Volume = &v
	return s
}

func (s *DescribeAppAgentTemplatesResponseBodyTemplatesTtsConfig) Validate() error {
	return dara.Validate(s)
}
