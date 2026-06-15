// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStorageSetDetailsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeStorageSetDetailsRequest
	GetClientToken() *string
	SetDiskIds(v string) *DescribeStorageSetDetailsRequest
	GetDiskIds() *string
	SetOwnerAccount(v string) *DescribeStorageSetDetailsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeStorageSetDetailsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeStorageSetDetailsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeStorageSetDetailsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeStorageSetDetailsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeStorageSetDetailsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeStorageSetDetailsRequest
	GetResourceOwnerId() *int64
	SetStorageSetId(v string) *DescribeStorageSetDetailsRequest
	GetStorageSetId() *string
	SetStorageSetPartitionNumber(v int32) *DescribeStorageSetDetailsRequest
	GetStorageSetPartitionNumber() *int32
}

type DescribeStorageSetDetailsRequest struct {
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DiskIds      *string `json:"DiskIds,omitempty" xml:"DiskIds,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	StorageSetId              *string `json:"StorageSetId,omitempty" xml:"StorageSetId,omitempty"`
	StorageSetPartitionNumber *int32  `json:"StorageSetPartitionNumber,omitempty" xml:"StorageSetPartitionNumber,omitempty"`
}

func (s DescribeStorageSetDetailsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeStorageSetDetailsRequest) GoString() string {
	return s.String()
}

func (s *DescribeStorageSetDetailsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeStorageSetDetailsRequest) GetDiskIds() *string {
	return s.DiskIds
}

func (s *DescribeStorageSetDetailsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeStorageSetDetailsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeStorageSetDetailsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeStorageSetDetailsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeStorageSetDetailsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeStorageSetDetailsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeStorageSetDetailsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeStorageSetDetailsRequest) GetStorageSetId() *string {
	return s.StorageSetId
}

func (s *DescribeStorageSetDetailsRequest) GetStorageSetPartitionNumber() *int32 {
	return s.StorageSetPartitionNumber
}

func (s *DescribeStorageSetDetailsRequest) SetClientToken(v string) *DescribeStorageSetDetailsRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeStorageSetDetailsRequest) SetDiskIds(v string) *DescribeStorageSetDetailsRequest {
	s.DiskIds = &v
	return s
}

func (s *DescribeStorageSetDetailsRequest) SetOwnerAccount(v string) *DescribeStorageSetDetailsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeStorageSetDetailsRequest) SetOwnerId(v int64) *DescribeStorageSetDetailsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeStorageSetDetailsRequest) SetPageNumber(v int32) *DescribeStorageSetDetailsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeStorageSetDetailsRequest) SetPageSize(v int32) *DescribeStorageSetDetailsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeStorageSetDetailsRequest) SetRegionId(v string) *DescribeStorageSetDetailsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeStorageSetDetailsRequest) SetResourceOwnerAccount(v string) *DescribeStorageSetDetailsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeStorageSetDetailsRequest) SetResourceOwnerId(v int64) *DescribeStorageSetDetailsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeStorageSetDetailsRequest) SetStorageSetId(v string) *DescribeStorageSetDetailsRequest {
	s.StorageSetId = &v
	return s
}

func (s *DescribeStorageSetDetailsRequest) SetStorageSetPartitionNumber(v int32) *DescribeStorageSetDetailsRequest {
	s.StorageSetPartitionNumber = &v
	return s
}

func (s *DescribeStorageSetDetailsRequest) Validate() error {
	return dara.Validate(s)
}
