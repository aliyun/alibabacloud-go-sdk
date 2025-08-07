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
	// The code of the commodity. Valid values:
	//
	// 	- polardb_sub: the subscription cluster in regions in the Chinese mainland
	//
	// 	- polardb_sub _intl: the subscription cluster in regions outside the Chinese mainland
	//
	// 	- polardb_payg: the pay-as-you-go cluster in regions in the Chinese mainland
	//
	// 	- polardb_payg_intl: the pay-as-you-go cluster in regions outside the Chinese mainland
	//
	// 	- polardb_sub_jushita: the subscription cluster for CloudTmall
	//
	// 	- polardb_payg_jushita: the pay-as-you-go cluster for CloudTmall
	//
	// 	- polardb_sub_cainiao: the subscription cluster for Cainiao
	//
	// 	- polardb_payg_cainiao: the pay-as-you-go cluster for Cainiao
	//
	// > 	- If you use an Alibaba Cloud account on the China site, you can view only the codes of the commodities that are available in the Chinese mainland.
	//
	// >	- If you are using an Alibaba Cloud international account, you can view only the codes of the commodities that are available outside the Chinese mainland.
	//
	// >	- If you use a CloudTmall account, you can view only the codes of the commodities that are available in CloudTmall.
	//
	// >	- If you use a Cainiao account, you can view only the codes of the commodities that are available in Cainiao.
	//
	// This parameter is required.
	//
	// example:
	//
	// polardb_sub
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The number of nodes. Valid values:
	//
	// 	- single: Standalone Edition.
	//
	// 	- cluster: Cluster Edition.
	//
	// 	- all: both Standalone Edition and Cluster Edition.
	//
	// example:
	//
	// cluster
	MasterHa *string `json:"MasterHa,omitempty" xml:"MasterHa,omitempty"`
	// The type of the order. Valid values:
	//
	// 	- BUY: The order is used to purchase a cluster.
	//
	// 	- UPGRADE: The order is used to change the specifications of a cluster.
	//
	// 	- RENEW: The order is used to renew a cluster.
	//
	// 	- CONVERT: The order is used to change the billing method of a cluster.
	//
	// example:
	//
	// BUY
	OrderType    *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the cluster.
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
