// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHttpApiRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentProtocols(v []*string) *CreateHttpApiRequest
	GetAgentProtocols() []*string
	SetAiProtocols(v []*string) *CreateHttpApiRequest
	GetAiProtocols() []*string
	SetAuthConfig(v *AuthConfig) *CreateHttpApiRequest
	GetAuthConfig() *AuthConfig
	SetBasePath(v string) *CreateHttpApiRequest
	GetBasePath() *string
	SetDeployConfigs(v []*HttpApiDeployConfig) *CreateHttpApiRequest
	GetDeployConfigs() []*HttpApiDeployConfig
	SetDescription(v string) *CreateHttpApiRequest
	GetDescription() *string
	SetEnableAuth(v bool) *CreateHttpApiRequest
	GetEnableAuth() *bool
	SetIngressConfig(v *CreateHttpApiRequestIngressConfig) *CreateHttpApiRequest
	GetIngressConfig() *CreateHttpApiRequestIngressConfig
	SetModelCategory(v string) *CreateHttpApiRequest
	GetModelCategory() *string
	SetName(v string) *CreateHttpApiRequest
	GetName() *string
	SetProtocols(v []*string) *CreateHttpApiRequest
	GetProtocols() []*string
	SetRemoveBasePathOnForward(v bool) *CreateHttpApiRequest
	GetRemoveBasePathOnForward() *bool
	SetResourceGroupId(v string) *CreateHttpApiRequest
	GetResourceGroupId() *string
	SetType(v string) *CreateHttpApiRequest
	GetType() *string
	SetVersionConfig(v *HttpApiVersionConfig) *CreateHttpApiRequest
	GetVersionConfig() *HttpApiVersionConfig
}

