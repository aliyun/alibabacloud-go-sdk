// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateProductContentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBulletPointCount(v int32) *GenerateProductContentRequest
	GetBulletPointCount() *int32
	SetContentType(v string) *GenerateProductContentRequest
	GetContentType() *string
	SetEnableTranslation(v bool) *GenerateProductContentRequest
	GetEnableTranslation() *bool
	SetMaxLength(v int32) *GenerateProductContentRequest
	GetMaxLength() *int32
	SetMinLength(v int32) *GenerateProductContentRequest
	GetMinLength() *int32
	SetOperation(v string) *GenerateProductContentRequest
	GetOperation() *string
	SetPlatform(v string) *GenerateProductContentRequest
	GetPlatform() *string
	SetProductAttributes(v string) *GenerateProductContentRequest
	GetProductAttributes() *string
	SetProductCategory(v string) *GenerateProductContentRequest
	GetProductCategory() *string
	SetSourceLanguage(v string) *GenerateProductContentRequest
	GetSourceLanguage() *string
	SetSourceText(v string) *GenerateProductContentRequest
	GetSourceText() *string
	SetTargetLanguage(v string) *GenerateProductContentRequest
	GetTargetLanguage() *string
}

type GenerateProductContentRequest struct {
	// Takes effect only when contentType is set to bulletPoint. Valid values: 1 to 5. Default value in generate mode: 1. In rewrite mode, if manually specified, it must match the number of non-empty paragraphs in sourceText. This parameter is ignored for other copy types.
	//
	// example:
	//
	// 5
	BulletPointCount *int32 `json:"BulletPointCount,omitempty" xml:"BulletPointCount,omitempty"`
	// The copy type. Valid values: title, highlight, keywords, bulletPoint, and description. When a platform is specified, you must select a type supported by the platform.
	//
	// This parameter is required.
	//
	// example:
	//
	// description
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// Specifies whether to enable translation. Default value: false. Valid values: true and false.
	//
	// example:
	//
	// True
	EnableTranslation *bool `json:"EnableTranslation,omitempty" xml:"EnableTranslation,omitempty"`
	// The maximum number of characters for the final complete output text. If specified, it must be an integer greater than or equal to 30. If a platform is specified and this parameter is not passed, the configured upper limit for the type on the platform is used. Example: 60.
	//
	// example:
	//
	// 1000
	MaxLength *int32 `json:"MaxLength,omitempty" xml:"MaxLength,omitempty"`
	// The minimum number of characters for the final complete output text. If specified, it must be an integer greater than or equal to 10. If a platform is specified and this parameter is not passed, the configured lower limit for the type on the platform is used.
	//
	// example:
	//
	// 500
	MinLength *int32 `json:"MinLength,omitempty" xml:"MinLength,omitempty"`
	// The operation type. Valid values: generate and rewrite.
	//
	// This parameter is required.
	//
	// example:
	//
	// generate
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// The target platform. Valid values: amazon, temu, tiktok, shopify, Wildberries, and tmall. If this parameter is not passed, the general scenario is used. Enum values are case-sensitive.
	//
	// example:
	//
	// amazon
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The product attribute materials. Default value: empty. The length is 0 to 500 characters.
	//
	// example:
	//
	// {"brand":"Supor","material":"304 stainless steel"}
	ProductAttributes *string `json:"ProductAttributes,omitempty" xml:"ProductAttributes,omitempty"`
	// The product category. Default value: an empty string. The length is 0 to 100 characters. It helps understand the category during generation and eliminates ambiguity during rewriting.
	//
	// example:
	//
	// Kitchen utensils
	ProductCategory *string `json:"ProductCategory,omitempty" xml:"ProductCategory,omitempty"`
	// The source language. Valid values: zh and en. Automatic language identification is not supported. Supported source languages: Chinese (zh) and English (en).
	//
	// This parameter is required.
	//
	// example:
	//
	// zh
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	// The product materials or copy to be rewritten. The length is 20 to 3,000 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// Stainless steel steamer 28cm 3-tier large capacity household kitchen essential steaming and boiling dual-use cookware
	SourceText *string `json:"SourceText,omitempty" xml:"SourceText,omitempty"`
	// Required when enableTranslation is set to true. For valid values, see Multilingual processing. If translation is disabled, this parameter is optional. If specified, it must be the same as sourceLanguage. Supported target languages: Chinese (zh), English (en), Russian (ru), German (de), French (fr), Korean (ko), Italian (it), Thai (th), Arabic (ar), Turkish (tr), Hebrew (he), Ukrainian (uk), Polish (pl), Japanese (ja), Dutch (nl), Portuguese (pt), Spanish (es), Malay (ms), Bulgarian (bg), Indonesian (id), and Vietnamese (vi).
	//
	// example:
	//
	// en
	TargetLanguage *string `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
}

func (s GenerateProductContentRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateProductContentRequest) GoString() string {
	return s.String()
}

func (s *GenerateProductContentRequest) GetBulletPointCount() *int32 {
	return s.BulletPointCount
}

func (s *GenerateProductContentRequest) GetContentType() *string {
	return s.ContentType
}

func (s *GenerateProductContentRequest) GetEnableTranslation() *bool {
	return s.EnableTranslation
}

func (s *GenerateProductContentRequest) GetMaxLength() *int32 {
	return s.MaxLength
}

func (s *GenerateProductContentRequest) GetMinLength() *int32 {
	return s.MinLength
}

func (s *GenerateProductContentRequest) GetOperation() *string {
	return s.Operation
}

func (s *GenerateProductContentRequest) GetPlatform() *string {
	return s.Platform
}

func (s *GenerateProductContentRequest) GetProductAttributes() *string {
	return s.ProductAttributes
}

func (s *GenerateProductContentRequest) GetProductCategory() *string {
	return s.ProductCategory
}

func (s *GenerateProductContentRequest) GetSourceLanguage() *string {
	return s.SourceLanguage
}

func (s *GenerateProductContentRequest) GetSourceText() *string {
	return s.SourceText
}

func (s *GenerateProductContentRequest) GetTargetLanguage() *string {
	return s.TargetLanguage
}

func (s *GenerateProductContentRequest) SetBulletPointCount(v int32) *GenerateProductContentRequest {
	s.BulletPointCount = &v
	return s
}

func (s *GenerateProductContentRequest) SetContentType(v string) *GenerateProductContentRequest {
	s.ContentType = &v
	return s
}

func (s *GenerateProductContentRequest) SetEnableTranslation(v bool) *GenerateProductContentRequest {
	s.EnableTranslation = &v
	return s
}

func (s *GenerateProductContentRequest) SetMaxLength(v int32) *GenerateProductContentRequest {
	s.MaxLength = &v
	return s
}

func (s *GenerateProductContentRequest) SetMinLength(v int32) *GenerateProductContentRequest {
	s.MinLength = &v
	return s
}

func (s *GenerateProductContentRequest) SetOperation(v string) *GenerateProductContentRequest {
	s.Operation = &v
	return s
}

func (s *GenerateProductContentRequest) SetPlatform(v string) *GenerateProductContentRequest {
	s.Platform = &v
	return s
}

func (s *GenerateProductContentRequest) SetProductAttributes(v string) *GenerateProductContentRequest {
	s.ProductAttributes = &v
	return s
}

func (s *GenerateProductContentRequest) SetProductCategory(v string) *GenerateProductContentRequest {
	s.ProductCategory = &v
	return s
}

func (s *GenerateProductContentRequest) SetSourceLanguage(v string) *GenerateProductContentRequest {
	s.SourceLanguage = &v
	return s
}

func (s *GenerateProductContentRequest) SetSourceText(v string) *GenerateProductContentRequest {
	s.SourceText = &v
	return s
}

func (s *GenerateProductContentRequest) SetTargetLanguage(v string) *GenerateProductContentRequest {
	s.TargetLanguage = &v
	return s
}

func (s *GenerateProductContentRequest) Validate() error {
	return dara.Validate(s)
}
