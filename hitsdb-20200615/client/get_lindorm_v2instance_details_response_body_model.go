// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLindormV2InstanceDetailsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAliUid(v int64) *GetLindormV2InstanceDetailsResponseBody
	GetAliUid() *int64
	SetArbiterVSwitchId(v string) *GetLindormV2InstanceDetailsResponseBody
	GetArbiterVSwitchId() *string
	SetArbiterZoneId(v string) *GetLindormV2InstanceDetailsResponseBody
	GetArbiterZoneId() *string
	SetAutoRenew(v bool) *GetLindormV2InstanceDetailsResponseBody
	GetAutoRenew() *bool
	SetCloudStorageSize(v int64) *GetLindormV2InstanceDetailsResponseBody
	GetCloudStorageSize() *int64
	SetColdStorage(v int32) *GetLindormV2InstanceDetailsResponseBody
	GetColdStorage() *int32
	SetCreateMilliseconds(v int64) *GetLindormV2InstanceDetailsResponseBody
	GetCreateMilliseconds() *int64
	SetDeletionProtection(v string) *GetLindormV2InstanceDetailsResponseBody
	GetDeletionProtection() *string
	SetDiskCategory(v string) *GetLindormV2InstanceDetailsResponseBody
	GetDiskCategory() *string
	SetDiskThreshold(v string) *GetLindormV2InstanceDetailsResponseBody
	GetDiskThreshold() *string
	SetDiskUsage(v string) *GetLindormV2InstanceDetailsResponseBody
	GetDiskUsage() *string
	SetEnableCompute(v bool) *GetLindormV2InstanceDetailsResponseBody
	GetEnableCompute() *bool
	SetEnableFs(v string) *GetLindormV2InstanceDetailsResponseBody
	GetEnableFs() *string
	SetEnableStoreTDE(v string) *GetLindormV2InstanceDetailsResponseBody
	GetEnableStoreTDE() *string
	SetEngineList(v []*GetLindormV2InstanceDetailsResponseBodyEngineList) *GetLindormV2InstanceDetailsResponseBody
	GetEngineList() []*GetLindormV2InstanceDetailsResponseBodyEngineList
	SetExpiredMilliseconds(v int64) *GetLindormV2InstanceDetailsResponseBody
	GetExpiredMilliseconds() *int64
	SetInitialRootPassword(v string) *GetLindormV2InstanceDetailsResponseBody
	GetInitialRootPassword() *string
	SetInstanceAlias(v string) *GetLindormV2InstanceDetailsResponseBody
	GetInstanceAlias() *string
	SetInstanceId(v string) *GetLindormV2InstanceDetailsResponseBody
	GetInstanceId() *string
	SetInstanceStatus(v string) *GetLindormV2InstanceDetailsResponseBody
	GetInstanceStatus() *string
	SetInstanceType(v string) *GetLindormV2InstanceDetailsResponseBody
	GetInstanceType() *string
	SetMaintainEndTime(v string) *GetLindormV2InstanceDetailsResponseBody
	GetMaintainEndTime() *string
	SetMaintainStartTime(v string) *GetLindormV2InstanceDetailsResponseBody
	GetMaintainStartTime() *string
	SetNetworkType(v string) *GetLindormV2InstanceDetailsResponseBody
	GetNetworkType() *string
	SetPayType(v string) *GetLindormV2InstanceDetailsResponseBody
	GetPayType() *string
	SetPrimaryVSwitchId(v string) *GetLindormV2InstanceDetailsResponseBody
	GetPrimaryVSwitchId() *string
	SetPrimaryZoneId(v string) *GetLindormV2InstanceDetailsResponseBody
	GetPrimaryZoneId() *string
	SetRegionId(v string) *GetLindormV2InstanceDetailsResponseBody
	GetRegionId() *string
	SetRequestId(v string) *GetLindormV2InstanceDetailsResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *GetLindormV2InstanceDetailsResponseBody
	GetResourceGroupId() *string
	SetServiceType(v string) *GetLindormV2InstanceDetailsResponseBody
	GetServiceType() *string
	SetStandbyVSwitchId(v string) *GetLindormV2InstanceDetailsResponseBody
	GetStandbyVSwitchId() *string
	SetStandbyZoneId(v string) *GetLindormV2InstanceDetailsResponseBody
	GetStandbyZoneId() *string
	SetStorageUsage(v *GetLindormV2InstanceDetailsResponseBodyStorageUsage) *GetLindormV2InstanceDetailsResponseBody
	GetStorageUsage() *GetLindormV2InstanceDetailsResponseBodyStorageUsage
	SetVpcId(v string) *GetLindormV2InstanceDetailsResponseBody
	GetVpcId() *string
	SetVswitchId(v string) *GetLindormV2InstanceDetailsResponseBody
	GetVswitchId() *string
	SetWhiteIpList(v []*GetLindormV2InstanceDetailsResponseBodyWhiteIpList) *GetLindormV2InstanceDetailsResponseBody
	GetWhiteIpList() []*GetLindormV2InstanceDetailsResponseBodyWhiteIpList
	SetZoneEngineInfoMap(v map[string]interface{}) *GetLindormV2InstanceDetailsResponseBody
	GetZoneEngineInfoMap() map[string]interface{}
	SetZoneId(v string) *GetLindormV2InstanceDetailsResponseBody
	GetZoneId() *string
}

