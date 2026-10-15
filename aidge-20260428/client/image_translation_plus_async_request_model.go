// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageTranslationPlusAsyncRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGlossary(v string) *ImageTranslationPlusAsyncRequest
	GetGlossary() *string
	SetImageUrl(v string) *ImageTranslationPlusAsyncRequest
	GetImageUrl() *string
	SetIncludingProductArea(v bool) *ImageTranslationPlusAsyncRequest
	GetIncludingProductArea() *bool
	SetSourceLanguage(v string) *ImageTranslationPlusAsyncRequest
	GetSourceLanguage() *string
	SetTargetLanguage(v string) *ImageTranslationPlusAsyncRequest
	GetTargetLanguage() *string
	SetTranslatingBrandInTheProduct(v bool) *ImageTranslationPlusAsyncRequest
	GetTranslatingBrandInTheProduct() *bool
	SetUseImageEditor(v bool) *ImageTranslationPlusAsyncRequest
	GetUseImageEditor() *bool
}

type ImageTranslationPlusAsyncRequest struct {
	// The intervention glossary ID. The intervention glossary ID must be created separately and provided in the input parameters. If the provided intervention glossary ID is empty, the translation results will not be modified.
	//
	// example:
	//
	// glossary_1
	Glossary *string `json:"Glossary,omitempty" xml:"Glossary,omitempty"`
	// • Image URL: Must be publicly accessible.
	//
	// • Format: PNG, JPEG, JPG, BMP, or WEBP.
	//
	// • Pixels: Both width and height must not exceed 4,000.
	//
	// • File size: The original file must be less than or equal to 10 MB.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/image.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// Specifies whether to translate the text on the image subject. This helps you protect information and prevent embedded information, such as product names, from being translated. In the example image, the text in the red box represents the text on the subject, and the text in the blue box represents the text not on the subject.
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
	// The target language code. The 12 minor language pairs listed above are recommended.
	//
	// This parameter is required.
	//
	// example:
	//
	// en
	TargetLanguage *string `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
	// Specifies whether to translate brand names in the image. This prevents brand names from being translated. Default value: false, which indicates that brand names are not translated. In the example image, \\"懒角落\\" represents a brand name.
	//
	// example:
	//
	// false
	TranslatingBrandInTheProduct *bool `json:"TranslatingBrandInTheProduct,omitempty" xml:"TranslatingBrandInTheProduct,omitempty"`
	// Specifies whether to return layout information such as the position, font, and color of the text. This can be used to obtain data when integrating the image editor for secondary editing. You can call the image translation result query API to obtain the complete parameter information.
	//
	// example:
	//
	// false
	UseImageEditor *bool `json:"UseImageEditor,omitempty" xml:"UseImageEditor,omitempty"`
}

func (s ImageTranslationPlusAsyncRequest) String() string {
	return dara.Prettify(s)
}

func (s ImageTranslationPlusAsyncRequest) GoString() string {
	return s.String()
}

func (s *ImageTranslationPlusAsyncRequest) GetGlossary() *string {
	return s.Glossary
}

func (s *ImageTranslationPlusAsyncRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *ImageTranslationPlusAsyncRequest) GetIncludingProductArea() *bool {
	return s.IncludingProductArea
}

func (s *ImageTranslationPlusAsyncRequest) GetSourceLanguage() *string {
	return s.SourceLanguage
}

func (s *ImageTranslationPlusAsyncRequest) GetTargetLanguage() *string {
	return s.TargetLanguage
}

func (s *ImageTranslationPlusAsyncRequest) GetTranslatingBrandInTheProduct() *bool {
	return s.TranslatingBrandInTheProduct
}

func (s *ImageTranslationPlusAsyncRequest) GetUseImageEditor() *bool {
	return s.UseImageEditor
}

func (s *ImageTranslationPlusAsyncRequest) SetGlossary(v string) *ImageTranslationPlusAsyncRequest {
	s.Glossary = &v
	return s
}

func (s *ImageTranslationPlusAsyncRequest) SetImageUrl(v string) *ImageTranslationPlusAsyncRequest {
	s.ImageUrl = &v
	return s
}

func (s *ImageTranslationPlusAsyncRequest) SetIncludingProductArea(v bool) *ImageTranslationPlusAsyncRequest {
	s.IncludingProductArea = &v
	return s
}

func (s *ImageTranslationPlusAsyncRequest) SetSourceLanguage(v string) *ImageTranslationPlusAsyncRequest {
	s.SourceLanguage = &v
	return s
}

func (s *ImageTranslationPlusAsyncRequest) SetTargetLanguage(v string) *ImageTranslationPlusAsyncRequest {
	s.TargetLanguage = &v
	return s
}

func (s *ImageTranslationPlusAsyncRequest) SetTranslatingBrandInTheProduct(v bool) *ImageTranslationPlusAsyncRequest {
	s.TranslatingBrandInTheProduct = &v
	return s
}

func (s *ImageTranslationPlusAsyncRequest) SetUseImageEditor(v bool) *ImageTranslationPlusAsyncRequest {
	s.UseImageEditor = &v
	return s
}

func (s *ImageTranslationPlusAsyncRequest) Validate() error {
	return dara.Validate(s)
}
