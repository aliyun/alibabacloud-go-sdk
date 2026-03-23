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
	// This parameter is required.
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// This parameter is required.
	DBNames              *string `json:"DBNames,omitempty" xml:"DBNames,omitempty"`
	Extensions           *string `json:"Extensions,omitempty" xml:"Extensions,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RiskConfirmed        *bool   `json:"RiskConfirmed,omitempty" xml:"RiskConfirmed,omitempty"`
	SourceDatabase       *string `json:"SourceDatabase,omitempty" xml:"SourceDatabase,omitempty"`
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
