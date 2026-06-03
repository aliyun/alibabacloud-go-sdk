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
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// none
	NlsConfig *string `json:"NlsConfig,omitempty" xml:"NlsConfig,omitempty"`
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// {}
	Script *DescribeScriptResponseBodyScript `json:"Script,omitempty" xml:"Script,omitempty" type:"Struct"`
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
	AgentId  *int64  `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	AgentLlm *bool   `json:"AgentLlm,omitempty" xml:"AgentLlm,omitempty"`
	// example:
	//
	// {\\"AppKey\\":\\"3GHttnsvir1FeWWb\\"}
	AsrConfig  *string `json:"AsrConfig,omitempty" xml:"AsrConfig,omitempty"`
	ChatConfig *string `json:"ChatConfig,omitempty" xml:"ChatConfig,omitempty"`
	// example:
	//
	// chatbot-cn-EJfqqa***
	ChatbotId *string `json:"ChatbotId,omitempty" xml:"ChatbotId,omitempty"`
	// example:
	//
	// DRAFTED
	DebugStatus *string `json:"DebugStatus,omitempty" xml:"DebugStatus,omitempty"`
	// example:
	//
	// true
	EmotionEnable *bool   `json:"EmotionEnable,omitempty" xml:"EmotionEnable,omitempty"`
	Industry      *string `json:"Industry,omitempty" xml:"Industry,omitempty"`
	// example:
	//
	// true
	IsDebugDrafted *bool `json:"IsDebugDrafted,omitempty" xml:"IsDebugDrafted,omitempty"`
	// example:
	//
	// true
	IsDrafted   *bool   `json:"IsDrafted,omitempty" xml:"IsDrafted,omitempty"`
	LabelConfig *string `json:"LabelConfig,omitempty" xml:"LabelConfig,omitempty"`
	// example:
	//
	// true
	LongWaitEnable *bool `json:"LongWaitEnable,omitempty" xml:"LongWaitEnable,omitempty"`
	// example:
	//
	// true
	MiniPlaybackEnable *bool `json:"MiniPlaybackEnable,omitempty" xml:"MiniPlaybackEnable,omitempty"`
	// example:
	//
	// true
	NewBargeInEnable  *bool                                       `json:"NewBargeInEnable,omitempty" xml:"NewBargeInEnable,omitempty"`
	NluEngine         *string                                     `json:"NluEngine,omitempty" xml:"NluEngine,omitempty"`
	NluProfile        *DescribeScriptResponseBodyScriptNluProfile `json:"NluProfile,omitempty" xml:"NluProfile,omitempty" type:"Struct"`
	Scene             *string                                     `json:"Scene,omitempty" xml:"Scene,omitempty"`
	ScriptDescription *string                                     `json:"ScriptDescription,omitempty" xml:"ScriptDescription,omitempty"`
	// example:
	//
	// 810b5872-57f0-4b27-80ab-7b3f4d8a6374
	ScriptId   *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	ScriptName *string `json:"ScriptName,omitempty" xml:"ScriptName,omitempty"`
	// example:
	//
	// DRAFTED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// {\\"voice\\":\\"xiaobei\\",\\"volume\\":\\"50\\",\\"speechRate\\":\\"-150\\",\\"pitchRate\\":\\"0\\"}
	TtsConfig *string `json:"TtsConfig,omitempty" xml:"TtsConfig,omitempty"`
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
	FcFunction           *string `json:"FcFunction,omitempty" xml:"FcFunction,omitempty"`
	FcHttpTriggerUrl     *string `json:"FcHttpTriggerUrl,omitempty" xml:"FcHttpTriggerUrl,omitempty"`
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
