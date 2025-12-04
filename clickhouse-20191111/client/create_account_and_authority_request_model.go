// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAccountAndAuthorityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountDescription(v string) *CreateAccountAndAuthorityRequest
	GetAccountDescription() *string
	SetAccountName(v string) *CreateAccountAndAuthorityRequest
	GetAccountName() *string
	SetAccountPassword(v string) *CreateAccountAndAuthorityRequest
	GetAccountPassword() *string
	SetAllowDatabases(v string) *CreateAccountAndAuthorityRequest
	GetAllowDatabases() *string
	SetAllowDictionaries(v string) *CreateAccountAndAuthorityRequest
	GetAllowDictionaries() *string
	SetDBClusterId(v string) *CreateAccountAndAuthorityRequest
	GetDBClusterId() *string
	SetDdlAuthority(v bool) *CreateAccountAndAuthorityRequest
	GetDdlAuthority() *bool
	SetDmlAuthority(v string) *CreateAccountAndAuthorityRequest
	GetDmlAuthority() *string
	SetOwnerAccount(v string) *CreateAccountAndAuthorityRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateAccountAndAuthorityRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateAccountAndAuthorityRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateAccountAndAuthorityRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateAccountAndAuthorityRequest
	GetResourceOwnerId() *int64
	SetTotalDatabases(v string) *CreateAccountAndAuthorityRequest
	GetTotalDatabases() *string
	SetTotalDictionaries(v string) *CreateAccountAndAuthorityRequest
	GetTotalDictionaries() *string
}

type CreateAccountAndAuthorityRequest struct {
	// The description of the database account.
	//
	// 	- The description cannot start with http:// or https://.
	//
	// 	- The description must be 0 to 256 characters in length.
	//
	// example:
	//
	// ceshi
	AccountDescription *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	// The database account.
	//
	// 	- The name must be unique within the cluster.
	//
	// 	- The name can contain lowercase letters, digits, and underscores (_).
	//
	// 	- The name must start with a lowercase letter and end with a lowercase letter or digit.
	//
	// 	- The name must be 2 to 64 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The password of your database account.
	//
	// >
	//
	// 	- The password contains at least three types of the following characters: uppercase letters, lowercase letters, digits, and special characters.
	//
	// 	- The password can contain the following special characters: ! @ # $ % ^ & \\	- ( ) _ + - =
	//
	// 	- The password must be 8 to 32 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456Aa
	AccountPassword *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	// The databases to which you want to grant permissions. Separate databases with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// db1
	AllowDatabases *string `json:"AllowDatabases,omitempty" xml:"AllowDatabases,omitempty"`
	// The dictionaries to which you want to grant permissions. Separate dictionaries with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// dt1
	AllowDictionaries *string `json:"AllowDictionaries,omitempty" xml:"AllowDictionaries,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp1p816075e21****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Specifies whether to grant the DDL permissions to the database account. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	DdlAuthority *bool `json:"DdlAuthority,omitempty" xml:"DdlAuthority,omitempty"`
	// Specifies whether to grant DML permissions to the database account. Valid values:
	//
	// 	- **all**
	//
	// 	- **readOnly,modify**
	//
	// This parameter is required.
	//
	// example:
	//
	// all
	DmlAuthority *string `json:"DmlAuthority,omitempty" xml:"DmlAuthority,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID You can call the [DescribeRegions](https://help.aliyun.com/document_detail/170875.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// All databases. Separate databases with commas (,).
	//
	// example:
	//
	// db1,db2
	TotalDatabases *string `json:"TotalDatabases,omitempty" xml:"TotalDatabases,omitempty"`
	// All dictionaries. Separate dictionaries with commas (,).
	//
	// example:
	//
	// dt1,dt2
	TotalDictionaries *string `json:"TotalDictionaries,omitempty" xml:"TotalDictionaries,omitempty"`
}

func (s CreateAccountAndAuthorityRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAccountAndAuthorityRequest) GoString() string {
	return s.String()
}

func (s *CreateAccountAndAuthorityRequest) GetAccountDescription() *string {
	return s.AccountDescription
}

func (s *CreateAccountAndAuthorityRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *CreateAccountAndAuthorityRequest) GetAccountPassword() *string {
	return s.AccountPassword
}

func (s *CreateAccountAndAuthorityRequest) GetAllowDatabases() *string {
	return s.AllowDatabases
}

func (s *CreateAccountAndAuthorityRequest) GetAllowDictionaries() *string {
	return s.AllowDictionaries
}

func (s *CreateAccountAndAuthorityRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateAccountAndAuthorityRequest) GetDdlAuthority() *bool {
	return s.DdlAuthority
}

func (s *CreateAccountAndAuthorityRequest) GetDmlAuthority() *string {
	return s.DmlAuthority
}

func (s *CreateAccountAndAuthorityRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateAccountAndAuthorityRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateAccountAndAuthorityRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateAccountAndAuthorityRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateAccountAndAuthorityRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateAccountAndAuthorityRequest) GetTotalDatabases() *string {
	return s.TotalDatabases
}

func (s *CreateAccountAndAuthorityRequest) GetTotalDictionaries() *string {
	return s.TotalDictionaries
}

func (s *CreateAccountAndAuthorityRequest) SetAccountDescription(v string) *CreateAccountAndAuthorityRequest {
	s.AccountDescription = &v
	return s
}

func (s *CreateAccountAndAuthorityRequest) SetAccountName(v string) *CreateAccountAndAuthorityRequest {
	s.AccountName = &v
	return s
}

func (s *CreateAccountAndAuthorityRequest) SetAccountPassword(v string) *CreateAccountAndAuthorityRequest {
	s.AccountPassword = &v
	return s
}

func (s *CreateAccountAndAuthorityRequest) SetAllowDatabases(v string) *CreateAccountAndAuthorityRequest {
	s.AllowDatabases = &v
	return s
}

func (s *CreateAccountAndAuthorityRequest) SetAllowDictionaries(v string) *CreateAccountAndAuthorityRequest {
	s.AllowDictionaries = &v
	return s
}

func (s *CreateAccountAndAuthorityRequest) SetDBClusterId(v string) *CreateAccountAndAuthorityRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateAccountAndAuthorityRequest) SetDdlAuthority(v bool) *CreateAccountAndAuthorityRequest {
	s.DdlAuthority = &v
	return s
}

func (s *CreateAccountAndAuthorityRequest) SetDmlAuthority(v string) *CreateAccountAndAuthorityRequest {
	s.DmlAuthority = &v
	return s
}

func (s *CreateAccountAndAuthorityRequest) SetOwnerAccount(v string) *CreateAccountAndAuthorityRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateAccountAndAuthorityRequest) SetOwnerId(v int64) *CreateAccountAndAuthorityRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateAccountAndAuthorityRequest) SetRegionId(v string) *CreateAccountAndAuthorityRequest {
	s.RegionId = &v
	return s
}

func (s *CreateAccountAndAuthorityRequest) SetResourceOwnerAccount(v string) *CreateAccountAndAuthorityRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateAccountAndAuthorityRequest) SetResourceOwnerId(v int64) *CreateAccountAndAuthorityRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateAccountAndAuthorityRequest) SetTotalDatabases(v string) *CreateAccountAndAuthorityRequest {
	s.TotalDatabases = &v
	return s
}

func (s *CreateAccountAndAuthorityRequest) SetTotalDictionaries(v string) *CreateAccountAndAuthorityRequest {
	s.TotalDictionaries = &v
	return s
}

func (s *CreateAccountAndAuthorityRequest) Validate() error {
	return dara.Validate(s)
}
