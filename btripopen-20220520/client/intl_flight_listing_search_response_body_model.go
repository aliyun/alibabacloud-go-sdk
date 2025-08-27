// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightListingSearchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *IntlFlightListingSearchResponseBody
	GetCode() *string
	SetMessage(v string) *IntlFlightListingSearchResponseBody
	GetMessage() *string
	SetModule(v *IntlFlightListingSearchResponseBodyModule) *IntlFlightListingSearchResponseBody
	GetModule() *IntlFlightListingSearchResponseBodyModule
	SetRequestId(v string) *IntlFlightListingSearchResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *IntlFlightListingSearchResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *IntlFlightListingSearchResponseBody
	GetTraceId() *string
}

type IntlFlightListingSearchResponseBody struct {
	// example:
	//
	// SUCCESS
	Code    *string                                    `json:"code,omitempty" xml:"code,omitempty"`
	Message *string                                    `json:"message,omitempty" xml:"message,omitempty"`
	Module  *IntlFlightListingSearchResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// example:
	//
	// 407543AF-2BD9-5890-BD92-9D1AB7218B27
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 210bc4b116835992457938931db4de
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s IntlFlightListingSearchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightListingSearchResponseBody) GoString() string {
	return s.String()
}

func (s *IntlFlightListingSearchResponseBody) GetCode() *string {
	return s.Code
}

func (s *IntlFlightListingSearchResponseBody) GetMessage() *string {
	return s.Message
}

func (s *IntlFlightListingSearchResponseBody) GetModule() *IntlFlightListingSearchResponseBodyModule {
	return s.Module
}

func (s *IntlFlightListingSearchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *IntlFlightListingSearchResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *IntlFlightListingSearchResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *IntlFlightListingSearchResponseBody) SetCode(v string) *IntlFlightListingSearchResponseBody {
	s.Code = &v
	return s
}

func (s *IntlFlightListingSearchResponseBody) SetMessage(v string) *IntlFlightListingSearchResponseBody {
	s.Message = &v
	return s
}

func (s *IntlFlightListingSearchResponseBody) SetModule(v *IntlFlightListingSearchResponseBodyModule) *IntlFlightListingSearchResponseBody {
	s.Module = v
	return s
}

func (s *IntlFlightListingSearchResponseBody) SetRequestId(v string) *IntlFlightListingSearchResponseBody {
	s.RequestId = &v
	return s
}

func (s *IntlFlightListingSearchResponseBody) SetSuccess(v bool) *IntlFlightListingSearchResponseBody {
	s.Success = &v
	return s
}

func (s *IntlFlightListingSearchResponseBody) SetTraceId(v string) *IntlFlightListingSearchResponseBody {
	s.TraceId = &v
	return s
}

func (s *IntlFlightListingSearchResponseBody) Validate() error {
	return dara.Validate(s)
}

type IntlFlightListingSearchResponseBodyModule struct {
	FlightItemList []*IntlFlightListingSearchResponseBodyModuleFlightItemList `json:"flight_item_list,omitempty" xml:"flight_item_list,omitempty" type:"Repeated"`
	// example:
	//
	// false
	NeedContinue *bool `json:"need_continue,omitempty" xml:"need_continue,omitempty"`
	// example:
	//
	// ASDFASDFASDFASDFASDF
	QueryRecordId *string `json:"query_record_id,omitempty" xml:"query_record_id,omitempty"`
	// example:
	//
	// ee229f2d-1835-4199-bfe6-fd14afe8645e
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
}

func (s IntlFlightListingSearchResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightListingSearchResponseBodyModule) GoString() string {
	return s.String()
}

func (s *IntlFlightListingSearchResponseBodyModule) GetFlightItemList() []*IntlFlightListingSearchResponseBodyModuleFlightItemList {
	return s.FlightItemList
}

func (s *IntlFlightListingSearchResponseBodyModule) GetNeedContinue() *bool {
	return s.NeedContinue
}

func (s *IntlFlightListingSearchResponseBodyModule) GetQueryRecordId() *string {
	return s.QueryRecordId
}

func (s *IntlFlightListingSearchResponseBodyModule) GetToken() *string {
	return s.Token
}

