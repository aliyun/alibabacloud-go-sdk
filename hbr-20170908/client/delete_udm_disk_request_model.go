// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUdmDiskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDiskId(v string) *DeleteUdmDiskRequest
	GetDiskId() *string
}

type DeleteUdmDiskRequest struct {
	// The disk ID.
	//
	// example:
	//
	// d-bp15************xy70
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
}

func (s DeleteUdmDiskRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteUdmDiskRequest) GoString() string {
	return s.String()
}

func (s *DeleteUdmDiskRequest) GetDiskId() *string {
	return s.DiskId
}

func (s *DeleteUdmDiskRequest) SetDiskId(v string) *DeleteUdmDiskRequest {
	s.DiskId = &v
	return s
}

func (s *DeleteUdmDiskRequest) Validate() error {
	return dara.Validate(s)
}
