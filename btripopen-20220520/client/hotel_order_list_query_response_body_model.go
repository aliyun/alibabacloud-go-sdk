// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelOrderListQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *HotelOrderListQueryResponseBody
	GetCode() *string
	SetMessage(v string) *HotelOrderListQueryResponseBody
	GetMessage() *string
	SetModule(v []*HotelOrderListQueryResponseBodyModule) *HotelOrderListQueryResponseBody
	GetModule() []*HotelOrderListQueryResponseBodyModule
	SetPageInfo(v *HotelOrderListQueryResponseBodyPageInfo) *HotelOrderListQueryResponseBody
	GetPageInfo() *HotelOrderListQueryResponseBodyPageInfo
	SetRequestId(v string) *HotelOrderListQueryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *HotelOrderListQueryResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *HotelOrderListQueryResponseBody
	GetTraceId() *string
}

type HotelOrderListQueryResponseBody struct {
	// example:
	//
	// SUCCESS
	Code     *string                                  `json:"code,omitempty" xml:"code,omitempty"`
	Message  *string                                  `json:"message,omitempty" xml:"message,omitempty"`
	Module   []*HotelOrderListQueryResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Repeated"`
	PageInfo *HotelOrderListQueryResponseBodyPageInfo `json:"page_info,omitempty" xml:"page_info,omitempty" type:"Struct"`
	// example:
	//
	// C61ECFF6-606B-5F66-B81D-D77369043A5F
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

func (s HotelOrderListQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderListQueryResponseBody) GoString() string {
	return s.String()
}

func (s *HotelOrderListQueryResponseBody) GetCode() *string {
	return s.Code
}

func (s *HotelOrderListQueryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *HotelOrderListQueryResponseBody) GetModule() []*HotelOrderListQueryResponseBodyModule {
	return s.Module
}

func (s *HotelOrderListQueryResponseBody) GetPageInfo() *HotelOrderListQueryResponseBodyPageInfo {
	return s.PageInfo
}

func (s *HotelOrderListQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *HotelOrderListQueryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *HotelOrderListQueryResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *HotelOrderListQueryResponseBody) SetCode(v string) *HotelOrderListQueryResponseBody {
	s.Code = &v
	return s
}

func (s *HotelOrderListQueryResponseBody) SetMessage(v string) *HotelOrderListQueryResponseBody {
	s.Message = &v
	return s
}

func (s *HotelOrderListQueryResponseBody) SetModule(v []*HotelOrderListQueryResponseBodyModule) *HotelOrderListQueryResponseBody {
	s.Module = v
	return s
}

func (s *HotelOrderListQueryResponseBody) SetPageInfo(v *HotelOrderListQueryResponseBodyPageInfo) *HotelOrderListQueryResponseBody {
	s.PageInfo = v
	return s
}

func (s *HotelOrderListQueryResponseBody) SetRequestId(v string) *HotelOrderListQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *HotelOrderListQueryResponseBody) SetSuccess(v bool) *HotelOrderListQueryResponseBody {
	s.Success = &v
	return s
}

func (s *HotelOrderListQueryResponseBody) SetTraceId(v string) *HotelOrderListQueryResponseBody {
	s.TraceId = &v
	return s
}

func (s *HotelOrderListQueryResponseBody) Validate() error {
	return dara.Validate(s)
}

