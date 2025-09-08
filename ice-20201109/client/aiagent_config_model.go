// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAIAgentConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAmbientSoundConfig(v *AIAgentConfigAmbientSoundConfig) *AIAgentConfig
	GetAmbientSoundConfig() *AIAgentConfigAmbientSoundConfig
	SetAsrConfig(v *AIAgentConfigAsrConfig) *AIAgentConfig
	GetAsrConfig() *AIAgentConfigAsrConfig
	SetAvatarConfig(v *AIAgentConfigAvatarConfig) *AIAgentConfig
	GetAvatarConfig() *AIAgentConfigAvatarConfig
	SetAvatarUrl(v string) *AIAgentConfig
	GetAvatarUrl() *string
	SetAvatarUrlType(v string) *AIAgentConfig
	GetAvatarUrlType() *string
	SetEnableIntelligentSegment(v bool) *AIAgentConfig
	GetEnableIntelligentSegment() *bool
	SetEnablePushToTalk(v bool) *AIAgentConfig
	GetEnablePushToTalk() *bool
	SetExperimentalConfig(v string) *AIAgentConfig
	GetExperimentalConfig() *string
	SetGracefulShutdown(v bool) *AIAgentConfig
	GetGracefulShutdown() *bool
	SetGreeting(v string) *AIAgentConfig
	GetGreeting() *string
	SetInterruptConfig(v *AIAgentConfigInterruptConfig) *AIAgentConfig
	GetInterruptConfig() *AIAgentConfigInterruptConfig
	SetLlmConfig(v *AIAgentConfigLlmConfig) *AIAgentConfig
	GetLlmConfig() *AIAgentConfigLlmConfig
	SetMaxIdleTime(v int32) *AIAgentConfig
	GetMaxIdleTime() *int32
	SetTtsConfig(v *AIAgentConfigTtsConfig) *AIAgentConfig
	GetTtsConfig() *AIAgentConfigTtsConfig
	SetTurnDetectionConfig(v *AIAgentConfigTurnDetectionConfig) *AIAgentConfig
	GetTurnDetectionConfig() *AIAgentConfigTurnDetectionConfig
	SetUserOfflineTimeout(v int32) *AIAgentConfig
	GetUserOfflineTimeout() *int32
	SetUserOnlineTimeout(v int32) *AIAgentConfig
	GetUserOnlineTimeout() *int32
	SetVcrConfig(v *AIAgentConfigVcrConfig) *AIAgentConfig
	GetVcrConfig() *AIAgentConfigVcrConfig
	SetVoiceprintConfig(v *AIAgentConfigVoiceprintConfig) *AIAgentConfig
	GetVoiceprintConfig() *AIAgentConfigVoiceprintConfig
	SetVolume(v int64) *AIAgentConfig
	GetVolume() *int64
	SetWakeUpQuery(v string) *AIAgentConfig
	GetWakeUpQuery() *string
	SetWorkflowOverrideParams(v string) *AIAgentConfig
	GetWorkflowOverrideParams() *string
}

