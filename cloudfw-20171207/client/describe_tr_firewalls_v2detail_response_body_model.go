// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTrFirewallsV2DetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCenId(v string) *DescribeTrFirewallsV2DetailResponseBody
	GetCenId() *string
	SetFirewallDescription(v string) *DescribeTrFirewallsV2DetailResponseBody
	GetFirewallDescription() *string
	SetFirewallEniId(v string) *DescribeTrFirewallsV2DetailResponseBody
	GetFirewallEniId() *string
	SetFirewallEniVpcId(v string) *DescribeTrFirewallsV2DetailResponseBody
	GetFirewallEniVpcId() *string
	SetFirewallEniVswitchId(v string) *DescribeTrFirewallsV2DetailResponseBody
	GetFirewallEniVswitchId() *string
	SetFirewallId(v string) *DescribeTrFirewallsV2DetailResponseBody
	GetFirewallId() *string
	SetFirewallName(v string) *DescribeTrFirewallsV2DetailResponseBody
	GetFirewallName() *string
	SetFirewallStatus(v string) *DescribeTrFirewallsV2DetailResponseBody
	GetFirewallStatus() *string
	SetFirewallSubnetCidr(v string) *DescribeTrFirewallsV2DetailResponseBody
	GetFirewallSubnetCidr() *string
	SetFirewallSwitchStatus(v string) *DescribeTrFirewallsV2DetailResponseBody
	GetFirewallSwitchStatus() *string
	SetFirewallVpcCidr(v string) *DescribeTrFirewallsV2DetailResponseBody
	GetFirewallVpcCidr() *string
	SetRegionNo(v string) *DescribeTrFirewallsV2DetailResponseBody
	GetRegionNo() *string
	SetRequestId(v string) *DescribeTrFirewallsV2DetailResponseBody
	GetRequestId() *string
	SetRouteMode(v string) *DescribeTrFirewallsV2DetailResponseBody
	GetRouteMode() *string
	SetTrAttachmentId(v string) *DescribeTrFirewallsV2DetailResponseBody
	GetTrAttachmentId() *string
	SetTrAttachmentMasterCidr(v string) *DescribeTrFirewallsV2DetailResponseBody
	GetTrAttachmentMasterCidr() *string
	SetTrAttachmentMasterZone(v string) *DescribeTrFirewallsV2DetailResponseBody
	GetTrAttachmentMasterZone() *string
	SetTrAttachmentSlaveCidr(v string) *DescribeTrFirewallsV2DetailResponseBody
	GetTrAttachmentSlaveCidr() *string
	SetTrAttachmentSlaveZone(v string) *DescribeTrFirewallsV2DetailResponseBody
	GetTrAttachmentSlaveZone() *string
	SetTransitRouterId(v string) *DescribeTrFirewallsV2DetailResponseBody
	GetTransitRouterId() *string
}

