// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCreateInstancePriceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateInstanceRequest(v *QueryCreateInstancePriceRequestCreateInstanceRequest) *QueryCreateInstancePriceRequest
	GetCreateInstanceRequest() *QueryCreateInstancePriceRequestCreateInstanceRequest
}

type QueryCreateInstancePriceRequest struct {
	// This parameter is required.
	CreateInstanceRequest *QueryCreateInstancePriceRequestCreateInstanceRequest `json:"CreateInstanceRequest,omitempty" xml:"CreateInstanceRequest,omitempty" type:"Struct"`
}

func (s QueryCreateInstancePriceRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryCreateInstancePriceRequest) GoString() string {
	return s.String()
}

func (s *QueryCreateInstancePriceRequest) GetCreateInstanceRequest() *QueryCreateInstancePriceRequestCreateInstanceRequest {
	return s.CreateInstanceRequest
}

func (s *QueryCreateInstancePriceRequest) SetCreateInstanceRequest(v *QueryCreateInstancePriceRequestCreateInstanceRequest) *QueryCreateInstancePriceRequest {
	s.CreateInstanceRequest = v
	return s
}

func (s *QueryCreateInstancePriceRequest) Validate() error {
	if s.CreateInstanceRequest != nil {
		if err := s.CreateInstanceRequest.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryCreateInstancePriceRequestCreateInstanceRequest struct {
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
	// if can be null:
	// false
	Ha *bool `json:"Ha,omitempty" xml:"Ha,omitempty"`
	// if can be null:
	// true
	HaResourceSpec *QueryCreateInstancePriceRequestCreateInstanceRequestHaResourceSpec `json:"HaResourceSpec,omitempty" xml:"HaResourceSpec,omitempty" type:"Struct"`
	// example:
	//
	// vvp1
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// Month
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// example:
	//
	// 5ef***
	PromotionCode *string `json:"PromotionCode,omitempty" xml:"PromotionCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	Region       *string                                                           `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceSpec *QueryCreateInstancePriceRequestCreateInstanceRequestResourceSpec `json:"ResourceSpec,omitempty" xml:"ResourceSpec,omitempty" type:"Struct"`
	Storage      *QueryCreateInstancePriceRequestCreateInstanceRequestStorage      `json:"Storage,omitempty" xml:"Storage,omitempty" type:"Struct"`
	// example:
	//
	// true
	UsePromotionCode *bool     `json:"UsePromotionCode,omitempty" xml:"UsePromotionCode,omitempty"`
	VSwitchIds       []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	// VPC ID。
	//
	// example:
	//
	// vpc-2ze9*******nxfmfcdi
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// example:
	//
	// cn-beijing-g
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s QueryCreateInstancePriceRequestCreateInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryCreateInstancePriceRequestCreateInstanceRequest) GoString() string {
	return s.String()
}

func (s *QueryCreateInstancePriceRequestCreateInstanceRequest) GetArchitectureType() *string {
	return s.ArchitectureType
}

func (s *QueryCreateInstancePriceRequestCreateInstanceRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *QueryCreateInstancePriceRequestCreateInstanceRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *QueryCreateInstancePriceRequestCreateInstanceRequest) GetDuration() *int32 {
	return s.Duration
}

func (s *QueryCreateInstancePriceRequestCreateInstanceRequest) GetExtra() *string {
	return s.Extra
}

func (s *QueryCreateInstancePriceRequestCreateInstanceRequest) GetHa() *bool {
	return s.Ha
}

func (s *QueryCreateInstancePriceRequestCreateInstanceRequest) GetHaResourceSpec() *QueryCreateInstancePriceRequestCreateInstanceRequestHaResourceSpec {
	return s.HaResourceSpec
}

func (s *QueryCreateInstancePriceRequestCreateInstanceRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *QueryCreateInstancePriceRequestCreateInstanceRequest) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *QueryCreateInstancePriceRequestCreateInstanceRequest) GetPromotionCode() *string {
	return s.PromotionCode
}

func (s *QueryCreateInstancePriceRequestCreateInstanceRequest) GetRegion() *string {
	return s.Region
}

func (s *QueryCreateInstancePriceRequestCreateInstanceRequest) GetResourceSpec() *QueryCreateInstancePriceRequestCreateInstanceRequestResourceSpec {
	return s.ResourceSpec
}

func (s *QueryCreateInstancePriceRequestCreateInstanceRequest) GetStorage() *QueryCreateInstancePriceRequestCreateInstanceRequestStorage {
	return s.Storage
}

func (s *QueryCreateInstancePriceRequestCreateInstanceRequest) GetUsePromotionCode() *bool {
	return s.UsePromotionCode
}

func (s *QueryCreateInstancePriceRequestCreateInstanceRequest) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *QueryCreateInstancePriceRequestCreateInstanceRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *QueryCreateInstancePriceRequestCreateInstanceRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *QueryCreateInstancePriceRequestCreateInstanceRequest) SetArchitectureType(v string) *QueryCreateInstancePriceRequestCreateInstanceRequest {
	s.ArchitectureType = &v
	return s
}

