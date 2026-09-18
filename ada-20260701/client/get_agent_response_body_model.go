// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *GetAgentResponseBody
	GetAgentId() *string
	SetAgentVersion(v int64) *GetAgentResponseBody
	GetAgentVersion() *int64
	SetCanDelete(v bool) *GetAgentResponseBody
	GetCanDelete() *bool
	SetCanModify(v bool) *GetAgentResponseBody
	GetCanModify() *bool
	SetCreatedAt(v int64) *GetAgentResponseBody
	GetCreatedAt() *int64
	SetDescription(v string) *GetAgentResponseBody
	GetDescription() *string
	SetDisplayName(v string) *GetAgentResponseBody
	GetDisplayName() *string
	SetKnowledgeBases(v interface{}) *GetAgentResponseBody
	GetKnowledgeBases() interface{}
	SetMetadata(v interface{}) *GetAgentResponseBody
	GetMetadata() interface{}
	SetModel(v interface{}) *GetAgentResponseBody
	GetModel() interface{}
	SetName(v string) *GetAgentResponseBody
	GetName() *string
	SetOfficial(v bool) *GetAgentResponseBody
	GetOfficial() *bool
	SetRequestId(v string) *GetAgentResponseBody
	GetRequestId() *string
	SetSkills(v interface{}) *GetAgentResponseBody
	GetSkills() interface{}
	SetStatus(v string) *GetAgentResponseBody
	GetStatus() *string
	SetSystemPrompt(v string) *GetAgentResponseBody
	GetSystemPrompt() *string
	SetTools(v interface{}) *GetAgentResponseBody
	GetTools() interface{}
	SetUpdatedAt(v int64) *GetAgentResponseBody
	GetUpdatedAt() *int64
	SetVisibility(v string) *GetAgentResponseBody
	GetVisibility() *string
}

type GetAgentResponseBody struct {
	// Agent ID。
	//
	// example:
	//
	// agent_00000000000000000000000000000001
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// The current configuration revision number.
	//
	// example:
	//
	// 2
	AgentVersion *int64 `json:"AgentVersion,omitempty" xml:"AgentVersion,omitempty"`
	// Indicates whether the current identity can delete the agent.
	//
	// example:
	//
	// true
	CanDelete *bool `json:"CanDelete,omitempty" xml:"CanDelete,omitempty"`
	// Indicates whether the current identity can modify the agent.
	//
	// example:
	//
	// true
	CanModify *bool `json:"CanModify,omitempty" xml:"CanModify,omitempty"`
	// The creation time. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1788332400000
	CreatedAt *int64 `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// The description of the agent. This field may not be returned if it is not configured.
	//
	// example:
	//
	// Analyzes code changes and generates CR review comments
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The display name of the agent. This field may not be returned if it is not configured.
	//
	// example:
	//
	// CR Code Review Agent
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The list of knowledge base references. The list contains at most one element.
	//
	// example:
	//
	// [{"name":"code-review-guidelines"}]
	KnowledgeBases interface{} `json:"KnowledgeBases,omitempty" xml:"KnowledgeBases,omitempty"`
	// The display metadata of the agent. For specific fields, see "Supplementary description of response elements".
	//
	// example:
	//
	// {"iconUrl":"https://example.com/icons/code-review-agent.png"}
	Metadata interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// The saved model configuration. This field is returned only for official agents. The value supports an object array and is compatible with legacy single objects and strings. An empty array returns [ \\]. Object arrays preserve the original order, duplicate names, and object fields.
	//
	// example:
	//
	// [{"name":"base","default":true},{"name":"base"},{"name":"base","settings":{"thinking":false,"topK":5}}]
	Model interface{} `json:"Model,omitempty" xml:"Model,omitempty"`
	// The name of the agent.
	//
	// example:
	//
	// code-review-agent
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Indicates whether the agent is an official agent provided by the platform.
	//
	// example:
	//
	// false
	Official *bool `json:"Official,omitempty" xml:"Official,omitempty"`
	// The request ID, which is used for Tracing Analysis and troubleshooting.
	//
	// example:
	//
	// 0A1B2C3D-4E5F-6789-ABCD-EF0123456789
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of skill references. For specific fields, see "Supplementary description of response elements".
	//
	// example:
	//
	// [{"name":"code-review"}]
	Skills interface{} `json:"Skills,omitempty" xml:"Skills,omitempty"`
	// The status of the agent. The default status of a newly created agent is `draft`.
	//
	// example:
	//
	// draft
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The system prompt. This field may not be returned if it is not configured.
	//
	// example:
	//
	// Check the correctness, security, and maintainability of code changes in the CR, and provide review comments by severity
	SystemPrompt *string `json:"SystemPrompt,omitempty" xml:"SystemPrompt,omitempty"`
	// The list of MCP Server and Connector name references. For element fields, see the following section.
	//
	// example:
	//
	// [{"mcpServerName":"code-repository-mcp"},{"connectorName":"code-review-data"}]
	Tools interface{} `json:"Tools,omitempty" xml:"Tools,omitempty"`
	// The most recent update time. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1788332700000
	UpdatedAt *int64 `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	// The visibility scope of the agent. Valid values: `user` and `tenant`.
	//
	// example:
	//
	// user
	Visibility *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s GetAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAgentResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgentResponseBody) GetAgentId() *string {
	return s.AgentId
}

