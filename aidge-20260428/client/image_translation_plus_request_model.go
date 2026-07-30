// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageTranslationPlusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGlossary(v string) *ImageTranslationPlusRequest
	GetGlossary() *string
	SetImageUrl(v string) *ImageTranslationPlusRequest
	GetImageUrl() *string
	SetIncludingProductArea(v bool) *ImageTranslationPlusRequest
	GetIncludingProductArea() *bool
	SetSourceLanguage(v string) *ImageTranslationPlusRequest
	GetSourceLanguage() *string
	SetTargetLanguage(v string) *ImageTranslationPlusRequest
	GetTargetLanguage() *string
	SetTranslatingBrandInTheProduct(v bool) *ImageTranslationPlusRequest
	GetTranslatingBrandInTheProduct() *bool
	SetUseImageEditor(v bool) *ImageTranslationPlusRequest
	GetUseImageEditor() *bool
}

type ImageTranslationPlusRequest struct {
	// The ID of the intervention glossary. This parameter is optional.
	//
	// example:
	//
	// glossary-001
	Glossary *string `json:"Glossary,omitempty" xml:"Glossary,omitempty"`
	// The URL of the original image. This parameter is required.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://img.alicdn.com/example.png
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// Specifies whether to translate text on the product body. This parameter is optional. Default value: false.
	//
	// example:
	//
	// false
	IncludingProductArea *bool `json:"IncludingProductArea,omitempty" xml:"IncludingProductArea,omitempty"`
	// The source language. This parameter is required.
	//
	// This parameter is required.
	//
	// example:
	//
	// zh
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	// The target language. This parameter is required.
	//
	// This parameter is required.
	//
	// example:
	//
	// en
	TargetLanguage *string `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
	// Specifies whether to translate brand text on the product. This parameter is optional. Default value: false.
	//
	// example:
	//
	// false
	TranslatingBrandInTheProduct *bool `json:"TranslatingBrandInTheProduct,omitempty" xml:"TranslatingBrandInTheProduct,omitempty"`
	// Specifies whether to use the image translation editor protocol. This parameter is optional.
	//
	// example:
	//
	// false
	UseImageEditor *bool `json:"UseImageEditor,omitempty" xml:"UseImageEditor,omitempty"`
}

func (s ImageTranslationPlusRequest) String() string {
	return dara.Prettify(s)
}

func (s ImageTranslationPlusRequest) GoString() string {
	return s.String()
}

func (s *ImageTranslationPlusRequest) GetGlossary() *string {
	return s.Glossary
}

func (s *ImageTranslationPlusRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *ImageTranslationPlusRequest) GetIncludingProductArea() *bool {
	return s.IncludingProductArea
}

func (s *ImageTranslationPlusRequest) GetSourceLanguage() *string {
	return s.SourceLanguage
}

func (s *ImageTranslationPlusRequest) GetTargetLanguage() *string {
	return s.TargetLanguage
}

func (s *ImageTranslationPlusRequest) GetTranslatingBrandInTheProduct() *bool {
	return s.TranslatingBrandInTheProduct
}

func (s *ImageTranslationPlusRequest) GetUseImageEditor() *bool {
	return s.UseImageEditor
}

func (s *ImageTranslationPlusRequest) SetGlossary(v string) *ImageTranslationPlusRequest {
	s.Glossary = &v
	return s
}

func (s *ImageTranslationPlusRequest) SetImageUrl(v string) *ImageTranslationPlusRequest {
	s.ImageUrl = &v
	return s
}

func (s *ImageTranslationPlusRequest) SetIncludingProductArea(v bool) *ImageTranslationPlusRequest {
	s.IncludingProductArea = &v
	return s
}

func (s *ImageTranslationPlusRequest) SetSourceLanguage(v string) *ImageTranslationPlusRequest {
	s.SourceLanguage = &v
	return s
}

func (s *ImageTranslationPlusRequest) SetTargetLanguage(v string) *ImageTranslationPlusRequest {
	s.TargetLanguage = &v
	return s
}

func (s *ImageTranslationPlusRequest) SetTranslatingBrandInTheProduct(v bool) *ImageTranslationPlusRequest {
	s.TranslatingBrandInTheProduct = &v
	return s
}

func (s *ImageTranslationPlusRequest) SetUseImageEditor(v bool) *ImageTranslationPlusRequest {
	s.UseImageEditor = &v
	return s
}

func (s *ImageTranslationPlusRequest) Validate() error {
	return dara.Validate(s)
}
