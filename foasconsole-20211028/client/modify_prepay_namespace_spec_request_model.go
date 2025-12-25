// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPrepayNamespaceSpecRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyPrepayNamespaceSpecRequest
	GetInstanceId() *string
	SetNamespace(v string) *ModifyPrepayNamespaceSpecRequest
	GetNamespace() *string
	SetRegion(v string) *ModifyPrepayNamespaceSpecRequest
	GetRegion() *string
	SetResourceSpec(v *ModifyPrepayNamespaceSpecRequestResourceSpec) *ModifyPrepayNamespaceSpecRequest
	GetResourceSpec() *ModifyPrepayNamespaceSpecRequestResourceSpec
}

type ModifyPrepayNamespaceSpecRequest struct {
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
	// di-593440219799842
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-shenzhen
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// This parameter is required.
	ResourceSpec *ModifyPrepayNamespaceSpecRequestResourceSpec `json:"ResourceSpec,omitempty" xml:"ResourceSpec,omitempty" type:"Struct"`
}

func (s ModifyPrepayNamespaceSpecRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyPrepayNamespaceSpecRequest) GoString() string {
	return s.String()
}

func (s *ModifyPrepayNamespaceSpecRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyPrepayNamespaceSpecRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *ModifyPrepayNamespaceSpecRequest) GetRegion() *string {
	return s.Region
}

func (s *ModifyPrepayNamespaceSpecRequest) GetResourceSpec() *ModifyPrepayNamespaceSpecRequestResourceSpec {
	return s.ResourceSpec
}

func (s *ModifyPrepayNamespaceSpecRequest) SetInstanceId(v string) *ModifyPrepayNamespaceSpecRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyPrepayNamespaceSpecRequest) SetNamespace(v string) *ModifyPrepayNamespaceSpecRequest {
	s.Namespace = &v
	return s
}

func (s *ModifyPrepayNamespaceSpecRequest) SetRegion(v string) *ModifyPrepayNamespaceSpecRequest {
	s.Region = &v
	return s
}

func (s *ModifyPrepayNamespaceSpecRequest) SetResourceSpec(v *ModifyPrepayNamespaceSpecRequestResourceSpec) *ModifyPrepayNamespaceSpecRequest {
	s.ResourceSpec = v
	return s
}

func (s *ModifyPrepayNamespaceSpecRequest) Validate() error {
	if s.ResourceSpec != nil {
		if err := s.ResourceSpec.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyPrepayNamespaceSpecRequestResourceSpec struct {
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

func (s ModifyPrepayNamespaceSpecRequestResourceSpec) String() string {
	return dara.Prettify(s)
}

func (s ModifyPrepayNamespaceSpecRequestResourceSpec) GoString() string {
	return s.String()
}

func (s *ModifyPrepayNamespaceSpecRequestResourceSpec) GetCpu() *int32 {
	return s.Cpu
}

func (s *ModifyPrepayNamespaceSpecRequestResourceSpec) GetMemoryGB() *int32 {
	return s.MemoryGB
}

func (s *ModifyPrepayNamespaceSpecRequestResourceSpec) SetCpu(v int32) *ModifyPrepayNamespaceSpecRequestResourceSpec {
	s.Cpu = &v
	return s
}

func (s *ModifyPrepayNamespaceSpecRequestResourceSpec) SetMemoryGB(v int32) *ModifyPrepayNamespaceSpecRequestResourceSpec {
	s.MemoryGB = &v
	return s
}

func (s *ModifyPrepayNamespaceSpecRequestResourceSpec) Validate() error {
	return dara.Validate(s)
}
