// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGatewayPolicyConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGatewayPolicyConfig(v *GetGatewayPolicyConfigResponseBodyGatewayPolicyConfig) *GetGatewayPolicyConfigResponseBody
	GetGatewayPolicyConfig() *GetGatewayPolicyConfigResponseBodyGatewayPolicyConfig
	SetRequestId(v string) *GetGatewayPolicyConfigResponseBody
	GetRequestId() *string
}

type GetGatewayPolicyConfigResponseBody struct {
	GatewayPolicyConfig *GetGatewayPolicyConfigResponseBodyGatewayPolicyConfig `json:"GatewayPolicyConfig,omitempty" xml:"GatewayPolicyConfig,omitempty" type:"Struct"`
	// example:
	//
	// 2A48EB1D-D645-5758-91AF-EDF8E36E257B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetGatewayPolicyConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayPolicyConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetGatewayPolicyConfigResponseBody) GetGatewayPolicyConfig() *GetGatewayPolicyConfigResponseBodyGatewayPolicyConfig {
	return s.GatewayPolicyConfig
}

func (s *GetGatewayPolicyConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetGatewayPolicyConfigResponseBody) SetGatewayPolicyConfig(v *GetGatewayPolicyConfigResponseBodyGatewayPolicyConfig) *GetGatewayPolicyConfigResponseBody {
	s.GatewayPolicyConfig = v
	return s
}

func (s *GetGatewayPolicyConfigResponseBody) SetRequestId(v string) *GetGatewayPolicyConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGatewayPolicyConfigResponseBody) Validate() error {
	if s.GatewayPolicyConfig != nil {
		if err := s.GatewayPolicyConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetGatewayPolicyConfigResponseBodyGatewayPolicyConfig struct {
	// example:
	//
	// ENFORCE
	EnforcementMode *string `json:"EnforcementMode,omitempty" xml:"EnforcementMode,omitempty"`
	// example:
	//
	// acs:agentidentity:cn-beijing:123456:policyset/default-policy-set
	PolicySetArn *string `json:"PolicySetArn,omitempty" xml:"PolicySetArn,omitempty"`
}

func (s GetGatewayPolicyConfigResponseBodyGatewayPolicyConfig) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayPolicyConfigResponseBodyGatewayPolicyConfig) GoString() string {
	return s.String()
}

func (s *GetGatewayPolicyConfigResponseBodyGatewayPolicyConfig) GetEnforcementMode() *string {
	return s.EnforcementMode
}

func (s *GetGatewayPolicyConfigResponseBodyGatewayPolicyConfig) GetPolicySetArn() *string {
	return s.PolicySetArn
}

func (s *GetGatewayPolicyConfigResponseBodyGatewayPolicyConfig) SetEnforcementMode(v string) *GetGatewayPolicyConfigResponseBodyGatewayPolicyConfig {
	s.EnforcementMode = &v
	return s
}

func (s *GetGatewayPolicyConfigResponseBodyGatewayPolicyConfig) SetPolicySetArn(v string) *GetGatewayPolicyConfigResponseBodyGatewayPolicyConfig {
	s.PolicySetArn = &v
	return s
}

func (s *GetGatewayPolicyConfigResponseBodyGatewayPolicyConfig) Validate() error {
	return dara.Validate(s)
}
