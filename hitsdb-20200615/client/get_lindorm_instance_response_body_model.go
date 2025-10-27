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
	// 16-digit AliUid of the Alibaba Cloud primary account (main account).
	//
	// example:
	//
	// 164901546557****
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// Multi-AZ instance, coordinating Availability Zone virtual switch ID, which must be located in the Availability Zone corresponding to ArbiterZoneId.
	//
	// example:
	//
	// vsw-uf6664pqjawb87k36****
	ArbiterVSwitchId *string `json:"ArbiterVSwitchId,omitempty" xml:"ArbiterVSwitchId,omitempty"`
	// Multi-zone instance, coordinating Availability Zone ID.
	//
	// example:
	//
	// cn-shanghai-g
	ArbiterZoneId *string `json:"ArbiterZoneId,omitempty" xml:"ArbiterZoneId,omitempty"`
	// The architecture of the instance. Valid values:
	//
	// 	- **1.0**: The instance is deployed in a single zone.
	//
	// 	- **2.0**: The instance is deployed across multiple zones.
	//
	// example:
	//
	// 1.0
	ArchVersion *string `json:"ArchVersion,omitempty" xml:"ArchVersion,omitempty"`
	// The Archive storage size of the instance.
	//
	// example:
	//
	// 0GB
	ArchiveStorage *int32 `json:"ArchiveStorage,omitempty" xml:"ArchiveStorage,omitempty"`
	// Indicates whether auto-renewal is enabled, with the following returns:
	//
	// - **true**: Enabled.
	//
	// - **false**: Disabled.
	//
	// > This parameter is returned when the instance\\"s payment type is prepaid.
	//
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The Capacity storage size of the instance.
	//
	// example:
	//
	// 0GB
	ColdStorage *int32 `json:"ColdStorage,omitempty" xml:"ColdStorage,omitempty"`
	// The disk type of the core nodes. This parameter is returned only for multi-zone instances. Valid values:
	//
	// 	- **cloud_efficiency**: This instance uses the Standard type of storage.
	//
	// 	- **cloud_ssd**: This instance uses the Performance type of storage.
	//
	// 	- **cloud_essd**: This instance uses ESSDs for storage.
	//
	// 	- **cloud_essd_pl0**: This instance uses PL0 ESSDs for storage.
	//
	// example:
	//
	// cloud_efficiency
	CoreDiskCategory *string `json:"CoreDiskCategory,omitempty" xml:"CoreDiskCategory,omitempty"`
	// Multi-zone instance, number of core nodes.
	//
	// example:
	//
	// 4
	CoreNum *int32 `json:"CoreNum,omitempty" xml:"CoreNum,omitempty"`
	// Multi-zone instance, core single-node disk capacity.
	//
	// example:
	//
	// 400
	CoreSingleStorage *int32 `json:"CoreSingleStorage,omitempty" xml:"CoreSingleStorage,omitempty"`
	// Multi-zone instance, core node specification.
	//
	// example:
	//
	// lindorm.g.xlarge
	CoreSpec *string `json:"CoreSpec,omitempty" xml:"CoreSpec,omitempty"`
	// The timestamp in milliseconds between the instance creation time and 1970-01-01 00:00:00.
	//
	// example:
	//
	// 1627290664000
	CreateMilliseconds *int64 `json:"CreateMilliseconds,omitempty" xml:"CreateMilliseconds,omitempty"`
	// The storage capacity of the disk of a single log node. This parameter is returned only for multi-zone instances.
	//
	// example:
	//
	// 2021-07-26 17:10:26
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether deletion protection is enabled, returning:
	//
	// - **true**: Enabled.
	//
	// - **false**: Disabled.
	//
	// example:
	//
	// false
	DeletionProtection *string `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// The storage type of the instance. Valid values:
	//
	// 	- **cloud_efficiency**: This instance uses the Standard type of storage.
	//
	// 	- **cloud_ssd**: This instance uses the Performance type of storage.
	//
	// 	- **cloud_essd**: This instance uses ESSDs for storage.
	//
	// 	- **cloud_essd_pl0**: This instance uses PL0 ESSDs for storage.
	//
	// 	- **capacity_cloud_storage**: This instance uses the Capacity type of storage.
	//
	// 	- **local_ssd_pro**: This instance uses local SSDs for storage.
	//
	// 	- **local_hdd_pro**: This instance uses local HDDs for storage.
	//
	// example:
	//
	// cloud_efficiency
	DiskCategory *string `json:"DiskCategory,omitempty" xml:"DiskCategory,omitempty"`
	// The threshold for disk space.
	//
	// example:
	//
	// 80%
	DiskThreshold *string `json:"DiskThreshold,omitempty" xml:"DiskThreshold,omitempty"`
	// Disk space usage rate.
	//
	// example:
	//
	// 0.0%
	DiskUsage *string `json:"DiskUsage,omitempty" xml:"DiskUsage,omitempty"`
	// Indicates whether LBlob is enabled for the instance. Valid values:
	//
	// true: LBlob is enabled for the instance. false: LBlob is not enabled for the instance.
	//
	// example:
	//
	// true
	EnableBlob *bool `json:"EnableBlob,omitempty" xml:"EnableBlob,omitempty"`
	// Indicates whether the data subscription feature for the instance is enabled. Returns:
	//
	// - **true**: Enabled.
	//
	// - **false**: Not enabled.
	//
	// example:
	//
	// false
	EnableCdc *bool `json:"EnableCdc,omitempty" xml:"EnableCdc,omitempty"`
	// Indicates whether the instance\\"s compute engine is enabled, returning:
	//
	// - **true**: Enabled.
	//
	// - **false**: Not enabled.
	//
	// example:
	//
	// true
	EnableCompute *bool `json:"EnableCompute,omitempty" xml:"EnableCompute,omitempty"`
	// Indicates whether the Key Management Service (KMS) is enabled, returning:
	//
	// - **true**: Enabled.
	//
	// - **false**: Disabled.
	//
	// example:
	//
	// false
	EnableKms *bool `json:"EnableKms,omitempty" xml:"EnableKms,omitempty"`
	// Indicates whether LindormTable supports the Thrift and CQL protocols. If these protocols are not supported. You can call the SwitchLProxyService operation to enable or disable the support on these protocols for LindormTable.
	//
	// True: LindormTable supports the Thrift and CQL protocols.
	//
	// False: LindormTable does not support the Thrift and CQL protocols.
	//
	// example:
	//
	// False
	EnableLProxy *bool `json:"EnableLProxy,omitempty" xml:"EnableLProxy,omitempty"`
	// Indicates whether the LTS engine is activated for the instance. Valid values:
	//
	// 	- **true**: The LTS engine is activated for the instance.
	//
	// 	- **false**: The LTS engine is not activated for the instance.
	//
	// example:
	//
	// true
	EnableLTS *bool `json:"EnableLTS,omitempty" xml:"EnableLTS,omitempty"`
	// Indicates whether LindormTable of the instance supports LindormSQL V3 that is compatible with MySQL. By default, LindormTable of instances that are purchased after October 24, 2023 supports LindormSQL V3. If your instance is purchased before this date and want to enable LindormSQL V3, contact the technical support.
	//
	// 	- True: LindormTable supports LindormSQL V3.
	//
	// 	- False: LindormTable does not support LindormSQL V3.
	//
	// example:
	//
	// True
	EnableLsqlVersionV3 *bool `json:"EnableLsqlVersionV3,omitempty" xml:"EnableLsqlVersionV3,omitempty"`
	// Indicates whether AI control nodes are enabled for the instance.
	//
	// 	- True: AI control nodes are enabled for the instance.
	//
	// 	- False: AI control nodes are not enabled for the instance.
	//
	// example:
	//
	// False
	EnableMLCtrl *bool `json:"EnableMLCtrl,omitempty" xml:"EnableMLCtrl,omitempty"`
	// Indicates whether SSL link encryption is enabled, returning:
	//
	// - **true**: Enabled.
	//
	// - **false**: Disabled.
	//
	// example:
	//
	// false
	EnableSSL *bool `json:"EnableSSL,omitempty" xml:"EnableSSL,omitempty"`
	// Whether to enable the Compute Engine History Server.
	//
	// example:
	//
	// true
	EnableShs *bool `json:"EnableShs,omitempty" xml:"EnableShs,omitempty"`
	// Indicates whether the Transparent Data Encryption (TDE) is enabled, returning:
	//
	// - **true**: Enabled.
	//
	// - **false**: Disabled.
	//
	// example:
	//
	// false
	EnableStoreTDE *bool `json:"EnableStoreTDE,omitempty" xml:"EnableStoreTDE,omitempty"`
	// Indicates whether the instance has the stream engine enabled. Return values:
	//
	// - **true**: Stream engine is enabled.
	//
	// - **false**: Stream engine is not enabled.
	//
	// example:
	//
	// true
	EnableStream *bool `json:"EnableStream,omitempty" xml:"EnableStream,omitempty"`
	// The list of engines supported by the instance.
	EngineList []*GetLindormInstanceResponseBodyEngineList `json:"EngineList,omitempty" xml:"EngineList,omitempty" type:"Repeated"`
	// Supported engine types, the return value is obtained by performing addition operations on the values of the following engine types.
	//
	// - 1: Search Engine
	//
	// - 2: Time Series Engine
	//
	// - 4: Wide Table Engine
	//
	// - 8: File Engine
	//
	// > For example: If EngineType is 15, where 15 = 8 + 4 + 2 + 1, it indicates that the instance supports Search Engine, Time Series Engine, Wide Table Engine, and File Engine. If EngineType is 6, where 6 = 4 + 2, it signifies that the instance supports Time Series Engine and Wide Table Engine.
	//
	// example:
	//
	// 15
	EngineType *int32 `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// Expiration time of the instance, format: **yyyy-MM-dd HH:mm:ss**.
	//
	// > This parameter is only returned when the payment type is pre-paid.
	//
	// example:
	//
	// 2021-08-27 00:00:00
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The millisecond value between the instance expiration time and 1970-01-01 00:00:00.
	//
	// example:
	//
	// 1629993600000
	ExpiredMilliseconds *int64 `json:"ExpiredMilliseconds,omitempty" xml:"ExpiredMilliseconds,omitempty"`
	// Instance name.
	//
	// example:
	//
	// test0726
	InstanceAlias *string `json:"InstanceAlias,omitempty" xml:"InstanceAlias,omitempty"`
	// Instance ID.
	//
	// example:
	//
	// ld-bp1o3y0yme2i2****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The status of the instance. Valid values:
	//
	// 	- **CREATING**: The instance is being created.
	//
	// 	- **ACTIVATION**: The instance is running.
	//
	// 	- **COLD_EXPANDING**: The Capacity storage of the instance is being scaled up.
	//
	// 	- **MINOR_VERSION_TRANSING**: The minor version of the instance is being updated.
	//
	// 	- **RESIZING**: The nodes in the instance are being scaled up.
	//
	// 	- **SHRINKING**: The nodes in the instance are being scaled down.
	//
	// 	- **CLASS_CHANGING**: The specification of the instance is being changed.
	//
	// 	- **SSL_SWITCHING: SSL**: The SSL configurations of the instance are being changed.
	//
	// 	- **CDC_OPENING**: Data subscription is being enabled for the instance.
	//
	// 	- **TRANSFER**: The data of the instance is being transferred.
	//
	// 	- **DATABASE_TRANSFER**: The data of the instance is being transferred to databases.
	//
	// 	- **GUARD_CREATING**: A disaster recovery instance is being created.
	//
	// 	- **BACKUP_RECOVERING**: The data of the instance is being restored from a backup.
	//
	// 	- **DATABASE_IMPORTING**: Data is being imported to the instance.
	//
	// 	- **NET_MODIFYING**: The network configurations of the instance are being changed.
	//
	// 	- **NET_SWITCHING**: The network of the instance is being switched between a virtual private cloud (VPC) and the Internet.
	//
	// 	- **NET_CREATING**: The connection to the instance is being created.
	//
	// 	- **NET_DELETING**: The connection to the instance is being deleted.
	//
	// 	- **DELETING**: The instance is being deleted.
	//
	// 	- **RESTARTING**: The instance is restarting.
	//
	// 	- **LOCKED**: The instance is locked because it expires.
	//
	// example:
	//
	// ACTIVATION
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// Instance\\"s storage capacity.
	//
	// example:
	//
	// 480
	InstanceStorage *string `json:"InstanceStorage,omitempty" xml:"InstanceStorage,omitempty"`
	// Multi-zone instance, log node disk type, returns:
	//
	// - **cloud_efficiency**: Standard cloud storage.
	//
	// - **cloud_ssd**: Performance cloud storage.
	//
	// example:
	//
	// cloud_ssd
	LogDiskCategory *string `json:"LogDiskCategory,omitempty" xml:"LogDiskCategory,omitempty"`
	// Multi-zone instance, number of log nodes.
	//
	// example:
	//
	// 4
	LogNum *int32 `json:"LogNum,omitempty" xml:"LogNum,omitempty"`
	// The storage capacity of the disk of a single log node. This parameter is returned only for multi-zone instances.
	//
	// example:
	//
	// 400GB
	LogSingleStorage *int32 `json:"LogSingleStorage,omitempty" xml:"LogSingleStorage,omitempty"`
	// Multi-zone instance, log node specification.
	//
	// example:
	//
	// lindorm.sn1.large
	LogSpec *string `json:"LogSpec,omitempty" xml:"LogSpec,omitempty"`
	// Maintainable end time.
	//
	// example:
	//
	// 20:00Z
	MaintainEndTime *string `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	// Maintainable start time.
	//
	// example:
	//
	// 00:00Z
	MaintainStartTime *string `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	// Multi-zone combinations. For support details on zone combinations, please refer to the product page.
	//
	// - **ap-southeast-5abc-aliyun**: Indonesia (Jakarta) A+B+C.
	//
	// - **cn-hangzhou-ehi-aliyun**: East China 1 (Hangzhou) E+H+I.
	//
	// - **cn-beijing-acd-aliyun**: North China 2 (Beijing) A+C+D.
	//
	// - **ap-southeast-1-abc-aliyun**: Singapore A+B+C.
	//
	// - **cn-zhangjiakou-abc-aliyun**: North China 3 (Zhangjiakou) A+B+C.
	//
	// - **cn-shanghai-efg-aliyun**: East China 2 (Shanghai) E+F+G.
	//
	// - **cn-shanghai-abd-aliyun**: East China 2 (Shanghai) A+B+D.
	//
	// - **cn-hangzhou-bef-aliyun**: East China 1 (Hangzhou) B+E+F.
	//
	// - **cn-hangzhou-bce-aliyun**: East China 1 (Hangzhou) B+C+E.
	//
	// - **cn-beijing-fgh-aliyun**: North China 2 (Beijing) F+G+H.
	//
	// - **cn-shenzhen-abc-aliyun**: South China 1 (Shenzhen) A+B+C.
	//
	// example:
	//
	// cn-shanghai-efg-aliyun
	MultiZoneCombination *string `json:"MultiZoneCombination,omitempty" xml:"MultiZoneCombination,omitempty"`
	// Instance\\"s network type.
	//
	// example:
	//
	// vpc
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// PREPAY: subscription.
	//
	// POSTPAY: pay-as-you-go.
	//
	// example:
	//
	// POSTPAY
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// Multi-zone instance, the virtual switch ID of the primary availability zone, which must be in the availability zone corresponding to PrimaryZoneId.
	//
	// example:
	//
	// vsw-uf6fdqa7c0pipnqzq****
	PrimaryVSwitchId *string `json:"PrimaryVSwitchId,omitempty" xml:"PrimaryVSwitchId,omitempty"`
	// Multi-zone instance, availability zone ID of the primary zone.
	//
	// example:
	//
	// cn-shanghai-e
	PrimaryZoneId *string `json:"PrimaryZoneId,omitempty" xml:"PrimaryZoneId,omitempty"`
	// Region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 633F1BE4-C8DA-5744-8FDF-A3075C3FE37F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Resource group ID.
	//
	// example:
	//
	// rg-aek2wvd6oia****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Instance type, valid values:
	//
	// - **lindorm**: represents a Lindorm single-zone instance.
	//
	// - **lindorm_multizone**: represents a Lindorm multi-zone instance.
	//
	// - **serverless_lindorm**: represents a Lindorm Serverless instance.
	//
	// - **lindorm_standalone**: represents a Lindorm standalone instance.
	//
	// - **lts**: represents the Lindorm Data Channel Service type.
	//
	// example:
	//
	// lindorm
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// Multi-zone instance, the virtual switch ID of the backup availability zone, which must be in the availability zone corresponding to StandbyZoneId.
	//
	// example:
	//
	// vsw-2zec0kcn08cgdtr6****
	StandbyVSwitchId *string `json:"StandbyVSwitchId,omitempty" xml:"StandbyVSwitchId,omitempty"`
	// Multi-zone instance, backup availability zone\\"s availability zone ID.
	//
	// example:
	//
	// cn-shanghai-f
	StandbyZoneId *string `json:"StandbyZoneId,omitempty" xml:"StandbyZoneId,omitempty"`
	// The type of the log nodes. This parameter is returned only for multi-zone instances.
	//
	// example:
	//
	// vpc-bp1n3i15v90el48nx****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The number of the log nodes. This parameter is returned only for multi-zone instances.
	//
	// example:
	//
	// vsw-bp1vbjzmod9q3l9eo****
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	// Availability Zone ID.
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
	ArbiterCoreCount *string `json:"ArbiterCoreCount,omitempty" xml:"ArbiterCoreCount,omitempty"`
	// The number of engine nodes.
	//
	// example:
	//
	// 2
	CoreCount *string `json:"CoreCount,omitempty" xml:"CoreCount,omitempty"`
	// The number of CPU cores on the engine node.
	//
	// example:
	//
	// 4
	CpuCount *string `json:"CpuCount,omitempty" xml:"CpuCount,omitempty"`
	// The engine type. Valid values:
	//
	// 	- **lindorm**: LindormTable.
	//
	// 	- **tsdb**: LindormTSDB.
	//
	// 	- **solr**: LindormSearch.
	//
	// 	- **store**: LindormDFS.
	//
	// 	- **bds**: Lindorm Tunnel Service (LTS).
	//
	// 	- **compute**: Lindorm Distributed Processing System (LDPS).
	//
	// example:
	//
	// lindorm
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// Indicates whether the version of the engine is the latest. Valid values:
	//
	// 	- **true**: The version of the engine is the latest.
	//
	// 	- **false**: The version of the engine is not the latest.
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
	// The memory size of the engine nodes.
	//
	// example:
	//
	// 8GB
	MemorySize       *string `json:"MemorySize,omitempty" xml:"MemorySize,omitempty"`
	PrimaryCoreCount *string `json:"PrimaryCoreCount,omitempty" xml:"PrimaryCoreCount,omitempty"`
	// The specification of the engine node.
	//
	// example:
	//
	// lindorm.g.2xlarge
	Specification    *string `json:"Specification,omitempty" xml:"Specification,omitempty"`
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
