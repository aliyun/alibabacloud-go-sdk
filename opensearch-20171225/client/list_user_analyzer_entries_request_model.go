// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserAnalyzerEntriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListUserAnalyzerEntriesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListUserAnalyzerEntriesRequest
	GetPageSize() *int32
	SetWord(v string) *ListUserAnalyzerEntriesRequest
	GetWord() *string
}

type ListUserAnalyzerEntriesRequest struct {
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The key to be used to query entries.
	//
	// example:
	//
	// kevintest
	Word *string `json:"word,omitempty" xml:"word,omitempty"`
}

func (s ListUserAnalyzerEntriesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUserAnalyzerEntriesRequest) GoString() string {
	return s.String()
}

func (s *ListUserAnalyzerEntriesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListUserAnalyzerEntriesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListUserAnalyzerEntriesRequest) GetWord() *string {
	return s.Word
}

func (s *ListUserAnalyzerEntriesRequest) SetPageNumber(v int32) *ListUserAnalyzerEntriesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListUserAnalyzerEntriesRequest) SetPageSize(v int32) *ListUserAnalyzerEntriesRequest {
	s.PageSize = &v
	return s
}

func (s *ListUserAnalyzerEntriesRequest) SetWord(v string) *ListUserAnalyzerEntriesRequest {
	s.Word = &v
	return s
}

func (s *ListUserAnalyzerEntriesRequest) Validate() error {
	return dara.Validate(s)
}
