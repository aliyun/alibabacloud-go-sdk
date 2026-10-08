// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCDisksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDisks(v []*DescribeRCDisksResponseBodyDisks) *DescribeRCDisksResponseBody
	GetDisks() []*DescribeRCDisksResponseBodyDisks
	SetPageNumber(v int64) *DescribeRCDisksResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeRCDisksResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeRCDisksResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeRCDisksResponseBody
	GetTotalCount() *int64
}

type DescribeRCDisksResponseBody struct {
	// The list of disk information.
	Disks []*DescribeRCDisksResponseBodyDisks `json:"Disks,omitempty" xml:"Disks,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8B993DA9-5272-5414-94E3-4CA8BA0146C2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 12
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeRCDisksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCDisksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRCDisksResponseBody) GetDisks() []*DescribeRCDisksResponseBodyDisks {
	return s.Disks
}

func (s *DescribeRCDisksResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeRCDisksResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeRCDisksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRCDisksResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeRCDisksResponseBody) SetDisks(v []*DescribeRCDisksResponseBodyDisks) *DescribeRCDisksResponseBody {
	s.Disks = v
	return s
}

func (s *DescribeRCDisksResponseBody) SetPageNumber(v int64) *DescribeRCDisksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeRCDisksResponseBody) SetPageSize(v int64) *DescribeRCDisksResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeRCDisksResponseBody) SetRequestId(v string) *DescribeRCDisksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRCDisksResponseBody) SetTotalCount(v int64) *DescribeRCDisksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeRCDisksResponseBody) Validate() error {
	if s.Disks != nil {
		for _, item := range s.Disks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRCDisksResponseBodyDisks struct {
	// The time when the disk was attached.
	//
	// example:
	//
	// 2017-12-05T2340:00Z
	AttachedTime *string `json:"AttachedTime,omitempty" xml:"AttachedTime,omitempty"`
	// Indicates whether burst (performance bursting) is enabled. Valid values:
	//
	// true: Enabled.
	//
	// false: Disabled.
	//
	// This parameter is supported only when DiskCategory is set to cloud_auto. For more information, see ESSD AutoPL cloud disks.
	BurstingEnabled *bool `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	// The disk category. Valid values:
	//
	// - **cloud_efficiency**: ultra cloud disk.
	//
	// - **cloud_ssd**: standard SSD.
	//
	// - **cloud_essd**: ESSD cloud disk.
	//
	// - **cloud_auto**: premium performance disk.
	//
	// example:
	//
	// cloud_auto
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2024-10-22T02:41:37Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// Indicates whether automatic snapshots are deleted when the cloud disk is deleted. Valid values:
	//
	// - true: Automatic snapshots are deleted when the cloud disk is deleted.
	//
	// - false: Automatic snapshots are retained when the cloud disk is deleted.
	//
	// example:
	//
	// true
	DeleteAutoSnapshot *bool `json:"DeleteAutoSnapshot,omitempty" xml:"DeleteAutoSnapshot,omitempty"`
	// Indicates whether the disk is released when the instance is released. Valid values:
	//
	// - true: The disk is released when the instance is released.
	//
	// - false: The disk is retained when the instance is released.
	//
	// example:
	//
	// true
	DeleteWithInstance *bool `json:"DeleteWithInstance,omitempty" xml:"DeleteWithInstance,omitempty"`
	// The disk description.
	//
	// example:
	//
	// zd_test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The mount point of the disk.
	//
	// example:
	//
	// /dev/xvda
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
	// Billable methods of the disk.
	//
	// Only **PostPaid*	- is supported, which indicates the pay-as-you-go billing method.
	//
	// example:
	//
	// PostPaid
	DiskChargeType *string `json:"DiskChargeType,omitempty" xml:"DiskChargeType,omitempty"`
	// The disk ID.
	//
	// example:
	//
	// rcd-wz9f3peueu5npsl****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The disk name.
	//
	// example:
	//
	// fvt-ecs-bcfb3627
	DiskName *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	// Indicates whether only encrypted cloud disks are filtered. Valid values:
	//
	// - true: Only encrypted cloud disks are returned.
	//
	// - false (default): All cloud disks are returned.
	//
	// example:
	//
	// true
	Encrypted *bool `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	// A reserved parameter. You do not need to specify this parameter.
	//
	// example:
	//
	// none
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The provisioned read/write IOPS of the ESSD AutoPL cloud disk. Valid values: 0 to min{50000, 1000 × Capacity - Baseline performance}. Baseline performance = min{1,800 + 50 × Capacity, 50,000}.
	//
	// This parameter is supported only when `Category` is set to `cloud_auto`.
	//
	// example:
	//
	// 4000
	IOPS *int64 `json:"IOPS,omitempty" xml:"IOPS,omitempty"`
	// The image ID used to create the RDS Custom instance. This parameter has a value only for cloud disks created from an image. Otherwise, the value is empty. This value remains unchanged throughout the lifecycle of the cloud disk.
	//
	// example:
	//
	// m-2zeb24dw6wripjn2****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rc-e8w1cn7634kiam****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The performance level (PL) of the ESSD cloud disk. Valid values:
	//
	// - PL0: A single standard SSD can deliver up to 10,000 random read/write IOPS.
	//
	// - PL1: A single standard SSD can deliver up to 50,000 random read/write IOPS.
	//
	// - PL2: A single standard SSD can deliver up to 100,000 random read/write IOPS.
	//
	// - PL3: A single standard SSD can deliver up to 1,000,000 random read/write IOPS.
	//
	// example:
	//
	// PL0
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// Indicates whether the disk is detachable.
	Portable *bool `json:"Portable,omitempty" xml:"Portable,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the disk belongs.
	//
	// example:
	//
	// rg-aekzescnje5khnq
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The serial number of the disk.
	//
	// example:
	//
	// bp18um4r4f2fve2****
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// The disk size. Unit: GiB.
	//
	// example:
	//
	// 60
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The snapshot ID used to create the cloud disk.
	//
	// If no snapshot was specified when the cloud disk was created, this parameter is empty. This value remains unchanged throughout the lifecycle of the cloud disk.
	//
	// example:
	//
	// rcds-bp67acfmxazb4p****
	SourceSnapshotId *string `json:"SourceSnapshotId,omitempty" xml:"SourceSnapshotId,omitempty"`
	// The disk status. Valid values:
	//
	// - In_use: in use.
	//
	// - Available: to be attached.
	//
	// - Attaching: being attached.
	//
	// - Detaching: being detached.
	//
	// - Creating: being created.
	//
	// - ReIniting: being initialized.
	//
	// example:
	//
	// In_use
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the dedicated block storage cluster to which the cloud disk belongs. If the cloud disk is in a public cloud block storage cluster, this parameter is empty.
	//
	// example:
	//
	// dbsc-cn-zvp2rl601****
	StorageClusterId *string `json:"StorageClusterId,omitempty" xml:"StorageClusterId,omitempty"`
	// The storage set ID.
	//
	// example:
	//
	// ss-i-bp1j4i2jdf3owlhe****
	StorageSetId *string `json:"StorageSetId,omitempty" xml:"StorageSetId,omitempty"`
	// The tags.
	Tag []*DescribeRCDisksResponseBodyDisksTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The disk type. Valid values:
	//
	// - system: system cloud disk.
	//
	// - data: data cloud disk.
	//
	// example:
	//
	// data
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-j
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeRCDisksResponseBodyDisks) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCDisksResponseBodyDisks) GoString() string {
	return s.String()
}

