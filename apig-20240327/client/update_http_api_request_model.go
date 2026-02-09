// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHttpApiRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentProtocols(v []*string) *UpdateHttpApiRequest
	GetAgentProtocols() []*string
	SetAiProtocols(v []*string) *UpdateHttpApiRequest
	GetAiProtocols() []*string
	SetAuthConfig(v *AuthConfig) *UpdateHttpApiRequest
	GetAuthConfig() *AuthConfig
	SetBasePath(v string) *UpdateHttpApiRequest
	GetBasePath() *string
	SetDeployConfigs(v []*HttpApiDeployConfig) *UpdateHttpApiRequest
	GetDeployConfigs() []*HttpApiDeployConfig
	SetDescription(v string) *UpdateHttpApiRequest
	GetDescription() *string
	SetEnableAuth(v bool) *UpdateHttpApiRequest
	GetEnableAuth() *bool
	SetFirstByteTimeout(v int32) *UpdateHttpApiRequest
	GetFirstByteTimeout() *int32
	SetIngressConfig(v *UpdateHttpApiRequestIngressConfig) *UpdateHttpApiRequest
	GetIngressConfig() *UpdateHttpApiRequestIngressConfig
	SetOnlyChangeConfig(v bool) *UpdateHttpApiRequest
	GetOnlyChangeConfig() *bool
	SetProtocols(v []*string) *UpdateHttpApiRequest
	GetProtocols() []*string
	SetRemoveBasePathOnForward(v bool) *UpdateHttpApiRequest
	GetRemoveBasePathOnForward() *bool
	SetVersionConfig(v *HttpApiVersionConfig) *UpdateHttpApiRequest
	GetVersionConfig() *HttpApiVersionConfig
}

type UpdateHttpApiRequest struct {
	// The list of agent protocols
	AgentProtocols []*string `json:"agentProtocols,omitempty" xml:"agentProtocols,omitempty" type:"Repeated"`
	// The status code.
	AiProtocols []*string `json:"aiProtocols,omitempty" xml:"aiProtocols,omitempty" type:"Repeated"`
	// The authentication configuration
	AuthConfig *AuthConfig `json:"authConfig,omitempty" xml:"authConfig,omitempty"`
	// The list of API deployment configurations.
	//
	// This parameter is required.
	//
	// example:
	//
	// /v1
	BasePath *string `json:"basePath,omitempty" xml:"basePath,omitempty"`
	// The deployment configurations
	DeployConfigs []*HttpApiDeployConfig `json:"deployConfigs,omitempty" xml:"deployConfigs,omitempty" type:"Repeated"`
	// The environment ID.
	//
	// example:
	//
	// The source ID.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Whether authentication is enabled
	//
	// example:
	//
	// true
	EnableAuth *bool `json:"enableAuth,omitempty" xml:"enableAuth,omitempty"`
	// The first byte timeout in nanoseconds
	//
	// example:
	//
	// 30s
	FirstByteTimeout *int32 `json:"firstByteTimeout,omitempty" xml:"firstByteTimeout,omitempty"`
	// Specifies whether to enable authentication.
	IngressConfig *UpdateHttpApiRequestIngressConfig `json:"ingressConfig,omitempty" xml:"ingressConfig,omitempty" type:"Struct"`
	// Whether to only change configuration without redeployment
	//
	// example:
	//
	// true
	OnlyChangeConfig *bool `json:"onlyChangeConfig,omitempty" xml:"onlyChangeConfig,omitempty"`
	// The listened namespace.
	Protocols []*string `json:"protocols,omitempty" xml:"protocols,omitempty" type:"Repeated"`
	// Whether to remove base path when forwarding
	//
	// example:
	//
	// true
	RemoveBasePathOnForward *bool `json:"removeBasePathOnForward,omitempty" xml:"removeBasePathOnForward,omitempty"`
	// A deployment configuration.
	VersionConfig *HttpApiVersionConfig `json:"versionConfig,omitempty" xml:"versionConfig,omitempty"`
}

func (s UpdateHttpApiRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpApiRequest) GoString() string {
	return s.String()
}

func (s *UpdateHttpApiRequest) GetAgentProtocols() []*string {
	return s.AgentProtocols
}

func (s *UpdateHttpApiRequest) GetAiProtocols() []*string {
	return s.AiProtocols
}

func (s *UpdateHttpApiRequest) GetAuthConfig() *AuthConfig {
	return s.AuthConfig
}

func (s *UpdateHttpApiRequest) GetBasePath() *string {
	return s.BasePath
}

func (s *UpdateHttpApiRequest) GetDeployConfigs() []*HttpApiDeployConfig {
	return s.DeployConfigs
}

func (s *UpdateHttpApiRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateHttpApiRequest) GetEnableAuth() *bool {
	return s.EnableAuth
}

func (s *UpdateHttpApiRequest) GetFirstByteTimeout() *int32 {
	return s.FirstByteTimeout
}

func (s *UpdateHttpApiRequest) GetIngressConfig() *UpdateHttpApiRequestIngressConfig {
	return s.IngressConfig
}

func (s *UpdateHttpApiRequest) GetOnlyChangeConfig() *bool {
	return s.OnlyChangeConfig
}

func (s *UpdateHttpApiRequest) GetProtocols() []*string {
	return s.Protocols
}

