// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHttpApiRouteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackendConfig(v *CreateHttpApiRouteRequestBackendConfig) *CreateHttpApiRouteRequest
	GetBackendConfig() *CreateHttpApiRouteRequestBackendConfig
	SetDeployConfigs(v []*HttpApiDeployConfig) *CreateHttpApiRouteRequest
	GetDeployConfigs() []*HttpApiDeployConfig
	SetDescription(v string) *CreateHttpApiRouteRequest
	GetDescription() *string
	SetDomainIds(v []*string) *CreateHttpApiRouteRequest
	GetDomainIds() []*string
	SetEnvironmentId(v string) *CreateHttpApiRouteRequest
	GetEnvironmentId() *string
	SetMatch(v *HttpRouteMatch) *CreateHttpApiRouteRequest
	GetMatch() *HttpRouteMatch
	SetMcpRouteConfig(v *CreateHttpApiRouteRequestMcpRouteConfig) *CreateHttpApiRouteRequest
	GetMcpRouteConfig() *CreateHttpApiRouteRequestMcpRouteConfig
	SetName(v string) *CreateHttpApiRouteRequest
	GetName() *string
	SetPolicyConfigs(v []*HttpApiPolicyConfigs) *CreateHttpApiRouteRequest
	GetPolicyConfigs() []*HttpApiPolicyConfigs
}

type CreateHttpApiRouteRequest struct {
	// The backend service configuration of the route.
	BackendConfig *CreateHttpApiRouteRequestBackendConfig `json:"backendConfig,omitempty" xml:"backendConfig,omitempty" type:"Struct"`
	// Deprecated
	//
	// The API deployment configuration.
	DeployConfigs []*HttpApiDeployConfig `json:"deployConfigs,omitempty" xml:"deployConfigs,omitempty" type:"Repeated"`
	// The route description.
	//
	// example:
	//
	// User login route
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The domain name IDs.
	DomainIds []*string `json:"domainIds,omitempty" xml:"domainIds,omitempty" type:"Repeated"`
	// The environment ID.
	//
	// example:
	//
	// env-cpqnr6tlhtgubcv***
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	// The route match rule.
	Match *HttpRouteMatch `json:"match,omitempty" xml:"match,omitempty"`
	// The MCP route configuration.
	McpRouteConfig *CreateHttpApiRouteRequestMcpRouteConfig `json:"mcpRouteConfig,omitempty" xml:"mcpRouteConfig,omitempty" type:"Struct"`
	// The route name.
	//
	// example:
	//
	// login
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The policy type.
	PolicyConfigs []*HttpApiPolicyConfigs `json:"policyConfigs,omitempty" xml:"policyConfigs,omitempty" type:"Repeated"`
}

func (s CreateHttpApiRouteRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateHttpApiRouteRequest) GoString() string {
	return s.String()
}

func (s *CreateHttpApiRouteRequest) GetBackendConfig() *CreateHttpApiRouteRequestBackendConfig {
	return s.BackendConfig
}

func (s *CreateHttpApiRouteRequest) GetDeployConfigs() []*HttpApiDeployConfig {
	return s.DeployConfigs
}

func (s *CreateHttpApiRouteRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateHttpApiRouteRequest) GetDomainIds() []*string {
	return s.DomainIds
}

func (s *CreateHttpApiRouteRequest) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *CreateHttpApiRouteRequest) GetMatch() *HttpRouteMatch {
	return s.Match
}

func (s *CreateHttpApiRouteRequest) GetMcpRouteConfig() *CreateHttpApiRouteRequestMcpRouteConfig {
	return s.McpRouteConfig
}

func (s *CreateHttpApiRouteRequest) GetName() *string {
	return s.Name
}

func (s *CreateHttpApiRouteRequest) GetPolicyConfigs() []*HttpApiPolicyConfigs {
	return s.PolicyConfigs
}

func (s *CreateHttpApiRouteRequest) SetBackendConfig(v *CreateHttpApiRouteRequestBackendConfig) *CreateHttpApiRouteRequest {
	s.BackendConfig = v
	return s
}

func (s *CreateHttpApiRouteRequest) SetDeployConfigs(v []*HttpApiDeployConfig) *CreateHttpApiRouteRequest {
	s.DeployConfigs = v
	return s
}

func (s *CreateHttpApiRouteRequest) SetDescription(v string) *CreateHttpApiRouteRequest {
	s.Description = &v
	return s
}

func (s *CreateHttpApiRouteRequest) SetDomainIds(v []*string) *CreateHttpApiRouteRequest {
	s.DomainIds = v
	return s
}

func (s *CreateHttpApiRouteRequest) SetEnvironmentId(v string) *CreateHttpApiRouteRequest {
	s.EnvironmentId = &v
	return s
}

func (s *CreateHttpApiRouteRequest) SetMatch(v *HttpRouteMatch) *CreateHttpApiRouteRequest {
	s.Match = v
	return s
}

func (s *CreateHttpApiRouteRequest) SetMcpRouteConfig(v *CreateHttpApiRouteRequestMcpRouteConfig) *CreateHttpApiRouteRequest {
	s.McpRouteConfig = v
	return s
}

