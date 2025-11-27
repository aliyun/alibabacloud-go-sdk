// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreCheckCreateOrderForDeleteDBNodesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *PreCheckCreateOrderForDeleteDBNodesShrinkRequest
	GetAutoPay() *bool
	SetBusinessInfo(v string) *PreCheckCreateOrderForDeleteDBNodesShrinkRequest
	GetBusinessInfo() *string
	SetClientToken(v string) *PreCheckCreateOrderForDeleteDBNodesShrinkRequest
	GetClientToken() *string
	SetCommodityCode(v string) *PreCheckCreateOrderForDeleteDBNodesShrinkRequest
	GetCommodityCode() *string
	SetDBInstanceId(v string) *PreCheckCreateOrderForDeleteDBNodesShrinkRequest
	GetDBInstanceId() *string
	SetDBNodeIdShrink(v string) *PreCheckCreateOrderForDeleteDBNodesShrinkRequest
	GetDBNodeIdShrink() *string
	SetEngineVersion(v string) *PreCheckCreateOrderForDeleteDBNodesShrinkRequest
	GetEngineVersion() *string
	SetNodeType(v string) *PreCheckCreateOrderForDeleteDBNodesShrinkRequest
	GetNodeType() *string
	SetOwnerId(v int64) *PreCheckCreateOrderForDeleteDBNodesShrinkRequest
	GetOwnerId() *int64
	SetPromotionCode(v string) *PreCheckCreateOrderForDeleteDBNodesShrinkRequest
	GetPromotionCode() *string
	SetRegionId(v string) *PreCheckCreateOrderForDeleteDBNodesShrinkRequest
	GetRegionId() *string
	SetResource(v string) *PreCheckCreateOrderForDeleteDBNodesShrinkRequest
	GetResource() *string
	SetResourceOwnerAccount(v string) *PreCheckCreateOrderForDeleteDBNodesShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *PreCheckCreateOrderForDeleteDBNodesShrinkRequest
	GetResourceOwnerId() *int64
	SetZoneId(v string) *PreCheckCreateOrderForDeleteDBNodesShrinkRequest
	GetZoneId() *string
}

