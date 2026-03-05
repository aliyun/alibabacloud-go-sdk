// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstances(v *DescribeInstancesResponseBodyInstances) *DescribeInstancesResponseBody
	GetInstances() *DescribeInstancesResponseBodyInstances
	SetPageNumber(v int32) *DescribeInstancesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeInstancesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeInstancesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeInstancesResponseBody
	GetTotalCount() *int32
}

type DescribeInstancesResponseBody struct {
	Instances *DescribeInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1E83311F-0EE4-4922-A3BF-730B312B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of instances.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBody) GetInstances() *DescribeInstancesResponseBodyInstances {
	return s.Instances
}

func (s *DescribeInstancesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeInstancesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstancesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeInstancesResponseBody) SetInstances(v *DescribeInstancesResponseBodyInstances) *DescribeInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *DescribeInstancesResponseBody) SetPageNumber(v int32) *DescribeInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetPageSize(v int32) *DescribeInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetRequestId(v string) *DescribeInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetTotalCount(v int32) *DescribeInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeInstancesResponseBody) Validate() error {
	if s.Instances != nil {
		if err := s.Instances.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInstancesResponseBodyInstances struct {
	KVStoreInstance []*DescribeInstancesResponseBodyInstancesKVStoreInstance `json:"KVStoreInstance,omitempty" xml:"KVStoreInstance,omitempty" type:"Repeated"`
}

func (s DescribeInstancesResponseBodyInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstances) GetKVStoreInstance() []*DescribeInstancesResponseBodyInstancesKVStoreInstance {
	return s.KVStoreInstance
}

func (s *DescribeInstancesResponseBodyInstances) SetKVStoreInstance(v []*DescribeInstancesResponseBodyInstancesKVStoreInstance) *DescribeInstancesResponseBodyInstances {
	s.KVStoreInstance = v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) Validate() error {
	if s.KVStoreInstance != nil {
		for _, item := range s.KVStoreInstance {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstancesResponseBodyInstancesKVStoreInstance struct {
	ArchitectureType    *string                                                    `json:"ArchitectureType,omitempty" xml:"ArchitectureType,omitempty"`
	Bandwidth           *int64                                                     `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	Capacity            *int64                                                     `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	ChargeType          *string                                                    `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	CloudType           *string                                                    `json:"CloudType,omitempty" xml:"CloudType,omitempty"`
	ComputingType       *string                                                    `json:"ComputingType,omitempty" xml:"ComputingType,omitempty"`
	Config              *string                                                    `json:"Config,omitempty" xml:"Config,omitempty"`
	ConnectionDomain    *string                                                    `json:"ConnectionDomain,omitempty" xml:"ConnectionDomain,omitempty"`
	ConnectionMode      *string                                                    `json:"ConnectionMode,omitempty" xml:"ConnectionMode,omitempty"`
	Connections         *int64                                                     `json:"Connections,omitempty" xml:"Connections,omitempty"`
	CreateTime          *string                                                    `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DestroyTime         *string                                                    `json:"DestroyTime,omitempty" xml:"DestroyTime,omitempty"`
	EditionType         *string                                                    `json:"EditionType,omitempty" xml:"EditionType,omitempty"`
	EndTime             *string                                                    `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	EngineVersion       *string                                                    `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	GlobalInstanceId    *string                                                    `json:"GlobalInstanceId,omitempty" xml:"GlobalInstanceId,omitempty"`
	HasRenewChangeOrder *bool                                                      `json:"HasRenewChangeOrder,omitempty" xml:"HasRenewChangeOrder,omitempty"`
	InstanceClass       *string                                                    `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	InstanceId          *string                                                    `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName        *string                                                    `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	InstanceStatus      *string                                                    `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	InstanceType        *string                                                    `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	IsRds               *bool                                                      `json:"IsRds,omitempty" xml:"IsRds,omitempty"`
	NetworkType         *string                                                    `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	NodeType            *string                                                    `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	PackageType         *string                                                    `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	Port                *int64                                                     `json:"Port,omitempty" xml:"Port,omitempty"`
	PrivateIp           *string                                                    `json:"PrivateIp,omitempty" xml:"PrivateIp,omitempty"`
	QPS                 *int64                                                     `json:"QPS,omitempty" xml:"QPS,omitempty"`
	ReadOnlyCount       *string                                                    `json:"ReadOnlyCount,omitempty" xml:"ReadOnlyCount,omitempty"`
	RegionId            *string                                                    `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ReplacateId         *string                                                    `json:"ReplacateId,omitempty" xml:"ReplacateId,omitempty"`
	ReplicaCount        *int32                                                     `json:"ReplicaCount,omitempty" xml:"ReplicaCount,omitempty"`
	ResourceGroupId     *string                                                    `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SecondaryZoneId     *string                                                    `json:"SecondaryZoneId,omitempty" xml:"SecondaryZoneId,omitempty"`
	ShardClass          *string                                                    `json:"ShardClass,omitempty" xml:"ShardClass,omitempty"`
	ShardCount          *int32                                                     `json:"ShardCount,omitempty" xml:"ShardCount,omitempty"`
	SlaveReadOnlyCount  *int32                                                     `json:"SlaveReadOnlyCount,omitempty" xml:"SlaveReadOnlyCount,omitempty"`
	SlaveReplicaCount   *int32                                                     `json:"SlaveReplicaCount,omitempty" xml:"SlaveReplicaCount,omitempty"`
	Tags                *DescribeInstancesResponseBodyInstancesKVStoreInstanceTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	UserName            *string                                                    `json:"UserName,omitempty" xml:"UserName,omitempty"`
	VSwitchId           *string                                                    `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId               *string                                                    `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneId              *string                                                    `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesKVStoreInstance) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesKVStoreInstance) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) GetArchitectureType() *string {
	return s.ArchitectureType
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) GetBandwidth() *int64 {
	return s.Bandwidth
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) GetCapacity() *int64 {
	return s.Capacity
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) GetCloudType() *string {
	return s.CloudType
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) GetComputingType() *string {
	return s.ComputingType
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) GetConfig() *string {
	return s.Config
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) GetConnectionDomain() *string {
	return s.ConnectionDomain
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) GetConnectionMode() *string {
	return s.ConnectionMode
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) GetConnections() *int64 {
	return s.Connections
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) GetDestroyTime() *string {
	return s.DestroyTime
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) GetEditionType() *string {
	return s.EditionType
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) GetGlobalInstanceId() *string {
	return s.GlobalInstanceId
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) GetHasRenewChangeOrder() *bool {
	return s.HasRenewChangeOrder
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) GetInstanceClass() *string {
	return s.InstanceClass
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) GetInstanceStatus() *string {
	return s.InstanceStatus
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) GetIsRds() *bool {
	return s.IsRds
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) GetNodeType() *string {
	return s.NodeType
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) GetPackageType() *string {
	return s.PackageType
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) GetPort() *int64 {
	return s.Port
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) GetPrivateIp() *string {
	return s.PrivateIp
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) GetQPS() *int64 {
	return s.QPS
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) GetReadOnlyCount() *string {
	return s.ReadOnlyCount
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) GetReplacateId() *string {
	return s.ReplacateId
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) GetReplicaCount() *int32 {
	return s.ReplicaCount
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) GetSecondaryZoneId() *string {
	return s.SecondaryZoneId
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) GetShardClass() *string {
	return s.ShardClass
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) GetShardCount() *int32 {
	return s.ShardCount
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) GetSlaveReadOnlyCount() *int32 {
	return s.SlaveReadOnlyCount
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) GetSlaveReplicaCount() *int32 {
	return s.SlaveReplicaCount
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) GetTags() *DescribeInstancesResponseBodyInstancesKVStoreInstanceTags {
	return s.Tags
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) GetUserName() *string {
	return s.UserName
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetArchitectureType(v string) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.ArchitectureType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetBandwidth(v int64) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.Bandwidth = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetCapacity(v int64) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.Capacity = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetChargeType(v string) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.ChargeType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetCloudType(v string) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.CloudType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetComputingType(v string) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.ComputingType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetConfig(v string) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.Config = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetConnectionDomain(v string) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.ConnectionDomain = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetConnectionMode(v string) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.ConnectionMode = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetConnections(v int64) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.Connections = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetCreateTime(v string) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.CreateTime = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetDestroyTime(v string) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.DestroyTime = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetEditionType(v string) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.EditionType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetEndTime(v string) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.EndTime = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetEngineVersion(v string) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.EngineVersion = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetGlobalInstanceId(v string) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.GlobalInstanceId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetHasRenewChangeOrder(v bool) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.HasRenewChangeOrder = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetInstanceClass(v string) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.InstanceClass = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetInstanceId(v string) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetInstanceName(v string) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.InstanceName = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetInstanceStatus(v string) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetInstanceType(v string) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.InstanceType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetIsRds(v bool) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.IsRds = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetNetworkType(v string) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.NetworkType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetNodeType(v string) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.NodeType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetPackageType(v string) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.PackageType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetPort(v int64) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.Port = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetPrivateIp(v string) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.PrivateIp = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetQPS(v int64) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.QPS = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetReadOnlyCount(v string) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.ReadOnlyCount = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetRegionId(v string) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.RegionId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetReplacateId(v string) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.ReplacateId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetReplicaCount(v int32) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.ReplicaCount = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetResourceGroupId(v string) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetSecondaryZoneId(v string) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.SecondaryZoneId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetShardClass(v string) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.ShardClass = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetShardCount(v int32) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.ShardCount = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetSlaveReadOnlyCount(v int32) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.SlaveReadOnlyCount = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetSlaveReplicaCount(v int32) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.SlaveReplicaCount = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetTags(v *DescribeInstancesResponseBodyInstancesKVStoreInstanceTags) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.Tags = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetUserName(v string) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.UserName = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetVSwitchId(v string) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.VSwitchId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetVpcId(v string) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.VpcId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) SetZoneId(v string) *DescribeInstancesResponseBodyInstancesKVStoreInstance {
	s.ZoneId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstance) Validate() error {
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInstancesResponseBodyInstancesKVStoreInstanceTags struct {
	Tag []*DescribeInstancesResponseBodyInstancesKVStoreInstanceTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeInstancesResponseBodyInstancesKVStoreInstanceTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesKVStoreInstanceTags) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstanceTags) GetTag() []*DescribeInstancesResponseBodyInstancesKVStoreInstanceTagsTag {
	return s.Tag
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstanceTags) SetTag(v []*DescribeInstancesResponseBodyInstancesKVStoreInstanceTagsTag) *DescribeInstancesResponseBodyInstancesKVStoreInstanceTags {
	s.Tag = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstanceTags) Validate() error {
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

type DescribeInstancesResponseBodyInstancesKVStoreInstanceTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesKVStoreInstanceTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesKVStoreInstanceTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstanceTagsTag) GetKey() *string {
	return s.Key
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstanceTagsTag) GetValue() *string {
	return s.Value
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstanceTagsTag) SetKey(v string) *DescribeInstancesResponseBodyInstancesKVStoreInstanceTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstanceTagsTag) SetValue(v string) *DescribeInstancesResponseBodyInstancesKVStoreInstanceTagsTag {
	s.Value = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesKVStoreInstanceTagsTag) Validate() error {
	return dara.Validate(s)
}
