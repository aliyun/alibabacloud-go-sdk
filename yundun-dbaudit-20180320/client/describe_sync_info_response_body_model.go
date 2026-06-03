// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSyncInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceInfo(v *DescribeSyncInfoResponseBodyInstanceInfo) *DescribeSyncInfoResponseBody
	GetInstanceInfo() *DescribeSyncInfoResponseBodyInstanceInfo
	SetRequestId(v string) *DescribeSyncInfoResponseBody
	GetRequestId() *string
}

type DescribeSyncInfoResponseBody struct {
	InstanceInfo *DescribeSyncInfoResponseBodyInstanceInfo `json:"InstanceInfo,omitempty" xml:"InstanceInfo,omitempty" type:"Struct"`
	RequestId    *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSyncInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSyncInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSyncInfoResponseBody) GetInstanceInfo() *DescribeSyncInfoResponseBodyInstanceInfo {
	return s.InstanceInfo
}

func (s *DescribeSyncInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSyncInfoResponseBody) SetInstanceInfo(v *DescribeSyncInfoResponseBodyInstanceInfo) *DescribeSyncInfoResponseBody {
	s.InstanceInfo = v
	return s
}

func (s *DescribeSyncInfoResponseBody) SetRequestId(v string) *DescribeSyncInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSyncInfoResponseBody) Validate() error {
	if s.InstanceInfo != nil {
		if err := s.InstanceInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSyncInfoResponseBodyInstanceInfo struct {
	AccessType          *int32  `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	Aliuid              *int64  `json:"Aliuid,omitempty" xml:"Aliuid,omitempty"`
	CustomName          *string `json:"CustomName,omitempty" xml:"CustomName,omitempty"`
	EcsEip              *string `json:"EcsEip,omitempty" xml:"EcsEip,omitempty"`
	EcsInstanceId       *string `json:"EcsInstanceId,omitempty" xml:"EcsInstanceId,omitempty"`
	EcsInternetIp       *string `json:"EcsInternetIp,omitempty" xml:"EcsInternetIp,omitempty"`
	EcsIntranetIp       *string `json:"EcsIntranetIp,omitempty" xml:"EcsIntranetIp,omitempty"`
	EcsNetworkType      *string `json:"EcsNetworkType,omitempty" xml:"EcsNetworkType,omitempty"`
	EcsStatus           *string `json:"EcsStatus,omitempty" xml:"EcsStatus,omitempty"`
	EcsUuid             *string `json:"EcsUuid,omitempty" xml:"EcsUuid,omitempty"`
	ExpireTime          *int64  `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	ImageVersionName    *string `json:"ImageVersionName,omitempty" xml:"ImageVersionName,omitempty"`
	InstanceId          *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PlanCode            *string `json:"PlanCode,omitempty" xml:"PlanCode,omitempty"`
	PlanName            *string `json:"PlanName,omitempty" xml:"PlanName,omitempty"`
	PlanUpgradeStatus   *int32  `json:"PlanUpgradeStatus,omitempty" xml:"PlanUpgradeStatus,omitempty"`
	PlanUpgradeable     *string `json:"PlanUpgradeable,omitempty" xml:"PlanUpgradeable,omitempty"`
	ProductCode         *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductName         *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	PublicAccessControl *int32  `json:"PublicAccessControl,omitempty" xml:"PublicAccessControl,omitempty"`
	RegionName          *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	RegionNo            *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	Renewable           *bool   `json:"Renewable,omitempty" xml:"Renewable,omitempty"`
	StartTime           *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status              *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	UpgradeStatus       *int32  `json:"UpgradeStatus,omitempty" xml:"UpgradeStatus,omitempty"`
	VendorCode          *string `json:"VendorCode,omitempty" xml:"VendorCode,omitempty"`
	VswitchId           *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	ZoneNo              *string `json:"ZoneNo,omitempty" xml:"ZoneNo,omitempty"`
}

func (s DescribeSyncInfoResponseBodyInstanceInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeSyncInfoResponseBodyInstanceInfo) GoString() string {
	return s.String()
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) GetAccessType() *int32 {
	return s.AccessType
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) GetAliuid() *int64 {
	return s.Aliuid
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) GetCustomName() *string {
	return s.CustomName
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) GetEcsEip() *string {
	return s.EcsEip
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) GetEcsInstanceId() *string {
	return s.EcsInstanceId
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) GetEcsInternetIp() *string {
	return s.EcsInternetIp
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) GetEcsIntranetIp() *string {
	return s.EcsIntranetIp
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) GetEcsNetworkType() *string {
	return s.EcsNetworkType
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) GetEcsStatus() *string {
	return s.EcsStatus
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) GetEcsUuid() *string {
	return s.EcsUuid
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) GetImageVersionName() *string {
	return s.ImageVersionName
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) GetPlanCode() *string {
	return s.PlanCode
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) GetPlanName() *string {
	return s.PlanName
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) GetPlanUpgradeStatus() *int32 {
	return s.PlanUpgradeStatus
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) GetPlanUpgradeable() *string {
	return s.PlanUpgradeable
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) GetProductName() *string {
	return s.ProductName
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) GetPublicAccessControl() *int32 {
	return s.PublicAccessControl
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) GetRegionName() *string {
	return s.RegionName
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) GetRenewable() *bool {
	return s.Renewable
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) GetUpgradeStatus() *int32 {
	return s.UpgradeStatus
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) GetVendorCode() *string {
	return s.VendorCode
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) GetVswitchId() *string {
	return s.VswitchId
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) GetZoneNo() *string {
	return s.ZoneNo
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) SetAccessType(v int32) *DescribeSyncInfoResponseBodyInstanceInfo {
	s.AccessType = &v
	return s
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) SetAliuid(v int64) *DescribeSyncInfoResponseBodyInstanceInfo {
	s.Aliuid = &v
	return s
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) SetCustomName(v string) *DescribeSyncInfoResponseBodyInstanceInfo {
	s.CustomName = &v
	return s
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) SetEcsEip(v string) *DescribeSyncInfoResponseBodyInstanceInfo {
	s.EcsEip = &v
	return s
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) SetEcsInstanceId(v string) *DescribeSyncInfoResponseBodyInstanceInfo {
	s.EcsInstanceId = &v
	return s
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) SetEcsInternetIp(v string) *DescribeSyncInfoResponseBodyInstanceInfo {
	s.EcsInternetIp = &v
	return s
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) SetEcsIntranetIp(v string) *DescribeSyncInfoResponseBodyInstanceInfo {
	s.EcsIntranetIp = &v
	return s
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) SetEcsNetworkType(v string) *DescribeSyncInfoResponseBodyInstanceInfo {
	s.EcsNetworkType = &v
	return s
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) SetEcsStatus(v string) *DescribeSyncInfoResponseBodyInstanceInfo {
	s.EcsStatus = &v
	return s
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) SetEcsUuid(v string) *DescribeSyncInfoResponseBodyInstanceInfo {
	s.EcsUuid = &v
	return s
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) SetExpireTime(v int64) *DescribeSyncInfoResponseBodyInstanceInfo {
	s.ExpireTime = &v
	return s
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) SetImageVersionName(v string) *DescribeSyncInfoResponseBodyInstanceInfo {
	s.ImageVersionName = &v
	return s
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) SetInstanceId(v string) *DescribeSyncInfoResponseBodyInstanceInfo {
	s.InstanceId = &v
	return s
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) SetPlanCode(v string) *DescribeSyncInfoResponseBodyInstanceInfo {
	s.PlanCode = &v
	return s
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) SetPlanName(v string) *DescribeSyncInfoResponseBodyInstanceInfo {
	s.PlanName = &v
	return s
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) SetPlanUpgradeStatus(v int32) *DescribeSyncInfoResponseBodyInstanceInfo {
	s.PlanUpgradeStatus = &v
	return s
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) SetPlanUpgradeable(v string) *DescribeSyncInfoResponseBodyInstanceInfo {
	s.PlanUpgradeable = &v
	return s
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) SetProductCode(v string) *DescribeSyncInfoResponseBodyInstanceInfo {
	s.ProductCode = &v
	return s
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) SetProductName(v string) *DescribeSyncInfoResponseBodyInstanceInfo {
	s.ProductName = &v
	return s
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) SetPublicAccessControl(v int32) *DescribeSyncInfoResponseBodyInstanceInfo {
	s.PublicAccessControl = &v
	return s
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) SetRegionName(v string) *DescribeSyncInfoResponseBodyInstanceInfo {
	s.RegionName = &v
	return s
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) SetRegionNo(v string) *DescribeSyncInfoResponseBodyInstanceInfo {
	s.RegionNo = &v
	return s
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) SetRenewable(v bool) *DescribeSyncInfoResponseBodyInstanceInfo {
	s.Renewable = &v
	return s
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) SetStartTime(v int64) *DescribeSyncInfoResponseBodyInstanceInfo {
	s.StartTime = &v
	return s
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) SetStatus(v int32) *DescribeSyncInfoResponseBodyInstanceInfo {
	s.Status = &v
	return s
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) SetUpgradeStatus(v int32) *DescribeSyncInfoResponseBodyInstanceInfo {
	s.UpgradeStatus = &v
	return s
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) SetVendorCode(v string) *DescribeSyncInfoResponseBodyInstanceInfo {
	s.VendorCode = &v
	return s
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) SetVswitchId(v string) *DescribeSyncInfoResponseBodyInstanceInfo {
	s.VswitchId = &v
	return s
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) SetZoneNo(v string) *DescribeSyncInfoResponseBodyInstanceInfo {
	s.ZoneNo = &v
	return s
}

func (s *DescribeSyncInfoResponseBodyInstanceInfo) Validate() error {
	return dara.Validate(s)
}
