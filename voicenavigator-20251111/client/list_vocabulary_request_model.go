// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVocabularyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListVocabularyRequest
	GetInstanceId() *string
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
	// af81a389-91f0-4157-8d82-720edd02b66a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
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

func (s *ListVocabularyRequest) GetInstanceId() *string {
	return s.InstanceId
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

func (s *ListVocabularyRequest) SetInstanceId(v string) *ListVocabularyRequest {
	s.InstanceId = &v
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
