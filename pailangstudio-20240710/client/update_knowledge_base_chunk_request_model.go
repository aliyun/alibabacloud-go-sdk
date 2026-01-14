// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKnowledgeBaseChunkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChunkContent(v string) *UpdateKnowledgeBaseChunkRequest
	GetChunkContent() *string
	SetChunkStatus(v string) *UpdateKnowledgeBaseChunkRequest
	GetChunkStatus() *string
}

type UpdateKnowledgeBaseChunkRequest struct {
	// example:
	//
	// content
	ChunkContent *string `json:"ChunkContent,omitempty" xml:"ChunkContent,omitempty"`
	// example:
	//
	// Enable
	ChunkStatus *string `json:"ChunkStatus,omitempty" xml:"ChunkStatus,omitempty"`
}

func (s UpdateKnowledgeBaseChunkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeBaseChunkRequest) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeBaseChunkRequest) GetChunkContent() *string {
	return s.ChunkContent
}

func (s *UpdateKnowledgeBaseChunkRequest) GetChunkStatus() *string {
	return s.ChunkStatus
}

func (s *UpdateKnowledgeBaseChunkRequest) SetChunkContent(v string) *UpdateKnowledgeBaseChunkRequest {
	s.ChunkContent = &v
	return s
}

func (s *UpdateKnowledgeBaseChunkRequest) SetChunkStatus(v string) *UpdateKnowledgeBaseChunkRequest {
	s.ChunkStatus = &v
	return s
}

func (s *UpdateKnowledgeBaseChunkRequest) Validate() error {
	return dara.Validate(s)
}