type AIAgentConfig struct {
	AmbientSoundConfig       *AIAgentConfigAmbientSoundConfig  `json:"AmbientSoundConfig,omitempty" xml:"AmbientSoundConfig,omitempty" type:"Struct"`
	AsrConfig                *AIAgentConfigAsrConfig           `json:"AsrConfig,omitempty" xml:"AsrConfig,omitempty" type:"Struct"`
	AvatarConfig             *AIAgentConfigAvatarConfig        `json:"AvatarConfig,omitempty" xml:"AvatarConfig,omitempty" type:"Struct"`
	AvatarUrl                *string                           `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	AvatarUrlType            *string                           `json:"AvatarUrlType,omitempty" xml:"AvatarUrlType,omitempty"`
	EnableIntelligentSegment *bool                             `json:"EnableIntelligentSegment,omitempty" xml:"EnableIntelligentSegment,omitempty"`
	EnablePushToTalk         *bool                             `json:"EnablePushToTalk,omitempty" xml:"EnablePushToTalk,omitempty"`
	ExperimentalConfig       *string                           `json:"ExperimentalConfig,omitempty" xml:"ExperimentalConfig,omitempty"`
	GracefulShutdown         *bool                             `json:"GracefulShutdown,omitempty" xml:"GracefulShutdown,omitempty"`
	Greeting                 *string                           `json:"Greeting,omitempty" xml:"Greeting,omitempty"`
	InterruptConfig          *AIAgentConfigInterruptConfig     `json:"InterruptConfig,omitempty" xml:"InterruptConfig,omitempty" type:"Struct"`
	LlmConfig                *AIAgentConfigLlmConfig           `json:"LlmConfig,omitempty" xml:"LlmConfig,omitempty" type:"Struct"`
	MaxIdleTime              *int32                            `json:"MaxIdleTime,omitempty" xml:"MaxIdleTime,omitempty"`
	TtsConfig                *AIAgentConfigTtsConfig           `json:"TtsConfig,omitempty" xml:"TtsConfig,omitempty" type:"Struct"`
	TurnDetectionConfig      *AIAgentConfigTurnDetectionConfig `json:"TurnDetectionConfig,omitempty" xml:"TurnDetectionConfig,omitempty" type:"Struct"`
	UserOfflineTimeout       *int32                            `json:"UserOfflineTimeout,omitempty" xml:"UserOfflineTimeout,omitempty"`
	UserOnlineTimeout        *int32                            `json:"UserOnlineTimeout,omitempty" xml:"UserOnlineTimeout,omitempty"`
	VcrConfig                *AIAgentConfigVcrConfig           `json:"VcrConfig,omitempty" xml:"VcrConfig,omitempty" type:"Struct"`
	VoiceprintConfig         *AIAgentConfigVoiceprintConfig    `json:"VoiceprintConfig,omitempty" xml:"VoiceprintConfig,omitempty" type:"Struct"`
	Volume                   *int64                            `json:"Volume,omitempty" xml:"Volume,omitempty"`
	WakeUpQuery              *string                           `json:"WakeUpQuery,omitempty" xml:"WakeUpQuery,omitempty"`
	WorkflowOverrideParams   *string                           `json:"WorkflowOverrideParams,omitempty" xml:"WorkflowOverrideParams,omitempty"`
}

func (s AIAgentConfig) String() string {
	return dara.Prettify(s)
}

func (s AIAgentConfig) GoString() string {
	return s.String()
}

func (s *AIAgentConfig) GetAmbientSoundConfig() *AIAgentConfigAmbientSoundConfig {
	return s.AmbientSoundConfig
}

func (s *AIAgentConfig) GetAsrConfig() *AIAgentConfigAsrConfig {
	return s.AsrConfig
}

func (s *AIAgentConfig) GetAvatarConfig() *AIAgentConfigAvatarConfig {
	return s.AvatarConfig
}

func (s *AIAgentConfig) GetAvatarUrl() *string {
	return s.AvatarUrl
}

func (s *AIAgentConfig) GetAvatarUrlType() *string {
	return s.AvatarUrlType
}

func (s *AIAgentConfig) GetEnableIntelligentSegment() *bool {
	return s.EnableIntelligentSegment
}

func (s *AIAgentConfig) GetEnablePushToTalk() *bool {
	return s.EnablePushToTalk
}

func (s *AIAgentConfig) GetExperimentalConfig() *string {
	return s.ExperimentalConfig
}

func (s *AIAgentConfig) GetGracefulShutdown() *bool {
	return s.GracefulShutdown
}

func (s *AIAgentConfig) GetGreeting() *string {
	return s.Greeting
}

func (s *AIAgentConfig) GetInterruptConfig() *AIAgentConfigInterruptConfig {
	return s.InterruptConfig
}

func (s *AIAgentConfig) GetLlmConfig() *AIAgentConfigLlmConfig {
	return s.LlmConfig
}

func (s *AIAgentConfig) GetMaxIdleTime() *int32 {
	return s.MaxIdleTime
}

func (s *AIAgentConfig) GetTtsConfig() *AIAgentConfigTtsConfig {
	return s.TtsConfig
}

func (s *AIAgentConfig) GetTurnDetectionConfig() *AIAgentConfigTurnDetectionConfig {
	return s.TurnDetectionConfig
}

func (s *AIAgentConfig) GetUserOfflineTimeout() *int32 {
	return s.UserOfflineTimeout
}

func (s *AIAgentConfig) GetUserOnlineTimeout() *int32 {
	return s.UserOnlineTimeout
}

func (s *AIAgentConfig) GetVcrConfig() *AIAgentConfigVcrConfig {
	return s.VcrConfig
}

func (s *AIAgentConfig) GetVoiceprintConfig() *AIAgentConfigVoiceprintConfig {
	return s.VoiceprintConfig
}

func (s *AIAgentConfig) GetVolume() *int64 {
	return s.Volume
}

func (s *AIAgentConfig) GetWakeUpQuery() *string {
	return s.WakeUpQuery
}

func (s *AIAgentConfig) GetWorkflowOverrideParams() *string {
	return s.WorkflowOverrideParams
}

func (s *AIAgentConfig) SetAmbientSoundConfig(v *AIAgentConfigAmbientSoundConfig) *AIAgentConfig {
	s.AmbientSoundConfig = v
	return s
}

func (s *AIAgentConfig) SetAsrConfig(v *AIAgentConfigAsrConfig) *AIAgentConfig {
	s.AsrConfig = v
	return s
}

func (s *AIAgentConfig) SetAvatarConfig(v *AIAgentConfigAvatarConfig) *AIAgentConfig {
	s.AvatarConfig = v
	return s
}

func (s *AIAgentConfig) SetAvatarUrl(v string) *AIAgentConfig {
	s.AvatarUrl = &v
	return s
}

func (s *AIAgentConfig) SetAvatarUrlType(v string) *AIAgentConfig {
	s.AvatarUrlType = &v
	return s
}

func (s *AIAgentConfig) SetEnableIntelligentSegment(v bool) *AIAgentConfig {
	s.EnableIntelligentSegment = &v
	return s
}

func (s *AIAgentConfig) SetEnablePushToTalk(v bool) *AIAgentConfig {
	s.EnablePushToTalk = &v
	return s
}

func (s *AIAgentConfig) SetExperimentalConfig(v string) *AIAgentConfig {
	s.ExperimentalConfig = &v
	return s
}

func (s *AIAgentConfig) SetGracefulShutdown(v bool) *AIAgentConfig {
	s.GracefulShutdown = &v
	return s
}

func (s *AIAgentConfig) SetGreeting(v string) *AIAgentConfig {
	s.Greeting = &v
	return s
}

func (s *AIAgentConfig) SetInterruptConfig(v *AIAgentConfigInterruptConfig) *AIAgentConfig {
	s.InterruptConfig = v
	return s
}

func (s *AIAgentConfig) SetLlmConfig(v *AIAgentConfigLlmConfig) *AIAgentConfig {
	s.LlmConfig = v
	return s
}

func (s *AIAgentConfig) SetMaxIdleTime(v int32) *AIAgentConfig {
	s.MaxIdleTime = &v
	return s
}

func (s *AIAgentConfig) SetTtsConfig(v *AIAgentConfigTtsConfig) *AIAgentConfig {
	s.TtsConfig = v
	return s
}

func (s *AIAgentConfig) SetTurnDetectionConfig(v *AIAgentConfigTurnDetectionConfig) *AIAgentConfig {
	s.TurnDetectionConfig = v
	return s
}

func (s *AIAgentConfig) SetUserOfflineTimeout(v int32) *AIAgentConfig {
	s.UserOfflineTimeout = &v
	return s
}

func (s *AIAgentConfig) SetUserOnlineTimeout(v int32) *AIAgentConfig {
	s.UserOnlineTimeout = &v
	return s
}

func (s *AIAgentConfig) SetVcrConfig(v *AIAgentConfigVcrConfig) *AIAgentConfig {
	s.VcrConfig = v
	return s
}

func (s *AIAgentConfig) SetVoiceprintConfig(v *AIAgentConfigVoiceprintConfig) *AIAgentConfig {
	s.VoiceprintConfig = v
	return s
}

func (s *AIAgentConfig) SetVolume(v int64) *AIAgentConfig {
	s.Volume = &v
	return s
}

func (s *AIAgentConfig) SetWakeUpQuery(v string) *AIAgentConfig {
	s.WakeUpQuery = &v
	return s
}

func (s *AIAgentConfig) SetWorkflowOverrideParams(v string) *AIAgentConfig {
	s.WorkflowOverrideParams = &v
	return s
}

func (s *AIAgentConfig) Validate() error {
	return dara.Validate(s)
}

type AIAgentConfigAmbientSoundConfig struct {
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	Volume     *int32  `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s AIAgentConfigAmbientSoundConfig) String() string {
	return dara.Prettify(s)
}

