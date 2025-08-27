// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCarOrderListQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CarOrderListQueryResponseBody
	GetCode() *string
	SetMessage(v string) *CarOrderListQueryResponseBody
	GetMessage() *string
	SetModule(v []*CarOrderListQueryResponseBodyModule) *CarOrderListQueryResponseBody
	GetModule() []*CarOrderListQueryResponseBodyModule
	SetPageInfo(v *CarOrderListQueryResponseBodyPageInfo) *CarOrderListQueryResponseBody
	GetPageInfo() *CarOrderListQueryResponseBodyPageInfo
	SetRequestId(v string) *CarOrderListQueryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CarOrderListQueryResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *CarOrderListQueryResponseBody
	GetTraceId() *string
}

type CarOrderListQueryResponseBody struct {
	// example:
	//
	// SUCCESS
	Code     *string                                `json:"code,omitempty" xml:"code,omitempty"`
	Message  *string                                `json:"message,omitempty" xml:"message,omitempty"`
	Module   []*CarOrderListQueryResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Repeated"`
	PageInfo *CarOrderListQueryResponseBodyPageInfo `json:"page_info,omitempty" xml:"page_info,omitempty" type:"Struct"`
	// example:
	//
	// B72B39C8-32DE-558D-AD1C-D53F11F6ADFE
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

func (s CarOrderListQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CarOrderListQueryResponseBody) GoString() string {
	return s.String()
}

func (s *CarOrderListQueryResponseBody) GetCode() *string {
	return s.Code
}

func (s *CarOrderListQueryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CarOrderListQueryResponseBody) GetModule() []*CarOrderListQueryResponseBodyModule {
	return s.Module
}

func (s *CarOrderListQueryResponseBody) GetPageInfo() *CarOrderListQueryResponseBodyPageInfo {
	return s.PageInfo
}

func (s *CarOrderListQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CarOrderListQueryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CarOrderListQueryResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *CarOrderListQueryResponseBody) SetCode(v string) *CarOrderListQueryResponseBody {
	s.Code = &v
	return s
}

func (s *CarOrderListQueryResponseBody) SetMessage(v string) *CarOrderListQueryResponseBody {
	s.Message = &v
	return s
}

func (s *CarOrderListQueryResponseBody) SetModule(v []*CarOrderListQueryResponseBodyModule) *CarOrderListQueryResponseBody {
	s.Module = v
	return s
}

func (s *CarOrderListQueryResponseBody) SetPageInfo(v *CarOrderListQueryResponseBodyPageInfo) *CarOrderListQueryResponseBody {
	s.PageInfo = v
	return s
}

func (s *CarOrderListQueryResponseBody) SetRequestId(v string) *CarOrderListQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CarOrderListQueryResponseBody) SetSuccess(v bool) *CarOrderListQueryResponseBody {
	s.Success = &v
	return s
}

func (s *CarOrderListQueryResponseBody) SetTraceId(v string) *CarOrderListQueryResponseBody {
	s.TraceId = &v
	return s
}

func (s *CarOrderListQueryResponseBody) Validate() error {
	return dara.Validate(s)
}

