// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelOrderInfoQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *HotelOrderInfoQueryResponseBody
	GetCode() *string
	SetMessage(v string) *HotelOrderInfoQueryResponseBody
	GetMessage() *string
	SetModule(v *HotelOrderInfoQueryResponseBodyModule) *HotelOrderInfoQueryResponseBody
	GetModule() *HotelOrderInfoQueryResponseBodyModule
	SetRequestId(v string) *HotelOrderInfoQueryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *HotelOrderInfoQueryResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *HotelOrderInfoQueryResponseBody
	GetTraceId() *string
}

type HotelOrderInfoQueryResponseBody struct {
	// example:
	//
	// success
	Code    *string                                `json:"code,omitempty" xml:"code,omitempty"`
	Message *string                                `json:"message,omitempty" xml:"message,omitempty"`
	Module  *HotelOrderInfoQueryResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
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
	// 210f079e16603757182131635d866a
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s HotelOrderInfoQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderInfoQueryResponseBody) GoString() string {
	return s.String()
}

func (s *HotelOrderInfoQueryResponseBody) GetCode() *string {
	return s.Code
}

func (s *HotelOrderInfoQueryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *HotelOrderInfoQueryResponseBody) GetModule() *HotelOrderInfoQueryResponseBodyModule {
	return s.Module
}

func (s *HotelOrderInfoQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *HotelOrderInfoQueryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *HotelOrderInfoQueryResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *HotelOrderInfoQueryResponseBody) SetCode(v string) *HotelOrderInfoQueryResponseBody {
	s.Code = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBody) SetMessage(v string) *HotelOrderInfoQueryResponseBody {
	s.Message = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBody) SetModule(v *HotelOrderInfoQueryResponseBodyModule) *HotelOrderInfoQueryResponseBody {
	s.Module = v
	return s
}

func (s *HotelOrderInfoQueryResponseBody) SetRequestId(v string) *HotelOrderInfoQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBody) SetSuccess(v bool) *HotelOrderInfoQueryResponseBody {
	s.Success = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBody) SetTraceId(v string) *HotelOrderInfoQueryResponseBody {
	s.TraceId = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBody) Validate() error {
	return dara.Validate(s)
}

type HotelOrderInfoQueryResponseBodyModule struct {
	BaseOrderInfo        *HotelOrderInfoQueryResponseBodyModuleBaseOrderInfo          `json:"base_order_info,omitempty" xml:"base_order_info,omitempty" type:"Struct"`
	BookerInfo           *HotelOrderInfoQueryResponseBodyModuleBookerInfo             `json:"booker_info,omitempty" xml:"booker_info,omitempty" type:"Struct"`
	HotelInfo            *HotelOrderInfoQueryResponseBodyModuleHotelInfo              `json:"hotel_info,omitempty" xml:"hotel_info,omitempty" type:"Struct"`
	HotelOrderFeeInfo    *HotelOrderInfoQueryResponseBodyModuleHotelOrderFeeInfo      `json:"hotel_order_fee_info,omitempty" xml:"hotel_order_fee_info,omitempty" type:"Struct"`
	HotelOrderRefundInfo []*HotelOrderInfoQueryResponseBodyModuleHotelOrderRefundInfo `json:"hotel_order_refund_info,omitempty" xml:"hotel_order_refund_info,omitempty" type:"Repeated"`
	RoomTraverInfo       []*HotelOrderInfoQueryResponseBodyModuleRoomTraverInfo       `json:"room_traver_info,omitempty" xml:"room_traver_info,omitempty" type:"Repeated"`
}

func (s HotelOrderInfoQueryResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderInfoQueryResponseBodyModule) GoString() string {
	return s.String()
}

func (s *HotelOrderInfoQueryResponseBodyModule) GetBaseOrderInfo() *HotelOrderInfoQueryResponseBodyModuleBaseOrderInfo {
	return s.BaseOrderInfo
}

func (s *HotelOrderInfoQueryResponseBodyModule) GetBookerInfo() *HotelOrderInfoQueryResponseBodyModuleBookerInfo {
	return s.BookerInfo
}

func (s *HotelOrderInfoQueryResponseBodyModule) GetHotelInfo() *HotelOrderInfoQueryResponseBodyModuleHotelInfo {
	return s.HotelInfo
}

func (s *HotelOrderInfoQueryResponseBodyModule) GetHotelOrderFeeInfo() *HotelOrderInfoQueryResponseBodyModuleHotelOrderFeeInfo {
	return s.HotelOrderFeeInfo
}

func (s *HotelOrderInfoQueryResponseBodyModule) GetHotelOrderRefundInfo() []*HotelOrderInfoQueryResponseBodyModuleHotelOrderRefundInfo {
	return s.HotelOrderRefundInfo
}

func (s *HotelOrderInfoQueryResponseBodyModule) GetRoomTraverInfo() []*HotelOrderInfoQueryResponseBodyModuleRoomTraverInfo {
	return s.RoomTraverInfo
}

