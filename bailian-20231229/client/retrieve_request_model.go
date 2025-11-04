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
	EnableRewrite *bool     `json:"EnableRewrite,omitempty" xml:"EnableRewrite,omitempty"`
	Images        []*string `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	// The primary key ID of the knowledge base, which is the `Data.Id` parameter returned by the [CreateIndex](https://www.alibabacloud.com/help/en/model-studio/developer-reference/api-bailian-2023-12-29-createindex) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5pwe0m2g6t
	IndexId *string `json:"IndexId,omitempty" xml:"IndexId,omitempty"`
	// The input query prompt. The length and characters of the query are not limited.
	Query        *string                        `json:"Query,omitempty" xml:"Query,omitempty"`
	QueryHistory []*RetrieveRequestQueryHistory `json:"QueryHistory,omitempty" xml:"QueryHistory,omitempty" type:"Repeated"`
	// Ranking configurations.
	Rerank []*RetrieveRequestRerank `json:"Rerank,omitempty" xml:"Rerank,omitempty" type:"Repeated"`
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
	Rewrite []*RetrieveRequestRewrite `json:"Rewrite,omitempty" xml:"Rewrite,omitempty" type:"Repeated"`
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
	SearchFilters []map[string]*string `json:"SearchFilters,omitempty" xml:"SearchFilters,omitempty" type:"Repeated"`
	// The top K of keyword retrieval. Chunks that exactly match the keywords of the input text are retrieved from the knowledge base. This filters out irrelevant chunks and boosts accuracy. Valid values: 0 to 100. The sum of the `DenseSimilarityTopK` and `SparseSimilarityTopK` parameters must be less than or equal to 200.
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

type RetrieveRequestQueryHistory struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	Role    *string `json:"role,omitempty" xml:"role,omitempty"`
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
	// The name of the rank model. For more information, see [Create a knowledge base](https://www.alibabacloud.com/help/en/model-studio/user-guide/rag-knowledge-base). Valid values:
	//
	// 	- gte-rerank-hybrid: Recommended official model.
	//
	// 	- gte-rerank
	//
	// example:
	//
	// gte-rerank-hybrid
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
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

func (s *RetrieveRequestRerank) SetModelName(v string) *RetrieveRequestRerank {
	s.ModelName = &v
	return s
}

func (s *RetrieveRequestRerank) Validate() error {
	return dara.Validate(s)
}

type RetrieveRequestRewrite struct {
	// Conversation rewriting model name. The query rewriting model automatically adjusts the original prompt based on the context to improve retrieval performance. Valid value:
	//
	// 	- conv-rewrite-qwen-1.8b
	//
	// By default, this parameter is left empty, which means conv-rewrite-qwen-1.8b is used.
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
