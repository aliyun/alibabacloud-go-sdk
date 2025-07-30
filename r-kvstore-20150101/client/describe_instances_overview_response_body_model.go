// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstancesOverviewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstances(v []*DescribeInstancesOverviewResponseBodyInstances) *DescribeInstancesOverviewResponseBody
	GetInstances() []*DescribeInstancesOverviewResponseBodyInstances
	SetRequestId(v string) *DescribeInstancesOverviewResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeInstancesOverviewResponseBody
	GetTotalCount() *int32
}

type DescribeInstancesOverviewResponseBody struct {
	// The queried instances.
	Instances []*DescribeInstancesOverviewResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
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

func (s DescribeInstancesOverviewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesOverviewResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstancesOverviewResponseBody) GetInstances() []*DescribeInstancesOverviewResponseBodyInstances {
	return s.Instances
}

func (s *DescribeInstancesOverviewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstancesOverviewResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeInstancesOverviewResponseBody) SetInstances(v []*DescribeInstancesOverviewResponseBodyInstances) *DescribeInstancesOverviewResponseBody {
	s.Instances = v
	return s
}

func (s *DescribeInstancesOverviewResponseBody) SetRequestId(v string) *DescribeInstancesOverviewResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstancesOverviewResponseBody) SetTotalCount(v int32) *DescribeInstancesOverviewResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeInstancesOverviewResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeInstancesOverviewResponseBodyInstances struct {
	// The architecture of the instance. Valid values:
	//
	// 	- **cluster**: cluster architecture
	//
	// 	- **standard**: standard architecture
	//
	// 	- **rwsplit**: read/write splitting architecture
	//
	// example:
	//
	// cluster
	ArchitectureType *string `json:"ArchitectureType,omitempty" xml:"ArchitectureType,omitempty"`
	// The storage capacity of the instance. Unit: MB.
	//
	// example:
	//
	// 4096
	Capacity *int64 `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// 	- **PrePaid**: subscription
	//
	// 	- **PostPaid**: pay-as-you-go
	//
	// example:
	//
	// PostPaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The internal endpoint of the instance.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****.redis.rds.aliyuncs.com
	ConnectionDomain *string `json:"ConnectionDomain,omitempty" xml:"ConnectionDomain,omitempty"`
	// The time when the instance was created.
	//
	// example:
	//
	// 2018-11-07T08:49:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the subscription instance expires.
	//
	// example:
	//
	// 2022-06-13T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The engine version of the instance. Valid values: **2.8**, **4.0**, **5.0**, **6.0**, and **7.0**.
	//
	// example:
	//
	// 4.0
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The ID of the distributed instance.
	//
	// > This parameter is returned only when the instance is a child instance of a distributed instance.
	//
	// example:
	//
	// gr-bp14rkqrhac****
	GlobalInstanceId *string `json:"GlobalInstanceId,omitempty" xml:"GlobalInstanceId,omitempty"`
	// The instance type of the instance.
	//
	// example:
	//
	// redis.logic.sharding.2g.2db.0rodb.4proxy.default
	InstanceClass *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// apitest
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
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
	// 	- **SSLModifying**: The SSL certificate of the instance is being changed.
	//
	// 	- **MajorVersionUpgrading**: The major version of the instance is being upgraded. The instance remains accessible during the upgrade.
	//
	// example:
	//
	// Normal
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// The edition of the instance. Valid values:
	//
	// 	- **Tair**: Tair (Enterprise Edition)
	//
	// 	- **Redis**: Redis Open-Source Edition
	//
	// 	- **Memcache**
	//
	// example:
	//
	// Redis
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The network type of the instance. Valid values:
	//
	// 	- **CLASSIC**: classic network
	//
	// 	- **VPC**: VPC
	//
	// example:
	//
	// CLASSIC
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The private IP address of the instance.
	//
	// > This parameter is not returned when the instance is deployed in the classic network.
	//
	// example:
	//
	// 172.16.49.***
	PrivateIp *string `json:"PrivateIp,omitempty" xml:"PrivateIp,omitempty"`
	// The region ID of the instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the instance belongs.
	//
	// example:
	//
	// rg-acfmyiu4ekp****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Instance\\"s secondary zone id.
	//
	// > This parameter is only returned when the instance has a secondary zone ID.
	//
	// example:
	//
	// cn-hangzhou-g
	SecondaryZoneId *string `json:"SecondaryZoneId,omitempty" xml:"SecondaryZoneId,omitempty"`
	// The ID of the vSwitch to which the instance is connected.
	//
	// example:
	//
	// vsw-bp1e7clcw529l773d****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the VPC.
	//
	// example:
	//
	// vpc-bp1nme44gek34slfc****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zone ID of the instance.
	//
	// example:
	//
	// cn-hangzhou-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeInstancesOverviewResponseBodyInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesOverviewResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeInstancesOverviewResponseBodyInstances) GetArchitectureType() *string {
	return s.ArchitectureType
}

func (s *DescribeInstancesOverviewResponseBodyInstances) GetCapacity() *int64 {
	return s.Capacity
}

func (s *DescribeInstancesOverviewResponseBodyInstances) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeInstancesOverviewResponseBodyInstances) GetConnectionDomain() *string {
	return s.ConnectionDomain
}

func (s *DescribeInstancesOverviewResponseBodyInstances) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeInstancesOverviewResponseBodyInstances) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeInstancesOverviewResponseBodyInstances) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeInstancesOverviewResponseBodyInstances) GetGlobalInstanceId() *string {
	return s.GlobalInstanceId
}

func (s *DescribeInstancesOverviewResponseBodyInstances) GetInstanceClass() *string {
	return s.InstanceClass
}

func (s *DescribeInstancesOverviewResponseBodyInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstancesOverviewResponseBodyInstances) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeInstancesOverviewResponseBodyInstances) GetInstanceStatus() *string {
	return s.InstanceStatus
}

func (s *DescribeInstancesOverviewResponseBodyInstances) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeInstancesOverviewResponseBodyInstances) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeInstancesOverviewResponseBodyInstances) GetPrivateIp() *string {
	return s.PrivateIp
}

func (s *DescribeInstancesOverviewResponseBodyInstances) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInstancesOverviewResponseBodyInstances) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeInstancesOverviewResponseBodyInstances) GetSecondaryZoneId() *string {
	return s.SecondaryZoneId
}

func (s *DescribeInstancesOverviewResponseBodyInstances) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeInstancesOverviewResponseBodyInstances) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeInstancesOverviewResponseBodyInstances) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeInstancesOverviewResponseBodyInstances) SetArchitectureType(v string) *DescribeInstancesOverviewResponseBodyInstances {
	s.ArchitectureType = &v
	return s
}

func (s *DescribeInstancesOverviewResponseBodyInstances) SetCapacity(v int64) *DescribeInstancesOverviewResponseBodyInstances {
	s.Capacity = &v
	return s
}

func (s *DescribeInstancesOverviewResponseBodyInstances) SetChargeType(v string) *DescribeInstancesOverviewResponseBodyInstances {
	s.ChargeType = &v
	return s
}

func (s *DescribeInstancesOverviewResponseBodyInstances) SetConnectionDomain(v string) *DescribeInstancesOverviewResponseBodyInstances {
	s.ConnectionDomain = &v
	return s
}

func (s *DescribeInstancesOverviewResponseBodyInstances) SetCreateTime(v string) *DescribeInstancesOverviewResponseBodyInstances {
	s.CreateTime = &v
	return s
}

func (s *DescribeInstancesOverviewResponseBodyInstances) SetEndTime(v string) *DescribeInstancesOverviewResponseBodyInstances {
	s.EndTime = &v
	return s
}

func (s *DescribeInstancesOverviewResponseBodyInstances) SetEngineVersion(v string) *DescribeInstancesOverviewResponseBodyInstances {
	s.EngineVersion = &v
	return s
}

func (s *DescribeInstancesOverviewResponseBodyInstances) SetGlobalInstanceId(v string) *DescribeInstancesOverviewResponseBodyInstances {
	s.GlobalInstanceId = &v
	return s
}

func (s *DescribeInstancesOverviewResponseBodyInstances) SetInstanceClass(v string) *DescribeInstancesOverviewResponseBodyInstances {
	s.InstanceClass = &v
	return s
}

func (s *DescribeInstancesOverviewResponseBodyInstances) SetInstanceId(v string) *DescribeInstancesOverviewResponseBodyInstances {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstancesOverviewResponseBodyInstances) SetInstanceName(v string) *DescribeInstancesOverviewResponseBodyInstances {
	s.InstanceName = &v
	return s
}

func (s *DescribeInstancesOverviewResponseBodyInstances) SetInstanceStatus(v string) *DescribeInstancesOverviewResponseBodyInstances {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeInstancesOverviewResponseBodyInstances) SetInstanceType(v string) *DescribeInstancesOverviewResponseBodyInstances {
	s.InstanceType = &v
	return s
}

func (s *DescribeInstancesOverviewResponseBodyInstances) SetNetworkType(v string) *DescribeInstancesOverviewResponseBodyInstances {
	s.NetworkType = &v
	return s
}

func (s *DescribeInstancesOverviewResponseBodyInstances) SetPrivateIp(v string) *DescribeInstancesOverviewResponseBodyInstances {
	s.PrivateIp = &v
	return s
}

func (s *DescribeInstancesOverviewResponseBodyInstances) SetRegionId(v string) *DescribeInstancesOverviewResponseBodyInstances {
	s.RegionId = &v
	return s
}

func (s *DescribeInstancesOverviewResponseBodyInstances) SetResourceGroupId(v string) *DescribeInstancesOverviewResponseBodyInstances {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeInstancesOverviewResponseBodyInstances) SetSecondaryZoneId(v string) *DescribeInstancesOverviewResponseBodyInstances {
	s.SecondaryZoneId = &v
	return s
}

func (s *DescribeInstancesOverviewResponseBodyInstances) SetVSwitchId(v string) *DescribeInstancesOverviewResponseBodyInstances {
	s.VSwitchId = &v
	return s
}

func (s *DescribeInstancesOverviewResponseBodyInstances) SetVpcId(v string) *DescribeInstancesOverviewResponseBodyInstances {
	s.VpcId = &v
	return s
}

func (s *DescribeInstancesOverviewResponseBodyInstances) SetZoneId(v string) *DescribeInstancesOverviewResponseBodyInstances {
	s.ZoneId = &v
	return s
}

func (s *DescribeInstancesOverviewResponseBodyInstances) Validate() error {
	return dara.Validate(s)
}
