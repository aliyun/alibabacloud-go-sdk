// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightReShopPayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *IntlFlightReShopPayRequest
	GetOrderId() *string
	SetOutOrderId(v string) *IntlFlightReShopPayRequest
	GetOutOrderId() *string
	SetOutReShopApplyId(v string) *IntlFlightReShopPayRequest
	GetOutReShopApplyId() *string
	SetReShopApplyId(v string) *IntlFlightReShopPayRequest
	GetReShopApplyId() *string
}

type IntlFlightReShopPayRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1017035200059399795
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

func (s IntlFlightReShopPayRequest) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopPayRequest) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopPayRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *IntlFlightReShopPayRequest) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *IntlFlightReShopPayRequest) GetOutReShopApplyId() *string {
	return s.OutReShopApplyId
}

func (s *IntlFlightReShopPayRequest) GetReShopApplyId() *string {
	return s.ReShopApplyId
}

func (s *IntlFlightReShopPayRequest) SetOrderId(v string) *IntlFlightReShopPayRequest {
	s.OrderId = &v
	return s
}

func (s *IntlFlightReShopPayRequest) SetOutOrderId(v string) *IntlFlightReShopPayRequest {
	s.OutOrderId = &v
	return s
}

func (s *IntlFlightReShopPayRequest) SetOutReShopApplyId(v string) *IntlFlightReShopPayRequest {
	s.OutReShopApplyId = &v
	return s
}

func (s *IntlFlightReShopPayRequest) SetReShopApplyId(v string) *IntlFlightReShopPayRequest {
	s.ReShopApplyId = &v
	return s
}

func (s *IntlFlightReShopPayRequest) Validate() error {
	return dara.Validate(s)
}
