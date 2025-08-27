// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightModifyOtaSearchV2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *FlightModifyOtaSearchV2ResponseBody
	GetCode() *string
	SetMessage(v string) *FlightModifyOtaSearchV2ResponseBody
	GetMessage() *string
	SetModule(v *FlightModifyOtaSearchV2ResponseBodyModule) *FlightModifyOtaSearchV2ResponseBody
	GetModule() *FlightModifyOtaSearchV2ResponseBodyModule
	SetRequestId(v string) *FlightModifyOtaSearchV2ResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *FlightModifyOtaSearchV2ResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *FlightModifyOtaSearchV2ResponseBody
	GetTraceId() *string
}

type FlightModifyOtaSearchV2ResponseBody struct {
	// example:
	//
	// 200
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// module
	Module *FlightModifyOtaSearchV2ResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// requestId
	//
	// example:
	//
	// C61ECFF6-606B-5F66-B81D-D77369043A5F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// traceId
	//
	// example:
	//
	// 2136019116915615639457351e06ee
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s FlightModifyOtaSearchV2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyOtaSearchV2ResponseBody) GoString() string {
	return s.String()
}

func (s *FlightModifyOtaSearchV2ResponseBody) GetCode() *string {
	return s.Code
}

func (s *FlightModifyOtaSearchV2ResponseBody) GetMessage() *string {
	return s.Message
}

func (s *FlightModifyOtaSearchV2ResponseBody) GetModule() *FlightModifyOtaSearchV2ResponseBodyModule {
	return s.Module
}

func (s *FlightModifyOtaSearchV2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FlightModifyOtaSearchV2ResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *FlightModifyOtaSearchV2ResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *FlightModifyOtaSearchV2ResponseBody) SetCode(v string) *FlightModifyOtaSearchV2ResponseBody {
	s.Code = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBody) SetMessage(v string) *FlightModifyOtaSearchV2ResponseBody {
	s.Message = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBody) SetModule(v *FlightModifyOtaSearchV2ResponseBodyModule) *FlightModifyOtaSearchV2ResponseBody {
	s.Module = v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBody) SetRequestId(v string) *FlightModifyOtaSearchV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBody) SetSuccess(v bool) *FlightModifyOtaSearchV2ResponseBody {
	s.Success = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBody) SetTraceId(v string) *FlightModifyOtaSearchV2ResponseBody {
	s.TraceId = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBody) Validate() error {
	return dara.Validate(s)
}

type FlightModifyOtaSearchV2ResponseBodyModule struct {
	AgentInfos []*FlightModifyOtaSearchV2ResponseBodyModuleAgentInfos `json:"agentInfos,omitempty" xml:"agentInfos,omitempty" type:"Repeated"`
	AgentInfo  *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfo    `json:"agent_info,omitempty" xml:"agent_info,omitempty" type:"Struct"`
	// example:
	//
	// 2136019116915615639457351e06ee
	CacheKey           *string                                                          `json:"cache_key,omitempty" xml:"cache_key,omitempty"`
	FlightSegmentInfos [][]*FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos `json:"flight_segment_infos,omitempty" xml:"flight_segment_infos,omitempty" type:"Repeated"`
	PassengerCount     *FlightModifyOtaSearchV2ResponseBodyModulePassengerCount         `json:"passenger_count,omitempty" xml:"passenger_count,omitempty" type:"Struct"`
	// example:
	//
	// a2ffebfe733742aab5c491d960ba3d59
	SessionId *string `json:"session_id,omitempty" xml:"session_id,omitempty"`
}

func (s FlightModifyOtaSearchV2ResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyOtaSearchV2ResponseBodyModule) GoString() string {
	return s.String()
}

func (s *FlightModifyOtaSearchV2ResponseBodyModule) GetAgentInfos() []*FlightModifyOtaSearchV2ResponseBodyModuleAgentInfos {
	return s.AgentInfos
}

func (s *FlightModifyOtaSearchV2ResponseBodyModule) GetAgentInfo() *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfo {
	return s.AgentInfo
}

func (s *FlightModifyOtaSearchV2ResponseBodyModule) GetCacheKey() *string {
	return s.CacheKey
}

func (s *FlightModifyOtaSearchV2ResponseBodyModule) GetFlightSegmentInfos() [][]*FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos {
	return s.FlightSegmentInfos
}

func (s *FlightModifyOtaSearchV2ResponseBodyModule) GetPassengerCount() *FlightModifyOtaSearchV2ResponseBodyModulePassengerCount {
	return s.PassengerCount
}

func (s *FlightModifyOtaSearchV2ResponseBodyModule) GetSessionId() *string {
	return s.SessionId
}

