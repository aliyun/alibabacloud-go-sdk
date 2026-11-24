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
	// This field represents your identity and facilitates communication for various issues. If you are an internal Alibaba organization, we recommend that you specify this parameter based on the actual scenario, such as BU name-product or BU name-chat. If you are an external Alibaba partner, we recommend that you specify the full name of your company. The company name must be the same as the company name you used when registering your Alibaba Cloud account.
	//
	// example:
	//
	// MyCompany-Chat
	BizName *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	// The format type of the source text. This parameter is optional. Valid values:
	//
	// - text: Plain text format.
	//
	// - html: Web page format, which retains HTML tags.
	//
	// example:
	//
	// text
	FormatType *string `json:"FormatType,omitempty" xml:"FormatType,omitempty"`
	// The ID of the intervention glossary. This parameter is optional. You must create the glossary in the console and provide its ID. If the provided intervention glossary ID is empty, the translation results are not modified.
	//
	// example:
	//
	// glossary_1
	Glossary *string `json:"Glossary,omitempty" xml:"Glossary,omitempty"`
	// The source language code. If this parameter is not specified, the language is automatically detected. This parameter is optional. You can set this parameter to auto for language identification.
	//
	// <props="china">For supported language pairs, see [Language pair mapping table](https://www.alibabacloud.com/help/en/document_detail/3041883.html).
	//
	// <props="intl">For supported language pairs, see [Language pair mapping table](https://www.alibabacloud.com/help/zh/aidge/language-direction-mapping).
	//
	// example:
	//
	// auto
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	// The list of text to be translated. This parameter is required. The total character length cannot exceed 50,000, and the list length cannot exceed 50.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["Hello world"]
	SourceTextList []*string `json:"SourceTextList,omitempty" xml:"SourceTextList,omitempty" type:"Repeated"`
	// The target language code. This parameter is required.
	//
	// <props="china">For supported language pairs, see [Language pair mapping table](https://www.alibabacloud.com/help/en/document_detail/3041883.html).
	//
	// <props="intl">For supported language pairs, see [Language pair mapping table](https://www.alibabacloud.com/help/zh/aidge/language-direction-mapping).
	//
	// This parameter is required.
	//
	// example:
	//
	// ko
	TargetLanguage *string `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
	// The business scenario identifier. You can specify only one of the following values. After the value is specified, the translation engine calls the terminology library and style policy of the corresponding industry to output translations that better fit the industry. If this parameter is not specified or an invalid value is specified, the general translation policy is used. Valid values:
	//
	// - e-commerce-title: Cross-border e-commerce product title translation.
	//
	// - e-commerce-description: Cross-border e-commerce product description translation.
	//
	// - e-commerce-chat: Cross-border e-commerce conversation translation.
	//
	// - e-commerce-cpv: Cross-border e-commerce product CPV attribute translation.
	//
	// - novel: Novel translation.
	//
	// - game: Game translation.
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
