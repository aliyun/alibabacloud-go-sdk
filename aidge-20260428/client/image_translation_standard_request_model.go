// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageTranslationStandardRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGlossary(v string) *ImageTranslationStandardRequest
	GetGlossary() *string
	SetImageUrl(v string) *ImageTranslationStandardRequest
	GetImageUrl() *string
	SetIncludingProductArea(v bool) *ImageTranslationStandardRequest
	GetIncludingProductArea() *bool
	SetSourceLanguage(v string) *ImageTranslationStandardRequest
	GetSourceLanguage() *string
	SetTargetLanguage(v string) *ImageTranslationStandardRequest
	GetTargetLanguage() *string
	SetTranslatingBrandInTheProduct(v bool) *ImageTranslationStandardRequest
	GetTranslatingBrandInTheProduct() *bool
	SetUseImageEditor(v bool) *ImageTranslationStandardRequest
	GetUseImageEditor() *bool
}

type ImageTranslationStandardRequest struct {
	// The glossary ID. Optional. Create a glossary in the console and provide its ID. If the glossary ID is empty, the translation results are not modified.
	//
	// example:
	//
	// glossary_1
	Glossary *string `json:"Glossary,omitempty" xml:"Glossary,omitempty"`
	// - Image URL: Must be publicly accessible.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://images-na.ssl-images-amazon.com/images/I/41bKsNBDcwL.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// Specifies whether to translate text on the main subject of the image. Optional. Default value: false. This helps you protect information by avoiding translation of embedded information such as product names.
	//
	// example:
	//
	// false
	IncludingProductArea *bool `json:"IncludingProductArea,omitempty" xml:"IncludingProductArea,omitempty"`
	// The source language code. Required. For supported language pairs, see the supported language pair list.
	//
	// This parameter is required.
	//
	// example:
	//
	// en
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	// The target language code. Required. For supported language pairs, see the supported language pair list.
	//
	// This parameter is required.
	//
	// example:
	//
	// ko
	TargetLanguage *string `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
	// Specifies whether to translate brand names on the image. Optional. Default value: false. This helps you protect brand name information from being translated.
	//
	// example:
	//
	// false
	TranslatingBrandInTheProduct *bool `json:"TranslatingBrandInTheProduct,omitempty" xml:"TranslatingBrandInTheProduct,omitempty"`
	// Specifies whether to return layer information such as text position, font, and color. When set to true, layer information is returned for integration with image editors for secondary editing. Default value: false.
	//
	// example:
	//
	// false
	UseImageEditor *bool `json:"UseImageEditor,omitempty" xml:"UseImageEditor,omitempty"`
}

func (s ImageTranslationStandardRequest) String() string {
	return dara.Prettify(s)
}

func (s ImageTranslationStandardRequest) GoString() string {
	return s.String()
}

func (s *ImageTranslationStandardRequest) GetGlossary() *string {
	return s.Glossary
}

func (s *ImageTranslationStandardRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *ImageTranslationStandardRequest) GetIncludingProductArea() *bool {
	return s.IncludingProductArea
}

func (s *ImageTranslationStandardRequest) GetSourceLanguage() *string {
	return s.SourceLanguage
}

func (s *ImageTranslationStandardRequest) GetTargetLanguage() *string {
	return s.TargetLanguage
}

func (s *ImageTranslationStandardRequest) GetTranslatingBrandInTheProduct() *bool {
	return s.TranslatingBrandInTheProduct
}

func (s *ImageTranslationStandardRequest) GetUseImageEditor() *bool {
	return s.UseImageEditor
}

func (s *ImageTranslationStandardRequest) SetGlossary(v string) *ImageTranslationStandardRequest {
	s.Glossary = &v
	return s
}

func (s *ImageTranslationStandardRequest) SetImageUrl(v string) *ImageTranslationStandardRequest {
	s.ImageUrl = &v
	return s
}

func (s *ImageTranslationStandardRequest) SetIncludingProductArea(v bool) *ImageTranslationStandardRequest {
	s.IncludingProductArea = &v
	return s
}

func (s *ImageTranslationStandardRequest) SetSourceLanguage(v string) *ImageTranslationStandardRequest {
	s.SourceLanguage = &v
	return s
}

func (s *ImageTranslationStandardRequest) SetTargetLanguage(v string) *ImageTranslationStandardRequest {
	s.TargetLanguage = &v
	return s
}

func (s *ImageTranslationStandardRequest) SetTranslatingBrandInTheProduct(v bool) *ImageTranslationStandardRequest {
	s.TranslatingBrandInTheProduct = &v
	return s
}

func (s *ImageTranslationStandardRequest) SetUseImageEditor(v bool) *ImageTranslationStandardRequest {
	s.UseImageEditor = &v
	return s
}

func (s *ImageTranslationStandardRequest) Validate() error {
	return dara.Validate(s)
}
