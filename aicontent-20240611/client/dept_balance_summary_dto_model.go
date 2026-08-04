// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeptBalanceSummaryDTO interface {
	dara.Model
	String() string
	GoString() string
	SetMonthly(v *BalancePoolSummaryDTO) *DeptBalanceSummaryDTO
	GetMonthly() *BalancePoolSummaryDTO
	SetPermanent(v *BalancePoolSummaryDTO) *DeptBalanceSummaryDTO
	GetPermanent() *BalancePoolSummaryDTO
}

type DeptBalanceSummaryDTO struct {
	// example:
	//
	// {}
	Monthly *BalancePoolSummaryDTO `json:"monthly,omitempty" xml:"monthly,omitempty"`
	// example:
	//
	// {}
	Permanent *BalancePoolSummaryDTO `json:"permanent,omitempty" xml:"permanent,omitempty"`
}

func (s DeptBalanceSummaryDTO) String() string {
	return dara.Prettify(s)
}

func (s DeptBalanceSummaryDTO) GoString() string {
	return s.String()
}

func (s *DeptBalanceSummaryDTO) GetMonthly() *BalancePoolSummaryDTO {
	return s.Monthly
}

func (s *DeptBalanceSummaryDTO) GetPermanent() *BalancePoolSummaryDTO {
	return s.Permanent
}

func (s *DeptBalanceSummaryDTO) SetMonthly(v *BalancePoolSummaryDTO) *DeptBalanceSummaryDTO {
	s.Monthly = v
	return s
}

func (s *DeptBalanceSummaryDTO) SetPermanent(v *BalancePoolSummaryDTO) *DeptBalanceSummaryDTO {
	s.Permanent = v
	return s
}

func (s *DeptBalanceSummaryDTO) Validate() error {
	if s.Monthly != nil {
		if err := s.Monthly.Validate(); err != nil {
			return err
		}
	}
	if s.Permanent != nil {
		if err := s.Permanent.Validate(); err != nil {
			return err
		}
	}
	return nil
}
