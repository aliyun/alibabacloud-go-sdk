// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnrichResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnrichResponseBody
  GetRequestId() *string 
  SetData(v *EnrichResponseBodyData) *EnrichResponseBody
  GetData() *EnrichResponseBodyData 
  SetErrorCode(v string) *EnrichResponseBody
  GetErrorCode() *string 
  SetErrorData(v interface{}) *EnrichResponseBody
  GetErrorData() interface{} 
  SetErrorMsg(v string) *EnrichResponseBody
  GetErrorMsg() *string 
  SetStatus(v int32) *EnrichResponseBody
  GetStatus() *int32 
  SetSuccess(v bool) *EnrichResponseBody
  GetSuccess() *bool 
}

type EnrichResponseBody struct {
  // request ID
  // 
  // example:
  // 
  // 51593418-8C73-5E47-8BA8-3F1D4A00CC0B
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // data
  Data *EnrichResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
  // error code
  // 
  // example:
  // 
  // null
  ErrorCode *string `json:"error_code,omitempty" xml:"error_code,omitempty"`
  // error data
  // 
  // example:
  // 
  // null
  ErrorData interface{} `json:"error_data,omitempty" xml:"error_data,omitempty"`
  // error message
  // 
  // example:
  // 
  // null
  ErrorMsg *string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
  // http reqeust has been processed successfully，status code is 200
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

func (s EnrichResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnrichResponseBody) GoString() string {
  return s.String()
}

func (s *EnrichResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnrichResponseBody) GetData() *EnrichResponseBodyData  {
  return s.Data
}

func (s *EnrichResponseBody) GetErrorCode() *string  {
  return s.ErrorCode
}

func (s *EnrichResponseBody) GetErrorData() interface{}  {
  return s.ErrorData
}

func (s *EnrichResponseBody) GetErrorMsg() *string  {
  return s.ErrorMsg
}

func (s *EnrichResponseBody) GetStatus() *int32  {
  return s.Status
}

func (s *EnrichResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnrichResponseBody) SetRequestId(v string) *EnrichResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnrichResponseBody) SetData(v *EnrichResponseBodyData) *EnrichResponseBody {
  s.Data = v
  return s
}

func (s *EnrichResponseBody) SetErrorCode(v string) *EnrichResponseBody {
  s.ErrorCode = &v
  return s
}

func (s *EnrichResponseBody) SetErrorData(v interface{}) *EnrichResponseBody {
  s.ErrorData = v
  return s
}

func (s *EnrichResponseBody) SetErrorMsg(v string) *EnrichResponseBody {
  s.ErrorMsg = &v
  return s
}

func (s *EnrichResponseBody) SetStatus(v int32) *EnrichResponseBody {
  s.Status = &v
  return s
}

func (s *EnrichResponseBody) SetSuccess(v bool) *EnrichResponseBody {
  s.Success = &v
  return s
}

