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
	AgentId  *int64  `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	AgentLlm *bool   `json:"AgentLlm,omitempty" xml:"AgentLlm,omitempty"`
	// example:
	//
	// {\\"AppKey\\":\\"kknxKIhTTUcpCzYX\\"}
	AsrConfig  *string `json:"AsrConfig,omitempty" xml:"AsrConfig,omitempty"`
	ChatConfig *string `json:"ChatConfig,omitempty" xml:"ChatConfig,omitempty"`
	// example:
	//
	// chatbot-cn-iFZfi7eq6e
	ChatbotId *string `json:"ChatbotId,omitempty" xml:"ChatbotId,omitempty"`
	// example:
	//
	// true
	EmotionEnable *bool `json:"EmotionEnable,omitempty" xml:"EmotionEnable,omitempty"`
	// This parameter is required.
	Industry *string `json:"Industry,omitempty" xml:"Industry,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// c6320d3c-fa45-4011-b3b1-acdfabe3a8c6
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	LabelConfig *string `json:"LabelConfig,omitempty" xml:"LabelConfig,omitempty"`
	// example:
	//
	// true
	LongWaitEnable                   *bool   `json:"LongWaitEnable,omitempty" xml:"LongWaitEnable,omitempty"`
	MiniPlaybackConfigListJsonString *string `json:"MiniPlaybackConfigListJsonString,omitempty" xml:"MiniPlaybackConfigListJsonString,omitempty"`
	// example:
	//
	// true
	MiniPlaybackEnable *bool `json:"MiniPlaybackEnable,omitempty" xml:"MiniPlaybackEnable,omitempty"`
	// example:
	//
	// true
	NewBargeInEnable *bool   `json:"NewBargeInEnable,omitempty" xml:"NewBargeInEnable,omitempty"`
	NlsConfig        *string `json:"NlsConfig,omitempty" xml:"NlsConfig,omitempty"`
	NluAccessType    *string `json:"NluAccessType,omitempty" xml:"NluAccessType,omitempty"`
	NluEngine        *string `json:"NluEngine,omitempty" xml:"NluEngine,omitempty"`
	// This parameter is required.
	Scene             *string   `json:"Scene,omitempty" xml:"Scene,omitempty"`
	ScriptContent     []*string `json:"ScriptContent,omitempty" xml:"ScriptContent,omitempty" type:"Repeated"`
	ScriptDescription *string   `json:"ScriptDescription,omitempty" xml:"ScriptDescription,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// c153d0d8-ba04-41c0-8632-453944c9dd0b
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// This parameter is required.
	ScriptName     *string   `json:"ScriptName,omitempty" xml:"ScriptName,omitempty"`
	ScriptWaveform []*string `json:"ScriptWaveform,omitempty" xml:"ScriptWaveform,omitempty" type:"Repeated"`
	// example:
	//
	// {\\"voice\\":\\"siyue\\",\\"volume\\":\\"50\\",\\"speechRate\\":\\"-150\\",\\"pitchRate\\":\\"0\\"}
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
