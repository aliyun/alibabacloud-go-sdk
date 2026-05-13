// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBillingBillSummaryRespDTO interface {
	dara.Model
	String() string
	GoString() string
	SetPoints(v []*BillingBillSummaryPointDTO) *BillingBillSummaryRespDTO
	GetPoints() []*BillingBillSummaryPointDTO
	SetTotalAmount(v float32) *BillingBillSummaryRespDTO
	GetTotalAmount() *float32
}

type BillingBillSummaryRespDTO struct {
	Points []*BillingBillSummaryPointDTO `json:"points,omitempty" xml:"points,omitempty" type:"Repeated"`
	// example:
	//
	// 123.45
	TotalAmount *float32 `json:"totalAmount,omitempty" xml:"totalAmount,omitempty"`
}

func (s BillingBillSummaryRespDTO) String() string {
	return dara.Prettify(s)
}

func (s BillingBillSummaryRespDTO) GoString() string {
	return s.String()
}

func (s *BillingBillSummaryRespDTO) GetPoints() []*BillingBillSummaryPointDTO {
	return s.Points
}

func (s *BillingBillSummaryRespDTO) GetTotalAmount() *float32 {
	return s.TotalAmount
}

func (s *BillingBillSummaryRespDTO) SetPoints(v []*BillingBillSummaryPointDTO) *BillingBillSummaryRespDTO {
	s.Points = v
	return s
}

func (s *BillingBillSummaryRespDTO) SetTotalAmount(v float32) *BillingBillSummaryRespDTO {
	s.TotalAmount = &v
	return s
}

func (s *BillingBillSummaryRespDTO) Validate() error {
	if s.Points != nil {
		for _, item := range s.Points {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
