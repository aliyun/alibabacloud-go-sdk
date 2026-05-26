// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGatewayPolicyConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGatewayArn(v string) *GetGatewayPolicyConfigRequest
	GetGatewayArn() *string
	SetGatewayType(v string) *GetGatewayPolicyConfigRequest
	GetGatewayType() *string
}

type GetGatewayPolicyConfigRequest struct {
	// example:
	//
	// acs:agentidentity:cn-beijing:123456:gateway/my-gateway
	GatewayArn  *string `json:"GatewayArn,omitempty" xml:"GatewayArn,omitempty"`
	GatewayType *string `json:"GatewayType,omitempty" xml:"GatewayType,omitempty"`
}

func (s GetGatewayPolicyConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayPolicyConfigRequest) GoString() string {
	return s.String()
}

func (s *GetGatewayPolicyConfigRequest) GetGatewayArn() *string {
	return s.GatewayArn
}

func (s *GetGatewayPolicyConfigRequest) GetGatewayType() *string {
	return s.GatewayType
}

func (s *GetGatewayPolicyConfigRequest) SetGatewayArn(v string) *GetGatewayPolicyConfigRequest {
	s.GatewayArn = &v
	return s
}

func (s *GetGatewayPolicyConfigRequest) SetGatewayType(v string) *GetGatewayPolicyConfigRequest {
	s.GatewayType = &v
	return s
}

func (s *GetGatewayPolicyConfigRequest) Validate() error {
	return dara.Validate(s)
}
