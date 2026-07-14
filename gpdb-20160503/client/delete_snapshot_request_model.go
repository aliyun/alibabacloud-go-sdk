// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSnapshotRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLsn(v string) *DeleteSnapshotRequest
	GetLsn() *string
	SetProjectId(v string) *DeleteSnapshotRequest
	GetProjectId() *string
	SetRegionId(v string) *DeleteSnapshotRequest
	GetRegionId() *string
}

type DeleteSnapshotRequest struct {
	// The LSN of the snapshot to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0/3522648
	Lsn *string `json:"Lsn,omitempty" xml:"Lsn,omitempty"`
	// The Supabase project ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// spb-xxxx
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The region ID. Specifies the region in which to query or perform the operation.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteSnapshotRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSnapshotRequest) GoString() string {
	return s.String()
}

func (s *DeleteSnapshotRequest) GetLsn() *string {
	return s.Lsn
}

func (s *DeleteSnapshotRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *DeleteSnapshotRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteSnapshotRequest) SetLsn(v string) *DeleteSnapshotRequest {
	s.Lsn = &v
	return s
}

func (s *DeleteSnapshotRequest) SetProjectId(v string) *DeleteSnapshotRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteSnapshotRequest) SetRegionId(v string) *DeleteSnapshotRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteSnapshotRequest) Validate() error {
	return dara.Validate(s)
}
