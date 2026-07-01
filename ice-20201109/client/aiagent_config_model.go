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
	SetAutoSpeechConfig(v *AIAgentConfigAutoSpeechConfig) *AIAgentConfig
	GetAutoSpeechConfig() *AIAgentConfigAutoSpeechConfig
	SetAvatarConfig(v *AIAgentConfigAvatarConfig) *AIAgentConfig
	GetAvatarConfig() *AIAgentConfigAvatarConfig
	SetAvatarUrl(v string) *AIAgentConfig
	GetAvatarUrl() *string
	SetAvatarUrlType(v string) *AIAgentConfig
	GetAvatarUrlType() *string
	SetBackChannelingConfig(v []*AIAgentConfigBackChannelingConfig) *AIAgentConfig
	GetBackChannelingConfig() []*AIAgentConfigBackChannelingConfig
	SetBackChannelingConfigs(v []*AIAgentConfigBackChannelingConfigs) *AIAgentConfig
	GetBackChannelingConfigs() []*AIAgentConfigBackChannelingConfigs
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
	// Configuration for ambient sound during the call.
	AmbientSoundConfig *AIAgentConfigAmbientSoundConfig `json:"AmbientSoundConfig,omitempty" xml:"AmbientSoundConfig,omitempty" type:"Struct"`
	// Configuration for automatic speech recognition (ASR).
	AsrConfig *AIAgentConfigAsrConfig `json:"AsrConfig,omitempty" xml:"AsrConfig,omitempty" type:"Struct"`
	// Configuration for the agent\\"s automatic speech, including prompts for LLM latency and long periods of user silence.
	AutoSpeechConfig *AIAgentConfigAutoSpeechConfig `json:"AutoSpeechConfig,omitempty" xml:"AutoSpeechConfig,omitempty" type:"Struct"`
	// Configuration for the avatar. This takes effect only if the workflow includes an avatar node.
	AvatarConfig *AIAgentConfigAvatarConfig `json:"AvatarConfig,omitempty" xml:"AvatarConfig,omitempty" type:"Struct"`
	// The URL of the avatar to display during voice calls. If omitted, no avatar is displayed.
	//
	// example:
	//
	// http://example.com/a.jpg
	AvatarUrl *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	// The type of the avatar URL. By default, this parameter is not set.
	//
	// example:
	//
	// USER
	AvatarUrlType *string `json:"AvatarUrlType,omitempty" xml:"AvatarUrlType,omitempty"`
	// Deprecated
	//
	// 	Notice:
	//
	// 已废弃，请使用 BackChannelingConfigs
	BackChannelingConfig []*AIAgentConfigBackChannelingConfig `json:"BackChannelingConfig,omitempty" xml:"BackChannelingConfig,omitempty" type:"Repeated"`
	// Configuration for back-channeling. When enabled, the system plays short, responsive phrases at specific trigger points.
	BackChannelingConfigs []*AIAgentConfigBackChannelingConfigs `json:"BackChannelingConfigs,omitempty" xml:"BackChannelingConfigs,omitempty" type:"Repeated"`
	// Specifies whether to enable intelligent segmentation. When enabled, short user utterances are merged into a single sentence. Default: `true`.
	//
	// example:
	//
	// true
	EnableIntelligentSegment *bool `json:"EnableIntelligentSegment,omitempty" xml:"EnableIntelligentSegment,omitempty"`
	// Specifies whether to enable push-to-talk mode. Default: `false`.
	//
	// example:
	//
	// false
	EnablePushToTalk *bool `json:"EnablePushToTalk,omitempty" xml:"EnablePushToTalk,omitempty"`
	// Parameters for experimental features. Contact support for assistance.
	//
	// example:
	//
	// ""
	ExperimentalConfig *string `json:"ExperimentalConfig,omitempty" xml:"ExperimentalConfig,omitempty"`
	// Specifies whether to enable graceful shutdown. Default: `false`.
	//
	// If enabled, the AI agent completes its current utterance before disconnecting when the task is stopped. The agent will not speak for more than 10 seconds.
	//
	// example:
	//
	// false
	GracefulShutdown *bool `json:"GracefulShutdown,omitempty" xml:"GracefulShutdown,omitempty"`
	// The welcome message the AI agent plays when joining the session. Changes apply to subsequent sessions. If omitted, no welcome message is played.
	//
	// example:
	//
	// 你好
	Greeting *string `json:"Greeting,omitempty" xml:"Greeting,omitempty"`
	// Configuration for the speech interruption policy.
	InterruptConfig *AIAgentConfigInterruptConfig `json:"InterruptConfig,omitempty" xml:"InterruptConfig,omitempty" type:"Struct"`
	// Configuration for the large language model (LLM).
	LlmConfig *AIAgentConfigLlmConfig `json:"LlmConfig,omitempty" xml:"LlmConfig,omitempty" type:"Struct"`
	// The maximum idle duration in seconds before the AI agent disconnects. If the agent receives no user interaction within this period, it ends the task. Default: 600.
	//
	// example:
	//
	// 600
	MaxIdleTime *int32 `json:"MaxIdleTime,omitempty" xml:"MaxIdleTime,omitempty"`
	// Configuration for text-to-speech (TTS).
	TtsConfig *AIAgentConfigTtsConfig `json:"TtsConfig,omitempty" xml:"TtsConfig,omitempty" type:"Struct"`
	// Configuration for conversational turn detection.
	TurnDetectionConfig *AIAgentConfigTurnDetectionConfig `json:"TurnDetectionConfig,omitempty" xml:"TurnDetectionConfig,omitempty" type:"Struct"`
	// The duration in seconds the AI agent waits before terminating the task after a user leaves the session. Default: 5.
	//
	// example:
	//
	// 5
	UserOfflineTimeout *int32 `json:"UserOfflineTimeout,omitempty" xml:"UserOfflineTimeout,omitempty"`
	// The duration in seconds the AI agent waits for a user to join. If the user does not join within this time, the agent terminates the task. Default: 60.
	//
	// example:
	//
	// 60
	UserOnlineTimeout *int32 `json:"UserOnlineTimeout,omitempty" xml:"UserOnlineTimeout,omitempty"`
	// Configuration for video content recognition. This enables the system to send callbacks to the client about events detected in the video stream.
	VcrConfig *AIAgentConfigVcrConfig `json:"VcrConfig,omitempty" xml:"VcrConfig,omitempty" type:"Struct"`
	// Configuration for voiceprint recognition.
	VoiceprintConfig *AIAgentConfigVoiceprintConfig `json:"VoiceprintConfig,omitempty" xml:"VoiceprintConfig,omitempty" type:"Struct"`
	// The speaking volume of the AI agent.
	//
	// - If not set, the adaptive volume mode recommended by Alibaba Cloud is used by default.
	//
	// - If set, the value must be in the range of 0 to 400. The final output volume is calculated as: `(Workflow volume) 	- (volume / 100)`. For example:
	//
	// 1. If `volume` is 0, the output volume is 0.
	//
	// 2. If `volume` is 100, the output volume is the same as the original volume.
	//
	// 3. If `volume` is 200, the output volume is twice the original volume.
	//
	// example:
	//
	// 100
	Volume *int64 `json:"Volume,omitempty" xml:"Volume,omitempty"`
	// A user-provided command that the AI agent responds to immediately after the call starts.
	//
	// example:
	//
	// 今天天气怎么样？
	WakeUpQuery *string `json:"WakeUpQuery,omitempty" xml:"WakeUpQuery,omitempty"`
	// A JSON string containing parameters to override the default workflow configuration.
	//
	// example:
	//
	// {}
	WorkflowOverrideParams *string `json:"WorkflowOverrideParams,omitempty" xml:"WorkflowOverrideParams,omitempty"`
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

