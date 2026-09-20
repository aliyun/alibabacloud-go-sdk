// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRenewal(v bool) *DescribeInstanceResponseBody
	GetAutoRenewal() *bool
	SetBackupStatus(v string) *DescribeInstanceResponseBody
	GetBackupStatus() *string
	SetClusterId(v string) *DescribeInstanceResponseBody
	GetClusterId() *string
	SetClusterName(v string) *DescribeInstanceResponseBody
	GetClusterName() *string
	SetClusterType(v string) *DescribeInstanceResponseBody
	GetClusterType() *string
	SetColdStorageSize(v int32) *DescribeInstanceResponseBody
	GetColdStorageSize() *int32
	SetColdStorageStatus(v string) *DescribeInstanceResponseBody
	GetColdStorageStatus() *string
	SetConfirmMaintainTime(v string) *DescribeInstanceResponseBody
	GetConfirmMaintainTime() *string
	SetCoreDiskCount(v string) *DescribeInstanceResponseBody
	GetCoreDiskCount() *string
	SetCoreDiskSize(v int32) *DescribeInstanceResponseBody
	GetCoreDiskSize() *int32
	SetCoreDiskType(v string) *DescribeInstanceResponseBody
	GetCoreDiskType() *string
	SetCoreInstanceType(v string) *DescribeInstanceResponseBody
	GetCoreInstanceType() *string
	SetCoreNodeCount(v int32) *DescribeInstanceResponseBody
	GetCoreNodeCount() *int32
	SetCreatedTime(v string) *DescribeInstanceResponseBody
	GetCreatedTime() *string
	SetCreatedTimeUTC(v string) *DescribeInstanceResponseBody
	GetCreatedTimeUTC() *string
	SetDuration(v int32) *DescribeInstanceResponseBody
	GetDuration() *int32
	SetEnableHbaseProxy(v bool) *DescribeInstanceResponseBody
	GetEnableHbaseProxy() *bool
	SetEncryptionKey(v string) *DescribeInstanceResponseBody
	GetEncryptionKey() *string
	SetEncryptionType(v string) *DescribeInstanceResponseBody
	GetEncryptionType() *string
	SetEngine(v string) *DescribeInstanceResponseBody
	GetEngine() *string
	SetExpireTime(v string) *DescribeInstanceResponseBody
	GetExpireTime() *string
	SetExpireTimeUTC(v string) *DescribeInstanceResponseBody
	GetExpireTimeUTC() *string
	SetInitialRootPassword(v string) *DescribeInstanceResponseBody
	GetInitialRootPassword() *string
	SetInstanceId(v string) *DescribeInstanceResponseBody
	GetInstanceId() *string
	SetInstanceName(v string) *DescribeInstanceResponseBody
	GetInstanceName() *string
	SetIsDeletionProtection(v bool) *DescribeInstanceResponseBody
	GetIsDeletionProtection() *bool
	SetIsHa(v bool) *DescribeInstanceResponseBody
	GetIsHa() *bool
	SetIsLatestVersion(v bool) *DescribeInstanceResponseBody
	GetIsLatestVersion() *bool
	SetIsMultiModel(v bool) *DescribeInstanceResponseBody
	GetIsMultiModel() *bool
	SetLproxyMinorVersion(v string) *DescribeInstanceResponseBody
	GetLproxyMinorVersion() *string
	SetMaintainEndTime(v string) *DescribeInstanceResponseBody
	GetMaintainEndTime() *string
	SetMaintainStartTime(v string) *DescribeInstanceResponseBody
	GetMaintainStartTime() *string
	SetMajorVersion(v string) *DescribeInstanceResponseBody
	GetMajorVersion() *string
	SetMasterDiskSize(v int32) *DescribeInstanceResponseBody
	GetMasterDiskSize() *int32
	SetMasterDiskType(v string) *DescribeInstanceResponseBody
	GetMasterDiskType() *string
	SetMasterInstanceType(v string) *DescribeInstanceResponseBody
	GetMasterInstanceType() *string
	SetMasterNodeCount(v int32) *DescribeInstanceResponseBody
	GetMasterNodeCount() *int32
	SetMinorVersion(v string) *DescribeInstanceResponseBody
	GetMinorVersion() *string
	SetModuleId(v int32) *DescribeInstanceResponseBody
	GetModuleId() *int32
	SetModuleStackVersion(v string) *DescribeInstanceResponseBody
	GetModuleStackVersion() *string
	SetNeedUpgrade(v bool) *DescribeInstanceResponseBody
	GetNeedUpgrade() *bool
	SetNeedUpgradeComps(v *DescribeInstanceResponseBodyNeedUpgradeComps) *DescribeInstanceResponseBody
	GetNeedUpgradeComps() *DescribeInstanceResponseBodyNeedUpgradeComps
	SetNetworkType(v string) *DescribeInstanceResponseBody
	GetNetworkType() *string
	SetParentId(v string) *DescribeInstanceResponseBody
	GetParentId() *string
	SetPayType(v string) *DescribeInstanceResponseBody
	GetPayType() *string
	SetRegionId(v string) *DescribeInstanceResponseBody
	GetRegionId() *string
	SetRequestId(v string) *DescribeInstanceResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *DescribeInstanceResponseBody
	GetResourceGroupId() *string
	SetSingleZoneRiskAlert(v *DescribeInstanceResponseBodySingleZoneRiskAlert) *DescribeInstanceResponseBody
	GetSingleZoneRiskAlert() *DescribeInstanceResponseBodySingleZoneRiskAlert
	SetStatus(v string) *DescribeInstanceResponseBody
	GetStatus() *string
	SetTags(v *DescribeInstanceResponseBodyTags) *DescribeInstanceResponseBody
	GetTags() *DescribeInstanceResponseBodyTags
	SetTaskProgress(v string) *DescribeInstanceResponseBody
	GetTaskProgress() *string
	SetTaskStatus(v string) *DescribeInstanceResponseBody
	GetTaskStatus() *string
	SetVpcId(v string) *DescribeInstanceResponseBody
	GetVpcId() *string
	SetVswitchId(v string) *DescribeInstanceResponseBody
	GetVswitchId() *string
	SetZoneId(v string) *DescribeInstanceResponseBody
	GetZoneId() *string
}

