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
	SetFirstByteTimeout(v int32) *CreateHttpApiRequest
	GetFirstByteTimeout() *int32
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
	// Agent protocols
	AgentProtocols []*string `json:"agentProtocols,omitempty" xml:"agentProtocols,omitempty" type:"Repeated"`
	// $.parameters[0].schema.properties.authConfig.enumValueTitles
	AiProtocols []*string `json:"aiProtocols,omitempty" xml:"aiProtocols,omitempty" type:"Repeated"`
	// The request parameters for API creation.
	AuthConfig *AuthConfig `json:"authConfig,omitempty" xml:"authConfig,omitempty"`
	// $.parameters[0].schema.properties.deployConfigs.items.example
	//
	// example:
	//
	// /v1
	BasePath *string `json:"basePath,omitempty" xml:"basePath,omitempty"`
	// $.parameters[0].schema.example
	DeployConfigs []*HttpApiDeployConfig `json:"deployConfigs,omitempty" xml:"deployConfigs,omitempty" type:"Repeated"`
	// $.parameters[0].schema.properties.aiProtocols.items.description
	//
	// example:
	//
	// $.parameters[0].schema.properties.aiProtocols.items.example
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Create an API of HTTP type
	//
	// example:
	//
	// true
	EnableAuth *bool `json:"enableAuth,omitempty" xml:"enableAuth,omitempty"`
	// First byte timeout
	//
	// example:
	//
	// 30
	FirstByteTimeout *int32 `json:"firstByteTimeout,omitempty" xml:"firstByteTimeout,omitempty"`
	// $.parameters[0].schema.properties.deployConfigs.example
	IngressConfig *CreateHttpApiRequestIngressConfig `json:"ingressConfig,omitempty" xml:"ingressConfig,omitempty" type:"Struct"`
	// Model category
	//
	// example:
	//
	// llm/text-to-image
	ModelCategory *string `json:"modelCategory,omitempty" xml:"modelCategory,omitempty"`
	// $.parameters[0].schema.example
	//
	// This parameter is required.
	//
	// example:
	//
	// test-api
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// $.parameters[0].schema.properties.aiProtocols.description
	Protocols []*string `json:"protocols,omitempty" xml:"protocols,omitempty" type:"Repeated"`
	// Whether to remove base path when forwarding
	//
	// example:
	//
	// true
	RemoveBasePathOnForward *bool `json:"removeBasePathOnForward,omitempty" xml:"removeBasePathOnForward,omitempty"`
	// $.parameters[0].schema.properties.authConfig.example
	//
	// example:
	//
	// rg-xxx
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// $.parameters[0].schema.properties.deployConfigs.description
	//
	// example:
	//
	// Http
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// $.parameters[0].schema.properties.deployConfigs.items.enumValueTitles
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

func (s *CreateHttpApiRequest) GetFirstByteTimeout() *int32 {
	return s.FirstByteTimeout
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

func (s *CreateHttpApiRequest) SetFirstByteTimeout(v int32) *CreateHttpApiRequest {
	s.FirstByteTimeout = &v
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

type CreateHttpApiRequestIngressConfig struct {
	// Cluster ID.
	//
	// example:
	//
	// k7v5eobfzttudni2pw***
	ClusterId *string `json:"clusterId,omitempty" xml:"clusterId,omitempty"`
	// $.parameters[0].schema.properties.deployConfigs.enumValueTitles
	//
	// example:
	//
	// env-cq146allhtgk***
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	// $.parameters[0].schema.properties.enableAuth.example
	//
	// example:
	//
	// mse
	IngressClass *string `json:"ingressClass,omitempty" xml:"ingressClass,omitempty"`
	// $.parameters[0].schema.properties.authConfig.description
	//
	// example:
	//
	// false
	OverrideIngressIp *bool `json:"overrideIngressIp,omitempty" xml:"overrideIngressIp,omitempty"`
	// Deprecated
	//
	// $.parameters[0].schema.properties.enableAuth.description
	//
	// example:
	//
	// src-crdddallhtgtr***
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// $.parameters[0].schema.properties.enableAuth.enumValueTitles
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
