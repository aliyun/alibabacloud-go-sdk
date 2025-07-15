// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetSnapshotRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *ResetSnapshotRequest
	GetRegionId() *string
	SetSnapshotId(v string) *ResetSnapshotRequest
	GetSnapshotId() *string
}

type ResetSnapshotRequest struct {
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hzngahou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the snapshot.
	//
	// This parameter is required.
	//
	// example:
	//
	// s-2zeipxmnhej803x7****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s ResetSnapshotRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetSnapshotRequest) GoString() string {
	return s.String()
}

func (s *ResetSnapshotRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ResetSnapshotRequest) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *ResetSnapshotRequest) SetRegionId(v string) *ResetSnapshotRequest {
	s.RegionId = &v
	return s
}

func (s *ResetSnapshotRequest) SetSnapshotId(v string) *ResetSnapshotRequest {
	s.SnapshotId = &v
	return s
}

func (s *ResetSnapshotRequest) Validate() error {
	return dara.Validate(s)
}
