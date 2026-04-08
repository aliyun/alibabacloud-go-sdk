// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAiCacheConfig interface {
	dara.Model
	String() string
	GoString() string
	SetCacheKeyStrategy(v string) *AiCacheConfig
	GetCacheKeyStrategy() *string
	SetCacheMode(v string) *AiCacheConfig
	GetCacheMode() *string
	SetCacheTTL(v int32) *AiCacheConfig
	GetCacheTTL() *int32
	SetEmbeddingConfig(v *AiCacheConfigEmbeddingConfig) *AiCacheConfig
	GetEmbeddingConfig() *AiCacheConfigEmbeddingConfig
	SetPluginStatus(v *AiPluginStatus) *AiCacheConfig
	GetPluginStatus() *AiPluginStatus
	SetRedisConfig(v *AiPolicyRedisConfig) *AiCacheConfig
	GetRedisConfig() *AiPolicyRedisConfig
	SetVectorConfig(v *AiCacheConfigVectorConfig) *AiCacheConfig
	GetVectorConfig() *AiCacheConfigVectorConfig
}

type AiCacheConfig struct {
	CacheKeyStrategy *string                       `json:"cacheKeyStrategy,omitempty" xml:"cacheKeyStrategy,omitempty"`
	CacheMode        *string                       `json:"cacheMode,omitempty" xml:"cacheMode,omitempty"`
	CacheTTL         *int32                        `json:"cacheTTL,omitempty" xml:"cacheTTL,omitempty"`
	EmbeddingConfig  *AiCacheConfigEmbeddingConfig `json:"embeddingConfig,omitempty" xml:"embeddingConfig,omitempty" type:"Struct"`
	// if can be null:
	// true
	PluginStatus *AiPluginStatus `json:"pluginStatus,omitempty" xml:"pluginStatus,omitempty"`
	// if can be null:
	// true
	RedisConfig  *AiPolicyRedisConfig       `json:"redisConfig,omitempty" xml:"redisConfig,omitempty"`
	VectorConfig *AiCacheConfigVectorConfig `json:"vectorConfig,omitempty" xml:"vectorConfig,omitempty" type:"Struct"`
}

func (s AiCacheConfig) String() string {
	return dara.Prettify(s)
}

func (s AiCacheConfig) GoString() string {
	return s.String()
}

func (s *AiCacheConfig) GetCacheKeyStrategy() *string {
	return s.CacheKeyStrategy
}

func (s *AiCacheConfig) GetCacheMode() *string {
	return s.CacheMode
}

func (s *AiCacheConfig) GetCacheTTL() *int32 {
	return s.CacheTTL
}

func (s *AiCacheConfig) GetEmbeddingConfig() *AiCacheConfigEmbeddingConfig {
	return s.EmbeddingConfig
}

func (s *AiCacheConfig) GetPluginStatus() *AiPluginStatus {
	return s.PluginStatus
}

func (s *AiCacheConfig) GetRedisConfig() *AiPolicyRedisConfig {
	return s.RedisConfig
}

func (s *AiCacheConfig) GetVectorConfig() *AiCacheConfigVectorConfig {
	return s.VectorConfig
}

func (s *AiCacheConfig) SetCacheKeyStrategy(v string) *AiCacheConfig {
	s.CacheKeyStrategy = &v
	return s
}

func (s *AiCacheConfig) SetCacheMode(v string) *AiCacheConfig {
	s.CacheMode = &v
	return s
}

func (s *AiCacheConfig) SetCacheTTL(v int32) *AiCacheConfig {
	s.CacheTTL = &v
	return s
}

func (s *AiCacheConfig) SetEmbeddingConfig(v *AiCacheConfigEmbeddingConfig) *AiCacheConfig {
	s.EmbeddingConfig = v
	return s
}

func (s *AiCacheConfig) SetPluginStatus(v *AiPluginStatus) *AiCacheConfig {
	s.PluginStatus = v
	return s
}

func (s *AiCacheConfig) SetRedisConfig(v *AiPolicyRedisConfig) *AiCacheConfig {
	s.RedisConfig = v
	return s
}

func (s *AiCacheConfig) SetVectorConfig(v *AiCacheConfigVectorConfig) *AiCacheConfig {
	s.VectorConfig = v
	return s
}

