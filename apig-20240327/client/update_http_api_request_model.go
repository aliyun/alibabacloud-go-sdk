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
	AgentProtocols []*string `json:"agentProtocols,omitempty" xml:"agentProtocols,omitempty" type:"Repeated"`
	// The AI protocols.
	AiProtocols []*string `json:"aiProtocols,omitempty" xml:"aiProtocols,omitempty" type:"Repeated"`
	// The authentication configuration.
	AuthConfig *AuthConfig `json:"authConfig,omitempty" xml:"authConfig,omitempty"`
	// The API base path, which must start with a forward slash (/).
	//
	// This parameter is required.
	//
	// example:
	//
	// /v1
	BasePath *string `json:"basePath,omitempty" xml:"basePath,omitempty"`
	// The deployment configurations.
	DeployConfigs []*HttpApiDeployConfig `json:"deployConfigs,omitempty" xml:"deployConfigs,omitempty" type:"Repeated"`
	// The API description.
	//
	// example:
	//
	// API for testing
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Specifies whether to enable authentication.
	EnableAuth       *bool  `json:"enableAuth,omitempty" xml:"enableAuth,omitempty"`
	FirstByteTimeout *int32 `json:"firstByteTimeout,omitempty" xml:"firstByteTimeout,omitempty"`
	// The HTTP Ingress API configurations.
	IngressConfig    *UpdateHttpApiRequestIngressConfig `json:"ingressConfig,omitempty" xml:"ingressConfig,omitempty" type:"Struct"`
	OnlyChangeConfig *bool                              `json:"onlyChangeConfig,omitempty" xml:"onlyChangeConfig,omitempty"`
	// The protocols that are used to access the API.
	Protocols               []*string `json:"protocols,omitempty" xml:"protocols,omitempty" type:"Repeated"`
	RemoveBasePathOnForward *bool     `json:"removeBasePathOnForward,omitempty" xml:"removeBasePathOnForward,omitempty"`
	// The versioning configurations.
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
	return dara.Validate(s)
}

type UpdateHttpApiRequestIngressConfig struct {
	// The environment ID.
	//
	// example:
	//
	// env-cr6ql0tlhtgmc****
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	// The Ingress class for listening.
	//
	// example:
	//
	// mse
	IngressClass *string `json:"ingressClass,omitempty" xml:"ingressClass,omitempty"`
	// Specifies whether to update the address in Ingress Status.
	//
	// example:
	//
	// false
	OverrideIngressIp *bool `json:"overrideIngressIp,omitempty" xml:"overrideIngressIp,omitempty"`
	// The source ID.
	//
	// example:
	//
	// src-crdddallhtgtr****
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The namespace for listening.
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
