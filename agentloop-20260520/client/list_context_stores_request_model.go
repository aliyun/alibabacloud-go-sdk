// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListContextStoresRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContextStoreName(v string) *ListContextStoresRequest
	GetContextStoreName() *string
	SetContextType(v string) *ListContextStoresRequest
	GetContextType() *string
	SetMaxResults(v int32) *ListContextStoresRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListContextStoresRequest
	GetNextToken() *string
}

type ListContextStoresRequest struct {
	// Filters context stores by name. Exact match is supported. If this parameter is not specified, no filtering is applied.
	//
	// example:
	//
	// my-context-store
	ContextStoreName *string `json:"contextStoreName,omitempty" xml:"contextStoreName,omitempty"`
	// Filters context stores by type, such as experience or memory. If this parameter is not specified, no filtering is applied.
	//
	// example:
	//
	// experience
	ContextType *string `json:"contextType,omitempty" xml:"contextType,omitempty"`
	// The maximum number of context stores to return. Default value: 20. Maximum value: 100.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token. Set this parameter to the nextToken value returned in the previous response to retrieve the next page. Do not specify this parameter for the first request.
	//
	// example:
	//
	// MTIzNDU2Nzg5MA==
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListContextStoresRequest) String() string {
	return dara.Prettify(s)
}

func (s ListContextStoresRequest) GoString() string {
	return s.String()
}

func (s *ListContextStoresRequest) GetContextStoreName() *string {
	return s.ContextStoreName
}

func (s *ListContextStoresRequest) GetContextType() *string {
	return s.ContextType
}

func (s *ListContextStoresRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListContextStoresRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListContextStoresRequest) SetContextStoreName(v string) *ListContextStoresRequest {
	s.ContextStoreName = &v
	return s
}

func (s *ListContextStoresRequest) SetContextType(v string) *ListContextStoresRequest {
	s.ContextType = &v
	return s
}

func (s *ListContextStoresRequest) SetMaxResults(v int32) *ListContextStoresRequest {
	s.MaxResults = &v
	return s
}

func (s *ListContextStoresRequest) SetNextToken(v string) *ListContextStoresRequest {
	s.NextToken = &v
	return s
}

func (s *ListContextStoresRequest) Validate() error {
	return dara.Validate(s)
}
