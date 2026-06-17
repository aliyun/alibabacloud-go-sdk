// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatWithKnowledgeBaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ChatWithKnowledgeBaseRequest
	GetDBInstanceId() *string
	SetIncludeKnowledgeBaseResults(v bool) *ChatWithKnowledgeBaseRequest
	GetIncludeKnowledgeBaseResults() *bool
	SetKnowledgeParams(v *ChatWithKnowledgeBaseRequestKnowledgeParams) *ChatWithKnowledgeBaseRequest
	GetKnowledgeParams() *ChatWithKnowledgeBaseRequestKnowledgeParams
	SetModelParams(v *ChatWithKnowledgeBaseRequestModelParams) *ChatWithKnowledgeBaseRequest
	GetModelParams() *ChatWithKnowledgeBaseRequestModelParams
	SetOwnerId(v int64) *ChatWithKnowledgeBaseRequest
	GetOwnerId() *int64
	SetPromptParams(v string) *ChatWithKnowledgeBaseRequest
	GetPromptParams() *string
	SetRegionId(v string) *ChatWithKnowledgeBaseRequest
	GetRegionId() *string
}

type ChatWithKnowledgeBaseRequest struct {
	// The instance ID.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/196830.html) operation to view the details of all instances in a target region, including their instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// Whether to include the raw retrieval results from the knowledge base in the response. Default: `false`.
	//
	// example:
	//
	// false
	IncludeKnowledgeBaseResults *bool `json:"IncludeKnowledgeBaseResults,omitempty" xml:"IncludeKnowledgeBaseResults,omitempty"`
	// Parameters for knowledge retrieval. If omitted, the operation performs a standard chat without retrieving from a knowledge base.
	KnowledgeParams *ChatWithKnowledgeBaseRequestKnowledgeParams `json:"KnowledgeParams,omitempty" xml:"KnowledgeParams,omitempty" type:"Struct"`
	// The parameters for calling the large language model (LLM).
	//
	// This parameter is required.
	ModelParams *ChatWithKnowledgeBaseRequestModelParams `json:"ModelParams,omitempty" xml:"ModelParams,omitempty" type:"Struct"`
	OwnerId     *int64                                   `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// A custom system prompt template. If specified, it overrides the default prompt. The template must include the {{ text_chunks }}, {{ user_system_prompt }}, {{ graph_entities }}, and {{ graph_relations }} placeholders.
	//
	// example:
	//
	// "参考以下知识回答问题:{{ text_chunks }}"
	PromptParams *string `json:"PromptParams,omitempty" xml:"PromptParams,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ChatWithKnowledgeBaseRequest) String() string {
	return dara.Prettify(s)
}

func (s ChatWithKnowledgeBaseRequest) GoString() string {
	return s.String()
}

func (s *ChatWithKnowledgeBaseRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ChatWithKnowledgeBaseRequest) GetIncludeKnowledgeBaseResults() *bool {
	return s.IncludeKnowledgeBaseResults
}

func (s *ChatWithKnowledgeBaseRequest) GetKnowledgeParams() *ChatWithKnowledgeBaseRequestKnowledgeParams {
	return s.KnowledgeParams
}

func (s *ChatWithKnowledgeBaseRequest) GetModelParams() *ChatWithKnowledgeBaseRequestModelParams {
	return s.ModelParams
}

func (s *ChatWithKnowledgeBaseRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ChatWithKnowledgeBaseRequest) GetPromptParams() *string {
	return s.PromptParams
}

func (s *ChatWithKnowledgeBaseRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ChatWithKnowledgeBaseRequest) SetDBInstanceId(v string) *ChatWithKnowledgeBaseRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ChatWithKnowledgeBaseRequest) SetIncludeKnowledgeBaseResults(v bool) *ChatWithKnowledgeBaseRequest {
	s.IncludeKnowledgeBaseResults = &v
	return s
}

func (s *ChatWithKnowledgeBaseRequest) SetKnowledgeParams(v *ChatWithKnowledgeBaseRequestKnowledgeParams) *ChatWithKnowledgeBaseRequest {
	s.KnowledgeParams = v
	return s
}

func (s *ChatWithKnowledgeBaseRequest) SetModelParams(v *ChatWithKnowledgeBaseRequestModelParams) *ChatWithKnowledgeBaseRequest {
	s.ModelParams = v
	return s
}

func (s *ChatWithKnowledgeBaseRequest) SetOwnerId(v int64) *ChatWithKnowledgeBaseRequest {
	s.OwnerId = &v
	return s
}

func (s *ChatWithKnowledgeBaseRequest) SetPromptParams(v string) *ChatWithKnowledgeBaseRequest {
	s.PromptParams = &v
	return s
}