type DescribeInstanceResponseBody struct {
	// Indicates whether auto-renewal is enabled for the instance. Valid values:
	//
	// - **true**: Auto-renewal is enabled.
	//
	// - **false**: Auto-renewal is not enabled.
	//
	// > This parameter is returned only when PayType is set to Prepaid (subscription).
	//
	// example:
	//
	// false
	AutoRenewal *bool `json:"AutoRenewal,omitempty" xml:"AutoRenewal,omitempty"`
	// Indicates whether the backup feature is supported. Valid values:
	//
	// - **open**: The backup feature is supported.
	//
	// - **close**: The backup feature is not supported.
	//
	// example:
	//
	// open
	BackupStatus *string `json:"BackupStatus,omitempty" xml:"BackupStatus,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// ld-bp150tns0sjxs****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The instance name.
	//
	// example:
	//
	// testhbase
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The instance type. Valid values:
	//
	// - **cluster**: Cluster Edition.
	//
	// - **single**: single-node.
	//
	// example:
	//
	// cluster
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// The cold storage size. Unit: GB.
	//
	// example:
	//
	// 800
	ColdStorageSize *int32 `json:"ColdStorageSize,omitempty" xml:"ColdStorageSize,omitempty"`
	// Indicates whether the cold storage feature is supported. Valid values:
	//
	// - **open**: The cold storage feature is supported.
	//
	// - **close**: The cold storage feature is not supported.
	//
	// example:
	//
	// open
	ColdStorageStatus *string `json:"ColdStorageStatus,omitempty" xml:"ColdStorageStatus,omitempty"`
	// Indicates whether the O&M window of the instance has been confirmed for the first time. Valid values:
	//
	// - **true**: Confirmed.
	//
	// - **false**: Not confirmed.
	//
	// > The **Confirm the O&M window for the first time*	- dialog box appears only when you access the **Basic Information*	- page of the instance for the first time.
	//
	// example:
	//
	// true
	ConfirmMaintainTime *string `json:"ConfirmMaintainTime,omitempty" xml:"ConfirmMaintainTime,omitempty"`
	// The number of core node disks.
	//
	// example:
	//
	// 4
	CoreDiskCount *string `json:"CoreDiskCount,omitempty" xml:"CoreDiskCount,omitempty"`
	// The disk capacity of core nodes. Unit: GB.
	//
	// example:
	//
	// 100
	CoreDiskSize *int32 `json:"CoreDiskSize,omitempty" xml:"CoreDiskSize,omitempty"`
	// The disk type of core nodes. Valid values:
	//
	// - **cloud_efficiency**: ultra cloud disk.
	//
	// - **cloud_ssd**: standard SSD.
	//
	// - **local_hdd**: local HDD.
	//
	// - **local__ssd**: local SSD.
	//
	// example:
	//
	// cloud_ssd
	CoreDiskType *string `json:"CoreDiskType,omitempty" xml:"CoreDiskType,omitempty"`
	// The node specifications of core nodes.
	//
	// example:
	//
	// hbase.sn2.2xlarge
	CoreInstanceType *string `json:"CoreInstanceType,omitempty" xml:"CoreInstanceType,omitempty"`
	// The number of core nodes.
	//
	// example:
	//
	// 2
	CoreNodeCount *int32 `json:"CoreNodeCount,omitempty" xml:"CoreNodeCount,omitempty"`
	// The time when the instance was created.
	//
	// example:
	//
	// 2021-07-19T11:23:22
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The time when the instance was created, in UTC format.
	//
	// example:
	//
	// 2021-07-19T03:23:22Z
	CreatedTimeUTC *string `json:"CreatedTimeUTC,omitempty" xml:"CreatedTimeUTC,omitempty"`
	// The Unified Auto Renewal Cycle.
	//
	// - Monthly subscription: The auto-renewal epoch is 1 month.
	//
	// - Yearly subscription: The auto-renewal epoch is 1 year (12 months).
	//
	// > This parameter is returned only when PayType is set to Prepaid (subscription).
	//
	// example:
	//
	// 12
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// Indicates whether access from the HBase open source client is supported. Valid values:
	//
	// - **true**: Access is supported.
	//
	// - **false**: Access is not supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	EnableHbaseProxy *bool `json:"EnableHbaseProxy,omitempty" xml:"EnableHbaseProxy,omitempty"`
	// The encryption key.
	//
	// > This parameter is returned only when the encryption type is **CloudDisk**.
	//
	// example:
	//
	// 0d2470df-da7b-4786-b981-9a164dae****
	EncryptionKey *string `json:"EncryptionKey,omitempty" xml:"EncryptionKey,omitempty"`
	// The encryption type. Valid values:
	//
	// - **NoEncryption**: Encryption is not enabled.
	//
	// - **CloudDisk**: Cloud disk encryption is enabled.
	//
	// - **EncryptionKey**: The encryption key specified by the parameter.
	//
	// > Cloud disk encryption cannot be disabled after it is enabled.
	//
	// example:
	//
	// NoEncryption
	EncryptionType *string `json:"EncryptionType,omitempty" xml:"EncryptionType,omitempty"`
	// The database engine type. Valid values:
	//
	// - **hbase**: ApsaraDB for HBase Standard Edition or ApsaraDB for HBase single-node.
	//
	// - **hbaseue**: ApsaraDB for HBase Performance-enhanced Edition.
	//
	// - **serverlesshbase**: ApsaraDB for HBase Serverless Edition.
	//
	// - **bds**: BDS instance.
	//
	// example:
	//
	// hbaseue
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The time when the instance expires.
	//
	// example:
	//
	// 2022-02-24T00:00:00
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The time when the instance expires, in UTC format.
	//
	// example:
	//
	// 2022-02-23T16:00:00Z
	ExpireTimeUTC *string `json:"ExpireTimeUTC,omitempty" xml:"ExpireTimeUTC,omitempty"`
	// The initial default password.
	//
	// example:
	//
	// LFuVlAvSKsbo
	InitialRootPassword *string `json:"InitialRootPassword,omitempty" xml:"InitialRootPassword,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// ld-bp150tns0sjxs****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance name.
	//
	// example:
	//
	// testhbase
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// Indicates whether deletion protection is enabled. Valid values:
	//
	// - **true**: Enabled.
	//
	// - **false**: Not enabled.
	//
	// example:
	//
	// false
	IsDeletionProtection *bool `json:"IsDeletionProtection,omitempty" xml:"IsDeletionProtection,omitempty"`
	// Indicates whether the instance is configured for high availability. Valid values:
	//
	// - **true**: Configured for high availability.
	//
	// - **false**: Not configured for high availability.
	//
	// > - Cluster Edition instances are configured for high availability with default configurations and use 2 master nodes.
	//
	// - Single-node instances are configured with the actual active capacity.
	//
	// example:
	//
	// true
	IsHa *bool `json:"IsHa,omitempty" xml:"IsHa,omitempty"`
	// Indicates whether the instance is the latest version. Valid values:
	//
	// - **true**: The instance is the latest version.
	//
	// - **false**: The instance is not the latest version.
	//
	// example:
	//
	// true
	IsLatestVersion *bool `json:"IsLatestVersion,omitempty" xml:"IsLatestVersion,omitempty"`
	// Indicates whether the instance is a multi-model Cluster Edition instance. Valid values:
	//
	// - **true**: The instance is a multi-model Cluster Edition instance.
	//
	// - **false**: The instance is not a multi-model Cluster Edition instance.
	//
	// example:
	//
	// true
	IsMultiModel *bool `json:"IsMultiModel,omitempty" xml:"IsMultiModel,omitempty"`
	// The minor version of the LPROXY service.
	//
	// example:
	//
	// 2.3.2
	LproxyMinorVersion *string `json:"LproxyMinorVersion,omitempty" xml:"LproxyMinorVersion,omitempty"`
	// The end time of the O&M window.
	//
	// example:
	//
	// 22:00Z
	MaintainEndTime *string `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	// The start time of the O&M window.
	//
	// example:
	//
	// 18:00Z
	MaintainStartTime *string `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	// The major version number.
	//
	// example:
	//
	// 2.0
	MajorVersion *string `json:"MajorVersion,omitempty" xml:"MajorVersion,omitempty"`
	// The disk capacity of master nodes. Unit: GB.
	//
	// example:
	//
	// 0
	MasterDiskSize *int32 `json:"MasterDiskSize,omitempty" xml:"MasterDiskSize,omitempty"`
	// The disk type of master nodes. Valid values:
	//
	// - **cloud_efficiency**: ultra cloud disk.
	//
	// - **cloud_ssd**: standard SSD.
	//
	// > This parameter is returned for single-node instances.
	//
	// example:
	//
	// cloud_efficiency
	MasterDiskType *string `json:"MasterDiskType,omitempty" xml:"MasterDiskType,omitempty"`
	// The node specifications of master nodes.
	//
	// example:
	//
	// hbase.sn2.large
	MasterInstanceType *string `json:"MasterInstanceType,omitempty" xml:"MasterInstanceType,omitempty"`
	// The master node type. Valid values:
	//
	// - **0**: The master node is a single node.
	//
	// - **2**: The master node is in Cluster Edition.
	//
	// example:
	//
	// 2
	MasterNodeCount *int32 `json:"MasterNodeCount,omitempty" xml:"MasterNodeCount,omitempty"`
	// The minor version number of the instance.
	//
	// example:
	//
	// 2.2.9.1
	MinorVersion *string `json:"MinorVersion,omitempty" xml:"MinorVersion,omitempty"`
	// The module ID.
	//
	// example:
	//
	// 0
	ModuleId *int32 `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	// The module type version.
	//
	// example:
	//
	// phoenxi:4.0
	ModuleStackVersion *string `json:"ModuleStackVersion,omitempty" xml:"ModuleStackVersion,omitempty"`
	// Indicates whether the instance components need to be upgraded. Valid values:
	//
	// - **true**: Upgrade is required.
	//
	// - **false**: Upgrade is not required.
	//
	// example:
	//
	// false
	NeedUpgrade      *bool                                         `json:"NeedUpgrade,omitempty" xml:"NeedUpgrade,omitempty"`
	NeedUpgradeComps *DescribeInstanceResponseBodyNeedUpgradeComps `json:"NeedUpgradeComps,omitempty" xml:"NeedUpgradeComps,omitempty" type:"Struct"`
	// The network type. Valid values:
	//
	// - **VPC**: Virtual Private Cloud. If the network type is VPC, the VswitchId and VpcId parameters are returned.
	//
	// - **CLASSIC**: classic network.
	//
	// example:
	//
	// VPC
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The parent instance ID.
	//
	// example:
	//
	// ld-uf699153o1m2l****
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// - **Prepaid**: subscription.
	//
	// - **Postpaid**: pay-as-you-go.
	//
	// example:
	//
	// Prepaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The ID of the region in which the instance resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3F429923-B6F6-52C5-9C2A-5B8A8C6BBA66
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmyiu4ekp****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The single-zone risk alert information.
	SingleZoneRiskAlert *DescribeInstanceResponseBodySingleZoneRiskAlert `json:"SingleZoneRiskAlert,omitempty" xml:"SingleZoneRiskAlert,omitempty" type:"Struct"`
	// The instance status. Valid values:
	//
	// - **CREATING**: The instance is being created.
	//
	// - **ACTIVATION**: The instance is running.
	//
	// - **DELETING**: The instance is being deleted.
	//
	// - **RESTARTING**: The instance is being restarted.
	//
	// - **MINOR_VERSION_TRANSING**: A minor engine version update is in progress.
	//
	// example:
	//
	// ACTIVATION
	Status *string                           `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags   *DescribeInstanceResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The task progress of the instance, in percentage (%). Tasks initiated from the ApsaraDB for HBase console include specification changes, node scale-out, node scale-in, instance restart, and minor engine version updates.
	//
	// example:
	//
	// 25.00
	TaskProgress *string `json:"TaskProgress,omitempty" xml:"TaskProgress,omitempty"`
	// The task status. Valid values:
	//
	// - running: The task is running.
	//
	// - pause: The task is paused.
	//
	// - fail: The task is interrupted.
	//
	// - finish: The task is completed.
	//
	// example:
	//
	// running
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// The VPC ID. This parameter is returned when **NetworkType*	- is **2**.
	//
	// example:
	//
	// vpc-bp15s22y1a7sff5gj****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The vSwitch ID. This parameter is returned when **NetworkType*	- is **2**.
	//
	// example:
	//
	// vsw-bp1foll427ze3d4ps****
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-f
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBody) GetAutoRenewal() *bool {
	return s.AutoRenewal
}

