// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCarOrderQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CarOrderQueryResponseBody
	GetCode() *string
	SetMessage(v string) *CarOrderQueryResponseBody
	GetMessage() *string
	SetModule(v *CarOrderQueryResponseBodyModule) *CarOrderQueryResponseBody
	GetModule() *CarOrderQueryResponseBodyModule
	SetRequestId(v string) *CarOrderQueryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CarOrderQueryResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *CarOrderQueryResponseBody
	GetTraceId() *string
}

type CarOrderQueryResponseBody struct {
	// example:
	//
	// SUCCESS
	Code    *string                          `json:"code,omitempty" xml:"code,omitempty"`
	Message *string                          `json:"message,omitempty" xml:"message,omitempty"`
	Module  *CarOrderQueryResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// requestId
	//
	// example:
	//
	// A5009956-1077-52FB-B520-EA8C7E91D722
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// traceId
	//
	// example:
	//
	// 21041ce316577904808056433edbb2
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s CarOrderQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CarOrderQueryResponseBody) GoString() string {
	return s.String()
}

func (s *CarOrderQueryResponseBody) GetCode() *string {
	return s.Code
}

func (s *CarOrderQueryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CarOrderQueryResponseBody) GetModule() *CarOrderQueryResponseBodyModule {
	return s.Module
}

func (s *CarOrderQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CarOrderQueryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CarOrderQueryResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *CarOrderQueryResponseBody) SetCode(v string) *CarOrderQueryResponseBody {
	s.Code = &v
	return s
}

func (s *CarOrderQueryResponseBody) SetMessage(v string) *CarOrderQueryResponseBody {
	s.Message = &v
	return s
}

func (s *CarOrderQueryResponseBody) SetModule(v *CarOrderQueryResponseBodyModule) *CarOrderQueryResponseBody {
	s.Module = v
	return s
}

func (s *CarOrderQueryResponseBody) SetRequestId(v string) *CarOrderQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CarOrderQueryResponseBody) SetSuccess(v bool) *CarOrderQueryResponseBody {
	s.Success = &v
	return s
}

func (s *CarOrderQueryResponseBody) SetTraceId(v string) *CarOrderQueryResponseBody {
	s.TraceId = &v
	return s
}

func (s *CarOrderQueryResponseBody) Validate() error {
	return dara.Validate(s)
}

type CarOrderQueryResponseBodyModule struct {
	CarInfo       *CarOrderQueryResponseBodyModuleCarInfo         `json:"car_info,omitempty" xml:"car_info,omitempty" type:"Struct"`
	InvoiceInfo   *CarOrderQueryResponseBodyModuleInvoiceInfo     `json:"invoice_info,omitempty" xml:"invoice_info,omitempty" type:"Struct"`
	OrderBaseInfo *CarOrderQueryResponseBodyModuleOrderBaseInfo   `json:"order_base_info,omitempty" xml:"order_base_info,omitempty" type:"Struct"`
	PassengerList []*CarOrderQueryResponseBodyModulePassengerList `json:"passenger_list,omitempty" xml:"passenger_list,omitempty" type:"Repeated"`
	PriceInfoList []*CarOrderQueryResponseBodyModulePriceInfoList `json:"price_info_list,omitempty" xml:"price_info_list,omitempty" type:"Repeated"`
}

func (s CarOrderQueryResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s CarOrderQueryResponseBodyModule) GoString() string {
	return s.String()
}

func (s *CarOrderQueryResponseBodyModule) GetCarInfo() *CarOrderQueryResponseBodyModuleCarInfo {
	return s.CarInfo
}

func (s *CarOrderQueryResponseBodyModule) GetInvoiceInfo() *CarOrderQueryResponseBodyModuleInvoiceInfo {
	return s.InvoiceInfo
}

func (s *CarOrderQueryResponseBodyModule) GetOrderBaseInfo() *CarOrderQueryResponseBodyModuleOrderBaseInfo {
	return s.OrderBaseInfo
}

func (s *CarOrderQueryResponseBodyModule) GetPassengerList() []*CarOrderQueryResponseBodyModulePassengerList {
	return s.PassengerList
}

func (s *CarOrderQueryResponseBodyModule) GetPriceInfoList() []*CarOrderQueryResponseBodyModulePriceInfoList {
	return s.PriceInfoList
}

func (s *CarOrderQueryResponseBodyModule) SetCarInfo(v *CarOrderQueryResponseBodyModuleCarInfo) *CarOrderQueryResponseBodyModule {
	s.CarInfo = v
	return s
}

func (s *CarOrderQueryResponseBodyModule) SetInvoiceInfo(v *CarOrderQueryResponseBodyModuleInvoiceInfo) *CarOrderQueryResponseBodyModule {
	s.InvoiceInfo = v
	return s
}

func (s *CarOrderQueryResponseBodyModule) SetOrderBaseInfo(v *CarOrderQueryResponseBodyModuleOrderBaseInfo) *CarOrderQueryResponseBodyModule {
	s.OrderBaseInfo = v
	return s
}

