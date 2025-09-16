// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArchitectureType(v string) *CreateInstanceRequest
	GetArchitectureType() *string
	SetAutoRenew(v bool) *CreateInstanceRequest
	GetAutoRenew() *bool
	SetChargeType(v string) *CreateInstanceRequest
	GetChargeType() *string
	SetDuration(v int32) *CreateInstanceRequest
	GetDuration() *int32
	SetExtra(v string) *CreateInstanceRequest
	GetExtra() *string
	SetHa(v bool) *CreateInstanceRequest
	GetHa() *bool
	SetHaResourceSpec(v *CreateInstanceRequestHaResourceSpec) *CreateInstanceRequest
	GetHaResourceSpec() *CreateInstanceRequestHaResourceSpec
	SetHaVSwitchIds(v []*string) *CreateInstanceRequest
	GetHaVSwitchIds() []*string
	SetInstanceName(v string) *CreateInstanceRequest
	GetInstanceName() *string
	SetMonitorType(v string) *CreateInstanceRequest
	GetMonitorType() *string
	SetPricingCycle(v string) *CreateInstanceRequest
	GetPricingCycle() *string
	SetPromotionCode(v string) *CreateInstanceRequest
	GetPromotionCode() *string
	SetRegion(v string) *CreateInstanceRequest
	GetRegion() *string
	SetResourceGroupId(v string) *CreateInstanceRequest
	GetResourceGroupId() *string
	SetResourceSpec(v *CreateInstanceRequestResourceSpec) *CreateInstanceRequest
	GetResourceSpec() *CreateInstanceRequestResourceSpec
	SetStorage(v *CreateInstanceRequestStorage) *CreateInstanceRequest
	GetStorage() *CreateInstanceRequestStorage
	SetTag(v []*CreateInstanceRequestTag) *CreateInstanceRequest
	GetTag() []*CreateInstanceRequestTag
	SetUsePromotionCode(v bool) *CreateInstanceRequest
	GetUsePromotionCode() *bool
	SetVSwitchIds(v []*string) *CreateInstanceRequest
	GetVSwitchIds() []*string
	SetVpcId(v string) *CreateInstanceRequest
	GetVpcId() *string
}

