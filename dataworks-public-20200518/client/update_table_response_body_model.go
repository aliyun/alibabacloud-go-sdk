// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateTableResponseBody
	GetRequestId() *string
	SetTaskInfo(v *UpdateTableResponseBodyTaskInfo) *UpdateTableResponseBody
	GetTaskInfo() *UpdateTableResponseBodyTaskInfo
}

type UpdateTableResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// abcde
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the request task. After a request task is submitted, it is divided into multiple subtasks that are run in sequence. After the current subtask is complete, the next subtask starts to run. After all subtasks are complete, the request task is complete. If a request task is terminated due to one of the following issues, address the issue based on the error code and initiate the request task again:
	//
	// 	- The request task fails to be submitted.
	//
	// 	- After the request task is submitted, a subtask fails to run.
	TaskInfo *UpdateTableResponseBodyTaskInfo `json:"TaskInfo,omitempty" xml:"TaskInfo,omitempty" type:"Struct"`
}

func (s UpdateTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTableResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTableResponseBody) GetTaskInfo() *UpdateTableResponseBodyTaskInfo {
	return s.TaskInfo
}

func (s *UpdateTableResponseBody) SetRequestId(v string) *UpdateTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTableResponseBody) SetTaskInfo(v *UpdateTableResponseBodyTaskInfo) *UpdateTableResponseBody {
	s.TaskInfo = v
	return s
}

func (s *UpdateTableResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdateTableResponseBodyTaskInfo struct {
	// The details about the status of the current subtask. Valid values:
	//
	// 	- If the current subtask is successful, success is returned.
	//
	// 	- If the current subtask fails, the error details are displayed.
	//
	// example:
	//
	// success
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The ID of the subtask that you want to run. If this parameter is left empty, all subtasks are complete. You can call the [GetDDLJobStatus](https://help.aliyun.com/document_detail/185659.html) operation to query the status of the subtask based on the subtask ID.
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

func (s UpdateTableResponseBodyTaskInfo) String() string {
	return dara.Prettify(s)
}

func (s UpdateTableResponseBodyTaskInfo) GoString() string {
	return s.String()
}

func (s *UpdateTableResponseBodyTaskInfo) GetContent() *string {
	return s.Content
}

func (s *UpdateTableResponseBodyTaskInfo) GetNextTaskId() *string {
	return s.NextTaskId
}

func (s *UpdateTableResponseBodyTaskInfo) GetStatus() *string {
	return s.Status
}

func (s *UpdateTableResponseBodyTaskInfo) GetTaskId() *string {
	return s.TaskId
}

func (s *UpdateTableResponseBodyTaskInfo) SetContent(v string) *UpdateTableResponseBodyTaskInfo {
	s.Content = &v
	return s
}

func (s *UpdateTableResponseBodyTaskInfo) SetNextTaskId(v string) *UpdateTableResponseBodyTaskInfo {
	s.NextTaskId = &v
	return s
}

func (s *UpdateTableResponseBodyTaskInfo) SetStatus(v string) *UpdateTableResponseBodyTaskInfo {
	s.Status = &v
	return s
}

func (s *UpdateTableResponseBodyTaskInfo) SetTaskId(v string) *UpdateTableResponseBodyTaskInfo {
	s.TaskId = &v
	return s
}

func (s *UpdateTableResponseBodyTaskInfo) Validate() error {
	return dara.Validate(s)
}