func (s *CarOrderQueryResponseBodyModule) SetPassengerList(v []*CarOrderQueryResponseBodyModulePassengerList) *CarOrderQueryResponseBodyModule {
	s.PassengerList = v
	return s
}

func (s *CarOrderQueryResponseBodyModule) SetPriceInfoList(v []*CarOrderQueryResponseBodyModulePriceInfoList) *CarOrderQueryResponseBodyModule {
	s.PriceInfoList = v
	return s
}

func (s *CarOrderQueryResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type CarOrderQueryResponseBodyModuleCarInfo struct {
	// example:
	//
	// TRAVEL
	BusinessCategory *string `json:"business_category,omitempty" xml:"business_category,omitempty"`
	// example:
	//
	// 1669274251000
	CancelTime *int64  `json:"cancel_time,omitempty" xml:"cancel_time,omitempty"`
	CarInfo    *string `json:"car_info,omitempty" xml:"car_info,omitempty"`
	// example:
	//
	// 601
	CarLevel   *int32  `json:"car_level,omitempty" xml:"car_level,omitempty"`
	DriverCard *string `json:"driver_card,omitempty" xml:"driver_card,omitempty"`
	// example:
	//
	// 1669274251000
	DriverConfirmTime *int64  `json:"driver_confirm_time,omitempty" xml:"driver_confirm_time,omitempty"`
	DriverName        *string `json:"driver_name,omitempty" xml:"driver_name,omitempty"`
	// example:
	//
	// 2900
	EstimatePrice  *int64  `json:"estimate_price,omitempty" xml:"estimate_price,omitempty"`
	FromAddress    *string `json:"from_address,omitempty" xml:"from_address,omitempty"`
	FromCityAdCode *string `json:"from_city_ad_code,omitempty" xml:"from_city_ad_code,omitempty"`
	FromCityName   *string `json:"from_city_name,omitempty" xml:"from_city_name,omitempty"`
	// example:
	//
	// true
	IsSpecial *bool   `json:"is_special,omitempty" xml:"is_special,omitempty"`
	Memo      *string `json:"memo,omitempty" xml:"memo,omitempty"`
	// example:
	//
	// 1669274251000
	PayTime *int64 `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// example:
	//
	// 1669274251000
	PublishTime        *int64  `json:"publish_time,omitempty" xml:"publish_time,omitempty"`
	RealFromAddress    *string `json:"real_from_address,omitempty" xml:"real_from_address,omitempty"`
	RealFromCityAdCode *string `json:"real_from_city_ad_code,omitempty" xml:"real_from_city_ad_code,omitempty"`
	RealFromCityName   *string `json:"real_from_city_name,omitempty" xml:"real_from_city_name,omitempty"`
	RealToAddress      *string `json:"real_to_address,omitempty" xml:"real_to_address,omitempty"`
	RealToCityAdCode   *string `json:"real_to_city_ad_code,omitempty" xml:"real_to_city_ad_code,omitempty"`
	RealToCityName     *string `json:"real_to_city_name,omitempty" xml:"real_to_city_name,omitempty"`
	// example:
	//
	// 3
	ServiceType *int32 `json:"service_type,omitempty" xml:"service_type,omitempty"`
	// example:
	//
	// v_sp_t_1,v_sp_t_2,v_sp_t_3,v_sp_t_4,v_sp_t_5
	SpecialTypes *string `json:"special_types,omitempty" xml:"special_types,omitempty"`
	// example:
	//
	// 1669274251000
	TakenTime    *int64  `json:"taken_time,omitempty" xml:"taken_time,omitempty"`
	ToAddress    *string `json:"to_address,omitempty" xml:"to_address,omitempty"`
	ToCityAdCode *string `json:"to_city_ad_code,omitempty" xml:"to_city_ad_code,omitempty"`
	ToCityName   *string `json:"to_city_name,omitempty" xml:"to_city_name,omitempty"`
	// example:
	//
	// 12
	TravelDistance *string                                            `json:"travel_distance,omitempty" xml:"travel_distance,omitempty"`
	WayPoints      []*CarOrderQueryResponseBodyModuleCarInfoWayPoints `json:"way_points,omitempty" xml:"way_points,omitempty" type:"Repeated"`
}

func (s CarOrderQueryResponseBodyModuleCarInfo) String() string {
	return dara.Prettify(s)
}

func (s CarOrderQueryResponseBodyModuleCarInfo) GoString() string {
	return s.String()
}

func (s *CarOrderQueryResponseBodyModuleCarInfo) GetBusinessCategory() *string {
	return s.BusinessCategory
}

func (s *CarOrderQueryResponseBodyModuleCarInfo) GetCancelTime() *int64 {
	return s.CancelTime
}

func (s *CarOrderQueryResponseBodyModuleCarInfo) GetCarInfo() *string {
	return s.CarInfo
}

func (s *CarOrderQueryResponseBodyModuleCarInfo) GetCarLevel() *int32 {
	return s.CarLevel
}

func (s *CarOrderQueryResponseBodyModuleCarInfo) GetDriverCard() *string {
	return s.DriverCard
}

func (s *CarOrderQueryResponseBodyModuleCarInfo) GetDriverConfirmTime() *int64 {
	return s.DriverConfirmTime
}

func (s *CarOrderQueryResponseBodyModuleCarInfo) GetDriverName() *string {
	return s.DriverName
}

func (s *CarOrderQueryResponseBodyModuleCarInfo) GetEstimatePrice() *int64 {
	return s.EstimatePrice
}

func (s *CarOrderQueryResponseBodyModuleCarInfo) GetFromAddress() *string {
	return s.FromAddress
}

func (s *CarOrderQueryResponseBodyModuleCarInfo) GetFromCityAdCode() *string {
	return s.FromCityAdCode
}

func (s *CarOrderQueryResponseBodyModuleCarInfo) GetFromCityName() *string {
	return s.FromCityName
}

func (s *CarOrderQueryResponseBodyModuleCarInfo) GetIsSpecial() *bool {
	return s.IsSpecial
}

func (s *CarOrderQueryResponseBodyModuleCarInfo) GetMemo() *string {
	return s.Memo
}

func (s *CarOrderQueryResponseBodyModuleCarInfo) GetPayTime() *int64 {
	return s.PayTime
}

func (s *CarOrderQueryResponseBodyModuleCarInfo) GetPublishTime() *int64 {
	return s.PublishTime
}

func (s *CarOrderQueryResponseBodyModuleCarInfo) GetRealFromAddress() *string {
	return s.RealFromAddress
}

func (s *CarOrderQueryResponseBodyModuleCarInfo) GetRealFromCityAdCode() *string {
	return s.RealFromCityAdCode
}

func (s *CarOrderQueryResponseBodyModuleCarInfo) GetRealFromCityName() *string {
	return s.RealFromCityName
}

func (s *CarOrderQueryResponseBodyModuleCarInfo) GetRealToAddress() *string {
	return s.RealToAddress
}

func (s *CarOrderQueryResponseBodyModuleCarInfo) GetRealToCityAdCode() *string {
	return s.RealToCityAdCode
}

func (s *CarOrderQueryResponseBodyModuleCarInfo) GetRealToCityName() *string {
	return s.RealToCityName
}

func (s *CarOrderQueryResponseBodyModuleCarInfo) GetServiceType() *int32 {
	return s.ServiceType
}

func (s *CarOrderQueryResponseBodyModuleCarInfo) GetSpecialTypes() *string {
	return s.SpecialTypes
}

func (s *CarOrderQueryResponseBodyModuleCarInfo) GetTakenTime() *int64 {
	return s.TakenTime
}

func (s *CarOrderQueryResponseBodyModuleCarInfo) GetToAddress() *string {
	return s.ToAddress
}

func (s *CarOrderQueryResponseBodyModuleCarInfo) GetToCityAdCode() *string {
	return s.ToCityAdCode
}

func (s *CarOrderQueryResponseBodyModuleCarInfo) GetToCityName() *string {
	return s.ToCityName
}

func (s *CarOrderQueryResponseBodyModuleCarInfo) GetTravelDistance() *string {
	return s.TravelDistance
}

func (s *CarOrderQueryResponseBodyModuleCarInfo) GetWayPoints() []*CarOrderQueryResponseBodyModuleCarInfoWayPoints {
	return s.WayPoints
}

func (s *CarOrderQueryResponseBodyModuleCarInfo) SetBusinessCategory(v string) *CarOrderQueryResponseBodyModuleCarInfo {
	s.BusinessCategory = &v
	return s
}

func (s *CarOrderQueryResponseBodyModuleCarInfo) SetCancelTime(v int64) *CarOrderQueryResponseBodyModuleCarInfo {
	s.CancelTime = &v
	return s
}

func (s *CarOrderQueryResponseBodyModuleCarInfo) SetCarInfo(v string) *CarOrderQueryResponseBodyModuleCarInfo {
	s.CarInfo = &v
	return s
}

func (s *CarOrderQueryResponseBodyModuleCarInfo) SetCarLevel(v int32) *CarOrderQueryResponseBodyModuleCarInfo {
	s.CarLevel = &v
	return s
}

func (s *CarOrderQueryResponseBodyModuleCarInfo) SetDriverCard(v string) *CarOrderQueryResponseBodyModuleCarInfo {
	s.DriverCard = &v
	return s
}

func (s *CarOrderQueryResponseBodyModuleCarInfo) SetDriverConfirmTime(v int64) *CarOrderQueryResponseBodyModuleCarInfo {
	s.DriverConfirmTime = &v
	return s
}

func (s *CarOrderQueryResponseBodyModuleCarInfo) SetDriverName(v string) *CarOrderQueryResponseBodyModuleCarInfo {
	s.DriverName = &v
	return s
}

func (s *CarOrderQueryResponseBodyModuleCarInfo) SetEstimatePrice(v int64) *CarOrderQueryResponseBodyModuleCarInfo {
	s.EstimatePrice = &v
	return s
}

func (s *CarOrderQueryResponseBodyModuleCarInfo) SetFromAddress(v string) *CarOrderQueryResponseBodyModuleCarInfo {
	s.FromAddress = &v
	return s
}

func (s *CarOrderQueryResponseBodyModuleCarInfo) SetFromCityAdCode(v string) *CarOrderQueryResponseBodyModuleCarInfo {
	s.FromCityAdCode = &v
	return s
}

func (s *CarOrderQueryResponseBodyModuleCarInfo) SetFromCityName(v string) *CarOrderQueryResponseBodyModuleCarInfo {
	s.FromCityName = &v
	return s
}

func (s *CarOrderQueryResponseBodyModuleCarInfo) SetIsSpecial(v bool) *CarOrderQueryResponseBodyModuleCarInfo {
	s.IsSpecial = &v
	return s
}

func (s *CarOrderQueryResponseBodyModuleCarInfo) SetMemo(v string) *CarOrderQueryResponseBodyModuleCarInfo {
	s.Memo = &v
	return s
}

func (s *CarOrderQueryResponseBodyModuleCarInfo) SetPayTime(v int64) *CarOrderQueryResponseBodyModuleCarInfo {
	s.PayTime = &v
	return s
}

func (s *CarOrderQueryResponseBodyModuleCarInfo) SetPublishTime(v int64) *CarOrderQueryResponseBodyModuleCarInfo {
	s.PublishTime = &v
	return s
}

func (s *CarOrderQueryResponseBodyModuleCarInfo) SetRealFromAddress(v string) *CarOrderQueryResponseBodyModuleCarInfo {
	s.RealFromAddress = &v
	return s
}

func (s *CarOrderQueryResponseBodyModuleCarInfo) SetRealFromCityAdCode(v string) *CarOrderQueryResponseBodyModuleCarInfo {
	s.RealFromCityAdCode = &v
	return s
}

func (s *CarOrderQueryResponseBodyModuleCarInfo) SetRealFromCityName(v string) *CarOrderQueryResponseBodyModuleCarInfo {
	s.RealFromCityName = &v
	return s
}

func (s *CarOrderQueryResponseBodyModuleCarInfo) SetRealToAddress(v string) *CarOrderQueryResponseBodyModuleCarInfo {
	s.RealToAddress = &v
	return s
}

func (s *CarOrderQueryResponseBodyModuleCarInfo) SetRealToCityAdCode(v string) *CarOrderQueryResponseBodyModuleCarInfo {
	s.RealToCityAdCode = &v
	return s
}

func (s *CarOrderQueryResponseBodyModuleCarInfo) SetRealToCityName(v string) *CarOrderQueryResponseBodyModuleCarInfo {
	s.RealToCityName = &v
	return s
}

func (s *CarOrderQueryResponseBodyModuleCarInfo) SetServiceType(v int32) *CarOrderQueryResponseBodyModuleCarInfo {
	s.ServiceType = &v
	return s
}

func (s *CarOrderQueryResponseBodyModuleCarInfo) SetSpecialTypes(v string) *CarOrderQueryResponseBodyModuleCarInfo {
	s.SpecialTypes = &v
	return s
}

func (s *CarOrderQueryResponseBodyModuleCarInfo) SetTakenTime(v int64) *CarOrderQueryResponseBodyModuleCarInfo {
	s.TakenTime = &v
	return s
}

func (s *CarOrderQueryResponseBodyModuleCarInfo) SetToAddress(v string) *CarOrderQueryResponseBodyModuleCarInfo {
	s.ToAddress = &v
	return s
}

func (s *CarOrderQueryResponseBodyModuleCarInfo) SetToCityAdCode(v string) *CarOrderQueryResponseBodyModuleCarInfo {
	s.ToCityAdCode = &v
	return s
}

func (s *CarOrderQueryResponseBodyModuleCarInfo) SetToCityName(v string) *CarOrderQueryResponseBodyModuleCarInfo {
	s.ToCityName = &v
	return s
}

func (s *CarOrderQueryResponseBodyModuleCarInfo) SetTravelDistance(v string) *CarOrderQueryResponseBodyModuleCarInfo {
	s.TravelDistance = &v
	return s
}

func (s *CarOrderQueryResponseBodyModuleCarInfo) SetWayPoints(v []*CarOrderQueryResponseBodyModuleCarInfoWayPoints) *CarOrderQueryResponseBodyModuleCarInfo {
	s.WayPoints = v
	return s
}

func (s *CarOrderQueryResponseBodyModuleCarInfo) Validate() error {
	return dara.Validate(s)
}

type CarOrderQueryResponseBodyModuleCarInfoWayPoints struct {
	Address   *string `json:"address,omitempty" xml:"address,omitempty"`
	Index     *string `json:"index,omitempty" xml:"index,omitempty"`
	Latitude  *string `json:"latitude,omitempty" xml:"latitude,omitempty"`
	Longitude *string `json:"longitude,omitempty" xml:"longitude,omitempty"`
}

func (s CarOrderQueryResponseBodyModuleCarInfoWayPoints) String() string {
	return dara.Prettify(s)
}

func (s CarOrderQueryResponseBodyModuleCarInfoWayPoints) GoString() string {
	return s.String()
}

func (s *CarOrderQueryResponseBodyModuleCarInfoWayPoints) GetAddress() *string {
	return s.Address
}

func (s *CarOrderQueryResponseBodyModuleCarInfoWayPoints) GetIndex() *string {
	return s.Index
}

func (s *CarOrderQueryResponseBodyModuleCarInfoWayPoints) GetLatitude() *string {
	return s.Latitude
}

func (s *CarOrderQueryResponseBodyModuleCarInfoWayPoints) GetLongitude() *string {
	return s.Longitude
}

func (s *CarOrderQueryResponseBodyModuleCarInfoWayPoints) SetAddress(v string) *CarOrderQueryResponseBodyModuleCarInfoWayPoints {
	s.Address = &v
	return s
}

func (s *CarOrderQueryResponseBodyModuleCarInfoWayPoints) SetIndex(v string) *CarOrderQueryResponseBodyModuleCarInfoWayPoints {
	s.Index = &v
	return s
}

func (s *CarOrderQueryResponseBodyModuleCarInfoWayPoints) SetLatitude(v string) *CarOrderQueryResponseBodyModuleCarInfoWayPoints {
	s.Latitude = &v
	return s
}

func (s *CarOrderQueryResponseBodyModuleCarInfoWayPoints) SetLongitude(v string) *CarOrderQueryResponseBodyModuleCarInfoWayPoints {
	s.Longitude = &v
	return s
}

func (s *CarOrderQueryResponseBodyModuleCarInfoWayPoints) Validate() error {
	return dara.Validate(s)
}

type CarOrderQueryResponseBodyModuleInvoiceInfo struct {
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// xxxx
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s CarOrderQueryResponseBodyModuleInvoiceInfo) String() string {
	return dara.Prettify(s)
}

func (s CarOrderQueryResponseBodyModuleInvoiceInfo) GoString() string {
	return s.String()
}

func (s *CarOrderQueryResponseBodyModuleInvoiceInfo) GetId() *int64 {
	return s.Id
}

func (s *CarOrderQueryResponseBodyModuleInvoiceInfo) GetTitle() *string {
	return s.Title
}

func (s *CarOrderQueryResponseBodyModuleInvoiceInfo) SetId(v int64) *CarOrderQueryResponseBodyModuleInvoiceInfo {
	s.Id = &v
	return s
}

func (s *CarOrderQueryResponseBodyModuleInvoiceInfo) SetTitle(v string) *CarOrderQueryResponseBodyModuleInvoiceInfo {
	s.Title = &v
	return s
}

func (s *CarOrderQueryResponseBodyModuleInvoiceInfo) Validate() error {
	return dara.Validate(s)
}

type CarOrderQueryResponseBodyModuleOrderBaseInfo struct {
	// example:
	//
	// xxxxx
	ApplyId    *string `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	BtripCause *string `json:"btrip_cause,omitempty" xml:"btrip_cause,omitempty"`
	BtripTitle *string `json:"btrip_title,omitempty" xml:"btrip_title,omitempty"`
	// example:
	//
	// xxxxxxxx
	CorpId   *string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	CorpName *string `json:"corp_name,omitempty" xml:"corp_name,omitempty"`
	// example:
	//
	// 10101010
	DepartId   *string `json:"depart_id,omitempty" xml:"depart_id,omitempty"`
	DepartName *string `json:"depart_name,omitempty" xml:"depart_name,omitempty"`
	// example:
	//
	// 1669274251000
	GmtCreate *int64 `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// example:
	//
	// 1669274251000
	GmtModified *int64 `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// example:
	//
	// xxxxx
	ItineraryId *string `json:"itinerary_id,omitempty" xml:"itinerary_id,omitempty"`
	// example:
	//
	// 1012000000000000
	OrderId *int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// example:
	//
	// 5
	OrderStatus *int32 `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// example:
	//
	// 100000
	SubOrderId *int64 `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
	// example:
	//
	// thirdpart_1010101010
	ThirdDepartId *string `json:"third_depart_id,omitempty" xml:"third_depart_id,omitempty"`
	// example:
	//
	// xxxxx
	ThirdpartApplyId    *string `json:"thirdpart_apply_id,omitempty" xml:"thirdpart_apply_id,omitempty"`
	ThirdpartBusinessId *string `json:"thirdpart_business_id,omitempty" xml:"thirdpart_business_id,omitempty"`
	// example:
	//
	// xxxxx
	ThirdpartItineraryId *string `json:"thirdpart_itinerary_id,omitempty" xml:"thirdpart_itinerary_id,omitempty"`
	// example:
	//
	// xxxxxxxx
	UserId   *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserName *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s CarOrderQueryResponseBodyModuleOrderBaseInfo) String() string {
	return dara.Prettify(s)
}

func (s CarOrderQueryResponseBodyModuleOrderBaseInfo) GoString() string {
	return s.String()
}

func (s *CarOrderQueryResponseBodyModuleOrderBaseInfo) GetApplyId() *string {
	return s.ApplyId
}

func (s *CarOrderQueryResponseBodyModuleOrderBaseInfo) GetBtripCause() *string {
	return s.BtripCause
}

func (s *CarOrderQueryResponseBodyModuleOrderBaseInfo) GetBtripTitle() *string {
	return s.BtripTitle
}

func (s *CarOrderQueryResponseBodyModuleOrderBaseInfo) GetCorpId() *string {
	return s.CorpId
}

func (s *CarOrderQueryResponseBodyModuleOrderBaseInfo) GetCorpName() *string {
	return s.CorpName
}

func (s *CarOrderQueryResponseBodyModuleOrderBaseInfo) GetDepartId() *string {
	return s.DepartId
}

func (s *CarOrderQueryResponseBodyModuleOrderBaseInfo) GetDepartName() *string {
	return s.DepartName
}

func (s *CarOrderQueryResponseBodyModuleOrderBaseInfo) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *CarOrderQueryResponseBodyModuleOrderBaseInfo) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *CarOrderQueryResponseBodyModuleOrderBaseInfo) GetItineraryId() *string {
	return s.ItineraryId
}

