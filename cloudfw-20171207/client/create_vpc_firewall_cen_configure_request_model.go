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
	// Specifies whether to enable the VPC firewall. Valid values:
	//
	// 	- **open**: After you create the VPC firewall, the VPC firewall is automatically enabled. This is the default value.
	//
	// 	- **close**: After you create the VPC firewall, the VPC firewall is disabled. You can call the [ModifyVpcFirewallCenSwitchStatus](https://help.aliyun.com/document_detail/345780.html) operation to manually enable the VPC firewall.
	//
	// This parameter is required.
	//
	// example:
	//
	// open
	FirewallSwitch *string `json:"FirewallSwitch,omitempty" xml:"FirewallSwitch,omitempty"`
	// The CIDR block of the vSwitch that is automatically created for the VPC firewall. You must specify a CIDR block for the Cloud_Firewall_VSWITCH VPC that is automatically created for the VPC firewall for traffic redirection. The CIDR block does not conflict with your network plan. The subnet mask of the CIDR block must be less than or equal to 29 bits in length. The CIDR block of the vSwitch must be within the network segment of the VPC.
	//
	// If you do not specify a value, the CIDR block 10.219.219.216/29 is automatically allocated.
	//
	// >  This parameter takes effect only when you create a VPC firewall for the first time in the current CEN instance and region.
	FirewallVSwitchCidrBlock *string `json:"FirewallVSwitchCidrBlock,omitempty" xml:"FirewallVSwitchCidrBlock,omitempty"`
	// The CIDR block of the VPC that is automatically created for the VPC firewall. You must specify a CIDR block for the Cloud_Firewall_VPC VPC that is automatically created for the VPC firewall for traffic redirection. The subnet mask of the CIDR block must be less than or equal to 28 bits in length.
	//
	// If you do not specify a value, the CIDR block 10.0.0.0/8 is automatically allocated.
	//
	// >  This parameter takes effect only when you create a VPC firewall for the first time in the current CEN instance and region.
	//
	// example:
	//
	// 10.0.0.0/8
	FirewallVpcCidrBlock *string `json:"FirewallVpcCidrBlock,omitempty" xml:"FirewallVpcCidrBlock,omitempty"`
	// The ID of the backup availability zone to which the firewall belongs. The firewall will automatically switch to the backup availability zone to continue running only if the primary availability zone service is unavailable.
	//
	// If this parameter is not filled, the backup availability zone for the firewall will be automatically assigned.
	//
	// > This parameter is only effective when creating a VPC firewall for the first time in this CEN region.
	//
	// example:
	//
	// cn-hangzhou-b
	FirewallVpcStandbyZoneId *string `json:"FirewallVpcStandbyZoneId,omitempty" xml:"FirewallVpcStandbyZoneId,omitempty"`
	// The ID of the zone to which the vSwitch belongs. If your service is latency-sensitive, you can specify the same zone for the vSwitch of the firewall and the vSwitch of your business VPC to minimize latency.
	//
	// If you do not specify a value, a zone is automatically assigned for the vSwitch.
	//
	// >  This parameter takes effect only when you create a VPC firewall for the first time in the current CEN instance and region. For more information about zones that are supported by each region, see [Query zones](https://help.aliyun.com/document_detail/36064.html).
	//
	// example:
	//
	// cn-hangzhou-a
	FirewallVpcZoneId *string `json:"FirewallVpcZoneId,omitempty" xml:"FirewallVpcZoneId,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh**: Chinese (default)
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The UID of the member that is managed by your Alibaba Cloud account.
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
	// The ID of the vSwitch that is used to associate with the elastic network interface (ENI) required by the VPC firewall.
	//
	// example:
	//
	// vsw-qzeaol304m***
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The instance name of the VPC firewall.
	//
	// This parameter is required.
	//
	// example:
	//
	// Test instance
	VpcFirewallName *string `json:"VpcFirewallName,omitempty" xml:"VpcFirewallName,omitempty"`
	// The ID of the region to which the VPC belongs.
	//
	// > For more information about the regions, see [Supported regions](https://help.aliyun.com/document_detail/195657.html).
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
