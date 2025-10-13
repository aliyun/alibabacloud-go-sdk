// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchStartApplicationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BatchStartApplicationsResponseBody
	GetCode() *string
	SetData(v *BatchStartApplicationsResponseBodyData) *BatchStartApplicationsResponseBody
	GetData() *BatchStartApplicationsResponseBodyData
	SetErrorCode(v string) *BatchStartApplicationsResponseBody
	GetErrorCode() *string
	SetMessage(v string) *BatchStartApplicationsResponseBody
	GetMessage() *string
	SetRequestId(v string) *BatchStartApplicationsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *BatchStartApplicationsResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *BatchStartApplicationsResponseBody
	GetTraceId() *string
}

type BatchStartApplicationsResponseBody struct {
	// The HTTP status code. Take note of the following rules:
	//
	// - **2xx**: The call was successful.
	//
	// - **3xx**: The call was redirected.
	//
	// - **4xx**: The call failed.
	//
	// - **5xx**: A server error occurred.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error code.
	//
	// 	- If the request is successful, this parameter is not returned.****
	//
	// 	- This parameter is returned only if the request failed.***	- For more information, see **Error codes*	- in this topic.
	Data *BatchStartApplicationsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned if the request failed. Take note of the following rules:
	//
	// - The ErrorCode parameter is not returned if the request succeeds.
	//
	// - If the call fails, the ErrorCode parameter is returned. For more information, see the "Error codes" section of this topic.
	//
	// example:
	//
	// NULL
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The returned data.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the trace. It is used to query the details of a request.
	//
	// example:
	//
	// 91F93257-7A4A-4BD3-9A7E-2F6EAE6D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the application deployment is successful. Take note of the following rules:
	//
	// - **true**
	//
	// - **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The ID of the change order.
	//
	// example:
	//
	// 0a98a02315955564772843261e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s BatchStartApplicationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchStartApplicationsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchStartApplicationsResponseBody) GetCode() *string {
	return s.Code
}

func (s *BatchStartApplicationsResponseBody) GetData() *BatchStartApplicationsResponseBodyData {
	return s.Data
}

func (s *BatchStartApplicationsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *BatchStartApplicationsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BatchStartApplicationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchStartApplicationsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BatchStartApplicationsResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *BatchStartApplicationsResponseBody) SetCode(v string) *BatchStartApplicationsResponseBody {
	s.Code = &v
	return s
}

func (s *BatchStartApplicationsResponseBody) SetData(v *BatchStartApplicationsResponseBodyData) *BatchStartApplicationsResponseBody {
	s.Data = v
	return s
}

func (s *BatchStartApplicationsResponseBody) SetErrorCode(v string) *BatchStartApplicationsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *BatchStartApplicationsResponseBody) SetMessage(v string) *BatchStartApplicationsResponseBody {
	s.Message = &v
	return s
}

func (s *BatchStartApplicationsResponseBody) SetRequestId(v string) *BatchStartApplicationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchStartApplicationsResponseBody) SetSuccess(v bool) *BatchStartApplicationsResponseBody {
	s.Success = &v
	return s
}

func (s *BatchStartApplicationsResponseBody) SetTraceId(v string) *BatchStartApplicationsResponseBody {
	s.TraceId = &v
	return s
}

func (s *BatchStartApplicationsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BatchStartApplicationsResponseBodyData struct {
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
	// 01db03d3-3ee9-48b3-b3d0-dfce2d88****
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
}

func (s BatchStartApplicationsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s BatchStartApplicationsResponseBodyData) GoString() string {
	return s.String()
}

func (s *BatchStartApplicationsResponseBodyData) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *BatchStartApplicationsResponseBodyData) SetChangeOrderId(v string) *BatchStartApplicationsResponseBodyData {
	s.ChangeOrderId = &v
	return s
}

func (s *BatchStartApplicationsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
