// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcFirewallCenListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCenId(v string) *DescribeVpcFirewallCenListRequest
	GetCenId() *string
	SetCurrentPage(v string) *DescribeVpcFirewallCenListRequest
	GetCurrentPage() *string
	SetFirewallSwitchStatus(v string) *DescribeVpcFirewallCenListRequest
	GetFirewallSwitchStatus() *string
	SetLang(v string) *DescribeVpcFirewallCenListRequest
	GetLang() *string
	SetMemberUid(v string) *DescribeVpcFirewallCenListRequest
	GetMemberUid() *string
	SetNetworkInstanceId(v string) *DescribeVpcFirewallCenListRequest
	GetNetworkInstanceId() *string
	SetOwnerId(v string) *DescribeVpcFirewallCenListRequest
	GetOwnerId() *string
	SetPageSize(v string) *DescribeVpcFirewallCenListRequest
	GetPageSize() *string
	SetRegionNo(v string) *DescribeVpcFirewallCenListRequest
	GetRegionNo() *string
	SetRouteMode(v string) *DescribeVpcFirewallCenListRequest
	GetRouteMode() *string
	SetTransitRouterType(v string) *DescribeVpcFirewallCenListRequest
	GetTransitRouterType() *string
	SetVpcFirewallId(v string) *DescribeVpcFirewallCenListRequest
	GetVpcFirewallId() *string
	SetVpcFirewallName(v string) *DescribeVpcFirewallCenListRequest
	GetVpcFirewallName() *string
}

type DescribeVpcFirewallCenListRequest struct {
	// The ID of the CEN instance.
	//
	// example:
	//
	// cen-x5jayxou71ad73****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The number of the page to return.
	//
	// Pages start from page 1. Default value: 1.
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
	// 	- **configured**: The VPC firewall is configured but is not enabled.
	//
	// > If you do not specify this parameter, VPC firewalls in all states are queried.
	//
	// example:
	//
	// opened
	FirewallSwitchStatus *string `json:"FirewallSwitchStatus,omitempty" xml:"FirewallSwitchStatus,omitempty"`
	// The language of the content within the response. Valid values:
	//
	// 	- **zh**: Chinese (default)
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The UID of the member that is managed by your Alibaba Cloud account. The member is also an Alibaba Cloud account.
	//
	// example:
	//
	// 258039427902****
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The ID of the network instance.
	//
	// example:
	//
	// vpc-8vbwbo90rq0anm6t****
	NetworkInstanceId *string `json:"NetworkInstanceId,omitempty" xml:"NetworkInstanceId,omitempty"`
	OwnerId           *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of entries to return on each page.
	//
	// Default value: 10. Maximum value: 50.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the VPC.
	//
	// > For more information about the regions, see [Supported regions](https://help.aliyun.com/document_detail/195657.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The routing mode of the VPC firewall. Valid values:
	//
	// 	- **auto**: automatic mode
	//
	// 	- **manual**: manual mode
	//
	// > If you do not specify this parameter, VPC firewalls in all routing modes are queried.
	//
	// example:
	//
	// auto
	RouteMode *string `json:"RouteMode,omitempty" xml:"RouteMode,omitempty"`
	// The type of the transit router. Valid values:
	//
	// 	- **Basic**: Basic Edition transit router
	//
	// 	- **Enterprise**: Enterprise Edition transit router
	//
	// example:
	//
	// Basic
	TransitRouterType *string `json:"TransitRouterType,omitempty" xml:"TransitRouterType,omitempty"`
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
}

func (s DescribeVpcFirewallCenListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallCenListRequest) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallCenListRequest) GetCenId() *string {
	return s.CenId
}

func (s *DescribeVpcFirewallCenListRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribeVpcFirewallCenListRequest) GetFirewallSwitchStatus() *string {
	return s.FirewallSwitchStatus
}

func (s *DescribeVpcFirewallCenListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeVpcFirewallCenListRequest) GetMemberUid() *string {
	return s.MemberUid
}

func (s *DescribeVpcFirewallCenListRequest) GetNetworkInstanceId() *string {
	return s.NetworkInstanceId
}

func (s *DescribeVpcFirewallCenListRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DescribeVpcFirewallCenListRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeVpcFirewallCenListRequest) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeVpcFirewallCenListRequest) GetRouteMode() *string {
	return s.RouteMode
}

func (s *DescribeVpcFirewallCenListRequest) GetTransitRouterType() *string {
	return s.TransitRouterType
}

func (s *DescribeVpcFirewallCenListRequest) GetVpcFirewallId() *string {
	return s.VpcFirewallId
}

func (s *DescribeVpcFirewallCenListRequest) GetVpcFirewallName() *string {
	return s.VpcFirewallName
}

func (s *DescribeVpcFirewallCenListRequest) SetCenId(v string) *DescribeVpcFirewallCenListRequest {
	s.CenId = &v
	return s
}

func (s *DescribeVpcFirewallCenListRequest) SetCurrentPage(v string) *DescribeVpcFirewallCenListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeVpcFirewallCenListRequest) SetFirewallSwitchStatus(v string) *DescribeVpcFirewallCenListRequest {
	s.FirewallSwitchStatus = &v
	return s
}

func (s *DescribeVpcFirewallCenListRequest) SetLang(v string) *DescribeVpcFirewallCenListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeVpcFirewallCenListRequest) SetMemberUid(v string) *DescribeVpcFirewallCenListRequest {
	s.MemberUid = &v
	return s
}

func (s *DescribeVpcFirewallCenListRequest) SetNetworkInstanceId(v string) *DescribeVpcFirewallCenListRequest {
	s.NetworkInstanceId = &v
	return s
}

func (s *DescribeVpcFirewallCenListRequest) SetOwnerId(v string) *DescribeVpcFirewallCenListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVpcFirewallCenListRequest) SetPageSize(v string) *DescribeVpcFirewallCenListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVpcFirewallCenListRequest) SetRegionNo(v string) *DescribeVpcFirewallCenListRequest {
	s.RegionNo = &v
	return s
}

func (s *DescribeVpcFirewallCenListRequest) SetRouteMode(v string) *DescribeVpcFirewallCenListRequest {
	s.RouteMode = &v
	return s
}

func (s *DescribeVpcFirewallCenListRequest) SetTransitRouterType(v string) *DescribeVpcFirewallCenListRequest {
	s.TransitRouterType = &v
	return s
}

func (s *DescribeVpcFirewallCenListRequest) SetVpcFirewallId(v string) *DescribeVpcFirewallCenListRequest {
	s.VpcFirewallId = &v
	return s
}

func (s *DescribeVpcFirewallCenListRequest) SetVpcFirewallName(v string) *DescribeVpcFirewallCenListRequest {
	s.VpcFirewallName = &v
	return s
}

func (s *DescribeVpcFirewallCenListRequest) Validate() error {
	return dara.Validate(s)
}
