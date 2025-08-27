// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelPricePullResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *HotelPricePullResponseBody
	GetCode() *string
	SetMessage(v string) *HotelPricePullResponseBody
	GetMessage() *string
	SetModule(v *HotelPricePullResponseBodyModule) *HotelPricePullResponseBody
	GetModule() *HotelPricePullResponseBodyModule
	SetRequestId(v string) *HotelPricePullResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *HotelPricePullResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *HotelPricePullResponseBody
	GetTraceId() *string
}

type HotelPricePullResponseBody struct {
	// example:
	//
	// 0
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// None
	Message *string                           `json:"message,omitempty" xml:"message,omitempty"`
	Module  *HotelPricePullResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// example:
	//
	// 9BCDD5DE-E6CB-5C25-93B9-9BE178A0AA56
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// traceId
	//
	// example:
	//
	// 210f079e16603757182131635d866a
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s HotelPricePullResponseBody) String() string {
	return dara.Prettify(s)
}

func (s HotelPricePullResponseBody) GoString() string {
	return s.String()
}

func (s *HotelPricePullResponseBody) GetCode() *string {
	return s.Code
}

func (s *HotelPricePullResponseBody) GetMessage() *string {
	return s.Message
}

func (s *HotelPricePullResponseBody) GetModule() *HotelPricePullResponseBodyModule {
	return s.Module
}

func (s *HotelPricePullResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *HotelPricePullResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *HotelPricePullResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *HotelPricePullResponseBody) SetCode(v string) *HotelPricePullResponseBody {
	s.Code = &v
	return s
}

func (s *HotelPricePullResponseBody) SetMessage(v string) *HotelPricePullResponseBody {
	s.Message = &v
	return s
}

func (s *HotelPricePullResponseBody) SetModule(v *HotelPricePullResponseBodyModule) *HotelPricePullResponseBody {
	s.Module = v
	return s
}

func (s *HotelPricePullResponseBody) SetRequestId(v string) *HotelPricePullResponseBody {
	s.RequestId = &v
	return s
}

func (s *HotelPricePullResponseBody) SetSuccess(v bool) *HotelPricePullResponseBody {
	s.Success = &v
	return s
}

func (s *HotelPricePullResponseBody) SetTraceId(v string) *HotelPricePullResponseBody {
	s.TraceId = &v
	return s
}

func (s *HotelPricePullResponseBody) Validate() error {
	return dara.Validate(s)
}

type HotelPricePullResponseBodyModule struct {
	HotelPriceInfos []*HotelPricePullResponseBodyModuleHotelPriceInfos `json:"hotel_price_infos,omitempty" xml:"hotel_price_infos,omitempty" type:"Repeated"`
}

func (s HotelPricePullResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s HotelPricePullResponseBodyModule) GoString() string {
	return s.String()
}

func (s *HotelPricePullResponseBodyModule) GetHotelPriceInfos() []*HotelPricePullResponseBodyModuleHotelPriceInfos {
	return s.HotelPriceInfos
}

func (s *HotelPricePullResponseBodyModule) SetHotelPriceInfos(v []*HotelPricePullResponseBodyModuleHotelPriceInfos) *HotelPricePullResponseBodyModule {
	s.HotelPriceInfos = v
	return s
}

