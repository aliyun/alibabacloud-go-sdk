// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTaskPreparingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *TaskPreparingResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *TaskPreparingResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *TaskPreparingResponseBody
	GetMessage() *string
	SetRequestId(v string) *TaskPreparingResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TaskPreparingResponseBody
	GetSuccess() *bool
	SetTaskId(v string) *TaskPreparingResponseBody
	GetTaskId() *string
}

type TaskPreparingResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Succes
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// e4e2a770-b97b-465a-80d8-06dca008c503
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s TaskPreparingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TaskPreparingResponseBody) GoString() string {
	return s.String()
}

func (s *TaskPreparingResponseBody) GetCode() *string {
	return s.Code
}

func (s *TaskPreparingResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *TaskPreparingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TaskPreparingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TaskPreparingResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TaskPreparingResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *TaskPreparingResponseBody) SetCode(v string) *TaskPreparingResponseBody {
	s.Code = &v
	return s
}

func (s *TaskPreparingResponseBody) SetHttpStatusCode(v int32) *TaskPreparingResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *TaskPreparingResponseBody) SetMessage(v string) *TaskPreparingResponseBody {
	s.Message = &v
	return s
}

func (s *TaskPreparingResponseBody) SetRequestId(v string) *TaskPreparingResponseBody {
	s.RequestId = &v
	return s
}

func (s *TaskPreparingResponseBody) SetSuccess(v bool) *TaskPreparingResponseBody {
	s.Success = &v
	return s
}

func (s *TaskPreparingResponseBody) SetTaskId(v string) *TaskPreparingResponseBody {
	s.TaskId = &v
	return s
}

func (s *TaskPreparingResponseBody) Validate() error {
	return dara.Validate(s)
}
