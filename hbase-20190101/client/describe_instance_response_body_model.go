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
	// example:
	//
	// false
	AutoRenewal *bool `json:"AutoRenewal,omitempty" xml:"AutoRenewal,omitempty"`
	// example:
	//
	// open
	BackupStatus *string `json:"BackupStatus,omitempty" xml:"BackupStatus,omitempty"`
	// example:
	//
	// ld-bp150tns0sjxs****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// testhbase
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// example:
	//
	// cluster
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// example:
	//
	// 800
	ColdStorageSize *int32 `json:"ColdStorageSize,omitempty" xml:"ColdStorageSize,omitempty"`
	// example:
	//
	// open
	ColdStorageStatus *string `json:"ColdStorageStatus,omitempty" xml:"ColdStorageStatus,omitempty"`
	// example:
	//
	// true
	ConfirmMaintainTime *string `json:"ConfirmMaintainTime,omitempty" xml:"ConfirmMaintainTime,omitempty"`
	// example:
	//
	// 4
	CoreDiskCount *string `json:"CoreDiskCount,omitempty" xml:"CoreDiskCount,omitempty"`
	// example:
	//
	// 100
	CoreDiskSize *int32 `json:"CoreDiskSize,omitempty" xml:"CoreDiskSize,omitempty"`
	// example:
	//
	// cloud_ssd
	CoreDiskType *string `json:"CoreDiskType,omitempty" xml:"CoreDiskType,omitempty"`
	// example:
	//
	// hbase.sn2.2xlarge
	CoreInstanceType *string `json:"CoreInstanceType,omitempty" xml:"CoreInstanceType,omitempty"`
	// example:
	//
	// 2
	CoreNodeCount *int32 `json:"CoreNodeCount,omitempty" xml:"CoreNodeCount,omitempty"`
	// example:
	//
	// 2021-07-19T11:23:22
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// example:
	//
	// 2021-07-19T03:23:22Z
	CreatedTimeUTC *string `json:"CreatedTimeUTC,omitempty" xml:"CreatedTimeUTC,omitempty"`
	// example:
	//
	// 12
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// false
	EnableHbaseProxy *bool `json:"EnableHbaseProxy,omitempty" xml:"EnableHbaseProxy,omitempty"`
	// example:
	//
	// 0d2470df-da7b-4786-b981-9a164dae****
	EncryptionKey *string `json:"EncryptionKey,omitempty" xml:"EncryptionKey,omitempty"`
	// example:
	//
	// NoEncryption
	EncryptionType *string `json:"EncryptionType,omitempty" xml:"EncryptionType,omitempty"`
	// example:
	//
	// hbaseue
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// example:
	//
	// 2022-02-24T00:00:00
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// 2022-02-23T16:00:00Z
	ExpireTimeUTC *string `json:"ExpireTimeUTC,omitempty" xml:"ExpireTimeUTC,omitempty"`
	// example:
	//
	// ld-bp150tns0sjxs****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// testhbase
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// false
	IsDeletionProtection *bool `json:"IsDeletionProtection,omitempty" xml:"IsDeletionProtection,omitempty"`
	// example:
	//
	// true
	IsHa *bool `json:"IsHa,omitempty" xml:"IsHa,omitempty"`
	// example:
	//
	// true
	IsLatestVersion *bool `json:"IsLatestVersion,omitempty" xml:"IsLatestVersion,omitempty"`
	// example:
	//
	// true
	IsMultiModel *bool `json:"IsMultiModel,omitempty" xml:"IsMultiModel,omitempty"`
	// example:
	//
	// 2.3.2
	LproxyMinorVersion *string `json:"LproxyMinorVersion,omitempty" xml:"LproxyMinorVersion,omitempty"`
	// example:
	//
	// 22:00Z
	MaintainEndTime *string `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	// example:
	//
	// 18:00Z
	MaintainStartTime *string `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	// example:
	//
	// 2.0
	MajorVersion *string `json:"MajorVersion,omitempty" xml:"MajorVersion,omitempty"`
	// example:
	//
	// 0
	MasterDiskSize *int32 `json:"MasterDiskSize,omitempty" xml:"MasterDiskSize,omitempty"`
	// example:
	//
	// cloud_efficiency
	MasterDiskType *string `json:"MasterDiskType,omitempty" xml:"MasterDiskType,omitempty"`
	// example:
	//
	// hbase.sn2.large
	MasterInstanceType *string `json:"MasterInstanceType,omitempty" xml:"MasterInstanceType,omitempty"`
	// example:
	//
	// 2
	MasterNodeCount *int32 `json:"MasterNodeCount,omitempty" xml:"MasterNodeCount,omitempty"`
	// example:
	//
	// 2.2.9.1
	MinorVersion *string `json:"MinorVersion,omitempty" xml:"MinorVersion,omitempty"`
	// example:
	//
	// 0
	ModuleId *int32 `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	// example:
	//
	// phoenxi:4.0
	ModuleStackVersion *string `json:"ModuleStackVersion,omitempty" xml:"ModuleStackVersion,omitempty"`
	// example:
	//
	// false
	NeedUpgrade      *bool                                         `json:"NeedUpgrade,omitempty" xml:"NeedUpgrade,omitempty"`
	NeedUpgradeComps *DescribeInstanceResponseBodyNeedUpgradeComps `json:"NeedUpgradeComps,omitempty" xml:"NeedUpgradeComps,omitempty" type:"Struct"`
	// example:
	//
	// VPC
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// example:
	//
	// ld-uf699153o1m2l****
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// example:
	//
	// Prepaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 3F429923-B6F6-52C5-9C2A-5B8A8C6BBA66
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// rg-acfmyiu4ekp****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// ACTIVATION
	Status *string                           `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags   *DescribeInstanceResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// example:
	//
	// 25.00
	TaskProgress *string `json:"TaskProgress,omitempty" xml:"TaskProgress,omitempty"`
	TaskStatus   *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// example:
	//
	// vpc-bp15s22y1a7sff5gj****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// example:
	//
	// vsw-bp1foll427ze3d4ps****
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
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
	// example:
	//
	// test_key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// test_value
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