type CarOrderListQueryResponseBodyModule struct {
	// example:
	//
	// 117429516
	ApplyId *int64 `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	// example:
	//
	// 201802031353000525653
	ApplyShowId *string `json:"apply_show_id,omitempty" xml:"apply_show_id,omitempty"`
	BtripTitle  *string `json:"btrip_title,omitempty" xml:"btrip_title,omitempty"`
	// example:
	//
	// TRAVEL
	BusinessCategory *string `json:"business_category,omitempty" xml:"business_category,omitempty"`
	// example:
	//
	// 2022-07-04T16:13Z
	CancelTime *string `json:"cancel_time,omitempty" xml:"cancel_time,omitempty"`
	CarInfo    *string `json:"car_info,omitempty" xml:"car_info,omitempty"`
	// example:
	//
	// 1
	CarLevel *int32  `json:"car_level,omitempty" xml:"car_level,omitempty"`
	CorpId   *string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	CorpName *string `json:"corp_name,omitempty" xml:"corp_name,omitempty"`
	// example:
	//
	// 11376
	CostCenterId   *int64  `json:"cost_center_id,omitempty" xml:"cost_center_id,omitempty"`
	CostCenterName *string `json:"cost_center_name,omitempty" xml:"cost_center_name,omitempty"`
	// example:
	//
	// CT-134JHK
	CostCenterNumber *string `json:"cost_center_number,omitempty" xml:"cost_center_number,omitempty"`
	DeptId           *int64  `json:"dept_id,omitempty" xml:"dept_id,omitempty"`
	DeptName         *string `json:"dept_name,omitempty" xml:"dept_name,omitempty"`
	// example:
	//
	// 2022-07-04T16:13Z
	DriverConfirmTime *string `json:"driver_confirm_time,omitempty" xml:"driver_confirm_time,omitempty"`
	// example:
	//
	// 100
	EstimatePrice  *float64 `json:"estimate_price,omitempty" xml:"estimate_price,omitempty"`
	FromAddress    *string  `json:"from_address,omitempty" xml:"from_address,omitempty"`
	FromCityAdCode *string  `json:"from_city_ad_code,omitempty" xml:"from_city_ad_code,omitempty"`
	FromCityName   *string  `json:"from_city_name,omitempty" xml:"from_city_name,omitempty"`
	// example:
	//
	// 2022-07-04T16:13Z
	GmtCreate *string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// example:
	//
	// 2022-07-04T16:13Z
	GmtModified *string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// example:
	//
	// 3615085
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 11876
	InvoiceId    *int64  `json:"invoice_id,omitempty" xml:"invoice_id,omitempty"`
	InvoiceTitle *string `json:"invoice_title,omitempty" xml:"invoice_title,omitempty"`
	// example:
	//
	// true
	IsSpecial     *bool   `json:"is_special,omitempty" xml:"is_special,omitempty"`
	Memo          *string `json:"memo,omitempty" xml:"memo,omitempty"`
	OrderId       *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	OrderStatus   *int32  `json:"order_status,omitempty" xml:"order_status,omitempty"`
	PassengerName *string `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
	// example:
	//
	// 2022-07-04T16:13Z
	PayTime       *string                                             `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	PriceInfoList []*CarOrderListQueryResponseBodyModulePriceInfoList `json:"price_info_list,omitempty" xml:"price_info_list,omitempty" type:"Repeated"`
	// example:
	//
	// P- JI87KK
	ProjectCode *string `json:"project_code,omitempty" xml:"project_code,omitempty"`
	// example:
	//
	// 11546
	ProjectId    *int64  `json:"project_id,omitempty" xml:"project_id,omitempty"`
	ProjectTitle *string `json:"project_title,omitempty" xml:"project_title,omitempty"`
	// example:
	//
	// 2
	Provider *int32 `json:"provider,omitempty" xml:"provider,omitempty"`
	// example:
	//
	// 2022-07-04T16:13Z
	PublishTime        *string `json:"publish_time,omitempty" xml:"publish_time,omitempty"`
	RealFromAddress    *string `json:"real_from_address,omitempty" xml:"real_from_address,omitempty"`
	RealFromCityAdCode *string `json:"real_from_city_ad_code,omitempty" xml:"real_from_city_ad_code,omitempty"`
	RealFromCityName   *string `json:"real_from_city_name,omitempty" xml:"real_from_city_name,omitempty"`
	RealToAddress      *string `json:"real_to_address,omitempty" xml:"real_to_address,omitempty"`
	RealToCityAdCode   *string `json:"real_to_city_ad_code,omitempty" xml:"real_to_city_ad_code,omitempty"`
	RealToCityName     *string `json:"real_to_city_name,omitempty" xml:"real_to_city_name,omitempty"`
	// example:
	//
	// 1
	ServiceType  *int32    `json:"service_type,omitempty" xml:"service_type,omitempty"`
	SpecialTypes []*string `json:"special_types,omitempty" xml:"special_types,omitempty" type:"Repeated"`
	// example:
	//
	// 2022-07-04T16:13Z
	TakenTime *string `json:"taken_time,omitempty" xml:"taken_time,omitempty"`
	// example:
	//
	// CS-OIPK34H
	ThirdpartApplyId    *string `json:"thirdpart_apply_id,omitempty" xml:"thirdpart_apply_id,omitempty"`
	ThirdpartBusinessId *string `json:"thirdpart_business_id,omitempty" xml:"thirdpart_business_id,omitempty"`
	// example:
	//
	// CS-112JKDF
	ThirdpartItineraryId *string `json:"thirdpart_itinerary_id,omitempty" xml:"thirdpart_itinerary_id,omitempty"`
	ToAddress            *string `json:"to_address,omitempty" xml:"to_address,omitempty"`
	ToCityAdCode         *string `json:"to_city_ad_code,omitempty" xml:"to_city_ad_code,omitempty"`
	ToCityName           *string `json:"to_city_name,omitempty" xml:"to_city_name,omitempty"`
	// example:
	//
	// 1.2
	TravelDistance    *float64                                                `json:"travel_distance,omitempty" xml:"travel_distance,omitempty"`
	UserAffiliateList []*CarOrderListQueryResponseBodyModuleUserAffiliateList `json:"user_affiliate_list,omitempty" xml:"user_affiliate_list,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	UserConfirm *int32  `json:"user_confirm,omitempty" xml:"user_confirm,omitempty"`
	UserId      *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserName    *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s CarOrderListQueryResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s CarOrderListQueryResponseBodyModule) GoString() string {
	return s.String()
}

