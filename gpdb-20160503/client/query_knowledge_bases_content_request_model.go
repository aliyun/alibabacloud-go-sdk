// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryKnowledgeBasesContentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *QueryKnowledgeBasesContentRequest
	GetContent() *string
	SetDBInstanceId(v string) *QueryKnowledgeBasesContentRequest
	GetDBInstanceId() *string
	SetMergeMethod(v string) *QueryKnowledgeBasesContentRequest
	GetMergeMethod() *string
	SetMergeMethodArgs(v *QueryKnowledgeBasesContentRequestMergeMethodArgs) *QueryKnowledgeBasesContentRequest
	GetMergeMethodArgs() *QueryKnowledgeBasesContentRequestMergeMethodArgs
	SetOwnerId(v int64) *QueryKnowledgeBasesContentRequest
	GetOwnerId() *int64
	SetRegionId(v string) *QueryKnowledgeBasesContentRequest
	GetRegionId() *string
	SetRerankFactor(v float64) *QueryKnowledgeBasesContentRequest
	GetRerankFactor() *float64
	SetSourceCollection(v []*QueryKnowledgeBasesContentRequestSourceCollection) *QueryKnowledgeBasesContentRequest
	GetSourceCollection() []*QueryKnowledgeBasesContentRequestSourceCollection
	SetTopK(v int64) *QueryKnowledgeBasesContentRequest
	GetTopK() *int64
}