func (s *DescribeInstanceResponseBody) GetBackupStatus() *string {
	return s.BackupStatus
}

func (s *DescribeInstanceResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeInstanceResponseBody) GetClusterName() *string {
	return s.ClusterName
}

func (s *DescribeInstanceResponseBody) GetClusterType() *string {
	return s.ClusterType
}

func (s *DescribeInstanceResponseBody) GetColdStorageSize() *int32 {
	return s.ColdStorageSize
}

func (s *DescribeInstanceResponseBody) GetColdStorageStatus() *string {
	return s.ColdStorageStatus
}

func (s *DescribeInstanceResponseBody) GetConfirmMaintainTime() *string {
	return s.ConfirmMaintainTime
}

func (s *DescribeInstanceResponseBody) GetCoreDiskCount() *string {
	return s.CoreDiskCount
}

func (s *DescribeInstanceResponseBody) GetCoreDiskSize() *int32 {
	return s.CoreDiskSize
}

func (s *DescribeInstanceResponseBody) GetCoreDiskType() *string {
	return s.CoreDiskType
}

func (s *DescribeInstanceResponseBody) GetCoreInstanceType() *string {
	return s.CoreInstanceType
}

func (s *DescribeInstanceResponseBody) GetCoreNodeCount() *int32 {
	return s.CoreNodeCount
}

