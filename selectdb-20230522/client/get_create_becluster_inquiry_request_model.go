// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCreateBEClusterInquiryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCacheSize(v int64) *GetCreateBEClusterInquiryRequest
	GetCacheSize() *int64
	SetChargeType(v string) *GetCreateBEClusterInquiryRequest
	GetChargeType() *string
	SetCommodityCode(v string) *GetCreateBEClusterInquiryRequest
	GetCommodityCode() *string
	SetComputeSize(v int64) *GetCreateBEClusterInquiryRequest
	GetComputeSize() *int64
	SetDbInstanceId(v string) *GetCreateBEClusterInquiryRequest
	GetDbInstanceId() *string
	SetPreCacheSize(v int64) *GetCreateBEClusterInquiryRequest
	GetPreCacheSize() *int64
	SetPreComputeSize(v int64) *GetCreateBEClusterInquiryRequest
	GetPreComputeSize() *int64
	SetPricingCycle(v string) *GetCreateBEClusterInquiryRequest
	GetPricingCycle() *string
	SetQuantity(v int64) *GetCreateBEClusterInquiryRequest
	GetQuantity() *int64
	SetRegionId(v string) *GetCreateBEClusterInquiryRequest
	GetRegionId() *string
	SetResourceOwnerId(v int64) *GetCreateBEClusterInquiryRequest
	GetResourceOwnerId() *int64
}

type GetCreateBEClusterInquiryRequest struct {
	// The size of the elastic cache.
	//
	// example:
	//
	// 200
	CacheSize *int64 `json:"CacheSize,omitempty" xml:"CacheSize,omitempty"`
	// The billing method.
	//
	// Valid values:
	//
	// 	- PREPAY: subscription
	//
	// 	- POSTPAY: pay-as-you-go
	//
	// This parameter is required.
	//
	// example:
	//
	// POSTPAY
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The commodity code.
	//
	// Valid values:
	//
	// 	- selectdb_pre_public_intl: subscription commodity on the international site (alibabacloud.com)
	//
	// 	- selectdb_go_public_cn: pay-as-you-go commodity on the China site (aliyun.com)
	//
	// 	- selectdb_go_public_intl: pay-as-you-go commodity on the international site (alibabacloud.com)
	//
	// 	- selectdb_pre_public_cn: subscription commodity on the China site (aliyun.com).
	//
	// example:
	//
	// selectdb_go_public_cn
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The number of elastic CPU cores.
	//
	// example:
	//
	// 4
	ComputeSize *int64 `json:"ComputeSize,omitempty" xml:"ComputeSize,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// selectdb-cn-xxx
	DbInstanceId *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	// The size of the reserved cache.
	//
	// example:
	//
	// 200
	PreCacheSize *int64 `json:"PreCacheSize,omitempty" xml:"PreCacheSize,omitempty"`
	// The number of reserved CPU cores.
	//
	// example:
	//
	// 4
	PreComputeSize *int64 `json:"PreComputeSize,omitempty" xml:"PreComputeSize,omitempty"`
	// The billing cycle.
	//
	// Valid values:
	//
	// 	- Month
	//
	// 	- Year
	//
	// 	- Minute
	//
	// 	- Hour
	//
	// 	- Day
	//
	// This parameter is required.
	//
	// example:
	//
	// Hour
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// The number of clusters to be created.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Quantity *int64 `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetCreateBEClusterInquiryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCreateBEClusterInquiryRequest) GoString() string {
	return s.String()
}

func (s *GetCreateBEClusterInquiryRequest) GetCacheSize() *int64 {
	return s.CacheSize
}

func (s *GetCreateBEClusterInquiryRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *GetCreateBEClusterInquiryRequest) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *GetCreateBEClusterInquiryRequest) GetComputeSize() *int64 {
	return s.ComputeSize
}

func (s *GetCreateBEClusterInquiryRequest) GetDbInstanceId() *string {
	return s.DbInstanceId
}

func (s *GetCreateBEClusterInquiryRequest) GetPreCacheSize() *int64 {
	return s.PreCacheSize
}

func (s *GetCreateBEClusterInquiryRequest) GetPreComputeSize() *int64 {
	return s.PreComputeSize
}

func (s *GetCreateBEClusterInquiryRequest) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *GetCreateBEClusterInquiryRequest) GetQuantity() *int64 {
	return s.Quantity
}

func (s *GetCreateBEClusterInquiryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetCreateBEClusterInquiryRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetCreateBEClusterInquiryRequest) SetCacheSize(v int64) *GetCreateBEClusterInquiryRequest {
	s.CacheSize = &v
	return s
}

func (s *GetCreateBEClusterInquiryRequest) SetChargeType(v string) *GetCreateBEClusterInquiryRequest {
	s.ChargeType = &v
	return s
}

func (s *GetCreateBEClusterInquiryRequest) SetCommodityCode(v string) *GetCreateBEClusterInquiryRequest {
	s.CommodityCode = &v
	return s
}

func (s *GetCreateBEClusterInquiryRequest) SetComputeSize(v int64) *GetCreateBEClusterInquiryRequest {
	s.ComputeSize = &v
	return s
}

func (s *GetCreateBEClusterInquiryRequest) SetDbInstanceId(v string) *GetCreateBEClusterInquiryRequest {
	s.DbInstanceId = &v
	return s
}

func (s *GetCreateBEClusterInquiryRequest) SetPreCacheSize(v int64) *GetCreateBEClusterInquiryRequest {
	s.PreCacheSize = &v
	return s
}

func (s *GetCreateBEClusterInquiryRequest) SetPreComputeSize(v int64) *GetCreateBEClusterInquiryRequest {
	s.PreComputeSize = &v
	return s
}

func (s *GetCreateBEClusterInquiryRequest) SetPricingCycle(v string) *GetCreateBEClusterInquiryRequest {
	s.PricingCycle = &v
	return s
}

func (s *GetCreateBEClusterInquiryRequest) SetQuantity(v int64) *GetCreateBEClusterInquiryRequest {
	s.Quantity = &v
	return s
}

func (s *GetCreateBEClusterInquiryRequest) SetRegionId(v string) *GetCreateBEClusterInquiryRequest {
	s.RegionId = &v
	return s
}

func (s *GetCreateBEClusterInquiryRequest) SetResourceOwnerId(v int64) *GetCreateBEClusterInquiryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetCreateBEClusterInquiryRequest) Validate() error {
	return dara.Validate(s)
}
