// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RestartInstancesResponseBody
	GetCode() *string
	SetData(v *RestartInstancesResponseBodyData) *RestartInstancesResponseBody
	GetData() *RestartInstancesResponseBodyData
	SetErrorCode(v string) *RestartInstancesResponseBody
	GetErrorCode() *string
	SetMessage(v string) *RestartInstancesResponseBody
	GetMessage() *string
	SetRequestId(v string) *RestartInstancesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RestartInstancesResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *RestartInstancesResponseBody
	GetTraceId() *string
}

type RestartInstancesResponseBody struct {
	// The HTTP status code. Take note of the following rules:
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
	// The details of the application.
	Data *RestartInstancesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned if the request failed. Take note of the following rules:
	//
	// 	- The **ErrorCode*	- parameter is not returned if the request succeeds.
	//
	// 	- If the call fails, the **ErrorCode*	- parameter is returned. For more information, see the "**Error codes**" section of this topic.
	//
	// example:
	//
	// Null
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The additional information that is returned. Take note of the following rules:
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
	// Specifies whether the instances are successfully restarted. Take note of the following rules:
	//
	// 	- **true**
	//
	// 	- **false**
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

func (s RestartInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RestartInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *RestartInstancesResponseBody) GetCode() *string {
	return s.Code
}

func (s *RestartInstancesResponseBody) GetData() *RestartInstancesResponseBodyData {
	return s.Data
}

func (s *RestartInstancesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RestartInstancesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RestartInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RestartInstancesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RestartInstancesResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *RestartInstancesResponseBody) SetCode(v string) *RestartInstancesResponseBody {
	s.Code = &v
	return s
}

func (s *RestartInstancesResponseBody) SetData(v *RestartInstancesResponseBodyData) *RestartInstancesResponseBody {
	s.Data = v
	return s
}

func (s *RestartInstancesResponseBody) SetErrorCode(v string) *RestartInstancesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RestartInstancesResponseBody) SetMessage(v string) *RestartInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *RestartInstancesResponseBody) SetRequestId(v string) *RestartInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestartInstancesResponseBody) SetSuccess(v bool) *RestartInstancesResponseBody {
	s.Success = &v
	return s
}

func (s *RestartInstancesResponseBody) SetTraceId(v string) *RestartInstancesResponseBody {
	s.TraceId = &v
	return s
}

func (s *RestartInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}

type RestartInstancesResponseBodyData struct {
	// The ID of the change order.
	//
	// example:
	//
	// 5afa5b98-0c64-4637-983f-15eaa888****
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
}

func (s RestartInstancesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RestartInstancesResponseBodyData) GoString() string {
	return s.String()
}

func (s *RestartInstancesResponseBodyData) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *RestartInstancesResponseBodyData) SetChangeOrderId(v string) *RestartInstancesResponseBodyData {
	s.ChangeOrderId = &v
	return s
}

func (s *RestartInstancesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
