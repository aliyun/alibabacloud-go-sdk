// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryModifyInstancePriceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetModifyPrepayInstanceSpecRequest(v *QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequest) *QueryModifyInstancePriceRequest
	GetModifyPrepayInstanceSpecRequest() *QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequest
}

type QueryModifyInstancePriceRequest struct {
	// This parameter is required.
	ModifyPrepayInstanceSpecRequest *QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequest `json:"ModifyPrepayInstanceSpecRequest,omitempty" xml:"ModifyPrepayInstanceSpecRequest,omitempty" type:"Struct"`
}

func (s QueryModifyInstancePriceRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryModifyInstancePriceRequest) GoString() string {
	return s.String()
}

func (s *QueryModifyInstancePriceRequest) GetModifyPrepayInstanceSpecRequest() *QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequest {
	return s.ModifyPrepayInstanceSpecRequest
}

func (s *QueryModifyInstancePriceRequest) SetModifyPrepayInstanceSpecRequest(v *QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequest) *QueryModifyInstancePriceRequest {
	s.ModifyPrepayInstanceSpecRequest = v
	return s
}

func (s *QueryModifyInstancePriceRequest) Validate() error {
	if s.ModifyPrepayInstanceSpecRequest != nil {
		if err := s.ModifyPrepayInstanceSpecRequest.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequest struct {
	Ha *bool `json:"Ha,omitempty" xml:"Ha,omitempty"`
	// if can be null:
	// true
	HaResourceSpec *QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequestHaResourceSpec `json:"HaResourceSpec,omitempty" xml:"HaResourceSpec,omitempty" type:"Struct"`
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
	// sc_flinkserverless_public_cn-******
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// This parameter is required.
	ResourceSpec *QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequestResourceSpec `json:"ResourceSpec,omitempty" xml:"ResourceSpec,omitempty" type:"Struct"`
}

func (s QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequest) GoString() string {
	return s.String()
}

func (s *QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequest) GetHa() *bool {
	return s.Ha
}

func (s *QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequest) GetHaResourceSpec() *QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequestHaResourceSpec {
	return s.HaResourceSpec
}

func (s *QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequest) GetHaVSwitchIds() []*string {
	return s.HaVSwitchIds
}

func (s *QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequest) GetHaZoneId() *string {
	return s.HaZoneId
}

func (s *QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequest) GetRegion() *string {
	return s.Region
}

func (s *QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequest) GetResourceSpec() *QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequestResourceSpec {
	return s.ResourceSpec
}

func (s *QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequest) SetHa(v bool) *QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequest {
	s.Ha = &v
	return s
}

func (s *QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequest) SetHaResourceSpec(v *QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequestHaResourceSpec) *QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequest {
	s.HaResourceSpec = v
	return s
}

func (s *QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequest) SetHaVSwitchIds(v []*string) *QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequest {
	s.HaVSwitchIds = v
	return s
}

func (s *QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequest) SetHaZoneId(v string) *QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequest {
	s.HaZoneId = &v
	return s
}

func (s *QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequest) SetInstanceId(v string) *QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequest) SetRegion(v string) *QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequest {
	s.Region = &v
	return s
}

func (s *QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequest) SetResourceSpec(v *QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequestResourceSpec) *QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequest {
	s.ResourceSpec = v
	return s
}

func (s *QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequest) Validate() error {
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

type QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequestHaResourceSpec struct {
	Cpu      *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	MemoryGB *int32 `json:"MemoryGB,omitempty" xml:"MemoryGB,omitempty"`
}

func (s QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequestHaResourceSpec) String() string {
	return dara.Prettify(s)
}

func (s QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequestHaResourceSpec) GoString() string {
	return s.String()
}

func (s *QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequestHaResourceSpec) GetCpu() *int32 {
	return s.Cpu
}

func (s *QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequestHaResourceSpec) GetMemoryGB() *int32 {
	return s.MemoryGB
}

func (s *QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequestHaResourceSpec) SetCpu(v int32) *QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequestHaResourceSpec {
	s.Cpu = &v
	return s
}

func (s *QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequestHaResourceSpec) SetMemoryGB(v int32) *QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequestHaResourceSpec {
	s.MemoryGB = &v
	return s
}

func (s *QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequestHaResourceSpec) Validate() error {
	return dara.Validate(s)
}

type QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequestResourceSpec struct {
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

func (s QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequestResourceSpec) String() string {
	return dara.Prettify(s)
}

func (s QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequestResourceSpec) GoString() string {
	return s.String()
}

func (s *QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequestResourceSpec) GetCpu() *int32 {
	return s.Cpu
}

func (s *QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequestResourceSpec) GetMemoryGB() *int32 {
	return s.MemoryGB
}

func (s *QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequestResourceSpec) SetCpu(v int32) *QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequestResourceSpec {
	s.Cpu = &v
	return s
}

func (s *QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequestResourceSpec) SetMemoryGB(v int32) *QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequestResourceSpec {
	s.MemoryGB = &v
	return s
}

func (s *QueryModifyInstancePriceRequestModifyPrepayInstanceSpecRequestResourceSpec) Validate() error {
	return dara.Validate(s)
}
