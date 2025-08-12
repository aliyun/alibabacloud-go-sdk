// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMonitorGroupInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v int64) *ModifyMonitorGroupInstancesRequest
	GetGroupId() *int64
	SetInstances(v []*ModifyMonitorGroupInstancesRequestInstances) *ModifyMonitorGroupInstancesRequest
	GetInstances() []*ModifyMonitorGroupInstancesRequestInstances
	SetRegionId(v string) *ModifyMonitorGroupInstancesRequest
	GetRegionId() *string
}

type ModifyMonitorGroupInstancesRequest struct {
	// The ID of the application group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// This parameter is required.
	Instances []*ModifyMonitorGroupInstancesRequestInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	RegionId  *string                                        `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyMonitorGroupInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyMonitorGroupInstancesRequest) GoString() string {
	return s.String()
}

func (s *ModifyMonitorGroupInstancesRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *ModifyMonitorGroupInstancesRequest) GetInstances() []*ModifyMonitorGroupInstancesRequestInstances {
	return s.Instances
}

func (s *ModifyMonitorGroupInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyMonitorGroupInstancesRequest) SetGroupId(v int64) *ModifyMonitorGroupInstancesRequest {
	s.GroupId = &v
	return s
}

func (s *ModifyMonitorGroupInstancesRequest) SetInstances(v []*ModifyMonitorGroupInstancesRequestInstances) *ModifyMonitorGroupInstancesRequest {
	s.Instances = v
	return s
}

func (s *ModifyMonitorGroupInstancesRequest) SetRegionId(v string) *ModifyMonitorGroupInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyMonitorGroupInstancesRequest) Validate() error {
	return dara.Validate(s)
}

type ModifyMonitorGroupInstancesRequestInstances struct {
	// The abbreviation of the name of the service to which the instances to be added to the application group belong. Valid values:
	//
	// 	- ECS: Elastic Compute Service (ECS) instances provided by Alibaba Cloud and hosts not provided by Alibaba Cloud
	//
	// 	- RDS: ApsaraDB for RDS
	//
	// 	- ADS: AnalyticDB
	//
	// 	- SLB: Server Load Balancer (SLB)
	//
	// 	- VPC: Virtual Private Cloud (VPC)
	//
	// 	- APIGATEWAY: API Gateway
	//
	// 	- CDN: Alibaba Cloud Content Delivery Network (CDN)
	//
	// 	- CS: Container Service for Swarm
	//
	// 	- DCDN: Dynamic Route for CDN
	//
	// 	- DDoS: Anti-DDoS Pro
	//
	// 	- EIP: Elastic IP Address (EIP)
	//
	// 	- ELASTICSEARCH: Elasticsearch
	//
	// 	- EMR: E-MapReduce
	//
	// 	- ESS: Auto Scaling
	//
	// 	- HBASE: ApsaraDB for Hbase
	//
	// 	- IOT_EDGE: IoT Edge
	//
	// 	- K8S_POD: pods in Container Service for Kubernetes
	//
	// 	- KVSTORE_SHARDING: ApsaraDB for Redis of the cluster architecture
	//
	// 	- KVSTORE_SPLITRW: ApsaraDB for Redis of the read/write splitting architecture
	//
	// 	- KVSTORE_STANDARD: ApsaraDB for Redis of the standard architecture
	//
	// 	- MEMCACHE: ApsaraDB for Memcache
	//
	// 	- MNS: Message Service (MNS)
	//
	// 	- MONGODB: ApsaraDB for MongoDB of the replica set architecture
	//
	// 	- MONGODB_CLUSTER: ApsaraDB for MongoDB of the cluster architecture
	//
	// 	- MONGODB_SHARDING: ApsaraDB for MongoDB of the sharded cluster architecture
	//
	// 	- MQ_TOPIC: MNS topics
	//
	// 	- OCS: ApsaraDB for Memcache of earlier versions
	//
	// 	- OPENSEARCH: Open Search
	//
	// 	- OSS: Object Storage Service (OSS)
	//
	// 	- POLARDB: PolarDB
	//
	// 	- PETADATA: HybridDB for MySQL
	//
	// 	- SCDN: Secure Content Delivery Network (SCDN)
	//
	// 	- SHAREBANDWIDTHPACKAGES: EIP Bandwidth Plan
	//
	// 	- SLS: Log Service
	//
	// 	- VPN: VPN Gateway
	//
	//     Valid values of N: 1 to 2000.
	//
	// This parameter is required.
	//
	// example:
	//
	// ECS
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The ID of the instance. Valid values of N: 1 to 2000.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-a2d5q7pm12****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the instance. Valid values of N: 1 to 2000.
	//
	// This parameter is required.
	//
	// example:
	//
	// HostName
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The ID of the region where the instance resides. Valid values of N: 1 to 2000.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyMonitorGroupInstancesRequestInstances) String() string {
	return dara.Prettify(s)
}

func (s ModifyMonitorGroupInstancesRequestInstances) GoString() string {
	return s.String()
}

func (s *ModifyMonitorGroupInstancesRequestInstances) GetCategory() *string {
	return s.Category
}

func (s *ModifyMonitorGroupInstancesRequestInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyMonitorGroupInstancesRequestInstances) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ModifyMonitorGroupInstancesRequestInstances) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyMonitorGroupInstancesRequestInstances) SetCategory(v string) *ModifyMonitorGroupInstancesRequestInstances {
	s.Category = &v
	return s
}

func (s *ModifyMonitorGroupInstancesRequestInstances) SetInstanceId(v string) *ModifyMonitorGroupInstancesRequestInstances {
	s.InstanceId = &v
	return s
}

func (s *ModifyMonitorGroupInstancesRequestInstances) SetInstanceName(v string) *ModifyMonitorGroupInstancesRequestInstances {
	s.InstanceName = &v
	return s
}

func (s *ModifyMonitorGroupInstancesRequestInstances) SetRegionId(v string) *ModifyMonitorGroupInstancesRequestInstances {
	s.RegionId = &v
	return s
}

func (s *ModifyMonitorGroupInstancesRequestInstances) Validate() error {
	return dara.Validate(s)
}
