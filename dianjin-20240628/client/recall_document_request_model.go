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
	// Metadata filter conditions.
	Filters []*RecallDocumentRequestFilters `json:"filters,omitempty" xml:"filters,omitempty" type:"Repeated"`
	// Text.
	//
	// This parameter is required.
	//
	// example:
	//
	// 欧洲杯历史上有哪些球队因为球员的适应新文化而受益
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
	// Enable parent-child document chunk retrieval.
	//
	// - Parent-child document chunks: During document parsing, a complete semantic block, such as a paragraph or a section, might split into multiple document chunks. This depends on your chunking strategy. When you enable parent-child document retrieval, the system attempts to complete the semantic block of the retrieved document chunk. This makes the corpus more semantically complete when constructing prompts, improving answer completeness and accuracy.
	//
	// example:
	//
	// false
	Rearrangement *bool `json:"rearrangement,omitempty" xml:"rearrangement,omitempty"`
	// The number of document chunks to retrieve.
	//
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
	if s.Filters != nil {
		for _, item := range s.Filters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RecallDocumentRequestFilters struct {
	// AND expression, used to filter documents/document chunks.
	And []*RecallDocumentRequestFiltersAnd `json:"and,omitempty" xml:"and,omitempty" type:"Repeated"`
	// Document chunk type, used to filter document chunks, such as: Text, Graph, Table, FAQ.
	//
	// example:
	//
	// Text
	ChunkType *string `json:"chunkType,omitempty" xml:"chunkType,omitempty"`
	// Document ID list, used to filter documents/document chunks.
	DocIdList []*string `json:"docIdList,omitempty" xml:"docIdList,omitempty" type:"Repeated"`
	// Document library ID, used to filter documents/document chunks.
	//
	// This parameter is required.
	//
	// example:
	//
	// sdbjhvs
	LibraryId *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
	// OR expression, used to filter documents/document chunks.
	Or []*RecallDocumentRequestFiltersOr `json:"or,omitempty" xml:"or,omitempty" type:"Repeated"`
	// Document status list, used to filter documents.
	Status []*string `json:"status,omitempty" xml:"status,omitempty" type:"Repeated"`
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

type RecallDocumentRequestFiltersAnd struct {
	// Keyword weight.
	//
	// example:
	//
	// 20
	Boost *float32 `json:"boost,omitempty" xml:"boost,omitempty"`
	// The key of the metadata in the document library.
	//
	// example:
	//
	// docType
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The relationship between the value stored in the document library metadata key and the value you enter.
	//
	// - eq: The value stored in the document library metadata key equals the value you enter.
	//
	// - lte: The value stored in the document library metadata key is less than or equal to the value you enter.
	//
	// - gte: The value stored in the document library metadata key is greater than or equal to the value you enter.
	//
	// - lt: The value stored in the document library metadata key is less than the value you enter.
	//
	// - gt: The value stored in the document library metadata key is greater than the value you enter.
	//
	// - contains: The list of values stored in the document library metadata key contains the value you enter.
	//
	// example:
	//
	// contains
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// The value of the metadata you enter.
	//
	// example:
	//
	// 策略报告
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
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
	// Keyword weight.
	//
	// example:
	//
	// 30
	Boost *float32 `json:"boost,omitempty" xml:"boost,omitempty"`
	// The key of the metadata in the document library.
	//
	// example:
	//
	// researcher
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The relationship between the value stored in the document library metadata key and the value you enter.
	//
	// - eq: The value stored in the document library metadata key equals the value you enter.
	//
	// - lte: The value stored in the document library metadata key is less than or equal to the value you enter.
	//
	// - gte: The value stored in the document library metadata key is greater than or equal to the value you enter.
	//
	// - lt: The value stored in the document library metadata key is less than the value you enter.
	//
	// - gt: The value stored in the document library metadata key is greater than the value you enter.
	//
	// - contains: The list of values stored in the document library metadata key contains the value you enter.
	//
	// example:
	//
	// eq
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// The value of the metadata you enter.
	//
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
