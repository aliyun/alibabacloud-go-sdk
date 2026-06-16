// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageTranslationProRequest interface {
	dara.Model
	String() string
	GoString() string
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
	SetCallType(v string) *ImageTranslationProRequest
	GetCallType() *string
}

type ImageTranslationProRequest struct {
	// The ID of the intervention glossary. Optional. Create the glossary in the console and provide its ID. If the glossary ID is empty, the translation results are not modified.
	//
	// example:
	//
	// glossary_1
	Glossary *string `json:"Glossary,omitempty" xml:"Glossary,omitempty"`
	// The URL of the original image. Required. Image requirements: width and height must not exceed 4000 × 4000. File size must not exceed 10 MB. Supported formats: png, jpeg, jpg, bmp, and webp.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://img.alicdn.com/imgextra/i3/O1CN01HTDhDi28Fd85ZYs7H_!!6000000007903-0-tps-800-800.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// Specifies whether to translate text on the image subject. Optional. Default value: false. This helps protect information such as embedded product names from being translated.
	//
	// example:
	//
	// false
	IncludingProductArea *bool `json:"IncludingProductArea,omitempty" xml:"IncludingProductArea,omitempty"`
	// The source language code. Required. For supported language directions, see the supported language direction list.
	//
	// This parameter is required.
	//
	// example:
	//
	// en
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	// The target language code. Required. For supported language directions, see the supported language direction list.
	//
	// This parameter is required.
	//
	// example:
	//
	// ko
	TargetLanguage *string `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
	// Specifies whether to translate brand names on images. Optional. Default value: false. This helps protect brand name information from being translated.
	//
	// example:
	//
	// false
	TranslatingBrandInTheProduct *bool `json:"TranslatingBrandInTheProduct,omitempty" xml:"TranslatingBrandInTheProduct,omitempty"`
	// Specifies whether to return layout information such as text position, font, and color. Optional. Default value: false. This can be used for secondary editing when integrated with an image editor.
	//
	// example:
	//
	// false
	UseImageEditor *bool `json:"UseImageEditor,omitempty" xml:"UseImageEditor,omitempty"`
	// The call type. This is an internal system parameter. Optional.
	//
	// example:
	//
	// sync
	CallType *string `json:"callType,omitempty" xml:"callType,omitempty"`
}

func (s ImageTranslationProRequest) String() string {
	return dara.Prettify(s)
}

func (s ImageTranslationProRequest) GoString() string {
	return s.String()
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

func (s *ImageTranslationProRequest) GetCallType() *string {
	return s.CallType
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

func (s *ImageTranslationProRequest) SetCallType(v string) *ImageTranslationProRequest {
	s.CallType = &v
	return s
}

func (s *ImageTranslationProRequest) Validate() error {
	return dara.Validate(s)
}
