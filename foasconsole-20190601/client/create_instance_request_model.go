// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateInstanceRequest(v *CreateInstanceRequestCreateInstanceRequest) *CreateInstanceRequest
	GetCreateInstanceRequest() *CreateInstanceRequestCreateInstanceRequest
}

type CreateInstanceRequest struct {
	// This parameter is required.
	CreateInstanceRequest *CreateInstanceRequestCreateInstanceRequest `json:"CreateInstanceRequest,omitempty" xml:"CreateInstanceRequest,omitempty" type:"Struct"`
}

func (s CreateInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequest) GetCreateInstanceRequest() *CreateInstanceRequestCreateInstanceRequest {
	return s.CreateInstanceRequest
}

func (s *CreateInstanceRequest) SetCreateInstanceRequest(v *CreateInstanceRequestCreateInstanceRequest) *CreateInstanceRequest {
	s.CreateInstanceRequest = v
	return s
}

func (s *CreateInstanceRequest) Validate() error {
	if s.CreateInstanceRequest != nil {
		if err := s.CreateInstanceRequest.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateInstanceRequestCreateInstanceRequest struct {
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
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// *
	Extra *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// if can be null:
	// true
	Ha *bool `json:"Ha,omitempty" xml:"Ha,omitempty"`
	// if can be null:
	// true
	HaResourceSpec *CreateInstanceRequestCreateInstanceRequestHaResourceSpec `json:"HaResourceSpec,omitempty" xml:"HaResourceSpec,omitempty" type:"Struct"`
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
	// vvp1
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	MonitorType  *string `json:"MonitorType,omitempty" xml:"MonitorType,omitempty"`
	// example:
	//
	// Month
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// example:
	//
	// 5RT******
	PromotionCode *string `json:"PromotionCode,omitempty" xml:"PromotionCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// example:
	//
	// rg-******
	ResourceGroupId *string                                                 `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceSpec    *CreateInstanceRequestCreateInstanceRequestResourceSpec `json:"ResourceSpec,omitempty" xml:"ResourceSpec,omitempty" type:"Struct"`
	// This parameter is required.
	Storage *CreateInstanceRequestCreateInstanceRequestStorage `json:"Storage,omitempty" xml:"Storage,omitempty" type:"Struct"`
	// example:
	//
	// true
	UsePromotionCode *bool `json:"UsePromotionCode,omitempty" xml:"UsePromotionCode,omitempty"`
	// This parameter is required.
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	// VPC ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-2ze9*******nxfmfcdi
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing-g
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateInstanceRequestCreateInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequestCreateInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestCreateInstanceRequest) GetArchitectureType() *string {
	return s.ArchitectureType
}

func (s *CreateInstanceRequestCreateInstanceRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateInstanceRequestCreateInstanceRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *CreateInstanceRequestCreateInstanceRequest) GetDuration() *int32 {
	return s.Duration
}

func (s *CreateInstanceRequestCreateInstanceRequest) GetExtra() *string {
	return s.Extra
}

func (s *CreateInstanceRequestCreateInstanceRequest) GetHa() *bool {
	return s.Ha
}

func (s *CreateInstanceRequestCreateInstanceRequest) GetHaResourceSpec() *CreateInstanceRequestCreateInstanceRequestHaResourceSpec {
	return s.HaResourceSpec
}

func (s *CreateInstanceRequestCreateInstanceRequest) GetHaVSwitchIds() []*string {
	return s.HaVSwitchIds
}

func (s *CreateInstanceRequestCreateInstanceRequest) GetHaZoneId() *string {
	return s.HaZoneId
}

func (s *CreateInstanceRequestCreateInstanceRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CreateInstanceRequestCreateInstanceRequest) GetMonitorType() *string {
	return s.MonitorType
}

func (s *CreateInstanceRequestCreateInstanceRequest) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *CreateInstanceRequestCreateInstanceRequest) GetPromotionCode() *string {
	return s.PromotionCode
}

func (s *CreateInstanceRequestCreateInstanceRequest) GetRegion() *string {
	return s.Region
}

func (s *CreateInstanceRequestCreateInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateInstanceRequestCreateInstanceRequest) GetResourceSpec() *CreateInstanceRequestCreateInstanceRequestResourceSpec {
	return s.ResourceSpec
}

func (s *CreateInstanceRequestCreateInstanceRequest) GetStorage() *CreateInstanceRequestCreateInstanceRequestStorage {
	return s.Storage
}

func (s *CreateInstanceRequestCreateInstanceRequest) GetUsePromotionCode() *bool {
	return s.UsePromotionCode
}

func (s *CreateInstanceRequestCreateInstanceRequest) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *CreateInstanceRequestCreateInstanceRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateInstanceRequestCreateInstanceRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateInstanceRequestCreateInstanceRequest) SetArchitectureType(v string) *CreateInstanceRequestCreateInstanceRequest {
	s.ArchitectureType = &v
	return s
}