func (s *FlightModifyOtaSearchV2ResponseBodyModule) SetAgentInfos(v []*FlightModifyOtaSearchV2ResponseBodyModuleAgentInfos) *FlightModifyOtaSearchV2ResponseBodyModule {
	s.AgentInfos = v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModule) SetAgentInfo(v *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfo) *FlightModifyOtaSearchV2ResponseBodyModule {
	s.AgentInfo = v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModule) SetCacheKey(v string) *FlightModifyOtaSearchV2ResponseBodyModule {
	s.CacheKey = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModule) SetFlightSegmentInfos(v [][]*FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos) *FlightModifyOtaSearchV2ResponseBodyModule {
	s.FlightSegmentInfos = v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModule) SetPassengerCount(v *FlightModifyOtaSearchV2ResponseBodyModulePassengerCount) *FlightModifyOtaSearchV2ResponseBodyModule {
	s.PassengerCount = v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModule) SetSessionId(v string) *FlightModifyOtaSearchV2ResponseBodyModule {
	s.SessionId = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type FlightModifyOtaSearchV2ResponseBodyModuleAgentInfos struct {
	AttributeShowInfoMap map[string][]*ModuleAgentInfosAttributeShowInfoMapValue            `json:"attribute_show_info_map,omitempty" xml:"attribute_show_info_map,omitempty"`
	BestDiscount         *float64                                                           `json:"best_discount,omitempty" xml:"best_discount,omitempty"`
	CabinClassInfo       *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosCabinClassInfo `json:"cabin_class_info,omitempty" xml:"cabin_class_info,omitempty" type:"Struct"`
	CabinCode            *int32                                                             `json:"cabin_code,omitempty" xml:"cabin_code,omitempty"`
	CabinName            *string                                                            `json:"cabin_name,omitempty" xml:"cabin_name,omitempty"`
	ItemId               *string                                                            `json:"item_id,omitempty" xml:"item_id,omitempty"`
	ModifyTypeDesc       *string                                                            `json:"modify_type_desc,omitempty" xml:"modify_type_desc,omitempty"`
	ModifyTypeName       *string                                                            `json:"modify_type_name,omitempty" xml:"modify_type_name,omitempty"`
	PriceInfoDTO         *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTO   `json:"price_info_d_t_o,omitempty" xml:"price_info_d_t_o,omitempty" type:"Struct"`
	Quantity             *int32                                                             `json:"quantity,omitempty" xml:"quantity,omitempty"`
	SupportChildTicket   *bool                                                              `json:"support_child_ticket,omitempty" xml:"support_child_ticket,omitempty"`
}

func (s FlightModifyOtaSearchV2ResponseBodyModuleAgentInfos) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyOtaSearchV2ResponseBodyModuleAgentInfos) GoString() string {
	return s.String()
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfos) GetAttributeShowInfoMap() map[string][]*ModuleAgentInfosAttributeShowInfoMapValue {
	return s.AttributeShowInfoMap
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfos) GetBestDiscount() *float64 {
	return s.BestDiscount
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfos) GetCabinClassInfo() *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosCabinClassInfo {
	return s.CabinClassInfo
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfos) GetCabinCode() *int32 {
	return s.CabinCode
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfos) GetCabinName() *string {
	return s.CabinName
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfos) GetItemId() *string {
	return s.ItemId
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfos) GetModifyTypeDesc() *string {
	return s.ModifyTypeDesc
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfos) GetModifyTypeName() *string {
	return s.ModifyTypeName
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfos) GetPriceInfoDTO() *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTO {
	return s.PriceInfoDTO
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfos) GetQuantity() *int32 {
	return s.Quantity
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfos) GetSupportChildTicket() *bool {
	return s.SupportChildTicket
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfos) SetAttributeShowInfoMap(v map[string][]*ModuleAgentInfosAttributeShowInfoMapValue) *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfos {
	s.AttributeShowInfoMap = v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfos) SetBestDiscount(v float64) *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfos {
	s.BestDiscount = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfos) SetCabinClassInfo(v *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosCabinClassInfo) *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfos {
	s.CabinClassInfo = v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfos) SetCabinCode(v int32) *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfos {
	s.CabinCode = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfos) SetCabinName(v string) *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfos {
	s.CabinName = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfos) SetItemId(v string) *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfos {
	s.ItemId = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfos) SetModifyTypeDesc(v string) *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfos {
	s.ModifyTypeDesc = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfos) SetModifyTypeName(v string) *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfos {
	s.ModifyTypeName = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfos) SetPriceInfoDTO(v *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTO) *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfos {
	s.PriceInfoDTO = v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfos) SetQuantity(v int32) *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfos {
	s.Quantity = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfos) SetSupportChildTicket(v bool) *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfos {
	s.SupportChildTicket = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfos) Validate() error {
	return dara.Validate(s)
}

type FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosCabinClassInfo struct {
	CabinClass      *string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	ClassName       *string `json:"class_name,omitempty" xml:"class_name,omitempty"`
	InnerCabinClass *int32  `json:"inner_cabin_class,omitempty" xml:"inner_cabin_class,omitempty"`
	Quantity        *string `json:"quantity,omitempty" xml:"quantity,omitempty"`
}

func (s FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosCabinClassInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosCabinClassInfo) GoString() string {
	return s.String()
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosCabinClassInfo) GetCabinClass() *string {
	return s.CabinClass
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosCabinClassInfo) GetClassName() *string {
	return s.ClassName
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosCabinClassInfo) GetInnerCabinClass() *int32 {
	return s.InnerCabinClass
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosCabinClassInfo) GetQuantity() *string {
	return s.Quantity
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosCabinClassInfo) SetCabinClass(v string) *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosCabinClassInfo {
	s.CabinClass = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosCabinClassInfo) SetClassName(v string) *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosCabinClassInfo {
	s.ClassName = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosCabinClassInfo) SetInnerCabinClass(v int32) *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosCabinClassInfo {
	s.InnerCabinClass = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosCabinClassInfo) SetQuantity(v string) *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosCabinClassInfo {
	s.Quantity = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosCabinClassInfo) Validate() error {
	return dara.Validate(s)
}

type FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTO struct {
	AdultPrice              *int32                                                                             `json:"adult_price,omitempty" xml:"adult_price,omitempty"`
	AdultTax                *int32                                                                             `json:"adult_tax,omitempty" xml:"adult_tax,omitempty"`
	AdultTotalPrice         *int32                                                                             `json:"adult_total_price,omitempty" xml:"adult_total_price,omitempty"`
	BeforeControlPrice      *int32                                                                             `json:"before_control_price,omitempty" xml:"before_control_price,omitempty"`
	ChildPrice              *int32                                                                             `json:"child_price,omitempty" xml:"child_price,omitempty"`
	ChildTax                *int32                                                                             `json:"child_tax,omitempty" xml:"child_tax,omitempty"`
	ChildTotalPrice         *int32                                                                             `json:"child_total_price,omitempty" xml:"child_total_price,omitempty"`
	InfantPrice             *int32                                                                             `json:"infant_price,omitempty" xml:"infant_price,omitempty"`
	InfantTax               *int32                                                                             `json:"infant_tax,omitempty" xml:"infant_tax,omitempty"`
	InfantTotalPrice        *int32                                                                             `json:"infant_total_price,omitempty" xml:"infant_total_price,omitempty"`
	OriginalAdultPrice      *int32                                                                             `json:"original_adult_price,omitempty" xml:"original_adult_price,omitempty"`
	OriginalAdultTotalPrice *int32                                                                             `json:"original_adult_total_price,omitempty" xml:"original_adult_total_price,omitempty"`
	ReShopPriceInfoDTO      *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTOReShopPriceInfoDTO `json:"re_shop_price_info_d_t_o,omitempty" xml:"re_shop_price_info_d_t_o,omitempty" type:"Struct"`
}

