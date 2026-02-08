// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTaskExportTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateTaskExportTaskResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *CreateTaskExportTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateTaskExportTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateTaskExportTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateTaskExportTaskResponseBody
	GetSuccess() *bool
	SetTaskId(v string) *CreateTaskExportTaskResponseBody
	GetTaskId() *string
}

type CreateTaskExportTaskResponseBody struct {
	// Status code
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// API message
	//
	// example:
	//
	// []
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// C377C5FF-4F94-1B23-89D0-50C560623EE4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation succeeded
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Export job ID
	//
	// example:
	//
	// a7d6dcff1b8b40f4a8b769a9c24e7852
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateTaskExportTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskExportTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTaskExportTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateTaskExportTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateTaskExportTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateTaskExportTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTaskExportTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateTaskExportTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateTaskExportTaskResponseBody) SetCode(v string) *CreateTaskExportTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateTaskExportTaskResponseBody) SetHttpStatusCode(v int32) *CreateTaskExportTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateTaskExportTaskResponseBody) SetMessage(v string) *CreateTaskExportTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateTaskExportTaskResponseBody) SetRequestId(v string) *CreateTaskExportTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTaskExportTaskResponseBody) SetSuccess(v bool) *CreateTaskExportTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreateTaskExportTaskResponseBody) SetTaskId(v string) *CreateTaskExportTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateTaskExportTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