type QueryKnowledgeBasesContentRequest struct {
	// The text content for retrieval.
	//
	// This parameter is required.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
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
	// The method used to merge multiple knowledge bases. Default value: RRF. Valid values:
	//
	// 	- RRF
	//
	// 	- Weight
	//
	// example:
	//
	// RRF
	MergeMethod *string `json:"MergeMethod,omitempty" xml:"MergeMethod,omitempty"`
	// The parameters of the merge method for each SourceCollection.
	MergeMethodArgs *QueryKnowledgeBasesContentRequestMergeMethodArgs `json:"MergeMethodArgs,omitempty" xml:"MergeMethodArgs,omitempty" type:"Struct"`
	OwnerId         *int64                                            `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The rerank factor. If you specify this parameter, the vector retrieval results are reranked once more. Valid values: 1\\<RerankFactor<=5.
	//
	// >
	//
	// 	- If the document is segmented into sparse parts, reranking is inefficient.
	//
	// 	- We recommend that the number of reranked results (the ceiling of TopK × RerankFactor) not exceed 50.
	//
	// example:
	//
	// 2
	RerankFactor *float64 `json:"RerankFactor,omitempty" xml:"RerankFactor,omitempty"`
	// The information about collections to retrieve from.
	//
	// This parameter is required.
	SourceCollection []*QueryKnowledgeBasesContentRequestSourceCollection `json:"SourceCollection,omitempty" xml:"SourceCollection,omitempty" type:"Repeated"`
	// Set the number of top results to be returned after merging results from multiple path retrieval.
	//
	// example:
	//
	// 10
	TopK *int64 `json:"TopK,omitempty" xml:"TopK,omitempty"`
}

func (s QueryKnowledgeBasesContentRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryKnowledgeBasesContentRequest) GoString() string {
	return s.String()
}

func (s *QueryKnowledgeBasesContentRequest) GetContent() *string {
	return s.Content
}

func (s *QueryKnowledgeBasesContentRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *QueryKnowledgeBasesContentRequest) GetMergeMethod() *string {
	return s.MergeMethod
}

func (s *QueryKnowledgeBasesContentRequest) GetMergeMethodArgs() *QueryKnowledgeBasesContentRequestMergeMethodArgs {
	return s.MergeMethodArgs
}

func (s *QueryKnowledgeBasesContentRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryKnowledgeBasesContentRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *QueryKnowledgeBasesContentRequest) GetRerankFactor() *float64 {
	return s.RerankFactor
}

func (s *QueryKnowledgeBasesContentRequest) GetSourceCollection() []*QueryKnowledgeBasesContentRequestSourceCollection {
	return s.SourceCollection
}

func (s *QueryKnowledgeBasesContentRequest) GetTopK() *int64 {
	return s.TopK
}

func (s *QueryKnowledgeBasesContentRequest) SetContent(v string) *QueryKnowledgeBasesContentRequest {
	s.Content = &v
	return s
}

func (s *QueryKnowledgeBasesContentRequest) SetDBInstanceId(v string) *QueryKnowledgeBasesContentRequest {
	s.DBInstanceId = &v
	return s
}

func (s *QueryKnowledgeBasesContentRequest) SetMergeMethod(v string) *QueryKnowledgeBasesContentRequest {
	s.MergeMethod = &v
	return s
}

func (s *QueryKnowledgeBasesContentRequest) SetMergeMethodArgs(v *QueryKnowledgeBasesContentRequestMergeMethodArgs) *QueryKnowledgeBasesContentRequest {
	s.MergeMethodArgs = v
	return s
}

func (s *QueryKnowledgeBasesContentRequest) SetOwnerId(v int64) *QueryKnowledgeBasesContentRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryKnowledgeBasesContentRequest) SetRegionId(v string) *QueryKnowledgeBasesContentRequest {
	s.RegionId = &v
	return s
}

func (s *QueryKnowledgeBasesContentRequest) SetRerankFactor(v float64) *QueryKnowledgeBasesContentRequest {
	s.RerankFactor = &v
	return s
}

func (s *QueryKnowledgeBasesContentRequest) SetSourceCollection(v []*QueryKnowledgeBasesContentRequestSourceCollection) *QueryKnowledgeBasesContentRequest {
	s.SourceCollection = v
	return s
}

func (s *QueryKnowledgeBasesContentRequest) SetTopK(v int64) *QueryKnowledgeBasesContentRequest {
	s.TopK = &v
	return s
}

func (s *QueryKnowledgeBasesContentRequest) Validate() error {
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

type QueryKnowledgeBasesContentRequestMergeMethodArgs struct {
	// The parameter that can be configured when the MergeMethod parameter is set to RRF.
	Rrf *QueryKnowledgeBasesContentRequestMergeMethodArgsRrf `json:"Rrf,omitempty" xml:"Rrf,omitempty" type:"Struct"`
	// The parameter that you can configure when you set the MergeMethod parameter to Weight.
	Weight *QueryKnowledgeBasesContentRequestMergeMethodArgsWeight `json:"Weight,omitempty" xml:"Weight,omitempty" type:"Struct"`
}

func (s QueryKnowledgeBasesContentRequestMergeMethodArgs) String() string {
	return dara.Prettify(s)
}

func (s QueryKnowledgeBasesContentRequestMergeMethodArgs) GoString() string {
	return s.String()
}

func (s *QueryKnowledgeBasesContentRequestMergeMethodArgs) GetRrf() *QueryKnowledgeBasesContentRequestMergeMethodArgsRrf {
	return s.Rrf
}

func (s *QueryKnowledgeBasesContentRequestMergeMethodArgs) GetWeight() *QueryKnowledgeBasesContentRequestMergeMethodArgsWeight {
	return s.Weight
}

func (s *QueryKnowledgeBasesContentRequestMergeMethodArgs) SetRrf(v *QueryKnowledgeBasesContentRequestMergeMethodArgsRrf) *QueryKnowledgeBasesContentRequestMergeMethodArgs {
	s.Rrf = v
	return s
}

func (s *QueryKnowledgeBasesContentRequestMergeMethodArgs) SetWeight(v *QueryKnowledgeBasesContentRequestMergeMethodArgsWeight) *QueryKnowledgeBasesContentRequestMergeMethodArgs {
	s.Weight = v
	return s
}

func (s *QueryKnowledgeBasesContentRequestMergeMethodArgs) Validate() error {
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

type QueryKnowledgeBasesContentRequestMergeMethodArgsRrf struct {
	// The smoothing constant k in the formula to calculate the score: 1/(k + rank_i). The k constant must be a positive integer greater than 1.
	//
	// example:
	//
	// 60
	K *int64 `json:"K,omitempty" xml:"K,omitempty"`
}

func (s QueryKnowledgeBasesContentRequestMergeMethodArgsRrf) String() string {
	return dara.Prettify(s)
}

func (s QueryKnowledgeBasesContentRequestMergeMethodArgsRrf) GoString() string {
	return s.String()
}

func (s *QueryKnowledgeBasesContentRequestMergeMethodArgsRrf) GetK() *int64 {
	return s.K
}

func (s *QueryKnowledgeBasesContentRequestMergeMethodArgsRrf) SetK(v int64) *QueryKnowledgeBasesContentRequestMergeMethodArgsRrf {
	s.K = &v
	return s
}

func (s *QueryKnowledgeBasesContentRequestMergeMethodArgsRrf) Validate() error {
	return dara.Validate(s)
}

type QueryKnowledgeBasesContentRequestMergeMethodArgsWeight struct {
	// An array of weights for each SourceCollection.
	Weights []*float64 `json:"Weights,omitempty" xml:"Weights,omitempty" type:"Repeated"`
}

func (s QueryKnowledgeBasesContentRequestMergeMethodArgsWeight) String() string {
	return dara.Prettify(s)
}

func (s QueryKnowledgeBasesContentRequestMergeMethodArgsWeight) GoString() string {
	return s.String()
}

func (s *QueryKnowledgeBasesContentRequestMergeMethodArgsWeight) GetWeights() []*float64 {
	return s.Weights
}

func (s *QueryKnowledgeBasesContentRequestMergeMethodArgsWeight) SetWeights(v []*float64) *QueryKnowledgeBasesContentRequestMergeMethodArgsWeight {
	s.Weights = v
	return s
}

func (s *QueryKnowledgeBasesContentRequestMergeMethodArgsWeight) Validate() error {
	return dara.Validate(s)
}

type QueryKnowledgeBasesContentRequestSourceCollection struct {
	// The name of the document collection.
	//
	// >  You can call the [CreateDocumentCollection](https://help.aliyun.com/document_detail/2618448.html) operation to create a document collection and call the [ListDocumentCollections](https://help.aliyun.com/document_detail/2618452.html) operation to query a list of document collections.
	//
	// This parameter is required.
	//
	// example:
	//
	// knowledge22
	Collection *string `json:"Collection,omitempty" xml:"Collection,omitempty"`
	// The namespace.
	//
	// >  You can call the [CreateNamespace](https://help.aliyun.com/document_detail/2401495.html) operation to create a namespace and call the [ListNamespaces](https://help.aliyun.com/document_detail/2401502.html) operation to query a list of namespaces.
	//
	// example:
	//
	// ns_cloud_index
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The password of the namespace.
	//
	// >  The value of this parameter is specified when you call the [CreateNamespace](https://help.aliyun.com/document_detail/2401495.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// ns_password
	NamespacePassword *string `json:"NamespacePassword,omitempty" xml:"NamespacePassword,omitempty"`
	// The condition that is used to filter the data to be updated. Specify this parameter in a format that is the same as the WHERE clause.
	QueryParams *QueryKnowledgeBasesContentRequestSourceCollectionQueryParams `json:"QueryParams,omitempty" xml:"QueryParams,omitempty" type:"Struct"`
}

func (s QueryKnowledgeBasesContentRequestSourceCollection) String() string {
	return dara.Prettify(s)
}

func (s QueryKnowledgeBasesContentRequestSourceCollection) GoString() string {
	return s.String()
}

func (s *QueryKnowledgeBasesContentRequestSourceCollection) GetCollection() *string {
	return s.Collection
}

func (s *QueryKnowledgeBasesContentRequestSourceCollection) GetNamespace() *string {
	return s.Namespace
}

func (s *QueryKnowledgeBasesContentRequestSourceCollection) GetNamespacePassword() *string {
	return s.NamespacePassword
}

func (s *QueryKnowledgeBasesContentRequestSourceCollection) GetQueryParams() *QueryKnowledgeBasesContentRequestSourceCollectionQueryParams {
	return s.QueryParams
}

func (s *QueryKnowledgeBasesContentRequestSourceCollection) SetCollection(v string) *QueryKnowledgeBasesContentRequestSourceCollection {
	s.Collection = &v
	return s
}

func (s *QueryKnowledgeBasesContentRequestSourceCollection) SetNamespace(v string) *QueryKnowledgeBasesContentRequestSourceCollection {
	s.Namespace = &v
	return s
}

func (s *QueryKnowledgeBasesContentRequestSourceCollection) SetNamespacePassword(v string) *QueryKnowledgeBasesContentRequestSourceCollection {
	s.NamespacePassword = &v
	return s
}

func (s *QueryKnowledgeBasesContentRequestSourceCollection) SetQueryParams(v *QueryKnowledgeBasesContentRequestSourceCollectionQueryParams) *QueryKnowledgeBasesContentRequestSourceCollection {
	s.QueryParams = v
	return s
}

func (s *QueryKnowledgeBasesContentRequestSourceCollection) Validate() error {
	if s.QueryParams != nil {
		if err := s.QueryParams.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryKnowledgeBasesContentRequestSourceCollectionQueryParams struct {
	// The filter condition that is used to query data. Specify this parameter in a format that is the same as the WHERE clause. The parameter is an expression that returns a Boolean value of TRUE or FALSE. The condition can be a simple comparison using operators such as equal (=), not equal (<> or !=), greater than (>), less than (<), greater than or equal (>=), or less than or equal (<=). It can also be a more complex expression combining multiple conditions with logical operators (AND, OR, NOT), or use keywords such as IN, BETWEEN, and LIKE.
	//
	// >
	//
	// 	- For the syntax, see https://www.postgresqltutorial.com/postgresql-tutorial/postgresql-where/.
	//
	// example:
	//
	// id = \\"llm-52tvykqt6u67iw73_j6ovptwjk7_file_6ce3da1f7e69495d9f491f2180c86973_11967297\\"
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// Whether to enable knowledge graph enhancement. Default value: false.
	//
	// example:
	//
	// true
	GraphEnhance *bool `json:"GraphEnhance,omitempty" xml:"GraphEnhance,omitempty"`
	// Returns the top number of entities and relationship edges. Default value: 60.
	GraphSearchArgs *QueryKnowledgeBasesContentRequestSourceCollectionQueryParamsGraphSearchArgs `json:"GraphSearchArgs,omitempty" xml:"GraphSearchArgs,omitempty" type:"Struct"`
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
	// Offset for pagination.
	//
	// example:
	//
	// 20
	Offset *int32 `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// The fields by which to sort the results. This parameter is empty by default.
	//
	// The field must be either a metadata field or a default field in the table (e.g., id). Supported formats include:
	//
	// Single field, such as chunk_id. Multiple fields that are separated by commas (,), such as block_id,chunk_id. Descending order is supported, such as block_id DESC,chunk_id DESC.
	//
	// example:
	//
	// file_id,sort_num
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The retrieval window. If you specify this parameter, the context of the retrieved result is added in the output. Format: List\\<A, B>. Valid values: -10<=A<=0 and 0<=B<=10.
	//
	// >
	//
	// 	- We recommend that you specify this parameter if the source document is segmented into large numbers of pieces, which may result in loss of contextual information during retrieval.
	//
	// 	- Perform re-ranking before windowing.
	RecallWindow []*int64 `json:"RecallWindow,omitempty" xml:"RecallWindow,omitempty" type:"Repeated"`
	// The rerank factor. If you specify this parameter, the vector retrieval results are reranked once more. Valid values: 1\\<RerankFactor<=5.
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
	// 776
	TopK *int64 `json:"TopK,omitempty" xml:"TopK,omitempty"`
	// Specifies whether to use full-text retrieval (dual-path retrieval). The default value is false, which means only vector retrieval is used.
	//
	// example:
	//
	// false
	UseFullTextRetrieval *bool `json:"UseFullTextRetrieval,omitempty" xml:"UseFullTextRetrieval,omitempty"`
}

