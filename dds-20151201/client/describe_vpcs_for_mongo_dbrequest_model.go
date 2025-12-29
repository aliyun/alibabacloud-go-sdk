// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcsForMongoDBRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeVpcsForMongoDBRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeVpcsForMongoDBRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeVpcsForMongoDBRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeVpcsForMongoDBRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeVpcsForMongoDBRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeVpcsForMongoDBRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeVpcsForMongoDBRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeVpcsForMongoDBRequest
	GetResourceOwnerId() *int64
	SetZoneId(v string) *DescribeVpcsForMongoDBRequest
	GetZoneId() *string
}

type DescribeVpcsForMongoDBRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-acfmyiu4ekp****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeVpcsForMongoDBRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcsForMongoDBRequest) GoString() string {
	return s.String()
}

func (s *DescribeVpcsForMongoDBRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeVpcsForMongoDBRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVpcsForMongoDBRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeVpcsForMongoDBRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeVpcsForMongoDBRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeVpcsForMongoDBRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeVpcsForMongoDBRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeVpcsForMongoDBRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeVpcsForMongoDBRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeVpcsForMongoDBRequest) SetOwnerAccount(v string) *DescribeVpcsForMongoDBRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeVpcsForMongoDBRequest) SetOwnerId(v int64) *DescribeVpcsForMongoDBRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVpcsForMongoDBRequest) SetPageNumber(v int32) *DescribeVpcsForMongoDBRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeVpcsForMongoDBRequest) SetPageSize(v int32) *DescribeVpcsForMongoDBRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVpcsForMongoDBRequest) SetRegionId(v string) *DescribeVpcsForMongoDBRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeVpcsForMongoDBRequest) SetResourceGroupId(v string) *DescribeVpcsForMongoDBRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeVpcsForMongoDBRequest) SetResourceOwnerAccount(v string) *DescribeVpcsForMongoDBRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeVpcsForMongoDBRequest) SetResourceOwnerId(v int64) *DescribeVpcsForMongoDBRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeVpcsForMongoDBRequest) SetZoneId(v string) *DescribeVpcsForMongoDBRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeVpcsForMongoDBRequest) Validate() error {
	return dara.Validate(s)
}
