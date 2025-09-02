// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTableAddColumnResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateTableAddColumnResponseBody
	GetRequestId() *string
	SetTaskInfo(v *UpdateTableAddColumnResponseBodyTaskInfo) *UpdateTableAddColumnResponseBody
	GetTaskInfo() *UpdateTableAddColumnResponseBodyTaskInfo
}

type UpdateTableAddColumnResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// abc
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the request task. After a request task is submitted, it is divided into multiple subtasks that are run in sequence. After the current subtask is complete, the next subtask starts to run. After all subtasks are complete, the request task is complete. If a request task is aborted due to one of the following issues, address the issue based on the error code and initiate the request task again:
	//
	// 	- The request task fails to be submitted.
	//
	// 	- After the request task is submitted, a subtask fails to run.
	TaskInfo *UpdateTableAddColumnResponseBodyTaskInfo `json:"TaskInfo,omitempty" xml:"TaskInfo,omitempty" type:"Struct"`
}

func (s UpdateTableAddColumnResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTableAddColumnResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTableAddColumnResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTableAddColumnResponseBody) GetTaskInfo() *UpdateTableAddColumnResponseBodyTaskInfo {
	return s.TaskInfo
}

func (s *UpdateTableAddColumnResponseBody) SetRequestId(v string) *UpdateTableAddColumnResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTableAddColumnResponseBody) SetTaskInfo(v *UpdateTableAddColumnResponseBodyTaskInfo) *UpdateTableAddColumnResponseBody {
	s.TaskInfo = v
	return s
}

func (s *UpdateTableAddColumnResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdateTableAddColumnResponseBodyTaskInfo struct {
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
	// abc2
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UpdateTableAddColumnResponseBodyTaskInfo) String() string {
	return dara.Prettify(s)
}

func (s UpdateTableAddColumnResponseBodyTaskInfo) GoString() string {
	return s.String()
}

func (s *UpdateTableAddColumnResponseBodyTaskInfo) GetContent() *string {
	return s.Content
}

func (s *UpdateTableAddColumnResponseBodyTaskInfo) GetNextTaskId() *string {
	return s.NextTaskId
}

func (s *UpdateTableAddColumnResponseBodyTaskInfo) GetStatus() *string {
	return s.Status
}

func (s *UpdateTableAddColumnResponseBodyTaskInfo) GetTaskId() *string {
	return s.TaskId
}

func (s *UpdateTableAddColumnResponseBodyTaskInfo) SetContent(v string) *UpdateTableAddColumnResponseBodyTaskInfo {
	s.Content = &v
	return s
}

func (s *UpdateTableAddColumnResponseBodyTaskInfo) SetNextTaskId(v string) *UpdateTableAddColumnResponseBodyTaskInfo {
	s.NextTaskId = &v
	return s
}

func (s *UpdateTableAddColumnResponseBodyTaskInfo) SetStatus(v string) *UpdateTableAddColumnResponseBodyTaskInfo {
	s.Status = &v
	return s
}

func (s *UpdateTableAddColumnResponseBodyTaskInfo) SetTaskId(v string) *UpdateTableAddColumnResponseBodyTaskInfo {
	s.TaskId = &v
	return s
}

func (s *UpdateTableAddColumnResponseBodyTaskInfo) Validate() error {
	return dara.Validate(s)
}
