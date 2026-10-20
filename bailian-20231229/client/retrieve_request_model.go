// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetrieveRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDenseSimilarityTopK(v int32) *RetrieveRequest
	GetDenseSimilarityTopK() *int32
	SetEnableReranking(v bool) *RetrieveRequest
	GetEnableReranking() *bool
	SetEnableRewrite(v bool) *RetrieveRequest
	GetEnableRewrite() *bool
	SetExtra(v *RetrieveRequestExtra) *RetrieveRequest
	GetExtra() *RetrieveRequestExtra
	SetImages(v []*string) *RetrieveRequest
	GetImages() []*string
	SetIndexId(v string) *RetrieveRequest
	GetIndexId() *string
	SetQuery(v string) *RetrieveRequest
	GetQuery() *string
	SetQueryHistory(v []*RetrieveRequestQueryHistory) *RetrieveRequest
	GetQueryHistory() []*RetrieveRequestQueryHistory
	SetRerank(v []*RetrieveRequestRerank) *RetrieveRequest
	GetRerank() []*RetrieveRequestRerank
	SetRerankMinScore(v float32) *RetrieveRequest
	GetRerankMinScore() *float32
	SetRerankTopN(v int32) *RetrieveRequest
	GetRerankTopN() *int32
	SetRewrite(v []*RetrieveRequestRewrite) *RetrieveRequest
	GetRewrite() []*RetrieveRequestRewrite
	SetSaveRetrieverHistory(v bool) *RetrieveRequest
	GetSaveRetrieverHistory() *bool
	SetSearchFilters(v []map[string]*string) *RetrieveRequest
	GetSearchFilters() []map[string]*string
	SetSparseSimilarityTopK(v int32) *RetrieveRequest
	GetSparseSimilarityTopK() *int32
}

