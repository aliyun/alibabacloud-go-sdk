// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageTranslationProAsyncRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGlossary(v string) *ImageTranslationProAsyncRequest
	GetGlossary() *string
	SetImageUrl(v string) *ImageTranslationProAsyncRequest
	GetImageUrl() *string
	SetIncludingProductArea(v bool) *ImageTranslationProAsyncRequest
	GetIncludingProductArea() *bool
	SetSourceLanguage(v string) *ImageTranslationProAsyncRequest
	GetSourceLanguage() *string
	SetTargetLanguage(v string) *ImageTranslationProAsyncRequest
	GetTargetLanguage() *string
	SetTranslatingBrandInTheProduct(v bool) *ImageTranslationProAsyncRequest
	GetTranslatingBrandInTheProduct() *bool
	SetUseImageEditor(v bool) *ImageTranslationProAsyncRequest
	GetUseImageEditor() *bool
}

type ImageTranslationProAsyncRequest struct {
	// The ID of the intervention glossary. You must create the glossary in the console and provide its ID. If this parameter is left empty, the translation results are not modified.
	//
	// example:
	//
	// glossary_1
	Glossary *string `json:"Glossary,omitempty" xml:"Glossary,omitempty"`
	// • Image URL: Must be accessible over the Internet.
	//
	// • Format: PNG, JPEG, JPG, BMP, or WEBP.
	//
	// • Pixels: Both the width and height must not exceed 4,000.
	//
	// • File size: The original file size must be less than or equal to 10 MB.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/image.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// Specifies whether to translate the text on the product subject in the image. Set this parameter to false to protect the subject information, such as the product name, from being translated. Default value: false.
	//
	// example:
	//
	// false
	IncludingProductArea *bool `json:"IncludingProductArea,omitempty" xml:"IncludingProductArea,omitempty"`
	// The source language code. The value auto is supported for automatic detection.
	//
	// This parameter is required.
	//
	// example:
	//
	// zh
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	// The target language code. For supported language pairs, see the appendix.
	//
	// This parameter is required.
	//
	// example:
	//
	// en
	TargetLanguage *string `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
	// Specifies whether to translate the brand names in the image. Default value: false, which indicates that brand names are not translated.
	//
	// example:
	//
	// false
	TranslatingBrandInTheProduct *bool `json:"TranslatingBrandInTheProduct,omitempty" xml:"TranslatingBrandInTheProduct,omitempty"`
	// Specifies whether to return layout information such as the position, font, and color of the text. If you set this parameter to true, layer information is returned, which can be used for secondary editing in an image editor. Default value: false.
	//
	// example:
	//
	// false
	UseImageEditor *bool `json:"UseImageEditor,omitempty" xml:"UseImageEditor,omitempty"`
}

func (s ImageTranslationProAsyncRequest) String() string {
	return dara.Prettify(s)
}

func (s ImageTranslationProAsyncRequest) GoString() string {
	return s.String()
}

func (s *ImageTranslationProAsyncRequest) GetGlossary() *string {
	return s.Glossary
}

func (s *ImageTranslationProAsyncRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *ImageTranslationProAsyncRequest) GetIncludingProductArea() *bool {
	return s.IncludingProductArea
}

func (s *ImageTranslationProAsyncRequest) GetSourceLanguage() *string {
	return s.SourceLanguage
}

func (s *ImageTranslationProAsyncRequest) GetTargetLanguage() *string {
	return s.TargetLanguage
}

func (s *ImageTranslationProAsyncRequest) GetTranslatingBrandInTheProduct() *bool {
	return s.TranslatingBrandInTheProduct
}

func (s *ImageTranslationProAsyncRequest) GetUseImageEditor() *bool {
	return s.UseImageEditor
}

func (s *ImageTranslationProAsyncRequest) SetGlossary(v string) *ImageTranslationProAsyncRequest {
	s.Glossary = &v
	return s
}

func (s *ImageTranslationProAsyncRequest) SetImageUrl(v string) *ImageTranslationProAsyncRequest {
	s.ImageUrl = &v
	return s
}

func (s *ImageTranslationProAsyncRequest) SetIncludingProductArea(v bool) *ImageTranslationProAsyncRequest {
	s.IncludingProductArea = &v
	return s
}

func (s *ImageTranslationProAsyncRequest) SetSourceLanguage(v string) *ImageTranslationProAsyncRequest {
	s.SourceLanguage = &v
	return s
}

func (s *ImageTranslationProAsyncRequest) SetTargetLanguage(v string) *ImageTranslationProAsyncRequest {
	s.TargetLanguage = &v
	return s
}

func (s *ImageTranslationProAsyncRequest) SetTranslatingBrandInTheProduct(v bool) *ImageTranslationProAsyncRequest {
	s.TranslatingBrandInTheProduct = &v
	return s
}

func (s *ImageTranslationProAsyncRequest) SetUseImageEditor(v bool) *ImageTranslationProAsyncRequest {
	s.UseImageEditor = &v
	return s
}

func (s *ImageTranslationProAsyncRequest) Validate() error {
	return dara.Validate(s)
}
