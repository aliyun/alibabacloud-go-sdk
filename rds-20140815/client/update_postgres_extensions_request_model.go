// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePostgresExtensionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdatePostgresExtensionsRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *UpdatePostgresExtensionsRequest
	GetDBInstanceId() *string
	SetDBNames(v string) *UpdatePostgresExtensionsRequest
	GetDBNames() *string
	SetExtensions(v string) *UpdatePostgresExtensionsRequest
	GetExtensions() *string
	SetOwnerAccount(v string) *UpdatePostgresExtensionsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UpdatePostgresExtensionsRequest
	GetOwnerId() *int64
	SetResourceGroupId(v string) *UpdatePostgresExtensionsRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *UpdatePostgresExtensionsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdatePostgresExtensionsRequest
	GetResourceOwnerId() *int64
}

type UpdatePostgresExtensionsRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// This parameter is required.
	DBNames *string `json:"DBNames,omitempty" xml:"DBNames,omitempty"`
	// This parameter is required.
	Extensions           *string `json:"Extensions,omitempty" xml:"Extensions,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s UpdatePostgresExtensionsRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePostgresExtensionsRequest) GoString() string {
	return s.String()
}

func (s *UpdatePostgresExtensionsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdatePostgresExtensionsRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *UpdatePostgresExtensionsRequest) GetDBNames() *string {
	return s.DBNames
}

func (s *UpdatePostgresExtensionsRequest) GetExtensions() *string {
	return s.Extensions
}

func (s *UpdatePostgresExtensionsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UpdatePostgresExtensionsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdatePostgresExtensionsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *UpdatePostgresExtensionsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdatePostgresExtensionsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdatePostgresExtensionsRequest) SetClientToken(v string) *UpdatePostgresExtensionsRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdatePostgresExtensionsRequest) SetDBInstanceId(v string) *UpdatePostgresExtensionsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *UpdatePostgresExtensionsRequest) SetDBNames(v string) *UpdatePostgresExtensionsRequest {
	s.DBNames = &v
	return s
}

func (s *UpdatePostgresExtensionsRequest) SetExtensions(v string) *UpdatePostgresExtensionsRequest {
	s.Extensions = &v
	return s
}

func (s *UpdatePostgresExtensionsRequest) SetOwnerAccount(v string) *UpdatePostgresExtensionsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdatePostgresExtensionsRequest) SetOwnerId(v int64) *UpdatePostgresExtensionsRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdatePostgresExtensionsRequest) SetResourceGroupId(v string) *UpdatePostgresExtensionsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdatePostgresExtensionsRequest) SetResourceOwnerAccount(v string) *UpdatePostgresExtensionsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdatePostgresExtensionsRequest) SetResourceOwnerId(v int64) *UpdatePostgresExtensionsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdatePostgresExtensionsRequest) Validate() error {
	return dara.Validate(s)
}
