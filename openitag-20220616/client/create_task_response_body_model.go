// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CreateTaskResponseBody
	GetCode() *int32
	SetDetails(v string) *CreateTaskResponseBody
	GetDetails() *string
	SetErrorCode(v string) *CreateTaskResponseBody
	GetErrorCode() *string
	SetMessage(v string) *CreateTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateTaskResponseBody
	GetSuccess() *bool
	SetTaskId(v string) *CreateTaskResponseBody
	GetTaskId() *string
}

type CreateTaskResponseBody struct {
	// Total amount of data under the conditions of this request. This parameter is optional and does not need to be returned by default.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Details
	//
	// example:
	//
	// null
	Details *string `json:"Details,omitempty" xml:"Details,omitempty"`
	// Error code
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// Response message of the request
	//
	// This parameter is required.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// 90ABA848-AD74-1F6E-84BC-4182A7F1F29E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation succeeded.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Job ID
	//
	// example:
	//
	// 154***2518306500608
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTaskResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateTaskResponseBody) GetDetails() *string {
	return s.Details
}

func (s *CreateTaskResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateTaskResponseBody) SetCode(v int32) *CreateTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateTaskResponseBody) SetDetails(v string) *CreateTaskResponseBody {
	s.Details = &v
	return s
}

func (s *CreateTaskResponseBody) SetErrorCode(v string) *CreateTaskResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateTaskResponseBody) SetMessage(v string) *CreateTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateTaskResponseBody) SetRequestId(v string) *CreateTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTaskResponseBody) SetSuccess(v bool) *CreateTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreateTaskResponseBody) SetTaskId(v string) *CreateTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
