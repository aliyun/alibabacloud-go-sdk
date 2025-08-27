// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelOrderPreValidateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *HotelOrderPreValidateResponseBody
	GetCode() *string
	SetMessage(v string) *HotelOrderPreValidateResponseBody
	GetMessage() *string
	SetModule(v *HotelOrderPreValidateResponseBodyModule) *HotelOrderPreValidateResponseBody
	GetModule() *HotelOrderPreValidateResponseBodyModule
	SetRequestId(v string) *HotelOrderPreValidateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *HotelOrderPreValidateResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *HotelOrderPreValidateResponseBody
	GetTraceId() *string
}

type HotelOrderPreValidateResponseBody struct {
	// example:
	//
	// 0
	Code    *string                                  `json:"code,omitempty" xml:"code,omitempty"`
	Message *string                                  `json:"message,omitempty" xml:"message,omitempty"`
	Module  *HotelOrderPreValidateResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// requestId
	//
	// example:
	//
	// B72B39C8-32DE-558D-AD1C-D53F11F6ADFE
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 21041ce316577904808056433edbb2
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s HotelOrderPreValidateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderPreValidateResponseBody) GoString() string {
	return s.String()
}

func (s *HotelOrderPreValidateResponseBody) GetCode() *string {
	return s.Code
}

func (s *HotelOrderPreValidateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *HotelOrderPreValidateResponseBody) GetModule() *HotelOrderPreValidateResponseBodyModule {
	return s.Module
}

func (s *HotelOrderPreValidateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *HotelOrderPreValidateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *HotelOrderPreValidateResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *HotelOrderPreValidateResponseBody) SetCode(v string) *HotelOrderPreValidateResponseBody {
	s.Code = &v
	return s
}

func (s *HotelOrderPreValidateResponseBody) SetMessage(v string) *HotelOrderPreValidateResponseBody {
	s.Message = &v
	return s
}

func (s *HotelOrderPreValidateResponseBody) SetModule(v *HotelOrderPreValidateResponseBodyModule) *HotelOrderPreValidateResponseBody {
	s.Module = v
	return s
}

func (s *HotelOrderPreValidateResponseBody) SetRequestId(v string) *HotelOrderPreValidateResponseBody {
	s.RequestId = &v
	return s
}

func (s *HotelOrderPreValidateResponseBody) SetSuccess(v bool) *HotelOrderPreValidateResponseBody {
	s.Success = &v
	return s
}

func (s *HotelOrderPreValidateResponseBody) SetTraceId(v string) *HotelOrderPreValidateResponseBody {
	s.TraceId = &v
	return s
}

func (s *HotelOrderPreValidateResponseBody) Validate() error {
	return dara.Validate(s)
}

type HotelOrderPreValidateResponseBodyModule struct {
	// example:
	//
	// demo
	ExtendInfo  *string                                             `json:"extend_info,omitempty" xml:"extend_info,omitempty"`
	ItemInvoice *HotelOrderPreValidateResponseBodyModuleItemInvoice `json:"item_invoice,omitempty" xml:"item_invoice,omitempty" type:"Struct"`
	// example:
	//
	// fb5e1abf33924b6c912bd6d80deec0eb-4
	ItineraryNo   *string                                                 `json:"itinerary_no,omitempty" xml:"itinerary_no,omitempty"`
	PromotionInfo *HotelOrderPreValidateResponseBodyModulePromotionInfo   `json:"promotion_info,omitempty" xml:"promotion_info,omitempty" type:"Struct"`
	RatePlanDaily []*HotelOrderPreValidateResponseBodyModuleRatePlanDaily `json:"rate_plan_daily,omitempty" xml:"rate_plan_daily,omitempty" type:"Repeated"`
	// example:
	//
	// 5314280514218
	RatePlanId   *int64                                               `json:"rate_plan_id,omitempty" xml:"rate_plan_id,omitempty"`
	RatePlanInfo *HotelOrderPreValidateResponseBodyModuleRatePlanInfo `json:"rate_plan_info,omitempty" xml:"rate_plan_info,omitempty" type:"Struct"`
	// example:
	//
	// nonUltron_1673575241156_d91ea8ad16735752359161037bf6cf_c54d3768312a4b249b719f126377bf82
	ValidateResKey *string `json:"validate_res_key,omitempty" xml:"validate_res_key,omitempty"`
}

