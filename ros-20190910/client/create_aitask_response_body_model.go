// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAITaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateAITaskResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *CreateAITaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateAITaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateAITaskResponseBody
	GetRequestId() *string
	SetSuccess(v string) *CreateAITaskResponseBody
	GetSuccess() *string
	SetTaskId(v string) *CreateAITaskResponseBody
	GetTaskId() *string
}

type CreateAITaskResponseBody struct {
	// Error code.
	//
	// example:
	//
	// Forbidden
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Error message.
	//
	// example:
	//
	// You are not authorized to complete this action.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// B288A0BE-D927-4888-B0F7-B35EF84B6E6F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Values:
	//
	// - true: Call succeeded.
	//
	// - false: Call failed.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// AI task ID.
	//
	// example:
	//
	// t-asas*****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateAITaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAITaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAITaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateAITaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateAITaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateAITaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAITaskResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *CreateAITaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateAITaskResponseBody) SetCode(v string) *CreateAITaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAITaskResponseBody) SetHttpStatusCode(v int32) *CreateAITaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateAITaskResponseBody) SetMessage(v string) *CreateAITaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAITaskResponseBody) SetRequestId(v string) *CreateAITaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAITaskResponseBody) SetSuccess(v string) *CreateAITaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreateAITaskResponseBody) SetTaskId(v string) *CreateAITaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateAITaskResponseBody) Validate() error {
	return dara.Validate(s)
}
