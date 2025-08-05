// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcFirewallListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectSubType(v string) *DescribeVpcFirewallListRequest
	GetConnectSubType() *string
	SetCurrentPage(v string) *DescribeVpcFirewallListRequest
	GetCurrentPage() *string
	SetFirewallSwitchStatus(v string) *DescribeVpcFirewallListRequest
	GetFirewallSwitchStatus() *string
	SetLang(v string) *DescribeVpcFirewallListRequest
	GetLang() *string
	SetMemberUid(v string) *DescribeVpcFirewallListRequest
	GetMemberUid() *string
	SetPageSize(v string) *DescribeVpcFirewallListRequest
	GetPageSize() *string
	SetPeerUid(v string) *DescribeVpcFirewallListRequest
	GetPeerUid() *string
	SetRegionNo(v string) *DescribeVpcFirewallListRequest
	GetRegionNo() *string
	SetVpcFirewallId(v string) *DescribeVpcFirewallListRequest
	GetVpcFirewallId() *string
	SetVpcFirewallName(v string) *DescribeVpcFirewallListRequest
	GetVpcFirewallName() *string
	SetVpcId(v string) *DescribeVpcFirewallListRequest
	GetVpcId() *string
}

type DescribeVpcFirewallListRequest struct {
	// The sub-type of the connection. Valid values:
	//
	// 	- **vpc2vpc**: Express Connect connection
	//
	// 	- **vpcpeer**: peer connection
	//
	// example:
	//
	// vpcpeer
	ConnectSubType *string `json:"ConnectSubType,omitempty" xml:"ConnectSubType,omitempty"`
	// The number of the page to return.
	//
	// Pages start from page **1**. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The status of the VPC firewall. Valid values:
	//
	// 	- **opened**: The VPC firewall is enabled.
	//
	// 	- **closed**: The VPC firewall is disabled.
	//
	// 	- **notconfigured**: The VPC firewall is not configured.
	//
	// 	- **configured**: The VPC firewall is configured.
	//
	// > If you do not specify this parameter, VPC firewalls in all states are queried.
	//
	// example:
	//
	// opened
	FirewallSwitchStatus *string `json:"FirewallSwitchStatus,omitempty" xml:"FirewallSwitchStatus,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh**: Chinese (default)
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The UID of the member that is managed by your Alibaba Cloud account.
	//
	// example:
	//
	// 258039427902****
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The number of entries to return on each page.
	//
	// Default value: **10**. Maximum value: **50**.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The UID of the Alibaba Cloud account to which the peer VPC belongs.
	//
	// example:
	//
	// 258039427902****
	PeerUid *string `json:"PeerUid,omitempty" xml:"PeerUid,omitempty"`
	// The region ID of the VPC.
	//
	// > For more information about the regions, see [Supported regions](https://help.aliyun.com/document_detail/195657.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The instance ID of the VPC firewall.
	//
	// example:
	//
	// vfw-m5e7dbc4y****
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
	// The instance name of the VPC firewall.
	//
	// example:
	//
	// Test firewall
	VpcFirewallName *string `json:"VpcFirewallName,omitempty" xml:"VpcFirewallName,omitempty"`
	// The ID of the VPC.
	//
	// example:
	//
	// vpc-8vbwbo90rq0anm6t****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeVpcFirewallListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallListRequest) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallListRequest) GetConnectSubType() *string {
	return s.ConnectSubType
}

func (s *DescribeVpcFirewallListRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribeVpcFirewallListRequest) GetFirewallSwitchStatus() *string {
	return s.FirewallSwitchStatus
}

func (s *DescribeVpcFirewallListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeVpcFirewallListRequest) GetMemberUid() *string {
	return s.MemberUid
}

func (s *DescribeVpcFirewallListRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeVpcFirewallListRequest) GetPeerUid() *string {
	return s.PeerUid
}

func (s *DescribeVpcFirewallListRequest) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeVpcFirewallListRequest) GetVpcFirewallId() *string {
	return s.VpcFirewallId
}

func (s *DescribeVpcFirewallListRequest) GetVpcFirewallName() *string {
	return s.VpcFirewallName
}

func (s *DescribeVpcFirewallListRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeVpcFirewallListRequest) SetConnectSubType(v string) *DescribeVpcFirewallListRequest {
	s.ConnectSubType = &v
	return s
}

func (s *DescribeVpcFirewallListRequest) SetCurrentPage(v string) *DescribeVpcFirewallListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeVpcFirewallListRequest) SetFirewallSwitchStatus(v string) *DescribeVpcFirewallListRequest {
	s.FirewallSwitchStatus = &v
	return s
}

func (s *DescribeVpcFirewallListRequest) SetLang(v string) *DescribeVpcFirewallListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeVpcFirewallListRequest) SetMemberUid(v string) *DescribeVpcFirewallListRequest {
	s.MemberUid = &v
	return s
}

func (s *DescribeVpcFirewallListRequest) SetPageSize(v string) *DescribeVpcFirewallListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVpcFirewallListRequest) SetPeerUid(v string) *DescribeVpcFirewallListRequest {
	s.PeerUid = &v
	return s
}

func (s *DescribeVpcFirewallListRequest) SetRegionNo(v string) *DescribeVpcFirewallListRequest {
	s.RegionNo = &v
	return s
}

func (s *DescribeVpcFirewallListRequest) SetVpcFirewallId(v string) *DescribeVpcFirewallListRequest {
	s.VpcFirewallId = &v
	return s
}

func (s *DescribeVpcFirewallListRequest) SetVpcFirewallName(v string) *DescribeVpcFirewallListRequest {
	s.VpcFirewallName = &v
	return s
}

func (s *DescribeVpcFirewallListRequest) SetVpcId(v string) *DescribeVpcFirewallListRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeVpcFirewallListRequest) Validate() error {
	return dara.Validate(s)
}
