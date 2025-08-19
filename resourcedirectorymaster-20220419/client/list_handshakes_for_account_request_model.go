// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHandshakesForAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListHandshakesForAccountRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListHandshakesForAccountRequest
	GetPageSize() *int32
}

type ListHandshakesForAccountRequest struct {
	// The number of the page to return.
	//
	// Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListHandshakesForAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHandshakesForAccountRequest) GoString() string {
	return s.String()
}

func (s *ListHandshakesForAccountRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListHandshakesForAccountRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListHandshakesForAccountRequest) SetPageNumber(v int32) *ListHandshakesForAccountRequest {
	s.PageNumber = &v
	return s
}

func (s *ListHandshakesForAccountRequest) SetPageSize(v int32) *ListHandshakesForAccountRequest {
	s.PageSize = &v
	return s
}

func (s *ListHandshakesForAccountRequest) Validate() error {
	return dara.Validate(s)
}
