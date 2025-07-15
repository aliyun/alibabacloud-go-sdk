// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIPv6TranslatorsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllocateIpv4Addr(v string) *DescribeIPv6TranslatorsRequest
	GetAllocateIpv4Addr() *string
	SetAllocateIpv6Addr(v string) *DescribeIPv6TranslatorsRequest
	GetAllocateIpv6Addr() *string
	SetBusinessStatus(v string) *DescribeIPv6TranslatorsRequest
	GetBusinessStatus() *string
	SetIpv6TranslatorId(v string) *DescribeIPv6TranslatorsRequest
	GetIpv6TranslatorId() *string
	SetName(v string) *DescribeIPv6TranslatorsRequest
	GetName() *string
	SetOwnerAccount(v string) *DescribeIPv6TranslatorsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeIPv6TranslatorsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeIPv6TranslatorsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeIPv6TranslatorsRequest
	GetPageSize() *int32
	SetPayType(v string) *DescribeIPv6TranslatorsRequest
	GetPayType() *string
	SetRegionId(v string) *DescribeIPv6TranslatorsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeIPv6TranslatorsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeIPv6TranslatorsRequest
	GetResourceOwnerId() *int64
	SetSpec(v string) *DescribeIPv6TranslatorsRequest
	GetSpec() *string
	SetStatus(v string) *DescribeIPv6TranslatorsRequest
	GetStatus() *string
}

type DescribeIPv6TranslatorsRequest struct {
	// The IPv4 address allocated to the IPv6 Translation Service instance.
	//
	// example:
	//
	// 47.99.XX.XX
	AllocateIpv4Addr *string `json:"AllocateIpv4Addr,omitempty" xml:"AllocateIpv4Addr,omitempty"`
	// The IPv6 address allocated to the IPv6 Translation Service instance.
	//
	// example:
	//
	// 2400:3200:1600::XXXX
	AllocateIpv6Addr *string `json:"AllocateIpv6Addr,omitempty" xml:"AllocateIpv6Addr,omitempty"`
	// The business status of the IPv6 Translation Service instance. Valid values:
	//
	// 	- **Normal**
	//
	// 	- **FinancialLocked**
	//
	// example:
	//
	// Normal
	BusinessStatus *string `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
	// The ID of the IPv6 Translation Service instance.
	//
	// example:
	//
	// ipv6trans-bp1858ys****
	Ipv6TranslatorId *string `json:"Ipv6TranslatorId,omitempty" xml:"Ipv6TranslatorId,omitempty"`
	// The name of the IPv6 Translation Service instance.
	//
	// example:
	//
	// ipv6_1
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Maximum value: **50**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The billing method of the IPv6 Translation Service instance. Valid values:
	//
	// 	- **Prepay**: subscription
	//
	// 	- **Postpay**: pay-as-you-go
	//
	// example:
	//
	// Prepay
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The region of the IPv6 Translation Service instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The specification of the IPv6 Translation Service instance. Set the value to **small**.
	//
	// example:
	//
	// small
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
	// The status of the IPv6 Translation Service instance. Valid values:
	//
	// 	- **init**
	//
	// 	- **provisioning**
	//
	// 	- **active**
	//
	// 	- **updating**
	//
	// 	- **upgrading**
	//
	// 	- **deleting**
	//
	// 	- **deleted**
	//
	// example:
	//
	// active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeIPv6TranslatorsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeIPv6TranslatorsRequest) GoString() string {
	return s.String()
}

func (s *DescribeIPv6TranslatorsRequest) GetAllocateIpv4Addr() *string {
	return s.AllocateIpv4Addr
}

func (s *DescribeIPv6TranslatorsRequest) GetAllocateIpv6Addr() *string {
	return s.AllocateIpv6Addr
}

func (s *DescribeIPv6TranslatorsRequest) GetBusinessStatus() *string {
	return s.BusinessStatus
}

func (s *DescribeIPv6TranslatorsRequest) GetIpv6TranslatorId() *string {
	return s.Ipv6TranslatorId
}

func (s *DescribeIPv6TranslatorsRequest) GetName() *string {
	return s.Name
}

func (s *DescribeIPv6TranslatorsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeIPv6TranslatorsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeIPv6TranslatorsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeIPv6TranslatorsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeIPv6TranslatorsRequest) GetPayType() *string {
	return s.PayType
}

func (s *DescribeIPv6TranslatorsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeIPv6TranslatorsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeIPv6TranslatorsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeIPv6TranslatorsRequest) GetSpec() *string {
	return s.Spec
}

func (s *DescribeIPv6TranslatorsRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeIPv6TranslatorsRequest) SetAllocateIpv4Addr(v string) *DescribeIPv6TranslatorsRequest {
	s.AllocateIpv4Addr = &v
	return s
}

func (s *DescribeIPv6TranslatorsRequest) SetAllocateIpv6Addr(v string) *DescribeIPv6TranslatorsRequest {
	s.AllocateIpv6Addr = &v
	return s
}

func (s *DescribeIPv6TranslatorsRequest) SetBusinessStatus(v string) *DescribeIPv6TranslatorsRequest {
	s.BusinessStatus = &v
	return s
}

func (s *DescribeIPv6TranslatorsRequest) SetIpv6TranslatorId(v string) *DescribeIPv6TranslatorsRequest {
	s.Ipv6TranslatorId = &v
	return s
}

func (s *DescribeIPv6TranslatorsRequest) SetName(v string) *DescribeIPv6TranslatorsRequest {
	s.Name = &v
	return s
}

func (s *DescribeIPv6TranslatorsRequest) SetOwnerAccount(v string) *DescribeIPv6TranslatorsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeIPv6TranslatorsRequest) SetOwnerId(v int64) *DescribeIPv6TranslatorsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeIPv6TranslatorsRequest) SetPageNumber(v int32) *DescribeIPv6TranslatorsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeIPv6TranslatorsRequest) SetPageSize(v int32) *DescribeIPv6TranslatorsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeIPv6TranslatorsRequest) SetPayType(v string) *DescribeIPv6TranslatorsRequest {
	s.PayType = &v
	return s
}

func (s *DescribeIPv6TranslatorsRequest) SetRegionId(v string) *DescribeIPv6TranslatorsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeIPv6TranslatorsRequest) SetResourceOwnerAccount(v string) *DescribeIPv6TranslatorsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeIPv6TranslatorsRequest) SetResourceOwnerId(v int64) *DescribeIPv6TranslatorsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeIPv6TranslatorsRequest) SetSpec(v string) *DescribeIPv6TranslatorsRequest {
	s.Spec = &v
	return s
}

func (s *DescribeIPv6TranslatorsRequest) SetStatus(v string) *DescribeIPv6TranslatorsRequest {
	s.Status = &v
	return s
}

func (s *DescribeIPv6TranslatorsRequest) Validate() error {
	return dara.Validate(s)
}
