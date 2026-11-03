// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVpcFirewallCenConfigureRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCenId(v string) *CreateVpcFirewallCenConfigureRequest
	GetCenId() *string
	SetFirewallSwitch(v string) *CreateVpcFirewallCenConfigureRequest
	GetFirewallSwitch() *string
	SetFirewallVSwitchCidrBlock(v string) *CreateVpcFirewallCenConfigureRequest
	GetFirewallVSwitchCidrBlock() *string
	SetFirewallVSwitchZoneId(v string) *CreateVpcFirewallCenConfigureRequest
	GetFirewallVSwitchZoneId() *string
	SetFirewallVpcCidrBlock(v string) *CreateVpcFirewallCenConfigureRequest
	GetFirewallVpcCidrBlock() *string
	SetFirewallVpcStandbyZoneId(v string) *CreateVpcFirewallCenConfigureRequest
	GetFirewallVpcStandbyZoneId() *string
	SetFirewallVpcZoneId(v string) *CreateVpcFirewallCenConfigureRequest
	GetFirewallVpcZoneId() *string
	SetLang(v string) *CreateVpcFirewallCenConfigureRequest
	GetLang() *string
	SetMemberUid(v string) *CreateVpcFirewallCenConfigureRequest
	GetMemberUid() *string
	SetNetworkInstanceId(v string) *CreateVpcFirewallCenConfigureRequest
	GetNetworkInstanceId() *string
	SetVSwitchId(v string) *CreateVpcFirewallCenConfigureRequest
	GetVSwitchId() *string
	SetVpcFirewallName(v string) *CreateVpcFirewallCenConfigureRequest
	GetVpcFirewallName() *string
	SetVpcRegion(v string) *CreateVpcFirewallCenConfigureRequest
	GetVpcRegion() *string
}

type CreateVpcFirewallCenConfigureRequest struct {
	// The instance ID of the CEN.
	//
	// > Prerequisite: The CEN instance must be created by invoking the CreateCen operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// cen-x5jayxou71ad73****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The settings for the switch status of the virtual private cloud (VPC) firewalls at the border. Valid values:
	//
	// - **open*	- (default): Automatically enables protection after you create a VPC and the virtual private cloud (VPC) firewalls at the border are created.
	//
	// - **close**: Does not automatically enable protection after the virtual private cloud (VPC) firewalls at the border are created. You can invoke the [ModifyVpcFirewallCenSwitchStatus](https://help.aliyun.com/document_detail/345780.html) operation to enable protection.
	//
	// This parameter is required.
	//
	// example:
	//
	// open
	FirewallSwitch *string `json:"FirewallSwitch,omitempty" xml:"FirewallSwitch,omitempty"`
	// The vSwitch CIDR block used by the firewall. You must configure a CIDR block with a subnet mask of no more than 29 bits that does not conflict with your network planning to allocate to the vSwitch CIDR block required for the procedure of creating the firewall, in order to perform automatic creation of the vSwitch (Cloud_Firewall_VSWITCH) in the firewall secure VPC for traffic diversion and processing. The vSwitch CIDR block must be a subnet of the firewall VPC CIDR block.
	//
	// If you do not specify this parameter, the default is to automatically allocate the 10.219.219.216/29 CIDR block.
	//
	// > This parameter takes effect only when you create a VPC firewall in the local region of the CEN for the first time.
	//
	// example:
	//
	// 10.219.219.216/29
	FirewallVSwitchCidrBlock *string `json:"FirewallVSwitchCidrBlock,omitempty" xml:"FirewallVSwitchCidrBlock,omitempty"`
	// The ID of the zone to which the firewall vSwitch belongs.
	//
	// example:
	//
	// cn-hangzhou-i
	FirewallVSwitchZoneId *string `json:"FirewallVSwitchZoneId,omitempty" xml:"FirewallVSwitchZoneId,omitempty"`
	// The VPC CIDR block used by the firewall. You must configure a CIDR block with a subnet mask of no more than 28 bits to allocate to the VPC CIDR block required for the procedure of creating the firewall, in order to perform automatic creation of the firewall secure VPC (Cloud_Firewall_VPC) for traffic diversion and processing.
	//
	// If you do not specify this parameter, the default is to automatically allocate the 10.0.0.0/8 CIDR block.
	//
	// > This parameter takes effect only when you create a VPC firewall in the local region of the CEN for the first time.
	//
	// example:
	//
	// 10.0.0.0/8
	FirewallVpcCidrBlock *string `json:"FirewallVpcCidrBlock,omitempty" xml:"FirewallVpcCidrBlock,omitempty"`
	// The ID of the secondary zone to which the firewall belongs. Only when the primary zone service is not active, the firewall will perform automatic switchover to the secondary zone to continue running.
	//
	// If you do not specify this parameter, the default is to automatically allocate the secondary zone for the firewall.
	//
	//
	//
	// > This parameter takes effect only when you create a VPC firewall in the local region of the CEN for the first time.
	//
	// example:
	//
	// 10.219.219.216/29
	FirewallVpcStandbyZoneId *string `json:"FirewallVpcStandbyZoneId,omitempty" xml:"FirewallVpcStandbyZoneId,omitempty"`
	// The ID of the primary zone to which the firewall belongs.
	//
	// If your business is sensitive to latency, you can set the zone of the firewall to be the same as the zone of the vSwitch in your business VPC to reduce latency.
	//
	// If you do not specify this parameter, the system automatically assigns a zone for the firewall.
	//
	//
	//
	// > This parameter takes effect only when you create a VPC firewall in the local region of the CEN for the first time.
	//
	// example:
	//
	// cn-hangzhou-a
	FirewallVpcZoneId *string `json:"FirewallVpcZoneId,omitempty" xml:"FirewallVpcZoneId,omitempty"`
	// The language of the request and response. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The UID of the member accounts of the current Alibaba Cloud account.
	//
	// example:
	//
	// 258039427902****
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The instance ID of the VPC-connected instance for which you want to create virtual private cloud (VPC) firewalls at the border.
	//
	// > Prerequisite: The VPC must be added to the CEN specified by the CenId parameter by invoking the AttachCenChildInstance operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp10zlifxh6j0232w****
	NetworkInstanceId *string `json:"NetworkInstanceId,omitempty" xml:"NetworkInstanceId,omitempty"`
	// The ID of the vSwitch to which the Cloud Firewall interface belongs.
	//
	// example:
	//
	// vsw-qzeaol304m***
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The instance name of the virtual private cloud (VPC) firewalls at the border.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-firewall-test
	VpcFirewallName *string `json:"VpcFirewallName,omitempty" xml:"VpcFirewallName,omitempty"`
	// The region ID of the VPC for which you want to create virtual private cloud (VPC) firewalls at the border.
	//
	// > For more information about the regions supported by Cloud Firewall, see [Supported regions](https://help.aliyun.com/document_detail/195657.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	VpcRegion *string `json:"VpcRegion,omitempty" xml:"VpcRegion,omitempty"`
}

func (s CreateVpcFirewallCenConfigureRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVpcFirewallCenConfigureRequest) GoString() string {
	return s.String()
}

func (s *CreateVpcFirewallCenConfigureRequest) GetCenId() *string {
	return s.CenId
}

func (s *CreateVpcFirewallCenConfigureRequest) GetFirewallSwitch() *string {
	return s.FirewallSwitch
}

func (s *CreateVpcFirewallCenConfigureRequest) GetFirewallVSwitchCidrBlock() *string {
	return s.FirewallVSwitchCidrBlock
}

func (s *CreateVpcFirewallCenConfigureRequest) GetFirewallVSwitchZoneId() *string {
	return s.FirewallVSwitchZoneId
}

func (s *CreateVpcFirewallCenConfigureRequest) GetFirewallVpcCidrBlock() *string {
	return s.FirewallVpcCidrBlock
}

func (s *CreateVpcFirewallCenConfigureRequest) GetFirewallVpcStandbyZoneId() *string {
	return s.FirewallVpcStandbyZoneId
}

func (s *CreateVpcFirewallCenConfigureRequest) GetFirewallVpcZoneId() *string {
	return s.FirewallVpcZoneId
}

func (s *CreateVpcFirewallCenConfigureRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateVpcFirewallCenConfigureRequest) GetMemberUid() *string {
	return s.MemberUid
}

func (s *CreateVpcFirewallCenConfigureRequest) GetNetworkInstanceId() *string {
	return s.NetworkInstanceId
}

func (s *CreateVpcFirewallCenConfigureRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateVpcFirewallCenConfigureRequest) GetVpcFirewallName() *string {
	return s.VpcFirewallName
}

func (s *CreateVpcFirewallCenConfigureRequest) GetVpcRegion() *string {
	return s.VpcRegion
}

func (s *CreateVpcFirewallCenConfigureRequest) SetCenId(v string) *CreateVpcFirewallCenConfigureRequest {
	s.CenId = &v
	return s
}

func (s *CreateVpcFirewallCenConfigureRequest) SetFirewallSwitch(v string) *CreateVpcFirewallCenConfigureRequest {
	s.FirewallSwitch = &v
	return s
}

func (s *CreateVpcFirewallCenConfigureRequest) SetFirewallVSwitchCidrBlock(v string) *CreateVpcFirewallCenConfigureRequest {
	s.FirewallVSwitchCidrBlock = &v
	return s
}

func (s *CreateVpcFirewallCenConfigureRequest) SetFirewallVSwitchZoneId(v string) *CreateVpcFirewallCenConfigureRequest {
	s.FirewallVSwitchZoneId = &v
	return s
}

func (s *CreateVpcFirewallCenConfigureRequest) SetFirewallVpcCidrBlock(v string) *CreateVpcFirewallCenConfigureRequest {
	s.FirewallVpcCidrBlock = &v
	return s
}

func (s *CreateVpcFirewallCenConfigureRequest) SetFirewallVpcStandbyZoneId(v string) *CreateVpcFirewallCenConfigureRequest {
	s.FirewallVpcStandbyZoneId = &v
	return s
}

func (s *CreateVpcFirewallCenConfigureRequest) SetFirewallVpcZoneId(v string) *CreateVpcFirewallCenConfigureRequest {
	s.FirewallVpcZoneId = &v
	return s
}

func (s *CreateVpcFirewallCenConfigureRequest) SetLang(v string) *CreateVpcFirewallCenConfigureRequest {
	s.Lang = &v
	return s
}

func (s *CreateVpcFirewallCenConfigureRequest) SetMemberUid(v string) *CreateVpcFirewallCenConfigureRequest {
	s.MemberUid = &v
	return s
}

func (s *CreateVpcFirewallCenConfigureRequest) SetNetworkInstanceId(v string) *CreateVpcFirewallCenConfigureRequest {
	s.NetworkInstanceId = &v
	return s
}

func (s *CreateVpcFirewallCenConfigureRequest) SetVSwitchId(v string) *CreateVpcFirewallCenConfigureRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateVpcFirewallCenConfigureRequest) SetVpcFirewallName(v string) *CreateVpcFirewallCenConfigureRequest {
	s.VpcFirewallName = &v
	return s
}

func (s *CreateVpcFirewallCenConfigureRequest) SetVpcRegion(v string) *CreateVpcFirewallCenConfigureRequest {
	s.VpcRegion = &v
	return s
}

func (s *CreateVpcFirewallCenConfigureRequest) Validate() error {
	return dara.Validate(s)
}
