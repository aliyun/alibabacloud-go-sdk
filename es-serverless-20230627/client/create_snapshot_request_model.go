// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSnapshotRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIndices(v string) *CreateSnapshotRequest
	GetIndices() *string
	SetSnapshot(v string) *CreateSnapshotRequest
	GetSnapshot() *string
	SetDryRun(v bool) *CreateSnapshotRequest
	GetDryRun() *bool
}

type CreateSnapshotRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// product_info
	Indices *string `json:"indices,omitempty" xml:"indices,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// qingning
	Snapshot *string `json:"snapshot,omitempty" xml:"snapshot,omitempty"`
	// example:
	//
	// false
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s CreateSnapshotRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSnapshotRequest) GoString() string {
	return s.String()
}

func (s *CreateSnapshotRequest) GetIndices() *string {
	return s.Indices
}

func (s *CreateSnapshotRequest) GetSnapshot() *string {
	return s.Snapshot
}

func (s *CreateSnapshotRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateSnapshotRequest) SetIndices(v string) *CreateSnapshotRequest {
	s.Indices = &v
	return s
}

func (s *CreateSnapshotRequest) SetSnapshot(v string) *CreateSnapshotRequest {
	s.Snapshot = &v
	return s
}

func (s *CreateSnapshotRequest) SetDryRun(v bool) *CreateSnapshotRequest {
	s.DryRun = &v
	return s
}

func (s *CreateSnapshotRequest) Validate() error {
	return dara.Validate(s)
}