func (s QueryKnowledgeBasesContentRequestSourceCollectionQueryParams) String() string {
	return dara.Prettify(s)
}

func (s QueryKnowledgeBasesContentRequestSourceCollectionQueryParams) GoString() string {
	return s.String()
}

func (s *QueryKnowledgeBasesContentRequestSourceCollectionQueryParams) GetFilter() *string {
	return s.Filter
}

func (s *QueryKnowledgeBasesContentRequestSourceCollectionQueryParams) GetGraphEnhance() *bool {
	return s.GraphEnhance
}

func (s *QueryKnowledgeBasesContentRequestSourceCollectionQueryParams) GetGraphSearchArgs() *QueryKnowledgeBasesContentRequestSourceCollectionQueryParamsGraphSearchArgs {
	return s.GraphSearchArgs
}

func (s *QueryKnowledgeBasesContentRequestSourceCollectionQueryParams) GetHybridSearch() *string {
	return s.HybridSearch
}

func (s *QueryKnowledgeBasesContentRequestSourceCollectionQueryParams) GetHybridSearchArgs() map[string]interface{} {
	return s.HybridSearchArgs
}

func (s *QueryKnowledgeBasesContentRequestSourceCollectionQueryParams) GetMetrics() *string {
	return s.Metrics
}

func (s *QueryKnowledgeBasesContentRequestSourceCollectionQueryParams) GetOffset() *int32 {
	return s.Offset
}

