// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImageAnalyzeTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDocument(v *CreateImageAnalyzeTaskRequestDocument) *CreateImageAnalyzeTaskRequest
	GetDocument() *CreateImageAnalyzeTaskRequestDocument
}

type CreateImageAnalyzeTaskRequest struct {
	Document *CreateImageAnalyzeTaskRequestDocument `json:"document,omitempty" xml:"document,omitempty" type:"Struct"`
}

func (s CreateImageAnalyzeTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateImageAnalyzeTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateImageAnalyzeTaskRequest) GetDocument() *CreateImageAnalyzeTaskRequestDocument {
	return s.Document
}

func (s *CreateImageAnalyzeTaskRequest) SetDocument(v *CreateImageAnalyzeTaskRequestDocument) *CreateImageAnalyzeTaskRequest {
	s.Document = v
	return s
}

func (s *CreateImageAnalyzeTaskRequest) Validate() error {
	if s.Document != nil {
		if err := s.Document.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateImageAnalyzeTaskRequestDocument struct {
	Content  *string `json:"content,omitempty" xml:"content,omitempty"`
	FileName *string `json:"file_name,omitempty" xml:"file_name,omitempty"`
	FileType *string `json:"file_type,omitempty" xml:"file_type,omitempty"`
	Url      *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s CreateImageAnalyzeTaskRequestDocument) String() string {
	return dara.Prettify(s)
}

func (s CreateImageAnalyzeTaskRequestDocument) GoString() string {
	return s.String()
}

func (s *CreateImageAnalyzeTaskRequestDocument) GetContent() *string {
	return s.Content
}

func (s *CreateImageAnalyzeTaskRequestDocument) GetFileName() *string {
	return s.FileName
}

func (s *CreateImageAnalyzeTaskRequestDocument) GetFileType() *string {
	return s.FileType
}

func (s *CreateImageAnalyzeTaskRequestDocument) GetUrl() *string {
	return s.Url
}

func (s *CreateImageAnalyzeTaskRequestDocument) SetContent(v string) *CreateImageAnalyzeTaskRequestDocument {
	s.Content = &v
	return s
}

func (s *CreateImageAnalyzeTaskRequestDocument) SetFileName(v string) *CreateImageAnalyzeTaskRequestDocument {
	s.FileName = &v
	return s
}

func (s *CreateImageAnalyzeTaskRequestDocument) SetFileType(v string) *CreateImageAnalyzeTaskRequestDocument {
	s.FileType = &v
	return s
}

func (s *CreateImageAnalyzeTaskRequestDocument) SetUrl(v string) *CreateImageAnalyzeTaskRequestDocument {
	s.Url = &v
	return s
}

func (s *CreateImageAnalyzeTaskRequestDocument) Validate() error {
	return dara.Validate(s)
}
