// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightSearchListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAirlineCode(v string) *FlightSearchListRequest
	GetAirlineCode() *string
	SetArrCityCode(v string) *FlightSearchListRequest
	GetArrCityCode() *string
	SetArrCityName(v string) *FlightSearchListRequest
	GetArrCityName() *string
	SetArrDate(v string) *FlightSearchListRequest
	GetArrDate() *string
	SetCabinClass(v string) *FlightSearchListRequest
	GetCabinClass() *string
	SetDepCityCode(v string) *FlightSearchListRequest
	GetDepCityCode() *string
	SetDepCityName(v string) *FlightSearchListRequest
	GetDepCityName() *string
	SetDepDate(v string) *FlightSearchListRequest
	GetDepDate() *string
	SetFlightNo(v string) *FlightSearchListRequest
	GetFlightNo() *string
	SetNeedMultiClassPrice(v bool) *FlightSearchListRequest
	GetNeedMultiClassPrice() *bool
	SetTransferCityCode(v string) *FlightSearchListRequest
	GetTransferCityCode() *string
	SetTransferFlightNo(v string) *FlightSearchListRequest
	GetTransferFlightNo() *string
	SetTransferLeaveDate(v string) *FlightSearchListRequest
	GetTransferLeaveDate() *string
	SetTripType(v string) *FlightSearchListRequest
	GetTripType() *string
}

type FlightSearchListRequest struct {
	// example:
	//
	// CA
	AirlineCode *string `json:"airline_code,omitempty" xml:"airline_code,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// HGH
	ArrCityCode *string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	// example:
	//
	// BJS
	ArrCityName *string `json:"arr_city_name,omitempty" xml:"arr_city_name,omitempty"`
	// example:
	//
	// 2000-00-00 00:00:00
	ArrDate *string `json:"arr_date,omitempty" xml:"arr_date,omitempty"`
	// example:
	//
	// Y
	CabinClass *string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// HGH
	DepCityCode *string `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	// example:
	//
	// HGH
	DepCityName *string `json:"dep_city_name,omitempty" xml:"dep_city_name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2000-00-00 00:00:00
	DepDate *string `json:"dep_date,omitempty" xml:"dep_date,omitempty"`
	// example:
	//
	// CA2323
	FlightNo *string `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
	// example:
	//
	// false
	NeedMultiClassPrice *bool `json:"need_multi_class_price,omitempty" xml:"need_multi_class_price,omitempty"`
	// example:
	//
	// HGH
	TransferCityCode *string `json:"transfer_city_code,omitempty" xml:"transfer_city_code,omitempty"`
	// example:
	//
	// CA2323
	TransferFlightNo *string `json:"transfer_flight_no,omitempty" xml:"transfer_flight_no,omitempty"`
	// example:
	//
	// BJS
	TransferLeaveDate *string `json:"transfer_leave_date,omitempty" xml:"transfer_leave_date,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	TripType *string `json:"trip_type,omitempty" xml:"trip_type,omitempty"`
}

func (s FlightSearchListRequest) String() string {
	return dara.Prettify(s)
}

func (s FlightSearchListRequest) GoString() string {
	return s.String()
}

func (s *FlightSearchListRequest) GetAirlineCode() *string {
	return s.AirlineCode
}

func (s *FlightSearchListRequest) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *FlightSearchListRequest) GetArrCityName() *string {
	return s.ArrCityName
}

func (s *FlightSearchListRequest) GetArrDate() *string {
	return s.ArrDate
}

func (s *FlightSearchListRequest) GetCabinClass() *string {
	return s.CabinClass
}

func (s *FlightSearchListRequest) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *FlightSearchListRequest) GetDepCityName() *string {
	return s.DepCityName
}

func (s *FlightSearchListRequest) GetDepDate() *string {
	return s.DepDate
}

func (s *FlightSearchListRequest) GetFlightNo() *string {
	return s.FlightNo
}

func (s *FlightSearchListRequest) GetNeedMultiClassPrice() *bool {
	return s.NeedMultiClassPrice
}

func (s *FlightSearchListRequest) GetTransferCityCode() *string {
	return s.TransferCityCode
}

func (s *FlightSearchListRequest) GetTransferFlightNo() *string {
	return s.TransferFlightNo
}

func (s *FlightSearchListRequest) GetTransferLeaveDate() *string {
	return s.TransferLeaveDate
}

func (s *FlightSearchListRequest) GetTripType() *string {
	return s.TripType
}

func (s *FlightSearchListRequest) SetAirlineCode(v string) *FlightSearchListRequest {
	s.AirlineCode = &v
	return s
}

func (s *FlightSearchListRequest) SetArrCityCode(v string) *FlightSearchListRequest {
	s.ArrCityCode = &v
	return s
}

func (s *FlightSearchListRequest) SetArrCityName(v string) *FlightSearchListRequest {
	s.ArrCityName = &v
	return s
}

func (s *FlightSearchListRequest) SetArrDate(v string) *FlightSearchListRequest {
	s.ArrDate = &v
	return s
}

func (s *FlightSearchListRequest) SetCabinClass(v string) *FlightSearchListRequest {
	s.CabinClass = &v
	return s
}

func (s *FlightSearchListRequest) SetDepCityCode(v string) *FlightSearchListRequest {
	s.DepCityCode = &v
	return s
}

func (s *FlightSearchListRequest) SetDepCityName(v string) *FlightSearchListRequest {
	s.DepCityName = &v
	return s
}

func (s *FlightSearchListRequest) SetDepDate(v string) *FlightSearchListRequest {
	s.DepDate = &v
	return s
}

func (s *FlightSearchListRequest) SetFlightNo(v string) *FlightSearchListRequest {
	s.FlightNo = &v
	return s
}

func (s *FlightSearchListRequest) SetNeedMultiClassPrice(v bool) *FlightSearchListRequest {
	s.NeedMultiClassPrice = &v
	return s
}

func (s *FlightSearchListRequest) SetTransferCityCode(v string) *FlightSearchListRequest {
	s.TransferCityCode = &v
	return s
}

func (s *FlightSearchListRequest) SetTransferFlightNo(v string) *FlightSearchListRequest {
	s.TransferFlightNo = &v
	return s
}

func (s *FlightSearchListRequest) SetTransferLeaveDate(v string) *FlightSearchListRequest {
	s.TransferLeaveDate = &v
	return s
}

func (s *FlightSearchListRequest) SetTripType(v string) *FlightSearchListRequest {
	s.TripType = &v
	return s
}

func (s *FlightSearchListRequest) Validate() error {
	return dara.Validate(s)
}
