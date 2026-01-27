// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyImportTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyImportTaskResponseBody
	GetRequestId() *string
	SetStatus(v string) *ModifyImportTaskResponseBody
	GetStatus() *string
	SetTaskId(v int64) *ModifyImportTaskResponseBody
	GetTaskId() *int64
	SetTaskName(v string) *ModifyImportTaskResponseBody
	GetTaskName() *string
}

type ModifyImportTaskResponseBody struct {
	// example:
	//
	// 069EB9B1-DE12-54B9-8C20-822****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// IMPORTING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 41698****
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// task_1234
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s ModifyImportTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyImportTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyImportTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyImportTaskResponseBody) GetStatus() *string {
	return s.Status
}

func (s *ModifyImportTaskResponseBody) GetTaskId() *int64 {
	return s.TaskId
}

func (s *ModifyImportTaskResponseBody) GetTaskName() *string {
	return s.TaskName
}

func (s *ModifyImportTaskResponseBody) SetRequestId(v string) *ModifyImportTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyImportTaskResponseBody) SetStatus(v string) *ModifyImportTaskResponseBody {
	s.Status = &v
	return s
}

func (s *ModifyImportTaskResponseBody) SetTaskId(v int64) *ModifyImportTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *ModifyImportTaskResponseBody) SetTaskName(v string) *ModifyImportTaskResponseBody {
	s.TaskName = &v
	return s
}

func (s *ModifyImportTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
