// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSecurityProxyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBasicDefenseEnabled(v string) *CreateSecurityProxyRequest
	GetBasicDefenseEnabled() *string
	SetFirewallServiceMode(v string) *CreateSecurityProxyRequest
	GetFirewallServiceMode() *string
	SetFirewallServiceZones(v []*string) *CreateSecurityProxyRequest
	GetFirewallServiceZones() []*string
	SetFirewallSwitch(v string) *CreateSecurityProxyRequest
	GetFirewallSwitch() *string
	SetFwVswitchZoneId(v string) *CreateSecurityProxyRequest
	GetFwVswitchZoneId() *string
	SetLang(v string) *CreateSecurityProxyRequest
	GetLang() *string
	SetMasterVSwitchCidr(v string) *CreateSecurityProxyRequest
	GetMasterVSwitchCidr() *string
	SetMasterZoneId(v string) *CreateSecurityProxyRequest
	GetMasterZoneId() *string
	SetNatGatewayId(v string) *CreateSecurityProxyRequest
	GetNatGatewayId() *string
	SetNatMode(v string) *CreateSecurityProxyRequest
	GetNatMode() *string
	SetNatRouteEntryList(v []*CreateSecurityProxyRequestNatRouteEntryList) *CreateSecurityProxyRequest
	GetNatRouteEntryList() []*CreateSecurityProxyRequestNatRouteEntryList
	SetProxyName(v string) *CreateSecurityProxyRequest
	GetProxyName() *string
	SetRegionNo(v string) *CreateSecurityProxyRequest
	GetRegionNo() *string
	SetRuleClass(v string) *CreateSecurityProxyRequest
	GetRuleClass() *string
	SetRunMode(v string) *CreateSecurityProxyRequest
	GetRunMode() *string
	SetStandbyVSwitchCidr(v string) *CreateSecurityProxyRequest
	GetStandbyVSwitchCidr() *string
	SetStandbyZoneId(v string) *CreateSecurityProxyRequest
	GetStandbyZoneId() *string
	SetStrictMode(v int32) *CreateSecurityProxyRequest
	GetStrictMode() *int32
	SetThreatIntelEnabled(v string) *CreateSecurityProxyRequest
	GetThreatIntelEnabled() *string
	SetVirtualPatchEnabled(v string) *CreateSecurityProxyRequest
	GetVirtualPatchEnabled() *string
	SetVpcId(v string) *CreateSecurityProxyRequest
	GetVpcId() *string
	SetVswitchAuto(v string) *CreateSecurityProxyRequest
	GetVswitchAuto() *string
	SetVswitchCidr(v string) *CreateSecurityProxyRequest
	GetVswitchCidr() *string
	SetVswitchId(v string) *CreateSecurityProxyRequest
	GetVswitchId() *string
}

