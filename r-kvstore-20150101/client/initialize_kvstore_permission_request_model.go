// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitializeKvstorePermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *InitializeKvstorePermissionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *InitializeKvstorePermissionRequest
	GetOwnerId() *int64
	SetRegionId(v string) *InitializeKvstorePermissionRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *InitializeKvstorePermissionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *InitializeKvstorePermissionRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *InitializeKvstorePermissionRequest
	GetSecurityToken() *string
}

type InitializeKvstorePermissionRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s InitializeKvstorePermissionRequest) String() string {
	return dara.Prettify(s)
}

func (s InitializeKvstorePermissionRequest) GoString() string {
	return s.String()
}

func (s *InitializeKvstorePermissionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *InitializeKvstorePermissionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *InitializeKvstorePermissionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *InitializeKvstorePermissionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *InitializeKvstorePermissionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *InitializeKvstorePermissionRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *InitializeKvstorePermissionRequest) SetOwnerAccount(v string) *InitializeKvstorePermissionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *InitializeKvstorePermissionRequest) SetOwnerId(v int64) *InitializeKvstorePermissionRequest {
	s.OwnerId = &v
	return s
}

func (s *InitializeKvstorePermissionRequest) SetRegionId(v string) *InitializeKvstorePermissionRequest {
	s.RegionId = &v
	return s
}

func (s *InitializeKvstorePermissionRequest) SetResourceOwnerAccount(v string) *InitializeKvstorePermissionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *InitializeKvstorePermissionRequest) SetResourceOwnerId(v int64) *InitializeKvstorePermissionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *InitializeKvstorePermissionRequest) SetSecurityToken(v string) *InitializeKvstorePermissionRequest {
	s.SecurityToken = &v
	return s
}

func (s *InitializeKvstorePermissionRequest) Validate() error {
	return dara.Validate(s)
}
