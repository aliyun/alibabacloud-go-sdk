// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallableAgents(v []*UpdateAgentRequestCallableAgents) *UpdateAgentRequest
	GetCallableAgents() []*UpdateAgentRequestCallableAgents
	SetDescription(v string) *UpdateAgentRequest
	GetDescription() *string
	SetDisplayName(v string) *UpdateAgentRequest
	GetDisplayName() *string
	SetExpectedVersion(v int32) *UpdateAgentRequest
	GetExpectedVersion() *int32
	SetExpertSuites(v []*UpdateAgentRequestExpertSuites) *UpdateAgentRequest
	GetExpertSuites() []*UpdateAgentRequestExpertSuites
	SetMetadata(v map[string]interface{}) *UpdateAgentRequest
	GetMetadata() map[string]interface{}
	SetModel(v map[string]interface{}) *UpdateAgentRequest
	GetModel() map[string]interface{}
	SetName(v string) *UpdateAgentRequest
	GetName() *string
	SetPublishImmediately(v bool) *UpdateAgentRequest
	GetPublishImmediately() *bool
	SetSkills(v []*UpdateAgentRequestSkills) *UpdateAgentRequest
	GetSkills() []*UpdateAgentRequestSkills
	SetSystemPrompt(v string) *UpdateAgentRequest
	GetSystemPrompt() *string
	SetTools(v []*UpdateAgentRequestTools) *UpdateAgentRequest
	GetTools() []*UpdateAgentRequestTools
	SetVisibility(v string) *UpdateAgentRequest
	GetVisibility() *string
	SetVisibilityScope(v *UpdateAgentRequestVisibilityScope) *UpdateAgentRequest
	GetVisibilityScope() *UpdateAgentRequestVisibilityScope
}

type UpdateAgentRequest struct {
	// The list of callable sub-agent references. For non-empty references, the referenced sub-agents must be visible. The agent cannot reference itself or form a call loop. Omit this parameter if no modification is required. Do not use an empty array to express general retention or clearing semantics.
	//
	// example:
	//
	// [{"Name":"data-analysis"}]
	CallableAgents []*UpdateAgentRequestCallableAgents `json:"CallableAgents,omitempty" xml:"CallableAgents,omitempty" type:"Repeated"`
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
	// The agent definition version used for concurrency control. We recommend explicitly specifying the AgentVersion you just obtained from GetAgent for each call. The update is rejected if the version does not match the current version on the server. To ensure compatibility with different deployment versions, do not rely on the server-side processing when this parameter is omitted.
	//
	// example:
	//
	// 2
	ExpectedVersion *int32 `json:"ExpectedVersion,omitempty" xml:"ExpectedVersion,omitempty"`
	// The list of expert suite references. If omitted, the existing references are retained. An explicit empty array clears the references. For official suites, specify only Source=OFFICIAL and SuiteId. Version, VersionSelector, or client artifact fields are not accepted. For custom suites, use SuiteId, and you can explicitly specify an existing integer Version.
	//
	// example:
	//
	// [{"SuiteId":"suite_example","Version":3}]
	ExpertSuites []*UpdateAgentRequestExpertSuites `json:"ExpertSuites,omitempty" xml:"ExpertSuites,omitempty" type:"Repeated"`
	// The extended metadata. If omitted, the existing metadata is retained. If provided, the business metadata is updated, and the server retains the protected system metadata.
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
	// The name of the agent to be updated. Use the Name returned by GetAgent. This field is used to locate the target agent and is not used for renaming.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-agent
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Specifies whether to publish the agent immediately after the update. If this parameter is set to true, the operation saves a new draft version before requesting publication. If this parameter is set to false or omitted, it saves only the definition. Saving and publishing are not atomic operations. If the operation fails, call GetAgent to read the current version before retrying to avoid blindly repeating the update.
	//
	// example:
	//
	// false
	PublishImmediately *bool `json:"PublishImmediately,omitempty" xml:"PublishImmediately,omitempty"`
	// The list of skill references. Omit this parameter if no modification is required. Do not use an empty array to indicate retention or clearing. For official skills, specify only Source=OFFICIAL and SkillId. Name, Version, VersionSelector, or client artifact fields are not accepted. For custom skills, specify Name, and you can explicitly specify an existing integer Version.
	//
	// example:
	//
	// [{"Source":"OFFICIAL","SkillId":"dataworks-semantic"}]
	Skills []*UpdateAgentRequestSkills `json:"Skills,omitempty" xml:"Skills,omitempty" type:"Repeated"`
	// The system prompt.
	//
	// example:
	//
	// You are an excellent ETL engineer
	SystemPrompt *string `json:"SystemPrompt,omitempty" xml:"SystemPrompt,omitempty"`
	// The list of tool references. Omit this parameter if no modification is required. The current new backend retains existing tools even for an empty array. A non-empty array replaces the references. Each item must include Kind and the corresponding configuration, and cannot include only Name. Replace the MCP service and tool names in the example with actual values available to the current identity.
	//
	// example:
	//
	// [{"Kind":"MCP_TOOLSET","McpServerName":"mcp-a","McpItems":["only-this-tool"]}]
	Tools []*UpdateAgentRequestTools `json:"Tools,omitempty" xml:"Tools,omitempty" type:"Repeated"`
	// The visibility level.<br>TENANT: visible within the account.<br>PROJECT: visible to specified projects.<br>USER: visible to specified users.
	//
	// example:
	//
	// TENANT
	Visibility *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
	// The visibility scope. The corresponding field is used based on the Visibility value.
	VisibilityScope *UpdateAgentRequestVisibilityScope `json:"VisibilityScope,omitempty" xml:"VisibilityScope,omitempty" type:"Struct"`
}

