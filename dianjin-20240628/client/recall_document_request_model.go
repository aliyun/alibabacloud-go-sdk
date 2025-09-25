// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecallDocumentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilters(v []*RecallDocumentRequestFilters) *RecallDocumentRequest
	GetFilters() []*RecallDocumentRequestFilters
	SetQuery(v string) *RecallDocumentRequest
	GetQuery() *string
	SetRearrangement(v bool) *RecallDocumentRequest
	GetRearrangement() *bool
	SetTopK(v int32) *RecallDocumentRequest
	GetTopK() *int32
}

type RecallDocumentRequest struct {
	Filters []*RecallDocumentRequestFilters `json:"filters,omitempty" xml:"filters,omitempty" type:"Repeated"`
	// This parameter is required.
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
	// example:
	//
	// false
	Rearrangement *bool `json:"rearrangement,omitempty" xml:"rearrangement,omitempty"`
	// example:
	//
	// 10
	TopK *int32 `json:"topK,omitempty" xml:"topK,omitempty"`
}

func (s RecallDocumentRequest) String() string {
	return dara.Prettify(s)
}

func (s RecallDocumentRequest) GoString() string {
	return s.String()
}

func (s *RecallDocumentRequest) GetFilters() []*RecallDocumentRequestFilters {
	return s.Filters
}

func (s *RecallDocumentRequest) GetQuery() *string {
	return s.Query
}

func (s *RecallDocumentRequest) GetRearrangement() *bool {
	return s.Rearrangement
}

func (s *RecallDocumentRequest) GetTopK() *int32 {
	return s.TopK
}

func (s *RecallDocumentRequest) SetFilters(v []*RecallDocumentRequestFilters) *RecallDocumentRequest {
	s.Filters = v
	return s
}

func (s *RecallDocumentRequest) SetQuery(v string) *RecallDocumentRequest {
	s.Query = &v
	return s
}

func (s *RecallDocumentRequest) SetRearrangement(v bool) *RecallDocumentRequest {
	s.Rearrangement = &v
	return s
}

func (s *RecallDocumentRequest) SetTopK(v int32) *RecallDocumentRequest {
	s.TopK = &v
	return s
}

func (s *RecallDocumentRequest) Validate() error {
	return dara.Validate(s)
}

type RecallDocumentRequestFilters struct {
	And []*RecallDocumentRequestFiltersAnd `json:"and,omitempty" xml:"and,omitempty" type:"Repeated"`
	// example:
	//
	// Text
	ChunkType *string   `json:"chunkType,omitempty" xml:"chunkType,omitempty"`
	DocIdList []*string `json:"docIdList,omitempty" xml:"docIdList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// sdbjhvs
	LibraryId *string                           `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
	Or        []*RecallDocumentRequestFiltersOr `json:"or,omitempty" xml:"or,omitempty" type:"Repeated"`
	Status    []*string                         `json:"status,omitempty" xml:"status,omitempty" type:"Repeated"`
}

func (s RecallDocumentRequestFilters) String() string {
	return dara.Prettify(s)
}

func (s RecallDocumentRequestFilters) GoString() string {
	return s.String()
}

func (s *RecallDocumentRequestFilters) GetAnd() []*RecallDocumentRequestFiltersAnd {
	return s.And
}

func (s *RecallDocumentRequestFilters) GetChunkType() *string {
	return s.ChunkType
}

func (s *RecallDocumentRequestFilters) GetDocIdList() []*string {
	return s.DocIdList
}

func (s *RecallDocumentRequestFilters) GetLibraryId() *string {
	return s.LibraryId
}

func (s *RecallDocumentRequestFilters) GetOr() []*RecallDocumentRequestFiltersOr {
	return s.Or
}

func (s *RecallDocumentRequestFilters) GetStatus() []*string {
	return s.Status
}

func (s *RecallDocumentRequestFilters) SetAnd(v []*RecallDocumentRequestFiltersAnd) *RecallDocumentRequestFilters {
	s.And = v
	return s
}

