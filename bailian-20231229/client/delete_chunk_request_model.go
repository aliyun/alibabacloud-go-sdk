// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteChunkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChunkIds(v []*string) *DeleteChunkRequest
	GetChunkIds() []*string
	SetPipelineId(v string) *DeleteChunkRequest
	GetPipelineId() *string
}

type DeleteChunkRequest struct {
	// The list of text chunks to be deleted. You can specify up to 10 chunk IDs at a time.
	//
	// This parameter is required.
	ChunkIds []*string `json:"ChunkIds,omitempty" xml:"ChunkIds,omitempty" type:"Repeated"`
	// The knowledge base ID, which is the `Data.Id` parameter returned by **CreateIndex**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 79c0alxxxx
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
}

func (s DeleteChunkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteChunkRequest) GoString() string {
	return s.String()
}

func (s *DeleteChunkRequest) GetChunkIds() []*string {
	return s.ChunkIds
}

func (s *DeleteChunkRequest) GetPipelineId() *string {
	return s.PipelineId
}

func (s *DeleteChunkRequest) SetChunkIds(v []*string) *DeleteChunkRequest {
	s.ChunkIds = v
	return s
}

func (s *DeleteChunkRequest) SetPipelineId(v string) *DeleteChunkRequest {
	s.PipelineId = &v
	return s
}

func (s *DeleteChunkRequest) Validate() error {
	return dara.Validate(s)
}
