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
	Code      *string                                  `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                                  `json:"message,omitempty" xml:"message,omitempty"`
	Module    *HotelOrderPreValidateResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	RequestId *string                                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                                    `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string                                  `json:"traceId,omitempty" xml:"traceId,omitempty"`
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
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type HotelOrderPreValidateResponseBodyModule struct {
	ExtendInfo     *string                                                 `json:"extend_info,omitempty" xml:"extend_info,omitempty"`
	ItemInvoice    *HotelOrderPreValidateResponseBodyModuleItemInvoice     `json:"item_invoice,omitempty" xml:"item_invoice,omitempty" type:"Struct"`
	ItineraryNo    *string                                                 `json:"itinerary_no,omitempty" xml:"itinerary_no,omitempty"`
	PromotionInfo  *HotelOrderPreValidateResponseBodyModulePromotionInfo   `json:"promotion_info,omitempty" xml:"promotion_info,omitempty" type:"Struct"`
	RatePlanDaily  []*HotelOrderPreValidateResponseBodyModuleRatePlanDaily `json:"rate_plan_daily,omitempty" xml:"rate_plan_daily,omitempty" type:"Repeated"`
	RatePlanId     *int64                                                  `json:"rate_plan_id,omitempty" xml:"rate_plan_id,omitempty"`
	RatePlanInfo   *HotelOrderPreValidateResponseBodyModuleRatePlanInfo    `json:"rate_plan_info,omitempty" xml:"rate_plan_info,omitempty" type:"Struct"`
	ValidateResKey *string                                                 `json:"validate_res_key,omitempty" xml:"validate_res_key,omitempty"`
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
	if s.ItemInvoice != nil {
		if err := s.ItemInvoice.Validate(); err != nil {
			return err
		}
	}
	if s.PromotionInfo != nil {
		if err := s.PromotionInfo.Validate(); err != nil {
			return err
		}
	}
	if s.RatePlanDaily != nil {
		for _, item := range s.RatePlanDaily {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RatePlanInfo != nil {
		if err := s.RatePlanInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
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
	PromotionTotalPrice     *int64                                                                         `json:"promotion_total_price,omitempty" xml:"promotion_total_price,omitempty"`
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
	if s.PromotionDetailInfoList != nil {
		for _, item := range s.PromotionDetailInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type HotelOrderPreValidateResponseBodyModulePromotionInfoPromotionDetailInfoList struct {
	CheckStatus    *bool   `json:"check_status,omitempty" xml:"check_status,omitempty"`
	NeedCheck      *bool   `json:"need_check,omitempty" xml:"need_check,omitempty"`
	PromotionCode  *string `json:"promotion_code,omitempty" xml:"promotion_code,omitempty"`
	PromotionId    *string `json:"promotion_id,omitempty" xml:"promotion_id,omitempty"`
	PromotionName  *string `json:"promotion_name,omitempty" xml:"promotion_name,omitempty"`
	PromotionPrice *int64  `json:"promotion_price,omitempty" xml:"promotion_price,omitempty"`
	PromotionType  *string `json:"promotion_type,omitempty" xml:"promotion_type,omitempty"`
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
	Board                 *string `json:"board,omitempty" xml:"board,omitempty"`
	DiscountPrice         *string `json:"discount_price,omitempty" xml:"discount_price,omitempty"`
	MaxBookingNum         *int32  `json:"max_booking_num,omitempty" xml:"max_booking_num,omitempty"`
	Price                 *int64  `json:"price,omitempty" xml:"price,omitempty"`
	RateStartTime         *string `json:"rate_start_time,omitempty" xml:"rate_start_time,omitempty"`
	RoomCount             *int32  `json:"room_count,omitempty" xml:"room_count,omitempty"`
	RoundingDiscountPrice *string `json:"rounding_discount_price,omitempty" xml:"rounding_discount_price,omitempty"`
	RoundingPrice         *string `json:"rounding_price,omitempty" xml:"rounding_price,omitempty"`
	ServiceFee            *int64  `json:"service_fee,omitempty" xml:"service_fee,omitempty"`
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
	CertTypeList              []*string                                                                     `json:"cert_type_list,omitempty" xml:"cert_type_list,omitempty" type:"Repeated"`
	EarliestCheckInTime       *string                                                                       `json:"earliest_check_in_time,omitempty" xml:"earliest_check_in_time,omitempty"`
	HourItemArrivalTimeInfo   *HotelOrderPreValidateResponseBodyModuleRatePlanInfoHourItemArrivalTimeInfo   `json:"hour_item_arrival_time_info,omitempty" xml:"hour_item_arrival_time_info,omitempty" type:"Struct"`
	LatestCheckOutTime        *string                                                                       `json:"latest_check_out_time,omitempty" xml:"latest_check_out_time,omitempty"`
	MaxBookingNum             *int32                                                                        `json:"max_booking_num,omitempty" xml:"max_booking_num,omitempty"`
	MaxOccupancyNum           *int32                                                                        `json:"max_occupancy_num,omitempty" xml:"max_occupancy_num,omitempty"`
	NeedCertificate           *bool                                                                         `json:"need_certificate,omitempty" xml:"need_certificate,omitempty"`
	NeedEmail                 *bool                                                                         `json:"need_email,omitempty" xml:"need_email,omitempty"`
	NeedEnglishName           *bool                                                                         `json:"need_english_name,omitempty" xml:"need_english_name,omitempty"`
	RpType                    *int32                                                                        `json:"rp_type,omitempty" xml:"rp_type,omitempty"`
	TotalOrderPrice           *int64                                                                        `json:"total_order_price,omitempty" xml:"total_order_price,omitempty"`
	TotalRoomPrice            *int64                                                                        `json:"total_room_price,omitempty" xml:"total_room_price,omitempty"`
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

func (s *HotelOrderPreValidateResponseBodyModuleRatePlanInfo) GetCertTypeList() []*string {
	return s.CertTypeList
}

func (s *HotelOrderPreValidateResponseBodyModuleRatePlanInfo) GetEarliestCheckInTime() *string {
	return s.EarliestCheckInTime
}

func (s *HotelOrderPreValidateResponseBodyModuleRatePlanInfo) GetHourItemArrivalTimeInfo() *HotelOrderPreValidateResponseBodyModuleRatePlanInfoHourItemArrivalTimeInfo {
	return s.HourItemArrivalTimeInfo
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

func (s *HotelOrderPreValidateResponseBodyModuleRatePlanInfo) GetRpType() *int32 {
	return s.RpType
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

func (s *HotelOrderPreValidateResponseBodyModuleRatePlanInfo) SetCertTypeList(v []*string) *HotelOrderPreValidateResponseBodyModuleRatePlanInfo {
	s.CertTypeList = v
	return s
}

func (s *HotelOrderPreValidateResponseBodyModuleRatePlanInfo) SetEarliestCheckInTime(v string) *HotelOrderPreValidateResponseBodyModuleRatePlanInfo {
	s.EarliestCheckInTime = &v
	return s
}

func (s *HotelOrderPreValidateResponseBodyModuleRatePlanInfo) SetHourItemArrivalTimeInfo(v *HotelOrderPreValidateResponseBodyModuleRatePlanInfoHourItemArrivalTimeInfo) *HotelOrderPreValidateResponseBodyModuleRatePlanInfo {
	s.HourItemArrivalTimeInfo = v
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

func (s *HotelOrderPreValidateResponseBodyModuleRatePlanInfo) SetRpType(v int32) *HotelOrderPreValidateResponseBodyModuleRatePlanInfo {
	s.RpType = &v
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
	if s.BtripHotelCancelPolicyDTO != nil {
		if err := s.BtripHotelCancelPolicyDTO.Validate(); err != nil {
			return err
		}
	}
	if s.HourItemArrivalTimeInfo != nil {
		if err := s.HourItemArrivalTimeInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
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
	if s.BtripHotelCancelPolicyInfoDTOList != nil {
		for _, item := range s.BtripHotelCancelPolicyInfoDTOList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type HotelOrderPreValidateResponseBodyModuleRatePlanInfoBtripHotelCancelPolicyDTOBtripHotelCancelPolicyInfoDTOList struct {
	Hour  *int64 `json:"hour,omitempty" xml:"hour,omitempty"`
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

type HotelOrderPreValidateResponseBodyModuleRatePlanInfoHourItemArrivalTimeInfo struct {
	ActualLiveHour      *string `json:"actual_live_hour,omitempty" xml:"actual_live_hour,omitempty"`
	EarliestCheckInTime *string `json:"earliest_check_in_time,omitempty" xml:"earliest_check_in_time,omitempty"`
	LatestCheckInTime   *string `json:"latest_check_in_time,omitempty" xml:"latest_check_in_time,omitempty"`
	LatestCheckOutTime  *string `json:"latest_check_out_time,omitempty" xml:"latest_check_out_time,omitempty"`
	LiveHour            *string `json:"live_hour,omitempty" xml:"live_hour,omitempty"`
}

func (s HotelOrderPreValidateResponseBodyModuleRatePlanInfoHourItemArrivalTimeInfo) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderPreValidateResponseBodyModuleRatePlanInfoHourItemArrivalTimeInfo) GoString() string {
	return s.String()
}

func (s *HotelOrderPreValidateResponseBodyModuleRatePlanInfoHourItemArrivalTimeInfo) GetActualLiveHour() *string {
	return s.ActualLiveHour
}

func (s *HotelOrderPreValidateResponseBodyModuleRatePlanInfoHourItemArrivalTimeInfo) GetEarliestCheckInTime() *string {
	return s.EarliestCheckInTime
}

func (s *HotelOrderPreValidateResponseBodyModuleRatePlanInfoHourItemArrivalTimeInfo) GetLatestCheckInTime() *string {
	return s.LatestCheckInTime
}

func (s *HotelOrderPreValidateResponseBodyModuleRatePlanInfoHourItemArrivalTimeInfo) GetLatestCheckOutTime() *string {
	return s.LatestCheckOutTime
}

func (s *HotelOrderPreValidateResponseBodyModuleRatePlanInfoHourItemArrivalTimeInfo) GetLiveHour() *string {
	return s.LiveHour
}

func (s *HotelOrderPreValidateResponseBodyModuleRatePlanInfoHourItemArrivalTimeInfo) SetActualLiveHour(v string) *HotelOrderPreValidateResponseBodyModuleRatePlanInfoHourItemArrivalTimeInfo {
	s.ActualLiveHour = &v
	return s
}

func (s *HotelOrderPreValidateResponseBodyModuleRatePlanInfoHourItemArrivalTimeInfo) SetEarliestCheckInTime(v string) *HotelOrderPreValidateResponseBodyModuleRatePlanInfoHourItemArrivalTimeInfo {
	s.EarliestCheckInTime = &v
	return s
}

func (s *HotelOrderPreValidateResponseBodyModuleRatePlanInfoHourItemArrivalTimeInfo) SetLatestCheckInTime(v string) *HotelOrderPreValidateResponseBodyModuleRatePlanInfoHourItemArrivalTimeInfo {
	s.LatestCheckInTime = &v
	return s
}

func (s *HotelOrderPreValidateResponseBodyModuleRatePlanInfoHourItemArrivalTimeInfo) SetLatestCheckOutTime(v string) *HotelOrderPreValidateResponseBodyModuleRatePlanInfoHourItemArrivalTimeInfo {
	s.LatestCheckOutTime = &v
	return s
}

func (s *HotelOrderPreValidateResponseBodyModuleRatePlanInfoHourItemArrivalTimeInfo) SetLiveHour(v string) *HotelOrderPreValidateResponseBodyModuleRatePlanInfoHourItemArrivalTimeInfo {
	s.LiveHour = &v
	return s
}

func (s *HotelOrderPreValidateResponseBodyModuleRatePlanInfoHourItemArrivalTimeInfo) Validate() error {
	return dara.Validate(s)
}
