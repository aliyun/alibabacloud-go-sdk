// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSnapshotRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DeleteSnapshotRequest
	GetRegionId() *string
	SetSnapshotId(v []*string) *DeleteSnapshotRequest
	GetSnapshotId() []*string
}

type DeleteSnapshotRequest struct {
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The snapshot IDs. You can specify 1 to 100 IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// s-2ze81owrnv9pity4****
	SnapshotId []*string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty" type:"Repeated"`
}

func (s DeleteSnapshotRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSnapshotRequest) GoString() string {
	return s.String()
}

func (s *DeleteSnapshotRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteSnapshotRequest) GetSnapshotId() []*string {
	return s.SnapshotId
}

func (s *DeleteSnapshotRequest) SetRegionId(v string) *DeleteSnapshotRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteSnapshotRequest) SetSnapshotId(v []*string) *DeleteSnapshotRequest {
	s.SnapshotId = v
	return s
}

func (s *DeleteSnapshotRequest) Validate() error {
	return dara.Validate(s)
}
