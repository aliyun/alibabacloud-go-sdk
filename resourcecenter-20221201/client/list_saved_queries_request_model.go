// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSavedQueriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v string) *ListSavedQueriesRequest
	GetMaxResults() *string
	SetNextToken(v string) *ListSavedQueriesRequest
	GetNextToken() *string
}

type ListSavedQueriesRequest struct {
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

func (s ListSavedQueriesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSavedQueriesRequest) GoString() string {
	return s.String()
}

func (s *ListSavedQueriesRequest) GetMaxResults() *string {
	return s.MaxResults
}

func (s *ListSavedQueriesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSavedQueriesRequest) SetMaxResults(v string) *ListSavedQueriesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListSavedQueriesRequest) SetNextToken(v string) *ListSavedQueriesRequest {
	s.NextToken = &v
	return s
}

func (s *ListSavedQueriesRequest) Validate() error {
	return dara.Validate(s)
}