func (s *HotelOrderInfoQueryResponseBodyModule) SetBaseOrderInfo(v *HotelOrderInfoQueryResponseBodyModuleBaseOrderInfo) *HotelOrderInfoQueryResponseBodyModule {
	s.BaseOrderInfo = v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModule) SetBookerInfo(v *HotelOrderInfoQueryResponseBodyModuleBookerInfo) *HotelOrderInfoQueryResponseBodyModule {
	s.BookerInfo = v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModule) SetHotelInfo(v *HotelOrderInfoQueryResponseBodyModuleHotelInfo) *HotelOrderInfoQueryResponseBodyModule {
	s.HotelInfo = v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModule) SetHotelOrderFeeInfo(v *HotelOrderInfoQueryResponseBodyModuleHotelOrderFeeInfo) *HotelOrderInfoQueryResponseBodyModule {
	s.HotelOrderFeeInfo = v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModule) SetHotelOrderRefundInfo(v []*HotelOrderInfoQueryResponseBodyModuleHotelOrderRefundInfo) *HotelOrderInfoQueryResponseBodyModule {
	s.HotelOrderRefundInfo = v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModule) SetRoomTraverInfo(v []*HotelOrderInfoQueryResponseBodyModuleRoomTraverInfo) *HotelOrderInfoQueryResponseBodyModule {
	s.RoomTraverInfo = v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type HotelOrderInfoQueryResponseBodyModuleBaseOrderInfo struct {
	// example:
	//
	// 0
	BookMode *string `json:"book_mode,omitempty" xml:"book_mode,omitempty"`
	// example:
	//
	// 1430378
	BookerId   *string `json:"booker_id,omitempty" xml:"booker_id,omitempty"`
	BookerName *string `json:"booker_name,omitempty" xml:"booker_name,omitempty"`
	// example:
	//
	// btripkvxtn1321g49wtul
	BtripCorpId *string `json:"btrip_corp_id,omitempty" xml:"btrip_corp_id,omitempty"`
	// example:
	//
	// 2
	Category *int32 `json:"category,omitempty" xml:"category,omitempty"`
	// example:
	//
	// 1721145600000
	CheckInTime *int64 `json:"check_in_time,omitempty" xml:"check_in_time,omitempty"`
	// example:
	//
	// 1721145600000
	CheckOutTime *int64 `json:"check_out_time,omitempty" xml:"check_out_time,omitempty"`
	// example:
	//
	// true
	IsAgreementPrice *bool `json:"is_agreement_price,omitempty" xml:"is_agreement_price,omitempty"`
	// example:
	//
	// 4
	Nights *int32 `json:"nights,omitempty" xml:"nights,omitempty"`
	// example:
	//
	// 1721145600000
	OrderCreateTime *int64 `json:"order_create_time,omitempty" xml:"order_create_time,omitempty"`
	// example:
	//
	// 1012053198307958626
	OrderId *int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// example:
	//
	// 2
	OrderStatus     *int32  `json:"order_status,omitempty" xml:"order_status,omitempty"`
	OrderStatusDesc *string `json:"order_status_desc,omitempty" xml:"order_status_desc,omitempty"`
	// example:
	//
	// 1
	PayStatus *int32 `json:"pay_status,omitempty" xml:"pay_status,omitempty"`
	// example:
	//
	// 1721145600000
	PayTime *int64 `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// example:
	//
	// 2
	RoomNum *int32 `json:"room_num,omitempty" xml:"room_num,omitempty"`
	// example:
	//
	// 4
	SettleType *int32 `json:"settle_type,omitempty" xml:"settle_type,omitempty"`
	// example:
	//
	// 1
	TripMode *int32 `json:"trip_mode,omitempty" xml:"trip_mode,omitempty"`
}

func (s HotelOrderInfoQueryResponseBodyModuleBaseOrderInfo) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderInfoQueryResponseBodyModuleBaseOrderInfo) GoString() string {
	return s.String()
}

func (s *HotelOrderInfoQueryResponseBodyModuleBaseOrderInfo) GetBookMode() *string {
	return s.BookMode
}

func (s *HotelOrderInfoQueryResponseBodyModuleBaseOrderInfo) GetBookerId() *string {
	return s.BookerId
}

func (s *HotelOrderInfoQueryResponseBodyModuleBaseOrderInfo) GetBookerName() *string {
	return s.BookerName
}

func (s *HotelOrderInfoQueryResponseBodyModuleBaseOrderInfo) GetBtripCorpId() *string {
	return s.BtripCorpId
}

func (s *HotelOrderInfoQueryResponseBodyModuleBaseOrderInfo) GetCategory() *int32 {
	return s.Category
}

func (s *HotelOrderInfoQueryResponseBodyModuleBaseOrderInfo) GetCheckInTime() *int64 {
	return s.CheckInTime
}

func (s *HotelOrderInfoQueryResponseBodyModuleBaseOrderInfo) GetCheckOutTime() *int64 {
	return s.CheckOutTime
}

func (s *HotelOrderInfoQueryResponseBodyModuleBaseOrderInfo) GetIsAgreementPrice() *bool {
	return s.IsAgreementPrice
}

func (s *HotelOrderInfoQueryResponseBodyModuleBaseOrderInfo) GetNights() *int32 {
	return s.Nights
}

func (s *HotelOrderInfoQueryResponseBodyModuleBaseOrderInfo) GetOrderCreateTime() *int64 {
	return s.OrderCreateTime
}

func (s *HotelOrderInfoQueryResponseBodyModuleBaseOrderInfo) GetOrderId() *int64 {
	return s.OrderId
}

func (s *HotelOrderInfoQueryResponseBodyModuleBaseOrderInfo) GetOrderStatus() *int32 {
	return s.OrderStatus
}

func (s *HotelOrderInfoQueryResponseBodyModuleBaseOrderInfo) GetOrderStatusDesc() *string {
	return s.OrderStatusDesc
}

func (s *HotelOrderInfoQueryResponseBodyModuleBaseOrderInfo) GetPayStatus() *int32 {
	return s.PayStatus
}

func (s *HotelOrderInfoQueryResponseBodyModuleBaseOrderInfo) GetPayTime() *int64 {
	return s.PayTime
}

func (s *HotelOrderInfoQueryResponseBodyModuleBaseOrderInfo) GetRoomNum() *int32 {
	return s.RoomNum
}

func (s *HotelOrderInfoQueryResponseBodyModuleBaseOrderInfo) GetSettleType() *int32 {
	return s.SettleType
}

func (s *HotelOrderInfoQueryResponseBodyModuleBaseOrderInfo) GetTripMode() *int32 {
	return s.TripMode
}

func (s *HotelOrderInfoQueryResponseBodyModuleBaseOrderInfo) SetBookMode(v string) *HotelOrderInfoQueryResponseBodyModuleBaseOrderInfo {
	s.BookMode = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleBaseOrderInfo) SetBookerId(v string) *HotelOrderInfoQueryResponseBodyModuleBaseOrderInfo {
	s.BookerId = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleBaseOrderInfo) SetBookerName(v string) *HotelOrderInfoQueryResponseBodyModuleBaseOrderInfo {
	s.BookerName = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleBaseOrderInfo) SetBtripCorpId(v string) *HotelOrderInfoQueryResponseBodyModuleBaseOrderInfo {
	s.BtripCorpId = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleBaseOrderInfo) SetCategory(v int32) *HotelOrderInfoQueryResponseBodyModuleBaseOrderInfo {
	s.Category = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleBaseOrderInfo) SetCheckInTime(v int64) *HotelOrderInfoQueryResponseBodyModuleBaseOrderInfo {
	s.CheckInTime = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleBaseOrderInfo) SetCheckOutTime(v int64) *HotelOrderInfoQueryResponseBodyModuleBaseOrderInfo {
	s.CheckOutTime = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleBaseOrderInfo) SetIsAgreementPrice(v bool) *HotelOrderInfoQueryResponseBodyModuleBaseOrderInfo {
	s.IsAgreementPrice = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleBaseOrderInfo) SetNights(v int32) *HotelOrderInfoQueryResponseBodyModuleBaseOrderInfo {
	s.Nights = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleBaseOrderInfo) SetOrderCreateTime(v int64) *HotelOrderInfoQueryResponseBodyModuleBaseOrderInfo {
	s.OrderCreateTime = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleBaseOrderInfo) SetOrderId(v int64) *HotelOrderInfoQueryResponseBodyModuleBaseOrderInfo {
	s.OrderId = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleBaseOrderInfo) SetOrderStatus(v int32) *HotelOrderInfoQueryResponseBodyModuleBaseOrderInfo {
	s.OrderStatus = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleBaseOrderInfo) SetOrderStatusDesc(v string) *HotelOrderInfoQueryResponseBodyModuleBaseOrderInfo {
	s.OrderStatusDesc = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleBaseOrderInfo) SetPayStatus(v int32) *HotelOrderInfoQueryResponseBodyModuleBaseOrderInfo {
	s.PayStatus = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleBaseOrderInfo) SetPayTime(v int64) *HotelOrderInfoQueryResponseBodyModuleBaseOrderInfo {
	s.PayTime = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleBaseOrderInfo) SetRoomNum(v int32) *HotelOrderInfoQueryResponseBodyModuleBaseOrderInfo {
	s.RoomNum = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleBaseOrderInfo) SetSettleType(v int32) *HotelOrderInfoQueryResponseBodyModuleBaseOrderInfo {
	s.SettleType = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleBaseOrderInfo) SetTripMode(v int32) *HotelOrderInfoQueryResponseBodyModuleBaseOrderInfo {
	s.TripMode = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleBaseOrderInfo) Validate() error {
	return dara.Validate(s)
}

type HotelOrderInfoQueryResponseBodyModuleBookerInfo struct {
	// example:
	//
	// UN_APPLY
	BookerRole *string `json:"booker_role,omitempty" xml:"booker_role,omitempty"`
	// example:
	//
	// 13311112222@qq.com
	ContactEmail *string `json:"contact_email,omitempty" xml:"contact_email,omitempty"`
	// example:
	//
	// 13311112222
	ContactPhone *string `json:"contact_phone,omitempty" xml:"contact_phone,omitempty"`
	// example:
	//
	// open12gddn2kn1i47v10wRJNkMFx00
	CorpId     *string                                                    `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	Department *HotelOrderInfoQueryResponseBodyModuleBookerInfoDepartment `json:"department,omitempty" xml:"department,omitempty" type:"Struct"`
	// example:
	//
	// Tom
	EnName *string `json:"en_name,omitempty" xml:"en_name,omitempty"`
	// example:
	//
	// 1001
	JobNo *string `json:"job_no,omitempty" xml:"job_no,omitempty"`
	// example:
	//
	// true
	NeedApply *bool   `json:"need_apply,omitempty" xml:"need_apply,omitempty"`
	RealName  *string `json:"real_name,omitempty" xml:"real_name,omitempty"`
	// example:
	//
	// 1430378
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s HotelOrderInfoQueryResponseBodyModuleBookerInfo) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderInfoQueryResponseBodyModuleBookerInfo) GoString() string {
	return s.String()
}

