// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAgentShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallableAgentsShrink(v string) *UpdateAgentShrinkRequest
	GetCallableAgentsShrink() *string
	SetDescription(v string) *UpdateAgentShrinkRequest
	GetDescription() *string
	SetDisplayName(v string) *UpdateAgentShrinkRequest
	GetDisplayName() *string
	SetExpectedVersion(v int32) *UpdateAgentShrinkRequest
	GetExpectedVersion() *int32
	SetExpertSuitesShrink(v string) *UpdateAgentShrinkRequest
	GetExpertSuitesShrink() *string
	SetMetadataShrink(v string) *UpdateAgentShrinkRequest
	GetMetadataShrink() *string
	SetModelShrink(v string) *UpdateAgentShrinkRequest
	GetModelShrink() *string
	SetName(v string) *UpdateAgentShrinkRequest
	GetName() *string
	SetPublishImmediately(v bool) *UpdateAgentShrinkRequest
	GetPublishImmediately() *bool
	SetSkillsShrink(v string) *UpdateAgentShrinkRequest
	GetSkillsShrink() *string
	SetSystemPrompt(v string) *UpdateAgentShrinkRequest
	GetSystemPrompt() *string
	SetToolsShrink(v string) *UpdateAgentShrinkRequest
	GetToolsShrink() *string
	SetVisibility(v string) *UpdateAgentShrinkRequest
	GetVisibility() *string
	SetVisibilityScopeShrink(v string) *UpdateAgentShrinkRequest
	GetVisibilityScopeShrink() *string
}

type UpdateAgentShrinkRequest struct {
	// The list of callable sub-agent references. For non-empty references, the referenced sub-agents must be visible. The agent cannot reference itself or form a call loop. Omit this parameter if no modification is required. Do not use an empty array to express general retention or clearing semantics.
	//
	// example:
	//
	// [{"Name":"data-analysis"}]
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
	// The agent definition version used for concurrency control. We recommend that you explicitly specify the AgentVersion just obtained from GetAgent for each call. The update is rejected if the version does not match the current version on the server. To ensure compatibility with different deployment versions, do not rely on the server-side processing when this parameter is omitted.
	//
	// example:
	//
	// 2
	ExpectedVersion *int32 `json:"ExpectedVersion,omitempty" xml:"ExpectedVersion,omitempty"`
	// The list of expert suite references. If omitted, existing references are retained. An explicit empty array clears the references. For official suites, specify only Source=OFFICIAL and SuiteId. Version, VersionSelector, or client artifact fields are not accepted. For custom suites, use SuiteId, and you can explicitly specify an existing integer Version.
	//
	// example:
	//
	// [{"SuiteId":"suite_example","Version":3}]
	ExpertSuitesShrink *string `json:"ExpertSuites,omitempty" xml:"ExpertSuites,omitempty"`
	// The extended metadata. If omitted, existing metadata is retained. If passed, business metadata is updated, and the server retains protected system metadata.
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
	// The name of the agent to be updated. Use the Name returned by GetAgent. This field is used to locate the target agent and cannot be used to rename the agent.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-agent
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Specifies whether to publish the agent immediately after the update. Valid values: true: saves a new draft version first and then requests publishing. false or omitted: saves only the definition. Saving and publishing are not atomic operations. If the operation fails, call GetAgent to read the current version back before retrying, to avoid blindly repeating the update.
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
	SkillsShrink *string `json:"Skills,omitempty" xml:"Skills,omitempty"`
	// The system prompt.
	//
	// example:
	//
	// You are an excellent ETL engineer
	SystemPrompt *string `json:"SystemPrompt,omitempty" xml:"SystemPrompt,omitempty"`
	// The list of tool references. Omit this parameter if no modification is required. The current new backend retains existing tools even if an empty array is passed. A non-empty array replaces the references. Each item must include Kind and the corresponding configuration, and cannot include only Name. Replace the MCP service and tool names in the example with actual values available to the current identity.
	//
	// example:
	//
	// [{"Kind":"MCP_TOOLSET","McpServerName":"mcp-a","McpItems":["only-this-tool"]}]
	ToolsShrink *string `json:"Tools,omitempty" xml:"Tools,omitempty"`
	// The visibility level. Valid values: <br>`TENANT`: visible within the account. <br>`PROJECT`: visible to specified projects. <br>`USER`: visible to specified users.
	//
	// example:
	//
	// TENANT
	Visibility *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
	// The visibility scope. The corresponding field is used based on the Visibility parameter.
	VisibilityScopeShrink *string `json:"VisibilityScope,omitempty" xml:"VisibilityScope,omitempty"`
}

