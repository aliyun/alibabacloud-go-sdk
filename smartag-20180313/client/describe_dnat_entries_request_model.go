// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDnatEntriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeDnatEntriesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDnatEntriesRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeDnatEntriesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDnatEntriesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeDnatEntriesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeDnatEntriesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDnatEntriesRequest
	GetResourceOwnerId() *int64
	SetSagId(v string) *DescribeDnatEntriesRequest
	GetSagId() *string
	SetType(v string) *DescribeDnatEntriesRequest
	GetType() *string
}

type DescribeDnatEntriesRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: **10**. Maximum value: **50**.
	//
	// example:
	//
	// 12
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the SAG instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the SAG instance.
	//
	// >  Only SAG customer-premises equipment (CPE) instances are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-djgd*************
	SagId *string `json:"SagId,omitempty" xml:"SagId,omitempty"`
	// The type of the DNAT entry. Valid values:
	//
	// 	- **Intranet**: translates the IP address to a specific internal IP address. This is the default value.
	//
	// 	- **Internet**: translates the IP address to a specific public IP address.
	//
	// example:
	//
	// Intranet
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeDnatEntriesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnatEntriesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDnatEntriesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDnatEntriesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDnatEntriesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDnatEntriesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDnatEntriesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDnatEntriesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDnatEntriesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDnatEntriesRequest) GetSagId() *string {
	return s.SagId
}

func (s *DescribeDnatEntriesRequest) GetType() *string {
	return s.Type
}

func (s *DescribeDnatEntriesRequest) SetOwnerAccount(v string) *DescribeDnatEntriesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDnatEntriesRequest) SetOwnerId(v int64) *DescribeDnatEntriesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDnatEntriesRequest) SetPageNumber(v int32) *DescribeDnatEntriesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDnatEntriesRequest) SetPageSize(v int32) *DescribeDnatEntriesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDnatEntriesRequest) SetRegionId(v string) *DescribeDnatEntriesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDnatEntriesRequest) SetResourceOwnerAccount(v string) *DescribeDnatEntriesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDnatEntriesRequest) SetResourceOwnerId(v int64) *DescribeDnatEntriesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDnatEntriesRequest) SetSagId(v string) *DescribeDnatEntriesRequest {
	s.SagId = &v
	return s
}

func (s *DescribeDnatEntriesRequest) SetType(v string) *DescribeDnatEntriesRequest {
	s.Type = &v
	return s
}

func (s *DescribeDnatEntriesRequest) Validate() error {
	return dara.Validate(s)
}
