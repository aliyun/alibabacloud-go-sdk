// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArchitectureType(v string) *DescribeInstancesRequest
	GetArchitectureType() *string
	SetChargeType(v string) *DescribeInstancesRequest
	GetChargeType() *string
	SetEditionType(v string) *DescribeInstancesRequest
	GetEditionType() *string
	SetEngineVersion(v string) *DescribeInstancesRequest
	GetEngineVersion() *string
	SetExpired(v string) *DescribeInstancesRequest
	GetExpired() *string
	SetGlobalInstance(v bool) *DescribeInstancesRequest
	GetGlobalInstance() *bool
	SetInstanceClass(v string) *DescribeInstancesRequest
	GetInstanceClass() *string
	SetInstanceIds(v string) *DescribeInstancesRequest
	GetInstanceIds() *string
	SetInstanceStatus(v string) *DescribeInstancesRequest
	GetInstanceStatus() *string
	SetInstanceType(v string) *DescribeInstancesRequest
	GetInstanceType() *string
	SetNetworkType(v string) *DescribeInstancesRequest
	GetNetworkType() *string
	SetNodeType(v string) *DescribeInstancesRequest
	GetNodeType() *string
	SetOwnerAccount(v string) *DescribeInstancesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeInstancesRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeInstancesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeInstancesRequest
	GetPageSize() *int32
	SetPrivateIp(v string) *DescribeInstancesRequest
	GetPrivateIp() *string
	SetRegionId(v string) *DescribeInstancesRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeInstancesRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeInstancesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeInstancesRequest
	GetResourceOwnerId() *int64
	SetSearchKey(v string) *DescribeInstancesRequest
	GetSearchKey() *string
	SetSecurityToken(v string) *DescribeInstancesRequest
	GetSecurityToken() *string
	SetTag(v []*DescribeInstancesRequestTag) *DescribeInstancesRequest
	GetTag() []*DescribeInstancesRequestTag
	SetVSwitchId(v string) *DescribeInstancesRequest
	GetVSwitchId() *string
	SetVpcId(v string) *DescribeInstancesRequest
	GetVpcId() *string
	SetZoneId(v string) *DescribeInstancesRequest
	GetZoneId() *string
}

type DescribeInstancesRequest struct {
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
	// The database engine version of the instance. Valid values: **2.8**, **4.0**, **5.0**, **6.0**, and **7.0**.
	//
	// Enumerated values:
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
	// Specifies whether the instance has expired. Valid values:
	//
	// 	- **true**: The instance has expired.
	//
	// 	- **false**: The instance has not expired.
	//
	// example:
	//
	// false
	Expired *string `json:"Expired,omitempty" xml:"Expired,omitempty"`
	// Specifies whether to return the child instances of distributed instances. Valid values:
	//
	// 	- **true**: Only child instances are returned.
	//
	// 	- **false**: Child instances are not returned.
	//
	// example:
	//
	// true
	GlobalInstance *bool `json:"GlobalInstance,omitempty" xml:"GlobalInstance,omitempty"`
	// The instance type of the instance. For more information, see [Instance types](https://help.aliyun.com/document_detail/107984.html).
	//
	// example:
	//
	// redis.master.small.default
	InstanceClass *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	// The IDs of the instances that you want to query.
	//
	// >  If you want to specify multiple instance IDs, separate the instance IDs with commas (,). You can specify a maximum of 30 instance IDs in a single request.
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
	// 	- **Unavailable**: The instance is suspended.
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
	// The database engine. Valid values:
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
	// The network type. Valid values:
	//
	// 	- **CLASSIC**
	//
	// 	- **VPC**
	//
	// example:
	//
	// CLASSIC
	NetworkType  *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	NodeType     *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Pages start from page **1**. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Maximum value: **50**. Default value: **30**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The private IP address of the instance.
	//
	// example:
	//
	// 172.16.49.***
	PrivateIp *string `json:"PrivateIp,omitempty" xml:"PrivateIp,omitempty"`
	// The region ID of the instance.
	//
	// > When you call this operation and specify the **Tag*	- parameter, you must also specify this parameter.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the instance belongs.
	//
	// > You can query resource group IDs by using the Tair (Redis OSS-compatible) console or by calling the [ListResourceGroups](https://help.aliyun.com/document_detail/158855.html) operation. For more information, see [View basic information of a resource group](https://help.aliyun.com/document_detail/151181.html).
	//
	// example:
	//
	// rg-acfmyiu4ekp****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The keyword used for fuzzy search. The keyword can be based on an instance name or an instance ID.
	//
	// example:
	//
	// apitest
	SearchKey     *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The tags of the instance.
	Tag []*DescribeInstancesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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
	// cn-hongkong-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstancesRequest) GetArchitectureType() *string {
	return s.ArchitectureType
}

func (s *DescribeInstancesRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeInstancesRequest) GetEditionType() *string {
	return s.EditionType
}