func (s UpdateAgentShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateAgentShrinkRequest) GetCallableAgentsShrink() *string {
	return s.CallableAgentsShrink
}

func (s *UpdateAgentShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateAgentShrinkRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *UpdateAgentShrinkRequest) GetExpectedVersion() *int32 {
	return s.ExpectedVersion
}

func (s *UpdateAgentShrinkRequest) GetExpertSuitesShrink() *string {
	return s.ExpertSuitesShrink
}

func (s *UpdateAgentShrinkRequest) GetMetadataShrink() *string {
	return s.MetadataShrink
}

func (s *UpdateAgentShrinkRequest) GetModelShrink() *string {
	return s.ModelShrink
}

func (s *UpdateAgentShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdateAgentShrinkRequest) GetPublishImmediately() *bool {
	return s.PublishImmediately
}

func (s *UpdateAgentShrinkRequest) GetSkillsShrink() *string {
	return s.SkillsShrink
}

func (s *UpdateAgentShrinkRequest) GetSystemPrompt() *string {
	return s.SystemPrompt
}

func (s *UpdateAgentShrinkRequest) GetToolsShrink() *string {
	return s.ToolsShrink
}

func (s *UpdateAgentShrinkRequest) GetVisibility() *string {
	return s.Visibility
}

func (s *UpdateAgentShrinkRequest) GetVisibilityScopeShrink() *string {
	return s.VisibilityScopeShrink
}

func (s *UpdateAgentShrinkRequest) SetCallableAgentsShrink(v string) *UpdateAgentShrinkRequest {
	s.CallableAgentsShrink = &v
	return s
}

func (s *UpdateAgentShrinkRequest) SetDescription(v string) *UpdateAgentShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateAgentShrinkRequest) SetDisplayName(v string) *UpdateAgentShrinkRequest {
	s.DisplayName = &v
	return s
}

func (s *UpdateAgentShrinkRequest) SetExpectedVersion(v int32) *UpdateAgentShrinkRequest {
	s.ExpectedVersion = &v
	return s
}

func (s *UpdateAgentShrinkRequest) SetExpertSuitesShrink(v string) *UpdateAgentShrinkRequest {
	s.ExpertSuitesShrink = &v
	return s
}

func (s *UpdateAgentShrinkRequest) SetMetadataShrink(v string) *UpdateAgentShrinkRequest {
	s.MetadataShrink = &v
	return s
}

func (s *UpdateAgentShrinkRequest) SetModelShrink(v string) *UpdateAgentShrinkRequest {
	s.ModelShrink = &v
	return s
}

func (s *UpdateAgentShrinkRequest) SetName(v string) *UpdateAgentShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateAgentShrinkRequest) SetPublishImmediately(v bool) *UpdateAgentShrinkRequest {
	s.PublishImmediately = &v
	return s
}

func (s *UpdateAgentShrinkRequest) SetSkillsShrink(v string) *UpdateAgentShrinkRequest {
	s.SkillsShrink = &v
	return s
}

func (s *UpdateAgentShrinkRequest) SetSystemPrompt(v string) *UpdateAgentShrinkRequest {
	s.SystemPrompt = &v
	return s
}

func (s *UpdateAgentShrinkRequest) SetToolsShrink(v string) *UpdateAgentShrinkRequest {
	s.ToolsShrink = &v
	return s
}

func (s *UpdateAgentShrinkRequest) SetVisibility(v string) *UpdateAgentShrinkRequest {
	s.Visibility = &v
	return s
}

func (s *UpdateAgentShrinkRequest) SetVisibilityScopeShrink(v string) *UpdateAgentShrinkRequest {
	s.VisibilityScopeShrink = &v
	return s
}

func (s *UpdateAgentShrinkRequest) Validate() error {
	return dara.Validate(s)
}
