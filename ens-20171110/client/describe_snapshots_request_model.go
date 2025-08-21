// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSnapshotsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDiskId(v string) *DescribeSnapshotsRequest
	GetDiskId() *string
	SetEnsRegionId(v string) *DescribeSnapshotsRequest
	GetEnsRegionId() *string
	SetEnsRegionIds(v string) *DescribeSnapshotsRequest
	GetEnsRegionIds() *string
	SetInstanceId(v string) *DescribeSnapshotsRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *DescribeSnapshotsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSnapshotsRequest
	GetPageSize() *int32
	SetSnapshotId(v string) *DescribeSnapshotsRequest
	GetSnapshotId() *string
	SetSnapshotName(v string) *DescribeSnapshotsRequest
	GetSnapshotName() *string
}

type DescribeSnapshotsRequest struct {
	// The ID of the disk.
	//
	// example:
	//
	// d-bp67acfmxazb4p****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The ID of the ENS node. You can query the node ID by calling the [DescribeEnsRegions](https://help.aliyun.com/document_detail/2637662.html) operation.
	//
	// example:
	//
	// cn-hangzhou
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The node information.
	//
	// example:
	//
	// ["cn-suzhou-telecom","cn-chengdu-telecom"]
	EnsRegionIds *string `json:"EnsRegionIds,omitempty" xml:"EnsRegionIds,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// i-bp67acfmxazb4p****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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
	// The ID of the snapshot. The value can be a JSON array that consists of up to 100 snapshot IDs. Separate the snapshot IDs with commas (,).
	//
	// example:
	//
	// sp-bp67acfmxazb4p****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// The name of the snapshot. The name must be 2 to 128 characters in length. It must start with a letter and cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), and hyphens (-).
	//
	// example:
	//
	// testSnapshotName
	SnapshotName *string `json:"SnapshotName,omitempty" xml:"SnapshotName,omitempty"`
}

func (s DescribeSnapshotsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnapshotsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotsRequest) GetDiskId() *string {
	return s.DiskId
}

func (s *DescribeSnapshotsRequest) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeSnapshotsRequest) GetEnsRegionIds() *string {
	return s.EnsRegionIds
}

func (s *DescribeSnapshotsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeSnapshotsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSnapshotsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSnapshotsRequest) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *DescribeSnapshotsRequest) GetSnapshotName() *string {
	return s.SnapshotName
}

func (s *DescribeSnapshotsRequest) SetDiskId(v string) *DescribeSnapshotsRequest {
	s.DiskId = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetEnsRegionId(v string) *DescribeSnapshotsRequest {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetEnsRegionIds(v string) *DescribeSnapshotsRequest {
	s.EnsRegionIds = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetInstanceId(v string) *DescribeSnapshotsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetPageNumber(v int32) *DescribeSnapshotsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetPageSize(v int32) *DescribeSnapshotsRequest {
	s.PageSize = &v
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

func (s *DescribeSnapshotsRequest) Validate() error {
	return dara.Validate(s)
}
