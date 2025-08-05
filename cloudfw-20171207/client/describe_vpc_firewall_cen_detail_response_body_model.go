// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcFirewallCenDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConnectType(v string) *DescribeVpcFirewallCenDetailResponseBody
	GetConnectType() *string
	SetFirewallSwitchStatus(v string) *DescribeVpcFirewallCenDetailResponseBody
	GetFirewallSwitchStatus() *string
	SetFirewallVpc(v *DescribeVpcFirewallCenDetailResponseBodyFirewallVpc) *DescribeVpcFirewallCenDetailResponseBody
	GetFirewallVpc() *DescribeVpcFirewallCenDetailResponseBodyFirewallVpc
	SetLocalVpc(v *DescribeVpcFirewallCenDetailResponseBodyLocalVpc) *DescribeVpcFirewallCenDetailResponseBody
	GetLocalVpc() *DescribeVpcFirewallCenDetailResponseBodyLocalVpc
	SetRequestId(v string) *DescribeVpcFirewallCenDetailResponseBody
	GetRequestId() *string
	SetVpcFirewallId(v string) *DescribeVpcFirewallCenDetailResponseBody
	GetVpcFirewallId() *string
	SetVpcFirewallName(v string) *DescribeVpcFirewallCenDetailResponseBody
	GetVpcFirewallName() *string
}

type DescribeVpcFirewallCenDetailResponseBody struct {
	// The connection type of the VPC firewall. The value is fixed as **cen**, which indicates CEN instances.
	//
	// example:
	//
	// cen
	ConnectType *string `json:"ConnectType,omitempty" xml:"ConnectType,omitempty"`
	// The status of the VPC firewall. Valid values:
	//
	// 	- **opened**: enabled
	//
	// 	- **closed**: disabled
	//
	// 	- **notconfigured**: not configured
	//
	// example:
	//
	// opened
	FirewallSwitchStatus *string `json:"FirewallSwitchStatus,omitempty" xml:"FirewallSwitchStatus,omitempty"`
	// The firewall VPC.
	FirewallVpc *DescribeVpcFirewallCenDetailResponseBodyFirewallVpc `json:"FirewallVpc,omitempty" xml:"FirewallVpc,omitempty" type:"Struct"`
	// The details about the VPC.
	LocalVpc *DescribeVpcFirewallCenDetailResponseBodyLocalVpc `json:"LocalVpc,omitempty" xml:"LocalVpc,omitempty" type:"Struct"`
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
	// Test firewall
	VpcFirewallName *string `json:"VpcFirewallName,omitempty" xml:"VpcFirewallName,omitempty"`
}

func (s DescribeVpcFirewallCenDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallCenDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallCenDetailResponseBody) GetConnectType() *string {
	return s.ConnectType
}

func (s *DescribeVpcFirewallCenDetailResponseBody) GetFirewallSwitchStatus() *string {
	return s.FirewallSwitchStatus
}

func (s *DescribeVpcFirewallCenDetailResponseBody) GetFirewallVpc() *DescribeVpcFirewallCenDetailResponseBodyFirewallVpc {
	return s.FirewallVpc
}

func (s *DescribeVpcFirewallCenDetailResponseBody) GetLocalVpc() *DescribeVpcFirewallCenDetailResponseBodyLocalVpc {
	return s.LocalVpc
}

func (s *DescribeVpcFirewallCenDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVpcFirewallCenDetailResponseBody) GetVpcFirewallId() *string {
	return s.VpcFirewallId
}

func (s *DescribeVpcFirewallCenDetailResponseBody) GetVpcFirewallName() *string {
	return s.VpcFirewallName
}

