// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunLibraryChatGenerationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDocIdList(v []*string) *RunLibraryChatGenerationRequest
	GetDocIdList() []*string
	SetEnableFollowUp(v bool) *RunLibraryChatGenerationRequest
	GetEnableFollowUp() *bool
	SetEnableMultiQuery(v bool) *RunLibraryChatGenerationRequest
	GetEnableMultiQuery() *bool
	SetEnableOpenQa(v bool) *RunLibraryChatGenerationRequest
	GetEnableOpenQa() *bool
	SetFollowUpLlm(v string) *RunLibraryChatGenerationRequest
	GetFollowUpLlm() *string
	SetLibraryId(v string) *RunLibraryChatGenerationRequest
	GetLibraryId() *string
	SetLlmType(v string) *RunLibraryChatGenerationRequest
	GetLlmType() *string
	SetMultiQueryLlm(v string) *RunLibraryChatGenerationRequest
	GetMultiQueryLlm() *string
	SetQuery(v string) *RunLibraryChatGenerationRequest
	GetQuery() *string
	SetQueryCriteria(v *RunLibraryChatGenerationRequestQueryCriteria) *RunLibraryChatGenerationRequest
	GetQueryCriteria() *RunLibraryChatGenerationRequestQueryCriteria
	SetRerankType(v string) *RunLibraryChatGenerationRequest
	GetRerankType() *string
	SetSessionId(v string) *RunLibraryChatGenerationRequest
	GetSessionId() *string
	SetStream(v bool) *RunLibraryChatGenerationRequest
	GetStream() *bool
	SetSubQueryList(v []*string) *RunLibraryChatGenerationRequest
	GetSubQueryList() []*string
	SetTextSearchParameter(v *RunLibraryChatGenerationRequestTextSearchParameter) *RunLibraryChatGenerationRequest
	GetTextSearchParameter() *RunLibraryChatGenerationRequestTextSearchParameter
	SetTopK(v int32) *RunLibraryChatGenerationRequest
	GetTopK() *int32
	SetVectorSearchParameter(v *RunLibraryChatGenerationRequestVectorSearchParameter) *RunLibraryChatGenerationRequest
	GetVectorSearchParameter() *RunLibraryChatGenerationRequestVectorSearchParameter
	SetWithDocumentReference(v bool) *RunLibraryChatGenerationRequest
	GetWithDocumentReference() *bool
}