type GetLindormV2InstanceDetailsResponseBody struct {
	// The 16-digit AliUid of the Alibaba Cloud account.
	//
	// example:
	//
	// 164901546557****
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// For a multi-zone instance, this is the ID of the virtual switch in the arbiter zone. The vSwitch must be in the zone specified by ArbiterZoneId.
	//
	// example:
	//
	// vsw-uf6664pqjawb87k36****
	ArbiterVSwitchId *string `json:"ArbiterVSwitchId,omitempty" xml:"ArbiterVSwitchId,omitempty"`
	// For a multi-zone instance, this is the ID of the arbiter zone.
	//
	// example:
	//
	// cn-shanghai-g
	ArbiterZoneId *string `json:"ArbiterZoneId,omitempty" xml:"ArbiterZoneId,omitempty"`
	// Indicates whether auto-renewal is enabled. Valid values:
	//
	// - **true**: enabled.
	//
	// - **false**: disabled.
	//
	// > This parameter is returned only for subscription instances.
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The cloud storage capacity in GB.
	//
	// example:
	//
	// 480
	CloudStorageSize *int64 `json:"CloudStorageSize,omitempty" xml:"CloudStorageSize,omitempty"`
	// The capacity of the storage-optimized cloud storage.
	//
	// example:
	//
	// 800
	ColdStorage *int32 `json:"ColdStorage,omitempty" xml:"ColdStorage,omitempty"`
	// The number of milliseconds between the instance creation time and 00:00:00 on January 1, 1970.
	//
	// example:
	//
	// 1627290664000
	CreateMilliseconds *int64 `json:"CreateMilliseconds,omitempty" xml:"CreateMilliseconds,omitempty"`
	// Indicates whether deletion protection is enabled. Valid values:
	//
	// - **true**: enabled.
	//
	// - **false**: disabled.
	//
	// example:
	//
	// false
	DeletionProtection *string `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// The storage class. Valid values:
	//
	// - **StandardStorage**: Standard cloud storage.
	//
	// - **PerformanceStorage**: Performance cloud storage.
	//
	// example:
	//
	// PerformanceStorage
	DiskCategory *string `json:"DiskCategory,omitempty" xml:"DiskCategory,omitempty"`
	// The disk space threshold.
	//
	// example:
	//
	// 80%
	DiskThreshold *string `json:"DiskThreshold,omitempty" xml:"DiskThreshold,omitempty"`
	// The disk space usage.
	//
	// example:
	//
	// 0.0%
	DiskUsage *string `json:"DiskUsage,omitempty" xml:"DiskUsage,omitempty"`
	// Indicates whether the compute engine is enabled for the instance. Valid values:
	//
	// - **true**: enabled.
	//
	// - **false**: disabled.
	//
	// example:
	//
	// true
	EnableCompute *bool `json:"EnableCompute,omitempty" xml:"EnableCompute,omitempty"`
	// Indicates whether the file engine is enabled.
	//
	// example:
	//
	// false
	EnableFs *string `json:"EnableFs,omitempty" xml:"EnableFs,omitempty"`
	// Indicates whether transparent data encryption (TDE) is enabled for storage.
	//
	// example:
	//
	// false
	EnableStoreTDE *string `json:"EnableStoreTDE,omitempty" xml:"EnableStoreTDE,omitempty"`
	// The list of engine information.
	EngineList []*GetLindormV2InstanceDetailsResponseBodyEngineList `json:"EngineList,omitempty" xml:"EngineList,omitempty" type:"Repeated"`
	// The number of milliseconds between the instance expiration time and 00:00:00 on January 1, 1970.
	//
	// example:
	//
	// 1629993600000
	ExpiredMilliseconds *int64 `json:"ExpiredMilliseconds,omitempty" xml:"ExpiredMilliseconds,omitempty"`
	// The default initial password.
	//
	// example:
	//
	// *****
	InitialRootPassword *string `json:"InitialRootPassword,omitempty" xml:"InitialRootPassword,omitempty"`
	// The instance name.
	//
	// example:
	//
	// lindorm-test
	InstanceAlias *string `json:"InstanceAlias,omitempty" xml:"InstanceAlias,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// ld-bp1mq0tdzbx1m****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance status. Valid values:
	//
	// - **CREATING**: The instance is being created.
	//
	// - **ACTIVATION**: The instance is running.
	//
	// - **COLD_EXPANDING**: The storage-optimized cloud storage is being scaled out.
	//
	// - **MINOR_VERSION_TRANSING**: The minor version is being upgraded.
	//
	// - **RESIZING**: The nodes are being scaled out.
	//
	// - **SHRINKING**: The nodes are being scaled in.
	//
	// - **CLASS_CHANGING**: The instance specifications are being upgraded or downgraded.
	//
	// - **SSL_SWITCHING**: The SSL certificate is being changed.
	//
	// - **CDC_OPENING**: The data subscription feature is being enabled.
	//
	// - **TRANSFER**: Data is being migrated.
	//
	// - **DATABASE_TRANSFER**: Data is being migrated to the database.
	//
	// - **GUARD_CREATING**: A disaster recovery instance is being created.
	//
	// - **BACKUP_RECOVERING**: Data is being restored from a backup.
	//
	// - **DATABASE_IMPORTING**: Data is being imported.
	//
	// - **NET_MODIFYING**: The network is being modified.
	//
	// - **NET_SWITCHING**: The network is being switched between the private network and the public network.
	//
	// - **NET_CREATING**: A network connection is being created.
	//
	// - **NET_DELETING**: A network connection is being deleted.
	//
	// - **DELETING**: The instance is being deleted.
	//
	// - **RESTARTING**: The instance is being restarted.
	//
	// - **LOCKED**: The instance has expired and is locked.
	//
	// example:
	//
	// ACTIVATION
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// The instance type. Valid value:
	//
	// - basic: Production.
	//
	// example:
	//
	// basic
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The end time of the maintenance window.
	//
	// example:
	//
	// 20:00Z
	MaintainEndTime *string `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	// The start time of the maintenance window.
	//
	// example:
	//
	// 00:00Z
	MaintainStartTime *string `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	// The network type of the instance.
	//
	// example:
	//
	// VPC
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// - **PREPAY**: subscription.
	//
	// - **POSTPAY**: pay-as-you-go.
	//
	// example:
	//
	// POSTPAY
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// For a multi-zone instance, this is the ID of the virtual switch in the primary zone. The vSwitch must be in the zone specified by PrimaryZoneId.
	//
	// example:
	//
	// vsw-uf6fdqa7c0pipnqzq****
	PrimaryVSwitchId *string `json:"PrimaryVSwitchId,omitempty" xml:"PrimaryVSwitchId,omitempty"`
	// For a multi-zone instance, this is the ID of the primary zone.
	//
	// example:
	//
	// cn-shanghai-e
	PrimaryZoneId *string `json:"PrimaryZoneId,omitempty" xml:"PrimaryZoneId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1556DCB0-043A-4444-8BD9-CF4A68E7EE64
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aek2i6weeb4nfii
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The instance type. Valid values:
	//
	// - **lindorm_v2**: Lindorm V2 single-zone instance.
	//
	// - **lindorm_v2_multizone**: Lindorm V2 Multi-zone Deployment (Basic) instance.
	//
	// - **lindorm_v2_multizone_ha**: Lindorm V2 multi-zone High-availability Edition instance.
	//
	// example:
	//
	// lindorm_v2
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// For a multi-zone instance, this is the ID of the virtual switch in the secondary zone. The vSwitch must be in the zone specified by StandbyZoneId.
	//
	// example:
	//
	// vsw-2zec0kcn08cgdtr6****
	StandbyVSwitchId *string `json:"StandbyVSwitchId,omitempty" xml:"StandbyVSwitchId,omitempty"`
	// For a multi-zone instance, this is the ID of the secondary zone.
	//
	// example:
	//
	// cn-shanghai-f
	StandbyZoneId *string `json:"StandbyZoneId,omitempty" xml:"StandbyZoneId,omitempty"`
	// The instance storage usage.
	StorageUsage *GetLindormV2InstanceDetailsResponseBodyStorageUsage `json:"StorageUsage,omitempty" xml:"StorageUsage,omitempty" type:"Struct"`
	// The ID of the virtual private cloud (VPC) where the instance resides.
	//
	// example:
	//
	// vpc-bp1xxxxxxxxxxxxxxxxxx
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the virtual switch.
	//
	// example:
	//
	// vsw-bp1xxxxxxxxxxxxxxxxxx
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	// The instance access whitelist.
	WhiteIpList []*GetLindormV2InstanceDetailsResponseBodyWhiteIpList `json:"WhiteIpList,omitempty" xml:"WhiteIpList,omitempty" type:"Repeated"`
	// The deployment details of the engine in the zone.
	//
	// example:
	//
	// ZoneEngineInfoMap
	ZoneEngineInfoMap map[string]interface{} `json:"ZoneEngineInfoMap,omitempty" xml:"ZoneEngineInfoMap,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetLindormV2InstanceDetailsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLindormV2InstanceDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *GetLindormV2InstanceDetailsResponseBody) GetAliUid() *int64 {
	return s.AliUid
}

