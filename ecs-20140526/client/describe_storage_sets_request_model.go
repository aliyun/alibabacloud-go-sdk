// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStorageSetsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeStorageSetsRequest
	GetClientToken() *string
	SetOwnerAccount(v string) *DescribeStorageSetsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeStorageSetsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeStorageSetsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeStorageSetsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeStorageSetsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeStorageSetsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeStorageSetsRequest
	GetResourceOwnerId() *int64
	SetStorageSetIds(v string) *DescribeStorageSetsRequest
	GetStorageSetIds() *string
	SetStorageSetName(v string) *DescribeStorageSetsRequest
	GetStorageSetName() *string
	SetTag(v []*DescribeStorageSetsRequestTag) *DescribeStorageSetsRequest
	GetTag() []*DescribeStorageSetsRequestTag
	SetZoneId(v string) *DescribeStorageSetsRequest
	GetZoneId() *string
}

type DescribeStorageSetsRequest struct {
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	RegionId             *string                          `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string                          `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                           `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	StorageSetIds        *string                          `json:"StorageSetIds,omitempty" xml:"StorageSetIds,omitempty"`
	StorageSetName       *string                          `json:"StorageSetName,omitempty" xml:"StorageSetName,omitempty"`
	Tag                  []*DescribeStorageSetsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	ZoneId               *string                          `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeStorageSetsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeStorageSetsRequest) GoString() string {
	return s.String()
}

func (s *DescribeStorageSetsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeStorageSetsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeStorageSetsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeStorageSetsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeStorageSetsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeStorageSetsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeStorageSetsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeStorageSetsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeStorageSetsRequest) GetStorageSetIds() *string {
	return s.StorageSetIds
}

func (s *DescribeStorageSetsRequest) GetStorageSetName() *string {
	return s.StorageSetName
}

func (s *DescribeStorageSetsRequest) GetTag() []*DescribeStorageSetsRequestTag {
	return s.Tag
}

func (s *DescribeStorageSetsRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeStorageSetsRequest) SetClientToken(v string) *DescribeStorageSetsRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeStorageSetsRequest) SetOwnerAccount(v string) *DescribeStorageSetsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeStorageSetsRequest) SetOwnerId(v int64) *DescribeStorageSetsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeStorageSetsRequest) SetPageNumber(v int32) *DescribeStorageSetsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeStorageSetsRequest) SetPageSize(v int32) *DescribeStorageSetsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeStorageSetsRequest) SetRegionId(v string) *DescribeStorageSetsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeStorageSetsRequest) SetResourceOwnerAccount(v string) *DescribeStorageSetsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeStorageSetsRequest) SetResourceOwnerId(v int64) *DescribeStorageSetsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeStorageSetsRequest) SetStorageSetIds(v string) *DescribeStorageSetsRequest {
	s.StorageSetIds = &v
	return s
}

func (s *DescribeStorageSetsRequest) SetStorageSetName(v string) *DescribeStorageSetsRequest {
	s.StorageSetName = &v
	return s
}

func (s *DescribeStorageSetsRequest) SetTag(v []*DescribeStorageSetsRequestTag) *DescribeStorageSetsRequest {
	s.Tag = v
	return s
}

func (s *DescribeStorageSetsRequest) SetZoneId(v string) *DescribeStorageSetsRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeStorageSetsRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeStorageSetsRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeStorageSetsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeStorageSetsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeStorageSetsRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeStorageSetsRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeStorageSetsRequestTag) SetKey(v string) *DescribeStorageSetsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeStorageSetsRequestTag) SetValue(v string) *DescribeStorageSetsRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeStorageSetsRequestTag) Validate() error {
	return dara.Validate(s)
}
