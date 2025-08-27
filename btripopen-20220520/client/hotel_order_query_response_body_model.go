// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelOrderQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *HotelOrderQueryResponseBody
	GetCode() *string
	SetMessage(v string) *HotelOrderQueryResponseBody
	GetMessage() *string
	SetModule(v *HotelOrderQueryResponseBodyModule) *HotelOrderQueryResponseBody
	GetModule() *HotelOrderQueryResponseBodyModule
	SetRequestId(v string) *HotelOrderQueryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *HotelOrderQueryResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *HotelOrderQueryResponseBody
	GetTraceId() *string
}

type HotelOrderQueryResponseBody struct {
	// example:
	//
	// SUCCESS
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// module。
	Module *HotelOrderQueryResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
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

func (s HotelOrderQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderQueryResponseBody) GoString() string {
	return s.String()
}

func (s *HotelOrderQueryResponseBody) GetCode() *string {
	return s.Code
}

func (s *HotelOrderQueryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *HotelOrderQueryResponseBody) GetModule() *HotelOrderQueryResponseBodyModule {
	return s.Module
}

func (s *HotelOrderQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *HotelOrderQueryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *HotelOrderQueryResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *HotelOrderQueryResponseBody) SetCode(v string) *HotelOrderQueryResponseBody {
	s.Code = &v
	return s
}

func (s *HotelOrderQueryResponseBody) SetMessage(v string) *HotelOrderQueryResponseBody {
	s.Message = &v
	return s
}

func (s *HotelOrderQueryResponseBody) SetModule(v *HotelOrderQueryResponseBodyModule) *HotelOrderQueryResponseBody {
	s.Module = v
	return s
}

func (s *HotelOrderQueryResponseBody) SetRequestId(v string) *HotelOrderQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *HotelOrderQueryResponseBody) SetSuccess(v bool) *HotelOrderQueryResponseBody {
	s.Success = &v
	return s
}

func (s *HotelOrderQueryResponseBody) SetTraceId(v string) *HotelOrderQueryResponseBody {
	s.TraceId = &v
	return s
}

func (s *HotelOrderQueryResponseBody) Validate() error {
	return dara.Validate(s)
}

type HotelOrderQueryResponseBodyModule struct {
	HotelInfo     *HotelOrderQueryResponseBodyModuleHotelInfo       `json:"hotel_info,omitempty" xml:"hotel_info,omitempty" type:"Struct"`
	InvoiceInfo   *HotelOrderQueryResponseBodyModuleInvoiceInfo     `json:"invoice_info,omitempty" xml:"invoice_info,omitempty" type:"Struct"`
	OrderBaseInfo *HotelOrderQueryResponseBodyModuleOrderBaseInfo   `json:"order_base_info,omitempty" xml:"order_base_info,omitempty" type:"Struct"`
	PassengerList []*HotelOrderQueryResponseBodyModulePassengerList `json:"passenger_list,omitempty" xml:"passenger_list,omitempty" type:"Repeated"`
	PriceInfoList []*HotelOrderQueryResponseBodyModulePriceInfoList `json:"price_info_list,omitempty" xml:"price_info_list,omitempty" type:"Repeated"`
}

func (s HotelOrderQueryResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderQueryResponseBodyModule) GoString() string {
	return s.String()
}

func (s *HotelOrderQueryResponseBodyModule) GetHotelInfo() *HotelOrderQueryResponseBodyModuleHotelInfo {
	return s.HotelInfo
}

func (s *HotelOrderQueryResponseBodyModule) GetInvoiceInfo() *HotelOrderQueryResponseBodyModuleInvoiceInfo {
	return s.InvoiceInfo
}

func (s *HotelOrderQueryResponseBodyModule) GetOrderBaseInfo() *HotelOrderQueryResponseBodyModuleOrderBaseInfo {
	return s.OrderBaseInfo
}

func (s *HotelOrderQueryResponseBodyModule) GetPassengerList() []*HotelOrderQueryResponseBodyModulePassengerList {
	return s.PassengerList
}

func (s *HotelOrderQueryResponseBodyModule) GetPriceInfoList() []*HotelOrderQueryResponseBodyModulePriceInfoList {
	return s.PriceInfoList
}

func (s *HotelOrderQueryResponseBodyModule) SetHotelInfo(v *HotelOrderQueryResponseBodyModuleHotelInfo) *HotelOrderQueryResponseBodyModule {
	s.HotelInfo = v
	return s
}