func (s UpdateAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentRequest) GoString() string {
	return s.String()
}

func (s *UpdateAgentRequest) GetCallableAgents() []*UpdateAgentRequestCallableAgents {
	return s.CallableAgents
}

func (s *UpdateAgentRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateAgentRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *UpdateAgentRequest) GetExpectedVersion() *int32 {
	return s.ExpectedVersion
}

func (s *UpdateAgentRequest) GetExpertSuites() []*UpdateAgentRequestExpertSuites {
	return s.ExpertSuites
}

func (s *UpdateAgentRequest) GetMetadata() map[string]interface{} {
	return s.Metadata
}

func (s *UpdateAgentRequest) GetModel() map[string]interface{} {
	return s.Model
}

func (s *UpdateAgentRequest) GetName() *string {
	return s.Name
}

func (s *UpdateAgentRequest) GetPublishImmediately() *bool {
	return s.PublishImmediately
}

func (s *UpdateAgentRequest) GetSkills() []*UpdateAgentRequestSkills {
	return s.Skills
}

func (s *UpdateAgentRequest) GetSystemPrompt() *string {
	return s.SystemPrompt
}

func (s *UpdateAgentRequest) GetTools() []*UpdateAgentRequestTools {
	return s.Tools
}

func (s *UpdateAgentRequest) GetVisibility() *string {
	return s.Visibility
}

func (s *UpdateAgentRequest) GetVisibilityScope() *UpdateAgentRequestVisibilityScope {
	return s.VisibilityScope
}

func (s *UpdateAgentRequest) SetCallableAgents(v []*UpdateAgentRequestCallableAgents) *UpdateAgentRequest {
	s.CallableAgents = v
	return s
}

func (s *UpdateAgentRequest) SetDescription(v string) *UpdateAgentRequest {
	s.Description = &v
	return s
}

func (s *UpdateAgentRequest) SetDisplayName(v string) *UpdateAgentRequest {
	s.DisplayName = &v
	return s
}

func (s *UpdateAgentRequest) SetExpectedVersion(v int32) *UpdateAgentRequest {
	s.ExpectedVersion = &v
	return s
}

