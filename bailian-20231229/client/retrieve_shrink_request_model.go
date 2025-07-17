// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetrieveShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDenseSimilarityTopK(v int32) *RetrieveShrinkRequest
	GetDenseSimilarityTopK() *int32
	SetEnableReranking(v bool) *RetrieveShrinkRequest
	GetEnableReranking() *bool
	SetEnableRewrite(v bool) *RetrieveShrinkRequest
	GetEnableRewrite() *bool
	SetImagesShrink(v string) *RetrieveShrinkRequest
	GetImagesShrink() *string
	SetIndexId(v string) *RetrieveShrinkRequest
	GetIndexId() *string
	SetQuery(v string) *RetrieveShrinkRequest
	GetQuery() *string
	SetQueryHistoryShrink(v string) *RetrieveShrinkRequest
	GetQueryHistoryShrink() *string
	SetRerankShrink(v string) *RetrieveShrinkRequest
	GetRerankShrink() *string
	SetRerankMinScore(v float32) *RetrieveShrinkRequest
	GetRerankMinScore() *float32
	SetRerankTopN(v int32) *RetrieveShrinkRequest
	GetRerankTopN() *int32
	SetRewriteShrink(v string) *RetrieveShrinkRequest
	GetRewriteShrink() *string
	SetSaveRetrieverHistory(v bool) *RetrieveShrinkRequest
	GetSaveRetrieverHistory() *bool
	SetSearchFiltersShrink(v string) *RetrieveShrinkRequest
	GetSearchFiltersShrink() *string
	SetSparseSimilarityTopK(v int32) *RetrieveShrinkRequest
	GetSparseSimilarityTopK() *int32
}

type RetrieveShrinkRequest struct {
	// Vector retrieval top K. After generating vectors based on input text, the top K chunks in the knowledge base that are most similar to the vector representation of the input text are retrieved. Valid values: 0 to 100. The sum of the `DenseSimilarityTopK` and `SparseSimilarityTopK` parameters must be less than or equal to 200.
	//
	// Default value: 100.
	//
	// example:
	//
	// 100
	DenseSimilarityTopK *int32 `json:"DenseSimilarityTopK,omitempty" xml:"DenseSimilarityTopK,omitempty"`
	// Specifies whether to enable reranking. For more information, see [Create a knowledge base](https://www.alibabacloud.com/help/en/model-studio/user-guide/rag-knowledge-base). Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: true.
	//
	// example:
	//
	// true
	EnableReranking *bool `json:"EnableReranking,omitempty" xml:"EnableReranking,omitempty"`
	// Specifies whether to enable multi-round conversation rewriting. For more information, see [Create a knowledge base](https://www.alibabacloud.com/help/en/model-studio/user-guide/rag-knowledge-base). Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: false.
	//
	// example:
	//
	// false
	EnableRewrite *bool   `json:"EnableRewrite,omitempty" xml:"EnableRewrite,omitempty"`
	ImagesShrink  *string `json:"Images,omitempty" xml:"Images,omitempty"`
	// The primary key ID of the knowledge base, which is the `Data.Id` parameter returned by the [CreateIndex](https://www.alibabacloud.com/help/en/model-studio/developer-reference/api-bailian-2023-12-29-createindex) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5pwe0m2g6t
	IndexId *string `json:"IndexId,omitempty" xml:"IndexId,omitempty"`
	// The input query prompt. The length and characters of the query are not limited.
	Query              *string `json:"Query,omitempty" xml:"Query,omitempty"`
	QueryHistoryShrink *string `json:"QueryHistory,omitempty" xml:"QueryHistory,omitempty"`
	// Ranking configurations.
	RerankShrink *string `json:"Rerank,omitempty" xml:"Rerank,omitempty"`
	// Similarity Threshold The lowest similarity score of chunks that can be returned. This parameter is used to filter text chunks returned by the rank model. For more information, see [Create a knowledge base](https://www.alibabacloud.com/help/en/model-studio/user-guide/rag-knowledge-base). Valid values: [0.01-1.00]. The priority of this parameter is greater than the similarity threshold configured for the knowledge base.
	//
	// By default, this parameter is left empty. In this case, the similarity threshold of the knowledge base is used.
	//
	// example:
	//
	// 0.20
	RerankMinScore *float32 `json:"RerankMinScore,omitempty" xml:"RerankMinScore,omitempty"`
	// The top N return data after reranking. Valid values: 1 to 20. Default value: 5.
	//
	// example:
	//
	// 5
	RerankTopN *int32 `json:"RerankTopN,omitempty" xml:"RerankTopN,omitempty"`
	// Conversation rewriting configurations.
	RewriteShrink *string `json:"Rewrite,omitempty" xml:"Rewrite,omitempty"`
	// Specifies whether to save the retrieve test history. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: false.
	//
	// example:
	//
	// false
	SaveRetrieverHistory *bool `json:"SaveRetrieverHistory,omitempty" xml:"SaveRetrieverHistory,omitempty"`
	// Specifies complex filter conditions. For more information about the syntax of SearchFilters, see the SearchFilter syntax section of this topic.
	SearchFiltersShrink *string `json:"SearchFilters,omitempty" xml:"SearchFilters,omitempty"`
	// The top K of keyword retrieval. Chunks that exactly match the keywords of the input text are retrieved from the knowledge base. This filters out irrelevant chunks and boosts accuracy. Valid values: 0 to 100. The sum of the `DenseSimilarityTopK` and `SparseSimilarityTopK` parameters must be less than or equal to 200.
	//
	// Default value: 100.
	//
	// example:
	//
	// 100
	SparseSimilarityTopK *int32 `json:"SparseSimilarityTopK,omitempty" xml:"SparseSimilarityTopK,omitempty"`
}