func (s *HotelOrderQueryResponseBodyModule) SetInvoiceInfo(v *HotelOrderQueryResponseBodyModuleInvoiceInfo) *HotelOrderQueryResponseBodyModule {
	s.InvoiceInfo = v
	return s
}

func (s *HotelOrderQueryResponseBodyModule) SetOrderBaseInfo(v *HotelOrderQueryResponseBodyModuleOrderBaseInfo) *HotelOrderQueryResponseBodyModule {
	s.OrderBaseInfo = v
	return s
}

func (s *HotelOrderQueryResponseBodyModule) SetPassengerList(v []*HotelOrderQueryResponseBodyModulePassengerList) *HotelOrderQueryResponseBodyModule {
	s.PassengerList = v
	return s
}

func (s *HotelOrderQueryResponseBodyModule) SetPriceInfoList(v []*HotelOrderQueryResponseBodyModulePriceInfoList) *HotelOrderQueryResponseBodyModule {
	s.PriceInfoList = v
	return s
}

func (s *HotelOrderQueryResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type HotelOrderQueryResponseBodyModuleHotelInfo struct {
	BrandCode  *string `json:"brand_code,omitempty" xml:"brand_code,omitempty"`
	BrandGroup *string `json:"brand_group,omitempty" xml:"brand_group,omitempty"`
	BrandName  *string `json:"brand_name,omitempty" xml:"brand_name,omitempty"`
	// example:
	//
	// 1669344020
	CheckIn *int64 `json:"check_in,omitempty" xml:"check_in,omitempty"`
	// example:
	//
	// 1669344020
	CheckOut     *int64  `json:"check_out,omitempty" xml:"check_out,omitempty"`
	City         *string `json:"city,omitempty" xml:"city,omitempty"`
	CityAdCode   *string `json:"city_ad_code,omitempty" xml:"city_ad_code,omitempty"`
	CountryCode  *string `json:"country_code,omitempty" xml:"country_code,omitempty"`
	CountryName  *string `json:"country_name,omitempty" xml:"country_name,omitempty"`
	HotelAddress *string `json:"hotel_address,omitempty" xml:"hotel_address,omitempty"`
	HotelName    *string `json:"hotel_name,omitempty" xml:"hotel_name,omitempty"`
	HotelPhone   *string `json:"hotel_phone,omitempty" xml:"hotel_phone,omitempty"`
	// example:
	//
	// 11
	HotelSupportVatInvoiceType *int32 `json:"hotel_support_vat_invoice_type,omitempty" xml:"hotel_support_vat_invoice_type,omitempty"`
	// example:
	//
	// 1
	Night *int32 `json:"night,omitempty" xml:"night,omitempty"`
	// example:
	//
	// 1
	RoomNum  *int32  `json:"room_num,omitempty" xml:"room_num,omitempty"`
	RoomType *string `json:"room_type,omitempty" xml:"room_type,omitempty"`
}

func (s HotelOrderQueryResponseBodyModuleHotelInfo) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderQueryResponseBodyModuleHotelInfo) GoString() string {
	return s.String()
}

func (s *HotelOrderQueryResponseBodyModuleHotelInfo) GetBrandCode() *string {
	return s.BrandCode
}

func (s *HotelOrderQueryResponseBodyModuleHotelInfo) GetBrandGroup() *string {
	return s.BrandGroup
}

func (s *HotelOrderQueryResponseBodyModuleHotelInfo) GetBrandName() *string {
	return s.BrandName
}

func (s *HotelOrderQueryResponseBodyModuleHotelInfo) GetCheckIn() *int64 {
	return s.CheckIn
}

func (s *HotelOrderQueryResponseBodyModuleHotelInfo) GetCheckOut() *int64 {
	return s.CheckOut
}

func (s *HotelOrderQueryResponseBodyModuleHotelInfo) GetCity() *string {
	return s.City
}

func (s *HotelOrderQueryResponseBodyModuleHotelInfo) GetCityAdCode() *string {
	return s.CityAdCode
}

func (s *HotelOrderQueryResponseBodyModuleHotelInfo) GetCountryCode() *string {
	return s.CountryCode
}

func (s *HotelOrderQueryResponseBodyModuleHotelInfo) GetCountryName() *string {
	return s.CountryName
}

