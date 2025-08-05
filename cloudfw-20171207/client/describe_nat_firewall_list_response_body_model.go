// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNatFirewallListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNatFirewallList(v []*DescribeNatFirewallListResponseBodyNatFirewallList) *DescribeNatFirewallListResponseBody
	GetNatFirewallList() []*DescribeNatFirewallListResponseBodyNatFirewallList
	SetRequestId(v string) *DescribeNatFirewallListResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeNatFirewallListResponseBody
	GetTotalCount() *int32
}

type DescribeNatFirewallListResponseBody struct {
	// The NAT firewalls.
	NatFirewallList []*DescribeNatFirewallListResponseBodyNatFirewallList `json:"NatFirewallList,omitempty" xml:"NatFirewallList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 15FCCC52-1E23-57AE-B5EF-3E00A3******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of NAT firewalls.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeNatFirewallListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatFirewallListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNatFirewallListResponseBody) GetNatFirewallList() []*DescribeNatFirewallListResponseBodyNatFirewallList {
	return s.NatFirewallList
}

func (s *DescribeNatFirewallListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNatFirewallListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeNatFirewallListResponseBody) SetNatFirewallList(v []*DescribeNatFirewallListResponseBodyNatFirewallList) *DescribeNatFirewallListResponseBody {
	s.NatFirewallList = v
	return s
}

func (s *DescribeNatFirewallListResponseBody) SetRequestId(v string) *DescribeNatFirewallListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNatFirewallListResponseBody) SetTotalCount(v int32) *DescribeNatFirewallListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeNatFirewallListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeNatFirewallListResponseBodyNatFirewallList struct {
	// The UID of the Alibaba Cloud account.
	//
	// >  The value of this parameter indicates the management account to which the member is added.
	//
	// example:
	//
	// 19106481******
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The cause of the error.
	//
	// example:
	//
	// Create Failed.
	ErrorDetail *string `json:"ErrorDetail,omitempty" xml:"ErrorDetail,omitempty"`
	// The UID of the member in Cloud Firewall.
	//
	// example:
	//
	// 19106481******
	MemberUid *int64 `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The ID of the NAT gateway.
	//
	// example:
	//
	// ngw-uf6tnblxip4qcxg******
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The name of the NAT gateway.
	//
	// example:
	//
	// nat-******
	NatGatewayName *string `json:"NatGatewayName,omitempty" xml:"NatGatewayName,omitempty"`
	// The default route entries of the NAT gateway.
	NatRouteEntryList []*DescribeNatFirewallListResponseBodyNatFirewallListNatRouteEntryList `json:"NatRouteEntryList,omitempty" xml:"NatRouteEntryList,omitempty" type:"Repeated"`
	// The ID of the NAT firewall.
	//
	// example:
	//
	// proxy-nat30******
	ProxyId *string `json:"ProxyId,omitempty" xml:"ProxyId,omitempty"`
	// The name of the NAT firewall.
	//
	// example:
	//
	// proxy-******
	ProxyName *string `json:"ProxyName,omitempty" xml:"ProxyName,omitempty"`
	// The status of the NAT firewall. Valid values:
	//
	// 	- configuring
	//
	// 	- deleting
	//
	// 	- normal
	//
	// 	- abnormal
	//
	// 	- opening
	//
	// 	- closing
	//
	// 	- closed
	//
	// example:
	//
	// normal
	ProxyStatus *string `json:"ProxyStatus,omitempty" xml:"ProxyStatus,omitempty"`
	// The region ID of your Cloud Firewall.
	//
	// >  For more information about the supported regions of Cloud Firewall, see [Supported regions](https://help.aliyun.com/document_detail/195657.html).
	//
	// example:
	//
	// cn-qingdao
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Indicates whether the strict mode is enabled. Valid values: 1, which specifies yes, and 0, which specifies no.
	//
	// example:
	//
	// 0
	StrictMode *int32 `json:"StrictMode,omitempty" xml:"StrictMode,omitempty"`
	// The ID of the VPC.
	//
	// example:
	//
	// vpc-2ze26ya******
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The name of the VPC.
	//
	// example:
	//
	// vpc-******
	VpcName *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
}

