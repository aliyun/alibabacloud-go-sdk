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
	// The list of sub-agents that can be called by this agent.
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
	// The published expert suites bound to the agent. For personal or enterprise suites, use SuiteId, and optionally specify a fixed Version or VersionSelector. For official suites, you must explicitly specify Source=OFFICIAL and the SuiteId from the official catalog, and do not specify Version or VersionSelector. If not specified, an empty list is created, and all official suites are not loaded by default.
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
	// Specifies whether to publish the agent immediately after creation. A value of true publishes the created version. A value of false or omitting this parameter maintains the original behavior of creating a draft only. Saving and publishing are not atomic operations. If publishing fails, the draft may already exist. Before retrying, query the agent first to avoid duplicate creation.
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
	// You are a data analytics assistant.
	SystemPrompt *string `json:"SystemPrompt,omitempty" xml:"SystemPrompt,omitempty"`
	// The list of tools.
	//
	// example:
	//
	// -
	Tools []*CreateAgentRequestTools `json:"Tools,omitempty" xml:"Tools,omitempty" type:"Repeated"`
	// The visibility level. Valid values:<br>
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
	// The visibility scope. Specify the corresponding fields based on the Visibility value.
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
	// OFFICIAL indicates an official catalog resource. DB indicates a personal or enterprise resource. If omitted, the original database resource query semantics are maintained. For official resources, the artifact version is fixed when the agent is published.
	//
	// example:
	//
	// OFFICIAL
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// When passing expert suite references, each item must provide a SuiteId, which is validated by the server. When Source is OFFICIAL, this is the official catalog ID. Otherwise, this is the ID returned by expert suite management. The referenced resource must be visible to the caller.
	//
	// example:
	//
	// dataworks-data-analysis-extension
	SuiteId *string `json:"SuiteId,omitempty" xml:"SuiteId,omitempty"`
	// Used only for personal or enterprise suites. Specifies the fixed integer revision number of a published snapshot. This parameter is mutually exclusive with VersionSelector. Publishing a new version does not change the fixed reference. If both are omitted, the latest version is used for each new session. Do not specify this parameter when Source is OFFICIAL.
	//
	// example:
	//
	// 1
	Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
	// Used only for personal or enterprise resources. This parameter is mutually exclusive with Version. Supports exact semantic versions, latest, and full npm/node-semver range syntax, such as 1.2.*, ^1.2.3, ~1.2.3, >=1.0.0 <2.0.0, and 1.x || 2.x. If both are omitted, the latest version is used, and the highest matching published version is selected for each new session. Existing fixed integer references remain unchanged. Do not specify this parameter when Source is OFFICIAL.
	//
	// example:
	//
	// ^1.2.3
	VersionSelector *string `json:"VersionSelector,omitempty" xml:"VersionSelector,omitempty"`
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

func (s *CreateAgentRequestExpertSuites) GetVersionSelector() *string {
	return s.VersionSelector
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

func (s *CreateAgentRequestExpertSuites) SetVersionSelector(v string) *CreateAgentRequestExpertSuites {
	s.VersionSelector = &v
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
	// When Source is OFFICIAL, specify the official catalog ID. Do not specify Name, Version, or VersionSelector. Do not specify this field for personal or enterprise skills.
	//
	// example:
	//
	// dataworks-semantic
	SkillId *string `json:"SkillId,omitempty" xml:"SkillId,omitempty"`
	// OFFICIAL indicates an official catalog resource. DB indicates a personal or enterprise resource. If omitted, the original query-by-name semantics are maintained.
	//
	// example:
	//
	// OFFICIAL
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// Used only for personal or enterprise skills. Specifies the fixed integer revision number of a published snapshot. This parameter is mutually exclusive with VersionSelector. Publishing a new version does not change the fixed reference. If both are omitted, the latest version is used for each new session. Do not specify this parameter when Source is OFFICIAL.
	//
	// example:
	//
	// 1
	Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
	// Used only for personal or enterprise resources. This parameter is mutually exclusive with Version. Supports exact semantic versions, latest, and full npm/node-semver range syntax, such as 1.2.*, ^1.2.3, ~1.2.3, >=1.0.0 <2.0.0, and 1.x || 2.x. If both are omitted, the latest version is used, and the highest matching published version is selected for each new session. Existing fixed integer references remain unchanged. Do not specify this parameter when Source is OFFICIAL.
	//
	// example:
	//
	// ^1.2.3
	VersionSelector *string `json:"VersionSelector,omitempty" xml:"VersionSelector,omitempty"`
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

func (s *CreateAgentRequestSkills) GetVersionSelector() *string {
	return s.VersionSelector
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

func (s *CreateAgentRequestSkills) SetVersionSelector(v string) *CreateAgentRequestSkills {
	s.VersionSelector = &v
	return s
}

func (s *CreateAgentRequestSkills) Validate() error {
	return dara.Validate(s)
}

type CreateAgentRequestTools struct {
	// Required when Kind is BUILTIN_TOOLSET. Specifies the name of the built-in toolset supported by the target server. Omit this parameter when Kind is MCP_TOOLSET.
	//
	// example:
	//
	// web_search
	BuiltinName *string `json:"BuiltinName,omitempty" xml:"BuiltinName,omitempty"`
	// Required for valid tool references. BUILTIN_TOOLSET indicates a built-in toolset, and BuiltinName must be specified. MCP_TOOLSET indicates an MCP toolset, and McpServerName must be specified. The target server validates the parameter based on the tool type.
	//
	// example:
	//
	// MCP_TOOLSET
	Kind *string `json:"Kind,omitempty" xml:"Kind,omitempty"`
	// Optional when Kind is MCP_TOOLSET. If specified, only the listed tools are used. If omitted or left as an empty array, tools are not filtered by name. Each item must be the name of an actual tool in the target MCP service.
	//
	// example:
	//
	// ["only-this-tool"]
	McpItems []*string `json:"McpItems,omitempty" xml:"McpItems,omitempty" type:"Repeated"`
	// Required when Kind is MCP_TOOLSET. Specifies the MCP service name that the current caller has permission to reference and that has an active published version.
	//
	// example:
	//
	// mcp-a
	McpServerName *string `json:"McpServerName,omitempty" xml:"McpServerName,omitempty"`
	// The name of the MCP server.
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

func (s *CreateAgentRequestTools) GetBuiltinName() *string {
	return s.BuiltinName
}

func (s *CreateAgentRequestTools) GetKind() *string {
	return s.Kind
}

func (s *CreateAgentRequestTools) GetMcpItems() []*string {
	return s.McpItems
}

func (s *CreateAgentRequestTools) GetMcpServerName() *string {
	return s.McpServerName
}

func (s *CreateAgentRequestTools) GetName() *string {
	return s.Name
}

func (s *CreateAgentRequestTools) SetBuiltinName(v string) *CreateAgentRequestTools {
	s.BuiltinName = &v
	return s
}

func (s *CreateAgentRequestTools) SetKind(v string) *CreateAgentRequestTools {
	s.Kind = &v
	return s
}

func (s *CreateAgentRequestTools) SetMcpItems(v []*string) *CreateAgentRequestTools {
	s.McpItems = v
	return s
}

func (s *CreateAgentRequestTools) SetMcpServerName(v string) *CreateAgentRequestTools {
	s.McpServerName = &v
	return s
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