type CreateInstanceRequest struct {
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
	// true
	Ha *bool `json:"Ha,omitempty" xml:"Ha,omitempty"`
	// if can be null:
	// true
	HaResourceSpec *CreateInstanceRequestHaResourceSpec `json:"HaResourceSpec,omitempty" xml:"HaResourceSpec,omitempty" type:"Struct"`
	// if can be null:
	// true
	HaVSwitchIds []*string `json:"HaVSwitchIds,omitempty" xml:"HaVSwitchIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// rtc-e2e-test-pre
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	MonitorType  *string `json:"MonitorType,omitempty" xml:"MonitorType,omitempty"`
	// example:
	//
	// Month
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// example:
	//
	// 500043499350689
	PromotionCode *string `json:"PromotionCode,omitempty" xml:"PromotionCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	Region          *string                            `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceGroupId *string                            `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceSpec    *CreateInstanceRequestResourceSpec `json:"ResourceSpec,omitempty" xml:"ResourceSpec,omitempty" type:"Struct"`
	// This parameter is required.
	Storage          *CreateInstanceRequestStorage `json:"Storage,omitempty" xml:"Storage,omitempty" type:"Struct"`
	Tag              []*CreateInstanceRequestTag   `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	UsePromotionCode *bool                         `json:"UsePromotionCode,omitempty" xml:"UsePromotionCode,omitempty"`
	// This parameter is required.
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// vpc-2ze9xoh8qyt1rnxfmfcdi
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequest) GetArchitectureType() *string {
	return s.ArchitectureType
}

func (s *CreateInstanceRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateInstanceRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *CreateInstanceRequest) GetDuration() *int32 {
	return s.Duration
}

func (s *CreateInstanceRequest) GetExtra() *string {
	return s.Extra
}

func (s *CreateInstanceRequest) GetHa() *bool {
	return s.Ha
}

func (s *CreateInstanceRequest) GetHaResourceSpec() *CreateInstanceRequestHaResourceSpec {
	return s.HaResourceSpec
}

func (s *CreateInstanceRequest) GetHaVSwitchIds() []*string {
	return s.HaVSwitchIds
}

func (s *CreateInstanceRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CreateInstanceRequest) GetMonitorType() *string {
	return s.MonitorType
}

func (s *CreateInstanceRequest) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *CreateInstanceRequest) GetPromotionCode() *string {
	return s.PromotionCode
}

func (s *CreateInstanceRequest) GetRegion() *string {
	return s.Region
}

func (s *CreateInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateInstanceRequest) GetResourceSpec() *CreateInstanceRequestResourceSpec {
	return s.ResourceSpec
}

func (s *CreateInstanceRequest) GetStorage() *CreateInstanceRequestStorage {
	return s.Storage
}

func (s *CreateInstanceRequest) GetTag() []*CreateInstanceRequestTag {
	return s.Tag
}

func (s *CreateInstanceRequest) GetUsePromotionCode() *bool {
	return s.UsePromotionCode
}

func (s *CreateInstanceRequest) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *CreateInstanceRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateInstanceRequest) SetArchitectureType(v string) *CreateInstanceRequest {
	s.ArchitectureType = &v
	return s
}

func (s *CreateInstanceRequest) SetAutoRenew(v bool) *CreateInstanceRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateInstanceRequest) SetChargeType(v string) *CreateInstanceRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateInstanceRequest) SetDuration(v int32) *CreateInstanceRequest {
	s.Duration = &v
	return s
}

func (s *CreateInstanceRequest) SetExtra(v string) *CreateInstanceRequest {
	s.Extra = &v
	return s
}

func (s *CreateInstanceRequest) SetHa(v bool) *CreateInstanceRequest {
	s.Ha = &v
	return s
}

func (s *CreateInstanceRequest) SetHaResourceSpec(v *CreateInstanceRequestHaResourceSpec) *CreateInstanceRequest {
	s.HaResourceSpec = v
	return s
}

func (s *CreateInstanceRequest) SetHaVSwitchIds(v []*string) *CreateInstanceRequest {
	s.HaVSwitchIds = v
	return s
}

func (s *CreateInstanceRequest) SetInstanceName(v string) *CreateInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateInstanceRequest) SetMonitorType(v string) *CreateInstanceRequest {
	s.MonitorType = &v
	return s
}

func (s *CreateInstanceRequest) SetPricingCycle(v string) *CreateInstanceRequest {
	s.PricingCycle = &v
	return s
}

func (s *CreateInstanceRequest) SetPromotionCode(v string) *CreateInstanceRequest {
	s.PromotionCode = &v
	return s
}

func (s *CreateInstanceRequest) SetRegion(v string) *CreateInstanceRequest {
	s.Region = &v
	return s
}

func (s *CreateInstanceRequest) SetResourceGroupId(v string) *CreateInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateInstanceRequest) SetResourceSpec(v *CreateInstanceRequestResourceSpec) *CreateInstanceRequest {
	s.ResourceSpec = v
	return s
}

func (s *CreateInstanceRequest) SetStorage(v *CreateInstanceRequestStorage) *CreateInstanceRequest {
	s.Storage = v
	return s
}

func (s *CreateInstanceRequest) SetTag(v []*CreateInstanceRequestTag) *CreateInstanceRequest {
	s.Tag = v
	return s
}

func (s *CreateInstanceRequest) SetUsePromotionCode(v bool) *CreateInstanceRequest {
	s.UsePromotionCode = &v
	return s
}

func (s *CreateInstanceRequest) SetVSwitchIds(v []*string) *CreateInstanceRequest {
	s.VSwitchIds = v
	return s
}

func (s *CreateInstanceRequest) SetVpcId(v string) *CreateInstanceRequest {
	s.VpcId = &v
	return s
}

func (s *CreateInstanceRequest) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceRequestHaResourceSpec struct {
	Cpu      *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	MemoryGB *int32 `json:"MemoryGB,omitempty" xml:"MemoryGB,omitempty"`
}

func (s CreateInstanceRequestHaResourceSpec) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequestHaResourceSpec) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestHaResourceSpec) GetCpu() *int32 {
	return s.Cpu
}

func (s *CreateInstanceRequestHaResourceSpec) GetMemoryGB() *int32 {
	return s.MemoryGB
}

func (s *CreateInstanceRequestHaResourceSpec) SetCpu(v int32) *CreateInstanceRequestHaResourceSpec {
	s.Cpu = &v
	return s
}

func (s *CreateInstanceRequestHaResourceSpec) SetMemoryGB(v int32) *CreateInstanceRequestHaResourceSpec {
	s.MemoryGB = &v
	return s
}

func (s *CreateInstanceRequestHaResourceSpec) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceRequestResourceSpec struct {
	// example:
	//
	// 30
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// example:
	//
	// 120
	MemoryGB *int32 `json:"MemoryGB,omitempty" xml:"MemoryGB,omitempty"`
}

func (s CreateInstanceRequestResourceSpec) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequestResourceSpec) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestResourceSpec) GetCpu() *int32 {
	return s.Cpu
}

func (s *CreateInstanceRequestResourceSpec) GetMemoryGB() *int32 {
	return s.MemoryGB
}

func (s *CreateInstanceRequestResourceSpec) SetCpu(v int32) *CreateInstanceRequestResourceSpec {
	s.Cpu = &v
	return s
}

func (s *CreateInstanceRequestResourceSpec) SetMemoryGB(v int32) *CreateInstanceRequestResourceSpec {
	s.MemoryGB = &v
	return s
}

func (s *CreateInstanceRequestResourceSpec) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceRequestStorage struct {
	FullyManaged *bool                            `json:"FullyManaged,omitempty" xml:"FullyManaged,omitempty"`
	Oss          *CreateInstanceRequestStorageOss `json:"Oss,omitempty" xml:"Oss,omitempty" type:"Struct"`
}

func (s CreateInstanceRequestStorage) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequestStorage) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestStorage) GetFullyManaged() *bool {
	return s.FullyManaged
}

func (s *CreateInstanceRequestStorage) GetOss() *CreateInstanceRequestStorageOss {
	return s.Oss
}

func (s *CreateInstanceRequestStorage) SetFullyManaged(v bool) *CreateInstanceRequestStorage {
	s.FullyManaged = &v
	return s
}

func (s *CreateInstanceRequestStorage) SetOss(v *CreateInstanceRequestStorageOss) *CreateInstanceRequestStorage {
	s.Oss = v
	return s
}

func (s *CreateInstanceRequestStorage) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceRequestStorageOss struct {
	// example:
	//
	// oss-flink-cn-shanghai-260343971602724445
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
}

func (s CreateInstanceRequestStorageOss) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequestStorageOss) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestStorageOss) GetBucket() *string {
	return s.Bucket
}

func (s *CreateInstanceRequestStorageOss) SetBucket(v string) *CreateInstanceRequestStorageOss {
	s.Bucket = &v
	return s
}

func (s *CreateInstanceRequestStorageOss) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateInstanceRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequestTag) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateInstanceRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateInstanceRequestTag) SetKey(v string) *CreateInstanceRequestTag {
	s.Key = &v
	return s
}

func (s *CreateInstanceRequestTag) SetValue(v string) *CreateInstanceRequestTag {
	s.Value = &v
	return s
}

func (s *CreateInstanceRequestTag) Validate() error {
	return dara.Validate(s)
}
