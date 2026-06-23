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
	SetQ(v string) *ListCustomAgentsRequest
	GetQ() *string
	SetVisibility(v []*string) *ListCustomAgentsRequest
	GetVisibility() []*string
}

type ListCustomAgentsRequest struct {
	// The maximum number of entries to return on a single page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token. Omit this parameter for the first request. For subsequent requests, use the `NextToken` value from the previous response to retrieve the next page.
	//
	// example:
	//
	// 12345
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The search keyword for a fuzzy match by agent name.
	//
	// example:
	//
	// analysis
	Q *string `json:"Q,omitempty" xml:"Q,omitempty"`
	// Filters the results by visibility level. You can specify multiple levels.
	//
	// example:
	//
	// -
	Visibility []*string `json:"Visibility,omitempty" xml:"Visibility,omitempty" type:"Repeated"`
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

func (s *ListCustomAgentsRequest) GetQ() *string {
	return s.Q
}

func (s *ListCustomAgentsRequest) GetVisibility() []*string {
	return s.Visibility
}

func (s *ListCustomAgentsRequest) SetMaxResults(v int32) *ListCustomAgentsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListCustomAgentsRequest) SetNextToken(v string) *ListCustomAgentsRequest {
	s.NextToken = &v
	return s
}

func (s *ListCustomAgentsRequest) SetQ(v string) *ListCustomAgentsRequest {
	s.Q = &v
	return s
}

func (s *ListCustomAgentsRequest) SetVisibility(v []*string) *ListCustomAgentsRequest {
	s.Visibility = v
	return s
}

func (s *ListCustomAgentsRequest) Validate() error {
	return dara.Validate(s)
}