func (s *QueryKnowledgeBasesContentRequestSourceCollectionQueryParams) GetOrderBy() *string {
	return s.OrderBy
}

func (s *QueryKnowledgeBasesContentRequestSourceCollectionQueryParams) GetRecallWindow() []*int64 {
	return s.RecallWindow
}

func (s *QueryKnowledgeBasesContentRequestSourceCollectionQueryParams) GetRerankFactor() *float64 {
	return s.RerankFactor
}

func (s *QueryKnowledgeBasesContentRequestSourceCollectionQueryParams) GetTopK() *int64 {
	return s.TopK
}

func (s *QueryKnowledgeBasesContentRequestSourceCollectionQueryParams) GetUseFullTextRetrieval() *bool {
	return s.UseFullTextRetrieval
}

func (s *QueryKnowledgeBasesContentRequestSourceCollectionQueryParams) SetFilter(v string) *QueryKnowledgeBasesContentRequestSourceCollectionQueryParams {
	s.Filter = &v
	return s
}

func (s *QueryKnowledgeBasesContentRequestSourceCollectionQueryParams) SetGraphEnhance(v bool) *QueryKnowledgeBasesContentRequestSourceCollectionQueryParams {
	s.GraphEnhance = &v
	return s
}

func (s *QueryKnowledgeBasesContentRequestSourceCollectionQueryParams) SetGraphSearchArgs(v *QueryKnowledgeBasesContentRequestSourceCollectionQueryParamsGraphSearchArgs) *QueryKnowledgeBasesContentRequestSourceCollectionQueryParams {
	s.GraphSearchArgs = v
	return s
}

