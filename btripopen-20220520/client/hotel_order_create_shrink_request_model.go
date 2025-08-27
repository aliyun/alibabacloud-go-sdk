// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelOrderCreateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBtripUserId(v string) *HotelOrderCreateShrinkRequest
	GetBtripUserId() *string
	SetCheckIn(v string) *HotelOrderCreateShrinkRequest
	GetCheckIn() *string
	SetCheckOut(v string) *HotelOrderCreateShrinkRequest
	GetCheckOut() *string
	SetContractEmail(v string) *HotelOrderCreateShrinkRequest
	GetContractEmail() *string
	SetContractName(v string) *HotelOrderCreateShrinkRequest
	GetContractName() *string
	SetContractPhone(v string) *HotelOrderCreateShrinkRequest
	GetContractPhone() *string
	SetCorpPayPrice(v int64) *HotelOrderCreateShrinkRequest
	GetCorpPayPrice() *int64
	SetDisOrderId(v string) *HotelOrderCreateShrinkRequest
	GetDisOrderId() *string
	SetExtra(v string) *HotelOrderCreateShrinkRequest
	GetExtra() *string
	SetInvoiceInfoShrink(v string) *HotelOrderCreateShrinkRequest
	GetInvoiceInfoShrink() *string
	SetItemId(v int64) *HotelOrderCreateShrinkRequest
	GetItemId() *int64
	SetItineraryNo(v string) *HotelOrderCreateShrinkRequest
	GetItineraryNo() *string
	SetOccupantInfoListShrink(v string) *HotelOrderCreateShrinkRequest
	GetOccupantInfoListShrink() *string
	SetPersonPayPrice(v int64) *HotelOrderCreateShrinkRequest
	GetPersonPayPrice() *int64
	SetPromotionInfoShrink(v string) *HotelOrderCreateShrinkRequest
	GetPromotionInfoShrink() *string
	SetRatePlanId(v int64) *HotelOrderCreateShrinkRequest
	GetRatePlanId() *int64
	SetRoomId(v int64) *HotelOrderCreateShrinkRequest
	GetRoomId() *int64
	SetRoomNum(v int32) *HotelOrderCreateShrinkRequest
	GetRoomNum() *int32
	SetSellerId(v int64) *HotelOrderCreateShrinkRequest
	GetSellerId() *int64
	SetShid(v int64) *HotelOrderCreateShrinkRequest
	GetShid() *int64
	SetTotalOrderPrice(v int64) *HotelOrderCreateShrinkRequest
	GetTotalOrderPrice() *int64
	SetValidateResKey(v string) *HotelOrderCreateShrinkRequest
	GetValidateResKey() *string
}

type HotelOrderCreateShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123122
	BtripUserId *string `json:"btrip_user_id,omitempty" xml:"btrip_user_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2020-10-20
	CheckIn *string `json:"check_in,omitempty" xml:"check_in,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2020-10-20
	CheckOut *string `json:"check_out,omitempty" xml:"check_out,omitempty"`
	// example:
	//
	// demo
	ContractEmail *string `json:"contract_email,omitempty" xml:"contract_email,omitempty"`
	ContractName  *string `json:"contract_name,omitempty" xml:"contract_name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 19281772123
	ContractPhone *string `json:"contract_phone,omitempty" xml:"contract_phone,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 100
	CorpPayPrice *int64 `json:"corp_pay_price,omitempty" xml:"corp_pay_price,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dis123
	DisOrderId        *string `json:"dis_order_id,omitempty" xml:"dis_order_id,omitempty"`
	Extra             *string `json:"extra,omitempty" xml:"extra,omitempty"`
	InvoiceInfoShrink *string `json:"invoice_info,omitempty" xml:"invoice_info,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 671570615157
	ItemId *int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// fb5e1abf33924b6c912bd6d80deec0eb-1
	ItineraryNo *string `json:"itinerary_no,omitempty" xml:"itinerary_no,omitempty"`
	// This parameter is required.
	OccupantInfoListShrink *string `json:"occupant_info_list,omitempty" xml:"occupant_info_list,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1000
	PersonPayPrice      *int64  `json:"person_pay_price,omitempty" xml:"person_pay_price,omitempty"`
	PromotionInfoShrink *string `json:"promotion_info,omitempty" xml:"promotion_info,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1399417428510
	RatePlanId *int64 `json:"rate_plan_id,omitempty" xml:"rate_plan_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 187211
	RoomId *int64 `json:"room_id,omitempty" xml:"room_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	RoomNum *int32 `json:"room_num,omitempty" xml:"room_num,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2088441675613762
	SellerId *int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2198781
	Shid *int64 `json:"shid,omitempty" xml:"shid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 100
	TotalOrderPrice *int64 `json:"total_order_price,omitempty" xml:"total_order_price,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// nonUltron_1673575241156_d91ea8ad16735752359161037bf6cf_c54d3768312a4b249b719f126377bf82
	ValidateResKey *string `json:"validate_res_key,omitempty" xml:"validate_res_key,omitempty"`
}

