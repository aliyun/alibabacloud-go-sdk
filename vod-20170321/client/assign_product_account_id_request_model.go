// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssignProductAccountIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *AssignProductAccountIdRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *AssignProductAccountIdRequest
	GetOwnerId() *string
	SetResourceOwnerAccount(v string) *AssignProductAccountIdRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v string) *AssignProductAccountIdRequest
	GetResourceOwnerId() *string
	SetResourceRealOwnerId(v string) *AssignProductAccountIdRequest
	GetResourceRealOwnerId() *string
	SetStorageRegion(v string) *AssignProductAccountIdRequest
	GetStorageRegion() *string
}

type AssignProductAccountIdRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceRealOwnerId  *string `json:"ResourceRealOwnerId,omitempty" xml:"ResourceRealOwnerId,omitempty"`
	StorageRegion        *string `json:"StorageRegion,omitempty" xml:"StorageRegion,omitempty"`
}

func (s AssignProductAccountIdRequest) String() string {
	return dara.Prettify(s)
}

func (s AssignProductAccountIdRequest) GoString() string {
	return s.String()
}

func (s *AssignProductAccountIdRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AssignProductAccountIdRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *AssignProductAccountIdRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AssignProductAccountIdRequest) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *AssignProductAccountIdRequest) GetResourceRealOwnerId() *string {
	return s.ResourceRealOwnerId
}

func (s *AssignProductAccountIdRequest) GetStorageRegion() *string {
	return s.StorageRegion
}

func (s *AssignProductAccountIdRequest) SetOwnerAccount(v string) *AssignProductAccountIdRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AssignProductAccountIdRequest) SetOwnerId(v string) *AssignProductAccountIdRequest {
	s.OwnerId = &v
	return s
}

func (s *AssignProductAccountIdRequest) SetResourceOwnerAccount(v string) *AssignProductAccountIdRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AssignProductAccountIdRequest) SetResourceOwnerId(v string) *AssignProductAccountIdRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AssignProductAccountIdRequest) SetResourceRealOwnerId(v string) *AssignProductAccountIdRequest {
	s.ResourceRealOwnerId = &v
	return s
}

func (s *AssignProductAccountIdRequest) SetStorageRegion(v string) *AssignProductAccountIdRequest {
	s.StorageRegion = &v
	return s
}

func (s *AssignProductAccountIdRequest) Validate() error {
	return dara.Validate(s)
}
