// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyInvoiceTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillDate(v string) *ApplyInvoiceTaskShrinkRequest
	GetBillDate() *string
	SetInvoiceTaskListShrink(v string) *ApplyInvoiceTaskShrinkRequest
	GetInvoiceTaskListShrink() *string
}

type ApplyInvoiceTaskShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2022-12-01
	BillDate *string `json:"bill_date,omitempty" xml:"bill_date,omitempty"`
	// This parameter is required.
	InvoiceTaskListShrink *string `json:"invoice_task_list,omitempty" xml:"invoice_task_list,omitempty"`
}

func (s ApplyInvoiceTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ApplyInvoiceTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *ApplyInvoiceTaskShrinkRequest) GetBillDate() *string {
	return s.BillDate
}

func (s *ApplyInvoiceTaskShrinkRequest) GetInvoiceTaskListShrink() *string {
	return s.InvoiceTaskListShrink
}

func (s *ApplyInvoiceTaskShrinkRequest) SetBillDate(v string) *ApplyInvoiceTaskShrinkRequest {
	s.BillDate = &v
	return s
}

func (s *ApplyInvoiceTaskShrinkRequest) SetInvoiceTaskListShrink(v string) *ApplyInvoiceTaskShrinkRequest {
	s.InvoiceTaskListShrink = &v
	return s
}

func (s *ApplyInvoiceTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