func (s *CreateHttpApiRouteRequest) SetName(v string) *CreateHttpApiRouteRequest {
	s.Name = &v
	return s
}

func (s *CreateHttpApiRouteRequest) SetPolicyConfigs(v []*HttpApiPolicyConfigs) *CreateHttpApiRouteRequest {
	s.PolicyConfigs = v
	return s
}

func (s *CreateHttpApiRouteRequest) Validate() error {
	if s.BackendConfig != nil {
		if err := s.BackendConfig.Validate(); err != nil {
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
	if s.Match != nil {
		if err := s.Match.Validate(); err != nil {
			return err
		}
	}
	if s.McpRouteConfig != nil {
		if err := s.McpRouteConfig.Validate(); err != nil {
			return err
		}
	}
	if s.PolicyConfigs != nil {
		for _, item := range s.PolicyConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateHttpApiRouteRequestBackendConfig struct {
	// The backend service scenario. Valid values:
	//
	// - SingleService: single service.
	//
	// - MultiServiceByRatio: multiple services with ratio-based canary release.
	//
	// - Mock: mock service.
	//
	// - Redirect: redirect service.
	//
	// example:
	//
	// SingleService
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty"`
	// The list of backend services.
	Services []*CreateHttpApiRouteRequestBackendConfigServices `json:"services,omitempty" xml:"services,omitempty" type:"Repeated"`
}

func (s CreateHttpApiRouteRequestBackendConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateHttpApiRouteRequestBackendConfig) GoString() string {
	return s.String()
}

func (s *CreateHttpApiRouteRequestBackendConfig) GetScene() *string {
	return s.Scene
}

func (s *CreateHttpApiRouteRequestBackendConfig) GetServices() []*CreateHttpApiRouteRequestBackendConfigServices {
	return s.Services
}

func (s *CreateHttpApiRouteRequestBackendConfig) SetScene(v string) *CreateHttpApiRouteRequestBackendConfig {
	s.Scene = &v
	return s
}

func (s *CreateHttpApiRouteRequestBackendConfig) SetServices(v []*CreateHttpApiRouteRequestBackendConfigServices) *CreateHttpApiRouteRequestBackendConfig {
	s.Services = v
	return s
}

func (s *CreateHttpApiRouteRequestBackendConfig) Validate() error {
	if s.Services != nil {
		for _, item := range s.Services {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateHttpApiRouteRequestBackendConfigServices struct {
	// The service group. Used in the HTTP-to-Dubbo conversion scenario.
	//
	// example:
	//
	// DEFAULT_GROUP
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// The HTTP-to-Dubbo protocol conversion configuration. Only supported for SingleService MSE_NACOS DUBBO backends of HTTP APIs.
	//
	// example:
	//
	// {"dubboServiceName":"com.alibaba.nacos.example.dubbo.service.DemoService","dubboServiceVersion":"1.0.0","dubboServiceGroup":"DEV","methodMapList":[{"dubboMethodName":"sayName","httpMethod":"ALL_GET","methodPath":"/dubbo/sayName","passThroughAllHeaders":"PASS_ALL"}]}
	HttpDubboTranscoder *HttpDubboTranscoder `json:"httpDubboTranscoder,omitempty" xml:"httpDubboTranscoder,omitempty"`
	// The target model name. This field is shared by multiple model backend scenarios. The specific routing or model rewrite semantics are determined by backendConfig.scene. This field is required for the SemanticRouter scenario. If not specified in the AiAutoRouter scenario, the default model of the AI service is used.
	//
	// example:
	//
	// qwen-plus
	ModelName *string `json:"modelName,omitempty" xml:"modelName,omitempty"`
	// The service namespace. Used in the HTTP-to-Dubbo conversion scenario.
	//
	// example:
	//
	// public
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// The service port. Do not specify this parameter for dynamic ports.
	//
	// example:
	//
	// 8080
	Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
	// The service protocol. Valid values:
	//
	// - HTTP.
	//
	// - HTTPS.
	//
	// example:
	//
	// HTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// The service ID.
	//
	// example:
	//
	// svc-crbgq0dlhtgr***
	ServiceId *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	// The service source type. Used in the HTTP-to-Dubbo conversion scenario.
	//
	// example:
	//
	// MSE_NACOS
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// The service version. This parameter is valid only in the tag-based scenario.
	//
	// example:
	//
	// v1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// The traffic ratio percentage value.
	//
	// example:
	//
	// 49
	Weight *int32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s CreateHttpApiRouteRequestBackendConfigServices) String() string {
	return dara.Prettify(s)
}

func (s CreateHttpApiRouteRequestBackendConfigServices) GoString() string {
	return s.String()
}

func (s *CreateHttpApiRouteRequestBackendConfigServices) GetGroupName() *string {
	return s.GroupName
}

func (s *CreateHttpApiRouteRequestBackendConfigServices) GetHttpDubboTranscoder() *HttpDubboTranscoder {
	return s.HttpDubboTranscoder
}

func (s *CreateHttpApiRouteRequestBackendConfigServices) GetModelName() *string {
	return s.ModelName
}

func (s *CreateHttpApiRouteRequestBackendConfigServices) GetNamespace() *string {
	return s.Namespace
}

func (s *CreateHttpApiRouteRequestBackendConfigServices) GetPort() *int32 {
	return s.Port
}

func (s *CreateHttpApiRouteRequestBackendConfigServices) GetProtocol() *string {
	return s.Protocol
}

func (s *CreateHttpApiRouteRequestBackendConfigServices) GetServiceId() *string {
	return s.ServiceId
}

func (s *CreateHttpApiRouteRequestBackendConfigServices) GetSourceType() *string {
	return s.SourceType
}

func (s *CreateHttpApiRouteRequestBackendConfigServices) GetVersion() *string {
	return s.Version
}

func (s *CreateHttpApiRouteRequestBackendConfigServices) GetWeight() *int32 {
	return s.Weight
}

func (s *CreateHttpApiRouteRequestBackendConfigServices) SetGroupName(v string) *CreateHttpApiRouteRequestBackendConfigServices {
	s.GroupName = &v
	return s
}

func (s *CreateHttpApiRouteRequestBackendConfigServices) SetHttpDubboTranscoder(v *HttpDubboTranscoder) *CreateHttpApiRouteRequestBackendConfigServices {
	s.HttpDubboTranscoder = v
	return s
}

func (s *CreateHttpApiRouteRequestBackendConfigServices) SetModelName(v string) *CreateHttpApiRouteRequestBackendConfigServices {
	s.ModelName = &v
	return s
}

func (s *CreateHttpApiRouteRequestBackendConfigServices) SetNamespace(v string) *CreateHttpApiRouteRequestBackendConfigServices {
	s.Namespace = &v
	return s
}

func (s *CreateHttpApiRouteRequestBackendConfigServices) SetPort(v int32) *CreateHttpApiRouteRequestBackendConfigServices {
	s.Port = &v
	return s
}

func (s *CreateHttpApiRouteRequestBackendConfigServices) SetProtocol(v string) *CreateHttpApiRouteRequestBackendConfigServices {
	s.Protocol = &v
	return s
}

func (s *CreateHttpApiRouteRequestBackendConfigServices) SetServiceId(v string) *CreateHttpApiRouteRequestBackendConfigServices {
	s.ServiceId = &v
	return s
}

func (s *CreateHttpApiRouteRequestBackendConfigServices) SetSourceType(v string) *CreateHttpApiRouteRequestBackendConfigServices {
	s.SourceType = &v
	return s
}

func (s *CreateHttpApiRouteRequestBackendConfigServices) SetVersion(v string) *CreateHttpApiRouteRequestBackendConfigServices {
	s.Version = &v
	return s
}

func (s *CreateHttpApiRouteRequestBackendConfigServices) SetWeight(v int32) *CreateHttpApiRouteRequestBackendConfigServices {
	s.Weight = &v
	return s
}

func (s *CreateHttpApiRouteRequestBackendConfigServices) Validate() error {
	if s.HttpDubboTranscoder != nil {
		if err := s.HttpDubboTranscoder.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateHttpApiRouteRequestMcpRouteConfig struct {
	// The exposed URI path.
	//
	// example:
	//
	// /v1/chat/completions
	ExposedUriPath *string `json:"exposedUriPath,omitempty" xml:"exposedUriPath,omitempty"`
	// Specifies whether to enable MCP observability. Default value: false.
	//
	// example:
	//
	// false
	McpStatisticsEnable *bool `json:"mcpStatisticsEnable,omitempty" xml:"mcpStatisticsEnable,omitempty"`
	// The service protocol. Valid values:
	//
	// - TCP.
	//
	// - HTTP.
	//
	// - DUBBO.
	//
	// example:
	//
	// HTTP,HTTPS
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
}

func (s CreateHttpApiRouteRequestMcpRouteConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateHttpApiRouteRequestMcpRouteConfig) GoString() string {
	return s.String()
}

func (s *CreateHttpApiRouteRequestMcpRouteConfig) GetExposedUriPath() *string {
	return s.ExposedUriPath
}

func (s *CreateHttpApiRouteRequestMcpRouteConfig) GetMcpStatisticsEnable() *bool {
	return s.McpStatisticsEnable
}

func (s *CreateHttpApiRouteRequestMcpRouteConfig) GetProtocol() *string {
	return s.Protocol
}

func (s *CreateHttpApiRouteRequestMcpRouteConfig) SetExposedUriPath(v string) *CreateHttpApiRouteRequestMcpRouteConfig {
	s.ExposedUriPath = &v
	return s
}

func (s *CreateHttpApiRouteRequestMcpRouteConfig) SetMcpStatisticsEnable(v bool) *CreateHttpApiRouteRequestMcpRouteConfig {
	s.McpStatisticsEnable = &v
	return s
}

func (s *CreateHttpApiRouteRequestMcpRouteConfig) SetProtocol(v string) *CreateHttpApiRouteRequestMcpRouteConfig {
	s.Protocol = &v
	return s
}

func (s *CreateHttpApiRouteRequestMcpRouteConfig) Validate() error {
	return dara.Validate(s)
}
