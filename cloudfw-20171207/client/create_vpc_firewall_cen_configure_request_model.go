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
	// The instance ID of the CEN instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cen-x5jayxou71ad73****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The status of the virtual private cloud (VPC) firewall after you create a VPC firewall. Valid values:
	//
	// - **open*	- (default): The VPC firewall is automatically enabled after it is created.
	//
	// - **close**: The VPC firewall is not automatically enabled after it is created. You can invoke the [ModifyVpcFirewallCenSwitchStatus](https://help.aliyun.com/document_detail/345780.html) operation to enable the VPC firewall.
	//
	// This parameter is required.
	//
	// example:
	//
	// open
	FirewallSwitch *string `json:"FirewallSwitch,omitempty" xml:"FirewallSwitch,omitempty"`
	// The CIDR block of the vSwitch used by the firewall. Specify a CIDR block with a subnet mask of no more than 29 bits that does not conflict with your network planning. This CIDR block is allocated to the vSwitch that is created during the process to create a VPC firewall within the firewall security VPC (Cloud_Firewall_VSWITCH) for traffic redirection. The vSwitch CIDR block must be a subnet of the firewall VPC CIDR block.
	//
	// If you leave this parameter empty, the default CIDR block 10.219.219.216/29 is automatically allocated.
	//
	// > This parameter takes effect only when a VPC firewall is created for the first time in the local region of the CEN instance.
	FirewallVSwitchCidrBlock *string `json:"FirewallVSwitchCidrBlock,omitempty" xml:"FirewallVSwitchCidrBlock,omitempty"`
	// The CIDR block of the VPC used by the firewall. Specify a CIDR block with a subnet mask of no more than 28 bits. This CIDR block is allocated to the VPC that is created during the process to create a VPC firewall (Cloud_Firewall_VPC) for traffic redirection.
	//
	// If you leave this parameter empty, the default CIDR block 10.0.0.0/8 is automatically allocated.
	//
	// > This parameter takes effect only when a VPC firewall is created for the first time in the local region of the CEN instance.
	//
	// example:
	//
	// 10.0.0.0/8
	FirewallVpcCidrBlock *string `json:"FirewallVpcCidrBlock,omitempty" xml:"FirewallVpcCidrBlock,omitempty"`
	// The ID of the secondary zone to which the firewall belongs. The firewall performs an automatic switchover to the secondary zone to continue running only when the primary zone becomes unavailable.
	//
	// If you leave this parameter empty, a default secondary zone is automatically allocated for the firewall.
	//
	//
	//
	// > This parameter takes effect only when a VPC firewall is created for the first time in the local region of the CEN instance.
	//
	// example:
	//
	// cn-hangzhou-b
	FirewallVpcStandbyZoneId *string `json:"FirewallVpcStandbyZoneId,omitempty" xml:"FirewallVpcStandbyZoneId,omitempty"`
	// The ID of the primary active zone to which the firewall belongs. If your business is latency-sensitive, you can set the firewall zone to the same zone as the vSwitch of the business VPC to reduce latency.
	//
	// If you leave this parameter empty, a default zone is automatically allocated for the firewall.
	//
	//
	//
	// > This parameter takes effect only when a VPC firewall is created for the first time in the local region of the CEN instance.
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
	// The UID of the member account of the current Alibaba Cloud account.
	//
	// example:
	//
	// 258039427902****
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The instance ID of the VPC for which you want to create a VPC firewall.
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
	// The instance name of the virtual private cloud (VPC) firewall.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-firewall-test
	VpcFirewallName *string `json:"VpcFirewallName,omitempty" xml:"VpcFirewallName,omitempty"`
	// The region ID of the VPC for which you want to create a VPC firewall.
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
