// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterAvailableResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBNodeClass(v string) *DescribeDBClusterAvailableResourcesRequest
	GetDBNodeClass() *string
	SetDBType(v string) *DescribeDBClusterAvailableResourcesRequest
	GetDBType() *string
	SetDBVersion(v string) *DescribeDBClusterAvailableResourcesRequest
	GetDBVersion() *string
	SetOwnerAccount(v string) *DescribeDBClusterAvailableResourcesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDBClusterAvailableResourcesRequest
	GetOwnerId() *int64
	SetPayType(v string) *DescribeDBClusterAvailableResourcesRequest
	GetPayType() *string
	SetRegionId(v string) *DescribeDBClusterAvailableResourcesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeDBClusterAvailableResourcesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBClusterAvailableResourcesRequest
	GetResourceOwnerId() *int64
	SetZoneId(v string) *DescribeDBClusterAvailableResourcesRequest
	GetZoneId() *string
}

type DescribeDBClusterAvailableResourcesRequest struct {
	// The specifications of the node. For more information, see [Specifications of compute nodes](https://help.aliyun.com/document_detail/102542.html).
	//
	// example:
	//
	// polar.mysql.x4.large
	DBNodeClass *string `json:"DBNodeClass,omitempty" xml:"DBNodeClass,omitempty"`
	// The type of the database engine. Valid values:
	//
	// 	- **MySQL**
	//
	// 	- **PostgreSQL**
	//
	// 	- **Oracle**
	//
	// example:
	//
	// MySQL
	DBType *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	// The version of the database engine. Valid values for the MySQL database engine:
	//
	// 	- **5.6**
	//
	// 	- **5.7**
	//
	// 	- **8.0**
	//
	// Valid values for the PostgreSQL database engine:
	//
	// 	- **11**
	//
	// 	- **14**
	//
	// Valid value for the Oracle database engine: **11**
	//
	// > This parameter is required when you specify the **DBType*	- parameter.
	//
	// example:
	//
	// 5.6
	DBVersion    *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The billing method of the cluster. Valid values:
	//
	// 	- **Postpaid**: pay-as-you-go
	//
	// 	- **Prepaid**: subscription
	//
	// This parameter is required.
	//
	// example:
	//
	// Postpaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The region ID of the cluster. Default value: **cn-hangzhou**.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/98041.html) operation to query the available regions.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The zone ID of the cluster.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/98041.html) operation to query the available zones.
	//
	// example:
	//
	// cn-hangzhou-i
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDBClusterAvailableResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterAvailableResourcesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterAvailableResourcesRequest) GetDBNodeClass() *string {
	return s.DBNodeClass
}

func (s *DescribeDBClusterAvailableResourcesRequest) GetDBType() *string {
	return s.DBType
}

func (s *DescribeDBClusterAvailableResourcesRequest) GetDBVersion() *string {
	return s.DBVersion
}

func (s *DescribeDBClusterAvailableResourcesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDBClusterAvailableResourcesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBClusterAvailableResourcesRequest) GetPayType() *string {
	return s.PayType
}

func (s *DescribeDBClusterAvailableResourcesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBClusterAvailableResourcesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBClusterAvailableResourcesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBClusterAvailableResourcesRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeDBClusterAvailableResourcesRequest) SetDBNodeClass(v string) *DescribeDBClusterAvailableResourcesRequest {
	s.DBNodeClass = &v
	return s
}

func (s *DescribeDBClusterAvailableResourcesRequest) SetDBType(v string) *DescribeDBClusterAvailableResourcesRequest {
	s.DBType = &v
	return s
}

func (s *DescribeDBClusterAvailableResourcesRequest) SetDBVersion(v string) *DescribeDBClusterAvailableResourcesRequest {
	s.DBVersion = &v
	return s
}

func (s *DescribeDBClusterAvailableResourcesRequest) SetOwnerAccount(v string) *DescribeDBClusterAvailableResourcesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDBClusterAvailableResourcesRequest) SetOwnerId(v int64) *DescribeDBClusterAvailableResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBClusterAvailableResourcesRequest) SetPayType(v string) *DescribeDBClusterAvailableResourcesRequest {
	s.PayType = &v
	return s
}

func (s *DescribeDBClusterAvailableResourcesRequest) SetRegionId(v string) *DescribeDBClusterAvailableResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBClusterAvailableResourcesRequest) SetResourceOwnerAccount(v string) *DescribeDBClusterAvailableResourcesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBClusterAvailableResourcesRequest) SetResourceOwnerId(v int64) *DescribeDBClusterAvailableResourcesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBClusterAvailableResourcesRequest) SetZoneId(v string) *DescribeDBClusterAvailableResourcesRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeDBClusterAvailableResourcesRequest) Validate() error {
	return dara.Validate(s)
}
