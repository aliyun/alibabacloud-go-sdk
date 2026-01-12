// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHttpApiRouteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackendConfig(v *UpdateHttpApiRouteRequestBackendConfig) *UpdateHttpApiRouteRequest
	GetBackendConfig() *UpdateHttpApiRouteRequestBackendConfig
	SetDeployConfigs(v []*HttpApiDeployConfig) *UpdateHttpApiRouteRequest
	GetDeployConfigs() []*HttpApiDeployConfig
	SetDescription(v string) *UpdateHttpApiRouteRequest
	GetDescription() *string
	SetDomainIds(v []*string) *UpdateHttpApiRouteRequest
	GetDomainIds() []*string
	SetEnvironmentId(v string) *UpdateHttpApiRouteRequest
	GetEnvironmentId() *string
	SetMatch(v *HttpRouteMatch) *UpdateHttpApiRouteRequest
	GetMatch() *HttpRouteMatch
	SetMcpRouteConfig(v *UpdateHttpApiRouteRequestMcpRouteConfig) *UpdateHttpApiRouteRequest
	GetMcpRouteConfig() *UpdateHttpApiRouteRequestMcpRouteConfig
	SetName(v string) *UpdateHttpApiRouteRequest
	GetName() *string
	SetPolicyConfigs(v []*UpdateHttpApiRouteRequestPolicyConfigs) *UpdateHttpApiRouteRequest
	GetPolicyConfigs() []*UpdateHttpApiRouteRequestPolicyConfigs
}

type UpdateHttpApiRouteRequest struct {
	// The backend service configurations of the route.
	BackendConfig *UpdateHttpApiRouteRequestBackendConfig `json:"backendConfig,omitempty" xml:"backendConfig,omitempty" type:"Struct"`
	DeployConfigs []*HttpApiDeployConfig                  `json:"deployConfigs,omitempty" xml:"deployConfigs,omitempty" type:"Repeated"`
	// The route description.
	//
	// example:
	//
	// test route
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The domain IDs.
	DomainIds []*string `json:"domainIds,omitempty" xml:"domainIds,omitempty" type:"Repeated"`
	// The environment ID.
	//
	// example:
	//
	// env-cquqsollhtgid***
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	// The rules for matching the route.
	Match          *HttpRouteMatch                           `json:"match,omitempty" xml:"match,omitempty"`
	McpRouteConfig *UpdateHttpApiRouteRequestMcpRouteConfig  `json:"mcpRouteConfig,omitempty" xml:"mcpRouteConfig,omitempty" type:"Struct"`
	Name           *string                                   `json:"name,omitempty" xml:"name,omitempty"`
	PolicyConfigs  []*UpdateHttpApiRouteRequestPolicyConfigs `json:"policyConfigs,omitempty" xml:"policyConfigs,omitempty" type:"Repeated"`
}

func (s UpdateHttpApiRouteRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpApiRouteRequest) GoString() string {
	return s.String()
}

func (s *UpdateHttpApiRouteRequest) GetBackendConfig() *UpdateHttpApiRouteRequestBackendConfig {
	return s.BackendConfig
}

func (s *UpdateHttpApiRouteRequest) GetDeployConfigs() []*HttpApiDeployConfig {
	return s.DeployConfigs
}

func (s *UpdateHttpApiRouteRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateHttpApiRouteRequest) GetDomainIds() []*string {
	return s.DomainIds
}

func (s *UpdateHttpApiRouteRequest) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *UpdateHttpApiRouteRequest) GetMatch() *HttpRouteMatch {
	return s.Match
}

func (s *UpdateHttpApiRouteRequest) GetMcpRouteConfig() *UpdateHttpApiRouteRequestMcpRouteConfig {
	return s.McpRouteConfig
}

func (s *UpdateHttpApiRouteRequest) GetName() *string {
	return s.Name
}

func (s *UpdateHttpApiRouteRequest) GetPolicyConfigs() []*UpdateHttpApiRouteRequestPolicyConfigs {
	return s.PolicyConfigs
}

func (s *UpdateHttpApiRouteRequest) SetBackendConfig(v *UpdateHttpApiRouteRequestBackendConfig) *UpdateHttpApiRouteRequest {
	s.BackendConfig = v
	return s
}

func (s *UpdateHttpApiRouteRequest) SetDeployConfigs(v []*HttpApiDeployConfig) *UpdateHttpApiRouteRequest {
	s.DeployConfigs = v
	return s
}

func (s *UpdateHttpApiRouteRequest) SetDescription(v string) *UpdateHttpApiRouteRequest {
	s.Description = &v
	return s
}

func (s *UpdateHttpApiRouteRequest) SetDomainIds(v []*string) *UpdateHttpApiRouteRequest {
	s.DomainIds = v
	return s
}

func (s *UpdateHttpApiRouteRequest) SetEnvironmentId(v string) *UpdateHttpApiRouteRequest {
	s.EnvironmentId = &v
	return s
}

func (s *UpdateHttpApiRouteRequest) SetMatch(v *HttpRouteMatch) *UpdateHttpApiRouteRequest {
	s.Match = v
	return s
}

func (s *UpdateHttpApiRouteRequest) SetMcpRouteConfig(v *UpdateHttpApiRouteRequestMcpRouteConfig) *UpdateHttpApiRouteRequest {
	s.McpRouteConfig = v
	return s
}

func (s *UpdateHttpApiRouteRequest) SetName(v string) *UpdateHttpApiRouteRequest {
	s.Name = &v
	return s
}

func (s *UpdateHttpApiRouteRequest) SetPolicyConfigs(v []*UpdateHttpApiRouteRequestPolicyConfigs) *UpdateHttpApiRouteRequest {
	s.PolicyConfigs = v
	return s
}

