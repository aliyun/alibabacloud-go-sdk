// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallableAgentsShrink(v string) *CreateAgentShrinkRequest
	GetCallableAgentsShrink() *string
	SetDescription(v string) *CreateAgentShrinkRequest
	GetDescription() *string
	SetDisplayName(v string) *CreateAgentShrinkRequest
	GetDisplayName() *string
	SetExpertSuitesShrink(v string) *CreateAgentShrinkRequest
	GetExpertSuitesShrink() *string
	SetMetadataShrink(v string) *CreateAgentShrinkRequest
	GetMetadataShrink() *string
	SetModelShrink(v string) *CreateAgentShrinkRequest
	GetModelShrink() *string
	SetName(v string) *CreateAgentShrinkRequest
	GetName() *string
	SetPublishImmediately(v bool) *CreateAgentShrinkRequest
	GetPublishImmediately() *bool
	SetSkillsShrink(v string) *CreateAgentShrinkRequest
	GetSkillsShrink() *string
	SetSystemPrompt(v string) *CreateAgentShrinkRequest
	GetSystemPrompt() *string
	SetToolsShrink(v string) *CreateAgentShrinkRequest
	GetToolsShrink() *string
	SetVisibility(v string) *CreateAgentShrinkRequest
	GetVisibility() *string
	SetVisibilityScopeShrink(v string) *CreateAgentShrinkRequest
	GetVisibilityScopeShrink() *string
}

type CreateAgentShrinkRequest struct {
	// The list of child agents that can be called by this agent.
	//
	// example:
	//
	// -
	CallableAgentsShrink *string `json:"CallableAgents,omitempty" xml:"CallableAgents,omitempty"`
	// The description of the agent.
	//
	// example:
	//
	// Data analytics assistant
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The display name of the agent.
	//
	// example:
	//
	// My Assistant
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// Agent 绑定的已发布专家套件。个人或企业套件使用 SuiteId，可选 Version；官方套件必须显式指定 Source=OFFICIAL 和官方目录中的 SuiteId，不传 Version。未传时创建为空，不会默认加载所有官方套件。
	//
	// example:
	//
	// [{"Source":"OFFICIAL","SuiteId":"dataworks-data-analysis-extension"},{"SuiteId":"suite_example","Version":1}]
	ExpertSuitesShrink *string `json:"ExpertSuites,omitempty" xml:"ExpertSuites,omitempty"`
	// The extended metadata in key-value pairs.
	//
	// example:
	//
	// {}
	MetadataShrink *string `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// The model configuration.
	//
	// example:
	//
	// {
	//
	//           "modelName": "dataworks-public-bailian/qwen-max"
	//
	//         }
	ModelShrink *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// The name of the agent. The name must be unique within the current account.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-agent
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 是否在创建 Agent 后立即发布。true 发布创建的版本；false 或省略保持原有仅创建草稿行为。保存与发布不是原子操作，发布失败时草稿可能已存在；重试前应先查询 Agent，避免重复创建。
	//
	// example:
	//
	// false
	PublishImmediately *bool `json:"PublishImmediately,omitempty" xml:"PublishImmediately,omitempty"`
	// The list of skills.
	//
	// example:
	//
	// -
	SkillsShrink *string `json:"Skills,omitempty" xml:"Skills,omitempty"`
	// The system prompt.
	//
	// example:
	//
	// You are a data analytics assistant
	SystemPrompt *string `json:"SystemPrompt,omitempty" xml:"SystemPrompt,omitempty"`
	// The list of tools.
	//
	// example:
	//
	// -
	ToolsShrink *string `json:"Tools,omitempty" xml:"Tools,omitempty"`
	// The visibility level.<br>
	//
	// `TENANT`: Visible within the account.<br>
	//
	// `PROJECT`: Visible to specified projects.<br>
	//
	// `USER`: Visible to specified users.
	//
	// example:
	//
	// TENANT
	Visibility *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
	// The visibility scope. The corresponding field is used based on the visibility level.
	VisibilityScopeShrink *string `json:"VisibilityScope,omitempty" xml:"VisibilityScope,omitempty"`
}

func (s CreateAgentShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateAgentShrinkRequest) GetCallableAgentsShrink() *string {
	return s.CallableAgentsShrink
}

func (s *CreateAgentShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateAgentShrinkRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateAgentShrinkRequest) GetExpertSuitesShrink() *string {
	return s.ExpertSuitesShrink
}

func (s *CreateAgentShrinkRequest) GetMetadataShrink() *string {
	return s.MetadataShrink
}

func (s *CreateAgentShrinkRequest) GetModelShrink() *string {
	return s.ModelShrink
}

func (s *CreateAgentShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateAgentShrinkRequest) GetPublishImmediately() *bool {
	return s.PublishImmediately
}

func (s *CreateAgentShrinkRequest) GetSkillsShrink() *string {
	return s.SkillsShrink
}

func (s *CreateAgentShrinkRequest) GetSystemPrompt() *string {
	return s.SystemPrompt
}

func (s *CreateAgentShrinkRequest) GetToolsShrink() *string {
	return s.ToolsShrink
}

func (s *CreateAgentShrinkRequest) GetVisibility() *string {
	return s.Visibility
}

func (s *CreateAgentShrinkRequest) GetVisibilityScopeShrink() *string {
	return s.VisibilityScopeShrink
}

func (s *CreateAgentShrinkRequest) SetCallableAgentsShrink(v string) *CreateAgentShrinkRequest {
	s.CallableAgentsShrink = &v
	return s
}

func (s *CreateAgentShrinkRequest) SetDescription(v string) *CreateAgentShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateAgentShrinkRequest) SetDisplayName(v string) *CreateAgentShrinkRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateAgentShrinkRequest) SetExpertSuitesShrink(v string) *CreateAgentShrinkRequest {
	s.ExpertSuitesShrink = &v
	return s
}

func (s *CreateAgentShrinkRequest) SetMetadataShrink(v string) *CreateAgentShrinkRequest {
	s.MetadataShrink = &v
	return s
}

func (s *CreateAgentShrinkRequest) SetModelShrink(v string) *CreateAgentShrinkRequest {
	s.ModelShrink = &v
	return s
}

func (s *CreateAgentShrinkRequest) SetName(v string) *CreateAgentShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateAgentShrinkRequest) SetPublishImmediately(v bool) *CreateAgentShrinkRequest {
	s.PublishImmediately = &v
	return s
}

func (s *CreateAgentShrinkRequest) SetSkillsShrink(v string) *CreateAgentShrinkRequest {
	s.SkillsShrink = &v
	return s
}

func (s *CreateAgentShrinkRequest) SetSystemPrompt(v string) *CreateAgentShrinkRequest {
	s.SystemPrompt = &v
	return s
}

func (s *CreateAgentShrinkRequest) SetToolsShrink(v string) *CreateAgentShrinkRequest {
	s.ToolsShrink = &v
	return s
}

func (s *CreateAgentShrinkRequest) SetVisibility(v string) *CreateAgentShrinkRequest {
	s.Visibility = &v
	return s
}

func (s *CreateAgentShrinkRequest) SetVisibilityScopeShrink(v string) *CreateAgentShrinkRequest {
	s.VisibilityScopeShrink = &v
	return s
}

func (s *CreateAgentShrinkRequest) Validate() error {
	return dara.Validate(s)
}
