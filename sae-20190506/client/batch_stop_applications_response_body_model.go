// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchStopApplicationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BatchStopApplicationsResponseBody
	GetCode() *string
	SetData(v *BatchStopApplicationsResponseBodyData) *BatchStopApplicationsResponseBody
	GetData() *BatchStopApplicationsResponseBodyData
	SetErrorCode(v string) *BatchStopApplicationsResponseBody
	GetErrorCode() *string
	SetMessage(v string) *BatchStopApplicationsResponseBody
	GetMessage() *string
	SetRequestId(v string) *BatchStopApplicationsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *BatchStopApplicationsResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *BatchStopApplicationsResponseBody
	GetTraceId() *string
}

type BatchStopApplicationsResponseBody struct {
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
	// The ID of the change order.
	Data *BatchStopApplicationsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned if the request failed. Take note of the following rules:
	//
	// - The ErrorCode parameter is not returned if the request succeeds.
	//
	// - If the call fails, the ErrorCode parameter is returned. For more information, see the "Error codes" section of this topic.
	//
	// example:
	//
	// Null
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The ID of the trace. It can be used to query the details of a request.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The returned message.
	//
	// 	- **success*	- is returned when the request succeeds.
	//
	// 	- An error code is returned when the request fails.
	//
	// example:
	//
	// 7BD8F4C7-D84C-4D46-9885-8212997E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the application is created. Valid values
	//
	// - **true**
	//
	// - **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The returned data.
	//
	// example:
	//
	// 0bc3b6e215637275918588187d****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s BatchStopApplicationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchStopApplicationsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchStopApplicationsResponseBody) GetCode() *string {
	return s.Code
}

func (s *BatchStopApplicationsResponseBody) GetData() *BatchStopApplicationsResponseBodyData {
	return s.Data
}

func (s *BatchStopApplicationsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *BatchStopApplicationsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BatchStopApplicationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchStopApplicationsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BatchStopApplicationsResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *BatchStopApplicationsResponseBody) SetCode(v string) *BatchStopApplicationsResponseBody {
	s.Code = &v
	return s
}

func (s *BatchStopApplicationsResponseBody) SetData(v *BatchStopApplicationsResponseBodyData) *BatchStopApplicationsResponseBody {
	s.Data = v
	return s
}

func (s *BatchStopApplicationsResponseBody) SetErrorCode(v string) *BatchStopApplicationsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *BatchStopApplicationsResponseBody) SetMessage(v string) *BatchStopApplicationsResponseBody {
	s.Message = &v
	return s
}

func (s *BatchStopApplicationsResponseBody) SetRequestId(v string) *BatchStopApplicationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchStopApplicationsResponseBody) SetSuccess(v bool) *BatchStopApplicationsResponseBody {
	s.Success = &v
	return s
}

func (s *BatchStopApplicationsResponseBody) SetTraceId(v string) *BatchStopApplicationsResponseBody {
	s.TraceId = &v
	return s
}

func (s *BatchStopApplicationsResponseBody) Validate() error {
	return dara.Validate(s)
}

type BatchStopApplicationsResponseBodyData struct {
	// The error code.
	//
	// 	- If the request is successful, this parameter is not returned.****
	//
	// 	- This parameter is returned only if the request failed.***	- For more information, see the "**Error codes**" section in this topic.
	//
	// example:
	//
	// 4a815998-b468-4bea-b7d8-59f52a44****
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
}

func (s BatchStopApplicationsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s BatchStopApplicationsResponseBodyData) GoString() string {
	return s.String()
}

func (s *BatchStopApplicationsResponseBodyData) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *BatchStopApplicationsResponseBodyData) SetChangeOrderId(v string) *BatchStopApplicationsResponseBodyData {
	s.ChangeOrderId = &v
	return s
}

func (s *BatchStopApplicationsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
