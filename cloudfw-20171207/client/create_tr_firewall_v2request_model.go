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
	SetFirewallAttachmentZone(v string) *CreateTrFirewallV2Request
	GetFirewallAttachmentZone() *string
	SetFirewallDescription(v string) *CreateTrFirewallV2Request
	GetFirewallDescription() *string
	SetFirewallName(v string) *CreateTrFirewallV2Request
	GetFirewallName() *string
	SetFirewallServiceMode(v string) *CreateTrFirewallV2Request
	GetFirewallServiceMode() *string
	SetFirewallServiceZones(v []*string) *CreateTrFirewallV2Request
	GetFirewallServiceZones() []*string
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
	SetTrAttachmentZones(v []*string) *CreateTrFirewallV2Request
	GetTrAttachmentZones() []*string
	SetTransitRouterId(v string) *CreateTrFirewallV2Request
	GetTransitRouterId() *string
}

type CreateTrFirewallV2Request struct {
	// The ID of the CEN instance. Create a CEN instance in the CEN console first and make sure that an Enterprise Edition transit router has been created.
	//
	// > Note: Although this parameter is marked as not required in the schema, it is actually required. If this parameter is not specified, the ErrorParameters (400) error is returned.
	//
	// example:
	//
	// cen-4xbjup276au29r****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The zone ID used by the firewall connection.
	//
	// example:
	//
	// cn-hangzhou-h
	FirewallAttachmentZone *string `json:"FirewallAttachmentZone,omitempty" xml:"FirewallAttachmentZone,omitempty"`
	// The description of the firewall.
	//
	// example:
	//
	// vpc-firewall-description
	FirewallDescription *string `json:"FirewallDescription,omitempty" xml:"FirewallDescription,omitempty"`
	// The name of the Cloud Firewall instance.
	//
	// example:
	//
	// vpc-firewall-test
	FirewallName *string `json:"FirewallName,omitempty" xml:"FirewallName,omitempty"`
	// The deployment mode of the firewall service. Valid values:
	//
	// - **PrimaryStandby**: Primary/standby mode.
	//
	// - **MultiPrimary**: Active-active mode.
	//
	// > If this parameter is not specified, the system automatically selects a deployment mode based on the capabilities of the transit router. If an invalid value is specified, the ErrorFwServiceMode (-360437) error is returned. The MultiPrimary mode does not support specifying zones.
	//
	// example:
	//
	// PrimaryStandby
	FirewallServiceMode *string `json:"FirewallServiceMode,omitempty" xml:"FirewallServiceMode,omitempty"`
	// The list of zone IDs used by the firewall service.
	FirewallServiceZones []*string `json:"FirewallServiceZones,omitempty" xml:"FirewallServiceZones,omitempty" type:"Repeated"`
	// Deprecated
	//
	// The subnet CIDR block used to store the firewall elastic network interface (ENI) in the firewall VPC in automatic mode.
	//
	// example:
	//
	// 10.0.1.0/24
	FirewallSubnetCidr *string `json:"FirewallSubnetCidr,omitempty" xml:"FirewallSubnetCidr,omitempty"`
	// The CIDR block of the firewall VPC in automatic mode.
	//
	// example:
	//
	// 10.0.0.0/16
	FirewallVpcCidr *string `json:"FirewallVpcCidr,omitempty" xml:"FirewallVpcCidr,omitempty"`
	// The ID of the VPC in which the firewall ENI is created in manual mode.
	//
	// example:
	//
	// vpc-wz9r5qvryn0lg3atb****
	FirewallVpcId *string `json:"FirewallVpcId,omitempty" xml:"FirewallVpcId,omitempty"`
	// The ID of the vSwitch in which the firewall ENI is created in manual mode.
	//
	// example:
	//
	// vsw-uf6ydz3vqj77mr5l6****
	FirewallVswitchId *string `json:"FirewallVswitchId,omitempty" xml:"FirewallVswitchId,omitempty"`
	// The language of the response. Valid values:
	//
	// - **zh*	- (default): Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The region ID of the Enterprise Edition transit router.
	//
	// > Note: Although this parameter is marked as not required in the schema, it is actually required. If this parameter is not specified, the ErrorParameters (400) error is returned.
	//
	// example:
	//
	// cn-hangzhou
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The routing mode. Valid values: managed (automatic mode) and manual (manual mode). In managed mode, only FirewallVpcCidr is required. The FirewallSubnetCidr, TrAttachmentSlaveCidr, and TrAttachmentMasterCidr parameters are deprecated and do not need to be specified. In manual mode, specify FirewallVpcId, FirewallVswitchId, TrAttachmentSlaveZone, and TrAttachmentMasterZone.
	//
	// > Note: Although this parameter is marked as not required in the schema, it is actually required. If this parameter is not specified, the ErrorParameters (400) error is returned.
	//
	// example:
	//
	// managed
	RouteMode *string `json:"RouteMode,omitempty" xml:"RouteMode,omitempty"`
	// Deprecated
	//
	// The primary subnet CIDR block used to connect to the TR in the firewall VPC in automatic mode.
	//
	// example:
	//
	// 10.0.3.0/24
	TrAttachmentMasterCidr *string `json:"TrAttachmentMasterCidr,omitempty" xml:"TrAttachmentMasterCidr,omitempty"`
	// The primary zone of the vSwitch.
	//
	// example:
	//
	// cn-chengdu-a
	TrAttachmentMasterZone *string `json:"TrAttachmentMasterZone,omitempty" xml:"TrAttachmentMasterZone,omitempty"`
	// Deprecated
	//
	// The secondary subnet CIDR block used to connect to the TR in the firewall VPC in automatic mode.
	//
	// example:
	//
	// 10.0.0.16/28
	TrAttachmentSlaveCidr *string `json:"TrAttachmentSlaveCidr,omitempty" xml:"TrAttachmentSlaveCidr,omitempty"`
	// The secondary zone of the vSwitch.
	//
	// example:
	//
	// cn-chengdu-b
	TrAttachmentSlaveZone *string `json:"TrAttachmentSlaveZone,omitempty" xml:"TrAttachmentSlaveZone,omitempty"`
	// The list of zone IDs used by the TR connection.
	TrAttachmentZones []*string `json:"TrAttachmentZones,omitempty" xml:"TrAttachmentZones,omitempty" type:"Repeated"`
	// The ID of the Enterprise Edition transit router instance. The transit router must belong to the CEN instance specified by CenId.
	//
	// > Note: Although this parameter is marked as not required in the schema, it is actually required. If this parameter is not specified, the ErrorParameters (400) error is returned.
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

func (s *CreateTrFirewallV2Request) GetFirewallAttachmentZone() *string {
	return s.FirewallAttachmentZone
}

func (s *CreateTrFirewallV2Request) GetFirewallDescription() *string {
	return s.FirewallDescription
}

func (s *CreateTrFirewallV2Request) GetFirewallName() *string {
	return s.FirewallName
}

func (s *CreateTrFirewallV2Request) GetFirewallServiceMode() *string {
	return s.FirewallServiceMode
}

func (s *CreateTrFirewallV2Request) GetFirewallServiceZones() []*string {
	return s.FirewallServiceZones
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

func (s *CreateTrFirewallV2Request) GetTrAttachmentZones() []*string {
	return s.TrAttachmentZones
}

func (s *CreateTrFirewallV2Request) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *CreateTrFirewallV2Request) SetCenId(v string) *CreateTrFirewallV2Request {
	s.CenId = &v
	return s
}

func (s *CreateTrFirewallV2Request) SetFirewallAttachmentZone(v string) *CreateTrFirewallV2Request {
	s.FirewallAttachmentZone = &v
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

func (s *CreateTrFirewallV2Request) SetFirewallServiceMode(v string) *CreateTrFirewallV2Request {
	s.FirewallServiceMode = &v
	return s
}

func (s *CreateTrFirewallV2Request) SetFirewallServiceZones(v []*string) *CreateTrFirewallV2Request {
	s.FirewallServiceZones = v
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

func (s *CreateTrFirewallV2Request) SetTrAttachmentZones(v []*string) *CreateTrFirewallV2Request {
	s.TrAttachmentZones = v
	return s
}

func (s *CreateTrFirewallV2Request) SetTransitRouterId(v string) *CreateTrFirewallV2Request {
	s.TransitRouterId = &v
	return s
}

func (s *CreateTrFirewallV2Request) Validate() error {
	return dara.Validate(s)
}
