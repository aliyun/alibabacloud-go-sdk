// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStorageInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *GetStorageInfoRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *GetStorageInfoRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *GetStorageInfoRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *GetStorageInfoRequest
	GetResourceOwnerId() *string
	SetResourceRealOwnerId(v int64) *GetStorageInfoRequest
	GetResourceRealOwnerId() *int64
	SetStorageLocation(v string) *GetStorageInfoRequest
	GetStorageLocation() *string
}

type GetStorageInfoRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceRealOwnerId  *int64  `json:"ResourceRealOwnerId,omitempty" xml:"ResourceRealOwnerId,omitempty"`
	StorageLocation      *string `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
}

func (s GetStorageInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetStorageInfoRequest) GoString() string {
	return s.String()
}

func (s *GetStorageInfoRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetStorageInfoRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *GetStorageInfoRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetStorageInfoRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *GetStorageInfoRequest) GetResourceRealOwnerId() *int64 {
	return s.ResourceRealOwnerId
}

func (s *GetStorageInfoRequest) GetStorageLocation() *string {
	return s.StorageLocation
}

func (s *GetStorageInfoRequest) SetOwnerAccount(v string) *GetStorageInfoRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetStorageInfoRequest) SetOwnerId(v string) *GetStorageInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *GetStorageInfoRequest) SetResourceOwnerAccount(v string) *GetStorageInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetStorageInfoRequest) SetResourceOwnerId(v string) *GetStorageInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetStorageInfoRequest) SetResourceRealOwnerId(v int64) *GetStorageInfoRequest {
	s.ResourceRealOwnerId = &v
	return s
}

func (s *GetStorageInfoRequest) SetStorageLocation(v string) *GetStorageInfoRequest {
	s.StorageLocation = &v
	return s
}

func (s *GetStorageInfoRequest) Validate() error {
	return dara.Validate(s)
}