func (s *DescribeRCDisksResponseBodyDisks) GetAttachedTime() *string {
	return s.AttachedTime
}

func (s *DescribeRCDisksResponseBodyDisks) GetBurstingEnabled() *bool {
	return s.BurstingEnabled
}

func (s *DescribeRCDisksResponseBodyDisks) GetCategory() *string {
	return s.Category
}

func (s *DescribeRCDisksResponseBodyDisks) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeRCDisksResponseBodyDisks) GetDeleteAutoSnapshot() *bool {
	return s.DeleteAutoSnapshot
}

func (s *DescribeRCDisksResponseBodyDisks) GetDeleteWithInstance() *bool {
	return s.DeleteWithInstance
}

func (s *DescribeRCDisksResponseBodyDisks) GetDescription() *string {
	return s.Description
}

func (s *DescribeRCDisksResponseBodyDisks) GetDevice() *string {
	return s.Device
}

func (s *DescribeRCDisksResponseBodyDisks) GetDiskChargeType() *string {
	return s.DiskChargeType
}

func (s *DescribeRCDisksResponseBodyDisks) GetDiskId() *string {
	return s.DiskId
}

func (s *DescribeRCDisksResponseBodyDisks) GetDiskName() *string {
	return s.DiskName
}

func (s *DescribeRCDisksResponseBodyDisks) GetEncrypted() *bool {
	return s.Encrypted
}

