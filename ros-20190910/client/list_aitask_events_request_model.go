// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAITaskEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v string) *ListAITaskEventsRequest
	GetMaxResults() *string
	SetNextToken(v string) *ListAITaskEventsRequest
	GetNextToken() *string
	SetTaskId(v string) *ListAITaskEventsRequest
	GetTaskId() *string
}

type ListAITaskEventsRequest struct {
	// The maximum number of results to be returned from a single query when the NextToken parameter is used in the query.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 50.
	//
	// example:
	//
	// 50
	MaxResults *string `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// AAAAAdDWBF2****w==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the AI task.
	//
	// example:
	//
	// t-asasas*****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ListAITaskEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAITaskEventsRequest) GoString() string {
	return s.String()
}

func (s *ListAITaskEventsRequest) GetMaxResults() *string {
	return s.MaxResults
}

func (s *ListAITaskEventsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAITaskEventsRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *ListAITaskEventsRequest) SetMaxResults(v string) *ListAITaskEventsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAITaskEventsRequest) SetNextToken(v string) *ListAITaskEventsRequest {
	s.NextToken = &v
	return s
}

func (s *ListAITaskEventsRequest) SetTaskId(v string) *ListAITaskEventsRequest {
	s.TaskId = &v
	return s
}

func (s *ListAITaskEventsRequest) Validate() error {
	return dara.Validate(s)
}
