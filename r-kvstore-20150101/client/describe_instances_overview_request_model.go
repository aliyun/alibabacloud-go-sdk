// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstancesOverviewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArchitectureType(v string) *DescribeInstancesOverviewRequest
	GetArchitectureType() *string
	SetChargeType(v string) *DescribeInstancesOverviewRequest
	GetChargeType() *string
	SetEditionType(v string) *DescribeInstancesOverviewRequest
	GetEditionType() *string
	SetEngineVersion(v string) *DescribeInstancesOverviewRequest
	GetEngineVersion() *string
	SetInstanceClass(v string) *DescribeInstancesOverviewRequest
	GetInstanceClass() *string
	SetInstanceIds(v string) *DescribeInstancesOverviewRequest
	GetInstanceIds() *string
	SetInstanceStatus(v string) *DescribeInstancesOverviewRequest
	GetInstanceStatus() *string
	SetInstanceType(v string) *DescribeInstancesOverviewRequest
	GetInstanceType() *string
	SetNetworkType(v string) *DescribeInstancesOverviewRequest
	GetNetworkType() *string
	SetOwnerAccount(v string) *DescribeInstancesOverviewRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeInstancesOverviewRequest
	GetOwnerId() *int64
	SetPrivateIp(v string) *DescribeInstancesOverviewRequest
	GetPrivateIp() *string
	SetRegionId(v string) *DescribeInstancesOverviewRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeInstancesOverviewRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeInstancesOverviewRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeInstancesOverviewRequest
	GetResourceOwnerId() *int64
	SetSearchKey(v string) *DescribeInstancesOverviewRequest
	GetSearchKey() *string
	SetSecurityToken(v string) *DescribeInstancesOverviewRequest
	GetSecurityToken() *string
	SetVSwitchId(v string) *DescribeInstancesOverviewRequest
	GetVSwitchId() *string
	SetVpcId(v string) *DescribeInstancesOverviewRequest
	GetVpcId() *string
	SetZoneId(v string) *DescribeInstancesOverviewRequest
	GetZoneId() *string
}

