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
	// The maximum number of API keys to return. Default value: 20. Maximum value: 100.
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
