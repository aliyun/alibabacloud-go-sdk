// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAgent(v *GetAgentResponseBodyAgent) *GetAgentResponseBody
	GetAgent() *GetAgentResponseBodyAgent
	SetRequestId(v string) *GetAgentResponseBody
	GetRequestId() *string
}

type GetAgentResponseBody struct {
	// The agent details.
	Agent *GetAgentResponseBodyAgent `json:"Agent,omitempty" xml:"Agent,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 824F80BA-1778-5D8A-BAFF-668A4D9C4CC7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAgentResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgentResponseBody) GetAgent() *GetAgentResponseBodyAgent {
	return s.Agent
}

func (s *GetAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAgentResponseBody) SetAgent(v *GetAgentResponseBodyAgent) *GetAgentResponseBody {
	s.Agent = v
	return s
}

func (s *GetAgentResponseBody) SetRequestId(v string) *GetAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAgentResponseBody) Validate() error {
	if s.Agent != nil {
		if err := s.Agent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAgentResponseBodyAgent struct {
	// A list of callable sub-agents.
	CallableAgents []*GetAgentResponseBodyAgentCallableAgents `json:"CallableAgents,omitempty" xml:"CallableAgents,omitempty" type:"Repeated"`
	// The creator ID.
	//
	// example:
	//
	// 123456
	CreatorId *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// **The description.**
	//
	// example:
	//
	// 数据分析助手
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// **The display name.**
	//
	// example:
	//
	// 我的助手
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The creation time, as a Unix timestamp in milliseconds.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 1780555634000
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// The last modification time, as a Unix timestamp in milliseconds.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 1780555634000
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// **Additional metadata.**
	//
	// example:
	//
	// {}
	Metadata map[string]interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// **The model configuration.**
	Model *GetAgentResponseBodyAgentModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// The ID of the last modifier.
	//
	// example:
	//
	// 123456
	ModifierId *string `json:"ModifierId,omitempty" xml:"ModifierId,omitempty"`
	// **The agent name.**
	//
	// example:
	//
	// my-agent
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The runtime dependencies.
	RequiredRuntime []*string `json:"RequiredRuntime,omitempty" xml:"RequiredRuntime,omitempty" type:"Repeated"`
	// A list of skills.
	Skills []*GetAgentResponseBodyAgentSkills `json:"Skills,omitempty" xml:"Skills,omitempty" type:"Repeated"`
	// **The system prompt.**
	//
	// example:
	//
	// 你是一个数据分析助手。
	SystemPrompt *string `json:"SystemPrompt,omitempty" xml:"SystemPrompt,omitempty"`
	// **A list of tools.**
	Tools []*GetAgentResponseBodyAgentTools `json:"Tools,omitempty" xml:"Tools,omitempty" type:"Repeated"`
	// **The visibility level.**
	//
	// example:
	//
	// TENANT
	Visibility *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
	// The visibility scope.
	VisibilityScope *GetAgentResponseBodyAgentVisibilityScope `json:"VisibilityScope,omitempty" xml:"VisibilityScope,omitempty" type:"Struct"`
}

func (s GetAgentResponseBodyAgent) String() string {
	return dara.Prettify(s)
}

func (s GetAgentResponseBodyAgent) GoString() string {
	return s.String()
}

func (s *GetAgentResponseBodyAgent) GetCallableAgents() []*GetAgentResponseBodyAgentCallableAgents {
	return s.CallableAgents
}

func (s *GetAgentResponseBodyAgent) GetCreatorId() *string {
	return s.CreatorId
}

func (s *GetAgentResponseBodyAgent) GetDescription() *string {
	return s.Description
}

func (s *GetAgentResponseBodyAgent) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetAgentResponseBodyAgent) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *GetAgentResponseBodyAgent) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *GetAgentResponseBodyAgent) GetMetadata() map[string]interface{} {
	return s.Metadata
}

func (s *GetAgentResponseBodyAgent) GetModel() *GetAgentResponseBodyAgentModel {
	return s.Model
}

func (s *GetAgentResponseBodyAgent) GetModifierId() *string {
	return s.ModifierId
}

func (s *GetAgentResponseBodyAgent) GetName() *string {
	return s.Name
}

func (s *GetAgentResponseBodyAgent) GetRequiredRuntime() []*string {
	return s.RequiredRuntime
}

func (s *GetAgentResponseBodyAgent) GetSkills() []*GetAgentResponseBodyAgentSkills {
	return s.Skills
}

func (s *GetAgentResponseBodyAgent) GetSystemPrompt() *string {
	return s.SystemPrompt
}

func (s *GetAgentResponseBodyAgent) GetTools() []*GetAgentResponseBodyAgentTools {
	return s.Tools
}

func (s *GetAgentResponseBodyAgent) GetVisibility() *string {
	return s.Visibility
}

func (s *GetAgentResponseBodyAgent) GetVisibilityScope() *GetAgentResponseBodyAgentVisibilityScope {
	return s.VisibilityScope
}

func (s *GetAgentResponseBodyAgent) SetCallableAgents(v []*GetAgentResponseBodyAgentCallableAgents) *GetAgentResponseBodyAgent {
	s.CallableAgents = v
	return s
}

func (s *GetAgentResponseBodyAgent) SetCreatorId(v string) *GetAgentResponseBodyAgent {
	s.CreatorId = &v
	return s
}

func (s *GetAgentResponseBodyAgent) SetDescription(v string) *GetAgentResponseBodyAgent {
	s.Description = &v
	return s
}

func (s *GetAgentResponseBodyAgent) SetDisplayName(v string) *GetAgentResponseBodyAgent {
	s.DisplayName = &v
	return s
}

func (s *GetAgentResponseBodyAgent) SetGmtCreateTime(v string) *GetAgentResponseBodyAgent {
	s.GmtCreateTime = &v
	return s
}

func (s *GetAgentResponseBodyAgent) SetGmtModifiedTime(v string) *GetAgentResponseBodyAgent {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetAgentResponseBodyAgent) SetMetadata(v map[string]interface{}) *GetAgentResponseBodyAgent {
	s.Metadata = v
	return s
}

func (s *GetAgentResponseBodyAgent) SetModel(v *GetAgentResponseBodyAgentModel) *GetAgentResponseBodyAgent {
	s.Model = v
	return s
}

func (s *GetAgentResponseBodyAgent) SetModifierId(v string) *GetAgentResponseBodyAgent {
	s.ModifierId = &v
	return s
}

func (s *GetAgentResponseBodyAgent) SetName(v string) *GetAgentResponseBodyAgent {
	s.Name = &v
	return s
}

func (s *GetAgentResponseBodyAgent) SetRequiredRuntime(v []*string) *GetAgentResponseBodyAgent {
	s.RequiredRuntime = v
	return s
}

func (s *GetAgentResponseBodyAgent) SetSkills(v []*GetAgentResponseBodyAgentSkills) *GetAgentResponseBodyAgent {
	s.Skills = v
	return s
}

func (s *GetAgentResponseBodyAgent) SetSystemPrompt(v string) *GetAgentResponseBodyAgent {
	s.SystemPrompt = &v
	return s
}

func (s *GetAgentResponseBodyAgent) SetTools(v []*GetAgentResponseBodyAgentTools) *GetAgentResponseBodyAgent {
	s.Tools = v
	return s
}

func (s *GetAgentResponseBodyAgent) SetVisibility(v string) *GetAgentResponseBodyAgent {
	s.Visibility = &v
	return s
}

func (s *GetAgentResponseBodyAgent) SetVisibilityScope(v *GetAgentResponseBodyAgentVisibilityScope) *GetAgentResponseBodyAgent {
	s.VisibilityScope = v
	return s
}

func (s *GetAgentResponseBodyAgent) Validate() error {
	if s.CallableAgents != nil {
		for _, item := range s.CallableAgents {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Model != nil {
		if err := s.Model.Validate(); err != nil {
			return err
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

type GetAgentResponseBodyAgentCallableAgents struct {
	// The sub-agent display name.
	//
	// example:
	//
	// 子助手
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The sub-agent name.
	//
	// example:
	//
	// sub-agent
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The sub-agent source.
	//
	// example:
	//
	// custom
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The sub-agent version.
	//
	// example:
	//
	// -
	Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetAgentResponseBodyAgentCallableAgents) String() string {
	return dara.Prettify(s)
}

func (s GetAgentResponseBodyAgentCallableAgents) GoString() string {
	return s.String()
}

func (s *GetAgentResponseBodyAgentCallableAgents) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetAgentResponseBodyAgentCallableAgents) GetName() *string {
	return s.Name
}

func (s *GetAgentResponseBodyAgentCallableAgents) GetSource() *string {
	return s.Source
}

func (s *GetAgentResponseBodyAgentCallableAgents) GetVersion() *int32 {
	return s.Version
}

func (s *GetAgentResponseBodyAgentCallableAgents) SetDisplayName(v string) *GetAgentResponseBodyAgentCallableAgents {
	s.DisplayName = &v
	return s
}

func (s *GetAgentResponseBodyAgentCallableAgents) SetName(v string) *GetAgentResponseBodyAgentCallableAgents {
	s.Name = &v
	return s
}

func (s *GetAgentResponseBodyAgentCallableAgents) SetSource(v string) *GetAgentResponseBodyAgentCallableAgents {
	s.Source = &v
	return s
}

func (s *GetAgentResponseBodyAgentCallableAgents) SetVersion(v int32) *GetAgentResponseBodyAgentCallableAgents {
	s.Version = &v
	return s
}

func (s *GetAgentResponseBodyAgentCallableAgents) Validate() error {
	return dara.Validate(s)
}

type GetAgentResponseBodyAgentModel struct {
	// Additional configuration for the model.
	//
	// example:
	//
	// {}
	Config map[string]interface{} `json:"Config,omitempty" xml:"Config,omitempty"`
	// The maximum number of tokens to generate in one response.
	//
	// example:
	//
	// 8192
	MaxTokens *int32 `json:"MaxTokens,omitempty" xml:"MaxTokens,omitempty"`
	// The model name.
	//
	// example:
	//
	// qwen3-max
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// Indicates whether streaming output is enabled.
	//
	// example:
	//
	// true
	Stream *bool `json:"Stream,omitempty" xml:"Stream,omitempty"`
	// The temperature.
	//
	// example:
	//
	// 1
	Temperature *float64 `json:"Temperature,omitempty" xml:"Temperature,omitempty"`
	// The top-p.
	//
	// example:
	//
	// 1
	TopP *float64 `json:"TopP,omitempty" xml:"TopP,omitempty"`
}

func (s GetAgentResponseBodyAgentModel) String() string {
	return dara.Prettify(s)
}

func (s GetAgentResponseBodyAgentModel) GoString() string {
	return s.String()
}

func (s *GetAgentResponseBodyAgentModel) GetConfig() map[string]interface{} {
	return s.Config
}

func (s *GetAgentResponseBodyAgentModel) GetMaxTokens() *int32 {
	return s.MaxTokens
}

func (s *GetAgentResponseBodyAgentModel) GetModelName() *string {
	return s.ModelName
}

func (s *GetAgentResponseBodyAgentModel) GetStream() *bool {
	return s.Stream
}

func (s *GetAgentResponseBodyAgentModel) GetTemperature() *float64 {
	return s.Temperature
}

func (s *GetAgentResponseBodyAgentModel) GetTopP() *float64 {
	return s.TopP
}

func (s *GetAgentResponseBodyAgentModel) SetConfig(v map[string]interface{}) *GetAgentResponseBodyAgentModel {
	s.Config = v
	return s
}

func (s *GetAgentResponseBodyAgentModel) SetMaxTokens(v int32) *GetAgentResponseBodyAgentModel {
	s.MaxTokens = &v
	return s
}

func (s *GetAgentResponseBodyAgentModel) SetModelName(v string) *GetAgentResponseBodyAgentModel {
	s.ModelName = &v
	return s
}

func (s *GetAgentResponseBodyAgentModel) SetStream(v bool) *GetAgentResponseBodyAgentModel {
	s.Stream = &v
	return s
}

func (s *GetAgentResponseBodyAgentModel) SetTemperature(v float64) *GetAgentResponseBodyAgentModel {
	s.Temperature = &v
	return s
}

func (s *GetAgentResponseBodyAgentModel) SetTopP(v float64) *GetAgentResponseBodyAgentModel {
	s.TopP = &v
	return s
}

func (s *GetAgentResponseBodyAgentModel) Validate() error {
	return dara.Validate(s)
}

type GetAgentResponseBodyAgentSkills struct {
	// The skill name.
	//
	// example:
	//
	// my-skill
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The skill version.
	//
	// example:
	//
	// -
	Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetAgentResponseBodyAgentSkills) String() string {
	return dara.Prettify(s)
}

func (s GetAgentResponseBodyAgentSkills) GoString() string {
	return s.String()
}

func (s *GetAgentResponseBodyAgentSkills) GetName() *string {
	return s.Name
}

func (s *GetAgentResponseBodyAgentSkills) GetVersion() *int32 {
	return s.Version
}

func (s *GetAgentResponseBodyAgentSkills) SetName(v string) *GetAgentResponseBodyAgentSkills {
	s.Name = &v
	return s
}

func (s *GetAgentResponseBodyAgentSkills) SetVersion(v int32) *GetAgentResponseBodyAgentSkills {
	s.Version = &v
	return s
}

func (s *GetAgentResponseBodyAgentSkills) Validate() error {
	return dara.Validate(s)
}

type GetAgentResponseBodyAgentTools struct {
	// **The name of the built-in tool. This parameter applies only when `Kind` is set to `builtin`.**
	//
	// example:
	//
	// builtin_sql
	BuiltinName *string `json:"BuiltinName,omitempty" xml:"BuiltinName,omitempty"`
	// **The tool type.**
	//
	// example:
	//
	// builtin
	Kind *string `json:"Kind,omitempty" xml:"Kind,omitempty"`
	// **The selected MCP tool items. This parameter applies only when `Kind` is set to `mcp`.**
	McpItems []*string `json:"McpItems,omitempty" xml:"McpItems,omitempty" type:"Repeated"`
	// **The name of the associated MCP server. This parameter applies only when `Kind` is set to `mcp`.**
	//
	// example:
	//
	// server-name
	McpServerName *string `json:"McpServerName,omitempty" xml:"McpServerName,omitempty"`
}

func (s GetAgentResponseBodyAgentTools) String() string {
	return dara.Prettify(s)
}

func (s GetAgentResponseBodyAgentTools) GoString() string {
	return s.String()
}

func (s *GetAgentResponseBodyAgentTools) GetBuiltinName() *string {
	return s.BuiltinName
}

func (s *GetAgentResponseBodyAgentTools) GetKind() *string {
	return s.Kind
}

func (s *GetAgentResponseBodyAgentTools) GetMcpItems() []*string {
	return s.McpItems
}

func (s *GetAgentResponseBodyAgentTools) GetMcpServerName() *string {
	return s.McpServerName
}

func (s *GetAgentResponseBodyAgentTools) SetBuiltinName(v string) *GetAgentResponseBodyAgentTools {
	s.BuiltinName = &v
	return s
}

func (s *GetAgentResponseBodyAgentTools) SetKind(v string) *GetAgentResponseBodyAgentTools {
	s.Kind = &v
	return s
}

func (s *GetAgentResponseBodyAgentTools) SetMcpItems(v []*string) *GetAgentResponseBodyAgentTools {
	s.McpItems = v
	return s
}

func (s *GetAgentResponseBodyAgentTools) SetMcpServerName(v string) *GetAgentResponseBodyAgentTools {
	s.McpServerName = &v
	return s
}

func (s *GetAgentResponseBodyAgentTools) Validate() error {
	return dara.Validate(s)
}

type GetAgentResponseBodyAgentVisibilityScope struct {
	// A list of project IDs that can view the agent.
	ProjectIds []*string `json:"ProjectIds,omitempty" xml:"ProjectIds,omitempty" type:"Repeated"`
	// A list of user IDs that can view the agent.
	UserIds []*string `json:"UserIds,omitempty" xml:"UserIds,omitempty" type:"Repeated"`
}

func (s GetAgentResponseBodyAgentVisibilityScope) String() string {
	return dara.Prettify(s)
}

func (s GetAgentResponseBodyAgentVisibilityScope) GoString() string {
	return s.String()
}

func (s *GetAgentResponseBodyAgentVisibilityScope) GetProjectIds() []*string {
	return s.ProjectIds
}

func (s *GetAgentResponseBodyAgentVisibilityScope) GetUserIds() []*string {
	return s.UserIds
}

func (s *GetAgentResponseBodyAgentVisibilityScope) SetProjectIds(v []*string) *GetAgentResponseBodyAgentVisibilityScope {
	s.ProjectIds = v
	return s
}

func (s *GetAgentResponseBodyAgentVisibilityScope) SetUserIds(v []*string) *GetAgentResponseBodyAgentVisibilityScope {
	s.UserIds = v
	return s
}

func (s *GetAgentResponseBodyAgentVisibilityScope) Validate() error {
	return dara.Validate(s)
}
