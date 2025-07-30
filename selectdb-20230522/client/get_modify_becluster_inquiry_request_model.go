// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModifyBEClusterInquiryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCacheSize(v int64) *GetModifyBEClusterInquiryRequest
	GetCacheSize() *int64
	SetChargeType(v string) *GetModifyBEClusterInquiryRequest
	GetChargeType() *string
	SetClusterId(v string) *GetModifyBEClusterInquiryRequest
	GetClusterId() *string
	SetCommodityCode(v string) *GetModifyBEClusterInquiryRequest
	GetCommodityCode() *string
	SetComputeSize(v int64) *GetModifyBEClusterInquiryRequest
	GetComputeSize() *int64
	SetDbInstanceId(v string) *GetModifyBEClusterInquiryRequest
	GetDbInstanceId() *string
	SetModifyClusterChargeType(v bool) *GetModifyBEClusterInquiryRequest
	GetModifyClusterChargeType() *bool
	SetPreCacheSize(v int64) *GetModifyBEClusterInquiryRequest
	GetPreCacheSize() *int64
	SetPreComputeSize(v int64) *GetModifyBEClusterInquiryRequest
	GetPreComputeSize() *int64
	SetPricingCycle(v string) *GetModifyBEClusterInquiryRequest
	GetPricingCycle() *string
	SetQuantity(v int64) *GetModifyBEClusterInquiryRequest
	GetQuantity() *int64
	SetRegionId(v string) *GetModifyBEClusterInquiryRequest
	GetRegionId() *string
	SetResourceOwnerId(v int64) *GetModifyBEClusterInquiryRequest
	GetResourceOwnerId() *int64
}

type GetModifyBEClusterInquiryRequest struct {
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
	// The cluster ID.
	//
	// example:
	//
	// selectdb-xxx-be
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
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
	// 	- selectdb_pre_public_cn: subscription commodity on the China site (aliyun.com)
	//
	// This parameter is required.
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
	// Specifies whether to change the billing method of the cluster.
	//
	// example:
	//
	// true
	ModifyClusterChargeType *bool `json:"ModifyClusterChargeType,omitempty" xml:"ModifyClusterChargeType,omitempty"`
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
	// The number of clusters whose specifications are to be changed.
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

func (s GetModifyBEClusterInquiryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetModifyBEClusterInquiryRequest) GoString() string {
	return s.String()
}

func (s *GetModifyBEClusterInquiryRequest) GetCacheSize() *int64 {
	return s.CacheSize
}

func (s *GetModifyBEClusterInquiryRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *GetModifyBEClusterInquiryRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetModifyBEClusterInquiryRequest) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *GetModifyBEClusterInquiryRequest) GetComputeSize() *int64 {
	return s.ComputeSize
}

func (s *GetModifyBEClusterInquiryRequest) GetDbInstanceId() *string {
	return s.DbInstanceId
}

func (s *GetModifyBEClusterInquiryRequest) GetModifyClusterChargeType() *bool {
	return s.ModifyClusterChargeType
}

func (s *GetModifyBEClusterInquiryRequest) GetPreCacheSize() *int64 {
	return s.PreCacheSize
}

func (s *GetModifyBEClusterInquiryRequest) GetPreComputeSize() *int64 {
	return s.PreComputeSize
}

func (s *GetModifyBEClusterInquiryRequest) GetPricingCycle() *string {
	return s.PricingCycle
}

func (s *GetModifyBEClusterInquiryRequest) GetQuantity() *int64 {
	return s.Quantity
}

func (s *GetModifyBEClusterInquiryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetModifyBEClusterInquiryRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetModifyBEClusterInquiryRequest) SetCacheSize(v int64) *GetModifyBEClusterInquiryRequest {
	s.CacheSize = &v
	return s
}

func (s *GetModifyBEClusterInquiryRequest) SetChargeType(v string) *GetModifyBEClusterInquiryRequest {
	s.ChargeType = &v
	return s
}

func (s *GetModifyBEClusterInquiryRequest) SetClusterId(v string) *GetModifyBEClusterInquiryRequest {
	s.ClusterId = &v
	return s
}

func (s *GetModifyBEClusterInquiryRequest) SetCommodityCode(v string) *GetModifyBEClusterInquiryRequest {
	s.CommodityCode = &v
	return s
}

func (s *GetModifyBEClusterInquiryRequest) SetComputeSize(v int64) *GetModifyBEClusterInquiryRequest {
	s.ComputeSize = &v
	return s
}

func (s *GetModifyBEClusterInquiryRequest) SetDbInstanceId(v string) *GetModifyBEClusterInquiryRequest {
	s.DbInstanceId = &v
	return s
}

func (s *GetModifyBEClusterInquiryRequest) SetModifyClusterChargeType(v bool) *GetModifyBEClusterInquiryRequest {
	s.ModifyClusterChargeType = &v
	return s
}

func (s *GetModifyBEClusterInquiryRequest) SetPreCacheSize(v int64) *GetModifyBEClusterInquiryRequest {
	s.PreCacheSize = &v
	return s
}

func (s *GetModifyBEClusterInquiryRequest) SetPreComputeSize(v int64) *GetModifyBEClusterInquiryRequest {
	s.PreComputeSize = &v
	return s
}

func (s *GetModifyBEClusterInquiryRequest) SetPricingCycle(v string) *GetModifyBEClusterInquiryRequest {
	s.PricingCycle = &v
	return s
}

func (s *GetModifyBEClusterInquiryRequest) SetQuantity(v int64) *GetModifyBEClusterInquiryRequest {
	s.Quantity = &v
	return s
}

func (s *GetModifyBEClusterInquiryRequest) SetRegionId(v string) *GetModifyBEClusterInquiryRequest {
	s.RegionId = &v
	return s
}

func (s *GetModifyBEClusterInquiryRequest) SetResourceOwnerId(v int64) *GetModifyBEClusterInquiryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetModifyBEClusterInquiryRequest) Validate() error {
	return dara.Validate(s)
}