func (s *DescribeInstanceResponseBody) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *DescribeInstanceResponseBody) GetCreatedTimeUTC() *string {
	return s.CreatedTimeUTC
}

func (s *DescribeInstanceResponseBody) GetDuration() *int32 {
	return s.Duration
}

func (s *DescribeInstanceResponseBody) GetEnableHbaseProxy() *bool {
	return s.EnableHbaseProxy
}

func (s *DescribeInstanceResponseBody) GetEncryptionKey() *string {
	return s.EncryptionKey
}

func (s *DescribeInstanceResponseBody) GetEncryptionType() *string {
	return s.EncryptionType
}

func (s *DescribeInstanceResponseBody) GetEngine() *string {
	return s.Engine
}

func (s *DescribeInstanceResponseBody) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeInstanceResponseBody) GetExpireTimeUTC() *string {
	return s.ExpireTimeUTC
}

func (s *DescribeInstanceResponseBody) GetInitialRootPassword() *string {
	return s.InitialRootPassword
}

func (s *DescribeInstanceResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceResponseBody) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeInstanceResponseBody) GetIsDeletionProtection() *bool {
	return s.IsDeletionProtection
}

func (s *DescribeInstanceResponseBody) GetIsHa() *bool {
	return s.IsHa
}

func (s *DescribeInstanceResponseBody) GetIsLatestVersion() *bool {
	return s.IsLatestVersion
}

