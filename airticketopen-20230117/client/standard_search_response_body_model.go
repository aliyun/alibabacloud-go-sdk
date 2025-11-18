// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStandardSearchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StandardSearchResponseBody
	GetRequestId() *string
	SetData(v *StandardSearchResponseBodyData) *StandardSearchResponseBody
	GetData() *StandardSearchResponseBodyData
	SetErrorCode(v string) *StandardSearchResponseBody
	GetErrorCode() *string
	SetErrorData(v interface{}) *StandardSearchResponseBody
	GetErrorData() interface{}
	SetErrorMsg(v string) *StandardSearchResponseBody
	GetErrorMsg() *string
	SetStatus(v int32) *StandardSearchResponseBody
	GetStatus() *int32
	SetSuccess(v bool) *StandardSearchResponseBody
	GetSuccess() *bool
}

type StandardSearchResponseBody struct {
	// example:
	//
	// 51593418-8C73-5E47-8BA8-3F1D4A00CC0B
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *StandardSearchResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// null
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// example:
	//
	// null
	ErrorData interface{} `json:"error_data,omitempty" xml:"error_data,omitempty"`
	// example:
	//
	// null
	ErrorMsg *string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// example:
	//
	// 200
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s StandardSearchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StandardSearchResponseBody) GoString() string {
	return s.String()
}

func (s *StandardSearchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StandardSearchResponseBody) GetData() *StandardSearchResponseBodyData {
	return s.Data
}

func (s *StandardSearchResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *StandardSearchResponseBody) GetErrorData() interface{} {
	return s.ErrorData
}

func (s *StandardSearchResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *StandardSearchResponseBody) GetStatus() *int32 {
	return s.Status
}

func (s *StandardSearchResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StandardSearchResponseBody) SetRequestId(v string) *StandardSearchResponseBody {
	s.RequestId = &v
	return s
}

func (s *StandardSearchResponseBody) SetData(v *StandardSearchResponseBodyData) *StandardSearchResponseBody {
	s.Data = v
	return s
}

func (s *StandardSearchResponseBody) SetErrorCode(v string) *StandardSearchResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *StandardSearchResponseBody) SetErrorData(v interface{}) *StandardSearchResponseBody {
	s.ErrorData = v
	return s
}

func (s *StandardSearchResponseBody) SetErrorMsg(v string) *StandardSearchResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *StandardSearchResponseBody) SetStatus(v int32) *StandardSearchResponseBody {
	s.Status = &v
	return s
}

func (s *StandardSearchResponseBody) SetSuccess(v bool) *StandardSearchResponseBody {
	s.Success = &v
	return s
}