func (s HotelOrderCreateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderCreateShrinkRequest) GoString() string {
	return s.String()
}

func (s *HotelOrderCreateShrinkRequest) GetBtripUserId() *string {
	return s.BtripUserId
}

func (s *HotelOrderCreateShrinkRequest) GetCheckIn() *string {
	return s.CheckIn
}

func (s *HotelOrderCreateShrinkRequest) GetCheckOut() *string {
	return s.CheckOut
}

func (s *HotelOrderCreateShrinkRequest) GetContractEmail() *string {
	return s.ContractEmail
}

func (s *HotelOrderCreateShrinkRequest) GetContractName() *string {
	return s.ContractName
}

func (s *HotelOrderCreateShrinkRequest) GetContractPhone() *string {
	return s.ContractPhone
}

func (s *HotelOrderCreateShrinkRequest) GetCorpPayPrice() *int64 {
	return s.CorpPayPrice
}

func (s *HotelOrderCreateShrinkRequest) GetDisOrderId() *string {
	return s.DisOrderId
}

func (s *HotelOrderCreateShrinkRequest) GetExtra() *string {
	return s.Extra
}

func (s *HotelOrderCreateShrinkRequest) GetInvoiceInfoShrink() *string {
	return s.InvoiceInfoShrink
}

func (s *HotelOrderCreateShrinkRequest) GetItemId() *int64 {
	return s.ItemId
}

func (s *HotelOrderCreateShrinkRequest) GetItineraryNo() *string {
	return s.ItineraryNo
}

func (s *HotelOrderCreateShrinkRequest) GetOccupantInfoListShrink() *string {
	return s.OccupantInfoListShrink
}

func (s *HotelOrderCreateShrinkRequest) GetPersonPayPrice() *int64 {
	return s.PersonPayPrice
}

func (s *HotelOrderCreateShrinkRequest) GetPromotionInfoShrink() *string {
	return s.PromotionInfoShrink
}

func (s *HotelOrderCreateShrinkRequest) GetRatePlanId() *int64 {
	return s.RatePlanId
}

func (s *HotelOrderCreateShrinkRequest) GetRoomId() *int64 {
	return s.RoomId
}

func (s *HotelOrderCreateShrinkRequest) GetRoomNum() *int32 {
	return s.RoomNum
}

func (s *HotelOrderCreateShrinkRequest) GetSellerId() *int64 {
	return s.SellerId
}

func (s *HotelOrderCreateShrinkRequest) GetShid() *int64 {
	return s.Shid
}

func (s *HotelOrderCreateShrinkRequest) GetTotalOrderPrice() *int64 {
	return s.TotalOrderPrice
}

func (s *HotelOrderCreateShrinkRequest) GetValidateResKey() *string {
	return s.ValidateResKey
}