func (s *CarOrderQueryResponseBodyModuleOrderBaseInfo) GetOrderId() *int64 {
	return s.OrderId
}

func (s *CarOrderQueryResponseBodyModuleOrderBaseInfo) GetOrderStatus() *int32 {
	return s.OrderStatus
}

func (s *CarOrderQueryResponseBodyModuleOrderBaseInfo) GetSubOrderId() *int64 {
	return s.SubOrderId
}

func (s *CarOrderQueryResponseBodyModuleOrderBaseInfo) GetThirdDepartId() *string {
	return s.ThirdDepartId
}

func (s *CarOrderQueryResponseBodyModuleOrderBaseInfo) GetThirdpartApplyId() *string {
	return s.ThirdpartApplyId
}

func (s *CarOrderQueryResponseBodyModuleOrderBaseInfo) GetThirdpartBusinessId() *string {
	return s.ThirdpartBusinessId
}

func (s *CarOrderQueryResponseBodyModuleOrderBaseInfo) GetThirdpartItineraryId() *string {
	return s.ThirdpartItineraryId
}

func (s *CarOrderQueryResponseBodyModuleOrderBaseInfo) GetUserId() *string {
	return s.UserId
}

func (s *CarOrderQueryResponseBodyModuleOrderBaseInfo) GetUserName() *string {
	return s.UserName
}

