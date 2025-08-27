// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelGoodsQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *HotelGoodsQueryResponseBody
	GetCode() *string
	SetMessage(v string) *HotelGoodsQueryResponseBody
	GetMessage() *string
	SetModule(v *HotelGoodsQueryResponseBodyModule) *HotelGoodsQueryResponseBody
	GetModule() *HotelGoodsQueryResponseBodyModule
	SetRequestId(v string) *HotelGoodsQueryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *HotelGoodsQueryResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *HotelGoodsQueryResponseBody
	GetTraceId() *string
}

type HotelGoodsQueryResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// None
	Message *string                            `json:"message,omitempty" xml:"message,omitempty"`
	Module  *HotelGoodsQueryResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// example:
	//
	// A5009956-1077-52FB-B520-EA8C7E91D722
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 210e847f16611516748613869de4f6
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s HotelGoodsQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s HotelGoodsQueryResponseBody) GoString() string {
	return s.String()
}

func (s *HotelGoodsQueryResponseBody) GetCode() *string {
	return s.Code
}

func (s *HotelGoodsQueryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *HotelGoodsQueryResponseBody) GetModule() *HotelGoodsQueryResponseBodyModule {
	return s.Module
}

func (s *HotelGoodsQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *HotelGoodsQueryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *HotelGoodsQueryResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *HotelGoodsQueryResponseBody) SetCode(v string) *HotelGoodsQueryResponseBody {
	s.Code = &v
	return s
}

func (s *HotelGoodsQueryResponseBody) SetMessage(v string) *HotelGoodsQueryResponseBody {
	s.Message = &v
	return s
}

func (s *HotelGoodsQueryResponseBody) SetModule(v *HotelGoodsQueryResponseBodyModule) *HotelGoodsQueryResponseBody {
	s.Module = v
	return s
}

func (s *HotelGoodsQueryResponseBody) SetRequestId(v string) *HotelGoodsQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *HotelGoodsQueryResponseBody) SetSuccess(v bool) *HotelGoodsQueryResponseBody {
	s.Success = &v
	return s
}

func (s *HotelGoodsQueryResponseBody) SetTraceId(v string) *HotelGoodsQueryResponseBody {
	s.TraceId = &v
	return s
}

func (s *HotelGoodsQueryResponseBody) Validate() error {
	return dara.Validate(s)
}

type HotelGoodsQueryResponseBodyModule struct {
	Address             *string            `json:"address,omitempty" xml:"address,omitempty"`
	BookingInstructions map[string]*string `json:"booking_instructions,omitempty" xml:"booking_instructions,omitempty"`
	// example:
	//
	// true
	CanForeigner *bool `json:"can_foreigner,omitempty" xml:"can_foreigner,omitempty"`
	// example:
	//
	// 2022-05-15
	CheckIn *string `json:"check_in,omitempty" xml:"check_in,omitempty"`
	// example:
	//
	// 2022-05-15
	CheckOut             *string   `json:"check_out,omitempty" xml:"check_out,omitempty"`
	Descriptions         []*string `json:"descriptions,omitempty" xml:"descriptions,omitempty" type:"Repeated"`
	DinamicBannerPicUrls []*string `json:"dinamic_banner_pic_urls,omitempty" xml:"dinamic_banner_pic_urls,omitempty" type:"Repeated"`
	// example:
	//
	// 17:00
	EarlyArrivalTime *string `json:"early_arrival_time,omitempty" xml:"early_arrival_time,omitempty"`
	// example:
	//
	// 29382
	HotelId   *int64  `json:"hotel_id,omitempty" xml:"hotel_id,omitempty"`
	HotelName *string `json:"hotel_name,omitempty" xml:"hotel_name,omitempty"`
	// example:
	//
	// 12:00
	LateArrivalTime *string                                   `json:"late_arrival_time,omitempty" xml:"late_arrival_time,omitempty"`
	Rooms           []*HotelGoodsQueryResponseBodyModuleRooms `json:"rooms,omitempty" xml:"rooms,omitempty" type:"Repeated"`
	// example:
	//
	// demo
	SearchId *string `json:"search_id,omitempty" xml:"search_id,omitempty"`
}

func (s HotelGoodsQueryResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s HotelGoodsQueryResponseBodyModule) GoString() string {
	return s.String()
}

func (s *HotelGoodsQueryResponseBodyModule) GetAddress() *string {
	return s.Address
}

func (s *HotelGoodsQueryResponseBodyModule) GetBookingInstructions() map[string]*string {
	return s.BookingInstructions
}

func (s *HotelGoodsQueryResponseBodyModule) GetCanForeigner() *bool {
	return s.CanForeigner
}

func (s *HotelGoodsQueryResponseBodyModule) GetCheckIn() *string {
	return s.CheckIn
}

func (s *HotelGoodsQueryResponseBodyModule) GetCheckOut() *string {
	return s.CheckOut
}

func (s *HotelGoodsQueryResponseBodyModule) GetDescriptions() []*string {
	return s.Descriptions
}

func (s *HotelGoodsQueryResponseBodyModule) GetDinamicBannerPicUrls() []*string {
	return s.DinamicBannerPicUrls
}

func (s *HotelGoodsQueryResponseBodyModule) GetEarlyArrivalTime() *string {
	return s.EarlyArrivalTime
}

func (s *HotelGoodsQueryResponseBodyModule) GetHotelId() *int64 {
	return s.HotelId
}

func (s *HotelGoodsQueryResponseBodyModule) GetHotelName() *string {
	return s.HotelName
}

func (s *HotelGoodsQueryResponseBodyModule) GetLateArrivalTime() *string {
	return s.LateArrivalTime
}

func (s *HotelGoodsQueryResponseBodyModule) GetRooms() []*HotelGoodsQueryResponseBodyModuleRooms {
	return s.Rooms
}

func (s *HotelGoodsQueryResponseBodyModule) GetSearchId() *string {
	return s.SearchId
}

