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
	// The business scenario identifier. Optional. Valid values: e-commerce-title, e-commerce-description, e-commerce-chat, e-commerce-cpv, novel, game. If not specified or invalid, the general translation strategy is used by default.
	//
	// example:
	//
	// MyCompany-Chat
	BizName *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	// The format type of the source text. Optional. Supports text (plain text format) and html (web page format, preserving HTML tags).
	//
	// example:
	//
	// text
	FormatType *string `json:"FormatType,omitempty" xml:"FormatType,omitempty"`
	// The intervention glossary ID. Optional. The glossary must be created separately in the console and its ID provided. If the glossary ID is empty, the translation result is not modified.
	//
	// example:
	//
	// glossary_1
	Glossary *string `json:"Glossary,omitempty" xml:"Glossary,omitempty"`
	// The source language code. Optional. If not specified, the language is automatically detected. You can pass auto for language detection.
	//
	// example:
	//
	// auto
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	// The list of texts to translate. Required. The total character length cannot exceed 50,000, and the list length cannot exceed 50.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["Hello world"]
	SourceTextList []*string `json:"SourceTextList,omitempty" xml:"SourceTextList,omitempty" type:"Repeated"`
	// The target language code. Required. Supports more than 100 language directions. For details, refer to the supported language directions list.
	//
	// This parameter is required.
	//
	// example:
	//
	// ko
	TargetLanguage *string `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
	// The format of the translation text. **html*	- (web page format. This setting processes both the source text and translated text in HTML format). **text*	- (text format. This setting processes both the source text and translated result as plain text without format processing).
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
