// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePdfTranslateTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDocId(v string) *CreatePdfTranslateTaskRequest
	GetDocId() *string
	SetKnowledge(v string) *CreatePdfTranslateTaskRequest
	GetKnowledge() *string
	SetLibraryId(v string) *CreatePdfTranslateTaskRequest
	GetLibraryId() *string
	SetModelId(v string) *CreatePdfTranslateTaskRequest
	GetModelId() *string
	SetTranslateTo(v string) *CreatePdfTranslateTaskRequest
	GetTranslateTo() *string
}

type CreatePdfTranslateTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 873648346573245
	DocId     *string `json:"docId,omitempty" xml:"docId,omitempty"`
	Knowledge *string `json:"knowledge,omitempty" xml:"knowledge,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cjshcxxxx
	LibraryId *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// qwen-plus
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// example:
	//
	// 中文
	TranslateTo *string `json:"translateTo,omitempty" xml:"translateTo,omitempty"`
}

func (s CreatePdfTranslateTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePdfTranslateTaskRequest) GoString() string {
	return s.String()
}

func (s *CreatePdfTranslateTaskRequest) GetDocId() *string {
	return s.DocId
}

func (s *CreatePdfTranslateTaskRequest) GetKnowledge() *string {
	return s.Knowledge
}

func (s *CreatePdfTranslateTaskRequest) GetLibraryId() *string {
	return s.LibraryId
}

func (s *CreatePdfTranslateTaskRequest) GetModelId() *string {
	return s.ModelId
}

func (s *CreatePdfTranslateTaskRequest) GetTranslateTo() *string {
	return s.TranslateTo
}

func (s *CreatePdfTranslateTaskRequest) SetDocId(v string) *CreatePdfTranslateTaskRequest {
	s.DocId = &v
	return s
}

func (s *CreatePdfTranslateTaskRequest) SetKnowledge(v string) *CreatePdfTranslateTaskRequest {
	s.Knowledge = &v
	return s
}

func (s *CreatePdfTranslateTaskRequest) SetLibraryId(v string) *CreatePdfTranslateTaskRequest {
	s.LibraryId = &v
	return s
}

func (s *CreatePdfTranslateTaskRequest) SetModelId(v string) *CreatePdfTranslateTaskRequest {
	s.ModelId = &v
	return s
}

func (s *CreatePdfTranslateTaskRequest) SetTranslateTo(v string) *CreatePdfTranslateTaskRequest {
	s.TranslateTo = &v
	return s
}

func (s *CreatePdfTranslateTaskRequest) Validate() error {
	return dara.Validate(s)
}
