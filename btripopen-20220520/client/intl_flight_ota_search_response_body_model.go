// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightOtaSearchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *IntlFlightOtaSearchResponseBody
	GetCode() *string
	SetMessage(v string) *IntlFlightOtaSearchResponseBody
	GetMessage() *string
	SetModule(v *IntlFlightOtaSearchResponseBodyModule) *IntlFlightOtaSearchResponseBody
	GetModule() *IntlFlightOtaSearchResponseBodyModule
	SetRequestId(v string) *IntlFlightOtaSearchResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *IntlFlightOtaSearchResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *IntlFlightOtaSearchResponseBody
	GetTraceId() *string
}

type IntlFlightOtaSearchResponseBody struct {
	// example:
	//
	// 200
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// module。
	Module *IntlFlightOtaSearchResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
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
	// 2150435016896473589786246e03e9
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s IntlFlightOtaSearchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOtaSearchResponseBody) GoString() string {
	return s.String()
}

func (s *IntlFlightOtaSearchResponseBody) GetCode() *string {
	return s.Code
}

func (s *IntlFlightOtaSearchResponseBody) GetMessage() *string {
	return s.Message
}

func (s *IntlFlightOtaSearchResponseBody) GetModule() *IntlFlightOtaSearchResponseBodyModule {
	return s.Module
}

func (s *IntlFlightOtaSearchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *IntlFlightOtaSearchResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *IntlFlightOtaSearchResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *IntlFlightOtaSearchResponseBody) SetCode(v string) *IntlFlightOtaSearchResponseBody {
	s.Code = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBody) SetMessage(v string) *IntlFlightOtaSearchResponseBody {
	s.Message = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBody) SetModule(v *IntlFlightOtaSearchResponseBodyModule) *IntlFlightOtaSearchResponseBody {
	s.Module = v
	return s
}

func (s *IntlFlightOtaSearchResponseBody) SetRequestId(v string) *IntlFlightOtaSearchResponseBody {
	s.RequestId = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBody) SetSuccess(v bool) *IntlFlightOtaSearchResponseBody {
	s.Success = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBody) SetTraceId(v string) *IntlFlightOtaSearchResponseBody {
	s.TraceId = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBody) Validate() error {
	return dara.Validate(s)
}

type IntlFlightOtaSearchResponseBodyModule struct {
	FlightJourneyInfos []*IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfos `json:"flight_journey_infos,omitempty" xml:"flight_journey_infos,omitempty" type:"Repeated"`
	ItemList           []*IntlFlightOtaSearchResponseBodyModuleItemList           `json:"item_list,omitempty" xml:"item_list,omitempty" type:"Repeated"`
}

func (s IntlFlightOtaSearchResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOtaSearchResponseBodyModule) GoString() string {
	return s.String()
}

func (s *IntlFlightOtaSearchResponseBodyModule) GetFlightJourneyInfos() []*IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfos {
	return s.FlightJourneyInfos
}

func (s *IntlFlightOtaSearchResponseBodyModule) GetItemList() []*IntlFlightOtaSearchResponseBodyModuleItemList {
	return s.ItemList
}

