// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInterventionDictionaryEntriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListInterventionDictionaryEntriesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListInterventionDictionaryEntriesRequest
	GetPageSize() *int32
	SetWord(v string) *ListInterventionDictionaryEntriesRequest
	GetWord() *string
}

type ListInterventionDictionaryEntriesRequest struct {
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries returned per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The intervention entry.
	//
	// example:
	//
	// test
	Word *string `json:"word,omitempty" xml:"word,omitempty"`
}

func (s ListInterventionDictionaryEntriesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInterventionDictionaryEntriesRequest) GoString() string {
	return s.String()
}

func (s *ListInterventionDictionaryEntriesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListInterventionDictionaryEntriesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListInterventionDictionaryEntriesRequest) GetWord() *string {
	return s.Word
}

func (s *ListInterventionDictionaryEntriesRequest) SetPageNumber(v int32) *ListInterventionDictionaryEntriesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInterventionDictionaryEntriesRequest) SetPageSize(v int32) *ListInterventionDictionaryEntriesRequest {
	s.PageSize = &v
	return s
}

func (s *ListInterventionDictionaryEntriesRequest) SetWord(v string) *ListInterventionDictionaryEntriesRequest {
	s.Word = &v
	return s
}

func (s *ListInterventionDictionaryEntriesRequest) Validate() error {
	return dara.Validate(s)
}
