// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHttpApiPolicyConfigs interface {
	dara.Model
	String() string
	GoString() string
	SetAiCacheConfig(v *HttpApiPolicyConfigsAiCacheConfig) *HttpApiPolicyConfigs
	GetAiCacheConfig() *HttpApiPolicyConfigsAiCacheConfig
	SetAiFallbackConfig(v *HttpApiPolicyConfigsAiFallbackConfig) *HttpApiPolicyConfigs
	GetAiFallbackConfig() *HttpApiPolicyConfigsAiFallbackConfig
	SetAiNetworkSearchConfig(v *HttpApiPolicyConfigsAiNetworkSearchConfig) *HttpApiPolicyConfigs
	GetAiNetworkSearchConfig() *HttpApiPolicyConfigsAiNetworkSearchConfig
	SetAiSecurityGuardConfig(v *HttpApiPolicyConfigsAiSecurityGuardConfig) *HttpApiPolicyConfigs
	GetAiSecurityGuardConfig() *HttpApiPolicyConfigsAiSecurityGuardConfig
	SetAiStatisticsConfig(v *HttpApiPolicyConfigsAiStatisticsConfig) *HttpApiPolicyConfigs
	GetAiStatisticsConfig() *HttpApiPolicyConfigsAiStatisticsConfig
	SetAiTokenRateLimitConfig(v *HttpApiPolicyConfigsAiTokenRateLimitConfig) *HttpApiPolicyConfigs
	GetAiTokenRateLimitConfig() *HttpApiPolicyConfigsAiTokenRateLimitConfig
	SetAiToolSelectionConfig(v *HttpApiPolicyConfigsAiToolSelectionConfig) *HttpApiPolicyConfigs
	GetAiToolSelectionConfig() *HttpApiPolicyConfigsAiToolSelectionConfig
	SetEnable(v bool) *HttpApiPolicyConfigs
	GetEnable() *bool
	SetSemanticRouterConfig(v *HttpApiPolicyConfigsSemanticRouterConfig) *HttpApiPolicyConfigs
	GetSemanticRouterConfig() *HttpApiPolicyConfigsSemanticRouterConfig
	SetType(v string) *HttpApiPolicyConfigs
	GetType() *string
}

type HttpApiPolicyConfigs struct {
	AiCacheConfig          *HttpApiPolicyConfigsAiCacheConfig          `json:"aiCacheConfig,omitempty" xml:"aiCacheConfig,omitempty" type:"Struct"`
	AiFallbackConfig       *HttpApiPolicyConfigsAiFallbackConfig       `json:"aiFallbackConfig,omitempty" xml:"aiFallbackConfig,omitempty" type:"Struct"`
	AiNetworkSearchConfig  *HttpApiPolicyConfigsAiNetworkSearchConfig  `json:"aiNetworkSearchConfig,omitempty" xml:"aiNetworkSearchConfig,omitempty" type:"Struct"`
	AiSecurityGuardConfig  *HttpApiPolicyConfigsAiSecurityGuardConfig  `json:"aiSecurityGuardConfig,omitempty" xml:"aiSecurityGuardConfig,omitempty" type:"Struct"`
	AiStatisticsConfig     *HttpApiPolicyConfigsAiStatisticsConfig     `json:"aiStatisticsConfig,omitempty" xml:"aiStatisticsConfig,omitempty" type:"Struct"`
	AiTokenRateLimitConfig *HttpApiPolicyConfigsAiTokenRateLimitConfig `json:"aiTokenRateLimitConfig,omitempty" xml:"aiTokenRateLimitConfig,omitempty" type:"Struct"`
	AiToolSelectionConfig  *HttpApiPolicyConfigsAiToolSelectionConfig  `json:"aiToolSelectionConfig,omitempty" xml:"aiToolSelectionConfig,omitempty" type:"Struct"`
	Enable                 *bool                                       `json:"enable,omitempty" xml:"enable,omitempty"`
	SemanticRouterConfig   *HttpApiPolicyConfigsSemanticRouterConfig   `json:"semanticRouterConfig,omitempty" xml:"semanticRouterConfig,omitempty" type:"Struct"`
	Type                   *string                                     `json:"type,omitempty" xml:"type,omitempty"`
}

func (s HttpApiPolicyConfigs) String() string {
	return dara.Prettify(s)
}

func (s HttpApiPolicyConfigs) GoString() string {
	return s.String()
}

func (s *HttpApiPolicyConfigs) GetAiCacheConfig() *HttpApiPolicyConfigsAiCacheConfig {
	return s.AiCacheConfig
}

func (s *HttpApiPolicyConfigs) GetAiFallbackConfig() *HttpApiPolicyConfigsAiFallbackConfig {
	return s.AiFallbackConfig
}

func (s *HttpApiPolicyConfigs) GetAiNetworkSearchConfig() *HttpApiPolicyConfigsAiNetworkSearchConfig {
	return s.AiNetworkSearchConfig
}

func (s *HttpApiPolicyConfigs) GetAiSecurityGuardConfig() *HttpApiPolicyConfigsAiSecurityGuardConfig {
	return s.AiSecurityGuardConfig
}

func (s *HttpApiPolicyConfigs) GetAiStatisticsConfig() *HttpApiPolicyConfigsAiStatisticsConfig {
	return s.AiStatisticsConfig
}

func (s *HttpApiPolicyConfigs) GetAiTokenRateLimitConfig() *HttpApiPolicyConfigsAiTokenRateLimitConfig {
	return s.AiTokenRateLimitConfig
}

func (s *HttpApiPolicyConfigs) GetAiToolSelectionConfig() *HttpApiPolicyConfigsAiToolSelectionConfig {
	return s.AiToolSelectionConfig
}

func (s *HttpApiPolicyConfigs) GetEnable() *bool {
	return s.Enable
}

func (s *HttpApiPolicyConfigs) GetSemanticRouterConfig() *HttpApiPolicyConfigsSemanticRouterConfig {
	return s.SemanticRouterConfig
}

func (s *HttpApiPolicyConfigs) GetType() *string {
	return s.Type
}

func (s *HttpApiPolicyConfigs) SetAiCacheConfig(v *HttpApiPolicyConfigsAiCacheConfig) *HttpApiPolicyConfigs {
	s.AiCacheConfig = v
	return s
}

func (s *HttpApiPolicyConfigs) SetAiFallbackConfig(v *HttpApiPolicyConfigsAiFallbackConfig) *HttpApiPolicyConfigs {
	s.AiFallbackConfig = v
	return s
}

func (s *HttpApiPolicyConfigs) SetAiNetworkSearchConfig(v *HttpApiPolicyConfigsAiNetworkSearchConfig) *HttpApiPolicyConfigs {
	s.AiNetworkSearchConfig = v
	return s
}

func (s *HttpApiPolicyConfigs) SetAiSecurityGuardConfig(v *HttpApiPolicyConfigsAiSecurityGuardConfig) *HttpApiPolicyConfigs {
	s.AiSecurityGuardConfig = v
	return s
}

func (s *HttpApiPolicyConfigs) SetAiStatisticsConfig(v *HttpApiPolicyConfigsAiStatisticsConfig) *HttpApiPolicyConfigs {
	s.AiStatisticsConfig = v
	return s
}

func (s *HttpApiPolicyConfigs) SetAiTokenRateLimitConfig(v *HttpApiPolicyConfigsAiTokenRateLimitConfig) *HttpApiPolicyConfigs {
	s.AiTokenRateLimitConfig = v
	return s
}

func (s *HttpApiPolicyConfigs) SetAiToolSelectionConfig(v *HttpApiPolicyConfigsAiToolSelectionConfig) *HttpApiPolicyConfigs {
	s.AiToolSelectionConfig = v
	return s
}

func (s *HttpApiPolicyConfigs) SetEnable(v bool) *HttpApiPolicyConfigs {
	s.Enable = &v
	return s
}

func (s *HttpApiPolicyConfigs) SetSemanticRouterConfig(v *HttpApiPolicyConfigsSemanticRouterConfig) *HttpApiPolicyConfigs {
	s.SemanticRouterConfig = v
	return s
}

func (s *HttpApiPolicyConfigs) SetType(v string) *HttpApiPolicyConfigs {
	s.Type = &v
	return s
}