func (s AIAgentConfigAmbientSoundConfig) GoString() string {
	return s.String()
}

func (s *AIAgentConfigAmbientSoundConfig) GetResourceId() *string {
	return s.ResourceId
}

func (s *AIAgentConfigAmbientSoundConfig) GetVolume() *int32 {
	return s.Volume
}

func (s *AIAgentConfigAmbientSoundConfig) SetResourceId(v string) *AIAgentConfigAmbientSoundConfig {
	s.ResourceId = &v
	return s
}

func (s *AIAgentConfigAmbientSoundConfig) SetVolume(v int32) *AIAgentConfigAmbientSoundConfig {
	s.Volume = &v
	return s
}

func (s *AIAgentConfigAmbientSoundConfig) Validate() error {
	return dara.Validate(s)
}

type AIAgentConfigAsrConfig struct {
	AsrHotWords   []*string `json:"AsrHotWords,omitempty" xml:"AsrHotWords,omitempty" type:"Repeated"`
	AsrLanguageId *string   `json:"AsrLanguageId,omitempty" xml:"AsrLanguageId,omitempty"`
	AsrMaxSilence *int32    `json:"AsrMaxSilence,omitempty" xml:"AsrMaxSilence,omitempty"`
	CustomParams  *string   `json:"CustomParams,omitempty" xml:"CustomParams,omitempty"`
	VadDuration   *int32    `json:"VadDuration,omitempty" xml:"VadDuration,omitempty"`
	VadLevel      *int32    `json:"VadLevel,omitempty" xml:"VadLevel,omitempty"`
}

func (s AIAgentConfigAsrConfig) String() string {
	return dara.Prettify(s)
}

func (s AIAgentConfigAsrConfig) GoString() string {
	return s.String()
}

func (s *AIAgentConfigAsrConfig) GetAsrHotWords() []*string {
	return s.AsrHotWords
}

func (s *AIAgentConfigAsrConfig) GetAsrLanguageId() *string {
	return s.AsrLanguageId
}

func (s *AIAgentConfigAsrConfig) GetAsrMaxSilence() *int32 {
	return s.AsrMaxSilence
}

func (s *AIAgentConfigAsrConfig) GetCustomParams() *string {
	return s.CustomParams
}

func (s *AIAgentConfigAsrConfig) GetVadDuration() *int32 {
	return s.VadDuration
}

func (s *AIAgentConfigAsrConfig) GetVadLevel() *int32 {
	return s.VadLevel
}

func (s *AIAgentConfigAsrConfig) SetAsrHotWords(v []*string) *AIAgentConfigAsrConfig {
	s.AsrHotWords = v
	return s
}

func (s *AIAgentConfigAsrConfig) SetAsrLanguageId(v string) *AIAgentConfigAsrConfig {
	s.AsrLanguageId = &v
	return s
}

