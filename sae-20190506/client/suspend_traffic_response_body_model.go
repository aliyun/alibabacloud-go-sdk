// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSuspendTrafficResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SuspendTrafficResponseBody
	GetCode() *string
	SetData(v *SuspendTrafficResponseBodyData) *SuspendTrafficResponseBody
	GetData() *SuspendTrafficResponseBodyData
	SetErrorCode(v string) *SuspendTrafficResponseBody
	GetErrorCode() *string
	SetMessage(v string) *SuspendTrafficResponseBody
	GetMessage() *string
	SetRequestId(v string) *SuspendTrafficResponseBody
	GetRequestId() *string
	SetSuccess(v string) *SuspendTrafficResponseBody
	GetSuccess() *string
	SetTraceId(v string) *SuspendTrafficResponseBody
	GetTraceId() *string
}

type SuspendTrafficResponseBody struct {
	// The HTTP status code. Valid values:
	//
	// 	- **2xx**: The request was successful.
	//
	// 	- **3xx**: The request was redirected.
	//
	// 	- **4xx**: The request failed.
	//
	// 	- **5xx**: A server error occurred.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *SuspendTrafficResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// - The **ErrorCode*	- parameter is not returned when the request succeeds.
	//
	// - The **ErrorCode*	- parameter is returned when the request fails. For more information, see **Error codes*	- in this topic.
	//
	// example:
	//
	// Null
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The message returned. Valid values:
	//
	// 	- If the request was successful, **success*	- is returned.
	//
	// 	- If the request failed, an error code is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B4D805CA-926D-41B1-8E63-7AD0C1ED****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the traffic was removed. Valid values:
	//
	// - **true**: The traffic was removed.
	//
	// - **false**: The traffic failed to be removed.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID.
	//
	// example:
	//
	// 0a98a02315955564772843261e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s SuspendTrafficResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SuspendTrafficResponseBody) GoString() string {
	return s.String()
}

func (s *SuspendTrafficResponseBody) GetCode() *string {
	return s.Code
}

func (s *SuspendTrafficResponseBody) GetData() *SuspendTrafficResponseBodyData {
	return s.Data
}

func (s *SuspendTrafficResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *SuspendTrafficResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SuspendTrafficResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SuspendTrafficResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *SuspendTrafficResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *SuspendTrafficResponseBody) SetCode(v string) *SuspendTrafficResponseBody {
	s.Code = &v
	return s
}

func (s *SuspendTrafficResponseBody) SetData(v *SuspendTrafficResponseBodyData) *SuspendTrafficResponseBody {
	s.Data = v
	return s
}

func (s *SuspendTrafficResponseBody) SetErrorCode(v string) *SuspendTrafficResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SuspendTrafficResponseBody) SetMessage(v string) *SuspendTrafficResponseBody {
	s.Message = &v
	return s
}

func (s *SuspendTrafficResponseBody) SetRequestId(v string) *SuspendTrafficResponseBody {
	s.RequestId = &v
	return s
}

func (s *SuspendTrafficResponseBody) SetSuccess(v string) *SuspendTrafficResponseBody {
	s.Success = &v
	return s
}

func (s *SuspendTrafficResponseBody) SetTraceId(v string) *SuspendTrafficResponseBody {
	s.TraceId = &v
	return s
}

func (s *SuspendTrafficResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SuspendTrafficResponseBodyData struct {
	// The description of the returned code.
	//
	// example:
	//
	// success
	Msg *string `json:"msg,omitempty" xml:"msg,omitempty"`
	// Indicates whether the traffic was removed. Valid values:
	//
	// 	- **true**: The traffic was removed.
	//
	// 	- **false**: The traffic failed to be removed.
	//
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SuspendTrafficResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SuspendTrafficResponseBodyData) GoString() string {
	return s.String()
}

func (s *SuspendTrafficResponseBodyData) GetMsg() *string {
	return s.Msg
}

func (s *SuspendTrafficResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *SuspendTrafficResponseBodyData) SetMsg(v string) *SuspendTrafficResponseBodyData {
	s.Msg = &v
	return s
}

func (s *SuspendTrafficResponseBodyData) SetSuccess(v bool) *SuspendTrafficResponseBodyData {
	s.Success = &v
	return s
}

func (s *SuspendTrafficResponseBodyData) Validate() error {
	return dara.Validate(s)
}
