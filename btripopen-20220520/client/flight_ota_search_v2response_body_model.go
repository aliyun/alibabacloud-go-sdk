// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightOtaSearchV2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *FlightOtaSearchV2ResponseBody
	GetCode() *string
	SetMessage(v string) *FlightOtaSearchV2ResponseBody
	GetMessage() *string
	SetModule(v *FlightOtaSearchV2ResponseBodyModule) *FlightOtaSearchV2ResponseBody
	GetModule() *FlightOtaSearchV2ResponseBodyModule
	SetRequestId(v string) *FlightOtaSearchV2ResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *FlightOtaSearchV2ResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *FlightOtaSearchV2ResponseBody
	GetTraceId() *string
}

type FlightOtaSearchV2ResponseBody struct {
	// example:
	//
	// 200
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// module
	Module *FlightOtaSearchV2ResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
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

func (s FlightOtaSearchV2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaSearchV2ResponseBody) GoString() string {
	return s.String()
}

func (s *FlightOtaSearchV2ResponseBody) GetCode() *string {
	return s.Code
}

func (s *FlightOtaSearchV2ResponseBody) GetMessage() *string {
	return s.Message
}

func (s *FlightOtaSearchV2ResponseBody) GetModule() *FlightOtaSearchV2ResponseBodyModule {
	return s.Module
}

func (s *FlightOtaSearchV2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FlightOtaSearchV2ResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *FlightOtaSearchV2ResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *FlightOtaSearchV2ResponseBody) SetCode(v string) *FlightOtaSearchV2ResponseBody {
	s.Code = &v
	return s
}

func (s *FlightOtaSearchV2ResponseBody) SetMessage(v string) *FlightOtaSearchV2ResponseBody {
	s.Message = &v
	return s
}

func (s *FlightOtaSearchV2ResponseBody) SetModule(v *FlightOtaSearchV2ResponseBodyModule) *FlightOtaSearchV2ResponseBody {
	s.Module = v
	return s
}

func (s *FlightOtaSearchV2ResponseBody) SetRequestId(v string) *FlightOtaSearchV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *FlightOtaSearchV2ResponseBody) SetSuccess(v bool) *FlightOtaSearchV2ResponseBody {
	s.Success = &v
	return s
}

func (s *FlightOtaSearchV2ResponseBody) SetTraceId(v string) *FlightOtaSearchV2ResponseBody {
	s.TraceId = &v
	return s
}

func (s *FlightOtaSearchV2ResponseBody) Validate() error {
	return dara.Validate(s)
}

type FlightOtaSearchV2ResponseBodyModule struct {
	FlightJourneyInfos []*FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfos `json:"flight_journey_infos,omitempty" xml:"flight_journey_infos,omitempty" type:"Repeated"`
	ItemList           []*FlightOtaSearchV2ResponseBodyModuleItemList           `json:"item_list,omitempty" xml:"item_list,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	SearchMode *int32 `json:"search_mode,omitempty" xml:"search_mode,omitempty"`
	// example:
	//
	// 1
	TripType *int32 `json:"trip_type,omitempty" xml:"trip_type,omitempty"`
}

func (s FlightOtaSearchV2ResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaSearchV2ResponseBodyModule) GoString() string {
	return s.String()
}

func (s *FlightOtaSearchV2ResponseBodyModule) GetFlightJourneyInfos() []*FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfos {
	return s.FlightJourneyInfos
}

func (s *FlightOtaSearchV2ResponseBodyModule) GetItemList() []*FlightOtaSearchV2ResponseBodyModuleItemList {
	return s.ItemList
}

func (s *FlightOtaSearchV2ResponseBodyModule) GetSearchMode() *int32 {
	return s.SearchMode
}

func (s *FlightOtaSearchV2ResponseBodyModule) GetTripType() *int32 {
	return s.TripType
}

func (s *FlightOtaSearchV2ResponseBodyModule) SetFlightJourneyInfos(v []*FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfos) *FlightOtaSearchV2ResponseBodyModule {
	s.FlightJourneyInfos = v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModule) SetItemList(v []*FlightOtaSearchV2ResponseBodyModuleItemList) *FlightOtaSearchV2ResponseBodyModule {
	s.ItemList = v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModule) SetSearchMode(v int32) *FlightOtaSearchV2ResponseBodyModule {
	s.SearchMode = &v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModule) SetTripType(v int32) *FlightOtaSearchV2ResponseBodyModule {
	s.TripType = &v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfos struct {
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
	Duration           *int32                                                                     `json:"duration,omitempty" xml:"duration,omitempty"`
	Extensions         map[string]*string                                                         `json:"extensions,omitempty" xml:"extensions,omitempty"`
	FlightSegmentInfos []*FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos `json:"flight_segment_infos,omitempty" xml:"flight_segment_infos,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	JourneyIndex *int32 `json:"journey_index,omitempty" xml:"journey_index,omitempty"`
	// example:
	//
	// 0
	TransferTime *int32 `json:"transfer_time,omitempty" xml:"transfer_time,omitempty"`
}

