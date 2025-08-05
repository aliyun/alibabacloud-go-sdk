// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTrFirewallV2Request interface {
	dara.Model
	String() string
	GoString() string
	SetCenId(v string) *CreateTrFirewallV2Request
	GetCenId() *string
	SetFirewallDescription(v string) *CreateTrFirewallV2Request
	GetFirewallDescription() *string
	SetFirewallName(v string) *CreateTrFirewallV2Request
	GetFirewallName() *string
	SetFirewallSubnetCidr(v string) *CreateTrFirewallV2Request
	GetFirewallSubnetCidr() *string
	SetFirewallVpcCidr(v string) *CreateTrFirewallV2Request
	GetFirewallVpcCidr() *string
	SetFirewallVpcId(v string) *CreateTrFirewallV2Request
	GetFirewallVpcId() *string
	SetFirewallVswitchId(v string) *CreateTrFirewallV2Request
	GetFirewallVswitchId() *string
	SetLang(v string) *CreateTrFirewallV2Request
	GetLang() *string
	SetRegionNo(v string) *CreateTrFirewallV2Request
	GetRegionNo() *string
	SetRouteMode(v string) *CreateTrFirewallV2Request
	GetRouteMode() *string
	SetTrAttachmentMasterCidr(v string) *CreateTrFirewallV2Request
	GetTrAttachmentMasterCidr() *string
	SetTrAttachmentMasterZone(v string) *CreateTrFirewallV2Request
	GetTrAttachmentMasterZone() *string
	SetTrAttachmentSlaveCidr(v string) *CreateTrFirewallV2Request
	GetTrAttachmentSlaveCidr() *string
	SetTrAttachmentSlaveZone(v string) *CreateTrFirewallV2Request
	GetTrAttachmentSlaveZone() *string
	SetTransitRouterId(v string) *CreateTrFirewallV2Request
	GetTransitRouterId() *string
}

type CreateTrFirewallV2Request struct {
	// The ID of the Cloud Enterprise Network (CEN) instance.
	//
	// example:
	//
	// cen-4xbjup276au29r****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The description of the firewall.
	//
	// example:
	//
	// vfw-sz
	FirewallDescription *string `json:"FirewallDescription,omitempty" xml:"FirewallDescription,omitempty"`
	// The name of the firewall.
	//
	// example:
	//
	// cfw_test2
	FirewallName *string `json:"FirewallName,omitempty" xml:"FirewallName,omitempty"`
	// The subnet CIDR block of the VPC in which the ENI of the firewall is stored in automatic mode.
	//
	// example:
	//
	// 10.0.1.0/24
	FirewallSubnetCidr *string `json:"FirewallSubnetCidr,omitempty" xml:"FirewallSubnetCidr,omitempty"`
	// The CIDR block that is allocated to the VPC created for the VPC firewall in automatic mode.
	//
	// example:
	//
	// 10.0.0.0/16
	FirewallVpcCidr *string `json:"FirewallVpcCidr,omitempty" xml:"FirewallVpcCidr,omitempty"`
	// The ID of the VPC in which the ENI associated with the VPC firewall is created in manual mode.
	//
	// example:
	//
	// vpc-wz9r5qvryn0lg3atb****
	FirewallVpcId *string `json:"FirewallVpcId,omitempty" xml:"FirewallVpcId,omitempty"`
	// The ID of the vSwitch that is used to create the ENI in manual mode.
	//
	// example:
	//
	// vsw-uf6ydz3vqj77mr5l6****
	FirewallVswitchId *string `json:"FirewallVswitchId,omitempty" xml:"FirewallVswitchId,omitempty"`
	// The language of the content within the response. Valid values:
	//
	// 	- **zh**: Chinese (default)
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The region ID of the route router.
	//
	// example:
	//
	// cn-hangzhou
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The routing mode of the VPC firewall. Valid values:
	//
	// 	- **managed**: automatic mode
	//
	// 	- **manual**: manual mode
	//
	// example:
	//
	// managed
	RouteMode *string `json:"RouteMode,omitempty" xml:"RouteMode,omitempty"`
	// The primary subnet CIDR block that the VPC uses to connect to the transit router in automatic mode.
	//
	// example:
	//
	// 10.0.3.0/24
	TrAttachmentMasterCidr *string `json:"TrAttachmentMasterCidr,omitempty" xml:"TrAttachmentMasterCidr,omitempty"`
	// The primary zone for the vSwitch.
	//
	// example:
	//
	// cn-chengdu-a
	TrAttachmentMasterZone *string `json:"TrAttachmentMasterZone,omitempty" xml:"TrAttachmentMasterZone,omitempty"`
	// The secondary subnet CIDR block that the VPC uses to connect to the transit router in automatic mode.
	//
	// example:
	//
	// 10.0.0.16/28
	TrAttachmentSlaveCidr *string `json:"TrAttachmentSlaveCidr,omitempty" xml:"TrAttachmentSlaveCidr,omitempty"`
	// The secondary zone for the vSwitch.
	//
	// example:
	//
	// cn-chengdu-b
	TrAttachmentSlaveZone *string `json:"TrAttachmentSlaveZone,omitempty" xml:"TrAttachmentSlaveZone,omitempty"`
	// The ID of the transit router.
	//
	// example:
	//
	// tr-m5etmb2q7e0mxcur****
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
}

