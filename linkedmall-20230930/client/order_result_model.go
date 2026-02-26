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
	// example:
	//
	// 2023-09-01T12:00:00.000Z
	CreateDate *string `json:"createDate,omitempty" xml:"createDate,omitempty"`
	// example:
	//
	// 12****01
	DistributorId *string `json:"distributorId,omitempty" xml:"distributorId,omitempty"`
	// example:
	//
	// 1
	LogisticsStatus *string `json:"logisticsStatus,omitempty" xml:"logisticsStatus,omitempty"`
	// example:
	//
	// 100
	OrderAmount *int64 `json:"orderAmount,omitempty" xml:"orderAmount,omitempty"`
	// example:
	//
	// 系统关单
	OrderClosedReason *string `json:"orderClosedReason,omitempty" xml:"orderClosedReason,omitempty"`
	// example:
	//
	// 6692****5457
	OrderId       *string            `json:"orderId,omitempty" xml:"orderId,omitempty"`
	OrderLineList []*OrderLineResult `json:"orderLineList,omitempty" xml:"orderLineList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	OrderStatus *string `json:"orderStatus,omitempty" xml:"orderStatus,omitempty"`
	// example:
	//
	// D12***111
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
