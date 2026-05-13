// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBillingBillTierDTO interface {
	dara.Model
	String() string
	GoString() string
	SetDimValues(v string) *BillingBillTierDTO
	GetDimValues() *string
	SetPayableAmount(v float32) *BillingBillTierDTO
	GetPayableAmount() *float32
	SetValues(v string) *BillingBillTierDTO
	GetValues() *string
}

type BillingBillTierDTO struct {
	// example:
	//
	// {"context_tier": "0-32k"}
	DimValues *string `json:"dimValues,omitempty" xml:"dimValues,omitempty"`
	// example:
	//
	// 0.05
	PayableAmount *float32 `json:"payableAmount,omitempty" xml:"payableAmount,omitempty"`
	// example:
	//
	// {"input_tokens": 1000, "output_tokens": 500}
	Values *string `json:"values,omitempty" xml:"values,omitempty"`
}

func (s BillingBillTierDTO) String() string {
	return dara.Prettify(s)
}

func (s BillingBillTierDTO) GoString() string {
	return s.String()
}

func (s *BillingBillTierDTO) GetDimValues() *string {
	return s.DimValues
}

func (s *BillingBillTierDTO) GetPayableAmount() *float32 {
	return s.PayableAmount
}

func (s *BillingBillTierDTO) GetValues() *string {
	return s.Values
}

func (s *BillingBillTierDTO) SetDimValues(v string) *BillingBillTierDTO {
	s.DimValues = &v
	return s
}

func (s *BillingBillTierDTO) SetPayableAmount(v float32) *BillingBillTierDTO {
	s.PayableAmount = &v
	return s
}

func (s *BillingBillTierDTO) SetValues(v string) *BillingBillTierDTO {
	s.Values = &v
	return s
}

func (s *BillingBillTierDTO) Validate() error {
	return dara.Validate(s)
}
