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
	// The document ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 873648346573245
	DocId *string `json:"docId,omitempty" xml:"docId,omitempty"`
	// The domain knowledge referenced during translation.
	//
	// example:
	//
	// Net Profit
	//
	// English: Net Profit
	//
	// Chinese: Net profit (typically refers to the profit after deducting all expenses and taxes)
	Knowledge *string `json:"knowledge,omitempty" xml:"knowledge,omitempty"`
	// The document library ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cjshcxxxx
	LibraryId *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
	// The model ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// qwen-plus
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// The target language. Default value: Chinese.
	//
	// example:
	//
	// Chinese
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
