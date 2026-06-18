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
	EnableRewrite *bool                 `json:"EnableRewrite,omitempty" xml:"EnableRewrite,omitempty"`
	Extra         *RetrieveRequestExtra `json:"Extra,omitempty" xml:"Extra,omitempty" type:"Struct"`
	// The URLs of images to include in the query.
	Images []*string `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
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
	QueryHistory []*RetrieveRequestQueryHistory `json:"QueryHistory,omitempty" xml:"QueryHistory,omitempty" type:"Repeated"`
	// The reranking configurations.
	Rerank []*RetrieveRequestRerank `json:"Rerank,omitempty" xml:"Rerank,omitempty" type:"Repeated"`
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
	Rewrite []*RetrieveRequestRewrite `json:"Rewrite,omitempty" xml:"Rewrite,omitempty" type:"Repeated"`
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
	SearchFilters []map[string]*string `json:"SearchFilters,omitempty" xml:"SearchFilters,omitempty" type:"Repeated"`
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
	// The content of the message for the specified `role`.
	//
	// example:
	//
	// 代表一段文本。
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The role of the entity that sent the message.
	//
	// Valid values:
	//
	// - `user`: Indicates that the `content` is from the end user.
	//
	// - `assistant`: Indicates that the `content` is a response from the Model Studio application.
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
	// The reranking model to use. Specifying a model here overrides the default model selected when the knowledge base was created. Valid values:
	//
	// <props="china">
	//
	// - `qwen3-rerank-hybrid`: Performs reranking by using the qwen3-rerank (hybrid) model.
	//
	// - `qwen3-rerank`: Performs reranking by using the qwen3-rerank model.
	//
	// - `gte-rerank-hybrid`: Performs reranking by using the gte-rerank (hybrid) model.
	//
	// - `gte-rerank`: Performs reranking by using the gte-rerank model.
	//
	//
	//
	// <props="intl">
	//
	// - `gte-rerank-hybrid`: Performs reranking by using the gte-rerank (hybrid) model.
	//
	// - `gte-rerank`: Performs reranking by using the gte-rerank model.
	//
	//
	//
	// If you do not specify this parameter, the model configured for the knowledge base is used.
	//
	// <props="china">
	//
	// > Use `qwen3-rerank` for semantic ranking only. We recommend `qwen3-rerank-hybrid` if you require both semantic ranking and text matching features for higher relevance.
	//
	//
	//
	// <props="intl">
	//
	// > Use `gte-rerank` for semantic ranking only. We recommend `gte-rerank-hybrid` if you require both semantic ranking and text matching features for higher relevance.
	//
	//
	//
	// <props="china">
	//
	// > The `gte-rerank-hybrid` and `gte-rerank` models are no longer updated and are not recommended for use.
	//
	// example:
	//
	// gte-rerank-hybrid
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// <props="intl">
	//
	// This parameter is not yet available. Do not specify a value for it.
	//
	//
	//
	// <props="china">
	//
	// A natural language instruction to fine-tune the behavior of the reranking model.
	//
	// 	Notice:
	//
	// This parameter takes effect only when `RerankMode` is set to `custom`.
	RerankInstruct *string `json:"RerankInstruct,omitempty" xml:"RerankInstruct,omitempty"`
	// <props="china">
	//
	// Specifies the instruction intervention mode for the reranking model. This mode determines the model\\"s scoring preference.
	//
	// **Valid values:**
	//
	// - `qa`: (Default) Q\\&A mode. The model assigns higher scores to candidates that directly answer the query. Recommended for Q\\&A scenarios.
	//
	// - `similar`: Similarity mode. The model assigns higher scores to candidates with high content similarity to the query. Recommended for match-based retrieval scenarios.
	//
	// - `custom`: Custom mode. The model\\"s ranking behavior is determined by the instructions in the `RerankInstruct` parameter.
	//
	//
	//
	// <props="intl">
	//
	// This parameter is not yet available. Do not specify a value for it.
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
	// Specifies the model for conversational query rewriting, which automatically rewrites the original query based on the conversation context to improve retrieval results. Valid value:
	//
	// - `conv-rewrite-qwen-1.8b`: The only model currently supported for this feature.
	//
	// If this parameter is not specified, `conv-rewrite-qwen-1.8b` is used by default.
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
