// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelCopyImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageId(v string) *CancelCopyImageRequest
	GetImageId() *string
	SetRegionId(v string) *CancelCopyImageRequest
	GetRegionId() *string
}

type CancelCopyImageRequest struct {
	// The ID of the new image in the destination region.
	//
	// This parameter is required.
	//
	// example:
	//
	// m-gx2x1dhsmusr2****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The ID of the region to which the image is copied.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CancelCopyImageRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelCopyImageRequest) GoString() string {
	return s.String()
}

func (s *CancelCopyImageRequest) GetImageId() *string {
	return s.ImageId
}

func (s *CancelCopyImageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CancelCopyImageRequest) SetImageId(v string) *CancelCopyImageRequest {
	s.ImageId = &v
	return s
}

func (s *CancelCopyImageRequest) SetRegionId(v string) *CancelCopyImageRequest {
	s.RegionId = &v
	return s
}

func (s *CancelCopyImageRequest) Validate() error {
	return dara.Validate(s)
}
