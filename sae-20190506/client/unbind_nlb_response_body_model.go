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
	// The HTTP status code.
	//
	// - **2xx**: Success.
	//
	// - **3xx**: Redirection.
	//
	// - **4xx**: Client error.
	//
	// - **5xx**: Server error.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *UnbindNlbResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// - The **ErrorCode*	- parameter is not returned if the request is successful.
	//
	// - The **ErrorCode*	- parameter is returned if the request fails. For more information, see the **Error codes*	- section of this topic.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The returned message.
	//
	// - If the request is successful, **success*	- is returned.
	//
	// - If the request fails, an error code is returned.
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
	// Indicates whether the request was successful.
	//
	// - **true**: The request was successful.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID for the request. Use this ID to query request details.
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
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UnbindNlbResponseBodyData struct {
	// The change order ID. Use this ID to check the task\\"s execution status.
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
