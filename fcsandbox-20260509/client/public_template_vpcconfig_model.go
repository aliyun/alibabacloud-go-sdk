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
	// example:
	//
	// AliyunFCSandboxDefaultRole
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
	// example:
	//
	// sg-bp1gx7yj8ud5mabcde
	SecurityGroupId *string   `json:"securityGroupId,omitempty" xml:"securityGroupId,omitempty"`
	VSwitchIds      []*string `json:"vSwitchIds,omitempty" xml:"vSwitchIds,omitempty" type:"Repeated"`
	// example:
	//
	// vpc-bp1mwrqm3wkq7abcde
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
