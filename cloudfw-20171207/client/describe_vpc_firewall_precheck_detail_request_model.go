// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcFirewallPrecheckDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCenId(v string) *DescribeVpcFirewallPrecheckDetailRequest
	GetCenId() *string
	SetLang(v string) *DescribeVpcFirewallPrecheckDetailRequest
	GetLang() *string
	SetMemberUid(v string) *DescribeVpcFirewallPrecheckDetailRequest
	GetMemberUid() *string
	SetNetworkInstanceType(v string) *DescribeVpcFirewallPrecheckDetailRequest
	GetNetworkInstanceType() *string
	SetRegion(v string) *DescribeVpcFirewallPrecheckDetailRequest
	GetRegion() *string
	SetTransitRouterId(v string) *DescribeVpcFirewallPrecheckDetailRequest
	GetTransitRouterId() *string
	SetVpcId(v string) *DescribeVpcFirewallPrecheckDetailRequest
	GetVpcId() *string
}

type DescribeVpcFirewallPrecheckDetailRequest struct {
	// example:
	//
	// cen-hxsqf2bv6di1a****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 134388541648****
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// example:
	//
	// cen_firewall
	NetworkInstanceType *string `json:"NetworkInstanceType,omitempty" xml:"NetworkInstanceType,omitempty"`
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// example:
	//
	// tr-2vcn4u2g86tm72****
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
	// example:
	//
	// vpc-2zev8s8rxao33xt****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeVpcFirewallPrecheckDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallPrecheckDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallPrecheckDetailRequest) GetCenId() *string {
	return s.CenId
}

func (s *DescribeVpcFirewallPrecheckDetailRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeVpcFirewallPrecheckDetailRequest) GetMemberUid() *string {
	return s.MemberUid
}

func (s *DescribeVpcFirewallPrecheckDetailRequest) GetNetworkInstanceType() *string {
	return s.NetworkInstanceType
}

func (s *DescribeVpcFirewallPrecheckDetailRequest) GetRegion() *string {
	return s.Region
}

func (s *DescribeVpcFirewallPrecheckDetailRequest) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *DescribeVpcFirewallPrecheckDetailRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeVpcFirewallPrecheckDetailRequest) SetCenId(v string) *DescribeVpcFirewallPrecheckDetailRequest {
	s.CenId = &v
	return s
}

func (s *DescribeVpcFirewallPrecheckDetailRequest) SetLang(v string) *DescribeVpcFirewallPrecheckDetailRequest {
	s.Lang = &v
	return s
}

func (s *DescribeVpcFirewallPrecheckDetailRequest) SetMemberUid(v string) *DescribeVpcFirewallPrecheckDetailRequest {
	s.MemberUid = &v
	return s
}

func (s *DescribeVpcFirewallPrecheckDetailRequest) SetNetworkInstanceType(v string) *DescribeVpcFirewallPrecheckDetailRequest {
	s.NetworkInstanceType = &v
	return s
}

func (s *DescribeVpcFirewallPrecheckDetailRequest) SetRegion(v string) *DescribeVpcFirewallPrecheckDetailRequest {
	s.Region = &v
	return s
}

func (s *DescribeVpcFirewallPrecheckDetailRequest) SetTransitRouterId(v string) *DescribeVpcFirewallPrecheckDetailRequest {
	s.TransitRouterId = &v
	return s
}

func (s *DescribeVpcFirewallPrecheckDetailRequest) SetVpcId(v string) *DescribeVpcFirewallPrecheckDetailRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeVpcFirewallPrecheckDetailRequest) Validate() error {
	return dara.Validate(s)
}
