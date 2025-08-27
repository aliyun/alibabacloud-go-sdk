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
	// example:
	//
	// 0
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// operation success.
	Message *string                             `json:"message,omitempty" xml:"message,omitempty"`
	Module  *HotelOrderCreateResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// example:
	//
	// C61ECFF6-606B-5F66-B81D-D77369043A5F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 21041ce316577904808056433edbb2
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
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
	return dara.Validate(s)
}

type HotelOrderCreateResponseBodyModule struct {
	// example:
	//
	// 123
	BtripOrderId *int64  `json:"btrip_order_id,omitempty" xml:"btrip_order_id,omitempty"`
	PaymentNo    *string `json:"payment_no,omitempty" xml:"payment_no,omitempty"`
	// example:
	//
	// 100
	TotalPrice *int64 `json:"total_price,omitempty" xml:"total_price,omitempty"`
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
