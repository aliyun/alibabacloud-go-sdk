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
	SetRerankModel(v *QueryKnowledgeBasesContentRequestRerankModel) *QueryKnowledgeBasesContentRequest
	GetRerankModel() *QueryKnowledgeBasesContentRequestRerankModel
	SetSourceCollection(v []*QueryKnowledgeBasesContentRequestSourceCollection) *QueryKnowledgeBasesContentRequest
	GetSourceCollection() []*QueryKnowledgeBasesContentRequestSourceCollection
	SetTopK(v int64) *QueryKnowledgeBasesContentRequest
	GetTopK() *int64
}

type QueryKnowledgeBasesContentRequest struct {
	// The text content used for retrieval.
	//
	// This parameter is required.
	//
	// example:
	//
	// What is ADBPG?
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The instance ID.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the details of all AnalyticDB for PostgreSQL instances in a region, including instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The method used to merge results from multiple knowledge bases. Default value: RRF. Valid values:
	//
	// - RRF
	//
	// - Weight.
	//
	// example:
	//
	// RRF
	MergeMethod *string `json:"MergeMethod,omitempty" xml:"MergeMethod,omitempty"`
	// The parameters for the merge method of each SourceCollection.
	MergeMethodArgs *QueryKnowledgeBasesContentRequestMergeMethodArgs `json:"MergeMethodArgs,omitempty" xml:"MergeMethodArgs,omitempty" type:"Struct"`
	OwnerId         *int64                                            `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The reranking factor. If this parameter is not empty, the vector retrieval results are reranked. Valid values: 1 < RerankFactor <= 5.
	//
	// > - Reranking is slow when document chunks are sparse.
	//
	// > - The recommended reranking count (TopK × Factor, rounded up) should not exceed 50.
	//
	// example:
	//
	// 2
	RerankFactor *float64 `json:"RerankFactor,omitempty" xml:"RerankFactor,omitempty"`
	// The reranking model parameters for performing an additional reranking on the overall results after multi-channel merging.
	RerankModel *QueryKnowledgeBasesContentRequestRerankModel `json:"RerankModel,omitempty" xml:"RerankModel,omitempty" type:"Struct"`
	// The information about the multiple collections to retrieve.
	//
	// This parameter is required.
	SourceCollection []*QueryKnowledgeBasesContentRequestSourceCollection `json:"SourceCollection,omitempty" xml:"SourceCollection,omitempty" type:"Repeated"`
	// The number of top results to return after multi-channel recall merging.
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

func (s *QueryKnowledgeBasesContentRequest) GetRerankModel() *QueryKnowledgeBasesContentRequestRerankModel {
	return s.RerankModel
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

func (s *QueryKnowledgeBasesContentRequest) SetRerankModel(v *QueryKnowledgeBasesContentRequestRerankModel) *QueryKnowledgeBasesContentRequest {
	s.RerankModel = v
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

type QueryKnowledgeBasesContentRequestMergeMethodArgs struct {
	// The configurable parameters when MergeMethod is set to RRF.
	Rrf *QueryKnowledgeBasesContentRequestMergeMethodArgsRrf `json:"Rrf,omitempty" xml:"Rrf,omitempty" type:"Struct"`
	// The configurable parameters when MergeMethod is set to Weight.
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
	// The k constant in the scoring algorithm `1/(k+rank_i)`. The value must be a positive integer greater than 1.
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
	// The weight array for each SourceCollection.
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

type QueryKnowledgeBasesContentRequestRerankModel struct {
	// This parameter can be set when RerankModel.Name is set to qwen3-rerank. Specifies a custom ranking task type description that guides the model to adopt different ranking strategies.
	//
	// example:
	//
	// Given a web search query, retrieve relevant passages that answer the query
	Instruct *string `json:"Instruct,omitempty" xml:"Instruct,omitempty"`
	// The name of the reranking model. Valid values: qwen3-rerank, gte-rerank-v2.
	//
	// example:
	//
	// qwen3-rerank
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s QueryKnowledgeBasesContentRequestRerankModel) String() string {
	return dara.Prettify(s)
}

func (s QueryKnowledgeBasesContentRequestRerankModel) GoString() string {
	return s.String()
}

func (s *QueryKnowledgeBasesContentRequestRerankModel) GetInstruct() *string {
	return s.Instruct
}

func (s *QueryKnowledgeBasesContentRequestRerankModel) GetName() *string {
	return s.Name
}

func (s *QueryKnowledgeBasesContentRequestRerankModel) SetInstruct(v string) *QueryKnowledgeBasesContentRequestRerankModel {
	s.Instruct = &v
	return s
}

func (s *QueryKnowledgeBasesContentRequestRerankModel) SetName(v string) *QueryKnowledgeBasesContentRequestRerankModel {
	s.Name = &v
	return s
}

func (s *QueryKnowledgeBasesContentRequestRerankModel) Validate() error {
	return dara.Validate(s)
}

type QueryKnowledgeBasesContentRequestSourceCollection struct {
	// The name of the document collection.
	//
	// > The document collection is created by calling the [CreateDocumentCollection](https://help.aliyun.com/document_detail/2618448.html) operation. You can call the [ListDocumentCollections](https://help.aliyun.com/document_detail/2618452.html) operation to view existing document collections.
	//
	// This parameter is required.
	//
	// example:
	//
	// knowledge22
	Collection *string `json:"Collection,omitempty" xml:"Collection,omitempty"`
	// The namespace.
	//
	// > You can create a namespace by calling the [CreateNamespace](https://help.aliyun.com/document_detail/2401495.html) operation and view the list by calling the [ListNamespaces](https://help.aliyun.com/document_detail/2401502.html) operation.
	//
	// example:
	//
	// ns_cloud_index
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The password of the namespace.
	//
	// > This value is specified by the [CreateNamespace](https://help.aliyun.com/document_detail/2401495.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// ns_password
	NamespacePassword *string `json:"NamespacePassword,omitempty" xml:"NamespacePassword,omitempty"`
	// The filter conditions for the data to query, in SQL WHERE clause format.
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
	// The filter conditions for the data to query, in SQL WHERE clause format. This is an expression that returns a Boolean value (true or false). The conditions can be simple comparison operators such as equal to (=), not equal to (<> or !=), greater than (>), less than (<), greater than or equal to (>=), and less than or equal to (<=). They can also be more complex expressions combined with logical operators (AND, OR, NOT), as well as conditions that use keywords such as IN, BETWEEN, and LIKE.
	//
	// >
	//
	// > - For detailed syntax, refer to: https://www.postgresqltutorial.com/postgresql-tutorial/postgresql-where/.
	//
	// example:
	//
	// id = \\"llm-52tvykqt6u67iw73_j6ovptwjk7_file_6ce3da1f7e69495d9f491f2180c86973_11967297\\"
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// Specifies whether to enable knowledge graph enhancement. Default value: false.
	//
	// example:
	//
	// true
	GraphEnhance *bool `json:"GraphEnhance,omitempty" xml:"GraphEnhance,omitempty"`
	// The number of top entities and relationship edges to return. Default value: 60.
	GraphSearchArgs *QueryKnowledgeBasesContentRequestSourceCollectionQueryParamsGraphSearchArgs `json:"GraphSearchArgs,omitempty" xml:"GraphSearchArgs,omitempty" type:"Struct"`
	// The multi-channel recall algorithm. Default value: empty (the scores from dense vectors and full-text retrieval are directly compared and sorted).
	//
	// Valid values:
	//
	// - RRF: reciprocal rank fusion. A parameter k controls the fusion effect. For more information, see the HybridSearchArgs configuration.
	//
	// - Weight: weighted ranking. Parameters control the score weights of vector retrieval and full-text retrieval before sorting. For more information, see the HybridSearchArgs configuration.
	//
	// - Cascaded: full-text retrieval is performed first, followed by vector retrieval on the full-text results.
	//
	// example:
	//
	// Cascaded
	HybridSearch *string `json:"HybridSearch,omitempty" xml:"HybridSearch,omitempty"`
	// The algorithm parameters for multi-channel recall. RRF and Weight are supported. HybridPathsSetting specifies the recall paths: dense vectors (dense), sparse vectors (sparse), and full-text retrieval (fulltext). If this value is empty, dense vectors (dense) and full-text retrieval (fulltext) are used by default.
	//
	// - RRF: specifies the k constant in the scoring algorithm `1/(k+rank_i)`. The value must be a positive integer greater than 1. Format:
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
	//       - Formula: alpha 	- dense_score + (1-alpha) 	- fulltext_score. The alpha parameter specifies the score weight between dense vectors and full-text retrieval. Valid values: 0 to 1, where 0 indicates full-text retrieval only and 1 indicates dense vectors only:
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
	//   - Three-path recall pattern:
	//
	//      - Formula: normalized_dense 	- dense_score + normalized_sparse 	- sparse_score + normalized_fulltext 	- fulltext_score. dense, sparse, and fulltext represent the weights for dense vectors, sparse vectors, and full-text retrieval respectively. Valid values: greater than or equal to 0. The system automatically performs normalization on the weights to 0 to 1 (normalized_x = x / (dense + sparse + fulltext)).
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
	// The method used to build the vector index. Valid values:
	//
	// - l2: Euclidean distance.
	//
	// - ip: inner product distance.
	//
	// - cosine: cosine similarity.
	//
	// example:
	//
	// cosine
	Metrics *string `json:"Metrics,omitempty" xml:"Metrics,omitempty"`
	// The offset for paging query.
	//
	// example:
	//
	// 20
	Offset *int32 `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// The field used for sorting. Default value: empty.
	//
	// The field must belong to metadata or a default field in the table, such as id. Supported formats:
	//
	// A single field, such as chunk_id.
	//
	// Multiple fields separated by commas, such as block_id, chunk_id.
	//
	// Descending order is supported, such as block_id DESC, chunk_id DESC.
	//
	// example:
	//
	// file_id,sort_num
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The recall window. If this value is not empty, the context of the retrieval results is included. The format is a two-element array: List<A, B>, where -10 <= A <= 0 and 0 <= B <= 10.
	//
	// > - Use this parameter when document chunks are too small and retrieval may lose context information.
	//
	// > - Reranking takes priority over windowing. Reranking is performed first, followed by windowing.
	RecallWindow []*int64 `json:"RecallWindow,omitempty" xml:"RecallWindow,omitempty" type:"Repeated"`
	// The reranking factor. If this parameter is not empty, the vector retrieval results are reranked. Valid values: 1 < RerankFactor <= 5.
	//
	// > - Reranking is slow when document chunks are sparse.
	//
	// > - The recommended reranking count (TopK × Factor, rounded up) should not exceed 50.
	//
	// example:
	//
	// 2.0
	RerankFactor *float64 `json:"RerankFactor,omitempty" xml:"RerankFactor,omitempty"`
	// The reranking model parameters.
	RerankModel *QueryKnowledgeBasesContentRequestSourceCollectionQueryParamsRerankModel `json:"RerankModel,omitempty" xml:"RerankModel,omitempty" type:"Struct"`
	// The number of top results to return.
	//
	// example:
	//
	// 776
	TopK *int64 `json:"TopK,omitempty" xml:"TopK,omitempty"`
	// Specifies whether to use full-text retrieval (dual-path recall). Default value: false, which indicates that only vector retrieval is used.
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

func (s *QueryKnowledgeBasesContentRequestSourceCollectionQueryParams) GetRerankModel() *QueryKnowledgeBasesContentRequestSourceCollectionQueryParamsRerankModel {
	return s.RerankModel
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

func (s *QueryKnowledgeBasesContentRequestSourceCollectionQueryParams) SetRerankModel(v *QueryKnowledgeBasesContentRequestSourceCollectionQueryParamsRerankModel) *QueryKnowledgeBasesContentRequestSourceCollectionQueryParams {
	s.RerankModel = v
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
	if s.RerankModel != nil {
		if err := s.RerankModel.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryKnowledgeBasesContentRequestSourceCollectionQueryParamsGraphSearchArgs struct {
	// The number of top entities and relationship edges to return. Default value: 60.
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

type QueryKnowledgeBasesContentRequestSourceCollectionQueryParamsRerankModel struct {
	// This parameter can be set when RerankModel.Name is set to qwen3-rerank. Specifies a custom ranking task type description that guides the model to adopt different ranking strategies.
	//
	// example:
	//
	// Given a web search query, retrieve relevant passages that answer the query
	Instruct *string `json:"Instruct,omitempty" xml:"Instruct,omitempty"`
	// The name of the reranking model. Valid values: qwen3-rerank, gte-rerank-v2.
	//
	// example:
	//
	// qwen3-rerank
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RerankMetadataFields *string `json:"RerankMetadataFields,omitempty" xml:"RerankMetadataFields,omitempty"`
}

func (s QueryKnowledgeBasesContentRequestSourceCollectionQueryParamsRerankModel) String() string {
	return dara.Prettify(s)
}

func (s QueryKnowledgeBasesContentRequestSourceCollectionQueryParamsRerankModel) GoString() string {
	return s.String()
}

func (s *QueryKnowledgeBasesContentRequestSourceCollectionQueryParamsRerankModel) GetInstruct() *string {
	return s.Instruct
}

func (s *QueryKnowledgeBasesContentRequestSourceCollectionQueryParamsRerankModel) GetName() *string {
	return s.Name
}

func (s *QueryKnowledgeBasesContentRequestSourceCollectionQueryParamsRerankModel) GetRerankMetadataFields() *string {
	return s.RerankMetadataFields
}

func (s *QueryKnowledgeBasesContentRequestSourceCollectionQueryParamsRerankModel) SetInstruct(v string) *QueryKnowledgeBasesContentRequestSourceCollectionQueryParamsRerankModel {
	s.Instruct = &v
	return s
}

func (s *QueryKnowledgeBasesContentRequestSourceCollectionQueryParamsRerankModel) SetName(v string) *QueryKnowledgeBasesContentRequestSourceCollectionQueryParamsRerankModel {
	s.Name = &v
	return s
}

func (s *QueryKnowledgeBasesContentRequestSourceCollectionQueryParamsRerankModel) SetRerankMetadataFields(v string) *QueryKnowledgeBasesContentRequestSourceCollectionQueryParamsRerankModel {
	s.RerankMetadataFields = &v
	return s
}

func (s *QueryKnowledgeBasesContentRequestSourceCollectionQueryParamsRerankModel) Validate() error {
	return dara.Validate(s)
}