func (s *DescribeInstancesRequest) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeInstancesRequest) GetExpired() *string {
	return s.Expired
}

func (s *DescribeInstancesRequest) GetGlobalInstance() *bool {
	return s.GlobalInstance
}

func (s *DescribeInstancesRequest) GetInstanceClass() *string {
	return s.InstanceClass
}

func (s *DescribeInstancesRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *DescribeInstancesRequest) GetInstanceStatus() *string {
	return s.InstanceStatus
}

func (s *DescribeInstancesRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeInstancesRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeInstancesRequest) GetNodeType() *string {
	return s.NodeType
}

func (s *DescribeInstancesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeInstancesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeInstancesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeInstancesRequest) GetPrivateIp() *string {
	return s.PrivateIp
}

func (s *DescribeInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInstancesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeInstancesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeInstancesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeInstancesRequest) GetSearchKey() *string {
	return s.SearchKey
}

func (s *DescribeInstancesRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeInstancesRequest) GetTag() []*DescribeInstancesRequestTag {
	return s.Tag
}

func (s *DescribeInstancesRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeInstancesRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeInstancesRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeInstancesRequest) SetArchitectureType(v string) *DescribeInstancesRequest {
	s.ArchitectureType = &v
	return s
}

func (s *DescribeInstancesRequest) SetChargeType(v string) *DescribeInstancesRequest {
	s.ChargeType = &v
	return s
}

func (s *DescribeInstancesRequest) SetEditionType(v string) *DescribeInstancesRequest {
	s.EditionType = &v
	return s
}

func (s *DescribeInstancesRequest) SetEngineVersion(v string) *DescribeInstancesRequest {
	s.EngineVersion = &v
	return s
}

func (s *DescribeInstancesRequest) SetExpired(v string) *DescribeInstancesRequest {
	s.Expired = &v
	return s
}

func (s *DescribeInstancesRequest) SetGlobalInstance(v bool) *DescribeInstancesRequest {
	s.GlobalInstance = &v
	return s
}

func (s *DescribeInstancesRequest) SetInstanceClass(v string) *DescribeInstancesRequest {
	s.InstanceClass = &v
	return s
}

func (s *DescribeInstancesRequest) SetInstanceIds(v string) *DescribeInstancesRequest {
	s.InstanceIds = &v
	return s
}

func (s *DescribeInstancesRequest) SetInstanceStatus(v string) *DescribeInstancesRequest {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeInstancesRequest) SetInstanceType(v string) *DescribeInstancesRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeInstancesRequest) SetNetworkType(v string) *DescribeInstancesRequest {
	s.NetworkType = &v
	return s
}

func (s *DescribeInstancesRequest) SetNodeType(v string) *DescribeInstancesRequest {
	s.NodeType = &v
	return s
}

func (s *DescribeInstancesRequest) SetOwnerAccount(v string) *DescribeInstancesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeInstancesRequest) SetOwnerId(v int64) *DescribeInstancesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeInstancesRequest) SetPageNumber(v int32) *DescribeInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeInstancesRequest) SetPageSize(v int32) *DescribeInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInstancesRequest) SetPrivateIp(v string) *DescribeInstancesRequest {
	s.PrivateIp = &v
	return s
}

func (s *DescribeInstancesRequest) SetRegionId(v string) *DescribeInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInstancesRequest) SetResourceGroupId(v string) *DescribeInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeInstancesRequest) SetResourceOwnerAccount(v string) *DescribeInstancesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeInstancesRequest) SetResourceOwnerId(v int64) *DescribeInstancesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeInstancesRequest) SetSearchKey(v string) *DescribeInstancesRequest {
	s.SearchKey = &v
	return s
}

func (s *DescribeInstancesRequest) SetSecurityToken(v string) *DescribeInstancesRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeInstancesRequest) SetTag(v []*DescribeInstancesRequestTag) *DescribeInstancesRequest {
	s.Tag = v
	return s
}

func (s *DescribeInstancesRequest) SetVSwitchId(v string) *DescribeInstancesRequest {
	s.VSwitchId = &v
	return s
}

func (s *DescribeInstancesRequest) SetVpcId(v string) *DescribeInstancesRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeInstancesRequest) SetZoneId(v string) *DescribeInstancesRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeInstancesRequest) Validate() error {
	return dara.Validate(s)
}

type DescribeInstancesRequestTag struct {
	// The tag key. A tag is a key-value pair.
	//
	// >  A maximum of five key-value pairs can be specified at a time.
	//
	// example:
	//
	// Storage type
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// Local disk
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeInstancesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeInstancesRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeInstancesRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeInstancesRequestTag) SetKey(v string) *DescribeInstancesRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeInstancesRequestTag) SetValue(v string) *DescribeInstancesRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeInstancesRequestTag) Validate() error {
	return dara.Validate(s)
}
