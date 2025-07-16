// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *CreateDBRequest
	GetAccountName() *string
	SetAccountPrivilege(v string) *CreateDBRequest
	GetAccountPrivilege() *string
	SetCharset(v string) *CreateDBRequest
	GetCharset() *string
	SetDBInstanceName(v string) *CreateDBRequest
	GetDBInstanceName() *string
	SetDbDescription(v string) *CreateDBRequest
	GetDbDescription() *string
	SetDbName(v string) *CreateDBRequest
	GetDbName() *string
	SetMode(v string) *CreateDBRequest
	GetMode() *string
	SetRegionId(v string) *CreateDBRequest
	GetRegionId() *string
	SetSecurityAccountName(v string) *CreateDBRequest
	GetSecurityAccountName() *string
	SetSecurityAccountPassword(v string) *CreateDBRequest
	GetSecurityAccountPassword() *string
	SetStoragePoolName(v string) *CreateDBRequest
	GetStoragePoolName() *string
}

type CreateDBRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// testaccount
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// example:
	//
	// ReadWrite
	AccountPrivilege *string `json:"AccountPrivilege,omitempty" xml:"AccountPrivilege,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// utf8mb4
	Charset *string `json:"Charset,omitempty" xml:"Charset,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pxc-**************
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// db for test
	DbDescription *string `json:"DbDescription,omitempty" xml:"DbDescription,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testdb
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// example:
	//
	// auto
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// securityAccount
	SecurityAccountName *string `json:"SecurityAccountName,omitempty" xml:"SecurityAccountName,omitempty"`
	// example:
	//
	// securityPassword
	SecurityAccountPassword *string `json:"SecurityAccountPassword,omitempty" xml:"SecurityAccountPassword,omitempty"`
	StoragePoolName         *string `json:"StoragePoolName,omitempty" xml:"StoragePoolName,omitempty"`
}

func (s CreateDBRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDBRequest) GoString() string {
	return s.String()
}

func (s *CreateDBRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *CreateDBRequest) GetAccountPrivilege() *string {
	return s.AccountPrivilege
}

func (s *CreateDBRequest) GetCharset() *string {
	return s.Charset
}

func (s *CreateDBRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *CreateDBRequest) GetDbDescription() *string {
	return s.DbDescription
}

func (s *CreateDBRequest) GetDbName() *string {
	return s.DbName
}

func (s *CreateDBRequest) GetMode() *string {
	return s.Mode
}

func (s *CreateDBRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDBRequest) GetSecurityAccountName() *string {
	return s.SecurityAccountName
}

func (s *CreateDBRequest) GetSecurityAccountPassword() *string {
	return s.SecurityAccountPassword
}

func (s *CreateDBRequest) GetStoragePoolName() *string {
	return s.StoragePoolName
}

func (s *CreateDBRequest) SetAccountName(v string) *CreateDBRequest {
	s.AccountName = &v
	return s
}

func (s *CreateDBRequest) SetAccountPrivilege(v string) *CreateDBRequest {
	s.AccountPrivilege = &v
	return s
}

func (s *CreateDBRequest) SetCharset(v string) *CreateDBRequest {
	s.Charset = &v
	return s
}

func (s *CreateDBRequest) SetDBInstanceName(v string) *CreateDBRequest {
	s.DBInstanceName = &v
	return s
}

func (s *CreateDBRequest) SetDbDescription(v string) *CreateDBRequest {
	s.DbDescription = &v
	return s
}

func (s *CreateDBRequest) SetDbName(v string) *CreateDBRequest {
	s.DbName = &v
	return s
}

func (s *CreateDBRequest) SetMode(v string) *CreateDBRequest {
	s.Mode = &v
	return s
}

func (s *CreateDBRequest) SetRegionId(v string) *CreateDBRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDBRequest) SetSecurityAccountName(v string) *CreateDBRequest {
	s.SecurityAccountName = &v
	return s
}

func (s *CreateDBRequest) SetSecurityAccountPassword(v string) *CreateDBRequest {
	s.SecurityAccountPassword = &v
	return s
}

func (s *CreateDBRequest) SetStoragePoolName(v string) *CreateDBRequest {
	s.StoragePoolName = &v
	return s
}

func (s *CreateDBRequest) Validate() error {
	return dara.Validate(s)
}
