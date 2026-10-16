// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayPolicyConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnforcementMode(v string) *UpdateGatewayPolicyConfigRequest
	GetEnforcementMode() *string
	SetGatewayArn(v string) *UpdateGatewayPolicyConfigRequest
	GetGatewayArn() *string
	SetGatewayType(v string) *UpdateGatewayPolicyConfigRequest
	GetGatewayType() *string
}

type UpdateGatewayPolicyConfigRequest struct {
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
	// The ARN of the target gateway.
	//
	// example:
	//
	// acs:agentidentity:cn-beijing:123456:gateway/my-gateway
	GatewayArn *string `json:"GatewayArn,omitempty" xml:"GatewayArn,omitempty"`
	// The gateway type. Valid values:
	//
	// - AIGateway: Alibaba Cloud AI gateway.
	//
	// - AgentRunMCP: AgentRun MCP tool.
	//
	// If you do not specify this parameter, the default value is AIGateway.
	//
	// example:
	//
	// AIGateway
	GatewayType *string `json:"GatewayType,omitempty" xml:"GatewayType,omitempty"`
}

func (s UpdateGatewayPolicyConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayPolicyConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateGatewayPolicyConfigRequest) GetEnforcementMode() *string {
	return s.EnforcementMode
}

func (s *UpdateGatewayPolicyConfigRequest) GetGatewayArn() *string {
	return s.GatewayArn
}

func (s *UpdateGatewayPolicyConfigRequest) GetGatewayType() *string {
	return s.GatewayType
}

func (s *UpdateGatewayPolicyConfigRequest) SetEnforcementMode(v string) *UpdateGatewayPolicyConfigRequest {
	s.EnforcementMode = &v
	return s
}

func (s *UpdateGatewayPolicyConfigRequest) SetGatewayArn(v string) *UpdateGatewayPolicyConfigRequest {
	s.GatewayArn = &v
	return s
}

func (s *UpdateGatewayPolicyConfigRequest) SetGatewayType(v string) *UpdateGatewayPolicyConfigRequest {
	s.GatewayType = &v
	return s
}

func (s *UpdateGatewayPolicyConfigRequest) Validate() error {
	return dara.Validate(s)
}
