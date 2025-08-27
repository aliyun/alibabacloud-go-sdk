// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightModifyListingSearchV2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *FlightModifyListingSearchV2ResponseBody
	GetCode() *string
	SetMessage(v string) *FlightModifyListingSearchV2ResponseBody
	GetMessage() *string
	SetModule(v *FlightModifyListingSearchV2ResponseBodyModule) *FlightModifyListingSearchV2ResponseBody
	GetModule() *FlightModifyListingSearchV2ResponseBodyModule
	SetRequestId(v string) *FlightModifyListingSearchV2ResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *FlightModifyListingSearchV2ResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *FlightModifyListingSearchV2ResponseBody
	GetTraceId() *string
}

type FlightModifyListingSearchV2ResponseBody struct {
	// example:
	//
	// 200
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// module
	Module *FlightModifyListingSearchV2ResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
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
	// requestId
	//
	// example:
	//
	// 2136019116915615924561621e06ee
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s FlightModifyListingSearchV2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyListingSearchV2ResponseBody) GoString() string {
	return s.String()
}

func (s *FlightModifyListingSearchV2ResponseBody) GetCode() *string {
	return s.Code
}

func (s *FlightModifyListingSearchV2ResponseBody) GetMessage() *string {
	return s.Message
}

func (s *FlightModifyListingSearchV2ResponseBody) GetModule() *FlightModifyListingSearchV2ResponseBodyModule {
	return s.Module
}

func (s *FlightModifyListingSearchV2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FlightModifyListingSearchV2ResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *FlightModifyListingSearchV2ResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *FlightModifyListingSearchV2ResponseBody) SetCode(v string) *FlightModifyListingSearchV2ResponseBody {
	s.Code = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBody) SetMessage(v string) *FlightModifyListingSearchV2ResponseBody {
	s.Message = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBody) SetModule(v *FlightModifyListingSearchV2ResponseBodyModule) *FlightModifyListingSearchV2ResponseBody {
	s.Module = v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBody) SetRequestId(v string) *FlightModifyListingSearchV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBody) SetSuccess(v bool) *FlightModifyListingSearchV2ResponseBody {
	s.Success = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBody) SetTraceId(v string) *FlightModifyListingSearchV2ResponseBody {
	s.TraceId = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBody) Validate() error {
	return dara.Validate(s)
}

type FlightModifyListingSearchV2ResponseBodyModule struct {
	DirectFlightList []*FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList `json:"direct_flight_list,omitempty" xml:"direct_flight_list,omitempty" type:"Repeated"`
	NextReqWaitTime  *int64                                                           `json:"next_req_wait_time,omitempty" xml:"next_req_wait_time,omitempty"`
	Retry            *bool                                                            `json:"retry,omitempty" xml:"retry,omitempty"`
	SearchRetryToken *string                                                          `json:"search_retry_token,omitempty" xml:"search_retry_token,omitempty"`
	// example:
	//
	// a2ffebfe733742aab5c491d960ba3d59
	SessionId          *string                                                            `json:"session_id,omitempty" xml:"session_id,omitempty"`
	TransferFlightList []*FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList `json:"transfer_flight_list,omitempty" xml:"transfer_flight_list,omitempty" type:"Repeated"`
	TransferTitle      *string                                                            `json:"transfer_title,omitempty" xml:"transfer_title,omitempty"`
}

func (s FlightModifyListingSearchV2ResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyListingSearchV2ResponseBodyModule) GoString() string {
	return s.String()
}

func (s *FlightModifyListingSearchV2ResponseBodyModule) GetDirectFlightList() []*FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList {
	return s.DirectFlightList
}

func (s *FlightModifyListingSearchV2ResponseBodyModule) GetNextReqWaitTime() *int64 {
	return s.NextReqWaitTime
}

func (s *FlightModifyListingSearchV2ResponseBodyModule) GetRetry() *bool {
	return s.Retry
}

func (s *FlightModifyListingSearchV2ResponseBodyModule) GetSearchRetryToken() *string {
	return s.SearchRetryToken
}

func (s *FlightModifyListingSearchV2ResponseBodyModule) GetSessionId() *string {
	return s.SessionId
}

func (s *FlightModifyListingSearchV2ResponseBodyModule) GetTransferFlightList() []*FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList {
	return s.TransferFlightList
}

func (s *FlightModifyListingSearchV2ResponseBodyModule) GetTransferTitle() *string {
	return s.TransferTitle
}