func (s HotelOrderPreValidateResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderPreValidateResponseBodyModule) GoString() string {
	return s.String()
}

func (s *HotelOrderPreValidateResponseBodyModule) GetExtendInfo() *string {
	return s.ExtendInfo
}

func (s *HotelOrderPreValidateResponseBodyModule) GetItemInvoice() *HotelOrderPreValidateResponseBodyModuleItemInvoice {
	return s.ItemInvoice
}

func (s *HotelOrderPreValidateResponseBodyModule) GetItineraryNo() *string {
	return s.ItineraryNo
}

func (s *HotelOrderPreValidateResponseBodyModule) GetPromotionInfo() *HotelOrderPreValidateResponseBodyModulePromotionInfo {
	return s.PromotionInfo
}

func (s *HotelOrderPreValidateResponseBodyModule) GetRatePlanDaily() []*HotelOrderPreValidateResponseBodyModuleRatePlanDaily {
	return s.RatePlanDaily
}

func (s *HotelOrderPreValidateResponseBodyModule) GetRatePlanId() *int64 {
	return s.RatePlanId
}

func (s *HotelOrderPreValidateResponseBodyModule) GetRatePlanInfo() *HotelOrderPreValidateResponseBodyModuleRatePlanInfo {
	return s.RatePlanInfo
}

func (s *HotelOrderPreValidateResponseBodyModule) GetValidateResKey() *string {
	return s.ValidateResKey
}

func (s *HotelOrderPreValidateResponseBodyModule) SetExtendInfo(v string) *HotelOrderPreValidateResponseBodyModule {
	s.ExtendInfo = &v
	return s
}

func (s *HotelOrderPreValidateResponseBodyModule) SetItemInvoice(v *HotelOrderPreValidateResponseBodyModuleItemInvoice) *HotelOrderPreValidateResponseBodyModule {
	s.ItemInvoice = v
	return s
}

func (s *HotelOrderPreValidateResponseBodyModule) SetItineraryNo(v string) *HotelOrderPreValidateResponseBodyModule {
	s.ItineraryNo = &v
	return s
}

func (s *HotelOrderPreValidateResponseBodyModule) SetPromotionInfo(v *HotelOrderPreValidateResponseBodyModulePromotionInfo) *HotelOrderPreValidateResponseBodyModule {
	s.PromotionInfo = v
	return s
}

func (s *HotelOrderPreValidateResponseBodyModule) SetRatePlanDaily(v []*HotelOrderPreValidateResponseBodyModuleRatePlanDaily) *HotelOrderPreValidateResponseBodyModule {
	s.RatePlanDaily = v
	return s
}

func (s *HotelOrderPreValidateResponseBodyModule) SetRatePlanId(v int64) *HotelOrderPreValidateResponseBodyModule {
	s.RatePlanId = &v
	return s
}

func (s *HotelOrderPreValidateResponseBodyModule) SetRatePlanInfo(v *HotelOrderPreValidateResponseBodyModuleRatePlanInfo) *HotelOrderPreValidateResponseBodyModule {
	s.RatePlanInfo = v
	return s
}

func (s *HotelOrderPreValidateResponseBodyModule) SetValidateResKey(v string) *HotelOrderPreValidateResponseBodyModule {
	s.ValidateResKey = &v
	return s
}

func (s *HotelOrderPreValidateResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type HotelOrderPreValidateResponseBodyModuleItemInvoice struct {
	SupportSpecial *bool `json:"support_special,omitempty" xml:"support_special,omitempty"`
}

func (s HotelOrderPreValidateResponseBodyModuleItemInvoice) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderPreValidateResponseBodyModuleItemInvoice) GoString() string {
	return s.String()
}

func (s *HotelOrderPreValidateResponseBodyModuleItemInvoice) GetSupportSpecial() *bool {
	return s.SupportSpecial
}

func (s *HotelOrderPreValidateResponseBodyModuleItemInvoice) SetSupportSpecial(v bool) *HotelOrderPreValidateResponseBodyModuleItemInvoice {
	s.SupportSpecial = &v
	return s
}