func (s *HotelOrderInfoQueryResponseBodyModuleBookerInfo) GetBookerRole() *string {
	return s.BookerRole
}

func (s *HotelOrderInfoQueryResponseBodyModuleBookerInfo) GetContactEmail() *string {
	return s.ContactEmail
}

func (s *HotelOrderInfoQueryResponseBodyModuleBookerInfo) GetContactPhone() *string {
	return s.ContactPhone
}

func (s *HotelOrderInfoQueryResponseBodyModuleBookerInfo) GetCorpId() *string {
	return s.CorpId
}

func (s *HotelOrderInfoQueryResponseBodyModuleBookerInfo) GetDepartment() *HotelOrderInfoQueryResponseBodyModuleBookerInfoDepartment {
	return s.Department
}

func (s *HotelOrderInfoQueryResponseBodyModuleBookerInfo) GetEnName() *string {
	return s.EnName
}

func (s *HotelOrderInfoQueryResponseBodyModuleBookerInfo) GetJobNo() *string {
	return s.JobNo
}

func (s *HotelOrderInfoQueryResponseBodyModuleBookerInfo) GetNeedApply() *bool {
	return s.NeedApply
}

func (s *HotelOrderInfoQueryResponseBodyModuleBookerInfo) GetRealName() *string {
	return s.RealName
}

func (s *HotelOrderInfoQueryResponseBodyModuleBookerInfo) GetUserId() *string {
	return s.UserId
}

