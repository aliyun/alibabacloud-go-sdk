// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePostgresExtensionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribePostgresExtensionsRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *DescribePostgresExtensionsRequest
	GetDBInstanceId() *string
	SetDBName(v string) *DescribePostgresExtensionsRequest
	GetDBName() *string
	SetOwnerAccount(v string) *DescribePostgresExtensionsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribePostgresExtensionsRequest
	GetOwnerId() *int64
	SetResourceGroupId(v string) *DescribePostgresExtensionsRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribePostgresExtensionsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribePostgresExtensionsRequest
	GetResourceOwnerId() *int64
}

type DescribePostgresExtensionsRequest struct {
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
	// The database name. You can call the DescribeDatabases operation to query the database name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_db
	DBName       *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
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

func (s DescribePostgresExtensionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePostgresExtensionsRequest) GoString() string {
	return s.String()
}

func (s *DescribePostgresExtensionsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribePostgresExtensionsRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribePostgresExtensionsRequest) GetDBName() *string {
	return s.DBName
}

func (s *DescribePostgresExtensionsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribePostgresExtensionsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribePostgresExtensionsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribePostgresExtensionsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribePostgresExtensionsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribePostgresExtensionsRequest) SetClientToken(v string) *DescribePostgresExtensionsRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribePostgresExtensionsRequest) SetDBInstanceId(v string) *DescribePostgresExtensionsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribePostgresExtensionsRequest) SetDBName(v string) *DescribePostgresExtensionsRequest {
	s.DBName = &v
	return s
}

func (s *DescribePostgresExtensionsRequest) SetOwnerAccount(v string) *DescribePostgresExtensionsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribePostgresExtensionsRequest) SetOwnerId(v int64) *DescribePostgresExtensionsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribePostgresExtensionsRequest) SetResourceGroupId(v string) *DescribePostgresExtensionsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribePostgresExtensionsRequest) SetResourceOwnerAccount(v string) *DescribePostgresExtensionsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribePostgresExtensionsRequest) SetResourceOwnerId(v int64) *DescribePostgresExtensionsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribePostgresExtensionsRequest) Validate() error {
	return dara.Validate(s)
}
