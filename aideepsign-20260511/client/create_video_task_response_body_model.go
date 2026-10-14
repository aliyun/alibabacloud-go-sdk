// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVideoTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateVideoTaskResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *CreateVideoTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateVideoTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateVideoTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateVideoTaskResponseBody
	GetSuccess() *bool
	SetTaskId(v string) *CreateVideoTaskResponseBody
	GetTaskId() *string
}

type CreateVideoTaskResponseBody struct {
	// The return code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The **HTTP status code**. A value of 200 indicates that the request is successful.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The additional information. A value of success indicates that the request is successful.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// A1B2C3D4-E5F6-7890-ABCD-EF1234567890
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. Valid values:
	//
	// - true: The call is successful.
	//
	// - false: The call fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The task ID, which is used to query the generation result.
	//
	// example:
	//
	// f47ac10b-58cc-4372-a567-0e02b2c3d479
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateVideoTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVideoTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVideoTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateVideoTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateVideoTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateVideoTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVideoTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateVideoTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateVideoTaskResponseBody) SetCode(v string) *CreateVideoTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateVideoTaskResponseBody) SetHttpStatusCode(v int32) *CreateVideoTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateVideoTaskResponseBody) SetMessage(v string) *CreateVideoTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateVideoTaskResponseBody) SetRequestId(v string) *CreateVideoTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVideoTaskResponseBody) SetSuccess(v bool) *CreateVideoTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreateVideoTaskResponseBody) SetTaskId(v string) *CreateVideoTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateVideoTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
