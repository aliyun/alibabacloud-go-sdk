// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatUserSecAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgent(v string) *ChatUserSecAgentRequest
	GetAgent() *string
	SetAttachmentStagingId(v string) *ChatUserSecAgentRequest
	GetAttachmentStagingId() *string
	SetAttachments(v string) *ChatUserSecAgentRequest
	GetAttachments() *string
	SetChannel(v string) *ChatUserSecAgentRequest
	GetChannel() *string
	SetExecutionMode(v string) *ChatUserSecAgentRequest
	GetExecutionMode() *string
	SetExtraParams(v string) *ChatUserSecAgentRequest
	GetExtraParams() *string
	SetMemory(v bool) *ChatUserSecAgentRequest
	GetMemory() *bool
	SetModel(v string) *ChatUserSecAgentRequest
	GetModel() *string
	SetPrompt(v string) *ChatUserSecAgentRequest
	GetPrompt() *string
	SetResponseLanguage(v string) *ChatUserSecAgentRequest
	GetResponseLanguage() *string
	SetSessionId(v string) *ChatUserSecAgentRequest
	GetSessionId() *string
	SetSkill(v string) *ChatUserSecAgentRequest
	GetSkill() *string
	SetStream(v bool) *ChatUserSecAgentRequest
	GetStream() *bool
	SetTalkId(v string) *ChatUserSecAgentRequest
	GetTalkId() *string
	SetTarget(v string) *ChatUserSecAgentRequest
	GetTarget() *string
	SetTimeZone(v string) *ChatUserSecAgentRequest
	GetTimeZone() *string
	SetUserInputInfo(v string) *ChatUserSecAgentRequest
	GetUserInputInfo() *string
}

