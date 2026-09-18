// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAgents(v []interface{}) *ListAgentsResponseBody
	GetAgents() []interface{}
	SetMaxResults(v int32) *ListAgentsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListAgentsResponseBody
	GetNextToken() *string
	SetPageNumber(v int64) *ListAgentsResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *ListAgentsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListAgentsResponseBody
	GetRequestId() *string
	SetTotal(v int64) *ListAgentsResponseBody
	GetTotal() *int64
}

type ListAgentsResponseBody struct {
	// The list of agent summaries. For field details, see "Supplementary description of response parameters".
	Agents []interface{} `json:"Agents,omitempty" xml:"Agents,omitempty" type:"Repeated"`
	// The actual cursor-based pagination size used.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The next page token. An empty string is returned if there is no next page.
	//
	// example:
	//
	// eyJwYWdlIjoyfQ.example
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The current page size for page number-based pagination.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID, used for Tracing Analysis and troubleshooting.
	//
	// example:
	//
	// 0A1B2C3D-4E5F-6789-ABCD-EF0123456789
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of agents that match the filter conditions.
	//
	// example:
	//
	// 1
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListAgentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAgentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAgentsResponseBody) GetAgents() []interface{} {
	return s.Agents
}

func (s *ListAgentsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAgentsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAgentsResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListAgentsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListAgentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAgentsResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListAgentsResponseBody) SetAgents(v []interface{}) *ListAgentsResponseBody {
	s.Agents = v
	return s
}

func (s *ListAgentsResponseBody) SetMaxResults(v int32) *ListAgentsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAgentsResponseBody) SetNextToken(v string) *ListAgentsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAgentsResponseBody) SetPageNumber(v int64) *ListAgentsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAgentsResponseBody) SetPageSize(v int64) *ListAgentsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAgentsResponseBody) SetRequestId(v string) *ListAgentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAgentsResponseBody) SetTotal(v int64) *ListAgentsResponseBody {
	s.Total = &v
	return s
}

func (s *ListAgentsResponseBody) Validate() error {
	return dara.Validate(s)
}
