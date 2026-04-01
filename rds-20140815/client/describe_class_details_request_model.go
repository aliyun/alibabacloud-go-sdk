// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClassDetailsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClassCode(v string) *DescribeClassDetailsRequest
	GetClassCode() *string
	SetClientToken(v string) *DescribeClassDetailsRequest
	GetClientToken() *string
	SetCommodityCode(v string) *DescribeClassDetailsRequest
	GetCommodityCode() *string
	SetEngine(v string) *DescribeClassDetailsRequest
	GetEngine() *string
	SetEngineVersion(v string) *DescribeClassDetailsRequest
	GetEngineVersion() *string
	SetOwnerId(v int64) *DescribeClassDetailsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeClassDetailsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeClassDetailsRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeClassDetailsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeClassDetailsRequest
	GetResourceOwnerId() *int64
}

type DescribeClassDetailsRequest struct {
	// The code of the instance type.
	//
	// This parameter is required.
	//
	// example:
	//
	// rds.mysql.s3.large
	ClassCode *string `json:"ClassCode,omitempty" xml:"ClassCode,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the generated token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz*****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The commodity code of the instance. Valid values:
	//
	// 	- **bards_intl**: The instance is a pay-as-you-go primary instance.
	//
	// 	- **rds_intl**: The instance is a subscription primary instance.
	//
	// 	- **rords_intl**: The instance is a pay-as-you-go read-only instance.
	//
	// 	- **rds_rordspre_public_intl**: The instance is a subscription read-only instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// rds
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The type of the database engine.
	//
	// This parameter is required.
	//
	// example:
	//
	// MySQL
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The database engine version of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5.6
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call the DescribeRegions operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID. You can call the DescribeDBInstanceAttribute operation to query the resource group ID.
	//
	// example:
	//
	// rg-acfmy*****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeClassDetailsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeClassDetailsRequest) GoString() string {
	return s.String()
}

func (s *DescribeClassDetailsRequest) GetClassCode() *string {
	return s.ClassCode
}

func (s *DescribeClassDetailsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeClassDetailsRequest) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *DescribeClassDetailsRequest) GetEngine() *string {
	return s.Engine
}

func (s *DescribeClassDetailsRequest) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeClassDetailsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeClassDetailsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeClassDetailsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeClassDetailsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeClassDetailsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeClassDetailsRequest) SetClassCode(v string) *DescribeClassDetailsRequest {
	s.ClassCode = &v
	return s
}

func (s *DescribeClassDetailsRequest) SetClientToken(v string) *DescribeClassDetailsRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeClassDetailsRequest) SetCommodityCode(v string) *DescribeClassDetailsRequest {
	s.CommodityCode = &v
	return s
}

func (s *DescribeClassDetailsRequest) SetEngine(v string) *DescribeClassDetailsRequest {
	s.Engine = &v
	return s
}

func (s *DescribeClassDetailsRequest) SetEngineVersion(v string) *DescribeClassDetailsRequest {
	s.EngineVersion = &v
	return s
}

func (s *DescribeClassDetailsRequest) SetOwnerId(v int64) *DescribeClassDetailsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeClassDetailsRequest) SetRegionId(v string) *DescribeClassDetailsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeClassDetailsRequest) SetResourceGroupId(v string) *DescribeClassDetailsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeClassDetailsRequest) SetResourceOwnerAccount(v string) *DescribeClassDetailsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeClassDetailsRequest) SetResourceOwnerId(v int64) *DescribeClassDetailsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeClassDetailsRequest) Validate() error {
	return dara.Validate(s)
}
