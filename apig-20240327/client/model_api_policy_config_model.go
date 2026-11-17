// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelApiPolicyConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAiSecurityGuardConfig(v *AiSecurityGuardConfig) *ModelApiPolicyConfig
	GetAiSecurityGuardConfig() *AiSecurityGuardConfig
	SetAiStatisticsConfig(v *AiStatisticsConfig) *ModelApiPolicyConfig
	GetAiStatisticsConfig() *AiStatisticsConfig
	SetEnable(v bool) *ModelApiPolicyConfig
	GetEnable() *bool
	SetType(v string) *ModelApiPolicyConfig
	GetType() *string
}

type ModelApiPolicyConfig struct {
	// The AI security protection settings.
	AiSecurityGuardConfig *AiSecurityGuardConfig `json:"aiSecurityGuardConfig,omitempty" xml:"aiSecurityGuardConfig,omitempty"`
	// The AI observability and request log configuration.
	AiStatisticsConfig *AiStatisticsConfig `json:"aiStatisticsConfig,omitempty" xml:"aiStatisticsConfig,omitempty"`
	// Specifies whether to enable the policy.
	//
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// The policy type. Only AiStatistics and AiSecurityGuard are supported.
	//
	// example:
	//
	// AiStatistics
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ModelApiPolicyConfig) String() string {
	return dara.Prettify(s)
}

func (s ModelApiPolicyConfig) GoString() string {
	return s.String()
}

func (s *ModelApiPolicyConfig) GetAiSecurityGuardConfig() *AiSecurityGuardConfig {
	return s.AiSecurityGuardConfig
}

func (s *ModelApiPolicyConfig) GetAiStatisticsConfig() *AiStatisticsConfig {
	return s.AiStatisticsConfig
}

func (s *ModelApiPolicyConfig) GetEnable() *bool {
	return s.Enable
}

func (s *ModelApiPolicyConfig) GetType() *string {
	return s.Type
}

func (s *ModelApiPolicyConfig) SetAiSecurityGuardConfig(v *AiSecurityGuardConfig) *ModelApiPolicyConfig {
	s.AiSecurityGuardConfig = v
	return s
}

func (s *ModelApiPolicyConfig) SetAiStatisticsConfig(v *AiStatisticsConfig) *ModelApiPolicyConfig {
	s.AiStatisticsConfig = v
	return s
}

func (s *ModelApiPolicyConfig) SetEnable(v bool) *ModelApiPolicyConfig {
	s.Enable = &v
	return s
}

func (s *ModelApiPolicyConfig) SetType(v string) *ModelApiPolicyConfig {
	s.Type = &v
	return s
}

func (s *ModelApiPolicyConfig) Validate() error {
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
	return nil
}
