// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTrFirewallsV2ListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeTrFirewallsV2ListResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *DescribeTrFirewallsV2ListResponseBody
	GetTotalCount() *string
	SetVpcTrFirewalls(v []*DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls) *DescribeTrFirewallsV2ListResponseBody
	GetVpcTrFirewalls() []*DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls
}

type DescribeTrFirewallsV2ListResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1471E2EC-F706-5F11-A79B-BD583ACB8297
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 6
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The list of VPC firewalls.
	VpcTrFirewalls []*DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls `json:"VpcTrFirewalls,omitempty" xml:"VpcTrFirewalls,omitempty" type:"Repeated"`
}

func (s DescribeTrFirewallsV2ListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTrFirewallsV2ListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTrFirewallsV2ListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTrFirewallsV2ListResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *DescribeTrFirewallsV2ListResponseBody) GetVpcTrFirewalls() []*DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls {
	return s.VpcTrFirewalls
}

func (s *DescribeTrFirewallsV2ListResponseBody) SetRequestId(v string) *DescribeTrFirewallsV2ListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTrFirewallsV2ListResponseBody) SetTotalCount(v string) *DescribeTrFirewallsV2ListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeTrFirewallsV2ListResponseBody) SetVpcTrFirewalls(v []*DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls) *DescribeTrFirewallsV2ListResponseBody {
	s.VpcTrFirewalls = v
	return s
}

