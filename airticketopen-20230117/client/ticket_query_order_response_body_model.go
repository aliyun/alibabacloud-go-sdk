// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketQueryOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *TicketQueryOrderResponseBodyData) *TicketQueryOrderResponseBody
	GetData() *TicketQueryOrderResponseBodyData
	SetErrorCode(v string) *TicketQueryOrderResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *TicketQueryOrderResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *TicketQueryOrderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TicketQueryOrderResponseBody
	GetSuccess() *bool
}

type TicketQueryOrderResponseBody struct {
	Data *TicketQueryOrderResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// DistributorOrderIdInvalid
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 分销商订单号不合法
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// 51593418-8C73-5E47-8BA8-3F1D4A00CC0B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s TicketQueryOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TicketQueryOrderResponseBody) GoString() string {
	return s.String()
}

func (s *TicketQueryOrderResponseBody) GetData() *TicketQueryOrderResponseBodyData {
	return s.Data
}

func (s *TicketQueryOrderResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *TicketQueryOrderResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *TicketQueryOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TicketQueryOrderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TicketQueryOrderResponseBody) SetData(v *TicketQueryOrderResponseBodyData) *TicketQueryOrderResponseBody {
	s.Data = v
	return s
}

func (s *TicketQueryOrderResponseBody) SetErrorCode(v string) *TicketQueryOrderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *TicketQueryOrderResponseBody) SetErrorMsg(v string) *TicketQueryOrderResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *TicketQueryOrderResponseBody) SetRequestId(v string) *TicketQueryOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *TicketQueryOrderResponseBody) SetSuccess(v bool) *TicketQueryOrderResponseBody {
	s.Success = &v
	return s
}

func (s *TicketQueryOrderResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TicketQueryOrderResponseBodyData struct {
	Order    *TicketQueryOrderResponseBodyDataOrder      `json:"Order,omitempty" xml:"Order,omitempty" type:"Struct"`
	Vouchers []*TicketQueryOrderResponseBodyDataVouchers `json:"Vouchers,omitempty" xml:"Vouchers,omitempty" type:"Repeated"`
}

func (s TicketQueryOrderResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s TicketQueryOrderResponseBodyData) GoString() string {
	return s.String()
}

func (s *TicketQueryOrderResponseBodyData) GetOrder() *TicketQueryOrderResponseBodyDataOrder {
	return s.Order
}

func (s *TicketQueryOrderResponseBodyData) GetVouchers() []*TicketQueryOrderResponseBodyDataVouchers {
	return s.Vouchers
}

func (s *TicketQueryOrderResponseBodyData) SetOrder(v *TicketQueryOrderResponseBodyDataOrder) *TicketQueryOrderResponseBodyData {
	s.Order = v
	return s
}

func (s *TicketQueryOrderResponseBodyData) SetVouchers(v []*TicketQueryOrderResponseBodyDataVouchers) *TicketQueryOrderResponseBodyData {
	s.Vouchers = v
	return s
}

func (s *TicketQueryOrderResponseBodyData) Validate() error {
	if s.Order != nil {
		if err := s.Order.Validate(); err != nil {
			return err
		}
	}
	if s.Vouchers != nil {
		for _, item := range s.Vouchers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type TicketQueryOrderResponseBodyDataOrder struct {
	// example:
	//
	// 1
	FundStatus *int32 `json:"FundStatus,omitempty" xml:"FundStatus,omitempty"`
	// example:
	//
	// 123456
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// 1
	OrderStatus *int32 `json:"OrderStatus,omitempty" xml:"OrderStatus,omitempty"`
}

func (s TicketQueryOrderResponseBodyDataOrder) String() string {
	return dara.Prettify(s)
}

func (s TicketQueryOrderResponseBodyDataOrder) GoString() string {
	return s.String()
}

func (s *TicketQueryOrderResponseBodyDataOrder) GetFundStatus() *int32 {
	return s.FundStatus
}

func (s *TicketQueryOrderResponseBodyDataOrder) GetOrderId() *string {
	return s.OrderId
}

func (s *TicketQueryOrderResponseBodyDataOrder) GetOrderStatus() *int32 {
	return s.OrderStatus
}

func (s *TicketQueryOrderResponseBodyDataOrder) SetFundStatus(v int32) *TicketQueryOrderResponseBodyDataOrder {
	s.FundStatus = &v
	return s
}

func (s *TicketQueryOrderResponseBodyDataOrder) SetOrderId(v string) *TicketQueryOrderResponseBodyDataOrder {
	s.OrderId = &v
	return s
}

func (s *TicketQueryOrderResponseBodyDataOrder) SetOrderStatus(v int32) *TicketQueryOrderResponseBodyDataOrder {
	s.OrderStatus = &v
	return s
}

func (s *TicketQueryOrderResponseBodyDataOrder) Validate() error {
	return dara.Validate(s)
}

type TicketQueryOrderResponseBodyDataVouchers struct {
	// example:
	//
	// 1234567890
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 10
	TotalTimes *int32 `json:"TotalTimes,omitempty" xml:"TotalTimes,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// https://www.alitrip.com/1234567890.png
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s TicketQueryOrderResponseBodyDataVouchers) String() string {
	return dara.Prettify(s)
}

func (s TicketQueryOrderResponseBodyDataVouchers) GoString() string {
	return s.String()
}

func (s *TicketQueryOrderResponseBodyDataVouchers) GetCode() *string {
	return s.Code
}

func (s *TicketQueryOrderResponseBodyDataVouchers) GetTotalTimes() *int32 {
	return s.TotalTimes
}

func (s *TicketQueryOrderResponseBodyDataVouchers) GetType() *int32 {
	return s.Type
}

func (s *TicketQueryOrderResponseBodyDataVouchers) GetUrl() *string {
	return s.Url
}

func (s *TicketQueryOrderResponseBodyDataVouchers) SetCode(v string) *TicketQueryOrderResponseBodyDataVouchers {
	s.Code = &v
	return s
}

func (s *TicketQueryOrderResponseBodyDataVouchers) SetTotalTimes(v int32) *TicketQueryOrderResponseBodyDataVouchers {
	s.TotalTimes = &v
	return s
}

func (s *TicketQueryOrderResponseBodyDataVouchers) SetType(v int32) *TicketQueryOrderResponseBodyDataVouchers {
	s.Type = &v
	return s
}

func (s *TicketQueryOrderResponseBodyDataVouchers) SetUrl(v string) *TicketQueryOrderResponseBodyDataVouchers {
	s.Url = &v
	return s
}

func (s *TicketQueryOrderResponseBodyDataVouchers) Validate() error {
	return dara.Validate(s)
}
