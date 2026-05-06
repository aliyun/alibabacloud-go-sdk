// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVocabularyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessUnitId(v string) *ListVocabularyRequest
	GetBusinessUnitId() *string
	SetName(v string) *ListVocabularyRequest
	GetName() *string
	SetPageNumber(v int32) *ListVocabularyRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListVocabularyRequest
	GetPageSize() *int32
}

type ListVocabularyRequest struct {
	// example:
	//
	// llm-c11iig67g863rih8
	BusinessUnitId *string `json:"BusinessUnitId,omitempty" xml:"BusinessUnitId,omitempty"`
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListVocabularyRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVocabularyRequest) GoString() string {
	return s.String()
}

func (s *ListVocabularyRequest) GetBusinessUnitId() *string {
	return s.BusinessUnitId
}

func (s *ListVocabularyRequest) GetName() *string {
	return s.Name
}

func (s *ListVocabularyRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListVocabularyRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListVocabularyRequest) SetBusinessUnitId(v string) *ListVocabularyRequest {
	s.BusinessUnitId = &v
	return s
}

func (s *ListVocabularyRequest) SetName(v string) *ListVocabularyRequest {
	s.Name = &v
	return s
}

func (s *ListVocabularyRequest) SetPageNumber(v int32) *ListVocabularyRequest {
	s.PageNumber = &v
	return s
}

func (s *ListVocabularyRequest) SetPageSize(v int32) *ListVocabularyRequest {
	s.PageSize = &v
	return s
}

func (s *ListVocabularyRequest) Validate() error {
	return dara.Validate(s)
}
