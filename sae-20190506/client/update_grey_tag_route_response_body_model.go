// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGreyTagRouteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateGreyTagRouteResponseBody
	GetCode() *string
	SetData(v *UpdateGreyTagRouteResponseBodyData) *UpdateGreyTagRouteResponseBody
	GetData() *UpdateGreyTagRouteResponseBodyData
	SetErrorCode(v string) *UpdateGreyTagRouteResponseBody
	GetErrorCode() *string
	SetMessage(v string) *UpdateGreyTagRouteResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateGreyTagRouteResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateGreyTagRouteResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *UpdateGreyTagRouteResponseBody
	GetTraceId() *string
}

type UpdateGreyTagRouteResponseBody struct {
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
	// The information about the canary release rule.
	Data *UpdateGreyTagRouteResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code. Valid values:
	//
	// 	- If the call is successful, the **ErrorCode*	- parameter is not returned.
	//
	// 	- If the call fails, the **ErrorCode*	- parameter is returned. For more information, see the **Error codes*	- section in this topic.
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
	// 9D29CBD0-45D3-410B-9826-52F86F90****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the information of the change order was queried. Valid values:
	//
	// 	- **true**: The information was queried.
	//
	// 	- **false**: The information failed to be queried.
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

func (s UpdateGreyTagRouteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateGreyTagRouteResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGreyTagRouteResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateGreyTagRouteResponseBody) GetData() *UpdateGreyTagRouteResponseBodyData {
	return s.Data
}

func (s *UpdateGreyTagRouteResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateGreyTagRouteResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateGreyTagRouteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateGreyTagRouteResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateGreyTagRouteResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *UpdateGreyTagRouteResponseBody) SetCode(v string) *UpdateGreyTagRouteResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateGreyTagRouteResponseBody) SetData(v *UpdateGreyTagRouteResponseBodyData) *UpdateGreyTagRouteResponseBody {
	s.Data = v
	return s
}

func (s *UpdateGreyTagRouteResponseBody) SetErrorCode(v string) *UpdateGreyTagRouteResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateGreyTagRouteResponseBody) SetMessage(v string) *UpdateGreyTagRouteResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateGreyTagRouteResponseBody) SetRequestId(v string) *UpdateGreyTagRouteResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGreyTagRouteResponseBody) SetSuccess(v bool) *UpdateGreyTagRouteResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateGreyTagRouteResponseBody) SetTraceId(v string) *UpdateGreyTagRouteResponseBody {
	s.TraceId = &v
	return s
}

func (s *UpdateGreyTagRouteResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdateGreyTagRouteResponseBodyData struct {
	// The ID of the canary release rule. The ID is globally unique.
	//
	// example:
	//
	// 1
	GreyTagRouteId *int64 `json:"GreyTagRouteId,omitempty" xml:"GreyTagRouteId,omitempty"`
}

func (s UpdateGreyTagRouteResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateGreyTagRouteResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateGreyTagRouteResponseBodyData) GetGreyTagRouteId() *int64 {
	return s.GreyTagRouteId
}

func (s *UpdateGreyTagRouteResponseBodyData) SetGreyTagRouteId(v int64) *UpdateGreyTagRouteResponseBodyData {
	s.GreyTagRouteId = &v
	return s
}

func (s *UpdateGreyTagRouteResponseBodyData) Validate() error {
	return dara.Validate(s)
}
