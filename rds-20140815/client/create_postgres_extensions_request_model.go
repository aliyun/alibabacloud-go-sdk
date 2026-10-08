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
	// The user to which the extension belongs. Only privileged accounts are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_user
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID. You can call DescribeDBInstances to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pgm-gc7f1****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The database name of the instance. You can call DescribeDatabases to query the database name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_db
	DBNames *string `json:"DBNames,omitempty" xml:"DBNames,omitempty"`
	// The plugins to install. Separate multiple plugins with commas (,).
	//
	// If you do not specify the request parameter **SourceDatabase**, this parameter is required.
	//
	// example:
	//
	// citext,pg_profile
	Extensions   *string `json:"Extensions,omitempty" xml:"Extensions,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmy****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specifies whether to confirm the security risk of installing specific extensions on instances that run minor engine versions that are too early. After you confirm the risk, the extensions can be installed.
	//
	// Valid values:
	//
	// - true
	//
	// - false
	//
	// > For information about related risks, see [Restrictions on creating extensions in ApsaraDB RDS for PostgreSQL](https://help.aliyun.com/document_detail/2587815.html).
	//
	// example:
	//
	// true
	RiskConfirmed *bool `json:"RiskConfirmed,omitempty" xml:"RiskConfirmed,omitempty"`
	// The source database from which plugins are synchronized to the target database. If you do not specify the request parameter **Extensions**, this parameter is required.
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
