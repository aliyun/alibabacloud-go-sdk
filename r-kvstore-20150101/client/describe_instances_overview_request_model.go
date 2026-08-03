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
	SetNodeType(v string) *DescribeInstancesOverviewRequest
	GetNodeType() *string
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
	// The architecture type. Valid values:
	//
	// 	- **cluster**: Cluster Edition.
	//
	// 	- **standard**: Standard Edition.
	//
	// 	- **rwsplit**: read/write splitting edition.
	//
	// example:
	//
	// standard
	ArchitectureType *string `json:"ArchitectureType,omitempty" xml:"ArchitectureType,omitempty"`
	// The billing method. Valid values:
	//
	// 	- **PrePaid**: subscription.
	//
	// 	- **PostPaid**: pay-as-you-go.
	//
	// example:
	//
	// PostPaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The edition type. Valid values:
	//
	// 	- **Community**: Community Edition.
	//
	// 	- **Enterprise**: Enterprise Edition.
	//
	// example:
	//
	// Enterprise
	EditionType *string `json:"EditionType,omitempty" xml:"EditionType,omitempty"`
	// The Redis-compatible engine version of the instance. Valid values: **2.8**, **4.0**, **5.0**, **6.0**, and **7.0**.
	//
	// example:
	//
	// 6.0
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The instance type. For more information, see [Instance types](https://help.aliyun.com/document_detail/107984.html).
	//
	// example:
	//
	// redis.master.small.default
	InstanceClass *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	// The IDs of the instances that you want to query.
	//
	// > By default, all instances under the current account are queried. To specify multiple instance IDs, separate them with commas (,).
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The instance status. Valid values:
	//
	// 	- **Normal**: The instance is running.
	//
	// 	- **Creating**: The instance is being created.
	//
	// 	- **Changing**: The instance is being modified.
	//
	// 	- **Inactive**: The instance is disabled.
	//
	// 	- **Flushing**: The instance is being purged.
	//
	// 	- **Released**: The instance is released.
	//
	// 	- **Transforming**: The instance is being transformed.
	//
	// 	- **Migrating**: The instance is being migrated.
	//
	// 	- **BackupRecovering**: The instance is being restored from a backup.
	//
	// 	- **MinorVersionUpgrading**: A minor version upgrade is in progress.
	//
	// 	- **NetworkModifying**: The network configuration is being modified.
	//
	// 	- **SSLModifying**: The SSL configuration is being modified.
	//
	// 	- **MajorVersionUpgrading**: A major engine version upgrade is in progress. The instance can be accessed normally.
	//
	// > For more information about instance statuses, see [Instance states and impacts](https://help.aliyun.com/document_detail/200740.html).
	//
	// example:
	//
	// Normal
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// The instance type. Valid values:
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
	// The network type. Valid values:
	//
	// 	- **CLASSIC**: classic network.
	//
	// 	- **VPC**: virtual private cloud (VPC).
	//
	// example:
	//
	// CLASSIC
	NetworkType  *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	NodeType     *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The private IP address.
	//
	// example:
	//
	// 172.16.49.***
	PrivateIp *string `json:"PrivateIp,omitempty" xml:"PrivateIp,omitempty"`
	// The region ID of the instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/473763.html) operation to query the region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// > You can invoke the [ListResourceGroups](https://help.aliyun.com/document_detail/158855.html) operation or use the console to obtain the resource group ID. Related operations, see [View basic information of a resource group](https://help.aliyun.com/document_detail/151181.html).
	//
	// example:
	//
	// rg-acfmyiu4ekp****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The keyword used for fuzzy search by instance ID or instance description.
	//
	// example:
	//
	// apitest
	SearchKey     *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-bp1e7clcw529l773d****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-bp1nme44gek34slfc****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zone ID.
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

func (s *DescribeInstancesOverviewRequest) GetNodeType() *string {
	return s.NodeType
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

func (s *DescribeInstancesOverviewRequest) SetNodeType(v string) *DescribeInstancesOverviewRequest {
	s.NodeType = &v
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
