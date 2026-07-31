// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMmAppAndBindingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UpdateMmAppAndBindingRequest
	GetAppId() *string
	SetAppName(v string) *UpdateMmAppAndBindingRequest
	GetAppName() *string
	SetBindingConfig(v *UpdateMmAppAndBindingRequestBindingConfig) *UpdateMmAppAndBindingRequest
	GetBindingConfig() *UpdateMmAppAndBindingRequestBindingConfig
	SetConversationConfig(v *UpdateMmAppAndBindingRequestConversationConfig) *UpdateMmAppAndBindingRequest
	GetConversationConfig() *UpdateMmAppAndBindingRequestConversationConfig
	SetMemoryConfig(v *UpdateMmAppAndBindingRequestMemoryConfig) *UpdateMmAppAndBindingRequest
	GetMemoryConfig() *UpdateMmAppAndBindingRequestMemoryConfig
	SetModelConfig(v *UpdateMmAppAndBindingRequestModelConfig) *UpdateMmAppAndBindingRequest
	GetModelConfig() *UpdateMmAppAndBindingRequestModelConfig
	SetPrompt(v string) *UpdateMmAppAndBindingRequest
	GetPrompt() *string
	SetWorkspaceId(v string) *UpdateMmAppAndBindingRequest
	GetWorkspaceId() *string
}

type UpdateMmAppAndBindingRequest struct {
	// This parameter is required.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	AppName            *string                                         `json:"AppName,omitempty" xml:"AppName,omitempty"`
	BindingConfig      *UpdateMmAppAndBindingRequestBindingConfig      `json:"BindingConfig,omitempty" xml:"BindingConfig,omitempty" type:"Struct"`
	ConversationConfig *UpdateMmAppAndBindingRequestConversationConfig `json:"ConversationConfig,omitempty" xml:"ConversationConfig,omitempty" type:"Struct"`
	MemoryConfig       *UpdateMmAppAndBindingRequestMemoryConfig       `json:"MemoryConfig,omitempty" xml:"MemoryConfig,omitempty" type:"Struct"`
	ModelConfig        *UpdateMmAppAndBindingRequestModelConfig        `json:"ModelConfig,omitempty" xml:"ModelConfig,omitempty" type:"Struct"`
	Prompt             *string                                         `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateMmAppAndBindingRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMmAppAndBindingRequest) GoString() string {
	return s.String()
}

func (s *UpdateMmAppAndBindingRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateMmAppAndBindingRequest) GetAppName() *string {
	return s.AppName
}

func (s *UpdateMmAppAndBindingRequest) GetBindingConfig() *UpdateMmAppAndBindingRequestBindingConfig {
	return s.BindingConfig
}

func (s *UpdateMmAppAndBindingRequest) GetConversationConfig() *UpdateMmAppAndBindingRequestConversationConfig {
	return s.ConversationConfig
}

func (s *UpdateMmAppAndBindingRequest) GetMemoryConfig() *UpdateMmAppAndBindingRequestMemoryConfig {
	return s.MemoryConfig
}

func (s *UpdateMmAppAndBindingRequest) GetModelConfig() *UpdateMmAppAndBindingRequestModelConfig {
	return s.ModelConfig
}

func (s *UpdateMmAppAndBindingRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *UpdateMmAppAndBindingRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateMmAppAndBindingRequest) SetAppId(v string) *UpdateMmAppAndBindingRequest {
	s.AppId = &v
	return s
}

func (s *UpdateMmAppAndBindingRequest) SetAppName(v string) *UpdateMmAppAndBindingRequest {
	s.AppName = &v
	return s
}

func (s *UpdateMmAppAndBindingRequest) SetBindingConfig(v *UpdateMmAppAndBindingRequestBindingConfig) *UpdateMmAppAndBindingRequest {
	s.BindingConfig = v
	return s
}

func (s *UpdateMmAppAndBindingRequest) SetConversationConfig(v *UpdateMmAppAndBindingRequestConversationConfig) *UpdateMmAppAndBindingRequest {
	s.ConversationConfig = v
	return s
}

func (s *UpdateMmAppAndBindingRequest) SetMemoryConfig(v *UpdateMmAppAndBindingRequestMemoryConfig) *UpdateMmAppAndBindingRequest {
	s.MemoryConfig = v
	return s
}

func (s *UpdateMmAppAndBindingRequest) SetModelConfig(v *UpdateMmAppAndBindingRequestModelConfig) *UpdateMmAppAndBindingRequest {
	s.ModelConfig = v
	return s
}

func (s *UpdateMmAppAndBindingRequest) SetPrompt(v string) *UpdateMmAppAndBindingRequest {
	s.Prompt = &v
	return s
}

func (s *UpdateMmAppAndBindingRequest) SetWorkspaceId(v string) *UpdateMmAppAndBindingRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateMmAppAndBindingRequest) Validate() error {
	if s.BindingConfig != nil {
		if err := s.BindingConfig.Validate(); err != nil {
			return err
		}
	}
	if s.ConversationConfig != nil {
		if err := s.ConversationConfig.Validate(); err != nil {
			return err
		}
	}
	if s.MemoryConfig != nil {
		if err := s.MemoryConfig.Validate(); err != nil {
			return err
		}
	}
	if s.ModelConfig != nil {
		if err := s.ModelConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateMmAppAndBindingRequestBindingConfig struct {
	Agents    []*UpdateMmAppAndBindingRequestBindingConfigAgents   `json:"Agents,omitempty" xml:"Agents,omitempty" type:"Repeated"`
	Commands  []*UpdateMmAppAndBindingRequestBindingConfigCommands `json:"Commands,omitempty" xml:"Commands,omitempty" type:"Repeated"`
	Mcps      []*UpdateMmAppAndBindingRequestBindingConfigMcps     `json:"Mcps,omitempty" xml:"Mcps,omitempty" type:"Repeated"`
	Plugins   []*UpdateMmAppAndBindingRequestBindingConfigPlugins  `json:"Plugins,omitempty" xml:"Plugins,omitempty" type:"Repeated"`
	RagConfig *UpdateMmAppAndBindingRequestBindingConfigRagConfig  `json:"RagConfig,omitempty" xml:"RagConfig,omitempty" type:"Struct"`
}

func (s UpdateMmAppAndBindingRequestBindingConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateMmAppAndBindingRequestBindingConfig) GoString() string {
	return s.String()
}

func (s *UpdateMmAppAndBindingRequestBindingConfig) GetAgents() []*UpdateMmAppAndBindingRequestBindingConfigAgents {
	return s.Agents
}

func (s *UpdateMmAppAndBindingRequestBindingConfig) GetCommands() []*UpdateMmAppAndBindingRequestBindingConfigCommands {
	return s.Commands
}

func (s *UpdateMmAppAndBindingRequestBindingConfig) GetMcps() []*UpdateMmAppAndBindingRequestBindingConfigMcps {
	return s.Mcps
}

func (s *UpdateMmAppAndBindingRequestBindingConfig) GetPlugins() []*UpdateMmAppAndBindingRequestBindingConfigPlugins {
	return s.Plugins
}

func (s *UpdateMmAppAndBindingRequestBindingConfig) GetRagConfig() *UpdateMmAppAndBindingRequestBindingConfigRagConfig {
	return s.RagConfig
}

func (s *UpdateMmAppAndBindingRequestBindingConfig) SetAgents(v []*UpdateMmAppAndBindingRequestBindingConfigAgents) *UpdateMmAppAndBindingRequestBindingConfig {
	s.Agents = v
	return s
}

func (s *UpdateMmAppAndBindingRequestBindingConfig) SetCommands(v []*UpdateMmAppAndBindingRequestBindingConfigCommands) *UpdateMmAppAndBindingRequestBindingConfig {
	s.Commands = v
	return s
}

func (s *UpdateMmAppAndBindingRequestBindingConfig) SetMcps(v []*UpdateMmAppAndBindingRequestBindingConfigMcps) *UpdateMmAppAndBindingRequestBindingConfig {
	s.Mcps = v
	return s
}

func (s *UpdateMmAppAndBindingRequestBindingConfig) SetPlugins(v []*UpdateMmAppAndBindingRequestBindingConfigPlugins) *UpdateMmAppAndBindingRequestBindingConfig {
	s.Plugins = v
	return s
}

func (s *UpdateMmAppAndBindingRequestBindingConfig) SetRagConfig(v *UpdateMmAppAndBindingRequestBindingConfigRagConfig) *UpdateMmAppAndBindingRequestBindingConfig {
	s.RagConfig = v
	return s
}

func (s *UpdateMmAppAndBindingRequestBindingConfig) Validate() error {
	if s.Agents != nil {
		for _, item := range s.Agents {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Commands != nil {
		for _, item := range s.Commands {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Mcps != nil {
		for _, item := range s.Mcps {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Plugins != nil {
		for _, item := range s.Plugins {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RagConfig != nil {
		if err := s.RagConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateMmAppAndBindingRequestBindingConfigAgents struct {
	AgentCode           *string                                                   `json:"AgentCode,omitempty" xml:"AgentCode,omitempty"`
	AgentName           *string                                                   `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	AgentType           *string                                                   `json:"AgentType,omitempty" xml:"AgentType,omitempty"`
	CentralConfig       map[string]interface{}                                    `json:"CentralConfig,omitempty" xml:"CentralConfig,omitempty"`
	Description         *string                                                   `json:"Description,omitempty" xml:"Description,omitempty"`
	IntentFewShotConfig map[string][]*BindingConfigAgentsIntentFewShotConfigValue `json:"IntentFewShotConfig,omitempty" xml:"IntentFewShotConfig,omitempty"`
	OwnConfig           map[string]interface{}                                    `json:"OwnConfig,omitempty" xml:"OwnConfig,omitempty"`
}