func (s *AIAgentConfigAsrConfig) SetAsrMaxSilence(v int32) *AIAgentConfigAsrConfig {
	s.AsrMaxSilence = &v
	return s
}

func (s *AIAgentConfigAsrConfig) SetCustomParams(v string) *AIAgentConfigAsrConfig {
	s.CustomParams = &v
	return s
}

func (s *AIAgentConfigAsrConfig) SetVadDuration(v int32) *AIAgentConfigAsrConfig {
	s.VadDuration = &v
	return s
}

func (s *AIAgentConfigAsrConfig) SetVadLevel(v int32) *AIAgentConfigAsrConfig {
	s.VadLevel = &v
	return s
}

func (s *AIAgentConfigAsrConfig) Validate() error {
	return dara.Validate(s)
}

type AIAgentConfigAvatarConfig struct {
	AvatarId *string `json:"AvatarId,omitempty" xml:"AvatarId,omitempty"`
}

func (s AIAgentConfigAvatarConfig) String() string {
	return dara.Prettify(s)
}

func (s AIAgentConfigAvatarConfig) GoString() string {
	return s.String()
}

func (s *AIAgentConfigAvatarConfig) GetAvatarId() *string {
	return s.AvatarId
}

func (s *AIAgentConfigAvatarConfig) SetAvatarId(v string) *AIAgentConfigAvatarConfig {
	s.AvatarId = &v
	return s
}

func (s *AIAgentConfigAvatarConfig) Validate() error {
	return dara.Validate(s)
}

type AIAgentConfigInterruptConfig struct {
	EnableVoiceInterrupt *bool     `json:"EnableVoiceInterrupt,omitempty" xml:"EnableVoiceInterrupt,omitempty"`
	InterruptWords       []*string `json:"InterruptWords,omitempty" xml:"InterruptWords,omitempty" type:"Repeated"`
}

func (s AIAgentConfigInterruptConfig) String() string {
	return dara.Prettify(s)
}

func (s AIAgentConfigInterruptConfig) GoString() string {
	return s.String()
}

func (s *AIAgentConfigInterruptConfig) GetEnableVoiceInterrupt() *bool {
	return s.EnableVoiceInterrupt
}

func (s *AIAgentConfigInterruptConfig) GetInterruptWords() []*string {
	return s.InterruptWords
}

func (s *AIAgentConfigInterruptConfig) SetEnableVoiceInterrupt(v bool) *AIAgentConfigInterruptConfig {
	s.EnableVoiceInterrupt = &v
	return s
}

func (s *AIAgentConfigInterruptConfig) SetInterruptWords(v []*string) *AIAgentConfigInterruptConfig {
	s.InterruptWords = v
	return s
}

func (s *AIAgentConfigInterruptConfig) Validate() error {
	return dara.Validate(s)
}

type AIAgentConfigLlmConfig struct {
	BailianAppParams *string                              `json:"BailianAppParams,omitempty" xml:"BailianAppParams,omitempty"`
	FunctionMap      []*AIAgentConfigLlmConfigFunctionMap `json:"FunctionMap,omitempty" xml:"FunctionMap,omitempty" type:"Repeated"`
	LlmCompleteReply *bool                                `json:"LlmCompleteReply,omitempty" xml:"LlmCompleteReply,omitempty"`
	LlmHistory       []*AIAgentConfigLlmConfigLlmHistory  `json:"LlmHistory,omitempty" xml:"LlmHistory,omitempty" type:"Repeated"`
	LlmHistoryLimit  *int32                               `json:"LlmHistoryLimit,omitempty" xml:"LlmHistoryLimit,omitempty"`
	LlmSystemPrompt  *string                              `json:"LlmSystemPrompt,omitempty" xml:"LlmSystemPrompt,omitempty"`
	OpenAIExtraQuery *string                              `json:"OpenAIExtraQuery,omitempty" xml:"OpenAIExtraQuery,omitempty"`
	OutputMaxDelay   *int32                               `json:"OutputMaxDelay,omitempty" xml:"OutputMaxDelay,omitempty"`
	OutputMinLength  *int32                               `json:"OutputMinLength,omitempty" xml:"OutputMinLength,omitempty"`
}

func (s AIAgentConfigLlmConfig) String() string {
	return dara.Prettify(s)
}

func (s AIAgentConfigLlmConfig) GoString() string {
	return s.String()
}

func (s *AIAgentConfigLlmConfig) GetBailianAppParams() *string {
	return s.BailianAppParams
}

func (s *AIAgentConfigLlmConfig) GetFunctionMap() []*AIAgentConfigLlmConfigFunctionMap {
	return s.FunctionMap
}

func (s *AIAgentConfigLlmConfig) GetLlmCompleteReply() *bool {
	return s.LlmCompleteReply
}

func (s *AIAgentConfigLlmConfig) GetLlmHistory() []*AIAgentConfigLlmConfigLlmHistory {
	return s.LlmHistory
}

func (s *AIAgentConfigLlmConfig) GetLlmHistoryLimit() *int32 {
	return s.LlmHistoryLimit
}

func (s *AIAgentConfigLlmConfig) GetLlmSystemPrompt() *string {
	return s.LlmSystemPrompt
}

