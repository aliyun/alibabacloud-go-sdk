// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetrieveKnowledgeBaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v string) *RetrieveKnowledgeBaseRequest
	GetFilter() *string
	SetHybridSearch(v string) *RetrieveKnowledgeBaseRequest
	GetHybridSearch() *string
	SetHybridSearchArgs(v string) *RetrieveKnowledgeBaseRequest
	GetHybridSearchArgs() *string
	SetIncludeMetadataFields(v string) *RetrieveKnowledgeBaseRequest
	GetIncludeMetadataFields() *string
	SetIncludeVector(v bool) *RetrieveKnowledgeBaseRequest
	GetIncludeVector() *bool
	SetKbUuid(v string) *RetrieveKnowledgeBaseRequest
	GetKbUuid() *string
	SetMetrics(v string) *RetrieveKnowledgeBaseRequest
	GetMetrics() *string
	SetOffset(v int32) *RetrieveKnowledgeBaseRequest
	GetOffset() *int32
	SetOrderBy(v string) *RetrieveKnowledgeBaseRequest
	GetOrderBy() *string
	SetQuery(v string) *RetrieveKnowledgeBaseRequest
	GetQuery() *string
	SetRecallWindow(v string) *RetrieveKnowledgeBaseRequest
	GetRecallWindow() *string
	SetRerankFactor(v float64) *RetrieveKnowledgeBaseRequest
	GetRerankFactor() *float64
	SetTopK(v int32) *RetrieveKnowledgeBaseRequest
	GetTopK() *int32
}

type RetrieveKnowledgeBaseRequest struct {
	// example:
	//
	// title = \\"test\\" AND name like \\"test%\\"
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// example:
	//
	// RRF
	HybridSearch *string `json:"HybridSearch,omitempty" xml:"HybridSearch,omitempty"`
	// example:
	//
	// { \\"Weight\\": { \\"alpha\\": 0.5 } }
	HybridSearchArgs *string `json:"HybridSearchArgs,omitempty" xml:"HybridSearchArgs,omitempty"`
	// example:
	//
	// title,page
	IncludeMetadataFields *string `json:"IncludeMetadataFields,omitempty" xml:"IncludeMetadataFields,omitempty"`
	// example:
	//
	// false
	IncludeVector *bool `json:"IncludeVector,omitempty" xml:"IncludeVector,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// kb-***
	KbUuid *string `json:"KbUuid,omitempty" xml:"KbUuid,omitempty"`
	// example:
	//
	// cosine
	Metrics *string `json:"Metrics,omitempty" xml:"Metrics,omitempty"`
	// example:
	//
	// 0
	Offset *int32 `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// example:
	//
	// created_at
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// What is GraphRAG?
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// example:
	//
	// [-5,5]
	RecallWindow *string `json:"RecallWindow,omitempty" xml:"RecallWindow,omitempty"`
	// example:
	//
	// 2
	RerankFactor *float64 `json:"RerankFactor,omitempty" xml:"RerankFactor,omitempty"`
	// example:
	//
	// 10
	TopK *int32 `json:"TopK,omitempty" xml:"TopK,omitempty"`
}

func (s RetrieveKnowledgeBaseRequest) String() string {
	return dara.Prettify(s)
}

func (s RetrieveKnowledgeBaseRequest) GoString() string {
	return s.String()
}

func (s *RetrieveKnowledgeBaseRequest) GetFilter() *string {
	return s.Filter
}

func (s *RetrieveKnowledgeBaseRequest) GetHybridSearch() *string {
	return s.HybridSearch
}

func (s *RetrieveKnowledgeBaseRequest) GetHybridSearchArgs() *string {
	return s.HybridSearchArgs
}

func (s *RetrieveKnowledgeBaseRequest) GetIncludeMetadataFields() *string {
	return s.IncludeMetadataFields
}

func (s *RetrieveKnowledgeBaseRequest) GetIncludeVector() *bool {
	return s.IncludeVector
}

func (s *RetrieveKnowledgeBaseRequest) GetKbUuid() *string {
	return s.KbUuid
}

func (s *RetrieveKnowledgeBaseRequest) GetMetrics() *string {
	return s.Metrics
}

func (s *RetrieveKnowledgeBaseRequest) GetOffset() *int32 {
	return s.Offset
}

func (s *RetrieveKnowledgeBaseRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *RetrieveKnowledgeBaseRequest) GetQuery() *string {
	return s.Query
}

func (s *RetrieveKnowledgeBaseRequest) GetRecallWindow() *string {
	return s.RecallWindow
}

func (s *RetrieveKnowledgeBaseRequest) GetRerankFactor() *float64 {
	return s.RerankFactor
}

func (s *RetrieveKnowledgeBaseRequest) GetTopK() *int32 {
	return s.TopK
}

func (s *RetrieveKnowledgeBaseRequest) SetFilter(v string) *RetrieveKnowledgeBaseRequest {
	s.Filter = &v
	return s
}

func (s *RetrieveKnowledgeBaseRequest) SetHybridSearch(v string) *RetrieveKnowledgeBaseRequest {
	s.HybridSearch = &v
	return s
}

func (s *RetrieveKnowledgeBaseRequest) SetHybridSearchArgs(v string) *RetrieveKnowledgeBaseRequest {
	s.HybridSearchArgs = &v
	return s
}

func (s *RetrieveKnowledgeBaseRequest) SetIncludeMetadataFields(v string) *RetrieveKnowledgeBaseRequest {
	s.IncludeMetadataFields = &v
	return s
}

func (s *RetrieveKnowledgeBaseRequest) SetIncludeVector(v bool) *RetrieveKnowledgeBaseRequest {
	s.IncludeVector = &v
	return s
}

func (s *RetrieveKnowledgeBaseRequest) SetKbUuid(v string) *RetrieveKnowledgeBaseRequest {
	s.KbUuid = &v
	return s
}

func (s *RetrieveKnowledgeBaseRequest) SetMetrics(v string) *RetrieveKnowledgeBaseRequest {
	s.Metrics = &v
	return s
}

func (s *RetrieveKnowledgeBaseRequest) SetOffset(v int32) *RetrieveKnowledgeBaseRequest {
	s.Offset = &v
	return s
}

func (s *RetrieveKnowledgeBaseRequest) SetOrderBy(v string) *RetrieveKnowledgeBaseRequest {
	s.OrderBy = &v
	return s
}

func (s *RetrieveKnowledgeBaseRequest) SetQuery(v string) *RetrieveKnowledgeBaseRequest {
	s.Query = &v
	return s
}

func (s *RetrieveKnowledgeBaseRequest) SetRecallWindow(v string) *RetrieveKnowledgeBaseRequest {
	s.RecallWindow = &v
	return s
}

func (s *RetrieveKnowledgeBaseRequest) SetRerankFactor(v float64) *RetrieveKnowledgeBaseRequest {
	s.RerankFactor = &v
	return s
}

func (s *RetrieveKnowledgeBaseRequest) SetTopK(v int32) *RetrieveKnowledgeBaseRequest {
	s.TopK = &v
	return s
}

func (s *RetrieveKnowledgeBaseRequest) Validate() error {
	return dara.Validate(s)
}
