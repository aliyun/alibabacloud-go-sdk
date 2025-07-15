// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteImagesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageIdsShrink(v string) *DeleteImagesShrinkRequest
	GetImageIdsShrink() *string
}

type DeleteImagesShrinkRequest struct {
	// The IDs of the images.
	//
	// This parameter is required.
	ImageIdsShrink *string `json:"ImageIds,omitempty" xml:"ImageIds,omitempty"`
}

func (s DeleteImagesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteImagesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteImagesShrinkRequest) GetImageIdsShrink() *string {
	return s.ImageIdsShrink
}

func (s *DeleteImagesShrinkRequest) SetImageIdsShrink(v string) *DeleteImagesShrinkRequest {
	s.ImageIdsShrink = &v
	return s
}

func (s *DeleteImagesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
