// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightReShopOtaSearchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *IntlFlightReShopOtaSearchResponseBody
	GetCode() *string
	SetMessage(v string) *IntlFlightReShopOtaSearchResponseBody
	GetMessage() *string
	SetModule(v *IntlFlightReShopOtaSearchResponseBodyModule) *IntlFlightReShopOtaSearchResponseBody
	GetModule() *IntlFlightReShopOtaSearchResponseBodyModule
	SetRequestId(v string) *IntlFlightReShopOtaSearchResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *IntlFlightReShopOtaSearchResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *IntlFlightReShopOtaSearchResponseBody
	GetTraceId() *string
}

type IntlFlightReShopOtaSearchResponseBody struct {
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// 成功
	Message *string                                      `json:"message,omitempty" xml:"message,omitempty"`
	Module  *IntlFlightReShopOtaSearchResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
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

func (s IntlFlightReShopOtaSearchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopOtaSearchResponseBody) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopOtaSearchResponseBody) GetCode() *string {
	return s.Code
}

func (s *IntlFlightReShopOtaSearchResponseBody) GetMessage() *string {
	return s.Message
}

func (s *IntlFlightReShopOtaSearchResponseBody) GetModule() *IntlFlightReShopOtaSearchResponseBodyModule {
	return s.Module
}

func (s *IntlFlightReShopOtaSearchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *IntlFlightReShopOtaSearchResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *IntlFlightReShopOtaSearchResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *IntlFlightReShopOtaSearchResponseBody) SetCode(v string) *IntlFlightReShopOtaSearchResponseBody {
	s.Code = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBody) SetMessage(v string) *IntlFlightReShopOtaSearchResponseBody {
	s.Message = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBody) SetModule(v *IntlFlightReShopOtaSearchResponseBodyModule) *IntlFlightReShopOtaSearchResponseBody {
	s.Module = v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBody) SetRequestId(v string) *IntlFlightReShopOtaSearchResponseBody {
	s.RequestId = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBody) SetSuccess(v bool) *IntlFlightReShopOtaSearchResponseBody {
	s.Success = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBody) SetTraceId(v string) *IntlFlightReShopOtaSearchResponseBody {
	s.TraceId = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type IntlFlightReShopOtaSearchResponseBodyModule struct {
	FlightJourneyInfos []*IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfos `json:"flight_journey_infos,omitempty" xml:"flight_journey_infos,omitempty" type:"Repeated"`
	NeedContinue       *bool                                                            `json:"need_continue,omitempty" xml:"need_continue,omitempty"`
	NextReqWaitTime    *int32                                                           `json:"next_req_wait_time,omitempty" xml:"next_req_wait_time,omitempty"`
	ReShopItemList     []*IntlFlightReShopOtaSearchResponseBodyModuleReShopItemList     `json:"re_shop_item_list,omitempty" xml:"re_shop_item_list,omitempty" type:"Repeated"`
	Token              *string                                                          `json:"token,omitempty" xml:"token,omitempty"`
}

func (s IntlFlightReShopOtaSearchResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopOtaSearchResponseBodyModule) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopOtaSearchResponseBodyModule) GetFlightJourneyInfos() []*IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfos {
	return s.FlightJourneyInfos
}

func (s *IntlFlightReShopOtaSearchResponseBodyModule) GetNeedContinue() *bool {
	return s.NeedContinue
}

func (s *IntlFlightReShopOtaSearchResponseBodyModule) GetNextReqWaitTime() *int32 {
	return s.NextReqWaitTime
}

func (s *IntlFlightReShopOtaSearchResponseBodyModule) GetReShopItemList() []*IntlFlightReShopOtaSearchResponseBodyModuleReShopItemList {
	return s.ReShopItemList
}

func (s *IntlFlightReShopOtaSearchResponseBodyModule) GetToken() *string {
	return s.Token
}