func (s *HotelOrderCreateShrinkRequest) SetBtripUserId(v string) *HotelOrderCreateShrinkRequest {
	s.BtripUserId = &v
	return s
}

func (s *HotelOrderCreateShrinkRequest) SetCheckIn(v string) *HotelOrderCreateShrinkRequest {
	s.CheckIn = &v
	return s
}

func (s *HotelOrderCreateShrinkRequest) SetCheckOut(v string) *HotelOrderCreateShrinkRequest {
	s.CheckOut = &v
	return s
}

func (s *HotelOrderCreateShrinkRequest) SetContractEmail(v string) *HotelOrderCreateShrinkRequest {
	s.ContractEmail = &v
	return s
}

func (s *HotelOrderCreateShrinkRequest) SetContractName(v string) *HotelOrderCreateShrinkRequest {
	s.ContractName = &v
	return s
}

func (s *HotelOrderCreateShrinkRequest) SetContractPhone(v string) *HotelOrderCreateShrinkRequest {
	s.ContractPhone = &v
	return s
}

func (s *HotelOrderCreateShrinkRequest) SetCorpPayPrice(v int64) *HotelOrderCreateShrinkRequest {
	s.CorpPayPrice = &v
	return s
}

func (s *HotelOrderCreateShrinkRequest) SetDisOrderId(v string) *HotelOrderCreateShrinkRequest {
	s.DisOrderId = &v
	return s
}

func (s *HotelOrderCreateShrinkRequest) SetExtra(v string) *HotelOrderCreateShrinkRequest {
	s.Extra = &v
	return s
}

func (s *HotelOrderCreateShrinkRequest) SetInvoiceInfoShrink(v string) *HotelOrderCreateShrinkRequest {
	s.InvoiceInfoShrink = &v
	return s
}

func (s *HotelOrderCreateShrinkRequest) SetItemId(v int64) *HotelOrderCreateShrinkRequest {
	s.ItemId = &v
	return s
}

func (s *HotelOrderCreateShrinkRequest) SetItineraryNo(v string) *HotelOrderCreateShrinkRequest {
	s.ItineraryNo = &v
	return s
}

func (s *HotelOrderCreateShrinkRequest) SetOccupantInfoListShrink(v string) *HotelOrderCreateShrinkRequest {
	s.OccupantInfoListShrink = &v
	return s
}

func (s *HotelOrderCreateShrinkRequest) SetPersonPayPrice(v int64) *HotelOrderCreateShrinkRequest {
	s.PersonPayPrice = &v
	return s
}

func (s *HotelOrderCreateShrinkRequest) SetPromotionInfoShrink(v string) *HotelOrderCreateShrinkRequest {
	s.PromotionInfoShrink = &v
	return s
}

func (s *HotelOrderCreateShrinkRequest) SetRatePlanId(v int64) *HotelOrderCreateShrinkRequest {
	s.RatePlanId = &v
	return s
}

func (s *HotelOrderCreateShrinkRequest) SetRoomId(v int64) *HotelOrderCreateShrinkRequest {
	s.RoomId = &v
	return s
}

func (s *HotelOrderCreateShrinkRequest) SetRoomNum(v int32) *HotelOrderCreateShrinkRequest {
	s.RoomNum = &v
	return s
}

func (s *HotelOrderCreateShrinkRequest) SetSellerId(v int64) *HotelOrderCreateShrinkRequest {
	s.SellerId = &v
	return s
}

func (s *HotelOrderCreateShrinkRequest) SetShid(v int64) *HotelOrderCreateShrinkRequest {
	s.Shid = &v
	return s
}

func (s *HotelOrderCreateShrinkRequest) SetTotalOrderPrice(v int64) *HotelOrderCreateShrinkRequest {
	s.TotalOrderPrice = &v
	return s
}

func (s *HotelOrderCreateShrinkRequest) SetValidateResKey(v string) *HotelOrderCreateShrinkRequest {
	s.ValidateResKey = &v
	return s
}

func (s *HotelOrderCreateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
