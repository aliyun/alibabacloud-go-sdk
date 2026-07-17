// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLindormInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAliUid(v int64) *GetLindormInstanceResponseBody
	GetAliUid() *int64
	SetArbiterVSwitchId(v string) *GetLindormInstanceResponseBody
	GetArbiterVSwitchId() *string
	SetArbiterZoneId(v string) *GetLindormInstanceResponseBody
	GetArbiterZoneId() *string
	SetArchVersion(v string) *GetLindormInstanceResponseBody
	GetArchVersion() *string
	SetArchiveStorage(v int32) *GetLindormInstanceResponseBody
	GetArchiveStorage() *int32
	SetAutoRenew(v bool) *GetLindormInstanceResponseBody
	GetAutoRenew() *bool
	SetBackupInstance(v string) *GetLindormInstanceResponseBody
	GetBackupInstance() *string
	SetColdStorage(v int32) *GetLindormInstanceResponseBody
	GetColdStorage() *int32
	SetCoreDiskCategory(v string) *GetLindormInstanceResponseBody
	GetCoreDiskCategory() *string
	SetCoreNum(v int32) *GetLindormInstanceResponseBody
	GetCoreNum() *int32
	SetCoreSingleStorage(v int32) *GetLindormInstanceResponseBody
	GetCoreSingleStorage() *int32
	SetCoreSpec(v string) *GetLindormInstanceResponseBody
	GetCoreSpec() *string
	SetCreateMilliseconds(v int64) *GetLindormInstanceResponseBody
	GetCreateMilliseconds() *int64
	SetCreateTime(v string) *GetLindormInstanceResponseBody
	GetCreateTime() *string
	SetDeletionProtection(v string) *GetLindormInstanceResponseBody
	GetDeletionProtection() *string
	SetDiskCategory(v string) *GetLindormInstanceResponseBody
	GetDiskCategory() *string
	SetDiskThreshold(v string) *GetLindormInstanceResponseBody
	GetDiskThreshold() *string
	SetDiskUsage(v string) *GetLindormInstanceResponseBody
	GetDiskUsage() *string
	SetEnableBlob(v bool) *GetLindormInstanceResponseBody
	GetEnableBlob() *bool
	SetEnableCdc(v bool) *GetLindormInstanceResponseBody
	GetEnableCdc() *bool
	SetEnableCompute(v bool) *GetLindormInstanceResponseBody
	GetEnableCompute() *bool
	SetEnableKms(v bool) *GetLindormInstanceResponseBody
	GetEnableKms() *bool
	SetEnableLProxy(v bool) *GetLindormInstanceResponseBody
	GetEnableLProxy() *bool
	SetEnableLTS(v bool) *GetLindormInstanceResponseBody
	GetEnableLTS() *bool
	SetEnableLsqlVersionV3(v bool) *GetLindormInstanceResponseBody
	GetEnableLsqlVersionV3() *bool
	SetEnableMLCtrl(v bool) *GetLindormInstanceResponseBody
	GetEnableMLCtrl() *bool
	SetEnableSSL(v bool) *GetLindormInstanceResponseBody
	GetEnableSSL() *bool
	SetEnableShs(v bool) *GetLindormInstanceResponseBody
	GetEnableShs() *bool
	SetEnableStoreTDE(v bool) *GetLindormInstanceResponseBody
	GetEnableStoreTDE() *bool
	SetEnableStream(v bool) *GetLindormInstanceResponseBody
	GetEnableStream() *bool
	SetEngineList(v []*GetLindormInstanceResponseBodyEngineList) *GetLindormInstanceResponseBody
	GetEngineList() []*GetLindormInstanceResponseBodyEngineList
	SetEngineType(v int32) *GetLindormInstanceResponseBody
	GetEngineType() *int32
	SetExpireTime(v string) *GetLindormInstanceResponseBody
	GetExpireTime() *string
	SetExpiredMilliseconds(v int64) *GetLindormInstanceResponseBody
	GetExpiredMilliseconds() *int64
	SetInstanceAlias(v string) *GetLindormInstanceResponseBody
	GetInstanceAlias() *string
	SetInstanceId(v string) *GetLindormInstanceResponseBody
	GetInstanceId() *string
	SetInstanceStatus(v string) *GetLindormInstanceResponseBody
	GetInstanceStatus() *string
	SetInstanceStorage(v string) *GetLindormInstanceResponseBody
	GetInstanceStorage() *string
	SetLogDiskCategory(v string) *GetLindormInstanceResponseBody
	GetLogDiskCategory() *string
	SetLogNum(v int32) *GetLindormInstanceResponseBody
	GetLogNum() *int32
	SetLogSingleStorage(v int32) *GetLindormInstanceResponseBody
	GetLogSingleStorage() *int32
	SetLogSpec(v string) *GetLindormInstanceResponseBody
	GetLogSpec() *string
	SetMaintainEndTime(v string) *GetLindormInstanceResponseBody
	GetMaintainEndTime() *string
	SetMaintainStartTime(v string) *GetLindormInstanceResponseBody
	GetMaintainStartTime() *string
	SetMultiZoneCombination(v string) *GetLindormInstanceResponseBody
	GetMultiZoneCombination() *string
	SetNetworkType(v string) *GetLindormInstanceResponseBody
	GetNetworkType() *string
	SetPayType(v string) *GetLindormInstanceResponseBody
	GetPayType() *string
	SetPrimaryVSwitchId(v string) *GetLindormInstanceResponseBody
	GetPrimaryVSwitchId() *string
	SetPrimaryZoneId(v string) *GetLindormInstanceResponseBody
	GetPrimaryZoneId() *string
	SetRegionId(v string) *GetLindormInstanceResponseBody
	GetRegionId() *string
	SetRequestId(v string) *GetLindormInstanceResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *GetLindormInstanceResponseBody
	GetResourceGroupId() *string
	SetServiceType(v string) *GetLindormInstanceResponseBody
	GetServiceType() *string
	SetStandbyVSwitchId(v string) *GetLindormInstanceResponseBody
	GetStandbyVSwitchId() *string
	SetStandbyZoneId(v string) *GetLindormInstanceResponseBody
	GetStandbyZoneId() *string
	SetVpcId(v string) *GetLindormInstanceResponseBody
	GetVpcId() *string
	SetVswitchId(v string) *GetLindormInstanceResponseBody
	GetVswitchId() *string
	SetZoneId(v string) *GetLindormInstanceResponseBody
	GetZoneId() *string
}

