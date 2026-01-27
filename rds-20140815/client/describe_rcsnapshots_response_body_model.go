// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCSnapshotsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int64) *DescribeRCSnapshotsResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeRCSnapshotsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeRCSnapshotsResponseBody
	GetRequestId() *string
	SetSnapshots(v []*DescribeRCSnapshotsResponseBodySnapshots) *DescribeRCSnapshotsResponseBody
	GetSnapshots() []*DescribeRCSnapshotsResponseBodySnapshots
	SetTotalCount(v int64) *DescribeRCSnapshotsResponseBody
	GetTotalCount() *int64
}

type DescribeRCSnapshotsResponseBody struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 30
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9DAC759A-F4F0-5D02-8335-BC458C0CCB94
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of snapshots.
	Snapshots []*DescribeRCSnapshotsResponseBodySnapshots `json:"Snapshots,omitempty" xml:"Snapshots,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 7
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeRCSnapshotsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCSnapshotsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRCSnapshotsResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeRCSnapshotsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeRCSnapshotsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRCSnapshotsResponseBody) GetSnapshots() []*DescribeRCSnapshotsResponseBodySnapshots {
	return s.Snapshots
}

func (s *DescribeRCSnapshotsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeRCSnapshotsResponseBody) SetPageNumber(v int64) *DescribeRCSnapshotsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeRCSnapshotsResponseBody) SetPageSize(v int64) *DescribeRCSnapshotsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeRCSnapshotsResponseBody) SetRequestId(v string) *DescribeRCSnapshotsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRCSnapshotsResponseBody) SetSnapshots(v []*DescribeRCSnapshotsResponseBodySnapshots) *DescribeRCSnapshotsResponseBody {
	s.Snapshots = v
	return s
}

func (s *DescribeRCSnapshotsResponseBody) SetTotalCount(v int64) *DescribeRCSnapshotsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeRCSnapshotsResponseBody) Validate() error {
	if s.Snapshots != nil {
		for _, item := range s.Snapshots {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRCSnapshotsResponseBodySnapshots struct {
	// Indicates whether the snapshot can be shared and used to create or roll back a cloud disk. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Available *bool `json:"Available,omitempty" xml:"Available,omitempty"`
	// The snapshot type. Valid values:
	//
	// 	- Standard: standard snapshot
	//
	// 	- Flash: local snapshot This value will be deprecated. The local snapshot feature is replaced with the instant access feature.
	//
	// 	- archive: archived snapshot
	//
	// example:
	//
	// Standard
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The creation time. The time follows the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2024-10-18T09:37:14Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The snapshot description.
	//
	// example:
	//
	// zd_test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the snapshot was encrypted. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Encrypted *bool `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// none
	InstantAccess *bool `json:"InstantAccess,omitempty" xml:"InstantAccess,omitempty"`
	// The progress of the snapshot creation task in percentage.
	//
	// example:
	//
	// 100
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The snapshot ID.
	//
	// example:
	//
	// rcds-hc1zg51xobdg4****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// The snapshot name.
	//
	// example:
	//
	// s-2ze8klip00xcogcwer76
	SnapshotName *string `json:"SnapshotName,omitempty" xml:"SnapshotName,omitempty"`
	// The snapshot type. Valid values:
	//
	// 	- auto or timer: automatically created snapshot
	//
	// 	- user: manually created snapshot
	//
	// 	- all: all snapshot types
	//
	// example:
	//
	// auto
	SnapshotType *string `json:"SnapshotType,omitempty" xml:"SnapshotType,omitempty"`
	// The ID of the original disk. This parameter is retained even after the original disk for which the snapshot was created is released.
	//
	// example:
	//
	// rcd-bp67acfmxazb4ph****
	SourceDiskId *string `json:"SourceDiskId,omitempty" xml:"SourceDiskId,omitempty"`
	// The storage capacity of the original disk. Unit: GiB.
	//
	// example:
	//
	// 60
	SourceDiskSize *int64 `json:"SourceDiskSize,omitempty" xml:"SourceDiskSize,omitempty"`
	// The type of the original disk. Valid values:
	//
	// 	- SYSTEM: system disk
	//
	// 	- DATA: data disk
	//
	// example:
	//
	// data
	SourceDiskType *string `json:"SourceDiskType,omitempty" xml:"SourceDiskType,omitempty"`
	// The type of the source disk.
	//
	// >  This parameter will be removed in the future. To ensure future compatibility, we recommend that you use other parameters.
	//
	// example:
	//
	// disk
	SourceStorageType *string `json:"SourceStorageType,omitempty" xml:"SourceStorageType,omitempty"`
	// The snapshot status. Valid values:
	//
	// 	- progressing: The snapshot is being created.
	//
	// 	- accomplished: The snapshot is created.
	//
	// 	- failed: The snapshot fails to be created.
	//
	// example:
	//
	// progressing
	Status *string                                        `json:"Status,omitempty" xml:"Status,omitempty"`
	Tag    []*DescribeRCSnapshotsResponseBodySnapshotsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// Indicates whether the snapshot is used to create custom images or disks. Valid values:
	//
	// 	- image: The snapshot is used to create custom images.
	//
	// 	- disk: The snapshot is used to create disks.
	//
	// 	- image_disk: The snapshot is used to create custom images and data disks.
	//
	// 	- none: The snapshot is not used to create custom images or disks.
	//
	// example:
	//
	// none
	Usage *string `json:"Usage,omitempty" xml:"Usage,omitempty"`
}

