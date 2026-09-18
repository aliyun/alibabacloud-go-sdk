// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateAgentRequest
	GetDescription() *string
	SetDisplayName(v string) *CreateAgentRequest
	GetDisplayName() *string
	SetKnowledgeBases(v interface{}) *CreateAgentRequest
	GetKnowledgeBases() interface{}
	SetName(v string) *CreateAgentRequest
	GetName() *string
	SetSkills(v interface{}) *CreateAgentRequest
	GetSkills() interface{}
	SetSystemPrompt(v string) *CreateAgentRequest
	GetSystemPrompt() *string
	SetTools(v interface{}) *CreateAgentRequest
	GetTools() interface{}
	SetVisibility(v string) *CreateAgentRequest
	GetVisibility() *string
}

type CreateAgentRequest struct {
	// The description of the Agent.
	//
	// example:
	//
	// Analyzes code changes and generates CR review comments
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The display name of the Agent.
	//
	// example:
	//
	// CR Code Review Agent
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The knowledge base reference list, which contains at most one element.
	//
	// example:
	//
	// [{"name":"code-review-guidelines"}]
	KnowledgeBases interface{} `json:"KnowledgeBases,omitempty" xml:"KnowledgeBases,omitempty"`
	// The Agent name, which is also the unique identifier that cannot be modified after creation.
	//
	// This parameter is required.
	//
	// example:
	//
	// code-review-agent
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The Skill reference list. For specific fields, see "Supplementary description of request parameters".
	//
	// example:
	//
	// [{"name":"code-review"}]
	Skills interface{} `json:"Skills,omitempty" xml:"Skills,omitempty"`
	// The system prompt.
	//
	// example:
	//
	// Check the correctness, security, and maintainability of code changes in the CR, and provide review comments by severity
	SystemPrompt *string `json:"SystemPrompt,omitempty" xml:"SystemPrompt,omitempty"`
	// The reference list of MCP Server and Connector names. The same array supports both types of entries. Each entry specifies one type of reference. If items is omitted for an MCP Server, all public tools are included. Previously specified items retain their original values.
	//
	// example:
	//
	// [{"mcpServerName":"code-repository-mcp"},{"connectorName":"code-review-data"}]
	Tools interface{} `json:"Tools,omitempty" xml:"Tools,omitempty"`
	// The visibility scope of the Agent. Valid values:
	//
	// - user
	//
	// - tenant
	//
	// Default value: user.
	//
	// example:
	//
	// user
	Visibility *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s CreateAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentRequest) GoString() string {
	return s.String()
}

func (s *CreateAgentRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateAgentRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateAgentRequest) GetKnowledgeBases() interface{} {
	return s.KnowledgeBases
}

func (s *CreateAgentRequest) GetName() *string {
	return s.Name
}

func (s *CreateAgentRequest) GetSkills() interface{} {
	return s.Skills
}

func (s *CreateAgentRequest) GetSystemPrompt() *string {
	return s.SystemPrompt
}

func (s *CreateAgentRequest) GetTools() interface{} {
	return s.Tools
}

func (s *CreateAgentRequest) GetVisibility() *string {
	return s.Visibility
}

func (s *CreateAgentRequest) SetDescription(v string) *CreateAgentRequest {
	s.Description = &v
	return s
}

func (s *CreateAgentRequest) SetDisplayName(v string) *CreateAgentRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateAgentRequest) SetKnowledgeBases(v interface{}) *CreateAgentRequest {
	s.KnowledgeBases = v
	return s
}

func (s *CreateAgentRequest) SetName(v string) *CreateAgentRequest {
	s.Name = &v
	return s
}

func (s *CreateAgentRequest) SetSkills(v interface{}) *CreateAgentRequest {
	s.Skills = v
	return s
}

func (s *CreateAgentRequest) SetSystemPrompt(v string) *CreateAgentRequest {
	s.SystemPrompt = &v
	return s
}

func (s *CreateAgentRequest) SetTools(v interface{}) *CreateAgentRequest {
	s.Tools = v
	return s
}

func (s *CreateAgentRequest) SetVisibility(v string) *CreateAgentRequest {
	s.Visibility = &v
	return s
}

func (s *CreateAgentRequest) Validate() error {
	return dara.Validate(s)
}
