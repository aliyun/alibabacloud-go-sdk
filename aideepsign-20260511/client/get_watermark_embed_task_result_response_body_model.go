// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWatermarkEmbedTaskResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetWatermarkEmbedTaskResultResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetWatermarkEmbedTaskResultResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetWatermarkEmbedTaskResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetWatermarkEmbedTaskResultResponseBody
	GetRequestId() *string
	SetStatus(v string) *GetWatermarkEmbedTaskResultResponseBody
	GetStatus() *string
	SetSuccess(v bool) *GetWatermarkEmbedTaskResultResponseBody
	GetSuccess() *bool
	SetTaskId(v string) *GetWatermarkEmbedTaskResultResponseBody
	GetTaskId() *string
	SetUrl(v string) *GetWatermarkEmbedTaskResultResponseBody
	GetUrl() *string
}

type GetWatermarkEmbedTaskResultResponseBody struct {
	// **The business error code.*	- The value OK is returned if the request is successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// **The HTTP status code.*	- The value 200 is returned if the request is successful.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The failure reason when the task fails. This parameter is returned only when Status is failed.
	//
	// example:
	//
	// poll timeout: no result within 10m0s
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A1B2C3D4-E5F6-7890-ABCD-EF1234567890
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task status. The processing time is approximately equal to the media duration. Valid values:
	//
	// - pending: The task is waiting.
	//
	// - running: The task is running.
	//
	// - succeeded: The task succeeded.
	//
	// - failed: The task failed.
	//
	// example:
	//
	// succeeded
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The invisible watermark embedding task ID.
	//
	// example:
	//
	// f47ac10b-58cc-4372-a567-0e02b2c3d479
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The pre-signed URL for accessing the watermarked file. This parameter is returned only when Status is succeeded.
	//
	// example:
	//
	// https://oss-bucket.example.com/aideepsign/...
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetWatermarkEmbedTaskResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWatermarkEmbedTaskResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetWatermarkEmbedTaskResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetWatermarkEmbedTaskResultResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetWatermarkEmbedTaskResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetWatermarkEmbedTaskResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWatermarkEmbedTaskResultResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetWatermarkEmbedTaskResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetWatermarkEmbedTaskResultResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *GetWatermarkEmbedTaskResultResponseBody) GetUrl() *string {
	return s.Url
}

func (s *GetWatermarkEmbedTaskResultResponseBody) SetCode(v string) *GetWatermarkEmbedTaskResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetWatermarkEmbedTaskResultResponseBody) SetHttpStatusCode(v int32) *GetWatermarkEmbedTaskResultResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetWatermarkEmbedTaskResultResponseBody) SetMessage(v string) *GetWatermarkEmbedTaskResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetWatermarkEmbedTaskResultResponseBody) SetRequestId(v string) *GetWatermarkEmbedTaskResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWatermarkEmbedTaskResultResponseBody) SetStatus(v string) *GetWatermarkEmbedTaskResultResponseBody {
	s.Status = &v
	return s
}

func (s *GetWatermarkEmbedTaskResultResponseBody) SetSuccess(v bool) *GetWatermarkEmbedTaskResultResponseBody {
	s.Success = &v
	return s
}

func (s *GetWatermarkEmbedTaskResultResponseBody) SetTaskId(v string) *GetWatermarkEmbedTaskResultResponseBody {
	s.TaskId = &v
	return s
}

func (s *GetWatermarkEmbedTaskResultResponseBody) SetUrl(v string) *GetWatermarkEmbedTaskResultResponseBody {
	s.Url = &v
	return s
}

func (s *GetWatermarkEmbedTaskResultResponseBody) Validate() error {
	return dara.Validate(s)
}
