// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetrievalKnowledgeBaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKnowledgeBaseId(v string) *RetrievalKnowledgeBaseRequest
	GetKnowledgeBaseId() *string
	SetQueryText(v string) *RetrievalKnowledgeBaseRequest
	GetQueryText() *string
	SetRegionId(v string) *RetrievalKnowledgeBaseRequest
	GetRegionId() *string
	SetRerankEnabled(v bool) *RetrievalKnowledgeBaseRequest
	GetRerankEnabled() *bool
	SetScoreThreshold(v float64) *RetrievalKnowledgeBaseRequest
	GetScoreThreshold() *float64
	SetTopK(v int32) *RetrievalKnowledgeBaseRequest
	GetTopK() *int32
}

type RetrievalKnowledgeBaseRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pkb-xxxxx
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 财报
	QueryText *string `json:"QueryText,omitempty" xml:"QueryText,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// true
	RerankEnabled *bool `json:"RerankEnabled,omitempty" xml:"RerankEnabled,omitempty"`
	// example:
	//
	// 0.7
	ScoreThreshold *float64 `json:"ScoreThreshold,omitempty" xml:"ScoreThreshold,omitempty"`
	// example:
	//
	// 5
	TopK *int32 `json:"TopK,omitempty" xml:"TopK,omitempty"`
}

func (s RetrievalKnowledgeBaseRequest) String() string {
	return dara.Prettify(s)
}

func (s RetrievalKnowledgeBaseRequest) GoString() string {
	return s.String()
}

func (s *RetrievalKnowledgeBaseRequest) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *RetrievalKnowledgeBaseRequest) GetQueryText() *string {
	return s.QueryText
}

func (s *RetrievalKnowledgeBaseRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RetrievalKnowledgeBaseRequest) GetRerankEnabled() *bool {
	return s.RerankEnabled
}

func (s *RetrievalKnowledgeBaseRequest) GetScoreThreshold() *float64 {
	return s.ScoreThreshold
}

func (s *RetrievalKnowledgeBaseRequest) GetTopK() *int32 {
	return s.TopK
}

func (s *RetrievalKnowledgeBaseRequest) SetKnowledgeBaseId(v string) *RetrievalKnowledgeBaseRequest {
	s.KnowledgeBaseId = &v
	return s
}

func (s *RetrievalKnowledgeBaseRequest) SetQueryText(v string) *RetrievalKnowledgeBaseRequest {
	s.QueryText = &v
	return s
}

func (s *RetrievalKnowledgeBaseRequest) SetRegionId(v string) *RetrievalKnowledgeBaseRequest {
	s.RegionId = &v
	return s
}

func (s *RetrievalKnowledgeBaseRequest) SetRerankEnabled(v bool) *RetrievalKnowledgeBaseRequest {
	s.RerankEnabled = &v
	return s
}

func (s *RetrievalKnowledgeBaseRequest) SetScoreThreshold(v float64) *RetrievalKnowledgeBaseRequest {
	s.ScoreThreshold = &v
	return s
}

func (s *RetrievalKnowledgeBaseRequest) SetTopK(v int32) *RetrievalKnowledgeBaseRequest {
	s.TopK = &v
	return s
}

func (s *RetrievalKnowledgeBaseRequest) Validate() error {
	return dara.Validate(s)
}
