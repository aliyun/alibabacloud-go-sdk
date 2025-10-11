// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExampleQueriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExampleQueries(v []*ListExampleQueriesResponseBodyExampleQueries) *ListExampleQueriesResponseBody
	GetExampleQueries() []*ListExampleQueriesResponseBodyExampleQueries
	SetMaxResults(v string) *ListExampleQueriesResponseBody
	GetMaxResults() *string
	SetNextToken(v string) *ListExampleQueriesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListExampleQueriesResponseBody
	GetRequestId() *string
}

type ListExampleQueriesResponseBody struct {
	// The information about the sample query templates.
	ExampleQueries []*ListExampleQueriesResponseBodyExampleQueries `json:"ExampleQueries,omitempty" xml:"ExampleQueries,omitempty" type:"Repeated"`
	// The maximum number of entries per page.
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
	// The request ID.
	//
	// example:
	//
	// D696E6EF-3A6D-5770-801E-4982081FE4D0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListExampleQueriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListExampleQueriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListExampleQueriesResponseBody) GetExampleQueries() []*ListExampleQueriesResponseBodyExampleQueries {
	return s.ExampleQueries
}

func (s *ListExampleQueriesResponseBody) GetMaxResults() *string {
	return s.MaxResults
}

func (s *ListExampleQueriesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListExampleQueriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListExampleQueriesResponseBody) SetExampleQueries(v []*ListExampleQueriesResponseBodyExampleQueries) *ListExampleQueriesResponseBody {
	s.ExampleQueries = v
	return s
}

func (s *ListExampleQueriesResponseBody) SetMaxResults(v string) *ListExampleQueriesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListExampleQueriesResponseBody) SetNextToken(v string) *ListExampleQueriesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListExampleQueriesResponseBody) SetRequestId(v string) *ListExampleQueriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListExampleQueriesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListExampleQueriesResponseBodyExampleQueries struct {
	// The description of the template.
	//
	// example:
	//
	// Queries all resources on which you have permissions and sorts the resources by resource type and resource ID.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the template.
	//
	// example:
	//
	// Query of All Alibaba Cloud Resources
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the template.
	//
	// example:
	//
	// sq-0PfKy****
	QueryId *string `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
}

func (s ListExampleQueriesResponseBodyExampleQueries) String() string {
	return dara.Prettify(s)
}

func (s ListExampleQueriesResponseBodyExampleQueries) GoString() string {
	return s.String()
}

func (s *ListExampleQueriesResponseBodyExampleQueries) GetDescription() *string {
	return s.Description
}

func (s *ListExampleQueriesResponseBodyExampleQueries) GetName() *string {
	return s.Name
}

func (s *ListExampleQueriesResponseBodyExampleQueries) GetQueryId() *string {
	return s.QueryId
}

func (s *ListExampleQueriesResponseBodyExampleQueries) SetDescription(v string) *ListExampleQueriesResponseBodyExampleQueries {
	s.Description = &v
	return s
}

func (s *ListExampleQueriesResponseBodyExampleQueries) SetName(v string) *ListExampleQueriesResponseBodyExampleQueries {
	s.Name = &v
	return s
}

func (s *ListExampleQueriesResponseBodyExampleQueries) SetQueryId(v string) *ListExampleQueriesResponseBodyExampleQueries {
	s.QueryId = &v
	return s
}

func (s *ListExampleQueriesResponseBodyExampleQueries) Validate() error {
	return dara.Validate(s)
}
