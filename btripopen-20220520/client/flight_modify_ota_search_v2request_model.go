// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightModifyOtaSearchV2Request interface {
	dara.Model
	String() string
	GoString() string
	SetCabinClass(v []*int32) *FlightModifyOtaSearchV2Request
	GetCabinClass() []*int32
	SetDepDate(v []*string) *FlightModifyOtaSearchV2Request
	GetDepDate() []*string
	SetIsvName(v string) *FlightModifyOtaSearchV2Request
	GetIsvName() *string
	SetOrderId(v int64) *FlightModifyOtaSearchV2Request
	GetOrderId() *int64
	SetOutOrderId(v string) *FlightModifyOtaSearchV2Request
	GetOutOrderId() *string
	SetPassengerSegmentRelations(v []*FlightModifyOtaSearchV2RequestPassengerSegmentRelations) *FlightModifyOtaSearchV2Request
	GetPassengerSegmentRelations() []*FlightModifyOtaSearchV2RequestPassengerSegmentRelations
	SetSelectedSegments(v []*FlightModifyOtaSearchV2RequestSelectedSegments) *FlightModifyOtaSearchV2Request
	GetSelectedSegments() []*FlightModifyOtaSearchV2RequestSelectedSegments
	SetSessionId(v string) *FlightModifyOtaSearchV2Request
	GetSessionId() *string
	SetVoluntary(v bool) *FlightModifyOtaSearchV2Request
	GetVoluntary() *bool
}

type FlightModifyOtaSearchV2Request struct {
	CabinClass []*int32  `json:"cabin_class,omitempty" xml:"cabin_class,omitempty" type:"Repeated"`
	DepDate    []*string `json:"dep_date,omitempty" xml:"dep_date,omitempty" type:"Repeated"`
	// example:
	//
	// name
	IsvName *string `json:"isv_name,omitempty" xml:"isv_name,omitempty"`
	// example:
	//
	// 1017002195370467200
	OrderId *int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// example:
	//
	// 1017002195370467200
	OutOrderId                *string                                                    `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	PassengerSegmentRelations []*FlightModifyOtaSearchV2RequestPassengerSegmentRelations `json:"passenger_segment_relations,omitempty" xml:"passenger_segment_relations,omitempty" type:"Repeated"`
	SelectedSegments          []*FlightModifyOtaSearchV2RequestSelectedSegments          `json:"selected_segments,omitempty" xml:"selected_segments,omitempty" type:"Repeated"`
	// example:
	//
	// 590f17eca9374f20ac7e8ed8a7db2f35
	SessionId *string `json:"session_id,omitempty" xml:"session_id,omitempty"`
	// example:
	//
	// true
	Voluntary *bool `json:"voluntary,omitempty" xml:"voluntary,omitempty"`
}

func (s FlightModifyOtaSearchV2Request) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyOtaSearchV2Request) GoString() string {
	return s.String()
}

func (s *FlightModifyOtaSearchV2Request) GetCabinClass() []*int32 {
	return s.CabinClass
}

func (s *FlightModifyOtaSearchV2Request) GetDepDate() []*string {
	return s.DepDate
}

func (s *FlightModifyOtaSearchV2Request) GetIsvName() *string {
	return s.IsvName
}

func (s *FlightModifyOtaSearchV2Request) GetOrderId() *int64 {
	return s.OrderId
}

func (s *FlightModifyOtaSearchV2Request) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *FlightModifyOtaSearchV2Request) GetPassengerSegmentRelations() []*FlightModifyOtaSearchV2RequestPassengerSegmentRelations {
	return s.PassengerSegmentRelations
}

func (s *FlightModifyOtaSearchV2Request) GetSelectedSegments() []*FlightModifyOtaSearchV2RequestSelectedSegments {
	return s.SelectedSegments
}

func (s *FlightModifyOtaSearchV2Request) GetSessionId() *string {
	return s.SessionId
}

func (s *FlightModifyOtaSearchV2Request) GetVoluntary() *bool {
	return s.Voluntary
}

func (s *FlightModifyOtaSearchV2Request) SetCabinClass(v []*int32) *FlightModifyOtaSearchV2Request {
	s.CabinClass = v
	return s
}

func (s *FlightModifyOtaSearchV2Request) SetDepDate(v []*string) *FlightModifyOtaSearchV2Request {
	s.DepDate = v
	return s
}

func (s *FlightModifyOtaSearchV2Request) SetIsvName(v string) *FlightModifyOtaSearchV2Request {
	s.IsvName = &v
	return s
}

func (s *FlightModifyOtaSearchV2Request) SetOrderId(v int64) *FlightModifyOtaSearchV2Request {
	s.OrderId = &v
	return s
}

func (s *FlightModifyOtaSearchV2Request) SetOutOrderId(v string) *FlightModifyOtaSearchV2Request {
	s.OutOrderId = &v
	return s
}

func (s *FlightModifyOtaSearchV2Request) SetPassengerSegmentRelations(v []*FlightModifyOtaSearchV2RequestPassengerSegmentRelations) *FlightModifyOtaSearchV2Request {
	s.PassengerSegmentRelations = v
	return s
}

func (s *FlightModifyOtaSearchV2Request) SetSelectedSegments(v []*FlightModifyOtaSearchV2RequestSelectedSegments) *FlightModifyOtaSearchV2Request {
	s.SelectedSegments = v
	return s
}

func (s *FlightModifyOtaSearchV2Request) SetSessionId(v string) *FlightModifyOtaSearchV2Request {
	s.SessionId = &v
	return s
}

func (s *FlightModifyOtaSearchV2Request) SetVoluntary(v bool) *FlightModifyOtaSearchV2Request {
	s.Voluntary = &v
	return s
}

func (s *FlightModifyOtaSearchV2Request) Validate() error {
	return dara.Validate(s)
}

type FlightModifyOtaSearchV2RequestPassengerSegmentRelations struct {
	// example:
	//
	// 3243028
	PassengerId   *string   `json:"passenger_id,omitempty" xml:"passenger_id,omitempty"`
	SegmentIdList []*string `json:"segment_id_list,omitempty" xml:"segment_id_list,omitempty" type:"Repeated"`
}

func (s FlightModifyOtaSearchV2RequestPassengerSegmentRelations) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyOtaSearchV2RequestPassengerSegmentRelations) GoString() string {
	return s.String()
}

func (s *FlightModifyOtaSearchV2RequestPassengerSegmentRelations) GetPassengerId() *string {
	return s.PassengerId
}

func (s *FlightModifyOtaSearchV2RequestPassengerSegmentRelations) GetSegmentIdList() []*string {
	return s.SegmentIdList
}

func (s *FlightModifyOtaSearchV2RequestPassengerSegmentRelations) SetPassengerId(v string) *FlightModifyOtaSearchV2RequestPassengerSegmentRelations {
	s.PassengerId = &v
	return s
}

func (s *FlightModifyOtaSearchV2RequestPassengerSegmentRelations) SetSegmentIdList(v []*string) *FlightModifyOtaSearchV2RequestPassengerSegmentRelations {
	s.SegmentIdList = v
	return s
}

func (s *FlightModifyOtaSearchV2RequestPassengerSegmentRelations) Validate() error {
	return dara.Validate(s)
}

type FlightModifyOtaSearchV2RequestSelectedSegments struct {
	// example:
	//
	// BJS
	ArrCityCode *string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	// example:
	//
	// XIL
	DepCityCode *string `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	// example:
	//
	// 2023-09-18 09:10:00
	DepDateTime *string `json:"dep_date_time,omitempty" xml:"dep_date_time,omitempty"`
	// example:
	//
	// 0
	JourneySeq *int32 `json:"journey_seq,omitempty" xml:"journey_seq,omitempty"`
	// example:
	//
	// CA1110
	MarketingFlightNo *string `json:"marketing_flight_no,omitempty" xml:"marketing_flight_no,omitempty"`
	// example:
	//
	// MU8625
	OperatingFlightNo *string `json:"operating_flight_no,omitempty" xml:"operating_flight_no,omitempty"`
	// example:
	//
	// 0
	SegmentSeq *int32 `json:"segment_seq,omitempty" xml:"segment_seq,omitempty"`
}

