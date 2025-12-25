// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConvertInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDuration(v int32) *ConvertInstanceRequest
	GetDuration() *int32
	SetInstanceId(v string) *ConvertInstanceRequest
	GetInstanceId() *string
	SetIsAutoRenew(v bool) *ConvertInstanceRequest
	GetIsAutoRenew() *bool
	SetNamespaceResourceSpecs(v []*ConvertInstanceRequestNamespaceResourceSpecs) *ConvertInstanceRequest
	GetNamespaceResourceSpecs() []*ConvertInstanceRequestNamespaceResourceSpecs
	SetPricingCycle(v string) *ConvertInstanceRequest
	GetPricingCycle() *string
	SetPromotionCode(v string) *ConvertInstanceRequest
	GetPromotionCode() *string
	SetRegion(v string) *ConvertInstanceRequest
	GetRegion() *string
	SetUsePromotionCode(v bool) *ConvertInstanceRequest
	GetUsePromotionCode() *bool
}

type ConvertInstanceRequest struct {
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
	NamespaceResourceSpecs []*ConvertInstanceRequestNamespaceResourceSpecs `json:"NamespaceResourceSpecs,omitempty" xml:"NamespaceResourceSpecs,omitempty" type:"Repeated"`
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

func (s ConvertInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ConvertInstanceRequest) GoString() string {
	return s.String()
}

func (s *ConvertInstanceRequest) GetDuration() *int32 {
	return s.Duration
}

func (s *ConvertInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ConvertInstanceRequest) GetIsAutoRenew() *bool {
	return s.IsAutoRenew
}

func (s *ConvertInstanceRequest) GetNamespaceResourceSpecs() []*ConvertInstanceRequestNamespaceResourceSpecs {
	return s.NamespaceResourceSpecs
}

func (s *ConvertInstanceRequest) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *ConvertInstanceRequest) GetPromotionCode() *string {
	return s.PromotionCode
}

func (s *ConvertInstanceRequest) GetRegion() *string {
	return s.Region
}

func (s *ConvertInstanceRequest) GetUsePromotionCode() *bool {
	return s.UsePromotionCode
}

func (s *ConvertInstanceRequest) SetDuration(v int32) *ConvertInstanceRequest {
	s.Duration = &v
	return s
}

func (s *ConvertInstanceRequest) SetInstanceId(v string) *ConvertInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *ConvertInstanceRequest) SetIsAutoRenew(v bool) *ConvertInstanceRequest {
	s.IsAutoRenew = &v
	return s
}

func (s *ConvertInstanceRequest) SetNamespaceResourceSpecs(v []*ConvertInstanceRequestNamespaceResourceSpecs) *ConvertInstanceRequest {
	s.NamespaceResourceSpecs = v
	return s
}

func (s *ConvertInstanceRequest) SetPricingCycle(v string) *ConvertInstanceRequest {
	s.PricingCycle = &v
	return s
}

func (s *ConvertInstanceRequest) SetPromotionCode(v string) *ConvertInstanceRequest {
	s.PromotionCode = &v
	return s
}

func (s *ConvertInstanceRequest) SetRegion(v string) *ConvertInstanceRequest {
	s.Region = &v
	return s
}

func (s *ConvertInstanceRequest) SetUsePromotionCode(v bool) *ConvertInstanceRequest {
	s.UsePromotionCode = &v
	return s
}

func (s *ConvertInstanceRequest) Validate() error {
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

type ConvertInstanceRequestNamespaceResourceSpecs struct {
	// This parameter is required.
	//
	// example:
	//
	// ns-1
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// This parameter is required.
	ResourceSpec *ConvertInstanceRequestNamespaceResourceSpecsResourceSpec `json:"ResourceSpec,omitempty" xml:"ResourceSpec,omitempty" type:"Struct"`
}

func (s ConvertInstanceRequestNamespaceResourceSpecs) String() string {
	return dara.Prettify(s)
}

func (s ConvertInstanceRequestNamespaceResourceSpecs) GoString() string {
	return s.String()
}

func (s *ConvertInstanceRequestNamespaceResourceSpecs) GetNamespace() *string {
	return s.Namespace
}

func (s *ConvertInstanceRequestNamespaceResourceSpecs) GetResourceSpec() *ConvertInstanceRequestNamespaceResourceSpecsResourceSpec {
	return s.ResourceSpec
}

func (s *ConvertInstanceRequestNamespaceResourceSpecs) SetNamespace(v string) *ConvertInstanceRequestNamespaceResourceSpecs {
	s.Namespace = &v
	return s
}

func (s *ConvertInstanceRequestNamespaceResourceSpecs) SetResourceSpec(v *ConvertInstanceRequestNamespaceResourceSpecsResourceSpec) *ConvertInstanceRequestNamespaceResourceSpecs {
	s.ResourceSpec = v
	return s
}

func (s *ConvertInstanceRequestNamespaceResourceSpecs) Validate() error {
	if s.ResourceSpec != nil {
		if err := s.ResourceSpec.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ConvertInstanceRequestNamespaceResourceSpecsResourceSpec struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 4
	MemoryGB *int32 `json:"MemoryGB,omitempty" xml:"MemoryGB,omitempty"`
}

func (s ConvertInstanceRequestNamespaceResourceSpecsResourceSpec) String() string {
	return dara.Prettify(s)
}

func (s ConvertInstanceRequestNamespaceResourceSpecsResourceSpec) GoString() string {
	return s.String()
}

func (s *ConvertInstanceRequestNamespaceResourceSpecsResourceSpec) GetCpu() *int32 {
	return s.Cpu
}

func (s *ConvertInstanceRequestNamespaceResourceSpecsResourceSpec) GetMemoryGB() *int32 {
	return s.MemoryGB
}

func (s *ConvertInstanceRequestNamespaceResourceSpecsResourceSpec) SetCpu(v int32) *ConvertInstanceRequestNamespaceResourceSpecsResourceSpec {
	s.Cpu = &v
	return s
}

func (s *ConvertInstanceRequestNamespaceResourceSpecsResourceSpec) SetMemoryGB(v int32) *ConvertInstanceRequestNamespaceResourceSpecsResourceSpec {
	s.MemoryGB = &v
	return s
}

func (s *ConvertInstanceRequestNamespaceResourceSpecsResourceSpec) Validate() error {
	return dara.Validate(s)
}
