// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOrderResult interface {
	dara.Model
	String() string
	GoString() string
	SetCreateDate(v string) *OrderResult
	GetCreateDate() *string
	SetDistributorId(v string) *OrderResult
	GetDistributorId() *string
	SetLogisticsStatus(v string) *OrderResult
	GetLogisticsStatus() *string
	SetOrderAmount(v int64) *OrderResult
	GetOrderAmount() *int64
	SetOrderClosedReason(v string) *OrderResult
	GetOrderClosedReason() *string
	SetOrderId(v string) *OrderResult
	GetOrderId() *string
	SetOrderLineList(v []*OrderLineResult) *OrderResult
	GetOrderLineList() []*OrderLineResult
	SetOrderStatus(v string) *OrderResult
	GetOrderStatus() *string
	SetRequestId(v string) *OrderResult
	GetRequestId() *string
}

type OrderResult struct {
	// The order creation time.
	//
	// example:
	//
	// 2023-09-11T12:22:24.000+08:00
	CreateDate *string `json:"createDate,omitempty" xml:"createDate,omitempty"`
	// The distributor ID.
	//
	// example:
	//
	// 12****01
	DistributorId *string `json:"distributorId,omitempty" xml:"distributorId,omitempty"`
	// The logistics status. Valid values: 1 (Awaiting Seller\\"s Shipment), 2 (Awaiting Buyer\\"s Confirmation), 3 (Received), 4 (Returned), 5 (Partially Received), 6 (Partially Shipped), and 8 (Logistics Order Not Created).
	//
	// example:
	//
	// 1
	LogisticsStatus *string `json:"logisticsStatus,omitempty" xml:"logisticsStatus,omitempty"`
	// The order amount, in cents.
	//
	// example:
	//
	// 100
	OrderAmount *int64 `json:"orderAmount,omitempty" xml:"orderAmount,omitempty"`
	// The reason the order was closed.
	//
	// example:
	//
	// 系统关单
	OrderClosedReason *string `json:"orderClosedReason,omitempty" xml:"orderClosedReason,omitempty"`
	// The ID of the main order.
	//
	// example:
	//
	// 6692****5457
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
	// The list of sub-orders.
	OrderLineList []*OrderLineResult `json:"orderLineList,omitempty" xml:"orderLineList,omitempty" type:"Repeated"`
	// The order status. Valid values: 1 (Pending Payment), 2 (Paid), 4 (Closed with Refund), 6 (Transaction Successful), and 8 (Closed).
	//
	// example:
	//
	// 1
	OrderStatus *string `json:"orderStatus,omitempty" xml:"orderStatus,omitempty"`
	// The unique identifier for the request.
	//
	// example:
	//
	// 841471F6-5D61-1331-8C38-2****B55
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s OrderResult) String() string {
	return dara.Prettify(s)
}

func (s OrderResult) GoString() string {
	return s.String()
}

func (s *OrderResult) GetCreateDate() *string {
	return s.CreateDate
}

func (s *OrderResult) GetDistributorId() *string {
	return s.DistributorId
}

func (s *OrderResult) GetLogisticsStatus() *string {
	return s.LogisticsStatus
}

func (s *OrderResult) GetOrderAmount() *int64 {
	return s.OrderAmount
}

func (s *OrderResult) GetOrderClosedReason() *string {
	return s.OrderClosedReason
}

func (s *OrderResult) GetOrderId() *string {
	return s.OrderId
}

func (s *OrderResult) GetOrderLineList() []*OrderLineResult {
	return s.OrderLineList
}

func (s *OrderResult) GetOrderStatus() *string {
	return s.OrderStatus
}

func (s *OrderResult) GetRequestId() *string {
	return s.RequestId
}

func (s *OrderResult) SetCreateDate(v string) *OrderResult {
	s.CreateDate = &v
	return s
}

func (s *OrderResult) SetDistributorId(v string) *OrderResult {
	s.DistributorId = &v
	return s
}

func (s *OrderResult) SetLogisticsStatus(v string) *OrderResult {
	s.LogisticsStatus = &v
	return s
}

func (s *OrderResult) SetOrderAmount(v int64) *OrderResult {
	s.OrderAmount = &v
	return s
}

func (s *OrderResult) SetOrderClosedReason(v string) *OrderResult {
	s.OrderClosedReason = &v
	return s
}

func (s *OrderResult) SetOrderId(v string) *OrderResult {
	s.OrderId = &v
	return s
}

func (s *OrderResult) SetOrderLineList(v []*OrderLineResult) *OrderResult {
	s.OrderLineList = v
	return s
}

func (s *OrderResult) SetOrderStatus(v string) *OrderResult {
	s.OrderStatus = &v
	return s
}

func (s *OrderResult) SetRequestId(v string) *OrderResult {
	s.RequestId = &v
	return s
}

func (s *OrderResult) Validate() error {
	if s.OrderLineList != nil {
		for _, item := range s.OrderLineList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