func (s *GetLindormV2InstanceDetailsResponseBody) GetArbiterVSwitchId() *string {
	return s.ArbiterVSwitchId
}

func (s *GetLindormV2InstanceDetailsResponseBody) GetArbiterZoneId() *string {
	return s.ArbiterZoneId
}

func (s *GetLindormV2InstanceDetailsResponseBody) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *GetLindormV2InstanceDetailsResponseBody) GetCloudStorageSize() *int64 {
	return s.CloudStorageSize
}

func (s *GetLindormV2InstanceDetailsResponseBody) GetColdStorage() *int32 {
	return s.ColdStorage
}

func (s *GetLindormV2InstanceDetailsResponseBody) GetCreateMilliseconds() *int64 {
	return s.CreateMilliseconds
}

func (s *GetLindormV2InstanceDetailsResponseBody) GetDeletionProtection() *string {
	return s.DeletionProtection
}

func (s *GetLindormV2InstanceDetailsResponseBody) GetDiskCategory() *string {
	return s.DiskCategory
}

func (s *GetLindormV2InstanceDetailsResponseBody) GetDiskThreshold() *string {
	return s.DiskThreshold
}

func (s *GetLindormV2InstanceDetailsResponseBody) GetDiskUsage() *string {
	return s.DiskUsage
}

func (s *GetLindormV2InstanceDetailsResponseBody) GetEnableCompute() *bool {
	return s.EnableCompute
}

func (s *GetLindormV2InstanceDetailsResponseBody) GetEnableFs() *string {
	return s.EnableFs
}

func (s *GetLindormV2InstanceDetailsResponseBody) GetEnableStoreTDE() *string {
	return s.EnableStoreTDE
}

func (s *GetLindormV2InstanceDetailsResponseBody) GetEngineList() []*GetLindormV2InstanceDetailsResponseBodyEngineList {
	return s.EngineList
}

func (s *GetLindormV2InstanceDetailsResponseBody) GetExpiredMilliseconds() *int64 {
	return s.ExpiredMilliseconds
}

func (s *GetLindormV2InstanceDetailsResponseBody) GetInitialRootPassword() *string {
	return s.InitialRootPassword
}

func (s *GetLindormV2InstanceDetailsResponseBody) GetInstanceAlias() *string {
	return s.InstanceAlias
}

func (s *GetLindormV2InstanceDetailsResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetLindormV2InstanceDetailsResponseBody) GetInstanceStatus() *string {
	return s.InstanceStatus
}

func (s *GetLindormV2InstanceDetailsResponseBody) GetInstanceType() *string {
	return s.InstanceType
}

func (s *GetLindormV2InstanceDetailsResponseBody) GetMaintainEndTime() *string {
	return s.MaintainEndTime
}

func (s *GetLindormV2InstanceDetailsResponseBody) GetMaintainStartTime() *string {
	return s.MaintainStartTime
}

func (s *GetLindormV2InstanceDetailsResponseBody) GetNetworkType() *string {
	return s.NetworkType
}

func (s *GetLindormV2InstanceDetailsResponseBody) GetPayType() *string {
	return s.PayType
}

func (s *GetLindormV2InstanceDetailsResponseBody) GetPrimaryVSwitchId() *string {
	return s.PrimaryVSwitchId
}

func (s *GetLindormV2InstanceDetailsResponseBody) GetPrimaryZoneId() *string {
	return s.PrimaryZoneId
}

