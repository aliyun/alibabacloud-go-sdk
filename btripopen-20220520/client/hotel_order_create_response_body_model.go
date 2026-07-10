// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelOrderCreateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *HotelOrderCreateResponseBody
	GetCode() *string
	SetMessage(v string) *HotelOrderCreateResponseBody
	GetMessage() *string
	SetModule(v *HotelOrderCreateResponseBodyModule) *HotelOrderCreateResponseBody
	GetModule() *HotelOrderCreateResponseBodyModule
	SetRequestId(v string) *HotelOrderCreateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *HotelOrderCreateResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *HotelOrderCreateResponseBody
	GetTraceId() *string
}

type HotelOrderCreateResponseBody struct {
	Code      *string                             `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                             `json:"message,omitempty" xml:"message,omitempty"`
	Module    *HotelOrderCreateResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	RequestId *string                             `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                               `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string                             `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s HotelOrderCreateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderCreateResponseBody) GoString() string {
	return s.String()
}

func (s *HotelOrderCreateResponseBody) GetCode() *string {
	return s.Code
}

func (s *HotelOrderCreateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *HotelOrderCreateResponseBody) GetModule() *HotelOrderCreateResponseBodyModule {
	return s.Module
}

func (s *HotelOrderCreateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *HotelOrderCreateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *HotelOrderCreateResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *HotelOrderCreateResponseBody) SetCode(v string) *HotelOrderCreateResponseBody {
	s.Code = &v
	return s
}

func (s *HotelOrderCreateResponseBody) SetMessage(v string) *HotelOrderCreateResponseBody {
	s.Message = &v
	return s
}

func (s *HotelOrderCreateResponseBody) SetModule(v *HotelOrderCreateResponseBodyModule) *HotelOrderCreateResponseBody {
	s.Module = v
	return s
}

func (s *HotelOrderCreateResponseBody) SetRequestId(v string) *HotelOrderCreateResponseBody {
	s.RequestId = &v
	return s
}

func (s *HotelOrderCreateResponseBody) SetSuccess(v bool) *HotelOrderCreateResponseBody {
	s.Success = &v
	return s
}

func (s *HotelOrderCreateResponseBody) SetTraceId(v string) *HotelOrderCreateResponseBody {
	s.TraceId = &v
	return s
}

func (s *HotelOrderCreateResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type HotelOrderCreateResponseBodyModule struct {
	BtripOrderId *int64  `json:"btrip_order_id,omitempty" xml:"btrip_order_id,omitempty"`
	PaymentNo    *string `json:"payment_no,omitempty" xml:"payment_no,omitempty"`
	TotalPrice   *int64  `json:"total_price,omitempty" xml:"total_price,omitempty"`
}

func (s HotelOrderCreateResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderCreateResponseBodyModule) GoString() string {
	return s.String()
}

func (s *HotelOrderCreateResponseBodyModule) GetBtripOrderId() *int64 {
	return s.BtripOrderId
}

func (s *HotelOrderCreateResponseBodyModule) GetPaymentNo() *string {
	return s.PaymentNo
}

func (s *HotelOrderCreateResponseBodyModule) GetTotalPrice() *int64 {
	return s.TotalPrice
}

func (s *HotelOrderCreateResponseBodyModule) SetBtripOrderId(v int64) *HotelOrderCreateResponseBodyModule {
	s.BtripOrderId = &v
	return s
}

func (s *HotelOrderCreateResponseBodyModule) SetPaymentNo(v string) *HotelOrderCreateResponseBodyModule {
	s.PaymentNo = &v
	return s
}

func (s *HotelOrderCreateResponseBodyModule) SetTotalPrice(v int64) *HotelOrderCreateResponseBodyModule {
	s.TotalPrice = &v
	return s
}

func (s *HotelOrderCreateResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
