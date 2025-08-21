// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDiskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDiskId(v string) *DeleteDiskRequest
	GetDiskId() *string
}

type DeleteDiskRequest struct {
	// The ID of the disk.
	//
	// This parameter is required.
	//
	// example:
	//
	// d-5va95bg6i5f44kgkeuazyfcxm
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
}

func (s DeleteDiskRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDiskRequest) GoString() string {
	return s.String()
}

func (s *DeleteDiskRequest) GetDiskId() *string {
	return s.DiskId
}

func (s *DeleteDiskRequest) SetDiskId(v string) *DeleteDiskRequest {
	s.DiskId = &v
	return s
}

func (s *DeleteDiskRequest) Validate() error {
	return dara.Validate(s)
}