func (s FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTO) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTO) GoString() string {
	return s.String()
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTO) GetAdultPrice() *int32 {
	return s.AdultPrice
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTO) GetAdultTax() *int32 {
	return s.AdultTax
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTO) GetAdultTotalPrice() *int32 {
	return s.AdultTotalPrice
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTO) GetBeforeControlPrice() *int32 {
	return s.BeforeControlPrice
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTO) GetChildPrice() *int32 {
	return s.ChildPrice
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTO) GetChildTax() *int32 {
	return s.ChildTax
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTO) GetChildTotalPrice() *int32 {
	return s.ChildTotalPrice
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTO) GetInfantPrice() *int32 {
	return s.InfantPrice
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTO) GetInfantTax() *int32 {
	return s.InfantTax
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTO) GetInfantTotalPrice() *int32 {
	return s.InfantTotalPrice
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTO) GetOriginalAdultPrice() *int32 {
	return s.OriginalAdultPrice
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTO) GetOriginalAdultTotalPrice() *int32 {
	return s.OriginalAdultTotalPrice
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTO) GetReShopPriceInfoDTO() *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTOReShopPriceInfoDTO {
	return s.ReShopPriceInfoDTO
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTO) SetAdultPrice(v int32) *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTO {
	s.AdultPrice = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTO) SetAdultTax(v int32) *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTO {
	s.AdultTax = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTO) SetAdultTotalPrice(v int32) *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTO {
	s.AdultTotalPrice = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTO) SetBeforeControlPrice(v int32) *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTO {
	s.BeforeControlPrice = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTO) SetChildPrice(v int32) *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTO {
	s.ChildPrice = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTO) SetChildTax(v int32) *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTO {
	s.ChildTax = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTO) SetChildTotalPrice(v int32) *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTO {
	s.ChildTotalPrice = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTO) SetInfantPrice(v int32) *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTO {
	s.InfantPrice = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTO) SetInfantTax(v int32) *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTO {
	s.InfantTax = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTO) SetInfantTotalPrice(v int32) *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTO {
	s.InfantTotalPrice = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTO) SetOriginalAdultPrice(v int32) *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTO {
	s.OriginalAdultPrice = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTO) SetOriginalAdultTotalPrice(v int32) *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTO {
	s.OriginalAdultTotalPrice = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTO) SetReShopPriceInfoDTO(v *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTOReShopPriceInfoDTO) *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTO {
	s.ReShopPriceInfoDTO = v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTO) Validate() error {
	return dara.Validate(s)
}

type FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTOReShopPriceInfoDTO struct {
	ReShopAdultChangeFee *int32 `json:"re_shop_adult_change_fee,omitempty" xml:"re_shop_adult_change_fee,omitempty"`
	ReShopAdultPrice     *int32 `json:"re_shop_adult_price,omitempty" xml:"re_shop_adult_price,omitempty"`
	ReShopAdultPriceGap  *int32 `json:"re_shop_adult_price_gap,omitempty" xml:"re_shop_adult_price_gap,omitempty"`
	ReShopChildChangeFee *int32 `json:"re_shop_child_change_fee,omitempty" xml:"re_shop_child_change_fee,omitempty"`
	ReShopChildPrice     *int32 `json:"re_shop_child_price,omitempty" xml:"re_shop_child_price,omitempty"`
	ReShopChildPriceGap  *int32 `json:"re_shop_child_price_gap,omitempty" xml:"re_shop_child_price_gap,omitempty"`
	ReShopInfChangeFee   *int32 `json:"re_shop_inf_change_fee,omitempty" xml:"re_shop_inf_change_fee,omitempty"`
	ReShopInfPrice       *int32 `json:"re_shop_inf_price,omitempty" xml:"re_shop_inf_price,omitempty"`
	ReShopInfPriceGap    *int32 `json:"re_shop_inf_price_gap,omitempty" xml:"re_shop_inf_price_gap,omitempty"`
}

func (s FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTOReShopPriceInfoDTO) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTOReShopPriceInfoDTO) GoString() string {
	return s.String()
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTOReShopPriceInfoDTO) GetReShopAdultChangeFee() *int32 {
	return s.ReShopAdultChangeFee
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTOReShopPriceInfoDTO) GetReShopAdultPrice() *int32 {
	return s.ReShopAdultPrice
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTOReShopPriceInfoDTO) GetReShopAdultPriceGap() *int32 {
	return s.ReShopAdultPriceGap
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTOReShopPriceInfoDTO) GetReShopChildChangeFee() *int32 {
	return s.ReShopChildChangeFee
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTOReShopPriceInfoDTO) GetReShopChildPrice() *int32 {
	return s.ReShopChildPrice
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTOReShopPriceInfoDTO) GetReShopChildPriceGap() *int32 {
	return s.ReShopChildPriceGap
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTOReShopPriceInfoDTO) GetReShopInfChangeFee() *int32 {
	return s.ReShopInfChangeFee
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTOReShopPriceInfoDTO) GetReShopInfPrice() *int32 {
	return s.ReShopInfPrice
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTOReShopPriceInfoDTO) GetReShopInfPriceGap() *int32 {
	return s.ReShopInfPriceGap
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTOReShopPriceInfoDTO) SetReShopAdultChangeFee(v int32) *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTOReShopPriceInfoDTO {
	s.ReShopAdultChangeFee = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTOReShopPriceInfoDTO) SetReShopAdultPrice(v int32) *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTOReShopPriceInfoDTO {
	s.ReShopAdultPrice = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTOReShopPriceInfoDTO) SetReShopAdultPriceGap(v int32) *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTOReShopPriceInfoDTO {
	s.ReShopAdultPriceGap = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTOReShopPriceInfoDTO) SetReShopChildChangeFee(v int32) *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTOReShopPriceInfoDTO {
	s.ReShopChildChangeFee = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTOReShopPriceInfoDTO) SetReShopChildPrice(v int32) *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTOReShopPriceInfoDTO {
	s.ReShopChildPrice = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTOReShopPriceInfoDTO) SetReShopChildPriceGap(v int32) *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTOReShopPriceInfoDTO {
	s.ReShopChildPriceGap = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTOReShopPriceInfoDTO) SetReShopInfChangeFee(v int32) *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTOReShopPriceInfoDTO {
	s.ReShopInfChangeFee = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTOReShopPriceInfoDTO) SetReShopInfPrice(v int32) *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTOReShopPriceInfoDTO {
	s.ReShopInfPrice = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTOReShopPriceInfoDTO) SetReShopInfPriceGap(v int32) *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTOReShopPriceInfoDTO {
	s.ReShopInfPriceGap = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfosPriceInfoDTOReShopPriceInfoDTO) Validate() error {
	return dara.Validate(s)
}

