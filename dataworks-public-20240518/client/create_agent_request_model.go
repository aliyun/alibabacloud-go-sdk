// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallableAgents(v []*CreateAgentRequestCallableAgents) *CreateAgentRequest
	GetCallableAgents() []*CreateAgentRequestCallableAgents
	SetDescription(v string) *CreateAgentRequest
	GetDescription() *string
	SetDisplayName(v string) *CreateAgentRequest
	GetDisplayName() *string
	SetExpertSuites(v []*CreateAgentRequestExpertSuites) *CreateAgentRequest
	GetExpertSuites() []*CreateAgentRequestExpertSuites
	SetMetadata(v map[string]interface{}) *CreateAgentRequest
	GetMetadata() map[string]interface{}
	SetModel(v map[string]interface{}) *CreateAgentRequest
	GetModel() map[string]interface{}
	SetName(v string) *CreateAgentRequest
	GetName() *string
	SetPublishImmediately(v bool) *CreateAgentRequest
	GetPublishImmediately() *bool
	SetSkills(v []*CreateAgentRequestSkills) *CreateAgentRequest
	GetSkills() []*CreateAgentRequestSkills
	SetSystemPrompt(v string) *CreateAgentRequest
	GetSystemPrompt() *string
	SetTools(v []*CreateAgentRequestTools) *CreateAgentRequest
	GetTools() []*CreateAgentRequestTools
	SetVisibility(v string) *CreateAgentRequest
	GetVisibility() *string
	SetVisibilityScope(v *CreateAgentRequestVisibilityScope) *CreateAgentRequest
	GetVisibilityScope() *CreateAgentRequestVisibilityScope
}

type CreateAgentRequest struct {
	// The list of child agents that can be called by this agent.
	//
	// example:
	//
	// -
	CallableAgents []*CreateAgentRequestCallableAgents `json:"CallableAgents,omitempty" xml:"CallableAgents,omitempty" type:"Repeated"`
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
	ExpertSuites []*CreateAgentRequestExpertSuites `json:"ExpertSuites,omitempty" xml:"ExpertSuites,omitempty" type:"Repeated"`
	// The extended metadata in key-value pairs.
	//
	// example:
	//
	// {}
	Metadata map[string]interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// The model configuration.
	//
	// example:
	//
	// {
	//
	//           "modelName": "dataworks-public-bailian/qwen-max"
	//
	//         }
	Model map[string]interface{} `json:"Model,omitempty" xml:"Model,omitempty"`
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
	Skills []*CreateAgentRequestSkills `json:"Skills,omitempty" xml:"Skills,omitempty" type:"Repeated"`
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
	Tools []*CreateAgentRequestTools `json:"Tools,omitempty" xml:"Tools,omitempty" type:"Repeated"`
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
	VisibilityScope *CreateAgentRequestVisibilityScope `json:"VisibilityScope,omitempty" xml:"VisibilityScope,omitempty" type:"Struct"`
}

func (s CreateAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentRequest) GoString() string {
	return s.String()
}

func (s *CreateAgentRequest) GetCallableAgents() []*CreateAgentRequestCallableAgents {
	return s.CallableAgents
}

func (s *CreateAgentRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateAgentRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateAgentRequest) GetExpertSuites() []*CreateAgentRequestExpertSuites {
	return s.ExpertSuites
}

func (s *CreateAgentRequest) GetMetadata() map[string]interface{} {
	return s.Metadata
}

func (s *CreateAgentRequest) GetModel() map[string]interface{} {
	return s.Model
}

func (s *CreateAgentRequest) GetName() *string {
	return s.Name
}

func (s *CreateAgentRequest) GetPublishImmediately() *bool {
	return s.PublishImmediately
}

func (s *CreateAgentRequest) GetSkills() []*CreateAgentRequestSkills {
	return s.Skills
}

func (s *CreateAgentRequest) GetSystemPrompt() *string {
	return s.SystemPrompt
}

func (s *CreateAgentRequest) GetTools() []*CreateAgentRequestTools {
	return s.Tools
}

func (s *CreateAgentRequest) GetVisibility() *string {
	return s.Visibility
}

func (s *CreateAgentRequest) GetVisibilityScope() *CreateAgentRequestVisibilityScope {
	return s.VisibilityScope
}

