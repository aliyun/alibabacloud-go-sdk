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
	AgentId  *int64  `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	AgentLlm *bool   `json:"AgentLlm,omitempty" xml:"AgentLlm,omitempty"`
	// example:
	//
	// {\\"appKey\\":\\"kknxKIhTTUcpCzYX\\",\\"maxEndSilence\\":\\"400\\",\\"silenceTimeout\\":\\"5\\"}
	AsrConfig *string `json:"AsrConfig,omitempty" xml:"AsrConfig,omitempty"`
	// example:
	//
	// chatbot-cn-IfaUfqaUnb
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
	// c46001bc-3ead-4bfd-9a69-4b5b66a4a3f4
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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
	NewBargeInEnable *bool   `json:"NewBargeInEnable,omitempty" xml:"NewBargeInEnable,omitempty"`
	NluAccessType    *string `json:"NluAccessType,omitempty" xml:"NluAccessType,omitempty"`
	NluEngine        *string `json:"NluEngine,omitempty" xml:"NluEngine,omitempty"`
	// This parameter is required.
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// example:
	//
	// []
	ScriptContent     []*string `json:"ScriptContent,omitempty" xml:"ScriptContent,omitempty" type:"Repeated"`
	ScriptDescription *string   `json:"ScriptDescription,omitempty" xml:"ScriptDescription,omitempty"`
	// This parameter is required.
	ScriptName                 *string `json:"ScriptName,omitempty" xml:"ScriptName,omitempty"`
	ScriptNluProfileJsonString *string `json:"ScriptNluProfileJsonString,omitempty" xml:"ScriptNluProfileJsonString,omitempty"`
	// example:
	//
	// []
	ScriptWaveform []*string `json:"ScriptWaveform,omitempty" xml:"ScriptWaveform,omitempty" type:"Repeated"`
	// example:
	//
	// {\\"voice\\":\\"aixia\\",\\"volume\\":\\"50\\",\\"speechRate\\":\\"-150\\",\\"pitchRate\\":\\"0\\"}
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
