// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppAgentTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentSilenceConfig(v *ModifyAppAgentTemplateRequestAgentSilenceConfig) *ModifyAppAgentTemplateRequest
	GetAgentSilenceConfig() *ModifyAppAgentTemplateRequestAgentSilenceConfig
	SetAmbientSoundConfig(v *ModifyAppAgentTemplateRequestAmbientSoundConfig) *ModifyAppAgentTemplateRequest
	GetAmbientSoundConfig() *ModifyAppAgentTemplateRequestAmbientSoundConfig
	SetAppId(v string) *ModifyAppAgentTemplateRequest
	GetAppId() *string
	SetAsrConfig(v *ModifyAppAgentTemplateRequestAsrConfig) *ModifyAppAgentTemplateRequest
	GetAsrConfig() *ModifyAppAgentTemplateRequestAsrConfig
	SetBackChannelConfig(v *ModifyAppAgentTemplateRequestBackChannelConfig) *ModifyAppAgentTemplateRequest
	GetBackChannelConfig() *ModifyAppAgentTemplateRequestBackChannelConfig
	SetChatMode(v int32) *ModifyAppAgentTemplateRequest
	GetChatMode() *int32
	SetGreeting(v string) *ModifyAppAgentTemplateRequest
	GetGreeting() *string
	SetId(v string) *ModifyAppAgentTemplateRequest
	GetId() *string
	SetInterruptConfig(v *ModifyAppAgentTemplateRequestInterruptConfig) *ModifyAppAgentTemplateRequest
	GetInterruptConfig() *ModifyAppAgentTemplateRequestInterruptConfig
	SetInterruptMode(v int32) *ModifyAppAgentTemplateRequest
	GetInterruptMode() *int32
	SetLlmConfig(v *ModifyAppAgentTemplateRequestLlmConfig) *ModifyAppAgentTemplateRequest
	GetLlmConfig() *ModifyAppAgentTemplateRequestLlmConfig
	SetName(v string) *ModifyAppAgentTemplateRequest
	GetName() *string
	SetTtsConfig(v *ModifyAppAgentTemplateRequestTtsConfig) *ModifyAppAgentTemplateRequest
	GetTtsConfig() *ModifyAppAgentTemplateRequestTtsConfig
	SetType(v int32) *ModifyAppAgentTemplateRequest
	GetType() *int32
}