func (s *IntlFlightListingSearchResponseBodyModule) SetFlightItemList(v []*IntlFlightListingSearchResponseBodyModuleFlightItemList) *IntlFlightListingSearchResponseBodyModule {
	s.FlightItemList = v
	return s
}

func (s *IntlFlightListingSearchResponseBodyModule) SetNeedContinue(v bool) *IntlFlightListingSearchResponseBodyModule {
	s.NeedContinue = &v
	return s
}

func (s *IntlFlightListingSearchResponseBodyModule) SetQueryRecordId(v string) *IntlFlightListingSearchResponseBodyModule {
	s.QueryRecordId = &v
	return s
}

func (s *IntlFlightListingSearchResponseBodyModule) SetToken(v string) *IntlFlightListingSearchResponseBodyModule {
	s.Token = &v
	return s
}

func (s *IntlFlightListingSearchResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type IntlFlightListingSearchResponseBodyModuleFlightItemList struct {
	BestPriceItem      *IntlFlightListingSearchResponseBodyModuleFlightItemListBestPriceItem        `json:"best_price_item,omitempty" xml:"best_price_item,omitempty" type:"Struct"`
	FlightJourneyInfos []*IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfos `json:"flight_journey_infos,omitempty" xml:"flight_journey_infos,omitempty" type:"Repeated"`
}

func (s IntlFlightListingSearchResponseBodyModuleFlightItemList) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightListingSearchResponseBodyModuleFlightItemList) GoString() string {
	return s.String()
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemList) GetBestPriceItem() *IntlFlightListingSearchResponseBodyModuleFlightItemListBestPriceItem {
	return s.BestPriceItem
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemList) GetFlightJourneyInfos() []*IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfos {
	return s.FlightJourneyInfos
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemList) SetBestPriceItem(v *IntlFlightListingSearchResponseBodyModuleFlightItemListBestPriceItem) *IntlFlightListingSearchResponseBodyModuleFlightItemList {
	s.BestPriceItem = v
	return s
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemList) SetFlightJourneyInfos(v []*IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfos) *IntlFlightListingSearchResponseBodyModuleFlightItemList {
	s.FlightJourneyInfos = v
	return s
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemList) Validate() error {
	return dara.Validate(s)
}

type IntlFlightListingSearchResponseBodyModuleFlightItemListBestPriceItem struct {
	AgreementPriceCodes []*IntlFlightListingSearchResponseBodyModuleFlightItemListBestPriceItemAgreementPriceCodes `json:"agreement_price_codes,omitempty" xml:"agreement_price_codes,omitempty" type:"Repeated"`
	ItemType            *string                                                                                    `json:"item_type,omitempty" xml:"item_type,omitempty"`
	LabelList           []*IntlFlightListingSearchResponseBodyModuleFlightItemListBestPriceItemLabelList           `json:"label_list,omitempty" xml:"label_list,omitempty" type:"Repeated"`
	ShoppingItemMap     map[string]*ModuleFlightItemListBestPriceItemShoppingItemMapValue                          `json:"shopping_item_map,omitempty" xml:"shopping_item_map,omitempty"`
}

func (s IntlFlightListingSearchResponseBodyModuleFlightItemListBestPriceItem) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightListingSearchResponseBodyModuleFlightItemListBestPriceItem) GoString() string {
	return s.String()
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListBestPriceItem) GetAgreementPriceCodes() []*IntlFlightListingSearchResponseBodyModuleFlightItemListBestPriceItemAgreementPriceCodes {
	return s.AgreementPriceCodes
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListBestPriceItem) GetItemType() *string {
	return s.ItemType
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListBestPriceItem) GetLabelList() []*IntlFlightListingSearchResponseBodyModuleFlightItemListBestPriceItemLabelList {
	return s.LabelList
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListBestPriceItem) GetShoppingItemMap() map[string]*ModuleFlightItemListBestPriceItemShoppingItemMapValue {
	return s.ShoppingItemMap
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListBestPriceItem) SetAgreementPriceCodes(v []*IntlFlightListingSearchResponseBodyModuleFlightItemListBestPriceItemAgreementPriceCodes) *IntlFlightListingSearchResponseBodyModuleFlightItemListBestPriceItem {
	s.AgreementPriceCodes = v
	return s
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListBestPriceItem) SetItemType(v string) *IntlFlightListingSearchResponseBodyModuleFlightItemListBestPriceItem {
	s.ItemType = &v
	return s
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListBestPriceItem) SetLabelList(v []*IntlFlightListingSearchResponseBodyModuleFlightItemListBestPriceItemLabelList) *IntlFlightListingSearchResponseBodyModuleFlightItemListBestPriceItem {
	s.LabelList = v
	return s
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListBestPriceItem) SetShoppingItemMap(v map[string]*ModuleFlightItemListBestPriceItemShoppingItemMapValue) *IntlFlightListingSearchResponseBodyModuleFlightItemListBestPriceItem {
	s.ShoppingItemMap = v
	return s
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListBestPriceItem) Validate() error {
	return dara.Validate(s)
}

