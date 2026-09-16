// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKnowledgeBaseSearchChunk interface {
	dara.Model
	String() string
	GoString() string
	SetChunkSeq(v int32) *KnowledgeBaseSearchChunk
	GetChunkSeq() *int32
	SetContent(v string) *KnowledgeBaseSearchChunk
	GetContent() *string
	SetDocumentId(v string) *KnowledgeBaseSearchChunk
	GetDocumentId() *string
	SetFileName(v string) *KnowledgeBaseSearchChunk
	GetFileName() *string
	SetScore(v float64) *KnowledgeBaseSearchChunk
	GetScore() *float64
	SetScores(v *KnowledgeBaseSearchChunkScores) *KnowledgeBaseSearchChunk
	GetScores() *KnowledgeBaseSearchChunkScores
	SetSourceLocation(v string) *KnowledgeBaseSearchChunk
	GetSourceLocation() *string
	SetTitlePath(v string) *KnowledgeBaseSearchChunk
	GetTitlePath() *string
}

type KnowledgeBaseSearchChunk struct {
	// The sequence number of the chunk within the document.
	//
	// example:
	//
	// 12
	ChunkSeq *int32 `json:"ChunkSeq,omitempty" xml:"ChunkSeq,omitempty"`
	// The body content of the hit chunk.
	//
	// example:
	//
	// EventBridge supports routing events to multiple target services
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The ID of the document to which the hit chunk belongs.
	//
	// example:
	//
	// doc-bp1xxxxxxxxxxxx
	DocumentId *string `json:"DocumentId,omitempty" xml:"DocumentId,omitempty"`
	// The file name of the document to which the hit chunk belongs. This value has the same source as the FileName returned by GetDocument and can be used to render the reference source.
	//
	// example:
	//
	// product-handbook.pdf
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The retrieval relevance score. A higher score indicates higher relevance.
	//
	// example:
	//
	// 0.92
	Score *float64 `json:"Score,omitempty" xml:"Score,omitempty"`
	// The score details for each stage. Score fields that are not involved in the calculation are not returned.
	Scores *KnowledgeBaseSearchChunkScores `json:"Scores,omitempty" xml:"Scores,omitempty" type:"Struct"`
	// The location of the chunk in the original document. p.N indicates page N (PDF). s.N indicates slide N (PPT/PPTX).
	//
	// example:
	//
	// p.3
	SourceLocation *string `json:"SourceLocation,omitempty" xml:"SourceLocation,omitempty"`
	// The title path to which the chunk belongs, such as Chapter 1>1.1 Overview.
	//
	// example:
	//
	// Installation Guide>Prerequisites
	TitlePath *string `json:"TitlePath,omitempty" xml:"TitlePath,omitempty"`
}

func (s KnowledgeBaseSearchChunk) String() string {
	return dara.Prettify(s)
}

func (s KnowledgeBaseSearchChunk) GoString() string {
	return s.String()
}

func (s *KnowledgeBaseSearchChunk) GetChunkSeq() *int32 {
	return s.ChunkSeq
}

func (s *KnowledgeBaseSearchChunk) GetContent() *string {
	return s.Content
}

func (s *KnowledgeBaseSearchChunk) GetDocumentId() *string {
	return s.DocumentId
}

func (s *KnowledgeBaseSearchChunk) GetFileName() *string {
	return s.FileName
}

func (s *KnowledgeBaseSearchChunk) GetScore() *float64 {
	return s.Score
}

func (s *KnowledgeBaseSearchChunk) GetScores() *KnowledgeBaseSearchChunkScores {
	return s.Scores
}

func (s *KnowledgeBaseSearchChunk) GetSourceLocation() *string {
	return s.SourceLocation
}

func (s *KnowledgeBaseSearchChunk) GetTitlePath() *string {
	return s.TitlePath
}

func (s *KnowledgeBaseSearchChunk) SetChunkSeq(v int32) *KnowledgeBaseSearchChunk {
	s.ChunkSeq = &v
	return s
}

func (s *KnowledgeBaseSearchChunk) SetContent(v string) *KnowledgeBaseSearchChunk {
	s.Content = &v
	return s
}

func (s *KnowledgeBaseSearchChunk) SetDocumentId(v string) *KnowledgeBaseSearchChunk {
	s.DocumentId = &v
	return s
}

func (s *KnowledgeBaseSearchChunk) SetFileName(v string) *KnowledgeBaseSearchChunk {
	s.FileName = &v
	return s
}

func (s *KnowledgeBaseSearchChunk) SetScore(v float64) *KnowledgeBaseSearchChunk {
	s.Score = &v
	return s
}

func (s *KnowledgeBaseSearchChunk) SetScores(v *KnowledgeBaseSearchChunkScores) *KnowledgeBaseSearchChunk {
	s.Scores = v
	return s
}

func (s *KnowledgeBaseSearchChunk) SetSourceLocation(v string) *KnowledgeBaseSearchChunk {
	s.SourceLocation = &v
	return s
}

func (s *KnowledgeBaseSearchChunk) SetTitlePath(v string) *KnowledgeBaseSearchChunk {
	s.TitlePath = &v
	return s
}

func (s *KnowledgeBaseSearchChunk) Validate() error {
	if s.Scores != nil {
		if err := s.Scores.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type KnowledgeBaseSearchChunkScores struct {
	// The score after hybrid search fusion (reciprocal rank fusion or weighted normalization, depending on the active fusion algorithm). Value range: [0, 1].
	//
	// example:
	//
	// 0.78
	Fusion *float64 `json:"Fusion,omitempty" xml:"Fusion,omitempty"`
	// The normalized score of keyword (full-text) search. Value range: [0, 1].
	//
	// example:
	//
	// 0.62
	Keyword *float64 `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The score assigned by the rerank model. Value range: [0, 1].
	//
	// example:
	//
	// 0.91
	Rerank *float64 `json:"Rerank,omitempty" xml:"Rerank,omitempty"`
	// The similarity score of vector retrieval. Value range: [0, 1].
	//
	// example:
	//
	// 0.85
	Vector *float64 `json:"Vector,omitempty" xml:"Vector,omitempty"`
}

func (s KnowledgeBaseSearchChunkScores) String() string {
	return dara.Prettify(s)
}

func (s KnowledgeBaseSearchChunkScores) GoString() string {
	return s.String()
}

func (s *KnowledgeBaseSearchChunkScores) GetFusion() *float64 {
	return s.Fusion
}

func (s *KnowledgeBaseSearchChunkScores) GetKeyword() *float64 {
	return s.Keyword
}

func (s *KnowledgeBaseSearchChunkScores) GetRerank() *float64 {
	return s.Rerank
}

func (s *KnowledgeBaseSearchChunkScores) GetVector() *float64 {
	return s.Vector
}

func (s *KnowledgeBaseSearchChunkScores) SetFusion(v float64) *KnowledgeBaseSearchChunkScores {
	s.Fusion = &v
	return s
}

func (s *KnowledgeBaseSearchChunkScores) SetKeyword(v float64) *KnowledgeBaseSearchChunkScores {
	s.Keyword = &v
	return s
}

func (s *KnowledgeBaseSearchChunkScores) SetRerank(v float64) *KnowledgeBaseSearchChunkScores {
	s.Rerank = &v
	return s
}

func (s *KnowledgeBaseSearchChunkScores) SetVector(v float64) *KnowledgeBaseSearchChunkScores {
	s.Vector = &v
	return s
}

func (s *KnowledgeBaseSearchChunkScores) Validate() error {
	return dara.Validate(s)
}
