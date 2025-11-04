// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAIAgentOutboundCallConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAmbientSoundConfig(v *AIAgentOutboundCallConfigAmbientSoundConfig) *AIAgentOutboundCallConfig
	GetAmbientSoundConfig() *AIAgentOutboundCallConfigAmbientSoundConfig
	SetAsrConfig(v *AIAgentOutboundCallConfigAsrConfig) *AIAgentOutboundCallConfig
	GetAsrConfig() *AIAgentOutboundCallConfigAsrConfig
	SetEnableIntelligentSegment(v bool) *AIAgentOutboundCallConfig
	GetEnableIntelligentSegment() *bool
	SetExperimentalConfig(v string) *AIAgentOutboundCallConfig
	GetExperimentalConfig() *string
	SetGreeting(v string) *AIAgentOutboundCallConfig
	GetGreeting() *string
	SetGreetingDelay(v int32) *AIAgentOutboundCallConfig
	GetGreetingDelay() *int32
	SetInterruptConfig(v *AIAgentOutboundCallConfigInterruptConfig) *AIAgentOutboundCallConfig
	GetInterruptConfig() *AIAgentOutboundCallConfigInterruptConfig
	SetLlmConfig(v *AIAgentOutboundCallConfigLlmConfig) *AIAgentOutboundCallConfig
	GetLlmConfig() *AIAgentOutboundCallConfigLlmConfig
	SetTtsConfig(v *AIAgentOutboundCallConfigTtsConfig) *AIAgentOutboundCallConfig
	GetTtsConfig() *AIAgentOutboundCallConfigTtsConfig
	SetTurnDetectionConfig(v *AIAgentOutboundCallConfigTurnDetectionConfig) *AIAgentOutboundCallConfig
	GetTurnDetectionConfig() *AIAgentOutboundCallConfigTurnDetectionConfig
}

type AIAgentOutboundCallConfig struct {
	AmbientSoundConfig       *AIAgentOutboundCallConfigAmbientSoundConfig  `json:"AmbientSoundConfig,omitempty" xml:"AmbientSoundConfig,omitempty" type:"Struct"`
	AsrConfig                *AIAgentOutboundCallConfigAsrConfig           `json:"AsrConfig,omitempty" xml:"AsrConfig,omitempty" type:"Struct"`
	EnableIntelligentSegment *bool                                         `json:"EnableIntelligentSegment,omitempty" xml:"EnableIntelligentSegment,omitempty"`
	ExperimentalConfig       *string                                       `json:"ExperimentalConfig,omitempty" xml:"ExperimentalConfig,omitempty"`
	Greeting                 *string                                       `json:"Greeting,omitempty" xml:"Greeting,omitempty"`
	GreetingDelay            *int32                                        `json:"GreetingDelay,omitempty" xml:"GreetingDelay,omitempty"`
	InterruptConfig          *AIAgentOutboundCallConfigInterruptConfig     `json:"InterruptConfig,omitempty" xml:"InterruptConfig,omitempty" type:"Struct"`
	LlmConfig                *AIAgentOutboundCallConfigLlmConfig           `json:"LlmConfig,omitempty" xml:"LlmConfig,omitempty" type:"Struct"`
	TtsConfig                *AIAgentOutboundCallConfigTtsConfig           `json:"TtsConfig,omitempty" xml:"TtsConfig,omitempty" type:"Struct"`
	TurnDetectionConfig      *AIAgentOutboundCallConfigTurnDetectionConfig `json:"TurnDetectionConfig,omitempty" xml:"TurnDetectionConfig,omitempty" type:"Struct"`
}

func (s AIAgentOutboundCallConfig) String() string {
	return dara.Prettify(s)
}

func (s AIAgentOutboundCallConfig) GoString() string {
	return s.String()
}

func (s *AIAgentOutboundCallConfig) GetAmbientSoundConfig() *AIAgentOutboundCallConfigAmbientSoundConfig {
	return s.AmbientSoundConfig
}

func (s *AIAgentOutboundCallConfig) GetAsrConfig() *AIAgentOutboundCallConfigAsrConfig {
	return s.AsrConfig
}

