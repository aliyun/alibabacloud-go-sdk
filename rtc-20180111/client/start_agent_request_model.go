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
	// This parameter is required.
	//
	// example:
	//
	// aoe****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// yourChannelId
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// This parameter is required.
	RtcConfig *StartAgentRequestRtcConfig `json:"RtcConfig,omitempty" xml:"RtcConfig,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// yourTaskId
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 76dasgb****
	TemplateId      *string                           `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
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
	return dara.Validate(s)
}

type StartAgentRequestRtcConfig struct {
	TargetUserIds []*string `json:"TargetUserIds,omitempty" xml:"TargetUserIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 423341
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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

func (s *StartAgentRequestRtcConfig) SetTargetUserIds(v []*string) *StartAgentRequestRtcConfig {
	s.TargetUserIds = v
	return s
}

func (s *StartAgentRequestRtcConfig) SetUserId(v string) *StartAgentRequestRtcConfig {
	s.UserId = &v
	return s
}

func (s *StartAgentRequestRtcConfig) Validate() error {
	return dara.Validate(s)
}

type StartAgentRequestVoiceChatConfig struct {
	ASRConfig *StartAgentRequestVoiceChatConfigASRConfig `json:"ASRConfig,omitempty" xml:"ASRConfig,omitempty" type:"Struct"`
	// example:
	//
	// 1
	ChatMode *int32  `json:"ChatMode,omitempty" xml:"ChatMode,omitempty"`
	Greeting *string `json:"Greeting,omitempty" xml:"Greeting,omitempty"`
	// example:
	//
	// 1
	InterruptMode *int32                                     `json:"InterruptMode,omitempty" xml:"InterruptMode,omitempty"`
	LLMConfig     *StartAgentRequestVoiceChatConfigLLMConfig `json:"LLMConfig,omitempty" xml:"LLMConfig,omitempty" type:"Struct"`
	TTSConfig     *StartAgentRequestVoiceChatConfigTTSConfig `json:"TTSConfig,omitempty" xml:"TTSConfig,omitempty" type:"Struct"`
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

func (s *StartAgentRequestVoiceChatConfig) GetChatMode() *int32 {
	return s.ChatMode
}

func (s *StartAgentRequestVoiceChatConfig) GetGreeting() *string {
	return s.Greeting
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

func (s *StartAgentRequestVoiceChatConfig) SetASRConfig(v *StartAgentRequestVoiceChatConfigASRConfig) *StartAgentRequestVoiceChatConfig {
	s.ASRConfig = v
	return s
}

func (s *StartAgentRequestVoiceChatConfig) SetChatMode(v int32) *StartAgentRequestVoiceChatConfig {
	s.ChatMode = &v
	return s
}

func (s *StartAgentRequestVoiceChatConfig) SetGreeting(v string) *StartAgentRequestVoiceChatConfig {
	s.Greeting = &v
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

func (s *StartAgentRequestVoiceChatConfig) Validate() error {
	return dara.Validate(s)
}

type StartAgentRequestVoiceChatConfigASRConfig struct {
	LanguageHints []*string `json:"LanguageHints,omitempty" xml:"LanguageHints,omitempty" type:"Repeated"`
	// example:
	//
	// 800
	MaxSentenceSilence *int32 `json:"MaxSentenceSilence,omitempty" xml:"MaxSentenceSilence,omitempty"`
	// example:
	//
	// false
	SemanticPunctuationEnabled *bool `json:"SemanticPunctuationEnabled,omitempty" xml:"SemanticPunctuationEnabled,omitempty"`
	// example:
	//
	// zh
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
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

func (s *StartAgentRequestVoiceChatConfigASRConfig) SetVocabularyId(v string) *StartAgentRequestVoiceChatConfigASRConfig {
	s.VocabularyId = &v
	return s
}

func (s *StartAgentRequestVoiceChatConfigASRConfig) Validate() error {
	return dara.Validate(s)
}

type StartAgentRequestVoiceChatConfigLLMConfig struct {
	// example:
	//
	// xxxxxxxxxxx
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// example:
	//
	// 3
	HistoryDepth *int32 `json:"HistoryDepth,omitempty" xml:"HistoryDepth,omitempty"`
	// example:
	//
	// 500
	MaxToken *int32 `json:"MaxToken,omitempty" xml:"MaxToken,omitempty"`
	// example:
	//
	// qwen-plus
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// example:
	//
	// xxxx
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// example:
	//
	// 0.7
	Temperature *float64 `json:"Temperature,omitempty" xml:"Temperature,omitempty"`
	// example:
	//
	// 0.8
	TopP *float64 `json:"TopP,omitempty" xml:"TopP,omitempty"`
	// example:
	//
	// https://xxxxx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
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

func (s *StartAgentRequestVoiceChatConfigLLMConfig) GetHistoryDepth() *int32 {
	return s.HistoryDepth
}

func (s *StartAgentRequestVoiceChatConfigLLMConfig) GetMaxToken() *int32 {
	return s.MaxToken
}

func (s *StartAgentRequestVoiceChatConfigLLMConfig) GetModel() *string {
	return s.Model
}

func (s *StartAgentRequestVoiceChatConfigLLMConfig) GetPrompt() *string {
	return s.Prompt
}

func (s *StartAgentRequestVoiceChatConfigLLMConfig) GetTemperature() *float64 {
	return s.Temperature
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

func (s *StartAgentRequestVoiceChatConfigLLMConfig) SetPrompt(v string) *StartAgentRequestVoiceChatConfigLLMConfig {
	s.Prompt = &v
	return s
}

func (s *StartAgentRequestVoiceChatConfigLLMConfig) SetTemperature(v float64) *StartAgentRequestVoiceChatConfigLLMConfig {
	s.Temperature = &v
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
	// example:
	//
	// xxxxxx
	ApiKey         *string  `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	FilterBrackets []*int32 `json:"FilterBrackets,omitempty" xml:"FilterBrackets,omitempty" type:"Repeated"`
	// example:
	//
	// cosyvoice-v1
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// example:
	//
	// 1
	Pitch *float64 `json:"Pitch,omitempty" xml:"Pitch,omitempty"`
	// example:
	//
	// 1
	Rate *float64 `json:"Rate,omitempty" xml:"Rate,omitempty"`
	// example:
	//
	// aliyun
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
	// example:
	//
	// longxiaoxia
	Voice *string `json:"Voice,omitempty" xml:"Voice,omitempty"`
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