func (s *AiCacheConfig) Validate() error {
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

type AiCacheConfigEmbeddingConfig struct {
	ModelName *string `json:"modelName,omitempty" xml:"modelName,omitempty"`
	ServiceId *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	Timeout   *int32  `json:"timeout,omitempty" xml:"timeout,omitempty"`
	Type      *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s AiCacheConfigEmbeddingConfig) String() string {
	return dara.Prettify(s)
}

func (s AiCacheConfigEmbeddingConfig) GoString() string {
	return s.String()
}

func (s *AiCacheConfigEmbeddingConfig) GetModelName() *string {
	return s.ModelName
}

func (s *AiCacheConfigEmbeddingConfig) GetServiceId() *string {
	return s.ServiceId
}

func (s *AiCacheConfigEmbeddingConfig) GetTimeout() *int32 {
	return s.Timeout
}

func (s *AiCacheConfigEmbeddingConfig) GetType() *string {
	return s.Type
}

func (s *AiCacheConfigEmbeddingConfig) SetModelName(v string) *AiCacheConfigEmbeddingConfig {
	s.ModelName = &v
	return s
}

func (s *AiCacheConfigEmbeddingConfig) SetServiceId(v string) *AiCacheConfigEmbeddingConfig {
	s.ServiceId = &v
	return s
}

func (s *AiCacheConfigEmbeddingConfig) SetTimeout(v int32) *AiCacheConfigEmbeddingConfig {
	s.Timeout = &v
	return s
}

func (s *AiCacheConfigEmbeddingConfig) SetType(v string) *AiCacheConfigEmbeddingConfig {
	s.Type = &v
	return s
}

func (s *AiCacheConfigEmbeddingConfig) Validate() error {
	return dara.Validate(s)
}

type AiCacheConfigVectorConfig struct {
	ApiKey       *string  `json:"apiKey,omitempty" xml:"apiKey,omitempty"`
	CollectionId *string  `json:"collectionId,omitempty" xml:"collectionId,omitempty"`
	ServiceHost  *string  `json:"serviceHost,omitempty" xml:"serviceHost,omitempty"`
	Threshold    *float32 `json:"threshold,omitempty" xml:"threshold,omitempty"`
	Timeout      *int32   `json:"timeout,omitempty" xml:"timeout,omitempty"`
	Type         *string  `json:"type,omitempty" xml:"type,omitempty"`
}

func (s AiCacheConfigVectorConfig) String() string {
	return dara.Prettify(s)
}

func (s AiCacheConfigVectorConfig) GoString() string {
	return s.String()
}

func (s *AiCacheConfigVectorConfig) GetApiKey() *string {
	return s.ApiKey
}

func (s *AiCacheConfigVectorConfig) GetCollectionId() *string {
	return s.CollectionId
}

func (s *AiCacheConfigVectorConfig) GetServiceHost() *string {
	return s.ServiceHost
}

func (s *AiCacheConfigVectorConfig) GetThreshold() *float32 {
	return s.Threshold
}

func (s *AiCacheConfigVectorConfig) GetTimeout() *int32 {
	return s.Timeout
}

func (s *AiCacheConfigVectorConfig) GetType() *string {
	return s.Type
}

func (s *AiCacheConfigVectorConfig) SetApiKey(v string) *AiCacheConfigVectorConfig {
	s.ApiKey = &v
	return s
}

func (s *AiCacheConfigVectorConfig) SetCollectionId(v string) *AiCacheConfigVectorConfig {
	s.CollectionId = &v
	return s
}

func (s *AiCacheConfigVectorConfig) SetServiceHost(v string) *AiCacheConfigVectorConfig {
	s.ServiceHost = &v
	return s
}

func (s *AiCacheConfigVectorConfig) SetThreshold(v float32) *AiCacheConfigVectorConfig {
	s.Threshold = &v
	return s
}

func (s *AiCacheConfigVectorConfig) SetTimeout(v int32) *AiCacheConfigVectorConfig {
	s.Timeout = &v
	return s
}

func (s *AiCacheConfigVectorConfig) SetType(v string) *AiCacheConfigVectorConfig {
	s.Type = &v
	return s
}

func (s *AiCacheConfigVectorConfig) Validate() error {
	return dara.Validate(s)
}