func (s *HotelOrderInfoQueryResponseBodyModuleBookerInfo) SetBookerRole(v string) *HotelOrderInfoQueryResponseBodyModuleBookerInfo {
	s.BookerRole = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleBookerInfo) SetContactEmail(v string) *HotelOrderInfoQueryResponseBodyModuleBookerInfo {
	s.ContactEmail = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleBookerInfo) SetContactPhone(v string) *HotelOrderInfoQueryResponseBodyModuleBookerInfo {
	s.ContactPhone = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleBookerInfo) SetCorpId(v string) *HotelOrderInfoQueryResponseBodyModuleBookerInfo {
	s.CorpId = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleBookerInfo) SetDepartment(v *HotelOrderInfoQueryResponseBodyModuleBookerInfoDepartment) *HotelOrderInfoQueryResponseBodyModuleBookerInfo {
	s.Department = v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleBookerInfo) SetEnName(v string) *HotelOrderInfoQueryResponseBodyModuleBookerInfo {
	s.EnName = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleBookerInfo) SetJobNo(v string) *HotelOrderInfoQueryResponseBodyModuleBookerInfo {
	s.JobNo = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleBookerInfo) SetNeedApply(v bool) *HotelOrderInfoQueryResponseBodyModuleBookerInfo {
	s.NeedApply = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleBookerInfo) SetRealName(v string) *HotelOrderInfoQueryResponseBodyModuleBookerInfo {
	s.RealName = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleBookerInfo) SetUserId(v string) *HotelOrderInfoQueryResponseBodyModuleBookerInfo {
	s.UserId = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleBookerInfo) Validate() error {
	return dara.Validate(s)
}

type HotelOrderInfoQueryResponseBodyModuleBookerInfoDepartment struct {
	// example:
	//
	// 0rCUI20hjOsk0sTwlu
	CascadeDeptMask *string `json:"cascade_dept_mask,omitempty" xml:"cascade_dept_mask,omitempty"`
	CascadeDeptName *string `json:"cascade_dept_name,omitempty" xml:"cascade_dept_name,omitempty"`
	// example:
	//
	// 35
	DepartId   *string `json:"depart_id,omitempty" xml:"depart_id,omitempty"`
	DepartName *string `json:"depart_name,omitempty" xml:"depart_name,omitempty"`
	// example:
	//
	// 330000001815
	OutDepartId *string `json:"out_depart_id,omitempty" xml:"out_depart_id,omitempty"`
}

func (s HotelOrderInfoQueryResponseBodyModuleBookerInfoDepartment) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderInfoQueryResponseBodyModuleBookerInfoDepartment) GoString() string {
	return s.String()
}

func (s *HotelOrderInfoQueryResponseBodyModuleBookerInfoDepartment) GetCascadeDeptMask() *string {
	return s.CascadeDeptMask
}

func (s *HotelOrderInfoQueryResponseBodyModuleBookerInfoDepartment) GetCascadeDeptName() *string {
	return s.CascadeDeptName
}

func (s *HotelOrderInfoQueryResponseBodyModuleBookerInfoDepartment) GetDepartId() *string {
	return s.DepartId
}

func (s *HotelOrderInfoQueryResponseBodyModuleBookerInfoDepartment) GetDepartName() *string {
	return s.DepartName
}

func (s *HotelOrderInfoQueryResponseBodyModuleBookerInfoDepartment) GetOutDepartId() *string {
	return s.OutDepartId
}

func (s *HotelOrderInfoQueryResponseBodyModuleBookerInfoDepartment) SetCascadeDeptMask(v string) *HotelOrderInfoQueryResponseBodyModuleBookerInfoDepartment {
	s.CascadeDeptMask = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleBookerInfoDepartment) SetCascadeDeptName(v string) *HotelOrderInfoQueryResponseBodyModuleBookerInfoDepartment {
	s.CascadeDeptName = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleBookerInfoDepartment) SetDepartId(v string) *HotelOrderInfoQueryResponseBodyModuleBookerInfoDepartment {
	s.DepartId = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleBookerInfoDepartment) SetDepartName(v string) *HotelOrderInfoQueryResponseBodyModuleBookerInfoDepartment {
	s.DepartName = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleBookerInfoDepartment) SetOutDepartId(v string) *HotelOrderInfoQueryResponseBodyModuleBookerInfoDepartment {
	s.OutDepartId = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleBookerInfoDepartment) Validate() error {
	return dara.Validate(s)
}

type HotelOrderInfoQueryResponseBodyModuleHotelInfo struct {
	// example:
	//
	// 330100
	CityCode *string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	CityName *string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	// example:
	//
	// 1
	CountryCode *string `json:"country_code,omitempty" xml:"country_code,omitempty"`
	CountryName *string `json:"country_name,omitempty" xml:"country_name,omitempty"`
	// example:
	//
	// 330183
	DistrictCode *string `json:"district_code,omitempty" xml:"district_code,omitempty"`
	DistrictName *string `json:"district_name,omitempty" xml:"district_name,omitempty"`
	HotelAddress *string `json:"hotel_address,omitempty" xml:"hotel_address,omitempty"`
	// example:
	//
	// 3
	HotelBrandCode *string `json:"hotel_brand_code,omitempty" xml:"hotel_brand_code,omitempty"`
	HotelBrandName *string `json:"hotel_brand_name,omitempty" xml:"hotel_brand_name,omitempty"`
	// example:
	//
	// huazhu
	HotelGroup *string `json:"hotel_group,omitempty" xml:"hotel_group,omitempty"`
	// example:
	//
	// 55335212
	HotelId   *string `json:"hotel_id,omitempty" xml:"hotel_id,omitempty"`
	HotelName *string `json:"hotel_name,omitempty" xml:"hotel_name,omitempty"`
	// example:
	//
	// Test Hotel Name
	HotelNameEn *string `json:"hotel_name_en,omitempty" xml:"hotel_name_en,omitempty"`
	// example:
	//
	// 5
	Star *string `json:"star,omitempty" xml:"star,omitempty"`
}

func (s HotelOrderInfoQueryResponseBodyModuleHotelInfo) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderInfoQueryResponseBodyModuleHotelInfo) GoString() string {
	return s.String()
}

func (s *HotelOrderInfoQueryResponseBodyModuleHotelInfo) GetCityCode() *string {
	return s.CityCode
}

