// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightListingSearchV2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *FlightListingSearchV2ResponseBody
	GetCode() *string
	SetMessage(v string) *FlightListingSearchV2ResponseBody
	GetMessage() *string
	SetModule(v *FlightListingSearchV2ResponseBodyModule) *FlightListingSearchV2ResponseBody
	GetModule() *FlightListingSearchV2ResponseBodyModule
	SetRequestId(v string) *FlightListingSearchV2ResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *FlightListingSearchV2ResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *FlightListingSearchV2ResponseBody
	GetTraceId() *string
}

type FlightListingSearchV2ResponseBody struct {
	// example:
	//
	// 200
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// module
	Module *FlightListingSearchV2ResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// requestId
	//
	// example:
	//
	// 407543AF-2BD9-5890-BD92-9D1AB7218B27
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// traceId
	//
	// example:
	//
	// 210e842b16611337974412836dae27
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s FlightListingSearchV2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FlightListingSearchV2ResponseBody) GoString() string {
	return s.String()
}

func (s *FlightListingSearchV2ResponseBody) GetCode() *string {
	return s.Code
}

func (s *FlightListingSearchV2ResponseBody) GetMessage() *string {
	return s.Message
}

func (s *FlightListingSearchV2ResponseBody) GetModule() *FlightListingSearchV2ResponseBodyModule {
	return s.Module
}

func (s *FlightListingSearchV2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FlightListingSearchV2ResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *FlightListingSearchV2ResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *FlightListingSearchV2ResponseBody) SetCode(v string) *FlightListingSearchV2ResponseBody {
	s.Code = &v
	return s
}

func (s *FlightListingSearchV2ResponseBody) SetMessage(v string) *FlightListingSearchV2ResponseBody {
	s.Message = &v
	return s
}

func (s *FlightListingSearchV2ResponseBody) SetModule(v *FlightListingSearchV2ResponseBodyModule) *FlightListingSearchV2ResponseBody {
	s.Module = v
	return s
}

func (s *FlightListingSearchV2ResponseBody) SetRequestId(v string) *FlightListingSearchV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *FlightListingSearchV2ResponseBody) SetSuccess(v bool) *FlightListingSearchV2ResponseBody {
	s.Success = &v
	return s
}

func (s *FlightListingSearchV2ResponseBody) SetTraceId(v string) *FlightListingSearchV2ResponseBody {
	s.TraceId = &v
	return s
}

func (s *FlightListingSearchV2ResponseBody) Validate() error {
	return dara.Validate(s)
}

