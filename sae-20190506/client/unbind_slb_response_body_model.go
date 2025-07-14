// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindSlbResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UnbindSlbResponseBody
	GetCode() *string
	SetData(v *UnbindSlbResponseBodyData) *UnbindSlbResponseBody
	GetData() *UnbindSlbResponseBodyData
	SetErrorCode(v string) *UnbindSlbResponseBody
	GetErrorCode() *string
	SetMessage(v string) *UnbindSlbResponseBody
	GetMessage() *string
	SetRequestId(v string) *UnbindSlbResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UnbindSlbResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *UnbindSlbResponseBody
	GetTraceId() *string
}

type UnbindSlbResponseBody struct {
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
	Data *UnbindSlbResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code. Valid values:
	//
	// 	- If the call is successful, the **ErrorCode*	- parameter is not returned.
	//
	// 	- If the call fails, the **ErrorCode*	- parameter is returned. For more information, see the **Error codes*	- section in this topic.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The returned message. Valid values:
	//
	// 	- success: If the call is successful, **success*	- is returned.
	//
	// 	- An error code: If the call fails, an error code is returned.
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
	// Indicates whether the internal-facing or Internet-facing SLB instance was disassociated. Valid values:
	//
	// 	- **true**: The SLB instance was disassociated.
	//
	// 	- **false**: The SLB instance failed to be disassociated.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID that is used to query the details of the request.
	//
	// example:
	//
	// 0a98a02315955564772843261e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s UnbindSlbResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnbindSlbResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindSlbResponseBody) GetCode() *string {
	return s.Code
}

func (s *UnbindSlbResponseBody) GetData() *UnbindSlbResponseBodyData {
	return s.Data
}

func (s *UnbindSlbResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UnbindSlbResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UnbindSlbResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnbindSlbResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UnbindSlbResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *UnbindSlbResponseBody) SetCode(v string) *UnbindSlbResponseBody {
	s.Code = &v
	return s
}

func (s *UnbindSlbResponseBody) SetData(v *UnbindSlbResponseBodyData) *UnbindSlbResponseBody {
	s.Data = v
	return s
}

func (s *UnbindSlbResponseBody) SetErrorCode(v string) *UnbindSlbResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UnbindSlbResponseBody) SetMessage(v string) *UnbindSlbResponseBody {
	s.Message = &v
	return s
}

func (s *UnbindSlbResponseBody) SetRequestId(v string) *UnbindSlbResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindSlbResponseBody) SetSuccess(v bool) *UnbindSlbResponseBody {
	s.Success = &v
	return s
}

func (s *UnbindSlbResponseBody) SetTraceId(v string) *UnbindSlbResponseBody {
	s.TraceId = &v
	return s
}

func (s *UnbindSlbResponseBody) Validate() error {
	return dara.Validate(s)
}

type UnbindSlbResponseBodyData struct {
	// The ID of the change order. The ID can be used to query the status of the change task.
	//
	// example:
	//
	// 4a815998-b468-4bea-b7d8-59f52a44****
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
}

func (s UnbindSlbResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UnbindSlbResponseBodyData) GoString() string {
	return s.String()
}

func (s *UnbindSlbResponseBodyData) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *UnbindSlbResponseBodyData) SetChangeOrderId(v string) *UnbindSlbResponseBodyData {
	s.ChangeOrderId = &v
	return s
}

func (s *UnbindSlbResponseBodyData) Validate() error {
	return dara.Validate(s)
}
