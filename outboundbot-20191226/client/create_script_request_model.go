// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScriptRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v int64) *CreateScriptRequest
	GetAgentId() *int64
	SetAgentKey(v string) *CreateScriptRequest
	GetAgentKey() *string
	SetAgentLlm(v bool) *CreateScriptRequest
	GetAgentLlm() *bool
	SetAsrConfig(v string) *CreateScriptRequest
	GetAsrConfig() *string
	SetChatbotId(v string) *CreateScriptRequest
	GetChatbotId() *string
	SetEmotionEnable(v bool) *CreateScriptRequest
	GetEmotionEnable() *bool
	SetIndustry(v string) *CreateScriptRequest
	GetIndustry() *string
	SetInstanceId(v string) *CreateScriptRequest
	GetInstanceId() *string
	SetLongWaitEnable(v bool) *CreateScriptRequest
	GetLongWaitEnable() *bool
	SetMiniPlaybackEnable(v bool) *CreateScriptRequest
	GetMiniPlaybackEnable() *bool
	SetNewBargeInEnable(v bool) *CreateScriptRequest
	GetNewBargeInEnable() *bool
	SetNluAccessType(v string) *CreateScriptRequest
	GetNluAccessType() *string
	SetNluEngine(v string) *CreateScriptRequest
	GetNluEngine() *string
	SetScene(v string) *CreateScriptRequest
	GetScene() *string
	SetScriptContent(v []*string) *CreateScriptRequest
	GetScriptContent() []*string
	SetScriptDescription(v string) *CreateScriptRequest
	GetScriptDescription() *string
	SetScriptName(v string) *CreateScriptRequest
	GetScriptName() *string
	SetScriptNluProfileJsonString(v string) *CreateScriptRequest
	GetScriptNluProfileJsonString() *string
	SetScriptWaveform(v []*string) *CreateScriptRequest
	GetScriptWaveform() []*string
	SetTtsConfig(v string) *CreateScriptRequest
	GetTtsConfig() *string
}

