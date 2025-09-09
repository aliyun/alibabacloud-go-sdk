// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDrdsTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeDrdsTasksResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeDrdsTasksResponseBody
	GetSuccess() *bool
	SetTasks(v *DescribeDrdsTasksResponseBodyTasks) *DescribeDrdsTasksResponseBody
	GetTasks() *DescribeDrdsTasksResponseBodyTasks
}

type DescribeDrdsTasksResponseBody struct {
	// Indicates the ID of the request.
	//
	// example:
	//
	// CD412DF7-F21D-44CE-88FF-ED24917174A8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Indicates information about the tasks.
	Tasks *DescribeDrdsTasksResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Struct"`
}

func (s DescribeDrdsTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDrdsTasksResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeDrdsTasksResponseBody) GetTasks() *DescribeDrdsTasksResponseBodyTasks {
	return s.Tasks
}

func (s *DescribeDrdsTasksResponseBody) SetRequestId(v string) *DescribeDrdsTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsTasksResponseBody) SetSuccess(v bool) *DescribeDrdsTasksResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDrdsTasksResponseBody) SetTasks(v *DescribeDrdsTasksResponseBodyTasks) *DescribeDrdsTasksResponseBody {
	s.Tasks = v
	return s
}

func (s *DescribeDrdsTasksResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDrdsTasksResponseBodyTasks struct {
	Task []*DescribeDrdsTasksResponseBodyTasksTask `json:"Task,omitempty" xml:"Task,omitempty" type:"Repeated"`
}

func (s DescribeDrdsTasksResponseBodyTasks) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsTasksResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *DescribeDrdsTasksResponseBodyTasks) GetTask() []*DescribeDrdsTasksResponseBodyTasksTask {
	return s.Task
}

func (s *DescribeDrdsTasksResponseBodyTasks) SetTask(v []*DescribeDrdsTasksResponseBodyTasksTask) *DescribeDrdsTasksResponseBodyTasks {
	s.Task = v
	return s
}

func (s *DescribeDrdsTasksResponseBodyTasks) Validate() error {
	return dara.Validate(s)
}

type DescribeDrdsTasksResponseBodyTasksTask struct {
	// Indicates the content of a task.
	//
	// example:
	//
	// upgrade_instance
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// Indicates the ID of the task.
	//
	// example:
	//
	// 64148
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Indicates the state of the task.
	//
	// example:
	//
	// FAILED
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s DescribeDrdsTasksResponseBodyTasksTask) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsTasksResponseBodyTasksTask) GoString() string {
	return s.String()
}

func (s *DescribeDrdsTasksResponseBodyTasksTask) GetContent() *string {
	return s.Content
}

func (s *DescribeDrdsTasksResponseBodyTasksTask) GetId() *int64 {
	return s.Id
}

func (s *DescribeDrdsTasksResponseBodyTasksTask) GetState() *string {
	return s.State
}

func (s *DescribeDrdsTasksResponseBodyTasksTask) SetContent(v string) *DescribeDrdsTasksResponseBodyTasksTask {
	s.Content = &v
	return s
}

func (s *DescribeDrdsTasksResponseBodyTasksTask) SetId(v int64) *DescribeDrdsTasksResponseBodyTasksTask {
	s.Id = &v
	return s
}

func (s *DescribeDrdsTasksResponseBodyTasksTask) SetState(v string) *DescribeDrdsTasksResponseBodyTasksTask {
	s.State = &v
	return s
}

func (s *DescribeDrdsTasksResponseBodyTasksTask) Validate() error {
	return dara.Validate(s)
}