func (s FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfos) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfos) GoString() string {
	return s.String()
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfos) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfos) GetArrCityName() *string {
	return s.ArrCityName
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfos) GetArrTime() *string {
	return s.ArrTime
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfos) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfos) GetDepCityName() *string {
	return s.DepCityName
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfos) GetDepTime() *string {
	return s.DepTime
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfos) GetDuration() *int32 {
	return s.Duration
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfos) GetExtensions() map[string]*string {
	return s.Extensions
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfos) GetFlightSegmentInfos() []*FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	return s.FlightSegmentInfos
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfos) GetJourneyIndex() *int32 {
	return s.JourneyIndex
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfos) GetTransferTime() *int32 {
	return s.TransferTime
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfos) SetArrCityCode(v string) *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfos {
	s.ArrCityCode = &v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfos) SetArrCityName(v string) *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfos {
	s.ArrCityName = &v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfos) SetArrTime(v string) *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfos {
	s.ArrTime = &v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfos) SetDepCityCode(v string) *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfos {
	s.DepCityCode = &v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfos) SetDepCityName(v string) *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfos {
	s.DepCityName = &v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfos) SetDepTime(v string) *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfos {
	s.DepTime = &v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfos) SetDuration(v int32) *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfos {
	s.Duration = &v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfos) SetExtensions(v map[string]*string) *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfos {
	s.Extensions = v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfos) SetFlightSegmentInfos(v []*FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfos {
	s.FlightSegmentInfos = v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfos) SetJourneyIndex(v int32) *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfos {
	s.JourneyIndex = &v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfos) SetTransferTime(v int32) *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfos {
	s.TransferTime = &v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfos) Validate() error {
	return dara.Validate(s)
}

type FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos struct {
	AirlineInfo    *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo    `json:"airline_info,omitempty" xml:"airline_info,omitempty" type:"Struct"`
	ArrAirportInfo *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo `json:"arr_airport_info,omitempty" xml:"arr_airport_info,omitempty" type:"Struct"`
	// example:
	//
	// BJS
	ArrCityCode *string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	ArrCityName *string `json:"arr_city_name,omitempty" xml:"arr_city_name,omitempty"`
	// example:
	//
	// 2023-08-13 09:45
	ArrTime        *string                                                                                `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	BaggageDesc    *string                                                                                `json:"baggage_desc,omitempty" xml:"baggage_desc,omitempty"`
	DepAirportInfo *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo `json:"dep_airport_info,omitempty" xml:"dep_airport_info,omitempty" type:"Struct"`
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
	FlightNo        *string                                                                                 `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
	FlightShareInfo *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfo `json:"flight_share_info,omitempty" xml:"flight_share_info,omitempty" type:"Struct"`
	FlightSize      *string                                                                                 `json:"flight_size,omitempty" xml:"flight_size,omitempty"`
	FlightStopInfo  *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfo  `json:"flight_stop_info,omitempty" xml:"flight_stop_info,omitempty" type:"Struct"`
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

func (s FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GoString() string {
	return s.String()
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetAirlineInfo() *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo {
	return s.AirlineInfo
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetArrAirportInfo() *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo {
	return s.ArrAirportInfo
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetArrCityName() *string {
	return s.ArrCityName
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetArrTime() *string {
	return s.ArrTime
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetBaggageDesc() *string {
	return s.BaggageDesc
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetDepAirportInfo() *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo {
	return s.DepAirportInfo
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetDepCityName() *string {
	return s.DepCityName
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetDepTime() *string {
	return s.DepTime
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetDuration() *int32 {
	return s.Duration
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetExtraInfo() map[string]interface{} {
	return s.ExtraInfo
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetFlightNo() *string {
	return s.FlightNo
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetFlightShareInfo() *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfo {
	return s.FlightShareInfo
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetFlightSize() *string {
	return s.FlightSize
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetFlightStopInfo() *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfo {
	return s.FlightStopInfo
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetFlightType() *string {
	return s.FlightType
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetManufacturer() *string {
	return s.Manufacturer
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetMealDesc() *string {
	return s.MealDesc
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetMiles() *int32 {
	return s.Miles
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetOnTimeRate() *string {
	return s.OnTimeRate
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetOneMore() *int32 {
	return s.OneMore
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetOneMoreShow() *string {
	return s.OneMoreShow
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetSegmentIndex() *int32 {
	return s.SegmentIndex
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetShare() *bool {
	return s.Share
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetShortFlightSize() *string {
	return s.ShortFlightSize
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetStop() *bool {
	return s.Stop
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetTotalTime() *string {
	return s.TotalTime
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetTransferTime() *string {
	return s.TransferTime
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetTransferTimeNumber() *int32 {
	return s.TransferTimeNumber
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetAirlineInfo(v *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo) *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.AirlineInfo = v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetArrAirportInfo(v *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo) *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.ArrAirportInfo = v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetArrCityCode(v string) *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.ArrCityCode = &v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetArrCityName(v string) *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.ArrCityName = &v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetArrTime(v string) *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.ArrTime = &v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetBaggageDesc(v string) *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.BaggageDesc = &v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetDepAirportInfo(v *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo) *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.DepAirportInfo = v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetDepCityCode(v string) *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.DepCityCode = &v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetDepCityName(v string) *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.DepCityName = &v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetDepTime(v string) *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.DepTime = &v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetDuration(v int32) *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.Duration = &v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetExtraInfo(v map[string]interface{}) *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.ExtraInfo = v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetFlightNo(v string) *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.FlightNo = &v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetFlightShareInfo(v *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfo) *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.FlightShareInfo = v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetFlightSize(v string) *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.FlightSize = &v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetFlightStopInfo(v *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfo) *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.FlightStopInfo = v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetFlightType(v string) *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.FlightType = &v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetManufacturer(v string) *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.Manufacturer = &v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetMealDesc(v string) *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.MealDesc = &v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetMiles(v int32) *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.Miles = &v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetOnTimeRate(v string) *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.OnTimeRate = &v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetOneMore(v int32) *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.OneMore = &v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetOneMoreShow(v string) *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.OneMoreShow = &v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetSegmentIndex(v int32) *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.SegmentIndex = &v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetShare(v bool) *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.Share = &v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetShortFlightSize(v string) *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.ShortFlightSize = &v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetStop(v bool) *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.Stop = &v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetTotalTime(v string) *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.TotalTime = &v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetTransferTime(v string) *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.TransferTime = &v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetTransferTimeNumber(v int32) *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.TransferTimeNumber = &v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) Validate() error {
	return dara.Validate(s)
}

type FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo struct {
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

func (s FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo) GoString() string {
	return s.String()
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo) GetAirlineChineseName() *string {
	return s.AirlineChineseName
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo) GetAirlineChineseShortName() *string {
	return s.AirlineChineseShortName
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo) GetAirlineCode() *string {
	return s.AirlineCode
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo) GetAirlineIcon() *string {
	return s.AirlineIcon
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo) GetCheapFlight() *bool {
	return s.CheapFlight
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo) SetAirlineChineseName(v string) *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo {
	s.AirlineChineseName = &v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo) SetAirlineChineseShortName(v string) *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo {
	s.AirlineChineseShortName = &v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo) SetAirlineCode(v string) *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo {
	s.AirlineCode = &v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo) SetAirlineIcon(v string) *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo {
	s.AirlineIcon = &v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo) SetCheapFlight(v bool) *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo {
	s.CheapFlight = &v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo) Validate() error {
	return dara.Validate(s)
}

type FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo struct {
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

func (s FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo) GoString() string {
	return s.String()
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo) GetAirportCode() *string {
	return s.AirportCode
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo) GetAirportName() *string {
	return s.AirportName
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo) GetAirportNameColor() *string {
	return s.AirportNameColor
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo) GetAirportShortName() *string {
	return s.AirportShortName
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo) GetTerminal() *string {
	return s.Terminal
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo) SetAirportCode(v string) *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo {
	s.AirportCode = &v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo) SetAirportName(v string) *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo {
	s.AirportName = &v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo) SetAirportNameColor(v string) *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo {
	s.AirportNameColor = &v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo) SetAirportShortName(v string) *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo {
	s.AirportShortName = &v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo) SetTerminal(v string) *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo {
	s.Terminal = &v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo) Validate() error {
	return dara.Validate(s)
}

type FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo struct {
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

func (s FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo) GoString() string {
	return s.String()
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo) GetAirportCode() *string {
	return s.AirportCode
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo) GetAirportName() *string {
	return s.AirportName
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo) GetAirportNameColor() *string {
	return s.AirportNameColor
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo) GetAirportShortName() *string {
	return s.AirportShortName
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo) GetTerminal() *string {
	return s.Terminal
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo) SetAirportCode(v string) *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo {
	s.AirportCode = &v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo) SetAirportName(v string) *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo {
	s.AirportName = &v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo) SetAirportNameColor(v string) *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo {
	s.AirportNameColor = &v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo) SetAirportShortName(v string) *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo {
	s.AirportShortName = &v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo) SetTerminal(v string) *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo {
	s.Terminal = &v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo) Validate() error {
	return dara.Validate(s)
}

type FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfo struct {
	OperatingAirlineInfo *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo `json:"operating_airline_info,omitempty" xml:"operating_airline_info,omitempty" type:"Struct"`
	// example:
	//
	// CX601
	OperatingFlightNo *string `json:"operating_flight_no,omitempty" xml:"operating_flight_no,omitempty"`
}

func (s FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfo) GoString() string {
	return s.String()
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfo) GetOperatingAirlineInfo() *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo {
	return s.OperatingAirlineInfo
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfo) GetOperatingFlightNo() *string {
	return s.OperatingFlightNo
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfo) SetOperatingAirlineInfo(v *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfo {
	s.OperatingAirlineInfo = v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfo) SetOperatingFlightNo(v string) *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfo {
	s.OperatingFlightNo = &v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfo) Validate() error {
	return dara.Validate(s)
}

type FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo struct {
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

func (s FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) GoString() string {
	return s.String()
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) GetAirlineChineseName() *string {
	return s.AirlineChineseName
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) GetAirlineChineseShortName() *string {
	return s.AirlineChineseShortName
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) GetAirlineCode() *string {
	return s.AirlineCode
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) GetAirlineIcon() *string {
	return s.AirlineIcon
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) GetCheapFlight() *bool {
	return s.CheapFlight
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) SetAirlineChineseName(v string) *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo {
	s.AirlineChineseName = &v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) SetAirlineChineseShortName(v string) *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo {
	s.AirlineChineseShortName = &v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) SetAirlineCode(v string) *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo {
	s.AirlineCode = &v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) SetAirlineIcon(v string) *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo {
	s.AirlineIcon = &v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) SetCheapFlight(v bool) *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo {
	s.CheapFlight = &v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) Validate() error {
	return dara.Validate(s)
}

type FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfo struct {
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

func (s FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfo) GoString() string {
	return s.String()
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfo) GetStopAirport() *string {
	return s.StopAirport
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfo) GetStopAirportName() *string {
	return s.StopAirportName
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfo) GetStopArrTerm() *string {
	return s.StopArrTerm
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfo) GetStopArrTime() *string {
	return s.StopArrTime
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfo) GetStopCityCode() *string {
	return s.StopCityCode
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfo) GetStopCityName() *string {
	return s.StopCityName
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfo) GetStopCityNames() []*string {
	return s.StopCityNames
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfo) GetStopDepTerm() *string {
	return s.StopDepTerm
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfo) GetStopDepTime() *string {
	return s.StopDepTime
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfo) GetStopTime() *string {
	return s.StopTime
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfo) SetStopAirport(v string) *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfo {
	s.StopAirport = &v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfo) SetStopAirportName(v string) *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfo {
	s.StopAirportName = &v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfo) SetStopArrTerm(v string) *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfo {
	s.StopArrTerm = &v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfo) SetStopArrTime(v string) *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfo {
	s.StopArrTime = &v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfo) SetStopCityCode(v string) *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfo {
	s.StopCityCode = &v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfo) SetStopCityName(v string) *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfo {
	s.StopCityName = &v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfo) SetStopCityNames(v []*string) *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfo {
	s.StopCityNames = v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfo) SetStopDepTerm(v string) *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfo {
	s.StopDepTerm = &v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfo) SetStopDepTime(v string) *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfo {
	s.StopDepTime = &v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfo) SetStopTime(v string) *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfo {
	s.StopTime = &v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfo) Validate() error {
	return dara.Validate(s)
}

type FlightOtaSearchV2ResponseBodyModuleItemList struct {
	FlightRuleInfos map[string]*ModuleItemListFlightRuleInfosValue `json:"flight_rule_infos,omitempty" xml:"flight_rule_infos,omitempty"`
	// example:
	//
	// e50d380fc05942cc8ac57af8ae02f448_0
	ItemId             *string                                                `json:"item_id,omitempty" xml:"item_id,omitempty"`
	ShoppingItemMap    map[string]*ModuleItemListShoppingItemMapValue         `json:"shopping_item_map,omitempty" xml:"shopping_item_map,omitempty"`
	SubItemPositionMap map[string][]*ModuleItemListSubItemPositionMapValue    `json:"sub_item_position_map,omitempty" xml:"sub_item_position_map,omitempty"`
	SubItems           []*FlightOtaSearchV2ResponseBodyModuleItemListSubItems `json:"sub_items,omitempty" xml:"sub_items,omitempty" type:"Repeated"`
}

func (s FlightOtaSearchV2ResponseBodyModuleItemList) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaSearchV2ResponseBodyModuleItemList) GoString() string {
	return s.String()
}

func (s *FlightOtaSearchV2ResponseBodyModuleItemList) GetFlightRuleInfos() map[string]*ModuleItemListFlightRuleInfosValue {
	return s.FlightRuleInfos
}

func (s *FlightOtaSearchV2ResponseBodyModuleItemList) GetItemId() *string {
	return s.ItemId
}

func (s *FlightOtaSearchV2ResponseBodyModuleItemList) GetShoppingItemMap() map[string]*ModuleItemListShoppingItemMapValue {
	return s.ShoppingItemMap
}

func (s *FlightOtaSearchV2ResponseBodyModuleItemList) GetSubItemPositionMap() map[string][]*ModuleItemListSubItemPositionMapValue {
	return s.SubItemPositionMap
}

func (s *FlightOtaSearchV2ResponseBodyModuleItemList) GetSubItems() []*FlightOtaSearchV2ResponseBodyModuleItemListSubItems {
	return s.SubItems
}

func (s *FlightOtaSearchV2ResponseBodyModuleItemList) SetFlightRuleInfos(v map[string]*ModuleItemListFlightRuleInfosValue) *FlightOtaSearchV2ResponseBodyModuleItemList {
	s.FlightRuleInfos = v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleItemList) SetItemId(v string) *FlightOtaSearchV2ResponseBodyModuleItemList {
	s.ItemId = &v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleItemList) SetShoppingItemMap(v map[string]*ModuleItemListShoppingItemMapValue) *FlightOtaSearchV2ResponseBodyModuleItemList {
	s.ShoppingItemMap = v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleItemList) SetSubItemPositionMap(v map[string][]*ModuleItemListSubItemPositionMapValue) *FlightOtaSearchV2ResponseBodyModuleItemList {
	s.SubItemPositionMap = v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleItemList) SetSubItems(v []*FlightOtaSearchV2ResponseBodyModuleItemListSubItems) *FlightOtaSearchV2ResponseBodyModuleItemList {
	s.SubItems = v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleItemList) Validate() error {
	return dara.Validate(s)
}

type FlightOtaSearchV2ResponseBodyModuleItemListSubItems struct {
	ShoppingItemMap map[string]*ModuleItemListSubItemsShoppingItemMapValue `json:"shopping_item_map,omitempty" xml:"shopping_item_map,omitempty"`
	Tag             *string                                                `json:"tag,omitempty" xml:"tag,omitempty"`
	// example:
	//
	// 07df0bd9-f803-4a50-8449-f4bd675d9939
	UniqKey *string `json:"uniq_key,omitempty" xml:"uniq_key,omitempty"`
}

func (s FlightOtaSearchV2ResponseBodyModuleItemListSubItems) String() string {
	return dara.Prettify(s)
}

func (s FlightOtaSearchV2ResponseBodyModuleItemListSubItems) GoString() string {
	return s.String()
}

func (s *FlightOtaSearchV2ResponseBodyModuleItemListSubItems) GetShoppingItemMap() map[string]*ModuleItemListSubItemsShoppingItemMapValue {
	return s.ShoppingItemMap
}

func (s *FlightOtaSearchV2ResponseBodyModuleItemListSubItems) GetTag() *string {
	return s.Tag
}

func (s *FlightOtaSearchV2ResponseBodyModuleItemListSubItems) GetUniqKey() *string {
	return s.UniqKey
}

func (s *FlightOtaSearchV2ResponseBodyModuleItemListSubItems) SetShoppingItemMap(v map[string]*ModuleItemListSubItemsShoppingItemMapValue) *FlightOtaSearchV2ResponseBodyModuleItemListSubItems {
	s.ShoppingItemMap = v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleItemListSubItems) SetTag(v string) *FlightOtaSearchV2ResponseBodyModuleItemListSubItems {
	s.Tag = &v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleItemListSubItems) SetUniqKey(v string) *FlightOtaSearchV2ResponseBodyModuleItemListSubItems {
	s.UniqKey = &v
	return s
}

func (s *FlightOtaSearchV2ResponseBodyModuleItemListSubItems) Validate() error {
	return dara.Validate(s)
}
