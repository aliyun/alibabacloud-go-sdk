// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTairKVCacheCustomInstanceAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetArchitectureType(v string) *DescribeTairKVCacheCustomInstanceAttributeResponseBody
	GetArchitectureType() *string
	SetChargeType(v string) *DescribeTairKVCacheCustomInstanceAttributeResponseBody
	GetChargeType() *string
	SetCpu(v int64) *DescribeTairKVCacheCustomInstanceAttributeResponseBody
	GetCpu() *int64
	SetCreateTime(v string) *DescribeTairKVCacheCustomInstanceAttributeResponseBody
	GetCreateTime() *string
	SetDisks(v *DescribeTairKVCacheCustomInstanceAttributeResponseBodyDisks) *DescribeTairKVCacheCustomInstanceAttributeResponseBody
	GetDisks() *DescribeTairKVCacheCustomInstanceAttributeResponseBodyDisks
	SetEndTime(v string) *DescribeTairKVCacheCustomInstanceAttributeResponseBody
	GetEndTime() *string
	SetImageId(v string) *DescribeTairKVCacheCustomInstanceAttributeResponseBody
	GetImageId() *string
	SetInstanceClass(v string) *DescribeTairKVCacheCustomInstanceAttributeResponseBody
	GetInstanceClass() *string
	SetInstanceId(v string) *DescribeTairKVCacheCustomInstanceAttributeResponseBody
	GetInstanceId() *string
	SetInstanceName(v string) *DescribeTairKVCacheCustomInstanceAttributeResponseBody
	GetInstanceName() *string
	SetInstanceStatus(v string) *DescribeTairKVCacheCustomInstanceAttributeResponseBody
	GetInstanceStatus() *string
	SetInstanceType(v string) *DescribeTairKVCacheCustomInstanceAttributeResponseBody
	GetInstanceType() *string
	SetIsOrderCompleted(v bool) *DescribeTairKVCacheCustomInstanceAttributeResponseBody
	GetIsOrderCompleted() *bool
	SetMemory(v int64) *DescribeTairKVCacheCustomInstanceAttributeResponseBody
	GetMemory() *int64
	SetNetworkType(v string) *DescribeTairKVCacheCustomInstanceAttributeResponseBody
	GetNetworkType() *string
	SetPrivateIp(v string) *DescribeTairKVCacheCustomInstanceAttributeResponseBody
	GetPrivateIp() *string
	SetRegionId(v string) *DescribeTairKVCacheCustomInstanceAttributeResponseBody
	GetRegionId() *string
	SetRequestId(v string) *DescribeTairKVCacheCustomInstanceAttributeResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *DescribeTairKVCacheCustomInstanceAttributeResponseBody
	GetResourceGroupId() *string
	SetSecurityGroupId(v string) *DescribeTairKVCacheCustomInstanceAttributeResponseBody
	GetSecurityGroupId() *string
	SetStorage(v int64) *DescribeTairKVCacheCustomInstanceAttributeResponseBody
	GetStorage() *int64
	SetStorageType(v string) *DescribeTairKVCacheCustomInstanceAttributeResponseBody
	GetStorageType() *string
	SetTags(v *DescribeTairKVCacheCustomInstanceAttributeResponseBodyTags) *DescribeTairKVCacheCustomInstanceAttributeResponseBody
	GetTags() *DescribeTairKVCacheCustomInstanceAttributeResponseBodyTags
	SetUseEni(v bool) *DescribeTairKVCacheCustomInstanceAttributeResponseBody
	GetUseEni() *bool
	SetVSwitchId(v string) *DescribeTairKVCacheCustomInstanceAttributeResponseBody
	GetVSwitchId() *string
	SetVpcId(v string) *DescribeTairKVCacheCustomInstanceAttributeResponseBody
	GetVpcId() *string
	SetZoneId(v string) *DescribeTairKVCacheCustomInstanceAttributeResponseBody
	GetZoneId() *string
	SetZoneType(v string) *DescribeTairKVCacheCustomInstanceAttributeResponseBody
	GetZoneType() *string
}

