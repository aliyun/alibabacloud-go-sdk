// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcFirewallCenListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeVpcFirewallCenListResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeVpcFirewallCenListResponseBody
	GetTotalCount() *int32
	SetVpcFirewalls(v []*DescribeVpcFirewallCenListResponseBodyVpcFirewalls) *DescribeVpcFirewallCenListResponseBody
	GetVpcFirewalls() []*DescribeVpcFirewallCenListResponseBodyVpcFirewalls
}

type DescribeVpcFirewallCenListResponseBody struct {
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
	// The details of the VPC firewalls.
	VpcFirewalls []*DescribeVpcFirewallCenListResponseBodyVpcFirewalls `json:"VpcFirewalls,omitempty" xml:"VpcFirewalls,omitempty" type:"Repeated"`
}

func (s DescribeVpcFirewallCenListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallCenListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallCenListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVpcFirewallCenListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeVpcFirewallCenListResponseBody) GetVpcFirewalls() []*DescribeVpcFirewallCenListResponseBodyVpcFirewalls {
	return s.VpcFirewalls
}

func (s *DescribeVpcFirewallCenListResponseBody) SetRequestId(v string) *DescribeVpcFirewallCenListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBody) SetTotalCount(v int32) *DescribeVpcFirewallCenListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBody) SetVpcFirewalls(v []*DescribeVpcFirewallCenListResponseBodyVpcFirewalls) *DescribeVpcFirewallCenListResponseBody {
	s.VpcFirewalls = v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBody) Validate() error {
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

type DescribeVpcFirewallCenListResponseBodyVpcFirewalls struct {
	// The access control list (ACL) engine mode.
	AclConfig *DescribeVpcFirewallCenListResponseBodyVpcFirewallsAclConfig `json:"AclConfig,omitempty" xml:"AclConfig,omitempty" type:"Struct"`
	// The ID of the CEN instance.
	//
	// example:
	//
	// cen-x5jayxou71ad73****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The name of the CEN instance.
	//
	// example:
	//
	// cen-test
	CenName *string `json:"CenName,omitempty" xml:"CenName,omitempty"`
	// The connection type of the VPC firewall. The value is fixed as cen, which indicates a CEN instance.
	//
	// example:
	//
	// cen
	ConnectType *string `json:"ConnectType,omitempty" xml:"ConnectType,omitempty"`
	// The status of the VPC firewall. Valid values:
	//
	// - **opened**: enabled.
	//
	// - **closed**: disabled.
	//
	// - **notconfigured**: not configured.
	//
	// example:
	//
	// opened
	FirewallSwitchStatus *string `json:"FirewallSwitchStatus,omitempty" xml:"FirewallSwitchStatus,omitempty"`
	// The intrusion prevention system (IPS) configuration.
	IpsConfig *DescribeVpcFirewallCenListResponseBodyVpcFirewallsIpsConfig `json:"IpsConfig,omitempty" xml:"IpsConfig,omitempty" type:"Struct"`
	// The details of the VPC.
	LocalVpc *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc `json:"LocalVpc,omitempty" xml:"LocalVpc,omitempty" type:"Struct"`
	// The UID of the member account.
	//
	// example:
	//
	// 258039427902****
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// Indicates whether the firewall can be automatically created. Cloud Firewall automatically creates a firewall and protects VPC traffic based on route learning. Valid values:
	//
	// - **passed**: The firewall can be automatically created.
	//
	// - **failed**: The firewall cannot be automatically created.
	//
	// - **unknown**: The status is unknown.
	//
	// example:
	//
	// failed
	PrecheckStatus *string `json:"PrecheckStatus,omitempty" xml:"PrecheckStatus,omitempty"`
	// The status of the region. Valid values:
	//
	// - **enable**: VPC firewalls can be configured in the region.
	//
	// - **disable**: VPC firewalls cannot be configured in the region.
	//
	// example:
	//
	// enable
	RegionStatus *string `json:"RegionStatus,omitempty" xml:"RegionStatus,omitempty"`
	// The result code of the operation to create the VPC firewall. Valid values:
	//
	// - **Unauthorized**: The VPC firewall cannot be created because a network instance is not authorized.
	//
	// - **RegionDisable**: The VPC firewall cannot be created because the region of the network instance is not supported.
	//
	// - **OpsDisable**: The VPC firewall cannot be created.
	//
	// - **VbrNotSupport**: The VPC firewall cannot be created because the VBR in the CEN instance is not supported.
	//
	// - An empty string indicates that the VPC firewall can be created for the network instance.
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
	// vpc-firewall-test
	VpcFirewallName *string `json:"VpcFirewallName,omitempty" xml:"VpcFirewallName,omitempty"`
}

func (s DescribeVpcFirewallCenListResponseBodyVpcFirewalls) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallCenListResponseBodyVpcFirewalls) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewalls) GetAclConfig() *DescribeVpcFirewallCenListResponseBodyVpcFirewallsAclConfig {
	return s.AclConfig
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewalls) GetCenId() *string {
	return s.CenId
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewalls) GetCenName() *string {
	return s.CenName
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewalls) GetConnectType() *string {
	return s.ConnectType
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewalls) GetFirewallSwitchStatus() *string {
	return s.FirewallSwitchStatus
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewalls) GetIpsConfig() *DescribeVpcFirewallCenListResponseBodyVpcFirewallsIpsConfig {
	return s.IpsConfig
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewalls) GetLocalVpc() *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc {
	return s.LocalVpc
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewalls) GetMemberUid() *string {
	return s.MemberUid
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewalls) GetPrecheckStatus() *string {
	return s.PrecheckStatus
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewalls) GetRegionStatus() *string {
	return s.RegionStatus
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewalls) GetResultCode() *string {
	return s.ResultCode
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewalls) GetVpcFirewallId() *string {
	return s.VpcFirewallId
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewalls) GetVpcFirewallName() *string {
	return s.VpcFirewallName
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewalls) SetAclConfig(v *DescribeVpcFirewallCenListResponseBodyVpcFirewallsAclConfig) *DescribeVpcFirewallCenListResponseBodyVpcFirewalls {
	s.AclConfig = v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewalls) SetCenId(v string) *DescribeVpcFirewallCenListResponseBodyVpcFirewalls {
	s.CenId = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewalls) SetCenName(v string) *DescribeVpcFirewallCenListResponseBodyVpcFirewalls {
	s.CenName = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewalls) SetConnectType(v string) *DescribeVpcFirewallCenListResponseBodyVpcFirewalls {
	s.ConnectType = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewalls) SetFirewallSwitchStatus(v string) *DescribeVpcFirewallCenListResponseBodyVpcFirewalls {
	s.FirewallSwitchStatus = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewalls) SetIpsConfig(v *DescribeVpcFirewallCenListResponseBodyVpcFirewallsIpsConfig) *DescribeVpcFirewallCenListResponseBodyVpcFirewalls {
	s.IpsConfig = v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewalls) SetLocalVpc(v *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc) *DescribeVpcFirewallCenListResponseBodyVpcFirewalls {
	s.LocalVpc = v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewalls) SetMemberUid(v string) *DescribeVpcFirewallCenListResponseBodyVpcFirewalls {
	s.MemberUid = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewalls) SetPrecheckStatus(v string) *DescribeVpcFirewallCenListResponseBodyVpcFirewalls {
	s.PrecheckStatus = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewalls) SetRegionStatus(v string) *DescribeVpcFirewallCenListResponseBodyVpcFirewalls {
	s.RegionStatus = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewalls) SetResultCode(v string) *DescribeVpcFirewallCenListResponseBodyVpcFirewalls {
	s.ResultCode = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewalls) SetVpcFirewallId(v string) *DescribeVpcFirewallCenListResponseBodyVpcFirewalls {
	s.VpcFirewallId = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewalls) SetVpcFirewallName(v string) *DescribeVpcFirewallCenListResponseBodyVpcFirewalls {
	s.VpcFirewallName = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewalls) Validate() error {
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
	return nil
}