func (s *DescribeInstanceResponseBody) GetIsMultiModel() *bool {
	return s.IsMultiModel
}

func (s *DescribeInstanceResponseBody) GetLproxyMinorVersion() *string {
	return s.LproxyMinorVersion
}

func (s *DescribeInstanceResponseBody) GetMaintainEndTime() *string {
	return s.MaintainEndTime
}

func (s *DescribeInstanceResponseBody) GetMaintainStartTime() *string {
	return s.MaintainStartTime
}

func (s *DescribeInstanceResponseBody) GetMajorVersion() *string {
	return s.MajorVersion
}

func (s *DescribeInstanceResponseBody) GetMasterDiskSize() *int32 {
	return s.MasterDiskSize
}

func (s *DescribeInstanceResponseBody) GetMasterDiskType() *string {
	return s.MasterDiskType
}

func (s *DescribeInstanceResponseBody) GetMasterInstanceType() *string {
	return s.MasterInstanceType
}

func (s *DescribeInstanceResponseBody) GetMasterNodeCount() *int32 {
	return s.MasterNodeCount
}

func (s *DescribeInstanceResponseBody) GetMinorVersion() *string {
	return s.MinorVersion
}

func (s *DescribeInstanceResponseBody) GetModuleId() *int32 {
	return s.ModuleId
}

