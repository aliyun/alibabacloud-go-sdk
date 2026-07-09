// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImageDetectionTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateImageDetectionTaskResponseBody
	GetCode() *string
	SetHttpStatusCode(v string) *CreateImageDetectionTaskResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *CreateImageDetectionTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateImageDetectionTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateImageDetectionTaskResponseBody
	GetSuccess() *bool
	SetTaskId(v string) *CreateImageDetectionTaskResponseBody
	GetTaskId() *string
}

type CreateImageDetectionTaskResponseBody struct {
	// The business error code. The value `OK` is returned if the request was successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code. The value `200` is returned if the request was successful.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The additional information. The value `success` is returned if the request was successful.
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
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The task ID. You can use this ID to call `GetImageDetectionTaskResult` to query the result.
	//
	// example:
	//
	// f47ac10b-58cc-4372-a567-0e02b2c3d479
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateImageDetectionTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateImageDetectionTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateImageDetectionTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateImageDetectionTaskResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *CreateImageDetectionTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateImageDetectionTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateImageDetectionTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateImageDetectionTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateImageDetectionTaskResponseBody) SetCode(v string) *CreateImageDetectionTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateImageDetectionTaskResponseBody) SetHttpStatusCode(v string) *CreateImageDetectionTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateImageDetectionTaskResponseBody) SetMessage(v string) *CreateImageDetectionTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateImageDetectionTaskResponseBody) SetRequestId(v string) *CreateImageDetectionTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateImageDetectionTaskResponseBody) SetSuccess(v bool) *CreateImageDetectionTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreateImageDetectionTaskResponseBody) SetTaskId(v string) *CreateImageDetectionTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateImageDetectionTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
