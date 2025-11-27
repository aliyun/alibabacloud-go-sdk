// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageId(v string) *RemoveImageRequest
	GetImageId() *string
	SetImageType(v string) *RemoveImageRequest
	GetImageType() *string
}

type RemoveImageRequest struct {
	// The image ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// m-bp14wakr1rkxtb******
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The type of the images. Valid values:
	//
	// 	- VM: Virtual Machine Image
	//
	// 	- Container: container image
	//
	// example:
	//
	// VM
	ImageType *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
}

func (s RemoveImageRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveImageRequest) GoString() string {
	return s.String()
}

func (s *RemoveImageRequest) GetImageId() *string {
	return s.ImageId
}

func (s *RemoveImageRequest) GetImageType() *string {
	return s.ImageType
}

func (s *RemoveImageRequest) SetImageId(v string) *RemoveImageRequest {
	s.ImageId = &v
	return s
}

func (s *RemoveImageRequest) SetImageType(v string) *RemoveImageRequest {
	s.ImageType = &v
	return s
}

func (s *RemoveImageRequest) Validate() error {
	return dara.Validate(s)
}
