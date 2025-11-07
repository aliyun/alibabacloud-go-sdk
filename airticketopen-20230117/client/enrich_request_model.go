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
  // Number of adult passengers (1-9)
  // 
  // example:
  // 
  // 1
  Adults *int32 `json:"adults,omitempty" xml:"adults,omitempty"`
  // Cabin class: ALL_CABIN: All cabin classes; Y: Economy; FC: First Class and Business Class; S: Premium Economy; YS: Economy and Premium Economy; YSC: Economy, Premium Economy, and Business Class;
  // 
  // example:
  // 
  // ALL_CABIN
  CabinClass *string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
  // Number of child passengers (0-9)
  // 
  // example:
  // 
  // 1
  Children *int32 `json:"children,omitempty" xml:"children,omitempty"`
  // Number of infant passengers (0-9)
  // 
  // example:
  // 
  // 1
  Infants *int32 `json:"infants,omitempty" xml:"infants,omitempty"`
  // Trip information
  JourneyParamList []*EnrichRequestJourneyParamList `json:"journey_param_list,omitempty" xml:"journey_param_list,omitempty" type:"Repeated"`
  // The `solution_id` returned by the Search interface
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
  // Arrival city three-letter code (uppercase)
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // MFM
  ArrivalCity *string `json:"arrival_city,omitempty" xml:"arrival_city,omitempty"`
  // Departure city three-letter code (uppercase)
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // SHA
  DepartureCity *string `json:"departure_city,omitempty" xml:"departure_city,omitempty"`
  // Departure date (yyyyMMdd)
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 20230310
  DepartureDate *string `json:"departure_date,omitempty" xml:"departure_date,omitempty"`
  // Specified segment information for this trip
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
  // Flight arrival airport three-letter code (uppercase)
  // 
  // example:
  // 
  // MFM
  ArrivalAirport *string `json:"arrival_airport,omitempty" xml:"arrival_airport,omitempty"`
  // Flight arrival city three-letter code (uppercase)
  // 
  // example:
  // 
  // MFM
  ArrivalCity *string `json:"arrival_city,omitempty" xml:"arrival_city,omitempty"`
  // Cabin
  // 
  // example:
  // 
  // V
  Cabin *string `json:"cabin,omitempty" xml:"cabin,omitempty"`
  // Child cabin
  // 
  // example:
  // 
  // E
  ChildCabin *string `json:"child_cabin,omitempty" xml:"child_cabin,omitempty"`
  // Flight departure airport three-letter code (uppercase)
  // 
  // example:
  // 
  // PVG
  DepartureAirport *string `json:"departure_airport,omitempty" xml:"departure_airport,omitempty"`
  // Flight departure city three-letter code (uppercase)
  // 
  // example:
  // 
  // SHA
  DepartureCity *string `json:"departure_city,omitempty" xml:"departure_city,omitempty"`
  // String, flight departure date (yyyy-MM-dd), either departure_date or departure_time, with departure_time preferred for greater accuracy
  // 
  // example:
  // 
  // 2023-03-10
  DepartureDate *string `json:"departure_date,omitempty" xml:"departure_date,omitempty"`
  // String, flight departure date and time (yyyy-MM-dd HH:mm:ss)
  // 
  // example:
  // 
  // 2023-03-10 07:55:00
  DepartureTime *string `json:"departure_time,omitempty" xml:"departure_time,omitempty"`
  // Marketing flight number (e.g., KA5809)
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