func (s *DescribeVpcFirewallCenDetailResponseBody) SetConnectType(v string) *DescribeVpcFirewallCenDetailResponseBody {
	s.ConnectType = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBody) SetFirewallSwitchStatus(v string) *DescribeVpcFirewallCenDetailResponseBody {
	s.FirewallSwitchStatus = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBody) SetFirewallVpc(v *DescribeVpcFirewallCenDetailResponseBodyFirewallVpc) *DescribeVpcFirewallCenDetailResponseBody {
	s.FirewallVpc = v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBody) SetLocalVpc(v *DescribeVpcFirewallCenDetailResponseBodyLocalVpc) *DescribeVpcFirewallCenDetailResponseBody {
	s.LocalVpc = v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBody) SetRequestId(v string) *DescribeVpcFirewallCenDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBody) SetVpcFirewallId(v string) *DescribeVpcFirewallCenDetailResponseBody {
	s.VpcFirewallId = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBody) SetVpcFirewallName(v string) *DescribeVpcFirewallCenDetailResponseBody {
	s.VpcFirewallName = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeVpcFirewallCenDetailResponseBodyFirewallVpc struct {
	// Indicates whether you can specify a CIDR block when you create a VPC firewall for a Basic Edition transit router of a CEN instance. Valid values:
	//
	// 	- **1**: yes
	//
	// 	- **0**: no
	//
	// example:
	//
	// 0
	AllowConfiguration *int32 `json:"AllowConfiguration,omitempty" xml:"AllowConfiguration,omitempty"`
	// Firewall backup availability zone ID.
	//
	// example:
	//
	// cn-hangzhou-k
	StandbyZoneId *string `json:"StandbyZoneId,omitempty" xml:"StandbyZoneId,omitempty"`
	// The CIDR block of the VPC.
	//
	// example:
	//
	// 10.0.0.0/8
	VpcCidr *string `json:"VpcCidr,omitempty" xml:"VpcCidr,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-bp1245k5oagy2bp74****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The CIDR block of the vSwitch.
	//
	// example:
	//
	// 10.0.0.1/24
	VswitchCidr *string `json:"VswitchCidr,omitempty" xml:"VswitchCidr,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-bp1sqg9wms9wxcs1****
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	// The availability zone ID of the virtual switch.
	//
	// example:
	//
	// cn-hangzhou-i
	VswitchZoneId *string `json:"VswitchZoneId,omitempty" xml:"VswitchZoneId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-i
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeVpcFirewallCenDetailResponseBodyFirewallVpc) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallCenDetailResponseBodyFirewallVpc) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallCenDetailResponseBodyFirewallVpc) GetAllowConfiguration() *int32 {
	return s.AllowConfiguration
}

func (s *DescribeVpcFirewallCenDetailResponseBodyFirewallVpc) GetStandbyZoneId() *string {
	return s.StandbyZoneId
}

func (s *DescribeVpcFirewallCenDetailResponseBodyFirewallVpc) GetVpcCidr() *string {
	return s.VpcCidr
}

func (s *DescribeVpcFirewallCenDetailResponseBodyFirewallVpc) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeVpcFirewallCenDetailResponseBodyFirewallVpc) GetVswitchCidr() *string {
	return s.VswitchCidr
}

func (s *DescribeVpcFirewallCenDetailResponseBodyFirewallVpc) GetVswitchId() *string {
	return s.VswitchId
}

func (s *DescribeVpcFirewallCenDetailResponseBodyFirewallVpc) GetVswitchZoneId() *string {
	return s.VswitchZoneId
}

func (s *DescribeVpcFirewallCenDetailResponseBodyFirewallVpc) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeVpcFirewallCenDetailResponseBodyFirewallVpc) SetAllowConfiguration(v int32) *DescribeVpcFirewallCenDetailResponseBodyFirewallVpc {
	s.AllowConfiguration = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyFirewallVpc) SetStandbyZoneId(v string) *DescribeVpcFirewallCenDetailResponseBodyFirewallVpc {
	s.StandbyZoneId = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyFirewallVpc) SetVpcCidr(v string) *DescribeVpcFirewallCenDetailResponseBodyFirewallVpc {
	s.VpcCidr = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyFirewallVpc) SetVpcId(v string) *DescribeVpcFirewallCenDetailResponseBodyFirewallVpc {
	s.VpcId = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyFirewallVpc) SetVswitchCidr(v string) *DescribeVpcFirewallCenDetailResponseBodyFirewallVpc {
	s.VswitchCidr = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyFirewallVpc) SetVswitchId(v string) *DescribeVpcFirewallCenDetailResponseBodyFirewallVpc {
	s.VswitchId = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyFirewallVpc) SetVswitchZoneId(v string) *DescribeVpcFirewallCenDetailResponseBodyFirewallVpc {
	s.VswitchZoneId = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyFirewallVpc) SetZoneId(v string) *DescribeVpcFirewallCenDetailResponseBodyFirewallVpc {
	s.ZoneId = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyFirewallVpc) Validate() error {
	return dara.Validate(s)
}

type DescribeVpcFirewallCenDetailResponseBodyLocalVpc struct {
	// The ID of the connection between two network instances.
	//
	// example:
	//
	// tr-attach-sxig7bye51fid5****
	AttachmentId *string `json:"AttachmentId,omitempty" xml:"AttachmentId,omitempty"`
	// The name of the connection between two network instances.
	//
	// example:
	//
	// Local test
	AttachmentName *string `json:"AttachmentName,omitempty" xml:"AttachmentName,omitempty"`
	// An array consisting of the CIDR blocks that are protected by the VPC firewall.
	DefendCidrList []*string `json:"DefendCidrList,omitempty" xml:"DefendCidrList,omitempty" type:"Repeated"`
	// The Elastic Network Interfaces (ENIs).
	EniList []*DescribeVpcFirewallCenDetailResponseBodyLocalVpcEniList `json:"EniList,omitempty" xml:"EniList,omitempty" type:"Repeated"`
	// The ID of the specified vSwitch when the routing mode is manual.
	//
	// example:
	//
	// vsw-zeq4o875u****
	ManualVSwitchId *string `json:"ManualVSwitchId,omitempty" xml:"ManualVSwitchId,omitempty"`
	// The ID of the VPC for which the VPC firewall is created.
	//
	// example:
	//
	// vpc-2zefk9fbn8j7v585g****
	NetworkInstanceId *string `json:"NetworkInstanceId,omitempty" xml:"NetworkInstanceId,omitempty"`
	// The name of the network instance.
	//
	// example:
	//
	// Test VPC
	NetworkInstanceName *string `json:"NetworkInstanceName,omitempty" xml:"NetworkInstanceName,omitempty"`
	// The type of the network instance. The value is fixed as **VPC**.
	//
	// example:
	//
	// VPC
	NetworkInstanceType *string `json:"NetworkInstanceType,omitempty" xml:"NetworkInstanceType,omitempty"`
	// The UID of the Alibaba Cloud account to which the VPC belongs.
	//
	// example:
	//
	// 158039427902****
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region in which the VPC resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The routing mode. Valid values:
	//
	// 	- auto: automatic mode
	//
	// 	- manual: manual mode
	//
	// example:
	//
	// auto
	RouteMode *string `json:"RouteMode,omitempty" xml:"RouteMode,omitempty"`
	// Indicates whether the manual routing mode is supported. Valid values:
	//
	// 	- **1**: yes
	//
	// 	- **0**: no
	//
	// example:
	//
	// 0
	SupportManualMode *string `json:"SupportManualMode,omitempty" xml:"SupportManualMode,omitempty"`
	// The instance ID of the CEN transit router.
	//
	// example:
	//
	// tr-2zetwxskej633l3u1****
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
	// The edition of the CEN transit router. Valid values:
	//
	// 	- **Basic**: Basic Edition
	//
	// 	- **Enterprise**: Enterprise Edition
	//
	// example:
	//
	// Basic
	TransitRouterType *string `json:"TransitRouterType,omitempty" xml:"TransitRouterType,omitempty"`
	// An array that consists of the CIDR blocks of the VPC.
	VpcCidrTableList []*DescribeVpcFirewallCenDetailResponseBodyLocalVpcVpcCidrTableList `json:"VpcCidrTableList,omitempty" xml:"VpcCidrTableList,omitempty" type:"Repeated"`
	// The ID of the VPC.
	//
	// example:
	//
	// vpc-8vbwbo90rq0anm6t****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The name of the VPC.
	//
	// example:
	//
	// Test instance
	VpcName *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
}

func (s DescribeVpcFirewallCenDetailResponseBodyLocalVpc) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallCenDetailResponseBodyLocalVpc) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpc) GetAttachmentId() *string {
	return s.AttachmentId
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpc) GetAttachmentName() *string {
	return s.AttachmentName
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpc) GetDefendCidrList() []*string {
	return s.DefendCidrList
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpc) GetEniList() []*DescribeVpcFirewallCenDetailResponseBodyLocalVpcEniList {
	return s.EniList
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpc) GetManualVSwitchId() *string {
	return s.ManualVSwitchId
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpc) GetNetworkInstanceId() *string {
	return s.NetworkInstanceId
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpc) GetNetworkInstanceName() *string {
	return s.NetworkInstanceName
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpc) GetNetworkInstanceType() *string {
	return s.NetworkInstanceType
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpc) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpc) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpc) GetRouteMode() *string {
	return s.RouteMode
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpc) GetSupportManualMode() *string {
	return s.SupportManualMode
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpc) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpc) GetTransitRouterType() *string {
	return s.TransitRouterType
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpc) GetVpcCidrTableList() []*DescribeVpcFirewallCenDetailResponseBodyLocalVpcVpcCidrTableList {
	return s.VpcCidrTableList
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpc) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpc) GetVpcName() *string {
	return s.VpcName
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpc) SetAttachmentId(v string) *DescribeVpcFirewallCenDetailResponseBodyLocalVpc {
	s.AttachmentId = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpc) SetAttachmentName(v string) *DescribeVpcFirewallCenDetailResponseBodyLocalVpc {
	s.AttachmentName = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpc) SetDefendCidrList(v []*string) *DescribeVpcFirewallCenDetailResponseBodyLocalVpc {
	s.DefendCidrList = v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpc) SetEniList(v []*DescribeVpcFirewallCenDetailResponseBodyLocalVpcEniList) *DescribeVpcFirewallCenDetailResponseBodyLocalVpc {
	s.EniList = v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpc) SetManualVSwitchId(v string) *DescribeVpcFirewallCenDetailResponseBodyLocalVpc {
	s.ManualVSwitchId = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpc) SetNetworkInstanceId(v string) *DescribeVpcFirewallCenDetailResponseBodyLocalVpc {
	s.NetworkInstanceId = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpc) SetNetworkInstanceName(v string) *DescribeVpcFirewallCenDetailResponseBodyLocalVpc {
	s.NetworkInstanceName = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpc) SetNetworkInstanceType(v string) *DescribeVpcFirewallCenDetailResponseBodyLocalVpc {
	s.NetworkInstanceType = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpc) SetOwnerId(v string) *DescribeVpcFirewallCenDetailResponseBodyLocalVpc {
	s.OwnerId = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpc) SetRegionNo(v string) *DescribeVpcFirewallCenDetailResponseBodyLocalVpc {
	s.RegionNo = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpc) SetRouteMode(v string) *DescribeVpcFirewallCenDetailResponseBodyLocalVpc {
	s.RouteMode = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpc) SetSupportManualMode(v string) *DescribeVpcFirewallCenDetailResponseBodyLocalVpc {
	s.SupportManualMode = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpc) SetTransitRouterId(v string) *DescribeVpcFirewallCenDetailResponseBodyLocalVpc {
	s.TransitRouterId = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpc) SetTransitRouterType(v string) *DescribeVpcFirewallCenDetailResponseBodyLocalVpc {
	s.TransitRouterType = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpc) SetVpcCidrTableList(v []*DescribeVpcFirewallCenDetailResponseBodyLocalVpcVpcCidrTableList) *DescribeVpcFirewallCenDetailResponseBodyLocalVpc {
	s.VpcCidrTableList = v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpc) SetVpcId(v string) *DescribeVpcFirewallCenDetailResponseBodyLocalVpc {
	s.VpcId = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpc) SetVpcName(v string) *DescribeVpcFirewallCenDetailResponseBodyLocalVpc {
	s.VpcName = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpc) Validate() error {
	return dara.Validate(s)
}

type DescribeVpcFirewallCenDetailResponseBodyLocalVpcEniList struct {
	// The ID of the ENI that belongs to the VPC.
	//
	// example:
	//
	// eni-8vbhfosfqv2rff42****
	EniId *string `json:"EniId,omitempty" xml:"EniId,omitempty"`
	// The private IP address of the ENI that belongs to the VPC.
	//
	// example:
	//
	// 192.168.XX.XX
	EniPrivateIpAddress *string `json:"EniPrivateIpAddress,omitempty" xml:"EniPrivateIpAddress,omitempty"`
	// The ID of the vSwitch to which the ENI is connected.
	//
	// example:
	//
	// vsw-wz9viido7j436b0n1****
	EniVSwitchId *string `json:"EniVSwitchId,omitempty" xml:"EniVSwitchId,omitempty"`
}

func (s DescribeVpcFirewallCenDetailResponseBodyLocalVpcEniList) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallCenDetailResponseBodyLocalVpcEniList) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpcEniList) GetEniId() *string {
	return s.EniId
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpcEniList) GetEniPrivateIpAddress() *string {
	return s.EniPrivateIpAddress
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpcEniList) GetEniVSwitchId() *string {
	return s.EniVSwitchId
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpcEniList) SetEniId(v string) *DescribeVpcFirewallCenDetailResponseBodyLocalVpcEniList {
	s.EniId = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpcEniList) SetEniPrivateIpAddress(v string) *DescribeVpcFirewallCenDetailResponseBodyLocalVpcEniList {
	s.EniPrivateIpAddress = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpcEniList) SetEniVSwitchId(v string) *DescribeVpcFirewallCenDetailResponseBodyLocalVpcEniList {
	s.EniVSwitchId = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpcEniList) Validate() error {
	return dara.Validate(s)
}

type DescribeVpcFirewallCenDetailResponseBodyLocalVpcVpcCidrTableList struct {
	// The route entries for the VPC.
	RouteEntryList []*DescribeVpcFirewallCenDetailResponseBodyLocalVpcVpcCidrTableListRouteEntryList `json:"RouteEntryList,omitempty" xml:"RouteEntryList,omitempty" type:"Repeated"`
	// The route table ID of the VPC.
	//
	// example:
	//
	// vtb-1234
	RouteTableId *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
}

func (s DescribeVpcFirewallCenDetailResponseBodyLocalVpcVpcCidrTableList) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallCenDetailResponseBodyLocalVpcVpcCidrTableList) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpcVpcCidrTableList) GetRouteEntryList() []*DescribeVpcFirewallCenDetailResponseBodyLocalVpcVpcCidrTableListRouteEntryList {
	return s.RouteEntryList
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpcVpcCidrTableList) GetRouteTableId() *string {
	return s.RouteTableId
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpcVpcCidrTableList) SetRouteEntryList(v []*DescribeVpcFirewallCenDetailResponseBodyLocalVpcVpcCidrTableListRouteEntryList) *DescribeVpcFirewallCenDetailResponseBodyLocalVpcVpcCidrTableList {
	s.RouteEntryList = v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpcVpcCidrTableList) SetRouteTableId(v string) *DescribeVpcFirewallCenDetailResponseBodyLocalVpcVpcCidrTableList {
	s.RouteTableId = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpcVpcCidrTableList) Validate() error {
	return dara.Validate(s)
}

type DescribeVpcFirewallCenDetailResponseBodyLocalVpcVpcCidrTableListRouteEntryList struct {
	// The destination CIDR block of the VPC.
	//
	// example:
	//
	// 192.168.XX.XX/24
	DestinationCidr *string `json:"DestinationCidr,omitempty" xml:"DestinationCidr,omitempty"`
	// The instance ID of the next hop for the VPC.
	//
	// example:
	//
	// vrt-m5eb5me6c3l5sezae****
	NextHopInstanceId *string `json:"NextHopInstanceId,omitempty" xml:"NextHopInstanceId,omitempty"`
}

func (s DescribeVpcFirewallCenDetailResponseBodyLocalVpcVpcCidrTableListRouteEntryList) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallCenDetailResponseBodyLocalVpcVpcCidrTableListRouteEntryList) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpcVpcCidrTableListRouteEntryList) GetDestinationCidr() *string {
	return s.DestinationCidr
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpcVpcCidrTableListRouteEntryList) GetNextHopInstanceId() *string {
	return s.NextHopInstanceId
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpcVpcCidrTableListRouteEntryList) SetDestinationCidr(v string) *DescribeVpcFirewallCenDetailResponseBodyLocalVpcVpcCidrTableListRouteEntryList {
	s.DestinationCidr = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpcVpcCidrTableListRouteEntryList) SetNextHopInstanceId(v string) *DescribeVpcFirewallCenDetailResponseBodyLocalVpcVpcCidrTableListRouteEntryList {
	s.NextHopInstanceId = &v
	return s
}

func (s *DescribeVpcFirewallCenDetailResponseBodyLocalVpcVpcCidrTableListRouteEntryList) Validate() error {
	return dara.Validate(s)
}
