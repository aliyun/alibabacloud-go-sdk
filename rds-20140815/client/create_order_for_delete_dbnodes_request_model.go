// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrderForDeleteDBNodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *CreateOrderForDeleteDBNodesRequest
	GetAutoPay() *bool
	SetBusinessInfo(v string) *CreateOrderForDeleteDBNodesRequest
	GetBusinessInfo() *string
	SetClientToken(v string) *CreateOrderForDeleteDBNodesRequest
	GetClientToken() *string
	SetCommodityCode(v string) *CreateOrderForDeleteDBNodesRequest
	GetCommodityCode() *string
	SetDBInstanceId(v string) *CreateOrderForDeleteDBNodesRequest
	GetDBInstanceId() *string
	SetDBNodeId(v []*string) *CreateOrderForDeleteDBNodesRequest
	GetDBNodeId() []*string
	SetEngineVersion(v string) *CreateOrderForDeleteDBNodesRequest
	GetEngineVersion() *string
	SetNodeType(v string) *CreateOrderForDeleteDBNodesRequest
	GetNodeType() *string
	SetOwnerId(v int64) *CreateOrderForDeleteDBNodesRequest
	GetOwnerId() *int64
	SetPromotionCode(v string) *CreateOrderForDeleteDBNodesRequest
	GetPromotionCode() *string
	SetRegionId(v string) *CreateOrderForDeleteDBNodesRequest
	GetRegionId() *string
	SetResource(v string) *CreateOrderForDeleteDBNodesRequest
	GetResource() *string
	SetResourceGroupId(v string) *CreateOrderForDeleteDBNodesRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateOrderForDeleteDBNodesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateOrderForDeleteDBNodesRequest
	GetResourceOwnerId() *int64
	SetZoneId(v string) *CreateOrderForDeleteDBNodesRequest
	GetZoneId() *string
}

type CreateOrderForDeleteDBNodesRequest struct {
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
	DBNodeId []*string `json:"DBNodeId,omitempty" xml:"DBNodeId,omitempty" type:"Repeated"`
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

func (s CreateOrderForDeleteDBNodesRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOrderForDeleteDBNodesRequest) GoString() string {
	return s.String()
}

func (s *CreateOrderForDeleteDBNodesRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *CreateOrderForDeleteDBNodesRequest) GetBusinessInfo() *string {
	return s.BusinessInfo
}

func (s *CreateOrderForDeleteDBNodesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateOrderForDeleteDBNodesRequest) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *CreateOrderForDeleteDBNodesRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateOrderForDeleteDBNodesRequest) GetDBNodeId() []*string {
	return s.DBNodeId
}

func (s *CreateOrderForDeleteDBNodesRequest) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *CreateOrderForDeleteDBNodesRequest) GetNodeType() *string {
	return s.NodeType
}

func (s *CreateOrderForDeleteDBNodesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateOrderForDeleteDBNodesRequest) GetPromotionCode() *string {
	return s.PromotionCode
}

func (s *CreateOrderForDeleteDBNodesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateOrderForDeleteDBNodesRequest) GetResource() *string {
	return s.Resource
}

func (s *CreateOrderForDeleteDBNodesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateOrderForDeleteDBNodesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateOrderForDeleteDBNodesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateOrderForDeleteDBNodesRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateOrderForDeleteDBNodesRequest) SetAutoPay(v bool) *CreateOrderForDeleteDBNodesRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateOrderForDeleteDBNodesRequest) SetBusinessInfo(v string) *CreateOrderForDeleteDBNodesRequest {
	s.BusinessInfo = &v
	return s
}

func (s *CreateOrderForDeleteDBNodesRequest) SetClientToken(v string) *CreateOrderForDeleteDBNodesRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateOrderForDeleteDBNodesRequest) SetCommodityCode(v string) *CreateOrderForDeleteDBNodesRequest {
	s.CommodityCode = &v
	return s
}

func (s *CreateOrderForDeleteDBNodesRequest) SetDBInstanceId(v string) *CreateOrderForDeleteDBNodesRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateOrderForDeleteDBNodesRequest) SetDBNodeId(v []*string) *CreateOrderForDeleteDBNodesRequest {
	s.DBNodeId = v
	return s
}

func (s *CreateOrderForDeleteDBNodesRequest) SetEngineVersion(v string) *CreateOrderForDeleteDBNodesRequest {
	s.EngineVersion = &v
	return s
}

func (s *CreateOrderForDeleteDBNodesRequest) SetNodeType(v string) *CreateOrderForDeleteDBNodesRequest {
	s.NodeType = &v
	return s
}

func (s *CreateOrderForDeleteDBNodesRequest) SetOwnerId(v int64) *CreateOrderForDeleteDBNodesRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateOrderForDeleteDBNodesRequest) SetPromotionCode(v string) *CreateOrderForDeleteDBNodesRequest {
	s.PromotionCode = &v
	return s
}

func (s *CreateOrderForDeleteDBNodesRequest) SetRegionId(v string) *CreateOrderForDeleteDBNodesRequest {
	s.RegionId = &v
	return s
}

func (s *CreateOrderForDeleteDBNodesRequest) SetResource(v string) *CreateOrderForDeleteDBNodesRequest {
	s.Resource = &v
	return s
}

func (s *CreateOrderForDeleteDBNodesRequest) SetResourceGroupId(v string) *CreateOrderForDeleteDBNodesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateOrderForDeleteDBNodesRequest) SetResourceOwnerAccount(v string) *CreateOrderForDeleteDBNodesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateOrderForDeleteDBNodesRequest) SetResourceOwnerId(v int64) *CreateOrderForDeleteDBNodesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateOrderForDeleteDBNodesRequest) SetZoneId(v string) *CreateOrderForDeleteDBNodesRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateOrderForDeleteDBNodesRequest) Validate() error {
	return dara.Validate(s)
}