func (s *CarOrderQueryResponseBodyModuleOrderBaseInfo) SetApplyId(v string) *CarOrderQueryResponseBodyModuleOrderBaseInfo {
	s.ApplyId = &v
	return s
}

func (s *CarOrderQueryResponseBodyModuleOrderBaseInfo) SetBtripCause(v string) *CarOrderQueryResponseBodyModuleOrderBaseInfo {
	s.BtripCause = &v
	return s
}

func (s *CarOrderQueryResponseBodyModuleOrderBaseInfo) SetBtripTitle(v string) *CarOrderQueryResponseBodyModuleOrderBaseInfo {
	s.BtripTitle = &v
	return s
}

func (s *CarOrderQueryResponseBodyModuleOrderBaseInfo) SetCorpId(v string) *CarOrderQueryResponseBodyModuleOrderBaseInfo {
	s.CorpId = &v
	return s
}

func (s *CarOrderQueryResponseBodyModuleOrderBaseInfo) SetCorpName(v string) *CarOrderQueryResponseBodyModuleOrderBaseInfo {
	s.CorpName = &v
	return s
}

func (s *CarOrderQueryResponseBodyModuleOrderBaseInfo) SetDepartId(v string) *CarOrderQueryResponseBodyModuleOrderBaseInfo {
	s.DepartId = &v
	return s
}

