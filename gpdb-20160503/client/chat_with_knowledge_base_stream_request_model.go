// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatWithKnowledgeBaseStreamRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ChatWithKnowledgeBaseStreamRequest
	GetDBInstanceId() *string
	SetIncludeKnowledgeBaseResults(v bool) *ChatWithKnowledgeBaseStreamRequest
	GetIncludeKnowledgeBaseResults() *bool
	SetKnowledgeParams(v *ChatWithKnowledgeBaseStreamRequestKnowledgeParams) *ChatWithKnowledgeBaseStreamRequest
	GetKnowledgeParams() *ChatWithKnowledgeBaseStreamRequestKnowledgeParams
	SetModelParams(v *ChatWithKnowledgeBaseStreamRequestModelParams) *ChatWithKnowledgeBaseStreamRequest
	GetModelParams() *ChatWithKnowledgeBaseStreamRequestModelParams
	SetOwnerId(v int64) *ChatWithKnowledgeBaseStreamRequest
	GetOwnerId() *int64
	SetPromptParams(v string) *ChatWithKnowledgeBaseStreamRequest
	GetPromptParams() *string
	SetRegionId(v string) *ChatWithKnowledgeBaseStreamRequest
	GetRegionId() *string
}

type ChatWithKnowledgeBaseStreamRequest struct {
	// The instance ID.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the IDs of all AnalyticDB for PostgreSQL instances in a specified region.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// Specifies whether to include the retrieved knowledge base results in the response. Default value: `false`.
	//
	// example:
	//
	// false
	IncludeKnowledgeBaseResults *bool `json:"IncludeKnowledgeBaseResults,omitempty" xml:"IncludeKnowledgeBaseResults,omitempty"`
	// Parameters for knowledge retrieval. If omitted, the API performs a chat-only operation.
	KnowledgeParams *ChatWithKnowledgeBaseStreamRequestKnowledgeParams `json:"KnowledgeParams,omitempty" xml:"KnowledgeParams,omitempty" type:"Struct"`
	// An object that contains parameters for the Large Language Model (LLM) call.
	//
	// This parameter is required.
	ModelParams *ChatWithKnowledgeBaseStreamRequestModelParams `json:"ModelParams,omitempty" xml:"ModelParams,omitempty" type:"Struct"`
	OwnerId     *int64                                         `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// A template for the system prompt. It must include placeholders such as `{{text_chunks}}`, `{{user_system_prompt}}`, `{{graph_entities}}`, and `{{graph_relations}}`. If omitted, no custom prompt template is applied.
	//
	// example:
	//
	// "参考以下知识回答问题:{{ text_chunks }}"
	PromptParams *string `json:"PromptParams,omitempty" xml:"PromptParams,omitempty"`
	// The instance\\"s region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ChatWithKnowledgeBaseStreamRequest) String() string {
	return dara.Prettify(s)
}

func (s ChatWithKnowledgeBaseStreamRequest) GoString() string {
	return s.String()
}

func (s *ChatWithKnowledgeBaseStreamRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ChatWithKnowledgeBaseStreamRequest) GetIncludeKnowledgeBaseResults() *bool {
	return s.IncludeKnowledgeBaseResults
}

func (s *ChatWithKnowledgeBaseStreamRequest) GetKnowledgeParams() *ChatWithKnowledgeBaseStreamRequestKnowledgeParams {
	return s.KnowledgeParams
}

func (s *ChatWithKnowledgeBaseStreamRequest) GetModelParams() *ChatWithKnowledgeBaseStreamRequestModelParams {
	return s.ModelParams
}

func (s *ChatWithKnowledgeBaseStreamRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ChatWithKnowledgeBaseStreamRequest) GetPromptParams() *string {
	return s.PromptParams
}

func (s *ChatWithKnowledgeBaseStreamRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ChatWithKnowledgeBaseStreamRequest) SetDBInstanceId(v string) *ChatWithKnowledgeBaseStreamRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ChatWithKnowledgeBaseStreamRequest) SetIncludeKnowledgeBaseResults(v bool) *ChatWithKnowledgeBaseStreamRequest {
	s.IncludeKnowledgeBaseResults = &v
	return s
}

func (s *ChatWithKnowledgeBaseStreamRequest) SetKnowledgeParams(v *ChatWithKnowledgeBaseStreamRequestKnowledgeParams) *ChatWithKnowledgeBaseStreamRequest {
	s.KnowledgeParams = v
	return s
}

func (s *ChatWithKnowledgeBaseStreamRequest) SetModelParams(v *ChatWithKnowledgeBaseStreamRequestModelParams) *ChatWithKnowledgeBaseStreamRequest {
	s.ModelParams = v
	return s
}

func (s *ChatWithKnowledgeBaseStreamRequest) SetOwnerId(v int64) *ChatWithKnowledgeBaseStreamRequest {
	s.OwnerId = &v
	return s
}

func (s *ChatWithKnowledgeBaseStreamRequest) SetPromptParams(v string) *ChatWithKnowledgeBaseStreamRequest {
	s.PromptParams = &v
	return s
}

func (s *ChatWithKnowledgeBaseStreamRequest) SetRegionId(v string) *ChatWithKnowledgeBaseStreamRequest {
	s.RegionId = &v
	return s
}

func (s *ChatWithKnowledgeBaseStreamRequest) Validate() error {
	if s.KnowledgeParams != nil {
		if err := s.KnowledgeParams.Validate(); err != nil {
			return err
		}
	}
	if s.ModelParams != nil {
		if err := s.ModelParams.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ChatWithKnowledgeBaseStreamRequestKnowledgeParams struct {
	// Specifies the method for merging results from multiple knowledge bases. Default: `RRF`. Valid values:
	//
	// - `RRF`
	//
	// - `Weight`
	//
	// example:
	//
	// "RRF"
	MergeMethod *string `json:"MergeMethod,omitempty" xml:"MergeMethod,omitempty"`
	// The arguments for the result merging method.
	MergeMethodArgs *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsMergeMethodArgs `json:"MergeMethodArgs,omitempty" xml:"MergeMethodArgs,omitempty" type:"Struct"`
	// Specifies the factor for reranking vector search results. The value must be greater than 1 and less than or equal to 5.
	//
	// > - Reranking may be inefficient if document chunks are sparse.
	//
	// >
	//
	// > - The number of items to rerank, calculated as `ceil(TopK 	- RerankFactor)`, should not exceed 50.
	//
	// example:
	//
	// 5.0
	RerankFactor *float64 `json:"RerankFactor,omitempty" xml:"RerankFactor,omitempty"`
	// The rerank model to use.
	RerankModel *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsRerankModel `json:"RerankModel,omitempty" xml:"RerankModel,omitempty" type:"Struct"`
	// An array of knowledge bases to search.
	//
	// This parameter is required.
	SourceCollection []*ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollection `json:"SourceCollection,omitempty" xml:"SourceCollection,omitempty" type:"Repeated"`
	// The total number of top results to return after merging results from all collections.
	//
	// example:
	//
	// 10
	TopK *int64 `json:"TopK,omitempty" xml:"TopK,omitempty"`
}

func (s ChatWithKnowledgeBaseStreamRequestKnowledgeParams) String() string {
	return dara.Prettify(s)
}

func (s ChatWithKnowledgeBaseStreamRequestKnowledgeParams) GoString() string {
	return s.String()
}

func (s *ChatWithKnowledgeBaseStreamRequestKnowledgeParams) GetMergeMethod() *string {
	return s.MergeMethod
}

func (s *ChatWithKnowledgeBaseStreamRequestKnowledgeParams) GetMergeMethodArgs() *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsMergeMethodArgs {
	return s.MergeMethodArgs
}

func (s *ChatWithKnowledgeBaseStreamRequestKnowledgeParams) GetRerankFactor() *float64 {
	return s.RerankFactor
}

func (s *ChatWithKnowledgeBaseStreamRequestKnowledgeParams) GetRerankModel() *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsRerankModel {
	return s.RerankModel
}

func (s *ChatWithKnowledgeBaseStreamRequestKnowledgeParams) GetSourceCollection() []*ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollection {
	return s.SourceCollection
}

func (s *ChatWithKnowledgeBaseStreamRequestKnowledgeParams) GetTopK() *int64 {
	return s.TopK
}

func (s *ChatWithKnowledgeBaseStreamRequestKnowledgeParams) SetMergeMethod(v string) *ChatWithKnowledgeBaseStreamRequestKnowledgeParams {
	s.MergeMethod = &v
	return s
}

func (s *ChatWithKnowledgeBaseStreamRequestKnowledgeParams) SetMergeMethodArgs(v *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsMergeMethodArgs) *ChatWithKnowledgeBaseStreamRequestKnowledgeParams {
	s.MergeMethodArgs = v
	return s
}

func (s *ChatWithKnowledgeBaseStreamRequestKnowledgeParams) SetRerankFactor(v float64) *ChatWithKnowledgeBaseStreamRequestKnowledgeParams {
	s.RerankFactor = &v
	return s
}

func (s *ChatWithKnowledgeBaseStreamRequestKnowledgeParams) SetRerankModel(v *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsRerankModel) *ChatWithKnowledgeBaseStreamRequestKnowledgeParams {
	s.RerankModel = v
	return s
}

func (s *ChatWithKnowledgeBaseStreamRequestKnowledgeParams) SetSourceCollection(v []*ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollection) *ChatWithKnowledgeBaseStreamRequestKnowledgeParams {
	s.SourceCollection = v
	return s
}

func (s *ChatWithKnowledgeBaseStreamRequestKnowledgeParams) SetTopK(v int64) *ChatWithKnowledgeBaseStreamRequestKnowledgeParams {
	s.TopK = &v
	return s
}

func (s *ChatWithKnowledgeBaseStreamRequestKnowledgeParams) Validate() error {
	if s.MergeMethodArgs != nil {
		if err := s.MergeMethodArgs.Validate(); err != nil {
			return err
		}
	}
	if s.RerankModel != nil {
		if err := s.RerankModel.Validate(); err != nil {
			return err
		}
	}
	if s.SourceCollection != nil {
		for _, item := range s.SourceCollection {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ChatWithKnowledgeBaseStreamRequestKnowledgeParamsMergeMethodArgs struct {
	// Parameters for the `RRF` merge method.
	Rrf *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsMergeMethodArgsRrf `json:"Rrf,omitempty" xml:"Rrf,omitempty" type:"Struct"`
	// Parameters for the `Weight` merge method.
	Weight *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsMergeMethodArgsWeight `json:"Weight,omitempty" xml:"Weight,omitempty" type:"Struct"`
}

func (s ChatWithKnowledgeBaseStreamRequestKnowledgeParamsMergeMethodArgs) String() string {
	return dara.Prettify(s)
}

func (s ChatWithKnowledgeBaseStreamRequestKnowledgeParamsMergeMethodArgs) GoString() string {
	return s.String()
}

func (s *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsMergeMethodArgs) GetRrf() *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsMergeMethodArgsRrf {
	return s.Rrf
}

func (s *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsMergeMethodArgs) GetWeight() *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsMergeMethodArgsWeight {
	return s.Weight
}

func (s *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsMergeMethodArgs) SetRrf(v *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsMergeMethodArgsRrf) *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsMergeMethodArgs {
	s.Rrf = v
	return s
}

func (s *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsMergeMethodArgs) SetWeight(v *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsMergeMethodArgsWeight) *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsMergeMethodArgs {
	s.Weight = v
	return s
}

func (s *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsMergeMethodArgs) Validate() error {
	if s.Rrf != nil {
		if err := s.Rrf.Validate(); err != nil {
			return err
		}
	}
	if s.Weight != nil {
		if err := s.Weight.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ChatWithKnowledgeBaseStreamRequestKnowledgeParamsMergeMethodArgsRrf struct {
	// The constant `k` used in the reciprocal rank fusion (RRF) formula `1/(k + rank_i)`. The value must be an integer greater than 1.
	//
	// example:
	//
	// 60
	K *int64 `json:"K,omitempty" xml:"K,omitempty"`
}

func (s ChatWithKnowledgeBaseStreamRequestKnowledgeParamsMergeMethodArgsRrf) String() string {
	return dara.Prettify(s)
}

func (s ChatWithKnowledgeBaseStreamRequestKnowledgeParamsMergeMethodArgsRrf) GoString() string {
	return s.String()
}

func (s *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsMergeMethodArgsRrf) GetK() *int64 {
	return s.K
}

func (s *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsMergeMethodArgsRrf) SetK(v int64) *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsMergeMethodArgsRrf {
	s.K = &v
	return s
}

func (s *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsMergeMethodArgsRrf) Validate() error {
	return dara.Validate(s)
}

type ChatWithKnowledgeBaseStreamRequestKnowledgeParamsMergeMethodArgsWeight struct {
	// An array of weights for each `SourceCollection`.
	Weights []*float64 `json:"Weights,omitempty" xml:"Weights,omitempty" type:"Repeated"`
}

func (s ChatWithKnowledgeBaseStreamRequestKnowledgeParamsMergeMethodArgsWeight) String() string {
	return dara.Prettify(s)
}

func (s ChatWithKnowledgeBaseStreamRequestKnowledgeParamsMergeMethodArgsWeight) GoString() string {
	return s.String()
}

func (s *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsMergeMethodArgsWeight) GetWeights() []*float64 {
	return s.Weights
}

func (s *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsMergeMethodArgsWeight) SetWeights(v []*float64) *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsMergeMethodArgsWeight {
	s.Weights = v
	return s
}

func (s *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsMergeMethodArgsWeight) Validate() error {
	return dara.Validate(s)
}

type ChatWithKnowledgeBaseStreamRequestKnowledgeParamsRerankModel struct {
	// An instruction for the rerank model.
	//
	// example:
	//
	// Given a web search query, retrieve relevant passages that answer the query
	Instruct *string `json:"Instruct,omitempty" xml:"Instruct,omitempty"`
	// The name of the rerank model.
	//
	// example:
	//
	// qwen3-rerank
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ChatWithKnowledgeBaseStreamRequestKnowledgeParamsRerankModel) String() string {
	return dara.Prettify(s)
}

func (s ChatWithKnowledgeBaseStreamRequestKnowledgeParamsRerankModel) GoString() string {
	return s.String()
}

func (s *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsRerankModel) GetInstruct() *string {
	return s.Instruct
}

func (s *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsRerankModel) GetName() *string {
	return s.Name
}

func (s *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsRerankModel) SetInstruct(v string) *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsRerankModel {
	s.Instruct = &v
	return s
}

func (s *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsRerankModel) SetName(v string) *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsRerankModel {
	s.Name = &v
	return s
}

func (s *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsRerankModel) Validate() error {
	return dara.Validate(s)
}

type ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollection struct {
	// The name of the collection to search.
	//
	// This parameter is required.
	//
	// example:
	//
	// cloud_index_adb_50943_prod
	Collection *string `json:"Collection,omitempty" xml:"Collection,omitempty"`
	// The namespace that contains the collection.
	//
	// > You can call the [ListNamespaces](https://help.aliyun.com/document_detail/2401502.html) operation to view available namespaces.
	//
	// example:
	//
	// ddstar_vector
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The password for the specified namespace.
	//
	// > This value is specified in the `CreateNamespace` operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// namespacePassword
	NamespacePassword *string `json:"NamespacePassword,omitempty" xml:"NamespacePassword,omitempty"`
	// Parameters for the knowledge base query.
	QueryParams *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollectionQueryParams `json:"QueryParams,omitempty" xml:"QueryParams,omitempty" type:"Struct"`
}

func (s ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollection) String() string {
	return dara.Prettify(s)
}

func (s ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollection) GoString() string {
	return s.String()
}

func (s *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollection) GetCollection() *string {
	return s.Collection
}

func (s *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollection) GetNamespace() *string {
	return s.Namespace
}

func (s *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollection) GetNamespacePassword() *string {
	return s.NamespacePassword
}

func (s *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollection) GetQueryParams() *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollectionQueryParams {
	return s.QueryParams
}

func (s *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollection) SetCollection(v string) *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollection {
	s.Collection = &v
	return s
}

func (s *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollection) SetNamespace(v string) *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollection {
	s.Namespace = &v
	return s
}

func (s *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollection) SetNamespacePassword(v string) *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollection {
	s.NamespacePassword = &v
	return s
}

func (s *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollection) SetQueryParams(v *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollectionQueryParams) *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollection {
	s.QueryParams = v
	return s
}

func (s *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollection) Validate() error {
	if s.QueryParams != nil {
		if err := s.QueryParams.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollectionQueryParams struct {
	// A filter expression to apply to the search, similar to a SQL `WHERE` clause.
	//
	// example:
	//
	// method_id=\\"e41695f0-2851-40ac-b21d-dd337b60d71c\\"
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// Specifies whether to enable knowledge graph enhancement. Default value: `false`.
	//
	// example:
	//
	// true
	GraphEnhance *bool `json:"GraphEnhance,omitempty" xml:"GraphEnhance,omitempty"`
	// The parameters for knowledge graph search.
	GraphSearchArgs *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollectionQueryParamsGraphSearchArgs `json:"GraphSearchArgs,omitempty" xml:"GraphSearchArgs,omitempty" type:"Struct"`
	// Specifies the hybrid search algorithm. If omitted, the system performs a basic score comparison of vector search and full-text retrieval results.
	//
	// Valid values:
	//
	// - `RRF`: Reciprocal rank fusion. Configure the `k` parameter in `HybridSearchArgs`.
	//
	// - `Weight`: Weighted score fusion. Use the `alpha` parameter in `HybridSearchArgs` to control the balance between vector and full-text search scores.
	//
	// - `Cascaded`: First performs full-text retrieval, then runs a vector search on the results.
	//
	// example:
	//
	// Cascaded
	HybridSearch *string `json:"HybridSearch,omitempty" xml:"HybridSearch,omitempty"`
	// The arguments for the specified hybrid search algorithm. Supports `RRF` and `Weight`.
	//
	// - `RRF`: Specifies the constant `k` in the score calculation formula `1/(k+rank_i)`. `k` must be an integer greater than 1. Format:
	//
	// ```
	//
	// {
	//
	//    "RRF": {
	//
	//     "k": 60
	//
	//    }
	//
	// }
	//
	// ```
	//
	// - `Weight`: Calculates the final score using the formula `alpha 	- vector_score + (1 - alpha) 	- text_score`. The `alpha` parameter balances the scores, ranging from 0 (full-text only) to 1 (vector only). Format:
	//
	// ```
	//
	// {
	//
	//    "Weight": {
	//
	//     "alpha": 0.5
	//
	//    }
	//
	// }
	//
	// ```
	HybridSearchArgs map[string]interface{} `json:"HybridSearchArgs,omitempty" xml:"HybridSearchArgs,omitempty"`
	// The distance metric for vector search. Valid values:
	//
	// - `l2`: Euclidean distance.
	//
	// - `ip`: Inner product.
	//
	// - `cosine`: Cosine similarity.
	//
	// example:
	//
	// cosine
	Metrics *string `json:"Metrics,omitempty" xml:"Metrics,omitempty"`
	// The recall window. Specifies a window of context to include around retrieved chunks. The value must be a two-element array `[A, B]`, where -10 <= A <= 0 and 0 <= B <= 10.
	//
	// > - This parameter is useful when document chunks are small and a search might miss important surrounding context.
	//
	// >
	//
	// > - The window is applied after reranking.
	RecallWindow []*int64 `json:"RecallWindow,omitempty" xml:"RecallWindow,omitempty" type:"Repeated"`
	// The rerank factor. If specified, the system reranks the results from the vector search. The value must be greater than 1 and less than or equal to 5.
	//
	// > - Reranking may be inefficient if document chunks are sparse.
	//
	// >
	//
	// > - The number of items to rerank, calculated as `ceil(TopK 	- RerankFactor)`, should not exceed 50.
	//
	// example:
	//
	// 2.0
	RerankFactor *float64 `json:"RerankFactor,omitempty" xml:"RerankFactor,omitempty"`
	// The rerank model to use.
	RerankModel *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollectionQueryParamsRerankModel `json:"RerankModel,omitempty" xml:"RerankModel,omitempty" type:"Struct"`
	// The number of top results to return from this collection.
	//
	// example:
	//
	// 101
	TopK *int64 `json:"TopK,omitempty" xml:"TopK,omitempty"`
	// Specifies whether to use full-text retrieval for hybrid search. If `false` (the default), only vector search is performed.
	//
	// example:
	//
	// true
	UseFullTextRetrieval *bool `json:"UseFullTextRetrieval,omitempty" xml:"UseFullTextRetrieval,omitempty"`
}

