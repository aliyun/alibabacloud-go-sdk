// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTairKVCacheCustomInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstances(v *DescribeTairKVCacheCustomInstancesResponseBodyInstances) *DescribeTairKVCacheCustomInstancesResponseBody
	GetInstances() *DescribeTairKVCacheCustomInstancesResponseBodyInstances
	SetPageNumber(v int32) *DescribeTairKVCacheCustomInstancesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeTairKVCacheCustomInstancesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeTairKVCacheCustomInstancesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeTairKVCacheCustomInstancesResponseBody
	GetTotalCount() *int32
}

type DescribeTairKVCacheCustomInstancesResponseBody struct {
	Instances *DescribeTairKVCacheCustomInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Struct"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// B79C1A90-495B-4E99-A2AA-A4DB13B8****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 40
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeTairKVCacheCustomInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTairKVCacheCustomInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTairKVCacheCustomInstancesResponseBody) GetInstances() *DescribeTairKVCacheCustomInstancesResponseBodyInstances {
	return s.Instances
}

func (s *DescribeTairKVCacheCustomInstancesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeTairKVCacheCustomInstancesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeTairKVCacheCustomInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTairKVCacheCustomInstancesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeTairKVCacheCustomInstancesResponseBody) SetInstances(v *DescribeTairKVCacheCustomInstancesResponseBodyInstances) *DescribeTairKVCacheCustomInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *DescribeTairKVCacheCustomInstancesResponseBody) SetPageNumber(v int32) *DescribeTairKVCacheCustomInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstancesResponseBody) SetPageSize(v int32) *DescribeTairKVCacheCustomInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstancesResponseBody) SetRequestId(v string) *DescribeTairKVCacheCustomInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstancesResponseBody) SetTotalCount(v int32) *DescribeTairKVCacheCustomInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeTairKVCacheCustomInstancesResponseBodyInstances struct {
	KVStoreInstance []*DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstance `json:"KVStoreInstance,omitempty" xml:"KVStoreInstance,omitempty" type:"Repeated"`
}

func (s DescribeTairKVCacheCustomInstancesResponseBodyInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeTairKVCacheCustomInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeTairKVCacheCustomInstancesResponseBodyInstances) GetKVStoreInstance() []*DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstance {
	return s.KVStoreInstance
}

func (s *DescribeTairKVCacheCustomInstancesResponseBodyInstances) SetKVStoreInstance(v []*DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstance) *DescribeTairKVCacheCustomInstancesResponseBodyInstances {
	s.KVStoreInstance = v
	return s
}

func (s *DescribeTairKVCacheCustomInstancesResponseBodyInstances) Validate() error {
	return dara.Validate(s)
}

type DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstance struct {
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// example:
	//
	// 2018-11-07T08:49:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 2019-04-28T10:03:01Z
	DestroyTime *string `json:"DestroyTime,omitempty" xml:"DestroyTime,omitempty"`
	// example:
	//
	// 2024-05-21T00:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
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
	// testdb
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
	// VPC
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	PrivateIp   *string `json:"PrivateIp,omitempty" xml:"PrivateIp,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-acfmyiu4ekp****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 50
	Storage *int64 `json:"Storage,omitempty" xml:"Storage,omitempty"`
	// example:
	//
	// essd_pl1
	StorageType *string                                                                     `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	Tags        *DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstanceTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	UseEni      *bool                                                                       `json:"UseEni,omitempty" xml:"UseEni,omitempty"`
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
	// cn-hangzhou-e
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstance) String() string {
	return dara.Prettify(s)
}

func (s DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstance) GoString() string {
	return s.String()
}

func (s *DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstance) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstance) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstance) GetDestroyTime() *string {
	return s.DestroyTime
}

func (s *DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstance) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstance) GetInstanceClass() *string {
	return s.InstanceClass
}

func (s *DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstance) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstance) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstance) GetInstanceStatus() *string {
	return s.InstanceStatus
}

func (s *DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstance) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstance) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstance) GetPrivateIp() *string {
	return s.PrivateIp
}

func (s *DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstance) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstance) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstance) GetStorage() *int64 {
	return s.Storage
}

func (s *DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstance) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstance) GetTags() *DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstanceTags {
	return s.Tags
}

func (s *DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstance) GetUseEni() *bool {
	return s.UseEni
}

func (s *DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstance) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstance) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstance) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstance) SetChargeType(v string) *DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstance {
	s.ChargeType = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstance) SetCreateTime(v string) *DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstance {
	s.CreateTime = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstance) SetDestroyTime(v string) *DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstance {
	s.DestroyTime = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstance) SetEndTime(v string) *DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstance {
	s.EndTime = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstance) SetInstanceClass(v string) *DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstance {
	s.InstanceClass = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstance) SetInstanceId(v string) *DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstance {
	s.InstanceId = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstance) SetInstanceName(v string) *DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstance {
	s.InstanceName = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstance) SetInstanceStatus(v string) *DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstance {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstance) SetInstanceType(v string) *DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstance {
	s.InstanceType = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstance) SetNetworkType(v string) *DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstance {
	s.NetworkType = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstance) SetPrivateIp(v string) *DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstance {
	s.PrivateIp = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstance) SetRegionId(v string) *DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstance {
	s.RegionId = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstance) SetResourceGroupId(v string) *DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstance {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstance) SetStorage(v int64) *DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstance {
	s.Storage = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstance) SetStorageType(v string) *DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstance {
	s.StorageType = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstance) SetTags(v *DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstanceTags) *DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstance {
	s.Tags = v
	return s
}

func (s *DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstance) SetUseEni(v bool) *DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstance {
	s.UseEni = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstance) SetVSwitchId(v string) *DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstance {
	s.VSwitchId = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstance) SetVpcId(v string) *DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstance {
	s.VpcId = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstance) SetZoneId(v string) *DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstance {
	s.ZoneId = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstance) Validate() error {
	return dara.Validate(s)
}

type DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstanceTags struct {
	Tag []*DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstanceTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstanceTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstanceTags) GoString() string {
	return s.String()
}

func (s *DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstanceTags) GetTag() []*DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstanceTagsTag {
	return s.Tag
}

func (s *DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstanceTags) SetTag(v []*DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstanceTagsTag) *DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstanceTags {
	s.Tag = v
	return s
}

func (s *DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstanceTags) Validate() error {
	return dara.Validate(s)
}

type DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstanceTagsTag struct {
	// example:
	//
	// tagkey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// tagvalue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstanceTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstanceTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstanceTagsTag) GetKey() *string {
	return s.Key
}

func (s *DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstanceTagsTag) GetValue() *string {
	return s.Value
}

func (s *DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstanceTagsTag) SetKey(v string) *DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstanceTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstanceTagsTag) SetValue(v string) *DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstanceTagsTag {
	s.Value = &v
	return s
}

func (s *DescribeTairKVCacheCustomInstancesResponseBodyInstancesKVStoreInstanceTagsTag) Validate() error {
	return dara.Validate(s)
}