func (s *CarOrderQueryResponseBodyModuleOrderBaseInfo) SetDepartName(v string) *CarOrderQueryResponseBodyModuleOrderBaseInfo {
	s.DepartName = &v
	return s
}

func (s *CarOrderQueryResponseBodyModuleOrderBaseInfo) SetGmtCreate(v int64) *CarOrderQueryResponseBodyModuleOrderBaseInfo {
	s.GmtCreate = &v
	return s
}

func (s *CarOrderQueryResponseBodyModuleOrderBaseInfo) SetGmtModified(v int64) *CarOrderQueryResponseBodyModuleOrderBaseInfo {
	s.GmtModified = &v
	return s
}

func (s *CarOrderQueryResponseBodyModuleOrderBaseInfo) SetItineraryId(v string) *CarOrderQueryResponseBodyModuleOrderBaseInfo {
	s.ItineraryId = &v
	return s
}

func (s *CarOrderQueryResponseBodyModuleOrderBaseInfo) SetOrderId(v int64) *CarOrderQueryResponseBodyModuleOrderBaseInfo {
	s.OrderId = &v
	return s
}

func (s *CarOrderQueryResponseBodyModuleOrderBaseInfo) SetOrderStatus(v int32) *CarOrderQueryResponseBodyModuleOrderBaseInfo {
	s.OrderStatus = &v
	return s
}

