// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceSpecRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHa(v bool) *ModifyInstanceSpecRequest
	GetHa() *bool
	SetHaResourceSpec(v *ModifyInstanceSpecRequestHaResourceSpec) *ModifyInstanceSpecRequest
	GetHaResourceSpec() *ModifyInstanceSpecRequestHaResourceSpec
	SetHaVSwitchIds(v []*string) *ModifyInstanceSpecRequest
	GetHaVSwitchIds() []*string
	SetInstanceId(v string) *ModifyInstanceSpecRequest
	GetInstanceId() *string
	SetPromotionCode(v string) *ModifyInstanceSpecRequest
	GetPromotionCode() *string
	SetRegion(v string) *ModifyInstanceSpecRequest
	GetRegion() *string
	SetResourceSpec(v *ModifyInstanceSpecRequestResourceSpec) *ModifyInstanceSpecRequest
	GetResourceSpec() *ModifyInstanceSpecRequestResourceSpec
	SetUsePromotionCode(v bool) *ModifyInstanceSpecRequest
	GetUsePromotionCode() *bool
}

type ModifyInstanceSpecRequest struct {
	// example:
	//
	// true
	Ha             *bool                                    `json:"Ha,omitempty" xml:"Ha,omitempty"`
	HaResourceSpec *ModifyInstanceSpecRequestHaResourceSpec `json:"HaResourceSpec,omitempty" xml:"HaResourceSpec,omitempty" type:"Struct"`
	HaVSwitchIds   []*string                                `json:"HaVSwitchIds,omitempty" xml:"HaVSwitchIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// f-cn-wwo36qj4g06
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PromotionCode *string `json:"PromotionCode,omitempty" xml:"PromotionCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// This parameter is required.
	ResourceSpec     *ModifyInstanceSpecRequestResourceSpec `json:"ResourceSpec,omitempty" xml:"ResourceSpec,omitempty" type:"Struct"`
	UsePromotionCode *bool                                  `json:"UsePromotionCode,omitempty" xml:"UsePromotionCode,omitempty"`
}

func (s ModifyInstanceSpecRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceSpecRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceSpecRequest) GetHa() *bool {
	return s.Ha
}

func (s *ModifyInstanceSpecRequest) GetHaResourceSpec() *ModifyInstanceSpecRequestHaResourceSpec {
	return s.HaResourceSpec
}

func (s *ModifyInstanceSpecRequest) GetHaVSwitchIds() []*string {
	return s.HaVSwitchIds
}

func (s *ModifyInstanceSpecRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyInstanceSpecRequest) GetPromotionCode() *string {
	return s.PromotionCode
}

func (s *ModifyInstanceSpecRequest) GetRegion() *string {
	return s.Region
}

func (s *ModifyInstanceSpecRequest) GetResourceSpec() *ModifyInstanceSpecRequestResourceSpec {
	return s.ResourceSpec
}

func (s *ModifyInstanceSpecRequest) GetUsePromotionCode() *bool {
	return s.UsePromotionCode
}

func (s *ModifyInstanceSpecRequest) SetHa(v bool) *ModifyInstanceSpecRequest {
	s.Ha = &v
	return s
}

func (s *ModifyInstanceSpecRequest) SetHaResourceSpec(v *ModifyInstanceSpecRequestHaResourceSpec) *ModifyInstanceSpecRequest {
	s.HaResourceSpec = v
	return s
}

func (s *ModifyInstanceSpecRequest) SetHaVSwitchIds(v []*string) *ModifyInstanceSpecRequest {
	s.HaVSwitchIds = v
	return s
}

func (s *ModifyInstanceSpecRequest) SetInstanceId(v string) *ModifyInstanceSpecRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceSpecRequest) SetPromotionCode(v string) *ModifyInstanceSpecRequest {
	s.PromotionCode = &v
	return s
}

func (s *ModifyInstanceSpecRequest) SetRegion(v string) *ModifyInstanceSpecRequest {
	s.Region = &v
	return s
}

func (s *ModifyInstanceSpecRequest) SetResourceSpec(v *ModifyInstanceSpecRequestResourceSpec) *ModifyInstanceSpecRequest {
	s.ResourceSpec = v
	return s
}

func (s *ModifyInstanceSpecRequest) SetUsePromotionCode(v bool) *ModifyInstanceSpecRequest {
	s.UsePromotionCode = &v
	return s
}

func (s *ModifyInstanceSpecRequest) Validate() error {
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

type ModifyInstanceSpecRequestHaResourceSpec struct {
	// example:
	//
	// 10
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// example:
	//
	// 40
	MemoryGB *int32 `json:"MemoryGB,omitempty" xml:"MemoryGB,omitempty"`
}

func (s ModifyInstanceSpecRequestHaResourceSpec) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceSpecRequestHaResourceSpec) GoString() string {
	return s.String()
}

func (s *ModifyInstanceSpecRequestHaResourceSpec) GetCpu() *int32 {
	return s.Cpu
}

func (s *ModifyInstanceSpecRequestHaResourceSpec) GetMemoryGB() *int32 {
	return s.MemoryGB
}

func (s *ModifyInstanceSpecRequestHaResourceSpec) SetCpu(v int32) *ModifyInstanceSpecRequestHaResourceSpec {
	s.Cpu = &v
	return s
}

func (s *ModifyInstanceSpecRequestHaResourceSpec) SetMemoryGB(v int32) *ModifyInstanceSpecRequestHaResourceSpec {
	s.MemoryGB = &v
	return s
}

func (s *ModifyInstanceSpecRequestHaResourceSpec) Validate() error {
	return dara.Validate(s)
}

type ModifyInstanceSpecRequestResourceSpec struct {
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

func (s ModifyInstanceSpecRequestResourceSpec) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceSpecRequestResourceSpec) GoString() string {
	return s.String()
}

func (s *ModifyInstanceSpecRequestResourceSpec) GetCpu() *int32 {
	return s.Cpu
}

func (s *ModifyInstanceSpecRequestResourceSpec) GetMemoryGB() *int32 {
	return s.MemoryGB
}

func (s *ModifyInstanceSpecRequestResourceSpec) SetCpu(v int32) *ModifyInstanceSpecRequestResourceSpec {
	s.Cpu = &v
	return s
}

func (s *ModifyInstanceSpecRequestResourceSpec) SetMemoryGB(v int32) *ModifyInstanceSpecRequestResourceSpec {
	s.MemoryGB = &v
	return s
}

func (s *ModifyInstanceSpecRequestResourceSpec) Validate() error {
	return dara.Validate(s)
}
