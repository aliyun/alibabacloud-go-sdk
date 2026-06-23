// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallableAgents(v []*string) *CreateAgentRequest
	GetCallableAgents() []*string
	SetDescription(v string) *CreateAgentRequest
	GetDescription() *string
	SetDisplayName(v string) *CreateAgentRequest
	GetDisplayName() *string
	SetMetadata(v map[string]interface{}) *CreateAgentRequest
	GetMetadata() map[string]interface{}
	SetModel(v map[string]interface{}) *CreateAgentRequest
	GetModel() map[string]interface{}
	SetName(v string) *CreateAgentRequest
	GetName() *string
	SetSkills(v []*string) *CreateAgentRequest
	GetSkills() []*string
	SetSystemPrompt(v string) *CreateAgentRequest
	GetSystemPrompt() *string
	SetTools(v []*string) *CreateAgentRequest
	GetTools() []*string
	SetVisibility(v string) *CreateAgentRequest
	GetVisibility() *string
	SetVisibilityScope(v *CreateAgentRequestVisibilityScope) *CreateAgentRequest
	GetVisibilityScope() *CreateAgentRequestVisibilityScope
}

type CreateAgentRequest struct {
	// The list of sub-Agents that can be called by this Agent.
	//
	// example:
	//
	// -
	CallableAgents []*string `json:"CallableAgents,omitempty" xml:"CallableAgents,omitempty" type:"Repeated"`
	// The description of the Agent.
	//
	// example:
	//
	// 数据分析助手
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The display name of the Agent.
	//
	// example:
	//
	// 我的助手
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// Extended metadata (key-value pairs).
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
	// The name of the Agent. It must be unique under the current account.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-agent
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The list of skills.
	//
	// example:
	//
	// -
	Skills []*string `json:"Skills,omitempty" xml:"Skills,omitempty" type:"Repeated"`
	// The system prompt.
	//
	// example:
	//
	// 你是一个数据分析助手。
	SystemPrompt *string `json:"SystemPrompt,omitempty" xml:"SystemPrompt,omitempty"`
	// The list of tools.
	//
	// example:
	//
	// -
	Tools []*string `json:"Tools,omitempty" xml:"Tools,omitempty" type:"Repeated"`
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
	// The visibility scope. The corresponding field is selected based on Visibility.
	VisibilityScope *CreateAgentRequestVisibilityScope `json:"VisibilityScope,omitempty" xml:"VisibilityScope,omitempty" type:"Struct"`
}

func (s CreateAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentRequest) GoString() string {
	return s.String()
}

func (s *CreateAgentRequest) GetCallableAgents() []*string {
	return s.CallableAgents
}

func (s *CreateAgentRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateAgentRequest) GetDisplayName() *string {
	return s.DisplayName
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

func (s *CreateAgentRequest) GetSkills() []*string {
	return s.Skills
}

func (s *CreateAgentRequest) GetSystemPrompt() *string {
	return s.SystemPrompt
}

func (s *CreateAgentRequest) GetTools() []*string {
	return s.Tools
}

func (s *CreateAgentRequest) GetVisibility() *string {
	return s.Visibility
}

func (s *CreateAgentRequest) GetVisibilityScope() *CreateAgentRequestVisibilityScope {
	return s.VisibilityScope
}

func (s *CreateAgentRequest) SetCallableAgents(v []*string) *CreateAgentRequest {
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

func (s *CreateAgentRequest) SetSkills(v []*string) *CreateAgentRequest {
	s.Skills = v
	return s
}

func (s *CreateAgentRequest) SetSystemPrompt(v string) *CreateAgentRequest {
	s.SystemPrompt = &v
	return s
}

func (s *CreateAgentRequest) SetTools(v []*string) *CreateAgentRequest {
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
	if s.VisibilityScope != nil {
		if err := s.VisibilityScope.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAgentRequestVisibilityScope struct {
	// The list of visible project IDs. Takes effect when Visibility is `PROJECT`.
	ProjectIds []*string `json:"ProjectIds,omitempty" xml:"ProjectIds,omitempty" type:"Repeated"`
	// The list of visible user IDs. Takes effect when Visibility is `USER`.
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
