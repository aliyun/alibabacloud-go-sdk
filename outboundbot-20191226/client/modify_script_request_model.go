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
	// The ID of the robot workspace.
	//
	// example:
	//
	// 1198938
	AgentId *int64 `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// The access key for the robot workspace.
	//
	// example:
	//
	// 9137ab9c27044921860030adf8590ec4_p_outbound_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// Indicates whether the robot workspace is a Large Language Model (LLM) workspace.
	//
	// example:
	//
	// false
	AgentLlm *bool `json:"AgentLlm,omitempty" xml:"AgentLlm,omitempty"`
	// The ASR configuration. Parameter definitions:
	//
	// - **appKey**: The Alibaba Cloud account appKey.
	//
	// - **maxEndSilence**: The duration for voice endpoint detection.
	//
	// - **silenceTimeout**: The silence timeout. Unit: seconds. The system times out after the user is silent for N seconds.
	//
	// - **engine**: The service to invoke. Valid values: ali, xunfei.
	//
	// - **nlsServiceType**: The type of the invoked service.
	//
	//   - Managed: Public cloud NLS service.
	//
	//   - Authorized: Authorized NLS service.
	//
	//   - Provided: NLS service provided by the customer through AS/SK.
	//
	//   - Apes: Private cloud service.
	//
	// - **engineXunfei**: If the caller is xunfei, fill in the corresponding configuration.
	//
	// > If engine is set to ali and nlsServiceType is set to Authorized, a custom service is used, and the service provider is ali. If engine is set to ali and nlsServiceType is set to Managed, the default service is used. If engine is set to xunfei and nlsServiceType is set to Authorized, the service provider is xunfei. You must fill in the xunfei configuration, such as {"uuid":"ed2xxxxxxxxx","globalMaxEndSilence":700,"globalMaxEndSilenceEnable":true}.
	//
	// - **globalMaxEndSilence**: Silence detection. Unit: milliseconds.
	//
	// - **globalMaxEndSilenceEnable**: The switch for silence detection. Default value: enabled.
	//
	// - **speechNoiseThreshold**: The noise filtering threshold.
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
	// The call configuration.
	//
	// - silenceConfig: Silence configuration.
	//
	//   - silenceReplyTimeout: The timeout period for silence replies, in milliseconds.
	//
	//   - silenceTimeoutMaxCount: Hang up after several rounds of silence timeout.
	//
	//   - silenceTimeoutMaxCountEnable: Indicates whether to hang up on silence.
	//
	// - hangupConfig: Hang-up configuration.
	//
	//   - aiHangupEnable: AI hang-up. Valid values: true, false.
	//
	//   - delayHangup: Delayed hang-up. Maximum value: 10.
	//
	//   - hangupMaxRounds: Interaction rounds. Maximum value: 100.
	//
	//   - hangupMaxRoundsBroadcast: The script for hang-up broadcast.
	//
	//   - hangupMaxRoundsEnable: Determine the maximum number of interaction rounds. Valid values: true, false.
	//
	// - securityInterceptConfig: Security block configuration.
	//
	//   - broadcast: The script for block broadcast.
	//
	// - specialInterceptConfig: Special case block.
	//
	//   - specialInterceptEnable: The switch for special case block.
	//
	//   - specialIntercepts: Special cases.
	//
	//     - voiceAssistant: Voice assistant.
	//
	//     - extensionNumberTransfer: Extension number transfer.
	//
	// - transitionConfig: Transition phrase model configuration.
	//
	//   - transitionSwitch: The switch for the transition phrase model.
	//
	// example:
	//
	// {"silenceConfig":{"silenceReplyTimeout":499,"silenceTimeoutMaxCount":10,"silenceTimeoutMaxCountEnable":true},"hangupConfig":{"aiHangupEnable":false,"delayHangup":0,"hangupMaxRounds":20,"hangupMaxRoundsBroadcast":"感谢您的接听，祝您生活愉快，再见!","hangupMaxRoundsEnable":false},"securityInterceptConfig":{"broadcast":"您说的这个问题我不能回答您，您可以尝试询问其他问题。"},"specialInterceptConfig":{"specialInterceptEnable":false,"specialIntercepts":[{"code":"voiceAssistant"},{"code":"extensionNumberTransfer"}]},"transitionConfig":{"transitionSwitch":false}}
	ChatConfig *string `json:"ChatConfig,omitempty" xml:"ChatConfig,omitempty"`
	// The ID of the chatbot.
	//
	// example:
	//
	// chatbot-cn-iFZfi7eq6e
	ChatbotId *string `json:"ChatbotId,omitempty" xml:"ChatbotId,omitempty"`
	// The switch for emotion detection configuration.
	//
	// example:
	//
	// true
	EmotionEnable *bool `json:"EmotionEnable,omitempty" xml:"EmotionEnable,omitempty"`
	// The industry.
	//
	// This parameter is required.
	//
	// example:
	//
	// 电商
	Industry *string `json:"Industry,omitempty" xml:"Industry,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// c6320d3c-fa45-4011-b3b1-acdfabe3a8c6
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	LabelConfig *string `json:"LabelConfig,omitempty" xml:"LabelConfig,omitempty"`
	// The switch for intelligent sentence segmentation configuration (small model).
	//
	// example:
	//
	// true
	LongWaitEnable *bool `json:"LongWaitEnable,omitempty" xml:"LongWaitEnable,omitempty"`
	// The configuration for transition phrases. Customization is not supported temporarily. Do not pass this parameter by default. This parameter is deprecated.
	//
	// example:
	//
	// 无
	MiniPlaybackConfigListJsonString *string `json:"MiniPlaybackConfigListJsonString,omitempty" xml:"MiniPlaybackConfigListJsonString,omitempty"`
	// The switch for transition phrase configuration (small model).
	//
	// example:
	//
	// true
	MiniPlaybackEnable *bool `json:"MiniPlaybackEnable,omitempty" xml:"MiniPlaybackEnable,omitempty"`
	// The switch for graceful barge-in configuration (small model).
	//
	// example:
	//
	// true
	NewBargeInEnable *bool `json:"NewBargeInEnable,omitempty" xml:"NewBargeInEnable,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// 无
	NlsConfig *string `json:"NlsConfig,omitempty" xml:"NlsConfig,omitempty"`
	// The NLU access method (applicable only to Large Language Model (LLM) scenarios).
	//
	// Enumeration:
	//
	// - Managed: Access using an Alibaba Cloud public account.
	//
	// - This parameter is empty for non-LLM scenarios.
	//
	// example:
	//
	// Managed
	NluAccessType *string `json:"NluAccessType,omitempty" xml:"NluAccessType,omitempty"`
	// The NLU engine (applicable only to Large Language Model (LLM) scenarios).
	//
	// > After a scenario is created, you cannot modify the scenario mode.
	//
	// Enumeration:
	//
	// - Prompts: Text filling mode.
	//
	// - SSE_FUNCTION: Function Compute mode.
	//
	// - This parameter is empty for non-LLM scenarios.
	//
	// example:
	//
	// Prompts
	NluEngine *string `json:"NluEngine,omitempty" xml:"NluEngine,omitempty"`
	// The scenario information.
	//
	// This parameter is required.
	//
	// example:
	//
	// 回访
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// For notification instances, pass in the script list. This parameter is deprecated.
	ScriptContent []*string `json:"ScriptContent,omitempty" xml:"ScriptContent,omitempty" type:"Repeated"`
	// The description of the script.
	//
	// example:
	//
	// 返工回访话术
	ScriptDescription *string `json:"ScriptDescription,omitempty" xml:"ScriptDescription,omitempty"`
	// The ID of the script.
	//
	// This parameter is required.
	//
	// example:
	//
	// c153d0d8-ba04-41c0-8632-453944c9dd0b
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// The name of the script.
	//
	// This parameter is required.
	//
	// example:
	//
	// 续费提醒
	ScriptName *string `json:"ScriptName,omitempty" xml:"ScriptName,omitempty"`
	// For notification instances, pass in the script voice list. This parameter is deprecated.
	ScriptWaveform []*string `json:"ScriptWaveform,omitempty" xml:"ScriptWaveform,omitempty" type:"Repeated"`
	// The TTS configuration. Parameter definitions:
	//
	// - **voice**: The voice actor.
	//
	// - **volume**: The volume. Valid values: 0 to 100. Default value: 50.
	//
	// - **speechRate**: The speech rate. Valid values: -500 to 500. Default value: 0.
	//
	// - **pitchRate**: The pitch rate. Valid values: -500 to 500. Default value: 0.
	//
	// - **globalInterruptible**: The voice interruption configuration.
	//
	// - **engine**
	//
	//   **nlsServiceType**: The type of the invoked service.
	//
	//   - Managed: Public cloud NLS service.
	//
	//   - Authorized: Authorized NLS service.
	//
	//   - Provided: NLS service provided by the customer through AS/SK.
	//
	//   - Apes: Private cloud service.
	//
	// - **engineXunfei**: The configuration when the service provider is xunfei.
	//
	// > 1\\. If engine is set to ali and nlsServiceType is set to Authorized, a custom service is used. 2. If the service provider is ali, and engine is set to ali and nlsServiceType is set to Managed, the default service is used. 3. If engine is set to xunfei (applicable to small model scenarios) and nlsServiceType is set to Authorized, the service provider is xunfei. You must fill in the engineXunfei configuration, such as {"pitchRate":50,"speechRate":50,"voice":"aisjiuxu","volume":50}. 4. If engine is set to volc and nlsServiceType is set to Authorized, it indicates that Doubao is applicable.
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
