// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRoutineRelatedRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *ListRoutineRelatedRecordsRequest
	GetName() *string
	SetPageNumber(v int64) *ListRoutineRelatedRecordsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListRoutineRelatedRecordsRequest
	GetPageSize() *int64
	SetSearchKeyWord(v string) *ListRoutineRelatedRecordsRequest
	GetSearchKeyWord() *string
}

type ListRoutineRelatedRecordsRequest struct {
	// The name of the function.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The page number of the returned page. Default value: 1
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: an integer from 1 to 20.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The keyword used for fuzzy search.
	//
	// example:
	//
	// hello
	SearchKeyWord *string `json:"SearchKeyWord,omitempty" xml:"SearchKeyWord,omitempty"`
}

func (s ListRoutineRelatedRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRoutineRelatedRecordsRequest) GoString() string {
	return s.String()
}

func (s *ListRoutineRelatedRecordsRequest) GetName() *string {
	return s.Name
}

func (s *ListRoutineRelatedRecordsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListRoutineRelatedRecordsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListRoutineRelatedRecordsRequest) GetSearchKeyWord() *string {
	return s.SearchKeyWord
}

func (s *ListRoutineRelatedRecordsRequest) SetName(v string) *ListRoutineRelatedRecordsRequest {
	s.Name = &v
	return s
}

func (s *ListRoutineRelatedRecordsRequest) SetPageNumber(v int64) *ListRoutineRelatedRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRoutineRelatedRecordsRequest) SetPageSize(v int64) *ListRoutineRelatedRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *ListRoutineRelatedRecordsRequest) SetSearchKeyWord(v string) *ListRoutineRelatedRecordsRequest {
	s.SearchKeyWord = &v
	return s
}

func (s *ListRoutineRelatedRecordsRequest) Validate() error {
	return dara.Validate(s)
}