func (s *HotelOrderQueryResponseBodyModuleHotelInfo) GetHotelAddress() *string {
	return s.HotelAddress
}

func (s *HotelOrderQueryResponseBodyModuleHotelInfo) GetHotelName() *string {
	return s.HotelName
}

func (s *HotelOrderQueryResponseBodyModuleHotelInfo) GetHotelPhone() *string {
	return s.HotelPhone
}

func (s *HotelOrderQueryResponseBodyModuleHotelInfo) GetHotelSupportVatInvoiceType() *int32 {
	return s.HotelSupportVatInvoiceType
}

func (s *HotelOrderQueryResponseBodyModuleHotelInfo) GetNight() *int32 {
	return s.Night
}

func (s *HotelOrderQueryResponseBodyModuleHotelInfo) GetRoomNum() *int32 {
	return s.RoomNum
}

func (s *HotelOrderQueryResponseBodyModuleHotelInfo) GetRoomType() *string {
	return s.RoomType
}

func (s *HotelOrderQueryResponseBodyModuleHotelInfo) SetBrandCode(v string) *HotelOrderQueryResponseBodyModuleHotelInfo {
	s.BrandCode = &v
	return s
}

func (s *HotelOrderQueryResponseBodyModuleHotelInfo) SetBrandGroup(v string) *HotelOrderQueryResponseBodyModuleHotelInfo {
	s.BrandGroup = &v
	return s
}

func (s *HotelOrderQueryResponseBodyModuleHotelInfo) SetBrandName(v string) *HotelOrderQueryResponseBodyModuleHotelInfo {
	s.BrandName = &v
	return s
}

func (s *HotelOrderQueryResponseBodyModuleHotelInfo) SetCheckIn(v int64) *HotelOrderQueryResponseBodyModuleHotelInfo {
	s.CheckIn = &v
	return s
}

func (s *HotelOrderQueryResponseBodyModuleHotelInfo) SetCheckOut(v int64) *HotelOrderQueryResponseBodyModuleHotelInfo {
	s.CheckOut = &v
	return s
}

func (s *HotelOrderQueryResponseBodyModuleHotelInfo) SetCity(v string) *HotelOrderQueryResponseBodyModuleHotelInfo {
	s.City = &v
	return s
}

func (s *HotelOrderQueryResponseBodyModuleHotelInfo) SetCityAdCode(v string) *HotelOrderQueryResponseBodyModuleHotelInfo {
	s.CityAdCode = &v
	return s
}

func (s *HotelOrderQueryResponseBodyModuleHotelInfo) SetCountryCode(v string) *HotelOrderQueryResponseBodyModuleHotelInfo {
	s.CountryCode = &v
	return s
}

func (s *HotelOrderQueryResponseBodyModuleHotelInfo) SetCountryName(v string) *HotelOrderQueryResponseBodyModuleHotelInfo {
	s.CountryName = &v
	return s
}

func (s *HotelOrderQueryResponseBodyModuleHotelInfo) SetHotelAddress(v string) *HotelOrderQueryResponseBodyModuleHotelInfo {
	s.HotelAddress = &v
	return s
}

func (s *HotelOrderQueryResponseBodyModuleHotelInfo) SetHotelName(v string) *HotelOrderQueryResponseBodyModuleHotelInfo {
	s.HotelName = &v
	return s
}

func (s *HotelOrderQueryResponseBodyModuleHotelInfo) SetHotelPhone(v string) *HotelOrderQueryResponseBodyModuleHotelInfo {
	s.HotelPhone = &v
	return s
}

func (s *HotelOrderQueryResponseBodyModuleHotelInfo) SetHotelSupportVatInvoiceType(v int32) *HotelOrderQueryResponseBodyModuleHotelInfo {
	s.HotelSupportVatInvoiceType = &v
	return s
}

func (s *HotelOrderQueryResponseBodyModuleHotelInfo) SetNight(v int32) *HotelOrderQueryResponseBodyModuleHotelInfo {
	s.Night = &v
	return s
}

func (s *HotelOrderQueryResponseBodyModuleHotelInfo) SetRoomNum(v int32) *HotelOrderQueryResponseBodyModuleHotelInfo {
	s.RoomNum = &v
	return s
}

func (s *HotelOrderQueryResponseBodyModuleHotelInfo) SetRoomType(v string) *HotelOrderQueryResponseBodyModuleHotelInfo {
	s.RoomType = &v
	return s
}

