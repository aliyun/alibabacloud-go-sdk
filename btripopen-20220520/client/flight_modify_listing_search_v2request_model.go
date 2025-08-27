// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightModifyListingSearchV2Request interface {
	dara.Model
	String() string
	GoString() string
	SetCabinClass(v []*int32) *FlightModifyListingSearchV2Request
	GetCabinClass() []*int32
	SetDepDate(v []*string) *FlightModifyListingSearchV2Request
	GetDepDate() []*string
	SetInterfaceCallerIsSupportRetry(v bool) *FlightModifyListingSearchV2Request
	GetInterfaceCallerIsSupportRetry() *bool
	SetIsvName(v string) *FlightModifyListingSearchV2Request
	GetIsvName() *string
	SetOrderId(v int64) *FlightModifyListingSearchV2Request
	GetOrderId() *int64
	SetOutOrderId(v string) *FlightModifyListingSearchV2Request
	GetOutOrderId() *string
	SetPassengerSegmentRelations(v []*FlightModifyListingSearchV2RequestPassengerSegmentRelations) *FlightModifyListingSearchV2Request
	GetPassengerSegmentRelations() []*FlightModifyListingSearchV2RequestPassengerSegmentRelations
	SetSearchMode(v int32) *FlightModifyListingSearchV2Request
	GetSearchMode() *int32
	SetSearchRetryToken(v string) *FlightModifyListingSearchV2Request
	GetSearchRetryToken() *string
	SetSelectedSegments(v []*FlightModifyListingSearchV2RequestSelectedSegments) *FlightModifyListingSearchV2Request
	GetSelectedSegments() []*FlightModifyListingSearchV2RequestSelectedSegments
	SetSessionId(v string) *FlightModifyListingSearchV2Request
	GetSessionId() *string
	SetVoluntary(v bool) *FlightModifyListingSearchV2Request
	GetVoluntary() *bool
}

type FlightModifyListingSearchV2Request struct {
	CabinClass                    []*int32  `json:"cabin_class,omitempty" xml:"cabin_class,omitempty" type:"Repeated"`
	DepDate                       []*string `json:"dep_date,omitempty" xml:"dep_date,omitempty" type:"Repeated"`
	InterfaceCallerIsSupportRetry *bool     `json:"interface_caller_is_support_retry,omitempty" xml:"interface_caller_is_support_retry,omitempty"`
	// example:
	//
	// name
	IsvName *string `json:"isv_name,omitempty" xml:"isv_name,omitempty"`
	// example:
	//
	// 1017002195370467138
	OrderId *int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// example:
	//
	// 1017002195370467137
	OutOrderId                *string                                                        `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	PassengerSegmentRelations []*FlightModifyListingSearchV2RequestPassengerSegmentRelations `json:"passenger_segment_relations,omitempty" xml:"passenger_segment_relations,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	SearchMode       *int32                                                `json:"search_mode,omitempty" xml:"search_mode,omitempty"`
	SearchRetryToken *string                                               `json:"search_retry_token,omitempty" xml:"search_retry_token,omitempty"`
	SelectedSegments []*FlightModifyListingSearchV2RequestSelectedSegments `json:"selected_segments,omitempty" xml:"selected_segments,omitempty" type:"Repeated"`
	// example:
	//
	// a2ffebfe733742aab5c491d960ba3d59
	SessionId *string `json:"session_id,omitempty" xml:"session_id,omitempty"`
	// example:
	//
	// true
	Voluntary *bool `json:"voluntary,omitempty" xml:"voluntary,omitempty"`
}

func (s FlightModifyListingSearchV2Request) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyListingSearchV2Request) GoString() string {
	return s.String()
}

func (s *FlightModifyListingSearchV2Request) GetCabinClass() []*int32 {
	return s.CabinClass
}

func (s *FlightModifyListingSearchV2Request) GetDepDate() []*string {
	return s.DepDate
}

func (s *FlightModifyListingSearchV2Request) GetInterfaceCallerIsSupportRetry() *bool {
	return s.InterfaceCallerIsSupportRetry
}

