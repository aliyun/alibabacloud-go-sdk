// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDocumentAnalyzeTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDocument(v *CreateDocumentAnalyzeTaskRequestDocument) *CreateDocumentAnalyzeTaskRequest
	GetDocument() *CreateDocumentAnalyzeTaskRequestDocument
	SetOutput(v *CreateDocumentAnalyzeTaskRequestOutput) *CreateDocumentAnalyzeTaskRequest
	GetOutput() *CreateDocumentAnalyzeTaskRequestOutput
	SetStrategy(v *CreateDocumentAnalyzeTaskRequestStrategy) *CreateDocumentAnalyzeTaskRequest
	GetStrategy() *CreateDocumentAnalyzeTaskRequestStrategy
}

type CreateDocumentAnalyzeTaskRequest struct {
	Document *CreateDocumentAnalyzeTaskRequestDocument `json:"document,omitempty" xml:"document,omitempty" type:"Struct"`
	Output   *CreateDocumentAnalyzeTaskRequestOutput   `json:"output,omitempty" xml:"output,omitempty" type:"Struct"`
	Strategy *CreateDocumentAnalyzeTaskRequestStrategy `json:"strategy,omitempty" xml:"strategy,omitempty" type:"Struct"`
}

func (s CreateDocumentAnalyzeTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDocumentAnalyzeTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateDocumentAnalyzeTaskRequest) GetDocument() *CreateDocumentAnalyzeTaskRequestDocument {
	return s.Document
}

func (s *CreateDocumentAnalyzeTaskRequest) GetOutput() *CreateDocumentAnalyzeTaskRequestOutput {
	return s.Output
}

func (s *CreateDocumentAnalyzeTaskRequest) GetStrategy() *CreateDocumentAnalyzeTaskRequestStrategy {
	return s.Strategy
}

func (s *CreateDocumentAnalyzeTaskRequest) SetDocument(v *CreateDocumentAnalyzeTaskRequestDocument) *CreateDocumentAnalyzeTaskRequest {
	s.Document = v
	return s
}

func (s *CreateDocumentAnalyzeTaskRequest) SetOutput(v *CreateDocumentAnalyzeTaskRequestOutput) *CreateDocumentAnalyzeTaskRequest {
	s.Output = v
	return s
}

func (s *CreateDocumentAnalyzeTaskRequest) SetStrategy(v *CreateDocumentAnalyzeTaskRequestStrategy) *CreateDocumentAnalyzeTaskRequest {
	s.Strategy = v
	return s
}

func (s *CreateDocumentAnalyzeTaskRequest) Validate() error {
	if s.Document != nil {
		if err := s.Document.Validate(); err != nil {
			return err
		}
	}
	if s.Output != nil {
		if err := s.Output.Validate(); err != nil {
			return err
		}
	}
	if s.Strategy != nil {
		if err := s.Strategy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDocumentAnalyzeTaskRequestDocument struct {
	Content  *string `json:"content,omitempty" xml:"content,omitempty"`
	FileName *string `json:"file_name,omitempty" xml:"file_name,omitempty"`
	FileType *string `json:"file_type,omitempty" xml:"file_type,omitempty"`
	Url      *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s CreateDocumentAnalyzeTaskRequestDocument) String() string {
	return dara.Prettify(s)
}

func (s CreateDocumentAnalyzeTaskRequestDocument) GoString() string {
	return s.String()
}

func (s *CreateDocumentAnalyzeTaskRequestDocument) GetContent() *string {
	return s.Content
}

func (s *CreateDocumentAnalyzeTaskRequestDocument) GetFileName() *string {
	return s.FileName
}

func (s *CreateDocumentAnalyzeTaskRequestDocument) GetFileType() *string {
	return s.FileType
}

func (s *CreateDocumentAnalyzeTaskRequestDocument) GetUrl() *string {
	return s.Url
}

func (s *CreateDocumentAnalyzeTaskRequestDocument) SetContent(v string) *CreateDocumentAnalyzeTaskRequestDocument {
	s.Content = &v
	return s
}

func (s *CreateDocumentAnalyzeTaskRequestDocument) SetFileName(v string) *CreateDocumentAnalyzeTaskRequestDocument {
	s.FileName = &v
	return s
}

func (s *CreateDocumentAnalyzeTaskRequestDocument) SetFileType(v string) *CreateDocumentAnalyzeTaskRequestDocument {
	s.FileType = &v
	return s
}

func (s *CreateDocumentAnalyzeTaskRequestDocument) SetUrl(v string) *CreateDocumentAnalyzeTaskRequestDocument {
	s.Url = &v
	return s
}

func (s *CreateDocumentAnalyzeTaskRequestDocument) Validate() error {
	return dara.Validate(s)
}

type CreateDocumentAnalyzeTaskRequestOutput struct {
	ImageStorage *string `json:"image_storage,omitempty" xml:"image_storage,omitempty"`
}

func (s CreateDocumentAnalyzeTaskRequestOutput) String() string {
	return dara.Prettify(s)
}

func (s CreateDocumentAnalyzeTaskRequestOutput) GoString() string {
	return s.String()
}

func (s *CreateDocumentAnalyzeTaskRequestOutput) GetImageStorage() *string {
	return s.ImageStorage
}

func (s *CreateDocumentAnalyzeTaskRequestOutput) SetImageStorage(v string) *CreateDocumentAnalyzeTaskRequestOutput {
	s.ImageStorage = &v
	return s
}

func (s *CreateDocumentAnalyzeTaskRequestOutput) Validate() error {
	return dara.Validate(s)
}

type CreateDocumentAnalyzeTaskRequestStrategy struct {
	EnableSemantic *bool `json:"enable_semantic,omitempty" xml:"enable_semantic,omitempty"`
}

func (s CreateDocumentAnalyzeTaskRequestStrategy) String() string {
	return dara.Prettify(s)
}

func (s CreateDocumentAnalyzeTaskRequestStrategy) GoString() string {
	return s.String()
}

func (s *CreateDocumentAnalyzeTaskRequestStrategy) GetEnableSemantic() *bool {
	return s.EnableSemantic
}

func (s *CreateDocumentAnalyzeTaskRequestStrategy) SetEnableSemantic(v bool) *CreateDocumentAnalyzeTaskRequestStrategy {
	s.EnableSemantic = &v
	return s
}

func (s *CreateDocumentAnalyzeTaskRequestStrategy) Validate() error {
	return dara.Validate(s)
}
