// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUdmSnapshotsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDiskId(v string) *DescribeUdmSnapshotsRequest
	GetDiskId() *string
	SetEndTime(v int64) *DescribeUdmSnapshotsRequest
	GetEndTime() *int64
	SetInstanceId(v string) *DescribeUdmSnapshotsRequest
	GetInstanceId() *string
	SetJobId(v string) *DescribeUdmSnapshotsRequest
	GetJobId() *string
	SetSnapshotIds(v map[string]interface{}) *DescribeUdmSnapshotsRequest
	GetSnapshotIds() map[string]interface{}
	SetSourceType(v string) *DescribeUdmSnapshotsRequest
	GetSourceType() *string
	SetStartTime(v int64) *DescribeUdmSnapshotsRequest
	GetStartTime() *int64
	SetUdmRegionId(v string) *DescribeUdmSnapshotsRequest
	GetUdmRegionId() *string
}

type DescribeUdmSnapshotsRequest struct {
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
	SnapshotIds map[string]interface{} `json:"SnapshotIds,omitempty" xml:"SnapshotIds,omitempty"`
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

func (s DescribeUdmSnapshotsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUdmSnapshotsRequest) GoString() string {
	return s.String()
}

func (s *DescribeUdmSnapshotsRequest) GetDiskId() *string {
	return s.DiskId
}

func (s *DescribeUdmSnapshotsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeUdmSnapshotsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeUdmSnapshotsRequest) GetJobId() *string {
	return s.JobId
}

func (s *DescribeUdmSnapshotsRequest) GetSnapshotIds() map[string]interface{} {
	return s.SnapshotIds
}

func (s *DescribeUdmSnapshotsRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *DescribeUdmSnapshotsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeUdmSnapshotsRequest) GetUdmRegionId() *string {
	return s.UdmRegionId
}

func (s *DescribeUdmSnapshotsRequest) SetDiskId(v string) *DescribeUdmSnapshotsRequest {
	s.DiskId = &v
	return s
}

func (s *DescribeUdmSnapshotsRequest) SetEndTime(v int64) *DescribeUdmSnapshotsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeUdmSnapshotsRequest) SetInstanceId(v string) *DescribeUdmSnapshotsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeUdmSnapshotsRequest) SetJobId(v string) *DescribeUdmSnapshotsRequest {
	s.JobId = &v
	return s
}

func (s *DescribeUdmSnapshotsRequest) SetSnapshotIds(v map[string]interface{}) *DescribeUdmSnapshotsRequest {
	s.SnapshotIds = v
	return s
}

func (s *DescribeUdmSnapshotsRequest) SetSourceType(v string) *DescribeUdmSnapshotsRequest {
	s.SourceType = &v
	return s
}

func (s *DescribeUdmSnapshotsRequest) SetStartTime(v int64) *DescribeUdmSnapshotsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeUdmSnapshotsRequest) SetUdmRegionId(v string) *DescribeUdmSnapshotsRequest {
	s.UdmRegionId = &v
	return s
}

func (s *DescribeUdmSnapshotsRequest) Validate() error {
	return dara.Validate(s)
}
