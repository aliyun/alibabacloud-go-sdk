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
	SourceTextList []*string `json:"SourceTextList,omitempty" xml:"SourceTextList,omitempty" type:"Repeated"`
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