type ChatUserSecAgentRequest struct {
	// example:
	//
	// sec-ops-agent
	Agent *string `json:"Agent,omitempty" xml:"Agent,omitempty"`
	// 附件暂存 ID
	//
	// example:
	//
	// stg-6f1d9c8b7a2e4530
	AttachmentStagingId *string `json:"AttachmentStagingId,omitempty" xml:"AttachmentStagingId,omitempty"`
	// 附件列表 JSON 字符串
	//
	// example:
	//
	// ["oss/input.txt","oss/raw.txt"]
	Attachments *string `json:"Attachments,omitempty" xml:"Attachments,omitempty"`
	// 逻辑渠道名
	//
	// example:
	//
	// console
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	// 执行模式: single/team/role
	//
	// example:
	//
	// single
	ExecutionMode *string `json:"ExecutionMode,omitempty" xml:"ExecutionMode,omitempty"`
	// 扩展参数 JSON 字符串，如 execution_mode、target 等
	//
	// example:
	//
	// {"execution_mode":"team","target":"team:auto"}
	ExtraParams *string `json:"ExtraParams,omitempty" xml:"ExtraParams,omitempty"`
	Memory      *bool   `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// example:
	//
	// qwen-max
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// 用户提问；新会话时必填，恢复/交互时可空
	//
	// example:
	//
	// 帮我梳理最近 24 小时的高危告警并给出处置建议
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// example:
	//
	// zh-CN
	ResponseLanguage *string `json:"ResponseLanguage,omitempty" xml:"ResponseLanguage,omitempty"`
	// example:
	//
	// 5f2c1b9a8d3e4c7f
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// alert-analysis
	Skill  *string `json:"Skill,omitempty" xml:"Skill,omitempty"`
	Stream *bool   `json:"Stream,omitempty" xml:"Stream,omitempty"`
	// example:
	//
	// 9b1e7d2c4a6f8e30
	TalkId *string `json:"TalkId,omitempty" xml:"TalkId,omitempty"`
	// 执行目标
	//
	// example:
	//
	// sec-ops-team-01
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// example:
	//
	// Asia/Shanghai
	TimeZone *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
	// 会话恢复/交互提交信息 JSON 字符串
	//
	// example:
	//
	// {"sessionId":"session_example","talkId":"talk_example","formId":"interaction_example","formValues":{"q1":{"kind":"selected","optionIds":["q1_o1"]}},"formAction":"submit"}
	UserInputInfo *string `json:"UserInputInfo,omitempty" xml:"UserInputInfo,omitempty"`
}

func (s ChatUserSecAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s ChatUserSecAgentRequest) GoString() string {
	return s.String()
}

func (s *ChatUserSecAgentRequest) GetAgent() *string {
	return s.Agent
}

func (s *ChatUserSecAgentRequest) GetAttachmentStagingId() *string {
	return s.AttachmentStagingId
}

func (s *ChatUserSecAgentRequest) GetAttachments() *string {
	return s.Attachments
}

func (s *ChatUserSecAgentRequest) GetChannel() *string {
	return s.Channel
}

func (s *ChatUserSecAgentRequest) GetExecutionMode() *string {
	return s.ExecutionMode
}

func (s *ChatUserSecAgentRequest) GetExtraParams() *string {
	return s.ExtraParams
}

func (s *ChatUserSecAgentRequest) GetMemory() *bool {
	return s.Memory
}

func (s *ChatUserSecAgentRequest) GetModel() *string {
	return s.Model
}

func (s *ChatUserSecAgentRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *ChatUserSecAgentRequest) GetResponseLanguage() *string {
	return s.ResponseLanguage
}

func (s *ChatUserSecAgentRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *ChatUserSecAgentRequest) GetSkill() *string {
	return s.Skill
}

func (s *ChatUserSecAgentRequest) GetStream() *bool {
	return s.Stream
}

func (s *ChatUserSecAgentRequest) GetTalkId() *string {
	return s.TalkId
}

func (s *ChatUserSecAgentRequest) GetTarget() *string {
	return s.Target
}

func (s *ChatUserSecAgentRequest) GetTimeZone() *string {
	return s.TimeZone
}

func (s *ChatUserSecAgentRequest) GetUserInputInfo() *string {
	return s.UserInputInfo
}

func (s *ChatUserSecAgentRequest) SetAgent(v string) *ChatUserSecAgentRequest {
	s.Agent = &v
	return s
}

func (s *ChatUserSecAgentRequest) SetAttachmentStagingId(v string) *ChatUserSecAgentRequest {
	s.AttachmentStagingId = &v
	return s
}

func (s *ChatUserSecAgentRequest) SetAttachments(v string) *ChatUserSecAgentRequest {
	s.Attachments = &v
	return s
}

func (s *ChatUserSecAgentRequest) SetChannel(v string) *ChatUserSecAgentRequest {
	s.Channel = &v
	return s
}

func (s *ChatUserSecAgentRequest) SetExecutionMode(v string) *ChatUserSecAgentRequest {
	s.ExecutionMode = &v
	return s
}

func (s *ChatUserSecAgentRequest) SetExtraParams(v string) *ChatUserSecAgentRequest {
	s.ExtraParams = &v
	return s
}

func (s *ChatUserSecAgentRequest) SetMemory(v bool) *ChatUserSecAgentRequest {
	s.Memory = &v
	return s
}

func (s *ChatUserSecAgentRequest) SetModel(v string) *ChatUserSecAgentRequest {
	s.Model = &v
	return s
}

func (s *ChatUserSecAgentRequest) SetPrompt(v string) *ChatUserSecAgentRequest {
	s.Prompt = &v
	return s
}

func (s *ChatUserSecAgentRequest) SetResponseLanguage(v string) *ChatUserSecAgentRequest {
	s.ResponseLanguage = &v
	return s
}

func (s *ChatUserSecAgentRequest) SetSessionId(v string) *ChatUserSecAgentRequest {
	s.SessionId = &v
	return s
}

func (s *ChatUserSecAgentRequest) SetSkill(v string) *ChatUserSecAgentRequest {
	s.Skill = &v
	return s
}

func (s *ChatUserSecAgentRequest) SetStream(v bool) *ChatUserSecAgentRequest {
	s.Stream = &v
	return s
}

func (s *ChatUserSecAgentRequest) SetTalkId(v string) *ChatUserSecAgentRequest {
	s.TalkId = &v
	return s
}

func (s *ChatUserSecAgentRequest) SetTarget(v string) *ChatUserSecAgentRequest {
	s.Target = &v
	return s
}

func (s *ChatUserSecAgentRequest) SetTimeZone(v string) *ChatUserSecAgentRequest {
	s.TimeZone = &v
	return s
}

func (s *ChatUserSecAgentRequest) SetUserInputInfo(v string) *ChatUserSecAgentRequest {
	s.UserInputInfo = &v
	return s
}

func (s *ChatUserSecAgentRequest) Validate() error {
	return dara.Validate(s)
}
