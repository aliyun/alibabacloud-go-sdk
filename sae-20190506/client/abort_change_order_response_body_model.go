// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAbortChangeOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AbortChangeOrderResponseBody
	GetCode() *string
	SetData(v *AbortChangeOrderResponseBodyData) *AbortChangeOrderResponseBody
	GetData() *AbortChangeOrderResponseBodyData
	SetErrorCode(v string) *AbortChangeOrderResponseBody
	GetErrorCode() *string
	SetMessage(v string) *AbortChangeOrderResponseBody
	GetMessage() *string
	SetRequestId(v string) *AbortChangeOrderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AbortChangeOrderResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *AbortChangeOrderResponseBody
	GetTraceId() *string
}

type AbortChangeOrderResponseBody struct {
	// The API status code or POP error code. Valid values:
	//
	// - **2xx**: Success.
	//
	// - **3xx**: Redirection.
	//
	// - **4xx**: Request error.
	//
	// - **5xx**: Server error.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	Data *AbortChangeOrderResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code. Valid values:
	//
	// - If the request is successful, the **ErrorCode*	- parameter is not returned.
	//
	// - If the request fails, the **ErrorCode*	- parameter is returned. For more information, see the **Error codes*	- section in this topic.
	//
	// example:
	//
	// InvalidParameter.NotEmpty
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The additional information about the call result.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 91F93257-7A4A-4BD3-9A7E-2F6EAE6D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the change order was aborted successfully. Valid values:
	//
	// - **true**: The change order was aborted successfully.
	//
	// - **false**: The change order failed to be aborted.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID.
	//
	// example:
	//
	// 0a98a02315955564772843261e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s AbortChangeOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AbortChangeOrderResponseBody) GoString() string {
	return s.String()
}

func (s *AbortChangeOrderResponseBody) GetCode() *string {
	return s.Code
}

func (s *AbortChangeOrderResponseBody) GetData() *AbortChangeOrderResponseBodyData {
	return s.Data
}

func (s *AbortChangeOrderResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *AbortChangeOrderResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AbortChangeOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AbortChangeOrderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AbortChangeOrderResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *AbortChangeOrderResponseBody) SetCode(v string) *AbortChangeOrderResponseBody {
	s.Code = &v
	return s
}

func (s *AbortChangeOrderResponseBody) SetData(v *AbortChangeOrderResponseBodyData) *AbortChangeOrderResponseBody {
	s.Data = v
	return s
}

func (s *AbortChangeOrderResponseBody) SetErrorCode(v string) *AbortChangeOrderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *AbortChangeOrderResponseBody) SetMessage(v string) *AbortChangeOrderResponseBody {
	s.Message = &v
	return s
}

func (s *AbortChangeOrderResponseBody) SetRequestId(v string) *AbortChangeOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *AbortChangeOrderResponseBody) SetSuccess(v bool) *AbortChangeOrderResponseBody {
	s.Success = &v
	return s
}

func (s *AbortChangeOrderResponseBody) SetTraceId(v string) *AbortChangeOrderResponseBody {
	s.TraceId = &v
	return s
}

func (s *AbortChangeOrderResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AbortChangeOrderResponseBodyData struct {
	// The ID of the change order.
	//
	// example:
	//
	// be2e1c76-682b-4897-98d3-1d8d6478****
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
}

func (s AbortChangeOrderResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AbortChangeOrderResponseBodyData) GoString() string {
	return s.String()
}

func (s *AbortChangeOrderResponseBodyData) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *AbortChangeOrderResponseBodyData) SetChangeOrderId(v string) *AbortChangeOrderResponseBodyData {
	s.ChangeOrderId = &v
	return s
}

func (s *AbortChangeOrderResponseBodyData) Validate() error {
	return dara.Validate(s)
}
