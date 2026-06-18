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
	SetExtraShrink(v string) *RetrieveShrinkRequest
	GetExtraShrink() *string
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
	// The number of top-K similar text chunks to retrieve using vector retrieval. This is achieved by generating a vector representation of the query and searching the knowledge base for the K text chunks with the most similar vectors. The value must be an integer from 0 to 100. The sum of `DenseSimilarityTopK` and `SparseSimilarityTopK` must not exceed 200.
	//
	// Default value: 100.
	//
	// example:
	//
	// 100
	DenseSimilarityTopK *int32 `json:"DenseSimilarityTopK,omitempty" xml:"DenseSimilarityTopK,omitempty"`
	// Specifies whether to enable reranking. For more information, see [Knowledge base](https://help.aliyun.com/document_detail/2807740.html). Valid values:
	//
	// - `true`: Enables reranking.
	//
	// - `false`: Disables reranking.
	//
	// Default value: `true`.
	//
	// example:
	//
	// true
	EnableReranking *bool `json:"EnableReranking,omitempty" xml:"EnableReranking,omitempty"`
	// Specifies whether to enable <props="china">[conversational query rewriting](https://help.aliyun.com/model-studio/use-cases/rag-optimization#b7031e2ad6cji)<props="intl">[conversational query rewriting](https://www.alibabacloud.com/help/model-studio/use-cases/rag-optimization#b7031e2ad6cji).
	//
	// Valid values:
	//
	// - `true`: Enables conversational query rewriting.
	//
	// - `false`: Disables conversational query rewriting.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	EnableRewrite *bool   `json:"EnableRewrite,omitempty" xml:"EnableRewrite,omitempty"`
	ExtraShrink   *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// The URLs of images to include in the query.
	ImagesShrink *string `json:"Images,omitempty" xml:"Images,omitempty"`
	// The ID of the knowledge base. This is the `Data.Id` value returned by the `CreateIndex` operation.
	//
	// > - Ensure the specified knowledge base exists and has not been deleted.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5pwe0mxxxx
	IndexId *string `json:"IndexId,omitempty" xml:"IndexId,omitempty"`
	// The query, which is the original user prompt. There are no limits on the length of the query.
	//
	// example:
	//
	// 阿里云百炼平台介绍
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// The conversation history, used for <props="china">[conversational query rewriting](https://help.aliyun.com/model-studio/use-cases/rag-optimization#b7031e2ad6cji)<props="intl">[conversational query rewriting](https://www.alibabacloud.com/help/model-studio/use-cases/rag-optimization#b7031e2ad6cji). This parameter is effective only when `EnableRewrite` is set to `true`.
	QueryHistoryShrink *string `json:"QueryHistory,omitempty" xml:"QueryHistory,omitempty"`
	// The reranking configurations.
	RerankShrink *string `json:"Rerank,omitempty" xml:"Rerank,omitempty"`
	// The similarity threshold for reranking. Only text chunks with a similarity score greater than this value are returned. The value must be between 0.01 and 1.00, inclusive. This parameter overrides the similarity threshold setting of the knowledge base.
	//
	// If not specified, the threshold configured for the knowledge base is used.
	//
	// example:
	//
	// 0.20
	RerankMinScore *float32 `json:"RerankMinScore,omitempty" xml:"RerankMinScore,omitempty"`
	// The number of top-ranked text chunks to return after reranking. The value must be an integer from 1 to 20. Default value: 5.
	//
	// example:
	//
	// 5
	RerankTopN *int32 `json:"RerankTopN,omitempty" xml:"RerankTopN,omitempty"`
	// Configuration for conversational query rewriting.
	RewriteShrink *string `json:"Rewrite,omitempty" xml:"Rewrite,omitempty"`
	// Specifies whether to save the retrieval history for testing purposes. Valid values:
	//
	// - `true`: Saves the retrieval history.
	//
	// - `false`: Does not save the retrieval history.
	//
	// Default value: `false`.
	//
	// example:
	//
	// false
	SaveRetrieverHistory *bool `json:"SaveRetrieverHistory,omitempty" xml:"SaveRetrieverHistory,omitempty"`
	// Specifies custom retrieval conditions, such as tags, to filter semantic retrieval results and exclude irrelevant information.
	//
	// The filtering logic is applied only when the `is_displayed_chunk_content` parameter is set to `true`. For more information, see [SearchFilters for a knowledge base](https://help.aliyun.com/document_detail/2869641.html).
	SearchFiltersShrink *string `json:"SearchFilters,omitempty" xml:"SearchFilters,omitempty"`
	// The number of top-K text chunks to retrieve using keyword retrieval. This feature finds text chunks in the knowledge base that exactly match the keywords in the query. It helps filter out irrelevant text chunks and provide more accurate results.
	//
	// The value must be an integer from 0 to 100.
	//
	// The sum of `DenseSimilarityTopK` and `SparseSimilarityTopK` must not exceed 200.
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

func (s *RetrieveShrinkRequest) GetExtraShrink() *string {
	return s.ExtraShrink
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

func (s *RetrieveShrinkRequest) SetExtraShrink(v string) *RetrieveShrinkRequest {
	s.ExtraShrink = &v
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
