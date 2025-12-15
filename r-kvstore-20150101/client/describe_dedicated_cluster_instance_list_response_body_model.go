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
	// Details about the instances.
	Instances []*DescribeDedicatedClusterInstanceListResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 712CCF2A-16BD-411B-93F7-E978BEF2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
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
	// The default bandwidth of the instance. Unit: Mbit/s.
	//
	// example:
	//
	// 24
	BandWidth *int64 `json:"BandWidth,omitempty" xml:"BandWidth,omitempty"`
	// The architecture of the instance. Valid values:
	//
	// 	- **logic**: cluster
	//
	// 	- **normal**: standard
	//
	// example:
	//
	// logic
	CharacterType *string `json:"CharacterType,omitempty" xml:"CharacterType,omitempty"`
	// The ID of the dedicated cluster to which the instance belongs.
	//
	// example:
	//
	// dhg-rx71fc5ndh9o****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The name of the dedicated cluster to which the instance belongs.
	//
	// example:
	//
	// testname
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The private endpoint of the instance.
	//
	// example:
	//
	// r-t4ncdi1dgi0ja8****.redis.hangzhou.rds.aliyuncs.com
	ConnectionDomain *string `json:"ConnectionDomain,omitempty" xml:"ConnectionDomain,omitempty"`
	// The time when the instance was created. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-05-06T07:09:40Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The current bandwidth of the instance, which is the sum of the default bandwidth and any extra bandwidth that is purchased. Unit: Mbit/s.
	//
	// example:
	//
	// 50
	CurrentBandWidth *int64 `json:"CurrentBandWidth,omitempty" xml:"CurrentBandWidth,omitempty"`
	// An internal parameter used for the maintenance and management of instances.
	//
	// example:
	//
	// 4652****
	CustomId *string `json:"CustomId,omitempty" xml:"CustomId,omitempty"`
	// The database engine. The return value is **redis**.
	//
	// example:
	//
	// redis
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The database engine version of the instance. The return value is **5.0**.
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
	// The ID of the instance.
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
	// Details about the nodes.
	InstanceNodeList []*DescribeDedicatedClusterInstanceListResponseBodyInstancesInstanceNodeList `json:"InstanceNodeList,omitempty" xml:"InstanceNodeList,omitempty" type:"Repeated"`
	// The state of the instance. Valid values:
	//
	// 	- **Normal**: The instance is normal.
	//
	// 	- **Creating**: The instance is being created.
	//
	// 	- **Changing**: The configurations of the instance are being changed.
	//
	// 	- **Inactive**: The instance is disabled.
	//
	// 	- **Flushing**: The instance is being released.
	//
	// 	- **Released**: The instance is released.
	//
	// 	- **Transforming**: The billing method of the instance is being changed.
	//
	// 	- **Unavailable**: The instance is unavailable.
	//
	// 	- **Error**: The instance failed to be created.
	//
	// 	- **Migrating**: The instance is being migrated.
	//
	// 	- **BackupRecovering**: The instance is being restored from a backup.
	//
	// 	- **MinorVersionUpgrading**: The minor version of the instance is being updated.
	//
	// 	- **NetworkModifying**: The network type of the instance is being changed.
	//
	// 	- **SSLModifying**: The SSL configurations of the instance are being changed.
	//
	// 	- **MajorVersionUpgrading**: The major version of the instance is being upgraded. The instance remains accessible during the upgrade.
	//
	// example:
	//
	// Normal
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// The end time of the maintenance window. The time is in the *HH:mm*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 17:00Z
	MaintainEndTime *string `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	// The start time of the maintenance window. The time is in the *HH:mm*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 16:00Z
	MaintainStartTime *string `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	// The number of proxy nodes.
	//
	// >
	//
	// 	- If the return value is **0**, the proxy mode is disabled for the instance. If the return value is an integer greater than **0**, such as **1**, the proxy mode is enabled for the instance.
	//
	// 	- This parameter is returned only when the instance is a [cluster](https://help.aliyun.com/document_detail/52228.html) instance.
	//
	// example:
	//
	// 1
	ProxyCount *int32 `json:"ProxyCount,omitempty" xml:"ProxyCount,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of shards.
	//
	// >  This parameter is returned only when the instance is a [cluster](https://help.aliyun.com/document_detail/52228.html) instance.
	//
	// example:
	//
	// 3
	ShardCount *int32 `json:"ShardCount,omitempty" xml:"ShardCount,omitempty"`
	// The storage type of the instance. The return value is LOCAL_SSD, which indicates [enhanced SSDs (ESSDs)](https://help.aliyun.com/document_detail/122389.html).
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
	// The ID of the host in the dedicated cluster.
	//
	// example:
	//
	// ch-bp13vf0y9gx3c****
	DedicatedHostName *string `json:"DedicatedHostName,omitempty" xml:"DedicatedHostName,omitempty"`
	// The ID of the instance.
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
	// 	- **db**: data node.
	//
	// 	- **proxy**: proxy node.
	//
	// 	- **normal**: regular node. This value is returned when the instance runs in the standard architecture.
	//
	// example:
	//
	// normal
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// The port number that is used to connect to the node.
	//
	// example:
	//
	// 3001
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The role of the node. Valid values:
	//
	// 	- **master**: master node
	//
	// 	- **slave**: replica node
	//
	// example:
	//
	// master
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The zone ID of the node.
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
