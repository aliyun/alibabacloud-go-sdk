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
	SecurityGroupId *string `json:"securityGroupId,omitempty" xml:"securityGroupId,omitempty"`
	VSwitchId       *string `json:"vSwitchId,omitempty" xml:"vSwitchId,omitempty"`
	VpcId           *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
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
