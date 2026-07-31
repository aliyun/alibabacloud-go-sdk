// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPrepaidTextTranslateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizName(v string) *PrepaidTextTranslateShrinkRequest
	GetBizName() *string
	SetFormatType(v string) *PrepaidTextTranslateShrinkRequest
	GetFormatType() *string
	SetGlossary(v string) *PrepaidTextTranslateShrinkRequest
	GetGlossary() *string
	SetSourceLanguage(v string) *PrepaidTextTranslateShrinkRequest
	GetSourceLanguage() *string
	SetSourceTextListShrink(v string) *PrepaidTextTranslateShrinkRequest
	GetSourceTextListShrink() *string
	SetTargetLanguage(v string) *PrepaidTextTranslateShrinkRequest
	GetTargetLanguage() *string
	SetTranslateScene(v string) *PrepaidTextTranslateShrinkRequest
	GetTranslateScene() *string
}

type PrepaidTextTranslateShrinkRequest struct {
	// The business scenario identifier. This parameter is optional. Valid values: e-commerce-title, e-commerce-description, e-commerce-chat, e-commerce-cpv, novel, game. If not specified or an invalid value is passed, the general translation strategy is used by default.
	//
	// example:
	//
	// Alibaba-商品
	BizName *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	// The format type of the source text. This parameter is optional. Supports text (plain text format) and html (web page format, preserving HTML tags).
	//
	// example:
	//
	// text
	FormatType *string `json:"FormatType,omitempty" xml:"FormatType,omitempty"`
	// The intervention glossary ID. This parameter is optional. The glossary must be created separately in the console, and its ID must be provided. If the glossary ID is empty, the translation result is not modified.
	//
	// example:
	//
	// custom_glossary
	Glossary *string `json:"Glossary,omitempty" xml:"Glossary,omitempty"`
	// The source language code. This parameter is optional. If not specified, the language is automatically detected. You can pass auto for language detection.
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
	// The target language code. This parameter is required. More than 100 language directions are supported. For details, refer to the supported language directions list.
	//
	// This parameter is required.
	//
	// example:
	//
	// zh
	TargetLanguage *string `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
	// The format of the translation text. html (web page format. This setting processes both the source text and translated text in HTML format) or text (text format. This setting processes both the source text and translated result as plain text without format processing).
	//
	// example:
	//
	// e-commerce-title
	TranslateScene *string `json:"TranslateScene,omitempty" xml:"TranslateScene,omitempty"`
}

func (s PrepaidTextTranslateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s PrepaidTextTranslateShrinkRequest) GoString() string {
	return s.String()
}

func (s *PrepaidTextTranslateShrinkRequest) GetBizName() *string {
	return s.BizName
}

func (s *PrepaidTextTranslateShrinkRequest) GetFormatType() *string {
	return s.FormatType
}

func (s *PrepaidTextTranslateShrinkRequest) GetGlossary() *string {
	return s.Glossary
}

func (s *PrepaidTextTranslateShrinkRequest) GetSourceLanguage() *string {
	return s.SourceLanguage
}

func (s *PrepaidTextTranslateShrinkRequest) GetSourceTextListShrink() *string {
	return s.SourceTextListShrink
}

func (s *PrepaidTextTranslateShrinkRequest) GetTargetLanguage() *string {
	return s.TargetLanguage
}

func (s *PrepaidTextTranslateShrinkRequest) GetTranslateScene() *string {
	return s.TranslateScene
}

func (s *PrepaidTextTranslateShrinkRequest) SetBizName(v string) *PrepaidTextTranslateShrinkRequest {
	s.BizName = &v
	return s
}

func (s *PrepaidTextTranslateShrinkRequest) SetFormatType(v string) *PrepaidTextTranslateShrinkRequest {
	s.FormatType = &v
	return s
}

func (s *PrepaidTextTranslateShrinkRequest) SetGlossary(v string) *PrepaidTextTranslateShrinkRequest {
	s.Glossary = &v
	return s
}

func (s *PrepaidTextTranslateShrinkRequest) SetSourceLanguage(v string) *PrepaidTextTranslateShrinkRequest {
	s.SourceLanguage = &v
	return s
}

func (s *PrepaidTextTranslateShrinkRequest) SetSourceTextListShrink(v string) *PrepaidTextTranslateShrinkRequest {
	s.SourceTextListShrink = &v
	return s
}

func (s *PrepaidTextTranslateShrinkRequest) SetTargetLanguage(v string) *PrepaidTextTranslateShrinkRequest {
	s.TargetLanguage = &v
	return s
}

func (s *PrepaidTextTranslateShrinkRequest) SetTranslateScene(v string) *PrepaidTextTranslateShrinkRequest {
	s.TranslateScene = &v
	return s
}

func (s *PrepaidTextTranslateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
