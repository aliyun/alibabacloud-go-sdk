// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPrepayInstanceSpecRequest interface {
	dara.Model
	String() string
	GoString() string
	SetModifyPrepayInstanceSpecRequest(v *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequest) *ModifyPrepayInstanceSpecRequest
	GetModifyPrepayInstanceSpecRequest() *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequest
}

type ModifyPrepayInstanceSpecRequest struct {
	// This parameter is required.
	ModifyPrepayInstanceSpecRequest *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequest `json:"ModifyPrepayInstanceSpecRequest,omitempty" xml:"ModifyPrepayInstanceSpecRequest,omitempty" type:"Struct"`
}

func (s ModifyPrepayInstanceSpecRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyPrepayInstanceSpecRequest) GoString() string {
	return s.String()
}

func (s *ModifyPrepayInstanceSpecRequest) GetModifyPrepayInstanceSpecRequest() *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequest {
	return s.ModifyPrepayInstanceSpecRequest
}

func (s *ModifyPrepayInstanceSpecRequest) SetModifyPrepayInstanceSpecRequest(v *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequest) *ModifyPrepayInstanceSpecRequest {
	s.ModifyPrepayInstanceSpecRequest = v
	return s
}

func (s *ModifyPrepayInstanceSpecRequest) Validate() error {
	if s.ModifyPrepayInstanceSpecRequest != nil {
		if err := s.ModifyPrepayInstanceSpecRequest.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequest struct {
	// if can be null:
	// true
	Ha *bool `json:"Ha,omitempty" xml:"Ha,omitempty"`
	// if can be null:
	// true
	HaResourceSpec *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequestHaResourceSpec `json:"HaResourceSpec,omitempty" xml:"HaResourceSpec,omitempty" type:"Struct"`
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
	// sc_flinkserverlesspost_public_cn-0ju2bj2i104
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// This parameter is required.
	ResourceSpec *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequestResourceSpec `json:"ResourceSpec,omitempty" xml:"ResourceSpec,omitempty" type:"Struct"`
}

func (s ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequest) GoString() string {
	return s.String()
}

func (s *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequest) GetHa() *bool {
	return s.Ha
}

func (s *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequest) GetHaResourceSpec() *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequestHaResourceSpec {
	return s.HaResourceSpec
}

func (s *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequest) GetHaVSwitchIds() []*string {
	return s.HaVSwitchIds
}

func (s *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequest) GetHaZoneId() *string {
	return s.HaZoneId
}

func (s *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequest) GetRegion() *string {
	return s.Region
}

func (s *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequest) GetResourceSpec() *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequestResourceSpec {
	return s.ResourceSpec
}

func (s *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequest) SetHa(v bool) *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequest {
	s.Ha = &v
	return s
}

func (s *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequest) SetHaResourceSpec(v *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequestHaResourceSpec) *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequest {
	s.HaResourceSpec = v
	return s
}

func (s *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequest) SetHaVSwitchIds(v []*string) *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequest {
	s.HaVSwitchIds = v
	return s
}

func (s *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequest) SetHaZoneId(v string) *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequest {
	s.HaZoneId = &v
	return s
}

func (s *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequest) SetInstanceId(v string) *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequest) SetRegion(v string) *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequest {
	s.Region = &v
	return s
}

func (s *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequest) SetResourceSpec(v *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequestResourceSpec) *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequest {
	s.ResourceSpec = v
	return s
}

func (s *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequest) Validate() error {
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

type ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequestHaResourceSpec struct {
	Cpu      *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	MemoryGB *int32 `json:"MemoryGB,omitempty" xml:"MemoryGB,omitempty"`
}

func (s ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequestHaResourceSpec) String() string {
	return dara.Prettify(s)
}

func (s ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequestHaResourceSpec) GoString() string {
	return s.String()
}

func (s *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequestHaResourceSpec) GetCpu() *int32 {
	return s.Cpu
}

func (s *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequestHaResourceSpec) GetMemoryGB() *int32 {
	return s.MemoryGB
}

func (s *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequestHaResourceSpec) SetCpu(v int32) *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequestHaResourceSpec {
	s.Cpu = &v
	return s
}

func (s *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequestHaResourceSpec) SetMemoryGB(v int32) *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequestHaResourceSpec {
	s.MemoryGB = &v
	return s
}

func (s *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequestHaResourceSpec) Validate() error {
	return dara.Validate(s)
}

type ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequestResourceSpec struct {
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

func (s ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequestResourceSpec) String() string {
	return dara.Prettify(s)
}

func (s ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequestResourceSpec) GoString() string {
	return s.String()
}

func (s *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequestResourceSpec) GetCpu() *int32 {
	return s.Cpu
}

func (s *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequestResourceSpec) GetMemoryGB() *int32 {
	return s.MemoryGB
}

func (s *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequestResourceSpec) SetCpu(v int32) *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequestResourceSpec {
	s.Cpu = &v
	return s
}

func (s *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequestResourceSpec) SetMemoryGB(v int32) *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequestResourceSpec {
	s.MemoryGB = &v
	return s
}

func (s *ModifyPrepayInstanceSpecRequestModifyPrepayInstanceSpecRequestResourceSpec) Validate() error {
	return dara.Validate(s)
}
