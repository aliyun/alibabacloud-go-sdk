// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateAgentRequest
	GetDescription() *string
	SetDisplayName(v string) *UpdateAgentRequest
	GetDisplayName() *string
	SetExpectedVersion(v int64) *UpdateAgentRequest
	GetExpectedVersion() *int64
	SetKnowledgeBases(v interface{}) *UpdateAgentRequest
	GetKnowledgeBases() interface{}
	SetName(v string) *UpdateAgentRequest
	GetName() *string
	SetSkills(v interface{}) *UpdateAgentRequest
	GetSkills() interface{}
	SetSystemPrompt(v string) *UpdateAgentRequest
	GetSystemPrompt() *string
	SetTools(v interface{}) *UpdateAgentRequest
	GetTools() interface{}
	SetVisibility(v string) *UpdateAgentRequest
	GetVisibility() *string
}

type UpdateAgentRequest struct {
	// The new description. If not specified, the existing value is retained.
	//
	// example:
	//
	// Analyzes CR code changes and checks for correctness, security, and maintainability
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The new display name. If not specified, the existing value is retained.
	//
	// example:
	//
	// CR Code and Security Review Agent
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The expected current configuration revision number. This parameter is omitted by default. For concurrency protection, pass in the `AgentVersion` returned by the most recent `GetAgent` call.
	//
	// example:
	//
	// 2
	ExpectedVersion *int64 `json:"ExpectedVersion,omitempty" xml:"ExpectedVersion,omitempty"`
	// The list of knowledge base bindings, which contains at most one element. If not specified, the existing value is retained. A non-empty array replaces the entire value. Passing `[ ]` removes all bindings.
	//
	// example:
	//
	// [{"name":"code-review-guidelines"}]
	KnowledgeBases interface{} `json:"KnowledgeBases,omitempty" xml:"KnowledgeBases,omitempty"`
	// The name of the Agent to update.
	//
	// This parameter is required.
	//
	// example:
	//
	// code-review-agent
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The list of Skill references. If not specified, the existing value is retained. A non-empty array replaces the entire value. Passing `[ ]` removes all bindings.
	//
	// example:
	//
	// [{"name":"code-review"}]
	Skills interface{} `json:"Skills,omitempty" xml:"Skills,omitempty"`
	// The new system prompt. If not specified, the existing value is retained.
	//
	// example:
	//
	// Check the CR for code defects, security risks, and compatibility issues, and provide actionable suggestions for fixes
	SystemPrompt *string `json:"SystemPrompt,omitempty" xml:"SystemPrompt,omitempty"`
	// The list of MCP Server and Connector name references. The same array supports both types of entries. Each entry specifies one type of reference. If items is omitted for an MCP entry, all public tools are included. Previously specified items retain their existing values.
	//
	// example:
	//
	// [{"mcpServerName":"code-repository-mcp"},{"connectorName":"code-review-data"}]
	Tools interface{} `json:"Tools,omitempty" xml:"Tools,omitempty"`
	// The new visibility scope. Valid values:
	//
	// - user
	//
	// - tenant
	//
	// If not specified, the existing value is retained.
	//
	// example:
	//
	// tenant
	Visibility *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s UpdateAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentRequest) GoString() string {
	return s.String()
}

func (s *UpdateAgentRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateAgentRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *UpdateAgentRequest) GetExpectedVersion() *int64 {
	return s.ExpectedVersion
}

func (s *UpdateAgentRequest) GetKnowledgeBases() interface{} {
	return s.KnowledgeBases
}

func (s *UpdateAgentRequest) GetName() *string {
	return s.Name
}

func (s *UpdateAgentRequest) GetSkills() interface{} {
	return s.Skills
}

func (s *UpdateAgentRequest) GetSystemPrompt() *string {
	return s.SystemPrompt
}

func (s *UpdateAgentRequest) GetTools() interface{} {
	return s.Tools
}

func (s *UpdateAgentRequest) GetVisibility() *string {
	return s.Visibility
}

func (s *UpdateAgentRequest) SetDescription(v string) *UpdateAgentRequest {
	s.Description = &v
	return s
}

func (s *UpdateAgentRequest) SetDisplayName(v string) *UpdateAgentRequest {
	s.DisplayName = &v
	return s
}

func (s *UpdateAgentRequest) SetExpectedVersion(v int64) *UpdateAgentRequest {
	s.ExpectedVersion = &v
	return s
}

func (s *UpdateAgentRequest) SetKnowledgeBases(v interface{}) *UpdateAgentRequest {
	s.KnowledgeBases = v
	return s
}

func (s *UpdateAgentRequest) SetName(v string) *UpdateAgentRequest {
	s.Name = &v
	return s
}

func (s *UpdateAgentRequest) SetSkills(v interface{}) *UpdateAgentRequest {
	s.Skills = v
	return s
}

func (s *UpdateAgentRequest) SetSystemPrompt(v string) *UpdateAgentRequest {
	s.SystemPrompt = &v
	return s
}

func (s *UpdateAgentRequest) SetTools(v interface{}) *UpdateAgentRequest {
	s.Tools = v
	return s
}

func (s *UpdateAgentRequest) SetVisibility(v string) *UpdateAgentRequest {
	s.Visibility = &v
	return s
}

func (s *UpdateAgentRequest) Validate() error {
	return dara.Validate(s)
}
