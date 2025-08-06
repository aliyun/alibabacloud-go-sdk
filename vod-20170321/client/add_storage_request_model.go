// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddStorageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *AddStorageRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *AddStorageRequest
	GetOwnerId() *string
	SetResourceGroupId(v string) *AddStorageRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *AddStorageRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *AddStorageRequest
	GetResourceOwnerId() *string
	SetResourceRealOwnerId(v string) *AddStorageRequest
	GetResourceRealOwnerId() *string
	SetStorageLocation(v string) *AddStorageRequest
	GetStorageLocation() *string
	SetStorageRedundancyType(v string) *AddStorageRequest
	GetStorageRedundancyType() *string
	SetStorageRegion(v string) *AddStorageRequest
	GetStorageRegion() *string
	SetStorageType(v string) *AddStorageRequest
	GetStorageType() *string
}

type AddStorageRequest struct {
	OwnerAccount          *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId               *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceGroupId       *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount  *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId       *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceRealOwnerId   *string `json:"ResourceRealOwnerId,omitempty" xml:"ResourceRealOwnerId,omitempty"`
	StorageLocation       *string `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
	StorageRedundancyType *string `json:"StorageRedundancyType,omitempty" xml:"StorageRedundancyType,omitempty"`
	StorageRegion         *string `json:"StorageRegion,omitempty" xml:"StorageRegion,omitempty"`
	StorageType           *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s AddStorageRequest) String() string {
	return dara.Prettify(s)
}

func (s AddStorageRequest) GoString() string {
	return s.String()
}

func (s *AddStorageRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AddStorageRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *AddStorageRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *AddStorageRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddStorageRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *AddStorageRequest) GetResourceRealOwnerId() *string {
	return s.ResourceRealOwnerId
}

func (s *AddStorageRequest) GetStorageLocation() *string {
	return s.StorageLocation
}

func (s *AddStorageRequest) GetStorageRedundancyType() *string {
	return s.StorageRedundancyType
}

func (s *AddStorageRequest) GetStorageRegion() *string {
	return s.StorageRegion
}

func (s *AddStorageRequest) GetStorageType() *string {
	return s.StorageType
}

func (s *AddStorageRequest) SetOwnerAccount(v string) *AddStorageRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AddStorageRequest) SetOwnerId(v string) *AddStorageRequest {
	s.OwnerId = &v
	return s
}

func (s *AddStorageRequest) SetResourceGroupId(v string) *AddStorageRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *AddStorageRequest) SetResourceOwnerAccount(v string) *AddStorageRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddStorageRequest) SetResourceOwnerId(v string) *AddStorageRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddStorageRequest) SetResourceRealOwnerId(v string) *AddStorageRequest {
	s.ResourceRealOwnerId = &v
	return s
}

func (s *AddStorageRequest) SetStorageLocation(v string) *AddStorageRequest {
	s.StorageLocation = &v
	return s
}

func (s *AddStorageRequest) SetStorageRedundancyType(v string) *AddStorageRequest {
	s.StorageRedundancyType = &v
	return s
}

func (s *AddStorageRequest) SetStorageRegion(v string) *AddStorageRequest {
	s.StorageRegion = &v
	return s
}

func (s *AddStorageRequest) SetStorageType(v string) *AddStorageRequest {
	s.StorageType = &v
	return s
}

func (s *AddStorageRequest) Validate() error {
	return dara.Validate(s)
}
