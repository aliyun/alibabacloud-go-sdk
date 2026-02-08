// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyScriptRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v int64) *ModifyScriptRequest
	GetAgentId() *int64
	SetAgentKey(v string) *ModifyScriptRequest
	GetAgentKey() *string
	SetAgentLlm(v bool) *ModifyScriptRequest
	GetAgentLlm() *bool
	SetAsrConfig(v string) *ModifyScriptRequest
	GetAsrConfig() *string
	SetChatConfig(v string) *ModifyScriptRequest
	GetChatConfig() *string
	SetChatbotId(v string) *ModifyScriptRequest
	GetChatbotId() *string
	SetEmotionEnable(v bool) *ModifyScriptRequest
	GetEmotionEnable() *bool
	SetIndustry(v string) *ModifyScriptRequest
	GetIndustry() *string
	SetInstanceId(v string) *ModifyScriptRequest
	GetInstanceId() *string
	SetLabelConfig(v string) *ModifyScriptRequest
	GetLabelConfig() *string
	SetLongWaitEnable(v bool) *ModifyScriptRequest
	GetLongWaitEnable() *bool
	SetMiniPlaybackConfigListJsonString(v string) *ModifyScriptRequest
	GetMiniPlaybackConfigListJsonString() *string
	SetMiniPlaybackEnable(v bool) *ModifyScriptRequest
	GetMiniPlaybackEnable() *bool
	SetNewBargeInEnable(v bool) *ModifyScriptRequest
	GetNewBargeInEnable() *bool
	SetNlsConfig(v string) *ModifyScriptRequest
	GetNlsConfig() *string
	SetNluAccessType(v string) *ModifyScriptRequest
	GetNluAccessType() *string
	SetNluEngine(v string) *ModifyScriptRequest
	GetNluEngine() *string
	SetScene(v string) *ModifyScriptRequest
	GetScene() *string
	SetScriptContent(v []*string) *ModifyScriptRequest
	GetScriptContent() []*string
	SetScriptDescription(v string) *ModifyScriptRequest
	GetScriptDescription() *string
	SetScriptId(v string) *ModifyScriptRequest
	GetScriptId() *string
	SetScriptName(v string) *ModifyScriptRequest
	GetScriptName() *string
	SetScriptWaveform(v []*string) *ModifyScriptRequest
	GetScriptWaveform() []*string
	SetTtsConfig(v string) *ModifyScriptRequest
	GetTtsConfig() *string
}

