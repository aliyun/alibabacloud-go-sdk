// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublicTemplateVPCConfig interface {
	dara.Model
	String() string
	GoString() string
	SetRole(v string) *PublicTemplateVPCConfig
	GetRole() *string
	SetSecurityGroupId(v string) *PublicTemplateVPCConfig
	GetSecurityGroupId() *string
	SetVSwitchIds(v []*string) *PublicTemplateVPCConfig
	GetVSwitchIds() []*string
	SetVpcId(v string) *PublicTemplateVPCConfig
	GetVpcId() *string
}

type PublicTemplateVPCConfig struct {
	// The ARN of the RAM role that the sandbox instance assumes when accessing the VPC. The format is acs:ram::{account ID}:role/{role name}.
	//
	// example:
	//
	// acs:ram::1234567890123456:role/aliyunfcdefaultrole
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
	// The ID of the security group used by the sandbox instance.
	//
	// example:
	//
	// sg-****
	SecurityGroupId *string `json:"securityGroupId,omitempty" xml:"securityGroupId,omitempty"`
	// The list of vSwitch IDs that the sandbox instance connects to.
	VSwitchIds []*string `json:"vSwitchIds" xml:"vSwitchIds" type:"Repeated"`
	// The ID of the VPC that the sandbox instance connects to.
	//
	// example:
	//
	// vpc-****
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s PublicTemplateVPCConfig) String() string {
	return dara.Prettify(s)
}

func (s PublicTemplateVPCConfig) GoString() string {
	return s.String()
}

func (s *PublicTemplateVPCConfig) GetRole() *string {
	return s.Role
}

func (s *PublicTemplateVPCConfig) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *PublicTemplateVPCConfig) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *PublicTemplateVPCConfig) GetVpcId() *string {
	return s.VpcId
}

func (s *PublicTemplateVPCConfig) SetRole(v string) *PublicTemplateVPCConfig {
	s.Role = &v
	return s
}

func (s *PublicTemplateVPCConfig) SetSecurityGroupId(v string) *PublicTemplateVPCConfig {
	s.SecurityGroupId = &v
	return s
}

func (s *PublicTemplateVPCConfig) SetVSwitchIds(v []*string) *PublicTemplateVPCConfig {
	s.VSwitchIds = v
	return s
}

func (s *PublicTemplateVPCConfig) SetVpcId(v string) *PublicTemplateVPCConfig {
	s.VpcId = &v
	return s
}

func (s *PublicTemplateVPCConfig) Validate() error {
	return dara.Validate(s)
}
