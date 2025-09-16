// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCreateInstancePriceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArchitectureType(v string) *QueryCreateInstancePriceShrinkRequest
	GetArchitectureType() *string
	SetAutoRenew(v bool) *QueryCreateInstancePriceShrinkRequest
	GetAutoRenew() *bool
	SetChargeType(v string) *QueryCreateInstancePriceShrinkRequest
	GetChargeType() *string
	SetDuration(v int32) *QueryCreateInstancePriceShrinkRequest
	GetDuration() *int32
	SetExtra(v string) *QueryCreateInstancePriceShrinkRequest
	GetExtra() *string
	SetHa(v bool) *QueryCreateInstancePriceShrinkRequest
	GetHa() *bool
	SetHaResourceSpecShrink(v string) *QueryCreateInstancePriceShrinkRequest
	GetHaResourceSpecShrink() *string
	SetInstanceName(v string) *QueryCreateInstancePriceShrinkRequest
	GetInstanceName() *string
	SetPricingCycle(v string) *QueryCreateInstancePriceShrinkRequest
	GetPricingCycle() *string
	SetPromotionCode(v string) *QueryCreateInstancePriceShrinkRequest
	GetPromotionCode() *string
	SetRegion(v string) *QueryCreateInstancePriceShrinkRequest
	GetRegion() *string
	SetResourceSpecShrink(v string) *QueryCreateInstancePriceShrinkRequest
	GetResourceSpecShrink() *string
	SetStorageShrink(v string) *QueryCreateInstancePriceShrinkRequest
	GetStorageShrink() *string
	SetUsePromotionCode(v bool) *QueryCreateInstancePriceShrinkRequest
	GetUsePromotionCode() *bool
	SetVSwitchIdsShrink(v string) *QueryCreateInstancePriceShrinkRequest
	GetVSwitchIdsShrink() *string
	SetVpcId(v string) *QueryCreateInstancePriceShrinkRequest
	GetVpcId() *string
}

