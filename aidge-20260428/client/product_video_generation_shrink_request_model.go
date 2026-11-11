// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProductVideoGenerationShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAspectRatio(v string) *ProductVideoGenerationShrinkRequest
	GetAspectRatio() *string
	SetDuration(v int64) *ProductVideoGenerationShrinkRequest
	GetDuration() *int64
	SetProductImagesShrink(v string) *ProductVideoGenerationShrinkRequest
	GetProductImagesShrink() *string
	SetProductTitle(v string) *ProductVideoGenerationShrinkRequest
	GetProductTitle() *string
	SetResolution(v string) *ProductVideoGenerationShrinkRequest
	GetResolution() *string
	SetScript(v string) *ProductVideoGenerationShrinkRequest
	GetScript() *string
}

type ProductVideoGenerationShrinkRequest struct {
	// The aspect ratio. Default value: 9:16. Valid values: 16:9, 1:1, 3:4, and 9:16.
	//
	// example:
	//
	// 9:16
	AspectRatio *string `json:"aspect_ratio,omitempty" xml:"aspect_ratio,omitempty"`
	// The video duration in seconds. Default value: 15. If no script is provided, the value must be an integer from 5 to 15. If a script is provided, the value must be an integer from 5 to 30. A value of 15.0 is treated as 15.
	//
	// example:
	//
	// 15
	Duration *int64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// The public image URLs. At least one image URL is required. Each URL must start with http:// or https:// after leading and trailing whitespaces are removed.
	//
	// This parameter is required.
	ProductImagesShrink *string `json:"product_images,omitempty" xml:"product_images,omitempty"`
	// The product title. This parameter is required.
	//
	// This parameter is required.
	//
	// example:
	//
	// White minimalist ceramic mug
	ProductTitle *string `json:"product_title,omitempty" xml:"product_title,omitempty"`
	// The resolution. Default value: 720P. Valid values: 720P and 1080P. The value is case-insensitive and is automatically converted to uppercase by the service.
	//
	// example:
	//
	// 720P
	Resolution *string `json:"resolution,omitempty" xml:"resolution,omitempty"`
	// The existing storyboard script. If this parameter is not provided or is left blank, automatic planning is used.
	//
	// example:
	//
	// create a dog
	Script *string `json:"script,omitempty" xml:"script,omitempty"`
}

func (s ProductVideoGenerationShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ProductVideoGenerationShrinkRequest) GoString() string {
	return s.String()
}

func (s *ProductVideoGenerationShrinkRequest) GetAspectRatio() *string {
	return s.AspectRatio
}

func (s *ProductVideoGenerationShrinkRequest) GetDuration() *int64 {
	return s.Duration
}

func (s *ProductVideoGenerationShrinkRequest) GetProductImagesShrink() *string {
	return s.ProductImagesShrink
}

func (s *ProductVideoGenerationShrinkRequest) GetProductTitle() *string {
	return s.ProductTitle
}

func (s *ProductVideoGenerationShrinkRequest) GetResolution() *string {
	return s.Resolution
}

func (s *ProductVideoGenerationShrinkRequest) GetScript() *string {
	return s.Script
}

func (s *ProductVideoGenerationShrinkRequest) SetAspectRatio(v string) *ProductVideoGenerationShrinkRequest {
	s.AspectRatio = &v
	return s
}

func (s *ProductVideoGenerationShrinkRequest) SetDuration(v int64) *ProductVideoGenerationShrinkRequest {
	s.Duration = &v
	return s
}

func (s *ProductVideoGenerationShrinkRequest) SetProductImagesShrink(v string) *ProductVideoGenerationShrinkRequest {
	s.ProductImagesShrink = &v
	return s
}

func (s *ProductVideoGenerationShrinkRequest) SetProductTitle(v string) *ProductVideoGenerationShrinkRequest {
	s.ProductTitle = &v
	return s
}

func (s *ProductVideoGenerationShrinkRequest) SetResolution(v string) *ProductVideoGenerationShrinkRequest {
	s.Resolution = &v
	return s
}

func (s *ProductVideoGenerationShrinkRequest) SetScript(v string) *ProductVideoGenerationShrinkRequest {
	s.Script = &v
	return s
}

func (s *ProductVideoGenerationShrinkRequest) Validate() error {
	return dara.Validate(s)
}
