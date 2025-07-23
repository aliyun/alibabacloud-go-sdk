// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageId(v string) *GetImageRequest
	GetImageId() *string
}

type GetImageRequest struct {
	// The ID of the image.
	//
	// This parameter is required.
	//
	// example:
	//
	// image-wz9c5ths5dfuwx47****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
}

func (s GetImageRequest) String() string {
	return dara.Prettify(s)
}

func (s GetImageRequest) GoString() string {
	return s.String()
}

func (s *GetImageRequest) GetImageId() *string {
	return s.ImageId
}

func (s *GetImageRequest) SetImageId(v string) *GetImageRequest {
	s.ImageId = &v
	return s
}

func (s *GetImageRequest) Validate() error {
	return dara.Validate(s)
}