func (s *AIAgentConfig) GetAutoSpeechConfig() *AIAgentConfigAutoSpeechConfig {
	return s.AutoSpeechConfig
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

func (s *AIAgentConfig) GetBackChannelingConfig() []*AIAgentConfigBackChannelingConfig {
	return s.BackChannelingConfig
}

func (s *AIAgentConfig) GetBackChannelingConfigs() []*AIAgentConfigBackChannelingConfigs {
	return s.BackChannelingConfigs
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

func (s *AIAgentConfig) SetAutoSpeechConfig(v *AIAgentConfigAutoSpeechConfig) *AIAgentConfig {
	s.AutoSpeechConfig = v
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

func (s *AIAgentConfig) SetBackChannelingConfig(v []*AIAgentConfigBackChannelingConfig) *AIAgentConfig {
	s.BackChannelingConfig = v
	return s
}

func (s *AIAgentConfig) SetBackChannelingConfigs(v []*AIAgentConfigBackChannelingConfigs) *AIAgentConfig {
	s.BackChannelingConfigs = v
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
	if s.AvatarConfig != nil {
		if err := s.AvatarConfig.Validate(); err != nil {
			return err
		}
	}
	if s.BackChannelingConfig != nil {
		for _, item := range s.BackChannelingConfig {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
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
	if s.VcrConfig != nil {
		if err := s.VcrConfig.Validate(); err != nil {
			return err
		}
	}
	if s.VoiceprintConfig != nil {
		if err := s.VoiceprintConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AIAgentConfigAmbientSoundConfig struct {
	// The ID of the ambient sound resource. You can obtain this ID from the advanced settings of the agent configuration in the console.
	//
	// example:
	//
	// f67901c595834************
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The volume of the ambient sound. Range: 0–100. A value of 0 disables the sound.
	//
	// example:
	//
	// 50
	Volume *int32 `json:"Volume,omitempty" xml:"Volume,omitempty"`
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
	// A list of hotwords to improve ASR accuracy. You can specify a maximum of 128 hotwords.
	AsrHotWords []*string `json:"AsrHotWords,omitempty" xml:"AsrHotWords,omitempty" type:"Repeated"`
	// The language for ASR. Valid values:
	//
	// - `zh_mandarin`: Chinese (Mandarin)
	//
	// - `en`: English
	//
	// - `zh_en`: Chinese-English mixed
	//
	// - `es`: Spanish
	//
	// - `jp`: Japanese
	//
	// example:
	//
	// zh_mandarin
	AsrLanguageId *string `json:"AsrLanguageId,omitempty" xml:"AsrLanguageId,omitempty"`
	// The maximum duration of silence in milliseconds before the ASR engine finalizes an utterance. A pause longer than this value signals a sentence break. Range: 200–1200. Default: 400.
	//
	// example:
	//
	// 400
	AsrMaxSilence *int32 `json:"AsrMaxSilence,omitempty" xml:"AsrMaxSilence,omitempty"`
	// Passthrough parameters for proprietary ASR integrations.
	//
	// example:
	//
	// mode=fast&sample=16000&format=wav
	CustomParams *string `json:"CustomParams,omitempty" xml:"CustomParams,omitempty"`
	// The minimum duration in milliseconds of continuous user speech required to trigger an interruption. This controls interruption sensitivity. A value of 0 disables this feature. Range: 200–2000. A common range is 200–500 ms, which typically corresponds to 1 to 4 Chinese characters. If omitted, this feature is disabled.
	//
	// example:
	//
	// 300
	VadDuration *int32 `json:"VadDuration,omitempty" xml:"VadDuration,omitempty"`
	// The Voice Activity Detection (VAD) threshold for interruptions. Range: 0–11. Default: 11.
	//
	// - `0`: Disables VAD.
	//
	// - `1`–`10`: Sets the interruption sensitivity. A higher value makes the agent harder to interrupt.
	//
	// - `11`: An enhanced mode with lower audio distortion and stronger noise resistance.
	//
	// example:
	//
	// 11
	VadLevel *int32 `json:"VadLevel,omitempty" xml:"VadLevel,omitempty"`
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

type AIAgentConfigAutoSpeechConfig struct {
	// Configuration for prompts to play during LLM response latency.
	LlmPending *AIAgentConfigAutoSpeechConfigLlmPending `json:"LlmPending,omitempty" xml:"LlmPending,omitempty" type:"Struct"`
	// Configuration for prompts to play when the user is silent for an extended period.
	UserIdle *AIAgentConfigAutoSpeechConfigUserIdle `json:"UserIdle,omitempty" xml:"UserIdle,omitempty" type:"Struct"`
}

func (s AIAgentConfigAutoSpeechConfig) String() string {
	return dara.Prettify(s)
}

func (s AIAgentConfigAutoSpeechConfig) GoString() string {
	return s.String()
}

func (s *AIAgentConfigAutoSpeechConfig) GetLlmPending() *AIAgentConfigAutoSpeechConfigLlmPending {
	return s.LlmPending
}

func (s *AIAgentConfigAutoSpeechConfig) GetUserIdle() *AIAgentConfigAutoSpeechConfigUserIdle {
	return s.UserIdle
}

func (s *AIAgentConfigAutoSpeechConfig) SetLlmPending(v *AIAgentConfigAutoSpeechConfigLlmPending) *AIAgentConfigAutoSpeechConfig {
	s.LlmPending = v
	return s
}

func (s *AIAgentConfigAutoSpeechConfig) SetUserIdle(v *AIAgentConfigAutoSpeechConfigUserIdle) *AIAgentConfigAutoSpeechConfig {
	s.UserIdle = v
	return s
}

func (s *AIAgentConfigAutoSpeechConfig) Validate() error {
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

type AIAgentConfigAutoSpeechConfigLlmPending struct {
	// A collection of prompt messages. A maximum of 10 messages are supported, each up to 100 characters. The sum of all probabilities must be 100%.
	Messages []*AIAgentConfigAutoSpeechConfigLlmPendingMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Repeated"`
	// The mode for handling LLM latency prompts. `random`: Plays a random message from the list. `sequence`: Plays messages in order. This is a required field.
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The wait time threshold for LLM responses. If the threshold is exceeded, a prompt is played. This is a required field. Unit: ms. Range: 500–10000. Set this value based on the actual performance of your LLM.
	//
	// example:
	//
	// 3000
	WaitTime *int32 `json:"WaitTime,omitempty" xml:"WaitTime,omitempty"`
}

func (s AIAgentConfigAutoSpeechConfigLlmPending) String() string {
	return dara.Prettify(s)
}

func (s AIAgentConfigAutoSpeechConfigLlmPending) GoString() string {
	return s.String()
}

func (s *AIAgentConfigAutoSpeechConfigLlmPending) GetMessages() []*AIAgentConfigAutoSpeechConfigLlmPendingMessages {
	return s.Messages
}

func (s *AIAgentConfigAutoSpeechConfigLlmPending) GetMode() *string {
	return s.Mode
}

func (s *AIAgentConfigAutoSpeechConfigLlmPending) GetWaitTime() *int32 {
	return s.WaitTime
}

func (s *AIAgentConfigAutoSpeechConfigLlmPending) SetMessages(v []*AIAgentConfigAutoSpeechConfigLlmPendingMessages) *AIAgentConfigAutoSpeechConfigLlmPending {
	s.Messages = v
	return s
}

func (s *AIAgentConfigAutoSpeechConfigLlmPending) SetMode(v string) *AIAgentConfigAutoSpeechConfigLlmPending {
	s.Mode = &v
	return s
}

func (s *AIAgentConfigAutoSpeechConfigLlmPending) SetWaitTime(v int32) *AIAgentConfigAutoSpeechConfigLlmPending {
	s.WaitTime = &v
	return s
}

func (s *AIAgentConfigAutoSpeechConfigLlmPending) Validate() error {
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

type AIAgentConfigAutoSpeechConfigLlmPendingMessages struct {
	// The probability of this message being selected. Range: 0–1, corresponding to 0%–100%.
	//
	// example:
	//
	// 0.5
	Probability *float64 `json:"Probability,omitempty" xml:"Probability,omitempty"`
	// The text of the prompt message, up to 100 characters.
	//
	// example:
	//
	// 稍等一下
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s AIAgentConfigAutoSpeechConfigLlmPendingMessages) String() string {
	return dara.Prettify(s)
}

func (s AIAgentConfigAutoSpeechConfigLlmPendingMessages) GoString() string {
	return s.String()
}

func (s *AIAgentConfigAutoSpeechConfigLlmPendingMessages) GetProbability() *float64 {
	return s.Probability
}

func (s *AIAgentConfigAutoSpeechConfigLlmPendingMessages) GetText() *string {
	return s.Text
}

func (s *AIAgentConfigAutoSpeechConfigLlmPendingMessages) SetProbability(v float64) *AIAgentConfigAutoSpeechConfigLlmPendingMessages {
	s.Probability = &v
	return s
}

func (s *AIAgentConfigAutoSpeechConfigLlmPendingMessages) SetText(v string) *AIAgentConfigAutoSpeechConfigLlmPendingMessages {
	s.Text = &v
	return s
}

func (s *AIAgentConfigAutoSpeechConfigLlmPendingMessages) Validate() error {
	return dara.Validate(s)
}

type AIAgentConfigAutoSpeechConfigUserIdle struct {
	// A farewell message played before hanging up due to user inactivity.
	//
	// example:
	//
	// 我挂了，再见
	HangupEndWord *string `json:"HangupEndWord,omitempty" xml:"HangupEndWord,omitempty"`
	// The maximum number of times the prompt can be repeated. Range: 0–10. This is a required field. If the limit is exceeded, the call is terminated.
	//
	// example:
	//
	// 5
	MaxRepeats *int32 `json:"MaxRepeats,omitempty" xml:"MaxRepeats,omitempty"`
	// A collection of prompt messages. A maximum of 10 messages are supported, each up to 100 characters. The sum of all probabilities must be 100%.
	Messages []*AIAgentConfigAutoSpeechConfigUserIdleMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Repeated"`
	// The silence duration threshold in milliseconds. If the user is silent for longer than this period, a prompt is triggered. Range: 5000–600000. This is a required field.
	//
	// example:
	//
	// 5000
	WaitTime *int32 `json:"WaitTime,omitempty" xml:"WaitTime,omitempty"`
}

func (s AIAgentConfigAutoSpeechConfigUserIdle) String() string {
	return dara.Prettify(s)
}

func (s AIAgentConfigAutoSpeechConfigUserIdle) GoString() string {
	return s.String()
}

func (s *AIAgentConfigAutoSpeechConfigUserIdle) GetHangupEndWord() *string {
	return s.HangupEndWord
}

func (s *AIAgentConfigAutoSpeechConfigUserIdle) GetMaxRepeats() *int32 {
	return s.MaxRepeats
}

func (s *AIAgentConfigAutoSpeechConfigUserIdle) GetMessages() []*AIAgentConfigAutoSpeechConfigUserIdleMessages {
	return s.Messages
}

func (s *AIAgentConfigAutoSpeechConfigUserIdle) GetWaitTime() *int32 {
	return s.WaitTime
}

func (s *AIAgentConfigAutoSpeechConfigUserIdle) SetHangupEndWord(v string) *AIAgentConfigAutoSpeechConfigUserIdle {
	s.HangupEndWord = &v
	return s
}

func (s *AIAgentConfigAutoSpeechConfigUserIdle) SetMaxRepeats(v int32) *AIAgentConfigAutoSpeechConfigUserIdle {
	s.MaxRepeats = &v
	return s
}

func (s *AIAgentConfigAutoSpeechConfigUserIdle) SetMessages(v []*AIAgentConfigAutoSpeechConfigUserIdleMessages) *AIAgentConfigAutoSpeechConfigUserIdle {
	s.Messages = v
	return s
}

func (s *AIAgentConfigAutoSpeechConfigUserIdle) SetWaitTime(v int32) *AIAgentConfigAutoSpeechConfigUserIdle {
	s.WaitTime = &v
	return s
}

func (s *AIAgentConfigAutoSpeechConfigUserIdle) Validate() error {
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

type AIAgentConfigAutoSpeechConfigUserIdleMessages struct {
	// The probability of this message being selected. Range: 0–1, corresponding to 0%–100%.
	//
	// example:
	//
	// 0.5
	Probability *float64 `json:"Probability,omitempty" xml:"Probability,omitempty"`
	// The text of the prompt message, up to 100 characters.
	//
	// example:
	//
	// 您还在吗？
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s AIAgentConfigAutoSpeechConfigUserIdleMessages) String() string {
	return dara.Prettify(s)
}

func (s AIAgentConfigAutoSpeechConfigUserIdleMessages) GoString() string {
	return s.String()
}

func (s *AIAgentConfigAutoSpeechConfigUserIdleMessages) GetProbability() *float64 {
	return s.Probability
}

func (s *AIAgentConfigAutoSpeechConfigUserIdleMessages) GetText() *string {
	return s.Text
}

func (s *AIAgentConfigAutoSpeechConfigUserIdleMessages) SetProbability(v float64) *AIAgentConfigAutoSpeechConfigUserIdleMessages {
	s.Probability = &v
	return s
}

func (s *AIAgentConfigAutoSpeechConfigUserIdleMessages) SetText(v string) *AIAgentConfigAutoSpeechConfigUserIdleMessages {
	s.Text = &v
	return s
}

func (s *AIAgentConfigAutoSpeechConfigUserIdleMessages) Validate() error {
	return dara.Validate(s)
}

type AIAgentConfigAvatarConfig struct {
	// The model ID of the avatar.
	//
	// example:
	//
	// 5257
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

type AIAgentConfigBackChannelingConfig struct {
	// 是否启用附和功能。必填，取值 true/false。
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// 功能触发概率。范围 0.0–1.0。必填。
	//
	// example:
	//
	// 0.5
	Probability *float64 `json:"Probability,omitempty" xml:"Probability,omitempty"`
	// 附和触发的时机。可选值：
	//
	// - pause_detected（检测到说话短暂停顿）
	//
	// example:
	//
	// pause_detected
	TriggerStage *string `json:"TriggerStage,omitempty" xml:"TriggerStage,omitempty"`
	// 附和短语集合。最大 10 条，每条短语长度 ≤ 20 字符，概率总和为 1.0。
	Words []*AIAgentConfigBackChannelingConfigWords `json:"Words,omitempty" xml:"Words,omitempty" type:"Repeated"`
}

func (s AIAgentConfigBackChannelingConfig) String() string {
	return dara.Prettify(s)
}

func (s AIAgentConfigBackChannelingConfig) GoString() string {
	return s.String()
}

func (s *AIAgentConfigBackChannelingConfig) GetEnabled() *bool {
	return s.Enabled
}

func (s *AIAgentConfigBackChannelingConfig) GetProbability() *float64 {
	return s.Probability
}

func (s *AIAgentConfigBackChannelingConfig) GetTriggerStage() *string {
	return s.TriggerStage
}

func (s *AIAgentConfigBackChannelingConfig) GetWords() []*AIAgentConfigBackChannelingConfigWords {
	return s.Words
}

func (s *AIAgentConfigBackChannelingConfig) SetEnabled(v bool) *AIAgentConfigBackChannelingConfig {
	s.Enabled = &v
	return s
}

func (s *AIAgentConfigBackChannelingConfig) SetProbability(v float64) *AIAgentConfigBackChannelingConfig {
	s.Probability = &v
	return s
}

func (s *AIAgentConfigBackChannelingConfig) SetTriggerStage(v string) *AIAgentConfigBackChannelingConfig {
	s.TriggerStage = &v
	return s
}

func (s *AIAgentConfigBackChannelingConfig) SetWords(v []*AIAgentConfigBackChannelingConfigWords) *AIAgentConfigBackChannelingConfig {
	s.Words = v
	return s
}

func (s *AIAgentConfigBackChannelingConfig) Validate() error {
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

type AIAgentConfigBackChannelingConfigWords struct {
	// 本短语的触发概率，范围 0.0–1.0，必填。
	//
	// example:
	//
	// 0.3
	Probability *float64 `json:"Probability,omitempty" xml:"Probability,omitempty"`
	// 短语文本，长度 ≤ 20 字符，支持多语言。必填。
	//
	// example:
	//
	// 嗯嗯
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s AIAgentConfigBackChannelingConfigWords) String() string {
	return dara.Prettify(s)
}

func (s AIAgentConfigBackChannelingConfigWords) GoString() string {
	return s.String()
}

func (s *AIAgentConfigBackChannelingConfigWords) GetProbability() *float64 {
	return s.Probability
}

func (s *AIAgentConfigBackChannelingConfigWords) GetText() *string {
	return s.Text
}

func (s *AIAgentConfigBackChannelingConfigWords) SetProbability(v float64) *AIAgentConfigBackChannelingConfigWords {
	s.Probability = &v
	return s
}

func (s *AIAgentConfigBackChannelingConfigWords) SetText(v string) *AIAgentConfigBackChannelingConfigWords {
	s.Text = &v
	return s
}

func (s *AIAgentConfigBackChannelingConfigWords) Validate() error {
	return dara.Validate(s)
}

type AIAgentConfigBackChannelingConfigs struct {
	// Specifies whether to enable this back-channeling rule. This is a required field.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The trigger probability. Range: 0.0–1.0. This is a required field.
	//
	// example:
	//
	// 0.5
	Probability *float64 `json:"Probability,omitempty" xml:"Probability,omitempty"`
	// The trigger for the back-channeling. Valid value:
	//
	// - `pause_detected`: Triggered when a short pause in speech is detected.
	//
	// example:
	//
	// pause_detected
	TriggerStage *string `json:"TriggerStage,omitempty" xml:"TriggerStage,omitempty"`
	// A collection of acknowledgment phrases. You can specify a maximum of 10 phrases. Each phrase must be 20 characters or less, and the sum of their probabilities must be 1.0.
	Words []*AIAgentConfigBackChannelingConfigsWords `json:"Words,omitempty" xml:"Words,omitempty" type:"Repeated"`
}

func (s AIAgentConfigBackChannelingConfigs) String() string {
	return dara.Prettify(s)
}

func (s AIAgentConfigBackChannelingConfigs) GoString() string {
	return s.String()
}

func (s *AIAgentConfigBackChannelingConfigs) GetEnabled() *bool {
	return s.Enabled
}

func (s *AIAgentConfigBackChannelingConfigs) GetProbability() *float64 {
	return s.Probability
}

func (s *AIAgentConfigBackChannelingConfigs) GetTriggerStage() *string {
	return s.TriggerStage
}

func (s *AIAgentConfigBackChannelingConfigs) GetWords() []*AIAgentConfigBackChannelingConfigsWords {
	return s.Words
}

func (s *AIAgentConfigBackChannelingConfigs) SetEnabled(v bool) *AIAgentConfigBackChannelingConfigs {
	s.Enabled = &v
	return s
}

func (s *AIAgentConfigBackChannelingConfigs) SetProbability(v float64) *AIAgentConfigBackChannelingConfigs {
	s.Probability = &v
	return s
}

func (s *AIAgentConfigBackChannelingConfigs) SetTriggerStage(v string) *AIAgentConfigBackChannelingConfigs {
	s.TriggerStage = &v
	return s
}

func (s *AIAgentConfigBackChannelingConfigs) SetWords(v []*AIAgentConfigBackChannelingConfigsWords) *AIAgentConfigBackChannelingConfigs {
	s.Words = v
	return s
}

func (s *AIAgentConfigBackChannelingConfigs) Validate() error {
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

type AIAgentConfigBackChannelingConfigsWords struct {
	// 本短语的触发概率，范围 0.0–1.0，必填。
	//
	// example:
	//
	// 0.3
	Probability *float64 `json:"Probability,omitempty" xml:"Probability,omitempty"`
	// 短语文本，长度 ≤ 20 字符，支持多语言。必填。
	//
	// example:
	//
	// 嗯嗯
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s AIAgentConfigBackChannelingConfigsWords) String() string {
	return dara.Prettify(s)
}

func (s AIAgentConfigBackChannelingConfigsWords) GoString() string {
	return s.String()
}

func (s *AIAgentConfigBackChannelingConfigsWords) GetProbability() *float64 {
	return s.Probability
}

func (s *AIAgentConfigBackChannelingConfigsWords) GetText() *string {
	return s.Text
}

func (s *AIAgentConfigBackChannelingConfigsWords) SetProbability(v float64) *AIAgentConfigBackChannelingConfigsWords {
	s.Probability = &v
	return s
}

func (s *AIAgentConfigBackChannelingConfigsWords) SetText(v string) *AIAgentConfigBackChannelingConfigsWords {
	s.Text = &v
	return s
}

func (s *AIAgentConfigBackChannelingConfigsWords) Validate() error {
	return dara.Validate(s)
}

type AIAgentConfigInterruptConfig struct {
	// Specifies whether to enable speech interruption. Default: `true`.
	//
	// example:
	//
	// true
	EnableVoiceInterrupt *bool `json:"EnableVoiceInterrupt,omitempty" xml:"EnableVoiceInterrupt,omitempty"`
	// A list of specific words or phrases that trigger an interruption.
	InterruptWords []*string `json:"InterruptWords,omitempty" xml:"InterruptWords,omitempty" type:"Repeated"`
	// Specifies whether to include the interrupt words in the text sent to the LLM. Default: `false` (words are discarded).
	//
	// > For example, if "hold on" is an interrupt word and the user says "hold on, what is the weather like today?", setting this to `false` results in only "what is the weather like today?" being sent to the LLM.
	//
	// example:
	//
	// true
	KeepInterruptWordsForLLM *bool `json:"KeepInterruptWordsForLLM,omitempty" xml:"KeepInterruptWordsForLLM,omitempty"`
	// Specifies how to handle user speech that occurs during a non-interruptible section of the agent\\"s utterance.
	//
	// - `cache`: Caches the user\\"s speech and processes it in the next conversational turn.
	//
	// - `discard`: Discards the user\\"s speech.
	//
	// Default: `cache`.
	//
	// example:
	//
	// cache
	NoInterruptMode *string `json:"NoInterruptMode,omitempty" xml:"NoInterruptMode,omitempty"`
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

func (s *AIAgentConfigInterruptConfig) GetKeepInterruptWordsForLLM() *bool {
	return s.KeepInterruptWordsForLLM
}

func (s *AIAgentConfigInterruptConfig) GetNoInterruptMode() *string {
	return s.NoInterruptMode
}

func (s *AIAgentConfigInterruptConfig) SetEnableVoiceInterrupt(v bool) *AIAgentConfigInterruptConfig {
	s.EnableVoiceInterrupt = &v
	return s
}

func (s *AIAgentConfigInterruptConfig) SetInterruptWords(v []*string) *AIAgentConfigInterruptConfig {
	s.InterruptWords = v
	return s
}

func (s *AIAgentConfigInterruptConfig) SetKeepInterruptWordsForLLM(v bool) *AIAgentConfigInterruptConfig {
	s.KeepInterruptWordsForLLM = &v
	return s
}

func (s *AIAgentConfigInterruptConfig) SetNoInterruptMode(v string) *AIAgentConfigInterruptConfig {
	s.NoInterruptMode = &v
	return s
}

func (s *AIAgentConfigInterruptConfig) Validate() error {
	return dara.Validate(s)
}

type AIAgentConfigLlmConfig struct {
	// Parameters for Alibaba Cloud Model Studio, provided as a JSON string. For the parameter format, see
	//
	// [Alibaba Cloud Model Studio Parameters](https://help.aliyun.com/document_detail/2858132.html)
	//
	// example:
	//
	// "{\\"biz_params\\":{\\"user_defined_params\\":{\\"your_plugin_id\\":{\\"article_index\\":2}}},\\"memory_id\\":\\"your_memory_id\\",\\"image_list\\":[\\"https://your_image_url\\"],\\"rag_options\\":{\\"pipeline_ids\\":[\\"your_id\\"],\\"file_ids\\":[\\"文档ID1\\",\\"文档ID2\\"],\\"metadata_filter\\":{\\"name\\":\\"张三\\"},\\"structured_filter\\":{\\"key1\\":\\"value1\\",\\"key2\\":\\"value2\\"},\\"tags\\":[\\"标签1\\",\\"标签2\\"]}}"
	BailianAppParams *string `json:"BailianAppParams,omitempty" xml:"BailianAppParams,omitempty"`
	// Maps built-in agent functions to custom LLM functions. Currently, this only supports function calling for custom, OpenAI-compatible LLMs.
	FunctionMap []*AIAgentConfigLlmConfigFunctionMap `json:"FunctionMap,omitempty" xml:"FunctionMap,omitempty" type:"Repeated"`
	// Specifies whether the LLM message history is synchronized with the content played by the TTS. Default: `false`. When enabled, the saved LLM messages match the content actually played by the TTS.
	//
	// > When a user interrupts the agent, the `<ims_agent_interrupted>` tag is inserted into the message history at the point of interruption. This affects the next message sent to the LLM. For example:
	//
	// ```
	//
	// [
	//
	//   {"role": "user", "content": "Tell me a story."},
	//
	//   {"role": "assistant", "content": "Okay, I can tell you a story about the Three Kingdoms. Would you<ims_agent_interrupted> like that?"},
	//
	//   {"role": "user", "content": "Tell me a different one."}
	//
	// ]
	//
	// ```
	//
	// example:
	//
	// false
	HistorySyncWithTTS *bool `json:"HistorySyncWithTTS,omitempty" xml:"HistorySyncWithTTS,omitempty"`
	// When set to `true`, the agent sends the entire LLM response in a single message after it is fully generated, rather than streaming it. This setting does not affect the streaming of subtitles.
	//
	// example:
	//
	// true
	LlmCompleteReply *bool `json:"LlmCompleteReply,omitempty" xml:"LlmCompleteReply,omitempty"`
	// The conversation history context for the LLM/MLLM.
	LlmHistory []*AIAgentConfigLlmConfigLlmHistory `json:"LlmHistory,omitempty" xml:"LlmHistory,omitempty" type:"Repeated"`
	// The maximum number of recent conversational turns to include in the LLM/MLLM context. Default: 10.
	//
	// example:
	//
	// 10
	LlmHistoryLimit *int32 `json:"LlmHistoryLimit,omitempty" xml:"LlmHistoryLimit,omitempty"`
	// The system prompt for the LLM after the call starts.
	//
	// example:
	//
	// 你是一位友好且乐于助人的助手，专注于为用户提供准确的信息和建议。
	LlmSystemPrompt *string `json:"LlmSystemPrompt,omitempty" xml:"LlmSystemPrompt,omitempty"`
	// Additional query parameters for an OpenAI-compatible LLM. Parameters must be provided as a URL query string (e.g., `key1=value1&key2=value2`). All values must be strings.
	//
	// example:
	//
	// api-version=2024-02-01&api-key=sk-xxx
	OpenAIExtraQuery *string `json:"OpenAIExtraQuery,omitempty" xml:"OpenAIExtraQuery,omitempty"`
	// The maximum delay in milliseconds before buffered text is sent to the TTS engine, even if `OutputMinLength` is not met. Range: 1000–10000. A value of `0` or omitting this parameter disables the delay limit. Default: Not set.
	//
	// example:
	//
	// 2000
	OutputMaxDelay *int32 `json:"OutputMaxDelay,omitempty" xml:"OutputMaxDelay,omitempty"`
	// The minimum number of characters in a text chunk before it is sent to the TTS engine. Shorter chunks are buffered. Range: 0–100. A value of `0` or omitting this parameter disables buffering. Default: Not set.
	//
	// example:
	//
	// 5
	OutputMinLength *int32 `json:"OutputMinLength,omitempty" xml:"OutputMinLength,omitempty"`
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

func (s *AIAgentConfigLlmConfig) GetHistorySyncWithTTS() *bool {
	return s.HistorySyncWithTTS
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

func (s *AIAgentConfigLlmConfig) SetHistorySyncWithTTS(v bool) *AIAgentConfigLlmConfig {
	s.HistorySyncWithTTS = &v
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

type AIAgentConfigLlmConfigFunctionMap struct {
	// The name of a built-in function provided by the AI agent system. Currently, only `hangup` is supported.
	//
	// example:
	//
	// hangup
	Function *string `json:"Function,omitempty" xml:"Function,omitempty"`
	// The name of the custom LLM function that maps to the agent\\"s built-in function. For details on the custom LLM protocol, see [LLM Standard Interface](https://help.aliyun.com/document_detail/2839094.html).
	//
	// example:
	//
	// hangup
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
	// The text content of the message from this role.
	//
	// example:
	//
	// 你好
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The role of the participant in the conversation. Valid values:
	//
	// - `user`
	//
	// - `assistant`
	//
	// - `system`
	//
	// - `function`
	//
	// - `plugin`
	//
	// - `tool`
	//
	// example:
	//
	// user
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
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
	// This parameter applies only to the Minimax provider. Supported emotions include:
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
	// This parameter is for the minimax provider only. It enhances recognition for specific low-resource languages and dialects. If the language is unknown, set this to `auto` for automatic detection. By default, this parameter is not set. Supported values include:
	//
	// <details>
	//
	// <summary>
	//
	// Supported languages
	//
	// </summary>
	//
	// - Chinese
	//
	// - Chinese,Yue: Cantonese
	//
	// - English
	//
	// - Arabic
	//
	// - Russian
	//
	// - Spanish
	//
	// - French
	//
	// - Portuguese
	//
	// - German
	//
	// - Turkish
	//
	// - Dutch
	//
	// - Ukrainian
	//
	// - Vietnamese
	//
	// - Indonesian
	//
	// - Japanese
	//
	// - Italian
	//
	// - Korean
	//
	// - Thai
	//
	// - Polish
	//
	// - Romanian
	//
	// - Greek
	//
	// - Czech
	//
	// - Finnish
	//
	// - Hindi
	//
	// - auto
	//
	// </details>
	//
	// example:
	//
	// Chinese
	LanguageId *string `json:"LanguageId,omitempty" xml:"LanguageId,omitempty"`
	// This parameter applies only to the Minimax provider. Valid values:
	//
	// `speech-01-turbo`, `speech-02-turbo`
	//
	// example:
	//
	// speech-01-turbo
	ModelId *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	// A list of TTS pronunciation rules, executed in order. You can specify a maximum of 20 rules.
	PronunciationRules []*AIAgentConfigTtsConfigPronunciationRules `json:"PronunciationRules,omitempty" xml:"PronunciationRules,omitempty" type:"Repeated"`
	// The speech rate, where a value of 1.0 is normal speed. The supported range can vary by provider. For CosyVoice, the range is 0.5 to 2.0 (default: 1.0). For Minimax, the range is 0.5 to 2.0 (default: 1.0).
	//
	// example:
	//
	// 1.0
	SpeechRate *float64 `json:"SpeechRate,omitempty" xml:"SpeechRate,omitempty"`
	// The ID of the preset TTS voice. Changes apply to the next utterance. If omitted, the voice from the AI agent template is used. The ID can be a maximum of 64 characters. For available voices, see [Intelligent Voice Samples](https://help.aliyun.com/document_detail/449563.html).
	//
	// example:
	//
	// longcheng_v2
	VoiceId *string `json:"VoiceId,omitempty" xml:"VoiceId,omitempty"`
	// A list of available voices.
	VoiceIdList []*string `json:"VoiceIdList,omitempty" xml:"VoiceIdList,omitempty" type:"Repeated"`
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

type AIAgentConfigTtsConfigPronunciationRules struct {
	// The replacement pronunciation. It must be 1 to 9 Chinese characters long and cannot contain spaces.
	//
	// example:
	//
	// 幺幺零
	Pronunciation *string `json:"Pronunciation,omitempty" xml:"Pronunciation,omitempty"`
	// The type of pronunciation rule.
	//
	// Valid value:
	//
	// - `replacement`: Replaces the specified `Word` with the `Pronunciation`.
	//
	// example:
	//
	// replacement
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The word to be replaced. It must be 1 to 9 Chinese characters long and cannot contain spaces.
	//
	// example:
	//
	// 一一零
	Word *string `json:"Word,omitempty" xml:"Word,omitempty"`
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
	// Controls the agent\\"s response speed after detecting a user pause. This parameter applies only in `Semantic` mode. A higher setting results in a faster response but increases the risk of interrupting the user:
	//
	// - `Low`: Waits patiently with a maximum wait time of 6 seconds, reducing the risk of interruption.
	//
	// - `Medium`: A balanced wait time (up to 4 seconds), suitable for most scenarios.
	//
	// - `High`: Responds quickly (up to 2 seconds), which improves speed but may increase the risk of incorrect turn-taking.
	//
	// This field is empty by default.
	//
	// example:
	//
	// High
	Eagerness *string `json:"Eagerness,omitempty" xml:"Eagerness,omitempty"`
	// The conversational turn detection mode.
	//
	// - `Normal` (Default): The agent relies on pauses to detect the end of a user\\"s turn.
	//
	// - `Semantic`: The agent uses AI to analyze conversational context to determine if the user has finished speaking.
	//
	// example:
	//
	// Semantic
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The pause detection time in AI mode, in milliseconds. Default: -1.
	//
	// - -1: The AI automatically determines a suitable wait time.
	//
	// - 0–10000: A custom wait time. A range of 0–1500 ms is recommended.
	//
	// > This parameter has no effect in `Normal` mode.
	//
	// example:
	//
	// -1
	SemanticWaitDuration *int32 `json:"SemanticWaitDuration,omitempty" xml:"SemanticWaitDuration,omitempty"`
	// A list of keywords used to determine the end of a user\\"s conversational turn.
	TurnEndWords []*string `json:"TurnEndWords,omitempty" xml:"TurnEndWords,omitempty" type:"Repeated"`
}

func (s AIAgentConfigTurnDetectionConfig) String() string {
	return dara.Prettify(s)
}

func (s AIAgentConfigTurnDetectionConfig) GoString() string {
	return s.String()
}

func (s *AIAgentConfigTurnDetectionConfig) GetEagerness() *string {
	return s.Eagerness
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

func (s *AIAgentConfigTurnDetectionConfig) SetEagerness(v string) *AIAgentConfigTurnDetectionConfig {
	s.Eagerness = &v
	return s
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
	// Configuration for device identification.
	Equipment *AIAgentConfigVcrConfigEquipment `json:"Equipment,omitempty" xml:"Equipment,omitempty" type:"Struct"`
	// Configuration for head motion detection.
	HeadMotion *AIAgentConfigVcrConfigHeadMotion `json:"HeadMotion,omitempty" xml:"HeadMotion,omitempty" type:"Struct"`
	// Configuration for invalid frame detection.
	InvalidFrameMotion *AIAgentConfigVcrConfigInvalidFrameMotion `json:"InvalidFrameMotion,omitempty" xml:"InvalidFrameMotion,omitempty" type:"Struct"`
	// Configuration for look-away detection.
	LookAway *AIAgentConfigVcrConfigLookAway `json:"LookAway,omitempty" xml:"LookAway,omitempty" type:"Struct"`
	// Configuration for the people counting feature.
	PeopleCount *AIAgentConfigVcrConfigPeopleCount `json:"PeopleCount,omitempty" xml:"PeopleCount,omitempty" type:"Struct"`
	// Configuration for still frame detection.
	StillFrameMotion *AIAgentConfigVcrConfigStillFrameMotion `json:"StillFrameMotion,omitempty" xml:"StillFrameMotion,omitempty" type:"Struct"`
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
	if s.Equipment != nil {
		if err := s.Equipment.Validate(); err != nil {
			return err
		}
	}
	if s.HeadMotion != nil {
		if err := s.HeadMotion.Validate(); err != nil {
			return err
		}
	}
	if s.InvalidFrameMotion != nil {
		if err := s.InvalidFrameMotion.Validate(); err != nil {
			return err
		}
	}
	if s.LookAway != nil {
		if err := s.LookAway.Validate(); err != nil {
			return err
		}
	}
	if s.PeopleCount != nil {
		if err := s.PeopleCount.Validate(); err != nil {
			return err
		}
	}
	if s.StillFrameMotion != nil {
		if err := s.StillFrameMotion.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AIAgentConfigVcrConfigEquipment struct {
	// Specifies whether to enable device identification. Default: `false`.
	//
	// example:
	//
	// false
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
	// Specifies whether to enable head motion detection. Default: `false`.
	//
	// example:
	//
	// false
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
	// The duration in milliseconds that an invalid frame must persist before a notification is sent. If not specified, the setting from the console is used. Range: 200–5000.
	//
	// example:
	//
	// 3000
	CallbackDelay *int32 `json:"CallbackDelay,omitempty" xml:"CallbackDelay,omitempty"`
	// Specifies whether to enable invalid frame detection.
	//
	// example:
	//
	// false
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
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
	// Specifies whether to enable look-away detection. Default: `false`.
	//
	// example:
	//
	// true
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
	// Specifies whether to enable people counting. Default: `false`.
	//
	// example:
	//
	// false
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
	// The duration in milliseconds that a frame must remain still before a notification is sent. If not specified, the setting from the console is used. Range: 200–5000.
	//
	// example:
	//
	// 3000
	CallbackDelay *int32 `json:"CallbackDelay,omitempty" xml:"CallbackDelay,omitempty"`
	// Specifies whether to enable still frame detection. Default: `false`.
	//
	// example:
	//
	// false
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
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
	// The voiceprint registration mode. Default: `Explicit`.
	//
	// | Value      | Description                                                                                                         |
	//
	// | ---------- | ------------------------------------------------------------------------------------------------------------------- |
	//
	// | `Explicit` | In `Explicit` mode, the user must register their voiceprint in advance by using the voiceprint registration API.    |
	//
	// | `Implicit` | In `Implicit` mode, the system automatically collects user speech during the conversation to register a voiceprint. |
	//
	// example:
	//
	// Explicit
	RegistrationMode *string `json:"RegistrationMode,omitempty" xml:"RegistrationMode,omitempty"`
	// Specifies whether to enable voiceprint recognition. Default: `false`. If set to `true`, you must also provide a valid `VoiceprintId`.
	//
	// example:
	//
	// false
	UseVoiceprint *bool `json:"UseVoiceprint,omitempty" xml:"UseVoiceprint,omitempty"`
	// The unique identifier for the voiceprint. This is not set by default. The ID must correspond to a voiceprint registered using the voiceprint registration API. For more information, see [Register a voiceprint](https://help.aliyun.com/document_detail/2964738.html).
	//
	// example:
	//
	// zhixiaoxia
	VoiceprintId *string `json:"VoiceprintId,omitempty" xml:"VoiceprintId,omitempty"`
}

func (s AIAgentConfigVoiceprintConfig) String() string {
	return dara.Prettify(s)
}

func (s AIAgentConfigVoiceprintConfig) GoString() string {
	return s.String()
}

func (s *AIAgentConfigVoiceprintConfig) GetRegistrationMode() *string {
	return s.RegistrationMode
}

func (s *AIAgentConfigVoiceprintConfig) GetUseVoiceprint() *bool {
	return s.UseVoiceprint
}

func (s *AIAgentConfigVoiceprintConfig) GetVoiceprintId() *string {
	return s.VoiceprintId
}

func (s *AIAgentConfigVoiceprintConfig) SetRegistrationMode(v string) *AIAgentConfigVoiceprintConfig {
	s.RegistrationMode = &v
	return s
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
