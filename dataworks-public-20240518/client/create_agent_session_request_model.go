// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentSessionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *CreateAgentSessionRequest
	GetId() *string
	SetJsonrpc(v string) *CreateAgentSessionRequest
	GetJsonrpc() *string
	SetParams(v *CreateAgentSessionRequestParams) *CreateAgentSessionRequest
	GetParams() *CreateAgentSessionRequestParams
}

type CreateAgentSessionRequest struct {
	// The ID passed by the requester. The value is returned as is.
	//
	// example:
	//
	// 4758330557805415712
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The JSON-RPC version. This value is always 2.0.
	//
	// example:
	//
	// 2.0
	Jsonrpc *string `json:"Jsonrpc,omitempty" xml:"Jsonrpc,omitempty"`
	// The business parameters.
	Params *CreateAgentSessionRequestParams `json:"Params,omitempty" xml:"Params,omitempty" type:"Struct"`
}

func (s CreateAgentSessionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentSessionRequest) GoString() string {
	return s.String()
}

func (s *CreateAgentSessionRequest) GetId() *string {
	return s.Id
}

func (s *CreateAgentSessionRequest) GetJsonrpc() *string {
	return s.Jsonrpc
}

func (s *CreateAgentSessionRequest) GetParams() *CreateAgentSessionRequestParams {
	return s.Params
}

func (s *CreateAgentSessionRequest) SetId(v string) *CreateAgentSessionRequest {
	s.Id = &v
	return s
}

func (s *CreateAgentSessionRequest) SetJsonrpc(v string) *CreateAgentSessionRequest {
	s.Jsonrpc = &v
	return s
}

func (s *CreateAgentSessionRequest) SetParams(v *CreateAgentSessionRequestParams) *CreateAgentSessionRequest {
	s.Params = v
	return s
}