func (s *HotelOrderInfoQueryResponseBodyModuleHotelInfo) GetCityName() *string {
	return s.CityName
}

func (s *HotelOrderInfoQueryResponseBodyModuleHotelInfo) GetCountryCode() *string {
	return s.CountryCode
}

func (s *HotelOrderInfoQueryResponseBodyModuleHotelInfo) GetCountryName() *string {
	return s.CountryName
}

func (s *HotelOrderInfoQueryResponseBodyModuleHotelInfo) GetDistrictCode() *string {
	return s.DistrictCode
}

func (s *HotelOrderInfoQueryResponseBodyModuleHotelInfo) GetDistrictName() *string {
	return s.DistrictName
}

func (s *HotelOrderInfoQueryResponseBodyModuleHotelInfo) GetHotelAddress() *string {
	return s.HotelAddress
}

func (s *HotelOrderInfoQueryResponseBodyModuleHotelInfo) GetHotelBrandCode() *string {
	return s.HotelBrandCode
}

func (s *HotelOrderInfoQueryResponseBodyModuleHotelInfo) GetHotelBrandName() *string {
	return s.HotelBrandName
}

func (s *HotelOrderInfoQueryResponseBodyModuleHotelInfo) GetHotelGroup() *string {
	return s.HotelGroup
}

func (s *HotelOrderInfoQueryResponseBodyModuleHotelInfo) GetHotelId() *string {
	return s.HotelId
}

func (s *HotelOrderInfoQueryResponseBodyModuleHotelInfo) GetHotelName() *string {
	return s.HotelName
}

func (s *HotelOrderInfoQueryResponseBodyModuleHotelInfo) GetHotelNameEn() *string {
	return s.HotelNameEn
}

func (s *HotelOrderInfoQueryResponseBodyModuleHotelInfo) GetStar() *string {
	return s.Star
}

func (s *HotelOrderInfoQueryResponseBodyModuleHotelInfo) SetCityCode(v string) *HotelOrderInfoQueryResponseBodyModuleHotelInfo {
	s.CityCode = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleHotelInfo) SetCityName(v string) *HotelOrderInfoQueryResponseBodyModuleHotelInfo {
	s.CityName = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleHotelInfo) SetCountryCode(v string) *HotelOrderInfoQueryResponseBodyModuleHotelInfo {
	s.CountryCode = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleHotelInfo) SetCountryName(v string) *HotelOrderInfoQueryResponseBodyModuleHotelInfo {
	s.CountryName = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleHotelInfo) SetDistrictCode(v string) *HotelOrderInfoQueryResponseBodyModuleHotelInfo {
	s.DistrictCode = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleHotelInfo) SetDistrictName(v string) *HotelOrderInfoQueryResponseBodyModuleHotelInfo {
	s.DistrictName = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleHotelInfo) SetHotelAddress(v string) *HotelOrderInfoQueryResponseBodyModuleHotelInfo {
	s.HotelAddress = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleHotelInfo) SetHotelBrandCode(v string) *HotelOrderInfoQueryResponseBodyModuleHotelInfo {
	s.HotelBrandCode = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleHotelInfo) SetHotelBrandName(v string) *HotelOrderInfoQueryResponseBodyModuleHotelInfo {
	s.HotelBrandName = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleHotelInfo) SetHotelGroup(v string) *HotelOrderInfoQueryResponseBodyModuleHotelInfo {
	s.HotelGroup = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleHotelInfo) SetHotelId(v string) *HotelOrderInfoQueryResponseBodyModuleHotelInfo {
	s.HotelId = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleHotelInfo) SetHotelName(v string) *HotelOrderInfoQueryResponseBodyModuleHotelInfo {
	s.HotelName = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleHotelInfo) SetHotelNameEn(v string) *HotelOrderInfoQueryResponseBodyModuleHotelInfo {
	s.HotelNameEn = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleHotelInfo) SetStar(v string) *HotelOrderInfoQueryResponseBodyModuleHotelInfo {
	s.Star = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleHotelInfo) Validate() error {
	return dara.Validate(s)
}

type HotelOrderInfoQueryResponseBodyModuleHotelOrderFeeInfo struct {
	// example:
	//
	// 35000
	OrderAmount *int64 `json:"order_amount,omitempty" xml:"order_amount,omitempty"`
	// example:
	//
	// 0
	OtherFee *int64 `json:"other_fee,omitempty" xml:"other_fee,omitempty"`
	// example:
	//
	// 35000
	PayAmount *int64 `json:"pay_amount,omitempty" xml:"pay_amount,omitempty"`
	// example:
	//
	// 5000
	PromotionAmount *int64 `json:"promotion_amount,omitempty" xml:"promotion_amount,omitempty"`
	// example:
	//
	// 35000
	TotalRoomAmount *int64 `json:"total_room_amount,omitempty" xml:"total_room_amount,omitempty"`
}

func (s HotelOrderInfoQueryResponseBodyModuleHotelOrderFeeInfo) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderInfoQueryResponseBodyModuleHotelOrderFeeInfo) GoString() string {
	return s.String()
}

func (s *HotelOrderInfoQueryResponseBodyModuleHotelOrderFeeInfo) GetOrderAmount() *int64 {
	return s.OrderAmount
}

func (s *HotelOrderInfoQueryResponseBodyModuleHotelOrderFeeInfo) GetOtherFee() *int64 {
	return s.OtherFee
}

func (s *HotelOrderInfoQueryResponseBodyModuleHotelOrderFeeInfo) GetPayAmount() *int64 {
	return s.PayAmount
}

func (s *HotelOrderInfoQueryResponseBodyModuleHotelOrderFeeInfo) GetPromotionAmount() *int64 {
	return s.PromotionAmount
}

func (s *HotelOrderInfoQueryResponseBodyModuleHotelOrderFeeInfo) GetTotalRoomAmount() *int64 {
	return s.TotalRoomAmount
}