func (s *QueryKnowledgeBasesContentRequestSourceCollectionQueryParams) SetHybridSearch(v string) *QueryKnowledgeBasesContentRequestSourceCollectionQueryParams {
	s.HybridSearch = &v
	return s
}

func (s *QueryKnowledgeBasesContentRequestSourceCollectionQueryParams) SetHybridSearchArgs(v map[string]interface{}) *QueryKnowledgeBasesContentRequestSourceCollectionQueryParams {
	s.HybridSearchArgs = v
	return s
}

func (s *QueryKnowledgeBasesContentRequestSourceCollectionQueryParams) SetMetrics(v string) *QueryKnowledgeBasesContentRequestSourceCollectionQueryParams {
	s.Metrics = &v
	return s
}

func (s *QueryKnowledgeBasesContentRequestSourceCollectionQueryParams) SetOffset(v int32) *QueryKnowledgeBasesContentRequestSourceCollectionQueryParams {
	s.Offset = &v
	return s
}

func (s *QueryKnowledgeBasesContentRequestSourceCollectionQueryParams) SetOrderBy(v string) *QueryKnowledgeBasesContentRequestSourceCollectionQueryParams {
	s.OrderBy = &v
	return s
}

func (s *QueryKnowledgeBasesContentRequestSourceCollectionQueryParams) SetRecallWindow(v []*int64) *QueryKnowledgeBasesContentRequestSourceCollectionQueryParams {
	s.RecallWindow = v
	return s
}

func (s *QueryKnowledgeBasesContentRequestSourceCollectionQueryParams) SetRerankFactor(v float64) *QueryKnowledgeBasesContentRequestSourceCollectionQueryParams {
	s.RerankFactor = &v
	return s
}

func (s *QueryKnowledgeBasesContentRequestSourceCollectionQueryParams) SetTopK(v int64) *QueryKnowledgeBasesContentRequestSourceCollectionQueryParams {
	s.TopK = &v
	return s
}

func (s *QueryKnowledgeBasesContentRequestSourceCollectionQueryParams) SetUseFullTextRetrieval(v bool) *QueryKnowledgeBasesContentRequestSourceCollectionQueryParams {
	s.UseFullTextRetrieval = &v
	return s
}

func (s *QueryKnowledgeBasesContentRequestSourceCollectionQueryParams) Validate() error {
	if s.GraphSearchArgs != nil {
		if err := s.GraphSearchArgs.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryKnowledgeBasesContentRequestSourceCollectionQueryParamsGraphSearchArgs struct {
	// Returns the top number of entities and relationship edges. Default value: 60.
	//
	// example:
	//
	// 60
	GraphTopK *int64 `json:"GraphTopK,omitempty" xml:"GraphTopK,omitempty"`
}

func (s QueryKnowledgeBasesContentRequestSourceCollectionQueryParamsGraphSearchArgs) String() string {
	return dara.Prettify(s)
}

func (s QueryKnowledgeBasesContentRequestSourceCollectionQueryParamsGraphSearchArgs) GoString() string {
	return s.String()
}

func (s *QueryKnowledgeBasesContentRequestSourceCollectionQueryParamsGraphSearchArgs) GetGraphTopK() *int64 {
	return s.GraphTopK
}

func (s *QueryKnowledgeBasesContentRequestSourceCollectionQueryParamsGraphSearchArgs) SetGraphTopK(v int64) *QueryKnowledgeBasesContentRequestSourceCollectionQueryParamsGraphSearchArgs {
	s.GraphTopK = &v
	return s
}

func (s *QueryKnowledgeBasesContentRequestSourceCollectionQueryParamsGraphSearchArgs) Validate() error {
	return dara.Validate(s)
}
