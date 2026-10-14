// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImageIncomeTaxDetectionTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateImageIncomeTaxDetectionTaskResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *CreateImageIncomeTaxDetectionTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateImageIncomeTaxDetectionTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateImageIncomeTaxDetectionTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateImageIncomeTaxDetectionTaskResponseBody
	GetSuccess() *bool
	SetTaskId(v string) *CreateImageIncomeTaxDetectionTaskResponseBody
	GetTaskId() *string
}

type CreateImageIncomeTaxDetectionTaskResponseBody struct {
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
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
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
	// Indicates whether the call was successful. Valid values:
	//
	// - true: The call was successful.
	//
	// - false: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The income tax screenshot detection task ID. Use this ID to query the result by calling GetImageIncomeTaxDetectionTaskResult.
	//
	// example:
	//
	// f47ac10b-58cc-4372-a567-0e02b2c3d479
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateImageIncomeTaxDetectionTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateImageIncomeTaxDetectionTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateImageIncomeTaxDetectionTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateImageIncomeTaxDetectionTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateImageIncomeTaxDetectionTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateImageIncomeTaxDetectionTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateImageIncomeTaxDetectionTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateImageIncomeTaxDetectionTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateImageIncomeTaxDetectionTaskResponseBody) SetCode(v string) *CreateImageIncomeTaxDetectionTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateImageIncomeTaxDetectionTaskResponseBody) SetHttpStatusCode(v int32) *CreateImageIncomeTaxDetectionTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateImageIncomeTaxDetectionTaskResponseBody) SetMessage(v string) *CreateImageIncomeTaxDetectionTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateImageIncomeTaxDetectionTaskResponseBody) SetRequestId(v string) *CreateImageIncomeTaxDetectionTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateImageIncomeTaxDetectionTaskResponseBody) SetSuccess(v bool) *CreateImageIncomeTaxDetectionTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreateImageIncomeTaxDetectionTaskResponseBody) SetTaskId(v string) *CreateImageIncomeTaxDetectionTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateImageIncomeTaxDetectionTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
