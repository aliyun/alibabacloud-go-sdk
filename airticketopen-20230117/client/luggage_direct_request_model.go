// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLuggageDirectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFlightSegmentParamList(v []*LuggageDirectRequestFlightSegmentParamList) *LuggageDirectRequest
	GetFlightSegmentParamList() []*LuggageDirectRequestFlightSegmentParamList
}

type LuggageDirectRequest struct {
	FlightSegmentParamList []*LuggageDirectRequestFlightSegmentParamList `json:"flight_segment_param_list,omitempty" xml:"flight_segment_param_list,omitempty" type:"Repeated"`
}

func (s LuggageDirectRequest) String() string {
	return dara.Prettify(s)
}

func (s LuggageDirectRequest) GoString() string {
	return s.String()
}

func (s *LuggageDirectRequest) GetFlightSegmentParamList() []*LuggageDirectRequestFlightSegmentParamList {
	return s.FlightSegmentParamList
}

func (s *LuggageDirectRequest) SetFlightSegmentParamList(v []*LuggageDirectRequestFlightSegmentParamList) *LuggageDirectRequest {
	s.FlightSegmentParamList = v
	return s
}

func (s *LuggageDirectRequest) Validate() error {
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

type LuggageDirectRequestFlightSegmentParamList struct {
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

func (s LuggageDirectRequestFlightSegmentParamList) String() string {
	return dara.Prettify(s)
}

func (s LuggageDirectRequestFlightSegmentParamList) GoString() string {
	return s.String()
}

func (s *LuggageDirectRequestFlightSegmentParamList) GetArrivalAirport() *string {
	return s.ArrivalAirport
}

func (s *LuggageDirectRequestFlightSegmentParamList) GetArrivalTerminal() *string {
	return s.ArrivalTerminal
}

func (s *LuggageDirectRequestFlightSegmentParamList) GetArrivalTime() *int64 {
	return s.ArrivalTime
}

func (s *LuggageDirectRequestFlightSegmentParamList) GetCodeShare() *bool {
	return s.CodeShare
}

func (s *LuggageDirectRequestFlightSegmentParamList) GetDepartureAirport() *string {
	return s.DepartureAirport
}

func (s *LuggageDirectRequestFlightSegmentParamList) GetDepartureTerminal() *string {
	return s.DepartureTerminal
}

func (s *LuggageDirectRequestFlightSegmentParamList) GetDepartureTime() *int64 {
	return s.DepartureTime
}

func (s *LuggageDirectRequestFlightSegmentParamList) GetMarketingAirline() *string {
	return s.MarketingAirline
}

func (s *LuggageDirectRequestFlightSegmentParamList) GetMarketingFlightNo() *string {
	return s.MarketingFlightNo
}

func (s *LuggageDirectRequestFlightSegmentParamList) GetOperatingAirline() *string {
	return s.OperatingAirline
}

func (s *LuggageDirectRequestFlightSegmentParamList) GetStopCityList() *string {
	return s.StopCityList
}

func (s *LuggageDirectRequestFlightSegmentParamList) GetTicketingAirline() *string {
	return s.TicketingAirline
}

func (s *LuggageDirectRequestFlightSegmentParamList) SetArrivalAirport(v string) *LuggageDirectRequestFlightSegmentParamList {
	s.ArrivalAirport = &v
	return s
}

func (s *LuggageDirectRequestFlightSegmentParamList) SetArrivalTerminal(v string) *LuggageDirectRequestFlightSegmentParamList {
	s.ArrivalTerminal = &v
	return s
}

func (s *LuggageDirectRequestFlightSegmentParamList) SetArrivalTime(v int64) *LuggageDirectRequestFlightSegmentParamList {
	s.ArrivalTime = &v
	return s
}

func (s *LuggageDirectRequestFlightSegmentParamList) SetCodeShare(v bool) *LuggageDirectRequestFlightSegmentParamList {
	s.CodeShare = &v
	return s
}

func (s *LuggageDirectRequestFlightSegmentParamList) SetDepartureAirport(v string) *LuggageDirectRequestFlightSegmentParamList {
	s.DepartureAirport = &v
	return s
}

func (s *LuggageDirectRequestFlightSegmentParamList) SetDepartureTerminal(v string) *LuggageDirectRequestFlightSegmentParamList {
	s.DepartureTerminal = &v
	return s
}

func (s *LuggageDirectRequestFlightSegmentParamList) SetDepartureTime(v int64) *LuggageDirectRequestFlightSegmentParamList {
	s.DepartureTime = &v
	return s
}

func (s *LuggageDirectRequestFlightSegmentParamList) SetMarketingAirline(v string) *LuggageDirectRequestFlightSegmentParamList {
	s.MarketingAirline = &v
	return s
}

func (s *LuggageDirectRequestFlightSegmentParamList) SetMarketingFlightNo(v string) *LuggageDirectRequestFlightSegmentParamList {
	s.MarketingFlightNo = &v
	return s
}

func (s *LuggageDirectRequestFlightSegmentParamList) SetOperatingAirline(v string) *LuggageDirectRequestFlightSegmentParamList {
	s.OperatingAirline = &v
	return s
}

func (s *LuggageDirectRequestFlightSegmentParamList) SetStopCityList(v string) *LuggageDirectRequestFlightSegmentParamList {
	s.StopCityList = &v
	return s
}

func (s *LuggageDirectRequestFlightSegmentParamList) SetTicketingAirline(v string) *LuggageDirectRequestFlightSegmentParamList {
	s.TicketingAirline = &v
	return s
}

func (s *LuggageDirectRequestFlightSegmentParamList) Validate() error {
	return dara.Validate(s)
}