func (s RetrieveShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RetrieveShrinkRequest) GoString() string {
	return s.String()
}

func (s *RetrieveShrinkRequest) GetDenseSimilarityTopK() *int32 {
	return s.DenseSimilarityTopK
}

func (s *RetrieveShrinkRequest) GetEnableReranking() *bool {
	return s.EnableReranking
}

func (s *RetrieveShrinkRequest) GetEnableRewrite() *bool {
	return s.EnableRewrite
}

func (s *RetrieveShrinkRequest) GetImagesShrink() *string {
	return s.ImagesShrink
}

func (s *RetrieveShrinkRequest) GetIndexId() *string {
	return s.IndexId
}

func (s *RetrieveShrinkRequest) GetQuery() *string {
	return s.Query
}

func (s *RetrieveShrinkRequest) GetQueryHistoryShrink() *string {
	return s.QueryHistoryShrink
}

func (s *RetrieveShrinkRequest) GetRerankShrink() *string {
	return s.RerankShrink
}

func (s *RetrieveShrinkRequest) GetRerankMinScore() *float32 {
	return s.RerankMinScore
}

func (s *RetrieveShrinkRequest) GetRerankTopN() *int32 {
	return s.RerankTopN
}

func (s *RetrieveShrinkRequest) GetRewriteShrink() *string {
	return s.RewriteShrink
}

func (s *RetrieveShrinkRequest) GetSaveRetrieverHistory() *bool {
	return s.SaveRetrieverHistory
}

func (s *RetrieveShrinkRequest) GetSearchFiltersShrink() *string {
	return s.SearchFiltersShrink
}

func (s *RetrieveShrinkRequest) GetSparseSimilarityTopK() *int32 {
	return s.SparseSimilarityTopK
}

func (s *RetrieveShrinkRequest) SetDenseSimilarityTopK(v int32) *RetrieveShrinkRequest {
	s.DenseSimilarityTopK = &v
	return s
}

func (s *RetrieveShrinkRequest) SetEnableReranking(v bool) *RetrieveShrinkRequest {
	s.EnableReranking = &v
	return s
}

func (s *RetrieveShrinkRequest) SetEnableRewrite(v bool) *RetrieveShrinkRequest {
	s.EnableRewrite = &v
	return s
}

func (s *RetrieveShrinkRequest) SetImagesShrink(v string) *RetrieveShrinkRequest {
	s.ImagesShrink = &v
	return s
}

func (s *RetrieveShrinkRequest) SetIndexId(v string) *RetrieveShrinkRequest {
	s.IndexId = &v
	return s
}

func (s *RetrieveShrinkRequest) SetQuery(v string) *RetrieveShrinkRequest {
	s.Query = &v
	return s
}

func (s *RetrieveShrinkRequest) SetQueryHistoryShrink(v string) *RetrieveShrinkRequest {
	s.QueryHistoryShrink = &v
	return s
}

func (s *RetrieveShrinkRequest) SetRerankShrink(v string) *RetrieveShrinkRequest {
	s.RerankShrink = &v
	return s
}

func (s *RetrieveShrinkRequest) SetRerankMinScore(v float32) *RetrieveShrinkRequest {
	s.RerankMinScore = &v
	return s
}

func (s *RetrieveShrinkRequest) SetRerankTopN(v int32) *RetrieveShrinkRequest {
	s.RerankTopN = &v
	return s
}

func (s *RetrieveShrinkRequest) SetRewriteShrink(v string) *RetrieveShrinkRequest {
	s.RewriteShrink = &v
	return s
}

func (s *RetrieveShrinkRequest) SetSaveRetrieverHistory(v bool) *RetrieveShrinkRequest {
	s.SaveRetrieverHistory = &v
	return s
}

func (s *RetrieveShrinkRequest) SetSearchFiltersShrink(v string) *RetrieveShrinkRequest {
	s.SearchFiltersShrink = &v
	return s
}

func (s *RetrieveShrinkRequest) SetSparseSimilarityTopK(v int32) *RetrieveShrinkRequest {
	s.SparseSimilarityTopK = &v
	return s
}

func (s *RetrieveShrinkRequest) Validate() error {
	return dara.Validate(s)
}
