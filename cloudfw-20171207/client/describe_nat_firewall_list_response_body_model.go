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
	// The list of Cloud Firewalls.
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
	if s.NatFirewallList != nil {
		for _, item := range s.NatFirewallList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeNatFirewallListResponseBodyNatFirewallList struct {
	// The UID of the Alibaba Cloud account.
	//
	// > The management account of the Cloud Firewall member accounts.
	//
	// example:
	//
	// 19106481******
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The error details.
	//
	// example:
	//
	// Firewall creation failed
	ErrorDetail *string `json:"ErrorDetail,omitempty" xml:"ErrorDetail,omitempty"`
	// The deployment mode of the NAT firewall service. Valid values: **PrimaryStandby*	- (active/standby mode) and **MultiPrimary*	- (active-active mode).
	//
	// example:
	//
	// PrimaryStandby
	FirewallServiceMode *string `json:"FirewallServiceMode,omitempty" xml:"FirewallServiceMode,omitempty"`
	// The list of zone IDs used by the NAT firewall service.
	FirewallServiceZones []*string `json:"FirewallServiceZones,omitempty" xml:"FirewallServiceZones,omitempty" type:"Repeated"`
	// The UID of the Cloud Firewall member accounts.
	//
	// example:
	//
	// 19106481******
	MemberUid *int64 `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The ID of the NAT gateway to query.
	//
	// example:
	//
	// ngw-uf6tnblxip4qcxg******
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The name of the NAT gateway.
	//
	// example:
	//
	// nat-gateway-test
	NatGatewayName *string `json:"NatGatewayName,omitempty" xml:"NatGatewayName,omitempty"`
	// The list of default route entries for the NAT gateway.
	NatRouteEntryList []*DescribeNatFirewallListResponseBodyNatFirewallListNatRouteEntryList `json:"NatRouteEntryList,omitempty" xml:"NatRouteEntryList,omitempty" type:"Repeated"`
	// The NAT firewall ID.
	//
	// example:
	//
	// proxy-nat30******
	ProxyId *string `json:"ProxyId,omitempty" xml:"ProxyId,omitempty"`
	// The NAT firewall name.
	//
	// example:
	//
	// nat-firewall-test
	ProxyName *string `json:"ProxyName,omitempty" xml:"ProxyName,omitempty"`
	// The elastic network interface (ENI) ID used by the firewall.
	//
	// example:
	//
	// eni-bp127llmo4v5qju******
	ProxyNetworkInterfaceId *string `json:"ProxyNetworkInterfaceId,omitempty" xml:"ProxyNetworkInterfaceId,omitempty"`
	// The route table ID used by the firewall.
	//
	// example:
	//
	// vtb-bp1pmyga7p4j10a******
	ProxyRouteTableId *string `json:"ProxyRouteTableId,omitempty" xml:"ProxyRouteTableId,omitempty"`
	// The Cloud Firewall status. Valid values:
	//
	// - configuring: being created
	//
	// - deleting: being deleted
	//
	// - normal: normal
	//
	// - abnormal: abnormal
	//
	// - opening: being enabled
	//
	// - closing: being disabled
	//
	// - closed: disabled
	//
	// example:
	//
	// normal
	ProxyStatus *string `json:"ProxyStatus,omitempty" xml:"ProxyStatus,omitempty"`
	// The vSwitch ID used by the firewall.
	//
	// example:
	//
	// vsw-bp1amn3t1ktjjy8******
	ProxyVSwitchId *string `json:"ProxyVSwitchId,omitempty" xml:"ProxyVSwitchId,omitempty"`
	// The region ID of the Cloud Firewall.
	//
	// > For more information about the regions supported by Cloud Firewall, see [Supported regions](https://help.aliyun.com/document_detail/195657.html).
	//
	// example:
	//
	// cn-qingdao
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Indicates whether strict mode is enabled.
	//
	// - 1: Strict mode is enabled.
	//
	// - 0: Strict mode is disabled.
	//
	// example:
	//
	// 0
	StrictMode *int32 `json:"StrictMode,omitempty" xml:"StrictMode,omitempty"`
	// The VPC-connected instance ID.
	//
	// example:
	//
	// vpc-2ze26ya******
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The name of the VPC instance.
	//
	// example:
	//
	// vpc-test-instance
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

func (s *DescribeNatFirewallListResponseBodyNatFirewallList) GetFirewallServiceMode() *string {
	return s.FirewallServiceMode
}

func (s *DescribeNatFirewallListResponseBodyNatFirewallList) GetFirewallServiceZones() []*string {
	return s.FirewallServiceZones
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

func (s *DescribeNatFirewallListResponseBodyNatFirewallList) GetProxyNetworkInterfaceId() *string {
	return s.ProxyNetworkInterfaceId
}

func (s *DescribeNatFirewallListResponseBodyNatFirewallList) GetProxyRouteTableId() *string {
	return s.ProxyRouteTableId
}

func (s *DescribeNatFirewallListResponseBodyNatFirewallList) GetProxyStatus() *string {
	return s.ProxyStatus
}

func (s *DescribeNatFirewallListResponseBodyNatFirewallList) GetProxyVSwitchId() *string {
	return s.ProxyVSwitchId
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

func (s *DescribeNatFirewallListResponseBodyNatFirewallList) SetFirewallServiceMode(v string) *DescribeNatFirewallListResponseBodyNatFirewallList {
	s.FirewallServiceMode = &v
	return s
}

func (s *DescribeNatFirewallListResponseBodyNatFirewallList) SetFirewallServiceZones(v []*string) *DescribeNatFirewallListResponseBodyNatFirewallList {
	s.FirewallServiceZones = v
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

func (s *DescribeNatFirewallListResponseBodyNatFirewallList) SetProxyNetworkInterfaceId(v string) *DescribeNatFirewallListResponseBodyNatFirewallList {
	s.ProxyNetworkInterfaceId = &v
	return s
}

func (s *DescribeNatFirewallListResponseBodyNatFirewallList) SetProxyRouteTableId(v string) *DescribeNatFirewallListResponseBodyNatFirewallList {
	s.ProxyRouteTableId = &v
	return s
}

func (s *DescribeNatFirewallListResponseBodyNatFirewallList) SetProxyStatus(v string) *DescribeNatFirewallListResponseBodyNatFirewallList {
	s.ProxyStatus = &v
	return s
}

func (s *DescribeNatFirewallListResponseBodyNatFirewallList) SetProxyVSwitchId(v string) *DescribeNatFirewallListResponseBodyNatFirewallList {
	s.ProxyVSwitchId = &v
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
	if s.NatRouteEntryList != nil {
		for _, item := range s.NatRouteEntryList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeNatFirewallListResponseBodyNatFirewallListNatRouteEntryList struct {
	// The destination CIDR block of the default route.
	//
	// example:
	//
	// 0.0.0.0/0
	DestinationCidr *string `json:"DestinationCidr,omitempty" xml:"DestinationCidr,omitempty"`
	// The original next hop address of the NAT gateway.
	//
	// example:
	//
	// ngw-2ze0s284r9atg5******
	NextHopId *string `json:"NextHopId,omitempty" xml:"NextHopId,omitempty"`
	// The network type of the next hop. Valid values: NatGateway.
	//
	// example:
	//
	// NatGateway
	NextHopType *string `json:"NextHopType,omitempty" xml:"NextHopType,omitempty"`
	// The route table that contains the default route of the NAT gateway.
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
