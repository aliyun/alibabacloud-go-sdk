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
	// The cluster ID.
	//
	// >  You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the information about all AnalyticDB for PostgreSQL instances within a region, including instance IDs.
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
	KnowledgeParams *ChatWithKnowledgeBaseStreamRequestKnowledgeParams `json:"KnowledgeParams,omitempty" xml:"KnowledgeParams,omitempty" type:"Struct"`
	// The Large Language Model (LLM) invocation parameter object.
	//
	// This parameter is required.
	ModelParams *ChatWithKnowledgeBaseStreamRequestModelParams `json:"ModelParams,omitempty" xml:"ModelParams,omitempty" type:"Struct"`
	OwnerId     *int64                                         `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The system prompt template, which should include {{ text_chunks }},{{ user_system_prompt }},{{ graph_entities },{{ graph_relations }}. If any of these placeholders are not specified, the corresponding section should have no effect.
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
	// The method used to merge multiple knowledge base. Default value: RRF. Valid values:
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
	MergeMethodArgs *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsMergeMethodArgs `json:"MergeMethodArgs,omitempty" xml:"MergeMethodArgs,omitempty" type:"Struct"`
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
	// 5.0
	RerankFactor *float64 `json:"RerankFactor,omitempty" xml:"RerankFactor,omitempty"`
	// Knowledge base.
	//
	// This parameter is required.
	SourceCollection []*ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollection `json:"SourceCollection,omitempty" xml:"SourceCollection,omitempty" type:"Repeated"`
	// Specifies the number of top results to return after merging retrieved results from multiple vector collections.
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
	// The parameter to be configured when the MergeMethod parameter is set to RRF.
	Rrf *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsMergeMethodArgsRrf `json:"Rrf,omitempty" xml:"Rrf,omitempty" type:"Struct"`
	// The smoothing constant in the formula to calculate the score: 1/(k + rank_i). It must be a positive integer greater than 1.
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
	// Formula to calculate the score: 1/(k + rank_i). The k constant must be a positive integer greater than 1.
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
	// An array of weights for each SourceCollection.
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

type ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollection struct {
	// The name of the collection to be recalled.
	//
	// This parameter is required.
	//
	// example:
	//
	// cloud_index_adb_50943_prod
	Collection *string `json:"Collection,omitempty" xml:"Collection,omitempty"`
	// The namespace.
	//
	// >  You can call the [ListNamespaces](https://help.aliyun.com/document_detail/2401502.html) operation to query a list of namespaces.
	//
	// example:
	//
	// ddstar_vector
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The password for the namespace.
	//
	// > The value of this parameter is specified by the CreateNamespace operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// namespacePassword
	NamespacePassword *string `json:"NamespacePassword,omitempty" xml:"NamespacePassword,omitempty"`
	// Parameters related to the knowledge base retrieval.
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
	// The condition that is used to filter the data to be updated. Specify this parameter in a format that is the same as the WHERE clause.
	//
	// example:
	//
	// method_id=\\"e41695f0-2851-40ac-b21d-dd337b60d71c\\"
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// Whether to enable knowledge graph enhancement. Default value: false.
	//
	// example:
	//
	// true
	GraphEnhance *bool `json:"GraphEnhance,omitempty" xml:"GraphEnhance,omitempty"`
	// The knowledge graph retrieval parameters.
	GraphSearchArgs *ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollectionQueryParamsGraphSearchArgs `json:"GraphSearchArgs,omitempty" xml:"GraphSearchArgs,omitempty" type:"Struct"`
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
	// Cascaded
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
	// 	- Prioritized reranking with windowing, i.e., perform reranking first followed by windowing processing.
	RecallWindow []*int64 `json:"RecallWindow,omitempty" xml:"RecallWindow,omitempty" type:"Repeated"`
	// The rerank factor. If you specify this parameter, the retrieved results are reranked once again. Valid values: 1\\<RerankFactor<=5.
	//
	// >
	//
	// 	- If the document is segmented into sparse parts, reranking is inefficient.
	//
	// 	- We recommend that the number of reranked results (the ceiling of TopK × RerankFactor) not exceed 50.
	//
	// example:
	//
	// 2.0
	RerankFactor *float64 `json:"RerankFactor,omitempty" xml:"RerankFactor,omitempty"`
	// The number of top results.
	//
	// example:
	//
	// 101
	TopK *int64 `json:"TopK,omitempty" xml:"TopK,omitempty"`
	// Specifies whether to use full-text retrieval (dual-path retrieval). The default value is false, which means only vector retrieval is used.
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
	return nil
}

type ChatWithKnowledgeBaseStreamRequestKnowledgeParamsSourceCollectionQueryParamsGraphSearchArgs struct {
	// The number of top entities and relationship edges. Default value: 60.
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

type ChatWithKnowledgeBaseStreamRequestModelParams struct {
	// Maximum number of tokens to generate.
	//
	// example:
	//
	// 8192
	MaxTokens *int64 `json:"MaxTokens,omitempty" xml:"MaxTokens,omitempty"`
	// Message list.
	//
	// This parameter is required.
	Messages []*ChatWithKnowledgeBaseStreamRequestModelParamsMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Repeated"`
	// The model name. See [Model Studio Document](https://help.aliyun.com/zh/model-studio/compatibility-of-openai-with-dashscope?spm=openapi-amp.newDocPublishment.0.0.257c281fH8TtM8\\&scm=20140722.H_2833609._.OR_help-T_cn~zh-V_1#eadfc13038jd5) for the available models.
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
	// Tools.
	Tools []*ChatWithKnowledgeBaseStreamRequestModelParamsTools `json:"Tools,omitempty" xml:"Tools,omitempty" type:"Repeated"`
	// Top-p (nucleus) sampling threshold (0–1).
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
	// The information about a function.
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
