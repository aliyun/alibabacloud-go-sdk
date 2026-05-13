// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBillingBillSummaryPointDTO interface {
	dara.Model
	String() string
	GoString() string
	SetTimestamp(v int64) *BillingBillSummaryPointDTO
	GetTimestamp() *int64
	SetTotalAmount(v float32) *BillingBillSummaryPointDTO
	GetTotalAmount() *float32
}

type BillingBillSummaryPointDTO struct {
	// example:
	//
	// 1700000000
	Timestamp *int64 `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	// example:
	//
	// 1.23
	TotalAmount *float32 `json:"totalAmount,omitempty" xml:"totalAmount,omitempty"`
}

func (s BillingBillSummaryPointDTO) String() string {
	return dara.Prettify(s)
}

func (s BillingBillSummaryPointDTO) GoString() string {
	return s.String()
}

func (s *BillingBillSummaryPointDTO) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *BillingBillSummaryPointDTO) GetTotalAmount() *float32 {
	return s.TotalAmount
}

func (s *BillingBillSummaryPointDTO) SetTimestamp(v int64) *BillingBillSummaryPointDTO {
	s.Timestamp = &v
	return s
}

func (s *BillingBillSummaryPointDTO) SetTotalAmount(v float32) *BillingBillSummaryPointDTO {
	s.TotalAmount = &v
	return s
}

func (s *BillingBillSummaryPointDTO) Validate() error {
	return dara.Validate(s)
}