func (s *HotelOrderQueryResponseBodyModuleHotelInfo) Validate() error {
	return dara.Validate(s)
}

type HotelOrderQueryResponseBodyModuleInvoiceInfo struct {
	// example:
	//
	// 12345678
	Id    *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s HotelOrderQueryResponseBodyModuleInvoiceInfo) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderQueryResponseBodyModuleInvoiceInfo) GoString() string {
	return s.String()
}

func (s *HotelOrderQueryResponseBodyModuleInvoiceInfo) GetId() *int64 {
	return s.Id
}

func (s *HotelOrderQueryResponseBodyModuleInvoiceInfo) GetTitle() *string {
	return s.Title
}

func (s *HotelOrderQueryResponseBodyModuleInvoiceInfo) SetId(v int64) *HotelOrderQueryResponseBodyModuleInvoiceInfo {
	s.Id = &v
	return s
}

func (s *HotelOrderQueryResponseBodyModuleInvoiceInfo) SetTitle(v string) *HotelOrderQueryResponseBodyModuleInvoiceInfo {
	s.Title = &v
	return s
}

func (s *HotelOrderQueryResponseBodyModuleInvoiceInfo) Validate() error {
	return dara.Validate(s)
}

type HotelOrderQueryResponseBodyModuleOrderBaseInfo struct {
	// example:
	//
	// 12345678
	ApplyId    *string `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	BtripTitle *string `json:"btrip_title,omitempty" xml:"btrip_title,omitempty"`
	Category   *int32  `json:"category,omitempty" xml:"category,omitempty"`
	// example:
	//
	// iscorpId
	CorpId   *string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	CorpName *string `json:"corp_name,omitempty" xml:"corp_name,omitempty"`
	// example:
	//
	// 12345678
	DepartId       *string   `json:"depart_id,omitempty" xml:"depart_id,omitempty"`
	DepartName     *string   `json:"depart_name,omitempty" xml:"depart_name,omitempty"`
	ExceedApplyNos []*string `json:"exceed_apply_nos,omitempty" xml:"exceed_apply_nos,omitempty" type:"Repeated"`
	ExtendField    *string   `json:"extend_field,omitempty" xml:"extend_field,omitempty"`
	// example:
	//
	// 1669344020
	GmtCreate *int64 `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// example:
	//
	// 1669344020
	GmtModified *int64 `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// example:
	//
	// 1002145190081005400
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 12345678
	ItineraryId *string `json:"itinerary_id,omitempty" xml:"itinerary_id,omitempty"`
	// example:
	//
	// 1
	OrderStatus *int32 `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// example:
	//
	// 1
	OrderType *int32  `json:"order_type,omitempty" xml:"order_type,omitempty"`
	Supplier  *string `json:"supplier,omitempty" xml:"supplier,omitempty"`
	// example:
	//
	// 12345678
	ThirdpartApplyId    *string `json:"thirdpart_apply_id,omitempty" xml:"thirdpart_apply_id,omitempty"`
	ThirdpartBusinessId *string `json:"thirdpart_business_id,omitempty" xml:"thirdpart_business_id,omitempty"`
	// example:
	//
	// 12345678
	ThirdpartDepartId *string `json:"thirdpart_depart_id,omitempty" xml:"thirdpart_depart_id,omitempty"`
	// example:
	//
	// 12345678
	ThirdpartItineraryId *string `json:"thirdpart_itinerary_id,omitempty" xml:"thirdpart_itinerary_id,omitempty"`
	// example:
	//
	// 12345678
	UserId   *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserName *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s HotelOrderQueryResponseBodyModuleOrderBaseInfo) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderQueryResponseBodyModuleOrderBaseInfo) GoString() string {
	return s.String()
}

func (s *HotelOrderQueryResponseBodyModuleOrderBaseInfo) GetApplyId() *string {
	return s.ApplyId
}

func (s *HotelOrderQueryResponseBodyModuleOrderBaseInfo) GetBtripTitle() *string {
	return s.BtripTitle
}

func (s *HotelOrderQueryResponseBodyModuleOrderBaseInfo) GetCategory() *int32 {
	return s.Category
}

func (s *HotelOrderQueryResponseBodyModuleOrderBaseInfo) GetCorpId() *string {
	return s.CorpId
}

func (s *HotelOrderQueryResponseBodyModuleOrderBaseInfo) GetCorpName() *string {
	return s.CorpName
}