func (s ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollectionQueryParams) String() string {
	return dara.Prettify(s)
}

func (s ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollectionQueryParams) GoString() string {
	return s.String()
}

func (s *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollectionQueryParams) GetFilter() *string {
	return s.Filter
}

func (s *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollectionQueryParams) GetGraphEnhance() *bool {
	return s.GraphEnhance
}

func (s *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollectionQueryParams) GetGraphSearchArgs() *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollectionQueryParamsGraphSearchArgs {
	return s.GraphSearchArgs
}

func (s *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollectionQueryParams) GetHybridSearch() *string {
	return s.HybridSearch
}

func (s *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollectionQueryParams) GetHybridSearchArgs() map[string]interface{} {
	return s.HybridSearchArgs
}

func (s *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollectionQueryParams) GetMetrics() *string {
	return s.Metrics
}

func (s *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollectionQueryParams) GetRecallWindow() []*int64 {
	return s.RecallWindow
}

func (s *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollectionQueryParams) GetRerankFactor() *float64 {
	return s.RerankFactor
}

func (s *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollectionQueryParams) GetRerankModel() *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollectionQueryParamsRerankModel {
	return s.RerankModel
}

func (s *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollectionQueryParams) GetTopK() *int64 {
	return s.TopK
}

