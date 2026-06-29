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
	// Specifies whether to enable global (API-level) throttling rules. Default value: false.
	//
	// example:
	//
	// true
	EnableGlobalRules *bool `json:"enableGlobalRules,omitempty" xml:"enableGlobalRules,omitempty"`
	// The list of global throttling rules. Only the LimitByGlobal type is allowed.
	GlobalRules []*AiTokenRateLimitConfigRule `json:"globalRules,omitempty" xml:"globalRules,omitempty" type:"Repeated"`
	// The running status of the plugin.
	//
	// if can be null:
	// true
	PluginStatus *AiPluginStatus `json:"pluginStatus,omitempty" xml:"pluginStatus,omitempty"`
	// The Redis configuration.
	//
	// if can be null:
	// true
	RedisConfig *AiPolicyRedisConfig `json:"redisConfig,omitempty" xml:"redisConfig,omitempty"`
	// The list of throttling rules.
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
