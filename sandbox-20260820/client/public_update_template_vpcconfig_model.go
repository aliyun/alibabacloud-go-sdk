// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublicUpdateTemplateVPCConfig interface {
	dara.Model
	String() string
	GoString() string
	SetRole(v string) *PublicUpdateTemplateVPCConfig
	GetRole() *string
	SetSecurityGroupId(v string) *PublicUpdateTemplateVPCConfig
	GetSecurityGroupId() *string
	SetVSwitchIds(v []*string) *PublicUpdateTemplateVPCConfig
	GetVSwitchIds() []*string
	SetVpcId(v string) *PublicUpdateTemplateVPCConfig
	GetVpcId() *string
}

type PublicUpdateTemplateVPCConfig struct {
	Role            *string   `json:"role,omitempty" xml:"role,omitempty"`
	SecurityGroupId *string   `json:"securityGroupId,omitempty" xml:"securityGroupId,omitempty"`
	VSwitchIds      []*string `json:"vSwitchIds,omitempty" xml:"vSwitchIds,omitempty" type:"Repeated"`
	VpcId           *string   `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s PublicUpdateTemplateVPCConfig) String() string {
	return dara.Prettify(s)
}

func (s PublicUpdateTemplateVPCConfig) GoString() string {
	return s.String()
}

func (s *PublicUpdateTemplateVPCConfig) GetRole() *string {
	return s.Role
}

func (s *PublicUpdateTemplateVPCConfig) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *PublicUpdateTemplateVPCConfig) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *PublicUpdateTemplateVPCConfig) GetVpcId() *string {
	return s.VpcId
}

func (s *PublicUpdateTemplateVPCConfig) SetRole(v string) *PublicUpdateTemplateVPCConfig {
	s.Role = &v
	return s
}

func (s *PublicUpdateTemplateVPCConfig) SetSecurityGroupId(v string) *PublicUpdateTemplateVPCConfig {
	s.SecurityGroupId = &v
	return s
}

func (s *PublicUpdateTemplateVPCConfig) SetVSwitchIds(v []*string) *PublicUpdateTemplateVPCConfig {
	s.VSwitchIds = v
	return s
}

func (s *PublicUpdateTemplateVPCConfig) SetVpcId(v string) *PublicUpdateTemplateVPCConfig {
	s.VpcId = &v
	return s
}

func (s *PublicUpdateTemplateVPCConfig) Validate() error {
	return dara.Validate(s)
}