type ModifyScriptRequest struct {
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
	// Indicates whether the chatbot workspace is an LLM workspace
	//
	// example:
	//
	// false
	AgentLlm *bool `json:"AgentLlm,omitempty" xml:"AgentLlm,omitempty"`
	// ASR configuration, parameter definitions
	//
	// - **appKey**: Alibaba Cloud account appKey.
	//
	// - **maxEndSilence**: Speech endpoint detection duration.
	//
	// - **silenceTimeout**: Silence timeout in seconds. The session times out after the user remains silent for N seconds.
	//
	// - **engine**: Service to invoke; options are [`ali`, `xunfei`].
	//
	// - **nlsServiceType**: Service type to invoke:
	//
	//    - **Managed**: Public cloud NLS service.
	//
	//    - **Authorized**: Authorized NLS service.
	//
	//    - **Provided**: NLS service provided by the customer using AccessKey/SecretKey.
	//
	//    - **Apes**: Private Cloud service.
	//
	// - **engineXunfei**: Configuration required when the engine is set to `xunfei`.
	//
	// > When `engine` is set to `ali` and `nlsServiceType` is `Authorized`, a custom service provided by `ali` is used. When `engine` is `ali` and `nlsServiceType` is `Managed`, the default service is used. When `engine` is `xunfei` and `nlsServiceType` is `Authorized`, the service provider is `xunfei`, and the following `xunfei` configuration must be provided: `{"uuid":"ed2xxxxxxxxx","globalMaxEndSilence":700,"globalMaxEndSilenceEnable":true}`.
	//
	// - **globalMaxEndSilence**: Silence detection threshold, in milliseconds.
	//
	// - **globalMaxEndSilenceEnable**: Toggle for silence detection. Enabled by default.
	//
	// - **speechNoiseThreshold**: Noise filtering threshold.
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
	// Call configuration
	//
	// - **silenceConfig**: Silence configuration
	//
	//     - **silenceReplyTimeout**: Silence reply timeout, in milliseconds.
	//
	//     - **silenceTimeoutMaxCount**: Number of consecutive silence timeouts after which the call is disconnected.
	//
	//     - **silenceTimeoutMaxCountEnable**: Toggle to disconnect the call after repeated silence timeouts.
	//
	// - **hangupConfig**: Hang-up configuration
	//
	//     - **aiHangupEnable**: AI hang-up. Values: `true` or `false`.
	//
	//     - **delayHangup**: Delayed hang-up, maximum value is 10.
	//
	//     - **hangupMaxRounds**: Maximum number of interaction rounds, up to 100.
	//
	//     - **hangupMaxRoundsBroadcast**: Broadcast message played before hanging up due to maximum rounds.
	//
	//     - **hangupMaxRoundsEnable**: Toggle to enforce the maximum interaction rounds limit (`true`/`false`).
	//
	// - **securityInterceptConfig**: Security interception configuration
	//
	//     - **broadcast**: Broadcast message played upon interception.
	//
	// - **specialInterceptConfig**: Special-case interception
	//
	//     - **specialInterceptEnable**: Toggle for special-case interception.
	//
	//     - **specialIntercepts**: Special cases:
	//
	//         - **voiceAssistant**: Voice assistant.
	//
	//         - **extensionNumberTransfer**: Extension number transfer.
	//
	// - **transitionConfig**: Transition phrase model configuration
	//
	//     - **transitionSwitch**: Toggle for the transition phrase model.
	//
	// example:
	//
	// {"silenceConfig":{"silenceReplyTimeout":499,"silenceTimeoutMaxCount":10,"silenceTimeoutMaxCountEnable":true},"hangupConfig":{"aiHangupEnable":false,"delayHangup":0,"hangupMaxRounds":20,"hangupMaxRoundsBroadcast":"感谢您的接听，祝您生活愉快，再见!","hangupMaxRoundsEnable":false},"securityInterceptConfig":{"broadcast":"您说的这个问题我不能回答您，您可以尝试询问其他问题。"},"specialInterceptConfig":{"specialInterceptEnable":false,"specialIntercepts":[{"code":"voiceAssistant"},{"code":"extensionNumberTransfer"}]},"transitionConfig":{"transitionSwitch":false}}
	ChatConfig *string `json:"ChatConfig,omitempty" xml:"ChatConfig,omitempty"`
	// Chatbot ID
	//
	// example:
	//
	// chatbot-cn-iFZfi7eq6e
	ChatbotId *string `json:"ChatbotId,omitempty" xml:"ChatbotId,omitempty"`
	// Emotion detection configuration toggle
	//
	// example:
	//
	// true
	EmotionEnable *bool `json:"EmotionEnable,omitempty" xml:"EmotionEnable,omitempty"`
	// Industry
	//
	// This parameter is required.
	//
	// example:
	//
	// 电商
	Industry *string `json:"Industry,omitempty" xml:"Industry,omitempty"`
	// Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// c6320d3c-fa45-4011-b3b1-acdfabe3a8c6
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	LabelConfig *string `json:"LabelConfig,omitempty" xml:"LabelConfig,omitempty"`
	// Intelligent sentence segmentation toggle (small model)
	//
	// example:
	//
	// true
	LongWaitEnable *bool `json:"LongWaitEnable,omitempty" xml:"LongWaitEnable,omitempty"`
	// Tone transition configuration. Customization is not currently supported, and this parameter does not need to be passed by default. (Deprecated)
	//
	// example:
	//
	// 无
	MiniPlaybackConfigListJsonString *string `json:"MiniPlaybackConfigListJsonString,omitempty" xml:"MiniPlaybackConfigListJsonString,omitempty"`
	// Follow-up utterance configuration toggle (small model)
	//
	// example:
	//
	// true
	MiniPlaybackEnable *bool `json:"MiniPlaybackEnable,omitempty" xml:"MiniPlaybackEnable,omitempty"`
	// Graceful interruption toggle (for small models)
	//
	// example:
	//
	// true
	NewBargeInEnable *bool `json:"NewBargeInEnable,omitempty" xml:"NewBargeInEnable,omitempty"`
	// Deprecated
	//
	// example:
	//
	// 无
	NlsConfig *string `json:"NlsConfig,omitempty" xml:"NlsConfig,omitempty"`
	// NLU access method (applicable only to LLM scenarios)
	//
	// Enumeration:
	//
	// - Managed: Access using Alibaba public account.
	//
	// - Empty for non-LLM scenarios.
	//
	// example:
	//
	// Managed
	NluAccessType *string `json:"NluAccessType,omitempty" xml:"NluAccessType,omitempty"`
	// NLU engine (applicable only to LLM scenarios)
	//
	// > The scenario pattern cannot be modified after the scenario is created.
	//
	// Enumeration:
	//
	// - Prompts: Text-filling pattern.
	//
	// - SSE_FUNCTION: Function Compute pattern.
	//
	// - Empty for non-LLM scenarios.
	//
	// example:
	//
	// Prompts
	NluEngine *string `json:"NluEngine,omitempty" xml:"NluEngine,omitempty"`
	// Scenario information
	//
	// This parameter is required.
	//
	// example:
	//
	// 回访
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// For notification-type instances, pass a list of scripts.
	//
	// Deprecated.
	ScriptContent []*string `json:"ScriptContent,omitempty" xml:"ScriptContent,omitempty" type:"Repeated"`
	// Script description
	//
	// example:
	//
	// 返工回访话术
	ScriptDescription *string `json:"ScriptDescription,omitempty" xml:"ScriptDescription,omitempty"`
	// Script ID
	//
	// This parameter is required.
	//
	// example:
	//
	// c153d0d8-ba04-41c0-8632-453944c9dd0b
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// Script name
	//
	// This parameter is required.
	//
	// example:
	//
	// 续费提醒
	ScriptName *string `json:"ScriptName,omitempty" xml:"ScriptName,omitempty"`
	// Deprecated. Previously used to pass a list of voice prompts for notification-type instances.
	ScriptWaveform []*string `json:"ScriptWaveform,omitempty" xml:"ScriptWaveform,omitempty" type:"Repeated"`
	// TTS configuration parameters
	//
	// - **voice**: Speaker.
	//
	// - **volume**: Volume, range: 0–100. Default value: 50.
	//
	// - **speechRate**: Speech rate, range: -500–500. Default value: 0.
	//
	// - **pitchRate**: Pitch, range: -500–500. Default value: 0.
	//
	// - **globalInterruptible**: Speech interruption configuration.
	//
	// - **engine**: Invoked service; options: [ali, volc, xunfei]. xunfei is not supported in LLM scenarios.
	//
	// - **nlsServiceType**: Service type for invocation
	//
	//    - Managed: Public cloud NLS service
	//
	//    - Authorized: Authorized NLS service
	//
	//    - Provided: Customer-provided NLS service via AccessKey/SecretKey
	//
	//    - Apes: Private Cloud service
	//
	// - **engineXunfei**: Configuration when the service provider is xunfei.
	//
	// > 1. When engine is set to ali and nlsServiceType is Authorized, a custom service is used.
	//
	// > 2. When the service provider is ali, engine is ali, and nlsServiceType is Managed, the default service is used.
	//
	// > 3. When engine is xunfei (applicable only to small-model scenarios) and nlsServiceType is Authorized, the service provider is xunfei, and engineXunfei must be configured as {"pitchRate":50,"speechRate":50,"voice":"aisjiuxu","volume":50}.
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

func (s ModifyScriptRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyScriptRequest) GoString() string {
	return s.String()
}

func (s *ModifyScriptRequest) GetAgentId() *int64 {
	return s.AgentId
}

func (s *ModifyScriptRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *ModifyScriptRequest) GetAgentLlm() *bool {
	return s.AgentLlm
}

func (s *ModifyScriptRequest) GetAsrConfig() *string {
	return s.AsrConfig
}

func (s *ModifyScriptRequest) GetChatConfig() *string {
	return s.ChatConfig
}

func (s *ModifyScriptRequest) GetChatbotId() *string {
	return s.ChatbotId
}

func (s *ModifyScriptRequest) GetEmotionEnable() *bool {
	return s.EmotionEnable
}

func (s *ModifyScriptRequest) GetIndustry() *string {
	return s.Industry
}

func (s *ModifyScriptRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyScriptRequest) GetLabelConfig() *string {
	return s.LabelConfig
}

func (s *ModifyScriptRequest) GetLongWaitEnable() *bool {
	return s.LongWaitEnable
}

func (s *ModifyScriptRequest) GetMiniPlaybackConfigListJsonString() *string {
	return s.MiniPlaybackConfigListJsonString
}

func (s *ModifyScriptRequest) GetMiniPlaybackEnable() *bool {
	return s.MiniPlaybackEnable
}

func (s *ModifyScriptRequest) GetNewBargeInEnable() *bool {
	return s.NewBargeInEnable
}

func (s *ModifyScriptRequest) GetNlsConfig() *string {
	return s.NlsConfig
}

func (s *ModifyScriptRequest) GetNluAccessType() *string {
	return s.NluAccessType
}

func (s *ModifyScriptRequest) GetNluEngine() *string {
	return s.NluEngine
}

func (s *ModifyScriptRequest) GetScene() *string {
	return s.Scene
}

func (s *ModifyScriptRequest) GetScriptContent() []*string {
	return s.ScriptContent
}

func (s *ModifyScriptRequest) GetScriptDescription() *string {
	return s.ScriptDescription
}

func (s *ModifyScriptRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *ModifyScriptRequest) GetScriptName() *string {
	return s.ScriptName
}

func (s *ModifyScriptRequest) GetScriptWaveform() []*string {
	return s.ScriptWaveform
}

func (s *ModifyScriptRequest) GetTtsConfig() *string {
	return s.TtsConfig
}

func (s *ModifyScriptRequest) SetAgentId(v int64) *ModifyScriptRequest {
	s.AgentId = &v
	return s
}

func (s *ModifyScriptRequest) SetAgentKey(v string) *ModifyScriptRequest {
	s.AgentKey = &v
	return s
}

func (s *ModifyScriptRequest) SetAgentLlm(v bool) *ModifyScriptRequest {
	s.AgentLlm = &v
	return s
}

func (s *ModifyScriptRequest) SetAsrConfig(v string) *ModifyScriptRequest {
	s.AsrConfig = &v
	return s
}

func (s *ModifyScriptRequest) SetChatConfig(v string) *ModifyScriptRequest {
	s.ChatConfig = &v
	return s
}

func (s *ModifyScriptRequest) SetChatbotId(v string) *ModifyScriptRequest {
	s.ChatbotId = &v
	return s
}

func (s *ModifyScriptRequest) SetEmotionEnable(v bool) *ModifyScriptRequest {
	s.EmotionEnable = &v
	return s
}

func (s *ModifyScriptRequest) SetIndustry(v string) *ModifyScriptRequest {
	s.Industry = &v
	return s
}

func (s *ModifyScriptRequest) SetInstanceId(v string) *ModifyScriptRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyScriptRequest) SetLabelConfig(v string) *ModifyScriptRequest {
	s.LabelConfig = &v
	return s
}

