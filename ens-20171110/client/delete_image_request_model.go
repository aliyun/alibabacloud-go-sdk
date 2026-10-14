// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageId(v string) *DeleteImageRequest
	GetImageId() *string
}

type DeleteImageRequest struct {
	// The image ID. Only a single ID is supported.
	//
	// Only custom images are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// m-5taesrgwpo9zqj9cjqu792****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
}

func (s DeleteImageRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteImageRequest) GoString() string {
	return s.String()
}

func (s *DeleteImageRequest) GetImageId() *string {
	return s.ImageId
}

func (s *DeleteImageRequest) SetImageId(v string) *DeleteImageRequest {
	s.ImageId = &v
	return s
}

func (s *DeleteImageRequest) Validate() error {
	return dara.Validate(s)
}
