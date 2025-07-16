// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTranslateSearchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFormatType(v string) *TranslateSearchRequest
	GetFormatType() *string
	SetScene(v string) *TranslateSearchRequest
	GetScene() *string
	SetSourceLanguage(v string) *TranslateSearchRequest
	GetSourceLanguage() *string
	SetSourceText(v string) *TranslateSearchRequest
	GetSourceText() *string
	SetTargetLanguage(v string) *TranslateSearchRequest
	GetTargetLanguage() *string
}

type TranslateSearchRequest struct {
	// This parameter is required.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// text
	FormatType *string `json:"FormatType,omitempty" xml:"FormatType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// query
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// This parameter is required.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// zh
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 今天天气不错
	SourceText *string `json:"SourceText,omitempty" xml:"SourceText,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// en
	TargetLanguage *string `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
}

func (s TranslateSearchRequest) String() string {
	return dara.Prettify(s)
}

func (s TranslateSearchRequest) GoString() string {
	return s.String()
}

func (s *TranslateSearchRequest) GetFormatType() *string {
	return s.FormatType
}

func (s *TranslateSearchRequest) GetScene() *string {
	return s.Scene
}

func (s *TranslateSearchRequest) GetSourceLanguage() *string {
	return s.SourceLanguage
}

func (s *TranslateSearchRequest) GetSourceText() *string {
	return s.SourceText
}

func (s *TranslateSearchRequest) GetTargetLanguage() *string {
	return s.TargetLanguage
}

func (s *TranslateSearchRequest) SetFormatType(v string) *TranslateSearchRequest {
	s.FormatType = &v
	return s
}

func (s *TranslateSearchRequest) SetScene(v string) *TranslateSearchRequest {
	s.Scene = &v
	return s
}

func (s *TranslateSearchRequest) SetSourceLanguage(v string) *TranslateSearchRequest {
	s.SourceLanguage = &v
	return s
}

func (s *TranslateSearchRequest) SetSourceText(v string) *TranslateSearchRequest {
	s.SourceText = &v
	return s
}

func (s *TranslateSearchRequest) SetTargetLanguage(v string) *TranslateSearchRequest {
	s.TargetLanguage = &v
	return s
}

func (s *TranslateSearchRequest) Validate() error {
	return dara.Validate(s)
}
