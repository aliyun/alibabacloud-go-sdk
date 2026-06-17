// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClassListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommodityCode(v string) *DescribeClassListRequest
	GetCommodityCode() *string
	SetMasterHa(v string) *DescribeClassListRequest
	GetMasterHa() *string
	SetOrderType(v string) *DescribeClassListRequest
	GetOrderType() *string
	SetOwnerAccount(v string) *DescribeClassListRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeClassListRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeClassListRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeClassListRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeClassListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeClassListRequest
	GetResourceOwnerId() *int64
}

type DescribeClassListRequest struct {
	// The commodity code. Valid values:
	//
	// - polardb_sub: subscription for the Chinese mainland.
	//
	// - polardb_sub_intl: subscription for regions in Hong Kong (China) and outside the Chinese mainland.
	//
	// - polardb_payg: pay-as-you-go for the Chinese mainland.
	//
	// - polardb_payg_intl: pay-as-you-go for regions in Hong Kong (China) and outside the Chinese mainland.
	//
	// - polardb_sub_jushita: Jushita subscription.
	//
	// - polardb_payg_jushita: Jushita pay-as-you-go.
	//
	// - polardb_sub_cainiao: Cainiao subscription.
	//
	// - polardb_payg_cainiao: Cainiao pay-as-you-go.
	//
	// > 	- If you use an Alibaba Cloud China site account, you can view only the commodity codes for the Chinese mainland.
	//
	// >
	//
	// > 	- If you use an Alibaba Cloud international site account, you can view only the commodity codes for regions outside the Chinese mainland.
	//
	// >
	//
	// > 	- If you use a Jushita account, you can view only the commodity codes for Jushita.
	//
	// >
	//
	// > 	- If you use a Cainiao account, you can view only the commodity codes for Cainiao.
	//
	// This parameter is required.
	//
	// example:
	//
	// polardb_sub
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The number of nodes. Valid values:
	//
	// - single: single node.
	//
	// - cluster: cluster.
	//
	// - all: single node and cluster.
	//
	// example:
	//
	// cluster
	MasterHa *string `json:"MasterHa,omitempty" xml:"MasterHa,omitempty"`
	// The order type. Valid values:
	//
	// - BUY: new purchase.
	//
	// - UPGRADE: changes the configuration.
	//
	// - RENEW: renews the instance.
	//
	// - CONVERT: changes the billing method.
	//
	// example:
	//
	// BUY
	OrderType    *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// > This parameter is required if you use an Alibaba Cloud international site account.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-************
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeClassListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeClassListRequest) GoString() string {
	return s.String()
}

func (s *DescribeClassListRequest) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *DescribeClassListRequest) GetMasterHa() *string {
	return s.MasterHa
}

func (s *DescribeClassListRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *DescribeClassListRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeClassListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeClassListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeClassListRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeClassListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeClassListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeClassListRequest) SetCommodityCode(v string) *DescribeClassListRequest {
	s.CommodityCode = &v
	return s
}

func (s *DescribeClassListRequest) SetMasterHa(v string) *DescribeClassListRequest {
	s.MasterHa = &v
	return s
}

func (s *DescribeClassListRequest) SetOrderType(v string) *DescribeClassListRequest {
	s.OrderType = &v
	return s
}

func (s *DescribeClassListRequest) SetOwnerAccount(v string) *DescribeClassListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeClassListRequest) SetOwnerId(v int64) *DescribeClassListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeClassListRequest) SetRegionId(v string) *DescribeClassListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeClassListRequest) SetResourceGroupId(v string) *DescribeClassListRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeClassListRequest) SetResourceOwnerAccount(v string) *DescribeClassListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeClassListRequest) SetResourceOwnerId(v int64) *DescribeClassListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeClassListRequest) Validate() error {
	return dara.Validate(s)
}
