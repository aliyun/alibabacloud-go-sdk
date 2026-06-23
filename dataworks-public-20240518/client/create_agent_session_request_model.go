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
	// The request ID provided by the client. This ID is returned in the response without modification.
	//
	// example:
	//
	// 4758330557805415712
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The JSON-RPC version. The value is fixed at `2.0`.
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
	// The extended metadata, which includes information such as agent binding, session source, and session tags.
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
	// The agent configuration for this session. The value must be one of the agents returned by the `ListAgents` API.
	Agent *CreateAgentSessionRequestParamsMetaAgent `json:"Agent,omitempty" xml:"Agent,omitempty" type:"Struct"`
	// The configuration parameters for the session, such as filters based on session source and session tags.
	Config               *CreateAgentSessionRequestParamsMetaConfig               `json:"Config,omitempty" xml:"Config,omitempty" type:"Struct"`
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
	// The agent name to bind to the session. This parameter is required.
	//
	// example:
	//
	// chat_cli_chatbi
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
	// The identifier for the session source. This allows you to search for sessions by their source. For example, if you use an agent on multiple pages, such as Page A and Page B, you can use this parameter to filter and display only the sessions created on Page A. The identifier can be up to 128 characters and can contain letters, digits, hyphens (-), and underscores (_).
	//
	// example:
	//
	// openapi_sdk
	SessionSource *string `json:"SessionSource,omitempty" xml:"SessionSource,omitempty"`
	// A list of session tags. You can use these tags to search and filter sessions.
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
	// The session tag. You can use session tags to filter sessions. For example, if your application calls the API with a fixed RAM sub-account but maintains its own user account system, you can pass a user\\"s account ID as a tag. This allows you to filter the session list by your internal account IDs. The tag can be up to 128 characters and can contain letters, digits, hyphens (-), and underscores (_).
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
	// example:
	//
	// chat，cli
	ExecutionLane *string `json:"ExecutionLane,omitempty" xml:"ExecutionLane,omitempty"`
	// example:
	//
	// yolo
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
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

func (s *CreateAgentSessionRequestParamsMetaInitialConfigOptions) SetExecutionLane(v string) *CreateAgentSessionRequestParamsMetaInitialConfigOptions {
	s.ExecutionLane = &v
	return s
}

func (s *CreateAgentSessionRequestParamsMetaInitialConfigOptions) SetMode(v string) *CreateAgentSessionRequestParamsMetaInitialConfigOptions {
	s.Mode = &v
	return s
}

func (s *CreateAgentSessionRequestParamsMetaInitialConfigOptions) Validate() error {
	return dara.Validate(s)
}
