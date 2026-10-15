// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageMattingAsyncRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackGroundType(v string) *ImageMattingAsyncRequest
	GetBackGroundType() *string
	SetBgColor(v string) *ImageMattingAsyncRequest
	GetBgColor() *string
	SetImageUrl(v string) *ImageMattingAsyncRequest
	GetImageUrl() *string
	SetTargetHeight(v int32) *ImageMattingAsyncRequest
	GetTargetHeight() *int32
	SetTargetWidth(v int32) *ImageMattingAsyncRequest
	GetTargetWidth() *int32
}

type ImageMattingAsyncRequest struct {
	// Specifies the background format of the returned image. Valid values:
	//
	// - WHITE_BACKGROUND: Returns an image with a white background.
	//
	// - TRANSPARENT: Returns an image with a transparent background.
	//
	// This parameter is required.
	//
	// example:
	//
	// WHITE_BACKGROUND
	BackGroundType *string `json:"BackGroundType,omitempty" xml:"BackGroundType,omitempty"`
	// The BGR color value of the background. Default value: white.
	//
	// example:
	//
	// 255,255,255
	BgColor *string `json:"BgColor,omitempty" xml:"BgColor,omitempty"`
	// The URL of the original image. The image must be in JPG, JPEG, PNG, BMP, or WEBP format. The resolution must be between 100 × 100 and 3000 × 3000 pixels, and the file size cannot exceed 10 MB.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://ae01.alicdn.com/kf/S342f0070dc9f4be09a6cbed34e90dc8fs.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The expected height of the cropped image. If you do not specify this parameter, the original height of the image is used. Unit: pixels. Valid values: 100 to 5000.
	//
	// example:
	//
	// 800
	TargetHeight *int32 `json:"TargetHeight,omitempty" xml:"TargetHeight,omitempty"`
	// The expected width of the cropped image. If you do not specify this parameter, the original width of the image is used. Unit: pixels. Valid values: 100 to 5000.
	//
	// example:
	//
	// 800
	TargetWidth *int32 `json:"TargetWidth,omitempty" xml:"TargetWidth,omitempty"`
}

func (s ImageMattingAsyncRequest) String() string {
	return dara.Prettify(s)
}

func (s ImageMattingAsyncRequest) GoString() string {
	return s.String()
}

func (s *ImageMattingAsyncRequest) GetBackGroundType() *string {
	return s.BackGroundType
}

func (s *ImageMattingAsyncRequest) GetBgColor() *string {
	return s.BgColor
}

func (s *ImageMattingAsyncRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *ImageMattingAsyncRequest) GetTargetHeight() *int32 {
	return s.TargetHeight
}

func (s *ImageMattingAsyncRequest) GetTargetWidth() *int32 {
	return s.TargetWidth
}

func (s *ImageMattingAsyncRequest) SetBackGroundType(v string) *ImageMattingAsyncRequest {
	s.BackGroundType = &v
	return s
}

func (s *ImageMattingAsyncRequest) SetBgColor(v string) *ImageMattingAsyncRequest {
	s.BgColor = &v
	return s
}

func (s *ImageMattingAsyncRequest) SetImageUrl(v string) *ImageMattingAsyncRequest {
	s.ImageUrl = &v
	return s
}

func (s *ImageMattingAsyncRequest) SetTargetHeight(v int32) *ImageMattingAsyncRequest {
	s.TargetHeight = &v
	return s
}

func (s *ImageMattingAsyncRequest) SetTargetWidth(v int32) *ImageMattingAsyncRequest {
	s.TargetWidth = &v
	return s
}

func (s *ImageMattingAsyncRequest) Validate() error {
	return dara.Validate(s)
}