func (s DescribeRCSnapshotsResponseBodySnapshots) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCSnapshotsResponseBodySnapshots) GoString() string {
	return s.String()
}

func (s *DescribeRCSnapshotsResponseBodySnapshots) GetAvailable() *bool {
	return s.Available
}

func (s *DescribeRCSnapshotsResponseBodySnapshots) GetCategory() *string {
	return s.Category
}

func (s *DescribeRCSnapshotsResponseBodySnapshots) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeRCSnapshotsResponseBodySnapshots) GetDescription() *string {
	return s.Description
}

func (s *DescribeRCSnapshotsResponseBodySnapshots) GetEncrypted() *bool {
	return s.Encrypted
}

func (s *DescribeRCSnapshotsResponseBodySnapshots) GetInstantAccess() *bool {
	return s.InstantAccess
}

func (s *DescribeRCSnapshotsResponseBodySnapshots) GetProgress() *string {
	return s.Progress
}

func (s *DescribeRCSnapshotsResponseBodySnapshots) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRCSnapshotsResponseBodySnapshots) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeRCSnapshotsResponseBodySnapshots) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *DescribeRCSnapshotsResponseBodySnapshots) GetSnapshotName() *string {
	return s.SnapshotName
}

func (s *DescribeRCSnapshotsResponseBodySnapshots) GetSnapshotType() *string {
	return s.SnapshotType
}

func (s *DescribeRCSnapshotsResponseBodySnapshots) GetSourceDiskId() *string {
	return s.SourceDiskId
}

func (s *DescribeRCSnapshotsResponseBodySnapshots) GetSourceDiskSize() *int64 {
	return s.SourceDiskSize
}

func (s *DescribeRCSnapshotsResponseBodySnapshots) GetSourceDiskType() *string {
	return s.SourceDiskType
}

func (s *DescribeRCSnapshotsResponseBodySnapshots) GetSourceStorageType() *string {
	return s.SourceStorageType
}

func (s *DescribeRCSnapshotsResponseBodySnapshots) GetStatus() *string {
	return s.Status
}

func (s *DescribeRCSnapshotsResponseBodySnapshots) GetTag() []*DescribeRCSnapshotsResponseBodySnapshotsTag {
	return s.Tag
}

func (s *DescribeRCSnapshotsResponseBodySnapshots) GetUsage() *string {
	return s.Usage
}

func (s *DescribeRCSnapshotsResponseBodySnapshots) SetAvailable(v bool) *DescribeRCSnapshotsResponseBodySnapshots {
	s.Available = &v
	return s
}

func (s *DescribeRCSnapshotsResponseBodySnapshots) SetCategory(v string) *DescribeRCSnapshotsResponseBodySnapshots {
	s.Category = &v
	return s
}

func (s *DescribeRCSnapshotsResponseBodySnapshots) SetCreationTime(v string) *DescribeRCSnapshotsResponseBodySnapshots {
	s.CreationTime = &v
	return s
}

