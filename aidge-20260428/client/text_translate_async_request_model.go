// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTextTranslateAsyncRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *TextTranslateAsyncRequest
	GetBizId() *string
	SetBizName(v string) *TextTranslateAsyncRequest
	GetBizName() *string
	SetFormatType(v string) *TextTranslateAsyncRequest
	GetFormatType() *string
	SetGlossary(v string) *TextTranslateAsyncRequest
	GetGlossary() *string
	SetSourceLanguage(v string) *TextTranslateAsyncRequest
	GetSourceLanguage() *string
	SetSourceTextList(v []*string) *TextTranslateAsyncRequest
	GetSourceTextList() []*string
	SetTargetLanguage(v string) *TextTranslateAsyncRequest
	GetTargetLanguage() *string
	SetTranslateScene(v string) *TextTranslateAsyncRequest
	GetTranslateScene() *string
}

type TextTranslateAsyncRequest struct {
	// The business ID of the caller. The maximum length is 128 characters. This ID is passed through in the translation results.
	//
	// example:
	//
	// order-20260928-001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// The business name of the caller. This field represents your identity and facilitates subsequent communication. If you are an internal Alibaba organization, specify this parameter based on the actual scenario, such as BU name-product or BU name-chat. If you are an external partner, specify the full name of your company. The company name must be the same as the one used when registering your Alibaba Cloud account.
	//
	// example:
	//
	// BUxxx-chat
	BizName *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	// The format type of the source text. Valid values: text and html.
	//
	// example:
	//
	// text
	FormatType *string `json:"FormatType,omitempty" xml:"FormatType,omitempty"`
	// The ID of the intervention glossary. You must create the glossary in the console and provide its ID in the request parameters. If the intervention glossary ID is empty, the translation results are not modified.
	//
	// example:
	//
	// " glossary_1"
	Glossary *string `json:"Glossary,omitempty" xml:"Glossary,omitempty"`
	// The source language code. For more information, see the supported language list below. If this parameter is not specified, the language is automatically detected.
	//
	// example:
	//
	// auto
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	// The list of texts to be translated.
	//
	// This parameter is required.
	SourceTextList []*string `json:"SourceTextList,omitempty" xml:"SourceTextList,omitempty" type:"Repeated"`
	// The target language code. For more information, see the supported language pair list below.
	//
	// This parameter is required.
	//
	// example:
	//
	// en
	TargetLanguage *string `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
	// The business scenario identifier. You can specify only one of the following values. After you specify a value, the translation engine uses the terminology library and style policy of the corresponding industry to output industry-specific translations. If this parameter is not specified or an invalid value is specified, the general translation policy is used. Valid values:
	//
	// - e-commerce-title: Translates cross-border e-commerce product titles.
	//
	// - e-commerce-description: Translates cross-border e-commerce product descriptions.
	//
	// - e-commerce-chat: Translates cross-border e-commerce conversations.
	//
	// - e-commerce-cpv: Translates cross-border e-commerce product CPV attributes.
	//
	// - novel: Translates novels.
	//
	// - game: Translates games.
	//
	// example:
	//
	// e-commerce-title
	TranslateScene *string `json:"TranslateScene,omitempty" xml:"TranslateScene,omitempty"`
}

func (s TextTranslateAsyncRequest) String() string {
	return dara.Prettify(s)
}

func (s TextTranslateAsyncRequest) GoString() string {
	return s.String()
}

func (s *TextTranslateAsyncRequest) GetBizId() *string {
	return s.BizId
}

func (s *TextTranslateAsyncRequest) GetBizName() *string {
	return s.BizName
}

func (s *TextTranslateAsyncRequest) GetFormatType() *string {
	return s.FormatType
}

func (s *TextTranslateAsyncRequest) GetGlossary() *string {
	return s.Glossary
}

func (s *TextTranslateAsyncRequest) GetSourceLanguage() *string {
	return s.SourceLanguage
}

func (s *TextTranslateAsyncRequest) GetSourceTextList() []*string {
	return s.SourceTextList
}

func (s *TextTranslateAsyncRequest) GetTargetLanguage() *string {
	return s.TargetLanguage
}

func (s *TextTranslateAsyncRequest) GetTranslateScene() *string {
	return s.TranslateScene
}

func (s *TextTranslateAsyncRequest) SetBizId(v string) *TextTranslateAsyncRequest {
	s.BizId = &v
	return s
}

func (s *TextTranslateAsyncRequest) SetBizName(v string) *TextTranslateAsyncRequest {
	s.BizName = &v
	return s
}

func (s *TextTranslateAsyncRequest) SetFormatType(v string) *TextTranslateAsyncRequest {
	s.FormatType = &v
	return s
}

func (s *TextTranslateAsyncRequest) SetGlossary(v string) *TextTranslateAsyncRequest {
	s.Glossary = &v
	return s
}

func (s *TextTranslateAsyncRequest) SetSourceLanguage(v string) *TextTranslateAsyncRequest {
	s.SourceLanguage = &v
	return s
}

func (s *TextTranslateAsyncRequest) SetSourceTextList(v []*string) *TextTranslateAsyncRequest {
	s.SourceTextList = v
	return s
}

func (s *TextTranslateAsyncRequest) SetTargetLanguage(v string) *TextTranslateAsyncRequest {
	s.TargetLanguage = &v
	return s
}

func (s *TextTranslateAsyncRequest) SetTranslateScene(v string) *TextTranslateAsyncRequest {
	s.TranslateScene = &v
	return s
}

func (s *TextTranslateAsyncRequest) Validate() error {
	return dara.Validate(s)
}