type IntlFlightListingSearchResponseBodyModuleFlightItemListBestPriceItemAgreementPriceCodes struct {
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s IntlFlightListingSearchResponseBodyModuleFlightItemListBestPriceItemAgreementPriceCodes) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightListingSearchResponseBodyModuleFlightItemListBestPriceItemAgreementPriceCodes) GoString() string {
	return s.String()
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListBestPriceItemAgreementPriceCodes) GetCode() *string {
	return s.Code
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListBestPriceItemAgreementPriceCodes) GetType() *string {
	return s.Type
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListBestPriceItemAgreementPriceCodes) SetCode(v string) *IntlFlightListingSearchResponseBodyModuleFlightItemListBestPriceItemAgreementPriceCodes {
	s.Code = &v
	return s
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListBestPriceItemAgreementPriceCodes) SetType(v string) *IntlFlightListingSearchResponseBodyModuleFlightItemListBestPriceItemAgreementPriceCodes {
	s.Type = &v
	return s
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListBestPriceItemAgreementPriceCodes) Validate() error {
	return dara.Validate(s)
}

type IntlFlightListingSearchResponseBodyModuleFlightItemListBestPriceItemLabelList struct {
	LabelCode *int32  `json:"labelCode,omitempty" xml:"labelCode,omitempty"`
	LabelName *string `json:"labelName,omitempty" xml:"labelName,omitempty"`
}

func (s IntlFlightListingSearchResponseBodyModuleFlightItemListBestPriceItemLabelList) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightListingSearchResponseBodyModuleFlightItemListBestPriceItemLabelList) GoString() string {
	return s.String()
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListBestPriceItemLabelList) GetLabelCode() *int32 {
	return s.LabelCode
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListBestPriceItemLabelList) GetLabelName() *string {
	return s.LabelName
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListBestPriceItemLabelList) SetLabelCode(v int32) *IntlFlightListingSearchResponseBodyModuleFlightItemListBestPriceItemLabelList {
	s.LabelCode = &v
	return s
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListBestPriceItemLabelList) SetLabelName(v string) *IntlFlightListingSearchResponseBodyModuleFlightItemListBestPriceItemLabelList {
	s.LabelName = &v
	return s
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListBestPriceItemLabelList) Validate() error {
	return dara.Validate(s)
}

type IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfos struct {
	// example:
	//
	// HGH
	ArrCityCode *string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	ArrCityName *string `json:"arr_city_name,omitempty" xml:"arr_city_name,omitempty"`
	// example:
	//
	// 2023-08-13 09:45
	ArrTime *string `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	// example:
	//
	// BJS
	DepCityCode *string `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	DepCityName *string `json:"dep_city_name,omitempty" xml:"dep_city_name,omitempty"`
	// example:
	//
	// 2023-08-13 07:25
	DepTime *string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	// example:
	//
	// 60
	Duration           *int32                                                                                         `json:"duration,omitempty" xml:"duration,omitempty"`
	FlightSegmentInfos []*IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos `json:"flight_segment_infos,omitempty" xml:"flight_segment_infos,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	JourneyIndex *int32 `json:"journey_index,omitempty" xml:"journey_index,omitempty"`
	// example:
	//
	// 0
	TransferTime *int32 `json:"transfer_time,omitempty" xml:"transfer_time,omitempty"`
}

func (s IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfos) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfos) GoString() string {
	return s.String()
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfos) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfos) GetArrCityName() *string {
	return s.ArrCityName
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfos) GetArrTime() *string {
	return s.ArrTime
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfos) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfos) GetDepCityName() *string {
	return s.DepCityName
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfos) GetDepTime() *string {
	return s.DepTime
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfos) GetDuration() *int32 {
	return s.Duration
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfos) GetFlightSegmentInfos() []*IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos {
	return s.FlightSegmentInfos
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfos) GetJourneyIndex() *int32 {
	return s.JourneyIndex
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfos) GetTransferTime() *int32 {
	return s.TransferTime
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfos) SetArrCityCode(v string) *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfos {
	s.ArrCityCode = &v
	return s
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfos) SetArrCityName(v string) *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfos {
	s.ArrCityName = &v
	return s
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfos) SetArrTime(v string) *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfos {
	s.ArrTime = &v
	return s
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfos) SetDepCityCode(v string) *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfos {
	s.DepCityCode = &v
	return s
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfos) SetDepCityName(v string) *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfos {
	s.DepCityName = &v
	return s
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfos) SetDepTime(v string) *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfos {
	s.DepTime = &v
	return s
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfos) SetDuration(v int32) *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfos {
	s.Duration = &v
	return s
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfos) SetFlightSegmentInfos(v []*IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfos {
	s.FlightSegmentInfos = v
	return s
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfos) SetJourneyIndex(v int32) *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfos {
	s.JourneyIndex = &v
	return s
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfos) SetTransferTime(v int32) *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfos {
	s.TransferTime = &v
	return s
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfos) Validate() error {
	return dara.Validate(s)
}

type IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos struct {
	AirlineInfo    *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosAirlineInfo    `json:"airline_info,omitempty" xml:"airline_info,omitempty" type:"Struct"`
	ArrAirportInfo *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosArrAirportInfo `json:"arr_airport_info,omitempty" xml:"arr_airport_info,omitempty" type:"Struct"`
	// example:
	//
	// BJS
	ArrCityCode *string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	ArrCityName *string `json:"arr_city_name,omitempty" xml:"arr_city_name,omitempty"`
	// example:
	//
	// 2023-08-13 09:45
	ArrTime        *string                                                                                                    `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	DepAirportInfo *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosDepAirportInfo `json:"dep_airport_info,omitempty" xml:"dep_airport_info,omitempty" type:"Struct"`
	// example:
	//
	// HGH
	DepCityCode *string `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	DepCityName *string `json:"dep_city_name,omitempty" xml:"dep_city_name,omitempty"`
	// example:
	//
	// 2023-08-13 07:25
	DepTime *string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	// duration
	//
	// example:
	//
	// 140
	Duration *int32 `json:"duration,omitempty" xml:"duration,omitempty"`
	// example:
	//
	// HO1116
	FlightNo           *string                                                                                                          `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
	FlightShareInfo    *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfo      `json:"flight_share_info,omitempty" xml:"flight_share_info,omitempty" type:"Struct"`
	FlightSize         *string                                                                                                          `json:"flight_size,omitempty" xml:"flight_size,omitempty"`
	FlightStopInfoList []*IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoList `json:"flight_stop_info_list,omitempty" xml:"flight_stop_info_list,omitempty" type:"Repeated"`
	// example:
	//
	// 320
	FlightType   *string `json:"flight_type,omitempty" xml:"flight_type,omitempty"`
	Manufacturer *string `json:"manufacturer,omitempty" xml:"manufacturer,omitempty"`
	MealDesc     *string `json:"meal_desc,omitempty" xml:"meal_desc,omitempty"`
	// example:
	//
	// 0
	OneMore     *int32  `json:"one_more,omitempty" xml:"one_more,omitempty"`
	OneMoreShow *string `json:"one_more_show,omitempty" xml:"one_more_show,omitempty"`
	// example:
	//
	// 0
	SegmentIndex *int32 `json:"segment_index,omitempty" xml:"segment_index,omitempty"`
	// example:
	//
	// KN6728HGHPKX0725
	SegmentKey *string `json:"segment_key,omitempty" xml:"segment_key,omitempty"`
	// example:
	//
	// false
	Share           *bool   `json:"share,omitempty" xml:"share,omitempty"`
	ShortFlightSize *string `json:"short_flight_size,omitempty" xml:"short_flight_size,omitempty"`
	// example:
	//
	// false
	Stop      *bool   `json:"stop,omitempty" xml:"stop,omitempty"`
	TotalTime *string `json:"total_time,omitempty" xml:"total_time,omitempty"`
}

