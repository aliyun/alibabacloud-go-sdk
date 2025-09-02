// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateTableResponseBody
	GetRequestId() *string
	SetTaskInfo(v *CreateTableResponseBodyTaskInfo) *CreateTableResponseBody
	GetTaskInfo() *CreateTableResponseBodyTaskInfo
}

type CreateTableResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// abcdef
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the request task. After a request task is submitted, it is divided into multiple subtasks that are run in sequence. After the current subtask is complete, the next subtask starts to run. After all subtasks are complete, the request task is complete.
	//
	// If a request task is aborted due to one of the following issues, address the issue based on the error code and initiate the request task again:
	//
	// 	- The request task fails to be submitted.
	//
	// 	- After the request task is submitted, a subtask fails to run.
	TaskInfo *CreateTableResponseBodyTaskInfo `json:"TaskInfo,omitempty" xml:"TaskInfo,omitempty" type:"Struct"`
}

func (s CreateTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTableResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTableResponseBody) GetTaskInfo() *CreateTableResponseBodyTaskInfo {
	return s.TaskInfo
}

func (s *CreateTableResponseBody) SetRequestId(v string) *CreateTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTableResponseBody) SetTaskInfo(v *CreateTableResponseBodyTaskInfo) *CreateTableResponseBody {
	s.TaskInfo = v
	return s
}

func (s *CreateTableResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateTableResponseBodyTaskInfo struct {
	// The details about the status of the current subtask.
	//
	// 	- If the current subtask is successful, success is returned.
	//
	// 	- If the current subtask fails, the error details are displayed.
	//
	// example:
	//
	// success
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The ID of the subtask that you want to run. If this parameter is left empty, all subtasks are complete.
	//
	// example:
	//
	// abc1
	NextTaskId *string `json:"NextTaskId,omitempty" xml:"NextTaskId,omitempty"`
	// The status of the current subtask. Valid values:
	//
	// 	- operating: The subtask is running.
	//
	// 	- success: The subtask succeeds.
	//
	// 	- failure: The subtask fails to run. For more information about the error details, see the Content parameter.
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the current subtask.
	//
	// example:
	//
	// abc
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateTableResponseBodyTaskInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateTableResponseBodyTaskInfo) GoString() string {
	return s.String()
}

func (s *CreateTableResponseBodyTaskInfo) GetContent() *string {
	return s.Content
}

func (s *CreateTableResponseBodyTaskInfo) GetNextTaskId() *string {
	return s.NextTaskId
}

func (s *CreateTableResponseBodyTaskInfo) GetStatus() *string {
	return s.Status
}

func (s *CreateTableResponseBodyTaskInfo) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateTableResponseBodyTaskInfo) SetContent(v string) *CreateTableResponseBodyTaskInfo {
	s.Content = &v
	return s
}

func (s *CreateTableResponseBodyTaskInfo) SetNextTaskId(v string) *CreateTableResponseBodyTaskInfo {
	s.NextTaskId = &v
	return s
}

func (s *CreateTableResponseBodyTaskInfo) SetStatus(v string) *CreateTableResponseBodyTaskInfo {
	s.Status = &v
	return s
}

func (s *CreateTableResponseBodyTaskInfo) SetTaskId(v string) *CreateTableResponseBodyTaskInfo {
	s.TaskId = &v
	return s
}

func (s *CreateTableResponseBodyTaskInfo) Validate() error {
	return dara.Validate(s)
}
