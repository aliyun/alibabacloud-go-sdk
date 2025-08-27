// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTravelStandardRelateDeleteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *TravelStandardRelateDeleteResponseBody
	GetMessage() *string
	SetRequestId(v string) *TravelStandardRelateDeleteResponseBody
	GetRequestId() *string
	SetResultCode(v int32) *TravelStandardRelateDeleteResponseBody
	GetResultCode() *int32
	SetSuccess(v bool) *TravelStandardRelateDeleteResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *TravelStandardRelateDeleteResponseBody
	GetTraceId() *string
}

type TravelStandardRelateDeleteResponseBody struct {
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

func (s TravelStandardRelateDeleteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TravelStandardRelateDeleteResponseBody) GoString() string {
	return s.String()
}

func (s *TravelStandardRelateDeleteResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TravelStandardRelateDeleteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TravelStandardRelateDeleteResponseBody) GetResultCode() *int32 {
	return s.ResultCode
}

func (s *TravelStandardRelateDeleteResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TravelStandardRelateDeleteResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *TravelStandardRelateDeleteResponseBody) SetMessage(v string) *TravelStandardRelateDeleteResponseBody {
	s.Message = &v
	return s
}

func (s *TravelStandardRelateDeleteResponseBody) SetRequestId(v string) *TravelStandardRelateDeleteResponseBody {
	s.RequestId = &v
	return s
}

func (s *TravelStandardRelateDeleteResponseBody) SetResultCode(v int32) *TravelStandardRelateDeleteResponseBody {
	s.ResultCode = &v
	return s
}

func (s *TravelStandardRelateDeleteResponseBody) SetSuccess(v bool) *TravelStandardRelateDeleteResponseBody {
	s.Success = &v
	return s
}

func (s *TravelStandardRelateDeleteResponseBody) SetTraceId(v string) *TravelStandardRelateDeleteResponseBody {
	s.TraceId = &v
	return s
}

func (s *TravelStandardRelateDeleteResponseBody) Validate() error {
	return dara.Validate(s)
}