func (s *RecallDocumentRequestFilters) SetChunkType(v string) *RecallDocumentRequestFilters {
	s.ChunkType = &v
	return s
}

func (s *RecallDocumentRequestFilters) SetDocIdList(v []*string) *RecallDocumentRequestFilters {
	s.DocIdList = v
	return s
}

func (s *RecallDocumentRequestFilters) SetLibraryId(v string) *RecallDocumentRequestFilters {
	s.LibraryId = &v
	return s
}

func (s *RecallDocumentRequestFilters) SetOr(v []*RecallDocumentRequestFiltersOr) *RecallDocumentRequestFilters {
	s.Or = v
	return s
}

func (s *RecallDocumentRequestFilters) SetStatus(v []*string) *RecallDocumentRequestFilters {
	s.Status = v
	return s
}

func (s *RecallDocumentRequestFilters) Validate() error {
	return dara.Validate(s)
}

type RecallDocumentRequestFiltersAnd struct {
	// example:
	//
	// 20
	Boost *float32 `json:"boost,omitempty" xml:"boost,omitempty"`
	// example:
	//
	// docType
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// contains
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	Value    *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s RecallDocumentRequestFiltersAnd) String() string {
	return dara.Prettify(s)
}

func (s RecallDocumentRequestFiltersAnd) GoString() string {
	return s.String()
}

func (s *RecallDocumentRequestFiltersAnd) GetBoost() *float32 {
	return s.Boost
}

func (s *RecallDocumentRequestFiltersAnd) GetKey() *string {
	return s.Key
}

func (s *RecallDocumentRequestFiltersAnd) GetOperator() *string {
	return s.Operator
}

func (s *RecallDocumentRequestFiltersAnd) GetValue() *string {
	return s.Value
}

func (s *RecallDocumentRequestFiltersAnd) SetBoost(v float32) *RecallDocumentRequestFiltersAnd {
	s.Boost = &v
	return s
}

func (s *RecallDocumentRequestFiltersAnd) SetKey(v string) *RecallDocumentRequestFiltersAnd {
	s.Key = &v
	return s
}

func (s *RecallDocumentRequestFiltersAnd) SetOperator(v string) *RecallDocumentRequestFiltersAnd {
	s.Operator = &v
	return s
}

func (s *RecallDocumentRequestFiltersAnd) SetValue(v string) *RecallDocumentRequestFiltersAnd {
	s.Value = &v
	return s
}

func (s *RecallDocumentRequestFiltersAnd) Validate() error {
	return dara.Validate(s)
}

type RecallDocumentRequestFiltersOr struct {
	// example:
	//
	// 30
	Boost *float32 `json:"boost,omitempty" xml:"boost,omitempty"`
	// example:
	//
	// researcher
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// eq
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// example:
	//
	// zhangsan
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s RecallDocumentRequestFiltersOr) String() string {
	return dara.Prettify(s)
}

func (s RecallDocumentRequestFiltersOr) GoString() string {
	return s.String()
}

func (s *RecallDocumentRequestFiltersOr) GetBoost() *float32 {
	return s.Boost
}

func (s *RecallDocumentRequestFiltersOr) GetKey() *string {
	return s.Key
}

func (s *RecallDocumentRequestFiltersOr) GetOperator() *string {
	return s.Operator
}

func (s *RecallDocumentRequestFiltersOr) GetValue() *string {
	return s.Value
}

func (s *RecallDocumentRequestFiltersOr) SetBoost(v float32) *RecallDocumentRequestFiltersOr {
	s.Boost = &v
	return s
}

func (s *RecallDocumentRequestFiltersOr) SetKey(v string) *RecallDocumentRequestFiltersOr {
	s.Key = &v
	return s
}

func (s *RecallDocumentRequestFiltersOr) SetOperator(v string) *RecallDocumentRequestFiltersOr {
	s.Operator = &v
	return s
}

func (s *RecallDocumentRequestFiltersOr) SetValue(v string) *RecallDocumentRequestFiltersOr {
	s.Value = &v
	return s
}

func (s *RecallDocumentRequestFiltersOr) Validate() error {
	return dara.Validate(s)
}