type GetLindormInstanceResponseBody struct {
	// The UID of the Alibaba Cloud account.
	//
	// example:
	//
	// 164901546557****
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The ID of the vSwitch in the arbiter zone for the multi-zone instance. The vSwitch must be deployed in the zone that is specified by `ArbiterZoneId`.
	//
	// example:
	//
	// vsw-uf6664pqjawb87k36****
	ArbiterVSwitchId *string `json:"ArbiterVSwitchId,omitempty" xml:"ArbiterVSwitchId,omitempty"`
	// The arbiter zone ID of the multi-zone instance.
	//
	// example:
	//
	// cn-shanghai-g
	ArbiterZoneId *string `json:"ArbiterZoneId,omitempty" xml:"ArbiterZoneId,omitempty"`
	// The deployment architecture. Valid values:
	//
	// - **1.0**: single-zone deployment.
	//
	// - **2.0**: multi-zone deployment.
	//
	// example:
	//
	// 1.0
	ArchVersion *string `json:"ArchVersion,omitempty" xml:"ArchVersion,omitempty"`
	// The billable storage capacity of the archive storage. Unit: GB.
	//
	// example:
	//
	// 0GB
	ArchiveStorage *int32 `json:"ArchiveStorage,omitempty" xml:"ArchiveStorage,omitempty"`
	// Indicates whether auto-renewal is enabled for the instance. Valid values:
	//
	// - **true**: Enabled.
	//
	// - **false**: Disabled.
	//
	// > This parameter is returned only for subscription instances.
	//
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The ID of the backup instance.
	//
	// example:
	//
	// ld-xxxx
	BackupInstance *string `json:"BackupInstance,omitempty" xml:"BackupInstance,omitempty"`
	// The capacity of the cold storage.
	//
	// example:
	//
	// 0GB
	ColdStorage *int32 `json:"ColdStorage,omitempty" xml:"ColdStorage,omitempty"`
	// The disk type of the core nodes in a multi-zone instance. Valid values:
	//
	// - **cloud_efficiency**: Standard.
	//
	// - **cloud_ssd**: Performance.
	//
	// - **cloud_essd**: ESSD.
	//
	// - **cloud_essd_pl0**: ESSD PL0.
	//
	// example:
	//
	// cloud_efficiency
	CoreDiskCategory *string `json:"CoreDiskCategory,omitempty" xml:"CoreDiskCategory,omitempty"`
	// The number of core nodes in the multi-zone instance.
	//
	// example:
	//
	// 4
	CoreNum *int32 `json:"CoreNum,omitempty" xml:"CoreNum,omitempty"`
	// The storage capacity of a single core node in the multi-zone instance.
	//
	// example:
	//
	// 400
	CoreSingleStorage *int32 `json:"CoreSingleStorage,omitempty" xml:"CoreSingleStorage,omitempty"`
	// The specification of the core nodes in the multi-zone instance.
	//
	// example:
	//
	// lindorm.g.xlarge
	CoreSpec *string `json:"CoreSpec,omitempty" xml:"CoreSpec,omitempty"`
	// The time at which the instance was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1627290664000
	CreateMilliseconds *int64 `json:"CreateMilliseconds,omitempty" xml:"CreateMilliseconds,omitempty"`
	// The time at which the instance was created. The time is displayed in the **yyyy-MM-dd HH:mm:ss*	- format.
	//
	// example:
	//
	// 2021-07-26 17:10:26
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether release protection is enabled for the instance. Valid values:
	//
	// - **true**: Enabled.
	//
	// - **false**: Disabled.
	//
	// example:
	//
	// false
	DeletionProtection *string `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// The storage type. Valid values:
	//
	// - **cloud_efficiency**: Standard.
	//
	// - **cloud_ssd**: Performance.
	//
	// - **cloud_essd**: Enhanced SSD (ESSD).
	//
	// - **cloud_essd_pl0**: ESSD PL0.
	//
	// - **capacity_cloud_storage**: Capacity.
	//
	// - **local_ssd_pro**: local SSD.
	//
	// - **local_hdd_pro**: local HDD.
	//
	// example:
	//
	// cloud_efficiency
	DiskCategory *string `json:"DiskCategory,omitempty" xml:"DiskCategory,omitempty"`
	// The disk space threshold.
	//
	// example:
	//
	// 80%
	DiskThreshold *string `json:"DiskThreshold,omitempty" xml:"DiskThreshold,omitempty"`
	// The disk usage.
	//
	// example:
	//
	// 0.0%
	DiskUsage *string `json:"DiskUsage,omitempty" xml:"DiskUsage,omitempty"`
	// Indicates whether LBlob is enabled. Valid values:
	//
	// true: Enabled. false: Disabled.
	//
	// example:
	//
	// true
	EnableBlob *bool `json:"EnableBlob,omitempty" xml:"EnableBlob,omitempty"`
	// Indicates whether Change Data Capture (CDC) is enabled for the instance. Valid values:
	//
	// - **true**: Enabled.
	//
	// - **false**: Disabled.
	//
	// example:
	//
	// false
	EnableCdc *bool `json:"EnableCdc,omitempty" xml:"EnableCdc,omitempty"`
	// Indicates whether the compute engine is enabled for the instance. Valid values:
	//
	// - **true**: Enabled.
	//
	// - **false**: Disabled.
	//
	// example:
	//
	// true
	EnableCompute *bool `json:"EnableCompute,omitempty" xml:"EnableCompute,omitempty"`
	// Indicates whether Key Management Service (KMS) is enabled. Valid values:
	//
	// - **true**: Enabled.
	//
	// - **false**: Disabled.
	//
	// example:
	//
	// false
	EnableKms *bool `json:"EnableKms,omitempty" xml:"EnableKms,omitempty"`
	// Specifies whether the wide table engine supports the Thrift and CQL protocols. If this feature is disabled, you can call the SwitchLProxyService operation to enable it.
	//
	// true: Supported.
	//
	// false: Not supported.
	//
	// example:
	//
	// False
	EnableLProxy *bool `json:"EnableLProxy,omitempty" xml:"EnableLProxy,omitempty"`
	// Indicates whether the LTS engine is enabled for the instance. Valid values:
	//
	// - **true**: Enabled.
	//
	// - **false**: Disabled.
	//
	// example:
	//
	// true
	EnableLTS *bool `json:"EnableLTS,omitempty" xml:"EnableLTS,omitempty"`
	// Indicates whether LindormSQL V3.0, which is compatible with the MySQL protocol, is supported by the wide table engine.
	//
	// This feature is supported by default on instances created after October 24, 2023. For existing instances, contact technical support to enable this feature.
	//
	// - true: Supported.
	//
	// - false: Not supported.
	//
	// example:
	//
	// True
	EnableLsqlVersionV3 *bool `json:"EnableLsqlVersionV3,omitempty" xml:"EnableLsqlVersionV3,omitempty"`
	// Indicates whether the ML node is enabled. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// example:
	//
	// False
	EnableMLCtrl *bool `json:"EnableMLCtrl,omitempty" xml:"EnableMLCtrl,omitempty"`
	// Indicates whether SSL encryption is enabled. Valid values:
	//
	// - **true**: Enabled.
	//
	// - **false**: Disabled.
	//
	// example:
	//
	// false
	EnableSSL *bool `json:"EnableSSL,omitempty" xml:"EnableSSL,omitempty"`
	// Indicates whether the History Server is enabled for the compute engine.
	//
	// example:
	//
	// true
	EnableShs *bool `json:"EnableShs,omitempty" xml:"EnableShs,omitempty"`
	// Indicates whether Transparent Data Encryption (TDE) is enabled. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// example:
	//
	// false
	EnableStoreTDE *bool `json:"EnableStoreTDE,omitempty" xml:"EnableStoreTDE,omitempty"`
	// Indicates whether the stream engine is enabled for the instance. Valid values:
	//
	// - **true**: Enabled.
	//
	// - **false**: Disabled.
	//
	// example:
	//
	// true
	EnableStream *bool `json:"EnableStream,omitempty" xml:"EnableStream,omitempty"`
	// The information about the engines.
	EngineList []*GetLindormInstanceResponseBodyEngineList `json:"EngineList,omitempty" xml:"EngineList,omitempty" type:"Repeated"`
	// The types of the engines that are supported by the instance. The value of this parameter is the sum of the values of all supported engine types.
	//
	// - 1: search engine
	//
	// - 2: time series engine
	//
	// - 4: wide table engine
	//
	// - 8: file engine
	//
	// > For example, if the value of this parameter is 15, it indicates that the instance supports the search, time series, wide table, and file engines because 1 + 2 + 4 + 8 = 15. If the value of this parameter is 6, it indicates that the instance supports the time series and wide table engines because 2 + 4 = 6.
	//
	// example:
	//
	// 15
	EngineType *int32 `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// The expiration time of the instance. The time is displayed in the **yyyy-MM-dd HH:mm:ss*	- format.
	//
	// > This parameter is returned only for subscription instances.
	//
	// example:
	//
	// 2021-08-27 00:00:00
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The expiration time of the instance. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1629993600000
	ExpiredMilliseconds *int64 `json:"ExpiredMilliseconds,omitempty" xml:"ExpiredMilliseconds,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// test0726
	InstanceAlias *string `json:"InstanceAlias,omitempty" xml:"InstanceAlias,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// ld-bp1o3y0yme2i2****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The status of the instance. Valid values:
	//
	// - **CREATING**: The instance is being created.
	//
	// - **ACTIVATION**: The instance is running.
	//
	// - **COLD_EXPANDING**: The capacity of the cold storage is being expanded.
	//
	// - **MINOR_VERSION_TRANSITIONING**: The minor version of the instance is being changed.
	//
	// - **RESIZING**: The number of nodes is being changed.
	//
	// - **SHRINKING**: The number of nodes is being changed.
	//
	// - **CLASS_CHANGING**: The specification of the instance is being changed.
	//
	// - **SSL_SWITCHING**: SSL is being enabled or disabled.
	//
	// - **CDC_OPENING**: The CDC feature is being enabled.
	//
	// - **TRANSFER**: Data is being migrated.
	//
	// - **DATABASE_TRANSFER**: Data is being migrated.
	//
	// - **GUARD_CREATING**: A disaster recovery instance is being created.
	//
	// - **BACKUP_RECOVERING**: Data is being restored from a backup.
	//
	// - **DATABASE_IMPORTING**: Data is being imported.
	//
	// - **NET_MODIFYING**: The network type is being changed.
	//
	// - **NET_SWITCHING**: The network is being switched.
	//
	// - **NET_CREATING**: A network connection is being created.
	//
	// - **NET_DELETING**: A network connection is being deleted.
	//
	// - **DELETING**: The instance is being deleted.
	//
	// - **RESTARTING**: The instance is being restarted.
	//
	// - **LOCKED**: The instance is locked.
	//
	// example:
	//
	// ACTIVATION
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// The storage capacity of the instance.
	//
	// example:
	//
	// 480
	InstanceStorage *string `json:"InstanceStorage,omitempty" xml:"InstanceStorage,omitempty"`
	// The disk type of the log nodes in the multi-zone instance. Valid values:
	//
	// - **cloud_efficiency**: Standard.
	//
	// - **cloud_ssd**: Performance.
	//
	// example:
	//
	// cloud_ssd
	LogDiskCategory *string `json:"LogDiskCategory,omitempty" xml:"LogDiskCategory,omitempty"`
	// The number of log nodes in the multi-zone instance.
	//
	// example:
	//
	// 4
	LogNum *int32 `json:"LogNum,omitempty" xml:"LogNum,omitempty"`
	// The storage capacity of a single log node in the multi-zone instance.
	//
	// example:
	//
	// 400GB
	LogSingleStorage *int32 `json:"LogSingleStorage,omitempty" xml:"LogSingleStorage,omitempty"`
	// The specification of the log nodes in the multi-zone instance.
	//
	// example:
	//
	// lindorm.sn1.large
	LogSpec *string `json:"LogSpec,omitempty" xml:"LogSpec,omitempty"`
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
	// The combination of zones. For more information about the supported zone combinations, see the instance buy page.
	//
	// - **ap-southeast-5abc-aliyun**: Indonesia (Jakarta) Zone A, B, and C.
	//
	// - **cn-hangzhou-ehi-aliyun**: China (Hangzhou) Zone E, H, and I.
	//
	// - **cn-beijing-acd-aliyun**: China (Beijing) Zone A, C, and D.
	//
	// - **ap-southeast-1-abc-aliyun**: Singapore Zone A, B, and C.
	//
	// - **cn-zhangjiakou-abc-aliyun**: China (Zhangjiakou) Zone A, B, and C.
	//
	// - **cn-shanghai-efg-aliyun**: China (Shanghai) Zone E, F, and G.
	//
	// - **cn-shanghai-abd-aliyun**: China (Shanghai) Zone A, B, and D.
	//
	// - **cn-hangzhou-bef-aliyun**: China (Hangzhou) Zone B, E, and F.
	//
	// - **cn-hangzhou-bce-aliyun**: China (Hangzhou) Zone B, C, and E.
	//
	// - **cn-beijing-fgh-aliyun**: China (Beijing) Zone F, G, and H.
	//
	// - **cn-shenzhen-abc-aliyun**: China (Shenzhen) Zone A, B, and C.
	//
	// example:
	//
	// cn-shanghai-efg-aliyun
	MultiZoneCombination *string `json:"MultiZoneCombination,omitempty" xml:"MultiZoneCombination,omitempty"`
	// The network type of the instance.
	//
	// example:
	//
	// vpc
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// - **PREPAY**: subscription
	//
	// - **POSTPAY**: pay-as-you-go
	//
	// example:
	//
	// POSTPAY
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The ID of the vSwitch in the primary zone for the multi-zone instance. The vSwitch must be deployed in the zone that is specified by `PrimaryZoneId`.
	//
	// example:
	//
	// vsw-uf6fdqa7c0pipnqzq****
	PrimaryVSwitchId *string `json:"PrimaryVSwitchId,omitempty" xml:"PrimaryVSwitchId,omitempty"`
	// The primary zone ID of the multi-zone instance.
	//
	// example:
	//
	// cn-shanghai-e
	PrimaryZoneId *string `json:"PrimaryZoneId,omitempty" xml:"PrimaryZoneId,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 633F1BE4-C8DA-5744-8FDF-A3075C3FE37F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aek2wvd6oia****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The type of the instance. Valid values:
	//
	// - **lindorm**: a single-zone instance.
	//
	// - **lindorm_multizone**: a multi-zone instance.
	//
	// - **serverless_lindorm**: a serverless instance.
	//
	// - **lindorm_standalone**: a single-node instance.
	//
	// - **lts**: a Lindorm Tunnel Service (LTS) instance.
	//
	// example:
	//
	// lindorm
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// The ID of the vSwitch in the secondary zone for the multi-zone instance. The vSwitch must be deployed in the zone that is specified by `StandbyZoneId`.
	//
	// example:
	//
	// vsw-2zec0kcn08cgdtr6****
	StandbyVSwitchId *string `json:"StandbyVSwitchId,omitempty" xml:"StandbyVSwitchId,omitempty"`
	// The secondary zone ID of the multi-zone instance.
	//
	// example:
	//
	// cn-shanghai-f
	StandbyZoneId *string `json:"StandbyZoneId,omitempty" xml:"StandbyZoneId,omitempty"`
	// The ID of the virtual private cloud (VPC) to which the instance belongs.
	//
	// example:
	//
	// vpc-bp1n3i15v90el48nx****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the vSwitch.
	//
	// example:
	//
	// vsw-bp1vbjzmod9q3l9eo****
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	// The ID of the zone.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetLindormInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLindormInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetLindormInstanceResponseBody) GetAliUid() *int64 {
	return s.AliUid
}

func (s *GetLindormInstanceResponseBody) GetArbiterVSwitchId() *string {
	return s.ArbiterVSwitchId
}

func (s *GetLindormInstanceResponseBody) GetArbiterZoneId() *string {
	return s.ArbiterZoneId
}

func (s *GetLindormInstanceResponseBody) GetArchVersion() *string {
	return s.ArchVersion
}

func (s *GetLindormInstanceResponseBody) GetArchiveStorage() *int32 {
	return s.ArchiveStorage
}

func (s *GetLindormInstanceResponseBody) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *GetLindormInstanceResponseBody) GetBackupInstance() *string {
	return s.BackupInstance
}

func (s *GetLindormInstanceResponseBody) GetColdStorage() *int32 {
	return s.ColdStorage
}

func (s *GetLindormInstanceResponseBody) GetCoreDiskCategory() *string {
	return s.CoreDiskCategory
}

func (s *GetLindormInstanceResponseBody) GetCoreNum() *int32 {
	return s.CoreNum
}

func (s *GetLindormInstanceResponseBody) GetCoreSingleStorage() *int32 {
	return s.CoreSingleStorage
}

func (s *GetLindormInstanceResponseBody) GetCoreSpec() *string {
	return s.CoreSpec
}

func (s *GetLindormInstanceResponseBody) GetCreateMilliseconds() *int64 {
	return s.CreateMilliseconds
}

func (s *GetLindormInstanceResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetLindormInstanceResponseBody) GetDeletionProtection() *string {
	return s.DeletionProtection
}

func (s *GetLindormInstanceResponseBody) GetDiskCategory() *string {
	return s.DiskCategory
}

func (s *GetLindormInstanceResponseBody) GetDiskThreshold() *string {
	return s.DiskThreshold
}

func (s *GetLindormInstanceResponseBody) GetDiskUsage() *string {
	return s.DiskUsage
}

func (s *GetLindormInstanceResponseBody) GetEnableBlob() *bool {
	return s.EnableBlob
}

func (s *GetLindormInstanceResponseBody) GetEnableCdc() *bool {
	return s.EnableCdc
}

func (s *GetLindormInstanceResponseBody) GetEnableCompute() *bool {
	return s.EnableCompute
}

func (s *GetLindormInstanceResponseBody) GetEnableKms() *bool {
	return s.EnableKms
}

func (s *GetLindormInstanceResponseBody) GetEnableLProxy() *bool {
	return s.EnableLProxy
}

func (s *GetLindormInstanceResponseBody) GetEnableLTS() *bool {
	return s.EnableLTS
}

func (s *GetLindormInstanceResponseBody) GetEnableLsqlVersionV3() *bool {
	return s.EnableLsqlVersionV3
}

func (s *GetLindormInstanceResponseBody) GetEnableMLCtrl() *bool {
	return s.EnableMLCtrl
}

func (s *GetLindormInstanceResponseBody) GetEnableSSL() *bool {
	return s.EnableSSL
}

func (s *GetLindormInstanceResponseBody) GetEnableShs() *bool {
	return s.EnableShs
}

func (s *GetLindormInstanceResponseBody) GetEnableStoreTDE() *bool {
	return s.EnableStoreTDE
}

func (s *GetLindormInstanceResponseBody) GetEnableStream() *bool {
	return s.EnableStream
}

func (s *GetLindormInstanceResponseBody) GetEngineList() []*GetLindormInstanceResponseBodyEngineList {
	return s.EngineList
}

func (s *GetLindormInstanceResponseBody) GetEngineType() *int32 {
	return s.EngineType
}

func (s *GetLindormInstanceResponseBody) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *GetLindormInstanceResponseBody) GetExpiredMilliseconds() *int64 {
	return s.ExpiredMilliseconds
}

func (s *GetLindormInstanceResponseBody) GetInstanceAlias() *string {
	return s.InstanceAlias
}

func (s *GetLindormInstanceResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetLindormInstanceResponseBody) GetInstanceStatus() *string {
	return s.InstanceStatus
}

func (s *GetLindormInstanceResponseBody) GetInstanceStorage() *string {
	return s.InstanceStorage
}

func (s *GetLindormInstanceResponseBody) GetLogDiskCategory() *string {
	return s.LogDiskCategory
}

func (s *GetLindormInstanceResponseBody) GetLogNum() *int32 {
	return s.LogNum
}

func (s *GetLindormInstanceResponseBody) GetLogSingleStorage() *int32 {
	return s.LogSingleStorage
}

func (s *GetLindormInstanceResponseBody) GetLogSpec() *string {
	return s.LogSpec
}

func (s *GetLindormInstanceResponseBody) GetMaintainEndTime() *string {
	return s.MaintainEndTime
}

func (s *GetLindormInstanceResponseBody) GetMaintainStartTime() *string {
	return s.MaintainStartTime
}

func (s *GetLindormInstanceResponseBody) GetMultiZoneCombination() *string {
	return s.MultiZoneCombination
}

func (s *GetLindormInstanceResponseBody) GetNetworkType() *string {
	return s.NetworkType
}

func (s *GetLindormInstanceResponseBody) GetPayType() *string {
	return s.PayType
}

func (s *GetLindormInstanceResponseBody) GetPrimaryVSwitchId() *string {
	return s.PrimaryVSwitchId
}

func (s *GetLindormInstanceResponseBody) GetPrimaryZoneId() *string {
	return s.PrimaryZoneId
}

func (s *GetLindormInstanceResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *GetLindormInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLindormInstanceResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetLindormInstanceResponseBody) GetServiceType() *string {
	return s.ServiceType
}

func (s *GetLindormInstanceResponseBody) GetStandbyVSwitchId() *string {
	return s.StandbyVSwitchId
}

func (s *GetLindormInstanceResponseBody) GetStandbyZoneId() *string {
	return s.StandbyZoneId
}

func (s *GetLindormInstanceResponseBody) GetVpcId() *string {
	return s.VpcId
}

func (s *GetLindormInstanceResponseBody) GetVswitchId() *string {
	return s.VswitchId
}

func (s *GetLindormInstanceResponseBody) GetZoneId() *string {
	return s.ZoneId
}

func (s *GetLindormInstanceResponseBody) SetAliUid(v int64) *GetLindormInstanceResponseBody {
	s.AliUid = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetArbiterVSwitchId(v string) *GetLindormInstanceResponseBody {
	s.ArbiterVSwitchId = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetArbiterZoneId(v string) *GetLindormInstanceResponseBody {
	s.ArbiterZoneId = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetArchVersion(v string) *GetLindormInstanceResponseBody {
	s.ArchVersion = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetArchiveStorage(v int32) *GetLindormInstanceResponseBody {
	s.ArchiveStorage = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetAutoRenew(v bool) *GetLindormInstanceResponseBody {
	s.AutoRenew = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetBackupInstance(v string) *GetLindormInstanceResponseBody {
	s.BackupInstance = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetColdStorage(v int32) *GetLindormInstanceResponseBody {
	s.ColdStorage = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetCoreDiskCategory(v string) *GetLindormInstanceResponseBody {
	s.CoreDiskCategory = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetCoreNum(v int32) *GetLindormInstanceResponseBody {
	s.CoreNum = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetCoreSingleStorage(v int32) *GetLindormInstanceResponseBody {
	s.CoreSingleStorage = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetCoreSpec(v string) *GetLindormInstanceResponseBody {
	s.CoreSpec = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetCreateMilliseconds(v int64) *GetLindormInstanceResponseBody {
	s.CreateMilliseconds = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetCreateTime(v string) *GetLindormInstanceResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetDeletionProtection(v string) *GetLindormInstanceResponseBody {
	s.DeletionProtection = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetDiskCategory(v string) *GetLindormInstanceResponseBody {
	s.DiskCategory = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetDiskThreshold(v string) *GetLindormInstanceResponseBody {
	s.DiskThreshold = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetDiskUsage(v string) *GetLindormInstanceResponseBody {
	s.DiskUsage = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetEnableBlob(v bool) *GetLindormInstanceResponseBody {
	s.EnableBlob = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetEnableCdc(v bool) *GetLindormInstanceResponseBody {
	s.EnableCdc = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetEnableCompute(v bool) *GetLindormInstanceResponseBody {
	s.EnableCompute = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetEnableKms(v bool) *GetLindormInstanceResponseBody {
	s.EnableKms = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetEnableLProxy(v bool) *GetLindormInstanceResponseBody {
	s.EnableLProxy = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetEnableLTS(v bool) *GetLindormInstanceResponseBody {
	s.EnableLTS = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetEnableLsqlVersionV3(v bool) *GetLindormInstanceResponseBody {
	s.EnableLsqlVersionV3 = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetEnableMLCtrl(v bool) *GetLindormInstanceResponseBody {
	s.EnableMLCtrl = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetEnableSSL(v bool) *GetLindormInstanceResponseBody {
	s.EnableSSL = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetEnableShs(v bool) *GetLindormInstanceResponseBody {
	s.EnableShs = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetEnableStoreTDE(v bool) *GetLindormInstanceResponseBody {
	s.EnableStoreTDE = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetEnableStream(v bool) *GetLindormInstanceResponseBody {
	s.EnableStream = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetEngineList(v []*GetLindormInstanceResponseBodyEngineList) *GetLindormInstanceResponseBody {
	s.EngineList = v
	return s
}

func (s *GetLindormInstanceResponseBody) SetEngineType(v int32) *GetLindormInstanceResponseBody {
	s.EngineType = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetExpireTime(v string) *GetLindormInstanceResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetExpiredMilliseconds(v int64) *GetLindormInstanceResponseBody {
	s.ExpiredMilliseconds = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetInstanceAlias(v string) *GetLindormInstanceResponseBody {
	s.InstanceAlias = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetInstanceId(v string) *GetLindormInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetInstanceStatus(v string) *GetLindormInstanceResponseBody {
	s.InstanceStatus = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetInstanceStorage(v string) *GetLindormInstanceResponseBody {
	s.InstanceStorage = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetLogDiskCategory(v string) *GetLindormInstanceResponseBody {
	s.LogDiskCategory = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetLogNum(v int32) *GetLindormInstanceResponseBody {
	s.LogNum = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetLogSingleStorage(v int32) *GetLindormInstanceResponseBody {
	s.LogSingleStorage = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetLogSpec(v string) *GetLindormInstanceResponseBody {
	s.LogSpec = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetMaintainEndTime(v string) *GetLindormInstanceResponseBody {
	s.MaintainEndTime = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetMaintainStartTime(v string) *GetLindormInstanceResponseBody {
	s.MaintainStartTime = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetMultiZoneCombination(v string) *GetLindormInstanceResponseBody {
	s.MultiZoneCombination = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetNetworkType(v string) *GetLindormInstanceResponseBody {
	s.NetworkType = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetPayType(v string) *GetLindormInstanceResponseBody {
	s.PayType = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetPrimaryVSwitchId(v string) *GetLindormInstanceResponseBody {
	s.PrimaryVSwitchId = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetPrimaryZoneId(v string) *GetLindormInstanceResponseBody {
	s.PrimaryZoneId = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetRegionId(v string) *GetLindormInstanceResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetRequestId(v string) *GetLindormInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetResourceGroupId(v string) *GetLindormInstanceResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetServiceType(v string) *GetLindormInstanceResponseBody {
	s.ServiceType = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetStandbyVSwitchId(v string) *GetLindormInstanceResponseBody {
	s.StandbyVSwitchId = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetStandbyZoneId(v string) *GetLindormInstanceResponseBody {
	s.StandbyZoneId = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetVpcId(v string) *GetLindormInstanceResponseBody {
	s.VpcId = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetVswitchId(v string) *GetLindormInstanceResponseBody {
	s.VswitchId = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetZoneId(v string) *GetLindormInstanceResponseBody {
	s.ZoneId = &v
	return s
}

func (s *GetLindormInstanceResponseBody) Validate() error {
	if s.EngineList != nil {
		for _, item := range s.EngineList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetLindormInstanceResponseBodyEngineList struct {
	// The number of nodes in the arbiter zone.
	//
	// example:
	//
	// 2
	ArbiterCoreCount *string `json:"ArbiterCoreCount,omitempty" xml:"ArbiterCoreCount,omitempty"`
	// The number of engine nodes.
	//
	// example:
	//
	// 2
	CoreCount *string `json:"CoreCount,omitempty" xml:"CoreCount,omitempty"`
	// The number of vCPUs for the engine node.
	//
	// example:
	//
	// 4
	CpuCount *string `json:"CpuCount,omitempty" xml:"CpuCount,omitempty"`
	// The type of the engine. Valid values:
	//
	// - **lindorm**: the wide table engine.
	//
	// - **tsdb**: the time series engine.
	//
	// - **solr**: the search engine.
	//
	// - **store**: the file engine.
	//
	// - **bds**: the LTS engine.
	//
	// - **compute**: the compute engine.
	//
	// example:
	//
	// lindorm
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// Indicates whether the engine is of the latest version. Valid values:
	//
	// - **true**: The engine is of the latest version.
	//
	// - **false**: The engine is not of the latest version.
	//
	// example:
	//
	// false
	IsLastVersion *bool `json:"IsLastVersion,omitempty" xml:"IsLastVersion,omitempty"`
	// The latest version of the engine.
	//
	// example:
	//
	// 2.2.19.2
	LatestVersion *string `json:"LatestVersion,omitempty" xml:"LatestVersion,omitempty"`
	// The memory size of the engine node.
	//
	// example:
	//
	// 8GB
	MemorySize *string `json:"MemorySize,omitempty" xml:"MemorySize,omitempty"`
	// The number of nodes in the primary zone.
	//
	// example:
	//
	// 2
	PrimaryCoreCount *string `json:"PrimaryCoreCount,omitempty" xml:"PrimaryCoreCount,omitempty"`
	// The specification of the engine nodes.
	//
	// example:
	//
	// lindorm.g.2xlarge
	Specification *string `json:"Specification,omitempty" xml:"Specification,omitempty"`
	// The number of nodes in the secondary zone.
	//
	// example:
	//
	// 2
	StandbyCoreCount *string `json:"StandbyCoreCount,omitempty" xml:"StandbyCoreCount,omitempty"`
	// The version of the engine.
	//
	// example:
	//
	// 2.2.3
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetLindormInstanceResponseBodyEngineList) String() string {
	return dara.Prettify(s)
}

func (s GetLindormInstanceResponseBodyEngineList) GoString() string {
	return s.String()
}

func (s *GetLindormInstanceResponseBodyEngineList) GetArbiterCoreCount() *string {
	return s.ArbiterCoreCount
}

func (s *GetLindormInstanceResponseBodyEngineList) GetCoreCount() *string {
	return s.CoreCount
}

func (s *GetLindormInstanceResponseBodyEngineList) GetCpuCount() *string {
	return s.CpuCount
}

func (s *GetLindormInstanceResponseBodyEngineList) GetEngine() *string {
	return s.Engine
}

func (s *GetLindormInstanceResponseBodyEngineList) GetIsLastVersion() *bool {
	return s.IsLastVersion
}

func (s *GetLindormInstanceResponseBodyEngineList) GetLatestVersion() *string {
	return s.LatestVersion
}

func (s *GetLindormInstanceResponseBodyEngineList) GetMemorySize() *string {
	return s.MemorySize
}

func (s *GetLindormInstanceResponseBodyEngineList) GetPrimaryCoreCount() *string {
	return s.PrimaryCoreCount
}

func (s *GetLindormInstanceResponseBodyEngineList) GetSpecification() *string {
	return s.Specification
}

func (s *GetLindormInstanceResponseBodyEngineList) GetStandbyCoreCount() *string {
	return s.StandbyCoreCount
}

func (s *GetLindormInstanceResponseBodyEngineList) GetVersion() *string {
	return s.Version
}

func (s *GetLindormInstanceResponseBodyEngineList) SetArbiterCoreCount(v string) *GetLindormInstanceResponseBodyEngineList {
	s.ArbiterCoreCount = &v
	return s
}

func (s *GetLindormInstanceResponseBodyEngineList) SetCoreCount(v string) *GetLindormInstanceResponseBodyEngineList {
	s.CoreCount = &v
	return s
}

func (s *GetLindormInstanceResponseBodyEngineList) SetCpuCount(v string) *GetLindormInstanceResponseBodyEngineList {
	s.CpuCount = &v
	return s
}

func (s *GetLindormInstanceResponseBodyEngineList) SetEngine(v string) *GetLindormInstanceResponseBodyEngineList {
	s.Engine = &v
	return s
}

func (s *GetLindormInstanceResponseBodyEngineList) SetIsLastVersion(v bool) *GetLindormInstanceResponseBodyEngineList {
	s.IsLastVersion = &v
	return s
}

func (s *GetLindormInstanceResponseBodyEngineList) SetLatestVersion(v string) *GetLindormInstanceResponseBodyEngineList {
	s.LatestVersion = &v
	return s
}

func (s *GetLindormInstanceResponseBodyEngineList) SetMemorySize(v string) *GetLindormInstanceResponseBodyEngineList {
	s.MemorySize = &v
	return s
}

func (s *GetLindormInstanceResponseBodyEngineList) SetPrimaryCoreCount(v string) *GetLindormInstanceResponseBodyEngineList {
	s.PrimaryCoreCount = &v
	return s
}

func (s *GetLindormInstanceResponseBodyEngineList) SetSpecification(v string) *GetLindormInstanceResponseBodyEngineList {
	s.Specification = &v
	return s
}

func (s *GetLindormInstanceResponseBodyEngineList) SetStandbyCoreCount(v string) *GetLindormInstanceResponseBodyEngineList {
	s.StandbyCoreCount = &v
	return s
}

func (s *GetLindormInstanceResponseBodyEngineList) SetVersion(v string) *GetLindormInstanceResponseBodyEngineList {
	s.Version = &v
	return s
}

func (s *GetLindormInstanceResponseBodyEngineList) Validate() error {
	return dara.Validate(s)
}
