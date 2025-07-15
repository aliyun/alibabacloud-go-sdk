// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePriceForRenewDesktopOversoldGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOversoldGroupId(v string) *DescribePriceForRenewDesktopOversoldGroupRequest
	GetOversoldGroupId() *string
	SetPeriod(v int32) *DescribePriceForRenewDesktopOversoldGroupRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *DescribePriceForRenewDesktopOversoldGroupRequest
	GetPeriodUnit() *string
}

type DescribePriceForRenewDesktopOversoldGroupRequest struct {
	OversoldGroupId *string `json:"OversoldGroupId,omitempty" xml:"OversoldGroupId,omitempty"`
	Period          *int32  `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit      *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
}

func (s DescribePriceForRenewDesktopOversoldGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceForRenewDesktopOversoldGroupRequest) GoString() string {
	return s.String()
}

func (s *DescribePriceForRenewDesktopOversoldGroupRequest) GetOversoldGroupId() *string {
	return s.OversoldGroupId
}

func (s *DescribePriceForRenewDesktopOversoldGroupRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *DescribePriceForRenewDesktopOversoldGroupRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *DescribePriceForRenewDesktopOversoldGroupRequest) SetOversoldGroupId(v string) *DescribePriceForRenewDesktopOversoldGroupRequest {
	s.OversoldGroupId = &v
	return s
}

func (s *DescribePriceForRenewDesktopOversoldGroupRequest) SetPeriod(v int32) *DescribePriceForRenewDesktopOversoldGroupRequest {
	s.Period = &v
	return s
}

func (s *DescribePriceForRenewDesktopOversoldGroupRequest) SetPeriodUnit(v string) *DescribePriceForRenewDesktopOversoldGroupRequest {
	s.PeriodUnit = &v
	return s
}

func (s *DescribePriceForRenewDesktopOversoldGroupRequest) Validate() error {
	return dara.Validate(s)
}
