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
	// This parameter is required.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// example:
	//
	// RRF
	MergeMethod     *string                                           `json:"MergeMethod,omitempty" xml:"MergeMethod,omitempty"`
	MergeMethodArgs *QueryKnowledgeBasesContentRequestMergeMethodArgs `json:"MergeMethodArgs,omitempty" xml:"MergeMethodArgs,omitempty" type:"Struct"`
	OwnerId         *int64                                            `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 2
	RerankFactor *float64 `json:"RerankFactor,omitempty" xml:"RerankFactor,omitempty"`
	// This parameter is required.
	SourceCollection []*QueryKnowledgeBasesContentRequestSourceCollection `json:"SourceCollection,omitempty" xml:"SourceCollection,omitempty" type:"Repeated"`
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
	Rrf    *QueryKnowledgeBasesContentRequestMergeMethodArgsRrf    `json:"Rrf,omitempty" xml:"Rrf,omitempty" type:"Struct"`
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
	// This parameter is required.
	//
	// example:
	//
	// knowledge22
	Collection *string `json:"Collection,omitempty" xml:"Collection,omitempty"`
	// example:
	//
	// ns_cloud_index
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ns_password
	NamespacePassword *string                                                       `json:"NamespacePassword,omitempty" xml:"NamespacePassword,omitempty"`
	QueryParams       *QueryKnowledgeBasesContentRequestSourceCollectionQueryParams `json:"QueryParams,omitempty" xml:"QueryParams,omitempty" type:"Struct"`
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
	// example:
	//
	// id = \\"llm-52tvykqt6u67iw73_j6ovptwjk7_file_6ce3da1f7e69495d9f491f2180c86973_11967297\\"
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// example:
	//
	// true
	GraphEnhance    *bool                                                                        `json:"GraphEnhance,omitempty" xml:"GraphEnhance,omitempty"`
	GraphSearchArgs *QueryKnowledgeBasesContentRequestSourceCollectionQueryParamsGraphSearchArgs `json:"GraphSearchArgs,omitempty" xml:"GraphSearchArgs,omitempty" type:"Struct"`
	// example:
	//
	// Cascaded
	HybridSearch     *string                `json:"HybridSearch,omitempty" xml:"HybridSearch,omitempty"`
	HybridSearchArgs map[string]interface{} `json:"HybridSearchArgs,omitempty" xml:"HybridSearchArgs,omitempty"`
	// example:
	//
	// cosine
	Metrics *string `json:"Metrics,omitempty" xml:"Metrics,omitempty"`
	// example:
	//
	// 20
	Offset *int32 `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// example:
	//
	// file_id,sort_num
	OrderBy      *string  `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	RecallWindow []*int64 `json:"RecallWindow,omitempty" xml:"RecallWindow,omitempty" type:"Repeated"`
	// example:
	//
	// 2.0
	RerankFactor *float64 `json:"RerankFactor,omitempty" xml:"RerankFactor,omitempty"`
	// example:
	//
	// 776
	TopK *int64 `json:"TopK,omitempty" xml:"TopK,omitempty"`
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
