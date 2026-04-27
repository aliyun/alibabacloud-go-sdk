// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBillingCostBreakdownRespDTO interface {
	dara.Model
	String() string
	GoString() string
	SetGranularity(v string) *BillingCostBreakdownRespDTO
	GetGranularity() *string
	SetPage(v int32) *BillingCostBreakdownRespDTO
	GetPage() *int32
	SetPageSize(v int32) *BillingCostBreakdownRespDTO
	GetPageSize() *int32
	SetRows(v []*BillingCostBreakdownRowDTO) *BillingCostBreakdownRespDTO
	GetRows() []*BillingCostBreakdownRowDTO
	SetTotal(v int64) *BillingCostBreakdownRespDTO
	GetTotal() *int64
}

type BillingCostBreakdownRespDTO struct {
	// example:
	//
	// hourly
	Granularity *string `json:"granularity,omitempty" xml:"granularity,omitempty"`
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// []
	Rows []*BillingCostBreakdownRowDTO `json:"rows,omitempty" xml:"rows,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s BillingCostBreakdownRespDTO) String() string {
	return dara.Prettify(s)
}

func (s BillingCostBreakdownRespDTO) GoString() string {
	return s.String()
}

func (s *BillingCostBreakdownRespDTO) GetGranularity() *string {
	return s.Granularity
}

func (s *BillingCostBreakdownRespDTO) GetPage() *int32 {
	return s.Page
}

func (s *BillingCostBreakdownRespDTO) GetPageSize() *int32 {
	return s.PageSize
}

func (s *BillingCostBreakdownRespDTO) GetRows() []*BillingCostBreakdownRowDTO {
	return s.Rows
}

func (s *BillingCostBreakdownRespDTO) GetTotal() *int64 {
	return s.Total
}

func (s *BillingCostBreakdownRespDTO) SetGranularity(v string) *BillingCostBreakdownRespDTO {
	s.Granularity = &v
	return s
}

func (s *BillingCostBreakdownRespDTO) SetPage(v int32) *BillingCostBreakdownRespDTO {
	s.Page = &v
	return s
}

func (s *BillingCostBreakdownRespDTO) SetPageSize(v int32) *BillingCostBreakdownRespDTO {
	s.PageSize = &v
	return s
}

func (s *BillingCostBreakdownRespDTO) SetRows(v []*BillingCostBreakdownRowDTO) *BillingCostBreakdownRespDTO {
	s.Rows = v
	return s
}

func (s *BillingCostBreakdownRespDTO) SetTotal(v int64) *BillingCostBreakdownRespDTO {
	s.Total = &v
	return s
}

func (s *BillingCostBreakdownRespDTO) Validate() error {
	if s.Rows != nil {
		for _, item := range s.Rows {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
