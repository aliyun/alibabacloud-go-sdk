// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiId(v string) *ListCustomAgentRequest
	GetApiId() *string
	SetPageNumber(v int64) *ListCustomAgentRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListCustomAgentRequest
	GetPageSize() *int64
}

type ListCustomAgentRequest struct {
	// example:
	//
	// app-iBuGU1VxEY42zrQRQfNA****
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
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

func (s *ListCustomAgentRequest) GetApiId() *string {
	return s.ApiId
}

func (s *ListCustomAgentRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListCustomAgentRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListCustomAgentRequest) SetApiId(v string) *ListCustomAgentRequest {
	s.ApiId = &v
	return s
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