type DescribeTrFirewallsV2DetailResponseBody struct {
	// The ID of the Cloud Enterprise Network (CEN) instance.
	//
	// example:
	//
	// cen-37nddhri7jf0d2****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The description of the VPC firewall.
	//
	// example:
	//
	// VPC Firewall
	FirewallDescription *string `json:"FirewallDescription,omitempty" xml:"FirewallDescription,omitempty"`
	// The ID of the Elastic Network Interface (ENI) with which the VPC firewall is associated.
	//
	// example:
	//
	// eni-uf621u00nafypeex****
	FirewallEniId *string `json:"FirewallEniId,omitempty" xml:"FirewallEniId,omitempty"`
	// The ID of the VPC to which the ENI is attached.
	//
	// example:
	//
	// vpc-2zeppcci782zeh2bk****
	FirewallEniVpcId *string `json:"FirewallEniVpcId,omitempty" xml:"FirewallEniVpcId,omitempty"`
	// The ID of the vSwitch with which the ENI is associated.
	//
	// example:
	//
	// vsw-uf6ptq1kl1c1d9pw9****
	FirewallEniVswitchId *string `json:"FirewallEniVswitchId,omitempty" xml:"FirewallEniVswitchId,omitempty"`
	// The instance ID of the VPC firewall.
	//
	// example:
	//
	// vfw-tr-9c7c711abdfa4d80****
	FirewallId *string `json:"FirewallId,omitempty" xml:"FirewallId,omitempty"`
	// The name of the VPC firewall.
	//
	// example:
	//
	// cloudfirewall-manual
	FirewallName *string `json:"FirewallName,omitempty" xml:"FirewallName,omitempty"`
	// The status of the VPC firewall. Valid values:
	//
	// 	- Creating
	//
	// 	- Deleting
	//
	// 	- Ready
	//
	// example:
	//
	// Ready
	FirewallStatus *string `json:"FirewallStatus,omitempty" xml:"FirewallStatus,omitempty"`
	// The subnet CIDR block of the VPC in which the ENI of the firewall is stored in automatic mode.
	//
	// example:
	//
	// 10.0.1.0/24
	FirewallSubnetCidr *string `json:"FirewallSubnetCidr,omitempty" xml:"FirewallSubnetCidr,omitempty"`
	// The status of the VPC firewall. Valid values:
	//
	// 	- **opened**: The VPC firewall is enabled.
	//
	// 	- **closed**: The VPC firewall is disabled.
	//
	// 	- **notconfigured**: The VPC firewall is not created.
	//
	// 	- **configured**: The VPC firewall is created but is not enabled.
	//
	// 	- **creating**: The VPC firewall is being created.
	//
	// 	- **opening**: The VPC firewall is being enabled.
	//
	// 	- **deleting**: The VPC firewall is being deleted.
	//
	// > If you do not specify this parameter, VPC firewalls in all states are queried.
	//
	// example:
	//
	// opened
	FirewallSwitchStatus *string `json:"FirewallSwitchStatus,omitempty" xml:"FirewallSwitchStatus,omitempty"`
	// The CIDR block that is allocated to the VPC created for the VPC firewall in automatic mode.
	//
	// example:
	//
	// 10.0.0.0/16
	FirewallVpcCidr *string `json:"FirewallVpcCidr,omitempty" xml:"FirewallVpcCidr,omitempty"`
	// The region ID of the transit router.
	//
	// example:
	//
	// cn-shanghai
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7E53A7FB-3EB9-5E33-8E50-B8F417D1E02B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The routing mode of the VPC firewall. Valid values:
	//
	// 	- **managed**: automatic mode
	//
	// 	- **manual**: manual mode
	//
	// example:
	//
	// managed
	RouteMode      *string `json:"RouteMode,omitempty" xml:"RouteMode,omitempty"`
	TrAttachmentId *string `json:"TrAttachmentId,omitempty" xml:"TrAttachmentId,omitempty"`
	// The primary subnet CIDR block that the VPC uses to connect to the transit router in automatic mode.
	//
	// example:
	//
	// 10.0.2.0/24
	TrAttachmentMasterCidr *string `json:"TrAttachmentMasterCidr,omitempty" xml:"TrAttachmentMasterCidr,omitempty"`
	// In automatic mode, the primary availability zone of the subnet in the firewall VPC used for connecting to TR.
	//
	// example:
	//
	// cn-hangzhou-h
	TrAttachmentMasterZone *string `json:"TrAttachmentMasterZone,omitempty" xml:"TrAttachmentMasterZone,omitempty"`
	// The secondary subnet CIDR block that the VPC uses to connect to the transit router in automatic mode.
	//
	// example:
	//
	// 10.0.3.0/24
	TrAttachmentSlaveCidr *string `json:"TrAttachmentSlaveCidr,omitempty" xml:"TrAttachmentSlaveCidr,omitempty"`
	// In automatic mode, the backup availability zone for the subnet used to connect TR in the firewall VPC.
	//
	// example:
	//
	// cn-hangzhou-i
	TrAttachmentSlaveZone *string `json:"TrAttachmentSlaveZone,omitempty" xml:"TrAttachmentSlaveZone,omitempty"`
	// The ID of the transit router.
	//
	// example:
	//
	// tr-wz9y8sgug8b1xb416****
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
}

func (s DescribeTrFirewallsV2DetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTrFirewallsV2DetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTrFirewallsV2DetailResponseBody) GetCenId() *string {
	return s.CenId
}

func (s *DescribeTrFirewallsV2DetailResponseBody) GetFirewallDescription() *string {
	return s.FirewallDescription
}

func (s *DescribeTrFirewallsV2DetailResponseBody) GetFirewallEniId() *string {
	return s.FirewallEniId
}

func (s *DescribeTrFirewallsV2DetailResponseBody) GetFirewallEniVpcId() *string {
	return s.FirewallEniVpcId
}

func (s *DescribeTrFirewallsV2DetailResponseBody) GetFirewallEniVswitchId() *string {
	return s.FirewallEniVswitchId
}

func (s *DescribeTrFirewallsV2DetailResponseBody) GetFirewallId() *string {
	return s.FirewallId
}

func (s *DescribeTrFirewallsV2DetailResponseBody) GetFirewallName() *string {
	return s.FirewallName
}

func (s *DescribeTrFirewallsV2DetailResponseBody) GetFirewallStatus() *string {
	return s.FirewallStatus
}

func (s *DescribeTrFirewallsV2DetailResponseBody) GetFirewallSubnetCidr() *string {
	return s.FirewallSubnetCidr
}

func (s *DescribeTrFirewallsV2DetailResponseBody) GetFirewallSwitchStatus() *string {
	return s.FirewallSwitchStatus
}

func (s *DescribeTrFirewallsV2DetailResponseBody) GetFirewallVpcCidr() *string {
	return s.FirewallVpcCidr
}

func (s *DescribeTrFirewallsV2DetailResponseBody) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeTrFirewallsV2DetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTrFirewallsV2DetailResponseBody) GetRouteMode() *string {
	return s.RouteMode
}

func (s *DescribeTrFirewallsV2DetailResponseBody) GetTrAttachmentId() *string {
	return s.TrAttachmentId
}

