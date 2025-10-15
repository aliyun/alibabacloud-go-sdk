// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDocumentChunkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChunks(v []*UpdateDocumentChunkRequestChunks) *UpdateDocumentChunkRequest
	GetChunks() []*UpdateDocumentChunkRequestChunks
	SetLibraryId(v string) *UpdateDocumentChunkRequest
	GetLibraryId() *string
}

type UpdateDocumentChunkRequest struct {
	// This parameter is required.
	Chunks []*UpdateDocumentChunkRequestChunks `json:"chunks,omitempty" xml:"chunks,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// sjdgdsfg
	LibraryId *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
}

func (s UpdateDocumentChunkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDocumentChunkRequest) GoString() string {
	return s.String()
}

func (s *UpdateDocumentChunkRequest) GetChunks() []*UpdateDocumentChunkRequestChunks {
	return s.Chunks
}

func (s *UpdateDocumentChunkRequest) GetLibraryId() *string {
	return s.LibraryId
}

func (s *UpdateDocumentChunkRequest) SetChunks(v []*UpdateDocumentChunkRequestChunks) *UpdateDocumentChunkRequest {
	s.Chunks = v
	return s
}

func (s *UpdateDocumentChunkRequest) SetLibraryId(v string) *UpdateDocumentChunkRequest {
	s.LibraryId = &v
	return s
}

func (s *UpdateDocumentChunkRequest) Validate() error {
	if s.Chunks != nil {
		for _, item := range s.Chunks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateDocumentChunkRequestChunks struct {
	// This parameter is required.
	//
	// example:
	//
	// 1987834755763847
	ChunkId *string `json:"chunkId,omitempty" xml:"chunkId,omitempty"`
	// This parameter is required.
	ChunkText *string `json:"chunkText,omitempty" xml:"chunkText,omitempty"`
}

func (s UpdateDocumentChunkRequestChunks) String() string {
	return dara.Prettify(s)
}

func (s UpdateDocumentChunkRequestChunks) GoString() string {
	return s.String()
}

func (s *UpdateDocumentChunkRequestChunks) GetChunkId() *string {
	return s.ChunkId
}

func (s *UpdateDocumentChunkRequestChunks) GetChunkText() *string {
	return s.ChunkText
}

func (s *UpdateDocumentChunkRequestChunks) SetChunkId(v string) *UpdateDocumentChunkRequestChunks {
	s.ChunkId = &v
	return s
}

func (s *UpdateDocumentChunkRequestChunks) SetChunkText(v string) *UpdateDocumentChunkRequestChunks {
	s.ChunkText = &v
	return s
}

func (s *UpdateDocumentChunkRequestChunks) Validate() error {
	return dara.Validate(s)
}
