// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWaitApplyInvoiceTaskDetailQueryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillDate(v string) *WaitApplyInvoiceTaskDetailQueryRequest
	GetBillDate() *string
}

type WaitApplyInvoiceTaskDetailQueryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2022-12-01
	BillDate *string `json:"bill_date,omitempty" xml:"bill_date,omitempty"`
}

func (s WaitApplyInvoiceTaskDetailQueryRequest) String() string {
	return dara.Prettify(s)
}

func (s WaitApplyInvoiceTaskDetailQueryRequest) GoString() string {
	return s.String()
}

func (s *WaitApplyInvoiceTaskDetailQueryRequest) GetBillDate() *string {
	return s.BillDate
}

func (s *WaitApplyInvoiceTaskDetailQueryRequest) SetBillDate(v string) *WaitApplyInvoiceTaskDetailQueryRequest {
	s.BillDate = &v
	return s
}

func (s *WaitApplyInvoiceTaskDetailQueryRequest) Validate() error {
	return dara.Validate(s)
}
