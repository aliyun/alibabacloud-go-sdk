// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTemplateVPCConfig interface {
	dara.Model
	String() string
	GoString() string
	SetRole(v string) *CreateTemplateVPCConfig
	GetRole() *string
	SetSecurityGroupId(v string) *CreateTemplateVPCConfig
	GetSecurityGroupId() *string
	SetVSwitchIds(v []*string) *CreateTemplateVPCConfig
	GetVSwitchIds() []*string
	SetVpcId(v string) *CreateTemplateVPCConfig
	GetVpcId() *string
}

type CreateTemplateVPCConfig struct {
	// The ARN of the RAM role assumed by the sandboxed container instance when accessing the VPC. Format: acs:ram::{account ID}:role/{role name}.
	//
	// example:
	//
	// acs:ram::1234567890123456:role/aliyunfcdefaultrole
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
	// The ID of the security group used by the sandboxed container instance.
	//
	// example:
	//
	// sg-****
	SecurityGroupId *string `json:"securityGroupId,omitempty" xml:"securityGroupId,omitempty"`
	// The list of vSwitch IDs that the sandboxed container instance connects to. Configure vSwitches in multiple zones to improve availability.
	VSwitchIds []*string `json:"vSwitchIds" xml:"vSwitchIds" type:"Repeated"`
	// The ID of the VPC that the sandboxed container instance connects to.
	//
	// example:
	//
	// vpc-****
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s CreateTemplateVPCConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateTemplateVPCConfig) GoString() string {
	return s.String()
}

func (s *CreateTemplateVPCConfig) GetRole() *string {
	return s.Role
}

func (s *CreateTemplateVPCConfig) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateTemplateVPCConfig) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *CreateTemplateVPCConfig) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateTemplateVPCConfig) SetRole(v string) *CreateTemplateVPCConfig {
	s.Role = &v
	return s
}

func (s *CreateTemplateVPCConfig) SetSecurityGroupId(v string) *CreateTemplateVPCConfig {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateTemplateVPCConfig) SetVSwitchIds(v []*string) *CreateTemplateVPCConfig {
	s.VSwitchIds = v
	return s
}

func (s *CreateTemplateVPCConfig) SetVpcId(v string) *CreateTemplateVPCConfig {
	s.VpcId = &v
	return s
}

func (s *CreateTemplateVPCConfig) Validate() error {
	return dara.Validate(s)
}