func (s DescribeNatFirewallListResponseBodyNatFirewallList) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatFirewallListResponseBodyNatFirewallList) GoString() string {
	return s.String()
}

func (s *DescribeNatFirewallListResponseBodyNatFirewallList) GetAliUid() *int64 {
	return s.AliUid
}

func (s *DescribeNatFirewallListResponseBodyNatFirewallList) GetErrorDetail() *string {
	return s.ErrorDetail
}

func (s *DescribeNatFirewallListResponseBodyNatFirewallList) GetMemberUid() *int64 {
	return s.MemberUid
}

func (s *DescribeNatFirewallListResponseBodyNatFirewallList) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *DescribeNatFirewallListResponseBodyNatFirewallList) GetNatGatewayName() *string {
	return s.NatGatewayName
}

func (s *DescribeNatFirewallListResponseBodyNatFirewallList) GetNatRouteEntryList() []*DescribeNatFirewallListResponseBodyNatFirewallListNatRouteEntryList {
	return s.NatRouteEntryList
}

func (s *DescribeNatFirewallListResponseBodyNatFirewallList) GetProxyId() *string {
	return s.ProxyId
}

func (s *DescribeNatFirewallListResponseBodyNatFirewallList) GetProxyName() *string {
	return s.ProxyName
}

func (s *DescribeNatFirewallListResponseBodyNatFirewallList) GetProxyStatus() *string {
	return s.ProxyStatus
}

func (s *DescribeNatFirewallListResponseBodyNatFirewallList) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeNatFirewallListResponseBodyNatFirewallList) GetStrictMode() *int32 {
	return s.StrictMode
}

func (s *DescribeNatFirewallListResponseBodyNatFirewallList) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeNatFirewallListResponseBodyNatFirewallList) GetVpcName() *string {
	return s.VpcName
}

func (s *DescribeNatFirewallListResponseBodyNatFirewallList) SetAliUid(v int64) *DescribeNatFirewallListResponseBodyNatFirewallList {
	s.AliUid = &v
	return s
}

func (s *DescribeNatFirewallListResponseBodyNatFirewallList) SetErrorDetail(v string) *DescribeNatFirewallListResponseBodyNatFirewallList {
	s.ErrorDetail = &v
	return s
}

func (s *DescribeNatFirewallListResponseBodyNatFirewallList) SetMemberUid(v int64) *DescribeNatFirewallListResponseBodyNatFirewallList {
	s.MemberUid = &v
	return s
}

func (s *DescribeNatFirewallListResponseBodyNatFirewallList) SetNatGatewayId(v string) *DescribeNatFirewallListResponseBodyNatFirewallList {
	s.NatGatewayId = &v
	return s
}

func (s *DescribeNatFirewallListResponseBodyNatFirewallList) SetNatGatewayName(v string) *DescribeNatFirewallListResponseBodyNatFirewallList {
	s.NatGatewayName = &v
	return s
}

func (s *DescribeNatFirewallListResponseBodyNatFirewallList) SetNatRouteEntryList(v []*DescribeNatFirewallListResponseBodyNatFirewallListNatRouteEntryList) *DescribeNatFirewallListResponseBodyNatFirewallList {
	s.NatRouteEntryList = v
	return s
}

func (s *DescribeNatFirewallListResponseBodyNatFirewallList) SetProxyId(v string) *DescribeNatFirewallListResponseBodyNatFirewallList {
	s.ProxyId = &v
	return s
}

func (s *DescribeNatFirewallListResponseBodyNatFirewallList) SetProxyName(v string) *DescribeNatFirewallListResponseBodyNatFirewallList {
	s.ProxyName = &v
	return s
}

func (s *DescribeNatFirewallListResponseBodyNatFirewallList) SetProxyStatus(v string) *DescribeNatFirewallListResponseBodyNatFirewallList {
	s.ProxyStatus = &v
	return s
}

