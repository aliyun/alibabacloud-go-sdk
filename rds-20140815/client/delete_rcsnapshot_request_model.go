// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRCSnapshotRequest interface {
	dara.Model
	String() string
	GoString() string
	SetForce(v bool) *DeleteRCSnapshotRequest
	GetForce() *bool
	SetRegionId(v string) *DeleteRCSnapshotRequest
	GetRegionId() *string
	SetSnapshotId(v string) *DeleteRCSnapshotRequest
	GetSnapshotId() *string
}

type DeleteRCSnapshotRequest struct {
	// Specifies whether to forcefully delete the snapshot that is used to create cloud disks. Valid values:
	//
	// 	- **true**: forcefully deletes the snapshot After the snapshot is forcefully deleted, the cloud disks created from the snapshot cannot be re-initialized.
	//
	// 	- **false**(default): does not forcefully delete the snapshot.
	//
	// example:
	//
	// false
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The snapshot ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rcds-7mbefjzkqccvdev****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s DeleteRCSnapshotRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRCSnapshotRequest) GoString() string {
	return s.String()
}

func (s *DeleteRCSnapshotRequest) GetForce() *bool {
	return s.Force
}

func (s *DeleteRCSnapshotRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteRCSnapshotRequest) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *DeleteRCSnapshotRequest) SetForce(v bool) *DeleteRCSnapshotRequest {
	s.Force = &v
	return s
}

func (s *DeleteRCSnapshotRequest) SetRegionId(v string) *DeleteRCSnapshotRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteRCSnapshotRequest) SetSnapshotId(v string) *DeleteRCSnapshotRequest {
	s.SnapshotId = &v
	return s
}

func (s *DeleteRCSnapshotRequest) Validate() error {
	return dara.Validate(s)
}
