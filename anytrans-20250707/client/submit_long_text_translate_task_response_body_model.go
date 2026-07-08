// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitLongTextTranslateTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubmitLongTextTranslateTaskResponseBody
	GetCode() *string
	SetData(v *SubmitLongTextTranslateTaskResponseBodyData) *SubmitLongTextTranslateTaskResponseBody
	GetData() *SubmitLongTextTranslateTaskResponseBodyData
	SetHttpStatusCode(v string) *SubmitLongTextTranslateTaskResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *SubmitLongTextTranslateTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitLongTextTranslateTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitLongTextTranslateTaskResponseBody
	GetSuccess() *bool
}

type SubmitLongTextTranslateTaskResponseBody struct {
	// The result code of the API call.
	//
	// example:
	//
	// success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The object that contains the returned data.
	Data *SubmitLongTextTranslateTaskResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// A message that corresponds to the code.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The unique ID for the API request, used for tracing.
	//
	// example:
	//
	// C2D45266-3135-1A06-AD7F-69E782ED596F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the API call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SubmitLongTextTranslateTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitLongTextTranslateTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitLongTextTranslateTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitLongTextTranslateTaskResponseBody) GetData() *SubmitLongTextTranslateTaskResponseBodyData {
	return s.Data
}

func (s *SubmitLongTextTranslateTaskResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *SubmitLongTextTranslateTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitLongTextTranslateTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitLongTextTranslateTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitLongTextTranslateTaskResponseBody) SetCode(v string) *SubmitLongTextTranslateTaskResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitLongTextTranslateTaskResponseBody) SetData(v *SubmitLongTextTranslateTaskResponseBodyData) *SubmitLongTextTranslateTaskResponseBody {
	s.Data = v
	return s
}

func (s *SubmitLongTextTranslateTaskResponseBody) SetHttpStatusCode(v string) *SubmitLongTextTranslateTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SubmitLongTextTranslateTaskResponseBody) SetMessage(v string) *SubmitLongTextTranslateTaskResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitLongTextTranslateTaskResponseBody) SetRequestId(v string) *SubmitLongTextTranslateTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitLongTextTranslateTaskResponseBody) SetSuccess(v bool) *SubmitLongTextTranslateTaskResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitLongTextTranslateTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitLongTextTranslateTaskResponseBodyData struct {
	// The status of the translation task.
	//
	// example:
	//
	// 200
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The ID of the long-text translation task.
	//
	// example:
	//
	// a8f25f25-0b36-4349-857f-e19a43f69e51
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// A custom string passed from the request to the response unmodified. This is useful for tracking or correlating API calls.
	//
	// example:
	//
	// {"traceId":"trace_123456"}
	TrackingData *string `json:"trackingData,omitempty" xml:"trackingData,omitempty"`
}

func (s SubmitLongTextTranslateTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SubmitLongTextTranslateTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitLongTextTranslateTaskResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *SubmitLongTextTranslateTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *SubmitLongTextTranslateTaskResponseBodyData) GetTrackingData() *string {
	return s.TrackingData
}

func (s *SubmitLongTextTranslateTaskResponseBodyData) SetStatus(v string) *SubmitLongTextTranslateTaskResponseBodyData {
	s.Status = &v
	return s
}

func (s *SubmitLongTextTranslateTaskResponseBodyData) SetTaskId(v string) *SubmitLongTextTranslateTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *SubmitLongTextTranslateTaskResponseBodyData) SetTrackingData(v string) *SubmitLongTextTranslateTaskResponseBodyData {
	s.TrackingData = &v
	return s
}

func (s *SubmitLongTextTranslateTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