func (s *DescribeRCDisksResponseBodyDisks) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *DescribeRCDisksResponseBodyDisks) GetIOPS() *int64 {
	return s.IOPS
}

func (s *DescribeRCDisksResponseBodyDisks) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeRCDisksResponseBodyDisks) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeRCDisksResponseBodyDisks) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *DescribeRCDisksResponseBodyDisks) GetPortable() *bool {
	return s.Portable
}

func (s *DescribeRCDisksResponseBodyDisks) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRCDisksResponseBodyDisks) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeRCDisksResponseBodyDisks) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *DescribeRCDisksResponseBodyDisks) GetSize() *int64 {
	return s.Size
}

func (s *DescribeRCDisksResponseBodyDisks) GetSourceSnapshotId() *string {
	return s.SourceSnapshotId
}

func (s *DescribeRCDisksResponseBodyDisks) GetStatus() *string {
	return s.Status
}

func (s *DescribeRCDisksResponseBodyDisks) GetStorageClusterId() *string {
	return s.StorageClusterId
}

func (s *DescribeRCDisksResponseBodyDisks) GetStorageSetId() *string {
	return s.StorageSetId
}

func (s *DescribeRCDisksResponseBodyDisks) GetTag() []*DescribeRCDisksResponseBodyDisksTag {
	return s.Tag
}

func (s *DescribeRCDisksResponseBodyDisks) GetType() *string {
	return s.Type
}

func (s *DescribeRCDisksResponseBodyDisks) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeRCDisksResponseBodyDisks) SetAttachedTime(v string) *DescribeRCDisksResponseBodyDisks {
	s.AttachedTime = &v
	return s
}

func (s *DescribeRCDisksResponseBodyDisks) SetBurstingEnabled(v bool) *DescribeRCDisksResponseBodyDisks {
	s.BurstingEnabled = &v
	return s
}

func (s *DescribeRCDisksResponseBodyDisks) SetCategory(v string) *DescribeRCDisksResponseBodyDisks {
	s.Category = &v
	return s
}

func (s *DescribeRCDisksResponseBodyDisks) SetCreationTime(v string) *DescribeRCDisksResponseBodyDisks {
	s.CreationTime = &v
	return s
}

func (s *DescribeRCDisksResponseBodyDisks) SetDeleteAutoSnapshot(v bool) *DescribeRCDisksResponseBodyDisks {
	s.DeleteAutoSnapshot = &v
	return s
}

func (s *DescribeRCDisksResponseBodyDisks) SetDeleteWithInstance(v bool) *DescribeRCDisksResponseBodyDisks {
	s.DeleteWithInstance = &v
	return s
}

func (s *DescribeRCDisksResponseBodyDisks) SetDescription(v string) *DescribeRCDisksResponseBodyDisks {
	s.Description = &v
	return s
}

func (s *DescribeRCDisksResponseBodyDisks) SetDevice(v string) *DescribeRCDisksResponseBodyDisks {
	s.Device = &v
	return s
}

func (s *DescribeRCDisksResponseBodyDisks) SetDiskChargeType(v string) *DescribeRCDisksResponseBodyDisks {
	s.DiskChargeType = &v
	return s
}

func (s *DescribeRCDisksResponseBodyDisks) SetDiskId(v string) *DescribeRCDisksResponseBodyDisks {
	s.DiskId = &v
	return s
}

func (s *DescribeRCDisksResponseBodyDisks) SetDiskName(v string) *DescribeRCDisksResponseBodyDisks {
	s.DiskName = &v
	return s
}