func (s *HotelOrderPreValidateResponseBodyModuleItemInvoice) Validate() error {
	return dara.Validate(s)
}

type HotelOrderPreValidateResponseBodyModulePromotionInfo struct {
	ExtAttrMap              map[string]*string                                                             `json:"ext_attr_map,omitempty" xml:"ext_attr_map,omitempty"`
	PromotionDetailInfoList []*HotelOrderPreValidateResponseBodyModulePromotionInfoPromotionDetailInfoList `json:"promotion_detail_info_list,omitempty" xml:"promotion_detail_info_list,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	PromotionTotalPrice *int64 `json:"promotion_total_price,omitempty" xml:"promotion_total_price,omitempty"`
}

func (s HotelOrderPreValidateResponseBodyModulePromotionInfo) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderPreValidateResponseBodyModulePromotionInfo) GoString() string {
	return s.String()
}

func (s *HotelOrderPreValidateResponseBodyModulePromotionInfo) GetExtAttrMap() map[string]*string {
	return s.ExtAttrMap
}

func (s *HotelOrderPreValidateResponseBodyModulePromotionInfo) GetPromotionDetailInfoList() []*HotelOrderPreValidateResponseBodyModulePromotionInfoPromotionDetailInfoList {
	return s.PromotionDetailInfoList
}

func (s *HotelOrderPreValidateResponseBodyModulePromotionInfo) GetPromotionTotalPrice() *int64 {
	return s.PromotionTotalPrice
}

func (s *HotelOrderPreValidateResponseBodyModulePromotionInfo) SetExtAttrMap(v map[string]*string) *HotelOrderPreValidateResponseBodyModulePromotionInfo {
	s.ExtAttrMap = v
	return s
}

func (s *HotelOrderPreValidateResponseBodyModulePromotionInfo) SetPromotionDetailInfoList(v []*HotelOrderPreValidateResponseBodyModulePromotionInfoPromotionDetailInfoList) *HotelOrderPreValidateResponseBodyModulePromotionInfo {
	s.PromotionDetailInfoList = v
	return s
}

func (s *HotelOrderPreValidateResponseBodyModulePromotionInfo) SetPromotionTotalPrice(v int64) *HotelOrderPreValidateResponseBodyModulePromotionInfo {
	s.PromotionTotalPrice = &v
	return s
}

func (s *HotelOrderPreValidateResponseBodyModulePromotionInfo) Validate() error {
	return dara.Validate(s)
}

type HotelOrderPreValidateResponseBodyModulePromotionInfoPromotionDetailInfoList struct {
	// example:
	//
	// true
	CheckStatus *bool `json:"check_status,omitempty" xml:"check_status,omitempty"`
	// example:
	//
	// true
	NeedCheck     *bool   `json:"need_check,omitempty" xml:"need_check,omitempty"`
	PromotionCode *string `json:"promotion_code,omitempty" xml:"promotion_code,omitempty"`
	// example:
	//
	// 12893
	PromotionId *string `json:"promotion_id,omitempty" xml:"promotion_id,omitempty"`
	// example:
	//
	// demo
	PromotionName *string `json:"promotion_name,omitempty" xml:"promotion_name,omitempty"`
	// example:
	//
	// 100
	PromotionPrice *int64 `json:"promotion_price,omitempty" xml:"promotion_price,omitempty"`
	// example:
	//
	// 1
	PromotionType *string `json:"promotion_type,omitempty" xml:"promotion_type,omitempty"`
}

func (s HotelOrderPreValidateResponseBodyModulePromotionInfoPromotionDetailInfoList) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderPreValidateResponseBodyModulePromotionInfoPromotionDetailInfoList) GoString() string {
	return s.String()
}

func (s *HotelOrderPreValidateResponseBodyModulePromotionInfoPromotionDetailInfoList) GetCheckStatus() *bool {
	return s.CheckStatus
}

func (s *HotelOrderPreValidateResponseBodyModulePromotionInfoPromotionDetailInfoList) GetNeedCheck() *bool {
	return s.NeedCheck
}

func (s *HotelOrderPreValidateResponseBodyModulePromotionInfoPromotionDetailInfoList) GetPromotionCode() *string {
	return s.PromotionCode
}

func (s *HotelOrderPreValidateResponseBodyModulePromotionInfoPromotionDetailInfoList) GetPromotionId() *string {
	return s.PromotionId
}

func (s *HotelOrderPreValidateResponseBodyModulePromotionInfoPromotionDetailInfoList) GetPromotionName() *string {
	return s.PromotionName
}

func (s *HotelOrderPreValidateResponseBodyModulePromotionInfoPromotionDetailInfoList) GetPromotionPrice() *int64 {
	return s.PromotionPrice
}

func (s *HotelOrderPreValidateResponseBodyModulePromotionInfoPromotionDetailInfoList) GetPromotionType() *string {
	return s.PromotionType
}

func (s *HotelOrderPreValidateResponseBodyModulePromotionInfoPromotionDetailInfoList) SetCheckStatus(v bool) *HotelOrderPreValidateResponseBodyModulePromotionInfoPromotionDetailInfoList {
	s.CheckStatus = &v
	return s
}

func (s *HotelOrderPreValidateResponseBodyModulePromotionInfoPromotionDetailInfoList) SetNeedCheck(v bool) *HotelOrderPreValidateResponseBodyModulePromotionInfoPromotionDetailInfoList {
	s.NeedCheck = &v
	return s
}

func (s *HotelOrderPreValidateResponseBodyModulePromotionInfoPromotionDetailInfoList) SetPromotionCode(v string) *HotelOrderPreValidateResponseBodyModulePromotionInfoPromotionDetailInfoList {
	s.PromotionCode = &v
	return s
}

func (s *HotelOrderPreValidateResponseBodyModulePromotionInfoPromotionDetailInfoList) SetPromotionId(v string) *HotelOrderPreValidateResponseBodyModulePromotionInfoPromotionDetailInfoList {
	s.PromotionId = &v
	return s
}

func (s *HotelOrderPreValidateResponseBodyModulePromotionInfoPromotionDetailInfoList) SetPromotionName(v string) *HotelOrderPreValidateResponseBodyModulePromotionInfoPromotionDetailInfoList {
	s.PromotionName = &v
	return s
}

func (s *HotelOrderPreValidateResponseBodyModulePromotionInfoPromotionDetailInfoList) SetPromotionPrice(v int64) *HotelOrderPreValidateResponseBodyModulePromotionInfoPromotionDetailInfoList {
	s.PromotionPrice = &v
	return s
}

func (s *HotelOrderPreValidateResponseBodyModulePromotionInfoPromotionDetailInfoList) SetPromotionType(v string) *HotelOrderPreValidateResponseBodyModulePromotionInfoPromotionDetailInfoList {
	s.PromotionType = &v
	return s
}

func (s *HotelOrderPreValidateResponseBodyModulePromotionInfoPromotionDetailInfoList) Validate() error {
	return dara.Validate(s)
}

type HotelOrderPreValidateResponseBodyModuleRatePlanDaily struct {
	Board         *string `json:"board,omitempty" xml:"board,omitempty"`
	DiscountPrice *string `json:"discount_price,omitempty" xml:"discount_price,omitempty"`
	MaxBookingNum *int32  `json:"max_booking_num,omitempty" xml:"max_booking_num,omitempty"`
	// example:
	//
	// 100
	Price *int64 `json:"price,omitempty" xml:"price,omitempty"`
	// example:
	//
	// 2023-01-19
	RateStartTime *string `json:"rate_start_time,omitempty" xml:"rate_start_time,omitempty"`
	// example:
	//
	// 1
	RoomCount             *int32  `json:"room_count,omitempty" xml:"room_count,omitempty"`
	RoundingDiscountPrice *string `json:"rounding_discount_price,omitempty" xml:"rounding_discount_price,omitempty"`
	RoundingPrice         *string `json:"rounding_price,omitempty" xml:"rounding_price,omitempty"`
	// example:
	//
	// 100
	ServiceFee *int64 `json:"service_fee,omitempty" xml:"service_fee,omitempty"`
}

func (s HotelOrderPreValidateResponseBodyModuleRatePlanDaily) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderPreValidateResponseBodyModuleRatePlanDaily) GoString() string {
	return s.String()
}

func (s *HotelOrderPreValidateResponseBodyModuleRatePlanDaily) GetBoard() *string {
	return s.Board
}

func (s *HotelOrderPreValidateResponseBodyModuleRatePlanDaily) GetDiscountPrice() *string {
	return s.DiscountPrice
}

func (s *HotelOrderPreValidateResponseBodyModuleRatePlanDaily) GetMaxBookingNum() *int32 {
	return s.MaxBookingNum
}

func (s *HotelOrderPreValidateResponseBodyModuleRatePlanDaily) GetPrice() *int64 {
	return s.Price
}

func (s *HotelOrderPreValidateResponseBodyModuleRatePlanDaily) GetRateStartTime() *string {
	return s.RateStartTime
}

func (s *HotelOrderPreValidateResponseBodyModuleRatePlanDaily) GetRoomCount() *int32 {
	return s.RoomCount
}

func (s *HotelOrderPreValidateResponseBodyModuleRatePlanDaily) GetRoundingDiscountPrice() *string {
	return s.RoundingDiscountPrice
}

func (s *HotelOrderPreValidateResponseBodyModuleRatePlanDaily) GetRoundingPrice() *string {
	return s.RoundingPrice
}

func (s *HotelOrderPreValidateResponseBodyModuleRatePlanDaily) GetServiceFee() *int64 {
	return s.ServiceFee
}

func (s *HotelOrderPreValidateResponseBodyModuleRatePlanDaily) SetBoard(v string) *HotelOrderPreValidateResponseBodyModuleRatePlanDaily {
	s.Board = &v
	return s
}

func (s *HotelOrderPreValidateResponseBodyModuleRatePlanDaily) SetDiscountPrice(v string) *HotelOrderPreValidateResponseBodyModuleRatePlanDaily {
	s.DiscountPrice = &v
	return s
}

func (s *HotelOrderPreValidateResponseBodyModuleRatePlanDaily) SetMaxBookingNum(v int32) *HotelOrderPreValidateResponseBodyModuleRatePlanDaily {
	s.MaxBookingNum = &v
	return s
}

func (s *HotelOrderPreValidateResponseBodyModuleRatePlanDaily) SetPrice(v int64) *HotelOrderPreValidateResponseBodyModuleRatePlanDaily {
	s.Price = &v
	return s
}

func (s *HotelOrderPreValidateResponseBodyModuleRatePlanDaily) SetRateStartTime(v string) *HotelOrderPreValidateResponseBodyModuleRatePlanDaily {
	s.RateStartTime = &v
	return s
}

func (s *HotelOrderPreValidateResponseBodyModuleRatePlanDaily) SetRoomCount(v int32) *HotelOrderPreValidateResponseBodyModuleRatePlanDaily {
	s.RoomCount = &v
	return s
}

func (s *HotelOrderPreValidateResponseBodyModuleRatePlanDaily) SetRoundingDiscountPrice(v string) *HotelOrderPreValidateResponseBodyModuleRatePlanDaily {
	s.RoundingDiscountPrice = &v
	return s
}

func (s *HotelOrderPreValidateResponseBodyModuleRatePlanDaily) SetRoundingPrice(v string) *HotelOrderPreValidateResponseBodyModuleRatePlanDaily {
	s.RoundingPrice = &v
	return s
}

func (s *HotelOrderPreValidateResponseBodyModuleRatePlanDaily) SetServiceFee(v int64) *HotelOrderPreValidateResponseBodyModuleRatePlanDaily {
	s.ServiceFee = &v
	return s
}

func (s *HotelOrderPreValidateResponseBodyModuleRatePlanDaily) Validate() error {
	return dara.Validate(s)
}

type HotelOrderPreValidateResponseBodyModuleRatePlanInfo struct {
	BedDesc                   *string                                                                       `json:"bed_desc,omitempty" xml:"bed_desc,omitempty"`
	BtripHotelCancelPolicyDTO *HotelOrderPreValidateResponseBodyModuleRatePlanInfoBtripHotelCancelPolicyDTO `json:"btrip_hotel_cancel_policy_d_t_o,omitempty" xml:"btrip_hotel_cancel_policy_d_t_o,omitempty" type:"Struct"`
	// example:
	//
	// 12:00
	EarliestCheckInTime *string `json:"earliest_check_in_time,omitempty" xml:"earliest_check_in_time,omitempty"`
	// example:
	//
	// 17:00
	LatestCheckOutTime *string `json:"latest_check_out_time,omitempty" xml:"latest_check_out_time,omitempty"`
	MaxBookingNum      *int32  `json:"max_booking_num,omitempty" xml:"max_booking_num,omitempty"`
	// example:
	//
	// 1
	MaxOccupancyNum *int32 `json:"max_occupancy_num,omitempty" xml:"max_occupancy_num,omitempty"`
	// example:
	//
	// false
	NeedCertificate *bool `json:"need_certificate,omitempty" xml:"need_certificate,omitempty"`
	// example:
	//
	// false
	NeedEmail *bool `json:"need_email,omitempty" xml:"need_email,omitempty"`
	// example:
	//
	// false
	NeedEnglishName *bool `json:"need_english_name,omitempty" xml:"need_english_name,omitempty"`
	// example:
	//
	// 100
	TotalOrderPrice *int64 `json:"total_order_price,omitempty" xml:"total_order_price,omitempty"`
	// example:
	//
	// 100
	TotalRoomPrice *int64 `json:"total_room_price,omitempty" xml:"total_room_price,omitempty"`
}

func (s HotelOrderPreValidateResponseBodyModuleRatePlanInfo) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderPreValidateResponseBodyModuleRatePlanInfo) GoString() string {
	return s.String()
}

func (s *HotelOrderPreValidateResponseBodyModuleRatePlanInfo) GetBedDesc() *string {
	return s.BedDesc
}

func (s *HotelOrderPreValidateResponseBodyModuleRatePlanInfo) GetBtripHotelCancelPolicyDTO() *HotelOrderPreValidateResponseBodyModuleRatePlanInfoBtripHotelCancelPolicyDTO {
	return s.BtripHotelCancelPolicyDTO
}

func (s *HotelOrderPreValidateResponseBodyModuleRatePlanInfo) GetEarliestCheckInTime() *string {
	return s.EarliestCheckInTime
}

func (s *HotelOrderPreValidateResponseBodyModuleRatePlanInfo) GetLatestCheckOutTime() *string {
	return s.LatestCheckOutTime
}

func (s *HotelOrderPreValidateResponseBodyModuleRatePlanInfo) GetMaxBookingNum() *int32 {
	return s.MaxBookingNum
}

func (s *HotelOrderPreValidateResponseBodyModuleRatePlanInfo) GetMaxOccupancyNum() *int32 {
	return s.MaxOccupancyNum
}

func (s *HotelOrderPreValidateResponseBodyModuleRatePlanInfo) GetNeedCertificate() *bool {
	return s.NeedCertificate
}

func (s *HotelOrderPreValidateResponseBodyModuleRatePlanInfo) GetNeedEmail() *bool {
	return s.NeedEmail
}

func (s *HotelOrderPreValidateResponseBodyModuleRatePlanInfo) GetNeedEnglishName() *bool {
	return s.NeedEnglishName
}

func (s *HotelOrderPreValidateResponseBodyModuleRatePlanInfo) GetTotalOrderPrice() *int64 {
	return s.TotalOrderPrice
}

func (s *HotelOrderPreValidateResponseBodyModuleRatePlanInfo) GetTotalRoomPrice() *int64 {
	return s.TotalRoomPrice
}

func (s *HotelOrderPreValidateResponseBodyModuleRatePlanInfo) SetBedDesc(v string) *HotelOrderPreValidateResponseBodyModuleRatePlanInfo {
	s.BedDesc = &v
	return s
}

func (s *HotelOrderPreValidateResponseBodyModuleRatePlanInfo) SetBtripHotelCancelPolicyDTO(v *HotelOrderPreValidateResponseBodyModuleRatePlanInfoBtripHotelCancelPolicyDTO) *HotelOrderPreValidateResponseBodyModuleRatePlanInfo {
	s.BtripHotelCancelPolicyDTO = v
	return s
}

func (s *HotelOrderPreValidateResponseBodyModuleRatePlanInfo) SetEarliestCheckInTime(v string) *HotelOrderPreValidateResponseBodyModuleRatePlanInfo {
	s.EarliestCheckInTime = &v
	return s
}

func (s *HotelOrderPreValidateResponseBodyModuleRatePlanInfo) SetLatestCheckOutTime(v string) *HotelOrderPreValidateResponseBodyModuleRatePlanInfo {
	s.LatestCheckOutTime = &v
	return s
}

func (s *HotelOrderPreValidateResponseBodyModuleRatePlanInfo) SetMaxBookingNum(v int32) *HotelOrderPreValidateResponseBodyModuleRatePlanInfo {
	s.MaxBookingNum = &v
	return s
}

func (s *HotelOrderPreValidateResponseBodyModuleRatePlanInfo) SetMaxOccupancyNum(v int32) *HotelOrderPreValidateResponseBodyModuleRatePlanInfo {
	s.MaxOccupancyNum = &v
	return s
}

func (s *HotelOrderPreValidateResponseBodyModuleRatePlanInfo) SetNeedCertificate(v bool) *HotelOrderPreValidateResponseBodyModuleRatePlanInfo {
	s.NeedCertificate = &v
	return s
}

func (s *HotelOrderPreValidateResponseBodyModuleRatePlanInfo) SetNeedEmail(v bool) *HotelOrderPreValidateResponseBodyModuleRatePlanInfo {
	s.NeedEmail = &v
	return s
}

func (s *HotelOrderPreValidateResponseBodyModuleRatePlanInfo) SetNeedEnglishName(v bool) *HotelOrderPreValidateResponseBodyModuleRatePlanInfo {
	s.NeedEnglishName = &v
	return s
}

func (s *HotelOrderPreValidateResponseBodyModuleRatePlanInfo) SetTotalOrderPrice(v int64) *HotelOrderPreValidateResponseBodyModuleRatePlanInfo {
	s.TotalOrderPrice = &v
	return s
}

func (s *HotelOrderPreValidateResponseBodyModuleRatePlanInfo) SetTotalRoomPrice(v int64) *HotelOrderPreValidateResponseBodyModuleRatePlanInfo {
	s.TotalRoomPrice = &v
	return s
}

func (s *HotelOrderPreValidateResponseBodyModuleRatePlanInfo) Validate() error {
	return dara.Validate(s)
}

type HotelOrderPreValidateResponseBodyModuleRatePlanInfoBtripHotelCancelPolicyDTO struct {
	BtripHotelCancelPolicyInfoDTOList []*HotelOrderPreValidateResponseBodyModuleRatePlanInfoBtripHotelCancelPolicyDTOBtripHotelCancelPolicyInfoDTOList `json:"btrip_hotel_cancel_policy_info_d_t_o_list,omitempty" xml:"btrip_hotel_cancel_policy_info_d_t_o_list,omitempty" type:"Repeated"`
	CancelPolicyType                  *int32                                                                                                           `json:"cancel_policy_type,omitempty" xml:"cancel_policy_type,omitempty"`
	Content                           *string                                                                                                          `json:"content,omitempty" xml:"content,omitempty"`
	ShortDesc                         *string                                                                                                          `json:"short_desc,omitempty" xml:"short_desc,omitempty"`
}

func (s HotelOrderPreValidateResponseBodyModuleRatePlanInfoBtripHotelCancelPolicyDTO) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderPreValidateResponseBodyModuleRatePlanInfoBtripHotelCancelPolicyDTO) GoString() string {
	return s.String()
}

func (s *HotelOrderPreValidateResponseBodyModuleRatePlanInfoBtripHotelCancelPolicyDTO) GetBtripHotelCancelPolicyInfoDTOList() []*HotelOrderPreValidateResponseBodyModuleRatePlanInfoBtripHotelCancelPolicyDTOBtripHotelCancelPolicyInfoDTOList {
	return s.BtripHotelCancelPolicyInfoDTOList
}

func (s *HotelOrderPreValidateResponseBodyModuleRatePlanInfoBtripHotelCancelPolicyDTO) GetCancelPolicyType() *int32 {
	return s.CancelPolicyType
}

func (s *HotelOrderPreValidateResponseBodyModuleRatePlanInfoBtripHotelCancelPolicyDTO) GetContent() *string {
	return s.Content
}

func (s *HotelOrderPreValidateResponseBodyModuleRatePlanInfoBtripHotelCancelPolicyDTO) GetShortDesc() *string {
	return s.ShortDesc
}

func (s *HotelOrderPreValidateResponseBodyModuleRatePlanInfoBtripHotelCancelPolicyDTO) SetBtripHotelCancelPolicyInfoDTOList(v []*HotelOrderPreValidateResponseBodyModuleRatePlanInfoBtripHotelCancelPolicyDTOBtripHotelCancelPolicyInfoDTOList) *HotelOrderPreValidateResponseBodyModuleRatePlanInfoBtripHotelCancelPolicyDTO {
	s.BtripHotelCancelPolicyInfoDTOList = v
	return s
}

func (s *HotelOrderPreValidateResponseBodyModuleRatePlanInfoBtripHotelCancelPolicyDTO) SetCancelPolicyType(v int32) *HotelOrderPreValidateResponseBodyModuleRatePlanInfoBtripHotelCancelPolicyDTO {
	s.CancelPolicyType = &v
	return s
}

func (s *HotelOrderPreValidateResponseBodyModuleRatePlanInfoBtripHotelCancelPolicyDTO) SetContent(v string) *HotelOrderPreValidateResponseBodyModuleRatePlanInfoBtripHotelCancelPolicyDTO {
	s.Content = &v
	return s
}

func (s *HotelOrderPreValidateResponseBodyModuleRatePlanInfoBtripHotelCancelPolicyDTO) SetShortDesc(v string) *HotelOrderPreValidateResponseBodyModuleRatePlanInfoBtripHotelCancelPolicyDTO {
	s.ShortDesc = &v
	return s
}

func (s *HotelOrderPreValidateResponseBodyModuleRatePlanInfoBtripHotelCancelPolicyDTO) Validate() error {
	return dara.Validate(s)
}

type HotelOrderPreValidateResponseBodyModuleRatePlanInfoBtripHotelCancelPolicyDTOBtripHotelCancelPolicyInfoDTOList struct {
	// example:
	//
	// 1
	Hour *int64 `json:"hour,omitempty" xml:"hour,omitempty"`
	// example:
	//
	// 1
	Value *int64 `json:"value,omitempty" xml:"value,omitempty"`
}

func (s HotelOrderPreValidateResponseBodyModuleRatePlanInfoBtripHotelCancelPolicyDTOBtripHotelCancelPolicyInfoDTOList) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderPreValidateResponseBodyModuleRatePlanInfoBtripHotelCancelPolicyDTOBtripHotelCancelPolicyInfoDTOList) GoString() string {
	return s.String()
}

func (s *HotelOrderPreValidateResponseBodyModuleRatePlanInfoBtripHotelCancelPolicyDTOBtripHotelCancelPolicyInfoDTOList) GetHour() *int64 {
	return s.Hour
}

func (s *HotelOrderPreValidateResponseBodyModuleRatePlanInfoBtripHotelCancelPolicyDTOBtripHotelCancelPolicyInfoDTOList) GetValue() *int64 {
	return s.Value
}

func (s *HotelOrderPreValidateResponseBodyModuleRatePlanInfoBtripHotelCancelPolicyDTOBtripHotelCancelPolicyInfoDTOList) SetHour(v int64) *HotelOrderPreValidateResponseBodyModuleRatePlanInfoBtripHotelCancelPolicyDTOBtripHotelCancelPolicyInfoDTOList {
	s.Hour = &v
	return s
}

func (s *HotelOrderPreValidateResponseBodyModuleRatePlanInfoBtripHotelCancelPolicyDTOBtripHotelCancelPolicyInfoDTOList) SetValue(v int64) *HotelOrderPreValidateResponseBodyModuleRatePlanInfoBtripHotelCancelPolicyDTOBtripHotelCancelPolicyInfoDTOList {
	s.Value = &v
	return s
}

func (s *HotelOrderPreValidateResponseBodyModuleRatePlanInfoBtripHotelCancelPolicyDTOBtripHotelCancelPolicyInfoDTOList) Validate() error {
	return dara.Validate(s)
}