func (s *UpdateAgentRequest) SetExpertSuites(v []*UpdateAgentRequestExpertSuites) *UpdateAgentRequest {
	s.ExpertSuites = v
	return s
}

func (s *UpdateAgentRequest) SetMetadata(v map[string]interface{}) *UpdateAgentRequest {
	s.Metadata = v
	return s
}

func (s *UpdateAgentRequest) SetModel(v map[string]interface{}) *UpdateAgentRequest {
	s.Model = v
	return s
}

func (s *UpdateAgentRequest) SetName(v string) *UpdateAgentRequest {
	s.Name = &v
	return s
}

func (s *UpdateAgentRequest) SetPublishImmediately(v bool) *UpdateAgentRequest {
	s.PublishImmediately = &v
	return s
}

func (s *UpdateAgentRequest) SetSkills(v []*UpdateAgentRequestSkills) *UpdateAgentRequest {
	s.Skills = v
	return s
}

func (s *UpdateAgentRequest) SetSystemPrompt(v string) *UpdateAgentRequest {
	s.SystemPrompt = &v
	return s
}

func (s *UpdateAgentRequest) SetTools(v []*UpdateAgentRequestTools) *UpdateAgentRequest {
	s.Tools = v
	return s
}

func (s *UpdateAgentRequest) SetVisibility(v string) *UpdateAgentRequest {
	s.Visibility = &v
	return s
}

func (s *UpdateAgentRequest) SetVisibilityScope(v *UpdateAgentRequestVisibilityScope) *UpdateAgentRequest {
	s.VisibilityScope = v
	return s
}

func (s *UpdateAgentRequest) Validate() error {
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

type UpdateAgentRequestCallableAgents struct {
	// The name of the called sub-agent. The agent must be one that the current caller has permissions to use.
	//
	// example:
	//
	// data-analysis
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateAgentRequestCallableAgents) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentRequestCallableAgents) GoString() string {
	return s.String()
}

func (s *UpdateAgentRequestCallableAgents) GetName() *string {
	return s.Name
}

func (s *UpdateAgentRequestCallableAgents) SetName(v string) *UpdateAgentRequestCallableAgents {
	s.Name = &v
	return s
}

func (s *UpdateAgentRequestCallableAgents) Validate() error {
	return dara.Validate(s)
}

type UpdateAgentRequestExpertSuites struct {
	// The source of the expert suite. OFFICIAL: official suite. DB or omitted: custom suite. Both types of references use SuiteId, and the version fields apply only to custom resources.
	//
	// example:
	//
	// OFFICIAL
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The ID of the target expert suite. The suite must be one that the current caller has permissions to reference.
	//
	// example:
	//
	// suite_example
	SuiteId *string `json:"SuiteId,omitempty" xml:"SuiteId,omitempty"`
	// The published integer version number of the custom expert suite, used to explicitly select an existing version. Do not pass this field when Source is set to OFFICIAL. This parameter is mutually exclusive with VersionSelector. To obtain explicit and reproducible references, do not rely on the default selection when the version is omitted.
	//
	// example:
	//
	// 3
	Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
	// The semantic version selection expression of the custom expert suite, such as ^1.2.3. Use this parameter only when the target server has deployed the custom resource version selection capability. This parameter is mutually exclusive with Version. This field is not accepted when Source is set to OFFICIAL. To ensure compatibility with servers of unconfirmed versions, use an existing integer Version.
	//
	// example:
	//
	// ^1.2.3
	VersionSelector *string `json:"VersionSelector,omitempty" xml:"VersionSelector,omitempty"`
}

func (s UpdateAgentRequestExpertSuites) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentRequestExpertSuites) GoString() string {
	return s.String()
}

func (s *UpdateAgentRequestExpertSuites) GetSource() *string {
	return s.Source
}

func (s *UpdateAgentRequestExpertSuites) GetSuiteId() *string {
	return s.SuiteId
}

func (s *UpdateAgentRequestExpertSuites) GetVersion() *int32 {
	return s.Version
}

