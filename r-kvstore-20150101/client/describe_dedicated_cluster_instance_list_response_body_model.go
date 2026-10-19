// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDedicatedClusterInstanceListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstances(v []*DescribeDedicatedClusterInstanceListResponseBodyInstances) *DescribeDedicatedClusterInstanceListResponseBody
	GetInstances() []*DescribeDedicatedClusterInstanceListResponseBodyInstances
	SetPageNumber(v int32) *DescribeDedicatedClusterInstanceListResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDedicatedClusterInstanceListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeDedicatedClusterInstanceListResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeDedicatedClusterInstanceListResponseBody
	GetTotalCount() *int32
}

type DescribeDedicatedClusterInstanceListResponseBody struct {
	// The list of instance information.
	Instances []*DescribeDedicatedClusterInstanceListResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 25
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 712CCF2A-16BD-411B-93F7-E978BEF2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 0
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDedicatedClusterInstanceListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedClusterInstanceListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedClusterInstanceListResponseBody) GetInstances() []*DescribeDedicatedClusterInstanceListResponseBodyInstances {
	return s.Instances
}

func (s *DescribeDedicatedClusterInstanceListResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDedicatedClusterInstanceListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDedicatedClusterInstanceListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDedicatedClusterInstanceListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeDedicatedClusterInstanceListResponseBody) SetInstances(v []*DescribeDedicatedClusterInstanceListResponseBodyInstances) *DescribeDedicatedClusterInstanceListResponseBody {
	s.Instances = v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBody) SetPageNumber(v int32) *DescribeDedicatedClusterInstanceListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBody) SetPageSize(v int32) *DescribeDedicatedClusterInstanceListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBody) SetRequestId(v string) *DescribeDedicatedClusterInstanceListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBody) SetTotalCount(v int32) *DescribeDedicatedClusterInstanceListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBody) Validate() error {
	if s.Instances != nil {
		for _, item := range s.Instances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDedicatedClusterInstanceListResponseBodyInstances struct {
	// The default bandwidth of the instance. Unit: MB/s.
	//
	// example:
	//
	// 24
	BandWidth *int64 `json:"BandWidth,omitempty" xml:"BandWidth,omitempty"`
	// The architecture type. Valid values:
	//
	// 	- **logic**: Cluster architecture.
	//
	// 	- **normal**: Standard architecture.
	//
	// example:
	//
	// logic
	CharacterType *string `json:"CharacterType,omitempty" xml:"CharacterType,omitempty"`
	// The dedicated cluster ID.
	//
	// example:
	//
	// dhg-rx71fc5ndh9o****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The dedicated cluster name.
	//
	// example:
	//
	// testname
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The direct connection endpoint.
	//
	// example:
	//
	// r-t4ncdi1dgi0ja8****.redis.hangzhou.rds.aliyuncs.com
	ConnectionDomain *string `json:"ConnectionDomain,omitempty" xml:"ConnectionDomain,omitempty"`
	// The time when the instance was created. The time is in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-05-06T07:09:40Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The current bandwidth of the instance. Unit: MB/s. The value is the sum of the default bandwidth and the increased bandwidth.
	//
	// example:
	//
	// 50
	CurrentBandWidth *int64 `json:"CurrentBandWidth,omitempty" xml:"CurrentBandWidth,omitempty"`
	// An internal parameter used for instance maintenance.
	//
	// example:
	//
	// 4652****
	CustomId *string `json:"CustomId,omitempty" xml:"CustomId,omitempty"`
	// The database engine. The value is fixed as **redis**.
	//
	// example:
	//
	// redis
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The database engine version. The value is fixed as **5.0**.
	//
	// example:
	//
	// 5.0
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The instance type.
	//
	// example:
	//
	// redis.cluster.sharding.common.ce
	InstanceClass *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// r-bp1e7vl6ygf1yq****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance name.
	//
	// example:
	//
	// testdb
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The list of node information.
	InstanceNodeList []*DescribeDedicatedClusterInstanceListResponseBodyInstancesInstanceNodeList `json:"InstanceNodeList,omitempty" xml:"InstanceNodeList,omitempty" type:"Repeated"`
	// The instance status. Valid values:
	//
	// 	- **Normal**: Normal.
	//
	// 	- **Creating**: Creating.
	//
	// 	- **Changing**: Changing.
	//
	// 	- **Inactive**: Disabled.
	//
	// 	- **Flushing**: Purge in progress.
	//
	// 	- **Released**: Released.
	//
	// 	- **Transforming**: Transforming.
	//
	// 	- **Unavailable**: Unavailable.
	//
	// 	- **Error**: Creation failed.
	//
	// 	- **Migrating**: Migrating.
	//
	// 	- **BackupRecovering**: Restoring from backup.
	//
	// 	- **MinorVersionUpgrading**: Minor version upgrade in progress.
	//
	// 	- **NetworkModifying**: Changing the network.
	//
	// 	- **SSLModifying**: Changing SSL configurations.
	//
	// 	- **MajorVersionUpgrading**: Major engine version upgrade in progress. The instance remains accessible.
	//
	// example:
	//
	// Normal
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// The end time of the maintenance window. The time is in the <i>HH:mm</i>Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 17:00Z
	MaintainEndTime *string `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	// The start time of the maintenance window. The time is in the <i>HH:mm</i>Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 16:00Z
	MaintainStartTime *string `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	// The number of proxy nodes.
	//
	// > 	- A value of **0*	- indicates that the proxy mode is disabled for the instance. A value greater than **0*	- (such as **1**) indicates that the proxy mode is enabled for the instance.
	//
	// > 	- This parameter is returned only when the instance uses the [cluster architecture](https://help.aliyun.com/document_detail/52228.html).
	//
	// example:
	//
	// 1
	ProxyCount *int32 `json:"ProxyCount,omitempty" xml:"ProxyCount,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of data shards.
	//
	// > This parameter is returned only when the instance uses the [cluster architecture](https://help.aliyun.com/document_detail/52228.html).
	//
	// example:
	//
	// 3
	ShardCount *int32 `json:"ShardCount,omitempty" xml:"ShardCount,omitempty"`
	// The storage type. The value is fixed as LOCAL_SSD ([ESSD cloud disk (standard SSD)](https://help.aliyun.com/document_detail/122389.html)).
	//
	// example:
	//
	// LOCAL_SSD
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-t4n2clc70t3hqwsrr****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-t4nvrca24dczppq44****
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-e
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDedicatedClusterInstanceListResponseBodyInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedClusterInstanceListResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) GetBandWidth() *int64 {
	return s.BandWidth
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) GetCharacterType() *string {
	return s.CharacterType
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) GetClusterName() *string {
	return s.ClusterName
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) GetConnectionDomain() *string {
	return s.ConnectionDomain
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) GetCurrentBandWidth() *int64 {
	return s.CurrentBandWidth
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) GetCustomId() *string {
	return s.CustomId
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) GetEngine() *string {
	return s.Engine
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) GetInstanceClass() *string {
	return s.InstanceClass
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) GetInstanceNodeList() []*DescribeDedicatedClusterInstanceListResponseBodyInstancesInstanceNodeList {
	return s.InstanceNodeList
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) GetInstanceStatus() *string {
	return s.InstanceStatus
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) GetMaintainEndTime() *string {
	return s.MaintainEndTime
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) GetMaintainStartTime() *string {
	return s.MaintainStartTime
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) GetProxyCount() *int32 {
	return s.ProxyCount
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) GetShardCount() *int32 {
	return s.ShardCount
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) GetVswitchId() *string {
	return s.VswitchId
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) SetBandWidth(v int64) *DescribeDedicatedClusterInstanceListResponseBodyInstances {
	s.BandWidth = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) SetCharacterType(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstances {
	s.CharacterType = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) SetClusterId(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstances {
	s.ClusterId = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) SetClusterName(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstances {
	s.ClusterName = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) SetConnectionDomain(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstances {
	s.ConnectionDomain = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) SetCreateTime(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstances {
	s.CreateTime = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) SetCurrentBandWidth(v int64) *DescribeDedicatedClusterInstanceListResponseBodyInstances {
	s.CurrentBandWidth = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) SetCustomId(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstances {
	s.CustomId = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) SetEngine(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstances {
	s.Engine = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) SetEngineVersion(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstances {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) SetInstanceClass(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstances {
	s.InstanceClass = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) SetInstanceId(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstances {
	s.InstanceId = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) SetInstanceName(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstances {
	s.InstanceName = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) SetInstanceNodeList(v []*DescribeDedicatedClusterInstanceListResponseBodyInstancesInstanceNodeList) *DescribeDedicatedClusterInstanceListResponseBodyInstances {
	s.InstanceNodeList = v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) SetInstanceStatus(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstances {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) SetMaintainEndTime(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstances {
	s.MaintainEndTime = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) SetMaintainStartTime(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstances {
	s.MaintainStartTime = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) SetProxyCount(v int32) *DescribeDedicatedClusterInstanceListResponseBodyInstances {
	s.ProxyCount = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) SetRegionId(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstances {
	s.RegionId = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) SetShardCount(v int32) *DescribeDedicatedClusterInstanceListResponseBodyInstances {
	s.ShardCount = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) SetStorageType(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstances {
	s.StorageType = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) SetVpcId(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstances {
	s.VpcId = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) SetVswitchId(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstances {
	s.VswitchId = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) SetZoneId(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstances {
	s.ZoneId = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstances) Validate() error {
	if s.InstanceNodeList != nil {
		for _, item := range s.InstanceNodeList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDedicatedClusterInstanceListResponseBodyInstancesInstanceNodeList struct {
	// The host ID in the dedicated cluster.
	//
	// example:
	//
	// ch-bp13vf0y9gx3c****
	DedicatedHostName *string `json:"DedicatedHostName,omitempty" xml:"DedicatedHostName,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// r-bp1e7vl6ygf1yq****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The node ID.
	//
	// example:
	//
	// 17921111
	NodeId *int32 `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The IP address of the node.
	//
	// example:
	//
	// 10.0.33.***
	NodeIp *string `json:"NodeIp,omitempty" xml:"NodeIp,omitempty"`
	// The node type. Valid values:
	//
	// 	- **db**: Data node.
	//
	// 	- **proxy**: Proxy node.
	//
	// 	- **normal**: Regular node. This value is returned for instances that use the standard architecture.
	//
	// example:
	//
	// normal
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// The port of the node.
	//
	// example:
	//
	// 3001
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The node role. Valid values:
	//
	// 	- **master**: Primary node.
	//
	// 	- **slave**: Secondary node.
	//
	// example:
	//
	// master
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The zone to which the node belongs.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDedicatedClusterInstanceListResponseBodyInstancesInstanceNodeList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedClusterInstanceListResponseBodyInstancesInstanceNodeList) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstancesInstanceNodeList) GetDedicatedHostName() *string {
	return s.DedicatedHostName
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstancesInstanceNodeList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstancesInstanceNodeList) GetNodeId() *int32 {
	return s.NodeId
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstancesInstanceNodeList) GetNodeIp() *string {
	return s.NodeIp
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstancesInstanceNodeList) GetNodeType() *string {
	return s.NodeType
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstancesInstanceNodeList) GetPort() *int32 {
	return s.Port
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstancesInstanceNodeList) GetRole() *string {
	return s.Role
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstancesInstanceNodeList) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstancesInstanceNodeList) SetDedicatedHostName(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstancesInstanceNodeList {
	s.DedicatedHostName = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstancesInstanceNodeList) SetInstanceId(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstancesInstanceNodeList {
	s.InstanceId = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstancesInstanceNodeList) SetNodeId(v int32) *DescribeDedicatedClusterInstanceListResponseBodyInstancesInstanceNodeList {
	s.NodeId = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstancesInstanceNodeList) SetNodeIp(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstancesInstanceNodeList {
	s.NodeIp = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstancesInstanceNodeList) SetNodeType(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstancesInstanceNodeList {
	s.NodeType = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstancesInstanceNodeList) SetPort(v int32) *DescribeDedicatedClusterInstanceListResponseBodyInstancesInstanceNodeList {
	s.Port = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstancesInstanceNodeList) SetRole(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstancesInstanceNodeList {
	s.Role = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstancesInstanceNodeList) SetZoneId(v string) *DescribeDedicatedClusterInstanceListResponseBodyInstancesInstanceNodeList {
	s.ZoneId = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponseBodyInstancesInstanceNodeList) Validate() error {
	return dara.Validate(s)
}
