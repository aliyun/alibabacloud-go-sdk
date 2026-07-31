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
	SetMetadata(v map[string]interface{}) *CreateAgentRequest
	GetMetadata() map[string]interface{}
	SetModel(v map[string]interface{}) *CreateAgentRequest
	GetModel() map[string]interface{}
	SetName(v string) *CreateAgentRequest
	GetName() *string
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
	// The list of child Agents that can be called by this Agent.
	//
	// example:
	//
	// -
	CallableAgents []*CreateAgentRequestCallableAgents `json:"CallableAgents,omitempty" xml:"CallableAgents,omitempty" type:"Repeated"`
	// The description of the Agent.
	//
	// example:
	//
	// Data analytics assistant
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The display name of the Agent.
	//
	// example:
	//
	// MyAssistant.
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The extended metadata (key-value pairs).
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
	// The Agent name, which must be unique within the current account.
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
	Skills []*CreateAgentRequestSkills `json:"Skills,omitempty" xml:"Skills,omitempty" type:"Repeated"`
	// The system prompt.
	//
	// example:
	//
	// You are a data analytics assistant.
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
	// The visibility scope. The corresponding field is determined by the Visibility parameter.
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

func (s *CreateAgentRequest) GetMetadata() map[string]interface{} {
	return s.Metadata
}

func (s *CreateAgentRequest) GetModel() map[string]interface{} {
	return s.Model
}

func (s *CreateAgentRequest) GetName() *string {
	return s.Name
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
	// The Agent name.
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

type CreateAgentRequestSkills struct {
	// The skill name.
	//
	// example:
	//
	// skill-1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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

func (s *CreateAgentRequestSkills) SetName(v string) *CreateAgentRequestSkills {
	s.Name = &v
	return s
}

func (s *CreateAgentRequestSkills) Validate() error {
	return dara.Validate(s)
}

type CreateAgentRequestTools struct {
	// The McpServer name.
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
	// The list of project IDs that have visibility. This parameter takes effect when Visibility is set to `PROJECT`.
	ProjectIds []*string `json:"ProjectIds,omitempty" xml:"ProjectIds,omitempty" type:"Repeated"`
	// The list of user IDs that have visibility. This parameter takes effect when Visibility is set to `USER`.
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
