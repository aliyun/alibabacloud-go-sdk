// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWatermarkExtractTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateWatermarkExtractTaskResponseBody
	GetCode() *string
	SetHttpStatusCode(v string) *CreateWatermarkExtractTaskResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *CreateWatermarkExtractTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateWatermarkExtractTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateWatermarkExtractTaskResponseBody
	GetSuccess() *bool
	SetTaskId(v string) *CreateWatermarkExtractTaskResponseBody
	GetTaskId() *string
}

type CreateWatermarkExtractTaskResponseBody struct {
	// The business error code. OK is returned if the request is successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code. `200` is returned if the request is successful.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The additional information. "success" is returned if the request is successful.
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
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The dark watermark extraction task ID, which is used to call GetWatermarkExtractTaskResult to query the result.
	//
	// example:
	//
	// f47ac10b-58cc-4372-a567-0e02b2c3d479
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateWatermarkExtractTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateWatermarkExtractTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWatermarkExtractTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateWatermarkExtractTaskResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *CreateWatermarkExtractTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateWatermarkExtractTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateWatermarkExtractTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateWatermarkExtractTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateWatermarkExtractTaskResponseBody) SetCode(v string) *CreateWatermarkExtractTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateWatermarkExtractTaskResponseBody) SetHttpStatusCode(v string) *CreateWatermarkExtractTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateWatermarkExtractTaskResponseBody) SetMessage(v string) *CreateWatermarkExtractTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateWatermarkExtractTaskResponseBody) SetRequestId(v string) *CreateWatermarkExtractTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWatermarkExtractTaskResponseBody) SetSuccess(v bool) *CreateWatermarkExtractTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreateWatermarkExtractTaskResponseBody) SetTaskId(v string) *CreateWatermarkExtractTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateWatermarkExtractTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
