// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAIAgentTemplateConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAvatarChat3D(v *AIAgentTemplateConfigAvatarChat3D) *AIAgentTemplateConfig
	GetAvatarChat3D() *AIAgentTemplateConfigAvatarChat3D
	SetVisionChat(v *AIAgentTemplateConfigVisionChat) *AIAgentTemplateConfig
	GetVisionChat() *AIAgentTemplateConfigVisionChat
	SetVoiceChat(v *AIAgentTemplateConfigVoiceChat) *AIAgentTemplateConfig
	GetVoiceChat() *AIAgentTemplateConfigVoiceChat
}

type AIAgentTemplateConfig struct {
	// 3D avatar parameters.
	AvatarChat3D *AIAgentTemplateConfigAvatarChat3D `json:"AvatarChat3D,omitempty" xml:"AvatarChat3D,omitempty" type:"Struct"`
	// Vision agent parameters.
	VisionChat *AIAgentTemplateConfigVisionChat `json:"VisionChat,omitempty" xml:"VisionChat,omitempty" type:"Struct"`
	// Voice chat parameters.
	VoiceChat *AIAgentTemplateConfigVoiceChat `json:"VoiceChat,omitempty" xml:"VoiceChat,omitempty" type:"Struct"`
}

func (s AIAgentTemplateConfig) String() string {
	return dara.Prettify(s)
}

func (s AIAgentTemplateConfig) GoString() string {
	return s.String()
}

func (s *AIAgentTemplateConfig) GetAvatarChat3D() *AIAgentTemplateConfigAvatarChat3D {
	return s.AvatarChat3D
}

func (s *AIAgentTemplateConfig) GetVisionChat() *AIAgentTemplateConfigVisionChat {
	return s.VisionChat
}

func (s *AIAgentTemplateConfig) GetVoiceChat() *AIAgentTemplateConfigVoiceChat {
	return s.VoiceChat
}

func (s *AIAgentTemplateConfig) SetAvatarChat3D(v *AIAgentTemplateConfigAvatarChat3D) *AIAgentTemplateConfig {
	s.AvatarChat3D = v
	return s
}

func (s *AIAgentTemplateConfig) SetVisionChat(v *AIAgentTemplateConfigVisionChat) *AIAgentTemplateConfig {
	s.VisionChat = v
	return s
}

func (s *AIAgentTemplateConfig) SetVoiceChat(v *AIAgentTemplateConfigVoiceChat) *AIAgentTemplateConfig {
	s.VoiceChat = v
	return s
}

