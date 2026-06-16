// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTextTranslateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
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
}

type TextTranslateShrinkRequest struct {
	// The format type of the source text. Optional. Valid values: text (plain text format) and html (web page format that preserves HTML tags).
	//
	// example:
	//
	// text
	FormatType *string `json:"FormatType,omitempty" xml:"FormatType,omitempty"`
	// The intervention glossary ID. Optional. Create the glossary in the console and provide its ID. If the glossary ID is empty, the translation results are not modified.
	//
	// example:
	//
	// glossary_1
	Glossary *string `json:"Glossary,omitempty" xml:"Glossary,omitempty"`
	// The source language code. Optional. If not specified, the language is automatically detected. Set to auto for automatic language detection.
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
	SourceTextListShrink *string `json:"SourceTextList,omitempty" xml:"SourceTextList,omitempty"`
	// The target language code. Required. More than 100 language directions are supported. For details, refer to the supported language directions list.
	//
	// This parameter is required.
	//
	// example:
	//
	// ko
	TargetLanguage *string `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
}

func (s TextTranslateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s TextTranslateShrinkRequest) GoString() string {
	return s.String()
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

func (s *TextTranslateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
