// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListManagedAgentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListManagedAgentsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListManagedAgentsRequest
	GetNextToken() *string
}

type ListManagedAgentsRequest struct {
	// The maximum number of entries to return per page. Valid values: 0 to 100. If this parameter is omitted or set to 0, the default value is 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token. Leave this parameter empty for the first call. For subsequent calls, use the nextToken value returned in the previous response.
	//
	// example:
	//
	// next-token-1
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListManagedAgentsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListManagedAgentsRequest) GoString() string {
	return s.String()
}

func (s *ListManagedAgentsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListManagedAgentsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListManagedAgentsRequest) SetMaxResults(v int32) *ListManagedAgentsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListManagedAgentsRequest) SetNextToken(v string) *ListManagedAgentsRequest {
	s.NextToken = &v
	return s
}

func (s *ListManagedAgentsRequest) Validate() error {
	return dara.Validate(s)
}
