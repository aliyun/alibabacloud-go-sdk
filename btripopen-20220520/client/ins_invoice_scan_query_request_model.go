// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsInvoiceScanQueryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillDate(v string) *InsInvoiceScanQueryRequest
	GetBillDate() *string
	SetBillId(v int64) *InsInvoiceScanQueryRequest
	GetBillId() *int64
	SetInvoiceSubTaskId(v int64) *InsInvoiceScanQueryRequest
	GetInvoiceSubTaskId() *int64
	SetPageNo(v int32) *InsInvoiceScanQueryRequest
	GetPageNo() *int32
	SetPageSize(v int32) *InsInvoiceScanQueryRequest
	GetPageSize() *int32
}

type InsInvoiceScanQueryRequest struct {
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
	PageSize *int32 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}

func (s InsInvoiceScanQueryRequest) String() string {
	return dara.Prettify(s)
}

func (s InsInvoiceScanQueryRequest) GoString() string {
	return s.String()
}

func (s *InsInvoiceScanQueryRequest) GetBillDate() *string {
	return s.BillDate
}

func (s *InsInvoiceScanQueryRequest) GetBillId() *int64 {
	return s.BillId
}

func (s *InsInvoiceScanQueryRequest) GetInvoiceSubTaskId() *int64 {
	return s.InvoiceSubTaskId
}

func (s *InsInvoiceScanQueryRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *InsInvoiceScanQueryRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *InsInvoiceScanQueryRequest) SetBillDate(v string) *InsInvoiceScanQueryRequest {
	s.BillDate = &v
	return s
}

func (s *InsInvoiceScanQueryRequest) SetBillId(v int64) *InsInvoiceScanQueryRequest {
	s.BillId = &v
	return s
}

func (s *InsInvoiceScanQueryRequest) SetInvoiceSubTaskId(v int64) *InsInvoiceScanQueryRequest {
	s.InvoiceSubTaskId = &v
	return s
}

func (s *InsInvoiceScanQueryRequest) SetPageNo(v int32) *InsInvoiceScanQueryRequest {
	s.PageNo = &v
	return s
}

func (s *InsInvoiceScanQueryRequest) SetPageSize(v int32) *InsInvoiceScanQueryRequest {
	s.PageSize = &v
	return s
}

func (s *InsInvoiceScanQueryRequest) Validate() error {
	return dara.Validate(s)
}
