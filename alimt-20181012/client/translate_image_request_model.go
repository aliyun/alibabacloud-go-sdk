// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTranslateImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExt(v string) *TranslateImageRequest
	GetExt() *string
	SetField(v string) *TranslateImageRequest
	GetField() *string
	SetImageBase64(v string) *TranslateImageRequest
	GetImageBase64() *string
	SetImageUrl(v string) *TranslateImageRequest
	GetImageUrl() *string
	SetSourceLanguage(v string) *TranslateImageRequest
	GetSourceLanguage() *string
	SetTargetLanguage(v string) *TranslateImageRequest
	GetTargetLanguage() *string
}

type TranslateImageRequest struct {
	// example:
	//
	// {"needEditorData": "false", "ignoreEntityRecognize": "true"}
	Ext *string `json:"Ext,omitempty" xml:"Ext,omitempty"`
	// example:
	//
	// general
	Field *string `json:"Field,omitempty" xml:"Field,omitempty"`
	// example:
	//
	// /9j/4...H/9k=
	ImageBase64 *string `json:"ImageBase64,omitempty" xml:"ImageBase64,omitempty"`
	// example:
	//
	// https://example.com/example.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// example:
	//
	// zh
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	// example:
	//
	// en
	TargetLanguage *string `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
}

func (s TranslateImageRequest) String() string {
	return dara.Prettify(s)
}

func (s TranslateImageRequest) GoString() string {
	return s.String()
}

func (s *TranslateImageRequest) GetExt() *string {
	return s.Ext
}

func (s *TranslateImageRequest) GetField() *string {
	return s.Field
}

func (s *TranslateImageRequest) GetImageBase64() *string {
	return s.ImageBase64
}

func (s *TranslateImageRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *TranslateImageRequest) GetSourceLanguage() *string {
	return s.SourceLanguage
}

func (s *TranslateImageRequest) GetTargetLanguage() *string {
	return s.TargetLanguage
}

func (s *TranslateImageRequest) SetExt(v string) *TranslateImageRequest {
	s.Ext = &v
	return s
}

func (s *TranslateImageRequest) SetField(v string) *TranslateImageRequest {
	s.Field = &v
	return s
}

func (s *TranslateImageRequest) SetImageBase64(v string) *TranslateImageRequest {
	s.ImageBase64 = &v
	return s
}

func (s *TranslateImageRequest) SetImageUrl(v string) *TranslateImageRequest {
	s.ImageUrl = &v
	return s
}

func (s *TranslateImageRequest) SetSourceLanguage(v string) *TranslateImageRequest {
	s.SourceLanguage = &v
	return s
}

func (s *TranslateImageRequest) SetTargetLanguage(v string) *TranslateImageRequest {
	s.TargetLanguage = &v
	return s
}

func (s *TranslateImageRequest) Validate() error {
	return dara.Validate(s)
}
