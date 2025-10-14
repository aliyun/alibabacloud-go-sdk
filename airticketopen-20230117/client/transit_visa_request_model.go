// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransitVisaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFlightSegmentParamList(v []*TransitVisaRequestFlightSegmentParamList) *TransitVisaRequest
	GetFlightSegmentParamList() []*TransitVisaRequestFlightSegmentParamList
}

type TransitVisaRequest struct {
	FlightSegmentParamList []*TransitVisaRequestFlightSegmentParamList `json:"flight_segment_param_list,omitempty" xml:"flight_segment_param_list,omitempty" type:"Repeated"`
}

func (s TransitVisaRequest) String() string {
	return dara.Prettify(s)
}

func (s TransitVisaRequest) GoString() string {
	return s.String()
}

func (s *TransitVisaRequest) GetFlightSegmentParamList() []*TransitVisaRequestFlightSegmentParamList {
	return s.FlightSegmentParamList
}

func (s *TransitVisaRequest) SetFlightSegmentParamList(v []*TransitVisaRequestFlightSegmentParamList) *TransitVisaRequest {
	s.FlightSegmentParamList = v
	return s
}

func (s *TransitVisaRequest) Validate() error {
	if s.FlightSegmentParamList != nil {
		for _, item := range s.FlightSegmentParamList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type TransitVisaRequestFlightSegmentParamList struct {
	// This parameter is required.
	//
	// example:
	//
	// SIN
	ArrivalAirport *string `json:"arrival_airport,omitempty" xml:"arrival_airport,omitempty"`
	// example:
	//
	// T1
	ArrivalTerminal *string `json:"arrival_terminal,omitempty" xml:"arrival_terminal,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1705285430445
	ArrivalTime *int64 `json:"arrival_time,omitempty" xml:"arrival_time,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	CodeShare *bool `json:"code_share,omitempty" xml:"code_share,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// HGH
	DepartureAirport *string `json:"departure_airport,omitempty" xml:"departure_airport,omitempty"`
	// example:
	//
	// T1
	DepartureTerminal *string `json:"departure_terminal,omitempty" xml:"departure_terminal,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1705285430445
	DepartureTime *int64 `json:"departure_time,omitempty" xml:"departure_time,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CZ
	MarketingAirline *string `json:"marketing_airline,omitempty" xml:"marketing_airline,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CZ616
	MarketingFlightNo *string `json:"marketing_flight_no,omitempty" xml:"marketing_flight_no,omitempty"`
	// example:
	//
	// CZ
	OperatingAirline *string `json:"operating_airline,omitempty" xml:"operating_airline,omitempty"`
	// example:
	//
	// SEL,HKG
	StopCityList *string `json:"stop_city_list,omitempty" xml:"stop_city_list,omitempty"`
	// example:
	//
	// CZ
	TicketingAirline *string `json:"ticketing_airline,omitempty" xml:"ticketing_airline,omitempty"`
}

func (s TransitVisaRequestFlightSegmentParamList) String() string {
	return dara.Prettify(s)
}

func (s TransitVisaRequestFlightSegmentParamList) GoString() string {
	return s.String()
}

func (s *TransitVisaRequestFlightSegmentParamList) GetArrivalAirport() *string {
	return s.ArrivalAirport
}

func (s *TransitVisaRequestFlightSegmentParamList) GetArrivalTerminal() *string {
	return s.ArrivalTerminal
}

func (s *TransitVisaRequestFlightSegmentParamList) GetArrivalTime() *int64 {
	return s.ArrivalTime
}

func (s *TransitVisaRequestFlightSegmentParamList) GetCodeShare() *bool {
	return s.CodeShare
}

func (s *TransitVisaRequestFlightSegmentParamList) GetDepartureAirport() *string {
	return s.DepartureAirport
}

func (s *TransitVisaRequestFlightSegmentParamList) GetDepartureTerminal() *string {
	return s.DepartureTerminal
}

func (s *TransitVisaRequestFlightSegmentParamList) GetDepartureTime() *int64 {
	return s.DepartureTime
}

func (s *TransitVisaRequestFlightSegmentParamList) GetMarketingAirline() *string {
	return s.MarketingAirline
}

func (s *TransitVisaRequestFlightSegmentParamList) GetMarketingFlightNo() *string {
	return s.MarketingFlightNo
}

func (s *TransitVisaRequestFlightSegmentParamList) GetOperatingAirline() *string {
	return s.OperatingAirline
}

func (s *TransitVisaRequestFlightSegmentParamList) GetStopCityList() *string {
	return s.StopCityList
}

func (s *TransitVisaRequestFlightSegmentParamList) GetTicketingAirline() *string {
	return s.TicketingAirline
}

func (s *TransitVisaRequestFlightSegmentParamList) SetArrivalAirport(v string) *TransitVisaRequestFlightSegmentParamList {
	s.ArrivalAirport = &v
	return s
}

func (s *TransitVisaRequestFlightSegmentParamList) SetArrivalTerminal(v string) *TransitVisaRequestFlightSegmentParamList {
	s.ArrivalTerminal = &v
	return s
}

func (s *TransitVisaRequestFlightSegmentParamList) SetArrivalTime(v int64) *TransitVisaRequestFlightSegmentParamList {
	s.ArrivalTime = &v
	return s
}

func (s *TransitVisaRequestFlightSegmentParamList) SetCodeShare(v bool) *TransitVisaRequestFlightSegmentParamList {
	s.CodeShare = &v
	return s
}

func (s *TransitVisaRequestFlightSegmentParamList) SetDepartureAirport(v string) *TransitVisaRequestFlightSegmentParamList {
	s.DepartureAirport = &v
	return s
}

func (s *TransitVisaRequestFlightSegmentParamList) SetDepartureTerminal(v string) *TransitVisaRequestFlightSegmentParamList {
	s.DepartureTerminal = &v
	return s
}

func (s *TransitVisaRequestFlightSegmentParamList) SetDepartureTime(v int64) *TransitVisaRequestFlightSegmentParamList {
	s.DepartureTime = &v
	return s
}

func (s *TransitVisaRequestFlightSegmentParamList) SetMarketingAirline(v string) *TransitVisaRequestFlightSegmentParamList {
	s.MarketingAirline = &v
	return s
}

func (s *TransitVisaRequestFlightSegmentParamList) SetMarketingFlightNo(v string) *TransitVisaRequestFlightSegmentParamList {
	s.MarketingFlightNo = &v
	return s
}

func (s *TransitVisaRequestFlightSegmentParamList) SetOperatingAirline(v string) *TransitVisaRequestFlightSegmentParamList {
	s.OperatingAirline = &v
	return s
}

func (s *TransitVisaRequestFlightSegmentParamList) SetStopCityList(v string) *TransitVisaRequestFlightSegmentParamList {
	s.StopCityList = &v
	return s
}

func (s *TransitVisaRequestFlightSegmentParamList) SetTicketingAirline(v string) *TransitVisaRequestFlightSegmentParamList {
	s.TicketingAirline = &v
	return s
}

func (s *TransitVisaRequestFlightSegmentParamList) Validate() error {
	return dara.Validate(s)
}
