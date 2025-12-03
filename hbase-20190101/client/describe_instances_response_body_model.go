// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstances(v *DescribeInstancesResponseBodyInstances) *DescribeInstancesResponseBody
	GetInstances() *DescribeInstancesResponseBodyInstances
	SetPageNumber(v int32) *DescribeInstancesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeInstancesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeInstancesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeInstancesResponseBody
	GetTotalCount() *int64
}

type DescribeInstancesResponseBody struct {
	Instances *DescribeInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Struct"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// EBECBF12-2E34-41BE-8DE9-FC3700D4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 18
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBody) GetInstances() *DescribeInstancesResponseBodyInstances {
	return s.Instances
}

func (s *DescribeInstancesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeInstancesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstancesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeInstancesResponseBody) SetInstances(v *DescribeInstancesResponseBodyInstances) *DescribeInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *DescribeInstancesResponseBody) SetPageNumber(v int32) *DescribeInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetPageSize(v int32) *DescribeInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetRequestId(v string) *DescribeInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetTotalCount(v int64) *DescribeInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeInstancesResponseBody) Validate() error {
	if s.Instances != nil {
		if err := s.Instances.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInstancesResponseBodyInstances struct {
	Instance []*DescribeInstancesResponseBodyInstancesInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Repeated"`
}

func (s DescribeInstancesResponseBodyInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstances) GetInstance() []*DescribeInstancesResponseBodyInstancesInstance {
	return s.Instance
}

func (s *DescribeInstancesResponseBodyInstances) SetInstance(v []*DescribeInstancesResponseBodyInstancesInstance) *DescribeInstancesResponseBodyInstances {
	s.Instance = v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) Validate() error {
	if s.Instance != nil {
		for _, item := range s.Instance {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstancesResponseBodyInstancesInstance struct {
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
	// hb-bp1u0639js2h7****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// test
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// example:
	//
	// cluster
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// example:
	//
	// open
	ColdStorageStatus *string `json:"ColdStorageStatus,omitempty" xml:"ColdStorageStatus,omitempty"`
	// example:
	//
	// 2
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
	// 2
	CoreNodeCount *int32 `json:"CoreNodeCount,omitempty" xml:"CoreNodeCount,omitempty"`
	// example:
	//
	// 2019-09-12T14:40:46
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// example:
	//
	// 2019-09-12T14:40:46Z
	CreatedTimeUTC *string `json:"CreatedTimeUTC,omitempty" xml:"CreatedTimeUTC,omitempty"`
	// example:
	//
	// 12
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// hbase
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// example:
	//
	// 2019-10-12T14:40:46
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// 2019-10-12T14:40:46Z
	ExpireTimeUTC *string `json:"ExpireTimeUTC,omitempty" xml:"ExpireTimeUTC,omitempty"`
	// example:
	//
	// hb-bp1u0639js2h7****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// test
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// true
	IsDeletionProtection *bool `json:"IsDeletionProtection,omitempty" xml:"IsDeletionProtection,omitempty"`
	// example:
	//
	// true
	IsHa *bool `json:"IsHa,omitempty" xml:"IsHa,omitempty"`
	// example:
	//
	// 2.0
	MajorVersion *string `json:"MajorVersion,omitempty" xml:"MajorVersion,omitempty"`
	// example:
	//
	// 100
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
	// 1.0
	ModuleStackVersion *string `json:"ModuleStackVersion,omitempty" xml:"ModuleStackVersion,omitempty"`
	// example:
	//
	// VPC
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// example:
	//
	// 2980****2123
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
	// rg-4f51d54g5****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// ACTIVATION
	Status *string                                             `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags   *DescribeInstancesResponseBodyInstancesInstanceTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// example:
	//
	// vpc-bp120k6ixs4eoghz*****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// example:
	//
	// vsw-bp191ipotq****dbqf
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	// example:
	//
	// cn-hangzhou-f
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesInstance) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstance) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetAutoRenewal() *bool {
	return s.AutoRenewal
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetBackupStatus() *string {
	return s.BackupStatus
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetClusterName() *string {
	return s.ClusterName
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetClusterType() *string {
	return s.ClusterType
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetColdStorageStatus() *string {
	return s.ColdStorageStatus
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetCoreDiskCount() *string {
	return s.CoreDiskCount
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetCoreDiskSize() *int32 {
	return s.CoreDiskSize
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetCoreDiskType() *string {
	return s.CoreDiskType
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetCoreInstanceType() *string {
	return s.CoreInstanceType
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetCoreNodeCount() *int32 {
	return s.CoreNodeCount
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetCreatedTimeUTC() *string {
	return s.CreatedTimeUTC
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetDuration() *int32 {
	return s.Duration
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetEngine() *string {
	return s.Engine
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetExpireTimeUTC() *string {
	return s.ExpireTimeUTC
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetIsDeletionProtection() *bool {
	return s.IsDeletionProtection
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetIsHa() *bool {
	return s.IsHa
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetMajorVersion() *string {
	return s.MajorVersion
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetMasterDiskSize() *int32 {
	return s.MasterDiskSize
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetMasterDiskType() *string {
	return s.MasterDiskType
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetMasterInstanceType() *string {
	return s.MasterInstanceType
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetMasterNodeCount() *int32 {
	return s.MasterNodeCount
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetModuleId() *int32 {
	return s.ModuleId
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetModuleStackVersion() *string {
	return s.ModuleStackVersion
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetParentId() *string {
	return s.ParentId
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetPayType() *string {
	return s.PayType
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetStatus() *string {
	return s.Status
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetTags() *DescribeInstancesResponseBodyInstancesInstanceTags {
	return s.Tags
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetVswitchId() *string {
	return s.VswitchId
}

func (s *DescribeInstancesResponseBodyInstancesInstance) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetAutoRenewal(v bool) *DescribeInstancesResponseBodyInstancesInstance {
	s.AutoRenewal = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetBackupStatus(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.BackupStatus = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetClusterId(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.ClusterId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetClusterName(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.ClusterName = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetClusterType(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.ClusterType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetColdStorageStatus(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.ColdStorageStatus = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetCoreDiskCount(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.CoreDiskCount = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetCoreDiskSize(v int32) *DescribeInstancesResponseBodyInstancesInstance {
	s.CoreDiskSize = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetCoreDiskType(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.CoreDiskType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetCoreInstanceType(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.CoreInstanceType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetCoreNodeCount(v int32) *DescribeInstancesResponseBodyInstancesInstance {
	s.CoreNodeCount = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetCreatedTime(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.CreatedTime = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetCreatedTimeUTC(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.CreatedTimeUTC = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetDuration(v int32) *DescribeInstancesResponseBodyInstancesInstance {
	s.Duration = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetEngine(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.Engine = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetExpireTime(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.ExpireTime = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetExpireTimeUTC(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.ExpireTimeUTC = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetInstanceId(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetInstanceName(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.InstanceName = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetIsDeletionProtection(v bool) *DescribeInstancesResponseBodyInstancesInstance {
	s.IsDeletionProtection = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetIsHa(v bool) *DescribeInstancesResponseBodyInstancesInstance {
	s.IsHa = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetMajorVersion(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.MajorVersion = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetMasterDiskSize(v int32) *DescribeInstancesResponseBodyInstancesInstance {
	s.MasterDiskSize = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetMasterDiskType(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.MasterDiskType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetMasterInstanceType(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.MasterInstanceType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetMasterNodeCount(v int32) *DescribeInstancesResponseBodyInstancesInstance {
	s.MasterNodeCount = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetModuleId(v int32) *DescribeInstancesResponseBodyInstancesInstance {
	s.ModuleId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetModuleStackVersion(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.ModuleStackVersion = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetNetworkType(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.NetworkType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetParentId(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.ParentId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetPayType(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.PayType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetRegionId(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.RegionId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetResourceGroupId(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetStatus(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.Status = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetTags(v *DescribeInstancesResponseBodyInstancesInstanceTags) *DescribeInstancesResponseBodyInstancesInstance {
	s.Tags = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetVpcId(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.VpcId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetVswitchId(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.VswitchId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) SetZoneId(v string) *DescribeInstancesResponseBodyInstancesInstance {
	s.ZoneId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstance) Validate() error {
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInstancesResponseBodyInstancesInstanceTags struct {
	Tag []*DescribeInstancesResponseBodyInstancesInstanceTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeInstancesResponseBodyInstancesInstanceTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstanceTags) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstanceTags) GetTag() []*DescribeInstancesResponseBodyInstancesInstanceTagsTag {
	return s.Tag
}

func (s *DescribeInstancesResponseBodyInstancesInstanceTags) SetTag(v []*DescribeInstancesResponseBodyInstancesInstanceTagsTag) *DescribeInstancesResponseBodyInstancesInstanceTags {
	s.Tag = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceTags) Validate() error {
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

type DescribeInstancesResponseBodyInstancesInstanceTagsTag struct {
	// example:
	//
	// test-key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// test-value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesInstanceTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesInstanceTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesInstanceTagsTag) GetKey() *string {
	return s.Key
}

func (s *DescribeInstancesResponseBodyInstancesInstanceTagsTag) GetValue() *string {
	return s.Value
}

func (s *DescribeInstancesResponseBodyInstancesInstanceTagsTag) SetKey(v string) *DescribeInstancesResponseBodyInstancesInstanceTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceTagsTag) SetValue(v string) *DescribeInstancesResponseBodyInstancesInstanceTagsTag {
	s.Value = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesInstanceTagsTag) Validate() error {
	return dara.Validate(s)
}