func (s UpdateMmAppAndBindingRequestBindingConfigAgents) String() string {
	return dara.Prettify(s)
}

func (s UpdateMmAppAndBindingRequestBindingConfigAgents) GoString() string {
	return s.String()
}

func (s *UpdateMmAppAndBindingRequestBindingConfigAgents) GetAgentCode() *string {
	return s.AgentCode
}

func (s *UpdateMmAppAndBindingRequestBindingConfigAgents) GetAgentName() *string {
	return s.AgentName
}

func (s *UpdateMmAppAndBindingRequestBindingConfigAgents) GetAgentType() *string {
	return s.AgentType
}

func (s *UpdateMmAppAndBindingRequestBindingConfigAgents) GetCentralConfig() map[string]interface{} {
	return s.CentralConfig
}

func (s *UpdateMmAppAndBindingRequestBindingConfigAgents) GetDescription() *string {
	return s.Description
}

func (s *UpdateMmAppAndBindingRequestBindingConfigAgents) GetIntentFewShotConfig() map[string][]*BindingConfigAgentsIntentFewShotConfigValue {
	return s.IntentFewShotConfig
}

func (s *UpdateMmAppAndBindingRequestBindingConfigAgents) GetOwnConfig() map[string]interface{} {
	return s.OwnConfig
}

func (s *UpdateMmAppAndBindingRequestBindingConfigAgents) SetAgentCode(v string) *UpdateMmAppAndBindingRequestBindingConfigAgents {
	s.AgentCode = &v
	return s
}

func (s *UpdateMmAppAndBindingRequestBindingConfigAgents) SetAgentName(v string) *UpdateMmAppAndBindingRequestBindingConfigAgents {
	s.AgentName = &v
	return s
}

