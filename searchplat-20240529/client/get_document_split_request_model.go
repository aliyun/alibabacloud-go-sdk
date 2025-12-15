// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocumentSplitRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDocument(v *GetDocumentSplitRequestDocument) *GetDocumentSplitRequest
	GetDocument() *GetDocumentSplitRequestDocument
	SetStrategy(v *GetDocumentSplitRequestStrategy) *GetDocumentSplitRequest
	GetStrategy() *GetDocumentSplitRequestStrategy
}

type GetDocumentSplitRequest struct {
	// This parameter is required.
	Document *GetDocumentSplitRequestDocument `json:"document,omitempty" xml:"document,omitempty" type:"Struct"`
	Strategy *GetDocumentSplitRequestStrategy `json:"strategy,omitempty" xml:"strategy,omitempty" type:"Struct"`
}

func (s GetDocumentSplitRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentSplitRequest) GoString() string {
	return s.String()
}

func (s *GetDocumentSplitRequest) GetDocument() *GetDocumentSplitRequestDocument {
	return s.Document
}

func (s *GetDocumentSplitRequest) GetStrategy() *GetDocumentSplitRequestStrategy {
	return s.Strategy
}

func (s *GetDocumentSplitRequest) SetDocument(v *GetDocumentSplitRequestDocument) *GetDocumentSplitRequest {
	s.Document = v
	return s
}

func (s *GetDocumentSplitRequest) SetStrategy(v *GetDocumentSplitRequestStrategy) *GetDocumentSplitRequest {
	s.Strategy = v
	return s
}

func (s *GetDocumentSplitRequest) Validate() error {
	if s.Document != nil {
		if err := s.Document.Validate(); err != nil {
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

type GetDocumentSplitRequestDocument struct {
	Content         *string `json:"content,omitempty" xml:"content,omitempty"`
	ContentEncoding *string `json:"content_encoding,omitempty" xml:"content_encoding,omitempty"`
	ContentType     *string `json:"content_type,omitempty" xml:"content_type,omitempty"`
}

func (s GetDocumentSplitRequestDocument) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentSplitRequestDocument) GoString() string {
	return s.String()
}

func (s *GetDocumentSplitRequestDocument) GetContent() *string {
	return s.Content
}

func (s *GetDocumentSplitRequestDocument) GetContentEncoding() *string {
	return s.ContentEncoding
}

func (s *GetDocumentSplitRequestDocument) GetContentType() *string {
	return s.ContentType
}

func (s *GetDocumentSplitRequestDocument) SetContent(v string) *GetDocumentSplitRequestDocument {
	s.Content = &v
	return s
}

func (s *GetDocumentSplitRequestDocument) SetContentEncoding(v string) *GetDocumentSplitRequestDocument {
	s.ContentEncoding = &v
	return s
}

func (s *GetDocumentSplitRequestDocument) SetContentType(v string) *GetDocumentSplitRequestDocument {
	s.ContentType = &v
	return s
}

func (s *GetDocumentSplitRequestDocument) Validate() error {
	return dara.Validate(s)
}

type GetDocumentSplitRequestStrategy struct {
	ComputeType  *string `json:"compute_type,omitempty" xml:"compute_type,omitempty"`
	MaxChunkSize *int64  `json:"max_chunk_size,omitempty" xml:"max_chunk_size,omitempty"`
	NeedSentence *bool   `json:"need_sentence,omitempty" xml:"need_sentence,omitempty"`
}

func (s GetDocumentSplitRequestStrategy) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentSplitRequestStrategy) GoString() string {
	return s.String()
}

func (s *GetDocumentSplitRequestStrategy) GetComputeType() *string {
	return s.ComputeType
}

func (s *GetDocumentSplitRequestStrategy) GetMaxChunkSize() *int64 {
	return s.MaxChunkSize
}

func (s *GetDocumentSplitRequestStrategy) GetNeedSentence() *bool {
	return s.NeedSentence
}

func (s *GetDocumentSplitRequestStrategy) SetComputeType(v string) *GetDocumentSplitRequestStrategy {
	s.ComputeType = &v
	return s
}

func (s *GetDocumentSplitRequestStrategy) SetMaxChunkSize(v int64) *GetDocumentSplitRequestStrategy {
	s.MaxChunkSize = &v
	return s
}

func (s *GetDocumentSplitRequestStrategy) SetNeedSentence(v bool) *GetDocumentSplitRequestStrategy {
	s.NeedSentence = &v
	return s
}

func (s *GetDocumentSplitRequestStrategy) Validate() error {
	return dara.Validate(s)
}
