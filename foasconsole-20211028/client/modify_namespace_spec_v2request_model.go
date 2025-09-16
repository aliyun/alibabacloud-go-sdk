// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNamespaceSpecV2Request interface {
	dara.Model
	String() string
	GoString() string
	SetElasticResourceSpec(v *ModifyNamespaceSpecV2RequestElasticResourceSpec) *ModifyNamespaceSpecV2Request
	GetElasticResourceSpec() *ModifyNamespaceSpecV2RequestElasticResourceSpec
	SetGuaranteedResourceSpec(v *ModifyNamespaceSpecV2RequestGuaranteedResourceSpec) *ModifyNamespaceSpecV2Request
	GetGuaranteedResourceSpec() *ModifyNamespaceSpecV2RequestGuaranteedResourceSpec
	SetHa(v bool) *ModifyNamespaceSpecV2Request
	GetHa() *bool
	SetInstanceId(v string) *ModifyNamespaceSpecV2Request
	GetInstanceId() *string
	SetNamespace(v string) *ModifyNamespaceSpecV2Request
	GetNamespace() *string
	SetRegion(v string) *ModifyNamespaceSpecV2Request
	GetRegion() *string
}

type ModifyNamespaceSpecV2Request struct {
	ElasticResourceSpec    *ModifyNamespaceSpecV2RequestElasticResourceSpec    `json:"ElasticResourceSpec,omitempty" xml:"ElasticResourceSpec,omitempty" type:"Struct"`
	GuaranteedResourceSpec *ModifyNamespaceSpecV2RequestGuaranteedResourceSpec `json:"GuaranteedResourceSpec,omitempty" xml:"GuaranteedResourceSpec,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// false
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
	// di-593439443804417
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s ModifyNamespaceSpecV2Request) String() string {
	return dara.Prettify(s)
}

func (s ModifyNamespaceSpecV2Request) GoString() string {
	return s.String()
}

func (s *ModifyNamespaceSpecV2Request) GetElasticResourceSpec() *ModifyNamespaceSpecV2RequestElasticResourceSpec {
	return s.ElasticResourceSpec
}

func (s *ModifyNamespaceSpecV2Request) GetGuaranteedResourceSpec() *ModifyNamespaceSpecV2RequestGuaranteedResourceSpec {
	return s.GuaranteedResourceSpec
}

func (s *ModifyNamespaceSpecV2Request) GetHa() *bool {
	return s.Ha
}

func (s *ModifyNamespaceSpecV2Request) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyNamespaceSpecV2Request) GetNamespace() *string {
	return s.Namespace
}

func (s *ModifyNamespaceSpecV2Request) GetRegion() *string {
	return s.Region
}

func (s *ModifyNamespaceSpecV2Request) SetElasticResourceSpec(v *ModifyNamespaceSpecV2RequestElasticResourceSpec) *ModifyNamespaceSpecV2Request {
	s.ElasticResourceSpec = v
	return s
}

func (s *ModifyNamespaceSpecV2Request) SetGuaranteedResourceSpec(v *ModifyNamespaceSpecV2RequestGuaranteedResourceSpec) *ModifyNamespaceSpecV2Request {
	s.GuaranteedResourceSpec = v
	return s
}

func (s *ModifyNamespaceSpecV2Request) SetHa(v bool) *ModifyNamespaceSpecV2Request {
	s.Ha = &v
	return s
}

func (s *ModifyNamespaceSpecV2Request) SetInstanceId(v string) *ModifyNamespaceSpecV2Request {
	s.InstanceId = &v
	return s
}

func (s *ModifyNamespaceSpecV2Request) SetNamespace(v string) *ModifyNamespaceSpecV2Request {
	s.Namespace = &v
	return s
}

func (s *ModifyNamespaceSpecV2Request) SetRegion(v string) *ModifyNamespaceSpecV2Request {
	s.Region = &v
	return s
}

func (s *ModifyNamespaceSpecV2Request) Validate() error {
	return dara.Validate(s)
}

type ModifyNamespaceSpecV2RequestElasticResourceSpec struct {
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
	// 52
	MemoryGB *int32 `json:"MemoryGB,omitempty" xml:"MemoryGB,omitempty"`
}

func (s ModifyNamespaceSpecV2RequestElasticResourceSpec) String() string {
	return dara.Prettify(s)
}

func (s ModifyNamespaceSpecV2RequestElasticResourceSpec) GoString() string {
	return s.String()
}

func (s *ModifyNamespaceSpecV2RequestElasticResourceSpec) GetCpu() *int32 {
	return s.Cpu
}

func (s *ModifyNamespaceSpecV2RequestElasticResourceSpec) GetMemoryGB() *int32 {
	return s.MemoryGB
}

func (s *ModifyNamespaceSpecV2RequestElasticResourceSpec) SetCpu(v int32) *ModifyNamespaceSpecV2RequestElasticResourceSpec {
	s.Cpu = &v
	return s
}

func (s *ModifyNamespaceSpecV2RequestElasticResourceSpec) SetMemoryGB(v int32) *ModifyNamespaceSpecV2RequestElasticResourceSpec {
	s.MemoryGB = &v
	return s
}

func (s *ModifyNamespaceSpecV2RequestElasticResourceSpec) Validate() error {
	return dara.Validate(s)
}

type ModifyNamespaceSpecV2RequestGuaranteedResourceSpec struct {
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

func (s ModifyNamespaceSpecV2RequestGuaranteedResourceSpec) String() string {
	return dara.Prettify(s)
}

func (s ModifyNamespaceSpecV2RequestGuaranteedResourceSpec) GoString() string {
	return s.String()
}

func (s *ModifyNamespaceSpecV2RequestGuaranteedResourceSpec) GetCpu() *int32 {
	return s.Cpu
}

func (s *ModifyNamespaceSpecV2RequestGuaranteedResourceSpec) GetMemoryGB() *int32 {
	return s.MemoryGB
}

func (s *ModifyNamespaceSpecV2RequestGuaranteedResourceSpec) SetCpu(v int32) *ModifyNamespaceSpecV2RequestGuaranteedResourceSpec {
	s.Cpu = &v
	return s
}

func (s *ModifyNamespaceSpecV2RequestGuaranteedResourceSpec) SetMemoryGB(v int32) *ModifyNamespaceSpecV2RequestGuaranteedResourceSpec {
	s.MemoryGB = &v
	return s
}

func (s *ModifyNamespaceSpecV2RequestGuaranteedResourceSpec) Validate() error {
	return dara.Validate(s)
}
