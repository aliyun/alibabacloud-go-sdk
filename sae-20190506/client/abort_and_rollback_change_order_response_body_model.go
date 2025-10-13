// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAbortAndRollbackChangeOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AbortAndRollbackChangeOrderResponseBody
	GetCode() *string
	SetData(v *AbortAndRollbackChangeOrderResponseBodyData) *AbortAndRollbackChangeOrderResponseBody
	GetData() *AbortAndRollbackChangeOrderResponseBodyData
	SetErrorCode(v string) *AbortAndRollbackChangeOrderResponseBody
	GetErrorCode() *string
	SetMessage(v string) *AbortAndRollbackChangeOrderResponseBody
	GetMessage() *string
	SetRequestId(v string) *AbortAndRollbackChangeOrderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AbortAndRollbackChangeOrderResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *AbortAndRollbackChangeOrderResponseBody
	GetTraceId() *string
}

type AbortAndRollbackChangeOrderResponseBody struct {
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
	// The details of the change order.
	Data *AbortAndRollbackChangeOrderResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code. Valid values:
	//
	// 	- If the call is successful, the **ErrorCode*	- parameter is not returned.
	//
	// 	- If the call fails, the **ErrorCode*	- parameter is returned. For more information, see the **Error codes*	- section in this topic.
	//
	// example:
	//
	// success
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The returned message.
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
	// Indicates whether the change order was terminated. Valid values:
	//
	// 	- **true**: The change order was terminated.
	//
	// 	- **false**: The change order failed to be terminated.
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

func (s AbortAndRollbackChangeOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AbortAndRollbackChangeOrderResponseBody) GoString() string {
	return s.String()
}

func (s *AbortAndRollbackChangeOrderResponseBody) GetCode() *string {
	return s.Code
}

func (s *AbortAndRollbackChangeOrderResponseBody) GetData() *AbortAndRollbackChangeOrderResponseBodyData {
	return s.Data
}

func (s *AbortAndRollbackChangeOrderResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *AbortAndRollbackChangeOrderResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AbortAndRollbackChangeOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AbortAndRollbackChangeOrderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AbortAndRollbackChangeOrderResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *AbortAndRollbackChangeOrderResponseBody) SetCode(v string) *AbortAndRollbackChangeOrderResponseBody {
	s.Code = &v
	return s
}

func (s *AbortAndRollbackChangeOrderResponseBody) SetData(v *AbortAndRollbackChangeOrderResponseBodyData) *AbortAndRollbackChangeOrderResponseBody {
	s.Data = v
	return s
}

func (s *AbortAndRollbackChangeOrderResponseBody) SetErrorCode(v string) *AbortAndRollbackChangeOrderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *AbortAndRollbackChangeOrderResponseBody) SetMessage(v string) *AbortAndRollbackChangeOrderResponseBody {
	s.Message = &v
	return s
}

func (s *AbortAndRollbackChangeOrderResponseBody) SetRequestId(v string) *AbortAndRollbackChangeOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *AbortAndRollbackChangeOrderResponseBody) SetSuccess(v bool) *AbortAndRollbackChangeOrderResponseBody {
	s.Success = &v
	return s
}

func (s *AbortAndRollbackChangeOrderResponseBody) SetTraceId(v string) *AbortAndRollbackChangeOrderResponseBody {
	s.TraceId = &v
	return s
}

func (s *AbortAndRollbackChangeOrderResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AbortAndRollbackChangeOrderResponseBodyData struct {
	// The ID of the change order.
	//
	// example:
	//
	// ba386059-69b1-4e65-b1e5-0682d9fa****
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
}

func (s AbortAndRollbackChangeOrderResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AbortAndRollbackChangeOrderResponseBodyData) GoString() string {
	return s.String()
}

func (s *AbortAndRollbackChangeOrderResponseBodyData) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *AbortAndRollbackChangeOrderResponseBodyData) SetChangeOrderId(v string) *AbortAndRollbackChangeOrderResponseBodyData {
	s.ChangeOrderId = &v
	return s
}

func (s *AbortAndRollbackChangeOrderResponseBodyData) Validate() error {
	return dara.Validate(s)
}