func (s *FlightModifyListingSearchV2ResponseBodyModule) SetDirectFlightList(v []*FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList) *FlightModifyListingSearchV2ResponseBodyModule {
	s.DirectFlightList = v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModule) SetNextReqWaitTime(v int64) *FlightModifyListingSearchV2ResponseBodyModule {
	s.NextReqWaitTime = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModule) SetRetry(v bool) *FlightModifyListingSearchV2ResponseBodyModule {
	s.Retry = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModule) SetSearchRetryToken(v string) *FlightModifyListingSearchV2ResponseBodyModule {
	s.SearchRetryToken = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModule) SetSessionId(v string) *FlightModifyListingSearchV2ResponseBodyModule {
	s.SessionId = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModule) SetTransferFlightList(v []*FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList) *FlightModifyListingSearchV2ResponseBodyModule {
	s.TransferFlightList = v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModule) SetTransferTitle(v string) *FlightModifyListingSearchV2ResponseBodyModule {
	s.TransferTitle = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList struct {
	AirlineInfo    *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListAirlineInfo    `json:"airline_info,omitempty" xml:"airline_info,omitempty" type:"Struct"`
	ArrAirportInfo *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListArrAirportInfo `json:"arr_airport_info,omitempty" xml:"arr_airport_info,omitempty" type:"Struct"`
	// example:
	//
	// BJS
	ArrCityCode *string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	// example:
	//
	// 2023-09-18 10:25:00
	ArrTime        *string                                                                      `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	CabinClass     *string                                                                      `json:"cabinClass,omitempty" xml:"cabinClass,omitempty"`
	CabinClassName *string                                                                      `json:"cabinClassName,omitempty" xml:"cabinClassName,omitempty"`
	DepAirportInfo *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListDepAirportInfo `json:"dep_airport_info,omitempty" xml:"dep_airport_info,omitempty" type:"Struct"`
	// example:
	//
	// XIL
	DepCityCode *string `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	// example:
	//
	// 2023-09-18 09:10:00
	DepTime *string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	// example:
	//
	// 240
	Duration *int32 `json:"duration,omitempty" xml:"duration,omitempty"`
	// example:
	//
	// CA1110
	FlightNo           *string                                                                          `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
	FlightShareInfo    *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightShareInfo    `json:"flight_share_info,omitempty" xml:"flight_share_info,omitempty" type:"Struct"`
	FlightSize         *string                                                                          `json:"flight_size,omitempty" xml:"flight_size,omitempty"`
	FlightStopInfo     *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightStopInfo     `json:"flight_stop_info,omitempty" xml:"flight_stop_info,omitempty" type:"Struct"`
	FlightTransferInfo *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightTransferInfo `json:"flight_transfer_info,omitempty" xml:"flight_transfer_info,omitempty" type:"Struct"`
	// example:
	//
	// ARJ
	FlightType *string `json:"flight_type,omitempty" xml:"flight_type,omitempty"`
	// example:
	//
	// 0
	JourneySeq *int32 `json:"journey_seq,omitempty" xml:"journey_seq,omitempty"`
	// example:
	//
	// 8
	LeftNum      *string                                                                    `json:"left_num,omitempty" xml:"left_num,omitempty"`
	Manufacturer *string                                                                    `json:"manufacturer,omitempty" xml:"manufacturer,omitempty"`
	MealDesc     *string                                                                    `json:"meal_desc,omitempty" xml:"meal_desc,omitempty"`
	PriceInfoDTO *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTO `json:"price_info_d_t_o,omitempty" xml:"price_info_d_t_o,omitempty" type:"Struct"`
	// example:
	//
	// 0
	SegmentSeq *int32 `json:"segment_seq,omitempty" xml:"segment_seq,omitempty"`
	// example:
	//
	// false
	Share           *bool   `json:"share,omitempty" xml:"share,omitempty"`
	ShortFlightSize *string `json:"short_flight_size,omitempty" xml:"short_flight_size,omitempty"`
	SpanDay         *string `json:"span_day,omitempty" xml:"span_day,omitempty"`
	// example:
	//
	// false
	Stop *bool `json:"stop,omitempty" xml:"stop,omitempty"`
	// example:
	//
	// false
	Transfer *bool `json:"transfer,omitempty" xml:"transfer,omitempty"`
}

func (s FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList) GoString() string {
	return s.String()
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList) GetAirlineInfo() *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListAirlineInfo {
	return s.AirlineInfo
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList) GetArrAirportInfo() *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListArrAirportInfo {
	return s.ArrAirportInfo
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList) GetArrTime() *string {
	return s.ArrTime
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList) GetCabinClass() *string {
	return s.CabinClass
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList) GetCabinClassName() *string {
	return s.CabinClassName
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList) GetDepAirportInfo() *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListDepAirportInfo {
	return s.DepAirportInfo
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList) GetDepTime() *string {
	return s.DepTime
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList) GetDuration() *int32 {
	return s.Duration
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList) GetFlightNo() *string {
	return s.FlightNo
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList) GetFlightShareInfo() *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightShareInfo {
	return s.FlightShareInfo
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList) GetFlightSize() *string {
	return s.FlightSize
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList) GetFlightStopInfo() *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightStopInfo {
	return s.FlightStopInfo
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList) GetFlightTransferInfo() *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightTransferInfo {
	return s.FlightTransferInfo
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList) GetFlightType() *string {
	return s.FlightType
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList) GetJourneySeq() *int32 {
	return s.JourneySeq
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList) GetLeftNum() *string {
	return s.LeftNum
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList) GetManufacturer() *string {
	return s.Manufacturer
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList) GetMealDesc() *string {
	return s.MealDesc
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList) GetPriceInfoDTO() *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTO {
	return s.PriceInfoDTO
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList) GetSegmentSeq() *int32 {
	return s.SegmentSeq
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList) GetShare() *bool {
	return s.Share
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList) GetShortFlightSize() *string {
	return s.ShortFlightSize
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList) GetSpanDay() *string {
	return s.SpanDay
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList) GetStop() *bool {
	return s.Stop
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList) GetTransfer() *bool {
	return s.Transfer
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList) SetAirlineInfo(v *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListAirlineInfo) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList {
	s.AirlineInfo = v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList) SetArrAirportInfo(v *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListArrAirportInfo) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList {
	s.ArrAirportInfo = v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList) SetArrCityCode(v string) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList {
	s.ArrCityCode = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList) SetArrTime(v string) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList {
	s.ArrTime = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList) SetCabinClass(v string) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList {
	s.CabinClass = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList) SetCabinClassName(v string) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList {
	s.CabinClassName = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList) SetDepAirportInfo(v *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListDepAirportInfo) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList {
	s.DepAirportInfo = v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList) SetDepCityCode(v string) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList {
	s.DepCityCode = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList) SetDepTime(v string) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList {
	s.DepTime = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList) SetDuration(v int32) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList {
	s.Duration = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList) SetFlightNo(v string) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList {
	s.FlightNo = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList) SetFlightShareInfo(v *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightShareInfo) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList {
	s.FlightShareInfo = v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList) SetFlightSize(v string) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList {
	s.FlightSize = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList) SetFlightStopInfo(v *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightStopInfo) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList {
	s.FlightStopInfo = v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList) SetFlightTransferInfo(v *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightTransferInfo) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList {
	s.FlightTransferInfo = v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList) SetFlightType(v string) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList {
	s.FlightType = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList) SetJourneySeq(v int32) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList {
	s.JourneySeq = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList) SetLeftNum(v string) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList {
	s.LeftNum = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList) SetManufacturer(v string) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList {
	s.Manufacturer = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList) SetMealDesc(v string) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList {
	s.MealDesc = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList) SetPriceInfoDTO(v *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTO) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList {
	s.PriceInfoDTO = v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList) SetSegmentSeq(v int32) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList {
	s.SegmentSeq = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList) SetShare(v bool) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList {
	s.Share = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList) SetShortFlightSize(v string) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList {
	s.ShortFlightSize = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList) SetSpanDay(v string) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList {
	s.SpanDay = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList) SetStop(v bool) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList {
	s.Stop = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList) SetTransfer(v bool) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList {
	s.Transfer = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightList) Validate() error {
	return dara.Validate(s)
}

type FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListAirlineInfo struct {
	AirlineChineseName      *string `json:"airline_chinese_name,omitempty" xml:"airline_chinese_name,omitempty"`
	AirlineChineseShortName *string `json:"airline_chinese_short_name,omitempty" xml:"airline_chinese_short_name,omitempty"`
	// example:
	//
	// CA
	AirlineCode *string `json:"airline_code,omitempty" xml:"airline_code,omitempty"`
	// example:
	//
	// //gw.alicdn.com/tfs/TB12fJAFHr1gK0jSZR0XXbP8XXa-450-450.png_80x80.jpg
	AirlineIcon *string `json:"airline_icon,omitempty" xml:"airline_icon,omitempty"`
	// example:
	//
	// false
	CheapFlight *bool `json:"cheap_flight,omitempty" xml:"cheap_flight,omitempty"`
}

func (s FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListAirlineInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListAirlineInfo) GoString() string {
	return s.String()
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListAirlineInfo) GetAirlineChineseName() *string {
	return s.AirlineChineseName
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListAirlineInfo) GetAirlineChineseShortName() *string {
	return s.AirlineChineseShortName
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListAirlineInfo) GetAirlineCode() *string {
	return s.AirlineCode
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListAirlineInfo) GetAirlineIcon() *string {
	return s.AirlineIcon
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListAirlineInfo) GetCheapFlight() *bool {
	return s.CheapFlight
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListAirlineInfo) SetAirlineChineseName(v string) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListAirlineInfo {
	s.AirlineChineseName = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListAirlineInfo) SetAirlineChineseShortName(v string) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListAirlineInfo {
	s.AirlineChineseShortName = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListAirlineInfo) SetAirlineCode(v string) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListAirlineInfo {
	s.AirlineCode = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListAirlineInfo) SetAirlineIcon(v string) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListAirlineInfo {
	s.AirlineIcon = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListAirlineInfo) SetCheapFlight(v bool) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListAirlineInfo {
	s.CheapFlight = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListAirlineInfo) Validate() error {
	return dara.Validate(s)
}

type FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListArrAirportInfo struct {
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

func (s FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListArrAirportInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListArrAirportInfo) GoString() string {
	return s.String()
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListArrAirportInfo) GetAirportCode() *string {
	return s.AirportCode
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListArrAirportInfo) GetAirportName() *string {
	return s.AirportName
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListArrAirportInfo) GetAirportShortName() *string {
	return s.AirportShortName
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListArrAirportInfo) GetTerminal() *string {
	return s.Terminal
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListArrAirportInfo) SetAirportCode(v string) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListArrAirportInfo {
	s.AirportCode = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListArrAirportInfo) SetAirportName(v string) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListArrAirportInfo {
	s.AirportName = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListArrAirportInfo) SetAirportShortName(v string) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListArrAirportInfo {
	s.AirportShortName = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListArrAirportInfo) SetTerminal(v string) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListArrAirportInfo {
	s.Terminal = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListArrAirportInfo) Validate() error {
	return dara.Validate(s)
}

type FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListDepAirportInfo struct {
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

func (s FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListDepAirportInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListDepAirportInfo) GoString() string {
	return s.String()
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListDepAirportInfo) GetAirportCode() *string {
	return s.AirportCode
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListDepAirportInfo) GetAirportName() *string {
	return s.AirportName
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListDepAirportInfo) GetAirportShortName() *string {
	return s.AirportShortName
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListDepAirportInfo) GetTerminal() *string {
	return s.Terminal
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListDepAirportInfo) SetAirportCode(v string) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListDepAirportInfo {
	s.AirportCode = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListDepAirportInfo) SetAirportName(v string) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListDepAirportInfo {
	s.AirportName = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListDepAirportInfo) SetAirportShortName(v string) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListDepAirportInfo {
	s.AirportShortName = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListDepAirportInfo) SetTerminal(v string) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListDepAirportInfo {
	s.Terminal = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListDepAirportInfo) Validate() error {
	return dara.Validate(s)
}

type FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightShareInfo struct {
	OperatingAirlineInfo *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightShareInfoOperatingAirlineInfo `json:"operating_airline_info,omitempty" xml:"operating_airline_info,omitempty" type:"Struct"`
	OperatingFlightNo    *string                                                                                           `json:"operating_flight_no,omitempty" xml:"operating_flight_no,omitempty"`
}

func (s FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightShareInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightShareInfo) GoString() string {
	return s.String()
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightShareInfo) GetOperatingAirlineInfo() *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightShareInfoOperatingAirlineInfo {
	return s.OperatingAirlineInfo
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightShareInfo) GetOperatingFlightNo() *string {
	return s.OperatingFlightNo
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightShareInfo) SetOperatingAirlineInfo(v *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightShareInfoOperatingAirlineInfo) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightShareInfo {
	s.OperatingAirlineInfo = v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightShareInfo) SetOperatingFlightNo(v string) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightShareInfo {
	s.OperatingFlightNo = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightShareInfo) Validate() error {
	return dara.Validate(s)
}

type FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightShareInfoOperatingAirlineInfo struct {
	AirlineChineseName      *string `json:"airline_chinese_name,omitempty" xml:"airline_chinese_name,omitempty"`
	AirlineChineseShortName *string `json:"airline_chinese_short_name,omitempty" xml:"airline_chinese_short_name,omitempty"`
	AirlineCode             *string `json:"airline_code,omitempty" xml:"airline_code,omitempty"`
	AirlineIcon             *string `json:"airline_icon,omitempty" xml:"airline_icon,omitempty"`
	// example:
	//
	// false
	CheapFlight *bool `json:"cheap_flight,omitempty" xml:"cheap_flight,omitempty"`
}

func (s FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightShareInfoOperatingAirlineInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightShareInfoOperatingAirlineInfo) GoString() string {
	return s.String()
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightShareInfoOperatingAirlineInfo) GetAirlineChineseName() *string {
	return s.AirlineChineseName
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightShareInfoOperatingAirlineInfo) GetAirlineChineseShortName() *string {
	return s.AirlineChineseShortName
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightShareInfoOperatingAirlineInfo) GetAirlineCode() *string {
	return s.AirlineCode
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightShareInfoOperatingAirlineInfo) GetAirlineIcon() *string {
	return s.AirlineIcon
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightShareInfoOperatingAirlineInfo) GetCheapFlight() *bool {
	return s.CheapFlight
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightShareInfoOperatingAirlineInfo) SetAirlineChineseName(v string) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightShareInfoOperatingAirlineInfo {
	s.AirlineChineseName = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightShareInfoOperatingAirlineInfo) SetAirlineChineseShortName(v string) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightShareInfoOperatingAirlineInfo {
	s.AirlineChineseShortName = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightShareInfoOperatingAirlineInfo) SetAirlineCode(v string) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightShareInfoOperatingAirlineInfo {
	s.AirlineCode = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightShareInfoOperatingAirlineInfo) SetAirlineIcon(v string) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightShareInfoOperatingAirlineInfo {
	s.AirlineIcon = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightShareInfoOperatingAirlineInfo) SetCheapFlight(v bool) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightShareInfoOperatingAirlineInfo {
	s.CheapFlight = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightShareInfoOperatingAirlineInfo) Validate() error {
	return dara.Validate(s)
}

type FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightStopInfo struct {
	StopAirport  *string `json:"stop_airport,omitempty" xml:"stop_airport,omitempty"`
	StopArrTerm  *string `json:"stop_arr_term,omitempty" xml:"stop_arr_term,omitempty"`
	StopArrTime  *string `json:"stop_arr_time,omitempty" xml:"stop_arr_time,omitempty"`
	StopCityCode *string `json:"stop_city_code,omitempty" xml:"stop_city_code,omitempty"`
	StopCityName *string `json:"stop_city_name,omitempty" xml:"stop_city_name,omitempty"`
	StopDepTerm  *string `json:"stop_dep_term,omitempty" xml:"stop_dep_term,omitempty"`
	StopDepTime  *string `json:"stop_dep_time,omitempty" xml:"stop_dep_time,omitempty"`
}

func (s FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightStopInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightStopInfo) GoString() string {
	return s.String()
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightStopInfo) GetStopAirport() *string {
	return s.StopAirport
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightStopInfo) GetStopArrTerm() *string {
	return s.StopArrTerm
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightStopInfo) GetStopArrTime() *string {
	return s.StopArrTime
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightStopInfo) GetStopCityCode() *string {
	return s.StopCityCode
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightStopInfo) GetStopCityName() *string {
	return s.StopCityName
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightStopInfo) GetStopDepTerm() *string {
	return s.StopDepTerm
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightStopInfo) GetStopDepTime() *string {
	return s.StopDepTime
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightStopInfo) SetStopAirport(v string) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightStopInfo {
	s.StopAirport = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightStopInfo) SetStopArrTerm(v string) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightStopInfo {
	s.StopArrTerm = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightStopInfo) SetStopArrTime(v string) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightStopInfo {
	s.StopArrTime = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightStopInfo) SetStopCityCode(v string) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightStopInfo {
	s.StopCityCode = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightStopInfo) SetStopCityName(v string) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightStopInfo {
	s.StopCityName = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightStopInfo) SetStopDepTerm(v string) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightStopInfo {
	s.StopDepTerm = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightStopInfo) SetStopDepTime(v string) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightStopInfo {
	s.StopDepTime = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightStopInfo) Validate() error {
	return dara.Validate(s)
}

type FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightTransferInfo struct {
	TransferAirlineInfo *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightTransferInfoTransferAirlineInfo `json:"transfer_airline_info,omitempty" xml:"transfer_airline_info,omitempty" type:"Struct"`
	TransferCityCode    *string                                                                                             `json:"transfer_city_code,omitempty" xml:"transfer_city_code,omitempty"`
	TransferCityName    *string                                                                                             `json:"transfer_city_name,omitempty" xml:"transfer_city_name,omitempty"`
	TransferDepTime     *string                                                                                             `json:"transfer_dep_time,omitempty" xml:"transfer_dep_time,omitempty"`
	TransferFlightNo    *string                                                                                             `json:"transfer_flight_no,omitempty" xml:"transfer_flight_no,omitempty"`
	TransferFlightSize  *string                                                                                             `json:"transfer_flight_size,omitempty" xml:"transfer_flight_size,omitempty"`
	// example:
	//
	// false
	TransferShare *bool `json:"transfer_share,omitempty" xml:"transfer_share,omitempty"`
	// example:
	//
	// 60
	TransferStopTime *int32 `json:"transfer_stop_time,omitempty" xml:"transfer_stop_time,omitempty"`
}

func (s FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightTransferInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightTransferInfo) GoString() string {
	return s.String()
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightTransferInfo) GetTransferAirlineInfo() *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightTransferInfoTransferAirlineInfo {
	return s.TransferAirlineInfo
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightTransferInfo) GetTransferCityCode() *string {
	return s.TransferCityCode
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightTransferInfo) GetTransferCityName() *string {
	return s.TransferCityName
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightTransferInfo) GetTransferDepTime() *string {
	return s.TransferDepTime
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightTransferInfo) GetTransferFlightNo() *string {
	return s.TransferFlightNo
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightTransferInfo) GetTransferFlightSize() *string {
	return s.TransferFlightSize
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightTransferInfo) GetTransferShare() *bool {
	return s.TransferShare
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightTransferInfo) GetTransferStopTime() *int32 {
	return s.TransferStopTime
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightTransferInfo) SetTransferAirlineInfo(v *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightTransferInfoTransferAirlineInfo) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightTransferInfo {
	s.TransferAirlineInfo = v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightTransferInfo) SetTransferCityCode(v string) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightTransferInfo {
	s.TransferCityCode = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightTransferInfo) SetTransferCityName(v string) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightTransferInfo {
	s.TransferCityName = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightTransferInfo) SetTransferDepTime(v string) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightTransferInfo {
	s.TransferDepTime = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightTransferInfo) SetTransferFlightNo(v string) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightTransferInfo {
	s.TransferFlightNo = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightTransferInfo) SetTransferFlightSize(v string) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightTransferInfo {
	s.TransferFlightSize = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightTransferInfo) SetTransferShare(v bool) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightTransferInfo {
	s.TransferShare = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightTransferInfo) SetTransferStopTime(v int32) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightTransferInfo {
	s.TransferStopTime = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightTransferInfo) Validate() error {
	return dara.Validate(s)
}

type FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightTransferInfoTransferAirlineInfo struct {
	AirlineChineseName      *string `json:"airline_chinese_name,omitempty" xml:"airline_chinese_name,omitempty"`
	AirlineChineseShortName *string `json:"airline_chinese_short_name,omitempty" xml:"airline_chinese_short_name,omitempty"`
	AirlineCode             *string `json:"airline_code,omitempty" xml:"airline_code,omitempty"`
	AirlineIcon             *string `json:"airline_icon,omitempty" xml:"airline_icon,omitempty"`
	// example:
	//
	// false
	CheapFlight *bool `json:"cheap_flight,omitempty" xml:"cheap_flight,omitempty"`
}

func (s FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightTransferInfoTransferAirlineInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightTransferInfoTransferAirlineInfo) GoString() string {
	return s.String()
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightTransferInfoTransferAirlineInfo) GetAirlineChineseName() *string {
	return s.AirlineChineseName
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightTransferInfoTransferAirlineInfo) GetAirlineChineseShortName() *string {
	return s.AirlineChineseShortName
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightTransferInfoTransferAirlineInfo) GetAirlineCode() *string {
	return s.AirlineCode
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightTransferInfoTransferAirlineInfo) GetAirlineIcon() *string {
	return s.AirlineIcon
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightTransferInfoTransferAirlineInfo) GetCheapFlight() *bool {
	return s.CheapFlight
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightTransferInfoTransferAirlineInfo) SetAirlineChineseName(v string) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightTransferInfoTransferAirlineInfo {
	s.AirlineChineseName = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightTransferInfoTransferAirlineInfo) SetAirlineChineseShortName(v string) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightTransferInfoTransferAirlineInfo {
	s.AirlineChineseShortName = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightTransferInfoTransferAirlineInfo) SetAirlineCode(v string) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightTransferInfoTransferAirlineInfo {
	s.AirlineCode = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightTransferInfoTransferAirlineInfo) SetAirlineIcon(v string) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightTransferInfoTransferAirlineInfo {
	s.AirlineIcon = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightTransferInfoTransferAirlineInfo) SetCheapFlight(v bool) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightTransferInfoTransferAirlineInfo {
	s.CheapFlight = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListFlightTransferInfoTransferAirlineInfo) Validate() error {
	return dara.Validate(s)
}

type FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTO struct {
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
	// 12000
	InfantPrice *int32 `json:"infant_price,omitempty" xml:"infant_price,omitempty"`
	// example:
	//
	// 0
	InfantTax *int32 `json:"infant_tax,omitempty" xml:"infant_tax,omitempty"`
	// example:
	//
	// 12000
	InfantTotalPrice *int32 `json:"infant_total_price,omitempty" xml:"infant_total_price,omitempty"`
	// example:
	//
	// 1300
	OriginalAdultPrice *int32 `json:"original_adult_price,omitempty" xml:"original_adult_price,omitempty"`
	// example:
	//
	// 12300
	OriginalAdultTotalPrice *int32                                                                                       `json:"original_adult_total_price,omitempty" xml:"original_adult_total_price,omitempty"`
	ReShopPriceInfoDTO      *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTOReShopPriceInfoDTO `json:"re_shop_price_info_d_t_o,omitempty" xml:"re_shop_price_info_d_t_o,omitempty" type:"Struct"`
}

func (s FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTO) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTO) GoString() string {
	return s.String()
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTO) GetAdultPrice() *int32 {
	return s.AdultPrice
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTO) GetAdultTax() *int32 {
	return s.AdultTax
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTO) GetAdultTotalPrice() *int32 {
	return s.AdultTotalPrice
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTO) GetBeforeControlPrice() *int32 {
	return s.BeforeControlPrice
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTO) GetChildPrice() *int32 {
	return s.ChildPrice
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTO) GetChildTax() *int32 {
	return s.ChildTax
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTO) GetChildTotalPrice() *int32 {
	return s.ChildTotalPrice
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTO) GetInfantPrice() *int32 {
	return s.InfantPrice
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTO) GetInfantTax() *int32 {
	return s.InfantTax
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTO) GetInfantTotalPrice() *int32 {
	return s.InfantTotalPrice
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTO) GetOriginalAdultPrice() *int32 {
	return s.OriginalAdultPrice
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTO) GetOriginalAdultTotalPrice() *int32 {
	return s.OriginalAdultTotalPrice
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTO) GetReShopPriceInfoDTO() *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTOReShopPriceInfoDTO {
	return s.ReShopPriceInfoDTO
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTO) SetAdultPrice(v int32) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTO {
	s.AdultPrice = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTO) SetAdultTax(v int32) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTO {
	s.AdultTax = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTO) SetAdultTotalPrice(v int32) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTO {
	s.AdultTotalPrice = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTO) SetBeforeControlPrice(v int32) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTO {
	s.BeforeControlPrice = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTO) SetChildPrice(v int32) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTO {
	s.ChildPrice = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTO) SetChildTax(v int32) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTO {
	s.ChildTax = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTO) SetChildTotalPrice(v int32) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTO {
	s.ChildTotalPrice = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTO) SetInfantPrice(v int32) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTO {
	s.InfantPrice = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTO) SetInfantTax(v int32) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTO {
	s.InfantTax = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTO) SetInfantTotalPrice(v int32) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTO {
	s.InfantTotalPrice = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTO) SetOriginalAdultPrice(v int32) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTO {
	s.OriginalAdultPrice = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTO) SetOriginalAdultTotalPrice(v int32) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTO {
	s.OriginalAdultTotalPrice = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTO) SetReShopPriceInfoDTO(v *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTOReShopPriceInfoDTO) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTO {
	s.ReShopPriceInfoDTO = v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTO) Validate() error {
	return dara.Validate(s)
}

type FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTOReShopPriceInfoDTO struct {
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

func (s FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTOReShopPriceInfoDTO) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTOReShopPriceInfoDTO) GoString() string {
	return s.String()
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTOReShopPriceInfoDTO) GetReShopAdultChangeFee() *int32 {
	return s.ReShopAdultChangeFee
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTOReShopPriceInfoDTO) GetReShopAdultPrice() *int32 {
	return s.ReShopAdultPrice
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTOReShopPriceInfoDTO) GetReShopAdultPriceGap() *int32 {
	return s.ReShopAdultPriceGap
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTOReShopPriceInfoDTO) GetReShopChildChangeFee() *int32 {
	return s.ReShopChildChangeFee
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTOReShopPriceInfoDTO) GetReShopChildPrice() *int32 {
	return s.ReShopChildPrice
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTOReShopPriceInfoDTO) GetReShopChildPriceGap() *int32 {
	return s.ReShopChildPriceGap
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTOReShopPriceInfoDTO) GetReShopInfChangeFee() *int32 {
	return s.ReShopInfChangeFee
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTOReShopPriceInfoDTO) GetReShopInfPrice() *int32 {
	return s.ReShopInfPrice
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTOReShopPriceInfoDTO) GetReShopInfPriceGap() *int32 {
	return s.ReShopInfPriceGap
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTOReShopPriceInfoDTO) SetReShopAdultChangeFee(v int32) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTOReShopPriceInfoDTO {
	s.ReShopAdultChangeFee = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTOReShopPriceInfoDTO) SetReShopAdultPrice(v int32) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTOReShopPriceInfoDTO {
	s.ReShopAdultPrice = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTOReShopPriceInfoDTO) SetReShopAdultPriceGap(v int32) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTOReShopPriceInfoDTO {
	s.ReShopAdultPriceGap = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTOReShopPriceInfoDTO) SetReShopChildChangeFee(v int32) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTOReShopPriceInfoDTO {
	s.ReShopChildChangeFee = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTOReShopPriceInfoDTO) SetReShopChildPrice(v int32) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTOReShopPriceInfoDTO {
	s.ReShopChildPrice = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTOReShopPriceInfoDTO) SetReShopChildPriceGap(v int32) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTOReShopPriceInfoDTO {
	s.ReShopChildPriceGap = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTOReShopPriceInfoDTO) SetReShopInfChangeFee(v int32) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTOReShopPriceInfoDTO {
	s.ReShopInfChangeFee = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTOReShopPriceInfoDTO) SetReShopInfPrice(v int32) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTOReShopPriceInfoDTO {
	s.ReShopInfPrice = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTOReShopPriceInfoDTO) SetReShopInfPriceGap(v int32) *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTOReShopPriceInfoDTO {
	s.ReShopInfPriceGap = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleDirectFlightListPriceInfoDTOReShopPriceInfoDTO) Validate() error {
	return dara.Validate(s)
}

type FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList struct {
	AirlineInfo    *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListAirlineInfo    `json:"airline_info,omitempty" xml:"airline_info,omitempty" type:"Struct"`
	ArrAirportInfo *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListArrAirportInfo `json:"arr_airport_info,omitempty" xml:"arr_airport_info,omitempty" type:"Struct"`
	ArrCityCode    *string                                                                        `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	ArrTime        *string                                                                        `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	CabinClass     *string                                                                        `json:"cabinClass,omitempty" xml:"cabinClass,omitempty"`
	CabinClassName *string                                                                        `json:"cabinClassName,omitempty" xml:"cabinClassName,omitempty"`
	DepAirportInfo *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListDepAirportInfo `json:"dep_airport_info,omitempty" xml:"dep_airport_info,omitempty" type:"Struct"`
	DepCityCode    *string                                                                        `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	DepTime        *string                                                                        `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	// example:
	//
	// 240
	Duration           *int32                                                                             `json:"duration,omitempty" xml:"duration,omitempty"`
	FlightNo           *string                                                                            `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
	FlightShareInfo    *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightShareInfo    `json:"flight_share_info,omitempty" xml:"flight_share_info,omitempty" type:"Struct"`
	FlightSize         *string                                                                            `json:"flight_size,omitempty" xml:"flight_size,omitempty"`
	FlightStopInfo     *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightStopInfo     `json:"flight_stop_info,omitempty" xml:"flight_stop_info,omitempty" type:"Struct"`
	FlightTransferInfo *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightTransferInfo `json:"flight_transfer_info,omitempty" xml:"flight_transfer_info,omitempty" type:"Struct"`
	FlightType         *string                                                                            `json:"flight_type,omitempty" xml:"flight_type,omitempty"`
	// example:
	//
	// 0
	JourneySeq *int32 `json:"journey_seq,omitempty" xml:"journey_seq,omitempty"`
	// example:
	//
	// 7
	LeftNum      *string                                                                      `json:"left_num,omitempty" xml:"left_num,omitempty"`
	Manufacturer *string                                                                      `json:"manufacturer,omitempty" xml:"manufacturer,omitempty"`
	MealDesc     *string                                                                      `json:"meal_desc,omitempty" xml:"meal_desc,omitempty"`
	PriceInfoDTO *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTO `json:"price_info_d_t_o,omitempty" xml:"price_info_d_t_o,omitempty" type:"Struct"`
	// example:
	//
	// 0
	SegmentSeq *int32 `json:"segment_seq,omitempty" xml:"segment_seq,omitempty"`
	// example:
	//
	// false
	Share           *bool   `json:"share,omitempty" xml:"share,omitempty"`
	ShortFlightSize *string `json:"short_flight_size,omitempty" xml:"short_flight_size,omitempty"`
	SpanDay         *string `json:"span_day,omitempty" xml:"span_day,omitempty"`
	// example:
	//
	// false
	Stop *bool `json:"stop,omitempty" xml:"stop,omitempty"`
	// example:
	//
	// false
	Transfer *bool `json:"transfer,omitempty" xml:"transfer,omitempty"`
}

func (s FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList) GoString() string {
	return s.String()
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList) GetAirlineInfo() *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListAirlineInfo {
	return s.AirlineInfo
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList) GetArrAirportInfo() *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListArrAirportInfo {
	return s.ArrAirportInfo
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList) GetArrTime() *string {
	return s.ArrTime
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList) GetCabinClass() *string {
	return s.CabinClass
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList) GetCabinClassName() *string {
	return s.CabinClassName
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList) GetDepAirportInfo() *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListDepAirportInfo {
	return s.DepAirportInfo
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList) GetDepTime() *string {
	return s.DepTime
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList) GetDuration() *int32 {
	return s.Duration
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList) GetFlightNo() *string {
	return s.FlightNo
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList) GetFlightShareInfo() *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightShareInfo {
	return s.FlightShareInfo
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList) GetFlightSize() *string {
	return s.FlightSize
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList) GetFlightStopInfo() *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightStopInfo {
	return s.FlightStopInfo
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList) GetFlightTransferInfo() *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightTransferInfo {
	return s.FlightTransferInfo
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList) GetFlightType() *string {
	return s.FlightType
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList) GetJourneySeq() *int32 {
	return s.JourneySeq
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList) GetLeftNum() *string {
	return s.LeftNum
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList) GetManufacturer() *string {
	return s.Manufacturer
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList) GetMealDesc() *string {
	return s.MealDesc
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList) GetPriceInfoDTO() *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTO {
	return s.PriceInfoDTO
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList) GetSegmentSeq() *int32 {
	return s.SegmentSeq
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList) GetShare() *bool {
	return s.Share
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList) GetShortFlightSize() *string {
	return s.ShortFlightSize
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList) GetSpanDay() *string {
	return s.SpanDay
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList) GetStop() *bool {
	return s.Stop
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList) GetTransfer() *bool {
	return s.Transfer
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList) SetAirlineInfo(v *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListAirlineInfo) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList {
	s.AirlineInfo = v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList) SetArrAirportInfo(v *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListArrAirportInfo) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList {
	s.ArrAirportInfo = v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList) SetArrCityCode(v string) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList {
	s.ArrCityCode = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList) SetArrTime(v string) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList {
	s.ArrTime = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList) SetCabinClass(v string) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList {
	s.CabinClass = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList) SetCabinClassName(v string) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList {
	s.CabinClassName = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList) SetDepAirportInfo(v *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListDepAirportInfo) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList {
	s.DepAirportInfo = v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList) SetDepCityCode(v string) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList {
	s.DepCityCode = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList) SetDepTime(v string) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList {
	s.DepTime = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList) SetDuration(v int32) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList {
	s.Duration = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList) SetFlightNo(v string) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList {
	s.FlightNo = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList) SetFlightShareInfo(v *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightShareInfo) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList {
	s.FlightShareInfo = v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList) SetFlightSize(v string) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList {
	s.FlightSize = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList) SetFlightStopInfo(v *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightStopInfo) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList {
	s.FlightStopInfo = v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList) SetFlightTransferInfo(v *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightTransferInfo) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList {
	s.FlightTransferInfo = v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList) SetFlightType(v string) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList {
	s.FlightType = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList) SetJourneySeq(v int32) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList {
	s.JourneySeq = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList) SetLeftNum(v string) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList {
	s.LeftNum = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList) SetManufacturer(v string) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList {
	s.Manufacturer = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList) SetMealDesc(v string) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList {
	s.MealDesc = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList) SetPriceInfoDTO(v *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTO) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList {
	s.PriceInfoDTO = v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList) SetSegmentSeq(v int32) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList {
	s.SegmentSeq = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList) SetShare(v bool) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList {
	s.Share = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList) SetShortFlightSize(v string) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList {
	s.ShortFlightSize = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList) SetSpanDay(v string) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList {
	s.SpanDay = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList) SetStop(v bool) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList {
	s.Stop = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList) SetTransfer(v bool) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList {
	s.Transfer = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightList) Validate() error {
	return dara.Validate(s)
}

type FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListAirlineInfo struct {
	AirlineChineseName      *string `json:"airline_chinese_name,omitempty" xml:"airline_chinese_name,omitempty"`
	AirlineChineseShortName *string `json:"airline_chinese_short_name,omitempty" xml:"airline_chinese_short_name,omitempty"`
	AirlineCode             *string `json:"airline_code,omitempty" xml:"airline_code,omitempty"`
	AirlineIcon             *string `json:"airline_icon,omitempty" xml:"airline_icon,omitempty"`
	CheapFlight             *bool   `json:"cheap_flight,omitempty" xml:"cheap_flight,omitempty"`
}

func (s FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListAirlineInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListAirlineInfo) GoString() string {
	return s.String()
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListAirlineInfo) GetAirlineChineseName() *string {
	return s.AirlineChineseName
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListAirlineInfo) GetAirlineChineseShortName() *string {
	return s.AirlineChineseShortName
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListAirlineInfo) GetAirlineCode() *string {
	return s.AirlineCode
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListAirlineInfo) GetAirlineIcon() *string {
	return s.AirlineIcon
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListAirlineInfo) GetCheapFlight() *bool {
	return s.CheapFlight
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListAirlineInfo) SetAirlineChineseName(v string) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListAirlineInfo {
	s.AirlineChineseName = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListAirlineInfo) SetAirlineChineseShortName(v string) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListAirlineInfo {
	s.AirlineChineseShortName = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListAirlineInfo) SetAirlineCode(v string) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListAirlineInfo {
	s.AirlineCode = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListAirlineInfo) SetAirlineIcon(v string) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListAirlineInfo {
	s.AirlineIcon = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListAirlineInfo) SetCheapFlight(v bool) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListAirlineInfo {
	s.CheapFlight = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListAirlineInfo) Validate() error {
	return dara.Validate(s)
}

type FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListArrAirportInfo struct {
	AirportCode      *string `json:"airport_code,omitempty" xml:"airport_code,omitempty"`
	AirportName      *string `json:"airport_name,omitempty" xml:"airport_name,omitempty"`
	AirportShortName *string `json:"airport_short_name,omitempty" xml:"airport_short_name,omitempty"`
	Terminal         *string `json:"terminal,omitempty" xml:"terminal,omitempty"`
}

func (s FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListArrAirportInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListArrAirportInfo) GoString() string {
	return s.String()
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListArrAirportInfo) GetAirportCode() *string {
	return s.AirportCode
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListArrAirportInfo) GetAirportName() *string {
	return s.AirportName
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListArrAirportInfo) GetAirportShortName() *string {
	return s.AirportShortName
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListArrAirportInfo) GetTerminal() *string {
	return s.Terminal
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListArrAirportInfo) SetAirportCode(v string) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListArrAirportInfo {
	s.AirportCode = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListArrAirportInfo) SetAirportName(v string) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListArrAirportInfo {
	s.AirportName = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListArrAirportInfo) SetAirportShortName(v string) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListArrAirportInfo {
	s.AirportShortName = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListArrAirportInfo) SetTerminal(v string) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListArrAirportInfo {
	s.Terminal = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListArrAirportInfo) Validate() error {
	return dara.Validate(s)
}

type FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListDepAirportInfo struct {
	AirportCode      *string `json:"airport_code,omitempty" xml:"airport_code,omitempty"`
	AirportName      *string `json:"airport_name,omitempty" xml:"airport_name,omitempty"`
	AirportShortName *string `json:"airport_short_name,omitempty" xml:"airport_short_name,omitempty"`
	Terminal         *string `json:"terminal,omitempty" xml:"terminal,omitempty"`
}

func (s FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListDepAirportInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListDepAirportInfo) GoString() string {
	return s.String()
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListDepAirportInfo) GetAirportCode() *string {
	return s.AirportCode
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListDepAirportInfo) GetAirportName() *string {
	return s.AirportName
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListDepAirportInfo) GetAirportShortName() *string {
	return s.AirportShortName
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListDepAirportInfo) GetTerminal() *string {
	return s.Terminal
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListDepAirportInfo) SetAirportCode(v string) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListDepAirportInfo {
	s.AirportCode = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListDepAirportInfo) SetAirportName(v string) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListDepAirportInfo {
	s.AirportName = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListDepAirportInfo) SetAirportShortName(v string) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListDepAirportInfo {
	s.AirportShortName = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListDepAirportInfo) SetTerminal(v string) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListDepAirportInfo {
	s.Terminal = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListDepAirportInfo) Validate() error {
	return dara.Validate(s)
}

type FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightShareInfo struct {
	OperatingAirlineInfo *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightShareInfoOperatingAirlineInfo `json:"operating_airline_info,omitempty" xml:"operating_airline_info,omitempty" type:"Struct"`
	OperatingFlightNo    *string                                                                                             `json:"operating_flight_no,omitempty" xml:"operating_flight_no,omitempty"`
}

func (s FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightShareInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightShareInfo) GoString() string {
	return s.String()
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightShareInfo) GetOperatingAirlineInfo() *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightShareInfoOperatingAirlineInfo {
	return s.OperatingAirlineInfo
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightShareInfo) GetOperatingFlightNo() *string {
	return s.OperatingFlightNo
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightShareInfo) SetOperatingAirlineInfo(v *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightShareInfoOperatingAirlineInfo) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightShareInfo {
	s.OperatingAirlineInfo = v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightShareInfo) SetOperatingFlightNo(v string) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightShareInfo {
	s.OperatingFlightNo = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightShareInfo) Validate() error {
	return dara.Validate(s)
}

type FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightShareInfoOperatingAirlineInfo struct {
	AirlineChineseName      *string `json:"airline_chinese_name,omitempty" xml:"airline_chinese_name,omitempty"`
	AirlineChineseShortName *string `json:"airline_chinese_short_name,omitempty" xml:"airline_chinese_short_name,omitempty"`
	AirlineCode             *string `json:"airline_code,omitempty" xml:"airline_code,omitempty"`
	AirlineIcon             *string `json:"airline_icon,omitempty" xml:"airline_icon,omitempty"`
	// example:
	//
	// false
	CheapFlight *bool `json:"cheap_flight,omitempty" xml:"cheap_flight,omitempty"`
}

func (s FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightShareInfoOperatingAirlineInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightShareInfoOperatingAirlineInfo) GoString() string {
	return s.String()
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightShareInfoOperatingAirlineInfo) GetAirlineChineseName() *string {
	return s.AirlineChineseName
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightShareInfoOperatingAirlineInfo) GetAirlineChineseShortName() *string {
	return s.AirlineChineseShortName
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightShareInfoOperatingAirlineInfo) GetAirlineCode() *string {
	return s.AirlineCode
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightShareInfoOperatingAirlineInfo) GetAirlineIcon() *string {
	return s.AirlineIcon
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightShareInfoOperatingAirlineInfo) GetCheapFlight() *bool {
	return s.CheapFlight
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightShareInfoOperatingAirlineInfo) SetAirlineChineseName(v string) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightShareInfoOperatingAirlineInfo {
	s.AirlineChineseName = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightShareInfoOperatingAirlineInfo) SetAirlineChineseShortName(v string) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightShareInfoOperatingAirlineInfo {
	s.AirlineChineseShortName = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightShareInfoOperatingAirlineInfo) SetAirlineCode(v string) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightShareInfoOperatingAirlineInfo {
	s.AirlineCode = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightShareInfoOperatingAirlineInfo) SetAirlineIcon(v string) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightShareInfoOperatingAirlineInfo {
	s.AirlineIcon = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightShareInfoOperatingAirlineInfo) SetCheapFlight(v bool) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightShareInfoOperatingAirlineInfo {
	s.CheapFlight = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightShareInfoOperatingAirlineInfo) Validate() error {
	return dara.Validate(s)
}

type FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightStopInfo struct {
	StopAirport  *string `json:"stop_airport,omitempty" xml:"stop_airport,omitempty"`
	StopArrTerm  *string `json:"stop_arr_term,omitempty" xml:"stop_arr_term,omitempty"`
	StopArrTime  *string `json:"stop_arr_time,omitempty" xml:"stop_arr_time,omitempty"`
	StopCityCode *string `json:"stop_city_code,omitempty" xml:"stop_city_code,omitempty"`
	StopCityName *string `json:"stop_city_name,omitempty" xml:"stop_city_name,omitempty"`
	StopDepTerm  *string `json:"stop_dep_term,omitempty" xml:"stop_dep_term,omitempty"`
	StopDepTime  *string `json:"stop_dep_time,omitempty" xml:"stop_dep_time,omitempty"`
}

func (s FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightStopInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightStopInfo) GoString() string {
	return s.String()
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightStopInfo) GetStopAirport() *string {
	return s.StopAirport
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightStopInfo) GetStopArrTerm() *string {
	return s.StopArrTerm
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightStopInfo) GetStopArrTime() *string {
	return s.StopArrTime
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightStopInfo) GetStopCityCode() *string {
	return s.StopCityCode
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightStopInfo) GetStopCityName() *string {
	return s.StopCityName
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightStopInfo) GetStopDepTerm() *string {
	return s.StopDepTerm
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightStopInfo) GetStopDepTime() *string {
	return s.StopDepTime
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightStopInfo) SetStopAirport(v string) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightStopInfo {
	s.StopAirport = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightStopInfo) SetStopArrTerm(v string) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightStopInfo {
	s.StopArrTerm = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightStopInfo) SetStopArrTime(v string) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightStopInfo {
	s.StopArrTime = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightStopInfo) SetStopCityCode(v string) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightStopInfo {
	s.StopCityCode = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightStopInfo) SetStopCityName(v string) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightStopInfo {
	s.StopCityName = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightStopInfo) SetStopDepTerm(v string) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightStopInfo {
	s.StopDepTerm = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightStopInfo) SetStopDepTime(v string) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightStopInfo {
	s.StopDepTime = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightStopInfo) Validate() error {
	return dara.Validate(s)
}

type FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightTransferInfo struct {
	TransferAirlineInfo *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightTransferInfoTransferAirlineInfo `json:"transfer_airline_info,omitempty" xml:"transfer_airline_info,omitempty" type:"Struct"`
	TransferCityCode    *string                                                                                               `json:"transfer_city_code,omitempty" xml:"transfer_city_code,omitempty"`
	TransferCityName    *string                                                                                               `json:"transfer_city_name,omitempty" xml:"transfer_city_name,omitempty"`
	TransferDepTime     *string                                                                                               `json:"transfer_dep_time,omitempty" xml:"transfer_dep_time,omitempty"`
	TransferFlightNo    *string                                                                                               `json:"transfer_flight_no,omitempty" xml:"transfer_flight_no,omitempty"`
	TransferFlightSize  *string                                                                                               `json:"transfer_flight_size,omitempty" xml:"transfer_flight_size,omitempty"`
	// example:
	//
	// false
	TransferShare *bool `json:"transfer_share,omitempty" xml:"transfer_share,omitempty"`
	// example:
	//
	// 20
	TransferStopTime *int32 `json:"transfer_stop_time,omitempty" xml:"transfer_stop_time,omitempty"`
}

func (s FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightTransferInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightTransferInfo) GoString() string {
	return s.String()
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightTransferInfo) GetTransferAirlineInfo() *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightTransferInfoTransferAirlineInfo {
	return s.TransferAirlineInfo
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightTransferInfo) GetTransferCityCode() *string {
	return s.TransferCityCode
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightTransferInfo) GetTransferCityName() *string {
	return s.TransferCityName
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightTransferInfo) GetTransferDepTime() *string {
	return s.TransferDepTime
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightTransferInfo) GetTransferFlightNo() *string {
	return s.TransferFlightNo
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightTransferInfo) GetTransferFlightSize() *string {
	return s.TransferFlightSize
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightTransferInfo) GetTransferShare() *bool {
	return s.TransferShare
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightTransferInfo) GetTransferStopTime() *int32 {
	return s.TransferStopTime
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightTransferInfo) SetTransferAirlineInfo(v *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightTransferInfoTransferAirlineInfo) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightTransferInfo {
	s.TransferAirlineInfo = v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightTransferInfo) SetTransferCityCode(v string) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightTransferInfo {
	s.TransferCityCode = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightTransferInfo) SetTransferCityName(v string) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightTransferInfo {
	s.TransferCityName = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightTransferInfo) SetTransferDepTime(v string) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightTransferInfo {
	s.TransferDepTime = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightTransferInfo) SetTransferFlightNo(v string) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightTransferInfo {
	s.TransferFlightNo = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightTransferInfo) SetTransferFlightSize(v string) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightTransferInfo {
	s.TransferFlightSize = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightTransferInfo) SetTransferShare(v bool) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightTransferInfo {
	s.TransferShare = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightTransferInfo) SetTransferStopTime(v int32) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightTransferInfo {
	s.TransferStopTime = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightTransferInfo) Validate() error {
	return dara.Validate(s)
}

type FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightTransferInfoTransferAirlineInfo struct {
	AirlineChineseName      *string `json:"airline_chinese_name,omitempty" xml:"airline_chinese_name,omitempty"`
	AirlineChineseShortName *string `json:"airline_chinese_short_name,omitempty" xml:"airline_chinese_short_name,omitempty"`
	AirlineCode             *string `json:"airline_code,omitempty" xml:"airline_code,omitempty"`
	AirlineIcon             *string `json:"airline_icon,omitempty" xml:"airline_icon,omitempty"`
	// example:
	//
	// false
	CheapFlight *bool `json:"cheap_flight,omitempty" xml:"cheap_flight,omitempty"`
}

func (s FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightTransferInfoTransferAirlineInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightTransferInfoTransferAirlineInfo) GoString() string {
	return s.String()
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightTransferInfoTransferAirlineInfo) GetAirlineChineseName() *string {
	return s.AirlineChineseName
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightTransferInfoTransferAirlineInfo) GetAirlineChineseShortName() *string {
	return s.AirlineChineseShortName
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightTransferInfoTransferAirlineInfo) GetAirlineCode() *string {
	return s.AirlineCode
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightTransferInfoTransferAirlineInfo) GetAirlineIcon() *string {
	return s.AirlineIcon
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightTransferInfoTransferAirlineInfo) GetCheapFlight() *bool {
	return s.CheapFlight
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightTransferInfoTransferAirlineInfo) SetAirlineChineseName(v string) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightTransferInfoTransferAirlineInfo {
	s.AirlineChineseName = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightTransferInfoTransferAirlineInfo) SetAirlineChineseShortName(v string) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightTransferInfoTransferAirlineInfo {
	s.AirlineChineseShortName = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightTransferInfoTransferAirlineInfo) SetAirlineCode(v string) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightTransferInfoTransferAirlineInfo {
	s.AirlineCode = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightTransferInfoTransferAirlineInfo) SetAirlineIcon(v string) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightTransferInfoTransferAirlineInfo {
	s.AirlineIcon = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightTransferInfoTransferAirlineInfo) SetCheapFlight(v bool) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightTransferInfoTransferAirlineInfo {
	s.CheapFlight = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListFlightTransferInfoTransferAirlineInfo) Validate() error {
	return dara.Validate(s)
}

type FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTO struct {
	// example:
	//
	// 1000
	AdultPrice *int32 `json:"adult_price,omitempty" xml:"adult_price,omitempty"`
	// example:
	//
	// 1000
	AdultTax *int32 `json:"adult_tax,omitempty" xml:"adult_tax,omitempty"`
	// example:
	//
	// 1000
	AdultTotalPrice *int32 `json:"adult_total_price,omitempty" xml:"adult_total_price,omitempty"`
	// example:
	//
	// 1000
	BeforeControlPrice *int32 `json:"before_control_price,omitempty" xml:"before_control_price,omitempty"`
	// example:
	//
	// 1000
	ChildPrice *int32 `json:"child_price,omitempty" xml:"child_price,omitempty"`
	// example:
	//
	// 1000
	ChildTax *int32 `json:"child_tax,omitempty" xml:"child_tax,omitempty"`
	// example:
	//
	// 1000
	ChildTotalPrice *int32 `json:"child_total_price,omitempty" xml:"child_total_price,omitempty"`
	// example:
	//
	// 1000
	InfantPrice *int32 `json:"infant_price,omitempty" xml:"infant_price,omitempty"`
	// example:
	//
	// 1000
	InfantTax *int32 `json:"infant_tax,omitempty" xml:"infant_tax,omitempty"`
	// example:
	//
	// 1000
	InfantTotalPrice *int32 `json:"infant_total_price,omitempty" xml:"infant_total_price,omitempty"`
	// example:
	//
	// 1000
	OriginalAdultPrice *int32 `json:"original_adult_price,omitempty" xml:"original_adult_price,omitempty"`
	// example:
	//
	// 1000
	OriginalAdultTotalPrice *int32                                                                                         `json:"original_adult_total_price,omitempty" xml:"original_adult_total_price,omitempty"`
	ReShopPriceInfoDTO      *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTOReShopPriceInfoDTO `json:"re_shop_price_info_d_t_o,omitempty" xml:"re_shop_price_info_d_t_o,omitempty" type:"Struct"`
}

func (s FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTO) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTO) GoString() string {
	return s.String()
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTO) GetAdultPrice() *int32 {
	return s.AdultPrice
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTO) GetAdultTax() *int32 {
	return s.AdultTax
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTO) GetAdultTotalPrice() *int32 {
	return s.AdultTotalPrice
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTO) GetBeforeControlPrice() *int32 {
	return s.BeforeControlPrice
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTO) GetChildPrice() *int32 {
	return s.ChildPrice
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTO) GetChildTax() *int32 {
	return s.ChildTax
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTO) GetChildTotalPrice() *int32 {
	return s.ChildTotalPrice
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTO) GetInfantPrice() *int32 {
	return s.InfantPrice
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTO) GetInfantTax() *int32 {
	return s.InfantTax
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTO) GetInfantTotalPrice() *int32 {
	return s.InfantTotalPrice
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTO) GetOriginalAdultPrice() *int32 {
	return s.OriginalAdultPrice
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTO) GetOriginalAdultTotalPrice() *int32 {
	return s.OriginalAdultTotalPrice
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTO) GetReShopPriceInfoDTO() *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTOReShopPriceInfoDTO {
	return s.ReShopPriceInfoDTO
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTO) SetAdultPrice(v int32) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTO {
	s.AdultPrice = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTO) SetAdultTax(v int32) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTO {
	s.AdultTax = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTO) SetAdultTotalPrice(v int32) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTO {
	s.AdultTotalPrice = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTO) SetBeforeControlPrice(v int32) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTO {
	s.BeforeControlPrice = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTO) SetChildPrice(v int32) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTO {
	s.ChildPrice = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTO) SetChildTax(v int32) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTO {
	s.ChildTax = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTO) SetChildTotalPrice(v int32) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTO {
	s.ChildTotalPrice = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTO) SetInfantPrice(v int32) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTO {
	s.InfantPrice = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTO) SetInfantTax(v int32) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTO {
	s.InfantTax = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTO) SetInfantTotalPrice(v int32) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTO {
	s.InfantTotalPrice = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTO) SetOriginalAdultPrice(v int32) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTO {
	s.OriginalAdultPrice = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTO) SetOriginalAdultTotalPrice(v int32) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTO {
	s.OriginalAdultTotalPrice = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTO) SetReShopPriceInfoDTO(v *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTOReShopPriceInfoDTO) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTO {
	s.ReShopPriceInfoDTO = v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTO) Validate() error {
	return dara.Validate(s)
}

type FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTOReShopPriceInfoDTO struct {
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

func (s FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTOReShopPriceInfoDTO) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTOReShopPriceInfoDTO) GoString() string {
	return s.String()
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTOReShopPriceInfoDTO) GetReShopAdultChangeFee() *int32 {
	return s.ReShopAdultChangeFee
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTOReShopPriceInfoDTO) GetReShopAdultPrice() *int32 {
	return s.ReShopAdultPrice
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTOReShopPriceInfoDTO) GetReShopAdultPriceGap() *int32 {
	return s.ReShopAdultPriceGap
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTOReShopPriceInfoDTO) GetReShopChildChangeFee() *int32 {
	return s.ReShopChildChangeFee
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTOReShopPriceInfoDTO) GetReShopChildPrice() *int32 {
	return s.ReShopChildPrice
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTOReShopPriceInfoDTO) GetReShopChildPriceGap() *int32 {
	return s.ReShopChildPriceGap
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTOReShopPriceInfoDTO) GetReShopInfChangeFee() *int32 {
	return s.ReShopInfChangeFee
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTOReShopPriceInfoDTO) GetReShopInfPrice() *int32 {
	return s.ReShopInfPrice
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTOReShopPriceInfoDTO) GetReShopInfPriceGap() *int32 {
	return s.ReShopInfPriceGap
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTOReShopPriceInfoDTO) SetReShopAdultChangeFee(v int32) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTOReShopPriceInfoDTO {
	s.ReShopAdultChangeFee = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTOReShopPriceInfoDTO) SetReShopAdultPrice(v int32) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTOReShopPriceInfoDTO {
	s.ReShopAdultPrice = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTOReShopPriceInfoDTO) SetReShopAdultPriceGap(v int32) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTOReShopPriceInfoDTO {
	s.ReShopAdultPriceGap = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTOReShopPriceInfoDTO) SetReShopChildChangeFee(v int32) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTOReShopPriceInfoDTO {
	s.ReShopChildChangeFee = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTOReShopPriceInfoDTO) SetReShopChildPrice(v int32) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTOReShopPriceInfoDTO {
	s.ReShopChildPrice = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTOReShopPriceInfoDTO) SetReShopChildPriceGap(v int32) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTOReShopPriceInfoDTO {
	s.ReShopChildPriceGap = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTOReShopPriceInfoDTO) SetReShopInfChangeFee(v int32) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTOReShopPriceInfoDTO {
	s.ReShopInfChangeFee = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTOReShopPriceInfoDTO) SetReShopInfPrice(v int32) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTOReShopPriceInfoDTO {
	s.ReShopInfPrice = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTOReShopPriceInfoDTO) SetReShopInfPriceGap(v int32) *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTOReShopPriceInfoDTO {
	s.ReShopInfPriceGap = &v
	return s
}

func (s *FlightModifyListingSearchV2ResponseBodyModuleTransferFlightListPriceInfoDTOReShopPriceInfoDTO) Validate() error {
	return dara.Validate(s)
}
