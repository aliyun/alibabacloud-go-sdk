// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightInventoryPriceCheckShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBtripUserId(v string) *IntlFlightInventoryPriceCheckShrinkRequest
	GetBtripUserId() *string
	SetBuyerName(v string) *IntlFlightInventoryPriceCheckShrinkRequest
	GetBuyerName() *string
	SetIsvName(v string) *IntlFlightInventoryPriceCheckShrinkRequest
	GetIsvName() *string
	SetOrderPrice(v int64) *IntlFlightInventoryPriceCheckShrinkRequest
	GetOrderPrice() *int64
	SetOtaItemId(v string) *IntlFlightInventoryPriceCheckShrinkRequest
	GetOtaItemId() *string
	SetPassengerListShrink(v string) *IntlFlightInventoryPriceCheckShrinkRequest
	GetPassengerListShrink() *string
}

type IntlFlightInventoryPriceCheckShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 10001
	BtripUserId *string `json:"btrip_user_id,omitempty" xml:"btrip_user_id,omitempty"`
	// example:
	//
	// ZHANG/SAN
	BuyerName *string `json:"buyer_name,omitempty" xml:"buyer_name,omitempty"`
	// example:
	//
	// ZJTD
	IsvName *string `json:"isv_name,omitempty" xml:"isv_name,omitempty"`
	// example:
	//
	// 102000
	OrderPrice *int64 `json:"order_price,omitempty" xml:"order_price,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 22111acaf9ea47c09ed0db6abc45be2d_0
	OtaItemId *string `json:"ota_item_id,omitempty" xml:"ota_item_id,omitempty"`
	// This parameter is required.
	PassengerListShrink *string `json:"passenger_list,omitempty" xml:"passenger_list,omitempty"`
}

func (s IntlFlightInventoryPriceCheckShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightInventoryPriceCheckShrinkRequest) GoString() string {
	return s.String()
}

func (s *IntlFlightInventoryPriceCheckShrinkRequest) GetBtripUserId() *string {
	return s.BtripUserId
}

func (s *IntlFlightInventoryPriceCheckShrinkRequest) GetBuyerName() *string {
	return s.BuyerName
}

func (s *IntlFlightInventoryPriceCheckShrinkRequest) GetIsvName() *string {
	return s.IsvName
}

func (s *IntlFlightInventoryPriceCheckShrinkRequest) GetOrderPrice() *int64 {
	return s.OrderPrice
}

func (s *IntlFlightInventoryPriceCheckShrinkRequest) GetOtaItemId() *string {
	return s.OtaItemId
}

func (s *IntlFlightInventoryPriceCheckShrinkRequest) GetPassengerListShrink() *string {
	return s.PassengerListShrink
}

func (s *IntlFlightInventoryPriceCheckShrinkRequest) SetBtripUserId(v string) *IntlFlightInventoryPriceCheckShrinkRequest {
	s.BtripUserId = &v
	return s
}

func (s *IntlFlightInventoryPriceCheckShrinkRequest) SetBuyerName(v string) *IntlFlightInventoryPriceCheckShrinkRequest {
	s.BuyerName = &v
	return s
}

func (s *IntlFlightInventoryPriceCheckShrinkRequest) SetIsvName(v string) *IntlFlightInventoryPriceCheckShrinkRequest {
	s.IsvName = &v
	return s
}

func (s *IntlFlightInventoryPriceCheckShrinkRequest) SetOrderPrice(v int64) *IntlFlightInventoryPriceCheckShrinkRequest {
	s.OrderPrice = &v
	return s
}

func (s *IntlFlightInventoryPriceCheckShrinkRequest) SetOtaItemId(v string) *IntlFlightInventoryPriceCheckShrinkRequest {
	s.OtaItemId = &v
	return s
}

func (s *IntlFlightInventoryPriceCheckShrinkRequest) SetPassengerListShrink(v string) *IntlFlightInventoryPriceCheckShrinkRequest {
	s.PassengerListShrink = &v
	return s
}

func (s *IntlFlightInventoryPriceCheckShrinkRequest) Validate() error {
	return dara.Validate(s)
}