func (s *HotelGoodsQueryResponseBodyModule) SetAddress(v string) *HotelGoodsQueryResponseBodyModule {
	s.Address = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModule) SetBookingInstructions(v map[string]*string) *HotelGoodsQueryResponseBodyModule {
	s.BookingInstructions = v
	return s
}

func (s *HotelGoodsQueryResponseBodyModule) SetCanForeigner(v bool) *HotelGoodsQueryResponseBodyModule {
	s.CanForeigner = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModule) SetCheckIn(v string) *HotelGoodsQueryResponseBodyModule {
	s.CheckIn = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModule) SetCheckOut(v string) *HotelGoodsQueryResponseBodyModule {
	s.CheckOut = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModule) SetDescriptions(v []*string) *HotelGoodsQueryResponseBodyModule {
	s.Descriptions = v
	return s
}

func (s *HotelGoodsQueryResponseBodyModule) SetDinamicBannerPicUrls(v []*string) *HotelGoodsQueryResponseBodyModule {
	s.DinamicBannerPicUrls = v
	return s
}

func (s *HotelGoodsQueryResponseBodyModule) SetEarlyArrivalTime(v string) *HotelGoodsQueryResponseBodyModule {
	s.EarlyArrivalTime = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModule) SetHotelId(v int64) *HotelGoodsQueryResponseBodyModule {
	s.HotelId = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModule) SetHotelName(v string) *HotelGoodsQueryResponseBodyModule {
	s.HotelName = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModule) SetLateArrivalTime(v string) *HotelGoodsQueryResponseBodyModule {
	s.LateArrivalTime = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModule) SetRooms(v []*HotelGoodsQueryResponseBodyModuleRooms) *HotelGoodsQueryResponseBodyModule {
	s.Rooms = v
	return s
}

