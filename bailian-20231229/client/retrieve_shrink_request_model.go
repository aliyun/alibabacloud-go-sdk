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
	// The top K value for AISearch. The system generates a vector for the input text and retrieves the top K text chunks in the knowledge base that are most similar to the vector representation. Valid values: 0 to 100.
	//
	// The sum of DenseSimilarityTopK and SparseSimilarityTopK must be less than or equal to 200.
	//
	// Default value: 100.
	//
	// example:
	//
	// 100
	DenseSimilarityTopK *int32 `json:"DenseSimilarityTopK,omitempty" xml:"DenseSimilarityTopK,omitempty"`
	// Specifies whether to enable reranking. For more information, see [Knowledge base](https://help.aliyun.com/document_detail/2807740.html). Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// Default value: true.
	//
	// example:
	//
	// true
	EnableReranking *bool `json:"EnableReranking,omitempty" xml:"EnableReranking,omitempty"`
	// Specifies whether to enable <props="china">[multi-turn conversation rewriting](https://help.aliyun.com/model-studio/use-cases/rag-optimization#b7031e2ad6cji)<props="intl">[multi-turn conversation rewriting](https://www.alibabacloud.com/help/model-studio/use-cases/rag-optimization#b7031e2ad6cji).
	//
	// Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	EnableRewrite *bool   `json:"EnableRewrite,omitempty" xml:"EnableRewrite,omitempty"`
	ExtraShrink   *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// The URLs of the images passed in the query.
	ImagesShrink *string `json:"Images,omitempty" xml:"Images,omitempty"`
	// The ID of the knowledge base, which is the Data.Id returned by the CreateIndex operation.
	//
	// > - Make sure that the corresponding knowledge base is created and not deleted.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5pwe0mxxxx
	IndexId *string `json:"IndexId,omitempty" xml:"IndexId,omitempty"`
	// The input text, which is the original input prompt. The length and characters of Query are not limited.
	//
	// example:
	//
	// Introduction to Alibaba Cloud Model Studio
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// <props="china">[Multi-turn conversation rewriting](https://help.aliyun.com/model-studio/use-cases/rag-optimization#b7031e2ad6cji)<props="intl">[Multi-turn conversation rewriting](https://www.alibabacloud.com/help/model-studio/use-cases/rag-optimization#b7031e2ad6cji) allows you to pass in your own managed conversation history. This parameter takes effect only when EnableRewrite is set to **true**. Otherwise, it does not take effect even if you specify it.
	QueryHistoryShrink *string `json:"QueryHistory,omitempty" xml:"QueryHistory,omitempty"`
	// The reranking configuration.
	RerankShrink *string `json:"Rerank,omitempty" xml:"Rerank,omitempty"`
	// The similarity threshold. This threshold specifies the minimum similarity score allowed for recalled text chunks and is used to filter the text chunks returned by the ranking model. Only text chunks with a score greater than this value are recalled. Valid values: 0.01 to 1.00. This parameter has a higher priority than the similarity threshold configured for the knowledge base.
	//
	// If no specific value is specified, the similarity threshold configured for the knowledge base is used by default.
	//
	// example:
	//
	// 0.20
	RerankMinScore *float32 `json:"RerankMinScore,omitempty" xml:"RerankMinScore,omitempty"`
	// The top N text chunks to return after reranking. Valid values: 1 to 20. Default value: 5.
	//
	// example:
	//
	// 5
	RerankTopN *int32 `json:"RerankTopN,omitempty" xml:"RerankTopN,omitempty"`
	// The multi-turn conversation rewriting configuration.
	RewriteShrink *string `json:"Rewrite,omitempty" xml:"Rewrite,omitempty"`
	// Specifies whether to save the historical text chunk recall test data. Valid values:
	//
	// - true: Save.
	//
	// - false: Do not save.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	SaveRetrieverHistory *bool `json:"SaveRetrieverHistory,omitempty" xml:"SaveRetrieverHistory,omitempty"`
	// The personalized retrieval conditions, such as tags, configured by using SearchFilter to filter semantic retrieval results and exclude information irrelevant to the query.
	//
	// The filtering logic takes effect when the is_displayed_chunk_content parameter is set to true. For more information, see [Knowledge base SearchFilters](https://help.aliyun.com/document_detail/2869641.html).
	SearchFiltersShrink *string `json:"SearchFilters,omitempty" xml:"SearchFilters,omitempty"`
	// The top K value for keyword retrieval. The system retrieves chunks from the knowledge base that are an exact match for the keywords in the input text. This helps you filter out irrelevant text chunks and provides more accurate results.
	//
	// Valid values: 0 to 100.
	//
	// The sum of DenseSimilarityTopK and SparseSimilarityTopK must be less than or equal to 200.
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
