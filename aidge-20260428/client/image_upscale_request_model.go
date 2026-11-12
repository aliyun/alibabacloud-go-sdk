// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageUpscaleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageUrl(v string) *ImageUpscaleRequest
	GetImageUrl() *string
	SetUpscaleFactor(v int32) *ImageUpscaleRequest
	GetUpscaleFactor() *int32
}

type ImageUpscaleRequest struct {
	// The URL of the original image. It must be a publicly accessible image URL. Private network addresses such as localhost are not supported. The width and height of the input image must be greater than 100 pixels and less than 3,000 pixels. The width and height of the upscaled image must not exceed 3,000 pixels.
	//
	// This parameter is required.
	//
	// example:
	//
	// http://example.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The upscaling factor of the image. Valid values: 2 and 4. Default value: 2.
	//
	// example:
	//
	// 2
	UpscaleFactor *int32 `json:"UpscaleFactor,omitempty" xml:"UpscaleFactor,omitempty"`
}

func (s ImageUpscaleRequest) String() string {
	return dara.Prettify(s)
}

func (s ImageUpscaleRequest) GoString() string {
	return s.String()
}

func (s *ImageUpscaleRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *ImageUpscaleRequest) GetUpscaleFactor() *int32 {
	return s.UpscaleFactor
}

func (s *ImageUpscaleRequest) SetImageUrl(v string) *ImageUpscaleRequest {
	s.ImageUrl = &v
	return s
}

func (s *ImageUpscaleRequest) SetUpscaleFactor(v int32) *ImageUpscaleRequest {
	s.UpscaleFactor = &v
	return s
}

func (s *ImageUpscaleRequest) Validate() error {
	return dara.Validate(s)
}
