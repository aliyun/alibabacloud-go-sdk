// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyElasticResourceSpecRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyElasticResourceSpecRequest
	GetInstanceId() *string
	SetRegion(v string) *ModifyElasticResourceSpecRequest
	GetRegion() *string
	SetResourceSpec(v *ModifyElasticResourceSpecRequestResourceSpec) *ModifyElasticResourceSpecRequest
	GetResourceSpec() *ModifyElasticResourceSpecRequestResourceSpec
}

type ModifyElasticResourceSpecRequest struct {
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
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// This parameter is required.
	ResourceSpec *ModifyElasticResourceSpecRequestResourceSpec `json:"ResourceSpec,omitempty" xml:"ResourceSpec,omitempty" type:"Struct"`
}

func (s ModifyElasticResourceSpecRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyElasticResourceSpecRequest) GoString() string {
	return s.String()
}

func (s *ModifyElasticResourceSpecRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyElasticResourceSpecRequest) GetRegion() *string {
	return s.Region
}

func (s *ModifyElasticResourceSpecRequest) GetResourceSpec() *ModifyElasticResourceSpecRequestResourceSpec {
	return s.ResourceSpec
}

func (s *ModifyElasticResourceSpecRequest) SetInstanceId(v string) *ModifyElasticResourceSpecRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyElasticResourceSpecRequest) SetRegion(v string) *ModifyElasticResourceSpecRequest {
	s.Region = &v
	return s
}

func (s *ModifyElasticResourceSpecRequest) SetResourceSpec(v *ModifyElasticResourceSpecRequestResourceSpec) *ModifyElasticResourceSpecRequest {
	s.ResourceSpec = v
	return s
}

func (s *ModifyElasticResourceSpecRequest) Validate() error {
	if s.ResourceSpec != nil {
		if err := s.ResourceSpec.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyElasticResourceSpecRequestResourceSpec struct {
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
	// 40
	MemoryGB *int32 `json:"MemoryGB,omitempty" xml:"MemoryGB,omitempty"`
}

func (s ModifyElasticResourceSpecRequestResourceSpec) String() string {
	return dara.Prettify(s)
}

func (s ModifyElasticResourceSpecRequestResourceSpec) GoString() string {
	return s.String()
}

func (s *ModifyElasticResourceSpecRequestResourceSpec) GetCpu() *int32 {
	return s.Cpu
}

func (s *ModifyElasticResourceSpecRequestResourceSpec) GetMemoryGB() *int32 {
	return s.MemoryGB
}

func (s *ModifyElasticResourceSpecRequestResourceSpec) SetCpu(v int32) *ModifyElasticResourceSpecRequestResourceSpec {
	s.Cpu = &v
	return s
}

func (s *ModifyElasticResourceSpecRequestResourceSpec) SetMemoryGB(v int32) *ModifyElasticResourceSpecRequestResourceSpec {
	s.MemoryGB = &v
	return s
}

func (s *ModifyElasticResourceSpecRequestResourceSpec) Validate() error {
	return dara.Validate(s)
}
