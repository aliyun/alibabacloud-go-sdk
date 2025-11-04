// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteChunkShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChunkIdsShrink(v string) *DeleteChunkShrinkRequest
	GetChunkIdsShrink() *string
	SetPipelineId(v string) *DeleteChunkShrinkRequest
	GetPipelineId() *string
}

type DeleteChunkShrinkRequest struct {
	// The list of text chunks to be deleted. You can specify up to 10 chunk IDs at a time.
	//
	// This parameter is required.
	ChunkIdsShrink *string `json:"ChunkIds,omitempty" xml:"ChunkIds,omitempty"`
	// The knowledge base ID, which is the `Data.Id` parameter returned by **CreateIndex**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 79c0alxxxx
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
}

func (s DeleteChunkShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteChunkShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteChunkShrinkRequest) GetChunkIdsShrink() *string {
	return s.ChunkIdsShrink
}

func (s *DeleteChunkShrinkRequest) GetPipelineId() *string {
	return s.PipelineId
}

func (s *DeleteChunkShrinkRequest) SetChunkIdsShrink(v string) *DeleteChunkShrinkRequest {
	s.ChunkIdsShrink = &v
	return s
}

func (s *DeleteChunkShrinkRequest) SetPipelineId(v string) *DeleteChunkShrinkRequest {
	s.PipelineId = &v
	return s
}

func (s *DeleteChunkShrinkRequest) Validate() error {
	return dara.Validate(s)
}
