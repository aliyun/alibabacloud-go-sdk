// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAccountAuthorityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *ModifyAccountAuthorityRequest
	GetAccountName() *string
	SetAllowDatabases(v string) *ModifyAccountAuthorityRequest
	GetAllowDatabases() *string
	SetAllowDictionaries(v string) *ModifyAccountAuthorityRequest
	GetAllowDictionaries() *string
	SetDBClusterId(v string) *ModifyAccountAuthorityRequest
	GetDBClusterId() *string
	SetDdlAuthority(v bool) *ModifyAccountAuthorityRequest
	GetDdlAuthority() *bool
	SetDmlAuthority(v string) *ModifyAccountAuthorityRequest
	GetDmlAuthority() *string
	SetOwnerAccount(v string) *ModifyAccountAuthorityRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyAccountAuthorityRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyAccountAuthorityRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyAccountAuthorityRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyAccountAuthorityRequest
	GetResourceOwnerId() *int64
	SetTotalDatabases(v string) *ModifyAccountAuthorityRequest
	GetTotalDatabases() *string
	SetTotalDictionaries(v string) *ModifyAccountAuthorityRequest
	GetTotalDictionaries() *string
}

type ModifyAccountAuthorityRequest struct {
	// The name of the database account.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
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
	// cc-bp108z124a8o7****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Specifies whether to grant DDL permissions to the database account. Valid values:
	//
	// 	- **true**: grants DDL permissions to the database account.
	//
	// 	- **false**: does not grant DDL permissions to the database account.
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
	// 	- **readonly,modify**
	//
	// This parameter is required.
	//
	// example:
	//
	// all
	DmlAuthority *string `json:"DmlAuthority,omitempty" xml:"DmlAuthority,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/170875.html) operation to query the most recent region list.
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

func (s ModifyAccountAuthorityRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccountAuthorityRequest) GoString() string {
	return s.String()
}

func (s *ModifyAccountAuthorityRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *ModifyAccountAuthorityRequest) GetAllowDatabases() *string {
	return s.AllowDatabases
}

func (s *ModifyAccountAuthorityRequest) GetAllowDictionaries() *string {
	return s.AllowDictionaries
}

func (s *ModifyAccountAuthorityRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyAccountAuthorityRequest) GetDdlAuthority() *bool {
	return s.DdlAuthority
}

func (s *ModifyAccountAuthorityRequest) GetDmlAuthority() *string {
	return s.DmlAuthority
}

func (s *ModifyAccountAuthorityRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyAccountAuthorityRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyAccountAuthorityRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyAccountAuthorityRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyAccountAuthorityRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyAccountAuthorityRequest) GetTotalDatabases() *string {
	return s.TotalDatabases
}

func (s *ModifyAccountAuthorityRequest) GetTotalDictionaries() *string {
	return s.TotalDictionaries
}

func (s *ModifyAccountAuthorityRequest) SetAccountName(v string) *ModifyAccountAuthorityRequest {
	s.AccountName = &v
	return s
}

func (s *ModifyAccountAuthorityRequest) SetAllowDatabases(v string) *ModifyAccountAuthorityRequest {
	s.AllowDatabases = &v
	return s
}

func (s *ModifyAccountAuthorityRequest) SetAllowDictionaries(v string) *ModifyAccountAuthorityRequest {
	s.AllowDictionaries = &v
	return s
}

func (s *ModifyAccountAuthorityRequest) SetDBClusterId(v string) *ModifyAccountAuthorityRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyAccountAuthorityRequest) SetDdlAuthority(v bool) *ModifyAccountAuthorityRequest {
	s.DdlAuthority = &v
	return s
}

func (s *ModifyAccountAuthorityRequest) SetDmlAuthority(v string) *ModifyAccountAuthorityRequest {
	s.DmlAuthority = &v
	return s
}

func (s *ModifyAccountAuthorityRequest) SetOwnerAccount(v string) *ModifyAccountAuthorityRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyAccountAuthorityRequest) SetOwnerId(v int64) *ModifyAccountAuthorityRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyAccountAuthorityRequest) SetRegionId(v string) *ModifyAccountAuthorityRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyAccountAuthorityRequest) SetResourceOwnerAccount(v string) *ModifyAccountAuthorityRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyAccountAuthorityRequest) SetResourceOwnerId(v int64) *ModifyAccountAuthorityRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyAccountAuthorityRequest) SetTotalDatabases(v string) *ModifyAccountAuthorityRequest {
	s.TotalDatabases = &v
	return s
}

func (s *ModifyAccountAuthorityRequest) SetTotalDictionaries(v string) *ModifyAccountAuthorityRequest {
	s.TotalDictionaries = &v
	return s
}

func (s *ModifyAccountAuthorityRequest) Validate() error {
	return dara.Validate(s)
}
