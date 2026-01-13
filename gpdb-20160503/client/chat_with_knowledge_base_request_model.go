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
	// The cluster ID.
	//
	// >  You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/196830.html) operation to query the information about all AnalyticDB for PostgreSQL instances within a region, including instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// Whether to return the retrieved result. Default value: false.
	//
	// example:
	//
	// false
	IncludeKnowledgeBaseResults *bool `json:"IncludeKnowledgeBaseResults,omitempty" xml:"IncludeKnowledgeBaseResults,omitempty"`
	// The knowledge retrieval parameter object. If you do not specify this parameter, only chat mode is enabled.
	KnowledgeParams *ChatWithKnowledgeBaseRequestKnowledgeParams `json:"KnowledgeParams,omitempty" xml:"KnowledgeParams,omitempty" type:"Struct"`
	// The Large Language Model (LLM) invocation parameter object.
	//
	// This parameter is required.
	ModelParams *ChatWithKnowledgeBaseRequestModelParams `json:"ModelParams,omitempty" xml:"ModelParams,omitempty" type:"Struct"`
	OwnerId     *int64                                   `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The system prompt template, which should include {{ text_chunks }},{{ user_system_prompt }},{{ graph_entities },{{ graph_relations }}. If any of these placeholders are not specified, the corresponding section should have no effect.
	PromptParams *string `json:"PromptParams,omitempty" xml:"PromptParams,omitempty"`
	// 实例所在的地域ID
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
	// The method used to merge multiple knowledge bases. Default value: RRF. Optional:
	//
	// 	- RRF
	//
	// 	- Weight
	//
	// example:
	//
	// "RRF"
	MergeMethod *string `json:"MergeMethod,omitempty" xml:"MergeMethod,omitempty"`
	// Parameters for multi-knowledge-base fusion.
	MergeMethodArgs *ChatWithKnowledgeBaseRequestKnowledgeParamsMergeMethodArgs `json:"MergeMethodArgs,omitempty" xml:"MergeMethodArgs,omitempty" type:"Struct"`
	// The rerank factor. If you specify this parameter, the search result is reranked once again. Valid values: 1\\<RerankFactor<=5.
	//
	// >
	//
	// 	- If the document is segmented into sparse parts, reranking is inefficient.
	//
	// 	- We recommend that the number of reranked results (the ceiling of TopK × RerankFactor) not exceed 50.
	//
	// example:
	//
	// 1.0001
	RerankFactor *float64 `json:"RerankFactor,omitempty" xml:"RerankFactor,omitempty"`
	// Knowledge base.
	//
	// This parameter is required.
	SourceCollection []*ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollection `json:"SourceCollection,omitempty" xml:"SourceCollection,omitempty" type:"Repeated"`
	// Specifies the number of top results to return after merging retrieved results from multiple vector collections.
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
	// The parameter that can be configured when the MergeMethod parameter is set to RRF.
	Rrf *ChatWithKnowledgeBaseRequestKnowledgeParamsMergeMethodArgsRrf `json:"Rrf,omitempty" xml:"Rrf,omitempty" type:"Struct"`
	// The parameter that you can configure when you set the MergeMethod parameter to Weight.
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
	// The smoothing constant k in the formula to calculate the score: 1/(k + rank_i). It must be a positive integer greater than 1.
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
	// An array of weights for each SourceCollection.
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

type ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollection struct {
	// The name of the collection to be recalled.
	//
	// This parameter is required.
	//
	// example:
	//
	// adbpg_document_collection
	Collection *string `json:"Collection,omitempty" xml:"Collection,omitempty"`
	// The name of the namespace. Default value: public.
	//
	// >  You can call the [CreateNamespace](https://help.aliyun.com/document_detail/2401495.html) operation to create a namespace and call the [ListNamespaces](https://help.aliyun.com/document_detail/2401502.html) operation to query a list of namespaces.
	//
	// example:
	//
	// dukang
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The password of the namespace.
	//
	// >  The value of this parameter is specified when you call the CreateNamespace operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// namespacePasswd
	NamespacePassword *string `json:"NamespacePassword,omitempty" xml:"NamespacePassword,omitempty"`
	// Parameters related to the knowledge base retrieval.
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
	// The condition that is used to filter the data to be updated. Specify this parameter in a format that is the same as the WHERE clause.
	//
	// example:
	//
	// id = \\"llm-t87l87fxuhn56woc_8anu8j2d3f_file_e74635e2cc314e838543e724f6b3b1f2_10658020\\"
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// Whether to enable knowledge graph enhancement. Default value: false.
	//
	// example:
	//
	// false
	GraphEnhance *bool `json:"GraphEnhance,omitempty" xml:"GraphEnhance,omitempty"`
	// Returns the top number of entities and relationship edges. Default value: 60.
	GraphSearchArgs *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParamsGraphSearchArgs `json:"GraphSearchArgs,omitempty" xml:"GraphSearchArgs,omitempty" type:"Struct"`
	// The dual-path retrieval algorithm. This parameter is empty by default, which specifies that scores of vector retrieval and full-text retrieval are directly compared and sorted together.
	//
	// Valid values:
	//
	// 	- RRF: The reciprocal rank fusion (RRF) algorithm uses a constant k to control the fusion effect. For more information, see the description of the HybridSearchArgs parameter.
	//
	// 	- Weight: This algorithm uses the alpha parameter to specify the proportion of the vector search score and the full-text search score and then sorts by weight. For more information, see the description of the HybridSearchArgs parameter.
	//
	// 	- Cascaded: This algorithm performs first full-text retrieval and then vector retrieval.
	//
	// example:
	//
	// RRF
	HybridSearch *string `json:"HybridSearch,omitempty" xml:"HybridSearch,omitempty"`
	// The parameters of the dual-path retrieval algorithm. RRF and Weight are supported at this time:
	//
	// 	- RRF: Specifies the smoothing constant k in the formula to calculate the score: `1/(k + rank_i)`. The k constant must be a positive integer greater than 1. The format:
	//
	// <!---->
	//
	//     {
	//
	//        "RRF": {
	//
	//         "k": 60
	//
	//        }
	//
	//     }
	//
	// 	- Weight: The score is computed as `alpha 	- vector_score + (1 - alpha) 	- text_score`. The parameter alpha controls the weighting between vector search and full-text search scores, with a valid range of [0, 1]. 0 specifies only full-text search score. 1 specifies only vector search score.
	//
	// <!---->
	//
	//     {
	//
	//        "Weight": {
	//
	//         "alpha": 0.5
	//
	//        }
	//
	//     }
	HybridSearchArgs map[string]interface{} `json:"HybridSearchArgs,omitempty" xml:"HybridSearchArgs,omitempty"`
	// The method that is used to create vector indexes. Valid values:
	//
	// 	- l2: Euclidean distance.
	//
	// 	- ip: Inner product distance.
	//
	// 	- cosine: Cosine similarity.
	//
	// example:
	//
	// cosine
	Metrics *string `json:"Metrics,omitempty" xml:"Metrics,omitempty"`
	// The retrieval window. If you specify this parameter, the context of the retrieved result is added in the output. Format: List\\<A, B>. Valid values: -10<=A<=0 and 0<=B<=10.
	//
	// >
	//
	// 	- We recommend that you specify this parameter if the source document is segmented into large numbers of pieces, which may result in loss of contextual information during retrieval.
	//
	// 	- Perform re-ranking before windowing.
	RecallWindow []*int64 `json:"RecallWindow,omitempty" xml:"RecallWindow,omitempty" type:"Repeated"`
	// The rerank factor. If you specify this parameter, the search result is reranked once again. Valid values: 1\\<RerankFactor<=5.
	//
	// >
	//
	// 	- If the document is segmented into sparse parts, reranking is inefficient.
	//
	// 	- We recommend that the number of reranked results (the ceiling of TopK × RerankFactor) not exceed 50.
	//
	// example:
	//
	// 1.5
	RerankFactor *float64 `json:"RerankFactor,omitempty" xml:"RerankFactor,omitempty"`
	// The number of top results.
	//
	// example:
	//
	// 10
	TopK *int64 `json:"TopK,omitempty" xml:"TopK,omitempty"`
	// Specifies whether to use full-text retrieval (dual-path retrieval). The default value is false, which means only vector retrieval is used.
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
	return nil
}

type ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParamsGraphSearchArgs struct {
	// Returns the top number of entities and relationship edges. Default value: 60.
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

type ChatWithKnowledgeBaseRequestModelParams struct {
	// Maximum number of tokens to generate.
	//
	// example:
	//
	// 8192
	MaxTokens *int64 `json:"MaxTokens,omitempty" xml:"MaxTokens,omitempty"`
	// Message list.
	//
	// This parameter is required.
	Messages []*ChatWithKnowledgeBaseRequestModelParamsMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Repeated"`
	// The model name. See [Model Studio Document](https://help.aliyun.com/zh/model-studio/compatibility-of-openai-with-dashscope?spm=a2c4g.11186623.help-menu-2400256.d_2_10_0.45b5516eZIciC8\\&scm=20140722.H_2833609._.OR_help-T_cn~zh-V_1#eadfc13038jd5) for the available models.
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
	// Presence penalty coefficient (-2.0 to 2.0).
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
	// Stop words.
	Stop []*string `json:"Stop,omitempty" xml:"Stop,omitempty" type:"Repeated"`
	// Sampling temperature (0~2).
	//
	// example:
	//
	// 0.6
	Temperature *float64 `json:"Temperature,omitempty" xml:"Temperature,omitempty"`
	// Tools
	Tools []*ChatWithKnowledgeBaseRequestModelParamsTools `json:"Tools,omitempty" xml:"Tools,omitempty" type:"Repeated"`
	// Top-p (nucleus) sampling threshold (0–1).
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
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The message role. Valid values:
	//
	// 	- system
	//
	// 	- user
	//
	// 	- assistant
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
	// The information about a function.
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
	// The description of the function.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the function.
	//
	// example:
	//
	// get_weather
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// JSON Schema for function parameters.
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
