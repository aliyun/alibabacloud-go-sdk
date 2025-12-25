// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCreateInstancePriceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArchitectureType(v string) *QueryCreateInstancePriceRequest
	GetArchitectureType() *string
	SetAutoRenew(v bool) *QueryCreateInstancePriceRequest
	GetAutoRenew() *bool
	SetChargeType(v string) *QueryCreateInstancePriceRequest
	GetChargeType() *string
	SetDuration(v int32) *QueryCreateInstancePriceRequest
	GetDuration() *int32
	SetExtra(v string) *QueryCreateInstancePriceRequest
	GetExtra() *string
	SetHa(v bool) *QueryCreateInstancePriceRequest
	GetHa() *bool
	SetHaResourceSpec(v *QueryCreateInstancePriceRequestHaResourceSpec) *QueryCreateInstancePriceRequest
	GetHaResourceSpec() *QueryCreateInstancePriceRequestHaResourceSpec
	SetInstanceName(v string) *QueryCreateInstancePriceRequest
	GetInstanceName() *string
	SetPricingCycle(v string) *QueryCreateInstancePriceRequest
	GetPricingCycle() *string
	SetPromotionCode(v string) *QueryCreateInstancePriceRequest
	GetPromotionCode() *string
	SetRegion(v string) *QueryCreateInstancePriceRequest
	GetRegion() *string
	SetResourceSpec(v *QueryCreateInstancePriceRequestResourceSpec) *QueryCreateInstancePriceRequest
	GetResourceSpec() *QueryCreateInstancePriceRequestResourceSpec
	SetStorage(v *QueryCreateInstancePriceRequestStorage) *QueryCreateInstancePriceRequest
	GetStorage() *QueryCreateInstancePriceRequestStorage
	SetUsePromotionCode(v bool) *QueryCreateInstancePriceRequest
	GetUsePromotionCode() *bool
	SetVSwitchIds(v []*string) *QueryCreateInstancePriceRequest
	GetVSwitchIds() []*string
	SetVpcId(v string) *QueryCreateInstancePriceRequest
	GetVpcId() *string
}

