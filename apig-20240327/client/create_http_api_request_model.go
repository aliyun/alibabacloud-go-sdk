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
	SetBelongGatewayId(v string) *CreateHttpApiRequest
	GetBelongGatewayId() *string
	SetDeployConfigs(v []*HttpApiDeployConfig) *CreateHttpApiRequest
	GetDeployConfigs() []*HttpApiDeployConfig
	SetDescription(v string) *CreateHttpApiRequest
	GetDescription() *string
	SetDryRun(v bool) *CreateHttpApiRequest
	GetDryRun() *bool
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
	SetStrategy(v string) *CreateHttpApiRequest
	GetStrategy() *string
	SetType(v string) *CreateHttpApiRequest
	GetType() *string
	SetVersionConfig(v *HttpApiVersionConfig) *CreateHttpApiRequest
	GetVersionConfig() *HttpApiVersionConfig
	SetClientToken(v string) *CreateHttpApiRequest
	GetClientToken() *string
}

type CreateHttpApiRequest struct {
	// The list of protocols supported by the agent. This parameter is required when type is set to Agent. You do not need to specify this parameter for other types.
	AgentProtocols []*string `json:"agentProtocols,omitempty" xml:"agentProtocols,omitempty" type:"Repeated"`
	// The list of AI API protocols. This parameter is required when type is set to LLM, and only one protocol can be specified. This parameter is required when type is set to Ai, and multiple protocols can be specified. You do not need to specify this parameter for other types. Example protocol entry: OpenAI/v1.
	AiProtocols []*string `json:"aiProtocols,omitempty" xml:"aiProtocols,omitempty" type:"Repeated"`
	// The authentication configuration. This parameter is required when enableAuth is set to true.
	AuthConfig *AuthConfig `json:"authConfig,omitempty" xml:"authConfig,omitempty"`
	// The base path of the API. The value must start with a forward slash (/), cannot exceed 256 bytes in length, and cannot contain spaces. This parameter is required when type is set to Rest. When type is set to LLM, Ai, or Agent, this parameter is optional and defaults to /.
	//
	// example:
	//
	// /v1
	BasePath *string `json:"basePath,omitempty" xml:"basePath,omitempty"`
	// The ID of the gateway to which the API belongs.
	//
	// example:
	//
	// gw-abc123xyz789
	BelongGatewayId *string `json:"belongGatewayId,omitempty" xml:"belongGatewayId,omitempty"`
	// The list of deployment configurations for the HTTP API. This parameter is required when type is set to LLM or Ai, and only one deployment configuration can be specified. This parameter is not validated at the request level for other types.
	DeployConfigs []*HttpApiDeployConfig `json:"deployConfigs,omitempty" xml:"deployConfigs,omitempty" type:"Repeated"`
	// The description of the API.
	//
	// example:
	//
	// Test API for integration
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Deprecated
	//
	// Specifies whether to perform a dry run without executing the operation.
	//
	// example:
	//
	// true
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
	// Specifies whether to enable authentication. This parameter is validated when type is set to LLM, Ai, or Agent. This parameter is not validated at the request level when type is set to Rest.
	//
	// example:
	//
	// true
	EnableAuth *bool `json:"enableAuth,omitempty" xml:"enableAuth,omitempty"`
	// The timeout period for waiting for the first byte from the backend.
	//
	// example:
	//
	// 30
	FirstByteTimeout *int32 `json:"firstByteTimeout,omitempty" xml:"firstByteTimeout,omitempty"`
	// The HTTP Ingress API configuration. This parameter is required and cannot be nil when type is set to HttpIngress. You do not need to specify this parameter for other types.
	IngressConfig *CreateHttpApiRequestIngressConfig `json:"ingressConfig,omitempty" xml:"ingressConfig,omitempty" type:"Struct"`
	// The AI model category. This parameter is optional when type is set to LLM or Ai. You do not need to specify this parameter for other types. Valid values:
	//
	// - Text: text generation.
	//
	// - Image: image generation.
	//
	// - Audio: audio processing.
	//
	// - Video: AI video generation.
	//
	// - MultiModal: multimodal.
	//
	// - Embedding: embedding.
	//
	// - Rerank: reranking.
	//
	// - Others: other.
	//
	// example:
	//
	// Text
	ModelCategory *string `json:"modelCategory,omitempty" xml:"modelCategory,omitempty"`
	// The name of the HTTP API, which identifies the API resource. Example: test-api.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-api
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The list of API access protocols.
	Protocols []*string `json:"protocols,omitempty" xml:"protocols,omitempty" type:"Repeated"`
	// Specifies whether to remove the base path when forwarding requests.
	//
	// example:
	//
	// true
	RemoveBasePathOnForward *bool `json:"removeBasePathOnForward,omitempty" xml:"removeBasePathOnForward,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-xxx
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The conflict merge strategy for import.
	//
	// example:
	//
	// ExistFirst
	Strategy *string `json:"strategy,omitempty" xml:"strategy,omitempty"`
	// The HTTP API type. Valid values:
	//
	// - Http: a standard HTTP API.
	//
	// - Rest: a RESTful API.
	//
	// - WebSocket: a WebSocket API.
	//
	// - HttpIngress: an HTTP API accessed through Ingress.
	//
	// - LLM: a large language model API.
	//
	// - Agent: an agent proxy API.
	//
	// This parameter is required.
	//
	// example:
	//
	// Http
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The API versioning configuration.
	VersionConfig *HttpApiVersionConfig `json:"versionConfig,omitempty" xml:"versionConfig,omitempty"`
	// The idempotency token, which is a globally unique value generated by the caller. We recommend that you use a UUID. The value cannot exceed 64 characters in length. Within approximately 24 hours after the first successful request, a duplicate request that carries the same ClientToken and identical request parameters directly returns the httpApiId created by the first request without creating a duplicate HTTP API. If the same ClientToken is carried but the request parameters are different, the IdempotentParameterMismatch error is returned. If the first request is still being processed, the IdempotentProcessing error is returned. If this parameter is not specified, idempotency control is not enabled, and the behavior is consistent with the existing version.
	//
	// example:
	//
	// 5f7a2c1e-9b3d-4e8f-a1c6-0d2b8e4f7a13
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
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

func (s *CreateHttpApiRequest) GetBelongGatewayId() *string {
	return s.BelongGatewayId
}

func (s *CreateHttpApiRequest) GetDeployConfigs() []*HttpApiDeployConfig {
	return s.DeployConfigs
}

func (s *CreateHttpApiRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateHttpApiRequest) GetDryRun() *bool {
	return s.DryRun
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

func (s *CreateHttpApiRequest) GetStrategy() *string {
	return s.Strategy
}

func (s *CreateHttpApiRequest) GetType() *string {
	return s.Type
}

func (s *CreateHttpApiRequest) GetVersionConfig() *HttpApiVersionConfig {
	return s.VersionConfig
}

func (s *CreateHttpApiRequest) GetClientToken() *string {
	return s.ClientToken
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

func (s *CreateHttpApiRequest) SetBelongGatewayId(v string) *CreateHttpApiRequest {
	s.BelongGatewayId = &v
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

func (s *CreateHttpApiRequest) SetDryRun(v bool) *CreateHttpApiRequest {
	s.DryRun = &v
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

func (s *CreateHttpApiRequest) SetStrategy(v string) *CreateHttpApiRequest {
	s.Strategy = &v
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

func (s *CreateHttpApiRequest) SetClientToken(v string) *CreateHttpApiRequest {
	s.ClientToken = &v
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
	// The cluster ID.
	//
	// example:
	//
	// k7v5eobfzttudni2pw***
	ClusterId *string `json:"clusterId,omitempty" xml:"clusterId,omitempty"`
	// The environment ID.
	//
	// example:
	//
	// env-cq146allhtgk***
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	// The Ingress class to listen on.
	//
	// example:
	//
	// mse
	IngressClass *string `json:"ingressClass,omitempty" xml:"ingressClass,omitempty"`
	// Specifies whether to update the address in the Ingress status.
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
	// The namespace to listen on.
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