func (s *UpdateAgentRequestExpertSuites) GetVersionSelector() *string {
	return s.VersionSelector
}

func (s *UpdateAgentRequestExpertSuites) SetSource(v string) *UpdateAgentRequestExpertSuites {
	s.Source = &v
	return s
}

func (s *UpdateAgentRequestExpertSuites) SetSuiteId(v string) *UpdateAgentRequestExpertSuites {
	s.SuiteId = &v
	return s
}

func (s *UpdateAgentRequestExpertSuites) SetVersion(v int32) *UpdateAgentRequestExpertSuites {
	s.Version = &v
	return s
}

func (s *UpdateAgentRequestExpertSuites) SetVersionSelector(v string) *UpdateAgentRequestExpertSuites {
	s.VersionSelector = &v
	return s
}

func (s *UpdateAgentRequestExpertSuites) Validate() error {
	return dara.Validate(s)
}

type UpdateAgentRequestSkills struct {
	// The name of the custom skill. Use the Name of a visible custom skill. For official skills, use SkillId.
	//
	// example:
	//
	// my-business-skill
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The resource ID of the official skill. This parameter is used when Source is set to OFFICIAL.
	//
	// example:
	//
	// dataworks-semantic
	SkillId *string `json:"SkillId,omitempty" xml:"SkillId,omitempty"`
	// The source of the skill. OFFICIAL: official skill, used with SkillId. DB or omitted: custom skill, used with Name.
	//
	// example:
	//
	// OFFICIAL
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The published integer version number of the custom skill, used to explicitly select an existing version. Do not pass this field when Source is set to OFFICIAL. This parameter is mutually exclusive with VersionSelector.
	//
	// example:
	//
	// 1
	Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
	// The semantic version selection expression of the custom skill, such as ^1.2.3. Use this parameter only when the target server has deployed the custom resource version selection capability. This parameter is mutually exclusive with Version. This field is not accepted when Source is set to OFFICIAL. To ensure compatibility with servers of unconfirmed versions, use an existing integer Version.
	//
	// example:
	//
	// ^1.2.3
	VersionSelector *string `json:"VersionSelector,omitempty" xml:"VersionSelector,omitempty"`
}

func (s UpdateAgentRequestSkills) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentRequestSkills) GoString() string {
	return s.String()
}

func (s *UpdateAgentRequestSkills) GetName() *string {
	return s.Name
}

func (s *UpdateAgentRequestSkills) GetSkillId() *string {
	return s.SkillId
}

func (s *UpdateAgentRequestSkills) GetSource() *string {
	return s.Source
}

func (s *UpdateAgentRequestSkills) GetVersion() *int32 {
	return s.Version
}

func (s *UpdateAgentRequestSkills) GetVersionSelector() *string {
	return s.VersionSelector
}

func (s *UpdateAgentRequestSkills) SetName(v string) *UpdateAgentRequestSkills {
	s.Name = &v
	return s
}

func (s *UpdateAgentRequestSkills) SetSkillId(v string) *UpdateAgentRequestSkills {
	s.SkillId = &v
	return s
}

func (s *UpdateAgentRequestSkills) SetSource(v string) *UpdateAgentRequestSkills {
	s.Source = &v
	return s
}

func (s *UpdateAgentRequestSkills) SetVersion(v int32) *UpdateAgentRequestSkills {
	s.Version = &v
	return s
}

func (s *UpdateAgentRequestSkills) SetVersionSelector(v string) *UpdateAgentRequestSkills {
	s.VersionSelector = &v
	return s
}

func (s *UpdateAgentRequestSkills) Validate() error {
	return dara.Validate(s)
}

