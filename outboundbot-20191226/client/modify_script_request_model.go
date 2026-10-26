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
	// Specifies whether the robot workspace is a large model workspace.
	//
	// example:
	//
	// false
	AgentLlm *bool `json:"AgentLlm,omitempty" xml:"AgentLlm,omitempty"`
	// The ASR configuration. Parameter definitions:
	//
	// - **appKey**: The appKey of the Alibaba Cloud account.
	//
	// - **maxEndSilence**: The duration of voice endpoint detection.
	//
	// - **silenceTimeout**: The silence timeout in seconds. The timeout occurs after the user is silent for N seconds.
	//
	// - **engine**: The service to call. Valid values: ali and xunfei.
	//
	// - **nlsServiceType**: The type of the service to call. Valid values:
	//
	//    - Managed: The NLS service on the public cloud.
	//
	//    - Authorized: The authorized NLS service.
	//
	//    - Provided: The NLS service provided by the customer through AccessKey ID and AccessKey secret.
	//
	//    - Apes: The private cloud service.
	//
	// - **engineXunfei**: The configuration required when the caller is xunfei.
	//
	// > When engine is set to ali and nlsServiceType is set to Authorized, a custom service is used, and the service provider is ali. When engine is set to ali and nlsServiceType is set to Managed, the default service is used. When engine is set to xunfei and nlsServiceType is set to Authorized, the service provider is xunfei, and you must enter the xunfei configuration: {"uuid":"ed2xxxxxxxxx","globalMaxEndSilence":700,"globalMaxEndSilenceEnable":true}.
	//
	// - **globalMaxEndSilence**: The silence detection duration in milliseconds.
	//
	// - **globalMaxEndSilenceEnable**: The silence detection switch. Enabled by default.
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
	// - silenceConfig: The silence configuration.
	//
	//     - silenceReplyTimeout: The silence reply timeout in milliseconds.
	//
	//     - silenceTimeoutMaxCount: The maximum number of silence timeout rounds before hanging up.
	//
	//     - silenceTimeoutMaxCountEnable: Specifies whether to hang up after silence timeouts.
	//
	// - hangupConfig: The hang-up configuration.
	//
	//     - aiHangupEnable: Specifies whether to enable AI hang-up. Valid values: true and false.
	//
	//     - delayHangup: The delayed hang-up time. Maximum value: 10.
	//
	//     - hangupMaxRounds: The maximum number of interaction rounds. Maximum value: 100.
	//
	//     - hangupMaxRoundsBroadcast: The broadcast script for hanging up.
	//
	//     - hangupMaxRoundsEnable: Specifies whether to enable the maximum interaction rounds check. Valid values: true and false.
	//
	// - securityInterceptConfig: The security interception configuration.
	//
	//     - broadcast: The interception broadcast script.
	//
	// - specialInterceptConfig: The special case interception configuration.
	//
	//     - specialInterceptEnable: Specifies whether to enable special case interception.
	//
	//     - specialIntercepts: The special cases.
	//
	//         - voiceAssistant: Voice assistant.
	//
	//         - extensionNumberTransfer: Extension number transfer.
	//
	// - transitionConfig: The backchannel model configuration.
	//
	//     - transitionSwitch: Specifies whether to enable the backchannel model.
	//
	// example:
	//
	// {"silenceConfig":{"silenceReplyTimeout":499,"silenceTimeoutMaxCount":10,"silenceTimeoutMaxCountEnable":true},"hangupConfig":{"aiHangupEnable":false,"delayHangup":0,"hangupMaxRounds":20,"hangupMaxRoundsBroadcast":"Thank you for taking my call. Have a great day. Goodbye!","hangupMaxRoundsEnable":false},"securityInterceptConfig":{"broadcast":"I cannot answer this question. You can try asking other questions."},"specialInterceptConfig":{"specialInterceptEnable":false,"specialIntercepts":[{"code":"voiceAssistant"},{"code":"extensionNumberTransfer"}]},"transitionConfig":{"transitionSwitch":false}}
	ChatConfig *string `json:"ChatConfig,omitempty" xml:"ChatConfig,omitempty"`
	// The robot ID.
	//
	// example:
	//
	// chatbot-cn-iFZfi7eq6e
	ChatbotId *string `json:"ChatbotId,omitempty" xml:"ChatbotId,omitempty"`
	// Specifies whether to enable the emotion recognition configuration.
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
	// E-commerce
	Industry *string `json:"Industry,omitempty" xml:"Industry,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c6320d3c-fa45-4011-b3b1-acdfabe3a8c6
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The label definition.
	//
	// > Used in large model flow canvas scenarios.
	//
	// example:
	//
	// {
	//
	// 	"labels": [{
	//
	// 		"description": "The customer\\"s final repayment intention after the conversation is completed",
	//
	// 		"name": "Whether promised to repay",
	//
	// 		"valueList": "[\\"Yes\\",\\"No\\"]"
	//
	// 	}]
	//
	// }
	LabelConfig *string `json:"LabelConfig,omitempty" xml:"LabelConfig,omitempty"`
	// Specifies whether to enable the intelligent sentence segmentation configuration for small models.
	//
	// example:
	//
	// true
	LongWaitEnable *bool `json:"LongWaitEnable,omitempty" xml:"LongWaitEnable,omitempty"`
	// The backchannel configuration. Customization is not supported at the moment. You do not need to pass this parameter by default. **[Deprecated]**
	//
	// example:
	//
	// None
	MiniPlaybackConfigListJsonString *string `json:"MiniPlaybackConfigListJsonString,omitempty" xml:"MiniPlaybackConfigListJsonString,omitempty"`
	// Specifies whether to enable the backchannel configuration for small models.
	//
	// example:
	//
	// true
	MiniPlaybackEnable *bool `json:"MiniPlaybackEnable,omitempty" xml:"MiniPlaybackEnable,omitempty"`
	// Specifies whether to enable the graceful interruption configuration for small models.
	//
	// example:
	//
	// true
	NewBargeInEnable *bool `json:"NewBargeInEnable,omitempty" xml:"NewBargeInEnable,omitempty"`
	// **[Deprecated]**
	//
	// example:
	//
	// None
	NlsConfig *string `json:"NlsConfig,omitempty" xml:"NlsConfig,omitempty"`
	// The NLU access method. This parameter is applicable only to large model scenarios.
	//
	// Valid values:
	//
	// - Managed: Uses the Alibaba Cloud public account for access.
	//
	// - Empty: For non-large model scenarios.
	//
	// example:
	//
	// Managed
	NluAccessType *string `json:"NluAccessType,omitempty" xml:"NluAccessType,omitempty"`
	// The NLU engine. This parameter is applicable only to large model scenarios.
	//
	// > After a scenario is created, the scenario mode cannot be modified.
	//
	// Valid values:
	//
	// - Prompts: The text input mode.
	//
	// - SSE_FUNCTION: The Function Compute mode.
	//
	// - Empty: For non-large model scenarios.
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
	// Follow-up visit
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// The script list for notification instances. **[Deprecated]**
	ScriptContent []*string `json:"ScriptContent,omitempty" xml:"ScriptContent,omitempty" type:"Repeated"`
	// The script description.
	//
	// example:
	//
	// Rework callback script
	ScriptDescription *string `json:"ScriptDescription,omitempty" xml:"ScriptDescription,omitempty"`
	// The script ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c153d0d8-ba04-41c0-8632-453944c9dd0b
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// The script name.
	//
	// This parameter is required.
	//
	// example:
	//
	// Renewal reminder
	ScriptName *string `json:"ScriptName,omitempty" xml:"ScriptName,omitempty"`
	// The script voice list for notification instances. **[Deprecated]**
	ScriptWaveform []*string `json:"ScriptWaveform,omitempty" xml:"ScriptWaveform,omitempty" type:"Repeated"`
	// The TTS configuration. Parameter definitions:
	//
	// - **voice**: The speaker.
	//
	// - **volume**: The volume. Valid values: 0 to 100. Default value: 50.
	//
	// - **speechRate**: The speech rate. Valid values: -500 to 500. Default value: 0.
	//
	// - **pitchRate**: The pitch. Valid values: -500 to 500. Default value: 0.
	//
	// - **openingPromptInterruptible**: The interruption configuration for the opening prompt.
	//
	// - **closingStatementInterruptible**: The interruption configuration for the closing statement.
	//
	// - **globalInterruptible**: The voice interruption configuration.
	//
	// - **engine**: The service to call. Valid values: ali, volc, and xunfei. xunfei is not supported in large model scenarios.
	//
	// - **nlsServiceType**: The type of the service to call. Valid values:
	//
	//    - Managed: The NLS service on the public cloud.
	//
	//    - Authorized: The authorized NLS service.
	//
	//    - Provided: The NLS service provided by the customer through AccessKey ID and AccessKey secret.
	//
	//    - Apes: The private cloud service.
	//
	// - **engineXunfei**: The configuration when the service provider is xunfei.
	//
	// > 1. When engine is set to ali and nlsServiceType is set to Authorized, a custom service is used. 2. When the service provider is ali, engine is set to ali, and nlsServiceType is set to Managed, the default service is used. 3. When engine is set to xunfei (applicable to small model scenarios) and nlsServiceType is set to Authorized, the service provider is xunfei, and you must enter the engineXunfei configuration: {"pitchRate":50,"speechRate":50,"voice":"aisjiuxu","volume":50}. 4. When engine is set to volc and nlsServiceType is set to Authorized, the applicable service is doubao.
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