type QueryCreateInstancePriceRequest struct {
	ArchitectureType *string `json:"ArchitectureType,omitempty" xml:"ArchitectureType,omitempty"`
	// example:
	//
	// true
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PRE
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// example:
	//
	// 1
	Duration *int32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Extra    *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	Ha       *bool   `json:"Ha,omitempty" xml:"Ha,omitempty"`
	// if can be null:
	// true
	HaResourceSpec *QueryCreateInstancePriceRequestHaResourceSpec `json:"HaResourceSpec,omitempty" xml:"HaResourceSpec,omitempty" type:"Struct"`
	// example:
	//
	// rtc-e2e-test-post
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// Month
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// example:
	//
	// 500041860100636
	PromotionCode *string `json:"PromotionCode,omitempty" xml:"PromotionCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	Region           *string                                      `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceSpec     *QueryCreateInstancePriceRequestResourceSpec `json:"ResourceSpec,omitempty" xml:"ResourceSpec,omitempty" type:"Struct"`
	Storage          *QueryCreateInstancePriceRequestStorage      `json:"Storage,omitempty" xml:"Storage,omitempty" type:"Struct"`
	UsePromotionCode *bool                                        `json:"UsePromotionCode,omitempty" xml:"UsePromotionCode,omitempty"`
	VSwitchIds       []*string                                    `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	// example:
	//
	// vpc-2ze9xoh8qyt1rnxfmfcdi
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s QueryCreateInstancePriceRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryCreateInstancePriceRequest) GoString() string {
	return s.String()
}

func (s *QueryCreateInstancePriceRequest) GetArchitectureType() *string {
	return s.ArchitectureType
}

func (s *QueryCreateInstancePriceRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *QueryCreateInstancePriceRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *QueryCreateInstancePriceRequest) GetDuration() *int32 {
	return s.Duration
}

func (s *QueryCreateInstancePriceRequest) GetExtra() *string {
	return s.Extra
}

func (s *QueryCreateInstancePriceRequest) GetHa() *bool {
	return s.Ha
}

func (s *QueryCreateInstancePriceRequest) GetHaResourceSpec() *QueryCreateInstancePriceRequestHaResourceSpec {
	return s.HaResourceSpec
}

func (s *QueryCreateInstancePriceRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *QueryCreateInstancePriceRequest) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *QueryCreateInstancePriceRequest) GetPromotionCode() *string {
	return s.PromotionCode
}

func (s *QueryCreateInstancePriceRequest) GetRegion() *string {
	return s.Region
}

func (s *QueryCreateInstancePriceRequest) GetResourceSpec() *QueryCreateInstancePriceRequestResourceSpec {
	return s.ResourceSpec
}

func (s *QueryCreateInstancePriceRequest) GetStorage() *QueryCreateInstancePriceRequestStorage {
	return s.Storage
}

func (s *QueryCreateInstancePriceRequest) GetUsePromotionCode() *bool {
	return s.UsePromotionCode
}

func (s *QueryCreateInstancePriceRequest) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *QueryCreateInstancePriceRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *QueryCreateInstancePriceRequest) SetArchitectureType(v string) *QueryCreateInstancePriceRequest {
	s.ArchitectureType = &v
	return s
}

func (s *QueryCreateInstancePriceRequest) SetAutoRenew(v bool) *QueryCreateInstancePriceRequest {
	s.AutoRenew = &v
	return s
}

func (s *QueryCreateInstancePriceRequest) SetChargeType(v string) *QueryCreateInstancePriceRequest {
	s.ChargeType = &v
	return s
}

func (s *QueryCreateInstancePriceRequest) SetDuration(v int32) *QueryCreateInstancePriceRequest {
	s.Duration = &v
	return s
}

func (s *QueryCreateInstancePriceRequest) SetExtra(v string) *QueryCreateInstancePriceRequest {
	s.Extra = &v
	return s
}

func (s *QueryCreateInstancePriceRequest) SetHa(v bool) *QueryCreateInstancePriceRequest {
	s.Ha = &v
	return s
}

func (s *QueryCreateInstancePriceRequest) SetHaResourceSpec(v *QueryCreateInstancePriceRequestHaResourceSpec) *QueryCreateInstancePriceRequest {
	s.HaResourceSpec = v
	return s
}

func (s *QueryCreateInstancePriceRequest) SetInstanceName(v string) *QueryCreateInstancePriceRequest {
	s.InstanceName = &v
	return s
}

func (s *QueryCreateInstancePriceRequest) SetPricingCycle(v string) *QueryCreateInstancePriceRequest {
	s.PricingCycle = &v
	return s
}

func (s *QueryCreateInstancePriceRequest) SetPromotionCode(v string) *QueryCreateInstancePriceRequest {
	s.PromotionCode = &v
	return s
}

func (s *QueryCreateInstancePriceRequest) SetRegion(v string) *QueryCreateInstancePriceRequest {
	s.Region = &v
	return s
}

func (s *QueryCreateInstancePriceRequest) SetResourceSpec(v *QueryCreateInstancePriceRequestResourceSpec) *QueryCreateInstancePriceRequest {
	s.ResourceSpec = v
	return s
}

func (s *QueryCreateInstancePriceRequest) SetStorage(v *QueryCreateInstancePriceRequestStorage) *QueryCreateInstancePriceRequest {
	s.Storage = v
	return s
}

func (s *QueryCreateInstancePriceRequest) SetUsePromotionCode(v bool) *QueryCreateInstancePriceRequest {
	s.UsePromotionCode = &v
	return s
}

func (s *QueryCreateInstancePriceRequest) SetVSwitchIds(v []*string) *QueryCreateInstancePriceRequest {
	s.VSwitchIds = v
	return s
}

func (s *QueryCreateInstancePriceRequest) SetVpcId(v string) *QueryCreateInstancePriceRequest {
	s.VpcId = &v
	return s
}

func (s *QueryCreateInstancePriceRequest) Validate() error {
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
	if s.Storage != nil {
		if err := s.Storage.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryCreateInstancePriceRequestHaResourceSpec struct {
	// if can be null:
	// false
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// if can be null:
	// false
	MemoryGB *int32 `json:"MemoryGB,omitempty" xml:"MemoryGB,omitempty"`
}

func (s QueryCreateInstancePriceRequestHaResourceSpec) String() string {
	return dara.Prettify(s)
}

func (s QueryCreateInstancePriceRequestHaResourceSpec) GoString() string {
	return s.String()
}

func (s *QueryCreateInstancePriceRequestHaResourceSpec) GetCpu() *int32 {
	return s.Cpu
}

func (s *QueryCreateInstancePriceRequestHaResourceSpec) GetMemoryGB() *int32 {
	return s.MemoryGB
}

func (s *QueryCreateInstancePriceRequestHaResourceSpec) SetCpu(v int32) *QueryCreateInstancePriceRequestHaResourceSpec {
	s.Cpu = &v
	return s
}

func (s *QueryCreateInstancePriceRequestHaResourceSpec) SetMemoryGB(v int32) *QueryCreateInstancePriceRequestHaResourceSpec {
	s.MemoryGB = &v
	return s
}

func (s *QueryCreateInstancePriceRequestHaResourceSpec) Validate() error {
	return dara.Validate(s)
}

type QueryCreateInstancePriceRequestResourceSpec struct {
	// example:
	//
	// 4
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// example:
	//
	// 16
	MemoryGB *int32 `json:"MemoryGB,omitempty" xml:"MemoryGB,omitempty"`
}

func (s QueryCreateInstancePriceRequestResourceSpec) String() string {
	return dara.Prettify(s)
}

func (s QueryCreateInstancePriceRequestResourceSpec) GoString() string {
	return s.String()
}

func (s *QueryCreateInstancePriceRequestResourceSpec) GetCpu() *int32 {
	return s.Cpu
}

func (s *QueryCreateInstancePriceRequestResourceSpec) GetMemoryGB() *int32 {
	return s.MemoryGB
}

func (s *QueryCreateInstancePriceRequestResourceSpec) SetCpu(v int32) *QueryCreateInstancePriceRequestResourceSpec {
	s.Cpu = &v
	return s
}

func (s *QueryCreateInstancePriceRequestResourceSpec) SetMemoryGB(v int32) *QueryCreateInstancePriceRequestResourceSpec {
	s.MemoryGB = &v
	return s
}

func (s *QueryCreateInstancePriceRequestResourceSpec) Validate() error {
	return dara.Validate(s)
}

type QueryCreateInstancePriceRequestStorage struct {
	Oss *QueryCreateInstancePriceRequestStorageOss `json:"Oss,omitempty" xml:"Oss,omitempty" type:"Struct"`
}

func (s QueryCreateInstancePriceRequestStorage) String() string {
	return dara.Prettify(s)
}

func (s QueryCreateInstancePriceRequestStorage) GoString() string {
	return s.String()
}

func (s *QueryCreateInstancePriceRequestStorage) GetOss() *QueryCreateInstancePriceRequestStorageOss {
	return s.Oss
}

func (s *QueryCreateInstancePriceRequestStorage) SetOss(v *QueryCreateInstancePriceRequestStorageOss) *QueryCreateInstancePriceRequestStorage {
	s.Oss = v
	return s
}

func (s *QueryCreateInstancePriceRequestStorage) Validate() error {
	if s.Oss != nil {
		if err := s.Oss.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryCreateInstancePriceRequestStorageOss struct {
	// example:
	//
	// quicktracing
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
}

func (s QueryCreateInstancePriceRequestStorageOss) String() string {
	return dara.Prettify(s)
}

func (s QueryCreateInstancePriceRequestStorageOss) GoString() string {
	return s.String()
}

func (s *QueryCreateInstancePriceRequestStorageOss) GetBucket() *string {
	return s.Bucket
}

func (s *QueryCreateInstancePriceRequestStorageOss) SetBucket(v string) *QueryCreateInstancePriceRequestStorageOss {
	s.Bucket = &v
	return s
}

func (s *QueryCreateInstancePriceRequestStorageOss) Validate() error {
	return dara.Validate(s)
}
