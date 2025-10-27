// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeOperatorPermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *RevokeOperatorPermissionRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *RevokeOperatorPermissionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RevokeOperatorPermissionRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *RevokeOperatorPermissionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RevokeOperatorPermissionRequest
	GetResourceOwnerId() *int64
}

type RevokeOperatorPermissionRequest struct {
	// The ID of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxxxxx
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s RevokeOperatorPermissionRequest) String() string {
	return dara.Prettify(s)
}

func (s RevokeOperatorPermissionRequest) GoString() string {
	return s.String()
}

func (s *RevokeOperatorPermissionRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *RevokeOperatorPermissionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RevokeOperatorPermissionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RevokeOperatorPermissionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RevokeOperatorPermissionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RevokeOperatorPermissionRequest) SetDBClusterId(v string) *RevokeOperatorPermissionRequest {
	s.DBClusterId = &v
	return s
}

func (s *RevokeOperatorPermissionRequest) SetOwnerAccount(v string) *RevokeOperatorPermissionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RevokeOperatorPermissionRequest) SetOwnerId(v int64) *RevokeOperatorPermissionRequest {
	s.OwnerId = &v
	return s
}

func (s *RevokeOperatorPermissionRequest) SetResourceOwnerAccount(v string) *RevokeOperatorPermissionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RevokeOperatorPermissionRequest) SetResourceOwnerId(v int64) *RevokeOperatorPermissionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RevokeOperatorPermissionRequest) Validate() error {
	return dara.Validate(s)
}
