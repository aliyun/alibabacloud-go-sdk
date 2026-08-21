// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTicketQueryRefundOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *TicketQueryRefundOrderResponseBodyData) *TicketQueryRefundOrderResponseBody
	GetData() *TicketQueryRefundOrderResponseBodyData
	SetErrorCode(v string) *TicketQueryRefundOrderResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *TicketQueryRefundOrderResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *TicketQueryRefundOrderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TicketQueryRefundOrderResponseBody
	GetSuccess() *bool
}

type TicketQueryRefundOrderResponseBody struct {
	Data *TicketQueryRefundOrderResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s TicketQueryRefundOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TicketQueryRefundOrderResponseBody) GoString() string {
	return s.String()
}

func (s *TicketQueryRefundOrderResponseBody) GetData() *TicketQueryRefundOrderResponseBodyData {
	return s.Data
}

func (s *TicketQueryRefundOrderResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *TicketQueryRefundOrderResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *TicketQueryRefundOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TicketQueryRefundOrderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TicketQueryRefundOrderResponseBody) SetData(v *TicketQueryRefundOrderResponseBodyData) *TicketQueryRefundOrderResponseBody {
	s.Data = v
	return s
}

func (s *TicketQueryRefundOrderResponseBody) SetErrorCode(v string) *TicketQueryRefundOrderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *TicketQueryRefundOrderResponseBody) SetErrorMsg(v string) *TicketQueryRefundOrderResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *TicketQueryRefundOrderResponseBody) SetRequestId(v string) *TicketQueryRefundOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *TicketQueryRefundOrderResponseBody) SetSuccess(v bool) *TicketQueryRefundOrderResponseBody {
	s.Success = &v
	return s
}

func (s *TicketQueryRefundOrderResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TicketQueryRefundOrderResponseBodyData struct {
	RefundOrders []*TicketQueryRefundOrderResponseBodyDataRefundOrders `json:"RefundOrders,omitempty" xml:"RefundOrders,omitempty" type:"Repeated"`
}

func (s TicketQueryRefundOrderResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s TicketQueryRefundOrderResponseBodyData) GoString() string {
	return s.String()
}

func (s *TicketQueryRefundOrderResponseBodyData) GetRefundOrders() []*TicketQueryRefundOrderResponseBodyDataRefundOrders {
	return s.RefundOrders
}

func (s *TicketQueryRefundOrderResponseBodyData) SetRefundOrders(v []*TicketQueryRefundOrderResponseBodyDataRefundOrders) *TicketQueryRefundOrderResponseBodyData {
	s.RefundOrders = v
	return s
}

func (s *TicketQueryRefundOrderResponseBodyData) Validate() error {
	if s.RefundOrders != nil {
		for _, item := range s.RefundOrders {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type TicketQueryRefundOrderResponseBodyDataRefundOrders struct {
	// example:
	//
	// 1
	FundStatus *int32 `json:"FundStatus,omitempty" xml:"FundStatus,omitempty"`
	// example:
	//
	// 1
	OrderStatus *int32 `json:"OrderStatus,omitempty" xml:"OrderStatus,omitempty"`
}

func (s TicketQueryRefundOrderResponseBodyDataRefundOrders) String() string {
	return dara.Prettify(s)
}

func (s TicketQueryRefundOrderResponseBodyDataRefundOrders) GoString() string {
	return s.String()
}

func (s *TicketQueryRefundOrderResponseBodyDataRefundOrders) GetFundStatus() *int32 {
	return s.FundStatus
}

func (s *TicketQueryRefundOrderResponseBodyDataRefundOrders) GetOrderStatus() *int32 {
	return s.OrderStatus
}

func (s *TicketQueryRefundOrderResponseBodyDataRefundOrders) SetFundStatus(v int32) *TicketQueryRefundOrderResponseBodyDataRefundOrders {
	s.FundStatus = &v
	return s
}

func (s *TicketQueryRefundOrderResponseBodyDataRefundOrders) SetOrderStatus(v int32) *TicketQueryRefundOrderResponseBodyDataRefundOrders {
	s.OrderStatus = &v
	return s
}

func (s *TicketQueryRefundOrderResponseBodyDataRefundOrders) Validate() error {
	return dara.Validate(s)
}
