// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScriptResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeScriptResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DescribeScriptResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DescribeScriptResponseBody
	GetMessage() *string
	SetNlsConfig(v string) *DescribeScriptResponseBody
	GetNlsConfig() *string
	SetRequestId(v string) *DescribeScriptResponseBody
	GetRequestId() *string
	SetScript(v *DescribeScriptResponseBodyScript) *DescribeScriptResponseBody
	GetScript() *DescribeScriptResponseBodyScript
	SetSuccess(v bool) *DescribeScriptResponseBody
	GetSuccess() *bool
}

type DescribeScriptResponseBody struct {
	// API status code
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// API message
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Deprecated
	//
	// example:
	//
	// none
	NlsConfig *string `json:"NlsConfig,omitempty" xml:"NlsConfig,omitempty"`
	// Request ID
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Script
	//
	// example:
	//
	// {}
	Script *DescribeScriptResponseBodyScript `json:"Script,omitempty" xml:"Script,omitempty" type:"Struct"`
	// Indicates whether the request succeeded
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeScriptResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeScriptResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScriptResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeScriptResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeScriptResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeScriptResponseBody) GetNlsConfig() *string {
	return s.NlsConfig
}

func (s *DescribeScriptResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeScriptResponseBody) GetScript() *DescribeScriptResponseBodyScript {
	return s.Script
}

func (s *DescribeScriptResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeScriptResponseBody) SetCode(v string) *DescribeScriptResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeScriptResponseBody) SetHttpStatusCode(v int32) *DescribeScriptResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeScriptResponseBody) SetMessage(v string) *DescribeScriptResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeScriptResponseBody) SetNlsConfig(v string) *DescribeScriptResponseBody {
	s.NlsConfig = &v
	return s
}

func (s *DescribeScriptResponseBody) SetRequestId(v string) *DescribeScriptResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScriptResponseBody) SetScript(v *DescribeScriptResponseBodyScript) *DescribeScriptResponseBody {
	s.Script = v
	return s
}

