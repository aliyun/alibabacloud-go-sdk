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
	// The Cloud Enterprise Network (CEN) instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cen-x5jayxou71ad73****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The switch status of the VPC border firewall. Valid values:
	//
	// - **open*	- (default): Protection is automatically enabled after the VPC border firewall is created.
	//
	// - **close**: Protection is not automatically enabled after the VPC border firewall is created. You can call the [ModifyVpcFirewallCenSwitchStatus](https://help.aliyun.com/document_detail/345780.html) operation to enable protection.
	//
	// This parameter is required.
	//
	// example:
	//
	// open
	FirewallSwitch *string `json:"FirewallSwitch,omitempty" xml:"FirewallSwitch,omitempty"`
	// The vSwitch CIDR block used by the firewall. You must configure a CIDR block with a subnet mask no larger than 29 bits that does not conflict with the network plan to allocate a vSwitch CIDR block for the firewall creation process. This is used to automatically create a vSwitch (Cloud_Firewall_VSWITCH) within the firewall security VPC for traffic redirection. The vSwitch CIDR block must be a subnet of the firewall VPC CIDR block.
	//
	// If this parameter is not specified, the CIDR block 10.219.219.216/29 is automatically allocated by default.
	//
	// > This parameter is only valid when a VPC firewall is created for the first time in this CEN region.
	FirewallVSwitchCidrBlock *string `json:"FirewallVSwitchCidrBlock,omitempty" xml:"FirewallVSwitchCidrBlock,omitempty"`
	// The VPC CIDR block used by the firewall. You must configure a CIDR block with a subnet mask no larger than 28 bits to allocate a VPC CIDR block for the firewall creation process. This is used to automatically create a firewall security VPC (Cloud_Firewall_VPC) for traffic redirection.
	//
	// If this parameter is not specified, the CIDR block 10.0.0.0/8 is automatically allocated by default.
	//
	// > This parameter is only valid when a VPC firewall is created for the first time in this CEN region.
	//
	// example:
	//
	// 10.0.0.0/8
	FirewallVpcCidrBlock *string `json:"FirewallVpcCidrBlock,omitempty" xml:"FirewallVpcCidrBlock,omitempty"`
	// The standby availability zone ID of the firewall. The firewall automatically switches to the standby availability zone and continues to run only when the primary availability zone service is unavailable.
	//
	// If this parameter is not specified, the firewall standby availability zone is automatically allocated by default.
	//
	//
	//
	// > This parameter is only valid when a VPC firewall is created for the first time in this CEN region.
	//
	// example:
	//
	// cn-hangzhou-b
	FirewallVpcStandbyZoneId *string `json:"FirewallVpcStandbyZoneId,omitempty" xml:"FirewallVpcStandbyZoneId,omitempty"`
	// The primary availability zone ID of the firewall.
	//
	// If your business is latency-sensitive, you can customize the firewall availability zone to be the same as your business VPC vSwitch availability zone to reduce latency.
	//
	// If this parameter is not specified, the firewall availability zone is automatically allocated by default.
	//
	//
	//
	// > This parameter is only valid when a VPC firewall is created for the first time in this CEN region.
	//
	// example:
	//
	// cn-hangzhou-a
	FirewallVpcZoneId *string `json:"FirewallVpcZoneId,omitempty" xml:"FirewallVpcZoneId,omitempty"`
	// The language type for the request and response messages. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The member account UID of the current Alibaba Cloud account.
	//
	// example:
	//
	// 258039427902****
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The VPC instance ID for which the VPC border firewall is created.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp10zlifxh6j0232w****
	NetworkInstanceId *string `json:"NetworkInstanceId,omitempty" xml:"NetworkInstanceId,omitempty"`
	// The vSwitch ID to which the Cloud Firewall interface belongs.
	//
	// example:
	//
	// vsw-qzeaol304m***
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The instance name of the VPC border firewall.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-firewall-test
	VpcFirewallName *string `json:"VpcFirewallName,omitempty" xml:"VpcFirewallName,omitempty"`
	// The region ID of the VPC for which the VPC border firewall is created.
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
