// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPrepaidTextTranslateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizName(v string) *PrepaidTextTranslateRequest
	GetBizName() *string
	SetFormatType(v string) *PrepaidTextTranslateRequest
	GetFormatType() *string
	SetGlossary(v string) *PrepaidTextTranslateRequest
	GetGlossary() *string
	SetSourceLanguage(v string) *PrepaidTextTranslateRequest
	GetSourceLanguage() *string
	SetSourceTextList(v []*string) *PrepaidTextTranslateRequest
	GetSourceTextList() []*string
	SetTargetLanguage(v string) *PrepaidTextTranslateRequest
	GetTargetLanguage() *string
	SetTranslateScene(v string) *PrepaidTextTranslateRequest
	GetTranslateScene() *string
}

type PrepaidTextTranslateRequest struct {
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
	SourceTextList []*string `json:"SourceTextList,omitempty" xml:"SourceTextList,omitempty" type:"Repeated"`
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

func (s PrepaidTextTranslateRequest) String() string {
	return dara.Prettify(s)
}

func (s PrepaidTextTranslateRequest) GoString() string {
	return s.String()
}

func (s *PrepaidTextTranslateRequest) GetBizName() *string {
	return s.BizName
}

func (s *PrepaidTextTranslateRequest) GetFormatType() *string {
	return s.FormatType
}

func (s *PrepaidTextTranslateRequest) GetGlossary() *string {
	return s.Glossary
}

func (s *PrepaidTextTranslateRequest) GetSourceLanguage() *string {
	return s.SourceLanguage
}

func (s *PrepaidTextTranslateRequest) GetSourceTextList() []*string {
	return s.SourceTextList
}

func (s *PrepaidTextTranslateRequest) GetTargetLanguage() *string {
	return s.TargetLanguage
}

func (s *PrepaidTextTranslateRequest) GetTranslateScene() *string {
	return s.TranslateScene
}

func (s *PrepaidTextTranslateRequest) SetBizName(v string) *PrepaidTextTranslateRequest {
	s.BizName = &v
	return s
}

func (s *PrepaidTextTranslateRequest) SetFormatType(v string) *PrepaidTextTranslateRequest {
	s.FormatType = &v
	return s
}

func (s *PrepaidTextTranslateRequest) SetGlossary(v string) *PrepaidTextTranslateRequest {
	s.Glossary = &v
	return s
}

func (s *PrepaidTextTranslateRequest) SetSourceLanguage(v string) *PrepaidTextTranslateRequest {
	s.SourceLanguage = &v
	return s
}

func (s *PrepaidTextTranslateRequest) SetSourceTextList(v []*string) *PrepaidTextTranslateRequest {
	s.SourceTextList = v
	return s
}

func (s *PrepaidTextTranslateRequest) SetTargetLanguage(v string) *PrepaidTextTranslateRequest {
	s.TargetLanguage = &v
	return s
}

func (s *PrepaidTextTranslateRequest) SetTranslateScene(v string) *PrepaidTextTranslateRequest {
	s.TranslateScene = &v
	return s
}

func (s *PrepaidTextTranslateRequest) Validate() error {
	return dara.Validate(s)
}
