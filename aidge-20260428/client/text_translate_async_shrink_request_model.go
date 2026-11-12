// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTextTranslateAsyncShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *TextTranslateAsyncShrinkRequest
	GetBizId() *string
	SetBizName(v string) *TextTranslateAsyncShrinkRequest
	GetBizName() *string
	SetFormatType(v string) *TextTranslateAsyncShrinkRequest
	GetFormatType() *string
	SetGlossary(v string) *TextTranslateAsyncShrinkRequest
	GetGlossary() *string
	SetSourceLanguage(v string) *TextTranslateAsyncShrinkRequest
	GetSourceLanguage() *string
	SetSourceTextListShrink(v string) *TextTranslateAsyncShrinkRequest
	GetSourceTextListShrink() *string
	SetTargetLanguage(v string) *TextTranslateAsyncShrinkRequest
	GetTargetLanguage() *string
	SetTranslateScene(v string) *TextTranslateAsyncShrinkRequest
	GetTranslateScene() *string
}

type TextTranslateAsyncShrinkRequest struct {
	// The caller business ID. The ID can be up to 128 characters in length and is passed through in the translation results. This parameter is used only for asynchronous text translation and is returned as is in the task results.
	//
	// example:
	//
	// order-20260928-001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// This field represents your identity and facilitates communication for various issues. If you are an internal Alibaba organization, we recommend that you specify this field based on the actual scenario, such as BU name-product or BU name-chat. If you are an external Alibaba partner, we recommend that you specify the full name of your company. The company name must be the same as the one you used when registering your Alibaba Cloud account.
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
	// The ID of the intervention glossary. You must create the glossary in the console and provide its ID in the input parameters. If the intervention glossary ID is empty, the translation results are not modified.
	//
	// example:
	//
	// " glossary_1"
	Glossary *string `json:"Glossary,omitempty" xml:"Glossary,omitempty"`
	// The source language code. For more information, see the supported language list below. If you do not specify this parameter, the source language is automatically detected.
	//
	// example:
	//
	// auto
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	// The list of text to be translated.
	//
	// This parameter is required.
	SourceTextListShrink *string `json:"SourceTextList,omitempty" xml:"SourceTextList,omitempty"`
	// The target language code. For more information, see the supported language pair list below.
	//
	// This parameter is required.
	//
	// example:
	//
	// en
	TargetLanguage *string `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
	// The business scenario identifier. You can specify only one of the following values. After you specify a value, the translation engine uses the terminology library and style policy of the corresponding industry to output industry-specific translations. If this field is not specified or is invalid, the general translation policy is used. Valid values:
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

func (s TextTranslateAsyncShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s TextTranslateAsyncShrinkRequest) GoString() string {
	return s.String()
}

func (s *TextTranslateAsyncShrinkRequest) GetBizId() *string {
	return s.BizId
}

func (s *TextTranslateAsyncShrinkRequest) GetBizName() *string {
	return s.BizName
}

func (s *TextTranslateAsyncShrinkRequest) GetFormatType() *string {
	return s.FormatType
}

func (s *TextTranslateAsyncShrinkRequest) GetGlossary() *string {
	return s.Glossary
}

func (s *TextTranslateAsyncShrinkRequest) GetSourceLanguage() *string {
	return s.SourceLanguage
}

func (s *TextTranslateAsyncShrinkRequest) GetSourceTextListShrink() *string {
	return s.SourceTextListShrink
}

func (s *TextTranslateAsyncShrinkRequest) GetTargetLanguage() *string {
	return s.TargetLanguage
}

func (s *TextTranslateAsyncShrinkRequest) GetTranslateScene() *string {
	return s.TranslateScene
}

func (s *TextTranslateAsyncShrinkRequest) SetBizId(v string) *TextTranslateAsyncShrinkRequest {
	s.BizId = &v
	return s
}

func (s *TextTranslateAsyncShrinkRequest) SetBizName(v string) *TextTranslateAsyncShrinkRequest {
	s.BizName = &v
	return s
}

func (s *TextTranslateAsyncShrinkRequest) SetFormatType(v string) *TextTranslateAsyncShrinkRequest {
	s.FormatType = &v
	return s
}

func (s *TextTranslateAsyncShrinkRequest) SetGlossary(v string) *TextTranslateAsyncShrinkRequest {
	s.Glossary = &v
	return s
}

func (s *TextTranslateAsyncShrinkRequest) SetSourceLanguage(v string) *TextTranslateAsyncShrinkRequest {
	s.SourceLanguage = &v
	return s
}

func (s *TextTranslateAsyncShrinkRequest) SetSourceTextListShrink(v string) *TextTranslateAsyncShrinkRequest {
	s.SourceTextListShrink = &v
	return s
}

func (s *TextTranslateAsyncShrinkRequest) SetTargetLanguage(v string) *TextTranslateAsyncShrinkRequest {
	s.TargetLanguage = &v
	return s
}

func (s *TextTranslateAsyncShrinkRequest) SetTranslateScene(v string) *TextTranslateAsyncShrinkRequest {
	s.TranslateScene = &v
	return s
}

func (s *TextTranslateAsyncShrinkRequest) Validate() error {
	return dara.Validate(s)
}
