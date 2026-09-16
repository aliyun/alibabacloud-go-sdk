// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKnowledgeBaseChunk interface {
	dara.Model
	String() string
	GoString() string
	SetChunkSeq(v int32) *KnowledgeBaseChunk
	GetChunkSeq() *int32
	SetChunkSize(v int32) *KnowledgeBaseChunk
	GetChunkSize() *int32
	SetContent(v string) *KnowledgeBaseChunk
	GetContent() *string
	SetCreatedAt(v string) *KnowledgeBaseChunk
	GetCreatedAt() *string
	SetDocumentId(v string) *KnowledgeBaseChunk
	GetDocumentId() *string
	SetEnabled(v bool) *KnowledgeBaseChunk
	GetEnabled() *bool
	SetFileName(v string) *KnowledgeBaseChunk
	GetFileName() *string
	SetSourceLocation(v string) *KnowledgeBaseChunk
	GetSourceLocation() *string
	SetTitlePath(v string) *KnowledgeBaseChunk
	GetTitlePath() *string
	SetUpdatedAt(v string) *KnowledgeBaseChunk
	GetUpdatedAt() *string
}

type KnowledgeBaseChunk struct {
	// The sequence number of the chunk within the document, starting from 1 and numbered consecutively.
	//
	// example:
	//
	// 12
	ChunkSeq *int32 `json:"ChunkSeq,omitempty" xml:"ChunkSeq,omitempty"`
	// The number of characters in the chunk content, measured in UTF-16 code units, consistent with MaxChunkSize. You can use this value to evaluate chunk saturation against the chunking configuration.
	//
	// example:
	//
	// 128
	ChunkSize *int32 `json:"ChunkSize,omitempty" xml:"ChunkSize,omitempty"`
	// The content of the chunk.
	//
	// example:
	//
	// EventBridge supports routing events to multiple target services
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The time when the chunk was created.
	//
	// example:
	//
	// 2026-08-24T10:00:00Z
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// The ID of the document to which the chunk belongs.
	//
	// example:
	//
	// doc-bp1xxxxxxxxxxxx
	DocumentId *string `json:"DocumentId,omitempty" xml:"DocumentId,omitempty"`
	// Indicates whether the chunk is enabled. Disabled chunks are excluded from retrieval.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The file name of the document to which the chunk belongs. This value is from the same source as the FileName returned by GetDocument.
	//
	// example:
	//
	// product-handbook.pdf
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The location of the chunk in the original document. The format varies by document type: for PDF, the value is p.PageNumber (such as p.3). For PPT/PPTX, the value is s.SlideNumber (such as s.2). For XLS/XLSX, the value is the sheet name. For other formats (such as txt, md, html, doc, or docx), this field is not returned if no source location is available.
	//
	// example:
	//
	// p.3
	SourceLocation *string `json:"SourceLocation,omitempty" xml:"SourceLocation,omitempty"`
	// The hierarchical title path of the chunk, connected by >. If no recognizable title exists in the original document, the value falls back to a summary of the first paragraph content (such as CONTENT). This field is for display purposes only.
	//
	// example:
	//
	// Installation Guide>Prerequisites
	TitlePath *string `json:"TitlePath,omitempty" xml:"TitlePath,omitempty"`
	// The time when the chunk was last updated.
	//
	// example:
	//
	// 2026-08-24T10:00:00Z
	UpdatedAt *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
}

func (s KnowledgeBaseChunk) String() string {
	return dara.Prettify(s)
}

func (s KnowledgeBaseChunk) GoString() string {
	return s.String()
}

func (s *KnowledgeBaseChunk) GetChunkSeq() *int32 {
	return s.ChunkSeq
}

func (s *KnowledgeBaseChunk) GetChunkSize() *int32 {
	return s.ChunkSize
}

func (s *KnowledgeBaseChunk) GetContent() *string {
	return s.Content
}

func (s *KnowledgeBaseChunk) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *KnowledgeBaseChunk) GetDocumentId() *string {
	return s.DocumentId
}

func (s *KnowledgeBaseChunk) GetEnabled() *bool {
	return s.Enabled
}

func (s *KnowledgeBaseChunk) GetFileName() *string {
	return s.FileName
}

func (s *KnowledgeBaseChunk) GetSourceLocation() *string {
	return s.SourceLocation
}

func (s *KnowledgeBaseChunk) GetTitlePath() *string {
	return s.TitlePath
}

func (s *KnowledgeBaseChunk) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *KnowledgeBaseChunk) SetChunkSeq(v int32) *KnowledgeBaseChunk {
	s.ChunkSeq = &v
	return s
}

func (s *KnowledgeBaseChunk) SetChunkSize(v int32) *KnowledgeBaseChunk {
	s.ChunkSize = &v
	return s
}

func (s *KnowledgeBaseChunk) SetContent(v string) *KnowledgeBaseChunk {
	s.Content = &v
	return s
}

func (s *KnowledgeBaseChunk) SetCreatedAt(v string) *KnowledgeBaseChunk {
	s.CreatedAt = &v
	return s
}

func (s *KnowledgeBaseChunk) SetDocumentId(v string) *KnowledgeBaseChunk {
	s.DocumentId = &v
	return s
}

func (s *KnowledgeBaseChunk) SetEnabled(v bool) *KnowledgeBaseChunk {
	s.Enabled = &v
	return s
}

func (s *KnowledgeBaseChunk) SetFileName(v string) *KnowledgeBaseChunk {
	s.FileName = &v
	return s
}

func (s *KnowledgeBaseChunk) SetSourceLocation(v string) *KnowledgeBaseChunk {
	s.SourceLocation = &v
	return s
}

func (s *KnowledgeBaseChunk) SetTitlePath(v string) *KnowledgeBaseChunk {
	s.TitlePath = &v
	return s
}

func (s *KnowledgeBaseChunk) SetUpdatedAt(v string) *KnowledgeBaseChunk {
	s.UpdatedAt = &v
	return s
}

func (s *KnowledgeBaseChunk) Validate() error {
	return dara.Validate(s)
}
