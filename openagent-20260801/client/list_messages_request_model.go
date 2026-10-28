// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMessagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v string) *ListMessagesRequest
	GetMaxResults() *string
	SetNextToken(v string) *ListMessagesRequest
	GetNextToken() *string
	SetRunId(v string) *ListMessagesRequest
	GetRunId() *string
}

type ListMessagesRequest struct {
	// The maximum number of timeline entries to return per page. The value must be a positive integer.
	//
	// example:
	//
	// 20
	MaxResults *string `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token. Do not specify this parameter for the first request. For subsequent requests, use the value returned from the previous page.
	//
	// example:
	//
	// example-next-token
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// Filters messages by the specified run ID.
	//
	// example:
	//
	// run_example_001
	RunId *string `json:"runId,omitempty" xml:"runId,omitempty"`
}

func (s ListMessagesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMessagesRequest) GoString() string {
	return s.String()
}

func (s *ListMessagesRequest) GetMaxResults() *string {
	return s.MaxResults
}

func (s *ListMessagesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMessagesRequest) GetRunId() *string {
	return s.RunId
}

func (s *ListMessagesRequest) SetMaxResults(v string) *ListMessagesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListMessagesRequest) SetNextToken(v string) *ListMessagesRequest {
	s.NextToken = &v
	return s
}

func (s *ListMessagesRequest) SetRunId(v string) *ListMessagesRequest {
	s.RunId = &v
	return s
}

func (s *ListMessagesRequest) Validate() error {
	return dara.Validate(s)
}
