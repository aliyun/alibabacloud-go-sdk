// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConvertInstanceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDuration(v int32) *ConvertInstanceShrinkRequest
	GetDuration() *int32
	SetInstanceId(v string) *ConvertInstanceShrinkRequest
	GetInstanceId() *string
	SetIsAutoRenew(v bool) *ConvertInstanceShrinkRequest
	GetIsAutoRenew() *bool
	SetNamespaceResourceSpecsShrink(v string) *ConvertInstanceShrinkRequest
	GetNamespaceResourceSpecsShrink() *string
	SetPricingCycle(v string) *ConvertInstanceShrinkRequest
	GetPricingCycle() *string
	SetPromotionCode(v string) *ConvertInstanceShrinkRequest
	GetPromotionCode() *string
	SetRegion(v string) *ConvertInstanceShrinkRequest
	GetRegion() *string
	SetUsePromotionCode(v bool) *ConvertInstanceShrinkRequest
	GetUsePromotionCode() *bool
}

type ConvertInstanceShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// f-cn-wwo36qj4g06
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	IsAutoRenew *bool `json:"IsAutoRenew,omitempty" xml:"IsAutoRenew,omitempty"`
	// This parameter is required.
	NamespaceResourceSpecsShrink *string `json:"NamespaceResourceSpecs,omitempty" xml:"NamespaceResourceSpecs,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Month
	PricingCycle  *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	PromotionCode *string `json:"PromotionCode,omitempty" xml:"PromotionCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	Region           *string `json:"Region,omitempty" xml:"Region,omitempty"`
	UsePromotionCode *bool   `json:"UsePromotionCode,omitempty" xml:"UsePromotionCode,omitempty"`
}

func (s ConvertInstanceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ConvertInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *ConvertInstanceShrinkRequest) GetDuration() *int32 {
	return s.Duration
}

func (s *ConvertInstanceShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ConvertInstanceShrinkRequest) GetIsAutoRenew() *bool {
	return s.IsAutoRenew
}

func (s *ConvertInstanceShrinkRequest) GetNamespaceResourceSpecsShrink() *string {
	return s.NamespaceResourceSpecsShrink
}

func (s *ConvertInstanceShrinkRequest) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *ConvertInstanceShrinkRequest) GetPromotionCode() *string {
	return s.PromotionCode
}

func (s *ConvertInstanceShrinkRequest) GetRegion() *string {
	return s.Region
}

func (s *ConvertInstanceShrinkRequest) GetUsePromotionCode() *bool {
	return s.UsePromotionCode
}

func (s *ConvertInstanceShrinkRequest) SetDuration(v int32) *ConvertInstanceShrinkRequest {
	s.Duration = &v
	return s
}

func (s *ConvertInstanceShrinkRequest) SetInstanceId(v string) *ConvertInstanceShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *ConvertInstanceShrinkRequest) SetIsAutoRenew(v bool) *ConvertInstanceShrinkRequest {
	s.IsAutoRenew = &v
	return s
}

func (s *ConvertInstanceShrinkRequest) SetNamespaceResourceSpecsShrink(v string) *ConvertInstanceShrinkRequest {
	s.NamespaceResourceSpecsShrink = &v
	return s
}

func (s *ConvertInstanceShrinkRequest) SetPricingCycle(v string) *ConvertInstanceShrinkRequest {
	s.PricingCycle = &v
	return s
}

func (s *ConvertInstanceShrinkRequest) SetPromotionCode(v string) *ConvertInstanceShrinkRequest {
	s.PromotionCode = &v
	return s
}

func (s *ConvertInstanceShrinkRequest) SetRegion(v string) *ConvertInstanceShrinkRequest {
	s.Region = &v
	return s
}

func (s *ConvertInstanceShrinkRequest) SetUsePromotionCode(v bool) *ConvertInstanceShrinkRequest {
	s.UsePromotionCode = &v
	return s
}

func (s *ConvertInstanceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
