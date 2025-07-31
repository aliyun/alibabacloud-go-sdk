// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstancesOverviewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChargeType(v string) *DescribeDBInstancesOverviewRequest
	GetChargeType() *string
	SetEngineVersion(v string) *DescribeDBInstancesOverviewRequest
	GetEngineVersion() *string
	SetInstanceClass(v string) *DescribeDBInstancesOverviewRequest
	GetInstanceClass() *string
	SetInstanceIds(v string) *DescribeDBInstancesOverviewRequest
	GetInstanceIds() *string
	SetInstanceStatus(v string) *DescribeDBInstancesOverviewRequest
	GetInstanceStatus() *string
	SetInstanceType(v string) *DescribeDBInstancesOverviewRequest
	GetInstanceType() *string
	SetNetworkType(v string) *DescribeDBInstancesOverviewRequest
	GetNetworkType() *string
	SetOwnerAccount(v string) *DescribeDBInstancesOverviewRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDBInstancesOverviewRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeDBInstancesOverviewRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeDBInstancesOverviewRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeDBInstancesOverviewRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBInstancesOverviewRequest
	GetResourceOwnerId() *int64
	SetShowTags(v bool) *DescribeDBInstancesOverviewRequest
	GetShowTags() *bool
	SetVSwitchId(v string) *DescribeDBInstancesOverviewRequest
	GetVSwitchId() *string
	SetVpcId(v string) *DescribeDBInstancesOverviewRequest
	GetVpcId() *string
	SetZoneId(v string) *DescribeDBInstancesOverviewRequest
	GetZoneId() *string
}

type DescribeDBInstancesOverviewRequest struct {
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
	// The database engine version of the instance. Valid values: **5.0**, **4.4**, **4.2**, **4.0**, and **3.4**.
	//
	// example:
	//
	// 4.0
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The instance type. The instance type varies based on the instance architecture. For more information about instance types supported by different instance architectures, see the following references:
	//
	// 	- [Standalone instance types](https://help.aliyun.com/document_detail/311407.html)
	//
	// 	- [Replica set instance types](https://help.aliyun.com/document_detail/311410.html)
	//
	// 	- [Sharded cluster instance types](https://help.aliyun.com/document_detail/311414.html)
	//
	// example:
	//
	// dds.mongo.2xlarge
	InstanceClass *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	// The ID of the instance for which you want to query the overview information.
	//
	// >
	//
	// 	- If you do not specify this parameter, the overview information of all instances within this account is queried.
	//
	// 	- Separate the instance IDs with commas (,).
	//
	// example:
	//
	// dds-bp12c5b040dc****
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The state of the instance. For more information about valid values, see [Instance states](https://help.aliyun.com/document_detail/63870.html).
	//
	// example:
	//
	// Running
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// The architecture of the instance. Valid values:
	//
	// 	- **sharding**: sharded cluster instance
	//
	// 	- **replicate**: replica set or standalone instance
	//
	// >
	//
	// 	- To query the overview information of a sharded cluster instance, you must set the parameter to **sharding**.
	//
	// 	- If you do not specify this parameter, the overview information of all instances within this account is queried.
	//
	// example:
	//
	// replicate
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The network type of the instance. Valid values:
	//
	// 	- **Classic**: classic network
	//
	// 	- **VPC**: Virtual Private Cloud (VPC)
	//
	// example:
	//
	// Classic
	NetworkType  *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/61933.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the instance belongs. For more information, see [View the basic information of a resource group](https://help.aliyun.com/document_detail/151181.html).
	//
	// example:
	//
	// rg-acfmyiu4ekp****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specifies whether to display instance tags. Default value: False.
	//
	// example:
	//
	// false
	ShowTags *bool `json:"ShowTags,omitempty" xml:"ShowTags,omitempty"`
	// The ID of the vSwitch to which the instance is connected.
	//
	// example:
	//
	// vsw-bp1e7clcw529l773d****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the VPC in which the instance is deployed.
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

func (s DescribeDBInstancesOverviewRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancesOverviewRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesOverviewRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeDBInstancesOverviewRequest) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeDBInstancesOverviewRequest) GetInstanceClass() *string {
	return s.InstanceClass
}

func (s *DescribeDBInstancesOverviewRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *DescribeDBInstancesOverviewRequest) GetInstanceStatus() *string {
	return s.InstanceStatus
}

func (s *DescribeDBInstancesOverviewRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeDBInstancesOverviewRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeDBInstancesOverviewRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDBInstancesOverviewRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBInstancesOverviewRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBInstancesOverviewRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDBInstancesOverviewRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBInstancesOverviewRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBInstancesOverviewRequest) GetShowTags() *bool {
	return s.ShowTags
}

func (s *DescribeDBInstancesOverviewRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeDBInstancesOverviewRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeDBInstancesOverviewRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeDBInstancesOverviewRequest) SetChargeType(v string) *DescribeDBInstancesOverviewRequest {
	s.ChargeType = &v
	return s
}

func (s *DescribeDBInstancesOverviewRequest) SetEngineVersion(v string) *DescribeDBInstancesOverviewRequest {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDBInstancesOverviewRequest) SetInstanceClass(v string) *DescribeDBInstancesOverviewRequest {
	s.InstanceClass = &v
	return s
}

func (s *DescribeDBInstancesOverviewRequest) SetInstanceIds(v string) *DescribeDBInstancesOverviewRequest {
	s.InstanceIds = &v
	return s
}

func (s *DescribeDBInstancesOverviewRequest) SetInstanceStatus(v string) *DescribeDBInstancesOverviewRequest {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeDBInstancesOverviewRequest) SetInstanceType(v string) *DescribeDBInstancesOverviewRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeDBInstancesOverviewRequest) SetNetworkType(v string) *DescribeDBInstancesOverviewRequest {
	s.NetworkType = &v
	return s
}

func (s *DescribeDBInstancesOverviewRequest) SetOwnerAccount(v string) *DescribeDBInstancesOverviewRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBInstancesOverviewRequest) SetOwnerId(v int64) *DescribeDBInstancesOverviewRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBInstancesOverviewRequest) SetRegionId(v string) *DescribeDBInstancesOverviewRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstancesOverviewRequest) SetResourceGroupId(v string) *DescribeDBInstancesOverviewRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBInstancesOverviewRequest) SetResourceOwnerAccount(v string) *DescribeDBInstancesOverviewRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBInstancesOverviewRequest) SetResourceOwnerId(v int64) *DescribeDBInstancesOverviewRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBInstancesOverviewRequest) SetShowTags(v bool) *DescribeDBInstancesOverviewRequest {
	s.ShowTags = &v
	return s
}

func (s *DescribeDBInstancesOverviewRequest) SetVSwitchId(v string) *DescribeDBInstancesOverviewRequest {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBInstancesOverviewRequest) SetVpcId(v string) *DescribeDBInstancesOverviewRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeDBInstancesOverviewRequest) SetZoneId(v string) *DescribeDBInstancesOverviewRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeDBInstancesOverviewRequest) Validate() error {
	return dara.Validate(s)
}
