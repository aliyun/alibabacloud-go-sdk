// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableArmsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DisableArmsResponseBody
	GetCode() *string
	SetData(v *DisableArmsResponseBodyData) *DisableArmsResponseBody
	GetData() *DisableArmsResponseBodyData
	SetErrorCode(v string) *DisableArmsResponseBody
	GetErrorCode() *string
	SetMessage(v string) *DisableArmsResponseBody
	GetMessage() *string
	SetRequestId(v string) *DisableArmsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DisableArmsResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *DisableArmsResponseBody
	GetTraceId() *string
}

type DisableArmsResponseBody struct {
	// The HTTP status code. Valid values:
	//
	// - **2xx**: The request was successful.
	//
	// - **3xx**: The request was redirected.
	//
	// - **4xx**: A client error occurred.
	//
	// - **5xx**: A server error occurred.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *DisableArmsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// - If the request is successful, the **ErrorCode*	- parameter is not returned.
	//
	// - If the request fails, the **ErrorCode*	- parameter is returned. For more information, see the **Error codes*	- section in this topic.
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
	// 6DA47906-6070-5A16-896D-67DCF38A6B7B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID that is used to query the details of a request.
	//
	// example:
	//
	// 0bc3b4ad17412276398692303e4cb1
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DisableArmsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableArmsResponseBody) GoString() string {
	return s.String()
}

func (s *DisableArmsResponseBody) GetCode() *string {
	return s.Code
}

func (s *DisableArmsResponseBody) GetData() *DisableArmsResponseBodyData {
	return s.Data
}

func (s *DisableArmsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DisableArmsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DisableArmsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableArmsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DisableArmsResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *DisableArmsResponseBody) SetCode(v string) *DisableArmsResponseBody {
	s.Code = &v
	return s
}

func (s *DisableArmsResponseBody) SetData(v *DisableArmsResponseBodyData) *DisableArmsResponseBody {
	s.Data = v
	return s
}

func (s *DisableArmsResponseBody) SetErrorCode(v string) *DisableArmsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DisableArmsResponseBody) SetMessage(v string) *DisableArmsResponseBody {
	s.Message = &v
	return s
}

func (s *DisableArmsResponseBody) SetRequestId(v string) *DisableArmsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableArmsResponseBody) SetSuccess(v bool) *DisableArmsResponseBody {
	s.Success = &v
	return s
}

func (s *DisableArmsResponseBody) SetTraceId(v string) *DisableArmsResponseBody {
	s.TraceId = &v
	return s
}

func (s *DisableArmsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DisableArmsResponseBodyData struct {
	// Indicates whether ARMS monitoring is enabled. Valid values:
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
}

func (s DisableArmsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DisableArmsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DisableArmsResponseBodyData) GetEnable() *bool {
	return s.Enable
}

func (s *DisableArmsResponseBodyData) SetEnable(v bool) *DisableArmsResponseBodyData {
	s.Enable = &v
	return s
}

func (s *DisableArmsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
