// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPrepayNamespaceSpecRequest interface {
	dara.Model
	String() string
	GoString() string
	SetModifyPrepayNamespaceSpecRequest(v *ModifyPrepayNamespaceSpecRequestModifyPrepayNamespaceSpecRequest) *ModifyPrepayNamespaceSpecRequest
	GetModifyPrepayNamespaceSpecRequest() *ModifyPrepayNamespaceSpecRequestModifyPrepayNamespaceSpecRequest
}

type ModifyPrepayNamespaceSpecRequest struct {
	// This parameter is required.
	ModifyPrepayNamespaceSpecRequest *ModifyPrepayNamespaceSpecRequestModifyPrepayNamespaceSpecRequest `json:"ModifyPrepayNamespaceSpecRequest,omitempty" xml:"ModifyPrepayNamespaceSpecRequest,omitempty" type:"Struct"`
}

func (s ModifyPrepayNamespaceSpecRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyPrepayNamespaceSpecRequest) GoString() string {
	return s.String()
}

func (s *ModifyPrepayNamespaceSpecRequest) GetModifyPrepayNamespaceSpecRequest() *ModifyPrepayNamespaceSpecRequestModifyPrepayNamespaceSpecRequest {
	return s.ModifyPrepayNamespaceSpecRequest
}

func (s *ModifyPrepayNamespaceSpecRequest) SetModifyPrepayNamespaceSpecRequest(v *ModifyPrepayNamespaceSpecRequestModifyPrepayNamespaceSpecRequest) *ModifyPrepayNamespaceSpecRequest {
	s.ModifyPrepayNamespaceSpecRequest = v
	return s
}

func (s *ModifyPrepayNamespaceSpecRequest) Validate() error {
	if s.ModifyPrepayNamespaceSpecRequest != nil {
		if err := s.ModifyPrepayNamespaceSpecRequest.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyPrepayNamespaceSpecRequestModifyPrepayNamespaceSpecRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// sc_flinkserverlesspost_public_cn-0ju2bj2****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ns-1
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// This parameter is required.
	ResourceSpec *ModifyPrepayNamespaceSpecRequestModifyPrepayNamespaceSpecRequestResourceSpec `json:"ResourceSpec,omitempty" xml:"ResourceSpec,omitempty" type:"Struct"`
}

func (s ModifyPrepayNamespaceSpecRequestModifyPrepayNamespaceSpecRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyPrepayNamespaceSpecRequestModifyPrepayNamespaceSpecRequest) GoString() string {
	return s.String()
}

func (s *ModifyPrepayNamespaceSpecRequestModifyPrepayNamespaceSpecRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyPrepayNamespaceSpecRequestModifyPrepayNamespaceSpecRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *ModifyPrepayNamespaceSpecRequestModifyPrepayNamespaceSpecRequest) GetRegion() *string {
	return s.Region
}

func (s *ModifyPrepayNamespaceSpecRequestModifyPrepayNamespaceSpecRequest) GetResourceSpec() *ModifyPrepayNamespaceSpecRequestModifyPrepayNamespaceSpecRequestResourceSpec {
	return s.ResourceSpec
}

func (s *ModifyPrepayNamespaceSpecRequestModifyPrepayNamespaceSpecRequest) SetInstanceId(v string) *ModifyPrepayNamespaceSpecRequestModifyPrepayNamespaceSpecRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyPrepayNamespaceSpecRequestModifyPrepayNamespaceSpecRequest) SetNamespace(v string) *ModifyPrepayNamespaceSpecRequestModifyPrepayNamespaceSpecRequest {
	s.Namespace = &v
	return s
}

func (s *ModifyPrepayNamespaceSpecRequestModifyPrepayNamespaceSpecRequest) SetRegion(v string) *ModifyPrepayNamespaceSpecRequestModifyPrepayNamespaceSpecRequest {
	s.Region = &v
	return s
}

func (s *ModifyPrepayNamespaceSpecRequestModifyPrepayNamespaceSpecRequest) SetResourceSpec(v *ModifyPrepayNamespaceSpecRequestModifyPrepayNamespaceSpecRequestResourceSpec) *ModifyPrepayNamespaceSpecRequestModifyPrepayNamespaceSpecRequest {
	s.ResourceSpec = v
	return s
}

func (s *ModifyPrepayNamespaceSpecRequestModifyPrepayNamespaceSpecRequest) Validate() error {
	if s.ResourceSpec != nil {
		if err := s.ResourceSpec.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyPrepayNamespaceSpecRequestModifyPrepayNamespaceSpecRequestResourceSpec struct {
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

func (s ModifyPrepayNamespaceSpecRequestModifyPrepayNamespaceSpecRequestResourceSpec) String() string {
	return dara.Prettify(s)
}

func (s ModifyPrepayNamespaceSpecRequestModifyPrepayNamespaceSpecRequestResourceSpec) GoString() string {
	return s.String()
}

func (s *ModifyPrepayNamespaceSpecRequestModifyPrepayNamespaceSpecRequestResourceSpec) GetCpu() *int32 {
	return s.Cpu
}

func (s *ModifyPrepayNamespaceSpecRequestModifyPrepayNamespaceSpecRequestResourceSpec) GetMemoryGB() *int32 {
	return s.MemoryGB
}

func (s *ModifyPrepayNamespaceSpecRequestModifyPrepayNamespaceSpecRequestResourceSpec) SetCpu(v int32) *ModifyPrepayNamespaceSpecRequestModifyPrepayNamespaceSpecRequestResourceSpec {
	s.Cpu = &v
	return s
}

func (s *ModifyPrepayNamespaceSpecRequestModifyPrepayNamespaceSpecRequestResourceSpec) SetMemoryGB(v int32) *ModifyPrepayNamespaceSpecRequestModifyPrepayNamespaceSpecRequestResourceSpec {
	s.MemoryGB = &v
	return s
}

func (s *ModifyPrepayNamespaceSpecRequestModifyPrepayNamespaceSpecRequestResourceSpec) Validate() error {
	return dara.Validate(s)
}
