// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCSecurityGroupListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeRCSecurityGroupListRequest
	GetRegionId() *string
	SetSecurityGroupId(v string) *DescribeRCSecurityGroupListRequest
	GetSecurityGroupId() *string
	SetVpcId(v string) *DescribeRCSecurityGroupListRequest
	GetVpcId() *string
}

type DescribeRCSecurityGroupListRequest struct {
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the security group.
	//
	// example:
	//
	// sg-2ze27hs990o2hn9****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The ID of the virtual private cloud (VPC) to which the security group belongs.
	//
	// example:
	//
	// vpc-bp1opxu1zkhn****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeRCSecurityGroupListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCSecurityGroupListRequest) GoString() string {
	return s.String()
}

func (s *DescribeRCSecurityGroupListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRCSecurityGroupListRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *DescribeRCSecurityGroupListRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeRCSecurityGroupListRequest) SetRegionId(v string) *DescribeRCSecurityGroupListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRCSecurityGroupListRequest) SetSecurityGroupId(v string) *DescribeRCSecurityGroupListRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeRCSecurityGroupListRequest) SetVpcId(v string) *DescribeRCSecurityGroupListRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeRCSecurityGroupListRequest) Validate() error {
	return dara.Validate(s)
}
