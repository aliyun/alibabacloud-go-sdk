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
	SetMemoryStoreName(v string) *ListMemoryStoresRequest
	GetMemoryStoreName() *string
	SetNextToken(v string) *ListMemoryStoresRequest
	GetNextToken() *string
}

type ListMemoryStoresRequest struct {
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// qianyi_test_1
	MemoryStoreName *string `json:"memoryStoreName,omitempty" xml:"memoryStoreName,omitempty"`
	// example:
	//
	// xxxxxxxxx
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

func (s *ListMemoryStoresRequest) GetMemoryStoreName() *string {
	return s.MemoryStoreName
}

func (s *ListMemoryStoresRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMemoryStoresRequest) SetMaxResults(v int32) *ListMemoryStoresRequest {
	s.MaxResults = &v
	return s
}

func (s *ListMemoryStoresRequest) SetMemoryStoreName(v string) *ListMemoryStoresRequest {
	s.MemoryStoreName = &v
	return s
}

func (s *ListMemoryStoresRequest) SetNextToken(v string) *ListMemoryStoresRequest {
	s.NextToken = &v
	return s
}

func (s *ListMemoryStoresRequest) Validate() error {
	return dara.Validate(s)
}