func (s *AIAgentConfigLlmConfig) GetOpenAIExtraQuery() *string {
	return s.OpenAIExtraQuery
}

func (s *AIAgentConfigLlmConfig) GetOutputMaxDelay() *int32 {
	return s.OutputMaxDelay
}

func (s *AIAgentConfigLlmConfig) GetOutputMinLength() *int32 {
	return s.OutputMinLength
}

func (s *AIAgentConfigLlmConfig) SetBailianAppParams(v string) *AIAgentConfigLlmConfig {
	s.BailianAppParams = &v
	return s
}

func (s *AIAgentConfigLlmConfig) SetFunctionMap(v []*AIAgentConfigLlmConfigFunctionMap) *AIAgentConfigLlmConfig {
	s.FunctionMap = v
	return s
}

func (s *AIAgentConfigLlmConfig) SetLlmCompleteReply(v bool) *AIAgentConfigLlmConfig {
	s.LlmCompleteReply = &v
	return s
}

func (s *AIAgentConfigLlmConfig) SetLlmHistory(v []*AIAgentConfigLlmConfigLlmHistory) *AIAgentConfigLlmConfig {
	s.LlmHistory = v
	return s
}

func (s *AIAgentConfigLlmConfig) SetLlmHistoryLimit(v int32) *AIAgentConfigLlmConfig {
	s.LlmHistoryLimit = &v
	return s
}

func (s *AIAgentConfigLlmConfig) SetLlmSystemPrompt(v string) *AIAgentConfigLlmConfig {
	s.LlmSystemPrompt = &v
	return s
}

func (s *AIAgentConfigLlmConfig) SetOpenAIExtraQuery(v string) *AIAgentConfigLlmConfig {
	s.OpenAIExtraQuery = &v
	return s
}

func (s *AIAgentConfigLlmConfig) SetOutputMaxDelay(v int32) *AIAgentConfigLlmConfig {
	s.OutputMaxDelay = &v
	return s
}

func (s *AIAgentConfigLlmConfig) SetOutputMinLength(v int32) *AIAgentConfigLlmConfig {
	s.OutputMinLength = &v
	return s
}

func (s *AIAgentConfigLlmConfig) Validate() error {
	return dara.Validate(s)
}

type AIAgentConfigLlmConfigFunctionMap struct {
	Function      *string `json:"Function,omitempty" xml:"Function,omitempty"`
	MatchFunction *string `json:"MatchFunction,omitempty" xml:"MatchFunction,omitempty"`
}

func (s AIAgentConfigLlmConfigFunctionMap) String() string {
	return dara.Prettify(s)
}

func (s AIAgentConfigLlmConfigFunctionMap) GoString() string {
	return s.String()
}

func (s *AIAgentConfigLlmConfigFunctionMap) GetFunction() *string {
	return s.Function
}

func (s *AIAgentConfigLlmConfigFunctionMap) GetMatchFunction() *string {
	return s.MatchFunction
}

func (s *AIAgentConfigLlmConfigFunctionMap) SetFunction(v string) *AIAgentConfigLlmConfigFunctionMap {
	s.Function = &v
	return s
}

func (s *AIAgentConfigLlmConfigFunctionMap) SetMatchFunction(v string) *AIAgentConfigLlmConfigFunctionMap {
	s.MatchFunction = &v
	return s
}

func (s *AIAgentConfigLlmConfigFunctionMap) Validate() error {
	return dara.Validate(s)
}

type AIAgentConfigLlmConfigLlmHistory struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Role    *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s AIAgentConfigLlmConfigLlmHistory) String() string {
	return dara.Prettify(s)
}

func (s AIAgentConfigLlmConfigLlmHistory) GoString() string {
	return s.String()
}

func (s *AIAgentConfigLlmConfigLlmHistory) GetContent() *string {
	return s.Content
}

func (s *AIAgentConfigLlmConfigLlmHistory) GetRole() *string {
	return s.Role
}

func (s *AIAgentConfigLlmConfigLlmHistory) SetContent(v string) *AIAgentConfigLlmConfigLlmHistory {
	s.Content = &v
	return s
}

func (s *AIAgentConfigLlmConfigLlmHistory) SetRole(v string) *AIAgentConfigLlmConfigLlmHistory {
	s.Role = &v
	return s
}

func (s *AIAgentConfigLlmConfigLlmHistory) Validate() error {
	return dara.Validate(s)
}

