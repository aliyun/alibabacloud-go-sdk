// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcFirewallDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidth(v int32) *DescribeVpcFirewallDetailResponseBody
	GetBandwidth() *int32
	SetConnectType(v string) *DescribeVpcFirewallDetailResponseBody
	GetConnectType() *string
	SetFirewallSwitchStatus(v string) *DescribeVpcFirewallDetailResponseBody
	GetFirewallSwitchStatus() *string
	SetLocalVpc(v *DescribeVpcFirewallDetailResponseBodyLocalVpc) *DescribeVpcFirewallDetailResponseBody
	GetLocalVpc() *DescribeVpcFirewallDetailResponseBodyLocalVpc
	SetMemberUid(v string) *DescribeVpcFirewallDetailResponseBody
	GetMemberUid() *string
	SetPeerVpc(v *DescribeVpcFirewallDetailResponseBodyPeerVpc) *DescribeVpcFirewallDetailResponseBody
	GetPeerVpc() *DescribeVpcFirewallDetailResponseBodyPeerVpc
	SetRequestId(v string) *DescribeVpcFirewallDetailResponseBody
	GetRequestId() *string
	SetVpcFirewallId(v string) *DescribeVpcFirewallDetailResponseBody
	GetVpcFirewallId() *string
	SetVpcFirewallName(v string) *DescribeVpcFirewallDetailResponseBody
	GetVpcFirewallName() *string
}

