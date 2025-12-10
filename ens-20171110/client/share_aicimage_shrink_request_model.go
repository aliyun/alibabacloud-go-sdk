// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iShareAICImageShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageId(v string) *ShareAICImageShrinkRequest
	GetImageId() *string
	SetUsersShrink(v string) *ShareAICImageShrinkRequest
	GetUsersShrink() *string
}

type ShareAICImageShrinkRequest struct {
	// The image name.
	//
	// This parameter is required.
	//
	// example:
	//
	// mykey
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The user groups.
	//
	// This parameter is required.
	UsersShrink *string `json:"Users,omitempty" xml:"Users,omitempty"`
}

func (s ShareAICImageShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ShareAICImageShrinkRequest) GoString() string {
	return s.String()
}

func (s *ShareAICImageShrinkRequest) GetImageId() *string {
	return s.ImageId
}

func (s *ShareAICImageShrinkRequest) GetUsersShrink() *string {
	return s.UsersShrink
}

func (s *ShareAICImageShrinkRequest) SetImageId(v string) *ShareAICImageShrinkRequest {
	s.ImageId = &v
	return s
}

func (s *ShareAICImageShrinkRequest) SetUsersShrink(v string) *ShareAICImageShrinkRequest {
	s.UsersShrink = &v
	return s
}

func (s *ShareAICImageShrinkRequest) Validate() error {
	return dara.Validate(s)
}
