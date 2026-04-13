// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChatsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrder(v string) *ListChatsRequest
	GetOrder() *string
	SetPageNumber(v int64) *ListChatsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListChatsRequest
	GetPageSize() *int64
	SetSortBy(v string) *ListChatsRequest
	GetSortBy() *string
	SetVerbose(v string) *ListChatsRequest
	GetVerbose() *string
}

type ListChatsRequest struct {
	// example:
	//
	// DESC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 60
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// GmtCreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// example:
	//
	// QuestionAndAnswer
	Verbose *string `json:"Verbose,omitempty" xml:"Verbose,omitempty"`
}

func (s ListChatsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListChatsRequest) GoString() string {
	return s.String()
}

func (s *ListChatsRequest) GetOrder() *string {
	return s.Order
}

func (s *ListChatsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListChatsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListChatsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListChatsRequest) GetVerbose() *string {
	return s.Verbose
}

func (s *ListChatsRequest) SetOrder(v string) *ListChatsRequest {
	s.Order = &v
	return s
}

func (s *ListChatsRequest) SetPageNumber(v int64) *ListChatsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListChatsRequest) SetPageSize(v int64) *ListChatsRequest {
	s.PageSize = &v
	return s
}

func (s *ListChatsRequest) SetSortBy(v string) *ListChatsRequest {
	s.SortBy = &v
	return s
}

func (s *ListChatsRequest) SetVerbose(v string) *ListChatsRequest {
	s.Verbose = &v
	return s
}

func (s *ListChatsRequest) Validate() error {
	return dara.Validate(s)
}