func (s *GetLindormV2InstanceDetailsResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *GetLindormV2InstanceDetailsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLindormV2InstanceDetailsResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetLindormV2InstanceDetailsResponseBody) GetServiceType() *string {
	return s.ServiceType
}

func (s *GetLindormV2InstanceDetailsResponseBody) GetStandbyVSwitchId() *string {
	return s.StandbyVSwitchId
}

func (s *GetLindormV2InstanceDetailsResponseBody) GetStandbyZoneId() *string {
	return s.StandbyZoneId
}

func (s *GetLindormV2InstanceDetailsResponseBody) GetStorageUsage() *GetLindormV2InstanceDetailsResponseBodyStorageUsage {
	return s.StorageUsage
}

func (s *GetLindormV2InstanceDetailsResponseBody) GetVpcId() *string {
	return s.VpcId
}

func (s *GetLindormV2InstanceDetailsResponseBody) GetVswitchId() *string {
	return s.VswitchId
}

func (s *GetLindormV2InstanceDetailsResponseBody) GetWhiteIpList() []*GetLindormV2InstanceDetailsResponseBodyWhiteIpList {
	return s.WhiteIpList
}

func (s *GetLindormV2InstanceDetailsResponseBody) GetZoneEngineInfoMap() map[string]interface{} {
	return s.ZoneEngineInfoMap
}

func (s *GetLindormV2InstanceDetailsResponseBody) GetZoneId() *string {
	return s.ZoneId
}

func (s *GetLindormV2InstanceDetailsResponseBody) SetAliUid(v int64) *GetLindormV2InstanceDetailsResponseBody {
	s.AliUid = &v
	return s
}

func (s *GetLindormV2InstanceDetailsResponseBody) SetArbiterVSwitchId(v string) *GetLindormV2InstanceDetailsResponseBody {
	s.ArbiterVSwitchId = &v
	return s
}

func (s *GetLindormV2InstanceDetailsResponseBody) SetArbiterZoneId(v string) *GetLindormV2InstanceDetailsResponseBody {
	s.ArbiterZoneId = &v
	return s
}

func (s *GetLindormV2InstanceDetailsResponseBody) SetAutoRenew(v bool) *GetLindormV2InstanceDetailsResponseBody {
	s.AutoRenew = &v
	return s
}

func (s *GetLindormV2InstanceDetailsResponseBody) SetCloudStorageSize(v int64) *GetLindormV2InstanceDetailsResponseBody {
	s.CloudStorageSize = &v
	return s
}

func (s *GetLindormV2InstanceDetailsResponseBody) SetColdStorage(v int32) *GetLindormV2InstanceDetailsResponseBody {
	s.ColdStorage = &v
	return s
}

func (s *GetLindormV2InstanceDetailsResponseBody) SetCreateMilliseconds(v int64) *GetLindormV2InstanceDetailsResponseBody {
	s.CreateMilliseconds = &v
	return s
}

func (s *GetLindormV2InstanceDetailsResponseBody) SetDeletionProtection(v string) *GetLindormV2InstanceDetailsResponseBody {
	s.DeletionProtection = &v
	return s
}

func (s *GetLindormV2InstanceDetailsResponseBody) SetDiskCategory(v string) *GetLindormV2InstanceDetailsResponseBody {
	s.DiskCategory = &v
	return s
}

func (s *GetLindormV2InstanceDetailsResponseBody) SetDiskThreshold(v string) *GetLindormV2InstanceDetailsResponseBody {
	s.DiskThreshold = &v
	return s
}

func (s *GetLindormV2InstanceDetailsResponseBody) SetDiskUsage(v string) *GetLindormV2InstanceDetailsResponseBody {
	s.DiskUsage = &v
	return s
}

func (s *GetLindormV2InstanceDetailsResponseBody) SetEnableCompute(v bool) *GetLindormV2InstanceDetailsResponseBody {
	s.EnableCompute = &v
	return s
}

func (s *GetLindormV2InstanceDetailsResponseBody) SetEnableFs(v string) *GetLindormV2InstanceDetailsResponseBody {
	s.EnableFs = &v
	return s
}

func (s *GetLindormV2InstanceDetailsResponseBody) SetEnableStoreTDE(v string) *GetLindormV2InstanceDetailsResponseBody {
	s.EnableStoreTDE = &v
	return s
}

func (s *GetLindormV2InstanceDetailsResponseBody) SetEngineList(v []*GetLindormV2InstanceDetailsResponseBodyEngineList) *GetLindormV2InstanceDetailsResponseBody {
	s.EngineList = v
	return s
}

func (s *GetLindormV2InstanceDetailsResponseBody) SetExpiredMilliseconds(v int64) *GetLindormV2InstanceDetailsResponseBody {
	s.ExpiredMilliseconds = &v
	return s
}

func (s *GetLindormV2InstanceDetailsResponseBody) SetInitialRootPassword(v string) *GetLindormV2InstanceDetailsResponseBody {
	s.InitialRootPassword = &v
	return s
}

func (s *GetLindormV2InstanceDetailsResponseBody) SetInstanceAlias(v string) *GetLindormV2InstanceDetailsResponseBody {
	s.InstanceAlias = &v
	return s
}

func (s *GetLindormV2InstanceDetailsResponseBody) SetInstanceId(v string) *GetLindormV2InstanceDetailsResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetLindormV2InstanceDetailsResponseBody) SetInstanceStatus(v string) *GetLindormV2InstanceDetailsResponseBody {
	s.InstanceStatus = &v
	return s
}

func (s *GetLindormV2InstanceDetailsResponseBody) SetInstanceType(v string) *GetLindormV2InstanceDetailsResponseBody {
	s.InstanceType = &v
	return s
}

func (s *GetLindormV2InstanceDetailsResponseBody) SetMaintainEndTime(v string) *GetLindormV2InstanceDetailsResponseBody {
	s.MaintainEndTime = &v
	return s
}

