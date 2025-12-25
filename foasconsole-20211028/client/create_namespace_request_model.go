// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNamespaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHa(v bool) *CreateNamespaceRequest
	GetHa() *bool
	SetInstanceId(v string) *CreateNamespaceRequest
	GetInstanceId() *string
	SetNamespace(v string) *CreateNamespaceRequest
	GetNamespace() *string
	SetRegion(v string) *CreateNamespaceRequest
	GetRegion() *string
	SetResourceSpec(v *CreateNamespaceRequestResourceSpec) *CreateNamespaceRequest
	GetResourceSpec() *CreateNamespaceRequestResourceSpec
}

type CreateNamespaceRequest struct {
	// if can be null:
	// true
	Ha *bool `json:"Ha,omitempty" xml:"Ha,omitempty"`
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
	// di-593440390152545
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-shenzhen
	Region       *string                             `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceSpec *CreateNamespaceRequestResourceSpec `json:"ResourceSpec,omitempty" xml:"ResourceSpec,omitempty" type:"Struct"`
}

func (s CreateNamespaceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateNamespaceRequest) GoString() string {
	return s.String()
}

func (s *CreateNamespaceRequest) GetHa() *bool {
	return s.Ha
}

func (s *CreateNamespaceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateNamespaceRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *CreateNamespaceRequest) GetRegion() *string {
	return s.Region
}

func (s *CreateNamespaceRequest) GetResourceSpec() *CreateNamespaceRequestResourceSpec {
	return s.ResourceSpec
}

func (s *CreateNamespaceRequest) SetHa(v bool) *CreateNamespaceRequest {
	s.Ha = &v
	return s
}

func (s *CreateNamespaceRequest) SetInstanceId(v string) *CreateNamespaceRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateNamespaceRequest) SetNamespace(v string) *CreateNamespaceRequest {
	s.Namespace = &v
	return s
}

func (s *CreateNamespaceRequest) SetRegion(v string) *CreateNamespaceRequest {
	s.Region = &v
	return s
}

func (s *CreateNamespaceRequest) SetResourceSpec(v *CreateNamespaceRequestResourceSpec) *CreateNamespaceRequest {
	s.ResourceSpec = v
	return s
}

func (s *CreateNamespaceRequest) Validate() error {
	if s.ResourceSpec != nil {
		if err := s.ResourceSpec.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateNamespaceRequestResourceSpec struct {
	// example:
	//
	// 30
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// example:
	//
	// 120
	MemoryGB *int32 `json:"MemoryGB,omitempty" xml:"MemoryGB,omitempty"`
}

func (s CreateNamespaceRequestResourceSpec) String() string {
	return dara.Prettify(s)
}

func (s CreateNamespaceRequestResourceSpec) GoString() string {
	return s.String()
}

func (s *CreateNamespaceRequestResourceSpec) GetCpu() *int32 {
	return s.Cpu
}

func (s *CreateNamespaceRequestResourceSpec) GetMemoryGB() *int32 {
	return s.MemoryGB
}

func (s *CreateNamespaceRequestResourceSpec) SetCpu(v int32) *CreateNamespaceRequestResourceSpec {
	s.Cpu = &v
	return s
}

func (s *CreateNamespaceRequestResourceSpec) SetMemoryGB(v int32) *CreateNamespaceRequestResourceSpec {
	s.MemoryGB = &v
	return s
}

func (s *CreateNamespaceRequestResourceSpec) Validate() error {
	return dara.Validate(s)
}
