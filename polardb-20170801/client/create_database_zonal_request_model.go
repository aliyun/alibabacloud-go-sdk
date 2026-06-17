// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDatabaseZonalRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *CreateDatabaseZonalRequest
	GetAccountName() *string
	SetAccountPrivilege(v string) *CreateDatabaseZonalRequest
	GetAccountPrivilege() *string
	SetCharacterSetName(v string) *CreateDatabaseZonalRequest
	GetCharacterSetName() *string
	SetClientToken(v string) *CreateDatabaseZonalRequest
	GetClientToken() *string
	SetCollate(v string) *CreateDatabaseZonalRequest
	GetCollate() *string
	SetCtype(v string) *CreateDatabaseZonalRequest
	GetCtype() *string
	SetDBClusterId(v string) *CreateDatabaseZonalRequest
	GetDBClusterId() *string
	SetDBDescription(v string) *CreateDatabaseZonalRequest
	GetDBDescription() *string
	SetDBName(v string) *CreateDatabaseZonalRequest
	GetDBName() *string
	SetOwnerAccount(v string) *CreateDatabaseZonalRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateDatabaseZonalRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CreateDatabaseZonalRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateDatabaseZonalRequest
	GetResourceOwnerId() *int64
}

type CreateDatabaseZonalRequest struct {
	// The name of the account that is authorized to access the database.
	//
	// example:
	//
	// testacc
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The permissions of the account. Valid values:
	//
	// - ReadWrite: read and write permissions.
	//
	// - ReadOnly: read-only permissions.
	//
	// - DMLOnly: DML permissions only.
	//
	// - DDLOnly: DDL permissions only.
	//
	// - ReadIndex: read-only and index permissions.
	//
	// If you do not specify this parameter, the default value is ReadWrite.
	//
	// example:
	//
	// ReadWrite
	AccountPrivilege *string `json:"AccountPrivilege,omitempty" xml:"AccountPrivilege,omitempty"`
	// The character set.
	//
	// This parameter is required.
	//
	// example:
	//
	// utf8
	CharacterSetName *string `json:"CharacterSetName,omitempty" xml:"CharacterSetName,omitempty"`
	// A client token to ensure request idempotence. The client generates this token. The token must be unique across requests. It is case-sensitive and can be up to 64 ASCII characters long.
	//
	// example:
	//
	// 6000170000591aed949d0f54a343f1a4233c1e7d1c5c******
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The locale setting. This specifies the collation for the new database.
	//
	// example:
	//
	// C
	Collate *string `json:"Collate,omitempty" xml:"Collate,omitempty"`
	// The locale setting. This specifies the character classification for the database.
	//
	// example:
	//
	// C
	Ctype *string `json:"Ctype,omitempty" xml:"Ctype,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-**************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The description of the database. The description must meet the following requirements:
	//
	// - It cannot start with `http://` or `https://`.
	//
	// - It must be 2 to 256 characters in length.
	//
	// example:
	//
	// testdesc
	DBDescription *string `json:"DBDescription,omitempty" xml:"DBDescription,omitempty"`
	// The name of the database. The name must meet the following requirements:
	//
	// - It must consist of lowercase letters, digits, hyphens (-), and underscores (_).
	//
	// - It must start with a letter and end with a letter or a digit. The name can be up to 64 characters long.
	//
	// This parameter is required.
	//
	// example:
	//
	// testDB
	DBName               *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateDatabaseZonalRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDatabaseZonalRequest) GoString() string {
	return s.String()
}

func (s *CreateDatabaseZonalRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *CreateDatabaseZonalRequest) GetAccountPrivilege() *string {
	return s.AccountPrivilege
}

func (s *CreateDatabaseZonalRequest) GetCharacterSetName() *string {
	return s.CharacterSetName
}

func (s *CreateDatabaseZonalRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateDatabaseZonalRequest) GetCollate() *string {
	return s.Collate
}

func (s *CreateDatabaseZonalRequest) GetCtype() *string {
	return s.Ctype
}

func (s *CreateDatabaseZonalRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateDatabaseZonalRequest) GetDBDescription() *string {
	return s.DBDescription
}

func (s *CreateDatabaseZonalRequest) GetDBName() *string {
	return s.DBName
}

func (s *CreateDatabaseZonalRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateDatabaseZonalRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateDatabaseZonalRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateDatabaseZonalRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateDatabaseZonalRequest) SetAccountName(v string) *CreateDatabaseZonalRequest {
	s.AccountName = &v
	return s
}

func (s *CreateDatabaseZonalRequest) SetAccountPrivilege(v string) *CreateDatabaseZonalRequest {
	s.AccountPrivilege = &v
	return s
}

func (s *CreateDatabaseZonalRequest) SetCharacterSetName(v string) *CreateDatabaseZonalRequest {
	s.CharacterSetName = &v
	return s
}

func (s *CreateDatabaseZonalRequest) SetClientToken(v string) *CreateDatabaseZonalRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDatabaseZonalRequest) SetCollate(v string) *CreateDatabaseZonalRequest {
	s.Collate = &v
	return s
}

func (s *CreateDatabaseZonalRequest) SetCtype(v string) *CreateDatabaseZonalRequest {
	s.Ctype = &v
	return s
}

func (s *CreateDatabaseZonalRequest) SetDBClusterId(v string) *CreateDatabaseZonalRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateDatabaseZonalRequest) SetDBDescription(v string) *CreateDatabaseZonalRequest {
	s.DBDescription = &v
	return s
}

func (s *CreateDatabaseZonalRequest) SetDBName(v string) *CreateDatabaseZonalRequest {
	s.DBName = &v
	return s
}

func (s *CreateDatabaseZonalRequest) SetOwnerAccount(v string) *CreateDatabaseZonalRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateDatabaseZonalRequest) SetOwnerId(v int64) *CreateDatabaseZonalRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDatabaseZonalRequest) SetResourceOwnerAccount(v string) *CreateDatabaseZonalRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateDatabaseZonalRequest) SetResourceOwnerId(v int64) *CreateDatabaseZonalRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateDatabaseZonalRequest) Validate() error {
	return dara.Validate(s)
}