type AIAgentConfigTtsConfig struct {
	Emotion            *string                                     `json:"Emotion,omitempty" xml:"Emotion,omitempty"`
	LanguageId         *string                                     `json:"LanguageId,omitempty" xml:"LanguageId,omitempty"`
	ModelId            *string                                     `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	PronunciationRules []*AIAgentConfigTtsConfigPronunciationRules `json:"PronunciationRules,omitempty" xml:"PronunciationRules,omitempty" type:"Repeated"`
	SpeechRate         *float64                                    `json:"SpeechRate,omitempty" xml:"SpeechRate,omitempty"`
	VoiceId            *string                                     `json:"VoiceId,omitempty" xml:"VoiceId,omitempty"`
	VoiceIdList        []*string                                   `json:"VoiceIdList,omitempty" xml:"VoiceIdList,omitempty" type:"Repeated"`
}

func (s AIAgentConfigTtsConfig) String() string {
	return dara.Prettify(s)
}

func (s AIAgentConfigTtsConfig) GoString() string {
	return s.String()
}

func (s *AIAgentConfigTtsConfig) GetEmotion() *string {
	return s.Emotion
}

func (s *AIAgentConfigTtsConfig) GetLanguageId() *string {
	return s.LanguageId
}

func (s *AIAgentConfigTtsConfig) GetModelId() *string {
	return s.ModelId
}

func (s *AIAgentConfigTtsConfig) GetPronunciationRules() []*AIAgentConfigTtsConfigPronunciationRules {
	return s.PronunciationRules
}

func (s *AIAgentConfigTtsConfig) GetSpeechRate() *float64 {
	return s.SpeechRate
}

func (s *AIAgentConfigTtsConfig) GetVoiceId() *string {
	return s.VoiceId
}

func (s *AIAgentConfigTtsConfig) GetVoiceIdList() []*string {
	return s.VoiceIdList
}

func (s *AIAgentConfigTtsConfig) SetEmotion(v string) *AIAgentConfigTtsConfig {
	s.Emotion = &v
	return s
}

func (s *AIAgentConfigTtsConfig) SetLanguageId(v string) *AIAgentConfigTtsConfig {
	s.LanguageId = &v
	return s
}

func (s *AIAgentConfigTtsConfig) SetModelId(v string) *AIAgentConfigTtsConfig {
	s.ModelId = &v
	return s
}

func (s *AIAgentConfigTtsConfig) SetPronunciationRules(v []*AIAgentConfigTtsConfigPronunciationRules) *AIAgentConfigTtsConfig {
	s.PronunciationRules = v
	return s
}

func (s *AIAgentConfigTtsConfig) SetSpeechRate(v float64) *AIAgentConfigTtsConfig {
	s.SpeechRate = &v
	return s
}

func (s *AIAgentConfigTtsConfig) SetVoiceId(v string) *AIAgentConfigTtsConfig {
	s.VoiceId = &v
	return s
}

func (s *AIAgentConfigTtsConfig) SetVoiceIdList(v []*string) *AIAgentConfigTtsConfig {
	s.VoiceIdList = v
	return s
}

func (s *AIAgentConfigTtsConfig) Validate() error {
	return dara.Validate(s)
}

type AIAgentConfigTtsConfigPronunciationRules struct {
	Pronunciation *string `json:"Pronunciation,omitempty" xml:"Pronunciation,omitempty"`
	Type          *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Word          *string `json:"Word,omitempty" xml:"Word,omitempty"`
}

func (s AIAgentConfigTtsConfigPronunciationRules) String() string {
	return dara.Prettify(s)
}

func (s AIAgentConfigTtsConfigPronunciationRules) GoString() string {
	return s.String()
}

func (s *AIAgentConfigTtsConfigPronunciationRules) GetPronunciation() *string {
	return s.Pronunciation
}

func (s *AIAgentConfigTtsConfigPronunciationRules) GetType() *string {
	return s.Type
}

func (s *AIAgentConfigTtsConfigPronunciationRules) GetWord() *string {
	return s.Word
}

func (s *AIAgentConfigTtsConfigPronunciationRules) SetPronunciation(v string) *AIAgentConfigTtsConfigPronunciationRules {
	s.Pronunciation = &v
	return s
}

func (s *AIAgentConfigTtsConfigPronunciationRules) SetType(v string) *AIAgentConfigTtsConfigPronunciationRules {
	s.Type = &v
	return s
}

func (s *AIAgentConfigTtsConfigPronunciationRules) SetWord(v string) *AIAgentConfigTtsConfigPronunciationRules {
	s.Word = &v
	return s
}

func (s *AIAgentConfigTtsConfigPronunciationRules) Validate() error {
	return dara.Validate(s)
}

type AIAgentConfigTurnDetectionConfig struct {
	Mode                 *string   `json:"Mode,omitempty" xml:"Mode,omitempty"`
	SemanticWaitDuration *int32    `json:"SemanticWaitDuration,omitempty" xml:"SemanticWaitDuration,omitempty"`
	TurnEndWords         []*string `json:"TurnEndWords,omitempty" xml:"TurnEndWords,omitempty" type:"Repeated"`
}

func (s AIAgentConfigTurnDetectionConfig) String() string {
	return dara.Prettify(s)
}

func (s AIAgentConfigTurnDetectionConfig) GoString() string {
	return s.String()
}

func (s *AIAgentConfigTurnDetectionConfig) GetMode() *string {
	return s.Mode
}

func (s *AIAgentConfigTurnDetectionConfig) GetSemanticWaitDuration() *int32 {
	return s.SemanticWaitDuration
}

func (s *AIAgentConfigTurnDetectionConfig) GetTurnEndWords() []*string {
	return s.TurnEndWords
}

func (s *AIAgentConfigTurnDetectionConfig) SetMode(v string) *AIAgentConfigTurnDetectionConfig {
	s.Mode = &v
	return s
}

func (s *AIAgentConfigTurnDetectionConfig) SetSemanticWaitDuration(v int32) *AIAgentConfigTurnDetectionConfig {
	s.SemanticWaitDuration = &v
	return s
}

func (s *AIAgentConfigTurnDetectionConfig) SetTurnEndWords(v []*string) *AIAgentConfigTurnDetectionConfig {
	s.TurnEndWords = v
	return s
}

func (s *AIAgentConfigTurnDetectionConfig) Validate() error {
	return dara.Validate(s)
}

type AIAgentConfigVcrConfig struct {
	Equipment          *AIAgentConfigVcrConfigEquipment          `json:"Equipment,omitempty" xml:"Equipment,omitempty" type:"Struct"`
	HeadMotion         *AIAgentConfigVcrConfigHeadMotion         `json:"HeadMotion,omitempty" xml:"HeadMotion,omitempty" type:"Struct"`
	InvalidFrameMotion *AIAgentConfigVcrConfigInvalidFrameMotion `json:"InvalidFrameMotion,omitempty" xml:"InvalidFrameMotion,omitempty" type:"Struct"`
	LookAway           *AIAgentConfigVcrConfigLookAway           `json:"LookAway,omitempty" xml:"LookAway,omitempty" type:"Struct"`
	PeopleCount        *AIAgentConfigVcrConfigPeopleCount        `json:"PeopleCount,omitempty" xml:"PeopleCount,omitempty" type:"Struct"`
	StillFrameMotion   *AIAgentConfigVcrConfigStillFrameMotion   `json:"StillFrameMotion,omitempty" xml:"StillFrameMotion,omitempty" type:"Struct"`
}

func (s AIAgentConfigVcrConfig) String() string {
	return dara.Prettify(s)
}

func (s AIAgentConfigVcrConfig) GoString() string {
	return s.String()
}

func (s *AIAgentConfigVcrConfig) GetEquipment() *AIAgentConfigVcrConfigEquipment {
	return s.Equipment
}

func (s *AIAgentConfigVcrConfig) GetHeadMotion() *AIAgentConfigVcrConfigHeadMotion {
	return s.HeadMotion
}

func (s *AIAgentConfigVcrConfig) GetInvalidFrameMotion() *AIAgentConfigVcrConfigInvalidFrameMotion {
	return s.InvalidFrameMotion
}

func (s *AIAgentConfigVcrConfig) GetLookAway() *AIAgentConfigVcrConfigLookAway {
	return s.LookAway
}

func (s *AIAgentConfigVcrConfig) GetPeopleCount() *AIAgentConfigVcrConfigPeopleCount {
	return s.PeopleCount
}

func (s *AIAgentConfigVcrConfig) GetStillFrameMotion() *AIAgentConfigVcrConfigStillFrameMotion {
	return s.StillFrameMotion
}

func (s *AIAgentConfigVcrConfig) SetEquipment(v *AIAgentConfigVcrConfigEquipment) *AIAgentConfigVcrConfig {
	s.Equipment = v
	return s
}

func (s *AIAgentConfigVcrConfig) SetHeadMotion(v *AIAgentConfigVcrConfigHeadMotion) *AIAgentConfigVcrConfig {
	s.HeadMotion = v
	return s
}

func (s *AIAgentConfigVcrConfig) SetInvalidFrameMotion(v *AIAgentConfigVcrConfigInvalidFrameMotion) *AIAgentConfigVcrConfig {
	s.InvalidFrameMotion = v
	return s
}

func (s *AIAgentConfigVcrConfig) SetLookAway(v *AIAgentConfigVcrConfigLookAway) *AIAgentConfigVcrConfig {
	s.LookAway = v
	return s
}

func (s *AIAgentConfigVcrConfig) SetPeopleCount(v *AIAgentConfigVcrConfigPeopleCount) *AIAgentConfigVcrConfig {
	s.PeopleCount = v
	return s
}

func (s *AIAgentConfigVcrConfig) SetStillFrameMotion(v *AIAgentConfigVcrConfigStillFrameMotion) *AIAgentConfigVcrConfig {
	s.StillFrameMotion = v
	return s
}

func (s *AIAgentConfigVcrConfig) Validate() error {
	return dara.Validate(s)
}

type AIAgentConfigVcrConfigEquipment struct {
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
}

func (s AIAgentConfigVcrConfigEquipment) String() string {
	return dara.Prettify(s)
}

func (s AIAgentConfigVcrConfigEquipment) GoString() string {
	return s.String()
}

func (s *AIAgentConfigVcrConfigEquipment) GetEnabled() *bool {
	return s.Enabled
}

func (s *AIAgentConfigVcrConfigEquipment) SetEnabled(v bool) *AIAgentConfigVcrConfigEquipment {
	s.Enabled = &v
	return s
}

func (s *AIAgentConfigVcrConfigEquipment) Validate() error {
	return dara.Validate(s)
}

type AIAgentConfigVcrConfigHeadMotion struct {
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
}

func (s AIAgentConfigVcrConfigHeadMotion) String() string {
	return dara.Prettify(s)
}

func (s AIAgentConfigVcrConfigHeadMotion) GoString() string {
	return s.String()
}

func (s *AIAgentConfigVcrConfigHeadMotion) GetEnabled() *bool {
	return s.Enabled
}

func (s *AIAgentConfigVcrConfigHeadMotion) SetEnabled(v bool) *AIAgentConfigVcrConfigHeadMotion {
	s.Enabled = &v
	return s
}

func (s *AIAgentConfigVcrConfigHeadMotion) Validate() error {
	return dara.Validate(s)
}

type AIAgentConfigVcrConfigInvalidFrameMotion struct {
	CallbackDelay *int32 `json:"CallbackDelay,omitempty" xml:"CallbackDelay,omitempty"`
	Enabled       *bool  `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
}