func (s *DescribeTrFirewallsV2DetailResponseBody) GetTrAttachmentMasterCidr() *string {
	return s.TrAttachmentMasterCidr
}

func (s *DescribeTrFirewallsV2DetailResponseBody) GetTrAttachmentMasterZone() *string {
	return s.TrAttachmentMasterZone
}

func (s *DescribeTrFirewallsV2DetailResponseBody) GetTrAttachmentSlaveCidr() *string {
	return s.TrAttachmentSlaveCidr
}

func (s *DescribeTrFirewallsV2DetailResponseBody) GetTrAttachmentSlaveZone() *string {
	return s.TrAttachmentSlaveZone
}

func (s *DescribeTrFirewallsV2DetailResponseBody) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *DescribeTrFirewallsV2DetailResponseBody) SetCenId(v string) *DescribeTrFirewallsV2DetailResponseBody {
	s.CenId = &v
	return s
}

func (s *DescribeTrFirewallsV2DetailResponseBody) SetFirewallDescription(v string) *DescribeTrFirewallsV2DetailResponseBody {
	s.FirewallDescription = &v
	return s
}

func (s *DescribeTrFirewallsV2DetailResponseBody) SetFirewallEniId(v string) *DescribeTrFirewallsV2DetailResponseBody {
	s.FirewallEniId = &v
	return s
}

func (s *DescribeTrFirewallsV2DetailResponseBody) SetFirewallEniVpcId(v string) *DescribeTrFirewallsV2DetailResponseBody {
	s.FirewallEniVpcId = &v
	return s
}

func (s *DescribeTrFirewallsV2DetailResponseBody) SetFirewallEniVswitchId(v string) *DescribeTrFirewallsV2DetailResponseBody {
	s.FirewallEniVswitchId = &v
	return s
}

func (s *DescribeTrFirewallsV2DetailResponseBody) SetFirewallId(v string) *DescribeTrFirewallsV2DetailResponseBody {
	s.FirewallId = &v
	return s
}

func (s *DescribeTrFirewallsV2DetailResponseBody) SetFirewallName(v string) *DescribeTrFirewallsV2DetailResponseBody {
	s.FirewallName = &v
	return s
}

func (s *DescribeTrFirewallsV2DetailResponseBody) SetFirewallStatus(v string) *DescribeTrFirewallsV2DetailResponseBody {
	s.FirewallStatus = &v
	return s
}

func (s *DescribeTrFirewallsV2DetailResponseBody) SetFirewallSubnetCidr(v string) *DescribeTrFirewallsV2DetailResponseBody {
	s.FirewallSubnetCidr = &v
	return s
}

func (s *DescribeTrFirewallsV2DetailResponseBody) SetFirewallSwitchStatus(v string) *DescribeTrFirewallsV2DetailResponseBody {
	s.FirewallSwitchStatus = &v
	return s
}

func (s *DescribeTrFirewallsV2DetailResponseBody) SetFirewallVpcCidr(v string) *DescribeTrFirewallsV2DetailResponseBody {
	s.FirewallVpcCidr = &v
	return s
}

func (s *DescribeTrFirewallsV2DetailResponseBody) SetRegionNo(v string) *DescribeTrFirewallsV2DetailResponseBody {
	s.RegionNo = &v
	return s
}

func (s *DescribeTrFirewallsV2DetailResponseBody) SetRequestId(v string) *DescribeTrFirewallsV2DetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTrFirewallsV2DetailResponseBody) SetRouteMode(v string) *DescribeTrFirewallsV2DetailResponseBody {
	s.RouteMode = &v
	return s
}

func (s *DescribeTrFirewallsV2DetailResponseBody) SetTrAttachmentId(v string) *DescribeTrFirewallsV2DetailResponseBody {
	s.TrAttachmentId = &v
	return s
}

func (s *DescribeTrFirewallsV2DetailResponseBody) SetTrAttachmentMasterCidr(v string) *DescribeTrFirewallsV2DetailResponseBody {
	s.TrAttachmentMasterCidr = &v
	return s
}

func (s *DescribeTrFirewallsV2DetailResponseBody) SetTrAttachmentMasterZone(v string) *DescribeTrFirewallsV2DetailResponseBody {
	s.TrAttachmentMasterZone = &v
	return s
}

func (s *DescribeTrFirewallsV2DetailResponseBody) SetTrAttachmentSlaveCidr(v string) *DescribeTrFirewallsV2DetailResponseBody {
	s.TrAttachmentSlaveCidr = &v
	return s
}

func (s *DescribeTrFirewallsV2DetailResponseBody) SetTrAttachmentSlaveZone(v string) *DescribeTrFirewallsV2DetailResponseBody {
	s.TrAttachmentSlaveZone = &v
	return s
}

func (s *DescribeTrFirewallsV2DetailResponseBody) SetTransitRouterId(v string) *DescribeTrFirewallsV2DetailResponseBody {
	s.TransitRouterId = &v
	return s
}

func (s *DescribeTrFirewallsV2DetailResponseBody) Validate() error {
	return dara.Validate(s)
}
