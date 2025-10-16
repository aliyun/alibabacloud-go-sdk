// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcFirewallListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeVpcFirewallListResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeVpcFirewallListResponseBody
	GetTotalCount() *int32
	SetVpcFirewalls(v []*DescribeVpcFirewallListResponseBodyVpcFirewalls) *DescribeVpcFirewallListResponseBody
	GetVpcFirewalls() []*DescribeVpcFirewallListResponseBodyVpcFirewalls
}

type DescribeVpcFirewallListResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 850A84D6-0DE4-4797-A1E8-00090125k8g2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of VPC firewalls.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The information about the VPC firewalls.
	VpcFirewalls []*DescribeVpcFirewallListResponseBodyVpcFirewalls `json:"VpcFirewalls,omitempty" xml:"VpcFirewalls,omitempty" type:"Repeated"`
}

func (s DescribeVpcFirewallListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVpcFirewallListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeVpcFirewallListResponseBody) GetVpcFirewalls() []*DescribeVpcFirewallListResponseBodyVpcFirewalls {
	return s.VpcFirewalls
}

func (s *DescribeVpcFirewallListResponseBody) SetRequestId(v string) *DescribeVpcFirewallListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBody) SetTotalCount(v int32) *DescribeVpcFirewallListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBody) SetVpcFirewalls(v []*DescribeVpcFirewallListResponseBodyVpcFirewalls) *DescribeVpcFirewallListResponseBody {
	s.VpcFirewalls = v
	return s
}

