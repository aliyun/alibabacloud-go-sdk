// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewARMServerInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRenew(v bool) *RenewARMServerInstanceRequest
	GetAutoRenew() *bool
	SetInstanceId(v string) *RenewARMServerInstanceRequest
	GetInstanceId() *string
	SetPeriod(v int32) *RenewARMServerInstanceRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *RenewARMServerInstanceRequest
	GetPeriodUnit() *string
}

type RenewARMServerInstanceRequest struct {
	// Specifies whether to enable auto-renewal for the premium bandwidth plan. Valid values:
	//
	// 	- **true**.
	//
	// 	- **false*	- (default).
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The ID of the instance that you want to renew.
	//
	// This parameter is required.
	//
	// example:
	//
	// yourInstance ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The renewal period. By default, instances are renewed on a monthly basis. Valid values: 1, 2, 3, 4, 5, 6, 7, 8, 9, and 12.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the renewal period. Valid values:
	//
	// 	- Month (default)
	//
	// 	- Year
	//
	// This parameter is required.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
}

func (s RenewARMServerInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RenewARMServerInstanceRequest) GoString() string {
	return s.String()
}

func (s *RenewARMServerInstanceRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *RenewARMServerInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RenewARMServerInstanceRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *RenewARMServerInstanceRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *RenewARMServerInstanceRequest) SetAutoRenew(v bool) *RenewARMServerInstanceRequest {
	s.AutoRenew = &v
	return s
}

func (s *RenewARMServerInstanceRequest) SetInstanceId(v string) *RenewARMServerInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *RenewARMServerInstanceRequest) SetPeriod(v int32) *RenewARMServerInstanceRequest {
	s.Period = &v
	return s
}

func (s *RenewARMServerInstanceRequest) SetPeriodUnit(v string) *RenewARMServerInstanceRequest {
	s.PeriodUnit = &v
	return s
}

func (s *RenewARMServerInstanceRequest) Validate() error {
	return dara.Validate(s)
}
