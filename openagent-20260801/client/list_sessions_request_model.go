// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSessionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v string) *ListSessionsRequest
	GetMaxResults() *string
	SetNextToken(v string) *ListSessionsRequest
	GetNextToken() *string
}

type ListSessionsRequest struct {
	// The maximum number of sessions to return per page. The value must be a positive integer.
	//
	// example:
	//
	// 20
	MaxResults *string `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token. Do not specify this parameter for the first request. For subsequent requests, pass the nextToken value returned in the previous response. Do not parse or construct this value.
	//
	// example:
	//
	// example-next-token
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListSessionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSessionsRequest) GoString() string {
	return s.String()
}

func (s *ListSessionsRequest) GetMaxResults() *string {
	return s.MaxResults
}

func (s *ListSessionsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSessionsRequest) SetMaxResults(v string) *ListSessionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListSessionsRequest) SetNextToken(v string) *ListSessionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListSessionsRequest) Validate() error {
	return dara.Validate(s)
}
