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
	Agent *GetAgentResponseBodyAgent `json:"Agent,omitempty" xml:"Agent,omitempty" type:"Struct"`
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
	CallableAgents []*GetAgentResponseBodyAgentCallableAgents `json:"CallableAgents,omitempty" xml:"CallableAgents,omitempty" type:"Repeated"`
	// example:
	//
	// 123456
	CreatorId   *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 1780555634000
	GmtModifiedTime *string                         `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	Metadata        map[string]interface{}          `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	Model           *GetAgentResponseBodyAgentModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// example:
	//
	// 123456
	ModifierId *string `json:"ModifierId,omitempty" xml:"ModifierId,omitempty"`
	// example:
	//
	// my-agent
	Name            *string                            `json:"Name,omitempty" xml:"Name,omitempty"`
	RequiredRuntime []*string                          `json:"RequiredRuntime,omitempty" xml:"RequiredRuntime,omitempty" type:"Repeated"`
	Skills          []*GetAgentResponseBodyAgentSkills `json:"Skills,omitempty" xml:"Skills,omitempty" type:"Repeated"`
	SystemPrompt    *string                            `json:"SystemPrompt,omitempty" xml:"SystemPrompt,omitempty"`
	Tools           []*GetAgentResponseBodyAgentTools  `json:"Tools,omitempty" xml:"Tools,omitempty" type:"Repeated"`
	// example:
	//
	// TENANT
	Visibility      *string                                   `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
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
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// sub-agent
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Source  *string `json:"Source,omitempty" xml:"Source,omitempty"`
	Version *int32  `json:"Version,omitempty" xml:"Version,omitempty"`
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
	Config    map[string]interface{} `json:"Config,omitempty" xml:"Config,omitempty"`
	MaxTokens *int32                 `json:"MaxTokens,omitempty" xml:"MaxTokens,omitempty"`
	// example:
	//
	// qwen3-max
	ModelName   *string  `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	Stream      *bool    `json:"Stream,omitempty" xml:"Stream,omitempty"`
	Temperature *float64 `json:"Temperature,omitempty" xml:"Temperature,omitempty"`
	TopP        *float64 `json:"TopP,omitempty" xml:"TopP,omitempty"`
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
	// example:
	//
	// my-skill
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Version *int32  `json:"Version,omitempty" xml:"Version,omitempty"`
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
	BuiltinName *string `json:"BuiltinName,omitempty" xml:"BuiltinName,omitempty"`
	// example:
	//
	// builtin
	Kind          *string   `json:"Kind,omitempty" xml:"Kind,omitempty"`
	McpItems      []*string `json:"McpItems,omitempty" xml:"McpItems,omitempty" type:"Repeated"`
	McpServerName *string   `json:"McpServerName,omitempty" xml:"McpServerName,omitempty"`
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
	ProjectIds []*string `json:"ProjectIds,omitempty" xml:"ProjectIds,omitempty" type:"Repeated"`
	UserIds    []*string `json:"UserIds,omitempty" xml:"UserIds,omitempty" type:"Repeated"`
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