func (s *GetLindormV2InstanceDetailsResponseBody) SetMaintainStartTime(v string) *GetLindormV2InstanceDetailsResponseBody {
	s.MaintainStartTime = &v
	return s
}

func (s *GetLindormV2InstanceDetailsResponseBody) SetNetworkType(v string) *GetLindormV2InstanceDetailsResponseBody {
	s.NetworkType = &v
	return s
}

func (s *GetLindormV2InstanceDetailsResponseBody) SetPayType(v string) *GetLindormV2InstanceDetailsResponseBody {
	s.PayType = &v
	return s
}

func (s *GetLindormV2InstanceDetailsResponseBody) SetPrimaryVSwitchId(v string) *GetLindormV2InstanceDetailsResponseBody {
	s.PrimaryVSwitchId = &v
	return s
}

func (s *GetLindormV2InstanceDetailsResponseBody) SetPrimaryZoneId(v string) *GetLindormV2InstanceDetailsResponseBody {
	s.PrimaryZoneId = &v
	return s
}

func (s *GetLindormV2InstanceDetailsResponseBody) SetRegionId(v string) *GetLindormV2InstanceDetailsResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetLindormV2InstanceDetailsResponseBody) SetRequestId(v string) *GetLindormV2InstanceDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLindormV2InstanceDetailsResponseBody) SetResourceGroupId(v string) *GetLindormV2InstanceDetailsResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *GetLindormV2InstanceDetailsResponseBody) SetServiceType(v string) *GetLindormV2InstanceDetailsResponseBody {
	s.ServiceType = &v
	return s
}

func (s *GetLindormV2InstanceDetailsResponseBody) SetStandbyVSwitchId(v string) *GetLindormV2InstanceDetailsResponseBody {
	s.StandbyVSwitchId = &v
	return s
}

func (s *GetLindormV2InstanceDetailsResponseBody) SetStandbyZoneId(v string) *GetLindormV2InstanceDetailsResponseBody {
	s.StandbyZoneId = &v
	return s
}

func (s *GetLindormV2InstanceDetailsResponseBody) SetStorageUsage(v *GetLindormV2InstanceDetailsResponseBodyStorageUsage) *GetLindormV2InstanceDetailsResponseBody {
	s.StorageUsage = v
	return s
}

func (s *GetLindormV2InstanceDetailsResponseBody) SetVpcId(v string) *GetLindormV2InstanceDetailsResponseBody {
	s.VpcId = &v
	return s
}

func (s *GetLindormV2InstanceDetailsResponseBody) SetVswitchId(v string) *GetLindormV2InstanceDetailsResponseBody {
	s.VswitchId = &v
	return s
}

func (s *GetLindormV2InstanceDetailsResponseBody) SetWhiteIpList(v []*GetLindormV2InstanceDetailsResponseBodyWhiteIpList) *GetLindormV2InstanceDetailsResponseBody {
	s.WhiteIpList = v
	return s
}

func (s *GetLindormV2InstanceDetailsResponseBody) SetZoneEngineInfoMap(v map[string]interface{}) *GetLindormV2InstanceDetailsResponseBody {
	s.ZoneEngineInfoMap = v
	return s
}

func (s *GetLindormV2InstanceDetailsResponseBody) SetZoneId(v string) *GetLindormV2InstanceDetailsResponseBody {
	s.ZoneId = &v
	return s
}