func (s *HotelGoodsQueryResponseBodyModule) SetSearchId(v string) *HotelGoodsQueryResponseBodyModule {
	s.SearchId = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type HotelGoodsQueryResponseBodyModuleRooms struct {
	// example:
	//
	// 27
	Area          *string `json:"area,omitempty" xml:"area,omitempty"`
	BedTypeString *string `json:"bed_type_string,omitempty" xml:"bed_type_string,omitempty"`
	// example:
	//
	// true
	ExtraBed *bool   `json:"extra_bed,omitempty" xml:"extra_bed,omitempty"`
	Facility *string `json:"facility,omitempty" xml:"facility,omitempty"`
	// example:
	//
	// 1,2,3,4,5,6
	Floor *string `json:"floor,omitempty" xml:"floor,omitempty"`
	// example:
	//
	// 1
	MaxOccupancy *int32  `json:"max_occupancy,omitempty" xml:"max_occupancy,omitempty"`
	Name         *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 0
	NetworkService *string                                        `json:"network_service,omitempty" xml:"network_service,omitempty"`
	Pics           *string                                        `json:"pics,omitempty" xml:"pics,omitempty"`
	Rates          []*HotelGoodsQueryResponseBodyModuleRoomsRates `json:"rates,omitempty" xml:"rates,omitempty" type:"Repeated"`
	// example:
	//
	// demo
	RoomDasc     *string                                              `json:"room_dasc,omitempty" xml:"room_dasc,omitempty"`
	RoomFacility []*string                                            `json:"room_facility,omitempty" xml:"room_facility,omitempty" type:"Repeated"`
	RoomService  []*HotelGoodsQueryResponseBodyModuleRoomsRoomService `json:"room_service,omitempty" xml:"room_service,omitempty" type:"Repeated"`
	// example:
	//
	// 100929
	Srid *int64 `json:"srid,omitempty" xml:"srid,omitempty"`
	// example:
	//
	// 1
	Status     *int32  `json:"status,omitempty" xml:"status,omitempty"`
	WindowType *string `json:"window_type,omitempty" xml:"window_type,omitempty"`
}

func (s HotelGoodsQueryResponseBodyModuleRooms) String() string {
	return dara.Prettify(s)
}

func (s HotelGoodsQueryResponseBodyModuleRooms) GoString() string {
	return s.String()
}

func (s *HotelGoodsQueryResponseBodyModuleRooms) GetArea() *string {
	return s.Area
}

func (s *HotelGoodsQueryResponseBodyModuleRooms) GetBedTypeString() *string {
	return s.BedTypeString
}

func (s *HotelGoodsQueryResponseBodyModuleRooms) GetExtraBed() *bool {
	return s.ExtraBed
}

func (s *HotelGoodsQueryResponseBodyModuleRooms) GetFacility() *string {
	return s.Facility
}

func (s *HotelGoodsQueryResponseBodyModuleRooms) GetFloor() *string {
	return s.Floor
}

func (s *HotelGoodsQueryResponseBodyModuleRooms) GetMaxOccupancy() *int32 {
	return s.MaxOccupancy
}

func (s *HotelGoodsQueryResponseBodyModuleRooms) GetName() *string {
	return s.Name
}

func (s *HotelGoodsQueryResponseBodyModuleRooms) GetNetworkService() *string {
	return s.NetworkService
}

func (s *HotelGoodsQueryResponseBodyModuleRooms) GetPics() *string {
	return s.Pics
}

func (s *HotelGoodsQueryResponseBodyModuleRooms) GetRates() []*HotelGoodsQueryResponseBodyModuleRoomsRates {
	return s.Rates
}

func (s *HotelGoodsQueryResponseBodyModuleRooms) GetRoomDasc() *string {
	return s.RoomDasc
}

func (s *HotelGoodsQueryResponseBodyModuleRooms) GetRoomFacility() []*string {
	return s.RoomFacility
}

func (s *HotelGoodsQueryResponseBodyModuleRooms) GetRoomService() []*HotelGoodsQueryResponseBodyModuleRoomsRoomService {
	return s.RoomService
}

func (s *HotelGoodsQueryResponseBodyModuleRooms) GetSrid() *int64 {
	return s.Srid
}

func (s *HotelGoodsQueryResponseBodyModuleRooms) GetStatus() *int32 {
	return s.Status
}

func (s *HotelGoodsQueryResponseBodyModuleRooms) GetWindowType() *string {
	return s.WindowType
}

func (s *HotelGoodsQueryResponseBodyModuleRooms) SetArea(v string) *HotelGoodsQueryResponseBodyModuleRooms {
	s.Area = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRooms) SetBedTypeString(v string) *HotelGoodsQueryResponseBodyModuleRooms {
	s.BedTypeString = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRooms) SetExtraBed(v bool) *HotelGoodsQueryResponseBodyModuleRooms {
	s.ExtraBed = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRooms) SetFacility(v string) *HotelGoodsQueryResponseBodyModuleRooms {
	s.Facility = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRooms) SetFloor(v string) *HotelGoodsQueryResponseBodyModuleRooms {
	s.Floor = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRooms) SetMaxOccupancy(v int32) *HotelGoodsQueryResponseBodyModuleRooms {
	s.MaxOccupancy = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRooms) SetName(v string) *HotelGoodsQueryResponseBodyModuleRooms {
	s.Name = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRooms) SetNetworkService(v string) *HotelGoodsQueryResponseBodyModuleRooms {
	s.NetworkService = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRooms) SetPics(v string) *HotelGoodsQueryResponseBodyModuleRooms {
	s.Pics = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRooms) SetRates(v []*HotelGoodsQueryResponseBodyModuleRoomsRates) *HotelGoodsQueryResponseBodyModuleRooms {
	s.Rates = v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRooms) SetRoomDasc(v string) *HotelGoodsQueryResponseBodyModuleRooms {
	s.RoomDasc = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRooms) SetRoomFacility(v []*string) *HotelGoodsQueryResponseBodyModuleRooms {
	s.RoomFacility = v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRooms) SetRoomService(v []*HotelGoodsQueryResponseBodyModuleRoomsRoomService) *HotelGoodsQueryResponseBodyModuleRooms {
	s.RoomService = v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRooms) SetSrid(v int64) *HotelGoodsQueryResponseBodyModuleRooms {
	s.Srid = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRooms) SetStatus(v int32) *HotelGoodsQueryResponseBodyModuleRooms {
	s.Status = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRooms) SetWindowType(v string) *HotelGoodsQueryResponseBodyModuleRooms {
	s.WindowType = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRooms) Validate() error {
	return dara.Validate(s)
}

type HotelGoodsQueryResponseBodyModuleRoomsRates struct {
	BedDesc              *string                                                            `json:"bed_desc,omitempty" xml:"bed_desc,omitempty"`
	BedType              *string                                                            `json:"bed_type,omitempty" xml:"bed_type,omitempty"`
	Breakfast            *string                                                            `json:"breakfast,omitempty" xml:"breakfast,omitempty"`
	BtripCancelRule      *HotelGoodsQueryResponseBodyModuleRoomsRatesBtripCancelRule        `json:"btrip_cancel_rule,omitempty" xml:"btrip_cancel_rule,omitempty" type:"Struct"`
	BtripHotelCancelDesc []*HotelGoodsQueryResponseBodyModuleRoomsRatesBtripHotelCancelDesc `json:"btrip_hotel_cancel_desc,omitempty" xml:"btrip_hotel_cancel_desc,omitempty" type:"Repeated"`
	// example:
	//
	// true
	CanSmoking       *bool   `json:"can_smoking,omitempty" xml:"can_smoking,omitempty"`
	CancelPolicyDesc *string `json:"cancel_policy_desc,omitempty" xml:"cancel_policy_desc,omitempty"`
	// example:
	//
	// 1
	CancelPolicyType *int32 `json:"cancel_policy_type,omitempty" xml:"cancel_policy_type,omitempty"`
	// example:
	//
	// 0
	CompanyAassist *string `json:"company_aassist,omitempty" xml:"company_aassist,omitempty"`
	// example:
	//
	// 0
	ConfirmType *int32 `json:"confirm_type,omitempty" xml:"confirm_type,omitempty"`
	// example:
	//
	// cny
	CurrencyCode *string `json:"currency_code,omitempty" xml:"currency_code,omitempty"`
	// example:
	//
	// demo
	DailyPriceFormatYuan *string `json:"daily_price_format_yuan,omitempty" xml:"daily_price_format_yuan,omitempty"`
	// example:
	//
	// 400
	DailyPriceView *string                                                  `json:"daily_price_view,omitempty" xml:"daily_price_view,omitempty"`
	DiscountDesc   *HotelGoodsQueryResponseBodyModuleRoomsRatesDiscountDesc `json:"discount_desc,omitempty" xml:"discount_desc,omitempty" type:"Struct"`
	// example:
	//
	// demo
	EndTimeDaily            *string                                                               `json:"end_time_daily,omitempty" xml:"end_time_daily,omitempty"`
	HotelDetailRatePriceDTO []*HotelGoodsQueryResponseBodyModuleRoomsRatesHotelDetailRatePriceDTO `json:"hotel_detail_rate_price_d_t_o,omitempty" xml:"hotel_detail_rate_price_d_t_o,omitempty" type:"Repeated"`
	// example:
	//
	// true
	InstantConfirm *bool   `json:"instant_confirm,omitempty" xml:"instant_confirm,omitempty"`
	InventoryDesc  *string `json:"inventory_desc,omitempty" xml:"inventory_desc,omitempty"`
	// example:
	//
	// 100
	InventoryPrice *string `json:"inventory_price,omitempty" xml:"inventory_price,omitempty"`
	// example:
	//
	// true
	IsBusinessPay4Goods *bool  `json:"is_business_pay4_goods,omitempty" xml:"is_business_pay4_goods,omitempty"`
	IsGuarantee         *int32 `json:"is_guarantee,omitempty" xml:"is_guarantee,omitempty"`
	// example:
	//
	// true
	IsNeedEmail *bool `json:"is_need_email,omitempty" xml:"is_need_email,omitempty"`
	// example:
	//
	// 612673015638
	ItemId *int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// example:
	//
	// demo
	LastCancelTime *string `json:"last_cancel_time,omitempty" xml:"last_cancel_time,omitempty"`
	// example:
	//
	// 2
	MaxOccupancy *int32 `json:"max_occupancy,omitempty" xml:"max_occupancy,omitempty"`
	// example:
	//
	// 2
	MinAdvHours *int32 `json:"min_adv_hours,omitempty" xml:"min_adv_hours,omitempty"`
	// example:
	//
	// 3
	MinDays *int32 `json:"min_days,omitempty" xml:"min_days,omitempty"`
	// example:
	//
	// 1
	Nod *int32 `json:"nod,omitempty" xml:"nod,omitempty"`
	// example:
	//
	// 2
	Nop *int32 `json:"nop,omitempty" xml:"nop,omitempty"`
	// example:
	//
	// 0
	OrderShipTime *int32 `json:"order_ship_time,omitempty" xml:"order_ship_time,omitempty"`
	// example:
	//
	// 1
	PaymentType *int32 `json:"payment_type,omitempty" xml:"payment_type,omitempty"`
	// example:
	//
	// 0
	PriceType *int32 `json:"price_type,omitempty" xml:"price_type,omitempty"`
	// example:
	//
	// demo
	PromotionInfo *string `json:"promotion_info,omitempty" xml:"promotion_info,omitempty"`
	// example:
	//
	// 4509447432148
	RateId       *int64  `json:"rate_id,omitempty" xml:"rate_id,omitempty"`
	RatePlanName *string `json:"rate_plan_name,omitempty" xml:"rate_plan_name,omitempty"`
	// example:
	//
	// 4509447432148
	RpId *int64 `json:"rp_id,omitempty" xml:"rp_id,omitempty"`
	// example:
	//
	// 4011822148
	SellerId *int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// example:
	//
	// demo
	StartTimeDaily *string `json:"start_time_daily,omitempty" xml:"start_time_daily,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 2321
	SupplierCode *string `json:"supplier_code,omitempty" xml:"supplier_code,omitempty"`
	// example:
	//
	// demo
	SupplierName *string `json:"supplier_name,omitempty" xml:"supplier_name,omitempty"`
	// example:
	//
	// true
	SupportSpecialInvoice *bool `json:"support_special_invoice,omitempty" xml:"support_special_invoice,omitempty"`
	// example:
	//
	// 40000
	UnroundingDailyPriceFormatYuan *string `json:"unrounding_daily_price_format_yuan,omitempty" xml:"unrounding_daily_price_format_yuan,omitempty"`
}

func (s HotelGoodsQueryResponseBodyModuleRoomsRates) String() string {
	return dara.Prettify(s)
}

func (s HotelGoodsQueryResponseBodyModuleRoomsRates) GoString() string {
	return s.String()
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) GetBedDesc() *string {
	return s.BedDesc
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) GetBedType() *string {
	return s.BedType
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) GetBreakfast() *string {
	return s.Breakfast
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) GetBtripCancelRule() *HotelGoodsQueryResponseBodyModuleRoomsRatesBtripCancelRule {
	return s.BtripCancelRule
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) GetBtripHotelCancelDesc() []*HotelGoodsQueryResponseBodyModuleRoomsRatesBtripHotelCancelDesc {
	return s.BtripHotelCancelDesc
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) GetCanSmoking() *bool {
	return s.CanSmoking
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) GetCancelPolicyDesc() *string {
	return s.CancelPolicyDesc
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) GetCancelPolicyType() *int32 {
	return s.CancelPolicyType
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) GetCompanyAassist() *string {
	return s.CompanyAassist
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) GetConfirmType() *int32 {
	return s.ConfirmType
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) GetCurrencyCode() *string {
	return s.CurrencyCode
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) GetDailyPriceFormatYuan() *string {
	return s.DailyPriceFormatYuan
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) GetDailyPriceView() *string {
	return s.DailyPriceView
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) GetDiscountDesc() *HotelGoodsQueryResponseBodyModuleRoomsRatesDiscountDesc {
	return s.DiscountDesc
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) GetEndTimeDaily() *string {
	return s.EndTimeDaily
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) GetHotelDetailRatePriceDTO() []*HotelGoodsQueryResponseBodyModuleRoomsRatesHotelDetailRatePriceDTO {
	return s.HotelDetailRatePriceDTO
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) GetInstantConfirm() *bool {
	return s.InstantConfirm
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) GetInventoryDesc() *string {
	return s.InventoryDesc
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) GetInventoryPrice() *string {
	return s.InventoryPrice
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) GetIsBusinessPay4Goods() *bool {
	return s.IsBusinessPay4Goods
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) GetIsGuarantee() *int32 {
	return s.IsGuarantee
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) GetIsNeedEmail() *bool {
	return s.IsNeedEmail
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) GetItemId() *int64 {
	return s.ItemId
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) GetLastCancelTime() *string {
	return s.LastCancelTime
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) GetMaxOccupancy() *int32 {
	return s.MaxOccupancy
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) GetMinAdvHours() *int32 {
	return s.MinAdvHours
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) GetMinDays() *int32 {
	return s.MinDays
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) GetNod() *int32 {
	return s.Nod
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) GetNop() *int32 {
	return s.Nop
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) GetOrderShipTime() *int32 {
	return s.OrderShipTime
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) GetPaymentType() *int32 {
	return s.PaymentType
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) GetPriceType() *int32 {
	return s.PriceType
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) GetPromotionInfo() *string {
	return s.PromotionInfo
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) GetRateId() *int64 {
	return s.RateId
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) GetRatePlanName() *string {
	return s.RatePlanName
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) GetRpId() *int64 {
	return s.RpId
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) GetSellerId() *int64 {
	return s.SellerId
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) GetStartTimeDaily() *string {
	return s.StartTimeDaily
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) GetStatus() *int32 {
	return s.Status
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) GetSupplierCode() *string {
	return s.SupplierCode
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) GetSupplierName() *string {
	return s.SupplierName
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) GetSupportSpecialInvoice() *bool {
	return s.SupportSpecialInvoice
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) GetUnroundingDailyPriceFormatYuan() *string {
	return s.UnroundingDailyPriceFormatYuan
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) SetBedDesc(v string) *HotelGoodsQueryResponseBodyModuleRoomsRates {
	s.BedDesc = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) SetBedType(v string) *HotelGoodsQueryResponseBodyModuleRoomsRates {
	s.BedType = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) SetBreakfast(v string) *HotelGoodsQueryResponseBodyModuleRoomsRates {
	s.Breakfast = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) SetBtripCancelRule(v *HotelGoodsQueryResponseBodyModuleRoomsRatesBtripCancelRule) *HotelGoodsQueryResponseBodyModuleRoomsRates {
	s.BtripCancelRule = v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) SetBtripHotelCancelDesc(v []*HotelGoodsQueryResponseBodyModuleRoomsRatesBtripHotelCancelDesc) *HotelGoodsQueryResponseBodyModuleRoomsRates {
	s.BtripHotelCancelDesc = v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) SetCanSmoking(v bool) *HotelGoodsQueryResponseBodyModuleRoomsRates {
	s.CanSmoking = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) SetCancelPolicyDesc(v string) *HotelGoodsQueryResponseBodyModuleRoomsRates {
	s.CancelPolicyDesc = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) SetCancelPolicyType(v int32) *HotelGoodsQueryResponseBodyModuleRoomsRates {
	s.CancelPolicyType = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) SetCompanyAassist(v string) *HotelGoodsQueryResponseBodyModuleRoomsRates {
	s.CompanyAassist = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) SetConfirmType(v int32) *HotelGoodsQueryResponseBodyModuleRoomsRates {
	s.ConfirmType = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) SetCurrencyCode(v string) *HotelGoodsQueryResponseBodyModuleRoomsRates {
	s.CurrencyCode = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) SetDailyPriceFormatYuan(v string) *HotelGoodsQueryResponseBodyModuleRoomsRates {
	s.DailyPriceFormatYuan = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) SetDailyPriceView(v string) *HotelGoodsQueryResponseBodyModuleRoomsRates {
	s.DailyPriceView = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) SetDiscountDesc(v *HotelGoodsQueryResponseBodyModuleRoomsRatesDiscountDesc) *HotelGoodsQueryResponseBodyModuleRoomsRates {
	s.DiscountDesc = v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) SetEndTimeDaily(v string) *HotelGoodsQueryResponseBodyModuleRoomsRates {
	s.EndTimeDaily = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) SetHotelDetailRatePriceDTO(v []*HotelGoodsQueryResponseBodyModuleRoomsRatesHotelDetailRatePriceDTO) *HotelGoodsQueryResponseBodyModuleRoomsRates {
	s.HotelDetailRatePriceDTO = v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) SetInstantConfirm(v bool) *HotelGoodsQueryResponseBodyModuleRoomsRates {
	s.InstantConfirm = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) SetInventoryDesc(v string) *HotelGoodsQueryResponseBodyModuleRoomsRates {
	s.InventoryDesc = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) SetInventoryPrice(v string) *HotelGoodsQueryResponseBodyModuleRoomsRates {
	s.InventoryPrice = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) SetIsBusinessPay4Goods(v bool) *HotelGoodsQueryResponseBodyModuleRoomsRates {
	s.IsBusinessPay4Goods = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) SetIsGuarantee(v int32) *HotelGoodsQueryResponseBodyModuleRoomsRates {
	s.IsGuarantee = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) SetIsNeedEmail(v bool) *HotelGoodsQueryResponseBodyModuleRoomsRates {
	s.IsNeedEmail = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) SetItemId(v int64) *HotelGoodsQueryResponseBodyModuleRoomsRates {
	s.ItemId = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) SetLastCancelTime(v string) *HotelGoodsQueryResponseBodyModuleRoomsRates {
	s.LastCancelTime = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) SetMaxOccupancy(v int32) *HotelGoodsQueryResponseBodyModuleRoomsRates {
	s.MaxOccupancy = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) SetMinAdvHours(v int32) *HotelGoodsQueryResponseBodyModuleRoomsRates {
	s.MinAdvHours = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) SetMinDays(v int32) *HotelGoodsQueryResponseBodyModuleRoomsRates {
	s.MinDays = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) SetNod(v int32) *HotelGoodsQueryResponseBodyModuleRoomsRates {
	s.Nod = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) SetNop(v int32) *HotelGoodsQueryResponseBodyModuleRoomsRates {
	s.Nop = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) SetOrderShipTime(v int32) *HotelGoodsQueryResponseBodyModuleRoomsRates {
	s.OrderShipTime = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) SetPaymentType(v int32) *HotelGoodsQueryResponseBodyModuleRoomsRates {
	s.PaymentType = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) SetPriceType(v int32) *HotelGoodsQueryResponseBodyModuleRoomsRates {
	s.PriceType = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) SetPromotionInfo(v string) *HotelGoodsQueryResponseBodyModuleRoomsRates {
	s.PromotionInfo = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) SetRateId(v int64) *HotelGoodsQueryResponseBodyModuleRoomsRates {
	s.RateId = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) SetRatePlanName(v string) *HotelGoodsQueryResponseBodyModuleRoomsRates {
	s.RatePlanName = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) SetRpId(v int64) *HotelGoodsQueryResponseBodyModuleRoomsRates {
	s.RpId = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) SetSellerId(v int64) *HotelGoodsQueryResponseBodyModuleRoomsRates {
	s.SellerId = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) SetStartTimeDaily(v string) *HotelGoodsQueryResponseBodyModuleRoomsRates {
	s.StartTimeDaily = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) SetStatus(v int32) *HotelGoodsQueryResponseBodyModuleRoomsRates {
	s.Status = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) SetSupplierCode(v string) *HotelGoodsQueryResponseBodyModuleRoomsRates {
	s.SupplierCode = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) SetSupplierName(v string) *HotelGoodsQueryResponseBodyModuleRoomsRates {
	s.SupplierName = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) SetSupportSpecialInvoice(v bool) *HotelGoodsQueryResponseBodyModuleRoomsRates {
	s.SupportSpecialInvoice = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) SetUnroundingDailyPriceFormatYuan(v string) *HotelGoodsQueryResponseBodyModuleRoomsRates {
	s.UnroundingDailyPriceFormatYuan = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRates) Validate() error {
	return dara.Validate(s)
}

type HotelGoodsQueryResponseBodyModuleRoomsRatesBtripCancelRule struct {
	BtripHotelCancelPolicyDTO *HotelGoodsQueryResponseBodyModuleRoomsRatesBtripCancelRuleBtripHotelCancelPolicyDTO `json:"btrip_hotel_cancel_policy_d_t_o,omitempty" xml:"btrip_hotel_cancel_policy_d_t_o,omitempty" type:"Struct"`
	CancelPolicyTitle         *string                                                                              `json:"cancel_policy_title,omitempty" xml:"cancel_policy_title,omitempty"`
	// example:
	//
	// 2023-02-27
	CheckIn *string `json:"check_in,omitempty" xml:"check_in,omitempty"`
}

func (s HotelGoodsQueryResponseBodyModuleRoomsRatesBtripCancelRule) String() string {
	return dara.Prettify(s)
}

func (s HotelGoodsQueryResponseBodyModuleRoomsRatesBtripCancelRule) GoString() string {
	return s.String()
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRatesBtripCancelRule) GetBtripHotelCancelPolicyDTO() *HotelGoodsQueryResponseBodyModuleRoomsRatesBtripCancelRuleBtripHotelCancelPolicyDTO {
	return s.BtripHotelCancelPolicyDTO
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRatesBtripCancelRule) GetCancelPolicyTitle() *string {
	return s.CancelPolicyTitle
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRatesBtripCancelRule) GetCheckIn() *string {
	return s.CheckIn
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRatesBtripCancelRule) SetBtripHotelCancelPolicyDTO(v *HotelGoodsQueryResponseBodyModuleRoomsRatesBtripCancelRuleBtripHotelCancelPolicyDTO) *HotelGoodsQueryResponseBodyModuleRoomsRatesBtripCancelRule {
	s.BtripHotelCancelPolicyDTO = v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRatesBtripCancelRule) SetCancelPolicyTitle(v string) *HotelGoodsQueryResponseBodyModuleRoomsRatesBtripCancelRule {
	s.CancelPolicyTitle = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRatesBtripCancelRule) SetCheckIn(v string) *HotelGoodsQueryResponseBodyModuleRoomsRatesBtripCancelRule {
	s.CheckIn = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRatesBtripCancelRule) Validate() error {
	return dara.Validate(s)
}

type HotelGoodsQueryResponseBodyModuleRoomsRatesBtripCancelRuleBtripHotelCancelPolicyDTO struct {
	BtripHotelCancelPolicyInfoDTOList []*HotelGoodsQueryResponseBodyModuleRoomsRatesBtripCancelRuleBtripHotelCancelPolicyDTOBtripHotelCancelPolicyInfoDTOList `json:"btrip_hotel_cancel_policy_info_d_t_o_list,omitempty" xml:"btrip_hotel_cancel_policy_info_d_t_o_list,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	CancelPolicyType *int32 `json:"cancel_policy_type,omitempty" xml:"cancel_policy_type,omitempty"`
}

func (s HotelGoodsQueryResponseBodyModuleRoomsRatesBtripCancelRuleBtripHotelCancelPolicyDTO) String() string {
	return dara.Prettify(s)
}

func (s HotelGoodsQueryResponseBodyModuleRoomsRatesBtripCancelRuleBtripHotelCancelPolicyDTO) GoString() string {
	return s.String()
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRatesBtripCancelRuleBtripHotelCancelPolicyDTO) GetBtripHotelCancelPolicyInfoDTOList() []*HotelGoodsQueryResponseBodyModuleRoomsRatesBtripCancelRuleBtripHotelCancelPolicyDTOBtripHotelCancelPolicyInfoDTOList {
	return s.BtripHotelCancelPolicyInfoDTOList
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRatesBtripCancelRuleBtripHotelCancelPolicyDTO) GetCancelPolicyType() *int32 {
	return s.CancelPolicyType
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRatesBtripCancelRuleBtripHotelCancelPolicyDTO) SetBtripHotelCancelPolicyInfoDTOList(v []*HotelGoodsQueryResponseBodyModuleRoomsRatesBtripCancelRuleBtripHotelCancelPolicyDTOBtripHotelCancelPolicyInfoDTOList) *HotelGoodsQueryResponseBodyModuleRoomsRatesBtripCancelRuleBtripHotelCancelPolicyDTO {
	s.BtripHotelCancelPolicyInfoDTOList = v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRatesBtripCancelRuleBtripHotelCancelPolicyDTO) SetCancelPolicyType(v int32) *HotelGoodsQueryResponseBodyModuleRoomsRatesBtripCancelRuleBtripHotelCancelPolicyDTO {
	s.CancelPolicyType = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRatesBtripCancelRuleBtripHotelCancelPolicyDTO) Validate() error {
	return dara.Validate(s)
}

type HotelGoodsQueryResponseBodyModuleRoomsRatesBtripCancelRuleBtripHotelCancelPolicyDTOBtripHotelCancelPolicyInfoDTOList struct {
	// example:
	//
	// 1
	Hour *int64 `json:"hour,omitempty" xml:"hour,omitempty"`
	// example:
	//
	// 20
	Value *int64 `json:"value,omitempty" xml:"value,omitempty"`
}

func (s HotelGoodsQueryResponseBodyModuleRoomsRatesBtripCancelRuleBtripHotelCancelPolicyDTOBtripHotelCancelPolicyInfoDTOList) String() string {
	return dara.Prettify(s)
}

func (s HotelGoodsQueryResponseBodyModuleRoomsRatesBtripCancelRuleBtripHotelCancelPolicyDTOBtripHotelCancelPolicyInfoDTOList) GoString() string {
	return s.String()
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRatesBtripCancelRuleBtripHotelCancelPolicyDTOBtripHotelCancelPolicyInfoDTOList) GetHour() *int64 {
	return s.Hour
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRatesBtripCancelRuleBtripHotelCancelPolicyDTOBtripHotelCancelPolicyInfoDTOList) GetValue() *int64 {
	return s.Value
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRatesBtripCancelRuleBtripHotelCancelPolicyDTOBtripHotelCancelPolicyInfoDTOList) SetHour(v int64) *HotelGoodsQueryResponseBodyModuleRoomsRatesBtripCancelRuleBtripHotelCancelPolicyDTOBtripHotelCancelPolicyInfoDTOList {
	s.Hour = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRatesBtripCancelRuleBtripHotelCancelPolicyDTOBtripHotelCancelPolicyInfoDTOList) SetValue(v int64) *HotelGoodsQueryResponseBodyModuleRoomsRatesBtripCancelRuleBtripHotelCancelPolicyDTOBtripHotelCancelPolicyInfoDTOList {
	s.Value = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRatesBtripCancelRuleBtripHotelCancelPolicyDTOBtripHotelCancelPolicyInfoDTOList) Validate() error {
	return dara.Validate(s)
}

type HotelGoodsQueryResponseBodyModuleRoomsRatesBtripHotelCancelDesc struct {
	Desc  *string `json:"desc,omitempty" xml:"desc,omitempty"`
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s HotelGoodsQueryResponseBodyModuleRoomsRatesBtripHotelCancelDesc) String() string {
	return dara.Prettify(s)
}

func (s HotelGoodsQueryResponseBodyModuleRoomsRatesBtripHotelCancelDesc) GoString() string {
	return s.String()
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRatesBtripHotelCancelDesc) GetDesc() *string {
	return s.Desc
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRatesBtripHotelCancelDesc) GetTitle() *string {
	return s.Title
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRatesBtripHotelCancelDesc) SetDesc(v string) *HotelGoodsQueryResponseBodyModuleRoomsRatesBtripHotelCancelDesc {
	s.Desc = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRatesBtripHotelCancelDesc) SetTitle(v string) *HotelGoodsQueryResponseBodyModuleRoomsRatesBtripHotelCancelDesc {
	s.Title = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRatesBtripHotelCancelDesc) Validate() error {
	return dara.Validate(s)
}

type HotelGoodsQueryResponseBodyModuleRoomsRatesDiscountDesc struct {
	CashReduceTotal *string                                                                  `json:"cash_reduce_total,omitempty" xml:"cash_reduce_total,omitempty"`
	DinamicLabel    *string                                                                  `json:"dinamic_label,omitempty" xml:"dinamic_label,omitempty"`
	DiscountDetail  []*HotelGoodsQueryResponseBodyModuleRoomsRatesDiscountDescDiscountDetail `json:"discount_detail,omitempty" xml:"discount_detail,omitempty" type:"Repeated"`
	SubTitle        *string                                                                  `json:"sub_title,omitempty" xml:"sub_title,omitempty"`
	Title           *string                                                                  `json:"title,omitempty" xml:"title,omitempty"`
}

func (s HotelGoodsQueryResponseBodyModuleRoomsRatesDiscountDesc) String() string {
	return dara.Prettify(s)
}

func (s HotelGoodsQueryResponseBodyModuleRoomsRatesDiscountDesc) GoString() string {
	return s.String()
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRatesDiscountDesc) GetCashReduceTotal() *string {
	return s.CashReduceTotal
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRatesDiscountDesc) GetDinamicLabel() *string {
	return s.DinamicLabel
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRatesDiscountDesc) GetDiscountDetail() []*HotelGoodsQueryResponseBodyModuleRoomsRatesDiscountDescDiscountDetail {
	return s.DiscountDetail
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRatesDiscountDesc) GetSubTitle() *string {
	return s.SubTitle
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRatesDiscountDesc) GetTitle() *string {
	return s.Title
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRatesDiscountDesc) SetCashReduceTotal(v string) *HotelGoodsQueryResponseBodyModuleRoomsRatesDiscountDesc {
	s.CashReduceTotal = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRatesDiscountDesc) SetDinamicLabel(v string) *HotelGoodsQueryResponseBodyModuleRoomsRatesDiscountDesc {
	s.DinamicLabel = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRatesDiscountDesc) SetDiscountDetail(v []*HotelGoodsQueryResponseBodyModuleRoomsRatesDiscountDescDiscountDetail) *HotelGoodsQueryResponseBodyModuleRoomsRatesDiscountDesc {
	s.DiscountDetail = v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRatesDiscountDesc) SetSubTitle(v string) *HotelGoodsQueryResponseBodyModuleRoomsRatesDiscountDesc {
	s.SubTitle = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRatesDiscountDesc) SetTitle(v string) *HotelGoodsQueryResponseBodyModuleRoomsRatesDiscountDesc {
	s.Title = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRatesDiscountDesc) Validate() error {
	return dara.Validate(s)
}

type HotelGoodsQueryResponseBodyModuleRoomsRatesDiscountDescDiscountDetail struct {
	LabelName []*string `json:"label_name,omitempty" xml:"label_name,omitempty" type:"Repeated"`
	MoneyDesc *string   `json:"money_desc,omitempty" xml:"money_desc,omitempty"`
}

func (s HotelGoodsQueryResponseBodyModuleRoomsRatesDiscountDescDiscountDetail) String() string {
	return dara.Prettify(s)
}

func (s HotelGoodsQueryResponseBodyModuleRoomsRatesDiscountDescDiscountDetail) GoString() string {
	return s.String()
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRatesDiscountDescDiscountDetail) GetLabelName() []*string {
	return s.LabelName
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRatesDiscountDescDiscountDetail) GetMoneyDesc() *string {
	return s.MoneyDesc
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRatesDiscountDescDiscountDetail) SetLabelName(v []*string) *HotelGoodsQueryResponseBodyModuleRoomsRatesDiscountDescDiscountDetail {
	s.LabelName = v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRatesDiscountDescDiscountDetail) SetMoneyDesc(v string) *HotelGoodsQueryResponseBodyModuleRoomsRatesDiscountDescDiscountDetail {
	s.MoneyDesc = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRatesDiscountDescDiscountDetail) Validate() error {
	return dara.Validate(s)
}

type HotelGoodsQueryResponseBodyModuleRoomsRatesHotelDetailRatePriceDTO struct {
	// example:
	//
	// 100
	BeforeDiscountPrice *int64 `json:"before_discount_price,omitempty" xml:"before_discount_price,omitempty"`
	// example:
	//
	// demo
	Breakfast *string `json:"breakfast,omitempty" xml:"breakfast,omitempty"`
	// example:
	//
	// 1000
	DiscountPrice *int64 `json:"discount_price,omitempty" xml:"discount_price,omitempty"`
	// example:
	//
	// 100
	LastDiscountsPrice         *int64 `json:"last_discounts_price,omitempty" xml:"last_discounts_price,omitempty"`
	LastDiscountsRoundingPrice *int64 `json:"last_discounts_rounding_price,omitempty" xml:"last_discounts_rounding_price,omitempty"`
	// example:
	//
	// 1
	LastNum *int32 `json:"last_num,omitempty" xml:"last_num,omitempty"`
	// example:
	//
	// 2023-03-25
	RateStartTime *string `json:"rate_start_time,omitempty" xml:"rate_start_time,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s HotelGoodsQueryResponseBodyModuleRoomsRatesHotelDetailRatePriceDTO) String() string {
	return dara.Prettify(s)
}

func (s HotelGoodsQueryResponseBodyModuleRoomsRatesHotelDetailRatePriceDTO) GoString() string {
	return s.String()
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRatesHotelDetailRatePriceDTO) GetBeforeDiscountPrice() *int64 {
	return s.BeforeDiscountPrice
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRatesHotelDetailRatePriceDTO) GetBreakfast() *string {
	return s.Breakfast
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRatesHotelDetailRatePriceDTO) GetDiscountPrice() *int64 {
	return s.DiscountPrice
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRatesHotelDetailRatePriceDTO) GetLastDiscountsPrice() *int64 {
	return s.LastDiscountsPrice
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRatesHotelDetailRatePriceDTO) GetLastDiscountsRoundingPrice() *int64 {
	return s.LastDiscountsRoundingPrice
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRatesHotelDetailRatePriceDTO) GetLastNum() *int32 {
	return s.LastNum
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRatesHotelDetailRatePriceDTO) GetRateStartTime() *string {
	return s.RateStartTime
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRatesHotelDetailRatePriceDTO) GetStatus() *int32 {
	return s.Status
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRatesHotelDetailRatePriceDTO) SetBeforeDiscountPrice(v int64) *HotelGoodsQueryResponseBodyModuleRoomsRatesHotelDetailRatePriceDTO {
	s.BeforeDiscountPrice = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRatesHotelDetailRatePriceDTO) SetBreakfast(v string) *HotelGoodsQueryResponseBodyModuleRoomsRatesHotelDetailRatePriceDTO {
	s.Breakfast = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRatesHotelDetailRatePriceDTO) SetDiscountPrice(v int64) *HotelGoodsQueryResponseBodyModuleRoomsRatesHotelDetailRatePriceDTO {
	s.DiscountPrice = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRatesHotelDetailRatePriceDTO) SetLastDiscountsPrice(v int64) *HotelGoodsQueryResponseBodyModuleRoomsRatesHotelDetailRatePriceDTO {
	s.LastDiscountsPrice = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRatesHotelDetailRatePriceDTO) SetLastDiscountsRoundingPrice(v int64) *HotelGoodsQueryResponseBodyModuleRoomsRatesHotelDetailRatePriceDTO {
	s.LastDiscountsRoundingPrice = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRatesHotelDetailRatePriceDTO) SetLastNum(v int32) *HotelGoodsQueryResponseBodyModuleRoomsRatesHotelDetailRatePriceDTO {
	s.LastNum = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRatesHotelDetailRatePriceDTO) SetRateStartTime(v string) *HotelGoodsQueryResponseBodyModuleRoomsRatesHotelDetailRatePriceDTO {
	s.RateStartTime = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRatesHotelDetailRatePriceDTO) SetStatus(v int32) *HotelGoodsQueryResponseBodyModuleRoomsRatesHotelDetailRatePriceDTO {
	s.Status = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRatesHotelDetailRatePriceDTO) Validate() error {
	return dara.Validate(s)
}

type HotelGoodsQueryResponseBodyModuleRoomsRoomService struct {
	// example:
	//
	// #000000
	Color *string `json:"color,omitempty" xml:"color,omitempty"`
	Desc  *string `json:"desc,omitempty" xml:"desc,omitempty"`
	// example:
	//
	// #4AA900
	HighlightColorColor *string `json:"highlight_color_color,omitempty" xml:"highlight_color_color,omitempty"`
	Title               *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s HotelGoodsQueryResponseBodyModuleRoomsRoomService) String() string {
	return dara.Prettify(s)
}

func (s HotelGoodsQueryResponseBodyModuleRoomsRoomService) GoString() string {
	return s.String()
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRoomService) GetColor() *string {
	return s.Color
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRoomService) GetDesc() *string {
	return s.Desc
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRoomService) GetHighlightColorColor() *string {
	return s.HighlightColorColor
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRoomService) GetTitle() *string {
	return s.Title
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRoomService) SetColor(v string) *HotelGoodsQueryResponseBodyModuleRoomsRoomService {
	s.Color = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRoomService) SetDesc(v string) *HotelGoodsQueryResponseBodyModuleRoomsRoomService {
	s.Desc = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRoomService) SetHighlightColorColor(v string) *HotelGoodsQueryResponseBodyModuleRoomsRoomService {
	s.HighlightColorColor = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRoomService) SetTitle(v string) *HotelGoodsQueryResponseBodyModuleRoomsRoomService {
	s.Title = &v
	return s
}

func (s *HotelGoodsQueryResponseBodyModuleRoomsRoomService) Validate() error {
	return dara.Validate(s)
}
