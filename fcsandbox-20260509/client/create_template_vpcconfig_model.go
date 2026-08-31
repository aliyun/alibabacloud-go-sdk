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
