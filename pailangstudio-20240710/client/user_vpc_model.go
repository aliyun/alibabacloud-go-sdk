// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUserVpc interface {
	dara.Model
	String() string
	GoString() string
	SetSecurityGroupId(v string) *UserVpc
	GetSecurityGroupId() *string
	SetVSwitchId(v string) *UserVpc
	GetVSwitchId() *string
	SetVpcId(v string) *UserVpc
	GetVpcId() *string
}

type UserVpc struct {
	// 用户安全组的ID
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// 用户交换机的ID
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// 用户VPC的ID
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s UserVpc) String() string {
	return dara.Prettify(s)
}

func (s UserVpc) GoString() string {
	return s.String()
}

func (s *UserVpc) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *UserVpc) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *UserVpc) GetVpcId() *string {
	return s.VpcId
}

func (s *UserVpc) SetSecurityGroupId(v string) *UserVpc {
	s.SecurityGroupId = &v
	return s
}

func (s *UserVpc) SetVSwitchId(v string) *UserVpc {
	s.VSwitchId = &v
	return s
}

func (s *UserVpc) SetVpcId(v string) *UserVpc {
	s.VpcId = &v
	return s
}

func (s *UserVpc) Validate() error {
	return dara.Validate(s)
}