type HotelOrderListQueryResponseBodyModule struct {
	// example:
	//
	// 22678
	ApplyId    *int64  `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	BtripTitle *string `json:"btrip_title,omitempty" xml:"btrip_title,omitempty"`
	Category   *int32  `json:"category,omitempty" xml:"category,omitempty"`
	// example:
	//
	// 2022-05-15T22:27Z
	CheckIn *string `json:"check_in,omitempty" xml:"check_in,omitempty"`
	// example:
	//
	// 2022-05-15T22:27Z
	CheckOut    *string                                          `json:"check_out,omitempty" xml:"check_out,omitempty"`
	City        *string                                          `json:"city,omitempty" xml:"city,omitempty"`
	CityAdCode  *string                                          `json:"city_ad_code,omitempty" xml:"city_ad_code,omitempty"`
	ContactName *string                                          `json:"contact_name,omitempty" xml:"contact_name,omitempty"`
	CorpId      *string                                          `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	CorpName    *string                                          `json:"corp_name,omitempty" xml:"corp_name,omitempty"`
	CostCenter  *HotelOrderListQueryResponseBodyModuleCostCenter `json:"cost_center,omitempty" xml:"cost_center,omitempty" type:"Struct"`
	CountryCode *string                                          `json:"country_code,omitempty" xml:"country_code,omitempty"`
	CountryName *string                                          `json:"country_name,omitempty" xml:"country_name,omitempty"`
	DepartId    *string                                          `json:"depart_id,omitempty" xml:"depart_id,omitempty"`
	DepartName  *string                                          `json:"depart_name,omitempty" xml:"depart_name,omitempty"`
	ExtendField *string                                          `json:"extend_field,omitempty" xml:"extend_field,omitempty"`
	// example:
	//
	// 2022-05-15T22:27Z
	GmtCreate *string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// example:
	//
	// 2022-05-15T22:27Z
	GmtModified *string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	Guest       *string `json:"guest,omitempty" xml:"guest,omitempty"`
	HotelName   *string `json:"hotel_name,omitempty" xml:"hotel_name,omitempty"`
	// example:
	//
	// 11
	HotelSupportVatInvoiceType *int32 `json:"hotel_support_vat_invoice_type,omitempty" xml:"hotel_support_vat_invoice_type,omitempty"`
	// example:
	//
	// 13764
	Id      *int64                                        `json:"id,omitempty" xml:"id,omitempty"`
	Invoice *HotelOrderListQueryResponseBodyModuleInvoice `json:"invoice,omitempty" xml:"invoice,omitempty" type:"Struct"`
	// example:
	//
	// 4
	Night *int32 `json:"night,omitempty" xml:"night,omitempty"`
	// example:
	//
	// 1
	OrderStatus     *int32  `json:"order_status,omitempty" xml:"order_status,omitempty"`
	OrderStatusDesc *string `json:"order_status_desc,omitempty" xml:"order_status_desc,omitempty"`
	// example:
	//
	// 1
	OrderType     *int32                                                `json:"order_type,omitempty" xml:"order_type,omitempty"`
	OrderTypeDesc *string                                               `json:"order_type_desc,omitempty" xml:"order_type_desc,omitempty"`
	PriceInfoList []*HotelOrderListQueryResponseBodyModulePriceInfoList `json:"price_info_list,omitempty" xml:"price_info_list,omitempty" type:"Repeated"`
	// example:
	//
	// CS-PROJECTCODE
	ProjectCode *string `json:"project_code,omitempty" xml:"project_code,omitempty"`
	// example:
	//
	// 13631
	ProjectId    *int64  `json:"project_id,omitempty" xml:"project_id,omitempty"`
	ProjectTitle *string `json:"project_title,omitempty" xml:"project_title,omitempty"`
	// example:
	//
	// 4
	RoomNum  *int32  `json:"room_num,omitempty" xml:"room_num,omitempty"`
	RoomType *string `json:"room_type,omitempty" xml:"room_type,omitempty"`
	Supplier *string `json:"supplier,omitempty" xml:"supplier,omitempty"`
	// example:
	//
	// CS-THIRDAPPLY
	ThirdpartApplyId    *string `json:"thirdpart_apply_id,omitempty" xml:"thirdpart_apply_id,omitempty"`
	ThirdpartBusinessId *string `json:"thirdpart_business_id,omitempty" xml:"thirdpart_business_id,omitempty"`
	// example:
	//
	// CS-ITINEARY
	ThirdpartItineraryId *string `json:"thirdpart_itinerary_id,omitempty" xml:"thirdpart_itinerary_id,omitempty"`
	// example:
	//
	// CS-THIRDPROJECT
	ThirdpartProjectId *string                                                   `json:"thirdpart_project_id,omitempty" xml:"thirdpart_project_id,omitempty"`
	UserAffiliateList  []*HotelOrderListQueryResponseBodyModuleUserAffiliateList `json:"user_affiliate_list,omitempty" xml:"user_affiliate_list,omitempty" type:"Repeated"`
	UserId             *string                                                   `json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserName           *string                                                   `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s HotelOrderListQueryResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderListQueryResponseBodyModule) GoString() string {
	return s.String()
}

func (s *HotelOrderListQueryResponseBodyModule) GetApplyId() *int64 {
	return s.ApplyId
}

func (s *HotelOrderListQueryResponseBodyModule) GetBtripTitle() *string {
	return s.BtripTitle
}

func (s *HotelOrderListQueryResponseBodyModule) GetCategory() *int32 {
	return s.Category
}

func (s *HotelOrderListQueryResponseBodyModule) GetCheckIn() *string {
	return s.CheckIn
}