func (s IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) GoString() string {
	return s.String()
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) GetAirlineInfo() *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosAirlineInfo {
	return s.AirlineInfo
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) GetArrAirportInfo() *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosArrAirportInfo {
	return s.ArrAirportInfo
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) GetArrCityName() *string {
	return s.ArrCityName
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) GetArrTime() *string {
	return s.ArrTime
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) GetDepAirportInfo() *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosDepAirportInfo {
	return s.DepAirportInfo
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) GetDepCityName() *string {
	return s.DepCityName
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) GetDepTime() *string {
	return s.DepTime
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) GetDuration() *int32 {
	return s.Duration
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) GetFlightNo() *string {
	return s.FlightNo
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) GetFlightShareInfo() *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfo {
	return s.FlightShareInfo
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) GetFlightSize() *string {
	return s.FlightSize
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) GetFlightStopInfoList() []*IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoList {
	return s.FlightStopInfoList
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) GetFlightType() *string {
	return s.FlightType
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) GetManufacturer() *string {
	return s.Manufacturer
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) GetMealDesc() *string {
	return s.MealDesc
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) GetOneMore() *int32 {
	return s.OneMore
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) GetOneMoreShow() *string {
	return s.OneMoreShow
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) GetSegmentIndex() *int32 {
	return s.SegmentIndex
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) GetSegmentKey() *string {
	return s.SegmentKey
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) GetShare() *bool {
	return s.Share
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) GetShortFlightSize() *string {
	return s.ShortFlightSize
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) GetStop() *bool {
	return s.Stop
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) GetTotalTime() *string {
	return s.TotalTime
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) SetAirlineInfo(v *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosAirlineInfo) *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos {
	s.AirlineInfo = v
	return s
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) SetArrAirportInfo(v *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosArrAirportInfo) *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos {
	s.ArrAirportInfo = v
	return s
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) SetArrCityCode(v string) *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos {
	s.ArrCityCode = &v
	return s
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) SetArrCityName(v string) *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos {
	s.ArrCityName = &v
	return s
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) SetArrTime(v string) *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos {
	s.ArrTime = &v
	return s
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) SetDepAirportInfo(v *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosDepAirportInfo) *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos {
	s.DepAirportInfo = v
	return s
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) SetDepCityCode(v string) *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos {
	s.DepCityCode = &v
	return s
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) SetDepCityName(v string) *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos {
	s.DepCityName = &v
	return s
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) SetDepTime(v string) *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos {
	s.DepTime = &v
	return s
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) SetDuration(v int32) *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos {
	s.Duration = &v
	return s
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) SetFlightNo(v string) *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos {
	s.FlightNo = &v
	return s
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) SetFlightShareInfo(v *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfo) *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos {
	s.FlightShareInfo = v
	return s
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) SetFlightSize(v string) *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos {
	s.FlightSize = &v
	return s
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) SetFlightStopInfoList(v []*IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos {
	s.FlightStopInfoList = v
	return s
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) SetFlightType(v string) *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos {
	s.FlightType = &v
	return s
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) SetManufacturer(v string) *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos {
	s.Manufacturer = &v
	return s
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) SetMealDesc(v string) *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos {
	s.MealDesc = &v
	return s
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) SetOneMore(v int32) *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos {
	s.OneMore = &v
	return s
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) SetOneMoreShow(v string) *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos {
	s.OneMoreShow = &v
	return s
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) SetSegmentIndex(v int32) *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos {
	s.SegmentIndex = &v
	return s
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) SetSegmentKey(v string) *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos {
	s.SegmentKey = &v
	return s
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) SetShare(v bool) *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos {
	s.Share = &v
	return s
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) SetShortFlightSize(v string) *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos {
	s.ShortFlightSize = &v
	return s
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) SetStop(v bool) *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos {
	s.Stop = &v
	return s
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) SetTotalTime(v string) *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos {
	s.TotalTime = &v
	return s
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) Validate() error {
	return dara.Validate(s)
}

type IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosAirlineInfo struct {
	// example:
	//
	// 9H
	AirlineCode *string `json:"airline_code,omitempty" xml:"airline_code,omitempty"`
	AirlineName *string `json:"airline_name,omitempty" xml:"airline_name,omitempty"`
	ShortName   *string `json:"short_name,omitempty" xml:"short_name,omitempty"`
}

func (s IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosAirlineInfo) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosAirlineInfo) GoString() string {
	return s.String()
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosAirlineInfo) GetAirlineCode() *string {
	return s.AirlineCode
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosAirlineInfo) GetAirlineName() *string {
	return s.AirlineName
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosAirlineInfo) GetShortName() *string {
	return s.ShortName
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosAirlineInfo) SetAirlineCode(v string) *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosAirlineInfo {
	s.AirlineCode = &v
	return s
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosAirlineInfo) SetAirlineName(v string) *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosAirlineInfo {
	s.AirlineName = &v
	return s
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosAirlineInfo) SetShortName(v string) *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosAirlineInfo {
	s.ShortName = &v
	return s
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosAirlineInfo) Validate() error {
	return dara.Validate(s)
}

type IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosArrAirportInfo struct {
	// example:
	//
	// PKX
	AirportCode      *string `json:"airport_code,omitempty" xml:"airport_code,omitempty"`
	AirportName      *string `json:"airport_name,omitempty" xml:"airport_name,omitempty"`
	AirportShortName *string `json:"airport_short_name,omitempty" xml:"airport_short_name,omitempty"`
	// example:
	//
	// --
	Terminal *string `json:"terminal,omitempty" xml:"terminal,omitempty"`
}

func (s IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosArrAirportInfo) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosArrAirportInfo) GoString() string {
	return s.String()
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosArrAirportInfo) GetAirportCode() *string {
	return s.AirportCode
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosArrAirportInfo) GetAirportName() *string {
	return s.AirportName
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosArrAirportInfo) GetAirportShortName() *string {
	return s.AirportShortName
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosArrAirportInfo) GetTerminal() *string {
	return s.Terminal
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosArrAirportInfo) SetAirportCode(v string) *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosArrAirportInfo {
	s.AirportCode = &v
	return s
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosArrAirportInfo) SetAirportName(v string) *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosArrAirportInfo {
	s.AirportName = &v
	return s
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosArrAirportInfo) SetAirportShortName(v string) *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosArrAirportInfo {
	s.AirportShortName = &v
	return s
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosArrAirportInfo) SetTerminal(v string) *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosArrAirportInfo {
	s.Terminal = &v
	return s
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosArrAirportInfo) Validate() error {
	return dara.Validate(s)
}

type IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosDepAirportInfo struct {
	// example:
	//
	// HGH
	AirportCode      *string `json:"airport_code,omitempty" xml:"airport_code,omitempty"`
	AirportName      *string `json:"airport_name,omitempty" xml:"airport_name,omitempty"`
	AirportShortName *string `json:"airport_short_name,omitempty" xml:"airport_short_name,omitempty"`
	// example:
	//
	// T3
	Terminal *string `json:"terminal,omitempty" xml:"terminal,omitempty"`
}

func (s IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosDepAirportInfo) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosDepAirportInfo) GoString() string {
	return s.String()
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosDepAirportInfo) GetAirportCode() *string {
	return s.AirportCode
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosDepAirportInfo) GetAirportName() *string {
	return s.AirportName
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosDepAirportInfo) GetAirportShortName() *string {
	return s.AirportShortName
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosDepAirportInfo) GetTerminal() *string {
	return s.Terminal
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosDepAirportInfo) SetAirportCode(v string) *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosDepAirportInfo {
	s.AirportCode = &v
	return s
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosDepAirportInfo) SetAirportName(v string) *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosDepAirportInfo {
	s.AirportName = &v
	return s
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosDepAirportInfo) SetAirportShortName(v string) *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosDepAirportInfo {
	s.AirportShortName = &v
	return s
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosDepAirportInfo) SetTerminal(v string) *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosDepAirportInfo {
	s.Terminal = &v
	return s
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosDepAirportInfo) Validate() error {
	return dara.Validate(s)
}

type IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfo struct {
	OperatingAirlineInfo *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo `json:"operating_airline_info,omitempty" xml:"operating_airline_info,omitempty" type:"Struct"`
	// example:
	//
	// CX601
	OperatingFlightNo *string `json:"operating_flight_no,omitempty" xml:"operating_flight_no,omitempty"`
}

func (s IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfo) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfo) GoString() string {
	return s.String()
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfo) GetOperatingAirlineInfo() *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo {
	return s.OperatingAirlineInfo
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfo) GetOperatingFlightNo() *string {
	return s.OperatingFlightNo
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfo) SetOperatingAirlineInfo(v *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfo {
	s.OperatingAirlineInfo = v
	return s
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfo) SetOperatingFlightNo(v string) *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfo {
	s.OperatingFlightNo = &v
	return s
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfo) Validate() error {
	return dara.Validate(s)
}

type IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo struct {
	// example:
	//
	// DR
	AirlineCode *string `json:"airline_code,omitempty" xml:"airline_code,omitempty"`
	AirlineName *string `json:"airline_name,omitempty" xml:"airline_name,omitempty"`
	ShortName   *string `json:"short_name,omitempty" xml:"short_name,omitempty"`
}

func (s IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) GoString() string {
	return s.String()
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) GetAirlineCode() *string {
	return s.AirlineCode
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) GetAirlineName() *string {
	return s.AirlineName
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) GetShortName() *string {
	return s.ShortName
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) SetAirlineCode(v string) *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo {
	s.AirlineCode = &v
	return s
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) SetAirlineName(v string) *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo {
	s.AirlineName = &v
	return s
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) SetShortName(v string) *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo {
	s.ShortName = &v
	return s
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) Validate() error {
	return dara.Validate(s)
}

type IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoList struct {
	StopAirport     *string `json:"stop_airport,omitempty" xml:"stop_airport,omitempty"`
	StopAirportName *string `json:"stop_airport_name,omitempty" xml:"stop_airport_name,omitempty"`
	StopArrTerm     *string `json:"stop_arr_term,omitempty" xml:"stop_arr_term,omitempty"`
	StopArrTime     *string `json:"stop_arr_time,omitempty" xml:"stop_arr_time,omitempty"`
	StopCityCode    *string `json:"stop_city_code,omitempty" xml:"stop_city_code,omitempty"`
	StopCityName    *string `json:"stop_city_name,omitempty" xml:"stop_city_name,omitempty"`
	StopDepTerm     *string `json:"stop_dep_term,omitempty" xml:"stop_dep_term,omitempty"`
	StopDepTime     *string `json:"stop_dep_time,omitempty" xml:"stop_dep_time,omitempty"`
	StopTime        *string `json:"stop_time,omitempty" xml:"stop_time,omitempty"`
}

func (s IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) GoString() string {
	return s.String()
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) GetStopAirport() *string {
	return s.StopAirport
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) GetStopAirportName() *string {
	return s.StopAirportName
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) GetStopArrTerm() *string {
	return s.StopArrTerm
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) GetStopArrTime() *string {
	return s.StopArrTime
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) GetStopCityCode() *string {
	return s.StopCityCode
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) GetStopCityName() *string {
	return s.StopCityName
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) GetStopDepTerm() *string {
	return s.StopDepTerm
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) GetStopDepTime() *string {
	return s.StopDepTime
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) GetStopTime() *string {
	return s.StopTime
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) SetStopAirport(v string) *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoList {
	s.StopAirport = &v
	return s
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) SetStopAirportName(v string) *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoList {
	s.StopAirportName = &v
	return s
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) SetStopArrTerm(v string) *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoList {
	s.StopArrTerm = &v
	return s
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) SetStopArrTime(v string) *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoList {
	s.StopArrTime = &v
	return s
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) SetStopCityCode(v string) *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoList {
	s.StopCityCode = &v
	return s
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) SetStopCityName(v string) *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoList {
	s.StopCityName = &v
	return s
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) SetStopDepTerm(v string) *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoList {
	s.StopDepTerm = &v
	return s
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) SetStopDepTime(v string) *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoList {
	s.StopDepTime = &v
	return s
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) SetStopTime(v string) *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoList {
	s.StopTime = &v
	return s
}

func (s *IntlFlightListingSearchResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) Validate() error {
	return dara.Validate(s)
}