type RunLibraryChatGenerationRequest struct {
	DocIdList []*string `json:"docIdList,omitempty" xml:"docIdList,omitempty" type:"Repeated"`
	// example:
	//
	// false
	EnableFollowUp *bool `json:"enableFollowUp,omitempty" xml:"enableFollowUp,omitempty"`
	// example:
	//
	// false
	EnableMultiQuery *bool `json:"enableMultiQuery,omitempty" xml:"enableMultiQuery,omitempty"`
	// example:
	//
	// false
	EnableOpenQa *bool `json:"enableOpenQa,omitempty" xml:"enableOpenQa,omitempty"`
	// example:
	//
	// qwen-max
	FollowUpLlm *string `json:"followUpLlm,omitempty" xml:"followUpLlm,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3akzl28vap
	LibraryId *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// qwen-max
	LlmType *string `json:"llmType,omitempty" xml:"llmType,omitempty"`
	// example:
	//
	// qwen-max
	MultiQueryLlm *string `json:"multiQueryLlm,omitempty" xml:"multiQueryLlm,omitempty"`
	// This parameter is required.
	Query         *string                                       `json:"query,omitempty" xml:"query,omitempty"`
	QueryCriteria *RunLibraryChatGenerationRequestQueryCriteria `json:"queryCriteria,omitempty" xml:"queryCriteria,omitempty" type:"Struct"`
	// example:
	//
	// linear
	RerankType *string `json:"rerankType,omitempty" xml:"rerankType,omitempty"`
	// sessionId
	//
	// example:
	//
	// null
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// example:
	//
	// false
	Stream              *bool                                               `json:"stream,omitempty" xml:"stream,omitempty"`
	SubQueryList        []*string                                           `json:"subQueryList,omitempty" xml:"subQueryList,omitempty" type:"Repeated"`
	TextSearchParameter *RunLibraryChatGenerationRequestTextSearchParameter `json:"textSearchParameter,omitempty" xml:"textSearchParameter,omitempty" type:"Struct"`
	// example:
	//
	// 1
	TopK                  *int32                                                `json:"topK,omitempty" xml:"topK,omitempty"`
	VectorSearchParameter *RunLibraryChatGenerationRequestVectorSearchParameter `json:"vectorSearchParameter,omitempty" xml:"vectorSearchParameter,omitempty" type:"Struct"`
	// example:
	//
	// false
	WithDocumentReference *bool `json:"withDocumentReference,omitempty" xml:"withDocumentReference,omitempty"`
}

func (s RunLibraryChatGenerationRequest) String() string {
	return dara.Prettify(s)
}

func (s RunLibraryChatGenerationRequest) GoString() string {
	return s.String()
}

func (s *RunLibraryChatGenerationRequest) GetDocIdList() []*string {
	return s.DocIdList
}

func (s *RunLibraryChatGenerationRequest) GetEnableFollowUp() *bool {
	return s.EnableFollowUp
}

func (s *RunLibraryChatGenerationRequest) GetEnableMultiQuery() *bool {
	return s.EnableMultiQuery
}

func (s *RunLibraryChatGenerationRequest) GetEnableOpenQa() *bool {
	return s.EnableOpenQa
}

func (s *RunLibraryChatGenerationRequest) GetFollowUpLlm() *string {
	return s.FollowUpLlm
}

func (s *RunLibraryChatGenerationRequest) GetLibraryId() *string {
	return s.LibraryId
}

func (s *RunLibraryChatGenerationRequest) GetLlmType() *string {
	return s.LlmType
}

func (s *RunLibraryChatGenerationRequest) GetMultiQueryLlm() *string {
	return s.MultiQueryLlm
}

func (s *RunLibraryChatGenerationRequest) GetQuery() *string {
	return s.Query
}

func (s *RunLibraryChatGenerationRequest) GetQueryCriteria() *RunLibraryChatGenerationRequestQueryCriteria {
	return s.QueryCriteria
}

func (s *RunLibraryChatGenerationRequest) GetRerankType() *string {
	return s.RerankType
}

func (s *RunLibraryChatGenerationRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *RunLibraryChatGenerationRequest) GetStream() *bool {
	return s.Stream
}

func (s *RunLibraryChatGenerationRequest) GetSubQueryList() []*string {
	return s.SubQueryList
}

func (s *RunLibraryChatGenerationRequest) GetTextSearchParameter() *RunLibraryChatGenerationRequestTextSearchParameter {
	return s.TextSearchParameter
}

func (s *RunLibraryChatGenerationRequest) GetTopK() *int32 {
	return s.TopK
}

func (s *RunLibraryChatGenerationRequest) GetVectorSearchParameter() *RunLibraryChatGenerationRequestVectorSearchParameter {
	return s.VectorSearchParameter
}

func (s *RunLibraryChatGenerationRequest) GetWithDocumentReference() *bool {
	return s.WithDocumentReference
}

func (s *RunLibraryChatGenerationRequest) SetDocIdList(v []*string) *RunLibraryChatGenerationRequest {
	s.DocIdList = v
	return s
}

func (s *RunLibraryChatGenerationRequest) SetEnableFollowUp(v bool) *RunLibraryChatGenerationRequest {
	s.EnableFollowUp = &v
	return s
}

func (s *RunLibraryChatGenerationRequest) SetEnableMultiQuery(v bool) *RunLibraryChatGenerationRequest {
	s.EnableMultiQuery = &v
	return s
}

func (s *RunLibraryChatGenerationRequest) SetEnableOpenQa(v bool) *RunLibraryChatGenerationRequest {
	s.EnableOpenQa = &v
	return s
}

func (s *RunLibraryChatGenerationRequest) SetFollowUpLlm(v string) *RunLibraryChatGenerationRequest {
	s.FollowUpLlm = &v
	return s
}

func (s *RunLibraryChatGenerationRequest) SetLibraryId(v string) *RunLibraryChatGenerationRequest {
	s.LibraryId = &v
	return s
}

func (s *RunLibraryChatGenerationRequest) SetLlmType(v string) *RunLibraryChatGenerationRequest {
	s.LlmType = &v
	return s
}

func (s *RunLibraryChatGenerationRequest) SetMultiQueryLlm(v string) *RunLibraryChatGenerationRequest {
	s.MultiQueryLlm = &v
	return s
}

func (s *RunLibraryChatGenerationRequest) SetQuery(v string) *RunLibraryChatGenerationRequest {
	s.Query = &v
	return s
}

func (s *RunLibraryChatGenerationRequest) SetQueryCriteria(v *RunLibraryChatGenerationRequestQueryCriteria) *RunLibraryChatGenerationRequest {
	s.QueryCriteria = v
	return s
}

func (s *RunLibraryChatGenerationRequest) SetRerankType(v string) *RunLibraryChatGenerationRequest {
	s.RerankType = &v
	return s
}

func (s *RunLibraryChatGenerationRequest) SetSessionId(v string) *RunLibraryChatGenerationRequest {
	s.SessionId = &v
	return s
}

func (s *RunLibraryChatGenerationRequest) SetStream(v bool) *RunLibraryChatGenerationRequest {
	s.Stream = &v
	return s
}

func (s *RunLibraryChatGenerationRequest) SetSubQueryList(v []*string) *RunLibraryChatGenerationRequest {
	s.SubQueryList = v
	return s
}

func (s *RunLibraryChatGenerationRequest) SetTextSearchParameter(v *RunLibraryChatGenerationRequestTextSearchParameter) *RunLibraryChatGenerationRequest {
	s.TextSearchParameter = v
	return s
}

func (s *RunLibraryChatGenerationRequest) SetTopK(v int32) *RunLibraryChatGenerationRequest {
	s.TopK = &v
	return s
}

func (s *RunLibraryChatGenerationRequest) SetVectorSearchParameter(v *RunLibraryChatGenerationRequestVectorSearchParameter) *RunLibraryChatGenerationRequest {
	s.VectorSearchParameter = v
	return s
}

func (s *RunLibraryChatGenerationRequest) SetWithDocumentReference(v bool) *RunLibraryChatGenerationRequest {
	s.WithDocumentReference = &v
	return s
}

func (s *RunLibraryChatGenerationRequest) Validate() error {
	if s.QueryCriteria != nil {
		if err := s.QueryCriteria.Validate(); err != nil {
			return err
		}
	}
	if s.TextSearchParameter != nil {
		if err := s.TextSearchParameter.Validate(); err != nil {
			return err
		}
	}
	if s.VectorSearchParameter != nil {
		if err := s.VectorSearchParameter.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunLibraryChatGenerationRequestQueryCriteria struct {
	And []*RunLibraryChatGenerationRequestQueryCriteriaAnd `json:"and,omitempty" xml:"and,omitempty" type:"Repeated"`
	Or  []*RunLibraryChatGenerationRequestQueryCriteriaOr  `json:"or,omitempty" xml:"or,omitempty" type:"Repeated"`
}

func (s RunLibraryChatGenerationRequestQueryCriteria) String() string {
	return dara.Prettify(s)
}

func (s RunLibraryChatGenerationRequestQueryCriteria) GoString() string {
	return s.String()
}

func (s *RunLibraryChatGenerationRequestQueryCriteria) GetAnd() []*RunLibraryChatGenerationRequestQueryCriteriaAnd {
	return s.And
}

func (s *RunLibraryChatGenerationRequestQueryCriteria) GetOr() []*RunLibraryChatGenerationRequestQueryCriteriaOr {
	return s.Or
}

func (s *RunLibraryChatGenerationRequestQueryCriteria) SetAnd(v []*RunLibraryChatGenerationRequestQueryCriteriaAnd) *RunLibraryChatGenerationRequestQueryCriteria {
	s.And = v
	return s
}

func (s *RunLibraryChatGenerationRequestQueryCriteria) SetOr(v []*RunLibraryChatGenerationRequestQueryCriteriaOr) *RunLibraryChatGenerationRequestQueryCriteria {
	s.Or = v
	return s
}

func (s *RunLibraryChatGenerationRequestQueryCriteria) Validate() error {
	if s.And != nil {
		for _, item := range s.And {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Or != nil {
		for _, item := range s.Or {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RunLibraryChatGenerationRequestQueryCriteriaAnd struct {
	// example:
	//
	// 0.5
	Boost *float32 `json:"boost,omitempty" xml:"boost,omitempty"`
	// example:
	//
	// city
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// eq
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	Value    *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s RunLibraryChatGenerationRequestQueryCriteriaAnd) String() string {
	return dara.Prettify(s)
}

func (s RunLibraryChatGenerationRequestQueryCriteriaAnd) GoString() string {
	return s.String()
}

func (s *RunLibraryChatGenerationRequestQueryCriteriaAnd) GetBoost() *float32 {
	return s.Boost
}

func (s *RunLibraryChatGenerationRequestQueryCriteriaAnd) GetKey() *string {
	return s.Key
}

func (s *RunLibraryChatGenerationRequestQueryCriteriaAnd) GetOperator() *string {
	return s.Operator
}

func (s *RunLibraryChatGenerationRequestQueryCriteriaAnd) GetValue() *string {
	return s.Value
}

func (s *RunLibraryChatGenerationRequestQueryCriteriaAnd) SetBoost(v float32) *RunLibraryChatGenerationRequestQueryCriteriaAnd {
	s.Boost = &v
	return s
}

func (s *RunLibraryChatGenerationRequestQueryCriteriaAnd) SetKey(v string) *RunLibraryChatGenerationRequestQueryCriteriaAnd {
	s.Key = &v
	return s
}

func (s *RunLibraryChatGenerationRequestQueryCriteriaAnd) SetOperator(v string) *RunLibraryChatGenerationRequestQueryCriteriaAnd {
	s.Operator = &v
	return s
}

func (s *RunLibraryChatGenerationRequestQueryCriteriaAnd) SetValue(v string) *RunLibraryChatGenerationRequestQueryCriteriaAnd {
	s.Value = &v
	return s
}

func (s *RunLibraryChatGenerationRequestQueryCriteriaAnd) Validate() error {
	return dara.Validate(s)
}

type RunLibraryChatGenerationRequestQueryCriteriaOr struct {
	// example:
	//
	// 0.5
	Boost *float32 `json:"boost,omitempty" xml:"boost,omitempty"`
	// example:
	//
	// city
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// eq
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	Value    *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s RunLibraryChatGenerationRequestQueryCriteriaOr) String() string {
	return dara.Prettify(s)
}

func (s RunLibraryChatGenerationRequestQueryCriteriaOr) GoString() string {
	return s.String()
}

func (s *RunLibraryChatGenerationRequestQueryCriteriaOr) GetBoost() *float32 {
	return s.Boost
}

func (s *RunLibraryChatGenerationRequestQueryCriteriaOr) GetKey() *string {
	return s.Key
}

func (s *RunLibraryChatGenerationRequestQueryCriteriaOr) GetOperator() *string {
	return s.Operator
}

func (s *RunLibraryChatGenerationRequestQueryCriteriaOr) GetValue() *string {
	return s.Value
}

func (s *RunLibraryChatGenerationRequestQueryCriteriaOr) SetBoost(v float32) *RunLibraryChatGenerationRequestQueryCriteriaOr {
	s.Boost = &v
	return s
}

func (s *RunLibraryChatGenerationRequestQueryCriteriaOr) SetKey(v string) *RunLibraryChatGenerationRequestQueryCriteriaOr {
	s.Key = &v
	return s
}

func (s *RunLibraryChatGenerationRequestQueryCriteriaOr) SetOperator(v string) *RunLibraryChatGenerationRequestQueryCriteriaOr {
	s.Operator = &v
	return s
}

func (s *RunLibraryChatGenerationRequestQueryCriteriaOr) SetValue(v string) *RunLibraryChatGenerationRequestQueryCriteriaOr {
	s.Value = &v
	return s
}

func (s *RunLibraryChatGenerationRequestQueryCriteriaOr) Validate() error {
	return dara.Validate(s)
}

type RunLibraryChatGenerationRequestTextSearchParameter struct {
	// example:
	//
	// 10
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// example:
	//
	// IkMaxWord
	SearchAnalyzerType *string `json:"searchAnalyzerType,omitempty" xml:"searchAnalyzerType,omitempty"`
}

func (s RunLibraryChatGenerationRequestTextSearchParameter) String() string {
	return dara.Prettify(s)
}

func (s RunLibraryChatGenerationRequestTextSearchParameter) GoString() string {
	return s.String()
}

func (s *RunLibraryChatGenerationRequestTextSearchParameter) GetLimit() *int32 {
	return s.Limit
}

func (s *RunLibraryChatGenerationRequestTextSearchParameter) GetSearchAnalyzerType() *string {
	return s.SearchAnalyzerType
}

func (s *RunLibraryChatGenerationRequestTextSearchParameter) SetLimit(v int32) *RunLibraryChatGenerationRequestTextSearchParameter {
	s.Limit = &v
	return s
}

func (s *RunLibraryChatGenerationRequestTextSearchParameter) SetSearchAnalyzerType(v string) *RunLibraryChatGenerationRequestTextSearchParameter {
	s.SearchAnalyzerType = &v
	return s
}

func (s *RunLibraryChatGenerationRequestTextSearchParameter) Validate() error {
	return dara.Validate(s)
}

type RunLibraryChatGenerationRequestVectorSearchParameter struct {
	// example:
	//
	// 10
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
}

func (s RunLibraryChatGenerationRequestVectorSearchParameter) String() string {
	return dara.Prettify(s)
}

func (s RunLibraryChatGenerationRequestVectorSearchParameter) GoString() string {
	return s.String()
}

func (s *RunLibraryChatGenerationRequestVectorSearchParameter) GetLimit() *int32 {
	return s.Limit
}

func (s *RunLibraryChatGenerationRequestVectorSearchParameter) SetLimit(v int32) *RunLibraryChatGenerationRequestVectorSearchParameter {
	s.Limit = &v
	return s
}

func (s *RunLibraryChatGenerationRequestVectorSearchParameter) Validate() error {
	return dara.Validate(s)
}
