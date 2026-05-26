// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachPolicySetFromGatewayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGatewayArn(v string) *DetachPolicySetFromGatewayRequest
	GetGatewayArn() *string
	SetGatewayType(v string) *DetachPolicySetFromGatewayRequest
	GetGatewayType() *string
	SetPolicySetName(v string) *DetachPolicySetFromGatewayRequest
	GetPolicySetName() *string
}

type DetachPolicySetFromGatewayRequest struct {
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

func (s DetachPolicySetFromGatewayRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachPolicySetFromGatewayRequest) GoString() string {
	return s.String()
}

func (s *DetachPolicySetFromGatewayRequest) GetGatewayArn() *string {
	return s.GatewayArn
}

func (s *DetachPolicySetFromGatewayRequest) GetGatewayType() *string {
	return s.GatewayType
}

func (s *DetachPolicySetFromGatewayRequest) GetPolicySetName() *string {
	return s.PolicySetName
}

func (s *DetachPolicySetFromGatewayRequest) SetGatewayArn(v string) *DetachPolicySetFromGatewayRequest {
	s.GatewayArn = &v
	return s
}

func (s *DetachPolicySetFromGatewayRequest) SetGatewayType(v string) *DetachPolicySetFromGatewayRequest {
	s.GatewayType = &v
	return s
}

func (s *DetachPolicySetFromGatewayRequest) SetPolicySetName(v string) *DetachPolicySetFromGatewayRequest {
	s.PolicySetName = &v
	return s
}

func (s *DetachPolicySetFromGatewayRequest) Validate() error {
	return dara.Validate(s)
}