func (s *GetLindormV2InstanceDetailsResponseBody) Validate() error {
	if s.EngineList != nil {
		for _, item := range s.EngineList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.StorageUsage != nil {
		if err := s.StorageUsage.Validate(); err != nil {
			return err
		}
	}
	if s.WhiteIpList != nil {
		for _, item := range s.WhiteIpList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetLindormV2InstanceDetailsResponseBodyEngineList struct {
	// A list of engine endpoints.
	ConnectAddressList []*GetLindormV2InstanceDetailsResponseBodyEngineListConnectAddressList `json:"ConnectAddressList,omitempty" xml:"ConnectAddressList,omitempty" type:"Repeated"`
	// The engine type. Valid values:
	//
	// - **TABLE**: LindormTable.
	//
	// - **TSDB**: LindormTSDB.
	//
	// - **LSEARCH**: Search engine.
	//
	// - **LTS**: LTS engine.
	//
	// - **LVECTOR**: Vector engine.
	//
	// - **LCOLUMN**: Column store.
	//
	// example:
	//
	// TABLE
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// Indicates whether the engine is the latest version. Valid values:
	//
	// - **true**: The engine is the latest version.
	//
	// - **false**: The engine is not the latest version.
	//
	// example:
	//
	// false
	IsLastVersion *bool `json:"IsLastVersion,omitempty" xml:"IsLastVersion,omitempty"`
	// The latest version number for the engine type.
	//
	// example:
	//
	// 2.2.19.2
	LatestVersion *string `json:"LatestVersion,omitempty" xml:"LatestVersion,omitempty"`
	// A list of engine node groups.
	NodeGroup []*GetLindormV2InstanceDetailsResponseBodyEngineListNodeGroup `json:"NodeGroup,omitempty" xml:"NodeGroup,omitempty" type:"Repeated"`
	// The version number of the engine type.
	//
	// example:
	//
	// 2.2.3
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetLindormV2InstanceDetailsResponseBodyEngineList) String() string {
	return dara.Prettify(s)
}

func (s GetLindormV2InstanceDetailsResponseBodyEngineList) GoString() string {
	return s.String()
}

func (s *GetLindormV2InstanceDetailsResponseBodyEngineList) GetConnectAddressList() []*GetLindormV2InstanceDetailsResponseBodyEngineListConnectAddressList {
	return s.ConnectAddressList
}

func (s *GetLindormV2InstanceDetailsResponseBodyEngineList) GetEngine() *string {
	return s.Engine
}

func (s *GetLindormV2InstanceDetailsResponseBodyEngineList) GetIsLastVersion() *bool {
	return s.IsLastVersion
}

func (s *GetLindormV2InstanceDetailsResponseBodyEngineList) GetLatestVersion() *string {
	return s.LatestVersion
}

func (s *GetLindormV2InstanceDetailsResponseBodyEngineList) GetNodeGroup() []*GetLindormV2InstanceDetailsResponseBodyEngineListNodeGroup {
	return s.NodeGroup
}

func (s *GetLindormV2InstanceDetailsResponseBodyEngineList) GetVersion() *string {
	return s.Version
}

func (s *GetLindormV2InstanceDetailsResponseBodyEngineList) SetConnectAddressList(v []*GetLindormV2InstanceDetailsResponseBodyEngineListConnectAddressList) *GetLindormV2InstanceDetailsResponseBodyEngineList {
	s.ConnectAddressList = v
	return s
}

func (s *GetLindormV2InstanceDetailsResponseBodyEngineList) SetEngine(v string) *GetLindormV2InstanceDetailsResponseBodyEngineList {
	s.Engine = &v
	return s
}

func (s *GetLindormV2InstanceDetailsResponseBodyEngineList) SetIsLastVersion(v bool) *GetLindormV2InstanceDetailsResponseBodyEngineList {
	s.IsLastVersion = &v
	return s
}

func (s *GetLindormV2InstanceDetailsResponseBodyEngineList) SetLatestVersion(v string) *GetLindormV2InstanceDetailsResponseBodyEngineList {
	s.LatestVersion = &v
	return s
}

func (s *GetLindormV2InstanceDetailsResponseBodyEngineList) SetNodeGroup(v []*GetLindormV2InstanceDetailsResponseBodyEngineListNodeGroup) *GetLindormV2InstanceDetailsResponseBodyEngineList {
	s.NodeGroup = v
	return s
}

func (s *GetLindormV2InstanceDetailsResponseBodyEngineList) SetVersion(v string) *GetLindormV2InstanceDetailsResponseBodyEngineList {
	s.Version = &v
	return s
}

func (s *GetLindormV2InstanceDetailsResponseBodyEngineList) Validate() error {
	if s.ConnectAddressList != nil {
		for _, item := range s.ConnectAddressList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.NodeGroup != nil {
		for _, item := range s.NodeGroup {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetLindormV2InstanceDetailsResponseBodyEngineListConnectAddressList struct {
	// The endpoint.
	//
	// example:
	//
	// ld-mxj9asg***-proxy-lindorm-vpc.lindorm.aliyuncs.com:33060
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The port number of the database endpoint.
	//
	// example:
	//
	// 33060
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The endpoint type.
	//
	// - INTRANET: VPC private endpoint.
	//
	// - INTERNET: Public endpoint.
	//
	// example:
	//
	// INTRANET
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetLindormV2InstanceDetailsResponseBodyEngineListConnectAddressList) String() string {
	return dara.Prettify(s)
}

func (s GetLindormV2InstanceDetailsResponseBodyEngineListConnectAddressList) GoString() string {
	return s.String()
}

func (s *GetLindormV2InstanceDetailsResponseBodyEngineListConnectAddressList) GetAddress() *string {
	return s.Address
}

func (s *GetLindormV2InstanceDetailsResponseBodyEngineListConnectAddressList) GetPort() *string {
	return s.Port
}

func (s *GetLindormV2InstanceDetailsResponseBodyEngineListConnectAddressList) GetType() *string {
	return s.Type
}

func (s *GetLindormV2InstanceDetailsResponseBodyEngineListConnectAddressList) SetAddress(v string) *GetLindormV2InstanceDetailsResponseBodyEngineListConnectAddressList {
	s.Address = &v
	return s
}

func (s *GetLindormV2InstanceDetailsResponseBodyEngineListConnectAddressList) SetPort(v string) *GetLindormV2InstanceDetailsResponseBodyEngineListConnectAddressList {
	s.Port = &v
	return s
}

func (s *GetLindormV2InstanceDetailsResponseBodyEngineListConnectAddressList) SetType(v string) *GetLindormV2InstanceDetailsResponseBodyEngineListConnectAddressList {
	s.Type = &v
	return s
}

func (s *GetLindormV2InstanceDetailsResponseBodyEngineListConnectAddressList) Validate() error {
	return dara.Validate(s)
}

type GetLindormV2InstanceDetailsResponseBodyEngineListNodeGroup struct {
	// Deprecated.
	//
	// example:
	//
	// caculated
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The number of vCPUs for the node.
	//
	// example:
	//
	// 32
	CpuCoreCount *int32 `json:"CpuCoreCount,omitempty" xml:"CpuCoreCount,omitempty"`
	// Indicates whether a local disk is attached to the node.
	//
	// example:
	//
	// false
	EnableAttachLocalDisk *bool `json:"EnableAttachLocalDisk,omitempty" xml:"EnableAttachLocalDisk,omitempty"`
	// The capacity of the local disk in GB.
	//
	// example:
	//
	// 100
	LocalDiskCapacity *int64 `json:"LocalDiskCapacity,omitempty" xml:"LocalDiskCapacity,omitempty"`
	// The type of the local cloud disk.
	//
	// - cloud_essd: performance cloud disk
	//
	// - cloud_efficiency: standard cloud disk
	//
	// example:
	//
	// cloud_essd
	LocalDiskCategory *string `json:"LocalDiskCategory,omitempty" xml:"LocalDiskCategory,omitempty"`
	// The memory size of the node in GiB.
	//
	// example:
	//
	// 64
	MemorySizeGiB *int32 `json:"MemorySizeGiB,omitempty" xml:"MemorySizeGiB,omitempty"`
	// The node specifications.
	//
	// If you select Performance cloud storage or Standard cloud storage, this parameter can be set to one of the following values:
	//
	// - lindorm.c.2xlarge: 8 vCPUs, 16 GB memory.
	//
	// - lindorm.g.2xlarge: 8 vCPUs, 32 GB memory.
	//
	// - lindorm.c.4xlarge: 16 vCPUs, 32 GB memory.
	//
	// - lindorm.g.4xlarge: 16 vCPUs, 64 GB memory.
	//
	// - lindorm.c.8xlarge: 32 vCPUs, 64 GB memory.
	//
	// - lindorm.g.8xlarge: 32 vCPUs, 128 GB memory.
	//
	// - lindorm.r.2xlarge: 8 vCPUs, 64 GB memory.
	//
	// - lindorm.r.4xlarge: 16 vCPUs, 128 GB memory.
	//
	// - lindorm.r.8xlarge: 32 vCPUs, 256 GB memory.
	//
	// If you select Local SSD, this parameter can be set to one of the following values:
	//
	// - lindorm.i4.xlarge: 4 vCPUs, 32 GB memory (I4).
	//
	// - lindorm.i4.2xlarge: 8 vCPUs, 64 GB memory (I4).
	//
	// - lindorm.i4.4xlarge: 16 vCPUs, 128 GB memory (I4).
	//
	// - lindorm.i4.8xlarge: 32 vCPUs, 256 GB memory (I4).
	//
	// - lindorm.i3.xlarge: 4 vCPUs, 32 GB memory (I3).
	//
	// - lindorm.i3.2xlarge: 8 vCPUs, 64 GB memory (I3).
	//
	// - lindorm.i3.4xlarge: 16 vCPUs, 128 GB memory (I3).
	//
	// - lindorm.i3.8xlarge: 32 vCPUs, 256 GB memory (I3).
	//
	// - lindorm.i2.xlarge: 4 vCPUs, 32 GB memory (I2).
	//
	// - lindorm.i2.2xlarge: 8 vCPUs, 64 GB memory (I2).
	//
	// - lindorm.i2.4xlarge: 16 vCPUs, 128 GB memory (I2).
	//
	// - lindorm.i2.8xlarge: 32 vCPUs, 256 GB memory (I2).
	//
	// If you select Big Data, this parameter can be set to one of the following values:
	//
	// - lindorm.sd3c.3xlarge: 14 vCPUs, 56 GB memory (D3C PRO).
	//
	// - lindorm.sd3c.7xlarge: 28 vCPUs, 112 GB memory (D3C PRO).
	//
	// - lindorm.sd3c.14xlarge: 56 vCPUs, 224 GB memory (D3C PRO).
	//
	// - lindorm.d2c.6xlarge: 24 vCPUs, 88 GB memory (D2C).
	//
	// - lindorm.d2c.12xlarge: 48 vCPUs, 176 GB memory (D2C).
	//
	// - lindorm.d2c.24xlarge: 96 vCPUs, 352 GB memory (D2C).
	//
	// - lindorm.d2s.5xlarge: 20 vCPUs, 88 GB memory (D2S).
	//
	// - lindorm.d2s.10xlarge: 40 vCPUs, 176 GB memory (D2S).
	//
	// - lindorm.d1.2xlarge: 8 vCPUs, 32 GB memory (D1NE).
	//
	// - lindorm.d1.4xlarge: 16 vCPUs, 64 GB memory (D1NE).
	//
	// - lindorm.d1.6xlarge: 24 vCPUs, 96 GB memory (D1NE).
	//
	// example:
	//
	// lindorm.g.2xlarge
	NodeSpec *string `json:"NodeSpec,omitempty" xml:"NodeSpec,omitempty"`
	// The number of nodes.
	//
	// example:
	//
	// 10
	Quantity *int32 `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	// The name of the node group. **Required**. This must be the same as the name used during creation.
	//
	// example:
	//
	// job_debug
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
	// The unique ID that corresponds to the delivery group ID.
	//
	// example:
	//
	// ecs.c6.large
	SpecId *string `json:"SpecId,omitempty" xml:"SpecId,omitempty"`
	// The node status.
	//
	// example:
	//
	// ACTIVATION
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetLindormV2InstanceDetailsResponseBodyEngineListNodeGroup) String() string {
	return dara.Prettify(s)
}

func (s GetLindormV2InstanceDetailsResponseBodyEngineListNodeGroup) GoString() string {
	return s.String()
}

func (s *GetLindormV2InstanceDetailsResponseBodyEngineListNodeGroup) GetCategory() *string {
	return s.Category
}

func (s *GetLindormV2InstanceDetailsResponseBodyEngineListNodeGroup) GetCpuCoreCount() *int32 {
	return s.CpuCoreCount
}

func (s *GetLindormV2InstanceDetailsResponseBodyEngineListNodeGroup) GetEnableAttachLocalDisk() *bool {
	return s.EnableAttachLocalDisk
}

func (s *GetLindormV2InstanceDetailsResponseBodyEngineListNodeGroup) GetLocalDiskCapacity() *int64 {
	return s.LocalDiskCapacity
}

func (s *GetLindormV2InstanceDetailsResponseBodyEngineListNodeGroup) GetLocalDiskCategory() *string {
	return s.LocalDiskCategory
}

func (s *GetLindormV2InstanceDetailsResponseBodyEngineListNodeGroup) GetMemorySizeGiB() *int32 {
	return s.MemorySizeGiB
}

func (s *GetLindormV2InstanceDetailsResponseBodyEngineListNodeGroup) GetNodeSpec() *string {
	return s.NodeSpec
}

func (s *GetLindormV2InstanceDetailsResponseBodyEngineListNodeGroup) GetQuantity() *int32 {
	return s.Quantity
}

func (s *GetLindormV2InstanceDetailsResponseBodyEngineListNodeGroup) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *GetLindormV2InstanceDetailsResponseBodyEngineListNodeGroup) GetSpecId() *string {
	return s.SpecId
}

func (s *GetLindormV2InstanceDetailsResponseBodyEngineListNodeGroup) GetStatus() *string {
	return s.Status
}

func (s *GetLindormV2InstanceDetailsResponseBodyEngineListNodeGroup) SetCategory(v string) *GetLindormV2InstanceDetailsResponseBodyEngineListNodeGroup {
	s.Category = &v
	return s
}

func (s *GetLindormV2InstanceDetailsResponseBodyEngineListNodeGroup) SetCpuCoreCount(v int32) *GetLindormV2InstanceDetailsResponseBodyEngineListNodeGroup {
	s.CpuCoreCount = &v
	return s
}

func (s *GetLindormV2InstanceDetailsResponseBodyEngineListNodeGroup) SetEnableAttachLocalDisk(v bool) *GetLindormV2InstanceDetailsResponseBodyEngineListNodeGroup {
	s.EnableAttachLocalDisk = &v
	return s
}

func (s *GetLindormV2InstanceDetailsResponseBodyEngineListNodeGroup) SetLocalDiskCapacity(v int64) *GetLindormV2InstanceDetailsResponseBodyEngineListNodeGroup {
	s.LocalDiskCapacity = &v
	return s
}

func (s *GetLindormV2InstanceDetailsResponseBodyEngineListNodeGroup) SetLocalDiskCategory(v string) *GetLindormV2InstanceDetailsResponseBodyEngineListNodeGroup {
	s.LocalDiskCategory = &v
	return s
}

func (s *GetLindormV2InstanceDetailsResponseBodyEngineListNodeGroup) SetMemorySizeGiB(v int32) *GetLindormV2InstanceDetailsResponseBodyEngineListNodeGroup {
	s.MemorySizeGiB = &v
	return s
}

func (s *GetLindormV2InstanceDetailsResponseBodyEngineListNodeGroup) SetNodeSpec(v string) *GetLindormV2InstanceDetailsResponseBodyEngineListNodeGroup {
	s.NodeSpec = &v
	return s
}

func (s *GetLindormV2InstanceDetailsResponseBodyEngineListNodeGroup) SetQuantity(v int32) *GetLindormV2InstanceDetailsResponseBodyEngineListNodeGroup {
	s.Quantity = &v
	return s
}

func (s *GetLindormV2InstanceDetailsResponseBodyEngineListNodeGroup) SetResourceGroupName(v string) *GetLindormV2InstanceDetailsResponseBodyEngineListNodeGroup {
	s.ResourceGroupName = &v
	return s
}

func (s *GetLindormV2InstanceDetailsResponseBodyEngineListNodeGroup) SetSpecId(v string) *GetLindormV2InstanceDetailsResponseBodyEngineListNodeGroup {
	s.SpecId = &v
	return s
}

func (s *GetLindormV2InstanceDetailsResponseBodyEngineListNodeGroup) SetStatus(v string) *GetLindormV2InstanceDetailsResponseBodyEngineListNodeGroup {
	s.Status = &v
	return s
}

func (s *GetLindormV2InstanceDetailsResponseBodyEngineListNodeGroup) Validate() error {
	return dara.Validate(s)
}

type GetLindormV2InstanceDetailsResponseBodyStorageUsage struct {
	// The instance storage usage.
	CapacityByDiskCategory []map[string]interface{} `json:"CapacityByDiskCategory,omitempty" xml:"CapacityByDiskCategory,omitempty" type:"Repeated"`
	// The storage usage of each engine.
	//
	// example:
	//
	// 16
	EngineUsage map[string]interface{} `json:"EngineUsage,omitempty" xml:"EngineUsage,omitempty"`
}

func (s GetLindormV2InstanceDetailsResponseBodyStorageUsage) String() string {
	return dara.Prettify(s)
}

func (s GetLindormV2InstanceDetailsResponseBodyStorageUsage) GoString() string {
	return s.String()
}

func (s *GetLindormV2InstanceDetailsResponseBodyStorageUsage) GetCapacityByDiskCategory() []map[string]interface{} {
	return s.CapacityByDiskCategory
}

func (s *GetLindormV2InstanceDetailsResponseBodyStorageUsage) GetEngineUsage() map[string]interface{} {
	return s.EngineUsage
}

func (s *GetLindormV2InstanceDetailsResponseBodyStorageUsage) SetCapacityByDiskCategory(v []map[string]interface{}) *GetLindormV2InstanceDetailsResponseBodyStorageUsage {
	s.CapacityByDiskCategory = v
	return s
}

func (s *GetLindormV2InstanceDetailsResponseBodyStorageUsage) SetEngineUsage(v map[string]interface{}) *GetLindormV2InstanceDetailsResponseBodyStorageUsage {
	s.EngineUsage = v
	return s
}

func (s *GetLindormV2InstanceDetailsResponseBodyStorageUsage) Validate() error {
	return dara.Validate(s)
}

type GetLindormV2InstanceDetailsResponseBodyWhiteIpList struct {
	// The group name. The name can contain only letters, digits, and underscores (_).
	//
	// example:
	//
	// swhy
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The IP addresses in the whitelist.
	//
	// example:
	//
	// [\\"10.2.0.0/18\\", \\"10.0.0.0/24\\", \\"119.23.188.139/32\\"]
	IpList *string `json:"IpList,omitempty" xml:"IpList,omitempty"`
}

func (s GetLindormV2InstanceDetailsResponseBodyWhiteIpList) String() string {
	return dara.Prettify(s)
}

func (s GetLindormV2InstanceDetailsResponseBodyWhiteIpList) GoString() string {
	return s.String()
}

func (s *GetLindormV2InstanceDetailsResponseBodyWhiteIpList) GetGroupName() *string {
	return s.GroupName
}

func (s *GetLindormV2InstanceDetailsResponseBodyWhiteIpList) GetIpList() *string {
	return s.IpList
}

func (s *GetLindormV2InstanceDetailsResponseBodyWhiteIpList) SetGroupName(v string) *GetLindormV2InstanceDetailsResponseBodyWhiteIpList {
	s.GroupName = &v
	return s
}

func (s *GetLindormV2InstanceDetailsResponseBodyWhiteIpList) SetIpList(v string) *GetLindormV2InstanceDetailsResponseBodyWhiteIpList {
	s.IpList = &v
	return s
}

func (s *GetLindormV2InstanceDetailsResponseBodyWhiteIpList) Validate() error {
	return dara.Validate(s)
}
