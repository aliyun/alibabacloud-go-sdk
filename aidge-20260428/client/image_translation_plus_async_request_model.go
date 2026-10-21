// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageTranslationPlusAsyncRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *ImageTranslationPlusAsyncRequest
	GetBizId() *string
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
	// The business order ID of the caller. This ID is recorded when the request is accepted and returned as is when you query the asynchronous task result.
	//
	// example:
	//
	// order-20260928-001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// The intervention glossary ID. You must create the intervention glossary separately and provide its ID in the input parameters. If the provided intervention glossary ID is empty, the translation results are not modified.
	//
	// example:
	//
	// glossary_1
	Glossary *string `json:"Glossary,omitempty" xml:"Glossary,omitempty"`
	// 	- Image URL: The image must be accessible over the Internet.
	//
	// 	- Format: PNG, JPEG, JPG, BMP, or WEBP.
	//
	// 	- Resolution: The width and height cannot exceed 4,000 pixels.
	//
	// 	- File size: The original file size cannot exceed 10 MB.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/image.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// Specifies whether to translate the text on the image subject. This helps you protect information and prevent embedded information, such as the product name, from being translated. In the example image, the text in the red box represents the text on the subject, and the text in the blue box represents the text not on the subject.
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
	// The target language code. We recommend that you use the 12 minor language pairs listed above.
	//
	// This parameter is required.
	//
	// example:
	//
	// en
	TargetLanguage *string `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
	// Specifies whether to translate the brand name in the image. This prevents the brand name from being translated. Default value: false. In the example image, \\"懒角落\\" represents the brand name.
	//
	// example:
	//
	// false
	TranslatingBrandInTheProduct *bool `json:"TranslatingBrandInTheProduct,omitempty" xml:"TranslatingBrandInTheProduct,omitempty"`
	// Specifies whether to return layout information such as the position, font, and color of the text. This can be used to obtain data when you integrate the image editor for secondary editing. You can call the image translation result data query API operation to obtain the complete parameter information.
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

func (s *ImageTranslationPlusAsyncRequest) GetBizId() *string {
	return s.BizId
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

func (s *ImageTranslationPlusAsyncRequest) SetBizId(v string) *ImageTranslationPlusAsyncRequest {
	s.BizId = &v
	return s
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
