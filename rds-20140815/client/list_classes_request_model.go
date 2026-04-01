// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClassesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ListClassesRequest
	GetClientToken() *string
	SetCommodityCode(v string) *ListClassesRequest
	GetCommodityCode() *string
	SetDBInstanceId(v string) *ListClassesRequest
	GetDBInstanceId() *string
	SetEngine(v string) *ListClassesRequest
	GetEngine() *string
	SetOrderType(v string) *ListClassesRequest
	GetOrderType() *string
	SetOwnerId(v int64) *ListClassesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ListClassesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ListClassesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListClassesRequest
	GetResourceOwnerId() *int64
}

type ListClassesRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCzxxxxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The commodity code of the instances.
	//
	// 	- **bards_intl**: The instances are pay-as-you-go primary instances.
	//
	// 	- **rds_intl**: The instances are subscription primary instances.
	//
	// 	- **rords_intl**: The instances are pay-as-you-go read-only instances.
	//
	// 	- **rds_rordspre_public_intl**: The instances are subscription read-only instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// bards_intl
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// >  If you set the **CommodityCode*	- parameter to the commodity code of read-only instances, you must specify this parameter.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
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
	// example:
	//
	// MySQL
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The type of order that you want to query. Valid values:
	//
	// 	- **BUY**: specifies the query orders that are used to purchase instances.
	//
	// 	- **UPGRADE**: specifies the query orders that are used to change the specifications of instances.
	//
	// 	- **RENEW**: specifies the query orders that are used to renew instances.
	//
	// 	- **CONVERT**: specifies the query orders that are used to change the billing methods of instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// BUY
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call the DescribeRegions operation to query the most recent region list.
	//
	// >  If you are using an Alibaba Cloud account on the International site (alibabacloud.com), you must specify this parameter.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListClassesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListClassesRequest) GoString() string {
	return s.String()
}

func (s *ListClassesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ListClassesRequest) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *ListClassesRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ListClassesRequest) GetEngine() *string {
	return s.Engine
}

func (s *ListClassesRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *ListClassesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListClassesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListClassesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListClassesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListClassesRequest) SetClientToken(v string) *ListClassesRequest {
	s.ClientToken = &v
	return s
}

func (s *ListClassesRequest) SetCommodityCode(v string) *ListClassesRequest {
	s.CommodityCode = &v
	return s
}

func (s *ListClassesRequest) SetDBInstanceId(v string) *ListClassesRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ListClassesRequest) SetEngine(v string) *ListClassesRequest {
	s.Engine = &v
	return s
}

func (s *ListClassesRequest) SetOrderType(v string) *ListClassesRequest {
	s.OrderType = &v
	return s
}

func (s *ListClassesRequest) SetOwnerId(v int64) *ListClassesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListClassesRequest) SetRegionId(v string) *ListClassesRequest {
	s.RegionId = &v
	return s
}

func (s *ListClassesRequest) SetResourceOwnerAccount(v string) *ListClassesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListClassesRequest) SetResourceOwnerId(v int64) *ListClassesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListClassesRequest) Validate() error {
	return dara.Validate(s)
}
