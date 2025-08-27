// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightOtaItemDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *IntlFlightOtaItemDetailResponseBody
	GetCode() *string
	SetMessage(v string) *IntlFlightOtaItemDetailResponseBody
	GetMessage() *string
	SetModule(v *IntlFlightOtaItemDetailResponseBodyModule) *IntlFlightOtaItemDetailResponseBody
	GetModule() *IntlFlightOtaItemDetailResponseBodyModule
	SetRequestId(v string) *IntlFlightOtaItemDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *IntlFlightOtaItemDetailResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *IntlFlightOtaItemDetailResponseBody
	GetTraceId() *string
}

type IntlFlightOtaItemDetailResponseBody struct {
	// example:
	//
	// SUCCESS
	Code    *string                                    `json:"code,omitempty" xml:"code,omitempty"`
	Message *string                                    `json:"message,omitempty" xml:"message,omitempty"`
	Module  *IntlFlightOtaItemDetailResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
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

func (s IntlFlightOtaItemDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOtaItemDetailResponseBody) GoString() string {
	return s.String()
}

func (s *IntlFlightOtaItemDetailResponseBody) GetCode() *string {
	return s.Code
}

func (s *IntlFlightOtaItemDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *IntlFlightOtaItemDetailResponseBody) GetModule() *IntlFlightOtaItemDetailResponseBodyModule {
	return s.Module
}

func (s *IntlFlightOtaItemDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *IntlFlightOtaItemDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *IntlFlightOtaItemDetailResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *IntlFlightOtaItemDetailResponseBody) SetCode(v string) *IntlFlightOtaItemDetailResponseBody {
	s.Code = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBody) SetMessage(v string) *IntlFlightOtaItemDetailResponseBody {
	s.Message = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBody) SetModule(v *IntlFlightOtaItemDetailResponseBodyModule) *IntlFlightOtaItemDetailResponseBody {
	s.Module = v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBody) SetRequestId(v string) *IntlFlightOtaItemDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBody) SetSuccess(v bool) *IntlFlightOtaItemDetailResponseBody {
	s.Success = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBody) SetTraceId(v string) *IntlFlightOtaItemDetailResponseBody {
	s.TraceId = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type IntlFlightOtaItemDetailResponseBodyModule struct {
	FlightJourneyInfos []*IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfos `json:"flight_journey_infos,omitempty" xml:"flight_journey_infos,omitempty" type:"Repeated"`
	GroupItem          *IntlFlightOtaItemDetailResponseBodyModuleGroupItem            `json:"group_item,omitempty" xml:"group_item,omitempty" type:"Struct"`
	ShutterDocs        []*IntlFlightOtaItemDetailResponseBodyModuleShutterDocs        `json:"shutter_docs,omitempty" xml:"shutter_docs,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	TripType *int32 `json:"trip_type,omitempty" xml:"trip_type,omitempty"`
}

func (s IntlFlightOtaItemDetailResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOtaItemDetailResponseBodyModule) GoString() string {
	return s.String()
}

func (s *IntlFlightOtaItemDetailResponseBodyModule) GetFlightJourneyInfos() []*IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfos {
	return s.FlightJourneyInfos
}

func (s *IntlFlightOtaItemDetailResponseBodyModule) GetGroupItem() *IntlFlightOtaItemDetailResponseBodyModuleGroupItem {
	return s.GroupItem
}

func (s *IntlFlightOtaItemDetailResponseBodyModule) GetShutterDocs() []*IntlFlightOtaItemDetailResponseBodyModuleShutterDocs {
	return s.ShutterDocs
}

func (s *IntlFlightOtaItemDetailResponseBodyModule) GetTripType() *int32 {
	return s.TripType
}

func (s *IntlFlightOtaItemDetailResponseBodyModule) SetFlightJourneyInfos(v []*IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfos) *IntlFlightOtaItemDetailResponseBodyModule {
	s.FlightJourneyInfos = v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModule) SetGroupItem(v *IntlFlightOtaItemDetailResponseBodyModuleGroupItem) *IntlFlightOtaItemDetailResponseBodyModule {
	s.GroupItem = v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModule) SetShutterDocs(v []*IntlFlightOtaItemDetailResponseBodyModuleShutterDocs) *IntlFlightOtaItemDetailResponseBodyModule {
	s.ShutterDocs = v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModule) SetTripType(v int32) *IntlFlightOtaItemDetailResponseBodyModule {
	s.TripType = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfos struct {
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
	Duration           *int32                                                                           `json:"duration,omitempty" xml:"duration,omitempty"`
	Extensions         map[string]*string                                                               `json:"extensions,omitempty" xml:"extensions,omitempty"`
	FlightSegmentInfos []*IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos `json:"flight_segment_infos,omitempty" xml:"flight_segment_infos,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	JourneyIndex *int32 `json:"journey_index,omitempty" xml:"journey_index,omitempty"`
	// example:
	//
	// 0
	TransferTime *int32 `json:"transfer_time,omitempty" xml:"transfer_time,omitempty"`
}

func (s IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfos) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfos) GoString() string {
	return s.String()
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfos) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfos) GetArrCityName() *string {
	return s.ArrCityName
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfos) GetArrTime() *string {
	return s.ArrTime
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfos) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfos) GetDepCityName() *string {
	return s.DepCityName
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfos) GetDepTime() *string {
	return s.DepTime
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfos) GetDuration() *int32 {
	return s.Duration
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfos) GetExtensions() map[string]*string {
	return s.Extensions
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfos) GetFlightSegmentInfos() []*IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	return s.FlightSegmentInfos
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfos) GetJourneyIndex() *int32 {
	return s.JourneyIndex
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfos) GetTransferTime() *int32 {
	return s.TransferTime
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfos) SetArrCityCode(v string) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfos {
	s.ArrCityCode = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfos) SetArrCityName(v string) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfos {
	s.ArrCityName = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfos) SetArrTime(v string) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfos {
	s.ArrTime = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfos) SetDepCityCode(v string) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfos {
	s.DepCityCode = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfos) SetDepCityName(v string) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfos {
	s.DepCityName = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfos) SetDepTime(v string) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfos {
	s.DepTime = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfos) SetDuration(v int32) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfos {
	s.Duration = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfos) SetExtensions(v map[string]*string) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfos {
	s.Extensions = v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfos) SetFlightSegmentInfos(v []*IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfos {
	s.FlightSegmentInfos = v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfos) SetJourneyIndex(v int32) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfos {
	s.JourneyIndex = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfos) SetTransferTime(v int32) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfos {
	s.TransferTime = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfos) Validate() error {
	return dara.Validate(s)
}

type IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos struct {
	AirlineInfo    *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo    `json:"airline_info,omitempty" xml:"airline_info,omitempty" type:"Struct"`
	ArrAirportInfo *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo `json:"arr_airport_info,omitempty" xml:"arr_airport_info,omitempty" type:"Struct"`
	// example:
	//
	// BJS
	ArrCityCode *string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	ArrCityName *string `json:"arr_city_name,omitempty" xml:"arr_city_name,omitempty"`
	// example:
	//
	// 2023-08-13 09:45
	ArrTime        *string                                                                                      `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	BaggageDesc    *string                                                                                      `json:"baggage_desc,omitempty" xml:"baggage_desc,omitempty"`
	DepAirportInfo *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo `json:"dep_airport_info,omitempty" xml:"dep_airport_info,omitempty" type:"Struct"`
	// example:
	//
	// HGH
	DepCityCode *string `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	DepCityName *string `json:"dep_city_name,omitempty" xml:"dep_city_name,omitempty"`
	// example:
	//
	// 2023-08-13 07:25
	DepTime *string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	// example:
	//
	// 140
	Duration *int32 `json:"duration,omitempty" xml:"duration,omitempty"`
	// example:
	//
	// MU5131
	FlightNo        *string                                                                                       `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
	FlightShareInfo *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfo `json:"flight_share_info,omitempty" xml:"flight_share_info,omitempty" type:"Struct"`
	FlightSize      *string                                                                                       `json:"flight_size,omitempty" xml:"flight_size,omitempty"`
	FlightStopInfo  *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfo  `json:"flight_stop_info,omitempty" xml:"flight_stop_info,omitempty" type:"Struct"`
	// example:
	//
	// 320
	FlightType        *string                                                                                         `json:"flight_type,omitempty" xml:"flight_type,omitempty"`
	LuggageDirectInfo *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosLuggageDirectInfo `json:"luggage_direct_info,omitempty" xml:"luggage_direct_info,omitempty" type:"Struct"`
	Manufacturer      *string                                                                                         `json:"manufacturer,omitempty" xml:"manufacturer,omitempty"`
	MealDesc          *string                                                                                         `json:"meal_desc,omitempty" xml:"meal_desc,omitempty"`
	// example:
	//
	// 80%
	OnTimeRate *string `json:"on_time_rate,omitempty" xml:"on_time_rate,omitempty"`
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
	SegmentKey        *string                                                                                         `json:"segment_key,omitempty" xml:"segment_key,omitempty"`
	SegmentVisaRemark *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark `json:"segment_visa_remark,omitempty" xml:"segment_visa_remark,omitempty" type:"Struct"`
	// example:
	//
	// false
	Share           *bool   `json:"share,omitempty" xml:"share,omitempty"`
	ShortFlightSize *string `json:"short_flight_size,omitempty" xml:"short_flight_size,omitempty"`
	// example:
	//
	// false
	Stop         *bool   `json:"stop,omitempty" xml:"stop,omitempty"`
	TotalTime    *string `json:"total_time,omitempty" xml:"total_time,omitempty"`
	TransferTime *string `json:"transfer_time,omitempty" xml:"transfer_time,omitempty"`
	// example:
	//
	// 80
	TransferTimeNumber *int32 `json:"transfer_time_number,omitempty" xml:"transfer_time_number,omitempty"`
}

func (s IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GoString() string {
	return s.String()
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetAirlineInfo() *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo {
	return s.AirlineInfo
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetArrAirportInfo() *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo {
	return s.ArrAirportInfo
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetArrCityName() *string {
	return s.ArrCityName
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetArrTime() *string {
	return s.ArrTime
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetBaggageDesc() *string {
	return s.BaggageDesc
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetDepAirportInfo() *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo {
	return s.DepAirportInfo
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetDepCityName() *string {
	return s.DepCityName
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetDepTime() *string {
	return s.DepTime
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetDuration() *int32 {
	return s.Duration
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetFlightNo() *string {
	return s.FlightNo
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetFlightShareInfo() *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfo {
	return s.FlightShareInfo
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetFlightSize() *string {
	return s.FlightSize
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetFlightStopInfo() *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfo {
	return s.FlightStopInfo
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetFlightType() *string {
	return s.FlightType
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetLuggageDirectInfo() *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosLuggageDirectInfo {
	return s.LuggageDirectInfo
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetManufacturer() *string {
	return s.Manufacturer
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetMealDesc() *string {
	return s.MealDesc
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetOnTimeRate() *string {
	return s.OnTimeRate
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetOneMore() *int32 {
	return s.OneMore
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetOneMoreShow() *string {
	return s.OneMoreShow
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetSegmentIndex() *int32 {
	return s.SegmentIndex
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetSegmentKey() *string {
	return s.SegmentKey
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetSegmentVisaRemark() *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark {
	return s.SegmentVisaRemark
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetShare() *bool {
	return s.Share
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetShortFlightSize() *string {
	return s.ShortFlightSize
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetStop() *bool {
	return s.Stop
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetTotalTime() *string {
	return s.TotalTime
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetTransferTime() *string {
	return s.TransferTime
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetTransferTimeNumber() *int32 {
	return s.TransferTimeNumber
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetAirlineInfo(v *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.AirlineInfo = v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetArrAirportInfo(v *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.ArrAirportInfo = v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetArrCityCode(v string) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.ArrCityCode = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetArrCityName(v string) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.ArrCityName = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetArrTime(v string) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.ArrTime = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetBaggageDesc(v string) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.BaggageDesc = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetDepAirportInfo(v *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.DepAirportInfo = v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetDepCityCode(v string) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.DepCityCode = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetDepCityName(v string) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.DepCityName = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetDepTime(v string) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.DepTime = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetDuration(v int32) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.Duration = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetFlightNo(v string) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.FlightNo = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetFlightShareInfo(v *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfo) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.FlightShareInfo = v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetFlightSize(v string) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.FlightSize = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetFlightStopInfo(v *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfo) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.FlightStopInfo = v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetFlightType(v string) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.FlightType = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetLuggageDirectInfo(v *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosLuggageDirectInfo) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.LuggageDirectInfo = v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetManufacturer(v string) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.Manufacturer = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetMealDesc(v string) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.MealDesc = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetOnTimeRate(v string) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.OnTimeRate = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetOneMore(v int32) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.OneMore = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetOneMoreShow(v string) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.OneMoreShow = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetSegmentIndex(v int32) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.SegmentIndex = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetSegmentKey(v string) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.SegmentKey = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetSegmentVisaRemark(v *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.SegmentVisaRemark = v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetShare(v bool) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.Share = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetShortFlightSize(v string) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.ShortFlightSize = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetStop(v bool) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.Stop = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetTotalTime(v string) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.TotalTime = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetTransferTime(v string) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.TransferTime = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetTransferTimeNumber(v int32) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.TransferTimeNumber = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) Validate() error {
	return dara.Validate(s)
}

type IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo struct {
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

func (s IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo) GoString() string {
	return s.String()
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo) GetAirlineChineseName() *string {
	return s.AirlineChineseName
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo) GetAirlineChineseShortName() *string {
	return s.AirlineChineseShortName
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo) GetAirlineCode() *string {
	return s.AirlineCode
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo) GetAirlineIcon() *string {
	return s.AirlineIcon
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo) GetCheapFlight() *bool {
	return s.CheapFlight
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo) SetAirlineChineseName(v string) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo {
	s.AirlineChineseName = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo) SetAirlineChineseShortName(v string) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo {
	s.AirlineChineseShortName = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo) SetAirlineCode(v string) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo {
	s.AirlineCode = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo) SetAirlineIcon(v string) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo {
	s.AirlineIcon = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo) SetCheapFlight(v bool) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo {
	s.CheapFlight = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo) Validate() error {
	return dara.Validate(s)
}

type IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo struct {
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

func (s IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo) GoString() string {
	return s.String()
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo) GetAirportCode() *string {
	return s.AirportCode
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo) GetAirportName() *string {
	return s.AirportName
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo) GetAirportShortName() *string {
	return s.AirportShortName
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo) GetTerminal() *string {
	return s.Terminal
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo) SetAirportCode(v string) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo {
	s.AirportCode = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo) SetAirportName(v string) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo {
	s.AirportName = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo) SetAirportShortName(v string) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo {
	s.AirportShortName = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo) SetTerminal(v string) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo {
	s.Terminal = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo) Validate() error {
	return dara.Validate(s)
}

type IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo struct {
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

func (s IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo) GoString() string {
	return s.String()
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo) GetAirportCode() *string {
	return s.AirportCode
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo) GetAirportName() *string {
	return s.AirportName
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo) GetAirportShortName() *string {
	return s.AirportShortName
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo) GetTerminal() *string {
	return s.Terminal
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo) SetAirportCode(v string) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo {
	s.AirportCode = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo) SetAirportName(v string) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo {
	s.AirportName = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo) SetAirportShortName(v string) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo {
	s.AirportShortName = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo) SetTerminal(v string) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo {
	s.Terminal = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo) Validate() error {
	return dara.Validate(s)
}

type IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfo struct {
	OperatingAirlineInfo *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo `json:"operating_airline_info,omitempty" xml:"operating_airline_info,omitempty" type:"Struct"`
	// example:
	//
	// CX601
	OperatingFlightNo *string `json:"operating_flight_no,omitempty" xml:"operating_flight_no,omitempty"`
}

func (s IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfo) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfo) GoString() string {
	return s.String()
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfo) GetOperatingAirlineInfo() *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo {
	return s.OperatingAirlineInfo
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfo) GetOperatingFlightNo() *string {
	return s.OperatingFlightNo
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfo) SetOperatingAirlineInfo(v *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfo {
	s.OperatingAirlineInfo = v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfo) SetOperatingFlightNo(v string) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfo {
	s.OperatingFlightNo = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfo) Validate() error {
	return dara.Validate(s)
}

type IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo struct {
	AirlineChineseName      *string `json:"airline_chinese_name,omitempty" xml:"airline_chinese_name,omitempty"`
	AirlineChineseShortName *string `json:"airline_chinese_short_name,omitempty" xml:"airline_chinese_short_name,omitempty"`
	// example:
	//
	// KN
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

func (s IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) GoString() string {
	return s.String()
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) GetAirlineChineseName() *string {
	return s.AirlineChineseName
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) GetAirlineChineseShortName() *string {
	return s.AirlineChineseShortName
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) GetAirlineCode() *string {
	return s.AirlineCode
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) GetAirlineIcon() *string {
	return s.AirlineIcon
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) GetCheapFlight() *bool {
	return s.CheapFlight
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) SetAirlineChineseName(v string) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo {
	s.AirlineChineseName = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) SetAirlineChineseShortName(v string) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo {
	s.AirlineChineseShortName = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) SetAirlineCode(v string) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo {
	s.AirlineCode = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) SetAirlineIcon(v string) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo {
	s.AirlineIcon = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) SetCheapFlight(v bool) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo {
	s.CheapFlight = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) Validate() error {
	return dara.Validate(s)
}

type IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfo struct {
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

func (s IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfo) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfo) GoString() string {
	return s.String()
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfo) GetStopAirport() *string {
	return s.StopAirport
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfo) GetStopAirportName() *string {
	return s.StopAirportName
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfo) GetStopArrTerm() *string {
	return s.StopArrTerm
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfo) GetStopArrTime() *string {
	return s.StopArrTime
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfo) GetStopCityCode() *string {
	return s.StopCityCode
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfo) GetStopCityName() *string {
	return s.StopCityName
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfo) GetStopCityNames() []*string {
	return s.StopCityNames
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfo) GetStopDepTerm() *string {
	return s.StopDepTerm
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfo) GetStopDepTime() *string {
	return s.StopDepTime
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfo) GetStopTime() *string {
	return s.StopTime
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfo) SetStopAirport(v string) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfo {
	s.StopAirport = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfo) SetStopAirportName(v string) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfo {
	s.StopAirportName = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfo) SetStopArrTerm(v string) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfo {
	s.StopArrTerm = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfo) SetStopArrTime(v string) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfo {
	s.StopArrTime = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfo) SetStopCityCode(v string) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfo {
	s.StopCityCode = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfo) SetStopCityName(v string) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfo {
	s.StopCityName = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfo) SetStopCityNames(v []*string) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfo {
	s.StopCityNames = v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfo) SetStopDepTerm(v string) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfo {
	s.StopDepTerm = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfo) SetStopDepTime(v string) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfo {
	s.StopDepTime = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfo) SetStopTime(v string) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfo {
	s.StopTime = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfo) Validate() error {
	return dara.Validate(s)
}

type IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosLuggageDirectInfo struct {
	// example:
	//
	// 1
	DepCityLuggageDirect *int32 `json:"dep_city_luggage_direct,omitempty" xml:"dep_city_luggage_direct,omitempty"`
	// example:
	//
	// 0
	StopCityLuggageDirect *int32 `json:"stop_city_luggage_direct,omitempty" xml:"stop_city_luggage_direct,omitempty"`
}

func (s IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosLuggageDirectInfo) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosLuggageDirectInfo) GoString() string {
	return s.String()
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosLuggageDirectInfo) GetDepCityLuggageDirect() *int32 {
	return s.DepCityLuggageDirect
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosLuggageDirectInfo) GetStopCityLuggageDirect() *int32 {
	return s.StopCityLuggageDirect
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosLuggageDirectInfo) SetDepCityLuggageDirect(v int32) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosLuggageDirectInfo {
	s.DepCityLuggageDirect = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosLuggageDirectInfo) SetStopCityLuggageDirect(v int32) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosLuggageDirectInfo {
	s.StopCityLuggageDirect = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosLuggageDirectInfo) Validate() error {
	return dara.Validate(s)
}

type IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark struct {
	DepCityVisaRemark *string `json:"dep_city_visa_remark,omitempty" xml:"dep_city_visa_remark,omitempty"`
	// example:
	//
	// 1
	DepCityVisaType     *int32    `json:"dep_city_visa_type,omitempty" xml:"dep_city_visa_type,omitempty"`
	StopCityVisaRemarks []*string `json:"stop_city_visa_remarks,omitempty" xml:"stop_city_visa_remarks,omitempty" type:"Repeated"`
	StopCityVisaTypes   []*int32  `json:"stop_city_visa_types,omitempty" xml:"stop_city_visa_types,omitempty" type:"Repeated"`
}

func (s IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark) GoString() string {
	return s.String()
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark) GetDepCityVisaRemark() *string {
	return s.DepCityVisaRemark
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark) GetDepCityVisaType() *int32 {
	return s.DepCityVisaType
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark) GetStopCityVisaRemarks() []*string {
	return s.StopCityVisaRemarks
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark) GetStopCityVisaTypes() []*int32 {
	return s.StopCityVisaTypes
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark) SetDepCityVisaRemark(v string) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark {
	s.DepCityVisaRemark = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark) SetDepCityVisaType(v int32) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark {
	s.DepCityVisaType = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark) SetStopCityVisaRemarks(v []*string) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark {
	s.StopCityVisaRemarks = v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark) SetStopCityVisaTypes(v []*int32) *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark {
	s.StopCityVisaTypes = v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark) Validate() error {
	return dara.Validate(s)
}

type IntlFlightOtaItemDetailResponseBodyModuleGroupItem struct {
	AgreementPriceCodes []*string                                                               `json:"agreement_price_codes,omitempty" xml:"agreement_price_codes,omitempty" type:"Repeated"`
	FlightRuleInfoList  []*IntlFlightOtaItemDetailResponseBodyModuleGroupItemFlightRuleInfoList `json:"flight_rule_info_list,omitempty" xml:"flight_rule_info_list,omitempty" type:"Repeated"`
	// example:
	//
	// 669845158418
	ItemId             *string                                                       `json:"item_id,omitempty" xml:"item_id,omitempty"`
	ItemType           *string                                                       `json:"item_type,omitempty" xml:"item_type,omitempty"`
	ShoppingItemMap    map[string]*ModuleGroupItemShoppingItemMapValue               `json:"shopping_item_map,omitempty" xml:"shopping_item_map,omitempty"`
	SubItemPositionMap map[string][]*ModuleGroupItemSubItemPositionMapValue          `json:"sub_item_position_map,omitempty" xml:"sub_item_position_map,omitempty"`
	SubItems           []*IntlFlightOtaItemDetailResponseBodyModuleGroupItemSubItems `json:"sub_items,omitempty" xml:"sub_items,omitempty" type:"Repeated"`
}

func (s IntlFlightOtaItemDetailResponseBodyModuleGroupItem) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOtaItemDetailResponseBodyModuleGroupItem) GoString() string {
	return s.String()
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleGroupItem) GetAgreementPriceCodes() []*string {
	return s.AgreementPriceCodes
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleGroupItem) GetFlightRuleInfoList() []*IntlFlightOtaItemDetailResponseBodyModuleGroupItemFlightRuleInfoList {
	return s.FlightRuleInfoList
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleGroupItem) GetItemId() *string {
	return s.ItemId
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleGroupItem) GetItemType() *string {
	return s.ItemType
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleGroupItem) GetShoppingItemMap() map[string]*ModuleGroupItemShoppingItemMapValue {
	return s.ShoppingItemMap
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleGroupItem) GetSubItemPositionMap() map[string][]*ModuleGroupItemSubItemPositionMapValue {
	return s.SubItemPositionMap
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleGroupItem) GetSubItems() []*IntlFlightOtaItemDetailResponseBodyModuleGroupItemSubItems {
	return s.SubItems
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleGroupItem) SetAgreementPriceCodes(v []*string) *IntlFlightOtaItemDetailResponseBodyModuleGroupItem {
	s.AgreementPriceCodes = v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleGroupItem) SetFlightRuleInfoList(v []*IntlFlightOtaItemDetailResponseBodyModuleGroupItemFlightRuleInfoList) *IntlFlightOtaItemDetailResponseBodyModuleGroupItem {
	s.FlightRuleInfoList = v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleGroupItem) SetItemId(v string) *IntlFlightOtaItemDetailResponseBodyModuleGroupItem {
	s.ItemId = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleGroupItem) SetItemType(v string) *IntlFlightOtaItemDetailResponseBodyModuleGroupItem {
	s.ItemType = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleGroupItem) SetShoppingItemMap(v map[string]*ModuleGroupItemShoppingItemMapValue) *IntlFlightOtaItemDetailResponseBodyModuleGroupItem {
	s.ShoppingItemMap = v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleGroupItem) SetSubItemPositionMap(v map[string][]*ModuleGroupItemSubItemPositionMapValue) *IntlFlightOtaItemDetailResponseBodyModuleGroupItem {
	s.SubItemPositionMap = v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleGroupItem) SetSubItems(v []*IntlFlightOtaItemDetailResponseBodyModuleGroupItemSubItems) *IntlFlightOtaItemDetailResponseBodyModuleGroupItem {
	s.SubItems = v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleGroupItem) Validate() error {
	return dara.Validate(s)
}

type IntlFlightOtaItemDetailResponseBodyModuleGroupItemFlightRuleInfoList struct {
	FlightRuleInfo  *IntlFlightOtaItemDetailResponseBodyModuleGroupItemFlightRuleInfoListFlightRuleInfo  `json:"flight_rule_info,omitempty" xml:"flight_rule_info,omitempty" type:"Struct"`
	SegmentPosition *IntlFlightOtaItemDetailResponseBodyModuleGroupItemFlightRuleInfoListSegmentPosition `json:"segment_position,omitempty" xml:"segment_position,omitempty" type:"Struct"`
}

func (s IntlFlightOtaItemDetailResponseBodyModuleGroupItemFlightRuleInfoList) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOtaItemDetailResponseBodyModuleGroupItemFlightRuleInfoList) GoString() string {
	return s.String()
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleGroupItemFlightRuleInfoList) GetFlightRuleInfo() *IntlFlightOtaItemDetailResponseBodyModuleGroupItemFlightRuleInfoListFlightRuleInfo {
	return s.FlightRuleInfo
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleGroupItemFlightRuleInfoList) GetSegmentPosition() *IntlFlightOtaItemDetailResponseBodyModuleGroupItemFlightRuleInfoListSegmentPosition {
	return s.SegmentPosition
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleGroupItemFlightRuleInfoList) SetFlightRuleInfo(v *IntlFlightOtaItemDetailResponseBodyModuleGroupItemFlightRuleInfoListFlightRuleInfo) *IntlFlightOtaItemDetailResponseBodyModuleGroupItemFlightRuleInfoList {
	s.FlightRuleInfo = v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleGroupItemFlightRuleInfoList) SetSegmentPosition(v *IntlFlightOtaItemDetailResponseBodyModuleGroupItemFlightRuleInfoListSegmentPosition) *IntlFlightOtaItemDetailResponseBodyModuleGroupItemFlightRuleInfoList {
	s.SegmentPosition = v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleGroupItemFlightRuleInfoList) Validate() error {
	return dara.Validate(s)
}

type IntlFlightOtaItemDetailResponseBodyModuleGroupItemFlightRuleInfoListFlightRuleInfo struct {
	BaggageDesc          *string `json:"baggage_desc,omitempty" xml:"baggage_desc,omitempty"`
	RefundChangeRuleDesc *string `json:"refund_change_rule_desc,omitempty" xml:"refund_change_rule_desc,omitempty"`
}

func (s IntlFlightOtaItemDetailResponseBodyModuleGroupItemFlightRuleInfoListFlightRuleInfo) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOtaItemDetailResponseBodyModuleGroupItemFlightRuleInfoListFlightRuleInfo) GoString() string {
	return s.String()
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleGroupItemFlightRuleInfoListFlightRuleInfo) GetBaggageDesc() *string {
	return s.BaggageDesc
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleGroupItemFlightRuleInfoListFlightRuleInfo) GetRefundChangeRuleDesc() *string {
	return s.RefundChangeRuleDesc
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleGroupItemFlightRuleInfoListFlightRuleInfo) SetBaggageDesc(v string) *IntlFlightOtaItemDetailResponseBodyModuleGroupItemFlightRuleInfoListFlightRuleInfo {
	s.BaggageDesc = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleGroupItemFlightRuleInfoListFlightRuleInfo) SetRefundChangeRuleDesc(v string) *IntlFlightOtaItemDetailResponseBodyModuleGroupItemFlightRuleInfoListFlightRuleInfo {
	s.RefundChangeRuleDesc = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleGroupItemFlightRuleInfoListFlightRuleInfo) Validate() error {
	return dara.Validate(s)
}

type IntlFlightOtaItemDetailResponseBodyModuleGroupItemFlightRuleInfoListSegmentPosition struct {
	JourneyIndex *int32 `json:"journey_index,omitempty" xml:"journey_index,omitempty"`
	SegmentIndex *int32 `json:"segment_index,omitempty" xml:"segment_index,omitempty"`
}

func (s IntlFlightOtaItemDetailResponseBodyModuleGroupItemFlightRuleInfoListSegmentPosition) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOtaItemDetailResponseBodyModuleGroupItemFlightRuleInfoListSegmentPosition) GoString() string {
	return s.String()
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleGroupItemFlightRuleInfoListSegmentPosition) GetJourneyIndex() *int32 {
	return s.JourneyIndex
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleGroupItemFlightRuleInfoListSegmentPosition) GetSegmentIndex() *int32 {
	return s.SegmentIndex
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleGroupItemFlightRuleInfoListSegmentPosition) SetJourneyIndex(v int32) *IntlFlightOtaItemDetailResponseBodyModuleGroupItemFlightRuleInfoListSegmentPosition {
	s.JourneyIndex = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleGroupItemFlightRuleInfoListSegmentPosition) SetSegmentIndex(v int32) *IntlFlightOtaItemDetailResponseBodyModuleGroupItemFlightRuleInfoListSegmentPosition {
	s.SegmentIndex = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleGroupItemFlightRuleInfoListSegmentPosition) Validate() error {
	return dara.Validate(s)
}

type IntlFlightOtaItemDetailResponseBodyModuleGroupItemSubItems struct {
	BaggageRule      *IntlFlightOtaItemDetailResponseBodyModuleGroupItemSubItemsBaggageRule      `json:"baggage_rule,omitempty" xml:"baggage_rule,omitempty" type:"Struct"`
	RefundChangeRule *IntlFlightOtaItemDetailResponseBodyModuleGroupItemSubItemsRefundChangeRule `json:"refund_change_rule,omitempty" xml:"refund_change_rule,omitempty" type:"Struct"`
	SegmentKeys      []*string                                                                   `json:"segment_keys,omitempty" xml:"segment_keys,omitempty" type:"Repeated"`
	ShoppingItemMap  map[string]*ModuleGroupItemSubItemsShoppingItemMapValue                     `json:"shopping_item_map,omitempty" xml:"shopping_item_map,omitempty"`
	// example:
	//
	// 23412341234124
	UniqKey *string `json:"uniq_key,omitempty" xml:"uniq_key,omitempty"`
}

func (s IntlFlightOtaItemDetailResponseBodyModuleGroupItemSubItems) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOtaItemDetailResponseBodyModuleGroupItemSubItems) GoString() string {
	return s.String()
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleGroupItemSubItems) GetBaggageRule() *IntlFlightOtaItemDetailResponseBodyModuleGroupItemSubItemsBaggageRule {
	return s.BaggageRule
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleGroupItemSubItems) GetRefundChangeRule() *IntlFlightOtaItemDetailResponseBodyModuleGroupItemSubItemsRefundChangeRule {
	return s.RefundChangeRule
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleGroupItemSubItems) GetSegmentKeys() []*string {
	return s.SegmentKeys
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleGroupItemSubItems) GetShoppingItemMap() map[string]*ModuleGroupItemSubItemsShoppingItemMapValue {
	return s.ShoppingItemMap
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleGroupItemSubItems) GetUniqKey() *string {
	return s.UniqKey
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleGroupItemSubItems) SetBaggageRule(v *IntlFlightOtaItemDetailResponseBodyModuleGroupItemSubItemsBaggageRule) *IntlFlightOtaItemDetailResponseBodyModuleGroupItemSubItems {
	s.BaggageRule = v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleGroupItemSubItems) SetRefundChangeRule(v *IntlFlightOtaItemDetailResponseBodyModuleGroupItemSubItemsRefundChangeRule) *IntlFlightOtaItemDetailResponseBodyModuleGroupItemSubItems {
	s.RefundChangeRule = v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleGroupItemSubItems) SetSegmentKeys(v []*string) *IntlFlightOtaItemDetailResponseBodyModuleGroupItemSubItems {
	s.SegmentKeys = v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleGroupItemSubItems) SetShoppingItemMap(v map[string]*ModuleGroupItemSubItemsShoppingItemMapValue) *IntlFlightOtaItemDetailResponseBodyModuleGroupItemSubItems {
	s.ShoppingItemMap = v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleGroupItemSubItems) SetUniqKey(v string) *IntlFlightOtaItemDetailResponseBodyModuleGroupItemSubItems {
	s.UniqKey = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleGroupItemSubItems) Validate() error {
	return dara.Validate(s)
}

type IntlFlightOtaItemDetailResponseBodyModuleGroupItemSubItemsBaggageRule struct {
	BaggageDigest  *string                                                             `json:"baggage_digest,omitempty" xml:"baggage_digest,omitempty"`
	BaggageInfoMap map[string][]*ModuleGroupItemSubItemsBaggageRuleBaggageInfoMapValue `json:"baggage_info_map,omitempty" xml:"baggage_info_map,omitempty"`
	// example:
	//
	// true
	StructuredBaggage *bool `json:"structured_baggage,omitempty" xml:"structured_baggage,omitempty"`
}

func (s IntlFlightOtaItemDetailResponseBodyModuleGroupItemSubItemsBaggageRule) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOtaItemDetailResponseBodyModuleGroupItemSubItemsBaggageRule) GoString() string {
	return s.String()
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleGroupItemSubItemsBaggageRule) GetBaggageDigest() *string {
	return s.BaggageDigest
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleGroupItemSubItemsBaggageRule) GetBaggageInfoMap() map[string][]*ModuleGroupItemSubItemsBaggageRuleBaggageInfoMapValue {
	return s.BaggageInfoMap
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleGroupItemSubItemsBaggageRule) GetStructuredBaggage() *bool {
	return s.StructuredBaggage
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleGroupItemSubItemsBaggageRule) SetBaggageDigest(v string) *IntlFlightOtaItemDetailResponseBodyModuleGroupItemSubItemsBaggageRule {
	s.BaggageDigest = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleGroupItemSubItemsBaggageRule) SetBaggageInfoMap(v map[string][]*ModuleGroupItemSubItemsBaggageRuleBaggageInfoMapValue) *IntlFlightOtaItemDetailResponseBodyModuleGroupItemSubItemsBaggageRule {
	s.BaggageInfoMap = v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleGroupItemSubItemsBaggageRule) SetStructuredBaggage(v bool) *IntlFlightOtaItemDetailResponseBodyModuleGroupItemSubItemsBaggageRule {
	s.StructuredBaggage = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleGroupItemSubItemsBaggageRule) Validate() error {
	return dara.Validate(s)
}

type IntlFlightOtaItemDetailResponseBodyModuleGroupItemSubItemsRefundChangeRule struct {
	// example:
	//
	// false
	CancelFeeInd *bool `json:"cancel_fee_ind,omitempty" xml:"cancel_fee_ind,omitempty"`
	// example:
	//
	// false
	ChangeFeeInd        *bool                                                                         `json:"change_fee_ind,omitempty" xml:"change_fee_ind,omitempty"`
	OfferPenaltyInfoMap map[string][]*ModuleGroupItemSubItemsRefundChangeRuleOfferPenaltyInfoMapValue `json:"offer_penalty_info_map,omitempty" xml:"offer_penalty_info_map,omitempty"`
	RefundChangeDigest  *string                                                                       `json:"refund_change_digest,omitempty" xml:"refund_change_digest,omitempty"`
	// example:
	//
	// false
	StructuredRefund *bool `json:"structured_refund,omitempty" xml:"structured_refund,omitempty"`
}

func (s IntlFlightOtaItemDetailResponseBodyModuleGroupItemSubItemsRefundChangeRule) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOtaItemDetailResponseBodyModuleGroupItemSubItemsRefundChangeRule) GoString() string {
	return s.String()
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleGroupItemSubItemsRefundChangeRule) GetCancelFeeInd() *bool {
	return s.CancelFeeInd
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleGroupItemSubItemsRefundChangeRule) GetChangeFeeInd() *bool {
	return s.ChangeFeeInd
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleGroupItemSubItemsRefundChangeRule) GetOfferPenaltyInfoMap() map[string][]*ModuleGroupItemSubItemsRefundChangeRuleOfferPenaltyInfoMapValue {
	return s.OfferPenaltyInfoMap
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleGroupItemSubItemsRefundChangeRule) GetRefundChangeDigest() *string {
	return s.RefundChangeDigest
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleGroupItemSubItemsRefundChangeRule) GetStructuredRefund() *bool {
	return s.StructuredRefund
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleGroupItemSubItemsRefundChangeRule) SetCancelFeeInd(v bool) *IntlFlightOtaItemDetailResponseBodyModuleGroupItemSubItemsRefundChangeRule {
	s.CancelFeeInd = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleGroupItemSubItemsRefundChangeRule) SetChangeFeeInd(v bool) *IntlFlightOtaItemDetailResponseBodyModuleGroupItemSubItemsRefundChangeRule {
	s.ChangeFeeInd = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleGroupItemSubItemsRefundChangeRule) SetOfferPenaltyInfoMap(v map[string][]*ModuleGroupItemSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) *IntlFlightOtaItemDetailResponseBodyModuleGroupItemSubItemsRefundChangeRule {
	s.OfferPenaltyInfoMap = v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleGroupItemSubItemsRefundChangeRule) SetRefundChangeDigest(v string) *IntlFlightOtaItemDetailResponseBodyModuleGroupItemSubItemsRefundChangeRule {
	s.RefundChangeDigest = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleGroupItemSubItemsRefundChangeRule) SetStructuredRefund(v bool) *IntlFlightOtaItemDetailResponseBodyModuleGroupItemSubItemsRefundChangeRule {
	s.StructuredRefund = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleGroupItemSubItemsRefundChangeRule) Validate() error {
	return dara.Validate(s)
}

type IntlFlightOtaItemDetailResponseBodyModuleShutterDocs struct {
	Contents  []*string `json:"contents,omitempty" xml:"contents,omitempty" type:"Repeated"`
	MainTitle *string   `json:"main_title,omitempty" xml:"main_title,omitempty"`
}

func (s IntlFlightOtaItemDetailResponseBodyModuleShutterDocs) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOtaItemDetailResponseBodyModuleShutterDocs) GoString() string {
	return s.String()
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleShutterDocs) GetContents() []*string {
	return s.Contents
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleShutterDocs) GetMainTitle() *string {
	return s.MainTitle
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleShutterDocs) SetContents(v []*string) *IntlFlightOtaItemDetailResponseBodyModuleShutterDocs {
	s.Contents = v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleShutterDocs) SetMainTitle(v string) *IntlFlightOtaItemDetailResponseBodyModuleShutterDocs {
	s.MainTitle = &v
	return s
}

func (s *IntlFlightOtaItemDetailResponseBodyModuleShutterDocs) Validate() error {
	return dara.Validate(s)
}