func (s *DescribeNatFirewallListResponseBodyNatFirewallList) SetRegionId(v string) *DescribeNatFirewallListResponseBodyNatFirewallList {
	s.RegionId = &v
	return s
}

func (s *DescribeNatFirewallListResponseBodyNatFirewallList) SetStrictMode(v int32) *DescribeNatFirewallListResponseBodyNatFirewallList {
	s.StrictMode = &v
	return s
}

func (s *DescribeNatFirewallListResponseBodyNatFirewallList) SetVpcId(v string) *DescribeNatFirewallListResponseBodyNatFirewallList {
	s.VpcId = &v
	return s
}

func (s *DescribeNatFirewallListResponseBodyNatFirewallList) SetVpcName(v string) *DescribeNatFirewallListResponseBodyNatFirewallList {
	s.VpcName = &v
	return s
}

func (s *DescribeNatFirewallListResponseBodyNatFirewallList) Validate() error {
	return dara.Validate(s)
}

type DescribeNatFirewallListResponseBodyNatFirewallListNatRouteEntryList struct {
	// The destination CIDR block of the default route.
	//
	// example:
	//
	// 0.0.0.0/0
	DestinationCidr *string `json:"DestinationCidr,omitempty" xml:"DestinationCidr,omitempty"`
	// The next hop of the original NAT gateway.
	//
	// example:
	//
	// ngw-2ze0s284r9atg5******
	NextHopId *string `json:"NextHopId,omitempty" xml:"NextHopId,omitempty"`
	// The network type of the next hop. The value is fixed as NatGateway.
	//
	// example:
	//
	// NatGateway
	NextHopType *string `json:"NextHopType,omitempty" xml:"NextHopType,omitempty"`
	// The route table to which the default route of the NAT gateway belongs.
	//
	// example:
	//
	// vtb-bp18o0gb******
	RouteTableId *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
}

func (s DescribeNatFirewallListResponseBodyNatFirewallListNatRouteEntryList) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatFirewallListResponseBodyNatFirewallListNatRouteEntryList) GoString() string {
	return s.String()
}

func (s *DescribeNatFirewallListResponseBodyNatFirewallListNatRouteEntryList) GetDestinationCidr() *string {
	return s.DestinationCidr
}

func (s *DescribeNatFirewallListResponseBodyNatFirewallListNatRouteEntryList) GetNextHopId() *string {
	return s.NextHopId
}

func (s *DescribeNatFirewallListResponseBodyNatFirewallListNatRouteEntryList) GetNextHopType() *string {
	return s.NextHopType
}

func (s *DescribeNatFirewallListResponseBodyNatFirewallListNatRouteEntryList) GetRouteTableId() *string {
	return s.RouteTableId
}

func (s *DescribeNatFirewallListResponseBodyNatFirewallListNatRouteEntryList) SetDestinationCidr(v string) *DescribeNatFirewallListResponseBodyNatFirewallListNatRouteEntryList {
	s.DestinationCidr = &v
	return s
}

func (s *DescribeNatFirewallListResponseBodyNatFirewallListNatRouteEntryList) SetNextHopId(v string) *DescribeNatFirewallListResponseBodyNatFirewallListNatRouteEntryList {
	s.NextHopId = &v
	return s
}

func (s *DescribeNatFirewallListResponseBodyNatFirewallListNatRouteEntryList) SetNextHopType(v string) *DescribeNatFirewallListResponseBodyNatFirewallListNatRouteEntryList {
	s.NextHopType = &v
	return s
}

func (s *DescribeNatFirewallListResponseBodyNatFirewallListNatRouteEntryList) SetRouteTableId(v string) *DescribeNatFirewallListResponseBodyNatFirewallListNatRouteEntryList {
	s.RouteTableId = &v
	return s
}

func (s *DescribeNatFirewallListResponseBodyNatFirewallListNatRouteEntryList) Validate() error {
	return dara.Validate(s)
}
