// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStorageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DeleteStorageRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *DeleteStorageRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *DeleteStorageRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *DeleteStorageRequest
	GetResourceOwnerId() *string
	SetResourceRealOwnerId(v int64) *DeleteStorageRequest
	GetResourceRealOwnerId() *int64
	SetStorageLocation(v string) *DeleteStorageRequest
	GetStorageLocation() *string
}

type DeleteStorageRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceRealOwnerId  *int64  `json:"ResourceRealOwnerId,omitempty" xml:"ResourceRealOwnerId,omitempty"`
	// This parameter is required.
	StorageLocation *string `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
}

func (s DeleteStorageRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteStorageRequest) GoString() string {
	return s.String()
}

func (s *DeleteStorageRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteStorageRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DeleteStorageRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteStorageRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *DeleteStorageRequest) GetResourceRealOwnerId() *int64 {
	return s.ResourceRealOwnerId
}

func (s *DeleteStorageRequest) GetStorageLocation() *string {
	return s.StorageLocation
}

func (s *DeleteStorageRequest) SetOwnerAccount(v string) *DeleteStorageRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteStorageRequest) SetOwnerId(v string) *DeleteStorageRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteStorageRequest) SetResourceOwnerAccount(v string) *DeleteStorageRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteStorageRequest) SetResourceOwnerId(v string) *DeleteStorageRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteStorageRequest) SetResourceRealOwnerId(v int64) *DeleteStorageRequest {
	s.ResourceRealOwnerId = &v
	return s
}

func (s *DeleteStorageRequest) SetStorageLocation(v string) *DeleteStorageRequest {
	s.StorageLocation = &v
	return s
}

func (s *DeleteStorageRequest) Validate() error {
	return dara.Validate(s)
}
