// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstancesForCloneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeDBInstancesForCloneRequest
	GetClientToken() *string
	SetConnectionMode(v string) *DescribeDBInstancesForCloneRequest
	GetConnectionMode() *string
	SetCurrentInstanceId(v string) *DescribeDBInstancesForCloneRequest
	GetCurrentInstanceId() *string
	SetDBInstanceClass(v string) *DescribeDBInstancesForCloneRequest
	GetDBInstanceClass() *string
	SetDBInstanceId(v string) *DescribeDBInstancesForCloneRequest
	GetDBInstanceId() *string
	SetDBInstanceStatus(v string) *DescribeDBInstancesForCloneRequest
	GetDBInstanceStatus() *string
	SetDBInstanceType(v string) *DescribeDBInstancesForCloneRequest
	GetDBInstanceType() *string
	SetEngine(v string) *DescribeDBInstancesForCloneRequest
	GetEngine() *string
	SetEngineVersion(v string) *DescribeDBInstancesForCloneRequest
	GetEngineVersion() *string
	SetExpired(v string) *DescribeDBInstancesForCloneRequest
	GetExpired() *string
	SetInstanceNetworkType(v string) *DescribeDBInstancesForCloneRequest
	GetInstanceNetworkType() *string
	SetNodeType(v string) *DescribeDBInstancesForCloneRequest
	GetNodeType() *string
	SetOwnerAccount(v string) *DescribeDBInstancesForCloneRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDBInstancesForCloneRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeDBInstancesForCloneRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDBInstancesForCloneRequest
	GetPageSize() *int32
	SetPayType(v string) *DescribeDBInstancesForCloneRequest
	GetPayType() *string
	SetRegionId(v string) *DescribeDBInstancesForCloneRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeDBInstancesForCloneRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeDBInstancesForCloneRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBInstancesForCloneRequest
	GetResourceOwnerId() *int64
	SetSearchKey(v string) *DescribeDBInstancesForCloneRequest
	GetSearchKey() *string
	SetVSwitchId(v string) *DescribeDBInstancesForCloneRequest
	GetVSwitchId() *string
	SetVpcId(v string) *DescribeDBInstancesForCloneRequest
	GetVpcId() *string
	SetZoneId(v string) *DescribeDBInstancesForCloneRequest
	GetZoneId() *string
	SetProxyId(v string) *DescribeDBInstancesForCloneRequest
	GetProxyId() *string
}

type DescribeDBInstancesForCloneRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCzxxxxxxxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The connection mode of the instance. Valid values:
	//
	// 	- **Standard**: standard mode
	//
	// 	- **Safe**: database proxy mode
	//
	// By default, this operation queries the instances that use any of the supported connection modes.
	//
	// example:
	//
	// Standard
	ConnectionMode *string `json:"ConnectionMode,omitempty" xml:"ConnectionMode,omitempty"`
	// The ID of the current instance.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxxxxx
	CurrentInstanceId *string `json:"CurrentInstanceId,omitempty" xml:"CurrentInstanceId,omitempty"`
	// The instance type of the instance. For more information, see [Instance types](https://help.aliyun.com/document_detail/26312.html).
	//
	// example:
	//
	// mysql.n1.micro.1
	DBInstanceClass *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The status of the instance. For more information, see [Instance state table](https://help.aliyun.com/document_detail/26315.html).
	//
	// example:
	//
	// Running
	DBInstanceStatus *string `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty"`
	// The role of the instance that you want to query. Valid values:
	//
	// 	- **Primary**: primary instance
	//
	// 	- **Readonly**: read-only instance
	//
	// 	- **Guard**: disaster recovery instance
	//
	// 	- **Temp**: temporary instance
	//
	// By default, this operation queries the instances of all roles.
	//
	// example:
	//
	// Primary
	DBInstanceType *string `json:"DBInstanceType,omitempty" xml:"DBInstanceType,omitempty"`
	// The database engine of the instance. Valid values:
	//
	// 	- MySQL
	//
	// 	- SQLServer
	//
	// 	- PostgreSQL
	//
	// 	- PPAS
	//
	// 	- MariaDB
	//
	// By default, this operation queries the instances that run any of the supported database engine types.
	//
	// example:
	//
	// MySQL
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The version of the database engine.
	//
	// example:
	//
	// 5.7
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// Specifies whether the instance expires. Valid values:
	//
	// 	- **True**: queries the instances that have expired.
	//
	// 	- **False**: does not query instances that have expired.
	//
	// example:
	//
	// True
	Expired *string `json:"Expired,omitempty" xml:"Expired,omitempty"`
	// The network type of the instance. Valid values:
	//
	// 	- **Classic**
	//
	// 	- **VPC**
	//
	// example:
	//
	// Classic
	InstanceNetworkType *string `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	// The type of the database node. Valid values:
	//
	// 	- **Master**: the primary node
	//
	// 	- **Slave**: the secondary node
	//
	// example:
	//
	// Master
	NodeType     *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values: **1 to 100**.
	//
	// Default value: **30**.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// 	- **Postpaid**: pay-as-you-go
	//
	// 	- **Prepaid**: subscription
	//
	// By default, this operation queries the instances that use any of the supported billing methods.
	//
	// example:
	//
	// Postpaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmy*****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The keyword that is used for the search. The keyword can be part of an instance ID or an instance description.
	//
	// example:
	//
	// rm-uf6w
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	// The ID of the vSwitch.
	//
	// example:
	//
	// vsw-j6csw46bgrgkxxxxxxxxxx
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the virtual private cloud (VPC).
	//
	// example:
	//
	// vpc-j6cjvqms29yxxxxxxxxxx
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zone ID of the instance.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	// The ID of the proxy mode.
	//
	// example:
	//
	// API
	ProxyId *string `json:"proxyId,omitempty" xml:"proxyId,omitempty"`
}

func (s DescribeDBInstancesForCloneRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancesForCloneRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesForCloneRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeDBInstancesForCloneRequest) GetConnectionMode() *string {
	return s.ConnectionMode
}

func (s *DescribeDBInstancesForCloneRequest) GetCurrentInstanceId() *string {
	return s.CurrentInstanceId
}

func (s *DescribeDBInstancesForCloneRequest) GetDBInstanceClass() *string {
	return s.DBInstanceClass
}

func (s *DescribeDBInstancesForCloneRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstancesForCloneRequest) GetDBInstanceStatus() *string {
	return s.DBInstanceStatus
}

func (s *DescribeDBInstancesForCloneRequest) GetDBInstanceType() *string {
	return s.DBInstanceType
}

func (s *DescribeDBInstancesForCloneRequest) GetEngine() *string {
	return s.Engine
}

func (s *DescribeDBInstancesForCloneRequest) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeDBInstancesForCloneRequest) GetExpired() *string {
	return s.Expired
}

func (s *DescribeDBInstancesForCloneRequest) GetInstanceNetworkType() *string {
	return s.InstanceNetworkType
}

func (s *DescribeDBInstancesForCloneRequest) GetNodeType() *string {
	return s.NodeType
}

func (s *DescribeDBInstancesForCloneRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDBInstancesForCloneRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBInstancesForCloneRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDBInstancesForCloneRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDBInstancesForCloneRequest) GetPayType() *string {
	return s.PayType
}

func (s *DescribeDBInstancesForCloneRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBInstancesForCloneRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDBInstancesForCloneRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBInstancesForCloneRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBInstancesForCloneRequest) GetSearchKey() *string {
	return s.SearchKey
}

func (s *DescribeDBInstancesForCloneRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeDBInstancesForCloneRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeDBInstancesForCloneRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeDBInstancesForCloneRequest) GetProxyId() *string {
	return s.ProxyId
}

func (s *DescribeDBInstancesForCloneRequest) SetClientToken(v string) *DescribeDBInstancesForCloneRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeDBInstancesForCloneRequest) SetConnectionMode(v string) *DescribeDBInstancesForCloneRequest {
	s.ConnectionMode = &v
	return s
}

func (s *DescribeDBInstancesForCloneRequest) SetCurrentInstanceId(v string) *DescribeDBInstancesForCloneRequest {
	s.CurrentInstanceId = &v
	return s
}

func (s *DescribeDBInstancesForCloneRequest) SetDBInstanceClass(v string) *DescribeDBInstancesForCloneRequest {
	s.DBInstanceClass = &v
	return s
}

func (s *DescribeDBInstancesForCloneRequest) SetDBInstanceId(v string) *DescribeDBInstancesForCloneRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstancesForCloneRequest) SetDBInstanceStatus(v string) *DescribeDBInstancesForCloneRequest {
	s.DBInstanceStatus = &v
	return s
}

func (s *DescribeDBInstancesForCloneRequest) SetDBInstanceType(v string) *DescribeDBInstancesForCloneRequest {
	s.DBInstanceType = &v
	return s
}

func (s *DescribeDBInstancesForCloneRequest) SetEngine(v string) *DescribeDBInstancesForCloneRequest {
	s.Engine = &v
	return s
}

func (s *DescribeDBInstancesForCloneRequest) SetEngineVersion(v string) *DescribeDBInstancesForCloneRequest {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDBInstancesForCloneRequest) SetExpired(v string) *DescribeDBInstancesForCloneRequest {
	s.Expired = &v
	return s
}

func (s *DescribeDBInstancesForCloneRequest) SetInstanceNetworkType(v string) *DescribeDBInstancesForCloneRequest {
	s.InstanceNetworkType = &v
	return s
}

func (s *DescribeDBInstancesForCloneRequest) SetNodeType(v string) *DescribeDBInstancesForCloneRequest {
	s.NodeType = &v
	return s
}

func (s *DescribeDBInstancesForCloneRequest) SetOwnerAccount(v string) *DescribeDBInstancesForCloneRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBInstancesForCloneRequest) SetOwnerId(v int64) *DescribeDBInstancesForCloneRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBInstancesForCloneRequest) SetPageNumber(v int32) *DescribeDBInstancesForCloneRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBInstancesForCloneRequest) SetPageSize(v int32) *DescribeDBInstancesForCloneRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDBInstancesForCloneRequest) SetPayType(v string) *DescribeDBInstancesForCloneRequest {
	s.PayType = &v
	return s
}

func (s *DescribeDBInstancesForCloneRequest) SetRegionId(v string) *DescribeDBInstancesForCloneRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstancesForCloneRequest) SetResourceGroupId(v string) *DescribeDBInstancesForCloneRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBInstancesForCloneRequest) SetResourceOwnerAccount(v string) *DescribeDBInstancesForCloneRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBInstancesForCloneRequest) SetResourceOwnerId(v int64) *DescribeDBInstancesForCloneRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBInstancesForCloneRequest) SetSearchKey(v string) *DescribeDBInstancesForCloneRequest {
	s.SearchKey = &v
	return s
}

func (s *DescribeDBInstancesForCloneRequest) SetVSwitchId(v string) *DescribeDBInstancesForCloneRequest {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBInstancesForCloneRequest) SetVpcId(v string) *DescribeDBInstancesForCloneRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeDBInstancesForCloneRequest) SetZoneId(v string) *DescribeDBInstancesForCloneRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeDBInstancesForCloneRequest) SetProxyId(v string) *DescribeDBInstancesForCloneRequest {
	s.ProxyId = &v
	return s
}

func (s *DescribeDBInstancesForCloneRequest) Validate() error {
	return dara.Validate(s)
}
