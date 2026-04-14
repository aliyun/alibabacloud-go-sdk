// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOneMetaKnowledgeBaseChunk interface {
	dara.Model
	String() string
	GoString() string
	SetChunkMtime(v string) *OneMetaKnowledgeBaseChunk
	GetChunkMtime() *string
	SetChunkTitle(v string) *OneMetaKnowledgeBaseChunk
	GetChunkTitle() *string
	SetContent(v string) *OneMetaKnowledgeBaseChunk
	GetContent() *string
	SetDocName(v string) *OneMetaKnowledgeBaseChunk
	GetDocName() *string
	SetId(v string) *OneMetaKnowledgeBaseChunk
	GetId() *string
}

type OneMetaKnowledgeBaseChunk struct {
	ChunkMtime *string `json:"ChunkMtime,omitempty" xml:"ChunkMtime,omitempty"`
	ChunkTitle *string `json:"ChunkTitle,omitempty" xml:"ChunkTitle,omitempty"`
	Content    *string `json:"Content,omitempty" xml:"Content,omitempty"`
	DocName    *string `json:"DocName,omitempty" xml:"DocName,omitempty"`
	Id         *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s OneMetaKnowledgeBaseChunk) String() string {
	return dara.Prettify(s)
}

func (s OneMetaKnowledgeBaseChunk) GoString() string {
	return s.String()
}

func (s *OneMetaKnowledgeBaseChunk) GetChunkMtime() *string {
	return s.ChunkMtime
}

func (s *OneMetaKnowledgeBaseChunk) GetChunkTitle() *string {
	return s.ChunkTitle
}

func (s *OneMetaKnowledgeBaseChunk) GetContent() *string {
	return s.Content
}

func (s *OneMetaKnowledgeBaseChunk) GetDocName() *string {
	return s.DocName
}

func (s *OneMetaKnowledgeBaseChunk) GetId() *string {
	return s.Id
}

func (s *OneMetaKnowledgeBaseChunk) SetChunkMtime(v string) *OneMetaKnowledgeBaseChunk {
	s.ChunkMtime = &v
	return s
}

func (s *OneMetaKnowledgeBaseChunk) SetChunkTitle(v string) *OneMetaKnowledgeBaseChunk {
	s.ChunkTitle = &v
	return s
}

func (s *OneMetaKnowledgeBaseChunk) SetContent(v string) *OneMetaKnowledgeBaseChunk {
	s.Content = &v
	return s
}

func (s *OneMetaKnowledgeBaseChunk) SetDocName(v string) *OneMetaKnowledgeBaseChunk {
	s.DocName = &v
	return s
}

func (s *OneMetaKnowledgeBaseChunk) SetId(v string) *OneMetaKnowledgeBaseChunk {
	s.Id = &v
	return s
}

func (s *OneMetaKnowledgeBaseChunk) Validate() error {
	return dara.Validate(s)
}