type DescribeTairKVCacheCustomInstanceAttributeResponseBody struct {
	// example:
	//
	// tair_custom
	ArchitectureType *string `json:"ArchitectureType,omitempty" xml:"ArchitectureType,omitempty"`
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// example:
	//
	// 2
	Cpu *int64 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// example:
	//
	// 2024-02-21T08:23Z
	CreateTime *string                                                      `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Disks      *DescribeTairKVCacheCustomInstanceAttributeResponseBodyDisks `json:"Disks,omitempty" xml:"Disks,omitempty" type:"Struct"`
	// example:
	//
	// 2024-05-28T00:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// m-bp10k5694v6yfevajw**
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// example:
	//
	// tair.gpu.test.16g
	InstanceClass *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	// example:
	//
	// tc-bp1zxszhcgatnx****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// newinstancename
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// Normal
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// example:
	//
	// TairCustom
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// example:
	//
	// true
	IsOrderCompleted *bool `json:"IsOrderCompleted,omitempty" xml:"IsOrderCompleted,omitempty"`
	// example:
	//
	// 262144
	Memory *int64 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// example:
	//
	// VPC
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// example:
	//
	// 172.16.49.***
	PrivateIp *string `json:"PrivateIp,omitempty" xml:"PrivateIp,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 2BE6E619-A657-42E3-AD2D-18F8428A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// rg-acfmyiu4ekp****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// sg-bpcfmyiu4ekp****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// example:
	//
	// 60
	Storage *int64 `json:"Storage,omitempty" xml:"Storage,omitempty"`
	// example:
	//
	// essd_pl1
	StorageType *string                                                     `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	Tags        *DescribeTairKVCacheCustomInstanceAttributeResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	UseEni      *bool                                                       `json:"UseEni,omitempty" xml:"UseEni,omitempty"`
	// example:
	//
	// vsw-bp1e7clcw529l773d****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// example:
	//
	// vpc-bp1nme44gek34slfc****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// example:
	//
	// cn-hangzhou-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	// example:
	//
	// singlezone
	ZoneType *string `json:"ZoneType,omitempty" xml:"ZoneType,omitempty"`
}

func (s DescribeTairKVCacheCustomInstanceAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTairKVCacheCustomInstanceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponseBody) GetArchitectureType() *string {
	return s.ArchitectureType
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponseBody) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponseBody) GetCpu() *int64 {
	return s.Cpu
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponseBody) GetDisks() *DescribeTairKVCacheCustomInstanceAttributeResponseBodyDisks {
	return s.Disks
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponseBody) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponseBody) GetInstanceClass() *string {
	return s.InstanceClass
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponseBody) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponseBody) GetInstanceStatus() *string {
	return s.InstanceStatus
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponseBody) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponseBody) GetIsOrderCompleted() *bool {
	return s.IsOrderCompleted
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponseBody) GetMemory() *int64 {
	return s.Memory
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponseBody) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponseBody) GetPrivateIp() *string {
	return s.PrivateIp
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponseBody) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponseBody) GetStorage() *int64 {
	return s.Storage
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponseBody) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponseBody) GetTags() *DescribeTairKVCacheCustomInstanceAttributeResponseBodyTags {
	return s.Tags
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponseBody) GetUseEni() *bool {
	return s.UseEni
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponseBody) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponseBody) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponseBody) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponseBody) GetZoneType() *string {
	return s.ZoneType
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponseBody) SetArchitectureType(v string) *DescribeTairKVCacheCustomInstanceAttributeResponseBody {
	s.ArchitectureType = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponseBody) SetChargeType(v string) *DescribeTairKVCacheCustomInstanceAttributeResponseBody {
	s.ChargeType = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponseBody) SetCpu(v int64) *DescribeTairKVCacheCustomInstanceAttributeResponseBody {
	s.Cpu = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponseBody) SetCreateTime(v string) *DescribeTairKVCacheCustomInstanceAttributeResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponseBody) SetDisks(v *DescribeTairKVCacheCustomInstanceAttributeResponseBodyDisks) *DescribeTairKVCacheCustomInstanceAttributeResponseBody {
	s.Disks = v
	return s
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponseBody) SetEndTime(v string) *DescribeTairKVCacheCustomInstanceAttributeResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponseBody) SetImageId(v string) *DescribeTairKVCacheCustomInstanceAttributeResponseBody {
	s.ImageId = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponseBody) SetInstanceClass(v string) *DescribeTairKVCacheCustomInstanceAttributeResponseBody {
	s.InstanceClass = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponseBody) SetInstanceId(v string) *DescribeTairKVCacheCustomInstanceAttributeResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponseBody) SetInstanceName(v string) *DescribeTairKVCacheCustomInstanceAttributeResponseBody {
	s.InstanceName = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponseBody) SetInstanceStatus(v string) *DescribeTairKVCacheCustomInstanceAttributeResponseBody {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponseBody) SetInstanceType(v string) *DescribeTairKVCacheCustomInstanceAttributeResponseBody {
	s.InstanceType = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponseBody) SetIsOrderCompleted(v bool) *DescribeTairKVCacheCustomInstanceAttributeResponseBody {
	s.IsOrderCompleted = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponseBody) SetMemory(v int64) *DescribeTairKVCacheCustomInstanceAttributeResponseBody {
	s.Memory = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponseBody) SetNetworkType(v string) *DescribeTairKVCacheCustomInstanceAttributeResponseBody {
	s.NetworkType = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponseBody) SetPrivateIp(v string) *DescribeTairKVCacheCustomInstanceAttributeResponseBody {
	s.PrivateIp = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponseBody) SetRegionId(v string) *DescribeTairKVCacheCustomInstanceAttributeResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponseBody) SetRequestId(v string) *DescribeTairKVCacheCustomInstanceAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponseBody) SetResourceGroupId(v string) *DescribeTairKVCacheCustomInstanceAttributeResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponseBody) SetSecurityGroupId(v string) *DescribeTairKVCacheCustomInstanceAttributeResponseBody {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponseBody) SetStorage(v int64) *DescribeTairKVCacheCustomInstanceAttributeResponseBody {
	s.Storage = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponseBody) SetStorageType(v string) *DescribeTairKVCacheCustomInstanceAttributeResponseBody {
	s.StorageType = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponseBody) SetTags(v *DescribeTairKVCacheCustomInstanceAttributeResponseBodyTags) *DescribeTairKVCacheCustomInstanceAttributeResponseBody {
	s.Tags = v
	return s
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponseBody) SetUseEni(v bool) *DescribeTairKVCacheCustomInstanceAttributeResponseBody {
	s.UseEni = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponseBody) SetVSwitchId(v string) *DescribeTairKVCacheCustomInstanceAttributeResponseBody {
	s.VSwitchId = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponseBody) SetVpcId(v string) *DescribeTairKVCacheCustomInstanceAttributeResponseBody {
	s.VpcId = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponseBody) SetZoneId(v string) *DescribeTairKVCacheCustomInstanceAttributeResponseBody {
	s.ZoneId = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponseBody) SetZoneType(v string) *DescribeTairKVCacheCustomInstanceAttributeResponseBody {
	s.ZoneType = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeTairKVCacheCustomInstanceAttributeResponseBodyDisks struct {
	Disk []*DescribeTairKVCacheCustomInstanceAttributeResponseBodyDisksDisk `json:"Disk,omitempty" xml:"Disk,omitempty" type:"Repeated"`
}

func (s DescribeTairKVCacheCustomInstanceAttributeResponseBodyDisks) String() string {
	return dara.Prettify(s)
}

func (s DescribeTairKVCacheCustomInstanceAttributeResponseBodyDisks) GoString() string {
	return s.String()
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponseBodyDisks) GetDisk() []*DescribeTairKVCacheCustomInstanceAttributeResponseBodyDisksDisk {
	return s.Disk
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponseBodyDisks) SetDisk(v []*DescribeTairKVCacheCustomInstanceAttributeResponseBodyDisksDisk) *DescribeTairKVCacheCustomInstanceAttributeResponseBodyDisks {
	s.Disk = v
	return s
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponseBodyDisks) Validate() error {
	return dara.Validate(s)
}

type DescribeTairKVCacheCustomInstanceAttributeResponseBodyDisksDisk struct {
	// example:
	//
	// d-5v1aggi3ffoxufb57**
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// example:
	//
	// 100
	Size *string `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// data
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeTairKVCacheCustomInstanceAttributeResponseBodyDisksDisk) String() string {
	return dara.Prettify(s)
}

