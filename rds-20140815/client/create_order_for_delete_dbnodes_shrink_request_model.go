// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrderForDeleteDBNodesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *CreateOrderForDeleteDBNodesShrinkRequest
	GetAutoPay() *bool
	SetBusinessInfo(v string) *CreateOrderForDeleteDBNodesShrinkRequest
	GetBusinessInfo() *string
	SetClientToken(v string) *CreateOrderForDeleteDBNodesShrinkRequest
	GetClientToken() *string
	SetCommodityCode(v string) *CreateOrderForDeleteDBNodesShrinkRequest
	GetCommodityCode() *string
	SetDBInstanceId(v string) *CreateOrderForDeleteDBNodesShrinkRequest
	GetDBInstanceId() *string
	SetDBNodeIdShrink(v string) *CreateOrderForDeleteDBNodesShrinkRequest
	GetDBNodeIdShrink() *string
	SetEngineVersion(v string) *CreateOrderForDeleteDBNodesShrinkRequest
	GetEngineVersion() *string
	SetNodeType(v string) *CreateOrderForDeleteDBNodesShrinkRequest
	GetNodeType() *string
	SetOwnerId(v int64) *CreateOrderForDeleteDBNodesShrinkRequest
	GetOwnerId() *int64
	SetPromotionCode(v string) *CreateOrderForDeleteDBNodesShrinkRequest
	GetPromotionCode() *string
	SetRegionId(v string) *CreateOrderForDeleteDBNodesShrinkRequest
	GetRegionId() *string
	SetResource(v string) *CreateOrderForDeleteDBNodesShrinkRequest
	GetResource() *string
	SetResourceGroupId(v string) *CreateOrderForDeleteDBNodesShrinkRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateOrderForDeleteDBNodesShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateOrderForDeleteDBNodesShrinkRequest
	GetResourceOwnerId() *int64
	SetZoneId(v string) *CreateOrderForDeleteDBNodesShrinkRequest
	GetZoneId() *string
}

type CreateOrderForDeleteDBNodesShrinkRequest struct {
	// Specifies whether to automatically complete the payment. Valid values:
	//
	// 1.  **true**: You must make sure that your account balance is sufficient.
	//
	// 2.  **false**: An unpaid order is generated.
	//
	// >  Default value: true. If your account balance is insufficient, you can set the AutoPay parameter to false to generate an unpaid order. Then, you can log on to the ApsaraDB RDS console to complete the payment.
	//
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The additional business information about the instance.
	//
	// example:
	//
	// {\\"shopCartItemId\\":\\"25******\\",\\"produceDriver\\":\\"NoOrder\\",\\"aliyun_shopcart_order_source\\":\\"fromShopcart\\",\\"shopCartId\\":\\"10190203suffix20230509******\\"}
	BusinessInfo *string `json:"BusinessInfo,omitempty" xml:"BusinessInfo,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCzxxxxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The commodity code. Valid values:
	//
	// 	- **bards**: The instance is a pay-as-you-go primary instance.
	//
	// 	- **rds**: The instance is a subscription primary instance.
	//
	// 	- **rords**: The instance is a pay-as-you-go read-only instance.
	//
	// 	- **rds_rordspre_public_cn**: The instance is a subscription read-only instance.
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
	// bards
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The instance ID. You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/610396.html) operation to query the ID of the instance.
	//
	// example:
	//
	// rm-8vb******
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// An array that consists of information about the ID of the node.
	DBNodeIdShrink *string `json:"DBNodeId,omitempty" xml:"DBNodeId,omitempty"`
	// The database engine version of the instance. Valid values:
	//
	// Valid values if you set Engine to MySQL: **5.5, 5.6, 5.7, and 8.0**
	//
	// example:
	//
	// 5.7
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The type of the database node. Valid values:
	//
	// 	- **Master**: the primary node
	//
	// 	- **Slave**: the secondary node
	//
	// example:
	//
	// Master
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The coupon code.
	//
	// example:
	//
	// aliwood-1688-mobile-promotion
	PromotionCode *string `json:"PromotionCode,omitempty" xml:"PromotionCode,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/610399.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resources.
	//
	// example:
	//
	// buy
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmy*****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-a
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateOrderForDeleteDBNodesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOrderForDeleteDBNodesShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateOrderForDeleteDBNodesShrinkRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *CreateOrderForDeleteDBNodesShrinkRequest) GetBusinessInfo() *string {
	return s.BusinessInfo
}

func (s *CreateOrderForDeleteDBNodesShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateOrderForDeleteDBNodesShrinkRequest) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *CreateOrderForDeleteDBNodesShrinkRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateOrderForDeleteDBNodesShrinkRequest) GetDBNodeIdShrink() *string {
	return s.DBNodeIdShrink
}

func (s *CreateOrderForDeleteDBNodesShrinkRequest) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *CreateOrderForDeleteDBNodesShrinkRequest) GetNodeType() *string {
	return s.NodeType
}

func (s *CreateOrderForDeleteDBNodesShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateOrderForDeleteDBNodesShrinkRequest) GetPromotionCode() *string {
	return s.PromotionCode
}

func (s *CreateOrderForDeleteDBNodesShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateOrderForDeleteDBNodesShrinkRequest) GetResource() *string {
	return s.Resource
}

func (s *CreateOrderForDeleteDBNodesShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateOrderForDeleteDBNodesShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateOrderForDeleteDBNodesShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateOrderForDeleteDBNodesShrinkRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateOrderForDeleteDBNodesShrinkRequest) SetAutoPay(v bool) *CreateOrderForDeleteDBNodesShrinkRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateOrderForDeleteDBNodesShrinkRequest) SetBusinessInfo(v string) *CreateOrderForDeleteDBNodesShrinkRequest {
	s.BusinessInfo = &v
	return s
}

func (s *CreateOrderForDeleteDBNodesShrinkRequest) SetClientToken(v string) *CreateOrderForDeleteDBNodesShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateOrderForDeleteDBNodesShrinkRequest) SetCommodityCode(v string) *CreateOrderForDeleteDBNodesShrinkRequest {
	s.CommodityCode = &v
	return s
}

func (s *CreateOrderForDeleteDBNodesShrinkRequest) SetDBInstanceId(v string) *CreateOrderForDeleteDBNodesShrinkRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateOrderForDeleteDBNodesShrinkRequest) SetDBNodeIdShrink(v string) *CreateOrderForDeleteDBNodesShrinkRequest {
	s.DBNodeIdShrink = &v
	return s
}

func (s *CreateOrderForDeleteDBNodesShrinkRequest) SetEngineVersion(v string) *CreateOrderForDeleteDBNodesShrinkRequest {
	s.EngineVersion = &v
	return s
}

func (s *CreateOrderForDeleteDBNodesShrinkRequest) SetNodeType(v string) *CreateOrderForDeleteDBNodesShrinkRequest {
	s.NodeType = &v
	return s
}

func (s *CreateOrderForDeleteDBNodesShrinkRequest) SetOwnerId(v int64) *CreateOrderForDeleteDBNodesShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateOrderForDeleteDBNodesShrinkRequest) SetPromotionCode(v string) *CreateOrderForDeleteDBNodesShrinkRequest {
	s.PromotionCode = &v
	return s
}

func (s *CreateOrderForDeleteDBNodesShrinkRequest) SetRegionId(v string) *CreateOrderForDeleteDBNodesShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateOrderForDeleteDBNodesShrinkRequest) SetResource(v string) *CreateOrderForDeleteDBNodesShrinkRequest {
	s.Resource = &v
	return s
}

func (s *CreateOrderForDeleteDBNodesShrinkRequest) SetResourceGroupId(v string) *CreateOrderForDeleteDBNodesShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateOrderForDeleteDBNodesShrinkRequest) SetResourceOwnerAccount(v string) *CreateOrderForDeleteDBNodesShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateOrderForDeleteDBNodesShrinkRequest) SetResourceOwnerId(v int64) *CreateOrderForDeleteDBNodesShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateOrderForDeleteDBNodesShrinkRequest) SetZoneId(v string) *CreateOrderForDeleteDBNodesShrinkRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateOrderForDeleteDBNodesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
