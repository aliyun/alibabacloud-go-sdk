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
	// The configurations for ambient sound.
	AmbientSoundConfig *AIAgentOutboundCallConfigAmbientSoundConfig `json:"AmbientSoundConfig,omitempty" xml:"AmbientSoundConfig,omitempty" type:"Struct"`
	// The automatic speech recognition (ASR) configurations.
	AsrConfig *AIAgentOutboundCallConfigAsrConfig `json:"AsrConfig,omitempty" xml:"AsrConfig,omitempty" type:"Struct"`
	// The configurations for the automatic speech module of the AI agent, which includes prompts during LLM delays and inquiries during prolonged user silence.
	AutoSpeechConfig *AIAgentOutboundCallConfigAutoSpeechConfig `json:"AutoSpeechConfig,omitempty" xml:"AutoSpeechConfig,omitempty" type:"Struct"`
	// Deprecated
	//
	// 	Notice:
	//
	// This parameter is deprecated. Use `BackChannelingConfigs` instead.
	BackChannelingConfig *AIAgentOutboundCallConfigBackChannelingConfig `json:"BackChannelingConfig,omitempty" xml:"BackChannelingConfig,omitempty" type:"Struct"`
	// The configurations for the back-channeling feature module. If you enable this feature, the system randomly plays short and affirmative phrases at specific trigger points.
	BackChannelingConfigs []*AIAgentOutboundCallConfigBackChannelingConfigs `json:"BackChannelingConfigs,omitempty" xml:"BackChannelingConfigs,omitempty" type:"Repeated"`
	// Specifies whether to enable intelligent segmentation. If you enable this feature, short and consecutive speech segments from the user are merged into a complete sentence. Default value: `true`.
	//
	// example:
	//
	// true
	EnableIntelligentSegment *bool `json:"EnableIntelligentSegment,omitempty" xml:"EnableIntelligentSegment,omitempty"`
	// The parameters for experimental features. If you have any requirements, contact technical support.
	//
	// example:
	//
	// ""
	ExperimentalConfig *string `json:"ExperimentalConfig,omitempty" xml:"ExperimentalConfig,omitempty"`
	// The welcome message. This change takes effect in the next call session. If this parameter is not set, no welcome message is played.
	//
	// example:
	//
	// 你好
	Greeting *string `json:"Greeting,omitempty" xml:"Greeting,omitempty"`
	// The delay before the welcome message is played. Unit: ms. Default value: 0. Valid values: 0 to 5000.
	//
	// example:
	//
	// 0
	GreetingDelay *int32 `json:"GreetingDelay,omitempty" xml:"GreetingDelay,omitempty"`
	// The speech interruption policy configurations.
	InterruptConfig *AIAgentOutboundCallConfigInterruptConfig `json:"InterruptConfig,omitempty" xml:"InterruptConfig,omitempty" type:"Struct"`
	// The configurations of the large language model (LLM).
	LlmConfig *AIAgentOutboundCallConfigLlmConfig `json:"LlmConfig,omitempty" xml:"LlmConfig,omitempty" type:"Struct"`
	// The maximum wait time for interaction with the AI agent. If the wait time is exceeded, the AI agent goes offline. Unit: seconds. Default value: 600.
	//
	// example:
	//
	// 600
	MaxIdleTime *int32 `json:"MaxIdleTime,omitempty" xml:"MaxIdleTime,omitempty"`
	// The text-to-speech (TTS) configurations.
	TtsConfig *AIAgentOutboundCallConfigTtsConfig `json:"TtsConfig,omitempty" xml:"TtsConfig,omitempty" type:"Struct"`
	// The configurations for conversational turn detection.
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
	// The ID of the ambient sound. You can obtain the ID from the advanced configurations of the AI agent on the console.
	//
	// example:
	//
	// f67901c595834************
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The volume of the ambient sound. Valid values: 0 to 100. A value of 0 disables the sound.
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
	// The list of hotwords for ASR. You can specify a maximum of 128 hotwords in the list.
	AsrHotWords []*string `json:"AsrHotWords,omitempty" xml:"AsrHotWords,omitempty" type:"Repeated"`
	// The language ID for ASR. Valid values:
	//
	// - `zh_mandarin`: Chinese
	//
	// - `en`: English
	//
	// - `zh_en`: Chinese and English
	//
	// - `es`: Spanish
	//
	// - `jp`: Japanese
	//
	// example:
	//
	// zh_mandarin
	AsrLanguageId *string `json:"AsrLanguageId,omitempty" xml:"AsrLanguageId,omitempty"`
	// The sentence segmentation threshold. If the duration of a silence exceeds this threshold, the system determines that the sentence is complete. Valid values: 200 to 1200. Unit: ms. Default value: 400.
	//
	// example:
	//
	// 400
	AsrMaxSilence *int32 `json:"AsrMaxSilence,omitempty" xml:"AsrMaxSilence,omitempty"`
	// The passthrough parameters for proprietary ASR.
	//
	// example:
	//
	// mode=fast&sample=16000&format=wav
	CustomParams *string `json:"CustomParams,omitempty" xml:"CustomParams,omitempty"`
	// The minimum duration threshold for VAD. This parameter controls the interruption sensitivity. A value of 0 indicates that this feature is disabled. Valid values: 200 to 2000. Unit: ms. A value from 200 to 500 corresponds to 1 to 4 words. The default value is empty, which indicates that this parameter is not in effect.
	//
	// example:
	//
	// 300
	VadDuration *int32 `json:"VadDuration,omitempty" xml:"VadDuration,omitempty"`
	// The interruption threshold for voice activity detection (VAD). Valid values: 0 to 11. Default value: 11.
	//
	// - A value of 0 disables the VAD feature.
	//
	// - A value from 1 to 10 indicates that the higher the value, the less sensitive the interruption.
	//
	// - A value of 11 provides a significantly different experience from the previous values. It lowers audio distortion during conversations and enhances resistance to interference.
	//
	// example:
	//
	// 11
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
	// The configurations for broadcasts during LLM response delays.
	LlmPending *AIAgentOutboundCallConfigAutoSpeechConfigLlmPending `json:"LlmPending,omitempty" xml:"LlmPending,omitempty" type:"Struct"`
	// The configurations for inquiry broadcasts during prolonged user silence.
	UserIdle *AIAgentOutboundCallConfigAutoSpeechConfigUserIdle `json:"UserIdle,omitempty" xml:"UserIdle,omitempty" type:"Struct"`
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
	// The collection of inquiry prompts. You can specify a maximum of 10 prompts. Each prompt can be up to 100 characters in length. The sum of the probabilities of all prompts must be 100%.
	Messages []*AIAgentOutboundCallConfigAutoSpeechConfigLlmPendingMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Repeated"`
	Mode     *string                                                        `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The wait time threshold for LLM responses. This parameter is required. A broadcast prompt is triggered if this threshold is exceeded. Unit: ms. Valid values: 500 to 10000. You need to configure this parameter based on the actual usage of the LLM.
	//
	// example:
	//
	// 3000
	WaitTime *int32 `json:"WaitTime,omitempty" xml:"WaitTime,omitempty"`
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

func (s *AIAgentOutboundCallConfigAutoSpeechConfigLlmPending) GetMode() *string {
	return s.Mode
}

func (s *AIAgentOutboundCallConfigAutoSpeechConfigLlmPending) GetWaitTime() *int32 {
	return s.WaitTime
}

func (s *AIAgentOutboundCallConfigAutoSpeechConfigLlmPending) SetMessages(v []*AIAgentOutboundCallConfigAutoSpeechConfigLlmPendingMessages) *AIAgentOutboundCallConfigAutoSpeechConfigLlmPending {
	s.Messages = v
	return s
}

func (s *AIAgentOutboundCallConfigAutoSpeechConfigLlmPending) SetMode(v string) *AIAgentOutboundCallConfigAutoSpeechConfigLlmPending {
	s.Mode = &v
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
	// The selection probability of the prompt. Valid values: 0 to 1, which corresponds to 0% to 100%.
	//
	// example:
	//
	// 0.5
	Probability *float64 `json:"Probability,omitempty" xml:"Probability,omitempty"`
	// The text of the inquiry prompt. The text can be up to 100 characters in length.
	//
	// example:
	//
	// 稍等一下
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
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
	// example:
	//
	// 我挂了，再见
	HangupEndWord *string `json:"HangupEndWord,omitempty" xml:"HangupEndWord,omitempty"`
	// The maximum number of inquiries. This parameter is required. Valid values: 0 to 10. After the maximum number of inquiries is reached, no more inquiries are triggered, and the call is disconnected.
	//
	// example:
	//
	// 5
	MaxRepeats *int32 `json:"MaxRepeats,omitempty" xml:"MaxRepeats,omitempty"`
	// The collection of inquiry prompts. You can specify a maximum of 10 prompts. Each prompt can be up to 100 characters in length. The sum of the probabilities of all prompts must be 100%.
	Messages []*AIAgentOutboundCallConfigAutoSpeechConfigUserIdleMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Repeated"`
	// The silence duration threshold. This parameter is required. An inquiry is triggered if this threshold is exceeded. Unit: ms. Valid values: 5000 to 600000.
	//
	// example:
	//
	// 5000
	WaitTime *int32 `json:"WaitTime,omitempty" xml:"WaitTime,omitempty"`
}