func (s *DescribeRCSnapshotsResponseBodySnapshots) SetDescription(v string) *DescribeRCSnapshotsResponseBodySnapshots {
	s.Description = &v
	return s
}

func (s *DescribeRCSnapshotsResponseBodySnapshots) SetEncrypted(v bool) *DescribeRCSnapshotsResponseBodySnapshots {
	s.Encrypted = &v
	return s
}

func (s *DescribeRCSnapshotsResponseBodySnapshots) SetInstantAccess(v bool) *DescribeRCSnapshotsResponseBodySnapshots {
	s.InstantAccess = &v
	return s
}

func (s *DescribeRCSnapshotsResponseBodySnapshots) SetProgress(v string) *DescribeRCSnapshotsResponseBodySnapshots {
	s.Progress = &v
	return s
}

func (s *DescribeRCSnapshotsResponseBodySnapshots) SetRegionId(v string) *DescribeRCSnapshotsResponseBodySnapshots {
	s.RegionId = &v
	return s
}

func (s *DescribeRCSnapshotsResponseBodySnapshots) SetResourceGroupId(v string) *DescribeRCSnapshotsResponseBodySnapshots {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeRCSnapshotsResponseBodySnapshots) SetSnapshotId(v string) *DescribeRCSnapshotsResponseBodySnapshots {
	s.SnapshotId = &v
	return s
}

func (s *DescribeRCSnapshotsResponseBodySnapshots) SetSnapshotName(v string) *DescribeRCSnapshotsResponseBodySnapshots {
	s.SnapshotName = &v
	return s
}

func (s *DescribeRCSnapshotsResponseBodySnapshots) SetSnapshotType(v string) *DescribeRCSnapshotsResponseBodySnapshots {
	s.SnapshotType = &v
	return s
}

func (s *DescribeRCSnapshotsResponseBodySnapshots) SetSourceDiskId(v string) *DescribeRCSnapshotsResponseBodySnapshots {
	s.SourceDiskId = &v
	return s
}

func (s *DescribeRCSnapshotsResponseBodySnapshots) SetSourceDiskSize(v int64) *DescribeRCSnapshotsResponseBodySnapshots {
	s.SourceDiskSize = &v
	return s
}

func (s *DescribeRCSnapshotsResponseBodySnapshots) SetSourceDiskType(v string) *DescribeRCSnapshotsResponseBodySnapshots {
	s.SourceDiskType = &v
	return s
}

func (s *DescribeRCSnapshotsResponseBodySnapshots) SetSourceStorageType(v string) *DescribeRCSnapshotsResponseBodySnapshots {
	s.SourceStorageType = &v
	return s
}

func (s *DescribeRCSnapshotsResponseBodySnapshots) SetStatus(v string) *DescribeRCSnapshotsResponseBodySnapshots {
	s.Status = &v
	return s
}

func (s *DescribeRCSnapshotsResponseBodySnapshots) SetTag(v []*DescribeRCSnapshotsResponseBodySnapshotsTag) *DescribeRCSnapshotsResponseBodySnapshots {
	s.Tag = v
	return s
}

func (s *DescribeRCSnapshotsResponseBodySnapshots) SetUsage(v string) *DescribeRCSnapshotsResponseBodySnapshots {
	s.Usage = &v
	return s
}

func (s *DescribeRCSnapshotsResponseBodySnapshots) Validate() error {
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

type DescribeRCSnapshotsResponseBodySnapshotsTag struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeRCSnapshotsResponseBodySnapshotsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCSnapshotsResponseBodySnapshotsTag) GoString() string {
	return s.String()
}

func (s *DescribeRCSnapshotsResponseBodySnapshotsTag) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeRCSnapshotsResponseBodySnapshotsTag) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeRCSnapshotsResponseBodySnapshotsTag) SetTagKey(v string) *DescribeRCSnapshotsResponseBodySnapshotsTag {
	s.TagKey = &v
	return s
}

func (s *DescribeRCSnapshotsResponseBodySnapshotsTag) SetTagValue(v string) *DescribeRCSnapshotsResponseBodySnapshotsTag {
	s.TagValue = &v
	return s
}

func (s *DescribeRCSnapshotsResponseBodySnapshotsTag) Validate() error {
	return dara.Validate(s)
}
