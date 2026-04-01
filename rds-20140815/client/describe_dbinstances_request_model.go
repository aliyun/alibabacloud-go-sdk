// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *DescribeDBInstancesRequest
	GetCategory() *string
	SetClientToken(v string) *DescribeDBInstancesRequest
	GetClientToken() *string
	SetConnectionMode(v string) *DescribeDBInstancesRequest
	GetConnectionMode() *string
	SetConnectionString(v string) *DescribeDBInstancesRequest
	GetConnectionString() *string
	SetDBInstanceClass(v string) *DescribeDBInstancesRequest
	GetDBInstanceClass() *string
	SetDBInstanceId(v string) *DescribeDBInstancesRequest
	GetDBInstanceId() *string
	SetDBInstanceStatus(v string) *DescribeDBInstancesRequest
	GetDBInstanceStatus() *string
	SetDBInstanceType(v string) *DescribeDBInstancesRequest
	GetDBInstanceType() *string
	SetDedicatedHostGroupId(v string) *DescribeDBInstancesRequest
	GetDedicatedHostGroupId() *string
	SetDedicatedHostId(v string) *DescribeDBInstancesRequest
	GetDedicatedHostId() *string
	SetEngine(v string) *DescribeDBInstancesRequest
	GetEngine() *string
	SetEngineVersion(v string) *DescribeDBInstancesRequest
	GetEngineVersion() *string
	SetExpired(v string) *DescribeDBInstancesRequest
	GetExpired() *string
	SetFilter(v string) *DescribeDBInstancesRequest
	GetFilter() *string
	SetInstanceLevel(v int32) *DescribeDBInstancesRequest
	GetInstanceLevel() *int32
	SetInstanceNetworkType(v string) *DescribeDBInstancesRequest
	GetInstanceNetworkType() *string
	SetMaxResults(v int32) *DescribeDBInstancesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeDBInstancesRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *DescribeDBInstancesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDBInstancesRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeDBInstancesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDBInstancesRequest
	GetPageSize() *int32
	SetPayType(v string) *DescribeDBInstancesRequest
	GetPayType() *string
	SetQueryAutoRenewal(v bool) *DescribeDBInstancesRequest
	GetQueryAutoRenewal() *bool
	SetRegionId(v string) *DescribeDBInstancesRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeDBInstancesRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeDBInstancesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBInstancesRequest
	GetResourceOwnerId() *int64
	SetSearchKey(v string) *DescribeDBInstancesRequest
	GetSearchKey() *string
	SetTags(v string) *DescribeDBInstancesRequest
	GetTags() *string
	SetVSwitchId(v string) *DescribeDBInstancesRequest
	GetVSwitchId() *string
	SetVpcId(v string) *DescribeDBInstancesRequest
	GetVpcId() *string
	SetZoneId(v string) *DescribeDBInstancesRequest
	GetZoneId() *string
	SetProxyId(v string) *DescribeDBInstancesRequest
	GetProxyId() *string
}

