// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArchitectureType(v string) *CreateInstanceShrinkRequest
	GetArchitectureType() *string
	SetAutoRenew(v bool) *CreateInstanceShrinkRequest
	GetAutoRenew() *bool
	SetChargeType(v string) *CreateInstanceShrinkRequest
	GetChargeType() *string
	SetDuration(v int32) *CreateInstanceShrinkRequest
	GetDuration() *int32
	SetExtra(v string) *CreateInstanceShrinkRequest
	GetExtra() *string
	SetHa(v bool) *CreateInstanceShrinkRequest
	GetHa() *bool
	SetHaResourceSpecShrink(v string) *CreateInstanceShrinkRequest
	GetHaResourceSpecShrink() *string
	SetHaVSwitchIdsShrink(v string) *CreateInstanceShrinkRequest
	GetHaVSwitchIdsShrink() *string
	SetInstanceName(v string) *CreateInstanceShrinkRequest
	GetInstanceName() *string
	SetMonitorType(v string) *CreateInstanceShrinkRequest
	GetMonitorType() *string
	SetPricingCycle(v string) *CreateInstanceShrinkRequest
	GetPricingCycle() *string
	SetPromotionCode(v string) *CreateInstanceShrinkRequest
	GetPromotionCode() *string
	SetRegion(v string) *CreateInstanceShrinkRequest
	GetRegion() *string
	SetResourceGroupId(v string) *CreateInstanceShrinkRequest
	GetResourceGroupId() *string
	SetResourceSpecShrink(v string) *CreateInstanceShrinkRequest
	GetResourceSpecShrink() *string
	SetStorageShrink(v string) *CreateInstanceShrinkRequest
	GetStorageShrink() *string
	SetTagShrink(v string) *CreateInstanceShrinkRequest
	GetTagShrink() *string
	SetUsePromotionCode(v bool) *CreateInstanceShrinkRequest
	GetUsePromotionCode() *bool
	SetVSwitchIdsShrink(v string) *CreateInstanceShrinkRequest
	GetVSwitchIdsShrink() *string
	SetVpcId(v string) *CreateInstanceShrinkRequest
	GetVpcId() *string
}

type CreateInstanceShrinkRequest struct {
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
	// if can be null:
	// true
	Ha *bool `json:"Ha,omitempty" xml:"Ha,omitempty"`
	// if can be null:
	// true
	HaResourceSpecShrink *string `json:"HaResourceSpec,omitempty" xml:"HaResourceSpec,omitempty"`
	// if can be null:
	// true
	HaVSwitchIdsShrink *string `json:"HaVSwitchIds,omitempty" xml:"HaVSwitchIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// rtc-e2e-test-pre
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	MonitorType  *string `json:"MonitorType,omitempty" xml:"MonitorType,omitempty"`
	// example:
	//
	// Month
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// example:
	//
	// 500043499350689
	PromotionCode *string `json:"PromotionCode,omitempty" xml:"PromotionCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	Region             *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceGroupId    *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceSpecShrink *string `json:"ResourceSpec,omitempty" xml:"ResourceSpec,omitempty"`
	// This parameter is required.
	StorageShrink    *string `json:"Storage,omitempty" xml:"Storage,omitempty"`
	TagShrink        *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	UsePromotionCode *bool   `json:"UsePromotionCode,omitempty" xml:"UsePromotionCode,omitempty"`
	// This parameter is required.
	VSwitchIdsShrink *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vpc-2ze9xoh8qyt1rnxfmfcdi
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateInstanceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceShrinkRequest) GetArchitectureType() *string {
	return s.ArchitectureType
}

func (s *CreateInstanceShrinkRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateInstanceShrinkRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *CreateInstanceShrinkRequest) GetDuration() *int32 {
	return s.Duration
}

func (s *CreateInstanceShrinkRequest) GetExtra() *string {
	return s.Extra
}

func (s *CreateInstanceShrinkRequest) GetHa() *bool {
	return s.Ha
}

func (s *CreateInstanceShrinkRequest) GetHaResourceSpecShrink() *string {
	return s.HaResourceSpecShrink
}

func (s *CreateInstanceShrinkRequest) GetHaVSwitchIdsShrink() *string {
	return s.HaVSwitchIdsShrink
}