func (s *StandardSearchResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type StandardSearchResponseBodyData struct {
	SolutionList []*StandardSearchResponseBodyDataSolutionList `json:"solution_list,omitempty" xml:"solution_list,omitempty" type:"Repeated"`
}

func (s StandardSearchResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s StandardSearchResponseBodyData) GoString() string {
	return s.String()
}

func (s *StandardSearchResponseBodyData) GetSolutionList() []*StandardSearchResponseBodyDataSolutionList {
	return s.SolutionList
}

func (s *StandardSearchResponseBodyData) SetSolutionList(v []*StandardSearchResponseBodyDataSolutionList) *StandardSearchResponseBodyData {
	s.SolutionList = v
	return s
}

func (s *StandardSearchResponseBodyData) Validate() error {
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

type StandardSearchResponseBodyDataSolutionList struct {
	// example:
	//
	// 300
	AdultPrice *float64 `json:"adult_price,omitempty" xml:"adult_price,omitempty"`
	// example:
	//
	// 30
	AdultTax *float64 `json:"adult_tax,omitempty" xml:"adult_tax,omitempty"`
	// example:
	//
	// 200
	ChildPrice *float64 `json:"child_price,omitempty" xml:"child_price,omitempty"`
	// example:
	//
	// 20
	ChildTax *float64 `json:"child_tax,omitempty" xml:"child_tax,omitempty"`
	// example:
	//
	// 100
	InfantPrice *float64 `json:"infant_price,omitempty" xml:"infant_price,omitempty"`
	// example:
	//
	// 10
	InfantTax                          *float64                                                                        `json:"infant_tax,omitempty" xml:"infant_tax,omitempty"`
	JourneyList                        []*StandardSearchResponseBodyDataSolutionListJourneyList                        `json:"journey_list,omitempty" xml:"journey_list,omitempty" type:"Repeated"`
	SegmentBaggageCheckInInfoList      []*StandardSearchResponseBodyDataSolutionListSegmentBaggageCheckInInfoList      `json:"segment_baggage_check_in_info_list,omitempty" xml:"segment_baggage_check_in_info_list,omitempty" type:"Repeated"`
	SegmentBaggageMappingList          []*StandardSearchResponseBodyDataSolutionListSegmentBaggageMappingList          `json:"segment_baggage_mapping_list,omitempty" xml:"segment_baggage_mapping_list,omitempty" type:"Repeated"`
	SegmentRefundChangeRuleMappingList []*StandardSearchResponseBodyDataSolutionListSegmentRefundChangeRuleMappingList `json:"segment_refund_change_rule_mapping_list,omitempty" xml:"segment_refund_change_rule_mapping_list,omitempty" type:"Repeated"`
	SolutionAttribute                  *StandardSearchResponseBodyDataSolutionListSolutionAttribute                    `json:"solution_attribute,omitempty" xml:"solution_attribute,omitempty" type:"Struct"`
	// solution_id
	//
	// example:
	//
	// eJwz8DeySEo0NjQ01TU3TUxxx
	SolutionId *string `json:"solution_id,omitempty" xml:"solution_id,omitempty"`
}

func (s StandardSearchResponseBodyDataSolutionList) String() string {
	return dara.Prettify(s)
}

func (s StandardSearchResponseBodyDataSolutionList) GoString() string {
	return s.String()
}

func (s *StandardSearchResponseBodyDataSolutionList) GetAdultPrice() *float64 {
	return s.AdultPrice
}

func (s *StandardSearchResponseBodyDataSolutionList) GetAdultTax() *float64 {
	return s.AdultTax
}

func (s *StandardSearchResponseBodyDataSolutionList) GetChildPrice() *float64 {
	return s.ChildPrice
}

func (s *StandardSearchResponseBodyDataSolutionList) GetChildTax() *float64 {
	return s.ChildTax
}

func (s *StandardSearchResponseBodyDataSolutionList) GetInfantPrice() *float64 {
	return s.InfantPrice
}

func (s *StandardSearchResponseBodyDataSolutionList) GetInfantTax() *float64 {
	return s.InfantTax
}

func (s *StandardSearchResponseBodyDataSolutionList) GetJourneyList() []*StandardSearchResponseBodyDataSolutionListJourneyList {
	return s.JourneyList
}

func (s *StandardSearchResponseBodyDataSolutionList) GetSegmentBaggageCheckInInfoList() []*StandardSearchResponseBodyDataSolutionListSegmentBaggageCheckInInfoList {
	return s.SegmentBaggageCheckInInfoList
}

func (s *StandardSearchResponseBodyDataSolutionList) GetSegmentBaggageMappingList() []*StandardSearchResponseBodyDataSolutionListSegmentBaggageMappingList {
	return s.SegmentBaggageMappingList
}

func (s *StandardSearchResponseBodyDataSolutionList) GetSegmentRefundChangeRuleMappingList() []*StandardSearchResponseBodyDataSolutionListSegmentRefundChangeRuleMappingList {
	return s.SegmentRefundChangeRuleMappingList
}

func (s *StandardSearchResponseBodyDataSolutionList) GetSolutionAttribute() *StandardSearchResponseBodyDataSolutionListSolutionAttribute {
	return s.SolutionAttribute
}

func (s *StandardSearchResponseBodyDataSolutionList) GetSolutionId() *string {
	return s.SolutionId
}

func (s *StandardSearchResponseBodyDataSolutionList) SetAdultPrice(v float64) *StandardSearchResponseBodyDataSolutionList {
	s.AdultPrice = &v
	return s
}

func (s *StandardSearchResponseBodyDataSolutionList) SetAdultTax(v float64) *StandardSearchResponseBodyDataSolutionList {
	s.AdultTax = &v
	return s
}

func (s *StandardSearchResponseBodyDataSolutionList) SetChildPrice(v float64) *StandardSearchResponseBodyDataSolutionList {
	s.ChildPrice = &v
	return s
}

func (s *StandardSearchResponseBodyDataSolutionList) SetChildTax(v float64) *StandardSearchResponseBodyDataSolutionList {
	s.ChildTax = &v
	return s
}

func (s *StandardSearchResponseBodyDataSolutionList) SetInfantPrice(v float64) *StandardSearchResponseBodyDataSolutionList {
	s.InfantPrice = &v
	return s
}

func (s *StandardSearchResponseBodyDataSolutionList) SetInfantTax(v float64) *StandardSearchResponseBodyDataSolutionList {
	s.InfantTax = &v
	return s
}

func (s *StandardSearchResponseBodyDataSolutionList) SetJourneyList(v []*StandardSearchResponseBodyDataSolutionListJourneyList) *StandardSearchResponseBodyDataSolutionList {
	s.JourneyList = v
	return s
}

func (s *StandardSearchResponseBodyDataSolutionList) SetSegmentBaggageCheckInInfoList(v []*StandardSearchResponseBodyDataSolutionListSegmentBaggageCheckInInfoList) *StandardSearchResponseBodyDataSolutionList {
	s.SegmentBaggageCheckInInfoList = v
	return s
}

func (s *StandardSearchResponseBodyDataSolutionList) SetSegmentBaggageMappingList(v []*StandardSearchResponseBodyDataSolutionListSegmentBaggageMappingList) *StandardSearchResponseBodyDataSolutionList {
	s.SegmentBaggageMappingList = v
	return s
}

func (s *StandardSearchResponseBodyDataSolutionList) SetSegmentRefundChangeRuleMappingList(v []*StandardSearchResponseBodyDataSolutionListSegmentRefundChangeRuleMappingList) *StandardSearchResponseBodyDataSolutionList {
	s.SegmentRefundChangeRuleMappingList = v
	return s
}

func (s *StandardSearchResponseBodyDataSolutionList) SetSolutionAttribute(v *StandardSearchResponseBodyDataSolutionListSolutionAttribute) *StandardSearchResponseBodyDataSolutionList {
	s.SolutionAttribute = v
	return s
}

func (s *StandardSearchResponseBodyDataSolutionList) SetSolutionId(v string) *StandardSearchResponseBodyDataSolutionList {
	s.SolutionId = &v
	return s
}

func (s *StandardSearchResponseBodyDataSolutionList) Validate() error {
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

type StandardSearchResponseBodyDataSolutionListJourneyList struct {
	SegmentList []*StandardSearchResponseBodyDataSolutionListJourneyListSegmentList `json:"segment_list,omitempty" xml:"segment_list,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	TransferCount *int32 `json:"transfer_count,omitempty" xml:"transfer_count,omitempty"`
}

func (s StandardSearchResponseBodyDataSolutionListJourneyList) String() string {
	return dara.Prettify(s)
}

func (s StandardSearchResponseBodyDataSolutionListJourneyList) GoString() string {
	return s.String()
}

func (s *StandardSearchResponseBodyDataSolutionListJourneyList) GetSegmentList() []*StandardSearchResponseBodyDataSolutionListJourneyListSegmentList {
	return s.SegmentList
}

func (s *StandardSearchResponseBodyDataSolutionListJourneyList) GetTransferCount() *int32 {
	return s.TransferCount
}

func (s *StandardSearchResponseBodyDataSolutionListJourneyList) SetSegmentList(v []*StandardSearchResponseBodyDataSolutionListJourneyListSegmentList) *StandardSearchResponseBodyDataSolutionListJourneyList {
	s.SegmentList = v
	return s
}

func (s *StandardSearchResponseBodyDataSolutionListJourneyList) SetTransferCount(v int32) *StandardSearchResponseBodyDataSolutionListJourneyList {
	s.TransferCount = &v
	return s
}

func (s *StandardSearchResponseBodyDataSolutionListJourneyList) Validate() error {
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

type StandardSearchResponseBodyDataSolutionListJourneyListSegmentList struct {
	// example:
	//
	// MFM
	ArrivalAirport *string `json:"arrival_airport,omitempty" xml:"arrival_airport,omitempty"`
	// example:
	//
	// MFM
	ArrivalCity *string `json:"arrival_city,omitempty" xml:"arrival_city,omitempty"`
	// example:
	//
	// T1
	ArrivalTerminal *string `json:"arrival_terminal,omitempty" xml:"arrival_terminal,omitempty"`
	// example:
	//
	// 2023-03-10 10:40:00
	ArrivalTime *string `json:"arrival_time,omitempty" xml:"arrival_time,omitempty"`
	// example:
	//
	// 7
	Availability *string `json:"availability,omitempty" xml:"availability,omitempty"`
	// example:
	//
	// V
	Cabin *string `json:"cabin,omitempty" xml:"cabin,omitempty"`
	// example:
	//
	// Y
	CabinClass *string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	// example:
	//
	// false
	CodeShare *bool `json:"code_share,omitempty" xml:"code_share,omitempty"`
	// example:
	//
	// PVG
	DepartureAirport *string `json:"departure_airport,omitempty" xml:"departure_airport,omitempty"`
	// example:
	//
	// SHA
	DepartureCity *string `json:"departure_city,omitempty" xml:"departure_city,omitempty"`
	// example:
	//
	// T2
	DepartureTerminal *string `json:"departure_terminal,omitempty" xml:"departure_terminal,omitempty"`
	// example:
	//
	// 2023-03-10 07:55:00
	DepartureTime *string `json:"departure_time,omitempty" xml:"departure_time,omitempty"`
	// example:
	//
	// 32Q
	EquipType *string `json:"equip_type,omitempty" xml:"equip_type,omitempty"`
	// example:
	//
	// 165
	FlightDuration *int32 `json:"flight_duration,omitempty" xml:"flight_duration,omitempty"`
	// example:
	//
	// HO
	MarketingAirline *string `json:"marketing_airline,omitempty" xml:"marketing_airline,omitempty"`
	// example:
	//
	// HO1295
	MarketingFlightNo *string `json:"marketing_flight_no,omitempty" xml:"marketing_flight_no,omitempty"`
	// example:
	//
	// 1295
	MarketingFlightNoInt *int32 `json:"marketing_flight_no_int,omitempty" xml:"marketing_flight_no_int,omitempty"`
	// example:
	//
	// HO
	OperatingAirline *string `json:"operating_airline,omitempty" xml:"operating_airline,omitempty"`
	// example:
	//
	// HO1295
	OperatingFlightNo *string `json:"operating_flight_no,omitempty" xml:"operating_flight_no,omitempty"`
	// example:
	//
	// HO1295-PVG-MFM-20230310
	SegmentId *string `json:"segment_id,omitempty" xml:"segment_id,omitempty"`
	// example:
	//
	// SEL,HKG
	StopCityList *string `json:"stop_city_list,omitempty" xml:"stop_city_list,omitempty"`
	// example:
	//
	// 0
	StopQuantity *int32 `json:"stop_quantity,omitempty" xml:"stop_quantity,omitempty"`
}

func (s StandardSearchResponseBodyDataSolutionListJourneyListSegmentList) String() string {
	return dara.Prettify(s)
}

func (s StandardSearchResponseBodyDataSolutionListJourneyListSegmentList) GoString() string {
	return s.String()
}

func (s *StandardSearchResponseBodyDataSolutionListJourneyListSegmentList) GetArrivalAirport() *string {
	return s.ArrivalAirport
}

func (s *StandardSearchResponseBodyDataSolutionListJourneyListSegmentList) GetArrivalCity() *string {
	return s.ArrivalCity
}

func (s *StandardSearchResponseBodyDataSolutionListJourneyListSegmentList) GetArrivalTerminal() *string {
	return s.ArrivalTerminal
}

func (s *StandardSearchResponseBodyDataSolutionListJourneyListSegmentList) GetArrivalTime() *string {
	return s.ArrivalTime
}

func (s *StandardSearchResponseBodyDataSolutionListJourneyListSegmentList) GetAvailability() *string {
	return s.Availability
}

func (s *StandardSearchResponseBodyDataSolutionListJourneyListSegmentList) GetCabin() *string {
	return s.Cabin
}

func (s *StandardSearchResponseBodyDataSolutionListJourneyListSegmentList) GetCabinClass() *string {
	return s.CabinClass
}

func (s *StandardSearchResponseBodyDataSolutionListJourneyListSegmentList) GetCodeShare() *bool {
	return s.CodeShare
}

func (s *StandardSearchResponseBodyDataSolutionListJourneyListSegmentList) GetDepartureAirport() *string {
	return s.DepartureAirport
}

func (s *StandardSearchResponseBodyDataSolutionListJourneyListSegmentList) GetDepartureCity() *string {
	return s.DepartureCity
}

func (s *StandardSearchResponseBodyDataSolutionListJourneyListSegmentList) GetDepartureTerminal() *string {
	return s.DepartureTerminal
}

func (s *StandardSearchResponseBodyDataSolutionListJourneyListSegmentList) GetDepartureTime() *string {
	return s.DepartureTime
}

func (s *StandardSearchResponseBodyDataSolutionListJourneyListSegmentList) GetEquipType() *string {
	return s.EquipType
}

func (s *StandardSearchResponseBodyDataSolutionListJourneyListSegmentList) GetFlightDuration() *int32 {
	return s.FlightDuration
}

func (s *StandardSearchResponseBodyDataSolutionListJourneyListSegmentList) GetMarketingAirline() *string {
	return s.MarketingAirline
}

func (s *StandardSearchResponseBodyDataSolutionListJourneyListSegmentList) GetMarketingFlightNo() *string {
	return s.MarketingFlightNo
}

func (s *StandardSearchResponseBodyDataSolutionListJourneyListSegmentList) GetMarketingFlightNoInt() *int32 {
	return s.MarketingFlightNoInt
}

func (s *StandardSearchResponseBodyDataSolutionListJourneyListSegmentList) GetOperatingAirline() *string {
	return s.OperatingAirline
}

func (s *StandardSearchResponseBodyDataSolutionListJourneyListSegmentList) GetOperatingFlightNo() *string {
	return s.OperatingFlightNo
}

func (s *StandardSearchResponseBodyDataSolutionListJourneyListSegmentList) GetSegmentId() *string {
	return s.SegmentId
}

func (s *StandardSearchResponseBodyDataSolutionListJourneyListSegmentList) GetStopCityList() *string {
	return s.StopCityList
}

func (s *StandardSearchResponseBodyDataSolutionListJourneyListSegmentList) GetStopQuantity() *int32 {
	return s.StopQuantity
}

func (s *StandardSearchResponseBodyDataSolutionListJourneyListSegmentList) SetArrivalAirport(v string) *StandardSearchResponseBodyDataSolutionListJourneyListSegmentList {
	s.ArrivalAirport = &v
	return s
}

func (s *StandardSearchResponseBodyDataSolutionListJourneyListSegmentList) SetArrivalCity(v string) *StandardSearchResponseBodyDataSolutionListJourneyListSegmentList {
	s.ArrivalCity = &v
	return s
}

func (s *StandardSearchResponseBodyDataSolutionListJourneyListSegmentList) SetArrivalTerminal(v string) *StandardSearchResponseBodyDataSolutionListJourneyListSegmentList {
	s.ArrivalTerminal = &v
	return s
}

func (s *StandardSearchResponseBodyDataSolutionListJourneyListSegmentList) SetArrivalTime(v string) *StandardSearchResponseBodyDataSolutionListJourneyListSegmentList {
	s.ArrivalTime = &v
	return s
}

func (s *StandardSearchResponseBodyDataSolutionListJourneyListSegmentList) SetAvailability(v string) *StandardSearchResponseBodyDataSolutionListJourneyListSegmentList {
	s.Availability = &v
	return s
}

func (s *StandardSearchResponseBodyDataSolutionListJourneyListSegmentList) SetCabin(v string) *StandardSearchResponseBodyDataSolutionListJourneyListSegmentList {
	s.Cabin = &v
	return s
}

func (s *StandardSearchResponseBodyDataSolutionListJourneyListSegmentList) SetCabinClass(v string) *StandardSearchResponseBodyDataSolutionListJourneyListSegmentList {
	s.CabinClass = &v
	return s
}

func (s *StandardSearchResponseBodyDataSolutionListJourneyListSegmentList) SetCodeShare(v bool) *StandardSearchResponseBodyDataSolutionListJourneyListSegmentList {
	s.CodeShare = &v
	return s
}

func (s *StandardSearchResponseBodyDataSolutionListJourneyListSegmentList) SetDepartureAirport(v string) *StandardSearchResponseBodyDataSolutionListJourneyListSegmentList {
	s.DepartureAirport = &v
	return s
}

func (s *StandardSearchResponseBodyDataSolutionListJourneyListSegmentList) SetDepartureCity(v string) *StandardSearchResponseBodyDataSolutionListJourneyListSegmentList {
	s.DepartureCity = &v
	return s
}

func (s *StandardSearchResponseBodyDataSolutionListJourneyListSegmentList) SetDepartureTerminal(v string) *StandardSearchResponseBodyDataSolutionListJourneyListSegmentList {
	s.DepartureTerminal = &v
	return s
}

func (s *StandardSearchResponseBodyDataSolutionListJourneyListSegmentList) SetDepartureTime(v string) *StandardSearchResponseBodyDataSolutionListJourneyListSegmentList {
	s.DepartureTime = &v
	return s
}

func (s *StandardSearchResponseBodyDataSolutionListJourneyListSegmentList) SetEquipType(v string) *StandardSearchResponseBodyDataSolutionListJourneyListSegmentList {
	s.EquipType = &v
	return s
}

func (s *StandardSearchResponseBodyDataSolutionListJourneyListSegmentList) SetFlightDuration(v int32) *StandardSearchResponseBodyDataSolutionListJourneyListSegmentList {
	s.FlightDuration = &v
	return s
}

func (s *StandardSearchResponseBodyDataSolutionListJourneyListSegmentList) SetMarketingAirline(v string) *StandardSearchResponseBodyDataSolutionListJourneyListSegmentList {
	s.MarketingAirline = &v
	return s
}

func (s *StandardSearchResponseBodyDataSolutionListJourneyListSegmentList) SetMarketingFlightNo(v string) *StandardSearchResponseBodyDataSolutionListJourneyListSegmentList {
	s.MarketingFlightNo = &v
	return s
}

func (s *StandardSearchResponseBodyDataSolutionListJourneyListSegmentList) SetMarketingFlightNoInt(v int32) *StandardSearchResponseBodyDataSolutionListJourneyListSegmentList {
	s.MarketingFlightNoInt = &v
	return s
}

func (s *StandardSearchResponseBodyDataSolutionListJourneyListSegmentList) SetOperatingAirline(v string) *StandardSearchResponseBodyDataSolutionListJourneyListSegmentList {
	s.OperatingAirline = &v
	return s
}

func (s *StandardSearchResponseBodyDataSolutionListJourneyListSegmentList) SetOperatingFlightNo(v string) *StandardSearchResponseBodyDataSolutionListJourneyListSegmentList {
	s.OperatingFlightNo = &v
	return s
}

func (s *StandardSearchResponseBodyDataSolutionListJourneyListSegmentList) SetSegmentId(v string) *StandardSearchResponseBodyDataSolutionListJourneyListSegmentList {
	s.SegmentId = &v
	return s
}

func (s *StandardSearchResponseBodyDataSolutionListJourneyListSegmentList) SetStopCityList(v string) *StandardSearchResponseBodyDataSolutionListJourneyListSegmentList {
	s.StopCityList = &v
	return s
}

func (s *StandardSearchResponseBodyDataSolutionListJourneyListSegmentList) SetStopQuantity(v int32) *StandardSearchResponseBodyDataSolutionListJourneyListSegmentList {
	s.StopQuantity = &v
	return s
}

func (s *StandardSearchResponseBodyDataSolutionListJourneyListSegmentList) Validate() error {
	return dara.Validate(s)
}

type StandardSearchResponseBodyDataSolutionListSegmentBaggageCheckInInfoList struct {
	// example:
	//
	// 1
	LuggageDirectInfoType *int32    `json:"luggage_direct_info_type,omitempty" xml:"luggage_direct_info_type,omitempty"`
	SegmentIdList         []*string `json:"segment_id_list,omitempty" xml:"segment_id_list,omitempty" type:"Repeated"`
}

func (s StandardSearchResponseBodyDataSolutionListSegmentBaggageCheckInInfoList) String() string {
	return dara.Prettify(s)
}

func (s StandardSearchResponseBodyDataSolutionListSegmentBaggageCheckInInfoList) GoString() string {
	return s.String()
}

func (s *StandardSearchResponseBodyDataSolutionListSegmentBaggageCheckInInfoList) GetLuggageDirectInfoType() *int32 {
	return s.LuggageDirectInfoType
}

func (s *StandardSearchResponseBodyDataSolutionListSegmentBaggageCheckInInfoList) GetSegmentIdList() []*string {
	return s.SegmentIdList
}

func (s *StandardSearchResponseBodyDataSolutionListSegmentBaggageCheckInInfoList) SetLuggageDirectInfoType(v int32) *StandardSearchResponseBodyDataSolutionListSegmentBaggageCheckInInfoList {
	s.LuggageDirectInfoType = &v
	return s
}

func (s *StandardSearchResponseBodyDataSolutionListSegmentBaggageCheckInInfoList) SetSegmentIdList(v []*string) *StandardSearchResponseBodyDataSolutionListSegmentBaggageCheckInInfoList {
	s.SegmentIdList = v
	return s
}

func (s *StandardSearchResponseBodyDataSolutionListSegmentBaggageCheckInInfoList) Validate() error {
	return dara.Validate(s)
}

type StandardSearchResponseBodyDataSolutionListSegmentBaggageMappingList struct {
	PassengerBaggageAllowanceMapping map[string]*DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue `json:"passenger_baggage_allowance_mapping,omitempty" xml:"passenger_baggage_allowance_mapping,omitempty"`
	SegmentIdList                    []*string                                                                                  `json:"segment_id_list,omitempty" xml:"segment_id_list,omitempty" type:"Repeated"`
}

func (s StandardSearchResponseBodyDataSolutionListSegmentBaggageMappingList) String() string {
	return dara.Prettify(s)
}

func (s StandardSearchResponseBodyDataSolutionListSegmentBaggageMappingList) GoString() string {
	return s.String()
}

func (s *StandardSearchResponseBodyDataSolutionListSegmentBaggageMappingList) GetPassengerBaggageAllowanceMapping() map[string]*DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue {
	return s.PassengerBaggageAllowanceMapping
}

func (s *StandardSearchResponseBodyDataSolutionListSegmentBaggageMappingList) GetSegmentIdList() []*string {
	return s.SegmentIdList
}

func (s *StandardSearchResponseBodyDataSolutionListSegmentBaggageMappingList) SetPassengerBaggageAllowanceMapping(v map[string]*DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) *StandardSearchResponseBodyDataSolutionListSegmentBaggageMappingList {
	s.PassengerBaggageAllowanceMapping = v
	return s
}

func (s *StandardSearchResponseBodyDataSolutionListSegmentBaggageMappingList) SetSegmentIdList(v []*string) *StandardSearchResponseBodyDataSolutionListSegmentBaggageMappingList {
	s.SegmentIdList = v
	return s
}

func (s *StandardSearchResponseBodyDataSolutionListSegmentBaggageMappingList) Validate() error {
	return dara.Validate(s)
}

type StandardSearchResponseBodyDataSolutionListSegmentRefundChangeRuleMappingList struct {
	RefundChangeRuleMap map[string]*DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValue `json:"refund_change_rule_map,omitempty" xml:"refund_change_rule_map,omitempty"`
	SegmentIdList       []*string                                                                              `json:"segment_id_list,omitempty" xml:"segment_id_list,omitempty" type:"Repeated"`
}

func (s StandardSearchResponseBodyDataSolutionListSegmentRefundChangeRuleMappingList) String() string {
	return dara.Prettify(s)
}

func (s StandardSearchResponseBodyDataSolutionListSegmentRefundChangeRuleMappingList) GoString() string {
	return s.String()
}

func (s *StandardSearchResponseBodyDataSolutionListSegmentRefundChangeRuleMappingList) GetRefundChangeRuleMap() map[string]*DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValue {
	return s.RefundChangeRuleMap
}

func (s *StandardSearchResponseBodyDataSolutionListSegmentRefundChangeRuleMappingList) GetSegmentIdList() []*string {
	return s.SegmentIdList
}

func (s *StandardSearchResponseBodyDataSolutionListSegmentRefundChangeRuleMappingList) SetRefundChangeRuleMap(v map[string]*DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValue) *StandardSearchResponseBodyDataSolutionListSegmentRefundChangeRuleMappingList {
	s.RefundChangeRuleMap = v
	return s
}

func (s *StandardSearchResponseBodyDataSolutionListSegmentRefundChangeRuleMappingList) SetSegmentIdList(v []*string) *StandardSearchResponseBodyDataSolutionListSegmentRefundChangeRuleMappingList {
	s.SegmentIdList = v
	return s
}

func (s *StandardSearchResponseBodyDataSolutionListSegmentRefundChangeRuleMappingList) Validate() error {
	return dara.Validate(s)
}

type StandardSearchResponseBodyDataSolutionListSolutionAttribute struct {
	IssueTimeInfo *StandardSearchResponseBodyDataSolutionListSolutionAttributeIssueTimeInfo `json:"issue_time_info,omitempty" xml:"issue_time_info,omitempty" type:"Struct"`
	// example:
	//
	// 1
	SupplySourceType *string `json:"supply_source_type,omitempty" xml:"supply_source_type,omitempty"`
}

func (s StandardSearchResponseBodyDataSolutionListSolutionAttribute) String() string {
	return dara.Prettify(s)
}

func (s StandardSearchResponseBodyDataSolutionListSolutionAttribute) GoString() string {
	return s.String()
}

func (s *StandardSearchResponseBodyDataSolutionListSolutionAttribute) GetIssueTimeInfo() *StandardSearchResponseBodyDataSolutionListSolutionAttributeIssueTimeInfo {
	return s.IssueTimeInfo
}

func (s *StandardSearchResponseBodyDataSolutionListSolutionAttribute) GetSupplySourceType() *string {
	return s.SupplySourceType
}

func (s *StandardSearchResponseBodyDataSolutionListSolutionAttribute) SetIssueTimeInfo(v *StandardSearchResponseBodyDataSolutionListSolutionAttributeIssueTimeInfo) *StandardSearchResponseBodyDataSolutionListSolutionAttribute {
	s.IssueTimeInfo = v
	return s
}

func (s *StandardSearchResponseBodyDataSolutionListSolutionAttribute) SetSupplySourceType(v string) *StandardSearchResponseBodyDataSolutionListSolutionAttribute {
	s.SupplySourceType = &v
	return s
}

func (s *StandardSearchResponseBodyDataSolutionListSolutionAttribute) Validate() error {
	if s.IssueTimeInfo != nil {
		if err := s.IssueTimeInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type StandardSearchResponseBodyDataSolutionListSolutionAttributeIssueTimeInfo struct {
	IssueTicketType *int32 `json:"issue_ticket_type,omitempty" xml:"issue_ticket_type,omitempty"`
	IssueTimeLimit  *int32 `json:"issue_time_limit,omitempty" xml:"issue_time_limit,omitempty"`
}

func (s StandardSearchResponseBodyDataSolutionListSolutionAttributeIssueTimeInfo) String() string {
	return dara.Prettify(s)
}

func (s StandardSearchResponseBodyDataSolutionListSolutionAttributeIssueTimeInfo) GoString() string {
	return s.String()
}

func (s *StandardSearchResponseBodyDataSolutionListSolutionAttributeIssueTimeInfo) GetIssueTicketType() *int32 {
	return s.IssueTicketType
}

func (s *StandardSearchResponseBodyDataSolutionListSolutionAttributeIssueTimeInfo) GetIssueTimeLimit() *int32 {
	return s.IssueTimeLimit
}

func (s *StandardSearchResponseBodyDataSolutionListSolutionAttributeIssueTimeInfo) SetIssueTicketType(v int32) *StandardSearchResponseBodyDataSolutionListSolutionAttributeIssueTimeInfo {
	s.IssueTicketType = &v
	return s
}

func (s *StandardSearchResponseBodyDataSolutionListSolutionAttributeIssueTimeInfo) SetIssueTimeLimit(v int32) *StandardSearchResponseBodyDataSolutionListSolutionAttributeIssueTimeInfo {
	s.IssueTimeLimit = &v
	return s
}

func (s *StandardSearchResponseBodyDataSolutionListSolutionAttributeIssueTimeInfo) Validate() error {
	return dara.Validate(s)
}
