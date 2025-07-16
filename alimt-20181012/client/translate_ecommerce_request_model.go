// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTranslateECommerceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContext(v string) *TranslateECommerceRequest
	GetContext() *string
	SetFormatType(v string) *TranslateECommerceRequest
	GetFormatType() *string
	SetScene(v string) *TranslateECommerceRequest
	GetScene() *string
	SetSourceLanguage(v string) *TranslateECommerceRequest
	GetSourceLanguage() *string
	SetSourceText(v string) *TranslateECommerceRequest
	GetSourceText() *string
	SetTargetLanguage(v string) *TranslateECommerceRequest
	GetTargetLanguage() *string
}

type TranslateECommerceRequest struct {
	// example:
	//
	// context信息
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
	// social
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// en
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Hello
	SourceText *string `json:"SourceText,omitempty" xml:"SourceText,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// zh
	TargetLanguage *string `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
}

func (s TranslateECommerceRequest) String() string {
	return dara.Prettify(s)
}

func (s TranslateECommerceRequest) GoString() string {
	return s.String()
}

func (s *TranslateECommerceRequest) GetContext() *string {
	return s.Context
}

func (s *TranslateECommerceRequest) GetFormatType() *string {
	return s.FormatType
}

func (s *TranslateECommerceRequest) GetScene() *string {
	return s.Scene
}

func (s *TranslateECommerceRequest) GetSourceLanguage() *string {
	return s.SourceLanguage
}

func (s *TranslateECommerceRequest) GetSourceText() *string {
	return s.SourceText
}

func (s *TranslateECommerceRequest) GetTargetLanguage() *string {
	return s.TargetLanguage
}

func (s *TranslateECommerceRequest) SetContext(v string) *TranslateECommerceRequest {
	s.Context = &v
	return s
}

func (s *TranslateECommerceRequest) SetFormatType(v string) *TranslateECommerceRequest {
	s.FormatType = &v
	return s
}

func (s *TranslateECommerceRequest) SetScene(v string) *TranslateECommerceRequest {
	s.Scene = &v
	return s
}

func (s *TranslateECommerceRequest) SetSourceLanguage(v string) *TranslateECommerceRequest {
	s.SourceLanguage = &v
	return s
}

func (s *TranslateECommerceRequest) SetSourceText(v string) *TranslateECommerceRequest {
	s.SourceText = &v
	return s
}

func (s *TranslateECommerceRequest) SetTargetLanguage(v string) *TranslateECommerceRequest {
	s.TargetLanguage = &v
	return s
}

func (s *TranslateECommerceRequest) Validate() error {
	return dara.Validate(s)
}