func (s *HotelPricePullResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type HotelPricePullResponseBodyModuleHotelPriceInfos struct {
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// example:
	//
	// 64389015
	HotelId   *string                                                 `json:"hotel_id,omitempty" xml:"hotel_id,omitempty"`
	HotelName *string                                                 `json:"hotel_name,omitempty" xml:"hotel_name,omitempty"`
	Rooms     []*HotelPricePullResponseBodyModuleHotelPriceInfosRooms `json:"rooms,omitempty" xml:"rooms,omitempty" type:"Repeated"`
	// example:
	//
	// 9BCDD5DE-E6CB-5C25-93B9-9BE178A0AA56
	SearchId *string `json:"search_id,omitempty" xml:"search_id,omitempty"`
}

func (s HotelPricePullResponseBodyModuleHotelPriceInfos) String() string {
	return dara.Prettify(s)
}

func (s HotelPricePullResponseBodyModuleHotelPriceInfos) GoString() string {
	return s.String()
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfos) GetAddress() *string {
	return s.Address
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfos) GetHotelId() *string {
	return s.HotelId
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfos) GetHotelName() *string {
	return s.HotelName
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfos) GetRooms() []*HotelPricePullResponseBodyModuleHotelPriceInfosRooms {
	return s.Rooms
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfos) GetSearchId() *string {
	return s.SearchId
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfos) SetAddress(v string) *HotelPricePullResponseBodyModuleHotelPriceInfos {
	s.Address = &v
	return s
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfos) SetHotelId(v string) *HotelPricePullResponseBodyModuleHotelPriceInfos {
	s.HotelId = &v
	return s
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfos) SetHotelName(v string) *HotelPricePullResponseBodyModuleHotelPriceInfos {
	s.HotelName = &v
	return s
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfos) SetRooms(v []*HotelPricePullResponseBodyModuleHotelPriceInfosRooms) *HotelPricePullResponseBodyModuleHotelPriceInfos {
	s.Rooms = v
	return s
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfos) SetSearchId(v string) *HotelPricePullResponseBodyModuleHotelPriceInfos {
	s.SearchId = &v
	return s
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfos) Validate() error {
	return dara.Validate(s)
}

type HotelPricePullResponseBodyModuleHotelPriceInfosRooms struct {
	// example:
	//
	// 32
	Area          *string `json:"area,omitempty" xml:"area,omitempty"`
	Bed           *string `json:"bed,omitempty" xml:"bed,omitempty"`
	BedTypeString *string `json:"bed_type_string,omitempty" xml:"bed_type_string,omitempty"`
	// example:
	//
	// {\\"bathtub\\":true}
	Facility *string `json:"facility,omitempty" xml:"facility,omitempty"`
	// example:
	//
	// 1,2,3,4,5,6
	Floor *string `json:"floor,omitempty" xml:"floor,omitempty"`
	// example:
	//
	// 2
	MaxOccupancy *int32 `json:"max_occupancy,omitempty" xml:"max_occupancy,omitempty"`
	// example:
	//
	// 0
	NetworkService *string                                                      `json:"network_service,omitempty" xml:"network_service,omitempty"`
	Pics           []*string                                                    `json:"pics,omitempty" xml:"pics,omitempty" type:"Repeated"`
	Rates          []*HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates `json:"rates,omitempty" xml:"rates,omitempty" type:"Repeated"`
	// example:
	//
	// 64681618
	RoomId   *string `json:"room_id,omitempty" xml:"room_id,omitempty"`
	RoomName *string `json:"room_name,omitempty" xml:"room_name,omitempty"`
	// example:
	//
	// 0
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 1
	WindowType *string `json:"window_type,omitempty" xml:"window_type,omitempty"`
}

func (s HotelPricePullResponseBodyModuleHotelPriceInfosRooms) String() string {
	return dara.Prettify(s)
}

func (s HotelPricePullResponseBodyModuleHotelPriceInfosRooms) GoString() string {
	return s.String()
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRooms) GetArea() *string {
	return s.Area
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRooms) GetBed() *string {
	return s.Bed
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRooms) GetBedTypeString() *string {
	return s.BedTypeString
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRooms) GetFacility() *string {
	return s.Facility
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRooms) GetFloor() *string {
	return s.Floor
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRooms) GetMaxOccupancy() *int32 {
	return s.MaxOccupancy
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRooms) GetNetworkService() *string {
	return s.NetworkService
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRooms) GetPics() []*string {
	return s.Pics
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRooms) GetRates() []*HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates {
	return s.Rates
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRooms) GetRoomId() *string {
	return s.RoomId
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRooms) GetRoomName() *string {
	return s.RoomName
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRooms) GetStatus() *int32 {
	return s.Status
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRooms) GetWindowType() *string {
	return s.WindowType
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRooms) SetArea(v string) *HotelPricePullResponseBodyModuleHotelPriceInfosRooms {
	s.Area = &v
	return s
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRooms) SetBed(v string) *HotelPricePullResponseBodyModuleHotelPriceInfosRooms {
	s.Bed = &v
	return s
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRooms) SetBedTypeString(v string) *HotelPricePullResponseBodyModuleHotelPriceInfosRooms {
	s.BedTypeString = &v
	return s
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRooms) SetFacility(v string) *HotelPricePullResponseBodyModuleHotelPriceInfosRooms {
	s.Facility = &v
	return s
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRooms) SetFloor(v string) *HotelPricePullResponseBodyModuleHotelPriceInfosRooms {
	s.Floor = &v
	return s
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRooms) SetMaxOccupancy(v int32) *HotelPricePullResponseBodyModuleHotelPriceInfosRooms {
	s.MaxOccupancy = &v
	return s
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRooms) SetNetworkService(v string) *HotelPricePullResponseBodyModuleHotelPriceInfosRooms {
	s.NetworkService = &v
	return s
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRooms) SetPics(v []*string) *HotelPricePullResponseBodyModuleHotelPriceInfosRooms {
	s.Pics = v
	return s
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRooms) SetRates(v []*HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates) *HotelPricePullResponseBodyModuleHotelPriceInfosRooms {
	s.Rates = v
	return s
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRooms) SetRoomId(v string) *HotelPricePullResponseBodyModuleHotelPriceInfosRooms {
	s.RoomId = &v
	return s
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRooms) SetRoomName(v string) *HotelPricePullResponseBodyModuleHotelPriceInfosRooms {
	s.RoomName = &v
	return s
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRooms) SetStatus(v int32) *HotelPricePullResponseBodyModuleHotelPriceInfosRooms {
	s.Status = &v
	return s
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRooms) SetWindowType(v string) *HotelPricePullResponseBodyModuleHotelPriceInfosRooms {
	s.WindowType = &v
	return s
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRooms) Validate() error {
	return dara.Validate(s)
}

type HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates struct {
	// example:
	//
	// demo
	Breakfast *string `json:"breakfast,omitempty" xml:"breakfast,omitempty"`
	// example:
	//
	// 1
	BreakfastCount         *int32                                                                           `json:"breakfast_count,omitempty" xml:"breakfast_count,omitempty"`
	BtripHotelCancelPolicy *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRatesBtripHotelCancelPolicy `json:"btrip_hotel_cancel_policy,omitempty" xml:"btrip_hotel_cancel_policy,omitempty" type:"Struct"`
	CancelPolicyDesc       *string                                                                          `json:"cancel_policy_desc,omitempty" xml:"cancel_policy_desc,omitempty"`
	// example:
	//
	// 1
	CompanyAassist *string `json:"company_aassist,omitempty" xml:"company_aassist,omitempty"`
	// example:
	//
	// CNY
	CurrencyCode *string `json:"currency_code,omitempty" xml:"currency_code,omitempty"`
	// example:
	//
	// true
	InstantConfirm *bool `json:"instant_confirm,omitempty" xml:"instant_confirm,omitempty"`
	// example:
	//
	// 721700504622
	ItemId *string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// example:
	//
	// 4
	MaxAdvHours *int32 `json:"max_adv_hours,omitempty" xml:"max_adv_hours,omitempty"`
	// example:
	//
	// 4344
	MaxDays *int32 `json:"max_days,omitempty" xml:"max_days,omitempty"`
	// example:
	//
	// 0
	MinAdvHours *int32 `json:"min_adv_hours,omitempty" xml:"min_adv_hours,omitempty"`
	// example:
	//
	// 0
	MinDays *int32 `json:"min_days,omitempty" xml:"min_days,omitempty"`
	// example:
	//
	// 0
	Nod *int32 `json:"nod,omitempty" xml:"nod,omitempty"`
	// example:
	//
	// 1
	Nop *int32 `json:"nop,omitempty" xml:"nop,omitempty"`
	// example:
	//
	// 1
	PaymentType *int32 `json:"payment_type,omitempty" xml:"payment_type,omitempty"`
	// example:
	//
	// 30000
	Price *int64 `json:"price,omitempty" xml:"price,omitempty"`
	// example:
	//
	// demo
	PromotionInfo *string `json:"promotion_info,omitempty" xml:"promotion_info,omitempty"`
	// example:
	//
	// 4
	Quota      *int32                                                                 `json:"quota,omitempty" xml:"quota,omitempty"`
	RateDailys []*HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRatesRateDailys `json:"rate_dailys,omitempty" xml:"rate_dailys,omitempty" type:"Repeated"`
	// example:
	//
	// 4509447432148
	RateId       *string `json:"rate_id,omitempty" xml:"rate_id,omitempty"`
	RatePlanName *string `json:"rate_plan_name,omitempty" xml:"rate_plan_name,omitempty"`
	// example:
	//
	// 4509447432148
	RpId *string `json:"rp_id,omitempty" xml:"rp_id,omitempty"`
	// example:
	//
	// 2829486701
	SellerId *string `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// example:
	//
	// true
	SupportSpecialInvoice *bool `json:"support_special_invoice,omitempty" xml:"support_special_invoice,omitempty"`
}

func (s HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates) String() string {
	return dara.Prettify(s)
}

func (s HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates) GoString() string {
	return s.String()
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates) GetBreakfast() *string {
	return s.Breakfast
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates) GetBreakfastCount() *int32 {
	return s.BreakfastCount
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates) GetBtripHotelCancelPolicy() *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRatesBtripHotelCancelPolicy {
	return s.BtripHotelCancelPolicy
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates) GetCancelPolicyDesc() *string {
	return s.CancelPolicyDesc
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates) GetCompanyAassist() *string {
	return s.CompanyAassist
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates) GetCurrencyCode() *string {
	return s.CurrencyCode
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates) GetInstantConfirm() *bool {
	return s.InstantConfirm
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates) GetItemId() *string {
	return s.ItemId
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates) GetMaxAdvHours() *int32 {
	return s.MaxAdvHours
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates) GetMaxDays() *int32 {
	return s.MaxDays
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates) GetMinAdvHours() *int32 {
	return s.MinAdvHours
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates) GetMinDays() *int32 {
	return s.MinDays
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates) GetNod() *int32 {
	return s.Nod
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates) GetNop() *int32 {
	return s.Nop
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates) GetPaymentType() *int32 {
	return s.PaymentType
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates) GetPrice() *int64 {
	return s.Price
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates) GetPromotionInfo() *string {
	return s.PromotionInfo
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates) GetQuota() *int32 {
	return s.Quota
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates) GetRateDailys() []*HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRatesRateDailys {
	return s.RateDailys
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates) GetRateId() *string {
	return s.RateId
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates) GetRatePlanName() *string {
	return s.RatePlanName
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates) GetRpId() *string {
	return s.RpId
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates) GetSellerId() *string {
	return s.SellerId
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates) GetSupportSpecialInvoice() *bool {
	return s.SupportSpecialInvoice
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates) SetBreakfast(v string) *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates {
	s.Breakfast = &v
	return s
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates) SetBreakfastCount(v int32) *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates {
	s.BreakfastCount = &v
	return s
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates) SetBtripHotelCancelPolicy(v *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRatesBtripHotelCancelPolicy) *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates {
	s.BtripHotelCancelPolicy = v
	return s
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates) SetCancelPolicyDesc(v string) *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates {
	s.CancelPolicyDesc = &v
	return s
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates) SetCompanyAassist(v string) *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates {
	s.CompanyAassist = &v
	return s
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates) SetCurrencyCode(v string) *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates {
	s.CurrencyCode = &v
	return s
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates) SetInstantConfirm(v bool) *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates {
	s.InstantConfirm = &v
	return s
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates) SetItemId(v string) *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates {
	s.ItemId = &v
	return s
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates) SetMaxAdvHours(v int32) *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates {
	s.MaxAdvHours = &v
	return s
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates) SetMaxDays(v int32) *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates {
	s.MaxDays = &v
	return s
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates) SetMinAdvHours(v int32) *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates {
	s.MinAdvHours = &v
	return s
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates) SetMinDays(v int32) *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates {
	s.MinDays = &v
	return s
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates) SetNod(v int32) *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates {
	s.Nod = &v
	return s
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates) SetNop(v int32) *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates {
	s.Nop = &v
	return s
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates) SetPaymentType(v int32) *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates {
	s.PaymentType = &v
	return s
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates) SetPrice(v int64) *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates {
	s.Price = &v
	return s
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates) SetPromotionInfo(v string) *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates {
	s.PromotionInfo = &v
	return s
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates) SetQuota(v int32) *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates {
	s.Quota = &v
	return s
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates) SetRateDailys(v []*HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRatesRateDailys) *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates {
	s.RateDailys = v
	return s
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates) SetRateId(v string) *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates {
	s.RateId = &v
	return s
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates) SetRatePlanName(v string) *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates {
	s.RatePlanName = &v
	return s
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates) SetRpId(v string) *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates {
	s.RpId = &v
	return s
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates) SetSellerId(v string) *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates {
	s.SellerId = &v
	return s
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates) SetSupportSpecialInvoice(v bool) *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates {
	s.SupportSpecialInvoice = &v
	return s
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRates) Validate() error {
	return dara.Validate(s)
}

type HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRatesBtripHotelCancelPolicy struct {
	BtripHotelCancelPolicyInfoDTOList []*HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRatesBtripHotelCancelPolicyBtripHotelCancelPolicyInfoDTOList `json:"btrip_hotel_cancel_policy_info_d_t_o_list,omitempty" xml:"btrip_hotel_cancel_policy_info_d_t_o_list,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	CancelPolicyType *int32 `json:"cancel_policy_type,omitempty" xml:"cancel_policy_type,omitempty"`
}

func (s HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRatesBtripHotelCancelPolicy) String() string {
	return dara.Prettify(s)
}

func (s HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRatesBtripHotelCancelPolicy) GoString() string {
	return s.String()
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRatesBtripHotelCancelPolicy) GetBtripHotelCancelPolicyInfoDTOList() []*HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRatesBtripHotelCancelPolicyBtripHotelCancelPolicyInfoDTOList {
	return s.BtripHotelCancelPolicyInfoDTOList
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRatesBtripHotelCancelPolicy) GetCancelPolicyType() *int32 {
	return s.CancelPolicyType
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRatesBtripHotelCancelPolicy) SetBtripHotelCancelPolicyInfoDTOList(v []*HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRatesBtripHotelCancelPolicyBtripHotelCancelPolicyInfoDTOList) *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRatesBtripHotelCancelPolicy {
	s.BtripHotelCancelPolicyInfoDTOList = v
	return s
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRatesBtripHotelCancelPolicy) SetCancelPolicyType(v int32) *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRatesBtripHotelCancelPolicy {
	s.CancelPolicyType = &v
	return s
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRatesBtripHotelCancelPolicy) Validate() error {
	return dara.Validate(s)
}

type HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRatesBtripHotelCancelPolicyBtripHotelCancelPolicyInfoDTOList struct {
	// example:
	//
	// 1
	Hour *int64 `json:"hour,omitempty" xml:"hour,omitempty"`
	// example:
	//
	// 20
	Value *int64 `json:"value,omitempty" xml:"value,omitempty"`
}

func (s HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRatesBtripHotelCancelPolicyBtripHotelCancelPolicyInfoDTOList) String() string {
	return dara.Prettify(s)
}

func (s HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRatesBtripHotelCancelPolicyBtripHotelCancelPolicyInfoDTOList) GoString() string {
	return s.String()
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRatesBtripHotelCancelPolicyBtripHotelCancelPolicyInfoDTOList) GetHour() *int64 {
	return s.Hour
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRatesBtripHotelCancelPolicyBtripHotelCancelPolicyInfoDTOList) GetValue() *int64 {
	return s.Value
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRatesBtripHotelCancelPolicyBtripHotelCancelPolicyInfoDTOList) SetHour(v int64) *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRatesBtripHotelCancelPolicyBtripHotelCancelPolicyInfoDTOList {
	s.Hour = &v
	return s
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRatesBtripHotelCancelPolicyBtripHotelCancelPolicyInfoDTOList) SetValue(v int64) *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRatesBtripHotelCancelPolicyBtripHotelCancelPolicyInfoDTOList {
	s.Value = &v
	return s
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRatesBtripHotelCancelPolicyBtripHotelCancelPolicyInfoDTOList) Validate() error {
	return dara.Validate(s)
}

type HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRatesRateDailys struct {
	// example:
	//
	// 30000
	DiscountPrice *int64 `json:"discount_price,omitempty" xml:"discount_price,omitempty"`
	// example:
	//
	// 62800
	LastDiscountsPrice *int64 `json:"last_discounts_price,omitempty" xml:"last_discounts_price,omitempty"`
	// example:
	//
	// 62800
	Price *int64 `json:"price,omitempty" xml:"price,omitempty"`
	// example:
	//
	// 2023-10-17
	StartDate *string `json:"start_date,omitempty" xml:"start_date,omitempty"`
}

func (s HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRatesRateDailys) String() string {
	return dara.Prettify(s)
}

func (s HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRatesRateDailys) GoString() string {
	return s.String()
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRatesRateDailys) GetDiscountPrice() *int64 {
	return s.DiscountPrice
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRatesRateDailys) GetLastDiscountsPrice() *int64 {
	return s.LastDiscountsPrice
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRatesRateDailys) GetPrice() *int64 {
	return s.Price
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRatesRateDailys) GetStartDate() *string {
	return s.StartDate
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRatesRateDailys) SetDiscountPrice(v int64) *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRatesRateDailys {
	s.DiscountPrice = &v
	return s
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRatesRateDailys) SetLastDiscountsPrice(v int64) *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRatesRateDailys {
	s.LastDiscountsPrice = &v
	return s
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRatesRateDailys) SetPrice(v int64) *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRatesRateDailys {
	s.Price = &v
	return s
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRatesRateDailys) SetStartDate(v string) *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRatesRateDailys {
	s.StartDate = &v
	return s
}

func (s *HotelPricePullResponseBodyModuleHotelPriceInfosRoomsRatesRateDailys) Validate() error {
	return dara.Validate(s)
}
