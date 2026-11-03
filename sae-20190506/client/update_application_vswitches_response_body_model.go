// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationVswitchesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateApplicationVswitchesResponseBody
	GetCode() *string
	SetErrorCode(v string) *UpdateApplicationVswitchesResponseBody
	GetErrorCode() *string
	SetMessage(v string) *UpdateApplicationVswitchesResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateApplicationVswitchesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateApplicationVswitchesResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *UpdateApplicationVswitchesResponseBody
	GetTraceId() *string
}

type UpdateApplicationVswitchesResponseBody struct {
	// The API status code or POP error code. Valid values:
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
	// The error code.
	//
	// - If the request is successful, the **ErrorCode*	- parameter is not returned.
	//
	// - If the request fails, the **ErrorCode*	- parameter is returned. For more information, see the **Error codes*	- section in this topic.
	//
	// example:
	//
	// InvalidParameter.NotEmpty
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The additional information.
	//
	// - If the request is successful, **success*	- is returned.
	//
	// - If an exception occurs, a specific error code is returned.
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
	// Indicates whether the application list is retrieved. Valid values:
	//
	// - **true**: The application list is retrieved.
	//
	// - **false**: The application list fails to be retrieved.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID used to precisely query call information.
	//
	// example:
	//
	// 0a98a02315955564772843261e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s UpdateApplicationVswitchesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationVswitchesResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateApplicationVswitchesResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateApplicationVswitchesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateApplicationVswitchesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateApplicationVswitchesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateApplicationVswitchesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateApplicationVswitchesResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *UpdateApplicationVswitchesResponseBody) SetCode(v string) *UpdateApplicationVswitchesResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateApplicationVswitchesResponseBody) SetErrorCode(v string) *UpdateApplicationVswitchesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateApplicationVswitchesResponseBody) SetMessage(v string) *UpdateApplicationVswitchesResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateApplicationVswitchesResponseBody) SetRequestId(v string) *UpdateApplicationVswitchesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateApplicationVswitchesResponseBody) SetSuccess(v bool) *UpdateApplicationVswitchesResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateApplicationVswitchesResponseBody) SetTraceId(v string) *UpdateApplicationVswitchesResponseBody {
	s.TraceId = &v
	return s
}

func (s *UpdateApplicationVswitchesResponseBody) Validate() error {
	return dara.Validate(s)
}