type CreateSecurityProxyRequest struct {
	// Specifies whether to enable the basic protection policies of the IPS. If this parameter is not specified, the system default (enabled) is used. Valid values:
	//
	// - **0**: disabled.
	//
	// - **1**: enabled.
	//
	// example:
	//
	// 1
	BasicDefenseEnabled *string `json:"BasicDefenseEnabled,omitempty" xml:"BasicDefenseEnabled,omitempty"`
	// The deployment mode of the firewall service. Valid values:
	//
	// - **PrimaryStandby**: primary/secondary mode.
	//
	// - **MultiPrimary**: active-active mode.
	//
	// example:
	//
	// PrimaryStandby
	FirewallServiceMode *string `json:"FirewallServiceMode,omitempty" xml:"FirewallServiceMode,omitempty"`
	// The list of zone IDs used by the firewall service.
	FirewallServiceZones []*string `json:"FirewallServiceZones,omitempty" xml:"FirewallServiceZones,omitempty" type:"Repeated"`
	// The security protection switch. Valid values:
	//
	// - **open**: enabled.
	//
	// - **close**: disabled.
	//
	// example:
	//
	// close
	FirewallSwitch *string `json:"FirewallSwitch,omitempty" xml:"FirewallSwitch,omitempty"`
	// The zone of the firewall vSwitch.
	//
	// example:
	//
	// cn-beijing-b
	FwVswitchZoneId *string `json:"FwVswitchZoneId,omitempty" xml:"FwVswitchZoneId,omitempty"`
	// The language of the response message. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The vSwitch CIDR block of the GWLB endpoint (GWLBe) in the active zone. This parameter applies only to NAT firewalls that use the GWLB architecture. For the GWLB architecture, this parameter is required along with MasterZoneId, StandbyZoneId, and StandbyVSwitchCidr. If any of these parameters is missing, ErrorParameters is returned. The vSwitch is created by the firewall creation node based on this vSwitch CIDR block. This parameter is not read for the ENI architecture.
	//
	// example:
	//
	// 192.168.0.0/24
	MasterVSwitchCidr *string `json:"MasterVSwitchCidr,omitempty" xml:"MasterVSwitchCidr,omitempty"`
	// The ID of the primary zone. This parameter applies only to NAT firewalls that use the GWLB architecture. For the GWLB architecture, this parameter is required along with MasterVSwitchCidr, StandbyZoneId, and StandbyVSwitchCidr. If any of these parameters is missing, ErrorParameters is returned. The value must be the same as the zone where the NAT gateway resides. This parameter is not read for the ENI architecture.
	//
	// example:
	//
	// cn-hangzhou-h
	MasterZoneId *string `json:"MasterZoneId,omitempty" xml:"MasterZoneId,omitempty"`
	// The ID of the NAT gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// ngw-bp1okz6k7******
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The architecture type of the NAT firewall. Valid values: **eni*	- (traditional ENI architecture) and **gwlb*	- (new GWLB architecture). If this parameter is omitted or left empty, the firewall is created by using the ENI architecture, and no implicit selection is performed on the server side. You can obtain the recommended architecture from the response of the DescribeNatFirewallArchitecture operation and pass it back when you create the firewall.
	//
	// example:
	//
	// gwlb
	NatMode *string `json:"NatMode,omitempty" xml:"NatMode,omitempty"`
	// The list of routes to be switched for the NAT gateway.
	NatRouteEntryList []*CreateSecurityProxyRequestNatRouteEntryList `json:"NatRouteEntryList,omitempty" xml:"NatRouteEntryList,omitempty" type:"Repeated"`
	// The name of the NAT firewall. The name must be 4 to 50 characters in length and can contain uppercase and lowercase letters, Chinese characters, digits, and underscores (_). It cannot start with an underscore.
	//
	// This parameter is required.
	//
	// example:
	//
	// nat-firewall
	ProxyName *string `json:"ProxyName,omitempty" xml:"ProxyName,omitempty"`
	// The region ID of the VPC.
	//
	// > For more information about the regions supported by Cloud Firewall, see [Supported regions](https://help.aliyun.com/document_detail/195657.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The level of the IPS rule group. This parameter is valid only in block mode. If this parameter is not specified, the system default (medium) is used. Valid values:
	//
	// - **0**: none. The value is forced to 0 in monitor mode.
	//
	// - **1**: loose.
	//
	// - **2**: medium.
	//
	// - **3**: strict.
	//
	// example:
	//
	// 2
	RuleClass *string `json:"RuleClass,omitempty" xml:"RuleClass,omitempty"`
	// The IPS defense mode. This parameter is supported only by NAT firewalls that use the GWLB architecture. If this parameter is not specified, the system default (Block Mode) is used. Valid values:
	//
	// - **0**: monitor mode. Only alerting is triggered, and traffic is not blocked.
	//
	// - **1**: Block Mode.
	//
	// example:
	//
	// 1
	RunMode *string `json:"RunMode,omitempty" xml:"RunMode,omitempty"`
	// The vSwitch CIDR block of the GWLBe in the standby zone. This parameter applies only to NAT firewalls that use the GWLB architecture. For the GWLB architecture, this parameter is required along with MasterZoneId, MasterVSwitchCidr, and StandbyZoneId. If any of these parameters is missing, ErrorParameters is returned. The vSwitch is created by the firewall creation node based on this vSwitch CIDR block. This parameter is not read for the ENI architecture.
	//
	// example:
	//
	// 192.168.1.0/24
	StandbyVSwitchCidr *string `json:"StandbyVSwitchCidr,omitempty" xml:"StandbyVSwitchCidr,omitempty"`
	// The ID of the secondary zone. This parameter applies only to NAT firewalls that use the GWLB architecture. For the GWLB architecture, this parameter is required along with MasterZoneId, MasterVSwitchCidr, and StandbyVSwitchCidr. If any of these parameters is missing, ErrorParameters is returned. The value cannot be the same as MasterZoneId. If they are the same, ErrorZoneIdInvalid is returned. This parameter is not read for the ENI architecture.
	//
	// example:
	//
	// cn-hangzhou-i
	StandbyZoneId *string `json:"StandbyZoneId,omitempty" xml:"StandbyZoneId,omitempty"`
	// Specifies whether to enable strict mode. Valid values:
	//
	// - **1**: enables strict mode.
	//
	// - **0**: disables strict mode.
	//
	// example:
	//
	// 0
	StrictMode *int32 `json:"StrictMode,omitempty" xml:"StrictMode,omitempty"`
	// Specifies whether to enable the threat intelligence rules of the IPS. If this parameter is not specified, the system default (enabled) is used. Valid values:
	//
	// - **0**: disabled.
	//
	// - **1**: enabled.
	//
	// example:
	//
	// 1
	ThreatIntelEnabled *string `json:"ThreatIntelEnabled,omitempty" xml:"ThreatIntelEnabled,omitempty"`
	// Specifies whether to enable the Virtual Patches of the IPS. If this parameter is not specified, the system default (enabled) is used. Valid values:
	//
	// - **0**: disabled.
	//
	// - **1**: enabled.
	//
	// example:
	//
	// 1
	VirtualPatchEnabled *string `json:"VirtualPatchEnabled,omitempty" xml:"VirtualPatchEnabled,omitempty"`
	// The instance ID of the VPC.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-uf6b5lyul0x******
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// Specifies whether to use the automatic mode for the vSwitch. Valid values:
	//
	// - **true**: automatic mode.
	//
	// - **false**: manual mode.
	//
	// > Default value: **true**. If VswitchAuto is set to true, VswitchCidr is required and must be a valid CIDR block. If VswitchAuto is set to false, VswitchId is required.
	//
	// example:
	//
	// true
	VswitchAuto *string `json:"VswitchAuto,omitempty" xml:"VswitchAuto,omitempty"`
	// The CIDR block of the vSwitch. This parameter is required when the automatic mode is used for the vSwitch.
	//
	// example:
	//
	// 0.0.0.0/0
	VswitchCidr *string `json:"VswitchCidr,omitempty" xml:"VswitchCidr,omitempty"`
	// The ID of the vSwitch. This parameter is required when the manual mode is used for the vSwitch.
	//
	// example:
	//
	// vsw-bp1sqg9w******
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
}

