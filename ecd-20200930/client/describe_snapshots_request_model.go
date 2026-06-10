// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSnapshotsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreator(v string) *DescribeSnapshotsRequest
	GetCreator() *string
	SetDesktopId(v string) *DescribeSnapshotsRequest
	GetDesktopId() *string
	SetDesktopName(v string) *DescribeSnapshotsRequest
	GetDesktopName() *string
	SetEndTime(v string) *DescribeSnapshotsRequest
	GetEndTime() *string
	SetMaxResults(v int32) *DescribeSnapshotsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeSnapshotsRequest
	GetNextToken() *string
	SetOsType(v string) *DescribeSnapshotsRequest
	GetOsType() *string
	SetRegionId(v string) *DescribeSnapshotsRequest
	GetRegionId() *string
	SetSnapshotId(v string) *DescribeSnapshotsRequest
	GetSnapshotId() *string
	SetSnapshotName(v string) *DescribeSnapshotsRequest
	GetSnapshotName() *string
	SetSnapshotType(v string) *DescribeSnapshotsRequest
	GetSnapshotType() *string
	SetSourceDiskType(v string) *DescribeSnapshotsRequest
	GetSourceDiskType() *string
	SetStartTime(v string) *DescribeSnapshotsRequest
	GetStartTime() *string
}

type DescribeSnapshotsRequest struct {
	// The creator.
	//
	// example:
	//
	// Administrator
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The ID of the cloud desktop.
	//
	// example:
	//
	// ecd-gx2x1dhsmucyy****
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// The name of the cloud desktop.
	//
	// example:
	//
	// testName
	DesktopName *string `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	// The end time to query for snapshots. The time follows the [ISO 8601](t10049.xdita#) standard and is in UTC. The format is `yyyy-mm-ddthh:mm:ssz`.
	//
	// example:
	//
	// 2020-11-31T06:32:31Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The number of entries per page.
	//
	// - Maximum value: 100.
	//
	// - Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next page of results. This is the NextToken value from the previous API call.
	//
	// example:
	//
	// 8051af8d01b5479bec9f5ddf02e4a8fbd0ab6e7e43f8****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The operating system type.
	//
	// example:
	//
	// Windows
	OsType *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	// The ID of the region. Call [](t2167755.xdita#)to get a list of regions that support Elastic Desktop Service (EDS).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The snapshot ID.
	//
	// example:
	//
	// s-2ze81owrnv9pity4****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// The display name of the snapshot. The name must be 2 to 127 characters long. It must start with a letter. It can contain digits, underscores (_), and hyphens (-). The name cannot start with `auto` to avoid naming conflicts with automatic snapshots.
	//
	// example:
	//
	// Test data disk
	SnapshotName *string `json:"SnapshotName,omitempty" xml:"SnapshotName,omitempty"`
	// The snapshot type.
	//
	// example:
	//
	// user
	SnapshotType *string `json:"SnapshotType,omitempty" xml:"SnapshotType,omitempty"`
	// The disk from which to create the snapshot.
	//
	// > The value is case-insensitive.
	//
	// example:
	//
	// system
	SourceDiskType *string `json:"SourceDiskType,omitempty" xml:"SourceDiskType,omitempty"`
	// The start time to query for snapshots. The time follows the [ISO 8601](t10049.xdita#) standard and is in UTC. The format is `yyyy-mm-ddthh:mm:ssz`.
	//
	// example:
	//
	// 2020-11-30T06:32:31Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeSnapshotsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnapshotsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotsRequest) GetCreator() *string {
	return s.Creator
}

func (s *DescribeSnapshotsRequest) GetDesktopId() *string {
	return s.DesktopId
}

func (s *DescribeSnapshotsRequest) GetDesktopName() *string {
	return s.DesktopName
}

func (s *DescribeSnapshotsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeSnapshotsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeSnapshotsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeSnapshotsRequest) GetOsType() *string {
	return s.OsType
}

func (s *DescribeSnapshotsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSnapshotsRequest) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *DescribeSnapshotsRequest) GetSnapshotName() *string {
	return s.SnapshotName
}

func (s *DescribeSnapshotsRequest) GetSnapshotType() *string {
	return s.SnapshotType
}

func (s *DescribeSnapshotsRequest) GetSourceDiskType() *string {
	return s.SourceDiskType
}

func (s *DescribeSnapshotsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeSnapshotsRequest) SetCreator(v string) *DescribeSnapshotsRequest {
	s.Creator = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetDesktopId(v string) *DescribeSnapshotsRequest {
	s.DesktopId = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetDesktopName(v string) *DescribeSnapshotsRequest {
	s.DesktopName = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetEndTime(v string) *DescribeSnapshotsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetMaxResults(v int32) *DescribeSnapshotsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetNextToken(v string) *DescribeSnapshotsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetOsType(v string) *DescribeSnapshotsRequest {
	s.OsType = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetRegionId(v string) *DescribeSnapshotsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetSnapshotId(v string) *DescribeSnapshotsRequest {
	s.SnapshotId = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetSnapshotName(v string) *DescribeSnapshotsRequest {
	s.SnapshotName = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetSnapshotType(v string) *DescribeSnapshotsRequest {
	s.SnapshotType = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetSourceDiskType(v string) *DescribeSnapshotsRequest {
	s.SourceDiskType = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetStartTime(v string) *DescribeSnapshotsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSnapshotsRequest) Validate() error {
	return dara.Validate(s)
}