func (s CreateTrFirewallV2Request) String() string {
	return dara.Prettify(s)
}

func (s CreateTrFirewallV2Request) GoString() string {
	return s.String()
}

func (s *CreateTrFirewallV2Request) GetCenId() *string {
	return s.CenId
}

func (s *CreateTrFirewallV2Request) GetFirewallDescription() *string {
	return s.FirewallDescription
}

func (s *CreateTrFirewallV2Request) GetFirewallName() *string {
	return s.FirewallName
}

func (s *CreateTrFirewallV2Request) GetFirewallSubnetCidr() *string {
	return s.FirewallSubnetCidr
}

func (s *CreateTrFirewallV2Request) GetFirewallVpcCidr() *string {
	return s.FirewallVpcCidr
}

func (s *CreateTrFirewallV2Request) GetFirewallVpcId() *string {
	return s.FirewallVpcId
}

func (s *CreateTrFirewallV2Request) GetFirewallVswitchId() *string {
	return s.FirewallVswitchId
}

func (s *CreateTrFirewallV2Request) GetLang() *string {
	return s.Lang
}

func (s *CreateTrFirewallV2Request) GetRegionNo() *string {
	return s.RegionNo
}

func (s *CreateTrFirewallV2Request) GetRouteMode() *string {
	return s.RouteMode
}

func (s *CreateTrFirewallV2Request) GetTrAttachmentMasterCidr() *string {
	return s.TrAttachmentMasterCidr
}

func (s *CreateTrFirewallV2Request) GetTrAttachmentMasterZone() *string {
	return s.TrAttachmentMasterZone
}

func (s *CreateTrFirewallV2Request) GetTrAttachmentSlaveCidr() *string {
	return s.TrAttachmentSlaveCidr
}

func (s *CreateTrFirewallV2Request) GetTrAttachmentSlaveZone() *string {
	return s.TrAttachmentSlaveZone
}

func (s *CreateTrFirewallV2Request) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *CreateTrFirewallV2Request) SetCenId(v string) *CreateTrFirewallV2Request {
	s.CenId = &v
	return s
}

func (s *CreateTrFirewallV2Request) SetFirewallDescription(v string) *CreateTrFirewallV2Request {
	s.FirewallDescription = &v
	return s
}

func (s *CreateTrFirewallV2Request) SetFirewallName(v string) *CreateTrFirewallV2Request {
	s.FirewallName = &v
	return s
}

func (s *CreateTrFirewallV2Request) SetFirewallSubnetCidr(v string) *CreateTrFirewallV2Request {
	s.FirewallSubnetCidr = &v
	return s
}

func (s *CreateTrFirewallV2Request) SetFirewallVpcCidr(v string) *CreateTrFirewallV2Request {
	s.FirewallVpcCidr = &v
	return s
}

func (s *CreateTrFirewallV2Request) SetFirewallVpcId(v string) *CreateTrFirewallV2Request {
	s.FirewallVpcId = &v
	return s
}

func (s *CreateTrFirewallV2Request) SetFirewallVswitchId(v string) *CreateTrFirewallV2Request {
	s.FirewallVswitchId = &v
	return s
}

func (s *CreateTrFirewallV2Request) SetLang(v string) *CreateTrFirewallV2Request {
	s.Lang = &v
	return s
}

func (s *CreateTrFirewallV2Request) SetRegionNo(v string) *CreateTrFirewallV2Request {
	s.RegionNo = &v
	return s
}

func (s *CreateTrFirewallV2Request) SetRouteMode(v string) *CreateTrFirewallV2Request {
	s.RouteMode = &v
	return s
}

func (s *CreateTrFirewallV2Request) SetTrAttachmentMasterCidr(v string) *CreateTrFirewallV2Request {
	s.TrAttachmentMasterCidr = &v
	return s
}

func (s *CreateTrFirewallV2Request) SetTrAttachmentMasterZone(v string) *CreateTrFirewallV2Request {
	s.TrAttachmentMasterZone = &v
	return s
}

func (s *CreateTrFirewallV2Request) SetTrAttachmentSlaveCidr(v string) *CreateTrFirewallV2Request {
	s.TrAttachmentSlaveCidr = &v
	return s
}

func (s *CreateTrFirewallV2Request) SetTrAttachmentSlaveZone(v string) *CreateTrFirewallV2Request {
	s.TrAttachmentSlaveZone = &v
	return s
}

func (s *CreateTrFirewallV2Request) SetTransitRouterId(v string) *CreateTrFirewallV2Request {
	s.TransitRouterId = &v
	return s
}

func (s *CreateTrFirewallV2Request) Validate() error {
	return dara.Validate(s)
}
