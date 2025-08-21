// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSnapshotsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeSnapshotsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSnapshotsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeSnapshotsResponseBody
	GetRequestId() *string
	SetSnapshots(v []*DescribeSnapshotsResponseBodySnapshots) *DescribeSnapshotsResponseBody
	GetSnapshots() []*DescribeSnapshotsResponseBodySnapshots
	SetTotalCount(v int32) *DescribeSnapshotsResponseBody
	GetTotalCount() *int32
}

type DescribeSnapshotsResponseBody struct {
	// The page number of the returned page.
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
	// 9635ED2E-3103-1606-84D4-9F8E816B19F9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the snapshots.
	Snapshots []*DescribeSnapshotsResponseBodySnapshots `json:"Snapshots,omitempty" xml:"Snapshots,omitempty" type:"Repeated"`
	// The total number of snapshots.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSnapshotsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnapshotsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSnapshotsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSnapshotsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSnapshotsResponseBody) GetSnapshots() []*DescribeSnapshotsResponseBodySnapshots {
	return s.Snapshots
}

func (s *DescribeSnapshotsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeSnapshotsResponseBody) SetPageNumber(v int32) *DescribeSnapshotsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSnapshotsResponseBody) SetPageSize(v int32) *DescribeSnapshotsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSnapshotsResponseBody) SetRequestId(v string) *DescribeSnapshotsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSnapshotsResponseBody) SetSnapshots(v []*DescribeSnapshotsResponseBodySnapshots) *DescribeSnapshotsResponseBody {
	s.Snapshots = v
	return s
}

func (s *DescribeSnapshotsResponseBody) SetTotalCount(v int32) *DescribeSnapshotsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSnapshotsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeSnapshotsResponseBodySnapshots struct {
	// The creation time. The time follows the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-08-20T14:52:28Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the snapshot.
	//
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the edge node.
	//
	// example:
	//
	// cn-beijing-15
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The capacity of the disk. Unit: MiB.
	//
	// example:
	//
	// 40
	Size *string `json:"Size,omitempty" xml:"Size,omitempty"`
	// The ID of the snapshot.
	//
	// example:
	//
	// s-bp67acfmxazb4p****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// The name of the snapshot. This parameter is returned only if a snapshot name was specified when the snapshot was created.
	//
	// example:
	//
	// testSnapshotName
	SnapshotName *string `json:"SnapshotName,omitempty" xml:"SnapshotName,omitempty"`
	// The type of the disk. Valid value:
	//
	// 	- cloud_efficiency: ultra disk
	//
	// 	- cloud_ssd: all-flash disk
	//
	// 	- local_hdd: local HDD
	//
	// 	- local_ssd: local SSD
	//
	// example:
	//
	// cloud_efficiency
	SourceDiskCategory *string `json:"SourceDiskCategory,omitempty" xml:"SourceDiskCategory,omitempty"`
	// The ID of the source disk. This parameter is retained even after the source disk for which the snapshot was created is released.
	//
	// example:
	//
	// d-bp67acfmxazb4ph****
	SourceDiskId *string `json:"SourceDiskId,omitempty" xml:"SourceDiskId,omitempty"`
	// The type of the disk. Valid value:
	//
	// 	- 1: system disk
	//
	// 	- 2: data disk
	//
	// example:
	//
	// 1
	SourceDiskType *string `json:"SourceDiskType,omitempty" xml:"SourceDiskType,omitempty"`
	// The ID of the source edge node.
	//
	// example:
	//
	// cn-hangzhou-27
	SourceEnsRegionId *string `json:"SourceEnsRegionId,omitempty" xml:"SourceEnsRegionId,omitempty"`
	// The ID of the source snapshot.
	//
	// example:
	//
	// s-bpdfer893jfkdqe****
	SourceSnapshotId *string `json:"SourceSnapshotId,omitempty" xml:"SourceSnapshotId,omitempty"`
	// The status of the snapshot. Valid value:
	//
	// 	- creating: The snapshot is being created.
	//
	// 	- Available: The snapshot is available.
	//
	// 	- deleting: The snapshot is being deleted.
	//
	// 	- error: An error occurred on the snapshot.
	//
	// example:
	//
	// available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeSnapshotsResponseBodySnapshots) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnapshotsResponseBodySnapshots) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotsResponseBodySnapshots) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeSnapshotsResponseBodySnapshots) GetDescription() *string {
	return s.Description
}

func (s *DescribeSnapshotsResponseBodySnapshots) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeSnapshotsResponseBodySnapshots) GetSize() *string {
	return s.Size
}

func (s *DescribeSnapshotsResponseBodySnapshots) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *DescribeSnapshotsResponseBodySnapshots) GetSnapshotName() *string {
	return s.SnapshotName
}

func (s *DescribeSnapshotsResponseBodySnapshots) GetSourceDiskCategory() *string {
	return s.SourceDiskCategory
}

func (s *DescribeSnapshotsResponseBodySnapshots) GetSourceDiskId() *string {
	return s.SourceDiskId
}

func (s *DescribeSnapshotsResponseBodySnapshots) GetSourceDiskType() *string {
	return s.SourceDiskType
}

func (s *DescribeSnapshotsResponseBodySnapshots) GetSourceEnsRegionId() *string {
	return s.SourceEnsRegionId
}

func (s *DescribeSnapshotsResponseBodySnapshots) GetSourceSnapshotId() *string {
	return s.SourceSnapshotId
}

func (s *DescribeSnapshotsResponseBodySnapshots) GetStatus() *string {
	return s.Status
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetCreationTime(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.CreationTime = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetDescription(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.Description = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetEnsRegionId(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetSize(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.Size = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetSnapshotId(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.SnapshotId = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetSnapshotName(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.SnapshotName = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetSourceDiskCategory(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.SourceDiskCategory = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetSourceDiskId(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.SourceDiskId = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetSourceDiskType(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.SourceDiskType = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetSourceEnsRegionId(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.SourceEnsRegionId = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetSourceSnapshotId(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.SourceSnapshotId = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetStatus(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.Status = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) Validate() error {
	return dara.Validate(s)
}
