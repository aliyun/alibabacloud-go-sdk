// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRoutineCodeVersionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *ListRoutineCodeVersionsRequest
	GetName() *string
	SetPageNumber(v int64) *ListRoutineCodeVersionsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListRoutineCodeVersionsRequest
	GetPageSize() *int64
	SetSearchKeyWord(v string) *ListRoutineCodeVersionsRequest
	GetSearchKeyWord() *string
}

type ListRoutineCodeVersionsRequest struct {
	// The function name.
	//
	// This parameter is required.
	//
	// example:
	//
	// ListRoutineCodeVersions
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The page number of the returned page. Default value: 1. Valid values: 1 and 2.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 20.
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

func (s ListRoutineCodeVersionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRoutineCodeVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListRoutineCodeVersionsRequest) GetName() *string {
	return s.Name
}

func (s *ListRoutineCodeVersionsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListRoutineCodeVersionsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListRoutineCodeVersionsRequest) GetSearchKeyWord() *string {
	return s.SearchKeyWord
}

func (s *ListRoutineCodeVersionsRequest) SetName(v string) *ListRoutineCodeVersionsRequest {
	s.Name = &v
	return s
}

func (s *ListRoutineCodeVersionsRequest) SetPageNumber(v int64) *ListRoutineCodeVersionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRoutineCodeVersionsRequest) SetPageSize(v int64) *ListRoutineCodeVersionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListRoutineCodeVersionsRequest) SetSearchKeyWord(v string) *ListRoutineCodeVersionsRequest {
	s.SearchKeyWord = &v
	return s
}

func (s *ListRoutineCodeVersionsRequest) Validate() error {
	return dara.Validate(s)
}