func (s *CarOrderListQueryResponseBodyModule) GetApplyId() *int64 {
	return s.ApplyId
}

func (s *CarOrderListQueryResponseBodyModule) GetApplyShowId() *string {
	return s.ApplyShowId
}

func (s *CarOrderListQueryResponseBodyModule) GetBtripTitle() *string {
	return s.BtripTitle
}

func (s *CarOrderListQueryResponseBodyModule) GetBusinessCategory() *string {
	return s.BusinessCategory
}

func (s *CarOrderListQueryResponseBodyModule) GetCancelTime() *string {
	return s.CancelTime
}

func (s *CarOrderListQueryResponseBodyModule) GetCarInfo() *string {
	return s.CarInfo
}

func (s *CarOrderListQueryResponseBodyModule) GetCarLevel() *int32 {
	return s.CarLevel
}

func (s *CarOrderListQueryResponseBodyModule) GetCorpId() *string {
	return s.CorpId
}

func (s *CarOrderListQueryResponseBodyModule) GetCorpName() *string {
	return s.CorpName
}

func (s *CarOrderListQueryResponseBodyModule) GetCostCenterId() *int64 {
	return s.CostCenterId
}

func (s *CarOrderListQueryResponseBodyModule) GetCostCenterName() *string {
	return s.CostCenterName
}

func (s *CarOrderListQueryResponseBodyModule) GetCostCenterNumber() *string {
	return s.CostCenterNumber
}

func (s *CarOrderListQueryResponseBodyModule) GetDeptId() *int64 {
	return s.DeptId
}

func (s *CarOrderListQueryResponseBodyModule) GetDeptName() *string {
	return s.DeptName
}

func (s *CarOrderListQueryResponseBodyModule) GetDriverConfirmTime() *string {
	return s.DriverConfirmTime
}

func (s *CarOrderListQueryResponseBodyModule) GetEstimatePrice() *float64 {
	return s.EstimatePrice
}

func (s *CarOrderListQueryResponseBodyModule) GetFromAddress() *string {
	return s.FromAddress
}

func (s *CarOrderListQueryResponseBodyModule) GetFromCityAdCode() *string {
	return s.FromCityAdCode
}

func (s *CarOrderListQueryResponseBodyModule) GetFromCityName() *string {
	return s.FromCityName
}

func (s *CarOrderListQueryResponseBodyModule) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *CarOrderListQueryResponseBodyModule) GetGmtModified() *string {
	return s.GmtModified
}

func (s *CarOrderListQueryResponseBodyModule) GetId() *int64 {
	return s.Id
}

