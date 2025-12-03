// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMultiZoneClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetArbiterVSwitchIds(v string) *DescribeMultiZoneClusterResponseBody
	GetArbiterVSwitchIds() *string
	SetArbiterZoneId(v string) *DescribeMultiZoneClusterResponseBody
	GetArbiterZoneId() *string
	SetAutoRenewal(v bool) *DescribeMultiZoneClusterResponseBody
	GetAutoRenewal() *bool
	SetClusterId(v string) *DescribeMultiZoneClusterResponseBody
	GetClusterId() *string
	SetClusterName(v string) *DescribeMultiZoneClusterResponseBody
	GetClusterName() *string
	SetColdStorageSize(v int32) *DescribeMultiZoneClusterResponseBody
	GetColdStorageSize() *int32
	SetCoreDiskCount(v string) *DescribeMultiZoneClusterResponseBody
	GetCoreDiskCount() *string
	SetCoreDiskSize(v int32) *DescribeMultiZoneClusterResponseBody
	GetCoreDiskSize() *int32
	SetCoreDiskType(v string) *DescribeMultiZoneClusterResponseBody
	GetCoreDiskType() *string
	SetCoreInstanceType(v string) *DescribeMultiZoneClusterResponseBody
	GetCoreInstanceType() *string
	SetCoreNodeCount(v int32) *DescribeMultiZoneClusterResponseBody
	GetCoreNodeCount() *int32
	SetCreatedTime(v string) *DescribeMultiZoneClusterResponseBody
	GetCreatedTime() *string
	SetCreatedTimeUTC(v string) *DescribeMultiZoneClusterResponseBody
	GetCreatedTimeUTC() *string
	SetDuration(v int32) *DescribeMultiZoneClusterResponseBody
	GetDuration() *int32
	SetEncryptionKey(v string) *DescribeMultiZoneClusterResponseBody
	GetEncryptionKey() *string
	SetEncryptionType(v string) *DescribeMultiZoneClusterResponseBody
	GetEncryptionType() *string
	SetEngine(v string) *DescribeMultiZoneClusterResponseBody
	GetEngine() *string
	SetExpireTime(v string) *DescribeMultiZoneClusterResponseBody
	GetExpireTime() *string
	SetExpireTimeUTC(v string) *DescribeMultiZoneClusterResponseBody
	GetExpireTimeUTC() *string
	SetInstanceId(v string) *DescribeMultiZoneClusterResponseBody
	GetInstanceId() *string
	SetInstanceName(v string) *DescribeMultiZoneClusterResponseBody
	GetInstanceName() *string
	SetIsDeletionProtection(v bool) *DescribeMultiZoneClusterResponseBody
	GetIsDeletionProtection() *bool
	SetLogDiskCount(v string) *DescribeMultiZoneClusterResponseBody
	GetLogDiskCount() *string
	SetLogDiskSize(v int32) *DescribeMultiZoneClusterResponseBody
	GetLogDiskSize() *int32
	SetLogDiskType(v string) *DescribeMultiZoneClusterResponseBody
	GetLogDiskType() *string
	SetLogInstanceType(v string) *DescribeMultiZoneClusterResponseBody
	GetLogInstanceType() *string
	SetLogNodeCount(v int32) *DescribeMultiZoneClusterResponseBody
	GetLogNodeCount() *int32
	SetMaintainEndTime(v string) *DescribeMultiZoneClusterResponseBody
	GetMaintainEndTime() *string
	SetMaintainStartTime(v string) *DescribeMultiZoneClusterResponseBody
	GetMaintainStartTime() *string
	SetMajorVersion(v string) *DescribeMultiZoneClusterResponseBody
	GetMajorVersion() *string
	SetMasterDiskSize(v int32) *DescribeMultiZoneClusterResponseBody
	GetMasterDiskSize() *int32
	SetMasterDiskType(v string) *DescribeMultiZoneClusterResponseBody
	GetMasterDiskType() *string
	SetMasterInstanceType(v string) *DescribeMultiZoneClusterResponseBody
	GetMasterInstanceType() *string
	SetMasterNodeCount(v int32) *DescribeMultiZoneClusterResponseBody
	GetMasterNodeCount() *int32
	SetModuleId(v int32) *DescribeMultiZoneClusterResponseBody
	GetModuleId() *int32
	SetModuleStackVersion(v string) *DescribeMultiZoneClusterResponseBody
	GetModuleStackVersion() *string
	SetMultiZoneCombination(v string) *DescribeMultiZoneClusterResponseBody
	GetMultiZoneCombination() *string
	SetMultiZoneInstanceModels(v *DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModels) *DescribeMultiZoneClusterResponseBody
	GetMultiZoneInstanceModels() *DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModels
	SetNetworkType(v string) *DescribeMultiZoneClusterResponseBody
	GetNetworkType() *string
	SetParentId(v string) *DescribeMultiZoneClusterResponseBody
	GetParentId() *string
	SetPayType(v string) *DescribeMultiZoneClusterResponseBody
	GetPayType() *string
	SetPrimaryVSwitchIds(v string) *DescribeMultiZoneClusterResponseBody
	GetPrimaryVSwitchIds() *string
	SetPrimaryZoneId(v string) *DescribeMultiZoneClusterResponseBody
	GetPrimaryZoneId() *string
	SetRegionId(v string) *DescribeMultiZoneClusterResponseBody
	GetRegionId() *string
	SetRequestId(v string) *DescribeMultiZoneClusterResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *DescribeMultiZoneClusterResponseBody
	GetResourceGroupId() *string
	SetStandbyVSwitchIds(v string) *DescribeMultiZoneClusterResponseBody
	GetStandbyVSwitchIds() *string
	SetStandbyZoneId(v string) *DescribeMultiZoneClusterResponseBody
	GetStandbyZoneId() *string
	SetStatus(v string) *DescribeMultiZoneClusterResponseBody
	GetStatus() *string
	SetTags(v *DescribeMultiZoneClusterResponseBodyTags) *DescribeMultiZoneClusterResponseBody
	GetTags() *DescribeMultiZoneClusterResponseBodyTags
	SetTaskProgress(v string) *DescribeMultiZoneClusterResponseBody
	GetTaskProgress() *string
	SetTaskStatus(v string) *DescribeMultiZoneClusterResponseBody
	GetTaskStatus() *string
	SetVpcId(v string) *DescribeMultiZoneClusterResponseBody
	GetVpcId() *string
}

