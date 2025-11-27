// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePostgresExtensionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *CreatePostgresExtensionsRequest
	GetAccountName() *string
	SetClientToken(v string) *CreatePostgresExtensionsRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *CreatePostgresExtensionsRequest
	GetDBInstanceId() *string
	SetDBNames(v string) *CreatePostgresExtensionsRequest
	GetDBNames() *string
	SetExtensions(v string) *CreatePostgresExtensionsRequest
	GetExtensions() *string
	SetOwnerAccount(v string) *CreatePostgresExtensionsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreatePostgresExtensionsRequest
	GetOwnerId() *int64
	SetResourceGroupId(v string) *CreatePostgresExtensionsRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreatePostgresExtensionsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreatePostgresExtensionsRequest
	GetResourceOwnerId() *int64
	SetRiskConfirmed(v bool) *CreatePostgresExtensionsRequest
	GetRiskConfirmed() *bool
	SetSourceDatabase(v string) *CreatePostgresExtensionsRequest
	GetSourceDatabase() *string
}

type CreatePostgresExtensionsRequest struct {
	// The account of the user who owns the extension. Only privileged accounts are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_user
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the generated token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
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
	// pgm-gc7f1****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The database name. You can call the DescribeDatabases operation to query the database name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_db
	DBNames *string `json:"DBNames,omitempty" xml:"DBNames,omitempty"`
	// The extension that you want to install. If you want to install multiple extensions, separate them with commas (,). If you do not specify the **SourceDatabase*	- parameter, you must specify this parameter.
	//
	// example:
	//
	// citext,pg_profile
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
	// The risk description that you need to confirm. If your instance runs an outdated minor engine version, installing specific extensions on the instance poses security risks. Proceed with the installation only after you acknowledge these risks. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// >  For more information about the risks, see [Limits on extension creation for ApsaraDB RDS for PostgreSQL instances](https://help.aliyun.com/document_detail/2587815.html).
	//
	// example:
	//
	// true
	RiskConfirmed *bool `json:"RiskConfirmed,omitempty" xml:"RiskConfirmed,omitempty"`
	// The source database from which you want to synchronize the extension to the destination database. If you do not specify the **Extensions*	- parameter, you must specify this parameter.
	//
	// example:
	//
	// source_db
	SourceDatabase *string `json:"SourceDatabase,omitempty" xml:"SourceDatabase,omitempty"`
}

func (s CreatePostgresExtensionsRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePostgresExtensionsRequest) GoString() string {
	return s.String()
}

func (s *CreatePostgresExtensionsRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *CreatePostgresExtensionsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreatePostgresExtensionsRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreatePostgresExtensionsRequest) GetDBNames() *string {
	return s.DBNames
}

func (s *CreatePostgresExtensionsRequest) GetExtensions() *string {
	return s.Extensions
}

func (s *CreatePostgresExtensionsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreatePostgresExtensionsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreatePostgresExtensionsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreatePostgresExtensionsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreatePostgresExtensionsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreatePostgresExtensionsRequest) GetRiskConfirmed() *bool {
	return s.RiskConfirmed
}

func (s *CreatePostgresExtensionsRequest) GetSourceDatabase() *string {
	return s.SourceDatabase
}

func (s *CreatePostgresExtensionsRequest) SetAccountName(v string) *CreatePostgresExtensionsRequest {
	s.AccountName = &v
	return s
}

func (s *CreatePostgresExtensionsRequest) SetClientToken(v string) *CreatePostgresExtensionsRequest {
	s.ClientToken = &v
	return s
}

func (s *CreatePostgresExtensionsRequest) SetDBInstanceId(v string) *CreatePostgresExtensionsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreatePostgresExtensionsRequest) SetDBNames(v string) *CreatePostgresExtensionsRequest {
	s.DBNames = &v
	return s
}

func (s *CreatePostgresExtensionsRequest) SetExtensions(v string) *CreatePostgresExtensionsRequest {
	s.Extensions = &v
	return s
}

func (s *CreatePostgresExtensionsRequest) SetOwnerAccount(v string) *CreatePostgresExtensionsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreatePostgresExtensionsRequest) SetOwnerId(v int64) *CreatePostgresExtensionsRequest {
	s.OwnerId = &v
	return s
}

func (s *CreatePostgresExtensionsRequest) SetResourceGroupId(v string) *CreatePostgresExtensionsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreatePostgresExtensionsRequest) SetResourceOwnerAccount(v string) *CreatePostgresExtensionsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreatePostgresExtensionsRequest) SetResourceOwnerId(v int64) *CreatePostgresExtensionsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreatePostgresExtensionsRequest) SetRiskConfirmed(v bool) *CreatePostgresExtensionsRequest {
	s.RiskConfirmed = &v
	return s
}

func (s *CreatePostgresExtensionsRequest) SetSourceDatabase(v string) *CreatePostgresExtensionsRequest {
	s.SourceDatabase = &v
	return s
}

func (s *CreatePostgresExtensionsRequest) Validate() error {
	return dara.Validate(s)
}