type FlightModifyOtaSearchV2ResponseBodyModuleAgentInfo struct {
	AttributeShowInfoMap map[string][]*ModuleAgentInfoAttributeShowInfoMapValue `json:"attribute_show_info_map,omitempty" xml:"attribute_show_info_map,omitempty"`
	// example:
	//
	// 10
	BestDiscount   *float64                                                          `json:"best_discount,omitempty" xml:"best_discount,omitempty"`
	CabinClassInfo *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoCabinClassInfo `json:"cabin_class_info,omitempty" xml:"cabin_class_info,omitempty" type:"Struct"`
	// example:
	//
	// 1
	CabinCode *int32  `json:"cabin_code,omitempty" xml:"cabin_code,omitempty"`
	CabinName *string `json:"cabin_name,omitempty" xml:"cabin_name,omitempty"`
	// item_id
	//
	// example:
	//
	// c85124c527fc4b26b86d0c043ddc08d3_0
	ItemId         *string                                                         `json:"item_id,omitempty" xml:"item_id,omitempty"`
	ModifyTypeDesc *string                                                         `json:"modify_type_desc,omitempty" xml:"modify_type_desc,omitempty"`
	ModifyTypeName *string                                                         `json:"modify_type_name,omitempty" xml:"modify_type_name,omitempty"`
	PriceInfoDTO   *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTO `json:"price_info_d_t_o,omitempty" xml:"price_info_d_t_o,omitempty" type:"Struct"`
	// example:
	//
	// 8
	Quantity *int32 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// example:
	//
	// true
	SupportChildTicket *bool `json:"support_child_ticket,omitempty" xml:"support_child_ticket,omitempty"`
}

func (s FlightModifyOtaSearchV2ResponseBodyModuleAgentInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyOtaSearchV2ResponseBodyModuleAgentInfo) GoString() string {
	return s.String()
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfo) GetAttributeShowInfoMap() map[string][]*ModuleAgentInfoAttributeShowInfoMapValue {
	return s.AttributeShowInfoMap
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfo) GetBestDiscount() *float64 {
	return s.BestDiscount
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfo) GetCabinClassInfo() *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoCabinClassInfo {
	return s.CabinClassInfo
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfo) GetCabinCode() *int32 {
	return s.CabinCode
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfo) GetCabinName() *string {
	return s.CabinName
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfo) GetItemId() *string {
	return s.ItemId
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfo) GetModifyTypeDesc() *string {
	return s.ModifyTypeDesc
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfo) GetModifyTypeName() *string {
	return s.ModifyTypeName
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfo) GetPriceInfoDTO() *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTO {
	return s.PriceInfoDTO
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfo) GetQuantity() *int32 {
	return s.Quantity
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfo) GetSupportChildTicket() *bool {
	return s.SupportChildTicket
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfo) SetAttributeShowInfoMap(v map[string][]*ModuleAgentInfoAttributeShowInfoMapValue) *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfo {
	s.AttributeShowInfoMap = v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfo) SetBestDiscount(v float64) *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfo {
	s.BestDiscount = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfo) SetCabinClassInfo(v *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoCabinClassInfo) *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfo {
	s.CabinClassInfo = v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfo) SetCabinCode(v int32) *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfo {
	s.CabinCode = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfo) SetCabinName(v string) *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfo {
	s.CabinName = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfo) SetItemId(v string) *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfo {
	s.ItemId = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfo) SetModifyTypeDesc(v string) *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfo {
	s.ModifyTypeDesc = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfo) SetModifyTypeName(v string) *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfo {
	s.ModifyTypeName = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfo) SetPriceInfoDTO(v *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTO) *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfo {
	s.PriceInfoDTO = v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfo) SetQuantity(v int32) *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfo {
	s.Quantity = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfo) SetSupportChildTicket(v bool) *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfo {
	s.SupportChildTicket = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfo) Validate() error {
	return dara.Validate(s)
}

type FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoCabinClassInfo struct {
	CabinClass *string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	ClassName  *string `json:"class_name,omitempty" xml:"class_name,omitempty"`
	// inner_cabin_class
	//
	// example:
	//
	// 1
	InnerCabinClass *int32  `json:"inner_cabin_class,omitempty" xml:"inner_cabin_class,omitempty"`
	Quantity        *string `json:"quantity,omitempty" xml:"quantity,omitempty"`
}

func (s FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoCabinClassInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoCabinClassInfo) GoString() string {
	return s.String()
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoCabinClassInfo) GetCabinClass() *string {
	return s.CabinClass
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoCabinClassInfo) GetClassName() *string {
	return s.ClassName
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoCabinClassInfo) GetInnerCabinClass() *int32 {
	return s.InnerCabinClass
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoCabinClassInfo) GetQuantity() *string {
	return s.Quantity
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoCabinClassInfo) SetCabinClass(v string) *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoCabinClassInfo {
	s.CabinClass = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoCabinClassInfo) SetClassName(v string) *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoCabinClassInfo {
	s.ClassName = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoCabinClassInfo) SetInnerCabinClass(v int32) *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoCabinClassInfo {
	s.InnerCabinClass = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoCabinClassInfo) SetQuantity(v string) *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoCabinClassInfo {
	s.Quantity = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoCabinClassInfo) Validate() error {
	return dara.Validate(s)
}

type FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTO struct {
	// example:
	//
	// 126000
	AdultPrice *int32 `json:"adult_price,omitempty" xml:"adult_price,omitempty"`
	// example:
	//
	// 11000
	AdultTax *int32 `json:"adult_tax,omitempty" xml:"adult_tax,omitempty"`
	// example:
	//
	// 137000
	AdultTotalPrice *int32 `json:"adult_total_price,omitempty" xml:"adult_total_price,omitempty"`
	// example:
	//
	// 126000
	BeforeControlPrice *int32 `json:"before_control_price,omitempty" xml:"before_control_price,omitempty"`
	// example:
	//
	// 64000
	ChildPrice *int32 `json:"child_price,omitempty" xml:"child_price,omitempty"`
	// example:
	//
	// 2000
	ChildTax *int32 `json:"child_tax,omitempty" xml:"child_tax,omitempty"`
	// example:
	//
	// 66000
	ChildTotalPrice *int32 `json:"child_total_price,omitempty" xml:"child_total_price,omitempty"`
	// example:
	//
	// 120
	InfantPrice *int32 `json:"infant_price,omitempty" xml:"infant_price,omitempty"`
	// example:
	//
	// 0
	InfantTax *int32 `json:"infant_tax,omitempty" xml:"infant_tax,omitempty"`
	// example:
	//
	// 120
	InfantTotalPrice *int32 `json:"infant_total_price,omitempty" xml:"infant_total_price,omitempty"`
	// example:
	//
	// 1300
	OriginalAdultPrice *int32 `json:"original_adult_price,omitempty" xml:"original_adult_price,omitempty"`
	// example:
	//
	// 12300
	OriginalAdultTotalPrice *int32                                                                            `json:"original_adult_total_price,omitempty" xml:"original_adult_total_price,omitempty"`
	ReShopPriceInfoDTO      *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTOReShopPriceInfoDTO `json:"re_shop_price_info_d_t_o,omitempty" xml:"re_shop_price_info_d_t_o,omitempty" type:"Struct"`
}

