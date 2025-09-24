// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewResourcePackageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDuration(v int32) *RenewResourcePackageRequest
	GetDuration() *int32
	SetEffectiveDate(v string) *RenewResourcePackageRequest
	GetEffectiveDate() *string
	SetInstanceId(v string) *RenewResourcePackageRequest
	GetInstanceId() *string
	SetOwnerId(v int64) *RenewResourcePackageRequest
	GetOwnerId() *int64
	SetPricingCycle(v string) *RenewResourcePackageRequest
	GetPricingCycle() *string
}

type RenewResourcePackageRequest struct {
	// The renewal period of the resource plan.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The time when the resource plan takes effect. If you leave this parameter empty, the resource plan immediately takes effect by default.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2020-02-02T12:00:00Z
	EffectiveDate *string `json:"EffectiveDate,omitempty" xml:"EffectiveDate,omitempty"`
	// The ID of the resource plan.
	//
	// This parameter is required.
	//
	// example:
	//
	// lskd****sljhsdj
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The unit of the validity period for the resource plan. Valid values:
	//
	// 	- Month
	//
	// 	- Year
	//
	// Default value: Month.
	//
	// This parameter is required.
	//
	// example:
	//
	// Month
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
}

func (s RenewResourcePackageRequest) String() string {
	return dara.Prettify(s)
}

func (s RenewResourcePackageRequest) GoString() string {
	return s.String()
}

func (s *RenewResourcePackageRequest) GetDuration() *int32 {
	return s.Duration
}

func (s *RenewResourcePackageRequest) GetEffectiveDate() *string {
	return s.EffectiveDate
}

func (s *RenewResourcePackageRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RenewResourcePackageRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RenewResourcePackageRequest) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *RenewResourcePackageRequest) SetDuration(v int32) *RenewResourcePackageRequest {
	s.Duration = &v
	return s
}

func (s *RenewResourcePackageRequest) SetEffectiveDate(v string) *RenewResourcePackageRequest {
	s.EffectiveDate = &v
	return s
}

func (s *RenewResourcePackageRequest) SetInstanceId(v string) *RenewResourcePackageRequest {
	s.InstanceId = &v
	return s
}

func (s *RenewResourcePackageRequest) SetOwnerId(v int64) *RenewResourcePackageRequest {
	s.OwnerId = &v
	return s
}

func (s *RenewResourcePackageRequest) SetPricingCycle(v string) *RenewResourcePackageRequest {
	s.PricingCycle = &v
	return s
}

func (s *RenewResourcePackageRequest) Validate() error {
	return dara.Validate(s)
}
