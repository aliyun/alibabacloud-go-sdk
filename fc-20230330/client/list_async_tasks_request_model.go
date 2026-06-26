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
	// Specifies whether to return the input parameters of the asynchronous task.
	//
	// - true: If this parameter is set to true, the `invocationPayload` field is returned.
	//
	// - false: If this parameter is set to false, the `invocationPayload` field is not returned.
	//
	// > The `invocationPayload` field specifies the input parameters of the function for the asynchronous task.
	//
	// example:
	//
	// true
	IncludePayload *bool `json:"includePayload,omitempty" xml:"includePayload,omitempty"`
	// The number of asynchronous tasks to return. The default value is 20. The value must be in the range of [1, 100].
	//
	// example:
	//
	// 10
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// The pagination token to return more results. You do not need to provide this parameter for the first query. Obtain the token for a subsequent query from the response to the previous query.
	//
	// example:
	//
	// MTIzNCNhYmM=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The prefix of the asynchronous task ID. The system returns a list of asynchronous tasks that match the prefix.
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
	// The sorting order of the returned asynchronous tasks.
	//
	// - asc: ascending order
	//
	// - desc: descending order
	//
	// example:
	//
	// asc
	SortOrderByTime *string `json:"sortOrderByTime,omitempty" xml:"sortOrderByTime,omitempty"`
	// The start of the time range when the asynchronous task was started.
	//
	// example:
	//
	// 1640966400000
	StartedTimeBegin *int64 `json:"startedTimeBegin,omitempty" xml:"startedTimeBegin,omitempty"`
	// The end of the time range when the asynchronous task was started.
	//
	// example:
	//
	// 1640966400000
	StartedTimeEnd *int64 `json:"startedTimeEnd,omitempty" xml:"startedTimeEnd,omitempty"`
	// The execution status of the asynchronous task.
	//
	// - Enqueued: The asynchronous message is enqueued and waits for processing.
	//
	// - Dequeued: The asynchronous message is dequeued and waits to be triggered.
	//
	// - Running: The invocation is in progress.
	//
	// - Succeeded: The invocation succeeded.
	//
	// - Failed: The invocation failed.
	//
	// - Stopped: The invocation was stopped.
	//
	// - Stopping: The invocation is being stopped.
	//
	// - Expired: The task was discarded because its configured maximum queuing duration was exceeded. The task was not executed.
	//
	// - Invalid: The execution is invalid because the function was deleted or for other reasons. The task was not executed.
	//
	// - Retrying: The asynchronous invocation is being retried because of an execution error.
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