type FlightListingSearchV2ResponseBodyModule struct {
	FlightItemList []*FlightListingSearchV2ResponseBodyModuleFlightItemList `json:"flight_item_list,omitempty" xml:"flight_item_list,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	SearchMode *int32 `json:"search_mode,omitempty" xml:"search_mode,omitempty"`
	// example:
	//
	// 1
	TripType *int32 `json:"trip_type,omitempty" xml:"trip_type,omitempty"`
}

func (s FlightListingSearchV2ResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s FlightListingSearchV2ResponseBodyModule) GoString() string {
	return s.String()
}

func (s *FlightListingSearchV2ResponseBodyModule) GetFlightItemList() []*FlightListingSearchV2ResponseBodyModuleFlightItemList {
	return s.FlightItemList
}

func (s *FlightListingSearchV2ResponseBodyModule) GetSearchMode() *int32 {
	return s.SearchMode
}

func (s *FlightListingSearchV2ResponseBodyModule) GetTripType() *int32 {
	return s.TripType
}

func (s *FlightListingSearchV2ResponseBodyModule) SetFlightItemList(v []*FlightListingSearchV2ResponseBodyModuleFlightItemList) *FlightListingSearchV2ResponseBodyModule {
	s.FlightItemList = v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModule) SetSearchMode(v int32) *FlightListingSearchV2ResponseBodyModule {
	s.SearchMode = &v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModule) SetTripType(v int32) *FlightListingSearchV2ResponseBodyModule {
	s.TripType = &v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type FlightListingSearchV2ResponseBodyModuleFlightItemList struct {
	BestPriceItem      *FlightListingSearchV2ResponseBodyModuleFlightItemListBestPriceItem        `json:"best_price_item,omitempty" xml:"best_price_item,omitempty" type:"Struct"`
	FlightJourneyInfos []*FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfos `json:"flight_journey_infos,omitempty" xml:"flight_journey_infos,omitempty" type:"Repeated"`
	ItemList           []*FlightListingSearchV2ResponseBodyModuleFlightItemListItemList           `json:"item_list,omitempty" xml:"item_list,omitempty" type:"Repeated"`
}

func (s FlightListingSearchV2ResponseBodyModuleFlightItemList) String() string {
	return dara.Prettify(s)
}

func (s FlightListingSearchV2ResponseBodyModuleFlightItemList) GoString() string {
	return s.String()
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemList) GetBestPriceItem() *FlightListingSearchV2ResponseBodyModuleFlightItemListBestPriceItem {
	return s.BestPriceItem
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemList) GetFlightJourneyInfos() []*FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfos {
	return s.FlightJourneyInfos
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemList) GetItemList() []*FlightListingSearchV2ResponseBodyModuleFlightItemListItemList {
	return s.ItemList
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemList) SetBestPriceItem(v *FlightListingSearchV2ResponseBodyModuleFlightItemListBestPriceItem) *FlightListingSearchV2ResponseBodyModuleFlightItemList {
	s.BestPriceItem = v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemList) SetFlightJourneyInfos(v []*FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfos) *FlightListingSearchV2ResponseBodyModuleFlightItemList {
	s.FlightJourneyInfos = v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemList) SetItemList(v []*FlightListingSearchV2ResponseBodyModuleFlightItemListItemList) *FlightListingSearchV2ResponseBodyModuleFlightItemList {
	s.ItemList = v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemList) Validate() error {
	return dara.Validate(s)
}

type FlightListingSearchV2ResponseBodyModuleFlightItemListBestPriceItem struct {
	FlightRuleInfos map[string]*ModuleFlightItemListBestPriceItemFlightRuleInfosValue `json:"flight_rule_infos,omitempty" xml:"flight_rule_infos,omitempty"`
	// example:
	//
	// e50d380fc05942cc8ac57af8ae02f448_0
	ItemId             *string                                                                       `json:"item_id,omitempty" xml:"item_id,omitempty"`
	ShoppingItemMap    map[string]*ModuleFlightItemListBestPriceItemShoppingItemMapValue             `json:"shopping_item_map,omitempty" xml:"shopping_item_map,omitempty"`
	SubItemPositionMap map[string][]*ModuleFlightItemListBestPriceItemSubItemPositionMapValue        `json:"sub_item_position_map,omitempty" xml:"sub_item_position_map,omitempty"`
	SubItems           []*FlightListingSearchV2ResponseBodyModuleFlightItemListBestPriceItemSubItems `json:"sub_items,omitempty" xml:"sub_items,omitempty" type:"Repeated"`
}

func (s FlightListingSearchV2ResponseBodyModuleFlightItemListBestPriceItem) String() string {
	return dara.Prettify(s)
}

func (s FlightListingSearchV2ResponseBodyModuleFlightItemListBestPriceItem) GoString() string {
	return s.String()
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListBestPriceItem) GetFlightRuleInfos() map[string]*ModuleFlightItemListBestPriceItemFlightRuleInfosValue {
	return s.FlightRuleInfos
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListBestPriceItem) GetItemId() *string {
	return s.ItemId
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListBestPriceItem) GetShoppingItemMap() map[string]*ModuleFlightItemListBestPriceItemShoppingItemMapValue {
	return s.ShoppingItemMap
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListBestPriceItem) GetSubItemPositionMap() map[string][]*ModuleFlightItemListBestPriceItemSubItemPositionMapValue {
	return s.SubItemPositionMap
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListBestPriceItem) GetSubItems() []*FlightListingSearchV2ResponseBodyModuleFlightItemListBestPriceItemSubItems {
	return s.SubItems
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListBestPriceItem) SetFlightRuleInfos(v map[string]*ModuleFlightItemListBestPriceItemFlightRuleInfosValue) *FlightListingSearchV2ResponseBodyModuleFlightItemListBestPriceItem {
	s.FlightRuleInfos = v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListBestPriceItem) SetItemId(v string) *FlightListingSearchV2ResponseBodyModuleFlightItemListBestPriceItem {
	s.ItemId = &v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListBestPriceItem) SetShoppingItemMap(v map[string]*ModuleFlightItemListBestPriceItemShoppingItemMapValue) *FlightListingSearchV2ResponseBodyModuleFlightItemListBestPriceItem {
	s.ShoppingItemMap = v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListBestPriceItem) SetSubItemPositionMap(v map[string][]*ModuleFlightItemListBestPriceItemSubItemPositionMapValue) *FlightListingSearchV2ResponseBodyModuleFlightItemListBestPriceItem {
	s.SubItemPositionMap = v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListBestPriceItem) SetSubItems(v []*FlightListingSearchV2ResponseBodyModuleFlightItemListBestPriceItemSubItems) *FlightListingSearchV2ResponseBodyModuleFlightItemListBestPriceItem {
	s.SubItems = v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListBestPriceItem) Validate() error {
	return dara.Validate(s)
}

type FlightListingSearchV2ResponseBodyModuleFlightItemListBestPriceItemSubItems struct {
	ShoppingItemMap map[string]*ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValue `json:"shopping_item_map,omitempty" xml:"shopping_item_map,omitempty"`
	// example:
	//
	// 07df0bd9-f803-4a50-8449-f4bd675d9939
	UniqKey *string `json:"uniq_key,omitempty" xml:"uniq_key,omitempty"`
}

func (s FlightListingSearchV2ResponseBodyModuleFlightItemListBestPriceItemSubItems) String() string {
	return dara.Prettify(s)
}

func (s FlightListingSearchV2ResponseBodyModuleFlightItemListBestPriceItemSubItems) GoString() string {
	return s.String()
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListBestPriceItemSubItems) GetShoppingItemMap() map[string]*ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValue {
	return s.ShoppingItemMap
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListBestPriceItemSubItems) GetUniqKey() *string {
	return s.UniqKey
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListBestPriceItemSubItems) SetShoppingItemMap(v map[string]*ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValue) *FlightListingSearchV2ResponseBodyModuleFlightItemListBestPriceItemSubItems {
	s.ShoppingItemMap = v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListBestPriceItemSubItems) SetUniqKey(v string) *FlightListingSearchV2ResponseBodyModuleFlightItemListBestPriceItemSubItems {
	s.UniqKey = &v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListBestPriceItemSubItems) Validate() error {
	return dara.Validate(s)
}

type FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfos struct {
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
	// 140
	Duration           *int32                                                                                       `json:"duration,omitempty" xml:"duration,omitempty"`
	Extensions         map[string]*string                                                                           `json:"extensions,omitempty" xml:"extensions,omitempty"`
	FlightSegmentInfos []*FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos `json:"flight_segment_infos,omitempty" xml:"flight_segment_infos,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	JourneyIndex *int32 `json:"journey_index,omitempty" xml:"journey_index,omitempty"`
	// example:
	//
	// 0
	TransferTime *int32 `json:"transfer_time,omitempty" xml:"transfer_time,omitempty"`
}