func (s *CreateAgentSessionRequest) Validate() error {
	if s.Params != nil {
		if err := s.Params.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAgentSessionRequestParams struct {
	// The extended metadata that carries agent binding, session source, and tags.
	Meta *CreateAgentSessionRequestParamsMeta `json:"Meta,omitempty" xml:"Meta,omitempty" type:"Struct"`
}

func (s CreateAgentSessionRequestParams) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentSessionRequestParams) GoString() string {
	return s.String()
}

func (s *CreateAgentSessionRequestParams) GetMeta() *CreateAgentSessionRequestParamsMeta {
	return s.Meta
}

func (s *CreateAgentSessionRequestParams) SetMeta(v *CreateAgentSessionRequestParamsMeta) *CreateAgentSessionRequestParams {
	s.Meta = v
	return s
}

func (s *CreateAgentSessionRequestParams) Validate() error {
	if s.Meta != nil {
		if err := s.Meta.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAgentSessionRequestParamsMeta struct {
	// The agent configuration for the session. Valid values are the results returned by the ListAgents operation.
	Agent *CreateAgentSessionRequestParamsMetaAgent `json:"Agent,omitempty" xml:"Agent,omitempty" type:"Struct"`
	// The session parameter settings, such as filter parameter settings based on the session source and session tags.
	Config *CreateAgentSessionRequestParamsMetaConfig `json:"Config,omitempty" xml:"Config,omitempty" type:"Struct"`
	// The advanced parameter settings for the agent execution environment.
	InitialConfigOptions *CreateAgentSessionRequestParamsMetaInitialConfigOptions `json:"InitialConfigOptions,omitempty" xml:"InitialConfigOptions,omitempty" type:"Struct"`
}

func (s CreateAgentSessionRequestParamsMeta) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentSessionRequestParamsMeta) GoString() string {
	return s.String()
}

func (s *CreateAgentSessionRequestParamsMeta) GetAgent() *CreateAgentSessionRequestParamsMetaAgent {
	return s.Agent
}

func (s *CreateAgentSessionRequestParamsMeta) GetConfig() *CreateAgentSessionRequestParamsMetaConfig {
	return s.Config
}

func (s *CreateAgentSessionRequestParamsMeta) GetInitialConfigOptions() *CreateAgentSessionRequestParamsMetaInitialConfigOptions {
	return s.InitialConfigOptions
}

func (s *CreateAgentSessionRequestParamsMeta) SetAgent(v *CreateAgentSessionRequestParamsMetaAgent) *CreateAgentSessionRequestParamsMeta {
	s.Agent = v
	return s
}

func (s *CreateAgentSessionRequestParamsMeta) SetConfig(v *CreateAgentSessionRequestParamsMetaConfig) *CreateAgentSessionRequestParamsMeta {
	s.Config = v
	return s
}

func (s *CreateAgentSessionRequestParamsMeta) SetInitialConfigOptions(v *CreateAgentSessionRequestParamsMetaInitialConfigOptions) *CreateAgentSessionRequestParamsMeta {
	s.InitialConfigOptions = v
	return s
}

func (s *CreateAgentSessionRequestParamsMeta) Validate() error {
	if s.Agent != nil {
		if err := s.Agent.Validate(); err != nil {
			return err
		}
	}
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	if s.InitialConfigOptions != nil {
		if err := s.InitialConfigOptions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAgentSessionRequestParamsMetaAgent struct {
	// The name of the agent bound to the session. This parameter is required. Valid values:
	//
	// 	- dataworks_data_agent: DataWorks built-in agent - Data Agent. It provides AI capabilities for the entire data development lifecycle, including data integration, development, O&M, governance, and analysis.
	//
	// 	- dataworks_chatbi_agent: DataWorks built-in agent - ChatBI. Based on natural language processing and intelligent analysis technologies, it automates the entire analysis process from requirement parsing, data extraction, and automatic code generation to visual report output through conversational interaction.
	//
	// 	- dataworks_ai_assistant_agent: DataWorks built-in agent - AI Assistant. An enterprise-grade AI assistant for DataWorks built on open-source frameworks such as OpenClaw and Hermes Agent.
	//
	// example:
	//
	// dataworks_data_agent
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
}

func (s CreateAgentSessionRequestParamsMetaAgent) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentSessionRequestParamsMetaAgent) GoString() string {
	return s.String()
}

func (s *CreateAgentSessionRequestParamsMetaAgent) GetAgentName() *string {
	return s.AgentName
}

func (s *CreateAgentSessionRequestParamsMetaAgent) SetAgentName(v string) *CreateAgentSessionRequestParamsMetaAgent {
	s.AgentName = &v
	return s
}

func (s *CreateAgentSessionRequestParamsMetaAgent) Validate() error {
	return dara.Validate(s)
}

type CreateAgentSessionRequestParamsMetaConfig struct {
	// The session source identifier, which facilitates retrieval by source. For example, if an agent is used on both Page A and Page B, and you want to display only sessions created on Page A, you can filter sessions based on this parameter. The value can be up to 128 characters in length and can contain letters, digits, hyphens (-), and underscores (_).
	//
	// example:
	//
	// openapi_sdk
	SessionSource *string `json:"SessionSource,omitempty" xml:"SessionSource,omitempty"`
	// The list of session tags. You can search and filter sessions based on session tags.
	SessionTags []*CreateAgentSessionRequestParamsMetaConfigSessionTags `json:"SessionTags,omitempty" xml:"SessionTags,omitempty" type:"Repeated"`
}

func (s CreateAgentSessionRequestParamsMetaConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentSessionRequestParamsMetaConfig) GoString() string {
	return s.String()
}

func (s *CreateAgentSessionRequestParamsMetaConfig) GetSessionSource() *string {
	return s.SessionSource
}

func (s *CreateAgentSessionRequestParamsMetaConfig) GetSessionTags() []*CreateAgentSessionRequestParamsMetaConfigSessionTags {
	return s.SessionTags
}

func (s *CreateAgentSessionRequestParamsMetaConfig) SetSessionSource(v string) *CreateAgentSessionRequestParamsMetaConfig {
	s.SessionSource = &v
	return s
}

func (s *CreateAgentSessionRequestParamsMetaConfig) SetSessionTags(v []*CreateAgentSessionRequestParamsMetaConfigSessionTags) *CreateAgentSessionRequestParamsMetaConfig {
	s.SessionTags = v
	return s
}

func (s *CreateAgentSessionRequestParamsMetaConfig) Validate() error {
	if s.SessionTags != nil {
		for _, item := range s.SessionTags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateAgentSessionRequestParamsMetaConfigSessionTags struct {
	// The session tag. You can filter sessions based on session tags. For example, if a fixed RAM user is used to call an API operation but the calling system has its own account system, you need to filter the session list based on the account ID. In this case, you can pass the account ID of the calling system as the session tag. The value can be up to 128 characters in length and can contain letters, digits, hyphens (-), and underscores (_).
	//
	// example:
	//
	// chatbi
	SessionTagCode *string `json:"SessionTagCode,omitempty" xml:"SessionTagCode,omitempty"`
}

func (s CreateAgentSessionRequestParamsMetaConfigSessionTags) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentSessionRequestParamsMetaConfigSessionTags) GoString() string {
	return s.String()
}

func (s *CreateAgentSessionRequestParamsMetaConfigSessionTags) GetSessionTagCode() *string {
	return s.SessionTagCode
}

func (s *CreateAgentSessionRequestParamsMetaConfigSessionTags) SetSessionTagCode(v string) *CreateAgentSessionRequestParamsMetaConfigSessionTags {
	s.SessionTagCode = &v
	return s
}

func (s *CreateAgentSessionRequestParamsMetaConfigSessionTags) Validate() error {
	return dara.Validate(s)
}

type CreateAgentSessionRequestParamsMetaInitialConfigOptions struct {
	// The execution mode. Valid values:
	//
	// 	- chat: conversation mode only. This mode is suitable for simple scenarios such as data querying and Q&A. It is fast and consumes fewer tokens, but cannot handle complex problems.
	//
	// 	- cli: sandbox mode. This mode is suitable for complex scenarios such as data analytics, data processing, and code writing. It can handle complex problems and allows the model to independently analyze and process problems. However, it is slower and consumes more tokens than conversation mode.
	//
	// example:
	//
	// chat，cli
	ExecutionLane *string `json:"ExecutionLane,omitempty" xml:"ExecutionLane,omitempty"`
	// The authorization mode for script execution. OpenAPI currently supports only the yolo mode. Valid values:
	//
	// 	- yolo: automatic authorization. No manual intervention is required, and the model can process tasks automatically.
	//
	// example:
	//
	// yolo
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The model ID. Use the full model identifier supported by the target environment. If this parameter is omitted, the default model is used. If this parameter is specified, the server verifies its availability and applies it before the session is created.
	//
	// example:
	//
	// qwen3.7-max(openai)
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// The DataWorks workspace ID. This parameter is used to initialize the project context of the session. If this parameter is omitted, the session is processed without a project context. You can subsequently correct or switch the project context by using PromptAgentSession.
	//
	// example:
	//
	// 12345
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The reasoning effort configuration. The value must be supported by the target model. A value of none disables reasoning, and a value of default uses the default configuration of the model. If an unsupported value is specified or the configuration application has failed, the session creation is marked as failed. If this parameter is omitted, the existing reasoning effort configuration is not modified. A successful configuration updates the default value of the daemon workspace and affects not only the current request.
	//
	// example:
	//
	// none
	ReasoningEffort *string `json:"ReasoningEffort,omitempty" xml:"ReasoningEffort,omitempty"`
	// The ID of the resource group used for initialization.
	//
	// example:
	//
	// Serverless_res_group_123123123_564657857
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The names of the custom skills to be loaded. Separate multiple names with commas (,).
	//
	// example:
	//
	// skill-name-1,skill-name2,skill-name-3
	Skills *string `json:"Skills,omitempty" xml:"Skills,omitempty"`
}

func (s CreateAgentSessionRequestParamsMetaInitialConfigOptions) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentSessionRequestParamsMetaInitialConfigOptions) GoString() string {
	return s.String()
}

func (s *CreateAgentSessionRequestParamsMetaInitialConfigOptions) GetExecutionLane() *string {
	return s.ExecutionLane
}

func (s *CreateAgentSessionRequestParamsMetaInitialConfigOptions) GetMode() *string {
	return s.Mode
}

func (s *CreateAgentSessionRequestParamsMetaInitialConfigOptions) GetModel() *string {
	return s.Model
}

func (s *CreateAgentSessionRequestParamsMetaInitialConfigOptions) GetProjectId() *string {
	return s.ProjectId
}

func (s *CreateAgentSessionRequestParamsMetaInitialConfigOptions) GetReasoningEffort() *string {
	return s.ReasoningEffort
}

func (s *CreateAgentSessionRequestParamsMetaInitialConfigOptions) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateAgentSessionRequestParamsMetaInitialConfigOptions) GetSkills() *string {
	return s.Skills
}

func (s *CreateAgentSessionRequestParamsMetaInitialConfigOptions) SetExecutionLane(v string) *CreateAgentSessionRequestParamsMetaInitialConfigOptions {
	s.ExecutionLane = &v
	return s
}

func (s *CreateAgentSessionRequestParamsMetaInitialConfigOptions) SetMode(v string) *CreateAgentSessionRequestParamsMetaInitialConfigOptions {
	s.Mode = &v
	return s
}

func (s *CreateAgentSessionRequestParamsMetaInitialConfigOptions) SetModel(v string) *CreateAgentSessionRequestParamsMetaInitialConfigOptions {
	s.Model = &v
	return s
}

func (s *CreateAgentSessionRequestParamsMetaInitialConfigOptions) SetProjectId(v string) *CreateAgentSessionRequestParamsMetaInitialConfigOptions {
	s.ProjectId = &v
	return s
}

func (s *CreateAgentSessionRequestParamsMetaInitialConfigOptions) SetReasoningEffort(v string) *CreateAgentSessionRequestParamsMetaInitialConfigOptions {
	s.ReasoningEffort = &v
	return s
}

func (s *CreateAgentSessionRequestParamsMetaInitialConfigOptions) SetResourceGroupId(v string) *CreateAgentSessionRequestParamsMetaInitialConfigOptions {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateAgentSessionRequestParamsMetaInitialConfigOptions) SetSkills(v string) *CreateAgentSessionRequestParamsMetaInitialConfigOptions {
	s.Skills = &v
	return s
}

func (s *CreateAgentSessionRequestParamsMetaInitialConfigOptions) Validate() error {
	return dara.Validate(s)
}
