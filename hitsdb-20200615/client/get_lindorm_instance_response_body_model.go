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
	SetSingleZoneRiskAlert(v *GetLindormInstanceResponseBodySingleZoneRiskAlert) *GetLindormInstanceResponseBody
	GetSingleZoneRiskAlert() *GetLindormInstanceResponseBodySingleZoneRiskAlert
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
	// The 16-character AliUid of the Alibaba Cloud account.
	//
	// example:
	//
	// 164901546557****
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The ID of the vSwitch in the arbiter zone for a multi-zone instance. The vSwitch must belong to the zone specified by ArbiterZoneId.
	//
	// example:
	//
	// vsw-uf6664pqjawb87k36****
	ArbiterVSwitchId *string `json:"ArbiterVSwitchId,omitempty" xml:"ArbiterVSwitchId,omitempty"`
	// The ID of the arbiter zone for a multi-zone instance.
	//
	// example:
	//
	// cn-shanghai-g
	ArbiterZoneId *string `json:"ArbiterZoneId,omitempty" xml:"ArbiterZoneId,omitempty"`
	// The deployment architecture. Valid values:
	//
	// - **1.0**: single-zone.
	//
	// - **2.0**: multi-zone.
	//
	// example:
	//
	// 1.0
	ArchVersion *string `json:"ArchVersion,omitempty" xml:"ArchVersion,omitempty"`
	// The billable storage capacity of the archive storage type. Unit: GB.
	//
	// example:
	//
	// 0
	ArchiveStorage *int32 `json:"ArchiveStorage,omitempty" xml:"ArchiveStorage,omitempty"`
	// Indicates whether auto-renewal is enabled. Valid values:
	//
	// - **true**: enabled.
	//
	// - **false**: disabled.
	//
	// > This parameter is returned only when the billing method of the instance is subscription.
	//
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The backup instance.
	//
	// example:
	//
	// ld-xxxx
	BackupInstance *string `json:"BackupInstance,omitempty" xml:"BackupInstance,omitempty"`
	// Storage capacity of storage capacity cloud storage.
	//
	// example:
	//
	// 0
	ColdStorage *int32 `json:"ColdStorage,omitempty" xml:"ColdStorage,omitempty"`
	// The disk type of the core node for a multi-zone instance. Valid values:
	//
	// - **cloud_efficiency**: standard cloud storage.
	//
	// - **cloud_ssd**: performance cloud storage.
	//
	// - **cloud_essd**: enhanced performance cloud storage.
	//
	// - **cloud_essd_pl0**: performance cloud storage PL0.
	//
	// example:
	//
	// cloud_efficiency
	CoreDiskCategory *string `json:"CoreDiskCategory,omitempty" xml:"CoreDiskCategory,omitempty"`
	// The number of core nodes for a multi-zone instance.
	//
	// example:
	//
	// 4
	CoreNum *int32 `json:"CoreNum,omitempty" xml:"CoreNum,omitempty"`
	// The disk capacity of a single core node for a multi-zone instance.
	//
	// example:
	//
	// 400
	CoreSingleStorage *int32 `json:"CoreSingleStorage,omitempty" xml:"CoreSingleStorage,omitempty"`
	// The core node specifications for a multi-zone instance.
	//
	// example:
	//
	// lindorm.g.xlarge
	CoreSpec *string `json:"CoreSpec,omitempty" xml:"CoreSpec,omitempty"`
	// The creation time of the instance, represented as the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1627290664000
	CreateMilliseconds *int64 `json:"CreateMilliseconds,omitempty" xml:"CreateMilliseconds,omitempty"`
	// The time when the instance was created. The time is displayed in the **yyyy-MM-dd HH:mm:ss*	- format.
	//
	// example:
	//
	// 2021-07-26 17:10:26
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
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
	// The storage type. Valid values:
	//
	// - **cloud_efficiency**: standard cloud storage.
	//
	// - **cloud_ssd**: performance cloud storage.
	//
	// - **cloud_essd**: enhanced performance cloud storage.
	//
	// - **cloud_essd_pl0**: performance cloud storage PL0.
	//
	// - **capacity_cloud_storage**: capacity cloud storage.
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
	// The disk space usage.
	//
	// example:
	//
	// 0.0%
	DiskUsage *string `json:"DiskUsage,omitempty" xml:"DiskUsage,omitempty"`
	// Indicates whether LBlob is enabled for the instance. Valid values:
	//
	// - true: enabled.
	//
	// - false: disabled.
	//
	// example:
	//
	// true
	EnableBlob *bool `json:"EnableBlob,omitempty" xml:"EnableBlob,omitempty"`
	// Indicates whether the data subscription feature is enabled for the instance. Valid values:
	//
	// - **true**: enabled.
	//
	// - **false**: disabled.
	//
	// example:
	//
	// false
	EnableCdc *bool `json:"EnableCdc,omitempty" xml:"EnableCdc,omitempty"`
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
	// Indicates whether Key Management Service (KMS) is enabled. Valid values:
	//
	// - **true**: enabled.
	//
	// - **false**: disabled.
	//
	// example:
	//
	// false
	EnableKms *bool `json:"EnableKms,omitempty" xml:"EnableKms,omitempty"`
	// Indicates whether LindormTable supports the Thrift and CQL protocols. If not supported, you can call the SwitchLProxyService operation to enable or disable the support. Valid values:
	//
	// - True: supported.
	//
	// - False: not supported.
	//
	// example:
	//
	// False
	EnableLProxy *bool `json:"EnableLProxy,omitempty" xml:"EnableLProxy,omitempty"`
	// Indicates whether Lindorm Tunnel Service is enabled for the instance. Valid values:
	//
	// - **true**: enabled.
	//
	// - **false**: disabled.
	//
	// example:
	//
	// true
	EnableLTS *bool `json:"EnableLTS,omitempty" xml:"EnableLTS,omitempty"`
	// Indicates whether LindormTable supports LindormSQL-V3, which is compatible with the MySQL protocol. Instances purchased after October 24, 2023 support this feature by default. For existing instances, contact technical support for evaluation before enabling this feature. Valid values:
	//
	// - True: supported.
	//
	// - False: not supported.
	//
	// example:
	//
	// True
	EnableLsqlVersionV3 *bool `json:"EnableLsqlVersionV3,omitempty" xml:"EnableLsqlVersionV3,omitempty"`
	// Indicates whether the AI control node is enabled. Valid values:
	//
	// - True: enabled.
	//
	// - False: disabled.
	//
	// example:
	//
	// False
	EnableMLCtrl *bool `json:"EnableMLCtrl,omitempty" xml:"EnableMLCtrl,omitempty"`
	// Indicates whether SSL encryption is enabled. Valid values:
	//
	// - **true**: enabled.
	//
	// - **false**: disabled.
	//
	// example:
	//
	// false
	EnableSSL *bool `json:"EnableSSL,omitempty" xml:"EnableSSL,omitempty"`
	// Indicates whether History Server for the compute engine is enabled.
	//
	// example:
	//
	// true
	EnableShs *bool `json:"EnableShs,omitempty" xml:"EnableShs,omitempty"`
	// Indicates whether storage encryption is enabled. Valid values:
	//
	// - true: enabled.
	//
	// - false: disabled.
	//
	// example:
	//
	// false
	EnableStoreTDE *bool `json:"EnableStoreTDE,omitempty" xml:"EnableStoreTDE,omitempty"`
	// Indicates whether the streaming engine is enabled for the instance. Valid values:
	//
	// - **true**: enabled.
	//
	// - **false**: disabled.
	//
	// example:
	//
	// true
	EnableStream *bool `json:"EnableStream,omitempty" xml:"EnableStream,omitempty"`
	// The list of engine information.
	EngineList []*GetLindormInstanceResponseBodyEngineList `json:"EngineList,omitempty" xml:"EngineList,omitempty" type:"Repeated"`
	// The supported engine types. The return value is the sum of the values of the following engine types:
	//
	// - 1: search engine
	//
	// - 2: LindormTSDB
	//
	// - 4: LindormTable
	//
	// - 8: file engine
	//
	// > For example, if the value of EngineType is 15 (15 = 8 + 4 + 2 + 1), the instance supports the search engine, LindormTSDB, LindormTable, and file engine. If the value of EngineType is 6 (6 = 4 + 2), the instance supports LindormTSDB and LindormTable.
	//
	// example:
	//
	// 15
	EngineType *int32 `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// The expiration time of the instance. The time is displayed in the **yyyy-MM-dd HH:mm:ss*	- format.
	//
	// > This parameter is returned only when the billing method is subscription.
	//
	// example:
	//
	// 2021-08-27 00:00:00
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The time when the instance expires, represented as the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
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
	// The instance ID.
	//
	// example:
	//
	// ld-bp1o3y0yme2i2****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance status. Valid values:
	//
	// - **CREATING**: The instance is being created.
	//
	// - **ACTIVATION**: The instance is running.
	//
	// - **COLD_EXPANDING**: Capacity cloud storage expansion in progress.
	//
	// - **MINOR_VERSION_TRANSING**: The minor version is being upgraded.
	//
	// - **RESIZING**: The nodes are being scaled up.
	//
	// - **SHRINKING**: The nodes are being scaled down.
	//
	// - **CLASS_CHANGING**: Instance specification upgrade or downgrade in progress.
	//
	// - **SSL_SWITCHING**: The SSL certificate is being changed.
	//
	// - **CDC_OPENING**: The data subscription feature is being enabled.
	//
	// - **TRANSFER**: Data migration in progress.
	//
	// - **DATABASE_TRANSFER**: Data is being migrated to the database.
	//
	// - **GUARD_CREATING**: A disaster recovery instance is being created.
	//
	// - **BACKUP_RECOVERING**: The backup is being restored.
	//
	// - **DATABASE_IMPORTING**: Data import in progress.
	//
	// - **NET_MODIFYING**: The network is being modified.
	//
	// - **NET_SWITCHING**: The network is being switched between the internal network and the Internet.
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
	// The storage capacity of the instance.
	//
	// example:
	//
	// 480
	InstanceStorage *string `json:"InstanceStorage,omitempty" xml:"InstanceStorage,omitempty"`
	// The disk type of the log node for a multi-zone instance. Valid values:
	//
	// - **cloud_efficiency**: standard cloud storage.
	//
	// - **cloud_ssd**: performance cloud storage.
	//
	// example:
	//
	// cloud_ssd
	LogDiskCategory *string `json:"LogDiskCategory,omitempty" xml:"LogDiskCategory,omitempty"`
	// The number of log nodes for a multi-zone instance.
	//
	// example:
	//
	// 4
	LogNum *int32 `json:"LogNum,omitempty" xml:"LogNum,omitempty"`
	// The disk capacity of a single log node for a multi-zone instance.
	//
	// example:
	//
	// 400GB
	LogSingleStorage *int32 `json:"LogSingleStorage,omitempty" xml:"LogSingleStorage,omitempty"`
	// The log node specifications for a multi-zone instance.
	//
	// example:
	//
	// lindorm.sn1.large
	LogSpec *string `json:"LogSpec,omitempty" xml:"LogSpec,omitempty"`
	// The end time of the maintenance window. The time is displayed in the HH:mmZ format. UTC time is used.
	//
	// example:
	//
	// 20:00Z
	MaintainEndTime *string `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	// The start time of the maintenance window. The time is displayed in the HH:mmZ format. UTC time is used.
	//
	// example:
	//
	// 00:00Z
	MaintainStartTime *string `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	// The multi-zone combination. You can go to the buy page to view the supported zone combinations. Valid values:
	//
	// - **ap-southeast-5abc-aliyun**: Indonesia (Jakarta) Zones A, B, and C.
	//
	// - **cn-hangzhou-ehi-aliyun**: China (Hangzhou) Zones E, H, and I.
	//
	// - **cn-beijing-acd-aliyun**: China (Beijing) Zones A, C, and D.
	//
	// - **ap-southeast-1-abc-aliyun**: Singapore Zones A, B, and C.
	//
	// - **cn-zhangjiakou-abc-aliyun**: China (Zhangjiakou) Zones A, B, and C.
	//
	// - **cn-shanghai-efg-aliyun**: China (Shanghai) Zones E, F, and G.
	//
	// - **cn-shanghai-abd-aliyun**: China (Shanghai) Zones A, B, and D.
	//
	// - **cn-hangzhou-bef-aliyun**: China (Hangzhou) Zones B, E, and F.
	//
	// - **cn-hangzhou-bce-aliyun**: China (Hangzhou) Zones B, C, and E.
	//
	// - **cn-beijing-fgh-aliyun**: China (Beijing) Zones F, G, and H.
	//
	// - **cn-shenzhen-abc-aliyun**: China (Shenzhen) Zones A, B, and C.
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
	// - **PREPAY**: subscription.
	//
	// - **POSTPAY**: pay-as-you-go.
	//
	// example:
	//
	// POSTPAY
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The ID of the vSwitch in the primary zone for a multi-zone instance. The vSwitch must belong to the zone specified by PrimaryZoneId.
	//
	// example:
	//
	// vsw-uf6fdqa7c0pipnqzq****
	PrimaryVSwitchId *string `json:"PrimaryVSwitchId,omitempty" xml:"PrimaryVSwitchId,omitempty"`
	// The ID of the primary zone for a multi-zone instance.
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
	// The instance type. Valid values:
	//
	// - **lindorm**: Lindorm single-zone instance.
	//
	// - **lindorm_multizone**: Lindorm multi-zone instance.
	//
	// - **serverless_lindorm**: Lindorm Serverless instance.
	//
	// - **lindorm_standalone**: Lindorm single-node instance.
	//
	// - **lts**: Lindorm Tunnel Service.
	//
	// example:
	//
	// lindorm
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// The risk alert information for the single-zone deployment.
	SingleZoneRiskAlert *GetLindormInstanceResponseBodySingleZoneRiskAlert `json:"SingleZoneRiskAlert,omitempty" xml:"SingleZoneRiskAlert,omitempty" type:"Struct"`
	// The ID of the vSwitch in the secondary zone for a multi-zone instance. The vSwitch must belong to the zone specified by StandbyZoneId.
	//
	// example:
	//
	// vsw-2zec0kcn08cgdtr6****
	StandbyVSwitchId *string `json:"StandbyVSwitchId,omitempty" xml:"StandbyVSwitchId,omitempty"`
	// The ID of the secondary zone for a multi-zone instance.
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
	// The zone ID.
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

func (s *GetLindormInstanceResponseBody) GetSingleZoneRiskAlert() *GetLindormInstanceResponseBodySingleZoneRiskAlert {
	return s.SingleZoneRiskAlert
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

func (s *GetLindormInstanceResponseBody) SetSingleZoneRiskAlert(v *GetLindormInstanceResponseBodySingleZoneRiskAlert) *GetLindormInstanceResponseBody {
	s.SingleZoneRiskAlert = v
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
	if s.SingleZoneRiskAlert != nil {
		if err := s.SingleZoneRiskAlert.Validate(); err != nil {
			return err
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
	// The number of CPUs for the engine node.
	//
	// example:
	//
	// 4
	CpuCount *string `json:"CpuCount,omitempty" xml:"CpuCount,omitempty"`
	// The engine type. Valid values:
	//
	// - **lindorm**: LindormTable.
	//
	// - **tsdb**: LindormTSDB.
	//
	// - **solr**: search engine.
	//
	// - **store**: file engine.
	//
	// - **bds**: Lindorm Tunnel Service engine.
	//
	// - **compute**: compute engine.
	//
	// example:
	//
	// lindorm
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// Indicates whether the engine is of the latest version. Valid values:
	//
	// - **true**: latest version.
	//
	// - **false**: not the latest version.
	//
	// example:
	//
	// false
	IsLastVersion *bool `json:"IsLastVersion,omitempty" xml:"IsLastVersion,omitempty"`
	// The latest version number of the engine.
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
	// The engine node specifications.
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
	// The version number of the engine.
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

type GetLindormInstanceResponseBodySingleZoneRiskAlert struct {
	// The confirmation date for risk handling. The date is displayed in the yyyy-MM-dd format.
	//
	// example:
	//
	// 2026-09-24
	ConfirmDate *string `json:"ConfirmDate,omitempty" xml:"ConfirmDate,omitempty"`
	// The risk handling type for the single-zone deployment.
	//
	// example:
	//
	// NO_SET
	DispositionType *string `json:"DispositionType,omitempty" xml:"DispositionType,omitempty"`
	// Indicates whether a risk alert for the single-zone deployment is required.
	//
	// example:
	//
	// true
	NeedAlert *bool `json:"NeedAlert,omitempty" xml:"NeedAlert,omitempty"`
	// The planned completion date for risk handling. The date is displayed in the yyyy-MM-dd format.
	//
	// example:
	//
	// 2026-09-30
	PlannedCompletionDate *string `json:"PlannedCompletionDate,omitempty" xml:"PlannedCompletionDate,omitempty"`
}

func (s GetLindormInstanceResponseBodySingleZoneRiskAlert) String() string {
	return dara.Prettify(s)
}

func (s GetLindormInstanceResponseBodySingleZoneRiskAlert) GoString() string {
	return s.String()
}

func (s *GetLindormInstanceResponseBodySingleZoneRiskAlert) GetConfirmDate() *string {
	return s.ConfirmDate
}

func (s *GetLindormInstanceResponseBodySingleZoneRiskAlert) GetDispositionType() *string {
	return s.DispositionType
}

func (s *GetLindormInstanceResponseBodySingleZoneRiskAlert) GetNeedAlert() *bool {
	return s.NeedAlert
}

func (s *GetLindormInstanceResponseBodySingleZoneRiskAlert) GetPlannedCompletionDate() *string {
	return s.PlannedCompletionDate
}

func (s *GetLindormInstanceResponseBodySingleZoneRiskAlert) SetConfirmDate(v string) *GetLindormInstanceResponseBodySingleZoneRiskAlert {
	s.ConfirmDate = &v
	return s
}

func (s *GetLindormInstanceResponseBodySingleZoneRiskAlert) SetDispositionType(v string) *GetLindormInstanceResponseBodySingleZoneRiskAlert {
	s.DispositionType = &v
	return s
}

func (s *GetLindormInstanceResponseBodySingleZoneRiskAlert) SetNeedAlert(v bool) *GetLindormInstanceResponseBodySingleZoneRiskAlert {
	s.NeedAlert = &v
	return s
}

func (s *GetLindormInstanceResponseBodySingleZoneRiskAlert) SetPlannedCompletionDate(v string) *GetLindormInstanceResponseBodySingleZoneRiskAlert {
	s.PlannedCompletionDate = &v
	return s
}

func (s *GetLindormInstanceResponseBodySingleZoneRiskAlert) Validate() error {
	return dara.Validate(s)
}
