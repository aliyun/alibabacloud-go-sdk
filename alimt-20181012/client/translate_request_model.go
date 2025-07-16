// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTranslateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContext(v string) *TranslateRequest
	GetContext() *string
	SetFormatType(v string) *TranslateRequest
	GetFormatType() *string
	SetScene(v string) *TranslateRequest
	GetScene() *string
	SetSourceLanguage(v string) *TranslateRequest
	GetSourceLanguage() *string
	SetSourceText(v string) *TranslateRequest
	GetSourceText() *string
	SetTargetLanguage(v string) *TranslateRequest
	GetTargetLanguage() *string
}

type TranslateRequest struct {
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
	// This parameter is required.
	//
	// example:
	//
	// title
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

func (s TranslateRequest) String() string {
	return dara.Prettify(s)
}

func (s TranslateRequest) GoString() string {
	return s.String()
}

func (s *TranslateRequest) GetContext() *string {
	return s.Context
}

func (s *TranslateRequest) GetFormatType() *string {
	return s.FormatType
}

func (s *TranslateRequest) GetScene() *string {
	return s.Scene
}

func (s *TranslateRequest) GetSourceLanguage() *string {
	return s.SourceLanguage
}

func (s *TranslateRequest) GetSourceText() *string {
	return s.SourceText
}

func (s *TranslateRequest) GetTargetLanguage() *string {
	return s.TargetLanguage
}

func (s *TranslateRequest) SetContext(v string) *TranslateRequest {
	s.Context = &v
	return s
}

func (s *TranslateRequest) SetFormatType(v string) *TranslateRequest {
	s.FormatType = &v
	return s
}

func (s *TranslateRequest) SetScene(v string) *TranslateRequest {
	s.Scene = &v
	return s
}

func (s *TranslateRequest) SetSourceLanguage(v string) *TranslateRequest {
	s.SourceLanguage = &v
	return s
}

func (s *TranslateRequest) SetSourceText(v string) *TranslateRequest {
	s.SourceText = &v
	return s
}

func (s *TranslateRequest) SetTargetLanguage(v string) *TranslateRequest {
	s.TargetLanguage = &v
	return s
}

func (s *TranslateRequest) Validate() error {
	return dara.Validate(s)
}