func (s *CarOrderQueryResponseBodyModuleOrderBaseInfo) SetSubOrderId(v int64) *CarOrderQueryResponseBodyModuleOrderBaseInfo {
	s.SubOrderId = &v
	return s
}

func (s *CarOrderQueryResponseBodyModuleOrderBaseInfo) SetThirdDepartId(v string) *CarOrderQueryResponseBodyModuleOrderBaseInfo {
	s.ThirdDepartId = &v
	return s
}

func (s *CarOrderQueryResponseBodyModuleOrderBaseInfo) SetThirdpartApplyId(v string) *CarOrderQueryResponseBodyModuleOrderBaseInfo {
	s.ThirdpartApplyId = &v
	return s
}

func (s *CarOrderQueryResponseBodyModuleOrderBaseInfo) SetThirdpartBusinessId(v string) *CarOrderQueryResponseBodyModuleOrderBaseInfo {
	s.ThirdpartBusinessId = &v
	return s
}

func (s *CarOrderQueryResponseBodyModuleOrderBaseInfo) SetThirdpartItineraryId(v string) *CarOrderQueryResponseBodyModuleOrderBaseInfo {
	s.ThirdpartItineraryId = &v
	return s
}

func (s *CarOrderQueryResponseBodyModuleOrderBaseInfo) SetUserId(v string) *CarOrderQueryResponseBodyModuleOrderBaseInfo {
	s.UserId = &v
	return s
}

