// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSnatEntriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeSnatEntriesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSnatEntriesRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeSnatEntriesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSnatEntriesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeSnatEntriesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeSnatEntriesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSnatEntriesRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *DescribeSnatEntriesRequest
	GetSmartAGId() *string
}

type DescribeSnatEntriesRequest struct {
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
	// 2
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region where the SAG instance is deployed.
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
	// This parameter is required.
	//
	// example:
	//
	// sag-whfn************
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
}

func (s DescribeSnatEntriesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnatEntriesRequest) GoString() string {
	return s.String()
}

func (s *DescribeSnatEntriesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSnatEntriesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSnatEntriesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSnatEntriesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSnatEntriesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSnatEntriesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSnatEntriesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSnatEntriesRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *DescribeSnatEntriesRequest) SetOwnerAccount(v string) *DescribeSnatEntriesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSnatEntriesRequest) SetOwnerId(v int64) *DescribeSnatEntriesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSnatEntriesRequest) SetPageNumber(v int32) *DescribeSnatEntriesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSnatEntriesRequest) SetPageSize(v int32) *DescribeSnatEntriesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSnatEntriesRequest) SetRegionId(v string) *DescribeSnatEntriesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSnatEntriesRequest) SetResourceOwnerAccount(v string) *DescribeSnatEntriesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSnatEntriesRequest) SetResourceOwnerId(v int64) *DescribeSnatEntriesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSnatEntriesRequest) SetSmartAGId(v string) *DescribeSnatEntriesRequest {
	s.SmartAGId = &v
	return s
}

func (s *DescribeSnatEntriesRequest) Validate() error {
	return dara.Validate(s)
}