func (s *IntlFlightOtaSearchResponseBodyModule) SetFlightJourneyInfos(v []*IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfos) *IntlFlightOtaSearchResponseBodyModule {
	s.FlightJourneyInfos = v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModule) SetItemList(v []*IntlFlightOtaSearchResponseBodyModuleItemList) *IntlFlightOtaSearchResponseBodyModule {
	s.ItemList = v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfos struct {
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
	Duration           *int32                                                                       `json:"duration,omitempty" xml:"duration,omitempty"`
	FlightSegmentInfos []*IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos `json:"flight_segment_infos,omitempty" xml:"flight_segment_infos,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	JourneyIndex *int32 `json:"journey_index,omitempty" xml:"journey_index,omitempty"`
	// example:
	//
	// 0
	TransferTime *int32 `json:"transfer_time,omitempty" xml:"transfer_time,omitempty"`
}

func (s IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfos) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfos) GoString() string {
	return s.String()
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfos) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfos) GetArrCityName() *string {
	return s.ArrCityName
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfos) GetArrTime() *string {
	return s.ArrTime
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfos) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfos) GetDepCityName() *string {
	return s.DepCityName
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfos) GetDepTime() *string {
	return s.DepTime
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfos) GetDuration() *int32 {
	return s.Duration
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfos) GetFlightSegmentInfos() []*IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	return s.FlightSegmentInfos
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfos) GetJourneyIndex() *int32 {
	return s.JourneyIndex
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfos) GetTransferTime() *int32 {
	return s.TransferTime
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfos) SetArrCityCode(v string) *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfos {
	s.ArrCityCode = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfos) SetArrCityName(v string) *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfos {
	s.ArrCityName = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfos) SetArrTime(v string) *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfos {
	s.ArrTime = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfos) SetDepCityCode(v string) *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfos {
	s.DepCityCode = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfos) SetDepCityName(v string) *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfos {
	s.DepCityName = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfos) SetDepTime(v string) *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfos {
	s.DepTime = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfos) SetDuration(v int32) *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfos {
	s.Duration = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfos) SetFlightSegmentInfos(v []*IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfos {
	s.FlightSegmentInfos = v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfos) SetJourneyIndex(v int32) *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfos {
	s.JourneyIndex = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfos) SetTransferTime(v int32) *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfos {
	s.TransferTime = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfos) Validate() error {
	return dara.Validate(s)
}

type IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos struct {
	AirlineInfo    *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo    `json:"airline_info,omitempty" xml:"airline_info,omitempty" type:"Struct"`
	ArrAirportInfo *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo `json:"arr_airport_info,omitempty" xml:"arr_airport_info,omitempty" type:"Struct"`
	// example:
	//
	// BJS
	ArrCityCode *string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	ArrCityName *string `json:"arr_city_name,omitempty" xml:"arr_city_name,omitempty"`
	// example:
	//
	// 2023-08-13 09:45
	ArrTime        *string                                                                                  `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	DepAirportInfo *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo `json:"dep_airport_info,omitempty" xml:"dep_airport_info,omitempty" type:"Struct"`
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
	FlightNo           *string                                                                                        `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
	FlightShareInfo    *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfo      `json:"flight_share_info,omitempty" xml:"flight_share_info,omitempty" type:"Struct"`
	FlightSize         *string                                                                                        `json:"flight_size,omitempty" xml:"flight_size,omitempty"`
	FlightStopInfoList []*IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoList `json:"flight_stop_info_list,omitempty" xml:"flight_stop_info_list,omitempty" type:"Repeated"`
	// example:
	//
	// 320
	FlightType        *string                                                                                     `json:"flight_type,omitempty" xml:"flight_type,omitempty"`
	LuggageDirectInfo *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosLuggageDirectInfo `json:"luggage_direct_info,omitempty" xml:"luggage_direct_info,omitempty" type:"Struct"`
	Manufacturer      *string                                                                                     `json:"manufacturer,omitempty" xml:"manufacturer,omitempty"`
	MealDesc          *string                                                                                     `json:"meal_desc,omitempty" xml:"meal_desc,omitempty"`
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
	SegmentKey        *string                                                                                     `json:"segment_key,omitempty" xml:"segment_key,omitempty"`
	SegmentVisaRemark *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark `json:"segment_visa_remark,omitempty" xml:"segment_visa_remark,omitempty" type:"Struct"`
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

func (s IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GoString() string {
	return s.String()
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetAirlineInfo() *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo {
	return s.AirlineInfo
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetArrAirportInfo() *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo {
	return s.ArrAirportInfo
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetArrCityName() *string {
	return s.ArrCityName
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetArrTime() *string {
	return s.ArrTime
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetDepAirportInfo() *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo {
	return s.DepAirportInfo
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetDepCityName() *string {
	return s.DepCityName
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetDepTime() *string {
	return s.DepTime
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetDuration() *int32 {
	return s.Duration
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetFlightNo() *string {
	return s.FlightNo
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetFlightShareInfo() *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfo {
	return s.FlightShareInfo
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetFlightSize() *string {
	return s.FlightSize
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetFlightStopInfoList() []*IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoList {
	return s.FlightStopInfoList
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetFlightType() *string {
	return s.FlightType
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetLuggageDirectInfo() *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosLuggageDirectInfo {
	return s.LuggageDirectInfo
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetManufacturer() *string {
	return s.Manufacturer
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetMealDesc() *string {
	return s.MealDesc
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetOneMore() *int32 {
	return s.OneMore
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetOneMoreShow() *string {
	return s.OneMoreShow
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetSegmentIndex() *int32 {
	return s.SegmentIndex
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetSegmentKey() *string {
	return s.SegmentKey
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetSegmentVisaRemark() *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark {
	return s.SegmentVisaRemark
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetShare() *bool {
	return s.Share
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetShortFlightSize() *string {
	return s.ShortFlightSize
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetStop() *bool {
	return s.Stop
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) GetTotalTime() *string {
	return s.TotalTime
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetAirlineInfo(v *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo) *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.AirlineInfo = v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetArrAirportInfo(v *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo) *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.ArrAirportInfo = v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetArrCityCode(v string) *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.ArrCityCode = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetArrCityName(v string) *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.ArrCityName = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetArrTime(v string) *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.ArrTime = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetDepAirportInfo(v *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo) *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.DepAirportInfo = v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetDepCityCode(v string) *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.DepCityCode = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetDepCityName(v string) *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.DepCityName = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetDepTime(v string) *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.DepTime = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetDuration(v int32) *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.Duration = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetFlightNo(v string) *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.FlightNo = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetFlightShareInfo(v *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfo) *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.FlightShareInfo = v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetFlightSize(v string) *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.FlightSize = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetFlightStopInfoList(v []*IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.FlightStopInfoList = v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetFlightType(v string) *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.FlightType = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetLuggageDirectInfo(v *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosLuggageDirectInfo) *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.LuggageDirectInfo = v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetManufacturer(v string) *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.Manufacturer = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetMealDesc(v string) *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.MealDesc = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetOneMore(v int32) *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.OneMore = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetOneMoreShow(v string) *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.OneMoreShow = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetSegmentIndex(v int32) *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.SegmentIndex = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetSegmentKey(v string) *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.SegmentKey = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetSegmentVisaRemark(v *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark) *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.SegmentVisaRemark = v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetShare(v bool) *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.Share = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetShortFlightSize(v string) *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.ShortFlightSize = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetStop(v bool) *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.Stop = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) SetTotalTime(v string) *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos {
	s.TotalTime = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfos) Validate() error {
	return dara.Validate(s)
}

type IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo struct {
	// example:
	//
	// MU
	AirlineCode *string `json:"airline_code,omitempty" xml:"airline_code,omitempty"`
	AirlineName *string `json:"airline_name,omitempty" xml:"airline_name,omitempty"`
	ShortName   *string `json:"short_name,omitempty" xml:"short_name,omitempty"`
}

func (s IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo) GoString() string {
	return s.String()
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo) GetAirlineCode() *string {
	return s.AirlineCode
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo) GetAirlineName() *string {
	return s.AirlineName
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo) GetShortName() *string {
	return s.ShortName
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo) SetAirlineCode(v string) *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo {
	s.AirlineCode = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo) SetAirlineName(v string) *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo {
	s.AirlineName = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo) SetShortName(v string) *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo {
	s.ShortName = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosAirlineInfo) Validate() error {
	return dara.Validate(s)
}

type IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo struct {
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

func (s IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo) GoString() string {
	return s.String()
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo) GetAirportCode() *string {
	return s.AirportCode
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo) GetAirportName() *string {
	return s.AirportName
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo) GetAirportShortName() *string {
	return s.AirportShortName
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo) GetTerminal() *string {
	return s.Terminal
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo) SetAirportCode(v string) *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo {
	s.AirportCode = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo) SetAirportName(v string) *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo {
	s.AirportName = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo) SetAirportShortName(v string) *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo {
	s.AirportShortName = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo) SetTerminal(v string) *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo {
	s.Terminal = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosArrAirportInfo) Validate() error {
	return dara.Validate(s)
}

type IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo struct {
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

func (s IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo) GoString() string {
	return s.String()
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo) GetAirportCode() *string {
	return s.AirportCode
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo) GetAirportName() *string {
	return s.AirportName
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo) GetAirportShortName() *string {
	return s.AirportShortName
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo) GetTerminal() *string {
	return s.Terminal
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo) SetAirportCode(v string) *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo {
	s.AirportCode = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo) SetAirportName(v string) *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo {
	s.AirportName = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo) SetAirportShortName(v string) *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo {
	s.AirportShortName = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo) SetTerminal(v string) *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo {
	s.Terminal = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosDepAirportInfo) Validate() error {
	return dara.Validate(s)
}

type IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfo struct {
	OperatingAirlineInfo *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo `json:"operating_airline_info,omitempty" xml:"operating_airline_info,omitempty" type:"Struct"`
	// example:
	//
	// CX601
	OperatingFlightNo *string `json:"operating_flight_no,omitempty" xml:"operating_flight_no,omitempty"`
}

func (s IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfo) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfo) GoString() string {
	return s.String()
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfo) GetOperatingAirlineInfo() *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo {
	return s.OperatingAirlineInfo
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfo) GetOperatingFlightNo() *string {
	return s.OperatingFlightNo
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfo) SetOperatingAirlineInfo(v *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfo {
	s.OperatingAirlineInfo = v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfo) SetOperatingFlightNo(v string) *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfo {
	s.OperatingFlightNo = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfo) Validate() error {
	return dara.Validate(s)
}

type IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo struct {
	// example:
	//
	// DR
	AirlineCode *string `json:"airline_code,omitempty" xml:"airline_code,omitempty"`
	AirlineName *string `json:"airline_name,omitempty" xml:"airline_name,omitempty"`
	ShortName   *string `json:"short_name,omitempty" xml:"short_name,omitempty"`
}

func (s IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) GoString() string {
	return s.String()
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) GetAirlineCode() *string {
	return s.AirlineCode
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) GetAirlineName() *string {
	return s.AirlineName
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) GetShortName() *string {
	return s.ShortName
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) SetAirlineCode(v string) *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo {
	s.AirlineCode = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) SetAirlineName(v string) *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo {
	s.AirlineName = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) SetShortName(v string) *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo {
	s.ShortName = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightShareInfoOperatingAirlineInfo) Validate() error {
	return dara.Validate(s)
}

type IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoList struct {
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

func (s IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) GoString() string {
	return s.String()
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) GetStopAirport() *string {
	return s.StopAirport
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) GetStopAirportName() *string {
	return s.StopAirportName
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) GetStopArrTerm() *string {
	return s.StopArrTerm
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) GetStopArrTime() *string {
	return s.StopArrTime
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) GetStopCityCode() *string {
	return s.StopCityCode
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) GetStopCityName() *string {
	return s.StopCityName
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) GetStopDepTerm() *string {
	return s.StopDepTerm
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) GetStopDepTime() *string {
	return s.StopDepTime
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) GetStopTime() *string {
	return s.StopTime
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) SetStopAirport(v string) *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoList {
	s.StopAirport = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) SetStopAirportName(v string) *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoList {
	s.StopAirportName = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) SetStopArrTerm(v string) *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoList {
	s.StopArrTerm = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) SetStopArrTime(v string) *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoList {
	s.StopArrTime = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) SetStopCityCode(v string) *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoList {
	s.StopCityCode = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) SetStopCityName(v string) *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoList {
	s.StopCityName = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) SetStopDepTerm(v string) *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoList {
	s.StopDepTerm = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) SetStopDepTime(v string) *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoList {
	s.StopDepTime = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) SetStopTime(v string) *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoList {
	s.StopTime = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosFlightStopInfoList) Validate() error {
	return dara.Validate(s)
}

type IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosLuggageDirectInfo struct {
	// example:
	//
	// 1
	DepCityLuggageDirect *int32 `json:"dep_city_luggage_direct,omitempty" xml:"dep_city_luggage_direct,omitempty"`
	// example:
	//
	// 0
	StopCityLuggageDirect *int32 `json:"stop_city_luggage_direct,omitempty" xml:"stop_city_luggage_direct,omitempty"`
}

func (s IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosLuggageDirectInfo) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosLuggageDirectInfo) GoString() string {
	return s.String()
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosLuggageDirectInfo) GetDepCityLuggageDirect() *int32 {
	return s.DepCityLuggageDirect
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosLuggageDirectInfo) GetStopCityLuggageDirect() *int32 {
	return s.StopCityLuggageDirect
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosLuggageDirectInfo) SetDepCityLuggageDirect(v int32) *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosLuggageDirectInfo {
	s.DepCityLuggageDirect = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosLuggageDirectInfo) SetStopCityLuggageDirect(v int32) *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosLuggageDirectInfo {
	s.StopCityLuggageDirect = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosLuggageDirectInfo) Validate() error {
	return dara.Validate(s)
}

type IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark struct {
	DepCityVisaRemark *string `json:"dep_city_visa_remark,omitempty" xml:"dep_city_visa_remark,omitempty"`
	// example:
	//
	// 1
	DepCityVisaType     *int32    `json:"dep_city_visa_type,omitempty" xml:"dep_city_visa_type,omitempty"`
	StopCityVisaRemarks []*string `json:"stop_city_visa_remarks,omitempty" xml:"stop_city_visa_remarks,omitempty" type:"Repeated"`
	StopCityVisaTypes   []*int32  `json:"stop_city_visa_types,omitempty" xml:"stop_city_visa_types,omitempty" type:"Repeated"`
}

func (s IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark) GoString() string {
	return s.String()
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark) GetDepCityVisaRemark() *string {
	return s.DepCityVisaRemark
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark) GetDepCityVisaType() *int32 {
	return s.DepCityVisaType
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark) GetStopCityVisaRemarks() []*string {
	return s.StopCityVisaRemarks
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark) GetStopCityVisaTypes() []*int32 {
	return s.StopCityVisaTypes
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark) SetDepCityVisaRemark(v string) *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark {
	s.DepCityVisaRemark = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark) SetDepCityVisaType(v int32) *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark {
	s.DepCityVisaType = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark) SetStopCityVisaRemarks(v []*string) *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark {
	s.StopCityVisaRemarks = v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark) SetStopCityVisaTypes(v []*int32) *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark {
	s.StopCityVisaTypes = v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleFlightJourneyInfosFlightSegmentInfosSegmentVisaRemark) Validate() error {
	return dara.Validate(s)
}

type IntlFlightOtaSearchResponseBodyModuleItemList struct {
	AgreementPriceCodes []*IntlFlightOtaSearchResponseBodyModuleItemListAgreementPriceCodes `json:"agreement_price_codes,omitempty" xml:"agreement_price_codes,omitempty" type:"Repeated"`
	// example:
	//
	// e50d380fc05942cc8ac57af8ae02f448_0
	ItemId          *string                                                   `json:"item_id,omitempty" xml:"item_id,omitempty"`
	ItemType        *string                                                   `json:"item_type,omitempty" xml:"item_type,omitempty"`
	LabelList       []*IntlFlightOtaSearchResponseBodyModuleItemListLabelList `json:"label_list,omitempty" xml:"label_list,omitempty" type:"Repeated"`
	ShoppingItemMap map[string]*ModuleItemListShoppingItemMapValue            `json:"shopping_item_map,omitempty" xml:"shopping_item_map,omitempty"`
	SubItems        []*IntlFlightOtaSearchResponseBodyModuleItemListSubItems  `json:"sub_items,omitempty" xml:"sub_items,omitempty" type:"Repeated"`
}

func (s IntlFlightOtaSearchResponseBodyModuleItemList) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOtaSearchResponseBodyModuleItemList) GoString() string {
	return s.String()
}

func (s *IntlFlightOtaSearchResponseBodyModuleItemList) GetAgreementPriceCodes() []*IntlFlightOtaSearchResponseBodyModuleItemListAgreementPriceCodes {
	return s.AgreementPriceCodes
}

func (s *IntlFlightOtaSearchResponseBodyModuleItemList) GetItemId() *string {
	return s.ItemId
}

func (s *IntlFlightOtaSearchResponseBodyModuleItemList) GetItemType() *string {
	return s.ItemType
}

func (s *IntlFlightOtaSearchResponseBodyModuleItemList) GetLabelList() []*IntlFlightOtaSearchResponseBodyModuleItemListLabelList {
	return s.LabelList
}

func (s *IntlFlightOtaSearchResponseBodyModuleItemList) GetShoppingItemMap() map[string]*ModuleItemListShoppingItemMapValue {
	return s.ShoppingItemMap
}

func (s *IntlFlightOtaSearchResponseBodyModuleItemList) GetSubItems() []*IntlFlightOtaSearchResponseBodyModuleItemListSubItems {
	return s.SubItems
}

func (s *IntlFlightOtaSearchResponseBodyModuleItemList) SetAgreementPriceCodes(v []*IntlFlightOtaSearchResponseBodyModuleItemListAgreementPriceCodes) *IntlFlightOtaSearchResponseBodyModuleItemList {
	s.AgreementPriceCodes = v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleItemList) SetItemId(v string) *IntlFlightOtaSearchResponseBodyModuleItemList {
	s.ItemId = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleItemList) SetItemType(v string) *IntlFlightOtaSearchResponseBodyModuleItemList {
	s.ItemType = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleItemList) SetLabelList(v []*IntlFlightOtaSearchResponseBodyModuleItemListLabelList) *IntlFlightOtaSearchResponseBodyModuleItemList {
	s.LabelList = v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleItemList) SetShoppingItemMap(v map[string]*ModuleItemListShoppingItemMapValue) *IntlFlightOtaSearchResponseBodyModuleItemList {
	s.ShoppingItemMap = v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleItemList) SetSubItems(v []*IntlFlightOtaSearchResponseBodyModuleItemListSubItems) *IntlFlightOtaSearchResponseBodyModuleItemList {
	s.SubItems = v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleItemList) Validate() error {
	return dara.Validate(s)
}

type IntlFlightOtaSearchResponseBodyModuleItemListAgreementPriceCodes struct {
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	Type *int32  `json:"type,omitempty" xml:"type,omitempty"`
}

func (s IntlFlightOtaSearchResponseBodyModuleItemListAgreementPriceCodes) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOtaSearchResponseBodyModuleItemListAgreementPriceCodes) GoString() string {
	return s.String()
}

func (s *IntlFlightOtaSearchResponseBodyModuleItemListAgreementPriceCodes) GetCode() *string {
	return s.Code
}

func (s *IntlFlightOtaSearchResponseBodyModuleItemListAgreementPriceCodes) GetType() *int32 {
	return s.Type
}

func (s *IntlFlightOtaSearchResponseBodyModuleItemListAgreementPriceCodes) SetCode(v string) *IntlFlightOtaSearchResponseBodyModuleItemListAgreementPriceCodes {
	s.Code = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleItemListAgreementPriceCodes) SetType(v int32) *IntlFlightOtaSearchResponseBodyModuleItemListAgreementPriceCodes {
	s.Type = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleItemListAgreementPriceCodes) Validate() error {
	return dara.Validate(s)
}

type IntlFlightOtaSearchResponseBodyModuleItemListLabelList struct {
	LabelCode *int32  `json:"labelCode,omitempty" xml:"labelCode,omitempty"`
	LabelName *string `json:"labelName,omitempty" xml:"labelName,omitempty"`
}

func (s IntlFlightOtaSearchResponseBodyModuleItemListLabelList) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOtaSearchResponseBodyModuleItemListLabelList) GoString() string {
	return s.String()
}

func (s *IntlFlightOtaSearchResponseBodyModuleItemListLabelList) GetLabelCode() *int32 {
	return s.LabelCode
}

func (s *IntlFlightOtaSearchResponseBodyModuleItemListLabelList) GetLabelName() *string {
	return s.LabelName
}

func (s *IntlFlightOtaSearchResponseBodyModuleItemListLabelList) SetLabelCode(v int32) *IntlFlightOtaSearchResponseBodyModuleItemListLabelList {
	s.LabelCode = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleItemListLabelList) SetLabelName(v string) *IntlFlightOtaSearchResponseBodyModuleItemListLabelList {
	s.LabelName = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleItemListLabelList) Validate() error {
	return dara.Validate(s)
}

type IntlFlightOtaSearchResponseBodyModuleItemListSubItems struct {
	BaggageRule         *IntlFlightOtaSearchResponseBodyModuleItemListSubItemsBaggageRule           `json:"baggage_rule,omitempty" xml:"baggage_rule,omitempty" type:"Struct"`
	RefundChangeRule    *IntlFlightOtaSearchResponseBodyModuleItemListSubItemsRefundChangeRule      `json:"refund_change_rule,omitempty" xml:"refund_change_rule,omitempty" type:"Struct"`
	SegmentKeys         []*string                                                                   `json:"segment_keys,omitempty" xml:"segment_keys,omitempty" type:"Repeated"`
	SegmentPositionList []*IntlFlightOtaSearchResponseBodyModuleItemListSubItemsSegmentPositionList `json:"segment_position_list,omitempty" xml:"segment_position_list,omitempty" type:"Repeated"`
	ShoppingItemMap     map[string]*ModuleItemListSubItemsShoppingItemMapValue                      `json:"shopping_item_map,omitempty" xml:"shopping_item_map,omitempty"`
	// example:
	//
	// 07df0bd9-f803-4a50-8449-f4bd675d9939
	UniqKey *string `json:"uniq_key,omitempty" xml:"uniq_key,omitempty"`
}

func (s IntlFlightOtaSearchResponseBodyModuleItemListSubItems) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOtaSearchResponseBodyModuleItemListSubItems) GoString() string {
	return s.String()
}

func (s *IntlFlightOtaSearchResponseBodyModuleItemListSubItems) GetBaggageRule() *IntlFlightOtaSearchResponseBodyModuleItemListSubItemsBaggageRule {
	return s.BaggageRule
}

func (s *IntlFlightOtaSearchResponseBodyModuleItemListSubItems) GetRefundChangeRule() *IntlFlightOtaSearchResponseBodyModuleItemListSubItemsRefundChangeRule {
	return s.RefundChangeRule
}

func (s *IntlFlightOtaSearchResponseBodyModuleItemListSubItems) GetSegmentKeys() []*string {
	return s.SegmentKeys
}

func (s *IntlFlightOtaSearchResponseBodyModuleItemListSubItems) GetSegmentPositionList() []*IntlFlightOtaSearchResponseBodyModuleItemListSubItemsSegmentPositionList {
	return s.SegmentPositionList
}

func (s *IntlFlightOtaSearchResponseBodyModuleItemListSubItems) GetShoppingItemMap() map[string]*ModuleItemListSubItemsShoppingItemMapValue {
	return s.ShoppingItemMap
}

func (s *IntlFlightOtaSearchResponseBodyModuleItemListSubItems) GetUniqKey() *string {
	return s.UniqKey
}

func (s *IntlFlightOtaSearchResponseBodyModuleItemListSubItems) SetBaggageRule(v *IntlFlightOtaSearchResponseBodyModuleItemListSubItemsBaggageRule) *IntlFlightOtaSearchResponseBodyModuleItemListSubItems {
	s.BaggageRule = v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleItemListSubItems) SetRefundChangeRule(v *IntlFlightOtaSearchResponseBodyModuleItemListSubItemsRefundChangeRule) *IntlFlightOtaSearchResponseBodyModuleItemListSubItems {
	s.RefundChangeRule = v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleItemListSubItems) SetSegmentKeys(v []*string) *IntlFlightOtaSearchResponseBodyModuleItemListSubItems {
	s.SegmentKeys = v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleItemListSubItems) SetSegmentPositionList(v []*IntlFlightOtaSearchResponseBodyModuleItemListSubItemsSegmentPositionList) *IntlFlightOtaSearchResponseBodyModuleItemListSubItems {
	s.SegmentPositionList = v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleItemListSubItems) SetShoppingItemMap(v map[string]*ModuleItemListSubItemsShoppingItemMapValue) *IntlFlightOtaSearchResponseBodyModuleItemListSubItems {
	s.ShoppingItemMap = v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleItemListSubItems) SetUniqKey(v string) *IntlFlightOtaSearchResponseBodyModuleItemListSubItems {
	s.UniqKey = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleItemListSubItems) Validate() error {
	return dara.Validate(s)
}

type IntlFlightOtaSearchResponseBodyModuleItemListSubItemsBaggageRule struct {
	BaggageDigest       *string                                                                 `json:"baggage_digest,omitempty" xml:"baggage_digest,omitempty"`
	OfferBaggageInfoMap map[string][]*ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue `json:"offer_baggage_info_map,omitempty" xml:"offer_baggage_info_map,omitempty"`
	// example:
	//
	// true
	StructuredBaggage *bool `json:"structured_baggage,omitempty" xml:"structured_baggage,omitempty"`
}

func (s IntlFlightOtaSearchResponseBodyModuleItemListSubItemsBaggageRule) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOtaSearchResponseBodyModuleItemListSubItemsBaggageRule) GoString() string {
	return s.String()
}

func (s *IntlFlightOtaSearchResponseBodyModuleItemListSubItemsBaggageRule) GetBaggageDigest() *string {
	return s.BaggageDigest
}

func (s *IntlFlightOtaSearchResponseBodyModuleItemListSubItemsBaggageRule) GetOfferBaggageInfoMap() map[string][]*ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue {
	return s.OfferBaggageInfoMap
}

func (s *IntlFlightOtaSearchResponseBodyModuleItemListSubItemsBaggageRule) GetStructuredBaggage() *bool {
	return s.StructuredBaggage
}

func (s *IntlFlightOtaSearchResponseBodyModuleItemListSubItemsBaggageRule) SetBaggageDigest(v string) *IntlFlightOtaSearchResponseBodyModuleItemListSubItemsBaggageRule {
	s.BaggageDigest = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleItemListSubItemsBaggageRule) SetOfferBaggageInfoMap(v map[string][]*ModuleItemListSubItemsBaggageRuleOfferBaggageInfoMapValue) *IntlFlightOtaSearchResponseBodyModuleItemListSubItemsBaggageRule {
	s.OfferBaggageInfoMap = v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleItemListSubItemsBaggageRule) SetStructuredBaggage(v bool) *IntlFlightOtaSearchResponseBodyModuleItemListSubItemsBaggageRule {
	s.StructuredBaggage = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleItemListSubItemsBaggageRule) Validate() error {
	return dara.Validate(s)
}

type IntlFlightOtaSearchResponseBodyModuleItemListSubItemsRefundChangeRule struct {
	// example:
	//
	// true
	CancelFeeInd *bool `json:"cancel_fee_ind,omitempty" xml:"cancel_fee_ind,omitempty"`
	// example:
	//
	// true
	ChangeFeeInd        *bool                                                                        `json:"change_fee_ind,omitempty" xml:"change_fee_ind,omitempty"`
	ChangeRuleDesc      *string                                                                      `json:"change_rule_desc,omitempty" xml:"change_rule_desc,omitempty"`
	OfferPenaltyInfoMap map[string][]*ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue `json:"offer_penalty_info_map,omitempty" xml:"offer_penalty_info_map,omitempty"`
	RefundChangeDigest  *string                                                                      `json:"refund_change_digest,omitempty" xml:"refund_change_digest,omitempty"`
	RefundRuleDesc      *string                                                                      `json:"refund_rule_desc,omitempty" xml:"refund_rule_desc,omitempty"`
	// example:
	//
	// false
	StructuredRefund *bool `json:"structured_refund,omitempty" xml:"structured_refund,omitempty"`
}

func (s IntlFlightOtaSearchResponseBodyModuleItemListSubItemsRefundChangeRule) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOtaSearchResponseBodyModuleItemListSubItemsRefundChangeRule) GoString() string {
	return s.String()
}

func (s *IntlFlightOtaSearchResponseBodyModuleItemListSubItemsRefundChangeRule) GetCancelFeeInd() *bool {
	return s.CancelFeeInd
}

func (s *IntlFlightOtaSearchResponseBodyModuleItemListSubItemsRefundChangeRule) GetChangeFeeInd() *bool {
	return s.ChangeFeeInd
}

func (s *IntlFlightOtaSearchResponseBodyModuleItemListSubItemsRefundChangeRule) GetChangeRuleDesc() *string {
	return s.ChangeRuleDesc
}

func (s *IntlFlightOtaSearchResponseBodyModuleItemListSubItemsRefundChangeRule) GetOfferPenaltyInfoMap() map[string][]*ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue {
	return s.OfferPenaltyInfoMap
}

func (s *IntlFlightOtaSearchResponseBodyModuleItemListSubItemsRefundChangeRule) GetRefundChangeDigest() *string {
	return s.RefundChangeDigest
}

func (s *IntlFlightOtaSearchResponseBodyModuleItemListSubItemsRefundChangeRule) GetRefundRuleDesc() *string {
	return s.RefundRuleDesc
}

func (s *IntlFlightOtaSearchResponseBodyModuleItemListSubItemsRefundChangeRule) GetStructuredRefund() *bool {
	return s.StructuredRefund
}

func (s *IntlFlightOtaSearchResponseBodyModuleItemListSubItemsRefundChangeRule) SetCancelFeeInd(v bool) *IntlFlightOtaSearchResponseBodyModuleItemListSubItemsRefundChangeRule {
	s.CancelFeeInd = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleItemListSubItemsRefundChangeRule) SetChangeFeeInd(v bool) *IntlFlightOtaSearchResponseBodyModuleItemListSubItemsRefundChangeRule {
	s.ChangeFeeInd = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleItemListSubItemsRefundChangeRule) SetChangeRuleDesc(v string) *IntlFlightOtaSearchResponseBodyModuleItemListSubItemsRefundChangeRule {
	s.ChangeRuleDesc = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleItemListSubItemsRefundChangeRule) SetOfferPenaltyInfoMap(v map[string][]*ModuleItemListSubItemsRefundChangeRuleOfferPenaltyInfoMapValue) *IntlFlightOtaSearchResponseBodyModuleItemListSubItemsRefundChangeRule {
	s.OfferPenaltyInfoMap = v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleItemListSubItemsRefundChangeRule) SetRefundChangeDigest(v string) *IntlFlightOtaSearchResponseBodyModuleItemListSubItemsRefundChangeRule {
	s.RefundChangeDigest = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleItemListSubItemsRefundChangeRule) SetRefundRuleDesc(v string) *IntlFlightOtaSearchResponseBodyModuleItemListSubItemsRefundChangeRule {
	s.RefundRuleDesc = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleItemListSubItemsRefundChangeRule) SetStructuredRefund(v bool) *IntlFlightOtaSearchResponseBodyModuleItemListSubItemsRefundChangeRule {
	s.StructuredRefund = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleItemListSubItemsRefundChangeRule) Validate() error {
	return dara.Validate(s)
}

type IntlFlightOtaSearchResponseBodyModuleItemListSubItemsSegmentPositionList struct {
	JourneyIndex *int32 `json:"journey_index,omitempty" xml:"journey_index,omitempty"`
	SegmentIndex *int32 `json:"segment_index,omitempty" xml:"segment_index,omitempty"`
}

func (s IntlFlightOtaSearchResponseBodyModuleItemListSubItemsSegmentPositionList) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOtaSearchResponseBodyModuleItemListSubItemsSegmentPositionList) GoString() string {
	return s.String()
}

func (s *IntlFlightOtaSearchResponseBodyModuleItemListSubItemsSegmentPositionList) GetJourneyIndex() *int32 {
	return s.JourneyIndex
}

func (s *IntlFlightOtaSearchResponseBodyModuleItemListSubItemsSegmentPositionList) GetSegmentIndex() *int32 {
	return s.SegmentIndex
}

func (s *IntlFlightOtaSearchResponseBodyModuleItemListSubItemsSegmentPositionList) SetJourneyIndex(v int32) *IntlFlightOtaSearchResponseBodyModuleItemListSubItemsSegmentPositionList {
	s.JourneyIndex = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleItemListSubItemsSegmentPositionList) SetSegmentIndex(v int32) *IntlFlightOtaSearchResponseBodyModuleItemListSubItemsSegmentPositionList {
	s.SegmentIndex = &v
	return s
}

func (s *IntlFlightOtaSearchResponseBodyModuleItemListSubItemsSegmentPositionList) Validate() error {
	return dara.Validate(s)
}
