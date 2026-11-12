// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSnapshotsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListSnapshotsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListSnapshotsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListSnapshotsResponseBody
	GetRequestId() *string
	SetSnapshots(v []*ListSnapshotsResponseBodySnapshots) *ListSnapshotsResponseBody
	GetSnapshots() []*ListSnapshotsResponseBodySnapshots
	SetTotalCount(v int32) *ListSnapshotsResponseBody
	GetTotalCount() *int32
}

type ListSnapshotsResponseBody struct {
	// The page number of the snapshot list.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page for a paged query.
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
	// The snapshot information.
	Snapshots []*ListSnapshotsResponseBodySnapshots `json:"Snapshots,omitempty" xml:"Snapshots,omitempty" type:"Repeated"`
	// The total number of snapshots.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSnapshotsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSnapshotsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSnapshotsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSnapshotsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSnapshotsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSnapshotsResponseBody) GetSnapshots() []*ListSnapshotsResponseBodySnapshots {
	return s.Snapshots
}

func (s *ListSnapshotsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListSnapshotsResponseBody) SetPageNumber(v int32) *ListSnapshotsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListSnapshotsResponseBody) SetPageSize(v int32) *ListSnapshotsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListSnapshotsResponseBody) SetRequestId(v string) *ListSnapshotsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSnapshotsResponseBody) SetSnapshots(v []*ListSnapshotsResponseBodySnapshots) *ListSnapshotsResponseBody {
	s.Snapshots = v
	return s
}

func (s *ListSnapshotsResponseBody) SetTotalCount(v int32) *ListSnapshotsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSnapshotsResponseBody) Validate() error {
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

type ListSnapshotsResponseBodySnapshots struct {
	// The time when the snapshot was created. The time follows the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-03-09T07:12:49Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The instance ID of the simple application server.
	//
	// Note: This field has a value for system cloud disk snapshots and is empty for data cloud disk snapshots.
	//
	// example:
	//
	// 2ad1ae67295445f598017499dc****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The progress of snapshot creation.
	//
	// example:
	//
	// 100%
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The remarks of the snapshot.
	//
	// example:
	//
	// test-Remark
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The ID of the resource group to which the snapshot belongs.
	//
	// example:
	//
	// rg-aek2bti7cf7****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The time of the most recent rollback. The time is in the UTC (Z) time zone in the YYYY-MM-DDThh:mm:ssZ format.
	//
	// example:
	//
	// 2021-03-09T07:12:49Z
	RollbackTime *string `json:"RollbackTime,omitempty" xml:"RollbackTime,omitempty"`
	// The snapshot ID.
	//
	// example:
	//
	// s-bp16oazlsold4dks****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// The snapshot name.
	//
	// example:
	//
	// test-SnapshotName
	SnapshotName *string `json:"SnapshotName,omitempty" xml:"SnapshotName,omitempty"`
	// The ID of the source disk. This value is retained even if the source disk of the snapshot has been released.
	//
	// example:
	//
	// d-bp14wq0149cpp2xy****
	SourceDiskId *string `json:"SourceDiskId,omitempty" xml:"SourceDiskId,omitempty"`
	// The type of the source disk for the snapshot. Valid values:
	//
	// - System: system cloud disk.
	//
	// - Data: data cloud disk.
	//
	// example:
	//
	// System
	SourceDiskType *string `json:"SourceDiskType,omitempty" xml:"SourceDiskType,omitempty"`
	// The snapshot status. Valid values:
	//
	// - progressing: The snapshot is being created.
	//
	// - accomplished: The snapshot is created.
	//
	// - failed: The snapshot failed to be created.
	//
	// - disabled: Rollback is not supported.
	//
	// example:
	//
	// Accomplished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags of the snapshot.
	Tags []*ListSnapshotsResponseBodySnapshotsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListSnapshotsResponseBodySnapshots) String() string {
	return dara.Prettify(s)
}

func (s ListSnapshotsResponseBodySnapshots) GoString() string {
	return s.String()
}

func (s *ListSnapshotsResponseBodySnapshots) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListSnapshotsResponseBodySnapshots) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListSnapshotsResponseBodySnapshots) GetProgress() *string {
	return s.Progress
}

func (s *ListSnapshotsResponseBodySnapshots) GetRegionId() *string {
	return s.RegionId
}

func (s *ListSnapshotsResponseBodySnapshots) GetRemark() *string {
	return s.Remark
}

