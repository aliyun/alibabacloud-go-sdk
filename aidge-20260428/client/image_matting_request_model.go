// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageMattingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackGroundType(v string) *ImageMattingRequest
	GetBackGroundType() *string
	SetBgColor(v string) *ImageMattingRequest
	GetBgColor() *string
	SetImageUrl(v string) *ImageMattingRequest
	GetImageUrl() *string
	SetTargetHeight(v int32) *ImageMattingRequest
	GetTargetHeight() *int32
	SetTargetWidth(v int32) *ImageMattingRequest
	GetTargetWidth() *int32
}

type ImageMattingRequest struct {
	// The URL of the image to process.
	//
	// This parameter is required.
	//
	// example:
	//
	// WHITE_BACKGROUND
	BackGroundType *string `json:"BackGroundType,omitempty" xml:"BackGroundType,omitempty"`
	// The target image height in pixels.
	//
	// example:
	//
	// 255,255,255
	BgColor *string `json:"BgColor,omitempty" xml:"BgColor,omitempty"`
	// The URL of the original image. The image must be in JPG, JPEG, PNG, BMP, or WEBP format. The resolution must be between 256 × 256 and 3000 × 3000 pixels. The file size cannot exceed 10 MB.<br>**Example**: `"https://ae01.alicdn.com/kf/S342f0070dc9f4be09a6cbed34e90dc8fs.jpg"`.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://sc02.alicdn.com/kf/H09364d2c7fc942e685d3b0f656261b24Q.png
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The target image width in pixels.
	//
	// example:
	//
	// 800
	TargetHeight *int32 `json:"TargetHeight,omitempty" xml:"TargetHeight,omitempty"`
	// The background type of the returned image. Valid values:
	//
	// - WHITE_BACKGROUND: white background.
	//
	// - TRANSPARENT: transparent background.
	//
	// example:
	//
	// 800
	TargetWidth *int32 `json:"TargetWidth,omitempty" xml:"TargetWidth,omitempty"`
}

func (s ImageMattingRequest) String() string {
	return dara.Prettify(s)
}

func (s ImageMattingRequest) GoString() string {
	return s.String()
}

func (s *ImageMattingRequest) GetBackGroundType() *string {
	return s.BackGroundType
}

func (s *ImageMattingRequest) GetBgColor() *string {
	return s.BgColor
}

func (s *ImageMattingRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *ImageMattingRequest) GetTargetHeight() *int32 {
	return s.TargetHeight
}

func (s *ImageMattingRequest) GetTargetWidth() *int32 {
	return s.TargetWidth
}

func (s *ImageMattingRequest) SetBackGroundType(v string) *ImageMattingRequest {
	s.BackGroundType = &v
	return s
}

func (s *ImageMattingRequest) SetBgColor(v string) *ImageMattingRequest {
	s.BgColor = &v
	return s
}

func (s *ImageMattingRequest) SetImageUrl(v string) *ImageMattingRequest {
	s.ImageUrl = &v
	return s
}

func (s *ImageMattingRequest) SetTargetHeight(v int32) *ImageMattingRequest {
	s.TargetHeight = &v
	return s
}

func (s *ImageMattingRequest) SetTargetWidth(v int32) *ImageMattingRequest {
	s.TargetWidth = &v
	return s
}

func (s *ImageMattingRequest) Validate() error {
	return dara.Validate(s)
}
