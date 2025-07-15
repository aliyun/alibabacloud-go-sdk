// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewDesktopOversoldGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOversoldGroupId(v string) *RenewDesktopOversoldGroupRequest
	GetOversoldGroupId() *string
	SetPeriod(v int32) *RenewDesktopOversoldGroupRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *RenewDesktopOversoldGroupRequest
	GetPeriodUnit() *string
}

type RenewDesktopOversoldGroupRequest struct {
	OversoldGroupId *string `json:"OversoldGroupId,omitempty" xml:"OversoldGroupId,omitempty"`
	Period          *int32  `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit      *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
}

func (s RenewDesktopOversoldGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s RenewDesktopOversoldGroupRequest) GoString() string {
	return s.String()
}

func (s *RenewDesktopOversoldGroupRequest) GetOversoldGroupId() *string {
	return s.OversoldGroupId
}

func (s *RenewDesktopOversoldGroupRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *RenewDesktopOversoldGroupRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *RenewDesktopOversoldGroupRequest) SetOversoldGroupId(v string) *RenewDesktopOversoldGroupRequest {
	s.OversoldGroupId = &v
	return s
}

func (s *RenewDesktopOversoldGroupRequest) SetPeriod(v int32) *RenewDesktopOversoldGroupRequest {
	s.Period = &v
	return s
}

func (s *RenewDesktopOversoldGroupRequest) SetPeriodUnit(v string) *RenewDesktopOversoldGroupRequest {
	s.PeriodUnit = &v
	return s
}

func (s *RenewDesktopOversoldGroupRequest) Validate() error {
	return dara.Validate(s)
}
