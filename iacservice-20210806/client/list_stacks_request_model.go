// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStacksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *ListStacksRequest
	GetKeyword() *string
	SetMaxResults(v int32) *ListStacksRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListStacksRequest
	GetNextToken() *string
	SetPageNumber(v int32) *ListStacksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListStacksRequest
	GetPageSize() *int32
	SetStatus(v string) *ListStacksRequest
	GetStatus() *string
}

type ListStacksRequest struct {
	// example:
	//
	// key
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// LC4NJL3Ru2bIiRdnbADPQp4dD+2BRJj42DLT6GrZysw=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// Deployed
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListStacksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListStacksRequest) GoString() string {
	return s.String()
}

func (s *ListStacksRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListStacksRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListStacksRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListStacksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListStacksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListStacksRequest) GetStatus() *string {
	return s.Status
}

func (s *ListStacksRequest) SetKeyword(v string) *ListStacksRequest {
	s.Keyword = &v
	return s
}

func (s *ListStacksRequest) SetMaxResults(v int32) *ListStacksRequest {
	s.MaxResults = &v
	return s
}

func (s *ListStacksRequest) SetNextToken(v string) *ListStacksRequest {
	s.NextToken = &v
	return s
}

func (s *ListStacksRequest) SetPageNumber(v int32) *ListStacksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListStacksRequest) SetPageSize(v int32) *ListStacksRequest {
	s.PageSize = &v
	return s
}

func (s *ListStacksRequest) SetStatus(v string) *ListStacksRequest {
	s.Status = &v
	return s
}

func (s *ListStacksRequest) Validate() error {
	return dara.Validate(s)
}
