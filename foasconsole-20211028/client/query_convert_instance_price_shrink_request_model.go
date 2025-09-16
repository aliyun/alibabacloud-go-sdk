// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryConvertInstancePriceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDuration(v int32) *QueryConvertInstancePriceShrinkRequest
	GetDuration() *int32
	SetInstanceId(v string) *QueryConvertInstancePriceShrinkRequest
	GetInstanceId() *string
	SetIsAutoRenew(v bool) *QueryConvertInstancePriceShrinkRequest
	GetIsAutoRenew() *bool
	SetNamespaceResourceSpecsShrink(v string) *QueryConvertInstancePriceShrinkRequest
	GetNamespaceResourceSpecsShrink() *string
	SetPricingCycle(v string) *QueryConvertInstancePriceShrinkRequest
	GetPricingCycle() *string
	SetPromotionCode(v string) *QueryConvertInstancePriceShrinkRequest
	GetPromotionCode() *string
	SetRegion(v string) *QueryConvertInstancePriceShrinkRequest
	GetRegion() *string
	SetUsePromotionCode(v bool) *QueryConvertInstancePriceShrinkRequest
	GetUsePromotionCode() *bool
}

type QueryConvertInstancePriceShrinkRequest struct {
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

func (s QueryConvertInstancePriceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryConvertInstancePriceShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryConvertInstancePriceShrinkRequest) GetDuration() *int32 {
	return s.Duration
}

func (s *QueryConvertInstancePriceShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryConvertInstancePriceShrinkRequest) GetIsAutoRenew() *bool {
	return s.IsAutoRenew
}

func (s *QueryConvertInstancePriceShrinkRequest) GetNamespaceResourceSpecsShrink() *string {
	return s.NamespaceResourceSpecsShrink
}

func (s *QueryConvertInstancePriceShrinkRequest) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *QueryConvertInstancePriceShrinkRequest) GetPromotionCode() *string {
	return s.PromotionCode
}

func (s *QueryConvertInstancePriceShrinkRequest) GetRegion() *string {
	return s.Region
}

func (s *QueryConvertInstancePriceShrinkRequest) GetUsePromotionCode() *bool {
	return s.UsePromotionCode
}

func (s *QueryConvertInstancePriceShrinkRequest) SetDuration(v int32) *QueryConvertInstancePriceShrinkRequest {
	s.Duration = &v
	return s
}

func (s *QueryConvertInstancePriceShrinkRequest) SetInstanceId(v string) *QueryConvertInstancePriceShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryConvertInstancePriceShrinkRequest) SetIsAutoRenew(v bool) *QueryConvertInstancePriceShrinkRequest {
	s.IsAutoRenew = &v
	return s
}

func (s *QueryConvertInstancePriceShrinkRequest) SetNamespaceResourceSpecsShrink(v string) *QueryConvertInstancePriceShrinkRequest {
	s.NamespaceResourceSpecsShrink = &v
	return s
}

func (s *QueryConvertInstancePriceShrinkRequest) SetPricingCycle(v string) *QueryConvertInstancePriceShrinkRequest {
	s.PricingCycle = &v
	return s
}

func (s *QueryConvertInstancePriceShrinkRequest) SetPromotionCode(v string) *QueryConvertInstancePriceShrinkRequest {
	s.PromotionCode = &v
	return s
}

func (s *QueryConvertInstancePriceShrinkRequest) SetRegion(v string) *QueryConvertInstancePriceShrinkRequest {
	s.Region = &v
	return s
}

func (s *QueryConvertInstancePriceShrinkRequest) SetUsePromotionCode(v bool) *QueryConvertInstancePriceShrinkRequest {
	s.UsePromotionCode = &v
	return s
}

func (s *QueryConvertInstancePriceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
