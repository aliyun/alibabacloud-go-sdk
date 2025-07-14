// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *StartApplicationResponseBody
	GetCode() *string
	SetData(v *StartApplicationResponseBodyData) *StartApplicationResponseBody
	GetData() *StartApplicationResponseBodyData
	SetErrorCode(v string) *StartApplicationResponseBody
	GetErrorCode() *string
	SetMessage(v string) *StartApplicationResponseBody
	GetMessage() *string
	SetRequestId(v string) *StartApplicationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StartApplicationResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *StartApplicationResponseBody
	GetTraceId() *string
}

type StartApplicationResponseBody struct {
	// The HTTP status code. Valid values:
	//
	// 	- **2xx**: The call was successful.
	//
	// 	- **3xx**: The call was redirected.
	//
	// 	- **4xx**: The call failed.
	//
	// 	- **5xx**: A server error occurred.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned result.
	Data *StartApplicationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned. Take note of the following rules:
	//
	// 	- If the call is successful, **ErrorCode*	- is not returned.
	//
	// 	- If the call fails, **ErrorCode*	- is returned. For more information, see the "**Error codes**" section in this topic.
	//
	// example:
	//
	// Null
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The returned message. Take note of the following rules:
	//
	// 	- If the call is successful, **success*	- is returned.
	//
	// 	- If the call fails, an error code is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7BD8F4C7-D84C-4D46-9885-8212997E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the application is started. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID that is used to query the details of the request.
	//
	// example:
	//
	// 0bc3b6e215637275918588187d****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s StartApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *StartApplicationResponseBody) GetCode() *string {
	return s.Code
}

func (s *StartApplicationResponseBody) GetData() *StartApplicationResponseBodyData {
	return s.Data
}

func (s *StartApplicationResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *StartApplicationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StartApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartApplicationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StartApplicationResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *StartApplicationResponseBody) SetCode(v string) *StartApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *StartApplicationResponseBody) SetData(v *StartApplicationResponseBodyData) *StartApplicationResponseBody {
	s.Data = v
	return s
}

func (s *StartApplicationResponseBody) SetErrorCode(v string) *StartApplicationResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *StartApplicationResponseBody) SetMessage(v string) *StartApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *StartApplicationResponseBody) SetRequestId(v string) *StartApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartApplicationResponseBody) SetSuccess(v bool) *StartApplicationResponseBody {
	s.Success = &v
	return s
}

func (s *StartApplicationResponseBody) SetTraceId(v string) *StartApplicationResponseBody {
	s.TraceId = &v
	return s
}

func (s *StartApplicationResponseBody) Validate() error {
	return dara.Validate(s)
}

type StartApplicationResponseBodyData struct {
	// The ID of the change order.
	//
	// example:
	//
	// 4a815998-b468-4bea-b7d8-59f52a44****
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
}

func (s StartApplicationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s StartApplicationResponseBodyData) GoString() string {
	return s.String()
}

func (s *StartApplicationResponseBodyData) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *StartApplicationResponseBodyData) SetChangeOrderId(v string) *StartApplicationResponseBodyData {
	s.ChangeOrderId = &v
	return s
}

func (s *StartApplicationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
