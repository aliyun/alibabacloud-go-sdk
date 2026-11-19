// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryMemoryTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RetryMemoryTaskResponseBody
	GetCode() *string
	SetData(v *RetryMemoryTaskResponseBodyData) *RetryMemoryTaskResponseBody
	GetData() *RetryMemoryTaskResponseBodyData
	SetHttpStatusCode(v int32) *RetryMemoryTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *RetryMemoryTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *RetryMemoryTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RetryMemoryTaskResponseBody
	GetSuccess() *bool
}

type RetryMemoryTaskResponseBody struct {
	// code
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// data
	Data *RetryMemoryTaskResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// httpStatusCode
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// message
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// requestId
	//
	// example:
	//
	// 01A00000-0000-0000-0000-000000000001
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// success
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s RetryMemoryTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RetryMemoryTaskResponseBody) GoString() string {
	return s.String()
}

func (s *RetryMemoryTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *RetryMemoryTaskResponseBody) GetData() *RetryMemoryTaskResponseBodyData {
	return s.Data
}

func (s *RetryMemoryTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *RetryMemoryTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RetryMemoryTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RetryMemoryTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RetryMemoryTaskResponseBody) SetCode(v string) *RetryMemoryTaskResponseBody {
	s.Code = &v
	return s
}

func (s *RetryMemoryTaskResponseBody) SetData(v *RetryMemoryTaskResponseBodyData) *RetryMemoryTaskResponseBody {
	s.Data = v
	return s
}

func (s *RetryMemoryTaskResponseBody) SetHttpStatusCode(v int32) *RetryMemoryTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RetryMemoryTaskResponseBody) SetMessage(v string) *RetryMemoryTaskResponseBody {
	s.Message = &v
	return s
}

func (s *RetryMemoryTaskResponseBody) SetRequestId(v string) *RetryMemoryTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *RetryMemoryTaskResponseBody) SetSuccess(v bool) *RetryMemoryTaskResponseBody {
	s.Success = &v
	return s
}

func (s *RetryMemoryTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RetryMemoryTaskResponseBodyData struct {
	// status
	//
	// This parameter is required.
	//
	// example:
	//
	// SUCCEEDED
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// taskId
	//
	// This parameter is required.
	//
	// example:
	//
	// task-2f5f9c5a8d0e4c3b
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s RetryMemoryTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RetryMemoryTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *RetryMemoryTaskResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *RetryMemoryTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *RetryMemoryTaskResponseBodyData) SetStatus(v string) *RetryMemoryTaskResponseBodyData {
	s.Status = &v
	return s
}

func (s *RetryMemoryTaskResponseBodyData) SetTaskId(v string) *RetryMemoryTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *RetryMemoryTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
