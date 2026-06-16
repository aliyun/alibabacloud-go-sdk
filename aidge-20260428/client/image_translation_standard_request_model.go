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
}

type ImageTranslationStandardRequest struct {
	// The ID of the intervention glossary. This parameter is optional. Create the glossary in the console and provide its ID. If the glossary ID is empty, the translation results are not modified.
	//
	// example:
	//
	// glossary_1
	Glossary *string `json:"Glossary,omitempty" xml:"Glossary,omitempty"`
	// The URL of the original image. This parameter is required. Image requirements: the width and height cannot exceed 4000 × 4000 pixels, the file size cannot exceed 10 MB, and the supported formats are png, jpeg, jpg, bmp, and webp.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://images-na.ssl-images-amazon.com/images/I/41bKsNBDcwL.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// Specifies whether to translate text on the product subject in the image. This parameter is optional. Default value: false. This helps protect information by preventing translation of embedded information such as product names.
	//
	// example:
	//
	// false
	IncludingProductArea *bool `json:"IncludingProductArea,omitempty" xml:"IncludingProductArea,omitempty"`
	// The source language code. This parameter is required. For supported language directions, see the supported language directions list.
	//
	// This parameter is required.
	//
	// example:
	//
	// en
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	// The target language code. This parameter is required. For supported language directions, see the supported language directions list.
	//
	// This parameter is required.
	//
	// example:
	//
	// ko
	TargetLanguage *string `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
	// Specifies whether to translate brand names on the image. This parameter is optional. Default value: false. This helps protect brand name information from being translated.
	//
	// example:
	//
	// false
	TranslatingBrandInTheProduct *bool `json:"TranslatingBrandInTheProduct,omitempty" xml:"TranslatingBrandInTheProduct,omitempty"`
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

func (s *ImageTranslationStandardRequest) Validate() error {
	return dara.Validate(s)
}
