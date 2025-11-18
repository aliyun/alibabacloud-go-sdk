// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SearchResponseBody
	GetRequestId() *string
	SetData(v *SearchResponseBodyData) *SearchResponseBody
	GetData() *SearchResponseBodyData
	SetErrorCode(v string) *SearchResponseBody
	GetErrorCode() *string
	SetErrorData(v interface{}) *SearchResponseBody
	GetErrorData() interface{}
	SetErrorMsg(v string) *SearchResponseBody
	GetErrorMsg() *string
	SetStatus(v int32) *SearchResponseBody
	GetStatus() *int32
	SetSuccess(v bool) *SearchResponseBody
	GetSuccess() *bool
}

type SearchResponseBody struct {
	// request ID
	//
	// example:
	//
	// 2236993B-7BE7-5F92-B179-21FF08570165
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// data
	Data *SearchResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// error code
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// error data
	//
	// example:
	//
	// {}
	ErrorData interface{} `json:"error_data,omitempty" xml:"error_data,omitempty"`
	// error message
	//
	// example:
	//
	// ""
	ErrorMsg *string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// http request has been processed successfully，status code is 200
	//
	// example:
	//
	// 200
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// true represents success, false represents failure
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SearchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchResponseBody) GoString() string {
	return s.String()
}

func (s *SearchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchResponseBody) GetData() *SearchResponseBodyData {
	return s.Data
}

func (s *SearchResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *SearchResponseBody) GetErrorData() interface{} {
	return s.ErrorData
}

func (s *SearchResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *SearchResponseBody) GetStatus() *int32 {
	return s.Status
}

func (s *SearchResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SearchResponseBody) SetRequestId(v string) *SearchResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchResponseBody) SetData(v *SearchResponseBodyData) *SearchResponseBody {
	s.Data = v
	return s
}

func (s *SearchResponseBody) SetErrorCode(v string) *SearchResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SearchResponseBody) SetErrorData(v interface{}) *SearchResponseBody {
	s.ErrorData = v
	return s
}

func (s *SearchResponseBody) SetErrorMsg(v string) *SearchResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *SearchResponseBody) SetStatus(v int32) *SearchResponseBody {
	s.Status = &v
	return s
}

func (s *SearchResponseBody) SetSuccess(v bool) *SearchResponseBody {
	s.Success = &v
	return s
}