func (s *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollectionQueryParams) GetUseFullTextRetrieval() *bool {
	return s.UseFullTextRetrieval
}

func (s *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollectionQueryParams) SetFilter(v string) *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollectionQueryParams {
	s.Filter = &v
	return s
}

func (s *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollectionQueryParams) SetGraphEnhance(v bool) *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollectionQueryParams {
	s.GraphEnhance = &v
	return s
}

func (s *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollectionQueryParams) SetGraphSearchArgs(v *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollectionQueryParamsGraphSearchArgs) *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollectionQueryParams {
	s.GraphSearchArgs = v
	return s
}

func (s *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollectionQueryParams) SetHybridSearch(v string) *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollectionQueryParams {
	s.HybridSearch = &v
	return s
}

func (s *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollectionQueryParams) SetHybridSearchArgs(v map[string]interface{}) *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollectionQueryParams {
	s.HybridSearchArgs = v
	return s
}

func (s *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollectionQueryParams) SetMetrics(v string) *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollectionQueryParams {
	s.Metrics = &v
	return s
}

func (s *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollectionQueryParams) SetRecallWindow(v []*int64) *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollectionQueryParams {
	s.RecallWindow = v
	return s
}

func (s *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollectionQueryParams) SetRerankFactor(v float64) *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollectionQueryParams {
	s.RerankFactor = &v
	return s
}