func (s *ChatWithKnowledgeBaseRequest) SetRegionId(v string) *ChatWithKnowledgeBaseRequest {
	s.RegionId = &v
	return s
}

func (s *ChatWithKnowledgeBaseRequest) Validate() error {
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

type ChatWithKnowledgeBaseRequestKnowledgeParams struct {
	// The method for merging results from multiple knowledge bases. Default: `RRF`. Valid values:
	//
	// - `RRF`
	//
	// - `Weight`
	//
	// example:
	//
	// "RRF"
	MergeMethod *string `json:"MergeMethod,omitempty" xml:"MergeMethod,omitempty"`
	// The parameters for the merge method.
	MergeMethodArgs *ChatWithKnowledgeBaseRequestKnowledgeParamsMergeMethodArgs `json:"MergeMethodArgs,omitempty" xml:"MergeMethodArgs,omitempty" type:"Struct"`
	// The reranking factor. Specify this to rerank the initial vector retrieval results for improved relevance. Valid range: (1, 5].
	//
	// > - Reranking may be less efficient if document chunks are sparse.
	//
	// >
	//
	// > - We recommend that the number of items to rerank, calculated as `Ceiling(TopK 	- RerankFactor)`, does not exceed 50.
	//
	// example:
	//
	// 1.0001
	RerankFactor *float64 `json:"RerankFactor,omitempty" xml:"RerankFactor,omitempty"`
	// The model to use for reranking.
	RerankModel *ChatWithKnowledgeBaseRequestKnowledgeParamsRerankModel `json:"RerankModel,omitempty" xml:"RerankModel,omitempty" type:"Struct"`
	// An array of knowledge base collections to query.
	//
	// This parameter is required.
	SourceCollection []*ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollection `json:"SourceCollection,omitempty" xml:"SourceCollection,omitempty" type:"Repeated"`
	// The number of top results to return after the results from multiple vector collection recalls are merged.
	//
	// example:
	//
	// 10
	TopK *int64 `json:"TopK,omitempty" xml:"TopK,omitempty"`
}

func (s ChatWithKnowledgeBaseRequestKnowledgeParams) String() string {
	return dara.Prettify(s)
}

func (s ChatWithKnowledgeBaseRequestKnowledgeParams) GoString() string {
	return s.String()
}

func (s *ChatWithKnowledgeBaseRequestKnowledgeParams) GetMergeMethod() *string {
	return s.MergeMethod
}

func (s *ChatWithKnowledgeBaseRequestKnowledgeParams) GetMergeMethodArgs() *ChatWithKnowledgeBaseRequestKnowledgeParamsMergeMethodArgs {
	return s.MergeMethodArgs
}

func (s *ChatWithKnowledgeBaseRequestKnowledgeParams) GetRerankFactor() *float64 {
	return s.RerankFactor
}

func (s *ChatWithKnowledgeBaseRequestKnowledgeParams) GetRerankModel() *ChatWithKnowledgeBaseRequestKnowledgeParamsRerankModel {
	return s.RerankModel
}

func (s *ChatWithKnowledgeBaseRequestKnowledgeParams) GetSourceCollection() []*ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollection {
	return s.SourceCollection
}

func (s *ChatWithKnowledgeBaseRequestKnowledgeParams) GetTopK() *int64 {
	return s.TopK
}

func (s *ChatWithKnowledgeBaseRequestKnowledgeParams) SetMergeMethod(v string) *ChatWithKnowledgeBaseRequestKnowledgeParams {
	s.MergeMethod = &v
	return s
}

func (s *ChatWithKnowledgeBaseRequestKnowledgeParams) SetMergeMethodArgs(v *ChatWithKnowledgeBaseRequestKnowledgeParamsMergeMethodArgs) *ChatWithKnowledgeBaseRequestKnowledgeParams {
	s.MergeMethodArgs = v
	return s
}

func (s *ChatWithKnowledgeBaseRequestKnowledgeParams) SetRerankFactor(v float64) *ChatWithKnowledgeBaseRequestKnowledgeParams {
	s.RerankFactor = &v
	return s
}

func (s *ChatWithKnowledgeBaseRequestKnowledgeParams) SetRerankModel(v *ChatWithKnowledgeBaseRequestKnowledgeParamsRerankModel) *ChatWithKnowledgeBaseRequestKnowledgeParams {
	s.RerankModel = v
	return s
}

func (s *ChatWithKnowledgeBaseRequestKnowledgeParams) SetSourceCollection(v []*ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollection) *ChatWithKnowledgeBaseRequestKnowledgeParams {
	s.SourceCollection = v
	return s
}

func (s *ChatWithKnowledgeBaseRequestKnowledgeParams) SetTopK(v int64) *ChatWithKnowledgeBaseRequestKnowledgeParams {
	s.TopK = &v
	return s
}

func (s *ChatWithKnowledgeBaseRequestKnowledgeParams) Validate() error {
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

type ChatWithKnowledgeBaseRequestKnowledgeParamsMergeMethodArgs struct {
	// Parameters to use when `MergeMethod` is set to `RRF`.
	Rrf *ChatWithKnowledgeBaseRequestKnowledgeParamsMergeMethodArgsRrf `json:"Rrf,omitempty" xml:"Rrf,omitempty" type:"Struct"`
	// Parameters to use when `MergeMethod` is set to `Weight`.
	Weight *ChatWithKnowledgeBaseRequestKnowledgeParamsMergeMethodArgsWeight `json:"Weight,omitempty" xml:"Weight,omitempty" type:"Struct"`
}

func (s ChatWithKnowledgeBaseRequestKnowledgeParamsMergeMethodArgs) String() string {
	return dara.Prettify(s)
}

func (s ChatWithKnowledgeBaseRequestKnowledgeParamsMergeMethodArgs) GoString() string {
	return s.String()
}

func (s *ChatWithKnowledgeBaseRequestKnowledgeParamsMergeMethodArgs) GetRrf() *ChatWithKnowledgeBaseRequestKnowledgeParamsMergeMethodArgsRrf {
	return s.Rrf
}

func (s *ChatWithKnowledgeBaseRequestKnowledgeParamsMergeMethodArgs) GetWeight() *ChatWithKnowledgeBaseRequestKnowledgeParamsMergeMethodArgsWeight {
	return s.Weight
}

func (s *ChatWithKnowledgeBaseRequestKnowledgeParamsMergeMethodArgs) SetRrf(v *ChatWithKnowledgeBaseRequestKnowledgeParamsMergeMethodArgsRrf) *ChatWithKnowledgeBaseRequestKnowledgeParamsMergeMethodArgs {
	s.Rrf = v
	return s
}

func (s *ChatWithKnowledgeBaseRequestKnowledgeParamsMergeMethodArgs) SetWeight(v *ChatWithKnowledgeBaseRequestKnowledgeParamsMergeMethodArgsWeight) *ChatWithKnowledgeBaseRequestKnowledgeParamsMergeMethodArgs {
	s.Weight = v
	return s
}

func (s *ChatWithKnowledgeBaseRequestKnowledgeParamsMergeMethodArgs) Validate() error {
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

type ChatWithKnowledgeBaseRequestKnowledgeParamsMergeMethodArgsRrf struct {
	// The constant `k` in the reciprocal rank fusion formula `1/(k + rank_i)`. The value must be a positive integer greater than 1.
	//
	// example:
	//
	// 60
	K *int64 `json:"K,omitempty" xml:"K,omitempty"`
}

func (s ChatWithKnowledgeBaseRequestKnowledgeParamsMergeMethodArgsRrf) String() string {
	return dara.Prettify(s)
}

func (s ChatWithKnowledgeBaseRequestKnowledgeParamsMergeMethodArgsRrf) GoString() string {
	return s.String()
}

func (s *ChatWithKnowledgeBaseRequestKnowledgeParamsMergeMethodArgsRrf) GetK() *int64 {
	return s.K
}

func (s *ChatWithKnowledgeBaseRequestKnowledgeParamsMergeMethodArgsRrf) SetK(v int64) *ChatWithKnowledgeBaseRequestKnowledgeParamsMergeMethodArgsRrf {
	s.K = &v
	return s
}

func (s *ChatWithKnowledgeBaseRequestKnowledgeParamsMergeMethodArgsRrf) Validate() error {
	return dara.Validate(s)
}

type ChatWithKnowledgeBaseRequestKnowledgeParamsMergeMethodArgsWeight struct {
	// An array of weights corresponding to each collection specified in `SourceCollection`.
	Weights []*float64 `json:"Weights,omitempty" xml:"Weights,omitempty" type:"Repeated"`
}

func (s ChatWithKnowledgeBaseRequestKnowledgeParamsMergeMethodArgsWeight) String() string {
	return dara.Prettify(s)
}

func (s ChatWithKnowledgeBaseRequestKnowledgeParamsMergeMethodArgsWeight) GoString() string {
	return s.String()
}

func (s *ChatWithKnowledgeBaseRequestKnowledgeParamsMergeMethodArgsWeight) GetWeights() []*float64 {
	return s.Weights
}

func (s *ChatWithKnowledgeBaseRequestKnowledgeParamsMergeMethodArgsWeight) SetWeights(v []*float64) *ChatWithKnowledgeBaseRequestKnowledgeParamsMergeMethodArgsWeight {
	s.Weights = v
	return s
}

func (s *ChatWithKnowledgeBaseRequestKnowledgeParamsMergeMethodArgsWeight) Validate() error {
	return dara.Validate(s)
}

type ChatWithKnowledgeBaseRequestKnowledgeParamsRerankModel struct {
	// The instruction or prompt for the reranking model.
	//
	// example:
	//
	// Given a web search query, retrieve relevant passages that answer the query
	Instruct *string `json:"Instruct,omitempty" xml:"Instruct,omitempty"`
	// The name of the reranking model.
	//
	// example:
	//
	// qwen3-rerank
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ChatWithKnowledgeBaseRequestKnowledgeParamsRerankModel) String() string {
	return dara.Prettify(s)
}

func (s ChatWithKnowledgeBaseRequestKnowledgeParamsRerankModel) GoString() string {
	return s.String()
}

func (s *ChatWithKnowledgeBaseRequestKnowledgeParamsRerankModel) GetInstruct() *string {
	return s.Instruct
}

func (s *ChatWithKnowledgeBaseRequestKnowledgeParamsRerankModel) GetName() *string {
	return s.Name
}

func (s *ChatWithKnowledgeBaseRequestKnowledgeParamsRerankModel) SetInstruct(v string) *ChatWithKnowledgeBaseRequestKnowledgeParamsRerankModel {
	s.Instruct = &v
	return s
}

func (s *ChatWithKnowledgeBaseRequestKnowledgeParamsRerankModel) SetName(v string) *ChatWithKnowledgeBaseRequestKnowledgeParamsRerankModel {
	s.Name = &v
	return s
}

func (s *ChatWithKnowledgeBaseRequestKnowledgeParamsRerankModel) Validate() error {
	return dara.Validate(s)
}

type ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollection struct {
	// The name of the collection to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// adbpg_document_collection
	Collection *string `json:"Collection,omitempty" xml:"Collection,omitempty"`
	// The namespace where the collection resides. Default: `public`.
	//
	// > You can create a namespace by calling the [CreateNamespace](https://help.aliyun.com/document_detail/2401495.html) operation and view existing namespaces by calling the [ListNamespaces](https://help.aliyun.com/document_detail/2401502.html) operation.
	//
	// example:
	//
	// dukang
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The password for the specified namespace.
	//
	// > This password is set when you call the `CreateNamespace` operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// namespacePasswd
	NamespacePassword *string `json:"NamespacePassword,omitempty" xml:"NamespacePassword,omitempty"`
	// Retrieval parameters for this knowledge base collection.
	QueryParams *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParams `json:"QueryParams,omitempty" xml:"QueryParams,omitempty" type:"Struct"`
}

func (s ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollection) String() string {
	return dara.Prettify(s)
}

func (s ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollection) GoString() string {
	return s.String()
}

func (s *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollection) GetCollection() *string {
	return s.Collection
}

func (s *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollection) GetNamespace() *string {
	return s.Namespace
}

func (s *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollection) GetNamespacePassword() *string {
	return s.NamespacePassword
}

func (s *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollection) GetQueryParams() *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParams {
	return s.QueryParams
}

func (s *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollection) SetCollection(v string) *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollection {
	s.Collection = &v
	return s
}

func (s *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollection) SetNamespace(v string) *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollection {
	s.Namespace = &v
	return s
}

func (s *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollection) SetNamespacePassword(v string) *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollection {
	s.NamespacePassword = &v
	return s
}

func (s *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollection) SetQueryParams(v *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParams) *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollection {
	s.QueryParams = v
	return s
}

func (s *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollection) Validate() error {
	if s.QueryParams != nil {
		if err := s.QueryParams.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParams struct {
	// A filter to apply to the search, specified as a SQL `WHERE` clause.
	//
	// example:
	//
	// id = \\"llm-t87l87fxuhn56woc_8anu8j2d3f_file_e74635e2cc314e838543e724f6b3b1f2_10658020\\"
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// Whether to enhance the search with a knowledge graph. Default: `false`.
	//
	// example:
	//
	// false
	GraphEnhance *bool `json:"GraphEnhance,omitempty" xml:"GraphEnhance,omitempty"`
	// Parameters for the knowledge graph search, used when `GraphEnhance` is `true`.
	GraphSearchArgs *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParamsGraphSearchArgs `json:"GraphSearchArgs,omitempty" xml:"GraphSearchArgs,omitempty" type:"Struct"`
	// The multi-channel recall algorithm. If omitted, the system directly compares and sorts scores from dense vector retrieval and full-text search.
	//
	// Valid values:
	//
	// - `RRF`: Uses reciprocal rank fusion. The fusion effect is controlled by the `k` parameter in `HybridSearchArgs`.
	//
	// - `Weight`: Uses weighted sorting. The weights for vector retrieval and full-text search scores are controlled by parameters in `HybridSearchArgs`.
	//
	// - `Cascaded`: Performs a full-text search first, followed by a vector retrieval on the results.
	//
	// example:
	//
	// RRF
	HybridSearch *string `json:"HybridSearch,omitempty" xml:"HybridSearch,omitempty"`
	// Parameters for the multi-channel recall algorithm. `RRF` and `Weight` are supported. The `HybridPathsSetting` parameter can specify the recall channels: `dense` (dense vector), `sparse` (sparse vector), and `fulltext` (full-text search). If this parameter is empty, `dense` and `fulltext` are used by default.
	//
	// - `RRF`: Specifies the constant `k` in the formula `1/(k+rank_i)`. The value must be a positive integer greater than 1. Format:
	//
	// ```
	//
	// {
	//
	//   "HybridPathsSetting": {
	//
	//     "paths": "dense,fulltext"
	//
	//   },
	//
	//   "RRF": {
	//
	//     "k": 60
	//
	//   }
	//
	// }
	//
	// ```
	//
	// - `Weight`:
	//
	//   - Two-channel recall (if `HybridPathsSetting` is not specified, only `alpha` is required):
	//
	//     - Formula: `alpha 	- dense_score + (1-alpha) 	- fulltext_score`. The `alpha` parameter represents the weight of the dense vector score relative to the full-text search score. The value must be in the range [0, 1]. A value of 0 uses only full-text search. A value of 1 uses only dense vector retrieval.
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
	//
	// - Three-channel recall:
	//
	//   - Formula: `normalized_dense 	- dense_score + normalized_sparse 	- sparse_score + normalized_fulltext 	- fulltext_score`. The `dense`, `sparse`, and `fulltext` parameters represent the weights for each channel and must be greater than or equal to 0. The system automatically normalizes these weights (for example, `normalized_x = x / (dense + sparse + fulltext)`).
	//
	// ```
	//
	// {
	//
	//   "HybridPathsSetting": {
	//
	//      "paths": "dense,sparse,fulltext"
	//
	//    },
	//
	//   "Weight": {
	//
	//     "dense": 0.5,
	//
	//     "sparse": 0.3,
	//
	//     "fulltext": 0.2
	//
	//   }
	//
	// }
	//
	// ```
	HybridSearchArgs map[string]interface{} `json:"HybridSearchArgs,omitempty" xml:"HybridSearchArgs,omitempty"`
	// The distance metric used for vector indexing. Valid values:
	//
	// - `l2`: euclidean distance.
	//
	// - `ip`: Inner product (dot product) distance.
	//
	// - `cosine`: cosine similarity.
	//
	// example:
	//
	// cosine
	Metrics *string `json:"Metrics,omitempty" xml:"Metrics,omitempty"`
	// The recall window. If specified, expands the context around retrieved text chunks. Must be an array of two integers, `[A, B]`, where `A` is the number of preceding chunks to include (from -10 to 0) and `B` is the number of following chunks (from 0 to 10).
	//
	// > - We recommend that you use this parameter when document chunks are highly fragmented and retrieval might result in a loss of context.
	//
	// >
	//
	// > - Reranking is performed before windowing is applied.
	RecallWindow []*int64 `json:"RecallWindow,omitempty" xml:"RecallWindow,omitempty" type:"Repeated"`
	// The reranking factor for this collection, which overrides the top-level `RerankFactor`. If specified, it reranks the initial retrieval results to improve relevance. Valid range: (1, 5].
	//
	// > - Reranking may be less efficient if document chunks are sparse.
	//
	// >
	//
	// > - We recommend that the number of items to rerank, calculated as `Ceiling(TopK 	- RerankFactor)`, does not exceed 50.
	//
	// example:
	//
	// 1.5
	RerankFactor *float64 `json:"RerankFactor,omitempty" xml:"RerankFactor,omitempty"`
	// The model to use for reranking.
	RerankModel *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParamsRerankModel `json:"RerankModel,omitempty" xml:"RerankModel,omitempty" type:"Struct"`
	// The number of top results to return from this collection before merging.
	//
	// example:
	//
	// 10
	TopK *int64 `json:"TopK,omitempty" xml:"TopK,omitempty"`
	// Whether to enable full-text search in addition to vector retrieval. Default: `false` (uses only vector retrieval).
	//
	// example:
	//
	// true
	UseFullTextRetrieval *bool `json:"UseFullTextRetrieval,omitempty" xml:"UseFullTextRetrieval,omitempty"`
}

func (s ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParams) String() string {
	return dara.Prettify(s)
}

func (s ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParams) GoString() string {
	return s.String()
}

func (s *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParams) GetFilter() *string {
	return s.Filter
}

func (s *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParams) GetGraphEnhance() *bool {
	return s.GraphEnhance
}

func (s *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParams) GetGraphSearchArgs() *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParamsGraphSearchArgs {
	return s.GraphSearchArgs
}

func (s *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParams) GetHybridSearch() *string {
	return s.HybridSearch
}

func (s *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParams) GetHybridSearchArgs() map[string]interface{} {
	return s.HybridSearchArgs
}

func (s *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParams) GetMetrics() *string {
	return s.Metrics
}

func (s *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParams) GetRecallWindow() []*int64 {
	return s.RecallWindow
}

func (s *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParams) GetRerankFactor() *float64 {
	return s.RerankFactor
}

func (s *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParams) GetRerankModel() *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParamsRerankModel {
	return s.RerankModel
}

func (s *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParams) GetTopK() *int64 {
	return s.TopK
}

func (s *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParams) GetUseFullTextRetrieval() *bool {
	return s.UseFullTextRetrieval
}

func (s *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParams) SetFilter(v string) *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParams {
	s.Filter = &v
	return s
}

func (s *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParams) SetGraphEnhance(v bool) *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParams {
	s.GraphEnhance = &v
	return s
}

func (s *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParams) SetGraphSearchArgs(v *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParamsGraphSearchArgs) *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParams {
	s.GraphSearchArgs = v
	return s
}

func (s *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParams) SetHybridSearch(v string) *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParams {
	s.HybridSearch = &v
	return s
}

func (s *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParams) SetHybridSearchArgs(v map[string]interface{}) *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParams {
	s.HybridSearchArgs = v
	return s
}

func (s *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParams) SetMetrics(v string) *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParams {
	s.Metrics = &v
	return s
}

func (s *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParams) SetRecallWindow(v []*int64) *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParams {
	s.RecallWindow = v
	return s
}

func (s *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParams) SetRerankFactor(v float64) *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParams {
	s.RerankFactor = &v
	return s
}

func (s *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParams) SetRerankModel(v *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParamsRerankModel) *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParams {
	s.RerankModel = v
	return s
}

func (s *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParams) SetTopK(v int64) *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParams {
	s.TopK = &v
	return s
}

func (s *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParams) SetUseFullTextRetrieval(v bool) *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParams {
	s.UseFullTextRetrieval = &v
	return s
}

func (s *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParams) Validate() error {
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

type ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParamsGraphSearchArgs struct {
	// The maximum number of entities and relationship edges to return from the knowledge graph search. Default: 60.
	//
	// example:
	//
	// 60
	GraphTopK *int64 `json:"GraphTopK,omitempty" xml:"GraphTopK,omitempty"`
}

func (s ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParamsGraphSearchArgs) String() string {
	return dara.Prettify(s)
}

func (s ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParamsGraphSearchArgs) GoString() string {
	return s.String()
}

func (s *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParamsGraphSearchArgs) GetGraphTopK() *int64 {
	return s.GraphTopK
}

func (s *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParamsGraphSearchArgs) SetGraphTopK(v int64) *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParamsGraphSearchArgs {
	s.GraphTopK = &v
	return s
}

func (s *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParamsGraphSearchArgs) Validate() error {
	return dara.Validate(s)
}

type ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParamsRerankModel struct {
	// The instruction or prompt for the reranking model.
	//
	// example:
	//
	// Given a web search query, retrieve relevant passages that answer the query
	Instruct *string `json:"Instruct,omitempty" xml:"Instruct,omitempty"`
	// The name of the reranking model.
	//
	// example:
	//
	// qwen3-rerank
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParamsRerankModel) String() string {
	return dara.Prettify(s)
}

func (s ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParamsRerankModel) GoString() string {
	return s.String()
}

func (s *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParamsRerankModel) GetInstruct() *string {
	return s.Instruct
}

func (s *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParamsRerankModel) GetName() *string {
	return s.Name
}

func (s *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParamsRerankModel) SetInstruct(v string) *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParamsRerankModel {
	s.Instruct = &v
	return s
}

func (s *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParamsRerankModel) SetName(v string) *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParamsRerankModel {
	s.Name = &v
	return s
}

func (s *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParamsRerankModel) Validate() error {
	return dara.Validate(s)
}

type ChatWithKnowledgeBaseRequestModelParams struct {
	// The maximum number of tokens to generate.
	//
	// example:
	//
	// 8192
	MaxTokens *int64 `json:"MaxTokens,omitempty" xml:"MaxTokens,omitempty"`
	// The list of messages that form the conversation history.
	//
	// This parameter is required.
	Messages []*ChatWithKnowledgeBaseRequestModelParamsMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Repeated"`
	// The name of the large language model to use. For a list of available models, see the [Model Studio documentation](https://help.aliyun.com/zh/model-studio/compatibility-of-openai-with-dashscope?spm=a2c4g.11186623.help-menu-2400256.d_2_10_0.45b5516eZIciC8\\&scm=20140722.H_2833609._.OR_help-T_cn~zh-V_1#eadfc13038jd5).
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
	// The presence penalty. Valid range: [-2.0, 2.0].
	//
	// example:
	//
	// 1.0
	PresencePenalty *float64 `json:"PresencePenalty,omitempty" xml:"PresencePenalty,omitempty"`
	// The random seed.
	//
	// example:
	//
	// 42
	Seed *int64 `json:"Seed,omitempty" xml:"Seed,omitempty"`
	// A list of stop words.
	Stop []*string `json:"Stop,omitempty" xml:"Stop,omitempty" type:"Repeated"`
	// The sampling temperature. Valid range: (0, 2.0].
	//
	// example:
	//
	// 0.6
	Temperature *float64 `json:"Temperature,omitempty" xml:"Temperature,omitempty"`
	// The list of tools.
	Tools []*ChatWithKnowledgeBaseRequestModelParamsTools `json:"Tools,omitempty" xml:"Tools,omitempty" type:"Repeated"`
	// The probability threshold for nucleus sampling. Valid range: (0, 1.0).
	//
	// example:
	//
	// 0.9
	TopP *float64 `json:"TopP,omitempty" xml:"TopP,omitempty"`
}

func (s ChatWithKnowledgeBaseRequestModelParams) String() string {
	return dara.Prettify(s)
}

func (s ChatWithKnowledgeBaseRequestModelParams) GoString() string {
	return s.String()
}

func (s *ChatWithKnowledgeBaseRequestModelParams) GetMaxTokens() *int64 {
	return s.MaxTokens
}

func (s *ChatWithKnowledgeBaseRequestModelParams) GetMessages() []*ChatWithKnowledgeBaseRequestModelParamsMessages {
	return s.Messages
}

func (s *ChatWithKnowledgeBaseRequestModelParams) GetModel() *string {
	return s.Model
}

func (s *ChatWithKnowledgeBaseRequestModelParams) GetN() *int64 {
	return s.N
}

func (s *ChatWithKnowledgeBaseRequestModelParams) GetPresencePenalty() *float64 {
	return s.PresencePenalty
}

func (s *ChatWithKnowledgeBaseRequestModelParams) GetSeed() *int64 {
	return s.Seed
}

func (s *ChatWithKnowledgeBaseRequestModelParams) GetStop() []*string {
	return s.Stop
}

func (s *ChatWithKnowledgeBaseRequestModelParams) GetTemperature() *float64 {
	return s.Temperature
}

func (s *ChatWithKnowledgeBaseRequestModelParams) GetTools() []*ChatWithKnowledgeBaseRequestModelParamsTools {
	return s.Tools
}

func (s *ChatWithKnowledgeBaseRequestModelParams) GetTopP() *float64 {
	return s.TopP
}

func (s *ChatWithKnowledgeBaseRequestModelParams) SetMaxTokens(v int64) *ChatWithKnowledgeBaseRequestModelParams {
	s.MaxTokens = &v
	return s
}

func (s *ChatWithKnowledgeBaseRequestModelParams) SetMessages(v []*ChatWithKnowledgeBaseRequestModelParamsMessages) *ChatWithKnowledgeBaseRequestModelParams {
	s.Messages = v
	return s
}

func (s *ChatWithKnowledgeBaseRequestModelParams) SetModel(v string) *ChatWithKnowledgeBaseRequestModelParams {
	s.Model = &v
	return s
}

func (s *ChatWithKnowledgeBaseRequestModelParams) SetN(v int64) *ChatWithKnowledgeBaseRequestModelParams {
	s.N = &v
	return s
}

func (s *ChatWithKnowledgeBaseRequestModelParams) SetPresencePenalty(v float64) *ChatWithKnowledgeBaseRequestModelParams {
	s.PresencePenalty = &v
	return s
}

func (s *ChatWithKnowledgeBaseRequestModelParams) SetSeed(v int64) *ChatWithKnowledgeBaseRequestModelParams {
	s.Seed = &v
	return s
}

func (s *ChatWithKnowledgeBaseRequestModelParams) SetStop(v []*string) *ChatWithKnowledgeBaseRequestModelParams {
	s.Stop = v
	return s
}

func (s *ChatWithKnowledgeBaseRequestModelParams) SetTemperature(v float64) *ChatWithKnowledgeBaseRequestModelParams {
	s.Temperature = &v
	return s
}

func (s *ChatWithKnowledgeBaseRequestModelParams) SetTools(v []*ChatWithKnowledgeBaseRequestModelParamsTools) *ChatWithKnowledgeBaseRequestModelParams {
	s.Tools = v
	return s
}

func (s *ChatWithKnowledgeBaseRequestModelParams) SetTopP(v float64) *ChatWithKnowledgeBaseRequestModelParams {
	s.TopP = &v
	return s
}

func (s *ChatWithKnowledgeBaseRequestModelParams) Validate() error {
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

type ChatWithKnowledgeBaseRequestModelParamsMessages struct {
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

func (s ChatWithKnowledgeBaseRequestModelParamsMessages) String() string {
	return dara.Prettify(s)
}

func (s ChatWithKnowledgeBaseRequestModelParamsMessages) GoString() string {
	return s.String()
}

func (s *ChatWithKnowledgeBaseRequestModelParamsMessages) GetContent() *string {
	return s.Content
}

func (s *ChatWithKnowledgeBaseRequestModelParamsMessages) GetRole() *string {
	return s.Role
}

func (s *ChatWithKnowledgeBaseRequestModelParamsMessages) SetContent(v string) *ChatWithKnowledgeBaseRequestModelParamsMessages {
	s.Content = &v
	return s
}

func (s *ChatWithKnowledgeBaseRequestModelParamsMessages) SetRole(v string) *ChatWithKnowledgeBaseRequestModelParamsMessages {
	s.Role = &v
	return s
}

func (s *ChatWithKnowledgeBaseRequestModelParamsMessages) Validate() error {
	return dara.Validate(s)
}

type ChatWithKnowledgeBaseRequestModelParamsTools struct {
	// The function information.
	Function *ChatWithKnowledgeBaseRequestModelParamsToolsFunction `json:"Function,omitempty" xml:"Function,omitempty" type:"Struct"`
}

func (s ChatWithKnowledgeBaseRequestModelParamsTools) String() string {
	return dara.Prettify(s)
}

func (s ChatWithKnowledgeBaseRequestModelParamsTools) GoString() string {
	return s.String()
}

func (s *ChatWithKnowledgeBaseRequestModelParamsTools) GetFunction() *ChatWithKnowledgeBaseRequestModelParamsToolsFunction {
	return s.Function
}

func (s *ChatWithKnowledgeBaseRequestModelParamsTools) SetFunction(v *ChatWithKnowledgeBaseRequestModelParamsToolsFunction) *ChatWithKnowledgeBaseRequestModelParamsTools {
	s.Function = v
	return s
}

func (s *ChatWithKnowledgeBaseRequestModelParamsTools) Validate() error {
	if s.Function != nil {
		if err := s.Function.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ChatWithKnowledgeBaseRequestModelParamsToolsFunction struct {
	// The description of the function tool.
	//
	// example:
	//
	// 获取天气。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the function tool.
	//
	// example:
	//
	// get_weather
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The JSON schema of the function parameters.
	//
	// example:
	//
	// {"type": "object", ...}
	Parameters interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
}

func (s ChatWithKnowledgeBaseRequestModelParamsToolsFunction) String() string {
	return dara.Prettify(s)
}

func (s ChatWithKnowledgeBaseRequestModelParamsToolsFunction) GoString() string {
	return s.String()
}

func (s *ChatWithKnowledgeBaseRequestModelParamsToolsFunction) GetDescription() *string {
	return s.Description
}

func (s *ChatWithKnowledgeBaseRequestModelParamsToolsFunction) GetName() *string {
	return s.Name
}

func (s *ChatWithKnowledgeBaseRequestModelParamsToolsFunction) GetParameters() interface{} {
	return s.Parameters
}

func (s *ChatWithKnowledgeBaseRequestModelParamsToolsFunction) SetDescription(v string) *ChatWithKnowledgeBaseRequestModelParamsToolsFunction {
	s.Description = &v
	return s
}

func (s *ChatWithKnowledgeBaseRequestModelParamsToolsFunction) SetName(v string) *ChatWithKnowledgeBaseRequestModelParamsToolsFunction {
	s.Name = &v
	return s
}

func (s *ChatWithKnowledgeBaseRequestModelParamsToolsFunction) SetParameters(v interface{}) *ChatWithKnowledgeBaseRequestModelParamsToolsFunction {
	s.Parameters = v
	return s
}

func (s *ChatWithKnowledgeBaseRequestModelParamsToolsFunction) Validate() error {
	return dara.Validate(s)
}
