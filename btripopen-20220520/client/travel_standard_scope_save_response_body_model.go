// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTravelStandardScopeSaveResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *TravelStandardScopeSaveResponseBody
	GetMessage() *string
	SetRequestId(v string) *TravelStandardScopeSaveResponseBody
	GetRequestId() *string
	SetResultCode(v int32) *TravelStandardScopeSaveResponseBody
	GetResultCode() *int32
	SetSuccess(v bool) *TravelStandardScopeSaveResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *TravelStandardScopeSaveResponseBody
	GetTraceId() *string
}

type TravelStandardScopeSaveResponseBody struct {
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 05F9C201-1B53-5905-94AB-0D7444D8466A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 0
	ResultCode *int32 `json:"resultCode,omitempty" xml:"resultCode,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// traceId
	//
	// example:
	//
	// 21041aa317070996148671005d0a0b
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s TravelStandardScopeSaveResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TravelStandardScopeSaveResponseBody) GoString() string {
	return s.String()
}

func (s *TravelStandardScopeSaveResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TravelStandardScopeSaveResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TravelStandardScopeSaveResponseBody) GetResultCode() *int32 {
	return s.ResultCode
}

func (s *TravelStandardScopeSaveResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TravelStandardScopeSaveResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *TravelStandardScopeSaveResponseBody) SetMessage(v string) *TravelStandardScopeSaveResponseBody {
	s.Message = &v
	return s
}

func (s *TravelStandardScopeSaveResponseBody) SetRequestId(v string) *TravelStandardScopeSaveResponseBody {
	s.RequestId = &v
	return s
}

func (s *TravelStandardScopeSaveResponseBody) SetResultCode(v int32) *TravelStandardScopeSaveResponseBody {
	s.ResultCode = &v
	return s
}

func (s *TravelStandardScopeSaveResponseBody) SetSuccess(v bool) *TravelStandardScopeSaveResponseBody {
	s.Success = &v
	return s
}

func (s *TravelStandardScopeSaveResponseBody) SetTraceId(v string) *TravelStandardScopeSaveResponseBody {
	s.TraceId = &v
	return s
}

func (s *TravelStandardScopeSaveResponseBody) Validate() error {
	return dara.Validate(s)
}
