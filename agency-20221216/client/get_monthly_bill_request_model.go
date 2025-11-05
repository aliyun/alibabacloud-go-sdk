// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMonthlyBillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillOwner(v string) *GetMonthlyBillRequest
	GetBillOwner() *string
	SetBillType(v string) *GetMonthlyBillRequest
	GetBillType() *string
	SetMonth(v string) *GetMonthlyBillRequest
	GetMonth() *string
}

type GetMonthlyBillRequest struct {
	// Bill Owner type. Value Range:</br>
	//
	// 1: Master account</br>
	//
	// 2: Sub account</br>
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	BillOwner *string `json:"BillOwner,omitempty" xml:"BillOwner,omitempty"`
	// Value Range:
	//
	// - MonthlyInvoice
	//
	// - MonthRefundInvoice
	//
	// - MonthlySummary (Deprecated)
	//
	// - MonthlyInstanceAddAdjustBill
	//
	// - MonthlyInstanceRefundBill
	//
	// - MonthlyAddAdjustInvoce
	//
	// - MonthlyRefundAdjustInvoce
	//
	// - MonthlyInstanceConsumeV2
	//
	// - MarginReportV2
	//
	// This parameter is required.
	//
	// example:
	//
	// MonthlyInvoice
	BillType *string `json:"BillType,omitempty" xml:"BillType,omitempty"`
	// Billing Month, Format is YYYY-MM
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-11
	Month *string `json:"Month,omitempty" xml:"Month,omitempty"`
}

func (s GetMonthlyBillRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMonthlyBillRequest) GoString() string {
	return s.String()
}

func (s *GetMonthlyBillRequest) GetBillOwner() *string {
	return s.BillOwner
}

func (s *GetMonthlyBillRequest) GetBillType() *string {
	return s.BillType
}

func (s *GetMonthlyBillRequest) GetMonth() *string {
	return s.Month
}

func (s *GetMonthlyBillRequest) SetBillOwner(v string) *GetMonthlyBillRequest {
	s.BillOwner = &v
	return s
}

func (s *GetMonthlyBillRequest) SetBillType(v string) *GetMonthlyBillRequest {
	s.BillType = &v
	return s
}

func (s *GetMonthlyBillRequest) SetMonth(v string) *GetMonthlyBillRequest {
	s.Month = &v
	return s
}

func (s *GetMonthlyBillRequest) Validate() error {
	return dara.Validate(s)
}