func (s AIAgentConfigVcrConfigInvalidFrameMotion) String() string {
	return dara.Prettify(s)
}

func (s AIAgentConfigVcrConfigInvalidFrameMotion) GoString() string {
	return s.String()
}

func (s *AIAgentConfigVcrConfigInvalidFrameMotion) GetCallbackDelay() *int32 {
	return s.CallbackDelay
}

func (s *AIAgentConfigVcrConfigInvalidFrameMotion) GetEnabled() *bool {
	return s.Enabled
}

func (s *AIAgentConfigVcrConfigInvalidFrameMotion) SetCallbackDelay(v int32) *AIAgentConfigVcrConfigInvalidFrameMotion {
	s.CallbackDelay = &v
	return s
}

func (s *AIAgentConfigVcrConfigInvalidFrameMotion) SetEnabled(v bool) *AIAgentConfigVcrConfigInvalidFrameMotion {
	s.Enabled = &v
	return s
}

func (s *AIAgentConfigVcrConfigInvalidFrameMotion) Validate() error {
	return dara.Validate(s)
}

type AIAgentConfigVcrConfigLookAway struct {
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
}

func (s AIAgentConfigVcrConfigLookAway) String() string {
	return dara.Prettify(s)
}

func (s AIAgentConfigVcrConfigLookAway) GoString() string {
	return s.String()
}

