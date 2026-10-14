// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWatermarkEmbedTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateWatermarkEmbedTaskResponseBody
	GetCode() *string
	SetHttpStatusCode(v string) *CreateWatermarkEmbedTaskResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *CreateWatermarkEmbedTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateWatermarkEmbedTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateWatermarkEmbedTaskResponseBody
	GetSuccess() *bool
	SetTaskId(v string) *CreateWatermarkEmbedTaskResponseBody
	GetTaskId() *string
}

type CreateWatermarkEmbedTaskResponseBody struct {
	// The business error code. The value OK is returned if the request is successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code. The value 200 is returned if the request is successful.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The response message.
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
	// The invisible watermark embedding task ID. Use this ID to call GetWatermarkEmbedTaskResult to query the result.
	//
	// example:
	//
	// f47ac10b-58cc-4372-a567-0e02b2c3d479
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateWatermarkEmbedTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateWatermarkEmbedTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWatermarkEmbedTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateWatermarkEmbedTaskResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *CreateWatermarkEmbedTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateWatermarkEmbedTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateWatermarkEmbedTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateWatermarkEmbedTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateWatermarkEmbedTaskResponseBody) SetCode(v string) *CreateWatermarkEmbedTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateWatermarkEmbedTaskResponseBody) SetHttpStatusCode(v string) *CreateWatermarkEmbedTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateWatermarkEmbedTaskResponseBody) SetMessage(v string) *CreateWatermarkEmbedTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateWatermarkEmbedTaskResponseBody) SetRequestId(v string) *CreateWatermarkEmbedTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWatermarkEmbedTaskResponseBody) SetSuccess(v bool) *CreateWatermarkEmbedTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreateWatermarkEmbedTaskResponseBody) SetTaskId(v string) *CreateWatermarkEmbedTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateWatermarkEmbedTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