func (s *DescribeScriptResponseBody) SetSuccess(v bool) *DescribeScriptResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeScriptResponseBody) Validate() error {
	if s.Script != nil {
		if err := s.Script.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeScriptResponseBodyScript struct {
	// Access key for the robot workspace
	//
	// example:
	//
	// 1218333
	AgentId *int64 `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// Access key for the chatbot workspace
	//
	// example:
	//
	// 14791f5f226b4878b3d9b676a0291234
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// Indicates whether the chatbot workspace is an LLM workspace.
	//
	// example:
	//
	// true
	AgentLlm *bool `json:"AgentLlm,omitempty" xml:"AgentLlm,omitempty"`
	// ASR configuration
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
	// Chat configuration
	//
	// example:
	//
	// {"silenceConfig":{"silenceReplyTimeout":499,"silenceTimeoutMaxCount":10,"silenceTimeoutMaxCountEnable":true},"hangupConfig":{"aiHangupEnable":false,"delayHangup":0,"hangupMaxRounds":20,"hangupMaxRoundsBroadcast":"感谢您的接听，祝您生活愉快，再见!","hangupMaxRoundsEnable":false},"securityInterceptConfig":{"broadcast":"您说的这个问题我不能回答您，您可以尝试询问其他问题。"},"specialInterceptConfig":{"specialInterceptEnable":false,"specialIntercepts":[{"code":"voiceAssistant"},{"code":"extensionNumberTransfer"}]},"transitionConfig":{"transitionSwitch":false}}
	ChatConfig *string `json:"ChatConfig,omitempty" xml:"ChatConfig,omitempty"`
	// Chatbot ID
	//
	// example:
	//
	// d31794e2a51f47d2901b4094d88311d7
	ChatbotId *string `json:"ChatbotId,omitempty" xml:"ChatbotId,omitempty"`
	// Debug status of the script
	//
	// example:
	//
	// DRAFTED
	DebugStatus *string `json:"DebugStatus,omitempty" xml:"DebugStatus,omitempty"`
	// Toggle for emotion detection configuration
	//
	// example:
	//
	// true
	EmotionEnable *bool `json:"EmotionEnable,omitempty" xml:"EmotionEnable,omitempty"`
	// Industry
	//
	// example:
	//
	// 金融
	Industry *string `json:"Industry,omitempty" xml:"Industry,omitempty"`
	// Indicates whether the debug version is in Draft status.
	//
	// example:
	//
	// true
	IsDebugDrafted *bool `json:"IsDebugDrafted,omitempty" xml:"IsDebugDrafted,omitempty"`
	// Indicates whether the script is in Draft status.
	//
	// example:
	//
	// true
	IsDrafted   *bool   `json:"IsDrafted,omitempty" xml:"IsDrafted,omitempty"`
	LabelConfig *string `json:"LabelConfig,omitempty" xml:"LabelConfig,omitempty"`
	// Toggle for Intelligent Sentence Segmentation configuration
	//
	// example:
	//
	// true
	LongWaitEnable *bool `json:"LongWaitEnable,omitempty" xml:"LongWaitEnable,omitempty"`
	// Toggle for tone continuity configuration
	//
	// example:
	//
	// true
	MiniPlaybackEnable *bool `json:"MiniPlaybackEnable,omitempty" xml:"MiniPlaybackEnable,omitempty"`
	// Graceful interruption configuration toggle
	//
	// example:
	//
	// true
	NewBargeInEnable *bool `json:"NewBargeInEnable,omitempty" xml:"NewBargeInEnable,omitempty"`
	// NLU engine (applicable only to LLM scenarios)
	//
	// - Prompts – Text input pattern
	//
	// - SSE_FUNCTION – Function Compute service pattern
	//
	// example:
	//
	// Prompts
	NluEngine *string `json:"NluEngine,omitempty" xml:"NluEngine,omitempty"`
	// Configuration for Function Compute service pattern
	NluProfile *DescribeScriptResponseBodyScriptNluProfile `json:"NluProfile,omitempty" xml:"NluProfile,omitempty" type:"Struct"`
	// Scenario
	//
	// example:
	//
	// 催收
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// Script description
	//
	// example:
	//
	// 催收话术
	ScriptDescription *string `json:"ScriptDescription,omitempty" xml:"ScriptDescription,omitempty"`
	// Script ID
	//
	// example:
	//
	// 810b5872-57f0-4b27-80ab-7b3f4d8a6374
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// Script Name
	//
	// example:
	//
	// 催收话术
	ScriptName *string `json:"ScriptName,omitempty" xml:"ScriptName,omitempty"`
	// Script status
	//
	// example:
	//
	// DRAFTED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// TTS configuration for the script
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
	// Update Time
	//
	// example:
	//
	// 1578881227000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeScriptResponseBodyScript) String() string {
	return dara.Prettify(s)
}

func (s DescribeScriptResponseBodyScript) GoString() string {
	return s.String()
}

func (s *DescribeScriptResponseBodyScript) GetAgentId() *int64 {
	return s.AgentId
}

func (s *DescribeScriptResponseBodyScript) GetAgentKey() *string {
	return s.AgentKey
}

func (s *DescribeScriptResponseBodyScript) GetAgentLlm() *bool {
	return s.AgentLlm
}

func (s *DescribeScriptResponseBodyScript) GetAsrConfig() *string {
	return s.AsrConfig
}

func (s *DescribeScriptResponseBodyScript) GetChatConfig() *string {
	return s.ChatConfig
}

func (s *DescribeScriptResponseBodyScript) GetChatbotId() *string {
	return s.ChatbotId
}

func (s *DescribeScriptResponseBodyScript) GetDebugStatus() *string {
	return s.DebugStatus
}

func (s *DescribeScriptResponseBodyScript) GetEmotionEnable() *bool {
	return s.EmotionEnable
}

func (s *DescribeScriptResponseBodyScript) GetIndustry() *string {
	return s.Industry
}

func (s *DescribeScriptResponseBodyScript) GetIsDebugDrafted() *bool {
	return s.IsDebugDrafted
}

func (s *DescribeScriptResponseBodyScript) GetIsDrafted() *bool {
	return s.IsDrafted
}

func (s *DescribeScriptResponseBodyScript) GetLabelConfig() *string {
	return s.LabelConfig
}

func (s *DescribeScriptResponseBodyScript) GetLongWaitEnable() *bool {
	return s.LongWaitEnable
}

func (s *DescribeScriptResponseBodyScript) GetMiniPlaybackEnable() *bool {
	return s.MiniPlaybackEnable
}

func (s *DescribeScriptResponseBodyScript) GetNewBargeInEnable() *bool {
	return s.NewBargeInEnable
}

func (s *DescribeScriptResponseBodyScript) GetNluEngine() *string {
	return s.NluEngine
}

func (s *DescribeScriptResponseBodyScript) GetNluProfile() *DescribeScriptResponseBodyScriptNluProfile {
	return s.NluProfile
}

func (s *DescribeScriptResponseBodyScript) GetScene() *string {
	return s.Scene
}

func (s *DescribeScriptResponseBodyScript) GetScriptDescription() *string {
	return s.ScriptDescription
}

func (s *DescribeScriptResponseBodyScript) GetScriptId() *string {
	return s.ScriptId
}

func (s *DescribeScriptResponseBodyScript) GetScriptName() *string {
	return s.ScriptName
}

func (s *DescribeScriptResponseBodyScript) GetStatus() *string {
	return s.Status
}

func (s *DescribeScriptResponseBodyScript) GetTtsConfig() *string {
	return s.TtsConfig
}

func (s *DescribeScriptResponseBodyScript) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *DescribeScriptResponseBodyScript) SetAgentId(v int64) *DescribeScriptResponseBodyScript {
	s.AgentId = &v
	return s
}

func (s *DescribeScriptResponseBodyScript) SetAgentKey(v string) *DescribeScriptResponseBodyScript {
	s.AgentKey = &v
	return s
}

func (s *DescribeScriptResponseBodyScript) SetAgentLlm(v bool) *DescribeScriptResponseBodyScript {
	s.AgentLlm = &v
	return s
}

func (s *DescribeScriptResponseBodyScript) SetAsrConfig(v string) *DescribeScriptResponseBodyScript {
	s.AsrConfig = &v
	return s
}

func (s *DescribeScriptResponseBodyScript) SetChatConfig(v string) *DescribeScriptResponseBodyScript {
	s.ChatConfig = &v
	return s
}

func (s *DescribeScriptResponseBodyScript) SetChatbotId(v string) *DescribeScriptResponseBodyScript {
	s.ChatbotId = &v
	return s
}

func (s *DescribeScriptResponseBodyScript) SetDebugStatus(v string) *DescribeScriptResponseBodyScript {
	s.DebugStatus = &v
	return s
}

func (s *DescribeScriptResponseBodyScript) SetEmotionEnable(v bool) *DescribeScriptResponseBodyScript {
	s.EmotionEnable = &v
	return s
}

func (s *DescribeScriptResponseBodyScript) SetIndustry(v string) *DescribeScriptResponseBodyScript {
	s.Industry = &v
	return s
}

func (s *DescribeScriptResponseBodyScript) SetIsDebugDrafted(v bool) *DescribeScriptResponseBodyScript {
	s.IsDebugDrafted = &v
	return s
}

func (s *DescribeScriptResponseBodyScript) SetIsDrafted(v bool) *DescribeScriptResponseBodyScript {
	s.IsDrafted = &v
	return s
}

func (s *DescribeScriptResponseBodyScript) SetLabelConfig(v string) *DescribeScriptResponseBodyScript {
	s.LabelConfig = &v
	return s
}

func (s *DescribeScriptResponseBodyScript) SetLongWaitEnable(v bool) *DescribeScriptResponseBodyScript {
	s.LongWaitEnable = &v
	return s
}

func (s *DescribeScriptResponseBodyScript) SetMiniPlaybackEnable(v bool) *DescribeScriptResponseBodyScript {
	s.MiniPlaybackEnable = &v
	return s
}

func (s *DescribeScriptResponseBodyScript) SetNewBargeInEnable(v bool) *DescribeScriptResponseBodyScript {
	s.NewBargeInEnable = &v
	return s
}

func (s *DescribeScriptResponseBodyScript) SetNluEngine(v string) *DescribeScriptResponseBodyScript {
	s.NluEngine = &v
	return s
}

func (s *DescribeScriptResponseBodyScript) SetNluProfile(v *DescribeScriptResponseBodyScriptNluProfile) *DescribeScriptResponseBodyScript {
	s.NluProfile = v
	return s
}

func (s *DescribeScriptResponseBodyScript) SetScene(v string) *DescribeScriptResponseBodyScript {
	s.Scene = &v
	return s
}

func (s *DescribeScriptResponseBodyScript) SetScriptDescription(v string) *DescribeScriptResponseBodyScript {
	s.ScriptDescription = &v
	return s
}

func (s *DescribeScriptResponseBodyScript) SetScriptId(v string) *DescribeScriptResponseBodyScript {
	s.ScriptId = &v
	return s
}

func (s *DescribeScriptResponseBodyScript) SetScriptName(v string) *DescribeScriptResponseBodyScript {
	s.ScriptName = &v
	return s
}

func (s *DescribeScriptResponseBodyScript) SetStatus(v string) *DescribeScriptResponseBodyScript {
	s.Status = &v
	return s
}

func (s *DescribeScriptResponseBodyScript) SetTtsConfig(v string) *DescribeScriptResponseBodyScript {
	s.TtsConfig = &v
	return s
}

func (s *DescribeScriptResponseBodyScript) SetUpdateTime(v int64) *DescribeScriptResponseBodyScript {
	s.UpdateTime = &v
	return s
}

func (s *DescribeScriptResponseBodyScript) Validate() error {
	if s.NluProfile != nil {
		if err := s.NluProfile.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeScriptResponseBodyScriptNluProfile struct {
	// Service Name of the function
	//
	// example:
	//
	// sanfang_test
	FcFunction *string `json:"FcFunction,omitempty" xml:"FcFunction,omitempty"`
	// Function Compute trigger
	//
	// example:
	//
	// http://sanfang_test-xxxxxx.cn-shanghai-vpc.fcapp.run
	FcHttpTriggerUrl *string `json:"FcHttpTriggerUrl,omitempty" xml:"FcHttpTriggerUrl,omitempty"`
	// Region of the Function Compute service
	//
	// example:
	//
	// cn-shanghai
	FcRegion             *string `json:"FcRegion,omitempty" xml:"FcRegion,omitempty"`
	SupportBeebotPrompts *bool   `json:"SupportBeebotPrompts,omitempty" xml:"SupportBeebotPrompts,omitempty"`
}

func (s DescribeScriptResponseBodyScriptNluProfile) String() string {
	return dara.Prettify(s)
}

func (s DescribeScriptResponseBodyScriptNluProfile) GoString() string {
	return s.String()
}

func (s *DescribeScriptResponseBodyScriptNluProfile) GetFcFunction() *string {
	return s.FcFunction
}

func (s *DescribeScriptResponseBodyScriptNluProfile) GetFcHttpTriggerUrl() *string {
	return s.FcHttpTriggerUrl
}

func (s *DescribeScriptResponseBodyScriptNluProfile) GetFcRegion() *string {
	return s.FcRegion
}

func (s *DescribeScriptResponseBodyScriptNluProfile) GetSupportBeebotPrompts() *bool {
	return s.SupportBeebotPrompts
}

func (s *DescribeScriptResponseBodyScriptNluProfile) SetFcFunction(v string) *DescribeScriptResponseBodyScriptNluProfile {
	s.FcFunction = &v
	return s
}

func (s *DescribeScriptResponseBodyScriptNluProfile) SetFcHttpTriggerUrl(v string) *DescribeScriptResponseBodyScriptNluProfile {
	s.FcHttpTriggerUrl = &v
	return s
}

func (s *DescribeScriptResponseBodyScriptNluProfile) SetFcRegion(v string) *DescribeScriptResponseBodyScriptNluProfile {
	s.FcRegion = &v
	return s
}

func (s *DescribeScriptResponseBodyScriptNluProfile) SetSupportBeebotPrompts(v bool) *DescribeScriptResponseBodyScriptNluProfile {
	s.SupportBeebotPrompts = &v
	return s
}

func (s *DescribeScriptResponseBodyScriptNluProfile) Validate() error {
	return dara.Validate(s)
}