func (s *AIAgentOutboundCallConfig) GetEnableIntelligentSegment() *bool {
	return s.EnableIntelligentSegment
}

func (s *AIAgentOutboundCallConfig) GetExperimentalConfig() *string {
	return s.ExperimentalConfig
}

func (s *AIAgentOutboundCallConfig) GetGreeting() *string {
	return s.Greeting
}

func (s *AIAgentOutboundCallConfig) GetGreetingDelay() *int32 {
	return s.GreetingDelay
}

func (s *AIAgentOutboundCallConfig) GetInterruptConfig() *AIAgentOutboundCallConfigInterruptConfig {
	return s.InterruptConfig
}

func (s *AIAgentOutboundCallConfig) GetLlmConfig() *AIAgentOutboundCallConfigLlmConfig {
	return s.LlmConfig
}

func (s *AIAgentOutboundCallConfig) GetTtsConfig() *AIAgentOutboundCallConfigTtsConfig {
	return s.TtsConfig
}

func (s *AIAgentOutboundCallConfig) GetTurnDetectionConfig() *AIAgentOutboundCallConfigTurnDetectionConfig {
	return s.TurnDetectionConfig
}

func (s *AIAgentOutboundCallConfig) SetAmbientSoundConfig(v *AIAgentOutboundCallConfigAmbientSoundConfig) *AIAgentOutboundCallConfig {
	s.AmbientSoundConfig = v
	return s
}

func (s *AIAgentOutboundCallConfig) SetAsrConfig(v *AIAgentOutboundCallConfigAsrConfig) *AIAgentOutboundCallConfig {
	s.AsrConfig = v
	return s
}

func (s *AIAgentOutboundCallConfig) SetEnableIntelligentSegment(v bool) *AIAgentOutboundCallConfig {
	s.EnableIntelligentSegment = &v
	return s
}

func (s *AIAgentOutboundCallConfig) SetExperimentalConfig(v string) *AIAgentOutboundCallConfig {
	s.ExperimentalConfig = &v
	return s
}

func (s *AIAgentOutboundCallConfig) SetGreeting(v string) *AIAgentOutboundCallConfig {
	s.Greeting = &v
	return s
}

func (s *AIAgentOutboundCallConfig) SetGreetingDelay(v int32) *AIAgentOutboundCallConfig {
	s.GreetingDelay = &v
	return s
}

func (s *AIAgentOutboundCallConfig) SetInterruptConfig(v *AIAgentOutboundCallConfigInterruptConfig) *AIAgentOutboundCallConfig {
	s.InterruptConfig = v
	return s
}

func (s *AIAgentOutboundCallConfig) SetLlmConfig(v *AIAgentOutboundCallConfigLlmConfig) *AIAgentOutboundCallConfig {
	s.LlmConfig = v
	return s
}

func (s *AIAgentOutboundCallConfig) SetTtsConfig(v *AIAgentOutboundCallConfigTtsConfig) *AIAgentOutboundCallConfig {
	s.TtsConfig = v
	return s
}

func (s *AIAgentOutboundCallConfig) SetTurnDetectionConfig(v *AIAgentOutboundCallConfigTurnDetectionConfig) *AIAgentOutboundCallConfig {
	s.TurnDetectionConfig = v
	return s
}

