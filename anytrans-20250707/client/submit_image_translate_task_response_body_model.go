// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitImageTranslateTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubmitImageTranslateTaskResponseBody
	GetCode() *string
	SetData(v *SubmitImageTranslateTaskResponseBodyData) *SubmitImageTranslateTaskResponseBody
	GetData() *SubmitImageTranslateTaskResponseBodyData
	SetHttpStatusCode(v string) *SubmitImageTranslateTaskResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *SubmitImageTranslateTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitImageTranslateTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitImageTranslateTaskResponseBody
	GetSuccess() *bool
}

type SubmitImageTranslateTaskResponseBody struct {
	// The response error code.
	//
	// example:
	//
	// success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	Data *SubmitImageTranslateTaskResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID, used for tracing API calls.
	//
	// example:
	//
	// 42FF90E5-5D40-5797-AAF6-8A4D837CCCD5
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SubmitImageTranslateTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitImageTranslateTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitImageTranslateTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitImageTranslateTaskResponseBody) GetData() *SubmitImageTranslateTaskResponseBodyData {
	return s.Data
}

func (s *SubmitImageTranslateTaskResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *SubmitImageTranslateTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitImageTranslateTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitImageTranslateTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitImageTranslateTaskResponseBody) SetCode(v string) *SubmitImageTranslateTaskResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitImageTranslateTaskResponseBody) SetData(v *SubmitImageTranslateTaskResponseBodyData) *SubmitImageTranslateTaskResponseBody {
	s.Data = v
	return s
}

func (s *SubmitImageTranslateTaskResponseBody) SetHttpStatusCode(v string) *SubmitImageTranslateTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SubmitImageTranslateTaskResponseBody) SetMessage(v string) *SubmitImageTranslateTaskResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitImageTranslateTaskResponseBody) SetRequestId(v string) *SubmitImageTranslateTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitImageTranslateTaskResponseBody) SetSuccess(v bool) *SubmitImageTranslateTaskResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitImageTranslateTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitImageTranslateTaskResponseBodyData struct {
	// The status of the translation task.
	//
	// example:
	//
	// success
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The ID of the image translation task.
	//
	// example:
	//
	// 2746f4be-cff2-465e-a2c6-12bff30ce0f9
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// User-defined passthrough data returned unmodified in the response, which is useful for tracking purposes.
	//
	// example:
	//
	// {"traceId":"trace_123456"}
	TrackingData *string `json:"trackingData,omitempty" xml:"trackingData,omitempty"`
}

func (s SubmitImageTranslateTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SubmitImageTranslateTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitImageTranslateTaskResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *SubmitImageTranslateTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *SubmitImageTranslateTaskResponseBodyData) GetTrackingData() *string {
	return s.TrackingData
}

func (s *SubmitImageTranslateTaskResponseBodyData) SetStatus(v string) *SubmitImageTranslateTaskResponseBodyData {
	s.Status = &v
	return s
}

func (s *SubmitImageTranslateTaskResponseBodyData) SetTaskId(v string) *SubmitImageTranslateTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *SubmitImageTranslateTaskResponseBodyData) SetTrackingData(v string) *SubmitImageTranslateTaskResponseBodyData {
	s.TrackingData = &v
	return s
}

func (s *SubmitImageTranslateTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
