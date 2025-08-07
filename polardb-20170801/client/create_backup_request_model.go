// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBackupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateBackupRequest
	GetClientToken() *string
	SetDBClusterId(v string) *CreateBackupRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *CreateBackupRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateBackupRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CreateBackupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateBackupRequest
	GetResourceOwnerId() *int64
}

type CreateBackupRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that it is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. The token is case-sensitive.
	//
	// example:
	//
	// 6000170000591aed949d0f54a343f1a4233c1e7d1c5c******
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-*************
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateBackupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBackupRequest) GoString() string {
	return s.String()
}

func (s *CreateBackupRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateBackupRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateBackupRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateBackupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateBackupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateBackupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateBackupRequest) SetClientToken(v string) *CreateBackupRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateBackupRequest) SetDBClusterId(v string) *CreateBackupRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateBackupRequest) SetOwnerAccount(v string) *CreateBackupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateBackupRequest) SetOwnerId(v int64) *CreateBackupRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateBackupRequest) SetResourceOwnerAccount(v string) *CreateBackupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateBackupRequest) SetResourceOwnerId(v int64) *CreateBackupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateBackupRequest) Validate() error {
	return dara.Validate(s)
}
