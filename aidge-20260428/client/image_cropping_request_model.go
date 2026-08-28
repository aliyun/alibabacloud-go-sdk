// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageCroppingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageUrl(v string) *ImageCroppingRequest
	GetImageUrl() *string
	SetTargetHeight(v int32) *ImageCroppingRequest
	GetTargetHeight() *int32
	SetTargetWidth(v int32) *ImageCroppingRequest
	GetTargetWidth() *int32
}

type ImageCroppingRequest struct {
	// The URL of the image to process.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/image.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The target height.
	//
	// This parameter is required.
	//
	// example:
	//
	// 800
	TargetHeight *int32 `json:"TargetHeight,omitempty" xml:"TargetHeight,omitempty"`
	// The target width.
	//
	// This parameter is required.
	//
	// example:
	//
	// 800
	TargetWidth *int32 `json:"TargetWidth,omitempty" xml:"TargetWidth,omitempty"`
}

func (s ImageCroppingRequest) String() string {
	return dara.Prettify(s)
}

func (s ImageCroppingRequest) GoString() string {
	return s.String()
}

func (s *ImageCroppingRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *ImageCroppingRequest) GetTargetHeight() *int32 {
	return s.TargetHeight
}

func (s *ImageCroppingRequest) GetTargetWidth() *int32 {
	return s.TargetWidth
}

func (s *ImageCroppingRequest) SetImageUrl(v string) *ImageCroppingRequest {
	s.ImageUrl = &v
	return s
}

func (s *ImageCroppingRequest) SetTargetHeight(v int32) *ImageCroppingRequest {
	s.TargetHeight = &v
	return s
}

func (s *ImageCroppingRequest) SetTargetWidth(v int32) *ImageCroppingRequest {
	s.TargetWidth = &v
	return s
}

func (s *ImageCroppingRequest) Validate() error {
	return dara.Validate(s)
}
