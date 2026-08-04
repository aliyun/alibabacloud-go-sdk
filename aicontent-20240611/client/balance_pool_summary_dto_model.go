// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBalancePoolSummaryDTO interface {
	dara.Model
	String() string
	GoString() string
	SetAllocated(v float64) *BalancePoolSummaryDTO
	GetAllocated() *float64
	SetAvailable(v float64) *BalancePoolSummaryDTO
	GetAvailable() *float64
	SetTotal(v float64) *BalancePoolSummaryDTO
	GetTotal() *float64
}

type BalancePoolSummaryDTO struct {
	// example:
	//
	// 40.00
	Allocated *float64 `json:"allocated,omitempty" xml:"allocated,omitempty"`
	// example:
	//
	// 60.00
	Available *float64 `json:"available,omitempty" xml:"available,omitempty"`
	// example:
	//
	// 100.00
	Total *float64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s BalancePoolSummaryDTO) String() string {
	return dara.Prettify(s)
}

func (s BalancePoolSummaryDTO) GoString() string {
	return s.String()
}

func (s *BalancePoolSummaryDTO) GetAllocated() *float64 {
	return s.Allocated
}

func (s *BalancePoolSummaryDTO) GetAvailable() *float64 {
	return s.Available
}

func (s *BalancePoolSummaryDTO) GetTotal() *float64 {
	return s.Total
}

func (s *BalancePoolSummaryDTO) SetAllocated(v float64) *BalancePoolSummaryDTO {
	s.Allocated = &v
	return s
}

func (s *BalancePoolSummaryDTO) SetAvailable(v float64) *BalancePoolSummaryDTO {
	s.Available = &v
	return s
}

func (s *BalancePoolSummaryDTO) SetTotal(v float64) *BalancePoolSummaryDTO {
	s.Total = &v
	return s
}

func (s *BalancePoolSummaryDTO) Validate() error {
	return dara.Validate(s)
}