func (s *DescribeTrFirewallsV2ListResponseBody) Validate() error {
	if s.VpcTrFirewalls != nil {
		for _, item := range s.VpcTrFirewalls {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls struct {
	// The mode of the access control list (ACL) engine.
	AclConfig *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsAclConfig `json:"AclConfig,omitempty" xml:"AclConfig,omitempty" type:"Struct"`
	// The ID of the CEN instance.
	//
	// example:
	//
	// cen-03f8s0z052ka3v****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The name of the CEN instance.
	//
	// example:
	//
	// cen_swas
	CenName *string `json:"CenName,omitempty" xml:"CenName,omitempty"`
	// The payer for the transit router (TR) instance that is created for the VPC firewall. Valid values:
	//
	// - **PayByCloudFirewall**: Cloud Firewall
	//
	// - **PayByCenOwner**: The account that owns the CEN instance
	//
	// example:
	//
	// PayByCenOwner
	CloudFirewallVpcOrderType *string `json:"CloudFirewallVpcOrderType,omitempty" xml:"CloudFirewallVpcOrderType,omitempty"`
	// The instance ID of the VPC firewall.
	//
	// example:
	//
	// vfw-tr-99bc4f0fc88b4d00****
	FirewallId *string `json:"FirewallId,omitempty" xml:"FirewallId,omitempty"`
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
	// The configurations of the intrusion prevention system (IPS).
	IpsConfig *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsIpsConfig `json:"IpsConfig,omitempty" xml:"IpsConfig,omitempty" type:"Struct"`
	// The ID of the Alibaba Cloud account that owns the VPC.
	//
	// example:
	//
	// 171761785151****
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Indicates whether the VPC firewall can be automatically created. Valid values:
	//
	// - **passed**: The VPC firewall can be automatically created.
	//
	// - **failed**: The VPC firewall cannot be automatically created.
	//
	// - **unknown**: The status is unknown.
	//
	// example:
	//
	// passed
	PrecheckStatus *string `json:"PrecheckStatus,omitempty" xml:"PrecheckStatus,omitempty"`
	// The list of protected resources.
	ProtectedResource *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsProtectedResource `json:"ProtectedResource,omitempty" xml:"ProtectedResource,omitempty" type:"Struct"`
	// The region ID of the transit router instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The status of the region. Valid values:
	//
	// - **enable**: The VPC firewall is available in the region.
	//
	// - **disable**: The VPC firewall is not available in the region.
	//
	// example:
	//
	// enable
	RegionStatus *string `json:"RegionStatus,omitempty" xml:"RegionStatus,omitempty"`
	// The result code of the operation to create the VPC firewall. Valid values:
	//
	// - **RegionDisable**: The VPC firewall is not supported in the region where the network instance resides. The VPC firewall cannot be created.
	//
	// - An empty string: The VPC firewall can be created for the network instance.
	//
	// example:
	//
	// RegionDisable
	ResultCode *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	// The routing mode. Valid values:
	//
	// - **managed**: automatic mode
	//
	// - **manual**: manual mode
	//
	// example:
	//
	// managed
	RouteMode *string `json:"RouteMode,omitempty" xml:"RouteMode,omitempty"`
	// The instance ID of the transit router.
	//
	// example:
	//
	// tr-2vcmhjs88nil55fvu****
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
	// The list of unprotected resources.
	UnprotectedResource *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsUnprotectedResource `json:"UnprotectedResource,omitempty" xml:"UnprotectedResource,omitempty" type:"Struct"`
	// The instance name of the VPC firewall.
	//
	// example:
	//
	// vpc-firewall-test
	VpcFirewallName *string `json:"VpcFirewallName,omitempty" xml:"VpcFirewallName,omitempty"`
}

func (s DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls) String() string {
	return dara.Prettify(s)
}

func (s DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls) GoString() string {
	return s.String()
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls) GetAclConfig() *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsAclConfig {
	return s.AclConfig
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls) GetCenId() *string {
	return s.CenId
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls) GetCenName() *string {
	return s.CenName
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls) GetCloudFirewallVpcOrderType() *string {
	return s.CloudFirewallVpcOrderType
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls) GetFirewallId() *string {
	return s.FirewallId
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls) GetFirewallSwitchStatus() *string {
	return s.FirewallSwitchStatus
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls) GetIpsConfig() *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsIpsConfig {
	return s.IpsConfig
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls) GetPrecheckStatus() *string {
	return s.PrecheckStatus
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls) GetProtectedResource() *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsProtectedResource {
	return s.ProtectedResource
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls) GetRegionStatus() *string {
	return s.RegionStatus
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls) GetResultCode() *string {
	return s.ResultCode
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls) GetRouteMode() *string {
	return s.RouteMode
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls) GetUnprotectedResource() *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsUnprotectedResource {
	return s.UnprotectedResource
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls) GetVpcFirewallName() *string {
	return s.VpcFirewallName
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls) SetAclConfig(v *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsAclConfig) *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls {
	s.AclConfig = v
	return s
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls) SetCenId(v string) *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls {
	s.CenId = &v
	return s
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls) SetCenName(v string) *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls {
	s.CenName = &v
	return s
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls) SetCloudFirewallVpcOrderType(v string) *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls {
	s.CloudFirewallVpcOrderType = &v
	return s
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls) SetFirewallId(v string) *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls {
	s.FirewallId = &v
	return s
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls) SetFirewallSwitchStatus(v string) *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls {
	s.FirewallSwitchStatus = &v
	return s
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls) SetIpsConfig(v *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsIpsConfig) *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls {
	s.IpsConfig = v
	return s
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls) SetOwnerId(v int64) *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls {
	s.OwnerId = &v
	return s
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls) SetPrecheckStatus(v string) *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls {
	s.PrecheckStatus = &v
	return s
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls) SetProtectedResource(v *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsProtectedResource) *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls {
	s.ProtectedResource = v
	return s
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls) SetRegionNo(v string) *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls {
	s.RegionNo = &v
	return s
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls) SetRegionStatus(v string) *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls {
	s.RegionStatus = &v
	return s
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls) SetResultCode(v string) *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls {
	s.ResultCode = &v
	return s
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls) SetRouteMode(v string) *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls {
	s.RouteMode = &v
	return s
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls) SetTransitRouterId(v string) *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls {
	s.TransitRouterId = &v
	return s
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls) SetUnprotectedResource(v *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsUnprotectedResource) *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls {
	s.UnprotectedResource = v
	return s
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls) SetVpcFirewallName(v string) *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls {
	s.VpcFirewallName = &v
	return s
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewalls) Validate() error {
	if s.AclConfig != nil {
		if err := s.AclConfig.Validate(); err != nil {
			return err
		}
	}
	if s.IpsConfig != nil {
		if err := s.IpsConfig.Validate(); err != nil {
			return err
		}
	}
	if s.ProtectedResource != nil {
		if err := s.ProtectedResource.Validate(); err != nil {
			return err
		}
	}
	if s.UnprotectedResource != nil {
		if err := s.UnprotectedResource.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsAclConfig struct {
	// Indicates whether the strict mode is enabled.
	//
	// - 1: enabled
	//
	// - 0: disabled
	//
	// example:
	//
	// 1
	StrictMode *int32 `json:"StrictMode,omitempty" xml:"StrictMode,omitempty"`
}

func (s DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsAclConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsAclConfig) GoString() string {
	return s.String()
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsAclConfig) GetStrictMode() *int32 {
	return s.StrictMode
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsAclConfig) SetStrictMode(v int32) *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsAclConfig {
	s.StrictMode = &v
	return s
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsAclConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsIpsConfig struct {
	// Indicates whether to enable the basic protection feature. Valid values:
	//
	// - **1**: enabled
	//
	// - **0**: disabled
	//
	// example:
	//
	// 1
	BasicRules *int32 `json:"BasicRules,omitempty" xml:"BasicRules,omitempty"`
	// Indicates whether to enable virtual patching. Valid values:
	//
	// - **1**: enabled
	//
	// - **0**: disabled
	//
	// example:
	//
	// 1
	EnableAllPatch *int32 `json:"EnableAllPatch,omitempty" xml:"EnableAllPatch,omitempty"`
	// The IPS rule group. Valid values:
	//
	// - **1**: loose
	//
	// - **2**: medium
	//
	// - **3**: strict
	//
	// example:
	//
	// 3
	RuleClass *int32 `json:"RuleClass,omitempty" xml:"RuleClass,omitempty"`
	// The IPS mode. Valid values:
	//
	// - **1**: block mode
	//
	// - **0**: monitor mode
	//
	// example:
	//
	// 1
	RunMode *int32 `json:"RunMode,omitempty" xml:"RunMode,omitempty"`
}

func (s DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsIpsConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsIpsConfig) GoString() string {
	return s.String()
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsIpsConfig) GetBasicRules() *int32 {
	return s.BasicRules
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsIpsConfig) GetEnableAllPatch() *int32 {
	return s.EnableAllPatch
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsIpsConfig) GetRuleClass() *int32 {
	return s.RuleClass
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsIpsConfig) GetRunMode() *int32 {
	return s.RunMode
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsIpsConfig) SetBasicRules(v int32) *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsIpsConfig {
	s.BasicRules = &v
	return s
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsIpsConfig) SetEnableAllPatch(v int32) *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsIpsConfig {
	s.EnableAllPatch = &v
	return s
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsIpsConfig) SetRuleClass(v int32) *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsIpsConfig {
	s.RuleClass = &v
	return s
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsIpsConfig) SetRunMode(v int32) *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsIpsConfig {
	s.RunMode = &v
	return s
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsIpsConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsProtectedResource struct {
	// The number of protected resources.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The list of protected Express Connect Router (ECR) instances.
	EcrList []*string `json:"EcrList,omitempty" xml:"EcrList,omitempty" type:"Repeated"`
	// The list of protected peer transit routers.
	PeerTrList []*string `json:"PeerTrList,omitempty" xml:"PeerTrList,omitempty" type:"Repeated"`
	// The list of protected virtual border routers (VBRs).
	VbrList []*string `json:"VbrList,omitempty" xml:"VbrList,omitempty" type:"Repeated"`
	// The list of protected VPCs.
	VpcList []*string `json:"VpcList,omitempty" xml:"VpcList,omitempty" type:"Repeated"`
	// The list of protected VPN gateways.
	VpnList []*string `json:"VpnList,omitempty" xml:"VpnList,omitempty" type:"Repeated"`
}

func (s DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsProtectedResource) String() string {
	return dara.Prettify(s)
}

func (s DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsProtectedResource) GoString() string {
	return s.String()
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsProtectedResource) GetCount() *int32 {
	return s.Count
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsProtectedResource) GetEcrList() []*string {
	return s.EcrList
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsProtectedResource) GetPeerTrList() []*string {
	return s.PeerTrList
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsProtectedResource) GetVbrList() []*string {
	return s.VbrList
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsProtectedResource) GetVpcList() []*string {
	return s.VpcList
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsProtectedResource) GetVpnList() []*string {
	return s.VpnList
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsProtectedResource) SetCount(v int32) *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsProtectedResource {
	s.Count = &v
	return s
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsProtectedResource) SetEcrList(v []*string) *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsProtectedResource {
	s.EcrList = v
	return s
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsProtectedResource) SetPeerTrList(v []*string) *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsProtectedResource {
	s.PeerTrList = v
	return s
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsProtectedResource) SetVbrList(v []*string) *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsProtectedResource {
	s.VbrList = v
	return s
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsProtectedResource) SetVpcList(v []*string) *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsProtectedResource {
	s.VpcList = v
	return s
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsProtectedResource) SetVpnList(v []*string) *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsProtectedResource {
	s.VpnList = v
	return s
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsProtectedResource) Validate() error {
	return dara.Validate(s)
}

type DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsUnprotectedResource struct {
	// The number of unprotected resources.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The list of unprotected Express Connect Router (ECR) instances.
	EcrList []*string `json:"EcrList,omitempty" xml:"EcrList,omitempty" type:"Repeated"`
	// The list of unprotected peer transit routers.
	PeerTrList []*string `json:"PeerTrList,omitempty" xml:"PeerTrList,omitempty" type:"Repeated"`
	// The list of unprotected virtual border routers (VBRs).
	VbrList []*string `json:"VbrList,omitempty" xml:"VbrList,omitempty" type:"Repeated"`
	// The list of unprotected VPCs.
	VpcList []*string `json:"VpcList,omitempty" xml:"VpcList,omitempty" type:"Repeated"`
	// The list of unprotected VPN gateways.
	VpnList []*string `json:"VpnList,omitempty" xml:"VpnList,omitempty" type:"Repeated"`
}

func (s DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsUnprotectedResource) String() string {
	return dara.Prettify(s)
}

func (s DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsUnprotectedResource) GoString() string {
	return s.String()
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsUnprotectedResource) GetCount() *int32 {
	return s.Count
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsUnprotectedResource) GetEcrList() []*string {
	return s.EcrList
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsUnprotectedResource) GetPeerTrList() []*string {
	return s.PeerTrList
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsUnprotectedResource) GetVbrList() []*string {
	return s.VbrList
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsUnprotectedResource) GetVpcList() []*string {
	return s.VpcList
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsUnprotectedResource) GetVpnList() []*string {
	return s.VpnList
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsUnprotectedResource) SetCount(v int32) *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsUnprotectedResource {
	s.Count = &v
	return s
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsUnprotectedResource) SetEcrList(v []*string) *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsUnprotectedResource {
	s.EcrList = v
	return s
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsUnprotectedResource) SetPeerTrList(v []*string) *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsUnprotectedResource {
	s.PeerTrList = v
	return s
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsUnprotectedResource) SetVbrList(v []*string) *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsUnprotectedResource {
	s.VbrList = v
	return s
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsUnprotectedResource) SetVpcList(v []*string) *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsUnprotectedResource {
	s.VpcList = v
	return s
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsUnprotectedResource) SetVpnList(v []*string) *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsUnprotectedResource {
	s.VpnList = v
	return s
}

func (s *DescribeTrFirewallsV2ListResponseBodyVpcTrFirewallsUnprotectedResource) Validate() error {
	return dara.Validate(s)
}