func (s *ModifyScriptRequest) SetLongWaitEnable(v bool) *ModifyScriptRequest {
	s.LongWaitEnable = &v
	return s
}

func (s *ModifyScriptRequest) SetMiniPlaybackConfigListJsonString(v string) *ModifyScriptRequest {
	s.MiniPlaybackConfigListJsonString = &v
	return s
}

func (s *ModifyScriptRequest) SetMiniPlaybackEnable(v bool) *ModifyScriptRequest {
	s.MiniPlaybackEnable = &v
	return s
}

func (s *ModifyScriptRequest) SetNewBargeInEnable(v bool) *ModifyScriptRequest {
	s.NewBargeInEnable = &v
	return s
}

func (s *ModifyScriptRequest) SetNlsConfig(v string) *ModifyScriptRequest {
	s.NlsConfig = &v
	return s
}

func (s *ModifyScriptRequest) SetNluAccessType(v string) *ModifyScriptRequest {
	s.NluAccessType = &v
	return s
}

func (s *ModifyScriptRequest) SetNluEngine(v string) *ModifyScriptRequest {
	s.NluEngine = &v
	return s
}

func (s *ModifyScriptRequest) SetScene(v string) *ModifyScriptRequest {
	s.Scene = &v
	return s
}

func (s *ModifyScriptRequest) SetScriptContent(v []*string) *ModifyScriptRequest {
	s.ScriptContent = v
	return s
}

func (s *ModifyScriptRequest) SetScriptDescription(v string) *ModifyScriptRequest {
	s.ScriptDescription = &v
	return s
}

func (s *ModifyScriptRequest) SetScriptId(v string) *ModifyScriptRequest {
	s.ScriptId = &v
	return s
}

func (s *ModifyScriptRequest) SetScriptName(v string) *ModifyScriptRequest {
	s.ScriptName = &v
	return s
}

func (s *ModifyScriptRequest) SetScriptWaveform(v []*string) *ModifyScriptRequest {
	s.ScriptWaveform = v
	return s
}

func (s *ModifyScriptRequest) SetTtsConfig(v string) *ModifyScriptRequest {
	s.TtsConfig = &v
	return s
}

func (s *ModifyScriptRequest) Validate() error {
	return dara.Validate(s)
}
