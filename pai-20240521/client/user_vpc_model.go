// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUserVpc interface {
	dara.Model
	String() string
	GoString() string
	SetDefaultForwardInfo(v *ForwardInfo) *UserVpc
	GetDefaultForwardInfo() *ForwardInfo
	SetDefaultRoute(v string) *UserVpc
	GetDefaultRoute() *string
	SetExtendedCIDRs(v []*string) *UserVpc
	GetExtendedCIDRs() []*string
	SetRoleArn(v string) *UserVpc
	GetRoleArn() *string
	SetSecurityGroupId(v string) *UserVpc
	GetSecurityGroupId() *string
	SetSwitchId(v string) *UserVpc
	GetSwitchId() *string
	SetVpcId(v string) *UserVpc
	GetVpcId() *string
}

type UserVpc struct {
	DefaultForwardInfo *ForwardInfo `json:"DefaultForwardInfo,omitempty" xml:"DefaultForwardInfo,omitempty"`
	DefaultRoute       *string      `json:"DefaultRoute,omitempty" xml:"DefaultRoute,omitempty"`
	ExtendedCIDRs      []*string    `json:"ExtendedCIDRs,omitempty" xml:"ExtendedCIDRs,omitempty" type:"Repeated"`
	RoleArn            *string      `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	SecurityGroupId    *string      `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	SwitchId           *string      `json:"SwitchId,omitempty" xml:"SwitchId,omitempty"`
	VpcId              *string      `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s UserVpc) String() string {
	return dara.Prettify(s)
}

func (s UserVpc) GoString() string {
	return s.String()
}

func (s *UserVpc) GetDefaultForwardInfo() *ForwardInfo {
	return s.DefaultForwardInfo
}

func (s *UserVpc) GetDefaultRoute() *string {
	return s.DefaultRoute
}

func (s *UserVpc) GetExtendedCIDRs() []*string {
	return s.ExtendedCIDRs
}

func (s *UserVpc) GetRoleArn() *string {
	return s.RoleArn
}

func (s *UserVpc) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *UserVpc) GetSwitchId() *string {
	return s.SwitchId
}

func (s *UserVpc) GetVpcId() *string {
	return s.VpcId
}

func (s *UserVpc) SetDefaultForwardInfo(v *ForwardInfo) *UserVpc {
	s.DefaultForwardInfo = v
	return s
}

func (s *UserVpc) SetDefaultRoute(v string) *UserVpc {
	s.DefaultRoute = &v
	return s
}

func (s *UserVpc) SetExtendedCIDRs(v []*string) *UserVpc {
	s.ExtendedCIDRs = v
	return s
}

func (s *UserVpc) SetRoleArn(v string) *UserVpc {
	s.RoleArn = &v
	return s
}

func (s *UserVpc) SetSecurityGroupId(v string) *UserVpc {
	s.SecurityGroupId = &v
	return s
}

func (s *UserVpc) SetSwitchId(v string) *UserVpc {
	s.SwitchId = &v
	return s
}

func (s *UserVpc) SetVpcId(v string) *UserVpc {
	s.VpcId = &v
	return s
}

func (s *UserVpc) Validate() error {
	if s.DefaultForwardInfo != nil {
		if err := s.DefaultForwardInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}
