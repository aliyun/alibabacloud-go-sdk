// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHttpApiPolicyConfigs interface {
	dara.Model
	String() string
	GoString() string
	SetAiCacheConfig(v *AiCacheConfig) *HttpApiPolicyConfigs
	GetAiCacheConfig() *AiCacheConfig
	SetAiFallbackConfig(v *AiFallbackConfig) *HttpApiPolicyConfigs
	GetAiFallbackConfig() *AiFallbackConfig
	SetAiNetworkSearchConfig(v *AiNetworkSearchConfig) *HttpApiPolicyConfigs
	GetAiNetworkSearchConfig() *AiNetworkSearchConfig
	SetAiSecurityGuardConfig(v *AiSecurityGuardConfig) *HttpApiPolicyConfigs
	GetAiSecurityGuardConfig() *AiSecurityGuardConfig
	SetAiStatisticsConfig(v *AiStatisticsConfig) *HttpApiPolicyConfigs
	GetAiStatisticsConfig() *AiStatisticsConfig
	SetAiTokenRateLimitConfig(v *AiTokenRateLimitConfig) *HttpApiPolicyConfigs
	GetAiTokenRateLimitConfig() *AiTokenRateLimitConfig
	SetAiToolSelectionConfig(v *AiToolSelectionConfig) *HttpApiPolicyConfigs
	GetAiToolSelectionConfig() *AiToolSelectionConfig
	SetEnable(v bool) *HttpApiPolicyConfigs
	GetEnable() *bool
	SetPolicyReference(v *HttpApiPolicyReference) *HttpApiPolicyConfigs
	GetPolicyReference() *HttpApiPolicyReference
	SetSemanticRouterConfig(v *HttpApiPolicyConfigsSemanticRouterConfig) *HttpApiPolicyConfigs
	GetSemanticRouterConfig() *HttpApiPolicyConfigsSemanticRouterConfig
	SetType(v string) *HttpApiPolicyConfigs
	GetType() *string
}

type HttpApiPolicyConfigs struct {
	// The AI cache configuration.
	//
	// if can be null:
	// true
	AiCacheConfig *AiCacheConfig `json:"aiCacheConfig,omitempty" xml:"aiCacheConfig,omitempty"`
	// The AI fallback configuration.
	//
	// if can be null:
	// false
	AiFallbackConfig *AiFallbackConfig `json:"aiFallbackConfig,omitempty" xml:"aiFallbackConfig,omitempty"`
	// The AI web search configuration.
	//
	// if can be null:
	// true
	AiNetworkSearchConfig *AiNetworkSearchConfig `json:"aiNetworkSearchConfig,omitempty" xml:"aiNetworkSearchConfig,omitempty"`
	// The AI security protection configuration.
	//
	// if can be null:
	// false
	AiSecurityGuardConfig *AiSecurityGuardConfig `json:"aiSecurityGuardConfig,omitempty" xml:"aiSecurityGuardConfig,omitempty"`
	// The AI statistics configuration.
	//
	// if can be null:
	// false
	AiStatisticsConfig *AiStatisticsConfig `json:"aiStatisticsConfig,omitempty" xml:"aiStatisticsConfig,omitempty"`
	// Deprecated
	//
	// The AI token rate limiting configuration.
	//
	// if can be null:
	// false
	AiTokenRateLimitConfig *AiTokenRateLimitConfig `json:"aiTokenRateLimitConfig,omitempty" xml:"aiTokenRateLimitConfig,omitempty"`
	// The AI tool selection configuration.
	//
	// if can be null:
	// true
	AiToolSelectionConfig *AiToolSelectionConfig `json:"aiToolSelectionConfig,omitempty" xml:"aiToolSelectionConfig,omitempty"`
	// Indicates whether the policy is enabled.
	//
	// example:
	//
	// false
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// The read-only compatible reference. GetHttpApi returns policyId/policyAttachmentId for ModelAPI AiTokenRateLimit. This must be stripped before write path persistence and is not used as a bind/unbind instruction.
	//
	// if can be null:
	// true
	PolicyReference *HttpApiPolicyReference `json:"policyReference,omitempty" xml:"policyReference,omitempty"`
	// The semantic routing configuration.
	//
	// if can be null:
	// false
	SemanticRouterConfig *HttpApiPolicyConfigsSemanticRouterConfig `json:"semanticRouterConfig,omitempty" xml:"semanticRouterConfig,omitempty" type:"Struct"`
	// The policy template type.
	//
	// example:
	//
	// K8S
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s HttpApiPolicyConfigs) String() string {
	return dara.Prettify(s)
}

