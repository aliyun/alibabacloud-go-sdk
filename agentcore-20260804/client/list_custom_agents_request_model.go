// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomAgentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListCustomAgentsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListCustomAgentsRequest
	GetNextToken() *string
}

type ListCustomAgentsRequest struct {
	// The maximum number of entries to return per page. Valid values: 1 to 100. If this parameter is omitted or set to 0, the default value 20 is used.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The opaque token used to query the next page. Omit this parameter in the first request. In subsequent requests, pass the nextToken value returned in the previous response as is, and keep the same workspace. Do not construct the token yourself.
	//
	// example:
	//
	// djE6YWdlbnRjb3JlLmFnZW50Lmxpc3QuY3VzdG9tOjIw
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListCustomAgentsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCustomAgentsRequest) GoString() string {
	return s.String()
}

func (s *ListCustomAgentsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListCustomAgentsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListCustomAgentsRequest) SetMaxResults(v int32) *ListCustomAgentsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListCustomAgentsRequest) SetNextToken(v string) *ListCustomAgentsRequest {
	s.NextToken = &v
	return s
}

func (s *ListCustomAgentsRequest) Validate() error {
	return dara.Validate(s)
}
