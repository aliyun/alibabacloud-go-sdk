// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChunkItem interface {
	dara.Model
	String() string
	GoString() string
	SetChunkContent(v string) *ChunkItem
	GetChunkContent() *string
	SetChunkId(v string) *ChunkItem
	GetChunkId() *string
	SetDocElsIds(v []*string) *ChunkItem
	GetDocElsIds() []*string
	SetDocId(v string) *ChunkItem
	GetDocId() *string
	SetDocName(v string) *ChunkItem
	GetDocName() *string
	SetDocUrl(v string) *ChunkItem
	GetDocUrl() *string
	SetRerankScore(v float32) *ChunkItem
	GetRerankScore() *float32
	SetScore(v float32) *ChunkItem
	GetScore() *float32
	SetWeightedScore(v float32) *ChunkItem
	GetWeightedScore() *float32
}

type ChunkItem struct {
	ChunkContent  *string   `json:"chunkContent,omitempty" xml:"chunkContent,omitempty"`
	ChunkId       *string   `json:"chunkId,omitempty" xml:"chunkId,omitempty"`
	DocElsIds     []*string `json:"docElsIds,omitempty" xml:"docElsIds,omitempty" type:"Repeated"`
	DocId         *string   `json:"docId,omitempty" xml:"docId,omitempty"`
	DocName       *string   `json:"docName,omitempty" xml:"docName,omitempty"`
	DocUrl        *string   `json:"docUrl,omitempty" xml:"docUrl,omitempty"`
	RerankScore   *float32  `json:"rerankScore,omitempty" xml:"rerankScore,omitempty"`
	Score         *float32  `json:"score,omitempty" xml:"score,omitempty"`
	WeightedScore *float32  `json:"weightedScore,omitempty" xml:"weightedScore,omitempty"`
}

func (s ChunkItem) String() string {
	return dara.Prettify(s)
}

func (s ChunkItem) GoString() string {
	return s.String()
}

func (s *ChunkItem) GetChunkContent() *string {
	return s.ChunkContent
}

func (s *ChunkItem) GetChunkId() *string {
	return s.ChunkId
}

func (s *ChunkItem) GetDocElsIds() []*string {
	return s.DocElsIds
}

func (s *ChunkItem) GetDocId() *string {
	return s.DocId
}

func (s *ChunkItem) GetDocName() *string {
	return s.DocName
}

func (s *ChunkItem) GetDocUrl() *string {
	return s.DocUrl
}

func (s *ChunkItem) GetRerankScore() *float32 {
	return s.RerankScore
}

func (s *ChunkItem) GetScore() *float32 {
	return s.Score
}

func (s *ChunkItem) GetWeightedScore() *float32 {
	return s.WeightedScore
}

func (s *ChunkItem) SetChunkContent(v string) *ChunkItem {
	s.ChunkContent = &v
	return s
}

func (s *ChunkItem) SetChunkId(v string) *ChunkItem {
	s.ChunkId = &v
	return s
}

func (s *ChunkItem) SetDocElsIds(v []*string) *ChunkItem {
	s.DocElsIds = v
	return s
}

func (s *ChunkItem) SetDocId(v string) *ChunkItem {
	s.DocId = &v
	return s
}

func (s *ChunkItem) SetDocName(v string) *ChunkItem {
	s.DocName = &v
	return s
}

func (s *ChunkItem) SetDocUrl(v string) *ChunkItem {
	s.DocUrl = &v
	return s
}

func (s *ChunkItem) SetRerankScore(v float32) *ChunkItem {
	s.RerankScore = &v
	return s
}

func (s *ChunkItem) SetScore(v float32) *ChunkItem {
	s.Score = &v
	return s
}

func (s *ChunkItem) SetWeightedScore(v float32) *ChunkItem {
	s.WeightedScore = &v
	return s
}

func (s *ChunkItem) Validate() error {
	return dara.Validate(s)
}