type DescribeDBInstancesRequest struct {
	// The RDS edition of the instance. Valid values:
	//
	// 	- **Basic**: RDS Basic Edition
	//
	// 	- **HighAvailability**: RDS High-availability Edition
	//
	// 	- **cluster**: RDS Cluster Edition
	//
	// 	- **serverless_basic**: RDS Serverless Basic Edition
	//
	// example:
	//
	// cluster
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCzxxxxxxx
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
	// The endpoint of the instance. You must specify this parameter only when you want to query a single instance.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxx.mysql.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The instance type of the instance. For information, see [Primary ApsaraDB RDS instance types](https://help.aliyun.com/document_detail/26312.html).
	//
	// example:
	//
	// rds.mys2.small
	DBInstanceClass *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The status of the instance. For more information, see [Instance states](https://help.aliyun.com/document_detail/26315.html).
	//
	// example:
	//
	// Running
	DBInstanceStatus *string `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty"`
	// The role of the instance. Valid values:
	//
	// 	- **Primary**: primary instance
	//
	// 	- **Readonly**: read-only instance
	//
	// 	- **Guard**: disaster recovery instance
	//
	// 	- **Temp**: temporary instance
	//
	// By default, this operation returns the instances that assume any of the supported roles.
	//
	// example:
	//
	// Primary
	DBInstanceType *string `json:"DBInstanceType,omitempty" xml:"DBInstanceType,omitempty"`
	// The dedicated cluster ID.
	//
	// example:
	//
	// dhg-7a9xxxxxxxx
	DedicatedHostGroupId *string `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	// The host ID of the instance in the dedicated cluster.
	//
	// example:
	//
	// i-bpxxxxxxx
	DedicatedHostId *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	// The database engine of the instance. Valid values:
	//
	// 	- **MySQL**
	//
	// 	- **SQLServer**
	//
	// 	- **PostgreSQL**
	//
	// 	- **MariaDB**
	//
	// By default, this operation returns the instances that run any of the supported database engines.
	//
	// example:
	//
	// MySQL
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The database engine version.
	//
	// example:
	//
	// 5.7
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// Specifies whether the instances have expired. Valid values:
	//
	// 	- **True**
	//
	// 	- **False**
	//
	// example:
	//
	// True
	Expired *string `json:"Expired,omitempty" xml:"Expired,omitempty"`
	// The JSON string that consists of filter condition parameters and their values.
	//
	// example:
	//
	// {"babelfishEnabled":"true"}
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// Specifies whether to return the RDS edition of the instance by using the Category parameter. Valid values:
	//
	// 	- **0**: returns the RDS edition of the instance.
	//
	// 	- **1**: does not return the RDS edition of the instance.
	//
	// example:
	//
	// 0
	InstanceLevel *int32 `json:"InstanceLevel,omitempty" xml:"InstanceLevel,omitempty"`
	// The network type of the instance. Valid values:
	//
	// 	- **VPC**
	//
	// 	- **Classic**
	//
	// By default, this operation returns the instances that reside in any of the supported network types.
	//
	// example:
	//
	// Classic
	InstanceNetworkType *string `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	// The number of entries to return per page. Valid values: **1 to 100**.
	//
	// Default value: **30**.
	//
	// > If you specify this parameter, **PageSize*	- and **PageNumber*	- are unavailable.
	//
	// example:
	//
	// 30
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to display the next page. You must set this parameter to the value that is returned from the most recent call of the **DescribeDBInstances*	- operation for **NextToken**. If the returned entries are displayed on multiple pages, the next page can be displayed when you call this operation again with this parameter specified.
	//
	// example:
	//
	// o7PORW5o2TJg**********
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Pages start from 1.
	//
	// Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values: **1*	- to **100**.
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
	// example:
	//
	// Postpaid
	PayType          *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	QueryAutoRenewal *bool   `json:"QueryAutoRenewal,omitempty" xml:"QueryAutoRenewal,omitempty"`
	// The region ID. You can call the DescribeRegions operation to query the most recent region list.
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
	// rg-acfmyxxxxx
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The keyword that is used for fuzzy search. The keyword can be part of an instance ID or an instance description.
	//
	// example:
	//
	// rm-uf6w
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	// The tag that is added to the instance. Each tag is a key-value pair that consists of two fields: TagKey and TagValue. You can specify a maximum of five tags in the following format for each request: {"key1":"value1","key2":"value2"...}.
	//
	// example:
	//
	// {"key1":"value1"}
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-uf6adz52c2pxxxxxxxxxx
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-uf6f7l4fg90xxxxxxxxxx
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zone ID of the instance.
	//
	// example:
	//
	// cn-hangzhou-a
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	// A deprecated parameter. You do not need to configure this parameter.
	//
	// example:
	//
	// API
	ProxyId *string `json:"proxyId,omitempty" xml:"proxyId,omitempty"`
}

func (s DescribeDBInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesRequest) GetCategory() *string {
	return s.Category
}

func (s *DescribeDBInstancesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeDBInstancesRequest) GetConnectionMode() *string {
	return s.ConnectionMode
}

func (s *DescribeDBInstancesRequest) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *DescribeDBInstancesRequest) GetDBInstanceClass() *string {
	return s.DBInstanceClass
}

func (s *DescribeDBInstancesRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstancesRequest) GetDBInstanceStatus() *string {
	return s.DBInstanceStatus
}

func (s *DescribeDBInstancesRequest) GetDBInstanceType() *string {
	return s.DBInstanceType
}

func (s *DescribeDBInstancesRequest) GetDedicatedHostGroupId() *string {
	return s.DedicatedHostGroupId
}

