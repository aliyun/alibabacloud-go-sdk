// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDefaultUploadStorageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *SetDefaultUploadStorageRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *SetDefaultUploadStorageRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *SetDefaultUploadStorageRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *SetDefaultUploadStorageRequest
	GetResourceOwnerId() *string
	SetResourceRealOwnerId(v int64) *SetDefaultUploadStorageRequest
	GetResourceRealOwnerId() *int64
	SetStorageLocation(v string) *SetDefaultUploadStorageRequest
	GetStorageLocation() *string
}

type SetDefaultUploadStorageRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceRealOwnerId  *int64  `json:"ResourceRealOwnerId,omitempty" xml:"ResourceRealOwnerId,omitempty"`
	// This parameter is required.
	StorageLocation *string `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
}

func (s SetDefaultUploadStorageRequest) String() string {
	return dara.Prettify(s)
}

func (s SetDefaultUploadStorageRequest) GoString() string {
	return s.String()
}

func (s *SetDefaultUploadStorageRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SetDefaultUploadStorageRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *SetDefaultUploadStorageRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SetDefaultUploadStorageRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *SetDefaultUploadStorageRequest) GetResourceRealOwnerId() *int64 {
	return s.ResourceRealOwnerId
}

func (s *SetDefaultUploadStorageRequest) GetStorageLocation() *string {
	return s.StorageLocation
}

func (s *SetDefaultUploadStorageRequest) SetOwnerAccount(v string) *SetDefaultUploadStorageRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SetDefaultUploadStorageRequest) SetOwnerId(v string) *SetDefaultUploadStorageRequest {
	s.OwnerId = &v
	return s
}

func (s *SetDefaultUploadStorageRequest) SetResourceOwnerAccount(v string) *SetDefaultUploadStorageRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SetDefaultUploadStorageRequest) SetResourceOwnerId(v string) *SetDefaultUploadStorageRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SetDefaultUploadStorageRequest) SetResourceRealOwnerId(v int64) *SetDefaultUploadStorageRequest {
	s.ResourceRealOwnerId = &v
	return s
}

func (s *SetDefaultUploadStorageRequest) SetStorageLocation(v string) *SetDefaultUploadStorageRequest {
	s.StorageLocation = &v
	return s
}

func (s *SetDefaultUploadStorageRequest) Validate() error {
	return dara.Validate(s)
}