func (s *CreateAgentRequest) SetCallableAgents(v []*CreateAgentRequestCallableAgents) *CreateAgentRequest {
	s.CallableAgents = v
	return s
}

func (s *CreateAgentRequest) SetDescription(v string) *CreateAgentRequest {
	s.Description = &v
	return s
}

func (s *CreateAgentRequest) SetDisplayName(v string) *CreateAgentRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateAgentRequest) SetExpertSuites(v []*CreateAgentRequestExpertSuites) *CreateAgentRequest {
	s.ExpertSuites = v
	return s
}

func (s *CreateAgentRequest) SetMetadata(v map[string]interface{}) *CreateAgentRequest {
	s.Metadata = v
	return s
}

func (s *CreateAgentRequest) SetModel(v map[string]interface{}) *CreateAgentRequest {
	s.Model = v
	return s
}

func (s *CreateAgentRequest) SetName(v string) *CreateAgentRequest {
	s.Name = &v
	return s
}

func (s *CreateAgentRequest) SetPublishImmediately(v bool) *CreateAgentRequest {
	s.PublishImmediately = &v
	return s
}

func (s *CreateAgentRequest) SetSkills(v []*CreateAgentRequestSkills) *CreateAgentRequest {
	s.Skills = v
	return s
}

func (s *CreateAgentRequest) SetSystemPrompt(v string) *CreateAgentRequest {
	s.SystemPrompt = &v
	return s
}

func (s *CreateAgentRequest) SetTools(v []*CreateAgentRequestTools) *CreateAgentRequest {
	s.Tools = v
	return s
}

func (s *CreateAgentRequest) SetVisibility(v string) *CreateAgentRequest {
	s.Visibility = &v
	return s
}

func (s *CreateAgentRequest) SetVisibilityScope(v *CreateAgentRequestVisibilityScope) *CreateAgentRequest {
	s.VisibilityScope = v
	return s
}

