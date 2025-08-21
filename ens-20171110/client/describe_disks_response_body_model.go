// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDisksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeDisksResponseBody
	GetCode() *int32
	SetDisks(v *DescribeDisksResponseBodyDisks) *DescribeDisksResponseBody
	GetDisks() *DescribeDisksResponseBodyDisks
	SetPageNumber(v int32) *DescribeDisksResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDisksResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeDisksResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeDisksResponseBody
	GetTotalCount() *int32
}

type DescribeDisksResponseBody struct {
	// The returned service code. 0 indicates that the request was successful.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the disks.
	Disks *DescribeDisksResponseBodyDisks `json:"Disks,omitempty" xml:"Disks,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page. Maximum value: **50**. Default value: **10**.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BA984
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned pages.
	//
	// example:
	//
	// 13
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDisksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDisksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDisksResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeDisksResponseBody) GetDisks() *DescribeDisksResponseBodyDisks {
	return s.Disks
}

func (s *DescribeDisksResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDisksResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDisksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDisksResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeDisksResponseBody) SetCode(v int32) *DescribeDisksResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeDisksResponseBody) SetDisks(v *DescribeDisksResponseBodyDisks) *DescribeDisksResponseBody {
	s.Disks = v
	return s
}

func (s *DescribeDisksResponseBody) SetPageNumber(v int32) *DescribeDisksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDisksResponseBody) SetPageSize(v int32) *DescribeDisksResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDisksResponseBody) SetRequestId(v string) *DescribeDisksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDisksResponseBody) SetTotalCount(v int32) *DescribeDisksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDisksResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDisksResponseBodyDisks struct {
	Disks []*DescribeDisksResponseBodyDisksDisks `json:"Disks,omitempty" xml:"Disks,omitempty" type:"Repeated"`
}

func (s DescribeDisksResponseBodyDisks) String() string {
	return dara.Prettify(s)
}

func (s DescribeDisksResponseBodyDisks) GoString() string {
	return s.String()
}

func (s *DescribeDisksResponseBodyDisks) GetDisks() []*DescribeDisksResponseBodyDisksDisks {
	return s.Disks
}

func (s *DescribeDisksResponseBodyDisks) SetDisks(v []*DescribeDisksResponseBodyDisksDisks) *DescribeDisksResponseBodyDisks {
	s.Disks = v
	return s
}

func (s *DescribeDisksResponseBodyDisks) Validate() error {
	return dara.Validate(s)
}

type DescribeDisksResponseBodyDisksDisks struct {
	// The category of the disk.
	//
	// 	- cloud_efficiency: ultra disk.
	//
	// 	- cloud_ssd: all-flash disk.
	//
	// 	- local_hdd: local HDD.
	//
	// 	- local_ssd: local SSD.
	//
	// example:
	//
	// local_ssd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The time when the disk was created. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-11-11T14:34:55+08:00
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// Indicates whether the disk is released when the instance to which the disk is attached is released. Valid values:
	//
	// 	- true: The disk is released when the associated instance is released.
	//
	// 	- false: The disk is retained when the associated instance is released.
	//
	// example:
	//
	// false
	DeleteWithInstance *bool `json:"DeleteWithInstance,omitempty" xml:"DeleteWithInstance,omitempty"`
	// The namespace description.
	//
	// example:
	//
	// disk-description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The billing method of the cloud disk or local disk. Valid values:
	//
	// 	- **prepaid**: subscription.
	//
	// 	- **postpaid**: pay-as-you-go.
	//
	// example:
	//
	// prepaid
	DiskChargeType *string `json:"DiskChargeType,omitempty" xml:"DiskChargeType,omitempty"`
	// The ID of the disk.
	//
	// example:
	//
	// d-5svum1dx1w4a4spr54lgr****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The name of the disk.
	//
	// example:
	//
	// fvt-ecs-5cf0****
	DiskName *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	// Indicates whether the cloud disk is encrypted. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// False
	Encrypted *bool `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	// The ID of the Key Management Service (KMS) key that is used for the cloud disk.
	//
	// example:
	//
	// 05467897a-4262-4802-b8cb-00d3fb40****
	EncryptedKeyId *string `json:"EncryptedKeyId,omitempty" xml:"EncryptedKeyId,omitempty"`
	// The ID of the edge node.
	//
	// example:
	//
	// cn-guangzhou-10
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// i-5t77rb0yoz79m28ku60sx****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// Edge Prod Environment Streaming Machine -1063
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// Indicates whether the cloud disk or local disk is removable. Valid values:
	//
	// 	- true: The disk is removable. A removable disk can independently exist and can be attached to or detached from an instance within the same zone.
	//
	// 	- false: The disk is not removable. A disk that is not removable cannot independently exist or be attached to or detached from an instance within the same zone.
	//
	// If disks are of the following categories or types, the **Portable*	- value is **false*	- and the disks have the same lifecycle as their attached instances:
	//
	// 	- Local HDDs
	//
	// 	- Local SSDs
	//
	// 	- Data disks that use the subscription billing method
	//
	// example:
	//
	// true
	Portable *bool `json:"Portable,omitempty" xml:"Portable,omitempty"`
	// The serial number.
	//
	// example:
	//
	// 123
	SerialId *string `json:"SerialId,omitempty" xml:"SerialId,omitempty"`
	// The size of the disk. Unit: MiB.
	//
	// example:
	//
	// 20
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The ID of the snapshot.
	//
	// example:
	//
	// s-bp67acfmxazb4p****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// The status of the disk. Valid values:
	//
	// 	- In-use: The disk is in use.
	//
	// 	- Available: The disk can be attached.
	//
	// 	- Attaching: The disk is being attached.
	//
	// 	- Detaching: The disk is being detached.
	//
	// 	- Creating: The disk is being created.
	//
	// 	- ReIniting: The disk is being reset.
	//
	// example:
	//
	// Available
	Status *string                                  `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags   *DescribeDisksResponseBodyDisksDisksTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The type of the cloud disk or local disk. Valid values:
	//
	// 	- 1: system disk.
	//
	// 	- 2: data disk.
	//
	// example:
	//
	// 1
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeDisksResponseBodyDisksDisks) String() string {
	return dara.Prettify(s)
}

func (s DescribeDisksResponseBodyDisksDisks) GoString() string {
	return s.String()
}

func (s *DescribeDisksResponseBodyDisksDisks) GetCategory() *string {
	return s.Category
}

func (s *DescribeDisksResponseBodyDisksDisks) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeDisksResponseBodyDisksDisks) GetDeleteWithInstance() *bool {
	return s.DeleteWithInstance
}

func (s *DescribeDisksResponseBodyDisksDisks) GetDescription() *string {
	return s.Description
}

func (s *DescribeDisksResponseBodyDisksDisks) GetDiskChargeType() *string {
	return s.DiskChargeType
}

func (s *DescribeDisksResponseBodyDisksDisks) GetDiskId() *string {
	return s.DiskId
}

func (s *DescribeDisksResponseBodyDisksDisks) GetDiskName() *string {
	return s.DiskName
}

func (s *DescribeDisksResponseBodyDisksDisks) GetEncrypted() *bool {
	return s.Encrypted
}

func (s *DescribeDisksResponseBodyDisksDisks) GetEncryptedKeyId() *string {
	return s.EncryptedKeyId
}

func (s *DescribeDisksResponseBodyDisksDisks) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeDisksResponseBodyDisksDisks) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDisksResponseBodyDisksDisks) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeDisksResponseBodyDisksDisks) GetPortable() *bool {
	return s.Portable
}

func (s *DescribeDisksResponseBodyDisksDisks) GetSerialId() *string {
	return s.SerialId
}

func (s *DescribeDisksResponseBodyDisksDisks) GetSize() *int32 {
	return s.Size
}

func (s *DescribeDisksResponseBodyDisksDisks) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *DescribeDisksResponseBodyDisksDisks) GetStatus() *string {
	return s.Status
}

func (s *DescribeDisksResponseBodyDisksDisks) GetTags() *DescribeDisksResponseBodyDisksDisksTags {
	return s.Tags
}

func (s *DescribeDisksResponseBodyDisksDisks) GetType() *string {
	return s.Type
}

func (s *DescribeDisksResponseBodyDisksDisks) SetCategory(v string) *DescribeDisksResponseBodyDisksDisks {
	s.Category = &v
	return s
}

func (s *DescribeDisksResponseBodyDisksDisks) SetCreationTime(v string) *DescribeDisksResponseBodyDisksDisks {
	s.CreationTime = &v
	return s
}

func (s *DescribeDisksResponseBodyDisksDisks) SetDeleteWithInstance(v bool) *DescribeDisksResponseBodyDisksDisks {
	s.DeleteWithInstance = &v
	return s
}

func (s *DescribeDisksResponseBodyDisksDisks) SetDescription(v string) *DescribeDisksResponseBodyDisksDisks {
	s.Description = &v
	return s
}

func (s *DescribeDisksResponseBodyDisksDisks) SetDiskChargeType(v string) *DescribeDisksResponseBodyDisksDisks {
	s.DiskChargeType = &v
	return s
}

func (s *DescribeDisksResponseBodyDisksDisks) SetDiskId(v string) *DescribeDisksResponseBodyDisksDisks {
	s.DiskId = &v
	return s
}

func (s *DescribeDisksResponseBodyDisksDisks) SetDiskName(v string) *DescribeDisksResponseBodyDisksDisks {
	s.DiskName = &v
	return s
}

func (s *DescribeDisksResponseBodyDisksDisks) SetEncrypted(v bool) *DescribeDisksResponseBodyDisksDisks {
	s.Encrypted = &v
	return s
}

func (s *DescribeDisksResponseBodyDisksDisks) SetEncryptedKeyId(v string) *DescribeDisksResponseBodyDisksDisks {
	s.EncryptedKeyId = &v
	return s
}

func (s *DescribeDisksResponseBodyDisksDisks) SetEnsRegionId(v string) *DescribeDisksResponseBodyDisksDisks {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeDisksResponseBodyDisksDisks) SetInstanceId(v string) *DescribeDisksResponseBodyDisksDisks {
	s.InstanceId = &v
	return s
}

func (s *DescribeDisksResponseBodyDisksDisks) SetInstanceName(v string) *DescribeDisksResponseBodyDisksDisks {
	s.InstanceName = &v
	return s
}

func (s *DescribeDisksResponseBodyDisksDisks) SetPortable(v bool) *DescribeDisksResponseBodyDisksDisks {
	s.Portable = &v
	return s
}

func (s *DescribeDisksResponseBodyDisksDisks) SetSerialId(v string) *DescribeDisksResponseBodyDisksDisks {
	s.SerialId = &v
	return s
}

func (s *DescribeDisksResponseBodyDisksDisks) SetSize(v int32) *DescribeDisksResponseBodyDisksDisks {
	s.Size = &v
	return s
}

func (s *DescribeDisksResponseBodyDisksDisks) SetSnapshotId(v string) *DescribeDisksResponseBodyDisksDisks {
	s.SnapshotId = &v
	return s
}

func (s *DescribeDisksResponseBodyDisksDisks) SetStatus(v string) *DescribeDisksResponseBodyDisksDisks {
	s.Status = &v
	return s
}

func (s *DescribeDisksResponseBodyDisksDisks) SetTags(v *DescribeDisksResponseBodyDisksDisksTags) *DescribeDisksResponseBodyDisksDisks {
	s.Tags = v
	return s
}

func (s *DescribeDisksResponseBodyDisksDisks) SetType(v string) *DescribeDisksResponseBodyDisksDisks {
	s.Type = &v
	return s
}

func (s *DescribeDisksResponseBodyDisksDisks) Validate() error {
	return dara.Validate(s)
}

type DescribeDisksResponseBodyDisksDisksTags struct {
	Tags []*DescribeDisksResponseBodyDisksDisksTagsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DescribeDisksResponseBodyDisksDisksTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeDisksResponseBodyDisksDisksTags) GoString() string {
	return s.String()
}

func (s *DescribeDisksResponseBodyDisksDisksTags) GetTags() []*DescribeDisksResponseBodyDisksDisksTagsTags {
	return s.Tags
}

func (s *DescribeDisksResponseBodyDisksDisksTags) SetTags(v []*DescribeDisksResponseBodyDisksDisksTagsTags) *DescribeDisksResponseBodyDisksDisksTags {
	s.Tags = v
	return s
}

func (s *DescribeDisksResponseBodyDisksDisksTags) Validate() error {
	return dara.Validate(s)
}

type DescribeDisksResponseBodyDisksDisksTagsTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDisksResponseBodyDisksDisksTagsTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeDisksResponseBodyDisksDisksTagsTags) GoString() string {
	return s.String()
}

func (s *DescribeDisksResponseBodyDisksDisksTagsTags) GetKey() *string {
	return s.Key
}

func (s *DescribeDisksResponseBodyDisksDisksTagsTags) GetValue() *string {
	return s.Value
}

func (s *DescribeDisksResponseBodyDisksDisksTagsTags) SetKey(v string) *DescribeDisksResponseBodyDisksDisksTagsTags {
	s.Key = &v
	return s
}

func (s *DescribeDisksResponseBodyDisksDisksTagsTags) SetValue(v string) *DescribeDisksResponseBodyDisksDisksTagsTags {
	s.Value = &v
	return s
}

func (s *DescribeDisksResponseBodyDisksDisksTagsTags) Validate() error {
	return dara.Validate(s)
}