func (s *UpdateHttpApiRequest) GetRemoveBasePathOnForward() *bool {
	return s.RemoveBasePathOnForward
}

func (s *UpdateHttpApiRequest) GetVersionConfig() *HttpApiVersionConfig {
	return s.VersionConfig
}

func (s *UpdateHttpApiRequest) SetAgentProtocols(v []*string) *UpdateHttpApiRequest {
	s.AgentProtocols = v
	return s
}

func (s *UpdateHttpApiRequest) SetAiProtocols(v []*string) *UpdateHttpApiRequest {
	s.AiProtocols = v
	return s
}

func (s *UpdateHttpApiRequest) SetAuthConfig(v *AuthConfig) *UpdateHttpApiRequest {
	s.AuthConfig = v
	return s
}

func (s *UpdateHttpApiRequest) SetBasePath(v string) *UpdateHttpApiRequest {
	s.BasePath = &v
	return s
}

func (s *UpdateHttpApiRequest) SetDeployConfigs(v []*HttpApiDeployConfig) *UpdateHttpApiRequest {
	s.DeployConfigs = v
	return s
}

func (s *UpdateHttpApiRequest) SetDescription(v string) *UpdateHttpApiRequest {
	s.Description = &v
	return s
}

func (s *UpdateHttpApiRequest) SetEnableAuth(v bool) *UpdateHttpApiRequest {
	s.EnableAuth = &v
	return s
}

func (s *UpdateHttpApiRequest) SetFirstByteTimeout(v int32) *UpdateHttpApiRequest {
	s.FirstByteTimeout = &v
	return s
}

func (s *UpdateHttpApiRequest) SetIngressConfig(v *UpdateHttpApiRequestIngressConfig) *UpdateHttpApiRequest {
	s.IngressConfig = v
	return s
}

func (s *UpdateHttpApiRequest) SetOnlyChangeConfig(v bool) *UpdateHttpApiRequest {
	s.OnlyChangeConfig = &v
	return s
}

func (s *UpdateHttpApiRequest) SetProtocols(v []*string) *UpdateHttpApiRequest {
	s.Protocols = v
	return s
}

func (s *UpdateHttpApiRequest) SetRemoveBasePathOnForward(v bool) *UpdateHttpApiRequest {
	s.RemoveBasePathOnForward = &v
	return s
}

func (s *UpdateHttpApiRequest) SetVersionConfig(v *HttpApiVersionConfig) *UpdateHttpApiRequest {
	s.VersionConfig = v
	return s
}

func (s *UpdateHttpApiRequest) Validate() error {
	if s.AuthConfig != nil {
		if err := s.AuthConfig.Validate(); err != nil {
			return err
		}
	}
	if s.DeployConfigs != nil {
		for _, item := range s.DeployConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.IngressConfig != nil {
		if err := s.IngressConfig.Validate(); err != nil {
			return err
		}
	}
	if s.VersionConfig != nil {
		if err := s.VersionConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateHttpApiRequestIngressConfig struct {
	// The authentication configuration.
	//
	// example:
	//
	// env-cr6ql0tlhtgmc****
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	// The response parameters.
	//
	// example:
	//
	// mse
	IngressClass *string `json:"ingressClass,omitempty" xml:"ingressClass,omitempty"`
	// The returned message.
	//
	// example:
	//
	// false
	OverrideIngressIp *bool `json:"overrideIngressIp,omitempty" xml:"overrideIngressIp,omitempty"`
	// json
	//
	// example:
	//
	// src-crdddallhtgtr****
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// default
	WatchNamespace *string `json:"watchNamespace,omitempty" xml:"watchNamespace,omitempty"`
}

func (s UpdateHttpApiRequestIngressConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpApiRequestIngressConfig) GoString() string {
	return s.String()
}

func (s *UpdateHttpApiRequestIngressConfig) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *UpdateHttpApiRequestIngressConfig) GetIngressClass() *string {
	return s.IngressClass
}

func (s *UpdateHttpApiRequestIngressConfig) GetOverrideIngressIp() *bool {
	return s.OverrideIngressIp
}

func (s *UpdateHttpApiRequestIngressConfig) GetSourceId() *string {
	return s.SourceId
}

func (s *UpdateHttpApiRequestIngressConfig) GetWatchNamespace() *string {
	return s.WatchNamespace
}

func (s *UpdateHttpApiRequestIngressConfig) SetEnvironmentId(v string) *UpdateHttpApiRequestIngressConfig {
	s.EnvironmentId = &v
	return s
}

func (s *UpdateHttpApiRequestIngressConfig) SetIngressClass(v string) *UpdateHttpApiRequestIngressConfig {
	s.IngressClass = &v
	return s
}

func (s *UpdateHttpApiRequestIngressConfig) SetOverrideIngressIp(v bool) *UpdateHttpApiRequestIngressConfig {
	s.OverrideIngressIp = &v
	return s
}

func (s *UpdateHttpApiRequestIngressConfig) SetSourceId(v string) *UpdateHttpApiRequestIngressConfig {
	s.SourceId = &v
	return s
}

func (s *UpdateHttpApiRequestIngressConfig) SetWatchNamespace(v string) *UpdateHttpApiRequestIngressConfig {
	s.WatchNamespace = &v
	return s
}

func (s *UpdateHttpApiRequestIngressConfig) Validate() error {
	return dara.Validate(s)
}