func (s *HttpApiPolicyConfigs) Validate() error {
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

type HttpApiPolicyConfigsAiCacheConfig struct {
	CacheKeyStrategy *string                                           `json:"cacheKeyStrategy,omitempty" xml:"cacheKeyStrategy,omitempty"`
	CacheMode        *string                                           `json:"cacheMode,omitempty" xml:"cacheMode,omitempty"`
	CacheTTL         *int32                                            `json:"cacheTTL,omitempty" xml:"cacheTTL,omitempty"`
	EmbeddingConfig  *HttpApiPolicyConfigsAiCacheConfigEmbeddingConfig `json:"embeddingConfig,omitempty" xml:"embeddingConfig,omitempty" type:"Struct"`
	PluginStatus     *HttpApiPolicyConfigsAiCacheConfigPluginStatus    `json:"pluginStatus,omitempty" xml:"pluginStatus,omitempty" type:"Struct"`
	RedisConfig      *HttpApiPolicyConfigsAiCacheConfigRedisConfig     `json:"redisConfig,omitempty" xml:"redisConfig,omitempty" type:"Struct"`
	VectorConfig     *HttpApiPolicyConfigsAiCacheConfigVectorConfig    `json:"vectorConfig,omitempty" xml:"vectorConfig,omitempty" type:"Struct"`
}

func (s HttpApiPolicyConfigsAiCacheConfig) String() string {
	return dara.Prettify(s)
}

func (s HttpApiPolicyConfigsAiCacheConfig) GoString() string {
	return s.String()
}

func (s *HttpApiPolicyConfigsAiCacheConfig) GetCacheKeyStrategy() *string {
	return s.CacheKeyStrategy
}

func (s *HttpApiPolicyConfigsAiCacheConfig) GetCacheMode() *string {
	return s.CacheMode
}

func (s *HttpApiPolicyConfigsAiCacheConfig) GetCacheTTL() *int32 {
	return s.CacheTTL
}

func (s *HttpApiPolicyConfigsAiCacheConfig) GetEmbeddingConfig() *HttpApiPolicyConfigsAiCacheConfigEmbeddingConfig {
	return s.EmbeddingConfig
}

func (s *HttpApiPolicyConfigsAiCacheConfig) GetPluginStatus() *HttpApiPolicyConfigsAiCacheConfigPluginStatus {
	return s.PluginStatus
}

func (s *HttpApiPolicyConfigsAiCacheConfig) GetRedisConfig() *HttpApiPolicyConfigsAiCacheConfigRedisConfig {
	return s.RedisConfig
}

func (s *HttpApiPolicyConfigsAiCacheConfig) GetVectorConfig() *HttpApiPolicyConfigsAiCacheConfigVectorConfig {
	return s.VectorConfig
}

func (s *HttpApiPolicyConfigsAiCacheConfig) SetCacheKeyStrategy(v string) *HttpApiPolicyConfigsAiCacheConfig {
	s.CacheKeyStrategy = &v
	return s
}

func (s *HttpApiPolicyConfigsAiCacheConfig) SetCacheMode(v string) *HttpApiPolicyConfigsAiCacheConfig {
	s.CacheMode = &v
	return s
}

func (s *HttpApiPolicyConfigsAiCacheConfig) SetCacheTTL(v int32) *HttpApiPolicyConfigsAiCacheConfig {
	s.CacheTTL = &v
	return s
}

func (s *HttpApiPolicyConfigsAiCacheConfig) SetEmbeddingConfig(v *HttpApiPolicyConfigsAiCacheConfigEmbeddingConfig) *HttpApiPolicyConfigsAiCacheConfig {
	s.EmbeddingConfig = v
	return s
}

func (s *HttpApiPolicyConfigsAiCacheConfig) SetPluginStatus(v *HttpApiPolicyConfigsAiCacheConfigPluginStatus) *HttpApiPolicyConfigsAiCacheConfig {
	s.PluginStatus = v
	return s
}

func (s *HttpApiPolicyConfigsAiCacheConfig) SetRedisConfig(v *HttpApiPolicyConfigsAiCacheConfigRedisConfig) *HttpApiPolicyConfigsAiCacheConfig {
	s.RedisConfig = v
	return s
}

func (s *HttpApiPolicyConfigsAiCacheConfig) SetVectorConfig(v *HttpApiPolicyConfigsAiCacheConfigVectorConfig) *HttpApiPolicyConfigsAiCacheConfig {
	s.VectorConfig = v
	return s
}

func (s *HttpApiPolicyConfigsAiCacheConfig) Validate() error {
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

type HttpApiPolicyConfigsAiCacheConfigEmbeddingConfig struct {
	ModelName *string `json:"modelName,omitempty" xml:"modelName,omitempty"`
	ServiceId *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	Timeout   *int32  `json:"timeout,omitempty" xml:"timeout,omitempty"`
	Type      *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s HttpApiPolicyConfigsAiCacheConfigEmbeddingConfig) String() string {
	return dara.Prettify(s)
}

func (s HttpApiPolicyConfigsAiCacheConfigEmbeddingConfig) GoString() string {
	return s.String()
}

func (s *HttpApiPolicyConfigsAiCacheConfigEmbeddingConfig) GetModelName() *string {
	return s.ModelName
}

func (s *HttpApiPolicyConfigsAiCacheConfigEmbeddingConfig) GetServiceId() *string {
	return s.ServiceId
}

func (s *HttpApiPolicyConfigsAiCacheConfigEmbeddingConfig) GetTimeout() *int32 {
	return s.Timeout
}

func (s *HttpApiPolicyConfigsAiCacheConfigEmbeddingConfig) GetType() *string {
	return s.Type
}

func (s *HttpApiPolicyConfigsAiCacheConfigEmbeddingConfig) SetModelName(v string) *HttpApiPolicyConfigsAiCacheConfigEmbeddingConfig {
	s.ModelName = &v
	return s
}

func (s *HttpApiPolicyConfigsAiCacheConfigEmbeddingConfig) SetServiceId(v string) *HttpApiPolicyConfigsAiCacheConfigEmbeddingConfig {
	s.ServiceId = &v
	return s
}

func (s *HttpApiPolicyConfigsAiCacheConfigEmbeddingConfig) SetTimeout(v int32) *HttpApiPolicyConfigsAiCacheConfigEmbeddingConfig {
	s.Timeout = &v
	return s
}

func (s *HttpApiPolicyConfigsAiCacheConfigEmbeddingConfig) SetType(v string) *HttpApiPolicyConfigsAiCacheConfigEmbeddingConfig {
	s.Type = &v
	return s
}

func (s *HttpApiPolicyConfigsAiCacheConfigEmbeddingConfig) Validate() error {
	return dara.Validate(s)
}

type HttpApiPolicyConfigsAiCacheConfigPluginStatus struct {
	ErrorLogs      map[string]*string `json:"errorLogs,omitempty" xml:"errorLogs,omitempty"`
	PluginId       *string            `json:"pluginId,omitempty" xml:"pluginId,omitempty"`
	ServiceHealthy *bool              `json:"serviceHealthy,omitempty" xml:"serviceHealthy,omitempty"`
}

func (s HttpApiPolicyConfigsAiCacheConfigPluginStatus) String() string {
	return dara.Prettify(s)
}

func (s HttpApiPolicyConfigsAiCacheConfigPluginStatus) GoString() string {
	return s.String()
}

func (s *HttpApiPolicyConfigsAiCacheConfigPluginStatus) GetErrorLogs() map[string]*string {
	return s.ErrorLogs
}

func (s *HttpApiPolicyConfigsAiCacheConfigPluginStatus) GetPluginId() *string {
	return s.PluginId
}

func (s *HttpApiPolicyConfigsAiCacheConfigPluginStatus) GetServiceHealthy() *bool {
	return s.ServiceHealthy
}

func (s *HttpApiPolicyConfigsAiCacheConfigPluginStatus) SetErrorLogs(v map[string]*string) *HttpApiPolicyConfigsAiCacheConfigPluginStatus {
	s.ErrorLogs = v
	return s
}

func (s *HttpApiPolicyConfigsAiCacheConfigPluginStatus) SetPluginId(v string) *HttpApiPolicyConfigsAiCacheConfigPluginStatus {
	s.PluginId = &v
	return s
}

func (s *HttpApiPolicyConfigsAiCacheConfigPluginStatus) SetServiceHealthy(v bool) *HttpApiPolicyConfigsAiCacheConfigPluginStatus {
	s.ServiceHealthy = &v
	return s
}

func (s *HttpApiPolicyConfigsAiCacheConfigPluginStatus) Validate() error {
	return dara.Validate(s)
}

type HttpApiPolicyConfigsAiCacheConfigRedisConfig struct {
	DatabaseNumber *int32  `json:"databaseNumber,omitempty" xml:"databaseNumber,omitempty"`
	Host           *string `json:"host,omitempty" xml:"host,omitempty"`
	Password       *string `json:"password,omitempty" xml:"password,omitempty"`
	Port           *int32  `json:"port,omitempty" xml:"port,omitempty"`
	Timeout        *int32  `json:"timeout,omitempty" xml:"timeout,omitempty"`
	Username       *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s HttpApiPolicyConfigsAiCacheConfigRedisConfig) String() string {
	return dara.Prettify(s)
}

func (s HttpApiPolicyConfigsAiCacheConfigRedisConfig) GoString() string {
	return s.String()
}

func (s *HttpApiPolicyConfigsAiCacheConfigRedisConfig) GetDatabaseNumber() *int32 {
	return s.DatabaseNumber
}

func (s *HttpApiPolicyConfigsAiCacheConfigRedisConfig) GetHost() *string {
	return s.Host
}

func (s *HttpApiPolicyConfigsAiCacheConfigRedisConfig) GetPassword() *string {
	return s.Password
}

func (s *HttpApiPolicyConfigsAiCacheConfigRedisConfig) GetPort() *int32 {
	return s.Port
}

func (s *HttpApiPolicyConfigsAiCacheConfigRedisConfig) GetTimeout() *int32 {
	return s.Timeout
}

func (s *HttpApiPolicyConfigsAiCacheConfigRedisConfig) GetUsername() *string {
	return s.Username
}

func (s *HttpApiPolicyConfigsAiCacheConfigRedisConfig) SetDatabaseNumber(v int32) *HttpApiPolicyConfigsAiCacheConfigRedisConfig {
	s.DatabaseNumber = &v
	return s
}

func (s *HttpApiPolicyConfigsAiCacheConfigRedisConfig) SetHost(v string) *HttpApiPolicyConfigsAiCacheConfigRedisConfig {
	s.Host = &v
	return s
}

func (s *HttpApiPolicyConfigsAiCacheConfigRedisConfig) SetPassword(v string) *HttpApiPolicyConfigsAiCacheConfigRedisConfig {
	s.Password = &v
	return s
}

func (s *HttpApiPolicyConfigsAiCacheConfigRedisConfig) SetPort(v int32) *HttpApiPolicyConfigsAiCacheConfigRedisConfig {
	s.Port = &v
	return s
}

func (s *HttpApiPolicyConfigsAiCacheConfigRedisConfig) SetTimeout(v int32) *HttpApiPolicyConfigsAiCacheConfigRedisConfig {
	s.Timeout = &v
	return s
}

func (s *HttpApiPolicyConfigsAiCacheConfigRedisConfig) SetUsername(v string) *HttpApiPolicyConfigsAiCacheConfigRedisConfig {
	s.Username = &v
	return s
}

func (s *HttpApiPolicyConfigsAiCacheConfigRedisConfig) Validate() error {
	return dara.Validate(s)
}

type HttpApiPolicyConfigsAiCacheConfigVectorConfig struct {
	ApiKey       *string  `json:"apiKey,omitempty" xml:"apiKey,omitempty"`
	CollectionId *string  `json:"collectionId,omitempty" xml:"collectionId,omitempty"`
	ServiceHost  *string  `json:"serviceHost,omitempty" xml:"serviceHost,omitempty"`
	Threshold    *float32 `json:"threshold,omitempty" xml:"threshold,omitempty"`
	Timeout      *int32   `json:"timeout,omitempty" xml:"timeout,omitempty"`
	Type         *string  `json:"type,omitempty" xml:"type,omitempty"`
}

func (s HttpApiPolicyConfigsAiCacheConfigVectorConfig) String() string {
	return dara.Prettify(s)
}

func (s HttpApiPolicyConfigsAiCacheConfigVectorConfig) GoString() string {
	return s.String()
}

func (s *HttpApiPolicyConfigsAiCacheConfigVectorConfig) GetApiKey() *string {
	return s.ApiKey
}

func (s *HttpApiPolicyConfigsAiCacheConfigVectorConfig) GetCollectionId() *string {
	return s.CollectionId
}

func (s *HttpApiPolicyConfigsAiCacheConfigVectorConfig) GetServiceHost() *string {
	return s.ServiceHost
}

func (s *HttpApiPolicyConfigsAiCacheConfigVectorConfig) GetThreshold() *float32 {
	return s.Threshold
}

func (s *HttpApiPolicyConfigsAiCacheConfigVectorConfig) GetTimeout() *int32 {
	return s.Timeout
}

func (s *HttpApiPolicyConfigsAiCacheConfigVectorConfig) GetType() *string {
	return s.Type
}

func (s *HttpApiPolicyConfigsAiCacheConfigVectorConfig) SetApiKey(v string) *HttpApiPolicyConfigsAiCacheConfigVectorConfig {
	s.ApiKey = &v
	return s
}

func (s *HttpApiPolicyConfigsAiCacheConfigVectorConfig) SetCollectionId(v string) *HttpApiPolicyConfigsAiCacheConfigVectorConfig {
	s.CollectionId = &v
	return s
}

func (s *HttpApiPolicyConfigsAiCacheConfigVectorConfig) SetServiceHost(v string) *HttpApiPolicyConfigsAiCacheConfigVectorConfig {
	s.ServiceHost = &v
	return s
}

func (s *HttpApiPolicyConfigsAiCacheConfigVectorConfig) SetThreshold(v float32) *HttpApiPolicyConfigsAiCacheConfigVectorConfig {
	s.Threshold = &v
	return s
}

func (s *HttpApiPolicyConfigsAiCacheConfigVectorConfig) SetTimeout(v int32) *HttpApiPolicyConfigsAiCacheConfigVectorConfig {
	s.Timeout = &v
	return s
}

func (s *HttpApiPolicyConfigsAiCacheConfigVectorConfig) SetType(v string) *HttpApiPolicyConfigsAiCacheConfigVectorConfig {
	s.Type = &v
	return s
}

func (s *HttpApiPolicyConfigsAiCacheConfigVectorConfig) Validate() error {
	return dara.Validate(s)
}

type HttpApiPolicyConfigsAiFallbackConfig struct {
	OnlyRedirectUpstreamCode *bool                                                 `json:"onlyRedirectUpstreamCode,omitempty" xml:"onlyRedirectUpstreamCode,omitempty"`
	RouteEmbedded            *bool                                                 `json:"routeEmbedded,omitempty" xml:"routeEmbedded,omitempty"`
	ServiceConfigs           []*HttpApiPolicyConfigsAiFallbackConfigServiceConfigs `json:"serviceConfigs,omitempty" xml:"serviceConfigs,omitempty" type:"Repeated"`
}

func (s HttpApiPolicyConfigsAiFallbackConfig) String() string {
	return dara.Prettify(s)
}

func (s HttpApiPolicyConfigsAiFallbackConfig) GoString() string {
	return s.String()
}

func (s *HttpApiPolicyConfigsAiFallbackConfig) GetOnlyRedirectUpstreamCode() *bool {
	return s.OnlyRedirectUpstreamCode
}

func (s *HttpApiPolicyConfigsAiFallbackConfig) GetRouteEmbedded() *bool {
	return s.RouteEmbedded
}

func (s *HttpApiPolicyConfigsAiFallbackConfig) GetServiceConfigs() []*HttpApiPolicyConfigsAiFallbackConfigServiceConfigs {
	return s.ServiceConfigs
}

func (s *HttpApiPolicyConfigsAiFallbackConfig) SetOnlyRedirectUpstreamCode(v bool) *HttpApiPolicyConfigsAiFallbackConfig {
	s.OnlyRedirectUpstreamCode = &v
	return s
}

func (s *HttpApiPolicyConfigsAiFallbackConfig) SetRouteEmbedded(v bool) *HttpApiPolicyConfigsAiFallbackConfig {
	s.RouteEmbedded = &v
	return s
}

func (s *HttpApiPolicyConfigsAiFallbackConfig) SetServiceConfigs(v []*HttpApiPolicyConfigsAiFallbackConfigServiceConfigs) *HttpApiPolicyConfigsAiFallbackConfig {
	s.ServiceConfigs = v
	return s
}

func (s *HttpApiPolicyConfigsAiFallbackConfig) Validate() error {
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

type HttpApiPolicyConfigsAiFallbackConfigServiceConfigs struct {
	Name                 *string `json:"name,omitempty" xml:"name,omitempty"`
	PassThroughModelName *bool   `json:"passThroughModelName,omitempty" xml:"passThroughModelName,omitempty"`
	ServiceId            *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	TargetModelName      *string `json:"targetModelName,omitempty" xml:"targetModelName,omitempty"`
}

func (s HttpApiPolicyConfigsAiFallbackConfigServiceConfigs) String() string {
	return dara.Prettify(s)
}

func (s HttpApiPolicyConfigsAiFallbackConfigServiceConfigs) GoString() string {
	return s.String()
}

func (s *HttpApiPolicyConfigsAiFallbackConfigServiceConfigs) GetName() *string {
	return s.Name
}

func (s *HttpApiPolicyConfigsAiFallbackConfigServiceConfigs) GetPassThroughModelName() *bool {
	return s.PassThroughModelName
}

func (s *HttpApiPolicyConfigsAiFallbackConfigServiceConfigs) GetServiceId() *string {
	return s.ServiceId
}

func (s *HttpApiPolicyConfigsAiFallbackConfigServiceConfigs) GetTargetModelName() *string {
	return s.TargetModelName
}

func (s *HttpApiPolicyConfigsAiFallbackConfigServiceConfigs) SetName(v string) *HttpApiPolicyConfigsAiFallbackConfigServiceConfigs {
	s.Name = &v
	return s
}

func (s *HttpApiPolicyConfigsAiFallbackConfigServiceConfigs) SetPassThroughModelName(v bool) *HttpApiPolicyConfigsAiFallbackConfigServiceConfigs {
	s.PassThroughModelName = &v
	return s
}

func (s *HttpApiPolicyConfigsAiFallbackConfigServiceConfigs) SetServiceId(v string) *HttpApiPolicyConfigsAiFallbackConfigServiceConfigs {
	s.ServiceId = &v
	return s
}

func (s *HttpApiPolicyConfigsAiFallbackConfigServiceConfigs) SetTargetModelName(v string) *HttpApiPolicyConfigsAiFallbackConfigServiceConfigs {
	s.TargetModelName = &v
	return s
}

func (s *HttpApiPolicyConfigsAiFallbackConfigServiceConfigs) Validate() error {
	return dara.Validate(s)
}

type HttpApiPolicyConfigsAiNetworkSearchConfig struct {
	DefaultEnable      *bool                                                        `json:"defaultEnable,omitempty" xml:"defaultEnable,omitempty"`
	DefaultLang        *string                                                      `json:"defaultLang,omitempty" xml:"defaultLang,omitempty"`
	NeedReference      *bool                                                        `json:"needReference,omitempty" xml:"needReference,omitempty"`
	PluginStatus       *HttpApiPolicyConfigsAiNetworkSearchConfigPluginStatus       `json:"pluginStatus,omitempty" xml:"pluginStatus,omitempty" type:"Struct"`
	ReferenceFormat    *string                                                      `json:"referenceFormat,omitempty" xml:"referenceFormat,omitempty"`
	ReferenceLocation  *string                                                      `json:"referenceLocation,omitempty" xml:"referenceLocation,omitempty"`
	SearchEngineConfig *HttpApiPolicyConfigsAiNetworkSearchConfigSearchEngineConfig `json:"searchEngineConfig,omitempty" xml:"searchEngineConfig,omitempty" type:"Struct"`
	SearchFrom         []*HttpApiPolicyConfigsAiNetworkSearchConfigSearchFrom       `json:"searchFrom,omitempty" xml:"searchFrom,omitempty" type:"Repeated"`
	SearchRewrite      *HttpApiPolicyConfigsAiNetworkSearchConfigSearchRewrite      `json:"searchRewrite,omitempty" xml:"searchRewrite,omitempty" type:"Struct"`
}

func (s HttpApiPolicyConfigsAiNetworkSearchConfig) String() string {
	return dara.Prettify(s)
}

func (s HttpApiPolicyConfigsAiNetworkSearchConfig) GoString() string {
	return s.String()
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfig) GetDefaultEnable() *bool {
	return s.DefaultEnable
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfig) GetDefaultLang() *string {
	return s.DefaultLang
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfig) GetNeedReference() *bool {
	return s.NeedReference
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfig) GetPluginStatus() *HttpApiPolicyConfigsAiNetworkSearchConfigPluginStatus {
	return s.PluginStatus
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfig) GetReferenceFormat() *string {
	return s.ReferenceFormat
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfig) GetReferenceLocation() *string {
	return s.ReferenceLocation
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfig) GetSearchEngineConfig() *HttpApiPolicyConfigsAiNetworkSearchConfigSearchEngineConfig {
	return s.SearchEngineConfig
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfig) GetSearchFrom() []*HttpApiPolicyConfigsAiNetworkSearchConfigSearchFrom {
	return s.SearchFrom
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfig) GetSearchRewrite() *HttpApiPolicyConfigsAiNetworkSearchConfigSearchRewrite {
	return s.SearchRewrite
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfig) SetDefaultEnable(v bool) *HttpApiPolicyConfigsAiNetworkSearchConfig {
	s.DefaultEnable = &v
	return s
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfig) SetDefaultLang(v string) *HttpApiPolicyConfigsAiNetworkSearchConfig {
	s.DefaultLang = &v
	return s
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfig) SetNeedReference(v bool) *HttpApiPolicyConfigsAiNetworkSearchConfig {
	s.NeedReference = &v
	return s
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfig) SetPluginStatus(v *HttpApiPolicyConfigsAiNetworkSearchConfigPluginStatus) *HttpApiPolicyConfigsAiNetworkSearchConfig {
	s.PluginStatus = v
	return s
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfig) SetReferenceFormat(v string) *HttpApiPolicyConfigsAiNetworkSearchConfig {
	s.ReferenceFormat = &v
	return s
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfig) SetReferenceLocation(v string) *HttpApiPolicyConfigsAiNetworkSearchConfig {
	s.ReferenceLocation = &v
	return s
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfig) SetSearchEngineConfig(v *HttpApiPolicyConfigsAiNetworkSearchConfigSearchEngineConfig) *HttpApiPolicyConfigsAiNetworkSearchConfig {
	s.SearchEngineConfig = v
	return s
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfig) SetSearchFrom(v []*HttpApiPolicyConfigsAiNetworkSearchConfigSearchFrom) *HttpApiPolicyConfigsAiNetworkSearchConfig {
	s.SearchFrom = v
	return s
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfig) SetSearchRewrite(v *HttpApiPolicyConfigsAiNetworkSearchConfigSearchRewrite) *HttpApiPolicyConfigsAiNetworkSearchConfig {
	s.SearchRewrite = v
	return s
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfig) Validate() error {
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

type HttpApiPolicyConfigsAiNetworkSearchConfigPluginStatus struct {
	ErrorLogs      map[string]*string `json:"errorLogs,omitempty" xml:"errorLogs,omitempty"`
	PluginId       *string            `json:"pluginId,omitempty" xml:"pluginId,omitempty"`
	ServiceHealthy *bool              `json:"serviceHealthy,omitempty" xml:"serviceHealthy,omitempty"`
}

func (s HttpApiPolicyConfigsAiNetworkSearchConfigPluginStatus) String() string {
	return dara.Prettify(s)
}

func (s HttpApiPolicyConfigsAiNetworkSearchConfigPluginStatus) GoString() string {
	return s.String()
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfigPluginStatus) GetErrorLogs() map[string]*string {
	return s.ErrorLogs
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfigPluginStatus) GetPluginId() *string {
	return s.PluginId
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfigPluginStatus) GetServiceHealthy() *bool {
	return s.ServiceHealthy
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfigPluginStatus) SetErrorLogs(v map[string]*string) *HttpApiPolicyConfigsAiNetworkSearchConfigPluginStatus {
	s.ErrorLogs = v
	return s
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfigPluginStatus) SetPluginId(v string) *HttpApiPolicyConfigsAiNetworkSearchConfigPluginStatus {
	s.PluginId = &v
	return s
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfigPluginStatus) SetServiceHealthy(v bool) *HttpApiPolicyConfigsAiNetworkSearchConfigPluginStatus {
	s.ServiceHealthy = &v
	return s
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfigPluginStatus) Validate() error {
	return dara.Validate(s)
}

type HttpApiPolicyConfigsAiNetworkSearchConfigSearchEngineConfig struct {
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

func (s HttpApiPolicyConfigsAiNetworkSearchConfigSearchEngineConfig) String() string {
	return dara.Prettify(s)
}

func (s HttpApiPolicyConfigsAiNetworkSearchConfigSearchEngineConfig) GoString() string {
	return s.String()
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfigSearchEngineConfig) GetApiKey() *string {
	return s.ApiKey
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfigSearchEngineConfig) GetContentMode() *string {
	return s.ContentMode
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfigSearchEngineConfig) GetCount() *int32 {
	return s.Count
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfigSearchEngineConfig) GetEndpoint() *string {
	return s.Endpoint
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfigSearchEngineConfig) GetIndustry() *string {
	return s.Industry
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfigSearchEngineConfig) GetOptionArgs() map[string]*string {
	return s.OptionArgs
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfigSearchEngineConfig) GetStart() *int32 {
	return s.Start
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfigSearchEngineConfig) GetTimeRange() *string {
	return s.TimeRange
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfigSearchEngineConfig) GetTimeoutMillisecond() *int32 {
	return s.TimeoutMillisecond
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfigSearchEngineConfig) GetType() *string {
	return s.Type
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfigSearchEngineConfig) SetApiKey(v string) *HttpApiPolicyConfigsAiNetworkSearchConfigSearchEngineConfig {
	s.ApiKey = &v
	return s
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfigSearchEngineConfig) SetContentMode(v string) *HttpApiPolicyConfigsAiNetworkSearchConfigSearchEngineConfig {
	s.ContentMode = &v
	return s
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfigSearchEngineConfig) SetCount(v int32) *HttpApiPolicyConfigsAiNetworkSearchConfigSearchEngineConfig {
	s.Count = &v
	return s
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfigSearchEngineConfig) SetEndpoint(v string) *HttpApiPolicyConfigsAiNetworkSearchConfigSearchEngineConfig {
	s.Endpoint = &v
	return s
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfigSearchEngineConfig) SetIndustry(v string) *HttpApiPolicyConfigsAiNetworkSearchConfigSearchEngineConfig {
	s.Industry = &v
	return s
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfigSearchEngineConfig) SetOptionArgs(v map[string]*string) *HttpApiPolicyConfigsAiNetworkSearchConfigSearchEngineConfig {
	s.OptionArgs = v
	return s
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfigSearchEngineConfig) SetStart(v int32) *HttpApiPolicyConfigsAiNetworkSearchConfigSearchEngineConfig {
	s.Start = &v
	return s
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfigSearchEngineConfig) SetTimeRange(v string) *HttpApiPolicyConfigsAiNetworkSearchConfigSearchEngineConfig {
	s.TimeRange = &v
	return s
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfigSearchEngineConfig) SetTimeoutMillisecond(v int32) *HttpApiPolicyConfigsAiNetworkSearchConfigSearchEngineConfig {
	s.TimeoutMillisecond = &v
	return s
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfigSearchEngineConfig) SetType(v string) *HttpApiPolicyConfigsAiNetworkSearchConfigSearchEngineConfig {
	s.Type = &v
	return s
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfigSearchEngineConfig) Validate() error {
	return dara.Validate(s)
}

type HttpApiPolicyConfigsAiNetworkSearchConfigSearchFrom struct {
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

func (s HttpApiPolicyConfigsAiNetworkSearchConfigSearchFrom) String() string {
	return dara.Prettify(s)
}

func (s HttpApiPolicyConfigsAiNetworkSearchConfigSearchFrom) GoString() string {
	return s.String()
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfigSearchFrom) GetApiKey() *string {
	return s.ApiKey
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfigSearchFrom) GetContentMode() *string {
	return s.ContentMode
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfigSearchFrom) GetCount() *int32 {
	return s.Count
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfigSearchFrom) GetEndpoint() *string {
	return s.Endpoint
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfigSearchFrom) GetIndustry() *string {
	return s.Industry
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfigSearchFrom) GetOptionArgs() map[string]*string {
	return s.OptionArgs
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfigSearchFrom) GetStart() *int32 {
	return s.Start
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfigSearchFrom) GetTimeRange() *string {
	return s.TimeRange
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfigSearchFrom) GetTimeoutMillisecond() *int32 {
	return s.TimeoutMillisecond
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfigSearchFrom) GetType() *string {
	return s.Type
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfigSearchFrom) SetApiKey(v string) *HttpApiPolicyConfigsAiNetworkSearchConfigSearchFrom {
	s.ApiKey = &v
	return s
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfigSearchFrom) SetContentMode(v string) *HttpApiPolicyConfigsAiNetworkSearchConfigSearchFrom {
	s.ContentMode = &v
	return s
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfigSearchFrom) SetCount(v int32) *HttpApiPolicyConfigsAiNetworkSearchConfigSearchFrom {
	s.Count = &v
	return s
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfigSearchFrom) SetEndpoint(v string) *HttpApiPolicyConfigsAiNetworkSearchConfigSearchFrom {
	s.Endpoint = &v
	return s
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfigSearchFrom) SetIndustry(v string) *HttpApiPolicyConfigsAiNetworkSearchConfigSearchFrom {
	s.Industry = &v
	return s
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfigSearchFrom) SetOptionArgs(v map[string]*string) *HttpApiPolicyConfigsAiNetworkSearchConfigSearchFrom {
	s.OptionArgs = v
	return s
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfigSearchFrom) SetStart(v int32) *HttpApiPolicyConfigsAiNetworkSearchConfigSearchFrom {
	s.Start = &v
	return s
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfigSearchFrom) SetTimeRange(v string) *HttpApiPolicyConfigsAiNetworkSearchConfigSearchFrom {
	s.TimeRange = &v
	return s
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfigSearchFrom) SetTimeoutMillisecond(v int32) *HttpApiPolicyConfigsAiNetworkSearchConfigSearchFrom {
	s.TimeoutMillisecond = &v
	return s
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfigSearchFrom) SetType(v string) *HttpApiPolicyConfigsAiNetworkSearchConfigSearchFrom {
	s.Type = &v
	return s
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfigSearchFrom) Validate() error {
	return dara.Validate(s)
}

type HttpApiPolicyConfigsAiNetworkSearchConfigSearchRewrite struct {
	Enable             *bool   `json:"enable,omitempty" xml:"enable,omitempty"`
	MaxCount           *int32  `json:"maxCount,omitempty" xml:"maxCount,omitempty"`
	ModelName          *string `json:"modelName,omitempty" xml:"modelName,omitempty"`
	ServiceId          *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	TimeoutMillisecond *int32  `json:"timeoutMillisecond,omitempty" xml:"timeoutMillisecond,omitempty"`
}

func (s HttpApiPolicyConfigsAiNetworkSearchConfigSearchRewrite) String() string {
	return dara.Prettify(s)
}

func (s HttpApiPolicyConfigsAiNetworkSearchConfigSearchRewrite) GoString() string {
	return s.String()
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfigSearchRewrite) GetEnable() *bool {
	return s.Enable
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfigSearchRewrite) GetMaxCount() *int32 {
	return s.MaxCount
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfigSearchRewrite) GetModelName() *string {
	return s.ModelName
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfigSearchRewrite) GetServiceId() *string {
	return s.ServiceId
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfigSearchRewrite) GetTimeoutMillisecond() *int32 {
	return s.TimeoutMillisecond
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfigSearchRewrite) SetEnable(v bool) *HttpApiPolicyConfigsAiNetworkSearchConfigSearchRewrite {
	s.Enable = &v
	return s
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfigSearchRewrite) SetMaxCount(v int32) *HttpApiPolicyConfigsAiNetworkSearchConfigSearchRewrite {
	s.MaxCount = &v
	return s
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfigSearchRewrite) SetModelName(v string) *HttpApiPolicyConfigsAiNetworkSearchConfigSearchRewrite {
	s.ModelName = &v
	return s
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfigSearchRewrite) SetServiceId(v string) *HttpApiPolicyConfigsAiNetworkSearchConfigSearchRewrite {
	s.ServiceId = &v
	return s
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfigSearchRewrite) SetTimeoutMillisecond(v int32) *HttpApiPolicyConfigsAiNetworkSearchConfigSearchRewrite {
	s.TimeoutMillisecond = &v
	return s
}

func (s *HttpApiPolicyConfigsAiNetworkSearchConfigSearchRewrite) Validate() error {
	return dara.Validate(s)
}

type HttpApiPolicyConfigsAiSecurityGuardConfig struct {
	BufferLimit                  *int32                                                                   `json:"bufferLimit,omitempty" xml:"bufferLimit,omitempty"`
	CheckRequest                 *bool                                                                    `json:"checkRequest,omitempty" xml:"checkRequest,omitempty"`
	CheckRequestImage            *bool                                                                    `json:"checkRequestImage,omitempty" xml:"checkRequestImage,omitempty"`
	CheckResponse                *bool                                                                    `json:"checkResponse,omitempty" xml:"checkResponse,omitempty"`
	CheckResponseImage           *bool                                                                    `json:"checkResponseImage,omitempty" xml:"checkResponseImage,omitempty"`
	ConsumerRequestCheckService  []*HttpApiPolicyConfigsAiSecurityGuardConfigConsumerRequestCheckService  `json:"consumerRequestCheckService,omitempty" xml:"consumerRequestCheckService,omitempty" type:"Repeated"`
	ConsumerResponseCheckService []*HttpApiPolicyConfigsAiSecurityGuardConfigConsumerResponseCheckService `json:"consumerResponseCheckService,omitempty" xml:"consumerResponseCheckService,omitempty" type:"Repeated"`
	ConsumerRiskLevel            []*HttpApiPolicyConfigsAiSecurityGuardConfigConsumerRiskLevel            `json:"consumerRiskLevel,omitempty" xml:"consumerRiskLevel,omitempty" type:"Repeated"`
	PluginStatus                 *HttpApiPolicyConfigsAiSecurityGuardConfigPluginStatus                   `json:"pluginStatus,omitempty" xml:"pluginStatus,omitempty" type:"Struct"`
	RequestCheckService          *string                                                                  `json:"requestCheckService,omitempty" xml:"requestCheckService,omitempty"`
	RequestImageCheckService     *string                                                                  `json:"requestImageCheckService,omitempty" xml:"requestImageCheckService,omitempty"`
	ResponseCheckService         *string                                                                  `json:"responseCheckService,omitempty" xml:"responseCheckService,omitempty"`
	ResponseImageCheckService    *string                                                                  `json:"responseImageCheckService,omitempty" xml:"responseImageCheckService,omitempty"`
	RiskAlertLevel               *string                                                                  `json:"riskAlertLevel,omitempty" xml:"riskAlertLevel,omitempty"`
	RiskConfig                   []*HttpApiPolicyConfigsAiSecurityGuardConfigRiskConfig                   `json:"riskConfig,omitempty" xml:"riskConfig,omitempty" type:"Repeated"`
	ServiceAddress               *string                                                                  `json:"serviceAddress,omitempty" xml:"serviceAddress,omitempty"`
}

func (s HttpApiPolicyConfigsAiSecurityGuardConfig) String() string {
	return dara.Prettify(s)
}

func (s HttpApiPolicyConfigsAiSecurityGuardConfig) GoString() string {
	return s.String()
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfig) GetBufferLimit() *int32 {
	return s.BufferLimit
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfig) GetCheckRequest() *bool {
	return s.CheckRequest
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfig) GetCheckRequestImage() *bool {
	return s.CheckRequestImage
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfig) GetCheckResponse() *bool {
	return s.CheckResponse
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfig) GetCheckResponseImage() *bool {
	return s.CheckResponseImage
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfig) GetConsumerRequestCheckService() []*HttpApiPolicyConfigsAiSecurityGuardConfigConsumerRequestCheckService {
	return s.ConsumerRequestCheckService
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfig) GetConsumerResponseCheckService() []*HttpApiPolicyConfigsAiSecurityGuardConfigConsumerResponseCheckService {
	return s.ConsumerResponseCheckService
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfig) GetConsumerRiskLevel() []*HttpApiPolicyConfigsAiSecurityGuardConfigConsumerRiskLevel {
	return s.ConsumerRiskLevel
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfig) GetPluginStatus() *HttpApiPolicyConfigsAiSecurityGuardConfigPluginStatus {
	return s.PluginStatus
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfig) GetRequestCheckService() *string {
	return s.RequestCheckService
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfig) GetRequestImageCheckService() *string {
	return s.RequestImageCheckService
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfig) GetResponseCheckService() *string {
	return s.ResponseCheckService
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfig) GetResponseImageCheckService() *string {
	return s.ResponseImageCheckService
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfig) GetRiskAlertLevel() *string {
	return s.RiskAlertLevel
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfig) GetRiskConfig() []*HttpApiPolicyConfigsAiSecurityGuardConfigRiskConfig {
	return s.RiskConfig
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfig) GetServiceAddress() *string {
	return s.ServiceAddress
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfig) SetBufferLimit(v int32) *HttpApiPolicyConfigsAiSecurityGuardConfig {
	s.BufferLimit = &v
	return s
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfig) SetCheckRequest(v bool) *HttpApiPolicyConfigsAiSecurityGuardConfig {
	s.CheckRequest = &v
	return s
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfig) SetCheckRequestImage(v bool) *HttpApiPolicyConfigsAiSecurityGuardConfig {
	s.CheckRequestImage = &v
	return s
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfig) SetCheckResponse(v bool) *HttpApiPolicyConfigsAiSecurityGuardConfig {
	s.CheckResponse = &v
	return s
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfig) SetCheckResponseImage(v bool) *HttpApiPolicyConfigsAiSecurityGuardConfig {
	s.CheckResponseImage = &v
	return s
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfig) SetConsumerRequestCheckService(v []*HttpApiPolicyConfigsAiSecurityGuardConfigConsumerRequestCheckService) *HttpApiPolicyConfigsAiSecurityGuardConfig {
	s.ConsumerRequestCheckService = v
	return s
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfig) SetConsumerResponseCheckService(v []*HttpApiPolicyConfigsAiSecurityGuardConfigConsumerResponseCheckService) *HttpApiPolicyConfigsAiSecurityGuardConfig {
	s.ConsumerResponseCheckService = v
	return s
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfig) SetConsumerRiskLevel(v []*HttpApiPolicyConfigsAiSecurityGuardConfigConsumerRiskLevel) *HttpApiPolicyConfigsAiSecurityGuardConfig {
	s.ConsumerRiskLevel = v
	return s
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfig) SetPluginStatus(v *HttpApiPolicyConfigsAiSecurityGuardConfigPluginStatus) *HttpApiPolicyConfigsAiSecurityGuardConfig {
	s.PluginStatus = v
	return s
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfig) SetRequestCheckService(v string) *HttpApiPolicyConfigsAiSecurityGuardConfig {
	s.RequestCheckService = &v
	return s
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfig) SetRequestImageCheckService(v string) *HttpApiPolicyConfigsAiSecurityGuardConfig {
	s.RequestImageCheckService = &v
	return s
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfig) SetResponseCheckService(v string) *HttpApiPolicyConfigsAiSecurityGuardConfig {
	s.ResponseCheckService = &v
	return s
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfig) SetResponseImageCheckService(v string) *HttpApiPolicyConfigsAiSecurityGuardConfig {
	s.ResponseImageCheckService = &v
	return s
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfig) SetRiskAlertLevel(v string) *HttpApiPolicyConfigsAiSecurityGuardConfig {
	s.RiskAlertLevel = &v
	return s
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfig) SetRiskConfig(v []*HttpApiPolicyConfigsAiSecurityGuardConfigRiskConfig) *HttpApiPolicyConfigsAiSecurityGuardConfig {
	s.RiskConfig = v
	return s
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfig) SetServiceAddress(v string) *HttpApiPolicyConfigsAiSecurityGuardConfig {
	s.ServiceAddress = &v
	return s
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfig) Validate() error {
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

type HttpApiPolicyConfigsAiSecurityGuardConfigConsumerRequestCheckService struct {
	MatchType                *string `json:"matchType,omitempty" xml:"matchType,omitempty"`
	ModalityType             *string `json:"modalityType,omitempty" xml:"modalityType,omitempty"`
	Name                     *string `json:"name,omitempty" xml:"name,omitempty"`
	RequestCheckService      *string `json:"requestCheckService,omitempty" xml:"requestCheckService,omitempty"`
	RequestImageCheckService *string `json:"requestImageCheckService,omitempty" xml:"requestImageCheckService,omitempty"`
}

func (s HttpApiPolicyConfigsAiSecurityGuardConfigConsumerRequestCheckService) String() string {
	return dara.Prettify(s)
}

func (s HttpApiPolicyConfigsAiSecurityGuardConfigConsumerRequestCheckService) GoString() string {
	return s.String()
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfigConsumerRequestCheckService) GetMatchType() *string {
	return s.MatchType
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfigConsumerRequestCheckService) GetModalityType() *string {
	return s.ModalityType
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfigConsumerRequestCheckService) GetName() *string {
	return s.Name
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfigConsumerRequestCheckService) GetRequestCheckService() *string {
	return s.RequestCheckService
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfigConsumerRequestCheckService) GetRequestImageCheckService() *string {
	return s.RequestImageCheckService
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfigConsumerRequestCheckService) SetMatchType(v string) *HttpApiPolicyConfigsAiSecurityGuardConfigConsumerRequestCheckService {
	s.MatchType = &v
	return s
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfigConsumerRequestCheckService) SetModalityType(v string) *HttpApiPolicyConfigsAiSecurityGuardConfigConsumerRequestCheckService {
	s.ModalityType = &v
	return s
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfigConsumerRequestCheckService) SetName(v string) *HttpApiPolicyConfigsAiSecurityGuardConfigConsumerRequestCheckService {
	s.Name = &v
	return s
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfigConsumerRequestCheckService) SetRequestCheckService(v string) *HttpApiPolicyConfigsAiSecurityGuardConfigConsumerRequestCheckService {
	s.RequestCheckService = &v
	return s
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfigConsumerRequestCheckService) SetRequestImageCheckService(v string) *HttpApiPolicyConfigsAiSecurityGuardConfigConsumerRequestCheckService {
	s.RequestImageCheckService = &v
	return s
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfigConsumerRequestCheckService) Validate() error {
	return dara.Validate(s)
}

type HttpApiPolicyConfigsAiSecurityGuardConfigConsumerResponseCheckService struct {
	MatchType                 *string `json:"matchType,omitempty" xml:"matchType,omitempty"`
	ModalityType              *string `json:"modalityType,omitempty" xml:"modalityType,omitempty"`
	Name                      *string `json:"name,omitempty" xml:"name,omitempty"`
	ResponseCheckService      *string `json:"responseCheckService,omitempty" xml:"responseCheckService,omitempty"`
	ResponseImageCheckService *string `json:"responseImageCheckService,omitempty" xml:"responseImageCheckService,omitempty"`
}

func (s HttpApiPolicyConfigsAiSecurityGuardConfigConsumerResponseCheckService) String() string {
	return dara.Prettify(s)
}

func (s HttpApiPolicyConfigsAiSecurityGuardConfigConsumerResponseCheckService) GoString() string {
	return s.String()
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfigConsumerResponseCheckService) GetMatchType() *string {
	return s.MatchType
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfigConsumerResponseCheckService) GetModalityType() *string {
	return s.ModalityType
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfigConsumerResponseCheckService) GetName() *string {
	return s.Name
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfigConsumerResponseCheckService) GetResponseCheckService() *string {
	return s.ResponseCheckService
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfigConsumerResponseCheckService) GetResponseImageCheckService() *string {
	return s.ResponseImageCheckService
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfigConsumerResponseCheckService) SetMatchType(v string) *HttpApiPolicyConfigsAiSecurityGuardConfigConsumerResponseCheckService {
	s.MatchType = &v
	return s
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfigConsumerResponseCheckService) SetModalityType(v string) *HttpApiPolicyConfigsAiSecurityGuardConfigConsumerResponseCheckService {
	s.ModalityType = &v
	return s
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfigConsumerResponseCheckService) SetName(v string) *HttpApiPolicyConfigsAiSecurityGuardConfigConsumerResponseCheckService {
	s.Name = &v
	return s
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfigConsumerResponseCheckService) SetResponseCheckService(v string) *HttpApiPolicyConfigsAiSecurityGuardConfigConsumerResponseCheckService {
	s.ResponseCheckService = &v
	return s
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfigConsumerResponseCheckService) SetResponseImageCheckService(v string) *HttpApiPolicyConfigsAiSecurityGuardConfigConsumerResponseCheckService {
	s.ResponseImageCheckService = &v
	return s
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfigConsumerResponseCheckService) Validate() error {
	return dara.Validate(s)
}

type HttpApiPolicyConfigsAiSecurityGuardConfigConsumerRiskLevel struct {
	Level     *string `json:"level,omitempty" xml:"level,omitempty"`
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty"`
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
	Type      *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s HttpApiPolicyConfigsAiSecurityGuardConfigConsumerRiskLevel) String() string {
	return dara.Prettify(s)
}

func (s HttpApiPolicyConfigsAiSecurityGuardConfigConsumerRiskLevel) GoString() string {
	return s.String()
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfigConsumerRiskLevel) GetLevel() *string {
	return s.Level
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfigConsumerRiskLevel) GetMatchType() *string {
	return s.MatchType
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfigConsumerRiskLevel) GetName() *string {
	return s.Name
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfigConsumerRiskLevel) GetType() *string {
	return s.Type
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfigConsumerRiskLevel) SetLevel(v string) *HttpApiPolicyConfigsAiSecurityGuardConfigConsumerRiskLevel {
	s.Level = &v
	return s
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfigConsumerRiskLevel) SetMatchType(v string) *HttpApiPolicyConfigsAiSecurityGuardConfigConsumerRiskLevel {
	s.MatchType = &v
	return s
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfigConsumerRiskLevel) SetName(v string) *HttpApiPolicyConfigsAiSecurityGuardConfigConsumerRiskLevel {
	s.Name = &v
	return s
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfigConsumerRiskLevel) SetType(v string) *HttpApiPolicyConfigsAiSecurityGuardConfigConsumerRiskLevel {
	s.Type = &v
	return s
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfigConsumerRiskLevel) Validate() error {
	return dara.Validate(s)
}

type HttpApiPolicyConfigsAiSecurityGuardConfigPluginStatus struct {
	ErrorLogs      map[string]*string `json:"errorLogs,omitempty" xml:"errorLogs,omitempty"`
	PluginId       *string            `json:"pluginId,omitempty" xml:"pluginId,omitempty"`
	ServiceHealthy *bool              `json:"serviceHealthy,omitempty" xml:"serviceHealthy,omitempty"`
}

func (s HttpApiPolicyConfigsAiSecurityGuardConfigPluginStatus) String() string {
	return dara.Prettify(s)
}

func (s HttpApiPolicyConfigsAiSecurityGuardConfigPluginStatus) GoString() string {
	return s.String()
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfigPluginStatus) GetErrorLogs() map[string]*string {
	return s.ErrorLogs
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfigPluginStatus) GetPluginId() *string {
	return s.PluginId
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfigPluginStatus) GetServiceHealthy() *bool {
	return s.ServiceHealthy
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfigPluginStatus) SetErrorLogs(v map[string]*string) *HttpApiPolicyConfigsAiSecurityGuardConfigPluginStatus {
	s.ErrorLogs = v
	return s
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfigPluginStatus) SetPluginId(v string) *HttpApiPolicyConfigsAiSecurityGuardConfigPluginStatus {
	s.PluginId = &v
	return s
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfigPluginStatus) SetServiceHealthy(v bool) *HttpApiPolicyConfigsAiSecurityGuardConfigPluginStatus {
	s.ServiceHealthy = &v
	return s
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfigPluginStatus) Validate() error {
	return dara.Validate(s)
}

type HttpApiPolicyConfigsAiSecurityGuardConfigRiskConfig struct {
	ConsumerRules *HttpApiPolicyConfigsAiSecurityGuardConfigRiskConfigConsumerRules `json:"consumerRules,omitempty" xml:"consumerRules,omitempty" type:"Struct"`
	Level         *string                                                           `json:"level,omitempty" xml:"level,omitempty"`
	Type          *string                                                           `json:"type,omitempty" xml:"type,omitempty"`
}

func (s HttpApiPolicyConfigsAiSecurityGuardConfigRiskConfig) String() string {
	return dara.Prettify(s)
}

func (s HttpApiPolicyConfigsAiSecurityGuardConfigRiskConfig) GoString() string {
	return s.String()
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfigRiskConfig) GetConsumerRules() *HttpApiPolicyConfigsAiSecurityGuardConfigRiskConfigConsumerRules {
	return s.ConsumerRules
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfigRiskConfig) GetLevel() *string {
	return s.Level
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfigRiskConfig) GetType() *string {
	return s.Type
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfigRiskConfig) SetConsumerRules(v *HttpApiPolicyConfigsAiSecurityGuardConfigRiskConfigConsumerRules) *HttpApiPolicyConfigsAiSecurityGuardConfigRiskConfig {
	s.ConsumerRules = v
	return s
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfigRiskConfig) SetLevel(v string) *HttpApiPolicyConfigsAiSecurityGuardConfigRiskConfig {
	s.Level = &v
	return s
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfigRiskConfig) SetType(v string) *HttpApiPolicyConfigsAiSecurityGuardConfigRiskConfig {
	s.Type = &v
	return s
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfigRiskConfig) Validate() error {
	if s.ConsumerRules != nil {
		if err := s.ConsumerRules.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type HttpApiPolicyConfigsAiSecurityGuardConfigRiskConfigConsumerRules struct {
	MatchType *string `json:"matchType,omitempty" xml:"matchType,omitempty"`
	Pattern   *string `json:"pattern,omitempty" xml:"pattern,omitempty"`
}

func (s HttpApiPolicyConfigsAiSecurityGuardConfigRiskConfigConsumerRules) String() string {
	return dara.Prettify(s)
}

func (s HttpApiPolicyConfigsAiSecurityGuardConfigRiskConfigConsumerRules) GoString() string {
	return s.String()
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfigRiskConfigConsumerRules) GetMatchType() *string {
	return s.MatchType
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfigRiskConfigConsumerRules) GetPattern() *string {
	return s.Pattern
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfigRiskConfigConsumerRules) SetMatchType(v string) *HttpApiPolicyConfigsAiSecurityGuardConfigRiskConfigConsumerRules {
	s.MatchType = &v
	return s
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfigRiskConfigConsumerRules) SetPattern(v string) *HttpApiPolicyConfigsAiSecurityGuardConfigRiskConfigConsumerRules {
	s.Pattern = &v
	return s
}

func (s *HttpApiPolicyConfigsAiSecurityGuardConfigRiskConfigConsumerRules) Validate() error {
	return dara.Validate(s)
}

type HttpApiPolicyConfigsAiStatisticsConfig struct {
	LogRequestContent  *bool `json:"logRequestContent,omitempty" xml:"logRequestContent,omitempty"`
	LogResponseContent *bool `json:"logResponseContent,omitempty" xml:"logResponseContent,omitempty"`
}

func (s HttpApiPolicyConfigsAiStatisticsConfig) String() string {
	return dara.Prettify(s)
}

func (s HttpApiPolicyConfigsAiStatisticsConfig) GoString() string {
	return s.String()
}

func (s *HttpApiPolicyConfigsAiStatisticsConfig) GetLogRequestContent() *bool {
	return s.LogRequestContent
}

func (s *HttpApiPolicyConfigsAiStatisticsConfig) GetLogResponseContent() *bool {
	return s.LogResponseContent
}

func (s *HttpApiPolicyConfigsAiStatisticsConfig) SetLogRequestContent(v bool) *HttpApiPolicyConfigsAiStatisticsConfig {
	s.LogRequestContent = &v
	return s
}

func (s *HttpApiPolicyConfigsAiStatisticsConfig) SetLogResponseContent(v bool) *HttpApiPolicyConfigsAiStatisticsConfig {
	s.LogResponseContent = &v
	return s
}

func (s *HttpApiPolicyConfigsAiStatisticsConfig) Validate() error {
	return dara.Validate(s)
}

type HttpApiPolicyConfigsAiTokenRateLimitConfig struct {
	EnableGlobalRules *bool                                                    `json:"enableGlobalRules,omitempty" xml:"enableGlobalRules,omitempty"`
	GlobalRules       []*HttpApiPolicyConfigsAiTokenRateLimitConfigGlobalRules `json:"globalRules,omitempty" xml:"globalRules,omitempty" type:"Repeated"`
	PluginStatus      *HttpApiPolicyConfigsAiTokenRateLimitConfigPluginStatus  `json:"pluginStatus,omitempty" xml:"pluginStatus,omitempty" type:"Struct"`
	RedisConfig       *HttpApiPolicyConfigsAiTokenRateLimitConfigRedisConfig   `json:"redisConfig,omitempty" xml:"redisConfig,omitempty" type:"Struct"`
	Rules             []*HttpApiPolicyConfigsAiTokenRateLimitConfigRules       `json:"rules,omitempty" xml:"rules,omitempty" type:"Repeated"`
}

func (s HttpApiPolicyConfigsAiTokenRateLimitConfig) String() string {
	return dara.Prettify(s)
}

func (s HttpApiPolicyConfigsAiTokenRateLimitConfig) GoString() string {
	return s.String()
}

func (s *HttpApiPolicyConfigsAiTokenRateLimitConfig) GetEnableGlobalRules() *bool {
	return s.EnableGlobalRules
}

func (s *HttpApiPolicyConfigsAiTokenRateLimitConfig) GetGlobalRules() []*HttpApiPolicyConfigsAiTokenRateLimitConfigGlobalRules {
	return s.GlobalRules
}

func (s *HttpApiPolicyConfigsAiTokenRateLimitConfig) GetPluginStatus() *HttpApiPolicyConfigsAiTokenRateLimitConfigPluginStatus {
	return s.PluginStatus
}

func (s *HttpApiPolicyConfigsAiTokenRateLimitConfig) GetRedisConfig() *HttpApiPolicyConfigsAiTokenRateLimitConfigRedisConfig {
	return s.RedisConfig
}

func (s *HttpApiPolicyConfigsAiTokenRateLimitConfig) GetRules() []*HttpApiPolicyConfigsAiTokenRateLimitConfigRules {
	return s.Rules
}

func (s *HttpApiPolicyConfigsAiTokenRateLimitConfig) SetEnableGlobalRules(v bool) *HttpApiPolicyConfigsAiTokenRateLimitConfig {
	s.EnableGlobalRules = &v
	return s
}

func (s *HttpApiPolicyConfigsAiTokenRateLimitConfig) SetGlobalRules(v []*HttpApiPolicyConfigsAiTokenRateLimitConfigGlobalRules) *HttpApiPolicyConfigsAiTokenRateLimitConfig {
	s.GlobalRules = v
	return s
}

func (s *HttpApiPolicyConfigsAiTokenRateLimitConfig) SetPluginStatus(v *HttpApiPolicyConfigsAiTokenRateLimitConfigPluginStatus) *HttpApiPolicyConfigsAiTokenRateLimitConfig {
	s.PluginStatus = v
	return s
}

func (s *HttpApiPolicyConfigsAiTokenRateLimitConfig) SetRedisConfig(v *HttpApiPolicyConfigsAiTokenRateLimitConfigRedisConfig) *HttpApiPolicyConfigsAiTokenRateLimitConfig {
	s.RedisConfig = v
	return s
}

func (s *HttpApiPolicyConfigsAiTokenRateLimitConfig) SetRules(v []*HttpApiPolicyConfigsAiTokenRateLimitConfigRules) *HttpApiPolicyConfigsAiTokenRateLimitConfig {
	s.Rules = v
	return s
}

func (s *HttpApiPolicyConfigsAiTokenRateLimitConfig) Validate() error {
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

type HttpApiPolicyConfigsAiTokenRateLimitConfigGlobalRules struct {
	LimitMode  *string `json:"limitMode,omitempty" xml:"limitMode,omitempty"`
	LimitType  *string `json:"limitType,omitempty" xml:"limitType,omitempty"`
	LimitValue *string `json:"limitValue,omitempty" xml:"limitValue,omitempty"`
	MatchKey   *string `json:"matchKey,omitempty" xml:"matchKey,omitempty"`
	MatchType  *string `json:"matchType,omitempty" xml:"matchType,omitempty"`
	MatchValue *string `json:"matchValue,omitempty" xml:"matchValue,omitempty"`
}

func (s HttpApiPolicyConfigsAiTokenRateLimitConfigGlobalRules) String() string {
	return dara.Prettify(s)
}

func (s HttpApiPolicyConfigsAiTokenRateLimitConfigGlobalRules) GoString() string {
	return s.String()
}

func (s *HttpApiPolicyConfigsAiTokenRateLimitConfigGlobalRules) GetLimitMode() *string {
	return s.LimitMode
}

func (s *HttpApiPolicyConfigsAiTokenRateLimitConfigGlobalRules) GetLimitType() *string {
	return s.LimitType
}

func (s *HttpApiPolicyConfigsAiTokenRateLimitConfigGlobalRules) GetLimitValue() *string {
	return s.LimitValue
}

func (s *HttpApiPolicyConfigsAiTokenRateLimitConfigGlobalRules) GetMatchKey() *string {
	return s.MatchKey
}

func (s *HttpApiPolicyConfigsAiTokenRateLimitConfigGlobalRules) GetMatchType() *string {
	return s.MatchType
}

func (s *HttpApiPolicyConfigsAiTokenRateLimitConfigGlobalRules) GetMatchValue() *string {
	return s.MatchValue
}

func (s *HttpApiPolicyConfigsAiTokenRateLimitConfigGlobalRules) SetLimitMode(v string) *HttpApiPolicyConfigsAiTokenRateLimitConfigGlobalRules {
	s.LimitMode = &v
	return s
}

func (s *HttpApiPolicyConfigsAiTokenRateLimitConfigGlobalRules) SetLimitType(v string) *HttpApiPolicyConfigsAiTokenRateLimitConfigGlobalRules {
	s.LimitType = &v
	return s
}

func (s *HttpApiPolicyConfigsAiTokenRateLimitConfigGlobalRules) SetLimitValue(v string) *HttpApiPolicyConfigsAiTokenRateLimitConfigGlobalRules {
	s.LimitValue = &v
	return s
}

func (s *HttpApiPolicyConfigsAiTokenRateLimitConfigGlobalRules) SetMatchKey(v string) *HttpApiPolicyConfigsAiTokenRateLimitConfigGlobalRules {
	s.MatchKey = &v
	return s
}

func (s *HttpApiPolicyConfigsAiTokenRateLimitConfigGlobalRules) SetMatchType(v string) *HttpApiPolicyConfigsAiTokenRateLimitConfigGlobalRules {
	s.MatchType = &v
	return s
}

func (s *HttpApiPolicyConfigsAiTokenRateLimitConfigGlobalRules) SetMatchValue(v string) *HttpApiPolicyConfigsAiTokenRateLimitConfigGlobalRules {
	s.MatchValue = &v
	return s
}

func (s *HttpApiPolicyConfigsAiTokenRateLimitConfigGlobalRules) Validate() error {
	return dara.Validate(s)
}

type HttpApiPolicyConfigsAiTokenRateLimitConfigPluginStatus struct {
	ErrorLogs      map[string]*string `json:"errorLogs,omitempty" xml:"errorLogs,omitempty"`
	PluginId       *string            `json:"pluginId,omitempty" xml:"pluginId,omitempty"`
	ServiceHealthy *bool              `json:"serviceHealthy,omitempty" xml:"serviceHealthy,omitempty"`
}

func (s HttpApiPolicyConfigsAiTokenRateLimitConfigPluginStatus) String() string {
	return dara.Prettify(s)
}

func (s HttpApiPolicyConfigsAiTokenRateLimitConfigPluginStatus) GoString() string {
	return s.String()
}

func (s *HttpApiPolicyConfigsAiTokenRateLimitConfigPluginStatus) GetErrorLogs() map[string]*string {
	return s.ErrorLogs
}

func (s *HttpApiPolicyConfigsAiTokenRateLimitConfigPluginStatus) GetPluginId() *string {
	return s.PluginId
}

func (s *HttpApiPolicyConfigsAiTokenRateLimitConfigPluginStatus) GetServiceHealthy() *bool {
	return s.ServiceHealthy
}

func (s *HttpApiPolicyConfigsAiTokenRateLimitConfigPluginStatus) SetErrorLogs(v map[string]*string) *HttpApiPolicyConfigsAiTokenRateLimitConfigPluginStatus {
	s.ErrorLogs = v
	return s
}

func (s *HttpApiPolicyConfigsAiTokenRateLimitConfigPluginStatus) SetPluginId(v string) *HttpApiPolicyConfigsAiTokenRateLimitConfigPluginStatus {
	s.PluginId = &v
	return s
}

func (s *HttpApiPolicyConfigsAiTokenRateLimitConfigPluginStatus) SetServiceHealthy(v bool) *HttpApiPolicyConfigsAiTokenRateLimitConfigPluginStatus {
	s.ServiceHealthy = &v
	return s
}

func (s *HttpApiPolicyConfigsAiTokenRateLimitConfigPluginStatus) Validate() error {
	return dara.Validate(s)
}

type HttpApiPolicyConfigsAiTokenRateLimitConfigRedisConfig struct {
	DatabaseNumber *int32  `json:"databaseNumber,omitempty" xml:"databaseNumber,omitempty"`
	Host           *string `json:"host,omitempty" xml:"host,omitempty"`
	Password       *string `json:"password,omitempty" xml:"password,omitempty"`
	Port           *int32  `json:"port,omitempty" xml:"port,omitempty"`
	Timeout        *int32  `json:"timeout,omitempty" xml:"timeout,omitempty"`
	Username       *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s HttpApiPolicyConfigsAiTokenRateLimitConfigRedisConfig) String() string {
	return dara.Prettify(s)
}

func (s HttpApiPolicyConfigsAiTokenRateLimitConfigRedisConfig) GoString() string {
	return s.String()
}

func (s *HttpApiPolicyConfigsAiTokenRateLimitConfigRedisConfig) GetDatabaseNumber() *int32 {
	return s.DatabaseNumber
}

func (s *HttpApiPolicyConfigsAiTokenRateLimitConfigRedisConfig) GetHost() *string {
	return s.Host
}

func (s *HttpApiPolicyConfigsAiTokenRateLimitConfigRedisConfig) GetPassword() *string {
	return s.Password
}

func (s *HttpApiPolicyConfigsAiTokenRateLimitConfigRedisConfig) GetPort() *int32 {
	return s.Port
}

func (s *HttpApiPolicyConfigsAiTokenRateLimitConfigRedisConfig) GetTimeout() *int32 {
	return s.Timeout
}

func (s *HttpApiPolicyConfigsAiTokenRateLimitConfigRedisConfig) GetUsername() *string {
	return s.Username
}

func (s *HttpApiPolicyConfigsAiTokenRateLimitConfigRedisConfig) SetDatabaseNumber(v int32) *HttpApiPolicyConfigsAiTokenRateLimitConfigRedisConfig {
	s.DatabaseNumber = &v
	return s
}

func (s *HttpApiPolicyConfigsAiTokenRateLimitConfigRedisConfig) SetHost(v string) *HttpApiPolicyConfigsAiTokenRateLimitConfigRedisConfig {
	s.Host = &v
	return s
}

func (s *HttpApiPolicyConfigsAiTokenRateLimitConfigRedisConfig) SetPassword(v string) *HttpApiPolicyConfigsAiTokenRateLimitConfigRedisConfig {
	s.Password = &v
	return s
}

func (s *HttpApiPolicyConfigsAiTokenRateLimitConfigRedisConfig) SetPort(v int32) *HttpApiPolicyConfigsAiTokenRateLimitConfigRedisConfig {
	s.Port = &v
	return s
}

func (s *HttpApiPolicyConfigsAiTokenRateLimitConfigRedisConfig) SetTimeout(v int32) *HttpApiPolicyConfigsAiTokenRateLimitConfigRedisConfig {
	s.Timeout = &v
	return s
}

func (s *HttpApiPolicyConfigsAiTokenRateLimitConfigRedisConfig) SetUsername(v string) *HttpApiPolicyConfigsAiTokenRateLimitConfigRedisConfig {
	s.Username = &v
	return s
}

func (s *HttpApiPolicyConfigsAiTokenRateLimitConfigRedisConfig) Validate() error {
	return dara.Validate(s)
}

type HttpApiPolicyConfigsAiTokenRateLimitConfigRules struct {
	LimitMode  *string `json:"limitMode,omitempty" xml:"limitMode,omitempty"`
	LimitType  *string `json:"limitType,omitempty" xml:"limitType,omitempty"`
	LimitValue *string `json:"limitValue,omitempty" xml:"limitValue,omitempty"`
	MatchKey   *string `json:"matchKey,omitempty" xml:"matchKey,omitempty"`
	MatchType  *string `json:"matchType,omitempty" xml:"matchType,omitempty"`
	MatchValue *string `json:"matchValue,omitempty" xml:"matchValue,omitempty"`
}

func (s HttpApiPolicyConfigsAiTokenRateLimitConfigRules) String() string {
	return dara.Prettify(s)
}

func (s HttpApiPolicyConfigsAiTokenRateLimitConfigRules) GoString() string {
	return s.String()
}

func (s *HttpApiPolicyConfigsAiTokenRateLimitConfigRules) GetLimitMode() *string {
	return s.LimitMode
}

func (s *HttpApiPolicyConfigsAiTokenRateLimitConfigRules) GetLimitType() *string {
	return s.LimitType
}

func (s *HttpApiPolicyConfigsAiTokenRateLimitConfigRules) GetLimitValue() *string {
	return s.LimitValue
}

func (s *HttpApiPolicyConfigsAiTokenRateLimitConfigRules) GetMatchKey() *string {
	return s.MatchKey
}

func (s *HttpApiPolicyConfigsAiTokenRateLimitConfigRules) GetMatchType() *string {
	return s.MatchType
}

func (s *HttpApiPolicyConfigsAiTokenRateLimitConfigRules) GetMatchValue() *string {
	return s.MatchValue
}

func (s *HttpApiPolicyConfigsAiTokenRateLimitConfigRules) SetLimitMode(v string) *HttpApiPolicyConfigsAiTokenRateLimitConfigRules {
	s.LimitMode = &v
	return s
}

func (s *HttpApiPolicyConfigsAiTokenRateLimitConfigRules) SetLimitType(v string) *HttpApiPolicyConfigsAiTokenRateLimitConfigRules {
	s.LimitType = &v
	return s
}

func (s *HttpApiPolicyConfigsAiTokenRateLimitConfigRules) SetLimitValue(v string) *HttpApiPolicyConfigsAiTokenRateLimitConfigRules {
	s.LimitValue = &v
	return s
}

func (s *HttpApiPolicyConfigsAiTokenRateLimitConfigRules) SetMatchKey(v string) *HttpApiPolicyConfigsAiTokenRateLimitConfigRules {
	s.MatchKey = &v
	return s
}

func (s *HttpApiPolicyConfigsAiTokenRateLimitConfigRules) SetMatchType(v string) *HttpApiPolicyConfigsAiTokenRateLimitConfigRules {
	s.MatchType = &v
	return s
}

func (s *HttpApiPolicyConfigsAiTokenRateLimitConfigRules) SetMatchValue(v string) *HttpApiPolicyConfigsAiTokenRateLimitConfigRules {
	s.MatchValue = &v
	return s
}

func (s *HttpApiPolicyConfigsAiTokenRateLimitConfigRules) Validate() error {
	return dara.Validate(s)
}

type HttpApiPolicyConfigsAiToolSelectionConfig struct {
	EnableConditions *HttpApiPolicyConfigsAiToolSelectionConfigEnableConditions `json:"enableConditions,omitempty" xml:"enableConditions,omitempty" type:"Struct"`
	PluginStatus     *HttpApiPolicyConfigsAiToolSelectionConfigPluginStatus     `json:"pluginStatus,omitempty" xml:"pluginStatus,omitempty" type:"Struct"`
	QueryRewriting   *HttpApiPolicyConfigsAiToolSelectionConfigQueryRewriting   `json:"queryRewriting,omitempty" xml:"queryRewriting,omitempty" type:"Struct"`
	ToolReranking    *HttpApiPolicyConfigsAiToolSelectionConfigToolReranking    `json:"toolReranking,omitempty" xml:"toolReranking,omitempty" type:"Struct"`
}

func (s HttpApiPolicyConfigsAiToolSelectionConfig) String() string {
	return dara.Prettify(s)
}

func (s HttpApiPolicyConfigsAiToolSelectionConfig) GoString() string {
	return s.String()
}

func (s *HttpApiPolicyConfigsAiToolSelectionConfig) GetEnableConditions() *HttpApiPolicyConfigsAiToolSelectionConfigEnableConditions {
	return s.EnableConditions
}

func (s *HttpApiPolicyConfigsAiToolSelectionConfig) GetPluginStatus() *HttpApiPolicyConfigsAiToolSelectionConfigPluginStatus {
	return s.PluginStatus
}

func (s *HttpApiPolicyConfigsAiToolSelectionConfig) GetQueryRewriting() *HttpApiPolicyConfigsAiToolSelectionConfigQueryRewriting {
	return s.QueryRewriting
}

func (s *HttpApiPolicyConfigsAiToolSelectionConfig) GetToolReranking() *HttpApiPolicyConfigsAiToolSelectionConfigToolReranking {
	return s.ToolReranking
}

func (s *HttpApiPolicyConfigsAiToolSelectionConfig) SetEnableConditions(v *HttpApiPolicyConfigsAiToolSelectionConfigEnableConditions) *HttpApiPolicyConfigsAiToolSelectionConfig {
	s.EnableConditions = v
	return s
}

func (s *HttpApiPolicyConfigsAiToolSelectionConfig) SetPluginStatus(v *HttpApiPolicyConfigsAiToolSelectionConfigPluginStatus) *HttpApiPolicyConfigsAiToolSelectionConfig {
	s.PluginStatus = v
	return s
}

func (s *HttpApiPolicyConfigsAiToolSelectionConfig) SetQueryRewriting(v *HttpApiPolicyConfigsAiToolSelectionConfigQueryRewriting) *HttpApiPolicyConfigsAiToolSelectionConfig {
	s.QueryRewriting = v
	return s
}

func (s *HttpApiPolicyConfigsAiToolSelectionConfig) SetToolReranking(v *HttpApiPolicyConfigsAiToolSelectionConfigToolReranking) *HttpApiPolicyConfigsAiToolSelectionConfig {
	s.ToolReranking = v
	return s
}

func (s *HttpApiPolicyConfigsAiToolSelectionConfig) Validate() error {
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

type HttpApiPolicyConfigsAiToolSelectionConfigEnableConditions struct {
	ToolCountThreshold *int32 `json:"toolCountThreshold,omitempty" xml:"toolCountThreshold,omitempty"`
}

func (s HttpApiPolicyConfigsAiToolSelectionConfigEnableConditions) String() string {
	return dara.Prettify(s)
}

func (s HttpApiPolicyConfigsAiToolSelectionConfigEnableConditions) GoString() string {
	return s.String()
}

func (s *HttpApiPolicyConfigsAiToolSelectionConfigEnableConditions) GetToolCountThreshold() *int32 {
	return s.ToolCountThreshold
}

func (s *HttpApiPolicyConfigsAiToolSelectionConfigEnableConditions) SetToolCountThreshold(v int32) *HttpApiPolicyConfigsAiToolSelectionConfigEnableConditions {
	s.ToolCountThreshold = &v
	return s
}

func (s *HttpApiPolicyConfigsAiToolSelectionConfigEnableConditions) Validate() error {
	return dara.Validate(s)
}

type HttpApiPolicyConfigsAiToolSelectionConfigPluginStatus struct {
	ErrorLogs      map[string]*string `json:"errorLogs,omitempty" xml:"errorLogs,omitempty"`
	PluginId       *string            `json:"pluginId,omitempty" xml:"pluginId,omitempty"`
	ServiceHealthy *bool              `json:"serviceHealthy,omitempty" xml:"serviceHealthy,omitempty"`
}

func (s HttpApiPolicyConfigsAiToolSelectionConfigPluginStatus) String() string {
	return dara.Prettify(s)
}

func (s HttpApiPolicyConfigsAiToolSelectionConfigPluginStatus) GoString() string {
	return s.String()
}

func (s *HttpApiPolicyConfigsAiToolSelectionConfigPluginStatus) GetErrorLogs() map[string]*string {
	return s.ErrorLogs
}

func (s *HttpApiPolicyConfigsAiToolSelectionConfigPluginStatus) GetPluginId() *string {
	return s.PluginId
}

func (s *HttpApiPolicyConfigsAiToolSelectionConfigPluginStatus) GetServiceHealthy() *bool {
	return s.ServiceHealthy
}

func (s *HttpApiPolicyConfigsAiToolSelectionConfigPluginStatus) SetErrorLogs(v map[string]*string) *HttpApiPolicyConfigsAiToolSelectionConfigPluginStatus {
	s.ErrorLogs = v
	return s
}

func (s *HttpApiPolicyConfigsAiToolSelectionConfigPluginStatus) SetPluginId(v string) *HttpApiPolicyConfigsAiToolSelectionConfigPluginStatus {
	s.PluginId = &v
	return s
}

func (s *HttpApiPolicyConfigsAiToolSelectionConfigPluginStatus) SetServiceHealthy(v bool) *HttpApiPolicyConfigsAiToolSelectionConfigPluginStatus {
	s.ServiceHealthy = &v
	return s
}

func (s *HttpApiPolicyConfigsAiToolSelectionConfigPluginStatus) Validate() error {
	return dara.Validate(s)
}

type HttpApiPolicyConfigsAiToolSelectionConfigQueryRewriting struct {
	ContextSelection  *HttpApiPolicyConfigsAiToolSelectionConfigQueryRewritingContextSelection  `json:"contextSelection,omitempty" xml:"contextSelection,omitempty" type:"Struct"`
	Enabled           *bool                                                                     `json:"enabled,omitempty" xml:"enabled,omitempty"`
	FallbackStrategy  *string                                                                   `json:"fallbackStrategy,omitempty" xml:"fallbackStrategy,omitempty"`
	MaxOutputTokens   *int32                                                                    `json:"maxOutputTokens,omitempty" xml:"maxOutputTokens,omitempty"`
	ModelService      *HttpApiPolicyConfigsAiToolSelectionConfigQueryRewritingModelService      `json:"modelService,omitempty" xml:"modelService,omitempty" type:"Struct"`
	PromptConfig      *HttpApiPolicyConfigsAiToolSelectionConfigQueryRewritingPromptConfig      `json:"promptConfig,omitempty" xml:"promptConfig,omitempty" type:"Struct"`
	TriggerConditions *HttpApiPolicyConfigsAiToolSelectionConfigQueryRewritingTriggerConditions `json:"triggerConditions,omitempty" xml:"triggerConditions,omitempty" type:"Struct"`
}

func (s HttpApiPolicyConfigsAiToolSelectionConfigQueryRewriting) String() string {
	return dara.Prettify(s)
}

func (s HttpApiPolicyConfigsAiToolSelectionConfigQueryRewriting) GoString() string {
	return s.String()
}

func (s *HttpApiPolicyConfigsAiToolSelectionConfigQueryRewriting) GetContextSelection() *HttpApiPolicyConfigsAiToolSelectionConfigQueryRewritingContextSelection {
	return s.ContextSelection
}

func (s *HttpApiPolicyConfigsAiToolSelectionConfigQueryRewriting) GetEnabled() *bool {
	return s.Enabled
}

func (s *HttpApiPolicyConfigsAiToolSelectionConfigQueryRewriting) GetFallbackStrategy() *string {
	return s.FallbackStrategy
}

func (s *HttpApiPolicyConfigsAiToolSelectionConfigQueryRewriting) GetMaxOutputTokens() *int32 {
	return s.MaxOutputTokens
}

func (s *HttpApiPolicyConfigsAiToolSelectionConfigQueryRewriting) GetModelService() *HttpApiPolicyConfigsAiToolSelectionConfigQueryRewritingModelService {
	return s.ModelService
}

func (s *HttpApiPolicyConfigsAiToolSelectionConfigQueryRewriting) GetPromptConfig() *HttpApiPolicyConfigsAiToolSelectionConfigQueryRewritingPromptConfig {
	return s.PromptConfig
}

func (s *HttpApiPolicyConfigsAiToolSelectionConfigQueryRewriting) GetTriggerConditions() *HttpApiPolicyConfigsAiToolSelectionConfigQueryRewritingTriggerConditions {
	return s.TriggerConditions
}

func (s *HttpApiPolicyConfigsAiToolSelectionConfigQueryRewriting) SetContextSelection(v *HttpApiPolicyConfigsAiToolSelectionConfigQueryRewritingContextSelection) *HttpApiPolicyConfigsAiToolSelectionConfigQueryRewriting {
	s.ContextSelection = v
	return s
}

func (s *HttpApiPolicyConfigsAiToolSelectionConfigQueryRewriting) SetEnabled(v bool) *HttpApiPolicyConfigsAiToolSelectionConfigQueryRewriting {
	s.Enabled = &v
	return s
}

func (s *HttpApiPolicyConfigsAiToolSelectionConfigQueryRewriting) SetFallbackStrategy(v string) *HttpApiPolicyConfigsAiToolSelectionConfigQueryRewriting {
	s.FallbackStrategy = &v
	return s
}

func (s *HttpApiPolicyConfigsAiToolSelectionConfigQueryRewriting) SetMaxOutputTokens(v int32) *HttpApiPolicyConfigsAiToolSelectionConfigQueryRewriting {
	s.MaxOutputTokens = &v
	return s
}

func (s *HttpApiPolicyConfigsAiToolSelectionConfigQueryRewriting) SetModelService(v *HttpApiPolicyConfigsAiToolSelectionConfigQueryRewritingModelService) *HttpApiPolicyConfigsAiToolSelectionConfigQueryRewriting {
	s.ModelService = v
	return s
}

func (s *HttpApiPolicyConfigsAiToolSelectionConfigQueryRewriting) SetPromptConfig(v *HttpApiPolicyConfigsAiToolSelectionConfigQueryRewritingPromptConfig) *HttpApiPolicyConfigsAiToolSelectionConfigQueryRewriting {
	s.PromptConfig = v
	return s
}

func (s *HttpApiPolicyConfigsAiToolSelectionConfigQueryRewriting) SetTriggerConditions(v *HttpApiPolicyConfigsAiToolSelectionConfigQueryRewritingTriggerConditions) *HttpApiPolicyConfigsAiToolSelectionConfigQueryRewriting {
	s.TriggerConditions = v
	return s
}

func (s *HttpApiPolicyConfigsAiToolSelectionConfigQueryRewriting) Validate() error {
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

type HttpApiPolicyConfigsAiToolSelectionConfigQueryRewritingContextSelection struct {
	Type  *string `json:"type,omitempty" xml:"type,omitempty"`
	Value *int32  `json:"value,omitempty" xml:"value,omitempty"`
}

func (s HttpApiPolicyConfigsAiToolSelectionConfigQueryRewritingContextSelection) String() string {
	return dara.Prettify(s)
}

func (s HttpApiPolicyConfigsAiToolSelectionConfigQueryRewritingContextSelection) GoString() string {
	return s.String()
}

func (s *HttpApiPolicyConfigsAiToolSelectionConfigQueryRewritingContextSelection) GetType() *string {
	return s.Type
}

func (s *HttpApiPolicyConfigsAiToolSelectionConfigQueryRewritingContextSelection) GetValue() *int32 {
	return s.Value
}

func (s *HttpApiPolicyConfigsAiToolSelectionConfigQueryRewritingContextSelection) SetType(v string) *HttpApiPolicyConfigsAiToolSelectionConfigQueryRewritingContextSelection {
	s.Type = &v
	return s
}

func (s *HttpApiPolicyConfigsAiToolSelectionConfigQueryRewritingContextSelection) SetValue(v int32) *HttpApiPolicyConfigsAiToolSelectionConfigQueryRewritingContextSelection {
	s.Value = &v
	return s
}

func (s *HttpApiPolicyConfigsAiToolSelectionConfigQueryRewritingContextSelection) Validate() error {
	return dara.Validate(s)
}

type HttpApiPolicyConfigsAiToolSelectionConfigQueryRewritingModelService struct {
	ModelName          *string `json:"modelName,omitempty" xml:"modelName,omitempty"`
	ServiceId          *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	TimeoutMillisecond *int32  `json:"timeoutMillisecond,omitempty" xml:"timeoutMillisecond,omitempty"`
}

func (s HttpApiPolicyConfigsAiToolSelectionConfigQueryRewritingModelService) String() string {
	return dara.Prettify(s)
}

func (s HttpApiPolicyConfigsAiToolSelectionConfigQueryRewritingModelService) GoString() string {
	return s.String()
}

func (s *HttpApiPolicyConfigsAiToolSelectionConfigQueryRewritingModelService) GetModelName() *string {
	return s.ModelName
}

func (s *HttpApiPolicyConfigsAiToolSelectionConfigQueryRewritingModelService) GetServiceId() *string {
	return s.ServiceId
}

func (s *HttpApiPolicyConfigsAiToolSelectionConfigQueryRewritingModelService) GetTimeoutMillisecond() *int32 {
	return s.TimeoutMillisecond
}

func (s *HttpApiPolicyConfigsAiToolSelectionConfigQueryRewritingModelService) SetModelName(v string) *HttpApiPolicyConfigsAiToolSelectionConfigQueryRewritingModelService {
	s.ModelName = &v
	return s
}

func (s *HttpApiPolicyConfigsAiToolSelectionConfigQueryRewritingModelService) SetServiceId(v string) *HttpApiPolicyConfigsAiToolSelectionConfigQueryRewritingModelService {
	s.ServiceId = &v
	return s
}

func (s *HttpApiPolicyConfigsAiToolSelectionConfigQueryRewritingModelService) SetTimeoutMillisecond(v int32) *HttpApiPolicyConfigsAiToolSelectionConfigQueryRewritingModelService {
	s.TimeoutMillisecond = &v
	return s
}

func (s *HttpApiPolicyConfigsAiToolSelectionConfigQueryRewritingModelService) Validate() error {
	return dara.Validate(s)
}

type HttpApiPolicyConfigsAiToolSelectionConfigQueryRewritingPromptConfig struct {
	CustomPrompt *string `json:"customPrompt,omitempty" xml:"customPrompt,omitempty"`
	Type         *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s HttpApiPolicyConfigsAiToolSelectionConfigQueryRewritingPromptConfig) String() string {
	return dara.Prettify(s)
}

func (s HttpApiPolicyConfigsAiToolSelectionConfigQueryRewritingPromptConfig) GoString() string {
	return s.String()
}

func (s *HttpApiPolicyConfigsAiToolSelectionConfigQueryRewritingPromptConfig) GetCustomPrompt() *string {
	return s.CustomPrompt
}

func (s *HttpApiPolicyConfigsAiToolSelectionConfigQueryRewritingPromptConfig) GetType() *string {
	return s.Type
}

func (s *HttpApiPolicyConfigsAiToolSelectionConfigQueryRewritingPromptConfig) SetCustomPrompt(v string) *HttpApiPolicyConfigsAiToolSelectionConfigQueryRewritingPromptConfig {
	s.CustomPrompt = &v
	return s
}

func (s *HttpApiPolicyConfigsAiToolSelectionConfigQueryRewritingPromptConfig) SetType(v string) *HttpApiPolicyConfigsAiToolSelectionConfigQueryRewritingPromptConfig {
	s.Type = &v
	return s
}

func (s *HttpApiPolicyConfigsAiToolSelectionConfigQueryRewritingPromptConfig) Validate() error {
	return dara.Validate(s)
}

type HttpApiPolicyConfigsAiToolSelectionConfigQueryRewritingTriggerConditions struct {
	MessageCountThreshold *int32 `json:"messageCountThreshold,omitempty" xml:"messageCountThreshold,omitempty"`
}

func (s HttpApiPolicyConfigsAiToolSelectionConfigQueryRewritingTriggerConditions) String() string {
	return dara.Prettify(s)
}

func (s HttpApiPolicyConfigsAiToolSelectionConfigQueryRewritingTriggerConditions) GoString() string {
	return s.String()
}

func (s *HttpApiPolicyConfigsAiToolSelectionConfigQueryRewritingTriggerConditions) GetMessageCountThreshold() *int32 {
	return s.MessageCountThreshold
}

func (s *HttpApiPolicyConfigsAiToolSelectionConfigQueryRewritingTriggerConditions) SetMessageCountThreshold(v int32) *HttpApiPolicyConfigsAiToolSelectionConfigQueryRewritingTriggerConditions {
	s.MessageCountThreshold = &v
	return s
}

func (s *HttpApiPolicyConfigsAiToolSelectionConfigQueryRewritingTriggerConditions) Validate() error {
	return dara.Validate(s)
}

type HttpApiPolicyConfigsAiToolSelectionConfigToolReranking struct {
	FallbackStrategy *string                                                             `json:"fallbackStrategy,omitempty" xml:"fallbackStrategy,omitempty"`
	FilteringMethod  *string                                                             `json:"filteringMethod,omitempty" xml:"filteringMethod,omitempty"`
	ModelService     *HttpApiPolicyConfigsAiToolSelectionConfigToolRerankingModelService `json:"modelService,omitempty" xml:"modelService,omitempty" type:"Struct"`
	ScoreThreshold   *float32                                                            `json:"scoreThreshold,omitempty" xml:"scoreThreshold,omitempty"`
	TopKPercent      *int32                                                              `json:"topKPercent,omitempty" xml:"topKPercent,omitempty"`
	TopNCount        *int32                                                              `json:"topNCount,omitempty" xml:"topNCount,omitempty"`
}

func (s HttpApiPolicyConfigsAiToolSelectionConfigToolReranking) String() string {
	return dara.Prettify(s)
}

func (s HttpApiPolicyConfigsAiToolSelectionConfigToolReranking) GoString() string {
	return s.String()
}

func (s *HttpApiPolicyConfigsAiToolSelectionConfigToolReranking) GetFallbackStrategy() *string {
	return s.FallbackStrategy
}

func (s *HttpApiPolicyConfigsAiToolSelectionConfigToolReranking) GetFilteringMethod() *string {
	return s.FilteringMethod
}

func (s *HttpApiPolicyConfigsAiToolSelectionConfigToolReranking) GetModelService() *HttpApiPolicyConfigsAiToolSelectionConfigToolRerankingModelService {
	return s.ModelService
}

func (s *HttpApiPolicyConfigsAiToolSelectionConfigToolReranking) GetScoreThreshold() *float32 {
	return s.ScoreThreshold
}

func (s *HttpApiPolicyConfigsAiToolSelectionConfigToolReranking) GetTopKPercent() *int32 {
	return s.TopKPercent
}

func (s *HttpApiPolicyConfigsAiToolSelectionConfigToolReranking) GetTopNCount() *int32 {
	return s.TopNCount
}

func (s *HttpApiPolicyConfigsAiToolSelectionConfigToolReranking) SetFallbackStrategy(v string) *HttpApiPolicyConfigsAiToolSelectionConfigToolReranking {
	s.FallbackStrategy = &v
	return s
}

func (s *HttpApiPolicyConfigsAiToolSelectionConfigToolReranking) SetFilteringMethod(v string) *HttpApiPolicyConfigsAiToolSelectionConfigToolReranking {
	s.FilteringMethod = &v
	return s
}

func (s *HttpApiPolicyConfigsAiToolSelectionConfigToolReranking) SetModelService(v *HttpApiPolicyConfigsAiToolSelectionConfigToolRerankingModelService) *HttpApiPolicyConfigsAiToolSelectionConfigToolReranking {
	s.ModelService = v
	return s
}

func (s *HttpApiPolicyConfigsAiToolSelectionConfigToolReranking) SetScoreThreshold(v float32) *HttpApiPolicyConfigsAiToolSelectionConfigToolReranking {
	s.ScoreThreshold = &v
	return s
}

func (s *HttpApiPolicyConfigsAiToolSelectionConfigToolReranking) SetTopKPercent(v int32) *HttpApiPolicyConfigsAiToolSelectionConfigToolReranking {
	s.TopKPercent = &v
	return s
}

func (s *HttpApiPolicyConfigsAiToolSelectionConfigToolReranking) SetTopNCount(v int32) *HttpApiPolicyConfigsAiToolSelectionConfigToolReranking {
	s.TopNCount = &v
	return s
}

func (s *HttpApiPolicyConfigsAiToolSelectionConfigToolReranking) Validate() error {
	if s.ModelService != nil {
		if err := s.ModelService.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type HttpApiPolicyConfigsAiToolSelectionConfigToolRerankingModelService struct {
	ModelName          *string `json:"modelName,omitempty" xml:"modelName,omitempty"`
	ServiceId          *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	TimeoutMillisecond *int32  `json:"timeoutMillisecond,omitempty" xml:"timeoutMillisecond,omitempty"`
}

func (s HttpApiPolicyConfigsAiToolSelectionConfigToolRerankingModelService) String() string {
	return dara.Prettify(s)
}

func (s HttpApiPolicyConfigsAiToolSelectionConfigToolRerankingModelService) GoString() string {
	return s.String()
}

func (s *HttpApiPolicyConfigsAiToolSelectionConfigToolRerankingModelService) GetModelName() *string {
	return s.ModelName
}

func (s *HttpApiPolicyConfigsAiToolSelectionConfigToolRerankingModelService) GetServiceId() *string {
	return s.ServiceId
}

func (s *HttpApiPolicyConfigsAiToolSelectionConfigToolRerankingModelService) GetTimeoutMillisecond() *int32 {
	return s.TimeoutMillisecond
}

func (s *HttpApiPolicyConfigsAiToolSelectionConfigToolRerankingModelService) SetModelName(v string) *HttpApiPolicyConfigsAiToolSelectionConfigToolRerankingModelService {
	s.ModelName = &v
	return s
}

func (s *HttpApiPolicyConfigsAiToolSelectionConfigToolRerankingModelService) SetServiceId(v string) *HttpApiPolicyConfigsAiToolSelectionConfigToolRerankingModelService {
	s.ServiceId = &v
	return s
}

func (s *HttpApiPolicyConfigsAiToolSelectionConfigToolRerankingModelService) SetTimeoutMillisecond(v int32) *HttpApiPolicyConfigsAiToolSelectionConfigToolRerankingModelService {
	s.TimeoutMillisecond = &v
	return s
}

func (s *HttpApiPolicyConfigsAiToolSelectionConfigToolRerankingModelService) Validate() error {
	return dara.Validate(s)
}

type HttpApiPolicyConfigsSemanticRouterConfig struct {
	TimeoutMillisecond *int32 `json:"timeoutMillisecond,omitempty" xml:"timeoutMillisecond,omitempty"`
}

func (s HttpApiPolicyConfigsSemanticRouterConfig) String() string {
	return dara.Prettify(s)
}

func (s HttpApiPolicyConfigsSemanticRouterConfig) GoString() string {
	return s.String()
}

func (s *HttpApiPolicyConfigsSemanticRouterConfig) GetTimeoutMillisecond() *int32 {
	return s.TimeoutMillisecond
}

func (s *HttpApiPolicyConfigsSemanticRouterConfig) SetTimeoutMillisecond(v int32) *HttpApiPolicyConfigsSemanticRouterConfig {
	s.TimeoutMillisecond = &v
	return s
}

func (s *HttpApiPolicyConfigsSemanticRouterConfig) Validate() error {
	return dara.Validate(s)
}
