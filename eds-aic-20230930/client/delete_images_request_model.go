// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteImagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageIds(v []*string) *DeleteImagesRequest
	GetImageIds() []*string
}

type DeleteImagesRequest struct {
	// The IDs of the images.
	//
	// This parameter is required.
	ImageIds []*string `json:"ImageIds,omitempty" xml:"ImageIds,omitempty" type:"Repeated"`
}

func (s DeleteImagesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteImagesRequest) GoString() string {
	return s.String()
}

func (s *DeleteImagesRequest) GetImageIds() []*string {
	return s.ImageIds
}

func (s *DeleteImagesRequest) SetImageIds(v []*string) *DeleteImagesRequest {
	s.ImageIds = v
	return s
}

func (s *DeleteImagesRequest) Validate() error {
	return dara.Validate(s)
}