type DescribeMultiZoneClusterResponseBody struct {
	// example:
	//
	// vsw-t4nax9mp3wk0czn****
	ArbiterVSwitchIds *string `json:"ArbiterVSwitchIds,omitempty" xml:"ArbiterVSwitchIds,omitempty"`
	// example:
	//
	// ap-southeast-1c
	ArbiterZoneId *string `json:"ArbiterZoneId,omitempty" xml:"ArbiterZoneId,omitempty"`
	// example:
	//
	// false
	AutoRenewal *bool `json:"AutoRenewal,omitempty" xml:"AutoRenewal,omitempty"`
	// example:
	//
	// ld-t4nn71xa0yn****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// mz_test
	ClusterName     *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	ColdStorageSize *int32  `json:"ColdStorageSize,omitempty" xml:"ColdStorageSize,omitempty"`
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
	// cloud_efficiency
	CoreDiskType *string `json:"CoreDiskType,omitempty" xml:"CoreDiskType,omitempty"`
	// example:
	//
	// hbase.sn1.large
	CoreInstanceType *string `json:"CoreInstanceType,omitempty" xml:"CoreInstanceType,omitempty"`
	// example:
	//
	// 6
	CoreNodeCount *int32 `json:"CoreNodeCount,omitempty" xml:"CoreNodeCount,omitempty"`
	// example:
	//
	// 2020-10-15T18:04:52
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// example:
	//
	// 2020-10-15T10:04:52Z
	CreatedTimeUTC *string `json:"CreatedTimeUTC,omitempty" xml:"CreatedTimeUTC,omitempty"`
	// example:
	//
	// 1
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 2a****
	EncryptionKey *string `json:"EncryptionKey,omitempty" xml:"EncryptionKey,omitempty"`
	// example:
	//
	// CloudDisk
	EncryptionType *string `json:"EncryptionType,omitempty" xml:"EncryptionType,omitempty"`
	// example:
	//
	// hbaseue
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// example:
	//
	// 2020-11-16T08:00:00
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// 2020-11-16T00:00:00Z
	ExpireTimeUTC *string `json:"ExpireTimeUTC,omitempty" xml:"ExpireTimeUTC,omitempty"`
	// example:
	//
	// ld-t4nn71xa0yn****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// mz_test
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// false
	IsDeletionProtection *bool `json:"IsDeletionProtection,omitempty" xml:"IsDeletionProtection,omitempty"`
	// example:
	//
	// 4
	LogDiskCount *string `json:"LogDiskCount,omitempty" xml:"LogDiskCount,omitempty"`
	// example:
	//
	// 100
	LogDiskSize *int32 `json:"LogDiskSize,omitempty" xml:"LogDiskSize,omitempty"`
	// example:
	//
	// cloud_efficiency
	LogDiskType *string `json:"LogDiskType,omitempty" xml:"LogDiskType,omitempty"`
	// example:
	//
	// hbase.sn1.large
	LogInstanceType *string `json:"LogInstanceType,omitempty" xml:"LogInstanceType,omitempty"`
	// example:
	//
	// 4
	LogNodeCount *int32 `json:"LogNodeCount,omitempty" xml:"LogNodeCount,omitempty"`
	// example:
	//
	// 06:00:00
	MaintainEndTime *string `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	// example:
	//
	// 02:00:00
	MaintainStartTime *string `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	// example:
	//
	// 2.0
	MajorVersion *string `json:"MajorVersion,omitempty" xml:"MajorVersion,omitempty"`
	// example:
	//
	// 50
	MasterDiskSize *int32 `json:"MasterDiskSize,omitempty" xml:"MasterDiskSize,omitempty"`
	// example:
	//
	// cloud_efficiency
	MasterDiskType *string `json:"MasterDiskType,omitempty" xml:"MasterDiskType,omitempty"`
	// example:
	//
	// hbase.sn1.large
	MasterInstanceType *string `json:"MasterInstanceType,omitempty" xml:"MasterInstanceType,omitempty"`
	// example:
	//
	// 2
	MasterNodeCount *int32 `json:"MasterNodeCount,omitempty" xml:"MasterNodeCount,omitempty"`
	// example:
	//
	// 0
	ModuleId *int32 `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	// example:
	//
	// 2.0
	ModuleStackVersion *string `json:"ModuleStackVersion,omitempty" xml:"ModuleStackVersion,omitempty"`
	// example:
	//
	// ap-southeast-1-abc-aliyun
	MultiZoneCombination    *string                                                      `json:"MultiZoneCombination,omitempty" xml:"MultiZoneCombination,omitempty"`
	MultiZoneInstanceModels *DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModels `json:"MultiZoneInstanceModels,omitempty" xml:"MultiZoneInstanceModels,omitempty" type:"Struct"`
	// example:
	//
	// VPC
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// example:
	//
	// ld-fls1gf31y5s35****
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// example:
	//
	// Prepaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// example:
	//
	// vsw-t4n3s1zd2gtidg****
	PrimaryVSwitchIds *string `json:"PrimaryVSwitchIds,omitempty" xml:"PrimaryVSwitchIds,omitempty"`
	// example:
	//
	// ap-southeast-1a
	PrimaryZoneId *string `json:"PrimaryZoneId,omitempty" xml:"PrimaryZoneId,omitempty"`
	// example:
	//
	// ap-southeast-1
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// A02C0E6D-3A47-4FA0-BA7E-60793CE256DA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// rg-lk51f5fer315e****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// vsw-t4nvvk7xur3rdi****
	StandbyVSwitchIds *string `json:"StandbyVSwitchIds,omitempty" xml:"StandbyVSwitchIds,omitempty"`
	// example:
	//
	// ap-southeast-1b
	StandbyZoneId *string `json:"StandbyZoneId,omitempty" xml:"StandbyZoneId,omitempty"`
	// example:
	//
	// ACTIVATION
	Status       *string                                   `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags         *DescribeMultiZoneClusterResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	TaskProgress *string                                   `json:"TaskProgress,omitempty" xml:"TaskProgress,omitempty"`
	TaskStatus   *string                                   `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// example:
	//
	// vpc-t4nx81tmlixcq5****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeMultiZoneClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMultiZoneClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneClusterResponseBody) GetArbiterVSwitchIds() *string {
	return s.ArbiterVSwitchIds
}

func (s *DescribeMultiZoneClusterResponseBody) GetArbiterZoneId() *string {
	return s.ArbiterZoneId
}

func (s *DescribeMultiZoneClusterResponseBody) GetAutoRenewal() *bool {
	return s.AutoRenewal
}

func (s *DescribeMultiZoneClusterResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeMultiZoneClusterResponseBody) GetClusterName() *string {
	return s.ClusterName
}

func (s *DescribeMultiZoneClusterResponseBody) GetColdStorageSize() *int32 {
	return s.ColdStorageSize
}

func (s *DescribeMultiZoneClusterResponseBody) GetCoreDiskCount() *string {
	return s.CoreDiskCount
}

func (s *DescribeMultiZoneClusterResponseBody) GetCoreDiskSize() *int32 {
	return s.CoreDiskSize
}

func (s *DescribeMultiZoneClusterResponseBody) GetCoreDiskType() *string {
	return s.CoreDiskType
}

func (s *DescribeMultiZoneClusterResponseBody) GetCoreInstanceType() *string {
	return s.CoreInstanceType
}

func (s *DescribeMultiZoneClusterResponseBody) GetCoreNodeCount() *int32 {
	return s.CoreNodeCount
}

func (s *DescribeMultiZoneClusterResponseBody) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *DescribeMultiZoneClusterResponseBody) GetCreatedTimeUTC() *string {
	return s.CreatedTimeUTC
}

func (s *DescribeMultiZoneClusterResponseBody) GetDuration() *int32 {
	return s.Duration
}

func (s *DescribeMultiZoneClusterResponseBody) GetEncryptionKey() *string {
	return s.EncryptionKey
}

func (s *DescribeMultiZoneClusterResponseBody) GetEncryptionType() *string {
	return s.EncryptionType
}

func (s *DescribeMultiZoneClusterResponseBody) GetEngine() *string {
	return s.Engine
}

func (s *DescribeMultiZoneClusterResponseBody) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeMultiZoneClusterResponseBody) GetExpireTimeUTC() *string {
	return s.ExpireTimeUTC
}

func (s *DescribeMultiZoneClusterResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeMultiZoneClusterResponseBody) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeMultiZoneClusterResponseBody) GetIsDeletionProtection() *bool {
	return s.IsDeletionProtection
}

func (s *DescribeMultiZoneClusterResponseBody) GetLogDiskCount() *string {
	return s.LogDiskCount
}

func (s *DescribeMultiZoneClusterResponseBody) GetLogDiskSize() *int32 {
	return s.LogDiskSize
}

func (s *DescribeMultiZoneClusterResponseBody) GetLogDiskType() *string {
	return s.LogDiskType
}

func (s *DescribeMultiZoneClusterResponseBody) GetLogInstanceType() *string {
	return s.LogInstanceType
}

func (s *DescribeMultiZoneClusterResponseBody) GetLogNodeCount() *int32 {
	return s.LogNodeCount
}

func (s *DescribeMultiZoneClusterResponseBody) GetMaintainEndTime() *string {
	return s.MaintainEndTime
}

func (s *DescribeMultiZoneClusterResponseBody) GetMaintainStartTime() *string {
	return s.MaintainStartTime
}

func (s *DescribeMultiZoneClusterResponseBody) GetMajorVersion() *string {
	return s.MajorVersion
}

func (s *DescribeMultiZoneClusterResponseBody) GetMasterDiskSize() *int32 {
	return s.MasterDiskSize
}

func (s *DescribeMultiZoneClusterResponseBody) GetMasterDiskType() *string {
	return s.MasterDiskType
}

func (s *DescribeMultiZoneClusterResponseBody) GetMasterInstanceType() *string {
	return s.MasterInstanceType
}

func (s *DescribeMultiZoneClusterResponseBody) GetMasterNodeCount() *int32 {
	return s.MasterNodeCount
}

func (s *DescribeMultiZoneClusterResponseBody) GetModuleId() *int32 {
	return s.ModuleId
}

func (s *DescribeMultiZoneClusterResponseBody) GetModuleStackVersion() *string {
	return s.ModuleStackVersion
}

func (s *DescribeMultiZoneClusterResponseBody) GetMultiZoneCombination() *string {
	return s.MultiZoneCombination
}

func (s *DescribeMultiZoneClusterResponseBody) GetMultiZoneInstanceModels() *DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModels {
	return s.MultiZoneInstanceModels
}

func (s *DescribeMultiZoneClusterResponseBody) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeMultiZoneClusterResponseBody) GetParentId() *string {
	return s.ParentId
}

func (s *DescribeMultiZoneClusterResponseBody) GetPayType() *string {
	return s.PayType
}

func (s *DescribeMultiZoneClusterResponseBody) GetPrimaryVSwitchIds() *string {
	return s.PrimaryVSwitchIds
}

func (s *DescribeMultiZoneClusterResponseBody) GetPrimaryZoneId() *string {
	return s.PrimaryZoneId
}

func (s *DescribeMultiZoneClusterResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeMultiZoneClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMultiZoneClusterResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeMultiZoneClusterResponseBody) GetStandbyVSwitchIds() *string {
	return s.StandbyVSwitchIds
}

func (s *DescribeMultiZoneClusterResponseBody) GetStandbyZoneId() *string {
	return s.StandbyZoneId
}

func (s *DescribeMultiZoneClusterResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeMultiZoneClusterResponseBody) GetTags() *DescribeMultiZoneClusterResponseBodyTags {
	return s.Tags
}

func (s *DescribeMultiZoneClusterResponseBody) GetTaskProgress() *string {
	return s.TaskProgress
}

func (s *DescribeMultiZoneClusterResponseBody) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *DescribeMultiZoneClusterResponseBody) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeMultiZoneClusterResponseBody) SetArbiterVSwitchIds(v string) *DescribeMultiZoneClusterResponseBody {
	s.ArbiterVSwitchIds = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetArbiterZoneId(v string) *DescribeMultiZoneClusterResponseBody {
	s.ArbiterZoneId = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetAutoRenewal(v bool) *DescribeMultiZoneClusterResponseBody {
	s.AutoRenewal = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetClusterId(v string) *DescribeMultiZoneClusterResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetClusterName(v string) *DescribeMultiZoneClusterResponseBody {
	s.ClusterName = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetColdStorageSize(v int32) *DescribeMultiZoneClusterResponseBody {
	s.ColdStorageSize = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetCoreDiskCount(v string) *DescribeMultiZoneClusterResponseBody {
	s.CoreDiskCount = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetCoreDiskSize(v int32) *DescribeMultiZoneClusterResponseBody {
	s.CoreDiskSize = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetCoreDiskType(v string) *DescribeMultiZoneClusterResponseBody {
	s.CoreDiskType = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetCoreInstanceType(v string) *DescribeMultiZoneClusterResponseBody {
	s.CoreInstanceType = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetCoreNodeCount(v int32) *DescribeMultiZoneClusterResponseBody {
	s.CoreNodeCount = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetCreatedTime(v string) *DescribeMultiZoneClusterResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetCreatedTimeUTC(v string) *DescribeMultiZoneClusterResponseBody {
	s.CreatedTimeUTC = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetDuration(v int32) *DescribeMultiZoneClusterResponseBody {
	s.Duration = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetEncryptionKey(v string) *DescribeMultiZoneClusterResponseBody {
	s.EncryptionKey = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetEncryptionType(v string) *DescribeMultiZoneClusterResponseBody {
	s.EncryptionType = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetEngine(v string) *DescribeMultiZoneClusterResponseBody {
	s.Engine = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetExpireTime(v string) *DescribeMultiZoneClusterResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetExpireTimeUTC(v string) *DescribeMultiZoneClusterResponseBody {
	s.ExpireTimeUTC = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetInstanceId(v string) *DescribeMultiZoneClusterResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetInstanceName(v string) *DescribeMultiZoneClusterResponseBody {
	s.InstanceName = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetIsDeletionProtection(v bool) *DescribeMultiZoneClusterResponseBody {
	s.IsDeletionProtection = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetLogDiskCount(v string) *DescribeMultiZoneClusterResponseBody {
	s.LogDiskCount = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetLogDiskSize(v int32) *DescribeMultiZoneClusterResponseBody {
	s.LogDiskSize = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetLogDiskType(v string) *DescribeMultiZoneClusterResponseBody {
	s.LogDiskType = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetLogInstanceType(v string) *DescribeMultiZoneClusterResponseBody {
	s.LogInstanceType = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetLogNodeCount(v int32) *DescribeMultiZoneClusterResponseBody {
	s.LogNodeCount = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetMaintainEndTime(v string) *DescribeMultiZoneClusterResponseBody {
	s.MaintainEndTime = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetMaintainStartTime(v string) *DescribeMultiZoneClusterResponseBody {
	s.MaintainStartTime = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetMajorVersion(v string) *DescribeMultiZoneClusterResponseBody {
	s.MajorVersion = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetMasterDiskSize(v int32) *DescribeMultiZoneClusterResponseBody {
	s.MasterDiskSize = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetMasterDiskType(v string) *DescribeMultiZoneClusterResponseBody {
	s.MasterDiskType = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetMasterInstanceType(v string) *DescribeMultiZoneClusterResponseBody {
	s.MasterInstanceType = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetMasterNodeCount(v int32) *DescribeMultiZoneClusterResponseBody {
	s.MasterNodeCount = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetModuleId(v int32) *DescribeMultiZoneClusterResponseBody {
	s.ModuleId = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetModuleStackVersion(v string) *DescribeMultiZoneClusterResponseBody {
	s.ModuleStackVersion = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetMultiZoneCombination(v string) *DescribeMultiZoneClusterResponseBody {
	s.MultiZoneCombination = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetMultiZoneInstanceModels(v *DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModels) *DescribeMultiZoneClusterResponseBody {
	s.MultiZoneInstanceModels = v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetNetworkType(v string) *DescribeMultiZoneClusterResponseBody {
	s.NetworkType = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetParentId(v string) *DescribeMultiZoneClusterResponseBody {
	s.ParentId = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetPayType(v string) *DescribeMultiZoneClusterResponseBody {
	s.PayType = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetPrimaryVSwitchIds(v string) *DescribeMultiZoneClusterResponseBody {
	s.PrimaryVSwitchIds = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetPrimaryZoneId(v string) *DescribeMultiZoneClusterResponseBody {
	s.PrimaryZoneId = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetRegionId(v string) *DescribeMultiZoneClusterResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetRequestId(v string) *DescribeMultiZoneClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetResourceGroupId(v string) *DescribeMultiZoneClusterResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetStandbyVSwitchIds(v string) *DescribeMultiZoneClusterResponseBody {
	s.StandbyVSwitchIds = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetStandbyZoneId(v string) *DescribeMultiZoneClusterResponseBody {
	s.StandbyZoneId = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetStatus(v string) *DescribeMultiZoneClusterResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetTags(v *DescribeMultiZoneClusterResponseBodyTags) *DescribeMultiZoneClusterResponseBody {
	s.Tags = v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetTaskProgress(v string) *DescribeMultiZoneClusterResponseBody {
	s.TaskProgress = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetTaskStatus(v string) *DescribeMultiZoneClusterResponseBody {
	s.TaskStatus = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) SetVpcId(v string) *DescribeMultiZoneClusterResponseBody {
	s.VpcId = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBody) Validate() error {
	if s.MultiZoneInstanceModels != nil {
		if err := s.MultiZoneInstanceModels.Validate(); err != nil {
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

type DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModels struct {
	MultiZoneInstanceModel []*DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModelsMultiZoneInstanceModel `json:"MultiZoneInstanceModel,omitempty" xml:"MultiZoneInstanceModel,omitempty" type:"Repeated"`
}

func (s DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModels) String() string {
	return dara.Prettify(s)
}

func (s DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModels) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModels) GetMultiZoneInstanceModel() []*DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModelsMultiZoneInstanceModel {
	return s.MultiZoneInstanceModel
}

func (s *DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModels) SetMultiZoneInstanceModel(v []*DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModelsMultiZoneInstanceModel) *DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModels {
	s.MultiZoneInstanceModel = v
	return s
}

func (s *DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModels) Validate() error {
	if s.MultiZoneInstanceModel != nil {
		for _, item := range s.MultiZoneInstanceModel {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModelsMultiZoneInstanceModel struct {
	HdfsMinorVersion *string `json:"HdfsMinorVersion,omitempty" xml:"HdfsMinorVersion,omitempty"`
	// example:
	//
	// ld-t4nn71xa0yn****-az-a
	InsName             *string `json:"InsName,omitempty" xml:"InsName,omitempty"`
	IsHdfsLatestVersion *string `json:"IsHdfsLatestVersion,omitempty" xml:"IsHdfsLatestVersion,omitempty"`
	// example:
	//
	// true
	IsLatestVersion        *bool   `json:"IsLatestVersion,omitempty" xml:"IsLatestVersion,omitempty"`
	LatestHdfsMinorVersion *string `json:"LatestHdfsMinorVersion,omitempty" xml:"LatestHdfsMinorVersion,omitempty"`
	LatestMinorVersion     *string `json:"LatestMinorVersion,omitempty" xml:"LatestMinorVersion,omitempty"`
	// example:
	//
	// 2.1.24
	MinorVersion *string `json:"MinorVersion,omitempty" xml:"MinorVersion,omitempty"`
	// example:
	//
	// primary
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// example:
	//
	// ACTIVATION
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModelsMultiZoneInstanceModel) String() string {
	return dara.Prettify(s)
}

func (s DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModelsMultiZoneInstanceModel) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModelsMultiZoneInstanceModel) GetHdfsMinorVersion() *string {
	return s.HdfsMinorVersion
}

func (s *DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModelsMultiZoneInstanceModel) GetInsName() *string {
	return s.InsName
}

func (s *DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModelsMultiZoneInstanceModel) GetIsHdfsLatestVersion() *string {
	return s.IsHdfsLatestVersion
}

func (s *DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModelsMultiZoneInstanceModel) GetIsLatestVersion() *bool {
	return s.IsLatestVersion
}

func (s *DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModelsMultiZoneInstanceModel) GetLatestHdfsMinorVersion() *string {
	return s.LatestHdfsMinorVersion
}

func (s *DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModelsMultiZoneInstanceModel) GetLatestMinorVersion() *string {
	return s.LatestMinorVersion
}

func (s *DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModelsMultiZoneInstanceModel) GetMinorVersion() *string {
	return s.MinorVersion
}

func (s *DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModelsMultiZoneInstanceModel) GetRole() *string {
	return s.Role
}

func (s *DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModelsMultiZoneInstanceModel) GetStatus() *string {
	return s.Status
}

func (s *DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModelsMultiZoneInstanceModel) SetHdfsMinorVersion(v string) *DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModelsMultiZoneInstanceModel {
	s.HdfsMinorVersion = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModelsMultiZoneInstanceModel) SetInsName(v string) *DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModelsMultiZoneInstanceModel {
	s.InsName = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModelsMultiZoneInstanceModel) SetIsHdfsLatestVersion(v string) *DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModelsMultiZoneInstanceModel {
	s.IsHdfsLatestVersion = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModelsMultiZoneInstanceModel) SetIsLatestVersion(v bool) *DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModelsMultiZoneInstanceModel {
	s.IsLatestVersion = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModelsMultiZoneInstanceModel) SetLatestHdfsMinorVersion(v string) *DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModelsMultiZoneInstanceModel {
	s.LatestHdfsMinorVersion = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModelsMultiZoneInstanceModel) SetLatestMinorVersion(v string) *DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModelsMultiZoneInstanceModel {
	s.LatestMinorVersion = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModelsMultiZoneInstanceModel) SetMinorVersion(v string) *DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModelsMultiZoneInstanceModel {
	s.MinorVersion = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModelsMultiZoneInstanceModel) SetRole(v string) *DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModelsMultiZoneInstanceModel {
	s.Role = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModelsMultiZoneInstanceModel) SetStatus(v string) *DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModelsMultiZoneInstanceModel {
	s.Status = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBodyMultiZoneInstanceModelsMultiZoneInstanceModel) Validate() error {
	return dara.Validate(s)
}

type DescribeMultiZoneClusterResponseBodyTags struct {
	Tag []*DescribeMultiZoneClusterResponseBodyTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeMultiZoneClusterResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeMultiZoneClusterResponseBodyTags) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneClusterResponseBodyTags) GetTag() []*DescribeMultiZoneClusterResponseBodyTagsTag {
	return s.Tag
}

func (s *DescribeMultiZoneClusterResponseBodyTags) SetTag(v []*DescribeMultiZoneClusterResponseBodyTagsTag) *DescribeMultiZoneClusterResponseBodyTags {
	s.Tag = v
	return s
}

func (s *DescribeMultiZoneClusterResponseBodyTags) Validate() error {
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

type DescribeMultiZoneClusterResponseBodyTagsTag struct {
	// example:
	//
	// test_key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// test_value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeMultiZoneClusterResponseBodyTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeMultiZoneClusterResponseBodyTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneClusterResponseBodyTagsTag) GetKey() *string {
	return s.Key
}

func (s *DescribeMultiZoneClusterResponseBodyTagsTag) GetValue() *string {
	return s.Value
}

func (s *DescribeMultiZoneClusterResponseBodyTagsTag) SetKey(v string) *DescribeMultiZoneClusterResponseBodyTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBodyTagsTag) SetValue(v string) *DescribeMultiZoneClusterResponseBodyTagsTag {
	s.Value = &v
	return s
}

func (s *DescribeMultiZoneClusterResponseBodyTagsTag) Validate() error {
	return dara.Validate(s)
}
