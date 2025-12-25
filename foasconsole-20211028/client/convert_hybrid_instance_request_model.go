// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConvertHybridInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ConvertHybridInstanceRequest
	GetInstanceId() *string
	SetRegion(v string) *ConvertHybridInstanceRequest
	GetRegion() *string
	SetResourceSpec(v *ConvertHybridInstanceRequestResourceSpec) *ConvertHybridInstanceRequest
	GetResourceSpec() *ConvertHybridInstanceRequestResourceSpec
}

type ConvertHybridInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// sc_flinkserverless_public_cn-7e22ae5sess
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-shenzhen
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// This parameter is required.
	ResourceSpec *ConvertHybridInstanceRequestResourceSpec `json:"ResourceSpec,omitempty" xml:"ResourceSpec,omitempty" type:"Struct"`
}

func (s ConvertHybridInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ConvertHybridInstanceRequest) GoString() string {
	return s.String()
}

func (s *ConvertHybridInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ConvertHybridInstanceRequest) GetRegion() *string {
	return s.Region
}

func (s *ConvertHybridInstanceRequest) GetResourceSpec() *ConvertHybridInstanceRequestResourceSpec {
	return s.ResourceSpec
}

func (s *ConvertHybridInstanceRequest) SetInstanceId(v string) *ConvertHybridInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *ConvertHybridInstanceRequest) SetRegion(v string) *ConvertHybridInstanceRequest {
	s.Region = &v
	return s
}

func (s *ConvertHybridInstanceRequest) SetResourceSpec(v *ConvertHybridInstanceRequestResourceSpec) *ConvertHybridInstanceRequest {
	s.ResourceSpec = v
	return s
}

func (s *ConvertHybridInstanceRequest) Validate() error {
	if s.ResourceSpec != nil {
		if err := s.ResourceSpec.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ConvertHybridInstanceRequestResourceSpec struct {
	// This parameter is required.
	//
	// example:
	//
	// 10
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 40GB
	MemoryGB *int32 `json:"MemoryGB,omitempty" xml:"MemoryGB,omitempty"`
}

func (s ConvertHybridInstanceRequestResourceSpec) String() string {
	return dara.Prettify(s)
}

func (s ConvertHybridInstanceRequestResourceSpec) GoString() string {
	return s.String()
}

func (s *ConvertHybridInstanceRequestResourceSpec) GetCpu() *int32 {
	return s.Cpu
}

func (s *ConvertHybridInstanceRequestResourceSpec) GetMemoryGB() *int32 {
	return s.MemoryGB
}

func (s *ConvertHybridInstanceRequestResourceSpec) SetCpu(v int32) *ConvertHybridInstanceRequestResourceSpec {
	s.Cpu = &v
	return s
}

func (s *ConvertHybridInstanceRequestResourceSpec) SetMemoryGB(v int32) *ConvertHybridInstanceRequestResourceSpec {
	s.MemoryGB = &v
	return s
}

func (s *ConvertHybridInstanceRequestResourceSpec) Validate() error {
	return dara.Validate(s)
}