func (s *HotelOrderListQueryResponseBodyModule) GetCheckOut() *string {
	return s.CheckOut
}

func (s *HotelOrderListQueryResponseBodyModule) GetCity() *string {
	return s.City
}

func (s *HotelOrderListQueryResponseBodyModule) GetCityAdCode() *string {
	return s.CityAdCode
}

func (s *HotelOrderListQueryResponseBodyModule) GetContactName() *string {
	return s.ContactName
}

func (s *HotelOrderListQueryResponseBodyModule) GetCorpId() *string {
	return s.CorpId
}

func (s *HotelOrderListQueryResponseBodyModule) GetCorpName() *string {
	return s.CorpName
}

func (s *HotelOrderListQueryResponseBodyModule) GetCostCenter() *HotelOrderListQueryResponseBodyModuleCostCenter {
	return s.CostCenter
}

func (s *HotelOrderListQueryResponseBodyModule) GetCountryCode() *string {
	return s.CountryCode
}

func (s *HotelOrderListQueryResponseBodyModule) GetCountryName() *string {
	return s.CountryName
}

func (s *HotelOrderListQueryResponseBodyModule) GetDepartId() *string {
	return s.DepartId
}

func (s *HotelOrderListQueryResponseBodyModule) GetDepartName() *string {
	return s.DepartName
}

func (s *HotelOrderListQueryResponseBodyModule) GetExtendField() *string {
	return s.ExtendField
}

func (s *HotelOrderListQueryResponseBodyModule) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *HotelOrderListQueryResponseBodyModule) GetGmtModified() *string {
	return s.GmtModified
}

func (s *HotelOrderListQueryResponseBodyModule) GetGuest() *string {
	return s.Guest
}

func (s *HotelOrderListQueryResponseBodyModule) GetHotelName() *string {
	return s.HotelName
}

func (s *HotelOrderListQueryResponseBodyModule) GetHotelSupportVatInvoiceType() *int32 {
	return s.HotelSupportVatInvoiceType
}

func (s *HotelOrderListQueryResponseBodyModule) GetId() *int64 {
	return s.Id
}

func (s *HotelOrderListQueryResponseBodyModule) GetInvoice() *HotelOrderListQueryResponseBodyModuleInvoice {
	return s.Invoice
}

func (s *HotelOrderListQueryResponseBodyModule) GetNight() *int32 {
	return s.Night
}

func (s *HotelOrderListQueryResponseBodyModule) GetOrderStatus() *int32 {
	return s.OrderStatus
}

func (s *HotelOrderListQueryResponseBodyModule) GetOrderStatusDesc() *string {
	return s.OrderStatusDesc
}

func (s *HotelOrderListQueryResponseBodyModule) GetOrderType() *int32 {
	return s.OrderType
}

func (s *HotelOrderListQueryResponseBodyModule) GetOrderTypeDesc() *string {
	return s.OrderTypeDesc
}

func (s *HotelOrderListQueryResponseBodyModule) GetPriceInfoList() []*HotelOrderListQueryResponseBodyModulePriceInfoList {
	return s.PriceInfoList
}

func (s *HotelOrderListQueryResponseBodyModule) GetProjectCode() *string {
	return s.ProjectCode
}

func (s *HotelOrderListQueryResponseBodyModule) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *HotelOrderListQueryResponseBodyModule) GetProjectTitle() *string {
	return s.ProjectTitle
}

func (s *HotelOrderListQueryResponseBodyModule) GetRoomNum() *int32 {
	return s.RoomNum
}

func (s *HotelOrderListQueryResponseBodyModule) GetRoomType() *string {
	return s.RoomType
}

func (s *HotelOrderListQueryResponseBodyModule) GetSupplier() *string {
	return s.Supplier
}

func (s *HotelOrderListQueryResponseBodyModule) GetThirdpartApplyId() *string {
	return s.ThirdpartApplyId
}

func (s *HotelOrderListQueryResponseBodyModule) GetThirdpartBusinessId() *string {
	return s.ThirdpartBusinessId
}

func (s *HotelOrderListQueryResponseBodyModule) GetThirdpartItineraryId() *string {
	return s.ThirdpartItineraryId
}

func (s *HotelOrderListQueryResponseBodyModule) GetThirdpartProjectId() *string {
	return s.ThirdpartProjectId
}

func (s *HotelOrderListQueryResponseBodyModule) GetUserAffiliateList() []*HotelOrderListQueryResponseBodyModuleUserAffiliateList {
	return s.UserAffiliateList
}

