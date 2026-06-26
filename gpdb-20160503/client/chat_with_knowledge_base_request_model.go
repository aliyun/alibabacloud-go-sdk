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
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/196830.html) operation to query the details of all instances in the target region, including the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// Specifies whether to return the recall results. Default value: false.
	//
	// example:
	//
	// false
	IncludeKnowledgeBaseResults *bool `json:"IncludeKnowledgeBaseResults,omitempty" xml:"IncludeKnowledgeBaseResults,omitempty"`
	// The knowledge retrieval parameter object. If not specified, only chat is performed.
	KnowledgeParams *ChatWithKnowledgeBaseRequestKnowledgeParams `json:"KnowledgeParams,omitempty" xml:"KnowledgeParams,omitempty" type:"Struct"`
	// The large language model (LLM) invocation parameter object.
	//
	// This parameter is required.
	ModelParams *ChatWithKnowledgeBaseRequestModelParams `json:"ModelParams,omitempty" xml:"ModelParams,omitempty" type:"Struct"`
	OwnerId     *int64                                   `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The system prompt template, which must include {{ text_chunks }}, {{ user_system_prompt }}, {{ graph_entities }}, and {{ graph_relations }}. If not specified, this part does not take effect.
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
	// The method for merging multiple knowledge bases. Default is RRF. Valid values:
	//
	// - RRF
	//
	// - Weight
	//
	// example:
	//
	// "RRF"
	MergeMethod *string `json:"MergeMethod,omitempty" xml:"MergeMethod,omitempty"`
	// The parameters for multi-knowledge base fusion.
	MergeMethodArgs *ChatWithKnowledgeBaseRequestKnowledgeParamsMergeMethodArgs `json:"MergeMethodArgs,omitempty" xml:"MergeMethodArgs,omitempty" type:"Struct"`
	// The reranking factor. When this value is not empty, the vector retrieval results are reranked. Value range: 1 < RerankFactor <= 5.
	//
	// > - Reranking is slow when document segmentation is sparse.
	//
	// > - It is recommended that the number of items to rerank (TopK 	- Factor, rounded up) does not exceed 50.
	//
	// example:
	//
	// 1.0001
	RerankFactor *float64                                                `json:"RerankFactor,omitempty" xml:"RerankFactor,omitempty"`
	RerankModel  *ChatWithKnowledgeBaseRequestKnowledgeParamsRerankModel `json:"RerankModel,omitempty" xml:"RerankModel,omitempty" type:"Struct"`
	// The list of knowledge bases.
	//
	// This parameter is required.
	SourceCollection []*ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollection `json:"SourceCollection,omitempty" xml:"SourceCollection,omitempty" type:"Repeated"`
	// The number of top results to return after merging the recall results from multiple vector collections.
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
	// The configurable parameters when MergeMethod is set to RRF.
	Rrf *ChatWithKnowledgeBaseRequestKnowledgeParamsMergeMethodArgsRrf `json:"Rrf,omitempty" xml:"Rrf,omitempty" type:"Struct"`
	// The configurable parameters when MergeMethod is set to Weight.
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
	// The k constant in the scoring algorithm 1/(k+rank_i). The value must be a positive integer greater than 1.
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
	// The weight array for each SourceCollection.
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
	// example:
	//
	// Given a web search query, retrieve relevant passages that answer the query
	Instruct *string `json:"Instruct,omitempty" xml:"Instruct,omitempty"`
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
	// The name of the collection to recall.
	//
	// This parameter is required.
	//
	// example:
	//
	// adbpg_document_collection
	Collection *string `json:"Collection,omitempty" xml:"Collection,omitempty"`
	// The namespace. Default value: public.
	//
	// > You can create a namespace by calling the [CreateNamespace](https://help.aliyun.com/document_detail/2401495.html) operation, and view the list by calling the [ListNamespaces](https://help.aliyun.com/document_detail/2401502.html) operation.
	//
	// example:
	//
	// dukang
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The password corresponding to the namespace.
	//
	// > This value is specified in the CreateNamespace operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// namespacePasswd
	NamespacePassword *string `json:"NamespacePassword,omitempty" xml:"NamespacePassword,omitempty"`
	// The parameters related to knowledge base retrieval.
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
	// The filter condition for the data to be updated, in SQL WHERE clause format.
	//
	// example:
	//
	// id = \\"llm-t87l87fxuhn56woc_8anu8j2d3f_file_e74635e2cc314e838543e724f6b3b1f2_10658020\\"
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// Specifies whether to enable knowledge graph enhancement. Default value: false.
	//
	// example:
	//
	// false
	GraphEnhance *bool `json:"GraphEnhance,omitempty" xml:"GraphEnhance,omitempty"`
	// The number of top entities and relationship edges to return. Default value: 60.
	GraphSearchArgs *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParamsGraphSearchArgs `json:"GraphSearchArgs,omitempty" xml:"GraphSearchArgs,omitempty" type:"Struct"`
	// The multi-path recall algorithm. Default is empty (i.e., directly compares and sorts the dense vector and full-text scores).
	//
	// Valid values:
	//
	// - RRF: Reciprocal Rank Fusion. Has a parameter k to control the fusion effect. See HybridSearchArgs configuration for details.
	//
	// - Weight: Weight-based sorting. Uses parameters to control the score weights of vector and full-text retrieval, then sorts. See HybridSearchArgs configuration for details.
	//
	// - Cascaded: First performs full-text retrieval, then performs vector retrieval on top of it.
	//
	// example:
	//
	// RRF
	HybridSearch *string `json:"HybridSearch,omitempty" xml:"HybridSearch,omitempty"`
	// The algorithm parameters for multi-path recall. Currently supports RRF and Weight. HybridPathsSetting can specify recall of dense vectors (dense), sparse vectors (sparse), and full-text retrieval (fulltext). If the value is empty, dense vectors (dense) and full-text retrieval (fulltext) are recalled by default.
	//
	// - RRF: Specifies the k constant in the scoring algorithm `1/(k+rank_i)`. The value must be a positive integer greater than 1. Format:
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
	// - Weight:
	//
	//    - Dual-path recall (without specifying HybridPathsSetting, only specifying alpha):
	//
	//       - Formula: alpha 	- dense_score + (1-alpha) 	- fulltext_score. The parameter alpha represents the score weight between dense vector and full-text retrieval, ranging from 0 to 1, where 0 means full-text only and 1 means dense vector only:
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
	//   - Three-path recall mode:
	//
	//      - Formula: normalized_dense 	- dense_score + normalized_sparse 	- sparse_score + normalized_fulltext 	- fulltext_score. Where dense, sparse, and fulltext represent the weights for dense vector, sparse vector, and full-text retrieval respectively, with values greater than or equal to 0. The system automatically normalizes the weights to the range 0-1 (i.e., normalized_x = x / (dense + sparse + fulltext)).
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
	// The method used when building the vector index. Valid values:
	//
	// - l2: Euclidean distance.
	//
	// - ip: Inner product distance.
	//
	// - cosine: Cosine similarity.
	//
	// example:
	//
	// cosine
	Metrics *string `json:"Metrics,omitempty" xml:"Metrics,omitempty"`
	// The recall window. When this value is not empty, additional context of the retrieval results is returned. The format is a 2-element array: List<A, B>, where -10 <= A <= 0 and 0 <= B <= 10.
	//
	// > - It is recommended to use this parameter when document segmentation is too granular and retrieval may lose contextual information.
	//
	// > - Reranking takes priority over windowing, meaning reranking is performed first, then windowing is applied.
	RecallWindow []*int64 `json:"RecallWindow,omitempty" xml:"RecallWindow,omitempty" type:"Repeated"`
	// The reranking factor. When this value is not empty, the vector retrieval results are reranked. Value range: 1 < RerankFactor <= 5.
	//
	// > - Reranking is slow when document segmentation is sparse.
	//
	// > - It is recommended that the number of items to rerank (TopK 	- Factor, rounded up) does not exceed 50.
	//
	// example:
	//
	// 1.5
	RerankFactor *float64                                                                           `json:"RerankFactor,omitempty" xml:"RerankFactor,omitempty"`
	RerankModel  *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParamsRerankModel `json:"RerankModel,omitempty" xml:"RerankModel,omitempty" type:"Struct"`
	// The number of top results to return.
	//
	// example:
	//
	// 10
	TopK *int64 `json:"TopK,omitempty" xml:"TopK,omitempty"`
	// Specifies whether to use full-text retrieval (dual-path recall). Default value: false, which means only vector retrieval is used.
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
	// The number of top entities and relationship edges to return. Default value: 60.
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
	// example:
	//
	// Given a web search query, retrieve relevant passages that answer the query
	Instruct *string `json:"Instruct,omitempty" xml:"Instruct,omitempty"`
	// example:
	//
	// qwen3-rerank
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RerankMetadataFields *string `json:"RerankMetadataFields,omitempty" xml:"RerankMetadataFields,omitempty"`
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

func (s *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParamsRerankModel) GetRerankMetadataFields() *string {
	return s.RerankMetadataFields
}

func (s *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParamsRerankModel) SetInstruct(v string) *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParamsRerankModel {
	s.Instruct = &v
	return s
}

func (s *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParamsRerankModel) SetName(v string) *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParamsRerankModel {
	s.Name = &v
	return s
}

func (s *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParamsRerankModel) SetRerankMetadataFields(v string) *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParamsRerankModel {
	s.RerankMetadataFields = &v
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
	// The message list.
	//
	// This parameter is required.
	Messages []*ChatWithKnowledgeBaseRequestModelParamsMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Repeated"`
	// The name of the large model to use. For valid values, see: [Bailian Help Documentation](https://help.aliyun.com/zh/model-studio/compatibility-of-openai-with-dashscope?spm=a2c4g.11186623.help-menu-2400256.d_2_10_0.45b5516eZIciC8&scm=20140722.H_2833609._.OR_help-T_cn~zh-V_1#eadfc13038jd5)
	//
	// This parameter is required.
	//
	// example:
	//
	// qwen-plus
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// The number of candidate replies to generate.
	//
	// example:
	//
	// 1
	N *int64 `json:"N,omitempty" xml:"N,omitempty"`
	// The presence penalty coefficient (-2.0 to 2.0).
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
	// The stop word list.
	Stop []*string `json:"Stop,omitempty" xml:"Stop,omitempty" type:"Repeated"`
	// The sampling temperature (0 to 2).
	//
	// example:
	//
	// 0.6
	Temperature *float64 `json:"Temperature,omitempty" xml:"Temperature,omitempty"`
	// The tool list.
	Tools []*ChatWithKnowledgeBaseRequestModelParamsTools `json:"Tools,omitempty" xml:"Tools,omitempty" type:"Repeated"`
	// The nucleus sampling probability threshold (0 to 1).
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
	// The message content.
	//
	// This parameter is required.
	//
	// example:
	//
	// You are a helpful assistant.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The message role. Valid values:
	//
	// - system
	//
	// - user
	//
	// - assistant
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
	// The function tool description.
	//
	// example:
	//
	// 获取天气。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The function tool name.
	//
	// example:
	//
	// get_weather
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The function parameters in JSON Schema format.
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