func (s *HotelOrderQueryResponseBodyModuleOrderBaseInfo) GetDepartId() *string {
	return s.DepartId
}

func (s *HotelOrderQueryResponseBodyModuleOrderBaseInfo) GetDepartName() *string {
	return s.DepartName
}

func (s *HotelOrderQueryResponseBodyModuleOrderBaseInfo) GetExceedApplyNos() []*string {
	return s.ExceedApplyNos
}

func (s *HotelOrderQueryResponseBodyModuleOrderBaseInfo) GetExtendField() *string {
	return s.ExtendField
}

func (s *HotelOrderQueryResponseBodyModuleOrderBaseInfo) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *HotelOrderQueryResponseBodyModuleOrderBaseInfo) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *HotelOrderQueryResponseBodyModuleOrderBaseInfo) GetId() *int64 {
	return s.Id
}

func (s *HotelOrderQueryResponseBodyModuleOrderBaseInfo) GetItineraryId() *string {
	return s.ItineraryId
}

func (s *HotelOrderQueryResponseBodyModuleOrderBaseInfo) GetOrderStatus() *int32 {
	return s.OrderStatus
}

func (s *HotelOrderQueryResponseBodyModuleOrderBaseInfo) GetOrderType() *int32 {
	return s.OrderType
}

func (s *HotelOrderQueryResponseBodyModuleOrderBaseInfo) GetSupplier() *string {
	return s.Supplier
}

func (s *HotelOrderQueryResponseBodyModuleOrderBaseInfo) GetThirdpartApplyId() *string {
	return s.ThirdpartApplyId
}

func (s *HotelOrderQueryResponseBodyModuleOrderBaseInfo) GetThirdpartBusinessId() *string {
	return s.ThirdpartBusinessId
}

func (s *HotelOrderQueryResponseBodyModuleOrderBaseInfo) GetThirdpartDepartId() *string {
	return s.ThirdpartDepartId
}

func (s *HotelOrderQueryResponseBodyModuleOrderBaseInfo) GetThirdpartItineraryId() *string {
	return s.ThirdpartItineraryId
}

func (s *HotelOrderQueryResponseBodyModuleOrderBaseInfo) GetUserId() *string {
	return s.UserId
}

func (s *HotelOrderQueryResponseBodyModuleOrderBaseInfo) GetUserName() *string {
	return s.UserName
}

func (s *HotelOrderQueryResponseBodyModuleOrderBaseInfo) SetApplyId(v string) *HotelOrderQueryResponseBodyModuleOrderBaseInfo {
	s.ApplyId = &v
	return s
}

func (s *HotelOrderQueryResponseBodyModuleOrderBaseInfo) SetBtripTitle(v string) *HotelOrderQueryResponseBodyModuleOrderBaseInfo {
	s.BtripTitle = &v
	return s
}

func (s *HotelOrderQueryResponseBodyModuleOrderBaseInfo) SetCategory(v int32) *HotelOrderQueryResponseBodyModuleOrderBaseInfo {
	s.Category = &v
	return s
}

func (s *HotelOrderQueryResponseBodyModuleOrderBaseInfo) SetCorpId(v string) *HotelOrderQueryResponseBodyModuleOrderBaseInfo {
	s.CorpId = &v
	return s
}

func (s *HotelOrderQueryResponseBodyModuleOrderBaseInfo) SetCorpName(v string) *HotelOrderQueryResponseBodyModuleOrderBaseInfo {
	s.CorpName = &v
	return s
}

func (s *HotelOrderQueryResponseBodyModuleOrderBaseInfo) SetDepartId(v string) *HotelOrderQueryResponseBodyModuleOrderBaseInfo {
	s.DepartId = &v
	return s
}

func (s *HotelOrderQueryResponseBodyModuleOrderBaseInfo) SetDepartName(v string) *HotelOrderQueryResponseBodyModuleOrderBaseInfo {
	s.DepartName = &v
	return s
}

func (s *HotelOrderQueryResponseBodyModuleOrderBaseInfo) SetExceedApplyNos(v []*string) *HotelOrderQueryResponseBodyModuleOrderBaseInfo {
	s.ExceedApplyNos = v
	return s
}

func (s *HotelOrderQueryResponseBodyModuleOrderBaseInfo) SetExtendField(v string) *HotelOrderQueryResponseBodyModuleOrderBaseInfo {
	s.ExtendField = &v
	return s
}

