// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAgentJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetModel(v string) *SubmitAgentJobRequest
	GetModel() *string
	SetNotifyUrl(v string) *SubmitAgentJobRequest
	GetNotifyUrl() *string
	SetPrompt(v string) *SubmitAgentJobRequest
	GetPrompt() *string
	SetSkill(v string) *SubmitAgentJobRequest
	GetSkill() *string
	SetUserData(v string) *SubmitAgentJobRequest
	GetUserData() *string
	SetWorkspaceId(v string) *SubmitAgentJobRequest
	GetWorkspaceId() *string
}

type SubmitAgentJobRequest struct {
	// The large language model (LLM) used to execute the agent task.
	//
	// example:
	//
	// qwen3.7-plus
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// The callback URL. Currently, only HTTP and HTTPS addresses are supported.
	//
	// example:
	//
	// https://api.ai-x.vip/callback
	NotifyUrl *string `json:"NotifyUrl,omitempty" xml:"NotifyUrl,omitempty"`
	// The prompt. Defined by the business as needed.
	//
	// This parameter is required.
	//
	// example:
	//
	// Compare the battle records of Yue Jin and Guan Yu in real history to determine who was stronger
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// The skill identifier, provided by the skill provider.
	//
	// example:
	//
	// wf://xxx
	Skill *string `json:"Skill,omitempty" xml:"Skill,omitempty"`
	// The custom user data. This value is returned as-is in the callback.
	//
	// example:
	//
	// {“x”: 1}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// ws_1151222932383236
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s SubmitAgentJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitAgentJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitAgentJobRequest) GetModel() *string {
	return s.Model
}

func (s *SubmitAgentJobRequest) GetNotifyUrl() *string {
	return s.NotifyUrl
}

func (s *SubmitAgentJobRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *SubmitAgentJobRequest) GetSkill() *string {
	return s.Skill
}

func (s *SubmitAgentJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitAgentJobRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *SubmitAgentJobRequest) SetModel(v string) *SubmitAgentJobRequest {
	s.Model = &v
	return s
}

func (s *SubmitAgentJobRequest) SetNotifyUrl(v string) *SubmitAgentJobRequest {
	s.NotifyUrl = &v
	return s
}

func (s *SubmitAgentJobRequest) SetPrompt(v string) *SubmitAgentJobRequest {
	s.Prompt = &v
	return s
}

func (s *SubmitAgentJobRequest) SetSkill(v string) *SubmitAgentJobRequest {
	s.Skill = &v
	return s
}

func (s *SubmitAgentJobRequest) SetUserData(v string) *SubmitAgentJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitAgentJobRequest) SetWorkspaceId(v string) *SubmitAgentJobRequest {
	s.WorkspaceId = &v
	return s
}

func (s *SubmitAgentJobRequest) Validate() error {
	return dara.Validate(s)
}
