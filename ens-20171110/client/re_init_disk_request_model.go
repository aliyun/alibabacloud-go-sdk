// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReInitDiskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDiskId(v string) *ReInitDiskRequest
	GetDiskId() *string
	SetImageId(v string) *ReInitDiskRequest
	GetImageId() *string
}

type ReInitDiskRequest struct {
	// The ID of the disk to be initialized. You can initialize only one disk at a time.
	//
	// This parameter is required.
	//
	// example:
	//
	// d-5r7v69e0bejrnzger09w71yjv
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The ID of the image to use to create the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// m-5rz3i231o531s4p4ozanxmgx7
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
}

func (s ReInitDiskRequest) String() string {
	return dara.Prettify(s)
}

func (s ReInitDiskRequest) GoString() string {
	return s.String()
}

func (s *ReInitDiskRequest) GetDiskId() *string {
	return s.DiskId
}

func (s *ReInitDiskRequest) GetImageId() *string {
	return s.ImageId
}

func (s *ReInitDiskRequest) SetDiskId(v string) *ReInitDiskRequest {
	s.DiskId = &v
	return s
}

func (s *ReInitDiskRequest) SetImageId(v string) *ReInitDiskRequest {
	s.ImageId = &v
	return s
}

func (s *ReInitDiskRequest) Validate() error {
	return dara.Validate(s)
}