func (s AIAgentOutboundCallConfigAutoSpeechConfigUserIdle) String() string {
	return dara.Prettify(s)
}

func (s AIAgentOutboundCallConfigAutoSpeechConfigUserIdle) GoString() string {
	return s.String()
}

func (s *AIAgentOutboundCallConfigAutoSpeechConfigUserIdle) GetHangupEndWord() *string {
	return s.HangupEndWord
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

func (s *AIAgentOutboundCallConfigAutoSpeechConfigUserIdle) SetHangupEndWord(v string) *AIAgentOutboundCallConfigAutoSpeechConfigUserIdle {
	s.HangupEndWord = &v
	return s
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
	// The selection probability of the prompt. Valid values: 0 to 1, which corresponds to 0% to 100%.
	//
	// example:
	//
	// 0.5
	Probability *float64 `json:"Probability,omitempty" xml:"Probability,omitempty"`
	// The text of the inquiry prompt. The text can be up to 100 characters in length.
	//
	// example:
	//
	// 您还在吗？
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
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
	// Specifies whether to enable the back-channeling feature. This parameter is required. Valid values: true and false.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The trigger probability. This parameter is required. Valid values: 0.0 to 1.0.
	//
	// example:
	//
	// 0.5
	Probability *float64 `json:"Probability,omitempty" xml:"Probability,omitempty"`
	// The trigger point for back-channeling. Valid value:
	//
	// - `pause_detected`: triggered when a short pause in speech is detected
	//
	// example:
	//
	// pause_detected
	TriggerStage *string `json:"TriggerStage,omitempty" xml:"TriggerStage,omitempty"`
	// The collection of back-channeling phrases. You can specify a maximum of 10 phrases. Each phrase can be up to 20 characters in length. The sum of the probabilities of all phrases must be 1.0.
	Words []*AIAgentOutboundCallConfigBackChannelingConfigsWords `json:"Words,omitempty" xml:"Words,omitempty" type:"Repeated"`
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
	// The selection probability of this phrase. This parameter is required. Valid values: 0.0 to 1.0.
	//
	// example:
	//
	// 0.3
	Probability *float64 `json:"Probability,omitempty" xml:"Probability,omitempty"`
	// The text of the phrase. This parameter is required. The text can be up to 20 characters in length and supports multiple languages.
	//
	// example:
	//
	// 嗯嗯
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
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
	// Specifies whether to support speech interruption. Default value: true.
	//
	// example:
	//
	// true
	EnableVoiceInterrupt *bool `json:"EnableVoiceInterrupt,omitempty" xml:"EnableVoiceInterrupt,omitempty"`
	// The specific words or phrases that trigger a conversation interruption.
	InterruptWords []*string `json:"InterruptWords,omitempty" xml:"InterruptWords,omitempty" type:"Repeated"`
	// example:
	//
	// true
	KeepInterruptWordsForLLM *bool `json:"KeepInterruptWordsForLLM,omitempty" xml:"KeepInterruptWordsForLLM,omitempty"`
	// The ASR processing policy in `NoInterruptMode`.
	//
	// - `cache`: caches the ASR text. The cached ASR text is processed in the next conversational turn.
	//
	// - `discard`: discards the ASR text.
	//
	// Default value: cache.
	//
	// example:
	//
	// cache
	NoInterruptMode *string `json:"NoInterruptMode,omitempty" xml:"NoInterruptMode,omitempty"`
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
	// The parameters for Alibaba Cloud Model Studio. For more information about the parameter format, see [Alibaba Cloud Model Studio parameters](https://help.aliyun.com/document_detail/2858132.html).
	//
	// example:
	//
	// "{\\"biz_params\\":{\\"user_defined_params\\":{\\"your_plugin_id\\":{\\"article_index\\":2}}},\\"memory_id\\":\\"your_memory_id\\",\\"image_list\\":[\\"https://your_image_url\\"],\\"rag_options\\":{\\"pipeline_ids\\":[\\"your_id\\"],\\"file_ids\\":[\\"文档ID1\\",\\"文档ID2\\"],\\"metadata_filter\\":{\\"name\\":\\"张三\\"},\\"structured_filter\\":{\\"key1\\":\\"value1\\",\\"key2\\":\\"value2\\"},\\"tags\\":[\\"标签1\\",\\"标签2\\"]}}"
	BailianAppParams *string `json:"BailianAppParams,omitempty" xml:"BailianAppParams,omitempty"`
	// The list of function mappings, which is used to map AI agent capabilities to LLM functions. This feature is supported only when function calls are used in custom LLMs that are compatible with the OpenAI protocol.
	FunctionMap []*AIAgentOutboundCallConfigLlmConfigFunctionMap `json:"FunctionMap,omitempty" xml:"FunctionMap,omitempty" type:"Repeated"`
	// Specifies whether to keep the LLM message history consistent with the TTS playback content. Default value: false. If you enable this feature, the saved LLM messages are consistent with the TTS playback content.
	//
	// example:
	//
	// false
	HistorySyncWithTTS *bool `json:"HistorySyncWithTTS,omitempty" xml:"HistorySyncWithTTS,omitempty"`
	// If you enable this feature, the system sends the complete LLM-generated result to the client after the generation is complete.
	//
	// example:
	//
	// true
	LlmCompleteReply *bool `json:"LlmCompleteReply,omitempty" xml:"LlmCompleteReply,omitempty"`
	// The conversation history of the LLM or MLLM.
	LlmHistory []*AIAgentOutboundCallConfigLlmConfigLlmHistory `json:"LlmHistory,omitempty" xml:"LlmHistory,omitempty" type:"Repeated"`
	// The maximum number of conversational turns to retain in the history of the LLM or multimodal large language model (MLLM). Default value: 10.
	//
	// example:
	//
	// 10
	LlmHistoryLimit *int32 `json:"LlmHistoryLimit,omitempty" xml:"LlmHistoryLimit,omitempty"`
	// The system prompt for the LLM after the call is initiated.
	//
	// example:
	//
	// 你是一位友好且乐于助人的助手，专注于为用户提供准确的信息和建议。
	LlmSystemPrompt *string `json:"LlmSystemPrompt,omitempty" xml:"LlmSystemPrompt,omitempty"`
	// The additional query parameters for an LLM that is compatible with the OpenAI protocol. The parameters must be in the key=value format. If you specify multiple parameters, separate them with ampersands (`&`). All values must be of the string type.
	//
	// example:
	//
	// api-version=2024-02-01&api-key=sk-xxx
	OpenAIExtraQuery *string `json:"OpenAIExtraQuery,omitempty" xml:"OpenAIExtraQuery,omitempty"`
	// The maximum delay for text output. If this threshold is exceeded, the cached text is forcibly output. Valid values: 1000 to 10000. Unit: ms. A value of 0 or empty indicates that this parameter is not in effect. Default value: empty.
	//
	// example:
	//
	// 2000
	OutputMaxDelay *string `json:"OutputMaxDelay,omitempty" xml:"OutputMaxDelay,omitempty"`
	// The minimum length of text output. The unit is characters. Text shorter than this length is cached and waits for concatenation. Valid values: 0 to 100. A value of 0 or empty indicates that this parameter is not in effect. Default value: empty.
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
	// The name of the built-in function provided by the AI agent in Alibaba Cloud. The value hangup is supported.
	//
	// example:
	//
	// hangup
	Function *string `json:"Function,omitempty" xml:"Function,omitempty"`
	// The name of the LLM function that corresponds to this function. This parameter is customized and used to call the corresponding function in the LLM. For more information about the protocol for custom LLMs, see [Standard LLM API](https://help.aliyun.com/document_detail/2839094.html).
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
	// The text of the conversation content that records the specific expressions or responses of the role in the conversation.
	//
	// example:
	//
	// 你好
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The role of the participant in the conversation. Valid values:
	//
	// - `user`: the user
	//
	// - `assistant`: the AI assistant
	//
	// - `system`: the system
	//
	// - `function`: a function
	//
	// - `plugin`: a plug-in
	//
	// - `tool`: a tool
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
	// Only MiniMax is supported. The following seven emotions are supported:
	//
	// - `happy`
	//
	// - `sad`
	//
	// - `angry`
	//
	// - `fearful`
	//
	// - `disgusted`
	//
	// - `surprised`
	//
	// - `calm`
	//
	// example:
	//
	// happy
	Emotion *string `json:"Emotion,omitempty" xml:"Emotion,omitempty"`
	// Only MiniMax is supported. The default value is empty. This parameter enhances the recognition of specific minority languages and dialects. After you set this parameter, the speech performance in the specified minority language or dialect scenarios is improved. If the minority language type is unknown, you can set this parameter to `auto` to enable the model to automatically determine the minority language type. Valid values:
	//
	// <details>
	//
	// <summary>
	//
	// Supported languages
	//
	// </summary>
	//
	// - `Chinese`: Chinese
	//
	// - `Chinese,Yue`: Cantonese
	//
	// - `English`: English
	//
	// - `Arabic`: Arabic
	//
	// - `Russian`: Russian
	//
	// - `Spanish`: Spanish
	//
	// - `French`: French
	//
	// - `Portuguese`: Portuguese
	//
	// - `German`: German
	//
	// - `Turkish`: Turkish
	//
	// - `Dutch`: Dutch
	//
	// - `Ukrainian`: Ukrainian
	//
	// - `Vietnamese`: Vietnamese
	//
	// - `Indonesian`: Indonesian
	//
	// - `Japanese`: Japanese
	//
	// - `Italian`: Italian
	//
	// - `Korean`: Korean
	//
	// - `Thai`: Thai
	//
	// - `Polish`: Polish
	//
	// - `Romanian`: Romanian
	//
	// - `Greek`: Greek
	//
	// - `Czech`: Czech
	//
	// - `Finnish`: Finnish
	//
	// - `Hindi`: Hindi
	//
	// - `auto`: Automatic detection
	//
	// </details>
	//
	// example:
	//
	// Chinese
	LanguageId *string `json:"LanguageId,omitempty" xml:"LanguageId,omitempty"`
	// Only MiniMax is supported. Valid values: `speech-01-turbo` and `speech-02-turbo`.
	//
	// example:
	//
	// speech-01-turbo
	ModelId *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	// The TTS pronunciation rules. You can specify a maximum of 20 rules in the array. The rules are executed in sequence.
	PronunciationRules []*AIAgentOutboundCallConfigTtsConfigPronunciationRules `json:"PronunciationRules,omitempty" xml:"PronunciationRules,omitempty" type:"Repeated"`
	// This parameter is supported on all platforms. For CosyVoice, the default value is 1.0 and the valid values are 0.5 to 2.0. For MiniMax, the default value is 1.0 and the valid values are 0.5 to 2.0.
	//
	// example:
	//
	// 1.0
	SpeechRate *float64 `json:"SpeechRate,omitempty" xml:"SpeechRate,omitempty"`
	// The voice ID. The change takes effect on the next sentence. If you do not specify this parameter, the voice ID configured in the AI agent template is used. This parameter is valid only for preset TTS voices. The value can be up to 64 characters in length. For more information about the valid values, see [Intelligent speech effect samples](https://help.aliyun.com/document_detail/449563.html).
	//
	// example:
	//
	// longcheng_v2
	VoiceId *string `json:"VoiceId,omitempty" xml:"VoiceId,omitempty"`
	// The list of available voices.
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
	// The target pronunciation. The pronunciation must be a Chinese character string of up to 10 characters in length and cannot contain spaces.
	//
	// example:
	//
	// 大石烂儿
	Pronunciation *string `json:"Pronunciation,omitempty" xml:"Pronunciation,omitempty"`
	// The type of the pronunciation rule. Valid value:
	//
	// - `replacement`: replaces the word with the specified pronunciation.
	//
	// example:
	//
	// replacement
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The word to be replaced. The word must be a Chinese character string of up to 10 characters in length and cannot contain spaces.
	//
	// example:
	//
	// 大栅栏
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
	// example:
	//
	// Low
	Eagerness *string `json:"Eagerness,omitempty" xml:"Eagerness,omitempty"`
	// The mode for conversational turn detection. Valid values:
	//
	// - `Normal`: a basic mode that does not use AI for semantic analysis.
	//
	// - `Semantic`: an AI-powered mode that determines whether the user has finished speaking based on the conversational context.
	//
	// Default value: `Normal`.
	//
	// example:
	//
	// Semantic
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The pause duration in AI mode that is used to determine whether a conversational turn has ended. Unit: ms. Default value: -1.
	//
	// - `-1`: The AI automatically determines an appropriate wait time.
	//
	// - `0-10000`: A custom wait time. We recommend that you set this parameter to a value from 0 to 1500.
	//
	// Note: This parameter is invalid in Normal mode.
	//
	// example:
	//
	// -1
	SemanticWaitDuration *int32 `json:"SemanticWaitDuration,omitempty" xml:"SemanticWaitDuration,omitempty"`
	// The list of keywords that are used to determine the end of a user\\"s conversational turn.
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
