// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnrichRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAdults(v int32) *EnrichRequest
  GetAdults() *int32 
  SetCabinClass(v string) *EnrichRequest
  GetCabinClass() *string 
  SetChildren(v int32) *EnrichRequest
  GetChildren() *int32 
  SetInfants(v int32) *EnrichRequest
  GetInfants() *int32 
  SetJourneyParamList(v []*EnrichRequestJourneyParamList) *EnrichRequest
  GetJourneyParamList() []*EnrichRequestJourneyParamList 
  SetSolutionId(v string) *EnrichRequest
  GetSolutionId() *string 
}

type EnrichRequest struct {
  // adult passenger amount 1-9
  // 
  // example:
  // 
  // 1
  Adults *int32 `json:"adults,omitempty" xml:"adults,omitempty"`
  // cabin class
  // 
  // 1. **ALL_CABIN*	- : all cabin class
  // 
  // 2. **Y*	- : economy class
  // 
  // 3. **FC*	- : first class and business class
  // 
  // 4. **S*	- : premium economy class
  // 
  // 5. **YS*	- : economy class and premium economy class
  // 
  // 6. **YSC*	- : economy class, premium economy class and business class
  // 
  // example:
  // 
  // ALL_CABIN
  CabinClass *string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
  // child passenger amount 0-9
  // 
  // example:
  // 
  // 1
  Children *int32 `json:"children,omitempty" xml:"children,omitempty"`
  // infant passenger amount 0-9
  // 
  // example:
  // 
  // 1
  Infants *int32 `json:"infants,omitempty" xml:"infants,omitempty"`
  // journey list
  JourneyParamList []*EnrichRequestJourneyParamList `json:"journey_param_list,omitempty" xml:"journey_param_list,omitempty" type:"Repeated"`
  // solution_id returned by Search
  // 
  // example:
  // 
  // eJwz8DeySEo0NjQ01TU3TU7TNTFINNO1SE5O0jVKM0hKNjEwTElLNYwz0A32cNT1dfPVNTIwMjYwNjRQ8/A3NLI01Q0Ic0cRBwBVFxJJ
  SolutionId *string `json:"solution_id,omitempty" xml:"solution_id,omitempty"`
}

func (s EnrichRequest) String() string {
  return dara.Prettify(s)
}

func (s EnrichRequest) GoString() string {
  return s.String()
}

func (s *EnrichRequest) GetAdults() *int32  {
  return s.Adults
}

func (s *EnrichRequest) GetCabinClass() *string  {
  return s.CabinClass
}

func (s *EnrichRequest) GetChildren() *int32  {
  return s.Children
}

func (s *EnrichRequest) GetInfants() *int32  {
  return s.Infants
}

func (s *EnrichRequest) GetJourneyParamList() []*EnrichRequestJourneyParamList  {
  return s.JourneyParamList
}

func (s *EnrichRequest) GetSolutionId() *string  {
  return s.SolutionId
}

func (s *EnrichRequest) SetAdults(v int32) *EnrichRequest {
  s.Adults = &v
  return s
}

func (s *EnrichRequest) SetCabinClass(v string) *EnrichRequest {
  s.CabinClass = &v
  return s
}

func (s *EnrichRequest) SetChildren(v int32) *EnrichRequest {
  s.Children = &v
  return s
}

func (s *EnrichRequest) SetInfants(v int32) *EnrichRequest {
  s.Infants = &v
  return s
}

func (s *EnrichRequest) SetJourneyParamList(v []*EnrichRequestJourneyParamList) *EnrichRequest {
  s.JourneyParamList = v
  return s
}

func (s *EnrichRequest) SetSolutionId(v string) *EnrichRequest {
  s.SolutionId = &v
  return s
}

