// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightReShopListSearchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *IntlFlightReShopListSearchResponseBody
	GetCode() *string
	SetMessage(v string) *IntlFlightReShopListSearchResponseBody
	GetMessage() *string
	SetModule(v *IntlFlightReShopListSearchResponseBodyModule) *IntlFlightReShopListSearchResponseBody
	GetModule() *IntlFlightReShopListSearchResponseBodyModule
	SetRequestId(v string) *IntlFlightReShopListSearchResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *IntlFlightReShopListSearchResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *IntlFlightReShopListSearchResponseBody
	GetTraceId() *string
}

type IntlFlightReShopListSearchResponseBody struct {
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// 成功
	Message *string                                       `json:"message,omitempty" xml:"message,omitempty"`
	Module  *IntlFlightReShopListSearchResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
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

func (s IntlFlightReShopListSearchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopListSearchResponseBody) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopListSearchResponseBody) GetCode() *string {
	return s.Code
}

func (s *IntlFlightReShopListSearchResponseBody) GetMessage() *string {
	return s.Message
}

func (s *IntlFlightReShopListSearchResponseBody) GetModule() *IntlFlightReShopListSearchResponseBodyModule {
	return s.Module
}

func (s *IntlFlightReShopListSearchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *IntlFlightReShopListSearchResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *IntlFlightReShopListSearchResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *IntlFlightReShopListSearchResponseBody) SetCode(v string) *IntlFlightReShopListSearchResponseBody {
	s.Code = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBody) SetMessage(v string) *IntlFlightReShopListSearchResponseBody {
	s.Message = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBody) SetModule(v *IntlFlightReShopListSearchResponseBodyModule) *IntlFlightReShopListSearchResponseBody {
	s.Module = v
	return s
}

func (s *IntlFlightReShopListSearchResponseBody) SetRequestId(v string) *IntlFlightReShopListSearchResponseBody {
	s.RequestId = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBody) SetSuccess(v bool) *IntlFlightReShopListSearchResponseBody {
	s.Success = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBody) SetTraceId(v string) *IntlFlightReShopListSearchResponseBody {
	s.TraceId = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type IntlFlightReShopListSearchResponseBodyModule struct {
	// example:
	//
	// false
	NeedContinue *bool `json:"need_continue,omitempty" xml:"need_continue,omitempty"`
	// example:
	//
	// 2000
	NextReqWaitTime *int32                                                        `json:"next_req_wait_time,omitempty" xml:"next_req_wait_time,omitempty"`
	ReShopItemList  []*IntlFlightReShopListSearchResponseBodyModuleReShopItemList `json:"re_shop_item_list,omitempty" xml:"re_shop_item_list,omitempty" type:"Repeated"`
	// example:
	//
	// 284e692fffdf71e8a49aebbe0657a625
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
}

func (s IntlFlightReShopListSearchResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopListSearchResponseBodyModule) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopListSearchResponseBodyModule) GetNeedContinue() *bool {
	return s.NeedContinue
}

func (s *IntlFlightReShopListSearchResponseBodyModule) GetNextReqWaitTime() *int32 {
	return s.NextReqWaitTime
}

func (s *IntlFlightReShopListSearchResponseBodyModule) GetReShopItemList() []*IntlFlightReShopListSearchResponseBodyModuleReShopItemList {
	return s.ReShopItemList
}

func (s *IntlFlightReShopListSearchResponseBodyModule) GetToken() *string {
	return s.Token
}