func (s *CarOrderQueryResponseBodyModuleOrderBaseInfo) SetUserName(v string) *CarOrderQueryResponseBodyModuleOrderBaseInfo {
	s.UserName = &v
	return s
}

func (s *CarOrderQueryResponseBodyModuleOrderBaseInfo) Validate() error {
	return dara.Validate(s)
}

type CarOrderQueryResponseBodyModulePassengerList struct {
	// example:
	//
	// costId
	CostCenterId *int64 `json:"cost_center_id,omitempty" xml:"cost_center_id,omitempty"`
	// example:
	//
	// costName
	CostCenterName *string `json:"cost_center_name,omitempty" xml:"cost_center_name,omitempty"`
	// example:
	//
	// costNumber
	CostCenterNumber *string `json:"cost_center_number,omitempty" xml:"cost_center_number,omitempty"`
	// example:
	//
	// projectCode1
	ProjectCode *string `json:"project_code,omitempty" xml:"project_code,omitempty"`
	// example:
	//
	// projectId
	ProjectId    *int64  `json:"project_id,omitempty" xml:"project_id,omitempty"`
	ProjectTitle *string `json:"project_title,omitempty" xml:"project_title,omitempty"`
	// example:
	//
	// thirdpartCostId
	ThirdpartCostCenterId *string `json:"thirdpart_cost_center_id,omitempty" xml:"thirdpart_cost_center_id,omitempty"`
	// example:
	//
	// thirdpartProjectId
	ThirdpartProjectId *string `json:"thirdpart_project_id,omitempty" xml:"thirdpart_project_id,omitempty"`
	// example:
	//
	// userId
	UserId   *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserName *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
	// example:
	//
	// 1
	UserType *int32 `json:"user_type,omitempty" xml:"user_type,omitempty"`
}

func (s CarOrderQueryResponseBodyModulePassengerList) String() string {
	return dara.Prettify(s)
}

func (s CarOrderQueryResponseBodyModulePassengerList) GoString() string {
	return s.String()
}

func (s *CarOrderQueryResponseBodyModulePassengerList) GetCostCenterId() *int64 {
	return s.CostCenterId
}

func (s *CarOrderQueryResponseBodyModulePassengerList) GetCostCenterName() *string {
	return s.CostCenterName
}

func (s *CarOrderQueryResponseBodyModulePassengerList) GetCostCenterNumber() *string {
	return s.CostCenterNumber
}

func (s *CarOrderQueryResponseBodyModulePassengerList) GetProjectCode() *string {
	return s.ProjectCode
}

func (s *CarOrderQueryResponseBodyModulePassengerList) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CarOrderQueryResponseBodyModulePassengerList) GetProjectTitle() *string {
	return s.ProjectTitle
}

func (s *CarOrderQueryResponseBodyModulePassengerList) GetThirdpartCostCenterId() *string {
	return s.ThirdpartCostCenterId
}

func (s *CarOrderQueryResponseBodyModulePassengerList) GetThirdpartProjectId() *string {
	return s.ThirdpartProjectId
}

func (s *CarOrderQueryResponseBodyModulePassengerList) GetUserId() *string {
	return s.UserId
}

func (s *CarOrderQueryResponseBodyModulePassengerList) GetUserName() *string {
	return s.UserName
}

func (s *CarOrderQueryResponseBodyModulePassengerList) GetUserType() *int32 {
	return s.UserType
}

func (s *CarOrderQueryResponseBodyModulePassengerList) SetCostCenterId(v int64) *CarOrderQueryResponseBodyModulePassengerList {
	s.CostCenterId = &v
	return s
}

func (s *CarOrderQueryResponseBodyModulePassengerList) SetCostCenterName(v string) *CarOrderQueryResponseBodyModulePassengerList {
	s.CostCenterName = &v
	return s
}

func (s *CarOrderQueryResponseBodyModulePassengerList) SetCostCenterNumber(v string) *CarOrderQueryResponseBodyModulePassengerList {
	s.CostCenterNumber = &v
	return s
}

