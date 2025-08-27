// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightItineraryScanQueryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillDate(v string) *FlightItineraryScanQueryRequest
	GetBillDate() *string
	SetBillId(v int64) *FlightItineraryScanQueryRequest
	GetBillId() *int64
	SetInvoiceSubTaskId(v int64) *FlightItineraryScanQueryRequest
	GetInvoiceSubTaskId() *int64
	SetItineraryNum(v string) *FlightItineraryScanQueryRequest
	GetItineraryNum() *string
	SetPageNo(v int32) *FlightItineraryScanQueryRequest
	GetPageNo() *int32
	SetPageSize(v int32) *FlightItineraryScanQueryRequest
	GetPageSize() *int32
	SetTicketNo(v string) *FlightItineraryScanQueryRequest
	GetTicketNo() *string
}

type FlightItineraryScanQueryRequest struct {
	// example:
	//
	// 2022-12-01
	BillDate *string `json:"bill_date,omitempty" xml:"bill_date,omitempty"`
	// example:
	//
	// 123
	BillId *int64 `json:"bill_id,omitempty" xml:"bill_id,omitempty"`
	// example:
	//
	// 456
	InvoiceSubTaskId *int64  `json:"invoice_sub_task_id,omitempty" xml:"invoice_sub_task_id,omitempty"`
	ItineraryNum     *string `json:"itinerary_num,omitempty" xml:"itinerary_num,omitempty"`
	// example:
	//
	// 1
	PageNo *int32 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// example:
	//
	// 20
	PageSize *int32  `json:"page_size,omitempty" xml:"page_size,omitempty"`
	TicketNo *string `json:"ticket_no,omitempty" xml:"ticket_no,omitempty"`
}

func (s FlightItineraryScanQueryRequest) String() string {
	return dara.Prettify(s)
}

func (s FlightItineraryScanQueryRequest) GoString() string {
	return s.String()
}

func (s *FlightItineraryScanQueryRequest) GetBillDate() *string {
	return s.BillDate
}

func (s *FlightItineraryScanQueryRequest) GetBillId() *int64 {
	return s.BillId
}

func (s *FlightItineraryScanQueryRequest) GetInvoiceSubTaskId() *int64 {
	return s.InvoiceSubTaskId
}

func (s *FlightItineraryScanQueryRequest) GetItineraryNum() *string {
	return s.ItineraryNum
}

func (s *FlightItineraryScanQueryRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *FlightItineraryScanQueryRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *FlightItineraryScanQueryRequest) GetTicketNo() *string {
	return s.TicketNo
}

func (s *FlightItineraryScanQueryRequest) SetBillDate(v string) *FlightItineraryScanQueryRequest {
	s.BillDate = &v
	return s
}

func (s *FlightItineraryScanQueryRequest) SetBillId(v int64) *FlightItineraryScanQueryRequest {
	s.BillId = &v
	return s
}

func (s *FlightItineraryScanQueryRequest) SetInvoiceSubTaskId(v int64) *FlightItineraryScanQueryRequest {
	s.InvoiceSubTaskId = &v
	return s
}

func (s *FlightItineraryScanQueryRequest) SetItineraryNum(v string) *FlightItineraryScanQueryRequest {
	s.ItineraryNum = &v
	return s
}

func (s *FlightItineraryScanQueryRequest) SetPageNo(v int32) *FlightItineraryScanQueryRequest {
	s.PageNo = &v
	return s
}

func (s *FlightItineraryScanQueryRequest) SetPageSize(v int32) *FlightItineraryScanQueryRequest {
	s.PageSize = &v
	return s
}

func (s *FlightItineraryScanQueryRequest) SetTicketNo(v string) *FlightItineraryScanQueryRequest {
	s.TicketNo = &v
	return s
}

func (s *FlightItineraryScanQueryRequest) Validate() error {
	return dara.Validate(s)
}
