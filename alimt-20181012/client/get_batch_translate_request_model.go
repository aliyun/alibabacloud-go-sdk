// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBatchTranslateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiType(v string) *GetBatchTranslateRequest
	GetApiType() *string
	SetFormatType(v string) *GetBatchTranslateRequest
	GetFormatType() *string
	SetScene(v string) *GetBatchTranslateRequest
	GetScene() *string
	SetSourceLanguage(v string) *GetBatchTranslateRequest
	GetSourceLanguage() *string
	SetSourceText(v string) *GetBatchTranslateRequest
	GetSourceText() *string
	SetTargetLanguage(v string) *GetBatchTranslateRequest
	GetTargetLanguage() *string
}

type GetBatchTranslateRequest struct {
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
	// {"11":"hello boy","12":"go home","13":"we can"}
	SourceText *string `json:"SourceText,omitempty" xml:"SourceText,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// zh
	TargetLanguage *string `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
}

func (s GetBatchTranslateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBatchTranslateRequest) GoString() string {
	return s.String()
}

func (s *GetBatchTranslateRequest) GetApiType() *string {
	return s.ApiType
}

func (s *GetBatchTranslateRequest) GetFormatType() *string {
	return s.FormatType
}

func (s *GetBatchTranslateRequest) GetScene() *string {
	return s.Scene
}

func (s *GetBatchTranslateRequest) GetSourceLanguage() *string {
	return s.SourceLanguage
}

func (s *GetBatchTranslateRequest) GetSourceText() *string {
	return s.SourceText
}

func (s *GetBatchTranslateRequest) GetTargetLanguage() *string {
	return s.TargetLanguage
}

func (s *GetBatchTranslateRequest) SetApiType(v string) *GetBatchTranslateRequest {
	s.ApiType = &v
	return s
}

func (s *GetBatchTranslateRequest) SetFormatType(v string) *GetBatchTranslateRequest {
	s.FormatType = &v
	return s
}

func (s *GetBatchTranslateRequest) SetScene(v string) *GetBatchTranslateRequest {
	s.Scene = &v
	return s
}

func (s *GetBatchTranslateRequest) SetSourceLanguage(v string) *GetBatchTranslateRequest {
	s.SourceLanguage = &v
	return s
}

func (s *GetBatchTranslateRequest) SetSourceText(v string) *GetBatchTranslateRequest {
	s.SourceText = &v
	return s
}

func (s *GetBatchTranslateRequest) SetTargetLanguage(v string) *GetBatchTranslateRequest {
	s.TargetLanguage = &v
	return s
}

func (s *GetBatchTranslateRequest) Validate() error {
	return dara.Validate(s)
}