func (s *CreateInstanceShrinkRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CreateInstanceShrinkRequest) GetMonitorType() *string {
	return s.MonitorType
}

func (s *CreateInstanceShrinkRequest) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *CreateInstanceShrinkRequest) GetPromotionCode() *string {
	return s.PromotionCode
}

func (s *CreateInstanceShrinkRequest) GetRegion() *string {
	return s.Region
}

func (s *CreateInstanceShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateInstanceShrinkRequest) GetResourceSpecShrink() *string {
	return s.ResourceSpecShrink
}

func (s *CreateInstanceShrinkRequest) GetStorageShrink() *string {
	return s.StorageShrink
}

func (s *CreateInstanceShrinkRequest) GetTagShrink() *string {
	return s.TagShrink
}

func (s *CreateInstanceShrinkRequest) GetUsePromotionCode() *bool {
	return s.UsePromotionCode
}

func (s *CreateInstanceShrinkRequest) GetVSwitchIdsShrink() *string {
	return s.VSwitchIdsShrink
}

func (s *CreateInstanceShrinkRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateInstanceShrinkRequest) SetArchitectureType(v string) *CreateInstanceShrinkRequest {
	s.ArchitectureType = &v
	return s
}

func (s *CreateInstanceShrinkRequest) SetAutoRenew(v bool) *CreateInstanceShrinkRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateInstanceShrinkRequest) SetChargeType(v string) *CreateInstanceShrinkRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateInstanceShrinkRequest) SetDuration(v int32) *CreateInstanceShrinkRequest {
	s.Duration = &v
	return s
}

func (s *CreateInstanceShrinkRequest) SetExtra(v string) *CreateInstanceShrinkRequest {
	s.Extra = &v
	return s
}

func (s *CreateInstanceShrinkRequest) SetHa(v bool) *CreateInstanceShrinkRequest {
	s.Ha = &v
	return s
}

func (s *CreateInstanceShrinkRequest) SetHaResourceSpecShrink(v string) *CreateInstanceShrinkRequest {
	s.HaResourceSpecShrink = &v
	return s
}

func (s *CreateInstanceShrinkRequest) SetHaVSwitchIdsShrink(v string) *CreateInstanceShrinkRequest {
	s.HaVSwitchIdsShrink = &v
	return s
}

func (s *CreateInstanceShrinkRequest) SetInstanceName(v string) *CreateInstanceShrinkRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateInstanceShrinkRequest) SetMonitorType(v string) *CreateInstanceShrinkRequest {
	s.MonitorType = &v
	return s
}

func (s *CreateInstanceShrinkRequest) SetPricingCycle(v string) *CreateInstanceShrinkRequest {
	s.PricingCycle = &v
	return s
}

func (s *CreateInstanceShrinkRequest) SetPromotionCode(v string) *CreateInstanceShrinkRequest {
	s.PromotionCode = &v
	return s
}

func (s *CreateInstanceShrinkRequest) SetRegion(v string) *CreateInstanceShrinkRequest {
	s.Region = &v
	return s
}

func (s *CreateInstanceShrinkRequest) SetResourceGroupId(v string) *CreateInstanceShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateInstanceShrinkRequest) SetResourceSpecShrink(v string) *CreateInstanceShrinkRequest {
	s.ResourceSpecShrink = &v
	return s
}

func (s *CreateInstanceShrinkRequest) SetStorageShrink(v string) *CreateInstanceShrinkRequest {
	s.StorageShrink = &v
	return s
}

func (s *CreateInstanceShrinkRequest) SetTagShrink(v string) *CreateInstanceShrinkRequest {
	s.TagShrink = &v
	return s
}

func (s *CreateInstanceShrinkRequest) SetUsePromotionCode(v bool) *CreateInstanceShrinkRequest {
	s.UsePromotionCode = &v
	return s
}

func (s *CreateInstanceShrinkRequest) SetVSwitchIdsShrink(v string) *CreateInstanceShrinkRequest {
	s.VSwitchIdsShrink = &v
	return s
}

func (s *CreateInstanceShrinkRequest) SetVpcId(v string) *CreateInstanceShrinkRequest {
	s.VpcId = &v
	return s
}

func (s *CreateInstanceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
