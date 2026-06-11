// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDocumentChunksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChunkIds(v []*string) *DeleteDocumentChunksRequest
	GetChunkIds() []*string
	SetDocumentName(v string) *DeleteDocumentChunksRequest
	GetDocumentName() *string
	SetKbUuid(v string) *DeleteDocumentChunksRequest
	GetKbUuid() *string
}

type DeleteDocumentChunksRequest struct {
	// A list of chunk IDs.
	//
	// This parameter is required.
	ChunkIds []*string `json:"ChunkIds,omitempty" xml:"ChunkIds,omitempty" type:"Repeated"`
	// The name of the document.
	//
	// This parameter is required.
	//
	// example:
	//
	// test.md
	DocumentName *string `json:"DocumentName,omitempty" xml:"DocumentName,omitempty"`
	// The ID of the knowledge base.
	//
	// This parameter is required.
	//
	// example:
	//
	// kb-***
	KbUuid *string `json:"KbUuid,omitempty" xml:"KbUuid,omitempty"`
}

func (s DeleteDocumentChunksRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDocumentChunksRequest) GoString() string {
	return s.String()
}

func (s *DeleteDocumentChunksRequest) GetChunkIds() []*string {
	return s.ChunkIds
}

func (s *DeleteDocumentChunksRequest) GetDocumentName() *string {
	return s.DocumentName
}

func (s *DeleteDocumentChunksRequest) GetKbUuid() *string {
	return s.KbUuid
}

func (s *DeleteDocumentChunksRequest) SetChunkIds(v []*string) *DeleteDocumentChunksRequest {
	s.ChunkIds = v
	return s
}

func (s *DeleteDocumentChunksRequest) SetDocumentName(v string) *DeleteDocumentChunksRequest {
	s.DocumentName = &v
	return s
}

func (s *DeleteDocumentChunksRequest) SetKbUuid(v string) *DeleteDocumentChunksRequest {
	s.KbUuid = &v
	return s
}

func (s *DeleteDocumentChunksRequest) Validate() error {
	return dara.Validate(s)
}
