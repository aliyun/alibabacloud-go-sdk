// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTextTranslateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizName(v string) *TextTranslateRequest
	GetBizName() *string
	SetFormatType(v string) *TextTranslateRequest
	GetFormatType() *string
	SetGlossary(v string) *TextTranslateRequest
	GetGlossary() *string
	SetSourceLanguage(v string) *TextTranslateRequest
	GetSourceLanguage() *string
	SetSourceTextList(v []*string) *TextTranslateRequest
	GetSourceTextList() []*string
	SetTargetLanguage(v string) *TextTranslateRequest
	GetTargetLanguage() *string
	SetTranslateScene(v string) *TextTranslateRequest
	GetTranslateScene() *string
}

type TextTranslateRequest struct {
	// This field represents your identity and facilitates communication for various issues.
	//
	// ● If you are an internal Alibaba organization, pass a value based on your actual scenario, such as BU name-product or BU name-chat.
	//
	// ● If you are an external Alibaba partner, pass the full name of your company. This company name must be consistent with the company name used when you registered your Alibaba Cloud account.
	//
	// example:
	//
	// MyCompany-Chat
	BizName *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	// The format type of the source text. This parameter is optional. Valid values: text (plain text format) and html (web page format that preserves HTML tags).
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
	// The source language code. If not specified, the language is automatically detected. This parameter is optional. You can pass auto for language detection. For supported language pairs, see [Language pair mapping table](https://www.alibabacloud.com/help/en/document_detail/3041883.html).
	//
	// example:
	//
	// auto
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	// The list of texts to be translated. This parameter is required. The total character length cannot exceed 50,000, and the list length cannot exceed 50.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["Hello world"]
	SourceTextList []*string `json:"SourceTextList,omitempty" xml:"SourceTextList,omitempty" type:"Repeated"`
	// The target language code. This parameter is required. For supported language pairs, see [Language pair mapping table](https://www.alibabacloud.com/help/en/document_detail/3041883.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ko
	TargetLanguage *string `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
	// The business scenario identifier. You can pass only one of the following values. When specified, the translation engine invokes the corresponding industry terminology library and style strategy to produce translations that better fit the industry. If this field is not specified or an invalid value is passed, the general translation strategy is used.
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

func (s TextTranslateRequest) String() string {
	return dara.Prettify(s)
}

func (s TextTranslateRequest) GoString() string {
	return s.String()
}

func (s *TextTranslateRequest) GetBizName() *string {
	return s.BizName
}

func (s *TextTranslateRequest) GetFormatType() *string {
	return s.FormatType
}

func (s *TextTranslateRequest) GetGlossary() *string {
	return s.Glossary
}

func (s *TextTranslateRequest) GetSourceLanguage() *string {
	return s.SourceLanguage
}

func (s *TextTranslateRequest) GetSourceTextList() []*string {
	return s.SourceTextList
}

func (s *TextTranslateRequest) GetTargetLanguage() *string {
	return s.TargetLanguage
}

func (s *TextTranslateRequest) GetTranslateScene() *string {
	return s.TranslateScene
}

func (s *TextTranslateRequest) SetBizName(v string) *TextTranslateRequest {
	s.BizName = &v
	return s
}

func (s *TextTranslateRequest) SetFormatType(v string) *TextTranslateRequest {
	s.FormatType = &v
	return s
}

func (s *TextTranslateRequest) SetGlossary(v string) *TextTranslateRequest {
	s.Glossary = &v
	return s
}

func (s *TextTranslateRequest) SetSourceLanguage(v string) *TextTranslateRequest {
	s.SourceLanguage = &v
	return s
}

func (s *TextTranslateRequest) SetSourceTextList(v []*string) *TextTranslateRequest {
	s.SourceTextList = v
	return s
}

func (s *TextTranslateRequest) SetTargetLanguage(v string) *TextTranslateRequest {
	s.TargetLanguage = &v
	return s
}

func (s *TextTranslateRequest) SetTranslateScene(v string) *TextTranslateRequest {
	s.TranslateScene = &v
	return s
}

func (s *TextTranslateRequest) Validate() error {
	return dara.Validate(s)
}