func (s *CarOrderListQueryResponseBodyModule) GetInvoiceId() *int64 {
	return s.InvoiceId
}

func (s *CarOrderListQueryResponseBodyModule) GetInvoiceTitle() *string {
	return s.InvoiceTitle
}

func (s *CarOrderListQueryResponseBodyModule) GetIsSpecial() *bool {
	return s.IsSpecial
}

func (s *CarOrderListQueryResponseBodyModule) GetMemo() *string {
	return s.Memo
}

func (s *CarOrderListQueryResponseBodyModule) GetOrderId() *string {
	return s.OrderId
}

func (s *CarOrderListQueryResponseBodyModule) GetOrderStatus() *int32 {
	return s.OrderStatus
}

func (s *CarOrderListQueryResponseBodyModule) GetPassengerName() *string {
	return s.PassengerName
}

func (s *CarOrderListQueryResponseBodyModule) GetPayTime() *string {
	return s.PayTime
}

func (s *CarOrderListQueryResponseBodyModule) GetPriceInfoList() []*CarOrderListQueryResponseBodyModulePriceInfoList {
	return s.PriceInfoList
}

func (s *CarOrderListQueryResponseBodyModule) GetProjectCode() *string {
	return s.ProjectCode
}

func (s *CarOrderListQueryResponseBodyModule) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CarOrderListQueryResponseBodyModule) GetProjectTitle() *string {
	return s.ProjectTitle
}

func (s *CarOrderListQueryResponseBodyModule) GetProvider() *int32 {
	return s.Provider
}

func (s *CarOrderListQueryResponseBodyModule) GetPublishTime() *string {
	return s.PublishTime
}

func (s *CarOrderListQueryResponseBodyModule) GetRealFromAddress() *string {
	return s.RealFromAddress
}

func (s *CarOrderListQueryResponseBodyModule) GetRealFromCityAdCode() *string {
	return s.RealFromCityAdCode
}

func (s *CarOrderListQueryResponseBodyModule) GetRealFromCityName() *string {
	return s.RealFromCityName
}

func (s *CarOrderListQueryResponseBodyModule) GetRealToAddress() *string {
	return s.RealToAddress
}

func (s *CarOrderListQueryResponseBodyModule) GetRealToCityAdCode() *string {
	return s.RealToCityAdCode
}

func (s *CarOrderListQueryResponseBodyModule) GetRealToCityName() *string {
	return s.RealToCityName
}

func (s *CarOrderListQueryResponseBodyModule) GetServiceType() *int32 {
	return s.ServiceType
}

func (s *CarOrderListQueryResponseBodyModule) GetSpecialTypes() []*string {
	return s.SpecialTypes
}

func (s *CarOrderListQueryResponseBodyModule) GetTakenTime() *string {
	return s.TakenTime
}

func (s *CarOrderListQueryResponseBodyModule) GetThirdpartApplyId() *string {
	return s.ThirdpartApplyId
}

func (s *CarOrderListQueryResponseBodyModule) GetThirdpartBusinessId() *string {
	return s.ThirdpartBusinessId
}

func (s *CarOrderListQueryResponseBodyModule) GetThirdpartItineraryId() *string {
	return s.ThirdpartItineraryId
}

func (s *CarOrderListQueryResponseBodyModule) GetToAddress() *string {
	return s.ToAddress
}

func (s *CarOrderListQueryResponseBodyModule) GetToCityAdCode() *string {
	return s.ToCityAdCode
}

func (s *CarOrderListQueryResponseBodyModule) GetToCityName() *string {
	return s.ToCityName
}

func (s *CarOrderListQueryResponseBodyModule) GetTravelDistance() *float64 {
	return s.TravelDistance
}

func (s *CarOrderListQueryResponseBodyModule) GetUserAffiliateList() []*CarOrderListQueryResponseBodyModuleUserAffiliateList {
	return s.UserAffiliateList
}

func (s *CarOrderListQueryResponseBodyModule) GetUserConfirm() *int32 {
	return s.UserConfirm
}

