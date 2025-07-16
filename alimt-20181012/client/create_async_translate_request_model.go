// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAsyncTranslateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiType(v string) *CreateAsyncTranslateRequest
	GetApiType() *string
	SetFormatType(v string) *CreateAsyncTranslateRequest
	GetFormatType() *string
	SetScene(v string) *CreateAsyncTranslateRequest
	GetScene() *string
	SetSourceLanguage(v string) *CreateAsyncTranslateRequest
	GetSourceLanguage() *string
	SetSourceText(v string) *CreateAsyncTranslateRequest
	GetSourceText() *string
	SetTargetLanguage(v string) *CreateAsyncTranslateRequest
	GetTargetLanguage() *string
}

type CreateAsyncTranslateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// translate_standard
	ApiType *string `json:"ApiType,omitempty" xml:"ApiType,omitempty"`
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
	// en
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Opinion: We have finally gotten some relief at the pump. But it may not last long
	SourceText *string `json:"SourceText,omitempty" xml:"SourceText,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// zh
	TargetLanguage *string `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
}

func (s CreateAsyncTranslateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAsyncTranslateRequest) GoString() string {
	return s.String()
}

func (s *CreateAsyncTranslateRequest) GetApiType() *string {
	return s.ApiType
}

func (s *CreateAsyncTranslateRequest) GetFormatType() *string {
	return s.FormatType
}

func (s *CreateAsyncTranslateRequest) GetScene() *string {
	return s.Scene
}

func (s *CreateAsyncTranslateRequest) GetSourceLanguage() *string {
	return s.SourceLanguage
}

func (s *CreateAsyncTranslateRequest) GetSourceText() *string {
	return s.SourceText
}

func (s *CreateAsyncTranslateRequest) GetTargetLanguage() *string {
	return s.TargetLanguage
}

func (s *CreateAsyncTranslateRequest) SetApiType(v string) *CreateAsyncTranslateRequest {
	s.ApiType = &v
	return s
}

func (s *CreateAsyncTranslateRequest) SetFormatType(v string) *CreateAsyncTranslateRequest {
	s.FormatType = &v
	return s
}

func (s *CreateAsyncTranslateRequest) SetScene(v string) *CreateAsyncTranslateRequest {
	s.Scene = &v
	return s
}

func (s *CreateAsyncTranslateRequest) SetSourceLanguage(v string) *CreateAsyncTranslateRequest {
	s.SourceLanguage = &v
	return s
}

func (s *CreateAsyncTranslateRequest) SetSourceText(v string) *CreateAsyncTranslateRequest {
	s.SourceText = &v
	return s
}

func (s *CreateAsyncTranslateRequest) SetTargetLanguage(v string) *CreateAsyncTranslateRequest {
	s.TargetLanguage = &v
	return s
}

func (s *CreateAsyncTranslateRequest) Validate() error {
	return dara.Validate(s)
}
