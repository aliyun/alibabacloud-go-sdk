// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNamespaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateNamespaceRequest(v *CreateNamespaceRequestCreateNamespaceRequest) *CreateNamespaceRequest
	GetCreateNamespaceRequest() *CreateNamespaceRequestCreateNamespaceRequest
}

type CreateNamespaceRequest struct {
	// This parameter is required.
	CreateNamespaceRequest *CreateNamespaceRequestCreateNamespaceRequest `json:"CreateNamespaceRequest,omitempty" xml:"CreateNamespaceRequest,omitempty" type:"Struct"`
}

func (s CreateNamespaceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateNamespaceRequest) GoString() string {
	return s.String()
}

func (s *CreateNamespaceRequest) GetCreateNamespaceRequest() *CreateNamespaceRequestCreateNamespaceRequest {
	return s.CreateNamespaceRequest
}

func (s *CreateNamespaceRequest) SetCreateNamespaceRequest(v *CreateNamespaceRequestCreateNamespaceRequest) *CreateNamespaceRequest {
	s.CreateNamespaceRequest = v
	return s
}

func (s *CreateNamespaceRequest) Validate() error {
	if s.CreateNamespaceRequest != nil {
		if err := s.CreateNamespaceRequest.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateNamespaceRequestCreateNamespaceRequest struct {
	// if can be null:
	// true
	Ha *bool `json:"Ha,omitempty" xml:"Ha,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 223493C7-FCA9-13D4-B75B-AF8B32F4****
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
	Region       *string                                                   `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceSpec *CreateNamespaceRequestCreateNamespaceRequestResourceSpec `json:"ResourceSpec,omitempty" xml:"ResourceSpec,omitempty" type:"Struct"`
}

func (s CreateNamespaceRequestCreateNamespaceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateNamespaceRequestCreateNamespaceRequest) GoString() string {
	return s.String()
}

func (s *CreateNamespaceRequestCreateNamespaceRequest) GetHa() *bool {
	return s.Ha
}

func (s *CreateNamespaceRequestCreateNamespaceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateNamespaceRequestCreateNamespaceRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *CreateNamespaceRequestCreateNamespaceRequest) GetRegion() *string {
	return s.Region
}

func (s *CreateNamespaceRequestCreateNamespaceRequest) GetResourceSpec() *CreateNamespaceRequestCreateNamespaceRequestResourceSpec {
	return s.ResourceSpec
}

func (s *CreateNamespaceRequestCreateNamespaceRequest) SetHa(v bool) *CreateNamespaceRequestCreateNamespaceRequest {
	s.Ha = &v
	return s
}

func (s *CreateNamespaceRequestCreateNamespaceRequest) SetInstanceId(v string) *CreateNamespaceRequestCreateNamespaceRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateNamespaceRequestCreateNamespaceRequest) SetNamespace(v string) *CreateNamespaceRequestCreateNamespaceRequest {
	s.Namespace = &v
	return s
}

func (s *CreateNamespaceRequestCreateNamespaceRequest) SetRegion(v string) *CreateNamespaceRequestCreateNamespaceRequest {
	s.Region = &v
	return s
}

func (s *CreateNamespaceRequestCreateNamespaceRequest) SetResourceSpec(v *CreateNamespaceRequestCreateNamespaceRequestResourceSpec) *CreateNamespaceRequestCreateNamespaceRequest {
	s.ResourceSpec = v
	return s
}

func (s *CreateNamespaceRequestCreateNamespaceRequest) Validate() error {
	if s.ResourceSpec != nil {
		if err := s.ResourceSpec.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateNamespaceRequestCreateNamespaceRequestResourceSpec struct {
	// example:
	//
	// 10
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// example:
	//
	// 40
	MemoryGB *int32 `json:"MemoryGB,omitempty" xml:"MemoryGB,omitempty"`
}

func (s CreateNamespaceRequestCreateNamespaceRequestResourceSpec) String() string {
	return dara.Prettify(s)
}

func (s CreateNamespaceRequestCreateNamespaceRequestResourceSpec) GoString() string {
	return s.String()
}

func (s *CreateNamespaceRequestCreateNamespaceRequestResourceSpec) GetCpu() *int32 {
	return s.Cpu
}

func (s *CreateNamespaceRequestCreateNamespaceRequestResourceSpec) GetMemoryGB() *int32 {
	return s.MemoryGB
}

func (s *CreateNamespaceRequestCreateNamespaceRequestResourceSpec) SetCpu(v int32) *CreateNamespaceRequestCreateNamespaceRequestResourceSpec {
	s.Cpu = &v
	return s
}

func (s *CreateNamespaceRequestCreateNamespaceRequestResourceSpec) SetMemoryGB(v int32) *CreateNamespaceRequestCreateNamespaceRequestResourceSpec {
	s.MemoryGB = &v
	return s
}

func (s *CreateNamespaceRequestCreateNamespaceRequestResourceSpec) Validate() error {
	return dara.Validate(s)
}