type CreateHttpApiRequest struct {
	AgentProtocols []*string `json:"agentProtocols,omitempty" xml:"agentProtocols,omitempty" type:"Repeated"`
	// The AI API protocols. Valid value:
	//
	// 	- OpenAI/v1
	AiProtocols []*string `json:"aiProtocols,omitempty" xml:"aiProtocols,omitempty" type:"Repeated"`
	// The authentication configurations.
	AuthConfig *AuthConfig `json:"authConfig,omitempty" xml:"authConfig,omitempty"`
	// The API base path, which must start with a forward slash (/).
	//
	// example:
	//
	// /v1
	BasePath *string `json:"basePath,omitempty" xml:"basePath,omitempty"`
	// The API deployment configurations. Currently, only AI APIs support deployment configurations, and only a single deployment configuration can be passed.
	DeployConfigs []*HttpApiDeployConfig `json:"deployConfigs,omitempty" xml:"deployConfigs,omitempty" type:"Repeated"`
	// The API description.
	//
	// example:
	//
	// API for testing
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Specifies whether to enable authentication.
	EnableAuth *bool `json:"enableAuth,omitempty" xml:"enableAuth,omitempty"`
	// The HTTP Ingress configurations.
	IngressConfig *CreateHttpApiRequestIngressConfig `json:"ingressConfig,omitempty" xml:"ingressConfig,omitempty" type:"Struct"`
	ModelCategory *string                            `json:"modelCategory,omitempty" xml:"modelCategory,omitempty"`
	// The API name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-api
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The protocols that are used to call the API.
	Protocols               []*string `json:"protocols,omitempty" xml:"protocols,omitempty" type:"Repeated"`
	RemoveBasePathOnForward *bool     `json:"removeBasePathOnForward,omitempty" xml:"removeBasePathOnForward,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aekzgvmlotionbi
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The API type. Valid values:
	//
	// 	- Http
	//
	// 	- Rest
	//
	// 	- WebSocket
	//
	// 	- HttpIngress
	//
	// example:
	//
	// Http
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The versioning configuration of the API.
	VersionConfig *HttpApiVersionConfig `json:"versionConfig,omitempty" xml:"versionConfig,omitempty"`
}

func (s CreateHttpApiRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateHttpApiRequest) GoString() string {
	return s.String()
}

func (s *CreateHttpApiRequest) GetAgentProtocols() []*string {
	return s.AgentProtocols
}

func (s *CreateHttpApiRequest) GetAiProtocols() []*string {
	return s.AiProtocols
}

func (s *CreateHttpApiRequest) GetAuthConfig() *AuthConfig {
	return s.AuthConfig
}

func (s *CreateHttpApiRequest) GetBasePath() *string {
	return s.BasePath
}

func (s *CreateHttpApiRequest) GetDeployConfigs() []*HttpApiDeployConfig {
	return s.DeployConfigs
}

func (s *CreateHttpApiRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateHttpApiRequest) GetEnableAuth() *bool {
	return s.EnableAuth
}

func (s *CreateHttpApiRequest) GetIngressConfig() *CreateHttpApiRequestIngressConfig {
	return s.IngressConfig
}

func (s *CreateHttpApiRequest) GetModelCategory() *string {
	return s.ModelCategory
}

func (s *CreateHttpApiRequest) GetName() *string {
	return s.Name
}

func (s *CreateHttpApiRequest) GetProtocols() []*string {
	return s.Protocols
}

func (s *CreateHttpApiRequest) GetRemoveBasePathOnForward() *bool {
	return s.RemoveBasePathOnForward
}

func (s *CreateHttpApiRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateHttpApiRequest) GetType() *string {
	return s.Type
}

func (s *CreateHttpApiRequest) GetVersionConfig() *HttpApiVersionConfig {
	return s.VersionConfig
}

func (s *CreateHttpApiRequest) SetAgentProtocols(v []*string) *CreateHttpApiRequest {
	s.AgentProtocols = v
	return s
}

func (s *CreateHttpApiRequest) SetAiProtocols(v []*string) *CreateHttpApiRequest {
	s.AiProtocols = v
	return s
}

func (s *CreateHttpApiRequest) SetAuthConfig(v *AuthConfig) *CreateHttpApiRequest {
	s.AuthConfig = v
	return s
}

func (s *CreateHttpApiRequest) SetBasePath(v string) *CreateHttpApiRequest {
	s.BasePath = &v
	return s
}

func (s *CreateHttpApiRequest) SetDeployConfigs(v []*HttpApiDeployConfig) *CreateHttpApiRequest {
	s.DeployConfigs = v
	return s
}

func (s *CreateHttpApiRequest) SetDescription(v string) *CreateHttpApiRequest {
	s.Description = &v
	return s
}

func (s *CreateHttpApiRequest) SetEnableAuth(v bool) *CreateHttpApiRequest {
	s.EnableAuth = &v
	return s
}

func (s *CreateHttpApiRequest) SetIngressConfig(v *CreateHttpApiRequestIngressConfig) *CreateHttpApiRequest {
	s.IngressConfig = v
	return s
}

func (s *CreateHttpApiRequest) SetModelCategory(v string) *CreateHttpApiRequest {
	s.ModelCategory = &v
	return s
}

func (s *CreateHttpApiRequest) SetName(v string) *CreateHttpApiRequest {
	s.Name = &v
	return s
}

func (s *CreateHttpApiRequest) SetProtocols(v []*string) *CreateHttpApiRequest {
	s.Protocols = v
	return s
}

func (s *CreateHttpApiRequest) SetRemoveBasePathOnForward(v bool) *CreateHttpApiRequest {
	s.RemoveBasePathOnForward = &v
	return s
}

func (s *CreateHttpApiRequest) SetResourceGroupId(v string) *CreateHttpApiRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateHttpApiRequest) SetType(v string) *CreateHttpApiRequest {
	s.Type = &v
	return s
}

func (s *CreateHttpApiRequest) SetVersionConfig(v *HttpApiVersionConfig) *CreateHttpApiRequest {
	s.VersionConfig = v
	return s
}

func (s *CreateHttpApiRequest) Validate() error {
	return dara.Validate(s)
}

type CreateHttpApiRequestIngressConfig struct {
	ClusterId *string `json:"clusterId,omitempty" xml:"clusterId,omitempty"`
	// The environment ID.
	//
	// example:
	//
	// env-cq146allhtgk***
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	// The Ingress Class for listening.
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
	// Deprecated
	//
	// The source ID.
	//
	// example:
	//
	// src-crdddallhtgtr***
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The namespace for listening.
	//
	// example:
	//
	// default
	WatchNamespace *string `json:"watchNamespace,omitempty" xml:"watchNamespace,omitempty"`
}

func (s CreateHttpApiRequestIngressConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateHttpApiRequestIngressConfig) GoString() string {
	return s.String()
}

func (s *CreateHttpApiRequestIngressConfig) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateHttpApiRequestIngressConfig) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *CreateHttpApiRequestIngressConfig) GetIngressClass() *string {
	return s.IngressClass
}

func (s *CreateHttpApiRequestIngressConfig) GetOverrideIngressIp() *bool {
	return s.OverrideIngressIp
}

func (s *CreateHttpApiRequestIngressConfig) GetSourceId() *string {
	return s.SourceId
}

func (s *CreateHttpApiRequestIngressConfig) GetWatchNamespace() *string {
	return s.WatchNamespace
}

func (s *CreateHttpApiRequestIngressConfig) SetClusterId(v string) *CreateHttpApiRequestIngressConfig {
	s.ClusterId = &v
	return s
}

func (s *CreateHttpApiRequestIngressConfig) SetEnvironmentId(v string) *CreateHttpApiRequestIngressConfig {
	s.EnvironmentId = &v
	return s
}

func (s *CreateHttpApiRequestIngressConfig) SetIngressClass(v string) *CreateHttpApiRequestIngressConfig {
	s.IngressClass = &v
	return s
}

func (s *CreateHttpApiRequestIngressConfig) SetOverrideIngressIp(v bool) *CreateHttpApiRequestIngressConfig {
	s.OverrideIngressIp = &v
	return s
}

func (s *CreateHttpApiRequestIngressConfig) SetSourceId(v string) *CreateHttpApiRequestIngressConfig {
	s.SourceId = &v
	return s
}

func (s *CreateHttpApiRequestIngressConfig) SetWatchNamespace(v string) *CreateHttpApiRequestIngressConfig {
	s.WatchNamespace = &v
	return s
}

func (s *CreateHttpApiRequestIngressConfig) Validate() error {
	return dara.Validate(s)
}
