// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAliasesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListAliasesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAliasesRequest
	GetPageSize() *int32
}

type ListAliasesRequest struct {
	// The number of the page to return.
	//
	// Pages start from page 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: 0 to 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListAliasesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAliasesRequest) GoString() string {
	return s.String()
}

func (s *ListAliasesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAliasesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAliasesRequest) SetPageNumber(v int32) *ListAliasesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAliasesRequest) SetPageSize(v int32) *ListAliasesRequest {
	s.PageSize = &v
	return s
}

func (s *ListAliasesRequest) Validate() error {
	return dara.Validate(s)
}
