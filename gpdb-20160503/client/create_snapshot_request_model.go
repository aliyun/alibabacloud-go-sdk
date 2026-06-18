// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSnapshotRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateSnapshotRequest
	GetClientToken() *string
	SetLsn(v string) *CreateSnapshotRequest
	GetLsn() *string
	SetProjectId(v string) *CreateSnapshotRequest
	GetProjectId() *string
	SetRegionId(v string) *CreateSnapshotRequest
	GetRegionId() *string
	SetSnapshotTimestamp(v string) *CreateSnapshotRequest
	GetSnapshotTimestamp() *string
}

type CreateSnapshotRequest struct {
	// The idempotence token. Ensures that repeated requests do not result in duplicate operations.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The LSN for the snapshot. You must specify either this parameter or SnapshotTimestamp. If this parameter is specified, the snapshot is created based on the specified LSN.
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
	// The region ID. Specifies the region in which to perform the operation.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The point in time for the snapshot. You must specify either this parameter or Lsn. If this parameter is specified, the snapshot is created based on the specified point in time.
	//
	// example:
	//
	// 2026-04-08T09:11:12Z
	SnapshotTimestamp *string `json:"SnapshotTimestamp,omitempty" xml:"SnapshotTimestamp,omitempty"`
}

func (s CreateSnapshotRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSnapshotRequest) GoString() string {
	return s.String()
}

func (s *CreateSnapshotRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateSnapshotRequest) GetLsn() *string {
	return s.Lsn
}

func (s *CreateSnapshotRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *CreateSnapshotRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateSnapshotRequest) GetSnapshotTimestamp() *string {
	return s.SnapshotTimestamp
}

func (s *CreateSnapshotRequest) SetClientToken(v string) *CreateSnapshotRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateSnapshotRequest) SetLsn(v string) *CreateSnapshotRequest {
	s.Lsn = &v
	return s
}

func (s *CreateSnapshotRequest) SetProjectId(v string) *CreateSnapshotRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateSnapshotRequest) SetRegionId(v string) *CreateSnapshotRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSnapshotRequest) SetSnapshotTimestamp(v string) *CreateSnapshotRequest {
	s.SnapshotTimestamp = &v
	return s
}

func (s *CreateSnapshotRequest) Validate() error {
	return dara.Validate(s)
}