func (s *CarOrderQueryResponseBodyModulePassengerList) SetProjectCode(v string) *CarOrderQueryResponseBodyModulePassengerList {
	s.ProjectCode = &v
	return s
}

func (s *CarOrderQueryResponseBodyModulePassengerList) SetProjectId(v int64) *CarOrderQueryResponseBodyModulePassengerList {
	s.ProjectId = &v
	return s
}

func (s *CarOrderQueryResponseBodyModulePassengerList) SetProjectTitle(v string) *CarOrderQueryResponseBodyModulePassengerList {
	s.ProjectTitle = &v
	return s
}

func (s *CarOrderQueryResponseBodyModulePassengerList) SetThirdpartCostCenterId(v string) *CarOrderQueryResponseBodyModulePassengerList {
	s.ThirdpartCostCenterId = &v
	return s
}

func (s *CarOrderQueryResponseBodyModulePassengerList) SetThirdpartProjectId(v string) *CarOrderQueryResponseBodyModulePassengerList {
	s.ThirdpartProjectId = &v
	return s
}

func (s *CarOrderQueryResponseBodyModulePassengerList) SetUserId(v string) *CarOrderQueryResponseBodyModulePassengerList {
	s.UserId = &v
	return s
}

func (s *CarOrderQueryResponseBodyModulePassengerList) SetUserName(v string) *CarOrderQueryResponseBodyModulePassengerList {
	s.UserName = &v
	return s
}

func (s *CarOrderQueryResponseBodyModulePassengerList) SetUserType(v int32) *CarOrderQueryResponseBodyModulePassengerList {
	s.UserType = &v
	return s
}

func (s *CarOrderQueryResponseBodyModulePassengerList) Validate() error {
	return dara.Validate(s)
}

type CarOrderQueryResponseBodyModulePriceInfoList struct {
	// example:
	//
	// 1
	CategoryCode *int32 `json:"category_code,omitempty" xml:"category_code,omitempty"`
	// example:
	//
	// 1669274251000
	GmtCreate *int64 `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// example:
	//
	// 4
	PayType *int32 `json:"pay_type,omitempty" xml:"pay_type,omitempty"`
	// example:
	//
	// 1000
	PersonPrice *int64 `json:"person_price,omitempty" xml:"person_price,omitempty"`
	// example:
	//
	// 2000
	Price *int64 `json:"price,omitempty" xml:"price,omitempty"`
	// example:
	//
	// 1012000000001
	TradeId *string `json:"trade_id,omitempty" xml:"trade_id,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CarOrderQueryResponseBodyModulePriceInfoList) String() string {
	return dara.Prettify(s)
}

func (s CarOrderQueryResponseBodyModulePriceInfoList) GoString() string {
	return s.String()
}

func (s *CarOrderQueryResponseBodyModulePriceInfoList) GetCategoryCode() *int32 {
	return s.CategoryCode
}

func (s *CarOrderQueryResponseBodyModulePriceInfoList) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *CarOrderQueryResponseBodyModulePriceInfoList) GetPayType() *int32 {
	return s.PayType
}

func (s *CarOrderQueryResponseBodyModulePriceInfoList) GetPersonPrice() *int64 {
	return s.PersonPrice
}

func (s *CarOrderQueryResponseBodyModulePriceInfoList) GetPrice() *int64 {
	return s.Price
}

func (s *CarOrderQueryResponseBodyModulePriceInfoList) GetTradeId() *string {
	return s.TradeId
}

func (s *CarOrderQueryResponseBodyModulePriceInfoList) GetType() *int32 {
	return s.Type
}

func (s *CarOrderQueryResponseBodyModulePriceInfoList) SetCategoryCode(v int32) *CarOrderQueryResponseBodyModulePriceInfoList {
	s.CategoryCode = &v
	return s
}

func (s *CarOrderQueryResponseBodyModulePriceInfoList) SetGmtCreate(v int64) *CarOrderQueryResponseBodyModulePriceInfoList {
	s.GmtCreate = &v
	return s
}

func (s *CarOrderQueryResponseBodyModulePriceInfoList) SetPayType(v int32) *CarOrderQueryResponseBodyModulePriceInfoList {
	s.PayType = &v
	return s
}

func (s *CarOrderQueryResponseBodyModulePriceInfoList) SetPersonPrice(v int64) *CarOrderQueryResponseBodyModulePriceInfoList {
	s.PersonPrice = &v
	return s
}

func (s *CarOrderQueryResponseBodyModulePriceInfoList) SetPrice(v int64) *CarOrderQueryResponseBodyModulePriceInfoList {
	s.Price = &v
	return s
}

func (s *CarOrderQueryResponseBodyModulePriceInfoList) SetTradeId(v string) *CarOrderQueryResponseBodyModulePriceInfoList {
	s.TradeId = &v
	return s
}

func (s *CarOrderQueryResponseBodyModulePriceInfoList) SetType(v int32) *CarOrderQueryResponseBodyModulePriceInfoList {
	s.Type = &v
	return s
}

func (s *CarOrderQueryResponseBodyModulePriceInfoList) Validate() error {
	return dara.Validate(s)
}
