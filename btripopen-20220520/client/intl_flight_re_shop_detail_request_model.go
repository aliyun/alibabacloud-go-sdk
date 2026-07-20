// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightReShopDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *IntlFlightReShopDetailRequest
	GetOrderId() *string
	SetOutOrderId(v string) *IntlFlightReShopDetailRequest
	GetOutOrderId() *string
	SetOutReShopApplyId(v string) *IntlFlightReShopDetailRequest
	GetOutReShopApplyId() *string
	SetReShopApplyId(v string) *IntlFlightReShopDetailRequest
	GetReShopApplyId() *string
}

type IntlFlightReShopDetailRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1017035200254689390
	OrderId *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// example:
	//
	// F11494210548838170624
	OutOrderId *string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// example:
	//
	// M132492719472
	OutReShopApplyId *string `json:"out_re_shop_apply_id,omitempty" xml:"out_re_shop_apply_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10002612002
	ReShopApplyId *string `json:"re_shop_apply_id,omitempty" xml:"re_shop_apply_id,omitempty"`
}

func (s IntlFlightReShopDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopDetailRequest) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopDetailRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *IntlFlightReShopDetailRequest) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *IntlFlightReShopDetailRequest) GetOutReShopApplyId() *string {
	return s.OutReShopApplyId
}

func (s *IntlFlightReShopDetailRequest) GetReShopApplyId() *string {
	return s.ReShopApplyId
}

func (s *IntlFlightReShopDetailRequest) SetOrderId(v string) *IntlFlightReShopDetailRequest {
	s.OrderId = &v
	return s
}

func (s *IntlFlightReShopDetailRequest) SetOutOrderId(v string) *IntlFlightReShopDetailRequest {
	s.OutOrderId = &v
	return s
}

func (s *IntlFlightReShopDetailRequest) SetOutReShopApplyId(v string) *IntlFlightReShopDetailRequest {
	s.OutReShopApplyId = &v
	return s
}

func (s *IntlFlightReShopDetailRequest) SetReShopApplyId(v string) *IntlFlightReShopDetailRequest {
	s.ReShopApplyId = &v
	return s
}

func (s *IntlFlightReShopDetailRequest) Validate() error {
	return dara.Validate(s)
}
