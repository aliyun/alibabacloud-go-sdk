// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListContextStoreAPIKeysRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListContextStoreAPIKeysRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListContextStoreAPIKeysRequest
	GetNextToken() *string
}

type ListContextStoreAPIKeysRequest struct {
	// The maximum number of entries to return in a single request. If this parameter is not specified, the service returns a default number of entries.
	//
	// example:
	//
	// 100
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token used to retrieve the next page of results. Leave this parameter empty to retrieve the first page.
	//
	// example:
	//
	// 615dfe7fd00f699ea94d5e63ba564aaf22450858c58f1387cc78f883b2254ab47232ce40fa95d9cb
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListContextStoreAPIKeysRequest) String() string {
	return dara.Prettify(s)
}

func (s ListContextStoreAPIKeysRequest) GoString() string {
	return s.String()
}

func (s *ListContextStoreAPIKeysRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListContextStoreAPIKeysRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListContextStoreAPIKeysRequest) SetMaxResults(v int32) *ListContextStoreAPIKeysRequest {
	s.MaxResults = &v
	return s
}

func (s *ListContextStoreAPIKeysRequest) SetNextToken(v string) *ListContextStoreAPIKeysRequest {
	s.NextToken = &v
	return s
}

func (s *ListContextStoreAPIKeysRequest) Validate() error {
	return dara.Validate(s)
}
