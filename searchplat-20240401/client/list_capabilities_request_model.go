// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCapabilitiesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListCapabilitiesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListCapabilitiesRequest
	GetNextToken() *string
	SetPageNumber(v string) *ListCapabilitiesRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListCapabilitiesRequest
	GetPageSize() *string
}

type ListCapabilitiesRequest struct {
	// The maximum number of entries per page for a paged query.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token.
	//
	// example:
	//
	// tMEiGtggHs0=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *string `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListCapabilitiesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCapabilitiesRequest) GoString() string {
	return s.String()
}

func (s *ListCapabilitiesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListCapabilitiesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListCapabilitiesRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListCapabilitiesRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListCapabilitiesRequest) SetMaxResults(v int32) *ListCapabilitiesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListCapabilitiesRequest) SetNextToken(v string) *ListCapabilitiesRequest {
	s.NextToken = &v
	return s
}

func (s *ListCapabilitiesRequest) SetPageNumber(v string) *ListCapabilitiesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCapabilitiesRequest) SetPageSize(v string) *ListCapabilitiesRequest {
	s.PageSize = &v
	return s
}

func (s *ListCapabilitiesRequest) Validate() error {
	return dara.Validate(s)
}
