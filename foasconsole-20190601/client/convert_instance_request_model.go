// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConvertInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConvertPostpayInstanceRequest(v *ConvertInstanceRequestConvertPostpayInstanceRequest) *ConvertInstanceRequest
	GetConvertPostpayInstanceRequest() *ConvertInstanceRequestConvertPostpayInstanceRequest
}

type ConvertInstanceRequest struct {
	// This parameter is required.
	ConvertPostpayInstanceRequest *ConvertInstanceRequestConvertPostpayInstanceRequest `json:"ConvertPostpayInstanceRequest,omitempty" xml:"ConvertPostpayInstanceRequest,omitempty" type:"Struct"`
}

func (s ConvertInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ConvertInstanceRequest) GoString() string {
	return s.String()
}

func (s *ConvertInstanceRequest) GetConvertPostpayInstanceRequest() *ConvertInstanceRequestConvertPostpayInstanceRequest {
	return s.ConvertPostpayInstanceRequest
}

func (s *ConvertInstanceRequest) SetConvertPostpayInstanceRequest(v *ConvertInstanceRequestConvertPostpayInstanceRequest) *ConvertInstanceRequest {
	s.ConvertPostpayInstanceRequest = v
	return s
}

func (s *ConvertInstanceRequest) Validate() error {
	if s.ConvertPostpayInstanceRequest != nil {
		if err := s.ConvertPostpayInstanceRequest.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ConvertInstanceRequestConvertPostpayInstanceRequest struct {
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
	// sc_flinkserverlesspost_public_cn-*******
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	IsAutoRenew *bool `json:"IsAutoRenew,omitempty" xml:"IsAutoRenew,omitempty"`
	// This parameter is required.
	NamespaceResourceSpecs []*ConvertInstanceRequestConvertPostpayInstanceRequestNamespaceResourceSpecs `json:"NamespaceResourceSpecs,omitempty" xml:"NamespaceResourceSpecs,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// Month
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s ConvertInstanceRequestConvertPostpayInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ConvertInstanceRequestConvertPostpayInstanceRequest) GoString() string {
	return s.String()
}

func (s *ConvertInstanceRequestConvertPostpayInstanceRequest) GetDuration() *int32 {
	return s.Duration
}

func (s *ConvertInstanceRequestConvertPostpayInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ConvertInstanceRequestConvertPostpayInstanceRequest) GetIsAutoRenew() *bool {
	return s.IsAutoRenew
}

func (s *ConvertInstanceRequestConvertPostpayInstanceRequest) GetNamespaceResourceSpecs() []*ConvertInstanceRequestConvertPostpayInstanceRequestNamespaceResourceSpecs {
	return s.NamespaceResourceSpecs
}

func (s *ConvertInstanceRequestConvertPostpayInstanceRequest) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *ConvertInstanceRequestConvertPostpayInstanceRequest) GetRegion() *string {
	return s.Region
}

func (s *ConvertInstanceRequestConvertPostpayInstanceRequest) SetDuration(v int32) *ConvertInstanceRequestConvertPostpayInstanceRequest {
	s.Duration = &v
	return s
}

func (s *ConvertInstanceRequestConvertPostpayInstanceRequest) SetInstanceId(v string) *ConvertInstanceRequestConvertPostpayInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *ConvertInstanceRequestConvertPostpayInstanceRequest) SetIsAutoRenew(v bool) *ConvertInstanceRequestConvertPostpayInstanceRequest {
	s.IsAutoRenew = &v
	return s
}

func (s *ConvertInstanceRequestConvertPostpayInstanceRequest) SetNamespaceResourceSpecs(v []*ConvertInstanceRequestConvertPostpayInstanceRequestNamespaceResourceSpecs) *ConvertInstanceRequestConvertPostpayInstanceRequest {
	s.NamespaceResourceSpecs = v
	return s
}

func (s *ConvertInstanceRequestConvertPostpayInstanceRequest) SetPricingCycle(v string) *ConvertInstanceRequestConvertPostpayInstanceRequest {
	s.PricingCycle = &v
	return s
}

func (s *ConvertInstanceRequestConvertPostpayInstanceRequest) SetRegion(v string) *ConvertInstanceRequestConvertPostpayInstanceRequest {
	s.Region = &v
	return s
}

func (s *ConvertInstanceRequestConvertPostpayInstanceRequest) Validate() error {
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

type ConvertInstanceRequestConvertPostpayInstanceRequestNamespaceResourceSpecs struct {
	// This parameter is required.
	//
	// example:
	//
	// ns-1
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// This parameter is required.
	ResourceSpec *ConvertInstanceRequestConvertPostpayInstanceRequestNamespaceResourceSpecsResourceSpec `json:"ResourceSpec,omitempty" xml:"ResourceSpec,omitempty" type:"Struct"`
}

func (s ConvertInstanceRequestConvertPostpayInstanceRequestNamespaceResourceSpecs) String() string {
	return dara.Prettify(s)
}

func (s ConvertInstanceRequestConvertPostpayInstanceRequestNamespaceResourceSpecs) GoString() string {
	return s.String()
}

func (s *ConvertInstanceRequestConvertPostpayInstanceRequestNamespaceResourceSpecs) GetNamespace() *string {
	return s.Namespace
}

func (s *ConvertInstanceRequestConvertPostpayInstanceRequestNamespaceResourceSpecs) GetResourceSpec() *ConvertInstanceRequestConvertPostpayInstanceRequestNamespaceResourceSpecsResourceSpec {
	return s.ResourceSpec
}

func (s *ConvertInstanceRequestConvertPostpayInstanceRequestNamespaceResourceSpecs) SetNamespace(v string) *ConvertInstanceRequestConvertPostpayInstanceRequestNamespaceResourceSpecs {
	s.Namespace = &v
	return s
}

func (s *ConvertInstanceRequestConvertPostpayInstanceRequestNamespaceResourceSpecs) SetResourceSpec(v *ConvertInstanceRequestConvertPostpayInstanceRequestNamespaceResourceSpecsResourceSpec) *ConvertInstanceRequestConvertPostpayInstanceRequestNamespaceResourceSpecs {
	s.ResourceSpec = v
	return s
}

func (s *ConvertInstanceRequestConvertPostpayInstanceRequestNamespaceResourceSpecs) Validate() error {
	if s.ResourceSpec != nil {
		if err := s.ResourceSpec.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ConvertInstanceRequestConvertPostpayInstanceRequestNamespaceResourceSpecsResourceSpec struct {
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

func (s ConvertInstanceRequestConvertPostpayInstanceRequestNamespaceResourceSpecsResourceSpec) String() string {
	return dara.Prettify(s)
}

func (s ConvertInstanceRequestConvertPostpayInstanceRequestNamespaceResourceSpecsResourceSpec) GoString() string {
	return s.String()
}

func (s *ConvertInstanceRequestConvertPostpayInstanceRequestNamespaceResourceSpecsResourceSpec) GetCpu() *int32 {
	return s.Cpu
}

func (s *ConvertInstanceRequestConvertPostpayInstanceRequestNamespaceResourceSpecsResourceSpec) GetMemoryGB() *int32 {
	return s.MemoryGB
}

func (s *ConvertInstanceRequestConvertPostpayInstanceRequestNamespaceResourceSpecsResourceSpec) SetCpu(v int32) *ConvertInstanceRequestConvertPostpayInstanceRequestNamespaceResourceSpecsResourceSpec {
	s.Cpu = &v
	return s
}

func (s *ConvertInstanceRequestConvertPostpayInstanceRequestNamespaceResourceSpecsResourceSpec) SetMemoryGB(v int32) *ConvertInstanceRequestConvertPostpayInstanceRequestNamespaceResourceSpecsResourceSpec {
	s.MemoryGB = &v
	return s
}

func (s *ConvertInstanceRequestConvertPostpayInstanceRequestNamespaceResourceSpecsResourceSpec) Validate() error {
	return dara.Validate(s)
}
