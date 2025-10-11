// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExampleQueriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v string) *ListExampleQueriesRequest
	GetMaxResults() *string
	SetNextToken(v string) *ListExampleQueriesRequest
	GetNextToken() *string
}

type ListExampleQueriesRequest struct {
	// The maximum number of entries per page.
	//
	// Valid values: 1 to 50.
	//
	// Default value: 50.
	//
	// example:
	//
	// 10
	MaxResults *string `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// eyJzZWFyY2hBZnRlcnMiOlsiMTAwMTU2Nzk4MTU1OSJd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListExampleQueriesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListExampleQueriesRequest) GoString() string {
	return s.String()
}

func (s *ListExampleQueriesRequest) GetMaxResults() *string {
	return s.MaxResults
}

func (s *ListExampleQueriesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListExampleQueriesRequest) SetMaxResults(v string) *ListExampleQueriesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListExampleQueriesRequest) SetNextToken(v string) *ListExampleQueriesRequest {
	s.NextToken = &v
	return s
}

func (s *ListExampleQueriesRequest) Validate() error {
	return dara.Validate(s)
}