func (s *IntlFlightReShopOtaSearchResponseBodyModule) SetFlightJourneyInfos(v []*IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfos) *IntlFlightReShopOtaSearchResponseBodyModule {
	s.FlightJourneyInfos = v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModule) SetNeedContinue(v bool) *IntlFlightReShopOtaSearchResponseBodyModule {
	s.NeedContinue = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModule) SetNextReqWaitTime(v int32) *IntlFlightReShopOtaSearchResponseBodyModule {
	s.NextReqWaitTime = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModule) SetReShopItemList(v []*IntlFlightReShopOtaSearchResponseBodyModuleReShopItemList) *IntlFlightReShopOtaSearchResponseBodyModule {
	s.ReShopItemList = v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModule) SetToken(v string) *IntlFlightReShopOtaSearchResponseBodyModule {
	s.Token = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModule) Validate() error {
	if s.FlightJourneyInfos != nil {
		for _, item := range s.FlightJourneyInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
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

type IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfos struct {
	ArrCityCode        *string                                                                            `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	ArrCityName        *string                                                                            `json:"arr_city_name,omitempty" xml:"arr_city_name,omitempty"`
	ArrTime            *string                                                                            `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	DepCityCode        *string                                                                            `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	DepCityName        *string                                                                            `json:"dep_city_name,omitempty" xml:"dep_city_name,omitempty"`
	DepTime            *string                                                                            `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	Duration           *int32                                                                             `json:"duration,omitempty" xml:"duration,omitempty"`
	FlightSegmentInfos []*IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos `json:"flight_segment_infos,omitempty" xml:"flight_segment_infos,omitempty" type:"Repeated"`
	JourneyIndex       *int32                                                                             `json:"journey_index,omitempty" xml:"journey_index,omitempty"`
	TransferTime       *int32                                                                             `json:"transfer_time,omitempty" xml:"transfer_time,omitempty"`
}

func (s IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfos) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfos) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfos) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfos) GetArrCityName() *string {
	return s.ArrCityName
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfos) GetArrTime() *string {
	return s.ArrTime
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfos) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfos) GetDepCityName() *string {
	return s.DepCityName
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfos) GetDepTime() *string {
	return s.DepTime
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfos) GetDuration() *int32 {
	return s.Duration
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfos) GetFlightSegmentInfos() []*IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	return s.FlightSegmentInfos
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfos) GetJourneyIndex() *int32 {
	return s.JourneyIndex
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfos) GetTransferTime() *int32 {
	return s.TransferTime
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfos) SetArrCityCode(v string) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfos {
	s.ArrCityCode = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfos) SetArrCityName(v string) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfos {
	s.ArrCityName = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfos) SetArrTime(v string) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfos {
	s.ArrTime = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfos) SetDepCityCode(v string) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfos {
	s.DepCityCode = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfos) SetDepCityName(v string) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfos {
	s.DepCityName = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfos) SetDepTime(v string) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfos {
	s.DepTime = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfos) SetDuration(v int32) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfos {
	s.Duration = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfos) SetFlightSegmentInfos(v []*IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfos {
	s.FlightSegmentInfos = v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfos) SetJourneyIndex(v int32) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfos {
	s.JourneyIndex = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfos) SetTransferTime(v int32) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfos {
	s.TransferTime = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfos) Validate() error {
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

type IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos struct {
	AirlineInfo          *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo          `json:"airline_info,omitempty" xml:"airline_info,omitempty" type:"Struct"`
	ArrAirportInfo       *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo       `json:"arr_airport_info,omitempty" xml:"arr_airport_info,omitempty" type:"Struct"`
	ArrCityCode          *string                                                                                              `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	ArrCityName          *string                                                                                              `json:"arr_city_name,omitempty" xml:"arr_city_name,omitempty"`
	ArrTime              *string                                                                                              `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	ArrTimeUTC           *string                                                                                              `json:"arr_time_u_t_c,omitempty" xml:"arr_time_u_t_c,omitempty"`
	DepAirportInfo       *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo       `json:"dep_airport_info,omitempty" xml:"dep_airport_info,omitempty" type:"Struct"`
	DepCityCode          *string                                                                                              `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	DepCityName          *string                                                                                              `json:"dep_city_name,omitempty" xml:"dep_city_name,omitempty"`
	DepTime              *string                                                                                              `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	DepTimeUTC           *string                                                                                              `json:"dep_time_u_t_c,omitempty" xml:"dep_time_u_t_c,omitempty"`
	Duration             *int32                                                                                               `json:"duration,omitempty" xml:"duration,omitempty"`
	FlightNo             *string                                                                                              `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
	FlightShareInfo      *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfo      `json:"flight_share_info,omitempty" xml:"flight_share_info,omitempty" type:"Struct"`
	FlightSize           *string                                                                                              `json:"flight_size,omitempty" xml:"flight_size,omitempty"`
	FlightStopInfoList   []*IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoList `json:"flight_stop_info_list,omitempty" xml:"flight_stop_info_list,omitempty" type:"Repeated"`
	FlightType           *string                                                                                              `json:"flight_type,omitempty" xml:"flight_type,omitempty"`
	JourneyIndex         *int32                                                                                               `json:"journey_index,omitempty" xml:"journey_index,omitempty"`
	LuggageDirectInfo    *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosLuggageDirectInfo    `json:"luggage_direct_info,omitempty" xml:"luggage_direct_info,omitempty" type:"Struct"`
	Manufacturer         *string                                                                                              `json:"manufacturer,omitempty" xml:"manufacturer,omitempty"`
	Meal                 *int32                                                                                               `json:"meal,omitempty" xml:"meal,omitempty"`
	MealDesc             *string                                                                                              `json:"meal_desc,omitempty" xml:"meal_desc,omitempty"`
	Miles                *int32                                                                                               `json:"miles,omitempty" xml:"miles,omitempty"`
	OnTimeRate           *string                                                                                              `json:"on_time_rate,omitempty" xml:"on_time_rate,omitempty"`
	OneMore              *int32                                                                                               `json:"one_more,omitempty" xml:"one_more,omitempty"`
	OneMoreShow          *string                                                                                              `json:"one_more_show,omitempty" xml:"one_more_show,omitempty"`
	OtherInfo            *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosOtherInfo            `json:"other_info,omitempty" xml:"other_info,omitempty" type:"Struct"`
	SegmentIndex         *int32                                                                                               `json:"segment_index,omitempty" xml:"segment_index,omitempty"`
	SegmentKey           *string                                                                                              `json:"segment_key,omitempty" xml:"segment_key,omitempty"`
	SegmentVisaRemark    *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark    `json:"segment_visa_remark,omitempty" xml:"segment_visa_remark,omitempty" type:"Struct"`
	Share                *bool                                                                                                `json:"share,omitempty" xml:"share,omitempty"`
	ShortFlightSize      *string                                                                                              `json:"short_flight_size,omitempty" xml:"short_flight_size,omitempty"`
	Stop                 *bool                                                                                                `json:"stop,omitempty" xml:"stop,omitempty"`
	TicketingAirlineInfo *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosTicketingAirlineInfo `json:"ticketing_airline_info,omitempty" xml:"ticketing_airline_info,omitempty" type:"Struct"`
	TotalTime            *string                                                                                              `json:"total_time,omitempty" xml:"total_time,omitempty"`
}

func (s IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetAirlineInfo() *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo {
	return s.AirlineInfo
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetArrAirportInfo() *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo {
	return s.ArrAirportInfo
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetArrCityName() *string {
	return s.ArrCityName
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetArrTime() *string {
	return s.ArrTime
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetArrTimeUTC() *string {
	return s.ArrTimeUTC
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetDepAirportInfo() *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo {
	return s.DepAirportInfo
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetDepCityName() *string {
	return s.DepCityName
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetDepTime() *string {
	return s.DepTime
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetDepTimeUTC() *string {
	return s.DepTimeUTC
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetDuration() *int32 {
	return s.Duration
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetFlightNo() *string {
	return s.FlightNo
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetFlightShareInfo() *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfo {
	return s.FlightShareInfo
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetFlightSize() *string {
	return s.FlightSize
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetFlightStopInfoList() []*IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoList {
	return s.FlightStopInfoList
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetFlightType() *string {
	return s.FlightType
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetJourneyIndex() *int32 {
	return s.JourneyIndex
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetLuggageDirectInfo() *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosLuggageDirectInfo {
	return s.LuggageDirectInfo
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetManufacturer() *string {
	return s.Manufacturer
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetMeal() *int32 {
	return s.Meal
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetMealDesc() *string {
	return s.MealDesc
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetMiles() *int32 {
	return s.Miles
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetOnTimeRate() *string {
	return s.OnTimeRate
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetOneMore() *int32 {
	return s.OneMore
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetOneMoreShow() *string {
	return s.OneMoreShow
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetOtherInfo() *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosOtherInfo {
	return s.OtherInfo
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetSegmentIndex() *int32 {
	return s.SegmentIndex
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetSegmentKey() *string {
	return s.SegmentKey
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetSegmentVisaRemark() *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark {
	return s.SegmentVisaRemark
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetShare() *bool {
	return s.Share
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetShortFlightSize() *string {
	return s.ShortFlightSize
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetStop() *bool {
	return s.Stop
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetTicketingAirlineInfo() *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosTicketingAirlineInfo {
	return s.TicketingAirlineInfo
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetTotalTime() *string {
	return s.TotalTime
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetAirlineInfo(v *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.AirlineInfo = v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetArrAirportInfo(v *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.ArrAirportInfo = v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetArrCityCode(v string) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.ArrCityCode = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetArrCityName(v string) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.ArrCityName = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetArrTime(v string) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.ArrTime = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetArrTimeUTC(v string) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.ArrTimeUTC = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetDepAirportInfo(v *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.DepAirportInfo = v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetDepCityCode(v string) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.DepCityCode = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetDepCityName(v string) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.DepCityName = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetDepTime(v string) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.DepTime = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetDepTimeUTC(v string) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.DepTimeUTC = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetDuration(v int32) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.Duration = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetFlightNo(v string) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.FlightNo = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetFlightShareInfo(v *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfo) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.FlightShareInfo = v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetFlightSize(v string) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.FlightSize = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetFlightStopInfoList(v []*IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.FlightStopInfoList = v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetFlightType(v string) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.FlightType = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetJourneyIndex(v int32) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.JourneyIndex = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetLuggageDirectInfo(v *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosLuggageDirectInfo) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.LuggageDirectInfo = v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetManufacturer(v string) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.Manufacturer = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetMeal(v int32) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.Meal = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetMealDesc(v string) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.MealDesc = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetMiles(v int32) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.Miles = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetOnTimeRate(v string) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.OnTimeRate = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetOneMore(v int32) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.OneMore = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetOneMoreShow(v string) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.OneMoreShow = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetOtherInfo(v *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosOtherInfo) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.OtherInfo = v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetSegmentIndex(v int32) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.SegmentIndex = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetSegmentKey(v string) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.SegmentKey = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetSegmentVisaRemark(v *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.SegmentVisaRemark = v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetShare(v bool) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.Share = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetShortFlightSize(v string) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.ShortFlightSize = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetStop(v bool) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.Stop = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetTicketingAirlineInfo(v *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosTicketingAirlineInfo) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.TicketingAirlineInfo = v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetTotalTime(v string) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.TotalTime = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) Validate() error {
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

type IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo struct {
	AirlineCode  *string `json:"airline_code,omitempty" xml:"airline_code,omitempty"`
	AirlineName  *string `json:"airline_name,omitempty" xml:"airline_name,omitempty"`
	CheapAirline *bool   `json:"cheap_airline,omitempty" xml:"cheap_airline,omitempty"`
	IconUrl      *string `json:"icon_url,omitempty" xml:"icon_url,omitempty"`
	ShortName    *string `json:"short_name,omitempty" xml:"short_name,omitempty"`
}

func (s IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo) GetAirlineCode() *string {
	return s.AirlineCode
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo) GetAirlineName() *string {
	return s.AirlineName
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo) GetCheapAirline() *bool {
	return s.CheapAirline
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo) GetIconUrl() *string {
	return s.IconUrl
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo) GetShortName() *string {
	return s.ShortName
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo) SetAirlineCode(v string) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo {
	s.AirlineCode = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo) SetAirlineName(v string) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo {
	s.AirlineName = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo) SetCheapAirline(v bool) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo {
	s.CheapAirline = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo) SetIconUrl(v string) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo {
	s.IconUrl = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo) SetShortName(v string) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo {
	s.ShortName = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo) Validate() error {
	return dara.Validate(s)
}

type IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo struct {
	AirportCode      *string `json:"airport_code,omitempty" xml:"airport_code,omitempty"`
	AirportName      *string `json:"airport_name,omitempty" xml:"airport_name,omitempty"`
	AirportShortName *string `json:"airport_short_name,omitempty" xml:"airport_short_name,omitempty"`
	Terminal         *string `json:"terminal,omitempty" xml:"terminal,omitempty"`
}

func (s IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo) GetAirportCode() *string {
	return s.AirportCode
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo) GetAirportName() *string {
	return s.AirportName
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo) GetAirportShortName() *string {
	return s.AirportShortName
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo) GetTerminal() *string {
	return s.Terminal
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo) SetAirportCode(v string) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo {
	s.AirportCode = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo) SetAirportName(v string) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo {
	s.AirportName = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo) SetAirportShortName(v string) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo {
	s.AirportShortName = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo) SetTerminal(v string) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo {
	s.Terminal = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo) Validate() error {
	return dara.Validate(s)
}

type IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo struct {
	AirportCode      *string `json:"airport_code,omitempty" xml:"airport_code,omitempty"`
	AirportName      *string `json:"airport_name,omitempty" xml:"airport_name,omitempty"`
	AirportShortName *string `json:"airport_short_name,omitempty" xml:"airport_short_name,omitempty"`
	Terminal         *string `json:"terminal,omitempty" xml:"terminal,omitempty"`
}

func (s IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo) GetAirportCode() *string {
	return s.AirportCode
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo) GetAirportName() *string {
	return s.AirportName
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo) GetAirportShortName() *string {
	return s.AirportShortName
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo) GetTerminal() *string {
	return s.Terminal
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo) SetAirportCode(v string) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo {
	s.AirportCode = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo) SetAirportName(v string) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo {
	s.AirportName = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo) SetAirportShortName(v string) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo {
	s.AirportShortName = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo) SetTerminal(v string) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo {
	s.Terminal = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo) Validate() error {
	return dara.Validate(s)
}

type IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfo struct {
	OperatingAirlineInfo *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo `json:"operating_airline_info,omitempty" xml:"operating_airline_info,omitempty" type:"Struct"`
	OperatingFlightNo    *string                                                                                                             `json:"operating_flight_no,omitempty" xml:"operating_flight_no,omitempty"`
}

func (s IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfo) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfo) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfo) GetOperatingAirlineInfo() *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo {
	return s.OperatingAirlineInfo
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfo) GetOperatingFlightNo() *string {
	return s.OperatingFlightNo
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfo) SetOperatingAirlineInfo(v *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfo {
	s.OperatingAirlineInfo = v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfo) SetOperatingFlightNo(v string) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfo {
	s.OperatingFlightNo = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfo) Validate() error {
	if s.OperatingAirlineInfo != nil {
		if err := s.OperatingAirlineInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo struct {
	AirlineCode  *string `json:"airline_code,omitempty" xml:"airline_code,omitempty"`
	AirlineName  *string `json:"airline_name,omitempty" xml:"airline_name,omitempty"`
	CheapAirline *bool   `json:"cheap_airline,omitempty" xml:"cheap_airline,omitempty"`
	IconUrl      *string `json:"icon_url,omitempty" xml:"icon_url,omitempty"`
	ShortName    *string `json:"short_name,omitempty" xml:"short_name,omitempty"`
}

func (s IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) GetAirlineCode() *string {
	return s.AirlineCode
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) GetAirlineName() *string {
	return s.AirlineName
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) GetCheapAirline() *bool {
	return s.CheapAirline
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) GetIconUrl() *string {
	return s.IconUrl
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) GetShortName() *string {
	return s.ShortName
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) SetAirlineCode(v string) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo {
	s.AirlineCode = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) SetAirlineName(v string) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo {
	s.AirlineName = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) SetCheapAirline(v bool) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo {
	s.CheapAirline = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) SetIconUrl(v string) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo {
	s.IconUrl = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) SetShortName(v string) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo {
	s.ShortName = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) Validate() error {
	return dara.Validate(s)
}

type IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoList struct {
	StopAirport           *string                                                                                                                 `json:"stop_airport,omitempty" xml:"stop_airport,omitempty"`
	StopAirportCountyInfo *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoListStopAirportCountyInfo `json:"stop_airport_county_info,omitempty" xml:"stop_airport_county_info,omitempty" type:"Struct"`
	StopAirportName       *string                                                                                                                 `json:"stop_airport_name,omitempty" xml:"stop_airport_name,omitempty"`
	StopArrTerm           *string                                                                                                                 `json:"stop_arr_term,omitempty" xml:"stop_arr_term,omitempty"`
	StopArrTime           *string                                                                                                                 `json:"stop_arr_time,omitempty" xml:"stop_arr_time,omitempty"`
	StopCityCode          *string                                                                                                                 `json:"stop_city_code,omitempty" xml:"stop_city_code,omitempty"`
	StopCityName          *string                                                                                                                 `json:"stop_city_name,omitempty" xml:"stop_city_name,omitempty"`
	StopDepTerm           *string                                                                                                                 `json:"stop_dep_term,omitempty" xml:"stop_dep_term,omitempty"`
	StopDepTime           *string                                                                                                                 `json:"stop_dep_time,omitempty" xml:"stop_dep_time,omitempty"`
	StopTime              *string                                                                                                                 `json:"stop_time,omitempty" xml:"stop_time,omitempty"`
}

func (s IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) GetStopAirport() *string {
	return s.StopAirport
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) GetStopAirportCountyInfo() *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoListStopAirportCountyInfo {
	return s.StopAirportCountyInfo
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) GetStopAirportName() *string {
	return s.StopAirportName
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) GetStopArrTerm() *string {
	return s.StopArrTerm
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) GetStopArrTime() *string {
	return s.StopArrTime
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) GetStopCityCode() *string {
	return s.StopCityCode
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) GetStopCityName() *string {
	return s.StopCityName
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) GetStopDepTerm() *string {
	return s.StopDepTerm
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) GetStopDepTime() *string {
	return s.StopDepTime
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) GetStopTime() *string {
	return s.StopTime
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) SetStopAirport(v string) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoList {
	s.StopAirport = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) SetStopAirportCountyInfo(v *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoListStopAirportCountyInfo) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoList {
	s.StopAirportCountyInfo = v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) SetStopAirportName(v string) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoList {
	s.StopAirportName = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) SetStopArrTerm(v string) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoList {
	s.StopArrTerm = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) SetStopArrTime(v string) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoList {
	s.StopArrTime = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) SetStopCityCode(v string) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoList {
	s.StopCityCode = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) SetStopCityName(v string) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoList {
	s.StopCityName = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) SetStopDepTerm(v string) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoList {
	s.StopDepTerm = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) SetStopDepTime(v string) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoList {
	s.StopDepTime = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) SetStopTime(v string) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoList {
	s.StopTime = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) Validate() error {
	if s.StopAirportCountyInfo != nil {
		if err := s.StopAirportCountyInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoListStopAirportCountyInfo struct {
	Adcode                *string `json:"adcode,omitempty" xml:"adcode,omitempty"`
	AirportCityCode       *string `json:"airport_city_code,omitempty" xml:"airport_city_code,omitempty"`
	AirportCityName       *string `json:"airport_city_name,omitempty" xml:"airport_city_name,omitempty"`
	AirportCode           *string `json:"airport_code,omitempty" xml:"airport_code,omitempty"`
	AirportName           *string `json:"airport_name,omitempty" xml:"airport_name,omitempty"`
	AirportParentCityName *string `json:"airport_parent_city_name,omitempty" xml:"airport_parent_city_name,omitempty"`
	CountyCityAdcode      *string `json:"county_city_adcode,omitempty" xml:"county_city_adcode,omitempty"`
	CountyCityName        *string `json:"county_city_name,omitempty" xml:"county_city_name,omitempty"`
	PrefectureCityAdcode  *string `json:"prefecture_city_adcode,omitempty" xml:"prefecture_city_adcode,omitempty"`
	PrefectureCityName    *string `json:"prefecture_city_name,omitempty" xml:"prefecture_city_name,omitempty"`
}

func (s IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoListStopAirportCountyInfo) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoListStopAirportCountyInfo) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoListStopAirportCountyInfo) GetAdcode() *string {
	return s.Adcode
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoListStopAirportCountyInfo) GetAirportCityCode() *string {
	return s.AirportCityCode
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoListStopAirportCountyInfo) GetAirportCityName() *string {
	return s.AirportCityName
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoListStopAirportCountyInfo) GetAirportCode() *string {
	return s.AirportCode
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoListStopAirportCountyInfo) GetAirportName() *string {
	return s.AirportName
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoListStopAirportCountyInfo) GetAirportParentCityName() *string {
	return s.AirportParentCityName
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoListStopAirportCountyInfo) GetCountyCityAdcode() *string {
	return s.CountyCityAdcode
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoListStopAirportCountyInfo) GetCountyCityName() *string {
	return s.CountyCityName
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoListStopAirportCountyInfo) GetPrefectureCityAdcode() *string {
	return s.PrefectureCityAdcode
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoListStopAirportCountyInfo) GetPrefectureCityName() *string {
	return s.PrefectureCityName
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoListStopAirportCountyInfo) SetAdcode(v string) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoListStopAirportCountyInfo {
	s.Adcode = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoListStopAirportCountyInfo) SetAirportCityCode(v string) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoListStopAirportCountyInfo {
	s.AirportCityCode = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoListStopAirportCountyInfo) SetAirportCityName(v string) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoListStopAirportCountyInfo {
	s.AirportCityName = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoListStopAirportCountyInfo) SetAirportCode(v string) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoListStopAirportCountyInfo {
	s.AirportCode = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoListStopAirportCountyInfo) SetAirportName(v string) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoListStopAirportCountyInfo {
	s.AirportName = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoListStopAirportCountyInfo) SetAirportParentCityName(v string) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoListStopAirportCountyInfo {
	s.AirportParentCityName = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoListStopAirportCountyInfo) SetCountyCityAdcode(v string) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoListStopAirportCountyInfo {
	s.CountyCityAdcode = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoListStopAirportCountyInfo) SetCountyCityName(v string) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoListStopAirportCountyInfo {
	s.CountyCityName = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoListStopAirportCountyInfo) SetPrefectureCityAdcode(v string) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoListStopAirportCountyInfo {
	s.PrefectureCityAdcode = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoListStopAirportCountyInfo) SetPrefectureCityName(v string) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoListStopAirportCountyInfo {
	s.PrefectureCityName = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoListStopAirportCountyInfo) Validate() error {
	return dara.Validate(s)
}

type IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosLuggageDirectInfo struct {
	DepCityLuggageDirect  *int32 `json:"dep_city_luggage_direct,omitempty" xml:"dep_city_luggage_direct,omitempty"`
	StopCityLuggageDirect *int32 `json:"stop_city_luggage_direct,omitempty" xml:"stop_city_luggage_direct,omitempty"`
}

func (s IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosLuggageDirectInfo) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosLuggageDirectInfo) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosLuggageDirectInfo) GetDepCityLuggageDirect() *int32 {
	return s.DepCityLuggageDirect
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosLuggageDirectInfo) GetStopCityLuggageDirect() *int32 {
	return s.StopCityLuggageDirect
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosLuggageDirectInfo) SetDepCityLuggageDirect(v int32) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosLuggageDirectInfo {
	s.DepCityLuggageDirect = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosLuggageDirectInfo) SetStopCityLuggageDirect(v int32) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosLuggageDirectInfo {
	s.StopCityLuggageDirect = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosLuggageDirectInfo) Validate() error {
	return dara.Validate(s)
}

type IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosOtherInfo struct {
	AircraftAge      *string `json:"aircraft_age,omitempty" xml:"aircraft_age,omitempty"`
	AvgDelayTime     *string `json:"avg_delay_time,omitempty" xml:"avg_delay_time,omitempty"`
	FlightCancelRate *string `json:"flight_cancel_rate,omitempty" xml:"flight_cancel_rate,omitempty"`
	JetBridgeRate    *string `json:"jet_bridge_rate,omitempty" xml:"jet_bridge_rate,omitempty"`
	OnTimeRate       *string `json:"on_time_rate,omitempty" xml:"on_time_rate,omitempty"`
	Wifi             *bool   `json:"wifi,omitempty" xml:"wifi,omitempty"`
}

func (s IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosOtherInfo) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosOtherInfo) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosOtherInfo) GetAircraftAge() *string {
	return s.AircraftAge
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosOtherInfo) GetAvgDelayTime() *string {
	return s.AvgDelayTime
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosOtherInfo) GetFlightCancelRate() *string {
	return s.FlightCancelRate
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosOtherInfo) GetJetBridgeRate() *string {
	return s.JetBridgeRate
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosOtherInfo) GetOnTimeRate() *string {
	return s.OnTimeRate
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosOtherInfo) GetWifi() *bool {
	return s.Wifi
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosOtherInfo) SetAircraftAge(v string) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosOtherInfo {
	s.AircraftAge = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosOtherInfo) SetAvgDelayTime(v string) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosOtherInfo {
	s.AvgDelayTime = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosOtherInfo) SetFlightCancelRate(v string) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosOtherInfo {
	s.FlightCancelRate = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosOtherInfo) SetJetBridgeRate(v string) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosOtherInfo {
	s.JetBridgeRate = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosOtherInfo) SetOnTimeRate(v string) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosOtherInfo {
	s.OnTimeRate = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosOtherInfo) SetWifi(v bool) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosOtherInfo {
	s.Wifi = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosOtherInfo) Validate() error {
	return dara.Validate(s)
}

type IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark struct {
	DepCityVisaRemark   *string   `json:"dep_city_visa_remark,omitempty" xml:"dep_city_visa_remark,omitempty"`
	DepCityVisaType     *int32    `json:"dep_city_visa_type,omitempty" xml:"dep_city_visa_type,omitempty"`
	StopCityVisaRemarks []*string `json:"stop_city_visa_remarks,omitempty" xml:"stop_city_visa_remarks,omitempty" type:"Repeated"`
	StopCityVisaTypes   []*int32  `json:"stop_city_visa_types,omitempty" xml:"stop_city_visa_types,omitempty" type:"Repeated"`
}

func (s IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark) GetDepCityVisaRemark() *string {
	return s.DepCityVisaRemark
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark) GetDepCityVisaType() *int32 {
	return s.DepCityVisaType
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark) GetStopCityVisaRemarks() []*string {
	return s.StopCityVisaRemarks
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark) GetStopCityVisaTypes() []*int32 {
	return s.StopCityVisaTypes
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark) SetDepCityVisaRemark(v string) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark {
	s.DepCityVisaRemark = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark) SetDepCityVisaType(v int32) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark {
	s.DepCityVisaType = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark) SetStopCityVisaRemarks(v []*string) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark {
	s.StopCityVisaRemarks = v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark) SetStopCityVisaTypes(v []*int32) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark {
	s.StopCityVisaTypes = v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark) Validate() error {
	return dara.Validate(s)
}

type IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosTicketingAirlineInfo struct {
	AirlineCode  *string `json:"airline_code,omitempty" xml:"airline_code,omitempty"`
	AirlineName  *string `json:"airline_name,omitempty" xml:"airline_name,omitempty"`
	CheapAirline *bool   `json:"cheap_airline,omitempty" xml:"cheap_airline,omitempty"`
	IconUrl      *string `json:"icon_url,omitempty" xml:"icon_url,omitempty"`
	ShortName    *string `json:"short_name,omitempty" xml:"short_name,omitempty"`
}

func (s IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosTicketingAirlineInfo) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosTicketingAirlineInfo) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosTicketingAirlineInfo) GetAirlineCode() *string {
	return s.AirlineCode
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosTicketingAirlineInfo) GetAirlineName() *string {
	return s.AirlineName
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosTicketingAirlineInfo) GetCheapAirline() *bool {
	return s.CheapAirline
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosTicketingAirlineInfo) GetIconUrl() *string {
	return s.IconUrl
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosTicketingAirlineInfo) GetShortName() *string {
	return s.ShortName
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosTicketingAirlineInfo) SetAirlineCode(v string) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosTicketingAirlineInfo {
	s.AirlineCode = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosTicketingAirlineInfo) SetAirlineName(v string) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosTicketingAirlineInfo {
	s.AirlineName = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosTicketingAirlineInfo) SetCheapAirline(v bool) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosTicketingAirlineInfo {
	s.CheapAirline = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosTicketingAirlineInfo) SetIconUrl(v string) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosTicketingAirlineInfo {
	s.IconUrl = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosTicketingAirlineInfo) SetShortName(v string) *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosTicketingAirlineInfo {
	s.ShortName = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosTicketingAirlineInfo) Validate() error {
	return dara.Validate(s)
}

type IntlFlightReShopOtaSearchResponseBodyModuleReShopItemList struct {
	ItemId          *string                                                              `json:"item_id,omitempty" xml:"item_id,omitempty"`
	ItemType        *string                                                              `json:"item_type,omitempty" xml:"item_type,omitempty"`
	ShoppingItemMap map[string]*ModuleReShopItemListShoppingItemMapValue                 `json:"shopping_item_map,omitempty" xml:"shopping_item_map,omitempty"`
	SubItems        []*IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItems `json:"sub_items,omitempty" xml:"sub_items,omitempty" type:"Repeated"`
}

func (s IntlFlightReShopOtaSearchResponseBodyModuleReShopItemList) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopOtaSearchResponseBodyModuleReShopItemList) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemList) GetItemId() *string {
	return s.ItemId
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemList) GetItemType() *string {
	return s.ItemType
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemList) GetShoppingItemMap() map[string]*ModuleReShopItemListShoppingItemMapValue {
	return s.ShoppingItemMap
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemList) GetSubItems() []*IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItems {
	return s.SubItems
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemList) SetItemId(v string) *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemList {
	s.ItemId = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemList) SetItemType(v string) *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemList {
	s.ItemType = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemList) SetShoppingItemMap(v map[string]*ModuleReShopItemListShoppingItemMapValue) *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemList {
	s.ShoppingItemMap = v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemList) SetSubItems(v []*IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItems) *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemList {
	s.SubItems = v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemList) Validate() error {
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

type IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItems struct {
	BaggageRule         *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItemsBaggageRule           `json:"baggage_rule,omitempty" xml:"baggage_rule,omitempty" type:"Struct"`
	DiscountNum         *float64                                                                                `json:"discount_num,omitempty" xml:"discount_num,omitempty"`
	RefundChangeRule    *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItemsRefundChangeRule      `json:"refund_change_rule,omitempty" xml:"refund_change_rule,omitempty" type:"Struct"`
	SegmentKeys         []*string                                                                               `json:"segment_keys,omitempty" xml:"segment_keys,omitempty" type:"Repeated"`
	SegmentPositionList []*IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItemsSegmentPositionList `json:"segment_position_list,omitempty" xml:"segment_position_list,omitempty" type:"Repeated"`
	ShoppingItemMap     map[string]*ModuleReShopItemListSubItemsShoppingItemMapValue                            `json:"shopping_item_map,omitempty" xml:"shopping_item_map,omitempty"`
	UniqKey             *string                                                                                 `json:"uniq_key,omitempty" xml:"uniq_key,omitempty"`
}

func (s IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItems) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItems) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItems) GetBaggageRule() *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItemsBaggageRule {
	return s.BaggageRule
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItems) GetDiscountNum() *float64 {
	return s.DiscountNum
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItems) GetRefundChangeRule() *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItemsRefundChangeRule {
	return s.RefundChangeRule
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItems) GetSegmentKeys() []*string {
	return s.SegmentKeys
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItems) GetSegmentPositionList() []*IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItemsSegmentPositionList {
	return s.SegmentPositionList
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItems) GetShoppingItemMap() map[string]*ModuleReShopItemListSubItemsShoppingItemMapValue {
	return s.ShoppingItemMap
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItems) GetUniqKey() *string {
	return s.UniqKey
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItems) SetBaggageRule(v *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItemsBaggageRule) *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItems {
	s.BaggageRule = v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItems) SetDiscountNum(v float64) *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItems {
	s.DiscountNum = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItems) SetRefundChangeRule(v *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItemsRefundChangeRule) *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItems {
	s.RefundChangeRule = v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItems) SetSegmentKeys(v []*string) *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItems {
	s.SegmentKeys = v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItems) SetSegmentPositionList(v []*IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItemsSegmentPositionList) *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItems {
	s.SegmentPositionList = v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItems) SetShoppingItemMap(v map[string]*ModuleReShopItemListSubItemsShoppingItemMapValue) *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItems {
	s.ShoppingItemMap = v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItems) SetUniqKey(v string) *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItems {
	s.UniqKey = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItems) Validate() error {
	if s.BaggageRule != nil {
		if err := s.BaggageRule.Validate(); err != nil {
			return err
		}
	}
	if s.RefundChangeRule != nil {
		if err := s.RefundChangeRule.Validate(); err != nil {
			return err
		}
	}
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

type IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItemsBaggageRule struct {
	BaggageDescShowColor *string                                                                       `json:"baggage_desc_show_color,omitempty" xml:"baggage_desc_show_color,omitempty"`
	BaggageDigest        *string                                                                       `json:"baggage_digest,omitempty" xml:"baggage_digest,omitempty"`
	BaggageRuleDesc      *string                                                                       `json:"baggage_rule_desc,omitempty" xml:"baggage_rule_desc,omitempty"`
	OfferBaggageInfoMap  map[string][]*ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue `json:"offer_baggage_info_map,omitempty" xml:"offer_baggage_info_map,omitempty"`
	StructuredBaggage    *bool                                                                         `json:"structured_baggage,omitempty" xml:"structured_baggage,omitempty"`
}

func (s IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItemsBaggageRule) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItemsBaggageRule) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItemsBaggageRule) GetBaggageDescShowColor() *string {
	return s.BaggageDescShowColor
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItemsBaggageRule) GetBaggageDigest() *string {
	return s.BaggageDigest
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItemsBaggageRule) GetBaggageRuleDesc() *string {
	return s.BaggageRuleDesc
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItemsBaggageRule) GetOfferBaggageInfoMap() map[string][]*ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue {
	return s.OfferBaggageInfoMap
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItemsBaggageRule) GetStructuredBaggage() *bool {
	return s.StructuredBaggage
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItemsBaggageRule) SetBaggageDescShowColor(v string) *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItemsBaggageRule {
	s.BaggageDescShowColor = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItemsBaggageRule) SetBaggageDigest(v string) *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItemsBaggageRule {
	s.BaggageDigest = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItemsBaggageRule) SetBaggageRuleDesc(v string) *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItemsBaggageRule {
	s.BaggageRuleDesc = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItemsBaggageRule) SetOfferBaggageInfoMap(v map[string][]*ModuleReShopItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItemsBaggageRule {
	s.OfferBaggageInfoMap = v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItemsBaggageRule) SetStructuredBaggage(v bool) *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItemsBaggageRule {
	s.StructuredBaggage = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItemsBaggageRule) Validate() error {
	return dara.Validate(s)
}

type IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItemsRefundChangeRule struct {
	CancelFeeInd         *bool                                                                              `json:"cancel_fee_ind,omitempty" xml:"cancel_fee_ind,omitempty"`
	ChangeFeeInd         *bool                                                                              `json:"change_fee_ind,omitempty" xml:"change_fee_ind,omitempty"`
	ChangeRuleDesc       *string                                                                            `json:"change_rule_desc,omitempty" xml:"change_rule_desc,omitempty"`
	ChangeRuleShowColor  *string                                                                            `json:"change_rule_show_color,omitempty" xml:"change_rule_show_color,omitempty"`
	OfferPenaltyInfoMap  map[string][]*ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue `json:"offer_penalty_info_map,omitempty" xml:"offer_penalty_info_map,omitempty"`
	RefundChangeDigest   *string                                                                            `json:"refund_change_digest,omitempty" xml:"refund_change_digest,omitempty"`
	RefundChangeRuleDesc *string                                                                            `json:"refund_change_rule_desc,omitempty" xml:"refund_change_rule_desc,omitempty"`
	RefundRuleDesc       *string                                                                            `json:"refund_rule_desc,omitempty" xml:"refund_rule_desc,omitempty"`
	RefundRuleShowColor  *string                                                                            `json:"refund_rule_show_color,omitempty" xml:"refund_rule_show_color,omitempty"`
	StructuredRefund     *bool                                                                              `json:"structured_refund,omitempty" xml:"structured_refund,omitempty"`
}

func (s IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItemsRefundChangeRule) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItemsRefundChangeRule) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItemsRefundChangeRule) GetCancelFeeInd() *bool {
	return s.CancelFeeInd
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItemsRefundChangeRule) GetChangeFeeInd() *bool {
	return s.ChangeFeeInd
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItemsRefundChangeRule) GetChangeRuleDesc() *string {
	return s.ChangeRuleDesc
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItemsRefundChangeRule) GetChangeRuleShowColor() *string {
	return s.ChangeRuleShowColor
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItemsRefundChangeRule) GetOfferPenaltyInfoMap() map[string][]*ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue {
	return s.OfferPenaltyInfoMap
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItemsRefundChangeRule) GetRefundChangeDigest() *string {
	return s.RefundChangeDigest
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItemsRefundChangeRule) GetRefundChangeRuleDesc() *string {
	return s.RefundChangeRuleDesc
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItemsRefundChangeRule) GetRefundRuleDesc() *string {
	return s.RefundRuleDesc
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItemsRefundChangeRule) GetRefundRuleShowColor() *string {
	return s.RefundRuleShowColor
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItemsRefundChangeRule) GetStructuredRefund() *bool {
	return s.StructuredRefund
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItemsRefundChangeRule) SetCancelFeeInd(v bool) *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItemsRefundChangeRule {
	s.CancelFeeInd = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItemsRefundChangeRule) SetChangeFeeInd(v bool) *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItemsRefundChangeRule {
	s.ChangeFeeInd = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItemsRefundChangeRule) SetChangeRuleDesc(v string) *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItemsRefundChangeRule {
	s.ChangeRuleDesc = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItemsRefundChangeRule) SetChangeRuleShowColor(v string) *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItemsRefundChangeRule {
	s.ChangeRuleShowColor = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItemsRefundChangeRule) SetOfferPenaltyInfoMap(v map[string][]*ModuleReShopItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItemsRefundChangeRule {
	s.OfferPenaltyInfoMap = v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItemsRefundChangeRule) SetRefundChangeDigest(v string) *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItemsRefundChangeRule {
	s.RefundChangeDigest = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItemsRefundChangeRule) SetRefundChangeRuleDesc(v string) *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItemsRefundChangeRule {
	s.RefundChangeRuleDesc = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItemsRefundChangeRule) SetRefundRuleDesc(v string) *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItemsRefundChangeRule {
	s.RefundRuleDesc = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItemsRefundChangeRule) SetRefundRuleShowColor(v string) *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItemsRefundChangeRule {
	s.RefundRuleShowColor = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItemsRefundChangeRule) SetStructuredRefund(v bool) *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItemsRefundChangeRule {
	s.StructuredRefund = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItemsRefundChangeRule) Validate() error {
	return dara.Validate(s)
}

type IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItemsSegmentPositionList struct {
	JourneyIndex *int32 `json:"journey_index,omitempty" xml:"journey_index,omitempty"`
	SegmentIndex *int32 `json:"segment_index,omitempty" xml:"segment_index,omitempty"`
}

func (s IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItemsSegmentPositionList) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItemsSegmentPositionList) GoString() string {
	return s.String()
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItemsSegmentPositionList) GetJourneyIndex() *int32 {
	return s.JourneyIndex
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItemsSegmentPositionList) GetSegmentIndex() *int32 {
	return s.SegmentIndex
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItemsSegmentPositionList) SetJourneyIndex(v int32) *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItemsSegmentPositionList {
	s.JourneyIndex = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItemsSegmentPositionList) SetSegmentIndex(v int32) *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItemsSegmentPositionList {
	s.SegmentIndex = &v
	return s
}

func (s *IntlFlightReShopOtaSearchResponseBodyModuleReShopItemListSubItemsSegmentPositionList) Validate() error {
	return dara.Validate(s)
}
