// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublicTemplateRegistryNetworkConfig interface {
	dara.Model
	String() string
	GoString() string
	SetSecurityGroupId(v string) *PublicTemplateRegistryNetworkConfig
	GetSecurityGroupId() *string
	SetVSwitchId(v string) *PublicTemplateRegistryNetworkConfig
	GetVSwitchId() *string
	SetVpcId(v string) *PublicTemplateRegistryNetworkConfig
	GetVpcId() *string
}

type PublicTemplateRegistryNetworkConfig struct {
	SecurityGroupId *string `json:"securityGroupId,omitempty" xml:"securityGroupId,omitempty"`
	VSwitchId       *string `json:"vSwitchId,omitempty" xml:"vSwitchId,omitempty"`
	VpcId           *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s PublicTemplateRegistryNetworkConfig) String() string {
	return dara.Prettify(s)
}

func (s PublicTemplateRegistryNetworkConfig) GoString() string {
	return s.String()
}

func (s *PublicTemplateRegistryNetworkConfig) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *PublicTemplateRegistryNetworkConfig) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *PublicTemplateRegistryNetworkConfig) GetVpcId() *string {
	return s.VpcId
}

func (s *PublicTemplateRegistryNetworkConfig) SetSecurityGroupId(v string) *PublicTemplateRegistryNetworkConfig {
	s.SecurityGroupId = &v
	return s
}

func (s *PublicTemplateRegistryNetworkConfig) SetVSwitchId(v string) *PublicTemplateRegistryNetworkConfig {
	s.VSwitchId = &v
	return s
}

func (s *PublicTemplateRegistryNetworkConfig) SetVpcId(v string) *PublicTemplateRegistryNetworkConfig {
	s.VpcId = &v
	return s
}

func (s *PublicTemplateRegistryNetworkConfig) Validate() error {
	return dara.Validate(s)
}