func (s *AIAgentTemplateConfig) Validate() error {
	if s.AvatarChat3D != nil {
		if err := s.AvatarChat3D.Validate(); err != nil {
			return err
		}
	}
	if s.VisionChat != nil {
		if err := s.VisionChat.Validate(); err != nil {
			return err
		}
	}
	if s.VoiceChat != nil {
		if err := s.VoiceChat.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AIAgentTemplateConfigAvatarChat3D struct {
	// A list of hot words to improve ASR accuracy. A maximum of 128 words is supported.
	AsrHotWords []*string `json:"AsrHotWords,omitempty" xml:"AsrHotWords,omitempty" type:"Repeated"`
	// The language ID for Automatic Speech Recognition (ASR). Possible values:
	//
	// - `zh_mandarin`: Chinese
	//
	// - `en`: English
	//
	// - `zh_en`: Chinese-English
	//
	// - `es`: Spanish
	//
	// - `jp`: Japanese
	//
	// example:
	//
	// zh_mandarin
	AsrLanguageId *string `json:"AsrLanguageId,omitempty" xml:"AsrLanguageId,omitempty"`
	// The maximum duration of silence in milliseconds before a sentence break is detected. Range: 200 to 1,200. Default: 400.
	//
	// example:
	//
	// 400
	AsrMaxSilence *int32 `json:"AsrMaxSilence,omitempty" xml:"AsrMaxSilence,omitempty"`
	// The ID of the avatar model.
	//
	// example:
	//
	// 1231
	AvatarId *string `json:"AvatarId,omitempty" xml:"AvatarId,omitempty"`
	// Parameters for Alibaba Cloud Bailian. For details, see [Bailian App Params](https://help.aliyun.com/document_detail/2858132.html).
	//
	// example:
	//
	// {}
	BailianAppParams *string `json:"BailianAppParams,omitempty" xml:"BailianAppParams,omitempty"`
	CharBreak        *bool   `json:"CharBreak,omitempty" xml:"CharBreak,omitempty"`
	// Specifies whether to enable intelligent segmentation. If enabled, this feature intelligently merges pauses in a user\\"s speech into a single, complete sentence. Default: `true`.
	//
	// example:
	//
	// true
	EnableIntelligentSegment *bool `json:"EnableIntelligentSegment,omitempty" xml:"EnableIntelligentSegment,omitempty"`
	// Specifies whether to enable Push-to-Talk mode. Default: `false`.
	//
	// example:
	//
	// false
	EnablePushToTalk *bool `json:"EnablePushToTalk,omitempty" xml:"EnablePushToTalk,omitempty"`
	// Specifies whether to enable voice interruption. Default: `true`.
	//
	// example:
	//
	// true
	EnableVoiceInterrupt *bool `json:"EnableVoiceInterrupt,omitempty" xml:"EnableVoiceInterrupt,omitempty"`
	// Specifies whether to enable graceful shutdown. Default: `false`.
	//
	// - If enabled, the agent finishes its current speech (up to 10 seconds) before stopping.
	//
	// example:
	//
	// false
	GracefulShutdown *bool `json:"GracefulShutdown,omitempty" xml:"GracefulShutdown,omitempty"`
	// The greeting message delivered when a user joins the session. If this parameter is omitted, the greeting configured in the agent template is used. Maximum length: 128 characters.
	//
	// example:
	//
	// 早上好，我的朋友！
	Greeting *string `json:"Greeting,omitempty" xml:"Greeting,omitempty"`
	// A list of specific words or phrases that trigger a conversation interruption.
	InterruptWords []*string `json:"InterruptWords,omitempty" xml:"InterruptWords,omitempty" type:"Repeated"`
	// The LLM/MLLM conversation history.
	LlmHistory []*AIAgentTemplateConfigAvatarChat3DLlmHistory `json:"LlmHistory,omitempty" xml:"LlmHistory,omitempty" type:"Repeated"`
	// The maximum number of conversation turns to retain in the LLM/MLLM history. Default: 10.
	//
	// example:
	//
	// 10
	LlmHistoryLimit *int32 `json:"LlmHistoryLimit,omitempty" xml:"LlmHistoryLimit,omitempty"`
	// The system prompt for the LLM, applied when the call starts.
	//
	// example:
	//
	// 你是一位友好且乐于助人的助手，专注于为用户提供准确的信息和建议。
	LlmSystemPrompt *string `json:"LlmSystemPrompt,omitempty" xml:"LlmSystemPrompt,omitempty"`
	// The maximum idle time in seconds with no interaction before the agent goes offline. Default: 600.
	//
	// example:
	//
	// 600
	MaxIdleTime *int32 `json:"MaxIdleTime,omitempty" xml:"MaxIdleTime,omitempty"`
	// Specifies whether to use voiceprint recognition. Default: `false`.
	//
	// example:
	//
	// false
	UseVoiceprint *bool `json:"UseVoiceprint,omitempty" xml:"UseVoiceprint,omitempty"`
	// The time in seconds that the agent waits after a user leaves before closing the task. Default: 5.
	//
	// example:
	//
	// 5
	UserOfflineTimeout *int32 `json:"UserOfflineTimeout,omitempty" xml:"UserOfflineTimeout,omitempty"`
	// The time in seconds that the agent waits for a user to join before closing the task. Default: 60.
	//
	// example:
	//
	// 60
	UserOnlineTimeout *int32 `json:"UserOnlineTimeout,omitempty" xml:"UserOnlineTimeout,omitempty"`
	// The interruption sensitivity threshold. A higher value makes it more difficult to interrupt the agent. Range: 0 to 11. Default: 11.
	//
	// - `0`: Disables VAD.
	//
	// - `1` to `10`: A higher value makes it more difficult to interrupt the agent.
	//
	// - `11`: Offers lower audio distortion and stronger resistance to interference.
	//
	// example:
	//
	// 0
	VadLevel *int32 `json:"VadLevel,omitempty" xml:"VadLevel,omitempty"`
	// The ID of the Text-to-Speech (TTS) voice. Changes take effect on the next utterance. If omitted, the default voice from the agent template is used. This parameter applies only to preset TTS voices. Maximum length: 64 characters. For available values, see [Intelligent voice effect samples](https://help.aliyun.com/document_detail/449563.html).
	//
	// example:
	//
	// zhixiaoxia
	VoiceId *string `json:"VoiceId,omitempty" xml:"VoiceId,omitempty"`
	// A list of available voices.
	VoiceIdList []*string `json:"VoiceIdList,omitempty" xml:"VoiceIdList,omitempty" type:"Repeated"`
	// The unique ID for voiceprint recognition. Default: not specified.
	//
	// example:
	//
	// uniqueId
	VoiceprintId *string `json:"VoiceprintId,omitempty" xml:"VoiceprintId,omitempty"`
	// The speaking volume of the agent.
	//
	// - If omitted, the system uses adaptive volume mode.
	//
	// - If specified, the valid range is 0 to 400. The output volume is calculated as: `Output Volume in Workflow` \\	- (`volume`/100). For example:
	//
	// 1. If `volume` is `0`, the output is silent.
	//
	// 2. If `volume` is `100`, the output volume is the original volume.
	//
	// 3. If `volume` is `200`, the output volume is twice the original volume.
	//
	// example:
	//
	// 100
	Volume *int64 `json:"Volume,omitempty" xml:"Volume,omitempty"`
	// An initial user query that the agent addresses immediately when the call starts.
	//
	// example:
	//
	// 今天天气怎么样？
	WakeUpQuery *string `json:"WakeUpQuery,omitempty" xml:"WakeUpQuery,omitempty"`
	// Workflow override parameters. Default: empty.
	//
	// example:
	//
	// {}
	WorkflowOverrideParams *string `json:"WorkflowOverrideParams,omitempty" xml:"WorkflowOverrideParams,omitempty"`
}

func (s AIAgentTemplateConfigAvatarChat3D) String() string {
	return dara.Prettify(s)
}

func (s AIAgentTemplateConfigAvatarChat3D) GoString() string {
	return s.String()
}

func (s *AIAgentTemplateConfigAvatarChat3D) GetAsrHotWords() []*string {
	return s.AsrHotWords
}

func (s *AIAgentTemplateConfigAvatarChat3D) GetAsrLanguageId() *string {
	return s.AsrLanguageId
}

func (s *AIAgentTemplateConfigAvatarChat3D) GetAsrMaxSilence() *int32 {
	return s.AsrMaxSilence
}

func (s *AIAgentTemplateConfigAvatarChat3D) GetAvatarId() *string {
	return s.AvatarId
}

func (s *AIAgentTemplateConfigAvatarChat3D) GetBailianAppParams() *string {
	return s.BailianAppParams
}

func (s *AIAgentTemplateConfigAvatarChat3D) GetCharBreak() *bool {
	return s.CharBreak
}

func (s *AIAgentTemplateConfigAvatarChat3D) GetEnableIntelligentSegment() *bool {
	return s.EnableIntelligentSegment
}

func (s *AIAgentTemplateConfigAvatarChat3D) GetEnablePushToTalk() *bool {
	return s.EnablePushToTalk
}

func (s *AIAgentTemplateConfigAvatarChat3D) GetEnableVoiceInterrupt() *bool {
	return s.EnableVoiceInterrupt
}

func (s *AIAgentTemplateConfigAvatarChat3D) GetGracefulShutdown() *bool {
	return s.GracefulShutdown
}

func (s *AIAgentTemplateConfigAvatarChat3D) GetGreeting() *string {
	return s.Greeting
}

func (s *AIAgentTemplateConfigAvatarChat3D) GetInterruptWords() []*string {
	return s.InterruptWords
}

func (s *AIAgentTemplateConfigAvatarChat3D) GetLlmHistory() []*AIAgentTemplateConfigAvatarChat3DLlmHistory {
	return s.LlmHistory
}

func (s *AIAgentTemplateConfigAvatarChat3D) GetLlmHistoryLimit() *int32 {
	return s.LlmHistoryLimit
}

func (s *AIAgentTemplateConfigAvatarChat3D) GetLlmSystemPrompt() *string {
	return s.LlmSystemPrompt
}

func (s *AIAgentTemplateConfigAvatarChat3D) GetMaxIdleTime() *int32 {
	return s.MaxIdleTime
}

func (s *AIAgentTemplateConfigAvatarChat3D) GetUseVoiceprint() *bool {
	return s.UseVoiceprint
}

func (s *AIAgentTemplateConfigAvatarChat3D) GetUserOfflineTimeout() *int32 {
	return s.UserOfflineTimeout
}

func (s *AIAgentTemplateConfigAvatarChat3D) GetUserOnlineTimeout() *int32 {
	return s.UserOnlineTimeout
}

func (s *AIAgentTemplateConfigAvatarChat3D) GetVadLevel() *int32 {
	return s.VadLevel
}

func (s *AIAgentTemplateConfigAvatarChat3D) GetVoiceId() *string {
	return s.VoiceId
}

func (s *AIAgentTemplateConfigAvatarChat3D) GetVoiceIdList() []*string {
	return s.VoiceIdList
}

func (s *AIAgentTemplateConfigAvatarChat3D) GetVoiceprintId() *string {
	return s.VoiceprintId
}

func (s *AIAgentTemplateConfigAvatarChat3D) GetVolume() *int64 {
	return s.Volume
}

func (s *AIAgentTemplateConfigAvatarChat3D) GetWakeUpQuery() *string {
	return s.WakeUpQuery
}

func (s *AIAgentTemplateConfigAvatarChat3D) GetWorkflowOverrideParams() *string {
	return s.WorkflowOverrideParams
}

func (s *AIAgentTemplateConfigAvatarChat3D) SetAsrHotWords(v []*string) *AIAgentTemplateConfigAvatarChat3D {
	s.AsrHotWords = v
	return s
}

func (s *AIAgentTemplateConfigAvatarChat3D) SetAsrLanguageId(v string) *AIAgentTemplateConfigAvatarChat3D {
	s.AsrLanguageId = &v
	return s
}

func (s *AIAgentTemplateConfigAvatarChat3D) SetAsrMaxSilence(v int32) *AIAgentTemplateConfigAvatarChat3D {
	s.AsrMaxSilence = &v
	return s
}

func (s *AIAgentTemplateConfigAvatarChat3D) SetAvatarId(v string) *AIAgentTemplateConfigAvatarChat3D {
	s.AvatarId = &v
	return s
}

func (s *AIAgentTemplateConfigAvatarChat3D) SetBailianAppParams(v string) *AIAgentTemplateConfigAvatarChat3D {
	s.BailianAppParams = &v
	return s
}

func (s *AIAgentTemplateConfigAvatarChat3D) SetCharBreak(v bool) *AIAgentTemplateConfigAvatarChat3D {
	s.CharBreak = &v
	return s
}

func (s *AIAgentTemplateConfigAvatarChat3D) SetEnableIntelligentSegment(v bool) *AIAgentTemplateConfigAvatarChat3D {
	s.EnableIntelligentSegment = &v
	return s
}

func (s *AIAgentTemplateConfigAvatarChat3D) SetEnablePushToTalk(v bool) *AIAgentTemplateConfigAvatarChat3D {
	s.EnablePushToTalk = &v
	return s
}

func (s *AIAgentTemplateConfigAvatarChat3D) SetEnableVoiceInterrupt(v bool) *AIAgentTemplateConfigAvatarChat3D {
	s.EnableVoiceInterrupt = &v
	return s
}

func (s *AIAgentTemplateConfigAvatarChat3D) SetGracefulShutdown(v bool) *AIAgentTemplateConfigAvatarChat3D {
	s.GracefulShutdown = &v
	return s
}

func (s *AIAgentTemplateConfigAvatarChat3D) SetGreeting(v string) *AIAgentTemplateConfigAvatarChat3D {
	s.Greeting = &v
	return s
}

func (s *AIAgentTemplateConfigAvatarChat3D) SetInterruptWords(v []*string) *AIAgentTemplateConfigAvatarChat3D {
	s.InterruptWords = v
	return s
}

func (s *AIAgentTemplateConfigAvatarChat3D) SetLlmHistory(v []*AIAgentTemplateConfigAvatarChat3DLlmHistory) *AIAgentTemplateConfigAvatarChat3D {
	s.LlmHistory = v
	return s
}

func (s *AIAgentTemplateConfigAvatarChat3D) SetLlmHistoryLimit(v int32) *AIAgentTemplateConfigAvatarChat3D {
	s.LlmHistoryLimit = &v
	return s
}

func (s *AIAgentTemplateConfigAvatarChat3D) SetLlmSystemPrompt(v string) *AIAgentTemplateConfigAvatarChat3D {
	s.LlmSystemPrompt = &v
	return s
}

func (s *AIAgentTemplateConfigAvatarChat3D) SetMaxIdleTime(v int32) *AIAgentTemplateConfigAvatarChat3D {
	s.MaxIdleTime = &v
	return s
}

func (s *AIAgentTemplateConfigAvatarChat3D) SetUseVoiceprint(v bool) *AIAgentTemplateConfigAvatarChat3D {
	s.UseVoiceprint = &v
	return s
}

func (s *AIAgentTemplateConfigAvatarChat3D) SetUserOfflineTimeout(v int32) *AIAgentTemplateConfigAvatarChat3D {
	s.UserOfflineTimeout = &v
	return s
}

func (s *AIAgentTemplateConfigAvatarChat3D) SetUserOnlineTimeout(v int32) *AIAgentTemplateConfigAvatarChat3D {
	s.UserOnlineTimeout = &v
	return s
}

func (s *AIAgentTemplateConfigAvatarChat3D) SetVadLevel(v int32) *AIAgentTemplateConfigAvatarChat3D {
	s.VadLevel = &v
	return s
}

func (s *AIAgentTemplateConfigAvatarChat3D) SetVoiceId(v string) *AIAgentTemplateConfigAvatarChat3D {
	s.VoiceId = &v
	return s
}

func (s *AIAgentTemplateConfigAvatarChat3D) SetVoiceIdList(v []*string) *AIAgentTemplateConfigAvatarChat3D {
	s.VoiceIdList = v
	return s
}

func (s *AIAgentTemplateConfigAvatarChat3D) SetVoiceprintId(v string) *AIAgentTemplateConfigAvatarChat3D {
	s.VoiceprintId = &v
	return s
}

func (s *AIAgentTemplateConfigAvatarChat3D) SetVolume(v int64) *AIAgentTemplateConfigAvatarChat3D {
	s.Volume = &v
	return s
}

func (s *AIAgentTemplateConfigAvatarChat3D) SetWakeUpQuery(v string) *AIAgentTemplateConfigAvatarChat3D {
	s.WakeUpQuery = &v
	return s
}

func (s *AIAgentTemplateConfigAvatarChat3D) SetWorkflowOverrideParams(v string) *AIAgentTemplateConfigAvatarChat3D {
	s.WorkflowOverrideParams = &v
	return s
}

func (s *AIAgentTemplateConfigAvatarChat3D) Validate() error {
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

type AIAgentTemplateConfigAvatarChat3DLlmHistory struct {
	// The text content of the message.
	//
	// example:
	//
	// 你好
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The role of the conversation participant. Valid values:
	//
	// - `user`: The user.
	//
	// - `assistant`: The AI assistant.
	//
	// - `system`: The system.
	//
	// - `function`: A function call.
	//
	// - `plugin`: A plugin.
	//
	// - `tool`: A tool.
	//
	// example:
	//
	// user
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s AIAgentTemplateConfigAvatarChat3DLlmHistory) String() string {
	return dara.Prettify(s)
}

func (s AIAgentTemplateConfigAvatarChat3DLlmHistory) GoString() string {
	return s.String()
}

func (s *AIAgentTemplateConfigAvatarChat3DLlmHistory) GetContent() *string {
	return s.Content
}

func (s *AIAgentTemplateConfigAvatarChat3DLlmHistory) GetRole() *string {
	return s.Role
}

func (s *AIAgentTemplateConfigAvatarChat3DLlmHistory) SetContent(v string) *AIAgentTemplateConfigAvatarChat3DLlmHistory {
	s.Content = &v
	return s
}

func (s *AIAgentTemplateConfigAvatarChat3DLlmHistory) SetRole(v string) *AIAgentTemplateConfigAvatarChat3DLlmHistory {
	s.Role = &v
	return s
}

func (s *AIAgentTemplateConfigAvatarChat3DLlmHistory) Validate() error {
	return dara.Validate(s)
}

type AIAgentTemplateConfigVisionChat struct {
	// A list of hot words to improve ASR accuracy. A maximum of 128 words is supported.
	AsrHotWords []*string `json:"AsrHotWords,omitempty" xml:"AsrHotWords,omitempty" type:"Repeated"`
	// The language ID for Automatic Speech Recognition (ASR). Possible values:
	//
	// - `zh_mandarin`: Chinese
	//
	// - `en`: English
	//
	// - `zh_en`: Chinese-English
	//
	// - `es`: Spanish
	//
	// - `jp`: Japanese
	//
	// example:
	//
	// zh_mandarin
	AsrLanguageId *string `json:"AsrLanguageId,omitempty" xml:"AsrLanguageId,omitempty"`
	// The maximum duration of silence in milliseconds before a sentence break is detected. Range: 200 to 1,200. Default: 400.
	//
	// example:
	//
	// 400
	AsrMaxSilence *int32 `json:"AsrMaxSilence,omitempty" xml:"AsrMaxSilence,omitempty"`
	// Parameters for Alibaba Cloud Bailian. For details, see [Bailian App Params](https://help.aliyun.com/document_detail/2858132.html).
	//
	// example:
	//
	// {}
	BailianAppParams *string `json:"BailianAppParams,omitempty" xml:"BailianAppParams,omitempty"`
	CharBreak        *bool   `json:"CharBreak,omitempty" xml:"CharBreak,omitempty"`
	// Specifies whether to enable intelligent segmentation. If enabled, this feature intelligently merges pauses in a user\\"s speech into a single, complete sentence. Default: `true`.
	//
	// example:
	//
	// true
	EnableIntelligentSegment *bool `json:"EnableIntelligentSegment,omitempty" xml:"EnableIntelligentSegment,omitempty"`
	// Specifies whether to enable Push-to-Talk mode. Default: `false`.
	//
	// example:
	//
	// false
	EnablePushToTalk *bool `json:"EnablePushToTalk,omitempty" xml:"EnablePushToTalk,omitempty"`
	// Specifies whether to enable voice interruption. Default: `true`.
	//
	// example:
	//
	// true
	EnableVoiceInterrupt *bool `json:"EnableVoiceInterrupt,omitempty" xml:"EnableVoiceInterrupt,omitempty"`
	// Specifies whether to enable graceful shutdown. Default: `false`.
	//
	// - If enabled, the agent finishes its current speech (up to 10 seconds) before stopping.
	//
	// example:
	//
	// false
	GracefulShutdown *bool `json:"GracefulShutdown,omitempty" xml:"GracefulShutdown,omitempty"`
	// The greeting message delivered when a user joins the session. If this parameter is omitted, the greeting configured in the agent template is used. Maximum length: 128 characters.
	//
	// example:
	//
	// 早上好，我的朋友！
	Greeting *string `json:"Greeting,omitempty" xml:"Greeting,omitempty"`
	// A list of specific words or phrases that trigger a conversation interruption.
	InterruptWords []*string `json:"InterruptWords,omitempty" xml:"InterruptWords,omitempty" type:"Repeated"`
	// The LLM/MLLM conversation history.
	LlmHistory []*AIAgentTemplateConfigVisionChatLlmHistory `json:"LlmHistory,omitempty" xml:"LlmHistory,omitempty" type:"Repeated"`
	// The maximum number of conversation turns to retain in the LLM/MLLM history. Default: 10.
	//
	// example:
	//
	// 10
	LlmHistoryLimit *int32 `json:"LlmHistoryLimit,omitempty" xml:"LlmHistoryLimit,omitempty"`
	// The system prompt for the LLM, applied when the call starts.
	//
	// example:
	//
	// 你是一位友好且乐于助人的助手，专注于为用户提供准确的信息和建议。
	LlmSystemPrompt *string `json:"LlmSystemPrompt,omitempty" xml:"LlmSystemPrompt,omitempty"`
	// The maximum idle time in seconds with no interaction before the agent goes offline. Default: 600.
	//
	// example:
	//
	// 600
	MaxIdleTime *int32 `json:"MaxIdleTime,omitempty" xml:"MaxIdleTime,omitempty"`
	// Specifies whether to use voiceprint recognition. Default: `false`.
	//
	// example:
	//
	// false
	UseVoiceprint *bool `json:"UseVoiceprint,omitempty" xml:"UseVoiceprint,omitempty"`
	// The time in seconds that the agent waits after a user leaves before closing the task. Default: 5.
	//
	// example:
	//
	// 5
	UserOfflineTimeout *int32 `json:"UserOfflineTimeout,omitempty" xml:"UserOfflineTimeout,omitempty"`
	// The time in seconds that the agent waits for a user to join before closing the task. Default: 60.
	//
	// example:
	//
	// 60
	UserOnlineTimeout *int32 `json:"UserOnlineTimeout,omitempty" xml:"UserOnlineTimeout,omitempty"`
	// The interruption sensitivity threshold. A higher value makes it more difficult to interrupt the agent. Range: 0 to 11. Default: 11.
	//
	// - `0`: Disables VAD.
	//
	// - `1` to `10`: A higher value makes it more difficult to interrupt the agent.
	//
	// - `11`: Offers lower audio distortion and stronger resistance to interference.
	//
	// example:
	//
	// 0
	VadLevel *int32 `json:"VadLevel,omitempty" xml:"VadLevel,omitempty"`
	// The ID of the Text-to-Speech (TTS) voice. Changes take effect on the next utterance. If omitted, the default voice from the agent template is used. This parameter applies only to preset TTS voices. Maximum length: 64 characters. For available values, see [Intelligent voice effect samples](https://help.aliyun.com/document_detail/449563.html).
	//
	// example:
	//
	// zhixiaoxia
	VoiceId *string `json:"VoiceId,omitempty" xml:"VoiceId,omitempty"`
	// A list of available voices.
	VoiceIdList []*string `json:"VoiceIdList,omitempty" xml:"VoiceIdList,omitempty" type:"Repeated"`
	// The unique ID for voiceprint recognition. Default: not specified.
	//
	// example:
	//
	// uniqueId
	VoiceprintId *string `json:"VoiceprintId,omitempty" xml:"VoiceprintId,omitempty"`
	// The speaking volume of the agent.
	//
	// - If omitted, the system uses adaptive volume mode.
	//
	// - If specified, the valid range is 0 to 400. The output volume is calculated as: `Output Volume in Workflow` \\	- (`volume`/100). For example:
	//
	// 1. If `volume` is `0`, the output is silent.
	//
	// 2. If `volume` is `100`, the output volume is the original volume.
	//
	// 3. If `volume` is `200`, the output volume is twice the original volume.
	//
	// example:
	//
	// 100
	Volume *int64 `json:"Volume,omitempty" xml:"Volume,omitempty"`
	// An initial user query that the agent addresses immediately when the call starts.
	//
	// example:
	//
	// 今天天气怎么样？
	WakeUpQuery *string `json:"WakeUpQuery,omitempty" xml:"WakeUpQuery,omitempty"`
	// Workflow override parameters. Default: empty.
	//
	// example:
	//
	// {}
	WorkflowOverrideParams *string `json:"WorkflowOverrideParams,omitempty" xml:"WorkflowOverrideParams,omitempty"`
}

func (s AIAgentTemplateConfigVisionChat) String() string {
	return dara.Prettify(s)
}

func (s AIAgentTemplateConfigVisionChat) GoString() string {
	return s.String()
}

func (s *AIAgentTemplateConfigVisionChat) GetAsrHotWords() []*string {
	return s.AsrHotWords
}

func (s *AIAgentTemplateConfigVisionChat) GetAsrLanguageId() *string {
	return s.AsrLanguageId
}

func (s *AIAgentTemplateConfigVisionChat) GetAsrMaxSilence() *int32 {
	return s.AsrMaxSilence
}

func (s *AIAgentTemplateConfigVisionChat) GetBailianAppParams() *string {
	return s.BailianAppParams
}

func (s *AIAgentTemplateConfigVisionChat) GetCharBreak() *bool {
	return s.CharBreak
}

func (s *AIAgentTemplateConfigVisionChat) GetEnableIntelligentSegment() *bool {
	return s.EnableIntelligentSegment
}

func (s *AIAgentTemplateConfigVisionChat) GetEnablePushToTalk() *bool {
	return s.EnablePushToTalk
}

func (s *AIAgentTemplateConfigVisionChat) GetEnableVoiceInterrupt() *bool {
	return s.EnableVoiceInterrupt
}

func (s *AIAgentTemplateConfigVisionChat) GetGracefulShutdown() *bool {
	return s.GracefulShutdown
}

func (s *AIAgentTemplateConfigVisionChat) GetGreeting() *string {
	return s.Greeting
}

func (s *AIAgentTemplateConfigVisionChat) GetInterruptWords() []*string {
	return s.InterruptWords
}

func (s *AIAgentTemplateConfigVisionChat) GetLlmHistory() []*AIAgentTemplateConfigVisionChatLlmHistory {
	return s.LlmHistory
}

func (s *AIAgentTemplateConfigVisionChat) GetLlmHistoryLimit() *int32 {
	return s.LlmHistoryLimit
}

func (s *AIAgentTemplateConfigVisionChat) GetLlmSystemPrompt() *string {
	return s.LlmSystemPrompt
}

func (s *AIAgentTemplateConfigVisionChat) GetMaxIdleTime() *int32 {
	return s.MaxIdleTime
}

func (s *AIAgentTemplateConfigVisionChat) GetUseVoiceprint() *bool {
	return s.UseVoiceprint
}

func (s *AIAgentTemplateConfigVisionChat) GetUserOfflineTimeout() *int32 {
	return s.UserOfflineTimeout
}

func (s *AIAgentTemplateConfigVisionChat) GetUserOnlineTimeout() *int32 {
	return s.UserOnlineTimeout
}

func (s *AIAgentTemplateConfigVisionChat) GetVadLevel() *int32 {
	return s.VadLevel
}

func (s *AIAgentTemplateConfigVisionChat) GetVoiceId() *string {
	return s.VoiceId
}

func (s *AIAgentTemplateConfigVisionChat) GetVoiceIdList() []*string {
	return s.VoiceIdList
}

func (s *AIAgentTemplateConfigVisionChat) GetVoiceprintId() *string {
	return s.VoiceprintId
}

func (s *AIAgentTemplateConfigVisionChat) GetVolume() *int64 {
	return s.Volume
}

func (s *AIAgentTemplateConfigVisionChat) GetWakeUpQuery() *string {
	return s.WakeUpQuery
}

func (s *AIAgentTemplateConfigVisionChat) GetWorkflowOverrideParams() *string {
	return s.WorkflowOverrideParams
}

func (s *AIAgentTemplateConfigVisionChat) SetAsrHotWords(v []*string) *AIAgentTemplateConfigVisionChat {
	s.AsrHotWords = v
	return s
}

func (s *AIAgentTemplateConfigVisionChat) SetAsrLanguageId(v string) *AIAgentTemplateConfigVisionChat {
	s.AsrLanguageId = &v
	return s
}

func (s *AIAgentTemplateConfigVisionChat) SetAsrMaxSilence(v int32) *AIAgentTemplateConfigVisionChat {
	s.AsrMaxSilence = &v
	return s
}

func (s *AIAgentTemplateConfigVisionChat) SetBailianAppParams(v string) *AIAgentTemplateConfigVisionChat {
	s.BailianAppParams = &v
	return s
}

func (s *AIAgentTemplateConfigVisionChat) SetCharBreak(v bool) *AIAgentTemplateConfigVisionChat {
	s.CharBreak = &v
	return s
}

func (s *AIAgentTemplateConfigVisionChat) SetEnableIntelligentSegment(v bool) *AIAgentTemplateConfigVisionChat {
	s.EnableIntelligentSegment = &v
	return s
}

func (s *AIAgentTemplateConfigVisionChat) SetEnablePushToTalk(v bool) *AIAgentTemplateConfigVisionChat {
	s.EnablePushToTalk = &v
	return s
}

func (s *AIAgentTemplateConfigVisionChat) SetEnableVoiceInterrupt(v bool) *AIAgentTemplateConfigVisionChat {
	s.EnableVoiceInterrupt = &v
	return s
}

func (s *AIAgentTemplateConfigVisionChat) SetGracefulShutdown(v bool) *AIAgentTemplateConfigVisionChat {
	s.GracefulShutdown = &v
	return s
}

func (s *AIAgentTemplateConfigVisionChat) SetGreeting(v string) *AIAgentTemplateConfigVisionChat {
	s.Greeting = &v
	return s
}

func (s *AIAgentTemplateConfigVisionChat) SetInterruptWords(v []*string) *AIAgentTemplateConfigVisionChat {
	s.InterruptWords = v
	return s
}

func (s *AIAgentTemplateConfigVisionChat) SetLlmHistory(v []*AIAgentTemplateConfigVisionChatLlmHistory) *AIAgentTemplateConfigVisionChat {
	s.LlmHistory = v
	return s
}

func (s *AIAgentTemplateConfigVisionChat) SetLlmHistoryLimit(v int32) *AIAgentTemplateConfigVisionChat {
	s.LlmHistoryLimit = &v
	return s
}

func (s *AIAgentTemplateConfigVisionChat) SetLlmSystemPrompt(v string) *AIAgentTemplateConfigVisionChat {
	s.LlmSystemPrompt = &v
	return s
}

func (s *AIAgentTemplateConfigVisionChat) SetMaxIdleTime(v int32) *AIAgentTemplateConfigVisionChat {
	s.MaxIdleTime = &v
	return s
}

func (s *AIAgentTemplateConfigVisionChat) SetUseVoiceprint(v bool) *AIAgentTemplateConfigVisionChat {
	s.UseVoiceprint = &v
	return s
}

func (s *AIAgentTemplateConfigVisionChat) SetUserOfflineTimeout(v int32) *AIAgentTemplateConfigVisionChat {
	s.UserOfflineTimeout = &v
	return s
}

func (s *AIAgentTemplateConfigVisionChat) SetUserOnlineTimeout(v int32) *AIAgentTemplateConfigVisionChat {
	s.UserOnlineTimeout = &v
	return s
}

func (s *AIAgentTemplateConfigVisionChat) SetVadLevel(v int32) *AIAgentTemplateConfigVisionChat {
	s.VadLevel = &v
	return s
}

func (s *AIAgentTemplateConfigVisionChat) SetVoiceId(v string) *AIAgentTemplateConfigVisionChat {
	s.VoiceId = &v
	return s
}

func (s *AIAgentTemplateConfigVisionChat) SetVoiceIdList(v []*string) *AIAgentTemplateConfigVisionChat {
	s.VoiceIdList = v
	return s
}

func (s *AIAgentTemplateConfigVisionChat) SetVoiceprintId(v string) *AIAgentTemplateConfigVisionChat {
	s.VoiceprintId = &v
	return s
}

func (s *AIAgentTemplateConfigVisionChat) SetVolume(v int64) *AIAgentTemplateConfigVisionChat {
	s.Volume = &v
	return s
}

func (s *AIAgentTemplateConfigVisionChat) SetWakeUpQuery(v string) *AIAgentTemplateConfigVisionChat {
	s.WakeUpQuery = &v
	return s
}

func (s *AIAgentTemplateConfigVisionChat) SetWorkflowOverrideParams(v string) *AIAgentTemplateConfigVisionChat {
	s.WorkflowOverrideParams = &v
	return s
}

func (s *AIAgentTemplateConfigVisionChat) Validate() error {
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

type AIAgentTemplateConfigVisionChatLlmHistory struct {
	// The text content of the message.
	//
	// example:
	//
	// 你好
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The role of the conversation participant. Valid values:
	//
	// - `user`: The user.
	//
	// - `assistant`: The AI assistant.
	//
	// - `system`: The system.
	//
	// - `function`: A function call.
	//
	// - `plugin`: A plugin.
	//
	// - `tool`: A tool.
	//
	// example:
	//
	// user
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s AIAgentTemplateConfigVisionChatLlmHistory) String() string {
	return dara.Prettify(s)
}

func (s AIAgentTemplateConfigVisionChatLlmHistory) GoString() string {
	return s.String()
}

func (s *AIAgentTemplateConfigVisionChatLlmHistory) GetContent() *string {
	return s.Content
}

func (s *AIAgentTemplateConfigVisionChatLlmHistory) GetRole() *string {
	return s.Role
}

func (s *AIAgentTemplateConfigVisionChatLlmHistory) SetContent(v string) *AIAgentTemplateConfigVisionChatLlmHistory {
	s.Content = &v
	return s
}

func (s *AIAgentTemplateConfigVisionChatLlmHistory) SetRole(v string) *AIAgentTemplateConfigVisionChatLlmHistory {
	s.Role = &v
	return s
}

func (s *AIAgentTemplateConfigVisionChatLlmHistory) Validate() error {
	return dara.Validate(s)
}

type AIAgentTemplateConfigVoiceChat struct {
	// A list of hot words to improve ASR accuracy. A maximum of 128 words is supported.
	AsrHotWords []*string `json:"AsrHotWords,omitempty" xml:"AsrHotWords,omitempty" type:"Repeated"`
	// The language ID for Automatic Speech Recognition (ASR).
	//
	// Possible values:
	//
	// - `zh_mandarin`: Chinese
	//
	// - `en`: English
	//
	// - `zh_en`: Chinese-English
	//
	// - `es`: Spanish
	//
	// - `jp`: Japanese
	//
	// example:
	//
	// zh_mandarin
	AsrLanguageId *string `json:"AsrLanguageId,omitempty" xml:"AsrLanguageId,omitempty"`
	// The maximum duration of silence in milliseconds before a sentence break is detected. Range: 200 to 1,200. Default: 400.
	//
	// example:
	//
	// 400
	AsrMaxSilence *int32 `json:"AsrMaxSilence,omitempty" xml:"AsrMaxSilence,omitempty"`
	// The URL of the agent\\"s avatar for voice chat. Default: none.
	//
	// example:
	//
	// http://example.com/a.jpg
	AvatarUrl *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	// The type of the agent\\"s avatar URL. Default: none.
	//
	// example:
	//
	// USER
	AvatarUrlType *string `json:"AvatarUrlType,omitempty" xml:"AvatarUrlType,omitempty"`
	// Parameters for Alibaba Cloud Bailian. For details, see [Bailian App Params](https://help.aliyun.com/document_detail/2858132.html).
	//
	// example:
	//
	// {}
	BailianAppParams *string `json:"BailianAppParams,omitempty" xml:"BailianAppParams,omitempty"`
	CharBreak        *bool   `json:"CharBreak,omitempty" xml:"CharBreak,omitempty"`
	// Specifies whether to enable intelligent segmentation. If enabled, this feature intelligently merges pauses in a user\\"s speech into a single, complete sentence. Default: `true`.
	//
	// example:
	//
	// true
	EnableIntelligentSegment *bool `json:"EnableIntelligentSegment,omitempty" xml:"EnableIntelligentSegment,omitempty"`
	// Specifies whether to enable Push-to-Talk mode. Default: `false`.
	//
	// example:
	//
	// false
	EnablePushToTalk *bool `json:"EnablePushToTalk,omitempty" xml:"EnablePushToTalk,omitempty"`
	// Specifies whether to enable voice interruption. Default: `true`.
	//
	// example:
	//
	// true
	EnableVoiceInterrupt *bool `json:"EnableVoiceInterrupt,omitempty" xml:"EnableVoiceInterrupt,omitempty"`
	// Specifies whether to enable graceful shutdown. Default: `false`.
	//
	// - If enabled, the agent finishes its current speech (up to 10 seconds) before stopping.
	//
	// example:
	//
	// false
	GracefulShutdown *bool `json:"GracefulShutdown,omitempty" xml:"GracefulShutdown,omitempty"`
	// The greeting message delivered when a user joins the session. If this parameter is omitted, the greeting configured in the agent template is used. Maximum length: 128 characters.
	//
	// example:
	//
	// 早上好，我的朋友
	Greeting *string `json:"Greeting,omitempty" xml:"Greeting,omitempty"`
	// A list of specific words or phrases that trigger a conversation interruption.
	InterruptWords []*string `json:"InterruptWords,omitempty" xml:"InterruptWords,omitempty" type:"Repeated"`
	// The LLM/MLLM conversation history.
	LlmHistory []*AIAgentTemplateConfigVoiceChatLlmHistory `json:"LlmHistory,omitempty" xml:"LlmHistory,omitempty" type:"Repeated"`
	// The maximum number of conversation turns to retain in the LLM/MLLM history. Default: 10.
	//
	// example:
	//
	// 10
	LlmHistoryLimit *int32 `json:"LlmHistoryLimit,omitempty" xml:"LlmHistoryLimit,omitempty"`
	// The system prompt for the LLM, applied when the call starts.
	//
	// example:
	//
	// 你是一位友好且乐于助人的助手，专注于为用户提供准确的信息和建议。
	LlmSystemPrompt *string `json:"LlmSystemPrompt,omitempty" xml:"LlmSystemPrompt,omitempty"`
	// The maximum idle time in seconds with no interaction before the agent goes offline. Default: 600.
	//
	// example:
	//
	// 600
	MaxIdleTime *int32 `json:"MaxIdleTime,omitempty" xml:"MaxIdleTime,omitempty"`
	// Specifies whether to use voiceprint recognition. Default: `false`.
	//
	// example:
	//
	// false
	UseVoiceprint *bool `json:"UseVoiceprint,omitempty" xml:"UseVoiceprint,omitempty"`
	// The time in seconds that the agent waits after a user leaves before closing the task. Default: 5.
	//
	// example:
	//
	// 5
	UserOfflineTimeout *int32 `json:"UserOfflineTimeout,omitempty" xml:"UserOfflineTimeout,omitempty"`
	// The time in seconds that the agent waits for a user to join before closing the task. Default: 60.
	//
	// example:
	//
	// 60
	UserOnlineTimeout *int32 `json:"UserOnlineTimeout,omitempty" xml:"UserOnlineTimeout,omitempty"`
	// The interruption sensitivity threshold. A higher value makes it more difficult to interrupt the agent. Range: 0 to 11. Default: 11.
	//
	// - `0`: Disables Voice Activity Detection (VAD).
	//
	// - `1` to `10`: A higher value makes it more difficult to interrupt the agent.
	//
	// - `11`: Offers lower audio distortion and stronger resistance to interference.
	//
	// example:
	//
	// 11
	VadLevel *int32 `json:"VadLevel,omitempty" xml:"VadLevel,omitempty"`
	// The ID of the Text-to-Speech (TTS) voice. Changes take effect on the next utterance. If omitted, the default voice from the agent template is used. This parameter applies only to preset TTS voices. Maximum length: 64 characters. For available values, see [Intelligent voice effect samples](https://help.aliyun.com/document_detail/449563.html).
	//
	// example:
	//
	// zhixiaoxia
	VoiceId *string `json:"VoiceId,omitempty" xml:"VoiceId,omitempty"`
	// A list of available voices.
	VoiceIdList []*string `json:"VoiceIdList,omitempty" xml:"VoiceIdList,omitempty" type:"Repeated"`
	// The unique ID for voiceprint recognition. Default: not specified.
	//
	// example:
	//
	// uniqueId
	VoiceprintId *string `json:"VoiceprintId,omitempty" xml:"VoiceprintId,omitempty"`
	// The speaking volume of the agent.
	//
	// - If omitted, the system uses adaptive volume mode.
	//
	// - If specified, the valid range is 0 to 400. The output volume is calculated as: `Output Volume in Workflow` \\	- (`volume`/100). For example:
	//
	// 1. If `volume` is `0`, the output is silent.
	//
	// 2. If `volume` is `100`, the output volume is the original volume.
	//
	// 3. If `volume` is `200`, the output volume is twice the original volume.
	//
	// example:
	//
	// 100
	Volume *int64 `json:"Volume,omitempty" xml:"Volume,omitempty"`
	// An initial user query that the agent addresses immediately when the call starts.
	//
	// example:
	//
	// 今天天气怎么样？
	WakeUpQuery *string `json:"WakeUpQuery,omitempty" xml:"WakeUpQuery,omitempty"`
	// Workflow override parameters. Default: empty.
	//
	// example:
	//
	// {}
	WorkflowOverrideParams *string `json:"WorkflowOverrideParams,omitempty" xml:"WorkflowOverrideParams,omitempty"`
}

func (s AIAgentTemplateConfigVoiceChat) String() string {
	return dara.Prettify(s)
}

func (s AIAgentTemplateConfigVoiceChat) GoString() string {
	return s.String()
}

func (s *AIAgentTemplateConfigVoiceChat) GetAsrHotWords() []*string {
	return s.AsrHotWords
}

func (s *AIAgentTemplateConfigVoiceChat) GetAsrLanguageId() *string {
	return s.AsrLanguageId
}

func (s *AIAgentTemplateConfigVoiceChat) GetAsrMaxSilence() *int32 {
	return s.AsrMaxSilence
}

func (s *AIAgentTemplateConfigVoiceChat) GetAvatarUrl() *string {
	return s.AvatarUrl
}

func (s *AIAgentTemplateConfigVoiceChat) GetAvatarUrlType() *string {
	return s.AvatarUrlType
}

func (s *AIAgentTemplateConfigVoiceChat) GetBailianAppParams() *string {
	return s.BailianAppParams
}

func (s *AIAgentTemplateConfigVoiceChat) GetCharBreak() *bool {
	return s.CharBreak
}

func (s *AIAgentTemplateConfigVoiceChat) GetEnableIntelligentSegment() *bool {
	return s.EnableIntelligentSegment
}

func (s *AIAgentTemplateConfigVoiceChat) GetEnablePushToTalk() *bool {
	return s.EnablePushToTalk
}

func (s *AIAgentTemplateConfigVoiceChat) GetEnableVoiceInterrupt() *bool {
	return s.EnableVoiceInterrupt
}

func (s *AIAgentTemplateConfigVoiceChat) GetGracefulShutdown() *bool {
	return s.GracefulShutdown
}

func (s *AIAgentTemplateConfigVoiceChat) GetGreeting() *string {
	return s.Greeting
}

func (s *AIAgentTemplateConfigVoiceChat) GetInterruptWords() []*string {
	return s.InterruptWords
}

func (s *AIAgentTemplateConfigVoiceChat) GetLlmHistory() []*AIAgentTemplateConfigVoiceChatLlmHistory {
	return s.LlmHistory
}

func (s *AIAgentTemplateConfigVoiceChat) GetLlmHistoryLimit() *int32 {
	return s.LlmHistoryLimit
}

func (s *AIAgentTemplateConfigVoiceChat) GetLlmSystemPrompt() *string {
	return s.LlmSystemPrompt
}

func (s *AIAgentTemplateConfigVoiceChat) GetMaxIdleTime() *int32 {
	return s.MaxIdleTime
}

func (s *AIAgentTemplateConfigVoiceChat) GetUseVoiceprint() *bool {
	return s.UseVoiceprint
}

func (s *AIAgentTemplateConfigVoiceChat) GetUserOfflineTimeout() *int32 {
	return s.UserOfflineTimeout
}

func (s *AIAgentTemplateConfigVoiceChat) GetUserOnlineTimeout() *int32 {
	return s.UserOnlineTimeout
}

func (s *AIAgentTemplateConfigVoiceChat) GetVadLevel() *int32 {
	return s.VadLevel
}

func (s *AIAgentTemplateConfigVoiceChat) GetVoiceId() *string {
	return s.VoiceId
}

func (s *AIAgentTemplateConfigVoiceChat) GetVoiceIdList() []*string {
	return s.VoiceIdList
}

func (s *AIAgentTemplateConfigVoiceChat) GetVoiceprintId() *string {
	return s.VoiceprintId
}

func (s *AIAgentTemplateConfigVoiceChat) GetVolume() *int64 {
	return s.Volume
}

func (s *AIAgentTemplateConfigVoiceChat) GetWakeUpQuery() *string {
	return s.WakeUpQuery
}

func (s *AIAgentTemplateConfigVoiceChat) GetWorkflowOverrideParams() *string {
	return s.WorkflowOverrideParams
}

func (s *AIAgentTemplateConfigVoiceChat) SetAsrHotWords(v []*string) *AIAgentTemplateConfigVoiceChat {
	s.AsrHotWords = v
	return s
}

func (s *AIAgentTemplateConfigVoiceChat) SetAsrLanguageId(v string) *AIAgentTemplateConfigVoiceChat {
	s.AsrLanguageId = &v
	return s
}

func (s *AIAgentTemplateConfigVoiceChat) SetAsrMaxSilence(v int32) *AIAgentTemplateConfigVoiceChat {
	s.AsrMaxSilence = &v
	return s
}

func (s *AIAgentTemplateConfigVoiceChat) SetAvatarUrl(v string) *AIAgentTemplateConfigVoiceChat {
	s.AvatarUrl = &v
	return s
}

func (s *AIAgentTemplateConfigVoiceChat) SetAvatarUrlType(v string) *AIAgentTemplateConfigVoiceChat {
	s.AvatarUrlType = &v
	return s
}

func (s *AIAgentTemplateConfigVoiceChat) SetBailianAppParams(v string) *AIAgentTemplateConfigVoiceChat {
	s.BailianAppParams = &v
	return s
}

func (s *AIAgentTemplateConfigVoiceChat) SetCharBreak(v bool) *AIAgentTemplateConfigVoiceChat {
	s.CharBreak = &v
	return s
}

func (s *AIAgentTemplateConfigVoiceChat) SetEnableIntelligentSegment(v bool) *AIAgentTemplateConfigVoiceChat {
	s.EnableIntelligentSegment = &v
	return s
}

func (s *AIAgentTemplateConfigVoiceChat) SetEnablePushToTalk(v bool) *AIAgentTemplateConfigVoiceChat {
	s.EnablePushToTalk = &v
	return s
}

func (s *AIAgentTemplateConfigVoiceChat) SetEnableVoiceInterrupt(v bool) *AIAgentTemplateConfigVoiceChat {
	s.EnableVoiceInterrupt = &v
	return s
}

func (s *AIAgentTemplateConfigVoiceChat) SetGracefulShutdown(v bool) *AIAgentTemplateConfigVoiceChat {
	s.GracefulShutdown = &v
	return s
}

func (s *AIAgentTemplateConfigVoiceChat) SetGreeting(v string) *AIAgentTemplateConfigVoiceChat {
	s.Greeting = &v
	return s
}

func (s *AIAgentTemplateConfigVoiceChat) SetInterruptWords(v []*string) *AIAgentTemplateConfigVoiceChat {
	s.InterruptWords = v
	return s
}

func (s *AIAgentTemplateConfigVoiceChat) SetLlmHistory(v []*AIAgentTemplateConfigVoiceChatLlmHistory) *AIAgentTemplateConfigVoiceChat {
	s.LlmHistory = v
	return s
}

func (s *AIAgentTemplateConfigVoiceChat) SetLlmHistoryLimit(v int32) *AIAgentTemplateConfigVoiceChat {
	s.LlmHistoryLimit = &v
	return s
}

func (s *AIAgentTemplateConfigVoiceChat) SetLlmSystemPrompt(v string) *AIAgentTemplateConfigVoiceChat {
	s.LlmSystemPrompt = &v
	return s
}

func (s *AIAgentTemplateConfigVoiceChat) SetMaxIdleTime(v int32) *AIAgentTemplateConfigVoiceChat {
	s.MaxIdleTime = &v
	return s
}

func (s *AIAgentTemplateConfigVoiceChat) SetUseVoiceprint(v bool) *AIAgentTemplateConfigVoiceChat {
	s.UseVoiceprint = &v
	return s
}

func (s *AIAgentTemplateConfigVoiceChat) SetUserOfflineTimeout(v int32) *AIAgentTemplateConfigVoiceChat {
	s.UserOfflineTimeout = &v
	return s
}

func (s *AIAgentTemplateConfigVoiceChat) SetUserOnlineTimeout(v int32) *AIAgentTemplateConfigVoiceChat {
	s.UserOnlineTimeout = &v
	return s
}

func (s *AIAgentTemplateConfigVoiceChat) SetVadLevel(v int32) *AIAgentTemplateConfigVoiceChat {
	s.VadLevel = &v
	return s
}

func (s *AIAgentTemplateConfigVoiceChat) SetVoiceId(v string) *AIAgentTemplateConfigVoiceChat {
	s.VoiceId = &v
	return s
}

func (s *AIAgentTemplateConfigVoiceChat) SetVoiceIdList(v []*string) *AIAgentTemplateConfigVoiceChat {
	s.VoiceIdList = v
	return s
}

func (s *AIAgentTemplateConfigVoiceChat) SetVoiceprintId(v string) *AIAgentTemplateConfigVoiceChat {
	s.VoiceprintId = &v
	return s
}

func (s *AIAgentTemplateConfigVoiceChat) SetVolume(v int64) *AIAgentTemplateConfigVoiceChat {
	s.Volume = &v
	return s
}

func (s *AIAgentTemplateConfigVoiceChat) SetWakeUpQuery(v string) *AIAgentTemplateConfigVoiceChat {
	s.WakeUpQuery = &v
	return s
}

func (s *AIAgentTemplateConfigVoiceChat) SetWorkflowOverrideParams(v string) *AIAgentTemplateConfigVoiceChat {
	s.WorkflowOverrideParams = &v
	return s
}

func (s *AIAgentTemplateConfigVoiceChat) Validate() error {
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

type AIAgentTemplateConfigVoiceChatLlmHistory struct {
	// The text content of the message.
	//
	// example:
	//
	// 你好
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The role of the conversation participant. Valid values:
	//
	// - `user`: The user.
	//
	// - `assistant`: The AI assistant.
	//
	// - `system`: The system.
	//
	// - `function`: A function call.
	//
	// - `plugin`: A plugin.
	//
	// - `tool`: A tool.
	//
	// example:
	//
	// user
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s AIAgentTemplateConfigVoiceChatLlmHistory) String() string {
	return dara.Prettify(s)
}

func (s AIAgentTemplateConfigVoiceChatLlmHistory) GoString() string {
	return s.String()
}

func (s *AIAgentTemplateConfigVoiceChatLlmHistory) GetContent() *string {
	return s.Content
}

func (s *AIAgentTemplateConfigVoiceChatLlmHistory) GetRole() *string {
	return s.Role
}

func (s *AIAgentTemplateConfigVoiceChatLlmHistory) SetContent(v string) *AIAgentTemplateConfigVoiceChatLlmHistory {
	s.Content = &v
	return s
}

func (s *AIAgentTemplateConfigVoiceChatLlmHistory) SetRole(v string) *AIAgentTemplateConfigVoiceChatLlmHistory {
	s.Role = &v
	return s
}

func (s *AIAgentTemplateConfigVoiceChatLlmHistory) Validate() error {
	return dara.Validate(s)
}
