// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindNlbResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UnbindNlbResponseBody
	GetCode() *string
	SetData(v *UnbindNlbResponseBodyData) *UnbindNlbResponseBody
	GetData() *UnbindNlbResponseBodyData
	SetErrorCode(v string) *UnbindNlbResponseBody
	GetErrorCode() *string
	SetMessage(v string) *UnbindNlbResponseBody
	GetMessage() *string
	SetRequestId(v string) *UnbindNlbResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UnbindNlbResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *UnbindNlbResponseBody
	GetTraceId() *string
}

type UnbindNlbResponseBody struct {
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
	Data *UnbindNlbResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The status code. Valid values:
	//
	// 	- If the request was successful, **ErrorCode*	- is not returned.
	//
	// 	- If the request failed, **ErrorCode*	- is returned. For more information, see **Error codes*	- section of this topic.
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
	// Id of the request
	//
	// example:
	//
	// 91F93257-7A4A-4BD3-9A7E-2F6EAE6D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the internal-facing or Internet-facing NLB instance was disassociated. Valid values:
	//
	// 	- **true**: The NLB instance was disassociated.
	//
	// 	- **false**: The NLB instance failed to be disassociated.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The ID of the trace. The ID is used to query the details of a request.
	//
	// example:
	//
	// 0a981dd515966966104121683d****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s UnbindNlbResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnbindNlbResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindNlbResponseBody) GetCode() *string {
	return s.Code
}

func (s *UnbindNlbResponseBody) GetData() *UnbindNlbResponseBodyData {
	return s.Data
}

func (s *UnbindNlbResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UnbindNlbResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UnbindNlbResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnbindNlbResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UnbindNlbResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *UnbindNlbResponseBody) SetCode(v string) *UnbindNlbResponseBody {
	s.Code = &v
	return s
}

func (s *UnbindNlbResponseBody) SetData(v *UnbindNlbResponseBodyData) *UnbindNlbResponseBody {
	s.Data = v
	return s
}

func (s *UnbindNlbResponseBody) SetErrorCode(v string) *UnbindNlbResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UnbindNlbResponseBody) SetMessage(v string) *UnbindNlbResponseBody {
	s.Message = &v
	return s
}

func (s *UnbindNlbResponseBody) SetRequestId(v string) *UnbindNlbResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindNlbResponseBody) SetSuccess(v bool) *UnbindNlbResponseBody {
	s.Success = &v
	return s
}

func (s *UnbindNlbResponseBody) SetTraceId(v string) *UnbindNlbResponseBody {
	s.TraceId = &v
	return s
}

func (s *UnbindNlbResponseBody) Validate() error {
	return dara.Validate(s)
}

type UnbindNlbResponseBodyData struct {
	// The ID of the change order. The ID can be used to query the status of the change task.
	//
	// example:
	//
	// ba386059-69b1-4e65-b1e5-0682d9fa****
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
}

func (s UnbindNlbResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UnbindNlbResponseBodyData) GoString() string {
	return s.String()
}

func (s *UnbindNlbResponseBodyData) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *UnbindNlbResponseBodyData) SetChangeOrderId(v string) *UnbindNlbResponseBodyData {
	s.ChangeOrderId = &v
	return s
}

func (s *UnbindNlbResponseBodyData) Validate() error {
	return dara.Validate(s)
}
