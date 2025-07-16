// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTranslateGeneralVpcRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContext(v string) *TranslateGeneralVpcRequest
	GetContext() *string
	SetFormatType(v string) *TranslateGeneralVpcRequest
	GetFormatType() *string
	SetScene(v string) *TranslateGeneralVpcRequest
	GetScene() *string
	SetSourceLanguage(v string) *TranslateGeneralVpcRequest
	GetSourceLanguage() *string
	SetSourceText(v string) *TranslateGeneralVpcRequest
	GetSourceText() *string
	SetTargetLanguage(v string) *TranslateGeneralVpcRequest
	GetTargetLanguage() *string
}

type TranslateGeneralVpcRequest struct {
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

func (s TranslateGeneralVpcRequest) String() string {
	return dara.Prettify(s)
}

func (s TranslateGeneralVpcRequest) GoString() string {
	return s.String()
}

func (s *TranslateGeneralVpcRequest) GetContext() *string {
	return s.Context
}

func (s *TranslateGeneralVpcRequest) GetFormatType() *string {
	return s.FormatType
}

func (s *TranslateGeneralVpcRequest) GetScene() *string {
	return s.Scene
}

func (s *TranslateGeneralVpcRequest) GetSourceLanguage() *string {
	return s.SourceLanguage
}

func (s *TranslateGeneralVpcRequest) GetSourceText() *string {
	return s.SourceText
}

func (s *TranslateGeneralVpcRequest) GetTargetLanguage() *string {
	return s.TargetLanguage
}

func (s *TranslateGeneralVpcRequest) SetContext(v string) *TranslateGeneralVpcRequest {
	s.Context = &v
	return s
}

func (s *TranslateGeneralVpcRequest) SetFormatType(v string) *TranslateGeneralVpcRequest {
	s.FormatType = &v
	return s
}

func (s *TranslateGeneralVpcRequest) SetScene(v string) *TranslateGeneralVpcRequest {
	s.Scene = &v
	return s
}

func (s *TranslateGeneralVpcRequest) SetSourceLanguage(v string) *TranslateGeneralVpcRequest {
	s.SourceLanguage = &v
	return s
}

func (s *TranslateGeneralVpcRequest) SetSourceText(v string) *TranslateGeneralVpcRequest {
	s.SourceText = &v
	return s
}

func (s *TranslateGeneralVpcRequest) SetTargetLanguage(v string) *TranslateGeneralVpcRequest {
	s.TargetLanguage = &v
	return s
}

func (s *TranslateGeneralVpcRequest) Validate() error {
	return dara.Validate(s)
}