type DescribeInstancesOverviewRequest struct {
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
	// standard
	ArchitectureType *string `json:"ArchitectureType,omitempty" xml:"ArchitectureType,omitempty"`
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
	// The edition of the instance. Valid values:
	//
	// 	- **Community**: Redis Open-Source Edition
	//
	// 	- **Enterprise**: Tair (Enterprise Edition)
	//
	// example:
	//
	// Enterprise
	EditionType *string `json:"EditionType,omitempty" xml:"EditionType,omitempty"`
	// The engine version of the instance. Valid values: **2.8**, **4.0**, **5.0**, **6.0**, and **7.0**.
	//
	// Valid values:
	//
	// 	- 1.0
	//
	// 	- 2.8
	//
	// 	- 4.0
	//
	// 	- 5.0
	//
	// 	- 6.0
	//
	// 	- 7.0
	//
	// example:
	//
	// 4.0
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The instance type of the instance. For more information, see [Instance types](https://help.aliyun.com/document_detail/107984.html).
	//
	// example:
	//
	// redis.master.small.default
	InstanceClass *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	// The IDs of instances.
	//
	// > By default, all instances that belong to this account are queried. If you specify multiple instance IDs, separate the instance IDs with commas (,).
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
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
	// > For more information about instance states, see [Instance states and impacts](https://help.aliyun.com/document_detail/200740.html).
	//
	// example:
	//
	// Normal
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// The category of the instance. Valid values:
	//
	// 	- **Tair**
	//
	// 	- **Redis**
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
	// 	- **VPC**: Virtual Private Cloud (VPC)
	//
	// example:
	//
	// CLASSIC
	NetworkType  *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The private IP address of the instance.
	//
	// example:
	//
	// 172.16.49.***
	PrivateIp *string `json:"PrivateIp,omitempty" xml:"PrivateIp,omitempty"`
	// The ID of the region in which the instances you want to query reside. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/473763.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the instances you want to query belong.
	//
	// > You can query resource group IDs by using the Tair (Redis OSS-compatible) console or by calling the [ListResourceGroups](https://help.aliyun.com/document_detail/158855.html) operation. For more information, see [View basic information of a resource group](https://help.aliyun.com/document_detail/151181.html).
	//
	// example:
	//
	// rg-acfmyiu4ekp****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The keyword used for fuzzy search. The keyword can be based on an instance ID or an instance description.
	//
	// example:
	//
	// apitest
	SearchKey     *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The ID of the vSwitch.
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

func (s DescribeInstancesOverviewRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesOverviewRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstancesOverviewRequest) GetArchitectureType() *string {
	return s.ArchitectureType
}

func (s *DescribeInstancesOverviewRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeInstancesOverviewRequest) GetEditionType() *string {
	return s.EditionType
}

func (s *DescribeInstancesOverviewRequest) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeInstancesOverviewRequest) GetInstanceClass() *string {
	return s.InstanceClass
}

func (s *DescribeInstancesOverviewRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *DescribeInstancesOverviewRequest) GetInstanceStatus() *string {
	return s.InstanceStatus
}

func (s *DescribeInstancesOverviewRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeInstancesOverviewRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeInstancesOverviewRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeInstancesOverviewRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeInstancesOverviewRequest) GetPrivateIp() *string {
	return s.PrivateIp
}

func (s *DescribeInstancesOverviewRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInstancesOverviewRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeInstancesOverviewRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeInstancesOverviewRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeInstancesOverviewRequest) GetSearchKey() *string {
	return s.SearchKey
}

func (s *DescribeInstancesOverviewRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeInstancesOverviewRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeInstancesOverviewRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeInstancesOverviewRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeInstancesOverviewRequest) SetArchitectureType(v string) *DescribeInstancesOverviewRequest {
	s.ArchitectureType = &v
	return s
}

func (s *DescribeInstancesOverviewRequest) SetChargeType(v string) *DescribeInstancesOverviewRequest {
	s.ChargeType = &v
	return s
}

func (s *DescribeInstancesOverviewRequest) SetEditionType(v string) *DescribeInstancesOverviewRequest {
	s.EditionType = &v
	return s
}

func (s *DescribeInstancesOverviewRequest) SetEngineVersion(v string) *DescribeInstancesOverviewRequest {
	s.EngineVersion = &v
	return s
}

func (s *DescribeInstancesOverviewRequest) SetInstanceClass(v string) *DescribeInstancesOverviewRequest {
	s.InstanceClass = &v
	return s
}

func (s *DescribeInstancesOverviewRequest) SetInstanceIds(v string) *DescribeInstancesOverviewRequest {
	s.InstanceIds = &v
	return s
}

func (s *DescribeInstancesOverviewRequest) SetInstanceStatus(v string) *DescribeInstancesOverviewRequest {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeInstancesOverviewRequest) SetInstanceType(v string) *DescribeInstancesOverviewRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeInstancesOverviewRequest) SetNetworkType(v string) *DescribeInstancesOverviewRequest {
	s.NetworkType = &v
	return s
}

func (s *DescribeInstancesOverviewRequest) SetOwnerAccount(v string) *DescribeInstancesOverviewRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeInstancesOverviewRequest) SetOwnerId(v int64) *DescribeInstancesOverviewRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeInstancesOverviewRequest) SetPrivateIp(v string) *DescribeInstancesOverviewRequest {
	s.PrivateIp = &v
	return s
}

func (s *DescribeInstancesOverviewRequest) SetRegionId(v string) *DescribeInstancesOverviewRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInstancesOverviewRequest) SetResourceGroupId(v string) *DescribeInstancesOverviewRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeInstancesOverviewRequest) SetResourceOwnerAccount(v string) *DescribeInstancesOverviewRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeInstancesOverviewRequest) SetResourceOwnerId(v int64) *DescribeInstancesOverviewRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeInstancesOverviewRequest) SetSearchKey(v string) *DescribeInstancesOverviewRequest {
	s.SearchKey = &v
	return s
}

func (s *DescribeInstancesOverviewRequest) SetSecurityToken(v string) *DescribeInstancesOverviewRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeInstancesOverviewRequest) SetVSwitchId(v string) *DescribeInstancesOverviewRequest {
	s.VSwitchId = &v
	return s
}

func (s *DescribeInstancesOverviewRequest) SetVpcId(v string) *DescribeInstancesOverviewRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeInstancesOverviewRequest) SetZoneId(v string) *DescribeInstancesOverviewRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeInstancesOverviewRequest) Validate() error {
	return dara.Validate(s)
}