func (s CreateSecurityProxyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSecurityProxyRequest) GoString() string {
	return s.String()
}

func (s *CreateSecurityProxyRequest) GetBasicDefenseEnabled() *string {
	return s.BasicDefenseEnabled
}

func (s *CreateSecurityProxyRequest) GetFirewallServiceMode() *string {
	return s.FirewallServiceMode
}

func (s *CreateSecurityProxyRequest) GetFirewallServiceZones() []*string {
	return s.FirewallServiceZones
}

func (s *CreateSecurityProxyRequest) GetFirewallSwitch() *string {
	return s.FirewallSwitch
}

func (s *CreateSecurityProxyRequest) GetFwVswitchZoneId() *string {
	return s.FwVswitchZoneId
}

func (s *CreateSecurityProxyRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateSecurityProxyRequest) GetMasterVSwitchCidr() *string {
	return s.MasterVSwitchCidr
}

func (s *CreateSecurityProxyRequest) GetMasterZoneId() *string {
	return s.MasterZoneId
}

func (s *CreateSecurityProxyRequest) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *CreateSecurityProxyRequest) GetNatMode() *string {
	return s.NatMode
}

func (s *CreateSecurityProxyRequest) GetNatRouteEntryList() []*CreateSecurityProxyRequestNatRouteEntryList {
	return s.NatRouteEntryList
}