func (s *IntlFlightReShopListSearchResponseBodyModule) SetNeedContinue(v bool) *IntlFlightReShopListSearchResponseBodyModule {
	s.NeedContinue = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModule) SetNextReqWaitTime(v int32) *IntlFlightReShopListSearchResponseBodyModule {
	s.NextReqWaitTime = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModule) SetReShopItemList(v []*IntlFlightReShopListSearchResponseBodyModuleReShopItemList) *IntlFlightReShopListSearchResponseBodyModule {
	s.ReShopItemList = v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModule) SetToken(v string) *IntlFlightReShopListSearchResponseBodyModule {
	s.Token = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModule) Validate() error {
	if s.ReShopItemList != nil {
		for _, item := range s.ReShopItemList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type IntlFlightReShopListSearchResponseBodyModuleReShopItemList struct {
	BestPriceItem      *IntlFlightReShopListSearchResponseBodyModuleReShopItemListBestPriceItem        `json:"best_price_item,omitempty" xml:"best_price_item,omitempty" type:"Struct"`
	FlightJourneyInfos []*IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfos `json:"flight_journey_infos,omitempty" xml:"flight_journey_infos,omitempty" type:"Repeated"`
}

func (s IntlFlightReShopListSearchResponseBodyModuleReShopItemList) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopListSearchResponseBodyModuleReShopItemList) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemList) GetBestPriceItem() *IntlFlightReShopListSearchResponseBodyModuleReShopItemListBestPriceItem {
	return s.BestPriceItem
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemList) GetFlightJourneyInfos() []*IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfos {
	return s.FlightJourneyInfos
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemList) SetBestPriceItem(v *IntlFlightReShopListSearchResponseBodyModuleReShopItemListBestPriceItem) *IntlFlightReShopListSearchResponseBodyModuleReShopItemList {
	s.BestPriceItem = v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemList) SetFlightJourneyInfos(v []*IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfos) *IntlFlightReShopListSearchResponseBodyModuleReShopItemList {
	s.FlightJourneyInfos = v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemList) Validate() error {
	if s.BestPriceItem != nil {
		if err := s.BestPriceItem.Validate(); err != nil {
			return err
		}
	}
	if s.FlightJourneyInfos != nil {
		for _, item := range s.FlightJourneyInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type IntlFlightReShopListSearchResponseBodyModuleReShopItemListBestPriceItem struct {
	// example:
	//
	// b83e3d6ebb8b44dfb94c763abc66c966_2
	ItemId *string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// example:
	//
	// normal
	ItemType        *string                                                                            `json:"item_type,omitempty" xml:"item_type,omitempty"`
	ShoppingItemMap map[string]*ModuleReShopItemListBestPriceItemShoppingItemMapValue                  `json:"shopping_item_map,omitempty" xml:"shopping_item_map,omitempty"`
	SubItems        []*IntlFlightReShopListSearchResponseBodyModuleReShopItemListBestPriceItemSubItems `json:"sub_items,omitempty" xml:"sub_items,omitempty" type:"Repeated"`
}

func (s IntlFlightReShopListSearchResponseBodyModuleReShopItemListBestPriceItem) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopListSearchResponseBodyModuleReShopItemListBestPriceItem) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListBestPriceItem) GetItemId() *string {
	return s.ItemId
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListBestPriceItem) GetItemType() *string {
	return s.ItemType
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListBestPriceItem) GetShoppingItemMap() map[string]*ModuleReShopItemListBestPriceItemShoppingItemMapValue {
	return s.ShoppingItemMap
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListBestPriceItem) GetSubItems() []*IntlFlightReShopListSearchResponseBodyModuleReShopItemListBestPriceItemSubItems {
	return s.SubItems
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListBestPriceItem) SetItemId(v string) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListBestPriceItem {
	s.ItemId = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListBestPriceItem) SetItemType(v string) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListBestPriceItem {
	s.ItemType = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListBestPriceItem) SetShoppingItemMap(v map[string]*ModuleReShopItemListBestPriceItemShoppingItemMapValue) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListBestPriceItem {
	s.ShoppingItemMap = v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListBestPriceItem) SetSubItems(v []*IntlFlightReShopListSearchResponseBodyModuleReShopItemListBestPriceItemSubItems) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListBestPriceItem {
	s.SubItems = v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListBestPriceItem) Validate() error {
	if s.SubItems != nil {
		for _, item := range s.SubItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type IntlFlightReShopListSearchResponseBodyModuleReShopItemListBestPriceItemSubItems struct {
	// example:
	//
	// 5.1
	DiscountNum         *float64                                                                                              `json:"discount_num,omitempty" xml:"discount_num,omitempty"`
	SegmentKeys         []*string                                                                                             `json:"segment_keys,omitempty" xml:"segment_keys,omitempty" type:"Repeated"`
	SegmentPositionList []*IntlFlightReShopListSearchResponseBodyModuleReShopItemListBestPriceItemSubItemsSegmentPositionList `json:"segment_position_list,omitempty" xml:"segment_position_list,omitempty" type:"Repeated"`
	ShoppingItemMap     map[string]*ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValue                             `json:"shopping_item_map,omitempty" xml:"shopping_item_map,omitempty"`
	// example:
	//
	// 07df0bd9-f803-4a50-8449-f4bd675d9939
	UniqKey *string `json:"uniq_key,omitempty" xml:"uniq_key,omitempty"`
}

func (s IntlFlightReShopListSearchResponseBodyModuleReShopItemListBestPriceItemSubItems) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopListSearchResponseBodyModuleReShopItemListBestPriceItemSubItems) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListBestPriceItemSubItems) GetDiscountNum() *float64 {
	return s.DiscountNum
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListBestPriceItemSubItems) GetSegmentKeys() []*string {
	return s.SegmentKeys
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListBestPriceItemSubItems) GetSegmentPositionList() []*IntlFlightReShopListSearchResponseBodyModuleReShopItemListBestPriceItemSubItemsSegmentPositionList {
	return s.SegmentPositionList
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListBestPriceItemSubItems) GetShoppingItemMap() map[string]*ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValue {
	return s.ShoppingItemMap
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListBestPriceItemSubItems) GetUniqKey() *string {
	return s.UniqKey
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListBestPriceItemSubItems) SetDiscountNum(v float64) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListBestPriceItemSubItems {
	s.DiscountNum = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListBestPriceItemSubItems) SetSegmentKeys(v []*string) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListBestPriceItemSubItems {
	s.SegmentKeys = v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListBestPriceItemSubItems) SetSegmentPositionList(v []*IntlFlightReShopListSearchResponseBodyModuleReShopItemListBestPriceItemSubItemsSegmentPositionList) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListBestPriceItemSubItems {
	s.SegmentPositionList = v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListBestPriceItemSubItems) SetShoppingItemMap(v map[string]*ModuleReShopItemListBestPriceItemSubItemsShoppingItemMapValue) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListBestPriceItemSubItems {
	s.ShoppingItemMap = v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListBestPriceItemSubItems) SetUniqKey(v string) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListBestPriceItemSubItems {
	s.UniqKey = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListBestPriceItemSubItems) Validate() error {
	if s.SegmentPositionList != nil {
		for _, item := range s.SegmentPositionList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type IntlFlightReShopListSearchResponseBodyModuleReShopItemListBestPriceItemSubItemsSegmentPositionList struct {
	// example:
	//
	// 0
	JourneyIndex *int32 `json:"journey_index,omitempty" xml:"journey_index,omitempty"`
	// example:
	//
	// 0
	SegmentIndex *int32 `json:"segment_index,omitempty" xml:"segment_index,omitempty"`
}

func (s IntlFlightReShopListSearchResponseBodyModuleReShopItemListBestPriceItemSubItemsSegmentPositionList) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopListSearchResponseBodyModuleReShopItemListBestPriceItemSubItemsSegmentPositionList) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListBestPriceItemSubItemsSegmentPositionList) GetJourneyIndex() *int32 {
	return s.JourneyIndex
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListBestPriceItemSubItemsSegmentPositionList) GetSegmentIndex() *int32 {
	return s.SegmentIndex
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListBestPriceItemSubItemsSegmentPositionList) SetJourneyIndex(v int32) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListBestPriceItemSubItemsSegmentPositionList {
	s.JourneyIndex = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListBestPriceItemSubItemsSegmentPositionList) SetSegmentIndex(v int32) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListBestPriceItemSubItemsSegmentPositionList {
	s.SegmentIndex = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListBestPriceItemSubItemsSegmentPositionList) Validate() error {
	return dara.Validate(s)
}

type IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfos struct {
	// example:
	//
	// HGH
	ArrCityCode *string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	// example:
	//
	// 杭州
	ArrCityName *string `json:"arr_city_name,omitempty" xml:"arr_city_name,omitempty"`
	// example:
	//
	// 2023-08-13 09:25
	ArrTime *string `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	// example:
	//
	// BJS
	DepCityCode *string `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	// example:
	//
	// 北京
	DepCityName *string `json:"dep_city_name,omitempty" xml:"dep_city_name,omitempty"`
	// example:
	//
	// 2023-08-13 07:25
	DepTime *string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	// example:
	//
	// 120
	Duration           *int32                                                                                            `json:"duration,omitempty" xml:"duration,omitempty"`
	FlightSegmentInfos []*IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos `json:"flight_segment_infos,omitempty" xml:"flight_segment_infos,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	JourneyIndex *int32 `json:"journey_index,omitempty" xml:"journey_index,omitempty"`
	// example:
	//
	// 0
	TransferTime *int32 `json:"transfer_time,omitempty" xml:"transfer_time,omitempty"`
}

func (s IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfos) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfos) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfos) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfos) GetArrCityName() *string {
	return s.ArrCityName
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfos) GetArrTime() *string {
	return s.ArrTime
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfos) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfos) GetDepCityName() *string {
	return s.DepCityName
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfos) GetDepTime() *string {
	return s.DepTime
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfos) GetDuration() *int32 {
	return s.Duration
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfos) GetFlightSegmentInfos() []*IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos {
	return s.FlightSegmentInfos
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfos) GetJourneyIndex() *int32 {
	return s.JourneyIndex
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfos) GetTransferTime() *int32 {
	return s.TransferTime
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfos) SetArrCityCode(v string) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfos {
	s.ArrCityCode = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfos) SetArrCityName(v string) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfos {
	s.ArrCityName = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfos) SetArrTime(v string) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfos {
	s.ArrTime = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfos) SetDepCityCode(v string) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfos {
	s.DepCityCode = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfos) SetDepCityName(v string) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfos {
	s.DepCityName = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfos) SetDepTime(v string) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfos {
	s.DepTime = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfos) SetDuration(v int32) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfos {
	s.Duration = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfos) SetFlightSegmentInfos(v []*IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfos {
	s.FlightSegmentInfos = v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfos) SetJourneyIndex(v int32) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfos {
	s.JourneyIndex = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfos) SetTransferTime(v int32) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfos {
	s.TransferTime = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfos) Validate() error {
	if s.FlightSegmentInfos != nil {
		for _, item := range s.FlightSegmentInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos struct {
	AirlineInfo    *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosAirlineInfo    `json:"airline_info,omitempty" xml:"airline_info,omitempty" type:"Struct"`
	ArrAirportInfo *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosArrAirportInfo `json:"arr_airport_info,omitempty" xml:"arr_airport_info,omitempty" type:"Struct"`
	// example:
	//
	// HGH
	ArrCityCode *string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	// example:
	//
	// 杭州
	ArrCityName *string `json:"arr_city_name,omitempty" xml:"arr_city_name,omitempty"`
	// example:
	//
	// 2023-08-13 09:25
	ArrTime *string `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	// example:
	//
	// 2022-06-06T12:56:34Z
	ArrTimeUTC     *string                                                                                                       `json:"arr_time_u_t_c,omitempty" xml:"arr_time_u_t_c,omitempty"`
	DepAirportInfo *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosDepAirportInfo `json:"dep_airport_info,omitempty" xml:"dep_airport_info,omitempty" type:"Struct"`
	// example:
	//
	// BJS
	DepCityCode *string `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	// example:
	//
	// 北京
	DepCityName *string `json:"dep_city_name,omitempty" xml:"dep_city_name,omitempty"`
	// example:
	//
	// 2023-08-13 07:25
	DepTime *string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	// example:
	//
	// 2022-06-06T12:56:34Z
	DepTimeUTC *string `json:"dep_time_u_t_c,omitempty" xml:"dep_time_u_t_c,omitempty"`
	// example:
	//
	// 320
	Duration *int32 `json:"duration,omitempty" xml:"duration,omitempty"`
	// example:
	//
	// CA2013
	FlightNo        *string                                                                                                        `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
	FlightShareInfo *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfo `json:"flight_share_info,omitempty" xml:"flight_share_info,omitempty" type:"Struct"`
	// example:
	//
	// 中型机
	FlightSize         *string                                                                                                             `json:"flight_size,omitempty" xml:"flight_size,omitempty"`
	FlightStopInfoList []*IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoList `json:"flight_stop_info_list,omitempty" xml:"flight_stop_info_list,omitempty" type:"Repeated"`
	// example:
	//
	// 320
	FlightType *string `json:"flight_type,omitempty" xml:"flight_type,omitempty"`
	// example:
	//
	// 0
	JourneyIndex      *int32                                                                                                           `json:"journey_index,omitempty" xml:"journey_index,omitempty"`
	LuggageDirectInfo *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosLuggageDirectInfo `json:"luggage_direct_info,omitempty" xml:"luggage_direct_info,omitempty" type:"Struct"`
	// example:
	//
	// 空客
	Manufacturer *string `json:"manufacturer,omitempty" xml:"manufacturer,omitempty"`
	// example:
	//
	// 1
	Meal *int32 `json:"meal,omitempty" xml:"meal,omitempty"`
	// example:
	//
	// 小食
	MealDesc *string `json:"meal_desc,omitempty" xml:"meal_desc,omitempty"`
	// example:
	//
	// 12
	Miles *int32 `json:"miles,omitempty" xml:"miles,omitempty"`
	// example:
	//
	// 90%
	OnTimeRate *string `json:"on_time_rate,omitempty" xml:"on_time_rate,omitempty"`
	// example:
	//
	// 0
	OneMore *int32 `json:"one_more,omitempty" xml:"one_more,omitempty"`
	// example:
	//
	// +1天
	OneMoreShow *string                                                                                                  `json:"one_more_show,omitempty" xml:"one_more_show,omitempty"`
	OtherInfo   *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosOtherInfo `json:"other_info,omitempty" xml:"other_info,omitempty" type:"Struct"`
	// example:
	//
	// 0
	SegmentIndex *int32 `json:"segment_index,omitempty" xml:"segment_index,omitempty"`
	// example:
	//
	// KN6728HGHPKX0725
	SegmentKey        *string                                                                                                          `json:"segment_key,omitempty" xml:"segment_key,omitempty"`
	SegmentVisaRemark *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark `json:"segment_visa_remark,omitempty" xml:"segment_visa_remark,omitempty" type:"Struct"`
	// example:
	//
	// true
	Share *bool `json:"share,omitempty" xml:"share,omitempty"`
	// example:
	//
	// 中
	ShortFlightSize *string `json:"short_flight_size,omitempty" xml:"short_flight_size,omitempty"`
	// example:
	//
	// true
	Stop                 *bool                                                                                                               `json:"stop,omitempty" xml:"stop,omitempty"`
	TicketingAirlineInfo *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosTicketingAirlineInfo `json:"ticketing_airline_info,omitempty" xml:"ticketing_airline_info,omitempty" type:"Struct"`
	// example:
	//
	// 5小时20分
	TotalTime *string `json:"total_time,omitempty" xml:"total_time,omitempty"`
}

