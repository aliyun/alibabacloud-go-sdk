// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUdmSnapshotsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDiskId(v string) *DescribeUdmSnapshotsShrinkRequest
	GetDiskId() *string
	SetEndTime(v int64) *DescribeUdmSnapshotsShrinkRequest
	GetEndTime() *int64
	SetInstanceId(v string) *DescribeUdmSnapshotsShrinkRequest
	GetInstanceId() *string
	SetJobId(v string) *DescribeUdmSnapshotsShrinkRequest
	GetJobId() *string
	SetSnapshotIdsShrink(v string) *DescribeUdmSnapshotsShrinkRequest
	GetSnapshotIdsShrink() *string
	SetSourceType(v string) *DescribeUdmSnapshotsShrinkRequest
	GetSourceType() *string
	SetStartTime(v int64) *DescribeUdmSnapshotsShrinkRequest
	GetStartTime() *int64
	SetUdmRegionId(v string) *DescribeUdmSnapshotsShrinkRequest
	GetUdmRegionId() *string
}

type DescribeUdmSnapshotsShrinkRequest struct {
	// The ID of the disk.
	//
	// example:
	//
	// d-bp1560750pclffpzxy70
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The end of the time range to query. The value must be a UNIX timestamp. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1643092168
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the ECS instance.
	//
	// example:
	//
	// i-bp18x2k7sw925ir7ofh8
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the backup job.
	//
	// example:
	//
	// job-*********************
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The list of backup snapshots.
	//
	// example:
	//
	// [\\"s-000e3vhhu62xsm6v92r0\\"]
	SnapshotIdsShrink *string `json:"SnapshotIds,omitempty" xml:"SnapshotIds,omitempty"`
	// The type of the data source. Valid values:
	//
	// 	- **UDM_ECS**: ECS instance backup
	//
	// 	- **UDM_ECS_DISK**: disk backup subtask of ECS instance backup
	//
	// 	- **UDM_DISK**: disk backup
	//
	// This parameter is required.
	//
	// example:
	//
	// UDM_ECS
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The beginning of the time range to query. The value must be a UNIX timestamp. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1642057551
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The ID of the region where the ECS instance resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	UdmRegionId *string `json:"UdmRegionId,omitempty" xml:"UdmRegionId,omitempty"`
}

func (s DescribeUdmSnapshotsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUdmSnapshotsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeUdmSnapshotsShrinkRequest) GetDiskId() *string {
	return s.DiskId
}

func (s *DescribeUdmSnapshotsShrinkRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeUdmSnapshotsShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeUdmSnapshotsShrinkRequest) GetJobId() *string {
	return s.JobId
}

func (s *DescribeUdmSnapshotsShrinkRequest) GetSnapshotIdsShrink() *string {
	return s.SnapshotIdsShrink
}

func (s *DescribeUdmSnapshotsShrinkRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *DescribeUdmSnapshotsShrinkRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeUdmSnapshotsShrinkRequest) GetUdmRegionId() *string {
	return s.UdmRegionId
}

func (s *DescribeUdmSnapshotsShrinkRequest) SetDiskId(v string) *DescribeUdmSnapshotsShrinkRequest {
	s.DiskId = &v
	return s
}

func (s *DescribeUdmSnapshotsShrinkRequest) SetEndTime(v int64) *DescribeUdmSnapshotsShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeUdmSnapshotsShrinkRequest) SetInstanceId(v string) *DescribeUdmSnapshotsShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeUdmSnapshotsShrinkRequest) SetJobId(v string) *DescribeUdmSnapshotsShrinkRequest {
	s.JobId = &v
	return s
}

func (s *DescribeUdmSnapshotsShrinkRequest) SetSnapshotIdsShrink(v string) *DescribeUdmSnapshotsShrinkRequest {
	s.SnapshotIdsShrink = &v
	return s
}

func (s *DescribeUdmSnapshotsShrinkRequest) SetSourceType(v string) *DescribeUdmSnapshotsShrinkRequest {
	s.SourceType = &v
	return s
}

func (s *DescribeUdmSnapshotsShrinkRequest) SetStartTime(v int64) *DescribeUdmSnapshotsShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeUdmSnapshotsShrinkRequest) SetUdmRegionId(v string) *DescribeUdmSnapshotsShrinkRequest {
	s.UdmRegionId = &v
	return s
}

func (s *DescribeUdmSnapshotsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
