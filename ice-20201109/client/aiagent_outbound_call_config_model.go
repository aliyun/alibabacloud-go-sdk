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
	SetAutoSpeechConfig(v *AIAgentOutboundCallConfigAutoSpeechConfig) *AIAgentOutboundCallConfig
	GetAutoSpeechConfig() *AIAgentOutboundCallConfigAutoSpeechConfig
	SetBackChannelingConfig(v *AIAgentOutboundCallConfigBackChannelingConfig) *AIAgentOutboundCallConfig
	GetBackChannelingConfig() *AIAgentOutboundCallConfigBackChannelingConfig
	SetBackChannelingConfigs(v []*AIAgentOutboundCallConfigBackChannelingConfigs) *AIAgentOutboundCallConfig
	GetBackChannelingConfigs() []*AIAgentOutboundCallConfigBackChannelingConfigs
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
	SetMaxIdleTime(v int32) *AIAgentOutboundCallConfig
	GetMaxIdleTime() *int32
	SetTtsConfig(v *AIAgentOutboundCallConfigTtsConfig) *AIAgentOutboundCallConfig
	GetTtsConfig() *AIAgentOutboundCallConfigTtsConfig
	SetTurnDetectionConfig(v *AIAgentOutboundCallConfigTurnDetectionConfig) *AIAgentOutboundCallConfig
	GetTurnDetectionConfig() *AIAgentOutboundCallConfigTurnDetectionConfig
}