func (s *QueryCreateInstancePriceRequestCreateInstanceRequest) SetAutoRenew(v bool) *QueryCreateInstancePriceRequestCreateInstanceRequest {
	s.AutoRenew = &v
	return s
}

func (s *QueryCreateInstancePriceRequestCreateInstanceRequest) SetChargeType(v string) *QueryCreateInstancePriceRequestCreateInstanceRequest {
	s.ChargeType = &v
	return s
}

func (s *QueryCreateInstancePriceRequestCreateInstanceRequest) SetDuration(v int32) *QueryCreateInstancePriceRequestCreateInstanceRequest {
	s.Duration = &v
	return s
}

func (s *QueryCreateInstancePriceRequestCreateInstanceRequest) SetExtra(v string) *QueryCreateInstancePriceRequestCreateInstanceRequest {
	s.Extra = &v
	return s
}

func (s *QueryCreateInstancePriceRequestCreateInstanceRequest) SetHa(v bool) *QueryCreateInstancePriceRequestCreateInstanceRequest {
	s.Ha = &v
	return s
}

func (s *QueryCreateInstancePriceRequestCreateInstanceRequest) SetHaResourceSpec(v *QueryCreateInstancePriceRequestCreateInstanceRequestHaResourceSpec) *QueryCreateInstancePriceRequestCreateInstanceRequest {
	s.HaResourceSpec = v
	return s
}

func (s *QueryCreateInstancePriceRequestCreateInstanceRequest) SetInstanceName(v string) *QueryCreateInstancePriceRequestCreateInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *QueryCreateInstancePriceRequestCreateInstanceRequest) SetPricingCycle(v string) *QueryCreateInstancePriceRequestCreateInstanceRequest {
	s.PricingCycle = &v
	return s
}

func (s *QueryCreateInstancePriceRequestCreateInstanceRequest) SetPromotionCode(v string) *QueryCreateInstancePriceRequestCreateInstanceRequest {
	s.PromotionCode = &v
	return s
}

func (s *QueryCreateInstancePriceRequestCreateInstanceRequest) SetRegion(v string) *QueryCreateInstancePriceRequestCreateInstanceRequest {
	s.Region = &v
	return s
}

func (s *QueryCreateInstancePriceRequestCreateInstanceRequest) SetResourceSpec(v *QueryCreateInstancePriceRequestCreateInstanceRequestResourceSpec) *QueryCreateInstancePriceRequestCreateInstanceRequest {
	s.ResourceSpec = v
	return s
}

func (s *QueryCreateInstancePriceRequestCreateInstanceRequest) SetStorage(v *QueryCreateInstancePriceRequestCreateInstanceRequestStorage) *QueryCreateInstancePriceRequestCreateInstanceRequest {
	s.Storage = v
	return s
}

func (s *QueryCreateInstancePriceRequestCreateInstanceRequest) SetUsePromotionCode(v bool) *QueryCreateInstancePriceRequestCreateInstanceRequest {
	s.UsePromotionCode = &v
	return s
}

func (s *QueryCreateInstancePriceRequestCreateInstanceRequest) SetVSwitchIds(v []*string) *QueryCreateInstancePriceRequestCreateInstanceRequest {
	s.VSwitchIds = v
	return s
}

func (s *QueryCreateInstancePriceRequestCreateInstanceRequest) SetVpcId(v string) *QueryCreateInstancePriceRequestCreateInstanceRequest {
	s.VpcId = &v
	return s
}

func (s *QueryCreateInstancePriceRequestCreateInstanceRequest) SetZoneId(v string) *QueryCreateInstancePriceRequestCreateInstanceRequest {
	s.ZoneId = &v
	return s
}

