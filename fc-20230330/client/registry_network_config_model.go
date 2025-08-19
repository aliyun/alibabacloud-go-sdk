// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegistryNetworkConfig interface {
	dara.Model
	String() string
	GoString() string
	SetSecurityGroupId(v string) *RegistryNetworkConfig
	GetSecurityGroupId() *string
	SetVSwitchId(v string) *RegistryNetworkConfig
	GetVSwitchId() *string
	SetVpcId(v string) *RegistryNetworkConfig
	GetVpcId() *string
}

type RegistryNetworkConfig struct {
	// example:
	//
	// sg-xxxxxxxxxxxxxx
	SecurityGroupId *string `json:"securityGroupId,omitempty" xml:"securityGroupId,omitempty"`
	// example:
	//
	// vsw-xxxxxxxxxxxxxx
	VSwitchId *string `json:"vSwitchId,omitempty" xml:"vSwitchId,omitempty"`
	// example:
	//
	// vpc-xxxxxxxxxxxxxx
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s RegistryNetworkConfig) String() string {
	return dara.Prettify(s)
}

func (s RegistryNetworkConfig) GoString() string {
	return s.String()
}

func (s *RegistryNetworkConfig) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *RegistryNetworkConfig) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *RegistryNetworkConfig) GetVpcId() *string {
	return s.VpcId
}

func (s *RegistryNetworkConfig) SetSecurityGroupId(v string) *RegistryNetworkConfig {
	s.SecurityGroupId = &v
	return s
}

func (s *RegistryNetworkConfig) SetVSwitchId(v string) *RegistryNetworkConfig {
	s.VSwitchId = &v
	return s
}

func (s *RegistryNetworkConfig) SetVpcId(v string) *RegistryNetworkConfig {
	s.VpcId = &v
	return s
}

func (s *RegistryNetworkConfig) Validate() error {
	return dara.Validate(s)
}
