// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserRoutinesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int64) *ListUserRoutinesRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListUserRoutinesRequest
	GetPageSize() *int64
	SetSearchKeyWord(v string) *ListUserRoutinesRequest
	GetSearchKeyWord() *string
}

type ListUserRoutinesRequest struct {
	// example:
	//
	// ListUserRoutines
	PageNumber    *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchKeyWord *string `json:"SearchKeyWord,omitempty" xml:"SearchKeyWord,omitempty"`
}

func (s ListUserRoutinesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUserRoutinesRequest) GoString() string {
	return s.String()
}

func (s *ListUserRoutinesRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListUserRoutinesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListUserRoutinesRequest) GetSearchKeyWord() *string {
	return s.SearchKeyWord
}

func (s *ListUserRoutinesRequest) SetPageNumber(v int64) *ListUserRoutinesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListUserRoutinesRequest) SetPageSize(v int64) *ListUserRoutinesRequest {
	s.PageSize = &v
	return s
}

func (s *ListUserRoutinesRequest) SetSearchKeyWord(v string) *ListUserRoutinesRequest {
	s.SearchKeyWord = &v
	return s
}

func (s *ListUserRoutinesRequest) Validate() error {
	return dara.Validate(s)
}