func (s *DescribeInstanceResponseBody) GetModuleStackVersion() *string {
	return s.ModuleStackVersion
}

func (s *DescribeInstanceResponseBody) GetNeedUpgrade() *bool {
	return s.NeedUpgrade
}

func (s *DescribeInstanceResponseBody) GetNeedUpgradeComps() *DescribeInstanceResponseBodyNeedUpgradeComps {
	return s.NeedUpgradeComps
}

func (s *DescribeInstanceResponseBody) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeInstanceResponseBody) GetParentId() *string {
	return s.ParentId
}

func (s *DescribeInstanceResponseBody) GetPayType() *string {
	return s.PayType
}

func (s *DescribeInstanceResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeInstanceResponseBody) GetSingleZoneRiskAlert() *DescribeInstanceResponseBodySingleZoneRiskAlert {
	return s.SingleZoneRiskAlert
}

func (s *DescribeInstanceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeInstanceResponseBody) GetTags() *DescribeInstanceResponseBodyTags {
	return s.Tags
}

func (s *DescribeInstanceResponseBody) GetTaskProgress() *string {
	return s.TaskProgress
}

func (s *DescribeInstanceResponseBody) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *DescribeInstanceResponseBody) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeInstanceResponseBody) GetVswitchId() *string {
	return s.VswitchId
}