type PreCheckCreateOrderForDeleteDBNodesShrinkRequest struct {
	// Specifies whether to automatically complete the payment. Valid values:
	//
	// 1.  **true**: automatically completes the payment. You must make sure that your account balance is sufficient.
	//
	// 2.  **false**: does not automatically complete the payment. An unpaid order is generated.
	//
	// >  The default value is true. If your account balance is insufficient, you can set the AutoPay parameter to false to generate an unpaid order. Then, you can log on to the ApsaraDB RDS console to complete the payment.
	//
	// example:
	//
	// True
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The additional business information about the instance.
	//
	// example:
	//
	// {\\"promotion_input_param\\":\\"{\\\\\\"promotionFilter\\\\\\":{},\\\\\\"promotionOptionCode\\\\\\":\\\\\\"youhui_quan\\\\\\"}\\"}
	BusinessInfo *string `json:"BusinessInfo,omitempty" xml:"BusinessInfo,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the generated token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOC****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The commodity code. Valid value:
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
	// rds
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rm-wz9rziy3he051if82
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The node IDs.
	DBNodeIdShrink *string `json:"DBNodeId,omitempty" xml:"DBNodeId,omitempty"`
	// The major engine version of the destination instance. The value of this parameter varies based on the value of **Engine**.
	//
	// 	- Valid values when Engine is set to MySQL: **5.5, 5.6, 5.7, and 8.0**
	//
	// 	- Valid values when Engine is set to SQLServer: **2008r2, 08r2_ent_ha, 2012, 2012_ent_ha, 2012_std_ha, 2012_web, 2014_std_ha, 2016_ent_ha, 2016_std_ha, 2016_web, 2017_std_ha, 2017_ent, 2019_std_ha, and 2019_ent**
	//
	// 	- Valid values when Engine is set to PostgreSQL: **10.0, 11.0, 12.0, 13.0, 14.0, and 15.0**
	//
	// example:
	//
	// 8.0
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The type of the database node. Valid value:
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
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/26243.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource of the instance.
	//
	// example:
	//
	// buy
	Resource             *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-a
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s PreCheckCreateOrderForDeleteDBNodesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s PreCheckCreateOrderForDeleteDBNodesShrinkRequest) GoString() string {
	return s.String()
}

func (s *PreCheckCreateOrderForDeleteDBNodesShrinkRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *PreCheckCreateOrderForDeleteDBNodesShrinkRequest) GetBusinessInfo() *string {
	return s.BusinessInfo
}

func (s *PreCheckCreateOrderForDeleteDBNodesShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *PreCheckCreateOrderForDeleteDBNodesShrinkRequest) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *PreCheckCreateOrderForDeleteDBNodesShrinkRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *PreCheckCreateOrderForDeleteDBNodesShrinkRequest) GetDBNodeIdShrink() *string {
	return s.DBNodeIdShrink
}

func (s *PreCheckCreateOrderForDeleteDBNodesShrinkRequest) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *PreCheckCreateOrderForDeleteDBNodesShrinkRequest) GetNodeType() *string {
	return s.NodeType
}

func (s *PreCheckCreateOrderForDeleteDBNodesShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *PreCheckCreateOrderForDeleteDBNodesShrinkRequest) GetPromotionCode() *string {
	return s.PromotionCode
}

func (s *PreCheckCreateOrderForDeleteDBNodesShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *PreCheckCreateOrderForDeleteDBNodesShrinkRequest) GetResource() *string {
	return s.Resource
}

func (s *PreCheckCreateOrderForDeleteDBNodesShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *PreCheckCreateOrderForDeleteDBNodesShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *PreCheckCreateOrderForDeleteDBNodesShrinkRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *PreCheckCreateOrderForDeleteDBNodesShrinkRequest) SetAutoPay(v bool) *PreCheckCreateOrderForDeleteDBNodesShrinkRequest {
	s.AutoPay = &v
	return s
}

func (s *PreCheckCreateOrderForDeleteDBNodesShrinkRequest) SetBusinessInfo(v string) *PreCheckCreateOrderForDeleteDBNodesShrinkRequest {
	s.BusinessInfo = &v
	return s
}

func (s *PreCheckCreateOrderForDeleteDBNodesShrinkRequest) SetClientToken(v string) *PreCheckCreateOrderForDeleteDBNodesShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *PreCheckCreateOrderForDeleteDBNodesShrinkRequest) SetCommodityCode(v string) *PreCheckCreateOrderForDeleteDBNodesShrinkRequest {
	s.CommodityCode = &v
	return s
}

func (s *PreCheckCreateOrderForDeleteDBNodesShrinkRequest) SetDBInstanceId(v string) *PreCheckCreateOrderForDeleteDBNodesShrinkRequest {
	s.DBInstanceId = &v
	return s
}

func (s *PreCheckCreateOrderForDeleteDBNodesShrinkRequest) SetDBNodeIdShrink(v string) *PreCheckCreateOrderForDeleteDBNodesShrinkRequest {
	s.DBNodeIdShrink = &v
	return s
}

func (s *PreCheckCreateOrderForDeleteDBNodesShrinkRequest) SetEngineVersion(v string) *PreCheckCreateOrderForDeleteDBNodesShrinkRequest {
	s.EngineVersion = &v
	return s
}

func (s *PreCheckCreateOrderForDeleteDBNodesShrinkRequest) SetNodeType(v string) *PreCheckCreateOrderForDeleteDBNodesShrinkRequest {
	s.NodeType = &v
	return s
}

func (s *PreCheckCreateOrderForDeleteDBNodesShrinkRequest) SetOwnerId(v int64) *PreCheckCreateOrderForDeleteDBNodesShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *PreCheckCreateOrderForDeleteDBNodesShrinkRequest) SetPromotionCode(v string) *PreCheckCreateOrderForDeleteDBNodesShrinkRequest {
	s.PromotionCode = &v
	return s
}

func (s *PreCheckCreateOrderForDeleteDBNodesShrinkRequest) SetRegionId(v string) *PreCheckCreateOrderForDeleteDBNodesShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *PreCheckCreateOrderForDeleteDBNodesShrinkRequest) SetResource(v string) *PreCheckCreateOrderForDeleteDBNodesShrinkRequest {
	s.Resource = &v
	return s
}

func (s *PreCheckCreateOrderForDeleteDBNodesShrinkRequest) SetResourceOwnerAccount(v string) *PreCheckCreateOrderForDeleteDBNodesShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *PreCheckCreateOrderForDeleteDBNodesShrinkRequest) SetResourceOwnerId(v int64) *PreCheckCreateOrderForDeleteDBNodesShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *PreCheckCreateOrderForDeleteDBNodesShrinkRequest) SetZoneId(v string) *PreCheckCreateOrderForDeleteDBNodesShrinkRequest {
	s.ZoneId = &v
	return s
}

func (s *PreCheckCreateOrderForDeleteDBNodesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
