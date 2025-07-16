// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTranslateImageBatchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomTaskId(v string) *TranslateImageBatchRequest
	GetCustomTaskId() *string
	SetExt(v string) *TranslateImageBatchRequest
	GetExt() *string
	SetField(v string) *TranslateImageBatchRequest
	GetField() *string
	SetImageUrls(v string) *TranslateImageBatchRequest
	GetImageUrls() *string
	SetSourceLanguage(v string) *TranslateImageBatchRequest
	GetSourceLanguage() *string
	SetTargetLanguage(v string) *TranslateImageBatchRequest
	GetTargetLanguage() *string
}

type TranslateImageBatchRequest struct {
	// example:
	//
	// my_awesome_task_1
	CustomTaskId *string `json:"CustomTaskId,omitempty" xml:"CustomTaskId,omitempty"`
	// example:
	//
	// {"needEditorData": "false", "ignoreEntityRecognize": "true"}
	Ext *string `json:"Ext,omitempty" xml:"Ext,omitempty"`
	// example:
	//
	// general
	Field *string `json:"Field,omitempty" xml:"Field,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/1.jpg,https://example.com/2.jpg,https://example.com/3.jpg
	ImageUrls *string `json:"ImageUrls,omitempty" xml:"ImageUrls,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// zh
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// en
	TargetLanguage *string `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
}

func (s TranslateImageBatchRequest) String() string {
	return dara.Prettify(s)
}

func (s TranslateImageBatchRequest) GoString() string {
	return s.String()
}

func (s *TranslateImageBatchRequest) GetCustomTaskId() *string {
	return s.CustomTaskId
}

func (s *TranslateImageBatchRequest) GetExt() *string {
	return s.Ext
}

func (s *TranslateImageBatchRequest) GetField() *string {
	return s.Field
}

func (s *TranslateImageBatchRequest) GetImageUrls() *string {
	return s.ImageUrls
}

func (s *TranslateImageBatchRequest) GetSourceLanguage() *string {
	return s.SourceLanguage
}

func (s *TranslateImageBatchRequest) GetTargetLanguage() *string {
	return s.TargetLanguage
}

func (s *TranslateImageBatchRequest) SetCustomTaskId(v string) *TranslateImageBatchRequest {
	s.CustomTaskId = &v
	return s
}

func (s *TranslateImageBatchRequest) SetExt(v string) *TranslateImageBatchRequest {
	s.Ext = &v
	return s
}

func (s *TranslateImageBatchRequest) SetField(v string) *TranslateImageBatchRequest {
	s.Field = &v
	return s
}

func (s *TranslateImageBatchRequest) SetImageUrls(v string) *TranslateImageBatchRequest {
	s.ImageUrls = &v
	return s
}

func (s *TranslateImageBatchRequest) SetSourceLanguage(v string) *TranslateImageBatchRequest {
	s.SourceLanguage = &v
	return s
}

func (s *TranslateImageBatchRequest) SetTargetLanguage(v string) *TranslateImageBatchRequest {
	s.TargetLanguage = &v
	return s
}

func (s *TranslateImageBatchRequest) Validate() error {
	return dara.Validate(s)
}