type QueryCreateInstancePriceShrinkRequest struct {
	ArchitectureType *string `json:"ArchitectureType,omitempty" xml:"ArchitectureType,omitempty"`
	// example:
	//
	// true
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PRE
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// example:
	//
	// 1
	Duration *int32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Extra    *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	Ha       *bool   `json:"Ha,omitempty" xml:"Ha,omitempty"`
	// if can be null:
	// true
	HaResourceSpecShrink *string `json:"HaResourceSpec,omitempty" xml:"HaResourceSpec,omitempty"`
	// example:
	//
	// rtc-e2e-test-post
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// Month
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// example:
	//
	// 500041860100636
	PromotionCode *string `json:"PromotionCode,omitempty" xml:"PromotionCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	Region             *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceSpecShrink *string `json:"ResourceSpec,omitempty" xml:"ResourceSpec,omitempty"`
	StorageShrink      *string `json:"Storage,omitempty" xml:"Storage,omitempty"`
	UsePromotionCode   *bool   `json:"UsePromotionCode,omitempty" xml:"UsePromotionCode,omitempty"`
	VSwitchIdsShrink   *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	// example:
	//
	// vpc-2ze9xoh8qyt1rnxfmfcdi
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s QueryCreateInstancePriceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryCreateInstancePriceShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryCreateInstancePriceShrinkRequest) GetArchitectureType() *string {
	return s.ArchitectureType
}

func (s *QueryCreateInstancePriceShrinkRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *QueryCreateInstancePriceShrinkRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *QueryCreateInstancePriceShrinkRequest) GetDuration() *int32 {
	return s.Duration
}

func (s *QueryCreateInstancePriceShrinkRequest) GetExtra() *string {
	return s.Extra
}

func (s *QueryCreateInstancePriceShrinkRequest) GetHa() *bool {
	return s.Ha
}

func (s *QueryCreateInstancePriceShrinkRequest) GetHaResourceSpecShrink() *string {
	return s.HaResourceSpecShrink
}

func (s *QueryCreateInstancePriceShrinkRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *QueryCreateInstancePriceShrinkRequest) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *QueryCreateInstancePriceShrinkRequest) GetPromotionCode() *string {
	return s.PromotionCode
}

func (s *QueryCreateInstancePriceShrinkRequest) GetRegion() *string {
	return s.Region
}

func (s *QueryCreateInstancePriceShrinkRequest) GetResourceSpecShrink() *string {
	return s.ResourceSpecShrink
}

func (s *QueryCreateInstancePriceShrinkRequest) GetStorageShrink() *string {
	return s.StorageShrink
}

func (s *QueryCreateInstancePriceShrinkRequest) GetUsePromotionCode() *bool {
	return s.UsePromotionCode
}

func (s *QueryCreateInstancePriceShrinkRequest) GetVSwitchIdsShrink() *string {
	return s.VSwitchIdsShrink
}

func (s *QueryCreateInstancePriceShrinkRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *QueryCreateInstancePriceShrinkRequest) SetArchitectureType(v string) *QueryCreateInstancePriceShrinkRequest {
	s.ArchitectureType = &v
	return s
}

func (s *QueryCreateInstancePriceShrinkRequest) SetAutoRenew(v bool) *QueryCreateInstancePriceShrinkRequest {
	s.AutoRenew = &v
	return s
}

func (s *QueryCreateInstancePriceShrinkRequest) SetChargeType(v string) *QueryCreateInstancePriceShrinkRequest {
	s.ChargeType = &v
	return s
}

func (s *QueryCreateInstancePriceShrinkRequest) SetDuration(v int32) *QueryCreateInstancePriceShrinkRequest {
	s.Duration = &v
	return s
}

func (s *QueryCreateInstancePriceShrinkRequest) SetExtra(v string) *QueryCreateInstancePriceShrinkRequest {
	s.Extra = &v
	return s
}

func (s *QueryCreateInstancePriceShrinkRequest) SetHa(v bool) *QueryCreateInstancePriceShrinkRequest {
	s.Ha = &v
	return s
}

func (s *QueryCreateInstancePriceShrinkRequest) SetHaResourceSpecShrink(v string) *QueryCreateInstancePriceShrinkRequest {
	s.HaResourceSpecShrink = &v
	return s
}

func (s *QueryCreateInstancePriceShrinkRequest) SetInstanceName(v string) *QueryCreateInstancePriceShrinkRequest {
	s.InstanceName = &v
	return s
}

func (s *QueryCreateInstancePriceShrinkRequest) SetPricingCycle(v string) *QueryCreateInstancePriceShrinkRequest {
	s.PricingCycle = &v
	return s
}

func (s *QueryCreateInstancePriceShrinkRequest) SetPromotionCode(v string) *QueryCreateInstancePriceShrinkRequest {
	s.PromotionCode = &v
	return s
}

func (s *QueryCreateInstancePriceShrinkRequest) SetRegion(v string) *QueryCreateInstancePriceShrinkRequest {
	s.Region = &v
	return s
}

func (s *QueryCreateInstancePriceShrinkRequest) SetResourceSpecShrink(v string) *QueryCreateInstancePriceShrinkRequest {
	s.ResourceSpecShrink = &v
	return s
}

func (s *QueryCreateInstancePriceShrinkRequest) SetStorageShrink(v string) *QueryCreateInstancePriceShrinkRequest {
	s.StorageShrink = &v
	return s
}

func (s *QueryCreateInstancePriceShrinkRequest) SetUsePromotionCode(v bool) *QueryCreateInstancePriceShrinkRequest {
	s.UsePromotionCode = &v
	return s
}

func (s *QueryCreateInstancePriceShrinkRequest) SetVSwitchIdsShrink(v string) *QueryCreateInstancePriceShrinkRequest {
	s.VSwitchIdsShrink = &v
	return s
}

func (s *QueryCreateInstancePriceShrinkRequest) SetVpcId(v string) *QueryCreateInstancePriceShrinkRequest {
	s.VpcId = &v
	return s
}

func (s *QueryCreateInstancePriceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
