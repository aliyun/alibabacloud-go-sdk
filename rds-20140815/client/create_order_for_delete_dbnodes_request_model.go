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
	AutoPay      *bool   `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	BusinessInfo *string `json:"BusinessInfo,omitempty" xml:"BusinessInfo,omitempty"`
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	CommodityCode        *string   `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	DBInstanceId         *string   `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DBNodeId             []*string `json:"DBNodeId,omitempty" xml:"DBNodeId,omitempty" type:"Repeated"`
	EngineVersion        *string   `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	NodeType             *string   `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PromotionCode        *string   `json:"PromotionCode,omitempty" xml:"PromotionCode,omitempty"`
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Resource             *string   `json:"Resource,omitempty" xml:"Resource,omitempty"`
	ResourceGroupId      *string   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ZoneId               *string   `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
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
