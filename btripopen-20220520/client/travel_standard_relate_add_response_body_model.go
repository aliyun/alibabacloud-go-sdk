// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTravelStandardRelateAddResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *TravelStandardRelateAddResponseBody
	GetMessage() *string
	SetRequestId(v string) *TravelStandardRelateAddResponseBody
	GetRequestId() *string
	SetResultCode(v int32) *TravelStandardRelateAddResponseBody
	GetResultCode() *int32
	SetSuccess(v bool) *TravelStandardRelateAddResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *TravelStandardRelateAddResponseBody
	GetTraceId() *string
}

type TravelStandardRelateAddResponseBody struct {
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 407543AF-2BD9-5890-BD92-9D1AB7218B27
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
	// 210bcc3a16583004579056128d33d7
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s TravelStandardRelateAddResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TravelStandardRelateAddResponseBody) GoString() string {
	return s.String()
}

func (s *TravelStandardRelateAddResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TravelStandardRelateAddResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TravelStandardRelateAddResponseBody) GetResultCode() *int32 {
	return s.ResultCode
}

func (s *TravelStandardRelateAddResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TravelStandardRelateAddResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *TravelStandardRelateAddResponseBody) SetMessage(v string) *TravelStandardRelateAddResponseBody {
	s.Message = &v
	return s
}

func (s *TravelStandardRelateAddResponseBody) SetRequestId(v string) *TravelStandardRelateAddResponseBody {
	s.RequestId = &v
	return s
}

func (s *TravelStandardRelateAddResponseBody) SetResultCode(v int32) *TravelStandardRelateAddResponseBody {
	s.ResultCode = &v
	return s
}

func (s *TravelStandardRelateAddResponseBody) SetSuccess(v bool) *TravelStandardRelateAddResponseBody {
	s.Success = &v
	return s
}

func (s *TravelStandardRelateAddResponseBody) SetTraceId(v string) *TravelStandardRelateAddResponseBody {
	s.TraceId = &v
	return s
}

func (s *TravelStandardRelateAddResponseBody) Validate() error {
	return dara.Validate(s)
}