func (s *SearchResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchResponseBodyData struct {
	// solution list
	SolutionList []*SearchResponseBodyDataSolutionList `json:"solution_list,omitempty" xml:"solution_list,omitempty" type:"Repeated"`
}

func (s SearchResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SearchResponseBodyData) GoString() string {
	return s.String()
}

func (s *SearchResponseBodyData) GetSolutionList() []*SearchResponseBodyDataSolutionList {
	return s.SolutionList
}

func (s *SearchResponseBodyData) SetSolutionList(v []*SearchResponseBodyDataSolutionList) *SearchResponseBodyData {
	s.SolutionList = v
	return s
}

func (s *SearchResponseBodyData) Validate() error {
	if s.SolutionList != nil {
		for _, item := range s.SolutionList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SearchResponseBodyDataSolutionList struct {
	// adult fare
	//
	// example:
	//
	// 600
	AdultPrice *float64 `json:"adult_price,omitempty" xml:"adult_price,omitempty"`
	// adult tax
	//
	// example:
	//
	// 11
	AdultTax *float64 `json:"adult_tax,omitempty" xml:"adult_tax,omitempty"`
	// child fare
	//
	// example:
	//
	// 500
	ChildPrice *float64 `json:"child_price,omitempty" xml:"child_price,omitempty"`
	// child tax
	//
	// example:
	//
	// 10
	ChildTax *float64 `json:"child_tax,omitempty" xml:"child_tax,omitempty"`
	// infant fare
	//
	// example:
	//
	// 400
	InfantPrice *float64 `json:"infant_price,omitempty" xml:"infant_price,omitempty"`
	// infant tax
	//
	// example:
	//
	// 9
	InfantTax *float64 `json:"infant_tax,omitempty" xml:"infant_tax,omitempty"`
	// segment list
	JourneyList []*SearchResponseBodyDataSolutionListJourneyList `json:"journey_list,omitempty" xml:"journey_list,omitempty" type:"Repeated"`
	// through check-in baggage policy
	SegmentBaggageCheckInInfoList []*SearchResponseBodyDataSolutionListSegmentBaggageCheckInInfoList `json:"segment_baggage_check_in_info_list,omitempty" xml:"segment_baggage_check_in_info_list,omitempty" type:"Repeated"`
	// baggage rule
	SegmentBaggageMappingList []*SearchResponseBodyDataSolutionListSegmentBaggageMappingList `json:"segment_baggage_mapping_list,omitempty" xml:"segment_baggage_mapping_list,omitempty" type:"Repeated"`
	// change and refund policy
	SegmentRefundChangeRuleMappingList []*SearchResponseBodyDataSolutionListSegmentRefundChangeRuleMappingList `json:"segment_refund_change_rule_mapping_list,omitempty" xml:"segment_refund_change_rule_mapping_list,omitempty" type:"Repeated"`
	// Quotation attributes
	SolutionAttribute *SearchResponseBodyDataSolutionListSolutionAttribute `json:"solution_attribute,omitempty" xml:"solution_attribute,omitempty" type:"Struct"`
	// solution ID
	//
	// example:
	//
	// eJwz8DeySEo0NjQ01TU3TU7TNTFINNO1SE5O0jVKM0hKNjEwTElLNYwz0A32cNT1dfPVNTIwMjYwNjRQ8/A3NLI01Q0Ic0cRBwBVFxJJ
	SolutionId *string `json:"solution_id,omitempty" xml:"solution_id,omitempty"`
}

func (s SearchResponseBodyDataSolutionList) String() string {
	return dara.Prettify(s)
}

func (s SearchResponseBodyDataSolutionList) GoString() string {
	return s.String()
}

func (s *SearchResponseBodyDataSolutionList) GetAdultPrice() *float64 {
	return s.AdultPrice
}

func (s *SearchResponseBodyDataSolutionList) GetAdultTax() *float64 {
	return s.AdultTax
}

func (s *SearchResponseBodyDataSolutionList) GetChildPrice() *float64 {
	return s.ChildPrice
}

func (s *SearchResponseBodyDataSolutionList) GetChildTax() *float64 {
	return s.ChildTax
}

func (s *SearchResponseBodyDataSolutionList) GetInfantPrice() *float64 {
	return s.InfantPrice
}

func (s *SearchResponseBodyDataSolutionList) GetInfantTax() *float64 {
	return s.InfantTax
}

func (s *SearchResponseBodyDataSolutionList) GetJourneyList() []*SearchResponseBodyDataSolutionListJourneyList {
	return s.JourneyList
}

func (s *SearchResponseBodyDataSolutionList) GetSegmentBaggageCheckInInfoList() []*SearchResponseBodyDataSolutionListSegmentBaggageCheckInInfoList {
	return s.SegmentBaggageCheckInInfoList
}

func (s *SearchResponseBodyDataSolutionList) GetSegmentBaggageMappingList() []*SearchResponseBodyDataSolutionListSegmentBaggageMappingList {
	return s.SegmentBaggageMappingList
}

func (s *SearchResponseBodyDataSolutionList) GetSegmentRefundChangeRuleMappingList() []*SearchResponseBodyDataSolutionListSegmentRefundChangeRuleMappingList {
	return s.SegmentRefundChangeRuleMappingList
}

func (s *SearchResponseBodyDataSolutionList) GetSolutionAttribute() *SearchResponseBodyDataSolutionListSolutionAttribute {
	return s.SolutionAttribute
}

func (s *SearchResponseBodyDataSolutionList) GetSolutionId() *string {
	return s.SolutionId
}

func (s *SearchResponseBodyDataSolutionList) SetAdultPrice(v float64) *SearchResponseBodyDataSolutionList {
	s.AdultPrice = &v
	return s
}

func (s *SearchResponseBodyDataSolutionList) SetAdultTax(v float64) *SearchResponseBodyDataSolutionList {
	s.AdultTax = &v
	return s
}

func (s *SearchResponseBodyDataSolutionList) SetChildPrice(v float64) *SearchResponseBodyDataSolutionList {
	s.ChildPrice = &v
	return s
}

func (s *SearchResponseBodyDataSolutionList) SetChildTax(v float64) *SearchResponseBodyDataSolutionList {
	s.ChildTax = &v
	return s
}

func (s *SearchResponseBodyDataSolutionList) SetInfantPrice(v float64) *SearchResponseBodyDataSolutionList {
	s.InfantPrice = &v
	return s
}

func (s *SearchResponseBodyDataSolutionList) SetInfantTax(v float64) *SearchResponseBodyDataSolutionList {
	s.InfantTax = &v
	return s
}

func (s *SearchResponseBodyDataSolutionList) SetJourneyList(v []*SearchResponseBodyDataSolutionListJourneyList) *SearchResponseBodyDataSolutionList {
	s.JourneyList = v
	return s
}

func (s *SearchResponseBodyDataSolutionList) SetSegmentBaggageCheckInInfoList(v []*SearchResponseBodyDataSolutionListSegmentBaggageCheckInInfoList) *SearchResponseBodyDataSolutionList {
	s.SegmentBaggageCheckInInfoList = v
	return s
}

func (s *SearchResponseBodyDataSolutionList) SetSegmentBaggageMappingList(v []*SearchResponseBodyDataSolutionListSegmentBaggageMappingList) *SearchResponseBodyDataSolutionList {
	s.SegmentBaggageMappingList = v
	return s
}

func (s *SearchResponseBodyDataSolutionList) SetSegmentRefundChangeRuleMappingList(v []*SearchResponseBodyDataSolutionListSegmentRefundChangeRuleMappingList) *SearchResponseBodyDataSolutionList {
	s.SegmentRefundChangeRuleMappingList = v
	return s
}

func (s *SearchResponseBodyDataSolutionList) SetSolutionAttribute(v *SearchResponseBodyDataSolutionListSolutionAttribute) *SearchResponseBodyDataSolutionList {
	s.SolutionAttribute = v
	return s
}

func (s *SearchResponseBodyDataSolutionList) SetSolutionId(v string) *SearchResponseBodyDataSolutionList {
	s.SolutionId = &v
	return s
}

func (s *SearchResponseBodyDataSolutionList) Validate() error {
	if s.JourneyList != nil {
		for _, item := range s.JourneyList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SegmentBaggageCheckInInfoList != nil {
		for _, item := range s.SegmentBaggageCheckInInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SegmentBaggageMappingList != nil {
		for _, item := range s.SegmentBaggageMappingList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SegmentRefundChangeRuleMappingList != nil {
		for _, item := range s.SegmentRefundChangeRuleMappingList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SolutionAttribute != nil {
		if err := s.SolutionAttribute.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchResponseBodyDataSolutionListJourneyList struct {
	// segment Info
	SegmentList []*SearchResponseBodyDataSolutionListJourneyListSegmentList `json:"segment_list,omitempty" xml:"segment_list,omitempty" type:"Repeated"`
	// number of transfers
	//
	// example:
	//
	// 0
	TransferCount *int32 `json:"transfer_count,omitempty" xml:"transfer_count,omitempty"`
}

func (s SearchResponseBodyDataSolutionListJourneyList) String() string {
	return dara.Prettify(s)
}

func (s SearchResponseBodyDataSolutionListJourneyList) GoString() string {
	return s.String()
}

func (s *SearchResponseBodyDataSolutionListJourneyList) GetSegmentList() []*SearchResponseBodyDataSolutionListJourneyListSegmentList {
	return s.SegmentList
}

func (s *SearchResponseBodyDataSolutionListJourneyList) GetTransferCount() *int32 {
	return s.TransferCount
}

func (s *SearchResponseBodyDataSolutionListJourneyList) SetSegmentList(v []*SearchResponseBodyDataSolutionListJourneyListSegmentList) *SearchResponseBodyDataSolutionListJourneyList {
	s.SegmentList = v
	return s
}

func (s *SearchResponseBodyDataSolutionListJourneyList) SetTransferCount(v int32) *SearchResponseBodyDataSolutionListJourneyList {
	s.TransferCount = &v
	return s
}

func (s *SearchResponseBodyDataSolutionListJourneyList) Validate() error {
	if s.SegmentList != nil {
		for _, item := range s.SegmentList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SearchResponseBodyDataSolutionListJourneyListSegmentList struct {
	// arrival airport code (capitalized)
	//
	// example:
	//
	// MFM
	ArrivalAirport *string `json:"arrival_airport,omitempty" xml:"arrival_airport,omitempty"`
	// arrival city code (capitalized)
	//
	// example:
	//
	// MFM
	ArrivalCity *string `json:"arrival_city,omitempty" xml:"arrival_city,omitempty"`
	// arrival terminal
	//
	// example:
	//
	// T1
	ArrivalTerminal *string `json:"arrival_terminal,omitempty" xml:"arrival_terminal,omitempty"`
	// arrival time in string format (yyyy-MM-dd HH:mm:ss)
	//
	// example:
	//
	// 2023-03-10 10:40:00
	ArrivalTime *string `json:"arrival_time,omitempty" xml:"arrival_time,omitempty"`
	// available seats (for reference only)
	//
	// example:
	//
	// 7
	Availability *string `json:"availability,omitempty" xml:"availability,omitempty"`
	// RBD
	//
	// example:
	//
	// V
	Cabin *string `json:"cabin,omitempty" xml:"cabin,omitempty"`
	// cabin class
	//
	// example:
	//
	// Y
	CabinClass *string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	// code share or not
	//
	// example:
	//
	// false
	CodeShare *bool `json:"code_share,omitempty" xml:"code_share,omitempty"`
	// departure airport code (capitalized)
	//
	// example:
	//
	// PVG
	DepartureAirport *string `json:"departure_airport,omitempty" xml:"departure_airport,omitempty"`
	// departure city code (capitalized)
	//
	// example:
	//
	// SHA
	DepartureCity *string `json:"departure_city,omitempty" xml:"departure_city,omitempty"`
	// departure terminal
	//
	// example:
	//
	// T2
	DepartureTerminal *string `json:"departure_terminal,omitempty" xml:"departure_terminal,omitempty"`
	// departure time in string format (yyyy-MM-dd HH:mm:ss)
	//
	// example:
	//
	// 2023-03-10 07:55:00
	DepartureTime *string `json:"departure_time,omitempty" xml:"departure_time,omitempty"`
	// equipment type
	//
	// example:
	//
	// 32Q
	EquipType *string `json:"equip_type,omitempty" xml:"equip_type,omitempty"`
	// flight time, unit: minute
	//
	// example:
	//
	// 165
	FlightDuration *int32 `json:"flight_duration,omitempty" xml:"flight_duration,omitempty"`
	// marketing airline code (ex.: KA)
	//
	// example:
	//
	// HO
	MarketingAirline *string `json:"marketing_airline,omitempty" xml:"marketing_airline,omitempty"`
	// marketing airline flight no. (ex.: KA5809)
	//
	// example:
	//
	// HO1295
	MarketingFlightNo *string `json:"marketing_flight_no,omitempty" xml:"marketing_flight_no,omitempty"`
	// marketing airline integer flight no. (ex.: 5809)
	//
	// example:
	//
	// 1259
	MarketingFlightNoInt *int32 `json:"marketing_flight_no_int,omitempty" xml:"marketing_flight_no_int,omitempty"`
	// operating airline code (ex.: CX)
	//
	// example:
	//
	// HO
	OperatingAirline *string `json:"operating_airline,omitempty" xml:"operating_airline,omitempty"`
	// operating airline flight no. (ex.: CX601)
	//
	// example:
	//
	// HO1295
	OperatingFlightNo *string `json:"operating_flight_no,omitempty" xml:"operating_flight_no,omitempty"`
	// segment ID format: flight no.+departure airport[IATA airport code]+arrival airport[IATA airport code]+departure time(MMdd)
	//
	// example:
	//
	// HO1295-PVG-MFM-20230310
	SegmentId *string `json:"segment_id,omitempty" xml:"segment_id,omitempty"`
	// stop city list
	//
	// when stop_quantity > 1, use “,” for seperation
	//
	// example:
	//
	// MFM,PVG
	StopCityList *string `json:"stop_city_list,omitempty" xml:"stop_city_list,omitempty"`
	// number of stops
	//
	// example:
	//
	// 0
	StopQuantity *int32 `json:"stop_quantity,omitempty" xml:"stop_quantity,omitempty"`
}

func (s SearchResponseBodyDataSolutionListJourneyListSegmentList) String() string {
	return dara.Prettify(s)
}

func (s SearchResponseBodyDataSolutionListJourneyListSegmentList) GoString() string {
	return s.String()
}

func (s *SearchResponseBodyDataSolutionListJourneyListSegmentList) GetArrivalAirport() *string {
	return s.ArrivalAirport
}

func (s *SearchResponseBodyDataSolutionListJourneyListSegmentList) GetArrivalCity() *string {
	return s.ArrivalCity
}

func (s *SearchResponseBodyDataSolutionListJourneyListSegmentList) GetArrivalTerminal() *string {
	return s.ArrivalTerminal
}

func (s *SearchResponseBodyDataSolutionListJourneyListSegmentList) GetArrivalTime() *string {
	return s.ArrivalTime
}

func (s *SearchResponseBodyDataSolutionListJourneyListSegmentList) GetAvailability() *string {
	return s.Availability
}

func (s *SearchResponseBodyDataSolutionListJourneyListSegmentList) GetCabin() *string {
	return s.Cabin
}

func (s *SearchResponseBodyDataSolutionListJourneyListSegmentList) GetCabinClass() *string {
	return s.CabinClass
}

func (s *SearchResponseBodyDataSolutionListJourneyListSegmentList) GetCodeShare() *bool {
	return s.CodeShare
}

func (s *SearchResponseBodyDataSolutionListJourneyListSegmentList) GetDepartureAirport() *string {
	return s.DepartureAirport
}

func (s *SearchResponseBodyDataSolutionListJourneyListSegmentList) GetDepartureCity() *string {
	return s.DepartureCity
}

func (s *SearchResponseBodyDataSolutionListJourneyListSegmentList) GetDepartureTerminal() *string {
	return s.DepartureTerminal
}

func (s *SearchResponseBodyDataSolutionListJourneyListSegmentList) GetDepartureTime() *string {
	return s.DepartureTime
}

func (s *SearchResponseBodyDataSolutionListJourneyListSegmentList) GetEquipType() *string {
	return s.EquipType
}

func (s *SearchResponseBodyDataSolutionListJourneyListSegmentList) GetFlightDuration() *int32 {
	return s.FlightDuration
}

func (s *SearchResponseBodyDataSolutionListJourneyListSegmentList) GetMarketingAirline() *string {
	return s.MarketingAirline
}

func (s *SearchResponseBodyDataSolutionListJourneyListSegmentList) GetMarketingFlightNo() *string {
	return s.MarketingFlightNo
}

func (s *SearchResponseBodyDataSolutionListJourneyListSegmentList) GetMarketingFlightNoInt() *int32 {
	return s.MarketingFlightNoInt
}

func (s *SearchResponseBodyDataSolutionListJourneyListSegmentList) GetOperatingAirline() *string {
	return s.OperatingAirline
}

func (s *SearchResponseBodyDataSolutionListJourneyListSegmentList) GetOperatingFlightNo() *string {
	return s.OperatingFlightNo
}

func (s *SearchResponseBodyDataSolutionListJourneyListSegmentList) GetSegmentId() *string {
	return s.SegmentId
}

func (s *SearchResponseBodyDataSolutionListJourneyListSegmentList) GetStopCityList() *string {
	return s.StopCityList
}

func (s *SearchResponseBodyDataSolutionListJourneyListSegmentList) GetStopQuantity() *int32 {
	return s.StopQuantity
}

func (s *SearchResponseBodyDataSolutionListJourneyListSegmentList) SetArrivalAirport(v string) *SearchResponseBodyDataSolutionListJourneyListSegmentList {
	s.ArrivalAirport = &v
	return s
}

func (s *SearchResponseBodyDataSolutionListJourneyListSegmentList) SetArrivalCity(v string) *SearchResponseBodyDataSolutionListJourneyListSegmentList {
	s.ArrivalCity = &v
	return s
}

func (s *SearchResponseBodyDataSolutionListJourneyListSegmentList) SetArrivalTerminal(v string) *SearchResponseBodyDataSolutionListJourneyListSegmentList {
	s.ArrivalTerminal = &v
	return s
}

func (s *SearchResponseBodyDataSolutionListJourneyListSegmentList) SetArrivalTime(v string) *SearchResponseBodyDataSolutionListJourneyListSegmentList {
	s.ArrivalTime = &v
	return s
}

func (s *SearchResponseBodyDataSolutionListJourneyListSegmentList) SetAvailability(v string) *SearchResponseBodyDataSolutionListJourneyListSegmentList {
	s.Availability = &v
	return s
}

func (s *SearchResponseBodyDataSolutionListJourneyListSegmentList) SetCabin(v string) *SearchResponseBodyDataSolutionListJourneyListSegmentList {
	s.Cabin = &v
	return s
}

func (s *SearchResponseBodyDataSolutionListJourneyListSegmentList) SetCabinClass(v string) *SearchResponseBodyDataSolutionListJourneyListSegmentList {
	s.CabinClass = &v
	return s
}

func (s *SearchResponseBodyDataSolutionListJourneyListSegmentList) SetCodeShare(v bool) *SearchResponseBodyDataSolutionListJourneyListSegmentList {
	s.CodeShare = &v
	return s
}

func (s *SearchResponseBodyDataSolutionListJourneyListSegmentList) SetDepartureAirport(v string) *SearchResponseBodyDataSolutionListJourneyListSegmentList {
	s.DepartureAirport = &v
	return s
}

func (s *SearchResponseBodyDataSolutionListJourneyListSegmentList) SetDepartureCity(v string) *SearchResponseBodyDataSolutionListJourneyListSegmentList {
	s.DepartureCity = &v
	return s
}

func (s *SearchResponseBodyDataSolutionListJourneyListSegmentList) SetDepartureTerminal(v string) *SearchResponseBodyDataSolutionListJourneyListSegmentList {
	s.DepartureTerminal = &v
	return s
}

func (s *SearchResponseBodyDataSolutionListJourneyListSegmentList) SetDepartureTime(v string) *SearchResponseBodyDataSolutionListJourneyListSegmentList {
	s.DepartureTime = &v
	return s
}

func (s *SearchResponseBodyDataSolutionListJourneyListSegmentList) SetEquipType(v string) *SearchResponseBodyDataSolutionListJourneyListSegmentList {
	s.EquipType = &v
	return s
}

func (s *SearchResponseBodyDataSolutionListJourneyListSegmentList) SetFlightDuration(v int32) *SearchResponseBodyDataSolutionListJourneyListSegmentList {
	s.FlightDuration = &v
	return s
}

func (s *SearchResponseBodyDataSolutionListJourneyListSegmentList) SetMarketingAirline(v string) *SearchResponseBodyDataSolutionListJourneyListSegmentList {
	s.MarketingAirline = &v
	return s
}

func (s *SearchResponseBodyDataSolutionListJourneyListSegmentList) SetMarketingFlightNo(v string) *SearchResponseBodyDataSolutionListJourneyListSegmentList {
	s.MarketingFlightNo = &v
	return s
}

func (s *SearchResponseBodyDataSolutionListJourneyListSegmentList) SetMarketingFlightNoInt(v int32) *SearchResponseBodyDataSolutionListJourneyListSegmentList {
	s.MarketingFlightNoInt = &v
	return s
}

func (s *SearchResponseBodyDataSolutionListJourneyListSegmentList) SetOperatingAirline(v string) *SearchResponseBodyDataSolutionListJourneyListSegmentList {
	s.OperatingAirline = &v
	return s
}

func (s *SearchResponseBodyDataSolutionListJourneyListSegmentList) SetOperatingFlightNo(v string) *SearchResponseBodyDataSolutionListJourneyListSegmentList {
	s.OperatingFlightNo = &v
	return s
}

func (s *SearchResponseBodyDataSolutionListJourneyListSegmentList) SetSegmentId(v string) *SearchResponseBodyDataSolutionListJourneyListSegmentList {
	s.SegmentId = &v
	return s
}

func (s *SearchResponseBodyDataSolutionListJourneyListSegmentList) SetStopCityList(v string) *SearchResponseBodyDataSolutionListJourneyListSegmentList {
	s.StopCityList = &v
	return s
}

func (s *SearchResponseBodyDataSolutionListJourneyListSegmentList) SetStopQuantity(v int32) *SearchResponseBodyDataSolutionListJourneyListSegmentList {
	s.StopQuantity = &v
	return s
}

func (s *SearchResponseBodyDataSolutionListJourneyListSegmentList) Validate() error {
	return dara.Validate(s)
}

type SearchResponseBodyDataSolutionListSegmentBaggageCheckInInfoList struct {
	// through check-in baggage policy type
	//
	// 1. baggage through check-in between segments
	//
	// 2. baggage re-check-in needed between segments
	//
	// 3. baggage through check-in at stop city ( applies for stop flight )
	//
	// 4. baggage re-checkin needed at stop city ( applies for stop flight )
	//
	// example:
	//
	// 1
	LuggageDirectInfoType *int32 `json:"luggage_direct_info_type,omitempty" xml:"luggage_direct_info_type,omitempty"`
	// segment id list.
	//
	// all the listed segment ids share the same baggage through check-in  policy
	SegmentIdList []*string `json:"segment_id_list,omitempty" xml:"segment_id_list,omitempty" type:"Repeated"`
}

func (s SearchResponseBodyDataSolutionListSegmentBaggageCheckInInfoList) String() string {
	return dara.Prettify(s)
}

func (s SearchResponseBodyDataSolutionListSegmentBaggageCheckInInfoList) GoString() string {
	return s.String()
}

func (s *SearchResponseBodyDataSolutionListSegmentBaggageCheckInInfoList) GetLuggageDirectInfoType() *int32 {
	return s.LuggageDirectInfoType
}

func (s *SearchResponseBodyDataSolutionListSegmentBaggageCheckInInfoList) GetSegmentIdList() []*string {
	return s.SegmentIdList
}

func (s *SearchResponseBodyDataSolutionListSegmentBaggageCheckInInfoList) SetLuggageDirectInfoType(v int32) *SearchResponseBodyDataSolutionListSegmentBaggageCheckInInfoList {
	s.LuggageDirectInfoType = &v
	return s
}

func (s *SearchResponseBodyDataSolutionListSegmentBaggageCheckInInfoList) SetSegmentIdList(v []*string) *SearchResponseBodyDataSolutionListSegmentBaggageCheckInInfoList {
	s.SegmentIdList = v
	return s
}

func (s *SearchResponseBodyDataSolutionListSegmentBaggageCheckInInfoList) Validate() error {
	return dara.Validate(s)
}

type SearchResponseBodyDataSolutionListSegmentBaggageMappingList struct {
	// baggage rule mapping, key is passenger type, value is baggage allowance details
	PassengerBaggageAllowanceMapping map[string]*DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue `json:"passenger_baggage_allowance_mapping,omitempty" xml:"passenger_baggage_allowance_mapping,omitempty"`
	// segment id list.
	//
	// all the listed segment id share the same baggage rule
	SegmentIdList []*string `json:"segment_id_list,omitempty" xml:"segment_id_list,omitempty" type:"Repeated"`
}

func (s SearchResponseBodyDataSolutionListSegmentBaggageMappingList) String() string {
	return dara.Prettify(s)
}

func (s SearchResponseBodyDataSolutionListSegmentBaggageMappingList) GoString() string {
	return s.String()
}

func (s *SearchResponseBodyDataSolutionListSegmentBaggageMappingList) GetPassengerBaggageAllowanceMapping() map[string]*DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue {
	return s.PassengerBaggageAllowanceMapping
}

func (s *SearchResponseBodyDataSolutionListSegmentBaggageMappingList) GetSegmentIdList() []*string {
	return s.SegmentIdList
}

func (s *SearchResponseBodyDataSolutionListSegmentBaggageMappingList) SetPassengerBaggageAllowanceMapping(v map[string]*DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) *SearchResponseBodyDataSolutionListSegmentBaggageMappingList {
	s.PassengerBaggageAllowanceMapping = v
	return s
}

func (s *SearchResponseBodyDataSolutionListSegmentBaggageMappingList) SetSegmentIdList(v []*string) *SearchResponseBodyDataSolutionListSegmentBaggageMappingList {
	s.SegmentIdList = v
	return s
}

func (s *SearchResponseBodyDataSolutionListSegmentBaggageMappingList) Validate() error {
	return dara.Validate(s)
}

type SearchResponseBodyDataSolutionListSegmentRefundChangeRuleMappingList struct {
	// change and refund policy mapping, key is passenger type, value is change and refund policy details
	RefundChangeRuleMap map[string]*DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValue `json:"refund_change_rule_map,omitempty" xml:"refund_change_rule_map,omitempty"`
	// segment id list.
	//
	// all the listed segment ids share the same change and refund policy
	SegmentIdList []*string `json:"segment_id_list,omitempty" xml:"segment_id_list,omitempty" type:"Repeated"`
}

func (s SearchResponseBodyDataSolutionListSegmentRefundChangeRuleMappingList) String() string {
	return dara.Prettify(s)
}

func (s SearchResponseBodyDataSolutionListSegmentRefundChangeRuleMappingList) GoString() string {
	return s.String()
}

func (s *SearchResponseBodyDataSolutionListSegmentRefundChangeRuleMappingList) GetRefundChangeRuleMap() map[string]*DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValue {
	return s.RefundChangeRuleMap
}

func (s *SearchResponseBodyDataSolutionListSegmentRefundChangeRuleMappingList) GetSegmentIdList() []*string {
	return s.SegmentIdList
}

func (s *SearchResponseBodyDataSolutionListSegmentRefundChangeRuleMappingList) SetRefundChangeRuleMap(v map[string]*DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValue) *SearchResponseBodyDataSolutionListSegmentRefundChangeRuleMappingList {
	s.RefundChangeRuleMap = v
	return s
}

func (s *SearchResponseBodyDataSolutionListSegmentRefundChangeRuleMappingList) SetSegmentIdList(v []*string) *SearchResponseBodyDataSolutionListSegmentRefundChangeRuleMappingList {
	s.SegmentIdList = v
	return s
}

func (s *SearchResponseBodyDataSolutionListSegmentRefundChangeRuleMappingList) Validate() error {
	return dara.Validate(s)
}

type SearchResponseBodyDataSolutionListSolutionAttribute struct {
	// Issue ticket time related
	IssueTimeInfo *SearchResponseBodyDataSolutionListSolutionAttributeIssueTimeInfo `json:"issue_time_info,omitempty" xml:"issue_time_info,omitempty" type:"Struct"`
	// Supply source type 1: Self-operated; 2: Agent; 3: Flagship store
	//
	// example:
	//
	// 1
	SupplySourceType *string `json:"supply_source_type,omitempty" xml:"supply_source_type,omitempty"`
}

func (s SearchResponseBodyDataSolutionListSolutionAttribute) String() string {
	return dara.Prettify(s)
}

func (s SearchResponseBodyDataSolutionListSolutionAttribute) GoString() string {
	return s.String()
}

func (s *SearchResponseBodyDataSolutionListSolutionAttribute) GetIssueTimeInfo() *SearchResponseBodyDataSolutionListSolutionAttributeIssueTimeInfo {
	return s.IssueTimeInfo
}

func (s *SearchResponseBodyDataSolutionListSolutionAttribute) GetSupplySourceType() *string {
	return s.SupplySourceType
}

func (s *SearchResponseBodyDataSolutionListSolutionAttribute) SetIssueTimeInfo(v *SearchResponseBodyDataSolutionListSolutionAttributeIssueTimeInfo) *SearchResponseBodyDataSolutionListSolutionAttribute {
	s.IssueTimeInfo = v
	return s
}

func (s *SearchResponseBodyDataSolutionListSolutionAttribute) SetSupplySourceType(v string) *SearchResponseBodyDataSolutionListSolutionAttribute {
	s.SupplySourceType = &v
	return s
}

func (s *SearchResponseBodyDataSolutionListSolutionAttribute) Validate() error {
	if s.IssueTimeInfo != nil {
		if err := s.IssueTimeInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchResponseBodyDataSolutionListSolutionAttributeIssueTimeInfo struct {
	// Issue ticket type: 1: after payment; 2: before departure; -1: unknown
	IssueTicketType *int32 `json:"issue_ticket_type,omitempty" xml:"issue_ticket_type,omitempty"`
	// Estimated issue ticket time, unit: minutes
	IssueTimeLimit *int32 `json:"issue_time_limit,omitempty" xml:"issue_time_limit,omitempty"`
}

func (s SearchResponseBodyDataSolutionListSolutionAttributeIssueTimeInfo) String() string {
	return dara.Prettify(s)
}

func (s SearchResponseBodyDataSolutionListSolutionAttributeIssueTimeInfo) GoString() string {
	return s.String()
}

func (s *SearchResponseBodyDataSolutionListSolutionAttributeIssueTimeInfo) GetIssueTicketType() *int32 {
	return s.IssueTicketType
}

func (s *SearchResponseBodyDataSolutionListSolutionAttributeIssueTimeInfo) GetIssueTimeLimit() *int32 {
	return s.IssueTimeLimit
}

func (s *SearchResponseBodyDataSolutionListSolutionAttributeIssueTimeInfo) SetIssueTicketType(v int32) *SearchResponseBodyDataSolutionListSolutionAttributeIssueTimeInfo {
	s.IssueTicketType = &v
	return s
}

func (s *SearchResponseBodyDataSolutionListSolutionAttributeIssueTimeInfo) SetIssueTimeLimit(v int32) *SearchResponseBodyDataSolutionListSolutionAttributeIssueTimeInfo {
	s.IssueTimeLimit = &v
	return s
}

func (s *SearchResponseBodyDataSolutionListSolutionAttributeIssueTimeInfo) Validate() error {
	return dara.Validate(s)
}
