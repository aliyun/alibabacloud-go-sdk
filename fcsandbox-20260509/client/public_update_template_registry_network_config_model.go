// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublicUpdateTemplateRegistryNetworkConfig interface {
	dara.Model
	String() string
	GoString() string
	SetSecurityGroupId(v string) *PublicUpdateTemplateRegistryNetworkConfig
	GetSecurityGroupId() *string
	SetVSwitchId(v string) *PublicUpdateTemplateRegistryNetworkConfig
	GetVSwitchId() *string
	SetVpcId(v string) *PublicUpdateTemplateRegistryNetworkConfig
	GetVpcId() *string
}

type PublicUpdateTemplateRegistryNetworkConfig struct {
	// The ID of the security group for repository access.
	//
	// example:
	//
	// sg-bp1abc123
	SecurityGroupId *string `json:"securityGroupId,omitempty" xml:"securityGroupId,omitempty"`
	// The ID of the vSwitch where the repository resides.
	//
	// example:
	//
	// vsw-bp1abc123
	VSwitchId *string `json:"vSwitchId,omitempty" xml:"vSwitchId,omitempty"`
	// The ID of the VPC where the repository resides.
	//
	// example:
	//
	// vpc-bp1abc123
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s PublicUpdateTemplateRegistryNetworkConfig) String() string {
	return dara.Prettify(s)
}

func (s PublicUpdateTemplateRegistryNetworkConfig) GoString() string {
	return s.String()
}

func (s *PublicUpdateTemplateRegistryNetworkConfig) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *PublicUpdateTemplateRegistryNetworkConfig) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *PublicUpdateTemplateRegistryNetworkConfig) GetVpcId() *string {
	return s.VpcId
}

func (s *PublicUpdateTemplateRegistryNetworkConfig) SetSecurityGroupId(v string) *PublicUpdateTemplateRegistryNetworkConfig {
	s.SecurityGroupId = &v
	return s
}

func (s *PublicUpdateTemplateRegistryNetworkConfig) SetVSwitchId(v string) *PublicUpdateTemplateRegistryNetworkConfig {
	s.VSwitchId = &v
	return s
}

func (s *PublicUpdateTemplateRegistryNetworkConfig) SetVpcId(v string) *PublicUpdateTemplateRegistryNetworkConfig {
	s.VpcId = &v
	return s
}

func (s *PublicUpdateTemplateRegistryNetworkConfig) Validate() error {
	return dara.Validate(s)
}
