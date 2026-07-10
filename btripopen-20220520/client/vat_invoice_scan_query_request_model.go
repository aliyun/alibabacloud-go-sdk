// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVatInvoiceScanQueryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillDate(v string) *VatInvoiceScanQueryRequest
	GetBillDate() *string
	SetBillId(v int64) *VatInvoiceScanQueryRequest
	GetBillId() *int64
	SetInvoiceSubTaskId(v int64) *VatInvoiceScanQueryRequest
	GetInvoiceSubTaskId() *int64
	SetPageNo(v int32) *VatInvoiceScanQueryRequest
	GetPageNo() *int32
	SetPageSize(v int32) *VatInvoiceScanQueryRequest
	GetPageSize() *int32
}

type VatInvoiceScanQueryRequest struct {
	// This parameter is required.
	BillDate         *string `json:"bill_date,omitempty" xml:"bill_date,omitempty"`
	BillId           *int64  `json:"bill_id,omitempty" xml:"bill_id,omitempty"`
	InvoiceSubTaskId *int64  `json:"invoice_sub_task_id,omitempty" xml:"invoice_sub_task_id,omitempty"`
	PageNo           *int32  `json:"page_no,omitempty" xml:"page_no,omitempty"`
	PageSize         *int32  `json:"page_size,omitempty" xml:"page_size,omitempty"`
}

func (s VatInvoiceScanQueryRequest) String() string {
	return dara.Prettify(s)
}

func (s VatInvoiceScanQueryRequest) GoString() string {
	return s.String()
}

func (s *VatInvoiceScanQueryRequest) GetBillDate() *string {
	return s.BillDate
}

func (s *VatInvoiceScanQueryRequest) GetBillId() *int64 {
	return s.BillId
}

func (s *VatInvoiceScanQueryRequest) GetInvoiceSubTaskId() *int64 {
	return s.InvoiceSubTaskId
}

func (s *VatInvoiceScanQueryRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *VatInvoiceScanQueryRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *VatInvoiceScanQueryRequest) SetBillDate(v string) *VatInvoiceScanQueryRequest {
	s.BillDate = &v
	return s
}

func (s *VatInvoiceScanQueryRequest) SetBillId(v int64) *VatInvoiceScanQueryRequest {
	s.BillId = &v
	return s
}

func (s *VatInvoiceScanQueryRequest) SetInvoiceSubTaskId(v int64) *VatInvoiceScanQueryRequest {
	s.InvoiceSubTaskId = &v
	return s
}

func (s *VatInvoiceScanQueryRequest) SetPageNo(v int32) *VatInvoiceScanQueryRequest {
	s.PageNo = &v
	return s
}

func (s *VatInvoiceScanQueryRequest) SetPageSize(v int32) *VatInvoiceScanQueryRequest {
	s.PageSize = &v
	return s
}

func (s *VatInvoiceScanQueryRequest) Validate() error {
	return dara.Validate(s)
}
