// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTimeRange(v *TimeRange) *ListTasksRequest
	GetEndTimeRange() *TimeRange
	SetMaxResults(v int64) *ListTasksRequest
	GetMaxResults() *int64
	SetNextToken(v string) *ListTasksRequest
	GetNextToken() *string
	SetOrder(v string) *ListTasksRequest
	GetOrder() *string
	SetProjectName(v string) *ListTasksRequest
	GetProjectName() *string
	SetRequestDefinition(v bool) *ListTasksRequest
	GetRequestDefinition() *bool
	SetSort(v string) *ListTasksRequest
	GetSort() *string
	SetStartTimeRange(v *TimeRange) *ListTasksRequest
	GetStartTimeRange() *TimeRange
	SetStatus(v string) *ListTasksRequest
	GetStatus() *string
	SetTagSelector(v string) *ListTasksRequest
	GetTagSelector() *string
	SetTaskTypes(v []*string) *ListTasksRequest
	GetTaskTypes() []*string
}

type ListTasksRequest struct {
	// The task end time range. You can specify this parameter to filter tasks that end within the specified range.
	EndTimeRange *TimeRange `json:"EndTimeRange,omitempty" xml:"EndTimeRange,omitempty"`
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
	StartTimeRange *TimeRange `json:"StartTimeRange,omitempty" xml:"StartTimeRange,omitempty"`
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
	TaskTypes []*string `json:"TaskTypes,omitempty" xml:"TaskTypes,omitempty" type:"Repeated"`
}

func (s ListTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTasksRequest) GoString() string {
	return s.String()
}

func (s *ListTasksRequest) GetEndTimeRange() *TimeRange {
	return s.EndTimeRange
}

func (s *ListTasksRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListTasksRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTasksRequest) GetOrder() *string {
	return s.Order
}

func (s *ListTasksRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *ListTasksRequest) GetRequestDefinition() *bool {
	return s.RequestDefinition
}

func (s *ListTasksRequest) GetSort() *string {
	return s.Sort
}

func (s *ListTasksRequest) GetStartTimeRange() *TimeRange {
	return s.StartTimeRange
}

func (s *ListTasksRequest) GetStatus() *string {
	return s.Status
}

func (s *ListTasksRequest) GetTagSelector() *string {
	return s.TagSelector
}

func (s *ListTasksRequest) GetTaskTypes() []*string {
	return s.TaskTypes
}

func (s *ListTasksRequest) SetEndTimeRange(v *TimeRange) *ListTasksRequest {
	s.EndTimeRange = v
	return s
}

func (s *ListTasksRequest) SetMaxResults(v int64) *ListTasksRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTasksRequest) SetNextToken(v string) *ListTasksRequest {
	s.NextToken = &v
	return s
}

func (s *ListTasksRequest) SetOrder(v string) *ListTasksRequest {
	s.Order = &v
	return s
}

func (s *ListTasksRequest) SetProjectName(v string) *ListTasksRequest {
	s.ProjectName = &v
	return s
}

func (s *ListTasksRequest) SetRequestDefinition(v bool) *ListTasksRequest {
	s.RequestDefinition = &v
	return s
}

func (s *ListTasksRequest) SetSort(v string) *ListTasksRequest {
	s.Sort = &v
	return s
}

func (s *ListTasksRequest) SetStartTimeRange(v *TimeRange) *ListTasksRequest {
	s.StartTimeRange = v
	return s
}

func (s *ListTasksRequest) SetStatus(v string) *ListTasksRequest {
	s.Status = &v
	return s
}

func (s *ListTasksRequest) SetTagSelector(v string) *ListTasksRequest {
	s.TagSelector = &v
	return s
}

func (s *ListTasksRequest) SetTaskTypes(v []*string) *ListTasksRequest {
	s.TaskTypes = v
	return s
}

func (s *ListTasksRequest) Validate() error {
	if s.EndTimeRange != nil {
		if err := s.EndTimeRange.Validate(); err != nil {
			return err
		}
	}
	if s.StartTimeRange != nil {
		if err := s.StartTimeRange.Validate(); err != nil {
			return err
		}
	}
	return nil
}
