// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLicenseOrdersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunOrderId(v string) *DescribeLicenseOrdersRequest
	GetAliyunOrderId() *string
	SetOwnerAccount(v string) *DescribeLicenseOrdersRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeLicenseOrdersRequest
	GetOwnerId() *int64
	SetPackageType(v string) *DescribeLicenseOrdersRequest
	GetPackageType() *string
	SetPageNumber(v int32) *DescribeLicenseOrdersRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeLicenseOrdersRequest
	GetPageSize() *int32
	SetPurchaseChannel(v string) *DescribeLicenseOrdersRequest
	GetPurchaseChannel() *string
	SetResourceOwnerAccount(v string) *DescribeLicenseOrdersRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeLicenseOrdersRequest
	GetResourceOwnerId() *int64
	SetVirtualOrder(v bool) *DescribeLicenseOrdersRequest
	GetVirtualOrder() *bool
}

type DescribeLicenseOrdersRequest struct {
	// The ID of the Alibaba Cloud order. The value can be the ID of a virtual order.
	//
	// example:
	//
	// 239618016570503
	AliyunOrderId *string `json:"AliyunOrderId,omitempty" xml:"AliyunOrderId,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The plan type. Valid values:
	//
	// 	- single_node_subscribe: Single-node Edition (Subscription).
	//
	// 	- single_node_long_term: Single-node Edition (Long-term).
	//
	// 	- primary_backup_subscribe: HA Edition (Subscription).
	//
	// 	- primary_backup_long_term: HA Edition (Long-term).
	//
	// 	- pre_generation_long_term: Pre-generated (Long-term).
	//
	// example:
	//
	// single_node_subscribe
	PackageType *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The purchase channel. Valid values: aliyun_market and aliyun_public. aliyun_market specifies Alibaba Cloud Marketplace. aliyun_public specifies the PolarDB buy page.
	//
	// example:
	//
	// aliyun_market
	PurchaseChannel      *string `json:"PurchaseChannel,omitempty" xml:"PurchaseChannel,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specifies whether to query only virtual orders.
	//
	// example:
	//
	// true
	VirtualOrder *bool `json:"VirtualOrder,omitempty" xml:"VirtualOrder,omitempty"`
}

func (s DescribeLicenseOrdersRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLicenseOrdersRequest) GoString() string {
	return s.String()
}

func (s *DescribeLicenseOrdersRequest) GetAliyunOrderId() *string {
	return s.AliyunOrderId
}

func (s *DescribeLicenseOrdersRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeLicenseOrdersRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLicenseOrdersRequest) GetPackageType() *string {
	return s.PackageType
}

func (s *DescribeLicenseOrdersRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeLicenseOrdersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeLicenseOrdersRequest) GetPurchaseChannel() *string {
	return s.PurchaseChannel
}

func (s *DescribeLicenseOrdersRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeLicenseOrdersRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeLicenseOrdersRequest) GetVirtualOrder() *bool {
	return s.VirtualOrder
}

func (s *DescribeLicenseOrdersRequest) SetAliyunOrderId(v string) *DescribeLicenseOrdersRequest {
	s.AliyunOrderId = &v
	return s
}

func (s *DescribeLicenseOrdersRequest) SetOwnerAccount(v string) *DescribeLicenseOrdersRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeLicenseOrdersRequest) SetOwnerId(v int64) *DescribeLicenseOrdersRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLicenseOrdersRequest) SetPackageType(v string) *DescribeLicenseOrdersRequest {
	s.PackageType = &v
	return s
}

func (s *DescribeLicenseOrdersRequest) SetPageNumber(v int32) *DescribeLicenseOrdersRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeLicenseOrdersRequest) SetPageSize(v int32) *DescribeLicenseOrdersRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLicenseOrdersRequest) SetPurchaseChannel(v string) *DescribeLicenseOrdersRequest {
	s.PurchaseChannel = &v
	return s
}

func (s *DescribeLicenseOrdersRequest) SetResourceOwnerAccount(v string) *DescribeLicenseOrdersRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeLicenseOrdersRequest) SetResourceOwnerId(v int64) *DescribeLicenseOrdersRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeLicenseOrdersRequest) SetVirtualOrder(v bool) *DescribeLicenseOrdersRequest {
	s.VirtualOrder = &v
	return s
}

func (s *DescribeLicenseOrdersRequest) Validate() error {
	return dara.Validate(s)
}