func (s *EnrichRequest) Validate() error {
  if s.JourneyParamList != nil {
    for _, item := range s.JourneyParamList {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

type EnrichRequestJourneyParamList struct {
  // arrival city code
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // MFM
  ArrivalCity *string `json:"arrival_city,omitempty" xml:"arrival_city,omitempty"`
  // departure city code
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // SHA
  DepartureCity *string `json:"departure_city,omitempty" xml:"departure_city,omitempty"`
  // departure date (eg: yyyyMMdd)
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 20230310
  DepartureDate *string `json:"departure_date,omitempty" xml:"departure_date,omitempty"`
  // segement param list
  // 
  // This parameter is required.
  SegmentParamList []*EnrichRequestJourneyParamListSegmentParamList `json:"segment_param_list,omitempty" xml:"segment_param_list,omitempty" type:"Repeated"`
}

func (s EnrichRequestJourneyParamList) String() string {
  return dara.Prettify(s)
}

func (s EnrichRequestJourneyParamList) GoString() string {
  return s.String()
}

func (s *EnrichRequestJourneyParamList) GetArrivalCity() *string  {
  return s.ArrivalCity
}

func (s *EnrichRequestJourneyParamList) GetDepartureCity() *string  {
  return s.DepartureCity
}

func (s *EnrichRequestJourneyParamList) GetDepartureDate() *string  {
  return s.DepartureDate
}

func (s *EnrichRequestJourneyParamList) GetSegmentParamList() []*EnrichRequestJourneyParamListSegmentParamList  {
  return s.SegmentParamList
}

func (s *EnrichRequestJourneyParamList) SetArrivalCity(v string) *EnrichRequestJourneyParamList {
  s.ArrivalCity = &v
  return s
}

func (s *EnrichRequestJourneyParamList) SetDepartureCity(v string) *EnrichRequestJourneyParamList {
  s.DepartureCity = &v
  return s
}

func (s *EnrichRequestJourneyParamList) SetDepartureDate(v string) *EnrichRequestJourneyParamList {
  s.DepartureDate = &v
  return s
}

func (s *EnrichRequestJourneyParamList) SetSegmentParamList(v []*EnrichRequestJourneyParamListSegmentParamList) *EnrichRequestJourneyParamList {
  s.SegmentParamList = v
  return s
}

func (s *EnrichRequestJourneyParamList) Validate() error {
  if s.SegmentParamList != nil {
    for _, item := range s.SegmentParamList {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

type EnrichRequestJourneyParamListSegmentParamList struct {
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
  // RBD
  // 
  // example:
  // 
  // V
  Cabin *string `json:"cabin,omitempty" xml:"cabin,omitempty"`
  // child RBD
  // 
  // example:
  // 
  // E
  ChildCabin *string `json:"child_cabin,omitempty" xml:"child_cabin,omitempty"`
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
  // example:
  // 
  // 2023-03-10
  DepartureDate *string `json:"departure_date,omitempty" xml:"departure_date,omitempty"`
  // departure time in string format (yyyy-MM-dd HH:mm:ss)
  // 
  // example:
  // 
  // 2023-03-10 07:55:00
  DepartureTime *string `json:"departure_time,omitempty" xml:"departure_time,omitempty"`
  // marketing flight no. (eg: KA5809)
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // HO1295
  MarketingFlightNo *string `json:"marketing_flight_no,omitempty" xml:"marketing_flight_no,omitempty"`
}

func (s EnrichRequestJourneyParamListSegmentParamList) String() string {
  return dara.Prettify(s)
}

func (s EnrichRequestJourneyParamListSegmentParamList) GoString() string {
  return s.String()
}

func (s *EnrichRequestJourneyParamListSegmentParamList) GetArrivalAirport() *string  {
  return s.ArrivalAirport
}

func (s *EnrichRequestJourneyParamListSegmentParamList) GetArrivalCity() *string  {
  return s.ArrivalCity
}

func (s *EnrichRequestJourneyParamListSegmentParamList) GetCabin() *string  {
  return s.Cabin
}

func (s *EnrichRequestJourneyParamListSegmentParamList) GetChildCabin() *string  {
  return s.ChildCabin
}

func (s *EnrichRequestJourneyParamListSegmentParamList) GetDepartureAirport() *string  {
  return s.DepartureAirport
}

func (s *EnrichRequestJourneyParamListSegmentParamList) GetDepartureCity() *string  {
  return s.DepartureCity
}

func (s *EnrichRequestJourneyParamListSegmentParamList) GetDepartureDate() *string  {
  return s.DepartureDate
}

func (s *EnrichRequestJourneyParamListSegmentParamList) GetDepartureTime() *string  {
  return s.DepartureTime
}

func (s *EnrichRequestJourneyParamListSegmentParamList) GetMarketingFlightNo() *string  {
  return s.MarketingFlightNo
}

func (s *EnrichRequestJourneyParamListSegmentParamList) SetArrivalAirport(v string) *EnrichRequestJourneyParamListSegmentParamList {
  s.ArrivalAirport = &v
  return s
}

func (s *EnrichRequestJourneyParamListSegmentParamList) SetArrivalCity(v string) *EnrichRequestJourneyParamListSegmentParamList {
  s.ArrivalCity = &v
  return s
}

func (s *EnrichRequestJourneyParamListSegmentParamList) SetCabin(v string) *EnrichRequestJourneyParamListSegmentParamList {
  s.Cabin = &v
  return s
}

func (s *EnrichRequestJourneyParamListSegmentParamList) SetChildCabin(v string) *EnrichRequestJourneyParamListSegmentParamList {
  s.ChildCabin = &v
  return s
}

func (s *EnrichRequestJourneyParamListSegmentParamList) SetDepartureAirport(v string) *EnrichRequestJourneyParamListSegmentParamList {
  s.DepartureAirport = &v
  return s
}

func (s *EnrichRequestJourneyParamListSegmentParamList) SetDepartureCity(v string) *EnrichRequestJourneyParamListSegmentParamList {
  s.DepartureCity = &v
  return s
}

func (s *EnrichRequestJourneyParamListSegmentParamList) SetDepartureDate(v string) *EnrichRequestJourneyParamListSegmentParamList {
  s.DepartureDate = &v
  return s
}

func (s *EnrichRequestJourneyParamListSegmentParamList) SetDepartureTime(v string) *EnrichRequestJourneyParamListSegmentParamList {
  s.DepartureTime = &v
  return s
}

func (s *EnrichRequestJourneyParamListSegmentParamList) SetMarketingFlightNo(v string) *EnrichRequestJourneyParamListSegmentParamList {
  s.MarketingFlightNo = &v
  return s
}

func (s *EnrichRequestJourneyParamListSegmentParamList) Validate() error {
  return dara.Validate(s)
}

