// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDisksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDisks(v []*ListDisksResponseBodyDisks) *ListDisksResponseBody
	GetDisks() []*ListDisksResponseBodyDisks
	SetPageNumber(v int32) *ListDisksResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDisksResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListDisksResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListDisksResponseBody
	GetTotalCount() *int32
}

type ListDisksResponseBody struct {
	// The queried disks.
	Disks []*ListDisksResponseBodyDisks `json:"Disks,omitempty" xml:"Disks,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F534F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDisksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDisksResponseBody) GoString() string {
	return s.String()
}

func (s *ListDisksResponseBody) GetDisks() []*ListDisksResponseBodyDisks {
	return s.Disks
}

func (s *ListDisksResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDisksResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDisksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDisksResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDisksResponseBody) SetDisks(v []*ListDisksResponseBodyDisks) *ListDisksResponseBody {
	s.Disks = v
	return s
}

func (s *ListDisksResponseBody) SetPageNumber(v int32) *ListDisksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListDisksResponseBody) SetPageSize(v int32) *ListDisksResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListDisksResponseBody) SetRequestId(v string) *ListDisksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDisksResponseBody) SetTotalCount(v int32) *ListDisksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDisksResponseBody) Validate() error {
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

type ListDisksResponseBodyDisks struct {
	// The category of the disk. Valid values:
	//
	// - ESSD: enhanced SSD (ESSD) of PL0
	//
	// - SSD: standard SSD
	//
	// - CLOUD_EFFICIENCY: ultra disk
	//
	// example:
	//
	// ESSD
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The time when the disk was created. The time follows the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-03-08T05:31:06Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The device name of the disk after the disk is attached to the simple application server.
	//
	// example:
	//
	// /dev/xvda
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
	// The billing method of the disk.
	//
	// example:
	//
	// PrePaid
	DiskChargeType *string `json:"DiskChargeType,omitempty" xml:"DiskChargeType,omitempty"`
	// The disk ID.
	//
	// example:
	//
	// d-bp14wq0149cpp2x****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The name of the disk.
	//
	// example:
	//
	// SystemDisk
	DiskName *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	// The disk type. Valid values:
	//
	// - system: system disk
	//
	// - data: data disk
	//
	// example:
	//
	// System
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// The ID of the simple application server to which the disk is attached.
	//
	// example:
	//
	// 2ad1ae67295445f598017499dc****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the simple application server.
	//
	// example:
	//
	// myInstance
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The remarks of the disk.
	//
	// example:
	//
	// remark
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The ID of the resource group to which the disk belongs.
	//
	// example:
	//
	// rg-aek2bti7cf7****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The size of the disk. Unit: GB.
	//
	// example:
	//
	// 50
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The status of the disk. Valid values:
	//
	// - ReIniting: The disk is being initialized.
	//
	// - Creating: The disk is being created.
	//
	// - In_use: The disk is in use.
	//
	// - Available: The disk can be attached.
	//
	// - Attaching: The disk is being attached.
	//
	// - Detaching: The disk is being detached.
	//
	// example:
	//
	// In_use
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags that are added to the disks.
	Tags []*ListDisksResponseBodyDisksTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListDisksResponseBodyDisks) String() string {
	return dara.Prettify(s)
}

func (s ListDisksResponseBodyDisks) GoString() string {
	return s.String()
}

func (s *ListDisksResponseBodyDisks) GetCategory() *string {
	return s.Category
}

func (s *ListDisksResponseBodyDisks) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListDisksResponseBodyDisks) GetDevice() *string {
	return s.Device
}

func (s *ListDisksResponseBodyDisks) GetDiskChargeType() *string {
	return s.DiskChargeType
}

func (s *ListDisksResponseBodyDisks) GetDiskId() *string {
	return s.DiskId
}

func (s *ListDisksResponseBodyDisks) GetDiskName() *string {
	return s.DiskName
}

func (s *ListDisksResponseBodyDisks) GetDiskType() *string {
	return s.DiskType
}

func (s *ListDisksResponseBodyDisks) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListDisksResponseBodyDisks) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListDisksResponseBodyDisks) GetRegionId() *string {
	return s.RegionId
}

func (s *ListDisksResponseBodyDisks) GetRemark() *string {
	return s.Remark
}

func (s *ListDisksResponseBodyDisks) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListDisksResponseBodyDisks) GetSize() *int32 {
	return s.Size
}

func (s *ListDisksResponseBodyDisks) GetStatus() *string {
	return s.Status
}

func (s *ListDisksResponseBodyDisks) GetTags() []*ListDisksResponseBodyDisksTags {
	return s.Tags
}

func (s *ListDisksResponseBodyDisks) SetCategory(v string) *ListDisksResponseBodyDisks {
	s.Category = &v
	return s
}

func (s *ListDisksResponseBodyDisks) SetCreationTime(v string) *ListDisksResponseBodyDisks {
	s.CreationTime = &v
	return s
}

func (s *ListDisksResponseBodyDisks) SetDevice(v string) *ListDisksResponseBodyDisks {
	s.Device = &v
	return s
}

func (s *ListDisksResponseBodyDisks) SetDiskChargeType(v string) *ListDisksResponseBodyDisks {
	s.DiskChargeType = &v
	return s
}

func (s *ListDisksResponseBodyDisks) SetDiskId(v string) *ListDisksResponseBodyDisks {
	s.DiskId = &v
	return s
}

func (s *ListDisksResponseBodyDisks) SetDiskName(v string) *ListDisksResponseBodyDisks {
	s.DiskName = &v
	return s
}

func (s *ListDisksResponseBodyDisks) SetDiskType(v string) *ListDisksResponseBodyDisks {
	s.DiskType = &v
	return s
}

func (s *ListDisksResponseBodyDisks) SetInstanceId(v string) *ListDisksResponseBodyDisks {
	s.InstanceId = &v
	return s
}

func (s *ListDisksResponseBodyDisks) SetInstanceName(v string) *ListDisksResponseBodyDisks {
	s.InstanceName = &v
	return s
}

func (s *ListDisksResponseBodyDisks) SetRegionId(v string) *ListDisksResponseBodyDisks {
	s.RegionId = &v
	return s
}

func (s *ListDisksResponseBodyDisks) SetRemark(v string) *ListDisksResponseBodyDisks {
	s.Remark = &v
	return s
}

func (s *ListDisksResponseBodyDisks) SetResourceGroupId(v string) *ListDisksResponseBodyDisks {
	s.ResourceGroupId = &v
	return s
}

func (s *ListDisksResponseBodyDisks) SetSize(v int32) *ListDisksResponseBodyDisks {
	s.Size = &v
	return s
}

func (s *ListDisksResponseBodyDisks) SetStatus(v string) *ListDisksResponseBodyDisks {
	s.Status = &v
	return s
}

func (s *ListDisksResponseBodyDisks) SetTags(v []*ListDisksResponseBodyDisksTags) *ListDisksResponseBodyDisks {
	s.Tags = v
	return s
}

func (s *ListDisksResponseBodyDisks) Validate() error {
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

type ListDisksResponseBodyDisksTags struct {
	// The tag key.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDisksResponseBodyDisksTags) String() string {
	return dara.Prettify(s)
}

func (s ListDisksResponseBodyDisksTags) GoString() string {
	return s.String()
}

func (s *ListDisksResponseBodyDisksTags) GetKey() *string {
	return s.Key
}

func (s *ListDisksResponseBodyDisksTags) GetValue() *string {
	return s.Value
}

func (s *ListDisksResponseBodyDisksTags) SetKey(v string) *ListDisksResponseBodyDisksTags {
	s.Key = &v
	return s
}

func (s *ListDisksResponseBodyDisksTags) SetValue(v string) *ListDisksResponseBodyDisksTags {
	s.Value = &v
	return s
}

func (s *ListDisksResponseBodyDisksTags) Validate() error {
	return dara.Validate(s)
}
