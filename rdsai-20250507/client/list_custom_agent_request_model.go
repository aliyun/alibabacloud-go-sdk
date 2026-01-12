// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int64) *ListCustomAgentRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListCustomAgentRequest
	GetPageSize() *int64
}

type ListCustomAgentRequest struct {
	// The operation that you want to perform. Set the value to **ListCustomAgent**.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListCustomAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCustomAgentRequest) GoString() string {
	return s.String()
}

func (s *ListCustomAgentRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListCustomAgentRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListCustomAgentRequest) SetPageNumber(v int64) *ListCustomAgentRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCustomAgentRequest) SetPageSize(v int64) *ListCustomAgentRequest {
	s.PageSize = &v
	return s
}

func (s *ListCustomAgentRequest) Validate() error {
	return dara.Validate(s)
}
