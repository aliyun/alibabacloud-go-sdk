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
	// Specifies whether to enable the auto-payment feature.
	//
	// Valid values:
	//
	// 	- true (default): enables the auto-payment feature. Make sure that your account balance is sufficient. Otherwise, an abnormal order is generated.
	//
	// 	- false: disables the auto-payment feature. In this case, an order is generated but you need to make the payment manually. You can log on to the EDS console and complete the payment based on the order ID on the Orders page.
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to enable auto-renewal.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The ID of the shared group.
	//
	// This parameter is required.
	//
	// example:
	//
	// dg-7724r1jitbjzc****
	DesktopGroupId *string `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
	// The renewal duration. Valid values of this parameter are determined by the value of the `PeriodUnit` parameter.
	//
	// 	- Valid values if you set the `PeriodUnit` parameter to `Month`: 1, 2, 3, and 6
	//
	// 	- Valid values if you set the `PeriodUnit` parameter to `Year`: 1, 2, 3, 4, and 5
	//
	// Default value: 1
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the renewal duration specified by the `Period` parameter.
	//
	// Valid values:
	//
	// 	- Month (default)
	//
	// 	- Year
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The region ID. You can call the [DescribeRegions](~~DescribeRegions~~) operation to query the regions supported by Elastic Desktop Service (EDS).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResellerOwnerUid *int64  `json:"ResellerOwnerUid,omitempty" xml:"ResellerOwnerUid,omitempty"`
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
