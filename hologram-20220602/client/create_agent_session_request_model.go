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
	// example:
	//
	// 2072736942627512345
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 2.0
	Jsonrpc *string                          `json:"Jsonrpc,omitempty" xml:"Jsonrpc,omitempty"`
	Params  *CreateAgentSessionRequestParams `json:"Params,omitempty" xml:"Params,omitempty" type:"Struct"`
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
	Agent                *CreateAgentSessionRequestParamsMetaAgent                `json:"Agent,omitempty" xml:"Agent,omitempty" type:"Struct"`
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
	// example:
	//
	// hologres
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
	// example:
	//
	// CHAT_HOLOGRES
	SessionSource *string                                                 `json:"SessionSource,omitempty" xml:"SessionSource,omitempty"`
	SessionTags   []*CreateAgentSessionRequestParamsMetaConfigSessionTags `json:"SessionTags,omitempty" xml:"SessionTags,omitempty" type:"Repeated"`
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
	// example:
	//
	// holo
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
	// cli
	ExecutionLane *string `json:"ExecutionLane,omitempty" xml:"ExecutionLane,omitempty"`
	// example:
	//
	// yolo
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// example:
	//
	// rg-acfmvqsnvkfr2sa
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
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