type ModifyAppAgentTemplateRequest struct {
	AgentSilenceConfig *ModifyAppAgentTemplateRequestAgentSilenceConfig `json:"AgentSilenceConfig,omitempty" xml:"AgentSilenceConfig,omitempty" type:"Struct"`
	AmbientSoundConfig *ModifyAppAgentTemplateRequestAmbientSoundConfig `json:"AmbientSoundConfig,omitempty" xml:"AmbientSoundConfig,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// ac7N****
	AppId             *string                                         `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AsrConfig         *ModifyAppAgentTemplateRequestAsrConfig         `json:"AsrConfig,omitempty" xml:"AsrConfig,omitempty" type:"Struct"`
	BackChannelConfig *ModifyAppAgentTemplateRequestBackChannelConfig `json:"BackChannelConfig,omitempty" xml:"BackChannelConfig,omitempty" type:"Struct"`
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
	Id              *string                                       `json:"Id,omitempty" xml:"Id,omitempty"`
	InterruptConfig *ModifyAppAgentTemplateRequestInterruptConfig `json:"InterruptConfig,omitempty" xml:"InterruptConfig,omitempty" type:"Struct"`
	// example:
	//
	// 2
	InterruptMode *int32                                  `json:"InterruptMode,omitempty" xml:"InterruptMode,omitempty"`
	LlmConfig     *ModifyAppAgentTemplateRequestLlmConfig `json:"LlmConfig,omitempty" xml:"LlmConfig,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 智能体模版
	Name      *string                                 `json:"Name,omitempty" xml:"Name,omitempty"`
	TtsConfig *ModifyAppAgentTemplateRequestTtsConfig `json:"TtsConfig,omitempty" xml:"TtsConfig,omitempty" type:"Struct"`
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ModifyAppAgentTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppAgentTemplateRequest) GoString() string {
	return s.String()
}

func (s *ModifyAppAgentTemplateRequest) GetAgentSilenceConfig() *ModifyAppAgentTemplateRequestAgentSilenceConfig {
	return s.AgentSilenceConfig
}

func (s *ModifyAppAgentTemplateRequest) GetAmbientSoundConfig() *ModifyAppAgentTemplateRequestAmbientSoundConfig {
	return s.AmbientSoundConfig
}

func (s *ModifyAppAgentTemplateRequest) GetAppId() *string {
	return s.AppId
}

func (s *ModifyAppAgentTemplateRequest) GetAsrConfig() *ModifyAppAgentTemplateRequestAsrConfig {
	return s.AsrConfig
}

func (s *ModifyAppAgentTemplateRequest) GetBackChannelConfig() *ModifyAppAgentTemplateRequestBackChannelConfig {
	return s.BackChannelConfig
}

func (s *ModifyAppAgentTemplateRequest) GetChatMode() *int32 {
	return s.ChatMode
}

func (s *ModifyAppAgentTemplateRequest) GetGreeting() *string {
	return s.Greeting
}

func (s *ModifyAppAgentTemplateRequest) GetId() *string {
	return s.Id
}

func (s *ModifyAppAgentTemplateRequest) GetInterruptConfig() *ModifyAppAgentTemplateRequestInterruptConfig {
	return s.InterruptConfig
}

func (s *ModifyAppAgentTemplateRequest) GetInterruptMode() *int32 {
	return s.InterruptMode
}

func (s *ModifyAppAgentTemplateRequest) GetLlmConfig() *ModifyAppAgentTemplateRequestLlmConfig {
	return s.LlmConfig
}

func (s *ModifyAppAgentTemplateRequest) GetName() *string {
	return s.Name
}

func (s *ModifyAppAgentTemplateRequest) GetTtsConfig() *ModifyAppAgentTemplateRequestTtsConfig {
	return s.TtsConfig
}

func (s *ModifyAppAgentTemplateRequest) GetType() *int32 {
	return s.Type
}

func (s *ModifyAppAgentTemplateRequest) SetAgentSilenceConfig(v *ModifyAppAgentTemplateRequestAgentSilenceConfig) *ModifyAppAgentTemplateRequest {
	s.AgentSilenceConfig = v
	return s
}

func (s *ModifyAppAgentTemplateRequest) SetAmbientSoundConfig(v *ModifyAppAgentTemplateRequestAmbientSoundConfig) *ModifyAppAgentTemplateRequest {
	s.AmbientSoundConfig = v
	return s
}

func (s *ModifyAppAgentTemplateRequest) SetAppId(v string) *ModifyAppAgentTemplateRequest {
	s.AppId = &v
	return s
}

func (s *ModifyAppAgentTemplateRequest) SetAsrConfig(v *ModifyAppAgentTemplateRequestAsrConfig) *ModifyAppAgentTemplateRequest {
	s.AsrConfig = v
	return s
}

func (s *ModifyAppAgentTemplateRequest) SetBackChannelConfig(v *ModifyAppAgentTemplateRequestBackChannelConfig) *ModifyAppAgentTemplateRequest {
	s.BackChannelConfig = v
	return s
}

func (s *ModifyAppAgentTemplateRequest) SetChatMode(v int32) *ModifyAppAgentTemplateRequest {
	s.ChatMode = &v
	return s
}

func (s *ModifyAppAgentTemplateRequest) SetGreeting(v string) *ModifyAppAgentTemplateRequest {
	s.Greeting = &v
	return s
}

func (s *ModifyAppAgentTemplateRequest) SetId(v string) *ModifyAppAgentTemplateRequest {
	s.Id = &v
	return s
}

func (s *ModifyAppAgentTemplateRequest) SetInterruptConfig(v *ModifyAppAgentTemplateRequestInterruptConfig) *ModifyAppAgentTemplateRequest {
	s.InterruptConfig = v
	return s
}

func (s *ModifyAppAgentTemplateRequest) SetInterruptMode(v int32) *ModifyAppAgentTemplateRequest {
	s.InterruptMode = &v
	return s
}

func (s *ModifyAppAgentTemplateRequest) SetLlmConfig(v *ModifyAppAgentTemplateRequestLlmConfig) *ModifyAppAgentTemplateRequest {
	s.LlmConfig = v
	return s
}

func (s *ModifyAppAgentTemplateRequest) SetName(v string) *ModifyAppAgentTemplateRequest {
	s.Name = &v
	return s
}

func (s *ModifyAppAgentTemplateRequest) SetTtsConfig(v *ModifyAppAgentTemplateRequestTtsConfig) *ModifyAppAgentTemplateRequest {
	s.TtsConfig = v
	return s
}

func (s *ModifyAppAgentTemplateRequest) SetType(v int32) *ModifyAppAgentTemplateRequest {
	s.Type = &v
	return s
}

func (s *ModifyAppAgentTemplateRequest) Validate() error {
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

type ModifyAppAgentTemplateRequestAgentSilenceConfig struct {
	// example:
	//
	// 30
	AlertTimeout *int32  `json:"AlertTimeout,omitempty" xml:"AlertTimeout,omitempty"`
	Content      *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 2
	Strategy *int32 `json:"Strategy,omitempty" xml:"Strategy,omitempty"`
	// example:
	//
	// 30
	WebhookTriggerTimeout *int32 `json:"WebhookTriggerTimeout,omitempty" xml:"WebhookTriggerTimeout,omitempty"`
}

func (s ModifyAppAgentTemplateRequestAgentSilenceConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppAgentTemplateRequestAgentSilenceConfig) GoString() string {
	return s.String()
}

func (s *ModifyAppAgentTemplateRequestAgentSilenceConfig) GetAlertTimeout() *int32 {
	return s.AlertTimeout
}

func (s *ModifyAppAgentTemplateRequestAgentSilenceConfig) GetContent() *string {
	return s.Content
}

func (s *ModifyAppAgentTemplateRequestAgentSilenceConfig) GetStrategy() *int32 {
	return s.Strategy
}

func (s *ModifyAppAgentTemplateRequestAgentSilenceConfig) GetWebhookTriggerTimeout() *int32 {
	return s.WebhookTriggerTimeout
}

func (s *ModifyAppAgentTemplateRequestAgentSilenceConfig) SetAlertTimeout(v int32) *ModifyAppAgentTemplateRequestAgentSilenceConfig {
	s.AlertTimeout = &v
	return s
}

func (s *ModifyAppAgentTemplateRequestAgentSilenceConfig) SetContent(v string) *ModifyAppAgentTemplateRequestAgentSilenceConfig {
	s.Content = &v
	return s
}

func (s *ModifyAppAgentTemplateRequestAgentSilenceConfig) SetStrategy(v int32) *ModifyAppAgentTemplateRequestAgentSilenceConfig {
	s.Strategy = &v
	return s
}

func (s *ModifyAppAgentTemplateRequestAgentSilenceConfig) SetWebhookTriggerTimeout(v int32) *ModifyAppAgentTemplateRequestAgentSilenceConfig {
	s.WebhookTriggerTimeout = &v
	return s
}

func (s *ModifyAppAgentTemplateRequestAgentSilenceConfig) Validate() error {
	return dara.Validate(s)
}

type ModifyAppAgentTemplateRequestAmbientSoundConfig struct {
	// example:
	//
	// white_noise
	SoundId *string `json:"SoundId,omitempty" xml:"SoundId,omitempty"`
	// example:
	//
	// 100
	Volume *int32 `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s ModifyAppAgentTemplateRequestAmbientSoundConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppAgentTemplateRequestAmbientSoundConfig) GoString() string {
	return s.String()
}

