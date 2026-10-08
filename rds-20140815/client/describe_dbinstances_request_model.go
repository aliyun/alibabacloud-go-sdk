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
	// The instance edition. Valid values:
	//
	// - **Basic**: Basic Edition
	//
	// - **HighAvailability**: High-availability Edition
	//
	// - **cluster**: Cluster Edition
	//
	// - **serverless_basic**: Serverless
	//
	// example:
	//
	// cluster
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The access mode of the instance. Valid values:
	//
	// 	- **Standard**: standard access mode
	//
	// 	- **Safe**: database proxy mode
	//
	// By default, instances in all access modes are returned.
	//
	// example:
	//
	// Standard
	ConnectionMode *string `json:"ConnectionMode,omitempty" xml:"ConnectionMode,omitempty"`
	// The endpoint of the instance. Use this endpoint to query the corresponding instance.
	//
	// example:
	//
	// rm-uf6wjk5****.mysql.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The instance type. For more information, see [Instance types](https://help.aliyun.com/document_detail/26312.html).
	//
	// example:
	//
	// rds.mys2.small
	DBInstanceClass *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rm-uf6wjk5****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The instance status. For more information, see [Instance states](https://help.aliyun.com/document_detail/26315.html).
	//
	// example:
	//
	// Running
	DBInstanceStatus *string `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty"`
	// The instance type. Valid values:
	//
	// 	- **Primary**: primary instance
	//
	// 	- **Readonly**: read-only instance
	//
	// 	- **Guard**: disaster recovery instance
	//
	// 	- **Temp**: temporary instance
	//
	// By default, instances of all types are returned.
	//
	// example:
	//
	// Primary
	DBInstanceType *string `json:"DBInstanceType,omitempty" xml:"DBInstanceType,omitempty"`
	// The dedicated cluster ID.
	//
	// example:
	//
	// dhg-7a9****
	DedicatedHostGroupId *string `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	// The host ID in the dedicated cluster.
	//
	// example:
	//
	// i-bp****
	DedicatedHostId *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	// The database engine. Valid values:
	//
	// 	- **MySQL**
	//
	// 	- **SQLServer**
	//
	// 	- **PostgreSQL**
	//
	// 	- **MariaDB**
	//
	// By default, instances of all database engines are returned.
	//
	// example:
	//
	// MySQL
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The database engine version.
	//
	// example:
	//
	// 8.0
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The expiration status of the instance. Valid values:
	//
	// 	- **True**: The instance has expired.
	//
	// 	- **False**: The instance has not expired.
	//
	// example:
	//
	// True
	Expired *string `json:"Expired,omitempty" xml:"Expired,omitempty"`
	// The JSON string that contains the instance filter conditions and their values.
	//
	// example:
	//
	// {"babelfishEnabled":"true"}
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// Specifies whether to return the instance edition (Category) information. Valid values:
	//
	// 	- **0**: does not return the information
	//
	// 	- **1**: returns the information
	//
	// example:
	//
	// 0
	InstanceLevel *int32 `json:"InstanceLevel,omitempty" xml:"InstanceLevel,omitempty"`
	// The network type of the instance. Valid values:
	//
	// 	- **VPC**: an instance in a virtual private cloud (VPC)
	//
	// 	- **Classic**: an instance in the classic network
	//
	// By default, instances of all network types are returned.
	//
	// example:
	//
	// Classic
	InstanceNetworkType *string `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	// The number of entries per page. Valid values: **1*	- to **100**.
	//
	// Default value: **30**.
	//
	// >If you specify this parameter, the **PageSize*	- and **PageNumber*	- parameters are unavailable.
	//
	// example:
	//
	// 30
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token. Set this parameter to the value of **NextToken*	- that is returned from the last call to the **DescribeDBInstances*	- operation. If the results span multiple pages, pass in this value to retrieve the next page.
	//
	// example:
	//
	// o7PORW5o2TJg****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Valid values: any value greater than 0 that does not exceed the maximum value of Integer.
	//
	// Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: **1*	- to **100**.
	//
	// Default value: **30**.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The billing method. Valid values:
	//
	// 	- **Postpaid**: pay-as-you-go
	//
	// 	- **Prepaid**: subscription
	//
	// example:
	//
	// Postpaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// A reserved parameter. You do not need to configure this parameter.
	//
	// example:
	//
	// test
	QueryAutoRenewal *bool `json:"QueryAutoRenewal,omitempty" xml:"QueryAutoRenewal,omitempty"`
	// The region ID. You can call DescribeRegions to query the available regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmy****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The keyword for fuzzy search based on the instance ID or instance description.
	//
	// example:
	//
	// rm-uf6w
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	// The tags that are bound to the instance, including TagKey and TagValue. You can specify up to five pairs of tags at a time. Format: {"key1":"value1","key2":"value2"...}. If the instance matches any of the specified tags, the instance information is returned.
	//
	// example:
	//
	// {"key1":"value1"}
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-uf6adz52c2p****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// VPC ID。
	//
	// example:
	//
	// vpc-uf6f7l4fg90****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zone ID.
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
