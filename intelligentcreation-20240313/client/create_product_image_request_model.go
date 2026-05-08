// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProductImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackgroundDescription(v string) *CreateProductImageRequest
	GetBackgroundDescription() *string
	SetBackgroundPriority(v int32) *CreateProductImageRequest
	GetBackgroundPriority() *int32
	SetBackgroundUrl(v string) *CreateProductImageRequest
	GetBackgroundUrl() *string
	SetHighlightText(v string) *CreateProductImageRequest
	GetHighlightText() *string
	SetImageCount(v int32) *CreateProductImageRequest
	GetImageCount() *int32
	SetImageUrl(v string) *CreateProductImageRequest
	GetImageUrl() *string
	SetSubTitle(v string) *CreateProductImageRequest
	GetSubTitle() *string
	SetTitle(v string) *CreateProductImageRequest
	GetTitle() *string
}

type CreateProductImageRequest struct {
	BackgroundDescription *string `json:"backgroundDescription,omitempty" xml:"backgroundDescription,omitempty"`
	// example:
	//
	// 75
	BackgroundPriority *int32 `json:"backgroundPriority,omitempty" xml:"backgroundPriority,omitempty"`
	// example:
	//
	// http://xxx/background.png
	BackgroundUrl *string `json:"backgroundUrl,omitempty" xml:"backgroundUrl,omitempty"`
	HighlightText *string `json:"highlightText,omitempty" xml:"highlightText,omitempty"`
	// example:
	//
	// 1
	ImageCount *int32 `json:"imageCount,omitempty" xml:"imageCount,omitempty"`
	// example:
	//
	// http://xxx/image.png
	ImageUrl *string `json:"imageUrl,omitempty" xml:"imageUrl,omitempty"`
	SubTitle *string `json:"subTitle,omitempty" xml:"subTitle,omitempty"`
	Title    *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s CreateProductImageRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateProductImageRequest) GoString() string {
	return s.String()
}

func (s *CreateProductImageRequest) GetBackgroundDescription() *string {
	return s.BackgroundDescription
}

func (s *CreateProductImageRequest) GetBackgroundPriority() *int32 {
	return s.BackgroundPriority
}

func (s *CreateProductImageRequest) GetBackgroundUrl() *string {
	return s.BackgroundUrl
}

func (s *CreateProductImageRequest) GetHighlightText() *string {
	return s.HighlightText
}

func (s *CreateProductImageRequest) GetImageCount() *int32 {
	return s.ImageCount
}

func (s *CreateProductImageRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *CreateProductImageRequest) GetSubTitle() *string {
	return s.SubTitle
}

func (s *CreateProductImageRequest) GetTitle() *string {
	return s.Title
}

func (s *CreateProductImageRequest) SetBackgroundDescription(v string) *CreateProductImageRequest {
	s.BackgroundDescription = &v
	return s
}

func (s *CreateProductImageRequest) SetBackgroundPriority(v int32) *CreateProductImageRequest {
	s.BackgroundPriority = &v
	return s
}

func (s *CreateProductImageRequest) SetBackgroundUrl(v string) *CreateProductImageRequest {
	s.BackgroundUrl = &v
	return s
}

func (s *CreateProductImageRequest) SetHighlightText(v string) *CreateProductImageRequest {
	s.HighlightText = &v
	return s
}

func (s *CreateProductImageRequest) SetImageCount(v int32) *CreateProductImageRequest {
	s.ImageCount = &v
	return s
}

func (s *CreateProductImageRequest) SetImageUrl(v string) *CreateProductImageRequest {
	s.ImageUrl = &v
	return s
}

func (s *CreateProductImageRequest) SetSubTitle(v string) *CreateProductImageRequest {
	s.SubTitle = &v
	return s
}

func (s *CreateProductImageRequest) SetTitle(v string) *CreateProductImageRequest {
	s.Title = &v
	return s
}

func (s *CreateProductImageRequest) Validate() error {
	return dara.Validate(s)
}
