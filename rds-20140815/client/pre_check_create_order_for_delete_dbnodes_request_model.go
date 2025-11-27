// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreCheckCreateOrderForDeleteDBNodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *PreCheckCreateOrderForDeleteDBNodesRequest
	GetAutoPay() *bool
	SetBusinessInfo(v string) *PreCheckCreateOrderForDeleteDBNodesRequest
	GetBusinessInfo() *string
	SetClientToken(v string) *PreCheckCreateOrderForDeleteDBNodesRequest
	GetClientToken() *string
	SetCommodityCode(v string) *PreCheckCreateOrderForDeleteDBNodesRequest
	GetCommodityCode() *string
	SetDBInstanceId(v string) *PreCheckCreateOrderForDeleteDBNodesRequest
	GetDBInstanceId() *string
	SetDBNodeId(v []*string) *PreCheckCreateOrderForDeleteDBNodesRequest
	GetDBNodeId() []*string
	SetEngineVersion(v string) *PreCheckCreateOrderForDeleteDBNodesRequest
	GetEngineVersion() *string
	SetNodeType(v string) *PreCheckCreateOrderForDeleteDBNodesRequest
	GetNodeType() *string
	SetOwnerId(v int64) *PreCheckCreateOrderForDeleteDBNodesRequest
	GetOwnerId() *int64
	SetPromotionCode(v string) *PreCheckCreateOrderForDeleteDBNodesRequest
	GetPromotionCode() *string
	SetRegionId(v string) *PreCheckCreateOrderForDeleteDBNodesRequest
	GetRegionId() *string
	SetResource(v string) *PreCheckCreateOrderForDeleteDBNodesRequest
	GetResource() *string
	SetResourceOwnerAccount(v string) *PreCheckCreateOrderForDeleteDBNodesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *PreCheckCreateOrderForDeleteDBNodesRequest
	GetResourceOwnerId() *int64
	SetZoneId(v string) *PreCheckCreateOrderForDeleteDBNodesRequest
	GetZoneId() *string
}

type PreCheckCreateOrderForDeleteDBNodesRequest struct {
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
	DBNodeId []*string `json:"DBNodeId,omitempty" xml:"DBNodeId,omitempty" type:"Repeated"`
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

func (s PreCheckCreateOrderForDeleteDBNodesRequest) String() string {
	return dara.Prettify(s)
}

func (s PreCheckCreateOrderForDeleteDBNodesRequest) GoString() string {
	return s.String()
}

func (s *PreCheckCreateOrderForDeleteDBNodesRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *PreCheckCreateOrderForDeleteDBNodesRequest) GetBusinessInfo() *string {
	return s.BusinessInfo
}

func (s *PreCheckCreateOrderForDeleteDBNodesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *PreCheckCreateOrderForDeleteDBNodesRequest) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *PreCheckCreateOrderForDeleteDBNodesRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *PreCheckCreateOrderForDeleteDBNodesRequest) GetDBNodeId() []*string {
	return s.DBNodeId
}

func (s *PreCheckCreateOrderForDeleteDBNodesRequest) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *PreCheckCreateOrderForDeleteDBNodesRequest) GetNodeType() *string {
	return s.NodeType
}

func (s *PreCheckCreateOrderForDeleteDBNodesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *PreCheckCreateOrderForDeleteDBNodesRequest) GetPromotionCode() *string {
	return s.PromotionCode
}

func (s *PreCheckCreateOrderForDeleteDBNodesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *PreCheckCreateOrderForDeleteDBNodesRequest) GetResource() *string {
	return s.Resource
}

func (s *PreCheckCreateOrderForDeleteDBNodesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *PreCheckCreateOrderForDeleteDBNodesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *PreCheckCreateOrderForDeleteDBNodesRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *PreCheckCreateOrderForDeleteDBNodesRequest) SetAutoPay(v bool) *PreCheckCreateOrderForDeleteDBNodesRequest {
	s.AutoPay = &v
	return s
}

func (s *PreCheckCreateOrderForDeleteDBNodesRequest) SetBusinessInfo(v string) *PreCheckCreateOrderForDeleteDBNodesRequest {
	s.BusinessInfo = &v
	return s
}

func (s *PreCheckCreateOrderForDeleteDBNodesRequest) SetClientToken(v string) *PreCheckCreateOrderForDeleteDBNodesRequest {
	s.ClientToken = &v
	return s
}

func (s *PreCheckCreateOrderForDeleteDBNodesRequest) SetCommodityCode(v string) *PreCheckCreateOrderForDeleteDBNodesRequest {
	s.CommodityCode = &v
	return s
}

func (s *PreCheckCreateOrderForDeleteDBNodesRequest) SetDBInstanceId(v string) *PreCheckCreateOrderForDeleteDBNodesRequest {
	s.DBInstanceId = &v
	return s
}

func (s *PreCheckCreateOrderForDeleteDBNodesRequest) SetDBNodeId(v []*string) *PreCheckCreateOrderForDeleteDBNodesRequest {
	s.DBNodeId = v
	return s
}

func (s *PreCheckCreateOrderForDeleteDBNodesRequest) SetEngineVersion(v string) *PreCheckCreateOrderForDeleteDBNodesRequest {
	s.EngineVersion = &v
	return s
}

func (s *PreCheckCreateOrderForDeleteDBNodesRequest) SetNodeType(v string) *PreCheckCreateOrderForDeleteDBNodesRequest {
	s.NodeType = &v
	return s
}

func (s *PreCheckCreateOrderForDeleteDBNodesRequest) SetOwnerId(v int64) *PreCheckCreateOrderForDeleteDBNodesRequest {
	s.OwnerId = &v
	return s
}

func (s *PreCheckCreateOrderForDeleteDBNodesRequest) SetPromotionCode(v string) *PreCheckCreateOrderForDeleteDBNodesRequest {
	s.PromotionCode = &v
	return s
}

func (s *PreCheckCreateOrderForDeleteDBNodesRequest) SetRegionId(v string) *PreCheckCreateOrderForDeleteDBNodesRequest {
	s.RegionId = &v
	return s
}

func (s *PreCheckCreateOrderForDeleteDBNodesRequest) SetResource(v string) *PreCheckCreateOrderForDeleteDBNodesRequest {
	s.Resource = &v
	return s
}

func (s *PreCheckCreateOrderForDeleteDBNodesRequest) SetResourceOwnerAccount(v string) *PreCheckCreateOrderForDeleteDBNodesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *PreCheckCreateOrderForDeleteDBNodesRequest) SetResourceOwnerId(v int64) *PreCheckCreateOrderForDeleteDBNodesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *PreCheckCreateOrderForDeleteDBNodesRequest) SetZoneId(v string) *PreCheckCreateOrderForDeleteDBNodesRequest {
	s.ZoneId = &v
	return s
}

func (s *PreCheckCreateOrderForDeleteDBNodesRequest) Validate() error {
	return dara.Validate(s)
}
