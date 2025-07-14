// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindSlbResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BindSlbResponseBody
	GetCode() *string
	SetData(v *BindSlbResponseBodyData) *BindSlbResponseBody
	GetData() *BindSlbResponseBodyData
	SetErrorCode(v string) *BindSlbResponseBody
	GetErrorCode() *string
	SetMessage(v string) *BindSlbResponseBody
	GetMessage() *string
	SetRequestId(v string) *BindSlbResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *BindSlbResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *BindSlbResponseBody
	GetTraceId() *string
}

type BindSlbResponseBody struct {
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
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *BindSlbResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// 	- The **ErrorCode*	- parameter is not returned when the request succeeds.
	//
	// 	- The **ErrorCode*	- parameter is returned when the request fails. For more information, see **Error codes*	- in this topic.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The returned message.
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
	// Indicates whether the SLB instance was successfully associated with the application. Valid values:
	//
	// 	- **true**: The SLB instance was successfully associated with the application.
	//
	// 	- **false**: The SLB instance could not be associated with the application.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The ID of the trace. It can be used to query the details of a request.
	//
	// example:
	//
	// 0a98a02315955564772843261e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s BindSlbResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BindSlbResponseBody) GoString() string {
	return s.String()
}

func (s *BindSlbResponseBody) GetCode() *string {
	return s.Code
}

func (s *BindSlbResponseBody) GetData() *BindSlbResponseBodyData {
	return s.Data
}

func (s *BindSlbResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *BindSlbResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BindSlbResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BindSlbResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BindSlbResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *BindSlbResponseBody) SetCode(v string) *BindSlbResponseBody {
	s.Code = &v
	return s
}

func (s *BindSlbResponseBody) SetData(v *BindSlbResponseBodyData) *BindSlbResponseBody {
	s.Data = v
	return s
}

func (s *BindSlbResponseBody) SetErrorCode(v string) *BindSlbResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *BindSlbResponseBody) SetMessage(v string) *BindSlbResponseBody {
	s.Message = &v
	return s
}

func (s *BindSlbResponseBody) SetRequestId(v string) *BindSlbResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindSlbResponseBody) SetSuccess(v bool) *BindSlbResponseBody {
	s.Success = &v
	return s
}

func (s *BindSlbResponseBody) SetTraceId(v string) *BindSlbResponseBody {
	s.TraceId = &v
	return s
}

func (s *BindSlbResponseBody) Validate() error {
	return dara.Validate(s)
}

type BindSlbResponseBodyData struct {
	// The ID of the change order. It can be used to query the task status.
	//
	// example:
	//
	// 01db03d3-3ee9-48b3-b3d0-dfce2d88****
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
}

func (s BindSlbResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s BindSlbResponseBodyData) GoString() string {
	return s.String()
}

func (s *BindSlbResponseBodyData) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *BindSlbResponseBodyData) SetChangeOrderId(v string) *BindSlbResponseBodyData {
	s.ChangeOrderId = &v
	return s
}

func (s *BindSlbResponseBodyData) Validate() error {
	return dara.Validate(s)
}
