// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPrepayInstanceSpecRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHa(v bool) *ModifyPrepayInstanceSpecRequest
	GetHa() *bool
	SetHaResourceSpec(v *ModifyPrepayInstanceSpecRequestHaResourceSpec) *ModifyPrepayInstanceSpecRequest
	GetHaResourceSpec() *ModifyPrepayInstanceSpecRequestHaResourceSpec
	SetHaVSwitchIds(v []*string) *ModifyPrepayInstanceSpecRequest
	GetHaVSwitchIds() []*string
	SetHaZoneId(v string) *ModifyPrepayInstanceSpecRequest
	GetHaZoneId() *string
	SetInstanceId(v string) *ModifyPrepayInstanceSpecRequest
	GetInstanceId() *string
	SetRegion(v string) *ModifyPrepayInstanceSpecRequest
	GetRegion() *string
	SetResourceSpec(v *ModifyPrepayInstanceSpecRequestResourceSpec) *ModifyPrepayInstanceSpecRequest
	GetResourceSpec() *ModifyPrepayInstanceSpecRequestResourceSpec
}

type ModifyPrepayInstanceSpecRequest struct {
	// if can be null:
	// true
	Ha *bool `json:"Ha,omitempty" xml:"Ha,omitempty"`
	// if can be null:
	// true
	HaResourceSpec *ModifyPrepayInstanceSpecRequestHaResourceSpec `json:"HaResourceSpec,omitempty" xml:"HaResourceSpec,omitempty" type:"Struct"`
	// if can be null:
	// true
	HaVSwitchIds []*string `json:"HaVSwitchIds,omitempty" xml:"HaVSwitchIds,omitempty" type:"Repeated"`
	// if can be null:
	// true
	HaZoneId *string `json:"HaZoneId,omitempty" xml:"HaZoneId,omitempty"`
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
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// This parameter is required.
	ResourceSpec *ModifyPrepayInstanceSpecRequestResourceSpec `json:"ResourceSpec,omitempty" xml:"ResourceSpec,omitempty" type:"Struct"`
}

func (s ModifyPrepayInstanceSpecRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyPrepayInstanceSpecRequest) GoString() string {
	return s.String()
}

func (s *ModifyPrepayInstanceSpecRequest) GetHa() *bool {
	return s.Ha
}

func (s *ModifyPrepayInstanceSpecRequest) GetHaResourceSpec() *ModifyPrepayInstanceSpecRequestHaResourceSpec {
	return s.HaResourceSpec
}

func (s *ModifyPrepayInstanceSpecRequest) GetHaVSwitchIds() []*string {
	return s.HaVSwitchIds
}

func (s *ModifyPrepayInstanceSpecRequest) GetHaZoneId() *string {
	return s.HaZoneId
}

func (s *ModifyPrepayInstanceSpecRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyPrepayInstanceSpecRequest) GetRegion() *string {
	return s.Region
}

func (s *ModifyPrepayInstanceSpecRequest) GetResourceSpec() *ModifyPrepayInstanceSpecRequestResourceSpec {
	return s.ResourceSpec
}

func (s *ModifyPrepayInstanceSpecRequest) SetHa(v bool) *ModifyPrepayInstanceSpecRequest {
	s.Ha = &v
	return s
}

func (s *ModifyPrepayInstanceSpecRequest) SetHaResourceSpec(v *ModifyPrepayInstanceSpecRequestHaResourceSpec) *ModifyPrepayInstanceSpecRequest {
	s.HaResourceSpec = v
	return s
}

func (s *ModifyPrepayInstanceSpecRequest) SetHaVSwitchIds(v []*string) *ModifyPrepayInstanceSpecRequest {
	s.HaVSwitchIds = v
	return s
}

func (s *ModifyPrepayInstanceSpecRequest) SetHaZoneId(v string) *ModifyPrepayInstanceSpecRequest {
	s.HaZoneId = &v
	return s
}

func (s *ModifyPrepayInstanceSpecRequest) SetInstanceId(v string) *ModifyPrepayInstanceSpecRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyPrepayInstanceSpecRequest) SetRegion(v string) *ModifyPrepayInstanceSpecRequest {
	s.Region = &v
	return s
}

func (s *ModifyPrepayInstanceSpecRequest) SetResourceSpec(v *ModifyPrepayInstanceSpecRequestResourceSpec) *ModifyPrepayInstanceSpecRequest {
	s.ResourceSpec = v
	return s
}

func (s *ModifyPrepayInstanceSpecRequest) Validate() error {
	if s.HaResourceSpec != nil {
		if err := s.HaResourceSpec.Validate(); err != nil {
			return err
		}
	}
	if s.ResourceSpec != nil {
		if err := s.ResourceSpec.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyPrepayInstanceSpecRequestHaResourceSpec struct {
	Cpu      *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	MemoryGB *int32 `json:"MemoryGB,omitempty" xml:"MemoryGB,omitempty"`
}

func (s ModifyPrepayInstanceSpecRequestHaResourceSpec) String() string {
	return dara.Prettify(s)
}

func (s ModifyPrepayInstanceSpecRequestHaResourceSpec) GoString() string {
	return s.String()
}

func (s *ModifyPrepayInstanceSpecRequestHaResourceSpec) GetCpu() *int32 {
	return s.Cpu
}

func (s *ModifyPrepayInstanceSpecRequestHaResourceSpec) GetMemoryGB() *int32 {
	return s.MemoryGB
}

func (s *ModifyPrepayInstanceSpecRequestHaResourceSpec) SetCpu(v int32) *ModifyPrepayInstanceSpecRequestHaResourceSpec {
	s.Cpu = &v
	return s
}

func (s *ModifyPrepayInstanceSpecRequestHaResourceSpec) SetMemoryGB(v int32) *ModifyPrepayInstanceSpecRequestHaResourceSpec {
	s.MemoryGB = &v
	return s
}

func (s *ModifyPrepayInstanceSpecRequestHaResourceSpec) Validate() error {
	return dara.Validate(s)
}

type ModifyPrepayInstanceSpecRequestResourceSpec struct {
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

func (s ModifyPrepayInstanceSpecRequestResourceSpec) String() string {
	return dara.Prettify(s)
}

func (s ModifyPrepayInstanceSpecRequestResourceSpec) GoString() string {
	return s.String()
}

func (s *ModifyPrepayInstanceSpecRequestResourceSpec) GetCpu() *int32 {
	return s.Cpu
}

func (s *ModifyPrepayInstanceSpecRequestResourceSpec) GetMemoryGB() *int32 {
	return s.MemoryGB
}

func (s *ModifyPrepayInstanceSpecRequestResourceSpec) SetCpu(v int32) *ModifyPrepayInstanceSpecRequestResourceSpec {
	s.Cpu = &v
	return s
}

func (s *ModifyPrepayInstanceSpecRequestResourceSpec) SetMemoryGB(v int32) *ModifyPrepayInstanceSpecRequestResourceSpec {
	s.MemoryGB = &v
	return s
}

func (s *ModifyPrepayInstanceSpecRequestResourceSpec) Validate() error {
	return dara.Validate(s)
}