func (s *DescribeDBInstancesRequest) GetDedicatedHostId() *string {
	return s.DedicatedHostId
}

func (s *DescribeDBInstancesRequest) GetEngine() *string {
	return s.Engine
}

func (s *DescribeDBInstancesRequest) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeDBInstancesRequest) GetExpired() *string {
	return s.Expired
}

func (s *DescribeDBInstancesRequest) GetFilter() *string {
	return s.Filter
}

func (s *DescribeDBInstancesRequest) GetInstanceLevel() *int32 {
	return s.InstanceLevel
}

func (s *DescribeDBInstancesRequest) GetInstanceNetworkType() *string {
	return s.InstanceNetworkType
}

func (s *DescribeDBInstancesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeDBInstancesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeDBInstancesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDBInstancesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBInstancesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDBInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDBInstancesRequest) GetPayType() *string {
	return s.PayType
}

func (s *DescribeDBInstancesRequest) GetQueryAutoRenewal() *bool {
	return s.QueryAutoRenewal
}

func (s *DescribeDBInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBInstancesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDBInstancesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBInstancesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBInstancesRequest) GetSearchKey() *string {
	return s.SearchKey
}

func (s *DescribeDBInstancesRequest) GetTags() *string {
	return s.Tags
}

func (s *DescribeDBInstancesRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeDBInstancesRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeDBInstancesRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeDBInstancesRequest) GetProxyId() *string {
	return s.ProxyId
}

func (s *DescribeDBInstancesRequest) SetCategory(v string) *DescribeDBInstancesRequest {
	s.Category = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetClientToken(v string) *DescribeDBInstancesRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetConnectionMode(v string) *DescribeDBInstancesRequest {
	s.ConnectionMode = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetConnectionString(v string) *DescribeDBInstancesRequest {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetDBInstanceClass(v string) *DescribeDBInstancesRequest {
	s.DBInstanceClass = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetDBInstanceId(v string) *DescribeDBInstancesRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetDBInstanceStatus(v string) *DescribeDBInstancesRequest {
	s.DBInstanceStatus = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetDBInstanceType(v string) *DescribeDBInstancesRequest {
	s.DBInstanceType = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetDedicatedHostGroupId(v string) *DescribeDBInstancesRequest {
	s.DedicatedHostGroupId = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetDedicatedHostId(v string) *DescribeDBInstancesRequest {
	s.DedicatedHostId = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetEngine(v string) *DescribeDBInstancesRequest {
	s.Engine = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetEngineVersion(v string) *DescribeDBInstancesRequest {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetExpired(v string) *DescribeDBInstancesRequest {
	s.Expired = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetFilter(v string) *DescribeDBInstancesRequest {
	s.Filter = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetInstanceLevel(v int32) *DescribeDBInstancesRequest {
	s.InstanceLevel = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetInstanceNetworkType(v string) *DescribeDBInstancesRequest {
	s.InstanceNetworkType = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetMaxResults(v int32) *DescribeDBInstancesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetNextToken(v string) *DescribeDBInstancesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetOwnerAccount(v string) *DescribeDBInstancesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetOwnerId(v int64) *DescribeDBInstancesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetPageNumber(v int32) *DescribeDBInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetPageSize(v int32) *DescribeDBInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetPayType(v string) *DescribeDBInstancesRequest {
	s.PayType = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetQueryAutoRenewal(v bool) *DescribeDBInstancesRequest {
	s.QueryAutoRenewal = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetRegionId(v string) *DescribeDBInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetResourceGroupId(v string) *DescribeDBInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetResourceOwnerAccount(v string) *DescribeDBInstancesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetResourceOwnerId(v int64) *DescribeDBInstancesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetSearchKey(v string) *DescribeDBInstancesRequest {
	s.SearchKey = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetTags(v string) *DescribeDBInstancesRequest {
	s.Tags = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetVSwitchId(v string) *DescribeDBInstancesRequest {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetVpcId(v string) *DescribeDBInstancesRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetZoneId(v string) *DescribeDBInstancesRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetProxyId(v string) *DescribeDBInstancesRequest {
	s.ProxyId = &v
	return s
}

func (s *DescribeDBInstancesRequest) Validate() error {
	return dara.Validate(s)
}
