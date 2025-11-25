// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFirewallVswitchResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFirewallType(v string) *DescribeFirewallVswitchResourcesRequest
	GetFirewallType() *string
	SetLang(v string) *DescribeFirewallVswitchResourcesRequest
	GetLang() *string
	SetRegionNo(v string) *DescribeFirewallVswitchResourcesRequest
	GetRegionNo() *string
	SetVpcId(v string) *DescribeFirewallVswitchResourcesRequest
	GetVpcId() *string
}

type DescribeFirewallVswitchResourcesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// internet
	FirewallType *string `json:"FirewallType,omitempty" xml:"FirewallType,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vpc-2ze36yb348axtnf****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeFirewallVswitchResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeFirewallVswitchResourcesRequest) GoString() string {
	return s.String()
}

func (s *DescribeFirewallVswitchResourcesRequest) GetFirewallType() *string {
	return s.FirewallType
}

func (s *DescribeFirewallVswitchResourcesRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeFirewallVswitchResourcesRequest) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeFirewallVswitchResourcesRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeFirewallVswitchResourcesRequest) SetFirewallType(v string) *DescribeFirewallVswitchResourcesRequest {
	s.FirewallType = &v
	return s
}

func (s *DescribeFirewallVswitchResourcesRequest) SetLang(v string) *DescribeFirewallVswitchResourcesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeFirewallVswitchResourcesRequest) SetRegionNo(v string) *DescribeFirewallVswitchResourcesRequest {
	s.RegionNo = &v
	return s
}

func (s *DescribeFirewallVswitchResourcesRequest) SetVpcId(v string) *DescribeFirewallVswitchResourcesRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeFirewallVswitchResourcesRequest) Validate() error {
	return dara.Validate(s)
}