type DescribeVpcFirewallDetailResponseBody struct {
	// The bandwidth of the Express Connect circuit. Unit: Mbit/s.
	//
	// example:
	//
	// 2
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The connection type of the VPC firewall. The value is fixed as **expressconnect**, which indicates Express Connect circuits.
	//
	// example:
	//
	// expressconnect
	ConnectType *string `json:"ConnectType,omitempty" xml:"ConnectType,omitempty"`
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
	// example:
	//
	// opened
	FirewallSwitchStatus *string `json:"FirewallSwitchStatus,omitempty" xml:"FirewallSwitchStatus,omitempty"`
	// The details about the local VPC.
	LocalVpc *DescribeVpcFirewallDetailResponseBodyLocalVpc `json:"LocalVpc,omitempty" xml:"LocalVpc,omitempty" type:"Struct"`
	// The UID of the member that is managed by your Alibaba Cloud account.
	//
	// example:
	//
	// 258039427902****
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The details about the peer VPC.
	PeerVpc *DescribeVpcFirewallDetailResponseBodyPeerVpc `json:"PeerVpc,omitempty" xml:"PeerVpc,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 850A84D6-0DE4-4797-A1E8-00090125g4d2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// tf-test
	VpcFirewallName *string `json:"VpcFirewallName,omitempty" xml:"VpcFirewallName,omitempty"`
}

func (s DescribeVpcFirewallDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallDetailResponseBody) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *DescribeVpcFirewallDetailResponseBody) GetConnectType() *string {
	return s.ConnectType
}

func (s *DescribeVpcFirewallDetailResponseBody) GetFirewallSwitchStatus() *string {
	return s.FirewallSwitchStatus
}

func (s *DescribeVpcFirewallDetailResponseBody) GetLocalVpc() *DescribeVpcFirewallDetailResponseBodyLocalVpc {
	return s.LocalVpc
}

func (s *DescribeVpcFirewallDetailResponseBody) GetMemberUid() *string {
	return s.MemberUid
}

func (s *DescribeVpcFirewallDetailResponseBody) GetPeerVpc() *DescribeVpcFirewallDetailResponseBodyPeerVpc {
	return s.PeerVpc
}

func (s *DescribeVpcFirewallDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVpcFirewallDetailResponseBody) GetVpcFirewallId() *string {
	return s.VpcFirewallId
}

func (s *DescribeVpcFirewallDetailResponseBody) GetVpcFirewallName() *string {
	return s.VpcFirewallName
}

func (s *DescribeVpcFirewallDetailResponseBody) SetBandwidth(v int32) *DescribeVpcFirewallDetailResponseBody {
	s.Bandwidth = &v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBody) SetConnectType(v string) *DescribeVpcFirewallDetailResponseBody {
	s.ConnectType = &v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBody) SetFirewallSwitchStatus(v string) *DescribeVpcFirewallDetailResponseBody {
	s.FirewallSwitchStatus = &v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBody) SetLocalVpc(v *DescribeVpcFirewallDetailResponseBodyLocalVpc) *DescribeVpcFirewallDetailResponseBody {
	s.LocalVpc = v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBody) SetMemberUid(v string) *DescribeVpcFirewallDetailResponseBody {
	s.MemberUid = &v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBody) SetPeerVpc(v *DescribeVpcFirewallDetailResponseBodyPeerVpc) *DescribeVpcFirewallDetailResponseBody {
	s.PeerVpc = v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBody) SetRequestId(v string) *DescribeVpcFirewallDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBody) SetVpcFirewallId(v string) *DescribeVpcFirewallDetailResponseBody {
	s.VpcFirewallId = &v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBody) SetVpcFirewallName(v string) *DescribeVpcFirewallDetailResponseBody {
	s.VpcFirewallName = &v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBody) Validate() error {
	if s.LocalVpc != nil {
		if err := s.LocalVpc.Validate(); err != nil {
			return err
		}
	}
	if s.PeerVpc != nil {
		if err := s.PeerVpc.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVpcFirewallDetailResponseBodyLocalVpc struct {
	// The ID of the ENI for the local VPC.
	//
	// example:
	//
	// eni-8vbhfosfqv2rff42****
	EniId *string `json:"EniId,omitempty" xml:"EniId,omitempty"`
	// The private IP address of the elastic network interface (ENI) for the local VPC.
	//
	// example:
	//
	// 192.168.XX.XX
	EniPrivateIpAddress *string `json:"EniPrivateIpAddress,omitempty" xml:"EniPrivateIpAddress,omitempty"`
	// The region ID of the local VPC.
	//
	// example:
	//
	// cn-hangzhou
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The router interface ID of the local VPC.
	//
	// example:
	//
	// vrt-m5eb5me6c3l5sezae****
	RouterInterfaceId *string `json:"RouterInterfaceId,omitempty" xml:"RouterInterfaceId,omitempty"`
	// The CIDR blocks of the local VPC.
	VpcCidrTableList []*DescribeVpcFirewallDetailResponseBodyLocalVpcVpcCidrTableList `json:"VpcCidrTableList,omitempty" xml:"VpcCidrTableList,omitempty" type:"Repeated"`
	// The ID of the local VPC.
	//
	// example:
	//
	// vpc-8vbwbo90rq0anm6t****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The name of the local VPC.
	//
	// example:
	//
	// Vitasoy
	VpcName *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
}

func (s DescribeVpcFirewallDetailResponseBodyLocalVpc) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallDetailResponseBodyLocalVpc) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallDetailResponseBodyLocalVpc) GetEniId() *string {
	return s.EniId
}

func (s *DescribeVpcFirewallDetailResponseBodyLocalVpc) GetEniPrivateIpAddress() *string {
	return s.EniPrivateIpAddress
}

func (s *DescribeVpcFirewallDetailResponseBodyLocalVpc) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeVpcFirewallDetailResponseBodyLocalVpc) GetRouterInterfaceId() *string {
	return s.RouterInterfaceId
}

func (s *DescribeVpcFirewallDetailResponseBodyLocalVpc) GetVpcCidrTableList() []*DescribeVpcFirewallDetailResponseBodyLocalVpcVpcCidrTableList {
	return s.VpcCidrTableList
}

func (s *DescribeVpcFirewallDetailResponseBodyLocalVpc) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeVpcFirewallDetailResponseBodyLocalVpc) GetVpcName() *string {
	return s.VpcName
}

func (s *DescribeVpcFirewallDetailResponseBodyLocalVpc) SetEniId(v string) *DescribeVpcFirewallDetailResponseBodyLocalVpc {
	s.EniId = &v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBodyLocalVpc) SetEniPrivateIpAddress(v string) *DescribeVpcFirewallDetailResponseBodyLocalVpc {
	s.EniPrivateIpAddress = &v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBodyLocalVpc) SetRegionNo(v string) *DescribeVpcFirewallDetailResponseBodyLocalVpc {
	s.RegionNo = &v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBodyLocalVpc) SetRouterInterfaceId(v string) *DescribeVpcFirewallDetailResponseBodyLocalVpc {
	s.RouterInterfaceId = &v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBodyLocalVpc) SetVpcCidrTableList(v []*DescribeVpcFirewallDetailResponseBodyLocalVpcVpcCidrTableList) *DescribeVpcFirewallDetailResponseBodyLocalVpc {
	s.VpcCidrTableList = v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBodyLocalVpc) SetVpcId(v string) *DescribeVpcFirewallDetailResponseBodyLocalVpc {
	s.VpcId = &v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBodyLocalVpc) SetVpcName(v string) *DescribeVpcFirewallDetailResponseBodyLocalVpc {
	s.VpcName = &v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBodyLocalVpc) Validate() error {
	if s.VpcCidrTableList != nil {
		for _, item := range s.VpcCidrTableList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVpcFirewallDetailResponseBodyLocalVpcVpcCidrTableList struct {
	// The route entries of the local VPC.
	RouteEntryList []*DescribeVpcFirewallDetailResponseBodyLocalVpcVpcCidrTableListRouteEntryList `json:"RouteEntryList,omitempty" xml:"RouteEntryList,omitempty" type:"Repeated"`
	// The ID of the route table for the local VPC.
	//
	// example:
	//
	// vtb-1234
	RouteTableId *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
}

func (s DescribeVpcFirewallDetailResponseBodyLocalVpcVpcCidrTableList) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallDetailResponseBodyLocalVpcVpcCidrTableList) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallDetailResponseBodyLocalVpcVpcCidrTableList) GetRouteEntryList() []*DescribeVpcFirewallDetailResponseBodyLocalVpcVpcCidrTableListRouteEntryList {
	return s.RouteEntryList
}

func (s *DescribeVpcFirewallDetailResponseBodyLocalVpcVpcCidrTableList) GetRouteTableId() *string {
	return s.RouteTableId
}

func (s *DescribeVpcFirewallDetailResponseBodyLocalVpcVpcCidrTableList) SetRouteEntryList(v []*DescribeVpcFirewallDetailResponseBodyLocalVpcVpcCidrTableListRouteEntryList) *DescribeVpcFirewallDetailResponseBodyLocalVpcVpcCidrTableList {
	s.RouteEntryList = v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBodyLocalVpcVpcCidrTableList) SetRouteTableId(v string) *DescribeVpcFirewallDetailResponseBodyLocalVpcVpcCidrTableList {
	s.RouteTableId = &v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBodyLocalVpcVpcCidrTableList) Validate() error {
	if s.RouteEntryList != nil {
		for _, item := range s.RouteEntryList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVpcFirewallDetailResponseBodyLocalVpcVpcCidrTableListRouteEntryList struct {
	// The destination CIDR block of the local VPC.
	//
	// example:
	//
	// 192.168.XX.XX/24
	DestinationCidr *string `json:"DestinationCidr,omitempty" xml:"DestinationCidr,omitempty"`
	// The instance ID of the next hop for the local VPC.
	//
	// example:
	//
	// vrt-m5eb5me6c3l5sezae****
	NextHopInstanceId *string `json:"NextHopInstanceId,omitempty" xml:"NextHopInstanceId,omitempty"`
}

func (s DescribeVpcFirewallDetailResponseBodyLocalVpcVpcCidrTableListRouteEntryList) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallDetailResponseBodyLocalVpcVpcCidrTableListRouteEntryList) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallDetailResponseBodyLocalVpcVpcCidrTableListRouteEntryList) GetDestinationCidr() *string {
	return s.DestinationCidr
}

func (s *DescribeVpcFirewallDetailResponseBodyLocalVpcVpcCidrTableListRouteEntryList) GetNextHopInstanceId() *string {
	return s.NextHopInstanceId
}

func (s *DescribeVpcFirewallDetailResponseBodyLocalVpcVpcCidrTableListRouteEntryList) SetDestinationCidr(v string) *DescribeVpcFirewallDetailResponseBodyLocalVpcVpcCidrTableListRouteEntryList {
	s.DestinationCidr = &v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBodyLocalVpcVpcCidrTableListRouteEntryList) SetNextHopInstanceId(v string) *DescribeVpcFirewallDetailResponseBodyLocalVpcVpcCidrTableListRouteEntryList {
	s.NextHopInstanceId = &v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBodyLocalVpcVpcCidrTableListRouteEntryList) Validate() error {
	return dara.Validate(s)
}

type DescribeVpcFirewallDetailResponseBodyPeerVpc struct {
	// The ID of the ENI for the peer VPC.
	//
	// example:
	//
	// eni-8vbhfosfqv2rff42****
	EniId *string `json:"EniId,omitempty" xml:"EniId,omitempty"`
	// The private IP address of the ENI for the peer VPC.
	//
	// example:
	//
	// 192.168.XX.XX
	EniPrivateIpAddress *string `json:"EniPrivateIpAddress,omitempty" xml:"EniPrivateIpAddress,omitempty"`
	// The region ID of the peer VPC.
	//
	// example:
	//
	// cn-hangzhou
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The router interface ID of the peer VPC.
	//
	// example:
	//
	// vrt-m5eb5me6c3l5sezae****
	RouterInterfaceId *string `json:"RouterInterfaceId,omitempty" xml:"RouterInterfaceId,omitempty"`
	// The CIDR blocks of the peer VPC.
	VpcCidrTableList []*DescribeVpcFirewallDetailResponseBodyPeerVpcVpcCidrTableList `json:"VpcCidrTableList,omitempty" xml:"VpcCidrTableList,omitempty" type:"Repeated"`
	// The ID of the peer VPC.
	//
	// example:
	//
	// vpc-90rq0anm6t8vbwbo****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The name of the peer VPC.
	//
	// example:
	//
	// zcy_prod
	VpcName *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
}

func (s DescribeVpcFirewallDetailResponseBodyPeerVpc) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallDetailResponseBodyPeerVpc) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallDetailResponseBodyPeerVpc) GetEniId() *string {
	return s.EniId
}

func (s *DescribeVpcFirewallDetailResponseBodyPeerVpc) GetEniPrivateIpAddress() *string {
	return s.EniPrivateIpAddress
}

func (s *DescribeVpcFirewallDetailResponseBodyPeerVpc) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeVpcFirewallDetailResponseBodyPeerVpc) GetRouterInterfaceId() *string {
	return s.RouterInterfaceId
}

func (s *DescribeVpcFirewallDetailResponseBodyPeerVpc) GetVpcCidrTableList() []*DescribeVpcFirewallDetailResponseBodyPeerVpcVpcCidrTableList {
	return s.VpcCidrTableList
}

func (s *DescribeVpcFirewallDetailResponseBodyPeerVpc) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeVpcFirewallDetailResponseBodyPeerVpc) GetVpcName() *string {
	return s.VpcName
}

func (s *DescribeVpcFirewallDetailResponseBodyPeerVpc) SetEniId(v string) *DescribeVpcFirewallDetailResponseBodyPeerVpc {
	s.EniId = &v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBodyPeerVpc) SetEniPrivateIpAddress(v string) *DescribeVpcFirewallDetailResponseBodyPeerVpc {
	s.EniPrivateIpAddress = &v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBodyPeerVpc) SetRegionNo(v string) *DescribeVpcFirewallDetailResponseBodyPeerVpc {
	s.RegionNo = &v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBodyPeerVpc) SetRouterInterfaceId(v string) *DescribeVpcFirewallDetailResponseBodyPeerVpc {
	s.RouterInterfaceId = &v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBodyPeerVpc) SetVpcCidrTableList(v []*DescribeVpcFirewallDetailResponseBodyPeerVpcVpcCidrTableList) *DescribeVpcFirewallDetailResponseBodyPeerVpc {
	s.VpcCidrTableList = v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBodyPeerVpc) SetVpcId(v string) *DescribeVpcFirewallDetailResponseBodyPeerVpc {
	s.VpcId = &v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBodyPeerVpc) SetVpcName(v string) *DescribeVpcFirewallDetailResponseBodyPeerVpc {
	s.VpcName = &v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBodyPeerVpc) Validate() error {
	if s.VpcCidrTableList != nil {
		for _, item := range s.VpcCidrTableList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVpcFirewallDetailResponseBodyPeerVpcVpcCidrTableList struct {
	// The route entries of the peer VPC.
	RouteEntryList []*DescribeVpcFirewallDetailResponseBodyPeerVpcVpcCidrTableListRouteEntryList `json:"RouteEntryList,omitempty" xml:"RouteEntryList,omitempty" type:"Repeated"`
	// The ID of the route table for the peer VPC.
	//
	// example:
	//
	// vtb-1256
	RouteTableId *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
}

func (s DescribeVpcFirewallDetailResponseBodyPeerVpcVpcCidrTableList) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallDetailResponseBodyPeerVpcVpcCidrTableList) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallDetailResponseBodyPeerVpcVpcCidrTableList) GetRouteEntryList() []*DescribeVpcFirewallDetailResponseBodyPeerVpcVpcCidrTableListRouteEntryList {
	return s.RouteEntryList
}

func (s *DescribeVpcFirewallDetailResponseBodyPeerVpcVpcCidrTableList) GetRouteTableId() *string {
	return s.RouteTableId
}

func (s *DescribeVpcFirewallDetailResponseBodyPeerVpcVpcCidrTableList) SetRouteEntryList(v []*DescribeVpcFirewallDetailResponseBodyPeerVpcVpcCidrTableListRouteEntryList) *DescribeVpcFirewallDetailResponseBodyPeerVpcVpcCidrTableList {
	s.RouteEntryList = v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBodyPeerVpcVpcCidrTableList) SetRouteTableId(v string) *DescribeVpcFirewallDetailResponseBodyPeerVpcVpcCidrTableList {
	s.RouteTableId = &v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBodyPeerVpcVpcCidrTableList) Validate() error {
	if s.RouteEntryList != nil {
		for _, item := range s.RouteEntryList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVpcFirewallDetailResponseBodyPeerVpcVpcCidrTableListRouteEntryList struct {
	// The destination CIDR block of the peer VPC.
	//
	// example:
	//
	// 192.168.XX.XX/24
	DestinationCidr *string `json:"DestinationCidr,omitempty" xml:"DestinationCidr,omitempty"`
	// The instance ID of the next hop for the peer VPC.
	//
	// example:
	//
	// vrt-m5eb5me6c3l5sezae****
	NextHopInstanceId *string `json:"NextHopInstanceId,omitempty" xml:"NextHopInstanceId,omitempty"`
}

func (s DescribeVpcFirewallDetailResponseBodyPeerVpcVpcCidrTableListRouteEntryList) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallDetailResponseBodyPeerVpcVpcCidrTableListRouteEntryList) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallDetailResponseBodyPeerVpcVpcCidrTableListRouteEntryList) GetDestinationCidr() *string {
	return s.DestinationCidr
}

func (s *DescribeVpcFirewallDetailResponseBodyPeerVpcVpcCidrTableListRouteEntryList) GetNextHopInstanceId() *string {
	return s.NextHopInstanceId
}

func (s *DescribeVpcFirewallDetailResponseBodyPeerVpcVpcCidrTableListRouteEntryList) SetDestinationCidr(v string) *DescribeVpcFirewallDetailResponseBodyPeerVpcVpcCidrTableListRouteEntryList {
	s.DestinationCidr = &v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBodyPeerVpcVpcCidrTableListRouteEntryList) SetNextHopInstanceId(v string) *DescribeVpcFirewallDetailResponseBodyPeerVpcVpcCidrTableListRouteEntryList {
	s.NextHopInstanceId = &v
	return s
}

func (s *DescribeVpcFirewallDetailResponseBodyPeerVpcVpcCidrTableListRouteEntryList) Validate() error {
	return dara.Validate(s)
}
