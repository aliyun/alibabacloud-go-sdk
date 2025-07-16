// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTranslateGeneralRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContext(v string) *TranslateGeneralRequest
	GetContext() *string
	SetFormatType(v string) *TranslateGeneralRequest
	GetFormatType() *string
	SetScene(v string) *TranslateGeneralRequest
	GetScene() *string
	SetSourceLanguage(v string) *TranslateGeneralRequest
	GetSourceLanguage() *string
	SetSourceText(v string) *TranslateGeneralRequest
	GetSourceText() *string
	SetTargetLanguage(v string) *TranslateGeneralRequest
	GetTargetLanguage() *string
}

type TranslateGeneralRequest struct {
	// example:
	//
	// {\\"appName\\":\\"alynx\\"}
	Context *string `json:"Context,omitempty" xml:"Context,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// text
	FormatType *string `json:"FormatType,omitempty" xml:"FormatType,omitempty"`
	// example:
	//
	// general
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// zh
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	// This parameter is required.
	SourceText *string `json:"SourceText,omitempty" xml:"SourceText,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// en
	TargetLanguage *string `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
}

func (s TranslateGeneralRequest) String() string {
	return dara.Prettify(s)
}

func (s TranslateGeneralRequest) GoString() string {
	return s.String()
}

func (s *TranslateGeneralRequest) GetContext() *string {
	return s.Context
}

func (s *TranslateGeneralRequest) GetFormatType() *string {
	return s.FormatType
}

func (s *TranslateGeneralRequest) GetScene() *string {
	return s.Scene
}

func (s *TranslateGeneralRequest) GetSourceLanguage() *string {
	return s.SourceLanguage
}

func (s *TranslateGeneralRequest) GetSourceText() *string {
	return s.SourceText
}

func (s *TranslateGeneralRequest) GetTargetLanguage() *string {
	return s.TargetLanguage
}

func (s *TranslateGeneralRequest) SetContext(v string) *TranslateGeneralRequest {
	s.Context = &v
	return s
}

func (s *TranslateGeneralRequest) SetFormatType(v string) *TranslateGeneralRequest {
	s.FormatType = &v
	return s
}

func (s *TranslateGeneralRequest) SetScene(v string) *TranslateGeneralRequest {
	s.Scene = &v
	return s
}

func (s *TranslateGeneralRequest) SetSourceLanguage(v string) *TranslateGeneralRequest {
	s.SourceLanguage = &v
	return s
}

func (s *TranslateGeneralRequest) SetSourceText(v string) *TranslateGeneralRequest {
	s.SourceText = &v
	return s
}

func (s *TranslateGeneralRequest) SetTargetLanguage(v string) *TranslateGeneralRequest {
	s.TargetLanguage = &v
	return s
}

func (s *TranslateGeneralRequest) Validate() error {
	return dara.Validate(s)
}