func (s *EnrichResponseBody) Validate() error {
  if s.Data != nil {
    if err := s.Data.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type EnrichResponseBodyData struct {
  // solution list
  SolutionList []*EnrichResponseBodyDataSolutionList `json:"solution_list,omitempty" xml:"solution_list,omitempty" type:"Repeated"`
}

func (s EnrichResponseBodyData) String() string {
  return dara.Prettify(s)
}

func (s EnrichResponseBodyData) GoString() string {
  return s.String()
}

func (s *EnrichResponseBodyData) GetSolutionList() []*EnrichResponseBodyDataSolutionList  {
  return s.SolutionList
}

func (s *EnrichResponseBodyData) SetSolutionList(v []*EnrichResponseBodyDataSolutionList) *EnrichResponseBodyData {
  s.SolutionList = v
  return s
}

func (s *EnrichResponseBodyData) Validate() error {
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

type EnrichResponseBodyDataSolutionList struct {
  // adult fare
  // 
  // example:
  // 
  // 500
  AdultPrice *float64 `json:"adult_price,omitempty" xml:"adult_price,omitempty"`
  // adult tax
  // 
  // example:
  // 
  // 100
  AdultTax *float64 `json:"adult_tax,omitempty" xml:"adult_tax,omitempty"`
  // child fare
  // 
  // example:
  // 
  // 100
  ChildPrice *float64 `json:"child_price,omitempty" xml:"child_price,omitempty"`
  // child tax
  // 
  // example:
  // 
  // 100
  ChildTax *float64 `json:"child_tax,omitempty" xml:"child_tax,omitempty"`
  // infant fare
  // 
  // example:
  // 
  // 500
  InfantPrice *float64 `json:"infant_price,omitempty" xml:"infant_price,omitempty"`
  // infant tax
  // 
  // example:
  // 
  // 100
  InfantTax *float64 `json:"infant_tax,omitempty" xml:"infant_tax,omitempty"`
  // journey list
  JourneyList []*EnrichResponseBodyDataSolutionListJourneyList `json:"journey_list,omitempty" xml:"journey_list,omitempty" type:"Repeated"`
  // through check-in baggage  policy
  SegmentBaggageCheckInInfoList []*EnrichResponseBodyDataSolutionListSegmentBaggageCheckInInfoList `json:"segment_baggage_check_in_info_list,omitempty" xml:"segment_baggage_check_in_info_list,omitempty" type:"Repeated"`
  // baggage rule
  SegmentBaggageMappingList []*EnrichResponseBodyDataSolutionListSegmentBaggageMappingList `json:"segment_baggage_mapping_list,omitempty" xml:"segment_baggage_mapping_list,omitempty" type:"Repeated"`
  // change and refund policy
  SegmentRefundChangeRuleMappingList []*EnrichResponseBodyDataSolutionListSegmentRefundChangeRuleMappingList `json:"segment_refund_change_rule_mapping_list,omitempty" xml:"segment_refund_change_rule_mapping_list,omitempty" type:"Repeated"`
  // Quotation Attributes
  SolutionAttribute *EnrichResponseBodyDataSolutionListSolutionAttribute `json:"solution_attribute,omitempty" xml:"solution_attribute,omitempty" type:"Struct"`
  // solution ID
  // 
  // example:
  // 
  // eJwz8DeySEo0NjQ01TU3TU7TNTFINNO1SE5O0jVKM0hKNjEwTElLNYwz0A32cNT1dfPVNTIwMjYwNjRQ8/A3NLI01Q0Ic0cRBwBVFxJJ
  SolutionId *string `json:"solution_id,omitempty" xml:"solution_id,omitempty"`
}

func (s EnrichResponseBodyDataSolutionList) String() string {
  return dara.Prettify(s)
}

func (s EnrichResponseBodyDataSolutionList) GoString() string {
  return s.String()
}

func (s *EnrichResponseBodyDataSolutionList) GetAdultPrice() *float64  {
  return s.AdultPrice
}

func (s *EnrichResponseBodyDataSolutionList) GetAdultTax() *float64  {
  return s.AdultTax
}

func (s *EnrichResponseBodyDataSolutionList) GetChildPrice() *float64  {
  return s.ChildPrice
}

func (s *EnrichResponseBodyDataSolutionList) GetChildTax() *float64  {
  return s.ChildTax
}

func (s *EnrichResponseBodyDataSolutionList) GetInfantPrice() *float64  {
  return s.InfantPrice
}

func (s *EnrichResponseBodyDataSolutionList) GetInfantTax() *float64  {
  return s.InfantTax
}

func (s *EnrichResponseBodyDataSolutionList) GetJourneyList() []*EnrichResponseBodyDataSolutionListJourneyList  {
  return s.JourneyList
}

func (s *EnrichResponseBodyDataSolutionList) GetSegmentBaggageCheckInInfoList() []*EnrichResponseBodyDataSolutionListSegmentBaggageCheckInInfoList  {
  return s.SegmentBaggageCheckInInfoList
}

func (s *EnrichResponseBodyDataSolutionList) GetSegmentBaggageMappingList() []*EnrichResponseBodyDataSolutionListSegmentBaggageMappingList  {
  return s.SegmentBaggageMappingList
}

func (s *EnrichResponseBodyDataSolutionList) GetSegmentRefundChangeRuleMappingList() []*EnrichResponseBodyDataSolutionListSegmentRefundChangeRuleMappingList  {
  return s.SegmentRefundChangeRuleMappingList
}

func (s *EnrichResponseBodyDataSolutionList) GetSolutionAttribute() *EnrichResponseBodyDataSolutionListSolutionAttribute  {
  return s.SolutionAttribute
}

func (s *EnrichResponseBodyDataSolutionList) GetSolutionId() *string  {
  return s.SolutionId
}

func (s *EnrichResponseBodyDataSolutionList) SetAdultPrice(v float64) *EnrichResponseBodyDataSolutionList {
  s.AdultPrice = &v
  return s
}

func (s *EnrichResponseBodyDataSolutionList) SetAdultTax(v float64) *EnrichResponseBodyDataSolutionList {
  s.AdultTax = &v
  return s
}

func (s *EnrichResponseBodyDataSolutionList) SetChildPrice(v float64) *EnrichResponseBodyDataSolutionList {
  s.ChildPrice = &v
  return s
}

func (s *EnrichResponseBodyDataSolutionList) SetChildTax(v float64) *EnrichResponseBodyDataSolutionList {
  s.ChildTax = &v
  return s
}

func (s *EnrichResponseBodyDataSolutionList) SetInfantPrice(v float64) *EnrichResponseBodyDataSolutionList {
  s.InfantPrice = &v
  return s
}

func (s *EnrichResponseBodyDataSolutionList) SetInfantTax(v float64) *EnrichResponseBodyDataSolutionList {
  s.InfantTax = &v
  return s
}

func (s *EnrichResponseBodyDataSolutionList) SetJourneyList(v []*EnrichResponseBodyDataSolutionListJourneyList) *EnrichResponseBodyDataSolutionList {
  s.JourneyList = v
  return s
}

func (s *EnrichResponseBodyDataSolutionList) SetSegmentBaggageCheckInInfoList(v []*EnrichResponseBodyDataSolutionListSegmentBaggageCheckInInfoList) *EnrichResponseBodyDataSolutionList {
  s.SegmentBaggageCheckInInfoList = v
  return s
}

func (s *EnrichResponseBodyDataSolutionList) SetSegmentBaggageMappingList(v []*EnrichResponseBodyDataSolutionListSegmentBaggageMappingList) *EnrichResponseBodyDataSolutionList {
  s.SegmentBaggageMappingList = v
  return s
}

func (s *EnrichResponseBodyDataSolutionList) SetSegmentRefundChangeRuleMappingList(v []*EnrichResponseBodyDataSolutionListSegmentRefundChangeRuleMappingList) *EnrichResponseBodyDataSolutionList {
  s.SegmentRefundChangeRuleMappingList = v
  return s
}

func (s *EnrichResponseBodyDataSolutionList) SetSolutionAttribute(v *EnrichResponseBodyDataSolutionListSolutionAttribute) *EnrichResponseBodyDataSolutionList {
  s.SolutionAttribute = v
  return s
}

func (s *EnrichResponseBodyDataSolutionList) SetSolutionId(v string) *EnrichResponseBodyDataSolutionList {
  s.SolutionId = &v
  return s
}

func (s *EnrichResponseBodyDataSolutionList) Validate() error {
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

type EnrichResponseBodyDataSolutionListJourneyList struct {
  // segment Info
  SegmentList []*EnrichResponseBodyDataSolutionListJourneyListSegmentList `json:"segment_list,omitempty" xml:"segment_list,omitempty" type:"Repeated"`
  // number of transfers
  // 
  // example:
  // 
  // 0
  TransferCount *int32 `json:"transfer_count,omitempty" xml:"transfer_count,omitempty"`
}

func (s EnrichResponseBodyDataSolutionListJourneyList) String() string {
  return dara.Prettify(s)
}

func (s EnrichResponseBodyDataSolutionListJourneyList) GoString() string {
  return s.String()
}

func (s *EnrichResponseBodyDataSolutionListJourneyList) GetSegmentList() []*EnrichResponseBodyDataSolutionListJourneyListSegmentList  {
  return s.SegmentList
}

func (s *EnrichResponseBodyDataSolutionListJourneyList) GetTransferCount() *int32  {
  return s.TransferCount
}

func (s *EnrichResponseBodyDataSolutionListJourneyList) SetSegmentList(v []*EnrichResponseBodyDataSolutionListJourneyListSegmentList) *EnrichResponseBodyDataSolutionListJourneyList {
  s.SegmentList = v
  return s
}

func (s *EnrichResponseBodyDataSolutionListJourneyList) SetTransferCount(v int32) *EnrichResponseBodyDataSolutionListJourneyList {
  s.TransferCount = &v
  return s
}

func (s *EnrichResponseBodyDataSolutionListJourneyList) Validate() error {
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

type EnrichResponseBodyDataSolutionListJourneyListSegmentList struct {
  // arrival airport code
  // 
  // example:
  // 
  // MFM
  ArrivalAirport *string `json:"arrival_airport,omitempty" xml:"arrival_airport,omitempty"`
  // arrival city code
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
  // arrival time (yyyy-MM-dd HH:mm:ss)
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
  // departure airport code
  // 
  // example:
  // 
  // PVG
  DepartureAirport *string `json:"departure_airport,omitempty" xml:"departure_airport,omitempty"`
  // departure city code
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
  // departure time (yyyy-MM-dd HH:mm:ss)
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
  // marketing airline code (eg: KA)
  // 
  // example:
  // 
  // HO
  MarketingAirline *string `json:"marketing_airline,omitempty" xml:"marketing_airline,omitempty"`
  // marketing airline flight no. (eg: KA5809)
  // 
  // example:
  // 
  // HO1295
  MarketingFlightNo *string `json:"marketing_flight_no,omitempty" xml:"marketing_flight_no,omitempty"`
  // marketing airline integer flight no. (eg: 5809)
  // 
  // example:
  // 
  // 1295
  MarketingFlightNoInt *int32 `json:"marketing_flight_no_int,omitempty" xml:"marketing_flight_no_int,omitempty"`
  // operating airline code (eg: CX)
  // 
  // example:
  // 
  // HO
  OperatingAirline *string `json:"operating_airline,omitempty" xml:"operating_airline,omitempty"`
  // operating airline flight no. (eg: CX601)
  // 
  // example:
  // 
  // HO1295
  OperatingFlightNo *string `json:"operating_flight_no,omitempty" xml:"operating_flight_no,omitempty"`
  // segment ID: flight no+departure airport+arrival airport+departure time(MMdd)
  // 
  // example:
  // 
  // HO1295-PVG-MFM-20230310
  SegmentId *string `json:"segment_id,omitempty" xml:"segment_id,omitempty"`
  // stop city list. 
  // 
  // when stop_quantity > 1 , use “,” for seperation
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

func (s EnrichResponseBodyDataSolutionListJourneyListSegmentList) String() string {
  return dara.Prettify(s)
}

func (s EnrichResponseBodyDataSolutionListJourneyListSegmentList) GoString() string {
  return s.String()
}

func (s *EnrichResponseBodyDataSolutionListJourneyListSegmentList) GetArrivalAirport() *string  {
  return s.ArrivalAirport
}

func (s *EnrichResponseBodyDataSolutionListJourneyListSegmentList) GetArrivalCity() *string  {
  return s.ArrivalCity
}

func (s *EnrichResponseBodyDataSolutionListJourneyListSegmentList) GetArrivalTerminal() *string  {
  return s.ArrivalTerminal
}

func (s *EnrichResponseBodyDataSolutionListJourneyListSegmentList) GetArrivalTime() *string  {
  return s.ArrivalTime
}

func (s *EnrichResponseBodyDataSolutionListJourneyListSegmentList) GetAvailability() *string  {
  return s.Availability
}

func (s *EnrichResponseBodyDataSolutionListJourneyListSegmentList) GetCabin() *string  {
  return s.Cabin
}

func (s *EnrichResponseBodyDataSolutionListJourneyListSegmentList) GetCabinClass() *string  {
  return s.CabinClass
}

func (s *EnrichResponseBodyDataSolutionListJourneyListSegmentList) GetCodeShare() *bool  {
  return s.CodeShare
}

func (s *EnrichResponseBodyDataSolutionListJourneyListSegmentList) GetDepartureAirport() *string  {
  return s.DepartureAirport
}

func (s *EnrichResponseBodyDataSolutionListJourneyListSegmentList) GetDepartureCity() *string  {
  return s.DepartureCity
}

func (s *EnrichResponseBodyDataSolutionListJourneyListSegmentList) GetDepartureTerminal() *string  {
  return s.DepartureTerminal
}

func (s *EnrichResponseBodyDataSolutionListJourneyListSegmentList) GetDepartureTime() *string  {
  return s.DepartureTime
}

func (s *EnrichResponseBodyDataSolutionListJourneyListSegmentList) GetEquipType() *string  {
  return s.EquipType
}

func (s *EnrichResponseBodyDataSolutionListJourneyListSegmentList) GetFlightDuration() *int32  {
  return s.FlightDuration
}

func (s *EnrichResponseBodyDataSolutionListJourneyListSegmentList) GetMarketingAirline() *string  {
  return s.MarketingAirline
}

func (s *EnrichResponseBodyDataSolutionListJourneyListSegmentList) GetMarketingFlightNo() *string  {
  return s.MarketingFlightNo
}

func (s *EnrichResponseBodyDataSolutionListJourneyListSegmentList) GetMarketingFlightNoInt() *int32  {
  return s.MarketingFlightNoInt
}

func (s *EnrichResponseBodyDataSolutionListJourneyListSegmentList) GetOperatingAirline() *string  {
  return s.OperatingAirline
}

func (s *EnrichResponseBodyDataSolutionListJourneyListSegmentList) GetOperatingFlightNo() *string  {
  return s.OperatingFlightNo
}

func (s *EnrichResponseBodyDataSolutionListJourneyListSegmentList) GetSegmentId() *string  {
  return s.SegmentId
}

func (s *EnrichResponseBodyDataSolutionListJourneyListSegmentList) GetStopCityList() *string  {
  return s.StopCityList
}

func (s *EnrichResponseBodyDataSolutionListJourneyListSegmentList) GetStopQuantity() *int32  {
  return s.StopQuantity
}

func (s *EnrichResponseBodyDataSolutionListJourneyListSegmentList) SetArrivalAirport(v string) *EnrichResponseBodyDataSolutionListJourneyListSegmentList {
  s.ArrivalAirport = &v
  return s
}

func (s *EnrichResponseBodyDataSolutionListJourneyListSegmentList) SetArrivalCity(v string) *EnrichResponseBodyDataSolutionListJourneyListSegmentList {
  s.ArrivalCity = &v
  return s
}

func (s *EnrichResponseBodyDataSolutionListJourneyListSegmentList) SetArrivalTerminal(v string) *EnrichResponseBodyDataSolutionListJourneyListSegmentList {
  s.ArrivalTerminal = &v
  return s
}

func (s *EnrichResponseBodyDataSolutionListJourneyListSegmentList) SetArrivalTime(v string) *EnrichResponseBodyDataSolutionListJourneyListSegmentList {
  s.ArrivalTime = &v
  return s
}

func (s *EnrichResponseBodyDataSolutionListJourneyListSegmentList) SetAvailability(v string) *EnrichResponseBodyDataSolutionListJourneyListSegmentList {
  s.Availability = &v
  return s
}

func (s *EnrichResponseBodyDataSolutionListJourneyListSegmentList) SetCabin(v string) *EnrichResponseBodyDataSolutionListJourneyListSegmentList {
  s.Cabin = &v
  return s
}

func (s *EnrichResponseBodyDataSolutionListJourneyListSegmentList) SetCabinClass(v string) *EnrichResponseBodyDataSolutionListJourneyListSegmentList {
  s.CabinClass = &v
  return s
}

func (s *EnrichResponseBodyDataSolutionListJourneyListSegmentList) SetCodeShare(v bool) *EnrichResponseBodyDataSolutionListJourneyListSegmentList {
  s.CodeShare = &v
  return s
}

func (s *EnrichResponseBodyDataSolutionListJourneyListSegmentList) SetDepartureAirport(v string) *EnrichResponseBodyDataSolutionListJourneyListSegmentList {
  s.DepartureAirport = &v
  return s
}

func (s *EnrichResponseBodyDataSolutionListJourneyListSegmentList) SetDepartureCity(v string) *EnrichResponseBodyDataSolutionListJourneyListSegmentList {
  s.DepartureCity = &v
  return s
}

func (s *EnrichResponseBodyDataSolutionListJourneyListSegmentList) SetDepartureTerminal(v string) *EnrichResponseBodyDataSolutionListJourneyListSegmentList {
  s.DepartureTerminal = &v
  return s
}

func (s *EnrichResponseBodyDataSolutionListJourneyListSegmentList) SetDepartureTime(v string) *EnrichResponseBodyDataSolutionListJourneyListSegmentList {
  s.DepartureTime = &v
  return s
}

func (s *EnrichResponseBodyDataSolutionListJourneyListSegmentList) SetEquipType(v string) *EnrichResponseBodyDataSolutionListJourneyListSegmentList {
  s.EquipType = &v
  return s
}

func (s *EnrichResponseBodyDataSolutionListJourneyListSegmentList) SetFlightDuration(v int32) *EnrichResponseBodyDataSolutionListJourneyListSegmentList {
  s.FlightDuration = &v
  return s
}

func (s *EnrichResponseBodyDataSolutionListJourneyListSegmentList) SetMarketingAirline(v string) *EnrichResponseBodyDataSolutionListJourneyListSegmentList {
  s.MarketingAirline = &v
  return s
}

func (s *EnrichResponseBodyDataSolutionListJourneyListSegmentList) SetMarketingFlightNo(v string) *EnrichResponseBodyDataSolutionListJourneyListSegmentList {
  s.MarketingFlightNo = &v
  return s
}

func (s *EnrichResponseBodyDataSolutionListJourneyListSegmentList) SetMarketingFlightNoInt(v int32) *EnrichResponseBodyDataSolutionListJourneyListSegmentList {
  s.MarketingFlightNoInt = &v
  return s
}

func (s *EnrichResponseBodyDataSolutionListJourneyListSegmentList) SetOperatingAirline(v string) *EnrichResponseBodyDataSolutionListJourneyListSegmentList {
  s.OperatingAirline = &v
  return s
}

func (s *EnrichResponseBodyDataSolutionListJourneyListSegmentList) SetOperatingFlightNo(v string) *EnrichResponseBodyDataSolutionListJourneyListSegmentList {
  s.OperatingFlightNo = &v
  return s
}

func (s *EnrichResponseBodyDataSolutionListJourneyListSegmentList) SetSegmentId(v string) *EnrichResponseBodyDataSolutionListJourneyListSegmentList {
  s.SegmentId = &v
  return s
}

func (s *EnrichResponseBodyDataSolutionListJourneyListSegmentList) SetStopCityList(v string) *EnrichResponseBodyDataSolutionListJourneyListSegmentList {
  s.StopCityList = &v
  return s
}

func (s *EnrichResponseBodyDataSolutionListJourneyListSegmentList) SetStopQuantity(v int32) *EnrichResponseBodyDataSolutionListJourneyListSegmentList {
  s.StopQuantity = &v
  return s
}

func (s *EnrichResponseBodyDataSolutionListJourneyListSegmentList) Validate() error {
  return dara.Validate(s)
}

type EnrichResponseBodyDataSolutionListSegmentBaggageCheckInInfoList struct {
  // through check-in baggage policy type
  // 
  // 1. baggage through check-in between segments
  // 
  // 2. baggage re-check-in needed between segments
  // 
  // 4. baggage through check-in at stop city ( applies for stop flight )
  // 
  // 3. baggage re-checkin needed at stop city ( applies for stop flight )
  // 
  // example:
  // 
  // 1
  LuggageDirectInfoType *int32 `json:"luggage_direct_info_type,omitempty" xml:"luggage_direct_info_type,omitempty"`
  // segment id list. all the listed segment ids share the same baggage through check-in policy
  SegmentIdList []*string `json:"segment_id_list,omitempty" xml:"segment_id_list,omitempty" type:"Repeated"`
}

func (s EnrichResponseBodyDataSolutionListSegmentBaggageCheckInInfoList) String() string {
  return dara.Prettify(s)
}

func (s EnrichResponseBodyDataSolutionListSegmentBaggageCheckInInfoList) GoString() string {
  return s.String()
}

func (s *EnrichResponseBodyDataSolutionListSegmentBaggageCheckInInfoList) GetLuggageDirectInfoType() *int32  {
  return s.LuggageDirectInfoType
}

func (s *EnrichResponseBodyDataSolutionListSegmentBaggageCheckInInfoList) GetSegmentIdList() []*string  {
  return s.SegmentIdList
}

func (s *EnrichResponseBodyDataSolutionListSegmentBaggageCheckInInfoList) SetLuggageDirectInfoType(v int32) *EnrichResponseBodyDataSolutionListSegmentBaggageCheckInInfoList {
  s.LuggageDirectInfoType = &v
  return s
}

func (s *EnrichResponseBodyDataSolutionListSegmentBaggageCheckInInfoList) SetSegmentIdList(v []*string) *EnrichResponseBodyDataSolutionListSegmentBaggageCheckInInfoList {
  s.SegmentIdList = v
  return s
}

func (s *EnrichResponseBodyDataSolutionListSegmentBaggageCheckInInfoList) Validate() error {
  return dara.Validate(s)
}

type EnrichResponseBodyDataSolutionListSegmentBaggageMappingList struct {
  // baggage rule mapping, key is passenger type, value is baggage allowance details
  PassengerBaggageAllowanceMapping map[string]*DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue `json:"passenger_baggage_allowance_mapping,omitempty" xml:"passenger_baggage_allowance_mapping,omitempty"`
  // segment id list. 
  // 
  // all the listed segment ids share the same baggage rule
  SegmentIdList []*string `json:"segment_id_list,omitempty" xml:"segment_id_list,omitempty" type:"Repeated"`
}

func (s EnrichResponseBodyDataSolutionListSegmentBaggageMappingList) String() string {
  return dara.Prettify(s)
}

func (s EnrichResponseBodyDataSolutionListSegmentBaggageMappingList) GoString() string {
  return s.String()
}

func (s *EnrichResponseBodyDataSolutionListSegmentBaggageMappingList) GetPassengerBaggageAllowanceMapping() map[string]*DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue  {
  return s.PassengerBaggageAllowanceMapping
}

func (s *EnrichResponseBodyDataSolutionListSegmentBaggageMappingList) GetSegmentIdList() []*string  {
  return s.SegmentIdList
}

func (s *EnrichResponseBodyDataSolutionListSegmentBaggageMappingList) SetPassengerBaggageAllowanceMapping(v map[string]*DataSolutionListSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) *EnrichResponseBodyDataSolutionListSegmentBaggageMappingList {
  s.PassengerBaggageAllowanceMapping = v
  return s
}

func (s *EnrichResponseBodyDataSolutionListSegmentBaggageMappingList) SetSegmentIdList(v []*string) *EnrichResponseBodyDataSolutionListSegmentBaggageMappingList {
  s.SegmentIdList = v
  return s
}

func (s *EnrichResponseBodyDataSolutionListSegmentBaggageMappingList) Validate() error {
  return dara.Validate(s)
}

type EnrichResponseBodyDataSolutionListSegmentRefundChangeRuleMappingList struct {
  // change and refund policy mapping, key is passenger type, value is change and refund policy detail
  RefundChangeRuleMap map[string]*DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValue `json:"refund_change_rule_map,omitempty" xml:"refund_change_rule_map,omitempty"`
  // segment id list. all the listed segment ids share the same change and refund policy
  SegmentIdList []*string `json:"segment_id_list,omitempty" xml:"segment_id_list,omitempty" type:"Repeated"`
}

func (s EnrichResponseBodyDataSolutionListSegmentRefundChangeRuleMappingList) String() string {
  return dara.Prettify(s)
}

func (s EnrichResponseBodyDataSolutionListSegmentRefundChangeRuleMappingList) GoString() string {
  return s.String()
}

func (s *EnrichResponseBodyDataSolutionListSegmentRefundChangeRuleMappingList) GetRefundChangeRuleMap() map[string]*DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValue  {
  return s.RefundChangeRuleMap
}

func (s *EnrichResponseBodyDataSolutionListSegmentRefundChangeRuleMappingList) GetSegmentIdList() []*string  {
  return s.SegmentIdList
}

func (s *EnrichResponseBodyDataSolutionListSegmentRefundChangeRuleMappingList) SetRefundChangeRuleMap(v map[string]*DataSolutionListSegmentRefundChangeRuleMappingListRefundChangeRuleMapValue) *EnrichResponseBodyDataSolutionListSegmentRefundChangeRuleMappingList {
  s.RefundChangeRuleMap = v
  return s
}

func (s *EnrichResponseBodyDataSolutionListSegmentRefundChangeRuleMappingList) SetSegmentIdList(v []*string) *EnrichResponseBodyDataSolutionListSegmentRefundChangeRuleMappingList {
  s.SegmentIdList = v
  return s
}

func (s *EnrichResponseBodyDataSolutionListSegmentRefundChangeRuleMappingList) Validate() error {
  return dara.Validate(s)
}

type EnrichResponseBodyDataSolutionListSolutionAttribute struct {
  IssueTimeInfo *EnrichResponseBodyDataSolutionListSolutionAttributeIssueTimeInfo `json:"issue_time_info,omitempty" xml:"issue_time_info,omitempty" type:"Struct"`
  // Supply source type 1:self-operated; 2:agent; 3:flagship store
  // 
  // example:
  // 
  // 1
  SupplySourceType *string `json:"supply_source_type,omitempty" xml:"supply_source_type,omitempty"`
}

func (s EnrichResponseBodyDataSolutionListSolutionAttribute) String() string {
  return dara.Prettify(s)
}

func (s EnrichResponseBodyDataSolutionListSolutionAttribute) GoString() string {
  return s.String()
}

func (s *EnrichResponseBodyDataSolutionListSolutionAttribute) GetIssueTimeInfo() *EnrichResponseBodyDataSolutionListSolutionAttributeIssueTimeInfo  {
  return s.IssueTimeInfo
}

func (s *EnrichResponseBodyDataSolutionListSolutionAttribute) GetSupplySourceType() *string  {
  return s.SupplySourceType
}

func (s *EnrichResponseBodyDataSolutionListSolutionAttribute) SetIssueTimeInfo(v *EnrichResponseBodyDataSolutionListSolutionAttributeIssueTimeInfo) *EnrichResponseBodyDataSolutionListSolutionAttribute {
  s.IssueTimeInfo = v
  return s
}

func (s *EnrichResponseBodyDataSolutionListSolutionAttribute) SetSupplySourceType(v string) *EnrichResponseBodyDataSolutionListSolutionAttribute {
  s.SupplySourceType = &v
  return s
}

func (s *EnrichResponseBodyDataSolutionListSolutionAttribute) Validate() error {
  if s.IssueTimeInfo != nil {
    if err := s.IssueTimeInfo.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type EnrichResponseBodyDataSolutionListSolutionAttributeIssueTimeInfo struct {
  IssueTicketType *int32 `json:"issue_ticket_type,omitempty" xml:"issue_ticket_type,omitempty"`
  IssueTimeLimit *int32 `json:"issue_time_limit,omitempty" xml:"issue_time_limit,omitempty"`
}

func (s EnrichResponseBodyDataSolutionListSolutionAttributeIssueTimeInfo) String() string {
  return dara.Prettify(s)
}

func (s EnrichResponseBodyDataSolutionListSolutionAttributeIssueTimeInfo) GoString() string {
  return s.String()
}

func (s *EnrichResponseBodyDataSolutionListSolutionAttributeIssueTimeInfo) GetIssueTicketType() *int32  {
  return s.IssueTicketType
}

func (s *EnrichResponseBodyDataSolutionListSolutionAttributeIssueTimeInfo) GetIssueTimeLimit() *int32  {
  return s.IssueTimeLimit
}

func (s *EnrichResponseBodyDataSolutionListSolutionAttributeIssueTimeInfo) SetIssueTicketType(v int32) *EnrichResponseBodyDataSolutionListSolutionAttributeIssueTimeInfo {
  s.IssueTicketType = &v
  return s
}

func (s *EnrichResponseBodyDataSolutionListSolutionAttributeIssueTimeInfo) SetIssueTimeLimit(v int32) *EnrichResponseBodyDataSolutionListSolutionAttributeIssueTimeInfo {
  s.IssueTimeLimit = &v
  return s
}

func (s *EnrichResponseBodyDataSolutionListSolutionAttributeIssueTimeInfo) Validate() error {
  return dara.Validate(s)
}