func (s *CreateInstanceRequestCreateInstanceRequest) SetAutoRenew(v bool) *CreateInstanceRequestCreateInstanceRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateInstanceRequestCreateInstanceRequest) SetChargeType(v string) *CreateInstanceRequestCreateInstanceRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateInstanceRequestCreateInstanceRequest) SetDuration(v int32) *CreateInstanceRequestCreateInstanceRequest {
	s.Duration = &v
	return s
}

func (s *CreateInstanceRequestCreateInstanceRequest) SetExtra(v string) *CreateInstanceRequestCreateInstanceRequest {
	s.Extra = &v
	return s
}

func (s *CreateInstanceRequestCreateInstanceRequest) SetHa(v bool) *CreateInstanceRequestCreateInstanceRequest {
	s.Ha = &v
	return s
}

func (s *CreateInstanceRequestCreateInstanceRequest) SetHaResourceSpec(v *CreateInstanceRequestCreateInstanceRequestHaResourceSpec) *CreateInstanceRequestCreateInstanceRequest {
	s.HaResourceSpec = v
	return s
}

func (s *CreateInstanceRequestCreateInstanceRequest) SetHaVSwitchIds(v []*string) *CreateInstanceRequestCreateInstanceRequest {
	s.HaVSwitchIds = v
	return s
}

func (s *CreateInstanceRequestCreateInstanceRequest) SetHaZoneId(v string) *CreateInstanceRequestCreateInstanceRequest {
	s.HaZoneId = &v
	return s
}

func (s *CreateInstanceRequestCreateInstanceRequest) SetInstanceName(v string) *CreateInstanceRequestCreateInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateInstanceRequestCreateInstanceRequest) SetMonitorType(v string) *CreateInstanceRequestCreateInstanceRequest {
	s.MonitorType = &v
	return s
}

func (s *CreateInstanceRequestCreateInstanceRequest) SetPricingCycle(v string) *CreateInstanceRequestCreateInstanceRequest {
	s.PricingCycle = &v
	return s
}

func (s *CreateInstanceRequestCreateInstanceRequest) SetPromotionCode(v string) *CreateInstanceRequestCreateInstanceRequest {
	s.PromotionCode = &v
	return s
}

func (s *CreateInstanceRequestCreateInstanceRequest) SetRegion(v string) *CreateInstanceRequestCreateInstanceRequest {
	s.Region = &v
	return s
}

func (s *CreateInstanceRequestCreateInstanceRequest) SetResourceGroupId(v string) *CreateInstanceRequestCreateInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateInstanceRequestCreateInstanceRequest) SetResourceSpec(v *CreateInstanceRequestCreateInstanceRequestResourceSpec) *CreateInstanceRequestCreateInstanceRequest {
	s.ResourceSpec = v
	return s
}

func (s *CreateInstanceRequestCreateInstanceRequest) SetStorage(v *CreateInstanceRequestCreateInstanceRequestStorage) *CreateInstanceRequestCreateInstanceRequest {
	s.Storage = v
	return s
}

func (s *CreateInstanceRequestCreateInstanceRequest) SetUsePromotionCode(v bool) *CreateInstanceRequestCreateInstanceRequest {
	s.UsePromotionCode = &v
	return s
}

func (s *CreateInstanceRequestCreateInstanceRequest) SetVSwitchIds(v []*string) *CreateInstanceRequestCreateInstanceRequest {
	s.VSwitchIds = v
	return s
}

func (s *CreateInstanceRequestCreateInstanceRequest) SetVpcId(v string) *CreateInstanceRequestCreateInstanceRequest {
	s.VpcId = &v
	return s
}

