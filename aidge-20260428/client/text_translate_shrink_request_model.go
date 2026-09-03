// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTextTranslateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizName(v string) *TextTranslateShrinkRequest
	GetBizName() *string
	SetFormatType(v string) *TextTranslateShrinkRequest
	GetFormatType() *string
	SetGlossary(v string) *TextTranslateShrinkRequest
	GetGlossary() *string
	SetSourceLanguage(v string) *TextTranslateShrinkRequest
	GetSourceLanguage() *string
	SetSourceTextListShrink(v string) *TextTranslateShrinkRequest
	GetSourceTextListShrink() *string
	SetTargetLanguage(v string) *TextTranslateShrinkRequest
	GetTargetLanguage() *string
	SetTranslateScene(v string) *TextTranslateShrinkRequest
	GetTranslateScene() *string
}

type TextTranslateShrinkRequest struct {
	// This field represents your identity and facilitates communication for various issues.
	//
	// ● If you are an Alibaba internal organization, specify a value based on your actual scenario, such as BU name-product or BU name-chat.
	//
	// ● If you are an external Alibaba partner, specify the full name of your company. This company name must be consistent with the company name used when you registered your Alibaba Cloud account.
	//
	// example:
	//
	// MyCompany-Chat
	BizName *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	// The format type of the source text. This parameter is optional. Valid values: text (plain text format) and html (web page format, which preserves HTML tags).
	//
	// example:
	//
	// text
	FormatType *string `json:"FormatType,omitempty" xml:"FormatType,omitempty"`
	// The intervention glossary ID. This parameter is optional. The glossary must be created separately in the console, and its ID must be provided. If the glossary ID is empty, the translation results are not modified.
	//
	// example:
	//
	// glossary_1
	Glossary *string `json:"Glossary,omitempty" xml:"Glossary,omitempty"`
	// The source language code. This parameter is optional. If not specified, the language is automatically detected. You can set this parameter to auto for language detection. For supported language pairs, see [Language pair mapping table](https://www.alibabacloud.com/help/en/document_detail/3041883.html).
	//
	// example:
	//
	// auto
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	// The list of texts to translate. This parameter is required. The total character length cannot exceed 50,000, and the list length cannot exceed 50.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["Hello world"]
	SourceTextListShrink *string `json:"SourceTextList,omitempty" xml:"SourceTextList,omitempty"`
	// The target language code. This parameter is required. For supported language pairs, see [Language pair mapping table](https://www.alibabacloud.com/help/en/document_detail/3041883.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ko
	TargetLanguage *string `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
	// The business scenario identifier. You can pass in only one of the following values. When specified, the translation engine invokes the corresponding industry terminology library and style strategy to produce translations that better fit the industry. If this field is not specified or an invalid value is passed, the general translation strategy is used.
	//
	// Valid values:
	//
	// ● e-commerce-title: cross-border e-commerce product title translation
	//
	// ● e-commerce-description: cross-border e-commerce product description translation
	//
	// ● e-commerce-chat: cross-border e-commerce conversation translation
	//
	// ● e-commerce-cpv: cross-border e-commerce product CPV attribute translation
	//
	// ● novel: novel translation
	//
	// ● game: game translation
	//
	// example:
	//
	// e-commerce-title
	TranslateScene *string `json:"TranslateScene,omitempty" xml:"TranslateScene,omitempty"`
}

func (s TextTranslateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s TextTranslateShrinkRequest) GoString() string {
	return s.String()
}

func (s *TextTranslateShrinkRequest) GetBizName() *string {
	return s.BizName
}

func (s *TextTranslateShrinkRequest) GetFormatType() *string {
	return s.FormatType
}

func (s *TextTranslateShrinkRequest) GetGlossary() *string {
	return s.Glossary
}

func (s *TextTranslateShrinkRequest) GetSourceLanguage() *string {
	return s.SourceLanguage
}

func (s *TextTranslateShrinkRequest) GetSourceTextListShrink() *string {
	return s.SourceTextListShrink
}

func (s *TextTranslateShrinkRequest) GetTargetLanguage() *string {
	return s.TargetLanguage
}

func (s *TextTranslateShrinkRequest) GetTranslateScene() *string {
	return s.TranslateScene
}

func (s *TextTranslateShrinkRequest) SetBizName(v string) *TextTranslateShrinkRequest {
	s.BizName = &v
	return s
}

func (s *TextTranslateShrinkRequest) SetFormatType(v string) *TextTranslateShrinkRequest {
	s.FormatType = &v
	return s
}

func (s *TextTranslateShrinkRequest) SetGlossary(v string) *TextTranslateShrinkRequest {
	s.Glossary = &v
	return s
}

func (s *TextTranslateShrinkRequest) SetSourceLanguage(v string) *TextTranslateShrinkRequest {
	s.SourceLanguage = &v
	return s
}

func (s *TextTranslateShrinkRequest) SetSourceTextListShrink(v string) *TextTranslateShrinkRequest {
	s.SourceTextListShrink = &v
	return s
}

func (s *TextTranslateShrinkRequest) SetTargetLanguage(v string) *TextTranslateShrinkRequest {
	s.TargetLanguage = &v
	return s
}

func (s *TextTranslateShrinkRequest) SetTranslateScene(v string) *TextTranslateShrinkRequest {
	s.TranslateScene = &v
	return s
}

func (s *TextTranslateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