func (s *AIAgentOutboundCallConfig) Validate() error {
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
	if s.TurnDetectionConfig != nil {
		if err := s.TurnDetectionConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AIAgentOutboundCallConfigAmbientSoundConfig struct {
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	Volume     *int32  `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s AIAgentOutboundCallConfigAmbientSoundConfig) String() string {
	return dara.Prettify(s)
}

func (s AIAgentOutboundCallConfigAmbientSoundConfig) GoString() string {
	return s.String()
}

func (s *AIAgentOutboundCallConfigAmbientSoundConfig) GetResourceId() *string {
	return s.ResourceId
}

func (s *AIAgentOutboundCallConfigAmbientSoundConfig) GetVolume() *int32 {
	return s.Volume
}

func (s *AIAgentOutboundCallConfigAmbientSoundConfig) SetResourceId(v string) *AIAgentOutboundCallConfigAmbientSoundConfig {
	s.ResourceId = &v
	return s
}

func (s *AIAgentOutboundCallConfigAmbientSoundConfig) SetVolume(v int32) *AIAgentOutboundCallConfigAmbientSoundConfig {
	s.Volume = &v
	return s
}

func (s *AIAgentOutboundCallConfigAmbientSoundConfig) Validate() error {
	return dara.Validate(s)
}

type AIAgentOutboundCallConfigAsrConfig struct {
	AsrHotWords   []*string `json:"AsrHotWords,omitempty" xml:"AsrHotWords,omitempty" type:"Repeated"`
	AsrLanguageId *string   `json:"AsrLanguageId,omitempty" xml:"AsrLanguageId,omitempty"`
	AsrMaxSilence *int32    `json:"AsrMaxSilence,omitempty" xml:"AsrMaxSilence,omitempty"`
	CustomParams  *string   `json:"CustomParams,omitempty" xml:"CustomParams,omitempty"`
	VadDuration   *int32    `json:"VadDuration,omitempty" xml:"VadDuration,omitempty"`
	VadLevel      *int32    `json:"VadLevel,omitempty" xml:"VadLevel,omitempty"`
}

func (s AIAgentOutboundCallConfigAsrConfig) String() string {
	return dara.Prettify(s)
}

func (s AIAgentOutboundCallConfigAsrConfig) GoString() string {
	return s.String()
}

func (s *AIAgentOutboundCallConfigAsrConfig) GetAsrHotWords() []*string {
	return s.AsrHotWords
}

func (s *AIAgentOutboundCallConfigAsrConfig) GetAsrLanguageId() *string {
	return s.AsrLanguageId
}

func (s *AIAgentOutboundCallConfigAsrConfig) GetAsrMaxSilence() *int32 {
	return s.AsrMaxSilence
}

func (s *AIAgentOutboundCallConfigAsrConfig) GetCustomParams() *string {
	return s.CustomParams
}

func (s *AIAgentOutboundCallConfigAsrConfig) GetVadDuration() *int32 {
	return s.VadDuration
}

func (s *AIAgentOutboundCallConfigAsrConfig) GetVadLevel() *int32 {
	return s.VadLevel
}

func (s *AIAgentOutboundCallConfigAsrConfig) SetAsrHotWords(v []*string) *AIAgentOutboundCallConfigAsrConfig {
	s.AsrHotWords = v
	return s
}

func (s *AIAgentOutboundCallConfigAsrConfig) SetAsrLanguageId(v string) *AIAgentOutboundCallConfigAsrConfig {
	s.AsrLanguageId = &v
	return s
}

func (s *AIAgentOutboundCallConfigAsrConfig) SetAsrMaxSilence(v int32) *AIAgentOutboundCallConfigAsrConfig {
	s.AsrMaxSilence = &v
	return s
}

func (s *AIAgentOutboundCallConfigAsrConfig) SetCustomParams(v string) *AIAgentOutboundCallConfigAsrConfig {
	s.CustomParams = &v
	return s
}

func (s *AIAgentOutboundCallConfigAsrConfig) SetVadDuration(v int32) *AIAgentOutboundCallConfigAsrConfig {
	s.VadDuration = &v
	return s
}

func (s *AIAgentOutboundCallConfigAsrConfig) SetVadLevel(v int32) *AIAgentOutboundCallConfigAsrConfig {
	s.VadLevel = &v
	return s
}

func (s *AIAgentOutboundCallConfigAsrConfig) Validate() error {
	return dara.Validate(s)
}

type AIAgentOutboundCallConfigInterruptConfig struct {
	EnableVoiceInterrupt *bool     `json:"EnableVoiceInterrupt,omitempty" xml:"EnableVoiceInterrupt,omitempty"`
	InterruptWords       []*string `json:"InterruptWords,omitempty" xml:"InterruptWords,omitempty" type:"Repeated"`
}

func (s AIAgentOutboundCallConfigInterruptConfig) String() string {
	return dara.Prettify(s)
}

func (s AIAgentOutboundCallConfigInterruptConfig) GoString() string {
	return s.String()
}

func (s *AIAgentOutboundCallConfigInterruptConfig) GetEnableVoiceInterrupt() *bool {
	return s.EnableVoiceInterrupt
}

func (s *AIAgentOutboundCallConfigInterruptConfig) GetInterruptWords() []*string {
	return s.InterruptWords
}

func (s *AIAgentOutboundCallConfigInterruptConfig) SetEnableVoiceInterrupt(v bool) *AIAgentOutboundCallConfigInterruptConfig {
	s.EnableVoiceInterrupt = &v
	return s
}

func (s *AIAgentOutboundCallConfigInterruptConfig) SetInterruptWords(v []*string) *AIAgentOutboundCallConfigInterruptConfig {
	s.InterruptWords = v
	return s
}

func (s *AIAgentOutboundCallConfigInterruptConfig) Validate() error {
	return dara.Validate(s)
}

type AIAgentOutboundCallConfigLlmConfig struct {
	BailianAppParams *string                                          `json:"BailianAppParams,omitempty" xml:"BailianAppParams,omitempty"`
	FunctionMap      []*AIAgentOutboundCallConfigLlmConfigFunctionMap `json:"FunctionMap,omitempty" xml:"FunctionMap,omitempty" type:"Repeated"`
	LlmCompleteReply *bool                                            `json:"LlmCompleteReply,omitempty" xml:"LlmCompleteReply,omitempty"`
	LlmHistory       []*AIAgentOutboundCallConfigLlmConfigLlmHistory  `json:"LlmHistory,omitempty" xml:"LlmHistory,omitempty" type:"Repeated"`
	LlmHistoryLimit  *int32                                           `json:"LlmHistoryLimit,omitempty" xml:"LlmHistoryLimit,omitempty"`
	LlmSystemPrompt  *string                                          `json:"LlmSystemPrompt,omitempty" xml:"LlmSystemPrompt,omitempty"`
	OpenAIExtraQuery *string                                          `json:"OpenAIExtraQuery,omitempty" xml:"OpenAIExtraQuery,omitempty"`
	OutputMaxDelay   *string                                          `json:"OutputMaxDelay,omitempty" xml:"OutputMaxDelay,omitempty"`
	OutputMinLength  *int32                                           `json:"OutputMinLength,omitempty" xml:"OutputMinLength,omitempty"`
}

func (s AIAgentOutboundCallConfigLlmConfig) String() string {
	return dara.Prettify(s)
}

func (s AIAgentOutboundCallConfigLlmConfig) GoString() string {
	return s.String()
}

func (s *AIAgentOutboundCallConfigLlmConfig) GetBailianAppParams() *string {
	return s.BailianAppParams
}

func (s *AIAgentOutboundCallConfigLlmConfig) GetFunctionMap() []*AIAgentOutboundCallConfigLlmConfigFunctionMap {
	return s.FunctionMap
}

func (s *AIAgentOutboundCallConfigLlmConfig) GetLlmCompleteReply() *bool {
	return s.LlmCompleteReply
}

func (s *AIAgentOutboundCallConfigLlmConfig) GetLlmHistory() []*AIAgentOutboundCallConfigLlmConfigLlmHistory {
	return s.LlmHistory
}

func (s *AIAgentOutboundCallConfigLlmConfig) GetLlmHistoryLimit() *int32 {
	return s.LlmHistoryLimit
}

func (s *AIAgentOutboundCallConfigLlmConfig) GetLlmSystemPrompt() *string {
	return s.LlmSystemPrompt
}

func (s *AIAgentOutboundCallConfigLlmConfig) GetOpenAIExtraQuery() *string {
	return s.OpenAIExtraQuery
}

func (s *AIAgentOutboundCallConfigLlmConfig) GetOutputMaxDelay() *string {
	return s.OutputMaxDelay
}

func (s *AIAgentOutboundCallConfigLlmConfig) GetOutputMinLength() *int32 {
	return s.OutputMinLength
}

func (s *AIAgentOutboundCallConfigLlmConfig) SetBailianAppParams(v string) *AIAgentOutboundCallConfigLlmConfig {
	s.BailianAppParams = &v
	return s
}

func (s *AIAgentOutboundCallConfigLlmConfig) SetFunctionMap(v []*AIAgentOutboundCallConfigLlmConfigFunctionMap) *AIAgentOutboundCallConfigLlmConfig {
	s.FunctionMap = v
	return s
}

func (s *AIAgentOutboundCallConfigLlmConfig) SetLlmCompleteReply(v bool) *AIAgentOutboundCallConfigLlmConfig {
	s.LlmCompleteReply = &v
	return s
}

func (s *AIAgentOutboundCallConfigLlmConfig) SetLlmHistory(v []*AIAgentOutboundCallConfigLlmConfigLlmHistory) *AIAgentOutboundCallConfigLlmConfig {
	s.LlmHistory = v
	return s
}

func (s *AIAgentOutboundCallConfigLlmConfig) SetLlmHistoryLimit(v int32) *AIAgentOutboundCallConfigLlmConfig {
	s.LlmHistoryLimit = &v
	return s
}

func (s *AIAgentOutboundCallConfigLlmConfig) SetLlmSystemPrompt(v string) *AIAgentOutboundCallConfigLlmConfig {
	s.LlmSystemPrompt = &v
	return s
}

func (s *AIAgentOutboundCallConfigLlmConfig) SetOpenAIExtraQuery(v string) *AIAgentOutboundCallConfigLlmConfig {
	s.OpenAIExtraQuery = &v
	return s
}

func (s *AIAgentOutboundCallConfigLlmConfig) SetOutputMaxDelay(v string) *AIAgentOutboundCallConfigLlmConfig {
	s.OutputMaxDelay = &v
	return s
}

func (s *AIAgentOutboundCallConfigLlmConfig) SetOutputMinLength(v int32) *AIAgentOutboundCallConfigLlmConfig {
	s.OutputMinLength = &v
	return s
}

func (s *AIAgentOutboundCallConfigLlmConfig) Validate() error {
	if s.FunctionMap != nil {
		for _, item := range s.FunctionMap {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.LlmHistory != nil {
		for _, item := range s.LlmHistory {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AIAgentOutboundCallConfigLlmConfigFunctionMap struct {
	Function      *string `json:"Function,omitempty" xml:"Function,omitempty"`
	MatchFunction *string `json:"MatchFunction,omitempty" xml:"MatchFunction,omitempty"`
}

func (s AIAgentOutboundCallConfigLlmConfigFunctionMap) String() string {
	return dara.Prettify(s)
}

func (s AIAgentOutboundCallConfigLlmConfigFunctionMap) GoString() string {
	return s.String()
}

func (s *AIAgentOutboundCallConfigLlmConfigFunctionMap) GetFunction() *string {
	return s.Function
}

func (s *AIAgentOutboundCallConfigLlmConfigFunctionMap) GetMatchFunction() *string {
	return s.MatchFunction
}

func (s *AIAgentOutboundCallConfigLlmConfigFunctionMap) SetFunction(v string) *AIAgentOutboundCallConfigLlmConfigFunctionMap {
	s.Function = &v
	return s
}

func (s *AIAgentOutboundCallConfigLlmConfigFunctionMap) SetMatchFunction(v string) *AIAgentOutboundCallConfigLlmConfigFunctionMap {
	s.MatchFunction = &v
	return s
}

func (s *AIAgentOutboundCallConfigLlmConfigFunctionMap) Validate() error {
	return dara.Validate(s)
}

type AIAgentOutboundCallConfigLlmConfigLlmHistory struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Role    *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s AIAgentOutboundCallConfigLlmConfigLlmHistory) String() string {
	return dara.Prettify(s)
}

func (s AIAgentOutboundCallConfigLlmConfigLlmHistory) GoString() string {
	return s.String()
}

func (s *AIAgentOutboundCallConfigLlmConfigLlmHistory) GetContent() *string {
	return s.Content
}

func (s *AIAgentOutboundCallConfigLlmConfigLlmHistory) GetRole() *string {
	return s.Role
}

func (s *AIAgentOutboundCallConfigLlmConfigLlmHistory) SetContent(v string) *AIAgentOutboundCallConfigLlmConfigLlmHistory {
	s.Content = &v
	return s
}

func (s *AIAgentOutboundCallConfigLlmConfigLlmHistory) SetRole(v string) *AIAgentOutboundCallConfigLlmConfigLlmHistory {
	s.Role = &v
	return s
}

func (s *AIAgentOutboundCallConfigLlmConfigLlmHistory) Validate() error {
	return dara.Validate(s)
}

type AIAgentOutboundCallConfigTtsConfig struct {
	Emotion            *string                                                 `json:"Emotion,omitempty" xml:"Emotion,omitempty"`
	LanguageId         *string                                                 `json:"LanguageId,omitempty" xml:"LanguageId,omitempty"`
	ModelId            *string                                                 `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	PronunciationRules []*AIAgentOutboundCallConfigTtsConfigPronunciationRules `json:"PronunciationRules,omitempty" xml:"PronunciationRules,omitempty" type:"Repeated"`
	SpeechRate         *float64                                                `json:"SpeechRate,omitempty" xml:"SpeechRate,omitempty"`
	VoiceId            *string                                                 `json:"VoiceId,omitempty" xml:"VoiceId,omitempty"`
	VoiceIdList        []*string                                               `json:"VoiceIdList,omitempty" xml:"VoiceIdList,omitempty" type:"Repeated"`
}

func (s AIAgentOutboundCallConfigTtsConfig) String() string {
	return dara.Prettify(s)
}

func (s AIAgentOutboundCallConfigTtsConfig) GoString() string {
	return s.String()
}

func (s *AIAgentOutboundCallConfigTtsConfig) GetEmotion() *string {
	return s.Emotion
}

func (s *AIAgentOutboundCallConfigTtsConfig) GetLanguageId() *string {
	return s.LanguageId
}

func (s *AIAgentOutboundCallConfigTtsConfig) GetModelId() *string {
	return s.ModelId
}

func (s *AIAgentOutboundCallConfigTtsConfig) GetPronunciationRules() []*AIAgentOutboundCallConfigTtsConfigPronunciationRules {
	return s.PronunciationRules
}

func (s *AIAgentOutboundCallConfigTtsConfig) GetSpeechRate() *float64 {
	return s.SpeechRate
}

func (s *AIAgentOutboundCallConfigTtsConfig) GetVoiceId() *string {
	return s.VoiceId
}

func (s *AIAgentOutboundCallConfigTtsConfig) GetVoiceIdList() []*string {
	return s.VoiceIdList
}

func (s *AIAgentOutboundCallConfigTtsConfig) SetEmotion(v string) *AIAgentOutboundCallConfigTtsConfig {
	s.Emotion = &v
	return s
}

func (s *AIAgentOutboundCallConfigTtsConfig) SetLanguageId(v string) *AIAgentOutboundCallConfigTtsConfig {
	s.LanguageId = &v
	return s
}

func (s *AIAgentOutboundCallConfigTtsConfig) SetModelId(v string) *AIAgentOutboundCallConfigTtsConfig {
	s.ModelId = &v
	return s
}

func (s *AIAgentOutboundCallConfigTtsConfig) SetPronunciationRules(v []*AIAgentOutboundCallConfigTtsConfigPronunciationRules) *AIAgentOutboundCallConfigTtsConfig {
	s.PronunciationRules = v
	return s
}

func (s *AIAgentOutboundCallConfigTtsConfig) SetSpeechRate(v float64) *AIAgentOutboundCallConfigTtsConfig {
	s.SpeechRate = &v
	return s
}

func (s *AIAgentOutboundCallConfigTtsConfig) SetVoiceId(v string) *AIAgentOutboundCallConfigTtsConfig {
	s.VoiceId = &v
	return s
}

func (s *AIAgentOutboundCallConfigTtsConfig) SetVoiceIdList(v []*string) *AIAgentOutboundCallConfigTtsConfig {
	s.VoiceIdList = v
	return s
}

func (s *AIAgentOutboundCallConfigTtsConfig) Validate() error {
	if s.PronunciationRules != nil {
		for _, item := range s.PronunciationRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AIAgentOutboundCallConfigTtsConfigPronunciationRules struct {
	Pronunciation *string `json:"Pronunciation,omitempty" xml:"Pronunciation,omitempty"`
	Type          *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Word          *string `json:"Word,omitempty" xml:"Word,omitempty"`
}

func (s AIAgentOutboundCallConfigTtsConfigPronunciationRules) String() string {
	return dara.Prettify(s)
}

func (s AIAgentOutboundCallConfigTtsConfigPronunciationRules) GoString() string {
	return s.String()
}

func (s *AIAgentOutboundCallConfigTtsConfigPronunciationRules) GetPronunciation() *string {
	return s.Pronunciation
}

func (s *AIAgentOutboundCallConfigTtsConfigPronunciationRules) GetType() *string {
	return s.Type
}

func (s *AIAgentOutboundCallConfigTtsConfigPronunciationRules) GetWord() *string {
	return s.Word
}

func (s *AIAgentOutboundCallConfigTtsConfigPronunciationRules) SetPronunciation(v string) *AIAgentOutboundCallConfigTtsConfigPronunciationRules {
	s.Pronunciation = &v
	return s
}

func (s *AIAgentOutboundCallConfigTtsConfigPronunciationRules) SetType(v string) *AIAgentOutboundCallConfigTtsConfigPronunciationRules {
	s.Type = &v
	return s
}

func (s *AIAgentOutboundCallConfigTtsConfigPronunciationRules) SetWord(v string) *AIAgentOutboundCallConfigTtsConfigPronunciationRules {
	s.Word = &v
	return s
}

func (s *AIAgentOutboundCallConfigTtsConfigPronunciationRules) Validate() error {
	return dara.Validate(s)
}

type AIAgentOutboundCallConfigTurnDetectionConfig struct {
	Mode                 *string   `json:"Mode,omitempty" xml:"Mode,omitempty"`
	SemanticWaitDuration *int32    `json:"SemanticWaitDuration,omitempty" xml:"SemanticWaitDuration,omitempty"`
	TurnEndWords         []*string `json:"TurnEndWords,omitempty" xml:"TurnEndWords,omitempty" type:"Repeated"`
}

func (s AIAgentOutboundCallConfigTurnDetectionConfig) String() string {
	return dara.Prettify(s)
}

func (s AIAgentOutboundCallConfigTurnDetectionConfig) GoString() string {
	return s.String()
}

func (s *AIAgentOutboundCallConfigTurnDetectionConfig) GetMode() *string {
	return s.Mode
}

func (s *AIAgentOutboundCallConfigTurnDetectionConfig) GetSemanticWaitDuration() *int32 {
	return s.SemanticWaitDuration
}

func (s *AIAgentOutboundCallConfigTurnDetectionConfig) GetTurnEndWords() []*string {
	return s.TurnEndWords
}

func (s *AIAgentOutboundCallConfigTurnDetectionConfig) SetMode(v string) *AIAgentOutboundCallConfigTurnDetectionConfig {
	s.Mode = &v
	return s
}

func (s *AIAgentOutboundCallConfigTurnDetectionConfig) SetSemanticWaitDuration(v int32) *AIAgentOutboundCallConfigTurnDetectionConfig {
	s.SemanticWaitDuration = &v
	return s
}

func (s *AIAgentOutboundCallConfigTurnDetectionConfig) SetTurnEndWords(v []*string) *AIAgentOutboundCallConfigTurnDetectionConfig {
	s.TurnEndWords = v
	return s
}

func (s *AIAgentOutboundCallConfigTurnDetectionConfig) Validate() error {
	return dara.Validate(s)
}
