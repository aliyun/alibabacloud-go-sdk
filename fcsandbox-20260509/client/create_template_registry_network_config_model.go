// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTemplateRegistryNetworkConfig interface {
	dara.Model
	String() string
	GoString() string
	SetSecurityGroupId(v string) *CreateTemplateRegistryNetworkConfig
	GetSecurityGroupId() *string
	SetVSwitchId(v string) *CreateTemplateRegistryNetworkConfig
	GetVSwitchId() *string
	SetVpcId(v string) *CreateTemplateRegistryNetworkConfig
	GetVpcId() *string
}

type CreateTemplateRegistryNetworkConfig struct {
	// The ID of the security group where the image repository resides.
	//
	// example:
	//
	// sg-bp1gx7yj8ud5mabcde
	SecurityGroupId *string `json:"securityGroupId,omitempty" xml:"securityGroupId,omitempty"`
	// The ID of the vSwitch where the image repository resides.
	//
	// example:
	//
	// vsw-bp1s5fnlk4jl2abcde
	VSwitchId *string `json:"vSwitchId,omitempty" xml:"vSwitchId,omitempty"`
	// The ID of the VPC where the image repository resides.
	//
	// example:
	//
	// vpc-bp1mwrqm3wkq7abcde
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s CreateTemplateRegistryNetworkConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateTemplateRegistryNetworkConfig) GoString() string {
	return s.String()
}

func (s *CreateTemplateRegistryNetworkConfig) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateTemplateRegistryNetworkConfig) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateTemplateRegistryNetworkConfig) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateTemplateRegistryNetworkConfig) SetSecurityGroupId(v string) *CreateTemplateRegistryNetworkConfig {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateTemplateRegistryNetworkConfig) SetVSwitchId(v string) *CreateTemplateRegistryNetworkConfig {
	s.VSwitchId = &v
	return s
}

func (s *CreateTemplateRegistryNetworkConfig) SetVpcId(v string) *CreateTemplateRegistryNetworkConfig {
	s.VpcId = &v
	return s
}

func (s *CreateTemplateRegistryNetworkConfig) Validate() error {
	return dara.Validate(s)
}
