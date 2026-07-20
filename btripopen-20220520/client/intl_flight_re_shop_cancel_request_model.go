// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightReShopCancelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *IntlFlightReShopCancelRequest
	GetOrderId() *string
	SetOutOrderId(v string) *IntlFlightReShopCancelRequest
	GetOutOrderId() *string
	SetOutReShopApplyId(v string) *IntlFlightReShopCancelRequest
	GetOutReShopApplyId() *string
	SetReShopApplyId(v string) *IntlFlightReShopCancelRequest
	GetReShopApplyId() *string
}

type IntlFlightReShopCancelRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1017035199763856322
	OrderId *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// example:
	//
	// JPT2025032400000001
	OutOrderId *string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// example:
	//
	// JPM20241024354
	OutReShopApplyId *string `json:"out_re_shop_apply_id,omitempty" xml:"out_re_shop_apply_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1019202345889633
	ReShopApplyId *string `json:"re_shop_apply_id,omitempty" xml:"re_shop_apply_id,omitempty"`
}

func (s IntlFlightReShopCancelRequest) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopCancelRequest) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopCancelRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *IntlFlightReShopCancelRequest) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *IntlFlightReShopCancelRequest) GetOutReShopApplyId() *string {
	return s.OutReShopApplyId
}

func (s *IntlFlightReShopCancelRequest) GetReShopApplyId() *string {
	return s.ReShopApplyId
}

func (s *IntlFlightReShopCancelRequest) SetOrderId(v string) *IntlFlightReShopCancelRequest {
	s.OrderId = &v
	return s
}

func (s *IntlFlightReShopCancelRequest) SetOutOrderId(v string) *IntlFlightReShopCancelRequest {
	s.OutOrderId = &v
	return s
}

func (s *IntlFlightReShopCancelRequest) SetOutReShopApplyId(v string) *IntlFlightReShopCancelRequest {
	s.OutReShopApplyId = &v
	return s
}

func (s *IntlFlightReShopCancelRequest) SetReShopApplyId(v string) *IntlFlightReShopCancelRequest {
	s.ReShopApplyId = &v
	return s
}

func (s *IntlFlightReShopCancelRequest) Validate() error {
	return dara.Validate(s)
}
