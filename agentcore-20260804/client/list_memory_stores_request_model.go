// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMemoryStoresRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListMemoryStoresRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListMemoryStoresRequest
	GetNextToken() *string
}

type ListMemoryStoresRequest struct {
	// The maximum number of records per page. Default value: 10. Valid values: 1 to 100.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token for the next page.
	//
	// example:
	//
	// eyJvZmZzZXQiOjEwfQ==
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListMemoryStoresRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMemoryStoresRequest) GoString() string {
	return s.String()
}

func (s *ListMemoryStoresRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListMemoryStoresRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMemoryStoresRequest) SetMaxResults(v int32) *ListMemoryStoresRequest {
	s.MaxResults = &v
	return s
}

func (s *ListMemoryStoresRequest) SetNextToken(v string) *ListMemoryStoresRequest {
	s.NextToken = &v
	return s
}

func (s *ListMemoryStoresRequest) Validate() error {
	return dara.Validate(s)
}
