// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *StopApplicationResponseBody
	GetCode() *string
	SetData(v *StopApplicationResponseBodyData) *StopApplicationResponseBody
	GetData() *StopApplicationResponseBodyData
	SetErrorCode(v string) *StopApplicationResponseBody
	GetErrorCode() *string
	SetMessage(v string) *StopApplicationResponseBody
	GetMessage() *string
	SetRequestId(v string) *StopApplicationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StopApplicationResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *StopApplicationResponseBody
	GetTraceId() *string
}

type StopApplicationResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error code.
	//
	// 	- If the request is successful, this parameter is not returned.****
	//
	// 	- This parameter is returned only if the request failed.***	- For more information, see **Error codes*	- in this topic.
	Data *StopApplicationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Indicates whether the specified application is stopped. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The returned data.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the trace. It can be used to query the details of a request.
	//
	// example:
	//
	// 91F93257-7A4A-4BD3-9A7E-2F6EAE6D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The ID of the change order.
	//
	// example:
	//
	// 0bc3b6e215637275918588187d****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s StopApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *StopApplicationResponseBody) GetCode() *string {
	return s.Code
}

func (s *StopApplicationResponseBody) GetData() *StopApplicationResponseBodyData {
	return s.Data
}

func (s *StopApplicationResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *StopApplicationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StopApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopApplicationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StopApplicationResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *StopApplicationResponseBody) SetCode(v string) *StopApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *StopApplicationResponseBody) SetData(v *StopApplicationResponseBodyData) *StopApplicationResponseBody {
	s.Data = v
	return s
}

func (s *StopApplicationResponseBody) SetErrorCode(v string) *StopApplicationResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *StopApplicationResponseBody) SetMessage(v string) *StopApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *StopApplicationResponseBody) SetRequestId(v string) *StopApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopApplicationResponseBody) SetSuccess(v bool) *StopApplicationResponseBody {
	s.Success = &v
	return s
}

func (s *StopApplicationResponseBody) SetTraceId(v string) *StopApplicationResponseBody {
	s.TraceId = &v
	return s
}

func (s *StopApplicationResponseBody) Validate() error {
	return dara.Validate(s)
}

type StopApplicationResponseBodyData struct {
	// The HTTP status code. Valid values:
	//
	// 	- **2xx**: indicates that the request was successful.
	//
	// 	- **3xx**: indicates that the request was redirected.
	//
	// 	- **4xx**: indicates that the request was invalid.
	//
	// 	- **5xx**: indicates that a server error occurred.
	//
	// example:
	//
	// 4a815998-b468-4bea-b7d8-59f52a44****
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
}

func (s StopApplicationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s StopApplicationResponseBodyData) GoString() string {
	return s.String()
}

func (s *StopApplicationResponseBodyData) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *StopApplicationResponseBodyData) SetChangeOrderId(v string) *StopApplicationResponseBodyData {
	s.ChangeOrderId = &v
	return s
}

func (s *StopApplicationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