func (s FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTO) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTO) GoString() string {
	return s.String()
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTO) GetAdultPrice() *int32 {
	return s.AdultPrice
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTO) GetAdultTax() *int32 {
	return s.AdultTax
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTO) GetAdultTotalPrice() *int32 {
	return s.AdultTotalPrice
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTO) GetBeforeControlPrice() *int32 {
	return s.BeforeControlPrice
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTO) GetChildPrice() *int32 {
	return s.ChildPrice
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTO) GetChildTax() *int32 {
	return s.ChildTax
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTO) GetChildTotalPrice() *int32 {
	return s.ChildTotalPrice
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTO) GetInfantPrice() *int32 {
	return s.InfantPrice
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTO) GetInfantTax() *int32 {
	return s.InfantTax
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTO) GetInfantTotalPrice() *int32 {
	return s.InfantTotalPrice
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTO) GetOriginalAdultPrice() *int32 {
	return s.OriginalAdultPrice
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTO) GetOriginalAdultTotalPrice() *int32 {
	return s.OriginalAdultTotalPrice
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTO) GetReShopPriceInfoDTO() *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTOReShopPriceInfoDTO {
	return s.ReShopPriceInfoDTO
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTO) SetAdultPrice(v int32) *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTO {
	s.AdultPrice = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTO) SetAdultTax(v int32) *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTO {
	s.AdultTax = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTO) SetAdultTotalPrice(v int32) *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTO {
	s.AdultTotalPrice = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTO) SetBeforeControlPrice(v int32) *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTO {
	s.BeforeControlPrice = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTO) SetChildPrice(v int32) *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTO {
	s.ChildPrice = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTO) SetChildTax(v int32) *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTO {
	s.ChildTax = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTO) SetChildTotalPrice(v int32) *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTO {
	s.ChildTotalPrice = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTO) SetInfantPrice(v int32) *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTO {
	s.InfantPrice = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTO) SetInfantTax(v int32) *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTO {
	s.InfantTax = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTO) SetInfantTotalPrice(v int32) *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTO {
	s.InfantTotalPrice = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTO) SetOriginalAdultPrice(v int32) *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTO {
	s.OriginalAdultPrice = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTO) SetOriginalAdultTotalPrice(v int32) *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTO {
	s.OriginalAdultTotalPrice = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTO) SetReShopPriceInfoDTO(v *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTOReShopPriceInfoDTO) *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTO {
	s.ReShopPriceInfoDTO = v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTO) Validate() error {
	return dara.Validate(s)
}

type FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTOReShopPriceInfoDTO struct {
	// example:
	//
	// -1
	ReShopAdultChangeFee *int32 `json:"re_shop_adult_change_fee,omitempty" xml:"re_shop_adult_change_fee,omitempty"`
	// example:
	//
	// -1
	ReShopAdultPrice *int32 `json:"re_shop_adult_price,omitempty" xml:"re_shop_adult_price,omitempty"`
	// example:
	//
	// -1
	ReShopAdultPriceGap *int32 `json:"re_shop_adult_price_gap,omitempty" xml:"re_shop_adult_price_gap,omitempty"`
	// example:
	//
	// -1
	ReShopChildChangeFee *int32 `json:"re_shop_child_change_fee,omitempty" xml:"re_shop_child_change_fee,omitempty"`
	// example:
	//
	// -1
	ReShopChildPrice *int32 `json:"re_shop_child_price,omitempty" xml:"re_shop_child_price,omitempty"`
	// example:
	//
	// -1
	ReShopChildPriceGap *int32 `json:"re_shop_child_price_gap,omitempty" xml:"re_shop_child_price_gap,omitempty"`
	// example:
	//
	// -1
	ReShopInfChangeFee *int32 `json:"re_shop_inf_change_fee,omitempty" xml:"re_shop_inf_change_fee,omitempty"`
	// example:
	//
	// -1
	ReShopInfPrice *int32 `json:"re_shop_inf_price,omitempty" xml:"re_shop_inf_price,omitempty"`
	// example:
	//
	// -1
	ReShopInfPriceGap *int32 `json:"re_shop_inf_price_gap,omitempty" xml:"re_shop_inf_price_gap,omitempty"`
}

func (s FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTOReShopPriceInfoDTO) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTOReShopPriceInfoDTO) GoString() string {
	return s.String()
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTOReShopPriceInfoDTO) GetReShopAdultChangeFee() *int32 {
	return s.ReShopAdultChangeFee
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTOReShopPriceInfoDTO) GetReShopAdultPrice() *int32 {
	return s.ReShopAdultPrice
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTOReShopPriceInfoDTO) GetReShopAdultPriceGap() *int32 {
	return s.ReShopAdultPriceGap
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTOReShopPriceInfoDTO) GetReShopChildChangeFee() *int32 {
	return s.ReShopChildChangeFee
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTOReShopPriceInfoDTO) GetReShopChildPrice() *int32 {
	return s.ReShopChildPrice
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTOReShopPriceInfoDTO) GetReShopChildPriceGap() *int32 {
	return s.ReShopChildPriceGap
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTOReShopPriceInfoDTO) GetReShopInfChangeFee() *int32 {
	return s.ReShopInfChangeFee
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTOReShopPriceInfoDTO) GetReShopInfPrice() *int32 {
	return s.ReShopInfPrice
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTOReShopPriceInfoDTO) GetReShopInfPriceGap() *int32 {
	return s.ReShopInfPriceGap
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTOReShopPriceInfoDTO) SetReShopAdultChangeFee(v int32) *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTOReShopPriceInfoDTO {
	s.ReShopAdultChangeFee = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTOReShopPriceInfoDTO) SetReShopAdultPrice(v int32) *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTOReShopPriceInfoDTO {
	s.ReShopAdultPrice = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTOReShopPriceInfoDTO) SetReShopAdultPriceGap(v int32) *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTOReShopPriceInfoDTO {
	s.ReShopAdultPriceGap = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTOReShopPriceInfoDTO) SetReShopChildChangeFee(v int32) *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTOReShopPriceInfoDTO {
	s.ReShopChildChangeFee = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTOReShopPriceInfoDTO) SetReShopChildPrice(v int32) *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTOReShopPriceInfoDTO {
	s.ReShopChildPrice = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTOReShopPriceInfoDTO) SetReShopChildPriceGap(v int32) *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTOReShopPriceInfoDTO {
	s.ReShopChildPriceGap = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTOReShopPriceInfoDTO) SetReShopInfChangeFee(v int32) *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTOReShopPriceInfoDTO {
	s.ReShopInfChangeFee = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTOReShopPriceInfoDTO) SetReShopInfPrice(v int32) *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTOReShopPriceInfoDTO {
	s.ReShopInfPrice = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTOReShopPriceInfoDTO) SetReShopInfPriceGap(v int32) *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTOReShopPriceInfoDTO {
	s.ReShopInfPriceGap = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleAgentInfoPriceInfoDTOReShopPriceInfoDTO) Validate() error {
	return dara.Validate(s)
}

type FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos struct {
	// example:
	//
	// 0
	JourneySeq *int32 `json:"journey_seq,omitempty" xml:"journey_seq,omitempty"`
	// example:
	//
	// 0
	SegmentSeq *int32 `json:"segment_seq,omitempty" xml:"segment_seq,omitempty"`
	// example:
	//
	// CA1110
	FlightNo *string `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
	// example:
	//
	// XIL
	DepCityCode *string `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	// example:
	//
	// BJS
	ArrCityCode    *string                                                                    `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	DepCityName    *string                                                                    `json:"dep_city_name,omitempty" xml:"dep_city_name,omitempty"`
	ArrCityName    *string                                                                    `json:"arr_city_name,omitempty" xml:"arr_city_name,omitempty"`
	DepAirportInfo *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosDepAirportInfo `json:"dep_airport_info,omitempty" xml:"dep_airport_info,omitempty" type:"Struct"`
	ArrAirportInfo *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosArrAirportInfo `json:"arr_airport_info,omitempty" xml:"arr_airport_info,omitempty" type:"Struct"`
	// example:
	//
	// 2023-09-18 09:10:00
	DepTime *string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	// example:
	//
	// 2023-09-18 10:25:00
	ArrTime     *string                                                                 `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	AirlineInfo *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosAirlineInfo `json:"airline_info,omitempty" xml:"airline_info,omitempty" type:"Struct"`
	// example:
	//
	// false
	Share            *bool                                                                        `json:"share,omitempty" xml:"share,omitempty"`
	FlightSharedInfo *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosFlightSharedInfo `json:"flight_shared_info,omitempty" xml:"flight_shared_info,omitempty" type:"Struct"`
	// example:
	//
	// false
	Stop           *bool                                                                      `json:"stop,omitempty" xml:"stop,omitempty"`
	FlightStopInfo *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosFlightStopInfo `json:"flight_stop_info,omitempty" xml:"flight_stop_info,omitempty" type:"Struct"`
	// example:
	//
	// 20
	TransferTime *int32 `json:"transfer_time,omitempty" xml:"transfer_time,omitempty"`
	// example:
	//
	// 75
	Duration     *int32  `json:"duration,omitempty" xml:"duration,omitempty"`
	Manufacturer *string `json:"manufacturer,omitempty" xml:"manufacturer,omitempty"`
	// example:
	//
	// ARJ
	FlightType *string `json:"flight_type,omitempty" xml:"flight_type,omitempty"`
	FlightSize *string `json:"flight_size,omitempty" xml:"flight_size,omitempty"`
	MealDesc   *string `json:"meal_desc,omitempty" xml:"meal_desc,omitempty"`
	OnTimeRate *string `json:"on_time_rate,omitempty" xml:"on_time_rate,omitempty"`
}

func (s FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos) GoString() string {
	return s.String()
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos) GetJourneySeq() *int32 {
	return s.JourneySeq
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos) GetSegmentSeq() *int32 {
	return s.SegmentSeq
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos) GetFlightNo() *string {
	return s.FlightNo
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos) GetDepCityName() *string {
	return s.DepCityName
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos) GetArrCityName() *string {
	return s.ArrCityName
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos) GetDepAirportInfo() *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosDepAirportInfo {
	return s.DepAirportInfo
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos) GetArrAirportInfo() *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosArrAirportInfo {
	return s.ArrAirportInfo
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos) GetDepTime() *string {
	return s.DepTime
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos) GetArrTime() *string {
	return s.ArrTime
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos) GetAirlineInfo() *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosAirlineInfo {
	return s.AirlineInfo
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos) GetShare() *bool {
	return s.Share
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos) GetFlightSharedInfo() *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosFlightSharedInfo {
	return s.FlightSharedInfo
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos) GetStop() *bool {
	return s.Stop
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos) GetFlightStopInfo() *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosFlightStopInfo {
	return s.FlightStopInfo
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos) GetTransferTime() *int32 {
	return s.TransferTime
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos) GetDuration() *int32 {
	return s.Duration
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos) GetManufacturer() *string {
	return s.Manufacturer
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos) GetFlightType() *string {
	return s.FlightType
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos) GetFlightSize() *string {
	return s.FlightSize
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos) GetMealDesc() *string {
	return s.MealDesc
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos) GetOnTimeRate() *string {
	return s.OnTimeRate
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos) SetJourneySeq(v int32) *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos {
	s.JourneySeq = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos) SetSegmentSeq(v int32) *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos {
	s.SegmentSeq = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos) SetFlightNo(v string) *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos {
	s.FlightNo = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos) SetDepCityCode(v string) *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos {
	s.DepCityCode = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos) SetArrCityCode(v string) *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos {
	s.ArrCityCode = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos) SetDepCityName(v string) *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos {
	s.DepCityName = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos) SetArrCityName(v string) *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos {
	s.ArrCityName = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos) SetDepAirportInfo(v *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosDepAirportInfo) *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos {
	s.DepAirportInfo = v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos) SetArrAirportInfo(v *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosArrAirportInfo) *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos {
	s.ArrAirportInfo = v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos) SetDepTime(v string) *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos {
	s.DepTime = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos) SetArrTime(v string) *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos {
	s.ArrTime = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos) SetAirlineInfo(v *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosAirlineInfo) *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos {
	s.AirlineInfo = v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos) SetShare(v bool) *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos {
	s.Share = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos) SetFlightSharedInfo(v *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosFlightSharedInfo) *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos {
	s.FlightSharedInfo = v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos) SetStop(v bool) *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos {
	s.Stop = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos) SetFlightStopInfo(v *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosFlightStopInfo) *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos {
	s.FlightStopInfo = v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos) SetTransferTime(v int32) *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos {
	s.TransferTime = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos) SetDuration(v int32) *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos {
	s.Duration = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos) SetManufacturer(v string) *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos {
	s.Manufacturer = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos) SetFlightType(v string) *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos {
	s.FlightType = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos) SetFlightSize(v string) *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos {
	s.FlightSize = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos) SetMealDesc(v string) *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos {
	s.MealDesc = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos) SetOnTimeRate(v string) *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos {
	s.OnTimeRate = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfos) Validate() error {
	return dara.Validate(s)
}

type FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosDepAirportInfo struct {
	// example:
	//
	// XIL
	AirportCode      *string `json:"airport_code,omitempty" xml:"airport_code,omitempty"`
	AirportName      *string `json:"airport_name,omitempty" xml:"airport_name,omitempty"`
	AirportShortName *string `json:"airport_short_name,omitempty" xml:"airport_short_name,omitempty"`
	// example:
	//
	// T3
	Terminal *string `json:"terminal,omitempty" xml:"terminal,omitempty"`
}

func (s FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosDepAirportInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosDepAirportInfo) GoString() string {
	return s.String()
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosDepAirportInfo) GetAirportCode() *string {
	return s.AirportCode
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosDepAirportInfo) GetAirportName() *string {
	return s.AirportName
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosDepAirportInfo) GetAirportShortName() *string {
	return s.AirportShortName
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosDepAirportInfo) GetTerminal() *string {
	return s.Terminal
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosDepAirportInfo) SetAirportCode(v string) *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosDepAirportInfo {
	s.AirportCode = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosDepAirportInfo) SetAirportName(v string) *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosDepAirportInfo {
	s.AirportName = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosDepAirportInfo) SetAirportShortName(v string) *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosDepAirportInfo {
	s.AirportShortName = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosDepAirportInfo) SetTerminal(v string) *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosDepAirportInfo {
	s.Terminal = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosDepAirportInfo) Validate() error {
	return dara.Validate(s)
}

type FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosArrAirportInfo struct {
	// example:
	//
	// PEK
	AirportCode      *string `json:"airport_code,omitempty" xml:"airport_code,omitempty"`
	AirportName      *string `json:"airport_name,omitempty" xml:"airport_name,omitempty"`
	AirportShortName *string `json:"airport_short_name,omitempty" xml:"airport_short_name,omitempty"`
	// example:
	//
	// T2
	Terminal *string `json:"terminal,omitempty" xml:"terminal,omitempty"`
}

func (s FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosArrAirportInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosArrAirportInfo) GoString() string {
	return s.String()
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosArrAirportInfo) GetAirportCode() *string {
	return s.AirportCode
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosArrAirportInfo) GetAirportName() *string {
	return s.AirportName
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosArrAirportInfo) GetAirportShortName() *string {
	return s.AirportShortName
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosArrAirportInfo) GetTerminal() *string {
	return s.Terminal
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosArrAirportInfo) SetAirportCode(v string) *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosArrAirportInfo {
	s.AirportCode = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosArrAirportInfo) SetAirportName(v string) *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosArrAirportInfo {
	s.AirportName = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosArrAirportInfo) SetAirportShortName(v string) *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosArrAirportInfo {
	s.AirportShortName = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosArrAirportInfo) SetTerminal(v string) *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosArrAirportInfo {
	s.Terminal = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosArrAirportInfo) Validate() error {
	return dara.Validate(s)
}

type FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosAirlineInfo struct {
	// example:
	//
	// CA
	AirlineCode             *string `json:"airline_code,omitempty" xml:"airline_code,omitempty"`
	AirlineChineseName      *string `json:"airline_chinese_name,omitempty" xml:"airline_chinese_name,omitempty"`
	AirlineChineseShortName *string `json:"airline_chinese_short_name,omitempty" xml:"airline_chinese_short_name,omitempty"`
	// example:
	//
	// //gw.alicdn.com/tfs/TB12fJAFHr1gK0jSZR0XXbP8XXa-450-450.png_80x80.jpg
	AirlineIcon *string `json:"airline_icon,omitempty" xml:"airline_icon,omitempty"`
	// example:
	//
	// false
	CheapFlight *bool `json:"cheap_flight,omitempty" xml:"cheap_flight,omitempty"`
}

func (s FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosAirlineInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosAirlineInfo) GoString() string {
	return s.String()
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosAirlineInfo) GetAirlineCode() *string {
	return s.AirlineCode
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosAirlineInfo) GetAirlineChineseName() *string {
	return s.AirlineChineseName
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosAirlineInfo) GetAirlineChineseShortName() *string {
	return s.AirlineChineseShortName
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosAirlineInfo) GetAirlineIcon() *string {
	return s.AirlineIcon
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosAirlineInfo) GetCheapFlight() *bool {
	return s.CheapFlight
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosAirlineInfo) SetAirlineCode(v string) *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosAirlineInfo {
	s.AirlineCode = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosAirlineInfo) SetAirlineChineseName(v string) *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosAirlineInfo {
	s.AirlineChineseName = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosAirlineInfo) SetAirlineChineseShortName(v string) *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosAirlineInfo {
	s.AirlineChineseShortName = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosAirlineInfo) SetAirlineIcon(v string) *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosAirlineInfo {
	s.AirlineIcon = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosAirlineInfo) SetCheapFlight(v bool) *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosAirlineInfo {
	s.CheapFlight = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosAirlineInfo) Validate() error {
	return dara.Validate(s)
}

type FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosFlightSharedInfo struct {
	OperatingFlightNo    *string                                                                                          `json:"operating_flight_no,omitempty" xml:"operating_flight_no,omitempty"`
	OperatingAirlineInfo *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosFlightSharedInfoOperatingAirlineInfo `json:"operating_airline_info,omitempty" xml:"operating_airline_info,omitempty" type:"Struct"`
}

func (s FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosFlightSharedInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosFlightSharedInfo) GoString() string {
	return s.String()
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosFlightSharedInfo) GetOperatingFlightNo() *string {
	return s.OperatingFlightNo
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosFlightSharedInfo) GetOperatingAirlineInfo() *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosFlightSharedInfoOperatingAirlineInfo {
	return s.OperatingAirlineInfo
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosFlightSharedInfo) SetOperatingFlightNo(v string) *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosFlightSharedInfo {
	s.OperatingFlightNo = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosFlightSharedInfo) SetOperatingAirlineInfo(v *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosFlightSharedInfoOperatingAirlineInfo) *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosFlightSharedInfo {
	s.OperatingAirlineInfo = v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosFlightSharedInfo) Validate() error {
	return dara.Validate(s)
}

type FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosFlightSharedInfoOperatingAirlineInfo struct {
	AirlineCode             *string `json:"airline_code,omitempty" xml:"airline_code,omitempty"`
	AirlineChineseName      *string `json:"airline_chinese_name,omitempty" xml:"airline_chinese_name,omitempty"`
	AirlineChineseShortName *string `json:"airline_chinese_short_name,omitempty" xml:"airline_chinese_short_name,omitempty"`
	AirlineIcon             *string `json:"airline_icon,omitempty" xml:"airline_icon,omitempty"`
	// example:
	//
	// false
	CheapFlight *bool `json:"cheap_flight,omitempty" xml:"cheap_flight,omitempty"`
}

