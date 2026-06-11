// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAiTokenRateLimitConfig interface {
	dara.Model
	String() string
	GoString() string
	SetEnableGlobalRules(v bool) *AiTokenRateLimitConfig
	GetEnableGlobalRules() *bool
	SetGlobalRules(v []*AiTokenRateLimitConfigRule) *AiTokenRateLimitConfig
	GetGlobalRules() []*AiTokenRateLimitConfigRule
	SetPluginStatus(v *AiPluginStatus) *AiTokenRateLimitConfig
	GetPluginStatus() *AiPluginStatus
	SetRedisConfig(v *AiPolicyRedisConfig) *AiTokenRateLimitConfig
	GetRedisConfig() *AiPolicyRedisConfig
	SetRules(v []*AiTokenRateLimitConfigRule) *AiTokenRateLimitConfig
	GetRules() []*AiTokenRateLimitConfigRule
}

type AiTokenRateLimitConfig struct {
	// Controls whether global rules are enabled. If set to `true`, the rules in `globalRules` are applied. Defaults to `false`.
	EnableGlobalRules *bool `json:"enableGlobalRules,omitempty" xml:"enableGlobalRules,omitempty"`
	// A list of global rate limit rules. These rules are applied when no specific rule in `rules` is matched.
	GlobalRules []*AiTokenRateLimitConfigRule `json:"globalRules,omitempty" xml:"globalRules,omitempty" type:"Repeated"`
	// Specifies the status of the plugin, such as `enabled` or `disabled`.
	//
	// if can be null:
	// true
	PluginStatus *AiPluginStatus `json:"pluginStatus,omitempty" xml:"pluginStatus,omitempty"`
	// Specifies the Redis configuration for distributed rate limiting.
	//
	// if can be null:
	// true
	RedisConfig *AiPolicyRedisConfig `json:"redisConfig,omitempty" xml:"redisConfig,omitempty"`
	// A list of specific rate limit rules.
	Rules []*AiTokenRateLimitConfigRule `json:"rules,omitempty" xml:"rules,omitempty" type:"Repeated"`
}

func (s AiTokenRateLimitConfig) String() string {
	return dara.Prettify(s)
}

func (s AiTokenRateLimitConfig) GoString() string {
	return s.String()
}

func (s *AiTokenRateLimitConfig) GetEnableGlobalRules() *bool {
	return s.EnableGlobalRules
}

func (s *AiTokenRateLimitConfig) GetGlobalRules() []*AiTokenRateLimitConfigRule {
	return s.GlobalRules
}

func (s *AiTokenRateLimitConfig) GetPluginStatus() *AiPluginStatus {
	return s.PluginStatus
}

func (s *AiTokenRateLimitConfig) GetRedisConfig() *AiPolicyRedisConfig {
	return s.RedisConfig
}

func (s *AiTokenRateLimitConfig) GetRules() []*AiTokenRateLimitConfigRule {
	return s.Rules
}

func (s *AiTokenRateLimitConfig) SetEnableGlobalRules(v bool) *AiTokenRateLimitConfig {
	s.EnableGlobalRules = &v
	return s
}

func (s *AiTokenRateLimitConfig) SetGlobalRules(v []*AiTokenRateLimitConfigRule) *AiTokenRateLimitConfig {
	s.GlobalRules = v
	return s
}

func (s *AiTokenRateLimitConfig) SetPluginStatus(v *AiPluginStatus) *AiTokenRateLimitConfig {
	s.PluginStatus = v
	return s
}

func (s *AiTokenRateLimitConfig) SetRedisConfig(v *AiPolicyRedisConfig) *AiTokenRateLimitConfig {
	s.RedisConfig = v
	return s
}

func (s *AiTokenRateLimitConfig) SetRules(v []*AiTokenRateLimitConfigRule) *AiTokenRateLimitConfig {
	s.Rules = v
	return s
}

func (s *AiTokenRateLimitConfig) Validate() error {
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
