// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryModifyInstancePriceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHa(v bool) *QueryModifyInstancePriceRequest
	GetHa() *bool
	SetHaResourceSpec(v *QueryModifyInstancePriceRequestHaResourceSpec) *QueryModifyInstancePriceRequest
	GetHaResourceSpec() *QueryModifyInstancePriceRequestHaResourceSpec
	SetHaVSwitchIds(v []*string) *QueryModifyInstancePriceRequest
	GetHaVSwitchIds() []*string
	SetInstanceId(v string) *QueryModifyInstancePriceRequest
	GetInstanceId() *string
	SetPromotionCode(v string) *QueryModifyInstancePriceRequest
	GetPromotionCode() *string
	SetRegion(v string) *QueryModifyInstancePriceRequest
	GetRegion() *string
	SetResourceSpec(v *QueryModifyInstancePriceRequestResourceSpec) *QueryModifyInstancePriceRequest
	GetResourceSpec() *QueryModifyInstancePriceRequestResourceSpec
	SetUsePromotionCode(v bool) *QueryModifyInstancePriceRequest
	GetUsePromotionCode() *bool
}

type QueryModifyInstancePriceRequest struct {
	Ha *bool `json:"Ha,omitempty" xml:"Ha,omitempty"`
	// if can be null:
	// true
	HaResourceSpec *QueryModifyInstancePriceRequestHaResourceSpec `json:"HaResourceSpec,omitempty" xml:"HaResourceSpec,omitempty" type:"Struct"`
	// if can be null:
	// true
	HaVSwitchIds []*string `json:"HaVSwitchIds,omitempty" xml:"HaVSwitchIds,omitempty" type:"Repeated"`
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
	ResourceSpec     *QueryModifyInstancePriceRequestResourceSpec `json:"ResourceSpec,omitempty" xml:"ResourceSpec,omitempty" type:"Struct"`
	UsePromotionCode *bool                                        `json:"UsePromotionCode,omitempty" xml:"UsePromotionCode,omitempty"`
}

func (s QueryModifyInstancePriceRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryModifyInstancePriceRequest) GoString() string {
	return s.String()
}

func (s *QueryModifyInstancePriceRequest) GetHa() *bool {
	return s.Ha
}

func (s *QueryModifyInstancePriceRequest) GetHaResourceSpec() *QueryModifyInstancePriceRequestHaResourceSpec {
	return s.HaResourceSpec
}

func (s *QueryModifyInstancePriceRequest) GetHaVSwitchIds() []*string {
	return s.HaVSwitchIds
}

func (s *QueryModifyInstancePriceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryModifyInstancePriceRequest) GetPromotionCode() *string {
	return s.PromotionCode
}

func (s *QueryModifyInstancePriceRequest) GetRegion() *string {
	return s.Region
}

func (s *QueryModifyInstancePriceRequest) GetResourceSpec() *QueryModifyInstancePriceRequestResourceSpec {
	return s.ResourceSpec
}

func (s *QueryModifyInstancePriceRequest) GetUsePromotionCode() *bool {
	return s.UsePromotionCode
}

func (s *QueryModifyInstancePriceRequest) SetHa(v bool) *QueryModifyInstancePriceRequest {
	s.Ha = &v
	return s
}

func (s *QueryModifyInstancePriceRequest) SetHaResourceSpec(v *QueryModifyInstancePriceRequestHaResourceSpec) *QueryModifyInstancePriceRequest {
	s.HaResourceSpec = v
	return s
}

func (s *QueryModifyInstancePriceRequest) SetHaVSwitchIds(v []*string) *QueryModifyInstancePriceRequest {
	s.HaVSwitchIds = v
	return s
}

func (s *QueryModifyInstancePriceRequest) SetInstanceId(v string) *QueryModifyInstancePriceRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryModifyInstancePriceRequest) SetPromotionCode(v string) *QueryModifyInstancePriceRequest {
	s.PromotionCode = &v
	return s
}

func (s *QueryModifyInstancePriceRequest) SetRegion(v string) *QueryModifyInstancePriceRequest {
	s.Region = &v
	return s
}

func (s *QueryModifyInstancePriceRequest) SetResourceSpec(v *QueryModifyInstancePriceRequestResourceSpec) *QueryModifyInstancePriceRequest {
	s.ResourceSpec = v
	return s
}

func (s *QueryModifyInstancePriceRequest) SetUsePromotionCode(v bool) *QueryModifyInstancePriceRequest {
	s.UsePromotionCode = &v
	return s
}

func (s *QueryModifyInstancePriceRequest) Validate() error {
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

type QueryModifyInstancePriceRequestHaResourceSpec struct {
	Cpu      *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	MemoryGB *int32 `json:"MemoryGB,omitempty" xml:"MemoryGB,omitempty"`
}

func (s QueryModifyInstancePriceRequestHaResourceSpec) String() string {
	return dara.Prettify(s)
}

func (s QueryModifyInstancePriceRequestHaResourceSpec) GoString() string {
	return s.String()
}

func (s *QueryModifyInstancePriceRequestHaResourceSpec) GetCpu() *int32 {
	return s.Cpu
}

func (s *QueryModifyInstancePriceRequestHaResourceSpec) GetMemoryGB() *int32 {
	return s.MemoryGB
}

func (s *QueryModifyInstancePriceRequestHaResourceSpec) SetCpu(v int32) *QueryModifyInstancePriceRequestHaResourceSpec {
	s.Cpu = &v
	return s
}

func (s *QueryModifyInstancePriceRequestHaResourceSpec) SetMemoryGB(v int32) *QueryModifyInstancePriceRequestHaResourceSpec {
	s.MemoryGB = &v
	return s
}

func (s *QueryModifyInstancePriceRequestHaResourceSpec) Validate() error {
	return dara.Validate(s)
}

type QueryModifyInstancePriceRequestResourceSpec struct {
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

func (s QueryModifyInstancePriceRequestResourceSpec) String() string {
	return dara.Prettify(s)
}

func (s QueryModifyInstancePriceRequestResourceSpec) GoString() string {
	return s.String()
}

func (s *QueryModifyInstancePriceRequestResourceSpec) GetCpu() *int32 {
	return s.Cpu
}

func (s *QueryModifyInstancePriceRequestResourceSpec) GetMemoryGB() *int32 {
	return s.MemoryGB
}

func (s *QueryModifyInstancePriceRequestResourceSpec) SetCpu(v int32) *QueryModifyInstancePriceRequestResourceSpec {
	s.Cpu = &v
	return s
}

func (s *QueryModifyInstancePriceRequestResourceSpec) SetMemoryGB(v int32) *QueryModifyInstancePriceRequestResourceSpec {
	s.MemoryGB = &v
	return s
}

func (s *QueryModifyInstancePriceRequestResourceSpec) Validate() error {
	return dara.Validate(s)
}
