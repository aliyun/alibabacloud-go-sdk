// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTraceAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *DeleteTraceAppResponseBody
	GetCode() *int64
	SetData(v string) *DeleteTraceAppResponseBody
	GetData() *string
	SetMessage(v string) *DeleteTraceAppResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteTraceAppResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteTraceAppResponseBody
	GetSuccess() *bool
}

type DeleteTraceAppResponseBody struct {
	// The status code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response in JSON format, including the HTTP status code, error code, response message, and trace ID.
	//
	// example:
	//
	// "{\\\\"code\\\\":200,\\\\"data\\\\":\\\\"{\\\\\\\\\\"code\\\\\\\\\\":200,\\\\\\\\\\"data\\\\\\\\\\":true,\\\\\\\\\\"errorCode\\\\\\\\\\":\\\\\\\\\\"The application is deleted\\\\\\\\\\",\\\\\\\\\\"message\\\\\\\\\\":\\\\\\\\\\"The application is deleted\\\\\\\\\\",\\\\\\\\\\"success\\\\\\\\\\":true,\\\\\\\\\\"traceId\\\\\\\\\\":\\\\\\\\\\"0bc0594d15954826692915817e\\*\\*\\*\\*\\\\\\\\\\"}\\\\",\\\\"errorCode\\\\":\\\\"The application is deleted\\\\",\\\\"message\\\\":\\\\"The application is deleted\\\\",\\\\"success\\\\":true,\\\\"traceId\\\\":\\\\"0ab2646915954826692568137d\\*\\*\\*\\*\\\\"}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// example:
	//
	// Internal error. Please try again. Contact the DingTalk service account if the issue                              persists after multiple retries.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 46355DD8-FC56-40C5-BFC6-269DE4F9****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// true: The request was successful.
	//
	// false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteTraceAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTraceAppResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTraceAppResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *DeleteTraceAppResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteTraceAppResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteTraceAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTraceAppResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteTraceAppResponseBody) SetCode(v int64) *DeleteTraceAppResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteTraceAppResponseBody) SetData(v string) *DeleteTraceAppResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteTraceAppResponseBody) SetMessage(v string) *DeleteTraceAppResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteTraceAppResponseBody) SetRequestId(v string) *DeleteTraceAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTraceAppResponseBody) SetSuccess(v bool) *DeleteTraceAppResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteTraceAppResponseBody) Validate() error {
	return dara.Validate(s)
}