func (s *QueryCreateInstancePriceRequestCreateInstanceRequest) Validate() error {
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

type QueryCreateInstancePriceRequestCreateInstanceRequestHaResourceSpec struct {
	Cpu      *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	MemoryGB *int32 `json:"MemoryGB,omitempty" xml:"MemoryGB,omitempty"`
}

func (s QueryCreateInstancePriceRequestCreateInstanceRequestHaResourceSpec) String() string {
	return dara.Prettify(s)
}

func (s QueryCreateInstancePriceRequestCreateInstanceRequestHaResourceSpec) GoString() string {
	return s.String()
}

func (s *QueryCreateInstancePriceRequestCreateInstanceRequestHaResourceSpec) GetCpu() *int32 {
	return s.Cpu
}

func (s *QueryCreateInstancePriceRequestCreateInstanceRequestHaResourceSpec) GetMemoryGB() *int32 {
	return s.MemoryGB
}

func (s *QueryCreateInstancePriceRequestCreateInstanceRequestHaResourceSpec) SetCpu(v int32) *QueryCreateInstancePriceRequestCreateInstanceRequestHaResourceSpec {
	s.Cpu = &v
	return s
}

func (s *QueryCreateInstancePriceRequestCreateInstanceRequestHaResourceSpec) SetMemoryGB(v int32) *QueryCreateInstancePriceRequestCreateInstanceRequestHaResourceSpec {
	s.MemoryGB = &v
	return s
}

func (s *QueryCreateInstancePriceRequestCreateInstanceRequestHaResourceSpec) Validate() error {
	return dara.Validate(s)
}

type QueryCreateInstancePriceRequestCreateInstanceRequestResourceSpec struct {
	// example:
	//
	// 10
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// example:
	//
	// 40
	MemoryGB *int32 `json:"MemoryGB,omitempty" xml:"MemoryGB,omitempty"`
}

func (s QueryCreateInstancePriceRequestCreateInstanceRequestResourceSpec) String() string {
	return dara.Prettify(s)
}

func (s QueryCreateInstancePriceRequestCreateInstanceRequestResourceSpec) GoString() string {
	return s.String()
}

func (s *QueryCreateInstancePriceRequestCreateInstanceRequestResourceSpec) GetCpu() *int32 {
	return s.Cpu
}

func (s *QueryCreateInstancePriceRequestCreateInstanceRequestResourceSpec) GetMemoryGB() *int32 {
	return s.MemoryGB
}

func (s *QueryCreateInstancePriceRequestCreateInstanceRequestResourceSpec) SetCpu(v int32) *QueryCreateInstancePriceRequestCreateInstanceRequestResourceSpec {
	s.Cpu = &v
	return s
}

func (s *QueryCreateInstancePriceRequestCreateInstanceRequestResourceSpec) SetMemoryGB(v int32) *QueryCreateInstancePriceRequestCreateInstanceRequestResourceSpec {
	s.MemoryGB = &v
	return s
}

func (s *QueryCreateInstancePriceRequestCreateInstanceRequestResourceSpec) Validate() error {
	return dara.Validate(s)
}

type QueryCreateInstancePriceRequestCreateInstanceRequestStorage struct {
	Oss *QueryCreateInstancePriceRequestCreateInstanceRequestStorageOss `json:"Oss,omitempty" xml:"Oss,omitempty" type:"Struct"`
}

func (s QueryCreateInstancePriceRequestCreateInstanceRequestStorage) String() string {
	return dara.Prettify(s)
}

func (s QueryCreateInstancePriceRequestCreateInstanceRequestStorage) GoString() string {
	return s.String()
}

func (s *QueryCreateInstancePriceRequestCreateInstanceRequestStorage) GetOss() *QueryCreateInstancePriceRequestCreateInstanceRequestStorageOss {
	return s.Oss
}

func (s *QueryCreateInstancePriceRequestCreateInstanceRequestStorage) SetOss(v *QueryCreateInstancePriceRequestCreateInstanceRequestStorageOss) *QueryCreateInstancePriceRequestCreateInstanceRequestStorage {
	s.Oss = v
	return s
}

func (s *QueryCreateInstancePriceRequestCreateInstanceRequestStorage) Validate() error {
	if s.Oss != nil {
		if err := s.Oss.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryCreateInstancePriceRequestCreateInstanceRequestStorageOss struct {
	// example:
	//
	// oss_flink
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
}

func (s QueryCreateInstancePriceRequestCreateInstanceRequestStorageOss) String() string {
	return dara.Prettify(s)
}

func (s QueryCreateInstancePriceRequestCreateInstanceRequestStorageOss) GoString() string {
	return s.String()
}

func (s *QueryCreateInstancePriceRequestCreateInstanceRequestStorageOss) GetBucket() *string {
	return s.Bucket
}

func (s *QueryCreateInstancePriceRequestCreateInstanceRequestStorageOss) SetBucket(v string) *QueryCreateInstancePriceRequestCreateInstanceRequestStorageOss {
	s.Bucket = &v
	return s
}

func (s *QueryCreateInstancePriceRequestCreateInstanceRequestStorageOss) Validate() error {
	return dara.Validate(s)
}