func (s *GetAgentResponseBody) GetAgentVersion() *int64 {
	return s.AgentVersion
}

func (s *GetAgentResponseBody) GetCanDelete() *bool {
	return s.CanDelete
}

func (s *GetAgentResponseBody) GetCanModify() *bool {
	return s.CanModify
}

func (s *GetAgentResponseBody) GetCreatedAt() *int64 {
	return s.CreatedAt
}

func (s *GetAgentResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetAgentResponseBody) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetAgentResponseBody) GetKnowledgeBases() interface{} {
	return s.KnowledgeBases
}

func (s *GetAgentResponseBody) GetMetadata() interface{} {
	return s.Metadata
}

func (s *GetAgentResponseBody) GetModel() interface{} {
	return s.Model
}

func (s *GetAgentResponseBody) GetName() *string {
	return s.Name
}

func (s *GetAgentResponseBody) GetOfficial() *bool {
	return s.Official
}

func (s *GetAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAgentResponseBody) GetSkills() interface{} {
	return s.Skills
}

func (s *GetAgentResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetAgentResponseBody) GetSystemPrompt() *string {
	return s.SystemPrompt
}

func (s *GetAgentResponseBody) GetTools() interface{} {
	return s.Tools
}

func (s *GetAgentResponseBody) GetUpdatedAt() *int64 {
	return s.UpdatedAt
}

func (s *GetAgentResponseBody) GetVisibility() *string {
	return s.Visibility
}

func (s *GetAgentResponseBody) SetAgentId(v string) *GetAgentResponseBody {
	s.AgentId = &v
	return s
}

func (s *GetAgentResponseBody) SetAgentVersion(v int64) *GetAgentResponseBody {
	s.AgentVersion = &v
	return s
}

func (s *GetAgentResponseBody) SetCanDelete(v bool) *GetAgentResponseBody {
	s.CanDelete = &v
	return s
}

func (s *GetAgentResponseBody) SetCanModify(v bool) *GetAgentResponseBody {
	s.CanModify = &v
	return s
}

func (s *GetAgentResponseBody) SetCreatedAt(v int64) *GetAgentResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *GetAgentResponseBody) SetDescription(v string) *GetAgentResponseBody {
	s.Description = &v
	return s
}

func (s *GetAgentResponseBody) SetDisplayName(v string) *GetAgentResponseBody {
	s.DisplayName = &v
	return s
}

func (s *GetAgentResponseBody) SetKnowledgeBases(v interface{}) *GetAgentResponseBody {
	s.KnowledgeBases = v
	return s
}

func (s *GetAgentResponseBody) SetMetadata(v interface{}) *GetAgentResponseBody {
	s.Metadata = v
	return s
}

func (s *GetAgentResponseBody) SetModel(v interface{}) *GetAgentResponseBody {
	s.Model = v
	return s
}

func (s *GetAgentResponseBody) SetName(v string) *GetAgentResponseBody {
	s.Name = &v
	return s
}

func (s *GetAgentResponseBody) SetOfficial(v bool) *GetAgentResponseBody {
	s.Official = &v
	return s
}

func (s *GetAgentResponseBody) SetRequestId(v string) *GetAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAgentResponseBody) SetSkills(v interface{}) *GetAgentResponseBody {
	s.Skills = v
	return s
}

func (s *GetAgentResponseBody) SetStatus(v string) *GetAgentResponseBody {
	s.Status = &v
	return s
}

func (s *GetAgentResponseBody) SetSystemPrompt(v string) *GetAgentResponseBody {
	s.SystemPrompt = &v
	return s
}

func (s *GetAgentResponseBody) SetTools(v interface{}) *GetAgentResponseBody {
	s.Tools = v
	return s
}

func (s *GetAgentResponseBody) SetUpdatedAt(v int64) *GetAgentResponseBody {
	s.UpdatedAt = &v
	return s
}

func (s *GetAgentResponseBody) SetVisibility(v string) *GetAgentResponseBody {
	s.Visibility = &v
	return s
}

func (s *GetAgentResponseBody) Validate() error {
	return dara.Validate(s)
}
