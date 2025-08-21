// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSnapshotRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSnapshotId(v string) *DeleteSnapshotRequest
	GetSnapshotId() *string
}

type DeleteSnapshotRequest struct {
	// The ID of the snapshot.
	//
	// This parameter is required.
	//
	// example:
	//
	// sp-bp1c0doj0taqyzzl****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s DeleteSnapshotRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSnapshotRequest) GoString() string {
	return s.String()
}

func (s *DeleteSnapshotRequest) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *DeleteSnapshotRequest) SetSnapshotId(v string) *DeleteSnapshotRequest {
	s.SnapshotId = &v
	return s
}

func (s *DeleteSnapshotRequest) Validate() error {
	return dara.Validate(s)
}
