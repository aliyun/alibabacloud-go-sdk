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
	// The list of instances. You can specify up to 2,000 instances.
	//
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

type ModifyMonitorGroupInstancesRequestInstances struct {
	// The cloud service to which the resource instance belongs. The following cloud services are supported:
	//
	// - ECS (including Alibaba Cloud and third-party hosts)
	//
	// - ApsaraDB RDS
	//
	// - AnalyticDB
	//
	// - SLB
	//
	// - VPC (Elastic IP)
	//
	// - API Gateway
	//
	// - Alibaba Cloud CDN
	//
	// - Container Service for Swarm
	//
	// - DCDN
	//
	// - Anti-DDoS
	//
	// - EIP
	//
	// - Elasticsearch
	//
	// - E-MapReduce
	//
	// - Auto Scaling
	//
	// - ApsaraDB for HBase
	//
	// - IoT Edge
	//
	// - Kubernetes pod
	//
	// - ApsaraDB for Redis (sharded cluster)
	//
	// - ApsaraDB for Redis (read/write splitting)
	//
	// - ApsaraDB for Redis (Standard Edition)
	//
	// - ApsaraDB for Memcache
	//
	// - MNS
	//
	// - ApsaraDB for MongoDB (replica set)
	//
	// - ApsaraDB for MongoDB (sharded cluster)
	//
	// - ApsaraDB for MongoDB (sharded cluster)
	//
	// - MNS topic
	//
	// - OCS (ApsaraDB for Memcache of earlier versions)
	//
	// - OpenSearch
	//
	// - OSS
	//
	// - PolarDB
	//
	// - HybridDB for MySQL
	//
	// - Internet Shared Bandwidth
	//
	// - SLS
	//
	// - VPN Gateway
	//
	// This parameter is required.
	//
	// example:
	//
	// ECS
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The ID of the resource instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-a2d5q7pm12****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// HostName
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The region ID.
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
