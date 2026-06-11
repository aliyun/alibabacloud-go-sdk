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
	// Robot workspace access Key
	//
	// example:
	//
	// 9137ab9c27044921860030adf8590ec4_p_outbound_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// Is the robot workspace a Large Language Model (LLM) workspace?
	//
	// example:
	//
	// false
	AgentLlm *bool `json:"AgentLlm,omitempty" xml:"AgentLlm,omitempty"`
	// ASR configuration. Parameter definitions:
	//
	// - **appKey**: Alibaba Cloud account appKey.
	//
	// - **maxEndSilence**: Voice endpoint detection duration.
	//
	// - **silenceTimeout**: Silence timeout. Unit: seconds. The user times out after N seconds of silence.
	//
	// - **engine**: Invoke service; [ali, xunfei]
	//
	// - **nlsServiceType**: Invoke service type [Managed, Authorized]
	//
	// - **engineXunfei**: If the caller is xunfei, enter the corresponding configuration.
	//
	// > If you select ali as the engine and Authorized as the nlsServiceType, a custom service is used, and the service provider is ali. If you select ali as the engine and Managed as the nlsServiceType, the default service is used. If you select xunfei as the engine and Authorized as the nlsServiceType, xunfei is the service provider. You must enter the xunfei configuration: {"uuid":"ed2xxxxxxxxx","globalMaxEndSilence":700,"globalMaxEndSilenceEnable":true}
	//
	// - **globalMaxEndSilence**: Silence detection. Unit: milliseconds.
	//
	// - **globalMaxEndSilenceEnable**: Silence detection switch. Enabled by default.
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
	// If the NluServiceType of the instance is Authorized or Provided, specify the ID of the chatbot instance to which the script needs to be attached using this field.
	//
	// example:
	//
	// chatbot-cn-IfaUfqaUnb
	ChatbotId *string `json:"ChatbotId,omitempty" xml:"ChatbotId,omitempty"`
	// Emotion detection configuration switch (applicable to small models)
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
	// Intelligent sentence segmentation configuration switch (applicable to small models)
	//
	// example:
	//
	// true
	LongWaitEnable *bool `json:"LongWaitEnable,omitempty" xml:"LongWaitEnable,omitempty"`
	// Connective phrase configuration switch (applicable to small models)
	//
	// example:
	//
	// true
	MiniPlaybackEnable *bool `json:"MiniPlaybackEnable,omitempty" xml:"MiniPlaybackEnable,omitempty"`
	// Graceful interruption configuration switch (applicable to small models)
	//
	// example:
	//
	// true
	NewBargeInEnable *bool `json:"NewBargeInEnable,omitempty" xml:"NewBargeInEnable,omitempty"`
	// NLU access method (applicable only to Large Language Model (LLM) scenarios). Enumeration: Managed - Access using an Alibaba public account. This field is empty for non-LLM scenarios.
	//
	// example:
	//
	// Managed
	NluAccessType *string `json:"NluAccessType,omitempty" xml:"NluAccessType,omitempty"`
	// NLU engine (applicable only to Large Language Model (LLM) scenarios). This field is empty for non-LLM scenarios.
	//
	// - Prompts - Large Language Model (LLM) scenario,
	//
	// - SSE_FUNCTION - Function Compute pattern.
	//
	// - BeeBot - Workflow pattern.
	//
	// example:
	//
	// Prompts
	NluEngine *string `json:"NluEngine,omitempty" xml:"NluEngine,omitempty"`
	// Scenario
	//
	// This parameter is required.
	//
	// example:
	//
	// 回访
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// For notification instances, pass in the script list. Deprecated.
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
	// Script name
	//
	// This parameter is required.
	//
	// example:
	//
	// 课程满意度回访
	ScriptName *string `json:"ScriptName,omitempty" xml:"ScriptName,omitempty"`
	// > If nluEngine is SSE_FUNCTION, you must pass in the corresponding configuration.
	//
	// Function Compute service pattern configuration
	//
	// - fcRegion: Function service region
	//
	// - fcFunction: Function service name
	//
	// - fcHttpTriggerUrl Function service trigger
	//
	// example:
	//
	// {"fcRegion":"cn-shanghai","fcFunction":"waihu_test","fcHttpTriggerUrl":"https://waihu-test.xxxxxxxxxxx.cn-shanghai-vpc.fcapp.run"}
	ScriptNluProfileJsonString *string `json:"ScriptNluProfileJsonString,omitempty" xml:"ScriptNluProfileJsonString,omitempty"`
	// For notification instances, pass in the script voice list. Deprecated.
	//
	// example:
	//
	// []
	ScriptWaveform []*string `json:"ScriptWaveform,omitempty" xml:"ScriptWaveform,omitempty" type:"Repeated"`
	// TTS configuration. Parameter definitions:
	//
	// - **voice**: Speaker.
	//
	// - **volume**: Volume. Value range: 0 to 100. Default value: 50.
	//
	// - **speechRate**: Speech rate. Value range: -500 to 500. Default value: 0.
	//
	// - **pitchRate**: Pitch rate. Value range: -500 to 500. Default value: 0.
	//
	// - **globalInterruptible**: Voice interruption configuration.
	//
	//   -**engine**: Invoke service; [ali, volc, xunfei]. Large Language Model (LLM) scenarios do not support xunfei.
	//
	// - **nlsServiceType**: Service type. [Managed, Authorized]
	//
	// - **engineXunfei**: Configuration when the service provider is xunfei.
	//
	// > 1\\. If you select ali as the engine and Authorized as the nlsServiceType, a custom service is used. 2. If the service provider is ali, and you select ali as the engine and Managed as the nlsServiceType, the default service is used. 3. If you select xunfei as the engine (applicable to small model scenarios) and Authorized as the nlsServiceType, xunfei is the service provider. You must enter the engineXunfei configuration: {"pitchRate":50,"speechRate":50,"voice":"aisjiuxu","volume":50}. 4. If you select volc as the engine and Authorized as the nlsServiceType, it applies to doubao.
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
