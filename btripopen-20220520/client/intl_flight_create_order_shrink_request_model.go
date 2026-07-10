// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightCreateOrderShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAsyncCreateOrderKey(v string) *IntlFlightCreateOrderShrinkRequest
	GetAsyncCreateOrderKey() *string
	SetAsyncCreateOrderMode(v bool) *IntlFlightCreateOrderShrinkRequest
	GetAsyncCreateOrderMode() *bool
	SetBtripUserId(v string) *IntlFlightCreateOrderShrinkRequest
	GetBtripUserId() *string
	SetBuyerName(v string) *IntlFlightCreateOrderShrinkRequest
	GetBuyerName() *string
	SetContactInfoShrink(v string) *IntlFlightCreateOrderShrinkRequest
	GetContactInfoShrink() *string
	SetExtraInfoShrink(v string) *IntlFlightCreateOrderShrinkRequest
	GetExtraInfoShrink() *string
	SetIsvName(v string) *IntlFlightCreateOrderShrinkRequest
	GetIsvName() *string
	SetOrderPrice(v int64) *IntlFlightCreateOrderShrinkRequest
	GetOrderPrice() *int64
	SetOtaItemId(v string) *IntlFlightCreateOrderShrinkRequest
	GetOtaItemId() *string
	SetOutOrderId(v string) *IntlFlightCreateOrderShrinkRequest
	GetOutOrderId() *string
	SetPassengerListShrink(v string) *IntlFlightCreateOrderShrinkRequest
	GetPassengerListShrink() *string
	SetRenderKey(v string) *IntlFlightCreateOrderShrinkRequest
	GetRenderKey() *string
}

type IntlFlightCreateOrderShrinkRequest struct {
	AsyncCreateOrderKey  *string `json:"async_create_order_key,omitempty" xml:"async_create_order_key,omitempty"`
	AsyncCreateOrderMode *bool   `json:"async_create_order_mode,omitempty" xml:"async_create_order_mode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10001
	BtripUserId *string `json:"btrip_user_id,omitempty" xml:"btrip_user_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ZHANG/SAN
	BuyerName *string `json:"buyer_name,omitempty" xml:"buyer_name,omitempty"`
	// This parameter is required.
	ContactInfoShrink *string `json:"contact_info,omitempty" xml:"contact_info,omitempty"`
	ExtraInfoShrink   *string `json:"extra_info,omitempty" xml:"extra_info,omitempty"`
	IsvName           *string `json:"isv_name,omitempty" xml:"isv_name,omitempty"`
	OrderPrice        *int64  `json:"order_price,omitempty" xml:"order_price,omitempty"`
	// This parameter is required.
	OtaItemId  *string `json:"ota_item_id,omitempty" xml:"ota_item_id,omitempty"`
	OutOrderId *string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// This parameter is required.
	PassengerListShrink *string `json:"passenger_list,omitempty" xml:"passenger_list,omitempty"`
	RenderKey           *string `json:"render_key,omitempty" xml:"render_key,omitempty"`
}

func (s IntlFlightCreateOrderShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightCreateOrderShrinkRequest) GoString() string {
	return s.String()
}

func (s *IntlFlightCreateOrderShrinkRequest) GetAsyncCreateOrderKey() *string {
	return s.AsyncCreateOrderKey
}

func (s *IntlFlightCreateOrderShrinkRequest) GetAsyncCreateOrderMode() *bool {
	return s.AsyncCreateOrderMode
}

func (s *IntlFlightCreateOrderShrinkRequest) GetBtripUserId() *string {
	return s.BtripUserId
}

func (s *IntlFlightCreateOrderShrinkRequest) GetBuyerName() *string {
	return s.BuyerName
}

func (s *IntlFlightCreateOrderShrinkRequest) GetContactInfoShrink() *string {
	return s.ContactInfoShrink
}

func (s *IntlFlightCreateOrderShrinkRequest) GetExtraInfoShrink() *string {
	return s.ExtraInfoShrink
}

func (s *IntlFlightCreateOrderShrinkRequest) GetIsvName() *string {
	return s.IsvName
}

func (s *IntlFlightCreateOrderShrinkRequest) GetOrderPrice() *int64 {
	return s.OrderPrice
}

func (s *IntlFlightCreateOrderShrinkRequest) GetOtaItemId() *string {
	return s.OtaItemId
}

func (s *IntlFlightCreateOrderShrinkRequest) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *IntlFlightCreateOrderShrinkRequest) GetPassengerListShrink() *string {
	return s.PassengerListShrink
}

func (s *IntlFlightCreateOrderShrinkRequest) GetRenderKey() *string {
	return s.RenderKey
}

func (s *IntlFlightCreateOrderShrinkRequest) SetAsyncCreateOrderKey(v string) *IntlFlightCreateOrderShrinkRequest {
	s.AsyncCreateOrderKey = &v
	return s
}

func (s *IntlFlightCreateOrderShrinkRequest) SetAsyncCreateOrderMode(v bool) *IntlFlightCreateOrderShrinkRequest {
	s.AsyncCreateOrderMode = &v
	return s
}

func (s *IntlFlightCreateOrderShrinkRequest) SetBtripUserId(v string) *IntlFlightCreateOrderShrinkRequest {
	s.BtripUserId = &v
	return s
}

func (s *IntlFlightCreateOrderShrinkRequest) SetBuyerName(v string) *IntlFlightCreateOrderShrinkRequest {
	s.BuyerName = &v
	return s
}

func (s *IntlFlightCreateOrderShrinkRequest) SetContactInfoShrink(v string) *IntlFlightCreateOrderShrinkRequest {
	s.ContactInfoShrink = &v
	return s
}

func (s *IntlFlightCreateOrderShrinkRequest) SetExtraInfoShrink(v string) *IntlFlightCreateOrderShrinkRequest {
	s.ExtraInfoShrink = &v
	return s
}

func (s *IntlFlightCreateOrderShrinkRequest) SetIsvName(v string) *IntlFlightCreateOrderShrinkRequest {
	s.IsvName = &v
	return s
}

func (s *IntlFlightCreateOrderShrinkRequest) SetOrderPrice(v int64) *IntlFlightCreateOrderShrinkRequest {
	s.OrderPrice = &v
	return s
}

func (s *IntlFlightCreateOrderShrinkRequest) SetOtaItemId(v string) *IntlFlightCreateOrderShrinkRequest {
	s.OtaItemId = &v
	return s
}

func (s *IntlFlightCreateOrderShrinkRequest) SetOutOrderId(v string) *IntlFlightCreateOrderShrinkRequest {
	s.OutOrderId = &v
	return s
}

func (s *IntlFlightCreateOrderShrinkRequest) SetPassengerListShrink(v string) *IntlFlightCreateOrderShrinkRequest {
	s.PassengerListShrink = &v
	return s
}

func (s *IntlFlightCreateOrderShrinkRequest) SetRenderKey(v string) *IntlFlightCreateOrderShrinkRequest {
	s.RenderKey = &v
	return s
}

func (s *IntlFlightCreateOrderShrinkRequest) Validate() error {
	return dara.Validate(s)
}