func (s *DescribeVpcFirewallListResponseBody) Validate() error {
	if s.VpcFirewalls != nil {
		for _, item := range s.VpcFirewalls {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVpcFirewallListResponseBodyVpcFirewalls struct {
	// ACL engine mode.
	AclConfig *DescribeVpcFirewallListResponseBodyVpcFirewallsAclConfig `json:"AclConfig,omitempty" xml:"AclConfig,omitempty" type:"Struct"`
	// The bandwidth of the Express Connect circuit. Unit: Mbit/s.
	//
	// example:
	//
	// 2
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
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
	// The connection type of the VPC firewall. The value is fixed as **expressconnect**, which indicates an Express Connect connection.
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
	// example:
	//
	// opened
	FirewallSwitchStatus *string `json:"FirewallSwitchStatus,omitempty" xml:"FirewallSwitchStatus,omitempty"`
	// The intrusion prevention system (IPS) configurations.
	IpsConfig *DescribeVpcFirewallListResponseBodyVpcFirewallsIpsConfig `json:"IpsConfig,omitempty" xml:"IpsConfig,omitempty" type:"Struct"`
	// The details about the local VPC.
	LocalVpc *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpc `json:"LocalVpc,omitempty" xml:"LocalVpc,omitempty" type:"Struct"`
	// The UID of the member that is managed by your Alibaba Cloud account.
	//
	// example:
	//
	// 258039427902****
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The details about the peer VPC.
	PeerVpc *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpc `json:"PeerVpc,omitempty" xml:"PeerVpc,omitempty" type:"Struct"`
	// Indicates whether you can create a VPC firewall in a specified region. Valid values:
	//
	// 	- **enable**: yes
	//
	// 	- **disable**: no
	//
	// example:
	//
	// enable
	RegionStatus *string `json:"RegionStatus,omitempty" xml:"RegionStatus,omitempty"`
	// The result code of the operation that creates the VPC firewall. Valid values:
	//
	// 	- **Unauthorized**: Cloud Firewall is not authorized to access a VPC for which the VPC firewall is created, and the VPC firewall cannot be created.
	//
	// 	- **RegionDisable**: VPC Firewall is not supported in the region of a VPC for which the VPC firewall is created, and the VPC firewall cannot be created.
	//
	// 	- **Empty string**: You can create a VPC firewall for the network instance.
	//
	// example:
	//
	// Unauthorized
	ResultCode *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
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

func (s DescribeVpcFirewallListResponseBodyVpcFirewalls) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallListResponseBodyVpcFirewalls) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewalls) GetAclConfig() *DescribeVpcFirewallListResponseBodyVpcFirewallsAclConfig {
	return s.AclConfig
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewalls) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewalls) GetConnectSubType() *string {
	return s.ConnectSubType
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewalls) GetConnectType() *string {
	return s.ConnectType
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewalls) GetFirewallSwitchStatus() *string {
	return s.FirewallSwitchStatus
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewalls) GetIpsConfig() *DescribeVpcFirewallListResponseBodyVpcFirewallsIpsConfig {
	return s.IpsConfig
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewalls) GetLocalVpc() *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpc {
	return s.LocalVpc
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewalls) GetMemberUid() *string {
	return s.MemberUid
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewalls) GetPeerVpc() *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpc {
	return s.PeerVpc
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewalls) GetRegionStatus() *string {
	return s.RegionStatus
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewalls) GetResultCode() *string {
	return s.ResultCode
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewalls) GetVpcFirewallId() *string {
	return s.VpcFirewallId
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewalls) GetVpcFirewallName() *string {
	return s.VpcFirewallName
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewalls) SetAclConfig(v *DescribeVpcFirewallListResponseBodyVpcFirewallsAclConfig) *DescribeVpcFirewallListResponseBodyVpcFirewalls {
	s.AclConfig = v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewalls) SetBandwidth(v int32) *DescribeVpcFirewallListResponseBodyVpcFirewalls {
	s.Bandwidth = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewalls) SetConnectSubType(v string) *DescribeVpcFirewallListResponseBodyVpcFirewalls {
	s.ConnectSubType = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewalls) SetConnectType(v string) *DescribeVpcFirewallListResponseBodyVpcFirewalls {
	s.ConnectType = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewalls) SetFirewallSwitchStatus(v string) *DescribeVpcFirewallListResponseBodyVpcFirewalls {
	s.FirewallSwitchStatus = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewalls) SetIpsConfig(v *DescribeVpcFirewallListResponseBodyVpcFirewallsIpsConfig) *DescribeVpcFirewallListResponseBodyVpcFirewalls {
	s.IpsConfig = v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewalls) SetLocalVpc(v *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpc) *DescribeVpcFirewallListResponseBodyVpcFirewalls {
	s.LocalVpc = v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewalls) SetMemberUid(v string) *DescribeVpcFirewallListResponseBodyVpcFirewalls {
	s.MemberUid = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewalls) SetPeerVpc(v *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpc) *DescribeVpcFirewallListResponseBodyVpcFirewalls {
	s.PeerVpc = v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewalls) SetRegionStatus(v string) *DescribeVpcFirewallListResponseBodyVpcFirewalls {
	s.RegionStatus = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewalls) SetResultCode(v string) *DescribeVpcFirewallListResponseBodyVpcFirewalls {
	s.ResultCode = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewalls) SetVpcFirewallId(v string) *DescribeVpcFirewallListResponseBodyVpcFirewalls {
	s.VpcFirewallId = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewalls) SetVpcFirewallName(v string) *DescribeVpcFirewallListResponseBodyVpcFirewalls {
	s.VpcFirewallName = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewalls) Validate() error {
	if s.AclConfig != nil {
		if err := s.AclConfig.Validate(); err != nil {
			return err
		}
	}
	if s.IpsConfig != nil {
		if err := s.IpsConfig.Validate(); err != nil {
			return err
		}
	}
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

type DescribeVpcFirewallListResponseBodyVpcFirewallsAclConfig struct {
	// Specifies whether to enable the strict mode. Valid values:
	//
	// 	- 1: yes
	//
	// 	- 0: no
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	StrictMode *int32 `json:"StrictMode,omitempty" xml:"StrictMode,omitempty"`
}

func (s DescribeVpcFirewallListResponseBodyVpcFirewallsAclConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallListResponseBodyVpcFirewallsAclConfig) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsAclConfig) GetStrictMode() *int32 {
	return s.StrictMode
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsAclConfig) SetStrictMode(v int32) *DescribeVpcFirewallListResponseBodyVpcFirewallsAclConfig {
	s.StrictMode = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsAclConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeVpcFirewallListResponseBodyVpcFirewallsIpsConfig struct {
	// Indicates whether basic protection is enabled. Valid values:
	//
	// 	- **1**: yes
	//
	// 	- **0**: no
	//
	// example:
	//
	// 1
	BasicRules *int32 `json:"BasicRules,omitempty" xml:"BasicRules,omitempty"`
	// Indicates whether virtual patching is enabled. Valid values:
	//
	// 	- **1**: yes
	//
	// 	- **0**: no
	//
	// example:
	//
	// 1
	EnableAllPatch *int32 `json:"EnableAllPatch,omitempty" xml:"EnableAllPatch,omitempty"`
	// The level of the rule group for the IPS. Valid values:
	//
	// 	- **1**: loose
	//
	// 	- **2**: medium
	//
	// 	- **3**: strict
	//
	// example:
	//
	// 1
	RuleClass *int32 `json:"RuleClass,omitempty" xml:"RuleClass,omitempty"`
	// The mode of the IPS. Valid values:
	//
	// 	- **1**: block mode
	//
	// 	- **0**: monitor mode
	//
	// example:
	//
	// 0
	RunMode *int32 `json:"RunMode,omitempty" xml:"RunMode,omitempty"`
}

func (s DescribeVpcFirewallListResponseBodyVpcFirewallsIpsConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallListResponseBodyVpcFirewallsIpsConfig) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsIpsConfig) GetBasicRules() *int32 {
	return s.BasicRules
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsIpsConfig) GetEnableAllPatch() *int32 {
	return s.EnableAllPatch
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsIpsConfig) GetRuleClass() *int32 {
	return s.RuleClass
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsIpsConfig) GetRunMode() *int32 {
	return s.RunMode
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsIpsConfig) SetBasicRules(v int32) *DescribeVpcFirewallListResponseBodyVpcFirewallsIpsConfig {
	s.BasicRules = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsIpsConfig) SetEnableAllPatch(v int32) *DescribeVpcFirewallListResponseBodyVpcFirewallsIpsConfig {
	s.EnableAllPatch = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsIpsConfig) SetRuleClass(v int32) *DescribeVpcFirewallListResponseBodyVpcFirewallsIpsConfig {
	s.RuleClass = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsIpsConfig) SetRunMode(v int32) *DescribeVpcFirewallListResponseBodyVpcFirewallsIpsConfig {
	s.RunMode = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsIpsConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpc struct {
	// Indicates whether Cloud Firewall is authorized to access the local VPC. The value is fixed as authorized, which indicates that Cloud Firewall is authorized to access the local VPC.
	//
	// example:
	//
	// authorized
	AuthorizationStatus *string `json:"AuthorizationStatus,omitempty" xml:"AuthorizationStatus,omitempty"`
	// The UID of the Alibaba Cloud account to which the local VPC belongs.
	//
	// example:
	//
	// 158039427902****
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the local VPC.
	//
	// example:
	//
	// cn-hangzhou
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// An array that consists of the CIDR blocks of the local VPC.
	VpcCidrTableList []*DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpcVpcCidrTableList `json:"VpcCidrTableList,omitempty" xml:"VpcCidrTableList,omitempty" type:"Repeated"`
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
	// Test instance
	VpcName *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
}

func (s DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpc) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpc) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpc) GetAuthorizationStatus() *string {
	return s.AuthorizationStatus
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpc) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpc) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpc) GetVpcCidrTableList() []*DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpcVpcCidrTableList {
	return s.VpcCidrTableList
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpc) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpc) GetVpcName() *string {
	return s.VpcName
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpc) SetAuthorizationStatus(v string) *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpc {
	s.AuthorizationStatus = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpc) SetOwnerId(v int64) *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpc {
	s.OwnerId = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpc) SetRegionNo(v string) *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpc {
	s.RegionNo = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpc) SetVpcCidrTableList(v []*DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpcVpcCidrTableList) *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpc {
	s.VpcCidrTableList = v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpc) SetVpcId(v string) *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpc {
	s.VpcId = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpc) SetVpcName(v string) *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpc {
	s.VpcName = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpc) Validate() error {
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

type DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpcVpcCidrTableList struct {
	// An array that consists of the route entries of the local VPC.
	RouteEntryList []*DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpcVpcCidrTableListRouteEntryList `json:"RouteEntryList,omitempty" xml:"RouteEntryList,omitempty" type:"Repeated"`
	// The ID of the route table for the local VPC.
	//
	// example:
	//
	// vtb-1234
	RouteTableId *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
}

func (s DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpcVpcCidrTableList) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpcVpcCidrTableList) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpcVpcCidrTableList) GetRouteEntryList() []*DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpcVpcCidrTableListRouteEntryList {
	return s.RouteEntryList
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpcVpcCidrTableList) GetRouteTableId() *string {
	return s.RouteTableId
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpcVpcCidrTableList) SetRouteEntryList(v []*DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpcVpcCidrTableListRouteEntryList) *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpcVpcCidrTableList {
	s.RouteEntryList = v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpcVpcCidrTableList) SetRouteTableId(v string) *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpcVpcCidrTableList {
	s.RouteTableId = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpcVpcCidrTableList) Validate() error {
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

type DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpcVpcCidrTableListRouteEntryList struct {
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

func (s DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpcVpcCidrTableListRouteEntryList) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpcVpcCidrTableListRouteEntryList) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpcVpcCidrTableListRouteEntryList) GetDestinationCidr() *string {
	return s.DestinationCidr
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpcVpcCidrTableListRouteEntryList) GetNextHopInstanceId() *string {
	return s.NextHopInstanceId
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpcVpcCidrTableListRouteEntryList) SetDestinationCidr(v string) *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpcVpcCidrTableListRouteEntryList {
	s.DestinationCidr = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpcVpcCidrTableListRouteEntryList) SetNextHopInstanceId(v string) *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpcVpcCidrTableListRouteEntryList {
	s.NextHopInstanceId = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsLocalVpcVpcCidrTableListRouteEntryList) Validate() error {
	return dara.Validate(s)
}

type DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpc struct {
	// Indicates whether Cloud Firewall is authorized to access the peer VPC. The value is fixed as **authorized**, which indicates that Cloud Firewall is authorized to access the peer VPC.
	//
	// example:
	//
	// authorized
	AuthorizationStatus *string `json:"AuthorizationStatus,omitempty" xml:"AuthorizationStatus,omitempty"`
	// The UID of the Alibaba Cloud account to which the peer VPC belongs.
	//
	// example:
	//
	// 158039427902****
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the peer VPC.
	//
	// example:
	//
	// cn-hangzhou
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// An array that consists of the CIDR blocks of the peer VPC.
	VpcCidrTableList []*DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpcVpcCidrTableList `json:"VpcCidrTableList,omitempty" xml:"VpcCidrTableList,omitempty" type:"Repeated"`
	// The ID of the peer VPC.
	//
	// example:
	//
	// vpc-8vbwbo90rq0anm6t****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The name of the peer VPC.
	//
	// example:
	//
	// Test VPC 2
	VpcName *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
}

func (s DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpc) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpc) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpc) GetAuthorizationStatus() *string {
	return s.AuthorizationStatus
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpc) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpc) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpc) GetVpcCidrTableList() []*DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpcVpcCidrTableList {
	return s.VpcCidrTableList
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpc) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpc) GetVpcName() *string {
	return s.VpcName
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpc) SetAuthorizationStatus(v string) *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpc {
	s.AuthorizationStatus = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpc) SetOwnerId(v int64) *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpc {
	s.OwnerId = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpc) SetRegionNo(v string) *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpc {
	s.RegionNo = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpc) SetVpcCidrTableList(v []*DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpcVpcCidrTableList) *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpc {
	s.VpcCidrTableList = v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpc) SetVpcId(v string) *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpc {
	s.VpcId = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpc) SetVpcName(v string) *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpc {
	s.VpcName = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpc) Validate() error {
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

type DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpcVpcCidrTableList struct {
	// An array that consists of the route entries of the peer VPC.
	RouteEntryList []*DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpcVpcCidrTableListRouteEntryList `json:"RouteEntryList,omitempty" xml:"RouteEntryList,omitempty" type:"Repeated"`
	// The ID of the route table for the peer VPC.
	//
	// example:
	//
	// vtb-1256
	RouteTableId *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
}

func (s DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpcVpcCidrTableList) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpcVpcCidrTableList) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpcVpcCidrTableList) GetRouteEntryList() []*DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpcVpcCidrTableListRouteEntryList {
	return s.RouteEntryList
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpcVpcCidrTableList) GetRouteTableId() *string {
	return s.RouteTableId
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpcVpcCidrTableList) SetRouteEntryList(v []*DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpcVpcCidrTableListRouteEntryList) *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpcVpcCidrTableList {
	s.RouteEntryList = v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpcVpcCidrTableList) SetRouteTableId(v string) *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpcVpcCidrTableList {
	s.RouteTableId = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpcVpcCidrTableList) Validate() error {
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

type DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpcVpcCidrTableListRouteEntryList struct {
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

func (s DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpcVpcCidrTableListRouteEntryList) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpcVpcCidrTableListRouteEntryList) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpcVpcCidrTableListRouteEntryList) GetDestinationCidr() *string {
	return s.DestinationCidr
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpcVpcCidrTableListRouteEntryList) GetNextHopInstanceId() *string {
	return s.NextHopInstanceId
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpcVpcCidrTableListRouteEntryList) SetDestinationCidr(v string) *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpcVpcCidrTableListRouteEntryList {
	s.DestinationCidr = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpcVpcCidrTableListRouteEntryList) SetNextHopInstanceId(v string) *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpcVpcCidrTableListRouteEntryList {
	s.NextHopInstanceId = &v
	return s
}

func (s *DescribeVpcFirewallListResponseBodyVpcFirewallsPeerVpcVpcCidrTableListRouteEntryList) Validate() error {
	return dara.Validate(s)
}
