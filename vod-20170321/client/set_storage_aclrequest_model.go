// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetStorageACLRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *SetStorageACLRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *SetStorageACLRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *SetStorageACLRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *SetStorageACLRequest
	GetResourceOwnerId() *string
	SetResourceRealOwnerId(v int64) *SetStorageACLRequest
	GetResourceRealOwnerId() *int64
	SetStorageACL(v string) *SetStorageACLRequest
	GetStorageACL() *string
	SetStorageLocation(v string) *SetStorageACLRequest
	GetStorageLocation() *string
}

type SetStorageACLRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceRealOwnerId  *int64  `json:"ResourceRealOwnerId,omitempty" xml:"ResourceRealOwnerId,omitempty"`
	// This parameter is required.
	StorageACL *string `json:"StorageACL,omitempty" xml:"StorageACL,omitempty"`
	// This parameter is required.
	StorageLocation *string `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
}

func (s SetStorageACLRequest) String() string {
	return dara.Prettify(s)
}

func (s SetStorageACLRequest) GoString() string {
	return s.String()
}

func (s *SetStorageACLRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SetStorageACLRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *SetStorageACLRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SetStorageACLRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *SetStorageACLRequest) GetResourceRealOwnerId() *int64 {
	return s.ResourceRealOwnerId
}

func (s *SetStorageACLRequest) GetStorageACL() *string {
	return s.StorageACL
}

func (s *SetStorageACLRequest) GetStorageLocation() *string {
	return s.StorageLocation
}

func (s *SetStorageACLRequest) SetOwnerAccount(v string) *SetStorageACLRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SetStorageACLRequest) SetOwnerId(v string) *SetStorageACLRequest {
	s.OwnerId = &v
	return s
}

func (s *SetStorageACLRequest) SetResourceOwnerAccount(v string) *SetStorageACLRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SetStorageACLRequest) SetResourceOwnerId(v string) *SetStorageACLRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SetStorageACLRequest) SetResourceRealOwnerId(v int64) *SetStorageACLRequest {
	s.ResourceRealOwnerId = &v
	return s
}

func (s *SetStorageACLRequest) SetStorageACL(v string) *SetStorageACLRequest {
	s.StorageACL = &v
	return s
}

func (s *SetStorageACLRequest) SetStorageLocation(v string) *SetStorageACLRequest {
	s.StorageLocation = &v
	return s
}

func (s *SetStorageACLRequest) Validate() error {
	return dara.Validate(s)
}
