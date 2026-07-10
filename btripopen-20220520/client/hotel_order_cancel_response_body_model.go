// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelOrderCancelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *HotelOrderCancelResponseBody
	GetCode() *string
	SetMessage(v string) *HotelOrderCancelResponseBody
	GetMessage() *string
	SetModule(v *HotelOrderCancelResponseBodyModule) *HotelOrderCancelResponseBody
	GetModule() *HotelOrderCancelResponseBodyModule
	SetRequestId(v string) *HotelOrderCancelResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *HotelOrderCancelResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *HotelOrderCancelResponseBody
	GetTraceId() *string
}

type HotelOrderCancelResponseBody struct {
	Code      *string                             `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                             `json:"message,omitempty" xml:"message,omitempty"`
	Module    *HotelOrderCancelResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	RequestId *string                             `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                               `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string                             `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s HotelOrderCancelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderCancelResponseBody) GoString() string {
	return s.String()
}

func (s *HotelOrderCancelResponseBody) GetCode() *string {
	return s.Code
}

func (s *HotelOrderCancelResponseBody) GetMessage() *string {
	return s.Message
}

func (s *HotelOrderCancelResponseBody) GetModule() *HotelOrderCancelResponseBodyModule {
	return s.Module
}

func (s *HotelOrderCancelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *HotelOrderCancelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *HotelOrderCancelResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *HotelOrderCancelResponseBody) SetCode(v string) *HotelOrderCancelResponseBody {
	s.Code = &v
	return s
}

func (s *HotelOrderCancelResponseBody) SetMessage(v string) *HotelOrderCancelResponseBody {
	s.Message = &v
	return s
}

func (s *HotelOrderCancelResponseBody) SetModule(v *HotelOrderCancelResponseBodyModule) *HotelOrderCancelResponseBody {
	s.Module = v
	return s
}

func (s *HotelOrderCancelResponseBody) SetRequestId(v string) *HotelOrderCancelResponseBody {
	s.RequestId = &v
	return s
}

func (s *HotelOrderCancelResponseBody) SetSuccess(v bool) *HotelOrderCancelResponseBody {
	s.Success = &v
	return s
}

func (s *HotelOrderCancelResponseBody) SetTraceId(v string) *HotelOrderCancelResponseBody {
	s.TraceId = &v
	return s
}

func (s *HotelOrderCancelResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type HotelOrderCancelResponseBodyModule struct {
	CancelSuccess *bool   `json:"cancel_success,omitempty" xml:"cancel_success,omitempty"`
	Code          *string `json:"code,omitempty" xml:"code,omitempty"`
	Desc          *string `json:"desc,omitempty" xml:"desc,omitempty"`
	ForfeitFee    *int64  `json:"forfeit_fee,omitempty" xml:"forfeit_fee,omitempty"`
}

func (s HotelOrderCancelResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderCancelResponseBodyModule) GoString() string {
	return s.String()
}

func (s *HotelOrderCancelResponseBodyModule) GetCancelSuccess() *bool {
	return s.CancelSuccess
}

func (s *HotelOrderCancelResponseBodyModule) GetCode() *string {
	return s.Code
}

func (s *HotelOrderCancelResponseBodyModule) GetDesc() *string {
	return s.Desc
}

func (s *HotelOrderCancelResponseBodyModule) GetForfeitFee() *int64 {
	return s.ForfeitFee
}

func (s *HotelOrderCancelResponseBodyModule) SetCancelSuccess(v bool) *HotelOrderCancelResponseBodyModule {
	s.CancelSuccess = &v
	return s
}

func (s *HotelOrderCancelResponseBodyModule) SetCode(v string) *HotelOrderCancelResponseBodyModule {
	s.Code = &v
	return s
}

func (s *HotelOrderCancelResponseBodyModule) SetDesc(v string) *HotelOrderCancelResponseBodyModule {
	s.Desc = &v
	return s
}

func (s *HotelOrderCancelResponseBodyModule) SetForfeitFee(v int64) *HotelOrderCancelResponseBodyModule {
	s.ForfeitFee = &v
	return s
}

func (s *HotelOrderCancelResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
