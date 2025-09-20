// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTasksShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTimeRangeShrink(v string) *ListTasksShrinkRequest
	GetEndTimeRangeShrink() *string
	SetMaxResults(v int64) *ListTasksShrinkRequest
	GetMaxResults() *int64
	SetNextToken(v string) *ListTasksShrinkRequest
	GetNextToken() *string
	SetOrder(v string) *ListTasksShrinkRequest
	GetOrder() *string
	SetProjectName(v string) *ListTasksShrinkRequest
	GetProjectName() *string
	SetRequestDefinition(v bool) *ListTasksShrinkRequest
	GetRequestDefinition() *bool
	SetSort(v string) *ListTasksShrinkRequest
	GetSort() *string
	SetStartTimeRangeShrink(v string) *ListTasksShrinkRequest
	GetStartTimeRangeShrink() *string
	SetStatus(v string) *ListTasksShrinkRequest
	GetStatus() *string
	SetTagSelector(v string) *ListTasksShrinkRequest
	GetTagSelector() *string
	SetTaskTypesShrink(v string) *ListTasksShrinkRequest
	GetTaskTypesShrink() *string
}

type ListTasksShrinkRequest struct {
	// The task end time range. You can specify this parameter to filter tasks that end within the specified range.
	EndTimeRangeShrink *string `json:"EndTimeRange,omitempty" xml:"EndTimeRange,omitempty"`
	// The maximum number of results to return. Valid value range: (0, 100]. Default value: 100.
	//
	// example:
	//
	// 1
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token.
	//
	// The pagination token is used in the next request to retrieve a new page of results if the total number of results exceeds the value of the MaxResults parameter. The next call to the operation returns results lexicographically after the NextToken parameter value.
	//
	// >  Leave this parameter empty in your first call to the operation.
	//
	// example:
	//
	// MTIzNDU2Nzg6aW1tdGVzdDpleGFtcGxlYnVja2V0OmRhdGFzZXQwMDE6b3NzOi8vZXhhbXBsZWJ1Y2tldC9zYW1wbGVvYmplY3QxLmpwZw==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The sort order. Valid values:
	//
	// 	- asc: in ascending order. This is the default value.
	//
	// 	- desc: in descending order.
	//
	// example:
	//
	// ASC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The name of the project.[](~~478153~~)
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// Specifies whether to return request parameters in the initial request to create the task. Default value: False.
	//
	// example:
	//
	// True
	RequestDefinition *bool `json:"RequestDefinition,omitempty" xml:"RequestDefinition,omitempty"`
	// The field used to sort the results by. Valid values:
	//
	// 	- TaskId: sorts the results by task ID. This is the default sort field.
	//
	// 	- StartTime: sorts the results by task start time.
	//
	// 	- StartTime: sorts the results by task end time.
	//
	// example:
	//
	// TaskId
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// The task start time range. You can specify this parameter to filter tasks that start within the specified range.
	StartTimeRangeShrink *string `json:"StartTimeRange,omitempty" xml:"StartTimeRange,omitempty"`
	// The task status. Valid values:
	//
	// 	- Running: The task is running.
	//
	// 	- Succeeded: The task is successful.
	//
	// 	- Failed: The task failed.
	//
	// example:
	//
	// Succeeded
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The custom tags of tasks.
	//
	// example:
	//
	// test=val1
	TagSelector *string `json:"TagSelector,omitempty" xml:"TagSelector,omitempty"`
	// The task types.
	TaskTypesShrink *string `json:"TaskTypes,omitempty" xml:"TaskTypes,omitempty"`
}

func (s ListTasksShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTasksShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListTasksShrinkRequest) GetEndTimeRangeShrink() *string {
	return s.EndTimeRangeShrink
}

func (s *ListTasksShrinkRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListTasksShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTasksShrinkRequest) GetOrder() *string {
	return s.Order
}

func (s *ListTasksShrinkRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *ListTasksShrinkRequest) GetRequestDefinition() *bool {
	return s.RequestDefinition
}

func (s *ListTasksShrinkRequest) GetSort() *string {
	return s.Sort
}

func (s *ListTasksShrinkRequest) GetStartTimeRangeShrink() *string {
	return s.StartTimeRangeShrink
}

func (s *ListTasksShrinkRequest) GetStatus() *string {
	return s.Status
}

func (s *ListTasksShrinkRequest) GetTagSelector() *string {
	return s.TagSelector
}

func (s *ListTasksShrinkRequest) GetTaskTypesShrink() *string {
	return s.TaskTypesShrink
}

func (s *ListTasksShrinkRequest) SetEndTimeRangeShrink(v string) *ListTasksShrinkRequest {
	s.EndTimeRangeShrink = &v
	return s
}

func (s *ListTasksShrinkRequest) SetMaxResults(v int64) *ListTasksShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTasksShrinkRequest) SetNextToken(v string) *ListTasksShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListTasksShrinkRequest) SetOrder(v string) *ListTasksShrinkRequest {
	s.Order = &v
	return s
}

func (s *ListTasksShrinkRequest) SetProjectName(v string) *ListTasksShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *ListTasksShrinkRequest) SetRequestDefinition(v bool) *ListTasksShrinkRequest {
	s.RequestDefinition = &v
	return s
}

func (s *ListTasksShrinkRequest) SetSort(v string) *ListTasksShrinkRequest {
	s.Sort = &v
	return s
}

func (s *ListTasksShrinkRequest) SetStartTimeRangeShrink(v string) *ListTasksShrinkRequest {
	s.StartTimeRangeShrink = &v
	return s
}

func (s *ListTasksShrinkRequest) SetStatus(v string) *ListTasksShrinkRequest {
	s.Status = &v
	return s
}

func (s *ListTasksShrinkRequest) SetTagSelector(v string) *ListTasksShrinkRequest {
	s.TagSelector = &v
	return s
}

func (s *ListTasksShrinkRequest) SetTaskTypesShrink(v string) *ListTasksShrinkRequest {
	s.TaskTypesShrink = &v
	return s
}

func (s *ListTasksShrinkRequest) Validate() error {
	return dara.Validate(s)
}
