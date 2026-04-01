// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePostgresExtensionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeletePostgresExtensionsRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *DeletePostgresExtensionsRequest
	GetDBInstanceId() *string
	SetDBNames(v string) *DeletePostgresExtensionsRequest
	GetDBNames() *string
	SetExtensions(v string) *DeletePostgresExtensionsRequest
	GetExtensions() *string
	SetOwnerAccount(v string) *DeletePostgresExtensionsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeletePostgresExtensionsRequest
	GetOwnerId() *int64
	SetResourceGroupId(v string) *DeletePostgresExtensionsRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DeletePostgresExtensionsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeletePostgresExtensionsRequest
	GetResourceOwnerId() *int64
}

type DeletePostgresExtensionsRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pgm-bp156o9ti493****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The database on which the extension is installed. If you want to specify multiple databases, separate the databases with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// test_db
	DBNames *string `json:"DBNames,omitempty" xml:"DBNames,omitempty"`
	// The name of the extension. If you want to specify multiple extensions, separate the extension names with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// citext
	Extensions   *string `json:"Extensions,omitempty" xml:"Extensions,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmy****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeletePostgresExtensionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePostgresExtensionsRequest) GoString() string {
	return s.String()
}

func (s *DeletePostgresExtensionsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeletePostgresExtensionsRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DeletePostgresExtensionsRequest) GetDBNames() *string {
	return s.DBNames
}

func (s *DeletePostgresExtensionsRequest) GetExtensions() *string {
	return s.Extensions
}

func (s *DeletePostgresExtensionsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeletePostgresExtensionsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeletePostgresExtensionsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DeletePostgresExtensionsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeletePostgresExtensionsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeletePostgresExtensionsRequest) SetClientToken(v string) *DeletePostgresExtensionsRequest {
	s.ClientToken = &v
	return s
}

func (s *DeletePostgresExtensionsRequest) SetDBInstanceId(v string) *DeletePostgresExtensionsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeletePostgresExtensionsRequest) SetDBNames(v string) *DeletePostgresExtensionsRequest {
	s.DBNames = &v
	return s
}

func (s *DeletePostgresExtensionsRequest) SetExtensions(v string) *DeletePostgresExtensionsRequest {
	s.Extensions = &v
	return s
}

func (s *DeletePostgresExtensionsRequest) SetOwnerAccount(v string) *DeletePostgresExtensionsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeletePostgresExtensionsRequest) SetOwnerId(v int64) *DeletePostgresExtensionsRequest {
	s.OwnerId = &v
	return s
}

func (s *DeletePostgresExtensionsRequest) SetResourceGroupId(v string) *DeletePostgresExtensionsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DeletePostgresExtensionsRequest) SetResourceOwnerAccount(v string) *DeletePostgresExtensionsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeletePostgresExtensionsRequest) SetResourceOwnerId(v int64) *DeletePostgresExtensionsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeletePostgresExtensionsRequest) Validate() error {
	return dara.Validate(s)
}
