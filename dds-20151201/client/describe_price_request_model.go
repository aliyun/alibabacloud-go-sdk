// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePriceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessInfo(v string) *DescribePriceRequest
	GetBusinessInfo() *string
	SetCommodityCode(v string) *DescribePriceRequest
	GetCommodityCode() *string
	SetCouponNo(v string) *DescribePriceRequest
	GetCouponNo() *string
	SetDBInstances(v string) *DescribePriceRequest
	GetDBInstances() *string
	SetOrderParamOut(v string) *DescribePriceRequest
	GetOrderParamOut() *string
	SetOrderType(v string) *DescribePriceRequest
	GetOrderType() *string
	SetOwnerAccount(v string) *DescribePriceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribePriceRequest
	GetOwnerId() *int64
	SetProductCode(v string) *DescribePriceRequest
	GetProductCode() *string
	SetRegionId(v string) *DescribePriceRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribePriceRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribePriceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribePriceRequest
	GetResourceOwnerId() *int64
}

type DescribePriceRequest struct {
	// The business information. This is an additional parameter.
	//
	// example:
	//
	// {"AccountPassword":"Pw123456","DBInstanceDescription":"test"}
	BusinessInfo *string `json:"BusinessInfo,omitempty" xml:"BusinessInfo,omitempty"`
	// The commodity code of the instance. Valid values:
	//
	// - **dds**: pay-as-you-go ReplicaSet instance.
	//
	// - **badds**: subscription ReplicaSet instance.
	//
	// - **dds_sharding**: pay-as-you-go sharded cluster instance.
	//
	// - **badds_sharding**: subscription sharded cluster instance.
	//
	// - **badds_sharding_intl**: subscription sharded cluster instance on the Alibaba Cloud International Website (www\\.alibabacloud.com).
	//
	// - **dds_sharding_intl**: pay-as-you-go sharded cluster instance on the Alibaba Cloud International Website (www\\.alibabacloud.com).
	//
	// - **badds_sharding_jp**: subscription sharded cluster instance on the Alibaba Cloud Japan Website.
	//
	// - **badds_intl**: subscription ReplicaSet instance on the Alibaba Cloud International Website (www\\.alibabacloud.com).
	//
	// - **dds_intl**: pay-as-you-go ReplicaSet instance on the Alibaba Cloud International Website (www\\.alibabacloud.com).
	//
	// example:
	//
	// badds
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// Specifies whether to use a coupon. Valid values:
	//
	// - **default*	- or **null*	- (default): A coupon is used.
	//
	// - **youhuiquan_promotion_option_id_for_blank**: A coupon is not used.
	//
	// example:
	//
	// default
	CouponNo *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
	// A JSON string that contains information about the instance. For more information about the parameters and JSON examples, see [DBInstances parameter of the DescribePrice operation](https://help.aliyun.com/document_detail/197291.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// [ { "DBInstanceId":"dds-bp1b6e54e7cc****", "RegionId":"cn-hangzhou", "ZoneId":"cn-hangzhou-h", "Engine":"MongoDB", "EngineVersion":" 5.0", "DBInstanceClass":"mdb.shard.2x.xlarge.d", "DBInstanceStorage":30, "ChargeType":"PrePaid", "Period":1, "StorageType":"cloud_essd1" } ]
	DBInstances *string `json:"DBInstances,omitempty" xml:"DBInstances,omitempty"`
	// Specifies whether to return the order parameters. Valid values:
	//
	// - **false*	- (default): The order parameters are not returned.
	//
	// - **true**: The order parameters are returned.
	//
	// example:
	//
	// true
	OrderParamOut *string `json:"OrderParamOut,omitempty" xml:"OrderParamOut,omitempty"`
	// The order type. Valid values:
	//
	// - **BUY**: Creates an instance.
	//
	// - **UPGRADE**: Changes the configuration of an instance.
	//
	// - **RENEW**: Renews an instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// BUY
	OrderType    *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The product code. The default value is **dds**.
	//
	// example:
	//
	// dds
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The region ID. Call the [DescribeRegions](https://help.aliyun.com/document_detail/61933.html) operation to query the region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID. For more information about resource groups, see [View basic information of a resource group](https://help.aliyun.com/document_detail/151181.html).
	//
	// example:
	//
	// rg-acfmyiu4ekp****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribePriceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceRequest) GoString() string {
	return s.String()
}

func (s *DescribePriceRequest) GetBusinessInfo() *string {
	return s.BusinessInfo
}

func (s *DescribePriceRequest) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *DescribePriceRequest) GetCouponNo() *string {
	return s.CouponNo
}

func (s *DescribePriceRequest) GetDBInstances() *string {
	return s.DBInstances
}

func (s *DescribePriceRequest) GetOrderParamOut() *string {
	return s.OrderParamOut
}

func (s *DescribePriceRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *DescribePriceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribePriceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribePriceRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribePriceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePriceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribePriceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribePriceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribePriceRequest) SetBusinessInfo(v string) *DescribePriceRequest {
	s.BusinessInfo = &v
	return s
}

func (s *DescribePriceRequest) SetCommodityCode(v string) *DescribePriceRequest {
	s.CommodityCode = &v
	return s
}

func (s *DescribePriceRequest) SetCouponNo(v string) *DescribePriceRequest {
	s.CouponNo = &v
	return s
}

func (s *DescribePriceRequest) SetDBInstances(v string) *DescribePriceRequest {
	s.DBInstances = &v
	return s
}

func (s *DescribePriceRequest) SetOrderParamOut(v string) *DescribePriceRequest {
	s.OrderParamOut = &v
	return s
}

func (s *DescribePriceRequest) SetOrderType(v string) *DescribePriceRequest {
	s.OrderType = &v
	return s
}

func (s *DescribePriceRequest) SetOwnerAccount(v string) *DescribePriceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribePriceRequest) SetOwnerId(v int64) *DescribePriceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribePriceRequest) SetProductCode(v string) *DescribePriceRequest {
	s.ProductCode = &v
	return s
}

func (s *DescribePriceRequest) SetRegionId(v string) *DescribePriceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePriceRequest) SetResourceGroupId(v string) *DescribePriceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribePriceRequest) SetResourceOwnerAccount(v string) *DescribePriceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribePriceRequest) SetResourceOwnerId(v int64) *DescribePriceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribePriceRequest) Validate() error {
	return dara.Validate(s)
}
