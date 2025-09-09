// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceTestTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *ListServiceTestTasksResponseBody
	GetCount() *int32
	SetMaxResults(v int32) *ListServiceTestTasksResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListServiceTestTasksResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListServiceTestTasksResponseBody
	GetRequestId() *string
	SetServiceTestTasks(v []*ListServiceTestTasksResponseBodyServiceTestTasks) *ListServiceTestTasksResponseBody
	GetServiceTestTasks() []*ListServiceTestTasksResponseBodyServiceTestTasks
}

type ListServiceTestTasksResponseBody struct {
	// The total number of entries returned.
	//
	// example:
	//
	// 9
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The number of items to return per page when paginating results. The maximum is 100, and the default is 20.
	//
	// example:
	//
	// 1
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// AAAAAfu+XtuBE55iRLHEYYuojI4=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 51945B04-6AA6-410D-93BA-236E0248B104
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The service test tasks.
	ServiceTestTasks []*ListServiceTestTasksResponseBodyServiceTestTasks `json:"ServiceTestTasks,omitempty" xml:"ServiceTestTasks,omitempty" type:"Repeated"`
}

func (s ListServiceTestTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListServiceTestTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceTestTasksResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *ListServiceTestTasksResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListServiceTestTasksResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListServiceTestTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListServiceTestTasksResponseBody) GetServiceTestTasks() []*ListServiceTestTasksResponseBodyServiceTestTasks {
	return s.ServiceTestTasks
}

func (s *ListServiceTestTasksResponseBody) SetCount(v int32) *ListServiceTestTasksResponseBody {
	s.Count = &v
	return s
}

func (s *ListServiceTestTasksResponseBody) SetMaxResults(v int32) *ListServiceTestTasksResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListServiceTestTasksResponseBody) SetNextToken(v string) *ListServiceTestTasksResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListServiceTestTasksResponseBody) SetRequestId(v string) *ListServiceTestTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceTestTasksResponseBody) SetServiceTestTasks(v []*ListServiceTestTasksResponseBodyServiceTestTasks) *ListServiceTestTasksResponseBody {
	s.ServiceTestTasks = v
	return s
}

func (s *ListServiceTestTasksResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListServiceTestTasksResponseBodyServiceTestTasks struct {
	// The time when the task was created.
	//
	// example:
	//
	// 2024-02-26T02:16:35Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// the status of service task.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The task ID.
	//
	// example:
	//
	// sttt-000h5nd4yrg59ucurzy1
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The name of the task.
	//
	// example:
	//
	// dadadad
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The task region id.
	//
	// example:
	//
	// cn-beijing
	TaskRegionId *string `json:"TaskRegionId,omitempty" xml:"TaskRegionId,omitempty"`
}

func (s ListServiceTestTasksResponseBodyServiceTestTasks) String() string {
	return dara.Prettify(s)
}

func (s ListServiceTestTasksResponseBodyServiceTestTasks) GoString() string {
	return s.String()
}

func (s *ListServiceTestTasksResponseBodyServiceTestTasks) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListServiceTestTasksResponseBodyServiceTestTasks) GetStatus() *string {
	return s.Status
}

func (s *ListServiceTestTasksResponseBodyServiceTestTasks) GetTaskId() *string {
	return s.TaskId
}

func (s *ListServiceTestTasksResponseBodyServiceTestTasks) GetTaskName() *string {
	return s.TaskName
}

func (s *ListServiceTestTasksResponseBodyServiceTestTasks) GetTaskRegionId() *string {
	return s.TaskRegionId
}

func (s *ListServiceTestTasksResponseBodyServiceTestTasks) SetCreateTime(v string) *ListServiceTestTasksResponseBodyServiceTestTasks {
	s.CreateTime = &v
	return s
}

func (s *ListServiceTestTasksResponseBodyServiceTestTasks) SetStatus(v string) *ListServiceTestTasksResponseBodyServiceTestTasks {
	s.Status = &v
	return s
}

func (s *ListServiceTestTasksResponseBodyServiceTestTasks) SetTaskId(v string) *ListServiceTestTasksResponseBodyServiceTestTasks {
	s.TaskId = &v
	return s
}

func (s *ListServiceTestTasksResponseBodyServiceTestTasks) SetTaskName(v string) *ListServiceTestTasksResponseBodyServiceTestTasks {
	s.TaskName = &v
	return s
}

func (s *ListServiceTestTasksResponseBodyServiceTestTasks) SetTaskRegionId(v string) *ListServiceTestTasksResponseBodyServiceTestTasks {
	s.TaskRegionId = &v
	return s
}

func (s *ListServiceTestTasksResponseBodyServiceTestTasks) Validate() error {
	return dara.Validate(s)
}