func (s *HotelOrderListQueryResponseBodyModule) GetUserId() *string {
	return s.UserId
}

func (s *HotelOrderListQueryResponseBodyModule) GetUserName() *string {
	return s.UserName
}

func (s *HotelOrderListQueryResponseBodyModule) SetApplyId(v int64) *HotelOrderListQueryResponseBodyModule {
	s.ApplyId = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) SetBtripTitle(v string) *HotelOrderListQueryResponseBodyModule {
	s.BtripTitle = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) SetCategory(v int32) *HotelOrderListQueryResponseBodyModule {
	s.Category = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) SetCheckIn(v string) *HotelOrderListQueryResponseBodyModule {
	s.CheckIn = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) SetCheckOut(v string) *HotelOrderListQueryResponseBodyModule {
	s.CheckOut = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) SetCity(v string) *HotelOrderListQueryResponseBodyModule {
	s.City = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) SetCityAdCode(v string) *HotelOrderListQueryResponseBodyModule {
	s.CityAdCode = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) SetContactName(v string) *HotelOrderListQueryResponseBodyModule {
	s.ContactName = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) SetCorpId(v string) *HotelOrderListQueryResponseBodyModule {
	s.CorpId = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) SetCorpName(v string) *HotelOrderListQueryResponseBodyModule {
	s.CorpName = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) SetCostCenter(v *HotelOrderListQueryResponseBodyModuleCostCenter) *HotelOrderListQueryResponseBodyModule {
	s.CostCenter = v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) SetCountryCode(v string) *HotelOrderListQueryResponseBodyModule {
	s.CountryCode = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) SetCountryName(v string) *HotelOrderListQueryResponseBodyModule {
	s.CountryName = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) SetDepartId(v string) *HotelOrderListQueryResponseBodyModule {
	s.DepartId = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) SetDepartName(v string) *HotelOrderListQueryResponseBodyModule {
	s.DepartName = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) SetExtendField(v string) *HotelOrderListQueryResponseBodyModule {
	s.ExtendField = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) SetGmtCreate(v string) *HotelOrderListQueryResponseBodyModule {
	s.GmtCreate = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) SetGmtModified(v string) *HotelOrderListQueryResponseBodyModule {
	s.GmtModified = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) SetGuest(v string) *HotelOrderListQueryResponseBodyModule {
	s.Guest = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) SetHotelName(v string) *HotelOrderListQueryResponseBodyModule {
	s.HotelName = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) SetHotelSupportVatInvoiceType(v int32) *HotelOrderListQueryResponseBodyModule {
	s.HotelSupportVatInvoiceType = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) SetId(v int64) *HotelOrderListQueryResponseBodyModule {
	s.Id = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) SetInvoice(v *HotelOrderListQueryResponseBodyModuleInvoice) *HotelOrderListQueryResponseBodyModule {
	s.Invoice = v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) SetNight(v int32) *HotelOrderListQueryResponseBodyModule {
	s.Night = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) SetOrderStatus(v int32) *HotelOrderListQueryResponseBodyModule {
	s.OrderStatus = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) SetOrderStatusDesc(v string) *HotelOrderListQueryResponseBodyModule {
	s.OrderStatusDesc = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) SetOrderType(v int32) *HotelOrderListQueryResponseBodyModule {
	s.OrderType = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) SetOrderTypeDesc(v string) *HotelOrderListQueryResponseBodyModule {
	s.OrderTypeDesc = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) SetPriceInfoList(v []*HotelOrderListQueryResponseBodyModulePriceInfoList) *HotelOrderListQueryResponseBodyModule {
	s.PriceInfoList = v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) SetProjectCode(v string) *HotelOrderListQueryResponseBodyModule {
	s.ProjectCode = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) SetProjectId(v int64) *HotelOrderListQueryResponseBodyModule {
	s.ProjectId = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) SetProjectTitle(v string) *HotelOrderListQueryResponseBodyModule {
	s.ProjectTitle = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) SetRoomNum(v int32) *HotelOrderListQueryResponseBodyModule {
	s.RoomNum = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) SetRoomType(v string) *HotelOrderListQueryResponseBodyModule {
	s.RoomType = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) SetSupplier(v string) *HotelOrderListQueryResponseBodyModule {
	s.Supplier = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) SetThirdpartApplyId(v string) *HotelOrderListQueryResponseBodyModule {
	s.ThirdpartApplyId = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) SetThirdpartBusinessId(v string) *HotelOrderListQueryResponseBodyModule {
	s.ThirdpartBusinessId = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) SetThirdpartItineraryId(v string) *HotelOrderListQueryResponseBodyModule {
	s.ThirdpartItineraryId = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) SetThirdpartProjectId(v string) *HotelOrderListQueryResponseBodyModule {
	s.ThirdpartProjectId = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) SetUserAffiliateList(v []*HotelOrderListQueryResponseBodyModuleUserAffiliateList) *HotelOrderListQueryResponseBodyModule {
	s.UserAffiliateList = v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) SetUserId(v string) *HotelOrderListQueryResponseBodyModule {
	s.UserId = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) SetUserName(v string) *HotelOrderListQueryResponseBodyModule {
	s.UserName = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type HotelOrderListQueryResponseBodyModuleCostCenter struct {
	CorpId *string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	// example:
	//
	// 14668
	Id   *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// CS-PNUY
	Number *string `json:"number,omitempty" xml:"number,omitempty"`
}

func (s HotelOrderListQueryResponseBodyModuleCostCenter) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderListQueryResponseBodyModuleCostCenter) GoString() string {
	return s.String()
}

