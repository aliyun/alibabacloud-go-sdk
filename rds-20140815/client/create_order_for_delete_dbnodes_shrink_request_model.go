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
	AutoPay      *bool   `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	BusinessInfo *string `json:"BusinessInfo,omitempty" xml:"BusinessInfo,omitempty"`
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	CommodityCode        *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DBNodeIdShrink       *string `json:"DBNodeId,omitempty" xml:"DBNodeId,omitempty"`
	EngineVersion        *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	NodeType             *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PromotionCode        *string `json:"PromotionCode,omitempty" xml:"PromotionCode,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Resource             *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
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
