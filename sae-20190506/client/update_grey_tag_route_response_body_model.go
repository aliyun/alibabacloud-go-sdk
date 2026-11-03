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
	// The API status code or POP error code. Valid values:
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
	// The canary release rule information.
	Data *UpdateGreyTagRouteResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// - If the request is successful, the **ErrorCode*	- field is not returned.
	//
	// - If the request fails, the **ErrorCode*	- field is returned. For more information, see the **Error codes*	- section in this topic.
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
	// The request ID.
	//
	// example:
	//
	// 9D29CBD0-45D3-410B-9826-52F86F90****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the change order information was queried successfully. Valid values:
	//
	// - **true**: The query was successful.
	//
	// - **false**: The query failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID, which is used to accurately query call information.
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
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateGreyTagRouteResponseBodyData struct {
	// The globally unique canary release rule ID.
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
