// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOrderListResult interface {
	dara.Model
	String() string
	GoString() string
	SetOrderList(v []*OrderResult) *OrderListResult
	GetOrderList() []*OrderResult
	SetRequestId(v string) *OrderListResult
	GetRequestId() *string
	SetTotal(v int32) *OrderListResult
	GetTotal() *int32
}

type OrderListResult struct {
	OrderList []*OrderResult `json:"orderList,omitempty" xml:"orderList,omitempty" type:"Repeated"`
	// example:
	//
	// 3239281273464326823
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 24
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s OrderListResult) String() string {
	return dara.Prettify(s)
}

func (s OrderListResult) GoString() string {
	return s.String()
}

func (s *OrderListResult) GetOrderList() []*OrderResult {
	return s.OrderList
}

func (s *OrderListResult) GetRequestId() *string {
	return s.RequestId
}

func (s *OrderListResult) GetTotal() *int32 {
	return s.Total
}

func (s *OrderListResult) SetOrderList(v []*OrderResult) *OrderListResult {
	s.OrderList = v
	return s
}

func (s *OrderListResult) SetRequestId(v string) *OrderListResult {
	s.RequestId = &v
	return s
}

func (s *OrderListResult) SetTotal(v int32) *OrderListResult {
	s.Total = &v
	return s
}

func (s *OrderListResult) Validate() error {
	if s.OrderList != nil {
		for _, item := range s.OrderList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
