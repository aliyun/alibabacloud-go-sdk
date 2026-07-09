// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSensitiveScanTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateSensitiveScanTaskResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *CreateSensitiveScanTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateSensitiveScanTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateSensitiveScanTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateSensitiveScanTaskResponseBody
	GetSuccess() *bool
	SetTaskId(v string) *CreateSensitiveScanTaskResponseBody
	GetTaskId() *string
}

type CreateSensitiveScanTaskResponseBody struct {
	// The internal error code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The additional information.
	//
	// example:
	//
	// null
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A1B2C3D4-E5F6-7890-ABCD-EF1234567890
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// - true: The call was successful.
	//
	// - false: The call failed.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The task ID. You can use this ID to call the GetSensitiveScanResult operation to query the result.
	//
	// example:
	//
	// f47ac10b-58cc-4372-a567-0e02b2c3d479
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateSensitiveScanTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSensitiveScanTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSensitiveScanTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateSensitiveScanTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateSensitiveScanTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateSensitiveScanTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSensitiveScanTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateSensitiveScanTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateSensitiveScanTaskResponseBody) SetCode(v string) *CreateSensitiveScanTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateSensitiveScanTaskResponseBody) SetHttpStatusCode(v int32) *CreateSensitiveScanTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateSensitiveScanTaskResponseBody) SetMessage(v string) *CreateSensitiveScanTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateSensitiveScanTaskResponseBody) SetRequestId(v string) *CreateSensitiveScanTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSensitiveScanTaskResponseBody) SetSuccess(v bool) *CreateSensitiveScanTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreateSensitiveScanTaskResponseBody) SetTaskId(v string) *CreateSensitiveScanTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateSensitiveScanTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
