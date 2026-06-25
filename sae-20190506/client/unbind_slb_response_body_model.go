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
	// - **2xx**: success
	//
	// - **3xx**: redirection
	//
	// - **4xx**: request error
	//
	// - **5xx**: server error
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	Data *UnbindSlbResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned error code.
	//
	// - This parameter is left empty if the request is successful.
	//
	// - If the request fails, this parameter contains an error code. For more information, see the **Error codes*	- section in this topic.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The returned message. Valid values:
	//
	// - If the request succeeds, **success*	- is returned.
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
	// Indicates whether the operation was successful.
	//
	// - **true**: The operation was successful.
	//
	// - **false**: The operation failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID of the request. You can use this ID to query call details.
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
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UnbindSlbResponseBodyData struct {
	// The change order ID. You can use this ID to query the status of the task.
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
