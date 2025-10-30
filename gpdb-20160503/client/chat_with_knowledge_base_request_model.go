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
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// example:
	//
	// false
	IncludeKnowledgeBaseResults *bool                                        `json:"IncludeKnowledgeBaseResults,omitempty" xml:"IncludeKnowledgeBaseResults,omitempty"`
	KnowledgeParams             *ChatWithKnowledgeBaseRequestKnowledgeParams `json:"KnowledgeParams,omitempty" xml:"KnowledgeParams,omitempty" type:"Struct"`
	// This parameter is required.
	ModelParams  *ChatWithKnowledgeBaseRequestModelParams `json:"ModelParams,omitempty" xml:"ModelParams,omitempty" type:"Struct"`
	OwnerId      *int64                                   `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PromptParams *string                                  `json:"PromptParams,omitempty" xml:"PromptParams,omitempty"`
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
	// example:
	//
	// "RRF"
	MergeMethod     *string                                                     `json:"MergeMethod,omitempty" xml:"MergeMethod,omitempty"`
	MergeMethodArgs *ChatWithKnowledgeBaseRequestKnowledgeParamsMergeMethodArgs `json:"MergeMethodArgs,omitempty" xml:"MergeMethodArgs,omitempty" type:"Struct"`
	// example:
	//
	// 1.0001
	RerankFactor *float64 `json:"RerankFactor,omitempty" xml:"RerankFactor,omitempty"`
	// This parameter is required.
	SourceCollection []*ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollection `json:"SourceCollection,omitempty" xml:"SourceCollection,omitempty" type:"Repeated"`
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
	Rrf    *ChatWithKnowledgeBaseRequestKnowledgeParamsMergeMethodArgsRrf    `json:"Rrf,omitempty" xml:"Rrf,omitempty" type:"Struct"`
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
	// This parameter is required.
	//
	// example:
	//
	// adbpg_document_collection
	Collection *string `json:"Collection,omitempty" xml:"Collection,omitempty"`
	// example:
	//
	// dukang
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// namespacePasswd
	NamespacePassword *string                                                                 `json:"NamespacePassword,omitempty" xml:"NamespacePassword,omitempty"`
	QueryParams       *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParams `json:"QueryParams,omitempty" xml:"QueryParams,omitempty" type:"Struct"`
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
	// example:
	//
	// id = \\"llm-t87l87fxuhn56woc_8anu8j2d3f_file_e74635e2cc314e838543e724f6b3b1f2_10658020\\"
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// example:
	//
	// false
	GraphEnhance    *bool                                                                                  `json:"GraphEnhance,omitempty" xml:"GraphEnhance,omitempty"`
	GraphSearchArgs *ChatWithKnowledgeBaseRequestKnowledgeParamsSourceCollectionQueryParamsGraphSearchArgs `json:"GraphSearchArgs,omitempty" xml:"GraphSearchArgs,omitempty" type:"Struct"`
	// example:
	//
	// RRF
	HybridSearch     *string                `json:"HybridSearch,omitempty" xml:"HybridSearch,omitempty"`
	HybridSearchArgs map[string]interface{} `json:"HybridSearchArgs,omitempty" xml:"HybridSearchArgs,omitempty"`
	// example:
	//
	// cosine
	Metrics      *string  `json:"Metrics,omitempty" xml:"Metrics,omitempty"`
	RecallWindow []*int64 `json:"RecallWindow,omitempty" xml:"RecallWindow,omitempty" type:"Repeated"`
	// example:
	//
	// 1.5
	RerankFactor *float64 `json:"RerankFactor,omitempty" xml:"RerankFactor,omitempty"`
	// example:
	//
	// 10
	TopK *int64 `json:"TopK,omitempty" xml:"TopK,omitempty"`
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
	// example:
	//
	// 8192
	MaxTokens *int64 `json:"MaxTokens,omitempty" xml:"MaxTokens,omitempty"`
	// This parameter is required.
	Messages []*ChatWithKnowledgeBaseRequestModelParamsMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// qwen-plus
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// example:
	//
	// 1
	N *int64 `json:"N,omitempty" xml:"N,omitempty"`
	// example:
	//
	// 1.0
	PresencePenalty *float64 `json:"PresencePenalty,omitempty" xml:"PresencePenalty,omitempty"`
	// example:
	//
	// 42
	Seed *int64    `json:"Seed,omitempty" xml:"Seed,omitempty"`
	Stop []*string `json:"Stop,omitempty" xml:"Stop,omitempty" type:"Repeated"`
	// example:
	//
	// 0.6
	Temperature *float64                                        `json:"Temperature,omitempty" xml:"Temperature,omitempty"`
	Tools       []*ChatWithKnowledgeBaseRequestModelParamsTools `json:"Tools,omitempty" xml:"Tools,omitempty" type:"Repeated"`
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
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
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
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// get_weather
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