func (s HttpApiPolicyConfigs) GoString() string {
	return s.String()
}

func (s *HttpApiPolicyConfigs) GetAiCacheConfig() *AiCacheConfig {
	return s.AiCacheConfig
}

func (s *HttpApiPolicyConfigs) GetAiFallbackConfig() *AiFallbackConfig {
	return s.AiFallbackConfig
}

func (s *HttpApiPolicyConfigs) GetAiNetworkSearchConfig() *AiNetworkSearchConfig {
	return s.AiNetworkSearchConfig
}

func (s *HttpApiPolicyConfigs) GetAiSecurityGuardConfig() *AiSecurityGuardConfig {
	return s.AiSecurityGuardConfig
}

func (s *HttpApiPolicyConfigs) GetAiStatisticsConfig() *AiStatisticsConfig {
	return s.AiStatisticsConfig
}

func (s *HttpApiPolicyConfigs) GetAiTokenRateLimitConfig() *AiTokenRateLimitConfig {
	return s.AiTokenRateLimitConfig
}

func (s *HttpApiPolicyConfigs) GetAiToolSelectionConfig() *AiToolSelectionConfig {
	return s.AiToolSelectionConfig
}

func (s *HttpApiPolicyConfigs) GetEnable() *bool {
	return s.Enable
}

func (s *HttpApiPolicyConfigs) GetPolicyReference() *HttpApiPolicyReference {
	return s.PolicyReference
}

func (s *HttpApiPolicyConfigs) GetSemanticRouterConfig() *HttpApiPolicyConfigsSemanticRouterConfig {
	return s.SemanticRouterConfig
}

func (s *HttpApiPolicyConfigs) GetType() *string {
	return s.Type
}

func (s *HttpApiPolicyConfigs) SetAiCacheConfig(v *AiCacheConfig) *HttpApiPolicyConfigs {
	s.AiCacheConfig = v
	return s
}

func (s *HttpApiPolicyConfigs) SetAiFallbackConfig(v *AiFallbackConfig) *HttpApiPolicyConfigs {
	s.AiFallbackConfig = v
	return s
}

func (s *HttpApiPolicyConfigs) SetAiNetworkSearchConfig(v *AiNetworkSearchConfig) *HttpApiPolicyConfigs {
	s.AiNetworkSearchConfig = v
	return s
}

func (s *HttpApiPolicyConfigs) SetAiSecurityGuardConfig(v *AiSecurityGuardConfig) *HttpApiPolicyConfigs {
	s.AiSecurityGuardConfig = v
	return s
}

func (s *HttpApiPolicyConfigs) SetAiStatisticsConfig(v *AiStatisticsConfig) *HttpApiPolicyConfigs {
	s.AiStatisticsConfig = v
	return s
}

func (s *HttpApiPolicyConfigs) SetAiTokenRateLimitConfig(v *AiTokenRateLimitConfig) *HttpApiPolicyConfigs {
	s.AiTokenRateLimitConfig = v
	return s
}

func (s *HttpApiPolicyConfigs) SetAiToolSelectionConfig(v *AiToolSelectionConfig) *HttpApiPolicyConfigs {
	s.AiToolSelectionConfig = v
	return s
}

func (s *HttpApiPolicyConfigs) SetEnable(v bool) *HttpApiPolicyConfigs {
	s.Enable = &v
	return s
}

func (s *HttpApiPolicyConfigs) SetPolicyReference(v *HttpApiPolicyReference) *HttpApiPolicyConfigs {
	s.PolicyReference = v
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
	if s.PolicyReference != nil {
		if err := s.PolicyReference.Validate(); err != nil {
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

type HttpApiPolicyConfigsSemanticRouterConfig struct {
	// The timeout period, in milliseconds.
	//
	// example:
	//
	// 2000
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