func (s *DescribeInstanceResponseBody) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeInstanceResponseBody) SetAutoRenewal(v bool) *DescribeInstanceResponseBody {
	s.AutoRenewal = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetBackupStatus(v string) *DescribeInstanceResponseBody {
	s.BackupStatus = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetClusterId(v string) *DescribeInstanceResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetClusterName(v string) *DescribeInstanceResponseBody {
	s.ClusterName = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetClusterType(v string) *DescribeInstanceResponseBody {
	s.ClusterType = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetColdStorageSize(v int32) *DescribeInstanceResponseBody {
	s.ColdStorageSize = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetColdStorageStatus(v string) *DescribeInstanceResponseBody {
	s.ColdStorageStatus = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetConfirmMaintainTime(v string) *DescribeInstanceResponseBody {
	s.ConfirmMaintainTime = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetCoreDiskCount(v string) *DescribeInstanceResponseBody {
	s.CoreDiskCount = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetCoreDiskSize(v int32) *DescribeInstanceResponseBody {
	s.CoreDiskSize = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetCoreDiskType(v string) *DescribeInstanceResponseBody {
	s.CoreDiskType = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetCoreInstanceType(v string) *DescribeInstanceResponseBody {
	s.CoreInstanceType = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetCoreNodeCount(v int32) *DescribeInstanceResponseBody {
	s.CoreNodeCount = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetCreatedTime(v string) *DescribeInstanceResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetCreatedTimeUTC(v string) *DescribeInstanceResponseBody {
	s.CreatedTimeUTC = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetDuration(v int32) *DescribeInstanceResponseBody {
	s.Duration = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetEnableHbaseProxy(v bool) *DescribeInstanceResponseBody {
	s.EnableHbaseProxy = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetEncryptionKey(v string) *DescribeInstanceResponseBody {
	s.EncryptionKey = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetEncryptionType(v string) *DescribeInstanceResponseBody {
	s.EncryptionType = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetEngine(v string) *DescribeInstanceResponseBody {
	s.Engine = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetExpireTime(v string) *DescribeInstanceResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetExpireTimeUTC(v string) *DescribeInstanceResponseBody {
	s.ExpireTimeUTC = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetInitialRootPassword(v string) *DescribeInstanceResponseBody {
	s.InitialRootPassword = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetInstanceId(v string) *DescribeInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetInstanceName(v string) *DescribeInstanceResponseBody {
	s.InstanceName = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetIsDeletionProtection(v bool) *DescribeInstanceResponseBody {
	s.IsDeletionProtection = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetIsHa(v bool) *DescribeInstanceResponseBody {
	s.IsHa = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetIsLatestVersion(v bool) *DescribeInstanceResponseBody {
	s.IsLatestVersion = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetIsMultiModel(v bool) *DescribeInstanceResponseBody {
	s.IsMultiModel = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetLproxyMinorVersion(v string) *DescribeInstanceResponseBody {
	s.LproxyMinorVersion = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetMaintainEndTime(v string) *DescribeInstanceResponseBody {
	s.MaintainEndTime = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetMaintainStartTime(v string) *DescribeInstanceResponseBody {
	s.MaintainStartTime = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetMajorVersion(v string) *DescribeInstanceResponseBody {
	s.MajorVersion = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetMasterDiskSize(v int32) *DescribeInstanceResponseBody {
	s.MasterDiskSize = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetMasterDiskType(v string) *DescribeInstanceResponseBody {
	s.MasterDiskType = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetMasterInstanceType(v string) *DescribeInstanceResponseBody {
	s.MasterInstanceType = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetMasterNodeCount(v int32) *DescribeInstanceResponseBody {
	s.MasterNodeCount = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetMinorVersion(v string) *DescribeInstanceResponseBody {
	s.MinorVersion = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetModuleId(v int32) *DescribeInstanceResponseBody {
	s.ModuleId = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetModuleStackVersion(v string) *DescribeInstanceResponseBody {
	s.ModuleStackVersion = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetNeedUpgrade(v bool) *DescribeInstanceResponseBody {
	s.NeedUpgrade = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetNeedUpgradeComps(v *DescribeInstanceResponseBodyNeedUpgradeComps) *DescribeInstanceResponseBody {
	s.NeedUpgradeComps = v
	return s
}

func (s *DescribeInstanceResponseBody) SetNetworkType(v string) *DescribeInstanceResponseBody {
	s.NetworkType = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetParentId(v string) *DescribeInstanceResponseBody {
	s.ParentId = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetPayType(v string) *DescribeInstanceResponseBody {
	s.PayType = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetRegionId(v string) *DescribeInstanceResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetRequestId(v string) *DescribeInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetResourceGroupId(v string) *DescribeInstanceResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetSingleZoneRiskAlert(v *DescribeInstanceResponseBodySingleZoneRiskAlert) *DescribeInstanceResponseBody {
	s.SingleZoneRiskAlert = v
	return s
}

func (s *DescribeInstanceResponseBody) SetStatus(v string) *DescribeInstanceResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetTags(v *DescribeInstanceResponseBodyTags) *DescribeInstanceResponseBody {
	s.Tags = v
	return s
}

func (s *DescribeInstanceResponseBody) SetTaskProgress(v string) *DescribeInstanceResponseBody {
	s.TaskProgress = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetTaskStatus(v string) *DescribeInstanceResponseBody {
	s.TaskStatus = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetVpcId(v string) *DescribeInstanceResponseBody {
	s.VpcId = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetVswitchId(v string) *DescribeInstanceResponseBody {
	s.VswitchId = &v
	return s
}

func (s *DescribeInstanceResponseBody) SetZoneId(v string) *DescribeInstanceResponseBody {
	s.ZoneId = &v
	return s
}

func (s *DescribeInstanceResponseBody) Validate() error {
	if s.NeedUpgradeComps != nil {
		if err := s.NeedUpgradeComps.Validate(); err != nil {
			return err
		}
	}
	if s.SingleZoneRiskAlert != nil {
		if err := s.SingleZoneRiskAlert.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInstanceResponseBodyNeedUpgradeComps struct {
	Comps []*string `json:"Comps,omitempty" xml:"Comps,omitempty" type:"Repeated"`
}

func (s DescribeInstanceResponseBodyNeedUpgradeComps) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceResponseBodyNeedUpgradeComps) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyNeedUpgradeComps) GetComps() []*string {
	return s.Comps
}

func (s *DescribeInstanceResponseBodyNeedUpgradeComps) SetComps(v []*string) *DescribeInstanceResponseBodyNeedUpgradeComps {
	s.Comps = v
	return s
}

func (s *DescribeInstanceResponseBodyNeedUpgradeComps) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceResponseBodySingleZoneRiskAlert struct {
	// The confirmation date.
	//
	// example:
	//
	// 2026-09-01
	ConfirmDate *string `json:"ConfirmDate,omitempty" xml:"ConfirmDate,omitempty"`
	// The disposition type.
	//
	// example:
	//
	// NO_SET: Not set.
	//
	// PLAN_MIGRATION: Plan to migrate to Lindorm multi-zone edition.
	//
	// EXTERNAL_BIZ_HA: Business-level disaster recovery
	DispositionType *string `json:"DispositionType,omitempty" xml:"DispositionType,omitempty"`
	// Indicates whether an alert is required.
	//
	// example:
	//
	// false
	NeedAlert *bool `json:"NeedAlert,omitempty" xml:"NeedAlert,omitempty"`
	// The planned completion date.
	//
	// example:
	//
	// 2027-01-01
	PlannedCompletionDate *string `json:"PlannedCompletionDate,omitempty" xml:"PlannedCompletionDate,omitempty"`
}

func (s DescribeInstanceResponseBodySingleZoneRiskAlert) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceResponseBodySingleZoneRiskAlert) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodySingleZoneRiskAlert) GetConfirmDate() *string {
	return s.ConfirmDate
}

func (s *DescribeInstanceResponseBodySingleZoneRiskAlert) GetDispositionType() *string {
	return s.DispositionType
}

func (s *DescribeInstanceResponseBodySingleZoneRiskAlert) GetNeedAlert() *bool {
	return s.NeedAlert
}

func (s *DescribeInstanceResponseBodySingleZoneRiskAlert) GetPlannedCompletionDate() *string {
	return s.PlannedCompletionDate
}

func (s *DescribeInstanceResponseBodySingleZoneRiskAlert) SetConfirmDate(v string) *DescribeInstanceResponseBodySingleZoneRiskAlert {
	s.ConfirmDate = &v
	return s
}

func (s *DescribeInstanceResponseBodySingleZoneRiskAlert) SetDispositionType(v string) *DescribeInstanceResponseBodySingleZoneRiskAlert {
	s.DispositionType = &v
	return s
}

func (s *DescribeInstanceResponseBodySingleZoneRiskAlert) SetNeedAlert(v bool) *DescribeInstanceResponseBodySingleZoneRiskAlert {
	s.NeedAlert = &v
	return s
}

func (s *DescribeInstanceResponseBodySingleZoneRiskAlert) SetPlannedCompletionDate(v string) *DescribeInstanceResponseBodySingleZoneRiskAlert {
	s.PlannedCompletionDate = &v
	return s
}

func (s *DescribeInstanceResponseBodySingleZoneRiskAlert) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceResponseBodyTags struct {
	Tag []*DescribeInstanceResponseBodyTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeInstanceResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceResponseBodyTags) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyTags) GetTag() []*DescribeInstanceResponseBodyTagsTag {
	return s.Tag
}

func (s *DescribeInstanceResponseBodyTags) SetTag(v []*DescribeInstanceResponseBodyTagsTag) *DescribeInstanceResponseBodyTags {
	s.Tag = v
	return s
}

func (s *DescribeInstanceResponseBodyTags) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstanceResponseBodyTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeInstanceResponseBodyTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceResponseBodyTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyTagsTag) GetKey() *string {
	return s.Key
}

func (s *DescribeInstanceResponseBodyTagsTag) GetValue() *string {
	return s.Value
}

func (s *DescribeInstanceResponseBodyTagsTag) SetKey(v string) *DescribeInstanceResponseBodyTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeInstanceResponseBodyTagsTag) SetValue(v string) *DescribeInstanceResponseBodyTagsTag {
	s.Value = &v
	return s
}

func (s *DescribeInstanceResponseBodyTagsTag) Validate() error {
	return dara.Validate(s)
}