func (s *HotelOrderQueryResponseBodyModuleOrderBaseInfo) SetGmtCreate(v int64) *HotelOrderQueryResponseBodyModuleOrderBaseInfo {
	s.GmtCreate = &v
	return s
}

func (s *HotelOrderQueryResponseBodyModuleOrderBaseInfo) SetGmtModified(v int64) *HotelOrderQueryResponseBodyModuleOrderBaseInfo {
	s.GmtModified = &v
	return s
}

func (s *HotelOrderQueryResponseBodyModuleOrderBaseInfo) SetId(v int64) *HotelOrderQueryResponseBodyModuleOrderBaseInfo {
	s.Id = &v
	return s
}

func (s *HotelOrderQueryResponseBodyModuleOrderBaseInfo) SetItineraryId(v string) *HotelOrderQueryResponseBodyModuleOrderBaseInfo {
	s.ItineraryId = &v
	return s
}

func (s *HotelOrderQueryResponseBodyModuleOrderBaseInfo) SetOrderStatus(v int32) *HotelOrderQueryResponseBodyModuleOrderBaseInfo {
	s.OrderStatus = &v
	return s
}

func (s *HotelOrderQueryResponseBodyModuleOrderBaseInfo) SetOrderType(v int32) *HotelOrderQueryResponseBodyModuleOrderBaseInfo {
	s.OrderType = &v
	return s
}

func (s *HotelOrderQueryResponseBodyModuleOrderBaseInfo) SetSupplier(v string) *HotelOrderQueryResponseBodyModuleOrderBaseInfo {
	s.Supplier = &v
	return s
}

func (s *HotelOrderQueryResponseBodyModuleOrderBaseInfo) SetThirdpartApplyId(v string) *HotelOrderQueryResponseBodyModuleOrderBaseInfo {
	s.ThirdpartApplyId = &v
	return s
}

func (s *HotelOrderQueryResponseBodyModuleOrderBaseInfo) SetThirdpartBusinessId(v string) *HotelOrderQueryResponseBodyModuleOrderBaseInfo {
	s.ThirdpartBusinessId = &v
	return s
}

func (s *HotelOrderQueryResponseBodyModuleOrderBaseInfo) SetThirdpartDepartId(v string) *HotelOrderQueryResponseBodyModuleOrderBaseInfo {
	s.ThirdpartDepartId = &v
	return s
}

func (s *HotelOrderQueryResponseBodyModuleOrderBaseInfo) SetThirdpartItineraryId(v string) *HotelOrderQueryResponseBodyModuleOrderBaseInfo {
	s.ThirdpartItineraryId = &v
	return s
}

func (s *HotelOrderQueryResponseBodyModuleOrderBaseInfo) SetUserId(v string) *HotelOrderQueryResponseBodyModuleOrderBaseInfo {
	s.UserId = &v
	return s
}

func (s *HotelOrderQueryResponseBodyModuleOrderBaseInfo) SetUserName(v string) *HotelOrderQueryResponseBodyModuleOrderBaseInfo {
	s.UserName = &v
	return s
}

func (s *HotelOrderQueryResponseBodyModuleOrderBaseInfo) Validate() error {
	return dara.Validate(s)
}

