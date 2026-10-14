// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWatermarkExtractTaskResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAvgAcc(v float64) *GetWatermarkExtractTaskResultResponseBody
	GetAvgAcc() *float64
	SetCode(v string) *GetWatermarkExtractTaskResultResponseBody
	GetCode() *string
	SetDetected(v bool) *GetWatermarkExtractTaskResultResponseBody
	GetDetected() *bool
	SetHttpStatusCode(v int32) *GetWatermarkExtractTaskResultResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetWatermarkExtractTaskResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetWatermarkExtractTaskResultResponseBody
	GetRequestId() *string
	SetStatus(v string) *GetWatermarkExtractTaskResultResponseBody
	GetStatus() *string
	SetSuccess(v bool) *GetWatermarkExtractTaskResultResponseBody
	GetSuccess() *bool
	SetTaskId(v string) *GetWatermarkExtractTaskResultResponseBody
	GetTaskId() *string
	SetWatermarkContent(v string) *GetWatermarkExtractTaskResultResponseBody
	GetWatermarkContent() *string
}

type GetWatermarkExtractTaskResultResponseBody struct {
	// The average extraction confidence, ranging from 0 to 1. This parameter is returned only when Status is succeeded.
	//
	// example:
	//
	// 0.99
	AvgAcc *float64 `json:"AvgAcc,omitempty" xml:"AvgAcc,omitempty"`
	// The business error code. OK is returned if the request is successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether a watermark is detected.
	//
	// example:
	//
	// true
	Detected *bool `json:"Detected,omitempty" xml:"Detected,omitempty"`
	// The HTTP status code. 200 is returned if the request is successful.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The failure reason when the task fails. This parameter is returned only when Status is failed.
	//
	// example:
	//
	// The request has failed due to a temporary failure of the server.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// A1B2C3D4-E5F6-7890-ABCD-EF1234567890
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task status. Valid values:
	//
	// - running: In progress.
	//
	// - succeeded: Completed.
	//
	// - failed: Failed.
	//
	// example:
	//
	// succeeded
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The dark watermark extraction task ID.
	//
	// example:
	//
	// f47ac10b-58cc-4372-a567-0e02b2c3d479
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The extracted watermark content. This parameter is returned only when `Status` is `succeeded` and `Detected=true`.
	//
	// example:
	//
	// wm-e2e-85970647
	WatermarkContent *string `json:"WatermarkContent,omitempty" xml:"WatermarkContent,omitempty"`
}

func (s GetWatermarkExtractTaskResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWatermarkExtractTaskResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetWatermarkExtractTaskResultResponseBody) GetAvgAcc() *float64 {
	return s.AvgAcc
}

func (s *GetWatermarkExtractTaskResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetWatermarkExtractTaskResultResponseBody) GetDetected() *bool {
	return s.Detected
}

func (s *GetWatermarkExtractTaskResultResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetWatermarkExtractTaskResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetWatermarkExtractTaskResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWatermarkExtractTaskResultResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetWatermarkExtractTaskResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetWatermarkExtractTaskResultResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *GetWatermarkExtractTaskResultResponseBody) GetWatermarkContent() *string {
	return s.WatermarkContent
}

func (s *GetWatermarkExtractTaskResultResponseBody) SetAvgAcc(v float64) *GetWatermarkExtractTaskResultResponseBody {
	s.AvgAcc = &v
	return s
}

func (s *GetWatermarkExtractTaskResultResponseBody) SetCode(v string) *GetWatermarkExtractTaskResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetWatermarkExtractTaskResultResponseBody) SetDetected(v bool) *GetWatermarkExtractTaskResultResponseBody {
	s.Detected = &v
	return s
}

func (s *GetWatermarkExtractTaskResultResponseBody) SetHttpStatusCode(v int32) *GetWatermarkExtractTaskResultResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetWatermarkExtractTaskResultResponseBody) SetMessage(v string) *GetWatermarkExtractTaskResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetWatermarkExtractTaskResultResponseBody) SetRequestId(v string) *GetWatermarkExtractTaskResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWatermarkExtractTaskResultResponseBody) SetStatus(v string) *GetWatermarkExtractTaskResultResponseBody {
	s.Status = &v
	return s
}

func (s *GetWatermarkExtractTaskResultResponseBody) SetSuccess(v bool) *GetWatermarkExtractTaskResultResponseBody {
	s.Success = &v
	return s
}

func (s *GetWatermarkExtractTaskResultResponseBody) SetTaskId(v string) *GetWatermarkExtractTaskResultResponseBody {
	s.TaskId = &v
	return s
}

func (s *GetWatermarkExtractTaskResultResponseBody) SetWatermarkContent(v string) *GetWatermarkExtractTaskResultResponseBody {
	s.WatermarkContent = &v
	return s
}

func (s *GetWatermarkExtractTaskResultResponseBody) Validate() error {
	return dara.Validate(s)
}