func (s *FlightModifyListingSearchV2Request) GetIsvName() *string {
	return s.IsvName
}

func (s *FlightModifyListingSearchV2Request) GetOrderId() *int64 {
	return s.OrderId
}

func (s *FlightModifyListingSearchV2Request) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *FlightModifyListingSearchV2Request) GetPassengerSegmentRelations() []*FlightModifyListingSearchV2RequestPassengerSegmentRelations {
	return s.PassengerSegmentRelations
}

func (s *FlightModifyListingSearchV2Request) GetSearchMode() *int32 {
	return s.SearchMode
}

func (s *FlightModifyListingSearchV2Request) GetSearchRetryToken() *string {
	return s.SearchRetryToken
}

func (s *FlightModifyListingSearchV2Request) GetSelectedSegments() []*FlightModifyListingSearchV2RequestSelectedSegments {
	return s.SelectedSegments
}

func (s *FlightModifyListingSearchV2Request) GetSessionId() *string {
	return s.SessionId
}

func (s *FlightModifyListingSearchV2Request) GetVoluntary() *bool {
	return s.Voluntary
}

func (s *FlightModifyListingSearchV2Request) SetCabinClass(v []*int32) *FlightModifyListingSearchV2Request {
	s.CabinClass = v
	return s
}

func (s *FlightModifyListingSearchV2Request) SetDepDate(v []*string) *FlightModifyListingSearchV2Request {
	s.DepDate = v
	return s
}

func (s *FlightModifyListingSearchV2Request) SetInterfaceCallerIsSupportRetry(v bool) *FlightModifyListingSearchV2Request {
	s.InterfaceCallerIsSupportRetry = &v
	return s
}

func (s *FlightModifyListingSearchV2Request) SetIsvName(v string) *FlightModifyListingSearchV2Request {
	s.IsvName = &v
	return s
}

func (s *FlightModifyListingSearchV2Request) SetOrderId(v int64) *FlightModifyListingSearchV2Request {
	s.OrderId = &v
	return s
}

func (s *FlightModifyListingSearchV2Request) SetOutOrderId(v string) *FlightModifyListingSearchV2Request {
	s.OutOrderId = &v
	return s
}

func (s *FlightModifyListingSearchV2Request) SetPassengerSegmentRelations(v []*FlightModifyListingSearchV2RequestPassengerSegmentRelations) *FlightModifyListingSearchV2Request {
	s.PassengerSegmentRelations = v
	return s
}

func (s *FlightModifyListingSearchV2Request) SetSearchMode(v int32) *FlightModifyListingSearchV2Request {
	s.SearchMode = &v
	return s
}

func (s *FlightModifyListingSearchV2Request) SetSearchRetryToken(v string) *FlightModifyListingSearchV2Request {
	s.SearchRetryToken = &v
	return s
}

func (s *FlightModifyListingSearchV2Request) SetSelectedSegments(v []*FlightModifyListingSearchV2RequestSelectedSegments) *FlightModifyListingSearchV2Request {
	s.SelectedSegments = v
	return s
}

func (s *FlightModifyListingSearchV2Request) SetSessionId(v string) *FlightModifyListingSearchV2Request {
	s.SessionId = &v
	return s
}

func (s *FlightModifyListingSearchV2Request) SetVoluntary(v bool) *FlightModifyListingSearchV2Request {
	s.Voluntary = &v
	return s
}

func (s *FlightModifyListingSearchV2Request) Validate() error {
	return dara.Validate(s)
}

type FlightModifyListingSearchV2RequestPassengerSegmentRelations struct {
	// example:
	//
	// 3243028
	PassengerId   *string   `json:"passenger_id,omitempty" xml:"passenger_id,omitempty"`
	SegmentIdList []*string `json:"segment_id_list,omitempty" xml:"segment_id_list,omitempty" type:"Repeated"`
}

func (s FlightModifyListingSearchV2RequestPassengerSegmentRelations) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyListingSearchV2RequestPassengerSegmentRelations) GoString() string {
	return s.String()
}

func (s *FlightModifyListingSearchV2RequestPassengerSegmentRelations) GetPassengerId() *string {
	return s.PassengerId
}