func (s *CreateAgentRequest) Validate() error {
	if s.CallableAgents != nil {
		for _, item := range s.CallableAgents {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ExpertSuites != nil {
		for _, item := range s.ExpertSuites {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Skills != nil {
		for _, item := range s.Skills {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tools != nil {
		for _, item := range s.Tools {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.VisibilityScope != nil {
		if err := s.VisibilityScope.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAgentRequestCallableAgents struct {
	// The name of the agent.
	//
	// example:
	//
	// agent-1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateAgentRequestCallableAgents) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentRequestCallableAgents) GoString() string {
	return s.String()
}

func (s *CreateAgentRequestCallableAgents) GetName() *string {
	return s.Name
}

func (s *CreateAgentRequestCallableAgents) SetName(v string) *CreateAgentRequestCallableAgents {
	s.Name = &v
	return s
}

func (s *CreateAgentRequestCallableAgents) Validate() error {
	return dara.Validate(s)
}

type CreateAgentRequestExpertSuites struct {
	// OFFICIAL 表示官方目录资源；DB 表示个人或企业资源。省略时保持原有数据库资源查询语义。官方资源在 Agent 发布时固定制品版本。
	//
	// example:
	//
	// OFFICIAL
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// 传入专家套件引用时，每项需提供 SuiteId，由服务端校验。Source=OFFICIAL 时为官方目录 ID；其他情况为专家套件管理返回的 ID。引用资源必须对调用者可见。
	//
	// example:
	//
	// dataworks-data-analysis-extension
	SuiteId *string `json:"SuiteId,omitempty" xml:"SuiteId,omitempty"`
	// 仅用于个人或企业套件，必须为已发布快照的整数版本号。省略时在 Agent 发布时固定当前已发布版本。Source=OFFICIAL 时不传此字段。
	//
	// example:
	//
	// 1
	Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s CreateAgentRequestExpertSuites) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentRequestExpertSuites) GoString() string {
	return s.String()
}

func (s *CreateAgentRequestExpertSuites) GetSource() *string {
	return s.Source
}

func (s *CreateAgentRequestExpertSuites) GetSuiteId() *string {
	return s.SuiteId
}

func (s *CreateAgentRequestExpertSuites) GetVersion() *int32 {
	return s.Version
}

func (s *CreateAgentRequestExpertSuites) SetSource(v string) *CreateAgentRequestExpertSuites {
	s.Source = &v
	return s
}

func (s *CreateAgentRequestExpertSuites) SetSuiteId(v string) *CreateAgentRequestExpertSuites {
	s.SuiteId = &v
	return s
}

func (s *CreateAgentRequestExpertSuites) SetVersion(v int32) *CreateAgentRequestExpertSuites {
	s.Version = &v
	return s
}

func (s *CreateAgentRequestExpertSuites) Validate() error {
	return dara.Validate(s)
}

type CreateAgentRequestSkills struct {
	// The name of the skill.
	//
	// example:
	//
	// skill-1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Source=OFFICIAL 时传官方目录 ID，不传 Name 或 Version。个人或企业 Skill 不传此字段。
	//
	// example:
	//
	// dataworks-semantic
	SkillId *string `json:"SkillId,omitempty" xml:"SkillId,omitempty"`
	// OFFICIAL 表示官方目录资源；DB 表示个人或企业资源，省略时保持原有按名称查询语义。
	//
	// example:
	//
	// OFFICIAL
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// 仅用于个人或企业 Skill，必须为已发布快照的整数版本。省略时在 Agent 发布时固定当前已发布版本；Source=OFFICIAL 时不传。
	//
	// example:
	//
	// 1
	Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s CreateAgentRequestSkills) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentRequestSkills) GoString() string {
	return s.String()
}

func (s *CreateAgentRequestSkills) GetName() *string {
	return s.Name
}

func (s *CreateAgentRequestSkills) GetSkillId() *string {
	return s.SkillId
}

func (s *CreateAgentRequestSkills) GetSource() *string {
	return s.Source
}

func (s *CreateAgentRequestSkills) GetVersion() *int32 {
	return s.Version
}

func (s *CreateAgentRequestSkills) SetName(v string) *CreateAgentRequestSkills {
	s.Name = &v
	return s
}

func (s *CreateAgentRequestSkills) SetSkillId(v string) *CreateAgentRequestSkills {
	s.SkillId = &v
	return s
}

func (s *CreateAgentRequestSkills) SetSource(v string) *CreateAgentRequestSkills {
	s.Source = &v
	return s
}

func (s *CreateAgentRequestSkills) SetVersion(v int32) *CreateAgentRequestSkills {
	s.Version = &v
	return s
}

func (s *CreateAgentRequestSkills) Validate() error {
	return dara.Validate(s)
}

type CreateAgentRequestTools struct {
	// The name of the McpServer.
	//
	// example:
	//
	// server-1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateAgentRequestTools) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentRequestTools) GoString() string {
	return s.String()
}

func (s *CreateAgentRequestTools) GetName() *string {
	return s.Name
}

func (s *CreateAgentRequestTools) SetName(v string) *CreateAgentRequestTools {
	s.Name = &v
	return s
}

func (s *CreateAgentRequestTools) Validate() error {
	return dara.Validate(s)
}

type CreateAgentRequestVisibilityScope struct {
	// The list of visible project IDs. This parameter takes effect when Visibility is set to `PROJECT`.
	ProjectIds []*string `json:"ProjectIds,omitempty" xml:"ProjectIds,omitempty" type:"Repeated"`
	// The list of visible user IDs. This parameter takes effect when Visibility is set to `USER`.
	UserIds []*string `json:"UserIds,omitempty" xml:"UserIds,omitempty" type:"Repeated"`
}

func (s CreateAgentRequestVisibilityScope) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentRequestVisibilityScope) GoString() string {
	return s.String()
}

func (s *CreateAgentRequestVisibilityScope) GetProjectIds() []*string {
	return s.ProjectIds
}

func (s *CreateAgentRequestVisibilityScope) GetUserIds() []*string {
	return s.UserIds
}

func (s *CreateAgentRequestVisibilityScope) SetProjectIds(v []*string) *CreateAgentRequestVisibilityScope {
	s.ProjectIds = v
	return s
}

func (s *CreateAgentRequestVisibilityScope) SetUserIds(v []*string) *CreateAgentRequestVisibilityScope {
	s.UserIds = v
	return s
}

func (s *CreateAgentRequestVisibilityScope) Validate() error {
	return dara.Validate(s)
}
