// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewDesktopGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *RenewDesktopGroupRequest
	GetAutoPay() *bool
	SetAutoRenew(v bool) *RenewDesktopGroupRequest
	GetAutoRenew() *bool
	SetDesktopGroupId(v string) *RenewDesktopGroupRequest
	GetDesktopGroupId() *string
	SetPeriod(v int32) *RenewDesktopGroupRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *RenewDesktopGroupRequest
	GetPeriodUnit() *string
	SetRegionId(v string) *RenewDesktopGroupRequest
	GetRegionId() *string
	SetResellerOwnerUid(v int64) *RenewDesktopGroupRequest
	GetResellerOwnerUid() *int64
}

type RenewDesktopGroupRequest struct {
	// Specifies whether to enable automatic payment.
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to enable auto-renewal.
	//
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The ID of the shared cloud computer.
	//
	// This parameter is required.
	//
	// example:
	//
	// dg-7724r1jitbjzc****
	DesktopGroupId *string `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
	// The renewal duration. Valid values of this parameter are determined by the value of the `PeriodUnit` parameter.
	//
	// - If `PeriodUnit` is set to `Month`, valid values are 1, 2, 3, and 6.
	//
	// - If `PeriodUnit` is set to `Year`, valid values are 1 to 5.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the renewal duration, which is the unit of the `Period` parameter.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The region ID. You can call [DescribeRegions](~~DescribeRegions~~) to query the regions supported by Elastic Desktop Service.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID of the resource ownership in the resale pattern. You do not need to specify this parameter if you are not using the resale pattern.
	//
	// example:
	//
	// 1422724566551XXX
	ResellerOwnerUid *int64 `json:"ResellerOwnerUid,omitempty" xml:"ResellerOwnerUid,omitempty"`
}

func (s RenewDesktopGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s RenewDesktopGroupRequest) GoString() string {
	return s.String()
}

func (s *RenewDesktopGroupRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *RenewDesktopGroupRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *RenewDesktopGroupRequest) GetDesktopGroupId() *string {
	return s.DesktopGroupId
}

func (s *RenewDesktopGroupRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *RenewDesktopGroupRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *RenewDesktopGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RenewDesktopGroupRequest) GetResellerOwnerUid() *int64 {
	return s.ResellerOwnerUid
}

func (s *RenewDesktopGroupRequest) SetAutoPay(v bool) *RenewDesktopGroupRequest {
	s.AutoPay = &v
	return s
}

func (s *RenewDesktopGroupRequest) SetAutoRenew(v bool) *RenewDesktopGroupRequest {
	s.AutoRenew = &v
	return s
}

func (s *RenewDesktopGroupRequest) SetDesktopGroupId(v string) *RenewDesktopGroupRequest {
	s.DesktopGroupId = &v
	return s
}

func (s *RenewDesktopGroupRequest) SetPeriod(v int32) *RenewDesktopGroupRequest {
	s.Period = &v
	return s
}

func (s *RenewDesktopGroupRequest) SetPeriodUnit(v string) *RenewDesktopGroupRequest {
	s.PeriodUnit = &v
	return s
}

func (s *RenewDesktopGroupRequest) SetRegionId(v string) *RenewDesktopGroupRequest {
	s.RegionId = &v
	return s
}

func (s *RenewDesktopGroupRequest) SetResellerOwnerUid(v int64) *RenewDesktopGroupRequest {
	s.ResellerOwnerUid = &v
	return s
}

func (s *RenewDesktopGroupRequest) Validate() error {
	return dara.Validate(s)
}