func (s FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfos) String() string {
	return dara.Prettify(s)
}

func (s FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfos) GoString() string {
	return s.String()
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfos) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfos) GetArrCityName() *string {
	return s.ArrCityName
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfos) GetArrTime() *string {
	return s.ArrTime
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfos) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfos) GetDepCityName() *string {
	return s.DepCityName
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfos) GetDepTime() *string {
	return s.DepTime
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfos) GetDuration() *int32 {
	return s.Duration
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfos) GetExtensions() map[string]*string {
	return s.Extensions
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfos) GetFlightSegmentInfos() []*FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos {
	return s.FlightSegmentInfos
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfos) GetJourneyIndex() *int32 {
	return s.JourneyIndex
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfos) GetTransferTime() *int32 {
	return s.TransferTime
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfos) SetArrCityCode(v string) *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfos {
	s.ArrCityCode = &v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfos) SetArrCityName(v string) *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfos {
	s.ArrCityName = &v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfos) SetArrTime(v string) *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfos {
	s.ArrTime = &v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfos) SetDepCityCode(v string) *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfos {
	s.DepCityCode = &v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfos) SetDepCityName(v string) *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfos {
	s.DepCityName = &v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfos) SetDepTime(v string) *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfos {
	s.DepTime = &v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfos) SetDuration(v int32) *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfos {
	s.Duration = &v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfos) SetExtensions(v map[string]*string) *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfos {
	s.Extensions = v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfos) SetFlightSegmentInfos(v []*FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfos {
	s.FlightSegmentInfos = v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfos) SetJourneyIndex(v int32) *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfos {
	s.JourneyIndex = &v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfos) SetTransferTime(v int32) *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfos {
	s.TransferTime = &v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfos) Validate() error {
	return dara.Validate(s)
}

type FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos struct {
	AirlineInfo    *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosAirlineInfo    `json:"airline_info,omitempty" xml:"airline_info,omitempty" type:"Struct"`
	ArrAirportInfo *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosArrAirportInfo `json:"arr_airport_info,omitempty" xml:"arr_airport_info,omitempty" type:"Struct"`
	// example:
	//
	// BJS
	ArrCityCode *string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	ArrCityName *string `json:"arr_city_name,omitempty" xml:"arr_city_name,omitempty"`
	// example:
	//
	// 2023-08-13 09:45
	ArrTime        *string                                                                                                  `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	BaggageDesc    *string                                                                                                  `json:"baggage_desc,omitempty" xml:"baggage_desc,omitempty"`
	DepAirportInfo *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosDepAirportInfo `json:"dep_airport_info,omitempty" xml:"dep_airport_info,omitempty" type:"Struct"`
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
	Duration  *int32                 `json:"duration,omitempty" xml:"duration,omitempty"`
	ExtraInfo map[string]interface{} `json:"extra_info,omitempty" xml:"extra_info,omitempty"`
	// example:
	//
	// MU5131
	FlightNo        *string                                                                                                   `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
	FlightShareInfo *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfo `json:"flight_share_info,omitempty" xml:"flight_share_info,omitempty" type:"Struct"`
	FlightSize      *string                                                                                                   `json:"flight_size,omitempty" xml:"flight_size,omitempty"`
	FlightStopInfo  *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfo  `json:"flight_stop_info,omitempty" xml:"flight_stop_info,omitempty" type:"Struct"`
	// example:
	//
	// 320
	FlightType   *string `json:"flight_type,omitempty" xml:"flight_type,omitempty"`
	Manufacturer *string `json:"manufacturer,omitempty" xml:"manufacturer,omitempty"`
	MealDesc     *string `json:"meal_desc,omitempty" xml:"meal_desc,omitempty"`
	Miles        *int32  `json:"miles,omitempty" xml:"miles,omitempty"`
	OnTimeRate   *string `json:"on_time_rate,omitempty" xml:"on_time_rate,omitempty"`
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
	// false
	Share           *bool   `json:"share,omitempty" xml:"share,omitempty"`
	ShortFlightSize *string `json:"short_flight_size,omitempty" xml:"short_flight_size,omitempty"`
	// example:
	//
	// false
	Stop               *bool   `json:"stop,omitempty" xml:"stop,omitempty"`
	TotalTime          *string `json:"total_time,omitempty" xml:"total_time,omitempty"`
	TransferTime       *string `json:"transfer_time,omitempty" xml:"transfer_time,omitempty"`
	TransferTimeNumber *int32  `json:"transfer_time_number,omitempty" xml:"transfer_time_number,omitempty"`
}

func (s FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) String() string {
	return dara.Prettify(s)
}

func (s FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) GoString() string {
	return s.String()
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) GetAirlineInfo() *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosAirlineInfo {
	return s.AirlineInfo
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) GetArrAirportInfo() *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosArrAirportInfo {
	return s.ArrAirportInfo
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) GetArrCityName() *string {
	return s.ArrCityName
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) GetArrTime() *string {
	return s.ArrTime
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) GetBaggageDesc() *string {
	return s.BaggageDesc
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) GetDepAirportInfo() *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosDepAirportInfo {
	return s.DepAirportInfo
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) GetDepCityName() *string {
	return s.DepCityName
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) GetDepTime() *string {
	return s.DepTime
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) GetDuration() *int32 {
	return s.Duration
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) GetExtraInfo() map[string]interface{} {
	return s.ExtraInfo
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) GetFlightNo() *string {
	return s.FlightNo
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) GetFlightShareInfo() *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfo {
	return s.FlightShareInfo
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) GetFlightSize() *string {
	return s.FlightSize
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) GetFlightStopInfo() *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfo {
	return s.FlightStopInfo
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) GetFlightType() *string {
	return s.FlightType
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) GetManufacturer() *string {
	return s.Manufacturer
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) GetMealDesc() *string {
	return s.MealDesc
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) GetMiles() *int32 {
	return s.Miles
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) GetOnTimeRate() *string {
	return s.OnTimeRate
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) GetOneMore() *int32 {
	return s.OneMore
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) GetOneMoreShow() *string {
	return s.OneMoreShow
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) GetSegmentIndex() *int32 {
	return s.SegmentIndex
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) GetShare() *bool {
	return s.Share
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) GetShortFlightSize() *string {
	return s.ShortFlightSize
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) GetStop() *bool {
	return s.Stop
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) GetTotalTime() *string {
	return s.TotalTime
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) GetTransferTime() *string {
	return s.TransferTime
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) GetTransferTimeNumber() *int32 {
	return s.TransferTimeNumber
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) SetAirlineInfo(v *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosAirlineInfo) *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos {
	s.AirlineInfo = v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) SetArrAirportInfo(v *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosArrAirportInfo) *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos {
	s.ArrAirportInfo = v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) SetArrCityCode(v string) *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos {
	s.ArrCityCode = &v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) SetArrCityName(v string) *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos {
	s.ArrCityName = &v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) SetArrTime(v string) *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos {
	s.ArrTime = &v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) SetBaggageDesc(v string) *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos {
	s.BaggageDesc = &v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) SetDepAirportInfo(v *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosDepAirportInfo) *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos {
	s.DepAirportInfo = v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) SetDepCityCode(v string) *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos {
	s.DepCityCode = &v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) SetDepCityName(v string) *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos {
	s.DepCityName = &v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) SetDepTime(v string) *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos {
	s.DepTime = &v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) SetDuration(v int32) *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos {
	s.Duration = &v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) SetExtraInfo(v map[string]interface{}) *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos {
	s.ExtraInfo = v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) SetFlightNo(v string) *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos {
	s.FlightNo = &v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) SetFlightShareInfo(v *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfo) *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos {
	s.FlightShareInfo = v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) SetFlightSize(v string) *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos {
	s.FlightSize = &v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) SetFlightStopInfo(v *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfo) *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos {
	s.FlightStopInfo = v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) SetFlightType(v string) *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos {
	s.FlightType = &v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) SetManufacturer(v string) *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos {
	s.Manufacturer = &v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) SetMealDesc(v string) *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos {
	s.MealDesc = &v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) SetMiles(v int32) *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos {
	s.Miles = &v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) SetOnTimeRate(v string) *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos {
	s.OnTimeRate = &v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) SetOneMore(v int32) *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos {
	s.OneMore = &v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) SetOneMoreShow(v string) *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos {
	s.OneMoreShow = &v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) SetSegmentIndex(v int32) *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos {
	s.SegmentIndex = &v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) SetShare(v bool) *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos {
	s.Share = &v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) SetShortFlightSize(v string) *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos {
	s.ShortFlightSize = &v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) SetStop(v bool) *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos {
	s.Stop = &v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) SetTotalTime(v string) *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos {
	s.TotalTime = &v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) SetTransferTime(v string) *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos {
	s.TransferTime = &v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) SetTransferTimeNumber(v int32) *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos {
	s.TransferTimeNumber = &v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfos) Validate() error {
	return dara.Validate(s)
}

type FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosAirlineInfo struct {
	AirlineChineseName      *string `json:"airline_chinese_name,omitempty" xml:"airline_chinese_name,omitempty"`
	AirlineChineseShortName *string `json:"airline_chinese_short_name,omitempty" xml:"airline_chinese_short_name,omitempty"`
	// example:
	//
	// MU
	AirlineCode *string `json:"airline_code,omitempty" xml:"airline_code,omitempty"`
	// example:
	//
	// https://gw.alicdn.com/tfs/TB15EXDFHY1gK0jSZTEXXXDQVXa-450-450.png_80x80.jpg
	AirlineIcon *string `json:"airline_icon,omitempty" xml:"airline_icon,omitempty"`
	// example:
	//
	// false
	CheapFlight *bool `json:"cheap_flight,omitempty" xml:"cheap_flight,omitempty"`
}

func (s FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosAirlineInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosAirlineInfo) GoString() string {
	return s.String()
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosAirlineInfo) GetAirlineChineseName() *string {
	return s.AirlineChineseName
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosAirlineInfo) GetAirlineChineseShortName() *string {
	return s.AirlineChineseShortName
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosAirlineInfo) GetAirlineCode() *string {
	return s.AirlineCode
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosAirlineInfo) GetAirlineIcon() *string {
	return s.AirlineIcon
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosAirlineInfo) GetCheapFlight() *bool {
	return s.CheapFlight
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosAirlineInfo) SetAirlineChineseName(v string) *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosAirlineInfo {
	s.AirlineChineseName = &v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosAirlineInfo) SetAirlineChineseShortName(v string) *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosAirlineInfo {
	s.AirlineChineseShortName = &v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosAirlineInfo) SetAirlineCode(v string) *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosAirlineInfo {
	s.AirlineCode = &v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosAirlineInfo) SetAirlineIcon(v string) *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosAirlineInfo {
	s.AirlineIcon = &v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosAirlineInfo) SetCheapFlight(v bool) *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosAirlineInfo {
	s.CheapFlight = &v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosAirlineInfo) Validate() error {
	return dara.Validate(s)
}

type FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosArrAirportInfo struct {
	// example:
	//
	// PKX
	AirportCode *string `json:"airport_code,omitempty" xml:"airport_code,omitempty"`
	AirportName *string `json:"airport_name,omitempty" xml:"airport_name,omitempty"`
	// example:
	//
	// #FF7300
	AirportNameColor *string `json:"airport_name_color,omitempty" xml:"airport_name_color,omitempty"`
	AirportShortName *string `json:"airport_short_name,omitempty" xml:"airport_short_name,omitempty"`
	// example:
	//
	// --
	Terminal *string `json:"terminal,omitempty" xml:"terminal,omitempty"`
}

func (s FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosArrAirportInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosArrAirportInfo) GoString() string {
	return s.String()
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosArrAirportInfo) GetAirportCode() *string {
	return s.AirportCode
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosArrAirportInfo) GetAirportName() *string {
	return s.AirportName
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosArrAirportInfo) GetAirportNameColor() *string {
	return s.AirportNameColor
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosArrAirportInfo) GetAirportShortName() *string {
	return s.AirportShortName
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosArrAirportInfo) GetTerminal() *string {
	return s.Terminal
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosArrAirportInfo) SetAirportCode(v string) *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosArrAirportInfo {
	s.AirportCode = &v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosArrAirportInfo) SetAirportName(v string) *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosArrAirportInfo {
	s.AirportName = &v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosArrAirportInfo) SetAirportNameColor(v string) *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosArrAirportInfo {
	s.AirportNameColor = &v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosArrAirportInfo) SetAirportShortName(v string) *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosArrAirportInfo {
	s.AirportShortName = &v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosArrAirportInfo) SetTerminal(v string) *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosArrAirportInfo {
	s.Terminal = &v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosArrAirportInfo) Validate() error {
	return dara.Validate(s)
}

type FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosDepAirportInfo struct {
	// example:
	//
	// HGH
	AirportCode *string `json:"airport_code,omitempty" xml:"airport_code,omitempty"`
	AirportName *string `json:"airport_name,omitempty" xml:"airport_name,omitempty"`
	// example:
	//
	// #FF7300
	AirportNameColor *string `json:"airport_name_color,omitempty" xml:"airport_name_color,omitempty"`
	AirportShortName *string `json:"airport_short_name,omitempty" xml:"airport_short_name,omitempty"`
	// example:
	//
	// T3
	Terminal *string `json:"terminal,omitempty" xml:"terminal,omitempty"`
}

func (s FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosDepAirportInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosDepAirportInfo) GoString() string {
	return s.String()
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosDepAirportInfo) GetAirportCode() *string {
	return s.AirportCode
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosDepAirportInfo) GetAirportName() *string {
	return s.AirportName
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosDepAirportInfo) GetAirportNameColor() *string {
	return s.AirportNameColor
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosDepAirportInfo) GetAirportShortName() *string {
	return s.AirportShortName
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosDepAirportInfo) GetTerminal() *string {
	return s.Terminal
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosDepAirportInfo) SetAirportCode(v string) *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosDepAirportInfo {
	s.AirportCode = &v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosDepAirportInfo) SetAirportName(v string) *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosDepAirportInfo {
	s.AirportName = &v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosDepAirportInfo) SetAirportNameColor(v string) *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosDepAirportInfo {
	s.AirportNameColor = &v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosDepAirportInfo) SetAirportShortName(v string) *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosDepAirportInfo {
	s.AirportShortName = &v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosDepAirportInfo) SetTerminal(v string) *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosDepAirportInfo {
	s.Terminal = &v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosDepAirportInfo) Validate() error {
	return dara.Validate(s)
}

type FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfo struct {
	OperatingAirlineInfo *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo `json:"operating_airline_info,omitempty" xml:"operating_airline_info,omitempty" type:"Struct"`
	// example:
	//
	// CX601
	OperatingFlightNo *string `json:"operating_flight_no,omitempty" xml:"operating_flight_no,omitempty"`
}

func (s FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfo) GoString() string {
	return s.String()
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfo) GetOperatingAirlineInfo() *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo {
	return s.OperatingAirlineInfo
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfo) GetOperatingFlightNo() *string {
	return s.OperatingFlightNo
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfo) SetOperatingAirlineInfo(v *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfo {
	s.OperatingAirlineInfo = v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfo) SetOperatingFlightNo(v string) *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfo {
	s.OperatingFlightNo = &v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfo) Validate() error {
	return dara.Validate(s)
}

type FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo struct {
	AirlineChineseName      *string `json:"airline_chinese_name,omitempty" xml:"airline_chinese_name,omitempty"`
	AirlineChineseShortName *string `json:"airline_chinese_short_name,omitempty" xml:"airline_chinese_short_name,omitempty"`
	// example:
	//
	// DR
	AirlineCode *string `json:"airline_code,omitempty" xml:"airline_code,omitempty"`
	// example:
	//
	// https://gw.alicdn.com/tfs/TB1gSXBFUT1gK0jSZFhXXaAtVXa-450-450.png_80x80.jpg
	AirlineIcon *string `json:"airline_icon,omitempty" xml:"airline_icon,omitempty"`
	// example:
	//
	// false
	CheapFlight *bool `json:"cheap_flight,omitempty" xml:"cheap_flight,omitempty"`
}

func (s FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) GoString() string {
	return s.String()
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) GetAirlineChineseName() *string {
	return s.AirlineChineseName
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) GetAirlineChineseShortName() *string {
	return s.AirlineChineseShortName
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) GetAirlineCode() *string {
	return s.AirlineCode
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) GetAirlineIcon() *string {
	return s.AirlineIcon
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) GetCheapFlight() *bool {
	return s.CheapFlight
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) SetAirlineChineseName(v string) *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo {
	s.AirlineChineseName = &v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) SetAirlineChineseShortName(v string) *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo {
	s.AirlineChineseShortName = &v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) SetAirlineCode(v string) *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo {
	s.AirlineCode = &v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) SetAirlineIcon(v string) *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo {
	s.AirlineIcon = &v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) SetCheapFlight(v bool) *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo {
	s.CheapFlight = &v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) Validate() error {
	return dara.Validate(s)
}

type FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfo struct {
	// example:
	//
	// HGH
	StopAirport     *string `json:"stop_airport,omitempty" xml:"stop_airport,omitempty"`
	StopAirportName *string `json:"stop_airport_name,omitempty" xml:"stop_airport_name,omitempty"`
	// example:
	//
	// T3
	StopArrTerm *string `json:"stop_arr_term,omitempty" xml:"stop_arr_term,omitempty"`
	// example:
	//
	// BJS
	StopArrTime *string `json:"stop_arr_time,omitempty" xml:"stop_arr_time,omitempty"`
	// example:
	//
	// HGH
	StopCityCode *string `json:"stop_city_code,omitempty" xml:"stop_city_code,omitempty"`
	// example:
	//
	// HGH
	StopCityName  *string   `json:"stop_city_name,omitempty" xml:"stop_city_name,omitempty"`
	StopCityNames []*string `json:"stop_city_names,omitempty" xml:"stop_city_names,omitempty" type:"Repeated"`
	// example:
	//
	// T4
	StopDepTerm *string `json:"stop_dep_term,omitempty" xml:"stop_dep_term,omitempty"`
	// example:
	//
	// 2023-08-13 07:25
	StopDepTime *string `json:"stop_dep_time,omitempty" xml:"stop_dep_time,omitempty"`
	StopTime    *string `json:"stop_time,omitempty" xml:"stop_time,omitempty"`
}

func (s FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfo) GoString() string {
	return s.String()
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfo) GetStopAirport() *string {
	return s.StopAirport
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfo) GetStopAirportName() *string {
	return s.StopAirportName
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfo) GetStopArrTerm() *string {
	return s.StopArrTerm
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfo) GetStopArrTime() *string {
	return s.StopArrTime
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfo) GetStopCityCode() *string {
	return s.StopCityCode
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfo) GetStopCityName() *string {
	return s.StopCityName
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfo) GetStopCityNames() []*string {
	return s.StopCityNames
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfo) GetStopDepTerm() *string {
	return s.StopDepTerm
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfo) GetStopDepTime() *string {
	return s.StopDepTime
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfo) GetStopTime() *string {
	return s.StopTime
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfo) SetStopAirport(v string) *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfo {
	s.StopAirport = &v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfo) SetStopAirportName(v string) *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfo {
	s.StopAirportName = &v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfo) SetStopArrTerm(v string) *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfo {
	s.StopArrTerm = &v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfo) SetStopArrTime(v string) *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfo {
	s.StopArrTime = &v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfo) SetStopCityCode(v string) *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfo {
	s.StopCityCode = &v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfo) SetStopCityName(v string) *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfo {
	s.StopCityName = &v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfo) SetStopCityNames(v []*string) *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfo {
	s.StopCityNames = v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfo) SetStopDepTerm(v string) *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfo {
	s.StopDepTerm = &v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfo) SetStopDepTime(v string) *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfo {
	s.StopDepTime = &v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfo) SetStopTime(v string) *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfo {
	s.StopTime = &v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfo) Validate() error {
	return dara.Validate(s)
}

type FlightListingSearchV2ResponseBodyModuleFlightItemListItemList struct {
	FlightRuleInfos map[string]*ModuleFlightItemListItemListFlightRuleInfosValue `json:"flight_rule_infos,omitempty" xml:"flight_rule_infos,omitempty"`
	// example:
	//
	// e50d380fc05942cc8ac57af8ae02f448_0
	ItemId             *string                                                                  `json:"item_id,omitempty" xml:"item_id,omitempty"`
	ShoppingItemMap    map[string]*ModuleFlightItemListItemListShoppingItemMapValue             `json:"shopping_item_map,omitempty" xml:"shopping_item_map,omitempty"`
	SubItemPositionMap map[string][]*ModuleFlightItemListItemListSubItemPositionMapValue        `json:"sub_item_position_map,omitempty" xml:"sub_item_position_map,omitempty"`
	SubItems           []*FlightListingSearchV2ResponseBodyModuleFlightItemListItemListSubItems `json:"sub_items,omitempty" xml:"sub_items,omitempty" type:"Repeated"`
}

func (s FlightListingSearchV2ResponseBodyModuleFlightItemListItemList) String() string {
	return dara.Prettify(s)
}

func (s FlightListingSearchV2ResponseBodyModuleFlightItemListItemList) GoString() string {
	return s.String()
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListItemList) GetFlightRuleInfos() map[string]*ModuleFlightItemListItemListFlightRuleInfosValue {
	return s.FlightRuleInfos
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListItemList) GetItemId() *string {
	return s.ItemId
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListItemList) GetShoppingItemMap() map[string]*ModuleFlightItemListItemListShoppingItemMapValue {
	return s.ShoppingItemMap
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListItemList) GetSubItemPositionMap() map[string][]*ModuleFlightItemListItemListSubItemPositionMapValue {
	return s.SubItemPositionMap
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListItemList) GetSubItems() []*FlightListingSearchV2ResponseBodyModuleFlightItemListItemListSubItems {
	return s.SubItems
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListItemList) SetFlightRuleInfos(v map[string]*ModuleFlightItemListItemListFlightRuleInfosValue) *FlightListingSearchV2ResponseBodyModuleFlightItemListItemList {
	s.FlightRuleInfos = v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListItemList) SetItemId(v string) *FlightListingSearchV2ResponseBodyModuleFlightItemListItemList {
	s.ItemId = &v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListItemList) SetShoppingItemMap(v map[string]*ModuleFlightItemListItemListShoppingItemMapValue) *FlightListingSearchV2ResponseBodyModuleFlightItemListItemList {
	s.ShoppingItemMap = v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListItemList) SetSubItemPositionMap(v map[string][]*ModuleFlightItemListItemListSubItemPositionMapValue) *FlightListingSearchV2ResponseBodyModuleFlightItemListItemList {
	s.SubItemPositionMap = v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListItemList) SetSubItems(v []*FlightListingSearchV2ResponseBodyModuleFlightItemListItemListSubItems) *FlightListingSearchV2ResponseBodyModuleFlightItemListItemList {
	s.SubItems = v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListItemList) Validate() error {
	return dara.Validate(s)
}

type FlightListingSearchV2ResponseBodyModuleFlightItemListItemListSubItems struct {
	ShoppingItemMap map[string]*ModuleFlightItemListItemListSubItemsShoppingItemMapValue `json:"shopping_item_map,omitempty" xml:"shopping_item_map,omitempty"`
	Tag             *string                                                              `json:"tag,omitempty" xml:"tag,omitempty"`
	// example:
	//
	// 07df0bd9-f803-4a50-8449-f4bd675d9939
	UniqKey *string `json:"uniq_key,omitempty" xml:"uniq_key,omitempty"`
}

func (s FlightListingSearchV2ResponseBodyModuleFlightItemListItemListSubItems) String() string {
	return dara.Prettify(s)
}

func (s FlightListingSearchV2ResponseBodyModuleFlightItemListItemListSubItems) GoString() string {
	return s.String()
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListItemListSubItems) GetShoppingItemMap() map[string]*ModuleFlightItemListItemListSubItemsShoppingItemMapValue {
	return s.ShoppingItemMap
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListItemListSubItems) GetTag() *string {
	return s.Tag
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListItemListSubItems) GetUniqKey() *string {
	return s.UniqKey
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListItemListSubItems) SetShoppingItemMap(v map[string]*ModuleFlightItemListItemListSubItemsShoppingItemMapValue) *FlightListingSearchV2ResponseBodyModuleFlightItemListItemListSubItems {
	s.ShoppingItemMap = v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListItemListSubItems) SetTag(v string) *FlightListingSearchV2ResponseBodyModuleFlightItemListItemListSubItems {
	s.Tag = &v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListItemListSubItems) SetUniqKey(v string) *FlightListingSearchV2ResponseBodyModuleFlightItemListItemListSubItems {
	s.UniqKey = &v
	return s
}

func (s *FlightListingSearchV2ResponseBodyModuleFlightItemListItemListSubItems) Validate() error {
	return dara.Validate(s)
}