type RetrieveRequest struct {
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
	EnableRewrite *bool                 `json:"EnableRewrite,omitempty" xml:"EnableRewrite,omitempty"`
	Extra         *RetrieveRequestExtra `json:"Extra,omitempty" xml:"Extra,omitempty" type:"Struct"`
	// The URLs of the images passed in the query.
	Images []*string `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
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
	QueryHistory []*RetrieveRequestQueryHistory `json:"QueryHistory,omitempty" xml:"QueryHistory,omitempty" type:"Repeated"`
	// The reranking configuration.
	Rerank []*RetrieveRequestRerank `json:"Rerank,omitempty" xml:"Rerank,omitempty" type:"Repeated"`
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
	Rewrite []*RetrieveRequestRewrite `json:"Rewrite,omitempty" xml:"Rewrite,omitempty" type:"Repeated"`
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
	SearchFilters []map[string]*string `json:"SearchFilters,omitempty" xml:"SearchFilters,omitempty" type:"Repeated"`
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

func (s RetrieveRequest) String() string {
	return dara.Prettify(s)
}

func (s RetrieveRequest) GoString() string {
	return s.String()
}

func (s *RetrieveRequest) GetDenseSimilarityTopK() *int32 {
	return s.DenseSimilarityTopK
}

func (s *RetrieveRequest) GetEnableReranking() *bool {
	return s.EnableReranking
}

func (s *RetrieveRequest) GetEnableRewrite() *bool {
	return s.EnableRewrite
}

func (s *RetrieveRequest) GetExtra() *RetrieveRequestExtra {
	return s.Extra
}

func (s *RetrieveRequest) GetImages() []*string {
	return s.Images
}

func (s *RetrieveRequest) GetIndexId() *string {
	return s.IndexId
}

func (s *RetrieveRequest) GetQuery() *string {
	return s.Query
}

func (s *RetrieveRequest) GetQueryHistory() []*RetrieveRequestQueryHistory {
	return s.QueryHistory
}

func (s *RetrieveRequest) GetRerank() []*RetrieveRequestRerank {
	return s.Rerank
}

func (s *RetrieveRequest) GetRerankMinScore() *float32 {
	return s.RerankMinScore
}

func (s *RetrieveRequest) GetRerankTopN() *int32 {
	return s.RerankTopN
}

func (s *RetrieveRequest) GetRewrite() []*RetrieveRequestRewrite {
	return s.Rewrite
}

func (s *RetrieveRequest) GetSaveRetrieverHistory() *bool {
	return s.SaveRetrieverHistory
}

func (s *RetrieveRequest) GetSearchFilters() []map[string]*string {
	return s.SearchFilters
}

func (s *RetrieveRequest) GetSparseSimilarityTopK() *int32 {
	return s.SparseSimilarityTopK
}

func (s *RetrieveRequest) SetDenseSimilarityTopK(v int32) *RetrieveRequest {
	s.DenseSimilarityTopK = &v
	return s
}

func (s *RetrieveRequest) SetEnableReranking(v bool) *RetrieveRequest {
	s.EnableReranking = &v
	return s
}

func (s *RetrieveRequest) SetEnableRewrite(v bool) *RetrieveRequest {
	s.EnableRewrite = &v
	return s
}

func (s *RetrieveRequest) SetExtra(v *RetrieveRequestExtra) *RetrieveRequest {
	s.Extra = v
	return s
}

func (s *RetrieveRequest) SetImages(v []*string) *RetrieveRequest {
	s.Images = v
	return s
}

func (s *RetrieveRequest) SetIndexId(v string) *RetrieveRequest {
	s.IndexId = &v
	return s
}

func (s *RetrieveRequest) SetQuery(v string) *RetrieveRequest {
	s.Query = &v
	return s
}

func (s *RetrieveRequest) SetQueryHistory(v []*RetrieveRequestQueryHistory) *RetrieveRequest {
	s.QueryHistory = v
	return s
}

func (s *RetrieveRequest) SetRerank(v []*RetrieveRequestRerank) *RetrieveRequest {
	s.Rerank = v
	return s
}

func (s *RetrieveRequest) SetRerankMinScore(v float32) *RetrieveRequest {
	s.RerankMinScore = &v
	return s
}

func (s *RetrieveRequest) SetRerankTopN(v int32) *RetrieveRequest {
	s.RerankTopN = &v
	return s
}

func (s *RetrieveRequest) SetRewrite(v []*RetrieveRequestRewrite) *RetrieveRequest {
	s.Rewrite = v
	return s
}

func (s *RetrieveRequest) SetSaveRetrieverHistory(v bool) *RetrieveRequest {
	s.SaveRetrieverHistory = &v
	return s
}

func (s *RetrieveRequest) SetSearchFilters(v []map[string]*string) *RetrieveRequest {
	s.SearchFilters = v
	return s
}

func (s *RetrieveRequest) SetSparseSimilarityTopK(v int32) *RetrieveRequest {
	s.SparseSimilarityTopK = &v
	return s
}

func (s *RetrieveRequest) Validate() error {
	if s.Extra != nil {
		if err := s.Extra.Validate(); err != nil {
			return err
		}
	}
	if s.QueryHistory != nil {
		for _, item := range s.QueryHistory {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Rerank != nil {
		for _, item := range s.Rerank {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Rewrite != nil {
		for _, item := range s.Rewrite {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RetrieveRequestExtra struct {
	UniqueId *string `json:"uniqueId,omitempty" xml:"uniqueId,omitempty"`
}

func (s RetrieveRequestExtra) String() string {
	return dara.Prettify(s)
}

func (s RetrieveRequestExtra) GoString() string {
	return s.String()
}

func (s *RetrieveRequestExtra) GetUniqueId() *string {
	return s.UniqueId
}

func (s *RetrieveRequestExtra) SetUniqueId(v string) *RetrieveRequestExtra {
	s.UniqueId = &v
	return s
}

func (s *RetrieveRequestExtra) Validate() error {
	return dara.Validate(s)
}

type RetrieveRequestQueryHistory struct {
	// The question or answer content of the corresponding role.
	//
	// example:
	//
	// A piece of text
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The role.
	//
	// Valid values:
	//
	// - user: If this role is passed, content indicates the text that you input to the Alibaba Cloud Model Studio application.
	//
	// - assistant: If this role is passed, content indicates the reply from the Alibaba Cloud Model Studio application.
	//
	// example:
	//
	// user
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s RetrieveRequestQueryHistory) String() string {
	return dara.Prettify(s)
}

func (s RetrieveRequestQueryHistory) GoString() string {
	return s.String()
}

func (s *RetrieveRequestQueryHistory) GetContent() *string {
	return s.Content
}

func (s *RetrieveRequestQueryHistory) GetRole() *string {
	return s.Role
}

func (s *RetrieveRequestQueryHistory) SetContent(v string) *RetrieveRequestQueryHistory {
	s.Content = &v
	return s
}

func (s *RetrieveRequestQueryHistory) SetRole(v string) *RetrieveRequestQueryHistory {
	s.Role = &v
	return s
}

func (s *RetrieveRequestQueryHistory) Validate() error {
	return dara.Validate(s)
}

type RetrieveRequestRerank struct {
	// The specified reranking model. The value specified here overrides the reranking model selected when the knowledge base was created. Valid values:
	//
	// <props="china">
	//
	// - qwen3-rerank-hybrid: qwen3-rerank (hybrid) reranking.
	//
	// - qwen3-rerank: qwen3-rerank reranking.
	//
	// - gte-rerank-hybrid: gte-rerank (hybrid) reranking.
	//
	// - gte-rerank: gte-rerank reranking.
	//
	//
	//
	// <props="intl">
	//
	// - gte-rerank-hybrid: Official reranking.
	//
	// - gte-rerank: gte-rerank reranking.
	//
	//
	//
	//
	//
	//
	// Default value: null. The reranking model selected when the knowledge base was created is used.
	//
	//
	// <props="china">
	//
	// > If you only need semantic reranking, use qwen3-rerank. If you need both semantic reranking and text matching features to ensure relevance, use qwen3-rerank-hybrid.
	//
	// >
	//
	//
	//
	//
	//
	//
	// <props="intl">
	//
	// > If you only need semantic reranking, use gte-rerank. If you need both semantic reranking and text matching features to ensure relevance, use gte-rerank-hybrid.
	//
	// >
	//
	//
	//
	//
	//
	//
	// <props="china">
	//
	// > gte-rerank-hybrid and gte-rerank will no longer be updated and are not recommended.
	//
	// >
	//
	// example:
	//
	// gte-rerank-hybrid
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// <props="intl">This parameter is not available. Do not specify this parameter.
	//
	// <props="china">A natural language instruction used to fine-tune the behavior of the reranking model.
	//
	// <notice>This parameter takes effect only when rerank_mode is set to custom.</notice>
	//
	// example:
	//
	// Given a web search query, retrieve relevant passages that answer the query.
	RerankInstruct *string `json:"RerankInstruct,omitempty" xml:"RerankInstruct,omitempty"`
	// <props="china">
	//
	// The instruction intervention mode of the reranking model, which determines its scoring preference.
	//
	// **Valid values:**
	//
	// - **qa**: (Default) Q&A mode. The model tends to assign higher scores to candidates that can directly answer the query. Recommended for Q&A scenarios.
	//
	// - **similar**: Similarity mode. The model tends to assign higher scores to candidates that are highly consistent with the query content. Recommended for matching and retrieval scenarios.
	//
	// - **custom**: Custom mode. The reranking behavior of the model is determined by the instruction in the rerank_instruct parameter.
	//
	//
	//
	// <props="intl">This parameter is not available. Do not specify this parameter.
	//
	// example:
	//
	// qa
	RerankMode *string `json:"RerankMode,omitempty" xml:"RerankMode,omitempty"`
}

func (s RetrieveRequestRerank) String() string {
	return dara.Prettify(s)
}

func (s RetrieveRequestRerank) GoString() string {
	return s.String()
}

func (s *RetrieveRequestRerank) GetModelName() *string {
	return s.ModelName
}

func (s *RetrieveRequestRerank) GetRerankInstruct() *string {
	return s.RerankInstruct
}

func (s *RetrieveRequestRerank) GetRerankMode() *string {
	return s.RerankMode
}

func (s *RetrieveRequestRerank) SetModelName(v string) *RetrieveRequestRerank {
	s.ModelName = &v
	return s
}

func (s *RetrieveRequestRerank) SetRerankInstruct(v string) *RetrieveRequestRerank {
	s.RerankInstruct = &v
	return s
}

func (s *RetrieveRequestRerank) SetRerankMode(v string) *RetrieveRequestRerank {
	s.RerankMode = &v
	return s
}

func (s *RetrieveRequestRerank) Validate() error {
	return dara.Validate(s)
}

type RetrieveRequestRewrite struct {
	// The name of the multi-turn conversation rewriting model. The model automatically adjusts the original input prompt (user question) based on the conversation context to improve retrieval performance. Valid values:
	//
	// - conv-rewrite-qwen-1.8b: The conv-rewrite-qwen-1.8b model. Currently, only this model is supported.
	//
	// Default value: null. The conv-rewrite-qwen-1.8b model is used.
	//
	// example:
	//
	// conv-rewrite-qwen-1.8b
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
}

func (s RetrieveRequestRewrite) String() string {
	return dara.Prettify(s)
}

func (s RetrieveRequestRewrite) GoString() string {
	return s.String()
}

func (s *RetrieveRequestRewrite) GetModelName() *string {
	return s.ModelName
}

func (s *RetrieveRequestRewrite) SetModelName(v string) *RetrieveRequestRewrite {
	s.ModelName = &v
	return s
}

func (s *RetrieveRequestRewrite) Validate() error {
	return dara.Validate(s)
}
