// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthorizedAccountsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListAuthorizedAccountsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAuthorizedAccountsRequest
	GetPageSize() *int32
}

type ListAuthorizedAccountsRequest struct {
	// The page number.
	//
	// Page starts from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 5
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListAuthorizedAccountsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizedAccountsRequest) GoString() string {
	return s.String()
}

func (s *ListAuthorizedAccountsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAuthorizedAccountsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAuthorizedAccountsRequest) SetPageNumber(v int32) *ListAuthorizedAccountsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAuthorizedAccountsRequest) SetPageSize(v int32) *ListAuthorizedAccountsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAuthorizedAccountsRequest) Validate() error {
	return dara.Validate(s)
}
