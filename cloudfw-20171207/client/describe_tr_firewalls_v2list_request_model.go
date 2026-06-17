// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTrFirewallsV2ListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCenId(v string) *DescribeTrFirewallsV2ListRequest
	GetCenId() *string
	SetCurrentPage(v int32) *DescribeTrFirewallsV2ListRequest
	GetCurrentPage() *int32
	SetFirewallId(v string) *DescribeTrFirewallsV2ListRequest
	GetFirewallId() *string
	SetFirewallName(v string) *DescribeTrFirewallsV2ListRequest
	GetFirewallName() *string
	SetFirewallSwitchStatus(v string) *DescribeTrFirewallsV2ListRequest
	GetFirewallSwitchStatus() *string
	SetLang(v string) *DescribeTrFirewallsV2ListRequest
	GetLang() *string
	SetOwnerId(v string) *DescribeTrFirewallsV2ListRequest
	GetOwnerId() *string
	SetPageSize(v int32) *DescribeTrFirewallsV2ListRequest
	GetPageSize() *int32
	SetRegionNo(v string) *DescribeTrFirewallsV2ListRequest
	GetRegionNo() *string
	SetRouteMode(v string) *DescribeTrFirewallsV2ListRequest
	GetRouteMode() *string
	SetTransitRouterId(v string) *DescribeTrFirewallsV2ListRequest
	GetTransitRouterId() *string
}

type DescribeTrFirewallsV2ListRequest struct {
	// The ID of the Cloud Enterprise Network (CEN) instance.
	//
	// example:
	//
	// cen-rig0t5zi96crkl****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The page number for a paged query. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The instance ID of the VPC firewall.
	//
	// example:
	//
	// vfw-tr-f1799baa9e254651****
	FirewallId *string `json:"FirewallId,omitempty" xml:"FirewallId,omitempty"`
	// The instance name of the VPC firewall.
	//
	// example:
	//
	// vpc-firewall-test
	FirewallName *string `json:"FirewallName,omitempty" xml:"FirewallName,omitempty"`
	// The status of the VPC firewall. Valid values:
	//
	// - **opened**: Enabled
	//
	// - **closed**: Disabled
	//
	// - **notconfigured**: The VPC firewall is not configured.
	//
	// - **configured**: The VPC firewall is configured.
	//
	// - **creating**: The VPC firewall is being created.
	//
	// - **opening**: The VPC firewall is being enabled.
	//
	// - **deleting**: The VPC firewall is being deleted.
	//
	// > If you do not specify this parameter, VPC firewalls in all states are queried.
	//
	// example:
	//
	// opened
	FirewallSwitchStatus *string `json:"FirewallSwitchStatus,omitempty" xml:"FirewallSwitchStatus,omitempty"`
	// The language of the response message. Valid values:
	//
	// - **zh*	- (default): Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang    *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The maximum number of entries to return on each page in a paged query. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the transit router instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The routing mode. Valid values:
	//
	// - **managed**: automatic mode.
	//
	// - **manual**: manual mode.
	//
	// > If you do not specify this parameter, VPC firewalls in all routing modes are queried.
	//
	// example:
	//
	// managed
	RouteMode *string `json:"RouteMode,omitempty" xml:"RouteMode,omitempty"`
	// The instance ID of the transit router.
	//
	// example:
	//
	// tr-uf6egtvyaedvt20xl****
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
}

func (s DescribeTrFirewallsV2ListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTrFirewallsV2ListRequest) GoString() string {
	return s.String()
}

func (s *DescribeTrFirewallsV2ListRequest) GetCenId() *string {
	return s.CenId
}

func (s *DescribeTrFirewallsV2ListRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeTrFirewallsV2ListRequest) GetFirewallId() *string {
	return s.FirewallId
}

func (s *DescribeTrFirewallsV2ListRequest) GetFirewallName() *string {
	return s.FirewallName
}

func (s *DescribeTrFirewallsV2ListRequest) GetFirewallSwitchStatus() *string {
	return s.FirewallSwitchStatus
}

func (s *DescribeTrFirewallsV2ListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeTrFirewallsV2ListRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DescribeTrFirewallsV2ListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeTrFirewallsV2ListRequest) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeTrFirewallsV2ListRequest) GetRouteMode() *string {
	return s.RouteMode
}

func (s *DescribeTrFirewallsV2ListRequest) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *DescribeTrFirewallsV2ListRequest) SetCenId(v string) *DescribeTrFirewallsV2ListRequest {
	s.CenId = &v
	return s
}

func (s *DescribeTrFirewallsV2ListRequest) SetCurrentPage(v int32) *DescribeTrFirewallsV2ListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeTrFirewallsV2ListRequest) SetFirewallId(v string) *DescribeTrFirewallsV2ListRequest {
	s.FirewallId = &v
	return s
}

func (s *DescribeTrFirewallsV2ListRequest) SetFirewallName(v string) *DescribeTrFirewallsV2ListRequest {
	s.FirewallName = &v
	return s
}

func (s *DescribeTrFirewallsV2ListRequest) SetFirewallSwitchStatus(v string) *DescribeTrFirewallsV2ListRequest {
	s.FirewallSwitchStatus = &v
	return s
}

func (s *DescribeTrFirewallsV2ListRequest) SetLang(v string) *DescribeTrFirewallsV2ListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeTrFirewallsV2ListRequest) SetOwnerId(v string) *DescribeTrFirewallsV2ListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeTrFirewallsV2ListRequest) SetPageSize(v int32) *DescribeTrFirewallsV2ListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTrFirewallsV2ListRequest) SetRegionNo(v string) *DescribeTrFirewallsV2ListRequest {
	s.RegionNo = &v
	return s
}

func (s *DescribeTrFirewallsV2ListRequest) SetRouteMode(v string) *DescribeTrFirewallsV2ListRequest {
	s.RouteMode = &v
	return s
}

func (s *DescribeTrFirewallsV2ListRequest) SetTransitRouterId(v string) *DescribeTrFirewallsV2ListRequest {
	s.TransitRouterId = &v
	return s
}

func (s *DescribeTrFirewallsV2ListRequest) Validate() error {
	return dara.Validate(s)
}
