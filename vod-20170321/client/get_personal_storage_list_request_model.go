// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPersonalStorageListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxKeys(v string) *GetPersonalStorageListRequest
	GetMaxKeys() *string
	SetOwnerAccount(v string) *GetPersonalStorageListRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *GetPersonalStorageListRequest
	GetOwnerId() *string
	SetPrefix(v string) *GetPersonalStorageListRequest
	GetPrefix() *string
	SetResourceOwnerAccount(v string) *GetPersonalStorageListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *GetPersonalStorageListRequest
	GetResourceOwnerId() *string
	SetResourceRealOwnerId(v int64) *GetPersonalStorageListRequest
	GetResourceRealOwnerId() *int64
	SetStorageRegion(v string) *GetPersonalStorageListRequest
	GetStorageRegion() *string
}

type GetPersonalStorageListRequest struct {
	MaxKeys              *string `json:"MaxKeys,omitempty" xml:"MaxKeys,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Prefix               *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceRealOwnerId  *int64  `json:"ResourceRealOwnerId,omitempty" xml:"ResourceRealOwnerId,omitempty"`
	StorageRegion        *string `json:"StorageRegion,omitempty" xml:"StorageRegion,omitempty"`
}

func (s GetPersonalStorageListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPersonalStorageListRequest) GoString() string {
	return s.String()
}

func (s *GetPersonalStorageListRequest) GetMaxKeys() *string {
	return s.MaxKeys
}

func (s *GetPersonalStorageListRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetPersonalStorageListRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *GetPersonalStorageListRequest) GetPrefix() *string {
	return s.Prefix
}

func (s *GetPersonalStorageListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetPersonalStorageListRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *GetPersonalStorageListRequest) GetResourceRealOwnerId() *int64 {
	return s.ResourceRealOwnerId
}

func (s *GetPersonalStorageListRequest) GetStorageRegion() *string {
	return s.StorageRegion
}

func (s *GetPersonalStorageListRequest) SetMaxKeys(v string) *GetPersonalStorageListRequest {
	s.MaxKeys = &v
	return s
}

func (s *GetPersonalStorageListRequest) SetOwnerAccount(v string) *GetPersonalStorageListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetPersonalStorageListRequest) SetOwnerId(v string) *GetPersonalStorageListRequest {
	s.OwnerId = &v
	return s
}

func (s *GetPersonalStorageListRequest) SetPrefix(v string) *GetPersonalStorageListRequest {
	s.Prefix = &v
	return s
}

func (s *GetPersonalStorageListRequest) SetResourceOwnerAccount(v string) *GetPersonalStorageListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetPersonalStorageListRequest) SetResourceOwnerId(v string) *GetPersonalStorageListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetPersonalStorageListRequest) SetResourceRealOwnerId(v int64) *GetPersonalStorageListRequest {
	s.ResourceRealOwnerId = &v
	return s
}

func (s *GetPersonalStorageListRequest) SetStorageRegion(v string) *GetPersonalStorageListRequest {
	s.StorageRegion = &v
	return s
}

func (s *GetPersonalStorageListRequest) Validate() error {
	return dara.Validate(s)
}
