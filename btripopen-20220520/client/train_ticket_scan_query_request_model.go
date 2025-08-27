// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainTicketScanQueryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillDate(v string) *TrainTicketScanQueryRequest
	GetBillDate() *string
	SetBillId(v int64) *TrainTicketScanQueryRequest
	GetBillId() *int64
	SetInvoiceSubTaskId(v int64) *TrainTicketScanQueryRequest
	GetInvoiceSubTaskId() *int64
	SetPageNo(v int32) *TrainTicketScanQueryRequest
	GetPageNo() *int32
	SetPageSize(v int32) *TrainTicketScanQueryRequest
	GetPageSize() *int32
	SetSerialNumber(v string) *TrainTicketScanQueryRequest
	GetSerialNumber() *string
	SetTicketNo(v string) *TrainTicketScanQueryRequest
	GetTicketNo() *string
}

type TrainTicketScanQueryRequest struct {
	// This parameter is required.
	//
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
	InvoiceSubTaskId *int64 `json:"invoice_sub_task_id,omitempty" xml:"invoice_sub_task_id,omitempty"`
	// example:
	//
	// 1
	PageNo *int32 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// example:
	//
	// 20
	PageSize     *int32  `json:"page_size,omitempty" xml:"page_size,omitempty"`
	SerialNumber *string `json:"serial_number,omitempty" xml:"serial_number,omitempty"`
	TicketNo     *string `json:"ticket_no,omitempty" xml:"ticket_no,omitempty"`
}

func (s TrainTicketScanQueryRequest) String() string {
	return dara.Prettify(s)
}

func (s TrainTicketScanQueryRequest) GoString() string {
	return s.String()
}

func (s *TrainTicketScanQueryRequest) GetBillDate() *string {
	return s.BillDate
}

func (s *TrainTicketScanQueryRequest) GetBillId() *int64 {
	return s.BillId
}

func (s *TrainTicketScanQueryRequest) GetInvoiceSubTaskId() *int64 {
	return s.InvoiceSubTaskId
}

func (s *TrainTicketScanQueryRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *TrainTicketScanQueryRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *TrainTicketScanQueryRequest) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *TrainTicketScanQueryRequest) GetTicketNo() *string {
	return s.TicketNo
}

func (s *TrainTicketScanQueryRequest) SetBillDate(v string) *TrainTicketScanQueryRequest {
	s.BillDate = &v
	return s
}

func (s *TrainTicketScanQueryRequest) SetBillId(v int64) *TrainTicketScanQueryRequest {
	s.BillId = &v
	return s
}

func (s *TrainTicketScanQueryRequest) SetInvoiceSubTaskId(v int64) *TrainTicketScanQueryRequest {
	s.InvoiceSubTaskId = &v
	return s
}

func (s *TrainTicketScanQueryRequest) SetPageNo(v int32) *TrainTicketScanQueryRequest {
	s.PageNo = &v
	return s
}

func (s *TrainTicketScanQueryRequest) SetPageSize(v int32) *TrainTicketScanQueryRequest {
	s.PageSize = &v
	return s
}

func (s *TrainTicketScanQueryRequest) SetSerialNumber(v string) *TrainTicketScanQueryRequest {
	s.SerialNumber = &v
	return s
}

func (s *TrainTicketScanQueryRequest) SetTicketNo(v string) *TrainTicketScanQueryRequest {
	s.TicketNo = &v
	return s
}

func (s *TrainTicketScanQueryRequest) Validate() error {
	return dara.Validate(s)
}