func (s *HotelOrderListQueryResponseBodyModuleCostCenter) GetCorpId() *string {
	return s.CorpId
}

func (s *HotelOrderListQueryResponseBodyModuleCostCenter) GetId() *int64 {
	return s.Id
}

func (s *HotelOrderListQueryResponseBodyModuleCostCenter) GetName() *string {
	return s.Name
}

func (s *HotelOrderListQueryResponseBodyModuleCostCenter) GetNumber() *string {
	return s.Number
}

func (s *HotelOrderListQueryResponseBodyModuleCostCenter) SetCorpId(v string) *HotelOrderListQueryResponseBodyModuleCostCenter {
	s.CorpId = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModuleCostCenter) SetId(v int64) *HotelOrderListQueryResponseBodyModuleCostCenter {
	s.Id = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModuleCostCenter) SetName(v string) *HotelOrderListQueryResponseBodyModuleCostCenter {
	s.Name = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModuleCostCenter) SetNumber(v string) *HotelOrderListQueryResponseBodyModuleCostCenter {
	s.Number = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModuleCostCenter) Validate() error {
	return dara.Validate(s)
}

type HotelOrderListQueryResponseBodyModuleInvoice struct {
	// example:
	//
	// 133568
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 1
	InvoiceType *int32  `json:"invoice_type,omitempty" xml:"invoice_type,omitempty"`
	Title       *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s HotelOrderListQueryResponseBodyModuleInvoice) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderListQueryResponseBodyModuleInvoice) GoString() string {
	return s.String()
}

func (s *HotelOrderListQueryResponseBodyModuleInvoice) GetId() *int64 {
	return s.Id
}

func (s *HotelOrderListQueryResponseBodyModuleInvoice) GetInvoiceType() *int32 {
	return s.InvoiceType
}

func (s *HotelOrderListQueryResponseBodyModuleInvoice) GetTitle() *string {
	return s.Title
}

func (s *HotelOrderListQueryResponseBodyModuleInvoice) SetId(v int64) *HotelOrderListQueryResponseBodyModuleInvoice {
	s.Id = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModuleInvoice) SetInvoiceType(v int32) *HotelOrderListQueryResponseBodyModuleInvoice {
	s.InvoiceType = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModuleInvoice) SetTitle(v string) *HotelOrderListQueryResponseBodyModuleInvoice {
	s.Title = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModuleInvoice) Validate() error {
	return dara.Validate(s)
}

type HotelOrderListQueryResponseBodyModulePriceInfoList struct {
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
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
	// 2022-05-15T22:27Z
	GmtCreate     *string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	PassengerName *string `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
	// example:
	//
	// 1
	PayType *int32 `json:"pay_type,omitempty" xml:"pay_type,omitempty"`
	// example:
	//
	// 100
	Price *float64 `json:"price,omitempty" xml:"price,omitempty"`
	// example:
	//
	// cs1546728
	TradeId *string `json:"trade_id,omitempty" xml:"trade_id,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s HotelOrderListQueryResponseBodyModulePriceInfoList) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderListQueryResponseBodyModulePriceInfoList) GoString() string {
	return s.String()
}

func (s *HotelOrderListQueryResponseBodyModulePriceInfoList) GetCategory() *string {
	return s.Category
}