func (s IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos) GetAirlineInfo() *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosAirlineInfo {
	return s.AirlineInfo
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos) GetArrAirportInfo() *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosArrAirportInfo {
	return s.ArrAirportInfo
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos) GetArrCityName() *string {
	return s.ArrCityName
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos) GetArrTime() *string {
	return s.ArrTime
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos) GetArrTimeUTC() *string {
	return s.ArrTimeUTC
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos) GetDepAirportInfo() *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosDepAirportInfo {
	return s.DepAirportInfo
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos) GetDepCityName() *string {
	return s.DepCityName
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos) GetDepTime() *string {
	return s.DepTime
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos) GetDepTimeUTC() *string {
	return s.DepTimeUTC
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos) GetDuration() *int32 {
	return s.Duration
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos) GetFlightNo() *string {
	return s.FlightNo
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos) GetFlightShareInfo() *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfo {
	return s.FlightShareInfo
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos) GetFlightSize() *string {
	return s.FlightSize
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos) GetFlightStopInfoList() []*IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoList {
	return s.FlightStopInfoList
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos) GetFlightType() *string {
	return s.FlightType
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos) GetJourneyIndex() *int32 {
	return s.JourneyIndex
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos) GetLuggageDirectInfo() *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosLuggageDirectInfo {
	return s.LuggageDirectInfo
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos) GetManufacturer() *string {
	return s.Manufacturer
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos) GetMeal() *int32 {
	return s.Meal
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos) GetMealDesc() *string {
	return s.MealDesc
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos) GetMiles() *int32 {
	return s.Miles
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos) GetOnTimeRate() *string {
	return s.OnTimeRate
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos) GetOneMore() *int32 {
	return s.OneMore
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos) GetOneMoreShow() *string {
	return s.OneMoreShow
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos) GetOtherInfo() *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosOtherInfo {
	return s.OtherInfo
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos) GetSegmentIndex() *int32 {
	return s.SegmentIndex
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos) GetSegmentKey() *string {
	return s.SegmentKey
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos) GetSegmentVisaRemark() *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark {
	return s.SegmentVisaRemark
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos) GetShare() *bool {
	return s.Share
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos) GetShortFlightSize() *string {
	return s.ShortFlightSize
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos) GetStop() *bool {
	return s.Stop
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos) GetTicketingAirlineInfo() *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosTicketingAirlineInfo {
	return s.TicketingAirlineInfo
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos) GetTotalTime() *string {
	return s.TotalTime
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos) SetAirlineInfo(v *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosAirlineInfo) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos {
	s.AirlineInfo = v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos) SetArrAirportInfo(v *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosArrAirportInfo) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos {
	s.ArrAirportInfo = v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos) SetArrCityCode(v string) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos {
	s.ArrCityCode = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos) SetArrCityName(v string) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos {
	s.ArrCityName = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos) SetArrTime(v string) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos {
	s.ArrTime = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos) SetArrTimeUTC(v string) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos {
	s.ArrTimeUTC = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos) SetDepAirportInfo(v *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosDepAirportInfo) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos {
	s.DepAirportInfo = v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos) SetDepCityCode(v string) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos {
	s.DepCityCode = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos) SetDepCityName(v string) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos {
	s.DepCityName = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos) SetDepTime(v string) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos {
	s.DepTime = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos) SetDepTimeUTC(v string) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos {
	s.DepTimeUTC = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos) SetDuration(v int32) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos {
	s.Duration = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos) SetFlightNo(v string) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos {
	s.FlightNo = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos) SetFlightShareInfo(v *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfo) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos {
	s.FlightShareInfo = v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos) SetFlightSize(v string) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos {
	s.FlightSize = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos) SetFlightStopInfoList(v []*IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos {
	s.FlightStopInfoList = v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos) SetFlightType(v string) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos {
	s.FlightType = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos) SetJourneyIndex(v int32) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos {
	s.JourneyIndex = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos) SetLuggageDirectInfo(v *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosLuggageDirectInfo) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos {
	s.LuggageDirectInfo = v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos) SetManufacturer(v string) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos {
	s.Manufacturer = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos) SetMeal(v int32) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos {
	s.Meal = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos) SetMealDesc(v string) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos {
	s.MealDesc = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos) SetMiles(v int32) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos {
	s.Miles = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos) SetOnTimeRate(v string) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos {
	s.OnTimeRate = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos) SetOneMore(v int32) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos {
	s.OneMore = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos) SetOneMoreShow(v string) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos {
	s.OneMoreShow = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos) SetOtherInfo(v *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosOtherInfo) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos {
	s.OtherInfo = v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos) SetSegmentIndex(v int32) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos {
	s.SegmentIndex = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos) SetSegmentKey(v string) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos {
	s.SegmentKey = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos) SetSegmentVisaRemark(v *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos {
	s.SegmentVisaRemark = v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos) SetShare(v bool) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos {
	s.Share = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos) SetShortFlightSize(v string) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos {
	s.ShortFlightSize = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos) SetStop(v bool) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos {
	s.Stop = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos) SetTicketingAirlineInfo(v *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosTicketingAirlineInfo) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos {
	s.TicketingAirlineInfo = v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos) SetTotalTime(v string) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos {
	s.TotalTime = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfos) Validate() error {
	if s.AirlineInfo != nil {
		if err := s.AirlineInfo.Validate(); err != nil {
			return err
		}
	}
	if s.ArrAirportInfo != nil {
		if err := s.ArrAirportInfo.Validate(); err != nil {
			return err
		}
	}
	if s.DepAirportInfo != nil {
		if err := s.DepAirportInfo.Validate(); err != nil {
			return err
		}
	}
	if s.FlightShareInfo != nil {
		if err := s.FlightShareInfo.Validate(); err != nil {
			return err
		}
	}
	if s.FlightStopInfoList != nil {
		for _, item := range s.FlightStopInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.LuggageDirectInfo != nil {
		if err := s.LuggageDirectInfo.Validate(); err != nil {
			return err
		}
	}
	if s.OtherInfo != nil {
		if err := s.OtherInfo.Validate(); err != nil {
			return err
		}
	}
	if s.SegmentVisaRemark != nil {
		if err := s.SegmentVisaRemark.Validate(); err != nil {
			return err
		}
	}
	if s.TicketingAirlineInfo != nil {
		if err := s.TicketingAirlineInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosAirlineInfo struct {
	// example:
	//
	// CA
	AirlineCode *string `json:"airline_code,omitempty" xml:"airline_code,omitempty"`
	// example:
	//
	// 中国国航
	AirlineName *string `json:"airline_name,omitempty" xml:"airline_name,omitempty"`
	// example:
	//
	// false
	CheapAirline *bool `json:"cheap_airline,omitempty" xml:"cheap_airline,omitempty"`
	// example:
	//
	// https://
	IconUrl *string `json:"icon_url,omitempty" xml:"icon_url,omitempty"`
	// example:
	//
	// 国航
	ShortName *string `json:"short_name,omitempty" xml:"short_name,omitempty"`
}

func (s IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosAirlineInfo) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosAirlineInfo) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosAirlineInfo) GetAirlineCode() *string {
	return s.AirlineCode
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosAirlineInfo) GetAirlineName() *string {
	return s.AirlineName
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosAirlineInfo) GetCheapAirline() *bool {
	return s.CheapAirline
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosAirlineInfo) GetIconUrl() *string {
	return s.IconUrl
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosAirlineInfo) GetShortName() *string {
	return s.ShortName
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosAirlineInfo) SetAirlineCode(v string) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosAirlineInfo {
	s.AirlineCode = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosAirlineInfo) SetAirlineName(v string) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosAirlineInfo {
	s.AirlineName = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosAirlineInfo) SetCheapAirline(v bool) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosAirlineInfo {
	s.CheapAirline = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosAirlineInfo) SetIconUrl(v string) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosAirlineInfo {
	s.IconUrl = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosAirlineInfo) SetShortName(v string) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosAirlineInfo {
	s.ShortName = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosAirlineInfo) Validate() error {
	return dara.Validate(s)
}

type IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosArrAirportInfo struct {
	// example:
	//
	// HGH
	AirportCode *string `json:"airport_code,omitempty" xml:"airport_code,omitempty"`
	// example:
	//
	// 萧山国际机场
	AirportName *string `json:"airport_name,omitempty" xml:"airport_name,omitempty"`
	// example:
	//
	// 萧山
	AirportShortName *string `json:"airport_short_name,omitempty" xml:"airport_short_name,omitempty"`
	// example:
	//
	// T3
	Terminal *string `json:"terminal,omitempty" xml:"terminal,omitempty"`
}

func (s IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosArrAirportInfo) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosArrAirportInfo) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosArrAirportInfo) GetAirportCode() *string {
	return s.AirportCode
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosArrAirportInfo) GetAirportName() *string {
	return s.AirportName
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosArrAirportInfo) GetAirportShortName() *string {
	return s.AirportShortName
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosArrAirportInfo) GetTerminal() *string {
	return s.Terminal
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosArrAirportInfo) SetAirportCode(v string) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosArrAirportInfo {
	s.AirportCode = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosArrAirportInfo) SetAirportName(v string) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosArrAirportInfo {
	s.AirportName = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosArrAirportInfo) SetAirportShortName(v string) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosArrAirportInfo {
	s.AirportShortName = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosArrAirportInfo) SetTerminal(v string) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosArrAirportInfo {
	s.Terminal = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosArrAirportInfo) Validate() error {
	return dara.Validate(s)
}

type IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosDepAirportInfo struct {
	// example:
	//
	// PKX
	AirportCode *string `json:"airport_code,omitempty" xml:"airport_code,omitempty"`
	// example:
	//
	// 大兴国际机场
	AirportName *string `json:"airport_name,omitempty" xml:"airport_name,omitempty"`
	// example:
	//
	// 大兴
	AirportShortName *string `json:"airport_short_name,omitempty" xml:"airport_short_name,omitempty"`
	// example:
	//
	// T3
	Terminal *string `json:"terminal,omitempty" xml:"terminal,omitempty"`
}

func (s IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosDepAirportInfo) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosDepAirportInfo) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosDepAirportInfo) GetAirportCode() *string {
	return s.AirportCode
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosDepAirportInfo) GetAirportName() *string {
	return s.AirportName
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosDepAirportInfo) GetAirportShortName() *string {
	return s.AirportShortName
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosDepAirportInfo) GetTerminal() *string {
	return s.Terminal
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosDepAirportInfo) SetAirportCode(v string) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosDepAirportInfo {
	s.AirportCode = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosDepAirportInfo) SetAirportName(v string) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosDepAirportInfo {
	s.AirportName = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosDepAirportInfo) SetAirportShortName(v string) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosDepAirportInfo {
	s.AirportShortName = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosDepAirportInfo) SetTerminal(v string) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosDepAirportInfo {
	s.Terminal = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosDepAirportInfo) Validate() error {
	return dara.Validate(s)
}

type IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfo struct {
	OperatingAirlineInfo *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo `json:"operating_airline_info,omitempty" xml:"operating_airline_info,omitempty" type:"Struct"`
	// example:
	//
	// CX601
	OperatingFlightNo *string `json:"operating_flight_no,omitempty" xml:"operating_flight_no,omitempty"`
}

func (s IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfo) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfo) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfo) GetOperatingAirlineInfo() *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo {
	return s.OperatingAirlineInfo
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfo) GetOperatingFlightNo() *string {
	return s.OperatingFlightNo
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfo) SetOperatingAirlineInfo(v *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfo {
	s.OperatingAirlineInfo = v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfo) SetOperatingFlightNo(v string) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfo {
	s.OperatingFlightNo = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfo) Validate() error {
	if s.OperatingAirlineInfo != nil {
		if err := s.OperatingAirlineInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo struct {
	// example:
	//
	// CA
	AirlineCode *string `json:"airline_code,omitempty" xml:"airline_code,omitempty"`
	// example:
	//
	// 中国国航
	AirlineName *string `json:"airline_name,omitempty" xml:"airline_name,omitempty"`
	// example:
	//
	// false
	CheapAirline *bool `json:"cheap_airline,omitempty" xml:"cheap_airline,omitempty"`
	// example:
	//
	// https://
	IconUrl *string `json:"icon_url,omitempty" xml:"icon_url,omitempty"`
	// example:
	//
	// 国航
	ShortName *string `json:"short_name,omitempty" xml:"short_name,omitempty"`
}

func (s IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) GetAirlineCode() *string {
	return s.AirlineCode
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) GetAirlineName() *string {
	return s.AirlineName
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) GetCheapAirline() *bool {
	return s.CheapAirline
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) GetIconUrl() *string {
	return s.IconUrl
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) GetShortName() *string {
	return s.ShortName
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) SetAirlineCode(v string) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo {
	s.AirlineCode = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) SetAirlineName(v string) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo {
	s.AirlineName = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) SetCheapAirline(v bool) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo {
	s.CheapAirline = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) SetIconUrl(v string) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo {
	s.IconUrl = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) SetShortName(v string) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo {
	s.ShortName = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) Validate() error {
	return dara.Validate(s)
}

type IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoList struct {
	// example:
	//
	// HGH
	StopAirport           *string                                                                                                                                `json:"stop_airport,omitempty" xml:"stop_airport,omitempty"`
	StopAirportCountyInfo *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoListStopAirportCountyInfo `json:"stop_airport_county_info,omitempty" xml:"stop_airport_county_info,omitempty" type:"Struct"`
	// example:
	//
	// 萧山国际机场
	StopAirportName *string `json:"stop_airport_name,omitempty" xml:"stop_airport_name,omitempty"`
	// example:
	//
	// T1
	StopArrTerm *string `json:"stop_arr_term,omitempty" xml:"stop_arr_term,omitempty"`
	// example:
	//
	// 2023-08-13 07:25
	StopArrTime *string `json:"stop_arr_time,omitempty" xml:"stop_arr_time,omitempty"`
	// example:
	//
	// BJS
	StopCityCode *string `json:"stop_city_code,omitempty" xml:"stop_city_code,omitempty"`
	// example:
	//
	// 北京
	StopCityName *string `json:"stop_city_name,omitempty" xml:"stop_city_name,omitempty"`
	// example:
	//
	// T1
	StopDepTerm *string `json:"stop_dep_term,omitempty" xml:"stop_dep_term,omitempty"`
	// example:
	//
	// 2023-08-13 09:25
	StopDepTime *string `json:"stop_dep_time,omitempty" xml:"stop_dep_time,omitempty"`
	// example:
	//
	// 60
	StopTime *string `json:"stop_time,omitempty" xml:"stop_time,omitempty"`
}

func (s IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) GetStopAirport() *string {
	return s.StopAirport
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) GetStopAirportCountyInfo() *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoListStopAirportCountyInfo {
	return s.StopAirportCountyInfo
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) GetStopAirportName() *string {
	return s.StopAirportName
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) GetStopArrTerm() *string {
	return s.StopArrTerm
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) GetStopArrTime() *string {
	return s.StopArrTime
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) GetStopCityCode() *string {
	return s.StopCityCode
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) GetStopCityName() *string {
	return s.StopCityName
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) GetStopDepTerm() *string {
	return s.StopDepTerm
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) GetStopDepTime() *string {
	return s.StopDepTime
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) GetStopTime() *string {
	return s.StopTime
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) SetStopAirport(v string) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoList {
	s.StopAirport = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) SetStopAirportCountyInfo(v *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoListStopAirportCountyInfo) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoList {
	s.StopAirportCountyInfo = v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) SetStopAirportName(v string) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoList {
	s.StopAirportName = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) SetStopArrTerm(v string) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoList {
	s.StopArrTerm = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) SetStopArrTime(v string) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoList {
	s.StopArrTime = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) SetStopCityCode(v string) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoList {
	s.StopCityCode = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) SetStopCityName(v string) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoList {
	s.StopCityName = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) SetStopDepTerm(v string) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoList {
	s.StopDepTerm = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) SetStopDepTime(v string) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoList {
	s.StopDepTime = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) SetStopTime(v string) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoList {
	s.StopTime = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) Validate() error {
	if s.StopAirportCountyInfo != nil {
		if err := s.StopAirportCountyInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoListStopAirportCountyInfo struct {
	// example:
	//
	// 330182
	Adcode *string `json:"adcode,omitempty" xml:"adcode,omitempty"`
	// example:
	//
	// HGH
	AirportCityCode *string `json:"airport_city_code,omitempty" xml:"airport_city_code,omitempty"`
	// example:
	//
	// 杭州
	AirportCityName *string `json:"airport_city_name,omitempty" xml:"airport_city_name,omitempty"`
	// example:
	//
	// HGH
	AirportCode *string `json:"airport_code,omitempty" xml:"airport_code,omitempty"`
	// example:
	//
	// 萧山国际机场
	AirportName *string `json:"airport_name,omitempty" xml:"airport_name,omitempty"`
	// example:
	//
	// 杭州
	AirportParentCityName *string `json:"airport_parent_city_name,omitempty" xml:"airport_parent_city_name,omitempty"`
	// example:
	//
	// 杭州
	CountyCityAdcode *string `json:"county_city_adcode,omitempty" xml:"county_city_adcode,omitempty"`
	// example:
	//
	// 330182
	CountyCityName *string `json:"county_city_name,omitempty" xml:"county_city_name,omitempty"`
	// example:
	//
	// 330182
	PrefectureCityAdcode *string `json:"prefecture_city_adcode,omitempty" xml:"prefecture_city_adcode,omitempty"`
	// example:
	//
	// 杭州
	PrefectureCityName *string `json:"prefecture_city_name,omitempty" xml:"prefecture_city_name,omitempty"`
}

func (s IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoListStopAirportCountyInfo) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoListStopAirportCountyInfo) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoListStopAirportCountyInfo) GetAdcode() *string {
	return s.Adcode
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoListStopAirportCountyInfo) GetAirportCityCode() *string {
	return s.AirportCityCode
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoListStopAirportCountyInfo) GetAirportCityName() *string {
	return s.AirportCityName
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoListStopAirportCountyInfo) GetAirportCode() *string {
	return s.AirportCode
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoListStopAirportCountyInfo) GetAirportName() *string {
	return s.AirportName
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoListStopAirportCountyInfo) GetAirportParentCityName() *string {
	return s.AirportParentCityName
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoListStopAirportCountyInfo) GetCountyCityAdcode() *string {
	return s.CountyCityAdcode
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoListStopAirportCountyInfo) GetCountyCityName() *string {
	return s.CountyCityName
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoListStopAirportCountyInfo) GetPrefectureCityAdcode() *string {
	return s.PrefectureCityAdcode
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoListStopAirportCountyInfo) GetPrefectureCityName() *string {
	return s.PrefectureCityName
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoListStopAirportCountyInfo) SetAdcode(v string) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoListStopAirportCountyInfo {
	s.Adcode = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoListStopAirportCountyInfo) SetAirportCityCode(v string) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoListStopAirportCountyInfo {
	s.AirportCityCode = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoListStopAirportCountyInfo) SetAirportCityName(v string) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoListStopAirportCountyInfo {
	s.AirportCityName = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoListStopAirportCountyInfo) SetAirportCode(v string) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoListStopAirportCountyInfo {
	s.AirportCode = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoListStopAirportCountyInfo) SetAirportName(v string) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoListStopAirportCountyInfo {
	s.AirportName = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoListStopAirportCountyInfo) SetAirportParentCityName(v string) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoListStopAirportCountyInfo {
	s.AirportParentCityName = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoListStopAirportCountyInfo) SetCountyCityAdcode(v string) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoListStopAirportCountyInfo {
	s.CountyCityAdcode = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoListStopAirportCountyInfo) SetCountyCityName(v string) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoListStopAirportCountyInfo {
	s.CountyCityName = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoListStopAirportCountyInfo) SetPrefectureCityAdcode(v string) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoListStopAirportCountyInfo {
	s.PrefectureCityAdcode = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoListStopAirportCountyInfo) SetPrefectureCityName(v string) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoListStopAirportCountyInfo {
	s.PrefectureCityName = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosFlightStopInfoListStopAirportCountyInfo) Validate() error {
	return dara.Validate(s)
}

type IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosLuggageDirectInfo struct {
	// example:
	//
	// 1
	DepCityLuggageDirect *int32 `json:"dep_city_luggage_direct,omitempty" xml:"dep_city_luggage_direct,omitempty"`
	// example:
	//
	// 1
	StopCityLuggageDirect *int32 `json:"stop_city_luggage_direct,omitempty" xml:"stop_city_luggage_direct,omitempty"`
}

func (s IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosLuggageDirectInfo) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosLuggageDirectInfo) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosLuggageDirectInfo) GetDepCityLuggageDirect() *int32 {
	return s.DepCityLuggageDirect
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosLuggageDirectInfo) GetStopCityLuggageDirect() *int32 {
	return s.StopCityLuggageDirect
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosLuggageDirectInfo) SetDepCityLuggageDirect(v int32) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosLuggageDirectInfo {
	s.DepCityLuggageDirect = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosLuggageDirectInfo) SetStopCityLuggageDirect(v int32) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosLuggageDirectInfo {
	s.StopCityLuggageDirect = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosLuggageDirectInfo) Validate() error {
	return dara.Validate(s)
}

type IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosOtherInfo struct {
	// example:
	//
	// 2.5年
	AircraftAge *string `json:"aircraft_age,omitempty" xml:"aircraft_age,omitempty"`
	// example:
	//
	// 平均延误58分钟
	AvgDelayTime *string `json:"avg_delay_time,omitempty" xml:"avg_delay_time,omitempty"`
	// example:
	//
	// 10%
	FlightCancelRate *string `json:"flight_cancel_rate,omitempty" xml:"flight_cancel_rate,omitempty"`
	// example:
	//
	// 10%
	JetBridgeRate *string `json:"jet_bridge_rate,omitempty" xml:"jet_bridge_rate,omitempty"`
	// example:
	//
	// 到达准点率90%
	OnTimeRate *string `json:"on_time_rate,omitempty" xml:"on_time_rate,omitempty"`
	// example:
	//
	// true
	Wifi *bool `json:"wifi,omitempty" xml:"wifi,omitempty"`
}

func (s IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosOtherInfo) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosOtherInfo) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosOtherInfo) GetAircraftAge() *string {
	return s.AircraftAge
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosOtherInfo) GetAvgDelayTime() *string {
	return s.AvgDelayTime
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosOtherInfo) GetFlightCancelRate() *string {
	return s.FlightCancelRate
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosOtherInfo) GetJetBridgeRate() *string {
	return s.JetBridgeRate
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosOtherInfo) GetOnTimeRate() *string {
	return s.OnTimeRate
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosOtherInfo) GetWifi() *bool {
	return s.Wifi
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosOtherInfo) SetAircraftAge(v string) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosOtherInfo {
	s.AircraftAge = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosOtherInfo) SetAvgDelayTime(v string) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosOtherInfo {
	s.AvgDelayTime = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosOtherInfo) SetFlightCancelRate(v string) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosOtherInfo {
	s.FlightCancelRate = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosOtherInfo) SetJetBridgeRate(v string) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosOtherInfo {
	s.JetBridgeRate = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosOtherInfo) SetOnTimeRate(v string) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosOtherInfo {
	s.OnTimeRate = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosOtherInfo) SetWifi(v bool) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosOtherInfo {
	s.Wifi = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosOtherInfo) Validate() error {
	return dara.Validate(s)
}

type IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark struct {
	// example:
	//
	// 前往菲律宾的旅客，请确保持有往返纸质行程单以及纸质签证办理登记手续，否则可能会被当地政府拒绝入境
	DepCityVisaRemark *string `json:"dep_city_visa_remark,omitempty" xml:"dep_city_visa_remark,omitempty"`
	// example:
	//
	// 0
	DepCityVisaType     *int32    `json:"dep_city_visa_type,omitempty" xml:"dep_city_visa_type,omitempty"`
	StopCityVisaRemarks []*string `json:"stop_city_visa_remarks,omitempty" xml:"stop_city_visa_remarks,omitempty" type:"Repeated"`
	StopCityVisaTypes   []*int32  `json:"stop_city_visa_types,omitempty" xml:"stop_city_visa_types,omitempty" type:"Repeated"`
}

func (s IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark) GetDepCityVisaRemark() *string {
	return s.DepCityVisaRemark
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark) GetDepCityVisaType() *int32 {
	return s.DepCityVisaType
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark) GetStopCityVisaRemarks() []*string {
	return s.StopCityVisaRemarks
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark) GetStopCityVisaTypes() []*int32 {
	return s.StopCityVisaTypes
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark) SetDepCityVisaRemark(v string) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark {
	s.DepCityVisaRemark = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark) SetDepCityVisaType(v int32) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark {
	s.DepCityVisaType = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark) SetStopCityVisaRemarks(v []*string) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark {
	s.StopCityVisaRemarks = v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark) SetStopCityVisaTypes(v []*int32) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark {
	s.StopCityVisaTypes = v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark) Validate() error {
	return dara.Validate(s)
}

type IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosTicketingAirlineInfo struct {
	// example:
	//
	// CA
	AirlineCode *string `json:"airline_code,omitempty" xml:"airline_code,omitempty"`
	// example:
	//
	// 中国国航
	AirlineName *string `json:"airline_name,omitempty" xml:"airline_name,omitempty"`
	// example:
	//
	// false
	CheapAirline *bool `json:"cheap_airline,omitempty" xml:"cheap_airline,omitempty"`
	// example:
	//
	// https://
	IconUrl *string `json:"icon_url,omitempty" xml:"icon_url,omitempty"`
	// example:
	//
	// 国航
	ShortName *string `json:"short_name,omitempty" xml:"short_name,omitempty"`
}

func (s IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosTicketingAirlineInfo) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosTicketingAirlineInfo) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosTicketingAirlineInfo) GetAirlineCode() *string {
	return s.AirlineCode
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosTicketingAirlineInfo) GetAirlineName() *string {
	return s.AirlineName
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosTicketingAirlineInfo) GetCheapAirline() *bool {
	return s.CheapAirline
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosTicketingAirlineInfo) GetIconUrl() *string {
	return s.IconUrl
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosTicketingAirlineInfo) GetShortName() *string {
	return s.ShortName
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosTicketingAirlineInfo) SetAirlineCode(v string) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosTicketingAirlineInfo {
	s.AirlineCode = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosTicketingAirlineInfo) SetAirlineName(v string) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosTicketingAirlineInfo {
	s.AirlineName = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosTicketingAirlineInfo) SetCheapAirline(v bool) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosTicketingAirlineInfo {
	s.CheapAirline = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosTicketingAirlineInfo) SetIconUrl(v string) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosTicketingAirlineInfo {
	s.IconUrl = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosTicketingAirlineInfo) SetShortName(v string) *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosTicketingAirlineInfo {
	s.ShortName = &v
	return s
}

func (s *IntlFlightReShopListSearchResponseBodyModuleReShopItemListFlightJourneyInfosFlightSegmentInfosTicketingAirlineInfo) Validate() error {
	return dara.Validate(s)
}