func (s *UpdateHttpApiRouteRequest) Validate() error {
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

type UpdateHttpApiRouteRequestBackendConfig struct {
	// The backend service scenario.
	//
	// Valid values:
	//
	// 	- SingleService
	//
	// 	- MultiServiceByRatio
	//
	// 	- Redirect
	//
	// 	- Mock
	//
	// example:
	//
	// SingleService
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty"`
	// The backend services.
	Services []*UpdateHttpApiRouteRequestBackendConfigServices `json:"services,omitempty" xml:"services,omitempty" type:"Repeated"`
}

func (s UpdateHttpApiRouteRequestBackendConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpApiRouteRequestBackendConfig) GoString() string {
	return s.String()
}

func (s *UpdateHttpApiRouteRequestBackendConfig) GetScene() *string {
	return s.Scene
}

func (s *UpdateHttpApiRouteRequestBackendConfig) GetServices() []*UpdateHttpApiRouteRequestBackendConfigServices {
	return s.Services
}

func (s *UpdateHttpApiRouteRequestBackendConfig) SetScene(v string) *UpdateHttpApiRouteRequestBackendConfig {
	s.Scene = &v
	return s
}

func (s *UpdateHttpApiRouteRequestBackendConfig) SetServices(v []*UpdateHttpApiRouteRequestBackendConfigServices) *UpdateHttpApiRouteRequestBackendConfig {
	s.Services = v
	return s
}

func (s *UpdateHttpApiRouteRequestBackendConfig) Validate() error {
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

type UpdateHttpApiRouteRequestBackendConfigServices struct {
	// The service port. If you want to use a dynamic port, do not pass this parameter.
	//
	// example:
	//
	// 8080
	Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
	// The protocol. Valid values:
	//
	// 	- HTTP
	//
	// 	- HTTPS
	//
	// example:
	//
	// HTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// The service ID.
	//
	// example:
	//
	// svc-cr6pk4tlhtgm58e***
	ServiceId *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	// The service version.
	//
	// example:
	//
	// v1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// The percentage value of traffic.
	//
	// example:
	//
	// 49
	Weight *int32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s UpdateHttpApiRouteRequestBackendConfigServices) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpApiRouteRequestBackendConfigServices) GoString() string {
	return s.String()
}

func (s *UpdateHttpApiRouteRequestBackendConfigServices) GetPort() *int32 {
	return s.Port
}

func (s *UpdateHttpApiRouteRequestBackendConfigServices) GetProtocol() *string {
	return s.Protocol
}

func (s *UpdateHttpApiRouteRequestBackendConfigServices) GetServiceId() *string {
	return s.ServiceId
}

func (s *UpdateHttpApiRouteRequestBackendConfigServices) GetVersion() *string {
	return s.Version
}

func (s *UpdateHttpApiRouteRequestBackendConfigServices) GetWeight() *int32 {
	return s.Weight
}

func (s *UpdateHttpApiRouteRequestBackendConfigServices) SetPort(v int32) *UpdateHttpApiRouteRequestBackendConfigServices {
	s.Port = &v
	return s
}

func (s *UpdateHttpApiRouteRequestBackendConfigServices) SetProtocol(v string) *UpdateHttpApiRouteRequestBackendConfigServices {
	s.Protocol = &v
	return s
}

func (s *UpdateHttpApiRouteRequestBackendConfigServices) SetServiceId(v string) *UpdateHttpApiRouteRequestBackendConfigServices {
	s.ServiceId = &v
	return s
}

func (s *UpdateHttpApiRouteRequestBackendConfigServices) SetVersion(v string) *UpdateHttpApiRouteRequestBackendConfigServices {
	s.Version = &v
	return s
}

func (s *UpdateHttpApiRouteRequestBackendConfigServices) SetWeight(v int32) *UpdateHttpApiRouteRequestBackendConfigServices {
	s.Weight = &v
	return s
}

func (s *UpdateHttpApiRouteRequestBackendConfigServices) Validate() error {
	return dara.Validate(s)
}

type UpdateHttpApiRouteRequestMcpRouteConfig struct {
	ExposedUriPath      *string `json:"exposedUriPath,omitempty" xml:"exposedUriPath,omitempty"`
	McpStatisticsEnable *bool   `json:"mcpStatisticsEnable,omitempty" xml:"mcpStatisticsEnable,omitempty"`
	Protocol            *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
}

func (s UpdateHttpApiRouteRequestMcpRouteConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpApiRouteRequestMcpRouteConfig) GoString() string {
	return s.String()
}

func (s *UpdateHttpApiRouteRequestMcpRouteConfig) GetExposedUriPath() *string {
	return s.ExposedUriPath
}

func (s *UpdateHttpApiRouteRequestMcpRouteConfig) GetMcpStatisticsEnable() *bool {
	return s.McpStatisticsEnable
}

func (s *UpdateHttpApiRouteRequestMcpRouteConfig) GetProtocol() *string {
	return s.Protocol
}

func (s *UpdateHttpApiRouteRequestMcpRouteConfig) SetExposedUriPath(v string) *UpdateHttpApiRouteRequestMcpRouteConfig {
	s.ExposedUriPath = &v
	return s
}

func (s *UpdateHttpApiRouteRequestMcpRouteConfig) SetMcpStatisticsEnable(v bool) *UpdateHttpApiRouteRequestMcpRouteConfig {
	s.McpStatisticsEnable = &v
	return s
}

func (s *UpdateHttpApiRouteRequestMcpRouteConfig) SetProtocol(v string) *UpdateHttpApiRouteRequestMcpRouteConfig {
	s.Protocol = &v
	return s
}

func (s *UpdateHttpApiRouteRequestMcpRouteConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateHttpApiRouteRequestPolicyConfigs struct {
	AiCacheConfig          *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfig          `json:"aiCacheConfig,omitempty" xml:"aiCacheConfig,omitempty" type:"Struct"`
	AiFallbackConfig       *UpdateHttpApiRouteRequestPolicyConfigsAiFallbackConfig       `json:"aiFallbackConfig,omitempty" xml:"aiFallbackConfig,omitempty" type:"Struct"`
	AiNetworkSearchConfig  *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfig  `json:"aiNetworkSearchConfig,omitempty" xml:"aiNetworkSearchConfig,omitempty" type:"Struct"`
	AiSecurityGuardConfig  *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfig  `json:"aiSecurityGuardConfig,omitempty" xml:"aiSecurityGuardConfig,omitempty" type:"Struct"`
	AiStatisticsConfig     *UpdateHttpApiRouteRequestPolicyConfigsAiStatisticsConfig     `json:"aiStatisticsConfig,omitempty" xml:"aiStatisticsConfig,omitempty" type:"Struct"`
	AiTokenRateLimitConfig *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfig `json:"aiTokenRateLimitConfig,omitempty" xml:"aiTokenRateLimitConfig,omitempty" type:"Struct"`
	AiToolSelectionConfig  *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfig  `json:"aiToolSelectionConfig,omitempty" xml:"aiToolSelectionConfig,omitempty" type:"Struct"`
	Enable                 *bool                                                         `json:"enable,omitempty" xml:"enable,omitempty"`
	SemanticRouterConfig   *UpdateHttpApiRouteRequestPolicyConfigsSemanticRouterConfig   `json:"semanticRouterConfig,omitempty" xml:"semanticRouterConfig,omitempty" type:"Struct"`
	Type                   *string                                                       `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateHttpApiRouteRequestPolicyConfigs) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpApiRouteRequestPolicyConfigs) GoString() string {
	return s.String()
}

func (s *UpdateHttpApiRouteRequestPolicyConfigs) GetAiCacheConfig() *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfig {
	return s.AiCacheConfig
}

func (s *UpdateHttpApiRouteRequestPolicyConfigs) GetAiFallbackConfig() *UpdateHttpApiRouteRequestPolicyConfigsAiFallbackConfig {
	return s.AiFallbackConfig
}

func (s *UpdateHttpApiRouteRequestPolicyConfigs) GetAiNetworkSearchConfig() *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfig {
	return s.AiNetworkSearchConfig
}

func (s *UpdateHttpApiRouteRequestPolicyConfigs) GetAiSecurityGuardConfig() *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfig {
	return s.AiSecurityGuardConfig
}

func (s *UpdateHttpApiRouteRequestPolicyConfigs) GetAiStatisticsConfig() *UpdateHttpApiRouteRequestPolicyConfigsAiStatisticsConfig {
	return s.AiStatisticsConfig
}

func (s *UpdateHttpApiRouteRequestPolicyConfigs) GetAiTokenRateLimitConfig() *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfig {
	return s.AiTokenRateLimitConfig
}

func (s *UpdateHttpApiRouteRequestPolicyConfigs) GetAiToolSelectionConfig() *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfig {
	return s.AiToolSelectionConfig
}

func (s *UpdateHttpApiRouteRequestPolicyConfigs) GetEnable() *bool {
	return s.Enable
}

func (s *UpdateHttpApiRouteRequestPolicyConfigs) GetSemanticRouterConfig() *UpdateHttpApiRouteRequestPolicyConfigsSemanticRouterConfig {
	return s.SemanticRouterConfig
}

func (s *UpdateHttpApiRouteRequestPolicyConfigs) GetType() *string {
	return s.Type
}

func (s *UpdateHttpApiRouteRequestPolicyConfigs) SetAiCacheConfig(v *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfig) *UpdateHttpApiRouteRequestPolicyConfigs {
	s.AiCacheConfig = v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigs) SetAiFallbackConfig(v *UpdateHttpApiRouteRequestPolicyConfigsAiFallbackConfig) *UpdateHttpApiRouteRequestPolicyConfigs {
	s.AiFallbackConfig = v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigs) SetAiNetworkSearchConfig(v *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfig) *UpdateHttpApiRouteRequestPolicyConfigs {
	s.AiNetworkSearchConfig = v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigs) SetAiSecurityGuardConfig(v *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfig) *UpdateHttpApiRouteRequestPolicyConfigs {
	s.AiSecurityGuardConfig = v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigs) SetAiStatisticsConfig(v *UpdateHttpApiRouteRequestPolicyConfigsAiStatisticsConfig) *UpdateHttpApiRouteRequestPolicyConfigs {
	s.AiStatisticsConfig = v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigs) SetAiTokenRateLimitConfig(v *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfig) *UpdateHttpApiRouteRequestPolicyConfigs {
	s.AiTokenRateLimitConfig = v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigs) SetAiToolSelectionConfig(v *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfig) *UpdateHttpApiRouteRequestPolicyConfigs {
	s.AiToolSelectionConfig = v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigs) SetEnable(v bool) *UpdateHttpApiRouteRequestPolicyConfigs {
	s.Enable = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigs) SetSemanticRouterConfig(v *UpdateHttpApiRouteRequestPolicyConfigsSemanticRouterConfig) *UpdateHttpApiRouteRequestPolicyConfigs {
	s.SemanticRouterConfig = v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigs) SetType(v string) *UpdateHttpApiRouteRequestPolicyConfigs {
	s.Type = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigs) Validate() error {
	if s.AiCacheConfig != nil {
		if err := s.AiCacheConfig.Validate(); err != nil {
			return err
		}
	}
	if s.AiFallbackConfig != nil {
		if err := s.AiFallbackConfig.Validate(); err != nil {
			return err
		}
	}
	if s.AiNetworkSearchConfig != nil {
		if err := s.AiNetworkSearchConfig.Validate(); err != nil {
			return err
		}
	}
	if s.AiSecurityGuardConfig != nil {
		if err := s.AiSecurityGuardConfig.Validate(); err != nil {
			return err
		}
	}
	if s.AiStatisticsConfig != nil {
		if err := s.AiStatisticsConfig.Validate(); err != nil {
			return err
		}
	}
	if s.AiTokenRateLimitConfig != nil {
		if err := s.AiTokenRateLimitConfig.Validate(); err != nil {
			return err
		}
	}
	if s.AiToolSelectionConfig != nil {
		if err := s.AiToolSelectionConfig.Validate(); err != nil {
			return err
		}
	}
	if s.SemanticRouterConfig != nil {
		if err := s.SemanticRouterConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfig struct {
	CacheKeyStrategy *string                                                             `json:"cacheKeyStrategy,omitempty" xml:"cacheKeyStrategy,omitempty"`
	CacheMode        *string                                                             `json:"cacheMode,omitempty" xml:"cacheMode,omitempty"`
	CacheTTL         *int32                                                              `json:"cacheTTL,omitempty" xml:"cacheTTL,omitempty"`
	EmbeddingConfig  *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigEmbeddingConfig `json:"embeddingConfig,omitempty" xml:"embeddingConfig,omitempty" type:"Struct"`
	PluginStatus     *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigPluginStatus    `json:"pluginStatus,omitempty" xml:"pluginStatus,omitempty" type:"Struct"`
	RedisConfig      *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigRedisConfig     `json:"redisConfig,omitempty" xml:"redisConfig,omitempty" type:"Struct"`
	VectorConfig     *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigVectorConfig    `json:"vectorConfig,omitempty" xml:"vectorConfig,omitempty" type:"Struct"`
}

func (s UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfig) GoString() string {
	return s.String()
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfig) GetCacheKeyStrategy() *string {
	return s.CacheKeyStrategy
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfig) GetCacheMode() *string {
	return s.CacheMode
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfig) GetCacheTTL() *int32 {
	return s.CacheTTL
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfig) GetEmbeddingConfig() *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigEmbeddingConfig {
	return s.EmbeddingConfig
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfig) GetPluginStatus() *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigPluginStatus {
	return s.PluginStatus
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfig) GetRedisConfig() *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigRedisConfig {
	return s.RedisConfig
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfig) GetVectorConfig() *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigVectorConfig {
	return s.VectorConfig
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfig) SetCacheKeyStrategy(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfig {
	s.CacheKeyStrategy = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfig) SetCacheMode(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfig {
	s.CacheMode = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfig) SetCacheTTL(v int32) *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfig {
	s.CacheTTL = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfig) SetEmbeddingConfig(v *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigEmbeddingConfig) *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfig {
	s.EmbeddingConfig = v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfig) SetPluginStatus(v *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigPluginStatus) *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfig {
	s.PluginStatus = v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfig) SetRedisConfig(v *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigRedisConfig) *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfig {
	s.RedisConfig = v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfig) SetVectorConfig(v *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigVectorConfig) *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfig {
	s.VectorConfig = v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfig) Validate() error {
	if s.EmbeddingConfig != nil {
		if err := s.EmbeddingConfig.Validate(); err != nil {
			return err
		}
	}
	if s.PluginStatus != nil {
		if err := s.PluginStatus.Validate(); err != nil {
			return err
		}
	}
	if s.RedisConfig != nil {
		if err := s.RedisConfig.Validate(); err != nil {
			return err
		}
	}
	if s.VectorConfig != nil {
		if err := s.VectorConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigEmbeddingConfig struct {
	ModelName *string `json:"modelName,omitempty" xml:"modelName,omitempty"`
	ServiceId *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	Timeout   *int32  `json:"timeout,omitempty" xml:"timeout,omitempty"`
	Type      *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigEmbeddingConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigEmbeddingConfig) GoString() string {
	return s.String()
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigEmbeddingConfig) GetModelName() *string {
	return s.ModelName
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigEmbeddingConfig) GetServiceId() *string {
	return s.ServiceId
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigEmbeddingConfig) GetTimeout() *int32 {
	return s.Timeout
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigEmbeddingConfig) GetType() *string {
	return s.Type
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigEmbeddingConfig) SetModelName(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigEmbeddingConfig {
	s.ModelName = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigEmbeddingConfig) SetServiceId(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigEmbeddingConfig {
	s.ServiceId = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigEmbeddingConfig) SetTimeout(v int32) *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigEmbeddingConfig {
	s.Timeout = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigEmbeddingConfig) SetType(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigEmbeddingConfig {
	s.Type = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigEmbeddingConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigPluginStatus struct {
	ErrorLogs      map[string]*string `json:"errorLogs,omitempty" xml:"errorLogs,omitempty"`
	PluginId       *string            `json:"pluginId,omitempty" xml:"pluginId,omitempty"`
	ServiceHealthy *bool              `json:"serviceHealthy,omitempty" xml:"serviceHealthy,omitempty"`
}

func (s UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigPluginStatus) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigPluginStatus) GoString() string {
	return s.String()
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigPluginStatus) GetErrorLogs() map[string]*string {
	return s.ErrorLogs
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigPluginStatus) GetPluginId() *string {
	return s.PluginId
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigPluginStatus) GetServiceHealthy() *bool {
	return s.ServiceHealthy
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigPluginStatus) SetErrorLogs(v map[string]*string) *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigPluginStatus {
	s.ErrorLogs = v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigPluginStatus) SetPluginId(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigPluginStatus {
	s.PluginId = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigPluginStatus) SetServiceHealthy(v bool) *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigPluginStatus {
	s.ServiceHealthy = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigPluginStatus) Validate() error {
	return dara.Validate(s)
}

type UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigRedisConfig struct {
	DatabaseNumber *int32  `json:"databaseNumber,omitempty" xml:"databaseNumber,omitempty"`
	Host           *string `json:"host,omitempty" xml:"host,omitempty"`
	Password       *string `json:"password,omitempty" xml:"password,omitempty"`
	Port           *int32  `json:"port,omitempty" xml:"port,omitempty"`
	Timeout        *int32  `json:"timeout,omitempty" xml:"timeout,omitempty"`
	Username       *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigRedisConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigRedisConfig) GoString() string {
	return s.String()
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigRedisConfig) GetDatabaseNumber() *int32 {
	return s.DatabaseNumber
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigRedisConfig) GetHost() *string {
	return s.Host
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigRedisConfig) GetPassword() *string {
	return s.Password
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigRedisConfig) GetPort() *int32 {
	return s.Port
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigRedisConfig) GetTimeout() *int32 {
	return s.Timeout
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigRedisConfig) GetUsername() *string {
	return s.Username
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigRedisConfig) SetDatabaseNumber(v int32) *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigRedisConfig {
	s.DatabaseNumber = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigRedisConfig) SetHost(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigRedisConfig {
	s.Host = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigRedisConfig) SetPassword(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigRedisConfig {
	s.Password = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigRedisConfig) SetPort(v int32) *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigRedisConfig {
	s.Port = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigRedisConfig) SetTimeout(v int32) *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigRedisConfig {
	s.Timeout = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigRedisConfig) SetUsername(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigRedisConfig {
	s.Username = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigRedisConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigVectorConfig struct {
	ApiKey       *string  `json:"apiKey,omitempty" xml:"apiKey,omitempty"`
	CollectionId *string  `json:"collectionId,omitempty" xml:"collectionId,omitempty"`
	ServiceHost  *string  `json:"serviceHost,omitempty" xml:"serviceHost,omitempty"`
	Threshold    *float32 `json:"threshold,omitempty" xml:"threshold,omitempty"`
	// example:
	//
	// 6000
	Timeout *int32  `json:"timeout,omitempty" xml:"timeout,omitempty"`
	Type    *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigVectorConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigVectorConfig) GoString() string {
	return s.String()
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigVectorConfig) GetApiKey() *string {
	return s.ApiKey
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigVectorConfig) GetCollectionId() *string {
	return s.CollectionId
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigVectorConfig) GetServiceHost() *string {
	return s.ServiceHost
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigVectorConfig) GetThreshold() *float32 {
	return s.Threshold
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigVectorConfig) GetTimeout() *int32 {
	return s.Timeout
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigVectorConfig) GetType() *string {
	return s.Type
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigVectorConfig) SetApiKey(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigVectorConfig {
	s.ApiKey = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigVectorConfig) SetCollectionId(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigVectorConfig {
	s.CollectionId = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigVectorConfig) SetServiceHost(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigVectorConfig {
	s.ServiceHost = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigVectorConfig) SetThreshold(v float32) *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigVectorConfig {
	s.Threshold = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigVectorConfig) SetTimeout(v int32) *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigVectorConfig {
	s.Timeout = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigVectorConfig) SetType(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigVectorConfig {
	s.Type = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiCacheConfigVectorConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateHttpApiRouteRequestPolicyConfigsAiFallbackConfig struct {
	OnlyRedirectUpstreamCode *bool                                                                   `json:"onlyRedirectUpstreamCode,omitempty" xml:"onlyRedirectUpstreamCode,omitempty"`
	RouteEmbedded            *bool                                                                   `json:"routeEmbedded,omitempty" xml:"routeEmbedded,omitempty"`
	ServiceConfigs           []*UpdateHttpApiRouteRequestPolicyConfigsAiFallbackConfigServiceConfigs `json:"serviceConfigs,omitempty" xml:"serviceConfigs,omitempty" type:"Repeated"`
}

func (s UpdateHttpApiRouteRequestPolicyConfigsAiFallbackConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpApiRouteRequestPolicyConfigsAiFallbackConfig) GoString() string {
	return s.String()
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiFallbackConfig) GetOnlyRedirectUpstreamCode() *bool {
	return s.OnlyRedirectUpstreamCode
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiFallbackConfig) GetRouteEmbedded() *bool {
	return s.RouteEmbedded
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiFallbackConfig) GetServiceConfigs() []*UpdateHttpApiRouteRequestPolicyConfigsAiFallbackConfigServiceConfigs {
	return s.ServiceConfigs
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiFallbackConfig) SetOnlyRedirectUpstreamCode(v bool) *UpdateHttpApiRouteRequestPolicyConfigsAiFallbackConfig {
	s.OnlyRedirectUpstreamCode = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiFallbackConfig) SetRouteEmbedded(v bool) *UpdateHttpApiRouteRequestPolicyConfigsAiFallbackConfig {
	s.RouteEmbedded = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiFallbackConfig) SetServiceConfigs(v []*UpdateHttpApiRouteRequestPolicyConfigsAiFallbackConfigServiceConfigs) *UpdateHttpApiRouteRequestPolicyConfigsAiFallbackConfig {
	s.ServiceConfigs = v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiFallbackConfig) Validate() error {
	if s.ServiceConfigs != nil {
		for _, item := range s.ServiceConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateHttpApiRouteRequestPolicyConfigsAiFallbackConfigServiceConfigs struct {
	Name                 *string `json:"name,omitempty" xml:"name,omitempty"`
	PassThroughModelName *bool   `json:"passThroughModelName,omitempty" xml:"passThroughModelName,omitempty"`
	ServiceId            *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	TargetModelName      *string `json:"targetModelName,omitempty" xml:"targetModelName,omitempty"`
}

func (s UpdateHttpApiRouteRequestPolicyConfigsAiFallbackConfigServiceConfigs) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpApiRouteRequestPolicyConfigsAiFallbackConfigServiceConfigs) GoString() string {
	return s.String()
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiFallbackConfigServiceConfigs) GetName() *string {
	return s.Name
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiFallbackConfigServiceConfigs) GetPassThroughModelName() *bool {
	return s.PassThroughModelName
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiFallbackConfigServiceConfigs) GetServiceId() *string {
	return s.ServiceId
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiFallbackConfigServiceConfigs) GetTargetModelName() *string {
	return s.TargetModelName
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiFallbackConfigServiceConfigs) SetName(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiFallbackConfigServiceConfigs {
	s.Name = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiFallbackConfigServiceConfigs) SetPassThroughModelName(v bool) *UpdateHttpApiRouteRequestPolicyConfigsAiFallbackConfigServiceConfigs {
	s.PassThroughModelName = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiFallbackConfigServiceConfigs) SetServiceId(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiFallbackConfigServiceConfigs {
	s.ServiceId = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiFallbackConfigServiceConfigs) SetTargetModelName(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiFallbackConfigServiceConfigs {
	s.TargetModelName = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiFallbackConfigServiceConfigs) Validate() error {
	return dara.Validate(s)
}

type UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfig struct {
	DefaultEnable      *bool                                                                          `json:"defaultEnable,omitempty" xml:"defaultEnable,omitempty"`
	DefaultLang        *string                                                                        `json:"defaultLang,omitempty" xml:"defaultLang,omitempty"`
	NeedReference      *bool                                                                          `json:"needReference,omitempty" xml:"needReference,omitempty"`
	PluginStatus       *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigPluginStatus       `json:"pluginStatus,omitempty" xml:"pluginStatus,omitempty" type:"Struct"`
	ReferenceFormat    *string                                                                        `json:"referenceFormat,omitempty" xml:"referenceFormat,omitempty"`
	ReferenceLocation  *string                                                                        `json:"referenceLocation,omitempty" xml:"referenceLocation,omitempty"`
	SearchEngineConfig *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchEngineConfig `json:"searchEngineConfig,omitempty" xml:"searchEngineConfig,omitempty" type:"Struct"`
	SearchFrom         []*UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchFrom       `json:"searchFrom,omitempty" xml:"searchFrom,omitempty" type:"Repeated"`
	SearchRewrite      *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchRewrite      `json:"searchRewrite,omitempty" xml:"searchRewrite,omitempty" type:"Struct"`
}

func (s UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfig) GoString() string {
	return s.String()
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfig) GetDefaultEnable() *bool {
	return s.DefaultEnable
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfig) GetDefaultLang() *string {
	return s.DefaultLang
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfig) GetNeedReference() *bool {
	return s.NeedReference
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfig) GetPluginStatus() *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigPluginStatus {
	return s.PluginStatus
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfig) GetReferenceFormat() *string {
	return s.ReferenceFormat
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfig) GetReferenceLocation() *string {
	return s.ReferenceLocation
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfig) GetSearchEngineConfig() *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchEngineConfig {
	return s.SearchEngineConfig
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfig) GetSearchFrom() []*UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchFrom {
	return s.SearchFrom
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfig) GetSearchRewrite() *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchRewrite {
	return s.SearchRewrite
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfig) SetDefaultEnable(v bool) *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfig {
	s.DefaultEnable = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfig) SetDefaultLang(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfig {
	s.DefaultLang = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfig) SetNeedReference(v bool) *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfig {
	s.NeedReference = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfig) SetPluginStatus(v *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigPluginStatus) *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfig {
	s.PluginStatus = v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfig) SetReferenceFormat(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfig {
	s.ReferenceFormat = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfig) SetReferenceLocation(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfig {
	s.ReferenceLocation = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfig) SetSearchEngineConfig(v *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchEngineConfig) *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfig {
	s.SearchEngineConfig = v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfig) SetSearchFrom(v []*UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchFrom) *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfig {
	s.SearchFrom = v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfig) SetSearchRewrite(v *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchRewrite) *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfig {
	s.SearchRewrite = v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfig) Validate() error {
	if s.PluginStatus != nil {
		if err := s.PluginStatus.Validate(); err != nil {
			return err
		}
	}
	if s.SearchEngineConfig != nil {
		if err := s.SearchEngineConfig.Validate(); err != nil {
			return err
		}
	}
	if s.SearchFrom != nil {
		for _, item := range s.SearchFrom {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SearchRewrite != nil {
		if err := s.SearchRewrite.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigPluginStatus struct {
	ErrorLogs      map[string]*string `json:"errorLogs,omitempty" xml:"errorLogs,omitempty"`
	PluginId       *string            `json:"pluginId,omitempty" xml:"pluginId,omitempty"`
	ServiceHealthy *bool              `json:"serviceHealthy,omitempty" xml:"serviceHealthy,omitempty"`
}

func (s UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigPluginStatus) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigPluginStatus) GoString() string {
	return s.String()
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigPluginStatus) GetErrorLogs() map[string]*string {
	return s.ErrorLogs
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigPluginStatus) GetPluginId() *string {
	return s.PluginId
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigPluginStatus) GetServiceHealthy() *bool {
	return s.ServiceHealthy
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigPluginStatus) SetErrorLogs(v map[string]*string) *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigPluginStatus {
	s.ErrorLogs = v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigPluginStatus) SetPluginId(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigPluginStatus {
	s.PluginId = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigPluginStatus) SetServiceHealthy(v bool) *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigPluginStatus {
	s.ServiceHealthy = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigPluginStatus) Validate() error {
	return dara.Validate(s)
}

type UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchEngineConfig struct {
	ApiKey             *string            `json:"apiKey,omitempty" xml:"apiKey,omitempty"`
	ContentMode        *string            `json:"contentMode,omitempty" xml:"contentMode,omitempty"`
	Count              *int32             `json:"count,omitempty" xml:"count,omitempty"`
	Endpoint           *string            `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	Industry           *string            `json:"industry,omitempty" xml:"industry,omitempty"`
	OptionArgs         map[string]*string `json:"optionArgs,omitempty" xml:"optionArgs,omitempty"`
	Start              *int32             `json:"start,omitempty" xml:"start,omitempty"`
	TimeRange          *string            `json:"timeRange,omitempty" xml:"timeRange,omitempty"`
	TimeoutMillisecond *int32             `json:"timeoutMillisecond,omitempty" xml:"timeoutMillisecond,omitempty"`
	Type               *string            `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchEngineConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchEngineConfig) GoString() string {
	return s.String()
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchEngineConfig) GetApiKey() *string {
	return s.ApiKey
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchEngineConfig) GetContentMode() *string {
	return s.ContentMode
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchEngineConfig) GetCount() *int32 {
	return s.Count
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchEngineConfig) GetEndpoint() *string {
	return s.Endpoint
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchEngineConfig) GetIndustry() *string {
	return s.Industry
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchEngineConfig) GetOptionArgs() map[string]*string {
	return s.OptionArgs
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchEngineConfig) GetStart() *int32 {
	return s.Start
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchEngineConfig) GetTimeRange() *string {
	return s.TimeRange
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchEngineConfig) GetTimeoutMillisecond() *int32 {
	return s.TimeoutMillisecond
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchEngineConfig) GetType() *string {
	return s.Type
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchEngineConfig) SetApiKey(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchEngineConfig {
	s.ApiKey = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchEngineConfig) SetContentMode(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchEngineConfig {
	s.ContentMode = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchEngineConfig) SetCount(v int32) *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchEngineConfig {
	s.Count = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchEngineConfig) SetEndpoint(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchEngineConfig {
	s.Endpoint = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchEngineConfig) SetIndustry(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchEngineConfig {
	s.Industry = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchEngineConfig) SetOptionArgs(v map[string]*string) *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchEngineConfig {
	s.OptionArgs = v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchEngineConfig) SetStart(v int32) *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchEngineConfig {
	s.Start = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchEngineConfig) SetTimeRange(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchEngineConfig {
	s.TimeRange = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchEngineConfig) SetTimeoutMillisecond(v int32) *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchEngineConfig {
	s.TimeoutMillisecond = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchEngineConfig) SetType(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchEngineConfig {
	s.Type = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchEngineConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchFrom struct {
	ApiKey             *string            `json:"apiKey,omitempty" xml:"apiKey,omitempty"`
	ContentMode        *string            `json:"contentMode,omitempty" xml:"contentMode,omitempty"`
	Count              *int32             `json:"count,omitempty" xml:"count,omitempty"`
	Endpoint           *string            `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	Industry           *string            `json:"industry,omitempty" xml:"industry,omitempty"`
	OptionArgs         map[string]*string `json:"optionArgs,omitempty" xml:"optionArgs,omitempty"`
	Start              *int32             `json:"start,omitempty" xml:"start,omitempty"`
	TimeRange          *string            `json:"timeRange,omitempty" xml:"timeRange,omitempty"`
	TimeoutMillisecond *int32             `json:"timeoutMillisecond,omitempty" xml:"timeoutMillisecond,omitempty"`
	Type               *string            `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchFrom) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchFrom) GoString() string {
	return s.String()
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchFrom) GetApiKey() *string {
	return s.ApiKey
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchFrom) GetContentMode() *string {
	return s.ContentMode
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchFrom) GetCount() *int32 {
	return s.Count
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchFrom) GetEndpoint() *string {
	return s.Endpoint
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchFrom) GetIndustry() *string {
	return s.Industry
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchFrom) GetOptionArgs() map[string]*string {
	return s.OptionArgs
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchFrom) GetStart() *int32 {
	return s.Start
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchFrom) GetTimeRange() *string {
	return s.TimeRange
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchFrom) GetTimeoutMillisecond() *int32 {
	return s.TimeoutMillisecond
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchFrom) GetType() *string {
	return s.Type
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchFrom) SetApiKey(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchFrom {
	s.ApiKey = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchFrom) SetContentMode(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchFrom {
	s.ContentMode = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchFrom) SetCount(v int32) *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchFrom {
	s.Count = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchFrom) SetEndpoint(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchFrom {
	s.Endpoint = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchFrom) SetIndustry(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchFrom {
	s.Industry = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchFrom) SetOptionArgs(v map[string]*string) *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchFrom {
	s.OptionArgs = v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchFrom) SetStart(v int32) *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchFrom {
	s.Start = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchFrom) SetTimeRange(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchFrom {
	s.TimeRange = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchFrom) SetTimeoutMillisecond(v int32) *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchFrom {
	s.TimeoutMillisecond = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchFrom) SetType(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchFrom {
	s.Type = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchFrom) Validate() error {
	return dara.Validate(s)
}

type UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchRewrite struct {
	Enable             *bool   `json:"enable,omitempty" xml:"enable,omitempty"`
	MaxCount           *int32  `json:"maxCount,omitempty" xml:"maxCount,omitempty"`
	ModelName          *string `json:"modelName,omitempty" xml:"modelName,omitempty"`
	ServiceId          *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	TimeoutMillisecond *int32  `json:"timeoutMillisecond,omitempty" xml:"timeoutMillisecond,omitempty"`
}

func (s UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchRewrite) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchRewrite) GoString() string {
	return s.String()
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchRewrite) GetEnable() *bool {
	return s.Enable
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchRewrite) GetMaxCount() *int32 {
	return s.MaxCount
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchRewrite) GetModelName() *string {
	return s.ModelName
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchRewrite) GetServiceId() *string {
	return s.ServiceId
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchRewrite) GetTimeoutMillisecond() *int32 {
	return s.TimeoutMillisecond
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchRewrite) SetEnable(v bool) *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchRewrite {
	s.Enable = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchRewrite) SetMaxCount(v int32) *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchRewrite {
	s.MaxCount = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchRewrite) SetModelName(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchRewrite {
	s.ModelName = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchRewrite) SetServiceId(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchRewrite {
	s.ServiceId = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchRewrite) SetTimeoutMillisecond(v int32) *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchRewrite {
	s.TimeoutMillisecond = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiNetworkSearchConfigSearchRewrite) Validate() error {
	return dara.Validate(s)
}

type UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfig struct {
	BufferLimit                  *int32                                                                                     `json:"bufferLimit,omitempty" xml:"bufferLimit,omitempty"`
	CheckRequest                 *bool                                                                                      `json:"checkRequest,omitempty" xml:"checkRequest,omitempty"`
	CheckRequestImage            *bool                                                                                      `json:"checkRequestImage,omitempty" xml:"checkRequestImage,omitempty"`
	CheckResponse                *bool                                                                                      `json:"checkResponse,omitempty" xml:"checkResponse,omitempty"`
	CheckResponseImage           *bool                                                                                      `json:"checkResponseImage,omitempty" xml:"checkResponseImage,omitempty"`
	ConsumerRequestCheckService  []*UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigConsumerRequestCheckService  `json:"consumerRequestCheckService,omitempty" xml:"consumerRequestCheckService,omitempty" type:"Repeated"`
	ConsumerResponseCheckService []*UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigConsumerResponseCheckService `json:"consumerResponseCheckService,omitempty" xml:"consumerResponseCheckService,omitempty" type:"Repeated"`
	ConsumerRiskLevel            []*UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigConsumerRiskLevel            `json:"consumerRiskLevel,omitempty" xml:"consumerRiskLevel,omitempty" type:"Repeated"`
	PluginStatus                 *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigPluginStatus                   `json:"pluginStatus,omitempty" xml:"pluginStatus,omitempty" type:"Struct"`
	RequestCheckService          *string                                                                                    `json:"requestCheckService,omitempty" xml:"requestCheckService,omitempty"`
	RequestImageCheckService     *string                                                                                    `json:"requestImageCheckService,omitempty" xml:"requestImageCheckService,omitempty"`
	ResponseCheckService         *string                                                                                    `json:"responseCheckService,omitempty" xml:"responseCheckService,omitempty"`
	ResponseImageCheckService    *string                                                                                    `json:"responseImageCheckService,omitempty" xml:"responseImageCheckService,omitempty"`
	RiskAlertLevel               *string                                                                                    `json:"riskAlertLevel,omitempty" xml:"riskAlertLevel,omitempty"`
	RiskConfig                   []*UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigRiskConfig                   `json:"riskConfig,omitempty" xml:"riskConfig,omitempty" type:"Repeated"`
	ServiceAddress               *string                                                                                    `json:"serviceAddress,omitempty" xml:"serviceAddress,omitempty"`
}

func (s UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfig) GoString() string {
	return s.String()
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfig) GetBufferLimit() *int32 {
	return s.BufferLimit
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfig) GetCheckRequest() *bool {
	return s.CheckRequest
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfig) GetCheckRequestImage() *bool {
	return s.CheckRequestImage
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfig) GetCheckResponse() *bool {
	return s.CheckResponse
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfig) GetCheckResponseImage() *bool {
	return s.CheckResponseImage
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfig) GetConsumerRequestCheckService() []*UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigConsumerRequestCheckService {
	return s.ConsumerRequestCheckService
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfig) GetConsumerResponseCheckService() []*UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigConsumerResponseCheckService {
	return s.ConsumerResponseCheckService
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfig) GetConsumerRiskLevel() []*UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigConsumerRiskLevel {
	return s.ConsumerRiskLevel
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfig) GetPluginStatus() *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigPluginStatus {
	return s.PluginStatus
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfig) GetRequestCheckService() *string {
	return s.RequestCheckService
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfig) GetRequestImageCheckService() *string {
	return s.RequestImageCheckService
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfig) GetResponseCheckService() *string {
	return s.ResponseCheckService
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfig) GetResponseImageCheckService() *string {
	return s.ResponseImageCheckService
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfig) GetRiskAlertLevel() *string {
	return s.RiskAlertLevel
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfig) GetRiskConfig() []*UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigRiskConfig {
	return s.RiskConfig
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfig) GetServiceAddress() *string {
	return s.ServiceAddress
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfig) SetBufferLimit(v int32) *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfig {
	s.BufferLimit = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfig) SetCheckRequest(v bool) *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfig {
	s.CheckRequest = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfig) SetCheckRequestImage(v bool) *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfig {
	s.CheckRequestImage = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfig) SetCheckResponse(v bool) *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfig {
	s.CheckResponse = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfig) SetCheckResponseImage(v bool) *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfig {
	s.CheckResponseImage = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfig) SetConsumerRequestCheckService(v []*UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigConsumerRequestCheckService) *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfig {
	s.ConsumerRequestCheckService = v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfig) SetConsumerResponseCheckService(v []*UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigConsumerResponseCheckService) *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfig {
	s.ConsumerResponseCheckService = v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfig) SetConsumerRiskLevel(v []*UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigConsumerRiskLevel) *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfig {
	s.ConsumerRiskLevel = v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfig) SetPluginStatus(v *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigPluginStatus) *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfig {
	s.PluginStatus = v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfig) SetRequestCheckService(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfig {
	s.RequestCheckService = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfig) SetRequestImageCheckService(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfig {
	s.RequestImageCheckService = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfig) SetResponseCheckService(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfig {
	s.ResponseCheckService = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfig) SetResponseImageCheckService(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfig {
	s.ResponseImageCheckService = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfig) SetRiskAlertLevel(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfig {
	s.RiskAlertLevel = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfig) SetRiskConfig(v []*UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigRiskConfig) *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfig {
	s.RiskConfig = v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfig) SetServiceAddress(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfig {
	s.ServiceAddress = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfig) Validate() error {
	if s.ConsumerRequestCheckService != nil {
		for _, item := range s.ConsumerRequestCheckService {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ConsumerResponseCheckService != nil {
		for _, item := range s.ConsumerResponseCheckService {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ConsumerRiskLevel != nil {
		for _, item := range s.ConsumerRiskLevel {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PluginStatus != nil {
		if err := s.PluginStatus.Validate(); err != nil {
			return err
		}
	}
	if s.RiskConfig != nil {
		for _, item := range s.RiskConfig {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigConsumerRequestCheckService struct {
	MatchType                *string `json:"matchType,omitempty" xml:"matchType,omitempty"`
	ModalityType             *string `json:"modalityType,omitempty" xml:"modalityType,omitempty"`
	Name                     *string `json:"name,omitempty" xml:"name,omitempty"`
	RequestCheckService      *string `json:"requestCheckService,omitempty" xml:"requestCheckService,omitempty"`
	RequestImageCheckService *string `json:"requestImageCheckService,omitempty" xml:"requestImageCheckService,omitempty"`
}

func (s UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigConsumerRequestCheckService) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigConsumerRequestCheckService) GoString() string {
	return s.String()
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigConsumerRequestCheckService) GetMatchType() *string {
	return s.MatchType
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigConsumerRequestCheckService) GetModalityType() *string {
	return s.ModalityType
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigConsumerRequestCheckService) GetName() *string {
	return s.Name
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigConsumerRequestCheckService) GetRequestCheckService() *string {
	return s.RequestCheckService
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigConsumerRequestCheckService) GetRequestImageCheckService() *string {
	return s.RequestImageCheckService
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigConsumerRequestCheckService) SetMatchType(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigConsumerRequestCheckService {
	s.MatchType = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigConsumerRequestCheckService) SetModalityType(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigConsumerRequestCheckService {
	s.ModalityType = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigConsumerRequestCheckService) SetName(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigConsumerRequestCheckService {
	s.Name = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigConsumerRequestCheckService) SetRequestCheckService(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigConsumerRequestCheckService {
	s.RequestCheckService = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigConsumerRequestCheckService) SetRequestImageCheckService(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigConsumerRequestCheckService {
	s.RequestImageCheckService = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigConsumerRequestCheckService) Validate() error {
	return dara.Validate(s)
}

type UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigConsumerResponseCheckService struct {
	MatchType                 *string `json:"matchType,omitempty" xml:"matchType,omitempty"`
	ModalityType              *string `json:"modalityType,omitempty" xml:"modalityType,omitempty"`
	Name                      *string `json:"name,omitempty" xml:"name,omitempty"`
	ResponseCheckService      *string `json:"responseCheckService,omitempty" xml:"responseCheckService,omitempty"`
	ResponseImageCheckService *string `json:"responseImageCheckService,omitempty" xml:"responseImageCheckService,omitempty"`
}

func (s UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigConsumerResponseCheckService) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigConsumerResponseCheckService) GoString() string {
	return s.String()
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigConsumerResponseCheckService) GetMatchType() *string {
	return s.MatchType
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigConsumerResponseCheckService) GetModalityType() *string {
	return s.ModalityType
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigConsumerResponseCheckService) GetName() *string {
	return s.Name
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigConsumerResponseCheckService) GetResponseCheckService() *string {
	return s.ResponseCheckService
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigConsumerResponseCheckService) GetResponseImageCheckService() *string {
	return s.ResponseImageCheckService
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigConsumerResponseCheckService) SetMatchType(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigConsumerResponseCheckService {
	s.MatchType = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigConsumerResponseCheckService) SetModalityType(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigConsumerResponseCheckService {
	s.ModalityType = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigConsumerResponseCheckService) SetName(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigConsumerResponseCheckService {
	s.Name = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigConsumerResponseCheckService) SetResponseCheckService(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigConsumerResponseCheckService {
	s.ResponseCheckService = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigConsumerResponseCheckService) SetResponseImageCheckService(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigConsumerResponseCheckService {
	s.ResponseImageCheckService = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigConsumerResponseCheckService) Validate() error {
	return dara.Validate(s)
}

type UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigConsumerRiskLevel struct {
	Level     *string `json:"level,omitempty" xml:"level,omitempty"`
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty"`
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
	Type      *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigConsumerRiskLevel) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigConsumerRiskLevel) GoString() string {
	return s.String()
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigConsumerRiskLevel) GetLevel() *string {
	return s.Level
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigConsumerRiskLevel) GetMatchType() *string {
	return s.MatchType
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigConsumerRiskLevel) GetName() *string {
	return s.Name
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigConsumerRiskLevel) GetType() *string {
	return s.Type
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigConsumerRiskLevel) SetLevel(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigConsumerRiskLevel {
	s.Level = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigConsumerRiskLevel) SetMatchType(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigConsumerRiskLevel {
	s.MatchType = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigConsumerRiskLevel) SetName(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigConsumerRiskLevel {
	s.Name = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigConsumerRiskLevel) SetType(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigConsumerRiskLevel {
	s.Type = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigConsumerRiskLevel) Validate() error {
	return dara.Validate(s)
}

type UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigPluginStatus struct {
	ErrorLogs      map[string]*string `json:"errorLogs,omitempty" xml:"errorLogs,omitempty"`
	PluginId       *string            `json:"pluginId,omitempty" xml:"pluginId,omitempty"`
	ServiceHealthy *bool              `json:"serviceHealthy,omitempty" xml:"serviceHealthy,omitempty"`
}

func (s UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigPluginStatus) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigPluginStatus) GoString() string {
	return s.String()
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigPluginStatus) GetErrorLogs() map[string]*string {
	return s.ErrorLogs
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigPluginStatus) GetPluginId() *string {
	return s.PluginId
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigPluginStatus) GetServiceHealthy() *bool {
	return s.ServiceHealthy
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigPluginStatus) SetErrorLogs(v map[string]*string) *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigPluginStatus {
	s.ErrorLogs = v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigPluginStatus) SetPluginId(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigPluginStatus {
	s.PluginId = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigPluginStatus) SetServiceHealthy(v bool) *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigPluginStatus {
	s.ServiceHealthy = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigPluginStatus) Validate() error {
	return dara.Validate(s)
}

type UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigRiskConfig struct {
	ConsumerRules *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigRiskConfigConsumerRules `json:"consumerRules,omitempty" xml:"consumerRules,omitempty" type:"Struct"`
	Level         *string                                                                             `json:"level,omitempty" xml:"level,omitempty"`
	Type          *string                                                                             `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigRiskConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigRiskConfig) GoString() string {
	return s.String()
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigRiskConfig) GetConsumerRules() *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigRiskConfigConsumerRules {
	return s.ConsumerRules
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigRiskConfig) GetLevel() *string {
	return s.Level
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigRiskConfig) GetType() *string {
	return s.Type
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigRiskConfig) SetConsumerRules(v *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigRiskConfigConsumerRules) *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigRiskConfig {
	s.ConsumerRules = v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigRiskConfig) SetLevel(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigRiskConfig {
	s.Level = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigRiskConfig) SetType(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigRiskConfig {
	s.Type = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigRiskConfig) Validate() error {
	if s.ConsumerRules != nil {
		if err := s.ConsumerRules.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigRiskConfigConsumerRules struct {
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty"`
	Pattern   *string `json:"pattern,omitempty" xml:"pattern,omitempty"`
}

func (s UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigRiskConfigConsumerRules) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigRiskConfigConsumerRules) GoString() string {
	return s.String()
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigRiskConfigConsumerRules) GetMatchType() *string {
	return s.MatchType
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigRiskConfigConsumerRules) GetPattern() *string {
	return s.Pattern
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigRiskConfigConsumerRules) SetMatchType(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigRiskConfigConsumerRules {
	s.MatchType = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigRiskConfigConsumerRules) SetPattern(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigRiskConfigConsumerRules {
	s.Pattern = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiSecurityGuardConfigRiskConfigConsumerRules) Validate() error {
	return dara.Validate(s)
}

type UpdateHttpApiRouteRequestPolicyConfigsAiStatisticsConfig struct {
	LogRequestContent  *bool `json:"logRequestContent,omitempty" xml:"logRequestContent,omitempty"`
	LogResponseContent *bool `json:"logResponseContent,omitempty" xml:"logResponseContent,omitempty"`
}

func (s UpdateHttpApiRouteRequestPolicyConfigsAiStatisticsConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpApiRouteRequestPolicyConfigsAiStatisticsConfig) GoString() string {
	return s.String()
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiStatisticsConfig) GetLogRequestContent() *bool {
	return s.LogRequestContent
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiStatisticsConfig) GetLogResponseContent() *bool {
	return s.LogResponseContent
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiStatisticsConfig) SetLogRequestContent(v bool) *UpdateHttpApiRouteRequestPolicyConfigsAiStatisticsConfig {
	s.LogRequestContent = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiStatisticsConfig) SetLogResponseContent(v bool) *UpdateHttpApiRouteRequestPolicyConfigsAiStatisticsConfig {
	s.LogResponseContent = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiStatisticsConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfig struct {
	EnableGlobalRules *bool                                                                      `json:"enableGlobalRules,omitempty" xml:"enableGlobalRules,omitempty"`
	GlobalRules       []*UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigGlobalRules `json:"globalRules,omitempty" xml:"globalRules,omitempty" type:"Repeated"`
	PluginStatus      *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigPluginStatus  `json:"pluginStatus,omitempty" xml:"pluginStatus,omitempty" type:"Struct"`
	RedisConfig       *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigRedisConfig   `json:"redisConfig,omitempty" xml:"redisConfig,omitempty" type:"Struct"`
	Rules             []*UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigRules       `json:"rules,omitempty" xml:"rules,omitempty" type:"Repeated"`
}

func (s UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfig) GoString() string {
	return s.String()
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfig) GetEnableGlobalRules() *bool {
	return s.EnableGlobalRules
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfig) GetGlobalRules() []*UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigGlobalRules {
	return s.GlobalRules
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfig) GetPluginStatus() *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigPluginStatus {
	return s.PluginStatus
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfig) GetRedisConfig() *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigRedisConfig {
	return s.RedisConfig
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfig) GetRules() []*UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigRules {
	return s.Rules
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfig) SetEnableGlobalRules(v bool) *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfig {
	s.EnableGlobalRules = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfig) SetGlobalRules(v []*UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigGlobalRules) *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfig {
	s.GlobalRules = v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfig) SetPluginStatus(v *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigPluginStatus) *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfig {
	s.PluginStatus = v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfig) SetRedisConfig(v *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigRedisConfig) *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfig {
	s.RedisConfig = v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfig) SetRules(v []*UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigRules) *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfig {
	s.Rules = v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfig) Validate() error {
	if s.GlobalRules != nil {
		for _, item := range s.GlobalRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PluginStatus != nil {
		if err := s.PluginStatus.Validate(); err != nil {
			return err
		}
	}
	if s.RedisConfig != nil {
		if err := s.RedisConfig.Validate(); err != nil {
			return err
		}
	}
	if s.Rules != nil {
		for _, item := range s.Rules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigGlobalRules struct {
	LimitMode  *string `json:"limitMode,omitempty" xml:"limitMode,omitempty"`
	LimitType  *string `json:"limitType,omitempty" xml:"limitType,omitempty"`
	LimitValue *int32  `json:"limitValue,omitempty" xml:"limitValue,omitempty"`
	MatchKey   *string `json:"matchKey,omitempty" xml:"matchKey,omitempty"`
	MatchType  *string `json:"matchType,omitempty" xml:"matchType,omitempty"`
	MatchValue *string `json:"matchValue,omitempty" xml:"matchValue,omitempty"`
}

func (s UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigGlobalRules) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigGlobalRules) GoString() string {
	return s.String()
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigGlobalRules) GetLimitMode() *string {
	return s.LimitMode
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigGlobalRules) GetLimitType() *string {
	return s.LimitType
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigGlobalRules) GetLimitValue() *int32 {
	return s.LimitValue
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigGlobalRules) GetMatchKey() *string {
	return s.MatchKey
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigGlobalRules) GetMatchType() *string {
	return s.MatchType
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigGlobalRules) GetMatchValue() *string {
	return s.MatchValue
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigGlobalRules) SetLimitMode(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigGlobalRules {
	s.LimitMode = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigGlobalRules) SetLimitType(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigGlobalRules {
	s.LimitType = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigGlobalRules) SetLimitValue(v int32) *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigGlobalRules {
	s.LimitValue = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigGlobalRules) SetMatchKey(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigGlobalRules {
	s.MatchKey = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigGlobalRules) SetMatchType(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigGlobalRules {
	s.MatchType = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigGlobalRules) SetMatchValue(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigGlobalRules {
	s.MatchValue = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigGlobalRules) Validate() error {
	return dara.Validate(s)
}

type UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigPluginStatus struct {
	ErrorLogs      map[string]*string `json:"errorLogs,omitempty" xml:"errorLogs,omitempty"`
	PluginId       *string            `json:"pluginId,omitempty" xml:"pluginId,omitempty"`
	ServiceHealthy *bool              `json:"serviceHealthy,omitempty" xml:"serviceHealthy,omitempty"`
}

func (s UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigPluginStatus) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigPluginStatus) GoString() string {
	return s.String()
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigPluginStatus) GetErrorLogs() map[string]*string {
	return s.ErrorLogs
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigPluginStatus) GetPluginId() *string {
	return s.PluginId
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigPluginStatus) GetServiceHealthy() *bool {
	return s.ServiceHealthy
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigPluginStatus) SetErrorLogs(v map[string]*string) *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigPluginStatus {
	s.ErrorLogs = v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigPluginStatus) SetPluginId(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigPluginStatus {
	s.PluginId = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigPluginStatus) SetServiceHealthy(v bool) *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigPluginStatus {
	s.ServiceHealthy = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigPluginStatus) Validate() error {
	return dara.Validate(s)
}

type UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigRedisConfig struct {
	DatabaseNumber *int32  `json:"databaseNumber,omitempty" xml:"databaseNumber,omitempty"`
	Host           *string `json:"host,omitempty" xml:"host,omitempty"`
	Password       *string `json:"password,omitempty" xml:"password,omitempty"`
	Port           *int32  `json:"port,omitempty" xml:"port,omitempty"`
	Timeout        *int32  `json:"timeout,omitempty" xml:"timeout,omitempty"`
	Username       *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigRedisConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigRedisConfig) GoString() string {
	return s.String()
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigRedisConfig) GetDatabaseNumber() *int32 {
	return s.DatabaseNumber
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigRedisConfig) GetHost() *string {
	return s.Host
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigRedisConfig) GetPassword() *string {
	return s.Password
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigRedisConfig) GetPort() *int32 {
	return s.Port
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigRedisConfig) GetTimeout() *int32 {
	return s.Timeout
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigRedisConfig) GetUsername() *string {
	return s.Username
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigRedisConfig) SetDatabaseNumber(v int32) *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigRedisConfig {
	s.DatabaseNumber = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigRedisConfig) SetHost(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigRedisConfig {
	s.Host = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigRedisConfig) SetPassword(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigRedisConfig {
	s.Password = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigRedisConfig) SetPort(v int32) *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigRedisConfig {
	s.Port = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigRedisConfig) SetTimeout(v int32) *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigRedisConfig {
	s.Timeout = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigRedisConfig) SetUsername(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigRedisConfig {
	s.Username = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigRedisConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigRules struct {
	LimitMode  *string `json:"limitMode,omitempty" xml:"limitMode,omitempty"`
	LimitType  *string `json:"limitType,omitempty" xml:"limitType,omitempty"`
	LimitValue *int32  `json:"limitValue,omitempty" xml:"limitValue,omitempty"`
	MatchKey   *string `json:"matchKey,omitempty" xml:"matchKey,omitempty"`
	MatchType  *string `json:"matchType,omitempty" xml:"matchType,omitempty"`
	MatchValue *string `json:"matchValue,omitempty" xml:"matchValue,omitempty"`
}

func (s UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigRules) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigRules) GoString() string {
	return s.String()
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigRules) GetLimitMode() *string {
	return s.LimitMode
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigRules) GetLimitType() *string {
	return s.LimitType
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigRules) GetLimitValue() *int32 {
	return s.LimitValue
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigRules) GetMatchKey() *string {
	return s.MatchKey
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigRules) GetMatchType() *string {
	return s.MatchType
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigRules) GetMatchValue() *string {
	return s.MatchValue
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigRules) SetLimitMode(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigRules {
	s.LimitMode = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigRules) SetLimitType(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigRules {
	s.LimitType = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigRules) SetLimitValue(v int32) *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigRules {
	s.LimitValue = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigRules) SetMatchKey(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigRules {
	s.MatchKey = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigRules) SetMatchType(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigRules {
	s.MatchType = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigRules) SetMatchValue(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigRules {
	s.MatchValue = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiTokenRateLimitConfigRules) Validate() error {
	return dara.Validate(s)
}

type UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfig struct {
	EnableConditions *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigEnableConditions `json:"enableConditions,omitempty" xml:"enableConditions,omitempty" type:"Struct"`
	PluginStatus     *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigPluginStatus     `json:"pluginStatus,omitempty" xml:"pluginStatus,omitempty" type:"Struct"`
	QueryRewriting   *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewriting   `json:"queryRewriting,omitempty" xml:"queryRewriting,omitempty" type:"Struct"`
	ToolReranking    *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigToolReranking    `json:"toolReranking,omitempty" xml:"toolReranking,omitempty" type:"Struct"`
}

func (s UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfig) GoString() string {
	return s.String()
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfig) GetEnableConditions() *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigEnableConditions {
	return s.EnableConditions
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfig) GetPluginStatus() *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigPluginStatus {
	return s.PluginStatus
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfig) GetQueryRewriting() *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewriting {
	return s.QueryRewriting
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfig) GetToolReranking() *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigToolReranking {
	return s.ToolReranking
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfig) SetEnableConditions(v *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigEnableConditions) *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfig {
	s.EnableConditions = v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfig) SetPluginStatus(v *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigPluginStatus) *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfig {
	s.PluginStatus = v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfig) SetQueryRewriting(v *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewriting) *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfig {
	s.QueryRewriting = v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfig) SetToolReranking(v *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigToolReranking) *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfig {
	s.ToolReranking = v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfig) Validate() error {
	if s.EnableConditions != nil {
		if err := s.EnableConditions.Validate(); err != nil {
			return err
		}
	}
	if s.PluginStatus != nil {
		if err := s.PluginStatus.Validate(); err != nil {
			return err
		}
	}
	if s.QueryRewriting != nil {
		if err := s.QueryRewriting.Validate(); err != nil {
			return err
		}
	}
	if s.ToolReranking != nil {
		if err := s.ToolReranking.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigEnableConditions struct {
	ToolCountThreshold *int32 `json:"toolCountThreshold,omitempty" xml:"toolCountThreshold,omitempty"`
}

func (s UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigEnableConditions) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigEnableConditions) GoString() string {
	return s.String()
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigEnableConditions) GetToolCountThreshold() *int32 {
	return s.ToolCountThreshold
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigEnableConditions) SetToolCountThreshold(v int32) *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigEnableConditions {
	s.ToolCountThreshold = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigEnableConditions) Validate() error {
	return dara.Validate(s)
}

type UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigPluginStatus struct {
	ErrorLogs      map[string]*string `json:"errorLogs,omitempty" xml:"errorLogs,omitempty"`
	PluginId       *string            `json:"pluginId,omitempty" xml:"pluginId,omitempty"`
	ServiceHealthy *bool              `json:"serviceHealthy,omitempty" xml:"serviceHealthy,omitempty"`
}

func (s UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigPluginStatus) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigPluginStatus) GoString() string {
	return s.String()
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigPluginStatus) GetErrorLogs() map[string]*string {
	return s.ErrorLogs
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigPluginStatus) GetPluginId() *string {
	return s.PluginId
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigPluginStatus) GetServiceHealthy() *bool {
	return s.ServiceHealthy
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigPluginStatus) SetErrorLogs(v map[string]*string) *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigPluginStatus {
	s.ErrorLogs = v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigPluginStatus) SetPluginId(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigPluginStatus {
	s.PluginId = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigPluginStatus) SetServiceHealthy(v bool) *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigPluginStatus {
	s.ServiceHealthy = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigPluginStatus) Validate() error {
	return dara.Validate(s)
}

type UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewriting struct {
	ContextSelection  *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewritingContextSelection  `json:"contextSelection,omitempty" xml:"contextSelection,omitempty" type:"Struct"`
	Enabled           *bool                                                                                       `json:"enabled,omitempty" xml:"enabled,omitempty"`
	FallbackStrategy  *string                                                                                     `json:"fallbackStrategy,omitempty" xml:"fallbackStrategy,omitempty"`
	MaxOutputTokens   *int32                                                                                      `json:"maxOutputTokens,omitempty" xml:"maxOutputTokens,omitempty"`
	ModelService      *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewritingModelService      `json:"modelService,omitempty" xml:"modelService,omitempty" type:"Struct"`
	PromptConfig      *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewritingPromptConfig      `json:"promptConfig,omitempty" xml:"promptConfig,omitempty" type:"Struct"`
	TriggerConditions *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewritingTriggerConditions `json:"triggerConditions,omitempty" xml:"triggerConditions,omitempty" type:"Struct"`
}

func (s UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewriting) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewriting) GoString() string {
	return s.String()
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewriting) GetContextSelection() *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewritingContextSelection {
	return s.ContextSelection
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewriting) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewriting) GetFallbackStrategy() *string {
	return s.FallbackStrategy
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewriting) GetMaxOutputTokens() *int32 {
	return s.MaxOutputTokens
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewriting) GetModelService() *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewritingModelService {
	return s.ModelService
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewriting) GetPromptConfig() *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewritingPromptConfig {
	return s.PromptConfig
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewriting) GetTriggerConditions() *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewritingTriggerConditions {
	return s.TriggerConditions
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewriting) SetContextSelection(v *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewritingContextSelection) *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewriting {
	s.ContextSelection = v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewriting) SetEnabled(v bool) *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewriting {
	s.Enabled = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewriting) SetFallbackStrategy(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewriting {
	s.FallbackStrategy = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewriting) SetMaxOutputTokens(v int32) *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewriting {
	s.MaxOutputTokens = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewriting) SetModelService(v *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewritingModelService) *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewriting {
	s.ModelService = v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewriting) SetPromptConfig(v *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewritingPromptConfig) *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewriting {
	s.PromptConfig = v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewriting) SetTriggerConditions(v *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewritingTriggerConditions) *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewriting {
	s.TriggerConditions = v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewriting) Validate() error {
	if s.ContextSelection != nil {
		if err := s.ContextSelection.Validate(); err != nil {
			return err
		}
	}
	if s.ModelService != nil {
		if err := s.ModelService.Validate(); err != nil {
			return err
		}
	}
	if s.PromptConfig != nil {
		if err := s.PromptConfig.Validate(); err != nil {
			return err
		}
	}
	if s.TriggerConditions != nil {
		if err := s.TriggerConditions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewritingContextSelection struct {
	Type  *string `json:"type,omitempty" xml:"type,omitempty"`
	Value *int32  `json:"value,omitempty" xml:"value,omitempty"`
}

func (s UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewritingContextSelection) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewritingContextSelection) GoString() string {
	return s.String()
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewritingContextSelection) GetType() *string {
	return s.Type
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewritingContextSelection) GetValue() *int32 {
	return s.Value
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewritingContextSelection) SetType(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewritingContextSelection {
	s.Type = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewritingContextSelection) SetValue(v int32) *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewritingContextSelection {
	s.Value = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewritingContextSelection) Validate() error {
	return dara.Validate(s)
}

type UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewritingModelService struct {
	ModelName          *string `json:"modelName,omitempty" xml:"modelName,omitempty"`
	ServiceId          *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	TimeoutMillisecond *int32  `json:"timeoutMillisecond,omitempty" xml:"timeoutMillisecond,omitempty"`
}

func (s UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewritingModelService) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewritingModelService) GoString() string {
	return s.String()
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewritingModelService) GetModelName() *string {
	return s.ModelName
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewritingModelService) GetServiceId() *string {
	return s.ServiceId
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewritingModelService) GetTimeoutMillisecond() *int32 {
	return s.TimeoutMillisecond
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewritingModelService) SetModelName(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewritingModelService {
	s.ModelName = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewritingModelService) SetServiceId(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewritingModelService {
	s.ServiceId = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewritingModelService) SetTimeoutMillisecond(v int32) *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewritingModelService {
	s.TimeoutMillisecond = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewritingModelService) Validate() error {
	return dara.Validate(s)
}

type UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewritingPromptConfig struct {
	CustomPrompt *string `json:"customPrompt,omitempty" xml:"customPrompt,omitempty"`
	Type         *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewritingPromptConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewritingPromptConfig) GoString() string {
	return s.String()
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewritingPromptConfig) GetCustomPrompt() *string {
	return s.CustomPrompt
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewritingPromptConfig) GetType() *string {
	return s.Type
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewritingPromptConfig) SetCustomPrompt(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewritingPromptConfig {
	s.CustomPrompt = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewritingPromptConfig) SetType(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewritingPromptConfig {
	s.Type = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewritingPromptConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewritingTriggerConditions struct {
	MessageCountThreshold *int32 `json:"messageCountThreshold,omitempty" xml:"messageCountThreshold,omitempty"`
}

func (s UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewritingTriggerConditions) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewritingTriggerConditions) GoString() string {
	return s.String()
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewritingTriggerConditions) GetMessageCountThreshold() *int32 {
	return s.MessageCountThreshold
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewritingTriggerConditions) SetMessageCountThreshold(v int32) *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewritingTriggerConditions {
	s.MessageCountThreshold = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigQueryRewritingTriggerConditions) Validate() error {
	return dara.Validate(s)
}

type UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigToolReranking struct {
	FallbackStrategy *string                                                                               `json:"fallbackStrategy,omitempty" xml:"fallbackStrategy,omitempty"`
	FilteringMethod  *string                                                                               `json:"filteringMethod,omitempty" xml:"filteringMethod,omitempty"`
	ModelService     *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigToolRerankingModelService `json:"modelService,omitempty" xml:"modelService,omitempty" type:"Struct"`
	ScoreThreshold   *float32                                                                              `json:"scoreThreshold,omitempty" xml:"scoreThreshold,omitempty"`
	TopKPercent      *int32                                                                                `json:"topKPercent,omitempty" xml:"topKPercent,omitempty"`
	TopNCount        *int32                                                                                `json:"topNCount,omitempty" xml:"topNCount,omitempty"`
}

func (s UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigToolReranking) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigToolReranking) GoString() string {
	return s.String()
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigToolReranking) GetFallbackStrategy() *string {
	return s.FallbackStrategy
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigToolReranking) GetFilteringMethod() *string {
	return s.FilteringMethod
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigToolReranking) GetModelService() *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigToolRerankingModelService {
	return s.ModelService
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigToolReranking) GetScoreThreshold() *float32 {
	return s.ScoreThreshold
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigToolReranking) GetTopKPercent() *int32 {
	return s.TopKPercent
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigToolReranking) GetTopNCount() *int32 {
	return s.TopNCount
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigToolReranking) SetFallbackStrategy(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigToolReranking {
	s.FallbackStrategy = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigToolReranking) SetFilteringMethod(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigToolReranking {
	s.FilteringMethod = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigToolReranking) SetModelService(v *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigToolRerankingModelService) *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigToolReranking {
	s.ModelService = v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigToolReranking) SetScoreThreshold(v float32) *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigToolReranking {
	s.ScoreThreshold = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigToolReranking) SetTopKPercent(v int32) *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigToolReranking {
	s.TopKPercent = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigToolReranking) SetTopNCount(v int32) *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigToolReranking {
	s.TopNCount = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigToolReranking) Validate() error {
	if s.ModelService != nil {
		if err := s.ModelService.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigToolRerankingModelService struct {
	ModelName          *string `json:"modelName,omitempty" xml:"modelName,omitempty"`
	ServiceId          *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	TimeoutMillisecond *int32  `json:"timeoutMillisecond,omitempty" xml:"timeoutMillisecond,omitempty"`
}

func (s UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigToolRerankingModelService) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigToolRerankingModelService) GoString() string {
	return s.String()
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigToolRerankingModelService) GetModelName() *string {
	return s.ModelName
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigToolRerankingModelService) GetServiceId() *string {
	return s.ServiceId
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigToolRerankingModelService) GetTimeoutMillisecond() *int32 {
	return s.TimeoutMillisecond
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigToolRerankingModelService) SetModelName(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigToolRerankingModelService {
	s.ModelName = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigToolRerankingModelService) SetServiceId(v string) *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigToolRerankingModelService {
	s.ServiceId = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigToolRerankingModelService) SetTimeoutMillisecond(v int32) *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigToolRerankingModelService {
	s.TimeoutMillisecond = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsAiToolSelectionConfigToolRerankingModelService) Validate() error {
	return dara.Validate(s)
}

type UpdateHttpApiRouteRequestPolicyConfigsSemanticRouterConfig struct {
	TimeoutMillisecond *int32 `json:"timeoutMillisecond,omitempty" xml:"timeoutMillisecond,omitempty"`
}

func (s UpdateHttpApiRouteRequestPolicyConfigsSemanticRouterConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpApiRouteRequestPolicyConfigsSemanticRouterConfig) GoString() string {
	return s.String()
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsSemanticRouterConfig) GetTimeoutMillisecond() *int32 {
	return s.TimeoutMillisecond
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsSemanticRouterConfig) SetTimeoutMillisecond(v int32) *UpdateHttpApiRouteRequestPolicyConfigsSemanticRouterConfig {
	s.TimeoutMillisecond = &v
	return s
}

func (s *UpdateHttpApiRouteRequestPolicyConfigsSemanticRouterConfig) Validate() error {
	return dara.Validate(s)
}
