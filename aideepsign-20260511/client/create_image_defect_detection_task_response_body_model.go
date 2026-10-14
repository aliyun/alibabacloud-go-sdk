// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImageDefectDetectionTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateImageDefectDetectionTaskResponseBody
	GetCode() *string
	SetHttpStatusCode(v string) *CreateImageDefectDetectionTaskResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *CreateImageDefectDetectionTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateImageDefectDetectionTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateImageDefectDetectionTaskResponseBody
	GetSuccess() *bool
	SetTaskId(v string) *CreateImageDefectDetectionTaskResponseBody
	GetTaskId() *string
}

type CreateImageDefectDetectionTaskResponseBody struct {
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
	// The defect detection task ID, used for subsequent queries through the GetImageDefectDetectionTaskResult operation.
	//
	// example:
	//
	// f47ac10b-58cc-4372-a567-0e02b2c3d479
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateImageDefectDetectionTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateImageDefectDetectionTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateImageDefectDetectionTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateImageDefectDetectionTaskResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *CreateImageDefectDetectionTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateImageDefectDetectionTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateImageDefectDetectionTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateImageDefectDetectionTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateImageDefectDetectionTaskResponseBody) SetCode(v string) *CreateImageDefectDetectionTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateImageDefectDetectionTaskResponseBody) SetHttpStatusCode(v string) *CreateImageDefectDetectionTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateImageDefectDetectionTaskResponseBody) SetMessage(v string) *CreateImageDefectDetectionTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateImageDefectDetectionTaskResponseBody) SetRequestId(v string) *CreateImageDefectDetectionTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateImageDefectDetectionTaskResponseBody) SetSuccess(v bool) *CreateImageDefectDetectionTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreateImageDefectDetectionTaskResponseBody) SetTaskId(v string) *CreateImageDefectDetectionTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateImageDefectDetectionTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
