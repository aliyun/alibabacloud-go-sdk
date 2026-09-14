// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDedicatedBlockStorageClusterDisksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDisks(v *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisks) *DescribeDedicatedBlockStorageClusterDisksResponseBody
	GetDisks() *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisks
	SetNextToken(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBody
	GetRequestId() *string
}

type DescribeDedicatedBlockStorageClusterDisksResponseBody struct {
	// The collection of cloud disk information.
	Disks *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisks `json:"Disks,omitempty" xml:"Disks,omitempty" type:"Struct"`
	// The pagination token returned in this call.
	//
	// example:
	//
	// AAAAAdDWBF2
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 11B55F58-D3A4-4A9B-9596-342420D0****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDedicatedBlockStorageClusterDisksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedBlockStorageClusterDisksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBody) GetDisks() *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisks {
	return s.Disks
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBody) SetDisks(v *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisks) *DescribeDedicatedBlockStorageClusterDisksResponseBody {
	s.Disks = v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBody) SetNextToken(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBody) SetRequestId(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBody) Validate() error {
	if s.Disks != nil {
		if err := s.Disks.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDedicatedBlockStorageClusterDisksResponseBodyDisks struct {
	// The collection of cloud disk information.
	Disk []*DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk `json:"Disk,omitempty" xml:"Disk,omitempty" type:"Repeated"`
}

func (s DescribeDedicatedBlockStorageClusterDisksResponseBodyDisks) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedBlockStorageClusterDisksResponseBodyDisks) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisks) GetDisk() []*DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	return s.Disk
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisks) SetDisk(v []*DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisks {
	s.Disk = v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisks) Validate() error {
	if s.Disk != nil {
		for _, item := range s.Disk {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk struct {
	// The time when the cloud disk was last attached. The time follows the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddThh:mmZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-06-07T06:08:56Z
	AttachedTime *string `json:"AttachedTime,omitempty" xml:"AttachedTime,omitempty"`
	// This parameter is in invitational preview and is not publicly available.
	//
	// example:
	//
	// null
	BdfId *string `json:"BdfId,omitempty" xml:"BdfId,omitempty"`
	// Indicates whether the burst (performance burst) feature is enabled. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// This parameter is supported only when `DiskCategory` is set to `cloud_auto`. For more information, see [ESSD AutoPL cloud disks](https://help.aliyun.com/document_detail/368372.html).
	//
	// example:
	//
	// false
	BurstingEnabled *bool `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	// The category of the cloud disk or local disk is cloud_essd, which indicates an ESSD.
	//
	// example:
	//
	// cloud_essd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// Indicates whether automatic snapshots are deleted when the cloud disk is released. Valid values:
	//
	// - true: Automatic snapshots are deleted when the cloud disk is released.
	//
	// - false: Automatic snapshots are retained when the cloud disk is released.
	//
	// Snapshots created by calling [CreateSnapshot](https://help.aliyun.com/document_detail/25524.html) or by using the console are not affected by this parameter and are always retained.
	//
	// example:
	//
	// false
	DeleteAutoSnapshot *bool `json:"DeleteAutoSnapshot,omitempty" xml:"DeleteAutoSnapshot,omitempty"`
	// Indicates whether the cloud disk is released when the instance is released. Valid values:
	//
	// - true: The cloud disk is released when the instance is released.
	//
	// - false: The cloud disk is retained when the instance is released.
	//
	// example:
	//
	// true
	DeleteWithInstance *bool `json:"DeleteWithInstance,omitempty" xml:"DeleteWithInstance,omitempty"`
	// The cloud disk description.
	//
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the cloud disk was last detached.
	//
	// example:
	//
	// 2021-06-07T21:01:22Z
	DetachedTime *string `json:"DetachedTime,omitempty" xml:"DetachedTime,omitempty"`
	// The device name of the instance to which the cloud disk is attached, such as /dev/xvdb. Note the following items:
	//
	// - This parameter has a value only when the `Status` parameter is set to `In_use`. This parameter is empty in other states.
	//
	// - For cloud disks with the multi-attach feature enabled, this value is always empty. You can view all attachment information of the cloud disk from the returned `Attachment` list.
	//
	// > This parameter will be deprecated. To ensure code compatibility, do not use this parameter.
	//
	// example:
	//
	// /dev/xvdb
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
	// The billing method of the cloud disk. Valid values:
	//
	// - PrePaid: subscription.
	//
	// - PostPaid: pay-as-you-go.
	//
	// example:
	//
	// PrePaid
	DiskChargeType *string `json:"DiskChargeType,omitempty" xml:"DiskChargeType,omitempty"`
	// The cloud disk ID.
	//
	// example:
	//
	// d-bp67acfmxazb4p****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The cloud disk name.
	//
	// example:
	//
	// testDiskName
	DiskName *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	// Indicates whether the automatic snapshot policy feature is enabled for the cloud disk.
	//
	// >This parameter is deprecated. After a cloud disk is created, the automatic snapshot policy feature is enabled by default. You only need to associate an automatic snapshot policy with the cloud disk.
	//
	// example:
	//
	// false
	EnableAutoSnapshot *bool `json:"EnableAutoSnapshot,omitempty" xml:"EnableAutoSnapshot,omitempty"`
	// Indicates whether the cloud disk is encrypted.
	//
	// example:
	//
	// false
	Encrypted *bool `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	// The maximum number of read/write (I/O) operations per second. Unit: operations/s.
	//
	// example:
	//
	// 4000
	IOPS *int64 `json:"IOPS,omitempty" xml:"IOPS,omitempty"`
	// The ID of the image used to create the ECS instance. This parameter has a value only for cloud disks created from an image. Otherwise, this value is empty. This value remains unchanged throughout the lifecycle of the cloud disk.
	//
	// example:
	//
	// m-bp13aqm171qynt3u***
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The instance ID of the instance to which the cloud disk is mounted. Note the following items:
	//
	// - This parameter has a value only when the `Status` parameter is set to `In_use`. This parameter is empty in other states.
	//
	// - For cloud disks with the multi-attach attribute enabled, this value is always empty. You can view all mount information of the cloud disk from the returned `Attachment` list.
	//
	// example:
	//
	// i-bp67acfmxazb4q****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The KMS key ID used by the cloud disk.
	//
	// example:
	//
	// 0e478b7a-4262-4802-b8cb-00d3fb40****
	KMSKeyId *string `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	// The number of instances to which the shared storage is attached.
	//
	// example:
	//
	// 1
	MountInstanceNum *int32 `json:"MountInstanceNum,omitempty" xml:"MountInstanceNum,omitempty"`
	// Indicates whether the multi-attach feature is enabled for the cloud disk.
	//
	// example:
	//
	// Disabled
	MultiAttach *string `json:"MultiAttach,omitempty" xml:"MultiAttach,omitempty"`
	// The performance level of the ESSD. Valid values:
	//
	// - PL0: a maximum of 10,000 random read/write IOPS per cloud disk.
	//
	// - PL1: a maximum of 50,000 random read/write IOPS per cloud disk.
	//
	// - PL2: a maximum of 100,000 random read/write IOPS per cloud disk.
	//
	// - PL3: a maximum of 1,000,000 random read/write IOPS per cloud disk.
	//
	// example:
	//
	// PL0
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// Indicates whether the cloud disk is removable.
	//
	// example:
	//
	// false
	Portable *bool `json:"Portable,omitempty" xml:"Portable,omitempty"`
	// The provisioned read/write IOPS of the ESSD AutoPL cloud disk. Valid values: 0 to min{50000, 1000 × Capacity - Baseline performance}.
	//
	// Baseline performance = min{1,800 + 50 × Capacity, 50,000}.
	//
	// This parameter is supported only when `DiskCategory` is set to `cloud_auto`. For more information, see [ESSD AutoPL cloud disks](https://help.aliyun.com/document_detail/368372.html).
	//
	// example:
	//
	// 40000
	ProvisionedIops *int64 `json:"ProvisionedIops,omitempty" xml:"ProvisionedIops,omitempty"`
	// The region ID of the cloud disk.
	//
	// example:
	//
	// cn-heyuan
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The cloud disk size. Unit: GiB.
	//
	// example:
	//
	// 60
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The ID of the snapshot used to create the cloud disk.
	//
	// If no snapshot was specified when the cloud disk was created, this value is empty. This value remains unchanged throughout the lifecycle of the cloud disk.
	//
	// example:
	//
	// s-bp67acfmxazb4p****
	SourceSnapshotId *string `json:"SourceSnapshotId,omitempty" xml:"SourceSnapshotId,omitempty"`
	// The cloud disk status. For more information, see [Cloud disk status](https://help.aliyun.com/document_detail/25689.html). Valid values:
	//
	// -   In_use.
	//
	// -   Available.
	//
	// -   Attaching.
	//
	// -   Detaching.
	//
	// -   Creating.
	//
	// -   ReIniting.
	//
	// example:
	//
	// In_use
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the dedicated block storage cluster to which the cloud disk belongs. If the cloud disk belongs to a public cloud block storage cluster, this value is empty.
	//
	// example:
	//
	// dbsc-j5e1sf2vaf5he8m2****
	StorageClusterId *string `json:"StorageClusterId,omitempty" xml:"StorageClusterId,omitempty"`
	// The storage set ID.
	//
	// example:
	//
	// ss-i-bp1j4i2jdf3owlhe****
	StorageSetId *string `json:"StorageSetId,omitempty" xml:"StorageSetId,omitempty"`
	// The maximum number of partitions in the storage set.
	//
	// example:
	//
	// 11
	StorageSetPartitionNumber *int32 `json:"StorageSetPartitionNumber,omitempty" xml:"StorageSetPartitionNumber,omitempty"`
	// The tags of the cloud disk.
	Tags []*DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDiskTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The amount of data that can be transferred per unit of time. Unit: MB/s.
	//
	// example:
	//
	// 100
	Throughput *int64 `json:"Throughput,omitempty" xml:"Throughput,omitempty"`
	// The type of the cloud disk. Valid values:
	//
	// - system: system cloud disk.
	//
	// - data: data cloud disk.
	//
	// example:
	//
	// system
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The zone ID of the cloud disk.
	//
	// example:
	//
	// cn-heyuan-i
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) GetAttachedTime() *string {
	return s.AttachedTime
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) GetBdfId() *string {
	return s.BdfId
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) GetBurstingEnabled() *bool {
	return s.BurstingEnabled
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) GetCategory() *string {
	return s.Category
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) GetDeleteAutoSnapshot() *bool {
	return s.DeleteAutoSnapshot
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) GetDeleteWithInstance() *bool {
	return s.DeleteWithInstance
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) GetDescription() *string {
	return s.Description
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) GetDetachedTime() *string {
	return s.DetachedTime
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) GetDevice() *string {
	return s.Device
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) GetDiskChargeType() *string {
	return s.DiskChargeType
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) GetDiskId() *string {
	return s.DiskId
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) GetDiskName() *string {
	return s.DiskName
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) GetEnableAutoSnapshot() *bool {
	return s.EnableAutoSnapshot
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) GetEncrypted() *bool {
	return s.Encrypted
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) GetIOPS() *int64 {
	return s.IOPS
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) GetKMSKeyId() *string {
	return s.KMSKeyId
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) GetMountInstanceNum() *int32 {
	return s.MountInstanceNum
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) GetMultiAttach() *string {
	return s.MultiAttach
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) GetPortable() *bool {
	return s.Portable
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) GetProvisionedIops() *int64 {
	return s.ProvisionedIops
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) GetSize() *int32 {
	return s.Size
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) GetSourceSnapshotId() *string {
	return s.SourceSnapshotId
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) GetStatus() *string {
	return s.Status
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) GetStorageClusterId() *string {
	return s.StorageClusterId
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) GetStorageSetId() *string {
	return s.StorageSetId
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) GetStorageSetPartitionNumber() *int32 {
	return s.StorageSetPartitionNumber
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) GetTags() []*DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDiskTags {
	return s.Tags
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) GetThroughput() *int64 {
	return s.Throughput
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) GetType() *string {
	return s.Type
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetAttachedTime(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.AttachedTime = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetBdfId(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.BdfId = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetBurstingEnabled(v bool) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.BurstingEnabled = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetCategory(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.Category = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetDeleteAutoSnapshot(v bool) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.DeleteAutoSnapshot = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetDeleteWithInstance(v bool) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.DeleteWithInstance = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetDescription(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.Description = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetDetachedTime(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.DetachedTime = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetDevice(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.Device = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetDiskChargeType(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.DiskChargeType = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetDiskId(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.DiskId = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetDiskName(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.DiskName = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetEnableAutoSnapshot(v bool) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.EnableAutoSnapshot = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetEncrypted(v bool) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.Encrypted = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetIOPS(v int64) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.IOPS = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetImageId(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.ImageId = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetInstanceId(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.InstanceId = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetKMSKeyId(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.KMSKeyId = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetMountInstanceNum(v int32) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.MountInstanceNum = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetMultiAttach(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.MultiAttach = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetPerformanceLevel(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetPortable(v bool) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.Portable = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetProvisionedIops(v int64) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.ProvisionedIops = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetRegionId(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.RegionId = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetSize(v int32) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.Size = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetSourceSnapshotId(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.SourceSnapshotId = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetStatus(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.Status = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetStorageClusterId(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.StorageClusterId = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetStorageSetId(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.StorageSetId = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetStorageSetPartitionNumber(v int32) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.StorageSetPartitionNumber = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetTags(v []*DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDiskTags) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.Tags = v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetThroughput(v int64) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.Throughput = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetType(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.Type = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) SetZoneId(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk {
	s.ZoneId = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDisk) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDiskTags struct {
	// The tag key.
	//
	// example:
	//
	// TestKey
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// TestValue
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDiskTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDiskTags) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDiskTags) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDiskTags) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDiskTags) SetTagKey(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDiskTags {
	s.TagKey = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDiskTags) SetTagValue(v string) *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDiskTags {
	s.TagValue = &v
	return s
}

func (s *DescribeDedicatedBlockStorageClusterDisksResponseBodyDisksDiskTags) Validate() error {
	return dara.Validate(s)
}