func (s *AIAgentConfigVcrConfigLookAway) GetEnabled() *bool {
	return s.Enabled
}

func (s *AIAgentConfigVcrConfigLookAway) SetEnabled(v bool) *AIAgentConfigVcrConfigLookAway {
	s.Enabled = &v
	return s
}

func (s *AIAgentConfigVcrConfigLookAway) Validate() error {
	return dara.Validate(s)
}

type AIAgentConfigVcrConfigPeopleCount struct {
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
}

func (s AIAgentConfigVcrConfigPeopleCount) String() string {
	return dara.Prettify(s)
}

func (s AIAgentConfigVcrConfigPeopleCount) GoString() string {
	return s.String()
}

func (s *AIAgentConfigVcrConfigPeopleCount) GetEnabled() *bool {
	return s.Enabled
}

func (s *AIAgentConfigVcrConfigPeopleCount) SetEnabled(v bool) *AIAgentConfigVcrConfigPeopleCount {
	s.Enabled = &v
	return s
}

func (s *AIAgentConfigVcrConfigPeopleCount) Validate() error {
	return dara.Validate(s)
}

type AIAgentConfigVcrConfigStillFrameMotion struct {
	CallbackDelay *int32 `json:"CallbackDelay,omitempty" xml:"CallbackDelay,omitempty"`
	Enabled       *bool  `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
}

func (s AIAgentConfigVcrConfigStillFrameMotion) String() string {
	return dara.Prettify(s)
}

func (s AIAgentConfigVcrConfigStillFrameMotion) GoString() string {
	return s.String()
}

func (s *AIAgentConfigVcrConfigStillFrameMotion) GetCallbackDelay() *int32 {
	return s.CallbackDelay
}

func (s *AIAgentConfigVcrConfigStillFrameMotion) GetEnabled() *bool {
	return s.Enabled
}

func (s *AIAgentConfigVcrConfigStillFrameMotion) SetCallbackDelay(v int32) *AIAgentConfigVcrConfigStillFrameMotion {
	s.CallbackDelay = &v
	return s
}

func (s *AIAgentConfigVcrConfigStillFrameMotion) SetEnabled(v bool) *AIAgentConfigVcrConfigStillFrameMotion {
	s.Enabled = &v
	return s
}

func (s *AIAgentConfigVcrConfigStillFrameMotion) Validate() error {
	return dara.Validate(s)
}

type AIAgentConfigVoiceprintConfig struct {
	UseVoiceprint *bool   `json:"UseVoiceprint,omitempty" xml:"UseVoiceprint,omitempty"`
	VoiceprintId  *string `json:"VoiceprintId,omitempty" xml:"VoiceprintId,omitempty"`
}

func (s AIAgentConfigVoiceprintConfig) String() string {
	return dara.Prettify(s)
}

func (s AIAgentConfigVoiceprintConfig) GoString() string {
	return s.String()
}

func (s *AIAgentConfigVoiceprintConfig) GetUseVoiceprint() *bool {
	return s.UseVoiceprint
}

func (s *AIAgentConfigVoiceprintConfig) GetVoiceprintId() *string {
	return s.VoiceprintId
}

func (s *AIAgentConfigVoiceprintConfig) SetUseVoiceprint(v bool) *AIAgentConfigVoiceprintConfig {
	s.UseVoiceprint = &v
	return s
}

func (s *AIAgentConfigVoiceprintConfig) SetVoiceprintId(v string) *AIAgentConfigVoiceprintConfig {
	s.VoiceprintId = &v
	return s
}

func (s *AIAgentConfigVoiceprintConfig) Validate() error {
	return dara.Validate(s)
}
