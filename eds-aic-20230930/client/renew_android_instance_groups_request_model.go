// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewAndroidInstanceGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *RenewAndroidInstanceGroupsRequest
	GetAutoPay() *bool
	SetInstanceGroupIds(v []*string) *RenewAndroidInstanceGroupsRequest
	GetInstanceGroupIds() []*string
	SetPeriod(v int32) *RenewAndroidInstanceGroupsRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *RenewAndroidInstanceGroupsRequest
	GetPeriodUnit() *string
}

type RenewAndroidInstanceGroupsRequest struct {
	// Specifies whether to enable the auto-payment feature.
	//
	// Valid values:
	//
	// 	- true: enables the auto-payment feature. Ensure your account has sufficient balance to use this feature.
	//
	// 	- false: disables the auto-payment feature. You need to manually complete the payment process.
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The IDs of the instance groups.
	InstanceGroupIds []*string `json:"InstanceGroupIds,omitempty" xml:"InstanceGroupIds,omitempty" type:"Repeated"`
	// The duration of the renewal, measured in units defined by PeriodUnit.
	//
	// example:
	//
	// 6
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the renewal duration. Default value: Month.
	//
	// Valid values:
	//
	// 	- Month
	//
	// 	- Year
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
}

func (s RenewAndroidInstanceGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s RenewAndroidInstanceGroupsRequest) GoString() string {
	return s.String()
}

func (s *RenewAndroidInstanceGroupsRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *RenewAndroidInstanceGroupsRequest) GetInstanceGroupIds() []*string {
	return s.InstanceGroupIds
}

func (s *RenewAndroidInstanceGroupsRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *RenewAndroidInstanceGroupsRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *RenewAndroidInstanceGroupsRequest) SetAutoPay(v bool) *RenewAndroidInstanceGroupsRequest {
	s.AutoPay = &v
	return s
}

func (s *RenewAndroidInstanceGroupsRequest) SetInstanceGroupIds(v []*string) *RenewAndroidInstanceGroupsRequest {
	s.InstanceGroupIds = v
	return s
}

func (s *RenewAndroidInstanceGroupsRequest) SetPeriod(v int32) *RenewAndroidInstanceGroupsRequest {
	s.Period = &v
	return s
}

func (s *RenewAndroidInstanceGroupsRequest) SetPeriodUnit(v string) *RenewAndroidInstanceGroupsRequest {
	s.PeriodUnit = &v
	return s
}

func (s *RenewAndroidInstanceGroupsRequest) Validate() error {
	return dara.Validate(s)
}
