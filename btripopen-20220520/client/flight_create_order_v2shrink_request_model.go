// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightCreateOrderV2ShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAsyncCreateOrderKey(v string) *FlightCreateOrderV2ShrinkRequest
	GetAsyncCreateOrderKey() *string
	SetAsyncCreateOrderMode(v bool) *FlightCreateOrderV2ShrinkRequest
	GetAsyncCreateOrderMode() *bool
	SetBtripUserId(v string) *FlightCreateOrderV2ShrinkRequest
	GetBtripUserId() *string
	SetBuyerName(v string) *FlightCreateOrderV2ShrinkRequest
	GetBuyerName() *string
	SetContactInfoShrink(v string) *FlightCreateOrderV2ShrinkRequest
	GetContactInfoShrink() *string
	SetIsvName(v string) *FlightCreateOrderV2ShrinkRequest
	GetIsvName() *string
	SetOtaItemId(v string) *FlightCreateOrderV2ShrinkRequest
	GetOtaItemId() *string
	SetOutOrderId(v string) *FlightCreateOrderV2ShrinkRequest
	GetOutOrderId() *string
	SetTotalPriceCent(v int64) *FlightCreateOrderV2ShrinkRequest
	GetTotalPriceCent() *int64
	SetTravelersShrink(v string) *FlightCreateOrderV2ShrinkRequest
	GetTravelersShrink() *string
}

type FlightCreateOrderV2ShrinkRequest struct {
	AsyncCreateOrderKey *string `json:"async_create_order_key,omitempty" xml:"async_create_order_key,omitempty"`
	// example:
	//
	// false
	AsyncCreateOrderMode *bool   `json:"async_create_order_mode,omitempty" xml:"async_create_order_mode,omitempty"`
	BtripUserId          *string `json:"btrip_user_id,omitempty" xml:"btrip_user_id,omitempty"`
	BuyerName            *string `json:"buyer_name,omitempty" xml:"buyer_name,omitempty"`
	// This parameter is required.
	ContactInfoShrink *string `json:"contact_info,omitempty" xml:"contact_info,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cheshiapi
	IsvName *string `json:"isv_name,omitempty" xml:"isv_name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 7fb731deeb4510b86c17e8c8c25740_11
	OtaItemId *string `json:"ota_item_id,omitempty" xml:"ota_item_id,omitempty"`
	// This parameter is required.
	OutOrderId     *string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	TotalPriceCent *int64  `json:"total_price_cent,omitempty" xml:"total_price_cent,omitempty"`
	// This parameter is required.
	TravelersShrink *string `json:"travelers,omitempty" xml:"travelers,omitempty"`
}

func (s FlightCreateOrderV2ShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s FlightCreateOrderV2ShrinkRequest) GoString() string {
	return s.String()
}

func (s *FlightCreateOrderV2ShrinkRequest) GetAsyncCreateOrderKey() *string {
	return s.AsyncCreateOrderKey
}

func (s *FlightCreateOrderV2ShrinkRequest) GetAsyncCreateOrderMode() *bool {
	return s.AsyncCreateOrderMode
}

func (s *FlightCreateOrderV2ShrinkRequest) GetBtripUserId() *string {
	return s.BtripUserId
}

func (s *FlightCreateOrderV2ShrinkRequest) GetBuyerName() *string {
	return s.BuyerName
}

func (s *FlightCreateOrderV2ShrinkRequest) GetContactInfoShrink() *string {
	return s.ContactInfoShrink
}

func (s *FlightCreateOrderV2ShrinkRequest) GetIsvName() *string {
	return s.IsvName
}

func (s *FlightCreateOrderV2ShrinkRequest) GetOtaItemId() *string {
	return s.OtaItemId
}

func (s *FlightCreateOrderV2ShrinkRequest) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *FlightCreateOrderV2ShrinkRequest) GetTotalPriceCent() *int64 {
	return s.TotalPriceCent
}

func (s *FlightCreateOrderV2ShrinkRequest) GetTravelersShrink() *string {
	return s.TravelersShrink
}

func (s *FlightCreateOrderV2ShrinkRequest) SetAsyncCreateOrderKey(v string) *FlightCreateOrderV2ShrinkRequest {
	s.AsyncCreateOrderKey = &v
	return s
}

func (s *FlightCreateOrderV2ShrinkRequest) SetAsyncCreateOrderMode(v bool) *FlightCreateOrderV2ShrinkRequest {
	s.AsyncCreateOrderMode = &v
	return s
}

func (s *FlightCreateOrderV2ShrinkRequest) SetBtripUserId(v string) *FlightCreateOrderV2ShrinkRequest {
	s.BtripUserId = &v
	return s
}

func (s *FlightCreateOrderV2ShrinkRequest) SetBuyerName(v string) *FlightCreateOrderV2ShrinkRequest {
	s.BuyerName = &v
	return s
}

func (s *FlightCreateOrderV2ShrinkRequest) SetContactInfoShrink(v string) *FlightCreateOrderV2ShrinkRequest {
	s.ContactInfoShrink = &v
	return s
}

func (s *FlightCreateOrderV2ShrinkRequest) SetIsvName(v string) *FlightCreateOrderV2ShrinkRequest {
	s.IsvName = &v
	return s
}

func (s *FlightCreateOrderV2ShrinkRequest) SetOtaItemId(v string) *FlightCreateOrderV2ShrinkRequest {
	s.OtaItemId = &v
	return s
}

func (s *FlightCreateOrderV2ShrinkRequest) SetOutOrderId(v string) *FlightCreateOrderV2ShrinkRequest {
	s.OutOrderId = &v
	return s
}

func (s *FlightCreateOrderV2ShrinkRequest) SetTotalPriceCent(v int64) *FlightCreateOrderV2ShrinkRequest {
	s.TotalPriceCent = &v
	return s
}

func (s *FlightCreateOrderV2ShrinkRequest) SetTravelersShrink(v string) *FlightCreateOrderV2ShrinkRequest {
	s.TravelersShrink = &v
	return s
}

func (s *FlightCreateOrderV2ShrinkRequest) Validate() error {
	return dara.Validate(s)
}