func (s *HotelOrderInfoQueryResponseBodyModuleHotelOrderFeeInfo) SetOrderAmount(v int64) *HotelOrderInfoQueryResponseBodyModuleHotelOrderFeeInfo {
	s.OrderAmount = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleHotelOrderFeeInfo) SetOtherFee(v int64) *HotelOrderInfoQueryResponseBodyModuleHotelOrderFeeInfo {
	s.OtherFee = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleHotelOrderFeeInfo) SetPayAmount(v int64) *HotelOrderInfoQueryResponseBodyModuleHotelOrderFeeInfo {
	s.PayAmount = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleHotelOrderFeeInfo) SetPromotionAmount(v int64) *HotelOrderInfoQueryResponseBodyModuleHotelOrderFeeInfo {
	s.PromotionAmount = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleHotelOrderFeeInfo) SetTotalRoomAmount(v int64) *HotelOrderInfoQueryResponseBodyModuleHotelOrderFeeInfo {
	s.TotalRoomAmount = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleHotelOrderFeeInfo) Validate() error {
	return dara.Validate(s)
}

type HotelOrderInfoQueryResponseBodyModuleHotelOrderRefundInfo struct {
	// example:
	//
	// 1000
	CancelFine *int64 `json:"cancel_fine,omitempty" xml:"cancel_fine,omitempty"`
	// example:
	//
	// 1000000002578096
	RefundApplyId *int64 `json:"refund_apply_id,omitempty" xml:"refund_apply_id,omitempty"`
	// example:
	//
	// 1721702353700
	RefundEndTime *int64 `json:"refund_end_time,omitempty" xml:"refund_end_time,omitempty"`
	// example:
	//
	// 34000
	RefundPrice  *int64  `json:"refund_price,omitempty" xml:"refund_price,omitempty"`
	RefundReason *string `json:"refund_reason,omitempty" xml:"refund_reason,omitempty"`
	// example:
	//
	// 1721702353700
	RefundStartTime *int64 `json:"refund_start_time,omitempty" xml:"refund_start_time,omitempty"`
	// example:
	//
	// 5
	RefundType *int32 `json:"refund_type,omitempty" xml:"refund_type,omitempty"`
}

func (s HotelOrderInfoQueryResponseBodyModuleHotelOrderRefundInfo) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderInfoQueryResponseBodyModuleHotelOrderRefundInfo) GoString() string {
	return s.String()
}

func (s *HotelOrderInfoQueryResponseBodyModuleHotelOrderRefundInfo) GetCancelFine() *int64 {
	return s.CancelFine
}

func (s *HotelOrderInfoQueryResponseBodyModuleHotelOrderRefundInfo) GetRefundApplyId() *int64 {
	return s.RefundApplyId
}

func (s *HotelOrderInfoQueryResponseBodyModuleHotelOrderRefundInfo) GetRefundEndTime() *int64 {
	return s.RefundEndTime
}

func (s *HotelOrderInfoQueryResponseBodyModuleHotelOrderRefundInfo) GetRefundPrice() *int64 {
	return s.RefundPrice
}

func (s *HotelOrderInfoQueryResponseBodyModuleHotelOrderRefundInfo) GetRefundReason() *string {
	return s.RefundReason
}

func (s *HotelOrderInfoQueryResponseBodyModuleHotelOrderRefundInfo) GetRefundStartTime() *int64 {
	return s.RefundStartTime
}

func (s *HotelOrderInfoQueryResponseBodyModuleHotelOrderRefundInfo) GetRefundType() *int32 {
	return s.RefundType
}

func (s *HotelOrderInfoQueryResponseBodyModuleHotelOrderRefundInfo) SetCancelFine(v int64) *HotelOrderInfoQueryResponseBodyModuleHotelOrderRefundInfo {
	s.CancelFine = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleHotelOrderRefundInfo) SetRefundApplyId(v int64) *HotelOrderInfoQueryResponseBodyModuleHotelOrderRefundInfo {
	s.RefundApplyId = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleHotelOrderRefundInfo) SetRefundEndTime(v int64) *HotelOrderInfoQueryResponseBodyModuleHotelOrderRefundInfo {
	s.RefundEndTime = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleHotelOrderRefundInfo) SetRefundPrice(v int64) *HotelOrderInfoQueryResponseBodyModuleHotelOrderRefundInfo {
	s.RefundPrice = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleHotelOrderRefundInfo) SetRefundReason(v string) *HotelOrderInfoQueryResponseBodyModuleHotelOrderRefundInfo {
	s.RefundReason = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleHotelOrderRefundInfo) SetRefundStartTime(v int64) *HotelOrderInfoQueryResponseBodyModuleHotelOrderRefundInfo {
	s.RefundStartTime = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleHotelOrderRefundInfo) SetRefundType(v int32) *HotelOrderInfoQueryResponseBodyModuleHotelOrderRefundInfo {
	s.RefundType = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleHotelOrderRefundInfo) Validate() error {
	return dara.Validate(s)
}

type HotelOrderInfoQueryResponseBodyModuleRoomTraverInfo struct {
	// example:
	//
	// 1
	LiveRoomNo   *string                                                           `json:"live_room_no,omitempty" xml:"live_room_no,omitempty"`
	RoomTypeName *string                                                           `json:"room_type_name,omitempty" xml:"room_type_name,omitempty"`
	TraverInfos  []*HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfos `json:"traver_infos,omitempty" xml:"traver_infos,omitempty" type:"Repeated"`
}

func (s HotelOrderInfoQueryResponseBodyModuleRoomTraverInfo) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderInfoQueryResponseBodyModuleRoomTraverInfo) GoString() string {
	return s.String()
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfo) GetLiveRoomNo() *string {
	return s.LiveRoomNo
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfo) GetRoomTypeName() *string {
	return s.RoomTypeName
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfo) GetTraverInfos() []*HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfos {
	return s.TraverInfos
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfo) SetLiveRoomNo(v string) *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfo {
	s.LiveRoomNo = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfo) SetRoomTypeName(v string) *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfo {
	s.RoomTypeName = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfo) SetTraverInfos(v []*HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfos) *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfo {
	s.TraverInfos = v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfo) Validate() error {
	return dara.Validate(s)
}

type HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfos struct {
	ApplyInfo *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosApplyInfo `json:"apply_info,omitempty" xml:"apply_info,omitempty" type:"Struct"`
	// example:
	//
	// 342229200801010023
	CertNo *string `json:"cert_no,omitempty" xml:"cert_no,omitempty"`
	// example:
	//
	// 0
	CertType   *int32                                                                    `json:"cert_type,omitempty" xml:"cert_type,omitempty"`
	Department *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosDepartment `json:"department,omitempty" xml:"department,omitempty" type:"Struct"`
	// example:
	//
	// 1001
	JobNo *string `json:"job_no,omitempty" xml:"job_no,omitempty"`
	// example:
	//
	// 13311112222
	Telephone *string `json:"telephone,omitempty" xml:"telephone,omitempty"`
	// example:
	//
	// 1430378
	TravelerId   *string `json:"traveler_id,omitempty" xml:"traveler_id,omitempty"`
	TravelerName *string `json:"traveler_name,omitempty" xml:"traveler_name,omitempty"`
	// example:
	//
	// 1
	TravelerType   *int32                                                                        `json:"traveler_type,omitempty" xml:"traveler_type,omitempty"`
	TripCostCenter *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosTripCostCenter `json:"trip_cost_center,omitempty" xml:"trip_cost_center,omitempty" type:"Struct"`
	// example:
	//
	// 0
	UserType *int32 `json:"user_type,omitempty" xml:"user_type,omitempty"`
}

func (s HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfos) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfos) GoString() string {
	return s.String()
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfos) GetApplyInfo() *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosApplyInfo {
	return s.ApplyInfo
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfos) GetCertNo() *string {
	return s.CertNo
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfos) GetCertType() *int32 {
	return s.CertType
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfos) GetDepartment() *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosDepartment {
	return s.Department
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfos) GetJobNo() *string {
	return s.JobNo
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfos) GetTelephone() *string {
	return s.Telephone
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfos) GetTravelerId() *string {
	return s.TravelerId
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfos) GetTravelerName() *string {
	return s.TravelerName
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfos) GetTravelerType() *int32 {
	return s.TravelerType
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfos) GetTripCostCenter() *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosTripCostCenter {
	return s.TripCostCenter
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfos) GetUserType() *int32 {
	return s.UserType
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfos) SetApplyInfo(v *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosApplyInfo) *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfos {
	s.ApplyInfo = v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfos) SetCertNo(v string) *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfos {
	s.CertNo = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfos) SetCertType(v int32) *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfos {
	s.CertType = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfos) SetDepartment(v *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosDepartment) *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfos {
	s.Department = v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfos) SetJobNo(v string) *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfos {
	s.JobNo = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfos) SetTelephone(v string) *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfos {
	s.Telephone = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfos) SetTravelerId(v string) *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfos {
	s.TravelerId = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfos) SetTravelerName(v string) *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfos {
	s.TravelerName = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfos) SetTravelerType(v int32) *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfos {
	s.TravelerType = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfos) SetTripCostCenter(v *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosTripCostCenter) *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfos {
	s.TripCostCenter = v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfos) SetUserType(v int32) *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfos {
	s.UserType = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfos) Validate() error {
	return dara.Validate(s)
}

type HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosApplyInfo struct {
	// example:
	//
	// 1001
	ApplyBusinessId   *string `json:"apply_business_id,omitempty" xml:"apply_business_id,omitempty"`
	ApplyBusinessName *string `json:"apply_business_name,omitempty" xml:"apply_business_name,omitempty"`
	// example:
	//
	// 1424031910085891196
	ApplyId     *string                                                                               `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	ExceedApply []*HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosApplyInfoExceedApply `json:"exceed_apply,omitempty" xml:"exceed_apply,omitempty" type:"Repeated"`
	// example:
	//
	// ef5e74dc1f1640b08858fb043f64e477-8
	ItineraryNo *string `json:"itinerary_no,omitempty" xml:"itinerary_no,omitempty"`
}

func (s HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosApplyInfo) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosApplyInfo) GoString() string {
	return s.String()
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosApplyInfo) GetApplyBusinessId() *string {
	return s.ApplyBusinessId
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosApplyInfo) GetApplyBusinessName() *string {
	return s.ApplyBusinessName
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosApplyInfo) GetApplyId() *string {
	return s.ApplyId
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosApplyInfo) GetExceedApply() []*HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosApplyInfoExceedApply {
	return s.ExceedApply
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosApplyInfo) GetItineraryNo() *string {
	return s.ItineraryNo
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosApplyInfo) SetApplyBusinessId(v string) *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosApplyInfo {
	s.ApplyBusinessId = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosApplyInfo) SetApplyBusinessName(v string) *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosApplyInfo {
	s.ApplyBusinessName = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosApplyInfo) SetApplyId(v string) *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosApplyInfo {
	s.ApplyId = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosApplyInfo) SetExceedApply(v []*HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosApplyInfoExceedApply) *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosApplyInfo {
	s.ExceedApply = v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosApplyInfo) SetItineraryNo(v string) *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosApplyInfo {
	s.ItineraryNo = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosApplyInfo) Validate() error {
	return dara.Validate(s)
}

type HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosApplyInfoExceedApply struct {
	ExceedReason *string `json:"exceed_reason,omitempty" xml:"exceed_reason,omitempty"`
	// example:
	//
	// 16
	ExceedType *int32 `json:"exceed_type,omitempty" xml:"exceed_type,omitempty"`
	// example:
	//
	// 3321
	FlowNo *int64 `json:"flow_no,omitempty" xml:"flow_no,omitempty"`
	// example:
	//
	// 60853
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
}

func (s HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosApplyInfoExceedApply) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosApplyInfoExceedApply) GoString() string {
	return s.String()
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosApplyInfoExceedApply) GetExceedReason() *string {
	return s.ExceedReason
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosApplyInfoExceedApply) GetExceedType() *int32 {
	return s.ExceedType
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosApplyInfoExceedApply) GetFlowNo() *int64 {
	return s.FlowNo
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosApplyInfoExceedApply) GetId() *int64 {
	return s.Id
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosApplyInfoExceedApply) SetExceedReason(v string) *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosApplyInfoExceedApply {
	s.ExceedReason = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosApplyInfoExceedApply) SetExceedType(v int32) *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosApplyInfoExceedApply {
	s.ExceedType = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosApplyInfoExceedApply) SetFlowNo(v int64) *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosApplyInfoExceedApply {
	s.FlowNo = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosApplyInfoExceedApply) SetId(v int64) *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosApplyInfoExceedApply {
	s.Id = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosApplyInfoExceedApply) Validate() error {
	return dara.Validate(s)
}

type HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosDepartment struct {
	// example:
	//
	// 0rCUI20hjOsk0sTwlu
	CascadeDeptMask *string `json:"cascade_dept_mask,omitempty" xml:"cascade_dept_mask,omitempty"`
	CascadeDeptName *string `json:"cascade_dept_name,omitempty" xml:"cascade_dept_name,omitempty"`
	// example:
	//
	// 35
	DepartId   *string `json:"depart_id,omitempty" xml:"depart_id,omitempty"`
	DepartName *string `json:"depart_name,omitempty" xml:"depart_name,omitempty"`
	// example:
	//
	// 330000001815
	OutDepartId *string `json:"out_depart_id,omitempty" xml:"out_depart_id,omitempty"`
}

func (s HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosDepartment) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosDepartment) GoString() string {
	return s.String()
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosDepartment) GetCascadeDeptMask() *string {
	return s.CascadeDeptMask
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosDepartment) GetCascadeDeptName() *string {
	return s.CascadeDeptName
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosDepartment) GetDepartId() *string {
	return s.DepartId
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosDepartment) GetDepartName() *string {
	return s.DepartName
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosDepartment) GetOutDepartId() *string {
	return s.OutDepartId
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosDepartment) SetCascadeDeptMask(v string) *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosDepartment {
	s.CascadeDeptMask = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosDepartment) SetCascadeDeptName(v string) *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosDepartment {
	s.CascadeDeptName = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosDepartment) SetDepartId(v string) *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosDepartment {
	s.DepartId = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosDepartment) SetDepartName(v string) *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosDepartment {
	s.DepartName = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosDepartment) SetOutDepartId(v string) *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosDepartment {
	s.OutDepartId = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosDepartment) Validate() error {
	return dara.Validate(s)
}

type HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosTripCostCenter struct {
	// example:
	//
	// 323431
	CostCenterCode *string `json:"cost_center_code,omitempty" xml:"cost_center_code,omitempty"`
	// example:
	//
	// 2312
	CostCenterId   *string `json:"cost_center_id,omitempty" xml:"cost_center_id,omitempty"`
	CostCenterName *string `json:"cost_center_name,omitempty" xml:"cost_center_name,omitempty"`
	// example:
	//
	// {"extraKey":"extraVal"}
	ExternalExtField *string `json:"external_ext_field,omitempty" xml:"external_ext_field,omitempty"`
	// example:
	//
	// 1
	FeeType *int32 `json:"fee_type,omitempty" xml:"fee_type,omitempty"`
	// example:
	//
	// 123332
	InvoiceId    *int64  `json:"invoice_id,omitempty" xml:"invoice_id,omitempty"`
	InvoiceTitle *string `json:"invoice_title,omitempty" xml:"invoice_title,omitempty"`
	// example:
	//
	// 1002
	ProjectCode  *string `json:"project_code,omitempty" xml:"project_code,omitempty"`
	ProjectTitle *string `json:"project_title,omitempty" xml:"project_title,omitempty"`
}

func (s HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosTripCostCenter) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosTripCostCenter) GoString() string {
	return s.String()
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosTripCostCenter) GetCostCenterCode() *string {
	return s.CostCenterCode
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosTripCostCenter) GetCostCenterId() *string {
	return s.CostCenterId
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosTripCostCenter) GetCostCenterName() *string {
	return s.CostCenterName
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosTripCostCenter) GetExternalExtField() *string {
	return s.ExternalExtField
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosTripCostCenter) GetFeeType() *int32 {
	return s.FeeType
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosTripCostCenter) GetInvoiceId() *int64 {
	return s.InvoiceId
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosTripCostCenter) GetInvoiceTitle() *string {
	return s.InvoiceTitle
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosTripCostCenter) GetProjectCode() *string {
	return s.ProjectCode
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosTripCostCenter) GetProjectTitle() *string {
	return s.ProjectTitle
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosTripCostCenter) SetCostCenterCode(v string) *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosTripCostCenter {
	s.CostCenterCode = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosTripCostCenter) SetCostCenterId(v string) *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosTripCostCenter {
	s.CostCenterId = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosTripCostCenter) SetCostCenterName(v string) *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosTripCostCenter {
	s.CostCenterName = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosTripCostCenter) SetExternalExtField(v string) *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosTripCostCenter {
	s.ExternalExtField = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosTripCostCenter) SetFeeType(v int32) *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosTripCostCenter {
	s.FeeType = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosTripCostCenter) SetInvoiceId(v int64) *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosTripCostCenter {
	s.InvoiceId = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosTripCostCenter) SetInvoiceTitle(v string) *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosTripCostCenter {
	s.InvoiceTitle = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosTripCostCenter) SetProjectCode(v string) *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosTripCostCenter {
	s.ProjectCode = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosTripCostCenter) SetProjectTitle(v string) *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosTripCostCenter {
	s.ProjectTitle = &v
	return s
}

func (s *HotelOrderInfoQueryResponseBodyModuleRoomTraverInfoTraverInfosTripCostCenter) Validate() error {
	return dara.Validate(s)
}