func (s *FlightModifyListingSearchV2RequestPassengerSegmentRelations) GetSegmentIdList() []*string {
	return s.SegmentIdList
}

func (s *FlightModifyListingSearchV2RequestPassengerSegmentRelations) SetPassengerId(v string) *FlightModifyListingSearchV2RequestPassengerSegmentRelations {
	s.PassengerId = &v
	return s
}

func (s *FlightModifyListingSearchV2RequestPassengerSegmentRelations) SetSegmentIdList(v []*string) *FlightModifyListingSearchV2RequestPassengerSegmentRelations {
	s.SegmentIdList = v
	return s
}

func (s *FlightModifyListingSearchV2RequestPassengerSegmentRelations) Validate() error {
	return dara.Validate(s)
}

type FlightModifyListingSearchV2RequestSelectedSegments struct {
	// example:
	//
	// XIL
	ArrCityCode *string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	// example:
	//
	// BJS
	DepCityCode *string `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	// example:
	//
	// 2023-09-17 18:15:00
	DepDateTime *string `json:"dep_date_time,omitempty" xml:"dep_date_time,omitempty"`
	// example:
	//
	// 0
	JourneySeq *int32 `json:"journey_seq,omitempty" xml:"journey_seq,omitempty"`
	// example:
	//
	// CA8625
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

func (s FlightModifyListingSearchV2RequestSelectedSegments) String() string {
	return dara.Prettify(s)
}

func (s FlightModifyListingSearchV2RequestSelectedSegments) GoString() string {
	return s.String()
}

func (s *FlightModifyListingSearchV2RequestSelectedSegments) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *FlightModifyListingSearchV2RequestSelectedSegments) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *FlightModifyListingSearchV2RequestSelectedSegments) GetDepDateTime() *string {
	return s.DepDateTime
}

func (s *FlightModifyListingSearchV2RequestSelectedSegments) GetJourneySeq() *int32 {
	return s.JourneySeq
}

func (s *FlightModifyListingSearchV2RequestSelectedSegments) GetMarketingFlightNo() *string {
	return s.MarketingFlightNo
}

func (s *FlightModifyListingSearchV2RequestSelectedSegments) GetOperatingFlightNo() *string {
	return s.OperatingFlightNo
}

func (s *FlightModifyListingSearchV2RequestSelectedSegments) GetSegmentSeq() *int32 {
	return s.SegmentSeq
}

func (s *FlightModifyListingSearchV2RequestSelectedSegments) SetArrCityCode(v string) *FlightModifyListingSearchV2RequestSelectedSegments {
	s.ArrCityCode = &v
	return s
}

func (s *FlightModifyListingSearchV2RequestSelectedSegments) SetDepCityCode(v string) *FlightModifyListingSearchV2RequestSelectedSegments {
	s.DepCityCode = &v
	return s
}

func (s *FlightModifyListingSearchV2RequestSelectedSegments) SetDepDateTime(v string) *FlightModifyListingSearchV2RequestSelectedSegments {
	s.DepDateTime = &v
	return s
}

func (s *FlightModifyListingSearchV2RequestSelectedSegments) SetJourneySeq(v int32) *FlightModifyListingSearchV2RequestSelectedSegments {
	s.JourneySeq = &v
	return s
}

func (s *FlightModifyListingSearchV2RequestSelectedSegments) SetMarketingFlightNo(v string) *FlightModifyListingSearchV2RequestSelectedSegments {
	s.MarketingFlightNo = &v
	return s
}

func (s *FlightModifyListingSearchV2RequestSelectedSegments) SetOperatingFlightNo(v string) *FlightModifyListingSearchV2RequestSelectedSegments {
	s.OperatingFlightNo = &v
	return s
}

func (s *FlightModifyListingSearchV2RequestSelectedSegments) SetSegmentSeq(v int32) *FlightModifyListingSearchV2RequestSelectedSegments {
	s.SegmentSeq = &v
	return s
}

func (s *FlightModifyListingSearchV2RequestSelectedSegments) Validate() error {
	return dara.Validate(s)
}