func (s FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosFlightSharedInfoOperatingAirlineInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosFlightSharedInfoOperatingAirlineInfo) GoString() string {
	return s.String()
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosFlightSharedInfoOperatingAirlineInfo) GetAirlineCode() *string {
	return s.AirlineCode
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosFlightSharedInfoOperatingAirlineInfo) GetAirlineChineseName() *string {
	return s.AirlineChineseName
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosFlightSharedInfoOperatingAirlineInfo) GetAirlineChineseShortName() *string {
	return s.AirlineChineseShortName
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosFlightSharedInfoOperatingAirlineInfo) GetAirlineIcon() *string {
	return s.AirlineIcon
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosFlightSharedInfoOperatingAirlineInfo) GetCheapFlight() *bool {
	return s.CheapFlight
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosFlightSharedInfoOperatingAirlineInfo) SetAirlineCode(v string) *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosFlightSharedInfoOperatingAirlineInfo {
	s.AirlineCode = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosFlightSharedInfoOperatingAirlineInfo) SetAirlineChineseName(v string) *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosFlightSharedInfoOperatingAirlineInfo {
	s.AirlineChineseName = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosFlightSharedInfoOperatingAirlineInfo) SetAirlineChineseShortName(v string) *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosFlightSharedInfoOperatingAirlineInfo {
	s.AirlineChineseShortName = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosFlightSharedInfoOperatingAirlineInfo) SetAirlineIcon(v string) *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosFlightSharedInfoOperatingAirlineInfo {
	s.AirlineIcon = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosFlightSharedInfoOperatingAirlineInfo) SetCheapFlight(v bool) *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosFlightSharedInfoOperatingAirlineInfo {
	s.CheapFlight = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosFlightSharedInfoOperatingAirlineInfo) Validate() error {
	return dara.Validate(s)
}

type FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosFlightStopInfo struct {
	StopCityName *string `json:"stop_city_name,omitempty" xml:"stop_city_name,omitempty"`
	StopArrTime  *string `json:"stop_arr_time,omitempty" xml:"stop_arr_time,omitempty"`
	StopDepTime  *string `json:"stop_dep_time,omitempty" xml:"stop_dep_time,omitempty"`
	StopCityCode *string `json:"stop_city_code,omitempty" xml:"stop_city_code,omitempty"`
	StopAirport  *string `json:"stop_airport,omitempty" xml:"stop_airport,omitempty"`
	StopArrTerm  *string `json:"stop_arr_term,omitempty" xml:"stop_arr_term,omitempty"`
	StopDepTerm  *string `json:"stop_dep_term,omitempty" xml:"stop_dep_term,omitempty"`
}

func (s FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosFlightStopInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosFlightStopInfo) GoString() string {
	return s.String()
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosFlightStopInfo) GetStopCityName() *string {
	return s.StopCityName
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosFlightStopInfo) GetStopArrTime() *string {
	return s.StopArrTime
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosFlightStopInfo) GetStopDepTime() *string {
	return s.StopDepTime
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosFlightStopInfo) GetStopCityCode() *string {
	return s.StopCityCode
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosFlightStopInfo) GetStopAirport() *string {
	return s.StopAirport
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosFlightStopInfo) GetStopArrTerm() *string {
	return s.StopArrTerm
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosFlightStopInfo) GetStopDepTerm() *string {
	return s.StopDepTerm
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosFlightStopInfo) SetStopCityName(v string) *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosFlightStopInfo {
	s.StopCityName = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosFlightStopInfo) SetStopArrTime(v string) *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosFlightStopInfo {
	s.StopArrTime = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosFlightStopInfo) SetStopDepTime(v string) *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosFlightStopInfo {
	s.StopDepTime = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosFlightStopInfo) SetStopCityCode(v string) *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosFlightStopInfo {
	s.StopCityCode = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosFlightStopInfo) SetStopAirport(v string) *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosFlightStopInfo {
	s.StopAirport = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosFlightStopInfo) SetStopArrTerm(v string) *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosFlightStopInfo {
	s.StopArrTerm = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosFlightStopInfo) SetStopDepTerm(v string) *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosFlightStopInfo {
	s.StopDepTerm = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModuleFlightSegmentInfosFlightStopInfo) Validate() error {
	return dara.Validate(s)
}

type FlightModifyOtaSearchV2ResponseBodyModulePassengerCount struct {
	// example:
	//
	// 1
	AdultPassengerNum *int32 `json:"adult_passenger_num,omitempty" xml:"adult_passenger_num,omitempty"`
	// example:
	//
	// 0
	ChildPassengerNum *int32 `json:"child_passenger_num,omitempty" xml:"child_passenger_num,omitempty"`
	// example:
	//
	// 0
	InfantPassengerNum *int32 `json:"infant_passenger_num,omitempty" xml:"infant_passenger_num,omitempty"`
}

func (s FlightModifyOtaSearchV2ResponseBodyModulePassengerCount) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyOtaSearchV2ResponseBodyModulePassengerCount) GoString() string {
	return s.String()
}

func (s *FlightModifyOtaSearchV2ResponseBodyModulePassengerCount) GetAdultPassengerNum() *int32 {
	return s.AdultPassengerNum
}

func (s *FlightModifyOtaSearchV2ResponseBodyModulePassengerCount) GetChildPassengerNum() *int32 {
	return s.ChildPassengerNum
}

func (s *FlightModifyOtaSearchV2ResponseBodyModulePassengerCount) GetInfantPassengerNum() *int32 {
	return s.InfantPassengerNum
}

func (s *FlightModifyOtaSearchV2ResponseBodyModulePassengerCount) SetAdultPassengerNum(v int32) *FlightModifyOtaSearchV2ResponseBodyModulePassengerCount {
	s.AdultPassengerNum = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModulePassengerCount) SetChildPassengerNum(v int32) *FlightModifyOtaSearchV2ResponseBodyModulePassengerCount {
	s.ChildPassengerNum = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModulePassengerCount) SetInfantPassengerNum(v int32) *FlightModifyOtaSearchV2ResponseBodyModulePassengerCount {
	s.InfantPassengerNum = &v
	return s
}

func (s *FlightModifyOtaSearchV2ResponseBodyModulePassengerCount) Validate() error {
	return dara.Validate(s)
}
