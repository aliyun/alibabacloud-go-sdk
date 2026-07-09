// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImageTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateImageTaskResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *CreateImageTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateImageTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateImageTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateImageTaskResponseBody
	GetSuccess() *bool
	SetTaskId(v string) *CreateImageTaskResponseBody
	GetTaskId() *string
}

type CreateImageTaskResponseBody struct {
	// The business error code. The value `OK` is returned if the request succeeds.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code. The value `200` is returned if the request succeeds.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The additional information. The value `success` is returned if the request succeeds.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A1B2C3D4-E5F6-7890-ABCD-EF1234567890
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values: true and false.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The task ID. Use this ID to call the GetImageTaskResult operation to query the results.
	//
	// example:
	//
	// f47ac10b-58cc-4372-a567-0e02b2c3d479
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateImageTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateImageTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateImageTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateImageTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateImageTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateImageTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateImageTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateImageTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateImageTaskResponseBody) SetCode(v string) *CreateImageTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateImageTaskResponseBody) SetHttpStatusCode(v int32) *CreateImageTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateImageTaskResponseBody) SetMessage(v string) *CreateImageTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateImageTaskResponseBody) SetRequestId(v string) *CreateImageTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateImageTaskResponseBody) SetSuccess(v bool) *CreateImageTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreateImageTaskResponseBody) SetTaskId(v string) *CreateImageTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateImageTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
