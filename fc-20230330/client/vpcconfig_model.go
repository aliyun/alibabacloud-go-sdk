// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVPCConfig interface {
	dara.Model
	String() string
	GoString() string
	SetRole(v string) *VPCConfig
	GetRole() *string
	SetSecurityGroupId(v string) *VPCConfig
	GetSecurityGroupId() *string
	SetVSwitchIds(v []*string) *VPCConfig
	GetVSwitchIds() []*string
	SetVpcId(v string) *VPCConfig
	GetVpcId() *string
}

type VPCConfig struct {
	// example:
	//
	// acs:ram::188077086902****:role/fc-test
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
	// example:
	//
	// sg-bp18hj1wtxgy3b0***
	SecurityGroupId *string   `json:"securityGroupId,omitempty" xml:"securityGroupId,omitempty"`
	VSwitchIds      []*string `json:"vSwitchIds" xml:"vSwitchIds" type:"Repeated"`
	// example:
	//
	// vpc-***
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s VPCConfig) String() string {
	return dara.Prettify(s)
}

func (s VPCConfig) GoString() string {
	return s.String()
}

func (s *VPCConfig) GetRole() *string {
	return s.Role
}

func (s *VPCConfig) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *VPCConfig) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *VPCConfig) GetVpcId() *string {
	return s.VpcId
}

func (s *VPCConfig) SetRole(v string) *VPCConfig {
	s.Role = &v
	return s
}

func (s *VPCConfig) SetSecurityGroupId(v string) *VPCConfig {
	s.SecurityGroupId = &v
	return s
}

func (s *VPCConfig) SetVSwitchIds(v []*string) *VPCConfig {
	s.VSwitchIds = v
	return s
}

func (s *VPCConfig) SetVpcId(v string) *VPCConfig {
	s.VpcId = &v
	return s
}

func (s *VPCConfig) Validate() error {
	return dara.Validate(s)
}
