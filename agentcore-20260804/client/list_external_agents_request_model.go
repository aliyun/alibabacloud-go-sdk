// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExternalAgentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListExternalAgentsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListExternalAgentsRequest
	GetNextToken() *string
}

type ListExternalAgentsRequest struct {
	// The maximum number of entries to return per page. Default value: 20. Valid values: 1 to 100.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token. You do not need to specify this parameter for the first request. For subsequent requests, use the nextToken value returned in the previous response.
	//
	// example:
	//
	// next-token-1
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListExternalAgentsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListExternalAgentsRequest) GoString() string {
	return s.String()
}

func (s *ListExternalAgentsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListExternalAgentsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListExternalAgentsRequest) SetMaxResults(v int32) *ListExternalAgentsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListExternalAgentsRequest) SetNextToken(v string) *ListExternalAgentsRequest {
	s.NextToken = &v
	return s
}

func (s *ListExternalAgentsRequest) Validate() error {
	return dara.Validate(s)
}