type HotelOrderQueryResponseBodyModulePassengerList struct {
	ApplyId *string `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	// example:
	//
	// 12345678
	CostCenterId   *int64  `json:"cost_center_id,omitempty" xml:"cost_center_id,omitempty"`
	CostCenterName *string `json:"cost_center_name,omitempty" xml:"cost_center_name,omitempty"`
	// example:
	//
	// 12345678
	CostCenterNumber *string `json:"cost_center_number,omitempty" xml:"cost_center_number,omitempty"`
	ItineraryId      *string `json:"itinerary_id,omitempty" xml:"itinerary_id,omitempty"`
	OccupantType     *int32  `json:"occupant_type,omitempty" xml:"occupant_type,omitempty"`
	ProjectCode      *string `json:"project_code,omitempty" xml:"project_code,omitempty"`
	// example:
	//
	// 12345678
	ProjectId        *int64  `json:"project_id,omitempty" xml:"project_id,omitempty"`
	ProjectTitle     *string `json:"project_title,omitempty" xml:"project_title,omitempty"`
	ThirdpartApplyId *string `json:"thirdpart_apply_id,omitempty" xml:"thirdpart_apply_id,omitempty"`
	// example:
	//
	// 12345678
	ThirdpartCostCenterId *string `json:"thirdpart_cost_center_id,omitempty" xml:"thirdpart_cost_center_id,omitempty"`
	// example:
	//
	// 12345678
	ThirdpartProjectId *string `json:"thirdpart_project_id,omitempty" xml:"thirdpart_project_id,omitempty"`
	// example:
	//
	// 12345678
	UserId   *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserName *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
	// example:
	//
	// 0
	UserType *int32 `json:"user_type,omitempty" xml:"user_type,omitempty"`
}

func (s HotelOrderQueryResponseBodyModulePassengerList) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderQueryResponseBodyModulePassengerList) GoString() string {
	return s.String()
}

func (s *HotelOrderQueryResponseBodyModulePassengerList) GetApplyId() *string {
	return s.ApplyId
}

func (s *HotelOrderQueryResponseBodyModulePassengerList) GetCostCenterId() *int64 {
	return s.CostCenterId
}

func (s *HotelOrderQueryResponseBodyModulePassengerList) GetCostCenterName() *string {
	return s.CostCenterName
}

func (s *HotelOrderQueryResponseBodyModulePassengerList) GetCostCenterNumber() *string {
	return s.CostCenterNumber
}

func (s *HotelOrderQueryResponseBodyModulePassengerList) GetItineraryId() *string {
	return s.ItineraryId
}

func (s *HotelOrderQueryResponseBodyModulePassengerList) GetOccupantType() *int32 {
	return s.OccupantType
}

func (s *HotelOrderQueryResponseBodyModulePassengerList) GetProjectCode() *string {
	return s.ProjectCode
}

func (s *HotelOrderQueryResponseBodyModulePassengerList) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *HotelOrderQueryResponseBodyModulePassengerList) GetProjectTitle() *string {
	return s.ProjectTitle
}

func (s *HotelOrderQueryResponseBodyModulePassengerList) GetThirdpartApplyId() *string {
	return s.ThirdpartApplyId
}

func (s *HotelOrderQueryResponseBodyModulePassengerList) GetThirdpartCostCenterId() *string {
	return s.ThirdpartCostCenterId
}

func (s *HotelOrderQueryResponseBodyModulePassengerList) GetThirdpartProjectId() *string {
	return s.ThirdpartProjectId
}

func (s *HotelOrderQueryResponseBodyModulePassengerList) GetUserId() *string {
	return s.UserId
}

func (s *HotelOrderQueryResponseBodyModulePassengerList) GetUserName() *string {
	return s.UserName
}

func (s *HotelOrderQueryResponseBodyModulePassengerList) GetUserType() *int32 {
	return s.UserType
}

func (s *HotelOrderQueryResponseBodyModulePassengerList) SetApplyId(v string) *HotelOrderQueryResponseBodyModulePassengerList {
	s.ApplyId = &v
	return s
}

func (s *HotelOrderQueryResponseBodyModulePassengerList) SetCostCenterId(v int64) *HotelOrderQueryResponseBodyModulePassengerList {
	s.CostCenterId = &v
	return s
}

func (s *HotelOrderQueryResponseBodyModulePassengerList) SetCostCenterName(v string) *HotelOrderQueryResponseBodyModulePassengerList {
	s.CostCenterName = &v
	return s
}

func (s *HotelOrderQueryResponseBodyModulePassengerList) SetCostCenterNumber(v string) *HotelOrderQueryResponseBodyModulePassengerList {
	s.CostCenterNumber = &v
	return s
}

func (s *HotelOrderQueryResponseBodyModulePassengerList) SetItineraryId(v string) *HotelOrderQueryResponseBodyModulePassengerList {
	s.ItineraryId = &v
	return s
}

func (s *HotelOrderQueryResponseBodyModulePassengerList) SetOccupantType(v int32) *HotelOrderQueryResponseBodyModulePassengerList {
	s.OccupantType = &v
	return s
}

func (s *HotelOrderQueryResponseBodyModulePassengerList) SetProjectCode(v string) *HotelOrderQueryResponseBodyModulePassengerList {
	s.ProjectCode = &v
	return s
}

func (s *HotelOrderQueryResponseBodyModulePassengerList) SetProjectId(v int64) *HotelOrderQueryResponseBodyModulePassengerList {
	s.ProjectId = &v
	return s
}

func (s *HotelOrderQueryResponseBodyModulePassengerList) SetProjectTitle(v string) *HotelOrderQueryResponseBodyModulePassengerList {
	s.ProjectTitle = &v
	return s
}

func (s *HotelOrderQueryResponseBodyModulePassengerList) SetThirdpartApplyId(v string) *HotelOrderQueryResponseBodyModulePassengerList {
	s.ThirdpartApplyId = &v
	return s
}

func (s *HotelOrderQueryResponseBodyModulePassengerList) SetThirdpartCostCenterId(v string) *HotelOrderQueryResponseBodyModulePassengerList {
	s.ThirdpartCostCenterId = &v
	return s
}

func (s *HotelOrderQueryResponseBodyModulePassengerList) SetThirdpartProjectId(v string) *HotelOrderQueryResponseBodyModulePassengerList {
	s.ThirdpartProjectId = &v
	return s
}

func (s *HotelOrderQueryResponseBodyModulePassengerList) SetUserId(v string) *HotelOrderQueryResponseBodyModulePassengerList {
	s.UserId = &v
	return s
}

func (s *HotelOrderQueryResponseBodyModulePassengerList) SetUserName(v string) *HotelOrderQueryResponseBodyModulePassengerList {
	s.UserName = &v
	return s
}

func (s *HotelOrderQueryResponseBodyModulePassengerList) SetUserType(v int32) *HotelOrderQueryResponseBodyModulePassengerList {
	s.UserType = &v
	return s
}

func (s *HotelOrderQueryResponseBodyModulePassengerList) Validate() error {
	return dara.Validate(s)
}

type HotelOrderQueryResponseBodyModulePriceInfoList struct {
	// example:
	//
	// 1
	CategoryCode *int32 `json:"category_code,omitempty" xml:"category_code,omitempty"`
	// example:
	//
	// 1669344020
	GmtCreate *int64 `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// example:
	//
	// 1
	PayType *int32 `json:"pay_type,omitempty" xml:"pay_type,omitempty"`
	// example:
	//
	// 200
	Price *float64 `json:"price,omitempty" xml:"price,omitempty"`
	// example:
	//
	// 12345678910987654321
	TradeId *string `json:"trade_id,omitempty" xml:"trade_id,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s HotelOrderQueryResponseBodyModulePriceInfoList) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderQueryResponseBodyModulePriceInfoList) GoString() string {
	return s.String()
}

func (s *HotelOrderQueryResponseBodyModulePriceInfoList) GetCategoryCode() *int32 {
	return s.CategoryCode
}

func (s *HotelOrderQueryResponseBodyModulePriceInfoList) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *HotelOrderQueryResponseBodyModulePriceInfoList) GetPayType() *int32 {
	return s.PayType
}

func (s *HotelOrderQueryResponseBodyModulePriceInfoList) GetPrice() *float64 {
	return s.Price
}

func (s *HotelOrderQueryResponseBodyModulePriceInfoList) GetTradeId() *string {
	return s.TradeId
}

func (s *HotelOrderQueryResponseBodyModulePriceInfoList) GetType() *int32 {
	return s.Type
}

func (s *HotelOrderQueryResponseBodyModulePriceInfoList) SetCategoryCode(v int32) *HotelOrderQueryResponseBodyModulePriceInfoList {
	s.CategoryCode = &v
	return s
}

func (s *HotelOrderQueryResponseBodyModulePriceInfoList) SetGmtCreate(v int64) *HotelOrderQueryResponseBodyModulePriceInfoList {
	s.GmtCreate = &v
	return s
}

func (s *HotelOrderQueryResponseBodyModulePriceInfoList) SetPayType(v int32) *HotelOrderQueryResponseBodyModulePriceInfoList {
	s.PayType = &v
	return s
}

func (s *HotelOrderQueryResponseBodyModulePriceInfoList) SetPrice(v float64) *HotelOrderQueryResponseBodyModulePriceInfoList {
	s.Price = &v
	return s
}

func (s *HotelOrderQueryResponseBodyModulePriceInfoList) SetTradeId(v string) *HotelOrderQueryResponseBodyModulePriceInfoList {
	s.TradeId = &v
	return s
}

func (s *HotelOrderQueryResponseBodyModulePriceInfoList) SetType(v int32) *HotelOrderQueryResponseBodyModulePriceInfoList {
	s.Type = &v
	return s
}

func (s *HotelOrderQueryResponseBodyModulePriceInfoList) Validate() error {
	return dara.Validate(s)
}
