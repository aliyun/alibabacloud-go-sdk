// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageTranslationStandardAsyncRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGlossary(v string) *ImageTranslationStandardAsyncRequest
	GetGlossary() *string
	SetImageUrl(v string) *ImageTranslationStandardAsyncRequest
	GetImageUrl() *string
	SetIncludingProductArea(v bool) *ImageTranslationStandardAsyncRequest
	GetIncludingProductArea() *bool
	SetSourceLanguage(v string) *ImageTranslationStandardAsyncRequest
	GetSourceLanguage() *string
	SetTargetLanguage(v string) *ImageTranslationStandardAsyncRequest
	GetTargetLanguage() *string
	SetTranslatingBrandInTheProduct(v bool) *ImageTranslationStandardAsyncRequest
	GetTranslatingBrandInTheProduct() *bool
	SetUseImageEditor(v bool) *ImageTranslationStandardAsyncRequest
	GetUseImageEditor() *bool
}

type ImageTranslationStandardAsyncRequest struct {
	// The intervention glossary.
	//
	// example:
	//
	// glossary_1
	Glossary *string `json:"Glossary,omitempty" xml:"Glossary,omitempty"`
	// The URL of the image to be translated.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/image.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// Specifies whether to include the product area in the translation.
	//
	// example:
	//
	// false
	IncludingProductArea *bool `json:"IncludingProductArea,omitempty" xml:"IncludingProductArea,omitempty"`
	// The source language.
	//
	// This parameter is required.
	//
	// example:
	//
	// zh
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	// The target language.
	//
	// This parameter is required.
	//
	// example:
	//
	// en
	TargetLanguage *string `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
	// Specifies whether to translate brand text in the product area.
	//
	// example:
	//
	// false
	TranslatingBrandInTheProduct *bool `json:"TranslatingBrandInTheProduct,omitempty" xml:"TranslatingBrandInTheProduct,omitempty"`
	// Specifies whether to use the image editor protocol.
	//
	// example:
	//
	// false
	UseImageEditor *bool `json:"UseImageEditor,omitempty" xml:"UseImageEditor,omitempty"`
}

func (s ImageTranslationStandardAsyncRequest) String() string {
	return dara.Prettify(s)
}

func (s ImageTranslationStandardAsyncRequest) GoString() string {
	return s.String()
}

func (s *ImageTranslationStandardAsyncRequest) GetGlossary() *string {
	return s.Glossary
}

func (s *ImageTranslationStandardAsyncRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *ImageTranslationStandardAsyncRequest) GetIncludingProductArea() *bool {
	return s.IncludingProductArea
}

func (s *ImageTranslationStandardAsyncRequest) GetSourceLanguage() *string {
	return s.SourceLanguage
}

func (s *ImageTranslationStandardAsyncRequest) GetTargetLanguage() *string {
	return s.TargetLanguage
}

func (s *ImageTranslationStandardAsyncRequest) GetTranslatingBrandInTheProduct() *bool {
	return s.TranslatingBrandInTheProduct
}

func (s *ImageTranslationStandardAsyncRequest) GetUseImageEditor() *bool {
	return s.UseImageEditor
}

func (s *ImageTranslationStandardAsyncRequest) SetGlossary(v string) *ImageTranslationStandardAsyncRequest {
	s.Glossary = &v
	return s
}

func (s *ImageTranslationStandardAsyncRequest) SetImageUrl(v string) *ImageTranslationStandardAsyncRequest {
	s.ImageUrl = &v
	return s
}

func (s *ImageTranslationStandardAsyncRequest) SetIncludingProductArea(v bool) *ImageTranslationStandardAsyncRequest {
	s.IncludingProductArea = &v
	return s
}

func (s *ImageTranslationStandardAsyncRequest) SetSourceLanguage(v string) *ImageTranslationStandardAsyncRequest {
	s.SourceLanguage = &v
	return s
}

func (s *ImageTranslationStandardAsyncRequest) SetTargetLanguage(v string) *ImageTranslationStandardAsyncRequest {
	s.TargetLanguage = &v
	return s
}

func (s *ImageTranslationStandardAsyncRequest) SetTranslatingBrandInTheProduct(v bool) *ImageTranslationStandardAsyncRequest {
	s.TranslatingBrandInTheProduct = &v
	return s
}

func (s *ImageTranslationStandardAsyncRequest) SetUseImageEditor(v bool) *ImageTranslationStandardAsyncRequest {
	s.UseImageEditor = &v
	return s
}

func (s *ImageTranslationStandardAsyncRequest) Validate() error {
	return dara.Validate(s)
}
