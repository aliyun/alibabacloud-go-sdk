// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryConvertInstancePriceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConvertPostpayInstanceRequest(v *QueryConvertInstancePriceRequestConvertPostpayInstanceRequest) *QueryConvertInstancePriceRequest
	GetConvertPostpayInstanceRequest() *QueryConvertInstancePriceRequestConvertPostpayInstanceRequest
}

type QueryConvertInstancePriceRequest struct {
	// This parameter is required.
	ConvertPostpayInstanceRequest *QueryConvertInstancePriceRequestConvertPostpayInstanceRequest `json:"ConvertPostpayInstanceRequest,omitempty" xml:"ConvertPostpayInstanceRequest,omitempty" type:"Struct"`
}

func (s QueryConvertInstancePriceRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryConvertInstancePriceRequest) GoString() string {
	return s.String()
}

func (s *QueryConvertInstancePriceRequest) GetConvertPostpayInstanceRequest() *QueryConvertInstancePriceRequestConvertPostpayInstanceRequest {
	return s.ConvertPostpayInstanceRequest
}

func (s *QueryConvertInstancePriceRequest) SetConvertPostpayInstanceRequest(v *QueryConvertInstancePriceRequestConvertPostpayInstanceRequest) *QueryConvertInstancePriceRequest {
	s.ConvertPostpayInstanceRequest = v
	return s
}