func (s FlightModifyOtaSearchV2RequestSelectedSegments) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyOtaSearchV2RequestSelectedSegments) GoString() string {
	return s.String()
}

func (s *FlightModifyOtaSearchV2RequestSelectedSegments) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *FlightModifyOtaSearchV2RequestSelectedSegments) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *FlightModifyOtaSearchV2RequestSelectedSegments) GetDepDateTime() *string {
	return s.DepDateTime
}

func (s *FlightModifyOtaSearchV2RequestSelectedSegments) GetJourneySeq() *int32 {
	return s.JourneySeq
}

func (s *FlightModifyOtaSearchV2RequestSelectedSegments) GetMarketingFlightNo() *string {
	return s.MarketingFlightNo
}

func (s *FlightModifyOtaSearchV2RequestSelectedSegments) GetOperatingFlightNo() *string {
	return s.OperatingFlightNo
}

func (s *FlightModifyOtaSearchV2RequestSelectedSegments) GetSegmentSeq() *int32 {
	return s.SegmentSeq
}

func (s *FlightModifyOtaSearchV2RequestSelectedSegments) SetArrCityCode(v string) *FlightModifyOtaSearchV2RequestSelectedSegments {
	s.ArrCityCode = &v
	return s
}

func (s *FlightModifyOtaSearchV2RequestSelectedSegments) SetDepCityCode(v string) *FlightModifyOtaSearchV2RequestSelectedSegments {
	s.DepCityCode = &v
	return s
}

func (s *FlightModifyOtaSearchV2RequestSelectedSegments) SetDepDateTime(v string) *FlightModifyOtaSearchV2RequestSelectedSegments {
	s.DepDateTime = &v
	return s
}

func (s *FlightModifyOtaSearchV2RequestSelectedSegments) SetJourneySeq(v int32) *FlightModifyOtaSearchV2RequestSelectedSegments {
	s.JourneySeq = &v
	return s
}

func (s *FlightModifyOtaSearchV2RequestSelectedSegments) SetMarketingFlightNo(v string) *FlightModifyOtaSearchV2RequestSelectedSegments {
	s.MarketingFlightNo = &v
	return s
}

func (s *FlightModifyOtaSearchV2RequestSelectedSegments) SetOperatingFlightNo(v string) *FlightModifyOtaSearchV2RequestSelectedSegments {
	s.OperatingFlightNo = &v
	return s
}

func (s *FlightModifyOtaSearchV2RequestSelectedSegments) SetSegmentSeq(v int32) *FlightModifyOtaSearchV2RequestSelectedSegments {
	s.SegmentSeq = &v
	return s
}

func (s *FlightModifyOtaSearchV2RequestSelectedSegments) Validate() error {
	return dara.Validate(s)
}
