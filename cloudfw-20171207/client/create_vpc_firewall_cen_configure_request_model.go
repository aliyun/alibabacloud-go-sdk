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
	// The ID of the CEN instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cen-x5jayxou71ad73****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The status of the VPC firewall. Valid values:
	//
	// - **open*	- (default): The VPC firewall is enabled after it is created.
	//
	// - **close**: The VPC firewall is disabled after it is created. Call the [ModifyVpcFirewallCenSwitchStatus](https://help.aliyun.com/document_detail/345780.html) operation to enable the firewall.
	//
	// This parameter is required.
	//
	// example:
	//
	// open
	FirewallSwitch *string `json:"FirewallSwitch,omitempty" xml:"FirewallSwitch,omitempty"`
	// The CIDR block of the vSwitch that is used by the firewall. Specify a CIDR block with a subnet mask of 29 bits or less that does not conflict with your network plan. This CIDR block is allocated to the vSwitch that is automatically created in the firewall VPC (Cloud_Firewall_VSWITCH) for traffic redirection. The vSwitch CIDR block must be a subnet of the firewall VPC CIDR block. If you do not specify this parameter, the system automatically allocates the 10.219.219.216/29 CIDR block.
	//
	// If you leave this parameter empty, the CIDR block 10.219.219.216/29 is automatically allocated.
	//
	// > This parameter is valid only when you create a VPC firewall for the first time in the current region of the CEN.
	FirewallVSwitchCidrBlock *string `json:"FirewallVSwitchCidrBlock,omitempty" xml:"FirewallVSwitchCidrBlock,omitempty"`
	// The CIDR block of the VPC that is used by the firewall. Specify a CIDR block with a subnet mask of 28 bits or less. This CIDR block is allocated to the VPC that is automatically created for the firewall for traffic redirection. If you do not specify this parameter, the system automatically allocates the 10.0.0.0/8 CIDR block.
	//
	// If you leave this property empty, the CIDR block 10.0.0.0/8 is automatically allocated.
	//
	// > This parameter is valid only when you create a VPC firewall for the first time in the current region of the CEN.
	//
	// example:
	//
	// 10.0.0.0/8
	FirewallVpcCidrBlock *string `json:"FirewallVpcCidrBlock,omitempty" xml:"FirewallVpcCidrBlock,omitempty"`
	// The ID of the secondary zone for the firewall. If the service in the primary zone becomes unavailable, the firewall automatically switches to the secondary zone. If you do not specify this parameter, the system automatically assigns a secondary zone for the firewall.
	//
	// If you do not specify a value, a zone is automatically allocated to the VPC firewall.
	//
	// > This parameter is valid only when you create a VPC firewall for the first time in the current region of the CEN.
	//
	// example:
	//
	// cn-hangzhou-b
	FirewallVpcStandbyZoneId *string `json:"FirewallVpcStandbyZoneId,omitempty" xml:"FirewallVpcStandbyZoneId,omitempty"`
	// The ID of the primary zone for the firewall. If your business is sensitive to latency, specify the same zone for the firewall and the vSwitch of your business VPC to reduce latency. If you do not specify this parameter, the system automatically assigns a zone for the firewall.
	//
	// If you do not specify a value, a zone is automatically allocated to the VPC firewall.
	//
	// > This parameter is valid only when you create a VPC firewall for the first time in the current region of the CEN.
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
	// The UID of the member account.
	//
	// example:
	//
	// 258039427902****
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The ID of the VPC for which you want to create the VPC firewall.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp10zlifxh6j0232w****
	NetworkInstanceId *string `json:"NetworkInstanceId,omitempty" xml:"NetworkInstanceId,omitempty"`
	// The ID of the vSwitch that is used by the Cloud Firewall interface.
	//
	// example:
	//
	// vsw-qzeaol304m***
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The name of the VPC firewall instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-firewall-test
	VpcFirewallName *string `json:"VpcFirewallName,omitempty" xml:"VpcFirewallName,omitempty"`
	// The region ID of the VPC for which you want to create the VPC firewall.
	//
	// > For more information about the regions that Cloud Firewall supports, see [Supported regions](https://help.aliyun.com/document_detail/195657.html).
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
