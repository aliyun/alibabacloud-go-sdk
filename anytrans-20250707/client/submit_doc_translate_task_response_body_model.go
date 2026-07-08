// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitDocTranslateTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubmitDocTranslateTaskResponseBody
	GetCode() *string
	SetData(v *SubmitDocTranslateTaskResponseBodyData) *SubmitDocTranslateTaskResponseBody
	GetData() *SubmitDocTranslateTaskResponseBodyData
	SetHttpStatusCode(v string) *SubmitDocTranslateTaskResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *SubmitDocTranslateTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitDocTranslateTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitDocTranslateTaskResponseBody
	GetSuccess() *bool
}

type SubmitDocTranslateTaskResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response data.
	Data *SubmitDocTranslateTaskResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// The request ID, used to trace the API call.
	//
	// example:
	//
	// 377A48D7-7CFA-53F9-8CA2-14FE3F2774B6
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SubmitDocTranslateTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitDocTranslateTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitDocTranslateTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitDocTranslateTaskResponseBody) GetData() *SubmitDocTranslateTaskResponseBodyData {
	return s.Data
}

func (s *SubmitDocTranslateTaskResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *SubmitDocTranslateTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitDocTranslateTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitDocTranslateTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitDocTranslateTaskResponseBody) SetCode(v string) *SubmitDocTranslateTaskResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitDocTranslateTaskResponseBody) SetData(v *SubmitDocTranslateTaskResponseBodyData) *SubmitDocTranslateTaskResponseBody {
	s.Data = v
	return s
}

func (s *SubmitDocTranslateTaskResponseBody) SetHttpStatusCode(v string) *SubmitDocTranslateTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SubmitDocTranslateTaskResponseBody) SetMessage(v string) *SubmitDocTranslateTaskResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitDocTranslateTaskResponseBody) SetRequestId(v string) *SubmitDocTranslateTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitDocTranslateTaskResponseBody) SetSuccess(v bool) *SubmitDocTranslateTaskResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitDocTranslateTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitDocTranslateTaskResponseBodyData struct {
	// The task status.
	//
	// example:
	//
	// ready
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The document translation task ID.
	//
	// example:
	//
	// d3a2397bc2c14ab4a2e40a4f5b46241b
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// User-defined passthrough data. The translation service returns this data without processing it. This is useful for analytics tracking.
	//
	// example:
	//
	// {"traceId":"trace_123456"}
	TrackingData *string `json:"trackingData,omitempty" xml:"trackingData,omitempty"`
}

func (s SubmitDocTranslateTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SubmitDocTranslateTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitDocTranslateTaskResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *SubmitDocTranslateTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *SubmitDocTranslateTaskResponseBodyData) GetTrackingData() *string {
	return s.TrackingData
}

func (s *SubmitDocTranslateTaskResponseBodyData) SetStatus(v string) *SubmitDocTranslateTaskResponseBodyData {
	s.Status = &v
	return s
}

func (s *SubmitDocTranslateTaskResponseBodyData) SetTaskId(v string) *SubmitDocTranslateTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *SubmitDocTranslateTaskResponseBodyData) SetTrackingData(v string) *SubmitDocTranslateTaskResponseBodyData {
	s.TrackingData = &v
	return s
}

func (s *SubmitDocTranslateTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
