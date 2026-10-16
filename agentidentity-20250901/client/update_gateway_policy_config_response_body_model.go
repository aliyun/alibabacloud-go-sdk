// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayPolicyConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGatewayPolicyConfig(v *UpdateGatewayPolicyConfigResponseBodyGatewayPolicyConfig) *UpdateGatewayPolicyConfigResponseBody
	GetGatewayPolicyConfig() *UpdateGatewayPolicyConfigResponseBodyGatewayPolicyConfig
	SetRequestId(v string) *UpdateGatewayPolicyConfigResponseBody
	GetRequestId() *string
}

type UpdateGatewayPolicyConfigResponseBody struct {
	// The gateway policy configuration.
	GatewayPolicyConfig *UpdateGatewayPolicyConfigResponseBodyGatewayPolicyConfig `json:"GatewayPolicyConfig,omitempty" xml:"GatewayPolicyConfig,omitempty" type:"Struct"`
	// The unique identifier of the request.
	//
	// example:
	//
	// 2A48EB1D-D645-5758-91AF-EDF8E36E257B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateGatewayPolicyConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayPolicyConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGatewayPolicyConfigResponseBody) GetGatewayPolicyConfig() *UpdateGatewayPolicyConfigResponseBodyGatewayPolicyConfig {
	return s.GatewayPolicyConfig
}

func (s *UpdateGatewayPolicyConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateGatewayPolicyConfigResponseBody) SetGatewayPolicyConfig(v *UpdateGatewayPolicyConfigResponseBodyGatewayPolicyConfig) *UpdateGatewayPolicyConfigResponseBody {
	s.GatewayPolicyConfig = v
	return s
}

func (s *UpdateGatewayPolicyConfigResponseBody) SetRequestId(v string) *UpdateGatewayPolicyConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGatewayPolicyConfigResponseBody) Validate() error {
	if s.GatewayPolicyConfig != nil {
		if err := s.GatewayPolicyConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateGatewayPolicyConfigResponseBodyGatewayPolicyConfig struct {
	// The policy enforcement mode. Valid values:
	//
	// - ENFORCE: Intercepts non-compliant requests and allows compliant requests.
	//
	// - LOG_ONLY: Only logs non-compliant requests without intercepting them.
	//
	// example:
	//
	// ENFORCE
	EnforcementMode *string `json:"EnforcementMode,omitempty" xml:"EnforcementMode,omitempty"`
	// The ARN of the policy set.
	//
	// example:
	//
	// acs:agentidentity:cn-beijing:123456:policyset/default-policy-set
	PolicySetArn *string `json:"PolicySetArn,omitempty" xml:"PolicySetArn,omitempty"`
}

func (s UpdateGatewayPolicyConfigResponseBodyGatewayPolicyConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayPolicyConfigResponseBodyGatewayPolicyConfig) GoString() string {
	return s.String()
}

func (s *UpdateGatewayPolicyConfigResponseBodyGatewayPolicyConfig) GetEnforcementMode() *string {
	return s.EnforcementMode
}

func (s *UpdateGatewayPolicyConfigResponseBodyGatewayPolicyConfig) GetPolicySetArn() *string {
	return s.PolicySetArn
}

func (s *UpdateGatewayPolicyConfigResponseBodyGatewayPolicyConfig) SetEnforcementMode(v string) *UpdateGatewayPolicyConfigResponseBodyGatewayPolicyConfig {
	s.EnforcementMode = &v
	return s
}

func (s *UpdateGatewayPolicyConfigResponseBodyGatewayPolicyConfig) SetPolicySetArn(v string) *UpdateGatewayPolicyConfigResponseBodyGatewayPolicyConfig {
	s.PolicySetArn = &v
	return s
}

func (s *UpdateGatewayPolicyConfigResponseBodyGatewayPolicyConfig) Validate() error {
	return dara.Validate(s)
}
