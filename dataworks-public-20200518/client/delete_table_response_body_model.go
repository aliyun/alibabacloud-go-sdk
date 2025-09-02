// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteTableResponseBody
	GetRequestId() *string
	SetTaskInfo(v *DeleteTableResponseBodyTaskInfo) *DeleteTableResponseBody
	GetTaskInfo() *DeleteTableResponseBodyTaskInfo
}

type DeleteTableResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// abcde
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the task that is used to delete the table.
	TaskInfo *DeleteTableResponseBodyTaskInfo `json:"TaskInfo,omitempty" xml:"TaskInfo,omitempty" type:"Struct"`
}

func (s DeleteTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTableResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTableResponseBody) GetTaskInfo() *DeleteTableResponseBodyTaskInfo {
	return s.TaskInfo
}

func (s *DeleteTableResponseBody) SetRequestId(v string) *DeleteTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTableResponseBody) SetTaskInfo(v *DeleteTableResponseBodyTaskInfo) *DeleteTableResponseBody {
	s.TaskInfo = v
	return s
}

func (s *DeleteTableResponseBody) Validate() error {
	return dara.Validate(s)
}

type DeleteTableResponseBodyTaskInfo struct {
	// The content of the task.
	//
	// example:
	//
	// success
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The ID of the task that is running.
	//
	// example:
	//
	// abc1
	NextTaskId *string `json:"NextTaskId,omitempty" xml:"NextTaskId,omitempty"`
	// The status of the task that is complete.
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the task that is complete.
	//
	// example:
	//
	// abc
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DeleteTableResponseBodyTaskInfo) String() string {
	return dara.Prettify(s)
}

func (s DeleteTableResponseBodyTaskInfo) GoString() string {
	return s.String()
}

func (s *DeleteTableResponseBodyTaskInfo) GetContent() *string {
	return s.Content
}

func (s *DeleteTableResponseBodyTaskInfo) GetNextTaskId() *string {
	return s.NextTaskId
}

func (s *DeleteTableResponseBodyTaskInfo) GetStatus() *string {
	return s.Status
}

func (s *DeleteTableResponseBodyTaskInfo) GetTaskId() *string {
	return s.TaskId
}

func (s *DeleteTableResponseBodyTaskInfo) SetContent(v string) *DeleteTableResponseBodyTaskInfo {
	s.Content = &v
	return s
}

func (s *DeleteTableResponseBodyTaskInfo) SetNextTaskId(v string) *DeleteTableResponseBodyTaskInfo {
	s.NextTaskId = &v
	return s
}

func (s *DeleteTableResponseBodyTaskInfo) SetStatus(v string) *DeleteTableResponseBodyTaskInfo {
	s.Status = &v
	return s
}

func (s *DeleteTableResponseBodyTaskInfo) SetTaskId(v string) *DeleteTableResponseBodyTaskInfo {
	s.TaskId = &v
	return s
}

func (s *DeleteTableResponseBodyTaskInfo) Validate() error {
	return dara.Validate(s)
}