type AIAgentOutboundCallConfig struct {
	// Configuration for the ambient sound played during the call.
	AmbientSoundConfig *AIAgentOutboundCallConfigAmbientSoundConfig `json:"AmbientSoundConfig,omitempty" xml:"AmbientSoundConfig,omitempty" type:"Struct"`
	// The configuration for Automatic Speech Recognition (ASR).
	AsrConfig        *AIAgentOutboundCallConfigAsrConfig        `json:"AsrConfig,omitempty" xml:"AsrConfig,omitempty" type:"Struct"`
	AutoSpeechConfig *AIAgentOutboundCallConfigAutoSpeechConfig `json:"AutoSpeechConfig,omitempty" xml:"AutoSpeechConfig,omitempty" type:"Struct"`
	// Deprecated
	BackChannelingConfig  *AIAgentOutboundCallConfigBackChannelingConfig    `json:"BackChannelingConfig,omitempty" xml:"BackChannelingConfig,omitempty" type:"Struct"`
	BackChannelingConfigs []*AIAgentOutboundCallConfigBackChannelingConfigs `json:"BackChannelingConfigs,omitempty" xml:"BackChannelingConfigs,omitempty" type:"Repeated"`
	// If enabled, the system intelligently merges short, interim segments into a single sentence. Default value: true.
	//
	// example:
	//
	// true
	EnableIntelligentSegment *bool   `json:"EnableIntelligentSegment,omitempty" xml:"EnableIntelligentSegment,omitempty"`
	ExperimentalConfig       *string `json:"ExperimentalConfig,omitempty" xml:"ExperimentalConfig,omitempty"`
	// The welcome message that the agent says upon joining. Changes take effect in the next session. Default value: None.
	Greeting *string `json:"Greeting,omitempty" xml:"Greeting,omitempty"`
	// The delay before playing the welcome message. Unit: milliseconds. Valid values: 0 to 5000. Default value: 0.
	//
	// example:
	//
	// 0
	GreetingDelay *int32 `json:"GreetingDelay,omitempty" xml:"GreetingDelay,omitempty"`
	// The configuration for the speech interruption strategy.
	InterruptConfig *AIAgentOutboundCallConfigInterruptConfig `json:"InterruptConfig,omitempty" xml:"InterruptConfig,omitempty" type:"Struct"`
	// The configuration for the large language model (LLM).
	LlmConfig   *AIAgentOutboundCallConfigLlmConfig `json:"LlmConfig,omitempty" xml:"LlmConfig,omitempty" type:"Struct"`
	MaxIdleTime *int32                              `json:"MaxIdleTime,omitempty" xml:"MaxIdleTime,omitempty"`
	// The configuration for Text-to-Speech (TTS).
	TtsConfig *AIAgentOutboundCallConfigTtsConfig `json:"TtsConfig,omitempty" xml:"TtsConfig,omitempty" type:"Struct"`
	// The configuration for detecting the end of a user\\"s conversational turn.
	TurnDetectionConfig *AIAgentOutboundCallConfigTurnDetectionConfig `json:"TurnDetectionConfig,omitempty" xml:"TurnDetectionConfig,omitempty" type:"Struct"`
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

func (s *AIAgentOutboundCallConfig) GetAutoSpeechConfig() *AIAgentOutboundCallConfigAutoSpeechConfig {
	return s.AutoSpeechConfig
}

func (s *AIAgentOutboundCallConfig) GetBackChannelingConfig() *AIAgentOutboundCallConfigBackChannelingConfig {
	return s.BackChannelingConfig
}

func (s *AIAgentOutboundCallConfig) GetBackChannelingConfigs() []*AIAgentOutboundCallConfigBackChannelingConfigs {
	return s.BackChannelingConfigs
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

func (s *AIAgentOutboundCallConfig) GetMaxIdleTime() *int32 {
	return s.MaxIdleTime
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

func (s *AIAgentOutboundCallConfig) SetAutoSpeechConfig(v *AIAgentOutboundCallConfigAutoSpeechConfig) *AIAgentOutboundCallConfig {
	s.AutoSpeechConfig = v
	return s
}

func (s *AIAgentOutboundCallConfig) SetBackChannelingConfig(v *AIAgentOutboundCallConfigBackChannelingConfig) *AIAgentOutboundCallConfig {
	s.BackChannelingConfig = v
	return s
}

func (s *AIAgentOutboundCallConfig) SetBackChannelingConfigs(v []*AIAgentOutboundCallConfigBackChannelingConfigs) *AIAgentOutboundCallConfig {
	s.BackChannelingConfigs = v
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

func (s *AIAgentOutboundCallConfig) SetMaxIdleTime(v int32) *AIAgentOutboundCallConfig {
	s.MaxIdleTime = &v
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
	if s.AutoSpeechConfig != nil {
		if err := s.AutoSpeechConfig.Validate(); err != nil {
			return err
		}
	}
	if s.BackChannelingConfig != nil {
		if err := s.BackChannelingConfig.Validate(); err != nil {
			return err
		}
	}
	if s.BackChannelingConfigs != nil {
		for _, item := range s.BackChannelingConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
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
	// The ID of the ambient sound. This ID can be obtained from the advanced settings section of the agent configuration in the console.
	//
	// example:
	//
	// f67901c595834************
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The volume of the ambient sound. Valid values: [0, 100]. A value of 0 disables the ambient sound.
	//
	// example:
	//
	// 50
	Volume *int32 `json:"Volume,omitempty" xml:"Volume,omitempty"`
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
	// Hotwords for ASR to improve recognition accuracy. Maximum of 128 hotwords.
	AsrHotWords []*string `json:"AsrHotWords,omitempty" xml:"AsrHotWords,omitempty" type:"Repeated"`
	// The language ID for ASR. Valid values:
	//
	// 	- zh_mandarin: Chinese
	//
	// 	- en: English
	//
	// 	- zh_en: Chinese and English
	//
	// 	- es: Spanish
	//
	// 	- jp: Japanese
	//
	// example:
	//
	// zh_mandarin
	AsrLanguageId *string `json:"AsrLanguageId,omitempty" xml:"AsrLanguageId,omitempty"`
	// The silence threshold for sentence segmentation. A pause longer than this value is considered a sentence break. Unit: milliseconds. Default value: 400. Valid values: 200 to 1200.
	//
	// example:
	//
	// 400
	AsrMaxSilence *int32 `json:"AsrMaxSilence,omitempty" xml:"AsrMaxSilence,omitempty"`
	// Passthrough parameters for ASR.
	//
	// example:
	//
	// mode=fast&sample=16000&format=wav
	CustomParams *string `json:"CustomParams,omitempty" xml:"CustomParams,omitempty"`
	// The minimum duration for voice activity detection, in milliseconds. This parameter controls the sensitivity of interruptions, preventing the agent from cutting off user speech too early during short pauses. 0: Disables this feature. Valid values: 200 to 2000. Recommended: 200 to 500 ms, which typically corresponds to the length of 1 to 4 words. By default, this parameter is left empty, which indicates the feature is disabled.
	//
	// example:
	//
	// 300
	VadDuration *int32 `json:"VadDuration,omitempty" xml:"VadDuration,omitempty"`
	// The VAD threshold for interruption. A higher value makes it harder to trigger interruptions. Valid values: 0 to 10. Default value: 1. The value of 0 specifies to disable the VAD feature.
	//
	// example:
	//
	// 1
	VadLevel *int32 `json:"VadLevel,omitempty" xml:"VadLevel,omitempty"`
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

type AIAgentOutboundCallConfigAutoSpeechConfig struct {
	LlmPending *AIAgentOutboundCallConfigAutoSpeechConfigLlmPending `json:"LlmPending,omitempty" xml:"LlmPending,omitempty" type:"Struct"`
	UserIdle   *AIAgentOutboundCallConfigAutoSpeechConfigUserIdle   `json:"UserIdle,omitempty" xml:"UserIdle,omitempty" type:"Struct"`
}

func (s AIAgentOutboundCallConfigAutoSpeechConfig) String() string {
	return dara.Prettify(s)
}

func (s AIAgentOutboundCallConfigAutoSpeechConfig) GoString() string {
	return s.String()
}

func (s *AIAgentOutboundCallConfigAutoSpeechConfig) GetLlmPending() *AIAgentOutboundCallConfigAutoSpeechConfigLlmPending {
	return s.LlmPending
}

func (s *AIAgentOutboundCallConfigAutoSpeechConfig) GetUserIdle() *AIAgentOutboundCallConfigAutoSpeechConfigUserIdle {
	return s.UserIdle
}

func (s *AIAgentOutboundCallConfigAutoSpeechConfig) SetLlmPending(v *AIAgentOutboundCallConfigAutoSpeechConfigLlmPending) *AIAgentOutboundCallConfigAutoSpeechConfig {
	s.LlmPending = v
	return s
}

func (s *AIAgentOutboundCallConfigAutoSpeechConfig) SetUserIdle(v *AIAgentOutboundCallConfigAutoSpeechConfigUserIdle) *AIAgentOutboundCallConfigAutoSpeechConfig {
	s.UserIdle = v
	return s
}

func (s *AIAgentOutboundCallConfigAutoSpeechConfig) Validate() error {
	if s.LlmPending != nil {
		if err := s.LlmPending.Validate(); err != nil {
			return err
		}
	}
	if s.UserIdle != nil {
		if err := s.UserIdle.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AIAgentOutboundCallConfigAutoSpeechConfigLlmPending struct {
	Messages []*AIAgentOutboundCallConfigAutoSpeechConfigLlmPendingMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Repeated"`
	WaitTime *int32                                                         `json:"WaitTime,omitempty" xml:"WaitTime,omitempty"`
}

func (s AIAgentOutboundCallConfigAutoSpeechConfigLlmPending) String() string {
	return dara.Prettify(s)
}

func (s AIAgentOutboundCallConfigAutoSpeechConfigLlmPending) GoString() string {
	return s.String()
}

func (s *AIAgentOutboundCallConfigAutoSpeechConfigLlmPending) GetMessages() []*AIAgentOutboundCallConfigAutoSpeechConfigLlmPendingMessages {
	return s.Messages
}

func (s *AIAgentOutboundCallConfigAutoSpeechConfigLlmPending) GetWaitTime() *int32 {
	return s.WaitTime
}

func (s *AIAgentOutboundCallConfigAutoSpeechConfigLlmPending) SetMessages(v []*AIAgentOutboundCallConfigAutoSpeechConfigLlmPendingMessages) *AIAgentOutboundCallConfigAutoSpeechConfigLlmPending {
	s.Messages = v
	return s
}

func (s *AIAgentOutboundCallConfigAutoSpeechConfigLlmPending) SetWaitTime(v int32) *AIAgentOutboundCallConfigAutoSpeechConfigLlmPending {
	s.WaitTime = &v
	return s
}

func (s *AIAgentOutboundCallConfigAutoSpeechConfigLlmPending) Validate() error {
	if s.Messages != nil {
		for _, item := range s.Messages {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AIAgentOutboundCallConfigAutoSpeechConfigLlmPendingMessages struct {
	Probability *float64 `json:"Probability,omitempty" xml:"Probability,omitempty"`
	Text        *string  `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s AIAgentOutboundCallConfigAutoSpeechConfigLlmPendingMessages) String() string {
	return dara.Prettify(s)
}

func (s AIAgentOutboundCallConfigAutoSpeechConfigLlmPendingMessages) GoString() string {
	return s.String()
}

func (s *AIAgentOutboundCallConfigAutoSpeechConfigLlmPendingMessages) GetProbability() *float64 {
	return s.Probability
}

func (s *AIAgentOutboundCallConfigAutoSpeechConfigLlmPendingMessages) GetText() *string {
	return s.Text
}

func (s *AIAgentOutboundCallConfigAutoSpeechConfigLlmPendingMessages) SetProbability(v float64) *AIAgentOutboundCallConfigAutoSpeechConfigLlmPendingMessages {
	s.Probability = &v
	return s
}

func (s *AIAgentOutboundCallConfigAutoSpeechConfigLlmPendingMessages) SetText(v string) *AIAgentOutboundCallConfigAutoSpeechConfigLlmPendingMessages {
	s.Text = &v
	return s
}

func (s *AIAgentOutboundCallConfigAutoSpeechConfigLlmPendingMessages) Validate() error {
	return dara.Validate(s)
}

type AIAgentOutboundCallConfigAutoSpeechConfigUserIdle struct {
	MaxRepeats *int32                                                       `json:"MaxRepeats,omitempty" xml:"MaxRepeats,omitempty"`
	Messages   []*AIAgentOutboundCallConfigAutoSpeechConfigUserIdleMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Repeated"`
	WaitTime   *int32                                                       `json:"WaitTime,omitempty" xml:"WaitTime,omitempty"`
}

func (s AIAgentOutboundCallConfigAutoSpeechConfigUserIdle) String() string {
	return dara.Prettify(s)
}

func (s AIAgentOutboundCallConfigAutoSpeechConfigUserIdle) GoString() string {
	return s.String()
}

func (s *AIAgentOutboundCallConfigAutoSpeechConfigUserIdle) GetMaxRepeats() *int32 {
	return s.MaxRepeats
}

func (s *AIAgentOutboundCallConfigAutoSpeechConfigUserIdle) GetMessages() []*AIAgentOutboundCallConfigAutoSpeechConfigUserIdleMessages {
	return s.Messages
}

func (s *AIAgentOutboundCallConfigAutoSpeechConfigUserIdle) GetWaitTime() *int32 {
	return s.WaitTime
}

func (s *AIAgentOutboundCallConfigAutoSpeechConfigUserIdle) SetMaxRepeats(v int32) *AIAgentOutboundCallConfigAutoSpeechConfigUserIdle {
	s.MaxRepeats = &v
	return s
}

func (s *AIAgentOutboundCallConfigAutoSpeechConfigUserIdle) SetMessages(v []*AIAgentOutboundCallConfigAutoSpeechConfigUserIdleMessages) *AIAgentOutboundCallConfigAutoSpeechConfigUserIdle {
	s.Messages = v
	return s
}

func (s *AIAgentOutboundCallConfigAutoSpeechConfigUserIdle) SetWaitTime(v int32) *AIAgentOutboundCallConfigAutoSpeechConfigUserIdle {
	s.WaitTime = &v
	return s
}

func (s *AIAgentOutboundCallConfigAutoSpeechConfigUserIdle) Validate() error {
	if s.Messages != nil {
		for _, item := range s.Messages {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AIAgentOutboundCallConfigAutoSpeechConfigUserIdleMessages struct {
	Probability *float64 `json:"Probability,omitempty" xml:"Probability,omitempty"`
	Text        *string  `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s AIAgentOutboundCallConfigAutoSpeechConfigUserIdleMessages) String() string {
	return dara.Prettify(s)
}

func (s AIAgentOutboundCallConfigAutoSpeechConfigUserIdleMessages) GoString() string {
	return s.String()
}

func (s *AIAgentOutboundCallConfigAutoSpeechConfigUserIdleMessages) GetProbability() *float64 {
	return s.Probability
}

func (s *AIAgentOutboundCallConfigAutoSpeechConfigUserIdleMessages) GetText() *string {
	return s.Text
}

func (s *AIAgentOutboundCallConfigAutoSpeechConfigUserIdleMessages) SetProbability(v float64) *AIAgentOutboundCallConfigAutoSpeechConfigUserIdleMessages {
	s.Probability = &v
	return s
}

func (s *AIAgentOutboundCallConfigAutoSpeechConfigUserIdleMessages) SetText(v string) *AIAgentOutboundCallConfigAutoSpeechConfigUserIdleMessages {
	s.Text = &v
	return s
}

func (s *AIAgentOutboundCallConfigAutoSpeechConfigUserIdleMessages) Validate() error {
	return dara.Validate(s)
}

type AIAgentOutboundCallConfigBackChannelingConfig struct {
	Enabled      *bool                                               `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	Probability  *float64                                            `json:"Probability,omitempty" xml:"Probability,omitempty"`
	TriggerStage *string                                             `json:"TriggerStage,omitempty" xml:"TriggerStage,omitempty"`
	Words        *AIAgentOutboundCallConfigBackChannelingConfigWords `json:"Words,omitempty" xml:"Words,omitempty" type:"Struct"`
}

func (s AIAgentOutboundCallConfigBackChannelingConfig) String() string {
	return dara.Prettify(s)
}

func (s AIAgentOutboundCallConfigBackChannelingConfig) GoString() string {
	return s.String()
}

func (s *AIAgentOutboundCallConfigBackChannelingConfig) GetEnabled() *bool {
	return s.Enabled
}

func (s *AIAgentOutboundCallConfigBackChannelingConfig) GetProbability() *float64 {
	return s.Probability
}

func (s *AIAgentOutboundCallConfigBackChannelingConfig) GetTriggerStage() *string {
	return s.TriggerStage
}

func (s *AIAgentOutboundCallConfigBackChannelingConfig) GetWords() *AIAgentOutboundCallConfigBackChannelingConfigWords {
	return s.Words
}

func (s *AIAgentOutboundCallConfigBackChannelingConfig) SetEnabled(v bool) *AIAgentOutboundCallConfigBackChannelingConfig {
	s.Enabled = &v
	return s
}

func (s *AIAgentOutboundCallConfigBackChannelingConfig) SetProbability(v float64) *AIAgentOutboundCallConfigBackChannelingConfig {
	s.Probability = &v
	return s
}

func (s *AIAgentOutboundCallConfigBackChannelingConfig) SetTriggerStage(v string) *AIAgentOutboundCallConfigBackChannelingConfig {
	s.TriggerStage = &v
	return s
}

func (s *AIAgentOutboundCallConfigBackChannelingConfig) SetWords(v *AIAgentOutboundCallConfigBackChannelingConfigWords) *AIAgentOutboundCallConfigBackChannelingConfig {
	s.Words = v
	return s
}

func (s *AIAgentOutboundCallConfigBackChannelingConfig) Validate() error {
	if s.Words != nil {
		if err := s.Words.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AIAgentOutboundCallConfigBackChannelingConfigWords struct {
	Probability *float64 `json:"Probability,omitempty" xml:"Probability,omitempty"`
	Text        *string  `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s AIAgentOutboundCallConfigBackChannelingConfigWords) String() string {
	return dara.Prettify(s)
}

func (s AIAgentOutboundCallConfigBackChannelingConfigWords) GoString() string {
	return s.String()
}

func (s *AIAgentOutboundCallConfigBackChannelingConfigWords) GetProbability() *float64 {
	return s.Probability
}

func (s *AIAgentOutboundCallConfigBackChannelingConfigWords) GetText() *string {
	return s.Text
}

func (s *AIAgentOutboundCallConfigBackChannelingConfigWords) SetProbability(v float64) *AIAgentOutboundCallConfigBackChannelingConfigWords {
	s.Probability = &v
	return s
}

func (s *AIAgentOutboundCallConfigBackChannelingConfigWords) SetText(v string) *AIAgentOutboundCallConfigBackChannelingConfigWords {
	s.Text = &v
	return s
}

func (s *AIAgentOutboundCallConfigBackChannelingConfigWords) Validate() error {
	return dara.Validate(s)
}

type AIAgentOutboundCallConfigBackChannelingConfigs struct {
	Enabled      *bool                                                  `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	Probability  *float64                                               `json:"Probability,omitempty" xml:"Probability,omitempty"`
	TriggerStage *string                                                `json:"TriggerStage,omitempty" xml:"TriggerStage,omitempty"`
	Words        []*AIAgentOutboundCallConfigBackChannelingConfigsWords `json:"Words,omitempty" xml:"Words,omitempty" type:"Repeated"`
}

func (s AIAgentOutboundCallConfigBackChannelingConfigs) String() string {
	return dara.Prettify(s)
}

func (s AIAgentOutboundCallConfigBackChannelingConfigs) GoString() string {
	return s.String()
}

func (s *AIAgentOutboundCallConfigBackChannelingConfigs) GetEnabled() *bool {
	return s.Enabled
}

func (s *AIAgentOutboundCallConfigBackChannelingConfigs) GetProbability() *float64 {
	return s.Probability
}

func (s *AIAgentOutboundCallConfigBackChannelingConfigs) GetTriggerStage() *string {
	return s.TriggerStage
}

func (s *AIAgentOutboundCallConfigBackChannelingConfigs) GetWords() []*AIAgentOutboundCallConfigBackChannelingConfigsWords {
	return s.Words
}

func (s *AIAgentOutboundCallConfigBackChannelingConfigs) SetEnabled(v bool) *AIAgentOutboundCallConfigBackChannelingConfigs {
	s.Enabled = &v
	return s
}

func (s *AIAgentOutboundCallConfigBackChannelingConfigs) SetProbability(v float64) *AIAgentOutboundCallConfigBackChannelingConfigs {
	s.Probability = &v
	return s
}

func (s *AIAgentOutboundCallConfigBackChannelingConfigs) SetTriggerStage(v string) *AIAgentOutboundCallConfigBackChannelingConfigs {
	s.TriggerStage = &v
	return s
}

func (s *AIAgentOutboundCallConfigBackChannelingConfigs) SetWords(v []*AIAgentOutboundCallConfigBackChannelingConfigsWords) *AIAgentOutboundCallConfigBackChannelingConfigs {
	s.Words = v
	return s
}

func (s *AIAgentOutboundCallConfigBackChannelingConfigs) Validate() error {
	if s.Words != nil {
		for _, item := range s.Words {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AIAgentOutboundCallConfigBackChannelingConfigsWords struct {
	Probability *float64 `json:"Probability,omitempty" xml:"Probability,omitempty"`
	Text        *string  `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s AIAgentOutboundCallConfigBackChannelingConfigsWords) String() string {
	return dara.Prettify(s)
}

func (s AIAgentOutboundCallConfigBackChannelingConfigsWords) GoString() string {
	return s.String()
}

func (s *AIAgentOutboundCallConfigBackChannelingConfigsWords) GetProbability() *float64 {
	return s.Probability
}

func (s *AIAgentOutboundCallConfigBackChannelingConfigsWords) GetText() *string {
	return s.Text
}

func (s *AIAgentOutboundCallConfigBackChannelingConfigsWords) SetProbability(v float64) *AIAgentOutboundCallConfigBackChannelingConfigsWords {
	s.Probability = &v
	return s
}

func (s *AIAgentOutboundCallConfigBackChannelingConfigsWords) SetText(v string) *AIAgentOutboundCallConfigBackChannelingConfigsWords {
	s.Text = &v
	return s
}

func (s *AIAgentOutboundCallConfigBackChannelingConfigsWords) Validate() error {
	return dara.Validate(s)
}

type AIAgentOutboundCallConfigInterruptConfig struct {
	// Deprecated
	Eagerness *string `json:"Eagerness,omitempty" xml:"Eagerness,omitempty"`
	// Specifies whether to allow the user to interrupt the agent by speaking. Default value: true.
	//
	// example:
	//
	// true
	EnableVoiceInterrupt *bool `json:"EnableVoiceInterrupt,omitempty" xml:"EnableVoiceInterrupt,omitempty"`
	// Words or phrases that will trigger an interruption.
	InterruptWords           []*string `json:"InterruptWords,omitempty" xml:"InterruptWords,omitempty" type:"Repeated"`
	KeepInterruptWordsForLLM *bool     `json:"KeepInterruptWordsForLLM,omitempty" xml:"KeepInterruptWordsForLLM,omitempty"`
	NoInterruptMode          *string   `json:"NoInterruptMode,omitempty" xml:"NoInterruptMode,omitempty"`
}

func (s AIAgentOutboundCallConfigInterruptConfig) String() string {
	return dara.Prettify(s)
}

func (s AIAgentOutboundCallConfigInterruptConfig) GoString() string {
	return s.String()
}

func (s *AIAgentOutboundCallConfigInterruptConfig) GetEagerness() *string {
	return s.Eagerness
}

func (s *AIAgentOutboundCallConfigInterruptConfig) GetEnableVoiceInterrupt() *bool {
	return s.EnableVoiceInterrupt
}

func (s *AIAgentOutboundCallConfigInterruptConfig) GetInterruptWords() []*string {
	return s.InterruptWords
}

func (s *AIAgentOutboundCallConfigInterruptConfig) GetKeepInterruptWordsForLLM() *bool {
	return s.KeepInterruptWordsForLLM
}

func (s *AIAgentOutboundCallConfigInterruptConfig) GetNoInterruptMode() *string {
	return s.NoInterruptMode
}

func (s *AIAgentOutboundCallConfigInterruptConfig) SetEagerness(v string) *AIAgentOutboundCallConfigInterruptConfig {
	s.Eagerness = &v
	return s
}

func (s *AIAgentOutboundCallConfigInterruptConfig) SetEnableVoiceInterrupt(v bool) *AIAgentOutboundCallConfigInterruptConfig {
	s.EnableVoiceInterrupt = &v
	return s
}

func (s *AIAgentOutboundCallConfigInterruptConfig) SetInterruptWords(v []*string) *AIAgentOutboundCallConfigInterruptConfig {
	s.InterruptWords = v
	return s
}

func (s *AIAgentOutboundCallConfigInterruptConfig) SetKeepInterruptWordsForLLM(v bool) *AIAgentOutboundCallConfigInterruptConfig {
	s.KeepInterruptWordsForLLM = &v
	return s
}

func (s *AIAgentOutboundCallConfigInterruptConfig) SetNoInterruptMode(v string) *AIAgentOutboundCallConfigInterruptConfig {
	s.NoInterruptMode = &v
	return s
}

func (s *AIAgentOutboundCallConfigInterruptConfig) Validate() error {
	return dara.Validate(s)
}

type AIAgentOutboundCallConfigLlmConfig struct {
	// Alibaba Cloud Model Studio Application Center parameters. Reference: [Model Studio Application Center Parameter](https://help.aliyun.com/document_detail/2858132.html)
	BailianAppParams *string `json:"BailianAppParams,omitempty" xml:"BailianAppParams,omitempty"`
	// Maps agent capabilities to LLM functions. Only supports function calling with custom LLMs that adhere to the OpenAI protocol.
	FunctionMap        []*AIAgentOutboundCallConfigLlmConfigFunctionMap `json:"FunctionMap,omitempty" xml:"FunctionMap,omitempty" type:"Repeated"`
	HistorySyncWithTTS *bool                                            `json:"HistorySyncWithTTS,omitempty" xml:"HistorySyncWithTTS,omitempty"`
	// If true, the service sends the complete result from the LLM to the client in a single response after the generation process is finished.
	//
	// example:
	//
	// true
	LlmCompleteReply *bool `json:"LlmCompleteReply,omitempty" xml:"LlmCompleteReply,omitempty"`
	// The LLM/MLLM conversation history context.
	LlmHistory []*AIAgentOutboundCallConfigLlmConfigLlmHistory `json:"LlmHistory,omitempty" xml:"LlmHistory,omitempty" type:"Repeated"`
	// The maximum number of conversational turns to retain in the history. Default value: 10.
	//
	// example:
	//
	// 10
	LlmHistoryLimit *int32 `json:"LlmHistoryLimit,omitempty" xml:"LlmHistoryLimit,omitempty"`
	// The system prompt for the LLM.
	LlmSystemPrompt *string `json:"LlmSystemPrompt,omitempty" xml:"LlmSystemPrompt,omitempty"`
	// Additional query parameters to be sent to the OpenAI-protocol LLM, formatted as a URL query string (key=value pairs separated by &). All values must be strings.
	//
	// example:
	//
	// api-version=2024-02-01&api-key=sk-xxx
	OpenAIExtraQuery *string `json:"OpenAIExtraQuery,omitempty" xml:"OpenAIExtraQuery,omitempty"`
	// The maximum time (in milliseconds) to buffer text before it is forcibly sent to the client. Valid values: [1000,10000]. A value of 0 or an empty string (default) disables this limit.
	//
	// example:
	//
	// 2000
	OutputMaxDelay *string `json:"OutputMaxDelay,omitempty" xml:"OutputMaxDelay,omitempty"`
	// The minimum number of characters that must be buffered before a text chunk is sent. Valid values: [0, 100]. A value of 0 or an empty string (default) disables this limit.
	//
	// example:
	//
	// 5
	OutputMinLength *int32 `json:"OutputMinLength,omitempty" xml:"OutputMinLength,omitempty"`
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

func (s *AIAgentOutboundCallConfigLlmConfig) GetHistorySyncWithTTS() *bool {
	return s.HistorySyncWithTTS
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

func (s *AIAgentOutboundCallConfigLlmConfig) SetHistorySyncWithTTS(v bool) *AIAgentOutboundCallConfigLlmConfig {
	s.HistorySyncWithTTS = &v
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
	// The name of the built-in agent capability. Only hangup is supported.
	//
	// example:
	//
	// hangup
	Function *string `json:"Function,omitempty" xml:"Function,omitempty"`
	// The corresponding user-defined function name in your LLM. When the LLM calls this function, it will trigger the mapped agent capability.[](~~2839094~~)
	//
	// example:
	//
	// hangup
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
	// The actual text content of the message for that role.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The role of the participant in the conversation. Valid values:
	//
	// 	- user
	//
	// 	- assistant
	//
	// 	- system
	//
	// 	- function
	//
	// 	- plugin
	//
	// 	- tool
	//
	// example:
	//
	// user
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
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
	// Applies only to MiniMax models. Seven types of emotions are supported:
	//
	// 	- happy
	//
	// 	- sad
	//
	// 	- angry
	//
	// 	- fearful
	//
	// 	- disgusted
	//
	// 	- surprised
	//
	// 	- calm
	//
	// example:
	//
	// happy
	Emotion *string `json:"Emotion,omitempty" xml:"Emotion,omitempty"`
	// Applies only to MiniMax models. By default, this parameter is left empty. This enhances speech recognition accuracy for specific languages and dialects. If the language type is unknown, set it to auto to have the model automatically detect it. Valid values:
	//
	// **Supported languages**
	//
	// 	- Chinese
	//
	// 	- Chinese,Yue
	//
	// 	- English
	//
	// 	- Arabic
	//
	// 	- Russian
	//
	// 	- Spanish
	//
	// 	- French
	//
	// 	- Portuguese
	//
	// 	- German
	//
	// 	- Turkish
	//
	// 	- Dutch
	//
	// 	- Ukrainian
	//
	// 	- Vietnamese
	//
	// 	- Indonesian
	//
	// 	- Japanese
	//
	// 	- Italian
	//
	// 	- Korean
	//
	// 	- Thai
	//
	// 	- Polish
	//
	// 	- Romanian
	//
	// 	- Greek
	//
	// 	- Czech
	//
	// 	- Finnish
	//
	// 	- Hindi
	//
	// 	- auto
	//
	// example:
	//
	// Chinese
	LanguageId *string `json:"LanguageId,omitempty" xml:"LanguageId,omitempty"`
	// Applies only to MiniMax models. Valid values: speech-01-turbo and speech-02-turbo.
	//
	// example:
	//
	// speech-01-turbo
	ModelId *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	// The pronunciation rules, executed in order. Maximum of 20 rules.
	PronunciationRules []*AIAgentOutboundCallConfigTtsConfigPronunciationRules `json:"PronunciationRules,omitempty" xml:"PronunciationRules,omitempty" type:"Repeated"`
	// Supports all models. For CosyVoice, the default value is 1.0. Valid values: 0.5 to 2.0. For MiniMax, the default value is 1.0. Valid values: 0.5 to 2.0.
	//
	// example:
	//
	// 1.0
	SpeechRate *float64 `json:"SpeechRate,omitempty" xml:"SpeechRate,omitempty"`
	// The voice ID. Changes take effect on the next sentence. If not set, the system uses the default voice ID specified in the agent template. This parameter takes effect only for the preset TTS model. Max length: 64 characters. Refer to [Intelligent voice samples](https://help.aliyun.com/document_detail/449563.html) for options.
	//
	// example:
	//
	// longcheng_v2
	VoiceId *string `json:"VoiceId,omitempty" xml:"VoiceId,omitempty"`
	// Available voices.
	VoiceIdList []*string `json:"VoiceIdList,omitempty" xml:"VoiceIdList,omitempty" type:"Repeated"`
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
	// The target pronunciation. The value supports up to 10 Chinese characters. Other characters, including spaces, are not supported.
	Pronunciation *string `json:"Pronunciation,omitempty" xml:"Pronunciation,omitempty"`
	// The type of rule. Valid value:
	//
	// 	- replacement: replaces every occurrence of Word value with Pronunciation value.
	//
	// example:
	//
	// replacement
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The word to be replaced. The value supports up to 10 Chinese characters. Other characters, including spaces, are not supported.
	Word *string `json:"Word,omitempty" xml:"Word,omitempty"`
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
	Eagerness *string `json:"Eagerness,omitempty" xml:"Eagerness,omitempty"`
	// The mode of turn detection.
	//
	// 	- Normal: uses simple pause detection.
	//
	// 	- Semantic: uses AI to analyze context.
	//
	// Default value: Normal.
	//
	// example:
	//
	// Semantic
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// Specifies how long to wait after a user stops speaking for the agent to decide if the turn is over. Unit: milliseconds. Default value: -1.
	//
	// 	- \\-1: AI decides an appropriate wait time automatically.
	//
	// 	- 0 to 10000: A custom wait time. Recommended: 0 to 1500 ms.
	//
	// Note: In Normal mode, this field is ignored.
	//
	// example:
	//
	// -1
	SemanticWaitDuration *int32 `json:"SemanticWaitDuration,omitempty" xml:"SemanticWaitDuration,omitempty"`
	// Keywords that signify the end of the user\\"s turn.
	TurnEndWords []*string `json:"TurnEndWords,omitempty" xml:"TurnEndWords,omitempty" type:"Repeated"`
}

func (s AIAgentOutboundCallConfigTurnDetectionConfig) String() string {
	return dara.Prettify(s)
}

func (s AIAgentOutboundCallConfigTurnDetectionConfig) GoString() string {
	return s.String()
}

func (s *AIAgentOutboundCallConfigTurnDetectionConfig) GetEagerness() *string {
	return s.Eagerness
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

func (s *AIAgentOutboundCallConfigTurnDetectionConfig) SetEagerness(v string) *AIAgentOutboundCallConfigTurnDetectionConfig {
	s.Eagerness = &v
	return s
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