func (s *HotelOrderListQueryResponseBodyModulePriceInfoList) GetCategoryCode() *int32 {
	return s.CategoryCode
}

func (s *HotelOrderListQueryResponseBodyModulePriceInfoList) GetCategoryType() *int32 {
	return s.CategoryType
}

func (s *HotelOrderListQueryResponseBodyModulePriceInfoList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *HotelOrderListQueryResponseBodyModulePriceInfoList) GetPassengerName() *string {
	return s.PassengerName
}

func (s *HotelOrderListQueryResponseBodyModulePriceInfoList) GetPayType() *int32 {
	return s.PayType
}

func (s *HotelOrderListQueryResponseBodyModulePriceInfoList) GetPrice() *float64 {
	return s.Price
}

func (s *HotelOrderListQueryResponseBodyModulePriceInfoList) GetTradeId() *string {
	return s.TradeId
}

func (s *HotelOrderListQueryResponseBodyModulePriceInfoList) GetType() *int32 {
	return s.Type
}

func (s *HotelOrderListQueryResponseBodyModulePriceInfoList) SetCategory(v string) *HotelOrderListQueryResponseBodyModulePriceInfoList {
	s.Category = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModulePriceInfoList) SetCategoryCode(v int32) *HotelOrderListQueryResponseBodyModulePriceInfoList {
	s.CategoryCode = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModulePriceInfoList) SetCategoryType(v int32) *HotelOrderListQueryResponseBodyModulePriceInfoList {
	s.CategoryType = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModulePriceInfoList) SetGmtCreate(v string) *HotelOrderListQueryResponseBodyModulePriceInfoList {
	s.GmtCreate = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModulePriceInfoList) SetPassengerName(v string) *HotelOrderListQueryResponseBodyModulePriceInfoList {
	s.PassengerName = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModulePriceInfoList) SetPayType(v int32) *HotelOrderListQueryResponseBodyModulePriceInfoList {
	s.PayType = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModulePriceInfoList) SetPrice(v float64) *HotelOrderListQueryResponseBodyModulePriceInfoList {
	s.Price = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModulePriceInfoList) SetTradeId(v string) *HotelOrderListQueryResponseBodyModulePriceInfoList {
	s.TradeId = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModulePriceInfoList) SetType(v int32) *HotelOrderListQueryResponseBodyModulePriceInfoList {
	s.Type = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModulePriceInfoList) Validate() error {
	return dara.Validate(s)
}

type HotelOrderListQueryResponseBodyModuleUserAffiliateList struct {
	UserId   *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserName *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s HotelOrderListQueryResponseBodyModuleUserAffiliateList) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderListQueryResponseBodyModuleUserAffiliateList) GoString() string {
	return s.String()
}

func (s *HotelOrderListQueryResponseBodyModuleUserAffiliateList) GetUserId() *string {
	return s.UserId
}

func (s *HotelOrderListQueryResponseBodyModuleUserAffiliateList) GetUserName() *string {
	return s.UserName
}

func (s *HotelOrderListQueryResponseBodyModuleUserAffiliateList) SetUserId(v string) *HotelOrderListQueryResponseBodyModuleUserAffiliateList {
	s.UserId = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModuleUserAffiliateList) SetUserName(v string) *HotelOrderListQueryResponseBodyModuleUserAffiliateList {
	s.UserName = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyModuleUserAffiliateList) Validate() error {
	return dara.Validate(s)
}

type HotelOrderListQueryResponseBodyPageInfo struct {
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// example:
	//
	// 50
	TotalNumber *int32 `json:"total_number,omitempty" xml:"total_number,omitempty"`
}

func (s HotelOrderListQueryResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s HotelOrderListQueryResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *HotelOrderListQueryResponseBodyPageInfo) GetPage() *int32 {
	return s.Page
}

func (s *HotelOrderListQueryResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *HotelOrderListQueryResponseBodyPageInfo) GetTotalNumber() *int32 {
	return s.TotalNumber
}

func (s *HotelOrderListQueryResponseBodyPageInfo) SetPage(v int32) *HotelOrderListQueryResponseBodyPageInfo {
	s.Page = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyPageInfo) SetPageSize(v int32) *HotelOrderListQueryResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyPageInfo) SetTotalNumber(v int32) *HotelOrderListQueryResponseBodyPageInfo {
	s.TotalNumber = &v
	return s
}

func (s *HotelOrderListQueryResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