func (s *CreateInstanceRequestCreateInstanceRequest) SetZoneId(v string) *CreateInstanceRequestCreateInstanceRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateInstanceRequestCreateInstanceRequest) Validate() error {
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

type CreateInstanceRequestCreateInstanceRequestHaResourceSpec struct {
	Cpu      *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	MemoryGB *int32 `json:"MemoryGB,omitempty" xml:"MemoryGB,omitempty"`
}

func (s CreateInstanceRequestCreateInstanceRequestHaResourceSpec) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequestCreateInstanceRequestHaResourceSpec) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestCreateInstanceRequestHaResourceSpec) GetCpu() *int32 {
	return s.Cpu
}

func (s *CreateInstanceRequestCreateInstanceRequestHaResourceSpec) GetMemoryGB() *int32 {
	return s.MemoryGB
}

func (s *CreateInstanceRequestCreateInstanceRequestHaResourceSpec) SetCpu(v int32) *CreateInstanceRequestCreateInstanceRequestHaResourceSpec {
	s.Cpu = &v
	return s
}

func (s *CreateInstanceRequestCreateInstanceRequestHaResourceSpec) SetMemoryGB(v int32) *CreateInstanceRequestCreateInstanceRequestHaResourceSpec {
	s.MemoryGB = &v
	return s
}

func (s *CreateInstanceRequestCreateInstanceRequestHaResourceSpec) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceRequestCreateInstanceRequestResourceSpec struct {
	// example:
	//
	// 10
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// example:
	//
	// 40
	MemoryGB *int32 `json:"MemoryGB,omitempty" xml:"MemoryGB,omitempty"`
}

func (s CreateInstanceRequestCreateInstanceRequestResourceSpec) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequestCreateInstanceRequestResourceSpec) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestCreateInstanceRequestResourceSpec) GetCpu() *int32 {
	return s.Cpu
}

func (s *CreateInstanceRequestCreateInstanceRequestResourceSpec) GetMemoryGB() *int32 {
	return s.MemoryGB
}

func (s *CreateInstanceRequestCreateInstanceRequestResourceSpec) SetCpu(v int32) *CreateInstanceRequestCreateInstanceRequestResourceSpec {
	s.Cpu = &v
	return s
}

func (s *CreateInstanceRequestCreateInstanceRequestResourceSpec) SetMemoryGB(v int32) *CreateInstanceRequestCreateInstanceRequestResourceSpec {
	s.MemoryGB = &v
	return s
}

func (s *CreateInstanceRequestCreateInstanceRequestResourceSpec) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceRequestCreateInstanceRequestStorage struct {
	// This parameter is required.
	Oss *CreateInstanceRequestCreateInstanceRequestStorageOss `json:"Oss,omitempty" xml:"Oss,omitempty" type:"Struct"`
}

func (s CreateInstanceRequestCreateInstanceRequestStorage) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequestCreateInstanceRequestStorage) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestCreateInstanceRequestStorage) GetOss() *CreateInstanceRequestCreateInstanceRequestStorageOss {
	return s.Oss
}

func (s *CreateInstanceRequestCreateInstanceRequestStorage) SetOss(v *CreateInstanceRequestCreateInstanceRequestStorageOss) *CreateInstanceRequestCreateInstanceRequestStorage {
	s.Oss = v
	return s
}

func (s *CreateInstanceRequestCreateInstanceRequestStorage) Validate() error {
	if s.Oss != nil {
		if err := s.Oss.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateInstanceRequestCreateInstanceRequestStorageOss struct {
	// This parameter is required.
	//
	// example:
	//
	// oss_flink
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
}

func (s CreateInstanceRequestCreateInstanceRequestStorageOss) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequestCreateInstanceRequestStorageOss) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestCreateInstanceRequestStorageOss) GetBucket() *string {
	return s.Bucket
}

func (s *CreateInstanceRequestCreateInstanceRequestStorageOss) SetBucket(v string) *CreateInstanceRequestCreateInstanceRequestStorageOss {
	s.Bucket = &v
	return s
}

func (s *CreateInstanceRequestCreateInstanceRequestStorageOss) Validate() error {
	return dara.Validate(s)
}
