// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProductVideoGenerationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAspectRatio(v string) *ProductVideoGenerationRequest
	GetAspectRatio() *string
	SetDuration(v int64) *ProductVideoGenerationRequest
	GetDuration() *int64
	SetProductImages(v []*string) *ProductVideoGenerationRequest
	GetProductImages() []*string
	SetProductTitle(v string) *ProductVideoGenerationRequest
	GetProductTitle() *string
	SetResolution(v string) *ProductVideoGenerationRequest
	GetResolution() *string
	SetScript(v string) *ProductVideoGenerationRequest
	GetScript() *string
}

type ProductVideoGenerationRequest struct {
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
	ProductImages []*string `json:"product_images,omitempty" xml:"product_images,omitempty" type:"Repeated"`
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

func (s ProductVideoGenerationRequest) String() string {
	return dara.Prettify(s)
}

func (s ProductVideoGenerationRequest) GoString() string {
	return s.String()
}

func (s *ProductVideoGenerationRequest) GetAspectRatio() *string {
	return s.AspectRatio
}

func (s *ProductVideoGenerationRequest) GetDuration() *int64 {
	return s.Duration
}

func (s *ProductVideoGenerationRequest) GetProductImages() []*string {
	return s.ProductImages
}

func (s *ProductVideoGenerationRequest) GetProductTitle() *string {
	return s.ProductTitle
}

func (s *ProductVideoGenerationRequest) GetResolution() *string {
	return s.Resolution
}

func (s *ProductVideoGenerationRequest) GetScript() *string {
	return s.Script
}

func (s *ProductVideoGenerationRequest) SetAspectRatio(v string) *ProductVideoGenerationRequest {
	s.AspectRatio = &v
	return s
}

func (s *ProductVideoGenerationRequest) SetDuration(v int64) *ProductVideoGenerationRequest {
	s.Duration = &v
	return s
}

func (s *ProductVideoGenerationRequest) SetProductImages(v []*string) *ProductVideoGenerationRequest {
	s.ProductImages = v
	return s
}

func (s *ProductVideoGenerationRequest) SetProductTitle(v string) *ProductVideoGenerationRequest {
	s.ProductTitle = &v
	return s
}

func (s *ProductVideoGenerationRequest) SetResolution(v string) *ProductVideoGenerationRequest {
	s.Resolution = &v
	return s
}

func (s *ProductVideoGenerationRequest) SetScript(v string) *ProductVideoGenerationRequest {
	s.Script = &v
	return s
}

func (s *ProductVideoGenerationRequest) Validate() error {
	return dara.Validate(s)
}
