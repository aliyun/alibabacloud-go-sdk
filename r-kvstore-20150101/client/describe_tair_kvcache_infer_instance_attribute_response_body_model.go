// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTairKVCacheInferInstanceAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstances(v *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstances) *DescribeTairKVCacheInferInstanceAttributeResponseBody
	GetInstances() *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstances
	SetRequestId(v string) *DescribeTairKVCacheInferInstanceAttributeResponseBody
	GetRequestId() *string
}

type DescribeTairKVCacheInferInstanceAttributeResponseBody struct {
	Instances *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Struct"`
	RequestId *string                                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeTairKVCacheInferInstanceAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTairKVCacheInferInstanceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponseBody) GetInstances() *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstances {
	return s.Instances
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponseBody) SetInstances(v *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstances) *DescribeTairKVCacheInferInstanceAttributeResponseBody {
	s.Instances = v
	return s
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponseBody) SetRequestId(v string) *DescribeTairKVCacheInferInstanceAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponseBody) Validate() error {
	if s.Instances != nil {
		if err := s.Instances.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeTairKVCacheInferInstanceAttributeResponseBodyInstances struct {
	DBInstanceAttribute []*DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute `json:"DBInstanceAttribute,omitempty" xml:"DBInstanceAttribute,omitempty" type:"Repeated"`
}

func (s DescribeTairKVCacheInferInstanceAttributeResponseBodyInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeTairKVCacheInferInstanceAttributeResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstances) GetDBInstanceAttribute() []*DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	return s.DBInstanceAttribute
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstances) SetDBInstanceAttribute(v []*DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute) *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstances {
	s.DBInstanceAttribute = v
	return s
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstances) Validate() error {
	if s.DBInstanceAttribute != nil {
		for _, item := range s.DBInstanceAttribute {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute struct {
	ArchitectureType *string                                                                                `json:"ArchitectureType,omitempty" xml:"ArchitectureType,omitempty"`
	ChargeType       *string                                                                                `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	ComputeUnitNum   *int32                                                                                 `json:"ComputeUnitNum,omitempty" xml:"ComputeUnitNum,omitempty"`
	ConnectionString *string                                                                                `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	CreateTime       *string                                                                                `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	EndTime          *string                                                                                `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceClass    *string                                                                                `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	InstanceId       *string                                                                                `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName     *string                                                                                `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	InstanceStatus   *string                                                                                `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	InstanceType     *string                                                                                `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	IsDelete         *int32                                                                                 `json:"IsDelete,omitempty" xml:"IsDelete,omitempty"`
	IsOrderCompleted *string                                                                                `json:"IsOrderCompleted,omitempty" xml:"IsOrderCompleted,omitempty"`
	Model            *string                                                                                `json:"Model,omitempty" xml:"Model,omitempty"`
	ModelServiceNum  *int32                                                                                 `json:"ModelServiceNum,omitempty" xml:"ModelServiceNum,omitempty"`
	NetworkType      *string                                                                                `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	PrivateIp        *string                                                                                `json:"PrivateIp,omitempty" xml:"PrivateIp,omitempty"`
	RegionId         *string                                                                                `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ReplicaNum       *string                                                                                `json:"ReplicaNum,omitempty" xml:"ReplicaNum,omitempty"`
	ResourceGroupId  *string                                                                                `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Storage          *int64                                                                                 `json:"Storage,omitempty" xml:"Storage,omitempty"`
	Tags             *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttributeTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	VSwitchId        *string                                                                                `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId            *string                                                                                `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneId           *string                                                                                `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	ZoneType         *string                                                                                `json:"ZoneType,omitempty" xml:"ZoneType,omitempty"`
}

func (s DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute) String() string {
	return dara.Prettify(s)
}

func (s DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GoString() string {
	return s.String()
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetArchitectureType() *string {
	return s.ArchitectureType
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetComputeUnitNum() *int32 {
	return s.ComputeUnitNum
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetInstanceClass() *string {
	return s.InstanceClass
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetInstanceStatus() *string {
	return s.InstanceStatus
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetIsDelete() *int32 {
	return s.IsDelete
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetIsOrderCompleted() *string {
	return s.IsOrderCompleted
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetModel() *string {
	return s.Model
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetModelServiceNum() *int32 {
	return s.ModelServiceNum
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetPrivateIp() *string {
	return s.PrivateIp
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetReplicaNum() *string {
	return s.ReplicaNum
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetStorage() *int64 {
	return s.Storage
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetTags() *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttributeTags {
	return s.Tags
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute) GetZoneType() *string {
	return s.ZoneType
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetArchitectureType(v string) *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.ArchitectureType = &v
	return s
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetChargeType(v string) *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.ChargeType = &v
	return s
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetComputeUnitNum(v int32) *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.ComputeUnitNum = &v
	return s
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetConnectionString(v string) *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.ConnectionString = &v
	return s
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetCreateTime(v string) *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.CreateTime = &v
	return s
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetEndTime(v string) *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.EndTime = &v
	return s
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetInstanceClass(v string) *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.InstanceClass = &v
	return s
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetInstanceId(v string) *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.InstanceId = &v
	return s
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetInstanceName(v string) *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.InstanceName = &v
	return s
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetInstanceStatus(v string) *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetInstanceType(v string) *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.InstanceType = &v
	return s
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetIsDelete(v int32) *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.IsDelete = &v
	return s
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetIsOrderCompleted(v string) *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.IsOrderCompleted = &v
	return s
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetModel(v string) *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.Model = &v
	return s
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetModelServiceNum(v int32) *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.ModelServiceNum = &v
	return s
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetNetworkType(v string) *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.NetworkType = &v
	return s
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetPrivateIp(v string) *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.PrivateIp = &v
	return s
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetRegionId(v string) *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.RegionId = &v
	return s
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetReplicaNum(v string) *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.ReplicaNum = &v
	return s
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetResourceGroupId(v string) *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetStorage(v int64) *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.Storage = &v
	return s
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetTags(v *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttributeTags) *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.Tags = v
	return s
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetVSwitchId(v string) *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.VSwitchId = &v
	return s
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetVpcId(v string) *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.VpcId = &v
	return s
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetZoneId(v string) *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.ZoneId = &v
	return s
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute) SetZoneType(v string) *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute {
	s.ZoneType = &v
	return s
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttribute) Validate() error {
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttributeTags struct {
	Tag []*DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttributeTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttributeTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttributeTags) GoString() string {
	return s.String()
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttributeTags) GetTag() []*DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttributeTagsTag {
	return s.Tag
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttributeTags) SetTag(v []*DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttributeTagsTag) *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttributeTags {
	s.Tag = v
	return s
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttributeTags) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttributeTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttributeTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttributeTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttributeTagsTag) GetKey() *string {
	return s.Key
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttributeTagsTag) GetValue() *string {
	return s.Value
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttributeTagsTag) SetKey(v string) *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttributeTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttributeTagsTag) SetValue(v string) *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttributeTagsTag {
	s.Value = &v
	return s
}

func (s *DescribeTairKVCacheInferInstanceAttributeResponseBodyInstancesDBInstanceAttributeTagsTag) Validate() error {
	return dara.Validate(s)
}