func (s *QueryConvertInstancePriceRequest) Validate() error {
	if s.ConvertPostpayInstanceRequest != nil {
		if err := s.ConvertPostpayInstanceRequest.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryConvertInstancePriceRequestConvertPostpayInstanceRequest struct {
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
	// sc_flinkserverlesspost_public_cn-******
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	IsAutoRenew *bool `json:"IsAutoRenew,omitempty" xml:"IsAutoRenew,omitempty"`
	// This parameter is required.
	NamespaceResourceSpecs []*QueryConvertInstancePriceRequestConvertPostpayInstanceRequestNamespaceResourceSpecs `json:"NamespaceResourceSpecs,omitempty" xml:"NamespaceResourceSpecs,omitempty" type:"Repeated"`
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

func (s QueryConvertInstancePriceRequestConvertPostpayInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryConvertInstancePriceRequestConvertPostpayInstanceRequest) GoString() string {
	return s.String()
}

func (s *QueryConvertInstancePriceRequestConvertPostpayInstanceRequest) GetDuration() *int32 {
	return s.Duration
}

func (s *QueryConvertInstancePriceRequestConvertPostpayInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryConvertInstancePriceRequestConvertPostpayInstanceRequest) GetIsAutoRenew() *bool {
	return s.IsAutoRenew
}

func (s *QueryConvertInstancePriceRequestConvertPostpayInstanceRequest) GetNamespaceResourceSpecs() []*QueryConvertInstancePriceRequestConvertPostpayInstanceRequestNamespaceResourceSpecs {
	return s.NamespaceResourceSpecs
}

func (s *QueryConvertInstancePriceRequestConvertPostpayInstanceRequest) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *QueryConvertInstancePriceRequestConvertPostpayInstanceRequest) GetRegion() *string {
	return s.Region
}

func (s *QueryConvertInstancePriceRequestConvertPostpayInstanceRequest) SetDuration(v int32) *QueryConvertInstancePriceRequestConvertPostpayInstanceRequest {
	s.Duration = &v
	return s
}

func (s *QueryConvertInstancePriceRequestConvertPostpayInstanceRequest) SetInstanceId(v string) *QueryConvertInstancePriceRequestConvertPostpayInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryConvertInstancePriceRequestConvertPostpayInstanceRequest) SetIsAutoRenew(v bool) *QueryConvertInstancePriceRequestConvertPostpayInstanceRequest {
	s.IsAutoRenew = &v
	return s
}

func (s *QueryConvertInstancePriceRequestConvertPostpayInstanceRequest) SetNamespaceResourceSpecs(v []*QueryConvertInstancePriceRequestConvertPostpayInstanceRequestNamespaceResourceSpecs) *QueryConvertInstancePriceRequestConvertPostpayInstanceRequest {
	s.NamespaceResourceSpecs = v
	return s
}

func (s *QueryConvertInstancePriceRequestConvertPostpayInstanceRequest) SetPricingCycle(v string) *QueryConvertInstancePriceRequestConvertPostpayInstanceRequest {
	s.PricingCycle = &v
	return s
}

func (s *QueryConvertInstancePriceRequestConvertPostpayInstanceRequest) SetRegion(v string) *QueryConvertInstancePriceRequestConvertPostpayInstanceRequest {
	s.Region = &v
	return s
}

func (s *QueryConvertInstancePriceRequestConvertPostpayInstanceRequest) Validate() error {
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

type QueryConvertInstancePriceRequestConvertPostpayInstanceRequestNamespaceResourceSpecs struct {
	// This parameter is required.
	//
	// example:
	//
	// ns-1
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// This parameter is required.
	ResourceSpec *QueryConvertInstancePriceRequestConvertPostpayInstanceRequestNamespaceResourceSpecsResourceSpec `json:"ResourceSpec,omitempty" xml:"ResourceSpec,omitempty" type:"Struct"`
}

func (s QueryConvertInstancePriceRequestConvertPostpayInstanceRequestNamespaceResourceSpecs) String() string {
	return dara.Prettify(s)
}

func (s QueryConvertInstancePriceRequestConvertPostpayInstanceRequestNamespaceResourceSpecs) GoString() string {
	return s.String()
}

func (s *QueryConvertInstancePriceRequestConvertPostpayInstanceRequestNamespaceResourceSpecs) GetNamespace() *string {
	return s.Namespace
}

func (s *QueryConvertInstancePriceRequestConvertPostpayInstanceRequestNamespaceResourceSpecs) GetResourceSpec() *QueryConvertInstancePriceRequestConvertPostpayInstanceRequestNamespaceResourceSpecsResourceSpec {
	return s.ResourceSpec
}

func (s *QueryConvertInstancePriceRequestConvertPostpayInstanceRequestNamespaceResourceSpecs) SetNamespace(v string) *QueryConvertInstancePriceRequestConvertPostpayInstanceRequestNamespaceResourceSpecs {
	s.Namespace = &v
	return s
}

func (s *QueryConvertInstancePriceRequestConvertPostpayInstanceRequestNamespaceResourceSpecs) SetResourceSpec(v *QueryConvertInstancePriceRequestConvertPostpayInstanceRequestNamespaceResourceSpecsResourceSpec) *QueryConvertInstancePriceRequestConvertPostpayInstanceRequestNamespaceResourceSpecs {
	s.ResourceSpec = v
	return s
}

func (s *QueryConvertInstancePriceRequestConvertPostpayInstanceRequestNamespaceResourceSpecs) Validate() error {
	if s.ResourceSpec != nil {
		if err := s.ResourceSpec.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryConvertInstancePriceRequestConvertPostpayInstanceRequestNamespaceResourceSpecsResourceSpec struct {
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

func (s QueryConvertInstancePriceRequestConvertPostpayInstanceRequestNamespaceResourceSpecsResourceSpec) String() string {
	return dara.Prettify(s)
}

func (s QueryConvertInstancePriceRequestConvertPostpayInstanceRequestNamespaceResourceSpecsResourceSpec) GoString() string {
	return s.String()
}

func (s *QueryConvertInstancePriceRequestConvertPostpayInstanceRequestNamespaceResourceSpecsResourceSpec) GetCpu() *int32 {
	return s.Cpu
}

func (s *QueryConvertInstancePriceRequestConvertPostpayInstanceRequestNamespaceResourceSpecsResourceSpec) GetMemoryGB() *int32 {
	return s.MemoryGB
}

func (s *QueryConvertInstancePriceRequestConvertPostpayInstanceRequestNamespaceResourceSpecsResourceSpec) SetCpu(v int32) *QueryConvertInstancePriceRequestConvertPostpayInstanceRequestNamespaceResourceSpecsResourceSpec {
	s.Cpu = &v
	return s
}

func (s *QueryConvertInstancePriceRequestConvertPostpayInstanceRequestNamespaceResourceSpecsResourceSpec) SetMemoryGB(v int32) *QueryConvertInstancePriceRequestConvertPostpayInstanceRequestNamespaceResourceSpecsResourceSpec {
	s.MemoryGB = &v
	return s
}

func (s *QueryConvertInstancePriceRequestConvertPostpayInstanceRequestNamespaceResourceSpecsResourceSpec) Validate() error {
	return dara.Validate(s)
}