type CreateScriptRequest struct {
	// Robot workspace ID
	//
	// example:
	//
	// 1198938
	AgentId *int64 `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// Access key for the robot workspace
	//
	// example:
	//
	// 9137ab9c27044921860030adf8590ec4_p_outbound_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// Indicates whether the robot workspace is an LLM workspace
	//
	// example:
	//
	// false
	AgentLlm *bool `json:"AgentLlm,omitempty" xml:"AgentLlm,omitempty"`
	// ASR configuration and parameter definitions
	//
	// - **appKey**: Alibaba Cloud account appKey.
	//
	// - **maxEndSilence**: Speech endpoint detection duration.
	//
	// - **silenceTimeout**: Silence timeout, in seconds. The system times out after the user remains silent for N seconds.
	//
	// - **engine**: Service to invoke; valid values: [`ali`, `xunfei`]
	//
	// - **nlsServiceType**: Service type to invoke; valid values: [`Managed`, `Authorized`]
	//
	// - **engineXunfei**: Configuration required when the engine is set to `xunfei`.
	//
	// > When `engine` is set to `ali` and `nlsServiceType` is `Authorized`, a custom service provided by Alibaba Cloud is used. When `engine` is `ali` and `nlsServiceType` is `Managed`, the default service is used. When `engine` is `xunfei` and `nlsServiceType` is `Authorized`, the service provider is Xunfei, and the corresponding Xunfei configuration must be provided, for example: `{"uuid":"ed2xxxxxxxxx","globalMaxEndSilence":700,"globalMaxEndSilenceEnable":true}`
	//
	// - **globalMaxEndSilence**: Silence detection duration, in milliseconds.
	//
	// - **globalMaxEndSilenceEnable**: Silence detection toggle, enabled by default.
	//
	// - **speechNoiseThreshold**: Noise filtering threshold
	//
	// example:
	//
	// {
	//
	// 	"appKey": "oQDVNlE6fZ5mg46X",
	//
	// 	"engine": "ali",
	//
	// 	"engineXunfei": "",
	//
	// 	"globalMaxEndSilence": 700,
	//
	// 	"globalMaxEndSilenceEnable": true,
	//
	// 	"maxEndSilence": "500",
	//
	// 	"nlsServiceType": "Managed",
	//
	// 	"silenceTimeout": "5000",
	//
	// 	"speechNoiseThreshold": "0"
	//
	// }
	AsrConfig *string `json:"AsrConfig,omitempty" xml:"AsrConfig,omitempty"`
	// When the NluServiceType of the instance is Authorized or Provided, you must use this field to specify the instance ID of the chatbot to which the script is attached.
	//
	// example:
	//
	// chatbot-cn-IfaUfqaUnb
	ChatbotId *string `json:"ChatbotId,omitempty" xml:"ChatbotId,omitempty"`
	// Emotion detection configuration Toggle (applicable only in non-LLM scenarios)
	//
	// example:
	//
	// true
	EmotionEnable *bool `json:"EmotionEnable,omitempty" xml:"EmotionEnable,omitempty"`
	// Associated industry
	//
	// This parameter is required.
	//
	// example:
	//
	// 教育
	Industry *string `json:"Industry,omitempty" xml:"Industry,omitempty"`
	// Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// c46001bc-3ead-4bfd-9a69-4b5b66a4a3f4
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Intelligent sentence segmentation Toggle (applicable only in non-LLM scenarios)
	//
	// example:
	//
	// true
	LongWaitEnable *bool `json:"LongWaitEnable,omitempty" xml:"LongWaitEnable,omitempty"`
	// Follow-up utterance configuration Toggle (applicable only in non-LLM scenarios)
	//
	// example:
	//
	// true
	MiniPlaybackEnable *bool `json:"MiniPlaybackEnable,omitempty" xml:"MiniPlaybackEnable,omitempty"`
	// Graceful interruption toggle (applicable to small models)
	//
	// example:
	//
	// true
	NewBargeInEnable *bool `json:"NewBargeInEnable,omitempty" xml:"NewBargeInEnable,omitempty"`
	// NLU access method (applicable only to LLM scenarios). Enumeration: `Managed`—uses the public Alibaba Cloud account for access. Leave empty for non-LLM scenarios.
	//
	// example:
	//
	// Managed
	NluAccessType *string `json:"NluAccessType,omitempty" xml:"NluAccessType,omitempty"`
	// NLU engine (applicable only in LLM scenarios; empty for non-LLM scenarios).
	//
	// - Prompts: LLM scenario.
	//
	// - SSE_FUNCTION: Function Compute pattern.
	//
	// - BeeBot: pipeline pattern.
	//
	// example:
	//
	// Prompts
	NluEngine *string `json:"NluEngine,omitempty" xml:"NluEngine,omitempty"`
	// Associated scenario
	//
	// This parameter is required.
	//
	// example:
	//
	// 回访
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// For notification-type instances, pass in a list of scripts.
	//
	// Deprecated.
	//
	// example:
	//
	// []
	ScriptContent []*string `json:"ScriptContent,omitempty" xml:"ScriptContent,omitempty" type:"Repeated"`
	// Script description
	//
	// example:
	//
	// 课程回复话术
	ScriptDescription *string `json:"ScriptDescription,omitempty" xml:"ScriptDescription,omitempty"`
	// Script Name
	//
	// This parameter is required.
	//
	// example:
	//
	// 课程满意度回访
	ScriptName *string `json:"ScriptName,omitempty" xml:"ScriptName,omitempty"`
	// > When nluEngine is set to SSE_FUNUNCTION, you must provide the corresponding configuration.
	//
	// Function Compute service mode configuration:
	//
	// - fcRegion: Function Compute region
	//
	// - fcFunction: Function name
	//
	// - fcHttpTriggerUrl: Function trigger URL
	//
	// example:
	//
	// {"fcRegion":"cn-shanghai","fcFunction":"waihu_test","fcHttpTriggerUrl":"https://waihu-test.xxxxxxxxxxx.cn-shanghai-vpc.fcapp.run"}
	ScriptNluProfileJsonString *string `json:"ScriptNluProfileJsonString,omitempty" xml:"ScriptNluProfileJsonString,omitempty"`
	// For notification-type instances, specify a list of script audio files.
	//
	// Deprecated.
	//
	// example:
	//
	// []
	ScriptWaveform []*string `json:"ScriptWaveform,omitempty" xml:"ScriptWaveform,omitempty" type:"Repeated"`
	// TTS configuration parameters:
	//
	// - **voice**: Voice speaker.
	//
	// - **volume**: Volume, range: 0–100. Default Value: 50.
	//
	// - **speechRate**: Speech rate, range: -500–500. Default Value: 0.
	//
	// - **pitchRate**: Pitch, range: -500–500. Default Value: 0.
	//
	// - **globalInterruptible**: Speech interruption configuration.
	//
	// - **engine**: Invoked service; options: [ali, volc, xunfei]. xunfei is not supported in LLM scenarios.
	//
	// - **nlsServiceType**: Service Type; options: [Managed, Authorized].
	//
	// - **engineXunfei**: Configuration when the service provider is xunfei.
	//
	// > 1. When engine is ali and nlsServiceType is Authorized, a Custom service is used.
	//
	// > 2. When engine is ali and nlsServiceType is Managed, the default service is used.
	//
	// > 3. When engine is xunfei (applicable only in non-LLM scenarios) and nlsServiceType is Authorized, the service provider is xunfei, and engineXunfei must be configured as {"pitchRate":50,"speechRate":50,"voice":"aisjiuxu","volume":50}.
	//
	// > 4. When engine is volc and nlsServiceType is Authorized, the service used is Doubao.
	//
	// example:
	//
	// {
	//
	// 	"appKey": "oQDVNlE6fZ5mg46X",
	//
	// 	"engine": "ali",
	//
	// 	"engineXunfei": "",
	//
	// 	"globalInterruptible": true,
	//
	// 	"nlsServiceType": "Managed",
	//
	// 	"pitchRate": "0",
	//
	// 	"speechRate": "0",
	//
	// 	"voice": "zhiyuan",
	//
	// 	"volume": "50"
	//
	// }
	TtsConfig *string `json:"TtsConfig,omitempty" xml:"TtsConfig,omitempty"`
}

func (s CreateScriptRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateScriptRequest) GoString() string {
	return s.String()
}

func (s *CreateScriptRequest) GetAgentId() *int64 {
	return s.AgentId
}

func (s *CreateScriptRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *CreateScriptRequest) GetAgentLlm() *bool {
	return s.AgentLlm
}

func (s *CreateScriptRequest) GetAsrConfig() *string {
	return s.AsrConfig
}

func (s *CreateScriptRequest) GetChatbotId() *string {
	return s.ChatbotId
}

func (s *CreateScriptRequest) GetEmotionEnable() *bool {
	return s.EmotionEnable
}

func (s *CreateScriptRequest) GetIndustry() *string {
	return s.Industry
}

func (s *CreateScriptRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateScriptRequest) GetLongWaitEnable() *bool {
	return s.LongWaitEnable
}

func (s *CreateScriptRequest) GetMiniPlaybackEnable() *bool {
	return s.MiniPlaybackEnable
}

func (s *CreateScriptRequest) GetNewBargeInEnable() *bool {
	return s.NewBargeInEnable
}

func (s *CreateScriptRequest) GetNluAccessType() *string {
	return s.NluAccessType
}

func (s *CreateScriptRequest) GetNluEngine() *string {
	return s.NluEngine
}

func (s *CreateScriptRequest) GetScene() *string {
	return s.Scene
}

func (s *CreateScriptRequest) GetScriptContent() []*string {
	return s.ScriptContent
}

func (s *CreateScriptRequest) GetScriptDescription() *string {
	return s.ScriptDescription
}

func (s *CreateScriptRequest) GetScriptName() *string {
	return s.ScriptName
}

func (s *CreateScriptRequest) GetScriptNluProfileJsonString() *string {
	return s.ScriptNluProfileJsonString
}

func (s *CreateScriptRequest) GetScriptWaveform() []*string {
	return s.ScriptWaveform
}

func (s *CreateScriptRequest) GetTtsConfig() *string {
	return s.TtsConfig
}

func (s *CreateScriptRequest) SetAgentId(v int64) *CreateScriptRequest {
	s.AgentId = &v
	return s
}

func (s *CreateScriptRequest) SetAgentKey(v string) *CreateScriptRequest {
	s.AgentKey = &v
	return s
}

func (s *CreateScriptRequest) SetAgentLlm(v bool) *CreateScriptRequest {
	s.AgentLlm = &v
	return s
}

func (s *CreateScriptRequest) SetAsrConfig(v string) *CreateScriptRequest {
	s.AsrConfig = &v
	return s
}

func (s *CreateScriptRequest) SetChatbotId(v string) *CreateScriptRequest {
	s.ChatbotId = &v
	return s
}

func (s *CreateScriptRequest) SetEmotionEnable(v bool) *CreateScriptRequest {
	s.EmotionEnable = &v
	return s
}

func (s *CreateScriptRequest) SetIndustry(v string) *CreateScriptRequest {
	s.Industry = &v
	return s
}

func (s *CreateScriptRequest) SetInstanceId(v string) *CreateScriptRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateScriptRequest) SetLongWaitEnable(v bool) *CreateScriptRequest {
	s.LongWaitEnable = &v
	return s
}

func (s *CreateScriptRequest) SetMiniPlaybackEnable(v bool) *CreateScriptRequest {
	s.MiniPlaybackEnable = &v
	return s
}

func (s *CreateScriptRequest) SetNewBargeInEnable(v bool) *CreateScriptRequest {
	s.NewBargeInEnable = &v
	return s
}

func (s *CreateScriptRequest) SetNluAccessType(v string) *CreateScriptRequest {
	s.NluAccessType = &v
	return s
}

func (s *CreateScriptRequest) SetNluEngine(v string) *CreateScriptRequest {
	s.NluEngine = &v
	return s
}

func (s *CreateScriptRequest) SetScene(v string) *CreateScriptRequest {
	s.Scene = &v
	return s
}

func (s *CreateScriptRequest) SetScriptContent(v []*string) *CreateScriptRequest {
	s.ScriptContent = v
	return s
}

func (s *CreateScriptRequest) SetScriptDescription(v string) *CreateScriptRequest {
	s.ScriptDescription = &v
	return s
}

func (s *CreateScriptRequest) SetScriptName(v string) *CreateScriptRequest {
	s.ScriptName = &v
	return s
}

func (s *CreateScriptRequest) SetScriptNluProfileJsonString(v string) *CreateScriptRequest {
	s.ScriptNluProfileJsonString = &v
	return s
}

func (s *CreateScriptRequest) SetScriptWaveform(v []*string) *CreateScriptRequest {
	s.ScriptWaveform = v
	return s
}

func (s *CreateScriptRequest) SetTtsConfig(v string) *CreateScriptRequest {
	s.TtsConfig = &v
	return s
}

func (s *CreateScriptRequest) Validate() error {
	return dara.Validate(s)
}