func (s DescribeTairKVCacheCustomInstanceAttributeResponseBodyDisksDisk) GoString() string {
	return s.String()
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponseBodyDisksDisk) GetDiskId() *string {
	return s.DiskId
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponseBodyDisksDisk) GetSize() *string {
	return s.Size
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponseBodyDisksDisk) GetType() *string {
	return s.Type
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponseBodyDisksDisk) SetDiskId(v string) *DescribeTairKVCacheCustomInstanceAttributeResponseBodyDisksDisk {
	s.DiskId = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponseBodyDisksDisk) SetSize(v string) *DescribeTairKVCacheCustomInstanceAttributeResponseBodyDisksDisk {
	s.Size = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponseBodyDisksDisk) SetType(v string) *DescribeTairKVCacheCustomInstanceAttributeResponseBodyDisksDisk {
	s.Type = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponseBodyDisksDisk) Validate() error {
	return dara.Validate(s)
}

type DescribeTairKVCacheCustomInstanceAttributeResponseBodyTags struct {
	Tag []*DescribeTairKVCacheCustomInstanceAttributeResponseBodyTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeTairKVCacheCustomInstanceAttributeResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeTairKVCacheCustomInstanceAttributeResponseBodyTags) GoString() string {
	return s.String()
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponseBodyTags) GetTag() []*DescribeTairKVCacheCustomInstanceAttributeResponseBodyTagsTag {
	return s.Tag
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponseBodyTags) SetTag(v []*DescribeTairKVCacheCustomInstanceAttributeResponseBodyTagsTag) *DescribeTairKVCacheCustomInstanceAttributeResponseBodyTags {
	s.Tag = v
	return s
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponseBodyTags) Validate() error {
	return dara.Validate(s)
}

type DescribeTairKVCacheCustomInstanceAttributeResponseBodyTagsTag struct {
	// example:
	//
	// tag1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeTairKVCacheCustomInstanceAttributeResponseBodyTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeTairKVCacheCustomInstanceAttributeResponseBodyTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponseBodyTagsTag) GetKey() *string {
	return s.Key
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponseBodyTagsTag) GetValue() *string {
	return s.Value
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponseBodyTagsTag) SetKey(v string) *DescribeTairKVCacheCustomInstanceAttributeResponseBodyTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponseBodyTagsTag) SetValue(v string) *DescribeTairKVCacheCustomInstanceAttributeResponseBodyTagsTag {
	s.Value = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstanceAttributeResponseBodyTagsTag) Validate() error {
	return dara.Validate(s)
}
