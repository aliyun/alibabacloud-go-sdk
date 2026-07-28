// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAutopilotPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnabled(v bool) *UpdateAutopilotPolicyRequest
	GetEnabled() *bool
	SetPolicyConfig(v *AutopilotPolicy) *UpdateAutopilotPolicyRequest
	GetPolicyConfig() *AutopilotPolicy
}

type UpdateAutopilotPolicyRequest struct {
	// Specifies whether to enable automatic tuning. A value of true enables automatic tuning (ACTIVE), and a value of false disables tuning (DISABLED). If this parameter is not specified, the current status is not changed.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The tuning policy configuration. This parameter uses full PUT mode: when specified, the complete policy object replaces the existing configuration entirely (fields not included are cleared). If this parameter is not specified, the existing configuration is retained.
	PolicyConfig *AutopilotPolicy `json:"policyConfig,omitempty" xml:"policyConfig,omitempty"`
}

func (s UpdateAutopilotPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAutopilotPolicyRequest) GoString() string {
	return s.String()
}

func (s *UpdateAutopilotPolicyRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateAutopilotPolicyRequest) GetPolicyConfig() *AutopilotPolicy {
	return s.PolicyConfig
}

func (s *UpdateAutopilotPolicyRequest) SetEnabled(v bool) *UpdateAutopilotPolicyRequest {
	s.Enabled = &v
	return s
}

func (s *UpdateAutopilotPolicyRequest) SetPolicyConfig(v *AutopilotPolicy) *UpdateAutopilotPolicyRequest {
	s.PolicyConfig = v
	return s
}

func (s *UpdateAutopilotPolicyRequest) Validate() error {
	if s.PolicyConfig != nil {
		if err := s.PolicyConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