func (s *UpdateMmAppAndBindingRequestBindingConfigAgents) SetAgentType(v string) *UpdateMmAppAndBindingRequestBindingConfigAgents {
	s.AgentType = &v
	return s
}

func (s *UpdateMmAppAndBindingRequestBindingConfigAgents) SetCentralConfig(v map[string]interface{}) *UpdateMmAppAndBindingRequestBindingConfigAgents {
	s.CentralConfig = v
	return s
}

func (s *UpdateMmAppAndBindingRequestBindingConfigAgents) SetDescription(v string) *UpdateMmAppAndBindingRequestBindingConfigAgents {
	s.Description = &v
	return s
}

func (s *UpdateMmAppAndBindingRequestBindingConfigAgents) SetIntentFewShotConfig(v map[string][]*BindingConfigAgentsIntentFewShotConfigValue) *UpdateMmAppAndBindingRequestBindingConfigAgents {
	s.IntentFewShotConfig = v
	return s
}

func (s *UpdateMmAppAndBindingRequestBindingConfigAgents) SetOwnConfig(v map[string]interface{}) *UpdateMmAppAndBindingRequestBindingConfigAgents {
	s.OwnConfig = v
	return s
}

func (s *UpdateMmAppAndBindingRequestBindingConfigAgents) Validate() error {
	return dara.Validate(s)
}

type UpdateMmAppAndBindingRequestBindingConfigCommands struct {
	DomainCode *string                                                   `json:"DomainCode,omitempty" xml:"DomainCode,omitempty"`
	DomainName *string                                                   `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Tools      []*UpdateMmAppAndBindingRequestBindingConfigCommandsTools `json:"Tools,omitempty" xml:"Tools,omitempty" type:"Repeated"`
	Type       *string                                                   `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateMmAppAndBindingRequestBindingConfigCommands) String() string {
	return dara.Prettify(s)
}

func (s UpdateMmAppAndBindingRequestBindingConfigCommands) GoString() string {
	return s.String()
}

func (s *UpdateMmAppAndBindingRequestBindingConfigCommands) GetDomainCode() *string {
	return s.DomainCode
}

func (s *UpdateMmAppAndBindingRequestBindingConfigCommands) GetDomainName() *string {
	return s.DomainName
}

func (s *UpdateMmAppAndBindingRequestBindingConfigCommands) GetTools() []*UpdateMmAppAndBindingRequestBindingConfigCommandsTools {
	return s.Tools
}

func (s *UpdateMmAppAndBindingRequestBindingConfigCommands) GetType() *string {
	return s.Type
}

func (s *UpdateMmAppAndBindingRequestBindingConfigCommands) SetDomainCode(v string) *UpdateMmAppAndBindingRequestBindingConfigCommands {
	s.DomainCode = &v
	return s
}

func (s *UpdateMmAppAndBindingRequestBindingConfigCommands) SetDomainName(v string) *UpdateMmAppAndBindingRequestBindingConfigCommands {
	s.DomainName = &v
	return s
}

func (s *UpdateMmAppAndBindingRequestBindingConfigCommands) SetTools(v []*UpdateMmAppAndBindingRequestBindingConfigCommandsTools) *UpdateMmAppAndBindingRequestBindingConfigCommands {
	s.Tools = v
	return s
}

func (s *UpdateMmAppAndBindingRequestBindingConfigCommands) SetType(v string) *UpdateMmAppAndBindingRequestBindingConfigCommands {
	s.Type = &v
	return s
}