func (s *CarOrderListQueryResponseBodyModule) GetUserId() *string {
	return s.UserId
}

func (s *CarOrderListQueryResponseBodyModule) GetUserName() *string {
	return s.UserName
}

func (s *CarOrderListQueryResponseBodyModule) SetApplyId(v int64) *CarOrderListQueryResponseBodyModule {
	s.ApplyId = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetApplyShowId(v string) *CarOrderListQueryResponseBodyModule {
	s.ApplyShowId = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetBtripTitle(v string) *CarOrderListQueryResponseBodyModule {
	s.BtripTitle = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetBusinessCategory(v string) *CarOrderListQueryResponseBodyModule {
	s.BusinessCategory = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetCancelTime(v string) *CarOrderListQueryResponseBodyModule {
	s.CancelTime = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetCarInfo(v string) *CarOrderListQueryResponseBodyModule {
	s.CarInfo = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetCarLevel(v int32) *CarOrderListQueryResponseBodyModule {
	s.CarLevel = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetCorpId(v string) *CarOrderListQueryResponseBodyModule {
	s.CorpId = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetCorpName(v string) *CarOrderListQueryResponseBodyModule {
	s.CorpName = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetCostCenterId(v int64) *CarOrderListQueryResponseBodyModule {
	s.CostCenterId = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetCostCenterName(v string) *CarOrderListQueryResponseBodyModule {
	s.CostCenterName = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetCostCenterNumber(v string) *CarOrderListQueryResponseBodyModule {
	s.CostCenterNumber = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetDeptId(v int64) *CarOrderListQueryResponseBodyModule {
	s.DeptId = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetDeptName(v string) *CarOrderListQueryResponseBodyModule {
	s.DeptName = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetDriverConfirmTime(v string) *CarOrderListQueryResponseBodyModule {
	s.DriverConfirmTime = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetEstimatePrice(v float64) *CarOrderListQueryResponseBodyModule {
	s.EstimatePrice = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetFromAddress(v string) *CarOrderListQueryResponseBodyModule {
	s.FromAddress = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetFromCityAdCode(v string) *CarOrderListQueryResponseBodyModule {
	s.FromCityAdCode = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetFromCityName(v string) *CarOrderListQueryResponseBodyModule {
	s.FromCityName = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetGmtCreate(v string) *CarOrderListQueryResponseBodyModule {
	s.GmtCreate = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetGmtModified(v string) *CarOrderListQueryResponseBodyModule {
	s.GmtModified = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetId(v int64) *CarOrderListQueryResponseBodyModule {
	s.Id = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetInvoiceId(v int64) *CarOrderListQueryResponseBodyModule {
	s.InvoiceId = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetInvoiceTitle(v string) *CarOrderListQueryResponseBodyModule {
	s.InvoiceTitle = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetIsSpecial(v bool) *CarOrderListQueryResponseBodyModule {
	s.IsSpecial = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetMemo(v string) *CarOrderListQueryResponseBodyModule {
	s.Memo = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetOrderId(v string) *CarOrderListQueryResponseBodyModule {
	s.OrderId = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetOrderStatus(v int32) *CarOrderListQueryResponseBodyModule {
	s.OrderStatus = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetPassengerName(v string) *CarOrderListQueryResponseBodyModule {
	s.PassengerName = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetPayTime(v string) *CarOrderListQueryResponseBodyModule {
	s.PayTime = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetPriceInfoList(v []*CarOrderListQueryResponseBodyModulePriceInfoList) *CarOrderListQueryResponseBodyModule {
	s.PriceInfoList = v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetProjectCode(v string) *CarOrderListQueryResponseBodyModule {
	s.ProjectCode = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetProjectId(v int64) *CarOrderListQueryResponseBodyModule {
	s.ProjectId = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetProjectTitle(v string) *CarOrderListQueryResponseBodyModule {
	s.ProjectTitle = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetProvider(v int32) *CarOrderListQueryResponseBodyModule {
	s.Provider = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetPublishTime(v string) *CarOrderListQueryResponseBodyModule {
	s.PublishTime = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetRealFromAddress(v string) *CarOrderListQueryResponseBodyModule {
	s.RealFromAddress = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetRealFromCityAdCode(v string) *CarOrderListQueryResponseBodyModule {
	s.RealFromCityAdCode = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetRealFromCityName(v string) *CarOrderListQueryResponseBodyModule {
	s.RealFromCityName = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetRealToAddress(v string) *CarOrderListQueryResponseBodyModule {
	s.RealToAddress = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetRealToCityAdCode(v string) *CarOrderListQueryResponseBodyModule {
	s.RealToCityAdCode = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetRealToCityName(v string) *CarOrderListQueryResponseBodyModule {
	s.RealToCityName = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetServiceType(v int32) *CarOrderListQueryResponseBodyModule {
	s.ServiceType = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetSpecialTypes(v []*string) *CarOrderListQueryResponseBodyModule {
	s.SpecialTypes = v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetTakenTime(v string) *CarOrderListQueryResponseBodyModule {
	s.TakenTime = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetThirdpartApplyId(v string) *CarOrderListQueryResponseBodyModule {
	s.ThirdpartApplyId = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetThirdpartBusinessId(v string) *CarOrderListQueryResponseBodyModule {
	s.ThirdpartBusinessId = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetThirdpartItineraryId(v string) *CarOrderListQueryResponseBodyModule {
	s.ThirdpartItineraryId = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetToAddress(v string) *CarOrderListQueryResponseBodyModule {
	s.ToAddress = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetToCityAdCode(v string) *CarOrderListQueryResponseBodyModule {
	s.ToCityAdCode = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetToCityName(v string) *CarOrderListQueryResponseBodyModule {
	s.ToCityName = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetTravelDistance(v float64) *CarOrderListQueryResponseBodyModule {
	s.TravelDistance = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetUserAffiliateList(v []*CarOrderListQueryResponseBodyModuleUserAffiliateList) *CarOrderListQueryResponseBodyModule {
	s.UserAffiliateList = v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetUserConfirm(v int32) *CarOrderListQueryResponseBodyModule {
	s.UserConfirm = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetUserId(v string) *CarOrderListQueryResponseBodyModule {
	s.UserId = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) SetUserName(v string) *CarOrderListQueryResponseBodyModule {
	s.UserName = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type CarOrderListQueryResponseBodyModulePriceInfoList struct {
	// example:
	//
	// 1
	CategoryCode *int32 `json:"category_code,omitempty" xml:"category_code,omitempty"`
	// example:
	//
	// 1
	CategoryType *int32 `json:"category_type,omitempty" xml:"category_type,omitempty"`
	// example:
	//
	// 2022-07-04T16:13Z
	GmtCreate     *string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	PassengerName *string `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
	// example:
	//
	// 1
	PayType *int32 `json:"pay_type,omitempty" xml:"pay_type,omitempty"`
	// example:
	//
	// 100
	PersonPrice *float64 `json:"person_price,omitempty" xml:"person_price,omitempty"`
	// example:
	//
	// 100
	Price *float64 `json:"price,omitempty" xml:"price,omitempty"`
	// example:
	//
	// 8908076767
	TradeId *string `json:"trade_id,omitempty" xml:"trade_id,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CarOrderListQueryResponseBodyModulePriceInfoList) String() string {
	return dara.Prettify(s)
}

func (s CarOrderListQueryResponseBodyModulePriceInfoList) GoString() string {
	return s.String()
}

func (s *CarOrderListQueryResponseBodyModulePriceInfoList) GetCategoryCode() *int32 {
	return s.CategoryCode
}

func (s *CarOrderListQueryResponseBodyModulePriceInfoList) GetCategoryType() *int32 {
	return s.CategoryType
}

func (s *CarOrderListQueryResponseBodyModulePriceInfoList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *CarOrderListQueryResponseBodyModulePriceInfoList) GetPassengerName() *string {
	return s.PassengerName
}

func (s *CarOrderListQueryResponseBodyModulePriceInfoList) GetPayType() *int32 {
	return s.PayType
}

func (s *CarOrderListQueryResponseBodyModulePriceInfoList) GetPersonPrice() *float64 {
	return s.PersonPrice
}

func (s *CarOrderListQueryResponseBodyModulePriceInfoList) GetPrice() *float64 {
	return s.Price
}

func (s *CarOrderListQueryResponseBodyModulePriceInfoList) GetTradeId() *string {
	return s.TradeId
}

func (s *CarOrderListQueryResponseBodyModulePriceInfoList) GetType() *int32 {
	return s.Type
}

func (s *CarOrderListQueryResponseBodyModulePriceInfoList) SetCategoryCode(v int32) *CarOrderListQueryResponseBodyModulePriceInfoList {
	s.CategoryCode = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModulePriceInfoList) SetCategoryType(v int32) *CarOrderListQueryResponseBodyModulePriceInfoList {
	s.CategoryType = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModulePriceInfoList) SetGmtCreate(v string) *CarOrderListQueryResponseBodyModulePriceInfoList {
	s.GmtCreate = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModulePriceInfoList) SetPassengerName(v string) *CarOrderListQueryResponseBodyModulePriceInfoList {
	s.PassengerName = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModulePriceInfoList) SetPayType(v int32) *CarOrderListQueryResponseBodyModulePriceInfoList {
	s.PayType = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModulePriceInfoList) SetPersonPrice(v float64) *CarOrderListQueryResponseBodyModulePriceInfoList {
	s.PersonPrice = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModulePriceInfoList) SetPrice(v float64) *CarOrderListQueryResponseBodyModulePriceInfoList {
	s.Price = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModulePriceInfoList) SetTradeId(v string) *CarOrderListQueryResponseBodyModulePriceInfoList {
	s.TradeId = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModulePriceInfoList) SetType(v int32) *CarOrderListQueryResponseBodyModulePriceInfoList {
	s.Type = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModulePriceInfoList) Validate() error {
	return dara.Validate(s)
}

type CarOrderListQueryResponseBodyModuleUserAffiliateList struct {
	// example:
	//
	// 11342
	UserId   *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserName *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s CarOrderListQueryResponseBodyModuleUserAffiliateList) String() string {
	return dara.Prettify(s)
}

func (s CarOrderListQueryResponseBodyModuleUserAffiliateList) GoString() string {
	return s.String()
}

func (s *CarOrderListQueryResponseBodyModuleUserAffiliateList) GetUserId() *string {
	return s.UserId
}

func (s *CarOrderListQueryResponseBodyModuleUserAffiliateList) GetUserName() *string {
	return s.UserName
}

func (s *CarOrderListQueryResponseBodyModuleUserAffiliateList) SetUserId(v string) *CarOrderListQueryResponseBodyModuleUserAffiliateList {
	s.UserId = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModuleUserAffiliateList) SetUserName(v string) *CarOrderListQueryResponseBodyModuleUserAffiliateList {
	s.UserName = &v
	return s
}

func (s *CarOrderListQueryResponseBodyModuleUserAffiliateList) Validate() error {
	return dara.Validate(s)
}

type CarOrderListQueryResponseBodyPageInfo struct {
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// example:
	//
	// 1000
	TotalNumber *int32 `json:"total_number,omitempty" xml:"total_number,omitempty"`
}

func (s CarOrderListQueryResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s CarOrderListQueryResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *CarOrderListQueryResponseBodyPageInfo) GetPage() *int32 {
	return s.Page
}

func (s *CarOrderListQueryResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *CarOrderListQueryResponseBodyPageInfo) GetTotalNumber() *int32 {
	return s.TotalNumber
}

func (s *CarOrderListQueryResponseBodyPageInfo) SetPage(v int32) *CarOrderListQueryResponseBodyPageInfo {
	s.Page = &v
	return s
}

func (s *CarOrderListQueryResponseBodyPageInfo) SetPageSize(v int32) *CarOrderListQueryResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *CarOrderListQueryResponseBodyPageInfo) SetTotalNumber(v int32) *CarOrderListQueryResponseBodyPageInfo {
	s.TotalNumber = &v
	return s
}

func (s *CarOrderListQueryResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