func (s *DescribeRCDisksResponseBodyDisks) SetEncrypted(v bool) *DescribeRCDisksResponseBodyDisks {
	s.Encrypted = &v
	return s
}

func (s *DescribeRCDisksResponseBodyDisks) SetExpiredTime(v string) *DescribeRCDisksResponseBodyDisks {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeRCDisksResponseBodyDisks) SetIOPS(v int64) *DescribeRCDisksResponseBodyDisks {
	s.IOPS = &v
	return s
}

func (s *DescribeRCDisksResponseBodyDisks) SetImageId(v string) *DescribeRCDisksResponseBodyDisks {
	s.ImageId = &v
	return s
}

func (s *DescribeRCDisksResponseBodyDisks) SetInstanceId(v string) *DescribeRCDisksResponseBodyDisks {
	s.InstanceId = &v
	return s
}

func (s *DescribeRCDisksResponseBodyDisks) SetPerformanceLevel(v string) *DescribeRCDisksResponseBodyDisks {
	s.PerformanceLevel = &v
	return s
}

func (s *DescribeRCDisksResponseBodyDisks) SetPortable(v bool) *DescribeRCDisksResponseBodyDisks {
	s.Portable = &v
	return s
}

func (s *DescribeRCDisksResponseBodyDisks) SetRegionId(v string) *DescribeRCDisksResponseBodyDisks {
	s.RegionId = &v
	return s
}

func (s *DescribeRCDisksResponseBodyDisks) SetResourceGroupId(v string) *DescribeRCDisksResponseBodyDisks {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeRCDisksResponseBodyDisks) SetSerialNumber(v string) *DescribeRCDisksResponseBodyDisks {
	s.SerialNumber = &v
	return s
}

func (s *DescribeRCDisksResponseBodyDisks) SetSize(v int64) *DescribeRCDisksResponseBodyDisks {
	s.Size = &v
	return s
}

func (s *DescribeRCDisksResponseBodyDisks) SetSourceSnapshotId(v string) *DescribeRCDisksResponseBodyDisks {
	s.SourceSnapshotId = &v
	return s
}

func (s *DescribeRCDisksResponseBodyDisks) SetStatus(v string) *DescribeRCDisksResponseBodyDisks {
	s.Status = &v
	return s
}

func (s *DescribeRCDisksResponseBodyDisks) SetStorageClusterId(v string) *DescribeRCDisksResponseBodyDisks {
	s.StorageClusterId = &v
	return s
}

func (s *DescribeRCDisksResponseBodyDisks) SetStorageSetId(v string) *DescribeRCDisksResponseBodyDisks {
	s.StorageSetId = &v
	return s
}

func (s *DescribeRCDisksResponseBodyDisks) SetTag(v []*DescribeRCDisksResponseBodyDisksTag) *DescribeRCDisksResponseBodyDisks {
	s.Tag = v
	return s
}

func (s *DescribeRCDisksResponseBodyDisks) SetType(v string) *DescribeRCDisksResponseBodyDisks {
	s.Type = &v
	return s
}

func (s *DescribeRCDisksResponseBodyDisks) SetZoneId(v string) *DescribeRCDisksResponseBodyDisks {
	s.ZoneId = &v
	return s
}

func (s *DescribeRCDisksResponseBodyDisks) Validate() error {
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

type DescribeRCDisksResponseBodyDisksTag struct {
	// The tag key.
	//
	// example:
	//
	// testkey1
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// testvalue1
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeRCDisksResponseBodyDisksTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCDisksResponseBodyDisksTag) GoString() string {
	return s.String()
}

func (s *DescribeRCDisksResponseBodyDisksTag) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeRCDisksResponseBodyDisksTag) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeRCDisksResponseBodyDisksTag) SetTagKey(v string) *DescribeRCDisksResponseBodyDisksTag {
	s.TagKey = &v
	return s
}

func (s *DescribeRCDisksResponseBodyDisksTag) SetTagValue(v string) *DescribeRCDisksResponseBodyDisksTag {
	s.TagValue = &v
	return s
}

func (s *DescribeRCDisksResponseBodyDisksTag) Validate() error {
	return dara.Validate(s)
}