func (s *CreateSecurityProxyRequest) GetProxyName() *string {
	return s.ProxyName
}

func (s *CreateSecurityProxyRequest) GetRegionNo() *string {
	return s.RegionNo
}

func (s *CreateSecurityProxyRequest) GetRuleClass() *string {
	return s.RuleClass
}

func (s *CreateSecurityProxyRequest) GetRunMode() *string {
	return s.RunMode
}

func (s *CreateSecurityProxyRequest) GetStandbyVSwitchCidr() *string {
	return s.StandbyVSwitchCidr
}

func (s *CreateSecurityProxyRequest) GetStandbyZoneId() *string {
	return s.StandbyZoneId
}

func (s *CreateSecurityProxyRequest) GetStrictMode() *int32 {
	return s.StrictMode
}

func (s *CreateSecurityProxyRequest) GetThreatIntelEnabled() *string {
	return s.ThreatIntelEnabled
}

func (s *CreateSecurityProxyRequest) GetVirtualPatchEnabled() *string {
	return s.VirtualPatchEnabled
}

func (s *CreateSecurityProxyRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateSecurityProxyRequest) GetVswitchAuto() *string {
	return s.VswitchAuto
}

func (s *CreateSecurityProxyRequest) GetVswitchCidr() *string {
	return s.VswitchCidr
}

func (s *CreateSecurityProxyRequest) GetVswitchId() *string {
	return s.VswitchId
}

func (s *CreateSecurityProxyRequest) SetBasicDefenseEnabled(v string) *CreateSecurityProxyRequest {
	s.BasicDefenseEnabled = &v
	return s
}

func (s *CreateSecurityProxyRequest) SetFirewallServiceMode(v string) *CreateSecurityProxyRequest {
	s.FirewallServiceMode = &v
	return s
}

func (s *CreateSecurityProxyRequest) SetFirewallServiceZones(v []*string) *CreateSecurityProxyRequest {
	s.FirewallServiceZones = v
	return s
}

func (s *CreateSecurityProxyRequest) SetFirewallSwitch(v string) *CreateSecurityProxyRequest {
	s.FirewallSwitch = &v
	return s
}

func (s *CreateSecurityProxyRequest) SetFwVswitchZoneId(v string) *CreateSecurityProxyRequest {
	s.FwVswitchZoneId = &v
	return s
}

func (s *CreateSecurityProxyRequest) SetLang(v string) *CreateSecurityProxyRequest {
	s.Lang = &v
	return s
}

func (s *CreateSecurityProxyRequest) SetMasterVSwitchCidr(v string) *CreateSecurityProxyRequest {
	s.MasterVSwitchCidr = &v
	return s
}

func (s *CreateSecurityProxyRequest) SetMasterZoneId(v string) *CreateSecurityProxyRequest {
	s.MasterZoneId = &v
	return s
}

func (s *CreateSecurityProxyRequest) SetNatGatewayId(v string) *CreateSecurityProxyRequest {
	s.NatGatewayId = &v
	return s
}

func (s *CreateSecurityProxyRequest) SetNatMode(v string) *CreateSecurityProxyRequest {
	s.NatMode = &v
	return s
}

func (s *CreateSecurityProxyRequest) SetNatRouteEntryList(v []*CreateSecurityProxyRequestNatRouteEntryList) *CreateSecurityProxyRequest {
	s.NatRouteEntryList = v
	return s
}

func (s *CreateSecurityProxyRequest) SetProxyName(v string) *CreateSecurityProxyRequest {
	s.ProxyName = &v
	return s
}

func (s *CreateSecurityProxyRequest) SetRegionNo(v string) *CreateSecurityProxyRequest {
	s.RegionNo = &v
	return s
}

func (s *CreateSecurityProxyRequest) SetRuleClass(v string) *CreateSecurityProxyRequest {
	s.RuleClass = &v
	return s
}

func (s *CreateSecurityProxyRequest) SetRunMode(v string) *CreateSecurityProxyRequest {
	s.RunMode = &v
	return s
}

