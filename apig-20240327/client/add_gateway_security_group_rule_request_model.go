// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddGatewaySecurityGroupRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *AddGatewaySecurityGroupRuleRequest
	GetDescription() *string
	SetPortRanges(v []*string) *AddGatewaySecurityGroupRuleRequest
	GetPortRanges() []*string
	SetSecurityGroupId(v string) *AddGatewaySecurityGroupRuleRequest
	GetSecurityGroupId() *string
}

type AddGatewaySecurityGroupRuleRequest struct {
	// The description of security group rule N.
	//
	// example:
	//
	// ProductCenterSG
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The port ranges.
	PortRanges []*string `json:"portRanges,omitempty" xml:"portRanges,omitempty" type:"Repeated"`
	// The security group ID.
	//
	// example:
	//
	// sg-wz929kxhcdp****
	SecurityGroupId *string `json:"securityGroupId,omitempty" xml:"securityGroupId,omitempty"`
}

func (s AddGatewaySecurityGroupRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s AddGatewaySecurityGroupRuleRequest) GoString() string {
	return s.String()
}

func (s *AddGatewaySecurityGroupRuleRequest) GetDescription() *string {
	return s.Description
}

func (s *AddGatewaySecurityGroupRuleRequest) GetPortRanges() []*string {
	return s.PortRanges
}

func (s *AddGatewaySecurityGroupRuleRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *AddGatewaySecurityGroupRuleRequest) SetDescription(v string) *AddGatewaySecurityGroupRuleRequest {
	s.Description = &v
	return s
}

func (s *AddGatewaySecurityGroupRuleRequest) SetPortRanges(v []*string) *AddGatewaySecurityGroupRuleRequest {
	s.PortRanges = v
	return s
}

func (s *AddGatewaySecurityGroupRuleRequest) SetSecurityGroupId(v string) *AddGatewaySecurityGroupRuleRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *AddGatewaySecurityGroupRuleRequest) Validate() error {
	return dara.Validate(s)
}