func (s *ListSnapshotsResponseBodySnapshots) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListSnapshotsResponseBodySnapshots) GetRollbackTime() *string {
	return s.RollbackTime
}

func (s *ListSnapshotsResponseBodySnapshots) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *ListSnapshotsResponseBodySnapshots) GetSnapshotName() *string {
	return s.SnapshotName
}

func (s *ListSnapshotsResponseBodySnapshots) GetSourceDiskId() *string {
	return s.SourceDiskId
}

func (s *ListSnapshotsResponseBodySnapshots) GetSourceDiskType() *string {
	return s.SourceDiskType
}

func (s *ListSnapshotsResponseBodySnapshots) GetStatus() *string {
	return s.Status
}

func (s *ListSnapshotsResponseBodySnapshots) GetTags() []*ListSnapshotsResponseBodySnapshotsTags {
	return s.Tags
}

func (s *ListSnapshotsResponseBodySnapshots) SetCreationTime(v string) *ListSnapshotsResponseBodySnapshots {
	s.CreationTime = &v
	return s
}

func (s *ListSnapshotsResponseBodySnapshots) SetInstanceId(v string) *ListSnapshotsResponseBodySnapshots {
	s.InstanceId = &v
	return s
}

func (s *ListSnapshotsResponseBodySnapshots) SetProgress(v string) *ListSnapshotsResponseBodySnapshots {
	s.Progress = &v
	return s
}

func (s *ListSnapshotsResponseBodySnapshots) SetRegionId(v string) *ListSnapshotsResponseBodySnapshots {
	s.RegionId = &v
	return s
}

func (s *ListSnapshotsResponseBodySnapshots) SetRemark(v string) *ListSnapshotsResponseBodySnapshots {
	s.Remark = &v
	return s
}

func (s *ListSnapshotsResponseBodySnapshots) SetResourceGroupId(v string) *ListSnapshotsResponseBodySnapshots {
	s.ResourceGroupId = &v
	return s
}

func (s *ListSnapshotsResponseBodySnapshots) SetRollbackTime(v string) *ListSnapshotsResponseBodySnapshots {
	s.RollbackTime = &v
	return s
}

func (s *ListSnapshotsResponseBodySnapshots) SetSnapshotId(v string) *ListSnapshotsResponseBodySnapshots {
	s.SnapshotId = &v
	return s
}

func (s *ListSnapshotsResponseBodySnapshots) SetSnapshotName(v string) *ListSnapshotsResponseBodySnapshots {
	s.SnapshotName = &v
	return s
}

func (s *ListSnapshotsResponseBodySnapshots) SetSourceDiskId(v string) *ListSnapshotsResponseBodySnapshots {
	s.SourceDiskId = &v
	return s
}

func (s *ListSnapshotsResponseBodySnapshots) SetSourceDiskType(v string) *ListSnapshotsResponseBodySnapshots {
	s.SourceDiskType = &v
	return s
}

func (s *ListSnapshotsResponseBodySnapshots) SetStatus(v string) *ListSnapshotsResponseBodySnapshots {
	s.Status = &v
	return s
}

func (s *ListSnapshotsResponseBodySnapshots) SetTags(v []*ListSnapshotsResponseBodySnapshotsTags) *ListSnapshotsResponseBodySnapshots {
	s.Tags = v
	return s
}

func (s *ListSnapshotsResponseBodySnapshots) Validate() error {
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

type ListSnapshotsResponseBodySnapshotsTags struct {
	// The tag key of the snapshot.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the snapshot.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListSnapshotsResponseBodySnapshotsTags) String() string {
	return dara.Prettify(s)
}

func (s ListSnapshotsResponseBodySnapshotsTags) GoString() string {
	return s.String()
}

func (s *ListSnapshotsResponseBodySnapshotsTags) GetKey() *string {
	return s.Key
}

func (s *ListSnapshotsResponseBodySnapshotsTags) GetValue() *string {
	return s.Value
}

func (s *ListSnapshotsResponseBodySnapshotsTags) SetKey(v string) *ListSnapshotsResponseBodySnapshotsTags {
	s.Key = &v
	return s
}

func (s *ListSnapshotsResponseBodySnapshotsTags) SetValue(v string) *ListSnapshotsResponseBodySnapshotsTags {
	s.Value = &v
	return s
}

func (s *ListSnapshotsResponseBodySnapshotsTags) Validate() error {
	return dara.Validate(s)
}