func (s *CreateSecurityProxyRequest) SetStandbyVSwitchCidr(v string) *CreateSecurityProxyRequest {
	s.StandbyVSwitchCidr = &v
	return s
}

func (s *CreateSecurityProxyRequest) SetStandbyZoneId(v string) *CreateSecurityProxyRequest {
	s.StandbyZoneId = &v
	return s
}

func (s *CreateSecurityProxyRequest) SetStrictMode(v int32) *CreateSecurityProxyRequest {
	s.StrictMode = &v
	return s
}

func (s *CreateSecurityProxyRequest) SetThreatIntelEnabled(v string) *CreateSecurityProxyRequest {
	s.ThreatIntelEnabled = &v
	return s
}

func (s *CreateSecurityProxyRequest) SetVirtualPatchEnabled(v string) *CreateSecurityProxyRequest {
	s.VirtualPatchEnabled = &v
	return s
}

func (s *CreateSecurityProxyRequest) SetVpcId(v string) *CreateSecurityProxyRequest {
	s.VpcId = &v
	return s
}

func (s *CreateSecurityProxyRequest) SetVswitchAuto(v string) *CreateSecurityProxyRequest {
	s.VswitchAuto = &v
	return s
}

func (s *CreateSecurityProxyRequest) SetVswitchCidr(v string) *CreateSecurityProxyRequest {
	s.VswitchCidr = &v
	return s
}

func (s *CreateSecurityProxyRequest) SetVswitchId(v string) *CreateSecurityProxyRequest {
	s.VswitchId = &v
	return s
}

func (s *CreateSecurityProxyRequest) Validate() error {
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

type CreateSecurityProxyRequestNatRouteEntryList struct {
	// The destination CIDR block of the default route.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0.0.0.0/0
	DestinationCidr *string `json:"DestinationCidr,omitempty" xml:"DestinationCidr,omitempty"`
	// The next hop ID of the original NAT gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// ngw-bp1okz6******
	NextHopId *string `json:"NextHopId,omitempty" xml:"NextHopId,omitempty"`
	// The network type of the next hop. Valid values: NatGateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// NatGateway
	NextHopType *string `json:"NextHopType,omitempty" xml:"NextHopType,omitempty"`
	// The ID of the route table to which the default route of the NAT gateway belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// vtb-2ze1******
	RouteTableId *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
}

func (s CreateSecurityProxyRequestNatRouteEntryList) String() string {
	return dara.Prettify(s)
}

func (s CreateSecurityProxyRequestNatRouteEntryList) GoString() string {
	return s.String()
}

func (s *CreateSecurityProxyRequestNatRouteEntryList) GetDestinationCidr() *string {
	return s.DestinationCidr
}

func (s *CreateSecurityProxyRequestNatRouteEntryList) GetNextHopId() *string {
	return s.NextHopId
}

func (s *CreateSecurityProxyRequestNatRouteEntryList) GetNextHopType() *string {
	return s.NextHopType
}

func (s *CreateSecurityProxyRequestNatRouteEntryList) GetRouteTableId() *string {
	return s.RouteTableId
}

func (s *CreateSecurityProxyRequestNatRouteEntryList) SetDestinationCidr(v string) *CreateSecurityProxyRequestNatRouteEntryList {
	s.DestinationCidr = &v
	return s
}

func (s *CreateSecurityProxyRequestNatRouteEntryList) SetNextHopId(v string) *CreateSecurityProxyRequestNatRouteEntryList {
	s.NextHopId = &v
	return s
}

func (s *CreateSecurityProxyRequestNatRouteEntryList) SetNextHopType(v string) *CreateSecurityProxyRequestNatRouteEntryList {
	s.NextHopType = &v
	return s
}

func (s *CreateSecurityProxyRequestNatRouteEntryList) SetRouteTableId(v string) *CreateSecurityProxyRequestNatRouteEntryList {
	s.RouteTableId = &v
	return s
}

func (s *CreateSecurityProxyRequestNatRouteEntryList) Validate() error {
	return dara.Validate(s)
}
