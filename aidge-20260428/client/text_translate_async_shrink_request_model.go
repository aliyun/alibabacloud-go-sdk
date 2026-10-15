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
	// example:
	//
	// order-20260928-001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// BUxxx-chat
	BizName *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	// example:
	//
	// text
	FormatType *string `json:"FormatType,omitempty" xml:"FormatType,omitempty"`
	// example:
	//
	// " glossary_1"
	Glossary *string `json:"Glossary,omitempty" xml:"Glossary,omitempty"`
	// example:
	//
	// auto
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	// This parameter is required.
	SourceTextListShrink *string `json:"SourceTextList,omitempty" xml:"SourceTextList,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// en
	TargetLanguage *string `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
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
