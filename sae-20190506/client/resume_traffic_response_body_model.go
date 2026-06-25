// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeTrafficResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ResumeTrafficResponseBody
	GetCode() *string
	SetData(v *ResumeTrafficResponseBodyData) *ResumeTrafficResponseBody
	GetData() *ResumeTrafficResponseBodyData
	SetErrorCode(v string) *ResumeTrafficResponseBody
	GetErrorCode() *string
	SetMessage(v string) *ResumeTrafficResponseBody
	GetMessage() *string
	SetRequestId(v string) *ResumeTrafficResponseBody
	GetRequestId() *string
	SetSuccess(v string) *ResumeTrafficResponseBody
	GetSuccess() *string
	SetTraceId(v string) *ResumeTrafficResponseBody
	GetTraceId() *string
}

type ResumeTrafficResponseBody struct {
	// The API status code or POP error code. Valid values:
	//
	// - **2xx**: The request was successful.
	//
	// - **3xx**: Redirection.
	//
	// - **4xx**: A client-side error occurred.
	//
	// - **5xx**: A server-side error occurred.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned results.
	Data *ResumeTrafficResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// - This parameter is empty if the request is successful.
	//
	// - If the request fails, this parameter contains an error code. For more information, see the "**Error codes**" section of this topic.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The returned message.
	//
	// - If the request is successful, **success*	- is returned.
	//
	// - If the request fails, an error message is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 91F93257-7A4A-4BD3-9A7E-2F6EAE6D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the traffic was resumed. **True*	- indicates that the traffic was resumed, and **False*	- indicates that the traffic was not resumed.
	//
	// example:
	//
	// True
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID of the request. You can use this ID to troubleshoot the request.
	//
	// example:
	//
	// ac1a0b2215622920113732501e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s ResumeTrafficResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResumeTrafficResponseBody) GoString() string {
	return s.String()
}

func (s *ResumeTrafficResponseBody) GetCode() *string {
	return s.Code
}

func (s *ResumeTrafficResponseBody) GetData() *ResumeTrafficResponseBodyData {
	return s.Data
}

func (s *ResumeTrafficResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ResumeTrafficResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ResumeTrafficResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResumeTrafficResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *ResumeTrafficResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *ResumeTrafficResponseBody) SetCode(v string) *ResumeTrafficResponseBody {
	s.Code = &v
	return s
}

func (s *ResumeTrafficResponseBody) SetData(v *ResumeTrafficResponseBodyData) *ResumeTrafficResponseBody {
	s.Data = v
	return s
}

func (s *ResumeTrafficResponseBody) SetErrorCode(v string) *ResumeTrafficResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ResumeTrafficResponseBody) SetMessage(v string) *ResumeTrafficResponseBody {
	s.Message = &v
	return s
}

func (s *ResumeTrafficResponseBody) SetRequestId(v string) *ResumeTrafficResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResumeTrafficResponseBody) SetSuccess(v string) *ResumeTrafficResponseBody {
	s.Success = &v
	return s
}

func (s *ResumeTrafficResponseBody) SetTraceId(v string) *ResumeTrafficResponseBody {
	s.TraceId = &v
	return s
}

func (s *ResumeTrafficResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ResumeTrafficResponseBodyData struct {
	// Details about the operation\\"s result.
	//
	// example:
	//
	// succeed in handling request
	Msg *string `json:"msg,omitempty" xml:"msg,omitempty"`
	// Indicates whether the traffic was resumed.
	//
	// - **True**: The traffic was resumed.
	//
	// - **False**: The traffic was not resumed.
	//
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ResumeTrafficResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ResumeTrafficResponseBodyData) GoString() string {
	return s.String()
}

func (s *ResumeTrafficResponseBodyData) GetMsg() *string {
	return s.Msg
}

func (s *ResumeTrafficResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *ResumeTrafficResponseBodyData) SetMsg(v string) *ResumeTrafficResponseBodyData {
	s.Msg = &v
	return s
}

func (s *ResumeTrafficResponseBodyData) SetSuccess(v bool) *ResumeTrafficResponseBodyData {
	s.Success = &v
	return s
}

func (s *ResumeTrafficResponseBodyData) Validate() error {
	return dara.Validate(s)
}