func (s *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollectionQueryParams) SetRerankModel(v *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollectionQueryParamsRerankModel) *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollectionQueryParams {
	s.RerankModel = v
	return s
}

func (s *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollectionQueryParams) SetTopK(v int64) *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollectionQueryParams {
	s.TopK = &v
	return s
}

func (s *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollectionQueryParams) SetUseFullTextRetrieval(v bool) *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollectionQueryParams {
	s.UseFullTextRetrieval = &v
	return s
}

func (s *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollectionQueryParams) Validate() error {
	if s.GraphSearchArgs != nil {
		if err := s.GraphSearchArgs.Validate(); err != nil {
			return err
		}
	}
	if s.RerankModel != nil {
		if err := s.RerankModel.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollectionQueryParamsGraphSearchArgs struct {
	// The number of top entities and relationship edges to return. Default value: `60`.
	//
	// example:
	//
	// 60
	GraphTopK *int64 `json:"GraphTopK,omitempty" xml:"GraphTopK,omitempty"`
}

func (s ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollectionQueryParamsGraphSearchArgs) String() string {
	return dara.Prettify(s)
}

func (s ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollectionQueryParamsGraphSearchArgs) GoString() string {
	return s.String()
}

func (s *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollectionQueryParamsGraphSearchArgs) GetGraphTopK() *int64 {
	return s.GraphTopK
}

func (s *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollectionQueryParamsGraphSearchArgs) SetGraphTopK(v int64) *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollectionQueryParamsGraphSearchArgs {
	s.GraphTopK = &v
	return s
}

func (s *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollectionQueryParamsGraphSearchArgs) Validate() error {
	return dara.Validate(s)
}

type ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollectionQueryParamsRerankModel struct {
	// An instruction for the rerank model.
	//
	// example:
	//
	// Given a web search query, retrieve relevant passages that answer the query
	Instruct *string `json:"Instruct,omitempty" xml:"Instruct,omitempty"`
	// The name of the rerank model.
	//
	// example:
	//
	// qwen3-rerank
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollectionQueryParamsRerankModel) String() string {
	return dara.Prettify(s)
}

func (s ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollectionQueryParamsRerankModel) GoString() string {
	return s.String()
}

func (s *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollectionQueryParamsRerankModel) GetInstruct() *string {
	return s.Instruct
}

func (s *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollectionQueryParamsRerankModel) GetName() *string {
	return s.Name
}

func (s *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollectionQueryParamsRerankModel) SetInstruct(v string) *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollectionQueryParamsRerankModel {
	s.Instruct = &v
	return s
}

func (s *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollectionQueryParamsRerankModel) SetName(v string) *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollectionQueryParamsRerankModel {
	s.Name = &v
	return s
}

func (s *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollectionQueryParamsRerankModel) Validate() error {
	return dara.Validate(s)
}

type ChatWithKnowledgeBaseStreamRequestModelParams struct {
	// The maximum number of tokens to generate.
	//
	// example:
	//
	// 8192
	MaxTokens *int64 `json:"MaxTokens,omitempty" xml:"MaxTokens,omitempty"`
	// A list of messages in the conversation.
	//
	// This parameter is required.
	Messages []*ChatWithKnowledgeBaseStreamRequestModelParamsMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Repeated"`
	// The name of the Large Language Model to use. For a list of available models, refer to the [Model Studio documentation](https://help.aliyun.com/zh/model-studio/compatibility-of-openai-with-dashscope?spm=openapi-amp.newDocPublishment.0.0.257c281fH8TtM8\\&scm=20140722.H_2833609._.OR_help-T_cn~zh-V_1#eadfc13038jd5).
	//
	// This parameter is required.
	//
	// example:
	//
	// qwen-plus
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// The number of candidate responses to generate.
	//
	// example:
	//
	// 1
	N *int64 `json:"N,omitempty" xml:"N,omitempty"`
	// The presence penalty. A value between -2.0 and 2.0.
	//
	// example:
	//
	// 1.0
	PresencePenalty *float64 `json:"PresencePenalty,omitempty" xml:"PresencePenalty,omitempty"`
	// The random seed for sampling.
	//
	// example:
	//
	// 42
	Seed *int64 `json:"Seed,omitempty" xml:"Seed,omitempty"`
	// A list of stop sequences.
	Stop []*string `json:"Stop,omitempty" xml:"Stop,omitempty" type:"Repeated"`
	// The sampling temperature. A value between 0 and 2.
	//
	// example:
	//
	// 0.6
	Temperature *float64 `json:"Temperature,omitempty" xml:"Temperature,omitempty"`
	// A list of tools the model can call.
	Tools []*ChatWithKnowledgeBaseStreamRequestModelParamsTools `json:"Tools,omitempty" xml:"Tools,omitempty" type:"Repeated"`
	// The nucleus sampling probability threshold. A value between 0 and 1.
	//
	// example:
	//
	// 0.9
	TopP *float64 `json:"TopP,omitempty" xml:"TopP,omitempty"`
}

func (s ChatWithKnowledgeBaseStreamRequestModelParams) String() string {
	return dara.Prettify(s)
}

func (s ChatWithKnowledgeBaseStreamRequestModelParams) GoString() string {
	return s.String()
}

func (s *ChatWithKnowledgeBaseStreamRequestModelParams) GetMaxTokens() *int64 {
	return s.MaxTokens
}

func (s *ChatWithKnowledgeBaseStreamRequestModelParams) GetMessages() []*ChatWithKnowledgeBaseStreamRequestModelParamsMessages {
	return s.Messages
}

func (s *ChatWithKnowledgeBaseStreamRequestModelParams) GetModel() *string {
	return s.Model
}

func (s *ChatWithKnowledgeBaseStreamRequestModelParams) GetN() *int64 {
	return s.N
}

func (s *ChatWithKnowledgeBaseStreamRequestModelParams) GetPresencePenalty() *float64 {
	return s.PresencePenalty
}

func (s *ChatWithKnowledgeBaseStreamRequestModelParams) GetSeed() *int64 {
	return s.Seed
}

func (s *ChatWithKnowledgeBaseStreamRequestModelParams) GetStop() []*string {
	return s.Stop
}

func (s *ChatWithKnowledgeBaseStreamRequestModelParams) GetTemperature() *float64 {
	return s.Temperature
}

func (s *ChatWithKnowledgeBaseStreamRequestModelParams) GetTools() []*ChatWithKnowledgeBaseStreamRequestModelParamsTools {
	return s.Tools
}

func (s *ChatWithKnowledgeBaseStreamRequestModelParams) GetTopP() *float64 {
	return s.TopP
}

func (s *ChatWithKnowledgeBaseStreamRequestModelParams) SetMaxTokens(v int64) *ChatWithKnowledgeBaseStreamRequestModelParams {
	s.MaxTokens = &v
	return s
}

func (s *ChatWithKnowledgeBaseStreamRequestModelParams) SetMessages(v []*ChatWithKnowledgeBaseStreamRequestModelParamsMessages) *ChatWithKnowledgeBaseStreamRequestModelParams {
	s.Messages = v
	return s
}

func (s *ChatWithKnowledgeBaseStreamRequestModelParams) SetModel(v string) *ChatWithKnowledgeBaseStreamRequestModelParams {
	s.Model = &v
	return s
}

func (s *ChatWithKnowledgeBaseStreamRequestModelParams) SetN(v int64) *ChatWithKnowledgeBaseStreamRequestModelParams {
	s.N = &v
	return s
}

func (s *ChatWithKnowledgeBaseStreamRequestModelParams) SetPresencePenalty(v float64) *ChatWithKnowledgeBaseStreamRequestModelParams {
	s.PresencePenalty = &v
	return s
}

func (s *ChatWithKnowledgeBaseStreamRequestModelParams) SetSeed(v int64) *ChatWithKnowledgeBaseStreamRequestModelParams {
	s.Seed = &v
	return s
}

func (s *ChatWithKnowledgeBaseStreamRequestModelParams) SetStop(v []*string) *ChatWithKnowledgeBaseStreamRequestModelParams {
	s.Stop = v
	return s
}

func (s *ChatWithKnowledgeBaseStreamRequestModelParams) SetTemperature(v float64) *ChatWithKnowledgeBaseStreamRequestModelParams {
	s.Temperature = &v
	return s
}

func (s *ChatWithKnowledgeBaseStreamRequestModelParams) SetTools(v []*ChatWithKnowledgeBaseStreamRequestModelParamsTools) *ChatWithKnowledgeBaseStreamRequestModelParams {
	s.Tools = v
	return s
}

func (s *ChatWithKnowledgeBaseStreamRequestModelParams) SetTopP(v float64) *ChatWithKnowledgeBaseStreamRequestModelParams {
	s.TopP = &v
	return s
}

func (s *ChatWithKnowledgeBaseStreamRequestModelParams) Validate() error {
	if s.Messages != nil {
		for _, item := range s.Messages {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tools != nil {
		for _, item := range s.Tools {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ChatWithKnowledgeBaseStreamRequestModelParamsMessages struct {
	// The content of the message.
	//
	// This parameter is required.
	//
	// example:
	//
	// You are a helpful assistant.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The role of the message author. Valid values:
	//
	// - `system`
	//
	// - `user`
	//
	// - `assistant`
	//
	// This parameter is required.
	//
	// example:
	//
	// user
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s ChatWithKnowledgeBaseStreamRequestModelParamsMessages) String() string {
	return dara.Prettify(s)
}

func (s ChatWithKnowledgeBaseStreamRequestModelParamsMessages) GoString() string {
	return s.String()
}

func (s *ChatWithKnowledgeBaseStreamRequestModelParamsMessages) GetContent() *string {
	return s.Content
}

func (s *ChatWithKnowledgeBaseStreamRequestModelParamsMessages) GetRole() *string {
	return s.Role
}

func (s *ChatWithKnowledgeBaseStreamRequestModelParamsMessages) SetContent(v string) *ChatWithKnowledgeBaseStreamRequestModelParamsMessages {
	s.Content = &v
	return s
}

func (s *ChatWithKnowledgeBaseStreamRequestModelParamsMessages) SetRole(v string) *ChatWithKnowledgeBaseStreamRequestModelParamsMessages {
	s.Role = &v
	return s
}

func (s *ChatWithKnowledgeBaseStreamRequestModelParamsMessages) Validate() error {
	return dara.Validate(s)
}

type ChatWithKnowledgeBaseStreamRequestModelParamsTools struct {
	// The function information.
	Function *ChatWithKnowledgeBaseStreamRequestModelParamsToolsFunction `json:"Function,omitempty" xml:"Function,omitempty" type:"Struct"`
}

func (s ChatWithKnowledgeBaseStreamRequestModelParamsTools) String() string {
	return dara.Prettify(s)
}

func (s ChatWithKnowledgeBaseStreamRequestModelParamsTools) GoString() string {
	return s.String()
}

func (s *ChatWithKnowledgeBaseStreamRequestModelParamsTools) GetFunction() *ChatWithKnowledgeBaseStreamRequestModelParamsToolsFunction {
	return s.Function
}

func (s *ChatWithKnowledgeBaseStreamRequestModelParamsTools) SetFunction(v *ChatWithKnowledgeBaseStreamRequestModelParamsToolsFunction) *ChatWithKnowledgeBaseStreamRequestModelParamsTools {
	s.Function = v
	return s
}

func (s *ChatWithKnowledgeBaseStreamRequestModelParamsTools) Validate() error {
	if s.Function != nil {
		if err := s.Function.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ChatWithKnowledgeBaseStreamRequestModelParamsToolsFunction struct {
	// A description of the function tool.
	//
	// example:
	//
	// Get weather.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the function tool.
	//
	// example:
	//
	// get_weather
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The parameters of the function, described as a JSON Schema object.
	//
	// example:
	//
	// {"type": "object", ...}
	Parameters interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
}

func (s ChatWithKnowledgeBaseStreamRequestModelParamsToolsFunction) String() string {
	return dara.Prettify(s)
}

func (s ChatWithKnowledgeBaseStreamRequestModelParamsToolsFunction) GoString() string {
	return s.String()
}

func (s *ChatWithKnowledgeBaseStreamRequestModelParamsToolsFunction) GetDescription() *string {
	return s.Description
}

func (s *ChatWithKnowledgeBaseStreamRequestModelParamsToolsFunction) GetName() *string {
	return s.Name
}

func (s *ChatWithKnowledgeBaseStreamRequestModelParamsToolsFunction) GetParameters() interface{} {
	return s.Parameters
}

func (s *ChatWithKnowledgeBaseStreamRequestModelParamsToolsFunction) SetDescription(v string) *ChatWithKnowledgeBaseStreamRequestModelParamsToolsFunction {
	s.Description = &v
	return s
}

func (s *ChatWithKnowledgeBaseStreamRequestModelParamsToolsFunction) SetName(v string) *ChatWithKnowledgeBaseStreamRequestModelParamsToolsFunction {
	s.Name = &v
	return s
}

func (s *ChatWithKnowledgeBaseStreamRequestModelParamsToolsFunction) SetParameters(v interface{}) *ChatWithKnowledgeBaseStreamRequestModelParamsToolsFunction {
	s.Parameters = v
	return s
}

func (s *ChatWithKnowledgeBaseStreamRequestModelParamsToolsFunction) Validate() error {
	return dara.Validate(s)
}
