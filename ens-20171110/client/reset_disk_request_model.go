// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetDiskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDiskId(v string) *ResetDiskRequest
	GetDiskId() *string
	SetSnapshotId(v string) *ResetDiskRequest
	GetSnapshotId() *string
}

type ResetDiskRequest struct {
	// The ID of the disk that you want to roll back.
	//
	// This parameter is required.
	//
	// example:
	//
	// d-bp199lyny9b3****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The ID of the snapshot that you want to use to roll back the disk.
	//
	// This parameter is required.
	//
	// example:
	//
	// sp-bp199lyny9b3****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s ResetDiskRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetDiskRequest) GoString() string {
	return s.String()
}

func (s *ResetDiskRequest) GetDiskId() *string {
	return s.DiskId
}

func (s *ResetDiskRequest) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *ResetDiskRequest) SetDiskId(v string) *ResetDiskRequest {
	s.DiskId = &v
	return s
}

func (s *ResetDiskRequest) SetSnapshotId(v string) *ResetDiskRequest {
	s.SnapshotId = &v
	return s
}

func (s *ResetDiskRequest) Validate() error {
	return dara.Validate(s)
}
