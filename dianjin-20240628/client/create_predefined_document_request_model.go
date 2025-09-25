// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePredefinedDocumentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChunks(v []*CreatePredefinedDocumentRequestChunks) *CreatePredefinedDocumentRequest
	GetChunks() []*CreatePredefinedDocumentRequestChunks
	SetLibraryId(v string) *CreatePredefinedDocumentRequest
	GetLibraryId() *string
	SetMetadata(v map[string]interface{}) *CreatePredefinedDocumentRequest
	GetMetadata() map[string]interface{}
	SetTitle(v string) *CreatePredefinedDocumentRequest
	GetTitle() *string
}

type CreatePredefinedDocumentRequest struct {
	Chunks []*CreatePredefinedDocumentRequestChunks `json:"chunks,omitempty" xml:"chunks,omitempty" type:"Repeated"`
	// example:
	//
	// a1b2c3
	LibraryId *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
	// example:
	//
	// {"a": "1"}
	Metadata map[string]interface{} `json:"metadata,omitempty" xml:"metadata,omitempty"`
	// example:
	//
	// 测试文档
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s CreatePredefinedDocumentRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePredefinedDocumentRequest) GoString() string {
	return s.String()
}

func (s *CreatePredefinedDocumentRequest) GetChunks() []*CreatePredefinedDocumentRequestChunks {
	return s.Chunks
}

func (s *CreatePredefinedDocumentRequest) GetLibraryId() *string {
	return s.LibraryId
}

func (s *CreatePredefinedDocumentRequest) GetMetadata() map[string]interface{} {
	return s.Metadata
}

func (s *CreatePredefinedDocumentRequest) GetTitle() *string {
	return s.Title
}

func (s *CreatePredefinedDocumentRequest) SetChunks(v []*CreatePredefinedDocumentRequestChunks) *CreatePredefinedDocumentRequest {
	s.Chunks = v
	return s
}

func (s *CreatePredefinedDocumentRequest) SetLibraryId(v string) *CreatePredefinedDocumentRequest {
	s.LibraryId = &v
	return s
}

func (s *CreatePredefinedDocumentRequest) SetMetadata(v map[string]interface{}) *CreatePredefinedDocumentRequest {
	s.Metadata = v
	return s
}

func (s *CreatePredefinedDocumentRequest) SetTitle(v string) *CreatePredefinedDocumentRequest {
	s.Title = &v
	return s
}

func (s *CreatePredefinedDocumentRequest) Validate() error {
	return dara.Validate(s)
}

type CreatePredefinedDocumentRequestChunks struct {
	// example:
	//
	// {"a": "1"}
	ChunkMeta map[string]interface{} `json:"chunkMeta,omitempty" xml:"chunkMeta,omitempty"`
	// example:
	//
	// 1
	ChunkOrder *int32 `json:"chunkOrder,omitempty" xml:"chunkOrder,omitempty"`
	// example:
	//
	// 这是一段测试文本
	ChunkText *string `json:"chunkText,omitempty" xml:"chunkText,omitempty"`
	// example:
	//
	// text
	ChunkType *string `json:"chunkType,omitempty" xml:"chunkType,omitempty"`
}

func (s CreatePredefinedDocumentRequestChunks) String() string {
	return dara.Prettify(s)
}

func (s CreatePredefinedDocumentRequestChunks) GoString() string {
	return s.String()
}

func (s *CreatePredefinedDocumentRequestChunks) GetChunkMeta() map[string]interface{} {
	return s.ChunkMeta
}

func (s *CreatePredefinedDocumentRequestChunks) GetChunkOrder() *int32 {
	return s.ChunkOrder
}

func (s *CreatePredefinedDocumentRequestChunks) GetChunkText() *string {
	return s.ChunkText
}

func (s *CreatePredefinedDocumentRequestChunks) GetChunkType() *string {
	return s.ChunkType
}

func (s *CreatePredefinedDocumentRequestChunks) SetChunkMeta(v map[string]interface{}) *CreatePredefinedDocumentRequestChunks {
	s.ChunkMeta = v
	return s
}

func (s *CreatePredefinedDocumentRequestChunks) SetChunkOrder(v int32) *CreatePredefinedDocumentRequestChunks {
	s.ChunkOrder = &v
	return s
}

func (s *CreatePredefinedDocumentRequestChunks) SetChunkText(v string) *CreatePredefinedDocumentRequestChunks {
	s.ChunkText = &v
	return s
}

func (s *CreatePredefinedDocumentRequestChunks) SetChunkType(v string) *CreatePredefinedDocumentRequestChunks {
	s.ChunkType = &v
	return s
}

func (s *CreatePredefinedDocumentRequestChunks) Validate() error {
	return dara.Validate(s)
}