func (s *ModifyAppAgentTemplateRequestAmbientSoundConfig) GetSoundId() *string {
	return s.SoundId
}

func (s *ModifyAppAgentTemplateRequestAmbientSoundConfig) GetVolume() *int32 {
	return s.Volume
}

func (s *ModifyAppAgentTemplateRequestAmbientSoundConfig) SetSoundId(v string) *ModifyAppAgentTemplateRequestAmbientSoundConfig {
	s.SoundId = &v
	return s
}

func (s *ModifyAppAgentTemplateRequestAmbientSoundConfig) SetVolume(v int32) *ModifyAppAgentTemplateRequestAmbientSoundConfig {
	s.Volume = &v
	return s
}

func (s *ModifyAppAgentTemplateRequestAmbientSoundConfig) Validate() error {
	return dara.Validate(s)
}

type ModifyAppAgentTemplateRequestAsrConfig struct {
	// example:
	//
	// 500
	MaxSentenceSilence *int32 `json:"MaxSentenceSilence,omitempty" xml:"MaxSentenceSilence,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// STT
	Name      *string                                          `json:"Name,omitempty" xml:"Name,omitempty"`
	VadConfig *ModifyAppAgentTemplateRequestAsrConfigVadConfig `json:"VadConfig,omitempty" xml:"VadConfig,omitempty" type:"Struct"`
	// example:
	//
	// 0bb1776b1745123332074d1b6b
	VocabularyId *string                                              `json:"VocabularyId,omitempty" xml:"VocabularyId,omitempty"`
	WordWeights  []*ModifyAppAgentTemplateRequestAsrConfigWordWeights `json:"WordWeights,omitempty" xml:"WordWeights,omitempty" type:"Repeated"`
}

func (s ModifyAppAgentTemplateRequestAsrConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppAgentTemplateRequestAsrConfig) GoString() string {
	return s.String()
}

func (s *ModifyAppAgentTemplateRequestAsrConfig) GetMaxSentenceSilence() *int32 {
	return s.MaxSentenceSilence
}

func (s *ModifyAppAgentTemplateRequestAsrConfig) GetName() *string {
	return s.Name
}

func (s *ModifyAppAgentTemplateRequestAsrConfig) GetVadConfig() *ModifyAppAgentTemplateRequestAsrConfigVadConfig {
	return s.VadConfig
}

func (s *ModifyAppAgentTemplateRequestAsrConfig) GetVocabularyId() *string {
	return s.VocabularyId
}

func (s *ModifyAppAgentTemplateRequestAsrConfig) GetWordWeights() []*ModifyAppAgentTemplateRequestAsrConfigWordWeights {
	return s.WordWeights
}

func (s *ModifyAppAgentTemplateRequestAsrConfig) SetMaxSentenceSilence(v int32) *ModifyAppAgentTemplateRequestAsrConfig {
	s.MaxSentenceSilence = &v
	return s
}

func (s *ModifyAppAgentTemplateRequestAsrConfig) SetName(v string) *ModifyAppAgentTemplateRequestAsrConfig {
	s.Name = &v
	return s
}

func (s *ModifyAppAgentTemplateRequestAsrConfig) SetVadConfig(v *ModifyAppAgentTemplateRequestAsrConfigVadConfig) *ModifyAppAgentTemplateRequestAsrConfig {
	s.VadConfig = v
	return s
}

func (s *ModifyAppAgentTemplateRequestAsrConfig) SetVocabularyId(v string) *ModifyAppAgentTemplateRequestAsrConfig {
	s.VocabularyId = &v
	return s
}

func (s *ModifyAppAgentTemplateRequestAsrConfig) SetWordWeights(v []*ModifyAppAgentTemplateRequestAsrConfigWordWeights) *ModifyAppAgentTemplateRequestAsrConfig {
	s.WordWeights = v
	return s
}

func (s *ModifyAppAgentTemplateRequestAsrConfig) Validate() error {
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

type ModifyAppAgentTemplateRequestAsrConfigVadConfig struct {
	// example:
	//
	// 1000
	InterruptSpeechDuration *int32 `json:"InterruptSpeechDuration,omitempty" xml:"InterruptSpeechDuration,omitempty"`
}

func (s ModifyAppAgentTemplateRequestAsrConfigVadConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppAgentTemplateRequestAsrConfigVadConfig) GoString() string {
	return s.String()
}

func (s *ModifyAppAgentTemplateRequestAsrConfigVadConfig) GetInterruptSpeechDuration() *int32 {
	return s.InterruptSpeechDuration
}

func (s *ModifyAppAgentTemplateRequestAsrConfigVadConfig) SetInterruptSpeechDuration(v int32) *ModifyAppAgentTemplateRequestAsrConfigVadConfig {
	s.InterruptSpeechDuration = &v
	return s
}

func (s *ModifyAppAgentTemplateRequestAsrConfigVadConfig) Validate() error {
	return dara.Validate(s)
}

type ModifyAppAgentTemplateRequestAsrConfigWordWeights struct {
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 4
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 苹果
	Word *string `json:"Word,omitempty" xml:"Word,omitempty"`
}

func (s ModifyAppAgentTemplateRequestAsrConfigWordWeights) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppAgentTemplateRequestAsrConfigWordWeights) GoString() string {
	return s.String()
}

func (s *ModifyAppAgentTemplateRequestAsrConfigWordWeights) GetLang() *string {
	return s.Lang
}

func (s *ModifyAppAgentTemplateRequestAsrConfigWordWeights) GetWeight() *int32 {
	return s.Weight
}

func (s *ModifyAppAgentTemplateRequestAsrConfigWordWeights) GetWord() *string {
	return s.Word
}

func (s *ModifyAppAgentTemplateRequestAsrConfigWordWeights) SetLang(v string) *ModifyAppAgentTemplateRequestAsrConfigWordWeights {
	s.Lang = &v
	return s
}

func (s *ModifyAppAgentTemplateRequestAsrConfigWordWeights) SetWeight(v int32) *ModifyAppAgentTemplateRequestAsrConfigWordWeights {
	s.Weight = &v
	return s
}

func (s *ModifyAppAgentTemplateRequestAsrConfigWordWeights) SetWord(v string) *ModifyAppAgentTemplateRequestAsrConfigWordWeights {
	s.Word = &v
	return s
}

func (s *ModifyAppAgentTemplateRequestAsrConfigWordWeights) Validate() error {
	return dara.Validate(s)
}

type ModifyAppAgentTemplateRequestBackChannelConfig struct {
	UserTurnEnd *bool `json:"UserTurnEnd,omitempty" xml:"UserTurnEnd,omitempty"`
}

func (s ModifyAppAgentTemplateRequestBackChannelConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppAgentTemplateRequestBackChannelConfig) GoString() string {
	return s.String()
}

func (s *ModifyAppAgentTemplateRequestBackChannelConfig) GetUserTurnEnd() *bool {
	return s.UserTurnEnd
}

func (s *ModifyAppAgentTemplateRequestBackChannelConfig) SetUserTurnEnd(v bool) *ModifyAppAgentTemplateRequestBackChannelConfig {
	s.UserTurnEnd = &v
	return s
}

func (s *ModifyAppAgentTemplateRequestBackChannelConfig) Validate() error {
	return dara.Validate(s)
}

type ModifyAppAgentTemplateRequestInterruptConfig struct {
	SemanticsInterrupt *bool `json:"SemanticsInterrupt,omitempty" xml:"SemanticsInterrupt,omitempty"`
}

func (s ModifyAppAgentTemplateRequestInterruptConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppAgentTemplateRequestInterruptConfig) GoString() string {
	return s.String()
}

func (s *ModifyAppAgentTemplateRequestInterruptConfig) GetSemanticsInterrupt() *bool {
	return s.SemanticsInterrupt
}

func (s *ModifyAppAgentTemplateRequestInterruptConfig) SetSemanticsInterrupt(v bool) *ModifyAppAgentTemplateRequestInterruptConfig {
	s.SemanticsInterrupt = &v
	return s
}

func (s *ModifyAppAgentTemplateRequestInterruptConfig) Validate() error {
	return dara.Validate(s)
}

type ModifyAppAgentTemplateRequestLlmConfig struct {
	AgentAppId *string `json:"AgentAppId,omitempty" xml:"AgentAppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ak-1213123123132123131
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// example:
	//
	// 8
	HistoryDepth *int32 `json:"HistoryDepth,omitempty" xml:"HistoryDepth,omitempty"`
	// example:
	//
	// 1024
	MaxToken *int32 `json:"MaxToken,omitempty" xml:"MaxToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// example:
	//
	// 0.7
	Temperature *float32 `json:"Temperature,omitempty" xml:"Temperature,omitempty"`
	// example:
	//
	// 0.8
	TopP *float32 `json:"TopP,omitempty" xml:"TopP,omitempty"`
	// example:
	//
	// https://llm.example.aliyuns.com
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// example:
	//
	// thirdparty
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s ModifyAppAgentTemplateRequestLlmConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppAgentTemplateRequestLlmConfig) GoString() string {
	return s.String()
}

func (s *ModifyAppAgentTemplateRequestLlmConfig) GetAgentAppId() *string {
	return s.AgentAppId
}

func (s *ModifyAppAgentTemplateRequestLlmConfig) GetApiKey() *string {
	return s.ApiKey
}

func (s *ModifyAppAgentTemplateRequestLlmConfig) GetHistoryDepth() *int32 {
	return s.HistoryDepth
}

func (s *ModifyAppAgentTemplateRequestLlmConfig) GetMaxToken() *int32 {
	return s.MaxToken
}

func (s *ModifyAppAgentTemplateRequestLlmConfig) GetName() *string {
	return s.Name
}

func (s *ModifyAppAgentTemplateRequestLlmConfig) GetPrompt() *string {
	return s.Prompt
}

func (s *ModifyAppAgentTemplateRequestLlmConfig) GetTemperature() *float32 {
	return s.Temperature
}

func (s *ModifyAppAgentTemplateRequestLlmConfig) GetTopP() *float32 {
	return s.TopP
}

func (s *ModifyAppAgentTemplateRequestLlmConfig) GetUrl() *string {
	return s.Url
}

func (s *ModifyAppAgentTemplateRequestLlmConfig) GetVendor() *string {
	return s.Vendor
}

func (s *ModifyAppAgentTemplateRequestLlmConfig) SetAgentAppId(v string) *ModifyAppAgentTemplateRequestLlmConfig {
	s.AgentAppId = &v
	return s
}

func (s *ModifyAppAgentTemplateRequestLlmConfig) SetApiKey(v string) *ModifyAppAgentTemplateRequestLlmConfig {
	s.ApiKey = &v
	return s
}

func (s *ModifyAppAgentTemplateRequestLlmConfig) SetHistoryDepth(v int32) *ModifyAppAgentTemplateRequestLlmConfig {
	s.HistoryDepth = &v
	return s
}

func (s *ModifyAppAgentTemplateRequestLlmConfig) SetMaxToken(v int32) *ModifyAppAgentTemplateRequestLlmConfig {
	s.MaxToken = &v
	return s
}

func (s *ModifyAppAgentTemplateRequestLlmConfig) SetName(v string) *ModifyAppAgentTemplateRequestLlmConfig {
	s.Name = &v
	return s
}

func (s *ModifyAppAgentTemplateRequestLlmConfig) SetPrompt(v string) *ModifyAppAgentTemplateRequestLlmConfig {
	s.Prompt = &v
	return s
}

func (s *ModifyAppAgentTemplateRequestLlmConfig) SetTemperature(v float32) *ModifyAppAgentTemplateRequestLlmConfig {
	s.Temperature = &v
	return s
}

func (s *ModifyAppAgentTemplateRequestLlmConfig) SetTopP(v float32) *ModifyAppAgentTemplateRequestLlmConfig {
	s.TopP = &v
	return s
}

func (s *ModifyAppAgentTemplateRequestLlmConfig) SetUrl(v string) *ModifyAppAgentTemplateRequestLlmConfig {
	s.Url = &v
	return s
}

func (s *ModifyAppAgentTemplateRequestLlmConfig) SetVendor(v string) *ModifyAppAgentTemplateRequestLlmConfig {
	s.Vendor = &v
	return s
}

func (s *ModifyAppAgentTemplateRequestLlmConfig) Validate() error {
	return dara.Validate(s)
}

type ModifyAppAgentTemplateRequestTtsConfig struct {
	// This parameter is required.
	//
	// example:
	//
	// ak-1213123123132123131
	ApiKey         *string  `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	FilterBrackets []*int32 `json:"FilterBrackets,omitempty" xml:"FilterBrackets,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// Tts
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 0.8
	Pitch *float32 `json:"Pitch,omitempty" xml:"Pitch,omitempty"`
	// example:
	//
	// 0.8
	Rate *float32 `json:"Rate,omitempty" xml:"Rate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// aliyun
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
	// example:
	//
	// longwan
	Voice *string `json:"Voice,omitempty" xml:"Voice,omitempty"`
	// example:
	//
	// 50
	Volume *int32 `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s ModifyAppAgentTemplateRequestTtsConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppAgentTemplateRequestTtsConfig) GoString() string {
	return s.String()
}

func (s *ModifyAppAgentTemplateRequestTtsConfig) GetApiKey() *string {
	return s.ApiKey
}

func (s *ModifyAppAgentTemplateRequestTtsConfig) GetFilterBrackets() []*int32 {
	return s.FilterBrackets
}

func (s *ModifyAppAgentTemplateRequestTtsConfig) GetName() *string {
	return s.Name
}

func (s *ModifyAppAgentTemplateRequestTtsConfig) GetPitch() *float32 {
	return s.Pitch
}

func (s *ModifyAppAgentTemplateRequestTtsConfig) GetRate() *float32 {
	return s.Rate
}

func (s *ModifyAppAgentTemplateRequestTtsConfig) GetVendor() *string {
	return s.Vendor
}

func (s *ModifyAppAgentTemplateRequestTtsConfig) GetVoice() *string {
	return s.Voice
}

func (s *ModifyAppAgentTemplateRequestTtsConfig) GetVolume() *int32 {
	return s.Volume
}

func (s *ModifyAppAgentTemplateRequestTtsConfig) SetApiKey(v string) *ModifyAppAgentTemplateRequestTtsConfig {
	s.ApiKey = &v
	return s
}

func (s *ModifyAppAgentTemplateRequestTtsConfig) SetFilterBrackets(v []*int32) *ModifyAppAgentTemplateRequestTtsConfig {
	s.FilterBrackets = v
	return s
}

func (s *ModifyAppAgentTemplateRequestTtsConfig) SetName(v string) *ModifyAppAgentTemplateRequestTtsConfig {
	s.Name = &v
	return s
}

func (s *ModifyAppAgentTemplateRequestTtsConfig) SetPitch(v float32) *ModifyAppAgentTemplateRequestTtsConfig {
	s.Pitch = &v
	return s
}

func (s *ModifyAppAgentTemplateRequestTtsConfig) SetRate(v float32) *ModifyAppAgentTemplateRequestTtsConfig {
	s.Rate = &v
	return s
}

func (s *ModifyAppAgentTemplateRequestTtsConfig) SetVendor(v string) *ModifyAppAgentTemplateRequestTtsConfig {
	s.Vendor = &v
	return s
}

func (s *ModifyAppAgentTemplateRequestTtsConfig) SetVoice(v string) *ModifyAppAgentTemplateRequestTtsConfig {
	s.Voice = &v
	return s
}

func (s *ModifyAppAgentTemplateRequestTtsConfig) SetVolume(v int32) *ModifyAppAgentTemplateRequestTtsConfig {
	s.Volume = &v
	return s
}

func (s *ModifyAppAgentTemplateRequestTtsConfig) Validate() error {
	return dara.Validate(s)
}
