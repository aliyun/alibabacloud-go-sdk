// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightRefundPreCalV2Request interface {
	dara.Model
	String() string
	GoString() string
	SetIsvName(v string) *FlightRefundPreCalV2Request
	GetIsvName() *string
	SetOrderId(v string) *FlightRefundPreCalV2Request
	GetOrderId() *string
	SetOutOrderId(v string) *FlightRefundPreCalV2Request
	GetOutOrderId() *string
	SetPassengerSegmentRelations(v []*FlightRefundPreCalV2RequestPassengerSegmentRelations) *FlightRefundPreCalV2Request
	GetPassengerSegmentRelations() []*FlightRefundPreCalV2RequestPassengerSegmentRelations
	SetPreCalType(v int32) *FlightRefundPreCalV2Request
	GetPreCalType() *int32
	SetTicketNos(v []*string) *FlightRefundPreCalV2Request
	GetTicketNos() []*string
	SetVoluntary(v bool) *FlightRefundPreCalV2Request
	GetVoluntary() *bool
}

type FlightRefundPreCalV2Request struct {
	// example:
	//
	// cheshiapi
	IsvName *string `json:"isv_name,omitempty" xml:"isv_name,omitempty"`
	// example:
	//
	// 3454043907950204159
	OrderId *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// example:
	//
	// 1017002195370467137
	OutOrderId                *string                                                 `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	PassengerSegmentRelations []*FlightRefundPreCalV2RequestPassengerSegmentRelations `json:"passenger_segment_relations,omitempty" xml:"passenger_segment_relations,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	PreCalType *int32    `json:"pre_cal_type,omitempty" xml:"pre_cal_type,omitempty"`
	TicketNos  []*string `json:"ticket_nos,omitempty" xml:"ticket_nos,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Voluntary *bool `json:"voluntary,omitempty" xml:"voluntary,omitempty"`
}

func (s FlightRefundPreCalV2Request) String() string {
	return dara.Prettify(s)
}

func (s FlightRefundPreCalV2Request) GoString() string {
	return s.String()
}

func (s *FlightRefundPreCalV2Request) GetIsvName() *string {
	return s.IsvName
}

func (s *FlightRefundPreCalV2Request) GetOrderId() *string {
	return s.OrderId
}

func (s *FlightRefundPreCalV2Request) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *FlightRefundPreCalV2Request) GetPassengerSegmentRelations() []*FlightRefundPreCalV2RequestPassengerSegmentRelations {
	return s.PassengerSegmentRelations
}

func (s *FlightRefundPreCalV2Request) GetPreCalType() *int32 {
	return s.PreCalType
}

func (s *FlightRefundPreCalV2Request) GetTicketNos() []*string {
	return s.TicketNos
}

func (s *FlightRefundPreCalV2Request) GetVoluntary() *bool {
	return s.Voluntary
}

func (s *FlightRefundPreCalV2Request) SetIsvName(v string) *FlightRefundPreCalV2Request {
	s.IsvName = &v
	return s
}

func (s *FlightRefundPreCalV2Request) SetOrderId(v string) *FlightRefundPreCalV2Request {
	s.OrderId = &v
	return s
}

func (s *FlightRefundPreCalV2Request) SetOutOrderId(v string) *FlightRefundPreCalV2Request {
	s.OutOrderId = &v
	return s
}

func (s *FlightRefundPreCalV2Request) SetPassengerSegmentRelations(v []*FlightRefundPreCalV2RequestPassengerSegmentRelations) *FlightRefundPreCalV2Request {
	s.PassengerSegmentRelations = v
	return s
}

func (s *FlightRefundPreCalV2Request) SetPreCalType(v int32) *FlightRefundPreCalV2Request {
	s.PreCalType = &v
	return s
}

func (s *FlightRefundPreCalV2Request) SetTicketNos(v []*string) *FlightRefundPreCalV2Request {
	s.TicketNos = v
	return s
}

func (s *FlightRefundPreCalV2Request) SetVoluntary(v bool) *FlightRefundPreCalV2Request {
	s.Voluntary = &v
	return s
}

func (s *FlightRefundPreCalV2Request) Validate() error {
	return dara.Validate(s)
}

type FlightRefundPreCalV2RequestPassengerSegmentRelations struct {
	// example:
	//
	// 3243028
	PassengerId   *string   `json:"passenger_id,omitempty" xml:"passenger_id,omitempty"`
	SegmentIdList []*string `json:"segment_id_list,omitempty" xml:"segment_id_list,omitempty" type:"Repeated"`
}

func (s FlightRefundPreCalV2RequestPassengerSegmentRelations) String() string {
	return dara.Prettify(s)
}

func (s FlightRefundPreCalV2RequestPassengerSegmentRelations) GoString() string {
	return s.String()
}

func (s *FlightRefundPreCalV2RequestPassengerSegmentRelations) GetPassengerId() *string {
	return s.PassengerId
}

func (s *FlightRefundPreCalV2RequestPassengerSegmentRelations) GetSegmentIdList() []*string {
	return s.SegmentIdList
}

func (s *FlightRefundPreCalV2RequestPassengerSegmentRelations) SetPassengerId(v string) *FlightRefundPreCalV2RequestPassengerSegmentRelations {
	s.PassengerId = &v
	return s
}

func (s *FlightRefundPreCalV2RequestPassengerSegmentRelations) SetSegmentIdList(v []*string) *FlightRefundPreCalV2RequestPassengerSegmentRelations {
	s.SegmentIdList = v
	return s
}

func (s *FlightRefundPreCalV2RequestPassengerSegmentRelations) Validate() error {
	return dara.Validate(s)
}
