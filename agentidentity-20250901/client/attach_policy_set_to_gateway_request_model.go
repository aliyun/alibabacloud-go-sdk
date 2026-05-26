// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachPolicySetToGatewayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnforcementMode(v string) *AttachPolicySetToGatewayRequest
	GetEnforcementMode() *string
	SetGatewayArn(v string) *AttachPolicySetToGatewayRequest
	GetGatewayArn() *string
	SetGatewayType(v string) *AttachPolicySetToGatewayRequest
	GetGatewayType() *string
	SetPolicySetName(v string) *AttachPolicySetToGatewayRequest
	GetPolicySetName() *string
}

type AttachPolicySetToGatewayRequest struct {
	// example:
	//
	// ENFORCE
	EnforcementMode *string `json:"EnforcementMode,omitempty" xml:"EnforcementMode,omitempty"`
	// example:
	//
	// acs:agentidentity:cn-beijing:123456:gateway/my-gateway
	GatewayArn  *string `json:"GatewayArn,omitempty" xml:"GatewayArn,omitempty"`
	GatewayType *string `json:"GatewayType,omitempty" xml:"GatewayType,omitempty"`
	// example:
	//
	// default-policy-set
	PolicySetName *string `json:"PolicySetName,omitempty" xml:"PolicySetName,omitempty"`
}

func (s AttachPolicySetToGatewayRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachPolicySetToGatewayRequest) GoString() string {
	return s.String()
}

func (s *AttachPolicySetToGatewayRequest) GetEnforcementMode() *string {
	return s.EnforcementMode
}

func (s *AttachPolicySetToGatewayRequest) GetGatewayArn() *string {
	return s.GatewayArn
}

func (s *AttachPolicySetToGatewayRequest) GetGatewayType() *string {
	return s.GatewayType
}

func (s *AttachPolicySetToGatewayRequest) GetPolicySetName() *string {
	return s.PolicySetName
}

func (s *AttachPolicySetToGatewayRequest) SetEnforcementMode(v string) *AttachPolicySetToGatewayRequest {
	s.EnforcementMode = &v
	return s
}

func (s *AttachPolicySetToGatewayRequest) SetGatewayArn(v string) *AttachPolicySetToGatewayRequest {
	s.GatewayArn = &v
	return s
}

func (s *AttachPolicySetToGatewayRequest) SetGatewayType(v string) *AttachPolicySetToGatewayRequest {
	s.GatewayType = &v
	return s
}

func (s *AttachPolicySetToGatewayRequest) SetPolicySetName(v string) *AttachPolicySetToGatewayRequest {
	s.PolicySetName = &v
	return s
}

func (s *AttachPolicySetToGatewayRequest) Validate() error {
	return dara.Validate(s)
}
