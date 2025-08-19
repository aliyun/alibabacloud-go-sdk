// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAsyncTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIncludePayload(v bool) *ListAsyncTasksRequest
	GetIncludePayload() *bool
	SetLimit(v int32) *ListAsyncTasksRequest
	GetLimit() *int32
	SetNextToken(v string) *ListAsyncTasksRequest
	GetNextToken() *string
	SetPrefix(v string) *ListAsyncTasksRequest
	GetPrefix() *string
	SetQualifier(v string) *ListAsyncTasksRequest
	GetQualifier() *string
	SetSortOrderByTime(v string) *ListAsyncTasksRequest
	GetSortOrderByTime() *string
	SetStartedTimeBegin(v int64) *ListAsyncTasksRequest
	GetStartedTimeBegin() *int64
	SetStartedTimeEnd(v int64) *ListAsyncTasksRequest
	GetStartedTimeEnd() *int64
	SetStatus(v string) *ListAsyncTasksRequest
	GetStatus() *string
}

type ListAsyncTasksRequest struct {
	// Specifies whether to return input parameters of the asynchronous tasks. Valid values:
	//
	// 	- true: returns the `invocationPayload` parameter in the response.
	//
	// 	- false: does not return the `invocationPayload` parameter in the response.
	//
	// >  The `invocationPayload` parameter indicates the input parameters of an asynchronous task.
	//
	// example:
	//
	// true
	IncludePayload *bool `json:"includePayload,omitempty" xml:"includePayload,omitempty"`
	// The number of asynchronous tasks to return. The default value is 20. Valid values: [1,100].
	//
	// example:
	//
	// 10
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// MTIzNCNhYmM=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The ID prefix of asynchronous tasks. If this parameter is specified, a list of asynchronous tasks whose IDs match the prefix is returned.
	//
	// example:
	//
	// job-
	Prefix *string `json:"prefix,omitempty" xml:"prefix,omitempty"`
	// The version or alias of the function.
	//
	// example:
	//
	// LATEST
	Qualifier *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
	// The order in which the returned asynchronous tasks are sorted.
	//
	// 	- asc: in ascending order.
	//
	// 	- desc: in descending order.
	//
	// example:
	//
	// asc
	SortOrderByTime *string `json:"sortOrderByTime,omitempty" xml:"sortOrderByTime,omitempty"`
	// The start time of the period during which the asynchronous tasks are initiated.
	//
	// example:
	//
	// 1640966400000
	StartedTimeBegin *int64 `json:"startedTimeBegin,omitempty" xml:"startedTimeBegin,omitempty"`
	// The end time of the period during which the asynchronous tasks are initiated.
	//
	// example:
	//
	// 1640966400000
	StartedTimeEnd *int64 `json:"startedTimeEnd,omitempty" xml:"startedTimeEnd,omitempty"`
	// The state of asynchronous tasks. The following items list the states of an asynchronous task:
	//
	// 	- Enqueued: The asynchronous invocation is enqueued and waiting to be executed.
	//
	// 	- Dequeued: The asynchronous invocation is dequeued and waiting to be triggered.
	//
	// 	- Running: The invocation is being executed.
	//
	// 	- Succeeded: The invocation is successful.
	//
	// 	- Failed: The invocation fails.
	//
	// 	- Stopped: The invocation is terminated.
	//
	// 	- Stopping: The invocation is being terminated.
	//
	// 	- Expired: The maximum validity period of messages is specified for asynchronous invocation. The invocation is discarded and not executed because the specified maximum validity period of messages expires.
	//
	// 	- Invalid: The invocation is invalid and not executed due to specific reasons. For example, the function is deleted.
	//
	// 	- Retrying: The asynchronous invocation is being retried due to an execution error.
	//
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListAsyncTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAsyncTasksRequest) GoString() string {
	return s.String()
}

func (s *ListAsyncTasksRequest) GetIncludePayload() *bool {
	return s.IncludePayload
}

func (s *ListAsyncTasksRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *ListAsyncTasksRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAsyncTasksRequest) GetPrefix() *string {
	return s.Prefix
}

func (s *ListAsyncTasksRequest) GetQualifier() *string {
	return s.Qualifier
}

func (s *ListAsyncTasksRequest) GetSortOrderByTime() *string {
	return s.SortOrderByTime
}

func (s *ListAsyncTasksRequest) GetStartedTimeBegin() *int64 {
	return s.StartedTimeBegin
}

func (s *ListAsyncTasksRequest) GetStartedTimeEnd() *int64 {
	return s.StartedTimeEnd
}

func (s *ListAsyncTasksRequest) GetStatus() *string {
	return s.Status
}

func (s *ListAsyncTasksRequest) SetIncludePayload(v bool) *ListAsyncTasksRequest {
	s.IncludePayload = &v
	return s
}

func (s *ListAsyncTasksRequest) SetLimit(v int32) *ListAsyncTasksRequest {
	s.Limit = &v
	return s
}

func (s *ListAsyncTasksRequest) SetNextToken(v string) *ListAsyncTasksRequest {
	s.NextToken = &v
	return s
}

func (s *ListAsyncTasksRequest) SetPrefix(v string) *ListAsyncTasksRequest {
	s.Prefix = &v
	return s
}

func (s *ListAsyncTasksRequest) SetQualifier(v string) *ListAsyncTasksRequest {
	s.Qualifier = &v
	return s
}

func (s *ListAsyncTasksRequest) SetSortOrderByTime(v string) *ListAsyncTasksRequest {
	s.SortOrderByTime = &v
	return s
}

func (s *ListAsyncTasksRequest) SetStartedTimeBegin(v int64) *ListAsyncTasksRequest {
	s.StartedTimeBegin = &v
	return s
}

func (s *ListAsyncTasksRequest) SetStartedTimeEnd(v int64) *ListAsyncTasksRequest {
	s.StartedTimeEnd = &v
	return s
}

func (s *ListAsyncTasksRequest) SetStatus(v string) *ListAsyncTasksRequest {
	s.Status = &v
	return s
}

func (s *ListAsyncTasksRequest) Validate() error {
	return dara.Validate(s)
}
