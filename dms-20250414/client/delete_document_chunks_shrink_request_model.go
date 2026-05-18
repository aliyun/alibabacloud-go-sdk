// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDocumentChunksShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChunkIdsShrink(v string) *DeleteDocumentChunksShrinkRequest
	GetChunkIdsShrink() *string
	SetDocumentName(v string) *DeleteDocumentChunksShrinkRequest
	GetDocumentName() *string
	SetKbUuid(v string) *DeleteDocumentChunksShrinkRequest
	GetKbUuid() *string
}

type DeleteDocumentChunksShrinkRequest struct {
	// This parameter is required.
	ChunkIdsShrink *string `json:"ChunkIds,omitempty" xml:"ChunkIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test.md
	DocumentName *string `json:"DocumentName,omitempty" xml:"DocumentName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// kb-***
	KbUuid *string `json:"KbUuid,omitempty" xml:"KbUuid,omitempty"`
}

func (s DeleteDocumentChunksShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDocumentChunksShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteDocumentChunksShrinkRequest) GetChunkIdsShrink() *string {
	return s.ChunkIdsShrink
}

func (s *DeleteDocumentChunksShrinkRequest) GetDocumentName() *string {
	return s.DocumentName
}

func (s *DeleteDocumentChunksShrinkRequest) GetKbUuid() *string {
	return s.KbUuid
}

func (s *DeleteDocumentChunksShrinkRequest) SetChunkIdsShrink(v string) *DeleteDocumentChunksShrinkRequest {
	s.ChunkIdsShrink = &v
	return s
}

func (s *DeleteDocumentChunksShrinkRequest) SetDocumentName(v string) *DeleteDocumentChunksShrinkRequest {
	s.DocumentName = &v
	return s
}

func (s *DeleteDocumentChunksShrinkRequest) SetKbUuid(v string) *DeleteDocumentChunksShrinkRequest {
	s.KbUuid = &v
	return s
}

func (s *DeleteDocumentChunksShrinkRequest) Validate() error {
	return dara.Validate(s)
}
