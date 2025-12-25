// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryConvertInstancePriceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDuration(v int32) *QueryConvertInstancePriceRequest
	GetDuration() *int32
	SetInstanceId(v string) *QueryConvertInstancePriceRequest
	GetInstanceId() *string
	SetIsAutoRenew(v bool) *QueryConvertInstancePriceRequest
	GetIsAutoRenew() *bool
	SetNamespaceResourceSpecs(v []*QueryConvertInstancePriceRequestNamespaceResourceSpecs) *QueryConvertInstancePriceRequest
	GetNamespaceResourceSpecs() []*QueryConvertInstancePriceRequestNamespaceResourceSpecs
	SetPricingCycle(v string) *QueryConvertInstancePriceRequest
	GetPricingCycle() *string
	SetPromotionCode(v string) *QueryConvertInstancePriceRequest
	GetPromotionCode() *string
	SetRegion(v string) *QueryConvertInstancePriceRequest
	GetRegion() *string
	SetUsePromotionCode(v bool) *QueryConvertInstancePriceRequest
	GetUsePromotionCode() *bool
}

type QueryConvertInstancePriceRequest struct {
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
	NamespaceResourceSpecs []*QueryConvertInstancePriceRequestNamespaceResourceSpecs `json:"NamespaceResourceSpecs,omitempty" xml:"NamespaceResourceSpecs,omitempty" type:"Repeated"`
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

func (s QueryConvertInstancePriceRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryConvertInstancePriceRequest) GoString() string {
	return s.String()
}

func (s *QueryConvertInstancePriceRequest) GetDuration() *int32 {
	return s.Duration
}

func (s *QueryConvertInstancePriceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryConvertInstancePriceRequest) GetIsAutoRenew() *bool {
	return s.IsAutoRenew
}

func (s *QueryConvertInstancePriceRequest) GetNamespaceResourceSpecs() []*QueryConvertInstancePriceRequestNamespaceResourceSpecs {
	return s.NamespaceResourceSpecs
}

func (s *QueryConvertInstancePriceRequest) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *QueryConvertInstancePriceRequest) GetPromotionCode() *string {
	return s.PromotionCode
}

func (s *QueryConvertInstancePriceRequest) GetRegion() *string {
	return s.Region
}

func (s *QueryConvertInstancePriceRequest) GetUsePromotionCode() *bool {
	return s.UsePromotionCode
}

func (s *QueryConvertInstancePriceRequest) SetDuration(v int32) *QueryConvertInstancePriceRequest {
	s.Duration = &v
	return s
}

func (s *QueryConvertInstancePriceRequest) SetInstanceId(v string) *QueryConvertInstancePriceRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryConvertInstancePriceRequest) SetIsAutoRenew(v bool) *QueryConvertInstancePriceRequest {
	s.IsAutoRenew = &v
	return s
}

func (s *QueryConvertInstancePriceRequest) SetNamespaceResourceSpecs(v []*QueryConvertInstancePriceRequestNamespaceResourceSpecs) *QueryConvertInstancePriceRequest {
	s.NamespaceResourceSpecs = v
	return s
}

func (s *QueryConvertInstancePriceRequest) SetPricingCycle(v string) *QueryConvertInstancePriceRequest {
	s.PricingCycle = &v
	return s
}

func (s *QueryConvertInstancePriceRequest) SetPromotionCode(v string) *QueryConvertInstancePriceRequest {
	s.PromotionCode = &v
	return s
}

func (s *QueryConvertInstancePriceRequest) SetRegion(v string) *QueryConvertInstancePriceRequest {
	s.Region = &v
	return s
}

func (s *QueryConvertInstancePriceRequest) SetUsePromotionCode(v bool) *QueryConvertInstancePriceRequest {
	s.UsePromotionCode = &v
	return s
}

func (s *QueryConvertInstancePriceRequest) Validate() error {
	if s.NamespaceResourceSpecs != nil {
		for _, item := range s.NamespaceResourceSpecs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryConvertInstancePriceRequestNamespaceResourceSpecs struct {
	// This parameter is required.
	//
	// example:
	//
	// lm-test-default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// This parameter is required.
	ResourceSpec *QueryConvertInstancePriceRequestNamespaceResourceSpecsResourceSpec `json:"ResourceSpec,omitempty" xml:"ResourceSpec,omitempty" type:"Struct"`
}

func (s QueryConvertInstancePriceRequestNamespaceResourceSpecs) String() string {
	return dara.Prettify(s)
}

func (s QueryConvertInstancePriceRequestNamespaceResourceSpecs) GoString() string {
	return s.String()
}

func (s *QueryConvertInstancePriceRequestNamespaceResourceSpecs) GetNamespace() *string {
	return s.Namespace
}

func (s *QueryConvertInstancePriceRequestNamespaceResourceSpecs) GetResourceSpec() *QueryConvertInstancePriceRequestNamespaceResourceSpecsResourceSpec {
	return s.ResourceSpec
}

func (s *QueryConvertInstancePriceRequestNamespaceResourceSpecs) SetNamespace(v string) *QueryConvertInstancePriceRequestNamespaceResourceSpecs {
	s.Namespace = &v
	return s
}

func (s *QueryConvertInstancePriceRequestNamespaceResourceSpecs) SetResourceSpec(v *QueryConvertInstancePriceRequestNamespaceResourceSpecsResourceSpec) *QueryConvertInstancePriceRequestNamespaceResourceSpecs {
	s.ResourceSpec = v
	return s
}

func (s *QueryConvertInstancePriceRequestNamespaceResourceSpecs) Validate() error {
	if s.ResourceSpec != nil {
		if err := s.ResourceSpec.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryConvertInstancePriceRequestNamespaceResourceSpecsResourceSpec struct {
	// This parameter is required.
	//
	// example:
	//
	// 6
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 24
	MemoryGB *int32 `json:"MemoryGB,omitempty" xml:"MemoryGB,omitempty"`
}

func (s QueryConvertInstancePriceRequestNamespaceResourceSpecsResourceSpec) String() string {
	return dara.Prettify(s)
}

func (s QueryConvertInstancePriceRequestNamespaceResourceSpecsResourceSpec) GoString() string {
	return s.String()
}

func (s *QueryConvertInstancePriceRequestNamespaceResourceSpecsResourceSpec) GetCpu() *int32 {
	return s.Cpu
}

func (s *QueryConvertInstancePriceRequestNamespaceResourceSpecsResourceSpec) GetMemoryGB() *int32 {
	return s.MemoryGB
}

func (s *QueryConvertInstancePriceRequestNamespaceResourceSpecsResourceSpec) SetCpu(v int32) *QueryConvertInstancePriceRequestNamespaceResourceSpecsResourceSpec {
	s.Cpu = &v
	return s
}

func (s *QueryConvertInstancePriceRequestNamespaceResourceSpecsResourceSpec) SetMemoryGB(v int32) *QueryConvertInstancePriceRequestNamespaceResourceSpecsResourceSpec {
	s.MemoryGB = &v
	return s
}

func (s *QueryConvertInstancePriceRequestNamespaceResourceSpecsResourceSpec) Validate() error {
	return dara.Validate(s)
}