type DescribeVpcFirewallCenListResponseBodyVpcFirewallsAclConfig struct {
	// Indicates whether the strict mode is enabled. Valid values:
	//
	// - 1: enabled
	//
	// - 0: disabled
	//
	// example:
	//
	// 1
	StrictMode *int32 `json:"StrictMode,omitempty" xml:"StrictMode,omitempty"`
}

func (s DescribeVpcFirewallCenListResponseBodyVpcFirewallsAclConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallCenListResponseBodyVpcFirewallsAclConfig) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsAclConfig) GetStrictMode() *int32 {
	return s.StrictMode
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsAclConfig) SetStrictMode(v int32) *DescribeVpcFirewallCenListResponseBodyVpcFirewallsAclConfig {
	s.StrictMode = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsAclConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeVpcFirewallCenListResponseBodyVpcFirewallsIpsConfig struct {
	// Indicates whether basic protection is enabled. Valid values:
	//
	// - **1**: enabled.
	//
	// - **0**: disabled.
	//
	// example:
	//
	// 1
	BasicRules *int32 `json:"BasicRules,omitempty" xml:"BasicRules,omitempty"`
	// Indicates whether virtual patching is enabled. Valid values:
	//
	// - **1**: enabled.
	//
	// - **0**: disabled.
	//
	// example:
	//
	// 1
	EnableAllPatch *int32 `json:"EnableAllPatch,omitempty" xml:"EnableAllPatch,omitempty"`
	// The type of the IPS rule group. Valid values:
	//
	// - **1**: loose.
	//
	// - **2**: medium.
	//
	// - **3**: strict.
	//
	// example:
	//
	// 1
	RuleClass *int32 `json:"RuleClass,omitempty" xml:"RuleClass,omitempty"`
	// The mode of the IPS. Valid values:
	//
	// - **1**: Block Mode.
	//
	// - **0**: Monitor mode.
	//
	// example:
	//
	// 0
	RunMode *int32 `json:"RunMode,omitempty" xml:"RunMode,omitempty"`
}

func (s DescribeVpcFirewallCenListResponseBodyVpcFirewallsIpsConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallCenListResponseBodyVpcFirewallsIpsConfig) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsIpsConfig) GetBasicRules() *int32 {
	return s.BasicRules
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsIpsConfig) GetEnableAllPatch() *int32 {
	return s.EnableAllPatch
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsIpsConfig) GetRuleClass() *int32 {
	return s.RuleClass
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsIpsConfig) GetRunMode() *int32 {
	return s.RunMode
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsIpsConfig) SetBasicRules(v int32) *DescribeVpcFirewallCenListResponseBodyVpcFirewallsIpsConfig {
	s.BasicRules = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsIpsConfig) SetEnableAllPatch(v int32) *DescribeVpcFirewallCenListResponseBodyVpcFirewallsIpsConfig {
	s.EnableAllPatch = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsIpsConfig) SetRuleClass(v int32) *DescribeVpcFirewallCenListResponseBodyVpcFirewallsIpsConfig {
	s.RuleClass = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsIpsConfig) SetRunMode(v int32) *DescribeVpcFirewallCenListResponseBodyVpcFirewallsIpsConfig {
	s.RunMode = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsIpsConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc struct {
	// The authorization status of the VPC. The value is fixed as **authorized**. This value indicates that the VPC is authorized.
	//
	// example:
	//
	// authorized
	AuthorizationStatus *string `json:"AuthorizationStatus,omitempty" xml:"AuthorizationStatus,omitempty"`
	// The CIDR blocks that are protected by the VPC firewall.
	DefendCidrList []*string `json:"DefendCidrList,omitempty" xml:"DefendCidrList,omitempty" type:"Repeated"`
	// The ID of the vSwitch that is specified for the manual routing mode.
	//
	// example:
	//
	// vsw-zeq4o875u****
	ManualVSwitchId *string `json:"ManualVSwitchId,omitempty" xml:"ManualVSwitchId,omitempty"`
	// The ID of the network instance.
	//
	// example:
	//
	// vpc-2zefk9fbn8j7v585g****
	NetworkInstanceId *string `json:"NetworkInstanceId,omitempty" xml:"NetworkInstanceId,omitempty"`
	// The name of the network instance.
	//
	// example:
	//
	// network-instance-test
	NetworkInstanceName *string `json:"NetworkInstanceName,omitempty" xml:"NetworkInstanceName,omitempty"`
	// The type of the network instance. Valid values:
	//
	// - **VPC**: Virtual Private Cloud.
	//
	// - **VBR**: Virtual Border Router.
	//
	// - **CCN**: Cloud Connect Network.
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
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the VPC.
	//
	// example:
	//
	// cn-hangzhou
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The routing mode. Valid values:
	//
	// - **auto**: automatic.
	//
	// - **manual**: manual.
	//
	// example:
	//
	// auto
	RouteMode *string `json:"RouteMode,omitempty" xml:"RouteMode,omitempty"`
	// Indicates whether the manual routing mode is supported. Valid values:
	//
	// - **1**: yes.
	//
	// - **0**: no.
	//
	// example:
	//
	// 0
	SupportManualMode *string `json:"SupportManualMode,omitempty" xml:"SupportManualMode,omitempty"`
	// The edition of the CEN transit router. Valid values:
	//
	// - **Basic**: Basic Edition.
	//
	// - **Enterprise**: Enterprise Edition.
	//
	// example:
	//
	// Basic
	TransitRouterType *string `json:"TransitRouterType,omitempty" xml:"TransitRouterType,omitempty"`
	// The CIDR blocks of the VPC.
	VpcCidrTableList []*DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpcVpcCidrTableList `json:"VpcCidrTableList,omitempty" xml:"VpcCidrTableList,omitempty" type:"Repeated"`
	// The instance ID of the VPC.
	//
	// example:
	//
	// vpc-8vbwbo90rq0anm6t****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The instance name of the VPC.
	//
	// example:
	//
	// vpc-instance
	VpcName *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
}

func (s DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc) GetAuthorizationStatus() *string {
	return s.AuthorizationStatus
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc) GetDefendCidrList() []*string {
	return s.DefendCidrList
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc) GetManualVSwitchId() *string {
	return s.ManualVSwitchId
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc) GetNetworkInstanceId() *string {
	return s.NetworkInstanceId
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc) GetNetworkInstanceName() *string {
	return s.NetworkInstanceName
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc) GetNetworkInstanceType() *string {
	return s.NetworkInstanceType
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc) GetRouteMode() *string {
	return s.RouteMode
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc) GetSupportManualMode() *string {
	return s.SupportManualMode
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc) GetTransitRouterType() *string {
	return s.TransitRouterType
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc) GetVpcCidrTableList() []*DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpcVpcCidrTableList {
	return s.VpcCidrTableList
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc) GetVpcName() *string {
	return s.VpcName
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc) SetAuthorizationStatus(v string) *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc {
	s.AuthorizationStatus = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc) SetDefendCidrList(v []*string) *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc {
	s.DefendCidrList = v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc) SetManualVSwitchId(v string) *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc {
	s.ManualVSwitchId = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc) SetNetworkInstanceId(v string) *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc {
	s.NetworkInstanceId = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc) SetNetworkInstanceName(v string) *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc {
	s.NetworkInstanceName = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc) SetNetworkInstanceType(v string) *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc {
	s.NetworkInstanceType = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc) SetOwnerId(v int64) *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc {
	s.OwnerId = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc) SetRegionNo(v string) *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc {
	s.RegionNo = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc) SetRouteMode(v string) *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc {
	s.RouteMode = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc) SetSupportManualMode(v string) *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc {
	s.SupportManualMode = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc) SetTransitRouterType(v string) *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc {
	s.TransitRouterType = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc) SetVpcCidrTableList(v []*DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpcVpcCidrTableList) *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc {
	s.VpcCidrTableList = v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc) SetVpcId(v string) *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc {
	s.VpcId = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc) SetVpcName(v string) *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc {
	s.VpcName = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpc) Validate() error {
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

type DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpcVpcCidrTableList struct {
	// The routes of the VPC.
	RouteEntryList []*DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpcVpcCidrTableListRouteEntryList `json:"RouteEntryList,omitempty" xml:"RouteEntryList,omitempty" type:"Repeated"`
	// The ID of the route table for the VPC.
	//
	// example:
	//
	// vtb-1234
	RouteTableId *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
}

func (s DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpcVpcCidrTableList) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpcVpcCidrTableList) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpcVpcCidrTableList) GetRouteEntryList() []*DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpcVpcCidrTableListRouteEntryList {
	return s.RouteEntryList
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpcVpcCidrTableList) GetRouteTableId() *string {
	return s.RouteTableId
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpcVpcCidrTableList) SetRouteEntryList(v []*DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpcVpcCidrTableListRouteEntryList) *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpcVpcCidrTableList {
	s.RouteEntryList = v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpcVpcCidrTableList) SetRouteTableId(v string) *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpcVpcCidrTableList {
	s.RouteTableId = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpcVpcCidrTableList) Validate() error {
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

type DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpcVpcCidrTableListRouteEntryList struct {
	// The destination CIDR block of the VPC.
	//
	// example:
	//
	// 192.168.XX.XX/24
	DestinationCidr *string `json:"DestinationCidr,omitempty" xml:"DestinationCidr,omitempty"`
	// The ID of the next hop instance in the VPC.
	//
	// example:
	//
	// vrt-m5eb5me6c3l5sezae****
	NextHopInstanceId *string `json:"NextHopInstanceId,omitempty" xml:"NextHopInstanceId,omitempty"`
}

func (s DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpcVpcCidrTableListRouteEntryList) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpcVpcCidrTableListRouteEntryList) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpcVpcCidrTableListRouteEntryList) GetDestinationCidr() *string {
	return s.DestinationCidr
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpcVpcCidrTableListRouteEntryList) GetNextHopInstanceId() *string {
	return s.NextHopInstanceId
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpcVpcCidrTableListRouteEntryList) SetDestinationCidr(v string) *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpcVpcCidrTableListRouteEntryList {
	s.DestinationCidr = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpcVpcCidrTableListRouteEntryList) SetNextHopInstanceId(v string) *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpcVpcCidrTableListRouteEntryList {
	s.NextHopInstanceId = &v
	return s
}

func (s *DescribeVpcFirewallCenListResponseBodyVpcFirewallsLocalVpcVpcCidrTableListRouteEntryList) Validate() error {
	return dara.Validate(s)
}