type UpdateAgentRequestTools struct {
	// The name of the built-in toolset. This parameter is required when Kind is set to BUILTIN_TOOLSET. Use a name supported by the target server. Omit this parameter when Kind is set to MCP_TOOLSET.
	//
	// example:
	//
	// web_search
	BuiltinName *string `json:"BuiltinName,omitempty" xml:"BuiltinName,omitempty"`
	// The tool type. A valid tool reference must specify BUILTIN_TOOLSET or MCP_TOOLSET. For BUILTIN_TOOLSET, BuiltinName must also be specified. For MCP_TOOLSET, McpServerName must also be specified. The target server validates the reference based on the tool type.
	//
	// example:
	//
	// MCP_TOOLSET
	Kind *string `json:"Kind,omitempty" xml:"Kind,omitempty"`
	// The list of selected MCP tool names. If this parameter is omitted or an empty array is provided, no filtering by tool name is performed. Each item must be an actual tool name in the target service.
	//
	// example:
	//
	// ["only-this-tool"]
	McpItems []*string `json:"McpItems,omitempty" xml:"McpItems,omitempty" type:"Repeated"`
	// The MCP service name. This parameter is required when Kind is set to MCP_TOOLSET. The service must be one that the current invoker has permissions to reference and that has an active published version.
	//
	// example:
	//
	// mcp-a
	McpServerName *string `json:"McpServerName,omitempty" xml:"McpServerName,omitempty"`
	// A field retained from the legacy model. Specifying only Name does not constitute a valid tool configuration and cannot replace Kind, BuiltinName, or McpServerName.
	//
	// example:
	//
	// search_tables
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateAgentRequestTools) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentRequestTools) GoString() string {
	return s.String()
}

func (s *UpdateAgentRequestTools) GetBuiltinName() *string {
	return s.BuiltinName
}

func (s *UpdateAgentRequestTools) GetKind() *string {
	return s.Kind
}

func (s *UpdateAgentRequestTools) GetMcpItems() []*string {
	return s.McpItems
}

func (s *UpdateAgentRequestTools) GetMcpServerName() *string {
	return s.McpServerName
}

func (s *UpdateAgentRequestTools) GetName() *string {
	return s.Name
}

func (s *UpdateAgentRequestTools) SetBuiltinName(v string) *UpdateAgentRequestTools {
	s.BuiltinName = &v
	return s
}

func (s *UpdateAgentRequestTools) SetKind(v string) *UpdateAgentRequestTools {
	s.Kind = &v
	return s
}

func (s *UpdateAgentRequestTools) SetMcpItems(v []*string) *UpdateAgentRequestTools {
	s.McpItems = v
	return s
}

func (s *UpdateAgentRequestTools) SetMcpServerName(v string) *UpdateAgentRequestTools {
	s.McpServerName = &v
	return s
}

func (s *UpdateAgentRequestTools) SetName(v string) *UpdateAgentRequestTools {
	s.Name = &v
	return s
}

func (s *UpdateAgentRequestTools) Validate() error {
	return dara.Validate(s)
}

type UpdateAgentRequestVisibilityScope struct {
	// The list of visible project IDs. This parameter takes effect when Visibility is set to PROJECT.
	ProjectIds []*string `json:"ProjectIds,omitempty" xml:"ProjectIds,omitempty" type:"Repeated"`
	// The list of visible user IDs. This parameter is used when Visibility is set to USER.
	UserIds []*string `json:"UserIds,omitempty" xml:"UserIds,omitempty" type:"Repeated"`
}

func (s UpdateAgentRequestVisibilityScope) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentRequestVisibilityScope) GoString() string {
	return s.String()
}

func (s *UpdateAgentRequestVisibilityScope) GetProjectIds() []*string {
	return s.ProjectIds
}

func (s *UpdateAgentRequestVisibilityScope) GetUserIds() []*string {
	return s.UserIds
}

func (s *UpdateAgentRequestVisibilityScope) SetProjectIds(v []*string) *UpdateAgentRequestVisibilityScope {
	s.ProjectIds = v
	return s
}

func (s *UpdateAgentRequestVisibilityScope) SetUserIds(v []*string) *UpdateAgentRequestVisibilityScope {
	s.UserIds = v
	return s
}

func (s *UpdateAgentRequestVisibilityScope) Validate() error {
	return dara.Validate(s)
}
