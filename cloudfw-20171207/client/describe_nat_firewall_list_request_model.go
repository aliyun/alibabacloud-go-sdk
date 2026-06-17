// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNatFirewallListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeNatFirewallListRequest
	GetLang() *string
	SetMemberUid(v int64) *DescribeNatFirewallListRequest
	GetMemberUid() *int64
	SetNatGatewayId(v string) *DescribeNatFirewallListRequest
	GetNatGatewayId() *string
	SetPageNo(v int64) *DescribeNatFirewallListRequest
	GetPageNo() *int64
	SetPageSize(v int64) *DescribeNatFirewallListRequest
	GetPageSize() *int64
	SetProxyId(v string) *DescribeNatFirewallListRequest
	GetProxyId() *string
	SetProxyName(v string) *DescribeNatFirewallListRequest
	GetProxyName() *string
	SetRegionNo(v string) *DescribeNatFirewallListRequest
	GetRegionNo() *string
	SetStatus(v string) *DescribeNatFirewallListRequest
	GetStatus() *string
	SetVpcId(v string) *DescribeNatFirewallListRequest
	GetVpcId() *string
}

type DescribeNatFirewallListRequest struct {
	// The language of the response message. Valid values:
	//
	// - **zh*	- (default): Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The UID of the member account.
	//
	// example:
	//
	// 147783******
	MemberUid *int64 `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The ID of the NAT Gateway.
	//
	// example:
	//
	// ngw-bp123456g******
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of NAT firewalls to return on each page for a paged query.
	//
	// The default value is **10**. The maximum value is **50**.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the NAT firewall.
	//
	// example:
	//
	// proxy-nat97a******
	ProxyId *string `json:"ProxyId,omitempty" xml:"ProxyId,omitempty"`
	// The name of the NAT firewall. The name must be 4 to 50 characters in length. It can contain letters, digits, and underscores (_). It cannot start with an underscore (_).
	//
	// example:
	//
	// nat-firewall
	ProxyName *string `json:"ProxyName,omitempty" xml:"ProxyName,omitempty"`
	// The region ID of the VPC.
	//
	// > For more information about the regions where Cloud Firewall is available, see [Supported regions](https://help.aliyun.com/document_detail/195657.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The status of the Cloud Firewall. Valid values:
	//
	// - configuring: The firewall is being created.
	//
	// - deleting: The firewall is being deleted.
	//
	// - normal: The firewall is working as expected.
	//
	// - abnormal: The firewall is not working as expected.
	//
	// - opening: The firewall is being enabled.
	//
	// - closing: The firewall is being disabled.
	//
	// - closed: The firewall is disabled.
	//
	// example:
	//
	// normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the VPC instance.
	//
	// example:
	//
	// vpc-8vbwbo90rq0anm6t****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeNatFirewallListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatFirewallListRequest) GoString() string {
	return s.String()
}

func (s *DescribeNatFirewallListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeNatFirewallListRequest) GetMemberUid() *int64 {
	return s.MemberUid
}

func (s *DescribeNatFirewallListRequest) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *DescribeNatFirewallListRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *DescribeNatFirewallListRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeNatFirewallListRequest) GetProxyId() *string {
	return s.ProxyId
}

func (s *DescribeNatFirewallListRequest) GetProxyName() *string {
	return s.ProxyName
}

func (s *DescribeNatFirewallListRequest) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeNatFirewallListRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeNatFirewallListRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeNatFirewallListRequest) SetLang(v string) *DescribeNatFirewallListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeNatFirewallListRequest) SetMemberUid(v int64) *DescribeNatFirewallListRequest {
	s.MemberUid = &v
	return s
}

func (s *DescribeNatFirewallListRequest) SetNatGatewayId(v string) *DescribeNatFirewallListRequest {
	s.NatGatewayId = &v
	return s
}

func (s *DescribeNatFirewallListRequest) SetPageNo(v int64) *DescribeNatFirewallListRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeNatFirewallListRequest) SetPageSize(v int64) *DescribeNatFirewallListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeNatFirewallListRequest) SetProxyId(v string) *DescribeNatFirewallListRequest {
	s.ProxyId = &v
	return s
}

func (s *DescribeNatFirewallListRequest) SetProxyName(v string) *DescribeNatFirewallListRequest {
	s.ProxyName = &v
	return s
}

func (s *DescribeNatFirewallListRequest) SetRegionNo(v string) *DescribeNatFirewallListRequest {
	s.RegionNo = &v
	return s
}

func (s *DescribeNatFirewallListRequest) SetStatus(v string) *DescribeNatFirewallListRequest {
	s.Status = &v
	return s
}

func (s *DescribeNatFirewallListRequest) SetVpcId(v string) *DescribeNatFirewallListRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeNatFirewallListRequest) Validate() error {
	return dara.Validate(s)
}
