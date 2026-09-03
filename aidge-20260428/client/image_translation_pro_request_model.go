// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageTranslationProRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAsync(v bool) *ImageTranslationProRequest
	GetAsync() *bool
	SetGlossary(v string) *ImageTranslationProRequest
	GetGlossary() *string
	SetImageUrl(v string) *ImageTranslationProRequest
	GetImageUrl() *string
	SetIncludingProductArea(v bool) *ImageTranslationProRequest
	GetIncludingProductArea() *bool
	SetSourceLanguage(v string) *ImageTranslationProRequest
	GetSourceLanguage() *string
	SetTargetLanguage(v string) *ImageTranslationProRequest
	GetTargetLanguage() *string
	SetTranslatingBrandInTheProduct(v bool) *ImageTranslationProRequest
	GetTranslatingBrandInTheProduct() *bool
	SetUseImageEditor(v bool) *ImageTranslationProRequest
	GetUseImageEditor() *bool
}

type ImageTranslationProRequest struct {
	// Specifies whether to use asynchronous mode. Default value: false (synchronous mode). When set to true, the API immediately returns a TaskId. Use the query translation result API to retrieve the final result.
	//
	// example:
	//
	// true
	Async *bool `json:"Async,omitempty" xml:"Async,omitempty"`
	// The ID of the intervention glossary. This parameter is optional. Create the glossary separately in the console and provide its ID. If the glossary ID is empty, the translation results are not modified.
	//
	// example:
	//
	// glossary_1
	Glossary *string `json:"Glossary,omitempty" xml:"Glossary,omitempty"`
	// The URL of the original image. This parameter is required.
	//
	// Image requirements:
	//
	// - Image URL: Must be publicly accessible.
	//
	// - Format: png, jpeg, jpg, bmp, or webp.
	//
	// - Pixels: The width and height must not exceed 4000 pixels each.
	//
	// - File size: The original file must be 10 MB or smaller.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://img.alicdn.com/imgextra/i3/O1CN01HTDhDi28Fd85ZYs7H_!!6000000007903-0-tps-800-800.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// Specifies whether to translate text on the image subject. This parameter is optional. Default value: false. Setting this parameter to false helps protect information such as embedded product names from being translated.
	//
	// example:
	//
	// false
	IncludingProductArea *bool `json:"IncludingProductArea,omitempty" xml:"IncludingProductArea,omitempty"`
	// The source language code. This parameter is required. For supported language directions, see [Language Direction Mapping Table](https://www.alibabacloud.com/help/en/document_detail/3041883.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// en
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	// The target language code. This parameter is required. For supported language directions, see [Language Direction Mapping Table](https://www.alibabacloud.com/help/en/document_detail/3041883.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ko
	TargetLanguage *string `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
	// Specifies whether to translate brand names on the image. This parameter is optional. Default value: false. Setting this parameter to false helps protect brand name information from being translated.
	//
	// example:
	//
	// false
	TranslatingBrandInTheProduct *bool `json:"TranslatingBrandInTheProduct,omitempty" xml:"TranslatingBrandInTheProduct,omitempty"`
	// Specifies whether to return layout information such as text position, font, and color. When set to true, layer information is returned for secondary editing with an image editor. Default value: false.
	//
	// example:
	//
	// false
	UseImageEditor *bool `json:"UseImageEditor,omitempty" xml:"UseImageEditor,omitempty"`
}

func (s ImageTranslationProRequest) String() string {
	return dara.Prettify(s)
}

func (s ImageTranslationProRequest) GoString() string {
	return s.String()
}

func (s *ImageTranslationProRequest) GetAsync() *bool {
	return s.Async
}

func (s *ImageTranslationProRequest) GetGlossary() *string {
	return s.Glossary
}

func (s *ImageTranslationProRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *ImageTranslationProRequest) GetIncludingProductArea() *bool {
	return s.IncludingProductArea
}

func (s *ImageTranslationProRequest) GetSourceLanguage() *string {
	return s.SourceLanguage
}

func (s *ImageTranslationProRequest) GetTargetLanguage() *string {
	return s.TargetLanguage
}

func (s *ImageTranslationProRequest) GetTranslatingBrandInTheProduct() *bool {
	return s.TranslatingBrandInTheProduct
}

func (s *ImageTranslationProRequest) GetUseImageEditor() *bool {
	return s.UseImageEditor
}

func (s *ImageTranslationProRequest) SetAsync(v bool) *ImageTranslationProRequest {
	s.Async = &v
	return s
}

func (s *ImageTranslationProRequest) SetGlossary(v string) *ImageTranslationProRequest {
	s.Glossary = &v
	return s
}

func (s *ImageTranslationProRequest) SetImageUrl(v string) *ImageTranslationProRequest {
	s.ImageUrl = &v
	return s
}

func (s *ImageTranslationProRequest) SetIncludingProductArea(v bool) *ImageTranslationProRequest {
	s.IncludingProductArea = &v
	return s
}

func (s *ImageTranslationProRequest) SetSourceLanguage(v string) *ImageTranslationProRequest {
	s.SourceLanguage = &v
	return s
}

func (s *ImageTranslationProRequest) SetTargetLanguage(v string) *ImageTranslationProRequest {
	s.TargetLanguage = &v
	return s
}

func (s *ImageTranslationProRequest) SetTranslatingBrandInTheProduct(v bool) *ImageTranslationProRequest {
	s.TranslatingBrandInTheProduct = &v
	return s
}

func (s *ImageTranslationProRequest) SetUseImageEditor(v bool) *ImageTranslationProRequest {
	s.UseImageEditor = &v
	return s
}

func (s *ImageTranslationProRequest) Validate() error {
	return dara.Validate(s)
}