func (s *UpdateMmAppAndBindingRequestBindingConfigCommands) Validate() error {
	if s.Tools != nil {
		for _, item := range s.Tools {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateMmAppAndBindingRequestBindingConfigCommandsTools struct {
	ReplyMode       *string                                                               `json:"ReplyMode,omitempty" xml:"ReplyMode,omitempty"`
	ToolDescription *string                                                               `json:"ToolDescription,omitempty" xml:"ToolDescription,omitempty"`
	ToolExamples    []*UpdateMmAppAndBindingRequestBindingConfigCommandsToolsToolExamples `json:"ToolExamples,omitempty" xml:"ToolExamples,omitempty" type:"Repeated"`
	ToolId          *string                                                               `json:"ToolId,omitempty" xml:"ToolId,omitempty"`
	ToolName        *string                                                               `json:"ToolName,omitempty" xml:"ToolName,omitempty"`
	ToolParams      []*UpdateMmAppAndBindingRequestBindingConfigCommandsToolsToolParams   `json:"ToolParams,omitempty" xml:"ToolParams,omitempty" type:"Repeated"`
}

func (s UpdateMmAppAndBindingRequestBindingConfigCommandsTools) String() string {
	return dara.Prettify(s)
}

func (s UpdateMmAppAndBindingRequestBindingConfigCommandsTools) GoString() string {
	return s.String()
}

func (s *UpdateMmAppAndBindingRequestBindingConfigCommandsTools) GetReplyMode() *string {
	return s.ReplyMode
}

func (s *UpdateMmAppAndBindingRequestBindingConfigCommandsTools) GetToolDescription() *string {
	return s.ToolDescription
}

func (s *UpdateMmAppAndBindingRequestBindingConfigCommandsTools) GetToolExamples() []*UpdateMmAppAndBindingRequestBindingConfigCommandsToolsToolExamples {
	return s.ToolExamples
}

func (s *UpdateMmAppAndBindingRequestBindingConfigCommandsTools) GetToolId() *string {
	return s.ToolId
}

func (s *UpdateMmAppAndBindingRequestBindingConfigCommandsTools) GetToolName() *string {
	return s.ToolName
}

func (s *UpdateMmAppAndBindingRequestBindingConfigCommandsTools) GetToolParams() []*UpdateMmAppAndBindingRequestBindingConfigCommandsToolsToolParams {
	return s.ToolParams
}

func (s *UpdateMmAppAndBindingRequestBindingConfigCommandsTools) SetReplyMode(v string) *UpdateMmAppAndBindingRequestBindingConfigCommandsTools {
	s.ReplyMode = &v
	return s
}

func (s *UpdateMmAppAndBindingRequestBindingConfigCommandsTools) SetToolDescription(v string) *UpdateMmAppAndBindingRequestBindingConfigCommandsTools {
	s.ToolDescription = &v
	return s
}

func (s *UpdateMmAppAndBindingRequestBindingConfigCommandsTools) SetToolExamples(v []*UpdateMmAppAndBindingRequestBindingConfigCommandsToolsToolExamples) *UpdateMmAppAndBindingRequestBindingConfigCommandsTools {
	s.ToolExamples = v
	return s
}

func (s *UpdateMmAppAndBindingRequestBindingConfigCommandsTools) SetToolId(v string) *UpdateMmAppAndBindingRequestBindingConfigCommandsTools {
	s.ToolId = &v
	return s
}

func (s *UpdateMmAppAndBindingRequestBindingConfigCommandsTools) SetToolName(v string) *UpdateMmAppAndBindingRequestBindingConfigCommandsTools {
	s.ToolName = &v
	return s
}

func (s *UpdateMmAppAndBindingRequestBindingConfigCommandsTools) SetToolParams(v []*UpdateMmAppAndBindingRequestBindingConfigCommandsToolsToolParams) *UpdateMmAppAndBindingRequestBindingConfigCommandsTools {
	s.ToolParams = v
	return s
}

func (s *UpdateMmAppAndBindingRequestBindingConfigCommandsTools) Validate() error {
	if s.ToolExamples != nil {
		for _, item := range s.ToolExamples {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ToolParams != nil {
		for _, item := range s.ToolParams {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateMmAppAndBindingRequestBindingConfigCommandsToolsToolExamples struct {
	Parameters map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	Query      *string                `json:"Query,omitempty" xml:"Query,omitempty"`
}

func (s UpdateMmAppAndBindingRequestBindingConfigCommandsToolsToolExamples) String() string {
	return dara.Prettify(s)
}

func (s UpdateMmAppAndBindingRequestBindingConfigCommandsToolsToolExamples) GoString() string {
	return s.String()
}

func (s *UpdateMmAppAndBindingRequestBindingConfigCommandsToolsToolExamples) GetParameters() map[string]interface{} {
	return s.Parameters
}

func (s *UpdateMmAppAndBindingRequestBindingConfigCommandsToolsToolExamples) GetQuery() *string {
	return s.Query
}

func (s *UpdateMmAppAndBindingRequestBindingConfigCommandsToolsToolExamples) SetParameters(v map[string]interface{}) *UpdateMmAppAndBindingRequestBindingConfigCommandsToolsToolExamples {
	s.Parameters = v
	return s
}

func (s *UpdateMmAppAndBindingRequestBindingConfigCommandsToolsToolExamples) SetQuery(v string) *UpdateMmAppAndBindingRequestBindingConfigCommandsToolsToolExamples {
	s.Query = &v
	return s
}

func (s *UpdateMmAppAndBindingRequestBindingConfigCommandsToolsToolExamples) Validate() error {
	return dara.Validate(s)
}

type UpdateMmAppAndBindingRequestBindingConfigCommandsToolsToolParams struct {
	ParamDesc    *string `json:"ParamDesc,omitempty" xml:"ParamDesc,omitempty"`
	ParamExample *string `json:"ParamExample,omitempty" xml:"ParamExample,omitempty"`
	ParamName    *string `json:"ParamName,omitempty" xml:"ParamName,omitempty"`
	ParamType    *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	Required     *bool   `json:"Required,omitempty" xml:"Required,omitempty"`
}

func (s UpdateMmAppAndBindingRequestBindingConfigCommandsToolsToolParams) String() string {
	return dara.Prettify(s)
}

func (s UpdateMmAppAndBindingRequestBindingConfigCommandsToolsToolParams) GoString() string {
	return s.String()
}

func (s *UpdateMmAppAndBindingRequestBindingConfigCommandsToolsToolParams) GetParamDesc() *string {
	return s.ParamDesc
}

func (s *UpdateMmAppAndBindingRequestBindingConfigCommandsToolsToolParams) GetParamExample() *string {
	return s.ParamExample
}

func (s *UpdateMmAppAndBindingRequestBindingConfigCommandsToolsToolParams) GetParamName() *string {
	return s.ParamName
}

func (s *UpdateMmAppAndBindingRequestBindingConfigCommandsToolsToolParams) GetParamType() *string {
	return s.ParamType
}

func (s *UpdateMmAppAndBindingRequestBindingConfigCommandsToolsToolParams) GetRequired() *bool {
	return s.Required
}

func (s *UpdateMmAppAndBindingRequestBindingConfigCommandsToolsToolParams) SetParamDesc(v string) *UpdateMmAppAndBindingRequestBindingConfigCommandsToolsToolParams {
	s.ParamDesc = &v
	return s
}

func (s *UpdateMmAppAndBindingRequestBindingConfigCommandsToolsToolParams) SetParamExample(v string) *UpdateMmAppAndBindingRequestBindingConfigCommandsToolsToolParams {
	s.ParamExample = &v
	return s
}

func (s *UpdateMmAppAndBindingRequestBindingConfigCommandsToolsToolParams) SetParamName(v string) *UpdateMmAppAndBindingRequestBindingConfigCommandsToolsToolParams {
	s.ParamName = &v
	return s
}

func (s *UpdateMmAppAndBindingRequestBindingConfigCommandsToolsToolParams) SetParamType(v string) *UpdateMmAppAndBindingRequestBindingConfigCommandsToolsToolParams {
	s.ParamType = &v
	return s
}

func (s *UpdateMmAppAndBindingRequestBindingConfigCommandsToolsToolParams) SetRequired(v bool) *UpdateMmAppAndBindingRequestBindingConfigCommandsToolsToolParams {
	s.Required = &v
	return s
}

func (s *UpdateMmAppAndBindingRequestBindingConfigCommandsToolsToolParams) Validate() error {
	return dara.Validate(s)
}

type UpdateMmAppAndBindingRequestBindingConfigMcps struct {
	Code     *string   `json:"Code,omitempty" xml:"Code,omitempty"`
	ToolList []*string `json:"ToolList,omitempty" xml:"ToolList,omitempty" type:"Repeated"`
	Type     *string   `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateMmAppAndBindingRequestBindingConfigMcps) String() string {
	return dara.Prettify(s)
}

func (s UpdateMmAppAndBindingRequestBindingConfigMcps) GoString() string {
	return s.String()
}

func (s *UpdateMmAppAndBindingRequestBindingConfigMcps) GetCode() *string {
	return s.Code
}

func (s *UpdateMmAppAndBindingRequestBindingConfigMcps) GetToolList() []*string {
	return s.ToolList
}

func (s *UpdateMmAppAndBindingRequestBindingConfigMcps) GetType() *string {
	return s.Type
}

func (s *UpdateMmAppAndBindingRequestBindingConfigMcps) SetCode(v string) *UpdateMmAppAndBindingRequestBindingConfigMcps {
	s.Code = &v
	return s
}

func (s *UpdateMmAppAndBindingRequestBindingConfigMcps) SetToolList(v []*string) *UpdateMmAppAndBindingRequestBindingConfigMcps {
	s.ToolList = v
	return s
}

func (s *UpdateMmAppAndBindingRequestBindingConfigMcps) SetType(v string) *UpdateMmAppAndBindingRequestBindingConfigMcps {
	s.Type = &v
	return s
}

func (s *UpdateMmAppAndBindingRequestBindingConfigMcps) Validate() error {
	return dara.Validate(s)
}

type UpdateMmAppAndBindingRequestBindingConfigPlugins struct {
	PluginCode *string `json:"PluginCode,omitempty" xml:"PluginCode,omitempty"`
	PluginName *string `json:"PluginName,omitempty" xml:"PluginName,omitempty"`
	PluginType *string `json:"PluginType,omitempty" xml:"PluginType,omitempty"`
}

func (s UpdateMmAppAndBindingRequestBindingConfigPlugins) String() string {
	return dara.Prettify(s)
}

func (s UpdateMmAppAndBindingRequestBindingConfigPlugins) GoString() string {
	return s.String()
}

func (s *UpdateMmAppAndBindingRequestBindingConfigPlugins) GetPluginCode() *string {
	return s.PluginCode
}

func (s *UpdateMmAppAndBindingRequestBindingConfigPlugins) GetPluginName() *string {
	return s.PluginName
}

func (s *UpdateMmAppAndBindingRequestBindingConfigPlugins) GetPluginType() *string {
	return s.PluginType
}

func (s *UpdateMmAppAndBindingRequestBindingConfigPlugins) SetPluginCode(v string) *UpdateMmAppAndBindingRequestBindingConfigPlugins {
	s.PluginCode = &v
	return s
}

func (s *UpdateMmAppAndBindingRequestBindingConfigPlugins) SetPluginName(v string) *UpdateMmAppAndBindingRequestBindingConfigPlugins {
	s.PluginName = &v
	return s
}

func (s *UpdateMmAppAndBindingRequestBindingConfigPlugins) SetPluginType(v string) *UpdateMmAppAndBindingRequestBindingConfigPlugins {
	s.PluginType = &v
	return s
}

func (s *UpdateMmAppAndBindingRequestBindingConfigPlugins) Validate() error {
	return dara.Validate(s)
}

type UpdateMmAppAndBindingRequestBindingConfigRagConfig struct {
	EnableSearch          *bool               `json:"EnableSearch,omitempty" xml:"EnableSearch,omitempty"`
	KnowledgeBaseCodeList []*string           `json:"KnowledgeBaseCodeList,omitempty" xml:"KnowledgeBaseCodeList,omitempty" type:"Repeated"`
	PromptStrategy        *string             `json:"PromptStrategy,omitempty" xml:"PromptStrategy,omitempty"`
	RankWeights           map[string]*float64 `json:"RankWeights,omitempty" xml:"RankWeights,omitempty"`
	RetrieveMaxLength     *int32              `json:"RetrieveMaxLength,omitempty" xml:"RetrieveMaxLength,omitempty"`
	TopK                  *int32              `json:"TopK,omitempty" xml:"TopK,omitempty"`
}

func (s UpdateMmAppAndBindingRequestBindingConfigRagConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateMmAppAndBindingRequestBindingConfigRagConfig) GoString() string {
	return s.String()
}

func (s *UpdateMmAppAndBindingRequestBindingConfigRagConfig) GetEnableSearch() *bool {
	return s.EnableSearch
}

func (s *UpdateMmAppAndBindingRequestBindingConfigRagConfig) GetKnowledgeBaseCodeList() []*string {
	return s.KnowledgeBaseCodeList
}

func (s *UpdateMmAppAndBindingRequestBindingConfigRagConfig) GetPromptStrategy() *string {
	return s.PromptStrategy
}

func (s *UpdateMmAppAndBindingRequestBindingConfigRagConfig) GetRankWeights() map[string]*float64 {
	return s.RankWeights
}

func (s *UpdateMmAppAndBindingRequestBindingConfigRagConfig) GetRetrieveMaxLength() *int32 {
	return s.RetrieveMaxLength
}

func (s *UpdateMmAppAndBindingRequestBindingConfigRagConfig) GetTopK() *int32 {
	return s.TopK
}

func (s *UpdateMmAppAndBindingRequestBindingConfigRagConfig) SetEnableSearch(v bool) *UpdateMmAppAndBindingRequestBindingConfigRagConfig {
	s.EnableSearch = &v
	return s
}

func (s *UpdateMmAppAndBindingRequestBindingConfigRagConfig) SetKnowledgeBaseCodeList(v []*string) *UpdateMmAppAndBindingRequestBindingConfigRagConfig {
	s.KnowledgeBaseCodeList = v
	return s
}

func (s *UpdateMmAppAndBindingRequestBindingConfigRagConfig) SetPromptStrategy(v string) *UpdateMmAppAndBindingRequestBindingConfigRagConfig {
	s.PromptStrategy = &v
	return s
}

func (s *UpdateMmAppAndBindingRequestBindingConfigRagConfig) SetRankWeights(v map[string]*float64) *UpdateMmAppAndBindingRequestBindingConfigRagConfig {
	s.RankWeights = v
	return s
}

func (s *UpdateMmAppAndBindingRequestBindingConfigRagConfig) SetRetrieveMaxLength(v int32) *UpdateMmAppAndBindingRequestBindingConfigRagConfig {
	s.RetrieveMaxLength = &v
	return s
}

func (s *UpdateMmAppAndBindingRequestBindingConfigRagConfig) SetTopK(v int32) *UpdateMmAppAndBindingRequestBindingConfigRagConfig {
	s.TopK = &v
	return s
}

func (s *UpdateMmAppAndBindingRequestBindingConfigRagConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateMmAppAndBindingRequestConversationConfig struct {
	AsrModel         *string `json:"AsrModel,omitempty" xml:"AsrModel,omitempty"`
	OpenAsr          *bool   `json:"OpenAsr,omitempty" xml:"OpenAsr,omitempty"`
	OpenTts          *bool   `json:"OpenTts,omitempty" xml:"OpenTts,omitempty"`
	StopOrRejectFlag *bool   `json:"StopOrRejectFlag,omitempty" xml:"StopOrRejectFlag,omitempty"`
	TtsModel         *string `json:"TtsModel,omitempty" xml:"TtsModel,omitempty"`
}

func (s UpdateMmAppAndBindingRequestConversationConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateMmAppAndBindingRequestConversationConfig) GoString() string {
	return s.String()
}

func (s *UpdateMmAppAndBindingRequestConversationConfig) GetAsrModel() *string {
	return s.AsrModel
}

func (s *UpdateMmAppAndBindingRequestConversationConfig) GetOpenAsr() *bool {
	return s.OpenAsr
}

func (s *UpdateMmAppAndBindingRequestConversationConfig) GetOpenTts() *bool {
	return s.OpenTts
}

func (s *UpdateMmAppAndBindingRequestConversationConfig) GetStopOrRejectFlag() *bool {
	return s.StopOrRejectFlag
}

func (s *UpdateMmAppAndBindingRequestConversationConfig) GetTtsModel() *string {
	return s.TtsModel
}

func (s *UpdateMmAppAndBindingRequestConversationConfig) SetAsrModel(v string) *UpdateMmAppAndBindingRequestConversationConfig {
	s.AsrModel = &v
	return s
}

func (s *UpdateMmAppAndBindingRequestConversationConfig) SetOpenAsr(v bool) *UpdateMmAppAndBindingRequestConversationConfig {
	s.OpenAsr = &v
	return s
}

func (s *UpdateMmAppAndBindingRequestConversationConfig) SetOpenTts(v bool) *UpdateMmAppAndBindingRequestConversationConfig {
	s.OpenTts = &v
	return s
}

func (s *UpdateMmAppAndBindingRequestConversationConfig) SetStopOrRejectFlag(v bool) *UpdateMmAppAndBindingRequestConversationConfig {
	s.StopOrRejectFlag = &v
	return s
}

func (s *UpdateMmAppAndBindingRequestConversationConfig) SetTtsModel(v string) *UpdateMmAppAndBindingRequestConversationConfig {
	s.TtsModel = &v
	return s
}

func (s *UpdateMmAppAndBindingRequestConversationConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateMmAppAndBindingRequestMemoryConfig struct {
	Attributes []*UpdateMmAppAndBindingRequestMemoryConfigAttributes `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	Desc       *string                                               `json:"Desc,omitempty" xml:"Desc,omitempty"`
	Name       *string                                               `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateMmAppAndBindingRequestMemoryConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateMmAppAndBindingRequestMemoryConfig) GoString() string {
	return s.String()
}

func (s *UpdateMmAppAndBindingRequestMemoryConfig) GetAttributes() []*UpdateMmAppAndBindingRequestMemoryConfigAttributes {
	return s.Attributes
}

func (s *UpdateMmAppAndBindingRequestMemoryConfig) GetDesc() *string {
	return s.Desc
}

func (s *UpdateMmAppAndBindingRequestMemoryConfig) GetName() *string {
	return s.Name
}

func (s *UpdateMmAppAndBindingRequestMemoryConfig) SetAttributes(v []*UpdateMmAppAndBindingRequestMemoryConfigAttributes) *UpdateMmAppAndBindingRequestMemoryConfig {
	s.Attributes = v
	return s
}

func (s *UpdateMmAppAndBindingRequestMemoryConfig) SetDesc(v string) *UpdateMmAppAndBindingRequestMemoryConfig {
	s.Desc = &v
	return s
}

func (s *UpdateMmAppAndBindingRequestMemoryConfig) SetName(v string) *UpdateMmAppAndBindingRequestMemoryConfig {
	s.Name = &v
	return s
}

func (s *UpdateMmAppAndBindingRequestMemoryConfig) Validate() error {
	if s.Attributes != nil {
		for _, item := range s.Attributes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateMmAppAndBindingRequestMemoryConfigAttributes struct {
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateMmAppAndBindingRequestMemoryConfigAttributes) String() string {
	return dara.Prettify(s)
}

func (s UpdateMmAppAndBindingRequestMemoryConfigAttributes) GoString() string {
	return s.String()
}

func (s *UpdateMmAppAndBindingRequestMemoryConfigAttributes) GetDesc() *string {
	return s.Desc
}

func (s *UpdateMmAppAndBindingRequestMemoryConfigAttributes) GetName() *string {
	return s.Name
}

func (s *UpdateMmAppAndBindingRequestMemoryConfigAttributes) SetDesc(v string) *UpdateMmAppAndBindingRequestMemoryConfigAttributes {
	s.Desc = &v
	return s
}

func (s *UpdateMmAppAndBindingRequestMemoryConfigAttributes) SetName(v string) *UpdateMmAppAndBindingRequestMemoryConfigAttributes {
	s.Name = &v
	return s
}

func (s *UpdateMmAppAndBindingRequestMemoryConfigAttributes) Validate() error {
	return dara.Validate(s)
}

type UpdateMmAppAndBindingRequestModelConfig struct {
	EnableIntentRecognize *bool                                                      `json:"EnableIntentRecognize,omitempty" xml:"EnableIntentRecognize,omitempty"`
	EnableTransition      *bool                                                      `json:"EnableTransition,omitempty" xml:"EnableTransition,omitempty"`
	HistoryLimit          *int32                                                     `json:"HistoryLimit,omitempty" xml:"HistoryLimit,omitempty"`
	IntentOnlySwitch      *bool                                                      `json:"IntentOnlySwitch,omitempty" xml:"IntentOnlySwitch,omitempty"`
	ModelType             *string                                                    `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	OpenMemory            *bool                                                      `json:"OpenMemory,omitempty" xml:"OpenMemory,omitempty"`
	OpenWebSearch         *bool                                                      `json:"OpenWebSearch,omitempty" xml:"OpenWebSearch,omitempty"`
	SearchModel           *string                                                    `json:"SearchModel,omitempty" xml:"SearchModel,omitempty"`
	SearchStrategy        *string                                                    `json:"SearchStrategy,omitempty" xml:"SearchStrategy,omitempty"`
	TextModal             *string                                                    `json:"TextModal,omitempty" xml:"TextModal,omitempty"`
	UserPromptParams      []*UpdateMmAppAndBindingRequestModelConfigUserPromptParams `json:"UserPromptParams,omitempty" xml:"UserPromptParams,omitempty" type:"Repeated"`
	UserQueryParams       []*UpdateMmAppAndBindingRequestModelConfigUserQueryParams  `json:"userQueryParams,omitempty" xml:"userQueryParams,omitempty" type:"Repeated"`
}

func (s UpdateMmAppAndBindingRequestModelConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateMmAppAndBindingRequestModelConfig) GoString() string {
	return s.String()
}

func (s *UpdateMmAppAndBindingRequestModelConfig) GetEnableIntentRecognize() *bool {
	return s.EnableIntentRecognize
}

func (s *UpdateMmAppAndBindingRequestModelConfig) GetEnableTransition() *bool {
	return s.EnableTransition
}

func (s *UpdateMmAppAndBindingRequestModelConfig) GetHistoryLimit() *int32 {
	return s.HistoryLimit
}

func (s *UpdateMmAppAndBindingRequestModelConfig) GetIntentOnlySwitch() *bool {
	return s.IntentOnlySwitch
}

func (s *UpdateMmAppAndBindingRequestModelConfig) GetModelType() *string {
	return s.ModelType
}

func (s *UpdateMmAppAndBindingRequestModelConfig) GetOpenMemory() *bool {
	return s.OpenMemory
}

func (s *UpdateMmAppAndBindingRequestModelConfig) GetOpenWebSearch() *bool {
	return s.OpenWebSearch
}

func (s *UpdateMmAppAndBindingRequestModelConfig) GetSearchModel() *string {
	return s.SearchModel
}

func (s *UpdateMmAppAndBindingRequestModelConfig) GetSearchStrategy() *string {
	return s.SearchStrategy
}

func (s *UpdateMmAppAndBindingRequestModelConfig) GetTextModal() *string {
	return s.TextModal
}

func (s *UpdateMmAppAndBindingRequestModelConfig) GetUserPromptParams() []*UpdateMmAppAndBindingRequestModelConfigUserPromptParams {
	return s.UserPromptParams
}

func (s *UpdateMmAppAndBindingRequestModelConfig) GetUserQueryParams() []*UpdateMmAppAndBindingRequestModelConfigUserQueryParams {
	return s.UserQueryParams
}

func (s *UpdateMmAppAndBindingRequestModelConfig) SetEnableIntentRecognize(v bool) *UpdateMmAppAndBindingRequestModelConfig {
	s.EnableIntentRecognize = &v
	return s
}

func (s *UpdateMmAppAndBindingRequestModelConfig) SetEnableTransition(v bool) *UpdateMmAppAndBindingRequestModelConfig {
	s.EnableTransition = &v
	return s
}

func (s *UpdateMmAppAndBindingRequestModelConfig) SetHistoryLimit(v int32) *UpdateMmAppAndBindingRequestModelConfig {
	s.HistoryLimit = &v
	return s
}

func (s *UpdateMmAppAndBindingRequestModelConfig) SetIntentOnlySwitch(v bool) *UpdateMmAppAndBindingRequestModelConfig {
	s.IntentOnlySwitch = &v
	return s
}

func (s *UpdateMmAppAndBindingRequestModelConfig) SetModelType(v string) *UpdateMmAppAndBindingRequestModelConfig {
	s.ModelType = &v
	return s
}

func (s *UpdateMmAppAndBindingRequestModelConfig) SetOpenMemory(v bool) *UpdateMmAppAndBindingRequestModelConfig {
	s.OpenMemory = &v
	return s
}

func (s *UpdateMmAppAndBindingRequestModelConfig) SetOpenWebSearch(v bool) *UpdateMmAppAndBindingRequestModelConfig {
	s.OpenWebSearch = &v
	return s
}

func (s *UpdateMmAppAndBindingRequestModelConfig) SetSearchModel(v string) *UpdateMmAppAndBindingRequestModelConfig {
	s.SearchModel = &v
	return s
}

func (s *UpdateMmAppAndBindingRequestModelConfig) SetSearchStrategy(v string) *UpdateMmAppAndBindingRequestModelConfig {
	s.SearchStrategy = &v
	return s
}

func (s *UpdateMmAppAndBindingRequestModelConfig) SetTextModal(v string) *UpdateMmAppAndBindingRequestModelConfig {
	s.TextModal = &v
	return s
}

func (s *UpdateMmAppAndBindingRequestModelConfig) SetUserPromptParams(v []*UpdateMmAppAndBindingRequestModelConfigUserPromptParams) *UpdateMmAppAndBindingRequestModelConfig {
	s.UserPromptParams = v
	return s
}

func (s *UpdateMmAppAndBindingRequestModelConfig) SetUserQueryParams(v []*UpdateMmAppAndBindingRequestModelConfigUserQueryParams) *UpdateMmAppAndBindingRequestModelConfig {
	s.UserQueryParams = v
	return s
}

func (s *UpdateMmAppAndBindingRequestModelConfig) Validate() error {
	if s.UserPromptParams != nil {
		for _, item := range s.UserPromptParams {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.UserQueryParams != nil {
		for _, item := range s.UserQueryParams {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateMmAppAndBindingRequestModelConfigUserPromptParams struct {
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateMmAppAndBindingRequestModelConfigUserPromptParams) String() string {
	return dara.Prettify(s)
}

func (s UpdateMmAppAndBindingRequestModelConfigUserPromptParams) GoString() string {
	return s.String()
}

func (s *UpdateMmAppAndBindingRequestModelConfigUserPromptParams) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *UpdateMmAppAndBindingRequestModelConfigUserPromptParams) GetDescription() *string {
	return s.Description
}

func (s *UpdateMmAppAndBindingRequestModelConfigUserPromptParams) GetName() *string {
	return s.Name
}

func (s *UpdateMmAppAndBindingRequestModelConfigUserPromptParams) GetType() *string {
	return s.Type
}

func (s *UpdateMmAppAndBindingRequestModelConfigUserPromptParams) SetDefaultValue(v string) *UpdateMmAppAndBindingRequestModelConfigUserPromptParams {
	s.DefaultValue = &v
	return s
}

func (s *UpdateMmAppAndBindingRequestModelConfigUserPromptParams) SetDescription(v string) *UpdateMmAppAndBindingRequestModelConfigUserPromptParams {
	s.Description = &v
	return s
}

func (s *UpdateMmAppAndBindingRequestModelConfigUserPromptParams) SetName(v string) *UpdateMmAppAndBindingRequestModelConfigUserPromptParams {
	s.Name = &v
	return s
}

func (s *UpdateMmAppAndBindingRequestModelConfigUserPromptParams) SetType(v string) *UpdateMmAppAndBindingRequestModelConfigUserPromptParams {
	s.Type = &v
	return s
}

func (s *UpdateMmAppAndBindingRequestModelConfigUserPromptParams) Validate() error {
	return dara.Validate(s)
}

type UpdateMmAppAndBindingRequestModelConfigUserQueryParams struct {
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateMmAppAndBindingRequestModelConfigUserQueryParams) String() string {
	return dara.Prettify(s)
}

func (s UpdateMmAppAndBindingRequestModelConfigUserQueryParams) GoString() string {
	return s.String()
}

func (s *UpdateMmAppAndBindingRequestModelConfigUserQueryParams) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *UpdateMmAppAndBindingRequestModelConfigUserQueryParams) GetDescription() *string {
	return s.Description
}

func (s *UpdateMmAppAndBindingRequestModelConfigUserQueryParams) GetName() *string {
	return s.Name
}

func (s *UpdateMmAppAndBindingRequestModelConfigUserQueryParams) GetType() *string {
	return s.Type
}

func (s *UpdateMmAppAndBindingRequestModelConfigUserQueryParams) SetDefaultValue(v string) *UpdateMmAppAndBindingRequestModelConfigUserQueryParams {
	s.DefaultValue = &v
	return s
}

func (s *UpdateMmAppAndBindingRequestModelConfigUserQueryParams) SetDescription(v string) *UpdateMmAppAndBindingRequestModelConfigUserQueryParams {
	s.Description = &v
	return s
}

func (s *UpdateMmAppAndBindingRequestModelConfigUserQueryParams) SetName(v string) *UpdateMmAppAndBindingRequestModelConfigUserQueryParams {
	s.Name = &v
	return s
}

func (s *UpdateMmAppAndBindingRequestModelConfigUserQueryParams) SetType(v string) *UpdateMmAppAndBindingRequestModelConfigUserQueryParams {
	s.Type = &v
	return s
}

func (s *UpdateMmAppAndBindingRequestModelConfigUserQueryParams) Validate() error {
	return dara.Validate(s)
}
